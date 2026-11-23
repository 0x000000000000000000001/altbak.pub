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

### Memory-management fairness

Native references use conventional memory management and return owned, usable
results with the same lifetime requirements as the compiled implementations.
Hand-written arenas, pools, bulk-lifetime allocation, or cross-iteration storage
recycling introduced specifically to improve a benchmark score are not allowed.
Compiler/runtime optimizations that preserve the program's ownership and lifetime
contract remain valid, including allocation elimination and stack allocation.
Construction, copying, conversion and eventual reclamation costs must be measured
and reported; deferred work must not disappear beyond a timing boundary.

### Core vs extended tests (`srx/`)
To ensure fair and executable comparisons across all backends, the test suite is split into two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) executed on the configured core backends via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Core stresstest benchmark results (pure computational)

> [!NOTE]
> **Reading the benchmark tables**
>
> - **Stable** — the 14 core cases, measured on every column. `Total Stable` is their sum.
> - **Incl WIP** — the 23 shared WIP micro-benchmarks. They are measured on every JS/Go column today; other families show `—` until their WIP campaign runs. `Total Incl WIP` is their sum (`—` until measured).
> - **Excl WIP** — the three single-column WIP diagnostics (Array Indexing, JSON to Typed AST, JSON Decoding), listed in every table and measured for JavaScript, Go, Rust and C. They are shown for reference only and are **not** part of any total.
> - **Total Stable + Incl WIP** — the 37 comparable cases; `—` until the Incl WIP cases are measured.
> - In every core language table, each case has three rows: its time, `↳ /C`, and `↳ /JS`. Rows beginning with `↳` belong to the case or total above them.
> - In the core language tables, `/JS` is the ratio to the **first column of the JavaScript table** (hand-written PureScript, official compiler, JS on V8), and `/C` is the ratio to the **first column of the C (reference) table** (hand-written imperative C, `clang -O3`). Extended Results uses its own first column as the JavaScript reference and has no C reference. The JavaScript reference column shows `—` for `/JS`; the C reference column shows `/C = 1.00x`.
> - On a total row, `↳ /C` and `↳ /JS` are ratios of the sums of case timings, equivalently arithmetic means of case ratios weighted by the reference time. They describe this benchmark mix, not an end-to-end application measurement.
> - `μ` is the unweighted arithmetic mean of case ratios; `Q₂` is their median (second quartile; the mean of the two middle ratios for an even count); `G` is their geometric mean, `exp(mean(log(ratio)))`. Each eligible case has equal weight in these three summaries. `/C` and `/JS` identify the reference; lower is faster, and `1.00x` means equal time.
> - `μ`, `Q₂` and `G` use the same eligible pairs. Timings displayed as bounds (`< ...`) are excluded, even if a historical per-case ratio cell contains a number. A suffix such as `(n=9/14)` gives the number of eligible cases out of the block's full size; the current C reference has nine eligible Stable cases. These partial summaries must not be described as covering all 14 cases.
> - Matching archived timings are used before rounding; rounded ratio cells are never averaged. `≈` marks a summary that falls back to published timing precision because a matching original campaign is unavailable. [Archived timing catalogue and provenance](bin/benchmark/ratio-measurements.json).
> - `—` means the candidate or reference block has not been measured, no eligible pairs remain, or this is the JavaScript reference compared with itself. A missing WIP reference prevents a combined 37-case summary; it is not replaced by the Stable subset. Values shown as `< 0.01x` are bounds.
> - `μ` is sensitive to large ratios, `Q₂` describes the middle case, and `G` summarizes multiplicative factors. None substitutes for an application benchmark; very short reference timings also deserve caution.
> - Totals are sums of per-case medians, not a single measured run.

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
| **Stable** | | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.26 μs | ~ 0.19 μs | ~ 0.20 μs | ~ 0.22 μs |
| ↳ /C | 2.68x | 1.96x | 2.06x | 2.27x |
| ↳ /JS | — | 0.73x | 0.77x | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs | ~ 0.31 μs | ~ 0.31 μs | ~ 0.37 μs |
| ↳ /C | 3.73x | 3.73x | 3.73x | 4.46x |
| ↳ /JS | — | 1.00x | 1.00x | 1.19x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 9.59 μs | ~ 5.06 μs | ~ 19.06 μs | ~ 0.49 μs |
| ↳ /C | 188.04x | 99.22x | 373.73x | 9.61x |
| ↳ /JS | — | 0.53x | 1.99x | 0.05x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 189.49 μs | ~ 151.13 μs | ~ 65.91 μs | ~ 68.31 μs |
| ↳ /C | 5.53x | 4.41x | 1.92x | 1.99x |
| ↳ /JS | — | 0.80x | 0.35x | 0.36x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 74.18 μs | ~ 134.11 μs | ~ 42.07 μs | ~ 7.69 μs |
| ↳ /C | 21.66x | 39.16x | 12.28x | 2.25x |
| ↳ /JS | — | 1.81x | 0.57x | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 67.49 μs | ~ 64.40 μs | ~ 38.25 μs | ~ 38.48 μs |
| ↳ /C | 3.99x | 3.81x | 2.26x | 2.28x |
| ↳ /JS | — | 0.95x | 0.57x | 0.57x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1263.82 μs | ~ 1242.29 μs | ~ 1337.51 μs | ~ 49.63 μs |
| ↳ /C | 1263820.00x | 1242290.00x | 1337510.00x | 49630.00x |
| ↳ /JS | — | 0.98x | 1.06x | 0.04x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 63.39 μs | ~ 26.99 μs | ~ 86.33 μs | ~ 0.65 μs |
| ↳ /C | 58.10x | 24.74x | 79.13x | 0.60x |
| ↳ /JS | — | 0.43x | 1.36x | 0.01x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 89940.58 μs | ~ 47272.88 μs | ~ 24486.83 μs | ~ 27249.58 μs |
| ↳ /C | 9.42x | 4.95x | 2.56x | 2.85x |
| ↳ /JS | — | 0.53x | 0.27x | 0.30x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 22830.25 μs | ~ 7368.38 μs | ~ 4843 μs | ~ 2805.25 μs |
| ↳ /C | 22830250.00x | 7368380.00x | 4843000.00x | 2805250.00x |
| ↳ /JS | — | 0.32x | 0.21x | 0.12x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 48.23 μs | ~ 10.95 μs | ~ 46.87 μs | ~ 0.42 μs |
| ↳ /C | 48230.00x | 10950.00x | 46870.00x | 420.00x |
| ↳ /JS | — | 0.23x | 0.97x | 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 14608.29 μs | ~ 10604.88 μs | ~ 17513.12 μs | ~ 321.48 μs |
| ↳ /C | 14608290.00x | 10604880.00x | 17513120.00x | 321480.00x |
| ↳ /JS | — | 0.73x | 1.20x | 0.02x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.35 μs | ~ 3.75 μs | ~ 3.12 μs | ~ 0.50 μs |
| ↳ /C | 147.00x | 75.00x | 62.40x | 10.00x |
| ↳ /JS | — | 0.51x | 0.42x | 0.07x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.04 μs | ~ 0.02 μs | ~ 0.06 μs | ~ 0.02 μs |
| ↳ /C | 40.00x | 20.00x | 60.00x | 20.00x |
| ↳ /JS | — | 0.50x | 1.50x | 0.50x |
| **Total Stable** | <strong>~ 129.10 ms</strong> | <strong>~ 66.89 ms</strong> | <strong>~ 48.48 ms</strong> | <strong>~ 30.54 ms</strong> |
| ↳ /C | <strong>13.43x</strong> | <strong>6.96x</strong> | <strong>5.04x</strong> | <strong>3.18x</strong> |
| ↳ μ/C | <strong>≈ 49.34x</strong> (n=9/14) | <strong>≈ 28.79x</strong> (n=9/14) | <strong>≈ 60.85x</strong> (n=9/14) | <strong>≈ 4.07x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 9.42x</strong> (n=9/14) | <strong>≈ 4.95x</strong> (n=9/14) | <strong>≈ 3.94x</strong> (n=9/14) | <strong>≈ 2.28x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 16.14x</strong> (n=9/14) | <strong>≈ 11.81x</strong> (n=9/14) | <strong>≈ 10.79x</strong> (n=9/14) | <strong>≈ 2.96x</strong> (n=9/14) |
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
| ↳ /JS | — | 1.00x | 0.63x | 0.05x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 561.69 μs | ~ 517.87 μs | ~ 114.68 μs | ~ 78.30 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.92x | 0.20x | 0.14x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 51.63 μs | ~ 50.71 μs | ~ 41.04 μs | ~ 5.03 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.98x | 0.79x | 0.10x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.20 μs | ~ 31.05 μs | ~ 2.05 μs | ~ 2.40 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.03x | 0.07x | 0.08x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 80.61 μs | ~ 81.04 μs | ~ 49.83 μs | ~ 1.39 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 1.01x | 0.62x | 0.02x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 68.35 μs | ~ 65.14 μs | ~ 42.18 μs | ~ 1.07 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.95x | 0.62x | 0.02x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 68.01 μs | ~ 64.65 μs | ~ 42.45 μs | ~ 1.08 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.95x | 0.62x | 0.02x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.01 μs | ~ 0.01 μs | ~ 0.02 μs | ~ 0.02 μs |
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
| ↳ /JS | — | 0.23x | 0.01x | 0.01x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 513.20 μs | ~ 106.03 μs | ~ 31.67 μs | ~ 3.41 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.21x | 0.06x | 0.01x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 155.67 μs | ~ 102.38 μs | ~ 2.39 μs | ~ 2.56 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.66x | 0.02x | 0.02x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 211.84 μs | ~ 113.53 μs | ~ 31.14 μs | ~ 3.52 μs |
| ↳ /C | — | — | — | — |
| ↳ /JS | — | 0.54x | 0.15x | 0.02x |
| **Total Incl WIP** | <strong>~ 2.88 ms</strong> | <strong>~ 1.95 ms</strong> | <strong>~ 0.92 ms</strong> | <strong>~ 0.27 ms</strong> |
| ↳ /C | — | — | — | — |
| ↳ μ/C | — | — | — | — |
| ↳ Q₂/C | — | — | — | — |
| ↳ G/C | — | — | — | — |
| ↳ /JS | — | <strong>0.68x</strong> | <strong>0.32x</strong> | <strong>0.09x</strong> |
| ↳ μ/JS | — | <strong>1.02x</strong> | <strong>0.81x</strong> | <strong>0.27x</strong> |
| ↳ Q₂/JS | — | <strong>0.99x</strong> | <strong>0.76x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | — | <strong>0.91x</strong> | <strong>0.44x</strong> | <strong>0.09x</strong> |
| **Total Stable + Incl WIP** | <strong>~ 131.98 ms</strong> | <strong>~ 68.84 ms</strong> | <strong>~ 49.40 ms</strong> | <strong>~ 30.81 ms</strong> |
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
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.32 μs | ~ 0.52 μs | ~ 4.17 μs | ~ 0.32 μs | ~ 0.32 μs |
| ↳ /C | 3.17x | 5.36x | 42.99x | 3.30x | 3.30x |
| ↳ /JS | 1.22x | 2.00x | 16.04x | 1.23x | 1.23x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.20 μs | ~ 0.54 μs | ~ 5.41 μs | ~ 0.16 μs | ~ 0.16 μs |
| ↳ /C | 2.54x | 6.51x | 65.18x | 1.93x | 1.93x |
| ↳ /JS | 0.65x | 1.74x | 17.45x | 0.52x | 0.52x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 8.83 μs | ~ 28.86 μs | ~ 236.01 μs | ~ 24.39 μs | ~ 0.45 μs |
| ↳ /C | 176.53x | 565.88x | 4627.65x | 478.24x | 8.82x |
| ↳ /JS | 0.92x | 3.01x | 24.61x | 2.54x | 0.05x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 61.64 μs | ~ 355.46 μs | ~ 6281.96 μs | ~ 35.98 μs | ~ 35.50 μs |
| ↳ /C | 1.80x | 10.37x | 183.24x | 1.05x | 1.04x |
| ↳ /JS | 0.33x | 1.88x | 33.15x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 8.54 μs | ~ 186.36 μs | ~ 5461.52 μs | ~ 40.43 μs | ~ 14.87 μs |
| ↳ /C | 2.50x | 54.41x | 1594.60x | 11.80x | 4.34x |
| ↳ /JS | 0.12x | 2.51x | 73.63x | 0.55x | 0.20x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.16 μs | ~ 35.56 μs | ~ 402.04 μs | ~ 20.85 μs | ~ 19.87 μs |
| ↳ /C | 0.96x | 2.10x | 23.78x | 1.23x | 1.18x |
| ↳ /JS | 0.24x | 0.53x | 5.96x | 0.31x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 226.67 μs | ~ 1296.16 μs | ~ 3811.48 μs | ~ 1044.60 μs | ~ 22.45 μs |
| ↳ /C | — | 1296160.00x | 3811480.00x | 1044600.00x | 22450.00x |
| ↳ /JS | 0.18x | 1.03x | 3.02x | 0.83x | 0.02x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 73.41 μs | ~ 76.41 μs | ~ 1340.99 μs | ~ 110.87 μs | ~ 0.62 μs |
| ↳ /C | 67.35x | 70.04x | 1229.14x | 101.62x | 0.57x |
| ↳ /JS | 1.16x | 1.21x | 21.15x | 1.75x | 0.01x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 9763.54 μs | ~ 31239.75 μs | ~ 870404.62 μs | ~ 24630.67 μs | ~ 9004.48 μs |
| ↳ /C | 1.02x | 3.27x | 91.14x | 2.58x | 0.94x |
| ↳ /JS | 0.11x | 0.35x | 9.68x | 0.27x | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.04 μs | ~ 15617.54 μs | ~ 505069 μs | ~ 49160.92 μs | ~ 2244.45 μs |
| ↳ /C | — | 15617540.00x | 505069000.00x | 49160920.00x | 2244450.00x |
| ↳ /JS | < 0.01x | 0.68x | 22.12x | 2.15x | 0.10x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 91.18 μs | ~ 110.57 μs | ~ 303.56 μs | ~ 52.10 μs | ~ 0.33 μs |
| ↳ /C | — | 110570.00x | 303560.00x | 52100.00x | 330.00x |
| ↳ /JS | 1.89x | 2.29x | 6.29x | 1.08x | 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.76 μs | ~ 1566.95 μs | ~ 71378.08 μs | ~ 13865.38 μs | ~ 0.27 μs |
| ↳ /C | — | 1566950.00x | 71378080.00x | 13865380.00x | 270.00x |
| ↳ /JS | < 0.01x | 0.11x | 4.89x | 0.95x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.64 μs | ~ 32.70 μs | ~ 43.81 μs | ~ 4.90 μs | ~ 0.45 μs |
| ↳ /C | 52.76x | 654.00x | 876.20x | 98.00x | 9.00x |
| ↳ /JS | 0.36x | 4.45x | 5.96x | 0.67x | 0.06x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.14 μs | ~ 0.30 μs | ~ 0.50 μs | ~ 0.07 μs | ~ 0.03 μs |
| ↳ /C | — | 300.00x | 500.00x | 70.00x | 30.00x |
| ↳ /JS | 3.62x | 7.50x | 12.50x | 1.75x | 0.75x |
| **Total Stable** | <strong>~ 10.25 ms</strong> | <strong>~ 50.55 ms</strong> | <strong>~ 1464.74 ms</strong> | <strong>~ 88.99 ms</strong> | <strong>~ 11.34 ms</strong> |
| ↳ /C | <strong>1.07x</strong> | <strong>5.26x</strong> | <strong>152.42x</strong> | <strong>9.26x</strong> | <strong>1.18x</strong> |
| ↳ μ/C | <strong>≈ 34.29x</strong> (n=9/14) | <strong>≈ 153.72x</strong> (n=9/14) | <strong>≈ 981.23x</strong> (n=9/14) | <strong>≈ 78.82x</strong> (n=9/14) | <strong>≈ 3.48x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 2.54x</strong> (n=9/14) | <strong>≈ 10.37x</strong> (n=9/14) | <strong>≈ 183.25x</strong> (n=9/14) | <strong>≈ 3.24x</strong> (n=9/14) | <strong>≈ 2.06x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 6.55x</strong> (n=9/14) | <strong>≈ 24.85x</strong> (n=9/14) | <strong>≈ 275.36x</strong> (n=9/14) | <strong>≈ 10.22x</strong> (n=9/14) | <strong>≈ 2.26x</strong> (n=9/14) |
| ↳ /JS | <strong>0.08x</strong> | <strong>0.39x</strong> | <strong>11.35x</strong> | <strong>0.69x</strong> | <strong>0.09x</strong> |
| ↳ μ/JS | <strong>0.76x</strong> | <strong>≈ 2.07x</strong> | <strong>18.29x</strong> | <strong>1.06x</strong> | <strong>0.25x</strong> |
| ↳ Q₂/JS | <strong>0.34x</strong> | <strong>≈ 1.81x</strong> | <strong>14.10x</strong> | <strong>0.89x</strong> | <strong>0.10x</strong> |
| ↳ G/JS | <strong>0.11x</strong> | <strong>≈ 1.34x</strong> | <strong>12.76x</strong> | <strong>0.81x</strong> | <strong>0.06x</strong> |
| **Incl WIP** | | | | | |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 131.85 μs | ~ 1452.41 μs | ~ 2146.41 μs | ~ 77.41 μs | ~ 6.61 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.08x | 11.91x | 17.60x | 0.63x | 0.05x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 76.49 μs | ~ 828.92 μs | ~ 1330.39 μs | ~ 67.38 μs | ~ 6.85 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.60x | 6.50x | 10.43x | 0.53x | 0.05x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 4.95 μs | ~ 281.15 μs | ~ 133.05 μs | ~ 49.77 μs | ~ 2.31 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.07x | 4.00x | 1.89x | 0.71x | 0.03x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 154.75 μs | ~ 12300.38 μs | ~ 11626.08 μs | ~ 825.35 μs | ~ 10.36 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.28x | 21.90x | 20.70x | 1.47x | 0.02x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 16.07 μs | ~ 752.15 μs | ~ 98.90 μs | ~ 48.10 μs | ~ 2.53 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.31x | 14.57x | 1.92x | 0.93x | 0.05x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.56 μs | ~ 304.51 μs | ~ 372.03 μs | ~ 13.95 μs | ~ 1.25 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.01x | 10.08x | 12.32x | 0.46x | 0.04x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 17.47 μs | ~ 464.65 μs | ~ 476.59 μs | ~ 18.97 μs | ~ 0.87 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.22x | 5.76x | 5.91x | 0.24x | 0.01x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 3.85 μs | ~ 298.48 μs | ~ 303.08 μs | ~ 11.64 μs | ~ 0.74 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.06x | 4.37x | 4.43x | 0.17x | 0.01x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 3.80 μs | ~ 297.10 μs | ~ 307.26 μs | ~ 11.44 μs | ~ 0.74 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.06x | 4.37x | 4.52x | 0.17x | 0.01x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.01 μs | ~ 17.05 μs | ~ 3.47 μs | ~ 3.45 μs | ~ 3.45 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.99x | 1687.77x | 343.62x | 341.43x | 341.50x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 6.71 μs | ~ 79.96 μs | ~ 62.44 μs | ~ 7.20 μs | ~ 3.41 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.18x | 2.14x | 1.67x | 0.19x | 0.09x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.41 μs | ~ 130.75 μs | ~ 2.71 μs | ~ 0.17 μs | ~ 0.17 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 2.16x | 685.42x | 14.19x | 0.90x | 0.88x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 2.51 μs | ~ 140.42 μs | ~ 2.47 μs | ~ 2.31 μs | ~ 2.32 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.04x | 2.20x | 0.04x | 0.04x | 0.04x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 44.48 μs | ~ 105.92 μs | ~ 104.16 μs | ~ 146.89 μs | ~ 3.66 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.98x | 2.34x | 2.30x | 3.24x | 0.08x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 133.98 μs | ~ 318.05 μs | ~ 300.93 μs | ~ 32.68 μs | ~ 9.75 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.92x | 2.19x | 2.07x | 0.22x | 0.07x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 56.96 μs | ~ 216.07 μs | ~ 167.98 μs | ~ 74.43 μs | ~ 3.71 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.19x | 4.53x | 3.52x | 1.56x | 0.08x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.53 μs | ~ 5.46 μs | ~ 144.74 μs | ~ 0.72 μs | ~ 0.24 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.22x | 2.24x | 59.23x | 0.29x | 0.10x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 5.69 μs | ~ 13.30 μs | ~ 190.03 μs | ~ 7.20 μs | ~ 0.26 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.25x | 2.92x | 41.74x | 1.58x | 0.06x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 10.04 μs | ~ 22.40 μs | ~ 679.42 μs | ~ 12.85 μs | ~ 1.03 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.43x | 3.19x | 96.67x | 1.83x | 0.15x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 7.15 μs | ~ 1019.37 μs | ~ 3085.81 μs | ~ 2.25 μs | ~ 2.25 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.02x | 2.20x | 6.65x | < 0.01x | < 0.01x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 20.67 μs | ~ 1073.35 μs | ~ 4068.22 μs | ~ 2.33 μs | ~ 2.38 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.04x | 2.09x | 7.93x | < 0.01x | < 0.01x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 190.79 μs | ~ 2199.81 μs | ~ 3135.15 μs | ~ 2.26 μs | ~ 2.26 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 1.23x | 14.13x | 20.14x | 0.01x | 0.01x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 192.99 μs | ~ 4998.00 μs | ~ 4026.20 μs | ~ 2.36 μs | ~ 2.36 μs |
| ↳ /C | — | — | — | — | — |
| ↳ /JS | 0.91x | 23.59x | 19.01x | 0.01x | 0.01x |
| **Total Incl WIP** | <strong>~ 1.11 ms</strong> | <strong>~ 27.32 ms</strong> | <strong>~ 32.77 ms</strong> | <strong>~ 1.42 ms</strong> | <strong>~ 0.07 ms</strong> |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | <strong>0.39x</strong> | <strong>9.49x</strong> | <strong>11.38x</strong> | <strong>0.49x</strong> | <strong>0.02x</strong> |
| ↳ μ/JS | <strong>0.66x</strong> | <strong>109.58x</strong> | <strong>30.37x</strong> | <strong>15.51x</strong> | <strong>14.93x</strong> |
| ↳ Q₂/JS | <strong>0.60x</strong> | <strong>4.37x</strong> | <strong>7.93x</strong> | <strong>0.46x</strong> | <strong>0.05x</strong> |
| ↳ G/JS | <strong>0.33x</strong> | <strong>7.90x</strong> | <strong>7.95x</strong> | <strong>0.33x</strong> | <strong>0.05x</strong> |
| **Total Stable + Incl WIP** | <strong>~ 11.36 ms</strong> | <strong>~ 77.87 ms</strong> | <strong>~ 1497.51 ms</strong> | <strong>~ 90.41 ms</strong> | <strong>~ 11.41 ms</strong> |
| ↳ /C | — | — | — | — | — |
| ↳ μ/C | — | — | — | — | — |
| ↳ Q₂/C | — | — | — | — | — |
| ↳ G/C | — | — | — | — | — |
| ↳ /JS | <strong>0.09x</strong> | <strong>0.59x</strong> | <strong>11.35x</strong> | <strong>0.69x</strong> | <strong>0.09x</strong> |
| ↳ μ/JS | <strong>0.70x</strong> | <strong>≈ 68.90x</strong> | <strong>25.80x</strong> | <strong>10.04x</strong> | <strong>9.38x</strong> |
| ↳ Q₂/JS | <strong>0.36x</strong> | <strong>≈ 2.92x</strong> | <strong>10.43x</strong> | <strong>0.63x</strong> | <strong>0.05x</strong> |
| ↳ G/JS | <strong>0.22x</strong> | <strong>≈ 4.04x</strong> | <strong>9.51x</strong> | <strong>0.46x</strong> | <strong>0.05x</strong> |
| **Excl WIP** | | | | | |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 4630.79 μs | — | — | — | — |
| ↳ /C | 0.70x | — | — | — | — |
| ↳ /JS | 0.18x | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 20534.73 μs | — | — | — | — |
| ↳ /C | 1.94x | — | — | — | — |
| ↳ /JS | 0.25x | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 2220.35 μs | — | — | — | — |
| ↳ /C | 3.38x | — | — | — | — |
| ↳ /JS | 0.24x | — | — | — | — |

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| **Stable** | | | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.07 μs | ~ 0.08 μs | ~ 0.07 μs |
| ↳ /C | 0.72x | 0.82x | 0.72x |
| ↳ /JS | 0.27x | 0.31x | 0.27x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs | ~ 0.15 μs | ~ 0.14 μs |
| ↳ /C | 1.81x | 1.81x | 1.69x |
| ↳ /JS | 0.48x | 0.48x | 0.45x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 6.35 μs | ~ 3.76 μs | ~ 2.84 μs |
| ↳ /C | 124.51x | 73.73x | 55.69x |
| ↳ /JS | 0.66x | 0.39x | 0.30x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 277.12 μs | ~ 322.34 μs | ~ 329.88 μs |
| ↳ /C | 8.08x | 9.40x | 9.62x |
| ↳ /JS | 1.46x | 1.70x | 1.74x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 228.44 μs | ~ 48.07 μs | ~ 43.57 μs |
| ↳ /C | 66.70x | 14.04x | 12.72x |
| ↳ /JS | 3.08x | 0.65x | 0.59x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.99 μs | ~ 9.07 μs | ~ 9.14 μs |
| ↳ /C | 1.00x | 0.54x | 0.54x |
| ↳ /JS | 0.25x | 0.13x | 0.14x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 237.12 μs | ~ 235.53 μs | ~ 62.26 μs |
| ↳ /C | 237120.00x | 235530.00x | 62260.00x |
| ↳ /JS | 0.19x | 0.19x | 0.05x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 37.49 μs | ~ 22.68 μs | ~ 0.80 μs |
| ↳ /C | 34.36x | 20.79x | 0.73x |
| ↳ /JS | 0.59x | 0.36x | 0.01x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 22551 μs | ~ 15428 μs | ~ 11970 μs |
| ↳ /C | 2.36x | 1.62x | 1.25x |
| ↳ /JS | 0.25x | 0.17x | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 17812 μs | ~ 27373 μs | ~ 6730.50 μs |
| ↳ /C | 17812000.00x | 27373000.00x | 6730500.00x |
| ↳ /JS | 0.78x | 1.20x | 0.29x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 1.70 μs | ~ 2.02 μs | ~ 0.84 μs |
| ↳ /C | 1700.00x | 2020.00x | 840.00x |
| ↳ /JS | 0.04x | 0.04x | 0.02x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 2989.50 μs | ~ 1964 μs | ~ 0.61 μs |
| ↳ /C | 2989500.00x | 1964000.00x | 610.00x |
| ↳ /JS | 0.20x | 0.13x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.29 μs | ~ 5.04 μs | ~ 2.79 μs |
| ↳ /C | 145.80x | 100.80x | 55.80x |
| ↳ /JS | 0.99x | 0.69x | 0.38x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs | ~ 0.01 μs | ~ 0 μs |
| ↳ /C | — | 10.00x | — |
| ↳ /JS | — | 0.25x | — |
| **Total Stable** | <strong>~ 44.17 ms</strong> | <strong>~ 45.41 ms</strong> | <strong>~ 19.15 ms</strong> |
| ↳ /C | <strong>4.60x</strong> | <strong>4.73x</strong> | <strong>1.99x</strong> |
| ↳ μ/C | <strong>≈ 43.12x</strong> (n=9/14) | <strong>≈ 25.00x</strong> (n=9/14) | <strong>≈ 15.55x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 8.08x</strong> (n=9/14) | <strong>≈ 9.40x</strong> (n=9/14) | <strong>≈ 1.80x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 10.12x</strong> (n=9/14) | <strong>≈ 6.64x</strong> (n=9/14) | <strong>≈ 3.96x</strong> (n=9/14) |
| ↳ /JS | <strong>0.34x</strong> | <strong>0.35x</strong> | <strong>0.15x</strong> |
| ↳ μ/JS | <strong>0.71x</strong> (n=13/14) | <strong>0.48x</strong> | <strong>0.34x</strong> (n=13/14) |
| ↳ Q₂/JS | <strong>0.48x</strong> (n=13/14) | <strong>0.35x</strong> | <strong>0.27x</strong> (n=13/14) |
| ↳ G/JS | <strong>0.43x</strong> (n=13/14) | <strong>0.33x</strong> | <strong>0.09x</strong> (n=13/14) |
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
| ↳ /C | 1.13x | 1.13x | 1.03x |
| ↳ /JS | 0.42x | 0.42x | 0.38x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.21 μs | ~ 0.23 μs | ~ 0.24 μs |
| ↳ /C | 2.53x | 2.77x | 2.89x |
| ↳ /JS | 0.68x | 0.74x | 0.77x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 37.32 μs | ~ 5.22 μs | ~ 0.86 μs |
| ↳ /C | 731.76x | 102.35x | 16.86x |
| ↳ /JS | 3.89x | 0.54x | 0.09x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 1341.30 μs | ~ 103.95 μs | ~ 104.62 μs |
| ↳ /C | 39.12x | 3.03x | 3.05x |
| ↳ /JS | 7.08x | 0.55x | 0.55x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 404.77 μs | ~ 1276.40 μs | ~ 15.46 μs |
| ↳ /C | 118.18x | 372.67x | 4.51x |
| ↳ /JS | 5.46x | 17.21x | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.24 μs | ~ 18.81 μs | ~ 19.68 μs |
| ↳ /C | 1.14x | 1.11x | 1.16x |
| ↳ /JS | 0.29x | 0.28x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 520.08 μs | ~ 418.19 μs | ~ 0.01 μs |
| ↳ /C | 520080.00x | 418190.00x | 10.00x |
| ↳ /JS | 0.41x | 0.33x | < 0.01x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 148.04 μs | ~ 30.95 μs | ~ 36.68 μs |
| ↳ /C | 135.69x | 28.37x | 33.62x |
| ↳ /JS | 2.34x | 0.49x | 0.58x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16911.71 μs | ~ 14282.54 μs | ~ 14426.50 μs |
| ↳ /C | 1.77x | 1.50x | 1.51x |
| ↳ /JS | 0.19x | 0.16x | 0.16x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 66666.38 μs | ~ 52475.29 μs | ~ 0.01 μs |
| ↳ /C | 66666380.00x | 52475290.00x | 10.00x |
| ↳ /JS | 2.92x | 2.30x | < 0.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 20.41 μs | ~ 14.55 μs | ~ 7.61 μs |
| ↳ /C | 20410.00x | 14550.00x | 7610.00x |
| ↳ /JS | 0.42x | 0.30x | 0.16x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8739.06 μs | ~ 6780.02 μs | ~ 0.01 μs |
| ↳ /C | 8739060.00x | 6780020.00x | 10.00x |
| ↳ /JS | 0.60x | 0.46x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 27.82 μs | ~ 16.21 μs | ~ 1.04 μs |
| ↳ /C | 556.40x | 324.20x | 20.80x |
| ↳ /JS | 3.79x | 2.21x | 0.14x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.04 μs | ~ 0.04 μs | ~ 0.01 μs |
| ↳ /C | 40.00x | 40.00x | 10.00x |
| ↳ /JS | 1.00x | 1.00x | 0.25x |
| **Total Stable** | <strong>~ 94.84 ms</strong> | <strong>~ 75.42 ms</strong> | <strong>~ 14.61 ms</strong> |
| ↳ /C | <strong>9.87x</strong> | <strong>7.85x</strong> | <strong>1.52x</strong> |
| ↳ μ/C | <strong>≈ 178.09x</strong> (n=9/14) | <strong>≈ 93.32x</strong> (n=9/14) | <strong>≈ 9.55x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 39.13x</strong> (n=9/14) | <strong>≈ 3.03x</strong> (n=9/14) | <strong>≈ 3.05x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 22.57x</strong> (n=9/14) | <strong>≈ 12.20x</strong> (n=9/14) | <strong>≈ 4.58x</strong> (n=9/14) |
| ↳ /JS | <strong>0.73x</strong> | <strong>0.58x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>2.10x</strong> | <strong>1.93x</strong> | <strong>0.26x</strong> |
| ↳ Q₂/JS | <strong>0.82x</strong> | <strong>0.52x</strong> | <strong>0.18x</strong> |
| ↳ G/JS | <strong>1.14x</strong> | <strong>0.71x</strong> | <strong>0.02x</strong> |
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
| ↳ /C | 21.96x | 32.89x | 21.75x |
| ↳ /JS | 8.19x | 12.27x | 8.12x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 1.08 μs | ~ 2.26 μs | ~ 1.27 μs |
| ↳ /C | 13.01x | 27.23x | 15.30x |
| ↳ /JS | 3.48x | 7.29x | 4.10x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 137.53 μs | ~ 174.47 μs | ~ 0.63 μs |
| ↳ /C | 2696.67x | 3420.98x | 12.35x |
| ↳ /JS | 14.34x | 18.19x | 0.07x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 72.18 μs | ~ 49.87 μs | ~ 50.30 μs |
| ↳ /C | 2.11x | 1.45x | 1.47x |
| ↳ /JS | 0.38x | 0.26x | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1413.71 μs | ~ 1676.49 μs | ~ 491.74 μs |
| ↳ /C | 412.76x | 489.49x | 143.57x |
| ↳ /JS | 19.06x | 22.60x | 6.63x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 45.89 μs | ~ 100.16 μs | ~ 52.48 μs |
| ↳ /C | 2.71x | 5.92x | 3.10x |
| ↳ /JS | 0.68x | 1.48x | 0.78x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1941.41 μs | ~ 6824.38 μs | ~ 32.84 μs |
| ↳ /C | 1941410.00x | 6824380.00x | 32840.00x |
| ↳ /JS | 1.54x | 5.40x | 0.03x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 698.16 μs | ~ 1022.38 μs | ~ 1.89 μs |
| ↳ /C | 639.93x | 937.10x | 1.73x |
| ↳ /JS | 11.01x | 16.13x | 0.03x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 53877.92 μs | ~ 509843.88 μs | ~ 113830.46 μs |
| ↳ /C | 5.64x | 53.39x | 11.92x |
| ↳ /JS | 0.60x | 5.67x | 1.27x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 6237.63 μs | ~ 432965.96 μs | ~ 87761.96 μs |
| ↳ /C | 6237630.00x | 432965960.00x | 87761960.00x |
| ↳ /JS | 0.27x | 18.96x | 3.84x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 467.12 μs | ~ 560.06 μs | ~ 0.48 μs |
| ↳ /C | 467120.00x | 560060.00x | 480.00x |
| ↳ /JS | 9.69x | 11.61x | 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 791.27 μs | ~ 111183.96 μs | ~ 326.09 μs |
| ↳ /C | 791270.00x | 111183960.00x | 326090.00x |
| ↳ /JS | 0.05x | 7.61x | 0.02x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 89.94 μs | ~ 26.94 μs | ~ 0.80 μs |
| ↳ /C | 1798.80x | 538.80x | 16.00x |
| ↳ /JS | 12.24x | 3.67x | 0.11x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.84 μs | ~ 0.85 μs | ~ 0.08 μs |
| ↳ /C | 840.00x | 850.00x | 80.00x |
| ↳ /JS | 21.00x | 21.25x | 2.00x |
| **Total Stable** | <strong>~ 65.78 ms</strong> | <strong>~ 1064.43 ms</strong> | <strong>~ 202.55 ms</strong> |
| ↳ /C | <strong>6.84x</strong> | <strong>110.76x</strong> | <strong>21.08x</strong> |
| ↳ μ/C | <strong>≈ 627.62x</strong> (n=9/14) | <strong>≈ 619.70x</strong> (n=9/14) | <strong>≈ 25.29x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 21.30x</strong> (n=9/14) | <strong>≈ 53.39x</strong> (n=9/14) | <strong>≈ 12.60x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 61.26x</strong> (n=9/14) | <strong>≈ 89.22x</strong> (n=9/14) | <strong>≈ 9.90x</strong> (n=9/14) |
| ↳ /JS | <strong>0.51x</strong> | <strong>8.25x</strong> | <strong>1.57x</strong> |
| ↳ μ/JS | <strong>≈ 7.27x</strong> | <strong>≈ 10.83x</strong> | <strong>≈ 1.94x</strong> |
| ↳ Q₂/JS | <strong>≈ 5.84x</strong> | <strong>≈ 9.61x</strong> | <strong>≈ 0.52x</strong> |
| ↳ G/JS | <strong>≈ 2.62x</strong> | <strong>≈ 7.14x</strong> | <strong>≈ 0.36x</strong> |
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
| ↳ /C | 2.43x | 37.94x | 2.37x | 2.27x | 2.37x |
| ↳ /JS | 0.91x | 14.15x | 0.88x | 0.85x | 0.88x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.09 μs | ~ 11.78 μs | ~ 0.09 μs | ~ 0.12 μs | ~ 0.10 μs |
| ↳ /C | 1.05x | 141.93x | 1.08x | 1.45x | 1.20x |
| ↳ /JS | 0.28x | 38.00x | 0.29x | 0.39x | 0.32x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.11 μs | ~ 459.33 μs | ~ 17.07 μs | ~ 11.05 μs | ~ 0.51 μs |
| ↳ /C | 2.24x | 9006.47x | 334.71x | 216.67x | 10.00x |
| ↳ /JS | 0.01x | 47.90x | 1.78x | 1.15x | 0.05x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 29.88 μs | ~ 23547.50 μs | ~ 10.93 μs | ~ 34.70 μs | ~ 34.63 μs |
| ↳ /C | 0.87x | 686.86x | 0.32x | 1.01x | 1.01x |
| ↳ /JS | 0.16x | 124.27x | 0.06x | 0.18x | 0.18x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.14 μs | ~ 8728.98 μs | ~ 114.96 μs | ~ 107.97 μs | ~ 3.47 μs |
| ↳ /C | 0.92x | 2548.61x | 33.56x | 31.52x | 1.01x |
| ↳ /JS | 0.04x | 117.67x | 1.55x | 1.46x | 0.05x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.45 μs | ~ 625.00 μs | ~ 15.60 μs | ~ 18.59 μs | ~ 18.36 μs |
| ↳ /C | 0.97x | 36.96x | 0.92x | 1.10x | 1.09x |
| ↳ /JS | 0.24x | 9.26x | 0.23x | 0.28x | 0.27x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 159.52 μs | ~ 11792.17 μs | ~ 2249.51 μs | ~ 1033.78 μs | ~ 0.01 μs |
| ↳ /C | 159524.00x | 11792170.00x | 2249510.00x | 1033780.00x | 10.00x |
| ↳ /JS | 0.13x | 9.33x | 1.78x | 0.82x | < 0.01x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 144.00 μs | ~ 2194.96 μs | ~ 120.43 μs | ~ 78.88 μs | ~ 0.78 μs |
| ↳ /C | 131.99x | 2011.88x | 110.38x | 72.30x | 0.71x |
| ↳ /JS | 2.27x | 34.63x | 1.90x | 1.24x | 0.01x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8136.98 μs | ~ 659207.04 μs | ~ 50862.42 μs | ~ 28737.38 μs | ~ 28247.96 μs |
| ↳ /C | 0.85x | 69.03x | 5.33x | 3.01x | 2.96x |
| ↳ /JS | 0.09x | 7.33x | 0.57x | 0.32x | 0.31x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.03 μs | ~ 1588595.46 μs | ~ 7471.33 μs | ~ 0.01 μs | ~ 0.01 μs |
| ↳ /C | 32.00x | 1588595460.00x | 7471330.00x | 10.00x | 10.00x |
| ↳ /JS | < 0.01x | 69.58x | 0.33x | < 0.01x | < 0.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 44.57 μs | ~ 826.12 μs | ~ 73.12 μs | ~ 27.36 μs | ~ 0.01 μs |
| ↳ /C | 44568.00x | 826120.00x | 73120.00x | 27360.00x | 10.00x |
| ↳ /JS | 0.92x | 17.13x | 1.52x | 0.57x | < 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.01 μs | ~ 159056.38 μs | ~ 38342.83 μs | ~ 23675.83 μs | ~ 0.01 μs |
| ↳ /C | 11.00x | 159056380.00x | 38342830.00x | 23675830.00x | 10.00x |
| ↳ /JS | < 0.01x | 10.89x | 2.62x | 1.62x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.03 μs | ~ 298.70 μs | ~ 3.28 μs | ~ 0.58 μs | ~ 0.25 μs |
| ↳ /C | 0.68x | 5974.00x | 65.60x | 11.60x | 5.00x |
| ↳ /JS | < 0.01x | 40.64x | 0.45x | 0.08x | 0.03x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.07 μs | ~ 0.95 μs | ~ 0.10 μs | ~ 0.02 μs | ~ 0.01 μs |
| ↳ /C | 68.00x | 950.00x | 100.00x | 20.00x | 10.00x |
| ↳ /JS | 1.70x | 23.75x | 2.50x | 0.50x | 0.25x |
| **Total Stable** | <strong>~ 8.54 ms</strong> | <strong>~ 2455.35 ms</strong> | <strong>~ 99.28 ms</strong> | <strong>~ 53.73 ms</strong> | <strong>~ 28.31 ms</strong> |
| ↳ /C | <strong>0.89x</strong> | <strong>255.50x</strong> | <strong>10.33x</strong> | <strong>5.59x</strong> | <strong>2.95x</strong> |
| ↳ μ/C | <strong>≈ 15.78x</strong> (n=9/14) | <strong>≈ 2300.40x</strong> (n=9/14) | <strong>≈ 62.34x</strong> (n=9/14) | <strong>≈ 38.37x</strong> (n=9/14) | <strong>≈ 2.84x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 0.97x</strong> (n=9/14) | <strong>≈ 686.92x</strong> (n=9/14) | <strong>≈ 5.33x</strong> (n=9/14) | <strong>≈ 3.01x</strong> (n=9/14) | <strong>≈ 1.25x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 1.89x</strong> (n=9/14) | <strong>≈ 517.86x</strong> (n=9/14) | <strong>≈ 8.86x</strong> (n=9/14) | <strong>≈ 7.38x</strong> (n=9/14) | <strong>≈ 1.91x</strong> (n=9/14) |
| ↳ /JS | <strong>0.07x</strong> | <strong>19.02x</strong> | <strong>0.77x</strong> | <strong>0.42x</strong> | <strong>0.22x</strong> |
| ↳ μ/JS | <strong>≈ 0.48x</strong> | <strong>≈ 40.26x</strong> | <strong>≈ 1.17x</strong> | <strong>≈ 0.67x</strong> | <strong>≈ 0.17x</strong> |
| ↳ Q₂/JS | <strong>≈ 0.14x</strong> | <strong>≈ 28.74x</strong> | <strong>≈ 1.20x</strong> | <strong>≈ 0.52x</strong> | <strong>≈ 0.05x</strong> |
| ↳ G/JS | <strong>≈ 0.03x</strong> | <strong>≈ 26.82x</strong> | <strong>≈ 0.77x</strong> | <strong>≈ 0.20x</strong> | <strong>≈ 0.01x</strong> |
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
| ↳ /C | 27.94x | 9.69x | 8.76x |
| ↳ /JS | 10.42x | 3.62x | 3.27x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs | ~ 0.11 μs | ~ 0.02 μs |
| ↳ /C | 3.73x | 1.33x | 0.24x |
| ↳ /JS | 1.00x | 0.35x | 0.06x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 154.89 μs | ~ 35.16 μs | ~ 0.36 μs |
| ↳ /C | 3037.06x | 689.41x | 7.06x |
| ↳ /JS | 16.15x | 3.67x | 0.04x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 2844.59 μs | ~ 35.11 μs | ~ 33.85 μs |
| ↳ /C | 82.97x | 1.02x | 0.99x |
| ↳ /JS | 15.01x | 0.19x | 0.18x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1453.15 μs | ~ 567.68 μs | ~ 3.41 μs |
| ↳ /C | 424.28x | 165.75x | 1.00x |
| ↳ /JS | 19.59x | 7.65x | 0.05x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 221.50 μs | ~ 17.99 μs | ~ 16.54 μs |
| ↳ /C | 13.10x | 1.06x | 0.98x |
| ↳ /JS | 3.28x | 0.27x | 0.25x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 2630.26 μs | ~ 4240.89 μs | ~ 0.02 μs |
| ↳ /C | 2630260.00x | 4240890.00x | 20.00x |
| ↳ /JS | 2.08x | 3.36x | < 0.01x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 974.45 μs | ~ 257.70 μs | ~ 1.08 μs |
| ↳ /C | 893.17x | 236.21x | 0.99x |
| ↳ /JS | 15.37x | 4.07x | 0.02x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 613388.67 μs | ~ 66049.46 μs | ~ 22007.12 μs |
| ↳ /C | 64.23x | 6.92x | 2.30x |
| ↳ /JS | 6.82x | 0.73x | 0.24x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 265261 μs | ~ 27358.79 μs | ~ 0.02 μs |
| ↳ /C | 265261000.00x | 27358790.00x | 20.00x |
| ↳ /JS | 11.62x | 1.20x | < 0.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 293.30 μs | ~ 185.70 μs | ~ 0.02 μs |
| ↳ /C | 293300.00x | 185700.00x | 20.00x |
| ↳ /JS | 6.08x | 3.85x | < 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 36615.58 μs | ~ 31474.58 μs | ~ 0.02 μs |
| ↳ /C | 36615580.00x | 31474580.00x | 20.00x |
| ↳ /JS | 2.51x | 2.15x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 36.57 μs | ~ 1.10 μs | ~ 0.74 μs |
| ↳ /C | 731.40x | 22.00x | 14.80x |
| ↳ /JS | 4.98x | 0.15x | 0.10x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.48 μs | ~ 0.30 μs | ~ 0.02 μs |
| ↳ /C | 480.00x | 300.00x | 20.00x |
| ↳ /JS | 12.00x | 7.50x | 0.50x |
| **Total Stable** | <strong>~ 923.88 ms</strong> | <strong>~ 130.23 ms</strong> | <strong>~ 22.06 ms</strong> |
| ↳ /C | <strong>96.14x</strong> | <strong>13.55x</strong> | <strong>2.30x</strong> |
| ↳ μ/C | <strong>≈ 593.28x</strong> (n=9/14) | <strong>≈ 127.48x</strong> (n=9/14) | <strong>≈ 4.11x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 82.98x</strong> (n=9/14) | <strong>≈ 9.39x</strong> (n=9/14) | <strong>≈ 1.00x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 122.94x</strong> (n=9/14) | <strong>≈ 15.68x</strong> (n=9/14) | <strong>≈ 2.04x</strong> (n=9/14) |
| ↳ /JS | <strong>7.16x</strong> | <strong>1.01x</strong> | <strong>0.17x</strong> |
| ↳ μ/JS | <strong>9.04x</strong> | <strong>2.74x</strong> | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>8.63x</strong> | <strong>2.76x</strong> | <strong>0.06x</strong> |
| ↳ G/JS | <strong>6.70x</strong> | <strong>1.45x</strong> | <strong>0.01x</strong> |
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
| ↳ /C | 11.55x | 5.15x | 5.15x |
| ↳ /JS | 4.31x | 1.92x | 1.92x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.36 μs | ~ 0.12 μs | ~ 0.11 μs |
| ↳ /C | 4.34x | 1.45x | 1.33x |
| ↳ /JS | 1.16x | 0.39x | 0.35x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 560.47 μs | ~ 16.05 μs | ~ 3.05 μs |
| ↳ /C | 10989.61x | 314.71x | 59.80x |
| ↳ /JS | 58.44x | 1.67x | 0.32x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.45 μs | ~ 37.60 μs | ~ 49.49 μs |
| ↳ /C | 1.68x | 1.10x | 1.44x |
| ↳ /JS | 0.30x | 0.20x | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 2545.27 μs | ~ 87.33 μs | ~ 7.21 μs |
| ↳ /C | 743.14x | 25.50x | 2.11x |
| ↳ /JS | 34.31x | 1.18x | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.15 μs | ~ 7.70 μs | ~ 7.66 μs |
| ↳ /C | 0.90x | 0.46x | 0.45x |
| ↳ /JS | 0.22x | 0.11x | 0.11x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1678.49 μs | ~ 342.79 μs | ~ 25.24 μs |
| ↳ /C | 1678490.00x | 342790.00x | 25240.00x |
| ↳ /JS | 1.33x | 0.27x | 0.02x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 525.55 μs | ~ 45.26 μs | ~ 1.80 μs |
| ↳ /C | 481.71x | 41.48x | 1.65x |
| ↳ /JS | 8.29x | 0.71x | 0.03x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 19071.54 μs | ~ 36659.92 μs | ~ 9499.96 μs |
| ↳ /C | 2.00x | 3.84x | 0.99x |
| ↳ /JS | 0.21x | 0.41x | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2523.93 μs | ~ 23242.04 μs | ~ 2254.54 μs |
| ↳ /C | 2523930.00x | 23242040.00x | 2254540.00x |
| ↳ /JS | 0.11x | 1.02x | 0.10x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 169.78 μs | ~ 78.22 μs | ~ 3.21 μs |
| ↳ /C | 169780.00x | 78220.00x | 3210.00x |
| ↳ /JS | 3.52x | 1.62x | 0.07x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 6576.73 μs | ~ 9665.58 μs | ~ 235.04 μs |
| ↳ /C | 6576730.00x | 9665580.00x | 235040.00x |
| ↳ /JS | 0.45x | 0.66x | 0.02x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 62.07 μs | ~ 29.54 μs | ~ 15.63 μs |
| ↳ /C | 1241.40x | 590.80x | 312.60x |
| ↳ /JS | 8.44x | 4.02x | 2.13x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.24 μs | ~ 0.06 μs | ~ 0.02 μs |
| ↳ /C | 240.00x | 60.00x | 20.00x |
| ↳ /JS | 6.00x | 1.50x | 0.50x |
| **Total Stable** | <strong>~ 33.79 ms</strong> | <strong>~ 70.21 ms</strong> | <strong>~ 12.10 ms</strong> |
| ↳ /C | <strong>3.52x</strong> | <strong>7.31x</strong> | <strong>1.26x</strong> |
| ↳ μ/C | <strong>≈ 1521.94x</strong> (n=9/14) | <strong>≈ 110.08x</strong> (n=9/14) | <strong>≈ 42.96x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 11.20x</strong> (n=9/14) | <strong>≈ 5.00x</strong> (n=9/14) | <strong>≈ 1.65x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 44.99x</strong> (n=9/14) | <strong>≈ 11.25x</strong> (n=9/14) | <strong>≈ 4.06x</strong> (n=9/14) |
| ↳ /JS | <strong>0.26x</strong> | <strong>0.54x</strong> | <strong>0.09x</strong> |
| ↳ μ/JS | <strong>≈ 9.06x</strong> | <strong>≈ 1.12x</strong> | <strong>≈ 0.43x</strong> |
| ↳ Q₂/JS | <strong>≈ 2.42x</strong> | <strong>≈ 0.87x</strong> | <strong>≈ 0.11x</strong> |
| ↳ G/JS | <strong>≈ 1.97x</strong> | <strong>≈ 0.75x</strong> | <strong>≈ 0.15x</strong> |
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
| ↳ /C | 1.44x | 1.13x | 1.03x |
| ↳ /JS | 0.54x | 0.42x | 0.38x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.26 μs | ~ 0.10 μs | ~ 0.08 μs |
| ↳ /C | 3.13x | 1.20x | 0.96x |
| ↳ /JS | 0.84x | 0.32x | 0.26x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 5.03 μs | ~ 4.18 μs | ~ 0.57 μs |
| ↳ /C | 98.63x | 81.96x | 11.18x |
| ↳ /JS | 0.52x | 0.44x | 0.06x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 45.64 μs | ~ 38.54 μs | ~ 39.83 μs |
| ↳ /C | 1.33x | 1.12x | 1.16x |
| ↳ /JS | 0.24x | 0.20x | 0.21x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 73.28 μs | ~ 38.80 μs | ~ 4.18 μs |
| ↳ /C | 21.40x | 11.33x | 1.22x |
| ↳ /JS | 0.99x | 0.52x | 0.06x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 18.18 μs | ~ 7.60 μs | ~ 5.80 μs |
| ↳ /C | 1.08x | 0.45x | 0.34x |
| ↳ /JS | 0.27x | 0.11x | 0.09x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 310.51 μs | ~ 375.72 μs | ~ 0 μs |
| ↳ /C | 310510.00x | 375720.00x | — |
| ↳ /JS | 0.25x | 0.30x | — |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.29 μs | ~ 18.40 μs | ~ 0.54 μs |
| ↳ /C | 19.51x | 16.87x | 0.49x |
| ↳ /JS | 0.34x | 0.29x | 0.01x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16788.38 μs | ~ 12065.75 μs | ~ 11828.17 μs |
| ↳ /C | 1.76x | 1.26x | 1.24x |
| ↳ /JS | 0.19x | 0.13x | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.01 μs | ~ 14583.08 μs | ~ 0 μs |
| ↳ /C | 10.00x | 14583080.00x | — |
| ↳ /JS | < 0.01x | 0.64x | — |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 6.26 μs | ~ 17.25 μs | ~ 0 μs |
| ↳ /C | 6260.00x | 17250.00x | — |
| ↳ /JS | 0.13x | 0.36x | — |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5.24 μs | ~ 8062 μs | ~ 0.04 μs |
| ↳ /C | 5240.00x | 8062000.00x | 40.00x |
| ↳ /JS | < 0.01x | 0.55x | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.96 μs | ~ 3.06 μs | ~ 0.50 μs |
| ↳ /C | 59.20x | 61.20x | 10.00x |
| ↳ /JS | 0.40x | 0.42x | 0.07x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.01 μs | ~ 0.01 μs | ~ 0 μs |
| ↳ /C | 10.00x | 10.00x | — |
| ↳ /JS | 0.25x | 0.25x | — |
| **Total Stable** | <strong>~ 17.28 ms</strong> | <strong>~ 35.21 ms</strong> | <strong>~ 11.88 ms</strong> |
| ↳ /C | <strong>1.80x</strong> | <strong>3.66x</strong> | <strong>1.24x</strong> |
| ↳ μ/C | <strong>≈ 23.28x</strong> (n=9/14) | <strong>≈ 19.81x</strong> (n=9/14) | <strong>≈ 3.08x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 3.24x</strong> (n=9/14) | <strong>≈ 1.26x</strong> (n=9/14) | <strong>≈ 1.16x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 6.75x</strong> (n=9/14) | <strong>≈ 4.54x</strong> (n=9/14) | <strong>≈ 1.47x</strong> (n=9/14) |
| ↳ /JS | <strong>0.13x</strong> | <strong>0.27x</strong> | <strong>0.09x</strong> |
| ↳ μ/JS | <strong>0.36x</strong> | <strong>0.35x</strong> | <strong>0.13x</strong> (n=10/14) |
| ↳ Q₂/JS | <strong>0.29x</strong> | <strong>0.34x</strong> | <strong>0.08x</strong> (n=10/14) |
| ↳ G/JS | <strong>0.08x</strong> | <strong>0.31x</strong> | <strong>0.03x</strong> (n=10/14) |
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
| ↳ /C | 2.27x |
| ↳ /JS | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.45x |
| ↳ /JS | 0.39x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 4.28 μs |
| ↳ /C | 83.92x |
| ↳ /JS | 0.45x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 80.51 μs |
| ↳ /C | 2.35x |
| ↳ /JS | 0.42x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 158.94 μs |
| ↳ /C | 46.41x |
| ↳ /JS | 2.14x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.38 μs |
| ↳ /C | 1.15x |
| ↳ /JS | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 758.69 μs |
| ↳ /C | 758690.00x |
| ↳ /JS | 0.60x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 12.17 μs |
| ↳ /C | 11.15x |
| ↳ /JS | 0.19x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8746.50 μs |
| ↳ /C | 0.92x |
| ↳ /JS | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 35687 μs |
| ↳ /C | 35687000.00x |
| ↳ /JS | 1.56x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 5.21 μs |
| ↳ /C | 5210.00x |
| ↳ /JS | 0.11x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8982 μs |
| ↳ /C | 8982000.00x |
| ↳ /JS | 0.61x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 6.58 μs |
| ↳ /C | 131.60x |
| ↳ /JS | 0.90x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Stable** | <strong>~ 54.46 ms</strong> |
| ↳ /C | <strong>5.67x</strong> |
| ↳ μ/C | <strong>≈ 31.43x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 2.35x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 7.15x</strong> (n=9/14) |
| ↳ /JS | <strong>0.42x</strong> |
| ↳ μ/JS | <strong>0.66x</strong> (n=13/14) |
| ↳ Q₂/JS | <strong>0.45x</strong> (n=13/14) |
| ↳ G/JS | <strong>0.46x</strong> (n=13/14) |
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
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.06 μs |
| ↳ /C | 0.62x |
| ↳ /JS | 0.23x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.19 μs |
| ↳ /C | 2.29x |
| ↳ /JS | 0.61x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 3.58 μs |
| ↳ /C | 70.20x |
| ↳ /JS | 0.37x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 62.84 μs |
| ↳ /C | 1.83x |
| ↳ /JS | 0.33x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.68 μs |
| ↳ /C | 2.24x |
| ↳ /JS | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 6.75 μs |
| ↳ /C | 0.40x |
| ↳ /JS | 0.10x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 143.90 μs |
| ↳ /C | 143900.00x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.58 μs |
| ↳ /C | 19.78x |
| ↳ /JS | 0.34x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 11559 μs |
| ↳ /C | 1.21x |
| ↳ /JS | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 4883 μs |
| ↳ /C | 4883000.00x |
| ↳ /JS | 0.21x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.02 μs |
| ↳ /C | 20.00x |
| ↳ /JS | < 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.46 μs |
| ↳ /C | 460.00x |
| ↳ /JS | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 4.33 μs |
| ↳ /C | 86.60x |
| ↳ /JS | 0.59x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Stable** | <strong>~ 16.69 ms</strong> |
| ↳ /C | <strong>1.74x</strong> |
| ↳ μ/C | <strong>≈ 20.73x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 2.25x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 4.10x</strong> (n=9/14) |
| ↳ /JS | <strong>0.13x</strong> |
| ↳ μ/JS | <strong>0.24x</strong> (n=13/14) |
| ↳ Q₂/JS | <strong>0.21x</strong> (n=13/14) |
| ↳ G/JS | <strong>0.07x</strong> (n=13/14) |
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
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.05 μs |
| ↳ /C | 0.52x |
| ↳ /JS | 0.19x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.45x |
| ↳ /JS | 0.39x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 1.98 μs |
| ↳ /C | 38.82x |
| ↳ /JS | 0.21x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 49.82 μs |
| ↳ /C | 1.45x |
| ↳ /JS | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 15.85 μs |
| ↳ /C | 4.63x |
| ↳ /JS | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.05 μs |
| ↳ /C | 0.95x |
| ↳ /JS | 0.24x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 140.38 μs |
| ↳ /C | 140380.00x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 17.94 μs |
| ↳ /C | 16.44x |
| ↳ /JS | 0.28x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10787 μs |
| ↳ /C | 1.13x |
| ↳ /JS | 0.12x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 12448 μs |
| ↳ /C | 12448000.00x |
| ↳ /JS | 0.55x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 10.03 μs |
| ↳ /C | 10030.00x |
| ↳ /JS | 0.21x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5690 μs |
| ↳ /C | 5690000.00x |
| ↳ /JS | 0.39x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 5.89 μs |
| ↳ /C | 117.80x |
| ↳ /JS | 0.80x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs |
| ↳ /C | — |
| ↳ /JS | — |
| **Total Stable** | <strong>~ 29.18 ms</strong> |
| ↳ /C | <strong>3.04x</strong> |
| ↳ μ/C | <strong>≈ 20.45x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 1.55x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 4.22x</strong> (n=9/14) |
| ↳ /JS | <strong>0.23x</strong> |
| ↳ μ/JS | <strong>0.31x</strong> (n=13/14) |
| ↳ Q₂/JS | <strong>0.24x</strong> (n=13/14) |
| ↳ G/JS | <strong>0.26x</strong> (n=13/14) |
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

