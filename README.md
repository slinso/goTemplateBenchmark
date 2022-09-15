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
| ComplexGoDirectBuffer | 7,195,976   | 488   | 0    | 0              |
| ComplexGoStaticString | 312,416,414 | 12    | 0    | 0              |

```
comparing: go1.18.6 to go version go1.19 linux/amd64
name                      old time/op    new time/op    delta
ComplexGoDirectBuffer-24     491ns ± 0%     488ns ± 0%  -0.67%
ComplexGoStaticString-24    11.7ns ± 0%    11.7ns ± 0%  -0.09%

name                      old alloc/op   new alloc/op   delta
ComplexGoDirectBuffer-24     0.00B          0.00B        0.00%
ComplexGoStaticString-24     0.00B          0.00B        0.00%

name                      old allocs/op  new allocs/op  delta
ComplexGoDirectBuffer-24      0.00           0.00        0.00%
ComplexGoStaticString-24      0.00           0.00        0.00%
```

## simple benchmarks

### full featured template engines

| Name       | Runs      | µs/op  | B/op  | allocations/op |
| ---------- | --------- | ------ | ----- | -------------- |
| Ace        | 277,736   | 12.943 | 1,121 | 40             |
| Amber      | 384,963   | 8.808  | 849   | 36             |
| Golang     | 659,941   | 8.628  | 769   | 35             |
| GolangText | 1,435,208 | 2.516  | 128   | 7              |
| Handlebars | 253,663   | 14.007 | 3,648 | 78             |
| JetHTML    | 4,248,274 | 0.828  | 0     | 0              |
| Mustache   | 753,109   | 4.546  | 1,723 | 30             |
| Pongo2     | 520,420   | 5.975  | 2,075 | 32             |
| Soy        | 913,322   | 3.786  | 1,224 | 19             |

```
comparing: go1.18.6 to go version go1.19 linux/amd64
name           old time/op    new time/op    delta
Golang-24        8.71µs ± 0%    8.63µs ± 0%   -0.88%
GolangText-24    2.51µs ± 0%    2.52µs ± 0%   +0.20%
Ace-24           13.6µs ± 0%    12.9µs ± 0%   -4.56%
Amber-24         9.04µs ± 0%    8.81µs ± 0%   -2.58%
Mustache-24      4.21µs ± 0%    4.55µs ± 0%   +7.96%
Pongo2-24        5.88µs ± 0%    5.98µs ± 0%   +1.63%
Handlebars-24    14.1µs ± 0%    14.0µs ± 0%   -0.58%
Soy-24           3.63µs ± 0%    3.79µs ± 0%   +4.38%
JetHTML-24        900ns ± 0%     828ns ± 0%   -8.11%

name           old alloc/op   new alloc/op   delta
Golang-24          769B ± 0%      769B ± 0%    0.00%
GolangText-24      128B ± 0%      128B ± 0%    0.00%
Ace-24           1.12kB ± 0%    1.12kB ± 0%    0.00%
Amber-24           849B ± 0%      849B ± 0%    0.00%
Mustache-24      1.53kB ± 0%    1.72kB ± 0%  +12.61%
Pongo2-24        2.07kB ± 0%    2.08kB ± 0%   +0.05%
Handlebars-24    3.97kB ± 0%    3.65kB ± 0%   -8.02%
Soy-24           1.22kB ± 0%    1.22kB ± 0%    0.00%
JetHTML-24        0.00B          0.00B         0.00%

name           old allocs/op  new allocs/op  delta
Golang-24          35.0 ± 0%      35.0 ± 0%    0.00%
GolangText-24      7.00 ± 0%      7.00 ± 0%    0.00%
Ace-24             40.0 ± 0%      40.0 ± 0%    0.00%
Amber-24           36.0 ± 0%      36.0 ± 0%    0.00%
Mustache-24        29.0 ± 0%      30.0 ± 0%   +3.45%
Pongo2-24          32.0 ± 0%      32.0 ± 0%    0.00%
Handlebars-24      78.0 ± 0%      78.0 ± 0%    0.00%
Soy-24             19.0 ± 0%      19.0 ± 0%    0.00%
JetHTML-24         0.00           0.00         0.00%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op | allocations/op |
| ------------- | ---------- | ----- | ---- | -------------- |
| Ego           | 3,330,222  | 1.138 | 85   | 8              |
| Ftmpl         | 2,247,235  | 1.603 | 774  | 12             |
| Gorazor       | 4,196,112  | 0.844 | 512  | 5              |
| Hero          | 25,661,955 | 0.122 | 0    | 0              |
| Jade          | 38,546,743 | 0.093 | 0    | 0              |
| Quicktemplate | 20,088,501 | 0.181 | 0    | 0              |

```
comparing: go1.18.6 to go version go1.19 linux/amd64
name              old time/op    new time/op    delta
Ego-24              1.24µs ± 0%    1.14µs ± 0%   -8.15%
Quicktemplate-24     190ns ± 0%     181ns ± 0%   -4.53%
Ftmpl-24            1.67µs ± 0%    1.60µs ± 0%   -4.01%
Gorazor-24           823ns ± 0%     844ns ± 0%   +2.56%
Hero-24              126ns ± 0%     122ns ± 0%   -3.56%
Jade-24             92.7ns ± 0%    92.5ns ± 0%   -0.15%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%    0.00%
Quicktemplate-24     0.00B          0.00B         0.00%
Ftmpl-24            1.09kB ± 0%    0.77kB ± 0%  -29.25%
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
| ComplexGolang     | 49,461  | 76.020 | 6,645 | 290            |
| ComplexGolangText | 100,723 | 33.766 | 2,236 | 107            |
| ComplexJetHTML    | 304,092 | 13.188 | 534   | 5              |
| ComplexMustache   | 125,402 | 28.236 | 7,274 | 156            |

