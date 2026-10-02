# Purust: native codegen optimization campaign

## Scope and retained changes

This second campaign targets **`purust-aff`**, on the same frozen application
inputs as the [first campaign](2026-10-02-purust-compiler-optimization.md):
244 TAST modules, 138,448 type-table entries, 27,122,224 TAST bytes,
`--main Test.Main --threaded`.

Retained changes:

1. Specialize `sanitizeIdent` in Rust, reusing the owned buffer for ordinary
   ASCII identifiers and writing escaped identifiers into one output buffer.
   The original PureScript implementation remains the JavaScript fallback.
2. Remove the unused `globalConsumed` and `globalCaptured` references. Capture
   information already flows through local function arguments; the save/union/
   restore operations did not feed codegen decisions.
3. Generate modules concurrently through a bounded FIFO of fibers, returning
   immutable results and publishing them in input order. Generator construction
   is deferred until its worker executes. One supervisor covers optimization,
   generation and the final drain, including cancellation and joining children.
4. Reserve generation slots from the existing worker budget. The native default
   is half the budget, up to 4 generation slots: **4 PBO + 4 codegen workers** on
   this host. JavaScript and TAST loading remain sequential by default.

The self-rebuilt stage-2 compiler is installed and passes the full Aff suite.
The [final five-pair benchmark](2026-10-02-purust-aff-codegen-compilation.md)
measures **6,463 ms JavaScript / 8,247 ms native**, ratio **1.28×**. All 12 outputs
match exactly and the generated application passes 47 Aff checks. Both backends
slow down in the last two pairs; every sample is retained in the published data.

The [raw campaign archive](2026-10-02-purust-codegen-optimization.json) includes
the comparisons below, compiler/source hashes, regression and self-hosting logs.

## Protocol and provenance

The restored reference has **548 files**, checked against the previously
published manifest. TAST SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

At campaign start, the installed native executable differed from the published
one. Both were preserved and compared explicitly; the installed artifact's
origin is not inferred from its filename. The JavaScript bundle was unchanged.

| Initial native artifact | SHA-256 |
|---|---|
| Published first-campaign stage 2 | `78eb3406774d118c28d8e709713514d6ccb9148d5dfcf83cf29631fe945e4b37` |
| Installed at second-campaign start | `5f554b9e13734625f7313ab389f57c6b7e8dc55bae979ca73539030eccfb1d42` |

Host: Apple M4 Pro, 14 cores, 48 GiB RAM, arm64 macOS 26.1, Node v24.8.0,
Rust/Cargo 1.96.0. Native compiler profile: **O3, LTO disabled**.

Exploratory comparisons use `tools/compare-native.mjs`: one warm-up per variant,
three measured rounds, rotating the first variant; fresh process, output and
`.purmeta` state per invocation. Compiler builds and tests finish before timing.
Normal OS file caching, no CPU affinity. All samples are retained. The metric
is monotonic **`backend total`**, including loading/sorting, preparation, PBO,
Rust generation and emission. It excludes startup/exit, the `purs` frontend,
bootstrap, Cargo and application execution. `/usr/bin/time -l` records RSS.

Each comparison uses its own contemporaneous control. Times from different
campaigns are not combined to calculate speedups. Phase medians are independent
and need not sum to the total median.

Durable workspace: `var/benchmark/purust-codegen-20261002/`.

## Native identifier sanitizer

| Variant | Three samples (ms) | Median total (ms) |
|---|---|---:|
| Published reference | 9,644 / 9,642 / 9,656 | **9,644** |
| Installed reference | 9,599 / 9,634 / 9,647 | **9,634** |
| Native sanitizer | 8,514 / 8,581 / 8,620 | **8,581** |

The sanitizer reduces the median by **11.0%** against the published reference.
All **12 invocations** produce 496 identical Rust source files/Cargo manifests.
Median optimize + generate falls from 5,611 to 4,686 ms; finalize + emit from
1,549 to 1,400 ms. Raw campaign: `ident.json`.

The native implementation passes **133,148 differential cases** against the
compiled PureScript reference: every UTF-16 code unit, isolated surrogates,
supplementary characters, keywords, punctuation, deterministic mixed strings,
long inputs and reuse of the original ASCII buffer. Purust strings encode one
Rust scalar per UTF-16 unit; escaping uses `purust_char_to_code_unit`, preserving
the original CodeUnits semantics.

## Final controlled before/after comparison

After the final JS/native campaign and its application validation, the published
first-lot native compiler and installed second-lot stage 2 were compared directly
on the restored input snapshot. One warm-up per artifact, then **five measured
rounds**, alternating which compiler runs first. All 12 outputs have the same
496 files. Raw campaign: `before-after.json` in the archive.

| Native compiler | Five samples (ms) | Median total (ms) | Maximum RSS (MiB) |
|---|---|---:|---:|
| Published first-lot stage 2 | 10,553 / 10,557 / 10,395 / 10,419 / 10,321 | **10,419** | 465 |
| Second-lot stage 2 | 8,707 / 8,160 / 8,249 / 8,189 / 8,007 | **8,189** | 498 |

