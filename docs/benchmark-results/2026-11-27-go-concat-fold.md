# Go concat-fold fusion and independent mapped-fold accumulators

The baseline is the official effects/strings/sieve/recurrence measured working tree on compiler HEAD `62f362fbd2dc90a4ac86c0423d21a3424deaeaf4`. All 2,501 tracked/untracked repository files are copied, and all 210 measured source/binary hashes match the preceding publication. Its final executables, README, direct todo.before copy and existing dirty publications are retained. New changes are compared with that measured working tree, not HEAD. No commit or push is made.

Historical Go37 is 9.144346472 ms; Go/C indicators are 1.0418121256792527× / 5.455220132618879× / 4.310730630766185× / 2.4871203375367776× (total / arithmetic / median / geometric). C37 stays fixed at 8.777346938669927 ms. All 741 protected benchmark/reference files and 269 library/FFI sources retain their hashes.

## Empirical selection and delivered changes

A real concat-action profile attributes 89.09% cumulatively to ConcatStrings and 8.48% to checksum over 1.65 seconds of sampled CPU. The first ASCII candidate restarted the complete work on Unicode and regressed when Unicode appeared late. That candidate was rejected; its logs remain distinct. The selected suffix probe reports 10269.408 → 1211.136 ns across three paired short processes (8.48×), with one 2048-byte allocation → zero. These are diagnostic results, separate from the official campaign below.

StringConcatFold recognizes only a canonical saturated foldlArray/concatString producer immediately consumed by an existing certified closed codepoint-fold descriptor. Producer callback, initial string and input array are captured once in their original order. ASCII prefixes are folded directly. At the first non-ASCII byte, the current accumulator is retained, only the unconsumed suffix is concatenated with the original helper, and the existing UTF8/WTF8 decoder continues the fold. The consumed ASCII prefix cannot contain an unfinished encoding. Scalar arithmetic, noncommutative callbacks, empty/wide seeds and shared input arrays retain their semantics. Unknown, partial, captured, unsupported or contradictory forms keep the original path. No runtime decoder, metadata schema or public ABI changes.

The independent-accumulator probe measured map 4763 → 4103 ns and successful traverse 466 → 385.4 ns in three processes. Append regressed 1009 → 1145 ns, so it was explicitly excluded. IndexedAddLanes applies only after existing mapped-producer fusion when a closed mapper normalizes its result and the consumer is exactly IntAdd(accumulator, independent contribution). An Int annotation alone is insufficient; accumulator dependence inside Let/Branch and possibly partial arithmetic reject the transformation. Fewer than eight elements retain the exact original loop. Larger arrays use eight accumulators and combine with the existing IntAdd, preserving element/callback index order, multiplication behavior and a scalar tail. The bound uses index<=len−8. No input is mutated, and empty arrays preserve wide seeds.

Twelve files are transferred: seven compiler modules and five tests/fixtures, including seven new files. Complete verification checks all 2,508 final repository files and preserves prior delivered changes. The native bootstrap verifies 565 TAST modules and 350,870 types: JS bundle 3.4 s, frontend 47.1 s, generation 191.0 s and Go build 48.0 s. The final compiler is identified by 209 source hashes and 212 hashes including binaries. The existing PrivateSTRegions shadowing warning retains unchanged source provenance. Only the regenerable Go compilation cache was cleared before this build; sources, downloaded modules and frozen artifacts were preserved. Bootstrap began with 10.15 GiB free, above the 8 GiB minimum.

## Generated scope and validation

All 633 core sections are byte-identical. Exactly three of 683 AppAll sections change: concat, map and successful traverse, in the same shared package. Only their saved action bodies change. Preparation prefixes, opaque invocation lines, callbacks, headers, runtime, FFI, initialization, Bench, AppLibrary, registrations and timed adapter are unchanged. There are no other activations or package movements. Append and failing Maybe traversals remain byte-identical. The bounded generated-code review completed before timing.

The concat targeted/related suites pass 11 entries, including 13,254 generated native/public comparisons. A separate diagnostic oracle checks 83,676 cases against original ConcatStrings plus the library codepoint-array FFI. The nine lanes-related test entries include 2,760 generated native/boxed comparisons over short/boundary/large lengths, wide and empty seeds, exact IntMul, call order and immutability. The final main bundle passes all 20 combined test entries; earlier reruns of the same tests are not added again.

Concat source E2E passes initial Release and Dev on the isolated JS host and combined native host, with 21 observations per profile. A Release-only imported callback change (33→37) keeps ABI and gives identical warm/clean results. Twelve native phase logs and compiler hashes are retained. The lanes smoke compiles the three unchanged actual workloads in Release and Dev, compares their outputs with real source JavaScript, verifies activation for map/traverse and absence for append. Earlier test-harness failures (wrong Array String ABI and a curried/uncurried adapter mismatch) are retained; no production expectation was weakened.

Current before/after executables receive separate dynamic-input proofs:

- Concat: 268 fragment variants through Apply, actual FFI callback adaptation and forced fallback give 804 observations per build. The 76 UTF16 variants produce 228 Node-comparable observations; 192 raw UTF8/WTF8/malformed-byte partitions produce 576 native-only observations. The saved action is replayed with changing inputs, two preparation opaque calls and one per action. Its real initial string stays empty; general initial strings belong to the generator tests. The actual action has no native effect metadata, so no distinct native-metadata execution is claimed. The native oracle concatenates bytes and uses the unchanged FFI decoder with explicit Int32 checksum.
- Map and successful traverse each pass 159 canonical Int32/Node observations and 156 separate native-wide/BigInt observations per build: 630 observations per build across four corpora. Each keeps its actual normalizing mapper, checksum, preparation and action opacity. Shared inputs remain unchanged.

