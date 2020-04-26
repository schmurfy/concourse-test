#!/bin/bash

PATHS=`find . -name go.mod`
for PATH in $PATHS; do
  echo $PATH
  DIR="$(/usr/bin/dirname $PATH)"
  echo $DIR
  pushd $DIR
  go mod download
  popd
done
