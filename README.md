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
- [Gorazor](https://github.com/sipin/gorazor)
- [Quicktemplate](https://github.com/valyala/quicktemplate)
- [Hero](https://github.com/shiyanhui/hero)
- [Jade](https://github.com/Joker/jade)
- [templ](https://github.com/a-h/templ)
- [gomponents](https://github.com/maragudk/gomponents)
- [hb](https://github.com/gouniverse/hb)

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

minimal configuration, but does not make sense, as it would compare the same 2
go versions.

```
./bench.sh -c go
```

testing your setup - reduce test runtime and assuming you have those two go
binaries on your system.

```
./bench.sh -t 10ms -c go1.18 -g go1.19
```

example for a normal benchmark run on my system with the default 3s runtime per
template engine

```
./bench.sh -c go1.18.6 -g go1.19.1
```

## Results dev machine

local desktop: ryzen 3900x

## simple benchmarks
### full featured template engines
| Name       | Runs      | µs/op  | B/op  | allocations/op |
| ---------- | --------- | ------ | ----- | -------------- |
| Ace        | 267,153   | 12.628 | 1,121 | 40             |
| Amber      | 393,578   | 8.569  | 849   | 36             |
| Golang     | 628,294   | 7.883  | 769   | 35             |
| GolangText | 1,532,516 | 2.366  | 128   | 7              |
| Handlebars | 266,960   | 13.046 | 3,424 | 75             |
| JetHTML    | 4,546,153 | 0.773  | 0     | 0              |
| Mustache   | 776,263   | 4.380  | 1,723 | 30             |
| Pongo2     | 566,592   | 5.787  | 2,075 | 32             |
| Soy        | 955,003   | 3.509  | 1,224 | 19             |

```
comparing: go1.22.7 to go1.23.1
name           old time/op    new time/op    delta
Golang-24        8.56µs ± 0%    7.88µs ± 0%  -7.91%
GolangText-24    2.44µs ± 0%    2.37µs ± 0%  -3.11%
Ace-24           13.1µs ± 0%    12.6µs ± 0%  -3.65%
Amber-24         8.90µs ± 0%    8.57µs ± 0%  -3.70%
Mustache-24      4.27µs ± 0%    4.38µs ± 0%  +2.50%
Pongo2-24        5.72µs ± 0%    5.79µs ± 0%  +1.24%
Handlebars-24    13.1µs ± 0%    13.0µs ± 0%  -0.21%
Soy-24           3.47µs ± 0%    3.51µs ± 0%  +1.18%
JetHTML-24        769ns ± 0%     773ns ± 0%  +0.52%

name           old alloc/op   new alloc/op   delta
Golang-24          769B ± 0%      769B ± 0%   0.00%
GolangText-24      128B ± 0%      128B ± 0%   0.00%
Ace-24           1.12kB ± 0%    1.12kB ± 0%   0.00%
Amber-24           849B ± 0%      849B ± 0%   0.00%
Mustache-24      1.72kB ± 0%    1.72kB ± 0%   0.00%
Pongo2-24        2.08kB ± 0%    2.08kB ± 0%   0.00%
Handlebars-24    3.42kB ± 0%    3.42kB ± 0%   0.00%
Soy-24           1.22kB ± 0%    1.22kB ± 0%   0.00%
JetHTML-24        0.00B          0.00B        0.00%

name           old allocs/op  new allocs/op  delta
Golang-24          35.0 ± 0%      35.0 ± 0%   0.00%
GolangText-24      7.00 ± 0%      7.00 ± 0%   0.00%
Ace-24             40.0 ± 0%      40.0 ± 0%   0.00%
Amber-24           36.0 ± 0%      36.0 ± 0%   0.00%
Mustache-24        30.0 ± 0%      30.0 ± 0%   0.00%
Pongo2-24          32.0 ± 0%      32.0 ± 0%   0.00%
Handlebars-24      75.0 ± 0%      75.0 ± 0%   0.00%
Soy-24             19.0 ± 0%      19.0 ± 0%   0.00%
JetHTML-24         0.00           0.00        0.00%
```

### precompilation to Go code
| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 3,811,827  | 1.012 | 85    | 8              |
| Ftmpl         | 2,185,862  | 1.607 | 774   | 12             |
| Gomponents    | 654,105    | 4.875 | 1,112 | 56             |
| Gorazor       | 4,444,626  | 0.820 | 512   | 5              |
| HB            | 498,691    | 7.534 | 2,448 | 51             |
| Hero          | 30,477,858 | 0.120 | 0     | 0              |
| Jade          | 41,132,322 | 0.085 | 0     | 0              |
| Quicktemplate | 18,337,161 | 0.180 | 0     | 0              |
| Templ         | 6,214,070  | 0.557 | 96    | 2              |

```
comparing: go1.22.7 to go1.23.1
name              old time/op    new time/op    delta
Ego-24              1.00µs ± 0%    1.01µs ± 0%  +1.60%
HB-24               7.48µs ± 0%    7.53µs ± 0%  +0.72%
Quicktemplate-24     176ns ± 0%     180ns ± 0%  +2.45%
Ftmpl-24            1.66µs ± 0%    1.61µs ± 0%  -3.31%
Gorazor-24           829ns ± 0%     820ns ± 0%  -1.13%
Hero-24              118ns ± 0%     120ns ± 0%  +1.52%
Jade-24             87.0ns ± 0%    85.2ns ± 0%  -1.98%
Templ-24             543ns ± 0%     557ns ± 0%  +2.73%
Gomponents-24       4.87µs ± 0%    4.88µs ± 0%  +0.02%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%   0.00%
HB-24               2.45kB ± 0%    2.45kB ± 0%   0.00%
Quicktemplate-24     0.00B          0.00B        0.00%
Ftmpl-24              774B ± 0%      774B ± 0%   0.00%
Gorazor-24            512B ± 0%      512B ± 0%   0.00%
Hero-24              0.00B          0.00B        0.00%
Jade-24              0.00B          0.00B        0.00%
Templ-24             96.0B ± 0%     96.0B ± 0%   0.00%
Gomponents-24       1.11kB ± 0%    1.11kB ± 0%   0.00%

name              old allocs/op  new allocs/op  delta
Ego-24                8.00 ± 0%      8.00 ± 0%   0.00%
HB-24                 51.0 ± 0%      51.0 ± 0%   0.00%
Quicktemplate-24      0.00           0.00        0.00%
Ftmpl-24              12.0 ± 0%      12.0 ± 0%   0.00%
Gorazor-24            5.00 ± 0%      5.00 ± 0%   0.00%
Hero-24               0.00           0.00        0.00%
Jade-24               0.00           0.00        0.00%
Templ-24              2.00 ± 0%      2.00 ± 0%   0.00%
Gomponents-24         56.0 ± 0%      56.0 ± 0%   0.00%
```

## more complex test with template inheritance (if possible)
### full featured template engines
| Name              | Runs    | µs/op  | B/op  | allocations/op |
| ----------------- | ------- | ------ | ----- | -------------- |
| ComplexGolang     | 58,860  | 68.765 | 6,565 | 290            |
| ComplexGolangText | 112,902 | 30.105 | 2,236 | 107            |
| ComplexJetHTML    | 374,076 | 12.292 | 535   | 5              |
| ComplexMustache   | 149,064 | 26.179 | 7,276 | 156            |

```
comparing: go1.22.7 to go1.23.1
name                  old time/op    new time/op    delta
ComplexGolang-24        69.4µs ± 0%    68.8µs ± 0%  -0.95%
ComplexGolangText-24    29.0µs ± 0%    30.1µs ± 0%  +3.86%
ComplexMustache-24      25.7µs ± 0%    26.2µs ± 0%  +1.99%
ComplexJetHTML-24       12.2µs ± 0%    12.3µs ± 0%  +0.52%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        6.57kB ± 0%    6.57kB ± 0%   0.00%
ComplexGolangText-24    2.24kB ± 0%    2.24kB ± 0%   0.00%
ComplexMustache-24      7.28kB ± 0%    7.28kB ± 0%  +0.01%
ComplexJetHTML-24         534B ± 0%      535B ± 0%  +0.19%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           290 ± 0%       290 ± 0%   0.00%
ComplexGolangText-24       107 ± 0%       107 ± 0%   0.00%
ComplexMustache-24         156 ± 0%       156 ± 0%   0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%   0.00%
```

### precompilation to Go code
| Name                  | Runs      | µs/op | B/op  | allocations/op |
| --------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo            | 1,000,000 | 5.034 | 569   | 31             |
| ComplexFtmpl          | 469,363   | 7.254 | 3,536 | 38             |
| ComplexGoDirectBuffer | 6,702,937 | 0.520 | 0     | 0              |
| ComplexGorazor        | 606,170   | 5.572 | 3,688 | 24             |
| ComplexHero           | 3,914,155 | 0.926 | 0     | 0              |
| ComplexJade           | 4,583,869 | 0.742 | 0     | 0              |
| ComplexQuicktemplate  | 3,441,301 | 1.065 | 0     | 0              |
| ComplexTempl          | 1,325,791 | 2.781 | 408   | 11             |

```
comparing: go1.22.7 to go1.23.1
name                      old time/op    new time/op    delta
ComplexEgo-24               4.98µs ± 0%    5.03µs ± 0%  +1.12%
ComplexQuicktemplate-24     1.04µs ± 0%    1.06µs ± 0%  +2.60%
ComplexTempl-24             2.80µs ± 0%    2.78µs ± 0%  -0.54%
ComplexFtmpl-24             7.19µs ± 0%    7.25µs ± 0%  +0.86%
ComplexGorazor-24           5.47µs ± 0%    5.57µs ± 0%  +1.86%
ComplexHero-24               937ns ± 0%     926ns ± 0%  -1.11%
ComplexJade-24               738ns ± 0%     742ns ± 0%  +0.57%
ComplexGoDirectBuffer-24     520ns ± 0%     520ns ± 0%  +0.10%

name                      old alloc/op   new alloc/op   delta
ComplexEgo-24                 569B ± 0%      569B ± 0%   0.00%
ComplexQuicktemplate-24      0.00B          0.00B        0.00%
ComplexTempl-24               408B ± 0%      408B ± 0%   0.00%
ComplexFtmpl-24             3.54kB ± 0%    3.54kB ± 0%   0.00%
ComplexGorazor-24           3.69kB ± 0%    3.69kB ± 0%   0.00%
ComplexHero-24               0.00B          0.00B        0.00%
ComplexJade-24               0.00B          0.00B        0.00%
ComplexGoDirectBuffer-24     0.00B          0.00B        0.00%

name                      old allocs/op  new allocs/op  delta
ComplexEgo-24                 31.0 ± 0%      31.0 ± 0%   0.00%
ComplexQuicktemplate-24       0.00           0.00        0.00%
ComplexTempl-24               11.0 ± 0%      11.0 ± 0%   0.00%
ComplexFtmpl-24               38.0 ± 0%      38.0 ± 0%   0.00%
ComplexGorazor-24             24.0 ± 0%      24.0 ± 0%   0.00%
ComplexHero-24                0.00           0.00        0.00%
ComplexJade-24                0.00           0.00        0.00%
ComplexGoDirectBuffer-24      0.00           0.00        0.00%
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
