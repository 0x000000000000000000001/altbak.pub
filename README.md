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

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
AST Evaluation | ~ 0.26 μs | ~ 0.19 μs | ~ 0.20 μs | ~ 0.22 μs |
Fibonacci | ~ 0.31 μs | ~ 0.31 μs | ~ 0.31 μs | ~ 0.37 μs |
List Processing | ~ 9.59 μs | ~ 5.06 μs | ~ 19.06 μs | ~ 0.49 μs |
Tail Call Optimization | ~ 189.49 μs | ~ 151.13 μs | ~ 65.91 μs | ~ 68.31 μs |
Deep Record Updates | ~ 74.18 μs | ~ 134.11 μs | ~ 42.07 μs | ~ 7.69 μs |
Ackermann | ~ 67.49 μs | ~ 64.40 μs | ~ 38.25 μs | ~ 38.48 μs |
Church Numerals | ~ 1263.82 μs | ~ 1242.29 μs | ~ 1337.51 μs | ~ 49.63 μs |
Prime Sieve | ~ 63.39 μs | ~ 26.99 μs | ~ 86.33 μs | ~ 0.65 μs |
Red-Black Tree | ~ 89940.58 μs | ~ 47272.88 μs | ~ 24486.83 μs | ~ 27249.58 μs |
Polymorphism | ~ 22830.25 μs | ~ 7368.38 μs | ~ 4843 μs | ~ 2805.25 μs |
State Monad | ~ 48.23 μs | ~ 10.95 μs | ~ 46.87 μs | ~ 0.42 μs |
Lazy Evaluation | ~ 14608.29 μs | ~ 10604.88 μs | ~ 17513.12 μs | ~ 321.48 μs |
Array Processing | ~ 7.35 μs | ~ 3.75 μs | ~ 3.12 μs | ~ 0.50 μs |
RowToList | ~ 0.04 μs | ~ 0.02 μs | ~ 0.06 μs | ~ 0.02 μs |
Array Indexing (excluded WIP) | ~ 25.30 ms | — | — | — |
JSON to Typed AST (excluded WIP) | ~ 78.19 ms | — | — | — |
JSON Decoding (excluded WIP) | ~ 9.28 ms | — | — | — |
**Total Execution Time** | ~ 129.10 ms <br>(/C = 13.1x) | ~ 66.89 ms <br>(/C = 6.8x) | ~ 48.48 ms <br>(/C = 4.9x) | ~ 30.54 ms <br>(/C = 3.1x) |

Rows marked **excluded WIP** are not included in **Total Execution Time** or the **/C** ratios.

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) |
----------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
AST Evaluation | ~ 0.33 μs | ~ 4.17 μs | ~ 0.32 μs | ~ 0.32 μs |
Fibonacci | ~ 0.25 μs | ~ 5.41 μs | ~ 0.16 μs | ~ 0.16 μs |
List Processing | ~ 17.19 μs | ~ 236.01 μs | ~ 24.39 μs | ~ 0.45 μs |
Tail Call Optimization | ~ 57.18 μs | ~ 6281.96 μs | ~ 35.98 μs | ~ 35.50 μs |
Deep Record Updates | ~ 6.63 μs | ~ 5461.52 μs | ~ 40.43 μs | ~ 14.87 μs |
Ackermann | ~ 15.93 μs | ~ 402.04 μs | ~ 20.85 μs | ~ 19.87 μs |
Church Numerals | ~ 230.56 μs | ~ 3811.48 μs | ~ 1044.60 μs | ~ 22.45 μs |
Prime Sieve | ~ 79.35 μs | ~ 1340.99 μs | ~ 110.87 μs | ~ 0.62 μs |
Red-Black Tree | ~ 10161.25 μs | ~ 870404.62 μs | ~ 24630.67 μs | ~ 9004.48 μs |
Polymorphism | ~ 2292.36 μs | ~ 505069 μs | ~ 49160.92 μs | ~ 2244.45 μs |
State Monad | ~ 150.94 μs | ~ 303.56 μs | ~ 52.10 μs | ~ 0.33 μs |
Lazy Evaluation | ~ 247.11 μs | ~ 71378.08 μs | ~ 13865.38 μs | ~ 0.27 μs |
Array Processing | ~ 19.26 μs | ~ 43.81 μs | ~ 4.90 μs | ~ 0.45 μs |
RowToList | ~ 0.15 μs | ~ 0.50 μs | ~ 0.07 μs | ~ 0.03 μs |
Array Indexing (excluded WIP) | ~ 3.96 ms | — | — | — |
JSON to Typed AST (excluded WIP) | ~ 19.98 ms | — | — | — |
JSON Decoding (excluded WIP) | ~ 2.24 ms | — | — | — |
**Total Execution Time** | ~ 13.28 ms <br>(/C = 1.3x) | ~ 1464.74 ms <br>(/C = 148.8x) | ~ 88.99 ms <br>(/C = 9.0x) | ~ 11.34 ms <br>(/C = 1.2x) |

