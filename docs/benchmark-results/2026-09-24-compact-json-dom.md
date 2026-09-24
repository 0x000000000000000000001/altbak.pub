# Compact JSON objects and fused field accessors

This cycle targets the actual remaining native TAST path, then the generated
application decoder. The preserved starting point is the final build from
[the dictionary/callback follow-up](2026-09-24-json-codegen-follow-up.md).

The final paired campaigns reduce **TAST combined time by 43.3%** and
**application JSON combined time by 30.2%**. Combined allocations fall by
**38.1%** and **34.6%**, respectively. The official Go README cells are now
**36.31 ms for JSON → Typed AST** and **4.95 ms for JSON Decoding**.

## Final paired results

Each campaign uses six processes per workspace, alternating workspace order,
two warmups and five samples per phase. The application campaign additionally
alternates generated/specialised decoder passes in each process and rotates
phase order. TAST uses the preserved official retained-output driver. All times
are medians of process minima; comparisons are within the same campaign.

| Workload / phase | Before | Final | Delta |
|---|---:|---:|---:|
| TAST parse | 27.494 ms | 20.292 ms | −26.2% |
| TAST decode | 15.838 ms | 13.627 ms | −14.0% |
| **TAST combined** | **61.755 ms** | **35.045 ms** | **−43.3%** |
| Application JSON parse | 2.506 ms | 2.033 ms | −18.9% |
| Application JSON decode | 2.833 ms | 2.044 ms | −27.9% |
| **Application JSON combined** | **6.613 ms** | **4.615 ms** | **−30.2%** |

| Workload / allocated bytes per corpus | Before | Final | Delta |
|---|---:|---:|---:|
| TAST parse | 56,031,272 | 24,405,672 | −56.4% |
| TAST decode | 27,071,160 | 27,071,192 | +32 B |
| **TAST combined** | **83,102,368** | **51,476,800** | **−38.1%** |
| Application JSON parse | 4,612,800 | 2,435,696 | −47.2% |
| Application JSON decode | 5,004,680 | 3,850,656 | −23.1% |
| **Application JSON combined** | **9,617,480** | **6,286,352** | **−34.6%** |

Allocation cells are minima over the recorded samples, not live-heap sizes.
All three phases improve in all six final pairs for both workloads. Individual
combined deltas range from **−47.8% to −31.9% for TAST** and **−32.7% to −19.6%
for application JSON**. The application audit checks **all 1,800 sampled
outputs** against frozen fingerprints; the TAST pair checks **2,160 sampled
module outputs**, as well as warmups and setup. Each TAST process additionally
checks all twelve setup fingerprints against the frozen oracle through the
runner.

Background load remains relevant: one-minute load averages range **19.6–22.2**
in the application campaign and **5.7–19.6** in the final TAST campaign. The
TAST combined process minima span **33.902–44.847 ms** on the final build.
Its allocation count is unchanged between the parser and accessor stages;
subtracting the two campaigns' time cells would not isolate an accessor-fusion
effect on TAST. Decode-only timing can also change when the retained DOM and
GC working set change.

Artifacts: `paired-final-json/results.json`, `paired-final-tast/results.json`
and `summary.json` under `var/benchmark/json-parser-20260924/`.

### Isolated application accessor stage

The separate six-pair comparison from compact DOM to compact DOM plus fused
accessors measures **2.842 → 2.207 ms decode (−22.4%)** and
**5.940 → 5.235 ms combined (−11.9%)**. Both improve in every pair. Parse is the
control: **2.099 → 2.086 ms**, with identical **2,435,696 B** allocation.
Decode allocation falls **5,064,248 → 3,850,656 B (−24.0%)**; combined falls
**7,499,944 → 6,286,352 B (−16.2%)**. Load averages range **21.1–22.4**.

Artifact: `paired-fusion-json/results.json`. This ordinary-driver campaign has
a different allocator schedule from the alternating-reference campaign above.

### Aligned specialised reference

