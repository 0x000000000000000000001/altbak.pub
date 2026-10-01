# Purust Aff compilation: JavaScript versus native

## Published result

Measured on **2026-10-01**, using `Test.Main` from `purust-aff` with `--threaded`:

| Backend | Median `backend total` | Range across five runs |
|---|---:|---:|
| `purust.js` | **6 480 ms** | 6 401–6 964 ms |
| `purust-native` | **17 214 ms** | 16 467–17 401 ms |

The native compiler takes **2.66× the JavaScript time** on this workload.

| Pair, in execution order | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Warm-up, excluded from medians | 6 296 | 16 987 |
| 1 | 6 401 | 16 467 |
| 2 | 6 427 | 16 696 |
| 3 | 6 964 | 17 288 |
| 4 | 6 561 | 17 214 |
| 5 | 6 480 | 17 401 |

The [measurement archive](2026-10-01-purust-aff-compilation.json) contains all
stdout/stderr logs, individual phase and process-wall timings, compiler hashes,
per-file input hashes, checkout states and validation results.

## Protocol

- Freeze the existing TAST, its PureScript sources, adjacent Rust FFI files and
  Cargo declarations, and the installed compiler launcher, JS bundle and native
  executable in an isolated workspace. Verify their hashes after measurement
  and again after validation.
- Input: **244 TAST modules**, **138 448 type-table entries**, **27 122 224 bytes**
  of `corefn.json`; `builtWith` is `0.15.16`. Both backends process the full module
  set and emit a `Test.Main` entrypoint.
- Invoke the frozen `bin/purust` launcher with
  `--source output --main Test.Main --threaded --out <fresh-directory>`.
  Select JavaScript with `PURUST_JS=1` and native with `PURUST_JS=0`.
- JavaScript uses Node with `--expose-gc --stack-size=65536
  --max-old-space-size=16384`, matching the launcher.
- Start a fresh process for every invocation, with a fresh output directory and
  no `.purmeta` cache. Perform one warm-up per backend, then five sequential
  JS→native pairs. Keep every successful measured sample.
- Report the median of the five `[purust] backend total: … ms` values for each
  backend. This monotonic-clock metric includes TAST loading/sorting,
  preparation, PBO optimization, Rust generation and emission. It excludes
  process startup/exit, the `purs` frontend, compiler rebuilding, Cargo and
  application tests.
- Use normal OS file caching without a cache flush or explicit CPU affinity.
  The warm-ups are separate processes; they do not preserve a JavaScript VM
  between measured invocations.

Host: **Apple M4 Pro**, 14 CPU cores (10 performance, 4 efficiency), **48 GiB RAM**,
ARM64, macOS **26.1 (25B78)**. Node **v24.8.0**. Application validation uses
Rust/Cargo **1.96.0**. Backend measurements ran from 19:56 to 19:58 UTC;
validation finished at 19:59 UTC.

### Phase medians

| Phase | JavaScript (ms) | Native (ms) |
|---|---:|---:|
| Load TAST + sort | 632 | 1 985 |
| Prepare | 40 | 1 002 |
| Optimize + generate | 4 902 | 11 554 |
| Finalize + emit | 848 | 2 506 |

Each phase has its own median. These values need not sum to the total median;
the total also includes work between the nested phase timers.

## Validation

Every warm-up and measured invocation exited successfully. Using Purust's
`compareGeneratedSources`, **496 Rust source files and Cargo manifests were
byte-identical** to the first JavaScript output in all 12 invocations.

After timing finished, the fifth native output was built with
`cargo build --offline` in the default development profile. Its executable
exited with code 0, produced no stderr, and matched
`test/expected-main.stdout` exactly: **47 Aff checks passed**. Compiler and
input hashes remained unchanged throughout the campaign.

## Provenance and rerunning

Recorded checkout revisions at snapshot time:

- Purust: `16e9a60249267a37316dddda656eacb9b6d30b93`
- `purust-aff`: `823ebc23d529af9a124f382431ab7a2eee893b93`
- PBO: `0f41544464ec0f42e6cb0dd77b206852813f904f`

The JSON also records working-tree status, including the optimizer's untracked
native FFI files. The frozen artifact hashes identify the actual compilers
measured:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `bin/purust.js` | 2 156 571 | `c1eed286e854808837d7e4f5c8609a9a25526c95b888c19c1c5c46670b2b393b` |
| `bin/purust-native` | 55 978 832 | `7b53d20694b54d09f69fc890432ca36df865b25ee9e0db558e43130bcddecda5` |

TAST manifest SHA-256:
`b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.
This hashes the concatenated, path-sorted `<file-sha256>  <snapshot-path>\n`
entries for the 244 `corefn.json` files. The complete per-file manifest is in
the JSON archive.

Run the same protocol on the currently installed compilers and existing
`purust-aff/output` from the `altbak.pub` directory:

```sh
node bin/benchmark/compilation-purust-aff.mjs \
  ../purust/purust ../purust/purust-aff \
  /path/to/new-results.json
```

The [runner](../../bin/benchmark/compilation-purust-aff.mjs) preserves its
workspace, frozen inputs/compilers, generated sources and validation logs.
The output JSON records that workspace path. Set `TMPDIR` to choose its parent;
use a new result filename for each campaign.
