# JSON decoding: shared imported dictionaries and immediate callbacks

This compiler follow-up to the [record-plan work](2026-09-24-record-plan-follow-up.md)
reduces generated decode allocation from **6,272,904 to 5,004,680 bytes per
five-case corpus (−20.2%)**. The six-process paired campaign measures
**3.906 → 2.926 ms decode** and **8.321 → 7.818 ms combined**. Background load
and control timings vary substantially; the combined timing improvement is
less conclusive than the allocation reduction.

Artifacts are preserved under `var/benchmark/json-dec-codegen-20260924/`.
The baseline is the previous final implementation, including typed record
plans and the identity-object shortcut. The application and codec library
sources are unchanged; all eighteen archived codec source/test files match
their previous hashes.

## Compiler changes

### Reuse an existing imported-dictionary construction

Inspection of optimized IR found that `decodeEvent` rebuilt
`decodeForeignObject decodeJsonJson` on every call, even though the same
application already had a typed top-level getter. PBO had reintroduced the
application without its dictionary result annotation, so the existing
annotation-based lifting pass missed it.

`Gopurs.ClosedDictionaries` now recognizes exact applications of imported
globals corresponding to an existing typed dictionary binding. It independently
recovers the result type from the declared global types and monomorphic
arguments, checking the instantiated parameter types and binding annotation.
Missing/dynamic types, unresolved polymorphism, local-module dependencies,
locals, explicit type applications and annotated child expressions decline
this reuse. Qualified names and curried/uncurried application shapes are kept
distinct. Recursive groups and local recursive initialization scopes are
preserved. Top-level roots remain excluded from lifting through their annotation
wrappers; newly allocated getter names still avoid source-name collisions.

This is a small extension of dictionary sharing, with no application schema,
Argonaut symbol names or constructor-layout constants in the pass.

### Bound duplicated callback syntax

`Gopurs.ImmediateApplications` previously rejected any callback over 128 syntax
nodes, including a callback inserted only once in a successful `Either` branch.
It now bounds **extra copies** across all syntactic branch leaves. Zero/one
substitution permits a large callback; duplicated syntax retains a size budget.
Each lambda parameter still has at most one use, arguments must be movable
values, and recursive local scopes are excluded. Captured variables are
preserved through freshening; discarded callback bodies stay unevaluated.

In emitted `decodeEvent`, static `gopurs_runtime.Func` construction sites go
from **7 to 3**, and `Apply` sites from **5 to 3**. These are source-site counts,
not a complete runtime-call count or a CPU-time decomposition. Some immediate
callbacks remain because their parameters have multiple uses.

### Forwarder regression coverage

Focused tests also found that mechanical foreign-forwarder recognition accepted
any function named `unsafeCoerce`, including unrelated qualified or shadowed
names. Recognition now requires the qualified `Unsafe.Coerce.unsafeCoerce`
intrinsic, rejects shadowed unqualified foreign heads and duplicate parameter
names, and tests exact argument order, captures and recursive calls. This
tightens which wrappers remain polymorphic. It causes no additional emitted-Go
change in the JSON diagnostic.

## Path coverage and ownership

Separate untimed instrumented builds validate all seventeen setup outputs and
the five retained counted outputs. After initialization:

| Calls per five-case corpus | previous final | new final |
|---|---:|---:|
| object-decoder construction | **2,477** | **0** |
| identity-object decoding | 2,477 | 2,477 |
| custom container-element decoding | 2,477 | 2,477 |
| nested records recognized by a parent plan | 6,350 | 6,350 |
| compact records constructed | 6,355 | 6,355 |

The independent output object is still allocated on every identity-object
decode. Output records, arrays, ownership and callback behavior retain the
library implementation validated previously. The gain combines dictionary
reuse and callback simplification; there is no per-transformation timing
ablation.

## Paired retained-output campaign

`paired-final/results.json` compares `final-audit` with the preserved
`json-dec-plan-audit-20260924/identity-object-audit`. Both use the corrected
runner: two warm-ups, five samples per mode/phase, six processes per workspace,
alternating decoder/workspace order and rotating/reversing phase order.
`GOMAXPROCS=1`, `GOGC=100`; sources, binaries, manifests, corpus and oracle are
hashed. The retention buffer is created before timing; produced values remain
live through validation, which checks those exact values after timing and
allocation sampling.

| Phase / allocation | previous generated | new generated | specialised in new workspace |
|---|---:|---:|---:|
| parse | 2.709 ms | 3.047 ms | 2.770 ms |
| decode | **3.906 ms** | **2.926 ms** | **0.821 ms** |
| combined | **8.321 ms** | **7.818 ms** | **4.235 ms** |
| parse allocation, minimum pass | 4.613 MB | 4.613 MB | 4.613 MB |
| decode allocation, minimum pass | **6.273 MB** | **5.005 MB** | **1.844 MB** |
| combined allocation, minimum pass | 10.886 MB | 9.617 MB | 6.457 MB |

Timing cells are medians of process minima. Ratios of medians give **−25.1%
decode**, **−6.0% combined**, **−20.2% decode allocation** and **−11.7% combined
allocation**. The same-campaign generated/specialised decode ratios are about
**4.47× before / 3.56× after**; the previous campaign's approximately 5× ratio
should not be used as a paired baseline.

All **1,800 sampled outputs** are retained and fingerprinted, in addition to
setup and warm-up outputs. Every frozen fingerprint agrees, including the ten
invalid-schema results and their error strings.

