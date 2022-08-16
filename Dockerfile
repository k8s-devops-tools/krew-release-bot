# Container image that runs your code
FROM ghcr.io/k8s-devops-tools/krew-release-bot@sha256:ff6748c6746bb7cddfbc3d8f8b52b5f773f3cce0d60faa9a7a6cfb246eea77ab

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh

# Executes `entrypoint.sh` when the Docker container starts up
ENTRYPOINT ["/entrypoint.sh"]
