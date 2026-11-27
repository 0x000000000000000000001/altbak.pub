# Go total-traverse fusion and native record fast wrappers

The baseline is the complete measured Box + array-fusion + Tree/Sum working tree on compiler HEAD `62f362fbd2dc90a4ac86c0423d21a3424deaeaf4`. The snapshot preserves all tracked/untracked changes, 2,483 source files, all 207 source/binary hashes matching the previous archive, previous after-core/after-all executables, official README and a direct todo.before copy. Existing benchmark documents and the initial dirty diff are retained. New changes are compared with this frozen working tree, not HEAD. No commit or push is made.

The official 37-case Go total is 9.335966958 ms. Historical Go/C indicators are 1.063643379171× / 6.555264143569× / 4.546005242372× / 3.620462001691× (ratio of totals / arithmetic mean / median / geometric mean). This baseline retains the preceding full-suite fail-last traverse result rather than substituting its isolated diagnostic. C37 remains fixed at 8.777346938669927 ms. All 739 protected benchmark/reference files and 269 library/FFI sources retain their hashes.

## Empirical selection and compiler changes

An earlier controlled-copy probe of the actual record actions measured median decreases of 9.13% for medium records and 12.75% for large records, with zero allocations in both variants. The generated native wrapper initially exceeded Go's inlining budget (85, then 81 units). An equivalent unsigned interval test brings the actual generated public workers to cost 80 and makes them inline in their actual consumer actions. The isolated real-source smoke verifies that inlining, 198 wrapper/cold-worker comparisons and the prepared-action results 1072/1272. These worktree diagnostics use an older runtime baseline and are not substituted for the current official measurements.

The existing complete affine-record recognizer now permits a small public native fast path with a typed cold helper. Eligibility requires a self-recursive module worker with a native record result, one direct Int-field update and a constant delta of −1, 0 or 1. Field name, record size, module name and input values do not determine admission. Nested/multiple updates, other deltas, unknown calls, captured increments, unmatched recursion and helper-name collisions preserve the original generation.

The public native ABI and boxed wrapper remain unchanged. The cold helper contains the complete original emitted body. The fast interval is `0 < n && n <= 2147483647`, expressed equivalently as `uint64(n)-1 < 2147483647`. Within this interval, one final Int32 reduction preserves the original repeated modulo-32 addition, including wide native seeds and delta zero. Count zero keeps the cold body's unchanged wide value; all other out-of-range counts retain the original loop. Only a value-copy of the record parameter changes; borrowed fields/buffers and caller argument evaluation remain intact.

The traverse candidate follows the earlier isolated producer-fold probe (1214 → 474.8 ns, 8192 bytes → 0). That historical diagnostic suggested the opportunity but is not a new measurement of this implementation. Current evidence includes a real optimized-IR dump and an unchanged benchmark smoke returning 752500 with explicit activation and no intermediate array.

ArrayTraverseFold recognizes a canonical Array/Maybe traverse immediately consumed by a certified indexed integer fold. The adjacent callback must directly construct canonical Just with a total closed primitive Int payload; indexed callbacks are also supported. Captures, unknown calls, Nothing, mapper branches, custom dictionaries and uncurried/staged applications retain the original path. Each typed layer is checked, including consistent specialization of repeated type variables.

The consuming Let must have the canonical tag tests and failure branch, with its Just branch folding field zero of the same binding. An intervening Let/effect blocks the fusion. EffectPure is admitted only within an already executing effect block, preserving when computation occurs. Arguments and dictionaries are captured in their original order before any buffer read. The closed mapper executes before the fold callback for each index; source buffers are read-only. Existing integer narrowing, Number multiplication and imported-local renaming are retained. The existing arrayFold metadata and cache-v3 descriptor/absence keys are reused; no new schema is introduced.

Exactly 13 source/test/fixture files are transferred from the isolated worktrees: five compiler sources, four test files and four fixtures, including nine new files. Independent transfer checks verify all other baseline source files, including Box, array fusion and Tree/Sum. No runtime, FFI, benchmark or C reference is transferred or modified. The separate Effect.Ref and take/checksum investigations are outside this wave.

## Build, scope and correctness

