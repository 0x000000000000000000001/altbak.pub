# Native C/C++ references for the excluded diagnostics

The three `excluded WIP` rows in the README tables (Array Indexing, JSON
Decoding, JSON to Typed AST) had no C cell, so there was no measured bound for
"how close is the generated code to a hand-written native implementation". This
report adds those references, with the same frozen corpora, the same timing
protocol and the same in-process oracle checks as the existing Go and
JavaScript diagnostics.

## What the references are

| Row | Source | Build | Scope |
|---|---|---|---|
| Array Indexing | `bin/benchmark/array-indexing/driver.c` | `clang -O3` | the exact diagnostic kernel (cyclic indexed reads), in a `native` case (`int64_t[]`) and a `boxed` case (24-byte tagged `Value`, the same layout as `gopurs_runtime.Value`); C has one runtime representation, so the two cases differ only by the array element type |
| JSON Decoding | `bin/benchmark/json-diagnostic/decoding.cc` | `clang++ -O3 -std=c++17 -lsimdjson` | simdjson DOM parsing plus hand-written arena decoding (one monotonic arena per pass, reset instead of freed); fingerprints are recomputed canonically in-process and compared with the frozen oracle |
| JSON to Typed AST | `bin/benchmark/json-diagnostic/typed-ast.cc` | `clang++ -O3 -std=c++17 -lsimdjson` | simdjson parsing plus a hand-written decoder for the full typed module (type table, annotations, usage, declarations, fingerprint); the pure usage-validation pass is omitted and its Go share is reported below |

Both drivers use the protocol of the existing diagnostics: three independent
processes per runtime, two warm-up passes and five sampled passes per phase, the
minimum per process and the median across processes; `fingerprints` are computed
outside every timed interval, and each process panics on an oracle mismatch
before reporting. Newlines and indentation of the reporting code are shared with
the Go/JS drivers (`bin/benchmark/json-diagnostic.py` grew a `--runtime c` path;
`bin/benchmark/array-indexing.py` grew the same for its suite).

The C reference deliberately does **not** reproduce Argonaut's error-message
formatting for malformed inputs: the fixture excludes error cases from timing,
so the exact `AtKey`/`AtIndex`/`TypeMismatch` strings have no bearing on the
measured work. Malformed inputs still have to *fail*, which the driver records
and the runner checks.

## Array Indexing

8,388,608 indexed reads per batch, ten batches, median of three process minima
(`var/benchmark/array-indexing/campaign-c3-20260923`), cell reported for the
16,384-element case:

| Runtime | native | boxed | boxed/native |
|---|---:|---:|---:|
| Go (generated) | 4.2021 ms | 4.2079 ms | ×1.001 |
| JavaScript | 25.5312 ms | 25.3520 ms | ×0.993 |
| C (`clang -O3`) | **6.3420 ms** | **6.5890 ms** | ×1.039 |

Two earlier campaigns (`campaign-c-20260923`, `campaign-refs-20260923`) gave the
same ordering with up to 10% higher absolute times while an external build was
loading the machine.

Reading:

- On the **same serial kernel**, the Go-generated loop is about **1.6× faster
  than `clang -O3`** (0.50 against 0.79 ns per access). Go lowers the index
  wrap-around to a predicted branch; LLVM emits a conditional select on the
  flag chain, which lengthens the loop-carried dependency. This is a compiler
  artifact, not a runtime-representation difference.
- The **boxed representation costs Go 0.1%** over the native one (≤2% across
  runs). The representation the diagnostic was built to measure is therefore
  not a bottleneck in the Go backend.
- The kernel itself is serial by construction (each address depends on the
  previous index computation). An **exploratory reformulation** that keeps the
  same access order but makes the addressing independent — sum whole blocks,
  then the tail — lets `clang -O3` auto-vectorize and reaches **0.06 ns per
  access (≈0.51 ms)**; the checksums were verified identical to the cyclic
  kernel over sampled sizes and counts. That figure is **not** the row's cell
  (it is a different loop formulation, and a Go implementation could use it
  too), but it locates the remaining headroom: kernel shape, not runtime
  representation.

## JSON Decoding

The five timed cases of `test/fixtures/json-decoding`, 636 KB of JSON, five
sampled passes, median of three process minima
(`var/benchmark/json-dec-c-results-20260923`):

| Runtime | parse | decode | combined |
|---|---:|---:|---:|
| Go (generated) | 2,376.29 µs | 11,100.17 µs | 14,666.71 µs |
| JavaScript | 1,655.50 µs | 7,691.83 µs | 8,942.25 µs |
| C/C++ (`clang++ -O3`, simdjson, arena) | **377.38 µs** | **264.79 µs** | **644.96 µs** |

Ratios of C to Go: **×6.3 on parse**, **×41.9 on decode**, **×22.7 on
combined**; to JavaScript: ×4.4, ×29.1, ×13.9.

Reading:

