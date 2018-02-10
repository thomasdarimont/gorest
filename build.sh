#!/usr/bin/env bash

echo "fetching tools..."
go get github.com/golang/dep/cmd/dep
go get github.com/ahmetb/govvv

echo "ensuring dependencies..."
dep ensure

echo "building binary..."
export CGO_ENABLED=0
export GOOS=linux

go build -ldflags="$(govvv -flags -pkg $(go list ./actuator))" -o out/app
echo "done."

if [ ! "$1" = "test" ]; then
  exit 0
fi

echo "running tests..."
go test ./...
echo "tests completed."