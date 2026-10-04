# Complete compilation-table refresh — 4 October 2026

All **153 gopurs project cells** (51 projects × three hosts), both three-column
totals, and both **purust-aff** cells are refreshed from this campaign.

## Results

| Backend-only median | JavaScript host | Go host | Rust host |
| --- | ---: | ---: | ---: |
| b8x → Go | 62.48 s | 39.57 s | 31.02 s |
| 50 gopurs libraries → Go (sum of medians) | 304.78 s | 74.02 s | 61.20 s |
| Complete gopurs table (sum of medians) | 367.27 s | 113.58 s | 92.22 s |
| purust-aff → Rust | 6316 ms | — | 2173 ms |

## Protocol and provenance

- Apple M4 Pro, 14 logical CPUs, 48 GiB memory; Node v24.8.0; no explicit CPU affinity.
- One warmup and five measured fresh processes per host; serialized cases and
  six host permutations over the six rounds. The JavaScript/Rust pair alternates
  for Purust. Each cell is the median of its five successful measurements.
- The timer is the compiler's internal **backend total**: loading and sorting,
  preparation, PBO, generation/emission and final drain. Frontend compilation,
  compiler bootstrap, application builds/execution and process startup/exit are
  outside this timer. Phase samples and medians are retained in the JSON report;
  overlapping producer, cumulative emission and drain clocks are never added.
- gopurs uses workers **8/8/8/8**, pipeline enabled. The public launcher applies
  Go's **GOGC=off / GOMEMLIMIT=10GiB** policy only to the Go host. Purust uses
  its public defaults (JS 1/1; native PBO 8, generation 4). OS caches stay warm;
  generated output and PBO caches are reset between runs.
- All five compiler artifacts match the most recent qualified rebuild, including
  the division, deterministic module ordering and IEEE negative-zero fixes.
  Rust-hosted gopurs uses O3, ThinLTO, threaded Arc, mimalloc and self-trained PGO.
  Binary hashes, frozen compiler sources and earlier bootstrap/test evidence are
  linked in the JSON report and retained in the archive.
- The 50 library rows use their retained reference TAST/FFI corpora. Every host
  must match the previously qualified Go source oracle, including the explicitly
  corrected Prelude signed-zero oracle. Historical failed evidence is retained.
- b8x uses the full local Go-target application corresponding to the public
  b8x project. Its frontend is regenerated from frozen current sources because
  20 cached source inputs were stale; its automatic entry-point selection is
  retained. Purust Aff's cached frontend was verified against all 296 source
  hashes before freezing its TAST and Rust FFI.
- Every generation is checked byte-for-byte. A manifest is kept for every run;
  identical source bytes are archived once per case, with every run referring
  to that canonical output. Failed runs, if any, retain their own diagnostics
  and generated files. All frozen inputs and compilers are rechecked afterward.
- The first b8x series was interrupted by a machine/server restart with 16
  complete, output-identical generations retained. Its raw state and logs are
  archived separately. The published b8x medians use a complete new warmup and
  five-run series on the same frozen inputs; the 900 completed library runs
  retain their original, uninterrupted campaign.

## Validation

- **930 checked generations**, including **775 measured runs**;
  **290,766 exact generated files** checked.
- Generated Go applications for Aff, Arrays, Enums, Promise, Prelude and Strings
  build and execute successfully outside timing. Aff prints all 45 checks.
- All generated b8x Go packages and entry points build outside timing.
- The Purust Aff application builds offline and executes all 47 checks, with
  exact expected output. JavaScript and Rust hosts generate identical Rust/Cargo files.
- The README totals are **sums of per-project medians**, not timings of one
  multi-project invocation.

## All gopurs medians (milliseconds)

| Project | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
| b8x | 62481 | 39565 | 31018 |
| gopurs-aff | 7520 | 1940 | 1528 |
| gopurs-argonaut-core | 7702 | 1518 | 1264 |
| gopurs-arrays | 6732 | 3393 | 3373 |
| gopurs-assert | 570 | 166 | 115 |
| gopurs-avar | 7171 | 1628 | 1291 |
| gopurs-catenable-lists | 4474 | 1182 | 898 |
| gopurs-console | 4648 | 788 | 600 |
| gopurs-datetime | 6678 | 1788 | 1469 |
| gopurs-effect | 2012 | 171 | 120 |
| gopurs-enums | 5193 | 1052 | 823 |
| gopurs-exceptions | 2167 | 257 | 184 |
| gopurs-foldable-traversable | 5114 | 1042 | 862 |
| gopurs-foreign | 5265 | 1330 | 1068 |
| gopurs-foreign-object | 7502 | 1785 | 1538 |
| gopurs-free | 4887 | 1106 | 910 |
| gopurs-functions | 1944 | 187 | 127 |
| gopurs-integers | 4774 | 947 | 787 |
| gopurs-js-bigints | 2113 | 200 | 150 |
| gopurs-js-date | 7410 | 1581 | 1285 |
| gopurs-js-promise | 3517 | 573 | 445 |
| gopurs-js-promise-aff | 7645 | 1630 | 1344 |
| gopurs-js-uri | 4573 | 752 | 593 |
| gopurs-lazy | 3163 | 551 | 430 |
| gopurs-node-buffer | 6858 | 1351 | 1120 |
| gopurs-node-event-emitter | 9520 | 2849 | 2468 |
| gopurs-node-fs | 8837 | 1915 | 1585 |
| gopurs-node-http | 9943 | 1939 | 1626 |
| gopurs-node-net | 9112 | 1899 | 1576 |
| gopurs-node-path | 1984 | 169 | 120 |
| gopurs-node-process | 8308 | 1817 | 1499 |
| gopurs-node-streams | 10387 | 3167 | 2781 |
| gopurs-now | 8384 | 1822 | 1515 |
| gopurs-nullable | 6458 | 1264 | 1040 |
| gopurs-numbers | 4745 | 923 | 731 |
| gopurs-ordered-collections | 7625 | 2173 | 1937 |
| gopurs-partial | 4520 | 750 | 589 |
| gopurs-prelude | 5021 | 974 | 750 |
| gopurs-random | 4636 | 760 | 590 |
| gopurs-record | 4700 | 785 | 615 |
| gopurs-refs | 4554 | 759 | 597 |
| gopurs-run | 7720 | 1776 | 1440 |
| gopurs-spec | 13361 | 4434 | 3837 |
| gopurs-st | 4574 | 772 | 607 |
| gopurs-strings | 6078 | 1447 | 1274 |
| gopurs-strings-extra | 6042 | 2756 | 1943 |
| gopurs-unfoldable | 4561 | 782 | 610 |
| gopurs-unsafe-coerce | 4539 | 757 | 589 |
| gopurs-uuid | 9044 | 2785 | 2422 |
| gopurs-variant | 4937 | 1221 | 902 |
| gopurs-yoga-json | 15562 | 6403 | 5233 |

## Evidence

- [Machine-readable results](2026-10-04-compilation-refresh.json)
- Local archive: `var/benchmark/compilation-refresh-20261004/`
- Reproduction tools: `bin/benchmark/compilation-refresh/`
- Source qualification: [corrected package campaign](2026-10-04-gopurs-packages-fixes.md)
