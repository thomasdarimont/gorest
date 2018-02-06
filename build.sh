#!/usr/bin/env bash

go get -v github.com/ahmetb/govvv
CGO_ENABLED=0 GOOS=linux govvv build -o out/app

docker build -t gitlab:5000/gorest:latest .