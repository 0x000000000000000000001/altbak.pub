# Typed record plans: generic direct decoding for Argonaut

The [specialised-decoder audit](2026-09-24-specialized-decoder-audit.md) showed
that the same final PureScript values can be produced in 0.56 ms instead of
8.18 ms when the decoder reads the parser's DOM directly. That prototype was
hand-written for one schema. This change implements the same idea
**generically**, in the native libraries, with no application-specific code:

- primitive and container decoders (`Int`, `Number`, `String`, `Boolean`,
  `Maybe`, `Array`) carry an immutable schema tag;
- the tag travels with the field method and reaches the record plan;
- a plan whose field is tagged decodes it straight from the DOM instead of
  driving `FO.lookup`, `Either`, `Rebox` and dictionary machinery;
- everything untagged keeps the previous generic path, byte for byte.

No application PureScript changed. The generated Go for existing programs
differs only where the new call shapes appear.

## Mechanism

| Layer | Change |
|---|---|
| runtime | `DecodeKind` tags on decoder values (`WithDecodeKind`, `DecodeKindOf`); never changes behaviour |
| `Class.purs` | primitive/`Maybe`/`Array` instances and `DecodeJsonField` methods attach tags; `fieldStep` carries the field tag onto the step closure the plan receives |
| `Record.go` | plans classify each step: primitives, `Maybe`/`Array` (including custom elements), nested record plans; direct decoding from raw Go DOM values **or** boxed `Value` entries |
| errors | the exact constructors travel in a support record, so direct failures build the same `AtKey`/`Named`/`AtIndex`/`TypeMismatch`/`MissingValue` values |
| contract | typed fields memoize their `reflectSymbol` key; custom fields still evaluate their symbol twice, in the documented order |

The direct path activates only when the input is an object value; any other
input falls back to the original method. `decodeAny` produces errors without
`AtKey`, so a nested record contributes exactly one wrapper, like a generic
step.

## Fidelity

Every measured process must reproduce all seventeen frozen fingerprints,
including the exact `JsonDecodeError` values of the ten malformed cases. In
addition:

- `gopurs-argonaut-codecs` Go tests (with `-race`) and the Node parity test:
  `ok gopurs/record`, including new direct-path tests for raw values, boxed
  values, custom array elements, missing keys and type errors;
- compiler fixtures: `JsonRecordPlan` plus the ten record-related fixtures
  (`CompactRecordConsumers`, `NativeRecordBoxing/Returns/Sizes/Workers`,
  `NestedRecordUpdate`, `NestedRecordUpdateWildcards`, `NewtypeWithRecordUpdate`,
  `RecordTypeChangingUpdate`, `TypeWildcardsRecordExtension`) — 11 passed,
  0 failed; `JsonRecordPlan.go` snapshot updated for the new saturated call
  shapes;
- backend tests: `npm run test:ffi`, `npm run test:runner` — 0 failures;
- diagnostic protocol Runs Go, JS and C and compares all three to the oracle:
  no divergence.

## Results

JSON Decoding, five timed cases, 636 KB, official protocol (3 processes,
2 warm-up passes, 5 samples, minimum per process, median across processes,
`GOMAXPROCS=1`, `GOGC=100`):

| phase | before | typed plans | delta |
|---|---:|---:|---:|
| parse | 2.373 ms | 2.503 ms | noise |
| decode | **9.134 ms** | **5.666 ms** | **−38.0%** |
| combined | **12.738 ms** | **9.595 ms** | **−24.7%** |
| decode allocations | 13.49 MB | **8.99 MB** | **−33.3%** |
| combined allocations | 18.10 MB | **13.60 MB** | **−24.9%** |

The C reference must remain the same work in the same binary; its cells were
unchanged (combined 0.640 ms). The JS backend ignores the tags and keeps its
own timings.

## Reading

The gain is real and generic, but it is far below the hand-written ceiling:

| path | decode |
|---|---:|
| generic (before) | 9.13 ms |
| typed plans (this change) | 5.67 ms |
| hand-written specialised decoder | 0.53 ms |

The remaining factor of ~10 is concentrated where the tags cannot reach yet:

- **custom decoders** (`decodeEvent` in the diagnostic) still call
  `Data.Argonaut.Decode.Decoders.getField`, which knows nothing about tags, so
  their primitive fields (`tag`, `path`, `duration`, `orderId`, `items`) keep
  the generic path; `events` is the largest single field of the workload;
- the plan machinery itself still allocates per record (decoded slice, key
  list, dictionary record) and applies `Right` once;
- nested plans recurse through `decodeValue`, which boxes state per field.

Per-case decode time shows the cost is spread rather than dominated by one
case: `flat-record-arrays` 1.70 ms, `nested-records-and-variants` 1.77 ms,
`null-and-missing-optionals` 1.13 ms, `unicode-escaping-and-numbers` 0.77 ms.

## Next steps

1. Make `getField` / `getFieldOptional'` tag-aware through a native FFI. That
   is where every custom decoder spends its time, and it needs the same error
   support record the plan already uses.
2. Trim the plan's own allocations (build the final record without the
   intermediate slice, or specialize the insertion when labels are unique).
3. Compare against the TAST decoder: the compile path may benefit from the
   same tag discipline if its native reader keeps dictionaries in the hot
   loop.

## Provenance

- Libraries: `gopurs/gopurs-argonaut-codecs`
  (`Class.purs`, `Internal/Record.{purs,go,js}`, `Internal/Typed.{purs,go,js}`,
  `test/record-plan_test.go`, `test/record-plan.mjs`).
- Runtime: `gopurs/gopurs/runtime/runtime.go` (`DecodeKind`), regenerated into
  `src/Gopurs/Runtime.go`.
- Workspace and results: `var/benchmark/json-dec-typed-20260924`,
  `var/benchmark/json-dec-typed-full-20260924/results.json`; baseline
  `var/benchmark/json-dec-cache7-results-20260923/results.json`.
- Reproduction: `python3 bin/benchmark/json-diagnostic.py build --suite
  JsonDecoding --workspace var/benchmark/json-dec-typed-20260924` then
  `measure`.
