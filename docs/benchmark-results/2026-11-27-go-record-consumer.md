# Go immediate additive record consumption

Baseline: official concat-fold README, Go 37 **9.187287572 ms**, C 37 **8.777346938669927 ms**. Compiler HEAD `5a87a36a4a497b9a4c0d36448fc0e15d3af94e73` is clean after an external commit; all 212 measured source/binary hashes match the preceding publication. All 2,508 repository files, README, todo and executables are frozen. The 742 protected benchmark/reference files and 269 FFI/library files retain their hashes. No benchmark, harness, reference or runtime changes; no commit/push by this campaign.

The preceding three-process diagnostic found medium 16.06→8.037 ns and large 30.11→10.68 ns, with 2,016 bounded copied-action checks. Those diagnostic numbers are separate from the actual-build measurements below.

RecordConsumer derives temporary intramodule plans from certified additive recurrences. It recognizes immediate, complete IntAdd consumers of the updated and original records. For counts 1..Int32Max and a fully checked current shape, it reads fields once and computes the equivalent scalar result modulo 32. Unknown, captured, partial, contradictory, recursive or nonmatching forms retain the original complete expression. Effects and inputs are evaluated in their original order; no input mutation, ownership inference, public ABI or persistent cache descriptor is introduced.

Six files change: three compiler modules and three tests/fixture files. The final snapshot contains 2,512 files and 213 compiler source/binary hashes. Bootstrap validates 566 TAST modules/351,983 types; stages are 3.9s JS,51.8s frontend,204.4s generation,53.6s Go. The existing PrivateSTRegions warning is unchanged. To restore 8.30GiB free, three old campaign-owned bootstrap directories were compressed, verified entry-by-entry and then removed only from temporary storage;22,143 entries remain recoverable in separately hashed tar.gz archives. Sources, modules and proofs are preserved.

## Validation

Four unit entries include 11,340 generated Go comparisons. Source E2E passes 44 observations in Release and Dev on isolated JS and final main native hosts; a delta 1→2 transition keeps ABI and gives identical warm/clean Release results. Native hashes and 12 phase logs are retained.

Actual saved public getters receive 672 dynamic medium/large variants through Apply, the official FFI callback adaptation and forced fallback:2,016 observations per build, with one preparation and two invocation opaque calls. Three replays preserve captured inputs and immutability. There are 864 independent modulo 32 oracle results and 1,152 malformed-record panics. The eight bounded counts exclude−1 and nonzero values with low 32 bits zero, whose original fallback may run billions of iterations; native-wide inputs do not extend the C comparison domain.

Inputs and all results agree before/after. Exactly 576 panic rows differ only in process addresses: normalization replaces solely `(?<=UnsafePtr:)0x[0-9a-f]+`; every other character agrees. Raw outputs remain archived. Before stderr contains one Go dependency-download message; after stderr is empty. Separate actual opaque-boundary alias proofs pass 108 dictionary/raw-map plus 36 compact-small observations per build, including external mutation. Historical executions are not recounted.

Generated core 633/633 sections are identical. Exactly 3/683 library sections change: the three saved record actions. Each adds only a guarded scalar branch after the opaque calls; the complete original body remains byte-identical. Preparation, workers, registrations, harness, headers, runtime, FFI and initialization are unchanged. Post-timing allocation checks verify 1,002 returns per actual case/build: all three remain 0 allocations.

## Measurements

| Case | Paired before (ns) | After (ns) | Change | Go/C after |
| --- | ---: | ---: | ---: | ---: |
| record/set-small-1000 | 9.357 | 8.506 | -9.09% | 7.572× |
| record/set-medium-1000 | 19.381 | 9.768 | -49.60% | 8.702× |
| record/set-large-1000 | 33.669 | 12.064 | -64.17% | 8.897× |

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Total | 1.046704× | 1.089758× | 1.090567× |
| Arithmetic μ | 5.001127× | 5.219632× | 4.537797× |
| Median Q₂ | 4.077025× | 4.279354× | 4.338441× |
| Geometric G | 2.290405× | 2.393997× | 2.285193× |

Go 37 is 9.565181858→9.572284360ms: total **+0.074%**, arithmetic **-13.063%**, median **+1.381%**, geometric **-4.545%**. All primary values are retained. The three targets save 0.032069µs; unchanged RBTree moves 9467.000→9474.083µs (+7.083µs), so no overall elapsed-time gain is claimed. Against the official README, total/μ/Q₂/G change **+4.191%/-9.265%/+6.412%/-0.228%**; these historical movements are distinct from paired gains. No individual primary increase exceeds 5%, so no extra regression timing is run.

The same 37 C pairs use App 14+AppAll 23 medians from 12 processes in AB/BA/AB order. Each process uses three global/local warm-ups, calibrated≥10ms batches and≥10 batches; reported times are medians of three process minima. Release, GOGC 800, default GOMAXPROCS, noPGO; all other builds/tests/profiles/reviews stop during timing. Host isolation is not claimed. All 306 outputs and source guards pass; C/JS are not remeasured.

Only existing Go Release README cells are updated. [The JSON archive](2026-11-27-go-record-consumer.json) retains all samples, transfers, proofs, source snapshots and logs; working files remain under `var/benchmark/go-record-consumer-20261127`.
