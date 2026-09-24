# TAST type-table storage and resolution

This cycle starts from the integrated
[single-pass text decoder](2026-09-24-tast-single-pass-cursor.md). Sources,
preserved builds and measurements live in `var/benchmark/json-tast-types-20260924/`.

The compact-reference candidate improves the complete workload **5.0%** in its
first six-pair campaign and **8.5%** in confirmation, winning **11/12 pairs**.
Combined allocations decrease **37,711,288 → 29,708,136 B (−21.2%)**.
The joint buffer-reuse/compact-reference implementation is retained.
The fresh official combined result is **21.566500 ms**, published as
**21.57 ms** in both benchmark checkouts.

## Implementation

The canonical native decoder in PBO's `CoreFn/Json.go` changes; `Json/Text.go`
is regenerated with the existing checked generator. Both DOM and typed inputs
therefore use the same resolution algorithm and final constructors.

The `buffers` variant:

- Stores settled errors and types in one array, distinguished by the existing
  state bytes. After validating every table entry, the successful array becomes
  the public result directly.
- Compacts the pending-ID list in place, preserving its order. Force resolution
  continues to pick the first remaining pending ID.
- Checks argument dependencies before allocating their final arrays. An
  unresolved argument dominates argument errors; when all arguments are settled,
  their first error in source order wins.
- Removes temporary arrays of resolution results from rows and constraints.
  Their dependency/error checks precede final tuple construction. Function
  argument arrays are likewise constructed after return-type readiness checks.
- Avoids building an unused boxed array around successful argument results.

The `compact-refs` variant additionally shares variant-exclusive payload slots
in the private `ntRef` structure. A single reference slot carries a constructor,
return, element, row, tail or body ID; a single string slice carries ADT name
parts or quantified variables. Static types and mutually exclusive immediate or
deferred errors share one value slot. The flags retain their separate meanings
and the same error precedence.

The arm64 layout inventory records **360 → 152 bytes per `ntRef`**. Across the
corpus's **21,574 entries**, that removes **4,487,392 bytes** of reference storage
before allocator rounding. These are private decode-time records; final TAST
constructors retain their established representation.

The corpus contains **14,578 argument-ID arrays / 24,844 IDs**: their integer
payload is **198,752 bytes** before allocator rounding. The larger reduction in
this candidate comes from the per-entry reference layout, shared state storage
and eliminated resolution-result buffers. Argument IDs continue to be validated
once and retained as integers for subsequent dependency scans.

The final output representations, eager validation of all type-table entries,
cycle fallback and source-usage validation are exercised against the generated
PureScript implementation. Public output arrays own their storage.

## Validation

The existing frozen-module, 2,210 module-mutation, 22 directed type-table,
syntax/index and input-overwrite suites are supplemented by 2,010 deterministic
type-table graphs. They exercise forward, missing and repeated references,
cycles, rows, constraints and combinations of pending dependencies and errors.
Both native input paths are compared with the generated PureScript resolver.
A separate test mutates a returned type-table array to check independence from
other decodes and from its constructor arguments.

Both variants pass the full suite, including **1,493 exact module
errors / 717 successful module mutations** and the new graph differential:
**1,054 successful tables / 956 exact errors**. The graph, syntax, public-boundary,
concurrency and ownership checks also pass under `-race`.

## Measurement protocol

`typed-tast/rebuild.py` replaces the two native decoder bodies in a copy of the
preserved integrated workspace. It checks the public FFI signatures, keeps the
generated wrappers and diagnostic driver, and archives the exact source inputs.
The resulting binaries are built with `go build -pgo=off`.

Paired measurements use `GOMAXPROCS=1`, `GOGC=100`, alternating process order and
all six phase orders. Every timed output is retained and fingerprinted after
timing. Final-value construction, ownership, conversion and normal GC are timed;
fingerprinting and retention-buffer allocation are outside the interval.

## Paired results

Each row is a separate six-pair campaign, with two warmups and five retained
measured samples per phase and process. Cells are medians of process minima;
each campaign verifies **2,160 sampled outputs**.

