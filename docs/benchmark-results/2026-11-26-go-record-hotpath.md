# Go record hot-path optimization

The starting point is the measured **dirty working tree** from the previous string-checksum campaign, on HEAD `541e68a24aba929b6183105398799104783d903c`. Its eight uncommitted compiler changes are preserved, not counted as changes made by this campaign. The baseline freezes 2,475 tracked and untracked source files, the native and JS compiler binaries, the existing benchmark executables and official README. All 203 hashes of the previously measured compiler match the current state. The starting diff from HEAD, including untracked files, is retained separately in `compiler-before/diff-from-head-including-untracked.patch`. New compiler changes are computed against this complete working-tree snapshot, not against HEAD.

The official 37-case Go total is 9.440900945 ms. Historical Go/C indicators are 1.075598470810× / 9.718704982286× / 4.479246931911× / 4.374540972327× (ratio of totals / arithmetic mean / median / geometric mean). The C37 denominator remains frozen at 8.777346938669927 ms. The campaign protects 734 benchmark/reference files and 269 library/FFI sources. Benchmark and harness sources, library FFI and C/JS references remain unchanged.

## Probe and profiles before implementation

The probe calls the three real prepared record actions from the frozen generated package and checks their expected results (1006, 1072 and 1272). Those are three distinct expected cases, not a varied-input correctness corpus. Native update-only probes use a runtime CLI seed and are separate diagnostics: the existing affine fast path has already removed the repeated update loop for admitted positive counts.

The exact generated action exposes an avoidable generic FFI call for the opaque counter, although a certified immediate-effect worker already exists. A diagnostic copy changes only that call to the existing worker and caches the immutable literal-to-interface conversion using the existing convention. It keeps the opaque barrier, saved input action, record conversion, native update and original-record consumption. Under GOGC=800, actual/direct-count timings are 38.759/18.077 ns for small records, 42.693/21.958 ns for medium, and 61.197/39.651 ns for large. Each actual action allocates 64 bytes in three allocations; each diagnostic direct-count variant allocates zero. These are probe results, not official compiler campaign gains.

Final CPU profiles run for approximately three seconds with GOGC=800. Allocation profiles for two million large-record invocations attribute 121.46 MB (pprof units) to the original counter FFI wrapper/factory; the diagnostic direct-count profile has no hot-action allocation entry. Focused call-path profiles retain the full sampling denominator. An earlier exploratory global profile was dominated by `runtime.kevent` (79.40%); it is retained but does not establish application-wide CPU attribution. Likewise an initial `harness/Apply` probe accidentally constructed its closure inside measurement and is superseded by the prepared-action version (2.652 ns, zero allocation). Raw pprof data, sources, corrected observations and focused/global summaries are preserved separately.

The residual profiles point to record projection and original-record reads. They motivate later work; this campaign does not remove, reuse or move record guards across effects or assume aliases cannot mutate the record.

## Generic compiler change

`ImmediateFfiEffects.direct` shares the existing certified immediate-effect-worker path. `EffectExprs.bindEffect` now tries it when a known foreign factory application is consumed directly by an Effect bind, in addition to the existing local-binding form. Admission requires an effect block, an existing certified worker, exact saturated arity and curried application boundaries. Arguments are captured once in left-to-right order; the worker calls the factory and its returned effect thunk before executing the continuation. Only the existing immutable scalar-literal conversion is cached; dynamic arguments, factories, effects and results are not cached.

Saved, replayed, unknown, deferred, partial, overapplied and uncurried actions retain the generic path. No benchmark names, record fields or workload-specific constants enter the compiler transformation. The record representation, projection guards, arithmetic and runtime are unchanged. The prior string-checksum implementation remains byte-identical to the dirty starting snapshot.

The actual-record JS-host smoke returns all three expected values and verifies source/bundle provenance. Its generated counter call uses the immediate worker, while the saved input action and its replay stay in the original runtime path. Record guards, original-field reads and native update/sum bodies are preserved. The native opaque factory, effect worker and `opaqueValue` body are unchanged, including the `//go:noinline` barrier; passing 1000 still passes through the runtime opaque call. JS/native packaging aliases and exported record field spellings are normalized only for this preliminary cross-host source comparison; the final production builds receive a separate exact comparison.

## Native build and generated scope

Native bootstrap completes on the frozen source: frontend 38.8 s, Go generation 187.9 s, Go build 43.5 s; 557 modules and 346,695 TAST types verified. The 200 compiler-source hashes are unchanged across bootstrap. The existing `PrivateSTRegions.purs:266` shadowing warning is retained with identical source provenance. The prior runtime and string-checksum changes are unchanged from the initial dirty snapshot.

The exact generated comparison finds five changed core sections and 25 changed AppAll sections (633/633 and 682/683 section inventories). Headers, runtime and every FFI definition remain byte-identical. The added section holds an immutable argument getter; no existing worker moves packages. Eleven library actions change a foreign-effect call boundary: the three records, append, equality, both indexOf cases, split, code-point take, and both failing traversals. Sort and concat change only during preparation. The changed WorkloadArray modules are not called by the selected 37 cases; all 14 core actions remain unchanged. Calculations, record guards, array/string operations and the previously optimized checksum are preserved.

