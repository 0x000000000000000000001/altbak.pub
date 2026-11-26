# Go Array traverse Maybe optimization

The Go compiler now emits an element loop directly for fully applied canonical Array / Maybe traversals whose callback has an exact TAST `Int -> Maybe Int` type (or `Int -> Int -> Maybe Int` for indexed traversal) and whose complete body passes a structural totality proof. The canonical dictionary recognition remains in `Gopurs.ArrayTraverse`; the proof and loop emission live in `Gopurs.ArrayTraverseTotal`.

The proof accepts only canonical `Just` / `Nothing` leaves, branches, integer literals and already evaluated integer locals, add/subtract/multiply/negate/bit-not, integer comparisons, and total boolean expressions. It rejects all calls, foreign operations, array indexing, division/modulo, unknown local types, recursive contexts, and unsupported syntax. A pure function type alone is insufficient. Parameter types come from the exact callback TAST, while captured scalar locals must carry their matching TAST type. Every branch must pass, including a branch that appears unreachable.

The specialized loop evaluates the admitted callback's primitive branches directly. It creates no callback closure or per-element Maybe object. `Nothing` returns immediately; omitted later callbacks cannot throw, diverge or have effects under the proof. Output storage is allocated only after the first successful element, so failure on the first element does not allocate a full output array. The empty successful result remains a non-nil empty array, as before. All supplied dictionary/input arguments retain their original evaluation order and application stage. Input storage is never mutated.

Unknown or partial callbacks, custom instances and partially applied traversal calls keep the existing strict loop, including callbacks after a failure and their exceptions. Existing strict-callback tests remain unchanged.

The output retains its existing `[]Value` representation. This avoids introducing a second whole-array conversion for the current generic fold consumer. Success and late failure still allocate an output array and success still passes through the existing consumer; generic consumer specialization/fusion is a separate future optimization. No benchmark source, benchmark input, FFI library or C reference was changed.

The isolated preimplementation Go probe in `traverse-total-probe.go` established that early return before allocation addresses the intended first-failure cost. Its timings are diagnostic only; production comparisons must use the archived full benchmark campaign and official README baseline.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `foldable/traverse-maybe-1000` | 16.79 | 15.90 | 13.19 | 1.21× |
| `foldable/traverse-maybe-fail-first-1000` | 3.95 | 3.76 | 0.036 | 104.35× |
| `foldable/traverse-maybe-fail-last-1000` | 3.95 | 3.69 | 1.20 | 3.07× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.29× | 1.18× | 1.19× |
| Arithmetic mean μ | 132.26× | 124.10× | 47.41× |
| Median Q₂ | 7.65× | 6.88× | 6.06× |
| Geometric mean G | 10.20× | 9.43× | 8.08× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-traverse-maybe.json) and `var/benchmark/go-traverse-maybe-20261126`.

## Regression review

The primary 37-case combination (App14 and AppAll23) contains these increases above 5%; they are retained in the measurements:

- `string/split-1000`: +5.18%.

`string/split-1000` increased from 133.658852 to 140.581703 µs (+5.1795%) in the primary pair. These samples and their effect on all aggregates are retained.

A separate AB/BA/AB diagnostic called the same generated `prepare` and unchanged `Bench.runBench`, with GOGC=800, default GOMAXPROCS and no PGO. Its two probe executables differ from the original AppAll executables; the generated target/harness sections are byte-identical between variants, and the original App/AppAll sources, executables and primary result file remain unchanged. All probe builds completed before the six diagnostic processes. Six outputs were validated.

The diagnostic medians were 124.821617 µs before and 125.697906 µs after (+0.7020%). The increase above 5% was not reproduced in this isolated probe. This diagnostic does not replace the primary measurements or establish the cause of their difference.

The combined primary total increased from 10.397847069 to 10.423703235 ms (+0.25%). The core generated Go is byte-identical, and only the three traverse modules change in AppAll. The observed total increase is retained without attributing it to the compiler change or claiming a measured total speedup. The improvement in the arithmetic, median and geometric Go/C indicators is reported separately.
