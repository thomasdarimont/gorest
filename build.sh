#!/usr/bin/env bash

echo "fetching tools..."
go get github.com/golang/dep/cmd/dep
go get github.com/ahmetb/govvv

echo "ensure dependencies..."
dep ensure

echo "build binary..."
export CGO_ENABLED=0
export GOOS=linux

go build -ldflags="$(govvv -flags -pkg $(go list ./actuator))" -o out/app

echo "done."