```
comparing: go1.18.6 to go version go1.19 linux/amd64
name                  old time/op    new time/op    delta
ComplexGolang-24        75.7µs ± 0%    76.0µs ± 0%  +0.44%
ComplexGolangText-24    33.4µs ± 0%    33.8µs ± 0%  +1.00%
ComplexMustache-24      28.1µs ± 0%    28.2µs ± 0%  +0.39%
ComplexJetHTML-24       13.0µs ± 0%    13.2µs ± 0%  +1.34%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        6.64kB ± 0%    6.64kB ± 0%  +0.03%
ComplexGolangText-24    2.23kB ± 0%    2.24kB ± 0%  +0.04%
ComplexMustache-24      7.40kB ± 0%    7.27kB ± 0%  -1.66%
ComplexJetHTML-24         533B ± 0%      534B ± 0%  +0.19%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           290 ± 0%       290 ± 0%   0.00%
ComplexGolangText-24       107 ± 0%       107 ± 0%   0.00%
ComplexMustache-24         155 ± 0%       156 ± 0%  +0.65%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%   0.00%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 1,000,000 | 5.765 | 568   | 31             |
| ComplexFtmpl         | 507,601   | 7.100 | 3,535 | 38             |
| ComplexGorazor       | 669,070   | 5.338 | 3,728 | 22             |
| ComplexHero          | 3,630,922 | 0.961 | 0     | 0              |
| ComplexJade          | 4,504,540 | 0.771 | 0     | 0              |
| ComplexQuicktemplate | 3,619,738 | 0.997 | 0     | 0              |

```
comparing: go1.18.6 to go version go1.19 linux/amd64
name                     old time/op    new time/op    delta
ComplexEgo-24              5.82µs ± 0%    5.76µs ± 0%   -0.86%
ComplexQuicktemplate-24    1.39µs ± 0%    1.00µs ± 0%  -28.10%
ComplexFtmpl-24            7.52µs ± 0%    7.10µs ± 0%   -5.56%
ComplexGorazor-24          4.99µs ± 0%    5.34µs ± 0%   +6.97%
ComplexHero-24             1.06µs ± 0%    0.96µs ± 0%   -8.88%
ComplexJade-24              779ns ± 0%     771ns ± 0%   -1.05%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                568B ± 0%      568B ± 0%    0.00%
ComplexQuicktemplate-24     0.00B          0.00B         0.00%
ComplexFtmpl-24            4.91kB ± 0%    3.54kB ± 0%  -28.03%
ComplexGorazor-24          3.74kB ± 0%    3.73kB ± 0%   -0.43%
ComplexHero-24              0.00B          0.00B         0.00%
ComplexJade-24              0.00B          0.00B         0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                31.0 ± 0%      31.0 ± 0%    0.00%
ComplexQuicktemplate-24      0.00           0.00         0.00%
ComplexFtmpl-24              38.0 ± 0%      38.0 ± 0%    0.00%
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
