# Go record projection and grouped-read optimization

The baseline is the clean compiler at HEAD `80e360a1b7ab6126c1f2d14967c5d289d4bad688`; the benchmark repository is also clean at campaign start. It freezes 2,475 source files, the compiler binaries, official README and previously measured benchmark executables. All 203 source/binary hashes of the preceding record-hotpath campaign match before its after-core/after-all builds are copied as this campaign's before builds. Earlier string and immediate-FFI optimizations are already in this baseline. New changes are compared against the complete frozen state.

The official 37-case Go total is 9.430460051 ms. Historical Go/C indicators are 1.074408943488× / 7.792059822195× / 4.447449332664× / 4.039488081780× (ratio of totals / arithmetic mean / median / geometric mean). C37 stays frozen at 8.777346938669927 ms. The campaign protects 735 benchmark/reference files and 269 library/FFI sources. Benchmarks, harness sources, runtime, library FFI and C/JS references remain unchanged.

## Probes and profiles

The initial probe imports the real prepared record actions from the frozen bundle. Diagnostic variants use compact projection/group guards for small records and one final Int32 narrowing in the grouped sums for medium/large. It checks the three original results (1006, 1072, 1272); this is not a varied-input correctness corpus. Under GOGC=800, it measures small 20.729 → 10.431 ns, medium 22.677 → 20.577 ns and large 41.820 → 37.085 ns, all with zero measured allocations before and after. The source update loops already have an affine fast path, so these are end-to-end action costs, not costs of 1000 executed updates. The probe does not replace the official paired campaign.

The initial small-record CPU profile attributes 1.23 s / 2.69 s (45.72% of whole-process samples) to RecordGet. Its driver also selects an action from a map inside the profile loop, accounting for about 23%; that overhead is absent from the official harness and from the probe's testing.Benchmark loop, which captures the action beforehand. The initial profile is preserved with this limitation.

Four corrected two-second CPU profiles then use identical drivers on the actual frozen before/after bundles, select the action before the loop, and check small=1006 / large=1272. Before, RecordGet accounts for 52.63% of small's cumulative samples; after, it is absent from the displayed top with a 1% threshold. Small's visible residual costs include Apply (26.49% flat), FFI boxing (14.59% flat), the compact consumers and opaque barriers. After large, projection accounts for 21.76% cumulative, the grouped sum/guard 33.68%, and the update worker 7.77%. These are sampling shares of distinct processes, not nanoseconds or speedup ratios; Apply includes the driver boundary and saved-action replay. Raw pprof files, drivers, executable hashes and top outputs are retained in probe-baseline-corrected and probe-residual.

Two other paths were evaluated and not implemented. The actual traverse action allocates 8192 bytes once and spends time in its checksum. A four-accumulator probe initially failed Go's inline budget and gave a misleading slowdown; the corrected closure-shaped version showed about 7% at 2000/10000 elements, with 1000 variable. Both diagnostics are preserved. Its 740 distinct length/value variants, repeated twice, are not production validation. A balanced native record sum only changed medium 20.884 → 20.492 ns and large 36.928 → 36.813 ns; this was also rejected. No traverse, map, append or balanced-sum change is included.

## Compiler change and generated scope

The new RecordShapeGuard.integerReads helper shares complete shape checks between GoConversions and RecordReadGroups. For compact layouts 1–5, it verifies the tag, payload pointer, key pointer and every ordered label before positional reads. The two-field layout uses RecordFields2, supporting both boxed and integer payloads. Larger records retain their existing key/value lengths and ordered-label checks. Each consumption checks its own current shape; no guard or field is cached across an unknown call or another invocation, and the compiler adds no mutation. Nonmatching shapes and representations retain the original ordered fallback lookups.

GoConversions admits nonempty all-Int64 projections; mixed/composite projections keep their original conversion. RecordReadGroups extends its existing closed scalar regions to compact layouts. Only connected IntAdd subtrees in an admitted grouped-read fast path become native addition with a final int64(int32(...)); other arithmetic boundaries retain their semantics. Constant-only addition islands keep runtime calls: Go rejects an overflowing constant cast instead of applying runtime wrapping. Two generated/executed cases cover this issue. The entire fallback expression remains exact. The three compiler modules changed/added are GoConversions.purs, RecordReadGroups.purs and RecordShapeGuard.purs.

Native bootstrap verifies 558 modules and 346,791 TAST types; frontend / generation / Go build take 37.7 / 195.7 / 50.5 s. All 201 frozen source hashes remain unchanged across bootstrap; the final compiler inventory has 204 source/binary hashes. The existing PrivateSTRegions.purs:266 shadowing warning has unchanged source provenance. Build2's incremental JS bundle has no warnings; the full native frontend still reports the pre-existing warning.

Exact generated inventories remain 633/633 core and 683/683 AppAll sections. Two core sections and six AppAll sections change; headers and standalone files, including runtime and every FFI definition, are byte-identical. Four selected actions have executable changes: the three library records and core Deep Record Updates. The latter gains a compact e/f projection when consuming initial.b.d after the unchanged opaque counter, before the original native worker. Control_Bind changes only parentheses around indexed values. No section is added, removed or moved.

