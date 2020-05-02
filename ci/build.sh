#!/bin/sh

apk add --no-cache make

cd $1
make $1-linux
