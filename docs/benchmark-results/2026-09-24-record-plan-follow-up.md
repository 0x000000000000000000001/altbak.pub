# Record plans: contract fixes, measured coverage and retained-output comparisons

This follow-up corrects the earlier typed-plan validation and specialised
runner, then implements library-side optimizations identified in the actual
generated path. The application schema, frozen corpus and fingerprints are
unchanged. No general compiler specialization pass is introduced.

The final paired campaign measures **5.717 → 3.892 ms decode** and
**9.839 → 7.872 ms combined**, with **8.860 → 6.273 MB** decode allocations,
against the corrected baseline. The aligned specialised reference is about
**0.76–0.77 ms decode** on this campaign; a substantial gap remains.

All artifacts below are under
`var/benchmark/json-dec-plan-audit-20260924/`. Separate workspaces and result
directories preserve every measured stage.

## Correctness and measurement corrections

- **Symbol callbacks:** removed `fieldKeyCache`/`sync.Once`. A typed decoder
  does not prove its symbol callback constant. Lookup evaluates each symbol
  in field order; successful insertion reevaluates symbols in reverse order.
  Failure stops at the first field, with no insertion callbacks. Duplicate
  labels retain the original `Record.insert` precedence.
- **JavaScript FFI:** aligned `recordNilImpl` and `recordConsImpl` with the
  current PureScript signatures (3 and 11 construction arguments). Updated
  helper calls and added compiled PureScript integration checks.
- **Actual JS dependency:** the earlier official JS benchmark used registry
  codecs, so it did not exercise the modified `Record.js`. The official builder
  now explicitly selects local `gopurs-argonaut-codecs` for JS too. Go, JS and
  C reports pass their existing frozen-oracle checks.
- **Specialised runner:** every timed output is retained through validation.
  The retention buffer is allocated before timing, matching the official
  runner. Fingerprinting happens after timing and allocation sampling and
  checks the exact produced outputs, without rerunning decoding. Parsed-value
  setup hashes now encode the actual parsed values rather than hash input text.
- **Runner regression:** a deliberately bad first timed output followed by a
  valid last output must fail validation in parse/decode/combined modes. This
  catches the former last-result-sink/setup-only protocol.
- **Provenance and scheduling:** generated Go sources, source manifests,
  audit sources, binaries, corpus and oracle are hashed. Campaigns reject
  changed artifacts, sanitize inherited `DIAG_*`/profiling overrides, alternate
  starting decoder and workspace order, rotate/reverse phase order, and record
  UTC timestamps and load averages. `GOMAXPROCS=1`, `GOGC=100`.

These checks establish agreement for the exercised cases, not universal
Argonaut equivalence. C++ retains its documented narrower malformed-input
contract; no C++ lifetime alignment or usable simdjson-to-Go transfer result is
claimed here.

## Implemented optimizations

1. **Propagate record metadata through `decodeRecord`.** The ordinary JSON
   wrapper is tagged with its recognized row-decoding plan. Parents can then
   decode nested records directly, including records inside arrays and `Maybe`.
   Unknown `GDecodeJson` methods retain their ordinary implementation.
2. **Borrow array input views during decoding.** Raw `[]any` and boxed arrays
   are read directly; the plan no longer allocates and recursively boxes a
   separate native input slice before constructing the decoded output array.
   Output storage remains independent of the input and other calls.
3. **Compact small record construction.** Up to five fields use call-local
   scratch and a final `RecordDict1`–`RecordDict5`. Symbols are still reevaluated
   and duplicates handled; larger records retain the generic construction.
4. **Standard identity-object decoding.** An allocation profile of the first
   optimized stage identified the `Foreign.Object Json` traversal inside the
   custom event decoder as a large allocation site. `typedJson` identifies the
   standard identity decoder; `typedObject` selects a narrow shortcut only for
   that case. It allocates a fresh output map and boxes its entries, preserving
   container ownership. Arbitrary/custom element decoders retain traversal.

The object shortcut uses no schema or constructor-layout constants. Scalar and
container support constructors continue to come from the caller. In nested
records containing custom fields, `Box(raw)` preserves an already boxed object
instead of introducing an extra `Any(Value)` layer.

### Actual path coverage

`coverage.py` instruments a separate copy of emitted record FFI. These binaries
are never used for timings. All 17 setup outputs and the five retained counted
outputs match the oracle.

