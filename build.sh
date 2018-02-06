#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -o out/app

docker build -t gitlab:5000/gorest:latest .