#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "=== Building Erlang Backend ==="
cp spago-erl.yaml spago.yaml
npx spago fetch
# Overwrite legacy CommonJS files with empty ES modules to satisfy newer purs compiler
find .spago/p -name "*.js" -exec sh -c 'echo "export const __patched = 1;" > "$1"' _ {} \;
# Inject missing Erlang FFI files for PureScript 0.15 compatibility in purerl
find .spago/p/prelude -name "Reflectable.purs" | while read f; do
  printf -- "-module(data_reflectable@foreign).\n-export([unsafeCoerce/1]).\nunsafeCoerce(X) -> X.\n" > "${f%.purs}.erl"
done
find .spago/p/prelude -name "Symbol.purs" | while read f; do
  printf -- "-module(data_symbol@foreign).\n-export([unsafeCoerce/1]).\nunsafeCoerce(X) -> X.\n" > "${f%.purs}.erl"
done
PATH="$(pwd)/bin:$PATH" npx spago build
# erlc compiles the .erl files to .beam files
erlc -o output/App output/*/*.erl

echo "=== Executing Erlang Backend (BEAM) ==="
erl -noshell -pa output/App -eval '
    app@ps:main(),
    halt().
'
