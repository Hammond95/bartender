#!/usr/bin/env bash
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPTDIR/../bartender"

function usage() {
    echo "./buildall.sh <RELEASE>"
}

if [ "$#" -ne 1 ]; then
    usage
    exit 1
fi

RELEASE="$1"
[[ -z "$RELEASE" ]] && (echo "Please provide a valid release version." && exit 1)

echo "Using: $(go version)"

PROJECT="github.com/Hammond95/bartender/bartender"
COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')

echo "Starting builds for linux."
linux_archs=(amd64 arm64)
for arch in ${linux_archs[@]}; do
    env GOOS=linux GOARCH=${arch} go build \
        -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        -o ../bin/bartender-${RELEASE}-linux-${arch}
done

echo "Starting builds for darwin."
mac_archs=(amd64 arm64)
for arch in ${mac_archs[@]}; do
    env GOOS=darwin GOARCH=${arch} go build \
        -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        -o ../bin/bartender-${RELEASE}-darwin-${arch}
done

exit 0