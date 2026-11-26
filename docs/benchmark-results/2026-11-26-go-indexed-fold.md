# Go indexed Array fold optimization

`Gopurs.ArrayIndexedFold` specializes the canonical `Data.FoldableWithIndex.foldlWithIndexImpl` when supplied with a known adjacent lambda, an integer seed, and optionally its input array. The two-argument form preserves partial application: callback and seed are captured immediately, input is supplied later, and every invocation starts with a fresh accumulator. The three-argument form evaluates callback, seed and input in source order before the loop.

The callback must have the exact TAST signature `Int -> Int -> Int -> Int`. The observed optimized IR sometimes leaves the first lambda without an outer type while retaining the exact `Int -> Int -> Int` type of its suffix. That form is also admitted: the canonical indexed-fold contract proves the first index parameter is Int, and the remaining argument/result types still come from TAST. An incompatible explicit outer type, absent/unknown/mismatched suffix, or different accumulator/element type is rejected.

The native callback uses three int64 parameters and an int64 result. Its primitive operations go through the existing translator, retaining PureScript Int32 arithmetic, including division and modulo semantics. The loop invokes it once for every element, in ascending index order. No reassociation, short circuit, producer fusion or checksum formula is introduced.

The old FFI wrapper both copied the boxed input array and invoked Apply3 for each callback. Directly viewing the source is admitted only for callbacks whose complete body is proven read-only: Int/Boolean literals and evaluated locals, compatible primitive operations, branches and strict nonrecursive lets. The proof checks operand and result types consistently. Calls, foreign getters, accessors, indexing, unknown captures and recursive contexts are rejected. This conservatively preserves the old snapshot behavior when an unknown callback could mutate an aliased source. A computation between curried lambdas also retains the original path.

The optimized path reads existing []Value or native array storage without an extra whole-array conversion, keeps the accumulator native, and removes per-element Apply3. It does not mutate input or remove output arrays from earlier producers. Unknown callbacks, other dictionaries, other types, insufficient/surplus arguments and unsupported syntax retain the existing FFI path.

The diagnostic probe copies the actual frozen FFI wrapper, runtime and fold implementation into an isolated Go module. Its opaque command-line seed and variable arrays include overflow cases. The existing wrapper plus partial closure took approximately 12–14 ns/element; direct FFI callback approximately 2.4 ns/element; the equivalent native loop approximately 2.2 ns/element. These establish the bottleneck, not official gains: published comparisons must use the archived benchmark campaign and README baseline.

Compiler files changed: ArrayIndexedFold.purs and its call site in CallExprs.purs. New unit tests include 192 executions of generated Go checking full/partial application, nonassociative ascending order, empty input, Int32 wraparound, replay and unchanged source storage. Benchmark source, library FFI and C references are unchanged.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `array/map-10000` | 130.18 | 127.55 | 39.10 | 3.26× |
| `array/filter-10000` | 75.32 | 73.37 | 28.38 | 2.58× |
| `array/sort-descending-10000` | 152.12 | 148.44 | 62.19 | 2.39× |
| `array/append-1000` | 25.85 | 25.09 | 6.99 | 3.59× |
| `foldable/traverse-maybe-1000` | 13.19 | 12.82 | 3.82 | 3.35× |
| `foldable/traverse-maybe-fail-first-1000` | 0.036 | 0.035 | 0.034 | 1.03× |
| `foldable/traverse-maybe-fail-last-1000` | 1.20 | 1.13 | 1.13 | 1.00× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.19× | 1.14× | 1.11× |
| Arithmetic mean μ | 47.41× | 45.72× | 27.60× |
| Median Q₂ | 6.06× | 5.89× | 5.91× |
| Geometric mean G | 8.08× | 7.79× | 6.66× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-indexed-fold.json) and `var/benchmark/go-indexed-fold-20261126`.

## Regression review

The primary 37-case combination (App14 and AppAll23) contains these increases above 5%; they are retained in the measurements:

- `AST Evaluation`: +7.73%.

`AST Evaluation` increased from 0.026680 to 0.028742 µs (+7.7286%) in the primary App pair. These samples and their effect on all aggregates are retained. The complete generated core Go, including the action and harness, is byte-identical between builds.

A separate AB/BA/AB diagnostic called the original generated action and `Bench.runBench`, with GOGC=800, default GOMAXPROCS and no PGO. Its two probe executables differ from the original App executables. All probe builds completed before the six diagnostic processes, which validated six outputs. The original App/AppAll sources, executables and primary result file remain unchanged.

The diagnostic medians were 0.028475 µs before and 0.029287 µs after (+2.8516%). The increase above 5% was not reproduced in this isolated probe. This diagnostic does not replace the primary samples or establish the cause of their difference.

Across all 37 primary pairs, the measured total decreased from 10.047846498 to 9.766831050 ms (-2.80%). The arithmetic Go/C mean decreased from 45.721264 to 27.603507 and the geometric mean from 7.792013 to 6.659963. The median increased from 5.889470 to 5.909141 (+0.334%); it did not improve in this pair. The separate official historical column is retained for context.

Five workloads that reach the shared checksum improved substantially: map, filter, sort, append and successful traverse. The two measured failing traversals return `Nothing` before reaching that checksum. Their small timing changes and the variation in the unchanged core are retained without attributing them to the optimized fold. Only the shared checksum body changes in generated AppAll code; the seven consumers change solely to qualify the relocated helper's package.
