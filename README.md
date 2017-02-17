# goTemplateBenchmark
comparing the performance of different template engines

## full featured template engines
* [Ace](https://github.com/yosssi/ace)
* [Amber](https://github.com/eknkc/amber)
* [Go](https://golang.org/pkg/html/template)
* [Handlebars](https://github.com/aymerick/raymond)
* [Kasia](https://github.com/ziutek/kasia.go)
* [Mustache](https://github.com/hoisie/mustache)
* [Pongo2](https://github.com/flosch/pongo2)
* [Soy](https://github.com/robfig/soy)
* [Jet](https://github.com/CloudyKit/jet)

## precompilation to Go code
* [ego](https://github.com/benbjohnson/ego)
* [egon](https://github.com/commondream/egon)
* [egonslinso](https://github.com/SlinSo/egon)
* [ftmpl](https://github.com/tkrajina/ftmpl)
* [Gorazor](https://github.com/sipin/gorazor)
* [Quicktemplate](https://github.com/valyala/quicktemplate)
* [Hero](https://github.com/shiyanhui/hero)

## transpiling to Go Template
* [Damsel](https://github.com/dskinner/damsel)
I won't benchmark transpiling engines, because transpilation should just happen once at startup. If you cache the transpilation result, which is recommended, you would have the same performance numbers as html/template for rendering.


## Why?
Just for fun. Go Templates work nice out of the box and should be used for rendering from a security point of view.
If you care about performance you should cache the rendered output.

Sometimes there are templates that cannot be reasonably cached. Then you possibly need a really fast template engine with code generation.


## Results dev machine
Changed the environment to my local dev laptop: i7-6700T  16GB Mem
Golang: 1.8

### full featured template engines
| Name           |      Runs |  µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ace            |   300,000 | 15.124 | 5,210 |             77 |
| Amber          | 1,000,000 |  5.257 | 1,448 |             39 |
| Golang         | 1,000,000 |  5.171 | 1,368 |             38 |
| Handlebars     |   500,000 | 10.589 | 4,258 |             90 |
| **JetHTML**        | 5,000,000 |  1.167 |     0 |              0 |
| Kasia          | 1,000,000 |  3.216 | 1,192 |             26 |
| Mustache       | 1,000,000 |  3.365 | 1,568 |             28 |
| Pongo2         | 1,000,000 |  4.895 | 2,376 |             47 |
| Soy            | 1,000,000 |  3.178 | 1,384 |             26 |


### precompilation to Go code
| Name              |       Runs | µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ego               |  5,000,000 | 0.890 |    85 |              8 |
| Egon              |  2,000,000 | 1.997 |   309 |             22 |
| EgonSlinso        | 20,000,000 | 0.362 |     0 |              0 |
| Ftmpl             |  3,000,000 | 1.394 | 1,141 |             12 |
| Gorazor           |  3,000,000 | 1.212 |   613 |             11 |
| **Hero**              | 20,000,000 | 0.201 |     0 |              0 |
| Quicktemplate     | 20,000,000 | 0.324 |     0 |              0 |


### more complex test with template inheritance (if possible)
| Name                     |      Runs |  µs/op |   B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexEgo               | 1,000,000 |  4.300 |    656 |             41 |
| ComplexEgoSlinso         | 2,000,000 |  1.955 |    165 |              7 |
| ComplexEgon              |   500,000 |  9.098 |  1,617 |            101 |
| ComplexFtmpl             | 1,000,000 |  6.553 |  5,043 |             48 |
| ComplexFtmplInclude      | 1,000,000 |  6.490 |  5,043 |             48 |
| ComplexGolang            |   100,000 | 42.322 | 10,535 |            300 |
| ComplexGorazor           |   500,000 |  9.013 |  8,453 |             73 |
| **ComplexHero**              | 3,000,000 |  1.290 |    165 |              7 |
| ComplexJetHTML           |   500,000 |  9.914 |    546 |              5 |
| ComplexMustache          |   200,000 | 21.313 |  7,854 |            166 |
| ComplexQuicktemplate     | 2,000,000 |  1.920 |      0 |              0 |


## Results small VPS 
single CPU, 1GB RAM
Golang: 1.7.4

### full featured template engines
| Name           |    Runs |  µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ace            |  50,000 | 98.752 | 5,509 |             77 |
| Amber          | 100,000 | 39.983 | 2,050 |             39 |
| Golang         | 100,000 | 42.146 | 2,038 |             38 |
| Handlebars     | 100,000 | 67.894 | 4,256 |             90 |
| **JetHTML**        | 500,000 |  7.659 |   518 |              0 |
| Kasia          | 200,000 | 21.069 | 1,789 |             26 |
| Mustache       | 200,000 | 19.842 | 1,568 |             28 |
| Pongo2         | 200,000 | 25.075 | 2,949 |             46 |
| Soy            | 300,000 | 16.406 | 1,784 |             26 |

### precompilation to Go code
| Name              |      Runs |  µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ego               | 1,000,000 |  5.906 |   603 |              8 |
| Egon              |   300,000 | 10.358 | 1,172 |             22 |
| EgonSlinso        | 2,000,000 |  3.046 |   517 |              0 |
| Ftmpl             |   500,000 |  7.026 | 1,141 |             12 |
| Gorazor           | 1,000,000 |  7.210 |   613 |             11 |
| **Hero**              | 5,000,000 |  1.016 |     0 |              0 |
| Quicktemplate     | 1,000,000 |  3.579 |   999 |              0 |


### more complex test with template inheritance (if possible)
| Name                     |      Runs |   µs/op |   B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexEgo               |   200,000 |  27.553 |  3,037 |             41 |
| ComplexEgoSlinso         |   300,000 |  15.258 |  3,341 |              7 |
| ComplexEgon              |   100,000 |  51.385 |  3,998 |            101 |
| ComplexFtmpl             |   100,000 |  35.341 |  5,296 |             48 |
| ComplexFtmplInclude      |   200,000 |  36.389 |  5,296 |             48 |
| ComplexGolang            |    10,000 | 347.405 | 13,001 |            295 |
| ComplexGorazor           |    50,000 |  71.411 |  8,321 |             73 |
| **ComplexHero**              | 1,000,000 |   7.142 |    165 |              7 |
| ComplexJetHTML           |    50,000 |  73.631 |  2,807 |              5 |
| ComplexMustache          |    30,000 | 141.413 |  7,849 |            166 |
| ComplexQuicktemplate     |   300,000 |  15.086 |  3,153 |              0 |

## Security
All packages assume that template authors are trusted. If you allow custom templates you have to sanitize your user input e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday). Generally speaking I would suggest to sanitize every input not just HTML-input.

| Framework | Security | Comment |
| --------- | -------- | ------- |
| Ace | No | |
| amber | No | |
| ego | Partial (html.EscapeString) | only HTML, others need to be called manually |
| egon | Partial (html.EscapeString) | only HTML, others need to be called manually |
| egonslinso | Partial (html.EscapeString) | only HTML, others need to be called manually |
| ftmpl | Partial (html.EscapeString) | only HTML, others need to be called manually |
| Go | Yes | contextual escaping [html/template Security Model](https://golang.org/pkg/html/template/#hdr-Security_Model) |
| Gorazor | Partial (template.HTMLEscapeString) | only HTML, others need to be called manually |
| Handlebars | Partial (raymond.escape) | only HTML |
| Hero | Partial (html.EscapeString) | only HTML, others need to be called manually |
| Jet | Partial (html.EscapeString) | Autoescape for HTML, others need to be called manually |
| Kasia | Partial (kasia.WriteEscapedHtml) | only HTML |
| Mustache | Partial (template.HTMLEscape) | only HTML |
| Pongo2 | Partial (pongo2.filterEscape, pongo2.filterEscapejs) | autoescape only escapes HTML, others could be implemented as pongo filters |
| Quicktemplate | Partial (html.EscapeString) | only HTML, others need to be called manually |
| Soy | Partial (template.HTMLEscapeString, url.QueryEscape, template.JSEscapeString) | autoescape only escapes HTML, contextual escaping is defined as a project goal |
