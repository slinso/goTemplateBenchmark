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
| ComplexGoDirectBuffer | 7,715,667   | 456   | 0    | 0              |
| ComplexGoStaticString | 296,467,746 | 12    | 0    | 0              |

```
comparing: go1.15.8 to go version go1.16 linux/amd64
name                      old time/op    new time/op    delta
ComplexGoDirectBuffer-24     472ns ± 0%     456ns ± 0%  -3.45%
ComplexGoStaticString-24    12.4ns ± 0%    11.7ns ± 0%  -5.65%

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
| Ace         | 287,395   | 12.796 | 1,249 | 40             |
| Amber       | 461,578   | 8.303  | 977   | 36             |
| Golang      | 466,560   | 7.928  | 897   | 35             |
| GolangText  | 1,460,233 | 2.582  | 128   | 7              |
| Handlebars  | 269,463   | 14.150 | 3,967 | 78             |
| **JetHTML** | 3,455,011 | 1.008  | 0     | 0              |
| Mustache    | 834,798   | 4.270  | 1,530 | 29             |
| Pongo2      | 589,389   | 5.933  | 2,074 | 32             |
| Soy         | 929,354   | 3.935  | 1,320 | 20             |

```
comparing: go1.15.8 to go version go1.16 linux/amd64
name           old time/op    new time/op    delta
Golang-24        8.19µs ± 0%    7.93µs ± 0%   -3.14%
GolangText-24    2.54µs ± 0%    2.58µs ± 0%   +1.49%
Ace-24           13.0µs ± 0%    12.8µs ± 0%   -1.39%
Amber-24         8.62µs ± 0%    8.30µs ± 0%   -3.63%
Mustache-24      4.16µs ± 0%    4.27µs ± 0%   +2.74%
Pongo2-24        5.79µs ± 0%    5.93µs ± 0%   +2.40%
Handlebars-24    13.9µs ± 0%    14.2µs ± 0%   +1.67%
Soy-24           3.93µs ± 0%    3.93µs ± 0%   +0.20%
JetHTML-24       0.99µs ± 0%    1.01µs ± 0%   +1.31%

name           old alloc/op   new alloc/op   delta
Golang-24        1.02kB ± 0%    0.90kB ± 0%  -12.49%
GolangText-24      128B ± 0%      128B ± 0%    0.00%
Ace-24           1.38kB ± 0%    1.25kB ± 0%   -9.36%
Amber-24         1.10kB ± 0%    0.98kB ± 0%  -11.58%
Mustache-24      1.57kB ± 0%    1.53kB ± 0%   -2.55%
Pongo2-24        2.07kB ± 0%    2.07kB ± 0%    0.00%
Handlebars-24    3.98kB ± 0%    3.97kB ± 0%   -0.40%
Soy-24           1.35kB ± 0%    1.32kB ± 0%   -2.37%
JetHTML-24        0.00B          0.00B         0.00%

name           old allocs/op  new allocs/op  delta
Golang-24          35.0 ± 0%      35.0 ± 0%    0.00%
GolangText-24      7.00 ± 0%      7.00 ± 0%    0.00%
Ace-24             40.0 ± 0%      40.0 ± 0%    0.00%
Amber-24           36.0 ± 0%      36.0 ± 0%    0.00%
Mustache-24        29.0 ± 0%      29.0 ± 0%    0.00%
Pongo2-24          32.0 ± 0%      32.0 ± 0%    0.00%
Handlebars-24      78.0 ± 0%      78.0 ± 0%    0.00%
Soy-24             20.0 ± 0%      20.0 ± 0%    0.00%
JetHTML-24         0.00           0.00         0.00%
```

### precompilation to Go code

| Name          | Runs       | µs/op | B/op  | allocations/op |
| ------------- | ---------- | ----- | ----- | -------------- |
| Ego           | 2,782,954  | 1.314 | 85    | 8              |
| EgonSlinso    | 10,311,388 | 0.305 | 0     | 0              |
| Ftmpl         | 2,095,531  | 1.710 | 1,095 | 12             |
| Gorazor       | 4,343,186  | 0.830 | 512   | 5              |
| Hero          | 27,075,139 | 0.135 | 0     | 0              |
| **Jade**      | 41,323,922 | 0.088 | 0     | 0              |
| Quicktemplate | 13,077,282 | 0.246 | 0     | 0              |

```
comparing: go1.15.8 to go version go1.16 linux/amd64
name              old time/op    new time/op    delta
Ego-24              1.30µs ± 0%    1.31µs ± 0%  +1.47%
EgonSlinso-24        287ns ± 0%     305ns ± 0%  +6.31%
Quicktemplate-24     224ns ± 0%     246ns ± 0%  +9.96%
Ftmpl-24            1.90µs ± 0%    1.71µs ± 0%  -9.81%
Gorazor-24           794ns ± 0%     830ns ± 0%  +4.57%
Hero-24              141ns ± 0%     135ns ± 0%  -4.26%
Jade-24             84.1ns ± 0%    87.8ns ± 0%  +4.42%

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
| ComplexGolang      | 49,890  | 73.015 | 7,774 | 285            |
| ComplexGolangText  | 123,493 | 31.905 | 2,532 | 102            |
| **ComplexJetHTML** | 227,565 | 14.653 | 534   | 5              |
| ComplexMustache    | 129,079 | 27.913 | 7,400 | 155            |

