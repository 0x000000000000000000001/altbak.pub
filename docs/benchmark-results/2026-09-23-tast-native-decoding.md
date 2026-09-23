# Native TAST decoding: −95% decode time, both cells ahead of JS

The whole CoreFn decoder now runs as native Go behind the existing per-backend
FFI boundary: type-table resolution, the `decodeArray` element loop, the
source-usage validation, the annotation decoder and the module decoder
(`decodeModule` with `decodeModule'`, expressions, binders, literals, binds,
imports and declarations). The PureScript implementations remain the JavaScript
path through explicit fallback arguments, so the JS bundle keeps its previous
algorithms. On the frozen TAST diagnostic the Go cells improve from
**358.78 / 396.74 ms** to **15.68 / 57.32 ms** (decode / combined) and are now
**faster than the JS control** (60.78 / 81.70 ms).

## Result

Cumulative paired campaign from the published state to the final state, five
pairs, both workspaces regenerated, oracles validated in every process
(`var/benchmark/native-dec5-cumulative-20260923`):

| Phase | Before | After | Paired delta | Allocations |
|---|---:|---:|---:|---:|
| Parse | 26.655 ms | 26.222 ms | −1.62% (noise) | 56,031,272 (identical) |
| **Decode** | 336.613 ms | **15.156 ms** | **−95.50%** | 402,378,424 → **27,106,040 B (−93.26%)** |
| Combined | 386.411 ms | **48.128 ms** | **−87.54%** | 458,409,680 → 83,137,248 B (−81.86%) |

Per-step paired campaigns:

| Step | Decode | Allocations |
|---|---:|---:|
| Type table (`stf` → `native-tt`) | 363.408 → 186.058 ms (−48.80%) | −46.94% |
| `decodeArray` loop (`native-tt` → `native-arr`) | 184.690 → 165.835 ms (−10.21%) | −6.84% |
| Usage validation (`native-arr` → `native-usage2`) | 177.478 → 96.776 ms (−45.47%) | −41.03% |
| Annotations (`native-usage2` → `native-ann`) | 99.934 → 39.018 ms (−60.96%) | −47.09% |
| Module decoder (`native-ann` → `native-dec`) | 37.275 → 16.302 ms (−56.27%) | −53.71% |
| Decoder follow-up (`native-dec` → `native-dec5`) | 16.254 → 15.156 ms | −5.6% (allocation work) |

Official protocol on the final workspace (three processes per backend, median
of process minima, `var/benchmark/native-dec5-results-20260923`, quiet machine):

| Cell | Before | After |
|---|---:|---:|
| Go parse / decode / combined | 28.23 / 358.78 / 396.74 ms | 26.01 / **15.41** / **48.00 ms** |
| JS control parse / decode / combined | 20.86 / 60.64 / 81.77 ms | 20.98 / 59.54 / 78.19 ms |

The ratio to the JS control moves from **×5.92 to ×0.26** on decode (the Go
cell is 3.9× faster) and from **×4.85 to ×0.61** on combined (Go 1.6× faster).
Allocations per corpus: parse 56,031,272, **decode 27,106,040** (was
402,378,424), combined 83,137,248.

The JavaScript implementations are unchanged: each FFI function receives its
PureScript fallback as an argument and the JS FFI calls it. The JS movement
between sessions (±5%) is not attributed to these changes. The JS bundle is no
longer byte-identical, and README cells are updated from the official cells of
this session.

## Phase 1 follow-up: error paths and decoder allocation

- **Differential error-path test** (native vs PureScript, mutated corpus,
  deterministic): **2,210 mutations over six seeds — 1,493 error cases and 717
  successful decodes** — comparing error strings and fingerprints. It found and
  fixed five semantic divergences in the native decoder: `Object`/`Array`
  messages for `decodeJObject`/`decodeJArray`, the `AtKey` wrapping of
  `getFieldOptional'`, the `AtIndex` wrapping of `decodeModuleName`, the `Array`
  message for a non-array type table, and the missing per-key wrapping rule in
  `decodeReExports`. The final run passes with identical messages and
  fingerprints, on the integrated compiler output.
- **Allocation reductions**: native type-table entry (no `Box` round trip),
  removal of the per-call closure in `cndExpr`, and persistent scopes in the
  usage validation (a frame chain instead of copying a map per binding).
  Decode allocations per corpus: **28,724,952 → 27,106,040 B (−5.6%)**.
- The decoder re-measured on a quiet machine: official cells **15.41 / 48.00 ms**
  (see the Result section); the earlier sessions during an external
  `purs compile` load (load average ~20) showed decode between 15.4 and 16.7 ms
  and combined between 47.6 and 66.9 ms. The paired campaigns are the
  authoritative comparison.
- Documented couplings: the native decoder calls the generated
  `decodeSourceSpan` and `Map.fromFoldable` specialisations for two cold paths
  (source positions and the foreign-annotation map); both call sites carry a
  comment explaining why the coupling is accepted and that a rename fails the
  build loudly.

## Why the generated path was expensive

