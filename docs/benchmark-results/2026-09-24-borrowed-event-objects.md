# Borrowed event objects and an indexed TAST experiment

**Application decode improves by 23.9%, combined by 15.1%, and combined
allocation falls by 20.7% in the ordinary paired campaign.** The new official
Go JSON Decoding cell is **3.91 ms**. The indexed TAST candidate is rejected:
its end-to-end cost exceeds the compact-DOM path.

This cycle starts from the preserved final builds in
[compact JSON DOM and fused field accessors](2026-09-24-compact-json-dom.md).
It implements compiler-proven borrowing of temporary identity-decoded objects
and removes immediately applied object/tag callbacks. It also evaluates a
call-local structural index feeding the existing complete native TAST decoder.

## Retained application optimization

Six paired processes, alternating workspace order and rotating/reversing phase
order, two warmups and five samples per phase. Cells are medians of process
minima; allocation cells are minima over the campaign. `GOMAXPROCS=1`,
`GOGC=100`, PGO disabled, all exact timed outputs retained and fingerprinted.

| Application phase | Before | After | Delta |
|---|---:|---:|---:|
| Parse control | 2.124 ms | 2.218 ms | +4.4% |
| Decode | 2.320 ms | 1.766 ms | **−23.9%** |
| Combined | 4.801 ms | 4.078 ms | **−15.1%** |

| Allocated bytes per corpus | Before | After | Delta |
|---|---:|---:|---:|
| Parse | 2,435,696 | 2,435,696 | 0 |
| Decode | 3,850,656 | 2,549,944 | **−33.8%** |
| Combined | 6,286,352 | 4,985,640 | **−20.7%** |

Decode and combined improve in all six pairs. Individual combined deltas range
from −20.5% to −3.7%. One-minute load averages range from 11.4 to 14.4, so the
timing spread remains meaningful. Parser code and allocation are identical;
the parse-control movement is not a parser change. This campaign measures the
two compiler/library changes together, rather than assigning time savings to
the removed maps alone. It validates 900 sampled outputs, plus warmups and
setup, against the frozen corpus.

Artifacts below are relative to `var/benchmark/json-direct-20260924/`:
`paired-borrow-json/results.json`, `borrow-json/`, and `final-json/`.
The final rebuild differs only in Go formatting from the paired current build:
all 326 emitted Go files agree after `gofmt` (`distribution.json`).

### Aligned specialised reference

A second six-pair campaign also alternates generated/specialised decoding in
each process, using the retained-output audit and its corruption-detection
regression test. It validates all **1,800 sampled outputs**, plus warmups/setup.

| Decoder / phase | Before | After |
|---|---:|---:|
| Generated decode | 2.490 ms | **1.793 ms** |
| Specialised decode | 0.650 ms | 0.632 ms |
| Generated combined | 5.136 ms | **4.110 ms** |
| Specialised combined | 2.917 ms | 2.849 ms |

The generated/specialised decode ratio narrows **3.83× → 2.84×**. The final
specialised combined cell is **30.7% lower** than the generated cell in this
same campaign. This is schema-specific measured headroom, not a guaranteed
compiler-specialization target. Its allocations remain **1,843,976 B decode /
4,279,672 B combined**, versus **2,549,944 / 4,985,640 B** generated.

Generated decode improves in every pair; combined improves in five of six.
Load averages range 10.9–12.9. This runner has a different allocator schedule
from the ordinary paired driver; its medians must not replace the cells above.
Artifact: `paired-reference-json/results.json`.

### Official cells

Three processes per backend on the frozen corpus:

| Backend | Parse | Decode | Combined |
|---|---:|---:|---:|
| Go | 2.148 ms | 1.739 ms | **3.905 ms** |
| JS | 1.755 ms | 7.709 ms | 8.909 ms |
| C++ reference | 0.387 ms | 0.261 ms | 0.658 ms |

All existing oracle checks pass. README now publishes **3.91 ms** for Go JSON
Decoding. TAST keeps its production compact-DOM path and existing **36.31 ms**
cell. The three excluded rows retain their existing treatment outside totals
and `/C` ratios. Published historical reference cells retain their provenance.
Artifact: `official-final-json/results.json`.

### Compiler proof and native helper

`Gopurs.ImmediateApplications` now allows multiple uses of an already evaluated
local reference or scalar. It still rejects multiple uses of a substituted
callback lambda, moving evaluated calls/effects, recursive initialization
scopes, and excessive syntax growth across branches. This exposes the event decoder's object and tag
continuations to ordinary direct Go control flow and unboxed branch results.

The new `Gopurs.BorrowedObjects` pass runs before ownership/TCO preparation. It
recognizes the exact qualified standard
`decodeForeignObject decodeJsonJson` dictionary, including nonrecursive local
module getters. It accounts for every use of the result envelope and its
successful payload. Only synchronous, saturated public field readers can
consume that payload. Returning, storing, capturing, mutating or passing the
object to an opaque function prevents borrowing. The pass retains annotations
and rejects ambiguous qualifications, partial calls and recursive scopes.

The internal `borrowObject` helper checks the original decoder's native identity
tag again. Successful object inputs reuse their existing representation;
invalid inputs and unknown/custom decoders execute the original decoder.
Ordinary `typedObject` application still allocates independent object storage.
The JavaScript helper simply calls the original decoder.

Untimed instrumentation (`borrow-coverage/counts.json`) confirms:

- **2,477/2,477** event object decodes borrow their temporary container;
- **zero** owned identity-object copies in the timed application corpus;
- **7,431/7,431** public field accesses still take the fused accessor path;
- 2,477 custom event dispatches, 6,350 nested recognized record decodes and
  6,355 compact final output records remain.

