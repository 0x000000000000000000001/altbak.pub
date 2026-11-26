# Go fold arithmetic optimization

Baseline: clean compiler HEAD `c2c64c317f55919721c0b6ba84e8a50119c21cd2`. Campaign preserved the baseline and verified the 202 prior source/binary hashes. Only `PrimitiveExprs.purs` and its selection in `CodeGen.purs` change compiler behavior. Benchmark and harness sources, references, runtime and FFI are unchanged. The existing fold/callback machinery is retained; its emitted arithmetic changes.

The diagnostic copied the real native weightedChecksum worker into `probe/actual.go`, with opaque CLI seed and noinline only on outer measured entries. Variants separated callback integration, modulo reduction, a guarded multiply helper, and an immediately invoked guarded expression. All 224 checks matched the original runtime, across lengths 0/1/2/3/17/1000/10000 and signed32, unsigned32, small-positive, and noncanonical int64 inputs. All variants allocate zero. `observations.txt` retains the broader first measurements; `observations-final.txt` records the final IIFE variant. These are diagnostic timings, not official campaign results.

For 10,000 small-positive values the final probe measured the actual worker at 26.115 us, callback integration alone at 24.241 us, modulo alone at 23.277 us, named guarded multiplication at 8.150 us, and the IIFE guard plus modulo reduction at 4.896 us. With varied signed32 values, actual/IIFE were 37.022/4.720 us. Compiler inlining logs and `assembly-final.txt` show that the original callback is already integrated. The remaining expensive call is math.mod. The named multiply helper failed the ordinary inline budget, while the immediately invoked guard was inlined; this determined the implementation shape. No fusion, callback rewrite, or global runtime multiplication replacement was needed.

The existing static `boundedIntProduct` remains first. Its syntax proof now learns the bound of Euclidean modulo by a canonical positive literal and sharpens addition/subtraction of known bounds, capped at the existing Int32 normalization bound. Thus 1+(index mod2) has absolute bound2. Explicit incompatible annotations and unknown local bounds remain unknown.

The second path admits a factor whose absolute magnitude is proved at most 2^b, with b<=21. An IIFE evaluates both operands once in source order, including each operand's statements. It checks the other full int64 value against the inclusive range [-2^(53-b),2^(53-b)]. Inside that range both Number inputs and their product are exact integers within ±2^53, and the int64 product cannot overflow. Narrowing the exact product to int32 therefore matches the original JS multiplication followed by ToInt32. Outside the guard, the original runtime IntMul remains. The threshold21 ensures the guard's range includes all ordinary signed32/unsigned32 values; it is not a claim that all typed Int values are Int32. Zero-fill shifts, the division boundary2^31, and arbitrary noncanonical int64 values remain supported.

The modulo path replaces a canonical positive power-of-two literal divisor with Go's raw int64 low-bit mask. Euclidean modulo has the same result for negative and positive inputs. This is not PureScript's bitwise helper, which also normalizes to int32. Literal bounds are checked before both the 32-bit power-of-two test and the compiler's subtraction: 2^32,2^32+2 and2^53 are rejected. Zero, negative, non-power-of-two and unknown divisors retain their original operation.

Tests distinguish static exact products, guarded products and retained fallback. They cover factor magnitude21 and outside, swapped/signed factors, inclusive and adjacent ±2^52/±2^53 boundaries, unsigned32 and int64 extrema, noncanonical modulus literals, strict left/right evaluation and exceptions on either side. The actual unchanged optimized weightedChecksum IR is emitted in `probe-real-ir`; 48 compiled native/public wrapper checks passed with guard±2^52, mask1, and IntMul only as fallback. Independent source, frontend and runtime proofs and the complete before/after campaign are archived separately. The two related suites contain 66 unique tests (35 + 31); the 24 primitive tests are included in that total.

Before native bootstrap, the related validation passed 66 unique tests (35 related primitive/fold/array/record tests plus 31 integer/affine/range regressions). The earlier 24 primitive tests are a subset of those 66, not additional tests. The real native/public checksum additionally passed 48 compiled checks. The final three-workload smoke passed and checked the exact inclusive guard, raw modulo mask, sole fallback IntMul, and preserved native-array paths.