The actual-record review verifies both opaque calls, saved input action/replay, exact ordered fallback, unchanged update/sum workers and the opaqueValue noinline barrier. The generated runBench, calibrate, native measureBatch and callback adapter are byte-identical. Inputs, including the counter 1000, still pass through the same runtime opaque calls. All other selected actions, including traverse/map/append, are unchanged.

## Paired measurements

| Case | Official Go baseline (ns) | Paired before (ns) | After (ns) | Speedup | Go/C after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `record/set-small-1000` | 19.651 | 19.884 | 10.887 | 1.83× | 9.691386× |
| `record/set-medium-1000` | 23.482 | 23.589 | 20.794 | 1.13× | 18.524156× |
| `record/set-large-1000` | 41.110 | 42.275 | 37.383 | 1.13× | 27.568469× |
| `Deep Record Updates` | 21.906 | 22.133 | 16.391 | 1.35× | 0.005012× |

All four indicators use the same 37 Go/C pairs: the primary 14 App core medians and 23 AppAll library medians. The C denominator remains 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.074409× | 1.100262× | 1.101399× |
| Arithmetic mean μ | 7.792060× | 7.888220× | 7.522089× |
| Median Q₂ | 4.447449× | 4.484382× | 4.514085× |
| Geometric mean G | 4.039488× | 4.092295× | 3.968849× |

The four changed actions improve by 45.25%, 11.85%, 11.57% and 25.94% respectively. The paired arithmetic mean decreases 4.64% and geometric mean 3.02%. The total ratio increases 0.103% and median increases 0.662%; both increases are retained. Go total is 9.657378738 → 9.667362047 ms (+9.983309 µs). Unchanged Red-Black Tree alone increases 10.979500 µs; the remaining cases together decrease 0.996191 µs. Thus this campaign does not claim an overall elapsed-time improvement.

Relative to the official historical baseline, the final total is 2.512% higher and median 1.498% higher, while arithmetic mean is 3.465% lower and geometric mean 1.749% lower. This historical comparison is separate from the fresh paired gains. The largest remaining Go/C ratios are large records 27.57×, successful traverse 27.34×, append 21.89×, map 20.12× and medium records 18.52×. No improvement of unchanged actions is attributed to this code.

## Protocol and regressions

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, batches calibrated to at least 10 ms, and minimum of ten batches per process. Each published time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. Compilation, tests and profiles stop before official timing; the host is not isolated. Neither C nor JS is remeasured. All 306 measured outputs validate and source/binary guards pass around collection. Independent calculations reproduce all four indicators from the retained samples.

One primary increase exceeds 5%: foldable/traverse-maybe-fail-first-1000 rises 12.459 → 13.716 ns (+10.089%). Its generated action and harness are unchanged. A separate AB/BA/AB diagnostic reports 12.821 → 12.849 ns (+0.218%), with all six outputs validated. The increase over 5% does not reproduce there; this does not establish that the original difference was noise. The primary samples and all aggregates retain the original increase.

The secondary AppAll measurement of List Processing rises 4.740 → 5.014 ns (+5.781%). Its primary core measurement, used in the 37-case indicators and README, is 4.722 → 4.706 ns. Both sets are retained; the secondary value is not substituted into the primary results. No sample is removed or replaced by a diagnostic.

## Semantic validation and publication

Seventeen distinct Node unit tests pass: two projection tests, two grouped-read tests and thirteen related record/tuple tests. The first two suites contain respectively 239 and 870 generated-Go comparisons; those are checks within the tests, not additional distinct tests. They cover compact layouts, both two-field payloads, missing/nil/malformed or reordered representations, fallback errors/getters, mixed arithmetic, Int32 overflow and native wide-Int64 compatibility, alias mutation and temporary-name collisions. Constant-only addition islands retain runtime behavior.

Two source fixtures pass on both JS-host and native-host compilers, in Release and Dev: 37 projection observations and 25 grouped-read observations per profile. Generated snapshots assert actual fast-path activation, while source traces check evaluation order and effects. Both Go profiles match independently checked PureScript JavaScript output byte-for-byte. Wide/noncanonical Int64 and malformed-runtime checks are native compatibility cases, not claims about valid PureScript JavaScript inputs.

The three actual prepared record actions pass 1,134 immutable runtime controls per build through three invocation paths and five record representations. Saved actions replay changing runtime inputs; each preparation calls opaque once and each exercised action twice. Nonnegative bounded counts include 0, 1, 2, 7, 1000 and 65535. Negative countdowns are excluded because the unchanged fallback would execute billions of iterations.

Separate external-alias overlays verify 108 dictionary/map controls and 36 additional compact-two-field controls per build. The latter exercises the small action with both integer and boxed payloads; counts are 0/1/7. These overlays deliberately mutate through an alias inside the actual opaque count effect, while the compiler itself adds no mutation. The records must observe that mutation once at the original point of consumption. All three corpora have identical inputs, expected output and stdout before/after, with no stderr. The domains and invocation-path repetitions remain separately counted.

The archive retains baseline provenance, compiler/test changes, profiles, rejected probes, raw primary and diagnostic samples, manifests, commands and validation logs. Only existing Go Release README cells are updated; no table, reference column or status is added or changed. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-26-go-hotpaths.json) and var/benchmark/go-hotpaths-20261126.