Rows marked **excluded WIP** are not included in **Total Execution Time** or the **/C** ratios.

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation | ~ 0.07 μs | ~ 0.08 μs | ~ 0.07 μs |
Fibonacci | ~ 0.15 μs | ~ 0.15 μs | ~ 0.14 μs |
List Processing | ~ 6.35 μs | ~ 3.76 μs | ~ 2.84 μs |
Tail Call Optimization | ~ 277.12 μs | ~ 322.34 μs | ~ 329.88 μs |
Deep Record Updates | ~ 228.44 μs | ~ 48.07 μs | ~ 43.57 μs |
Ackermann | ~ 16.99 μs | ~ 9.07 μs | ~ 9.14 μs |
Church Numerals | ~ 237.12 μs | ~ 235.53 μs | ~ 62.26 μs |
Prime Sieve | ~ 37.49 μs | ~ 22.68 μs | ~ 0.80 μs |
Red-Black Tree | ~ 22551 μs | ~ 15428 μs | ~ 11970 μs |
Polymorphism | ~ 17812 μs | ~ 27373 μs | ~ 6730.50 μs |
State Monad | ~ 1.70 μs | ~ 2.02 μs | ~ 0.84 μs |
Lazy Evaluation | ~ 2989.50 μs | ~ 1964 μs | ~ 0.61 μs |
Array Processing | ~ 7.29 μs | ~ 5.04 μs | ~ 2.79 μs |
RowToList | ~ 0 μs | ~ 0.01 μs | ~ 0 μs |
**Total Execution Time** | ~ 44.17 ms <br>(/C = 4.5x) | ~ 45.41 ms <br>(/C = 4.6x) | ~ 19.15 ms <br>(/C = 1.9x) |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation | ~ 0.11 μs | ~ 0.11 μs | ~ 0.10 μs |
Fibonacci | ~ 0.21 μs | ~ 0.23 μs | ~ 0.24 μs |
List Processing | ~ 37.32 μs | ~ 5.22 μs | ~ 0.86 μs |
Tail Call Optimization | ~ 1341.30 μs | ~ 103.95 μs | ~ 104.62 μs |
Deep Record Updates | ~ 404.77 μs | ~ 1276.40 μs | ~ 15.46 μs |
Ackermann | ~ 19.24 μs | ~ 18.81 μs | ~ 19.68 μs |
Church Numerals | ~ 520.08 μs | ~ 418.19 μs | ~ 0.01 μs |
Prime Sieve | ~ 148.04 μs | ~ 30.95 μs | ~ 36.68 μs |
Red-Black Tree | ~ 16911.71 μs | ~ 14282.54 μs | ~ 14426.50 μs |
Polymorphism | ~ 66666.38 μs | ~ 52475.29 μs | ~ 0.01 μs |
State Monad | ~ 20.41 μs | ~ 14.55 μs | ~ 7.61 μs |
Lazy Evaluation | ~ 8739.06 μs | ~ 6780.02 μs | ~ 0.01 μs |
Array Processing | ~ 27.82 μs | ~ 16.21 μs | ~ 1.04 μs |
RowToList | ~ 0.04 μs | ~ 0.04 μs | ~ 0.01 μs |
**Total Execution Time** | ~ 94.84 ms <br>(/C = 9.6x) | ~ 75.42 ms <br>(/C = 7.7x) | ~ 14.61 ms <br>(/C = 1.5x) |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
AST Evaluation | ~ 2.30 μs | ~ 3.19 μs | ~ 2.11 μs |
Fibonacci | ~ 1.27 μs | ~ 2.26 μs | ~ 1.27 μs |
List Processing | ~ 159.73 μs | ~ 174.47 μs | ~ 0.63 μs |
Tail Call Optimization | ~ 76.11 μs | ~ 49.87 μs | ~ 50.30 μs |
Deep Record Updates | ~ 1541.84 μs | ~ 1676.49 μs | ~ 491.74 μs |
Ackermann | ~ 49.08 μs | ~ 100.16 μs | ~ 52.48 μs |
Church Numerals | ~ 2161.17 μs | ~ 6824.38 μs | ~ 32.84 μs |
Prime Sieve | ~ 787.77 μs | ~ 1022.38 μs | ~ 1.89 μs |
Red-Black Tree | ~ 58291.04 μs | ~ 509843.88 μs | ~ 113830.46 μs |
Polymorphism | ~ 6615.56 μs | ~ 432965.96 μs | ~ 87761.96 μs |
State Monad | ~ 501.49 μs | ~ 560.06 μs | ~ 0.48 μs |
Lazy Evaluation | ~ 810.15 μs | ~ 111183.96 μs | ~ 326.09 μs |
Array Processing | ~ 101.22 μs | ~ 26.94 μs | ~ 0.80 μs |
RowToList | ~ 0.87 μs | ~ 0.85 μs | ~ 0.08 μs |
**Total Execution Time** | ~ 71.10 ms <br>(/C = 7.2x) | ~ 1064.43 ms <br>(/C = 108.2x) | ~ 202.55 ms <br>(/C = 20.6x) |

