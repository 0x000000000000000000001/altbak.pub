# Go immediate effects, bounded codepoint folds, sieve and scalar recurrences

The baseline is the official traverse/records measured working tree on compiler HEAD `62f362fbd2dc90a4ac86c0423d21a3424deaeaf4`. The snapshot copies all 2,492 tracked/untracked repository files, verifies all 208 measured source/binary hashes against the preceding archive, and preserves its final executables, README, direct todo.before copy and existing dirty publications. New changes are compared with that measured working tree, not HEAD. No commit or push is made.

Historical Go37 is 9.264692240 ms. Its four Go/C indicators are 1.055523076020044× / 5.904443281100011× / 4.505682136354958× / 3.450820746379006× (total / arithmetic / median / geometric). C37 stays fixed at 8.777346938669927 ms. All 740 protected benchmark/reference files and 269 library/FFI sources retain their hashes. The preceding campaign’s complete replacement series is the baseline; its excluded contaminated series is never reused as an official measurement.

## Empirical selection and compiler changes

Four independently probed compiler changes are combined into one bootstrap. The diagnostic timings below describe their probes, not the final campaign.

The immediate-effect branch probe reduced the real Effect.Ref action from approximately 13.0 to 6.2 ns and from one 32-byte allocation to zero through the official callback adapter. The generic pass applies only where an Effect bind continuation already immediately executes its selected branch. Known effect leaves execute inside that selected block; unknown factories, callbacks and explicit EffectDefer retain construction followed by exactly one Apply. Conditions, short-circuiting, selected effects and exceptions keep their order. Saved branches outside that consumer are unchanged. No helper ABI or runtime implementation changes.

The bounded codepoint-fold probe compared actual FFI take plus checksum on 8,568 inputs, with three repetitions in one process: median 7264 → 2956 ns and two allocations → zero. CodePointTakeFold recognizes canonical `take count input` immediately consumed by a certified closed String-to-Int fold. It captures count then input even for nonpositive counts, scans at most the requested codepoints, and reuses the existing ASCII path and exact UTF8/WTF8 decoder. Original Int narrowing and Number multiplication semantics are retained. Unknown consumers, stored substrings, captures, effects and contradictory type annotations keep the original path; inputs remain read-only.

A distinct codePointFold descriptor records callback, seed and absence in imported metadata and emission keys. Worker codec v4 is required for typed reuse, while the physical cached fragment remains index 3. Older ownership metadata stays readable for migration. Descriptor closure and ABI are revalidated. Source cache transitions change callback and seed, then remove eligibility without changing the String-to-Int signature.

The sieve probe measured 532.9 → 211.0 ns in three processes, with one allocation → zero. Its 1,102 cases include 1,100 independent ordered-list oracles and two large prefixes compared with the old bitmap. The complete existing recognizer gains a path only for dynamic lower bound 2, under the existing upper-bound and canonical-seed guards. It skips even numbers, marks odd multiples starting at p² and uses a local 512-byte buffer for small prefixes. Larger prefixes may allocate. Other lower bounds retain the old general elimination algorithm: their selected heads can be composite, so this p² shortcut is not applied to them.

The scalar recurrence probe measured 59632 → 1.877 ns across three processes, with zero allocations and 1,032 comparisons. AffineLoops extends its existing count-down recurrence proof to the increment `counter mod d`, with a positive literal Int32 divisor and the same counter decrementing by one. For positive canonical counts and canonical Int32 seeds, q=n/d and r=n%d give q*d*(d−1)/2 + r*(r+1)/2, followed by the original Int32 reduction. Products fit Int64. Other inputs keep the complete original loop, including zero counts and wide native seeds. Strict global demands retain their established order. This computes algebra on opaque runtime inputs; it does not substitute a known benchmark input.

The final transfer has 38 files: 16 compiler modules and 22 test, fixture or cache-tooling files, including nine new files. Two compiler modules are new. The only overlap between candidate transfers is ListProgressionFold: CP.take inserts one metadata field without changing the delivered sieve body. The exhaustive transfer check verifies all 2,501 resulting repository files and preserves every prior delivered change. Runtime, FFI, benchmarks and C references are unchanged. Later concat and record-snapshot probes are outside this wave.

