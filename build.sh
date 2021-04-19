#!/bin/sh

rm -rf output
mkdir -p output

OUTPUT="webrcon-server"

VERSION_PACKAGE="github.com/forewing/webrcon-server/version"
LDFLAGS="-s -w"

if GIT_TAG=$(git describe --tags); then
    LDFLAGS="$LDFLAGS -X '$VERSION_PACKAGE.Version=$GIT_TAG'"
    OUTPUT="${OUTPUT}-${GIT_TAG}"
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

CMD_BASE="CGO_ENABLED=0 go build -ldflags \"${LDFLAGS}\""

if [ ! -n "$1" ] || [ ! $1 = "all" ]; then
    eval ${CMD_BASE}
    mv webrcon-server output/
    exit 0
fi

# Cross compile

compress_tar_gz(){
    tar caf $1.tar.gz $1
    mv $1.tar.gz output/
    rm ${BIN_FILENAME}
}

compress_zip(){
    zip -q -r $1.zip $1.exe
    mv $1.zip output/
    rm $1.exe
}

PLATFORMS=""
PLATFORMS="$PLATFORMS darwin/amd64 darwin/arm64"
PLATFORMS="$PLATFORMS linux/386 linux/amd64 linux/arm64"
PLATFORMS="$PLATFORMS windows/386 windows/amd64"

PLATFORMS_ARM="linux windows"


for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}"
    BIN_FILENAME_ORINGIN="${BIN_FILENAME}"

    if [ "${GOOS}" = "windows" ];
        then BIN_FILENAME="${BIN_FILENAME}.exe"
    fi

    GO_ENV="GOOS=${GOOS} GOARCH=${GOARCH}"
    CMD="${GO_ENV} ${CMD_BASE} -o ${BIN_FILENAME}"
    echo "${GO_ENV}"
    eval $CMD

    if [ "${GOOS}" = "windows" ]; then
        compress_zip ${BIN_FILENAME_ORINGIN}
    else
        compress_tar_gz ${BIN_FILENAME}
    fi
done

for GOOS in $PLATFORMS_ARM; do
    GOARCH="arm"
    # build for each ARM version
    for GOARM in 7 6 5; do
        BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}${GOARM}"
        BIN_FILENAME_ORINGIN="${BIN_FILENAME}"

        if [ "${GOOS}" = "windows" ];
            then BIN_FILENAME="${BIN_FILENAME}.exe"
        fi

        GO_ENV="GOARM=${GOARM} GOOS=${GOOS} GOARCH=${GOARCH}"
        CMD="${GO_ENV} ${CMD_BASE} -o ${BIN_FILENAME}"
        echo "${GO_ENV}"
        eval "${CMD}"

        if [ "${GOOS}" = "windows" ]; then
            compress_zip ${BIN_FILENAME_ORINGIN}
        else
            compress_tar_gz ${BIN_FILENAME}
        fi
    done
done
