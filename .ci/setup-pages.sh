#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "$SCRIPT_DIR/versions.env"

apt-get update -q
apt-get install -q -y --no-install-recommends npm imagemagick librsvg2-bin

echo "Installing hugo $HUGO_VERSION"
curl -Lo /tmp/hugo.deb "https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_linux-amd64.deb"
dpkg -i /tmp/hugo.deb

echo "Installing just $JUST_VERSION"
curl -Ls "https://github.com/casey/just/releases/download/$JUST_VERSION/just-$JUST_VERSION-x86_64-unknown-linux-musl.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin just

echo "Installing pandoc $PANDOC_VERSION"
curl -Lo /tmp/pandoc.deb "https://github.com/jgm/pandoc/releases/download/${PANDOC_VERSION}/pandoc-${PANDOC_VERSION}-1-amd64.deb"
dpkg -i /tmp/pandoc.deb

npm install -g markdown-link-check@3.11.2
