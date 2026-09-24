# Specialised JSON decoding: measured ceiling for the codegen path

The [comparison audit](2026-09-24-comparison-audit.md) left one question
unmeasured: how much of the 12.74 ms JSON Decoding cell is spent in
**generated-code overhead** rather than in decoding itself. The hand-written
Go decoder used there stopped at plain Go structs; it did not produce the
PureScript values the application actually consumes. This audit closes that
gap: a hand-written decoder for the same schema that consumes the same parsed
`Json` and returns the **same final values** as the generated Argonaut
decoder, measured A/B in the same process.

## What was compared

| | generated | specialised |
|---|---|---|
| input | parser's `Json` (TypeAny over the Go DOM) | same |
| output | `Either JsonDecodeError Payload` | same |
| fields | `RowList` order, same error wrappers | same |
| strings | `runtime.Str` over parser-owned strings | same |
| containers | `Array`, compact/dictionary records | same |

The specialised decoder (`bin/benchmark/json-diagnostic/specialized/decode.go`)
reads the parser's DOM directly (`map[string]any`, `[]any`, `string`,
`float64`, `bool`, `nil`) and calls the runtime constructors. It skips the
generic machinery the generated code has to run for every value: dictionary
dispatch, `FO.lookup` with its unbox/copy/box pass, intermediate `Either`
values at every level, `Rebox` conversions at every boundary, and the
per-record decode plans.

**Faithfulness is a hard gate.** All seventeen corpus fingerprints — the five
timed cases, two valid optional-field cases and ten malformed cases with their
exact `JsonDecodeError` values — are validated against
`test/fixtures/json-decoding/expected.json` in every run, before and after
timing. Any difference aborts the campaign.

## Protocol

- Source workspace: the official `json-dec-cache7-20260923` build; only the
  audit binary is changed (new decoder file, new entry point). `parse`,
  `stringify` and `fingerprint` are the generated ones.
- Both decoders run in the same process, in alternating passes, over the same
  pre-parsed documents, so allocator state, caches and background load are
  shared. Every pass decodes all five timed cases.
- 9 processes, 2 warm-up passes per mode and phase, 5 sampled passes, minimum
  per process, median across processes. `GOMAXPROCS=1`, `GOGC=100` unless
  stated, monotonic timing.
- Corpora: `var/benchmark/json-dec-specialized-20260924/run-gogc100/` and
  `run-gogc-off/`.

## Results

JSON Decoding, 5 timed cases, 636 KB combined:

| phase | GOGC | generated | specialised | ratio |
|---|---|---:|---:|---:|
| decode | 100 | **8.181 ms** | **0.557 ms** | **14.7×** |
| combined | 100 | **12.387 ms** | **3.445 ms** | **3.6×** |
| decode | off | 6.841 ms | 0.568 ms | 12.0× |
| combined | off | 9.301 ms | 3.026 ms | 3.1× |

Allocated bytes per corpus (minimum pass):

| phase | generated | specialised | ratio |
|---|---:|---:|---:|
| decode | 13.49 MB | 1.84 MB | 7.3× |
| combined | 18.10 MB | 6.46 MB | 2.8× |

The generated `combined` cell (12.387 ms) reproduces the published campaign
cell (12.738 ms) and the generated phase split (2.36 parse / 8.76 decode
official) within normal process spread, which cross-checks the A/B runner
against the official protocol.

## Reading

1. **The decode work is almost all generated-code overhead.** Reaching the
   same values through direct field access costs 0.56 ms against 8.18 ms —
   14.7× less, with 7.3× fewer allocated bytes. The gap is not the language,
   the parser or the final representation: it is dictionary dispatch, DOM
   re-boxing and intermediate `Either`/`Maybe` plumbing.
2. **The collector is a consequence, not the cause.** With `GOGC=off` the
   generated decoder recovers 1.34 ms (8.18 → 6.84) and the specialised one
   changes by ~0.01 ms. Cutting the allocation volume removes most of the GC
   share automatically.
3. **The complete path lands near 3.4 ms.** `combined` = parse (2.36 ms,
   unchanged) + specialised decode (0.56 ms) + ~0.5 ms of per-pass effects.
   The remaining door to 1.5–2.5 ms is the parser, not the decoder: at the
   measured parse cost the floor is ≈ 2.9 ms even with a free decoder.
4. **The published ratios now have a bound.** The 20× JSON Decoding ratio
   against the arena C++ reference is not a Go limit: the same final
   PureScript values can be produced in 3.4 ms against 0.65 ms C++ (≈ 5×), and
   that residual is dominated by parsing (2.4 ms Go against 0.4 ms simdjson).

## What this means for the compiler

The experiment bounds a concrete codegen transformation: **specialise the
`DecodeJson` path for the instance actually used at each call site**. The
ingredients the prototype shows to matter:

- read the parser's DOM directly instead of routing through `toObject` and
  `FO.lookup`;
- resolve each field's decoder statically (the closed-dictionary and
  monomorphisation work already computes most of this);
- construct final records, `Maybe`, ADTs and `Either` directly, without
  intermediate wrappers or `Rebox` round trips;
- keep the exact evaluation order and error wrappers, which the oracle gate
  shows is achievable.

What the prototype does **not** establish: that the transformation can be
derived automatically for arbitrary code. It covers one schema, one corpus and
no custom dictionaries. It is a target and a ceiling, not a working pass.

## Limits

- The specialised decoder is hand-written for this schema and reads the DOM
  shape produced by the native parser for an object root; other `Json`
  constructions are out of scope.
- Same-process A/B shares code pages and caches between the two decoders. The
  deliberate trade is comparability; the generated side reproduces the
  official cells within spread.
- The ~0.5 ms difference between `combined` and `parse + decode` was not
  decomposed (per-pass parse warming and allocator effects are candidates).
- Timings were taken on this machine under background load; the deterministic
  allocation totals and the oracle gate are the stable parts.

## Provenance

- `bin/benchmark/json-diagnostic/specialized/` (`decode.go`, `audit_main.go`,
  `run.py`, README).
- Raw runs and aggregates:
  `var/benchmark/json-dec-specialized-20260924/{run-gogc100,run-gogc-off}/results.json`.
- Oracle: `test/fixtures/json-decoding/expected.json`; corpus:
  `test/fixtures/json-decoding/corpus.json`.
