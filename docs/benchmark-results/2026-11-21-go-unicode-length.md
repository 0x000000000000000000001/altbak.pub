# Go Unicode code-unit length — November 21, 2026

`string/cu-length-unicode-1000` now takes **0.009666 µs**, down from
**2.522969 µs** in the matched before binary: **261× faster**. Its Go / JS ratio
falls from **250.62× to 0.96×**. This is the cost of reading the length of an
already constructed string; the construction cost is measured separately below.

## Matched measurements

Times are in microseconds. Each result is the median of three independent
process minima, with the four frozen programs run sequentially in rotating
order. All builds and diagnostic tests finished before this campaign.

| Program | Process 1 | Process 2 | Process 3 | Median |
| --- | ---: | ---: | ---: | ---: |
| Go before | 2.431630 | 2.536143 | 2.522969 | 2.522969 |
| Go after | 0.009719 | 0.009666 | 0.009598 | 0.009666 |
| Official JS | 0.010085 | 0.010067 | 0.009936 | 0.010067 |
| PBO JS | 0.008709 | 0.008348 | 0.008974 | 0.008709 |

The historical README records **2.66 µs / 266× JS**. The preceding
[library campaign](2026-11-21-go-library-ratios.md) reduced that to 2.562683 µs
but retained a 243× JS gap. This campaign remeasures that preceding binary as
its before control. Both current Go / JS ratios use the same freshly measured
JS denominator. The new Go / PBO ratio is **1.11×**; differences of about one
nanosecond here should not be interpreted as a broad backend ranking.

The unchanged benchmark prepares 1,000 repetitions of `a😀é` outside the timed
action and passes the resulting string through `Bench.opaque` on every call.
The returned length remains **4,000 UTF-16 code units**. Its source, fixtures,
opaque boundary and oracles were not modified.

## What changed

The runtime can retain a string's known UTF-16 length as immutable metadata
inside the existing `Value.IntVal`. The existing pointer still references the
same immutable bytes, and `Value` remains **24 bytes**. Metadata adds no
allocation and uses no global cache or mutable state.

- Boxed string literals receive their length at compile time.
- `JoinWith` counts the completed string once and retains its length. Counting
  the final bytes preserves the existing behavior when malformed fragments
  combine into a valid UTF-8 sequence at a join boundary.
- The compiler preserves this metadata across immediate string unbox/rebox
  conversions, evaluating the source expression once.
- `Data.String.CodeUnits.length` reads the metadata when present. Ordinary
  native-string boxing remains constant time; strings without metadata use the
  existing UTF-8/WTF-8 scan. Byte lengths above 2³¹−1 also use that fallback.

This is an optional representation improvement, not a guarantee of constant
time for every string. A genuine round trip through a native Go `string` can
discard metadata. Code-point length is a separate operation and is unchanged.

## Construction cost

The following native Go microbenchmark uses the production `JoinWith` and
`Length` implementations on the same 1,000 `a😀é` fragments. It compares them
with the previous `strings.Join` plus length scan. Both variants return a boxed
string to the same sink. Values are medians of three 200 ms measurements.

| Work timed together | Before, µs | After, µs |
| --- | ---: | ---: |
| Construction, no length read | 5.477 | 8.046 |
| Construction + one length read | 7.985 | 8.013 |
| Construction + ten length reads | 31.117 | 8.036 |

All six variants allocate **8,192 bytes in one allocation**. The eager scan
makes construction alone **47% slower** for this input. Construction plus one
read is essentially unchanged; ten reads amortize the scan and are **3.87×
faster** overall. Consequently, the 261× isolated length gain must not be
presented as a 261× improvement in constructing and processing strings.

## Validation and scope

All **444 outputs** passed the existing independent oracles: 14 core and 23 WIP
results in each of 12 processes. WIP results remain excluded from aggregate
scores, and no C timing was added. The 14-case core total is
**12.876 → 12.715 ms**, compared with the historical README's **13.28 ms**;
this small difference does not establish a separate core speedup.

The candidate is **not regression-free** in these measurements.
`record/set-medium-1000` rises from **5.800 to 6.221 µs (+7.27%)**.
A second alternating campaign with only the same two Go binaries confirms
**5.708 to 6.098 µs (+6.82%)**, with another 222 correct outputs. Its cause
has not been established. In that check, array map differs by +0.33%, append
by +0.95%, Effect ref write by +1.81% and Effect ref modify by +3.27%; the core
total is 12.562 → 12.641 ms. Both campaigns are retained in the archive.

Semantic checks passed for JS parity fixtures, all 65,536 two-byte inputs,
10,000 deterministic random byte strings, ASCII boundaries, supplementary
characters, isolated surrogates and malformed join boundaries. Runtime tests
cover metadata bounds and fallback, garbage-collector retention, zero-allocation
boxing, `Box`/`AnyToValue` preservation and the unchanged `Value` size.
Compiler tests execute generated conversions and cover literal lengths,
evaluation count and adjacent representation conversions. Race checks passed.

The JS compiler builds with zero warnings and errors. The native compiler was
rebuilt and used to produce the final measured application. Its first Go link
attempts failed during a period of low disk space; freeing this task's temporary
builds allowed the same generated source to build successfully. Only the
successful binary was measured.

## Evidence and reproduction

The [JSON archive](2026-11-21-go-unicode-length.json) contains the process
results, historical baseline, construction samples and raw log, frozen build
manifests, source hashes and the eight source/test changes relative to the
preceding campaign, plus the separate six-process Go side-effect check.
That preceding campaign's archive records the earlier
array/reference/string optimizations already present in both controls.
Full logs and binaries are in `var/benchmark/go-unicode-length-20261121/`.

Host: Apple M4 Pro, macOS 26.1, Go 1.27.0, Node 24.8.0. Application runs use
`GOGC=800`, PGO off, monolithic output and no explicit CPU affinity. The JS and
PBO payloads were built in the preceding campaign, verified unchanged by their
hashes, and remeasured here. The README keeps its historical multi-column table
and links this newer result separately.

```sh
# From altbak.pub, with the rebuilt native compiler installed:
GOPURS_GO_PACKAGES=0 GOPURS_GO_PACKAGE_OVERLAY=0 \
  python3 bin/native/driver.py go --wip-included --build-only \
  --build-dir path/to/new-build

# Reuse the frozen plan only while all recorded payload hashes match:
GOGC=800 python3 bin/benchmark/measure.py \
  --plan var/benchmark/go-unicode-length-20261121/plan.json \
  --output path/to/new-measurements

# From gopurs-strings, measure construction as well as length reads:
node test/code-units-native.mjs -run '^$' \
  -bench '^BenchmarkJoinedCodeUnits$' -benchtime=200ms -count=3 -benchmem
```
