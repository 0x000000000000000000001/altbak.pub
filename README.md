# 🌈 PureScript universal multi-runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. Backends consume either standard `CoreFn` or, for the local AOT backends, the typed `TAST/tcorefn` representation exported by a local PureScript fork (may be subject to an official PR soon).

## Comprehensive benchmarks
The benchmark suite runs a wide variety of computationally intensive tasks: AST evaluation, purely recursive Fibonacci, massive list processing, tail call optimization, deep record updates, Ackermann function, Church numerals, prime sieves, red black tree insertions, heavy polymorphism (type class dictionary lookups), State monad operations, deep lazy evaluation, heavy file I/O (10,000 synchronous writes and reads), and asynchronous `Aff` operations (via the native event loop). These tests apply massive pressure on the call stack, garbage collector, disk I/O, event loop, and runtime execution engine to measure the raw ability of the compiler and the underlying virtual machine. 

> [!IMPORTANT]
> **Worst-case scenario driven design**
> The code in these benchmarks is deliberately designed to be as naive, inefficient, and stressful as possible. The goal is to obtain the maximum possible performance gap between the compiled code and its native FFI equivalent. By studying these artificially worsened gaps, we can continually improve the code generation of our AOT compilers.
> 
> In the context of a real-world project, the idea is to provide the best possible performance ratio even when a developer makes a catastrophic design error. This mitigates the impact of such errors and delays as long as possible the need to drop down to FFI or use mutable abstractions for the rare algorithmic *hot paths* of a project.

### Core vs extended tests (`srx/`)
To ensure fair and executable comparisons across all backends, the test suite is split into two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) executed on the configured core backends via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Go optimization follow-up — 25 November 2026

A paired Go campaign, with benchmark sources and C/JS references unchanged, makes Polymorphism **2.26× faster** and RowToList **2.08× faster**. Across the 14 core cases, Go/C arithmetic mean falls from **5.045× to 2.611×**, geometric mean from **0.1356× to 0.1048×**; the total ratio and median remain practically stable. See the [results and four indicators](docs/benchmark-results/2026-11-25-go-small-actions.md) and [raw evidence](docs/benchmark-results/2026-11-25-go-small-actions.json). The earlier campaign and historical tables below retain their original measurements.

### Comparable Go, C and JavaScript programs — 25 November 2026

This fresh 14-case campaign uses the same PureScript sources for Go and official JavaScript, and a new [functional C translation](tmp/bench_functional.c) following the [source contract](tmp/fp_reference_contract.md). C allocation and release are included in each timed action. Compiler optimizations, including closed forms, remain allowed; these are complete-program comparisons with different runtime and allocation models.

All three backends were remeasured in three independent processes in rotating order; all 126 results and calibrated batches passed validation. Values below are medians of process minima, sorted by Go/C-functional before rounding. The [full report](docs/benchmark-results/2026-11-25-go-c-js-functional.md) gives both descending rankings, variability and reproducible [raw evidence](docs/benchmark-results/2026-11-25-go-c-js-functional.json).

| Case | Go (µs) | C functional (µs) | Official JS (µs) | Go / C functional | Go / JS |
| --- | ---: | ---: | ---: | ---: | ---: |
| Polymorphism | 0.032 | 0.00068 | 11494.62 | 47.41× | 0.0000028× |
| RowToList | 0.026 | 0.0017 | 0.036 | 15.90× | 0.74× |
| Tail Call Optimization | 60.19 | 15.07 | 185.99 | 3.99× | 0.32× |
| Fibonacci | 0.15 | 0.086 | 0.31 | 1.74× | 0.48× |
| Ackermann | 16.58 | 14.25 | 63.84 | 1.16× | 0.26× |
| State Monad | 92.85 | 95.59 | 46.62 | 0.97× | 1.99× |
| Red-Black Tree | 9472.17 | 53161.00 | 82980.62 | 0.18× | 0.11× |
| Church Numerals | 223.43 | 1879.75 | 1362.31 | 0.12× | 0.16× |
| Array Processing | 0.028 | 0.33 | 6.99 | 0.084× | 0.0040× |
| AST Evaluation | 0.051 | 0.73 | 0.26 | 0.071× | 0.20× |
| Deep Record Updates | 0.047 | 11.49 | 72.88 | 0.0041× | 0.00065× |
| Prime Sieve | 0.57 | 158.39 | 58.25 | 0.0036× | 0.0097× |
| List Processing | 0.028 | 21.02 | 8.70 | 0.0013× | 0.0032× |
| Lazy Evaluation | 0.76 | 20301.00 | 13813.88 | 0.000037× | 0.000055× |

Polymorphism and RowToList have very short optimized C kernels: their ratios are sensitive to harness and call overhead, not a cost per dictionary operation. Small Go ratios on other cases can reflect legitimate elimination of source-level work. State Go/JS is 1.99× for the ratio of medians, with 1.87–2.07× observed across cycles.

The historical multi-backend tables below retain their original timings and **imperative C reference**. Their `/C` ratios are distinct from the functional C comparison above; the historical timing catalogue is preserved.

### Core stresstest benchmark results (pure computational)

