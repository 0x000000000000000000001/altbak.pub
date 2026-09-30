# Dense JSON results and allocation-guided decoding

This follows the [generated-decoder/native-record milestone](2026-09-30-rust-json-native.md)
for **JsonDecoding**. JsonTypedAst remains deferred.

## Published result

The Rust cell is now **1 130.69 µs**, previously 1 372.79 µs. The final paired
campaign reruns the preserved previous Rust build and the published optimized
Go/C executables under the same protocol:

| Phase | Previous Rust, rerun | Rust owned-v11 | Published Go, rerun | C, rerun |
|---|---:|---:|---:|---:|
| parse | 1 722.000 µs | **1 659.229 µs** | 1 905.709 µs | 373.438 µs |
| decode | 688.042 µs | **661.375 µs** | 636.938 µs | 274.458 µs |
| combined | 1 342.167 µs | **1 130.688 µs** | 2 141.146 µs | 659.646 µs |

Combined takes **15.8% less time than the previous Rust build**, **47.2% less
than the published Go build**, and **1.71× the C time**, down from about 2×.
The decode-only phase is still 3.8% slower than Go in this campaign.

The independent official campaign confirms **1 122.833 µs** Rust combined;
JavaScript is 7 995.292 µs, reconstructed Go 9 980.334 µs and C 645.375 µs.
That reconstructed Go is distinct from the optimized published reference.
The published Go/C table cells keep their historical 2 241.42 / 657.69 µs.

Raw measurements, executable/input manifests, allocation profiles, stack
attribution and qualification metadata are archived in
[`2026-09-30-rust-json-packed.json`](2026-09-30-rust-json-packed.json).

## Representation

Generated workers eagerly construct `Vec`s of concrete records, primitive
values, or the existing typed ADT carriers. `Value::NativeArray` transports
them through the ordinary `UnknownType` boundary. An escaping record/ADT
element carries a shared array owner and an index. It retains the whole backing
array; it does not retain a JSON cursor or postpone decoding.

- The thin owner keeps `Value` at **24 bytes** on the measured 64-bit target.
  The initial fat-pointer/index prototype enlarged it to 32 bytes, increasing
  DOM memory and cancelling the allocation savings. That prototype is rejected.
- Scalar vectors live directly in the owner. Nonempty record/ADT vectors have
  one extra type-erasure box per array, rather than one wrapper per element.
- Indexing, `IntItems`, array iterators and folds preserve the backing storage.
  Reading record/ADT elements allocates no wrapper or intermediate vector.
  Legacy `unwrap_array` explicitly converts to boxed element views.
- Record field projection dispatches directly to the concrete array element,
  avoiding an intermediate virtual `NativeRecord` projection. Reflection,
  immutable updates and `Foreign.Object` conversion remain supported.
- The thin owner caches the array length beside its erased vector. Indexing
  returns an owning element view directly, with a bounds check and no virtual
  call just to discover whether the element is a record or an ADT.
- ADT arrays retain their existing `Rc<ADT>`/`Arc<ADT>` representation and
  identity while removing their outer `Value::Class` allocation. Parameterized
  ADTs are not yet fully specialized.
- A generated decoding call shares its payload-free `Nothing` values through
  one stack-local slot. The ordinary `Maybe` representation is retained, and
  the slot is released before the call returns. There is no global optional
  cache or CAF change; a `Weak` contract checks complete release.

`--no-json-arrays` isolates the array representation. Dictionary/method proofs,
error fallback, duplicate-key semantics and owned UTF-16 strings retain the
contracts of the preceding milestone.

## Allocation-guided follow-up

`bench/allocation-stacks.py` now captures all allocation/reallocation stacks in
one measured combined pass in a separate debug-info build. Its reentrancy guard
is entered **before** initializing the map or symbolizing a backtrace. This
resolves the earlier profiler blockage without attributing its precise cause
retroactively. The first complete capture records **43 840 requests**, matching
the preceding thin-owner stage profile. The final capture and stage counters
are rebuilt from the same canonical workspace and reconcile exactly below.

