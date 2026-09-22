# Direct native JSON parsing

The integrated Go parser reduces parsing time by **65% on both real corpora**.
The paired combined measurements improve by **25.33% for JSON Decoding** and
**14.63% for JSON to Typed AST**. The existing PureScript programs, decoders,
corpora and JS implementation are unchanged.

## Final application measurements

Five alternating before/after Go process pairs, three JS control processes.
Each cell is the median of process minima, in milliseconds. The combined phase
is measured independently; it need not equal the sum of separate phase minima.

### JSON Decoding

| Phase | Go before | Go after | Change | JS current |
|---|---:|---:|---:|---:|
| parse | 7.245 | **2.529** | -65.10% | 1.680 |
| decode | 13.403 | **13.590** | +1.39% | 7.656 |
| combined | 22.766 | **16.999** | -25.33% | 9.297 |

The combined Go/JS ratio falls from **2.45 to 1.83** in this session.
Parsing is now **1.51 times JS**, down from 4.31 times.

### JSON to Typed AST

| Phase | Go before | Go after | Change | JS current |
|---|---:|---:|---:|---:|
| parse | 76.631 | **26.764** | -65.07% | 20.656 |
| decode | 383.632 | **372.527** | -2.89% | 59.896 |
| combined | 494.614 | **422.267** | -14.63% | 76.797 |

The combined Go/JS ratio falls from **6.44 to 5.50**. Parsing is now
**1.30 times JS**, down from 3.71 times. Decode remains the dominant gap.
No decoder code changed. Its timing variation is not a separately established
optimization: decoded allocations are unchanged, and the per-process timing
ranges overlap. Do not add its observed delta to the parsing gain.

All five pairs improve both parse and combined phases for each corpus.
JSON combined ranges: **22.125–23.174 ms → 16.435–17.282 ms**.
TAST combined ranges: **458.173–495.776 ms → 391.868–426.424 ms**.

### Allocation and size

| Suite / phase | Bytes before | Bytes after | Change |
|---|---:|---:|---:|
| JsonDecoding, parse | 5,997,744 | 4,612,800 | -23.09% |
| JsonDecoding, decode | 19,714,272 | 19,714,272 | +0.00% |
| JsonDecoding, combined | 25,712,272 | 24,327,056 | -5.39% |
| JsonTypedAst, parse | 68,432,280 | 56,031,272 | -18.12% |
| JsonTypedAst, decode | 442,632,472 | 442,632,456 | -0.00% |
| JsonTypedAst, combined | 511,065,584 | 498,663,744 | -2.43% |

Binary sizes: JSON **6,403,954 → 6,420,994 bytes** (+0.27%);
TAST **9,024,514 → 9,058,066 bytes** (+0.37%).

Previous README cells were JSON **24.94 ms Go / 9.18 ms JS** and TAST
**498.37 ms Go / 83.42 ms JS**. Current measured cells are **17.00 / 9.30 ms**
and **422.27 / 76.80 ms**. The current Go controls were **22.77 / 494.61 ms**;
paired percentages use these controls, not older sessions. No historical value
was multiplied by an estimated gain. Excluded WIP rows stay outside totals.

## Implementation

Only the Go FFI of `Data.Argonaut.Parser` changes. It scans the input string
and constructs the existing `nil` / `bool` / `float64` / `string` / `[]any` /
`map[string]any` representation, without a separate whole-document validation
pass or conversion of the entire document to `[]byte`. The old stdlib path to
`any` was already specialized; this is not a claim that reflection was occurring
at every node, nor the removal of all runtime `Value` conversions.

Simple strings are cloned; escaped strings use an owned builder implementing
JSON escapes and UTF-16 surrogate pairing. Invalid UTF-8 and isolated surrogates
follow the existing Go replacement behavior. Keeping a tiny parsed string does
not retain the input document. The parser keeps `float64`, signed zero, duplicate
key behavior, non-nil empty arrays/objects, callback evaluation and local state.

Any unsuccessful fast parse falls back to `encoding/json.Unmarshal` on a fresh
destination, preserving exact errors and nesting rules. Invalid documents can
therefore perform extra work; the throughput gain targets successful parsing.
The pre-existing Go/JS differences for lone surrogates and floating overflow
are preserved. No dependency, schema specialization or public API is added.

