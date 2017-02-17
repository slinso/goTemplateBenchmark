#!/bin/bash -e

RUNTIME=3s

go test -bench . -benchmem -benchtime=${RUNTIME} > new.run
benchcmp old.run new.run
mv new.run old.run
