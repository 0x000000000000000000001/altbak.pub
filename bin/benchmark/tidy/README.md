# purescript-tidy application benchmark

This benchmark measures parsing, formatting and consuming the complete output
of [purescript-tidy](https://github.com/natefaubion/purescript-tidy). It is an
application measurement, outside the numeric core and extended-test totals.
This page describes the November 22, 2026 campaign; the
[recorded results](results.json) retain every sample and output check.

## Current results

| Variant | Median runtime | Cumulative allocated bytes |
| --- | ---: | ---: |
| Official JavaScript | 283.20 ms | ~820,184,792 (V8 estimate) |
| Go before native Tuple construction | 603.05 ms | 1,062,217,504 (`TotalAlloc`) |
| Go with integrated native Tuple construction | 556.44 ms | 927,906,880 (`TotalAlloc`) |

The integrated correction reduces Go runtime by **7.73%** and allocated
bytes by **12.64%** on this workload. Go allocation count falls from 22,908,104
to 19,642,993 (**14.25%**). Its ratio to official JS moves from **2.13×**
to **1.96×**. The published tables use the integrated build.

## Build and rerun

Build both application workspaces as documented in the adjacent checkout's
`GOPURS.md`, selecting official `purs` for `bench/` and the typed compiler for
`bench-go/`. Prepare the Go module dependencies with
`library-go-dependencies.mjs` before invoking gopurs, including on a fresh
output directory. Once all builds have finished, run from the `altbak.pub` checkout:

```sh
cd ../gopurs-tidy
python3 bench/measure.py --runs 24 --output bench/results/application-repeat.json
```

To reproduce the three-way comparison, preserve gopurs builds before and
after native Tuple construction integration. Generate both variants from the
same corrected TAST in the canonical `bench-go` workspace, moving the before
`output` directory and `tidy-bench` binary to a separate directory before
building the integrated version. Keep the typed PureScript compiler and all
application and dependency sources identical. The application's `GOPURS.md`
provides the application build commands. Then run:

```sh
python3 bench/measure.py --runs 24 \
  --before-go-output /path/to/before/output \
  --before-go-binary /path/to/before/tidy-bench \
  --output bench/results/native-tuples-integrated-repeat.json
```

After saving the report, remove the temporary baseline directory. The report
keeps its measurements and artifact hashes.

The harness verifies the fixed output oracle before timing and fingerprints
sources, generated code, dependencies and executable payloads. It rejects
changes during the campaign. The small allocation-driver build finishes
before any timing and its temporary workspace is removed automatically.
The report published here is a byte-for-byte copy of
`bench/results/native-tuples-integrated-2026-11-22.json`.

<details>
<summary>Technical reference: timing boundaries, allocations and validation</summary>

## Workload and timing

The fixture is the 514-line `bench/fixture.purs` in the application checkout,
a copy of `src/Tidy/Doc.purs`. Formatting uses page width
80 and indentation width 2. The output is 14,624 bytes. Each process performs
ten complete parse/format/consume iterations and prints the sum of their
32-bit FNV-1a hashes: `32412660080`. The consumer reads all 146,240 output bytes;
checking only the string length is not the published result.

JavaScript encodes each result with `Buffer.from(text, "utf8")` before hashing.
Go hashes its native UTF-8 string bytes directly. The input and formatted
fixture are ASCII. The published runtime executes the actual `Main.main`
applications; a separate measurement driver checks output and allocations.

Runtime is the wall-clock duration of a fresh process, including startup,
module loading and fixture reading. The campaign uses 24 executions per
variant, cycling through all six orders of JS, Go before native Tuple
construction and Go with the integrated correction. There is one excluded
warm-up per variant and no outlier removal. Compilation and profiling are completed outside timed runs.

Measurements use an Apple M4 Pro with ten performance cores and four efficiency
cores, Go 1.27.0, and Node 24.8.0 with `--stack-size=4000`. The JavaScript
frontend is the official PureScript 0.15.15 compiler; Go uses TAST from the
typed 0.15.17 fork with the TypeApp optimizer correction. Both Go variants use
byte-identical TAST, application sources and dependency inputs. They differ in
gopurs: the integrated version constructs native Tuples directly and preserves
the canonical conversion to struct-value representations. Compiler provenance,
working-tree source hashes, generated headers and artifact hashes are recorded
in the report. GC and runtime thread settings retain their defaults.

## Allocations and verification

Three separate allocation runs cover each variant, rotating their order.
Input loading is outside allocation counters. Go reports `runtime.MemStats.TotalAlloc` and
`Mallocs` deltas around the ten iterations.

The JavaScript estimate sums `allocated=` from `--trace-gc-nvp` between forced
GC boundaries. This field is the change in V8 object-heap size between GC
events, not a strict monotonic allocation counter. Invalid unsigned-wraparound
values are rejected; an earlier exploratory empty control demonstrated why
the estimate must not be treated as an exact counter. No empty control is
subtracted.
The [Node 24.8 implementation](https://github.com/nodejs/node/blob/v24.8.0/deps/v8/src/heap/gc-tracer.cc#L759)
defines the counter. This campaign reports that estimate only; Inspector
profiles from earlier builds are not evidence for the current JavaScript
artifact. UTF-8 buffer payload is accounted separately because it is external
to the V8 object heap. Neither backend's allocation figure measures peak or
retained memory.

The principal fixture has SHA-256
`8faa34f297eefb2ebb7aaedde6d71e23d7766963237a43a9af0d77c36ee9ac0b`.
Its formatted output has SHA-256
`c8f7710cf563c22508c32ced1e9bbd15a12f3ea547e88bee2410cd5dfbb3395e`.
JavaScript and both Go builds produce identical bytes and per-iteration
hashes on 35 inputs: the fixture, all 31 snapshot inputs, Unicode and layout
examples, and a malformed input whose expected result is empty. Snapshot
inputs use the benchmark's formatting options; this comparison does not run
their individual formatting directives. Every timed process must print the
expected ten-iteration checksum. The report preserves all checks and all
72 timed samples.

All 273 TAST JSON files are byte-identical between the two Go builds, preserving
type declarations, annotations and TypeApp nodes. The recorded validation also
includes 42 focused compiler tests; this campaign does not claim full
compiler-suite coverage.

## Table publication

The application section has its own runtime and allocation tables. It is not
part of the 14 core, 23 included-WIP or five extended cases, and does not change
their sums or ratios. The existing core and WIP publishers preserve this
section. Compiler timing is outside this campaign.

Earlier exploratory comparisons used JavaScript emitted by the typed
0.15.16 fork, with length-only and separate full-consumption drivers. Their
historical reports remain in the application's `bench/results/` directory.
They are not official-JavaScript measurements and do not populate the current
application tables. The current fresh builds, actual application entrypoints
and raw samples are recorded in `results.json`.

</details>

<details>
<summary>Earlier campaigns and their different baselines</summary>

The preceding TypeApp campaign measured 289.06 ms for official JavaScript,
621.78 ms for Go and 1,063,512,040 bytes of Go allocations; its report remains
in `bench/results/typeapp-go-2026-11-22.json`. The earlier typed 0.15.16 campaign
measured 281.23 ms for JavaScript, 550.39 ms for Go and approximately 940 MB of
Go allocations, retained in `bench/results/application-2026-11-22.json`.
These are historical references. The percentage gains above compare the
before and integrated Go variants measured together in the current campaign.

</details>
