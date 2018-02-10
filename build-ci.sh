#!/usr/bin/env bash

if ! grep docker /proc/1/cgroup -qa; then
    echo "Can only be executed within a docker environment"
    exit
fi

echo "preparing go env"
mkdir -p $GOPATH/src/github.com/thomasdarimont/gotraining
ln -sf $CI_PROJECT_DIR $GOPATH/src/github.com/thomasdarimont/gotraining
cd $GOPATH/src/github.com/thomasdarimont/gotraining/$CI_PROJECT_NAME
echo "bulding $CI_PROJECT_NAME..."
echo $(cat VERSION).$CI_PIPELINE_ID > VERSION

./build.sh test