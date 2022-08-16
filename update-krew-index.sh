#!/bin/bash

export KREW_RELEASE_BOT_VERSION=v0.0.43

curl -LO https://github.com/k8s-devops-tools/krew-release-bot/releases/download/${KREW_RELEASE_BOT_VERSION}/krew-release-bot_${KREW_RELEASE_BOT_VERSION}_linux_amd64.tar.gz
tar -xvf krew-release-bot_${KREW_RELEASE_BOT_VERSION}_linux_amd64.tar.gz
./krew-release-bot action 
