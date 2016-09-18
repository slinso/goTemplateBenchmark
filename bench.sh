#!/bin/bash
# generate github compatible output

echo "simple"
go test -bench "k(Ace|Amber|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=3s | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo ""
echo "simple precompiled"
go test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor)$" -benchmem -benchtime=3s | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo ""
echo "complex"
go test . -bench="Complex" -benchmem -benchtime=3s | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
