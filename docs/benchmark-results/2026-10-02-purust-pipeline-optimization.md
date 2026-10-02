# Purust: native pipeline optimization campaign

## Result and scope

This third campaign targets the same frozen `purust-aff`
application as the [second campaign](2026-10-02-purust-codegen-optimization.md):
244 TAST modules, 138,448 type-table entries, 27,122,224 TAST bytes,
`--main Test.Main --threaded`.

The qualified, self-rebuilt stage 2 is installed. The
[final five-pair comparison](2026-10-02-purust-aff-pipeline-compilation.md) gives
**6,292 ms JS / 5,743 ms native**, ratio **0.91×**: native takes **8.7% less time**
on this workload. The first goal of beating JS on `purust-aff` is reached.
A separate native before/after comparison gives **7,583 → 5,499 ms**, or
**27.5% less time / 1.38× speedup**, with maximum RSS **504 → 485 MiB**.

The [raw archive](2026-10-02-purust-pipeline-optimization.json) preserves every
comparison, diagnostics, source/artifact hashes, build failures and qualification
logs. The final JS/native samples and exploratory samples remain distinct.

Durable workspace: `var/benchmark/purust-pipeline-20261002/`.

Retained changes:

1. Read source/FFI contents once per module, reuse combined scan input and cache
   fixed foreign-declaration regexes. Standalone evidence is lower memory and
   allocation work, not a demonstrated total-time speedup.
2. Compute transitive module dependencies with canonical integer indexing and
   native bitsets, retaining the PureScript fallback and ordered result.
3. Resolve valid acyclic type tables directly in Rust with shared type ownership;
   delegate errors, missing references and cycles to the original decoder.
4. Borrow layout-fact AVL keys during native lookup instead of allocating
   temporary pair/comparison objects and cloning keys at each node.

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

## Final controlled before/after comparison

After the final JS/native campaign and application validation, the preceding
published second-lot native compiler and the installed third-lot stage 2 were
compared on the restored input snapshot. One warm-up each, then **five measured
rounds**, alternating which compiler runs first. All 12 outputs have the same
496 Rust sources/Cargo manifests. Raw: `before-after.json`.

| Native compiler | Five samples (ms) | Median total (ms) | Maximum RSS (MiB) |
|---|---|---:|---:|
| Published second-lot stage 2 | 7,696 / 7,591 / 7,583 / 7,512 / 7,339 | **7,583** | 504 |
| Third-lot stage 2 | 5,183 / 5,730 / 5,715 / 5,178 / 5,499 | **5,499** | 485 |

This lot reduces median time by **27.5%**, a **1.38× speedup**, with lower maximum
RSS. All samples are retained, including the two faster third-lot runs. These
measurements isolate the third lot against its predecessor. The README retains
the separate paired JS/native campaign's native median, **5,743 ms**.

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

The cumulative preparation/graph/type-table/layout candidate compiles and gives:

| Phase | Candidate allocation/reallocation requests | Candidate requested bytes |
|---|---:|---:|
| Load | 53,755,455 | 2,224,653,112 |
| Prepare | 13,163,439 | 4,014,332,668 |
| Optimize + generate | 540,509,978 | 20,790,965,657 |
| Finalize | 44,916,908 | 1,678,154,700 |
| **Backend total** | **660,042,927** | **29,042,761,492** |

This is **17.0% fewer requests / 15.5% fewer requested bytes**. Deferred PBO
attempts were 66, versus 63 in the baseline diagnostic, so optimizer work is
still scheduling-dependent. All 496 generated files match; instrumented source
and executable changes were restored exactly. Raw: `alloc-layout.json`.

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

The first full-compiler comparison produced:

| Variant | Three samples (ms) | Median total (ms) | Median load (ms) |
|---|---|---:|---:|
| Published reference | 14,497 / 16,179 / 14,838 | **14,838** | 3,911 |
| Preparation + graph | 14,438 / 12,377 / 11,355 | **12,377** | 3,709 |
| Preparation + graph + tables | 8,837 / 12,930 / 8,072 | **8,837** | 1,491 |

All 12 outputs contain 496 identical files. However, the one-minute load average
was **33–40**, compared with roughly 9–10 in the preceding graph comparison,
and times were strongly dispersed. Another gopurs compilation campaign was
observed immediately afterwards. These measurements are retained as a contended
exploratory run; an uncontended confirmation is required before claiming a final
gain. Raw: `types.json`, `types-contention-processes.txt`.

## Native layout-fact lookup candidate

Type representation queries use a persistent Set of pairs of strings. The Rust
lookup borrows existing AVL keys directly, avoiding temporary Tuple/Ordering
objects, generic comparator calls and repeated key copies. Module-name
normalization and UTF-16 string ordering remain unchanged; JavaScript uses the
PureScript reference.

