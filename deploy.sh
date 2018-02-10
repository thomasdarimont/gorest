#!/bin/bash

: "${ENV:=test}"
: "${CONTAINER:=gorest}"
: "${IMAGE:=gorest}"
: "${REGISTRY:=gitlab:5000}"
: "${TAG:=latest}"
: "${DEPLOY_USER:=vagrant}"
: "${DEPLOY_HOST:=test}"


if [ "$1" = "initssh" ]; then
  if ! grep docker /proc/1/cgroup -qa; then
      echo "Can only be executed within a docker environment"
      exit
  fi

  echo "init ssh"
  which ssh-agent || ( apt-get update -y && apt-get install openssh-client -y )
  eval $(ssh-agent -s)
  ssh-add <(echo "$SSH_PRIVATE_KEY")
  mkdir -p ~/.ssh && chmod 700 ~/.ssh
  echo "$SSH_SERVER_HOSTKEYS" > ~/.ssh/known_hosts
fi

echo "Deploy to $ENV environment"

ssh $DEPLOY_USER@$DEPLOY_HOST "docker stop $CONTAINER || true"
ssh $DEPLOY_USER@$DEPLOY_HOST "docker rm $CONTAINER || true"
ssh $DEPLOY_USER@$DEPLOY_HOST "docker run -d --rm --name $CONTAINER -p 8090:8080 $REGISTRY/$IMAGE:$TAG"

# ENV=test DEPLOY_HOST=test TAG=$CI_COMMIT_SHA ./deploy.sh
