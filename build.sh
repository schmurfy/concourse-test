#!/bin/sh

BASEPATH=`pwd`

# echo "Environment:"
# export

# echo "Current Path:"
# pwd

# echo "Files:"
# ls -ls ..
find $BASEPATH/project-modules -maxdepth 4

export GOPATH=$BASEPATH/project-modules/go

GOVERSION=`go version`
echo "Go version: $GOVERSION"

echo "Building app..."
go build -o test main.go
