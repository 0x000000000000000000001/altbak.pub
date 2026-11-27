# Go consuming-tree continuations and integer sums

The baseline is the complete measured Box + array-fusion working tree on compiler HEAD `62f362fbd2dc90a4ac86c0423d21a3424deaeaf4`. The snapshot preserves its tracked and untracked changes, 2,482 source files, all 206 source/binary hashes matching the preceding campaign, previous after-core/after-all executables, official README and a direct todo.before copy. Existing benchmark documentation and the starting dirty diff are preserved. New changes are compared with this working-tree snapshot, not HEAD. No commit or push is made.

The official 37-case Go total is 9.526946090 ms. Historical Go/C indicators are 1.085401563430× / 6.728766618378× / 4.489297041124× / 3.640574098466× (ratio of totals / arithmetic mean / median / geometric mean). C37 remains frozen at 8.777346938669927 ms. All 738 protected benchmark/reference files and 269 library/FFI sources retain their hashes.

## Probes and selected changes

The earlier CPU profile of the actual opaque-input Red-Black Tree action places 61.25% of flat samples in its consuming balance worker and 26.20% in insertion. That worker contains 40 terminal paths but only five distinct terminal bodies after local renaming. An isolated generated-code overlay shares those bodies through labels without changing decisions or their order. Three diagnostic pairs measure median 9.679901 → 9.317070 ms (−3.75%), retaining 100,000 allocations and 3,200,000 bytes per action. A candidate profile still places 62.64% of flat samples in balance; percentages from separate profiles are not interpreted as absolute time gains. Sources, raw profiles and observations are retained. The original exact shell launch arguments were not saved; supplied commands are reproduction instructions, not a reconstructed execution log.

The production compiler implements this structurally in Ownership.Continuations. It traverses a Branch decision region and emits each terminal independently with the unchanged ownership emitter, original parameter environment and fresh local-name counter. Structurally equal emitted statement arrays share a scoped labeled block. Condition order, short-circuit behavior, snapshots, reads, donor choice, writes, allocation and recursive calls retain their order. A tail-loop continuation keeps its existing target. The pass stops at lexical bindings and falls back if any condition or terminal fails its existing proof. Regions with fewer than four leaves or fewer than 32 duplicated top-level statements retain the original emitter. It does not change ownership analysis, representation or mutation permissions, and does not use benchmark names.

The separate integer-sum probe exposes a dependency chain in repeated IntAdd. Two processes run copies of the actual prepared foldl-sum action with its opaque calls intact: 4732.920/4731.564 ns originally versus 1666.345/1668.184 ns with eight accumulators, with zero allocations. One-, two- and four-accumulator alternatives and 3,680 independent big.Int oracle checks are retained. These are diagnostic timings; the official measurements below use the complete newly rebuilt compiler.

ScalarArrayFold keeps its canonical saturated foldlArray/intAdd recognizer and original callback/seed/input evaluation order. Eight independent int64 accumulators process consecutive chunks, followed by the remainder and a final Int32 reduction. Addition modulo 2^32 permits this reassociation, including native int64 inputs. Reads retain their order and buffers are not mutated. Inputs shorter than eight elements retain the original loop, including an empty array returning a wide seed unchanged. The threshold follows the unroll width. Unknown callbacks and string-fold behavior keep their existing paths.

Exactly five files differ from the frozen starting tree: the new Ownership/Continuations module, Ownership/Workers, ScalarArrayFold and their two focused test files. Independent transfer checks verify the preserved Box runtime and array-fusion/cache changes outside those five files. The separate record-wrapper and traverse investigations are not included in this wave.

## Build and semantic checks

The native bootstrap checks 560 TAST modules / 348,473 types: frontend 41.7 s, generation 194.1 s and Go compilation 43.8 s. All 204 final source hashes remain fixed; 207 source/binary hashes identify the measured compiler. The pre-existing PrivateSTRegions shadowing warning retains unchanged source provenance. Available disk space was checked before bootstrap (18.26 GiB, above the 8 GiB minimum).

The combined main JS bundle passes 34 Node entries: 32 ownership tests and two scalar-fold tests. The ownership additions cover retain/swap and donor reuse, recursive tail continuations, nullable roots and lexical-binding boundaries. The existing suite retains collision, invalidation, snapshot and double-donation checks. Scalar tests execute native and boxed generated Go paths against 1,840 independent big.Int comparisons, including every remainder width, lengths 0–10000, signed limits and wide seeds. These counts include the test harness's entries and are not conflated with the separate actual-action corpora.

The older isolated Tree smoke validates activation and semantics against its own runtime baseline. Final real-action proofs instead use this campaign's before/after native builds. On each build, 6,010 fresh consuming balance calls agree with the unchanged persistent balance oracle and pass alias/cycle checks. They cover rotations, depths 0–4, red/black/unknown outer tags, signed int64 keys and absent/present donors. Forests are fresh and disjoint as required by the existing consuming contract; the proof does not claim arbitrary aliased Go pointers are consumable. Both native and public entry points of the real opaque-input tree action return 22.

