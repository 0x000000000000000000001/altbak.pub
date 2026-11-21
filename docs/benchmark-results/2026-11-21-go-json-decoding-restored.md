# Go JSON decoding restored — November 21, 2026

The native compiler again generates the direct text-to-schema decoder.
The matched before/after comparison improves **9.526 → 2.172 ms (4.38×)**,
winning all six process pairs. This restores the historical **2.241417 ms**
level. The latest README refresh had reported **11.338583 ms**; that measurement
used a different session and phase schedule, so the speedup above uses the
remeasured before binary from the same paired campaign.

## Cause

The published historical result came from a preserved optimized executable.
The [September 30 archive](2026-09-30-rust-json-native.json) already distinguishes
that executable from a reconstructed generic Go build around 11.8 ms.
The November rebuild exposed this existing loss of specialization. The corpus
and oracle hashes still match the historical workload.

Two compiler recognition problems prevented the optimized path:

1. Monomorphization cloned an Argonaut field-reader wrapper and embedded its
   decoder dictionary in an imported function. In the actual application IR,
   the `items` read became a two-argument `Combinators.getField__…` call.
   `BorrowedObjects` correctly rejected this opaque use of the object: its proof
   requires a known saturated reader with the decoder operand visible.
2. After exposing that reader, `DecoderSchemas.Programs` still recognized only
   `decodeJson(dictionary)`. PBO had produced the equivalent
   `dictionary.decodeJson` accessor, so derivation stopped again.

The event decoder therefore remained an opaque custom callback. That made the
complete schema ineligible for direct text decoding, and the public
`decodeJsonStringWith` composition fell back to parsing a JSON DOM and decoding
it. The CPU profile records this fallback branch; allocation profiling places
most allocation below the custom event decoder and its dictionary/record-plan
construction. CPU attribution on this host was dominated by `kevent`/`madvise`,
so no precise CPU-time partition is claimed from that profile.

## Correction

`Gopurs.Monomorphization` retains the three canonical Argonaut field-reader
wrappers for ordinary PBO inlining. This exclusion is applied after transitive
collection: their bodies remain available, and other eligible functions still
specialize. It is separate from foreign forwarding and does not infer a decoder
from an instantiated type.

`DecoderSchemas.Programs` accepts the actual `decodeJson` property operand and
subjects its dictionary to the existing schema proof. Borrowing, callback,
ownership and error-forwarding rules remain unchanged.

The final native-generated application contains **five text-schema workers**
and **zero opaque custom schema callbacks**, versus no text-schema workers
before. The benchmark source and fixtures are unchanged. No memoized benchmark
result, reused parser or cross-iteration decoded value was introduced.

## Matched results

Each cell is the median of six independent process minima. Before/after process
order alternates, and all six parse/decode/combined phase orders are exercised.
The pinned binaries run sequentially after all builds finish, with
`GOMAXPROCS=1`, `GOGC=100` and PGO off. Two warmups precede five retained,
fingerprint-checked samples per phase. Phase minima must not be added together.

| Phase | Before, ms | After, ms |
| --- | ---: | ---: |
| Parse control | 1.876 | 1.882 |
| Decode already-parsed JSON | 7.184 | 0.669 |
| Complete text-to-value | **9.526** | **2.172** |

| Minimum allocated bytes per corpus | Before | After |
| --- | ---: | ---: |
| Parse | 2,435,696 | 2,435,696 |
| Decode | 16,230,288 | 1,954,560 |
| Complete text-to-value | **18,666,304** | **3,561,448** |

The parse control remains stable; combined time falls **77.2%** and combined
allocations fall **80.9%**. The historical 2.241417 ms result is a useful reference,
but its old executable is no longer available on this machine. The 3.1%
difference from that historical number is not a separately established gain.

## Validation and evidence

- The new method-accessor test fails before the fix and passes afterwards.
  Opaque/local dictionaries, other properties and extra applications remain
  rejected.
- The reader-wrapper regression verifies all three exclusions, preserved
  bodies, and continued specialization of unrelated functions and same-named
  functions in other modules.
- **66 compiler tests pass**, including generated DOM/text workers, exact
  errors, owned results and race checks. The compiler builds without warnings
  or errors, and its native executable was rebuilt and installed.
- All twelve measured processes match the existing 17-case oracle. The five
  timed fixtures produce **900 retained sample outputs** checked after timing;
  error cases are validated outside timing as before.

The [JSON archive](2026-11-21-go-json-decoding-restored.json) records all process
results, source changes, build and executable hashes, corpus/oracle hashes,
diagnostic probes and profile fingerprints. Full artifacts are under
`var/benchmark/go-json-regression-20261121/`. Host: Apple M4 Pro, macOS 26.1,
Go 1.27.0. The README JSON Decoding cell now uses this six-process result;
other cells and the core total retain their preceding measurements.

```sh
GOPURS_GO_PACKAGES=0 GOPURS_GO_PACKAGE_OVERLAY=0 \
  python3 bin/benchmark/json-diagnostic.py --runtime go --suite JsonDecoding \
  --workspace path/to/new-build build

PYTHONPATH="$PWD/bin/benchmark" \
  python3 bin/benchmark/json-diagnostic/paired.py --suite JsonDecoding \
  --baseline path/to/preserved-before --current path/to/new-build \
  --output path/to/new-results --processes 6
```
