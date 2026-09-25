# JSON references with conventional owned storage

The two JSON rows in the **C reference** table are hand-written **C++/simdjson**
programs. Their former benchmark-specific arenas have been replaced with ordinary
owned C++ representations. The README now explicitly excludes hand-added arenas,
pools and benchmark-specific cross-iteration storage recycling, and requires
construction, copying, conversion and eventual reclamation to be accounted for.

Artifacts are preserved under `var/benchmark/json-reference-owned-20260925/`.
The original sources and binaries are in `sources-before/` and `previous/`.
The existing Go production binaries are reused byte-for-byte; this cycle changes
the reference implementation and refreshes the comparison measurements.

The final paired construction results are **19.977959 / 10.823125 ms** (Go/C++)
for TAST and **2.241417 / 0.657688 ms** for application JSON: **1.85× / 3.41×**.
The stricter fresh-process lifecycle audit, including two retained decodes,
deferred garbage and final release, gives **2.32× / 4.09×**. The two scopes are
reported separately rather than presenting construction time as total lifetime
cost.

## Implementation and scope

- Application results own `std::string`, `std::vector` and `std::optional`
  storage, allocated through the ordinary standard-library allocator.
- TAST results own strings and vectors, `unique_ptr` expression/binder trees,
  and `shared_ptr` type nodes where the type table introduces real sharing.
  Raw table temporaries are destroyed during decoding; referenced types remain
  owned by the returned annotations and declarations.
- The TAST decoder now resolves **every type-table entry**, including unused
  entries, in the same ascending fixed-point order as PureScript. Missing
  dependencies and cycles use the same ordered `Any` fallback.
- The lexical source-usage validation pass now runs before returning the TAST.
  It checks annotation placement, globally unique binding IDs, lexical scope,
  shadowing, recursive/nonrecursive lets, patterns and qualified variable uses.
  IDs inherit the enclosing module's provenance. The previous reference omitted
  this pass and resolved types on demand.
- Each combined decode constructs a fresh simdjson parser and destroys its DOM
  before returning the owned result. There is no parser reused across documents
  or iterations. Parse-only samples retain a separate parser-owned document for
  every input, then fingerprint those actual timed documents.
- Every timed decode returns a complete result, retained through its canonical
  fingerprint check. Result destruction is timed separately after that check.

simdjson's ordinary parser internals remain those of the installed library;
its views never become strings or nodes in a returned typed result. There is no
custom allocator in either production reference. The old copies under the
historical `json-diagnostic/audit/` tooling are not the active reference drivers.

The application reference still checks malformed inputs by rejection rather
than reproducing Argonaut's exact error strings. Those inputs are not timed.
The Go implementation preserves its full public decoder/error contract.
Consequently, these remain concrete workload comparisons between different
parsers, representations and runtimes, rather than a measurement solely of
compiler quality.

## Validation and ownership

Against the preserved PureScript/JavaScript decoder and canonical fingerprints:

| Corpus | Cases | Successful values | Rejections |
| --- | ---: | ---: | ---: |
| Application | 817 | 95 | 722 |
| TAST | 1,365 | 862 | 503 |

Both normal and AddressSanitizer/UndefinedBehaviorSanitizer builds pass. The
application cases include the 17 frozen inputs plus deterministic mutations.
The TAST cases include all 12 frozen modules, 600 scope/shadowing combinations,
700 type-table graphs, and directed annotation/number/let/pattern cases.

The ownership audit destroys the DOM and overwrites/releases the input text
before fingerprinting the result. It then detaches an owned user or executable
subtree, destroys its parent, and fingerprints the detached value again. TAST
types referenced from that subtree must survive the module and temporary table.
No sanitizer findings or differential mismatches occurred.

## Measurement protocol

