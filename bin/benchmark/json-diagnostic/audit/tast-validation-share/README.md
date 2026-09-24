# TAST decode layers (audit)

`main.go` times the layers of the Go TAST decode on the frozen corpus:

- `DecodeModuleImpl` with the real `validateSourceUsageModule`,
- the same decoder with a no-op validate (prices the validation pass),
- `Call_Test_JsonTypedAst_decode` (native decode, no record materialisation),
- `Apply(Test.JsonTypedAst.decode, json)`, the entry point the diagnostic
  measures,

and authenticates the measured entry point against the frozen fingerprints.

It must run inside the generated module of a built JsonTypedAst workspace (the
imports are `gopurs/output/purescript` and `gopurs/output/gopurs_runtime`):

```sh
ws=var/benchmark/json-tast-cache7-20260923
mkdir -p "$ws/output/audit"
cp bin/benchmark/json-diagnostic/audit/tast-validation-share/main.go "$ws/output/audit/main.go"
cd "$ws/output"
gzip -dc ../../../test/fixtures/json-typed-ast/tast-corpus.json.gz > /tmp/tast-corpus.json
GOWORK=off GOGC=100 GOMAXPROCS=1 \
  DIAG_CORPUS=/tmp/tast-corpus.json \
  DIAG_EXPECTED=../../../test/fixtures/json-typed-ast/expected.json \
  go run ./audit
```

Result (medians over eight runs, 2026-09-24): validated 15,338 µs, unvalidated
13,200 µs, worker 14,085 µs, boundary 14,320 µs. See
`docs/benchmark-results/2026-09-24-comparison-audit.md`.