## Short probes and the escaped-string correction

An initial bounded parser-only probe, three AB/BA/AB pairs within one process,
showed JSON **6.700 → 2.369 ms** and TAST **90.236 → 33.771 ms**. It used a GC
before each corpus outside timing and is not the README protocol. The integrated
measurements above use the actual PureScript/FFI application and normal GC.

The first implementation delegated each escaped string to `json.Unmarshal`.
Synthetic checks exposed **14–31% regressions** on escaped strings and arrays,
including 1,021 extra allocations per 1,024-element escaped array. Direct
unescaping was implemented and revalidated before the final application build.

Final synthetic medians, three Go benchmark repetitions per case, 200 ms each,
`GOMAXPROCS=1`, `GOGC=100`, identical callbacks, no concurrent workload:

| Input shape | stdlib ns/op | Native ns/op | Change | Allocs before / after |
|---|---:|---:|---:|---:|
| small-scalar | 160.80 | 26.52 | -83.5% | 4 / 2 |
| ascii-string | 33517.00 | 25907.00 | -22.7% | 5 / 2 |
| unicode-string | 76224.00 | 37804.00 | -50.4% | 5 / 2 |
| escaped-string | 158306.00 | 89370.00 | -43.5% | 6 / 2 |
| escaped-array | 144671.00 | 54816.00 | -62.1% | 3088 / 2061 |
| slash-array | 157455.00 | 71130.00 | -54.8% | 3088 / 2061 |
| surrogate-array | 137619.00 | 53840.00 | -60.9% | 3088 / 2061 |
| number-array | 125418.00 | 64266.00 | -48.8% | 2064 / 1036 |
| object-array | 178819.00 | 51758.00 | -71.1% | 2829 / 2313 |

These are coverage probes, not universal speed guarantees or README cells.

## Validation and protocol

- Seven focused Go tests pass on the actual production FFI, with and without
  `-race`: **106 directed cases**, **5,357 generated/mutated/truncated cases**,
  and **323 JS cases** on the common semantic domain. Comparison checks full
  types/values, float bits, exact errors and one callback per call.
- Additional checks cover 9,999/10,000/10,001 nesting levels, concurrent calls,
  callback panics and retained string pointers outside the source buffer.
- The existing Argonaut PureScript suite is regenerated and executed with the
  final FFI in an isolated workspace; all groups pass, including its ten parser
  round-trips. Its sibling-cache-clearing script is not needed.
- Every timed application process validates the full fixed oracles: **17 JSON
  cases** and **12 TAST modules**. Five JSON cases and all twelve TAST modules
  are timed. File reads, hashes and validation remain outside timed sections.
- Five alternating Go process pairs per suite; two warmups and five samples per
  phase; `GOMAXPROCS=1`, `GOGC=100`, PGO off; no task build/test/profile concurrent
  with timing. Three JS processes per suite; before/after JS bundles are
  byte-identical. Go 1.27.0, darwin/arm64.
- Before binaries are the verified frozen final artifacts of the previous
  traversal campaign. Final applications are rebuilt from source with the native
  compiler and updated FFI. Source, compiler, binary and corpus fingerprints are
  checked and archived. The compiler generator itself is unchanged.

## Scope and remaining work

This closes the direct-parsing work for generated applications. The polymorphic
decoder ABI, boxed callbacks and TAST construction remain separate tasks. This
campaign neither rebuilds the self-hosted compiler with the new library nor
measures full b8x compilation; those require their own validation before claiming
a compiler-load or end-to-end b8x gain.

[Parser contract and test commands](../../../gopurs/gopurs-argonaut-core/docs/native-parser.md).
[Raw application results, manifests, synthetic results, sources and validation logs](2026-09-22-native-json-parser.json).
Local workspaces and the archived paired runner are under
`scratch/json-parser-20260922/`. The runner uses the versioned
`bin/benchmark/json-diagnostic.py` corpus, oracles and cell calculation.

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --workspace /absolute/new-json
python3 bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace /absolute/new-tast
python3 /absolute/campaign.py --suite JsonDecoding --before /absolute/before-json --after /absolute/new-json --output /absolute/new-json-results
python3 /absolute/campaign.py --suite JsonTypedAst --before /absolute/before-tast --after /absolute/new-tast --output /absolute/new-tast-results
```
