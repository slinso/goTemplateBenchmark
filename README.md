# goTemplateBenchmark
comparing the performance of different template engines
* [Ace](https://github.com/yosssi/ace)
* [Amber](https://github.com/eknkc/amber)
* [Damsel](https://github.com/dskinner/damsel)
* [ego](https://github.com/benbjohnson/ego)
* [egon](https://github.com/commondream/egon)
* [egonslinso](https://github.com/SlinSo/egon)
* [ftmpl](https://github.com/tkrajina/ftmpl)
* [Go](https://golang.org/pkg/html/template)
* [Gorazor](https://github.com/sipin/gorazor)
* [Handlebars](https://github.com/aymerick/raymond)
* [Kasia](https://github.com/ziutek/kasia.go)
* [Mustache](https://github.com/hoisie/mustache)
* [Pongo2](https://github.com/flosch/pongo2)
* [Soy](https://github.com/robfig/soy)

## Why?
first thought:
Just for fun. Go Templates work nice out of the box.
If you really care about performance you will usually cache the rendered output.

on second thought:
I have some templates that I cannot cache in my production code, thats why I'm interested in performant
HTML generation using templates. After trying the code generation based projects I liked ego most, but some
features where missing and generated code could be optimized further. That's why I created a fork
and included the results in this benchmark.

## Results
Tests run on a VPS 1 CPU und 512 MB Ram

### normal Template Engines
```
go test -bench "k[Ace|Amber|Damsel|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy]" -benchmem -benchtime=3s
PASS
BenchmarkGolang           100000             34999 ns/op            2078 B/op         38 allocs/op
BenchmarkAce               50000             78572 ns/op            5549 B/op         77 allocs/op
BenchmarkAmber            100000             37048 ns/op            2090 B/op         39 allocs/op
BenchmarkDamsel           100000             54046 ns/op            2440 B/op         50 allocs/op
BenchmarkMustache         200000             19565 ns/op            1648 B/op         28 allocs/op
BenchmarkPongo2           200000             23062 ns/op            2997 B/op         46 allocs/op
BenchmarkHandlebars       100000             65334 ns/op            4496 B/op         90 allocs/op
BenchmarkKasia            300000             15852 ns/op            2028 B/op         26 allocs/op
BenchmarkSoy              200000             19451 ns/op            1732 B/op         26 allocs/op
```

### Template Engines with manual precompilation
```
go test -bench "kEgo$|kEgon$|kEgonSlinso$|kFtmpl|kGorazor" -benchmem -benchtime=3s
PASS
BenchmarkEgo             1000000              5017 ns/op             645 B/op          8 allocs/op
BenchmarkEgon             500000             10755 ns/op             870 B/op         22 allocs/op
BenchmarkEgonSlinso      2000000              2619 ns/op             517 B/op          0 allocs/op
BenchmarkFtmpl           1000000              8400 ns/op            1152 B/op         12 allocs/op
BenchmarkGorazor          500000              7249 ns/op             656 B/op         11 allocs/op
```

### more complex test with template inheritance (if possible)
```
go test . -bench="Complex" -benchmem -benchtime=5s
PASS
BenchmarkComplexGolang             30000            276809 ns/op           12532 B/op        295 allocs/op
BenchmarkComplexEgo               300000             27030 ns/op            4039 B/op         41 allocs/op
BenchmarkComplexEgon              200000             50424 ns/op            4206 B/op        101 allocs/op
BenchmarkComplexEgoSlinso         500000             13977 ns/op            2145 B/op          7 allocs/op
BenchmarkComplexFtmpl             200000             38427 ns/op            5201 B/op         40 allocs/op
BenchmarkComplexFtmplFctCall      200000             40946 ns/op            5745 B/op         48 allocs/op
BenchmarkComplexMustache           50000            132716 ns/op            8451 B/op        166 allocs/op
BenchmarkComplexGorazor           100000             61630 ns/op            8577 B/op         73 allocs/op
```

## Security
All packages assume that template authors are trusted. If you allow custom templates you have to sanitize your user input e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday). Generally speaking I would suggest to sanitize every input not just HTML-input. 

| Framework | Security | Comment |
| --------- | -------- | ------- |
| Ace | No | |
| amber | No | |
| Damsel | No | |
| ego | Partial (html.EscapeString) | only HTML, others need to be called manually |
| egon | Partial (html.EscapeString) | only HTML, others need to be called manually |
| egonslinso | Partial (html.EscapeString) | only HTML, others need to be called manually |
| ftmpl | Partial (html.EscapeString) | only HTML, others need to be called manually |
| Go | Yes | contextual escaping [html/template Security Model](https://golang.org/pkg/html/template/#hdr-Security_Model) |
| Gorazor | Partial (template.HTMLEscapeString) | only HTML, others need to be called manually |
| Handlebars | Partial (raymond.escape) | only HTML |
| Kasia | Partial (kasia.WriteEscapedHtml) | only HTML |
| Mustache | Partial (template.HTMLEscape) | only HTML |
| Pongo2 | Partial (pongo2.filterEscape, pongo2.filterEscapejs) | autoescape only escapes HTML, others could be implemented as pongo filters |
| Soy | Partial (template.HTMLEscapeString, url.QueryEscape, template.JSEscapeString) | autoescape only escapes HTML, contextual escaping is defined as a project goal |
