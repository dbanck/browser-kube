#!/bin/bash

set -e
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" && pwd)"
DOCKER_REGISTRY=https://index.docker.io/v1/
if [[ -z $DOCKER_USERNAME ]]; then
    echo "Must have an enviroment variable with your docker password set named DOCKER_USERNAME"
    exit 1
fi

if [[ -z $DOCKER_PASSWORD ]]; then
    echo "Must have an enviroment variable with your docker password set named DOCKER_PASSWORD"
    exit 1
fi

echo "Creating docker config"

DOCKER_CONFIG="$PROJECT_ROOT/kind-docker-config"
rm -rf $DOCKER_CONFIG
mkdir -p "$DOCKER_CONFIG"
export DOCKER_CONFIG


cat <<EOF >"${DOCKER_CONFIG}/config.json"
{
 "auths": { "${DOCKER_REGISTRY}": {} }
}
EOF
echo -n $DOCKER_PASSWORD  | docker login -u $DOCKER_USERNAME --password-stdin $DOCKER_REGISTRY
