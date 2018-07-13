#!/usr/bin/env bash

set -e
echo "" > coverage.txt

packages=$(go list ./... | grep -v vendor)
go test -race -v $packages

for package in $packages; do
    go test -coverprofile=profile.out -covermode=atomic $package
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