C Benchmark             | Hand-written imperative C<br>↓<br>clang -O3<br><br>(WIP) |
----------------------- | -------------------- |
| **Stable** | |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.10 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.37x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.08 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.27x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.05 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.01x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 34.28 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.18x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.42 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.05x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.91 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.25x |
| [Church Numerals](src/Test/Church/Church.purs) | < 0.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | < 0.01x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 1.09 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.02x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 9550.00 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | < 0.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | < 0.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | < 0.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | < 0.01x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | < 0.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | < 0.01x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.05 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.01x |
| [RowToList](src/Test/RowToList/RowToList.purs) | < 0.01 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.03x |
| **Total Stable** | <strong>~ 9.61 ms</strong> |
| ↳ /C | <strong>1.00x</strong> |
| ↳ μ/C | <strong>≈ 1.00x</strong> (n=9/14) |
| ↳ Q₂/C | <strong>≈ 1.00x</strong> (n=9/14) |
| ↳ G/C | <strong>≈ 1.00x</strong> (n=9/14) |
| ↳ /JS | <strong>0.07x</strong> |
| ↳ μ/JS | <strong>≈ 0.14x</strong> (n=9/14) |
| ↳ Q₂/JS | <strong>≈ 0.11x</strong> (n=9/14) |
| ↳ G/JS | <strong>≈ 0.06x</strong> (n=9/14) |
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
| ↳ /JS | 0.07x |