## Build, generated scope and correctness

The combined bootstrap verifies 563 TAST modules and 350,178 types: JS bundle 9.9 s, frontend 47.5 s, generation 193.4 s and Go build 45.8 s. Its 207 final source hashes and 210 source/binary hashes identify the measured compiler. The pre-existing PrivateSTRegions shadowing warning has unchanged source provenance. The 11.44 GiB available before bootstrap passed the 8 GiB minimum.

Six of 633 core sections and eight of 683 AppAll sections change. Headers, package placement, runtime, FFI, entry points, initialization order, Bench, AppLibrary, registrations and the timing/batching adapter are byte-identical. The four selected actions change as intended, preserving opaque inputs and saved-action replay. Ref’s cold write worker is byte-identical; its branch-level closure disappears. CP.take’s original checksum worker and FFI take are unchanged. Primes’ original general path and TCO’s original fallback remain in place.

Four additional generic helpers change through immediate-effect branch lowering: monadRecEffect, monadRecST, ST.modify and Iterator.pushWhile. Their callback/state ordering, short-circuit conditions, mutation and pattern-match failures are preserved while immediately executed branch closures disappear. These are executable helper changes, not formatting. The bounded review records their bodies and does not infer that every measured variation comes from them.

The combined JS bundle passes 31 affine/sieve test entries. The sieve tests include 1,085 generated-Go oracle comparisons and allocation bounds; the recurrence tests add 546 generated-Go modulo cases and 42 rejected guard forms. Ref’s related and targeted suites pass 19 entries. CP.take’s related/cache suites pass 19 distinct entries, with 11,235 bounded generated-Go comparisons and 126 existing unbounded comparisons. Earlier repeated runs of these same tests are not added to the counts.

All four source fixtures pass on their isolated JS-host compilers and the final combined main native compiler. Sieve/TCO native tests contribute four Release/Dev entries and 13 phase logs. Ref compares 32 exact source-JS observations per profile in both hosts, including saved/replayed actions, factories and throwing conditions. CP.take checks 74 observations per ordinary profile and 256 for the traced unknown-callback profile, with three descriptor transitions in each of Release and Dev. Its six transitions retain ABI and compare warm with clean results; all 21 observable logs match across hosts. The manifests distinguish each compiler and source version.

The actual final executables receive new dynamic-input proofs in this wave:

- Sieve: 265 native ordered-elimination cases per build. TCO: 88 runtime JSON inputs through native and boxed workers per build, including terminating wide-native cases as a distinct domain. Two default opaque actions are also verified. Inputs and stdout are byte-identical across builds.
- Ref: 12 actions are prepared before any replay, then called four times each in reverse input order after the overlay input changes. All 48 observations per build agree exactly. The diagnostic overlay changes only the generated scalar input helper; public getters and official callback adaptation remain real. Positive counts at or above Int32max are not executed because the original cold write loop does not terminate there.
- CP.take: 2,772 input variants through Apply, actual callback conversion and forced generic fallback give 8,316 observations per build. The 1,008 UTF16/Node observations and 7,308 arbitrary-native-byte observations remain separate; native-only behavior does not extend the C comparison domain. Both preparation opaque calls and both action opaque calls are checked. Inputs, expected values, stdout and independent references match byte-for-byte before/after.

Source, executable and proof-artifact hashes pass for every corpus. These are current executions, not transferred historical test counts. Preserved diagnostic failures include two smoke-extractor mistakes, a corrected cache fragment-index migration typo before transfer, and a pre-existing unsafePerformEffect-in-pure-let ordering discrepancy. For that last diagnostic, old native, same-host hook-disabled and active outputs agree; the accepted saved-action fixture sequences construction explicitly in Effect and matches unmodified source JavaScript.

## Official paired measurements

