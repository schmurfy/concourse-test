#!/bin/sh

export GOPATH="`pwd`/../project-modules/go"
export PATH="$GOPATH/bin:$PATH"

cd $1
go mod download
go install github.com/gogo/protobuf/protoc-gen-gogo
make $1-linux
