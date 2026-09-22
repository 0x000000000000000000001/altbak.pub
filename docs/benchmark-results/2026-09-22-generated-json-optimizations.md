# Generated JSON decoder optimizations

The changes are integrated into gopurs and measured on **fresh native-compiler
output**, using altbak's real PureScript diagnostics. There is no handwritten
replacement decoder, generated-Go patch, schema shortcut or result cache in the
published measurements. Both diagnostics validate complete decoded values.

The official README baselines before this campaign were **45.82 ms Go / 8.65 ms
JS** for JSON Decoding and **567.61 / 79.06 ms** for JSON to Typed AST. Fresh Go
controls below are close to those historical references. JS was remeasured;
its change from the old README is not attributed to a Go compiler optimization.

## Measurements

Times are milliseconds per corpus. Each phase is measured independently, so
parse and decode minima need not sum to the combined minimum.

| Diagnostic / phase | Go before | Go final | JS | Go time reduction |
|---|---:|---:|---:|---:|
| JSON Decoding: parse | 7.298 | **7.013** | 1.623 | 3.90% |
| JSON Decoding: decode | 36.709 | **30.023** | 7.631 | 18.21% |
| JSON Decoding: combined | 45.321 | **37.187** | 9.638 | 17.95% |
| JSON to Typed AST: parse | 75.726 | **76.856** | 20.307 | -1.49% |
| JSON to Typed AST: decode | 459.039 | **460.791** | 62.123 | -0.38% |
| JSON to Typed AST: combined | 565.532 | **560.745** | 77.022 | 0.85% |

The final Go/JS ratio for parse + decode is **3.86×**
for JSON Decoding and **7.28×** for the TAST.
JSON improves clearly in the paired runs. **No reliable TAST throughput gain is
established**: its final before range is 549.627–578.677 ms and after range
553.190–566.683 ms; the 0.85% median reduction is within this variation. These
measurements do not establish parity with JS or the 8–9 ms architectural prototype.

The first campaign with a broader native-sum rule observed 540.110 ms for TAST
(4.77% lower than its control); its raw results are retained in the archive, but
that is not the final compiler rule. This is not a controlled TAST-only ablation,
so the difference between campaigns cannot be assigned to one rule. JSON's JS
bundle is unchanged between campaigns, yet its median moved from 8.555 to 9.638 ms;
do not credit that variation to Go or use it as evidence of a JS regression.

Allocated bytes, expressed in MiB per corpus; cumulative allocation, not retained
heap or peak RSS:

| Diagnostic / phase | Before MiB | Final MiB | Reduction |
|---|---:|---:|---:|
| JsonDecoding: decode | 56.729 | **44.335** | 21.85% |
| JsonDecoding: combined | 62.459 | **50.065** | 19.84% |
| JsonTypedAst: decode | 496.360 | **488.143** | 1.66% |
| JsonTypedAst: combined | 561.631 | **553.406** | 1.46% |

| Binary | Before bytes | Final bytes |
|---|---:|---:|
| JsonDecoding | 6,475,474 | 6,458,482 |
| JsonTypedAst | 9,252,642 | 8,995,202 |

## What changed

1. Standard indexed Array/Either traversal becomes one private result buffer and
   a loop. It still calls every callback, in order, even after an error, and
   returns the first error. Unknown/custom dictionaries keep the generic path.
2. A general pass removes immediately consumed unary lambdas, including those
   returned by branches. It only moves values, preserves evaluation order,
   renames transplanted binders and bounds duplication. The actual
   `gDecodeJsonCons` callback pattern disappears from generated Go.
3. Existing native Maybe/Either/Tuple results survive annotations that would
   otherwise box them into Value; field reads use native slots. **Annotations
   selecting a more precise native pointer layout still get that conversion.**
   Payloads remain Value; dynamic callbacks and public wrappers retain adapters.