### Extended benchmark results (I/O, mutability, async)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
[File I/O](srx/Test/FileOps/FileOps.purs) | ~ 466899.83 μs<br>(/C = —) | ~ 477117.79 μs<br>(/C = —)<br>(/JS = 1.02x) | ~ 475960.00 μs<br>(/C = —)<br>(/JS = 1.02x) | ~ 542377.75 μs<br>(/C = —)<br>(/JS = 1.16x) |
[STArray Operations](srx/Test/STArray.purs) | ~ 0.92 μs<br>(/C = —) | ~ 0.54 μs<br>(/C = —)<br>(/JS = 0.59x) | ~ 0.29 μs<br>(/C = —)<br>(/JS = 0.32x) | ~ 0.29 μs<br>(/C = —)<br>(/JS = 0.32x) |
[String Operations](srx/Test/StringOps.purs) | ~ 256.71 μs<br>(/C = —) | ~ 232.50 μs<br>(/C = —)<br>(/JS = 0.91x) | ~ 559.04 μs<br>(/C = —)<br>(/JS = 2.18x) | ~ 414.58 μs<br>(/C = —)<br>(/JS = 1.61x) |
[Aff Operations](srx/Test/AffOperations.purs) | ~ 10125.21 μs<br>(/C = —) | ~ 11581.83 μs<br>(/C = —)<br>(/JS = 1.14x) | ~ 10500.21 μs<br>(/C = —)<br>(/JS = 1.04x) | ~ 10903.21 μs<br>(/C = —)<br>(/JS = 1.08x) |
[Parallelism](srx/Test/Parallelism.purs) | ~ 14275052.04 μs<br>(/C = —) | ~ 14278260.46 μs<br>(/C = —) | ~ 965397.12 μs<br>(/C = —)<br>(/JS = 0.07x) | ~ 478303.67 μs<br>(/C = —)<br>(/JS = 0.03x) |
**Total Execution Time** | <strong>~ 14752.33 ms</strong><br>(/C = —) | <strong>~ 14767.19 ms</strong><br>(/C = —) | <strong>~ 1452.42 ms</strong><br>(/C = —)<br>(/JS = <strong>0.10x</strong>) | <strong>~ 1032.00 ms</strong><br>(/C = —)<br>(/JS = <strong>0.07x</strong>) |
↳ μ/C | — | — | — | — |
↳ Q₂/C | — | — | — | — |
↳ G/C | — | — | — | — |
↳ μ/JS | — | <strong>0.93x</strong> | <strong>0.92x</strong> | <strong>≈ 0.84x</strong> |
↳ Q₂/JS | — | <strong>1.00x</strong> | <strong>1.02x</strong> | <strong>≈ 1.08x</strong> |
↳ G/JS | — | <strong>0.91x</strong> | <strong>0.55x</strong> | <strong>≈ 0.46x</strong> |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with the Go/Rust multicore runtimes.