The native bootstrap checks 561 TAST modules / 349,468 types: frontend 42.1 s, generation 193.2 s and Go compilation 44.6 s. All 205 final source hashes remain fixed; 208 source/binary hashes identify the measured compiler. The pre-existing PrivateSTRegions shadowing warning retains unchanged source provenance. Disk space was checked before bootstrap (13.41 GiB, above the 8 GiB minimum).

All 633 core sections are byte-identical. Seven of the 683 AppAll sections change, with unchanged headers, runtime, FFI, entry points, initialization order and timing/batching harness. The three record workers gain a fast path and typed cold helper. Nine bounded function comparisons verify three preserved preparation bodies, three original cold bodies and three unchanged native ABIs, normalizing only function names and generated package/runtime aliases. Original opaque calls, saved-action capture/replay, projection guards, fallback lookup, sums and public boxing remain in place.

Package placement changes: the medium record's getter, preparation and fast worker move to the shared package while its cold helper stays in p_root. The small record's fast worker/prepared action remain shared, with its cold helper in p_root. The large record remains in p_root. AppLibrary changes only the qualification of the medium preparation getter; registrations, labels and order are unchanged. These layout changes can influence other timings, so an unchanged action is not assumed immune to full-binary context changes.

The successful traverse changes inside its timed saved action. Its opaque input is still read on every invocation, followed by the same ordered dictionary/mapper/input captures. The intermediate array, always-Just result and consuming checksum become one loop with the same mapper and checksum arithmetic. Failure-traverse action bodies remain unchanged; Nothing-producing and unknown callbacks do not gain a new shortcut. The four directly changed measured actions are successful traverse and the three record sizes.

The main combined JS bundle passes six focused Node entries: three record-wrapper tests with 10,314 generated-Go checks and three traverse tests with 1,480 checks. The earlier isolated runs of those same tests are not counted again. Separate related suites report 12 passing entries each. Record checks include delta −1/0/1, zero/wide counts and seeds, cold fallback, name collisions, value-copy behavior and borrowed-field identity. Traverse checks include canonical/dictionary admission, evaluation order, alias mutation before buffer reading, imported locals and wide-value multiplication.

The source fixtures pass on their isolated JS-host compilers and on the final combined main native compiler. The preliminary worktrees use their older runtime; the final native manifests identify the measured main version. Record source behavior is checked in Release and Dev; its increment1→2 warm/clean cache transition and obsolete-helper removal are checked in Release. Traverse checks three ABI-preserving callback/seed/eligibility transitions in each of Release and Dev: six transitions total, with 19 observations in its initial profile and 60 in the traced unknown-callback profile. Its 21 observable logs agree between JS-host and final native-host runs. Twelve final native record phase logs are retained.

The actual record actions are rerun before and after on this wave's executables: 1,134 immutable-input observations, 108 intentional external-alias observations and 36 separate compact-record alias observations per build. These reuse proven scripts, not historical executions. They exercise Apply, the actual FFI callback conversion and forced fallback. Inputs, expected values and stdout agree byte-for-byte across builds; source, binary and artifact hashes are checked.

The actual successful-traverse action similarly passes 159 canonical Int32 observations (53 inputs, independent JavaScript Number/Int32 oracle) and a distinct 156 native int64 observations (52 inputs, BigInt oracle) per build. Its real seed is zero; general seeds and fallback callbacks belong to the compiler fixtures. It retains two preparation opaque calls and one per action. Both domains preserve identical before/after inputs, expected values and stdout. Diagnostic overlays change the supplied inputs only; production FFI and benchmark sources remain unchanged.

## Measurement incident and complete replacement

During the first 12-process measurement series, a sibling source-review script unintentionally performed quadratic work by repeatedly cleaning the remaining Go file for every function. Its Python process consumed CPU for an estimated 30–50 seconds. The exact overlap was not timestamped. The whole first series is therefore excluded, rather than selecting or correcting individual samples. Its 37 raw files remain byte-identical under campaign-contaminated, with original-to-retained path relocation and hashes documented in measurement-contamination.json. Correctness checks remain valid; those timings are not used below or in the README.

After termination of that process and confirmation that all agents had stopped computational work, the entire same 12-process AB/BA/AB campaign was repeated on the unchanged executables. Only this replacement is official. The faulty script is reconstructed from its recorded tool-call source; no original launch hash exists, so no byte-for-byte attestation of that earlier script is claimed. The corrected bounded nine-function review ran after the complete timing window. Both versions and incident metadata are retained.

