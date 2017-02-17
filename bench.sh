#!/bin/bash -e
# generate github compatible output

RUNTIME=1s

cd ~/gocode/src/github.com/SlinSo/goTemplateBenchmark

echo "simple:"
go test -bench "k(Ace|Amber|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "simple precompiled:"
go test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "complex:"
go test . -bench="Complex" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
