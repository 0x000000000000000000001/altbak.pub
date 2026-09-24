# Comparison audit: is the C++ reference a fair yardstick?

**Interpretation update:** these are exploratory implementation references.
The native-Go timed paths retain a sink and validate setup outputs; the
original specialised runner had the same limitation. The [record-plan
follow-up](2026-09-24-record-plan-follow-up.md) aligns retention and validation
for generated versus specialised gopurs outputs. C++ lifecycle alignment,
native-Go timed-output validation, and owned simdjson-to-Go value transfer
remain separate work.

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

### 1.1 Separately measured C++ arena release

The C++ drivers build one arena per pass and release it after the timed work.
A copy of each driver was extended with an audit hook
(`bin/benchmark/json-diagnostic/audit/*-fullcycle.cc`,
`DIAG_C_FULL_CYCLE=1`) reports release time separately. Inspection shows the
main elapsed interval stops before fingerprinting and release; this does not
measure a complete retained-output lifecycle. The stock sources were unchanged.

| Suite | configuration | parse | decode | combined | release |
|---|---|---:|---:|---:|---:|
| JSON Decoding | stock | 381 µs | 267 µs | 652 µs | — |
| JSON Decoding | release audit | 388 µs | 264 µs | 679 µs | **~1 µs** |
| JSON to Typed AST | stock | 4,461 µs | 4,754 µs | 9,148 µs | — |
| JSON to Typed AST | release audit | 4,457 µs | 4,709 µs | 9,056 µs | **~1 µs** |

The measured release operation costs about **1 µs per pass** in these runs.
This does not by itself quantify lifetime asymmetry with the retained Go results.

### 1.2 Sensitivity to collector settings

`GOGC=off` against `GOGC=100`, same binary, interleaved:

| Path | phase | `GOGC=100` | `GOGC=off` | delta |
|---|---|---:|---:|---:|
| JSON Decoding (gopurs Go) | decode | 10,688 µs | 7,733 µs | **−27.7%** |
| JSON Decoding (gopurs Go) | combined | 14,260 µs | 10,555 µs | −26.0% |
| JSON to Typed AST (gopurs Go) | decode | 16,006 µs | 15,653 µs | −2.2% |

These observations show allocation/collector sensitivity under this schedule;
they are not isolated measurements of GC's time share. Default `GOGC` remains 100.

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

The observed validation delta was **≈ 2.1 ms**; the older report recorded
872.8 µs using another build/method. Fixed ordering and unequal repetitions
limit attribution of validation and boundary-materialisation costs; the
measurements do not establish a 0–2% upper bound for the boundary.

### 1.4 Identical type-table work

The C++ reference resolves types lazily and the Go decoder resolves every
entry; the audit counter in the TAST driver shows the lazy path still builds
**21,574 of 21,574 entries** on this corpus, so there is no scope advantage
there.

### 1.5 Verdict for section 1

These implementations differ in representation, lifetimes and validation
obligations. The probes identify costs worth investigating but do not fully
align those obligations. Error-message contracts differ outside timed work.

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
| gopurs generated (historical cache7) | 2,373 µs | 9,134 µs | 12,738 µs | 4.61 / 13.49 MB |
| gopurs generated, `GOGC=off` | — | 7,733 µs | 10,555 µs | — |
| C++ reference | 381 µs | 267 µs | 652 µs | arena |
| hand-written Go, DOM → typed walk | 7,434 µs | **303 µs** | 9,492 µs | 5.99 / 0.67 MB |
| hand-written Go, DOM → typed walk + string copy | 7,256 µs | 415 µs | 8,355 µs | 5.99 / 0.90 MB |
| `encoding/json` → structs | — | — | **4,187 µs** | 2.76 MB |
| `goccy/go-json` → structs | — | — | **1,511 µs** | 2.49 MB |
| `goccy/go-json` DOM | 3,136 µs | — | — | 5.96 MB |

Reading:

- The hand-written Go path supplies a useful fast reference (303 µs), but
  produces different representations and implements a narrower error contract.
- The **generic path allocates 20× more** than the typed walk (13.49 MB against
  0.67 MB per corpus) and pays it back through the collector.
- The hand-written Go benchmark reports **1.5 ms (goccy) to 4.2 ms (stdlib)**
  for combined work. This is not a full Argonaut-equivalent application result.
- The gopurs **parser is competitive** (2.37 ms) with the Go alternatives at
  building a DOM (3.1 ms goccy, 7.4 ms stdlib); the reverse ordering of the
  combined numbers shows where the time actually goes.

## 3. Parser references and the remaining usable-value transfer experiment

| Parser | parse (DOM) | direct to structs | notes |
|---|---:|---:|---|
| simdjson (C++, in process) | 381 µs | — | implementation reference |
| simdjson behind cgo | **481–493 µs** | — | one cgo call per document, node count computed in C |
| gopurs native parser | 2,373 µs | — | best pure-Go DOM result measured |
| `goccy/go-json` | 3,136 µs | 1,511 µs (combined) | |
| `encoding/json` | 7,434 µs | 4,187 µs (combined) | |

The cgo microbenchmark copies input, parses, counts **84,130 nodes including
object keys**, releases the C-side state and returns a count. Its 481–493 µs
does not include conversion into usable Go values. The **16.2–17.0 ns** crossing
measurement times a different minimal operation; multiplying it by node count
to obtain ~1.4 ms is not a measured transfer cost. Bulk transfer, copying,
ownership, lifetime and the C/C++ distribution cost still need an integrated
experiment before predicting application performance.

## 4. What this means for the comparison

- The current published cells put Go at about **×20** (JSON decoding combined,
  12.74 ms against 0.645 ms) and **×5.8** (TAST combined, 50.51 ms against
  8.759 ms) from the C references. Those ratios mix three different things:
  **implementation strategy**
  (generic dictionary-driven decoding against a schema-specialised decoder),
  **representation** (boxed values, records as maps, collector against arena
  structs), and **obligations** (usage validation, record materialisation).
- The hand-written Go and C++ paths motivate implementation experiments, with
  their representation and semantic differences made explicit.
- The cgo count-only result does not close the parser-to-usable-value gap.
- The older specialised-decoder figures do not isolate codegen overhead or
  establish a combined-time floor. Library-side tag propagation and record/
  array allocation remain measurable levers. General compiler specialization
  and parser/fusion experiments require their own retained-output validation.

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
  `test/fixtures/json-typed-ast/expected.json`; setup validation in the decoder
  audits. The simdjson-cgo count-only probe does not validate decoded values.
- Campaign cells quoted for gopurs: `var/benchmark/json-dec-cache7-results-20260923`
  and `var/benchmark/json-tast-cache7-results-20260923`; the frozen C++ cells
  are from `var/benchmark/json-dec-c-results-20260923` and
  `var/benchmark/json-tast-c-results-20260923`, reproduced within noise by the
  stock binaries in this audit.
