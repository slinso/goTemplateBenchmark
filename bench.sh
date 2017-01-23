#!/bin/bash
# generate github compatible output

RUNTIME=3s
NEXT=go1.8rc1

cd ~/gocode/src/github.com/SlinSo/goTemplateBenchmark

echo "simple:"
go test -bench "k(Ace|Amber|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | tee /tmp/old1.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "simple precompiled:"
go test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero)$" -benchmem -benchtime=${RUNTIME} | tee /tmp/old2.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "complex:"
go test . -bench="Complex" -benchmem -benchtime=${RUNTIME} | tee /tmp/old3.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'


########################################
############### NEXT ###################
########################################
echo "simple:"
${NEXT} test -bench "k(Ace|Amber|Golang|Handlebars|Kasia|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | tee /tmp/new1.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "simple precompiled:"
${NEXT} test -bench "k(Ego|Egon|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero)$" -benchmem -benchtime=${RUNTIME} | tee /tmp/new2.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

echo "complex:"
${NEXT} test . -bench="Complex" -benchmem -benchtime=${RUNTIME} | tee /tmp/new3.txt | pb | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'

benchcmp /tmp/old1.txt /tmp/new1.txt
benchcmp /tmp/old2.txt /tmp/new2.txt
benchcmp /tmp/old3.txt /tmp/new3.txt

