#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
export GOCACHE="${GOCACHE:-/private/tmp/altbak-solod-gocache}"
export GOMODCACHE="${GOMODCACHE:-/private/tmp/altbak-solod-modcache}"
export UBSAN_OPTIONS=halt_on_error=1
export ASAN_OPTIONS=detect_stack_use_after_return=1
if [[ ! -x ../bin/so ]]; then
  mkdir -p ../bin
  GOBIN="$(cd ../bin && pwd)" go install solod.dev/cmd/so@v0.3.0
fi
mkdir -p bin
python3 prepare.py
for variant in original syntax; do
  if ../bin/so translate -o "generated-$variant" "./$variant/kernels" >"$variant-translate.log" 2>&1; then
    echo "$variant translation succeeded; inspect generated C before using it."
  else
    cat "$variant-translate.log"
  fi
done
(cd original && go build -pgo=off -trimpath -o ../bin/go-original ./cmd)
(cd syntax && go build -pgo=off -trimpath -o ../bin/go-syntax ./cmd)
(cd arena && go build -pgo=off -trimpath -o ../bin/go-arena ./cmd)
../bin/so translate -o generated ./arena/kernels >arena-translate.log 2>&1

# Separate C translation units, no LTO, signed wrapping retained.
C_SOURCES=(driver.c generated/kernels.c generated/so/builtin/builtin.c)
C_FLAGS=(-std=gnu11 -fwrapv -DSO_PANIC_MODE=SO_PANIC_ABORT -I generated)
clang -O3 "${C_FLAGS[@]}" "${C_SOURCES[@]}" -o bin/solod-arena 2>clang-build.log
clang -O1 -g -fsanitize=address,undefined -fno-omit-frame-pointer \
  "${C_FLAGS[@]}" "${C_SOURCES[@]}" -o bin/solod-sanitize 2>clang-sanitize.log
{
  go version
  ../bin/so version
  clang --version
  git -C ../../.. rev-parse HEAD
} >versions.txt
python3 bench.py "$@"
