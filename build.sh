#!/usr/bin/env bash

echo "fetching tools..."
go get github.com/golang/dep/cmd/dep
go get github.com/ahmetb/govvv

echo "ensure dependencies..."
dep ensure

echo "build binary..."
CGO_ENABLED=0 GOOS=linux govvv build -o out/app

echo "done."