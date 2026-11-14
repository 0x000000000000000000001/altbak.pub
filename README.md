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
The historical benchmark tables distinguish two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) executed on the configured core backends via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Excluded WIP benchmarks

Normal runs execute the **14 core benchmarks**, without WIP cases. Use
`./bin/go/run --wip-included` to include the core suite, **23 library cases**
(array operations, `Traversable` with `Maybe`, Unicode strings, flat record
updates, and `ST.Ref` / `Effect.Ref` loops), plus Array Indexing, JSON to Typed
AST and JSON Decoding. JS and Rust support the same selection; JS/PBO includes
the core suite and the 23 library cases, with the three older diagnostics not
configured for that backend. psgo also includes the 14 core and 23 new cases.
JS and Go accept `--ffi --wip-included` or `--fficc --wip-included` for the
corresponding native references. Each diagnostic retains its measurement protocol.
`./bin/run --wip-included` applies this selection to JS, JS/PBO, Go, Rust and psgo;
the other configured backends run their core suite only.

Every new library case has its own row in each core results table below. Rows
grouped under **Excluded WIP** remain outside the historical **Total Execution Time**
and **/C** ratios, even when requested with `--wip-included`. For the new rows,
**pending** means a configured implementation is pending runtime qualification,
with no published measurements; **—** means no implementation is configured for
that column. Existing diagnostic measurements retain their recorded values.

The [suite documentation](docs/library-benchmarks.md) records provenance,
timing boundaries, expected results and individual backend commands. Each test
has its own `src/Test/<Name>/` directory containing its PureScript and FFI files.
WIP selection remains explicit: configured runners include these cases only with
`--wip-included`.

The 23 new rows in all nine JS/Go columns use medians from three independent
processes, measured on November 14, 2026. The [measurement archive](docs/benchmark-results/2026-11-14-library-js-go.md)
records all 621 validated WIP outputs, build profiles and source fingerprints.
The psgo column includes a documented local UTF-16 length correction.