The resulting changes are general compiler/runtime paths:

- Escaped strings copy complete plain byte runs, preserving the encoded UTF-16
  units. A validated tape token supplies an upper bound on the decoded byte
  length, so result construction reserves once instead of repeatedly growing
  the string.
- Already validated, short decimal `Int` spellings use bounded integer
  arithmetic. Exponents, decimal points, overflow and other spellings retain
  IEEE-754 parsing and the original integer range check.
- A constructor's static function and its public wrapper consume their owned
  parameters once. They no longer clone a `String` at both boundaries. Tests
  check the string buffer address, the single ADT allocation, shared argument
  identity and reusable partial applications.
- Eta-expanded aliases of global functions also transfer their synthetic
  parameters, removing the full input-text copy in the `decodeText` wrapper.
- A method's discriminator stays borrowed only when the proved success program
  never includes that read in a constructed result. Returned strings remain
  owned. Escaped discriminators and strings used both in a test and a result
  are covered. This extension uses the **ABI3** codec handshake; older ports
  retain the ordinary decoder.

Isolated paired experiments: text/string/integer changes **1 354.604 →
1 175.417 µs**, constructor ownership **1 190.605 → 1 172.126 µs**, borrowed
discriminators **1 159.896 → 1 141.646 µs** combined. These experiments guide the
final qualification; they are not substituted for its independent campaign.

Before the final alias-wrapper change, the complete stack capture drops from
**43 840 to 31 553** requests. Its disjoint attribution reconciles exactly:

| Attributed request sites | Before | After | Removed |
|---|---:|---:|---:|
| Escaped-string buffers | 7 885 | 1 577 | 6 308 |
| `View` constructor, including its argument copies | 5 254 | 1 752 | 3 502 |
| Discriminator-only string reads | 2 477 | 0 | 2 477 |
| **Total removed** | | | **12 287** |

The remaining `Just` carriers account for 6 308 requests in that capture. Their
payloads and ordinary identity are preserved; fully specialized parameterized
ADTs remain a subsequent representation step.

## Rejected scanner experiment

Using `trailing_zeros` to return the first SWAR match directly was functionally
correct but slower: **1 365.604 → 1 490.708 µs** combined against the same packed
array build. It is not retained. The original bounded eight-byte scanner remains;
the additional multi-special-byte and cross-lane borrow tests are retained.

The `packed-v1`/`packed-v2` fat-owner builds and `scan-v4`/`share-v5` direct-match
scanner builds are experimental artifacts, not published performance cells.

## Measurement and validation

The final snapshot is `owned-v11`, generated from the canonical drivers.
Workspaces, executable snapshots, input manifests and raw reports are preserved
under `var/benchmark/json-packed-20260930/`.

The paired protocol remains six phase-order permutations, two warm-ups, five
samples and the median of process minima, with `GOMAXPROCS=1` and `GOGC=100`.
The comparison uses the preserved **published optimized Go and C binaries**.
The complete PureScript consumer is timed separately with `consume-json.py`.

- Codegen: **81/81**. TAST: **43/43** with the final stable compiler bundle.
- JSON schemas: **nine tests in each of five modes** (ordinary, generated boxed,
  native records, dense arrays, threaded), with identical differential outputs.
  Scanner checks pass in both native and threaded modes.
- b8x: **286/286**, build `build-Rue4Pa`. The clean compiler rebuild succeeded;
  its Linux step was interrupted when the container was replaced. Repeating
  `bin/b --runtime rust`, then `bin/t --runtime rust -c`, completed successfully.
- All timed campaigns started after builds and validation processes completed.
  Rust/Go match all 17 oracle cases; C matches successful values exactly and
  rejects malformed inputs under the established reference contract.

### Allocation requests, combined, one timed case per process

| Case | Previous native-record build | Final owned-v11 | Reduction |
|---|---:|---:|---:|
| Empty arrays | 11 | 10 | 9.1% |
| Flat record arrays | 16 014 | 9 017 | 43.7% |
| Nested records and variants | 14 013 | 11 014 | 21.4% |
| Null and missing optionals | 11 263 | 4 514 | 59.9% |
| Unicode, escaping and numbers | 14 435 | 6 997 | 51.5% |

