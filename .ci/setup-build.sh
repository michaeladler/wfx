#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "$SCRIPT_DIR/versions.env"

apt-get update -q
apt-get install -q -y --no-install-recommends xz-utils zstd

echo "Installing just $JUST_VERSION"
curl -Ls "https://github.com/casey/just/releases/download/$JUST_VERSION/just-$JUST_VERSION-x86_64-unknown-linux-musl.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin just

echo "Installing zig $ZIG_VERSION"
curl -L "https://ziglang.org/download/${ZIG_VERSION}/zig-linux-x86_64-${ZIG_VERSION}.tar.xz" |
    tar --extract --xz --strip-components=1 --directory=/usr/local/bin

echo "Installing goreleaser $GORELEASER_VERSION"
curl -L "https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser-${GORELEASER_VERSION}-1-x86_64.pkg.tar.zst" |
    tar --extract --zstd --directory=/