The counts distinguish the eliminated temporary copies from final owned
application results. There is no change to the record-plan symbol callback
contract, the default GC policy, or the frozen application data.

## Indexed TAST prototype

The experiment in `bin/benchmark/json-diagnostic/direct-tast/` builds a separate
preserved diagnostic. Only its combined operation changes. Parse-only and
decode-only remain ordinary public-DOM controls.

The candidate validates the complete source document first, including ignored
members and depth/number-range failures. A twelve-byte token index records
source spans and the next sibling after each subtree. Call-local cursors feed
the existing native decoder, with last-value precedence for duplicate/unescaped
keys and no assumption about object member order. Decoded strings own their
bytes. Invalid input uses the existing parser's exact error path.

The index is temporary and never becomes a published `Json`. Final values use
the same constructors, eager type-table resolution and source-usage validation
as the ordinary native TAST path. The cold `sourceSpan` subtree is materialized
at its existing generic Argonaut tuple-decoder boundary. This boundary was
identified by the first smoke run and fixed in `indexed-span-tast/`.

This is an indexed-input prototype, not a general complete-decoder compiler
specializer. Its compatibility adapter still boxes cursors into `any` and
materializes intermediate arrays. Those construction, conversion, copying and
GC costs belong inside the combined measurement.

### Final paired result: rejected

The final build uses the same `go build -pgo=off ... ./main` flags as the
preserved baseline. Six paired processes, fixed official TAST phase order,
two warmups and five samples per phase, same GC policy and retained outputs:

| TAST phase | Compact DOM | Indexed candidate |
|---|---:|---:|
| Parse control | 20.949 ms | 20.593 ms |
| Decode control | 14.563 ms | 13.710 ms |
| **Combined** | **44.022 ms** | **58.495 ms (+32.9%)** |

Combined allocations rise **51,476,800 → 56,330,248 B (+9.4%)**. Parse/decode
control allocations stay **24,405,672 / 27,071,192 B**. Combined time regresses
in five of six pairs. One-minute load averages range 8.2–18.7; the baseline's
combined process minima span 35.120–57.374 ms, so this campaign is not a new
production README cell. All **2,160 sampled module outputs** are validated,
plus warmups/setup. The index is not enabled in PBO or the production compiler.

Artifacts: `indexed-final-tast/` and `paired-index-final-tast/results.json`.
The earlier `indexed-span-tast/` pilot used `-trimpath` and measured
35.659 → 50.834 ms combined; it is retained separately. The final and validated
pilot's **506 compiled Go files agree after formatting** (`index-equivalence.json`).
This confirms the final flag-aligned build uses the same tested algorithm.

## Validation

- Compiler IR tests cover borrowing eligibility and rejection of escaping,
  deferred, recursive, partially applied, shadowed/ambiguous and custom producer
  cases, as well as repeated-local/scalar beta reduction. Existing dictionary
  caching tests and generated-Go callback/capture/effect tests pass.
- Codec helper tests pass under `-race`, including native identity borrowing,
  invalid/custom fallback behavior and independently owned public results.
  `test/typed-plans.mjs` passes its freshly generated Go `-race` execution and
  JavaScript fallback execution, including optional fields, nested errors,
  custom decoders and symbol callback reevaluation.
- An additional **829-case application differential** (17 frozen, 800
  deterministic mutations, 12 directed event cases) gives identical old/new
  Go fingerprints. **818** also agree with JS. The remaining **11** isolate the
  pre-existing native-64-bit versus JS-32-bit `Data.Int` boundary: otherwise
  valid payloads with `version = −2147483649` or `2147483648`. The complete old
  and new Go outputs agree there too. Exact differing cases are preserved in
  `application-differential/`; these are not new optimization discrepancies or
  part of the frozen timed successes.
- The indexed TAST candidate passes all twelve original fingerprints and
  **2,210 deterministic mutations: 1,493 errors, 717 successes**, against both
  generated PureScript and the ordinary native map-DOM decoder. The existing
  22 type-table argument/error tests pass alongside that differential suite.
- The index has **2,031 directed/generated/truncated parser-contract cases**
  against `encoding/json`, plus owned-string, duplicate-key and concurrent
  independent-result checks; the index-specific tests also pass under `-race`.
- Native compiler bootstrap succeeds on **460 modules / 276,475 type-table
  entries**. The rebuilt native compiler regenerates **559/559 Go files
  byte-identically** to Node on its own compilation workload.

## Distribution and remaining work

The application binary changes **6,554,962 → 6,555,090 B (+128 B)**. The native
compiler changes **27,812,562 → 28,003,794 B (+191,232 B, 0.69%)**. Diagnostic
`go.mod` files stay byte-identical. The retained optimization adds no external
parser, CGo or installation dependency.
The rejected, flag-aligned TAST diagnostic changes
**8,789,058 → 8,823,218 B (+34,160 B, 0.39%)**.

The event optimization is a narrow compiler nonescape/read-only analysis plus
ordinary callback reduction. A general whole-schema decoder specializer has
not been implemented in this cycle. The rejected index experiment shows that
an index behind the existing `any`/array adapter is insufficient; a subsequent
direct TAST design must eliminate those adapters and temporary containers too.
These are design conclusions, not predicted times or a language performance
limit.

Logs and source hashes are retained in the cycle directory. Fingerprinting and
the retention-buffer allocation occur outside timing, while parse/decode
construction and normal Go GC occur inside. As in the preceding campaign,
there is no forced and charged final heap drain after every pass. The C/C++
lifecycle and usage-validation limitations in the
[comparison audit](2026-09-24-comparison-audit.md) still govern cross-language
ratios; no universal performance floor is inferred here.