No parser, PBO decoder semantics, TAST format or benchmark workload changed.
The earlier shared scalar record-reader work remains a separate, limited ABI
improvement; its nanosecond gains are not credited as JSON gains here.

## Ablation of native sum preservation

A separate three-process comparison changes only the preservation guard in a
copied compiler bundle, then regenerates and compiles Go. These ablations use
JSON Decoding; they are not handwritten decoder implementations.

| Policy | Combined ms | Decode allocation MiB |
|---|---:|---:|
| Broad native-sum rule | 37.464 | 44.229 |
| Native-sum preservation disabled | 37.343 | 44.357 |
| Preserve only when annotation expects Value | 37.125 | 44.335 |

Timing ranges overlap. Disabling preservation changes decode allocation by less
than 0.3% versus the broad rule: **no substantial JSON throughput win is established
for this rule alone**. The final rule is restricted to Value boundaries to keep
already precise native layouts. The bulk of this campaign's JSON gain comes from
the traversal and closure changes, not a claim that changing outer Either storage
solves the native ABI.

An earlier isolated `decodeMaybe` branch rewrite also saved **zero allocations**
on null, success and failure (5/6/5 before and after). Go already optimizes some
local box/unbox patterns; target allocations that survive across actual calls.

## Validation and protocol

- 32 targeted tests passed, including 11,055 traversal comparisons and the
  regression check preserving native pointer layouts with int64 payloads.
- Six PureScript fixtures compile and execute: Maybe FFI, partial applications,
  shared native records, record boxing, type-changing updates and Maybe branches.
  Snapshots are updated only after successful execution.
- All timed runs check the 17 fixed JSON cases and 12 TAST modules against their
  versioned oracles. Complete hashes are checked after every timed sample.
- 317 additional JSON cases match the old Go executable, including error paths.
  Ten existing differences with JS concern Int values outside signed 32-bit bounds;
  the native Int contract is unchanged. No timed case has that discrepancy.
- Three processes per variant, two warmups and five samples per phase. A cell is
  the median of the three process minima. Order is reversed/rotated; GOMAXPROCS=1,
  GOGC=100 and PGO off. File reading, hashing and result verification are untimed.
- All builds/tests finish before timing. The paired driver verifies unchanged
  benchmark/library/fork-compiler inputs, identical TAST hashes, and frozen
  executable hashes. Only the gopurs compiler executable differs between builds.
- The diagnostic runner now selects and fingerprints the local typed `purs`
  executable explicitly; a stock compiler with the same version number is not
  accepted merely on that version number.

Raw samples, build fingerprints, validation and measurement scripts are in
[2026-09-22-generated-json-optimizations.json](2026-09-22-generated-json-optimizations.json). Local frozen workspaces:
`/Users/0x1/Documents/htdocs/scratch/json-native-abi-20260922`.

To build and measure a fresh final diagnostic from altbak:

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --workspace /tmp/json-final-build
python3 bin/benchmark/json-diagnostic.py measure --suite JsonDecoding --workspace /tmp/json-final-build --output /tmp/json-final-measure
```

Use `--suite JsonTypedAst` with separate new directories for the TAST. The README
contains the final combined milliseconds, with excluded-WIP totals unchanged.

## Remaining architectural work

The next substantial step is a shared decoder callback whose **arguments and
success payload remain native through the dictionary, worker and traversal**.
The loop still receives boxed callback results and builds a `[]Value`; record
payloads and row operations still cross generic boundaries. Outer native sums
alone do not close that gap. Preserve unknown-instance/FFI adapters, immutable
rows and fields, partial calls and errors while introducing that calling
convention. Reprofile its real generated decoder before generalizing.

There is no new complete b8x measurement in this campaign. The 8–9 ms manual
schema engines demonstrate architectural potential, not the compiler result.
See gopurs's [implementation contract](../../../gopurs/gopurs/docs/json-decoder-optimizations.md)
and tracked next steps in `gopurs/todo.md`.
