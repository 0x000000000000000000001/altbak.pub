# Known-decoder specialization and typed-cursor TAST experiment

This cycle follows [borrowed event objects](2026-09-24-borrowed-event-objects.md).
The application baseline is the preserved
`json-direct-20260924/final-json`; the TAST baseline is the production compact-DOM
build `json-parser-20260924/fused-tast`.

Artifacts are under `var/benchmark/json-schema-20260924/`. The application build
selected for publication is `published-json`; the typed-cursor experiment is
`validated-typed-tast`. Earlier pilots, failed builds and the initial TAST driver
boundary failure are preserved separately.

## Application result: retain the specialization

The final six-pair campaign (`paired-published-json`) improves both decoding and
the combined operation in **all six pairs**:

| Phase | Borrowed-object baseline | Schema workers | Change |
|---|---:|---:|---:|
| Parse control | 2.411876 ms | 2.247521 ms | −6.8% |
| Decode | 1.833438 ms | **0.864021 ms** | **−52.9%** |
| Parse + decode | 4.546292 ms | **3.314354 ms** | **−27.1%** |

| Minimum allocated bytes per corpus | Baseline | Schema workers | Change |
|---|---:|---:|---:|
| Parse control | 2,435,696 | 2,435,696 | unchanged |
| Decode | 2,549,944 | **1,876,160** | **−26.4%** |
| Combined | 4,985,640 | **4,311,856** | **−13.5%** |

The parser implementation is unchanged; its timing variation is a control,
not a parser optimization. One-minute load averages during this campaign span
**20.66–23.56**. The earlier six-pair application campaign also improved decode
and combined in every pair (1.645 → 0.778 ms and 3.812 → 3.154 ms). Its emitter
predates the discarded-value compilation fixes and is recorded separately.

The official Go/JavaScript/C campaign (`official-published-json`) validates the
frozen oracle and reports Go **2.315250 / 0.877833 / 3.189625 ms** for
parse/decode/combined. README publishes the rounded **3.19 ms** combined cell.
This is a separate schedule from the paired campaign, not its final paired cell.

### Aligned usable-value reference

`paired-published-reference` alternates generated and hand-specialized decoding
within each process and pairs the old/new generated builds:

| Decode | Generated | Specialized | Generated / specialized |
|---|---:|---:|---:|
| Before | 1.842917 ms | 0.691854 ms | 2.66× |
| After | **0.917063 ms** | 0.705208 ms | **1.30×** |

In that same final campaign, combined time is **3.149084 ms generated** versus
**2.997375 ms specialized**: the reference is **4.8% lower**. Decode allocations
are 1,876,160 versus 1,843,976 bytes, a remaining difference of **32,184 bytes
(1.75% of the reference)**. Combined allocations are 4,311,856 versus 4,279,672
bytes. The final values, parser input representation and ownership costs are
aligned in this comparison.

The ordinary paired and aligned-reference campaigns validate respectively
**900** and **1,800** sampled outputs, in addition to warmups and setup.

## Compiler and codec implementation

`Gopurs.DecoderSchemas` runs after dictionary caching, immediate-application
reduction and borrowed-object analysis, before ownership preparation. It emits
native workers from **resolved decoder dictionaries and function bodies**:

- Standard scalar, array, optional and record dictionaries identify actual
  decoding operations. The result type alone never selects a decoder.
- Record fields execute in dictionary order. A symbol is folded only when its
  actual `reflectSymbol` method resolves to a one-argument lambda returning a
  literal string. Dynamic symbols keep the original record path, including
  lookup-time evaluation and successful reverse-order insertion.
- Small final records use compact runtime records. Duplicate constant labels
  preserve the head field's value and the reverse-insertion slot order; fields
  whose final values are overwritten still decode in the original order.
- A bounded interpreter recognizes read-only custom decoder bodies consisting
  of a proven identity-object borrow, saturated public field reads, forwarding
  of exact `Left` payloads, string-equality branches and actual constructor
  expressions. It derives the event branches from their definitions; no
  application field names, event labels or constructor tags are embedded in the
  compiler pass.
- Successful constructor expressions go back through the ordinary code
  generator, preserving native layouts, annotations, constructor identities and
  necessary reboxing. Public results retain their normal representation.
- Unknown field/element methods retain their callbacks. Captures, recursive
  initialization scopes, shadowing producers, unknown builders, altered error
  continuations and unsupported custom bodies decline the transformation.