Native bootstrap succeeded on the unchanged frozen source: frontend 37.1 s, generated Go 186.0 s, go build 43.4 s. Native E2E, runtime proofs and the complete benchmark campaign use this build. No official performance claim is inferred from the isolated probe.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `array/map-10000` | 28.86 | 26.19 | 8.60 | 3.04× |
| `array/append-1000` | 5.94 | 5.31 | 1.86 | 2.86× |
| `foldable/traverse-maybe-1000` | 3.10 | 2.90 | 1.35 | 2.14× |
| `foldable/traverse-maybe-fail-first-1000` | 0.034 | 0.033 | 0.033 | 1.00× |
| `foldable/traverse-maybe-fail-last-1000` | 0.80 | 0.76 | 0.76 | 0.99× |
| `array/filter-10000` | 24.30 | 21.56 | 11.62 | 1.86× |
| `array/sort-descending-10000` | 51.70 | 47.43 | 29.75 | 1.59× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.18× | 1.10× | 1.08× |
| Arithmetic mean μ | 18.58× | 17.33× | 13.78× |
| Median Q₂ | 5.93× | 5.81× | 5.78× |
| Geometric mean G | 6.17× | 5.80× | 5.13× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-fold-arithmetic.json) and `var/benchmark/go-fold-arithmetic-20261126`.

## Regression review

The primary 37-case combination (App14 and AppAll23) contains these increases above 5%; they are retained in the measurements:

- `AST Evaluation`: +5.77%.

The primary AST Evaluation median increased from 0.027097 to 0.028660 µs (+5.768%). One separate AB/BA/AB probe using the exact generated before/after action and original harness measured 0.029031 → 0.029059 µs (+0.0964%): the increase above 5% was not reproduced in that isolated probe. These are distinct probe executables; the primary samples and published AST value remain unchanged. This does not prove that the primary difference was noise.

The paired total decreases from 9.659525102 to 9.497361368 ms (−1.679%), the arithmetic mean by 20.467%, the median by 0.523% and the geometric mean by 11.492%. The total falls by 162.164 µs, of which 107.334 µs (66.19%) comes from Red-Black Tree, whose generated module is byte-identical and does not use WorkloadMap. No Red-Black Tree acceleration is attributed to this change. The unchanged heavy cases also vary, so the total reduction cannot be attributed entirely to the optimized fold. Map improves by 3.044×, append by 2.861×, successful traversal by 2.142×, filter by 1.856× and sort by 1.594×. Fail-first changes by −0.465% and fail-last by +1.307%; their failure branches do not execute the checksum. ST.Ref/Effect.Ref modify improve by about 1.26× after the same generic arithmetic lowering.

The exhaustive comparison records 12 changed core Go sections and 17 AppAll sections. Changes extend beyond the shared checksum to parity predicates, bounded multiplications in array/list/map workloads, date helpers, code-point construction, reference-modify expressions and calibration. Headers, runtime and other generated Go files are byte-identical. The source harness remains unchanged. Its generated calibrate function uses the same guarded lowering for iterations*2 after measureBatch returns; runBench and the timed measureBatch implementation are byte-identical. The isolated AST proof checks this distinction explicitly.

Validation includes 66 unique tests (35 related plus 31 integer regressions; the 24 primitive tests are a subset), 48 checks of the emitted real worker/getter, and 224 diagnostic-probe oracles. The source fixture checks 37 observations per JS/native Release/Dev profile. Five real prepare/action workloads pass 471 varied-input observations before and 471 after. The distinct wide-int64 worker/getter corpus passes 128 observations and 482 prefix results per build: 256 observations and 964 prefix results across both builds. Both corpora preserve exact input/output bytes, immutable inputs and replay; the wide oracle explicitly converts decimal-string BigInt inputs to Number before multiplication. These wide checks do not change the frozen C32 benchmark references.

The previous unsequenced factory-order diagnostic remains archived with its original source, outputs and provenance. It is a known prior JS discrepancy outside this arithmetic change, not a JavaScript-conformance pass. The current arithmetic fixture tests explicit operand order and exception priority separately. The pre-existing PrivateSTRegions.purs:266 warning is retained with unchanged-source provenance. The host is not isolated, and all measurements remain archived without replacement.
