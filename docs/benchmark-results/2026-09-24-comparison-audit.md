# Comparison audit: is the C++ reference a fair yardstick?

Three questions drove this audit:

1. do the two sides do the same work and pay the same lifetime costs,
2. what does a **hand-written Go decoder** for the same schema cost (the point
   between "gopurs generated" and "specialized C++"), and
3. what do Go parsers cost **down to usable values**, including simdjson behind
   cgo.

All numbers were measured on the frozen corpora with the diagnostic protocol
(two warm-up passes, five sampled passes, minimum per process; medians across
processes where noted), `GOMAXPROCS=1`, `GOGC=100` unless stated. The machine
carried background load; paired in-run deltas and allocation totals are the
stable parts.

## 1. Obligations and lifetimes

### 1.1 The C++ reference frees outside the timed interval — and it is negligible

The C++ drivers build one arena per pass and release it after the timed work.
A copy of each driver was extended with an audit hook
(`bin/benchmark/json-diagnostic/audit/*-fullcycle.cc`,
`DIAG_C_FULL_CYCLE=1`) that releases the arena and the payload vector **inside**
the pass and reports the release time; the stock sources and binaries are
unchanged, and the audit variant without the variable reproduces the stock
numbers exactly.

| Suite | configuration | parse | decode | combined | release |
|---|---|---:|---:|---:|---:|
| JSON Decoding | stock | 381 µs | 267 µs | 652 µs | — |
| JSON Decoding | freeing inside | 388 µs | 264 µs | 679 µs | **~1 µs** |
| JSON to Typed AST | stock | 4,461 µs | 4,754 µs | 9,148 µs | — |
| JSON to Typed AST | freeing inside | 4,457 µs | 4,709 µs | 9,056 µs | **~1 µs** |

Freeing the decoded structures costs about **1 µs per pass** in both suites
(the allocator reuses the pages immediately). The asymmetry exists but cannot
explain the gap.

### 1.2 The Go side pays collector work inside the interval — 27.7% on JSON decoding, ~2% on TAST

`GOGC=off` against `GOGC=100`, same binary, interleaved:

| Path | phase | `GOGC=100` | `GOGC=off` | delta |
|---|---|---:|---:|---:|
| JSON Decoding (gopurs Go) | decode | 10,688 µs | 7,733 µs | **−27.7%** |
| JSON Decoding (gopurs Go) | combined | 14,260 µs | 10,555 µs | −26.0% |
| JSON to Typed AST (gopurs Go) | decode | 16,006 µs | 15,653 µs | −2.2% |

The general JSON decode loop allocates heavily into a small live set, so the
collector dominates it; the TAST decoder builds long-lived native structures and
barely notices the setting.

### 1.3 The Go TAST decoder does two things the C++ reference does not

Measured with a purpose-built program in the generated module (see
`audit/tast-validation-share/`), against the frozen fingerprints
(`oracle_matches: true` in every run), medians over eight runs:

| Layer | time (12 modules) |
|---|---:|
| `DecodeModuleImpl` + real usage validation | 15,338 µs |
| same decoder, no-op validate | 13,200 µs |
| native worker, no record materialisation (`Call_Test_JsonTypedAst_decode`) | 14,085 µs |
| driver entry point (`Apply(decode, json)`) | 14,320 µs |

So the pure `validateSourceUsageModule` pass costs **≈ 2.1 ms (10–15% of the Go
decode)** — larger than the 872.8 µs (6.6%) recorded in the
`2026-09-23-native-c-references.md` report, which was measured on an older
build with a different method. The record materialisation at the PS boundary is
~0–2% at most.

### 1.4 Identical type-table work

The C++ reference resolves types lazily and the Go decoder resolves every
entry; the audit counter in the TAST driver shows the lazy path still builds
**21,574 of 21,574 entries** on this corpus, so there is no scope advantage
there.

### 1.5 Verdict for section 1

The comparison is an **implementation comparison**, not a language comparison,
and the measurable asymmetries are small relative to the gap: freeing ~1 µs,
GC 27.7% of the JSON decode (a real but partial factor), TAST validation
~10–15%, materialisation ~0–2%. Errors-message contracts differ but sit outside
the timed work.

## 2. A hand-written Go decoder for the same schema

`bin/benchmark/json-diagnostic/native-go/` implements the schema twice: a
two-phase `dom` mode (`encoding/json` into `any`, then a hand-written typed
walk, C++-style) and a one-step `std` mode (`encoding/json` straight into
structs, what an application writes), plus a `goccy` variant. All successful
fingerprints and parse fingerprints match the frozen oracle; the ten malformed
cases are rejected.

JSON Decoding, 636 KB, 5 timed cases (medians of three process minima):

| Path | parse | decode | combined | allocated/pass |
|---|---:|---:|---:|---:|
| gopurs generated (current) | 2,373 µs | 9,134 µs | 12,738 µs | 4.61 / 13.49 MB |
| gopurs generated, `GOGC=off` | — | 7,733 µs | 10,555 µs | — |
| C++ reference | 381 µs | 267 µs | 652 µs | arena |
| hand-written Go, DOM → typed walk | 7,434 µs | **303 µs** | 9,492 µs | 5.99 / 0.67 MB |
| hand-written Go, DOM → typed walk + string copy | 7,256 µs | 415 µs | 8,355 µs | 5.99 / 0.90 MB |
| `encoding/json` → structs | — | — | **4,187 µs** | 2.76 MB |
| `goccy/go-json` → structs | — | — | **1,511 µs** | 2.49 MB |
| `goccy/go-json` DOM | 3,136 µs | — | — | 5.96 MB |

