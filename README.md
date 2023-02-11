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

## special benchmarks

| Name                  | Runs        | ns/op | B/op | allocations/op |
| --------------------- | ----------- | ----- | ---- | -------------- |
| ComplexGoDirectBuffer | 8,153,427   | 428   | 0    | 0              |
| ComplexGoStaticString | 310,142,265 | 12    | 0    | 0              |

```
comparing: go1.19.5 to go1.20
name                      old time/op    new time/op    delta
ComplexGoDirectBuffer-24     519ns ± 0%     428ns ± 0%  -17.63%
ComplexGoStaticString-24    11.7ns ± 0%    11.9ns ± 0%   +1.96%

name                      old alloc/op   new alloc/op   delta
ComplexGoDirectBuffer-24     0.00B          0.00B         0.00%
ComplexGoStaticString-24     0.00B          0.00B         0.00%

name                      old allocs/op  new allocs/op  delta
ComplexGoDirectBuffer-24      0.00           0.00         0.00%
ComplexGoStaticString-24      0.00           0.00         0.00%
```

## simple benchmarks

### full featured template engines

| Name       | Runs      | µs/op  | B/op  | allocations/op |
| ---------- | --------- | ------ | ----- | -------------- |
| Ace        | 289,507   | 12.165 | 1,121 | 40             |
| Amber      | 381,030   | 8.535  | 849   | 36             |
| Golang     | 694,083   | 8.197  | 769   | 35             |
| GolangText | 1,500,374 | 2.355  | 128   | 7              |
| Handlebars | 266,238   | 13.062 | 3,648 | 78             |
| JetHTML    | 4,429,336 | 0.771  | 0     | 0              |
| Mustache   | 859,552   | 4.177  | 1,723 | 30             |
| Pongo2     | 840,604   | 5.441  | 2,075 | 32             |
| Soy        | 1,000,000 | 3.610  | 1,224 | 19             |

```
comparing: go1.19.5 to go1.20
name           old time/op    new time/op    delta
Golang-24        8.42µs ± 0%    8.20µs ± 0%  -2.61%
GolangText-24    2.47µs ± 0%    2.35µs ± 0%  -4.81%
Ace-24           12.8µs ± 0%    12.2µs ± 0%  -4.75%
Amber-24         8.57µs ± 0%    8.54µs ± 0%  -0.42%
Mustache-24      4.46µs ± 0%    4.18µs ± 0%  -6.41%
Pongo2-24        5.91µs ± 0%    5.44µs ± 0%  -7.98%
Handlebars-24    13.9µs ± 0%    13.1µs ± 0%  -6.26%
Soy-24           3.67µs ± 0%    3.61µs ± 0%  -1.77%
JetHTML-24        832ns ± 0%     771ns ± 0%  -7.39%

name           old alloc/op   new alloc/op   delta
Golang-24          769B ± 0%      769B ± 0%   0.00%
GolangText-24      128B ± 0%      128B ± 0%   0.00%
Ace-24           1.12kB ± 0%    1.12kB ± 0%   0.00%
Amber-24           849B ± 0%      849B ± 0%   0.00%
Mustache-24      1.72kB ± 0%    1.72kB ± 0%   0.00%
Pongo2-24        2.08kB ± 0%    2.08kB ± 0%   0.00%
Handlebars-24    3.65kB ± 0%    3.65kB ± 0%   0.00%
Soy-24           1.22kB ± 0%    1.22kB ± 0%   0.00%
JetHTML-24        0.00B          0.00B        0.00%

name           old allocs/op  new allocs/op  delta
Golang-24          35.0 ± 0%      35.0 ± 0%   0.00%
GolangText-24      7.00 ± 0%      7.00 ± 0%   0.00%
Ace-24             40.0 ± 0%      40.0 ± 0%   0.00%
Amber-24           36.0 ± 0%      36.0 ± 0%   0.00%
Mustache-24        30.0 ± 0%      30.0 ± 0%   0.00%
Pongo2-24          32.0 ± 0%      32.0 ± 0%   0.00%
Handlebars-24      78.0 ± 0%      78.0 ± 0%   0.00%
Soy-24             19.0 ± 0%      19.0 ± 0%   0.00%
JetHTML-24         0.00           0.00        0.00%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op | allocations/op |
| ------------- | ---------- | ----- | ---- | -------------- |
| Ego           | 3,272,616  | 1.151 | 85   | 8              |
| Ftmpl         | 2,179,855  | 1.610 | 774  | 12             |
| Gorazor       | 4,294,610  | 0.803 | 512  | 5              |
| Hero          | 24,494,606 | 0.123 | 0    | 0              |
| Jade          | 40,706,695 | 0.089 | 0    | 0              |
| Quicktemplate | 19,722,939 | 0.181 | 0    | 0              |

```
comparing: go1.19.5 to go1.20
name              old time/op    new time/op    delta
Ego-24              1.13µs ± 0%    1.15µs ± 0%   +2.22%
Quicktemplate-24     185ns ± 0%     181ns ± 0%   -1.95%
Ftmpl-24            1.46µs ± 0%    1.61µs ± 0%  +10.05%
Gorazor-24           804ns ± 0%     803ns ± 0%   -0.06%
Hero-24              119ns ± 0%     123ns ± 0%   +3.20%
Jade-24             91.5ns ± 0%    88.8ns ± 0%   -2.86%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%    0.00%
Quicktemplate-24     0.00B          0.00B         0.00%
Ftmpl-24              774B ± 0%      774B ± 0%    0.00%
Gorazor-24            512B ± 0%      512B ± 0%    0.00%
Hero-24              0.00B          0.00B         0.00%
Jade-24              0.00B          0.00B         0.00%

