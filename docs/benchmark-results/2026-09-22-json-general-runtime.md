# General JSON decoding: compact records and Value FFI

Three general changes reduce the ordinary Argonaut diagnostic's Go parse + decode
time from **61.042 to 45.816 ms (−24.94%)** in a paired campaign. Decoding alone
falls **27.14%**, with **26.68% fewer allocated bytes**. The remaining Go/JS gap
is **5.30×**. This is a useful improvement, not closure of that gap.

The TAST diagnostic benefits much less: **582.022 → 567.609 ms (−2.48%)** and
**2.95% fewer allocated bytes**. Its timing ranges overlap, so this campaign
does not establish a strong TAST throughput improvement.

## Changes

- `gopurs/runtime`: converting empty or single-field records to native class
  dictionaries avoids a temporary map and sorting keys. Existing constructor
  pointers and the general sorted conversion remain unchanged.
- `gopurs-prelude`: `Record.Unsafe.unsafeSet` uses a new immutable `RecordSet`
  runtime operation. Records with up to five fields retain their compact
  representation when extended; larger records copy their field storage.
  The source record and aliases remain unchanged. Map inputs keep a copying fallback.
- `gopurs-foldable-traversable` and `Gopurs.FfiBridge`: array traversal passes
  runtime `Value` objects and arrays directly across the FFI boundary. Callbacks
  already accepting/returning `Value` avoid intermediate boxing and forwarding
  closures. The traversal algorithm and evaluation order are preserved.

These apply to ordinary PureScript programs. No JSON-specific decoder, PBO
change, parser replacement, cache, or increase in worker count was introduced.

## Paired results

| Diagnostic / phase | Go before (ms) | Go after (ms) | JS before (ms) | JS after (ms) |
|---|---:|---:|---:|---:|
| JSON Decoding: parse | 6.886 | 7.286 | 1.616 | 1.618 |
| JSON Decoding: decode | 50.175 | **36.556** | 7.282 | 7.086 |
| JSON Decoding: combined | 61.042 | **45.816** | 8.793 | **8.647** |
| JSON → Typed AST: parse | 75.342 | 76.282 | 20.675 | 20.035 |
| JSON → Typed AST: decode | 466.219 | 457.856 | 59.305 | 57.565 |
| JSON → Typed AST: combined | 582.022 | **567.609** | 75.764 | **79.059** |

| Go allocated bytes per corpus, MiB | Before | After | Change |
|---|---:|---:|---:|
| JSON Decoding: decode | 77.368 | 56.729 | −26.68% |
| JSON Decoding: combined | 83.099 | 62.460 | −24.84% |
| JSON → Typed AST: decode | 521.440 | 504.106 | −3.32% |
| JSON → Typed AST: combined | 586.711 | 569.378 | −2.95% |

The README's previous official baselines were **57.43 ms Go / 8.26 ms JS** for
JSON Decoding and **573.63 / 75.90 ms** for JSON → Typed AST. Relative to those
historical Go cells, the new observations are −20.22% and −1.05%, respectively.
The paired results above are the primary attribution evidence; run-to-run
variation is visible in both runtimes. The parser was not changed.

## Protocol and validation

Apple M4 Pro, Go/Node versions and all source/binary hashes in the
[raw archive](2026-09-22-json-general-runtime.json). Three independent processes
per version and runtime; two warm-ups and five samples per phase in each process.
Cells are medians of process minima. GOMAXPROCS=1, GOGC=100, PGO disabled.
Order per round: before/Go, after/JS, after/Go, before/JS; reversed in round two.
No builds, tests, or profilers ran concurrently with these timings. File I/O and
fingerprinting are outside the timed interval; GC is included as in the existing driver.

All 24 processes matched the fixed JSON and decoded-value oracles: 17 general
JSON cases (five timed, ten errors, two optional-field cases), and 12 real TAST
modules. The general diagnostic's JS bundle is byte-identical before/after.
The TAST fingerprint helper had been merged into its main test module; this
rebuild repaired missing imports and a stale source-list entry. Its complete
fingerprint and timed parse/decode/main definitions were verified textually
unchanged against the frozen previous build. Its JS bundle consequently differs.

The JS and native compilers were rebuilt, including the installed native binary.
The full runtime suite passes with `-race`; 15 targeted FFI/record tests pass,
including deferred callbacks, Either error priority, ordered Cartesian products,
and State effects. Thirteen runner checks also pass. No new full b8x measurement
was performed; these results do not quantify whole-compiler acceleration.

## What the profile says to do next

The initial decode profile attributed about 23% of allocated bytes directly to
`RecordToMap`. Isolated prototypes separated dictionary conversion, record
insertion, and traversal FFI before combining them. The final decode allocation
profile puts `RecordToMap` at about 1.4%; allocations now remain spread across
generated callbacks, generic applications, `Either`/`Maybe` constructors, JSON
primitive adapters, array concatenation, and foreign-object copies. Cumulative
profile percentages overlap and must not be added or treated as promised gains.

The next substantial compiler experiment is **eliminating immediately applied
closures across case branches**. The generated `gDecodeJsonCons` still builds
a function in each `Either` branch, then immediately applies it to another
function that inserts a field and constructs `Right`. Distributing the application
into the branches and beta-reducing can remove those intermediates without
special-casing JSON. Measure this first; preserve single evaluation, error
priority, effects, escaping functions, and partial application.

Next candidates are producer/consumer fusion of short-lived `Either`/`Maybe`
values, and specialization of known array traversals to avoid intermediate
arrays and concatenations. These need distinct semantic checks and measurements.
There is no measured ×2 prediction. The TAST path needs its own renewed profile
before attributing its remaining time to any of these candidates.

Local reproducibility artifacts (frozen builds, comparison script, profiles,
prototype runs and build/test logs):
`/Users/0x1/Documents/htdocs/scratch/json-general-20260922`.