> [!NOTE]
> **Reading the benchmark tables**
>
> - **Groups:** Stable = 14 core tests; Incl WIP = 23 shared micro-benchmarks. Together they make 37 tests. Excl WIP contains three diagnostics, excluded from all totals.
> - **Ratios:** `/C` and `/JS` compare with the first C and JavaScript reference columns. **Below 1× is faster; above 1× is slower.** Extended Results uses its own JS reference and has no C reference.
> - **Totals:** sums of per-test median times. Total ratios compare these sums; they do not represent an application benchmark.
> - **Summaries:** `μ` = arithmetic mean, `Q₂` = median, `G` = geometric mean of the case ratios. Each case has equal weight. Very small reference times can produce large ratios.
> - **Missing data:** `—` means unavailable or not applicable. Bounds (`< …`) are excluded from summaries; partial coverage is identified. The 37-test total requires both Stable and Incl WIP.
> - **Display precision, in every table:** values ≥ 1 use two decimal places (`12.000045 → 12.00`); positive values < 1 keep two significant digits after leading zeros (`0.000045`, `0.009492 → 0.0095`). Zero is `0.00`.
> - **Calculations:** use archived timings before display rounding, including for threshold checks. `≈` marks a fallback to published precision. See the [timing archive](bin/benchmark/ratio-measurements.json) and [methodology](docs/benchmark-methodology.md).

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
| **Stable** | | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.26 μs | ~ 0.19 μs | ~ 0.20 μs | ~ 0.22 μs |
| ↳ /C | 2.89x | 2.09x | 2.17x | 2.43x |
| ↳ /JS | — | 0.72x | 0.75x | 0.84x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs | ~ 0.31 μs | ~ 0.31 μs | ~ 0.37 μs |
| ↳ /C | 4.01x | 4.03x | 4.08x | 4.81x |
| ↳ /JS | — | 1.00x | 1.02x | 1.20x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 9.59 μs | ~ 5.06 μs | ~ 19.06 μs | ~ 0.49 μs |
| ↳ /C | 207.46x | 109.43x | 412.28x | 10.68x |
| ↳ /JS | — | 0.53x | 1.99x | 0.051x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 189.49 μs | ~ 151.13 μs | ~ 65.91 μs | ~ 68.31 μs |
| ↳ /C | 5.90x | 4.70x | 2.05x | 2.13x |
| ↳ /JS | — | 0.80x | 0.35x | 0.36x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 74.18 μs | ~ 134.11 μs | ~ 42.07 μs | ~ 7.69 μs |
| ↳ /C | 22.80x | 41.23x | 12.93x | 2.36x |
| ↳ /JS | — | 1.81x | 0.57x | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 67.49 μs | ~ 64.40 μs | ~ 38.25 μs | ~ 38.48 μs |
| ↳ /C | 4.01x | 3.83x | 2.27x | 2.29x |
| ↳ /JS | — | 0.95x | 0.57x | 0.57x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1263.82 μs | ~ 1242.29 μs | ~ 1337.51 μs | ~ 49.63 μs |
| ↳ /C | 1864037.79x | 1832288.62x | 1972729.17x | 73200.54x |
| ↳ /JS | — | 0.98x | 1.06x | 0.039x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 63.39 μs | ~ 26.99 μs | ~ 86.33 μs | ~ 0.65 μs |
| ↳ /C | 63.04x | 26.84x | 85.85x | 0.65x |
| ↳ /JS | — | 0.43x | 1.36x | 0.010x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 89940.58 μs | ~ 47272.88 μs | ~ 24486.83 μs | ~ 27249.58 μs |
| ↳ /C | 9.93x | 5.22x | 2.70x | 3.01x |
| ↳ /JS | — | 0.53x | 0.27x | 0.30x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 22830.25 μs | ~ 7368.38 μs | ~ 4843.00 μs | ~ 2805.25 μs |
| ↳ /C | 34075000.00x | 10997574.63x | 7228358.21x | 4186940.30x |
| ↳ /JS | — | 0.32x | 0.21x | 0.12x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 48.23 μs | ~ 10.95 μs | ~ 46.87 μs | ~ 0.42 μs |
| ↳ /C | 72201.55x | 16387.33x | 70165.82x | 625.43x |
| ↳ /JS | — | 0.23x | 0.97x | 0.0087x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 14608.29 μs | ~ 10604.88 μs | ~ 17513.12 μs | ~ 321.48 μs |
| ↳ /C | 21868700.60x | 15875561.38x | 26217252.99x | 481254.30x |
| ↳ /JS | — | 0.73x | 1.20x | 0.022x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.35 μs | ~ 3.75 μs | ~ 3.12 μs | ~ 0.50 μs |
| ↳ /C | 161.24x | 82.28x | 68.47x | 10.93x |
| ↳ /JS | — | 0.51x | 0.42x | 0.068x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.042 μs | ~ 0.017 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 62.23x | 25.67x | 89.79x | 30.63x |
| ↳ /JS | — | 0.41x | 1.44x | 0.49x |
| **Total Stable** | <strong>~ 129.10 ms</strong> | <strong>~ 66.89 ms</strong> | <strong>~ 48.48 ms</strong> | <strong>~ 30.54 ms</strong> |
| ↳ /C | <strong>14.18x</strong> | <strong>7.34x</strong> | <strong>5.32x</strong> | <strong>3.35x</strong> |
| ↳ μ/C | <strong>4134320.25x</strong> | <strong>2051579.81x</strong> | <strong>2534942.06x</strong> | <strong>338720.75x</strong> |
| ↳ Q₂/C | <strong>62.64x</strong> | <strong>34.03x</strong> | <strong>77.16x</strong> | <strong>7.75x</strong> |
| ↳ G/C | <strong>601.80x</strong> | <strong>374.38x</strong> | <strong>433.42x</strong> | <strong>71.23x</strong> |
| ↳ /JS | — | <strong>0.52x</strong> | <strong>0.38x</strong> | <strong>0.24x</strong> |
| ↳ μ/JS | — | <strong>0.71x</strong> | <strong>0.87x</strong> | <strong>0.30x</strong> |
| ↳ Q₂/JS | — | <strong>0.63x</strong> | <strong>0.86x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | — | <strong>0.62x</strong> | <strong>0.72x</strong> | <strong>0.12x</strong> |
| **Incl WIP** | | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 121.94 μs | ~ 144.60 μs | ~ 92.20 μs | ~ 12.93 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.19x | 0.76x | 0.11x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 127.59 μs | ~ 130.46 μs | ~ 74.63 μs | ~ 18.58 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.02x | 0.58x | 0.15x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 70.25 μs | ~ 70.22 μs | ~ 44.01 μs | ~ 3.32 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.00x | 0.63x | 0.047x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 561.69 μs | ~ 517.87 μs | ~ 114.68 μs | ~ 78.30 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.92x | 0.20x | 0.14x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 51.63 μs | ~ 50.71 μs | ~ 41.04 μs | ~ 5.03 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.98x | 0.79x | 0.097x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.20 μs | ~ 31.05 μs | ~ 2.05 μs | ~ 2.40 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.03x | 0.068x | 0.079x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 80.61 μs | ~ 81.04 μs | ~ 49.83 μs | ~ 1.39 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.01x | 0.62x | 0.017x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 68.35 μs | ~ 65.14 μs | ~ 42.18 μs | ~ 1.07 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 68.01 μs | ~ 64.65 μs | ~ 42.45 μs | ~ 1.08 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.010 μs | ~ 0.0087 μs | ~ 0.019 μs | ~ 0.020 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.86x | 1.92x | 1.94x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 37.44 μs | ~ 37.44 μs | ~ 37.09 μs | ~ 8.54 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.00x | 0.99x | 0.23x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.99x | 1.02x | 1.00x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 63.80 μs | ~ 63.97 μs | ~ 63.94 μs | ~ 63.72 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.00x | 1.00x | 1.00x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 45.28 μs | ~ 44.96 μs | ~ 42.10 μs | ~ 21.01 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.99x | 0.93x | 0.46x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 145.33 μs | ~ 144.87 μs | ~ 138.40 μs | ~ 24.69 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.00x | 0.95x | 0.17x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 47.71 μs | ~ 46.52 μs | ~ 42.64 μs | ~ 11.85 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.97x | 0.89x | 0.25x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 2.44 μs | ~ 4.37 μs | ~ 4.32 μs | ~ 0.72 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.79x | 1.77x | 0.30x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 4.55 μs | ~ 9.34 μs | ~ 9.04 μs | ~ 0.71 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 2.05x | 1.98x | 0.16x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 7.03 μs | ~ 14.73 μs | ~ 14.38 μs | ~ 0.72 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 2.10x | 2.05x | 0.10x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 463.87 μs | ~ 106.02 μs | ~ 2.44 μs | ~ 2.48 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.23x | 0.0053x | 0.0053x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 513.20 μs | ~ 106.03 μs | ~ 31.67 μs | ~ 3.41 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.21x | 0.062x | 0.0066x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 155.67 μs | ~ 102.38 μs | ~ 2.39 μs | ~ 2.56 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.66x | 0.015x | 0.016x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 211.84 μs | ~ 113.53 μs | ~ 31.14 μs | ~ 3.52 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.54x | 0.15x | 0.017x |
| **Total Incl WIP** | <strong>~ 2.88 ms</strong> | <strong>~ 1.95 ms</strong> | <strong>~ 0.92 ms</strong> | <strong>~ 0.27 ms</strong> |
| ↳ /C | — | — | — | — |
| ↳ μ/C | — | — | — | — |
| ↳ Q₂/C | — | — | — | — |
| ↳ G/C | — | — | — | — |
| ↳ /JS | — | <strong>0.68x</strong> | <strong>0.32x</strong> | <strong>0.093x</strong> |
| ↳ μ/JS | — | <strong>1.02x</strong> | <strong>0.81x</strong> | <strong>0.27x</strong> |
| ↳ Q₂/JS | — | <strong>0.99x</strong> | <strong>0.76x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | — | <strong>0.91x</strong> | <strong>0.44x</strong> | <strong>0.092x</strong> |
| **Total Stable + Incl WIP** | <strong>~ 131.98 ms</strong> | <strong>~ 68.84 ms</strong> | <strong>~ 49.41 ms</strong> | <strong>~ 30.81 ms</strong> |
| ↳ /C | — | — | — | — |
| ↳ μ/C | — | — | — | — |
| ↳ Q₂/C | — | — | — | — |
| ↳ G/C | — | — | — | — |
| ↳ /JS | — | <strong>0.52x</strong> | <strong>0.37x</strong> | <strong>0.23x</strong> |
| ↳ μ/JS | — | <strong>0.90x</strong> | <strong>0.83x</strong> | <strong>0.28x</strong> |
| ↳ Q₂/JS | — | <strong>0.95x</strong> | <strong>0.76x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | — | <strong>0.79x</strong> | <strong>0.53x</strong> | <strong>0.10x</strong> |
| **Excl WIP** | | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 25303.33 μs | — | — | — |
| ↳ /C | 3.84x | — | — | — |
| ↳ /JS | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 80646.58 μs | — | — | — |
| ↳ /C | 7.64x | — | — | — |
| ↳ /JS | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 9283.94 μs | — | — | — |
| ↳ /C | 14.12x | — | — | — |
| ↳ /JS | — | — | — | — |

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs) <br>`--dev`<br>↓<br>Go<br>↓<br>go-build<br>`-N -l`<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) |
----------------------- | -------------------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
| **Stable** | | | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.053 μs | ~ 0.52 μs | ~ 4.17 μs | ~ 0.32 μs | ~ 0.32 μs |
| ↳ /C | 0.59x | 5.78x | 46.41x | 3.60x | 3.54x |
| ↳ /JS | 0.21x | 2.00x | 16.07x | 1.25x | 1.22x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs | ~ 0.54 μs | ~ 5.41 μs | ~ 0.16 μs | ~ 0.16 μs |
| ↳ /C | 2.00x | 7.00x | 70.07x | 2.02x | 2.13x |
| ↳ /JS | 0.50x | 1.74x | 17.46x | 0.50x | 0.53x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.027 μs | ~ 28.86 μs | ~ 236.01 μs | ~ 24.39 μs | ~ 0.45 μs |
| ↳ /C | 0.59x | 624.42x | 5106.27x | 527.64x | 9.72x |
| ↳ /JS | 0.0028x | 3.01x | 24.61x | 2.54x | 0.047x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 63.20 μs | ~ 355.46 μs | ~ 6281.96 μs | ~ 35.98 μs | ~ 35.50 μs |
| ↳ /C | 1.97x | 11.06x | 195.45x | 1.12x | 1.10x |
| ↳ /JS | 0.33x | 1.88x | 33.15x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 0.046 μs | ~ 186.36 μs | ~ 5461.52 μs | ~ 40.43 μs | ~ 14.87 μs |
| ↳ /C | 0.014x | 57.29x | 1678.95x | 12.43x | 4.57x |
| ↳ /JS | 0.00062x | 2.51x | 73.62x | 0.55x | 0.20x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.66 μs | ~ 35.56 μs | ~ 402.04 μs | ~ 20.85 μs | ~ 19.87 μs |
| ↳ /C | 0.99x | 2.11x | 23.89x | 1.24x | 1.18x |
| ↳ /JS | 0.25x | 0.53x | 5.96x | 0.31x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 232.04 μs | ~ 1296.16 μs | ~ 3811.48 μs | ~ 1044.60 μs | ~ 22.45 μs |
| ↳ /C | 342238.59x | 1911740.41x | 5621650.44x | 1540714.23x | 33109.26x |
| ↳ /JS | 0.18x | 1.03x | 3.02x | 0.83x | 0.018x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 0.59 μs | ~ 76.41 μs | ~ 1340.99 μs | ~ 110.87 μs | ~ 0.62 μs |
| ↳ /C | 0.59x | 75.99x | 1333.67x | 110.27x | 0.62x |
| ↳ /JS | 0.0094x | 1.21x | 21.16x | 1.75x | 0.0098x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10102.83 μs | ~ 31239.75 μs | ~ 870404.62 μs | ~ 24630.67 μs | ~ 9004.48 μs |
| ↳ /C | 1.12x | 3.45x | 96.14x | 2.72x | 0.99x |
| ↳ /JS | 0.11x | 0.35x | 9.68x | 0.27x | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.031 μs | ~ 15617.54 μs | ~ 505069.00 μs | ~ 49160.92 μs | ~ 2244.45 μs |
| ↳ /C | 46.59x | 23309761.19x | 753834328.36x | 73374501.49x | 3349930.04x |
| ↳ /JS | 0.0000014x | 0.68x | 22.12x | 2.15x | 0.098x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 91.18 μs | ~ 110.57 μs | ~ 303.56 μs | ~ 52.10 μs | ~ 0.33 μs |
| ↳ /C | 136491.67x | 165523.95x | 454429.03x | 78000.25x | 499.53x |
| ↳ /JS | 1.89x | 2.29x | 6.29x | 1.08x | 0.0069x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.77 μs | ~ 1566.95 μs | ~ 71378.08 μs | ~ 13865.38 μs | ~ 0.27 μs |
| ↳ /C | 1150.33x | 2345733.53x | 106853417.66x | 20756549.40x | 402.15x |
| ↳ /JS | 0.000053x | 0.11x | 4.89x | 0.95x | 0.000018x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.027 μs | ~ 32.70 μs | ~ 43.81 μs | ~ 4.90 μs | ~ 0.45 μs |
| ↳ /C | 0.60x | 717.28x | 961.02x | 107.55x | 9.90x |
| ↳ /JS | 0.0037x | 4.45x | 5.96x | 0.67x | 0.061x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.026 μs | ~ 0.30 μs | ~ 0.50 μs | ~ 0.073 μs | ~ 0.032 μs |
| ↳ /C | 39.29x | 449.10x | 754.60x | 108.58x | 47.84x |
| ↳ /JS | 0.63x | 7.22x | 12.13x | 1.74x | 0.77x |
| **Total Stable** | <strong>~ 10.51 ms</strong> | <strong>~ 50.55 ms</strong> | <strong>~ 1464.74 ms</strong> | <strong>~ 88.99 ms</strong> | <strong>~ 11.34 ms</strong> |
| ↳ /C | <strong>1.15x</strong> | <strong>5.55x</strong> | <strong>160.84x</strong> | <strong>9.77x</strong> | <strong>1.25x</strong> |
| ↳ μ/C | <strong>34283.92x</strong> | <strong>1981050.90x</strong> | <strong>61912435.14x</strong> | <strong>6839331.61x</strong> | <strong>241715.90x</strong> |
| ↳ Q₂/C | <strong>1.54x</strong> | <strong>262.55x</strong> | <strong>1147.34x</strong> | <strong>108.06x</strong> | <strong>7.15x</strong> |
| ↳ G/C | <strong>11.54x</strong> | <strong>806.72x</strong> | <strong>7681.72x</strong> | <strong>487.21x</strong> | <strong>34.08x</strong> |
| ↳ /JS | <strong>0.081x</strong> | <strong>0.39x</strong> | <strong>11.35x</strong> | <strong>0.69x</strong> | <strong>0.088x</strong> |
| ↳ μ/JS | <strong>0.29x</strong> | <strong>2.07x</strong> | <strong>18.29x</strong> | <strong>1.06x</strong> | <strong>0.25x</strong> |
| ↳ Q₂/JS | <strong>0.15x</strong> | <strong>1.81x</strong> | <strong>14.10x</strong> | <strong>0.89x</strong> | <strong>0.099x</strong> |
| ↳ G/JS | <strong>0.019x</strong> | <strong>1.34x</strong> | <strong>12.76x</strong> | <strong>0.81x</strong> | <strong>0.057x</strong> |
| **Incl WIP** | | | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 132.22 μs | ~ 1452.41 μs | ~ 2146.41 μs | ~ 77.41 μs | ~ 6.61 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.08x | 11.91x | 17.60x | 0.63x | 0.054x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 77.25 μs | ~ 828.92 μs | ~ 1330.39 μs | ~ 67.38 μs | ~ 6.85 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.61x | 6.50x | 10.43x | 0.53x | 0.054x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 4.85 μs | ~ 281.15 μs | ~ 133.05 μs | ~ 49.77 μs | ~ 2.31 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.069x | 4.00x | 1.89x | 0.71x | 0.033x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 159.57 μs | ~ 12300.38 μs | ~ 11626.08 μs | ~ 825.35 μs | ~ 10.36 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.28x | 21.90x | 20.70x | 1.47x | 0.018x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 3.85 μs | ~ 752.15 μs | ~ 98.90 μs | ~ 48.10 μs | ~ 2.53 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.075x | 14.57x | 1.92x | 0.93x | 0.049x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 26.69 μs | ~ 304.51 μs | ~ 372.03 μs | ~ 13.95 μs | ~ 1.25 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.88x | 10.08x | 12.32x | 0.46x | 0.042x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 16.24 μs | ~ 464.65 μs | ~ 476.59 μs | ~ 18.97 μs | ~ 0.87 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.20x | 5.76x | 5.91x | 0.24x | 0.011x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 3.86 μs | ~ 298.48 μs | ~ 303.08 μs | ~ 11.64 μs | ~ 0.74 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.056x | 4.37x | 4.43x | 0.17x | 0.011x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 3.78 μs | ~ 297.10 μs | ~ 307.26 μs | ~ 11.44 μs | ~ 0.74 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.056x | 4.37x | 4.52x | 0.17x | 0.011x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.0097 μs | ~ 17.05 μs | ~ 3.47 μs | ~ 3.45 μs | ~ 3.45 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.96x | 1687.77x | 343.62x | 341.43x | 341.50x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 6.70 μs | ~ 79.96 μs | ~ 62.44 μs | ~ 7.20 μs | ~ 3.41 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.18x | 2.14x | 1.67x | 0.19x | 0.091x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.33 μs | ~ 130.75 μs | ~ 2.71 μs | ~ 0.17 μs | ~ 0.17 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.71x | 685.42x | 14.19x | 0.90x | 0.88x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 2.47 μs | ~ 140.42 μs | ~ 2.47 μs | ~ 2.31 μs | ~ 2.32 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.039x | 2.20x | 0.039x | 0.036x | 0.036x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 45.35 μs | ~ 105.92 μs | ~ 104.16 μs | ~ 146.89 μs | ~ 3.66 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.00x | 2.34x | 2.30x | 3.24x | 0.081x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 138.52 μs | ~ 318.05 μs | ~ 300.93 μs | ~ 32.68 μs | ~ 9.75 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.95x | 2.19x | 2.07x | 0.22x | 0.067x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 58.46 μs | ~ 216.07 μs | ~ 167.98 μs | ~ 74.43 μs | ~ 3.71 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.23x | 4.53x | 3.52x | 1.56x | 0.078x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.042 μs | ~ 5.46 μs | ~ 144.74 μs | ~ 0.72 μs | ~ 0.24 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.017x | 2.24x | 59.23x | 0.29x | 0.10x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 0.14 μs | ~ 13.30 μs | ~ 190.03 μs | ~ 7.20 μs | ~ 0.26 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.031x | 2.92x | 41.74x | 1.58x | 0.057x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 0.40 μs | ~ 22.40 μs | ~ 679.42 μs | ~ 12.85 μs | ~ 1.03 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.057x | 3.19x | 96.67x | 1.83x | 0.15x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 6.26 μs | ~ 1019.37 μs | ~ 3085.81 μs | ~ 2.25 μs | ~ 2.25 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.013x | 2.20x | 6.65x | 0.0049x | 0.0049x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 20.77 μs | ~ 1073.35 μs | ~ 4068.22 μs | ~ 2.33 μs | ~ 2.38 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.040x | 2.09x | 7.93x | 0.0045x | 0.0046x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 201.66 μs | ~ 2199.81 μs | ~ 3135.15 μs | ~ 2.26 μs | ~ 2.26 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.30x | 14.13x | 20.14x | 0.014x | 0.014x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 182.88 μs | ~ 4998.00 μs | ~ 4026.20 μs | ~ 2.36 μs | ~ 2.36 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.86x | 23.59x | 19.01x | 0.011x | 0.011x |
| **Total Incl WIP** | <strong>~ 1.09 ms</strong> | <strong>~ 27.32 ms</strong> | <strong>~ 32.77 ms</strong> | <strong>~ 1.42 ms</strong> | <strong>~ 0.070 ms</strong> |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | <strong>0.38x</strong> | <strong>9.49x</strong> | <strong>11.38x</strong> | <strong>0.49x</strong> | <strong>0.024x</strong> |
| ↳ μ/JS | <strong>0.51x</strong> | <strong>109.58x</strong> | <strong>30.37x</strong> | <strong>15.51x</strong> | <strong>14.93x</strong> |
| ↳ Q₂/JS | <strong>0.20x</strong> | <strong>4.37x</strong> | <strong>7.93x</strong> | <strong>0.46x</strong> | <strong>0.049x</strong> |
| ↳ G/JS | <strong>0.20x</strong> | <strong>7.90x</strong> | <strong>7.95x</strong> | <strong>0.33x</strong> | <strong>0.054x</strong> |
| **Total Stable + Incl WIP** | <strong>~ 11.60 ms</strong> | <strong>~ 77.87 ms</strong> | <strong>~ 1497.51 ms</strong> | <strong>~ 90.41 ms</strong> | <strong>~ 11.41 ms</strong> |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | <strong>0.088x</strong> | <strong>0.59x</strong> | <strong>11.35x</strong> | <strong>0.69x</strong> | <strong>0.086x</strong> |
| ↳ μ/JS | <strong>0.43x</strong> | <strong>68.90x</strong> | <strong>25.80x</strong> | <strong>10.04x</strong> | <strong>9.38x</strong> |
| ↳ Q₂/JS | <strong>0.18x</strong> | <strong>2.92x</strong> | <strong>10.43x</strong> | <strong>0.63x</strong> | <strong>0.054x</strong> |
| ↳ G/JS | <strong>0.083x</strong> | <strong>4.04x</strong> | <strong>9.51x</strong> | <strong>0.46x</strong> | <strong>0.055x</strong> |
| **Excl WIP** | | | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 4616.79 μs | — | — | — | — |
| ↳ /C | 0.70x | — | — | — | — |
| ↳ /JS | 0.18x | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 20801.17 μs | — | — | — | — |
| ↳ /C | 1.97x | — | — | — | — |
| ↳ /JS | 0.26x | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 1403.38 μs | — | — | — | — |
| ↳ /C | 2.13x | — | — | — | — |
| ↳ /JS | 0.15x | — | — | — | — |

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.074 μs | ~ 0.076 μs | ~ 0.069 μs |
| ↳ /C | 0.82x | 0.84x | 0.77x |
| ↳ /JS | 0.28x | 0.29x | 0.27x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs | ~ 0.15 μs | ~ 0.14 μs |
| ↳ /C | 1.91x | 1.91x | 1.87x |
| ↳ /JS | 0.48x | 0.48x | 0.47x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 6.35 μs | ~ 3.76 μs | ~ 2.84 μs |
| ↳ /C | 137.43x | 81.33x | 61.35x |
| ↳ /JS | 0.66x | 0.39x | 0.30x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 277.12 μs | ~ 322.34 μs | ~ 329.88 μs |
| ↳ /C | 8.62x | 10.03x | 10.26x |
| ↳ /JS | 1.46x | 1.70x | 1.74x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 228.44 μs | ~ 48.07 μs | ~ 43.57 μs |
| ↳ /C | 70.23x | 14.78x | 13.40x |
| ↳ /JS | 3.08x | 0.65x | 0.59x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.99 μs | ~ 9.07 μs | ~ 9.14 μs |
| ↳ /C | 1.01x | 0.54x | 0.54x |
| ↳ /JS | 0.25x | 0.13x | 0.14x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 237.12 μs | ~ 235.53 μs | ~ 62.26 μs |
| ↳ /C | 349741.89x | 347391.22x | 91825.68x |
| ↳ /JS | 0.19x | 0.19x | 0.049x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 37.49 μs | ~ 22.68 μs | ~ 0.80 μs |
| ↳ /C | 37.29x | 22.56x | 0.79x |
| ↳ /JS | 0.59x | 0.36x | 0.013x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 22551.00 μs | ~ 15428.00 μs | ~ 11970.00 μs |
| ↳ /C | 2.49x | 1.70x | 1.32x |
| ↳ /JS | 0.25x | 0.17x | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 17812.00 μs | ~ 27373.00 μs | ~ 6730.50 μs |
| ↳ /C | 26585074.63x | 40855223.88x | 10045522.39x |
| ↳ /JS | 0.78x | 1.20x | 0.29x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 1.70 μs | ~ 2.02 μs | ~ 0.84 μs |
| ↳ /C | 2545.93x | 3030.19x | 1257.80x |
| ↳ /JS | 0.035x | 0.042x | 0.017x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 2989.50 μs | ~ 1964.00 μs | ~ 0.61 μs |
| ↳ /C | 4475299.40x | 2940119.76x | 912.79x |
| ↳ /JS | 0.20x | 0.13x | 0.000042x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.29 μs | ~ 5.04 μs | ~ 2.79 μs |
| ↳ /C | 159.94x | 110.59x | 61.27x |
| ↳ /JS | 0.99x | 0.69x | 0.38x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs | ~ 0.015 μs | ~ 0.0010 μs |
| ↳ /C | 1.49x | 21.82x | 1.50x |
| ↳ /JS | 0.024x | 0.35x | 0.024x |
| **Total Stable** | <strong>~ 44.17 ms</strong> | <strong>~ 45.41 ms</strong> | <strong>~ 19.15 ms</strong> |
| ↳ /C | <strong>4.85x</strong> | <strong>4.99x</strong> | <strong>2.10x</strong> |
| ↳ μ/C | <strong>2243791.65x</strong> | <strong>3153287.94x</strong> | <strong>724262.27x</strong> |
| ↳ Q₂/C | <strong>53.76x</strong> | <strong>22.19x</strong> | <strong>11.83x</strong> |
| ↳ G/C | <strong>209.29x</strong> | <strong>195.91x</strong> | <strong>50.33x</strong> |
| ↳ /JS | <strong>0.34x</strong> | <strong>0.35x</strong> | <strong>0.15x</strong> |
| ↳ μ/JS | <strong>0.66x</strong> | <strong>0.48x</strong> | <strong>0.31x</strong> |
| ↳ Q₂/JS | <strong>0.38x</strong> | <strong>0.35x</strong> | <strong>0.20x</strong> |
| ↳ G/JS | <strong>0.35x</strong> | <strong>0.33x</strong> | <strong>0.084x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.11 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.23x | 1.27x | 1.12x |
| ↳ /JS | 0.43x | 0.44x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.21 μs | ~ 0.23 μs | ~ 0.24 μs |
| ↳ /C | 2.71x | 3.00x | 3.14x |
| ↳ /JS | 0.67x | 0.75x | 0.78x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 37.32 μs | ~ 5.22 μs | ~ 0.86 μs |
| ↳ /C | 807.44x | 112.98x | 18.58x |
| ↳ /JS | 3.89x | 0.54x | 0.090x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 1341.30 μs | ~ 103.95 μs | ~ 104.62 μs |
| ↳ /C | 41.73x | 3.23x | 3.26x |
| ↳ /JS | 7.08x | 0.55x | 0.55x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 404.77 μs | ~ 1276.40 μs | ~ 15.46 μs |
| ↳ /C | 124.43x | 392.38x | 4.75x |
| ↳ /JS | 5.46x | 17.21x | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.24 μs | ~ 18.81 μs | ~ 19.68 μs |
| ↳ /C | 1.14x | 1.12x | 1.17x |
| ↳ /JS | 0.29x | 0.28x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 520.08 μs | ~ 418.19 μs | ~ 0.012 μs |
| ↳ /C | 767080.66x | 616801.48x | 17.57x |
| ↳ /JS | 0.41x | 0.33x | 0.0000094x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 148.04 μs | ~ 30.95 μs | ~ 36.68 μs |
| ↳ /C | 147.23x | 30.78x | 36.48x |
| ↳ /JS | 2.34x | 0.49x | 0.58x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16911.71 μs | ~ 14282.54 μs | ~ 14426.50 μs |
| ↳ /C | 1.87x | 1.58x | 1.59x |
| ↳ /JS | 0.19x | 0.16x | 0.16x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 66666.38 μs | ~ 52475.29 μs | ~ 0.0097 μs |
| ↳ /C | 99502052.24x | 78321329.85x | 14.51x |
| ↳ /JS | 2.92x | 2.30x | 0.00000043x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 20.41 μs | ~ 14.55 μs | ~ 7.61 μs |
| ↳ /C | 30554.61x | 21777.31x | 11398.44x |
| ↳ /JS | 0.42x | 0.30x | 0.16x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8739.06 μs | ~ 6780.02 μs | ~ 0.010 μs |
| ↳ /C | 13082428.89x | 10149732.04x | 15.01x |
| ↳ /JS | 0.60x | 0.46x | 0.00000069x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 27.82 μs | ~ 16.21 μs | ~ 1.04 μs |
| ↳ /C | 610.34x | 355.50x | 22.89x |
| ↳ /JS | 3.79x | 2.20x | 0.14x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.040 μs | ~ 0.044 μs | ~ 0.0095 μs |
| ↳ /C | 59.58x | 65.94x | 14.25x |
| ↳ /JS | 0.96x | 1.06x | 0.23x |
| **Total Stable** | <strong>~ 94.84 ms</strong> | <strong>~ 75.42 ms</strong> | <strong>~ 14.61 ms</strong> |
| ↳ /C | <strong>10.41x</strong> | <strong>8.28x</strong> | <strong>1.60x</strong> |
| ↳ μ/C | <strong>8098851.01x</strong> | <strong>6365043.46x</strong> | <strong>825.20x</strong> |
| ↳ Q₂/C | <strong>135.83x</strong> | <strong>89.46x</strong> | <strong>14.38x</strong> |
| ↳ G/C | <strong>683.68x</strong> | <strong>430.10x</strong> | <strong>11.77x</strong> |
| ↳ /JS | <strong>0.73x</strong> | <strong>0.58x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>2.10x</strong> | <strong>1.93x</strong> | <strong>0.26x</strong> |
| ↳ Q₂/JS | <strong>0.82x</strong> | <strong>0.52x</strong> | <strong>0.18x</strong> |
| ↳ G/JS | <strong>1.14x</strong> | <strong>0.71x</strong> | <strong>0.020x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.13 μs | ~ 3.19 μs | ~ 2.11 μs |
| ↳ /C | 23.68x | 35.47x | 23.46x |
| ↳ /JS | 8.20x | 12.28x | 8.12x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 1.08 μs | ~ 2.26 μs | ~ 1.27 μs |
| ↳ /C | 14.00x | 29.29x | 16.46x |
| ↳ /JS | 3.49x | 7.30x | 4.10x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 137.53 μs | ~ 174.47 μs | ~ 0.63 μs |
| ↳ /C | 2975.62x | 3774.85x | 13.63x |
| ↳ /JS | 14.34x | 18.20x | 0.066x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 72.18 μs | ~ 49.87 μs | ~ 50.30 μs |
| ↳ /C | 2.25x | 1.55x | 1.56x |
| ↳ /JS | 0.38x | 0.26x | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1413.71 μs | ~ 1676.49 μs | ~ 491.74 μs |
| ↳ /C | 434.60x | 515.38x | 151.17x |
| ↳ /JS | 19.06x | 22.60x | 6.63x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 45.89 μs | ~ 100.16 μs | ~ 52.48 μs |
| ↳ /C | 2.73x | 5.95x | 3.12x |
| ↳ /JS | 0.68x | 1.48x | 0.78x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1941.41 μs | ~ 6824.38 μs | ~ 32.84 μs |
| ↳ /C | 2863436.58x | 10065457.23x | 48436.58x |
| ↳ /JS | 1.54x | 5.40x | 0.026x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 698.16 μs | ~ 1022.38 μs | ~ 1.89 μs |
| ↳ /C | 694.35x | 1016.79x | 1.88x |
| ↳ /JS | 11.01x | 16.13x | 0.030x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 53877.92 μs | ~ 509843.88 μs | ~ 113830.46 μs |
| ↳ /C | 5.95x | 56.31x | 12.57x |
| ↳ /JS | 0.60x | 5.67x | 1.27x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 6237.63 μs | ~ 432965.96 μs | ~ 87761.96 μs |
| ↳ /C | 9309895.52x | 646217850.75x | 130988000.00x |
| ↳ /JS | 0.27x | 18.96x | 3.84x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 467.12 μs | ~ 560.06 μs | ~ 0.48 μs |
| ↳ /C | 699281.44x | 838413.17x | 718.56x |
| ↳ /JS | 9.69x | 11.61x | 0.010x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 791.27 μs | ~ 111183.96 μs | ~ 326.09 μs |
| ↳ /C | 1184535.93x | 166443053.89x | 488158.68x |
| ↳ /JS | 0.054x | 7.61x | 0.022x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 89.94 μs | ~ 26.94 μs | ~ 0.80 μs |
| ↳ /C | 1972.84x | 590.93x | 17.55x |
| ↳ /JS | 12.24x | 3.66x | 0.11x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.84 μs | ~ 0.85 μs | ~ 0.080 μs |
| ↳ /C | 1257.49x | 1272.46x | 119.76x |
| ↳ /JS | 20.21x | 20.45x | 1.92x |
| **Total Stable** | <strong>~ 65.78 ms</strong> | <strong>~ 1064.43 ms</strong> | <strong>~ 202.55 ms</strong> |
| ↳ /C | <strong>7.22x</strong> | <strong>116.88x</strong> | <strong>22.24x</strong> |
| ↳ μ/C | <strong>1004609.50x</strong> | <strong>58826576.72x</strong> | <strong>9394691.07x</strong> |
| ↳ Q₂/C | <strong>975.92x</strong> | <strong>803.86x</strong> | <strong>20.50x</strong> |
| ↳ G/C | <strong>1578.13x</strong> | <strong>4296.10x</strong> | <strong>214.24x</strong> |
| ↳ /JS | <strong>0.51x</strong> | <strong>8.24x</strong> | <strong>1.57x</strong> |
| ↳ μ/JS | <strong>7.27x</strong> | <strong>10.83x</strong> | <strong>1.94x</strong> |
| ↳ Q₂/JS | <strong>5.84x</strong> | <strong>9.61x</strong> | <strong>0.52x</strong> |
| ↳ G/JS | <strong>2.62x</strong> | <strong>7.14x</strong> | <strong>0.36x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
| **Stable** | | | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.24 μs | ~ 3.68 μs | ~ 0.23 μs | ~ 0.22 μs | ~ 0.23 μs |
| ↳ /C | 2.67x | 40.91x | 2.56x | 2.45x | 2.56x |
| ↳ /JS | 0.92x | 14.17x | 0.89x | 0.85x | 0.89x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.090 μs | ~ 11.78 μs | ~ 0.090 μs | ~ 0.12 μs | ~ 0.10 μs |
| ↳ /C | 1.17x | 152.69x | 1.17x | 1.56x | 1.30x |
| ↳ /JS | 0.29x | 38.05x | 0.29x | 0.39x | 0.32x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.11 μs | ~ 459.33 μs | ~ 17.07 μs | ~ 11.05 μs | ~ 0.51 μs |
| ↳ /C | 2.38x | 9938.12x | 369.33x | 239.08x | 11.03x |
| ↳ /JS | 0.011x | 47.90x | 1.78x | 1.15x | 0.053x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 29.88 μs | ~ 23547.50 μs | ~ 10.93 μs | ~ 34.70 μs | ~ 34.63 μs |
| ↳ /C | 0.93x | 732.64x | 0.34x | 1.08x | 1.08x |
| ↳ /JS | 0.16x | 124.27x | 0.058x | 0.18x | 0.18x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.14 μs | ~ 8728.98 μs | ~ 114.96 μs | ~ 107.97 μs | ~ 3.47 μs |
| ↳ /C | 0.97x | 2683.42x | 35.34x | 33.19x | 1.07x |
| ↳ /JS | 0.042x | 117.67x | 1.55x | 1.46x | 0.047x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.45 μs | ~ 625.00 μs | ~ 15.60 μs | ~ 18.59 μs | ~ 18.36 μs |
| ↳ /C | 0.98x | 37.14x | 0.93x | 1.10x | 1.09x |
| ↳ /JS | 0.24x | 9.26x | 0.23x | 0.28x | 0.27x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 159.52 μs | ~ 11792.17 μs | ~ 2249.51 μs | ~ 1033.78 μs | ~ 0.010 μs |
| ↳ /C | 235280.24x | 17392581.12x | 3317861.36x | 1524749.26x | 14.75x |
| ↳ /JS | 0.13x | 9.33x | 1.78x | 0.82x | 0.0000079x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 144.00 μs | ~ 2194.96 μs | ~ 120.43 μs | ~ 78.88 μs | ~ 0.78 μs |
| ↳ /C | 143.21x | 2182.97x | 119.77x | 78.45x | 0.78x |
| ↳ /JS | 2.27x | 34.63x | 1.90x | 1.24x | 0.012x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8136.98 μs | ~ 659207.04 μs | ~ 50862.42 μs | ~ 28737.38 μs | ~ 28247.96 μs |
| ↳ /C | 0.90x | 72.81x | 5.62x | 3.17x | 3.12x |
| ↳ /JS | 0.090x | 7.33x | 0.57x | 0.32x | 0.31x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.030 μs | ~ 1588595.46 μs | ~ 7471.33 μs | ~ 0.010 μs | ~ 0.010 μs |
| ↳ /C | 44.78x | 2371038000.00x | 11151238.81x | 14.93x | 14.93x |
| ↳ /JS | 0.0000013x | 69.58x | 0.33x | 0.00000044x | 0.00000044x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 44.57 μs | ~ 826.12 μs | ~ 73.12 μs | ~ 27.36 μs | ~ 0.010 μs |
| ↳ /C | 66721.56x | 1236706.59x | 109461.08x | 40958.08x | 14.97x |
| ↳ /JS | 0.92x | 17.13x | 1.52x | 0.57x | 0.00021x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.010 μs | ~ 159056.38 μs | ~ 38342.83 μs | ~ 23675.83 μs | ~ 0.010 μs |
| ↳ /C | 14.97x | 238108353.29x | 57399446.11x | 35442859.28x | 14.97x |
| ↳ /JS | 0.00000068x | 10.89x | 2.62x | 1.62x | 0.00000068x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.030 μs | ~ 298.70 μs | ~ 3.28 μs | ~ 0.58 μs | ~ 0.25 μs |
| ↳ /C | 0.66x | 6552.02x | 71.95x | 12.72x | 5.48x |
| ↳ /JS | 0.0041x | 40.64x | 0.45x | 0.079x | 0.034x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.070 μs | ~ 0.95 μs | ~ 0.10 μs | ~ 0.020 μs | ~ 0.010 μs |
| ↳ /C | 104.79x | 1422.16x | 149.70x | 29.94x | 14.97x |
| ↳ /JS | 1.68x | 22.85x | 2.41x | 0.48x | 0.24x |
| **Total Stable** | <strong>~ 8.54 ms</strong> | <strong>~ 2455.35 ms</strong> | <strong>~ 99.28 ms</strong> | <strong>~ 53.73 ms</strong> | <strong>~ 28.31 ms</strong> |
| ↳ /C | <strong>0.94x</strong> | <strong>269.61x</strong> | <strong>10.90x</strong> | <strong>5.90x</strong> | <strong>3.11x</strong> |
| ↳ μ/C | <strong>21594.30x</strong> | <strong>187699961.13x</strong> | <strong>5141340.29x</strong> | <strong>2643498.88x</strong> | <strong>7.29x</strong> |
| ↳ Q₂/C | <strong>2.52x</strong> | <strong>2433.19x</strong> | <strong>95.86x</strong> | <strong>22.43x</strong> | <strong>4.30x</strong> |
| ↳ G/C | <strong>18.63x</strong> | <strong>16141.85x</strong> | <strong>462.80x</strong> | <strong>118.90x</strong> | <strong>4.15x</strong> |
| ↳ /JS | <strong>0.066x</strong> | <strong>19.02x</strong> | <strong>0.77x</strong> | <strong>0.42x</strong> | <strong>0.22x</strong> |
| ↳ μ/JS | <strong>0.48x</strong> | <strong>40.26x</strong> | <strong>1.17x</strong> | <strong>0.67x</strong> | <strong>0.17x</strong> |
| ↳ Q₂/JS | <strong>0.14x</strong> | <strong>28.74x</strong> | <strong>1.20x</strong> | <strong>0.52x</strong> | <strong>0.050x</strong> |
| ↳ G/JS | <strong>0.031x</strong> | <strong>26.82x</strong> | <strong>0.77x</strong> | <strong>0.20x</strong> | <strong>0.0069x</strong> |
| **Incl WIP** | | | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| **Total Incl WIP** | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| ↳ μ/JS | — | — | — | — | — |
| ↳ Q₂/JS | — | — | — | — | — |
| ↳ G/JS | — | — | — | — | — |
| **Total Stable + Incl WIP** | — | — | — | — | — |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | — | — | — | — | — |
| ↳ μ/JS | — | — | — | — | — |
| ↳ Q₂/JS | — | — | — | — | — |
| ↳ G/JS | — | — | — | — | — |
| **Excl WIP** | | | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 6387.79 μs | — | — | — | — |
| ↳ /C | 0.97x | — | — | — | — |
| ↳ /JS | 0.25x | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 35351.17 μs | — | — | — | — |
| ↳ /C | 3.35x | — | — | — | — |
| ↳ /JS | 0.44x | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 1130.69 μs | — | — | — | — |
| ↳ /C | 1.72x | — | — | — | — |
| ↳ /JS | 0.12x | — | — | — | — |

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written FP-style C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written imperative C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.71 μs | ~ 0.94 μs | ~ 0.85 μs |
| ↳ /C | 30.15x | 10.44x | 9.48x |
| ↳ /JS | 10.44x | 3.62x | 3.28x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs | ~ 0.11 μs | ~ 0.024 μs |
| ↳ /C | 4.08x | 1.39x | 0.32x |
| ↳ /JS | 1.02x | 0.35x | 0.079x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 154.89 μs | ~ 35.16 μs | ~ 0.36 μs |
| ↳ /C | 3351.23x | 760.76x | 7.77x |
| ↳ /JS | 16.15x | 3.67x | 0.037x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 2844.59 μs | ~ 35.11 μs | ~ 33.85 μs |
| ↳ /C | 88.50x | 1.09x | 1.05x |
| ↳ /JS | 15.01x | 0.19x | 0.18x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1453.15 μs | ~ 567.68 μs | ~ 3.41 μs |
| ↳ /C | 446.72x | 174.51x | 1.05x |
| ↳ /JS | 19.59x | 7.65x | 0.046x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 221.50 μs | ~ 17.99 μs | ~ 16.54 μs |
| ↳ /C | 13.16x | 1.07x | 0.98x |
| ↳ /JS | 3.28x | 0.27x | 0.25x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 2630.26 μs | ~ 4240.89 μs | ~ 0.022 μs |
| ↳ /C | 3879439.90x | 6254992.99x | 32.15x |
| ↳ /JS | 2.08x | 3.36x | 0.000017x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 974.45 μs | ~ 257.70 μs | ~ 1.08 μs |
| ↳ /C | 969.13x | 256.30x | 1.07x |
| ↳ /JS | 15.37x | 4.07x | 0.017x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 613388.67 μs | ~ 66049.46 μs | ~ 22007.12 μs |
| ↳ /C | 67.75x | 7.30x | 2.43x |
| ↳ /JS | 6.82x | 0.73x | 0.24x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 265261.00 μs | ~ 27358.79 μs | ~ 0.022 μs |
| ↳ /C | 395911940.30x | 40834017.91x | 33.21x |
| ↳ /JS | 11.62x | 1.20x | 0.00000097x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 293.30 μs | ~ 185.70 μs | ~ 0.022 μs |
| ↳ /C | 439067.18x | 277992.84x | 33.45x |
| ↳ /JS | 6.08x | 3.85x | 0.00046x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 36615.58 μs | ~ 31474.58 μs | ~ 0.023 μs |
| ↳ /C | 54813748.50x | 47117640.72x | 35.02x |
| ↳ /JS | 2.51x | 2.15x | 0.0000016x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 36.57 μs | ~ 1.10 μs | ~ 0.74 μs |
| ↳ /C | 802.24x | 24.02x | 16.13x |
| ↳ /JS | 4.98x | 0.15x | 0.10x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.48 μs | ~ 0.30 μs | ~ 0.022 μs |
| ↳ /C | 722.08x | 444.00x | 33.52x |
| ↳ /JS | 11.60x | 7.13x | 0.54x |
| **Total Stable** | <strong>~ 923.88 ms</strong> | <strong>~ 130.23 ms</strong> | <strong>~ 22.06 ms</strong> |
| ↳ /C | <strong>101.45x</strong> | <strong>14.30x</strong> | <strong>2.42x</strong> |
| ↳ μ/C | <strong>32503620.78x</strong> | <strong>6749023.24x</strong> | <strong>14.83x</strong> |
| ↳ Q₂/C | <strong>762.16x</strong> | <strong>215.40x</strong> | <strong>8.62x</strong> |
| ↳ G/C | <strong>4033.62x</strong> | <strong>873.44x</strong> | <strong>5.76x</strong> |
| ↳ /JS | <strong>7.16x</strong> | <strong>1.01x</strong> | <strong>0.17x</strong> |
| ↳ μ/JS | <strong>9.04x</strong> | <strong>2.74x</strong> | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>8.63x</strong> | <strong>2.76x</strong> | <strong>0.062x</strong> |
| ↳ G/JS | <strong>6.70x</strong> | <strong>1.45x</strong> | <strong>0.0096x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 1.12 μs | ~ 0.50 μs | ~ 0.50 μs |
| ↳ /C | 12.45x | 5.56x | 5.56x |
| ↳ /JS | 4.31x | 1.92x | 1.92x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.36 μs | ~ 0.12 μs | ~ 0.11 μs |
| ↳ /C | 4.67x | 1.56x | 1.43x |
| ↳ /JS | 1.16x | 0.39x | 0.36x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 560.47 μs | ~ 16.05 μs | ~ 3.05 μs |
| ↳ /C | 12126.40x | 347.26x | 65.99x |
| ↳ /JS | 58.45x | 1.67x | 0.32x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.45 μs | ~ 37.60 μs | ~ 49.49 μs |
| ↳ /C | 1.79x | 1.17x | 1.54x |
| ↳ /JS | 0.30x | 0.20x | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 2545.27 μs | ~ 87.33 μs | ~ 7.21 μs |
| ↳ /C | 782.45x | 26.85x | 2.22x |
| ↳ /JS | 34.31x | 1.18x | 0.097x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.15 μs | ~ 7.70 μs | ~ 7.66 μs |
| ↳ /C | 0.90x | 0.46x | 0.46x |
| ↳ /JS | 0.22x | 0.11x | 0.11x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1678.49 μs | ~ 342.79 μs | ~ 25.24 μs |
| ↳ /C | 2475648.97x | 505589.97x | 37227.14x |
| ↳ /JS | 1.33x | 0.27x | 0.020x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 525.55 μs | ~ 45.26 μs | ~ 1.80 μs |
| ↳ /C | 522.68x | 45.01x | 1.79x |
| ↳ /JS | 8.29x | 0.71x | 0.028x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 19071.54 μs | ~ 36659.92 μs | ~ 9499.96 μs |
| ↳ /C | 2.11x | 4.05x | 1.05x |
| ↳ /JS | 0.21x | 0.41x | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2523.93 μs | ~ 23242.04 μs | ~ 2254.54 μs |
| ↳ /C | 3767059.70x | 34689611.94x | 3364985.07x |
| ↳ /JS | 0.11x | 1.02x | 0.099x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 169.78 μs | ~ 78.22 μs | ~ 3.21 μs |
| ↳ /C | 254161.68x | 117095.81x | 4805.39x |
| ↳ /JS | 3.52x | 1.62x | 0.067x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 6576.73 μs | ~ 9665.58 μs | ~ 235.04 μs |
| ↳ /C | 9845404.19x | 14469431.14x | 351856.29x |
| ↳ /JS | 0.45x | 0.66x | 0.016x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 62.07 μs | ~ 29.54 μs | ~ 15.63 μs |
| ↳ /C | 1361.51x | 647.96x | 342.85x |
| ↳ /JS | 8.44x | 4.02x | 2.13x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.24 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 359.28x | 89.82x | 29.94x |
| ↳ /JS | 5.77x | 1.44x | 0.48x |
| **Total Stable** | <strong>~ 33.79 ms</strong> | <strong>~ 70.21 ms</strong> | <strong>~ 12.10 ms</strong> |
| ↳ /C | <strong>3.71x</strong> | <strong>7.71x</strong> | <strong>1.33x</strong> |
| ↳ μ/C | <strong>1168389.20x</strong> | <strong>3555921.33x</strong> | <strong>268523.34x</strong> |
| ↳ Q₂/C | <strong>652.57x</strong> | <strong>67.42x</strong> | <strong>17.75x</strong> |
| ↳ G/C | <strong>1187.98x</strong> | <strong>449.16x</strong> | <strong>92.42x</strong> |
| ↳ /JS | <strong>0.26x</strong> | <strong>0.54x</strong> | <strong>0.094x</strong> |
| ↳ μ/JS | <strong>9.06x</strong> | <strong>1.12x</strong> | <strong>0.43x</strong> |
| ↳ Q₂/JS | <strong>2.42x</strong> | <strong>0.87x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | <strong>1.97x</strong> | <strong>0.75x</strong> | <strong>0.15x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.14 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.60x | 1.18x | 1.13x |
| ↳ /JS | 0.55x | 0.41x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.26 μs | ~ 0.099 μs | ~ 0.077 μs |
| ↳ /C | 3.36x | 1.28x | 0.99x |
| ↳ /JS | 0.84x | 0.32x | 0.25x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 5.03 μs | ~ 4.18 μs | ~ 0.57 μs |
| ↳ /C | 108.83x | 90.51x | 12.27x |
| ↳ /JS | 0.52x | 0.44x | 0.059x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 45.64 μs | ~ 38.54 μs | ~ 39.83 μs |
| ↳ /C | 1.42x | 1.20x | 1.24x |
| ↳ /JS | 0.24x | 0.20x | 0.21x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 73.28 μs | ~ 38.80 μs | ~ 4.18 μs |
| ↳ /C | 22.53x | 11.93x | 1.28x |
| ↳ /JS | 0.99x | 0.52x | 0.056x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 18.18 μs | ~ 7.60 μs | ~ 5.80 μs |
| ↳ /C | 1.08x | 0.45x | 0.34x |
| ↳ /JS | 0.27x | 0.11x | 0.086x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 310.51 μs | ~ 375.72 μs | ~ 0.0031 μs |
| ↳ /C | 457978.06x | 554153.62x | 4.60x |
| ↳ /JS | 0.25x | 0.30x | 0.0000025x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.29 μs | ~ 18.40 μs | ~ 0.54 μs |
| ↳ /C | 21.17x | 18.30x | 0.54x |
| ↳ /JS | 0.34x | 0.29x | 0.0086x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16788.38 μs | ~ 12065.75 μs | ~ 11828.17 μs |
| ↳ /C | 1.85x | 1.33x | 1.31x |
| ↳ /JS | 0.19x | 0.13x | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0054 μs | ~ 14583.08 μs | ~ 0.0034 μs |
| ↳ /C | 8.08x | 21765795.52x | 5.02x |
| ↳ /JS | 0.00000024x | 0.64x | 0.00000015x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 6.26 μs | ~ 17.25 μs | ~ 0.0033 μs |
| ↳ /C | 9366.28x | 25821.89x | 4.88x |
| ↳ /JS | 0.13x | 0.36x | 0.000068x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5.24 μs | ~ 8062.00 μs | ~ 0.037 μs |
| ↳ /C | 7847.98x | 12068862.28x | 55.24x |
| ↳ /JS | 0.00036x | 0.55x | 0.0000025x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.96 μs | ~ 3.06 μs | ~ 0.50 μs |
| ↳ /C | 64.88x | 67.19x | 10.91x |
| ↳ /JS | 0.40x | 0.42x | 0.068x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.013 μs | ~ 0.0062 μs | ~ 0.0023 μs |
| ↳ /C | 19.04x | 9.22x | 3.49x |
| ↳ /JS | 0.31x | 0.15x | 0.056x |
| **Total Stable** | <strong>~ 17.28 ms</strong> | <strong>~ 35.21 ms</strong> | <strong>~ 11.88 ms</strong> |
| ↳ /C | <strong>1.90x</strong> | <strong>3.87x</strong> | <strong>1.30x</strong> |
| ↳ μ/C | <strong>33960.44x</strong> | <strong>2458202.56x</strong> | <strong>7.37x</strong> |
| ↳ Q₂/C | <strong>20.11x</strong> | <strong>15.11x</strong> | <strong>2.40x</strong> |
| ↳ G/C | <strong>47.12x</strong> | <strong>183.78x</strong> | <strong>2.73x</strong> |
| ↳ /JS | <strong>0.13x</strong> | <strong>0.27x</strong> | <strong>0.092x</strong> |
| ↳ μ/JS | <strong>0.36x</strong> | <strong>0.35x</strong> | <strong>0.094x</strong> |
| ↳ Q₂/JS | <strong>0.29x</strong> | <strong>0.34x</strong> | <strong>0.058x</strong> |
| ↳ G/JS | <strong>0.078x</strong> | <strong>0.31x</strong> | <strong>0.0045x</strong> |
| **Incl WIP** | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| **Total Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Total Stable + Incl WIP** | — | — | — |
| ↳ /C | — | — | — |
| ↳ μ/C | — | — | — |
| ↳ Q₂/C | — | — | — |
| ↳ G/C | — | — | — |
| ↳ /JS | — | — | — |
| ↳ μ/JS | — | — | — |
| ↳ Q₂/JS | — | — | — |
| ↳ G/JS | — | — | — |
| **Excl WIP** | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| ↳ /C | — | — | — |
| ↳ /JS | — | — | — |