#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP) |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
AST Evaluation | ~ 0.25 μs | ~ 3.68 μs | ~ 0.23 μs | ~ 0.22 μs | ~ 0.23 μs |
Fibonacci | ~ 0.10 μs | ~ 11.78 μs | ~ 0.09 μs | ~ 0.12 μs | ~ 0.10 μs |
List Processing | ~ 18.57 μs | ~ 459.33 μs | ~ 17.07 μs | ~ 11.05 μs | ~ 0.51 μs |
Tail Call Optimization | ~ 30.24 μs | ~ 23547.50 μs | ~ 10.93 μs | ~ 34.70 μs | ~ 34.63 μs |
Deep Record Updates | ~ 3.27 μs | ~ 8728.98 μs | ~ 114.96 μs | ~ 107.97 μs | ~ 3.47 μs |
Ackermann | ~ 16.66 μs | ~ 625.00 μs | ~ 15.60 μs | ~ 18.59 μs | ~ 18.36 μs |
Church Numerals | ~ 169.57 μs | ~ 11792.17 μs | ~ 2249.51 μs | ~ 1033.78 μs | ~ 0.01 μs |
Prime Sieve | ~ 146.47 μs | ~ 2194.96 μs | ~ 120.43 μs | ~ 78.88 μs | ~ 0.78 μs |
Red-Black Tree | ~ 8449.35 μs | ~ 659207.04 μs | ~ 50862.42 μs | ~ 28737.38 μs | ~ 28247.96 μs |
Polymorphism | ~ 0.05 μs | ~ 1588595.46 μs | ~ 7471.33 μs | ~ 0.01 μs | ~ 0.01 μs |
State Monad | ~ 46.91 μs | ~ 826.12 μs | ~ 73.12 μs | ~ 27.36 μs | ~ 0.01 μs |
Lazy Evaluation | ~ 0.01 μs | ~ 159056.38 μs | ~ 38342.83 μs | ~ 23675.83 μs | ~ 0.01 μs |
Array Processing | ~ 0.04 μs | ~ 298.70 μs | ~ 3.28 μs | ~ 0.58 μs | ~ 0.25 μs |
RowToList | ~ 0.07 μs | ~ 0.95 μs | ~ 0.10 μs | ~ 0.02 μs | ~ 0.01 μs |
**Total Execution Time** | ~ 8.88 ms <br>(/C = 0.9x) | ~ 2455.35 ms <br>(/C = 249.5x) | ~ 99.28 ms <br>(/C = 10.1x) | ~ 53.73 ms <br>(/C = 5.5x) | ~ 28.31 ms <br>(/C = 2.9x) |

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written FP-style C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written imperative C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation | ~ 2.71 μs | ~ 0.94 μs | ~ 0.85 μs |
Fibonacci | ~ 0.31 μs | ~ 0.11 μs | ~ 0.02 μs |
List Processing | ~ 154.89 μs | ~ 35.16 μs | ~ 0.36 μs |
Tail Call Optimization | ~ 2844.59 μs | ~ 35.11 μs | ~ 33.85 μs |
Deep Record Updates | ~ 1453.15 μs | ~ 567.68 μs | ~ 3.41 μs |
Ackermann | ~ 221.50 μs | ~ 17.99 μs | ~ 16.54 μs |
Church Numerals | ~ 2630.26 μs | ~ 4240.89 μs | ~ 0.02 μs |
Prime Sieve | ~ 974.45 μs | ~ 257.70 μs | ~ 1.08 μs |
Red-Black Tree | ~ 613388.67 μs | ~ 66049.46 μs | ~ 22007.12 μs |
Polymorphism | ~ 265261 μs | ~ 27358.79 μs | ~ 0.02 μs |
State Monad | ~ 293.30 μs | ~ 185.70 μs | ~ 0.02 μs |
Lazy Evaluation | ~ 36615.58 μs | ~ 31474.58 μs | ~ 0.02 μs |
Array Processing | ~ 36.57 μs | ~ 1.10 μs | ~ 0.74 μs |
RowToList | ~ 0.48 μs | ~ 0.30 μs | ~ 0.02 μs |
**Total Execution Time** | ~ 923.88 ms <br>(/C = 93.9x) | ~ 130.23 ms <br>(/C = 13.2x) | ~ 22.06 ms <br>(/C = 2.2x) |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation | ~ 1.84 μs | ~ 0.49 μs | ~ 0.50 μs |
Fibonacci | ~ 2.44 μs | ~ 0.12 μs | ~ 0.12 μs |
List Processing | ~ 491.87 μs | ~ 16.25 μs | ~ 2.96 μs |
Tail Call Optimization | ~ 43.41 μs | ~ 38.43 μs | ~ 44.32 μs |
Deep Record Updates | ~ 3026.21 μs | ~ 91.30 μs | ~ 9.13 μs |
Ackermann | ~ 173.72 μs | ~ 40.02 μs | ~ 39.78 μs |
Church Numerals | ~ 1787.07 μs | ~ 352.32 μs | ~ 25.65 μs |
Prime Sieve | ~ 607.84 μs | ~ 157.57 μs | ~ 15.45 μs |
Red-Black Tree | ~ 41631.79 μs | ~ 37763.92 μs | ~ 10551.67 μs |
Polymorphism | ~ 2452.14 μs | ~ 25932.33 μs | ~ 2332.77 μs |
State Monad | ~ 308.24 μs | ~ 79.71 μs | ~ 3.32 μs |
Lazy Evaluation | ~ 6758.21 μs | ~ 10061 μs | ~ 233.03 μs |
Array Processing | ~ 59.64 μs | ~ 31.74 μs | ~ 16.09 μs |
RowToList | ~ 0.32 μs | ~ 0.06 μs | ~ 0.03 μs |
**Total Execution Time** | ~ 57.34 ms <br>(/C = 5.8x) | ~ 74.57 ms <br>(/C = 7.6x) | ~ 13.27 ms <br>(/C = 1.3x) |

