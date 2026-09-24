# TAST: single-pass validation/indexing and typed scalar reads

This cycle continues the typed-cursor experiment from
[schema decoder specialization](2026-09-24-schema-decoder-specialization.md).
Artifacts are preserved under `var/benchmark/json-tast-cursor-20260924/`.

The integrated one-pass decoder improves the complete JSON → TAST workload
**40.659750 → 22.376709 ms (−45.0%)** and allocations **26.7%** in the final
six-pair campaign. The initial candidate and its confirmation improve **43.4%**
and **44.5%**, respectively: **18/18 pairs** improve across the three campaigns.
The official Go result is **21.927333 ms**, published as **21.93 ms**.
This replaces the preceding two-pass typed-cursor candidate, whose elapsed-time
gain did not reproduce.

## Implementation

The typed-input decoder continues to reuse the production native decoder's
control flow, final constructors, eager type-table resolution and source-usage
validation. The JSON input layer changes:

- Syntax validation and structural indexing run in one pass. Ignored members
  receive the same syntax and finite-number checks as consumed values.
- The index retains twelve-byte tokens. Array tokens carry the element count;
  object tokens point to their last key, and key tokens link backwards within
  their object. Lookup stops at the first reverse match, preserving last-value
  precedence without scanning later fields.
- String tokens record whether unescaping or invalid-UTF-8 normalization is
  needed. Ordinary field-name comparisons use validated source slices directly.
- Small integer reads avoid `strconv.ParseFloat` while retaining exact values
  and negative zero. Decimal, exponent and larger integer forms use the original
  conversion, including rounding and overflow behavior.
- Six statically selected discriminator sites borrow temporary text. All names,
  paths, literals and other strings stored in the final module remain owned.

Schema decoding begins only after the complete document validates. Inputs beyond
the private index's size limit and syntax failures use the existing public
parser/decoder boundary. The cold `sourceSpan` subtree still uses the existing
Argonaut decoder through a materialized subtree.

`one-pass` preserves the initial one-pass/back-link implementation. `scalars`
adds direct short-integer reads and borrowed discriminators. `candidate` has the
same decoding engine as `scalars`, with a phase-order-aware measurement driver.

## Production integration

PBO now exposes `PureScript.Backend.Optimizer.CoreFn.Json.Text.parseModule`,
returning `Either String (Module Ann)`. `App.readCoreFnModule` calls this API.
The diagnostic's combined phase calls the same public API and returns the same
usable module as its DOM decoding phase.

`purescript-backend-optimizer-gopurs/bin/json-text/generate.py` derives the Go
implementation from the canonical native `CoreFn/Json.go` decoder and local
cursor/index templates. The generated file is distributed with the library;
there is no runtime schema compilation or dependency on the benchmark checkout.
`--check` verifies that it matches its inputs. Token-level comparison confirms
that all **173 engine declarations** match the measured prototype; integration
adds the public parser fallback, error rendering and ordinary FFI wrappers.

The native compiler bootstrap succeeds on **462 modules / 278,555 type-table
entries**. Native regeneration using the integrated module reader emits
**568/568 Go files byte-identically** to Node.

## Measurement controls

`dom-control` and `candidate` are built from the same preserved production DOM
workspace with identical driver changes and `go build -pgo=off ... ./main` flags.
The only decoding change is the combined path's typed-input implementation.

The driver accepts and reports `DIAG_PHASES`; both manifests declare the
capability. The paired runner rotates and reverses phase order only when both
TAST builds support it, and checks each reported order. Legacy TAST workspaces
retain their original fixed-order protocol.

Every sampled output is retained and fingerprinted after timing. Construction,
copying, conversion and normal Go GC are included. Fingerprinting, retention
buffer allocation and a forced final heap drain remain outside the interval.
The environment remains `GOMAXPROCS=1`, `GOGC=100`.

## Paired results

These are medians of six process-minimum times, with five retained measured
passes after two warmups per phase. Each campaign checks **2,160 sampled outputs**.

| Campaign | DOM combined | Typed combined | Change | Winning pairs |
| --- | ---: | ---: | ---: | ---: |
| Initial candidate | 40.104271 ms | 22.703271 ms | −43.4% | 6/6 |
| Confirmation | 42.160979 ms | 23.399625 ms | −44.5% | 6/6 |
| Integrated public API | 40.659750 ms | 22.376709 ms | −45.0% | 6/6 |

Both campaigns report minimum allocations of **51,476,800 → 37,585,832 B**
per twelve-module corpus, a reduction of **13,890,968 B (27.0%)**. The ordinary
parse/decode controls retain **24,405,672 / 27,071,192 B**, respectively.
Phase timings come from different allocator histories and are not additive.

The final integrated build allocates **37,711,288 B** per combined corpus:
**13,765,512 B (26.7%)** less than DOM, and **125,456 B** more than the experimental
driver. The public FFI/caller boundary and regenerated surrounding code are
included in this final result. The final parse/decode controls allocate
**24,405,672 / 27,068,888 B**.

The separate official three-process Go run reports parse/decode/combined
**20.241041 / 13.664167 / 21.927333 ms**, with **540 sampled outputs** checked.
The JavaScript public-API run also validates **540 sampled outputs** against
the same frozen oracle. Its measured combined cell is **85.969167 ms**;
this unpaired validation run does not estimate a JavaScript performance change.