In the final application campaign, the specialised decoder measures
**0.542 ms decode / 2.759 ms combined**, allocating
**1,843,976 / 4,279,672 B**. Its own baseline measures **0.739 / 3.450 ms**,
with **1,843,976 / 6,456,776 B**: it also benefits from the compact parsed DOM.
The generated/specialised decode ratio changes only **3.84× → 3.77×**, despite
the substantial generated-decoder improvement.

The final reference's combined cell is **40.2% lower** than generated decoding
in that same campaign. This measures the remaining gap for this schema and
these implementations. Phase medians should not be added into a combined-time
floor, and cells from earlier runner schedules should not be substituted here.

## Official cells and controls

Three processes per backend, ordinary official schedule, same frozen oracles:

| Workload / backend | Parse | Decode | Combined |
|---|---:|---:|---:|
| Application JSON / Go | 2.163 ms | 2.283 ms | **4.949 ms** |
| Application JSON / JS | 1.739 ms | 7.514 ms | 9.113 ms |
| Application JSON / C++ | 0.379 ms | 0.270 ms | 0.665 ms |
| TAST / Go | 20.354 ms | 14.196 ms | **36.315 ms** |
| TAST / JS | 21.024 ms | 66.347 ms | 88.030 ms |
| TAST / C++ | 4.433 ms | 4.711 ms | 9.155 ms |

All existing Go, JS and C/C++ oracle checks pass. The reference lifecycle/scope
limitations below still apply. The README uses the two Go combined cells;
the excluded rows remain outside benchmark totals and `/C` ratios. Published
C/C++ reference cells remain the existing historical values.

Artifacts: `official-final-json/results.json` and
`official-final-tast/results.json`.

## Implementation

### Owned compact DOM

The initial TAST allocation profile attributes about **72% of allocated bytes
to JSON object construction**. This is an allocation attribution, not a CPU
share. The frozen corpus contains 121,418 objects: most have two or three fields;
only fifteen have more than eight fields.

`gopurs-argonaut-core/src/Data/Argonaut/Parser.go` now constructs an owned,
fixed-size object for up to eight distinct keys. A pointer to a small struct
stores the exact number of inline key/value pairs, avoiding both the minimum
hash-table bucket and a separately allocated slice header. Larger objects use
ordinary native maps. Duplicate keys retain their last value, including keys
that become equal after unescaping.

Small arrays accumulate in call-local scratch and allocate their exact output
slice once. Arrays longer than eight elements switch to a growing native slice.
Every published container owns its storage. Strings still own their bytes;
extracting a small child does not retain the source document. No global cache,
arena, input-buffer lifetime or new parser dependency is introduced.

The first slice-backed object prototype is preserved in `compact-stage` and
`compact-{json,tast}`. It reduced TAST parse allocation to 29,138,304 B but kept
the small-object slice-header allocation. The retained inline-object/small-array
variant reduces it further to **24,405,672 B**, from **56,031,272 B**.
Standalone parser allocation counts fall from about **1,249,486 to 1,040,069
objects per corpus**. Those standalone profiles exclude generated FFI wrappers.

### Borrowed access and explicit materialization

`gopurs/runtime/runtime.go` provides the immutable `JSONObject` protocol and a
`JSONObjectView` over compact objects, existing native maps and boxed records.
The CoreFn decoder in PBO and Argonaut record plans read through this view.
They construct the same final typed values as before. The native TAST decoder
still resolves the type table eagerly and performs source-usage validation.

`Foreign.Object.lookup` reads the view directly. APIs requiring a map receive
a materialized container through `UnboxObject`/`RecordToMap`; JSON stringify and
comparison understand both representations. `thawST`, `freezeST`, insertion,
deletion, folds and enumeration remain usable through the public APIs. Their
conversion/copy costs are incurred when those operations run.

The standard `Foreign.Object Json` decoder still creates its independent output
map. The event decoder performs **2,477 such owned copies per corpus**. This
optimization does not turn that public result into a borrowed mutable map.

### Fused application field accessors

