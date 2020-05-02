#!/bin/sh

apk add --no-cache make

cd $1
make
