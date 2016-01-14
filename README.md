# goTemplateBenchmark
comparing the performance of different template engines
* [Ace](https://github.com/yosssi/ace)
* [Amber](https://github.com/eknkc/amber)
* [Damsel](https://github.com/dskinner/damsel)
* [ego](https://github.com/benbjohnson/ego)
* [ftmpl](https://github.com/tkrajina/ftmpl)
* [Go](https://golang.org/pkg/html/template)
* [Gorazor](https://github.com/sipin/gorazor)
* [Handlebars](https://github.com/aymerick/raymond)
* [Kasia](https://github.com/ziutek/kasia.go)
* [Mustache](https://github.com/hoisie/mustache)
* [Pongo2](https://github.com/flosch/pongo2)
* [Soy](https://github.com/robfig/soy)

## Why?
Just for fun. Go Templates work nice out of the box.
If you really care about performance you will usually cache the rendered output.

## Results
Tests run on a VPS 1 CPU und 512 MB Ram

### normal Template Engines
`go test -bench "Ace|Amber|Damsel|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy" -benchmem -benchtime=5s`

```
BenchmarkGolang            50000             27615 ns/op            2079 B/op         38 allocs/op
BenchmarkAce               20000             69353 ns/op            5625 B/op         77 allocs/op
BenchmarkAmber             50000             27355 ns/op            2091 B/op         39 allocs/op
BenchmarkDamsel            50000             47903 ns/op            2440 B/op         50 allocs/op
BenchmarkMustache         100000             14665 ns/op            1648 B/op         28 allocs/op
BenchmarkPongo2           100000             18367 ns/op            2997 B/op         46 allocs/op
BenchmarkHandlebars        30000             50151 ns/op            4497 B/op         90 allocs/op
BenchmarkKasia            200000             13861 ns/op            1829 B/op         26 allocs/op
BenchmarkSoy              100000             14204 ns/op            1732 B/op         26 allocs/op
```

### Template Engines with manual precompilation
`go test -bench "Ego|Ftmpl|Gorazor" -benchmem -benchtime=5s`

```
BenchmarkEgo              500000              4076 ns/op             646 B/op          8 allocs/op
BenchmarkFtmpl            300000              6799 ns/op            1152 B/op         12 allocs/op
BenchmarkGorazor          200000              5302 ns/op             656 B/op         11 allocs/op
```
*ftmpl* and *gorazor* performs worse than *ego* in the benchmark because the Buffer is defined inside the template function which returns a string. This is imho just a benchmark problem, because in a real application you just generate the HTML once and will print the string.
Other than that *ftmpl* adds nice type safety, which could be implemented in _ego_ as well.
After I refactored the generated *ftmpl* code to accept the Buffer as a parameter, the performance and allocs are on par with *ego*.

## TODO
- Makefile
- Dependency Management