# Go record update optimization

The existing `RecordAffineLoops` optimization already collapses the positive-count record-update loops into scalar arithmetic. The measured record-set bottleneck is repeated boxed-field lookup: large's action reads sixteen fields when converting its opaque input into a native record, then reads those sixteen fields again when consuming the unchanged original. `RecordGet` searches `RecordData.Keys` linearly.

This step changes only `GoConversions.unboxRecord`. For native target records with more than five fields, all represented by int64 and with unique labels, the conversion first checks that the boxed input is a non-null `TypeRecordData`, its complete key count matches, its value array is long enough, and every key matches the target field at that position. Only then are the integer payloads copied by position into the existing fresh native result. The whole shape check happens anew on each conversion.

TAST supplies the target record fields and scalar types. It does not establish the layout of an opaque or foreign value; the runtime guard proves that separately. Full unique-key agreement preserves first-match behavior for duplicate labels. Reordered, additional, missing or duplicate keys, maps, compact records, constructors, getters and other representations retain the original ordered `RecordGet` conversions. Composite or non-integer fields do not receive the guard. Integer payload reads cannot invoke a getter or mutate another field while copying.

The threshold follows `boxRecord`'s existing protocol: zero through five fields have compact layouts; larger records use `RecordData`. A compact-two-field diagnostic found a small extra-branch cost, so the compact path remains unchanged. No compiler usage annotation is used as an ownership proof. The source record is never mutated, no view is cached across calls, and no read is moved across a callback/effect. The second set of field reads remains unchanged and is a possible later optimization.

An isolated probe kept the actual emitted affine update worker and used an opaque runtime count and source values. Full-shape validation for the first conversion reduced the diagnostic action from approximately 388 ns to 182 ns. This identifies a plausible gain, not the official result; the benchmark campaign must compare against the frozen README baseline. Diagnostic slot-label-only variants were rejected because they do not preserve earlier duplicate keys.

Validation includes generated Go across 1/2/3/5/8/16-field records, canonical and reversed shapes, extra/duplicate/missing labels, short value arrays, compact/dynamic/foreign representations, constructor getter order, failures, one-time input evaluation and changed keys between calls. Related record regression tests passed 39/39. Benchmark code, FFI libraries, runtime implementation and C references are unchanged.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

This step optimizes conversion from boxed records to native structs. The separate field reads used to sum the original record remain unchanged and remain part of the measured work.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `record/set-small-1000` | 0.040 | 0.040 | 0.040 | 0.98× |
| `record/set-medium-1000` | 0.13 | 0.13 | 0.086 | 1.54× |
| `record/set-large-1000` | 0.36 | 0.37 | 0.21 | 1.76× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.11× | 1.10× | 1.10× |
| Arithmetic mean μ | 27.60× | 27.74× | 23.29× |
| Median Q₂ | 5.91× | 5.87× | 5.91× |
| Geometric mean G | 6.66× | 6.61× | 6.42× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-record-set.json) and `var/benchmark/go-record-set-20261126`.

## Regression review

No published case increased by more than 5% in the primary pair.

The medium and large record workloads improved by 1.5439× and 1.7641× respectively in the paired measurement. The small workload uses byte-identical generated code; its timing increased from 0.039537 to 0.040402 µs (+2.188%) and is retained as measured.

Across the 37 primary pairs, the total increased from 9.622728824 to 9.692447831 ms (+0.7245%), and the median Go/C ratio increased from 5.869695 to 5.908559 (+0.6621%). Neither indicator improved in this pair. The arithmetic mean decreased from 27.742516 to 23.291294, and the geometric mean from 6.608798 to 6.416601. These paired changes are distinguished from comparison with the official historical column; no overall measured speedup is claimed.

The core benchmark modules, harness and runtime are byte-identical. Core generated code as a whole is not identical: `Control_Bind.go` also gains the canonical projection for a seven-field date/time record (`day`, `hour`, `millisecond`, `minute`, `month`, `second`, `year`), with its previous fallback retained. AppAll adds that same change plus the medium and large record coercions. The observed small, total and median differences are retained without attributing their cause to the record-coercion change. No separate diagnostic timing was needed because no primary case increased by more than 5%.

Runtime input validation covers five record representations, including the exact boxed values produced by the original prepares, varied Int32 fields, repeated actions, and counts from 0 to 65,535. Input fields and keys remain unchanged. Negative counters are not claimed as exhaustively executed oracles: their original decrementing loops can require billions of iterations. The optimization changes record projection and preserves the existing update-loop behavior.

The compiler build retained one pre-existing warning in `PrivateSTRegions.purs` at line 266. The validation summary records the warning and verifies that this source file is unchanged from the frozen baseline.
