#!/bin/bash -e
shopt -s expand_aliases

#dep ensure -update
alias gg='go get -u -v'

gg github.com/tkrajina/ftmpl
go install github.com/tkrajina/ftmpl
gg github.com/sipin/gorazor
go install github.com/sipin/gorazor
gg github.com/valyala/quicktemplate
go install github.com/valyala/quicktemplate
gg github.com/benbjohnson/ego
go install github.com/benbjohnson/ego
gg github.com/shiyanhui/hero
go install github.com/shiyanhui/hero/hero
gg github.com/Joker/jade
go install github.com/Joker/jade/cmd/jade

qtc -dir quicktemplate
ftmpl ftmpl/
gorazor gorazor gorazor
hero -source hero/
jade -d jade/ jade/simple.jade
jade -d jade/ jade/index.jade
