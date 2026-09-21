# JSON to typed AST — September 22, 2026

The complete Go workload falls from **720.986 ms to 599.956 ms (−16.79%)**
in a fresh before/after campaign. Allocation falls by **19.98%**. Against the
[published September 21 baseline](2026-09-21-go-diagnostics.md#json-to-typed-ast)
of 717.160 ms, the reduction is 16.34%. Go still takes **6.99×** the freshly
measured JS time; this change reduces, but does not close, the gap.

## Measurements

Same frozen twelve-module corpus, source kernels, structural oracle and timing
protocol as the original diagnostic. Three processes per runtime, two warm-ups
and five measured corpus passes per phase; cells are the median of process
minima. GOMAXPROCS=1, GOGC=100 and PGO off for Go. Node retains its default
background threads. No builds or other agent benchmarks ran during measurement.

| Entire corpus | Go before | Go after | JS control | Go change |
|---|---:|---:|---:|---:|
| Parse | 76.483 ms | 75.850 ms | 20.219 ms | −0.83% |
| Decode parsed JSON | 616.670 ms | 507.676 ms | 62.653 ms | −17.67% |
| Parse + decode | **720.986 ms** | **599.956 ms** | **85.856 ms** | **−16.79%** |

| Go allocations per corpus | Before | After |
|---|---:|---:|
| Parse | 65.262 MiB | 65.262 MiB |
| Decode parsed JSON | 679.509 MiB | 530.676 MiB |
| Parse + decode | 744.781 MiB | 595.948 MiB |

The parser is unchanged; its small timing variation is not an optimization
claim. The JS bundle is byte-identical to the baseline build. Its 85.856 ms
control versus the historical 83.839 ms illustrates run-to-run variation.
Combined timing is measured independently, not added from phase minima.
Fingerprint work remains outside timing but allocates between batches, as before.

The three Go baseline processes ran first. After reconstruction and all test
builds completed, the final campaign ran Go/JS/JS/Go/Go/JS sequentially. These
numbers describe this benchmark, not the full compiler or its default parallel
loading configuration. No new complete b8x build was timed.

## Cause and correction

A separate Go benchmark profiles decoding already parsed JSON, excluding the
fingerprint serializer from its loop. Initially, `gopurs_runtime.Apply` accounts
directly for 26.31% of sampled allocated bytes. Much of that is temporary partial
application: `Apply6` through `Apply10` invoked `Apply` repeatedly even when all
arguments of a matching `FuncN` were already available. JSON dispatch uses a
seven-argument function frequently.

The runtime now calls saturated functions directly. Partial applications capture
the supplied arguments once; lower arities retain sequential application for
functions returning functions. Existing heap-escape guarantees are retained.
The optimization applies to `Apply2` through `Apply10`, up to `Func11`.

`gopurs-argonaut-core` also passes the seven arguments and return value of
`CaseJsonImpl` directly as `gopurs_runtime.Value`, avoiding the former interface
conversions. Its JSON dispatch semantics are unchanged. PBO decoding, validation,
the JSON parser and the test workload are unchanged.

An initial isolated dispatch-only experiment reduced decode time from about
610 to 509 ms and allocations from 712.5 to 577.7 MB. This is a separate Go
benchmark protocol, not a replacement for the table above. In the final profile,
direct allocations attributed to `Apply` fall from approximately 2.30 GiB to
271.51 MiB over the profiling run. Remaining allocations are distributed across
partial applications, records, `Either`/`Maybe`, traversals and FFI adapters.
Sampled cumulative costs overlap and must not be added.

## Validation and follow-up

- All JSON and complete decoded-AST fingerprints match the fixed reference across
  all nine before/after processes.
- Of 503 generated Go files, 501 are identical. Only the runtime and Argonaut
  dispatch FFI change; no files are added or removed.
- Runtime tests pass normally and under the race detector: 99 arity/application
  combinations, 27 panic cases, zero saturated-call allocations for arities 2–10,
  and the existing closure lifetime and concurrency suite.
- The native Argonaut suite passes: 32 assertions/properties, including 536
  QuickCheck cases. Nine runner/publication checks also pass.
- Both JS and native gopurs compilers were rebuilt; the native binary is installed
  locally. Benchmark manifests now also fingerprint native library sources, so
  changing a library FFI invalidates a stale build.

The next targeted investigation is the remaining type-table traversals and their
callback/array conversions. Repeated `traverse`/`sequence` stages and generated
array unbox/rebox pairs are candidates, not measured future gains. Decoder error
priority and all TAST information must remain intact.

## Reproduction and evidence

The normal entry points and source layout are unchanged:

```sh
./bin/go/run --test JsonTypedAst --build-only
./bin/go/run --test JsonTypedAst --run-only
./bin/js/run --test JsonTypedAst --build-only
./bin/js/run --test JsonTypedAst --run-only
```

The [measurement archive](2026-09-22-json-typed-ast.json) includes every sample,
source/binary hashes, log hashes, validation and generated-file comparison.
Raw campaigns are in `var/benchmark/json-diagnostic/{before,after}-20260922`.
The retained build is under
`run/bak/go/modes/test-JsonTypedAst/fast-apply-20260922/builds/20260921T222710.474340Z-21380`.
The final campaign used the underlying `bin/benchmark/json-diagnostic.py measure`
helper without `--runtime` to alternate both runtimes from that single build.
Isolated experiments, binaries, CPU/allocation profiles and library test logs
are retained in `../scratch/json-decode-20260922`.
