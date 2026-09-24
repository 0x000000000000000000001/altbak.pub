# Specialised-decoder A/B audit

Compares a schema-specific implementation reference with generated Argonaut
decoding. Neither implementation is a bound on automatic compiler performance.

- `decode.go` — hand-written decoder for the `Test.JsonDecoding` schema. It
  consumes the same parsed `Json` the generated decoder receives (the native
  parser's `any` DOM) and produces the same final PureScript values: compact
  records, `Maybe`, the `Event` ADT, `Either JsonDecodeError Payload`.
- `audit_main.go` — A/B runner. Both decoders run in the same process, in
  alternating passes, over the same parsed documents; every timed result is
  retained and fingerprint-checked against the frozen oracle outside the timed
  interval. The retention buffer is allocated before timing, like the official
  runner. Parse outputs are checked by encoding the actual parsed values.
- `audit_main_test.go` — regression that corrupts an early timed output while
  leaving the last output valid; validation must reject that exact early result.
- `main_audit.go` — replacement entry point for the audit workspace.
- `run.py` — copies a selected official workspace, grafts the runner, tests and
  builds `benchmark-audit`, and records source/binary hashes. Campaigns sanitize
  diagnostic environment variables, alternate starting decoder and workspace
  order, rotate/reverse phase order, and record machine/load observations.
- `coverage.py` / `coverage_main.go` — instrument a separate copy of generated
  record FFI, count actual direct/custom paths, and validate retained outputs.
  These binaries produce no benchmark timings.

The audit fails on any fingerprint difference, including printed errors for
the ten malformed corpus cases. Corpus agreement does not establish semantic
equivalence for arbitrary schemas, custom dictionaries or numeric policies.

Historical campaigns before the retained-output fix checked setup only and
kept a single timed-result sink. Their 0.557 ms figure must not be presented
as an aligned floor.

```sh
python3 bin/benchmark/json-diagnostic/specialized/run.py setup \
    --source YOUR_OFFICIAL_WORKSPACE \
    --workspace NEW_AUDIT_WORKSPACE
python3 bin/benchmark/json-diagnostic/specialized/run.py campaign \
    --workspace NEW_AUDIT_WORKSPACE \
    --compare-workspace PRESERVED_BASELINE_AUDIT_WORKSPACE \
    --processes 6 --output NEW_RESULTS_DIRECTORY
```

Current results: `docs/benchmark-results/2026-09-24-record-plan-follow-up.md`.
Historical results: `docs/benchmark-results/2026-09-24-specialized-decoder-audit.md`.
