#!/bin/bash
#Expect to run this from the root of the project as "./build/buildImage.sh"

GIT_COMMIT="$(git rev-parse HEAD)"
echo "GIT_COMMIT hash is $GIT_COMMIT"

DOCKER_TAG="$(date +%s)"
echo "DOCKER_TAG: $DOCKER_TAG"

#Write version file
VERSION_FILE="./version.txt"
echo "GIT_COMMIT: $GIT_COMMIT" > $VERSION_FILE

#Set config file reference
CONFIG_FILE="deploy/portfolio-service-conf.json"

docker build --build-arg "VERSION_FILE=$VERSION_FILE" --build-arg "CONFIG_FILE=$CONFIG_FILE" --tag "portfolio-service:$DOCKER_TAG" -f build/Dockerfile .