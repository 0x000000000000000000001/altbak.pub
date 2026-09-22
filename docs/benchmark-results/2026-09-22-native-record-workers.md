# First generated shared native record worker

The first step of the native shared calling-convention proposal is now
implemented in gopurs: a proven read-only open-row parameter can be passed as a
native projection of its known primitive fields. This experiment uses **real
PureScript, the typed compiler, and gopurs-generated Go**, without hand-editing
its generated worker or replacing the computation with a Go implementation.

It is an argument projection for scalar-returning readers, not yet a general
native JSON decoder or a new representation for every row-polymorphic operation.
No new JSON/TAST or b8x timing is established here.

## Measured results

Each pipeline constructs a closed record and calls the same row-polymorphic
reader. Layout A has four fields, B five including a nested record, and C six
including an array. The worker reads Int, String and Boolean fields. Other
fields remain observable to the caller and are checked by correctness tests.

| Substantial reader pipeline | Before (ns) | After (ns) | Before bytes / allocations | After bytes / allocations |
|---|---:|---:|---:|---:|
| A | 52.610 | **2.921** | 176 / 2 | **0 / 0** |
| B | 82.120 | **3.045** | 288 / 4 | **0 / 0** |
| C | 133.200 | **3.238** | 448 / 8 | **0 / 0** |

The whole measured producer/call pipeline is **18–41× faster** in these cases.
The baseline boxes the complete record, including unused nested/array extras,
then reads it through Value. The treatment passes native fields directly.
The Go compiler can additionally optimize known field values and eliminate
unneeded record construction; this is not an isolated getter-dispatch ratio.

The existing optimized small case was retained as a regression control:

| Small reader pipeline | Before (ns) | After (ns) | Bytes / allocations, both |
|---|---:|---:|---:|
| A | 1.816 | 1.825 | 0 / 0 |
| B | 1.716 | 1.703 | 0 / 0 |
| C | 2.606 | 2.591 | 0 / 0 |

These differences are around 1% or less and establish no meaningful throughput
change. The baseline had already inlined specialized copies into these callers;
comparing against only its unused boxed worker would have overstated the gain.

A further control varies **all three fields at every call** through a real
PureScript `runFields :: Int -> String -> Boolean -> Int` function. It removes
the constant-name/flag advantage of the first producers:

| Dynamic field pipeline | Before | After |
|---|---:|---:|
| Time | 56.300 ns | **3.102 ns** |
| Bytes / allocations | 176 / 2 | **0 / 0** |

**18.15× faster**. Before samples: 56.30, 54.24, 61.00 ns; after: 2.936, 3.185,
3.102 ns. Three alternating processes per executable, 100 ms per cell, same
GOMAXPROCS/GOGC/PGO settings. Each variant also passes 1,024 combinations of
numeric seed, name and flag, plus the original 768 checks. This confirms that
the gain does not depend solely on constant field values. The baseline allocates
RecordDict4 and a string wrapper and reads fields dynamically; the treatment
uses the generated native projection and shared worker.

Official README baselines remain **45.82 ms Go / 8.65 ms JS** for JSON Decoding
and **567.61 / 79.06 ms** for JSON to Typed AST. Those are different workloads;
the nanosecond pipelines above must not be substituted into those cells, or
used to claim that the full JSON benchmark has reached the 8–9 ms prototype.

## What the compiler now emits

One source definition:

```purescript
score :: forall r. { id :: Int, name :: String, active :: Boolean | r } -> Int
```

has one generated worker accepting:

```go
struct { active bool; id int64; name string }
```

The native scalar result is retained. The three original shape-specialized
worker clones disappear; all three generated callers use the same worker ABI.
The existing Get wrapper adapts boxed and partially applied calls. There is no
interface wrapper chain, map, or full record reconstruction at a native call.

This establishes one **generated source worker body** across layouts. Normal Go
inlining is allowed; it does not promise exactly one machine-code copy. Application
binary sizes, same build flags:

| Application | Before (bytes) | After (bytes) |
|---|---:|---:|
| Small | 3,506,626 | 3,523,138 |
| Substantial | 3,523,346 | 3,523,202 |

The small executable grows 0.47%; the substantial one is essentially unchanged.
Metadata, wrappers and linker decisions remain in these sizes. This first
three-layout test is not a guarantee for arbitrary application size.

## Guards and validation

Two independent proofs are used: source TAST before deciding to share rather
than specialize, then the transformed TCO body before choosing its native ABI.
Only reads of declared primitive fields qualify; the result must be scalar.
Returning, forwarding, updating or capturing the row, unknown fields/types, and
unsupported polymorphism retain the ordinary representation. Labels that cannot
safely form distinct Go fields also retain it (blank, unsupported characters, or
sanitization collisions). Full source-record
evaluation and caller immutability are preserved.

If a source candidate fails the later proof, its ordinary ABI stays correct,
but its removed specializations are not restored. A performance regression is
possible in that situation. The tests and measurements do not claim otherwise.

Validation completed:

- Both compiler builds rebuilt; native bootstrap uses 454 typed modules.
- Eighteen targeted codegen/runtime tests pass, including native scanner checks,
  imported workers, records, dynamic/partial applications, and escape fallbacks.
- Four fixtures pass with the rebuilt native compiler: NativeRecordWorkers,
  RecordTypeChangingUpdate, NativeRecordBoxing, LocalNativeReturns. The new
  fixture's snapshot is checked into the existing passing-snapshots directory.
- Original fixed inputs, complete records, unused extra fields and immutable
  updates are checked. Each of the four measured executables also passes 768
  input/layout checks against an independent scalar oracle.
- The ordinary baseline and treatment consume identical TAST and benchmark
  harnesses. Frozen compiler bundles, generated code and executable hashes are
  recorded. No generated Go computation is manually patched. After the final
  label guard and native rebuild, all computation Go files match the measured
  treatment; only the unused Spago version string differs (1.0.4 versus 1.0.3
  in the two local runners).

The fixture runner initially picked the standard purs from PATH; that output
had no TAST metadata and could not compile the Go constructors. The successful
runs explicitly used the local typed fork. No production workaround or compiler
change was made for that environment issue.

## Protocol and continuation

Three independent processes per variant, alternating/reversed order, each Go
benchmark calibrated for 200 ms; medians of ns/op samples. GOMAXPROCS=1,
GOGC=100, PGO disabled, Apple M4 Pro / Go 1.27.0. All builds, tests and other task
benchmarks had finished before timing. The harness varies the numeric seed,
retains a checksum, and validates outputs outside the timed region.

The next architectural proof should keep a native record through a worker's
result and an Either success/error boundary, then measure the real JSON decoder.
This scalar-returning projection does not yet establish that path. Array fusion
remains separately supported by the previous generated-code experiment.

[Results, scripts, fixtures and provenance](2026-09-22-native-record-workers.json).
The complete local experiment is in
`/Users/0x1/Documents/htdocs/scratch/json-native-workers-20260922`.
Run `python3 measure.py` there with the recorded executables to repeat the main
campaign. `dynamic-control.py` prepares and measures the dynamic-field control. Compiler scope and guards are documented in
`gopurs/gopurs/docs/native-record-workers.md`.
