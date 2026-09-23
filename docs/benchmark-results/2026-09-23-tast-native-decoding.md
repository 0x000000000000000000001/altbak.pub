# Native TAST decoding: −87.5% decode time

Four native paths now run behind the existing per-backend FFI boundary: CoreFn
type-table resolution, the `decodeArray` element loop, the CoreFn source-usage
validation, and the annotation decoder (`decodeAnnWithUsage`). The PureScript
implementations remain the JavaScript path through explicit fallback arguments,
so the JS bundle keeps its previous algorithms. On the frozen TAST diagnostic
the Go cells improve from **358.78 / 396.74 ms** to **40.11 / 107.58 ms**
(decode / combined); the Go decode cell is now **faster than the JS control**
(67.79 ms).

## Result

Cumulative paired campaign from the published state to the final state, five
pairs, both workspaces regenerated, oracles validated in every process
(`var/benchmark/native-ann-cumulative-20260923`):

| Phase | Before | After | Paired delta | Allocations |
|---|---:|---:|---:|---:|
| Parse | 28.423 ms | 27.582 ms | −2.96% (noise) | 56,031,272 (identical) |
| **Decode** | 366.611 ms | **45.805 ms** | **−87.51%** | 402,378,408 → **62,058,680 B (−84.58%)** |
| Combined | 393.177 ms | **101.585 ms** | **−74.16%** | 458,409,696 → 118,089,888 B (−74.24%) |

Per-step paired campaigns:

| Step | Decode | Allocations |
|---|---:|---:|
| Type table (`json-tast-stf-20260923` → `json-tast-native-tt-20260923`) | 363.408 → 186.058 ms (−48.80%) | −46.94% |
| `decodeArray` loop (`native-tt` → `native-arr`) | 184.690 → 165.835 ms (−10.21%) | −6.84% |
| Usage validation (`native-arr` → `native-usage2`) | 177.478 → 96.776 ms (−45.47%) | −41.03% |
| Annotation decoder (`native-usage2` → `native-ann`) | 99.934 → 39.018 ms (−60.96%) | −47.09% |

Official protocol on the final workspace (three processes per backend, median
of process minima, `var/benchmark/native-ann-results-20260923`):

| Cell | Before | After |
|---|---:|---:|
| Go parse / decode / combined | 28.23 / 358.78 / 396.74 ms | 29.19 / **40.11** / **107.58 ms** |
| JS control parse / decode / combined | 20.86 / 60.64 / 81.77 ms | 20.84 / 67.79 / 87.64 ms |

The ratio to JS moves from **×5.92 to ×0.59** on decode (the native cell is
1.7× faster than the JS control) and from **×4.85 to ×1.23** on combined. The
remaining combined gap is now dominated by parsing (Go 29.19 ms vs JS
20.84 ms) and by the combined phase's interaction between parse allocations and
decoding; the standalone decode cell is ahead of JS.

Allocations per corpus: parse 56,031,272, **decode 62,058,680** (was
402,378,424), combined 118,089,888.

The JavaScript implementations are unchanged: each FFI function receives its
PureScript fallback as an argument and the JS FFI calls it, so the JS bundle
runs the same algorithms with one extra call. The JS movement between sessions
(the cumulative paired campaign gives +6.48% decode, +3.41% combined) is not
attributed to these changes. The JS bundle is no longer byte-identical, and
README cells are updated from the official cells of this session.

## Why the generated path was expensive

A decode-only profile of the published state attributed **51.3% of allocations
and 30% of CPU samples to the type table** (21,574 types, ~11 KB and ~280
objects per type for a 15 MB retained AST). The generated code materialises the
ST monad, `Maybe (Either …)` layers and one closure per step, and every JSON
field read crosses `Object.lookup`/`caseJson` boxing. Later profiles showed two
more large transient blocks: the usage validation (`StateT` + `Map`/`Set`,
paid after decoding every module) and the annotation decoder, which ran for
every node of the AST and rebuilt `Ann` records, `Maybe` layers and usage
records through the same generic plumbing.

The scratch experiments also isolated the GC share: at `GOGC=1200` the
unchanged baseline decoded in 164.6 ms instead of 322.0 ms, and the native
variant in 88.3 ms instead of 163.1 ms. A large part of the previous cost was
GC pressure from transient allocation, not decoder work.

