#!/bin/sh

echo "pwd"
pwd

echo "Files:"
ls -ls ..
ls -l ../modules-cache


GOVERSION=`go version`
echo "Go version: $GOVERSION"

echo "Building app..."
go build -o test main.go
