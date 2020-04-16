#!/bin/sh

echo "Files:"
ls -ls

GOVERSION=`go version`
echo "Go version: $GOVERSION"

echo "Building app..."
go build -o test main.go