### Application benchmark results

App scenario | JS | Go | Go / JS |
--- | ---: | ---: | ---: |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (corpus parsing) | ~ 62.54 ms | ~ 167.35 ms | 2.68x |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (fixture parsing) | ~ 8.87 ms | ~ 22.77 ms | 2.57x |
[tidy](https://github.com/natefaubion/purescript-tidy) (parsing + formatting) | ~ 15.25 ms | ~ 35.03 ms | 2.30x |
[tidy](https://github.com/natefaubion/purescript-tidy) (10 iterations, current driver) | ~ 311.84 ms | ~ 344.87 ms | 1.11x |
[tidy](https://github.com/natefaubion/purescript-tidy) (10 iterations, historical baseline) | ~ 283.20 ms | ~ 556.44 ms | 1.96x |

### Compilation times...

#### ... to Go

Benchmark   | [gopurs](https://github.com/0x000000000000000000001/gopurs) JS binary<br>(official compiler)<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>`--dev`<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ------------------ | --------------------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 146.39 s | ~ 78.04 s <br>(/JS = 0.53x) | ~ 53.61 s <br>(/JS = 0.37x) | ~ 52.92 s <br>(/JS = 0.36x)
[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff) | ~ 16023 ms | ~ 6312 ms | ~ 5437 ms | ~ 5002 ms
[gopurs-argonaut-core](https://github.com/0x000000000000000000001/gopurs-argonaut-core) | ~ 15672 ms | ~ 5930 ms | ~ 4962 ms | ~ 4800 ms
[gopurs-arrays](https://github.com/0x000000000000000000001/gopurs-arrays) | ~ 12758 ms | ~ 5072 ms | ~ 3782 ms | ~ 4001 ms
[gopurs-assert](https://github.com/0x000000000000000000001/gopurs-assert) | ~ 2514 ms | ~ 1305 ms | ~ 1187 ms | ~ 922 ms
[gopurs-avar](https://github.com/0x000000000000000000001/gopurs-avar) | ~ 15097 ms | ~ 5865 ms | ~ 4971 ms | ~ 4770 ms
[gopurs-catenable-lists](https://github.com/0x000000000000000000001/gopurs-catenable-lists) | ~ 10188 ms | ~ 3918 ms | ~ 3387 ms | ~ 3024 ms
[gopurs-console](https://github.com/0x000000000000000000001/gopurs-console) | ~ 9476 ms | ~ 3207 ms | ~ 2880 ms | ~ 2513 ms
[gopurs-datetime](https://github.com/0x000000000000000000001/gopurs-datetime) | ~ 14497 ms | ~ 5560 ms | ~ 4768 ms | ~ 4324 ms
[gopurs-effect](https://github.com/0x000000000000000000001/gopurs-effect) | ~ 4161 ms | ~ 1390 ms | ~ 1296 ms | ~ 994 ms
[gopurs-enums](https://github.com/0x000000000000000000001/gopurs-enums) | ~ 10858 ms | ~ 3936 ms | ~ 3300 ms | ~ 3125 ms
[gopurs-exceptions](https://github.com/0x000000000000000000001/gopurs-exceptions) | ~ 4671 ms | ~ 1665 ms | ~ 1519 ms | ~ 1195 ms
[gopurs-foldable-traversable](https://github.com/0x000000000000000000001/gopurs-foldable-traversable) | ~ 10499 ms | ~ 3841 ms | ~ 3357 ms | ~ 3057 ms
[gopurs-foreign](https://github.com/0x000000000000000000001/gopurs-foreign) | ~ 12846 ms | ~ 5309 ms | ~ 4621 ms | ~ 4347 ms
[gopurs-foreign-object](https://github.com/0x000000000000000000001/gopurs-foreign-object) | ~ 15874 ms | ~ 6228 ms | ~ 5090 ms | ~ 5136 ms
[gopurs-free](https://github.com/0x000000000000000000001/gopurs-free) | ~ 11960 ms | ~ 4840 ms | ~ 4252 ms | ~ 3881 ms
[gopurs-functions](https://github.com/0x000000000000000000001/gopurs-functions) | ~ 4231 ms | ~ 1417 ms | ~ 1299 ms | ~ 960 ms
[gopurs-integers](https://github.com/0x000000000000000000001/gopurs-integers) | ~ 10024 ms | ~ 3470 ms | ~ 3098 ms | ~ 2707 ms
[gopurs-js-bigints](https://github.com/0x000000000000000000001/gopurs-js-bigints) | ~ 4625 ms | ~ 1596 ms | ~ 1443 ms | ~ 1179 ms
[gopurs-js-date](https://github.com/0x000000000000000000001/gopurs-js-date) | ~ 16044 ms | ~ 6227 ms | ~ 5297 ms | ~ 5007 ms
[gopurs-js-promise](https://github.com/0x000000000000000000001/gopurs-js-promise) | ~ 8189 ms | ~ 2787 ms | ~ 2531 ms | ~ 2196 ms
[gopurs-js-promise-aff](https://github.com/0x000000000000000000001/gopurs-js-promise-aff) | ~ 16617 ms | ~ 6434 ms | ~ 5511 ms | ~ 5158 ms
[gopurs-js-uri](https://github.com/0x000000000000000000001/gopurs-js-uri) | ~ 9730 ms | ~ 3275 ms | ~ 2948 ms | ~ 2569 ms
[gopurs-lazy](https://github.com/0x000000000000000000001/gopurs-lazy) | ~ 7605 ms | ~ 2758 ms | ~ 2526 ms | ~ 2178 ms
[gopurs-node-buffer](https://github.com/0x000000000000000000001/gopurs-node-buffer) | ~ 14716 ms | ~ 5351 ms | ~ 4743 ms | ~ 4366 ms
[gopurs-node-event-emitter](https://github.com/0x000000000000000000001/gopurs-node-event-emitter) | ~ 21845 ms | ~ 8936 ms | ~ 6631 ms | ~ 7607 ms
[gopurs-node-fs](https://github.com/0x000000000000000000001/gopurs-node-fs) | ~ 18796 ms | ~ 6999 ms | ~ 5912 ms | ~ 5685 ms
[gopurs-node-http](https://github.com/0x000000000000000000001/gopurs-node-http) | ~ 20650 ms | ~ 7475 ms | ~ 6248 ms | ~ 6173 ms
[gopurs-node-net](https://github.com/0x000000000000000000001/gopurs-node-net) | ~ 19243 ms | ~ 6945 ms | ~ 5865 ms | ~ 5630 ms
[gopurs-node-path](https://github.com/0x000000000000000000001/gopurs-node-path) | ~ 4257 ms | ~ 1302 ms | ~ 1204 ms | ~ 914 ms
[gopurs-node-process](https://github.com/0x000000000000000000001/gopurs-node-process) | ~ 17942 ms | ~ 6684 ms | ~ 5685 ms | ~ 5391 ms
[gopurs-node-streams](https://github.com/0x000000000000000000001/gopurs-node-streams) | ~ 23385 ms | ~ 9204 ms | ~ 6884 ms | ~ 7938 ms
[gopurs-now](https://github.com/0x000000000000000000001/gopurs-now) | ~ 18036 ms | ~ 6756 ms | ~ 5772 ms | ~ 5480 ms
[gopurs-nullable](https://github.com/0x000000000000000000001/gopurs-nullable) | ~ 14155 ms | ~ 5343 ms | ~ 4659 ms | ~ 4399 ms
[gopurs-numbers](https://github.com/0x000000000000000000001/gopurs-numbers) | ~ 10142 ms | ~ 3507 ms | ~ 3170 ms | ~ 2749 ms
[gopurs-ordered-collections](https://github.com/0x000000000000000000001/gopurs-ordered-collections) | ~ 16381 ms | ~ 6449 ms | ~ 5138 ms | ~ 5047 ms
[gopurs-partial](https://github.com/0x000000000000000000001/gopurs-partial) | ~ 9519 ms | ~ 3188 ms | ~ 2859 ms | ~ 2493 ms
[gopurs-prelude](https://github.com/0x000000000000000000001/gopurs-prelude) | ~ 10711 ms | ~ 3557 ms | ~ 3053 ms | ~ 2792 ms
[gopurs-random](https://github.com/0x000000000000000000001/gopurs-random) | ~ 9902 ms | ~ 3224 ms | ~ 2880 ms | ~ 2525 ms
[gopurs-record](https://github.com/0x000000000000000000001/gopurs-record) | ~ 10133 ms | ~ 3356 ms | ~ 3019 ms | ~ 2627 ms
[gopurs-refs](https://github.com/0x000000000000000000001/gopurs-refs) | ~ 9881 ms | ~ 3300 ms | ~ 2960 ms | ~ 2594 ms
[gopurs-run](https://github.com/0x000000000000000000001/gopurs-run) | ~ 17665 ms | ~ 6775 ms | ~ 5776 ms | ~ 5433 ms
[gopurs-spec](https://github.com/0x000000000000000000001/gopurs-spec) | ~ 30416 ms | ~ 12821 ms | ~ 8658 ms | ~ 10954 ms
[gopurs-st](https://github.com/0x000000000000000000001/gopurs-st) | ~ 9659 ms | ~ 3302 ms | ~ 2980 ms | ~ 2597 ms
[gopurs-strings](https://github.com/0x000000000000000000001/gopurs-strings) | ~ 12680 ms | ~ 4489 ms | ~ 3824 ms | ~ 3805 ms
[gopurs-strings-extra](https://github.com/0x000000000000000000001/gopurs-strings-extra) | ~ 13757 ms | ~ 6990 ms | ~ 6343 ms | ~ 4827 ms
[gopurs-unfoldable](https://github.com/0x000000000000000000001/gopurs-unfoldable) | ~ 9565 ms | ~ 3322 ms | ~ 2979 ms | ~ 2607 ms
[gopurs-unsafe-coerce](https://github.com/0x000000000000000000001/gopurs-unsafe-coerce) | ~ 9841 ms | ~ 3290 ms | ~ 2941 ms | ~ 2562 ms
[gopurs-uuid](https://github.com/0x000000000000000000001/gopurs-uuid) | ~ 20972 ms | ~ 8748 ms | ~ 6505 ms | ~ 7340 ms
[gopurs-variant](https://github.com/0x000000000000000000001/gopurs-variant) | ~ 11455 ms | ~ 4451 ms | ~ 3881 ms | ~ 3279 ms
[gopurs-yoga-json](https://github.com/0x000000000000000000001/gopurs-yoga-json) | ~ 36226 ms | ~ 16572 ms | ~ 11991 ms | ~ 13134 ms
**Total gopurs-*** | <strong>~ 666.09 s</strong> | <strong>~ 254.68 s</strong> <br>(/JS = <strong>0.38x</strong>) | <strong>~ 212.02 s</strong> <br>(/JS = <strong>0.32x</strong>) | <strong>~ 203.99 s</strong> <br>(/JS = <strong>0.31x</strong>)
**Total** | <strong>~ 812.48 s</strong> | <strong>~ 332.71 s</strong> <br>(/JS = <strong>0.41x</strong>) | <strong>~ 265.63 s</strong> <br>(/JS = <strong>0.33x</strong>) | <strong>~ 256.92 s</strong> <br>(/JS = <strong>0.32x</strong>)

#### ... to Rust

Benchmark   | [purust](https://github.com/0x000000000000000000001/purust) JS binary<br>(official compiler)<br>(WIP) | [purust](https://github.com/0x000000000000000000001/purust) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 44.31 s | ~ 20.62 s <br>(/JS = 0.47x)
[purust-aff](https://github.com/0x000000000000000000001/purust-aff) | ~ 6218 ms | ~ 2076 ms
[purust-argonaut-codecs](https://github.com/0x000000000000000000001/purust-argonaut-codecs) | ~ 6090 ms | ~ 2120 ms
[purust-argonaut-core](https://github.com/0x000000000000000000001/purust-argonaut-core) | ~ 5122 ms | ~ 1682 ms
[purust-arraybuffer-types](https://github.com/0x000000000000000000001/purust-arraybuffer-types) | ~ 591 ms | ~ 158 ms
[purust-arrays](https://github.com/0x000000000000000000001/purust-arrays) | ~ 15134 ms | ~ 4291 ms
[purust-assert](https://github.com/0x000000000000000000001/purust-assert) | ~ 631 ms | ~ 170 ms
[purust-avar](https://github.com/0x000000000000000000001/purust-avar) | ~ 5942 ms | ~ 1938 ms
[purust-catenable-lists](https://github.com/0x000000000000000000001/purust-catenable-lists) | ~ 3474 ms | ~ 1177 ms
[purust-console](https://github.com/0x000000000000000000001/purust-console) | ~ 664 ms | ~ 183 ms
[purust-datetime](https://github.com/0x000000000000000000001/purust-datetime) | ~ 5659 ms | ~ 1902 ms
[purust-effect](https://github.com/0x000000000000000000001/purust-effect) | ~ 1240 ms | ~ 475 ms
[purust-enums](https://github.com/0x000000000000000000001/purust-enums) | ~ 2950 ms | ~ 923 ms
[purust-exceptions](https://github.com/0x000000000000000000001/purust-exceptions) | ~ 924 ms | ~ 276 ms
[purust-exists](https://github.com/0x000000000000000000001/purust-exists) | ~ 629 ms | ~ 171 ms
[purust-foldable-traversable](https://github.com/0x000000000000000000001/purust-foldable-traversable) | ~ 3593 ms | ~ 1220 ms
[purust-foreign](https://github.com/0x000000000000000000001/purust-foreign) | ~ 4949 ms | ~ 1647 ms
[purust-foreign-object](https://github.com/0x000000000000000000001/purust-foreign-object) | ~ 6320 ms | ~ 2217 ms
[purust-free](https://github.com/0x000000000000000000001/purust-free) | ~ 4570 ms | ~ 1520 ms
[purust-functions](https://github.com/0x000000000000000000001/purust-functions) | ~ 1365 ms | ~ 508 ms
[purust-integers](https://github.com/0x000000000000000000001/purust-integers) | ~ 4359 ms | ~ 1398 ms
[purust-js-bigints](https://github.com/0x000000000000000000001/purust-js-bigints) | ~ 6620 ms | ~ 2336 ms
[purust-js-date](https://github.com/0x000000000000000000001/purust-js-date) | ~ 5642 ms | ~ 1933 ms
[purust-js-promise](https://github.com/0x000000000000000000001/purust-js-promise) | ~ 6029 ms | ~ 2018 ms
[purust-js-promise-aff](https://github.com/0x000000000000000000001/purust-js-promise-aff) | ~ 6165 ms | ~ 2078 ms
[purust-lazy](https://github.com/0x000000000000000000001/purust-lazy) | ~ 2140 ms | ~ 693 ms
[purust-node-buffer](https://github.com/0x000000000000000000001/purust-node-buffer) | ~ 4894 ms | ~ 1601 ms
[purust-node-child-process](https://github.com/0x000000000000000000001/purust-node-child-process) | ~ 7651 ms | ~ 2609 ms
[purust-node-event-emitter](https://github.com/0x000000000000000000001/purust-node-event-emitter) | ~ 8202 ms | ~ 2809 ms
[purust-node-fs](https://github.com/0x000000000000000000001/purust-node-fs) | ~ 7430 ms | ~ 2508 ms
[purust-node-http](https://github.com/0x000000000000000000001/purust-node-http) | ~ 7105 ms | ~ 2452 ms
[purust-node-net](https://github.com/0x000000000000000000001/purust-node-net) | ~ 6805 ms | ~ 2310 ms
[purust-node-os](https://github.com/0x000000000000000000001/purust-node-os) | ~ 6032 ms | ~ 2086 ms
[purust-node-path](https://github.com/0x000000000000000000001/purust-node-path) | ~ 650 ms | ~ 175 ms
[purust-node-process](https://github.com/0x000000000000000000001/purust-node-process) | ~ 7196 ms | ~ 2443 ms
[purust-node-streams](https://github.com/0x000000000000000000001/purust-node-streams) | ~ 9197 ms | ~ 3489 ms
[purust-now](https://github.com/0x000000000000000000001/purust-now) | ~ 6780 ms | ~ 2310 ms
[purust-nullable](https://github.com/0x000000000000000000001/purust-nullable) | ~ 4653 ms | ~ 1557 ms
[purust-numbers](https://github.com/0x000000000000000000001/purust-numbers) | ~ 4401 ms | ~ 1408 ms
[purust-ordered-collections](https://github.com/0x000000000000000000001/purust-ordered-collections) | ~ 8217 ms | ~ 2986 ms
[purust-partial](https://github.com/0x000000000000000000001/purust-partial) | ~ 871 ms | ~ 268 ms
[purust-prelude](https://github.com/0x000000000000000000001/purust-prelude) | ~ 2793 ms | ~ 985 ms
[purust-quickcheck](https://github.com/0x000000000000000000001/purust-quickcheck) | ~ 4753 ms | ~ 1624 ms
[purust-random](https://github.com/0x000000000000000000001/purust-random) | ~ 2660 ms | ~ 898 ms
[purust-record](https://github.com/0x000000000000000000001/purust-record) | ~ 2641 ms | ~ 896 ms
[purust-refs](https://github.com/0x000000000000000000001/purust-refs) | ~ 718 ms | ~ 196 ms
[purust-run](https://github.com/0x000000000000000000001/purust-run) | ~ 6914 ms | ~ 2262 ms
[purust-spec](https://github.com/0x000000000000000000001/purust-spec) | ~ 12717 ms | ~ 4633 ms
[purust-spec-discovery](https://github.com/0x000000000000000000001/purust-spec-discovery) | ~ 11401 ms | ~ 4318 ms
[purust-spec-node](https://github.com/0x000000000000000000001/purust-spec-node) | ~ 11931 ms | ~ 4307 ms
[purust-st](https://github.com/0x000000000000000000001/purust-st) | ~ 2554 ms | ~ 866 ms
[purust-strings](https://github.com/0x000000000000000000001/purust-strings) | ~ 6589 ms | ~ 1851 ms
[purust-strings-extra](https://github.com/0x000000000000000000001/purust-strings-extra) | ~ 7666 ms | ~ 11028 ms
[purust-unfoldable](https://github.com/0x000000000000000000001/purust-unfoldable) | ~ 2651 ms | ~ 920 ms
[purust-unsafe-coerce](https://github.com/0x000000000000000000001/purust-unsafe-coerce) | ~ 2089 ms | ~ 669 ms
[purust-uuid](https://github.com/0x000000000000000000001/purust-uuid) | ~ 7613 ms | ~ 2588 ms
[purust-variant](https://github.com/0x000000000000000000001/purust-variant) | ~ 4537 ms | ~ 1405 ms
[purust-yoga-json](https://github.com/0x000000000000000000001/purust-yoga-json) | ~ 16694 ms | ~ 14548 ms
**Total purust-*** | <strong>~ 304.03 s</strong> | <strong>~ 120.12 s</strong> <br>(/JS = <strong>0.40x</strong>)
**Total** | <strong>~ 348.34 s</strong> | <strong>~ 140.73 s</strong> <br>(/JS = <strong>0.40x</strong>)