#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>koka -O3<br><br>(WIP) |
----------------------- | ----- |
| **Stable** | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.22 μs |
| ↳ /C | 2.43x |
| ↳ /JS | 0.84x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.60x |
| ↳ /JS | 0.40x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 4.28 μs |
| ↳ /C | 92.50x |
| ↳ /JS | 0.45x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 80.51 μs |
| ↳ /C | 2.50x |
| ↳ /JS | 0.42x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 158.94 μs |
| ↳ /C | 48.86x |
| ↳ /JS | 2.14x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.38 μs |
| ↳ /C | 1.15x |
| ↳ /JS | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 758.69 μs |
| ↳ /C | 1119008.11x |
| ↳ /JS | 0.60x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 12.17 μs |
| ↳ /C | 12.10x |
| ↳ /JS | 0.19x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8746.50 μs |
| ↳ /C | 0.97x |
| ↳ /JS | 0.097x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 35687.00 μs |
| ↳ /C | 53264179.10x |
| ↳ /JS | 1.56x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 5.21 μs |
| ↳ /C | 7793.50x |
| ↳ /JS | 0.11x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8982.00 μs |
| ↳ /C | 13446107.78x |
| ↳ /JS | 0.61x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 6.58 μs |
| ↳ /C | 144.30x |
| ↳ /JS | 0.89x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs |
| ↳ /C | 1.55x |
| ↳ /JS | 0.025x |
| **Total Stable** | <strong>~ 54.46 ms</strong> |
| ↳ /C | <strong>5.98x</strong> |
| ↳ μ/C | <strong>4845528.32x</strong> |
| ↳ Q₂/C | <strong>30.48x</strong> |
| ↳ G/C | <strong>224.51x</strong> |
| ↳ /JS | <strong>0.42x</strong> |
| ↳ μ/JS | <strong>0.62x</strong> |
| ↳ Q₂/JS | <strong>0.44x</strong> |
| ↳ G/JS | <strong>0.37x</strong> |
| **Incl WIP** | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Total Stable + Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Excl WIP** | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2<br><br>(WIP) |
----------------------- | ------------------------ |
| **Stable** | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.056 μs |
| ↳ /C | 0.62x |
| ↳ /JS | 0.21x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.19 μs |
| ↳ /C | 2.52x |
| ↳ /JS | 0.63x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 3.58 μs |
| ↳ /C | 77.44x |
| ↳ /JS | 0.37x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 62.84 μs |
| ↳ /C | 1.96x |
| ↳ /JS | 0.33x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.68 μs |
| ↳ /C | 2.36x |
| ↳ /JS | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 6.75 μs |
| ↳ /C | 0.40x |
| ↳ /JS | 0.10x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 143.90 μs |
| ↳ /C | 212239.58x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.58 μs |
| ↳ /C | 21.46x |
| ↳ /JS | 0.34x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 11559.00 μs |
| ↳ /C | 1.28x |
| ↳ /JS | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 4883.00 μs |
| ↳ /C | 7288059.70x |
| ↳ /JS | 0.21x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.016 μs |
| ↳ /C | 23.40x |
| ↳ /JS | 0.00032x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.46 μs |
| ↳ /C | 682.76x |
| ↳ /JS | 0.000031x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 4.33 μs |
| ↳ /C | 94.91x |
| ↳ /JS | 0.59x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0041 μs |
| ↳ /C | 6.09x |
| ↳ /JS | 0.098x |
| **Total Stable** | <strong>~ 16.69 ms</strong> |
| ↳ /C | <strong>1.83x</strong> |
| ↳ μ/C | <strong>535801.03x</strong> |
| ↳ Q₂/C | <strong>13.78x</strong> |
| ↳ G/C | <strong>43.53x</strong> |
| ↳ /JS | <strong>0.13x</strong> |
| ↳ μ/JS | <strong>0.23x</strong> |
| ↳ Q₂/JS | <strong>0.17x</strong> |
| ↳ G/JS | <strong>0.072x</strong> |
| **Incl WIP** | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Total Stable + Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Excl WIP** | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3<br><br>(WIP) |
----------------------- | --------------------------- |
| **Stable** | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.049 μs |
| ↳ /C | 0.55x |
| ↳ /JS | 0.19x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.61x |
| ↳ /JS | 0.40x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 1.98 μs |
| ↳ /C | 42.77x |
| ↳ /JS | 0.21x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 49.82 μs |
| ↳ /C | 1.55x |
| ↳ /JS | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 15.85 μs |
| ↳ /C | 4.87x |
| ↳ /JS | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.05 μs |
| ↳ /C | 0.95x |
| ↳ /JS | 0.24x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 140.38 μs |
| ↳ /C | 207042.77x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 17.94 μs |
| ↳ /C | 17.84x |
| ↳ /JS | 0.28x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10787.00 μs |
| ↳ /C | 1.19x |
| ↳ /JS | 0.12x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 12448.00 μs |
| ↳ /C | 18579104.48x |
| ↳ /JS | 0.55x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 10.03 μs |
| ↳ /C | 15009.53x |
| ↳ /JS | 0.21x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5690.00 μs |
| ↳ /C | 8517964.07x |
| ↳ /JS | 0.39x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 5.89 μs |
| ↳ /C | 129.27x |
| ↳ /JS | 0.80x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0047 μs |
| ↳ /C | 7.00x |
| ↳ /JS | 0.11x |
| **Total Stable** | <strong>~ 29.18 ms</strong> |
| ↳ /C | <strong>3.20x</strong> |
| ↳ μ/C | <strong>1951380.60x</strong> |
| ↳ Q₂/C | <strong>12.42x</strong> |
| ↳ G/C | <strong>148.68x</strong> |
| ↳ /JS | <strong>0.23x</strong> |
| ↳ μ/JS | <strong>0.29x</strong> |
| ↳ Q₂/JS | <strong>0.23x</strong> |
| ↳ G/JS | <strong>0.25x</strong> |
| **Incl WIP** | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Total Stable + Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Excl WIP** | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |

