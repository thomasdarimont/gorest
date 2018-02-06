#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -o out/app

docker build -t tdlabs.go/app:1.0.0 .