| Calls per five-case corpus | corrected baseline | final |
|---|---:|---:|
| records decoded by their internal plan | 6,355 | 6,355 |
| nested `Record` kinds recognized by a parent | **0** | **6,350** |
| untagged record-field steps | 2,700 | **0** |
| custom container-element calls | 7,627 | **2,477** |
| compact three-field records | 0 | 3,650 |
| compact four-field records | 0 | 5 |
| compact five-field records | 0 | 2,700 |
| identity-object shortcuts | 0 | 2,477 |

The remaining 2,477 custom-element calls correspond to the event decoder. This
is a call count, not a measured fraction of CPU time. Accessors still perform
`FO.lookup` and surrounding `Either` handling.

## Measurements

### Final paired campaign

`paired-final/results.json` compares `identity-object-audit` with the preserved
`corrected-audit-aligned` baseline. The baseline already includes callback and
JS interface fixes. Six processes per workspace cover all phase permutations,
with two warm-ups and five samples per decoder/phase.

| Phase | baseline generated | final generated | specialised in final workspace |
|---|---:|---:|---:|
| parse | 2.699 ms | 2.751 ms | 2.675 ms |
| decode | **5.717 ms** | **3.892 ms** | **0.772 ms** |
| combined | **9.839 ms** | **7.872 ms** | **3.446 ms** |
| decode allocation, minimum pass | **8.860 MB** | **6.273 MB** | **1.844 MB** |
| combined allocation, minimum pass | 13.473 MB | 10.886 MB | 6.457 MB |

Ratios of medians: **−31.9% decode**, **−20.0% combined**, **−29.2% decode
allocation**. Median within-pair percentage changes are −30.0% and −19.3%
for decode and combined. All six pairs improve; decode improvements range
15.4–43.0%, combined 15.4–26.8%. One-minute load averages range 12.3–14.5.
The parse control is 1.9% higher by ratio of medians, with unchanged allocations.
Background variation remains; these results evaluate the whole optimization
bundle, not each transformation's independent time share.

All **1,800 sampled outputs** (12 processes × 3 phases × 2 modes × 5 samples ×
5 cases) are retained and fingerprinted after timing, in addition to warm-up
and setup validation. The specialised decoder in the baseline workspace
measures 0.757 ms decode / 3.551 ms combined. The generated/specialised decode
ratio in the final workspace is about **5.0×**.

### Preserved first optimization stage

`paired/results.json` compares `optimized-audit` with
`corrected-audit-aligned`: six processes per workspace, two warm-up passes and
five samples per mode/phase, median of process minima. This stage includes
record propagation, array views and compact records, before the identity-object
shortcut.

| Generated decoder | corrected baseline | first optimized stage |
|---|---:|---:|
| parse | 2.719 ms | 2.760 ms |
| decode | 5.573 ms | 4.565 ms |
| combined | 10.683 ms | 8.444 ms |
| decode allocation, minimum pass | 8.860 MB | 6.758 MB |
| combined allocation, minimum pass | 13.473 MB | 11.371 MB |

Ratios of timing medians give −18.1% decode and −21.0% combined. The median of
the six within-pair percentage changes is −29.6% and −17.2%, respectively;
these are different statistics. Every pair improves, but decode improvements
span 8.5–54.5% and one parse control is 139% slower. Recorded one-minute load
averages range roughly 27–32. This supports the direction and allocation
reduction more strongly than a precise timing effect size.

The aligned specialised reference measures 0.765–0.789 ms decode in these two
workspaces, with 1.844 MB allocated. The old 0.557 ms single-sink result is not
an aligned floor. Even the corrected reference remains specific to one schema
and hard-coded gopurs constructor layouts.

### Final official campaign

`final-official/results.json` uses the official three-process Go/JS/C schedule
and retained-output protocol with the identity-object shortcut included:

| Backend | parse | decode | combined |
|---|---:|---:|---:|
| Go | 2.634 ms | **3.825 ms** | **8.273 ms** |
| JS, local codecs | 1.743 ms | 7.233 ms | 8.938 ms |
| C++ reference | 0.391 ms | 0.269 ms | 0.658 ms |

These official cells and the paired audit cells use different schedules and
allocator histories. They should not be mixed to derive per-transformation
time shares or an additive parse/decode floor. In particular, the separate
stage-one official combined cell was 8.150 ms; its difference from 8.273 ms
does not isolate the identity-object shortcut's time effect.

