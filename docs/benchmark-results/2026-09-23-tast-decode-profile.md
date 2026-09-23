# Renewed TAST decode profile and b8x reference — 23 September 2026

This campaign re-establishes the reference cells with the current compiler and
attributes the TAST decode cost on the generated Go. Nothing was optimised in
this step; all numbers are measurements of the published state.

## Summary

- Reference cells reproduced: TAST decode **371.00 ms Go / 56.37 ms JS**,
  combined **423.14 / 79.50 ms**, decode allocations **442,632,472 bytes**
  (identical to the published record-plan campaign). b8x `b -c` completes in
  **250.01 s wall** with a **160.21 s** backend.
- Decode-only profiling: **442.6 MB and 10.88 M objects allocated per pass**
  while the decoded AST retains only about **15 MB**. **~96% of the allocated
  bytes and objects are transient.**
- The **type table path is 54.8% of decode allocations** (21,574 types;
  ~11 KB and ~280 objects per type). The settle loop itself is not quadratic:
  a simulation of the exact algorithm gives **2.47 resolution attempts per
  type** (10–12 rounds, 53,189 attempts).
- The decoder's own generated code is only **~5% of allocated objects**.
  The rest is runtime and foreign-interface glue: `gopurs_runtime` 19.6%,
  gopurs-st ST adapters 16.2%, `Either`/`Maybe` constructor accessors 8.9%,
  `Foreign.Object` adapters 4.8%, `StateT` 3.8%, `Data.Ord` 3.0%,
  `Record_Unsafe` 2.1%, `Partial.Unsafe` 1.3%.
- **First bounded experiment (scratch probe)**: replacing the ST foreign
  wrappers with `Value`-native helpers on the frozen generated Go gives
  **−13.4% decode time and −9.1% allocated bytes**, oracles passing. This
  measures the ceiling of the ST/foreign-interface layer; porting it into
  gopurs is the next step.

## Reference measurements

### TAST diagnostic (12 modules, 21,574 types, 5.55 MB of JSON)

Fresh workspace built with the installed native compiler; official protocol
(GOMAXPROCS=1, GOGC=100, PGO off, three processes per backend, two warmups and
five samples per phase, median of process minima). Oracles: all 12 modules and
their fingerprints.

| Phase | Go | JS |
|---|---:|---:|
| Parse | 25.24 ms | 20.41 ms |
| Decode | **371.00 ms** | **56.37 ms** |
| Combined | **423.14 ms** | **79.50 ms** |

Allocations per corpus: parse **56,031,272**, decode **442,632,472**, combined
**498,663,744 bytes**. Decode ratio **×6.58**; combined **×5.32**.

### Complete b8x build (`b -c`, native compiler rebuilt by the script)

| Cell | Value |
|---|---:|
| Exit code | 0 |
| Wall | **250.01 s** |
| Child user + system | 919.19 s + 40.76 s |
| Peak RSS | **5.70 GiB** |
| `load TAST + sort` | 9,638 ms |
| `transitive specializations` | 30,008 ms |
| `prepare + monomorphize` | 38,991 ms |
| `optimize + emit` | 111,579 ms |
| `backend total` | **160,212 ms** |

Compiler SHA-256 `4a10eaf6…`. This is the baseline for the next changes; the
previous published full build was 181.22 s backend / 267.70 s command, so the
current state already improves it. The load phase now uses the native default
of 8 workers.

## Decode-only profile

Scratch copy of the generated output (`scratch/tast-decode-20260923`), with the
reference FFI driver replaced so that the corpus is parsed once, the decoded
fingerprints come from the fixture oracle, and the decode runs in a loop.
`validate` mode re-checks all 12 module fingerprints in-process before any
profile is published.

| Run | Iters | min | median | Allocations per pass |
|---|---:|---:|---:|---:|
| CPU (`cpu.prof`) | 20 | 379.41 ms | 388.10 ms | — |
| Allocations (`mem.prof`, rate 16384) | 5 | 539.09 ms | 556.42 ms | **442,632,773 B** |
| Validate | 3 | 372.51 ms | 381.98 ms | — |

The allocation figure matches the reference cell within 300 bytes, so the
profiled workload is the published one. Objects: **10,884,656 per pass**.
The CPU profile is GC-dominated: ~**74%** of flat samples are in the runtime
allocator/GC, `madvise` alone 21%.

### Retained output

Decode once, drop the loop, keep the results alive, GC twice, then heap
profile: **22.4 MB live**, of which 7.0 MB is the parsed input retained by the
driver. The decoded AST is therefore around **15 MB** against 442.6 MB
allocated per pass.

### Allocation attribution (share of 2.04 GB sampled over 5 passes)

Cumulative shares include descendants and are **not additive**:

| Path | Share |
|---|---:|
| `decodeModule'` | 67.0% |
| — type table (`getFieldOptional' decodeTypeTable`) | **54.8%** |
| — `decodeExpr` / `decodeArray` | 24.6% / 23.6% |
| — `decodeAnn*` + usage (`StateT`) | ~15–20% |

Flat shares by generated file (objects):