All before/after inputs, expected outputs and stdout agree byte-for-byte. Concat’s independent references also agree. Source, benchmark binary and proof-artifact hashes are verified. Native-only domains do not extend the C comparison domain. These are new executions on this wave’s builds; previous array/string proof executions are not recounted.

## Official paired measurements

| Case | Official baseline (µs) | Paired before (µs) | After (µs) | Paired change | Go/C before → after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `string/concat-1000` | 10.188802 | 10.162272 | 1.188082 | -88.309% | 10.999449× → 1.285957× |
| `array/map-10000` | 4.805847 | 4.827799 | 4.016551 | -16.804% | 11.258794× → 9.366902× |
| `foldable/traverse-maybe-1000` | 0.479764 | 0.466764 | 0.377965 | -19.024% | 11.054240× → 8.951238× |

Post-timing AllocsPerRun(1000) calls the actual saved actions through the official callback adaptation, keeps their opaque inputs and verifies every return. Concat changes from one allocation to zero; map and traverse remain at zero. One initial check plus warm-up yields 1,002 checks per fixed case/build. This check temporarily uses one GOMAXPROCS and supplies no campaign timing or varied-input claim.

All four indicators use the same 37 Go/C pairs: primary App14 medians plus AppAll23 library medians. C37 remains 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.041812× | 1.030995× | 1.046704× |
| Arithmetic mean μ | 5.455220× | 5.375267× | 5.001127× |
| Median Q₂ | 4.310731× | 4.259982× | 4.077025× |
| Geometric mean G | 2.487120× | 2.451312× | 2.290405× |

The arithmetic mean decreases 6.960%, median 4.295% and geometric mean 6.564% within the fresh pair. The total ratio increases 1.524% and is retained: Go37 is 9.049401312 → 9.187287572 ms. The three changed targets save 9.874237 µs, while the other 34 cases add 147.760497 µs. Red-Black Tree alone moves 8946.521 → 9093.458 µs (+146.937 µs) in the entirely unchanged core build. No causal performance gain or regression is assigned to its source from this wave, and no overall elapsed-time gain is claimed for this pair.

Against the preceding official README, total is 0.470% higher while arithmetic, median and geometric ratios are 8.324%, 5.421% and 7.909% lower. No individual primary increase exceeds 5%, so no extra regression diagnostic is run. Every primary result, including the higher total, is retained in the README and archive.

## Cumulative official comparison and remaining work

The principal cumulative baseline is the official README before the overnight Runtime Box campaign (the go-hotpaths archive). It is distinct from that first campaign’s freshly measured before series. Both origins and archive hashes are verified in proofs/night-baseline.json.

| Measure | Official README at start | Current official result | Change |
| --- | ---: | ---: | ---: |
| Go37 total | 9.667362047 ms | 9.187287572 ms | -4.966% |
| Ratio of totals | 1.101399× | 1.046704× | -4.966% |
| Arithmetic mean μ | 7.522089× | 5.001127× | -33.514% |
| Median Q₂ | 4.514085× | 4.077025× | -9.682% |
| Geometric mean G | 3.968849× | 2.290405× | -42.290% |

The supplementary first measured before series was 9.699130421 ms with ratios 1.1050184627280728 / 7.584230129940671 / 4.5782497078524385 / 3.999178857649216. It is not described as the earlier official README. C is unchanged throughout. Cumulative movement includes host variation and effects of different binaries; causal claims remain tied to the separately retained paired campaigns. In particular, the previous TCO algebraic reduction strongly influences the geometric mean, and unchanged-tree movements are not attributed to unrelated optimizations.

The largest remaining current ratios are large records 23.76×, medium records 16.28×, append 12.10×, map 9.37× and successful traverse 8.95×. A separate record-consumer diagnostic is retained as future work: three short processes report medium 16.06 → 8.037 ns and large 30.11 → 10.68 ns, with zero allocations and 2,016 bounded correctness observations. An initial negative-count case reached the unchanged fallback’s long loop and was stopped at a 25-second timeout; that log is preserved and the corrected bounded corpus is identified. This is a diagnostic copy and design proposal, with no production compiler change or official ratio from it. General admission, alias/opaque boundaries, strict evaluation and cache/public-ABI preservation still require their own implementation proof.

## Protocol and publication

Twelve official processes, three per build per suite in AB/BA/AB order; three global and three local warm-ups, calibrated batches of at least 10 ms and at least ten batches per process. Each time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. Other builds, tests, profiles and computational reviews stop during all official timings. The host is not isolated. C and JS are not remeasured. All 306 expected outputs and source/binary guards pass. Independent recomputation reproduces the four indicators.

The archive preserves dirty-baseline provenance, exact transfers, compiler/test changes, accepted and rejected probes, bootstrap logs, source/cache/native fixtures, current action corpora, allocation checks, generated-code review, all official samples and the separately labelled future record probe. Only existing Go Release README cells are updated; table structure, references and status labels are unchanged. The removed ratio catalogue remains absent. No next campaign or production record change is started. See [the JSON archive](2026-11-27-go-concat-fold.json) and var/benchmark/go-concat-fold-20261127.