#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
AST Evaluation | ~ 0.14 μs | ~ 0.11 μs | ~ 0.10 μs |
Fibonacci | ~ 0.26 μs | ~ 0.10 μs | ~ 0.08 μs |
List Processing | ~ 5.03 μs | ~ 4.18 μs | ~ 0.57 μs |
Tail Call Optimization | ~ 45.64 μs | ~ 38.54 μs | ~ 39.83 μs |
Deep Record Updates | ~ 73.28 μs | ~ 38.80 μs | ~ 4.18 μs |
Ackermann | ~ 18.18 μs | ~ 7.60 μs | ~ 5.80 μs |
Church Numerals | ~ 310.51 μs | ~ 375.72 μs | ~ 0 μs |
Prime Sieve | ~ 21.29 μs | ~ 18.40 μs | ~ 0.54 μs |
Red-Black Tree | ~ 16788.38 μs | ~ 12065.75 μs | ~ 11828.17 μs |
Polymorphism | ~ 0.01 μs | ~ 14583.08 μs | ~ 0 μs |
State Monad | ~ 6.26 μs | ~ 17.25 μs | ~ 0 μs |
Lazy Evaluation | ~ 5.24 μs | ~ 8062 μs | ~ 0.04 μs |
Array Processing | ~ 2.96 μs | ~ 3.06 μs | ~ 0.50 μs |
RowToList | ~ 0.01 μs | ~ 0.01 μs | ~ 0 μs |
**Total Execution Time** | ~ 17.28 ms <br>(/C = 1.8x) | ~ 35.21 ms <br>(/C = 3.6x) | ~ 11.88 ms <br>(/C = 1.2x) |