The final whole-corpus stack capture records **31 548 requests / 3 426 649
requested bytes**: index 5 / 1 334 904, result 31 527 / 1 424 048, outside
16 / 667 697. These totals agree exactly with the disjoint stage counters.
The alias-wrapper change removes five input copies totalling 667 217 bytes
relative to the preceding stack-guided build. Per-case process totals include
their own driver buffer, hence their sum differs from the single-corpus total.

`instrument.py` now retains per-case allocation samples and manifests.
`stages-json.py` profiles disjoint index/result/outside regions in a separate
copy of the generated workspace. Allocation counts include alloc/realloc
requests and requested bytes, not live memory. Instrumented times are diagnostic
and are not substituted for benchmark cells.

`instrument.py` rebuilds TAST before generating the profiling binary and checks
the exact oracle for each individual case. A reused ABI2 profiling workspace
initially produced the ordinary decoder with the ABI3 compiler; that
`allocations-v10.json` experiment is excluded from candidate allocation claims.
The `stages-v10` and `stacks-v10` builds copied the fresh canonical workspace and
are unaffected. Final per-case counts come from the refreshed TAST build.

### Consumer-inclusive check

Both final builds time the complete ordinary PureScript encode/fingerprint
consumer, differing only in `--no-json-arrays`:

| Phase including consumer | Native records, ordinary arrays | Dense arrays |
|---|---:|---:|
| parse + encode | 7 318.417 µs | 7 393.312 µs |
| decode + fingerprint | 15 402.480 µs | 15 635.354 µs |
| combined + fingerprint | 15 867.146 µs | 15 966.688 µs |

The final isolated array comparison retains a **0.63% combined / 1.51%
decode-only cost**. No standalone array-consumer speedup is claimed. The earlier
`packed-v6` comparison showed a 3.4% combined regression; direct indexing reduced
that cost substantially. The arrays are retained for their allocation reduction
and eager ABI-compatible storage.

A separate paired check against the preserved earlier `consume6-records`
control measures **16 211.979 → 16 064.917 µs** combined and **15 713.833 →
15 518.959 µs** decode, including consumers. That control already includes
decode-local absences and is not the original published native-record binary.

## Reproduction

From `purust/purust`, with a built TAST compiler:

```sh
npm run build
python3 bench/build-json.py --workspace /path/to/work --snapshot /path/to/rust
python3 bench/paired.py --workspace /path/to/work \
  --rust /path/to/rust --baseline /path/to/previous-rust \
  --corpus ../../altbak.pub/test/fixtures/json-decoding/corpus.json \
  --go /path/to/published/benchmark-go --c /path/to/published/benchmark-c \
  --output /path/to/paired.json
python3 bench/consume-json.py --workspace /path/to/work --output /path/to/packed-consumer
python3 bench/consume-json.py --workspace /path/to/work --output /path/to/record-consumer --no-arrays
python3 bench/stages-json.py --workspace /path/to/work --output /path/to/stages
python3 bench/allocation-stacks.py --workspace /path/to/work --output /path/to/stacks
```

Run the consumer pair with `paired.py` too, targeting their
`rust-project/target/release/purust_output` executables. Time uninstrumented
binaries after builds and validations have completed.

## Implementation commits

- `purust` (`edge`): `3e86109` — direct native-array indexing, borrowed
  discriminators, owned constructor/function-alias arguments, contracts and
  allocation-stack diagnostics. Dense-array integration was committed earlier
  as `eeb229c` and `74e267a`.
- `purust-argonaut-core`: `6295fc0` — escaped-string buffers, bounded integer
  conversion and borrowed text comparison.
- `purust-argonaut-codecs`: `2737472` — ABI3 DOM/text discriminator predicates.
- The array consumer ports retain the earlier integration commits:
  arrays `22f2d55`, prelude `0fb4742`, foldable-traversable `e626300`,
  foreign `ee31447`, yoga-json `c9e097b` and spec `187fad4`.
