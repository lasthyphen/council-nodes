#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Building docker image based off of most recent local commits of avalanchego and dijeth"

AVALANCHE_REMOTE="git@github.com:lasthyphen/dijetsgo.git"
DIJETH_REMOTE="git@github.com:lasthyphen/dijeth.git"
DOCKERHUB_REPO="avaplatform/avalanchego"

DOCKER="${DOCKER:-docker}"
SCRIPT_DIRPATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
ROOT_DIRPATH="$(dirname "${SCRIPT_DIRPATH}")"

AVA_LABS_RELATIVE_PATH="src/github.com/ava-labs"
EXISTING_GOPATH="$GOPATH"

export GOPATH="$SCRIPT_DIRPATH/.build_image_gopath"
WORKPREFIX="$GOPATH/src/github.com/ava-labs"

# Clone the remotes and checkout the desired branch/commits
AVALANCHE_CLONE="$WORKPREFIX/avalanchego"
DIJETH_CLONE="$WORKPREFIX/dijeth"

# Replace the WORKPREFIX directory
rm -rf "$WORKPREFIX"
mkdir -p "$WORKPREFIX"


AVALANCHE_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/avalanchego" rev-parse --short HEAD)"
DIJETH_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/dijeth" rev-parse --short HEAD)"

git config --global credential.helper cache

git clone "$AVALANCHE_REMOTE" "$AVALANCHE_CLONE"
git -C "$AVALANCHE_CLONE" checkout "$AVALANCHE_COMMIT_HASH"

git clone "$DIJETH_REMOTE" "$DIJETH_CLONE"
git -C "$DIJETH_CLONE" checkout "$DIJETH_COMMIT_HASH"

CONCATENATED_HASHES="$AVALANCHE_COMMIT_HASH-$DIJETH_COMMIT_HASH"

"$DOCKER" build -t "$DOCKERHUB_REPO:$CONCATENATED_HASHES" "$WORKPREFIX" -f "$SCRIPT_DIRPATH/local.Dockerfile"
