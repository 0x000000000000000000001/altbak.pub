# Shared JSON record construction: integrated results

The standard Argonaut record decoder now constructs each output record once,
using one shared loop and an immutable field plan. On the general JSON corpus,
decoding is **15.09% faster**, with **16.02% fewer allocated bytes**. Independently
measured parse + decode improves **8.95%**. **No TAST improvement is established.**

These are generated applications built with the final rebuilt native compiler,
not the manually patched Go probes. The public `GDecodeJson` API is unchanged.

## Paired measurements

Times are milliseconds per corpus. The combined phase is measured separately;
its minimum is not the sum of the separate parse and decode minima.

| Corpus and phase | Go control | Go final | JS control |
|---|---:|---:|---:|
| General JSON, parse | 2.602 | 2.560 | 1.681 |
| General JSON, decode | 13.385 | **11.365** | 7.822 |
| General JSON, combined | 16.599 | **15.113** | 9.284 |
| TAST, parse | 26.854 | 26.551 | 20.727 |
| TAST, decode | 377.262 | 368.694 | 57.535 |
| TAST, combined | 428.852 | 424.097 | 80.259 |

General JSON decode allocations: **19,714,272 → 16,556,992 bytes**. Combined:
**24,327,056 → 21,169,776 bytes (−12.98%)**. Parsing allocations are unchanged.
TAST decode and combined allocations remain exactly **442,632,472 / 498,663,744
bytes**. Its timing ranges overlap substantially and the paired differences
have mixed signs; the small median differences are not an established gain.

The README now records **15.11 ms Go / 9.28 ms JS** for general JSON and
**424.10 / 80.26 ms** for TAST. Previously it recorded **17.00 / 9.30** and
**422.27 / 76.80**, from another session. Use the paired controls above to
attribute the change, rather than subtracting those historical cells. General
JSON remains about **1.63×** the JS combined time.

## Protocol and variability

- Go 1.27.0, darwin/arm64; GOMAXPROCS=1, GOGC=100, PGO off. Fresh builds and
  all tests completed before timings; no concurrent task builds or profiling.
- General JSON: **two series of five alternating process pairs**, aggregated
  across all ten control and ten treatment processes. Six identical-bundle JS
  controls. The second series confirmed the first rather than replacing it.
- TAST: five alternating pairs and three identical-bundle JS controls.
- Two warmups, five samples per phase per process; published cells are medians
  of process minima. Allocation cells are medians across all allocation samples.
- Every process checks all **17 JSON cases or 12 TAST modules**, including full
  parsed and decoded fingerprints. Five successful JSON cases are timed.
- The first final JSON series contained a treatment combined minimum of
  **32.785 ms**. It is retained in the archive and in the aggregate. Both series
  improve their median decode and combined times; nine of ten combined pairs
  favor the treatment. No cause is assigned to the observed timing variability.

Application binary sizes: JSON **6,420,994 → 6,456,386 bytes (+0.55%)**;
TAST **9,058,066 → 9,059,202 bytes**. The native compiler was rebuilt and
installed. No complete b8x compilation was measured.

## What changed

The local `gopurs-argonaut-codecs` Nil/Cons instances compose an internal field
plan. Its loop calls the real PureScript field decoder, checks its `Either`,
and stores the successful payload in the final buffer. A successful field
reuses its existing `Right` envelope; this is verified in generated Go. Only
the completed record is materialized, eliminating successive `Record.insert`
copies, their repeated symbol-dictionary coercions and recursive row results.

ADT checks and error decoration remain in PureScript. The FFI receives ordinary
callbacks and does not hardcode generated constructor names, tags or offsets.
Unknown tail dictionaries retain the original implementation. Symbol callback
order, first error, immutable results and extra tail fields are preserved.

The runtime adds `WithFunctionData` / `FunctionData[T]` with a distinct callable
tag and GC-visible ownership, without enlarging `Value` or adding a global
registry. All application arities support it. Foreign function classification
and Promise callbacks were adapted and tested as well.

The earlier generated-Go probe gained about 19–23% in decode time. The first
integrated CPS interface gained less and allocated a private continuation state
per record. The retained interface removes that state and reuses the field's
existing `Either`. Both integrated variants and the probes are archived;
their measurements are not mixed into the final campaign.

## Validation and remaining scope

- Compiler suite: 174 passed, two skipped, zero failures. The fixture runner's
  eight tests also pass after fixing selection of explicitly requested native
  libraries outside its bootstrap core set.
- Runtime: 121 arity/application combinations, partial and excess arguments,
  Uncurried calls, metadata typing/replacement, GC, concurrency and zero
  allocations on saturated Apply1..10; optimized and race builds pass.
- Record helper: six Go tests with race detection and JS fallback assertions,
  covering sizes 0/1/2/5/6/32, order, errors, name collisions, the second symbol
  evaluation, custom fallback, aliasing, GC, reentrancy and concurrent calls.
- `JsonRecordPlan.purs` executes successfully in native Go and JS. It exercises
  custom `DecodeJson` and `GDecodeJson` instances, missing/null fields, nested
  errors, callback order and a retained shared tail record.
- Foreign and Promise compatibility tests pass with race detection.

**Payloads and output elements remain `Value`.** This closes the repeated
record-construction step; it does not complete a native typed record/array ABI.
The TAST's own construction path still needs a renewed profile and a separate
optimization. These results establish no final performance ceiling.

[Raw campaigns, provenance, intermediate variants and reproduction script](2026-09-22-json-record-plans.json)
are preserved. The Argonaut package is an existing local overlay without its
own Git repository, so its complete changes are also preserved as an
[applicable patch](2026-09-22-json-record-plans-argonaut.patch), plus complete
changed source files in the JSON archive. Apply the patch from the original
`gopurs-argonaut-codecs` root with `patch -p1`.
