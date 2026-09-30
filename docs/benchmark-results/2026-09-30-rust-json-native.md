# Generated JSON decoders and native records in purust

This cycle specializes **JsonDecoding**, including the actual custom `Event`
decoder. It builds an eager, owned result across the existing polymorphic
`drive` boundary. JsonTypedAst remains deferred.

## Published result

The Rust JSON Decoding cell is now **1 372.79 µs**, previously 6 817.98 µs.
The final paired campaign compares fresh baseline/candidate Rust builds with
the actual preserved, published Go/C executables:

| Phase | Rust baseline | Rust generated/native | Go published, rerun | C reference, rerun |
|---|---:|---:|---:|---:|
| parse | 1 895.417 µs | **1 843.125 µs** | 2 219.375 µs | 391.667 µs |
| decode | 3 970.083 µs | **779.271 µs** | 794.396 µs | 286.730 µs |
| combined | 6 290.250 µs | **1 372.791 µs** | 2 518.500 µs | 686.625 µs |

Combined is **4.58× faster than the rebuilt Rust baseline** and takes **45.5%
less time than the published Go binary** in the same campaign. It remains
**2.00× the C reference**. The Go and C published table cells retain their
original campaign values, 2 241.42 and 657.69 µs respectively.

The independent official harness campaign succeeds too: Rust **1 400.459 µs**
combined, JavaScript 9 104.375 µs, current reconstructed Go 11 811.041 µs and C
672.375 µs. The reconstructed Go artifact is distinct from the optimized
published executable used for the comparison above.

Raw paired reports, build fingerprints, consumer measurements and the official
campaign are preserved in
[`2026-09-30-rust-json-native.json`](2026-09-30-rust-json-native.json).

## Implementation

- **Dictionary and method recognition.** `Purust.DecoderSchemas` checks the
  actual `DecodeJson`, `DecodeJsonField` and `NativeField` composition. A bounded
  symbolic interpreter extracts field reads, branches and constructor calls
  from custom methods, including the closures left by `bind` and `apply`.
  A result type alone never supplies a decoding rule.
- **Generated DOM workers.** Child results use `Option<Value>` by value; only
  the public root constructs `Either`. On failure, the original composition
  owns the error and any recovery. Opaque callbacks are not run speculatively.
  Selected DOM fields are read under one lock, released before recursion.
- **Eager native record carriers.** Generated structs store concrete primitive
  fields and owned strings. `Value::NativeRecord` transports them through
  `UnknownType`. Getters, borrowed scalar projections, immutable updates,
  reflection and `Foreign.Object` interoperation support that representation.
  Arrays, optionals and custom ADTs retain their ordinary, eagerly constructed
  carriers. No deferred whole-result materialization is used.
- **Text-to-result workers.** A call-local offset tape validates the entire
  input, including ignored fields, then feeds the same generated workers.
  Normalized duplicate keys are last-wins; field order is arbitrary. Strings
  retain internal UTF-16 code units, including isolated surrogates, and own
  their result storage. Deep or unsupported inputs use the ordinary path.
- **Scanner and ownership.** Plain strings scan eight byte lanes per machine
  word with safe, bounded loads. Number lexing is separate from conversion,
  so selected numbers are converted once. Last-use field cursors move into
  their decoder, and primitive constructor arguments use a typed by-value ABI.

## Protocol

The canonical PureScript and Rust drivers are copied by the official harness,
then built through `spago build`, purust and release Cargo (`opt-level=3`, thin
LTO). All commands have strict exit checking; executables are preserved before
the next experiment.

The paired campaign runs six phase-order permutations, two warm-ups and five
samples per phase per process. Cells are medians of the six process minima.
`GOMAXPROCS=1`, `GOGC=100`; the Go and C binaries are the preserved published
reference executables, not the current Go compiler's reconstructed benchmark.
Hashes cover executables, corpus and oracle before and after measurement.

Rust and Go reproduce all 17 decoding and JSON fingerprints exactly. C checks
the same successful results and rejects error cases, following its existing
reference contract; it does not reproduce Argonaut error messages.

The initial DOM/text/layout experiments inherited `GOMAXPROCS=14`. They are
exploratory measurements, not the publication protocol. `bench/paired.py` now
sets and checks the published environment and validates sample minima.

## Consumer-inclusive representation check

`bench/consume-json.py` places the complete ordinary PureScript
encode/fingerprint consumer inside the timed interval. This is a separate
diagnostic, not the published decoding cell. Both builds use the same scanner
and worker compiler, differing in `--no-json-layouts`.

| Phase including consumer | Ordinary records | Native records |
|---|---:|---:|
| parse + encode | 8 084.104 µs | 7 980.896 µs |
| decode + fingerprint | 17 433.042 µs | 16 720.063 µs |
| combined + fingerprint | 18 019.833 µs | **17 195.167 µs** |

The native representation remains faster with its consumer included. The
complete result is constructed before the published decoding timer stops;
compound fields are not reconstructed by later projections.
This isolated comparison predates the final last-use cursor/constructor move
optimization; it compares the representation change with identical worker code.

## Validation and reproduction

- Codegen: 81/81. TAST: 43/43; the prior `js-bigints` dependency blocker is
  resolved. Existing exact allocation and pointer-identity contracts remain.
- After the final cursor/constructor move change, all codegen tests and the
  four-mode schema fixture were rerun successfully.
- The schema fixture compares ordinary, generated-boxed, native-record and
  threaded modes, including 512 differential documents and ordinary PS
  consumers through an opaque polymorphic FFI boundary.
- Cases cover error recovery, repeated reads, immutable updates and aliases,
  missing/null fields, escaped duplicate keys, ignored malformed values,
  every UTF-16 code unit, truncated documents, deep nesting and long numbers.
- The scanner checks all byte values at each lane/tail position in both Rc
  and Arc builds.
- b8x: `bin/b --runtime rust -c && bin/t --runtime rust -c`, 286/286. This
  qualifies the native-record runtime; the subsequent changes are local to
  the JSON scanner and generated schema workers, covered by their tests.

From `purust/purust`:

```sh
npm run build
python3 bench/build-json.py --workspace /path/to/work --snapshot /path/to/preserved-rust
python3 bench/paired.py --workspace /path/to/work \
  --rust /path/to/preserved-rust --baseline /path/to/preserved-baseline \
  --corpus ../../altbak.pub/test/fixtures/json-decoding/corpus.json \
  --go /path/to/published/benchmark-go --c /path/to/published/benchmark-c \
  --output /path/to/paired.json
```

Workspaces are under `var/benchmark/json-native-20260930` and
`var/benchmark/json-decoding-native-final2-20260930`.

## Remaining work toward C

The next representation work concerns concrete array elements and parameterized
ADTs, integrated with ordinary consumers and their conversions. Further scanner
vectorization should follow measurement. There is no claimed allocation-stack
attribution or assumed speedup for these future steps.

## Commits

- `purust` (`edge`): `f203297` — compiler pass, native-record runtime, typed
  constructor arguments, tests and reproducible benchmark tooling.
- `purust-argonaut-core`: `dce8d39` — validated offset tape, scalar conversion,
  bounded word-at-a-time scanning and scanner tests.
- `purust-argonaut-codecs`: `faeba20` — DOM/text cursor ABI and by-value workers.