#### C (reference)

> **24 November 2026:** [Fresh 14-case reference, raw values and limitations](docs/benchmark-results/2026-11-24-c-reference.md). Tiny kernels are heavily simplified by clang; AST/RBT retain the historical allocation model. Excl WIP diagnostics keep their historical references.

C Benchmark             | Hand-written imperative C<br>↓<br>clang -O3<br><br>(WIP) |
----------------------- | -------------------- |
| **Stable** | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.090 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.35x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.077 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.25x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.046 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.0048x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 32.14 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.17x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.25 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.044x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.83 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.25x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 0.00068 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.00000054x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 1.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.016x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 9053.50 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.00067 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.000000029x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.00067 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.000014x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.00067 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.000000046x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.046 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.0062x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.00067 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.016x |
| **Total Stable** | <strong>~ 9.11 ms</strong> |
| ↳ /C | <strong>1.00x</strong> |
| ↳ μ/C | <strong>1.00x</strong> |
| ↳ Q₂/C | <strong>1.00x</strong> |
| ↳ G/C | <strong>1.00x</strong> |
| ↳ /JS | <strong>0.071x</strong> |
| ↳ μ/JS | <strong>0.086x</strong> |
| ↳ Q₂/JS | <strong>0.016x</strong> |
| ↳ G/JS | <strong>0.0017x</strong> |
| **Incl WIP** | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Total Stable + Incl WIP** | — |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | — |
| ↳ μ/JS | — |
| ↳ Q₂/JS | — |
| ↳ G/JS | — |
| **Excl WIP** | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 6589.00 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.26x |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 10560.13 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.13x |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 657.69 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.071x |

