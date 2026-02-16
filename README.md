# goTemplateBenchmark

comparing the performance of different template engines

## full featured template engines

- [Ace](https://github.com/yosssi/ace)
- [Amber](https://github.com/eknkc/amber)
- [Go](https://golang.org/pkg/html/template)
- [Handlebars](https://github.com/aymerick/raymond)
- [Mustache](https://github.com/hoisie/mustache)
- [Pongo2](https://github.com/flosch/pongo2)
- [Soy](https://github.com/robfig/soy)
- [Jet](https://github.com/CloudyKit/jet)

## precompilation to Go code

- [ego](https://github.com/benbjohnson/ego)
- [ftmpl](https://github.com/tkrajina/ftmpl)
- [Goh](https://github.com/OblivionOcean/Goh)
- [Gorazor](https://github.com/sipin/gorazor)
- [Quicktemplate](https://github.com/valyala/quicktemplate)
- [Hero](https://github.com/shiyanhui/hero)
- [Jade](https://github.com/Joker/jade)
- [templ](https://github.com/a-h/templ)
- [gomponents](https://github.com/maragudk/gomponents)
- [hb](https://github.com/dracory/hb)

## baseline benchmarks for comparison

- DirectBuffer - Use go to write the HTML by hand to the buffer with basic escaping

## transpiling to Go Template

- [Damsel](https://github.com/dskinner/damsel) I won't benchmark transpiling
  engines, because transpilation should just happen once at startup. If you
  cache the transpilation result, which is recommended, you would have the same
  performance numbers as html/template for rendering.

## Why?

Just for fun. Go Templates work nice out of the box and should be used for
rendering from a security point of view. If you care about performance you
should cache the rendered output.

Sometimes there are templates that cannot be reasonably cached. Then you might
need a really fast template engine with code generation.

## How to run the benchmarks

```
./bench.sh -c go
```

## Results dev machine

local desktop: ryzen 3900x

## simple benchmarks
### full featured template engines
| Name       | Runs      | µs/op | B/op  | allocations/op |
| ---------- | --------- | ----- | ----- | -------------- |
| Ace        | 501,468   | 6.896 | 1,073 | 32             |
| Amber      | 829,180   | 4.371 | 753   | 28             |
| Golang     | 829,586   | 4.226 | 673   | 27             |
| GolangText | 2,524,951 | 1.428 | 224   | 7              |
| Handlebars | 491,289   | 7.281 | 3,407 | 73             |
| JetHTML    | 4,463,096 | 0.786 | 0     | 0              |
| Mustache   | 1,595,821 | 2.244 | 1,611 | 24             |
| Pongo2     | 1,000,000 | 3.159 | 2,059 | 30             |
| Soy        | 1,839,766 | 1.968 | 1,224 | 19             |


### precompilation to Go code
| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 6,319,710  | 0.565 | 85    | 8              |
| Ftmpl         | 3,803,455  | 0.929 | 774   | 12             |
| Goh           | 41,145,136 | 0.084 | 0     | 0              |
| Gomponents    | 2,234,062  | 1.607 | 776   | 25             |
| Gorazor       | 7,170,680  | 0.501 | 512   | 5              |
| HB            | 1,387,557  | 2.599 | 2,064 | 44             |
| Hero          | 31,017,628 | 0.125 | 0     | 0              |
| Jade          | 25,551,121 | 0.140 | 0     | 0              |
| Quicktemplate | 19,508,510 | 0.181 | 0     | 0              |
| Templ         | 4,734,141  | 0.760 | 182   | 10             |


## more complex test with template inheritance (if possible)
### full featured template engines
| Name              | Runs    | µs/op  | B/op  | allocations/op |
| ----------------- | ------- | ------ | ----- | -------------- |
| ComplexGolang     | 99,990  | 36.024 | 5,474 | 209            |
| ComplexGolangText | 214,260 | 16.702 | 2,396 | 78             |
| ComplexJetHTML    | 462,312 | 7.612  | 535   | 5              |
| ComplexMustache   | 248,392 | 14.313 | 6,730 | 112            |


### precompilation to Go code
| Name                  | Runs      | µs/op | B/op  | allocations/op |
| --------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo            | 1,275,114 | 2.815 | 569   | 31             |
| ComplexFtmpl          | 856,627   | 4.137 | 3,535 | 38             |
| ComplexGoDirectBuffer | 5,595,393 | 0.656 | 0     | 0              |
| ComplexGoh            | 6,029,959 | 0.593 | 0     | 0              |
| ComplexGorazor        | 1,000,000 | 3.283 | 3,688 | 24             |
| ComplexHero           | 3,601,705 | 1.003 | 0     | 0              |
| ComplexJade           | 3,236,332 | 1.060 | 0     | 0              |
| ComplexQuicktemplate  | 3,428,479 | 1.033 | 0     | 0              |
| ComplexTempl          | 1,000,000 | 3.168 | 762   | 38             |

## Security

All packages assume that template authors are trusted. If you allow custom
templates you have to sanitize your user input e.g.
[bluemonday](https://github.com/microcosm-cc/bluemonday). Generally speaking I
would suggest to sanitize every input not just HTML-input.

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
