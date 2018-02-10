#!/usr/bin/env bash

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