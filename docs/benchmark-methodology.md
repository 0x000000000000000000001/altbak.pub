# How the benchmarks are measured

These small tests help identify compiler and runtime bottlenecks. They do not
predict the speed of a complete application. For the table labels and ratios,
see the [README legend](../README.md).

## What is being compared?

| Columns | Implementation |
| --- | --- |
| Compiled PureScript | The same PureScript program, compiled for each target. |
| Functional references and FP-style FFI | Hand-written translations preserving the functional structure. |
| Imperative references and FFI | Optimized implementations that may use different algorithms and mutable data. |

A ratio therefore measures the whole implementation: compiler, libraries,
runtime and source choices. It does not isolate compiler quality.

The current C column and `/C` ratios use the
[optimized imperative C reference](../tmp/bench_native.c), following its
[input and timing contract](../tmp/imperative_reference_contract.md). It may
use mutable trees, arrays and fused loops. Working storage is allocated and
released inside each timed action; no arena is preallocated outside timing.
The earlier [imperative measurements](benchmark-results/2026-11-24-c-reference.md)
and [functional C comparison](benchmark-results/2026-11-25-go-c-js-functional.md)
remain historical observations. Excl WIP diagnostics retain their own references.

All tables require opaque runtime inputs: a compiler must not replace the
driver's input read with its known literal or reuse an entire call's result
across a batch. Validation uses additional runtime inputs and checks optimized
output where available. A source-level identity function alone does not prove
an optimization barrier. Historical columns retain the evidence and limitations
of their original campaigns; they are not automatically certified by a new audit.

Algebraic optimization of the runtime input remains valid: multiplying an
unknown repetition count by 1000 differs from returning a precomputed answer
for the driver's chosen count. Source constants and input-independent work can
still be simplified. RowToList explicitly ignores its opaque integer and counts
a fixed record's fields, so its result need not vary with that integer.

These are comparisons of complete optimized implementations. Their allocations,
garbage collection and call adapters can differ, and can dominate very short
actions. A small time does not measure the cost of operations eliminated by
the compiler.

## How is a time obtained?

For the 14 core tests and 23 shared library microbenchmarks:

1. Build first, then freeze the sources, executable and runtime settings.
2. Warm up the program and calibrate repeated calls into timed batches.
3. Check the calculated results against independent expected values.
4. Keep the fastest of ten batches per process, then publish the median from
   three independent processes.

Compilation, process startup, logging and conversion of the final result to
text are outside the core timing interval. The computation runs on every call;
its result must be consumed. Generated call adapters remain included where used.

These are warmed-up minimum timings, not average request latency. Totals add
per-test medians. Small differences, especially on very short tests, may be noise.

## Which tests have a separate protocol?

- **JSON diagnostics** measure parsing, decoding and their combination separately.
  File reads and result verification are outside timing. See the
  [JSON report](benchmark-results/2026-11-24-go-json-decoding.md).
- **Array Indexing** measures several array sizes and reports time and allocations
  per access. See the [protocol](array-indexing-benchmark.md).
- **Extended tests** include their Effect/Aff adapters. File I/O uses the OS cache
  without an `fsync` guarantee. The async delay tests the timer; Parallelism
  compares JS's single-thread scheduler with Go/Rust's multicore runtimes.

The three diagnostics marked **Excl WIP** never enter the totals. The 23 shared
**Incl WIP** tests have their own [suite documentation](library-benchmarks.md).
Compare campaigns only when their inputs, timing boundaries and runtime settings
match; each report records those settings.

## Precision and evidence

Values ≥ 1 display two decimal places. Below 1, leading fractional zeros do not
count toward the two significant digits: `0.009492 → 0.0095`.
Calculations use archived timings before rounding. Missing values and unresolved
bounds such as `< 0.01` are excluded from ratio summaries.

The [campaign index](benchmark-results/README.md) helps find a report. The
[timing catalogue](../bin/benchmark/ratio-measurements.json) links each
measurement to its source. Campaign archives retain process timings, checked
outputs, source and executable hashes, commands and runtime settings.

## Reproducing a campaign

Build the chosen backend with `--build-only`, then measure its frozen executable
with `--run-only`. For the standard collector, a plan specifies those artifacts
and run commands:

```sh
python3 bin/benchmark/measure.py --plan path/to/plan.json --output path/to/new-campaign
```

Use the relevant campaign report for the exact commands and environment.
`./bin/run` alone does not reproduce every published reference or aggregate
three processes. Runner-specific details live with the
[collector](../bin/benchmark/measure.py), [Java runner](../bin/java/run) and
[PHP dependency preparation](../bin/php/driver.py).

For the 14-case Go/C-imperative/official-JS comparison, freeze the three builds
first, then use `bin/benchmark/compare-core.py --c-kind imperative`. It checks
their source and payload fingerprints, runs each backend in three separate processes in rotating
order, validates the 14 oracles and batch counts, and writes unrounded timings
plus the two descending ratio rankings. JavaScript is remeasured in that
campaign rather than taken from an older catalogue. The optional
`--c-kind functional` preserves the separate functional-reference protocol.

The experimental Fable Rust route also requires the adaptations described in
[its technical reference](../tmp/fable_rust/README.md).
