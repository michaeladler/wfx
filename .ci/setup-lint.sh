#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "$SCRIPT_DIR/versions.env"

echo "Installing just $JUST_VERSION"
curl -Ls "https://github.com/casey/just/releases/download/$JUST_VERSION/just-$JUST_VERSION-x86_64-unknown-linux-musl.tar.gz" |
    tar --extract --gzip --directory=/usr/local/bin just

echo "Installing staticcheck $STATICCHECK_VERSION"
curl -Ls "https://github.com/dominikh/go-tools/releases/download/$STATICCHECK_VERSION/staticcheck_linux_amd64.tar.gz" |
    tar --extract --gzip --strip-components=1 --directory=/usr/local/bin staticcheck/staticcheck
chmod 0755 /usr/local/bin/staticcheck

echo "Installing golangci-lint $GOLANGCI_LINT_VERSION"
curl -Lo /tmp/golangci-lint.deb "https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCI_LINT_VERSION}/golangci-lint-${GOLANGCI_LINT_VERSION}-linux-amd64.deb"
dpkg -i /tmp/golangci-lint.deb
