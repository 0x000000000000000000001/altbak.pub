# Go grouped record reads optimization

The previous record-set step guards the boxed-to-native conversion. Its final scalar consumer still performs a separate linear `RecordGet` for each original field. A diagnostic extracted the actual updated worker, guarded conversion and consumer from the frozen after-all Go source. Replacing only the final read group at its consumption point reduced that isolated action from roughly 193–219 ns to 26–27 ns. Opaque runtime inputs were retained; reversed shapes exercised the unchanged fallback. These are diagnostic numbers, not the official benchmark result.

`Gopurs.RecordReadGroups` is a final pass over the generated Go AST, before import collection and printing. It admits an integer expression containing exact runtime IntAdd/IntSub/IntMul/IntDiv/IntMod calls, integer casts and scalar leaves. Boxed record reads must already have been lowered to `RecordGet(local, literalLabel).IntVal`; native record fields and evaluated scalar locals remain unchanged. This native IntVal/operator boundary is the representation chosen by TAST-driven lowering. The real optimized IR sometimes has an unannotated Local under Accessor, so the pass uses the proven lowered operation.

One candidate contains reads from exactly one boxed local and more than five distinct labels. Labels are deduplicated and sorted for the canonical native-boxing order. A fresh local guard checks non-null TypeRecordData, the complete key count, sufficient value count and equality of every key. Only then are reads replaced by positional IntVal reads. The complete unique shape proof preserves first-match behavior with duplicate labels. Missing, extra, reordered or duplicate keys, maps, constructors, getters and unsupported representations execute the exact original expression.

The pass matches the largest admitted tree first and prints its original fallback without visiting it again. It does not move an expression through a binding, branch, callback, loop or call argument. Unknown calls end a region; independently safe arithmetic inside an argument may still be optimized at its original position. Multiple boxed objects are not grouped into one region. No shape/view is reused between expressions or invocations, and neither records nor their keys/values are mutated. Names are chosen fresh against every scalar/source variable used by the region. Compact records through five fields remain outside this optimization, matching the existing boxing protocol.

Compiler changes are limited to the new module and its declaration-pass call in CodeGen. Existing conversion guards, arithmetic order/Int32 behavior, benchmarks, FFI ports, runtime and C references are unchanged.

Validation: 47 related unit tests passed. The new Go matrix executes 196 comparisons covering canonical/reversed/extra/duplicate/missing/short/uninitialized layouts, maps and foreign representations, effectful constructor getters and panics, repeated projections, native prefixes, Int arithmetic including noncanonical int64 payloads, unknown mutators between regions, result-width preservation and temporary-name collisions. Additional replay checks mutate keys between calls to ensure no layout proof is cached.

The current top-down analysis can revisit subtrees in long pure arithmetic chains that contain no eligible record group. A cached summary could reduce compiler work if profiling establishes a material cost; this step introduces no additional tree-analysis machinery.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `record/set-small-1000` | 0.040 | 0.042 | 0.042 | 1.01× |
| `record/set-medium-1000` | 0.086 | 0.089 | 0.047 | 1.89× |
| `record/set-large-1000` | 0.21 | 0.22 | 0.067 | 3.27× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.10× | 1.14× | 1.15× |
| Arithmetic mean μ | 23.29× | 23.98× | 20.40× |
| Median Q₂ | 5.91× | 6.07× | 6.10× |
| Geometric mean G | 6.42× | 6.61× | 6.36× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-record-reads.json) and `var/benchmark/go-record-reads-20261126`.

## Regression review

The primary 37-case combination (App14 and AppAll23) contains these increases above 5%; they are retained in the measurements:

- `foldable/traverse-maybe-1000`: +10.63%.

The large and medium record workloads improved by 3.2681× and 1.8918× in the paired measurement. Small uses byte-identical generated code; its observed timing changed from 0.041794 to 0.041543 µs (-0.6006%) and is retained without attributing that difference to the optimization.

`foldable/traverse-maybe-1000` increased from 3.806997 to 4.211782 µs (+10.6327%) in the primary pair. A single separate AB/BA/AB diagnostic used the same generated prepare and `Bench.runBench`, whose source sections are byte-identical between variants, with GOGC=800, default GOMAXPROCS and no PGO. Its two probe executables differ from the original AppAll executables; both were built before the six diagnostic processes. Six outputs were validated. The diagnostic medians were 3.765625 and 3.713918 µs (-1.3731%): the increase above 5% was not reproduced in this isolated probe. The primary samples, original executables and aggregates remain unchanged. The diagnostic does not establish the cause of the primary difference.

Across all 37 primary pairs, the total increased from 9.980297287 to 10.058322563 ms (+0.7818%), and the median Go/C ratio increased from 6.066847 to 6.101720 (+0.5748%). Neither improved in this pair. The arithmetic mean decreased from 23.978909 to 20.400734, and the geometric mean from 6.614433 to 6.361891. These paired results are distinguished from comparison with the separate official historical column; no overall measured speedup is claimed.

The complete generated core Go is byte-identical, including the date/time projection previously added to `Control_Bind`. In AppAll, only the medium and large record modules change. Small, the runtime, the harness and other generated modules remain identical. The new guarded read groups preserve the previous arithmetic expression as their fallback and retain the earlier record-conversion optimization. Variation outside the two changed modules is recorded without claiming its cause.

Runtime input validation covers five record representations, including the exact boxed values returned by the original prepares, varied Int32 fields, repeated actions and counts from 0 to 65,535. Fields and keys remain unchanged. Negative counters are not claimed as exhaustively executed oracles because the original decrementing behavior can require billions of iterations; this pass changes reads inside expressions and preserves that update behavior.

Build validation records the pre-existing warning in `PrivateSTRegions.purs` at line 266 and verifies that the source is unchanged from the frozen baseline. An independent recomputation of all 37 medians, ratios and four indicators is archived in `proofs/independent-metrics-check.json`.
