#!/bin/bash

version=$1
if [ "$version" == "" ]; then
	echo "version not provided"
	exit 1
fi

echo $version

## push for github actions
# id=$(echo $(docker build -f Dockerfile --platform linux/amd64 --quiet=true . 2>/dev/null) | awk '{print $NF}')

prefix=k8s-devops-tools
name=krew-release-bot
name=$(echo "$name" | sed -r 's#/$##g')
registry=ghcr.io
image_name=$(echo "$prefix/$name:$version" | sed -r 's#/+#/#g')
remote_name=$(echo "$registry/$prefix/$name:$version" | sed -r 's#/+#/#g')


docker build -f Dockerfile --platform linux/amd64 --tag "$image_name" --tag "$remote_name" . 

echo $id
echo $name
echo $remote_name

# docker tag $id $dir
# docker tag $id "$remote_name"
docker push "$remote_name"

# docker build . -t k8s-devops-tools/krew-release-bot:$version -f Dockerfile --platform linux/amd64
# docker tag k8s-devops-tools/krew-release-bot:$version ghcr.io/k8s-devops-tools/krew-release-bot:$version
# docker push ghcr.io/k8s-devops-tools/krew-release-bot:$version
