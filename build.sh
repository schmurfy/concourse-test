#!/bin/sh

echo "pwd"
pwd

echo "Files:"
ls -ls ..
ls -l ../project-modules


GOVERSION=`go version`
echo "Go version: $GOVERSION"

echo "Building app..."
go build -o test main.go
