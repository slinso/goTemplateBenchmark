# goTemplateBenchmark

comparing the performance of different template engines

## full featured template engines

- [Ace](https://github.com/yosssi/ace)
- [Amber](https://github.com/eknkc/amber)
- [Go](https://golang.org/pkg/html/template)
- [Handlebars](https://github.com/aymerick/raymond)
- removed - [Kasia](https://github.com/ziutek/kasia.go)
- [Mustache](https://github.com/hoisie/mustache)
- [Pongo2](https://github.com/flosch/pongo2)
- [Soy](https://github.com/robfig/soy)
- [Jet](https://github.com/CloudyKit/jet)

## precompilation to Go code

- [ego](https://github.com/benbjohnson/ego)
- removed - [egon](https://github.com/commondream/egon)
- [egonslinso](https://github.com/SlinSo/egon)
- [ftmpl](https://github.com/tkrajina/ftmpl)
- [Gorazor](https://github.com/sipin/gorazor)
- [Quicktemplate](https://github.com/valyala/quicktemplate)
- [Hero](https://github.com/shiyanhui/hero)
- [Jade](https://github.com/Joker/jade)

## special benchmarks for comparison

- Go text/template (do not use this for HTML)
- StaticString - Use one static string for the whole Template to have a base time
- DirectBuffer - Use go to write the HTML by hand to the buffer

## transpiling to Go Template

- [Damsel](https://github.com/dskinner/damsel)
  I won't benchmark transpiling engines, because transpilation should just happen once at startup. If you cache the transpilation result, which is recommended, you would have the same performance numbers as html/template for rendering.

## Why?

Just for fun. Go Templates work nice out of the box and should be used for rendering from a security point of view.
If you care about performance you should cache the rendered output.

Sometimes there are templates that cannot be reasonably cached. Then you possibly need a really fast template engine with code generation.

## Results dev machine

Changed the environment to my local dev laptop: i7-6700T 16GB Mem
Golang: 1.11

## Changes with 1.11

- Pongo and Soy got about 25% improved

## Changes with 1.9

There are quite some impressive performance improvements. Almost all pre compilation engines gained 10%-20%.

## special benchmarks

| Name                  | Runs        | ns/op | B/op | allocations/op |
| --------------------- | ----------- | ----- | ---- | -------------- |
| ComplexGoDirectBuffer | 5,925,524   | 590   | 0    | 0              |
| ComplexGoStaticString | 198,012,804 | 19    | 0    | 0              |

```
comparing: 1.11 to 1.13.5
benchmark                            old ns/op     new ns/op     delta
BenchmarkComplexGoDirectBuffer-8     621           590           -4.99%
BenchmarkComplexGoStaticString-8     24.0          18.5          -22.92%
```

## simple benchmarks

### full featured template engines

| Name        | Runs      | µs/op | B/op  | allocations/op |
| ----------- | --------- | ----- | ----- | -------------- |
| Ace         | 403,213   | 8.873 | 1,392 | 42             |
| Amber       | 536,733   | 5.674 | 1,120 | 38             |
| Golang      | 612,609   | 5.401 | 1,040 | 37             |
| GolangText  | 2,128,748 | 1.783 | 144   | 9              |
| Handlebars  | 397,616   | 9.666 | 4,018 | 82             |
| **JetHTML** | 3,283,270 | 1.075 | 0     | 0              |
| Mustache    | 1,000,000 | 3.282 | 1,568 | 29             |
| Pongo2      | 828,104   | 4.159 | 2,072 | 32             |
| Soy         | 1,321,779 | 2.766 | 1,392 | 25             |

```
comparing: 1.11 to 1.13.5
benchmark                 old ns/op     new ns/op     delta
BenchmarkAce-8            8970          8873          -1.08%
BenchmarkAmber-8          5386          5674          +5.35%
BenchmarkGolang-8         5264          5401          +2.60%
BenchmarkGolangText-8     1664          1783          +7.15%
BenchmarkHandlebars-8     9567          9666          +1.03%
BenchmarkJetHTML-8        1093          1075          -1.65%
BenchmarkMustache-8       3357          3282          -2.23%
BenchmarkPongo2-8         3996          4159          +4.08%
BenchmarkSoy-8            2795          2766          -1.04%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 4,448,131  | 0.800 | 85    | 8              |
| EgonSlinso    | 10,317,644 | 0.324 | 0     | 0              |
| Ftmpl         | 2,765,876  | 1.314 | 1,094 | 12             |
| Gorazor       | 6,207,387  | 0.556 | 512   | 5              |
| Hero          | 20,645,748 | 0.173 | 0     | 0              |
| **Jade**      | 34,540,056 | 0.102 | 0     | 0              |
| Quicktemplate | 12,412,521 | 0.285 | 0     | 0              |

```
comparing: 1.11 to 1.13.5
ignoring BenchmarkEgon-8: before has 1 instances, after has 0
benchmark                    old ns/op     new ns/op     delta
BenchmarkEgo-8               805           800           -0.62%
BenchmarkEgonSlinso-8        340           324           -4.71%
BenchmarkFtmpl-8             1371          1314          -4.16%
BenchmarkGorazor-8           1094          556           -49.18%
BenchmarkHero-8              168           173           +2.98%
BenchmarkJade-8              103           102           -0.97%
BenchmarkQuicktemplate-8     289           285           -1.38%

benchmark              old allocs     new allocs     delta
BenchmarkGorazor-8     11             5              -54.55%

benchmark              old bytes     new bytes     delta
BenchmarkFtmpl-8       1141          1094          -4.12%
BenchmarkGorazor-8     613           512           -16.48%
```

## more complex test with template inheritance (if possible)

### full featured template engines

| Name               | Runs    | µs/op  | B/op  | allocations/op |
| ------------------ | ------- | ------ | ----- | -------------- |
| ComplexGolang      | 76,704  | 47.023 | 8,862 | 296            |
| ComplexGolangText  | 163,388 | 20.565 | 2,793 | 113            |
| **ComplexJetHTML** | 364,692 | 9.689  | 546   | 5              |
| ComplexMustache    | 166,411 | 20.409 | 7,558 | 155            |

```
comparing: 1.11 to 1.13.5
benchmark                        old ns/op     new ns/op     delta
BenchmarkComplexGolang-8         45252         47023         +3.91%
BenchmarkComplexGolangText-8     19980         20565         +2.93%
BenchmarkComplexJetHTML-8        10155         9689          -4.59%
BenchmarkComplexMustache-8       20542         20409         -0.65%

benchmark                      old allocs     new allocs     delta
BenchmarkComplexGolang-8       293            296            +1.02%
BenchmarkComplexMustache-8     161            155            -3.73%

benchmark                      old bytes     new bytes     delta
BenchmarkComplexGolang-8       10478         8862          -15.42%
BenchmarkComplexMustache-8     7813          7558          -3.26%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 754,148   | 4.218 | 656   | 36             |
| ComplexEgoSlinso     | 1,946,601 | 1.984 | 160   | 2              |
| ComplexFtmpl         | 732,259   | 5.750 | 4,995 | 43             |
| ComplexGorazor       | 946,285   | 4.020 | 3,056 | 34             |
| ComplexHero          | 2,798,424 | 1.332 | 0     | 0              |
| **ComplexJade**      | 3,491,661 | 0.919 | 0     | 0              |
| ComplexQuicktemplate | 2,219,455 | 1.664 | 0     | 0              |

```
comparing: 1.11 to 1.13.5
ignoring BenchmarkComplexEgon-8: before has 1 instances, after has 0
benchmark                           old ns/op     new ns/op     delta
BenchmarkComplexEgo-8               3899          4218          +8.18%
BenchmarkComplexEgoSlinso-8         1813          1984          +9.43%
BenchmarkComplexFtmpl-8             5970          5750          -3.69%
BenchmarkComplexGorazor-8           9010          4020          -55.38%
BenchmarkComplexHero-8              1225          1332          +8.73%
BenchmarkComplexJade-8              938           919           -2.03%
BenchmarkComplexQuicktemplate-8     1629          1664          +2.15%

benchmark                     old allocs     new allocs     delta
BenchmarkComplexGorazor-8     64             34             -46.88%

benchmark                     old bytes     new bytes     delta
BenchmarkComplexFtmpl-8       5042          4995          -0.93%
BenchmarkComplexGorazor-8     8444          3056          -63.81%
```

## Security

All packages assume that template authors are trusted. If you allow custom templates you have to sanitize your user input e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday). Generally speaking I would suggest to sanitize every input not just HTML-input.

### Attention: This part is not updated since 2016.

| Framework     | Security                                                                      | Comment                                                                                                      |
| ------------- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| Ace           | No                                                                            |                                                                                                              |
| amber         | No                                                                            |                                                                                                              |
| ego           | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| egon          | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| egonslinso    | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| ftmpl         | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| Go            | Yes                                                                           | contextual escaping [html/template Security Model](https://golang.org/pkg/html/template/#hdr-Security_Model) |
| Gorazor       | Partial (template.HTMLEscapeString)                                           | only HTML, others need to be called manually                                                                 |
| Handlebars    | Partial (raymond.escape)                                                      | only HTML                                                                                                    |
| Hero          | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| Jade          | Partial (html.EscapeString)                                                   | Autoescape for HTML, others need to be called manually                                                       |
| Jet           | Partial (html.EscapeString)                                                   | Autoescape for HTML, others need to be called manually                                                       |
| Kasia         | Partial (kasia.WriteEscapedHtml)                                              | only HTML                                                                                                    |
| Mustache      | Partial (template.HTMLEscape)                                                 | only HTML                                                                                                    |
| Pongo2        | Partial (pongo2.filterEscape, pongo2.filterEscapejs)                          | autoescape only escapes HTML, others could be implemented as pongo filters                                   |
| Quicktemplate | Partial (html.EscapeString)                                                   | only HTML, others need to be called manually                                                                 |
| Soy           | Partial (template.HTMLEscapeString, url.QueryEscape, template.JSEscapeString) | autoescape only escapes HTML, contextual escaping is defined as a project goal                               |
