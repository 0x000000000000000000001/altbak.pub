# Go diagnostics — September 21, 2026

These two measurements supplement the [Go table](../../README.md#go). They are
excluded from its historical fourteen-case total. The [full archive](2026-09-21-go-diagnostics.json)
contains process samples, checksums, source and executable hashes, tool versions,
settings and raw-log hashes. All twelve processes ran sequentially after builds.

## Array indexing

The same PureScript loop reads a runtime-created array, wraps its index and
accumulates a checked checksum. Construction is outside timing. Each measured
batch performs 8,388,608 reads; three warm-ups precede ten batches. The value
below is the median of three independent process minimum times, per read.

| Elements | Go native slice | Go boxed array | JS native-input case | JS boxed-input case |
|---|---:|---:|---:|---:|
| 16 | 0.476 ns | 0.500 ns | 2.994 ns | 3.001 ns |
| 1,024 | 0.459 ns | 0.473 ns | 2.992 ns | 3.013 ns |
| 16,384 | 0.451 ns | **0.472 ns** | 3.006 ns | **3.016 ns** |

The README cells report a complete batch of 8,388,608 reads in the boxed case
at 16,384 elements: **Go 3.962167 ms / JS 25.303333 ms**.
These are the same validated measurements expressed per batch; the table above
retains the per-read breakdown. This is loop throughput including index
wrapping/checksum work, not the latency
of an isolated memory access. JS uses ordinary arrays in both cases. The Go
harness calls native entry points generated from PureScript; it excludes a
public boxed adapter that would convert the whole native input once per call.

Go's median allocation is **128 bytes per complete batch**, independent of all
three lengths and both representations (about 0.0000153 bytes/read; maximum
observed batch: 144 bytes). It is not zero allocation. Generated-code inspection
confirms that the boxed loop indexes `[]Value` before selecting `.IntVal`, without
whole-array conversion. A bounded preflight and code audit reject the former
copying behavior. Every measured checksum is validated by both driver and collector.

Go uses GOGC=800, GOMAXPROCS=1 and PGO off. JS is the official PureScript emitter
from the same typed frontend, without PBO JS optimization. These results do not
constitute a comparison with the existing `Array Processing` row, which performs
`range → filter → foldl` and has a different timing boundary.

[Kernel and driver protocol](../array-indexing-benchmark.md).

## JSON to typed AST

The measurements below are the historical baseline. The README now uses the
[September 22 runtime correction and new campaign](2026-09-22-json-typed-ast.md).

The versioned fixture contains **12 real TAST/tcorefn modules, 5,545,093 input
bytes and 21,574 type-table entries**, selected from the frozen gopurs-aff corpus.
It includes arrays, lists, maps, ADTs, Aff and Test.Main. The archive lists each
module and its input hash; the compressed fixture is stored in `test/fixtures/json-typed-ast`.

| Entire corpus, sequential application loop | Go | JS | Go / JS | Go allocated per corpus |
|---|---:|---:|---:|---:|
| JSON parsing | 73.933 ms | 19.726 ms | 3.75× | 65.262 MiB |
| Typed AST decoding from parsed JSON | 607.495 ms | 61.491 ms | 9.88× | 679.509 MiB |
| Parse + decode, measured together | **717.160 ms** | **83.839 ms** | **8.55×** | **744.781 MiB** |

The combined cell is measured separately, not obtained by adding minima from
separate phases. Two warm-ups precede five corpus samples per phase and process;
the cell is the median of three independent process minimum times. Inputs are
read before timing, and each invocation rebuilds parsed/decoded state as required
by its phase. Decoding-only reuses the parsed JSON, never a previously decoded AST.

Go and JS execute the same PBO decoding source through their respective library
implementations (native Go overrides versus the standard JS packages). The Go
JSON parser is encoding/json; JS uses JSON.parse. Both retain generic callback
adapters in this diagnostic. File I/O, sorting, compiler optimization/emission and
fingerprint construction are outside the measured interval. Validation between
samples allocates memory, so these timings describe this documented harness and
its GC state, not a decomposition of a complete compiler invocation.

Go uses **GOGC=100 and GOMAXPROCS=1**. This restricts Go GC workers as well as
application code; Node retains its default runtime background threads. Thus the
ratio is not an equal-total-CPU comparison and must not be extrapolated to the
compiler's default parallel loading or its end-to-end Go/JS ratio. A separate
campaign is needed to measure those settings.

All decoded module structures agree across the six processes, including AST
constructors, annotations, types, declarations, foreign types, source spans and
usage facts. JSON round-trip hashes are also checked. Fingerprints use a tagged,
ordered structural encoding, normalize JSON escaping, and are checked outside
timing after every corpus sample. The corpus exercises successful decoding;
it is not a replacement for PBO's malformed-input tests.

The result identifies a remaining cost in the **typed decoding phase, including
its allocations and GC**. It does not by itself attribute that cost to a specific
PBO function or establish that replacing the JSON parser would remove the gap.

## Reproduction

The sources follow the existing layout: the PureScript indexing kernel is
`src/Test/ArrayIndexing.purs`; the extended compiler workload and its FFI are
`src/Test/JsonTypedAst.*`, with its fingerprint helper alongside them. Its fixed
inputs and expected hashes live in `test/fixtures/json-typed-ast`.

Use the existing backend entry points from the repository root, with the local
TAST fork of `purs` and installed dependencies. Build before measuring:

```sh
./bin/go/run --test ArrayIndexing --build-only
./bin/go/run --test ArrayIndexing --run-only
./bin/js/run --test ArrayIndexing --build-only
./bin/js/run --test ArrayIndexing --run-only
./bin/go/run --test JsonTypedAst --build-only
./bin/go/run --test JsonTypedAst --run-only
./bin/js/run --test JsonTypedAst --build-only
./bin/js/run --test JsonTypedAst --run-only
```

`./bin/run --test ArrayIndexing` and `./bin/run --test JsonTypedAst` select the two
supported backends, Go and JS. These opt-in cases retain the detailed diagnostic
protocols above; they do not silently become `Bench.runBench` core cases or join
the ordinary `--x` suite. Each selected runtime collects three processes.
The backend entry points delegate clock/allocation/phase measurement to helpers
under `bin/benchmark`. The JSON/PBO dependency is confined to its isolated build;
ordinary extended snapshots exclude those compiler-specific sources.

Optional `--build-dir PATH` selects an isolated workspace; `--output NEW_DIRECTORY`
selects a new measurement campaign. Default workspaces live under
`run/bak/<backend>/modes/test-<name>`. Builds currently prepare both Go/JS artifacts;
measurement runs only the requested backend. Do not run other builds or benchmarks
while measuring. The original raw logs remain under
`var/benchmark/{array-indexing,json-diagnostic}/campaign-20260921`.

The published measurement archive is immutable: it retains the original paths
under `test/diagnostics` used for that campaign. The subsequent structure correction
moves sources and updates module names and runner entry points, without changing
the kernels or measured intervals. New manifests belong to the new locations;
the old source hashes are not rewritten to pretend a new measurement occurred.

The [structure validation](2026-09-22-diagnostic-structure.json) records successful
execution through all four Go/JS entry points (three processes each), unchanged
JSON/AST reference hashes, and nineteen passing runner/publication checks.

The published September 20 altbak.pub baselines remain unchanged. In particular,
its Go Array Processing cell is 19.262085 μs for the complete original kernel;
its core total is 13.278496 ms. Neither is an earlier measurement of these new
operations, and no historical speedup is inferred from the new cells.
