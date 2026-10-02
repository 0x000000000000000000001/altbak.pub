# Purust Aff compilation after native optimization

## Published result

Measured on **2026-10-02**, using the self-rebuilt stage-2 compiler, with native
PBO parallelism enabled by default:

| Backend | Median `backend total` | Range across five runs |
|---|---:|---:|
| `purust.js`, sequential PBO | **6 774 ms** | 6 626–6 935 ms |
| `purust-native`, 8 PBO workers | **10 328 ms** | 10 268–10 504 ms |

The native compiler takes **1.52× the JavaScript time** on this workload.
The [previous published ratio](2026-10-01-purust-aff-compilation.md) was 2.66×.
The native compiler has improved but has not yet beaten JavaScript; a ratio of
0.3 remains a longer-term objective.

A separate five-round comparison of the original and optimized native binaries
measures **17 111 → 9 935 ms**, a **41.9% reduction** (1.72× speedup). See the
[controlled comparison](2026-10-02-purust-compiler-optimization.md#final-controlled-oldnew-comparison).
Those samples are separate from the JS/native table above.

| Pair, in execution order | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Warm-up, excluded | 6 936 | 10 562 |
| 1 | 6 792 | 10 268 |
| 2 | 6 626 | 10 504 |
| 3 | 6 935 | 10 306 |
| 4 | 6 739 | 10 366 |
| 5 | 6 774 | 10 328 |

The [raw archive](2026-10-02-purust-aff-compilation.json) contains every sample,
phase time, stdout/stderr log, compiler hash, input hash, checkout state and
application-validation result. The [optimization campaign report](2026-10-02-purust-compiler-optimization.md)
details the individual changes and their controlled comparisons.

## Configuration and protocol

- **Same application inputs as 2026-10-01:** 244 TAST modules, 138 448 type-table
  entries, 27 122 224 bytes of TAST, `builtWith=0.15.16`. All 544 application
  input files were checked against the original frozen manifest before the
  campaign; the complete new snapshot contains 548 files including compilers
  and comparison tooling.
- Flags: `--source output --main Test.Main --threaded --out <fresh-directory>`.
  Both backends process the complete module set.
- The native compiler is built with **`opt-level=3`, LTO disabled**, including
  the scanner optimization, native PBO Maps, and the corrected parallel builder.
  On this host it uses **8 PBO workers**. Rust generation remains sequential and
  ordered. TAST loading concurrency is 1.
- The JavaScript bundle is built from the same compiler/PBO sources and uses
  sequential PBO. Node flags are `--expose-gc --stack-size=65536
  --max-old-space-size=16384`.
- Freeze the launcher, JS bundle, native executable, TAST, PureScript sources,
  adjacent Rust FFI/Cargo declarations and expected application stdout.
- One warm-up per backend, then five sequential JS→native pairs. Fresh process,
  output directory and `.purmeta` state for each invocation. Retain every
  measured sample. Normal OS file caching, no CPU affinity or cache flush.
- Report the median `[purust] backend total`: TAST loading/sorting, preparation,
  PBO optimization, Rust generation and emission. Process startup/exit, frontend
  `purs`, compiler rebuilding, Cargo and application tests are excluded.

Host: **Apple M4 Pro**, 14 CPU cores (10 performance, 4 efficiency), **48 GiB RAM**,
arm64 macOS **26.1 (25B78)**. Node **v24.8.0**, Rust/Cargo **1.96.0**.
Timing ran from **07:43:30 to 07:45:16 UTC**; application validation finished at
**07:45:35 UTC**. Compiler builds and tests were outside the timing interval.

### Phase medians

| Phase | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Load TAST + sort | 644 | 1 843 |
| Prepare | 41 | 727 |
| Optimize + generate | 5 167 | 6 002 |
| Finalize + emit | 869 | 1 636 |

Each phase has its own median; they do not necessarily sum to the total median.
Separate worker diagnostics identify about 5.4 seconds of sequential codegen
callbacks as the main remaining bottleneck.

## Validation and provenance

All **12 invocations** succeeded and produced **496 byte-identical Rust source
files and Cargo manifests**, compared to the first JavaScript output.
After timing, the last native output built successfully with `cargo build
--offline` and passed all **47 Aff checks**: exact expected stdout, empty stderr,
exit code 0. Every frozen input/compiler hash was verified after measurement
and again after application execution.

The installed native executable comes from a validated self-reconstruction:
451 compiler TAST modules, 910 Rust sources/manifests identical between Node
and native stage 1, a successful stage-2 Cargo build, and a fresh-project stage-2
smoke test with 152 modules, 312 identical files, and `PURUST_NATIVE_OK 42`.
The installed compiler also passes the full Aff suite: 47 Aff checks, 5 Rust
unit tests, concurrency/Ref/AVar/lifetime checks and 9 error-reporting scenarios.

| Compiler artifact | Bytes | SHA-256 |
|---|---:|---|
| JavaScript | 2 203 891 | `3f71a4ab68c8655d4dce0a304f00379218a1811a8ba68dc6aa8a12b343c9b237` |
| Native stage 2 | 42 014 784 | `78eb3406774d118c28d8e709713514d6ccb9148d5dfcf83cf29631fe945e4b37` |

Snapshot checkout revisions: Purust `1ae554598c7dd1b952acb24f1f43630924d338fc`,
Aff `823ebc23d529af9a124f382431ab7a2eee893b93`, PBO
`0f41544464ec0f42e6cb0dd77b206852813f904f`. The JSON records the working-tree
changes and untracked Rust FFI files. The compiler artifact hashes identify the
actual binaries measured; the optimization archive also retains self-hosting
logs and source/generated-file hashes.

TAST manifest SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.

## Rerunning

From `altbak.pub`, with current installed compilers and existing Aff TAST:

```sh
TMPDIR=/path/to/durable/workspaces \
  node bin/benchmark/compilation-purust-aff.mjs \
  ../purust/purust ../purust/purust-aff /path/to/new-results.json
```

Use a new result filename. The runner retains frozen inputs, compilers, generated
outputs and logs; its result records the workspace path. This campaign's
workspace is `var/benchmark/purust-compiler-20261002/purust-aff-compilation-9JjOXI`.
