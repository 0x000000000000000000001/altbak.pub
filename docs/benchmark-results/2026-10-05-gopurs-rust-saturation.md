# Rust-hosted gopurs saturation campaign — 4–5 October 2026

All three gopurs hosts generate **Go**. The final production table covers
**50 libraries plus b8x**, with newly built JavaScript, Go and Rust compilers.

## Qualified results

| Backend-only median | JavaScript host | Go host | Rust host |
| --- | ---: | ---: | ---: |
| b8x → Go | 55.78 s | 33.04 s | 22.40 s |
| 50 libraries, sum of medians | 285.27 s | 64.55 s | 51.79 s |
| Complete table, sum of medians | 341.04 s | 97.58 s | 74.19 s |

The final Rust host is **24.0% below the Go host**
on the complete table. The prior published Rust total was
**92.22 s**. A separate contemporary paired
campaign measures the frozen published Rust binary at **91.46 s**
and the new production binary at **77.15 s**:
**15.6% lower**. Its paired
bootstrap ratio interval is **[0.8413, 0.8500]**.
The new Rust median is lower on **50/51 projects** in that paired campaign.

### Recorded backend phases

These values sum each project's phase median across the complete three-host
table, in seconds. Preparation includes monomorphization. PBO, generation/writes
and drain overlap; the rows must not be added to reconstruct the backend total.

| Phase, sum of project medians | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
| load TAST + sort | 23.88 | 3.28 | 2.88 |
| prepare + monomorphize | 26.47 | 26.07 | 23.36 |
| PBO producer | 230.73 | 60.79 | 39.97 |
| generation + writes (cumulative batches) | 287.66 | 56.03 | 38.73 |
| emitter drain | 59.64 | 7.39 | 7.18 |

## Selection and stopping decision

- **Direct syntax folds** — Eliminate temporary FreeMonoidTree construction; Arrays screening 3149 → 2068 ms and b8x 29798 → 28244 ms, with exact fold order and outputs.
- **Bulk directive removal** — Persistent-map intersection/difference preserves rank visibility and shadowed defaults; b8x screening 28117 → 26668 ms.
- **Dynamic preparation scheduling** — Small atomic work claims preserve ordered round merges; b8x screening 26211 → 23895 ms, with deferred/replayable bounded scheduling contracts.
- **Lexical closed-scope proof** — Avoid constructing free-variable sets for every subtree and reject cheaply before traversal; composition-v1 confirms Arrays 1900 → 1512 ms.
- **Import accumulator** — One deduplicated file-level accumulator; full-table confirmation gives Arrays 1511 → 1442 ms with a lower overall total.
- **Observed specialization misses, stable runtime layout** — Invalidate only on relevant new membership, with a sequenced per-transform lookup journal and unchanged common runtime. Final composition confirms 76138 → 75168 ms overall and Arrays 1447 → 1325 ms.

Other selection decisions:
- **Module-range filtering** — Removed after 51-case ablation: its 0.83% overall and 2.42% b8x improvements remain below the agreed practical thresholds.
- **Annotation rewrite memo** — All five screening cases regress by 1–2% despite passing contracts.
- **Lookup-first alone** — Neutral alone; retained only as the lookup interface needed by the observed-miss mechanism, with no independent gain claimed.
- **Precise cache with extra record shapes** — Complete-backend regressions despite faster preparation. The retained tuple variant keeps the shared generated runtime byte-identical; this does not prove a sole cause for the earlier regression.
- **Adaptive directive prefixes / shared range trees** — No gain at the practical thresholds in their paired screenings.

Screenings compare one mechanism with a contemporary control; their percentages
are not added together. The final composition receives its own confirmation and
fresh production PGO training. The priority emission/allocation and preparation/monomorphization mechanisms have been investigated. After the last newly winning mechanism, three successive distinct mechanism experiments failed the practical thresholds; subsequent composition confirmation is not a new mechanism. Final profiles show the already-tested directive filtering/removal and type-name/eligibility costs on b8x, and distributed syntax traversal, analysis, scope and code generation on Arrays/Aff. No new isolated mechanism with a convincing threshold-sized benefit emerged from this review. This is practical saturation for this campaign and corpus, not a proof that further optimization is impossible.