### Extended benchmark results (I/O, mutability, async)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
[File I/O](srx/Test/FileOps/FileOps.purs) | ~ 466899.83 μs | ~ 477117.79 μs | ~ 475960.00 μs | ~ 542377.75 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.02x | 1.02x | 1.16x |
[STArray Operations](srx/Test/STArray.purs) | ~ 0.92 μs | ~ 0.54 μs | ~ 0.29 μs | ~ 0.29 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.59x | 0.32x | 0.32x |
[String Operations](srx/Test/StringOps.purs) | ~ 256.71 μs | ~ 232.50 μs | ~ 559.04 μs | ~ 414.58 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.91x | 2.18x | 1.61x |
[Aff Operations](srx/Test/AffOperations.purs) | ~ 10125.21 μs | ~ 11581.83 μs | ~ 10500.21 μs | ~ 10903.21 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.14x | 1.04x | 1.08x |
[Parallelism](srx/Test/Parallelism.purs) | ~ 14275052.04 μs | ~ 14278260.46 μs | ~ 965397.12 μs | ~ 478303.67 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.00x | 0.070x | 0.030x |
**Total Execution Time** | <strong>~ 14752.33 ms</strong> | <strong>~ 14767.19 ms</strong> | <strong>~ 1452.42 ms</strong> | <strong>~ 1032.00 ms</strong> |
↳ /C | — | — | — | — |
↳ μ/C | — | — | — | — |
↳ Q₂/C | — | — | — | — |
↳ G/C | — | — | — | — |
↳ /JS | — | <strong>1.00x</strong> | <strong>0.10x</strong> | <strong>0.070x</strong> |
↳ μ/JS | — | <strong>0.93x</strong> | <strong>0.92x</strong> | <strong>0.84x</strong> |
↳ Q₂/JS | — | <strong>1.00x</strong> | <strong>1.02x</strong> | <strong>1.08x</strong> |
↳ G/JS | — | <strong>0.91x</strong> | <strong>0.55x</strong> | <strong>0.46x</strong> |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with the Go/Rust multicore runtimes.

