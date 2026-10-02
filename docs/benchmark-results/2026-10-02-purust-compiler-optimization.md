# Purust: first native compiler optimization campaign

## Scope and retained changes

This campaign targets compilation of **`purust-aff`**, using the exact application
inputs from the [2026-10-01 benchmark](2026-10-01-purust-aff-compilation.md):
244 TAST modules, 138,448 type-table entries, 27,122,224 TAST bytes, and
`--main Test.Main --threaded`.

The [final five-pair benchmark](2026-10-02-purust-aff-compilation.md) measures
**6,774 ms JavaScript / 10,328 ms native**, a native/JS ratio of **1.52×**.
All generated files match, and the generated application passes its 47 Aff checks.

Retained changes:

1. Build the native compiler at **`opt-level=3`**, with **LTO disabled**.
2. Optimize the Rust source scanner: recognize raw-string prefixes directly,
   cache the character regex, transform ownership without per-span regex
   compilation, and discover imports without constructing an unused output string.
3. Implement PBO's native Map operations on the existing persistent AVL layout.
   String, Int and Qualified Ident keys use borrowed native comparisons.
   EvalRef/TcoRef operations retain their supplied comparator, avoiding a crate
   dependency cycle, while removing temporary Ord/Eq dictionaries and generic
   traversal wrappers.
4. Connect the existing parallel PBO builder to an Aff scheduler with bounded
   work, supervised worker lifetime, rank-filtered visibility, and ordered,
   sequential Rust generation. The native default is **at most 8 workers**,
   capped by available CPUs; `PURUST_PBO_JOBS=1` selects the sequential builder.
   The JavaScript default remains 1.

The detailed [measurement archive](2026-10-02-purust-compiler-optimization.json)
includes every exploratory sample, binary hashes, peak RSS, PBO diagnostics,
allocation counts, and instrumentation sources. These exploratory comparisons
use **three measured rounds**; the final JS/native comparison uses five pairs.

## Inputs, host and protocol

The old temporary benchmark directory had disappeared. All **548 frozen files**
were restored from the checkouts and verified against the published manifest,
including both original compiler executables. They are retained under
`altbak.pub/var/benchmark/purust-compiler-20261002/reference/`.
The input TAST manifest hash remains
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

Host: Apple M4 Pro, 14 CPU cores, 48 GiB RAM, arm64 macOS 26.1; Node v24.8.0,
Rust/Cargo 1.96.0. No CPU affinity or OS cache flush. Each exploratory campaign
uses one warm-up per variant, then three sequential rounds, rotating the first
variant. Every invocation gets a fresh output directory and an empty `.purmeta`.
Compiler rebuilds and tests finish outside the uninstrumented timing campaigns.

The metric is the compiler's monotonic **`backend total`**, including TAST
loading/sorting, preparation, PBO optimization, Rust generation and emission.
It excludes process startup/exit, the `purs` frontend, bootstrap, Cargo and
application execution. `/usr/bin/time -l` additionally records process resource
usage. Phase medians are independent and need not sum to the total median.

Each row below belongs to its own interleaved campaign. Changes in the control
times between campaigns are retained; speedups should be calculated against
the control measured in the same campaign.

## Compiler optimization level

The same generated compiler Rust sources were compiled at levels 1, 2 and 3,
always with LTO disabled.

| Compiler level | Three samples (ms) | Median (ms) |
|---|---|---:|
| 1 | 19,196 / 17,657 / 18,509 | **18,509** |
| 2 | 18,061 / 17,394 / 17,030 | **17,394** |
| 3 | 16,510 / 15,707 / 16,063 | **16,063** |

Level 3 reduces the median by **13.2%** relative to level 1. All 12 invocations,
including warm-ups, produced 496 identical Rust sources/manifests.
`PURUST_NATIVE_OPT_LEVEL=1|2|3` overrides the new bootstrap default of 3;
`PURUST_NATIVE_OUTPUT` selects a separate executable destination for experiments.

## Final controlled old/new comparison

