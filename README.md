# 🌈 PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. This is made possible by the PureScript compiler architecture, which generates an intermediate representation (`CoreFn`) that can be consumed by various backends:

1. **JavaScript (V8)**: The premier asynchronous JavaScript engine ([default backend](https://github.com/purescript/purescript)).
2. **Arista ES (V8)**: A highly optimized, modern ECMAScript backend developed by Arista Networks for extreme performance (via the [`purs-backend-es` backend](https://github.com/aristanetworks/purescript-backend-optimizer)).
3. **Erlang (BEAM)**: The distributed, highly concurrent, and fault tolerant virtual machine (via the [`purerl` backend](https://github.com/purerl/purerl)).
4. **Chez Scheme**: One of the fastest Lisp compilers in the world for highly optimized native execution (via the [`purescm` backend](https://github.com/purescm/purescm)).
5. **Go**: An experimental Ahead-Of-Time (AOT) backend generating native Go binaries (via the experimental local `gopurs` backend).
6. **PHP**: Generating modern PHP 7.4+ syntax (via the experimental local `phpurs` backend). PureScript that transpiles to PHP, and targets 70% of the web (e.g. containerless VPS).
7. **Rust**: An experimental Ahead-Of-Time (AOT) backend generating native Rust binaries (via the experimental local `purust` backend).
8. **WebAssembly GC (Node/V8)**: The experimental local [`purescript-backend-wasm`](https://github.com/purs-wasm/purescript-backend-wasm) backend. The 14 core algorithms and their library operations run inside Wasm; JavaScript provides the benchmark clock, output, and opaque input barrier.

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

### Core stresstest benchmark results (pure computational)

Command: `./bin/run` (Runs all configured core backends, including Wasm). New tests will gradually be added.

Core measurements use numeric results, observable inputs/outputs and ten calibrated batches per process. Published cells use the median of three independent processes. Formatting and build time are excluded. The FP-style FFI mappings and imperative algorithms have distinct roles, documented in the [methodology](docs/benchmark-methodology.md).

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
AST Evaluation          | ~ 93 μs                | ~ 74 μs            | ~ 75 μs       | ~ 99 μs                 |
Fibonacci               | ~ 43 μs                | ~ 46 μs            | ~ 43 μs       | ~ 41 μs                 |
List Processing         | ~ 386 μs               | ~ 368 μs           | ~ 533 μs      | ~ 40 μs                 |
Tail Call Optimization  | ~ 1597 μs              | ~ 1550 μs          | ~ 251 μs      | ~ 592 μs                |
Deep Record Updates     | ~ 433 μs               | ~ 562 μs           | ~ 340 μs      | ~ 203 μs                |
Ackermann               | ~ 211 μs               | ~ 210 μs           | ~ 156 μs      | ~ 188 μs                |
Church Numerals         | ~ 1662 μs              | ~ 1570 μs          | ~ 2077 μs     | ~ 219 μs                |
Prime Sieve             | ~ 725 μs               | ~ 689 μs           | ~ 678 μs      | ~ 47 μs                 |
Red-Black Tree          | ~ 94574 μs             | ~ 53648 μs         | ~ 35485 μs    | ~ 27514 μs              |
Polymorphism            | ~ 9029 μs              | ~ 8111 μs          | ~ 12321 μs    | ~ 4353 μs               |
State Monad             | ~ 425 μs               | ~ 170 μs           | ~ 724 μs      | ~ 39 μs                 |
Lazy Evaluation         | ~ 16372 μs             | ~ 13986 μs         | ~ 20562 μs    | ~ 744 μs                |
Array Processing        | ~ 218 μs               | ~ 222 μs           | ~ 147 μs      | ~ 82 μs                 |
RowToList               | ~ 37 μs                | ~ 17 μs            | ~ 58 μs       | ~ 19 μs                 |
**Total Execution Time**    | ~ 125.77 ms            | ~ 81.21 ms         | ~ 73.45 ms    | ~ 34.19 ms              |

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) |
----------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
AST Evaluation          | ~ 0.79 μs                        | ~ 296 μs                              | ~ 20 μs       | ~ 1 μs                  |
Fibonacci               | ~ 0.62 μs                        | ~ 24 μs                               | ~ 2 μs        | ~ 0.58 μs               |
List Processing         | ~ 19.92 μs                       | ~ 586 μs                              | ~ 121 μs      | ~ 0.92 μs               |
Tail Call Optimization  | ~ 63.83 μs                       | ~ 13354 μs                            | ~ 39 μs       | ~ 42.12 μs              |
Deep Record Updates     | ~ 7.46 μs                        | ~ 9492 μs                             | ~ 373 μs      | ~ 20.96 μs              |
Ackermann               | ~ 18.17 μs                       | ~ 1267 μs                             | ~ 23 μs       | ~ 25.46 μs              |
Church Numerals         | ~ 233.67 μs                      | ~ 5731 μs                             | ~ 1002 μs     | ~ 29.21 μs              |
Prime Sieve             | ~ 92.29 μs                       | ~ 2549 μs                             | ~ 134 μs      | ~ 1.12 μs               |
Red-Black Tree          | ~ 8979.42 μs                     | ~ 870034 μs                           | ~ 26529 μs    | ~ 8752.42 μs            |
Polymorphism            | ~ 2217.29 μs                     | ~ 669801 μs                           | ~ 58848 μs    | ~ 2217.54 μs            |
State Monad             | ~ 114.92 μs                      | ~ 536 μs                              | ~ 127 μs      | ~ 0.67 μs               |
Lazy Evaluation         | ~ 229.79 μs                      | ~ 71586 μs                            | ~ 14816 μs    | ~ 0.79 μs               |
Array Processing        | ~ 14.21 μs                       | ~ 61 μs                               | ~ 7 μs        | ~ 0.79 μs               |
RowToList               | ~ 0.50 μs                        | ~ 2 μs                                | ~ 1 μs        | ~ 0.33 μs               |
**Total Execution Time**| ~ 11.99 ms                       | ~ 1645.32 ms                          | ~ 102.04 ms   | ~ 11.09 ms              |

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written imperative Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation          | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
Fibonacci               | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
List Processing         | ~ 6 μs          | ~ 4 μs           | ~ 2 μs                     |
Tail Call Optimization  | ~ 273 μs        | ~ 292 μs         | ~ 320 μs                   |
Deep Record Updates     | ~ 167 μs        | ~ 49 μs          | ~ 42 μs                    |
Ackermann               | ~ 14 μs         | ~ 10 μs          | ~ 9 μs                     |
Church Numerals         | ~ 231 μs        | ~ 238 μs         | ~ 59 μs                    |
Prime Sieve             | ~ 33 μs         | ~ 23 μs          | ~ 1 μs                     |
Red-Black Tree          | ~ 22524 μs      | ~ 15420 μs       | ~ 11657 μs                 |
Polymorphism            | ~ 16466 μs      | ~ 27543 μs       | ~ 6574 μs                  |
State Monad             | ~ 2 μs          | ~ 2 μs           | ~ 1 μs                     |
Lazy Evaluation         | ~ 2898 μs       | ~ 1757 μs        | ~ 1 μs                     |
Array Processing        | ~ 7 μs          | ~ 4 μs           | ~ 2 μs                     |
RowToList               | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
**Total Execution Time**| ~ 42.62 ms      | ~ 45.34 ms       | ~ 18.67 ms                 |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation          | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
Fibonacci               | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
List Processing         | ~ 36 μs         | ~ 5 μs           | ~ 1 μs                     |
Tail Call Optimization  | ~ 1188 μs       | ~ 115 μs         | ~ 113 μs                   |
Deep Record Updates     | ~ 343 μs        | ~ 1200 μs        | ~ 16 μs                    |
Ackermann               | ~ 23 μs         | ~ 20 μs          | ~ 19 μs                    |
Church Numerals         | ~ 456 μs        | ~ 389 μs         | ~ 0 μs                     |
Prime Sieve             | ~ 134 μs        | ~ 31 μs          | ~ 33 μs                    |
Red-Black Tree          | ~ 16408 μs      | ~ 13689 μs       | ~ 14408 μs                 |
Polymorphism            | ~ 62676 μs      | ~ 50585 μs       | ~ 0 μs                     |
State Monad             | ~ 19 μs         | ~ 12 μs          | ~ 7 μs                     |
Lazy Evaluation         | ~ 8504 μs       | ~ 6788 μs        | ~ 0 μs                     |
Array Processing        | ~ 25 μs         | ~ 14 μs          | ~ 1 μs                     |
RowToList               | ~ 0 μs          | ~ 0 μs           | ~ 0 μs                     |
**Total Execution Time**| ~ 89.81 ms      | ~ 72.85 ms       | ~ 14.60 ms                 |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend-JIT<br><br>(WIP) | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend-JIT<br><br>(WIP) | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend-JIT<br><br>(WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
AST Evaluation          | ~ 5 μs                 | ~ 5 μs         | ~ 4 μs                   |
Fibonacci               | ~ 3 μs                 | ~ 4 μs         | ~ 3 μs                   |
List Processing         | ~ 160 μs               | ~ 164 μs       | ~ 2 μs                   |
Tail Call Optimization  | ~ 79 μs                | ~ 49 μs        | ~ 55 μs                  |
Deep Record Updates     | ~ 1495 μs              | ~ 1697 μs      | ~ 497 μs                 |
Ackermann               | ~ 48 μs                | ~ 98 μs        | ~ 47 μs                  |
Church Numerals         | ~ 2154 μs              | ~ 6984 μs      | ~ 31 μs                  |
Prime Sieve             | ~ 422 μs               | ~ 905 μs       | ~ 3 μs                   |
Red-Black Tree          | ~ 107963 μs            | ~ 511297 μs    | ~ 118278 μs              |
Polymorphism            | ~ 6443 μs              | ~ 425507 μs    | ~ 87442 μs               |
State Monad             | ~ 445 μs               | ~ 537 μs       | ~ 2 μs                   |
Lazy Evaluation         | ~ 903 μs               | ~ 100248 μs    | ~ 307 μs                 |
Array Processing        | ~ 89 μs                | ~ 10 μs        | ~ 2 μs                   |
RowToList               | ~ 2 μs                 | ~ 2 μs         | ~ 1 μs                   |
**Total Execution Time**| ~ 120.21 ms               | ~ 1047.51 ms   | ~ 206.67 ms              |

#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc-O3<br><br>(WIP) | PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc-O3<br><br>(WIP) | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc-O3 | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc-O3<br><br>(WIP) | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc-O3<br><br>(WIP) |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
AST Evaluation          | ~ 0.267 μs              | ~ 4.261 μs                        | ~ 0.262 μs | ~ 3 μs       | ~ 1 μs                 |
Fibonacci               | ~ 0.086 μs              | ~ 12.119 μs                        | ~ 0.107 μs | ~ 2 μs       | ~ 1 μs                 |
List Processing         | ~ 31.318 μs             | ~ 524.905 μs                        | ~ 17.129 μs | ~ 26 μs      | ~ 1 μs                 |
Tail Call Optimization  | ~ 30.352 μs             | ~ 28936.250 μs                        | ~ 11.159 μs | ~ 108 μs     | ~ 35 μs                |
Deep Record Updates     | ~ 3.257 μs              | ~ 10249.292 μs                        | ~ 118.973 μs | ~ 362 μs     | ~ 4 μs                 |
Ackermann               | ~ 16.890 μs             | ~ 700.396 μs                        | ~ 17.876 μs | ~ 39 μs      | ~ 18 μs                |
Church Numerals         | ~ 172.145 μs            | ~ 11497.208 μs                        | ~ 2136.766 μs | ~ 441 μs     | ~ 1 μs                 |
Prime Sieve             | ~ 154.524 μs            | ~ 2491.573 μs                        | ~ 118.545 μs | ~ 194 μs     | ~ 1 μs                 |
Red-Black Tree          | ~ 8600.333 μs           | ~ 723517.000 μs                        | ~ 54844.875 μs | ~ 37225 μs   | ~ 36070 μs             |
Polymorphism            | ~ 0.041 μs              | ~ 1927174.250 μs                        | ~ 7580.479 μs | ~ 6688 μs    | ~ 1 μs                 |
State Monad             | ~ 55.963 μs             | ~ 907.341 μs                        | ~ 83.459 μs | ~ 110 μs     | ~ 1 μs                 |
Lazy Evaluation         | ~ 0.001 μs              | ~ 176696.292 μs                        | ~ 41533.292 μs | ~ 21884 μs   | ~ 0 μs                 |
Array Processing        | ~ 23.496 μs             | ~ 361.768 μs                        | ~ 3.197 μs | ~ 2 μs       | ~ 1 μs                 |
RowToList               | ~ 0.064 μs              | ~ 1.129 μs                        | ~ 0.112 μs | ~ 1 μs       | ~ 0 μs                 |
**Total Execution Time**| ~ 9.09 ms              | ~ 2883.07 ms                        | ~ 106.47 ms | ~ 67.08 ms   | ~ 36.13 ms             |

[^fable-native-rust]: Hand-written, typed [F# source](tmp/fable_rust/NativeBench.fs) translating the PureScript FP kernels, compiled directly by unmodified Fable 5.17.2, then rustc with optimization level 3 and mimalloc. Lists, arrays, Church functions, generic dictionaries, State closures, non-memoizing thunks and recursive row dictionaries follow the [functional reference contract](tmp/fp_reference_contract.md). Measurements report medians of three processes with calibrated batches. Reproduce with `python3 tmp/run_fable_native_benchmark.py --dotnet /path/to/dotnet10/dotnet --update-readme`.

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++-O3<br><br>(WIP) | Hand-written FP-style C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++-O3<br><br>(WIP) | Hand-written imperative C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++-O3<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation          | ~ 3.25 μs                 | ~ 1.29 μs                 | ~ 1.33 μs                       |
Fibonacci               | ~ 0.62 μs                 | ~ 0.42 μs                 | ~ 0.38 μs                       |
List Processing         | ~ 135.62 μs               | ~ 37.21 μs                | ~ 0.71 μs                       |
Tail Call Optimization  | ~ 2673.33 μs              | ~ 34.83 μs                | ~ 34.96 μs                      |
Deep Record Updates     | ~ 1448.08 μs              | ~ 601.38 μs               | ~ 3.92 μs                       |
Ackermann               | ~ 226.67 μs               | ~ 19.83 μs                | ~ 19.33 μs                      |
Church Numerals         | ~ 2532.12 μs              | ~ 4305.33 μs              | ~ 0.42 μs                       |
Prime Sieve             | ~ 904.12 μs               | ~ 256.62 μs               | ~ 1.50 μs                       |
Red-Black Tree          | ~ 632452.83 μs            | ~ 67446.50 μs             | ~ 21264.08 μs                   |
Polymorphism            | ~ 248167.42 μs            | ~ 27229.04 μs             | ~ 0.38 μs                       |
State Monad             | ~ 305.04 μs               | ~ 198.04 μs               | ~ 0.33 μs                       |
Lazy Evaluation         | ~ 36672.83 μs             | ~ 30937.71 μs             | ~ 0.33 μs                       |
Array Processing        | ~ 38.54 μs                | ~ 1.50 μs                 | ~ 1.04 μs                       |
RowToList               | ~ 1.00 μs                 | ~ 0.29 μs                 | ~ 0.29 μs                       |
**Total Execution Time**| ~ 925.56 ms               | ~ 131.07 ms               | ~ 21.33 ms                      |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet-Release<br><br>(WIP) | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet-Release<br><br>(WIP) | Hand-written imperative F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet-Release<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation          | ~ 73.38 μs                | ~ 77.17 μs                | ~ 74.79 μs                      |
Fibonacci               | ~ 1.75 μs                 | ~ 1.13 μs                 | ~ 1.13 μs                       |
List Processing         | ~ 183.79 μs               | ~ 14.00 μs                | ~ 2.92 μs                       |
Tail Call Optimization  | ~ 51.04 μs                | ~ 42.50 μs                | ~ 43.67 μs                      |
Deep Record Updates     | ~ 2150.33 μs              | ~ 86.46 μs                | ~ 16.46 μs                      |
Ackermann               | ~ 92.96 μs                | ~ 45.42 μs                | ~ 45.25 μs                      |
Church Numerals         | ~ 1501.13 μs              | ~ 339.63 μs               | ~ 27.00 μs                      |
Prime Sieve             | ~ 358.38 μs               | ~ 107.46 μs               | ~ 21.54 μs                      |
Red-Black Tree          | ~ 42239.79 μs             | ~ 35344.83 μs             | ~ 9967.83 μs                    |
Polymorphism            | ~ 281291.33 μs            | ~ 23057.75 μs             | ~ 2229.04 μs                    |
State Monad             | ~ 155.17 μs               | ~ 75.96 μs                | ~ 3.42 μs                       |
Lazy Evaluation         | ~ 6316.58 μs              | ~ 6407.75 μs              | ~ 232.96 μs                     |
Array Processing        | ~ 48.25 μs                | ~ 28.46 μs                | ~ 13.08 μs                      |
RowToList               | ~ 0.75 μs                 | ~ 0.71 μs                 | ~ 0.50 μs                       |
**Total Execution Time**| ~ 334.46 ms               | ~ 65.63 ms                | ~ 12.68 ms                      |

#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot-JIT<br><br>(WIP) | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot-JIT<br><br>(WIP) | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot-JIT<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
AST Evaluation          | ~ 56.00 μs                | ~ 78.46 μs               | ~ 108.29 μs                    |
Fibonacci               | ~ 5.25 μs                 | ~ 2.71 μs                | ~ 4.04 μs                      |
List Processing         | ~ 149.33 μs               | ~ 55.79 μs               | ~ 8.88 μs                      |
Tail Call Optimization  | ~ 41.67 μs                | ~ 36.88 μs               | ~ 38.08 μs                     |
Deep Record Updates     | ~ 162.42 μs               | ~ 132.67 μs              | ~ 21.00 μs                     |
Ackermann               | ~ 23.00 μs                | ~ 9.17 μs                | ~ 7.67 μs                      |
Church Numerals         | ~ 533.79 μs               | ~ 433.04 μs              | ~ 13.29 μs                     |
Prime Sieve             | ~ 166.83 μs               | ~ 64.50 μs               | ~ 9.25 μs                      |
Red-Black Tree          | ~ 16185.67 μs             | ~ 12026.29 μs            | ~ 9123.92 μs                   |
Polymorphism            | ~ 2.71 μs                 | ~ 144.63 μs              | ~ 2.92 μs                      |
State Monad             | ~ 61.54 μs                | ~ 24.96 μs               | ~ 6.79 μs                      |
Lazy Evaluation         | ~ 75.92 μs                | ~ 4681.33 μs             | ~ 2.75 μs                      |
Array Processing        | ~ 62.50 μs                | ~ 41.29 μs               | ~ 9.33 μs                      |
RowToList               | ~ 1.88 μs                 | ~ 2.75 μs                | ~ 2.46 μs                      |
**Total Execution Time**| ~ 17.53 ms                | ~ 17.73 ms               | ~ 9.36 ms                      |

#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>Koka<br>↓<br>koka-O3 |
----------------------- | ----- |
AST Evaluation          | ~ 0.243 μs |
Fibonacci               | ~ 0.125 μs |
List Processing         | ~ 4.421 μs |
Tail Call Optimization  | ~ 82.922 μs |
Deep Record Updates     | ~ 164.031 μs |
Ackermann               | ~ 20.072 μs |
Church Numerals         | ~ 757.312 μs |
Prime Sieve             | ~ 12.443 μs |
Red-Black Tree          | ~ 8996.500 μs |
Polymorphism            | ~ 38653.000 μs |
State Monad             | ~ 5.379 μs |
Lazy Evaluation         | ~ 9669.000 μs |
Array Processing        | ~ 6.971 μs |
RowToList               | ~ 0.001 μs |
**Total Execution Time**| ~ 58.37 ms |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>Haskell<br>↓<br>GHC-O2 |
----------------------- | ------------------------ |
AST Evaluation          | ~ 0.057 μs |
Fibonacci               | ~ 0.205 μs |
List Processing         | ~ 3.727 μs |
Tail Call Optimization  | ~ 65.871 μs |
Deep Record Updates     | ~ 8.016 μs |
Ackermann               | ~ 7.072 μs |
Church Numerals         | ~ 144.039 μs |
Prime Sieve             | ~ 21.848 μs |
Red-Black Tree          | ~ 11952.000 μs |
Polymorphism            | ~ 4860.000 μs |
State Monad             | ~ 0.016 μs |
Lazy Evaluation         | ~ 0.492 μs |
Array Processing        | ~ 4.482 μs |
RowToList               | ~ 0.004 μs |
**Total Execution Time**| ~ 17.07 ms |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>OCaml<br>↓<br>ocamlopt-O3 |
----------------------- | --------------------------- |
AST Evaluation          | ~ 0.050 μs |
Fibonacci               | ~ 0.124 μs |
List Processing         | ~ 1.988 μs |
Tail Call Optimization  | ~ 50.148 μs |
Deep Record Updates     | ~ 15.899 μs |
Ackermann               | ~ 16.357 μs |
Church Numerals         | ~ 141.281 μs |
Prime Sieve             | ~ 17.868 μs |
Red-Black Tree          | ~ 10819.000 μs |
Polymorphism            | ~ 12779.000 μs |
State Monad             | ~ 10.165 μs |
Lazy Evaluation         | ~ 5895.000 μs |
Array Processing        | ~ 5.999 μs |
RowToList               | ~ 0.005 μs |
**Total Execution Time**| ~ 29.75 ms |

#### C (reference)

C Benchmark             | Hand-written imperative C<br>↓<br>C<br>↓<br>clang-O3 |
----------------------- | -------------------- |
AST Evaluation          | ~ 0.096 μs |
Fibonacci               | ~ 0.082 μs |
List Processing         | ~ 0.049 μs |
Tail Call Optimization  | ~ 34.084 μs |
Deep Record Updates     | ~ 3.426 μs |
Ackermann               | ~ 16.467 μs |
Church Numerals         | ~ 0.001 μs |
Prime Sieve             | ~ 1.059 μs |
Red-Black Tree          | ~ 9083.000 μs |
Polymorphism            | ~ 0.001 μs |
State Monad             | ~ 0.001 μs |
Lazy Evaluation         | ~ 0.001 μs |
Array Processing        | ~ 0.048 μs |
RowToList               | ~ 0.001 μs |
**Total Execution Time**| ~ 9.14 ms |

The compiled columns measure the supplied PureScript kernels. FP references retain their functional structures; imperative references may change algorithms and allocation strategies. Automatic compiler optimization is allowed. These synthetic results characterize the listed programs and configurations, not the performance of arbitrary applications or the usefulness of FFI.

Core calculations are sequential. The extended Parallelism test below also measures the runtimes' use of multiple CPU cores. See [measurement boundaries, source fidelity and reproduction](docs/benchmark-methodology.md).

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8-JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8-JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc-O3<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
File I/O                | ~ 429223 μs   | ~ 479362 μs    | ~ 427978.75 μs  | ~ 451274 μs     
STArray Operations      | ~ 3 μs        | ~ 3 μs         | ~ 0.33 μs       | ~ 1 μs          
String Operations       | ~ 2 μs        | ~ 2 μs         | ~ 509.38 μs     | ~ 560 μs        
Aff Operations          | ~ 11482 μs    | ~ 11378 μs     | ~ 10094.21 μs   | ~ 10421 μs      
Parallelism             | ~ 15113637 μs | ~ 14690018 μs  | ~ 1183248.83 μs | ~ 423907 μs     
**Total Execution Time**| ~ 15554.34 ms | ~ 15180.76 ms  | ~ 1621.83 ms    | ~ 886.16 ms     

> [!NOTE]
> **Hardware Context**
> To accurately measure multi-core scaling, these extended benchmarks (specifically the 10 concurrent tasks in the *Parallelism* test) were executed on a machine equipped with **10 performance cores** (Apple M4 Pro).

## Repository structure and output files

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The compiled files generated by the backends for our module: `output/` (`.js` files for Node, `.erl`/`.beam` for Erlang, and Scheme libraries/executables).
- The state variables and raw benchmark results: `var/benchmark/`

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
