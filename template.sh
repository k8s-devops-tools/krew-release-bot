#! /bin/bash

VERSION=v0.0.43
docker run --rm -v `pwd`:/home/app k8s-devops-tools/krew-release-bot:$VERSION krew-release-bot template 