#!/bin/sh

rm -rf output
mkdir -p output

export GO111MODULE=on

# Build server
CGO_ENABLED=0 go build -ldflags "-s -w"
mv webrcon-server output/