The sum proof reuses the actual prepared action and its callback adapter on three call paths. Each build passes 159 canonical Int32 observations (53 inputs, independent JavaScript Number/Int32 oracle) and a separate 156 native int64 observations (52 inputs, independent BigInt modulo-32 oracle). The real action has seed zero; varying and wide seeds are covered by the generated-Go matrix. Before/after inputs, expected values and stdout agree byte-for-byte within each domain. Diagnostic opaque overlays change the supplied inputs only; production benchmarks, FFI, C and generated sources remain intact.

## Generated scope

The exhaustive comparison records five changed core sections and six AppAll sections, with unchanged populations of 633 and 683. Headers, initialization, runtime, FFI, entry points and the timed harness are byte-identical; no package or helper migration occurs. The RBTree section changes only the consuming balance worker, producing five labels for its 40 terminal paths. Its persistent oracle, wrappers, opaque input and insertion workers remain identical.

All other changes are the same scalar-sum transformation. The actual foldl-sum action executes it on every saved-action invocation. ArrayOps changes only its noncanonical-range fallback; the measured opaque900 input takes the unchanged closed-form branch. Data_Foldable changes a callable specialized sum, but its WorkloadMap alias has no action call site in this graph. WorkloadArrayInt and WorkloadMap contain changed executable sums but are outside the selected 37 cases. Thus RBTree and foldl-sum are the two directly activated measured targets. Other measured variations are not attributed to an activated optimization solely because their timing changes. Exact diffs and call-site attribution are retained in generated-other-modules-review.

## Paired measurements

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Paired change | Go/C before → after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `Red-Black Tree` | 9356.708500 | 9538.708500 | 9165.437500 | -3.913% | 1.095396× → 1.052531× |
| `array/foldl-sum-10000` | 4.557027 | 4.679830 | 1.669632 | -64.323% | 16.787856× → 5.989436× |

All four indicators use the same 37 Go/C pairs: primary App14 medians and AppAll23 library medians. C37 remains 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.085402× | 1.106428× | 1.063643× |
| Arithmetic mean μ | 6.728767× | 6.809345× | 6.555264× |
| Median Q₂ | 4.489297× | 4.446450× | 4.546005× |
| Geometric mean G | 3.640574× | 3.694998× | 3.620462× |

Relative to the fresh paired baseline, total decreases 3.867%, arithmetic mean 3.731% and geometric mean 2.017%. Median increases 2.239% and is retained. Go total changes 9.711498479 → 9.335966958 ms (−375.531521 µs). The two activated targets contribute −376.281198 µs, while the other 35 cases together add 0.749677 µs. These remaining variations are not assigned to the two source transformations.

Against the official historical baseline, total decreases 2.005%, arithmetic mean 2.579% and geometric mean 0.552%; median increases 1.263%. Historical target decreases are 2.044% for the tree and 63.361% for the sum. The paired before tree is slower than its historical baseline, so the historical and paired gains remain separate.

The highest remaining Go/C ratios are successful traverse 28.11×, large records 27.68×, medium records 17.56×, append 12.27× and Effect.Ref write 11.66×. The sum's 5.99× and the tree's 1.053× remain above the 1× objective; the aggregate improvement does not imply all four indicators improved.

## Retained regression and protocol

The primary fail-last traverse increases 0.752678 → 0.961502 µs (+27.744%), with all three after samples higher. Its generated action, runtime and harness are unchanged. A separate six-process AB/BA/AB diagnostic using its actual action and unchanged measurement loop reports 0.675845 → 0.676338 µs (+0.073%). It does not reproduce the increase. Because this diagnostic uses a separately linked executable, it does not establish the cause of the full-suite increase or erase it. The original 0.961502 µs remains in the four indicators, README and raw archive.

The secondary AppAll AST result also increases 6.772%; the primary App14 AST decreases 4.580%. Both populations and their samples remain retained, and the 37-case indicators continue to use the designated primary App14 value. No secondary value replaces a primary one. The targeted extra diagnostic is limited to the primary increase over 5%.

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, batches calibrated to at least 10 ms and the minimum of ten batches per process. Each published time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All compilation, tests and profiling stop before timing. The host is not isolated. C and JS are not remeasured. All 306 expected campaign outputs and source/binary guards pass. The six diagnostic outputs are separate. Independent recomputation reproduces the four indicators from retained samples.

The archive contains the frozen dirty baseline, scoped transfers, five-file compiler/test delta, raw CPU profiles and diagnostic variants, bootstrap logs, generated-code review, tree oracle and sum corpora, all primary/secondary measurements and separate regression diagnostic. Earlier array and record corpora are not recounted as new validation. Only existing Go Release README cells are updated; tables, references and status labels are preserved. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-27-go-tree-sum.json) and var/benchmark/go-tree-sum-20261127.