This lot reduces time by **21.4%**, a **1.27× speedup**, with higher maximum RSS.
These measurements isolate the second lot against its predecessor. They do not
combine sanitizer and parallel speedups from separate exploratory campaigns.
The README uses the final JS/native campaign's **8,247 ms** native median.

## Stateless and parallel generation

PBO's total budget is fixed at 8. Two generation slots reserve 2 slots from PBO,
and four reserve 4. One-slot generation runs on the coordinator with 8 PBO workers.

| Variant | Three samples (ms) | Median total (ms) | Median optimize + generate (ms) | Maximum RSS (MiB) |
|---|---|---:|---:|---:|
| Sanitizer-only control | 9,114 / 9,073 / 9,154 | **9,114** | 4,939 | 451 |
| Stateless, sequential generation | 8,985 / 8,979 / 9,346 | **8,985** | 4,886 | 455 |
| 6 PBO + 2 generation | 8,114 / 7,971 / 8,280 | **8,114** | 4,032 | 500 |
| 4 PBO + 4 generation | 7,707 / 8,078 / 7,741 | **7,741** | 3,846 | 502 |

All **16 invocations** produce 496 identical files. Four generation workers
reduce total time by **13.8%** against the new sequential path, and **15.1%**
against the sanitizer-only control, with higher peak memory. The sequential
reorganization's 1.4% difference is small relative to sample variability and
is not treated as a strong standalone performance result. Raw: `emission.json`.

Measured deferred PBO attempts are 117–118 with sequential generation, 92–96
with two generation slots, and 65–69 with four. Fewer optimizer slots also reduce
speculative retry work. `dispatched + fallback` counts attempts, including
retries. The parallel `codegen-ms` diagnostic measures coordinator time spent
enqueuing/waiting/publishing; it is **not the sum of worker CPU time** and does
not include the final drain. Total/phase timers do include the drain.

Configuration:

- `PURUST_PBO_JOBS=1..64`: worker budget; default at most 8, capped by CPU count.
- `PURUST_CODEGEN_JOBS=1..64`: requested generation concurrency, capped to leave
  at least one optimizer slot. At 1, generation is inline and PBO uses the full
  budget. Native default: half the budget, at most 4; JavaScript default: 1.
- `PURUST_PBO_JOBS=1` selects fully sequential operation.
- `PURUST_JOBS`: independent TAST loading concurrency, default 1.

## Regression checks

- Full codegen suite: **90 tests passed**, followed by **9 targeted tests** for
  the final emitter/scheduler/default-budget changes (including overlapping
  tests from the full suite).
- **43 TAST tests passed**. The crypto test initially failed because its Docker
  container was stopped; it passed after restart. Both logs are retained.
- Native emitter fixture: **12 scenarios** at 1/2/4 workers, **474 identical
  JS/native generated files**, executable result `EMISSION_NATIVE_OK 12 scenarios`.
  It checks the concurrency bound, ordered publication, final drain and child
  cleanup after generation, publication or producer failure.
- JS additionally checks synchronous failure while constructing a generator,
  observed at its ordered publication point, plus a shared PBO/codegen lifetime.
- Budget tests cover small CPU budgets, explicit overrides, invalid values and
  combined concurrency limits.

The first emitter bootstrap failed with `No space left on device`. Cargo caches
from completed earlier experiments were cleaned after preserving executables;
generated sources, measurements and logs remain available. The corrected
bootstrap succeeded in a new workspace. This failed build contributes no timing
samples to the comparisons above.

## Self-reconstruction and full Aff qualification

The compiler's **452 TAST modules and 282,280 type-table entries** were generated
with the typed frontend. JavaScript emitted stage 1, compiled at O3 without LTO.
Stage 1, at the new default **4 PBO + 4 codegen workers**, regenerated all **912
Rust sources/Cargo manifests byte for byte**. Cargo then built stage 2.

Stage 2 passed the independent fresh-project smoke test: **152 modules, 312
identical files and `PURUST_NATIVE_OK 42`**. The first smoke run's assertions all
passed, but directory cleanup returned `ENOTEMPTY`. The bootstrap now forwards
`--keep-workspace` to its smoke test, and temporary cleanup allows retries.
The completed stage 2 was smoke-tested again with its workspace retained, then
atomically installed. The original failure and successful recovery logs are
retained; neither compiler stage needed to be rebuilt for this cleanup issue.

The installed compiler passes `purust-aff/bin/test`: **47 Aff checks, 5 Rust unit
tests**, concurrent timer/Ref/AVar integration, child lifetime checks, and **9
error-reporting scenarios**. Native execution remains the default; `PURUST_JS=1`
selects the JavaScript compiler explicitly.

Installed native SHA-256:
`0a0f2afe52159a3c272beeaee625e68d100d5855bd3d99dc18ab6f353eca875e`.

## Archive verification

A separate archive pass rechecked all compiler/frozen-input hashes and total
medians, the **544 identical application inputs**, the **912 self-hosted files**,
and all **52 application outputs** across the sanitizer, emission, before/after
and final JS/native campaigns. Every output has the same **496 source/manifests**.
The installed executable's hash matches the native artifact measured in the
final campaign. Results are recorded in the archive's `verification` section.
