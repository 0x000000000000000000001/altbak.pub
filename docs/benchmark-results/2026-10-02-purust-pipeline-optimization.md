# Purust: native pipeline optimization campaign

## Status and scope

**In progress.** This third campaign targets the same frozen `purust-aff`
application as the [second campaign](2026-10-02-purust-codegen-optimization.md):
244 TAST modules, 138,448 type-table entries, 27,122,224 TAST bytes,
`--main Test.Main --threaded`. Exploratory results below are not the final
JS/native comparison.

Durable workspace: `var/benchmark/purust-pipeline-20261002/`.

## Protocol and provenance

The 548 reference files were restored and checked against the preceding
publication. TAST manifest SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

The installed native executable again differed from the published artifact.
Both were preserved; their identity is explicit in each comparison.

| Artifact at campaign start | SHA-256 |
|---|---|
| Published second-lot native compiler | `0a0f2afe52159a3c272beeaee625e68d100d5855bd3d99dc18ab6f353eca875e` |
| Installed native compiler | `0a45b38c92708a387a4ffeba6896e32af928a33fec8a0ecb40985de667d458ac` |
| JavaScript bundle, unchanged | `68922e52446516e0739f791c9a24df03ce8153a3c266e7d3a7f4b17c8c9c7735` |

Host: Apple M4 Pro, 14 cores, 48 GiB RAM, arm64 macOS 26.1, Node v24.8.0,
Rust/Cargo 1.96.0. Native compiler profile: **O3, LTO disabled**. Default
concurrency is **4 PBO + 4 generation workers**, with sequential TAST loading.

Exploratory comparisons run prebuilt artifacts sequentially, with one warm-up
per variant and three measured rounds rotating the first variant. Each run
uses a fresh process/output and empty `.purmeta`. Builds and tests finish before
measurement. The metric is monotonic **`backend total`**, including TAST loading
and sorting, preparation, PBO, generation, final worker drain and emission;
startup/exit, frontend `purs`, bootstrap, Cargo and application execution are
excluded. OS file caching is normal and CPU affinity is unset. Phase medians
are independent and do not necessarily sum to the total.

Candidate builds freeze both Purust and PBO source/FFI trees before generation.
Each build records source hashes. Times from distinct campaigns are not combined
to calculate speedups.

## Updated CPU and allocation diagnostics

The all-thread sample contains **16,715 nonblocked stacks with generated
PureScript frames**. Inclusive families include codegen 5,691, Map/Set 4,513,
allocation 3,852, decoding 1,552, type-table resolution 866, regex 767 and foreign
declarations 262. These overlapping wall-clock stack families are not additive
CPU percentages. Raw data: `cpu-baseline.sample.txt`, `cpu-summary.json`.

Per-thread allocation counters report:

| Phase | Allocation/reallocation requests | Requested bytes |
|---|---:|---:|
| Load | 119,029,955 | 4,841,515,960 |
| Prepare | 13,762,123 | 4,587,496,996 |
| Optimize + generate | 568,233,862 | 21,911,688,612 |
| Finalize | 86,652,810 | 2,712,609,161 |
| **Backend total** | **795,375,892** | **34,387,965,902** |

Requested bytes are cumulative, not live memory. There were 63 deferred PBO
attempts in this diagnostic, and all 496 generated files matched. Instrumented
times are excluded from performance comparisons. The profiling runner restored
its temporary source and executable changes byte for byte. Raw: `alloc-baseline.json`.

## Preparation experiment

This candidate reads each module's source/FFI once, reuses the combined Rust
string for two scans and caches fixed declaration/pub-use regexes with `OnceLock`.

| Variant | Three samples (ms) | Median total (ms) |
|---|---|---:|
| Published reference | 7,408 / 7,572 / 8,304 | **7,572** |
| Installed reference | 7,904 / 7,573 / 8,171 | **7,904** |
| Preparation candidate | 7,493 / 7,948 / 7,748 | **7,748** |

Preparation falls from 700 to 640 ms and maximum RSS from approximately 503 to
474 MiB, but total time is 2.3% higher than the published control. **No global
speedup is established for this candidate alone.** All 12 outputs contain 496
identical files. Raw: `preparation.json`.

## Compact transitive dependency closure

The module graph is canonically indexed with integers, including referenced
external names as leaves. Rust computes nonreflexive reachability with a dense
bitset matrix; cycles can introduce self-reachability just as before. Invalid
indices and graphs larger than 8,192 vertices use the PureScript fallback.
Results are converted back to the original ordered Map/Set representation.

| Variant | Three samples (ms) | Median total (ms) | Median finalize (ms) |
|---|---|---:|---:|
| Published reference | 7,883 / 7,975 / 8,447 | **7,975** | 1,467 |
| Preparation only | 7,748 / 7,973 / 9,280 | **7,973** | 1,423 |
| Preparation + graph | 7,118 / 6,841 / 7,108 | **7,108** | 747 |

The combined candidate reduces the total by **10.9%** against the published
control. The preparation-only result again shows no clear total improvement.
All **12 outputs × 496 files** match exactly. Raw: `graph.json`.

Differential tests pass **422 named PureScript graphs / 417 native graphs**:
independent reachability, cycles, self-edges, external leaves, Unicode, random
graphs, 64/128-bit boundaries and invalid/oversized fallback.

## Specialized type-table resolution

The Rust path parses reference forms once, settles pending indices in ascending
order and retains shared ownership of resolved types. It handles well-formed
acyclic tables. Errors, missing references and cycles delegate the whole original
input to `decodeTypeTablePS`, preserving error priority and cycle forcing.
Text parsing and module/annotation decoding still use the existing implementation.

The real-runtime differential harness passes **830 tables / 151,339 entries**,
including 648 fast paths, all constructors, forward/backward references, UTF-16
strings, errors/cycles, randomized DAGs and the 244 frozen application tables.
Error text and shared type references are checked. Raw: `type-table-retry.log`.

The first harness build was interrupted by `ENOSPC`; completed older Cargo
caches were removed after preserving their executables, sources and logs.
The retry passed. The failed run supplies no benchmark sample.