| Case | Official baseline (µs) | Paired before (µs) | After (µs) | Paired change | Go/C before → after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `Prime Sieve` | 0.514225 | 0.510709 | 0.211056 | -58.674% | 4.474875× → 1.849290× |
| `Tail Call Optimization` | 60.095539 | 58.715496 | 0.003538 | -99.994% | 1.825500× → 0.000110× |
| `effect/ref-write-10000` | 0.012754 | 0.012772 | 0.006285 | -50.791% | 11.390527× → 5.605188× |
| `string/cp-take-unicode-1500` | 7.002970 | 7.009460 | 3.115051 | -55.559% | 10.869108× → 4.830304× |

Post-timing AllocsPerRun(1000) checks the actual four saved actions through the official FunctionData/Unbox-Apply callback adaptation, with original opaque calls and every result verified. Allocation counts are Primes 1→0, Ref 1→0, CP.take 2→0 and TCO 0→0. Its initial check plus warm-up yield 1,002 expected-result checks per fixed case/build. This allocation check temporarily uses one GOMAXPROCS and supplies no campaign timing or varied-input claim.

All four indicators use the same 37 Go/C pairs: primary App14 medians plus AppAll23 library medians. C37 stays 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.055523× | 1.070218× | 1.041812× |
| Arithmetic mean μ | 5.904443× | 5.877456× | 5.455220× |
| Median Q₂ | 4.505682× | 4.474875× | 4.310731× |
| Geometric mean G | 3.450821× | 3.439383× | 2.487120× |

Relative to the fresh pair, total decreases 2.654%, arithmetic mean 7.184%, median 3.668% and geometric mean 27.687%. Go37 is 9.393676240 → 9.144346472 ms. Relative to the official historical baseline, the decreases are 1.299%, 7.608%, 4.327% and 27.927%, respectively.

The four direct targets save 62.912507 µs in the pair; other cases save 186.417261 µs. Red-Black Tree alone moves 9225.9165 → 9040.125 µs (185.7915 µs), with its own generated section and runtime unchanged. That movement is not attributed to these target optimizations. Its historical time was 9095.6665 µs. The geometric-mean reduction is strongly influenced by TCO’s algebraic elimination and very small final ratio; it does not mean that all remaining cases approach 1×. The final 3.538 ns TCO action remains sensitive to call/harness overhead and is not a cost per original iteration.

The largest remaining ratios are large records 24.10×, medium records 16.84×, append 12.08×, successful traverse 11.36×, map 11.21× and concat 11.03×. No later candidate implementation or diagnostic result is included in these numbers.

## Retained regressions, protocol and publication

Three primary increases above 5% remain in the official samples, aggregate indicators and README:

| Case | Primary before → after (µs) | Primary change | Separate diagnostic before → after (µs) | Diagnostic change |
| --- | ---: | ---: | ---: | ---: |
| `st/ref-write-10000` | 0.006104 → 0.006951 | +13.876% | 0.006172 → 0.006352 | +2.916% |
| `Church Numerals` | 0.003827 → 0.004642 | +21.296% | 0.003836 → 0.003859 | +0.600% |
| `Deep Record Updates` | 0.016025 → 0.016963 | +5.853% | 0.016825 → 0.017145 | +1.902% |

Each diagnostic uses its own six-process AB/BA/AB series with the original target and timed harness verified unchanged. None reproduces a rise above 5%; this does not establish the cause of the full-suite increases or erase them. The 18 diagnostic outputs and all original samples are retained separately.

The official campaign has 12 processes: three per build per suite, AB/BA/AB order, three global and three local warm-ups, batches calibrated to at least 10 ms and at least ten batches per process. Each time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. Other builds, tests, profiles and computational reviews were stopped during official timings and diagnostics. The host is not isolated. C and JS were not remeasured. All 306 official expected outputs and source/binary guards pass, and independent recomputation reproduces all four indicators.

The archive preserves the dirty baseline, exact scoped transfers, source/test changes, candidate probes and their limits, bootstrap logs, cache/source fixtures, final dynamic-input proofs, allocation checks, generated-code review, official samples and separate diagnostics. Only existing Go Release README cells are updated; table structure, reference values and status labels are retained. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-27-go-effects-strings-sieve.json) and var/benchmark/go-effects-strings-sieve-20261127.
