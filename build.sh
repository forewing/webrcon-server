#!/bin/sh

rm -rf output
mkdir -p output

VERSION_PACKAGE="github.com/forewing/webrcon-server/version"
LDFLAGS="-s -w"

if GIT_TAG=$(git describe --tags); then
    LDFLAGS="$LDFLAGS -X '$VERSION_PACKAGE.Version=$GIT_TAG'"
fi

if GIT_HASH=$(git rev-parse HEAD); then
    LDFLAGS="$LDFLAGS -X '$VERSION_PACKAGE.Hash=$GIT_HASH'"
fi

# Build server
CGO_ENABLED=0 go build -ldflags "$LDFLAGS"
mv webrcon-server output/
