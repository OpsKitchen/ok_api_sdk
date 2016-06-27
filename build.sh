#!/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo 'install must be run within its container folder' 1>&2
    exit 1
fi

cd ../

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

cd src
gofmt -w .
go build -o example

export GOPATH="$OLDGOPATH"

echo '
build completed
please run ./example

'