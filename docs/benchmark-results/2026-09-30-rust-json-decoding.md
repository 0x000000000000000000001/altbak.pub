# Native record plans for purust JSON decoding

The Rust benchmark table had no JSON Decoding cell. The Go and C cells were
produced by the preserved-build protocol of
`var/benchmark/json-reference-owned-20260925` (six phase-order permutations,
median of six process minima, one binary per backend). This cycle gives purust
the same cell and improves the backend by an order of magnitude on the decode
phase.

## Result and protocol

The preserved `benchmark-go` (text-schema build) and `benchmark-c` binaries from
the `json-reference-owned` cycle and the current purust build were re-run in one
session on the same corpus, with `DIAG_PHASES` cycling through the six phase
permutations. Each cell is the median of the six process minima. All reports
validate against the fixture oracle (`fingerprints` and `json_fingerprints`,
17 modules, error cases included).

| Phase | C reference | Go published | purust before | purust now |
|---|---:|---:|---:|---:|
| parse | 381 | 2 012 | 3 506 | **2 559** |
| decode | 268 | 739 | 40 506 | **11 401** |
| combined | 658 | 2 406 | 43 859 | **14 800** |

The published Go and C cells come from the same campaign protocol (2 241.42 and
657.69); the re-run reproduces them within session noise. The JavaScript cell
(9 283.94) is from the official harness; in the current harness the same corpus
measures 9 165. In the official harness run of this cycle the purust binary
reaches **14 083 combined** with the generic Go build at 11 294 in the same
sequence, which is the reproducible path
(`bin/benchmark/json-diagnostic.py build` and `measure --runtime rust`).

## Implementation

The generic record path rebuilt every record through the `GDecodeJson`
dictionary chain: one class dispatch, one field lookup and one record insert per
field, with an intermediate `Either` per field. The purust port of
`argonaut-codecs` now builds an immutable plan for a record row — field name,
native kind and the exact generic step of every field — and runs it natively:

- `FieldSpec` covers `Int`, `Number`, `String`, `Boolean`, `Json`, `Maybe`,
  `Array` and nested record plans; anything else (user instances) stays
  `Custom`.
- The runner decodes a field natively while the value has the declared shape
  and otherwise calls the field's step, so anything unrecognized keeps the
  ordinary behavior. Errors are produced only by the step, which owns the exact
  `AtKey` wrapping, so the fingerprints stay bit-identical.
- `getField`, `getFieldOptional` and `getFieldOptional'` gained the same
  treatment: missing keys, `null` and successful decodes are answered natively
  and failures keep the generic accessor.

Supporting changes:

- The parser keeps encounter order when building objects (canonical hashing
  sorts anyway), copies plain JSON strings in bulk, and parses ordinary numbers
  from a stack buffer.
- The runtime gained `RecordFields::into_entries`, `with_capacity`/`push` for
  duplicate-free native construction, and `SharedRecord` now borrows through a
  `RefCell`: values never cross threads, the cell only guards re-entrant
  mutation.

## Remaining gap

Decode still costs 15.4 times the published Go build. The profile attributes the
remainder to per-value allocation and representation costs: one owned string key
per decoded record field, `Maybe`/`Either` heap wrappers, `Value` moves through
the deep clone in `Foreign.Object.get`, and the decoder dictionaries rebuilt on
every call (about 2% of the phase). The next levers are shared/`Rc<str>` record
keys, unboxed representations for small non-recursive sums, a byte-level parser
and a dictionary hoisting pass.

## Commits

- purust: `3655f80` (runtime, edge branch) — record enumeration, growth and
  borrowing.
- purust ports: `e92885f`, `61fc4fb`, `9c80c06` (argonaut-codecs) and `0af0f61`
  (argonaut-core).
