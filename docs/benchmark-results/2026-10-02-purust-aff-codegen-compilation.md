# Purust Aff compilation after the codegen optimization lot

## Published result

Measured on **2026-10-02**, using the self-rebuilt, installed stage-2 compiler:

| Backend | Median `backend total` | Range across five runs |
|---|---:|---:|
| `purust.js`, sequential PBO/codegen | **6,463 ms** | 6,400–7,995 ms |
| `purust-native`, 4 PBO + 4 codegen workers | **8,247 ms** | 7,708–9,769 ms |

The native compiler takes **1.28× the JavaScript time** on this workload.
The [first optimization campaign](2026-10-02-purust-aff-compilation.md) published
1.52×. Native has improved but has not yet beaten JavaScript.

A separate controlled comparison against the first campaign's published native
binary gives **10,419 → 8,189 ms**, a **21.4% reduction / 1.27× speedup** for this
second lot. See the [optimization report](2026-10-02-purust-codegen-optimization.md#final-controlled-beforeafter-comparison).
These are separate samples; the README uses **8,247 ms** from the paired campaign.

| Pair, in execution order | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Warm-up, excluded | 6,383 | 7,689 |
| 1 | 6,463 | 7,749 |
| 2 | 6,400 | 7,708 |
| 3 | 6,421 | 8,247 |
| 4 | 7,457 | 9,297 |
| 5 | 7,995 | 9,769 |

Both backends slow down in the last two pairs. The recorded one-minute system
load also rises, from roughly 7 to 11; this does not establish the cause of the
slowdown. **All five samples are retained**, with no selection of faster runs.
The published ratio is the ratio of the two backend medians, following the
existing protocol. The ranges and individual pairs expose the observed drift.

The [raw measurements](2026-10-02-purust-aff-codegen-compilation.json) contain
every sample, phase, stdout/stderr log, compiler/input hash, checkout state and
application-validation result. The [codegen campaign archive](2026-10-02-purust-codegen-optimization.json)
includes exploratory comparisons, controlled before/after measurements,
self-hosting evidence, source hashes and regression logs.

## Inputs, configuration and protocol

- Same **244 TAST modules, 138,448 type-table entries and 27,122,224 TAST bytes**
  as the earlier reports, `builtWith=0.15.16`. All **544 application input files**
  were checked against the frozen reference before timing. The final snapshot
  has 548 files including compiler artifacts and comparison tooling.
- Flags: `--source output --main Test.Main --threaded --out <fresh-directory>`.
  Both backends process the complete module set.
- Native compiler: **O3, LTO disabled**, native identifier sanitization, stateless
  generation and a bounded ordered codegen pipeline, in addition to the first
  campaign's scanner/Map/PBO optimizations. The default budget on this host is
  8, split **4 PBO + 4 codegen workers**. TAST loading concurrency is 1.
- JavaScript bundle from the same compiler/PBO source tree; sequential PBO and
  codegen, original PureScript identifier sanitizer. Node flags: `--expose-gc
  --stack-size=65536 --max-old-space-size=16384`.
- Freeze the launcher, both compiler artifacts, TAST, PureScript sources,
  adjacent Rust FFI/Cargo declarations and expected application stdout.
- One warm-up per backend, then **five sequential JS→native pairs**. Fresh
  process, output directory and `.purmeta` state per invocation. Normal OS file
  caching, no CPU affinity or cache flush. Retain every measured sample.
- Metric: monotonic **`backend total`**, covering TAST loading/sorting,
  preparation, PBO optimization, Rust generation and emission, including the
  final codegen drain. Excludes process startup/exit, frontend `purs`, compiler
  rebuilding, Cargo and application execution.

Host: **Apple M4 Pro**, 14 CPU cores, **48 GiB RAM**, arm64 macOS **26.1 (25B78)**;
Node **v24.8.0**, Rust/Cargo **1.96.0**. Timing ran from **11:04:11 to 11:05:46 UTC**;
application validation finished at **11:06:13 UTC**. Compiler rebuilds and tests
ran outside the timing interval.

### Phase medians

| Phase | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Load TAST + sort | 601 | 1,714 |
| Prepare | 42 | 704 |
| Optimize + generate | 4,897 | 4,191 |
| Finalize + emit | 868 | 1,526 |

Each phase has its own median; they need not sum to the total median. The native
optimization/generation phase is now faster than JavaScript in this campaign,
while loading, preparation and finalization remain substantially slower.

## Validation and provenance

All **12 invocations** succeeded and produced the same **496 Rust source files
and Cargo manifests**, compared byte for byte. After all measurements, the last
native output built with `cargo build --offline` and passed **47 Aff checks**:
exact expected stdout, empty stderr, exit 0. The runner verified all frozen
hashes after measurement and again after application execution.

The subsequent archive verification rechecked hashes and medians, confirmed the
544 application inputs against the old manifest, and compared every output in
all four second-lot campaigns: **52 outputs, each with 496 identical files**.
It also confirmed that the measured native executable is the installed one.

The measured native binary is the installed stage 2: **452 compiler TAST modules,
282,280 types, 912 identical generated files** between JS and native stage 1,
then a successful stage-2 Cargo build. Stage 2 passed a fresh **152-module /
312-file** smoke test with `PURUST_NATIVE_OK 42`. An initial cleanup-only failure
was recovered by repeating the smoke test with its workspace retained, before
installation. The full Aff suite also passes: **47 checks, 5 Rust unit tests,
concurrency/Ref/AVar/lifetime checks and 9 error-reporting scenarios**.

| Compiler artifact | Bytes | SHA-256 |
|---|---:|---|
| JavaScript | 2,204,314 | `68922e52446516e0739f791c9a24df03ce8153a3c266e7d3a7f4b17c8c9c7735` |
| Native stage 2 | 42,137,872 | `0a0f2afe52159a3c272beeaee625e68d100d5855bd3d99dc18ab6f353eca875e` |

Snapshot revisions: Purust `f842990ec9769df49d6c18de9d3b25d6e070450c`, Aff
`823ebc23d529af9a124f382431ab7a2eee893b93`, PBO
`0f41544464ec0f42e6cb0dd77b206852813f904f`. Working-tree state is recorded in the
JSON, including PBO's modified builder and untracked Rust FFI files. Artifact
hashes identify the actual executables measured.

TAST manifest SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

## Rerunning

From `altbak.pub`, with installed compilers and existing Aff TAST:

```sh
TMPDIR=/path/to/durable/workspaces \
  node bin/benchmark/compilation-purust-aff.mjs \
  ../purust/purust ../purust/purust-aff /path/to/new-results.json
```

Use a new result filename. This campaign's retained workspace is
`var/benchmark/purust-codegen-20261002/purust-aff-compilation-n5wyEI`.
