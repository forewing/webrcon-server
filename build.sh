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
BuildUnix(){
    GOOS=$1 GOARCH=$2 CGO_ENABLED=0 go build -ldflags "$LDFLAGS"
    tar caf output/webrcon-server_$1_$2.tar.gz webrcon-server
    rm webrcon-server
}

BuildWindows(){
    GOOS=windows GOARCH=$1 CGO_ENABLED=0 go build -ldflags "$LDFLAGS"
    zip -r output/webrcon-server_windows_$1.zip webrcon-server.exe
    rm webrcon-server.exe
}

if [ -n "$1" ] && [ $1 = "cross" ]; then
    BuildUnix linux amd64
    BuildUnix linux arm64
    BuildUnix linux arm
    BuildUnix darwin amd64
    BuildUnix darwin arm64
    BuildWindows amd64
else
    CGO_ENABLED=0 go build -ldflags "$LDFLAGS"
    mv webrcon-server output/
fi