### Timing variability

One-minute load averages range **21.4–25.4**. Decode improves in five of six
pairs (deltas −57.4% to +10.7%); combined improves in four of six (−21.9% to
+55.9%). The median within-pair changes are −25.6% and −4.6%, respectively.
The parse control is **12.5% slower** by ratio of medians, with unchanged
allocation; two parse pairs differ by 80–91%. The specialised decode control
is 0.874 ms in the baseline workspace versus 0.821 ms in the new one.
These observations support the allocation reduction and decode improvement
more strongly than a precise end-to-end effect size.

The earlier two-process `pilot` also validates outputs and yields the same
minimum allocations, but its timings are more variable. Its figures are not
pooled into the final campaign or used to choose a favourable process subset.

## Official diagnostic campaigns

`final-official/results.json` uses the official three-process schedule, with
the new native compiler and local JavaScript codecs:

| JSON backend | parse | decode | combined |
|---|---:|---:|---:|
| Go | 2.434 ms | **2.777 ms** | **6.904 ms** |
| JavaScript | 1.722 ms | 7.462 ms | 8.624 ms |
| C++ reference | 0.373 ms | 0.258 ms | 0.630 ms |

All existing frozen-oracle checks pass. The C++ row retains the previously
documented reference contract and lifecycle; this follow-up does not establish
retained-output lifecycle alignment with C++. The official **6.904 ms** combined
Go result updates the README diagnostic cell. It has a different schedule and
allocator history from the paired audit's **7.818 ms** and cannot be substituted
into the paired improvement calculation.

`tast-final-official/results.json` validates all twelve TAST module fingerprints
for Go, JS and C. Go parse/decode/combined are **25.965 / 16.294 / 55.064 ms**.
Combined process minima are 65.558, 55.064 and 47.465 ms. Decode allocation is
27,071,160 bytes, versus 27,075,000 in the previous official campaign; parse
allocation is unchanged at 56,031,272 bytes. This is a functional recheck and
fresh measurement, not a paired claim of TAST improvement or neutrality.
The README TAST diagnostic cell is updated to **55.06 ms**. The three excluded
diagnostic rows remain outside benchmark totals and `/C` ratios.

## Validation and provenance

- **35 focused compiler/runtime tests pass**, covering dictionary qualification,
  captures, shadowing, effects, recursive scopes and collisions; forwarders;
  large/unused/duplicated callbacks; generated-Go evaluation order and failure;
  boxed records, imported workers, recursive initialization and closure lifetime.
  Existing closure/function-data checks include the Go race detector.
- **Four compiled fixtures pass:** `JsonRecordPlan`, `StaticDictionary`,
  `TypeClassesInOrder`, and `3558-UpToDateDictsForHigherOrderFns`.
  Only the `JsonRecordPlan` snapshot changes, from fresh-local renumbering.
- **`node test/typed-plans.mjs` passes** in the codec library, including local
  JavaScript and newly generated Go with `-race`: nested/optional/custom
  decoders, errors, constructed/raw JSON, numeric boundaries, retained results
  and independent object ownership.
- **Native bootstrap passes:** 459 TAST modules and **275,764 type-table entries**.
  Node-generated and rebuilt-native-generated JSON workspaces contain
  **326/326 byte-identical Go files**, including after the forwarder correction.
  Five Go files differ from the previous final generation.
- The audit's exact-timed-output validation regression passes in each audit
  build. Coverage and allocation-profile binaries are kept separate from
  timing builds.

`compiler-final` archives the changed compiler sources, focused tests, Node
bundle and rebuilt native binary. Native SHA-256:
`ee29d6732a9fba053079f796bb6202c5f2e2c3af641fb21929d5829c26da9e99`.
`before` retains the previous compiler artifacts and changed source modules.
`generation-verification.json`, `library-verification.json`, `paired-analysis.json`
and the build manifests record the comparisons; bootstrap logs are retained.

## Remaining costs

`allocation-profile/alloc-space.txt` samples 100 decode passes plus setup and
a validating pass. The owned identity-object copy is still the largest single
flat allocation site. Compact records, arrays, `Either`/`Maybe` constructors,
accessor conversions and remaining callbacks also remain visible. The repeated
object-decoder construction has disappeared from the large allocation sites.
These are allocation samples, not CPU shares; an event decoder's cumulative
share includes nested item decoding.

A next focused experiment can address multi-use immediate callbacks or fuse
temporary-object accessors with explicit ownership rules. The specialised
decoder remains schema-specific and is an implementation reference rather than
a demonstrated compiler target or language limit. Separate phase medians do
not establish an additive floor.

## Reproduction

Use new destination paths:

```sh
cd /path/to/gopurs
npm run build:native -- --keep-workspace
node --test tools/closed-dictionaries.test.mjs tools/foreign-forwarders.test.mjs tools/immediate-applications.test.mjs

cd /path/to/altbak.pub-gopurs
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --workspace NEW_BUILD
python3 bin/benchmark/json-diagnostic/specialized/run.py setup --source NEW_BUILD --workspace NEW_AUDIT
python3 bin/benchmark/json-diagnostic/specialized/run.py campaign --workspace NEW_AUDIT --compare-workspace PRESERVED_BASELINE_AUDIT --processes 6 --output NEW_PAIRED_RESULTS
python3 bin/benchmark/json-diagnostic/specialized/coverage.py --source NEW_BUILD --workspace NEW_COVERAGE
```