### Core stresstest benchmark results (pure computational)

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.26 μs<br>(/C = 2.68x) | ~ 0.19 μs<br>(/C = 1.96x)<br>(/JS = 0.73x) | ~ 0.20 μs<br>(/C = 2.06x)<br>(/JS = 0.77x) | ~ 0.22 μs<br>(/C = 2.27x)<br>(/JS = 0.85x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs<br>(/C = 3.73x) | ~ 0.31 μs<br>(/C = 3.73x) | ~ 0.31 μs<br>(/C = 3.73x) | ~ 0.37 μs<br>(/C = 4.46x)<br>(/JS = 1.19x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 9.59 μs<br>(/C = 188.04x) | ~ 5.06 μs<br>(/C = 99.22x)<br>(/JS = 0.53x) | ~ 19.06 μs<br>(/C = 373.73x)<br>(/JS = 1.99x) | ~ 0.49 μs<br>(/C = 9.61x)<br>(/JS = 0.05x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 189.49 μs<br>(/C = 5.53x) | ~ 151.13 μs<br>(/C = 4.41x)<br>(/JS = 0.80x) | ~ 65.91 μs<br>(/C = 1.92x)<br>(/JS = 0.35x) | ~ 68.31 μs<br>(/C = 1.99x)<br>(/JS = 0.36x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 74.18 μs<br>(/C = 21.66x) | ~ 134.11 μs<br>(/C = 39.16x)<br>(/JS = 1.81x) | ~ 42.07 μs<br>(/C = 12.28x)<br>(/JS = 0.57x) | ~ 7.69 μs<br>(/C = 2.25x)<br>(/JS = 0.10x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 67.49 μs<br>(/C = 3.99x) | ~ 64.40 μs<br>(/C = 3.81x)<br>(/JS = 0.95x) | ~ 38.25 μs<br>(/C = 2.26x)<br>(/JS = 0.57x) | ~ 38.48 μs<br>(/C = 2.28x)<br>(/JS = 0.57x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 1263.82 μs<br>(/C = 1263820.00x) | ~ 1242.29 μs<br>(/C = 1242290.00x)<br>(/JS = 0.98x) | ~ 1337.51 μs<br>(/C = 1337510.00x)<br>(/JS = 1.06x) | ~ 49.63 μs<br>(/C = 49630.00x)<br>(/JS = 0.04x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 63.39 μs<br>(/C = 58.10x) | ~ 26.99 μs<br>(/C = 24.74x)<br>(/JS = 0.43x) | ~ 86.33 μs<br>(/C = 79.13x)<br>(/JS = 1.36x) | ~ 0.65 μs<br>(/C = 0.60x)<br>(/JS = 0.01x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 89940.58 μs<br>(/C = 9.42x) | ~ 47272.88 μs<br>(/C = 4.95x)<br>(/JS = 0.53x) | ~ 24486.83 μs<br>(/C = 2.56x)<br>(/JS = 0.27x) | ~ 27249.58 μs<br>(/C = 2.85x)<br>(/JS = 0.30x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 22830.25 μs<br>(/C = 22830250.00x) | ~ 7368.38 μs<br>(/C = 7368380.00x)<br>(/JS = 0.32x) | ~ 4843 μs<br>(/C = 4843000.00x)<br>(/JS = 0.21x) | ~ 2805.25 μs<br>(/C = 2805250.00x)<br>(/JS = 0.12x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 48.23 μs<br>(/C = 48230.00x) | ~ 10.95 μs<br>(/C = 10950.00x)<br>(/JS = 0.23x) | ~ 46.87 μs<br>(/C = 46870.00x)<br>(/JS = 0.97x) | ~ 0.42 μs<br>(/C = 420.00x)<br>(/JS = 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 14608.29 μs<br>(/C = 14608290.00x) | ~ 10604.88 μs<br>(/C = 10604880.00x)<br>(/JS = 0.73x) | ~ 17513.12 μs<br>(/C = 17513120.00x)<br>(/JS = 1.20x) | ~ 321.48 μs<br>(/C = 321480.00x)<br>(/JS = 0.02x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.35 μs<br>(/C = 147.00x) | ~ 3.75 μs<br>(/C = 75.00x)<br>(/JS = 0.51x) | ~ 3.12 μs<br>(/C = 62.40x)<br>(/JS = 0.42x) | ~ 0.50 μs<br>(/C = 10.00x)<br>(/JS = 0.07x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.04 μs<br>(/C = 40.00x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = 0.50x) | ~ 0.06 μs<br>(/C = 60.00x)<br>(/JS = 1.50x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = 0.50x) |
**Total Execution Time** | ~ 129.10 ms<br>(/C = 13.43x) | ~ 66.89 ms<br>(/C = 6.96x)<br>(/JS = 0.52x) | ~ 48.48 ms<br>(/C = 5.04x)<br>(/JS = 0.38x) | ~ 30.54 ms<br>(/C = 3.18x)<br>(/JS = 0.24x) |
**Excluded WIP** | | | | |
[Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 25303.33 μs<br>(/C = 3.84x) | — | — | — |
[JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 80646.58 μs<br>(/C = 7.64x) | — | — | — |
[JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 9283.94 μs<br>(/C = 14.12x) | — | — | — |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 121.94 μs<br>(/C = —) | ~ 144.60 μs<br>(/C = —)<br>(/JS = 1.19x) | ~ 92.20 μs<br>(/C = —)<br>(/JS = 0.76x) | ~ 12.93 μs<br>(/C = —)<br>(/JS = 0.11x) |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 127.59 μs<br>(/C = —) | ~ 130.46 μs<br>(/C = —)<br>(/JS = 1.02x) | ~ 74.63 μs<br>(/C = —)<br>(/JS = 0.58x) | ~ 18.58 μs<br>(/C = —)<br>(/JS = 0.15x) |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 70.25 μs<br>(/C = —) | ~ 70.22 μs<br>(/C = —) | ~ 44.01 μs<br>(/C = —)<br>(/JS = 0.63x) | ~ 3.32 μs<br>(/C = —)<br>(/JS = 0.05x) |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 561.69 μs<br>(/C = —) | ~ 517.87 μs<br>(/C = —)<br>(/JS = 0.92x) | ~ 114.68 μs<br>(/C = —)<br>(/JS = 0.20x) | ~ 78.30 μs<br>(/C = —)<br>(/JS = 0.14x) |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 51.63 μs<br>(/C = —) | ~ 50.71 μs<br>(/C = —)<br>(/JS = 0.98x) | ~ 41.04 μs<br>(/C = —)<br>(/JS = 0.79x) | ~ 5.03 μs<br>(/C = —)<br>(/JS = 0.10x) |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.20 μs<br>(/C = —) | ~ 31.05 μs<br>(/C = —)<br>(/JS = 1.03x) | ~ 2.05 μs<br>(/C = —)<br>(/JS = 0.07x) | ~ 2.40 μs<br>(/C = —)<br>(/JS = 0.08x) |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 80.61 μs<br>(/C = —) | ~ 81.04 μs<br>(/C = —)<br>(/JS = 1.01x) | ~ 49.83 μs<br>(/C = —)<br>(/JS = 0.62x) | ~ 1.39 μs<br>(/C = —)<br>(/JS = 0.02x) |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 68.35 μs<br>(/C = —) | ~ 65.14 μs<br>(/C = —)<br>(/JS = 0.95x) | ~ 42.18 μs<br>(/C = —)<br>(/JS = 0.62x) | ~ 1.07 μs<br>(/C = —)<br>(/JS = 0.02x) |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 68.01 μs<br>(/C = —) | ~ 64.65 μs<br>(/C = —)<br>(/JS = 0.95x) | ~ 42.45 μs<br>(/C = —)<br>(/JS = 0.62x) | ~ 1.08 μs<br>(/C = —)<br>(/JS = 0.02x) |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.01 μs<br>(/C = —) | ~ 0.01 μs<br>(/C = —)<br>(/JS = 0.86x) | ~ 0.02 μs<br>(/C = —)<br>(/JS = 1.92x) | ~ 0.02 μs<br>(/C = —)<br>(/JS = 1.94x) |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 37.44 μs<br>(/C = —) | ~ 37.44 μs<br>(/C = —) | ~ 37.09 μs<br>(/C = —)<br>(/JS = 0.99x) | ~ 8.54 μs<br>(/C = —)<br>(/JS = 0.23x) |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.19 μs<br>(/C = —) | ~ 0.19 μs<br>(/C = —)<br>(/JS = 0.99x) | ~ 0.19 μs<br>(/C = —)<br>(/JS = 1.02x) | ~ 0.19 μs<br>(/C = —) |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 63.80 μs<br>(/C = —) | ~ 63.97 μs<br>(/C = —) | ~ 63.94 μs<br>(/C = —) | ~ 63.72 μs<br>(/C = —) |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 45.28 μs<br>(/C = —) | ~ 44.96 μs<br>(/C = —)<br>(/JS = 0.99x) | ~ 42.10 μs<br>(/C = —)<br>(/JS = 0.93x) | ~ 21.01 μs<br>(/C = —)<br>(/JS = 0.46x) |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 145.33 μs<br>(/C = —) | ~ 144.87 μs<br>(/C = —) | ~ 138.40 μs<br>(/C = —)<br>(/JS = 0.95x) | ~ 24.69 μs<br>(/C = —)<br>(/JS = 0.17x) |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 47.71 μs<br>(/C = —) | ~ 46.52 μs<br>(/C = —)<br>(/JS = 0.97x) | ~ 42.64 μs<br>(/C = —)<br>(/JS = 0.89x) | ~ 11.85 μs<br>(/C = —)<br>(/JS = 0.25x) |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 2.44 μs<br>(/C = —) | ~ 4.37 μs<br>(/C = —)<br>(/JS = 1.79x) | ~ 4.32 μs<br>(/C = —)<br>(/JS = 1.77x) | ~ 0.72 μs<br>(/C = —)<br>(/JS = 0.30x) |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 4.55 μs<br>(/C = —) | ~ 9.34 μs<br>(/C = —)<br>(/JS = 2.05x) | ~ 9.04 μs<br>(/C = —)<br>(/JS = 1.98x) | ~ 0.71 μs<br>(/C = —)<br>(/JS = 0.16x) |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 7.03 μs<br>(/C = —) | ~ 14.73 μs<br>(/C = —)<br>(/JS = 2.10x) | ~ 14.38 μs<br>(/C = —)<br>(/JS = 2.05x) | ~ 0.72 μs<br>(/C = —)<br>(/JS = 0.10x) |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 463.87 μs<br>(/C = —) | ~ 106.02 μs<br>(/C = —)<br>(/JS = 0.23x) | ~ 2.44 μs<br>(/C = —)<br>(/JS = 0.01x) | ~ 2.48 μs<br>(/C = —)<br>(/JS = 0.01x) |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 513.20 μs<br>(/C = —) | ~ 106.03 μs<br>(/C = —)<br>(/JS = 0.21x) | ~ 31.67 μs<br>(/C = —)<br>(/JS = 0.06x) | ~ 3.41 μs<br>(/C = —)<br>(/JS = 0.01x) |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 155.67 μs<br>(/C = —) | ~ 102.38 μs<br>(/C = —)<br>(/JS = 0.66x) | ~ 2.39 μs<br>(/C = —)<br>(/JS = 0.02x) | ~ 2.56 μs<br>(/C = —)<br>(/JS = 0.02x) |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 211.84 μs<br>(/C = —) | ~ 113.53 μs<br>(/C = —)<br>(/JS = 0.54x) | ~ 31.14 μs<br>(/C = —)<br>(/JS = 0.15x) | ~ 3.52 μs<br>(/C = —)<br>(/JS = 0.02x) |

Rows grouped under **Excluded WIP** are not included in **Total Execution Time** or the **/C** ratios.

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs) <br>`--dev`<br>↓<br>Go<br>↓<br>go-build<br>`-N -l`<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) |
----------------------- | -------------------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.33 μs<br>(/C = 3.40x)<br>(/JS = 1.27x) | ~ 0.52 μs<br>(/C = 5.36x)<br>(/JS = 2.00x) | ~ 4.17 μs<br>(/C = 42.99x)<br>(/JS = 16.04x) | ~ 0.32 μs<br>(/C = 3.30x)<br>(/JS = 1.23x) | ~ 0.32 μs<br>(/C = 3.30x)<br>(/JS = 1.23x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.25 μs<br>(/C = 3.01x)<br>(/JS = 0.81x) | ~ 0.54 μs<br>(/C = 6.51x)<br>(/JS = 1.74x) | ~ 5.41 μs<br>(/C = 65.18x)<br>(/JS = 17.45x) | ~ 0.16 μs<br>(/C = 1.93x)<br>(/JS = 0.52x) | ~ 0.16 μs<br>(/C = 1.93x)<br>(/JS = 0.52x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 17.19 μs<br>(/C = 337.06x)<br>(/JS = 1.79x) | ~ 28.86 μs<br>(/C = 565.88x)<br>(/JS = 3.01x) | ~ 236.01 μs<br>(/C = 4627.65x)<br>(/JS = 24.61x) | ~ 24.39 μs<br>(/C = 478.24x)<br>(/JS = 2.54x) | ~ 0.45 μs<br>(/C = 8.82x)<br>(/JS = 0.05x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.18 μs<br>(/C = 1.67x)<br>(/JS = 0.30x) | ~ 355.46 μs<br>(/C = 10.37x)<br>(/JS = 1.88x) | ~ 6281.96 μs<br>(/C = 183.24x)<br>(/JS = 33.15x) | ~ 35.98 μs<br>(/C = 1.05x)<br>(/JS = 0.19x) | ~ 35.50 μs<br>(/C = 1.04x)<br>(/JS = 0.19x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 6.63 μs<br>(/C = 1.94x)<br>(/JS = 0.09x) | ~ 186.36 μs<br>(/C = 54.41x)<br>(/JS = 2.51x) | ~ 5461.52 μs<br>(/C = 1594.60x)<br>(/JS = 73.63x) | ~ 40.43 μs<br>(/C = 11.80x)<br>(/JS = 0.55x) | ~ 14.87 μs<br>(/C = 4.34x)<br>(/JS = 0.20x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.93 μs<br>(/C = 0.94x)<br>(/JS = 0.24x) | ~ 35.56 μs<br>(/C = 2.10x)<br>(/JS = 0.53x) | ~ 402.04 μs<br>(/C = 23.78x)<br>(/JS = 5.96x) | ~ 20.85 μs<br>(/C = 1.23x)<br>(/JS = 0.31x) | ~ 19.87 μs<br>(/C = 1.18x)<br>(/JS = 0.29x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 230.56 μs<br>(/C = 230560.00x)<br>(/JS = 0.18x) | ~ 1296.16 μs<br>(/C = 1296160.00x)<br>(/JS = 1.03x) | ~ 3811.48 μs<br>(/C = 3811480.00x)<br>(/JS = 3.02x) | ~ 1044.60 μs<br>(/C = 1044600.00x)<br>(/JS = 0.83x) | ~ 22.45 μs<br>(/C = 22450.00x)<br>(/JS = 0.02x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 79.35 μs<br>(/C = 72.73x)<br>(/JS = 1.25x) | ~ 76.41 μs<br>(/C = 70.04x)<br>(/JS = 1.21x) | ~ 1340.99 μs<br>(/C = 1229.14x)<br>(/JS = 21.15x) | ~ 110.87 μs<br>(/C = 101.62x)<br>(/JS = 1.75x) | ~ 0.62 μs<br>(/C = 0.57x)<br>(/JS = 0.01x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10161.25 μs<br>(/C = 1.06x)<br>(/JS = 0.11x) | ~ 31239.75 μs<br>(/C = 3.27x)<br>(/JS = 0.35x) | ~ 870404.62 μs<br>(/C = 91.14x)<br>(/JS = 9.68x) | ~ 24630.67 μs<br>(/C = 2.58x)<br>(/JS = 0.27x) | ~ 9004.48 μs<br>(/C = 0.94x)<br>(/JS = 0.10x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2292.36 μs<br>(/C = 2292360.00x)<br>(/JS = 0.10x) | ~ 15617.54 μs<br>(/C = 15617540.00x)<br>(/JS = 0.68x) | ~ 505069 μs<br>(/C = 505069000.00x)<br>(/JS = 22.12x) | ~ 49160.92 μs<br>(/C = 49160920.00x)<br>(/JS = 2.15x) | ~ 2244.45 μs<br>(/C = 2244450.00x)<br>(/JS = 0.10x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 150.94 μs<br>(/C = 150940.00x)<br>(/JS = 3.13x) | ~ 110.57 μs<br>(/C = 110570.00x)<br>(/JS = 2.29x) | ~ 303.56 μs<br>(/C = 303560.00x)<br>(/JS = 6.29x) | ~ 52.10 μs<br>(/C = 52100.00x)<br>(/JS = 1.08x) | ~ 0.33 μs<br>(/C = 330.00x)<br>(/JS = 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 247.11 μs<br>(/C = 247110.00x)<br>(/JS = 0.02x) | ~ 1566.95 μs<br>(/C = 1566950.00x)<br>(/JS = 0.11x) | ~ 71378.08 μs<br>(/C = 71378080.00x)<br>(/JS = 4.89x) | ~ 13865.38 μs<br>(/C = 13865380.00x)<br>(/JS = 0.95x) | ~ 0.27 μs<br>(/C = 270.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 19.26 μs<br>(/C = 385.20x)<br>(/JS = 2.62x) | ~ 32.70 μs<br>(/C = 654.00x)<br>(/JS = 4.45x) | ~ 43.81 μs<br>(/C = 876.20x)<br>(/JS = 5.96x) | ~ 4.90 μs<br>(/C = 98.00x)<br>(/JS = 0.67x) | ~ 0.45 μs<br>(/C = 9.00x)<br>(/JS = 0.06x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.15 μs<br>(/C = 150.00x)<br>(/JS = 3.75x) | ~ 0.30 μs<br>(/C = 300.00x)<br>(/JS = 7.50x) | ~ 0.50 μs<br>(/C = 500.00x)<br>(/JS = 12.50x) | ~ 0.07 μs<br>(/C = 70.00x)<br>(/JS = 1.75x) | ~ 0.03 μs<br>(/C = 30.00x)<br>(/JS = 0.75x) |
**Total Execution Time** | ~ 13.28 ms<br>(/C = 1.38x)<br>(/JS = 0.10x) | ~ 50.55 ms<br>(/C = 5.26x)<br>(/JS = 0.39x) | ~ 1464.74 ms<br>(/C = 152.42x)<br>(/JS = 11.35x) | ~ 88.99 ms<br>(/C = 9.26x)<br>(/JS = 0.69x) | ~ 11.34 ms<br>(/C = 1.18x)<br>(/JS = 0.09x) |
**Excluded WIP** | | | | | |
[Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 3962.17 μs<br>(/C = 0.60x)<br>(/JS = 0.16x) | — | — | — | — |
[JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 20002.75 μs<br>(/C = 1.89x)<br>(/JS = 0.25x) | — | — | — | — |
[JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 2241.42 μs<br>(/C = 3.41x)<br>(/JS = 0.24x) | — | — | — | — |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 745.24 μs<br>(/C = —)<br>(/JS = 6.11x) | ~ 1452.41 μs<br>(/C = —)<br>(/JS = 11.91x) | ~ 2146.41 μs<br>(/C = —)<br>(/JS = 17.60x) | ~ 77.41 μs<br>(/C = —)<br>(/JS = 0.63x) | ~ 6.61 μs<br>(/C = —)<br>(/JS = 0.05x) |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 397.02 μs<br>(/C = —)<br>(/JS = 3.11x) | ~ 828.92 μs<br>(/C = —)<br>(/JS = 6.50x) | ~ 1330.39 μs<br>(/C = —)<br>(/JS = 10.43x) | ~ 67.38 μs<br>(/C = —)<br>(/JS = 0.53x) | ~ 6.85 μs<br>(/C = —)<br>(/JS = 0.05x) |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 65.06 μs<br>(/C = —)<br>(/JS = 0.93x) | ~ 281.15 μs<br>(/C = —)<br>(/JS = 4.00x) | ~ 133.05 μs<br>(/C = —)<br>(/JS = 1.89x) | ~ 49.77 μs<br>(/C = —)<br>(/JS = 0.71x) | ~ 2.31 μs<br>(/C = —)<br>(/JS = 0.03x) |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 6631.06 μs<br>(/C = —)<br>(/JS = 11.81x) | ~ 12300.38 μs<br>(/C = —)<br>(/JS = 21.90x) | ~ 11626.08 μs<br>(/C = —)<br>(/JS = 20.70x) | ~ 825.35 μs<br>(/C = —)<br>(/JS = 1.47x) | ~ 10.36 μs<br>(/C = —)<br>(/JS = 0.02x) |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 269.64 μs<br>(/C = —)<br>(/JS = 5.22x) | ~ 752.15 μs<br>(/C = —)<br>(/JS = 14.57x) | ~ 98.90 μs<br>(/C = —)<br>(/JS = 1.92x) | ~ 48.10 μs<br>(/C = —)<br>(/JS = 0.93x) | ~ 2.53 μs<br>(/C = —)<br>(/JS = 0.05x) |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 171.34 μs<br>(/C = —)<br>(/JS = 5.67x) | ~ 304.51 μs<br>(/C = —)<br>(/JS = 10.08x) | ~ 372.03 μs<br>(/C = —)<br>(/JS = 12.32x) | ~ 13.95 μs<br>(/C = —)<br>(/JS = 0.46x) | ~ 1.25 μs<br>(/C = —)<br>(/JS = 0.04x) |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 286.46 μs<br>(/C = —)<br>(/JS = 3.55x) | ~ 464.65 μs<br>(/C = —)<br>(/JS = 5.76x) | ~ 476.59 μs<br>(/C = —)<br>(/JS = 5.91x) | ~ 18.97 μs<br>(/C = —)<br>(/JS = 0.24x) | ~ 0.87 μs<br>(/C = —)<br>(/JS = 0.01x) |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 180.46 μs<br>(/C = —)<br>(/JS = 2.64x) | ~ 298.48 μs<br>(/C = —)<br>(/JS = 4.37x) | ~ 303.08 μs<br>(/C = —)<br>(/JS = 4.43x) | ~ 11.64 μs<br>(/C = —)<br>(/JS = 0.17x) | ~ 0.74 μs<br>(/C = —)<br>(/JS = 0.01x) |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 179.33 μs<br>(/C = —)<br>(/JS = 2.64x) | ~ 297.10 μs<br>(/C = —)<br>(/JS = 4.37x) | ~ 307.26 μs<br>(/C = —)<br>(/JS = 4.52x) | ~ 11.44 μs<br>(/C = —)<br>(/JS = 0.17x) | ~ 0.74 μs<br>(/C = —)<br>(/JS = 0.01x) |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 6.22 μs<br>(/C = —)<br>(/JS = 615.93x) | ~ 17.05 μs<br>(/C = —)<br>(/JS = 1687.77x) | ~ 3.47 μs<br>(/C = —)<br>(/JS = 343.62x) | ~ 3.45 μs<br>(/C = —)<br>(/JS = 341.43x) | ~ 3.45 μs<br>(/C = —)<br>(/JS = 341.50x) |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 49.37 μs<br>(/C = —)<br>(/JS = 1.32x) | ~ 79.96 μs<br>(/C = —)<br>(/JS = 2.14x) | ~ 62.44 μs<br>(/C = —)<br>(/JS = 1.67x) | ~ 7.20 μs<br>(/C = —)<br>(/JS = 0.19x) | ~ 3.41 μs<br>(/C = —)<br>(/JS = 0.09x) |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 35.34 μs<br>(/C = —)<br>(/JS = 185.27x) | ~ 130.75 μs<br>(/C = —)<br>(/JS = 685.42x) | ~ 2.71 μs<br>(/C = —)<br>(/JS = 14.19x) | ~ 0.17 μs<br>(/C = —)<br>(/JS = 0.90x) | ~ 0.17 μs<br>(/C = —)<br>(/JS = 0.88x) |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 37.48 μs<br>(/C = —)<br>(/JS = 0.59x) | ~ 140.42 μs<br>(/C = —)<br>(/JS = 2.20x) | ~ 2.47 μs<br>(/C = —)<br>(/JS = 0.04x) | ~ 2.31 μs<br>(/C = —)<br>(/JS = 0.04x) | ~ 2.32 μs<br>(/C = —)<br>(/JS = 0.04x) |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 56.18 μs<br>(/C = —)<br>(/JS = 1.24x) | ~ 105.92 μs<br>(/C = —)<br>(/JS = 2.34x) | ~ 104.16 μs<br>(/C = —)<br>(/JS = 2.30x) | ~ 146.89 μs<br>(/C = —)<br>(/JS = 3.24x) | ~ 3.66 μs<br>(/C = —)<br>(/JS = 0.08x) |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 158.15 μs<br>(/C = —)<br>(/JS = 1.09x) | ~ 318.05 μs<br>(/C = —)<br>(/JS = 2.19x) | ~ 300.93 μs<br>(/C = —)<br>(/JS = 2.07x) | ~ 32.68 μs<br>(/C = —)<br>(/JS = 0.22x) | ~ 9.75 μs<br>(/C = —)<br>(/JS = 0.07x) |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 120.71 μs<br>(/C = —)<br>(/JS = 2.53x) | ~ 216.07 μs<br>(/C = —)<br>(/JS = 4.53x) | ~ 167.98 μs<br>(/C = —)<br>(/JS = 3.52x) | ~ 74.43 μs<br>(/C = —)<br>(/JS = 1.56x) | ~ 3.71 μs<br>(/C = —)<br>(/JS = 0.08x) |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.51 μs<br>(/C = —)<br>(/JS = 0.21x) | ~ 5.46 μs<br>(/C = —)<br>(/JS = 2.24x) | ~ 144.74 μs<br>(/C = —)<br>(/JS = 59.23x) | ~ 0.72 μs<br>(/C = —)<br>(/JS = 0.29x) | ~ 0.24 μs<br>(/C = —)<br>(/JS = 0.10x) |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 5.78 μs<br>(/C = —)<br>(/JS = 1.27x) | ~ 13.30 μs<br>(/C = —)<br>(/JS = 2.92x) | ~ 190.03 μs<br>(/C = —)<br>(/JS = 41.74x) | ~ 7.20 μs<br>(/C = —)<br>(/JS = 1.58x) | ~ 0.26 μs<br>(/C = —)<br>(/JS = 0.06x) |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 9.58 μs<br>(/C = —)<br>(/JS = 1.36x) | ~ 22.40 μs<br>(/C = —)<br>(/JS = 3.19x) | ~ 679.42 μs<br>(/C = —)<br>(/JS = 96.67x) | ~ 12.85 μs<br>(/C = —)<br>(/JS = 1.83x) | ~ 1.03 μs<br>(/C = —)<br>(/JS = 0.15x) |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 536.37 μs<br>(/C = —)<br>(/JS = 1.16x) | ~ 1019.37 μs<br>(/C = —)<br>(/JS = 2.20x) | ~ 3085.81 μs<br>(/C = —)<br>(/JS = 6.65x) | ~ 2.25 μs<br>(/C = —)<br>(/JS = < 0.01x) | ~ 2.25 μs<br>(/C = —)<br>(/JS = < 0.01x) |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 538.57 μs<br>(/C = —)<br>(/JS = 1.05x) | ~ 1073.35 μs<br>(/C = —)<br>(/JS = 2.09x) | ~ 4068.22 μs<br>(/C = —)<br>(/JS = 7.93x) | ~ 2.33 μs<br>(/C = —)<br>(/JS = < 0.01x) | ~ 2.38 μs<br>(/C = —)<br>(/JS = < 0.01x) |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 1014.83 μs<br>(/C = —)<br>(/JS = 6.52x) | ~ 2199.81 μs<br>(/C = —)<br>(/JS = 14.13x) | ~ 3135.15 μs<br>(/C = —)<br>(/JS = 20.14x) | ~ 2.26 μs<br>(/C = —)<br>(/JS = 0.01x) | ~ 2.26 μs<br>(/C = —)<br>(/JS = 0.01x) |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 3195.80 μs<br>(/C = —)<br>(/JS = 15.09x) | ~ 4998.00 μs<br>(/C = —)<br>(/JS = 23.59x) | ~ 4026.20 μs<br>(/C = —)<br>(/JS = 19.01x) | ~ 2.36 μs<br>(/C = —)<br>(/JS = 0.01x) | ~ 2.36 μs<br>(/C = —)<br>(/JS = 0.01x) |

Rows grouped under **Excluded WIP** are not included in **Total Execution Time** or the **/C** ratios.

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.07 μs<br>(/C = 0.72x)<br>(/JS = 0.27x) | ~ 0.08 μs<br>(/C = 0.82x)<br>(/JS = 0.31x) | ~ 0.07 μs<br>(/C = 0.72x)<br>(/JS = 0.27x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs<br>(/C = 1.81x)<br>(/JS = 0.48x) | ~ 0.15 μs<br>(/C = 1.81x)<br>(/JS = 0.48x) | ~ 0.14 μs<br>(/C = 1.69x)<br>(/JS = 0.45x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 6.35 μs<br>(/C = 124.51x)<br>(/JS = 0.66x) | ~ 3.76 μs<br>(/C = 73.73x)<br>(/JS = 0.39x) | ~ 2.84 μs<br>(/C = 55.69x)<br>(/JS = 0.30x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 277.12 μs<br>(/C = 8.08x)<br>(/JS = 1.46x) | ~ 322.34 μs<br>(/C = 9.40x)<br>(/JS = 1.70x) | ~ 329.88 μs<br>(/C = 9.62x)<br>(/JS = 1.74x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 228.44 μs<br>(/C = 66.70x)<br>(/JS = 3.08x) | ~ 48.07 μs<br>(/C = 14.04x)<br>(/JS = 0.65x) | ~ 43.57 μs<br>(/C = 12.72x)<br>(/JS = 0.59x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.99 μs<br>(/C = 1.00x)<br>(/JS = 0.25x) | ~ 9.07 μs<br>(/C = 0.54x)<br>(/JS = 0.13x) | ~ 9.14 μs<br>(/C = 0.54x)<br>(/JS = 0.14x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 237.12 μs<br>(/C = 237120.00x)<br>(/JS = 0.19x) | ~ 235.53 μs<br>(/C = 235530.00x)<br>(/JS = 0.19x) | ~ 62.26 μs<br>(/C = 62260.00x)<br>(/JS = 0.05x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 37.49 μs<br>(/C = 34.36x)<br>(/JS = 0.59x) | ~ 22.68 μs<br>(/C = 20.79x)<br>(/JS = 0.36x) | ~ 0.80 μs<br>(/C = 0.73x)<br>(/JS = 0.01x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 22551 μs<br>(/C = 2.36x)<br>(/JS = 0.25x) | ~ 15428 μs<br>(/C = 1.62x)<br>(/JS = 0.17x) | ~ 11970 μs<br>(/C = 1.25x)<br>(/JS = 0.13x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 17812 μs<br>(/C = 17812000.00x)<br>(/JS = 0.78x) | ~ 27373 μs<br>(/C = 27373000.00x)<br>(/JS = 1.20x) | ~ 6730.50 μs<br>(/C = 6730500.00x)<br>(/JS = 0.29x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 1.70 μs<br>(/C = 1700.00x)<br>(/JS = 0.04x) | ~ 2.02 μs<br>(/C = 2020.00x)<br>(/JS = 0.04x) | ~ 0.84 μs<br>(/C = 840.00x)<br>(/JS = 0.02x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 2989.50 μs<br>(/C = 2989500.00x)<br>(/JS = 0.20x) | ~ 1964 μs<br>(/C = 1964000.00x)<br>(/JS = 0.13x) | ~ 0.61 μs<br>(/C = 610.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.29 μs<br>(/C = 145.80x)<br>(/JS = 0.99x) | ~ 5.04 μs<br>(/C = 100.80x)<br>(/JS = 0.69x) | ~ 2.79 μs<br>(/C = 55.80x)<br>(/JS = 0.38x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs<br>(/C = —)<br>(/JS = —) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = 0.25x) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
**Total Execution Time** | ~ 44.17 ms<br>(/C = 4.60x)<br>(/JS = 0.34x) | ~ 45.41 ms<br>(/C = 4.73x)<br>(/JS = 0.35x) | ~ 19.15 ms<br>(/C = 1.99x)<br>(/JS = 0.15x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.11 μs<br>(/C = 1.13x)<br>(/JS = 0.42x) | ~ 0.11 μs<br>(/C = 1.13x)<br>(/JS = 0.42x) | ~ 0.10 μs<br>(/C = 1.03x)<br>(/JS = 0.38x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.21 μs<br>(/C = 2.53x)<br>(/JS = 0.68x) | ~ 0.23 μs<br>(/C = 2.77x)<br>(/JS = 0.74x) | ~ 0.24 μs<br>(/C = 2.89x)<br>(/JS = 0.77x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 37.32 μs<br>(/C = 731.76x)<br>(/JS = 3.89x) | ~ 5.22 μs<br>(/C = 102.35x)<br>(/JS = 0.54x) | ~ 0.86 μs<br>(/C = 16.86x)<br>(/JS = 0.09x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 1341.30 μs<br>(/C = 39.12x)<br>(/JS = 7.08x) | ~ 103.95 μs<br>(/C = 3.03x)<br>(/JS = 0.55x) | ~ 104.62 μs<br>(/C = 3.05x)<br>(/JS = 0.55x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 404.77 μs<br>(/C = 118.18x)<br>(/JS = 5.46x) | ~ 1276.40 μs<br>(/C = 372.67x)<br>(/JS = 17.21x) | ~ 15.46 μs<br>(/C = 4.51x)<br>(/JS = 0.21x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.24 μs<br>(/C = 1.14x)<br>(/JS = 0.29x) | ~ 18.81 μs<br>(/C = 1.11x)<br>(/JS = 0.28x) | ~ 19.68 μs<br>(/C = 1.16x)<br>(/JS = 0.29x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 520.08 μs<br>(/C = 520080.00x)<br>(/JS = 0.41x) | ~ 418.19 μs<br>(/C = 418190.00x)<br>(/JS = 0.33x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 148.04 μs<br>(/C = 135.69x)<br>(/JS = 2.34x) | ~ 30.95 μs<br>(/C = 28.37x)<br>(/JS = 0.49x) | ~ 36.68 μs<br>(/C = 33.62x)<br>(/JS = 0.58x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16911.71 μs<br>(/C = 1.77x)<br>(/JS = 0.19x) | ~ 14282.54 μs<br>(/C = 1.50x)<br>(/JS = 0.16x) | ~ 14426.50 μs<br>(/C = 1.51x)<br>(/JS = 0.16x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 66666.38 μs<br>(/C = 66666380.00x)<br>(/JS = 2.92x) | ~ 52475.29 μs<br>(/C = 52475290.00x)<br>(/JS = 2.30x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 20.41 μs<br>(/C = 20410.00x)<br>(/JS = 0.42x) | ~ 14.55 μs<br>(/C = 14550.00x)<br>(/JS = 0.30x) | ~ 7.61 μs<br>(/C = 7610.00x)<br>(/JS = 0.16x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8739.06 μs<br>(/C = 8739060.00x)<br>(/JS = 0.60x) | ~ 6780.02 μs<br>(/C = 6780020.00x)<br>(/JS = 0.46x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 27.82 μs<br>(/C = 556.40x)<br>(/JS = 3.79x) | ~ 16.21 μs<br>(/C = 324.20x)<br>(/JS = 2.21x) | ~ 1.04 μs<br>(/C = 20.80x)<br>(/JS = 0.14x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.04 μs<br>(/C = 40.00x) | ~ 0.04 μs<br>(/C = 40.00x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = 0.25x) |
**Total Execution Time** | ~ 94.84 ms<br>(/C = 9.87x)<br>(/JS = 0.73x) | ~ 75.42 ms<br>(/C = 7.85x)<br>(/JS = 0.58x) | ~ 14.61 ms<br>(/C = 1.52x)<br>(/JS = 0.11x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.13 μs<br>(/C = 21.96x)<br>(/JS = 8.19x) | ~ 3.19 μs<br>(/C = 32.89x)<br>(/JS = 12.27x) | ~ 2.11 μs<br>(/C = 21.75x)<br>(/JS = 8.12x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 1.08 μs<br>(/C = 13.01x)<br>(/JS = 3.48x) | ~ 2.26 μs<br>(/C = 27.23x)<br>(/JS = 7.29x) | ~ 1.27 μs<br>(/C = 15.30x)<br>(/JS = 4.10x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 137.53 μs<br>(/C = 2696.67x)<br>(/JS = 14.34x) | ~ 174.47 μs<br>(/C = 3420.98x)<br>(/JS = 18.19x) | ~ 0.63 μs<br>(/C = 12.35x)<br>(/JS = 0.07x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 72.18 μs<br>(/C = 2.11x)<br>(/JS = 0.38x) | ~ 49.87 μs<br>(/C = 1.45x)<br>(/JS = 0.26x) | ~ 50.30 μs<br>(/C = 1.47x)<br>(/JS = 0.27x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 1413.71 μs<br>(/C = 412.76x)<br>(/JS = 19.06x) | ~ 1676.49 μs<br>(/C = 489.49x)<br>(/JS = 22.60x) | ~ 491.74 μs<br>(/C = 143.57x)<br>(/JS = 6.63x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 45.89 μs<br>(/C = 2.71x)<br>(/JS = 0.68x) | ~ 100.16 μs<br>(/C = 5.92x)<br>(/JS = 1.48x) | ~ 52.48 μs<br>(/C = 3.10x)<br>(/JS = 0.78x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 1941.41 μs<br>(/C = 1941410.00x)<br>(/JS = 1.54x) | ~ 6824.38 μs<br>(/C = 6824380.00x)<br>(/JS = 5.40x) | ~ 32.84 μs<br>(/C = 32840.00x)<br>(/JS = 0.03x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 698.16 μs<br>(/C = 639.93x)<br>(/JS = 11.01x) | ~ 1022.38 μs<br>(/C = 937.10x)<br>(/JS = 16.13x) | ~ 1.89 μs<br>(/C = 1.73x)<br>(/JS = 0.03x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 53877.92 μs<br>(/C = 5.64x)<br>(/JS = 0.60x) | ~ 509843.88 μs<br>(/C = 53.39x)<br>(/JS = 5.67x) | ~ 113830.46 μs<br>(/C = 11.92x)<br>(/JS = 1.27x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 6237.63 μs<br>(/C = 6237630.00x)<br>(/JS = 0.27x) | ~ 432965.96 μs<br>(/C = 432965960.00x)<br>(/JS = 18.96x) | ~ 87761.96 μs<br>(/C = 87761960.00x)<br>(/JS = 3.84x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 467.12 μs<br>(/C = 467120.00x)<br>(/JS = 9.69x) | ~ 560.06 μs<br>(/C = 560060.00x)<br>(/JS = 11.61x) | ~ 0.48 μs<br>(/C = 480.00x)<br>(/JS = 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 791.27 μs<br>(/C = 791270.00x)<br>(/JS = 0.05x) | ~ 111183.96 μs<br>(/C = 111183960.00x)<br>(/JS = 7.61x) | ~ 326.09 μs<br>(/C = 326090.00x)<br>(/JS = 0.02x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 89.94 μs<br>(/C = 1798.80x)<br>(/JS = 12.24x) | ~ 26.94 μs<br>(/C = 538.80x)<br>(/JS = 3.67x) | ~ 0.80 μs<br>(/C = 16.00x)<br>(/JS = 0.11x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.84 μs<br>(/C = 840.00x)<br>(/JS = 21.00x) | ~ 0.85 μs<br>(/C = 850.00x)<br>(/JS = 21.25x) | ~ 0.08 μs<br>(/C = 80.00x)<br>(/JS = 2.00x) |
**Total Execution Time** | ~ 65.78 ms<br>(/C = 6.84x)<br>(/JS = 0.51x) | ~ 1064.43 ms<br>(/C = 110.76x)<br>(/JS = 8.25x) | ~ 202.55 ms<br>(/C = 21.08x)<br>(/JS = 1.57x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.24 μs<br>(/C = 2.43x)<br>(/JS = 0.91x) | ~ 3.68 μs<br>(/C = 37.94x)<br>(/JS = 14.15x) | ~ 0.23 μs<br>(/C = 2.37x)<br>(/JS = 0.88x) | ~ 0.22 μs<br>(/C = 2.27x)<br>(/JS = 0.85x) | ~ 0.23 μs<br>(/C = 2.37x)<br>(/JS = 0.88x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.09 μs<br>(/C = 1.05x)<br>(/JS = 0.28x) | ~ 11.78 μs<br>(/C = 141.93x)<br>(/JS = 38.00x) | ~ 0.09 μs<br>(/C = 1.08x)<br>(/JS = 0.29x) | ~ 0.12 μs<br>(/C = 1.45x)<br>(/JS = 0.39x) | ~ 0.10 μs<br>(/C = 1.20x)<br>(/JS = 0.32x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.11 μs<br>(/C = 2.24x)<br>(/JS = 0.01x) | ~ 459.33 μs<br>(/C = 9006.47x)<br>(/JS = 47.90x) | ~ 17.07 μs<br>(/C = 334.71x)<br>(/JS = 1.78x) | ~ 11.05 μs<br>(/C = 216.67x)<br>(/JS = 1.15x) | ~ 0.51 μs<br>(/C = 10.00x)<br>(/JS = 0.05x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 29.88 μs<br>(/C = 0.87x)<br>(/JS = 0.16x) | ~ 23547.50 μs<br>(/C = 686.86x)<br>(/JS = 124.27x) | ~ 10.93 μs<br>(/C = 0.32x)<br>(/JS = 0.06x) | ~ 34.70 μs<br>(/C = 1.01x)<br>(/JS = 0.18x) | ~ 34.63 μs<br>(/C = 1.01x)<br>(/JS = 0.18x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.14 μs<br>(/C = 0.92x)<br>(/JS = 0.04x) | ~ 8728.98 μs<br>(/C = 2548.61x)<br>(/JS = 117.67x) | ~ 114.96 μs<br>(/C = 33.56x)<br>(/JS = 1.55x) | ~ 107.97 μs<br>(/C = 31.52x)<br>(/JS = 1.46x) | ~ 3.47 μs<br>(/C = 1.01x)<br>(/JS = 0.05x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.45 μs<br>(/C = 0.97x)<br>(/JS = 0.24x) | ~ 625.00 μs<br>(/C = 36.96x)<br>(/JS = 9.26x) | ~ 15.60 μs<br>(/C = 0.92x)<br>(/JS = 0.23x) | ~ 18.59 μs<br>(/C = 1.10x)<br>(/JS = 0.28x) | ~ 18.36 μs<br>(/C = 1.09x)<br>(/JS = 0.27x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 159.52 μs<br>(/C = 159524.00x)<br>(/JS = 0.13x) | ~ 11792.17 μs<br>(/C = 11792170.00x)<br>(/JS = 9.33x) | ~ 2249.51 μs<br>(/C = 2249510.00x)<br>(/JS = 1.78x) | ~ 1033.78 μs<br>(/C = 1033780.00x)<br>(/JS = 0.82x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 144.00 μs<br>(/C = 131.99x)<br>(/JS = 2.27x) | ~ 2194.96 μs<br>(/C = 2011.88x)<br>(/JS = 34.63x) | ~ 120.43 μs<br>(/C = 110.38x)<br>(/JS = 1.90x) | ~ 78.88 μs<br>(/C = 72.30x)<br>(/JS = 1.24x) | ~ 0.78 μs<br>(/C = 0.71x)<br>(/JS = 0.01x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8136.98 μs<br>(/C = 0.85x)<br>(/JS = 0.09x) | ~ 659207.04 μs<br>(/C = 69.03x)<br>(/JS = 7.33x) | ~ 50862.42 μs<br>(/C = 5.33x)<br>(/JS = 0.57x) | ~ 28737.38 μs<br>(/C = 3.01x)<br>(/JS = 0.32x) | ~ 28247.96 μs<br>(/C = 2.96x)<br>(/JS = 0.31x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.03 μs<br>(/C = 32.00x)<br>(/JS = < 0.01x) | ~ 1588595.46 μs<br>(/C = 1588595460.00x)<br>(/JS = 69.58x) | ~ 7471.33 μs<br>(/C = 7471330.00x)<br>(/JS = 0.33x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 44.57 μs<br>(/C = 44568.00x)<br>(/JS = 0.92x) | ~ 826.12 μs<br>(/C = 826120.00x)<br>(/JS = 17.13x) | ~ 73.12 μs<br>(/C = 73120.00x)<br>(/JS = 1.52x) | ~ 27.36 μs<br>(/C = 27360.00x)<br>(/JS = 0.57x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.01 μs<br>(/C = 11.00x)<br>(/JS = < 0.01x) | ~ 159056.38 μs<br>(/C = 159056380.00x)<br>(/JS = 10.89x) | ~ 38342.83 μs<br>(/C = 38342830.00x)<br>(/JS = 2.62x) | ~ 23675.83 μs<br>(/C = 23675830.00x)<br>(/JS = 1.62x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.03 μs<br>(/C = 0.68x)<br>(/JS = < 0.01x) | ~ 298.70 μs<br>(/C = 5974.00x)<br>(/JS = 40.64x) | ~ 3.28 μs<br>(/C = 65.60x)<br>(/JS = 0.45x) | ~ 0.58 μs<br>(/C = 11.60x)<br>(/JS = 0.08x) | ~ 0.25 μs<br>(/C = 5.00x)<br>(/JS = 0.03x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.07 μs<br>(/C = 68.00x)<br>(/JS = 1.70x) | ~ 0.95 μs<br>(/C = 950.00x)<br>(/JS = 23.75x) | ~ 0.10 μs<br>(/C = 100.00x)<br>(/JS = 2.50x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = 0.50x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = 0.25x) |
**Total Execution Time** | ~ 8.54 ms<br>(/C = 0.89x)<br>(/JS = 0.07x) | ~ 2455.35 ms<br>(/C = 255.50x)<br>(/JS = 19.02x) | ~ 99.28 ms<br>(/C = 10.33x)<br>(/JS = 0.77x) | ~ 53.73 ms<br>(/C = 5.59x)<br>(/JS = 0.42x) | ~ 28.31 ms<br>(/C = 2.95x)<br>(/JS = 0.22x) |
**Excluded WIP** | | | | | |
[Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 6387.79 μs<br>(/C = 0.97x)<br>(/JS = 0.25x) | — | — | — | — |
[JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 35351.17 μs<br>(/C = 3.35x)<br>(/JS = 0.44x) | — | — | — | — |
[JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 1130.69 μs<br>(/C = 1.72x)<br>(/JS = 0.12x) | — | — | — | — |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | pending | — | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | pending | — | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | pending | — | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | pending | — | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | pending | — | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | pending | — | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | pending | — | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | pending | — | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | pending | — | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | pending | — | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | pending | — | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | pending | — | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | pending | — | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | pending | — | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | pending | — | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | pending | — | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | pending | — | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | pending | — | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | pending | — | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | pending | — | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | pending | — | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | pending | — | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | pending | — | — | — | — |

Rows grouped under **Excluded WIP** are not included in **Total Execution Time** or the **/C** ratios.

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written FP-style C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written imperative C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.71 μs<br>(/C = 27.94x)<br>(/JS = 10.42x) | ~ 0.94 μs<br>(/C = 9.69x)<br>(/JS = 3.62x) | ~ 0.85 μs<br>(/C = 8.76x)<br>(/JS = 3.27x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs<br>(/C = 3.73x) | ~ 0.11 μs<br>(/C = 1.33x)<br>(/JS = 0.35x) | ~ 0.02 μs<br>(/C = 0.24x)<br>(/JS = 0.06x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 154.89 μs<br>(/C = 3037.06x)<br>(/JS = 16.15x) | ~ 35.16 μs<br>(/C = 689.41x)<br>(/JS = 3.67x) | ~ 0.36 μs<br>(/C = 7.06x)<br>(/JS = 0.04x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 2844.59 μs<br>(/C = 82.97x)<br>(/JS = 15.01x) | ~ 35.11 μs<br>(/C = 1.02x)<br>(/JS = 0.19x) | ~ 33.85 μs<br>(/C = 0.99x)<br>(/JS = 0.18x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 1453.15 μs<br>(/C = 424.28x)<br>(/JS = 19.59x) | ~ 567.68 μs<br>(/C = 165.75x)<br>(/JS = 7.65x) | ~ 3.41 μs<br>(/C = 1.00x)<br>(/JS = 0.05x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 221.50 μs<br>(/C = 13.10x)<br>(/JS = 3.28x) | ~ 17.99 μs<br>(/C = 1.06x)<br>(/JS = 0.27x) | ~ 16.54 μs<br>(/C = 0.98x)<br>(/JS = 0.25x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 2630.26 μs<br>(/C = 2630260.00x)<br>(/JS = 2.08x) | ~ 4240.89 μs<br>(/C = 4240890.00x)<br>(/JS = 3.36x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = < 0.01x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 974.45 μs<br>(/C = 893.17x)<br>(/JS = 15.37x) | ~ 257.70 μs<br>(/C = 236.21x)<br>(/JS = 4.07x) | ~ 1.08 μs<br>(/C = 0.99x)<br>(/JS = 0.02x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 613388.67 μs<br>(/C = 64.23x)<br>(/JS = 6.82x) | ~ 66049.46 μs<br>(/C = 6.92x)<br>(/JS = 0.73x) | ~ 22007.12 μs<br>(/C = 2.30x)<br>(/JS = 0.24x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 265261 μs<br>(/C = 265261000.00x)<br>(/JS = 11.62x) | ~ 27358.79 μs<br>(/C = 27358790.00x)<br>(/JS = 1.20x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = < 0.01x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 293.30 μs<br>(/C = 293300.00x)<br>(/JS = 6.08x) | ~ 185.70 μs<br>(/C = 185700.00x)<br>(/JS = 3.85x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = < 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 36615.58 μs<br>(/C = 36615580.00x)<br>(/JS = 2.51x) | ~ 31474.58 μs<br>(/C = 31474580.00x)<br>(/JS = 2.15x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 36.57 μs<br>(/C = 731.40x)<br>(/JS = 4.98x) | ~ 1.10 μs<br>(/C = 22.00x)<br>(/JS = 0.15x) | ~ 0.74 μs<br>(/C = 14.80x)<br>(/JS = 0.10x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.48 μs<br>(/C = 480.00x)<br>(/JS = 12.00x) | ~ 0.30 μs<br>(/C = 300.00x)<br>(/JS = 7.50x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = 0.50x) |
**Total Execution Time** | ~ 923.88 ms<br>(/C = 96.14x)<br>(/JS = 7.16x) | ~ 130.23 ms<br>(/C = 13.55x)<br>(/JS = 1.01x) | ~ 22.06 ms<br>(/C = 2.30x)<br>(/JS = 0.17x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 1.12 μs<br>(/C = 11.55x)<br>(/JS = 4.31x) | ~ 0.50 μs<br>(/C = 5.15x)<br>(/JS = 1.92x) | ~ 0.50 μs<br>(/C = 5.15x)<br>(/JS = 1.92x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.36 μs<br>(/C = 4.34x)<br>(/JS = 1.16x) | ~ 0.12 μs<br>(/C = 1.45x)<br>(/JS = 0.39x) | ~ 0.11 μs<br>(/C = 1.33x)<br>(/JS = 0.35x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 560.47 μs<br>(/C = 10989.61x)<br>(/JS = 58.44x) | ~ 16.05 μs<br>(/C = 314.71x)<br>(/JS = 1.67x) | ~ 3.05 μs<br>(/C = 59.80x)<br>(/JS = 0.32x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.45 μs<br>(/C = 1.68x)<br>(/JS = 0.30x) | ~ 37.60 μs<br>(/C = 1.10x)<br>(/JS = 0.20x) | ~ 49.49 μs<br>(/C = 1.44x)<br>(/JS = 0.26x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 2545.27 μs<br>(/C = 743.14x)<br>(/JS = 34.31x) | ~ 87.33 μs<br>(/C = 25.50x)<br>(/JS = 1.18x) | ~ 7.21 μs<br>(/C = 2.11x)<br>(/JS = 0.10x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.15 μs<br>(/C = 0.90x)<br>(/JS = 0.22x) | ~ 7.70 μs<br>(/C = 0.46x)<br>(/JS = 0.11x) | ~ 7.66 μs<br>(/C = 0.45x)<br>(/JS = 0.11x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 1678.49 μs<br>(/C = 1678490.00x)<br>(/JS = 1.33x) | ~ 342.79 μs<br>(/C = 342790.00x)<br>(/JS = 0.27x) | ~ 25.24 μs<br>(/C = 25240.00x)<br>(/JS = 0.02x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 525.55 μs<br>(/C = 481.71x)<br>(/JS = 8.29x) | ~ 45.26 μs<br>(/C = 41.48x)<br>(/JS = 0.71x) | ~ 1.80 μs<br>(/C = 1.65x)<br>(/JS = 0.03x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 19071.54 μs<br>(/C = 2.00x)<br>(/JS = 0.21x) | ~ 36659.92 μs<br>(/C = 3.84x)<br>(/JS = 0.41x) | ~ 9499.96 μs<br>(/C = 0.99x)<br>(/JS = 0.11x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2523.93 μs<br>(/C = 2523930.00x)<br>(/JS = 0.11x) | ~ 23242.04 μs<br>(/C = 23242040.00x)<br>(/JS = 1.02x) | ~ 2254.54 μs<br>(/C = 2254540.00x)<br>(/JS = 0.10x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 169.78 μs<br>(/C = 169780.00x)<br>(/JS = 3.52x) | ~ 78.22 μs<br>(/C = 78220.00x)<br>(/JS = 1.62x) | ~ 3.21 μs<br>(/C = 3210.00x)<br>(/JS = 0.07x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 6576.73 μs<br>(/C = 6576730.00x)<br>(/JS = 0.45x) | ~ 9665.58 μs<br>(/C = 9665580.00x)<br>(/JS = 0.66x) | ~ 235.04 μs<br>(/C = 235040.00x)<br>(/JS = 0.02x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 62.07 μs<br>(/C = 1241.40x)<br>(/JS = 8.44x) | ~ 29.54 μs<br>(/C = 590.80x)<br>(/JS = 4.02x) | ~ 15.63 μs<br>(/C = 312.60x)<br>(/JS = 2.13x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.24 μs<br>(/C = 240.00x)<br>(/JS = 6.00x) | ~ 0.06 μs<br>(/C = 60.00x)<br>(/JS = 1.50x) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = 0.50x) |
**Total Execution Time** | ~ 33.79 ms<br>(/C = 3.52x)<br>(/JS = 0.26x) | ~ 70.21 ms<br>(/C = 7.31x)<br>(/JS = 0.54x) | ~ 12.10 ms<br>(/C = 1.26x)<br>(/JS = 0.09x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.14 μs<br>(/C = 1.44x)<br>(/JS = 0.54x) | ~ 0.11 μs<br>(/C = 1.13x)<br>(/JS = 0.42x) | ~ 0.10 μs<br>(/C = 1.03x)<br>(/JS = 0.38x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.26 μs<br>(/C = 3.13x)<br>(/JS = 0.84x) | ~ 0.10 μs<br>(/C = 1.20x)<br>(/JS = 0.32x) | ~ 0.08 μs<br>(/C = 0.96x)<br>(/JS = 0.26x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 5.03 μs<br>(/C = 98.63x)<br>(/JS = 0.52x) | ~ 4.18 μs<br>(/C = 81.96x)<br>(/JS = 0.44x) | ~ 0.57 μs<br>(/C = 11.18x)<br>(/JS = 0.06x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 45.64 μs<br>(/C = 1.33x)<br>(/JS = 0.24x) | ~ 38.54 μs<br>(/C = 1.12x)<br>(/JS = 0.20x) | ~ 39.83 μs<br>(/C = 1.16x)<br>(/JS = 0.21x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 73.28 μs<br>(/C = 21.40x)<br>(/JS = 0.99x) | ~ 38.80 μs<br>(/C = 11.33x)<br>(/JS = 0.52x) | ~ 4.18 μs<br>(/C = 1.22x)<br>(/JS = 0.06x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 18.18 μs<br>(/C = 1.08x)<br>(/JS = 0.27x) | ~ 7.60 μs<br>(/C = 0.45x)<br>(/JS = 0.11x) | ~ 5.80 μs<br>(/C = 0.34x)<br>(/JS = 0.09x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 310.51 μs<br>(/C = 310510.00x)<br>(/JS = 0.25x) | ~ 375.72 μs<br>(/C = 375720.00x)<br>(/JS = 0.30x) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.29 μs<br>(/C = 19.51x)<br>(/JS = 0.34x) | ~ 18.40 μs<br>(/C = 16.87x)<br>(/JS = 0.29x) | ~ 0.54 μs<br>(/C = 0.49x)<br>(/JS = 0.01x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16788.38 μs<br>(/C = 1.76x)<br>(/JS = 0.19x) | ~ 12065.75 μs<br>(/C = 1.26x)<br>(/JS = 0.13x) | ~ 11828.17 μs<br>(/C = 1.24x)<br>(/JS = 0.13x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = < 0.01x) | ~ 14583.08 μs<br>(/C = 14583080.00x)<br>(/JS = 0.64x) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 6.26 μs<br>(/C = 6260.00x)<br>(/JS = 0.13x) | ~ 17.25 μs<br>(/C = 17250.00x)<br>(/JS = 0.36x) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5.24 μs<br>(/C = 5240.00x)<br>(/JS = < 0.01x) | ~ 8062 μs<br>(/C = 8062000.00x)<br>(/JS = 0.55x) | ~ 0.04 μs<br>(/C = 40.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.96 μs<br>(/C = 59.20x)<br>(/JS = 0.40x) | ~ 3.06 μs<br>(/C = 61.20x)<br>(/JS = 0.42x) | ~ 0.50 μs<br>(/C = 10.00x)<br>(/JS = 0.07x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = 0.25x) | ~ 0.01 μs<br>(/C = 10.00x)<br>(/JS = 0.25x) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
**Total Execution Time** | ~ 17.28 ms<br>(/C = 1.80x)<br>(/JS = 0.13x) | ~ 35.21 ms<br>(/C = 3.66x)<br>(/JS = 0.27x) | ~ 11.88 ms<br>(/C = 1.24x)<br>(/JS = 0.09x) |
**Excluded WIP** | | | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |

#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>koka -O3<br><br>(WIP) |
----------------------- | ----- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.22 μs<br>(/C = 2.27x)<br>(/JS = 0.85x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs<br>(/C = 1.45x)<br>(/JS = 0.39x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 4.28 μs<br>(/C = 83.92x)<br>(/JS = 0.45x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 80.51 μs<br>(/C = 2.35x)<br>(/JS = 0.42x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 158.94 μs<br>(/C = 46.41x)<br>(/JS = 2.14x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.38 μs<br>(/C = 1.15x)<br>(/JS = 0.29x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 758.69 μs<br>(/C = 758690.00x)<br>(/JS = 0.60x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 12.17 μs<br>(/C = 11.15x)<br>(/JS = 0.19x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8746.50 μs<br>(/C = 0.92x)<br>(/JS = 0.10x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 35687 μs<br>(/C = 35687000.00x)<br>(/JS = 1.56x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 5.21 μs<br>(/C = 5210.00x)<br>(/JS = 0.11x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8982 μs<br>(/C = 8982000.00x)<br>(/JS = 0.61x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 6.58 μs<br>(/C = 131.60x)<br>(/JS = 0.90x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
**Total Execution Time** | ~ 54.46 ms<br>(/C = 5.67x)<br>(/JS = 0.42x) |
**Excluded WIP** | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2<br><br>(WIP) |
----------------------- | ------------------------ |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.06 μs<br>(/C = 0.62x)<br>(/JS = 0.23x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.19 μs<br>(/C = 2.29x)<br>(/JS = 0.61x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 3.58 μs<br>(/C = 70.20x)<br>(/JS = 0.37x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 62.84 μs<br>(/C = 1.83x)<br>(/JS = 0.33x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.68 μs<br>(/C = 2.24x)<br>(/JS = 0.10x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 6.75 μs<br>(/C = 0.40x)<br>(/JS = 0.10x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 143.90 μs<br>(/C = 143900.00x)<br>(/JS = 0.11x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.58 μs<br>(/C = 19.78x)<br>(/JS = 0.34x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 11559 μs<br>(/C = 1.21x)<br>(/JS = 0.13x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 4883 μs<br>(/C = 4883000.00x)<br>(/JS = 0.21x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.02 μs<br>(/C = 20.00x)<br>(/JS = < 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.46 μs<br>(/C = 460.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 4.33 μs<br>(/C = 86.60x)<br>(/JS = 0.59x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
**Total Execution Time** | ~ 16.69 ms<br>(/C = 1.74x)<br>(/JS = 0.13x) |
**Excluded WIP** | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3<br><br>(WIP) |
----------------------- | --------------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.05 μs<br>(/C = 0.52x)<br>(/JS = 0.19x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs<br>(/C = 1.45x)<br>(/JS = 0.39x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 1.98 μs<br>(/C = 38.82x)<br>(/JS = 0.21x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 49.82 μs<br>(/C = 1.45x)<br>(/JS = 0.26x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 15.85 μs<br>(/C = 4.63x)<br>(/JS = 0.21x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.05 μs<br>(/C = 0.95x)<br>(/JS = 0.24x) |
[Church Numerals](src/Test/Church/Church.purs) | ~ 140.38 μs<br>(/C = 140380.00x)<br>(/JS = 0.11x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 17.94 μs<br>(/C = 16.44x)<br>(/JS = 0.28x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10787 μs<br>(/C = 1.13x)<br>(/JS = 0.12x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 12448 μs<br>(/C = 12448000.00x)<br>(/JS = 0.55x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 10.03 μs<br>(/C = 10030.00x)<br>(/JS = 0.21x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5690 μs<br>(/C = 5690000.00x)<br>(/JS = 0.39x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 5.89 μs<br>(/C = 117.80x)<br>(/JS = 0.80x) |
[RowToList](src/Test/RowToList/RowToList.purs) | ~ 0 μs<br>(/C = —)<br>(/JS = —) |
**Total Execution Time** | ~ 29.18 ms<br>(/C = 3.04x)<br>(/JS = 0.23x) |
**Excluded WIP** | |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |

#### C (reference)

C Benchmark             | Hand-written imperative C<br>↓<br>clang -O3<br><br>(WIP) |
----------------------- | -------------------- |
[AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.10 μs<br>(/C = 1.00x)<br>(/JS = 0.37x) |
[Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.08 μs<br>(/C = 1.00x)<br>(/JS = 0.27x) |
[List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.05 μs<br>(/C = 1.00x)<br>(/JS = 0.01x) |
[Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 34.28 μs<br>(/C = 1.00x)<br>(/JS = 0.18x) |
[Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.42 μs<br>(/C = 1.00x)<br>(/JS = 0.05x) |
[Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.91 μs<br>(/C = 1.00x)<br>(/JS = 0.25x) |
[Church Numerals](src/Test/Church/Church.purs) | < 0.01 μs<br>(/C = 1.00x)<br>(/JS = < 0.01x) |
[Prime Sieve](src/Test/Primes/Primes.purs) | ~ 1.09 μs<br>(/C = 1.00x)<br>(/JS = 0.02x) |
[Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 9550.00 μs<br>(/C = 1.00x)<br>(/JS = 0.11x) |
[Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | < 0.01 μs<br>(/C = 1.00x)<br>(/JS = < 0.01x) |
[State Monad](src/Test/StateMonad/StateMonad.purs) | < 0.01 μs<br>(/C = 1.00x)<br>(/JS = < 0.01x) |
[Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | < 0.01 μs<br>(/C = 1.00x)<br>(/JS = < 0.01x) |
[Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.05 μs<br>(/C = 1.00x)<br>(/JS = 0.01x) |
[RowToList](src/Test/RowToList/RowToList.purs) | < 0.01 μs<br>(/C = 1.00x)<br>(/JS = 0.03x) |
**Total Execution Time** | ~ 9.61 ms<br>(/C = 1.00x)<br>(/JS = 0.07x) |
**Excluded WIP** | |
[Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 6589.00 μs<br>(/C = 1.00x)<br>(/JS = 0.26x) |
[JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 10560.13 μs<br>(/C = 1.00x)<br>(/JS = 0.13x) |
[JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 657.69 μs<br>(/C = 1.00x)<br>(/JS = 0.07x) |
[array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
[array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
[array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
[array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
[array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
[array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
[foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
[foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
[foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
[string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
[string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
[string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
[string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
[string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
[string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
[string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
[record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
[record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
[record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
[st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
[st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
[effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
[effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |

### Extended benchmark results (I/O, mutability, async)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
[File I/O](srx/Test/FileOps/FileOps.purs) | ~ 466899.83 μs<br>(/C = —) | ~ 477117.79 μs<br>(/C = —)<br>(/JS = 1.02x) | ~ 442703.25 μs<br>(/C = —)<br>(/JS = 0.95x) | ~ 542377.75 μs<br>(/C = —)<br>(/JS = 1.16x) |
[STArray Operations](srx/Test/STArray.purs) | ~ 0.92 μs<br>(/C = —) | ~ 0.54 μs<br>(/C = —)<br>(/JS = 0.59x) | ~ 0.54 μs<br>(/C = —)<br>(/JS = 0.59x) | ~ 0.29 μs<br>(/C = —)<br>(/JS = 0.32x) |
[String Operations](srx/Test/StringOps.purs) | ~ 256.71 μs<br>(/C = —) | ~ 232.50 μs<br>(/C = —)<br>(/JS = 0.91x) | ~ 507.08 μs<br>(/C = —)<br>(/JS = 1.98x) | ~ 414.58 μs<br>(/C = —)<br>(/JS = 1.61x) |
[Aff Operations](srx/Test/AffOperations.purs) | ~ 10125.21 μs<br>(/C = —) | ~ 11581.83 μs<br>(/C = —)<br>(/JS = 1.14x) | ~ 11009.25 μs<br>(/C = —)<br>(/JS = 1.09x) | ~ 10903.21 μs<br>(/C = —)<br>(/JS = 1.08x) |
[Parallelism](srx/Test/Parallelism.purs) | ~ 14275052.04 μs<br>(/C = —) | ~ 14278260.46 μs<br>(/C = —) | ~ 1182150.08 μs<br>(/C = —)<br>(/JS = 0.08x) | ~ 478303.67 μs<br>(/C = —)<br>(/JS = 0.03x) |
**Total Execution Time** | ~ 14752.33 ms<br>(/C = —) | ~ 14767.19 ms<br>(/C = —) | ~ 1636.37 ms<br>(/C = —)<br>(/JS = 0.11x) | ~ 1032.00 ms<br>(/C = —)<br>(/JS = 0.07x) |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with the Go/Rust multicore runtimes.

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
**Total gopurs-*** | ~ 666.09 s | ~ 254.68 s <br>(/JS = 0.38x) | ~ 212.02 s <br>(/JS = 0.32x) | ~ 203.99 s <br>(/JS = 0.31x)
**Total** | ~ 812.48 s | ~ 332.71 s <br>(/JS = 0.41x) | ~ 265.63 s <br>(/JS = 0.33x) | ~ 256.92 s <br>(/JS = 0.32x)

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
**Total purust-*** | ~ 304.03 s | ~ 120.12 s <br>(/JS = 0.40x)
**Total** | ~ 348.34 s | ~ 140.73 s <br>(/JS = 0.40x)
