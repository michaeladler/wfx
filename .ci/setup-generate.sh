#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "$SCRIPT_DIR/versions.env"

apt-get update -q
apt-get install -q -y --no-install-recommends python3-yaml git-lfs apt-transport-https unzip

echo "Installing gofumpt $GOFUMPT_VERSION"
curl -Lso /usr/local/bin/gofumpt "https://github.com/mvdan/gofumpt/releases/download/v$GOFUMPT_VERSION/gofumpt_v${GOFUMPT_VERSION}_linux_amd64"
chmod 0755 /usr/local/bin/gofumpt

echo "Installing just $JUST_VERSION"
curl -Ls "https://github.com/casey/just/releases/download/$JUST_VERSION/just-$JUST_VERSION-x86_64-unknown-linux-musl.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin just

echo "Installing go-swagger $SWAGGER_VERSION"
curl -Lo /usr/local/bin/swagger "https://github.com/go-swagger/go-swagger/releases/download/v${SWAGGER_VERSION}/swagger_linux_amd64"
chmod 0755 /usr/local/bin/swagger

echo "Installing mockery $MOCKERY_VERSION"
curl -Ls "https://github.com/vektra/mockery/releases/download/v$MOCKERY_VERSION/mockery_${MOCKERY_VERSION}_Linux_x86_64.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin mockery

echo "Installing flatbuffers $FLATBUFFERS_VERSION"
curl -Lo /tmp/flatc.zip "https://github.com/google/flatbuffers/releases/download/v${FLATBUFFERS_VERSION}/Linux.flatc.binary.g++-10.zip"
unzip /tmp/flatc.zip flatc -d /usr/local/bin/
