# Value-level accessors: keep `caseJson*` out of monomorphization

First landed piece of the "fusion" target (avoid boxing values that are
immediately destructured). The dictionary caching
(`2026-09-23-closed-dictionary-caching.md`) removed the per-record plan rebuild;
this change removes the boxing round trip around the Argonaut accessors.

## Diagnosis

`Data.Argonaut.Core` defines the accessors as thin wrappers over value-level
FFI:

```purescript
foreign import _caseJsonNumber :: Json -> (Number -> Json) -> Json -> Json
caseJsonNumber :: forall a. a -> (Number -> a) -> Json -> a
caseJsonNumber d f j = unsafeCoerce (_caseJsonNumber (unsafeCoerce d) (unsafeCoerce f) j)
```

PBO's monomorphizer specializes `caseJsonNumber` at `a = Either
JsonDecodeError Number` (and `caseJsonString` the same way), substitutes the
static continuation (`Right`), and gives the specialization a **native ADT
return representation** because the instantiated type is a concrete ADT. The
worker then destructures the boxed `Either` it just produced into the native
struct, and the cached wrapper re-boxes it — a `boxed → native → boxed` round
trip with two extra allocations and two conversions per call:

```go
// worker (before)
_v := Apply3(Get_Data_Argonaut_Core__caseJsonNumber(), d_0, Get_Data_Either_Right(), j_2)
if _v is Right { return struct{V1: right.V0, V2: true} }
return struct{V0: left.V0, V2: false}
// wrapper (before)
_v := Call_..._caseJsonNumber__627409332(d_box, f_box, j_box)
if _v.V2 { return Right{_v.V1} }
return Left{_v.V0}
```

Measured on the JSON workspace before the change:
`Get_Data_Argonaut_Core_caseJsonNumber__….func1.1.1` 298.5 MB and
`caseJsonString__….func1.1.1` 238.0 MB flat allocation over 500 decode passes
(≈ 7.3% of decode allocations).

## Change

`Gopurs.Monomorphization.monomorphizeModulesWith` now excludes the accessor
family from specialization, next to the existing `sharedRecordWorkers` filter:

```purescript
boxedFfiAccessors = Set.fromFoldable
  [ "Data.Argonaut.Core.caseJsonNull", "Data.Argonaut.Core.caseJsonBoolean"
  , "Data.Argonaut.Core.caseJsonNumber", "Data.Argonaut.Core.caseJsonString"
  , "Data.Argonaut.Core.caseJsonArray", "Data.Argonaut.Core.caseJsonObject"
  ]
```

A specialization of a thin FFI forwarder cannot unbox anything — the FFI
boundary is value-level by construction — so the only effect was the
round trip. With the exclusion, the accessor is a plain pass-through:

```go
// worker (after)
return gopurs_runtime.Apply3(Get_Data_Argonaut_Core__caseJsonNumber(), d_0, f_1, j_2)
```

The four specializations (`caseJsonNumber__627409332`,
`caseJsonNumber__1623957941`, `caseJsonString__28168628`,
`caseJsonString__3780298677`) disappear from the generated workspace.

## Measured effects

JSON decoding diagnostics, same protocol as before; the machine carried
external load during the campaigns, so paired deltas and allocation totals are
the stable parts.

| Comparison | decode | combined |
|---|---:|---:|
| `cache3` → `cache5` (paired, isolates this change) | **−17.6%** | −6.1% |
| baseline → `cache5` (paired, cumulative) | **−19.1%** | −9.5% |

Allocations per corpus (deterministic):

| | baseline | after dictionary caching | after this change |
|---|---:|---:|---:|
| decode | 16,556,992 B | 15,037,392 B | **13,490,152 B** |
| combined | 21,169,776 B | 19,650,176 B | **18,102,936 B** |

Cumulative decode allocations: **−18.5%**.

Remaining decode allocation leaders after the change: `Right` 12.3%, `Just`
12.0% (semantically required in the Argonaut representation — the ABI-level
target), lookup adapters 10.8%, `decodeForeignObject` 7.7%, record plan 7.5%,
`decodeFieldId` 4.8%, `Box` 4.0%, record dictionaries 4.6%.

## Validation

- JSON diagnostics: exact successful fingerprints and the ten malformed-input
  expectations validated against the frozen oracle in every measured process.
- Compiler fixtures: `./bin/test ArrayTraverseEither CheckTypeClass
  CyclicInstances CompactRecordConsumers JsonRecordPlan NativeRecordBoxing
  NativeRecordWorkers NestedRecordUpdate Monad DerivingFunctor InstanceChain
  LetInInstance EmptyTypeClass --update-snapshots` — 13 passed, 0 failed; none
  of their snapshots changed (`DerivingFunctor.go` still carries the
  pre-existing stale diff from the 2026-09-23 emission work).
- JSON to Typed AST: allocations byte-identical within rounding (parse
  56,031,272 B, decode 27,075,384 B against 27,078,264 B, combined
  83,106,592 B against 83,109,472 B): neutral, that path does not use the
  accessors.
- Compiler bootstrap: the native compiler itself (459 modules) was regenerated
  with the change and executed.

## Limits

- The exclusion is name-based and deliberately narrow: it lists the six
  accessors known to be thin value-level FFI wrappers. Extending the same rule
  to other forwarders requires identifying them mechanically (a body-shape
  check) before trusting a wider filter.
- The JavaScript backend is unaffected (the filter lives in the Go backend's
  monomorphization wiring); semantics are unchanged because the polymorphic
  accessor is the ordinary definition.

## Provenance

- Change: `gopurs/gopurs/src/Gopurs/Monomorphization.purs`.
- Workspaces: `var/benchmark/json-dec-cache3-20260923` (before),
  `var/benchmark/json-dec-cache5-20260923` (after).
- Campaigns: `var/benchmark/json-dec-cache5-vs-cache3-results-20260923`,
  `var/benchmark/json-dec-cache5-vs-baseline-results-20260923`,
  `var/benchmark/json-dec-cache5-results-20260923` (runner-validated).
