# Go map/append producer-fold fusion

The baseline is the complete measured Box working tree on compiler HEAD `62f362fbd2dc90a4ac86c0423d21a3424deaeaf4`. Its existing runtime.go modification and untracked Box test are preserved. The snapshot captures 2,477 source files, all 205 compiler source/binary hashes matching the previous campaign, previous after-core/after-all executables, official README, preceding publication and a direct todo.before copy. New changes are compared with that frozen working tree, not HEAD. The starting diff, including untracked source, is retained. No commit or push is made.

The official 37-case Go total is 9.723599745 ms. Historical Go/C indicators are 1.107806244067× / 7.325857265806× / 4.534561961959× / 3.810954850826× (ratio of totals / arithmetic mean / median / geometric mean). C37 remains frozen at 8.777346938669927 ms. All 737 protected benchmark/reference files and 269 library/FFI sources retain their hashes.

## Probe and selected change

The isolated array probe finds an allocation and a second traversal between the producer and the common weighted checksum. Its final inlinable kernels measure map 9015 → 4965 ns (1.82×) and append 1938 → 1009 ns (1.92×), with no intermediate allocation. Each passage includes 3,700 comparisons against the frozen generated kernels and actual checksum. These diagnostic figures are not substituted for the paired campaign below. The first named helper was not inlined; its rejected measurements and compiler diagnostics are retained. Traverse also showed potential in the probe, but is not implemented in this wave and receives no attributed gain.

The compiler now fuses a canonical arrayMap or concatArray directly consumed by an eligible indexed integer fold. The imported fold descriptor records a literal seed and closed primitive callback. The mapper must also be a closed primitive lambda. Type annotations are checked at every layer; incompatible annotations, unknown calls, captures, recursive scopes, uncurried/staged calls and stored intermediate results retain the original path. This is a structural optimization without benchmark names or workload constants.

Map evaluates the source expression once, then maps and folds each element in order. Append captures both input expressions before reading either buffer, then visits left and right with continuous indices. Thus a right-input factory can still mutate an alias of the left buffer before consumption. Imported locals receive fresh names. Existing integer operations, wide-value fallbacks and checksum order are retained. Neither source buffer is mutated, and usage annotations are not interpreted as ownership proofs.

The worker cache stores the fold descriptor, including its absence, in FunctionInfo.arrayFold. The worker codec becomes gopurs/workers-v3; imported descriptors are revalidated and enter existing snapshot/emission keys. Old manifests remain readable for output management without reusing their workers. Cache tests change callback, seed and eligibility while preserving the source type and native ABI, comparing incremental and clean builds.

## Build, generated scope and validation

Exactly 26 scoped source/test files are transferred from the isolated worktree: 13 source and 13 tool/fixture files, including five new files. Eight main runtime/embedding/binary hashes are checked before bootstrap. The pre-Box worktree does not overwrite the already delivered Box runtime or its generated embeddings.

The native bootstrap checks 559 TAST modules / 347,923 types: frontend 42.7 s, generation 198.4 s and Go compilation 41.7 s. All 203 final source hashes remain fixed; 206 source/binary hashes identify the measured compiler. The pre-existing PrivateSTRegions shadowing warning retains unchanged source provenance. Available disk space is checked before bootstrap.

The exhaustive generated comparison finds all 633 core sections identical and exactly two changed sections among the 683 AppAll sections: Test_ArrayMap10000 and Test_ArrayAppend1000. There are no added, removed or moved sections. Headers, runtime, FFI, initialization and measurement harness remain identical. The original shared weightedChecksum worker remains identical. The action prefixes preserve preparation, opaque effect order and saved-action replay. The mapper and checksum bodies remain identical after local-variable renaming; map reads its original source in one loop, append in two loops with a left-length offset. Each former intermediate make disappears from generated source; measured allocations are reported separately below.