A decode-only profile of the published state attributed **51.3% of allocations
and 30% of CPU samples to the type table** (21,574 types, ~11 KB and ~280
objects per type for a 15 MB retained AST). The generated code materialises the
ST monad, `Maybe (Either …)` layers and one closure per step, and every JSON
field read crosses `Object.lookup`/`caseJson` boxing. Later profiles showed the
same pattern in the usage validation (`StateT` + `Map`/`Set` after decoding
every module), the annotation decoder (every node rebuilds `Ann` records,
`Maybe` layers and usage records) and the expression decoder itself.

The scratch experiments also isolated the GC share: at `GOGC=1200` the
unchanged baseline decoded in 164.6 ms instead of 322.0 ms, and the native
type-table variant in 88.3 ms instead of 163.1 ms. A large part of the previous
cost was GC pressure from transient allocation, not decoder work.

## Implementations

All in `purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/CoreFn`:

- **Type table** (`TypeTable.purs`, `Json.go`): native decode and fixed-point
  resolution with the same rounds, force phase, first-error and deferred
  argument rules.
- **`decodeArray`** (`Json.purs`, `Json.go`): direct element loop preserving
  the first-error `AtIndex` wrapping.
- **Usage validation** (`Usage.purs`, `Usage.go`): the same traversal and
  scope/seen bookkeeping as `validateSourceUsageModule`, without `StateT`,
  `Map`/`Set` or `Maybe`/`Either` plumbing.
- **Annotation decoder** (`Json.purs`, `Json.go`): canonical `Ann` records,
  metadata (`IsConstructor` with its raw constructor-type tag), the optional
  `ExprType` index and the usage records.
- **Module decoder** (`Json.purs`, `Json.go`): `decodeModule` native end to
  end, including expressions, binders, literals, binds, imports, data/class
  declarations, re-exports, comments and the foreign-annotation map. Cold
  helpers with intricate generated specialisations (source spans, the foreign
  `Map`) are called through their generated functions and converted to
  canonical records.

All five follow the existing per-backend FFI pattern in this repository (the
native JSON parser, the value-native ST interface). The Go code is hand-written
FFI, not compiler-generated code. Each `foreign import` receives its PureScript
implementation as its first argument; the Go backend ignores it and the JS
backend calls it.

## Validation

- Differential test of the native module decoder against the generated
  PureScript decoder on the frozen 12-module corpus: identical fingerprints,
  and both match the fixture hashes.
- Differential test of the native type table against the PureScript one: **20
  fixed cases plus 400 seeded random tables**, comparing error strings and
  every decoded type (cycles, out-of-range and negative references, deferred
  argument errors, `Row` tails, `ForAll`, `ConstrainedType`, `TypeApp` chains,
  missing fields).
- The compiler bootstrap decodes its own 458-module TAST with the native
  paths; three semantic bugs were caught there before any measurement (two in
  the usage port, one in the annotation port).
- The 12 module fingerprints pass in every campaign and official process.
- `gopurs` native bootstrap completes.

## General JSON suite (unchanged path)

The `JsonDecoding` suite does not import the backend optimizer, so these
changes cannot affect it. It was re-measured earlier in the session: Go
parse/decode/combined **2.57 / 12.44 / 16.59 ms**, JS **1.76 / 7.85 / 9.09 ms**.
README cells for that suite stay at the previous official values (Go
**15.11 ms**, JS **9.28 ms**).

## Provenance

- Sources: `CoreFn/Json.purs`, `CoreFn/Json.js`, `CoreFn/Json.go`,
  `CoreFn/TypeTable.purs`, `CoreFn/Usage.purs`, `CoreFn/Usage.js`,
  `CoreFn/Usage.go`.
- Compiler: `gopurs/bin/gopurs-native`.
- Workspaces: `var/benchmark/json-tast-stf-20260923` (before),
  `json-tast-native-tt-20260923`, `json-tast-native-arr-20260923`,
  `json-tast-native-usage2-20260923`, `json-tast-native-ann-20260923`,
  `json-tast-native-dec-20260923` (after).
- Scratch differential tests and exploratory variants:
  `scratch/tast-revolution-20260923/`.
- Reproduce:
  1. `cd gopurs/gopurs && npm run build:native`
  2. `cd altbak.pub-gopurs && python3 bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace <new>`
  3. `python3 bin/benchmark/json-diagnostic.py measure --suite JsonTypedAst --workspace <new> --output <out>`
  4. Paired: `python3 scratch/tast-decode-20260923/campaign.py --suite JsonTypedAst --before var/benchmark/json-tast-stf-20260923 --after <new> --output <dir> --pairs 5`

## Limits and next steps

- The native decoder's failure paths are not covered by the diagnostic corpus
  (all 12 modules are valid). They mirror the PureScript code structure; the
  compiler bootstrap exercises valid modules only.
- The remaining Go costs are the JSON parse phase and the rest of the compiler
  pipeline; the decoder is no longer the bottleneck of this diagnostic.
- b8x wall time was not measured for these changes.