The last non-retained mechanisms were:
- **Reuse closed-type eligibility** (`cached-eligibility`): b8x +1.01%, Spec +2.08%; five-case sum 33387 → 33786 ms.
- **Pruned contribution-suffix fold** (`suffix-contributions`): b8x −0.46%, paired interval crosses equality; five-case sum 33404 → 33291 ms, below practical thresholds.
- **Chunked type-name construction** (`mangle-chunks`): b8x +15.38%, Arrays +5.52%; five-case sum 33484 → 38253 ms.

Every proposal, including failed construction/test attempts, remains in the
machine-readable report and local archive with its source, generated-code and
binary fingerprints. Initial fixed-profile screening and final self-trained
PGO are recorded separately.

## Method

- Apple M4 Pro, 10 performance and 4 efficiency cores; no explicit affinity.
- Frozen TAST/FFI inputs from the qualified compilation refresh. Every generation
  is checked byte-for-byte against its canonical Go oracle; input manifests and
  sibling FFI inputs are rechecked between runs.
- One warmup and five measured fresh processes per project and host. Projects,
  builds and measurement campaigns are serialized. The three-host table cycles
  through all six host permutations; the old/new Rust comparison alternates.
- Workers **8/8/8/8**, pipeline enabled. Go alone uses
  **GOGC=off / GOMEMLIMIT=10GiB**. The public launcher still selects Go by default,
  JavaScript with `GOPURS_JS=1` and Rust with `GOPURS_RUST=1`.
- The clock is the internal **backend total**: loading/sorting, preparation,
  monomorphization, PBO, generation/writes and drain. Frontend work, compiler and
  application builds, application execution and process startup/exit are outside
  timing. Producer, emission and drain phases overlap and are not additive.
- Each total is a sum of per-project medians. Bootstrap resampling preserves
  paired rounds within each project. Practical selection thresholds are 1% of
  the full-table total, or 3% on an important case with no significant overall
  regression, above observed variability.
- The final three-host campaign was interrupted after 48 complete projects. The remaining 3 projects were restarted with a fresh warmup and five whole rounds, using identical compiler artifacts. The interrupted records and 5 superseded checked generations remain archived; they are excluded from the published medians and final coverage.
- Diagnostic sampling covers all threads, excluding known blocking leaves.
  Inclusive categories overlap and are not CPU percentages. Instrumented runs
  are never used as selection timings.
- Before the campaign, **28,837,597,998 logical bytes**
  of regenerable Purust intermediates were reclaimed after inventory. Preserved
  sources, executables and historical evidence were rehashed.

## Qualification

- **1530 final checked generations**, including
  **1275 measured runs** and
  **474,690 byte-exact generated files**,
  across the production table and separate old/new paired campaign.
- Production Rust: O3, ThinLTO, threaded Arc, mimalloc and fresh compiler-self
  PGO, with three checked training passes excluding Test.Main. Bootstrap outputs
  are compared with the unprofiled compiler; independent JavaScript/native Purust
  generation checks establish Rust source identity.
- Affected fold, directive visibility, monomorphization, scheduling, scope and
  native FFI contracts are retained with their exact compiler builds.
- Fresh-project, launcher, parser/race and three-host native regression fixtures
  pass. Aff is checked on all three hosts. Representative measured applications
  (Aff, Arrays, Spec, Yoga JSON, Enums, Promise, Prelude and Strings) build and run;
  all generated b8x packages and entry points build outside timing.
