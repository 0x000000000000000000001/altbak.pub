#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
export GOCACHE="${GOCACHE:-/private/tmp/altbak-solod-gocache}"
export GOMODCACHE="${GOMODCACHE:-/private/tmp/altbak-solod-modcache}"
mkdir -p bin
python3 prepare.py
(cd original && go build -pgo=off -trimpath -o ../bin/go-original ./cmd)
(cd pool && go build -pgo=off -trimpath -o ../bin/go-pool ./cmd)
{
  go version
  git -C ../../../.. rev-parse HEAD
} >versions.txt
python3 bench.py "$@"
