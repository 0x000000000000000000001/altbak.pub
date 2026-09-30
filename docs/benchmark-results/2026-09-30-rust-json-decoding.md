# Native plans and shared dictionaries for purust JSON decoding

This is the initial 6 817.98 µs milestone. The subsequent
[generated-decoder/native-record cycle](2026-09-30-rust-json-native.md) publishes
**1 372.79 µs** with the preserved optimized Go/C references and the corrected
single-threaded Go measurement environment.

The Rust benchmark table had no JSON Decoding cell. The Go and C cells were
produced by the preserved-build protocol of
`var/benchmark/json-reference-owned-20260925` (six phase-order permutations,
median of six process minima, one binary per backend). This cycle gives purust
the same cell and improves the backend by an order of magnitude on both phases.

## Result and protocol

The preserved `benchmark-go` (text-schema build) and `benchmark-c` binaries from
the `json-reference-owned` cycle and the current purust build were re-run in one
session on the same corpus, with `DIAG_PHASES` cycling through the six phase
permutations. Each cell is the median of the six process minima. All reports
validate against the fixture oracle (`fingerprints` and `json_fingerprints`,
17 modules, error cases included).

| Phase | C reference | Go published | purust before | purust now |
|---|---:|---:|---:|---:|
| parse | 385 | 2 282 | 3 506 | **2 116** |
| decode | 276 | 762 | 40 506 | **4 475** |
| combined | 678 | 2 397 | 43 859 | **6 818** |

The published Go and C cells come from the same campaign protocol (2 241.42 and
657.69); the re-run reproduces them within session noise. In the official
harness run of this cycle the purust binary reaches **6 097 combined** with the
generic Go build at 11 001 and JavaScript at 8 835 in the same sequence
(`bin/benchmark/json-diagnostic.py build` and `measure --runtime rust`).

Rust is **0.55 times** the generic Go build and **0.68 times** JavaScript in
that sequence; under the paired protocol the remaining gap is 5.9 times the
published Go decode and 10.1 times the C combined (decode 16.2 times).

## Implementation

The generic record path rebuilt every record through the `GDecodeJson`
dictionary chain: one class dispatch, one field lookup and one record insert per
field, with an intermediate `Either` per field. Six changes remove most of it:

- **Native construction plans** (`argonaut-codecs` port): a record row builds an
  immutable plan — field name, native kind and the exact generic step of every
  field. The runner decodes `Int`, `Number`, `String`, `Boolean`, `Json`,
  `Maybe`, `Array` and nested records natively, falls back to the step for
  anything else, and lets the step own every error, so fingerprints are
  bit-identical. `getField`, `getFieldOptional` and `getFieldOptional'` follow
  the same pattern; the `Maybe`/`Array`/`Foreign.Object` instances decode their
  well-formed shapes directly and keep the generic composition as the fallback.
- **Sparse generic records** (purust): `Record_a` held one `Option` field per
  class label of the program — 284 fields, about 6.8 KB, zeroed on every
  dictionary construction, which dominated custom decoders (the `Event` decoder
  allocated fifteen of them per event). It now stores the set fields in a small
  `Vec` behind `get_field`/`set_field`/`from_fields`; the codegen reads and
  writes through those accessors.
- **Shared field names**: `RecordFields` keys are `Rc<str>`, so a plan clones
  one key reference per field instead of allocating a string, object copies
  reuse the source keys, and repeated record elements share one key set.
- **Byte-level parser** (`argonaut-core` port): the input is scanned directly
  over the encoded bytes — bulk string copies, stack-buffer numbers, V8-style
  error positions computed on demand — with the whole-input UTF-16 vector gone.
- **Runtime layouts**: `RecordFields` gained `into_entries`,
  `with_capacity`, `push`, `entries_unsorted` and their shared-key variants for
  the native paths.

`SharedRecord` keeps its `Mutex`: threaded mode shares object handles through
`Arc`, so the borrow cannot become a `RefCell`. Codegen and TAST tests stay
green (`80/80` and `41/42`, the remaining case needs unrelated `js-bigints`
dependencies resolved).

## Remaining gap

The allocation counters did not establish a precise cost attribution: their
regions overlap, and the attempted backtrace capture produced no usable
stacks. Closures, boxed results and string copies were hypotheses, not a
measured partition of the remaining cost. The follow-up work therefore uses
generated decoder workers, concrete record carriers and a validated text
cursor, with each stage measured independently.

## Commits

- purust: `3655f80`, `438c54f`, `a6d058f` (runtime and codegen, edge branch).
- purust ports: `e92885f`, `83d1b11`, `61fc4fb`, `9c80c06`, `58b07c9`
  (argonaut-codecs), `0af0f61`, `468e8ad` (argonaut-core), `43b7b56` (arrays),
  `b7b0bee` (aff).
