#!/usr/bin/env /bin/bash

: "${IMAGE:=gorest}"
: "${REGISTRY:=gitlab:5000}"
: "${TAG:=latest}"

echo "docker: building: $REGISTRY/$IMAGE:$TAG"
docker build -t $REGISTRY/$IMAGE:$TAG .
echo "docker: build completed"

if [ ! "$1" = "push" ]; then
  exit 0
fi

echo "docker: pushing image: $REGISTRY/$IMAGE:$TAG"
docker push $REGISTRY/$IMAGE:$TAG
echo "docker: pushing image completed"

# TAG=$CI_COMMIT_SHA ./build-docker.sh