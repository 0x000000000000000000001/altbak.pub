# Go string checksum optimization

Baseline: clean compiler HEAD `541e68a24aba929b6183105398799104783d903c`. The campaign freezes 2,470 compiler source files, both previous measured Go executables, the official README and 733 protected benchmark/reference files, including archived reference documents, and 269 library/FFI source files. All 202 hashes of the compiler measured by the previous fold-arithmetic campaign match this starting state. The frozen C37 total remains 8.777346938669927 ms. Benchmark, harness and FFI sources and C/JS references are unchanged.

## Probe before implementation

The probe invokes the real frozen `Bench.StringFixtures.checksum` getter using the existing boxed runtime call. Its comparison separates conversion to a code-point array, primitive folding over that array, and a direct code-point traversal. Inputs include ASCII, non-ASCII BMP characters, supplementary characters, isolated surrogates encoded as WTF-8, surrogate pairs encoded as WTF-8, malformed bytes and replacement characters. A CLI seed supplies an extra suffix to each input. All 60 original/primitive/direct comparisons agree. The repetition count zero cases therefore contain the suffix; they are not empty-string tests.

For the Unicode fixture (1,500 code points plus the seed suffix), the original checksum measures 43.008 µs, 149,920 bytes and 3,007 allocations per operation. Primitive folding over the converted array takes 8.145 µs and one allocation; conversion alone takes 7.547 µs and one allocation. Direct traversal takes 3.280 µs with zero allocations. For 2,000 ASCII characters plus the suffix, original/direct take 51.157/2.399 µs with 4,007/0 allocations. These diagnostic timings determine the implementation direction; they are not official benchmark results or a predicted campaign speedup.

The source and observations are retained in `probe/main.go` and `probe/observations.txt`. No benchmark or reference was modified to obtain the probe.

A separate supplementary probe retains the first probe unchanged and compares direct decoding with an ASCII byte branch before the Unicode decoder. All 61 checks, now including an empty string, agree. Direct/ASCII-branch timings are 3.413/3.167 µs on the Unicode sample, 2.518/0.989 µs on the ASCII sample, 7.4/3.0 ns on the tiny sample, and 1.687/1.682 µs on the WTF-8 sample, with zero allocations for both variants. These are another diagnostic experiment, not paired official benchmark results. Sources, observations and provenance are retained separately in `probe-ascii/`.

## Compiler change

The new generic `CodePointFold` lowering recognizes an immediate primitive integer fold over `Data.String.CodePoints.toCodePointArray`, and the canonical nonrecursive composition of this conversion with a closed two-argument primitive callback and literal seed. The composition is exposed as a native `String -> Int` worker before worker discovery and cache keys. The immediate-call path captures callback, seed and string exactly once in their original order. Unknown, staged, recursive, incompatible or effectful callbacks and contradictory TAST annotations keep their original path. Unannotated captures remain conservative. The actual TAST `Prim.Function` application spelling is accepted as equivalent to the existing function representation; it is not type erasure.

The emitted loop handles ASCII bytes directly and calls the new general runtime decoder for other bytes. That decoder follows the unchanged FFI behavior for ordinary UTF-8, paired or isolated WTF-8 surrogates and malformed bytes. Arithmetic remains the existing callback lowering, including Int overflow and wide-value behavior. The optimization removes an immediately consumed fresh array from an immutable string; it neither mutates an aliased array nor relies on usage annotations as an ownership proof. Public conversion functions and unrecognized consumers remain unchanged. Implementation detail and guard evidence are preserved in `implementation.md`.

The final JS-host smoke compiles and executes the three unchanged benchmark preparations and actions: results are 64421000, 195000 and 295000. It verifies the actual generated checksum uses direct decoding with neither a code-point array conversion nor generic array fold, and verifies compiler source/bundle and protected-input hashes. The initial smoke was rejected for final validation: it exposed a missing TAST wrapper case, and its provenance guard detected source changes after the initial bundle. Its files are retained separately as preliminary evidence, not a successful final validation.

