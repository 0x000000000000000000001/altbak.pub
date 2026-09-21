# Type-table argument traversal — September 22, 2026

Fusing PBO's `resolveArgs` traversal gives a **modest Go improvement** on the
JSON/TAST diagnostic: **590.273 → 573.631 ms (−2.82%)**, with **9.237 MiB less
allocation per corpus (−1.55%)**. JS also improves. This step does not close the
Go/JS gap: the final ratio is 7.56×, versus 7.08× in the paired reference.

## Measurement

The same twelve frozen TAST modules, complete JSON/AST oracle and timed kernels
are used. Each version/runtime has three independent processes, two warm-ups
and five measured samples per phase. Cells are medians of process minima.
Both versions use the same Go runtime correction, Argonaut FFI, compiler emitter
and dependencies; the only changed production source is `CoreFn/TypeTable.purs`.
GOMAXPROCS=1, GOGC=100, PGO off; Node keeps its usual background threads.
Processes run sequentially, alternating before/after and reversing the order
in the middle round. No agent builds or other benchmarks run during measurement.

| Whole corpus | Go before | Go after | JS before | JS after |
|---|---:|---:|---:|---:|
| Parse | 73.697 ms | 77.189 ms | 19.929 ms | 20.116 ms |
| Decode parsed JSON | 492.369 ms | 475.123 ms | 61.034 ms | 58.586 ms |
| Parse + decode | **590.273 ms** | **573.631 ms** | **83.381 ms** | **75.899 ms** |

| Go allocation per corpus | Before | After |
|---|---:|---:|
| Parse | 65.262 MiB | 65.262 MiB |
| Decode parsed JSON | 530.676 MiB | 521.440 MiB |
| Parse + decode | 595.948 MiB | 586.711 MiB |

Decode alone improves by 3.50% in Go and 4.01% in JS; the separately timed
combined phase improves by 2.82% and 8.97%. The parser is unchanged and its
timing variation is not attributed to this patch. Fingerprinting remains outside
timing but allocates between samples. Small timing gains should be read alongside
the repeatable allocation reduction, not treated as universal speedup guarantees.

The immediately preceding [published measurement](2026-09-22-json-typed-ast.md)
was Go 599.956 ms / JS 85.856 ms. The fresh reference here is 590.273 / 83.381 ms;
we use the paired reference to avoid attributing that drift to the change.
The original September 21 diagnostic was Go 717.160 / JS 83.839 ms; it remains
historical context, not the control for this individual patch. No complete b8x
build or default parallel-loading performance was measured in this step.

## Change and correctness

`resolveArgs` previously traversed references in ST, sequenced the resulting
`Maybe` values, then sequenced the `Either` values. One loop now gathers resolved
types, records whether any reference is pending and remembers the first error.
Pending references retain priority over errors, including an error encountered
earlier in argument order. Forced missing references still become `Any`.
The result array stays local and is frozen only after its last mutation.

The first prototype was correct but allocated slightly more in Go. Its generic
`void`/mapping adapters and per-index `unsafePartial` calls survived compilation.
The retained version follows the existing decoder's explicit ST bind style and
moves the partiality boundary outside the bounded loop. Other resolver branches,
TAST information and decoder validation are unchanged.

- The existing JS type-table suite now covers all four `resolveArgs` callers,
  empty/repeated/forward arguments, cycles, missing references and error priority.
- A deterministic differential check compares 1,000 tables with the original
  decoder: 372 successes and 628 errors, all identical, with inputs unchanged.
- The native suite passes 22 cases, including six error/priority cases and
  repeated decoding in each successful case.
- All twelve processes in the final campaign match the fixed complete JSON/AST
  fingerprints. The first prototype's twelve processes also matched.

## Reproduction and next target

Use the existing `bin/go/run --test JsonTypedAst` and
`bin/js/run --test JsonTypedAst` entry points, building before measuring.
The source/fixture layout and measurement boundaries are unchanged. The README
continues to show only milliseconds in these cells and excludes this diagnostic
from the historical fourteen-case totals.

The [archive](2026-09-22-resolve-args.json) retains every final sample, manifests,
the first prototype's measurements, validation and integration checks. Scripts
and raw logs are in `../scratch/resolve-args-20260922`; `compare-direct.py`
reproduces the alternating campaign from the two recorded immutable builds.
The before/after manifests intentionally identify different PBO decoder sources.

This completes the `resolveArgs` substep. Larger remaining candidates are the
row/constraint traversals, their FFI callback/array adaptations, and generated
array conversions. Their gains remain to be measured separately.
