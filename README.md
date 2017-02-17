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
Golang: 1.8

### full featured template engines
| Name           |    Runs |   µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ace            |  10,000 | 108.212 | 5,208 |             77 |
| Amber          |  30,000 |  38.085 | 1,448 |             39 |
| Golang         |  50,000 |  45.915 | 1,368 |             38 |
| Handlebars     |  20,000 |  71.592 | 4,256 |             90 |
| **JetHTML**        | 200,000 |   6.628 |     0 |              0 |
| Kasia          | 100,000 |  19.993 | 1,192 |             26 |
| Mustache       | 100,000 |  20.597 | 1,568 |             28 |
| Pongo2         |  50,000 |  32.215 | 2,376 |             47 |
| Soy            | 100,000 |  18.052 | 1,384 |             26 |


### precompilation to Go code
| Name              |      Runs | µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ego               |   300,000 | 4.339 |    85 |              8 |
| Egon              |   200,000 | 9.939 |   309 |             22 |
| EgonSlinso        | 1,000,000 | 1.560 |     0 |              0 |
| Ftmpl             |   200,000 | 9.901 | 1,141 |             12 |
| Gorazor           |   200,000 | 6.419 |   613 |             11 |
| **Hero**              | 2,000,000 | 1.014 |     0 |              0 |
| Quicktemplate     | 1,000,000 | 1.421 |     0 |              0 |



### more complex test with template inheritance (if possible)
| Name                     |    Runs |   µs/op |   B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexEgo               |  50,000 |  25.561 |    656 |             41 |
| ComplexEgoSlinso         | 200,000 |  10.873 |    165 |              7 |
| ComplexEgon              |  20,000 |  54.666 |  1,616 |            101 |
| ComplexFtmpl             |  30,000 |  39.898 |  5,040 |             48 |
| ComplexFtmplInclude      |  50,000 |  38.558 |  5,040 |             48 |
| ComplexGolang            |   5,000 | 324.229 | 10,534 |            300 |
| ComplexGorazor           |  20,000 |  65.143 |  8,449 |             73 |
| **ComplexHero**              | 200,000 |   8.248 |    165 |              7 |
| ComplexJetHTML           |  20,000 |  57.318 |    545 |              5 |
| ComplexMustache          |  10,000 | 135.291 |  7,849 |            166 |
| ComplexQuicktemplate     | 200,000 |   9.650 |      0 |              0 |


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
