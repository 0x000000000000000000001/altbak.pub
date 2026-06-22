#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "=== Building JavaScript Backend ==="
cp spago-js.yaml spago.yaml
npx spago build

echo "=== Executing JavaScript Backend (Node.js) ==="
node -e "import('./output/App/index.js').then(m => m.main())"