## Official paired measurements

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Paired change | Go/C before → after |
| --- | ---: | ---: | ---: | ---: | ---: |
| `foldable/traverse-maybe-1000` | 1.187042 | 1.177770 | 0.479253 | -59.308% | 27.8928× → 11.3500× |
| `record/set-small-1000` | 0.009515 | 0.009597 | 0.009175 | -4.397% | 8.5431× → 8.1674× |
| `record/set-medium-1000` | 0.019712 | 0.020101 | 0.019089 | -5.035% | 17.9068× → 17.0053× |
| `record/set-large-1000` | 0.037540 | 0.037071 | 0.032764 | -11.618% | 27.3384× → 24.1621× |

A separate post-timing check calls the actual saved actions through runtime.Apply with their original opaque inputs and checks every result. AllocsPerRun records traverse changing from 1 to 0 allocations per action; all three record actions remain at 0. It performs 1,000 iterations plus a warm-up and initial check (1,002 checks per fixed case/build), temporarily using one GOMAXPROCS as specified by AllocsPerRun. These are four repeated expected cases, not a varied-input corpus. No timing from this check replaces the campaign.

All four indicators use the same 37 Go/C pairs: primary App14 medians and AppAll23 library medians. C37 remains 8.777346938669927 ms.

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.063643× | 1.055504× | 1.055523× |
| Arithmetic mean μ | 6.555264× | 6.496767× | 5.904443× |
| Median Q₂ | 4.546005× | 4.515829× | 4.505682× |
| Geometric mean G | 3.620462× | 3.584361× | 3.450821× |

Relative to the fresh paired baseline, arithmetic mean decreases 9.117%, median 0.225% and geometric mean 3.726%. The total ratio increases 0.00185% and is retained: Go total is 9.264520550 → 9.264692240 ms (+0.171690 µs). The four directly changed targets save 0.704258 µs; the other 33 cases together add 0.875948 µs. No overall elapsed-time improvement is claimed for this paired series.

Against the official historical baseline, total is 0.763% lower, arithmetic mean 9.928% lower, median 0.887% lower and geometric mean 4.686% lower. Most of the historical elapsed-time decrease comes from Red-Black Tree (9165.4375 → 9095.6665 µs), whose emitted core code and runtime are unchanged here. That historical movement is not attributed to these optimizations. Its paired decrease is only 0.917 µs. The largest remaining Go/C ratios are large records 24.16×, medium records 17.01×, append 12.35×, map 11.39×, Effect.Ref write 11.37× and successful traverse 11.35×.

## Retained regression, protocol and publication

The official primary string-concat result increases 9.772013 → 10.416585 µs (+6.596%) despite unchanged action/runtime/harness source. A separate six-process AB/BA/AB diagnostic reports 10.172485 → 10.391398 µs (+2.152%). It does not reproduce a rise above 5%, but does not establish the full-suite cause or dismiss its observed increase. The official 10.416585 µs remains in the four indicators, README and archive; diagnostic values never replace primary samples.

Three independent processes per build in AB/BA/AB order for each suite; three global and three local warm-ups, batches calibrated to at least 10 ms and the minimum of ten batches per process. Each official time is the median of three process minima. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All other compilation, tests, profiling and computational reviews stop before the replacement series and its diagnostics. The host is not isolated. C and JS are not remeasured. All 306 replacement outputs and source/binary guards pass; the six extra diagnostic outputs and 306 excluded-series outputs are separate. Independent recomputation reproduces all four indicators from the retained official samples.

The archive includes the full dirty-baseline provenance, scoped transfers, compiler/test changes, probes with their limits, bootstrap logs, cache/source fixtures, actual-action and alias proofs, allocations, package-movement review, the excluded first series and the complete replacement, plus the separate concat diagnostic. Earlier Tree/Sum/array executions are not recounted as current validation. Only existing Go Release README cells are updated, preserving tables, references and status labels. The removed ratio catalogue remains absent. See [the JSON archive](2026-11-27-go-traverse-records.json) and var/benchmark/go-traverse-records-20261127.