- This is the row where the Go backend is *behind* JavaScript and where the
  general JSON suite had not yet been attacked natively. The C reference puts
  the native bound at **one twenty-second of the Go time** and one fourteenth
  of the JavaScript time.
- The decode gap is much larger than the parse gap: the Go path decodes through
  the generic Argonaut machinery (type class dispatch, boxed `Maybe`/`Either`,
  per-value records and arrays, a collector that must trace all of it), while
  the C reference builds concrete structures in one arena reset per pass. The
  diagnostic's own boxed/native Array Indexing row shows that *reading* a
  boxed representation is nearly free in Go; the cost here is building it.
- Parse alone is 6.3× slower in Go than simdjson. Some of that is parser
  technology (SIMD scanning), but the 41.9× decode gap is not explained by the
  parser: a native decoder of this corpus is dominated by allocation.

## JSON to Typed AST

The twelve corpus modules (5.55 MB of module JSON, 21,574 type-table entries),
five sampled passes, median of three process minima
(`var/benchmark/json-tast-c-results-20260923`):

| Runtime | parse | decode | combined |
|---|---:|---:|---:|
| Go (generated) | 27,461.79 µs | 15,377.21 µs | 56,990.38 µs |
| JavaScript | 20,356.46 µs | 61,988.25 µs | 80,351.96 µs |
| C/C++ (`clang++ -O3`, simdjson, arena) | **4,259.25 µs** | **4,518.54 µs** | **8,758.63 µs** |

Ratios of C to Go: **×6.4 on parse**, **×3.4 on decode**, **×6.5 on
combined**; to JavaScript: ×4.8, ×13.7, ×9.2.

Scope: the reference decodes the complete typed module — type-table resolution
with references and unresolved fallbacks, annotations (with the deliberately
empty expression spans), usage facts, data and class declarations, the foreign
map and comments — and validates all twelve fingerprints against the frozen
oracle. The only part it does not reproduce is the *pure* usage-validation
pass (`validateSourceUsageModule`, which returns `Unit` and cannot change the
fingerprint): measured on the same Go build, that pass costs **872.8 µs of the
13,206.6 µs Go decode (6.6%)**, so an equivalent C pass would land near
**4.8 ms**, still about 3.2× faster than the Go decoder.

Reading:

- This is the row where the Go backend had already been moved to a hand-written
  native decoder, and the remaining decode headroom is correspondingly smaller
  (×3.4, against ×41.9 for general JSON decoding). The largest single gap left
  is **parsing**: the Go JSON parser plus the Argonaut representation costs
  27.5 ms for 5.55 MB where a SIMD parser needs 4.3 ms (×6.4).
- Building the full typed structure (21,574 resolved type-table entries plus
  every expression, binder, literal, annotation, declaration and comment) takes
  4.5 ms in C++ with one arena; the Go decoder's 15.4 ms are the boxed
  representation, generic record/array construction and the collector work that
  follows it — not the type-table algorithm, which is already native on both
  sides.
- The JavaScript backend is ×13.7 from the native bound on decode, consistent
  with the general JSON row: the difference is representation and allocation,
  not the algorithms.

## Provenance

- Array Indexing: published cell from
  `var/benchmark/array-indexing/campaign-c3-20260923` (build
  `var/benchmark/array-indexing/build-refs-20260923`; C driver
  `sha256:e547df0356fe59e4…`, C binary `sha256:c4d5315f0a2e02d1…` per the build
  manifest, clang 17.0.0, Go 1.27.0, Node v24.8.0). The exploratory block
  formulation was measured with a standalone equivalent-kernel program validated
  against the cyclic kernel.
- JSON Decoding: `var/benchmark/json-dec-c-results-20260923` (workspace
  `var/benchmark/json-dec-c-20260923`, C binary
  `sha256:e6c99be2af83453c7deaa945d4843df299d3b52cd4ece3d7b64325bffdc5ce22`,
  source `decoding.cc` `sha256:f87161a0168eea…`, clang++ 17.0.0, simdjson at
  `/opt/homebrew/opt/simdjson`, Go 1.27.0, purs from the local TAST fork).
- Machine state: absolute cells are session-sensitive. The first array-indexing
  campaigns ran under an external build load (load average up to ~26); the
  published campaign and the JSON decoding campaign ran with an otherwise idle
  machine. The paired in-campaign ratios are the stable part, as in the other
  reports.
- JSON to Typed AST: `var/benchmark/json-tast-c-results-20260923` (workspace
  `var/benchmark/json-tast-c-20260923`, C binary
  `sha256:57c0f644aff01b57…`, source `typed-ast.cc` `sha256:23fa680767a31d3c…`,
  clang++ 17.0.0, simdjson at `/opt/homebrew/opt/simdjson`, Go 1.27.0). The
  twelve fingerprints were first compared textually, module by module, with the
  PureScript fingerprint produced by the same Go build
  (`scratch/tast-revolution-20260923/dec5`), then validated through the frozen
  oracle in every measured process. The usage-validation share (6.6%) was
  measured with `TestValidationShare` on the same build.
