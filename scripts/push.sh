#!/usr/bin/env bash
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

version="$1"
if [[ -z "$version" ]]; then
    echo "Please provide a release version when calling the script!"
    exit 1
fi
image="gintonic"
destimage="gintonic"

echo "Login for DockerHub."
read -p "Username: " DOCKER_USERNAME
read -sp "Password: " DOCKER_PASSWORD

# Image is already built by Makefile call

docker login \
    -u $DOCKER_USERNAME \
    -p $DOCKER_PASSWORD

docker tag $image:$version mdelucadev/$image:$version
docker push mdelucadev/$image:$version

