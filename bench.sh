#!/bin/bash -e
# generate github compatible output

RUNTIME=3s

cd ~/gocode/src/github.com/SlinSo/goTemplateBenchmark

echo "### full featured template engines"
go test -bench "k(Ace|Amber|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "### precompilation to Go code"
go test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "### more complex test with template inheritance (if possible)"
go test . -bench="Complex" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
