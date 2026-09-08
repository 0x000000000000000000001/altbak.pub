#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
POC_ROOT="$PWD"
export GOCACHE="${GOCACHE:-/private/tmp/altbak-solod-gocache}"
export GOMODCACHE="${GOMODCACHE:-/private/tmp/altbak-solod-modcache}"
export UBSAN_OPTIONS=halt_on_error=1
mkdir -p bin
if [[ ! -x bin/so ]]; then
  GOBIN="$POC_ROOT/bin" go install solod.dev/cmd/so@v0.3.0
fi
python3 prepare.py

# These failures are expected evidence, retained alongside the results.
if ./bin/so translate -o full-generated ../../run/bak/go/output/Main >full-translate.log 2>&1; then
  echo 'Full translation unexpectedly succeeded; inspect full-generated.'
else
  cat full-translate.log
fi
if ./bin/so translate -o generated-original ./original/kernels >direct-kernels.log 2>&1; then
  echo 'Unadapted kernel translation unexpectedly succeeded.'
else
  cat direct-kernels.log
fi

(cd original && go build -pgo=off -trimpath -o ../bin/go-original ./cmd)
(cd adapted && go build -pgo=off -trimpath -o ../bin/go-adapted ./cmd)
./bin/so translate -o generated ./adapted/kernels >adapted-translate.log 2>&1

# clang compiles these files as separate translation units. No LTO, no
# fast-math, no source-level algorithm rewrites, preserve signed wrapping.
C_SOURCES=(driver.c generated/kernels.c generated/so/builtin/builtin.c)
C_FLAGS=(-std=gnu11 -fwrapv -DSO_PANIC_MODE=SO_PANIC_ABORT -I generated)
clang -O3 "${C_FLAGS[@]}" "${C_SOURCES[@]}" -o bin/solod-o3
clang -O1 -g -fsanitize=address,undefined -fno-omit-frame-pointer \
  "${C_FLAGS[@]}" "${C_SOURCES[@]}" -o bin/solod-sanitize

{
  go version
  ./bin/so version
  clang --version
  git -C ../.. rev-parse HEAD
} >versions.txt
python3 bench.py "$@"
