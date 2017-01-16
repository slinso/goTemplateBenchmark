#/bin/bash -e

glide cc
glide up

go get -u github.com/tkrajina/ftmpl
go get -u github.com/sipin/gorazor

qtc
ftmpl ftmpl/
gorazor gorazor gorazor