Native bootstrap succeeds with the same frozen compiler sources: frontend 41.4 s, generated Go 196.0 s, Go build 41.3 s; 557 modules and 346,688 TAST types verified. All 200 source hashes remain unchanged. The one existing `PrivateSTRegions.purs:266` shadowing warning is retained with an unchanged-source hash; the JS build has no warnings. Bootstrap commands, stage logs and provenance are archived.

The final production-source comparison finds all 633 core sections byte-identical. AppAll changes five sections: the shared checksum gains a native worker with a forwarding public getter, and the three string actions call that worker. Every generated header is unchanged. In each build, the sole remaining Go file delta is the new decoder helper in the runtime; removing that helper reproduces the previous runtime exactly. Both timed and calibration harness sections are byte-identical. The three action sources retain their opaque calls. Static absence of array allocation in this worker is checked separately from the measured probe allocations.

## Paired measurements

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup | Go/C after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `string/cp-take-unicode-1500` | 41.597004 | 41.432293 | 6.973104 | 5.94× | 10.81× |
| `string/concat-1000` | 54.605469 | 54.429688 | 10.281372 | 5.29× | 11.13× |
| `string/split-1000` | 123.478188 | 125.746422 | 10.453776 | 12.03× | 4.24× |

All figures below use the same 37 matched Go/C cases: the 14 primary App core medians and 23 AppAll library medians.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.082031× | 1.081318× | 1.075598× |
| Arithmetic mean μ | 13.781811× | 13.793293× | 9.718705× |
| Median Q₂ | 5.780516× | 5.785415× | 4.479247× |
| Geometric mean G | 5.134773× | 5.138753× | 4.374541× |

The measured Go total is 9.491100941 → 9.440900945 ms. The three string cases save 193.900151 µs together. The unchanged Red-Black Tree moves from 9123.812500 to 9267.375000 µs (+1.573%), offsetting most of that saving; the total decreases by 50.199996 µs. No Red-Black Tree change is attributed to this optimization. Differences from the historical baseline are likewise kept separate from the fresh paired gains.

Successful array traversal also measures 1.343318 → 1.149106 µs (-14.458%), despite byte-identical generated action and array-checksum code. That observed variation is retained in the aggregates but is not claimed as a string-checksum optimization gain. The causal performance claims concern the three changed string actions.

## Protocol and validation

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, calibrated batches of at least 10 ms and the minimum of ten batches per process. Published rows are medians of the three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before timing; the host is not isolated. C and JS references are not remeasured. All 306 measured outputs validate, and source/binary hashes are checked around collection. An independent calculation reproduces the four indicators from the retained samples.

No primary case or secondary AppAll case increases by more than 5%; no isolated regression remeasurement is needed. All primary samples are retained.

Validation passes 42 unique Node tests (13 targeted/related plus 29 other fold tests) and 10 runtime string tests. The 126 generated native/public-worker checks cover ABI and arithmetic but share the runtime decoder as an oracle; they are not independent Unicode evidence. Independent decoding evidence comes from the original frozen getter probes and the separate baseline/JavaScript corpus.

The final source fixture checks 49 observations per JS-host/native-host Release/Dev execution. Six eligible workers activate; unknown and unannotated captured callbacks retain the generic path. An imported eligible callback changed to an ineligible callback invalidates the caller correctly: incremental and clean outputs agree and the old native worker disappears. Final native and JS-host stdout are identical. The additional related-fold test log includes an earlier native compiler run before bootstrap; it is not substituted for these final native validations.

The actual three benchmark actions pass 1,215 varied-input observations before and 1,215 after, with exact input/expected/stdout equality, preserved opaque calls, immutable inputs and replay through three call paths. Each build includes 852 independent JS UTF-16 observations, 315 arbitrary-native-byte observations and 48 legacy native split observations. Malformed bytes and empty/nonboundary split patterns preserve baseline native behavior; they are not JavaScript-conformance claims or changes to the C reference domain. Empty strings, code-point boundaries, lone and paired surrogates, malformed sequences and Int32 checksum overflow are covered.

Sources, commands, raw measurements, manifests, compiler delta and validation logs are retained in the workspace and [the JSON archive](2026-11-26-go-string-checksum.json). Publication updates only existing Go Release cells in the README; the reference columns, statuses and table structure are preserved.