#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>koka -O3<br><br>(WIP) |
----------------------- | ----- |
AST Evaluation | ~ 0.22 μs |
Fibonacci | ~ 0.12 μs |
List Processing | ~ 4.28 μs |
Tail Call Optimization | ~ 80.51 μs |
Deep Record Updates | ~ 158.94 μs |
Ackermann | ~ 19.38 μs |
Church Numerals | ~ 758.69 μs |
Prime Sieve | ~ 12.17 μs |
Red-Black Tree | ~ 8746.50 μs |
Polymorphism | ~ 35687 μs |
State Monad | ~ 5.21 μs |
Lazy Evaluation | ~ 8982 μs |
Array Processing | ~ 6.58 μs |
RowToList | ~ 0 μs |
**Total Execution Time** | ~ 54.46 ms <br>(/C = 5.5x) |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2<br><br>(WIP) |
----------------------- | ------------------------ |
AST Evaluation | ~ 0.06 μs |
Fibonacci | ~ 0.19 μs |
List Processing | ~ 3.58 μs |
Tail Call Optimization | ~ 62.84 μs |
Deep Record Updates | ~ 7.68 μs |
Ackermann | ~ 6.75 μs |
Church Numerals | ~ 143.90 μs |
Prime Sieve | ~ 21.58 μs |
Red-Black Tree | ~ 11559 μs |
Polymorphism | ~ 4883 μs |
State Monad | ~ 0.02 μs |
Lazy Evaluation | ~ 0.46 μs |
Array Processing | ~ 4.33 μs |
RowToList | ~ 0 μs |
**Total Execution Time** | ~ 16.69 ms <br>(/C = 1.7x) |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3<br><br>(WIP) |
----------------------- | --------------------------- |
AST Evaluation | ~ 0.05 μs |
Fibonacci | ~ 0.12 μs |
List Processing | ~ 1.98 μs |
Tail Call Optimization | ~ 49.82 μs |
Deep Record Updates | ~ 15.85 μs |
Ackermann | ~ 16.05 μs |
Church Numerals | ~ 140.38 μs |
Prime Sieve | ~ 17.94 μs |
Red-Black Tree | ~ 10787 μs |
Polymorphism | ~ 12448 μs |
State Monad | ~ 10.03 μs |
Lazy Evaluation | ~ 5690 μs |
Array Processing | ~ 5.89 μs |
RowToList | ~ 0 μs |
**Total Execution Time** | ~ 29.18 ms <br>(/C = 3.0x) |

