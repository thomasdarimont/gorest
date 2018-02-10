#!/usr/bin/env bash

: "${ENV:=test}"
: "${CONTAINER:=gorest}"
: "${IMAGE:=gorest}"
: "${REGISTRY:=gitlab:5000}"
: "${TAG:=latest}"
: "${DEPLOY_USER:=vagrant}"
: "${DEPLOY_HOST:=test}"

echo "Deploy to $ENV environment"

ssh $DEPLOY_USER@$DEPLOY_HOST "docker stop $CONTAINER || true"
ssh $DEPLOY_USER@$DEPLOY_HOST "docker rm $CONTAINER || true"
ssh $DEPLOY_USER@$DEPLOY_HOST "docker run -d --rm --name $CONTAINER -p 8090:8080 $REGISTRY/$IMAGE:$TAG"

# ENV=test DEPLOY_HOST=test TAG=$CI_COMMIT_SHA ./deploy.sh
