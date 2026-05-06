#!/usr/bin/env bash
set -euo pipefail

GITHUB="https://api.github.com/repos"
REPO="$GITHUB/thalestmm/dots"

OS=$(uname -s)
ARCH=$(uname -m)

echo "$OS $ARCH"

GH_PAYLOAD=$(curl -s "$REPO/releases/latest")

CURRENT_VERSION=$(echo "$GH_PAYLOAD" | jq -r .tag_name)
echo "Current version: $CURRENT_VERSION"

RELEASE_ASSETS=$(echo "$GH_PAYLOAD" | jq -r .assets[].name)
echo "Release assets: $RELEASE_ASSETS"