`gopurs-argonaut-codecs` adds native implementations of `getField`,
`getFieldOptional` and `getFieldOptional'` behind the existing public functions.
A recognized decoder tag allows lookup and decoding to build just the final
accessor result, eliminating lookup's `Just` and a temporary `Either`.
Untagged/custom decoders use the complete original accessor. JavaScript uses
the passed PureScript fallback.

Missing fields, explicit nulls, nested `Maybe`s, `AtKey`/`AtIndex`/`Named` errors
and first-error behavior retain their previous meanings. Record-plan symbols
are still reevaluated during lookup and reverse-order insertion.
Value-level object access also removes avoidable interface boxes introduced by
the first view implementation.

Untimed instrumentation confirms **7,431/7,431 public field accesses use the
fused path** on the five timed application cases. The previous counts remain:
2,477 custom event decodes, 6,350 nested recognized record decodes, 6,355 compact
output records and zero per-corpus identity-object decoder constructions after
initialization. These are execution counts, not CPU percentages.

## Parser-stage paired campaign

Six processes per workspace, alternating workspace order, two warmups and five
samples per phase. Cells are medians of process minima. All timed outputs are
retained through the end of their pass and fingerprinted afterwards.

| Workload / phase | Before | Compact DOM | Delta |
|---|---:|---:|---:|
| TAST parse | 27.342 ms | 21.284 ms | −22.2% |
| TAST decode | 15.962 ms | 15.460 ms | −3.1% |
| TAST combined | 61.896 ms | 46.718 ms | −24.5% |
| Application JSON parse | 2.522 ms | 2.018 ms | −20.0% |
| Application JSON decode | 2.745 ms | 2.602 ms | −5.2% |
| Application JSON combined | 6.877 ms | 5.389 ms | −21.6% |

TAST combined improves in 5/6 pairs, with individual deltas from −40.6% to
+10.9%; application combined improves in all six pairs (−23.6% to −17.6%).
Parse improves in every pair for both workloads. One-minute load averages range
from 30.0–36.2 for TAST and 26.4–30.7 for application JSON. The small decode-time
movements at this stage are not evidence of a decoder algorithm improvement.

The intermediate application decode allocation rises by 59,568 B because of
avoidable object-view interface boxes. The final accessor stage removes those
boxes along with the lookup/result intermediates.

Artifacts: `paired-parser-tast/results.json` and `paired-parser-json/results.json`
under `var/benchmark/json-parser-20260924/`.

## Validation

- Parser differential tests under `-race`: **106 directed** and **5,357
  generated/mutated/truncated inputs** against `encoding/json`, plus **323
  valid shared-contract inputs** against `JSON.parse`.
- Boundary cases on both sides of the eight-key cutoff, duplicate/unescaped
  keys, retained recursive results, depth limit, invalid Unicode, negative zero,
  source-storage ownership, concurrent calls and callback panic propagation.
- Runtime `-race` tests cover lookup, null versus absence, independent map
  materializations, persistent record updates and existing representations.
- New `gopurs-argonaut-core/test/compact-dom.mjs` integration test passes in JS
  and freshly generated Go under `-race`: stringify/comparison, lookup, folds,
  enumeration, map/filter, insert/delete/union and ST freeze/thaw across sizes
  zero through twelve with nested values.
- Codec helper tests and `test/typed-plans.mjs` pass, including freshly generated
  Go under `-race`, optional/null/Maybe combinations, custom accessors, errors,
  nested plans, symbol reevaluation and result ownership.
- **2,210 deterministic TAST mutations**: **1,493 errors and 717 successes**,
  comparing generated PureScript decoding, native compact-DOM decoding and
  native map-DOM decoding. Exact errors and successful fingerprints agree.
  The twelve unmodified module fingerprints and 22 type-table argument/error
  cases also pass. These tests ran on both the parser-stage and final builds.
- Final compiler bootstrap: **459 modules, 275,790 type-table entries**.
  The rebuilt native compiler regenerates **564/564 Go files byte-identically**
  to Node, including its own compilation workload.
