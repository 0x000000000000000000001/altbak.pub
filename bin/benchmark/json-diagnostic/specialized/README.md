# Specialised-decoder A/B audit

Measures the ceiling of a schema-specialised decode path for the JSON Decoding
diagnostic, against the Argonaut decoder gopurs generates today.

- `decode.go` — hand-written decoder for the `Test.JsonDecoding` schema. It
  consumes the same parsed `Json` the generated decoder receives (the native
  parser's `any` DOM) and produces the same final PureScript values: compact
  records, `Maybe`, the `Event` ADT, `Either JsonDecodeError Payload`.
- `audit_main.go` — A/B runner. Both decoders run in the same process, in
  alternating passes, over the same parsed documents; every timed result is
  fingerprint-checked against the frozen oracle outside the timed interval.
- `main_audit.go` — replacement entry point for the audit workspace.
- `run.py` — copies the official `json-dec-cache7-20260923` workspace, grafts
  the runner, builds `benchmark-audit`, and runs the campaign.

The decoder must stay faithful, not merely fast: the audit fails on any
fingerprint difference, including the exact `JsonDecodeError` values of the ten
malformed corpus cases. That check passed on every measured run.

```sh
python3 bin/benchmark/json-diagnostic/specialized/run.py setup \
    --workspace var/benchmark/json-dec-specialized-20260924/work
python3 bin/benchmark/json-diagnostic/specialized/run.py campaign \
    --workspace var/benchmark/json-dec-specialized-20260924/work \
    --output var/benchmark/json-dec-specialized-20260924/run-gogc100
```

Results and interpretation: `docs/benchmark-results/2026-09-24-specialized-decoder-audit.md`.