After the final JS/native campaign, the original native artifact and the
installed stage-2 artifact were compared directly on the restored input snapshot.
One warm-up per artifact was followed by **five measured rounds**, alternating
which compiler runs first. All 12 outputs contain the same 496 files.

| Native compiler | Five samples (ms) | Median (ms) | Maximum RSS (MiB) |
|---|---|---:|---:|
| Original artifact | 17,074 / 17,111 / 17,110 / 18,715 / 17,899 | **17,111** | 312 |
| Optimized stage 2 | 9,928 / 9,975 / 9,935 / 9,880 / 10,004 | **9,935** | 458 |

The optimized compiler is **1.72× faster**, taking **41.9% less time**, with
higher peak memory. The old/new figures belong to this separate controlled
campaign. The README uses the final paired JS/native campaign's **10,328 ms**
native median; it does not substitute the lower median from this comparison.

## Scanner and native Maps

All variants in this comparison use level 3.

| Variant | Three samples (ms) | Median (ms) |
|---|---|---:|
| Control | 15,124 / 15,536 / 15,475 | **15,475** |
| Scanner | 14,939 / 15,199 / 14,963 | **14,963** |
| Scanner + Maps | 14,662 / 14,589 / 14,489 | **14,589** |

The combined reduction is **5.7%**. All 12 invocations produced 496 identical
files. Finalization/emission fell from a median of 2,056 to 1,657 ms; optimization
plus generation fell from 10,609 to 10,134 ms.

The native scanner passes **614 differential cases against JavaScript**, plus
concurrent calls. The tests include truncated tokens, escaped and raw strings,
nested comments, Unicode, character literals and lifetimes. They also exposed
and fixed a pre-existing ASCII/Unicode word-boundary difference.

Native Map tests compare against the generated PureScript implementation:
insertion, lookup, union, left bias, combining-function argument order, exact AVL
shape/heights/sizes, retained old versions, callback comparators, and Unicode
keys including isolated UTF-16 surrogates.

## Parallel PBO

An initial two-worker run produced four different Rust files. The parallel
builder was missing the sequential builder's forced inlining for untyped private
bindings. The fix applies the rule to the prepared current module and visible
finalized predecessors. A targeted test fails before the fix and passes at
2/4/8 workers afterward. The incorrect variant is excluded from performance
conclusions; its failure log is retained in the archive.

With the correction, all **16 invocations** produced the same 496 files.

| PBO workers | Three total samples (ms) | Median total (ms) | Median optimize + generate (ms) | Maximum RSS (MiB) |
|---|---|---:|---:|---:|
| 1 | 14,905 / 14,412 / 14,623 | **14,623** | 10,129 | 310 |
| 2 | 12,175 / 12,476 / 12,037 | **12,175** | 7,698 | 468 |
| 4 | 11,441 / 11,498 / 11,370 | **11,441** | 7,003 | 455 |
| 8 | 10,506 / 10,606 / 10,608 | **10,606** | 6,146 | 463 |

Eight workers reduce total time by **27.5%** relative to the optimized sequential
compiler. There are **113–121 deferred attempts** across the measured eight-worker
runs. `dispatched + fallback` counts all attempts, including retries; neither
counter alone is the module count. Sequential codegen callbacks still account
for about **5.4 seconds**. That callback duration includes their FFI reads and
module preparation for emission, and overlaps worker optimization.

The scheduler tests exercise out-of-order completion, rank visibility, directive
accumulation, private inlining, preparation/codegen errors, and completion of
worker finalizers before propagating an error.

## Parallel loading experiment

PBO is fixed at 8 workers; `PURUST_JOBS` varies TAST loading concurrency.
All 16 invocations again produce identical files.

| Loading workers | Median load (ms) | Median total (ms) | Maximum RSS (MiB) |
|---|---:|---:|---:|
| 1 | 1,963 | 10,913 | 463 |
| 2 | 1,854 | 10,509 | 472 |
| 4 | 1,854 | 10,493 | 508 |
| 8 | 1,828 | 10,512 | 535 |

