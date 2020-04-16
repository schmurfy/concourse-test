#!/bin/sh
GOVERSION=`go version`
echo "Go version: $GOVERSION"

echo "Building app..."
go build -o test main.go
