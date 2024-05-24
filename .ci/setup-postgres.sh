#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "$SCRIPT_DIR/versions.env"

echo "Installing just $JUST_VERSION"
curl -Ls "https://github.com/casey/just/releases/download/$JUST_VERSION/just-$JUST_VERSION-x86_64-unknown-linux-musl.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin just