Each campaign alternates Go/C++ process order across six pairs and exercises all
six parse/decode/combined phase orders. Each process takes two warmups and five
retained, fingerprint-checked samples per phase; cells are medians of the six
process minima. Go runs with `GOMAXPROCS=1`, `GOGC=100` and PGO disabled. Compiler,
library, generated-source, binary, corpus and oracle hashes are retained.

The combined interval includes parsing/indexing, conversion, copying, complete
typed-value construction, normal Go GC and C++ temporary/parser destruction.
Input loading, output-retention-buffer allocation and fingerprint computation
are outside that interval for both implementations. The separate parse/decode
controls have different allocation histories; their medians must not be added
to estimate combined time.

The first campaign, preserved in `initial-paired/`, observed:

| Workload | Go combined | C++ combined | Go/C++ |
| --- | ---: | ---: | ---: |
| JSON → TAST | 20.569188 ms | 10.915604 ms | 1.88× |
| Application JSON | 2.364417 ms | 0.666729 ms | 3.55× |

Its initial lifecycle audit used the earlier Go cleanup protocol. The final
audit additionally charges deferred decoder garbage before fingerprinting, as
described below; only that revised audit is used for final lifecycle claims.

## Final paired results

| Phase | TAST Go | TAST C++ | Application Go | Application C++ |
| --- | ---: | ---: | ---: | ---: |
| Parse control | 20.519396 ms | 4.361813 ms | 2.021500 ms | 0.377000 ms |
| Decode already-parsed JSON | 11.312584 ms | 6.590250 ms | 0.711459 ms | 0.278396 ms |
| Complete text-to-owned-value construction | **19.977959 ms** | **10.823125 ms** | **2.241417 ms** | **0.657688 ms** |
| Go/C++ construction ratio | **1.85×** | | **3.41×** | |

Both campaigns put TAST construction at roughly **1.85–1.88×** and application
construction at **3.41–3.55×** the owned C++ reference. All 12 paired observations
per workload retain the complete results and place C++ ahead of Go.

The README cells are refreshed to **19.98 ms / 2.24 ms** for Go and
**10,823.13 µs / 657.69 µs** for C++. The previous published C++ cells were
**8.758630 / 0.644960 ms**, and the previous Go cells **21.566500 / 2.629042 ms**
gave ratios of **2.46× / 4.08×**. Those older observations are historical, not a
paired isolation of arena removal: this rewrite also adds missing TAST work
and fresh-parser allocation, while the Go binary has not changed. The new Go
times are a measurement refresh, not another Go optimization gain.

Go allocation minima remain **24,405,672 / 19,065,736 / 29,708,136 B** for TAST
parse/decode/combined and **2,435,696 / 1,876,160 / 3,483,048 B** for application
JSON. Normal GC within those samples is included in elapsed time.

C++ separately reports `release_us` and the per-sample construction-plus-release
sum `lifecycle_us`. Medians of process minima for the combined sample's sum are
**12.901604 ms TAST / 0.713979 ms application**. These are not divided into the
Go construction cells; the fresh-process audit below is the aligned lifecycle
comparison.

## Lifecycle accounting

Six alternating pairs of **fresh processes** construct the public decoder, run
two complete decodes while retaining both output batches, then reclaim them.
The interval subtotal is computed within each process before taking the median.

The Go audit first clears input-loader scratch before constructing the decoder.
After both decodes, it **times a GC with all results still live**, before any
fingerprint allocation. This charges deferred parsing/decoding garbage rather
than hiding it in validation cleanup. It then fingerprints both batches, clears
only validation scratch outside timing while retaining the results, drops both
output roots and times final collection. C++ releases temporaries in its decode
intervals and times destruction of both retained result batches separately;
there is no decoder-getter construction or deferred garbage-collection phase.

Thus input loading, fingerprint computation, validation-scratch cleanup and
outer retention buffers are excluded, but decoder construction, temporary
reclamation, retention and final result reclamation are accounted for. These
are measured interval subtotals, not whole-process wall-clock times.