The codec's versioned `schemaDecoderABI1` marker gates generation. The native
helper checks the original decoder's tag and schema shape when the worker is
installed, retaining the original decoder if the guard fails. Metadata remains
immutable. Public field readers can invoke an installed worker directly.
The JavaScript library retains its ordinary decoding path.

Untimed instrumentation of the five timed application documents counts:

| Emitted operation | Invocations |
|---|---:|
| Record construction | 6,355 |
| Proven custom variant worker | 2,477 |
| Array worker | 4,635 |
| Present optional worker | 3,528 |

There are no generic record-plan decodes, custom event dispatches, identity-object
copies or public field-accessor calls left in this corpus's decoding pass. The
emitted event worker performs the same 7,431 field reads directly. Arbitrary
decoders outside this recognized subset remain supported.

## Typed-cursor TAST experiment

`bin/benchmark/json-diagnostic/typed-tast/` mechanically specializes the existing
native TAST decoder's JSON input operations. Its schema control flow, error
precedence, type-table resolution, constructors and source-usage validation come
from the same native decoder as the compact-DOM baseline.

The new inputs are concrete `tcCursor`, `tcArray` and `tcIterator` values. Array
iteration reads the token index directly, with no temporary JSON-element slice
and no `any` cursor boxes on the decoding path. Object lookup supports arbitrary
field order and last-value duplicate-key precedence. Final strings are copied
into owned storage. The validated index and source text are call-local.

The prototype retains full syntax validation, including ignored members and
numeric overflow checks. It uses the previous twelve-byte-token index, so it
still incurs both validation and indexing passes. The cold `sourceSpan`
boundary materializes its small subtree for the existing Argonaut tuple decoder.
The type table is resolved eagerly, and the ordinary source-usage validator runs
on the completed module.

Only the diagnostic's combined operation selects this experimental path. Parse
and decode controls retain the production DOM implementation. An initial driver
adapter returned the module's `Either` envelope instead of its payload; the
retained-output check rejected it before a complete campaign. The corrected
adapter unwraps the successful result exactly like `Test.JsonTypedAst.decode`.
Schema/parser failures use that callback's original public error boundary.

### TAST decision: keep the typed cursor experimental

Two independent six-pair campaigns produce conflicting elapsed-time results:

| Combined operation | Compact DOM | Typed cursor | Change | Faster pairs |
|---|---:|---:|---:|---:|
| First campaign | 45.989334 ms | 43.086896 ms | −6.3% | 4/6 |
| Confirmation | 38.551626 ms | 43.494521 ms | **+12.8%** | **1/6** |

Allocation minima agree in both campaigns: **51,476,800 → 43,439,824 bytes**
per combined corpus, a reduction of **8,036,976 bytes (15.6%)**. The parse/decode
controls allocate 24,405,672 / 27,071,192 bytes in both binaries.

Each campaign validates **2,160 sampled outputs**, for **4,320** across the two.
Load averages span 13.48–21.77 in the first and 25.54–36.45 in confirmation.
The retained source and independent ownership are correct, but these measurements
do not establish a repeatable elapsed-time improvement. The production path and
README's TAST cell therefore continue to use compact DOM. The prototype is
preserved for investigating validation/indexing passes, repeated key scans and
the cold materialization boundary; the allocation reduction is not a speedup
claim.

## Validation

- Compiler tests: **33/33**, including generated Go execution and a `-race`
  compilation/execution regression for overwritten duplicate-label fields.
- Codec helper tests pass under `-race`.
- The expanded `test/typed-plans.mjs` passes JavaScript and freshly generated Go
  under `-race`. It covers raw and constructed JSON, nested arrays/records,
  optionals, custom callbacks, variant branches, exact errors, direct variant
  fields, checked-but-discarded reads and object-only custom decoders.
- Application differential: **829/829** old/new Go results match. **818** also
  match JavaScript. The same eleven previously documented native-64-bit versus
  JavaScript-32-bit `Data.Int` boundary cases remain; there are no new differences.
- Typed TAST: the twelve frozen modules and **2,210** deterministic mutations
  agree with generated PureScript decoding: **1,493 exact errors**, **717
  successes**. All **22** directed type-table argument/error cases pass on the
  typed-input resolver itself.
- Index tests: **2,031** directed/generated/truncated cases, plus ownership and
  concurrent independence checks. Typed cursor reads add **1,006** cases,
  including normalized duplicate keys and owned strings. Index and cursor
  tests also pass under `-race`.
- Final native bootstrap: **461 modules, 278,496 type-table entries**. The native
  compiler regenerates **566/566 Go files byte-identically** to Node.