- Both specialised audit builds pass the regression that deliberately corrupts
  an exact timed output and requires the runner to reject it.

## Ownership, timing and remaining costs

All campaigns keep `GOMAXPROCS=1`, `GOGC=100` and diagnostic PGO disabled.
The runners sanitize profiling/GC overrides. File I/O, the pass's retention-buffer
allocation and fingerprinting are outside timing in both builds. Construction,
conversion and copying performed by parse/decode are inside timing; ordinary
Go GC costs occur under the same policy. The retained-output protocol does not
force and charge a complete final heap drain after each pass.

Distribution cost is modest in these builds: the application diagnostic binary
grows **6,492,498 → 6,554,962 B (+62,464 B, 1.0%)**, the TAST diagnostic
**8,746,178 → 8,789,058 B (+42,880 B, 0.5%)**, and the native compiler
**27,733,842 → 27,812,562 B (+78,720 B, 0.3%)**. Both diagnostic `go.mod` files
are byte-identical to their baselines. No CGo, SIMD runtime or external parser
installation is required by the new Go path.

The final allocation profile still identifies owned object copies, final
records, arrays, result envelopes, accessor reboxing and event callbacks.
Its setup/validation allocations are present too; its times are not used as
benchmark cells. CPU profiling remains unreliable on this machine, so no CPU
share is inferred from this allocation profile.

The schema-specialised reference now reads both compact objects and native
maps through the same borrowed-access implementation. Both comparison builds
use that source. It still constructs the normal final PureScript values; it is
an implementation reference, not an implemented general specialization pass.

The earlier reference-audit limits remain: C++ has a different retained-output
lifecycle and incomplete TAST usage-validation work; the native-Go reference
does not yet validate every timed output; simdjson-cgo node counting does not
measure transfer to owned usable Go values. Broader application benchmarking
beyond compiler bootstrap remains pending.

## Sources and reproduction

Preserved root: `var/benchmark/json-parser-20260924/`.

- `before.json`, `library-before/`, `compiler-before`: starting hashes and parser
  sources; original diagnostic binaries remain in `json-dec-codegen-20260924`.
- `parser-before/`, `parser-compact/`, `parser-inline/`: standalone allocation
  profiles and their exact source/build/corpus inputs.
- `final-{json,tast}`: parser-stage builds; `fused-{json,tast}`: final builds.
- `baseline-audit/`, `fused-audit/`, `fused-coverage/`, `fused-profile/`: retained
  reference builds, untimed path counts and the allocation profile.
- `tast-validation/`, `fused-tast-validation/`: differential test sources,
  generated workspaces and logs.
- `native-self-check/`, `bootstrap-final-logs/`: Node/native equivalence and
  compiler build records.
- `sources-final/`, `sources-final.sha256.json`: final library/runtime/PBO
  sources, tests and compiler binaries. The native libraries are not Git repos.

Final compiler SHA-256:
`7c8feaf81bfff0dc17794523818b94c41287b83191c15a2da593da03e5117c7b`.
Node bundle SHA-256:
`3d282efa10ed9f2551353d84867806f736ed8a23da7d50853046f2e92ddab034`.

Rebuild with `npm run build:native -- --keep-workspace` in `gopurs/gopurs`, then
use `bin/benchmark/json-diagnostic.py build --suite <suite> --workspace <new>`.
`bin/benchmark/json-diagnostic/paired.py` compares preserved official Go builds
via `--baseline`, `--current`, `--suite`, `--output` and `--processes 6`.
It pins binaries, manifests and generated Go before and after the campaign.
Application phases rotate/reverse; the preserved TAST driver keeps its fixed
parse/decode/combined order, which the report explicitly records.

The specialised comparison uses `specialized/run.py setup` for each source
workspace, then `campaign --workspace <final> --compare-workspace <baseline>
--output <new> --processes 6`. Official cells use the ordinary runner's
`measure --suite <suite> --workspace <final> --output <new>`.
