#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "=== Building Arista ES Backend ==="
npx purs-backend-es build

echo "=== Executing Arista ES Backend (Node.js) ==="
node -e "import('./output-es/App/index.js').then(m => m.main())"