Reading:

- The **typed decode work itself is cheap in Go**: 303 µs hand-written against
  9,134 µs through the generic Argonaut path — a factor of ~30 — and within 14%
  of the C++ decode (267 µs). Whatever the gap is, it is not the language.
- The **generic path allocates 20× more** than the typed walk (13.49 MB against
  0.67 MB per corpus) and pays it back through the collector.
- A complete hand-written Go application path costs **1.5 ms (goccy) to 4.2 ms
  (stdlib)** against 0.65 ms for the arena C++ reference — a factor of 2.3 to
  6.4, on a workload where gopurs currently needs 12.7 ms.
- The gopurs **parser is competitive** (2.37 ms) with the Go alternatives at
  building a DOM (3.1 ms goccy, 7.4 ms stdlib); the reverse ordering of the
  combined numbers shows where the time actually goes.

## 3. Parsers down to usable values

| Parser | parse (DOM) | direct to structs | notes |
|---|---:|---:|---|
| simdjson (C++, in process) | 381 µs | — | the reference bound |
| simdjson behind cgo | **481–493 µs** | — | one cgo call per document, node count computed in C |
| gopurs native parser | 2,373 µs | — | best pure-Go DOM result measured |
| `goccy/go-json` | 3,136 µs | 1,511 µs (combined) | |
| `encoding/json` | 7,434 µs | 4,187 µs (combined) | |

The cgo crossing itself was measured at **16.2–17.0 ns per call**; the corpus
holds **84,130 values**, so a naive per-value conversion through cgo would cost
≈ **1.4 ms** — as much as goccy's entire combined time. Any cgo integration
would therefore have to move data in bulk (a compact C-side serialisation or a
shared buffer), and would land in roughly the same range as the fast pure-Go
decoder, at the price of a C/C++ toolchain in the distribution.

## 4. What this means for the comparison

- The current published cells put Go at about **×20** (JSON decoding combined,
  12.74 ms against 0.645 ms) and **×5.8** (TAST combined, 50.51 ms against
  8.759 ms) from the C references. Those ratios mix three different things:
  **implementation strategy**
  (generic dictionary-driven decoding against a schema-specialised decoder),
  **representation** (boxed values, records as maps, collector against arena
  structs), and **obligations** (usage validation, record materialisation).
- The measured asymmetries explain only a fraction of it. The hand-written Go
  points show the decode *work* can be within ~10% of the C++ reference, and a
  full hand-written Go path within 2.3×, with today's gopurs at 12.7 ms.
- The parser is the one place where pure Go is **4–8× behind** the SIMD
  reference; cgo closes that to ~1.3× but cannot convert per value.
- Conclusion for planning: "Go cannot reach C" is **not demonstrated** by these
  numbers; what is demonstrated is that the remaining gap lives in codegen
  specialisation (removing generic dispatch and boxing from the decoded path)
  and, to a smaller degree, in parser technology. The [specialised-decoder
  audit](2026-09-24-specialized-decoder-audit.md) measured that claim directly:
  producing the same final PureScript values through a specialised path costs
  0.56 ms against 8.18 ms for the generated decoder (14.7×), which puts the
  complete path at 3.4 ms with today's parser. Reaching roughly 1.5–2.5 ms
  would additionally require a faster parser; specialising the decoder alone
  does not get there.

## Provenance

- Audit drivers and runners: `bin/benchmark/json-diagnostic/audit/`
  (`decoding-fullcycle.cc`, `typed-ast-fullcycle.cc`, `run.py`),
  `bin/benchmark/json-diagnostic/native-go/` (`main.go`, `run.py`),
  `bin/benchmark/json-diagnostic/simdjson-cgo/`,
  `bin/benchmark/json-diagnostic/audit/tast-validation-share/`.
  The stock `decoding.cc` and `typed-ast.cc` are unmodified.
- Results: `var/benchmark/json-audit-20260924/` (C++ stock/audit/freeing,
  ~1 µs release), `var/benchmark/json-dec-native-go-20260924/` (hand-written Go
  modes), `var/benchmark/json-dec-native-go-smoke-20260924/` (single-process
  smoke), cgo and TAST-layer numbers printed by the audit programs above.
- Frozen oracle: `test/fixtures/json-decoding/expected.json`,
  `test/fixtures/json-typed-ast/expected.json`; validated in every measured
  process of sections 2 and 3 and in the TAST layer program.
- Campaign cells quoted for gopurs: `var/benchmark/json-dec-cache7-results-20260923`
  and `var/benchmark/json-tast-cache7-results-20260923`; the frozen C++ cells
  are from `var/benchmark/json-dec-c-results-20260923` and
  `var/benchmark/json-tast-c-results-20260923`, reproduced within noise by the
  stock binaries in this audit.
