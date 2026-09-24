# Specialised JSON decoding: historical implementation reference

**Protocol correction:** the figures below come from the original runner. It
validated setup outputs, retained only the last timed output, and did not
fingerprint the exact timed outputs. They are historical observations, not an
aligned performance bound. The corrected retained-output comparison is in
[the follow-up audit](2026-09-24-record-plan-follow-up.md).

The [comparison audit](2026-09-24-comparison-audit.md) left one question
unmeasured: how much of the 12.74 ms JSON Decoding cell is spent in
**generated-code overhead** rather than in decoding itself. The hand-written
Go decoder used there stopped at plain Go structs; it did not produce the
PureScript values the application actually consumes. This audit investigates
that gap with a hand-written decoder for the same schema that consumes the same parsed
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
generic machinery, including dictionary dispatch, intermediate results, and
per-record plans. `FO.lookup` already has a no-whole-object-copy path through
`UnboxObject`; copying the whole DOM is not an established explanation of the
gap. The prototype also uses compact records and hard-coded constructor tags
and layouts, so this comparison does not isolate dispatch or compiler work.

All seventeen setup corpus fingerprints — the five
timed cases, two valid optional-field cases and ten malformed cases with their
exact `JsonDecodeError` values — are validated against
`test/fixtures/json-decoding/expected.json` in every run before timing. This
establishes agreement on the frozen cases, including their printed errors;
it does not establish general Argonaut semantic equivalence.

## Protocol

- Source workspace: the official `json-dec-cache7-20260923` build; only the
  audit binary is changed (new decoder file, new entry point). `parse`,
  `stringify` and `fingerprint` are the generated ones.
- Both decoders run in the same process, in alternating passes, over the same
  pre-parsed documents, so allocator state, caches and background load are
  shared. Every pass decodes all five timed cases, but only its last output is
  retained. This differs from the official retained-results protocol.
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

Numerical proximity to the separately measured official cells does not
validate the runner's lifetime or output-validation protocol.

## Reading

1. The prototype supplies a concrete implementation reference producing
   usable gopurs values. Its allocation difference motivates removing
   intermediate representations and improving record construction.
2. The separate `GOGC` campaigns show sensitivity to collector settings;
   their difference is not an isolated measurement of GC's time share.
   Default `GOGC` remains 100.
3. This runner did not measure a parse phase. Adding a historical parse
   timing to its decode timing does not decompose `combined` or establish a
   2.9–3.0 ms floor. The earlier 1.5–2.5 ms specialization-only prediction is
   withdrawn; parser improvement or parse/decode fusion would need measurement.
4. Neither this prototype nor the C++ reference establishes a language limit
   or a guaranteed compiler target. Retention, validation, ownership and
   representation costs must be aligned before interpreting the gap.

## What this means for the compiler

One candidate transformation is to **specialise the `DecodeJson` path for
known instances**. The prototype suggests these mechanisms to investigate:

- read the parser's DOM directly instead of routing through `toObject` and
  `FO.lookup`;
- resolve each field's decoder statically (the closed-dictionary and
  monomorphisation work already computes most of this);
- construct final records, `Maybe`, ADTs and `Either` directly, without
  intermediate wrappers or `Rebox` round trips;
- preserve evaluation order and error wrappers, validating beyond the corpus
  with custom dictionaries, representations and simultaneous failures.

What the prototype does **not** establish: that the transformation can be
derived automatically for arbitrary code. It covers one schema, one corpus and
no arbitrary custom dictionaries. No general compiler specialization pass has
been implemented. Library-side propagation and allocation changes remain
independent opportunities, as the follow-up demonstrates.

## Limits

- The specialised decoder is hand-written for this schema and reads the DOM
  shape produced by the native parser for an object root; other `Json`
  constructions are out of scope.
- Same-process A/B shares code pages and caches between the two decoders.
  The old runner started generated-first and kept a fixed phase order.
- The old setup JSON fingerprints hash input text, rather than the encoded
  parsed result; the corrected runner checks actual parsed values too.
- Timings were taken on this machine under background load; the deterministic
  allocation totals and the oracle gate are the stable parts.

## Provenance

- `bin/benchmark/json-diagnostic/specialized/` (`decode.go`, `audit_main.go`,
  `run.py`, README).
- Raw runs and aggregates:
  `var/benchmark/json-dec-specialized-20260924/{run-gogc100,run-gogc-off}/results.json`.
- Oracle: `test/fixtures/json-decoding/expected.json`; corpus:
  `test/fixtures/json-decoding/corpus.json`.
