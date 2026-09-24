# PHP scalar enum regions without a `Typed` annotation

## Mechanism

`Phpurs.EnumRegions` rewrites a local call graph into private workers and
private compact layouts only after proving that the region is closed, that its
inputs and result are scalar and that every reachable worker and layout is
known. The original candidate test recognized a region at

```purescript
Typed ty inner | scalar ty -> case peel inner of
  App _ _ -> ...
  UncurriedApp _ _ -> ...
```

The optimizer does not keep that annotation on every call. As the direct
argument of `EffectPure` the entry expression stays a bare application, so the
canonical numeric harness (`act = pure (depth (buildTree ...))`) opened no
region in `Test.RBTree` while the older string harness kept one because `show`
forced an `Int` annotation on the same call. The miss cost the 100k-insertion
red-black tree approximately half of its time and pushed the compiled column
total from the low 120s to the 260s of milliseconds.

The candidate test now consults the head worker's signature instead of the
annotation:

```purescript
candidateCall :: Context -> NeutralExpr -> Boolean
candidateCall ctx fn = case peel fn of
  Var (Qualified (Just mn) (Ident ident)) | mn == ctx.name ->
    case Map.lookup (Ident ident) ctx.bindings >>= signature of
      Just sig -> scalar sig.ret
      Nothing -> false
  _ -> false
```

`scan` attempts a region at `App` and `UncurriedApp` nodes whose head is a
local monomorphic worker returning a scalar. Everything else in the proof,
budgets and rewrite is unchanged, so annotations remain an optimization detail
rather than a correctness input.

## Validation

- `tests/codegen/enum-regions.mjs` gains two cases: an entry that is a bare
  application must open the same two-worker region, and a bare call returning
  the public ADT must stay refused. All nine codegen suites pass.
- With the fix, the complete generated PHP for the compiled mode differs only
  in `Test.RBTree/index.php`. The FFI modes' generated code is byte-identical
  to the September 20 build.
- Every measured process validates all fourteen numeric results and the batch
  calibration counts before a cell is accepted.

## Results

Campaign 2026-09-24 on the same host (Apple M4 Pro, PHP 8.5.4, Zend JIT 1255,
three processes, median per row, total sums the fourteen medians):

| column | 2026-09-20 | 2026-09-24 | /C |
|---|---:|---:|---:|
| compiled (`php-pure`) | 262.81 ms | **121.52 ms** | 12.3x |
| FP-style FFI (`php-ffi`) | 1057.50 ms | 1064.43 ms | 108.2x |
| imperative FFI (`php-fficc`) | 204.82 ms | 202.55 ms | 20.6x |

Per-process totals for the compiled column are 122.00, 121.47 and 120.80 ms.
The FFI columns moved within run-to-run variation while their generated PHP is
unchanged; they were remeasured so the whole PHP table shares one campaign and
one backend.

| row | 2026-09-20 | 2026-09-24 |
|---|---:|---:|
| AST Evaluation | 2.28 μs | 2.218241 μs |
| Fibonacci | 1.08 μs | 1.055331 μs |
| List Processing | 162.93 μs | 157.337234 μs |
| Tail Call Optimization | 74.79 μs | 74.868652 μs |
| Deep Record Updates | 1495.61 μs | 1514.411500 μs |
| Ackermann | 48.44 μs | 47.988281 μs |
| Church Numerals | 2114.47 μs | 2129.385375 μs |
| Prime Sieve | 467.09 μs | 765.263063 μs |
| Red-Black Tree | 250560.25 μs | **108946.584000 μs** |
| Polymorphism | 6484.65 μs | 6481.583500 μs |
| State Monad | 494.48 μs | 491.938781 μs |
| Lazy Evaluation | 802.46 μs | 808.585938 μs |
| Array Processing | 99.27 μs | 99.601891 μs |
| RowToList | 0.84 μs | 0.847051 μs |

Prime Sieve is the largest unrelated shift (467.09 → 765.26 μs) with identical
generated code, which is the known short-kernel variation described in the
methodology rather than a code-generation change. The red-black tree is the
only cell the fix moves, and it moves by 2.3x.

## Provenance and reproduction

- Fix: `phpurs/phpurs/src/Phpurs/EnumRegions.purs` (`candidateCall`);
  regression cases in `phpurs/phpurs/tests/codegen/enum-regions.mjs`.
- Campaigns: `var/benchmark/php-pure-20260924`, `var/benchmark/php-ffi-20260924`
  and `var/benchmark/php-fficc-20260924` retain the frozen plan, build
  manifests, source fingerprints, three raw processes per column and the
  validated `results.json`.
- Rebuild and measure one column with
  `python3 bin/php/campaign.py --mode pure` (or `--mode ffi` / `--mode fficc`)
  and republish it with
  `python3 bin/php/campaign.py --mode pure --campaign-dir <dir> --update-readme --publish-only`.
