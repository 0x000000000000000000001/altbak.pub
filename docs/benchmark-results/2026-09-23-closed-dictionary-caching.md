# Closed dictionary caching in the Go backend

## Motivation

The general-JSON profiling notes (`2026-09-23-json-general-profiling.md`)
measured **2,393 `RecordConsImpl` calls per decode pass**, ≈ 645 µs ≈ 6.5% of
the decode phase, and observed that every call received freshly allocated
dictionary and closure arguments. The per-record plan was therefore rebuilt for
every decoded record, and the "plan is shared by every invocation of the
dictionary" assumption in `Data/Argonaut/Decode/Internal/Record.go` never held.

The cause is structural: an instance dictionary built **inside a lambda** is
re-evaluated on every call. Only top-level bindings receive module-level
`sync.Once` getters (the `Get_...` functions), so dictionaries constructed in
function bodies — including every record decoder used by a custom
`DecodeJson` instance — were rebuilt on each evaluation.

## Implementation

New backend pass `Gopurs.ClosedDictionaries.cacheClosedDictionaries`, run
between `optimizeImmediateApplications` and `Ownership.prepare` in
`Gopurs.CodeGen.translateWithFunctions`.

A node is lifted when all of the following hold:

- it is application-shaped (`App`/`UncurriedApp` after stripping `Typed` and
  `TypeApp`),
- it is closed: `freeVars` reports no captured local (module-level references
  are allowed and expected),
- it contains no effect syntax (`PrimEffect`, `EffectBind`, `EffectPure`,
  `EffectDefer`, uncurried effect forms),
- its annotation is a class type present in `metadata.classDeclsFields`.

The node is moved into a new non-recursive module binding — an ordinary cached
getter, lazily initialised by `sync.Once`, exactly like a top-level binding —
and the original site becomes a typed `Var` reference to it. A top-level
binding body is never lifted again (it is already cached); only constructions
nested below it are.

Files:

- `src/Gopurs/ClosedDictionaries.purs` (new),
- `src/Gopurs/CodeGen.purs` (import and pipeline wiring).

## Validation

- **JSON diagnostics**: the runner validates every successful payload hash and
  the ten malformed-input expectations against the frozen oracle in every
  measured process; all pass with the pass enabled.
- **Compiler fixtures**: `./bin/test ArrayTraverseEither CheckTypeClass
  CyclicInstances CompactRecordConsumers JsonRecordPlan NativeRecordBoxing
  NativeRecordWorkers NestedRecordUpdate Monad DerivingFunctor InstanceChain
  LetInInstance EmptyTypeClass --update-snapshots` — 13 passed, 0 failed. The
  pass produced no getter in any of these modules and twelve snapshots are
  byte-identical to the committed ones. The single changed snapshot
  (`DerivingFunctor.go`, `functorM` record-clone code) is a **stale snapshot**:
  it was committed on 2026-09-21 while `src/Gopurs/GoCode.purs` changed on
  2026-09-23, so the diff predates this pass.
- **Compiler bootstrap**: the native compiler itself (459 modules) was
  regenerated with the pass and executed.

## Measured effects

JSON decoding (general-JSON diagnostic), paired counterbalanced campaign,
three processes per binary, `GOMAXPROCS=1`, `GOGC=100`, median of process
minima. The machine carried external load during the campaign; the paired
deltas are the stable part:

| Phase | baseline | with caching | delta |
|---|---:|---:|---:|
| parse | 2,500.5 µs | 2,373.6 µs | −5.1% (noise) |
| decode | 11,518.0 µs | 9,850.9 µs | **−14.5%** |
| combined | 15,150.7 µs | 14,329.1 µs | **−5.4%** |

Allocations per corpus:

| Phase | baseline | with caching | delta |
|---|---:|---:|---:|
| parse | 4,612,800 B | 4,612,800 B | 0 |
| decode | 16,556,992 B | 15,037,392 B | **−9.2%** |
| combined | 21,169,776 B | 19,650,176 B | **−7.2%** |

Isolated parse comparison (best of 4,000 passes, interleaved) shows no parse
change: 2,311–2,320 µs with caching against 2,324–2,330 µs baseline.

Instrumented copies of otherwise identical builds: `RecordConsImpl` calls drop
from **91,368 to 18** over the same 41 decode passes; only the nested
dictionary that the purchase branch actually rebuilds is lifted
(`Test_JsonDecoding` gains exactly one cached getter).

JSON to Typed AST: allocations are byte-identical (parse 56,031,272 B, decode
27,078,264 B, combined 83,109,472 B) and paired timings show no effect outside
that suite's large process spread (combined process minima range 46–59 ms in
both binaries). That decode path is already native, so the pass has nothing to
remove there; the change is neutral.

## Rejected extension: record literals and constructors

An extended variant also accepted class-typed record literals (`Lit (LitRecord
_)`) and constructor applications (`CtorSaturated`). In the JSON workspace it
raised the number of lifted getters from 19 to 568. A paired campaign against
the narrow pass showed only **−1.3%** of decode allocations (7,363 → 7,264 MB
per 500 passes, vs −9.2% from the narrow pass) while decode and combined times
moved **+1.4% / +2.2%**. The extra module-level getters cost more than the
sharing saves, so the narrow application-only shape is kept.

## Limits

- Effect syntax is excluded. A construction that smuggled mutable state through
  `unsafePerformEffect` could observe the new sharing; no such case appears in
  the validated corpora, and the existing top-level binding cache already
  assumes dictionaries are values.
- The lifted value is built lazily on first use: a site evaluated once pays one
  extra getter call and keeps the dictionary alive afterwards. Workloads that
  decode many documents amortise this immediately.
- Only class-typed, closed, application-shaped constructions are lifted.
  Record-literal dictionaries and partially applied dictionaries are left
  alone: the measured extension above shows the broader shape is not
  profitable here.

## Provenance

- Pass: `gopurs/gopurs/src/Gopurs/ClosedDictionaries.purs`,
  `src/Gopurs/CodeGen.purs`; compiler built by `npm run build:native`.
- Baseline binary: `var/benchmark/json-dec-baseline-20260923` (previous
  compiler), cached binary: `var/benchmark/json-dec-cache3-20260923` (final
  pass with the nested-only restriction).
- Campaigns: `var/benchmark/json-dec-cache3-paired-results-20260923`,
  `var/benchmark/json-tast-cache3-paired-results-20260923`; driver
  `src/Test/JsonDecoding.go` profile mode and `scratch/paired-json-campaign.py`.
- Instrumented counter builds: `scratch/json-dec-cache3-count-20260923`
  (patched copy, not a build source).
