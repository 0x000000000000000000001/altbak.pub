#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "=== Building Scheme Backend ==="
cp spago-scm.yaml spago.yaml
npx spago build
PATH="$(pwd)/bin:$PATH" npx purescm build
PATH="$(pwd)/bin:$PATH" npx purescm bundle-app --main App

echo "=== Executing Scheme Backend (Chez Scheme) ==="
export DYLD_LIBRARY_PATH="/opt/homebrew/lib:/opt/homebrew/opt/icu4c/lib"
chez --program output/main
