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
| Ace        | 255,001   | 12.738 | 1,121 | 40             |
| Amber      | 385,418   | 8.304  | 849   | 36             |
| Golang     | 495,836   | 8.168  | 769   | 35             |
| GolangText | 1,588,129 | 2.430  | 128   | 7              |
| Handlebars | 260,066   | 12.838 | 3,424 | 75             |
| JetHTML    | 4,728,830 | 0.771  | 0     | 0              |
| Mustache   | 869,660   | 4.276  | 1,723 | 30             |
| Pongo2     | 919,458   | 5.437  | 2,075 | 32             |
| Soy        | 1,000,000 | 3.399  | 1,224 | 19             |

```
comparing: go1.22.7 to go1.23.1
name           old time/op    new time/op    delta
Golang-24        8.04µs ± 0%    8.17µs ± 0%  +1.62%
GolangText-24    2.39µs ± 0%    2.43µs ± 0%  +1.63%
Ace-24           12.7µs ± 0%    12.7µs ± 0%  +0.69%
Amber-24         8.20µs ± 0%    8.30µs ± 0%  +1.27%
Mustache-24      4.31µs ± 0%    4.28µs ± 0%  -0.72%
Pongo2-24        5.64µs ± 0%    5.44µs ± 0%  -3.51%
Handlebars-24    13.4µs ± 0%    12.8µs ± 0%  -4.13%
Soy-24           3.45µs ± 0%    3.40µs ± 0%  -1.56%
JetHTML-24        769ns ± 0%     771ns ± 0%  +0.27%

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
| Ego           | 4,065,416  | 0.942 | 85    | 8              |
| Ftmpl         | 2,238,477  | 1.602 | 774   | 12             |
| Gomponents    | 1,000,000  | 4.785 | 1,112 | 56             |
| Gorazor       | 4,340,546  | 0.831 | 512   | 5              |
| Hero          | 30,667,665 | 0.116 | 0     | 0              |
| Jade          | 38,743,312 | 0.090 | 0     | 0              |
| Quicktemplate | 16,765,074 | 0.188 | 0     | 0              |

```
comparing: go1.22.7 to go1.23.1
name              old time/op    new time/op    delta
Ego-24              1.02µs ± 0%    0.94µs ± 0%  -7.67%
Quicktemplate-24     194ns ± 0%     188ns ± 0%  -2.99%
Ftmpl-24            1.66µs ± 0%    1.60µs ± 0%  -3.55%
Gorazor-24           847ns ± 0%     831ns ± 0%  -1.84%
Hero-24              125ns ± 0%     116ns ± 0%  -6.96%
Jade-24             88.8ns ± 0%    89.7ns ± 0%  +0.96%
Gomponents-24       4.72µs ± 0%    4.79µs ± 0%  +1.44%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%   0.00%
Quicktemplate-24     0.00B          0.00B        0.00%
Ftmpl-24              774B ± 0%      774B ± 0%   0.00%
Gorazor-24            512B ± 0%      512B ± 0%   0.00%
Hero-24              0.00B          0.00B        0.00%
Jade-24              0.00B          0.00B        0.00%
Gomponents-24       1.11kB ± 0%    1.11kB ± 0%   0.00%

