#!/bin/bash -e
# generate github compatible output

RUNTIME=3s

cd ~/gocode/src/github.com/SlinSo/goTemplateBenchmark

echo "### full featured template engines"
go test -bench "k(Ace|Amber|Golang|GolangText|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "### precompilation to Go code"
go test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade|GoDirectBuffer|GoCustomHtmlAPI)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

#echo "### more complex test with template inheritance (if possible)"
#go test . -bench="Complex" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "### more complex test with template inheritance (if possible)"
echo "### full featured template engines"
go test -bench "Complex(Ace|Amber|Golang|GolangText|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "### precompilation to Go code"
go test -bench "Complex(GoStaticString|Ego|Egon|EgoSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade|GoDirectBuffer|GoCustomHtmlAPI)$" -benchmem -benchtime=${RUNTIME} | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
