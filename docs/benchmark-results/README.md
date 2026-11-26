# Benchmark reports

These are **dated technical reports for compiler and runtime contributors**.
They preserve measurements, experiments, rejected candidates and validation
proofs. A result describes its campaign; it may have been superseded.
For current figures, start with the [main tables](../../README.md).

## Recent Go results

| Report | What to read it for |
| --- | --- |
| [Native effect callbacks — 26 November](2026-11-26-go-native-callbacks.md) | Typed native entries for cached integer effects, preserved generic callbacks, all four Go/C indicators and three-path opaque-input proofs. |
| [FFI argument conversions — 26 November](2026-11-26-go-ffi-calls.md) | Cached immutable argument conversions, allocation proofs, all four Go/C indicators and 37-case regression checks. |
| [State affine loop — 25 November](2026-11-25-go-state.md) | Paired Go optimization preserving strict global reads, all four Go/C indicators, opaque-input proofs and regression diagnostics. |
| [Church native integer products — 25 November](2026-11-25-go-church.md) | Paired Go optimization, bounded native multiplication preserving JS rounding, all four Go/C indicators and regression checks. |
| [Lazy — 25 November](2026-11-25-go-lazy.md) | Paired Go optimization, all four Go/C indicators against unchanged imperative C, exact Int32 arithmetic and opaque-input proofs. |
| [Church and State — 25 November](2026-11-25-go-church-state.md) | Paired Go optimizations, all four Go/C indicators against unchanged imperative C, opaque-input proofs and 37-case regression checks. |
| [Polymorphism and RowToList — 25 November](2026-11-25-go-small-actions.md) | Paired Go optimization, allocations, all four Go/C indicators against unchanged C references, and 37-case regression checks. |
| [Go / functional C / official JS — 25 November](2026-11-25-go-c-js-functional.md) | Fresh 14-case comparison with a structural C reference, both descending rankings, complete raw evidence and allocation limits. |
| [JSON Decoding integration — 24 November](2026-11-24-go-json-decoding.md) | Before/after results, allocations and remaining C gap. |
| [C reference precision — 24 November](2026-11-24-c-reference.md) | The 14 measured references and their comparison limits. |
| [AST, Fibonacci and records — 24 November](2026-11-24-go-c-priorities.md) | Three compiler optimizations, measured against the C reference then in use. |
| [Go/JS ratios — 24 November](2026-11-24-go-js-ratios.md) | Validation of the 37-case suite against JavaScript. |
| [Prime Sieve — 24 November](2026-11-24-go-prime-sieve-progression-fusion.md) | The progression-fusion optimization and its validation. |

## Earlier investigations

- **Runtime and libraries:** [Go opportunities audit](2026-11-21-go-opportunities.md),
  [integrated Go gains](2026-11-21-go-significant-gains.md),
  [library measurements](2026-11-14-library-js-go.md).
- **JSON:** [compact DOM](2026-09-24-compact-json-dom.md),
  [schema specialization](2026-09-24-schema-decoder-specialization.md),
  [owned C++ references](2026-09-25-json-owned-references.md),
  [Rust decoding](2026-09-30-rust-json-packed.md).
- **Compiler performance:** [Purust native optimization](2026-10-03-purust-native-optimization.md),
  [Rust-hosted gopurs allocation work](2026-10-03-gopurs-rust-allocation.md),
  [compilation-table refresh](2026-10-25-gopurs-compilation-refresh.md).

Long reports keep their implementation details and reproduction evidence in
expandable **Technical reference** sections. Their numerical tables and raw
archives are retained; the files in this folder retain the intervening history.