#### C (reference)

C Benchmark             | Hand-written imperative C<br>↓<br>clang -O3<br><br>(WIP) |
----------------------- | -------------------- |
AST Evaluation | ~ 0.10 μs |
Fibonacci | ~ 0.08 μs |
List Processing | ~ 0.05 μs |
Tail Call Optimization | ~ 33.78 μs |
Deep Record Updates | ~ 3.41 μs |
Ackermann | ~ 17.35 μs |
Church Numerals | ~ 0 μs |
Prime Sieve | ~ 1.07 μs |
Red-Black Tree | ~ 9788 μs |
Polymorphism | ~ 0 μs |
State Monad | ~ 0 μs |
Lazy Evaluation | ~ 0 μs |
Array Processing | ~ 0.05 μs |
RowToList | ~ 0 μs |
Array Indexing (excluded WIP) | ~ 6589.00 μs |
JSON to Typed AST (excluded WIP) | ~ 10823.13 μs |
JSON Decoding (excluded WIP) | ~ 657.69 μs |
**Total Execution Time** | ~ 9.84 ms |

### Extended benchmark results (I/O, mutability, async)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
File I/O | ~ 466899.83 μs | ~ 477117.79 μs | ~ 442703.25 μs | ~ 542377.75 μs |
STArray Operations | ~ 0.92 μs | ~ 0.54 μs | ~ 0.54 μs | ~ 0.29 μs |
String Operations | ~ 256.71 μs | ~ 232.50 μs | ~ 507.08 μs | ~ 414.58 μs |
Aff Operations | ~ 10125.21 μs | ~ 11581.83 μs | ~ 11009.25 μs | ~ 10903.21 μs |
Parallelism | ~ 14275052.04 μs | ~ 14278260.46 μs | ~ 1182150.08 μs | ~ 478303.67 μs |
**Total Execution Time** | ~ 14752.33 ms | ~ 14767.19 ms | ~ 1636.37 ms | ~ 1032.00 ms |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with Go/Rust's multicore runtimes.

### Compilation times

#### [gopurs](https://github.com/0x000000000000000000001/gopurs)

Benchmark   | gopurs.js | gopurs Go compiled
----------- | --------- | ------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | (WIP)     | (WIP)
[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff)  | ~ 7511 ms | ~ 4321 ms
[gopurs-argonaut-core](https://github.com/0x000000000000000000001/gopurs-argonaut-core)  | ~ 7411 ms | ~ 3344 ms
[gopurs-arrays](https://github.com/0x000000000000000000001/gopurs-arrays)  | ~ 6714 ms (failed) | ~ 4782 ms
[gopurs-assert](https://github.com/0x000000000000000000001/gopurs-assert)  | ~ 600 ms | ~ 316 ms
[gopurs-avar](https://github.com/0x000000000000000000001/gopurs-avar)  | ~ 6713 ms | ~ 3524 ms

> [!NOTE]
> On `gopurs-arrays`, the `gopurs.js` compiler crashes with `RangeError: Maximum call stack size exceeded` during `optimize + emit`; the Go-compiled compiler completes the same build. The reported time is the backend total at failure.

#### [purust](https://github.com/0x000000000000000000001/purust)

Benchmark   | purust.js | purust Rust compiled
----------- | --------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | (WIP)     | (WIP)
[purust-aff](https://github.com/0x000000000000000000001/purust-aff)  | (WIP)     | (WIP)