- The sampled-output audit's deliberately corrupted-result regression passes.
- The rebuilt production-DOM TAST diagnostic passes the frozen oracle and
  validates **180** retained sampled outputs (`production-tast-check`).

The application differential used `variant-json`, and the allocation profile
used `final-json`. Their 326 compiled Go files match after formatting. The
publication build additionally emits discarded-value statements to support
unused reads/overwritten labels; removing those statements and formatting yields
the same Go. Both comparisons are recorded in `source-parity.json` and
`published-source-parity.json`. The publication build has its own paired and
official campaigns.

## Decoder construction and distribution costs

The repeated-pass timings above use initialized dictionaries and workers. A
separate audit (`construction-audit/`) measures the public `decode` getter in six
fresh processes per build, before parser and fingerprint getters execute:

| One-time decoder construction | Baseline | Schema workers |
|---|---:|---:|
| Median elapsed time | 68.000 µs | 73.833 µs |
| Median allocated bytes | 42,400 | 42,632 |

Thus worker installation adds **232 bytes** in this initialization measurement.
Its small timing difference is descriptive, not a separate speed claim. The
audit also retains and fingerprints the first two decoding passes, **120 outputs**
in total. First-pass decoding allocates 2,568,392 bytes before and 1,876,336 bytes
after. Parsing, fingerprinting, result-buffer allocation and process startup are
outside those intervals. Decoder construction is not folded into the warmed
campaign cells or silently treated as free.

`distribution.json` records the exact binary and emitted-source costs:

| Artifact | Before | After |
|---|---:|---:|
| Application executable | 6,555,090 B | 6,645,186 B |
| Application compiled Go source | 8,367,765 B | 8,436,829 B |
| Native compiler executable | 28,106,162 B | 28,735,410 B |
| Experimental TAST executable | 8,789,058 B | 8,982,338 B |

The application executable grows **90,096 bytes (1.37%)**, including the guarded
fallback path. Diagnostic `go.mod` files are byte-identical; this cycle adds no
external Go dependency. Compiler sizes use this cycle's archived starting binary
and final native bootstrap, while diagnostic binaries use the aligned official
build flags.

## Remaining allocation sites

`allocation-profile/` contains a 300-pass decode allocation profile and both
`alloc_space` and `alloc_objects` summaries. The largest flat allocation sites
are the final `RecordDict5` and `RecordDict3` values; output array buffers,
`Array` wrappers, `Maybe` constructors and event-constructor reboxing follow.
The remaining extra allocation relative to the specialized reference is small
enough that final representation/conversion costs now matter.

This sampled process-lifetime profile also includes setup and the validating
pass. It is not an exact partition of timed bytes, and allocation shares do not
establish CPU shares. CPU attribution through `runtime/pprof` remains unreliable
on this machine.

## Measurement contract

The paired campaigns use preserved workspaces, alternating process order,
`GOMAXPROCS=1`, `GOGC=100`, two warmups and five retained-output samples per phase.
The application phase order rotates and reverses; the existing TAST driver has
a fixed phase order. Cells are medians of process minima, never sums of phase
medians. All diagnostic Go binaries use aligned `go build -pgo=off ... ./main`
flags.

Every sampled application and TAST result is retained and fingerprinted after
the timed interval. Timing includes construction, conversion, copying and normal
Go GC; it excludes fingerprinting, retention-buffer allocation and a forced
final heap drain. The aligned schema-specific reference constructs the same
usable PureScript values. Its measurements demonstrate implementation headroom,
not a language-performance floor.

The three excluded diagnostic rows remain outside benchmark totals and `/C`
ratios. Historical JavaScript and C/C++ reference cells retain their documented
scope. The outstanding native-Go/C++ lifecycle and owned simdjson-to-Go transfer
limitations from the [comparison audit](2026-09-24-comparison-audit.md) still apply.

## Provenance

`before.json` and `sources-before/` preserve the starting revisions and sources.
`after.json` and `sources-after/` preserve the final compiler, codec, benchmark,
fixture and report sources, including the actual compiler binaries. Each paired
campaign records both manifests, executable hashes, emitted Go hashes, frozen
corpus/oracle hashes, process schedule and complete measurement reports.

`finalize.py` verifies these archives against the working sources, rechecks the
preserved campaign binaries and emitted Go, validates the saved outputs against
the frozen oracles, and records `final-verification.json`. Failed/pilot builds
remain identifiable in their original directories and are not publication data.
