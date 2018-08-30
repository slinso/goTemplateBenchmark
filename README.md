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

## special benchmarks for comparison
* Go text/template (do not use this for HTML) 
* StaticString - Use one static string for the whole Template to have a base time
* DirectBuffer - Use go to write the HTML by hand to the buffer

## transpiling to Go Template
* [Damsel](https://github.com/dskinner/damsel)
I won't benchmark transpiling engines, because transpilation should just happen once at startup. If you cache the transpilation result, which is recommended, you would have the same performance numbers as html/template for rendering.


## Why?
Just for fun. Go Templates work nice out of the box and should be used for rendering from a security point of view.
If you care about performance you should cache the rendered output.

Sometimes there are templates that cannot be reasonably cached. Then you possibly need a really fast template engine with code generation.


## Results dev machine
Changed the environment to my local dev laptop: i7-6700T  16GB Mem
Golang: 1.11


## Changes with 1.11
* Pongo and Soy got about 25% improved

## Changes with 1.9
There are quite some impressive performance improvements. Almost all pre compilation engines gained 10%-20%.

### special cases
| Name           |      Runs |  µs/op |  B/op | allocations/op | 
| --- | --- | --- | --- | --- |                   
| SimpleGoDirectBuffer      | 50,000,000 | 0.097 |     0 |              0 |
| ComplexGoDirectBuffer     |  10,000,000 | 0.621 |     0 |              0 |
| ComplexGoStaticString     | 200,000,000 | 0.024 |     0 |              0 |


### full featured template engines
| Name           |      Runs |  µs/op |  B/op | allocations/op | 
| --- | --- | --- | --- | --- |                                                    
| Ace            |   500,000 |  8.597 | 1,712 |             42 |
| Amber          | 1,000,000 |  5.628 | 1,440 |             38 |                  
| Golang         | 1,000,000 | 5.247 | 1,360 |             37 |
| GolangText     | 3,000,000 | 1.930 |   144 |              9 |
| Handlebars     |   500,000 | 9.793 | 4,226 |             83 |
| **JetHTML**        | 3,000,000 | 1.238 |     0 |              0 |
| Kasia          | 1,000,000 | 3.091 | 1,184 |             25 |
| Mustache       | 1,000,000 |  3.117 | 1,568 |             28 |
| Pongo2         | 1,000,000 |  3.864 | 2,136 |             32 |
| Soy            | 2,000,000 |  2.572 | 1,392 |             25 |

### precompilation to Go code
| Name              |       Runs | µs/op |  B/op | allocations/op |                
| --- | --- | --- | --- | --- |                                                    
| Ego                 |  5,000,000 | 0.822 |    85 |              8 |
| Egon                |  3,000,000 | 1.478 |   149 |             12 |
| EgonSlinso          | 20,000,000 | 0.334 |     0 |              0 |
| Ftmpl               |  3,000,000 | 1.318 | 1,141 |             12 |
| Gorazor             |  5,000,000 | 1.096 |   613 |             11 |
| **Hero**                | 30,000,000 | 0.177 |     0 |              0 |
| Quicktemplate       | 20,000,000 | 0.290 |     0 |              0 |


### more complex test with template inheritance (if possible)
### full featured template engines
| Name                  |    Runs |  µs/op |   B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexGolang         | 100,000 | 45.252 | 10,478 |            293 |
| ComplexGolangText     | 200,000 | 19.980 |  2,793 |            113 |
| **ComplexJetHTML**        | 500,000 | 10.155 |    546 |              5 |
| ComplexMustache       | 200,000 | 20.542 |  7,813 |            161 |

### more complex test with template inheritance (if possible)
### precompilation to Go code
 
| Name                      |        Runs | µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexEgo                |   1,000,000 | 3.967 |   656 |             36 |
| ComplexEgoSlinso          |   3,000,000 | 1.746 |   160 |              2 |
| ComplexEgon               |   1,000,000 | 6.803 |   960 |             55 |
| ComplexFtmpl              |   1,000,000 | 5.593 | 5,042 |             43 |
| ComplexGorazor            |     500,000 | 8.741 | 8,444 |             64 |
| **ComplexHero**               |   3,000,000 | 1.192 |     0 |              0 |
| ComplexQuicktemplate      |   3,000,000 | 1.590 |     0 |              0 |


## Results small VPS 
single CPU, 1GB RAM
Golang: 1.8

### full featured template engines
| Name           |    Runs |   µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ace            |  20,000 | 107.627 | 5,208 |             77 |
| Amber          |  30,000 |  37.748 | 1,448 |             39 |
| Golang         |  30,000 |  41.577 | 1,368 |             38 |
| Handlebars     |  20,000 |  64.995 | 4,256 |             90 |
| JetHTML        | 200,000 |   7.530 |     0 |              0 |
| Kasia          | 100,000 |  19.165 | 1,192 |             26 |
| Mustache       | 100,000 |  19.828 | 1,568 |             28 |
| Pongo2         | 100,000 |  44.673 | 2,376 |             47 |
| Soy            | 100,000 |  18.591 | 1,384 |             26 |


### precompilation to Go code
| Name              |      Runs |  µs/op |  B/op | allocations/op |
| --- | --- | --- | --- | --- |
| Ego               | 1,000,000 |  4.810 |    85 |              8 |
| Egon              |   200,000 | 10.859 |   309 |             22 |
| EgonSlinso        | 1,000,000 |  1.894 |     0 |              0 |
| Ftmpl             |   200,000 |  9.350 | 1,141 |             12 |
| Gorazor           |   200,000 |  7.134 |   613 |             11 |
| **Hero**              | 1,000,000 |  1.252 |     0 |              0 |
| Quicktemplate     | 1,000,000 |  1.406 |     0 |              0 |


### more complex test with template inheritance (if possible)
| Name                     |    Runs |   µs/op |   B/op | allocations/op |
| --- | --- | --- | --- | --- |
| ComplexEgo               | 100,000 |  25.522 |    656 |             41 |
| ComplexEgoSlinso         | 200,000 |  10.051 |    165 |              7 |
| ComplexEgon              |  30,000 |  48.603 |  1,616 |            101 |
| ComplexFtmpl             |  50,000 |  38.582 |  5,040 |             48 |
| ComplexFtmplInclude      |  30,000 |  46.322 |  5,040 |             48 |
| ComplexGolang            |  10,000 | 298.372 | 10,531 |            300 |
| ComplexGorazor           |  20,000 |  70.508 |  8,449 |             73 |
| **ComplexHero**              | 200,000 |   8.085 |      0 |              0 |
| ComplexJetHTML           |  30,000 |  66.645 |    544 |              5 |
| ComplexMustache          |  10,000 | 141.660 |  7,849 |            166 |
| ComplexQuicktemplate     | 300,000 |  10.346 |      0 |              0 |


## Security
All packages assume that template authors are trusted. If you allow custom templates you have to sanitize your user input e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday). Generally speaking I would suggest to sanitize every input not just HTML-input.

### Attention: This part is not updated since 2016.

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
