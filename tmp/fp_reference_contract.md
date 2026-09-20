# Functional reference contract

The Haskell, Koka, OCaml and hand-written F# benchmarks translate the numeric
kernels in `src/Test/*.purs`. Matching the final integer is necessary but is not
sufficient: the source must express the same data structures and operations.
The C reference is deliberately imperative and is outside this contract.

| PureScript module | Input | Required source structure | Result |
| --- | --- | --- | --- |
| AstTree | 3 | Recursive expression construction and recursive evaluation of Val/Add/Mul/Sub. | 7 |
| Fib | 10 | The two recursive Fibonacci calls. | 55 |
| ListOps | 900 | Strict custom list range, accumulator-based even filter (reversed order), then generic fold. | 202950 |
| TCO | 100000 | Tail recursion adding `n mod 3`, with the zero base case. | 100000 |
| Records | 10000 | Recursive immutable updates of all three nested record levels. | 20000 |
| Ackermann | 3, 4 | The original recursive Ackermann equations. | 125 |
| Church | 10 | `fromInt`, `succC`, `mulC` and the `c100`/`c10k`/`c100k` function compositions. | 100000 |
| Primes | 500 | Strict linked-list sieve, capturing filter predicate, reverse, and sum. | 21536 |
| RBTree | 100000 | Persistent Okasaki tree, four balancing cases, descending insertions, then depth traversal. | 22 |
| Polymorphism | 10000000 | Generic Monoidish loop with `mempty = 1` and an integer addition instance/dictionary. | 10000000 |
| StateMonad | 20 | Pure State closures and result records, get/put/bind; rebuild the 60-depth computation with initial state zero inside each repetition. | 1200 |
| LazyEvaluation | 1000 | Explicit non-memoizing `unit -> value` closures; build and force a 1000-depth chain inside each repetition. | 1000000 |
| ArrayOps | 900 | Materialize a range array, filter into a result array, then fold; do not manually fuse the stages into a numeric loop. | 202950 |
| RowToList | opaque 10000 | Heterogeneous five-field record and type-indexed recursive keysNil/keysCons dictionary; do not write a literal result of 5. | 5 |

Normal compiler optimizations are allowed, including specialization, fusion,
constant folding, loop conversion, allocation reuse and lifting invariant
computations. The sources do not manually substitute arithmetic formulas or
cache the State/Lazy computations outside their repetition functions. A small
timing can therefore be a real compiler result even when the original source
contains closures. Harness barriers prevent sharing the result of an entire
timed call across the batch; they are not inserted into the algorithms to
disable normal optimization.

## Language mappings

- Haskell uses strict custom lists/records and explicit function wrappers for
  State/Lazy, rather than relying on implicit lazy evaluation. ArrayOps uses
  immutable `UArray` values; its adapters can allocate temporary lists.
- Koka uses native strict data and vectors. Its generic vector filter uses a
  temporary list because the installed vector library has no filter operation.
- OCaml uses native arrays. Its generic filter uses a private buffer and returns
  a fresh array, preserving the functional API without fusing filter and fold.
- F# uses its native array operations and typed immutable data. The four tree
  rotations use separate child matches to avoid a Fable nested-pattern bug.
- Haskell has a real Monoidish type class; Koka, OCaml and F# pass typed
  dictionaries explicitly. Typed heterogeneous rows and recursive dictionaries
  represent RowToList where the language has no PureScript row constraints.
- Native integer representations differ (machine integers in Haskell/OCaml,
  extensible integers in Koka, 32-bit integers in F#); all measured values are in
  the common range. These are comparisons of native compiler implementations,
  not identical memory layouts or runtimes.

## Reproduction and validation

```sh
python3 tmp/run_benchmarks.py --languages Haskell Koka OCaml --update-readme
python3 tmp/run_fable_native_benchmark.py --dotnet /path/to/dotnet10/dotnet --update-readme
```

Each launcher builds fresh outputs, checks all 14 expected values, then measures
three processes sequentially. Each cell is the median of the three per-process
minimum per-call batch times (10 calibrated batches, target at least 10 ms).
Source fingerprints, logs and unrounded results are retained under
`var/benchmark/native-references/` and `var/benchmark/fable-native-rust/`.
The first command leaves the C source and its published timings unchanged.

The earlier native FP results used numerical shortcuts in seven cases and are
superseded, not valid baselines for the corrected FP workloads. Their published
totals were Haskell 15.37 ms, Koka 32.20 ms, OCaml 12.33 ms and F# 57.46 ms. The
historical C reference (9.14 ms) and compiled Purust baseline (9.09 ms) retain
their separate roles; no speedup claim should be inferred from replacing the
old shortened workloads.