The gain is small and memory rises. The loading default remains **1**. This
experiment does not establish parallel scaling of the JSON decoder itself.

## CPU and allocation diagnostics

`sample <pid> 20 2` collected 7,858 stack samples on the thread executing the
baseline backend. Inclusive families include PBO (3,686), codegen (2,926),
allocation (1,949), Map/Set (1,739), String cloning (934), decoding (842), Value
cloning (802), and the scanner (545). These are overlapping sampled stacks,
**not additive CPU percentages or allocation counts**.

A separate global-allocator wrapper counts allocation/reallocation requests and
requested bytes. The allocator remains mimalloc. The counters are sampled at
the existing phase boundaries. Instrumented times are excluded from benchmarks.
Requested bytes are cumulative traffic, not live or peak memory.

| Sequential phase | Baseline requests | Scanner + Maps requests | Baseline requested bytes | Scanner + Maps requested bytes |
|---|---:|---:|---:|---:|
| Load | 119,029,955 | 119,029,955 | 4,841,515,960 | 4,841,515,960 |
| Prepare | 14,072,002 | 14,072,002 | 4,593,931,305 | 4,593,931,305 |
| Optimize + generate | 593,221,204 | 568,846,931 | 21,483,814,747 | 20,684,676,139 |
| Finalize | 114,841,308 | 104,694,804 | 3,736,407,669 | 3,021,150,991 |
| **Backend total** | **848,879,832** | **814,359,060** | **34,990,644,720** | **33,476,249,526** |

The total includes work between phase boundaries. Scanner and Maps remove
**34,520,772 requests** and **1,514,395,194 requested bytes**.

A second diagnostic compares the corrected parallel compiler at 1 and 8 PBO
workers, using padded per-thread counters to reduce instrumentation contention:

| PBO workers | Total requests | Total requested bytes |
|---|---:|---:|
| 1 | 814,359,065 | 33,476,287,250 |
| 8 | 942,478,653 | 37,853,052,893 |

The eight-worker diagnostic has 116 deferred attempts: **15.7% more requests**
and **13.1% more requested bytes**. These counts describe the recorded
instrumented schedule; retries vary with scheduling. Both instrumented
comparisons preserve the 496 generated files exactly.

## Self-reconstruction and regression checks

The optimized compiler was rebuilt through the complete Node → native stage 1 →
native stage 2 chain. The compiler input contains **451 TAST modules and 281,964
type-table entries**. With 8 PBO workers, stage 1 produced **910 Rust sources and
Cargo manifests byte-identical to JavaScript**. Cargo compiled stage 2 at level 3
with LTO disabled. Stage 2 then passed the fresh-project smoke test: 152 modules,
312 identical files, and executable result **`PURUST_NATIVE_OK 42`**. That stage 2
was atomically installed as `bin/purust-native`.

The codegen checks comprise **86 tests**, including the scanner and scheduler
regressions. Four initially failed because their Docker test container was
stopped; all four passed when rerun after starting it. The **43 TAST tests** also
pass after updating the list-pipeline assertion: it had required the recursive
closure self-capture removed by the earlier compiler fix. It now verifies a
direct worker call without that capture, followed by compilation and execution
of the pipeline fixtures.

The full `purust-aff/bin/test` suite passes using the installed native compiler:
47 Aff checks, 5 Rust unit tests, concurrent timer/Ref/AVar checks, child lifetime
checks, and 9 error-reporting scenarios. The archive retains the full log.

## Next measured targets

After parallel PBO, sequential codegen is the dominant remaining phase. The
sampled allocation stacks point to `sanitizeIdent`,
`codegenExprTypeWithValueEnums`, `boxUnbox`, `joinWith` and generic Map operations.
Priorities are native string builders and redundant generated allocations,
followed by isolating codegen's mutable state before parallel generation, then
specialized TAST decoding. Native array indexing already uses `array_get`
without converting the entire array.

Purust's `todo.md` records the implementation plan, completed steps, experiments
and remaining work.