Generated `Bench.runBench` also lowers its ten `measureBatch` invocations to the existing immediate worker. The source harness is unchanged. A structural check restores those ten wrappers and normalizes only generated local suffixes to reproduce the previous body exactly: same arguments, order, warm-ups, calibration, minimum-of-ten selection and output. `Bench_MeasureBatch`, its worker and callback adapter, and `calibrate` are byte-identical. Wrapper conversion happens before the unchanged native thunk starts its clock. Thus runBench is not claimed byte-identical, while the actual timed loop is.

## Paired measurements

| Case | Official Go baseline (ns) | Paired before (ns) | After (ns) | Speedup | Go/C after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `record/set-small-1000` | 40.514 | 40.198 | 19.651 | 2.05× | 17.49× |
| `record/set-medium-1000` | 45.379 | 45.265 | 23.482 | 1.93× | 20.92× |
| `record/set-large-1000` | 65.400 | 65.395 | 41.110 | 1.59× | 30.32× |
| `foldable/traverse-maybe-fail-first-1000` | 32.929 | 32.930 | 12.394 | 2.66× | 9.23× |

The separate final allocation check invokes the actual frozen before/after actions, not the hand-modified diagnostic variant. `testing.AllocsPerRun(10000)` reports **3 → 0 allocations per action for all three records**. Expected results are checked throughout; this is an allocation check with three distinct expected cases, not another timing campaign or additional varied-input coverage. Production source and executable hashes remain unchanged.

All four indicators below use the same 37 matched Go/C cases: primary App14 and AppAll23 library medians. C remains frozen at 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.075598× | 1.084131× | 1.074409× |
| Arithmetic mean μ | 9.718705× | 9.656941× | 7.792060× |
| Median Q₂ | 4.479247× | 4.440115× | 4.447449× |
| Geometric mean G | 4.374541× | 4.341120× | 4.039488× |

The measured Go total is 9.515797056 → 9.430460051 ms, a reduction of 85.337005 µs. The unchanged Red-Black Tree alone decreases by 85.791500 µs (-0.918%); the other cases together increase by 0.454495 µs. This total decrease is not claimed as a record optimization gain. The arithmetic mean decreases by 19.31% and geometric mean by 6.95%; the median increases by 0.17% and is retained as measured. Differences from the historical baseline remain separate from fresh paired gains.

## Measurement and regression review

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, calibrated batches of at least 10 ms and the minimum of ten batches per process. Published rows are medians of the three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before timing; the host is not isolated. Neither C nor JS is remeasured. All 306 measured outputs validate; hashes are checked around collection. Independent calculations reproduce the four indicators from the retained samples.

Two primary increases exceed 5%; their generated actions are unchanged and their primary measurements are retained:

- `effect/ref-write-10000`: 13.042 → 14.000 ns (+7.345%). A separate AB/BA/AB diagnostic measures 13.324 → 13.424 ns (+0.751%).

- `st/ref-modify-10000`: 6.761 → 7.139 ns (+5.591%). A separate AB/BA/AB diagnostic measures 6.875 → 6.914 ns (+0.567%).

Neither increase above 5% reproduces in the isolated diagnostic. This does not prove the original differences were noise. The diagnostic uses separate executables and preserves the original actions, native timed loop and callback adapter; its generated runBench call lowering outside the clock is checked explicitly. All 12 diagnostic outputs validate. No primary sample or aggregate is replaced.

## Semantic validation and publication

Twenty-one unique Node tests pass: nine FFI tests plus twelve related tests, with the two repeated immediate-effect-application tests counted once. The source fixture passes 14 observations per JS-host/native-host Release/Dev execution. It covers argument/factory/thunk/continuation order, unknown and saved actions, replay and changing arguments. This campaign does not claim an intermodule cache-transition test; the ABI and dependency-key protocol of existing workers are unchanged.

The real three record actions pass 1,134 varied-input checks before and 1,134 after, through three call paths and five record representations: prepared, shape, dictionary, reversed and map. Inputs remain immutable and saved actions replay with changed runtime values. Each preparation invokes opaque once, and each measured action twice. Field values include Int32 boundaries; nonnegative bounded counts cover 0, 1, 2, 7, 1000 and 65535. Negative countdowns are excluded because the unchanged reference action would need billions of iterations.

A separate alias corpus passes 108 checks before and 108 after. Its diagnostic opaque overlay intentionally mutates a shared input alias once when the counter effect executes. Dictionary/map representations and counts 0/1/7 verify that the subsequent record projection and both consumers observe that mutation at the correct point. Here the input is intentionally not immutable; the compiler adds no mutation. All inputs and expected/stdout bytes match across before/after for each corpus. Production benchmarks, FFI and reference files remain unchanged.

The final archive contains the dirty starting snapshot metadata and starting diff, the new task delta, raw profiles, sources, commands, raw samples, manifests and validation logs. Existing Go Release README cells are updated without adding tables or changing reference columns/statuses. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-26-go-record-hotpath.json) and `var/benchmark/go-record-hotpath-20261126`.
