#!/bin/sh

rm -rf output
mkdir -p output

export GO111MODULE=on

# Generate resources
go generate -x

# Build server
go build -ldflags "-s -w"
mv webrcon-server output/