| Measured interval | TAST Go | TAST C++ | Application Go | Application C++ |
| --- | ---: | ---: | ---: | ---: |
| Public decoder/getter construction | 0.003209 ms | 0 | 0.069521 ms | 0 |
| First complete corpus decode | 23.660584 ms | 12.708355 ms | 2.542167 ms | 0.836667 ms |
| Second complete corpus decode | 28.931188 ms | 11.853458 ms | 2.657208 ms | 0.749208 ms |
| Deferred garbage, outputs still retained | 13.502396 ms | 0 | 1.274250 ms | 0 |
| Final result release | 1.577459 ms | 4.690770 ms | 0.380730 ms | 0.127208 ms |
| Per-process subtotal, then median | **67.757584 ms** | **29.217563 ms** | **6.953561 ms** | **1.701875 ms** |
| Go/C++ lifecycle ratio | **2.32×** | | **4.09×** | |

Subtotals are not sums of independently calculated phase medians. This audit
deliberately forces the Go retained-output and final-release collections so
deferred work is visible; its ratios describe that two-batch retention protocol,
not an amortized steady-state GC policy or a universal runtime ratio. The first
fresh-process pair is noisier for both runtimes and is retained in the archive.

Go heap medians with both output batches retained are **30,891,328 B TAST /
4,993,912 B application**; after final release they are **5,706,200 / 831,600 B**.
These are Go heap observations, not cross-runtime memory-footprint comparisons.

## Distribution

| Executable | Previous C++ | Owned C++ | Preserved Go |
| --- | ---: | ---: | ---: |
| Application | 76,152 B | 77,208 B | 7,001,506 B |
| TAST | 131,832 B | 189,752 B | 9,037,394 B |

C++ dynamically links simdjson **3.13.0** (a **112,992 B** shared library), libc++
and system libraries. The Go executable includes its runtime and generated
decoder; `distribution.json` records both binaries' dynamic dependencies.
The standalone executable sizes therefore have different dependency boundaries.
No Go compiler, library or generated production code changes are part of this
reference rewrite.

## Reproduction and provenance

The reproducible tooling is `bin/benchmark/json-diagnostic/owned-reference/`:

```sh
python3 bin/benchmark/json-diagnostic/owned-reference/run.py build \
  --cycle var/benchmark/json-reference-owned-20260925
python3 bin/benchmark/json-diagnostic/owned-reference/run.py validate \
  --cycle var/benchmark/json-reference-owned-20260925
python3 bin/benchmark/json-diagnostic/owned-reference/run.py measure \
  --cycle var/benchmark/json-reference-owned-20260925 --campaign paired
python3 bin/benchmark/json-diagnostic/owned-reference/finalize.py \
  --cycle var/benchmark/json-reference-owned-20260925
```

Build, validation and measurement destinations must be new: existing evidence is
not overwritten. A new cycle needs a `before.json` manifest pointing to the
preserved production builds and the matching `previous/<suite>/` binaries and
manifests. The existing cycle can be verified directly with `finalize.py`;
for another measurement use a new `--campaign` name. Snapshot compilation uses
`clang++ -O3 -std=c++17`; sanitizer executables are used only for validation.

`owned/` holds final build manifests, source snapshots and binaries;
`validation/` holds full differential and ownership evidence; `paired/` holds
all final process results and lifecycle observations. The corresponding
`initial-*` directories preserve the first campaign and earlier cleanup audit.
The official `json-diagnostic.py` driver builds the new owned references by
default and describes their updated validation/lifetime scope.

Across both campaigns, **6,120 sampled outputs** and **816 lifecycle outputs**
are retained and fingerprint-checked. Warmup and initial validation results are
additional to those counts. `final-verification.json` checks all preserved
manifests, binary/source hashes, differential files, per-process minima,
fingerprints, phase orders and lifecycle subtotals.

All three **excluded WIP** diagnostics remain outside benchmark totals and the
aggregate `/C` ratios.
