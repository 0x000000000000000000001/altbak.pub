# General JSON diagnostics — baseline and profiling notes

Preparation for the decode prototypes announced in
`docs/benchmark-results/2026-09-23-native-c-references.md`: the current Go
pipeline was re-measured and profiled per phase. No measured code path was
changed by this work.

## Baseline

Workspace `var/benchmark/json-dec-baseline-20260923`, results
`var/benchmark/json-dec-baseline-results-20260923`, Go runtime only, unchanged
protocol (three processes, two warm-ups, five samples per phase, median of
process minima, `GOMAXPROCS=1`, `GOGC=100`, PGO off).

| Runtime | parse | decode | combined |
|---|---:|---:|---:|
| Go, re-measured (this workspace) | 2,366.2 µs | 11,232.6 µs | 14,294.3 µs |
| Go, published `json-dec-c-results-20260923` | 2,376.3 µs | 11,100.2 µs | 14,666.7 µs |

Per-corpus allocations are unchanged to the byte: **4,612,800 B parse**,
**16,556,992 B decode**, **21,169,776 B combined**. The re-measurement is
consistent with the published Go cells within run-to-run variation, so the
driver changes below did not alter the measured path.

## Method and platform caveat

- The driver (`src/Test/JsonDecoding.go`) grew an opt-in profile mode:
  `DIAG_PROFILE_PHASE`, `DIAG_PROFILE_PASSES`, `DIAG_CPU_PROFILE`,
  `DIAG_MEM_PROFILE`. Defaults preserve the campaign behaviour (verified by the
  baseline above).
- On this machine (Apple M4 Pro, macOS, Go 1.27.0 darwin/arm64), `runtime/pprof`
  CPU profiles are **not usable for this binary**: most samples land on
  `runtime.kevent` (65%) and `runtime.madvise` (15–20%) while `time` reports
  ~0.01 s of system time for the whole run. `GODEBUG=asyncpreemptoff=1` does
  not fix the attribution; a trivial spin benchmark profiles correctly, so the
  artefact is tied to this workload's runtime threads, not to the toolchain
  alone. CPU hotspots were read from `sample(1)` leaf counts instead; the
  allocation profiles (heap sampling) are reliable and were used as the
  quantitative basis.
- `sample` samples every thread, including blocked ones; idle leaves
  (`__psynch_cvwait`, `__semwait_signal`) were excluded when reading user-code
  distributions.

## Findings

1. **GC ≈ 14% of the decode phase** at `GOGC=100` (`GODEBUG=gctrace=1`, mean
   over 100 cycles), consistent across the run.
2. **Decode allocations (16.6 MB per pass) are a long tail.** Flat
   `alloc_space` leaders: `Right`/`Just` constructors ≈ 19.5%; record-decoder
   machinery ≈ 15% (plan rebuild, field steps, record dictionaries);
   foreign-object/lookup adapters ≈ 9%; `decodeForeignObject` construction
   ≈ 6%; `decodeArray` ≈ 3% plus setup; JSON accessors
   (`caseJsonNumber`/`caseJsonString`) ≈ 6.5%; `Box`/`Any` ≈ 4.4%.
3. **Measured local costs** (instrumented copy, one validating pass plus 40
   profiled passes):
   - `RecordConsImpl` (record-plan rebuild): 91,368 calls; ≈ 645 µs per pass,
     i.e. **≈ 6.5% of the decode phase**. All dictionary and closure arguments
     are freshly allocated on every call (91,368 calls, 91,368 distinct
     `reflectName`/`step` closures), so no runtime cache keyed by argument
     identity is possible. The whole dictionary chain is rebuilt once per
     decoded record — a codegen artefact, not a decoder-level cost.
   - `decodeArray` setup (composition of the specialised traverse): 30,456
     calls, ≈ 55 µs per pass (**≈ 0.6%**). The element loop itself is already
     generated Go.
   - `decodeForeignObject` setup: 104,040 calls, ≈ 161 µs per pass (**≈ 1.6%**).
4. **Reading.** The remaining decode gap is dominated by (a) per-instantiation
   dictionary construction inside lambdas, (b) boxed constructor traffic
   inherent to the current representation, and (c) allocation/GC on the
   non-generational collector. The largest single local target (the record-plan
   rebuild) cannot be fixed at runtime; the actionable library-level items
   (lookup adapters, foreign-object and array setup) are worth a few percent
   each.

## GC tuning lever (measured 24 September)

The decode loop is allocation-heavy with a small live heap, so the collection
frequency dominates. Same binary (`json-dec-cache3`), interleaved runs,
decode phase only unless stated:

| Configuration | cycles (100 passes) | GC CPU | decode | combined |
|---|---:|---:|---:|---:|
| `GOGC=100`, `GOMAXPROCS=1` | 196 | 14% | 10,832.7 µs | 14,405.9 µs |
| `GOGC=300`, `GOMAXPROCS=1` | 42 | 3% | 7,915.9 µs (**−26.9%**) | 11,044.2 µs (**−23.3%**) |
| `GOGC=600`, `GOMAXPROCS=1` | — | — | 7,502.1 µs (**−28%**) | 10,104.5 µs |
| `GOGC=100`, `GOMAXPROCS=4` | 305 | 3% | 8,189.9 µs | 10,826.1 µs |
| `GOGC=300`, `GOMAXPROCS=4` | 92 | 1% | 7,704.4 µs (**−5.9%**) | 9,784.5 µs (**−9.6%**) |

Reading: the `gctrace` CPU share (14%) understates the cost of frequent
collections — mutator assists, scan and sweep credit are not counted there,
while relaxing the target removes them. The lever is largest on a single P
(the campaign protocol) and still measurable with four Ps (an application
shape). The campaign protocol stays at `GOGC=100` for comparability with the
JavaScript and C references; a default for generated applications is a
separate memory-policy decision (heap growth versus RSS) and has not been
applied.



## Candidate prototypes, in order

The dictionary hoisting sketched here is now implemented and measured (see
`2026-09-23-closed-dictionary-caching.md`); the remaining candidates are:

1. **Codegen fusion of immediately destructured `Maybe`/`Either` results**
   (case-of-known-constructor). Covers the largest remaining allocation group —
   `Just`/`Right` constructors ≈ 20.8% of decode allocations — plus the CPS
   accessor wrappers (`caseJsonNumber`/`caseJsonString`, ≈ 7.3%) that rebuild a
   constructor per call.
2. **Value-level `Foreign.Object.lookup` / field-step path** (≈ 10% of decode
   allocations). Local to the argonaut-codecs FFI, verifiable against the frozen
   error oracle; expected time gain is modest because the `Maybe` construction
   itself remains.
3. **`GOGC` policy for generated applications** (see the lever above). A
   memory-policy decision, not a code change in the measured path.
4. **Native `decodeForeignObject`/`decodeArray`** (setup ≈ 2% combined); the
   bulk of their allocations is the decoded maps and arrays themselves.

Tooling kept for the next experiments: the profile-capable driver, the
`sample`-based workflow, and the instrumented copy in
`var/benchmark/json-dec-count-20260923` (hacked counters, not a build source).
