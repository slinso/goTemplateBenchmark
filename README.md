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
- StaticString - Use one static string for the whole Template to have a base
  time
- DirectBuffer - Use go to write the HTML by hand to the buffer

## transpiling to Go Template

- [Damsel](https://github.com/dskinner/damsel) I won't benchmark transpiling
  engines, because transpilation should just happen once at startup. If you
  cache the transpilation result, which is recommended, you would have the same
  performance numbers as html/template for rendering.

## Why?

Just for fun. Go Templates work nice out of the box and should be used for
rendering from a security point of view. If you care about performance you
should cache the rendered output.

Sometimes there are templates that cannot be reasonably cached. Then you
possibly need a really fast template engine with code generation.

## Results dev machine

local desktop: ryzen 3900x

## Changes with 1.15

- benchmarks run on a new system. I ran 1.14.13 and 1.15.5 on the same machine.

## Changes with 1.14

- quite a few slowdowns, but I did not do any profiling. needs more testing.

## Changes with 1.11

- Pongo and Soy got about 25% improved

## Changes with 1.9

There are quite some impressive performance improvements. Almost all pre
compilation engines gained 10%-20%.

## special benchmarks

| Name                  | Runs        | ns/op | B/op | allocations/op |
| --------------------- | ----------- | ----- | ---- | -------------- |
| ComplexGoDirectBuffer | 7,914,384   | 450   | 0    | 0              |
| ComplexGoStaticString | 287,497,974 | 12    | 0    | 0              |

```
comparing: go1.14.13 to go version go1.15.5 linux/amd64
name                      old time/op    new time/op    delta
ComplexGoDirectBuffer-24     465ns ± 0%     450ns ± 0%  -3.23%
ComplexGoStaticString-24    11.9ns ± 0%    12.4ns ± 0%  +4.20%

name                      old alloc/op   new alloc/op   delta
ComplexGoDirectBuffer-24     0.00B          0.00B        0.00%
ComplexGoStaticString-24     0.00B          0.00B        0.00%

name                      old allocs/op  new allocs/op  delta
ComplexGoDirectBuffer-24      0.00           0.00        0.00%
ComplexGoStaticString-24      0.00           0.00        0.00%
```

## simple benchmarks

### full featured template engines

| Name        | Runs      | µs/op  | B/op  | allocations/op |
| ----------- | --------- | ------ | ----- | -------------- |
| Ace         | 271,477   | 12.135 | 1,378 | 40             |
| Amber       | 427,455   | 8.165  | 1,105 | 36             |
| Golang      | 422,346   | 7.959  | 1,025 | 35             |
| GolangText  | 1,445,512 | 2.489  | 128   | 7              |
| Handlebars  | 257,184   | 13.704 | 3,983 | 78             |
| **JetHTML** | 3,632,552 | 0.971  | 0     | 0              |
| Mustache    | 878,670   | 4.198  | 1,570 | 29             |
| Pongo2      | 616,092   | 5.774  | 2,074 | 32             |
| Soy         | 941,029   | 3.968  | 1,352 | 20             |

```
comparing: go1.14.13 to go version go1.15.5 linux/amd64
name           old time/op    new time/op    delta
Golang-24        8.24µs ± 0%    7.96µs ± 0%   -3.40%
GolangText-24    2.69µs ± 0%    2.49µs ± 0%   -7.64%
Ace-24           12.5µs ± 0%    12.1µs ± 0%   -3.01%
Amber-24         8.40µs ± 0%    8.17µs ± 0%   -2.79%
Mustache-24      4.08µs ± 0%    4.20µs ± 0%   +2.82%
Pongo2-24        5.78µs ± 0%    5.77µs ± 0%   -0.03%
Handlebars-24    14.0µs ± 0%    13.7µs ± 0%   -2.27%
Soy-24           4.06µs ± 0%    3.97µs ± 0%   -2.36%
JetHTML-24        974ns ± 0%     971ns ± 0%   -0.31%

name           old alloc/op   new alloc/op   delta
Golang-24        1.04kB ± 0%    1.02kB ± 0%   -1.54%
GolangText-24      144B ± 0%      128B ± 0%  -11.11%
Ace-24           1.39kB ± 0%    1.38kB ± 0%   -1.15%
Amber-24         1.12kB ± 0%    1.10kB ± 0%   -1.43%
Mustache-24      1.57kB ± 0%    1.57kB ± 0%    0.00%
Pongo2-24        2.07kB ± 0%    2.07kB ± 0%    0.00%
Handlebars-24    4.02kB ± 0%    3.98kB ± 0%   -0.99%
Soy-24           1.39kB ± 0%    1.35kB ± 0%   -2.87%
JetHTML-24        0.00B          0.00B         0.00%

name           old allocs/op  new allocs/op  delta
Golang-24          37.0 ± 0%      35.0 ± 0%   -5.41%
GolangText-24      9.00 ± 0%      7.00 ± 0%  -22.22%
Ace-24             42.0 ± 0%      40.0 ± 0%   -4.76%
Amber-24           38.0 ± 0%      36.0 ± 0%   -5.26%
Mustache-24        29.0 ± 0%      29.0 ± 0%    0.00%
Pongo2-24          32.0 ± 0%      32.0 ± 0%    0.00%
Handlebars-24      82.0 ± 0%      78.0 ± 0%   -4.88%
Soy-24             25.0 ± 0%      20.0 ± 0%  -20.00%
JetHTML-24         0.00           0.00         0.00%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 2,723,541  | 1.307 | 85    | 8              |
| EgonSlinso    | 13,241,811 | 0.269 | 0     | 0              |
| Ftmpl         | 2,104,824  | 1.714 | 1,095 | 12             |
| Gorazor       | 4,454,704  | 0.802 | 512   | 5              |
| Hero          | 25,230,196 | 0.139 | 0     | 0              |
| **Jade**      | 43,579,422 | 0.090 | 0     | 0              |
| Quicktemplate | 13,199,523 | 0.229 | 0     | 0              |

```
comparing: go1.14.13 to go version go1.15.5 linux/amd64
name              old time/op    new time/op    delta
Ego-24              1.26µs ± 0%    1.31µs ± 0%  +3.73%
EgonSlinso-24        264ns ± 0%     269ns ± 0%  +1.89%
Quicktemplate-24     236ns ± 0%     229ns ± 0%  -2.97%
Ftmpl-24            1.73µs ± 0%    1.71µs ± 0%  -0.92%
Gorazor-24           747ns ± 0%     802ns ± 0%  +7.36%
Hero-24              144ns ± 0%     139ns ± 0%  -3.47%
Jade-24             90.7ns ± 0%    90.4ns ± 0%  -0.33%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%   0.00%
EgonSlinso-24        0.00B          0.00B        0.00%
Quicktemplate-24     0.00B          0.00B        0.00%
Ftmpl-24            1.09kB ± 0%    1.09kB ± 0%   0.00%
Gorazor-24            512B ± 0%      512B ± 0%   0.00%
Hero-24              0.00B          0.00B        0.00%
Jade-24              0.00B          0.00B        0.00%

name              old allocs/op  new allocs/op  delta
Ego-24                8.00 ± 0%      8.00 ± 0%   0.00%
EgonSlinso-24         0.00           0.00        0.00%
Quicktemplate-24      0.00           0.00        0.00%
Ftmpl-24              12.0 ± 0%      12.0 ± 0%   0.00%
Gorazor-24            5.00 ± 0%      5.00 ± 0%   0.00%
Hero-24               0.00           0.00        0.00%
Jade-24               0.00           0.00        0.00%
```

## more complex test with template inheritance (if possible)

### full featured template engines

| Name               | Runs    | µs/op  | B/op  | allocations/op |
| ------------------ | ------- | ------ | ----- | -------------- |
| ComplexGolang      | 51,696  | 69.767 | 8,783 | 285            |
| ComplexGolangText  | 115,188 | 30.380 | 2,708 | 102            |
| **ComplexJetHTML** | 292,257 | 15.095 | 550   | 5              |
| ComplexMustache    | 129,087 | 27.772 | 7,568 | 155            |

```
comparing: go1.14.13 to go version go1.15.5 linux/amd64
name                  old time/op    new time/op    delta
ComplexGolang-24        68.7µs ± 0%    69.8µs ± 0%  +1.56%
ComplexGolangText-24    31.5µs ± 0%    30.4µs ± 0%  -3.50%
ComplexMustache-24      26.4µs ± 0%    27.8µs ± 0%  +5.40%
ComplexJetHTML-24       15.2µs ± 0%    15.1µs ± 0%  -0.87%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        8.87kB ± 0%    8.78kB ± 0%  -0.99%
ComplexGolangText-24    2.80kB ± 0%    2.71kB ± 0%  -3.18%
ComplexMustache-24      7.57kB ± 0%    7.57kB ± 0%   0.00%
ComplexJetHTML-24         551B ± 0%      550B ± 0%  -0.18%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           296 ± 0%       285 ± 0%  -3.72%
ComplexGolangText-24       113 ± 0%       102 ± 0%  -9.73%
ComplexMustache-24         155 ± 0%       155 ± 0%   0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%   0.00%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 503,415   | 6.523 | 592   | 31             |
| ComplexEgoSlinso     | 1,348,783 | 2.670 | 160   | 2              |
| ComplexFtmpl         | 459,254   | 7.681 | 4,936 | 38             |
| ComplexGorazor       | 846,919   | 4.573 | 2,872 | 22             |
| ComplexHero          | 3,478,174 | 1.049 | 0     | 0              |
| **ComplexJade**      | 4,716,076 | 0.725 | 0     | 0              |
| ComplexQuicktemplate | 2,778,873 | 1.260 | 0     | 0              |

```
comparing: go1.14.13 to go version go1.15.5 linux/amd64
name                     old time/op    new time/op    delta
ComplexEgo-24              6.54µs ± 0%    6.52µs ± 0%   -0.24%
ComplexQuicktemplate-24    1.28µs ± 0%    1.26µs ± 0%   -1.95%
ComplexEgoSlinso-24        2.63µs ± 0%    2.67µs ± 0%   +1.52%
ComplexFtmpl-24            7.70µs ± 0%    7.68µs ± 0%   -0.27%
ComplexGorazor-24          4.37µs ± 0%    4.57µs ± 0%   +4.57%
ComplexHero-24             1.15µs ± 0%    1.05µs ± 0%   -8.86%
ComplexJade-24              737ns ± 0%     725ns ± 0%   -1.63%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                657B ± 0%      592B ± 0%   -9.89%
ComplexQuicktemplate-24     0.00B          0.00B         0.00%
ComplexEgoSlinso-24          160B ± 0%      160B ± 0%    0.00%
ComplexFtmpl-24            5.00kB ± 0%    4.94kB ± 0%   -1.30%
ComplexGorazor-24          2.87kB ± 0%    2.87kB ± 0%    0.00%
ComplexHero-24              0.00B          0.00B         0.00%
ComplexJade-24              0.00B          0.00B         0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                36.0 ± 0%      31.0 ± 0%  -13.89%
ComplexQuicktemplate-24      0.00           0.00         0.00%
ComplexEgoSlinso-24          2.00 ± 0%      2.00 ± 0%    0.00%
ComplexFtmpl-24              43.0 ± 0%      38.0 ± 0%  -11.63%
ComplexGorazor-24            22.0 ± 0%      22.0 ± 0%    0.00%
ComplexHero-24               0.00           0.00         0.00%
ComplexJade-24               0.00           0.00         0.00%
```

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
