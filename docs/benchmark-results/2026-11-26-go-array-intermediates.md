# Go array intermediates optimization

Baseline: compiler HEAD `fe6bb9d61bf889343fcea863cb13ca4b1341cfe1`, clean before this step. Campaign preserved the complete baseline and matched 200 previous hashes. No benchmark, FFI port, harness, C or JS reference was changed.

The isolated diagnostic (`probe/isolated`, opaque CLI seed 2147483607, lengths 0/1/17/1000/10000, Int32 boundary values and failure positions) confirmed that native intermediate buffers reduce construction/conversion cost. Observed diagnostic values are archived separately; official claims must use the alternating campaign and README baselines. Fusion was probed but is not implemented: all producer work still finishes before the fold starts.

`ClosedIndexedFolds` eta-expands a nonrecursive canonical partial integer indexed fold only when its callback is a closed adjacent lambda with the existing primitive-body proof, and its seed is an integer literal. The original IR carries stacked instantiated and generic TAST annotations. Every wrapper must specialize consistently to Array Int -> Int; Prim.Array is normalized, variables unify consistently within each annotation, and concrete contradictions are rejected. The fresh input level follows every existing local. Ordinary worker metadata, owner dependencies and emission keys remain responsible for publishing this native ABI; no external implementation summary or extra cache protocol was added.

`ModuleWorkers` preserves the generic Value wrapper using the same indexed-fold admission as the native worker. Its boxed input is viewed directly as []Value; it does not copy to []int64. The native worker consumes []int64 directly. This protects higher-order and boxed consumers from an additional array conversion.

`CallExprs` requests the native array representation only for a saturated unary Array Int worker. This retains the native map result through residual annotations and avoids moving an array snapshot across another argument's evaluation. Multiargument and overapplied calls keep their existing translation. `ArrayConcat` recognizes only the canonical saturated concatenation in this unary demand: both source values are captured before either buffer is read, then one fresh []int64 is produced. A right-hand factory can mutate the left buffer before that snapshot; the compiled Go test covers this ordering explicitly. Source storage is never reused for mutation.

`ArrayTraverseTotal` retains the existing complete totality proof and delayed allocation, but may produce the canonical Maybe(Array Int) pointer containing []int64. Native constructor choice is explicit in `AdtExprs`, restricted to the exact canonical type and matching pointer; Nothing is the typed nil Just leaf. Default sum layouts, explicit boxed consumers, unsupported types and unknown/partial callbacks keep their prior paths. Generic boundaries still perform normal boxing/conversions. This step does not prove that every escaping intermediate array can avoid those conversions.

Validation covers generated Go and AST rejection cases, Int32 arithmetic, empty input, replay and source immutability, zero allocations for the generic fold getter, mutable alias evaluation order, native/boxed Maybe conversion including Just [], concrete annotation contradictions, repeated type variables, and exact pointer/constructor mismatch rejection. The related suite also covers indexed/scalar folds, native callbacks, existing traversal strictness, array conversion fusion, local worker wrappers, tuple construction and native sums. Source E2E/cache tests, five-action opaque runtime proofs, real-workload smoke checks and full builds are archived separately.

A diagnostic smoke exposed two missing activations: inner generic annotations initially blocked eta-expansion, and a residual map annotation initially forced boxing. These were corrected with structural specialization and explicit unary representation demand, with regression tests. An early smoke overlapped a compiler-local warning rename and was retained as diagnostic only. A first native build was stopped before publication to narrow unary admission. The existing PrivateSTRegions warning is outside this change and remains untouched.

An independent unsplit factory-order fixture exposed an existing order discrepancy: the frozen baseline and this compiler produce the same 33 observations, while JS evaluates callback/seed factories earlier than map/input factories. The exact source and both outputs are retained by the semantics proof as a known prior issue. The separate explicitly sequenced fixture validates producer/fold phase order without changing benchmarks or references. This optimization does not claim to resolve that prior discrepancy.

Final compiler validation: 34 related unit/compiled-Go tests passed; JS bundle succeeded without new warnings. Native build succeeded (TAST 46.6 s, generated Go 206.7 s, go build 46.7 s), publishing bin/gopurs-native. Final smoke verified all three outputs, native map with zero box/unbox, native append and traversal, and the preserved generic fold getter.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `array/map-10000` | 40.21 | 41.34 | 28.86 | 1.43× |
| `array/append-1000` | 7.31 | 7.51 | 5.94 | 1.27× |
| `foldable/traverse-maybe-1000` | 4.21 | 4.15 | 3.10 | 1.34× |
| `foldable/traverse-maybe-fail-first-1000` | 0.034 | 0.035 | 0.034 | 1.01× |
| `foldable/traverse-maybe-fail-last-1000` | 1.20 | 1.21 | 0.80 | 1.50× |
| `array/filter-10000` | 30.27 | 32.13 | 24.30 | 1.32× |
| `array/sort-descending-10000` | 62.63 | 63.37 | 51.70 | 1.23× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.15× | 1.18× | 1.18× |
| Arithmetic mean μ | 20.40× | 20.71× | 18.58× |
| Median Q₂ | 6.10× | 6.14× | 5.93× |
| Geometric mean G | 6.36× | 6.48× | 6.17× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-array-intermediates.json) and `var/benchmark/go-array-intermediates-20261126`.

## Regression review

No published case increased by more than 5% in the primary pair.

The seven consumers changed in the generated Go. Map, append, successful traversal, fail-last traversal, filter and sort became faster in this pair; fail-first changed by −1.28%. The total remains essentially unchanged: 10.350721674 → 10.356686942 ms (+0.0576%). This is not an overall speedup. The arithmetic mean decreases by 10.305%, the median by 3.460% and the geometric mean by 4.893% relative to the paired before measurements.

Every generated core Go section is byte-identical. AppAll changes exactly the seven consumers and the shared weightedChecksum module; headers, runtime and other Go sources are unchanged. Map, append, traversal and filter retain native intermediate arrays through the checksum. Sort still converts its boxed sorted result to []int64, but no longer boxes it again before the fold. These buffers remain materialized: producer callbacks finish before fold callbacks. Static source occurrence counts are not allocation measurements.

No primary case increased by more than 5%, so no extra counter-measurement was run. AppAll's duplicate AST Evaluation observation increased by 6.102%; it remains archived as a diagnostic and does not replace the primary App value, which increased by 2.823%. All primary samples and aggregates are retained. The host is not isolated; unchanged-source timing variation does not establish a compiler regression.

The source fixture validates 33 observations in each JS/native Release/Dev profile, plus incremental-versus-clean cache transitions. Five actual prepare/action workloads additionally pass 471 varied-input checks before and 471 after, with byte-identical inputs and outputs, dynamic replay, immutable arrays and the declared opaque calls. The 5 new and 34 related tests include 312 generated-Go checks. The pre-existing PrivateSTRegions.purs:266 warning is preserved with unchanged-source provenance.

One separate, unsequenced factory fixture has a known pre-existing order discrepancy: the frozen compiler and final compiler produce identical 33 observations, but one trace differs from JS. Its original source, all outputs and explicit non-conformance status are archived under fixture-js-diagnostic-2. The passing fixture uses explicit sequencing to test producer/fold phases and cache transitions. This change does not claim to fix the prior discrepancy.
