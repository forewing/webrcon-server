#!/bin/sh

rm -rf output
mkdir -p output

(cd / && go get -u github.com/go-bindata/go-bindata/go-bindata)

export GO111MODULE=on

# Build resources
go-bindata -fs -prefix resources/ resources/

# Build server
go build -ldflags "-s -w"
mv webrcon-server output/
