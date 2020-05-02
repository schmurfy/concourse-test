#!/bin/sh

apk add --no-cache make jq protoc

export GOPATH="`pwd`/../project-modules/go"

go install github.com/gogo/protobuf/protoc-gen-gogo

cd $1
go mod download
make $1-linux
