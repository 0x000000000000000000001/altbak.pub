# Typed record plans: generic direct decoding for Argonaut

This report records the first typed-plan campaigns. The [follow-up
audit](2026-09-24-record-plan-follow-up.md) corrects the callback contract and
JS validation, measures actual nested-plan coverage, and provides a preserved
paired comparison. The historical specialised runner's 0.557 ms result used
a different retention/validation protocol and is not an aligned bound.

The implementation adds generic native-library mechanisms:

- primitive and container decoders (`Int`, `Number`, `String`, `Boolean`,
  `Maybe`, `Array`) carry an immutable schema tag;
- the tag travels with the field method and reaches the record plan;
- a plan whose field is tagged decodes it straight from the DOM instead of
  driving `FO.lookup`, `Either`, `Rebox` and dictionary machinery;
- untagged decoders use the ordinary callback path.

No application PureScript changed. The generated Go for existing programs
differs only where the new call shapes appear.

## Mechanism

| Layer | Change |
|---|---|
| runtime | `FunctionData` metadata carrier (unchanged); tags are plain package-local types |
| `Class.purs` | primitive/`Maybe`/`Array` instances and `DecodeJsonField` methods attach tags; `fieldStep` carries the field tag onto the step closure the plan receives |
| `Record.go` | tags (`typedKind`), plans classify primitive and `Maybe`/`Array` steps (including custom elements); a record-plan kind existed but the `decodeRecord` wrapper did not propagate its metadata |
| `Decoders.purs` | `getField`, `getFieldOptional` and `getFieldOptional'` route through `decodeFieldFast`, so every custom decoder using the standard accessors benefits from the same tags |
| errors | error constructors travel in a support record, captured once per decoder or plan; corpus agreement covers the exercised errors |
| contract (subsequently corrected) | typed fields memoized their symbol callback without proving that callback constant; the follow-up restores reevaluation for every field |

The tag constructors live in `Internal/Record.{purs,go,js}`. Consolidation
resolved the observed native bootstrap `undefined: TagDecoder` failure; the
earlier claim about its precise loader cause was not independently established.

The direct path activates only when the input is an object value; any other
input falls back to the original method. `decodeAny` produces errors without
`AtKey`, so a nested record contributes exactly one wrapper, like a generic
step.

## Fidelity

Every measured process must reproduce all seventeen frozen fingerprints,
including the exact `JsonDecodeError` values of the ten malformed cases. In
addition:

- `gopurs-argonaut-codecs` Go tests (with `-race`) passed for raw/boxed values,
  custom array elements, missing keys and type errors. The Node helper test
  used outdated argument counts, so its success did not validate the new FFI;
- compiler fixtures: `JsonRecordPlan` plus the ten record-related fixtures
  (`CompactRecordConsumers`, `NativeRecordBoxing/Returns/Sizes/Workers`,
  `NestedRecordUpdate`, `NestedRecordUpdateWildcards`, `NewtypeWithRecordUpdate`,
  `RecordTypeChangingUpdate`, `TypeWildcardsRecordExtension`) — 11 passed,
  0 failed; `JsonRecordPlan.go` snapshot updated for the new saturated call
  shapes;
- backend tests: `npm run test:ffi`, `npm run test:runner` — 0 failures;
- the official diagnostic found no oracle divergence, but its JS build used
  registry codecs rather than the modified local codecs. The follow-up uses
  local codecs and fixes the JS arities before comparing implementations.

## Results

JSON Decoding, five timed cases, 636 KB, official protocol (3 processes,
2 warm-up passes, 5 samples, minimum per process, median across processes,
`GOMAXPROCS=1`, `GOGC=100`):

| phase | before | typed plans | + typed accessors |
|---|---:|---:|---:|
| parse | 2.373 ms | 2.503 ms | 2.554 ms |
| decode | **9.134 ms** | 5.666 ms | **5.365 ms** |
| combined | **12.738 ms** | 9.595 ms | **8.670 ms** |
| decode allocations | 13.49 MB | 8.99 MB | **8.86 MB** |
| combined allocations | 18.10 MB | 13.60 MB | **13.47 MB** |

Cumulative deltas: decode **−41.3%**, combined **−31.9%**, decode allocations
**−34.3%**. The two increments are separate campaigns on the same machine, so
read the direction and the deterministic allocation totals rather than third
decimals. The C reference keeps the same work in the same binary; its cells
were unchanged (combined 0.640 ms). The JS backend ignores the tags and keeps
its own timings.

## Reading

Historical observations (different campaigns and runner protocols):

| path | decode |
|---|---:|
| generic (before) | 9.13 ms |
| typed plans and accessors (this change) | 5.37 ms |
| hand-written specialised decoder, old single-sink runner | 0.557 ms |

The code still contained these mechanisms; their individual time shares were
not isolated by these campaigns:

- the custom decoder body itself (`decodeEvent`) is a compiled PureScript
  function: its branches, `Either` plumbing and constructor calls stay on the
  generic path. Standard field accessors try tagged decoding after `FO.lookup`
  and retain surrounding `Either` handling;
- the plan machinery still allocates per record (decoded slice, key list,
  dictionary record) and applies `Right` once;
- `decodeCustom` wraps every custom element with `isRight`/`leftOf`/
  `rightValue` calls;
- record-wrapper metadata was missing, and native arrays could be boxed into
  an intermediate slice before allocating the decoded output slice.

The reported 5.67 → 5.37 ms decode and 9.60 → 8.67 ms combined changes came
from separate campaigns whose intermediate workspace/results were overwritten.
They cannot establish an accessor-only gain or show that compiled custom
decoder bodies dominate the remainder.

Per-case decode time shows the cost is spread rather than dominated by one
case: `flat-record-arrays` 1.70 ms, `nested-records-and-variants` 1.77 ms,
`null-and-missing-optionals` 1.13 ms, `unicode-escaping-and-numbers` 0.77 ms.

## Next steps

1. Trim the plan's own allocations (build the final record without the
   intermediate slice, or specialize the insertion when labels are unique).
2. Measure direct/fallback coverage, propagate record-wrapper tags, and avoid
   intermediate array boxing before assigning the remaining gap to compilation.
3. Validate the compiler's TAST path and consider custom-decoder specialization
   based on the remaining measured costs. See the follow-up for completed work.

## Provenance

- Libraries: `gopurs/gopurs-argonaut-codecs`
  (`Class.purs`, `Decoders.purs`, `Internal/Record.{purs,go,js}`,
  `test/record-plan_test.go`, `test/record-plan.mjs`).
- Runtime: `gopurs/gopurs/runtime/runtime.go` (no tag-specific code; the tag
  types live in `Record.go`).
- Workspace and results: `var/benchmark/json-dec-typed-20260924`,
  `var/benchmark/json-dec-typed-full-20260924/results.json`; baseline
  `var/benchmark/json-dec-cache7-results-20260923/results.json`.
- Reproduction: `python3 bin/benchmark/json-diagnostic.py build --suite
  JsonDecoding --workspace var/benchmark/json-dec-typed-20260924` then
  `measure`.