name              old allocs/op  new allocs/op  delta
Ego-24                8.00 ± 0%      8.00 ± 0%   0.00%
Quicktemplate-24      0.00           0.00        0.00%
Ftmpl-24              12.0 ± 0%      12.0 ± 0%   0.00%
Gorazor-24            5.00 ± 0%      5.00 ± 0%   0.00%
Hero-24               0.00           0.00        0.00%
Jade-24               0.00           0.00        0.00%
Gomponents-24         56.0 ± 0%      56.0 ± 0%   0.00%
```

## more complex test with template inheritance (if possible)
### full featured template engines
| Name              | Runs    | µs/op  | B/op  | allocations/op |
| ----------------- | ------- | ------ | ----- | -------------- |
| ComplexGolang     | 56,319  | 67.173 | 6,565 | 290            |
| ComplexGolangText | 129,277 | 29.535 | 2,236 | 107            |
| ComplexJetHTML    | 354,434 | 11.799 | 535   | 5              |
| ComplexMustache   | 140,676 | 25.785 | 7,275 | 156            |

```
comparing: go1.22.7 to go1.23.1
name                  old time/op    new time/op    delta
ComplexGolang-24        67.1µs ± 0%    67.2µs ± 0%  +0.06%
ComplexGolangText-24    29.4µs ± 0%    29.5µs ± 0%  +0.30%
ComplexMustache-24      25.4µs ± 0%    25.8µs ± 0%  +1.49%
ComplexJetHTML-24       12.4µs ± 0%    11.8µs ± 0%  -4.89%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        6.55kB ± 0%    6.57kB ± 0%  +0.24%
ComplexGolangText-24    2.24kB ± 0%    2.24kB ± 0%   0.00%
ComplexMustache-24      7.28kB ± 0%    7.28kB ± 0%  -0.01%
ComplexJetHTML-24         535B ± 0%      535B ± 0%   0.00%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           290 ± 0%       290 ± 0%   0.00%
ComplexGolangText-24       107 ± 0%       107 ± 0%   0.00%
ComplexMustache-24         156 ± 0%       156 ± 0%   0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%   0.00%
```

### precompilation to Go code
| Name                  | Runs      | µs/op | B/op  | allocations/op |
| --------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo            | 1,000,000 | 4.817 | 569   | 31             |
| ComplexFtmpl          | 567,518   | 7.049 | 3,536 | 38             |
| ComplexGorazor        | 606,054   | 5.445 | 3,688 | 24             |
| ComplexHero           | 3,695,972 | 0.947 | 0     | 0              |
| ComplexJade           | 4,579,771 | 0.743 | 0     | 0              |
| ComplexQuicktemplate  | 3,465,189 | 1.033 | 0     | 0              |
| ComplexGoDirectBuffer | 6,758,412 | 0.531 | 0     | 0              |

```
comparing: go1.22.7 to go1.23.1
name                     old time/op    new time/op    delta
ComplexEgo-24              5.03µs ± 0%    4.82µs ± 0%  -4.16%
ComplexQuicktemplate-24    1.06µs ± 0%    1.03µs ± 0%  -2.55%
ComplexFtmpl-24            7.25µs ± 0%    7.05µs ± 0%  -2.76%
ComplexGorazor-24          5.40µs ± 0%    5.45µs ± 0%  +0.91%
ComplexHero-24              954ns ± 0%     947ns ± 0%  -0.72%
ComplexJade-24              736ns ± 0%     742ns ± 0%  +0.90%
ComplexGoDirectBuffer-24    533ns ± 0%     530ns ± 0%  -0.41%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                569B ± 0%      569B ± 0%   0.00%
ComplexQuicktemplate-24     0.00B          0.00B        0.00%
ComplexFtmpl-24            3.54kB ± 0%    3.54kB ± 0%   0.00%
ComplexGorazor-24          3.69kB ± 0%    3.69kB ± 0%   0.00%
ComplexHero-24              0.00B          0.00B        0.00%
ComplexJade-24              0.00B          0.00B        0.00%
ComplexGoDirectBuffer-24    0.00B          0.00B        0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                31.0 ± 0%      31.0 ± 0%   0.00%
ComplexQuicktemplate-24      0.00           0.00        0.00%
ComplexFtmpl-24              38.0 ± 0%      38.0 ± 0%   0.00%
ComplexGorazor-24            24.0 ± 0%      24.0 ± 0%   0.00%
ComplexHero-24               0.00           0.00        0.00%
ComplexJade-24               0.00           0.00        0.00%
ComplexGoDirectBuffer-24     0.00           0.00        0.00%
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
