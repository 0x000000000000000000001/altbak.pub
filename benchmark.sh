#!/usr/bin/env bash
set -e

mkdir -p var/benchmark
echo "=== Cleaning ==="
rm -rf output output_tmp .spago var/benchmark/*

echo ""
echo "=== Building JavaScript Backend ==="
cp spago-js.yaml spago.yaml
npx spago build

echo ""
echo "=== Executing JavaScript Backend (Node.js) ==="
{ time node -e "import('./output/App/index.js').then(m => m.main())" ; } 2> var/benchmark/node_time.txt

echo ""
echo "=== Building Scheme Backend ==="
cp spago-scheme.yaml spago.yaml
npx spago build
PATH="$(pwd)/bin:$PATH" npx purescm build
PATH="$(pwd)/bin:$PATH" npx purescm bundle-app --main App

echo ""
echo "=== Executing Scheme Backend (Chez Scheme) ==="
export DYLD_LIBRARY_PATH="/opt/homebrew/lib:/opt/homebrew/opt/icu4c/lib"
{ time chez --program output/main ; } 2> var/benchmark/scheme_time.txt

echo ""
echo "=== Building Erlang Backend ==="
cp spago-erl.yaml spago.yaml
npx spago fetch
# Overwrite legacy CommonJS files with empty ES modules to satisfy newer purs compiler
find .spago/p -name "*.js" -exec sh -c 'echo "export {}" > "$1"' _ {} \;
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

echo ""
echo "=== Executing Erlang Backend (BEAM) ==="
# We call app@ps:main()
{ time erl -noshell -pa output/App -eval '
    app@ps:main(),
    halt().
' ; } 2> var/benchmark/erl_time.txt

echo ""
echo "=========================================================================================================="
echo "                                      BENCHMARK RESULTS                                                   "
echo "=========================================================================================================="
printf "%-35s | %-35s | %-35s\n" "Node.js (V8)" "Chez Scheme (Native)" "Erlang (BEAM)"
printf "%-35s | %-35s | %-35s\n" "-----------------------------------" "-----------------------------------" "-----------------------------------"
grep -E '^(real|user|sys)' var/benchmark/node_time.txt > var/benchmark/node_clean.txt
grep -E '^(real|user|sys)' var/benchmark/scheme_time.txt > var/benchmark/scheme_clean.txt
grep -E '^(real|user|sys)' var/benchmark/erl_time.txt > var/benchmark/erl_clean.txt

exec 3< var/benchmark/node_clean.txt
exec 4< var/benchmark/scheme_clean.txt
exec 5< var/benchmark/erl_clean.txt
while read -r n1 <&3 && read -r s1 <&4 && read -r e1 <&5; do
    printf "%-35s | %-35s | %-35s\n" "$n1" "$s1" "$e1"
done
exec 3<&-
exec 4<&-
exec 5<&-
echo "=========================================================================================================="
