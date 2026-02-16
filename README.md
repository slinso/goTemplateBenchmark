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
| Name       | Runs      | µs/op  | B/op  | allocations/op |
| ---------- | --------- | ------ | ----- | -------------- |
| Ace        | 256,797   | 13.017 | 1,121 | 40             |
| Amber      | 373,540   | 8.926  | 849   | 36             |
| Golang     | 604,209   | 8.650  | 769   | 35             |
| GolangText | 1,426,340 | 2.453  | 128   | 7              |
| Handlebars | 254,118   | 13.904 | 3,424 | 75             |
| JetHTML    | 4,334,586 | 0.794  | 0     | 0              |
| Mustache   | 810,091   | 4.483  | 1,723 | 30             |
| Pongo2     | 586,269   | 5.805  | 2,075 | 32             |
| Soy        | 1,000,000 | 3.545  | 1,224 | 19             |


### precompilation to Go code
| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 3,287,790  | 1.024 | 85    | 8              |
| Ftmpl         | 2,133,676  | 1.704 | 774   | 12             |
| Goh           | 41,564,152 | 0.085 | 0     | 0              |
| Gomponents    | 630,400    | 5.394 | 1,240 | 64             |
| Gorazor       | 4,209,751  | 0.852 | 512   | 5              |
| HB            | 735,676    | 4.749 | 1,984 | 36             |
| Hero          | 30,954,032 | 0.120 | 0     | 0              |
| Jade          | 38,540,751 | 0.083 | 0     | 0              |
| Quicktemplate | 17,429,122 | 0.183 | 0     | 0              |
| Templ         | 6,126,406  | 0.575 | 96    | 2              |


## more complex test with template inheritance (if possible)
### full featured template engines
| Name              | Runs    | µs/op  | B/op  | allocations/op |
| ----------------- | ------- | ------ | ----- | -------------- |
| ComplexGolang     | 49,784  | 72.503 | 6,565 | 290            |
| ComplexGolangText | 120,789 | 32.198 | 2,236 | 107            |
| ComplexJetHTML    | 287,542 | 12.108 | 535   | 5              |
| ComplexMustache   | 128,841 | 27.375 | 7,275 | 156            |


### precompilation to Go code
| Name                  | Runs      | µs/op | B/op  | allocations/op |
| --------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo            | 1,000,000 | 5.352 | 569   | 31             |
| ComplexFtmpl          | 512,026   | 7.545 | 3,536 | 38             |
| ComplexGoDirectBuffer | 6,203,253 | 0.543 | 0     | 0              |
| ComplexGoh            | 6,074,386 | 0.569 | 0     | 0              |
| ComplexGorazor        | 636,561   | 5.653 | 3,688 | 24             |
| ComplexHero           | 3,706,204 | 0.955 | 0     | 0              |
| ComplexJade           | 4,795,210 | 0.708 | 0     | 0              |
| ComplexQuicktemplate  | 3,466,336 | 1.012 | 0     | 0              |
| ComplexTempl          | 1,265,850 | 2.790 | 408   | 11             |

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