### TAST recheck

`tast-final-official/results.json` passes all twelve module fingerprints for
Go, JS and C. Go parse/decode/combined medians are **27.683 / 15.912 / 49.903 ms**.
Go combined process minima are 65.265, 49.903 and 49.654 ms. This is a fresh
functional/performance check, not a paired demonstration of neutrality or gain
against the historical 50.511 ms cell.

## Validation

- `node test/record-plan.mjs`: Go helper tests with `-race`, plus JS FFI helper
  calls. Coverage includes changing typed symbols, early failures, duplicate
  insertion labels, raw/boxed/custom nested records, arrays of arrays, retained
  results, empty objects/non-objects, and fresh identity-object storage.
- `node test/typed-plans.mjs`: the same PureScript application runs through the
  modified local JS FFI and generated Go with `-race`. It exercises standard
  accessors, missing versus null optional fields, nested records/containers,
  custom ADTs and object decoders, error precedence, raw and constructed JSON,
  and tagged-versus-ordinary integer decoding at numeric boundaries.
  The numeric test preserves each backend's existing policy; it does not claim
  JS/Go equality outside the portable integer range.
- `./bin/test JsonRecordPlan --update-snapshots`: existing custom-dictionary,
  callback-order and shared-tail fixture passes; snapshot updated.
- Specialised timed-output validation regression passes in every audit build.
- Official frozen-oracle JSON and TAST campaigns pass, including the modified
  local JS path.
- `npm run build:native -- --keep-workspace` succeeds, including its Node
  backend build and native bootstrap with the local libraries: 459 TAST modules,
  275,340 type-table entries. Running the rebuilt native compiler on the same
  final diagnostic produces **326/326 byte-identical Go source files**.
  `post-bootstrap/verification.json` records the old/new compiler hashes and
  comparison. Measurements use the preserved pre-bootstrap generator
  (`e31ed31e2ffb…`); the rebuilt compiler is `7bd90b6fabe2…`.

## Provenance and reproduction

Preserved directories include:

- `library-before`, `library-corrected`, `library-final` (the native library is
  not itself a Git checkout), final source hashes, and the generator binary
  `compiler-before-bootstrap`;
- `corrected`, `optimized`, `identity-object`: official generated workspaces,
  source manifests and binaries;
- `corrected-audit-aligned`, `optimized-audit`, `identity-object-audit`: complete
  generated Go sources, runner tests, audit build manifests and binaries;
- `corrected-coverage`, `optimized-coverage`, `identity-object-coverage`: exact
  instrumented sources, manifests and counts;
- `corrected-official`, `optimized-official`, `final-official`, `paired`, and
  `paired-final`, `tast-final-official`: raw process outputs and aggregates;
- `allocation-profile`, `allocation-profile-final`: allocation profiles over 100 decode passes,
  including setup and one validation pass. Its sampled allocation attribution
  is not CPU attribution.

The final allocation profile still identifies the owned temporary JSON-object
copy, compact final records, arrays, `Either`/`Maybe` constructors and dynamic
applications/closures. The custom event-decoder call tree includes nested item
decoding, so its cumulative allocation share cannot be read as the cost of its
own dispatch alone. A next compiler experiment can target repeated closed
decoder construction or fusion of temporary object/accessor operations, with
explicit ownership and callback proofs, rather than assuming that a general
specialization pass will recover the whole residual gap.

Example for a fresh final build and audit (all destination paths must be new):

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --workspace NEW_BUILD
python3 bin/benchmark/json-diagnostic.py measure --suite JsonDecoding --workspace NEW_BUILD --output NEW_OFFICIAL_RESULTS
python3 bin/benchmark/json-diagnostic/specialized/run.py setup --source NEW_BUILD --workspace NEW_AUDIT
python3 bin/benchmark/json-diagnostic/specialized/run.py campaign --workspace NEW_AUDIT --compare-workspace PRESERVED_BASELINE_AUDIT --processes 6 --output NEW_PAIRED_RESULTS
```

These experiments narrow concrete library costs. The residual includes custom
decoder bodies, dynamic application/constructor handling, field lookups and
required output allocation. A compiler transformation should be selected using
remaining profiles and preserve custom-dictionary behavior; no general pass or
guaranteed sub-millisecond target follows from the prototype alone.
