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

local dev laptop: i7-6700T 16GB Mem

## Changes with 1.14

- quite a few slowdowns, but I did not do any profiling. needs more testing.

## Changes with 1.11

- Pongo and Soy got about 25% improved

## Changes with 1.9

There are quite some impressive performance improvements. Almost all pre compilation engines gained 10%-20%.

## special benchmarks

| Name                  | Runs        | ns/op | B/op | allocations/op |
| --------------------- | ----------- | ----- | ---- | -------------- |
| ComplexGoDirectBuffer | 5,332,621   | 693   | 0    | 0              |
| ComplexGoStaticString | 175,642,789 | 21    | 0    | 0              |

```
comparing: 1.13.5 to 1.14
benchcmp is deprecated in favor of benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
benchmark                            old ns/op     new ns/op     delta
BenchmarkComplexGoDirectBuffer-8     590           693           +17.46%
BenchmarkComplexGoStaticString-8     18.5          21.0          +13.51%
```

## simple benchmarks

### full featured template engines

| Name        | Runs      | µs/op | B/op  | allocations/op |
| ----------- | --------- | ----- | ----- | -------------- |
| Ace         | 346,275   | 8.678 | 1,392 | 42             |
| Amber       | 573,464   | 5.387 | 1,120 | 38             |
| Golang      | 597,543   | 5.386 | 1,040 | 37             |
| GolangText  | 2,045,877 | 1.829 | 144   | 9              |
| Handlebars  | 382,017   | 9.852 | 4,019 | 82             |
| **JetHTML** | 3,230,307 | 1.116 | 0     | 0              |
| Mustache    | 1,000,000 | 3.067 | 1,568 | 29             |
| Pongo2      | 761,532   | 4.678 | 2,072 | 32             |
| Soy         | 962,260   | 3.157 | 1,392 | 25             |

```
comparing: 1.13.5 to 1.14
benchmark                 old ns/op     new ns/op     delta
BenchmarkAce-8            8873          8678          -2.20%
BenchmarkAmber-8          5674          5387          -5.06%
BenchmarkGolang-8         5401          5386          -0.28%
BenchmarkGolangText-8     1783          1829          +2.58%
BenchmarkHandlebars-8     9666          9852          +1.92%
BenchmarkJetHTML-8        1075          1116          +3.81%
BenchmarkMustache-8       3282          3067          -6.55%
BenchmarkPongo2-8         4159          4678          +12.48%
BenchmarkSoy-8            2766          3157          +14.14%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 3,587,478  | 0.921 | 85    | 8              |
| EgonSlinso    | 8,769,262  | 0.371 | 0     | 0              |
| Ftmpl         | 2,722,228  | 1.270 | 1,094 | 12             |
| Gorazor       | 5,341,615  | 0.606 | 512   | 5              |
| Hero          | 17,180,913 | 0.201 | 0     | 0              |
| **Jade**      | 28,891,586 | 0.139 | 0     | 0              |
| Quicktemplate | 10,127,134 | 0.338 | 0     | 0              |

```
comparing: 1.13.5 to 1.14
benchmark                    old ns/op     new ns/op     delta
BenchmarkEgo-8               800           921           +15.13%
BenchmarkEgonSlinso-8        324           371           +14.51%
BenchmarkFtmpl-8             1314          1270          -3.35%
BenchmarkGorazor-8           556           606           +8.99%
BenchmarkHero-8              173           201           +16.18%
BenchmarkJade-8              102           139           +36.27%
BenchmarkQuicktemplate-8     285           338           +18.60%
```

## more complex test with template inheritance (if possible)

### full featured template engines

| Name               | Runs    | µs/op  | B/op  | allocations/op |
| ------------------ | ------- | ------ | ----- | -------------- |
| ComplexGolang      | 78,091  | 47.037 | 8,862 | 296            |
| ComplexGolangText  | 160,622 | 27.318 | 2,793 | 113            |
| **ComplexJetHTML** | 266,722 | 11.525 | 546   | 5              |
| ComplexMustache    | 181,446 | 20.331 | 7,559 | 155            |

```
comparing: 1.13.5 to 1.14
benchmark                        old ns/op     new ns/op     delta
BenchmarkComplexGolang-8         47023         47037         +0.03%
BenchmarkComplexGolangText-8     20565         27318         +32.84%
BenchmarkComplexJetHTML-8        9689          11525         +18.95%
BenchmarkComplexMustache-8       20409         20331         -0.38%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 791,263   | 4.786 | 656   | 36             |
| ComplexEgoSlinso     | 1,727,679 | 2.084 | 160   | 2              |
| ComplexFtmpl         | 526,308   | 5.727 | 4,995 | 43             |
| ComplexGorazor       | 923,703   | 4.108 | 3,056 | 34             |
| ComplexHero          | 2,355,554 | 1.498 | 0     | 0              |
| **ComplexJade**      | 3,000,294 | 1.195 | 0     | 0              |
| ComplexQuicktemplate | 1,881,404 | 1.819 | 0     | 0              |

```
comparing: 1.13.5 to 1.14
benchmark                           old ns/op     new ns/op     delta
BenchmarkComplexEgo-8               4218          4786          +13.47%
BenchmarkComplexEgoSlinso-8         1984          2084          +5.04%
BenchmarkComplexFtmpl-8             5750          5727          -0.40%
BenchmarkComplexGorazor-8           4020          4108          +2.19%
BenchmarkComplexHero-8              1332          1498          +12.46%
BenchmarkComplexJade-8              919           1195          +30.03%
BenchmarkComplexQuicktemplate-8     1664          1819          +9.31%
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