name              old allocs/op  new allocs/op  delta
Ego-24                8.00 ± 0%      8.00 ± 0%    0.00%
Quicktemplate-24      0.00           0.00         0.00%
Ftmpl-24              12.0 ± 0%      12.0 ± 0%    0.00%
Gorazor-24            5.00 ± 0%      5.00 ± 0%    0.00%
Hero-24               0.00           0.00         0.00%
Jade-24               0.00           0.00         0.00%
```

## more complex test with template inheritance (if possible)

### full featured template engines

| Name              | Runs    | µs/op  | B/op  | allocations/op |
| ----------------- | ------- | ------ | ----- | -------------- |
| ComplexGolang     | 55,249  | 72.293 | 6,548 | 290            |
| ComplexGolangText | 115,640 | 31.731 | 2,236 | 107            |
| ComplexJetHTML    | 313,855 | 11.832 | 534   | 5              |
| ComplexMustache   | 134,226 | 26.156 | 7,274 | 156            |

```
comparing: go1.19.5 to go1.20
name                  old time/op    new time/op    delta
ComplexGolang-24        72.8µs ± 0%    72.3µs ± 0%  -0.73%
ComplexGolangText-24    31.8µs ± 0%    31.7µs ± 0%  -0.20%
ComplexMustache-24      26.7µs ± 0%    26.2µs ± 0%  -2.07%
ComplexJetHTML-24       11.7µs ± 0%    11.8µs ± 0%  +1.13%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        6.64kB ± 0%    6.55kB ± 0%  -1.46%
ComplexGolangText-24    2.24kB ± 0%    2.24kB ± 0%   0.00%
ComplexMustache-24      7.27kB ± 0%    7.27kB ± 0%  +0.01%
ComplexJetHTML-24         534B ± 0%      534B ± 0%   0.00%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           290 ± 0%       290 ± 0%   0.00%
ComplexGolangText-24       107 ± 0%       107 ± 0%   0.00%
ComplexMustache-24         156 ± 0%       156 ± 0%   0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%   0.00%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 837,608   | 5.777 | 568   | 31             |
| ComplexFtmpl         | 486,427   | 7.126 | 3,535 | 38             |
| ComplexGorazor       | 720,013   | 5.193 | 3,688 | 24             |
| ComplexHero          | 3,623,186 | 0.955 | 0     | 0              |
| ComplexJade          | 5,046,976 | 0.745 | 0     | 0              |
| ComplexQuicktemplate | 3,481,230 | 1.011 | 0     | 0              |

```
comparing: go1.19.5 to go1.20
name                     old time/op    new time/op    delta
ComplexEgo-24              5.67µs ± 0%    5.78µs ± 0%  +1.91%
ComplexQuicktemplate-24    1.03µs ± 0%    1.01µs ± 0%  -2.22%
ComplexFtmpl-24            7.08µs ± 0%    7.13µs ± 0%  +0.65%
ComplexGorazor-24          5.09µs ± 0%    5.19µs ± 0%  +2.10%
ComplexHero-24              964ns ± 0%     955ns ± 0%  -0.92%
ComplexJade-24              764ns ± 0%     745ns ± 0%  -2.47%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                568B ± 0%      568B ± 0%   0.00%
ComplexQuicktemplate-24     0.00B          0.00B        0.00%
ComplexFtmpl-24            3.54kB ± 0%    3.54kB ± 0%   0.00%
ComplexGorazor-24          3.73kB ± 0%    3.69kB ± 0%  -1.07%
ComplexHero-24              0.00B          0.00B        0.00%
ComplexJade-24              0.00B          0.00B        0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                31.0 ± 0%      31.0 ± 0%   0.00%
ComplexQuicktemplate-24      0.00           0.00        0.00%
ComplexFtmpl-24              38.0 ± 0%      38.0 ± 0%   0.00%
ComplexGorazor-24            22.0 ± 0%      24.0 ± 0%  +9.09%
ComplexHero-24               0.00           0.00        0.00%
ComplexJade-24               0.00           0.00        0.00%
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
