# Fused Either traversals: generated JSON and TAST decoders

The final native compiler reduces JSON decoding time by **35.73%** and
parse + decode by **21.56%** in the paired campaign. TAST decoding improves by
**10.12%**, and its combined phase by **6.20%**. These are measurements of the
existing PureScript benchmarks compiled by gopurs, without a manual schema
engine or changes to their workload.

## Final measurements

Five alternating before/after Go process pairs; each cell is the median of
process minima. Times are milliseconds. Combined phases are measured separately
and need not equal the sum of the two separate phase minima.

### JSON Decoding

| Phase | Go before | Go after | Change | JS current |
|---|---:|---:|---:|---:|
| parse | 7.460 | **7.534** | +1.00% | 1.689 |
| decode | 22.752 | **14.621** | -35.73% | 7.878 |
| combined | 31.791 | **24.938** | -21.56% | 9.182 |

The combined Go/JS ratio falls from **3.46 to 2.72** in this session.
The parser is unchanged; it remains about 7.5 ms versus 1.7 ms in JS.

### JSON to Typed AST

| Phase | Go before | Go after | Change | JS current |
|---|---:|---:|---:|---:|
| parse | 76.072 | **75.421** | -0.86% | 20.439 |
| decode | 428.356 | **384.991** | -10.12% | 61.216 |
| combined | 531.295 | **498.367** | -6.20% | 83.423 |

The TAST combined Go/JS ratio remains **5.97**. Its combined times vary more:
before **505.895–578.681 ms**, after **474.589–518.145 ms**. All five paired
combined measurements improve, but the median gain is modest. Decode ranges
are **398.392–452.149 ms** before and **383.592–389.568 ms** after.

The unchanged JSON JS bundle is measured three times. PBO changed for TAST,
so both JS versions are measured three times: combined **83.044 → 83.423 ms**;
this campaign does not demonstrate a JS combined improvement.

### Allocated bytes per corpus

| Suite / phase | Before | After | Change |
|---|---:|---:|---:|
| JsonDecoding, decode | 30,423,384 | 19,714,272 | -35.20% |
| JsonDecoding, combined | 36,421,504 | 25,712,168 | -29.40% |
| JsonTypedAst, decode | 476,629,064 | 442,632,472 | -7.13% |
| JsonTypedAst, combined | 545,062,296 | 511,065,584 | -6.24% |

Parser allocations are effectively unchanged. Binary sizes also remain stable:
JSON **6,403,746 → 6,403,954 bytes**; TAST **9,027,362 → 9,024,514 bytes**.

### README history

Previous published cells were JSON **27.87 ms Go / 8.97 ms JS**, TAST
**505.42 ms Go / 80.94 ms JS**. This campaign publishes its actual final cells:
JSON **24.94 / 9.18 ms**, TAST **498.37 / 83.42 ms**. The contemporaneous Go
controls were **31.79 / 531.30 ms**, so the paired percentages above must not
be inferred from the older README cells. No historical cell was multiplied by
an estimated gain. The excluded WIP rows remain outside totals and /C ratios.

## Integrated changes

1. **Object/Either traversal:** a single fresh output map replaces repeated
   immutable insertions that copied the accumulating map. The compiler accepts
   only the standard qualified Object traversal and Either dictionaries.
   The input stays immutable; callbacks keep their sorted native fold order,
   all run even after failure, and the first error is retained.
   `traverseWithIndexDefault` stays generic because its existing callback order
   differs. Partially applied workers allocate a fresh output on every call.
2. **Native adjacent callbacks:** Array/Either loops directly consume literal
   unary or indexed callbacks with their generated result representation.
   Native Either and record payloads can survive to the consumer. Staged lambdas
   and unknown callbacks retain the ordinary path.
3. **Nonindexed Array/Either coverage:** qualified traversal, the known static
   dictionary field, and the FFI worker with the exact standard methods are
   recognized. All **19** targeted TAST sites are replaced: 9 in Data.Traversable,
   6 in CoreFn.TypeTable, 4 in CoreFn.Json. Four of these loops call a native Go
   callback directly; one also retains a native Either result. The other results
   remain Value. This is partial coverage, not a complete native callback ABI.
4. **PBO field errors:** CoreFn.Json's `getField` and `getFieldOptional'` construct
   `AtKey` only after a decoder failure. Missing/null handling, nested errors,
   and decoder call counts stay unchanged. TypeTable helpers are untouched.

An intermediate paired campaign with changes 1–2 improved JSON but left TAST
neutral (combined **542.679 → 545.042 ms**, nearly identical allocations).
A bounded generated-code probe then established the value of the 19 nonindexed
sites before changes 3–4 were integrated. The final campaign measures the whole
change; individual probe percentages are not added together.

## Validation and protocol

- Native compiler bootstrapped from the local TAST fork; benchmark programs and
  JS controls rebuilt in fresh workspaces. Go 1.27.0, darwin/arm64.
- `GOMAXPROCS=1`, `GOGC=100`, PGO off, two warmups and five samples per phase;
  file reads, oracle validation and fingerprints outside timings.
- Original frozen control binaries and final treatment binaries alternate;
  no build, profile or test from this task runs during timed campaigns.
  Binary, source and corpus hashes are archived; final source fingerprints are
  checked before and after each campaign.
- Full independent oracles pass in every process: **17 JSON cases** and
  **12 TAST modules**, including errors and complete decoded values.
- Compiler tests: **171 pass, 2 skipped, 0 failures**. Array traversal tests
  include **11,055** strict traversal comparisons; Object tests include **270**.
- Four fixtures pass with the final native compiler and the actual TAST fork:
  ArrayTraverseEither, NativeTraverseCallback, ObjectTraverseEither and
  NativeArrayReboxing. The first three also pass through the JS compiler.
- PBO: **7** field-helper tests and **12** source-usage tests pass on freshly
  compiled JS. Helpers remain private to the production API.

## Remaining work

The full shared dictionary/callback ABI remains open: cached callbacks still
cross Value boundaries, final arrays are `[]Value`, and generic record results
can still be adapted at worker boundaries. CoreFn.Json's hand-written ST array
decoder is not a `traverse` and is not rewritten by these rules. A fresh TAST
profile should choose its next target. Cumulative `Apply` cost includes callees;
it does not identify the intrinsic dispatch cost by itself.

The parser was not changed. There is no new full b8x compilation measurement,
and no claim that the manual 8–9 ms prototype has been reproduced by gopurs.
The existing local gopurs-argonaut-codecs override is present in both controls
and treatments; this campaign does not change its package distribution.

## Reproduction and sources

[Raw results, manifests, paired runner and changed production sources](2026-09-22-fused-json-traversals.json).
[Compiler contract](../../../gopurs/gopurs/docs/json-traversals.md).
Local workspaces and logs: `scratch/json-shared-callback-20260922/`.
The archived runner imports the versioned `bin/benchmark/json-diagnostic.py`;
it uses the same corpus, oracles and cell calculation with a frozen control.

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --workspace /absolute/new-json
python3 bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace /absolute/new-tast
python3 /absolute/campaign.py --suite JsonDecoding --before /absolute/before-json --after /absolute/new-json --output /absolute/new-json-results
python3 /absolute/campaign.py --suite JsonTypedAst --before /absolute/before-tast --after /absolute/new-tast --output /absolute/new-tast-results
```
