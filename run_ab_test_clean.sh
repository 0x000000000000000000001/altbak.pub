#!/bin/bash
set -e

PS_DIR="/Users/0x1/Documents/htdocs/purescript"
OPT_DIR="/Users/0x1/Documents/htdocs/purescript-backend-optimizer"
GOPURS_DIR="/Users/0x1/Documents/htdocs/gopurs/gopurs"
ALTBAK_DIR="/Users/0x1/Documents/htdocs/altbak.pub"

echo "=== Switching to main branches ==="
cd "$PS_DIR" && git checkout master
cd "$OPT_DIR" && git checkout main
cd "$GOPURS_DIR" && git checkout main

echo "=== Cleaning caches (main) ==="
cd "$GOPURS_DIR" && rm -rf output .spago
cd "$ALTBAK_DIR" && rm -rf output .spago

echo "=== Building purescript (main) ==="
cd "$PS_DIR"
stack build
stack install

echo "=== Running altbak.pub (main) ==="
cd "$ALTBAK_DIR"
./bin/go/run -c > /tmp/old_go_run_output.txt 2>&1 || true
rm -rf /tmp/gopurs_old
cp -r output /tmp/gopurs_old

echo "=== Switching back to edge branches ==="
cd "$PS_DIR" && git checkout edge
cd "$OPT_DIR" && git checkout edge
cd "$GOPURS_DIR" && git checkout edge

echo "=== Cleaning caches (edge) ==="
cd "$GOPURS_DIR" && rm -rf output .spago
cd "$ALTBAK_DIR" && rm -rf output .spago

echo "=== Building purescript (edge) ==="
cd "$PS_DIR"
stack build
stack install

echo "=== Running altbak.pub (edge) ==="
cd "$ALTBAK_DIR"
./bin/go/run -c > /tmp/new_go_run_output.txt 2>&1 || true
rm -rf /tmp/gopurs_new
cp -r output /tmp/gopurs_new

echo "=== Generating Diff ==="
diff -r -u /tmp/gopurs_old /tmp/gopurs_new > /tmp/gopurs.diff || true

echo "=== Done! ==="