```
comparing: go1.15.8 to go version go1.16 linux/amd64
name                  old time/op    new time/op    delta
ComplexGolang-24        73.1µs ± 0%    73.0µs ± 0%   -0.13%
ComplexGolangText-24    32.2µs ± 0%    31.9µs ± 0%   -0.81%
ComplexMustache-24      27.7µs ± 0%    27.9µs ± 0%   +0.93%
ComplexJetHTML-24       15.8µs ± 0%    14.7µs ± 0%   -7.41%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        8.78kB ± 0%    7.77kB ± 0%  -11.49%
ComplexGolangText-24    2.71kB ± 0%    2.53kB ± 0%   -6.50%
ComplexMustache-24      7.57kB ± 0%    7.40kB ± 0%   -2.22%
ComplexJetHTML-24         550B ± 0%      534B ± 0%   -2.91%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           285 ± 0%       285 ± 0%    0.00%
ComplexGolangText-24       102 ± 0%       102 ± 0%    0.00%
ComplexMustache-24         155 ± 0%       155 ± 0%    0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%    0.00%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 536,416   | 6.340 | 568   | 31             |
| ComplexEgoSlinso     | 1,000,000 | 3.083 | 160   | 2              |
| ComplexFtmpl         | 480,680   | 7.638 | 4,912 | 38             |
| ComplexGorazor       | 832,575   | 4.436 | 2,720 | 21             |
| ComplexHero          | 3,439,156 | 1.000 | 0     | 0              |
| **ComplexJade**      | 4,865,839 | 0.705 | 0     | 0              |
| ComplexQuicktemplate | 2,812,012 | 1.258 | 0     | 0              |

```
comparing: go1.15.8 to go version go1.16 linux/amd64
name                     old time/op    new time/op    delta
ComplexEgo-24              6.55µs ± 0%    6.34µs ± 0%  -3.27%
ComplexQuicktemplate-24    1.18µs ± 0%    1.26µs ± 0%  +6.52%
ComplexEgoSlinso-24        2.83µs ± 0%    3.08µs ± 0%  +8.94%
ComplexFtmpl-24            7.69µs ± 0%    7.64µs ± 0%  -0.71%
ComplexGorazor-24          4.54µs ± 0%    4.44µs ± 0%  -2.33%
ComplexHero-24             1.01µs ± 0%    1.00µs ± 0%  -0.89%
ComplexJade-24              706ns ± 0%     705ns ± 0%  -0.08%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                592B ± 0%      568B ± 0%  -4.05%
ComplexQuicktemplate-24     0.00B          0.00B        0.00%
ComplexEgoSlinso-24          160B ± 0%      160B ± 0%   0.00%
ComplexFtmpl-24            4.94kB ± 0%    4.91kB ± 0%  -0.49%
ComplexGorazor-24          2.87kB ± 0%    2.72kB ± 0%  -5.29%
ComplexHero-24              0.00B          0.00B        0.00%
ComplexJade-24              0.00B          0.00B        0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                31.0 ± 0%      31.0 ± 0%   0.00%
ComplexQuicktemplate-24      0.00           0.00        0.00%
ComplexEgoSlinso-24          2.00 ± 0%      2.00 ± 0%   0.00%
ComplexFtmpl-24              38.0 ± 0%      38.0 ± 0%   0.00%
ComplexGorazor-24            22.0 ± 0%      21.0 ± 0%  -4.55%
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
