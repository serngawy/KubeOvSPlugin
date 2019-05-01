#!/usr/bin/env bash

echo "GOPATH= " $GOPATH
dep ensure
mkdir -p $GOPATH/bin

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $GOPATH/bin/ovskubeproxy