| Candidate / campaign | Baseline decode | Candidate decode | Baseline combined | Candidate combined | Combined winning pairs |
| --- | ---: | ---: | ---: | ---: | ---: |
| Buffer reuse | 14.402480 ms | 13.725209 ms | 23.351729 ms | 21.959729 ms | 5/6 |
| Compact references | 15.485750 ms | 13.179771 ms | 23.622354 ms | 22.431938 ms | 5/6 |
| Compact confirmation | 14.306917 ms | 12.698771 ms | 23.297438 ms | 21.310938 ms | 6/6 |

The buffer-only combined improvement is **6.0%**. The compact implementation's
decode phase improves **14.9% / 11.2%** across its two campaigns and wins
**11/12 decode pairs**. These separate campaigns establish the complete
candidate's benefit; their cells are not an isolated timing attribution for
individual changes.

Minimum allocations per corpus:

| Build | Parse | DOM decode | Typed combined |
| --- | ---: | ---: | ---: |
| Previous production | 24,405,672 B | 27,068,888 B | 37,711,288 B |
| Buffer reuse | 24,405,672 B | 23,547,272 B | 34,189,672 B |
| Compact references | 24,405,672 B | 19,065,736 B | 29,708,136 B |

The compact implementation saves **8,003,152 B** in both decoder paths, giving
**29.6%** lower DOM-decode allocations and **21.2%** lower combined allocations.
The input index uses the preceding one-pass implementation.

## Official result and lifecycle costs

The rebuilt official diagnostic reports **20.513917 / 12.890459 / 21.566500 ms**
for parse / DOM decode / typed combined. Its allocation minima match the compact
candidate exactly. Three fresh processes validate **540 sampled outputs**. The
previous published cell was **21.93 ms**; the paired campaigns above provide the
controlled incremental comparison.

The fresh-process lifecycle audit uses the same public decoder getter and keeps
two complete corpus results alive until both fingerprints have been checked.
Six pairs validate **288 complete modules**. Medians of the separate observations:

| Phase | Previous production | Compact references |
| --- | ---: | ---: |
| Getter construction | 6.625 µs / 32 B | 4.7915 µs / 32 B |
| First full-corpus decode | 27.767584 ms / 37,718,488 B | 26.992021 ms / 29,715,336 B |
| Second retained full-corpus decode | 41.370812 ms / 37,711,288 B | 29.088521 ms / 29,708,136 B |
| Final release collection | 1.667979 ms / 0 B | 1.793125 ms / 0 B |

Final collection is measured after an untimed collection clears fingerprint
scratch with both output batches still alive. The release observation increases
by **0.125146 ms**. Retained/released heap medians are
**35,216,952 / 5,699,648 B** for the baseline and
**35,217,176 / 5,699,784 B** for the candidate; the final representation has the
same storage requirements within this measurement's runtime variation.

Summing the four measured intervals **within each process**, then taking the
median of those sums, gives **70.790853 → 57.451104 ms**, winning **6/6 pairs**.
This lifecycle subtotal excludes input loading, retention-buffer allocation,
fingerprinting and its preparation collection. It has a different allocator
history from the official phase measurements and is reported independently.

## Compiler and distribution verification

The native bootstrap succeeds on **462 modules / 278,555 type-table entries**.
The rebuilt compiler regenerates **568/568 Go files byte-identically** to Node.
The fresh official diagnostic has **507/507 compiled Go files token-identical**
to the validated and measured compact-reference candidate, ignoring comments and
formatting (`published-parity.json`).

The diagnostic executable changes **9,037,490 → 9,037,394 B**. The compiler
changes **28,964,194 → 28,980,610 B (+0.06%)**. Generated diagnostic Go source
changes **38,573,400 → 38,574,779 B**. The `go.mod` files and bundled JavaScript
diagnostic are byte-identical to the preceding integrated build.

## Provenance

`before.json` pins the starting heads, clean-worktree states, source archive and
the previous integrated binary. Each candidate keeps its exact decoder source
copies, generated Go, build command and manifest. `verify-campaigns.py` checks
the paired artifacts, every report's oracle fingerprints and the aggregates.
`finalize.py` archives the retained sources and verifies the campaigns, official
results, lifecycle audit and compiler/source parity.

The three paired campaigns, official run and lifecycle audit account for
**7,308 retained, fingerprint-checked sampled outputs** in this cycle.

The three excluded diagnostic rows retain their exclusion from totals and `/C`
ratios. The canonical decoder and generated text specialization are integrated
locally; this cycle's source changes have not been committed or pushed.
