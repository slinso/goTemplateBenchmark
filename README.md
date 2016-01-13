# goTemplateBenchmark
comparing the performance of different template engines
* Golang html/template
* [ego](https://github.com/benbjohnson/ego)
* [ftmpl](https://github.com/tkrajina/ftmpl)
 
## Results
Tests run on a VPS 1 CPU und 512 MB Ram

`go test . -bench . -benchmem -benchtime=1s`

```
BenchmarkGolang-8         100000             19177 ns/op            2058 B/op         38 allocs/op
BenchmarkEgo-8            300000              3758 ns/op             991 B/op          8 allocs/op
BenchmarkFtmpl-8          300000              4468 ns/op            1152 B/op         12 allocs/op
```
*ftmpl* performs worse than _ego_ in the benchmark because the Buffer is defined inside the template function which returns a string. This is imho just a benchmark problem, because in a real application you just generate the HTML once and will print the string.
Other than that *ftmpl* adds nice type safety, which could be implemented in _ego_ as well.
After I refactored the generated *ftmpl* code to accept the Buffer as a parameter, the performance was on par with *ego*.
