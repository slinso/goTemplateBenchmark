#!/bin/bash -e
# generate github compatible output

# RUNTIME=3s

# # move old results to .old files
# for i in {1..4}; do
#     mv files/results-"${i}".new files/results-"${i}".old
# done

# # run benchmarks
# go test -bench "k(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | tee ./files/results-1.new
# go test -bench "k(Ego|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime=${RUNTIME} | tee ./files/results-2.new
# go test -bench "Complex(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime=${RUNTIME} | tee ./files/results-3.new
# go test -bench "Complex(Ego|EgoSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime=${RUNTIME} | tee ./files/results-4.new

# pretty print for readme.md
i=1
echo ""
echo "## simple benchmarks"
echo "### full featured template engines"
pb <./files/results-"${i}".new | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
echo ""
echo "\`\`\`"
benchcmp -changed files/results-"${i}".old files/results-"${i}".new | tee files/results-"${i}"-benchcmp.txt
echo "\`\`\`"

i=2
echo ""
echo "### precompilation to Go code"
pb <./files/results-"${i}".new | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
echo ""
echo "\`\`\`"
benchcmp -changed files/results-"${i}".old files/results-"${i}".new | tee files/results-"${i}"-benchcmp.txt
echo "\`\`\`"

i=3
echo ""
echo "## more complex test with template inheritance (if possible)"
echo "### full featured template engines"
pb <./files/results-"${i}".new | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
echo ""
echo "\`\`\`"
benchcmp -changed files/results-"${i}".old files/results-"${i}".new | tee files/results-"${i}"-benchcmp.txt
echo "\`\`\`"

i=4
echo ""
echo "### precompilation to Go code"
pb <./files/results-"${i}".new | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
echo ""
echo "\`\`\`"
benchcmp -changed files/results-"${i}".old files/results-"${i}".new | tee files/results-"${i}"-benchcmp.txt
echo "\`\`\`"