| File | Share |
|---|---:|
| `gopurs_runtime/runtime.go` (`Apply`, `Any`, `Array`, `RecordDict`) | 19.6% |
| `Control_Monad_ST_Uncurried_ffi.go` | 8.2% |
| `Control_Monad_ST_Internal_ffi.go` | 5.3% |
| `PureScript…CoreFn_TypeTable.go` (decoder itself) | 5.2% |
| `Foreign_Object_ffi.go` | 4.8% |
| `Data_Either.go` / `Data_Maybe.go` (constructor accessors) | 4.8% / 4.1% |
| `Control_Monad_State_Trans.go` | 3.8% |
| `Data_Ord_ffi.go`, `Data_Array_ST_ffi.go` | 3.0% / 2.7% |
| `Record_Unsafe_ffi.go`, `Data_Int_ffi.go`, `Data_Map_Internal_ffi.go`, `Partial_Unsafe_ffi.go` | 2.1% / 1.4% / 1.4% / 1.3% |

Largest flat symbols: interface boxing `Any` 6.2%, `Right` constructor 4.3%,
`Just` constructor 4.1%, `Apply`/`Apply2` partial applications 6.8%,
uncurried ST adapter closures ~4.7%, the `Data.Ord` comparison wrapper 3.0%,
`unsafeGet` 2.1%.

### What the generated ST machinery looks like

Every `Control.Monad.ST.Internal` operation crosses the foreign bridge, which
re-creates `any`-based adapter closures and boxes each argument and result. For
`bind_`, the generated wrapper builds two closures and four interface boxes per
call before `f(a(nil))(nil)` runs. The same pattern appears for `forImpl`,
`while`, `foreach`, `read`/`write`, and for the uncurried `runSTFnN` and
`Data.Array.ST` helpers. This family is the single largest removable block in
the profile; the decoder's own loops contribute far less.

## Interpretation

1. **The type table is the dominant path**, but its cost is not algorithmic:
   2.47 attempts per type. It is the per-attempt constant (monadic steps,
   closures, `Maybe`/`Either` layers, `Value` boxing) that costs ~11 KB per
   type.
2. **Rewriting the decoder's PureScript source has a low ceiling**: the
   decoder's own generated code is ~5% of allocated objects. Removing
   intermediates in `resolveType`/`resolveArgs` would trim a fraction of that.
3. **The generic lever is the runtime/foreign interface**: ST adapters,
   `any`-boxing, partial application and constructor accessors are ~80% of
   allocated objects. This is the point 12 direction of the todo, now
   quantified on the current code.
4. The first bounded experiment should be the **ST/foreign-interface layer**
   (about 16% of objects plus the associated boxing and partial application),
   because it is generic, testable and independent of the decoder's
   semantics. The scratch probe confirmed it: see below.

## First experiment: Value-native ST wrappers (scratch probe)

`scratch/tast-decode-20260923/build-probe.py` copies the frozen generated
output, adds `gopurs_runtime` helpers that take `Value`s directly, and
replaces the bodies of the `Control.Monad.ST.Internal` wrappers (`bind_`,
`map_`, `pure_`, `run`, `while`, `forImpl`, `foreach`) and of the uncurried
`runSTFn1..9` wrappers with direct `Apply`/`ApplyN` calls. The adapters, their
`any` closures and the argument/result boxes disappear; call sites and the
generated decoder are untouched. `validate` mode checks all 12 module
fingerprints with the probe binary before any timing.

Alternating A/B, per-process report, corpus and oracles identical:

| Mode | Baseline | Probe | Delta |
|---|---:|---:|---:|
| Decode min (20 passes) | 385.91 ms | **334.12 ms** | **−13.42%** |
| Decode median | 418.84 ms | 371.69 ms | −11.27% |
| Decode min (allocation mode) | 557.51 ms | **490.28 ms** | −12.06% |
| Allocations per pass | 442,632,766 B | **402,378,712 B** | **−9.09% (−40.25 MB)** |

Three CPU pairs and two allocation pairs; the probe wins every pair and the
allocation figure is exact (`TotalAlloc`), not sampled. The oracle passes.

Caveats: this is a **scratch probe on the frozen Go**, not an integrated
compiler change; it patches only the wrapper bodies, so `Data.Array.ST`
operations still convert `[]Value` ↔ `[]any` per call and the PBO decoder is
unchanged. It measures a ceiling for the ST/foreign-interface layer, not a
promised compiler result. The next step is to obtain the same effect from
gopurs itself — codegen intrinsics for the ST operations, or a `Value`-native
foreign ABI — then run the paired campaign and the complete b8x build.

## Provenance

- Corpus sha256 `ef5ed1bae6ae52d084a4b3fd9d1e9e90a5da2c3a2d1e6e3aa56122007e0b223a`
  (same as the published campaign).
- Reference workspace: `altbak.pub-gopurs/var/benchmark/json-tast-ref-20260923`
  (`benchmark` sha256 `53361072…`, `benchmark.mjs` sha256 `8917cb3f…`).
- b8x reference: `scratch/tast-decode-20260923/reference` (`build.log`,
  `result.json`).
- Profile build: `scratch/tast-decode-20260923/profile/output`,
  `decode-profile` sha256 `cd195a3e…`; CPU profile `35ce078a…`,
  allocation profile `cf9323d3…`, live profile `1ccbd59b…`.
- ST probe: `probe-binary` sha256 `6ec79544…`, built by `build-probe.py`;
  raw A/B reports `probe-ab-cpu.json` and `probe-ab-alloc.json`.
- Reproduce:
  1. `python3 altbak.pub-gopurs/bin/benchmark/json-diagnostic.py build --suite JsonTypedAst --workspace <new>`
  2. `python3 … measure --suite JsonTypedAst --workspace <ws> --output <out>`
  3. `scratch/tast-decode-20260923/build-profile.sh` then `run-profile.sh cpu|alloc|validate|live`.

No repository file was modified for these measurements.
