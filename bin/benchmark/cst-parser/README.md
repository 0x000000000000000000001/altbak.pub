# CST parsing application benchmark

The [recorded campaign](results.json), measured on 22 November 2026, compares
`PureScript.CST.parseModule` on a nine-module tidy source corpus and the
historical tidy fixture. A matched parse/format/consume control uses the same
compiled `Main` and dependencies as the parsing measurements. The
[application tables](../../../README.md#application-benchmark-results) report
the runtime and allocation results.

## Workloads and timing

The corpus is the sorted set of `gopurs-tidy/src/**/*.purs`: nine modules,
103,257 bytes, including the local `Tidy.Printer`. The fixture is
`gopurs-tidy/bench/fixture.purs`, a 514-line copy of `Tidy.Doc`, with SHA-256
`8faa34f297eefb2ebb7aaedde6d71e23d7766963237a43a9af0d77c36ee9ac0b`.

| Workload | Passes per process | JS per pass | Go per pass | Go / JS |
| --- | ---: | ---: | ---: | ---: |
| Nine-module corpus, parse only | 12 | 63.05 ms | 387.77 ms | 6.15x |
| Tidy fixture, parse only | 76 | 9.40 ms | 44.31 ms | 4.71x |
| Same fixture, parse + format + consume | 76 | 15.03 ms | 57.59 ms | 3.83x |

Times are medians of the in-process loop divided by its pass count. Native
JS and Go drivers read source files before starting their timers and invoke
the generated functions on every pass. Process startup, module initialization,
input I/O and output serialization are outside the primary measurement.
Fresh-process wall times are also retained in the JSON.

There are twelve runs per backend/workload in alternating JS/Go and Go/JS
order. Calibration selects one pass count for both backends. Calibration and
one subsequent warm-up process per backend/workload are excluded; no outliers
are removed. All primary loops exceed 0.5 seconds, with a minimum of 0.684869 s.
Each Node sample starts a new process, so JIT optimization within that loop
remains part of its measured cost.

Parsing takes about 77% of the Go formatter control's time. This is a ratio
between separate workloads, not an instrumented phase breakdown. It shows a
substantial parsing cost, without determining its internal cause. Subtracting
the two times would also include formatting, printing and UTF-8 consumption,
rather than isolate the printer. String and other existing FFI code is included.

## Historical comparison

The [published tidy reference](../tidy/README.md) measures ten iterations at
process level: **283.20 ms JS / 556.44 ms Go, or 1.96x**. That baseline retains
its own timing scope. The long-loop measurements above exclude startup and
amortize JIT work over more passes.

This campaign's separate ten-iteration control measures **315.05 ms JS /
555.63 ms Go, or 1.76x**, including startup. Its driver imports additional
modules, so this is contextual evidence rather than a byte-identical replay
or a performance improvement claim. The Go compiler executable matches the
integrated native-Tuple baseline by SHA-256:
`67b3683a4e261a2875a6456bc610f8e5543e7a567daf67fe59526f77171cdee6`.

## Correctness and allocations

The timed parser consumes `1 + 31 * imports + 131 * declarations` from the
constructed module; negative sentinels reject recovered or failed parses.
Every sample must match the verified checksum times its pass count. The
checksums per pass are 48077 for the corpus and 11801 for the fixture.

Before timing, both backends verify each corpus file, the fixture, Unicode,
an empty module, and recovered/fatal malformed inputs: fourteen cases total.
The oracle compares imports, top-level declarations and the four categories
visited by `CST.Traversal.foldMapModule` (binders, declarations, expressions,
types), and requires a complete `printModule` source round-trip for successful
inputs. These counts do not represent every CST element or token. Deep
traversal and round-trip printing are outside the timed loops.

The matched formatter uses tidy's reference options and consumes all UTF-8
output bytes with FNV-1a32. Its expected per-pass checksum is 3241266008 for
the 14,624-byte formatted fixture.

Three separate allocation runs per backend/workload follow runtime timing.
Go uses `runtime.MemStats.TotalAlloc` and `Mallocs` deltas after an initial GC.
JS sums V8 `allocated=` fields between forced GC boundaries using
`--trace-gc-nvp`; invalid wrapped values are rejected. This is a heap-growth
estimate, not an exact equivalent of Go's counter. Neither measures peak RSS
or retained memory. The format control separately records 14,624 bytes of
JS UTF-8 Buffer payload per pass. Allocation instrumentation is disabled in
the runtime campaign, and source loading is outside allocation counters.

## Builds and reproduction

The campaign used Apple M4 Pro, Node 24.8.0 (`--stack-size=4000`), Go 1.27.0,
Spago 1.0.3, official PureScript 0.15.15 for JS and the typed 0.15.17 fork for
Go, with gopurs 0.1.1. Both workspaces select package set 77.7.0 and
`language-cst-parser` 0.14.1. JS uses registry libraries; Go selects the local
51-library FFI port set. Both use the adjacent tidy checkout for the control,
and share one `Main.purs` through a symlink. GC and thread environment settings
are recorded in the report.

The benchmark sources live in the adjacent `bench-candidates` workspace:
`cst-app/` (Go), `cst-app-js/` (JS), and `cst-bench/` (measurement drivers).
From the `altbak.pub` root, after making those sources and built compilers
available:

```sh
cd ../bench-candidates/cst-bench
python3 build.py
python3 measure.py --runs 12 --output results/cst-repeat.json
```

The build helper recreates the two workspaces' generated output and prepares
the declared Go module dependency before running gopurs. The measurement
harness compiles its small native driver before timing, verifies outputs,
and fingerprints sources, generated artifacts and compiler provenance before
and after the campaign. Finish other builds before measuring.

`results.json` is a byte-for-byte copy of
`bench-candidates/cst-bench/results/application-2026-11-22.json`. It retains
72 primary timing samples, 18 allocation samples, 24 historical-control
samples, calibration, warm-ups, all oracle checks and source/artifact hashes.
This application measurement is separate from the core and extended totals.