### Application benchmark results (WIP)

App scenario | JS | Go | Go / JS |
--- | ---: | ---: | ---: |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (corpus parsing) | ~ 62.17 ms | ~ 128.79 ms | 2.07x |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (fixture parsing) | ~ 8.94 ms | ~ 17.67 ms | 1.98x |
[tidy](https://github.com/natefaubion/purescript-tidy) (parsing + formatting) | ~ 14.23 ms | ~ 23.65 ms | 1.66x |
[tidy](https://github.com/natefaubion/purescript-tidy) (10 iterations) | ~ 316.12 ms | ~ 239.07 ms | 0.76x |
[b8x](https://github.com/0x000000000000000000001/b8x.pub) (287 tests) | ~ 10.43 s | ~ 11.05 s | 1.06x |

### Compilation times...

#### ... to Go

Benchmark   | [gopurs](https://github.com/0x000000000000000000001/gopurs) JS binary<br>(official compiler)<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>with `--dev`<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ------------------ | --------------------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 146.39 s | ~ 78.04 s | ~ 53.61 s | ~ 52.92 s
↳ /JS | — | 0.53x | 0.37x | 0.36x
[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff) | ~ 16023.00 ms | ~ 6312.00 ms | ~ 5437.00 ms | ~ 5002.00 ms
↳ /JS | — | 0.39x | 0.34x | 0.31x
[gopurs-argonaut-core](https://github.com/0x000000000000000000001/gopurs-argonaut-core) | ~ 15672.00 ms | ~ 5930.00 ms | ~ 4962.00 ms | ~ 4800.00 ms
↳ /JS | — | 0.38x | 0.32x | 0.31x
[gopurs-arrays](https://github.com/0x000000000000000000001/gopurs-arrays) | ~ 12758.00 ms | ~ 5072.00 ms | ~ 3782.00 ms | ~ 4001.00 ms
↳ /JS | — | 0.40x | 0.30x | 0.31x
[gopurs-assert](https://github.com/0x000000000000000000001/gopurs-assert) | ~ 2514.00 ms | ~ 1305.00 ms | ~ 1187.00 ms | ~ 922.00 ms
↳ /JS | — | 0.52x | 0.47x | 0.37x
[gopurs-avar](https://github.com/0x000000000000000000001/gopurs-avar) | ~ 15097.00 ms | ~ 5865.00 ms | ~ 4971.00 ms | ~ 4770.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.32x
[gopurs-catenable-lists](https://github.com/0x000000000000000000001/gopurs-catenable-lists) | ~ 10188.00 ms | ~ 3918.00 ms | ~ 3387.00 ms | ~ 3024.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.30x
[gopurs-console](https://github.com/0x000000000000000000001/gopurs-console) | ~ 9476.00 ms | ~ 3207.00 ms | ~ 2880.00 ms | ~ 2513.00 ms
↳ /JS | — | 0.34x | 0.30x | 0.27x
[gopurs-datetime](https://github.com/0x000000000000000000001/gopurs-datetime) | ~ 14497.00 ms | ~ 5560.00 ms | ~ 4768.00 ms | ~ 4324.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.30x
[gopurs-effect](https://github.com/0x000000000000000000001/gopurs-effect) | ~ 4161.00 ms | ~ 1390.00 ms | ~ 1296.00 ms | ~ 994.00 ms
↳ /JS | — | 0.33x | 0.31x | 0.24x
[gopurs-enums](https://github.com/0x000000000000000000001/gopurs-enums) | ~ 10858.00 ms | ~ 3936.00 ms | ~ 3300.00 ms | ~ 3125.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.29x
[gopurs-exceptions](https://github.com/0x000000000000000000001/gopurs-exceptions) | ~ 4671.00 ms | ~ 1665.00 ms | ~ 1519.00 ms | ~ 1195.00 ms
↳ /JS | — | 0.36x | 0.33x | 0.26x
[gopurs-foldable-traversable](https://github.com/0x000000000000000000001/gopurs-foldable-traversable) | ~ 10499.00 ms | ~ 3841.00 ms | ~ 3357.00 ms | ~ 3057.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.29x
[gopurs-foreign](https://github.com/0x000000000000000000001/gopurs-foreign) | ~ 12846.00 ms | ~ 5309.00 ms | ~ 4621.00 ms | ~ 4347.00 ms
↳ /JS | — | 0.41x | 0.36x | 0.34x
[gopurs-foreign-object](https://github.com/0x000000000000000000001/gopurs-foreign-object) | ~ 15874.00 ms | ~ 6228.00 ms | ~ 5090.00 ms | ~ 5136.00 ms
↳ /JS | — | 0.39x | 0.32x | 0.32x
[gopurs-free](https://github.com/0x000000000000000000001/gopurs-free) | ~ 11960.00 ms | ~ 4840.00 ms | ~ 4252.00 ms | ~ 3881.00 ms
↳ /JS | — | 0.40x | 0.36x | 0.32x
[gopurs-functions](https://github.com/0x000000000000000000001/gopurs-functions) | ~ 4231.00 ms | ~ 1417.00 ms | ~ 1299.00 ms | ~ 960.00 ms
↳ /JS | — | 0.33x | 0.31x | 0.23x
[gopurs-integers](https://github.com/0x000000000000000000001/gopurs-integers) | ~ 10024.00 ms | ~ 3470.00 ms | ~ 3098.00 ms | ~ 2707.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-js-bigints](https://github.com/0x000000000000000000001/gopurs-js-bigints) | ~ 4625.00 ms | ~ 1596.00 ms | ~ 1443.00 ms | ~ 1179.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.25x
[gopurs-js-date](https://github.com/0x000000000000000000001/gopurs-js-date) | ~ 16044.00 ms | ~ 6227.00 ms | ~ 5297.00 ms | ~ 5007.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.31x
[gopurs-js-promise](https://github.com/0x000000000000000000001/gopurs-js-promise) | ~ 8189.00 ms | ~ 2787.00 ms | ~ 2531.00 ms | ~ 2196.00 ms
↳ /JS | — | 0.34x | 0.31x | 0.27x
[gopurs-js-promise-aff](https://github.com/0x000000000000000000001/gopurs-js-promise-aff) | ~ 16617.00 ms | ~ 6434.00 ms | ~ 5511.00 ms | ~ 5158.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.31x
[gopurs-js-uri](https://github.com/0x000000000000000000001/gopurs-js-uri) | ~ 9730.00 ms | ~ 3275.00 ms | ~ 2948.00 ms | ~ 2569.00 ms
↳ /JS | — | 0.34x | 0.30x | 0.26x
[gopurs-lazy](https://github.com/0x000000000000000000001/gopurs-lazy) | ~ 7605.00 ms | ~ 2758.00 ms | ~ 2526.00 ms | ~ 2178.00 ms
↳ /JS | — | 0.36x | 0.33x | 0.29x
[gopurs-node-buffer](https://github.com/0x000000000000000000001/gopurs-node-buffer) | ~ 14716.00 ms | ~ 5351.00 ms | ~ 4743.00 ms | ~ 4366.00 ms
↳ /JS | — | 0.36x | 0.32x | 0.30x
[gopurs-node-event-emitter](https://github.com/0x000000000000000000001/gopurs-node-event-emitter) | ~ 21845.00 ms | ~ 8936.00 ms | ~ 6631.00 ms | ~ 7607.00 ms
↳ /JS | — | 0.41x | 0.30x | 0.35x
[gopurs-node-fs](https://github.com/0x000000000000000000001/gopurs-node-fs) | ~ 18796.00 ms | ~ 6999.00 ms | ~ 5912.00 ms | ~ 5685.00 ms
↳ /JS | — | 0.37x | 0.31x | 0.30x
[gopurs-node-http](https://github.com/0x000000000000000000001/gopurs-node-http) | ~ 20650.00 ms | ~ 7475.00 ms | ~ 6248.00 ms | ~ 6173.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.30x
[gopurs-node-net](https://github.com/0x000000000000000000001/gopurs-node-net) | ~ 19243.00 ms | ~ 6945.00 ms | ~ 5865.00 ms | ~ 5630.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.29x
[gopurs-node-path](https://github.com/0x000000000000000000001/gopurs-node-path) | ~ 4257.00 ms | ~ 1302.00 ms | ~ 1204.00 ms | ~ 914.00 ms
↳ /JS | — | 0.31x | 0.28x | 0.21x
[gopurs-node-process](https://github.com/0x000000000000000000001/gopurs-node-process) | ~ 17942.00 ms | ~ 6684.00 ms | ~ 5685.00 ms | ~ 5391.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.30x
[gopurs-node-streams](https://github.com/0x000000000000000000001/gopurs-node-streams) | ~ 23385.00 ms | ~ 9204.00 ms | ~ 6884.00 ms | ~ 7938.00 ms
↳ /JS | — | 0.39x | 0.29x | 0.34x
[gopurs-now](https://github.com/0x000000000000000000001/gopurs-now) | ~ 18036.00 ms | ~ 6756.00 ms | ~ 5772.00 ms | ~ 5480.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.30x
[gopurs-nullable](https://github.com/0x000000000000000000001/gopurs-nullable) | ~ 14155.00 ms | ~ 5343.00 ms | ~ 4659.00 ms | ~ 4399.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.31x
[gopurs-numbers](https://github.com/0x000000000000000000001/gopurs-numbers) | ~ 10142.00 ms | ~ 3507.00 ms | ~ 3170.00 ms | ~ 2749.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-ordered-collections](https://github.com/0x000000000000000000001/gopurs-ordered-collections) | ~ 16381.00 ms | ~ 6449.00 ms | ~ 5138.00 ms | ~ 5047.00 ms
↳ /JS | — | 0.39x | 0.31x | 0.31x
[gopurs-partial](https://github.com/0x000000000000000000001/gopurs-partial) | ~ 9519.00 ms | ~ 3188.00 ms | ~ 2859.00 ms | ~ 2493.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-prelude](https://github.com/0x000000000000000000001/gopurs-prelude) | ~ 10711.00 ms | ~ 3557.00 ms | ~ 3053.00 ms | ~ 2792.00 ms
↳ /JS | — | 0.33x | 0.29x | 0.26x
[gopurs-random](https://github.com/0x000000000000000000001/gopurs-random) | ~ 9902.00 ms | ~ 3224.00 ms | ~ 2880.00 ms | ~ 2525.00 ms
↳ /JS | — | 0.33x | 0.29x | 0.25x
[gopurs-record](https://github.com/0x000000000000000000001/gopurs-record) | ~ 10133.00 ms | ~ 3356.00 ms | ~ 3019.00 ms | ~ 2627.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-refs](https://github.com/0x000000000000000000001/gopurs-refs) | ~ 9881.00 ms | ~ 3300.00 ms | ~ 2960.00 ms | ~ 2594.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-run](https://github.com/0x000000000000000000001/gopurs-run) | ~ 17665.00 ms | ~ 6775.00 ms | ~ 5776.00 ms | ~ 5433.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.31x
[gopurs-spec](https://github.com/0x000000000000000000001/gopurs-spec) | ~ 30416.00 ms | ~ 12821.00 ms | ~ 8658.00 ms | ~ 10954.00 ms
↳ /JS | — | 0.42x | 0.28x | 0.36x
[gopurs-st](https://github.com/0x000000000000000000001/gopurs-st) | ~ 9659.00 ms | ~ 3302.00 ms | ~ 2980.00 ms | ~ 2597.00 ms
↳ /JS | — | 0.34x | 0.31x | 0.27x
[gopurs-strings](https://github.com/0x000000000000000000001/gopurs-strings) | ~ 12680.00 ms | ~ 4489.00 ms | ~ 3824.00 ms | ~ 3805.00 ms
↳ /JS | — | 0.35x | 0.30x | 0.30x
[gopurs-strings-extra](https://github.com/0x000000000000000000001/gopurs-strings-extra) | ~ 13757.00 ms | ~ 6990.00 ms | ~ 6343.00 ms | ~ 4827.00 ms
↳ /JS | — | 0.51x | 0.46x | 0.35x
[gopurs-unfoldable](https://github.com/0x000000000000000000001/gopurs-unfoldable) | ~ 9565.00 ms | ~ 3322.00 ms | ~ 2979.00 ms | ~ 2607.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-unsafe-coerce](https://github.com/0x000000000000000000001/gopurs-unsafe-coerce) | ~ 9841.00 ms | ~ 3290.00 ms | ~ 2941.00 ms | ~ 2562.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-uuid](https://github.com/0x000000000000000000001/gopurs-uuid) | ~ 20972.00 ms | ~ 8748.00 ms | ~ 6505.00 ms | ~ 7340.00 ms
↳ /JS | — | 0.42x | 0.31x | 0.35x
[gopurs-variant](https://github.com/0x000000000000000000001/gopurs-variant) | ~ 11455.00 ms | ~ 4451.00 ms | ~ 3881.00 ms | ~ 3279.00 ms
↳ /JS | — | 0.39x | 0.34x | 0.29x
[gopurs-yoga-json](https://github.com/0x000000000000000000001/gopurs-yoga-json) | ~ 36226.00 ms | ~ 16572.00 ms | ~ 11991.00 ms | ~ 13134.00 ms
↳ /JS | — | 0.46x | 0.33x | 0.36x
**Total gopurs-*** | <strong>~ 666.09 s</strong> | <strong>~ 254.68 s</strong> | <strong>~ 212.02 s</strong> | <strong>~ 203.99 s</strong>
↳ /JS | — | <strong>0.38x</strong> | <strong>0.32x</strong> | <strong>0.31x</strong>
**Total** | <strong>~ 812.48 s</strong> | <strong>~ 332.71 s</strong> | <strong>~ 265.63 s</strong> | <strong>~ 256.92 s</strong>
↳ /JS | — | <strong>0.41x</strong> | <strong>0.33x</strong> | <strong>0.32x</strong>

#### ... to Rust

Benchmark   | [purust](https://github.com/0x000000000000000000001/purust) JS binary<br>(official compiler)<br>(WIP) | [purust](https://github.com/0x000000000000000000001/purust) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 44.31 s | ~ 20.62 s
↳ /JS | — | 0.47x
[purust-aff](https://github.com/0x000000000000000000001/purust-aff) | ~ 6218.00 ms | ~ 2076.00 ms
↳ /JS | — | 0.33x
[purust-argonaut-codecs](https://github.com/0x000000000000000000001/purust-argonaut-codecs) | ~ 6090.00 ms | ~ 2120.00 ms
↳ /JS | — | 0.35x
[purust-argonaut-core](https://github.com/0x000000000000000000001/purust-argonaut-core) | ~ 5122.00 ms | ~ 1682.00 ms
↳ /JS | — | 0.33x
[purust-arraybuffer-types](https://github.com/0x000000000000000000001/purust-arraybuffer-types) | ~ 591.00 ms | ~ 158.00 ms
↳ /JS | — | 0.27x
[purust-arrays](https://github.com/0x000000000000000000001/purust-arrays) | ~ 15134.00 ms | ~ 4291.00 ms
↳ /JS | — | 0.28x
[purust-assert](https://github.com/0x000000000000000000001/purust-assert) | ~ 631.00 ms | ~ 170.00 ms
↳ /JS | — | 0.27x
[purust-avar](https://github.com/0x000000000000000000001/purust-avar) | ~ 5942.00 ms | ~ 1938.00 ms
↳ /JS | — | 0.33x
[purust-catenable-lists](https://github.com/0x000000000000000000001/purust-catenable-lists) | ~ 3474.00 ms | ~ 1177.00 ms
↳ /JS | — | 0.34x
[purust-console](https://github.com/0x000000000000000000001/purust-console) | ~ 664.00 ms | ~ 183.00 ms
↳ /JS | — | 0.28x
[purust-datetime](https://github.com/0x000000000000000000001/purust-datetime) | ~ 5659.00 ms | ~ 1902.00 ms
↳ /JS | — | 0.34x
[purust-effect](https://github.com/0x000000000000000000001/purust-effect) | ~ 1240.00 ms | ~ 475.00 ms
↳ /JS | — | 0.38x
[purust-enums](https://github.com/0x000000000000000000001/purust-enums) | ~ 2950.00 ms | ~ 923.00 ms
↳ /JS | — | 0.31x
[purust-exceptions](https://github.com/0x000000000000000000001/purust-exceptions) | ~ 924.00 ms | ~ 276.00 ms
↳ /JS | — | 0.30x
[purust-exists](https://github.com/0x000000000000000000001/purust-exists) | ~ 629.00 ms | ~ 171.00 ms
↳ /JS | — | 0.27x
[purust-foldable-traversable](https://github.com/0x000000000000000000001/purust-foldable-traversable) | ~ 3593.00 ms | ~ 1220.00 ms
↳ /JS | — | 0.34x
[purust-foreign](https://github.com/0x000000000000000000001/purust-foreign) | ~ 4949.00 ms | ~ 1647.00 ms
↳ /JS | — | 0.33x
[purust-foreign-object](https://github.com/0x000000000000000000001/purust-foreign-object) | ~ 6320.00 ms | ~ 2217.00 ms
↳ /JS | — | 0.35x
[purust-free](https://github.com/0x000000000000000000001/purust-free) | ~ 4570.00 ms | ~ 1520.00 ms
↳ /JS | — | 0.33x
[purust-functions](https://github.com/0x000000000000000000001/purust-functions) | ~ 1365.00 ms | ~ 508.00 ms
↳ /JS | — | 0.37x
[purust-integers](https://github.com/0x000000000000000000001/purust-integers) | ~ 4359.00 ms | ~ 1398.00 ms
↳ /JS | — | 0.32x
[purust-js-bigints](https://github.com/0x000000000000000000001/purust-js-bigints) | ~ 6620.00 ms | ~ 2336.00 ms
↳ /JS | — | 0.35x
[purust-js-date](https://github.com/0x000000000000000000001/purust-js-date) | ~ 5642.00 ms | ~ 1933.00 ms
↳ /JS | — | 0.34x
[purust-js-promise](https://github.com/0x000000000000000000001/purust-js-promise) | ~ 6029.00 ms | ~ 2018.00 ms
↳ /JS | — | 0.33x
[purust-js-promise-aff](https://github.com/0x000000000000000000001/purust-js-promise-aff) | ~ 6165.00 ms | ~ 2078.00 ms
↳ /JS | — | 0.34x
[purust-lazy](https://github.com/0x000000000000000000001/purust-lazy) | ~ 2140.00 ms | ~ 693.00 ms
↳ /JS | — | 0.32x
[purust-node-buffer](https://github.com/0x000000000000000000001/purust-node-buffer) | ~ 4894.00 ms | ~ 1601.00 ms
↳ /JS | — | 0.33x
[purust-node-child-process](https://github.com/0x000000000000000000001/purust-node-child-process) | ~ 7651.00 ms | ~ 2609.00 ms
↳ /JS | — | 0.34x
[purust-node-event-emitter](https://github.com/0x000000000000000000001/purust-node-event-emitter) | ~ 8202.00 ms | ~ 2809.00 ms
↳ /JS | — | 0.34x
[purust-node-fs](https://github.com/0x000000000000000000001/purust-node-fs) | ~ 7430.00 ms | ~ 2508.00 ms
↳ /JS | — | 0.34x
[purust-node-http](https://github.com/0x000000000000000000001/purust-node-http) | ~ 7105.00 ms | ~ 2452.00 ms
↳ /JS | — | 0.35x
[purust-node-net](https://github.com/0x000000000000000000001/purust-node-net) | ~ 6805.00 ms | ~ 2310.00 ms
↳ /JS | — | 0.34x
[purust-node-os](https://github.com/0x000000000000000000001/purust-node-os) | ~ 6032.00 ms | ~ 2086.00 ms
↳ /JS | — | 0.35x
[purust-node-path](https://github.com/0x000000000000000000001/purust-node-path) | ~ 650.00 ms | ~ 175.00 ms
↳ /JS | — | 0.27x
[purust-node-process](https://github.com/0x000000000000000000001/purust-node-process) | ~ 7196.00 ms | ~ 2443.00 ms
↳ /JS | — | 0.34x
[purust-node-streams](https://github.com/0x000000000000000000001/purust-node-streams) | ~ 9197.00 ms | ~ 3489.00 ms
↳ /JS | — | 0.38x
[purust-now](https://github.com/0x000000000000000000001/purust-now) | ~ 6780.00 ms | ~ 2310.00 ms
↳ /JS | — | 0.34x
[purust-nullable](https://github.com/0x000000000000000000001/purust-nullable) | ~ 4653.00 ms | ~ 1557.00 ms
↳ /JS | — | 0.33x
[purust-numbers](https://github.com/0x000000000000000000001/purust-numbers) | ~ 4401.00 ms | ~ 1408.00 ms
↳ /JS | — | 0.32x
[purust-ordered-collections](https://github.com/0x000000000000000000001/purust-ordered-collections) | ~ 8217.00 ms | ~ 2986.00 ms
↳ /JS | — | 0.36x
[purust-partial](https://github.com/0x000000000000000000001/purust-partial) | ~ 871.00 ms | ~ 268.00 ms
↳ /JS | — | 0.31x
[purust-prelude](https://github.com/0x000000000000000000001/purust-prelude) | ~ 2793.00 ms | ~ 985.00 ms
↳ /JS | — | 0.35x
[purust-quickcheck](https://github.com/0x000000000000000000001/purust-quickcheck) | ~ 4753.00 ms | ~ 1624.00 ms
↳ /JS | — | 0.34x
[purust-random](https://github.com/0x000000000000000000001/purust-random) | ~ 2660.00 ms | ~ 898.00 ms
↳ /JS | — | 0.34x
[purust-record](https://github.com/0x000000000000000000001/purust-record) | ~ 2641.00 ms | ~ 896.00 ms
↳ /JS | — | 0.34x
[purust-refs](https://github.com/0x000000000000000000001/purust-refs) | ~ 718.00 ms | ~ 196.00 ms
↳ /JS | — | 0.27x
[purust-run](https://github.com/0x000000000000000000001/purust-run) | ~ 6914.00 ms | ~ 2262.00 ms
↳ /JS | — | 0.33x
[purust-spec](https://github.com/0x000000000000000000001/purust-spec) | ~ 12717.00 ms | ~ 4633.00 ms
↳ /JS | — | 0.36x
[purust-spec-discovery](https://github.com/0x000000000000000000001/purust-spec-discovery) | ~ 11401.00 ms | ~ 4318.00 ms
↳ /JS | — | 0.38x
[purust-spec-node](https://github.com/0x000000000000000000001/purust-spec-node) | ~ 11931.00 ms | ~ 4307.00 ms
↳ /JS | — | 0.36x
[purust-st](https://github.com/0x000000000000000000001/purust-st) | ~ 2554.00 ms | ~ 866.00 ms
↳ /JS | — | 0.34x
[purust-strings](https://github.com/0x000000000000000000001/purust-strings) | ~ 6589.00 ms | ~ 1851.00 ms
↳ /JS | — | 0.28x
[purust-strings-extra](https://github.com/0x000000000000000000001/purust-strings-extra) | ~ 7666.00 ms | ~ 11028.00 ms
↳ /JS | — | 1.44x
[purust-unfoldable](https://github.com/0x000000000000000000001/purust-unfoldable) | ~ 2651.00 ms | ~ 920.00 ms
↳ /JS | — | 0.35x
[purust-unsafe-coerce](https://github.com/0x000000000000000000001/purust-unsafe-coerce) | ~ 2089.00 ms | ~ 669.00 ms
↳ /JS | — | 0.32x
[purust-uuid](https://github.com/0x000000000000000000001/purust-uuid) | ~ 7613.00 ms | ~ 2588.00 ms
↳ /JS | — | 0.34x
[purust-variant](https://github.com/0x000000000000000000001/purust-variant) | ~ 4537.00 ms | ~ 1405.00 ms
↳ /JS | — | 0.31x
[purust-yoga-json](https://github.com/0x000000000000000000001/purust-yoga-json) | ~ 16694.00 ms | ~ 14548.00 ms
↳ /JS | — | 0.87x
**Total purust-*** | <strong>~ 304.03 s</strong> | <strong>~ 120.12 s</strong>
↳ /JS | — | <strong>0.40x</strong>
**Total** | <strong>~ 348.34 s</strong> | <strong>~ 140.73 s</strong>
↳ /JS | — | <strong>0.40x</strong>