The generated-Go matrix passes 1,850 checks of results, ordering and alias behavior, including nonzero seeds, empty inputs, wide integers, shared buffers and right-factory mutation. Wide native integer compatibility is distinct from the canonical Int32 domain used by the independent JavaScript oracle. Actual actions additionally pass 282 immutable-input checks and 60 alias-boundary checks per build through Apply, FFI conversion and forced fallback. Before/after inputs, expected values and stdout are byte-identical. The alias corpus intentionally mutates the left buffer in the second opaque effect; production FFI and inputs are unchanged.

The preliminary JS-host source E2E runs in the isolated worktree with the pre-Box runtime; its manifest explicitly records that provenance. The final native-host E2E runs on the main compiler combining Box and this fusion. Both cover Release and Dev and six cache transitions; their 21 observable output logs agree. The native test harness reports two entries (parent and native-parallel subtest), not two distinct fixture programs. Detailed unit/cache logs and their hashes are retained in array-validation-summary.json.

## Paired measurements

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Paired change | Go/C before → after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `array/map-10000` | 8.812663 | 8.685710 | 4.779236 | -44.98% | 20.2557× → 11.1455× |
| `array/append-1000` | 1.899007 | 1.864156 | 1.008235 | -45.91% | 22.1381× → 11.9735× |

A separate allocation check runs the actual saved actions through runtime.Apply with their original opaque calls, verifying every result. Map changes from 1 to 0 allocations per action, append from 2 to 1: one allocation is removed in each case. Append is not allocation-free. The first checker wrongly expected one baseline allocation for both actions and stopped after observing append's two allocations; that failed assertion and its source/output are retained. The corrected collector records the measured counts without assuming the answer. AllocsPerRun performs 1,000 iterations plus its warm-up, with an initial result check, for 1,002 checks of each fixed expected result per build. These are repeated checks of two cases, not a varied-input corpus, and no timing from this allocation check replaces the campaign.

All four indicators use the same 37 Go/C pairs: primary App14 medians and AppAll23 library medians. C37 remains 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.107806× | 1.086025× | 1.085402× |
| Arithmetic mean μ | 7.325857× | 7.247138× | 6.728767× |
| Median Q₂ | 4.534562× | 4.452006× | 4.489297× |
| Geometric mean G | 3.810955× | 3.763840× | 3.640574× |

Relative to the fresh paired baseline, total decreases 0.057%, arithmetic mean 7.153% and geometric mean 3.275%. Median increases 0.838% and is retained. Go total changes 9.532418866 → 9.526946090 ms, a decrease of 5.472776 µs. Map and append account for 4.762395 µs of that difference; the 35 unchanged cases together decrease 0.710381 µs. Their variations are not attributed to the fusion.

Against the historical baseline, total is 2.022% lower, arithmetic mean 8.150% lower, median 0.998% lower and geometric mean 4.471% lower. The larger historical total decrease is mostly Red-Black Tree (9,545.750 → 9,356.7085 µs, -189.0415 µs), whose generated code and runtime are unchanged here. That historical shift is not claimed as an optimization gain. The paired Red-Black Tree change is only -0.1875 µs.

The largest remaining Go/C ratios are successful traverse 27.70×, large records 27.63×, medium records 17.54× and foldl-sum 16.35×. Append and map are now 11.97× and 11.15×, still above the 1× goal. Their lowered arithmetic/geometric averages do not imply that every case improves.

## Protocol and publication

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, batches calibrated to at least 10 ms and the minimum of ten batches per process. Each published time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. Compilation, tests and profiling stop before timing. The host is not isolated. C and JS are not remeasured. All 306 expected outputs and source/binary guards pass. Independent recalculation reproduces the four indicators from the retained samples. No primary or secondary increase exceeds 5%, so no additional regression timing campaign is run. All increases remain in the retained raw measurements and publication.

The archive contains the complete initial working-tree provenance, scoped source/test transfer, compiler/cache changes, probes including rejected and unimplemented trials, bootstrap logs, source/cache fixtures, real-action input and alias proofs, measured allocations and raw timings. It does not recount previous record corpora as new validation. Only existing Go Release README cells are updated, without adding tables or changing references/statuses. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-27-go-array-fusion.json) and var/benchmark/go-array-fusion-20261127.