The real-runtime differential harness passes **12,996 lookups**, including
retained persistent versions and Unicode/surrogate boundaries, plus concurrent
lookups through eight readers. Raw: `data-layout-tests.log`.

The first integrated build caught a Set newtype ABI mismatch at the FFI
boundary. The PureScript signature now exposes the underlying Map explicitly,
with `Set.toMap` at the call site; the native lookup algorithm is unchanged.
Failed-attempt sources and logs are retained in `build-layout/failed-set-ffi/`
and `build-layout/cargo-build.log`.

## Four-variant confirmation and retained candidate

After concurrent compilation activity subsided, all four prebuilt compilers
were compared again. Raw: `layout.json`; process/load preflight:
`layout-preflight.txt`.

| Variant | Three samples (ms) | Median total (ms) | Load (ms) | Optimize + generate (ms) | Finalize (ms) | Max RSS (MiB) |
|---|---|---:|---:|---:|---:|---:|
| Published reference | 8,011 / 7,630 / 7,693 | **7,693** | 1,678 | 3,856 | 1,359 | 498 |
| Preparation + graph | 7,268 / 6,904 / 6,840 | **6,904** | 1,678 | 3,759 | 731 | 465 |
| + Native type tables | 6,394 / 6,028 / 6,039 | **6,039** | 814 | 3,783 | 716 | 486 |
| + Native layout lookup | 5,771 / 5,789 / 5,716 | **5,771** | 825 | 3,450 | 720 | 483 |

All **16 outputs × 496 files** match exactly. The cumulative candidate reduces
total time by **25.0%** against the published control. Native tables improve
their graph-only control by **12.5%**; native layout lookup improves the
type-table control by a further **4.4%**. These comparisons belong to the same
campaign and confirm the type-table result outside the earlier heavy contention.

The complete candidate passes self-hosting qualification and is installed. Preparation
caching/regex reuse is retained for its lower observed memory footprint and
allocation work, with **no standalone total-time speedup claimed**.

The corrected layout build was interrupted again by disk exhaustion. Completed
compiler caches were reclaimed with executables, source trees and logs retained;
the Cargo retry and allocation diagnostic then succeeded. Qualification reuses
the measured Node-built stage 1, first checking it against regeneration with the
final JS bundle, and frees its completed cache before building stage 2.

## Regression checks

The codegen suite passes **94 tests**: 90 in the first run and four after
restarting the Docker service required by their fixtures. Both the original
failures and successful retry are retained. **Eleven TAST regressions** covering
representations, FFI files and Cargo manifests also pass.
After correcting the FFI boundary, the final JS bundle rebuilt successfully and
**three targeted tests** for value enums, dependency closure and foreign types
passed.

## Self-reconstruction and full Aff qualification

The frozen compiler corpus contains **453 modules / 282,610 types**. The
measured Node-built stage 1 was reused. The final JS bundle regenerated the
same compiler sources, then stage 1 regenerated them again: **914 Rust source
files/Cargo manifests identical across all three outputs**.

After its completed Cargo cache was freed, the separate stage-2 workspace built
at **O3 without LTO**. Stage 2 passed a fresh-project smoke test: **152 modules,
312 identical files and `PURUST_NATIVE_OK 42`**. Source hashes were checked
against the frozen candidate before and after qualification. Stage 2 was then
installed atomically as the default native compiler.

Installed stage-2 SHA-256:
`96cf27f08bb888a0677d1cd267ded9c52a5537301587ceb1715572e2018bea26`.
Final JavaScript SHA-256:
`969bf49f5da903506d6821caff0342f0f50161c48156a41078cbd1b9721c0940`.

The installed compiler passes the complete Aff suite: **47 Aff checks, 5 Rust
unit tests, concurrent timers/Ref/AVar, child lifetime checks and 9 error-reporting
scenarios**. Logs: `self-host.log`, `build-layout/stage2-smoke.log` and
`aff-final-tests.log`. Self-compilation times are validation observations and
are not used to claim a `purust-aff` compilation speedup.

## Final archive verification

The archive runner independently rechecks every campaign's input/artifact
hashes, samples, phase medians and maximum RSS, then compares all generated
outputs. Verification passes for **548 reference files, 544 unchanged application
inputs, 76 benchmark outputs × 496 identical files**, plus two allocation outputs.
It also rechecks **914 identical self-hosted files** across original JS, final JS
and native stage 1, and confirms that the installed stage 2 is the measured one.
Raw verification: `verification.json` in the durable workspace and the
`verification` object in the published campaign archive.