- The first application qualification detected Yoga JSON writing a `.spec-results` report in the frozen-input working directory. All application commands succeeded and no pre-existing input changed. The report and failed qualification remain archived; the corrected harness reruns all applications in private working copies, checks Yoga JSON's recorded assertions and revalidates every frozen corpus. Both timing campaigns are preserved unchanged.

## Complete table (milliseconds)

| Project | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
| b8x | 55775 | 33037 | 22404 |
| gopurs-aff | 6716 | 1581 | 1253 |
| gopurs-argonaut-core | 7096 | 1375 | 1113 |
| gopurs-arrays | 5533 | 1624 | 1281 |
| gopurs-assert | 537 | 150 | 106 |
| gopurs-avar | 6552 | 1382 | 1124 |
| gopurs-catenable-lists | 4105 | 1025 | 791 |
| gopurs-console | 4345 | 707 | 554 |
| gopurs-datetime | 6113 | 1449 | 1127 |
| gopurs-effect | 1880 | 158 | 113 |
| gopurs-enums | 4800 | 922 | 732 |
| gopurs-exceptions | 2032 | 241 | 173 |
| gopurs-foldable-traversable | 4683 | 872 | 694 |
| gopurs-foreign | 4814 | 1159 | 949 |
| gopurs-foreign-object | 6914 | 1488 | 1227 |
| gopurs-free | 4695 | 1035 | 830 |
| gopurs-functions | 1895 | 183 | 124 |
| gopurs-integers | 4451 | 823 | 642 |
| gopurs-js-bigints | 2047 | 195 | 144 |
| gopurs-js-date | 7051 | 1450 | 1166 |
| gopurs-js-promise | 3431 | 542 | 427 |
| gopurs-js-promise-aff | 7188 | 1498 | 1201 |
| gopurs-js-uri | 4368 | 710 | 557 |
| gopurs-lazy | 3031 | 523 | 413 |
| gopurs-node-buffer | 6490 | 1239 | 1011 |
| gopurs-node-event-emitter | 9012 | 2589 | 2204 |
| gopurs-node-fs | 8371 | 1731 | 1406 |
| gopurs-node-http | 9304 | 1768 | 1460 |
| gopurs-node-net | 8542 | 1722 | 1396 |
| gopurs-node-path | 1881 | 161 | 113 |
| gopurs-node-process | 7713 | 1649 | 1330 |
| gopurs-node-streams | 9772 | 2784 | 2401 |
| gopurs-now | 7796 | 1652 | 1340 |
| gopurs-nullable | 6102 | 1155 | 951 |
| gopurs-numbers | 4457 | 837 | 656 |
| gopurs-ordered-collections | 7054 | 1720 | 1346 |
| gopurs-partial | 4320 | 706 | 552 |
| gopurs-prelude | 4742 | 860 | 675 |
| gopurs-random | 4388 | 710 | 556 |
| gopurs-record | 4525 | 732 | 573 |
| gopurs-refs | 4360 | 710 | 558 |
| gopurs-run | 7198 | 1621 | 1300 |
| gopurs-spec | 12510 | 3826 | 3269 |
| gopurs-st | 4352 | 717 | 564 |
| gopurs-strings | 5555 | 1090 | 924 |
| gopurs-strings-extra | 5879 | 2598 | 1794 |
| gopurs-unfoldable | 4372 | 725 | 569 |
| gopurs-unsafe-coerce | 4340 | 710 | 553 |
| gopurs-uuid | 8612 | 2591 | 2174 |
| gopurs-variant | 4541 | 1101 | 811 |
| gopurs-yoga-json | 14801 | 5751 | 4558 |

## Evidence

- [Machine-readable report](2026-10-05-gopurs-rust-saturation.json)
- Local archive: `var/benchmark/gopurs-rust-saturation-20261004/`
- Reproduction tools: `bin/benchmark/gopurs-rust-saturation/`
- [Previous qualified compilation table](2026-10-04-compilation-refresh.md)
- [Earlier gopurs optimization campaign](2026-10-04-gopurs-rust-night.md)