## Initialization, ownership and release

The fresh-process construction audit alternates six processes per build. It
measures the text-to-module getter/callback setup before decoding, then retains
two complete corpus results before fingerprinting: **288 checked modules**.

| Audit phase | DOM | Integrated |
| --- | ---: | ---: |
| Getter/callback construction | 4.2295 µs / 64 B | 5.0830 µs / 32 B |
| First complete pass | 41.703625 ms / 51,484,472 B | 26.115271 ms / 37,718,488 B |
| Second retained pass | 50.848813 ms / 51,476,816 B | 38.735688 ms / 37,711,288 B |
| Final release collection | 1.774458 ms | 1.598167 ms |

After fingerprinting, an untimed collection clears its scratch allocations while
both module batches remain alive. The audit then drops them and measures a final
collection separately. Median retained heaps are **35,298,828 / 35,216,848 B**;
after release they are **5,699,272 / 5,699,704 B**, respectively, with input strings
and decoder callbacks still live. This schedule differs from the steady-state
campaign and its cells must not be added to the official medians.

Getter construction occurs inside `main`. A separate `inittrace` inventory
covers earlier generated-package initialization: **10,376 B / 67 allocations**
for both builds. These two inventory processes also validate their two retained
passes (**48 additional outputs**); their single clock observations are not
used as a timing comparison. Per-document validation/index construction remains
inside every combined timed interval.

## Validation and preservation

The existing twelve-module oracle, 2,210 deterministic schema mutations,
type-table argument/error tests, index differential tests and ownership tests
are supplemented by 20,015 syntax/mutation cases, six depth-limit cases,
bit-exact numeric checks and complete-module input-overwrite checks.

Both the initial one-pass engine and the scalar-read refinement pass the full
differential suite: **1,493 exact errors and 717 successful mutations**, the
**twelve frozen modules**, and **22 directed type-table cases**. The index and
cursor tests also pass under `-race`.

The final candidate passes the scalar-bit, syntax and ownership checks under
`-race`. All **twelve complete modules** retain identical exhaustive fingerprints
after their private input buffers are overwritten. `candidate-parity.json`
verifies that its 508 compiled Go files differ from the fully differential-tested
scalar build only in the diagnostic's phase-order handling.

The integrated public API passes the same full differential suite. An additional
**2,025 parser/schema error cases** match `parseModulePS` exactly, including
invalid ignored members, number overflow, duplicate/escaped keys and malformed
Unicode. These public-boundary tests, the syntax/index tests, scalar-bit checks,
concurrency checks and complete-module input-overwrite tests also pass under
`-race` (`integrated-validation.log`).

An initial validation run encountered a nearly full disk. Verified APFS
copy-on-write clones replaced 39,673 identical file copies across the preserved
benchmark directories, sharing 4,086,797,414 logical bytes. Their contents were
checked with SHA-256, and all paths and workspaces were retained. The interrupted
validation log is preserved; subsequent validation runs use the recovered space.
The Go build cache later reached 54 GB and was cleared after another disk-full
build failure. The resumed integrated build succeeds; benchmark workspaces and
their pinned artifacts remain preserved.

## Distribution

The diagnostic executable grows **8,789,106 → 9,037,490 B (+2.83%)** relative to
the phase-order-aware DOM control. The native compiler grows
**28,735,410 → 28,964,194 B (+0.80%)**. The diagnostic `go.mod` files are
byte-identical. Generated diagnostic Go source grows
**38,033,804 → 38,573,400 B**, including the public module and its generated
type annotations. Exact hashes are in `integrated-distribution.json`.

## Scope

The retained application-decoder specialization remains independently enabled.
This cycle changes PBO's text-to-module entry point and the TAST diagnostic's
combined path. `decodeModule :: Json -> ...` remains the public decoder for
already parsed JSON and arbitrary callers. Public JSON object ownership is
preserved by the existing parser boundary.

The three diagnostic rows stay excluded from benchmark totals and `/C` ratios.
The existing C++ source-usage-validation and result-release scope differences
are documented in the [comparison audit](2026-09-24-comparison-audit.md); these
paired Go results do not establish a new native-reference ratio or timing floor.

## Provenance and reproduction

`sources-before/`, `sources-before-integration/` and `sources-after/` retain the
source snapshots. The paired campaign reports pin their binaries, emitted Go,
manifests, corpus, oracle, phase schedules and process load observations.
`finalize.py` verifies these artifacts, official Go/JS results, both lifecycle
audits, compiler regeneration parity and production/prototype token parity.

From the benchmark checkout, a fresh integrated build and its measurements use:

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace NEW_WORKSPACE
python3 bin/benchmark/json-diagnostic.py measure --suite JsonTypedAst --workspace NEW_WORKSPACE --runtime go --output NEW_RESULTS
python3 bin/benchmark/json-diagnostic/paired.py --suite JsonTypedAst --baseline var/benchmark/json-tast-cursor-20260924/dom-control --current NEW_WORKSPACE --output NEW_PAIRED_RESULTS
```

The public API, PBO module reader and rebuilt native compiler are integrated
locally. Changes are not yet committed or pushed.