## Implementations

- **Type table** (`CoreFn/TypeTable.purs`, `CoreFn/Json.go`): native decode and
  fixed-point resolution with the same rounds, force phase, first-error and
  deferred-argument rules.
- **`decodeArray`** (`CoreFn/Json.purs`, `CoreFn/Json.go`): direct element loop
  preserving the first-error `AtIndex` wrapping.
- **Usage validation** (`CoreFn/Usage.purs`, `CoreFn/Usage.go`): the same
  traversal and scope/seen bookkeeping as `validateSourceUsageModule`, without
  `StateT`, `Map`/`Set` or `Maybe`/`Either` plumbing.
- **Annotation decoder** (`CoreFn/Json.purs`, `CoreFn/Json.go`):
  `decodeAnnWithUsage` builds the canonical `Ann` record directly, including
  metadata (`IsConstructor` with its raw constructor-type tag), the optional
  `ExprType` index and the usage records.

All four follow the existing per-backend FFI pattern in this repository (the
native JSON parser, the value-native ST interface). The Go code is hand-written
FFI, not compiler-generated code. Each `foreign import` receives its PureScript
implementation as its first argument; the Go backend ignores it and the JS
backend calls it.

## Validation

- Differential test of the native type table against the PureScript one: **20
  fixed cases plus 400 seeded random tables**, comparing error strings and
  every decoded type (self- and two-entry cycles, out-of-range and negative
  references, deferred argument errors, `Row` tails, `ForAll`,
  `ConstrainedType`, `TypeApp` chains, missing fields).
- The compiler bootstrap decodes its own 458-module TAST with the native
  paths. Three semantic bugs were caught there before any measurement (two in
  the usage port, one in the annotation port).
- The 12 module fingerprints pass in every campaign and official process; the
  fingerprint encodes the full annotated AST, so the annotation representation
  is checked structurally.
- `gopurs` native bootstrap completes.

## General JSON suite (unchanged path)

The `JsonDecoding` suite does not import the backend optimizer, so these
changes cannot affect it. It was re-measured in this session anyway: Go
parse/decode/combined **2.57 / 12.44 / 16.59 ms**, JS **1.76 / 7.85 / 9.09 ms**.
README cells for that suite stay at the previous official values (Go
**15.11 ms**, JS **9.28 ms**).

## Provenance

- Sources: `CoreFn/Json.purs`, `CoreFn/Json.js`, `CoreFn/Json.go`,
  `CoreFn/TypeTable.purs`, `CoreFn/Usage.purs`, `CoreFn/Usage.js`,
  `CoreFn/Usage.go` in `purescript-backend-optimizer-gopurs`.
- Compiler: `gopurs/bin/gopurs-native`.
- Workspaces: `var/benchmark/json-tast-stf-20260923` (before),
  `json-tast-native-tt-20260923`, `json-tast-native-arr-20260923`,
  `json-tast-native-usage2-20260923`, `json-tast-native-ann-20260923` (after).
- Reports of the exploratory variants: `scratch/tast-revolution-20260923/`.
- Reproduce:
  1. `cd gopurs/gopurs && npm run build:native`
  2. `cd altbak.pub-gopurs && python3 bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace <new>`
  3. `python3 bin/benchmark/json-diagnostic.py measure --suite JsonTypedAst --workspace <new> --output <out>`
  4. Paired: `python3 scratch/tast-decode-20260923/campaign.py --suite JsonTypedAst --before var/benchmark/json-tast-stf-20260923 --after <new> --output <dir> --pairs 5`

## Limits and next steps

- The expression/binder decoder (`decodeExpr`, `decodeBinder`, `decodeLiteral`,
  `decodeBind`, `decodeModule'`) still runs as generated PureScript code. It is
  the remaining decode cost and the next candidate; parsing is now the larger
  half of the combined cell.
- The native paths' failure paths are not covered by the diagnostic corpus (all
  12 modules are valid). They mirror the PureScript code structure; the
  compiler bootstrap exercises valid modules only.
- b8x wall time was not measured for these changes.
