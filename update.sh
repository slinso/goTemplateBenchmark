#!/bin/bash -e
shopt -s expand_aliases

#dep ensure -update
alias gg='go get -u -v'
cd ~
gg github.com/tkrajina/ftmpl
gg github.com/sipin/gorazor
gg github.com/valyala/quicktemplate/qtc
gg github.com/benbjohnson/ego/...
gg github.com/shiyanhui/hero
gg github.com/Joker/jade/cmd/jade

qtc -dir quicktemplate
ftmpl ftmpl/
gorazor -prefix github.com/SlinSo/goTemplateBenchmark gorazor gorazor
hero -source hero/
jade -d jade/ jade/simple.jade
jade -d jade/ jade/index.jade
