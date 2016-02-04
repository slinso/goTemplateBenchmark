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
First thought:
Just for fun. Go Templates work nice out of the box.
If you really care about performance you will usually cache the rendered output.

on second thought:
I have some templates that I cannot cache in my production system, thats why I'm interested in performant
HTML generation from templates. After trying the code generation based projects I found ego best, but some
features where missing and generated code could be optimized further. That's why I created the fork
and included the results in this benchmark. Currently I'm working on the integration in a web framework like gin.

## Results
Tests run on a VPS 1 CPU und 512 MB Ram

### normal Template Engines
`go test -bench "Ace|Amber|Damsel|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy" -benchmem -benchtime=5s`

```
BenchmarkGolang           200000             37886 ns/op            2078 B/op         38 allocs/op
BenchmarkAce              100000             85394 ns/op            5548 B/op         77 allocs/op
BenchmarkAmber            200000             39391 ns/op            2090 B/op         39 allocs/op
BenchmarkDamsel           100000             62043 ns/op            2440 B/op         50 allocs/op
BenchmarkMustache         500000             19965 ns/op            1648 B/op         28 allocs/op
BenchmarkPongo2           200000             31120 ns/op            2997 B/op         46 allocs/op
BenchmarkHandlebars       100000             79132 ns/op            4496 B/op         90 allocs/op
BenchmarkKasia            500000             16305 ns/op            2187 B/op         26 allocs/op
BenchmarkSoy              300000             19098 ns/op            1832 B/op         26 allocs/op
```

### Template Engines with manual precompilation
`go test -bench "Ego|Ftmpl|Gorazor" -benchmem -benchtime=5s`

```
BenchmarkEgo             1000000              6085 ns/op             645 B/op          8 allocs/op
BenchmarkEgon             100000              7536 ns/op             999 B/op         22 allocs/op
BenchmarkEgonSlinso       500000              1279 ns/op             518 B/op          0 allocs/op
BenchmarkFtmpl           1000000              9398 ns/op            1152 B/op         12 allocs/op
BenchmarkGorazor         1000000              7885 ns/op             656 B/op         11 allocs/op
```
*ftmpl* and *gorazor* performs worse than *ego* in the benchmark because the Buffer is defined inside the template function which returns a string. This is imho just a benchmark problem, because in a real application you just generate the HTML once and will print the string.
Other than that *ftmpl* adds nice type safety, which could be implemented in _ego_ as well.
After I refactored the generated *ftmpl* code to accept the Buffer as a parameter, the performance and allocs are on par with *ego*.

### more complex test with template inheritance (where possible)
`go test . -bench="Complex" -benchmem -benchtime=500ms`

```
BenchmarkComplexGolang            5000            202169 ns/op           13515 B/op        295 allocs/op
BenchmarkComplexEgo              50000             16303 ns/op            3246 B/op         41 allocs/op
BenchmarkComplexEgon             20000             29604 ns/op            4805 B/op        101 allocs/op
BenchmarkComplexEgoSlinso       100000              8461 ns/op            2622 B/op          7 allocs/op
BenchmarkComplexFtmpl            30000             19625 ns/op            5204 B/op         40 allocs/op
BenchmarkComplexFtmplFctCall     30000             26941 ns/op            5749 B/op         48 allocs/op
BenchmarkComplexMustache         10000            141554 ns/op            8457 B/op        166 allocs/op
BenchmarkComplexGorazor          20000             44475 ns/op            8583 B/op         73 allocs/op
```

## TODO
- Makefile
