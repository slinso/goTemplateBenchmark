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

Sometimes there are templates that cannot be reasonably cached. Then you might
need a really fast template engine with code generation.

## Results dev machine

local desktop: ryzen 3900x

## special benchmarks

| Name                  | Runs        | ns/op | B/op | allocations/op |
| --------------------- | ----------- | ----- | ---- | -------------- |
| ComplexGoDirectBuffer | 8,293,862   | 432   | 0    | 0              |
| ComplexGoStaticString | 296,413,996 | 12    | 0    | 0              |

```
comparing: go1.16.7 to go version go1.17 linux/amd64
name                      old time/op    new time/op    delta
ComplexGoDirectBuffer-24     457ns ± 0%     432ns ± 0%  -5.36%
ComplexGoStaticString-24    12.5ns ± 0%    11.7ns ± 0%  -6.32%

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
| Ace        | 281,271   | 12.762 | 1,121 | 40             |
| Amber      | 465,156   | 8.673  | 849   | 36             |
| Golang     | 428,677   | 8.380  | 769   | 35             |
| GolangText | 1,337,016 | 2.641  | 128   | 7              |
| Handlebars | 251,836   | 14.019 | 3,967 | 78             |
| JetHTML    | 3,901,444 | 0.866  | 0     | 0              |
| Mustache   | 842,256   | 4.163  | 1,530 | 29             |
| Pongo2     | 599,228   | 5.833  | 2,074 | 32             |
| Soy        | 945,814   | 3.752  | 1,320 | 20             |

```
comparing: go1.16.7 to go version go1.17 linux/amd64
name           old time/op    new time/op    delta
Golang-24        8.07µs ± 0%    8.38µs ± 0%   +3.84%
GolangText-24    2.60µs ± 0%    2.64µs ± 0%   +1.54%
Ace-24           12.7µs ± 0%    12.8µs ± 0%   +0.22%
Amber-24         8.37µs ± 0%    8.67µs ± 0%   +3.58%
Mustache-24      4.25µs ± 0%    4.16µs ± 0%   -2.05%
Pongo2-24        6.03µs ± 0%    5.83µs ± 0%   -3.25%
Handlebars-24    14.0µs ± 0%    14.0µs ± 0%   -0.01%
Soy-24           4.05µs ± 0%    3.75µs ± 0%   -7.34%
JetHTML-24       1.00µs ± 0%    0.87µs ± 0%  -13.21%

name           old alloc/op   new alloc/op   delta
Golang-24          897B ± 0%      769B ± 0%  -14.27%
GolangText-24      128B ± 0%      128B ± 0%    0.00%
Ace-24           1.25kB ± 0%    1.12kB ± 0%  -10.25%
Amber-24           977B ± 0%      849B ± 0%  -13.10%
Mustache-24      1.53kB ± 0%    1.53kB ± 0%    0.00%
Pongo2-24        2.07kB ± 0%    2.07kB ± 0%    0.00%
Handlebars-24    3.97kB ± 0%    3.97kB ± 0%    0.00%
Soy-24           1.32kB ± 0%    1.32kB ± 0%    0.00%
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
| Ego           | 2,917,192  | 1.214 | 85    | 8              |
| EgonSlinso    | 18,751,664 | 0.192 | 0     | 0              |
| Ftmpl         | 2,187,069  | 1.670 | 1,095 | 12             |
| Gorazor       | 4,525,616  | 0.793 | 512   | 5              |
| Hero          | 28,386,926 | 0.127 | 0     | 0              |
| Jade          | 40,868,552 | 0.086 | 0     | 0              |
| Quicktemplate | 12,820,016 | 0.257 | 0     | 0              |

