# Purust Aff compilation after the pipeline optimization lot

## Published result

Measured on **2026-10-02**, using the self-rebuilt, installed stage-2 compiler:

| Backend | Median `backend total` | Range across five runs |
|---|---:|---:|
| `purust.js`, sequential PBO/codegen | **6,292 ms** | 6,239–6,303 ms |
| `purust-native`, 4 PBO + 4 codegen workers | **5,743 ms** | 5,178–5,796 ms |

**Native now takes 0.91× the JavaScript time: 8.7% less time on this workload.**
This reaches the first target of beating JS on the frozen `purust-aff` corpus.
The preceding [codegen campaign](2026-10-02-purust-aff-codegen-compilation.md)
published a native/JS ratio of 1.28×. The longer-term 0.3× target is not reached.

A separate controlled comparison against the preceding published native
compiler gives **7,583 → 5,499 ms**, a **27.5% reduction / 1.38× speedup** for
this third lot, with maximum RSS **504 → 485 MiB**. See the
[optimization report](2026-10-02-purust-pipeline-optimization.md#final-controlled-beforeafter-comparison).
These are distinct samples; the README uses **5,743 ms** from the paired campaign.

| Pair, in execution order | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Warm-up, excluded | 6,234 | 5,738 |
| 1 | 6,239 | 5,752 |
| 2 | 6,303 | 5,178 |
| 3 | 6,248 | 5,796 |
| 4 | 6,302 | 5,743 |
| 5 | 6,292 | 5,676 |

All five measured native runs are faster than their JS partners. The second
native sample is noticeably faster than the others; **all samples are retained**
and the statistic remains the median. The one-minute load average declines from
roughly 8 at warm-up to 5 by the last pair. Compiler builds and regression tests
completed before timing.

The [raw measurements](2026-10-02-purust-aff-pipeline-compilation.json) contain
every sample, phase, stdout/stderr log, compiler/input hash, checkout state and
application-validation result. The [pipeline campaign archive](2026-10-02-purust-pipeline-optimization.json)
includes exploratory comparisons, the contended type-table run, controlled
before/after measurements, allocation diagnostics, source hashes and regression
and self-hosting evidence.

## Inputs, configuration and protocol

- Same **244 TAST modules, 138,448 type-table entries and 27,122,224 TAST bytes**
  as the earlier reports, `builtWith=0.15.16`. All **544 application input files**
  were checked against the frozen reference before the final snapshot. The
  snapshot has 548 files including compiler artifacts and comparison tooling.
- Flags: `--source output --main Test.Main --threaded --out <fresh-directory>`.
  Both backends process the complete module set.
- Native compiler: **O3, LTO disabled**, with compact dependency closure,
  specialized acyclic type-table resolution, borrowed layout-fact lookups and
  preparation reuse, on top of the previous scanner/Map/sanitizer/PBO/codegen
  changes. The eight-slot budget remains **4 PBO + 4 codegen workers**. TAST
  loading concurrency is 1.
- JavaScript uses the same compiler/PBO source tree, sequential PBO/codegen and
  the PureScript fallbacks. Node flags: `--expose-gc --stack-size=65536
  --max-old-space-size=16384`.
- Freeze the launcher, both compilers, TAST, PureScript sources, adjacent Rust
  FFI/Cargo declarations and expected application stdout.
- One warm-up per backend, then **five sequential JS→native pairs**. Fresh
  process, output directory and `.purmeta` state per invocation. Normal OS file
  caching, no CPU affinity or cache flush. Retain every measured sample.
- Metric: monotonic **`backend total`**, covering TAST loading/sorting,
  preparation, PBO, Rust generation and emission, including the final codegen
  drain. Excludes process startup/exit, frontend `purs`, compiler rebuilding,
  Cargo and application execution.

Host: **Apple M4 Pro**, 14 CPU cores, **48 GiB RAM**, arm64 macOS **26.1 (25B78)**;
Node **v24.8.0**, Rust/Cargo **1.96.0**. Timing ran from **14:38:42 to 14:39:56 UTC**;
application validation finished at **14:40:13 UTC**.

### Phase medians

| Phase | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Load TAST + sort | 586 | 817 |
| Prepare | 38 | 608 |
| Optimize + generate | 4,795 | 3,465 |
| Finalize + emit | 815 | 729 |

Each phase has its own median; they need not sum to the total median. Native
optimization/generation and finalization are faster in this campaign, while
loading and especially preparation still take more time than JS.

## Validation and provenance

All **12 invocations** succeeded and produced the same **496 Rust source files
and Cargo manifests**, compared byte for byte. After measurement, the last
native output built with `cargo build --offline` and passed **47 Aff checks**:
exact expected stdout, empty stderr, exit 0. Frozen hashes were checked after
measurement and again after application execution.

The subsequent archive verification rechecked all artifact/input hashes and
medians, the 544 application inputs and **76 measured/warm-up outputs, each with
496 identical files**, plus both allocation-diagnostic outputs. It also checked
all 914 self-hosted files across the original JS, final JS and native stage-1
outputs, and confirmed that the measured executable is the installed one.

The measured binary is the installed stage 2. Its compiler corpus contains
**453 TAST modules / 282,610 types**. Initial JS, final JS and native stage 1
produced **914 identical Rust sources/manifests**, then Cargo built stage 2.
Stage 2 passed the fresh **152-module / 312-file** smoke test with
`PURUST_NATIVE_OK 42` and the full Aff suite: **47 checks, 5 Rust unit tests,
concurrency/Ref/AVar/lifetime checks and 9 error-reporting scenarios**.

Regression checks include **94 codegen tests**, **11 targeted TAST tests** and
**3 final targeted checks**. The native specializations additionally pass
dependency-graph differential tests, **830 type tables / 151,339 entries** and
**12,996 layout lookups**, including UTF-16, persistent versions and concurrent
readers. Failed disk-limited builds and Docker-dependent test retries are
documented in the optimization report and retained in the raw archive.

| Compiler artifact | Bytes | SHA-256 |
|---|---:|---|
| JavaScript | 2,204,875 | `969bf49f5da903506d6821caff0342f0f50161c48156a41078cbd1b9721c0940` |
| Native stage 2 | 42,455,488 | `96cf27f08bb888a0677d1cd267ded9c52a5537301587ceb1715572e2018bea26` |

Snapshot revisions: Purust `7698d87d0ea0c475569f6e06ee46b89707944578`, Aff
`823ebc23d529af9a124f382431ab7a2eee893b93`, PBO
`0f41544464ec0f42e6cb0dd77b206852813f904f`. Working-tree changes are recorded in
the JSON, including the final DataLayout FFI adjustment and PBO's modified
builder/untracked Rust ports. Artifact and source hashes identify the actual
compiler contents used.

TAST manifest SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

## Rerunning

From `altbak.pub`, with installed compilers and existing Aff TAST:

```sh
TMPDIR=/path/to/durable/workspaces \
  node bin/benchmark/compilation-purust-aff.mjs \
  ../purust/purust ../purust/purust-aff /path/to/new-results.json
```

Use a new result filename. This campaign's retained workspace is
`var/benchmark/purust-pipeline-20261002/purust-aff-compilation-gVB0bN`.
