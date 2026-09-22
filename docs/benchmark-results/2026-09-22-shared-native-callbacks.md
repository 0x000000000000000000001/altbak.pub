# Shared native callback results: fewer allocations, no demonstrated time gain

Three bounded probes on the real generated JSON decoder reduce decoded
allocations by up to **6.1%**, but do **not** establish a decoding speedup.
These callback changes are not integrated. The README remains at the validated
parser campaign's **17.00 ms Go / 9.30 ms JS** combined JSON result.

## Scope

The probe starts from the frozen post-parser `JsonDecoding` application and
changes generated Go only. It retains the real PureScript decoder and one shared
`gDecodeJsonCons` worker for every row layout; no schema-specific decoder is
introduced.

1. **Shared row callback:** `GDecodeJson.gDecodeJson` gets a native entry returning
   `Either` as the existing native shape: two `Value` slots and a success flag
   (56 bytes). Recursive calls use this entry. A boxed entry remains available
   for consumers using the existing `Value` calling convention.
2. **Shared field callbacks:** extend the native return to `Maybe(Either)` in
   `decodeFieldId` and `decodeFieldMaybe`, immediately consumed by that row worker.
   This also removes their output `Just` wrapper. Their input and underlying
   decoder calls still use `Value`.
3. **Compact result:** the two `Either` payloads have the same Go type, so share
   one `Value` slot plus a tag (32 bytes). This tests a possible cost of large
   native returns without changing decoded values or the worker algorithm.

**Record payloads remain `Value`; no native record layout or typed output array
is introduced.** FFI lookup still returns a boxed `Maybe`; field decoders still
return a boxed `Either` upstream; `Record.insert` and its dynamic-record copies
are unchanged. The native adapters use exact function-type assertions in this
isolated probe, not a finished general runtime ABI.

## Why the cumulative Apply profile is not a dispatch cost

The refreshed profile has `Apply2` at **99.82% cumulative allocated bytes**, but
only **1.08% own allocated bytes**. Its callees account for the rest, including
record construction, lookup wrappers and decoder results. Removing the dispatch
function cannot be interpreted as removing that cumulative percentage.

In the CPU view restricted to benchmark call stacks, `Apply2` contributes
**0.92% own CPU**. That percentage is relative to the selected 9.74 seconds of
samples, not the entire process profile or all GC work. The separate complete
CPU profile contains 18.13 seconds of samples. These are different denominators;
none establishes a 90% intrinsic dispatch bottleneck. Profile summaries and the
profile manifest are retained in the result archive.

## Paired results

Each row has its own two reversed Go process pairs: **control, treatment,
treatment, control**. Times are the median of the two process minima, with five
samples per process phase. Allocation bytes are the median across all samples.
All 17 fixed JSON oracles pass in every process; five successful cases are timed.

| Variant | Control decode ms | Native decode ms | Control bytes | Native bytes | Allocation change |
|---|---:|---:|---:|---:|---:|
| Shared row `Either` | 13.130 | 13.155 | 19,714,272 | 19,221,592 | -2.50% |
| Plus field `Maybe(Either)` | 14.318 | 15.266 | 19,714,272 | 18,513,488 | -6.09% |
| Compact native result | 13.096 | 13.156 | 19,714,272 | 18,513,488 | -6.09% |

Individual decode process minima, in milliseconds:

| Variant | Controls | Treatments |
|---|---|---|
| Shared row | 13.390, 12.870 | 13.233, 13.076 |
| Plus field | 13.262, 15.374 | 14.830, 15.701 |
| Compact result | 12.699, 13.492 | 13.123, 13.189 |

The middle campaign had substantial session variability: even unchanged parsing
moved from 2.365 to 2.899 ms between its control processes. Thermal or background
activity is a possible explanation, not a measured attribution. Two process
pairs are sufficient to reject a claim of a demonstrated large gain here, not
to establish a precise small regression. The compact campaign also provides no
positive timing evidence despite the stable allocation reduction.

Combined medians are respectively **15.931 → 15.648 ms**, **19.426 → 19.048 ms**,
and **15.732 → 16.364 ms**. They move in different directions; the combined phase
is independently measured and is not the sum of the separate phase minima.
These short probes must not replace the longer README measurement campaign or
be interpreted as a new Go/JS comparison. JS, TAST and b8x were not rerun.

## Provenance and validation

- Go **1.27.0**, darwin/arm64; `GOMAXPROCS=1`, `GOGC=100`, PGO off; two warmups
  and five measured samples per phase. Builds and profiling completed before
  timing. Each process validates the official corpus order and complete decoded
  and JSON fingerprints through `bin/benchmark/json-diagnostic.py`.
- Control: `scratch/json-parser-20260922/final-json/benchmark`. Its SHA-256 still
  matches the original manifest:
  `154675c3b349c812e02b13682bb92029770d155d6ccaac099328112629ac1018`.
- Comparing the archived manifest's 495 source fingerprints with current sources
  finds **only `gopurs/bin/gopurs-native` changed**. Its archived hash starts
  `5d446523`, current hash `e74a615c`; full hashes are in the JSON archive.
  This campaign deliberately compares a verified frozen binary and Go modified
  from that same frozen output. It does not claim to have regenerated either
  application with the currently installed compiler.
- Final compact binary: **6,421,746 bytes**, against control **6,420,994 bytes**.
  The first two probe binaries were overwritten between variants; their exact
  patches and raw reports are retained, but no historical binary hash or size
  is asserted for them.
- All **12 application processes** pass the **17-case oracle**. These checks
  exercise the benchmark's success/error semantics; they are not a complete
  regression suite for a new generic calling convention.

[Raw results, original manifest, source differences, final binary fingerprints
and reproduction sources](2026-09-22-shared-native-callbacks.json) are archived
alongside this report. Local scripts and generated output remain under
`scratch/json-shared-native-20260922/probe/`.

For local reproduction, `patch.py` recreates the shared-row variant from the
frozen output; run `extend.py` for the second variant, then `compact.py` for the
third. After each selected transformation:

```sh
# From scratch/json-shared-native-20260922/probe/output:
go build -pgo=off -o ../benchmark ./main
# From scratch/json-shared-native-20260922/probe:
python3 run.py
```

The four script bodies are included in the archive's `reproduction_sources`.
The scripts are specific to the fingerprinted generated code, and intentionally
fail or require review when that code changes.

## Decision

Do not integrate these adapters merely to remove the tested envelopes: their
allocation reduction has not translated into a demonstrated timing benefit.
The result **does not refute a complete native ABI**. That larger design would
also preserve arguments, record payloads and array elements through producers
and consumers, eliminating boundaries explicitly left boxed by this probe.
Its benefit remains unmeasured.
