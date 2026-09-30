# Native plans and sparse dictionaries for purust JSON decoding

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
| parse | 375 | 1 923 | 3 506 | **2 569** |
| decode | 279 | 666 | 40 506 | **4 433** |
| combined | 656 | 2 162 | 43 859 | **7 395** |

The published Go and C cells come from the same campaign protocol (2 241.42 and
657.69); the re-run reproduces them within session noise. In the official
harness run of this cycle the purust binary reaches **6 823 combined** with the
generic Go build at 11 416 and JavaScript at 8 738 in the same sequence
(`bin/benchmark/json-diagnostic.py build` and `measure --runtime rust`).

Decode is now **0.55 times** the JavaScript build measured in the same harness
sequence (4 113 against 7 427) and combined **0.78 times** (6 823 against
8 738); under the paired protocol the remaining gap is 6.7 times the published
Go build (4 433 against 666) and 15.9 times the C reference.

## Implementation

The generic record path rebuilt every record through the `GDecodeJson`
dictionary chain: one class dispatch, one field lookup and one record insert per
field, with an intermediate `Either` per field. Four changes remove most of it:

- **Native construction plans** (`argonaut-codecs` port): a record row builds an
  immutable plan — field name, native kind and the exact generic step of every
  field. The runner decodes `Int`, `Number`, `String`, `Boolean`, `Json`,
  `Maybe`, `Array` and nested records natively, falls back to the step for
  anything else, and lets the step own every error, so fingerprints are
  bit-identical. `getField`, `getFieldOptional` and `getFieldOptional'` follow
  the same pattern. The `Maybe`/`Array`/`Foreign.Object` instances decode their
  well-formed shapes directly and keep the generic composition as the fallback.
- **Sparse generic records** (purust): `Record_a` held one `Option` field per
  class label of the program — 284 fields, about 6.8 KB, zeroed on every
  dictionary construction, which dominated custom decoders (the `Event` decoder
  allocated fifteen of them per event). It now stores the set fields in a small
  `Vec` behind `get_field`/`set_field`/`from_fields`; the codegen reads and
  writes through those accessors.
- **Parser** (`argonaut-core` port): encounter-order object construction
  (canonical hashing sorts anyway), bulk copies for plain strings and
  stack-buffer number parsing.
- **Runtime layouts**: `RecordFields` gained `into_entries`, `with_capacity`,
  `push` and `entries_unsorted` for the native construction paths.

`SharedRecord` keeps its `Mutex`: threaded mode shares object handles through
`Arc`, so the borrow cannot become a `RefCell`. The codegen and TAST tests were
updated to the sparse construction helpers (`80/80` and `41/42`, the remaining
TAST case needs unrelated `js-bigints` dependencies resolved).

## Remaining gap

Decode still costs 6.7 times the published Go build. The profile attributes the
remainder to per-value representation: one owned string key per decoded record
field, one string copy per decoded string value, `Maybe`/`Either` heap wrappers,
and the `Mutex` in every `Foreign.Object` lookup. The next levers are shared
record keys (`Rc<str>`), unboxed representations for small non-recursive sums,
a byte-level parser and dictionary hoisting.

## Commits

- purust: `3655f80`, `438c54f` (runtime and codegen, edge branch).
- purust ports: `e92885f`, `83d1b11`, `61fc4fb`, `9c80c06` (argonaut-codecs),
  `0af0f61` (argonaut-core), `43b7b56` (arrays), `b7b0bee` (aff).