```
comparing: go1.16.7 to go version go1.17 linux/amd64
name              old time/op    new time/op    delta
Ego-24              1.32µs ± 0%    1.21µs ± 0%   -8.10%
EgonSlinso-24        315ns ± 0%     192ns ± 0%  -38.94%
Quicktemplate-24     263ns ± 0%     257ns ± 0%   -2.06%
Ftmpl-24            1.70µs ± 0%    1.67µs ± 0%   -1.76%
Gorazor-24           827ns ± 0%     793ns ± 0%   -4.12%
Hero-24              129ns ± 0%     127ns ± 0%   -1.24%
Jade-24             85.8ns ± 0%    86.2ns ± 0%   +0.43%

name              old alloc/op   new alloc/op   delta
Ego-24               85.0B ± 0%     85.0B ± 0%    0.00%
EgonSlinso-24        0.00B          0.00B         0.00%
Quicktemplate-24     0.00B          0.00B         0.00%
Ftmpl-24            1.09kB ± 0%    1.09kB ± 0%    0.00%
Gorazor-24            512B ± 0%      512B ± 0%    0.00%
Hero-24              0.00B          0.00B         0.00%
Jade-24              0.00B          0.00B         0.00%

name              old allocs/op  new allocs/op  delta
Ego-24                8.00 ± 0%      8.00 ± 0%    0.00%
EgonSlinso-24         0.00           0.00         0.00%
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
| ComplexGolang     | 49,048  | 72.884 | 6,643 | 290            |
| ComplexGolangText | 103,266 | 33.945 | 2,235 | 107            |
| ComplexJetHTML    | 239,065 | 14.866 | 534   | 5              |
| ComplexMustache   | 131,143 | 28.025 | 7,399 | 155            |

```
comparing: go1.16.7 to go version go1.17 linux/amd64
name                  old time/op    new time/op    delta
ComplexGolang-24        71.9µs ± 0%    72.9µs ± 0%   +1.42%
ComplexGolangText-24    30.6µs ± 0%    33.9µs ± 0%  +11.04%
ComplexMustache-24      27.7µs ± 0%    28.0µs ± 0%   +1.15%
ComplexJetHTML-24       14.8µs ± 0%    14.9µs ± 0%   +0.32%

name                  old alloc/op   new alloc/op   delta
ComplexGolang-24        7.77kB ± 0%    6.64kB ± 0%  -14.54%
ComplexGolangText-24    2.53kB ± 0%    2.23kB ± 0%  -11.73%
ComplexMustache-24      7.40kB ± 0%    7.40kB ± 0%   -0.01%
ComplexJetHTML-24         534B ± 0%      534B ± 0%    0.00%

name                  old allocs/op  new allocs/op  delta
ComplexGolang-24           285 ± 0%       290 ± 0%   +1.75%
ComplexGolangText-24       102 ± 0%       107 ± 0%   +4.90%
ComplexMustache-24         155 ± 0%       155 ± 0%    0.00%
ComplexJetHTML-24         5.00 ± 0%      5.00 ± 0%    0.00%
```

### precompilation to Go code

| Name                 | Runs      | µs/op | B/op  | allocations/op |
| -------------------- | --------- | ----- | ----- | -------------- |
| ComplexEgo           | 616,015   | 6.026 | 568   | 31             |
| ComplexEgoSlinso     | 1,475,832 | 2.414 | 160   | 2              |
| ComplexFtmpl         | 493,070   | 7.440 | 4,912 | 38             |
| ComplexGorazor       | 870,666   | 4.248 | 2,720 | 21             |
| ComplexHero          | 3,547,309 | 1.000 | 0     | 0              |
| ComplexJade          | 4,559,787 | 0.744 | 0     | 0              |
| ComplexQuicktemplate | 3,535,020 | 0.952 | 0     | 0              |

```
comparing: go1.16.7 to go version go1.17 linux/amd64
name                     old time/op    new time/op    delta
ComplexEgo-24              6.33µs ± 0%    6.03µs ± 0%   -4.85%
ComplexQuicktemplate-24    1.37µs ± 0%    0.95µs ± 0%  -30.51%
ComplexEgoSlinso-24        2.88µs ± 0%    2.41µs ± 0%  -16.24%
ComplexFtmpl-24            7.66µs ± 0%    7.44µs ± 0%   -2.81%
ComplexGorazor-24          4.52µs ± 0%    4.25µs ± 0%   -6.12%
ComplexHero-24             1.01µs ± 0%    1.00µs ± 0%   -1.19%
ComplexJade-24              674ns ± 0%     744ns ± 0%  +10.32%

name                     old alloc/op   new alloc/op   delta
ComplexEgo-24                568B ± 0%      568B ± 0%    0.00%
ComplexQuicktemplate-24     0.00B          0.00B         0.00%
ComplexEgoSlinso-24          160B ± 0%      160B ± 0%    0.00%
ComplexFtmpl-24            4.91kB ± 0%    4.91kB ± 0%    0.00%
ComplexGorazor-24          2.72kB ± 0%    2.72kB ± 0%    0.00%
ComplexHero-24              0.00B          0.00B         0.00%
ComplexJade-24              0.00B          0.00B         0.00%

name                     old allocs/op  new allocs/op  delta
ComplexEgo-24                31.0 ± 0%      31.0 ± 0%    0.00%
ComplexQuicktemplate-24      0.00           0.00         0.00%
ComplexEgoSlinso-24          2.00 ± 0%      2.00 ± 0%    0.00%
ComplexFtmpl-24              38.0 ± 0%      38.0 ± 0%    0.00%
ComplexGorazor-24            21.0 ± 0%      21.0 ± 0%    0.00%
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
