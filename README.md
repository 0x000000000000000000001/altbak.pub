# PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. This is made possible by the PureScript compiler architecture, which generates an intermediate representation (`CoreFn`) that can be consumed by various backends:

The local AOT backends use the fork's enriched typed representation (TAST / `tcorefn`), including type annotations and data/class layouts. Upstream backends such as psgo, pscpp and purs-wasm consume the standard CoreFn expected by their own toolchains.

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

Core tables remeasured on 15 September 2026 after correcting the native workloads and runners. Each of the 30 columns passed all 14 expected outputs in three independent processes. Every process uses three global warmup suites, three warmups per test, then the best of ten measurements. Each table row is the median of those three process results; the total is the sum of its 14 displayed rows (before final rounding).

Rust uses O3/debug=false in all three columns (previous runner default: O1/debug=true); Go and psgo use GOGC=800 without PGO; C++ uses C++11/O3/NDEBUG without LTO. Exact runtimes, flags, source hashes and logs are archived in the [correction report](scratch/benchmark-correction-20260914/report.md), [validated measurements](scratch/benchmark-correction-20260914/results.json) and [reproduction commands](scratch/benchmark-correction-20260914/measurement-plan.json). The [previous README](scratch/benchmark-correction-20260914/previous-readme.md) is historical evidence; differences from invalid workloads or older profiles are not compiler speedups.

Compiled Gopurs was subsequently remeasured in five alternating processes after fixing a monomorphization regression: its column uses the median of those five results. See the [Go regression report](scratch/go-regression-20260915/report.md) and [measurements](scratch/go-regression-20260915/results.json).

#### JavaScript

JS Benchmark            | Compiled JS ([official](https://github.com/purescript/purescript)) | Compiled JS ([Arista](https://github.com/aristanetworks/purescript-backend-optimizer)) | Native FP-style JS FFI (WIP) | Native hand-written JS FFI (WIP) |
--- | --- | --- | --- | --- |
AST Evaluation | ~ 2.29 μs | ~ 1.83 μs | ~ 1.67 μs | ~ 1.92 μs |
Fibonacci | ~ 2.04 μs | ~ 1.67 μs | ~ 1.88 μs | ~ 2.00 μs |
List Processing | ~ 52.04 μs | ~ 12.08 μs | ~ 25.42 μs | ~ 1.17 μs |
Tail Call Optimization | ~ 177.54 μs | ~ 200.13 μs | ~ 50.58 μs | ~ 57.21 μs |
Deep Record Updates | ~ 71.88 μs | ~ 142.17 μs | ~ 45.63 μs | ~ 10.38 μs |
Ackermann | ~ 63.71 μs | ~ 58.58 μs | ~ 36.88 μs | ~ 40.54 μs |
Church Numerals | ~ 1252.08 μs | ~ 1259.46 μs | ~ 1247.75 μs | ~ 35.29 μs |
Prime Sieve | ~ 129.42 μs | ~ 74.08 μs | ~ 83.29 μs | ~ 6.62 μs |
Red-Black Tree | ~ 89448.83 μs | ~ 49863.88 μs | ~ 25283.08 μs | ~ 28190.38 μs |
Polymorphism | ~ 23021.92 μs | ~ 7476.71 μs | ~ 4645.92 μs | ~ 2717.67 μs |
State Monad | ~ 70.29 μs | ~ 14.25 μs | ~ 69.00 μs | ~ 0.96 μs |
Lazy Evaluation | ~ 14776.42 μs | ~ 11184.04 μs | ~ 18282.38 μs | ~ 310.67 μs |
Array Processing | ~ 27.29 μs | ~ 18.75 μs | ~ 10.87 μs | ~ 2.04 μs |
RowToList | ~ 0.58 μs | ~ 0.08 μs | ~ 0.71 μs | ~ 0.38 μs |
**Total Execution Time** | ~ 129.10 ms | ~ 70.31 ms | ~ 49.79 ms | ~ 31.38 ms |

#### Go

Go Benchmark            | Compiled Go ([gopurs](https://github.com/0x000000000000000000001/gopurs), mature WIP) | Compiled Go ([psgo](https://github.com/i-am-the-slime/purescript-native)) | Native FP-style Go FFI (WIP) | Native hand-written Go FFI (WIP) |
--- | --- | --- | --- | --- |
AST Evaluation | ~ 1.00 μs | ~ 4.71 μs | ~ 0.88 μs | ~ 0.88 μs |
Fibonacci | ~ 0.58 μs | ~ 5.21 μs | ~ 0.50 μs | ~ 0.50 μs |
List Processing | ~ 14.62 μs | ~ 235.88 μs | ~ 22.12 μs | ~ 0.79 μs |
Tail Call Optimization | ~ 51.29 μs | ~ 6340.50 μs | ~ 33.12 μs | ~ 33.17 μs |
Deep Record Updates | ~ 6.54 μs | ~ 5549.00 μs | ~ 40.83 μs | ~ 16.50 μs |
Ackermann | ~ 16.96 μs | ~ 396.08 μs | ~ 19.75 μs | ~ 20.00 μs |
Church Numerals | ~ 431.38 μs | ~ 3767.25 μs | ~ 868.71 μs | ~ 22.88 μs |
Prime Sieve | ~ 85.42 μs | ~ 1291.42 μs | ~ 104.83 μs | ~ 1.08 μs |
Red-Black Tree | ~ 23825.71 μs | ~ 894408.75 μs | ~ 24653.71 μs | ~ 9294.42 μs |
Polymorphism | ~ 2221.08 μs | ~ 528156.42 μs | ~ 58008.92 μs | ~ 2324.58 μs |
State Monad | ~ 122.42 μs | ~ 290.29 μs | ~ 39.08 μs | ~ 0.75 μs |
Lazy Evaluation | ~ 229.08 μs | ~ 74337.29 μs | ~ 14243.21 μs | ~ 0.75 μs |
Array Processing | ~ 15.04 μs | ~ 42.54 μs | ~ 2.92 μs | ~ 0.79 μs |
RowToList | ~ 0.50 μs | ~ 0.83 μs | ~ 0.38 μs | ~ 0.33 μs |
**Total Execution Time** | ~ 27.02 ms | ~ 1514.83 ms | ~ 98.04 ms | ~ 11.72 ms |


#### Scheme

Scheme Benchmark        | Compiled Scheme ([pscm](https://github.com/purescm/purescm)) | Native FP-style Scheme FFI (WIP) | Native hand-written Scheme FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
Fibonacci | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
List Processing | ~ 6.00 μs | ~ 4.00 μs | ~ 2.00 μs |
Tail Call Optimization | ~ 276.00 μs | ~ 313.00 μs | ~ 315.00 μs |
Deep Record Updates | ~ 162.00 μs | ~ 47.00 μs | ~ 42.00 μs |
Ackermann | ~ 15.00 μs | ~ 10.00 μs | ~ 9.00 μs |
Church Numerals | ~ 234.00 μs | ~ 229.00 μs | ~ 60.00 μs |
Prime Sieve | ~ 33.00 μs | ~ 23.00 μs | ~ 1.00 μs |
Red-Black Tree | ~ 23619.00 μs | ~ 15929.00 μs | ~ 12050.00 μs |
Polymorphism | ~ 18474.00 μs | ~ 28059.00 μs | ~ 6713.00 μs |
State Monad | ~ 2.00 μs | ~ 2.00 μs | ~ 1.00 μs |
Lazy Evaluation | ~ 2953.00 μs | ~ 1888.00 μs | ~ 0.00 μs |
Array Processing | ~ 7.00 μs | ~ 4.00 μs | ~ 3.00 μs |
RowToList | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
**Total Execution Time** | ~ 45.78 ms | ~ 46.51 ms | ~ 19.20 ms |

#### Erlang

Erlang Benchmark        | Compiled Erlang ([purerl](https://github.com/purerl/purerl)) | Native FP-style Erlang FFI (WIP) | Native hand-written Erlang FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
Fibonacci | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
List Processing | ~ 31.00 μs | ~ 5.00 μs | ~ 1.00 μs |
Tail Call Optimization | ~ 1241.00 μs | ~ 113.00 μs | ~ 91.00 μs |
Deep Record Updates | ~ 407.00 μs | ~ 1278.00 μs | ~ 16.00 μs |
Ackermann | ~ 22.00 μs | ~ 21.00 μs | ~ 22.00 μs |
Church Numerals | ~ 474.00 μs | ~ 392.00 μs | ~ 0.00 μs |
Prime Sieve | ~ 136.00 μs | ~ 35.00 μs | ~ 33.00 μs |
Red-Black Tree | ~ 16852.00 μs | ~ 14300.00 μs | ~ 15173.00 μs |
Polymorphism | ~ 62673.00 μs | ~ 51986.00 μs | ~ 0.00 μs |
State Monad | ~ 20.00 μs | ~ 14.00 μs | ~ 7.00 μs |
Lazy Evaluation | ~ 8768.00 μs | ~ 6739.00 μs | ~ 0.00 μs |
Array Processing | ~ 24.00 μs | ~ 14.00 μs | ~ 1.00 μs |
RowToList | ~ 0.00 μs | ~ 0.00 μs | ~ 0.00 μs |
**Total Execution Time** | ~ 90.65 ms | ~ 74.90 ms | ~ 15.34 ms |

#### PHP

PHP Benchmark           | Compiled PHP ([phpurs](https://github.com/0x000000000000000000001/phpurs), WIP) | Native FP-style PHP FFI (WIP) | Native hand-written PHP FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 5.33 μs | ~ 5.29 μs | ~ 3.71 μs |
Fibonacci | ~ 2.83 μs | ~ 4.13 μs | ~ 3.21 μs |
List Processing | ~ 155.83 μs | ~ 161.46 μs | ~ 2.17 μs |
Tail Call Optimization | ~ 75.71 μs | ~ 52.29 μs | ~ 51.21 μs |
Deep Record Updates | ~ 1515.50 μs | ~ 1747.17 μs | ~ 489.38 μs |
Ackermann | ~ 45.46 μs | ~ 98.50 μs | ~ 51.29 μs |
Church Numerals | ~ 2223.67 μs | ~ 6997.42 μs | ~ 34.42 μs |
Prime Sieve | ~ 438.96 μs | ~ 882.04 μs | ~ 3.25 μs |
Red-Black Tree | ~ 113127.50 μs | ~ 530606.50 μs | ~ 120196.21 μs |
Polymorphism | ~ 6464.46 μs | ~ 435328.13 μs | ~ 89940.38 μs |
State Monad | ~ 445.00 μs | ~ 531.54 μs | ~ 1.79 μs |
Lazy Evaluation | ~ 808.17 μs | ~ 105235.88 μs | ~ 331.17 μs |
Array Processing | ~ 98.96 μs | ~ 11.17 μs | ~ 2.08 μs |
RowToList | ~ 1.38 μs | ~ 2.13 μs | ~ 1.33 μs |
**Total Execution Time** | ~ 125.41 ms | ~ 1081.66 ms | ~ 211.11 ms |

#### Rust

Rust Benchmark          | Compiled Rust ([purust](https://github.com/0x000000000000000000001/purust), WIP) | Native FP-style Rust FFI (WIP) | Native hand-written Rust FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 1.12 μs | ~ 1.29 μs | ~ 1.29 μs |
Fibonacci | ~ 0.88 μs | ~ 1.00 μs | ~ 0.87 μs |
List Processing | ~ 32.92 μs | ~ 11.54 μs | ~ 1.25 μs |
Tail Call Optimization | ~ 29.71 μs | ~ 31.25 μs | ~ 35.75 μs |
Deep Record Updates | ~ 385.38 μs | ~ 98.75 μs | ~ 4.25 μs |
Ackermann | ~ 17.46 μs | ~ 17.38 μs | ~ 19.42 μs |
Church Numerals | ~ 164.42 μs | ~ 1196.54 μs | ~ 0.87 μs |
Prime Sieve | ~ 152.08 μs | ~ 81.46 μs | ~ 1.71 μs |
Red-Black Tree | ~ 8913.29 μs | ~ 31941.79 μs | ~ 31817.46 μs |
Polymorphism | ~ 0.88 μs | ~ 6925.00 μs | ~ 0.88 μs |
State Monad | ~ 55.29 μs | ~ 30.42 μs | ~ 0.75 μs |
Lazy Evaluation | ~ 0.75 μs | ~ 21006.88 μs | ~ 0.75 μs |
Array Processing | ~ 24.33 μs | ~ 1.79 μs | ~ 1.00 μs |
RowToList | ~ 0.83 μs | ~ 0.83 μs | ~ 0.75 μs |
**Total Execution Time** | ~ 9.78 ms | ~ 61.35 ms | ~ 31.89 ms |

#### C++

C++ Benchmark | Compiled C++ ([pscpp](https://github.com/purescript-native/purescript), WIP) | Native FP-style C++ FFI (WIP) | Native hand-written C++ FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 3.12 μs | ~ 1.33 μs | ~ 1.33 μs |
Fibonacci | ~ 0.58 μs | ~ 0.42 μs | ~ 0.33 μs |
List Processing | ~ 137.17 μs | ~ 33.83 μs | ~ 0.75 μs |
Tail Call Optimization | ~ 2713.04 μs | ~ 32.00 μs | ~ 30.71 μs |
Deep Record Updates | ~ 1477.50 μs | ~ 561.08 μs | ~ 3.79 μs |
Ackermann | ~ 222.88 μs | ~ 18.54 μs | ~ 16.96 μs |
Church Numerals | ~ 2609.58 μs | ~ 4372.12 μs | ~ 0.33 μs |
Prime Sieve | ~ 961.29 μs | ~ 267.29 μs | ~ 1.46 μs |
Red-Black Tree | ~ 650654.88 μs | ~ 70032.25 μs | ~ 21051.00 μs |
Polymorphism | ~ 249959.92 μs | ~ 27346.62 μs | ~ 0.38 μs |
State Monad | ~ 300.67 μs | ~ 199.67 μs | ~ 0.33 μs |
Lazy Evaluation | ~ 37244.25 μs | ~ 31668.92 μs | ~ 0.29 μs |
Array Processing | ~ 35.08 μs | ~ 1.46 μs | ~ 1.04 μs |
RowToList | ~ 0.83 μs | ~ 0.29 μs | ~ 0.29 μs |
**Total Execution Time** | ~ 946.32 ms | ~ 134.54 ms | ~ 21.11 ms |

#### F#/C#

F#/C# Benchmark         | Compiled F#/C# ([sharpurs](https://github.com/0x000000000000000000001/sharpurs), WIP) | Native FP-style F#/C# FFI (WIP) | Native hand-written F#/C# FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 76.38 μs | ~ 72.63 μs | ~ 71.67 μs |
Fibonacci | ~ 2.04 μs | ~ 1.46 μs | ~ 1.33 μs |
List Processing | ~ 180.79 μs | ~ 13.58 μs | ~ 3.38 μs |
Tail Call Optimization | ~ 55.17 μs | ~ 37.54 μs | ~ 43.71 μs |
Deep Record Updates | ~ 2171.58 μs | ~ 81.08 μs | ~ 14.88 μs |
Ackermann | ~ 96.25 μs | ~ 42.46 μs | ~ 45.33 μs |
Church Numerals | ~ 1571.88 μs | ~ 339.96 μs | ~ 27.17 μs |
Prime Sieve | ~ 369.83 μs | ~ 100.08 μs | ~ 17.42 μs |
Red-Black Tree | ~ 43232.67 μs | ~ 37477.83 μs | ~ 10625.17 μs |
Polymorphism | ~ 292475.58 μs | ~ 23654.04 μs | ~ 2306.38 μs |
State Monad | ~ 162.46 μs | ~ 79.21 μs | ~ 5.00 μs |
Lazy Evaluation | ~ 6140.04 μs | ~ 13693.58 μs | ~ 236.04 μs |
Array Processing | ~ 55.50 μs | ~ 24.83 μs | ~ 13.75 μs |
RowToList | ~ 0.75 μs | ~ 0.63 μs | ~ 0.67 μs |
**Total Execution Time** | ~ 346.59 ms | ~ 75.62 ms | ~ 13.41 ms |

#### Java

Java Benchmark          | Compiled Java ([javapurs](https://github.com/0x000000000000000000001/javapurs), WIP) | Native FP-style Java FFI (WIP) | Native hand-written Java FFI (WIP) |
--- | --- | --- | --- |
AST Evaluation | ~ 55.17 μs | ~ 81.63 μs | ~ 104.79 μs |
Fibonacci | ~ 5.08 μs | ~ 3.21 μs | ~ 4.00 μs |
List Processing | ~ 149.17 μs | ~ 57.63 μs | ~ 8.17 μs |
Tail Call Optimization | ~ 41.96 μs | ~ 37.25 μs | ~ 38.67 μs |
Deep Record Updates | ~ 158.96 μs | ~ 133.29 μs | ~ 27.17 μs |
Ackermann | ~ 20.63 μs | ~ 10.63 μs | ~ 9.04 μs |
Church Numerals | ~ 527.63 μs | ~ 462.54 μs | ~ 2.92 μs |
Prime Sieve | ~ 176.96 μs | ~ 64.00 μs | ~ 9.21 μs |
Red-Black Tree | ~ 16675.67 μs | ~ 12187.67 μs | ~ 8965.83 μs |
Polymorphism | ~ 2.46 μs | ~ 145.50 μs | ~ 2.67 μs |
State Monad | ~ 66.33 μs | ~ 28.04 μs | ~ 7.13 μs |
Lazy Evaluation | ~ 79.88 μs | ~ 4864.54 μs | ~ 2.54 μs |
Array Processing | ~ 70.92 μs | ~ 40.50 μs | ~ 8.00 μs |
RowToList | ~ 1.79 μs | ~ 2.46 μs | ~ 2.38 μs |
**Total Execution Time** | ~ 18.03 ms | ~ 18.12 ms | ~ 9.19 ms |

#### WebAssembly GC

Wasm Benchmark | Compiled Wasm GC (purs-wasm / Node) |
--- | --- |
AST Evaluation | ~ 0.58 μs |
Fibonacci | ~ 1.25 μs |
List Processing | ~ 15.21 μs |
Tail Call Optimization | ~ 110.42 μs |
Deep Record Updates | ~ 278.88 μs |
Ackermann | ~ 15.46 μs |
Church Numerals | ~ 1491.08 μs |
Prime Sieve | ~ 34.50 μs |
Red-Black Tree | ~ 558526.63 μs |
Polymorphism | ~ 12262.88 μs |
State Monad | ~ 13.25 μs |
Lazy Evaluation | ~ 7035.25 μs |
Array Processing | ~ 46.83 μs |
RowToList | ~ 0.42 μs |
**Total Execution Time** | ~ 579.83 ms |

### Reading the three implementation columns

The compiled column executes code generated from the PureScript tests. The native FP column is a manual translation preserving the corresponding functional structures: closures, persistent data, records and dictionaries. The handwritten column may use imperative updates or algebraic simplifications while preserving the benchmark argument and result. Its timings therefore include algorithmic choices as well as language and runtime costs.

The objective is to identify expensive generated patterns and improve the compiler where representative measurements justify it. A compiled implementation can beat a particular handwritten implementation on this suite. Such a result applies to these programs, inputs, compiler flags and runtime versions; it does not establish a universal advantage over manually optimized code or make computational FFI unnecessary.

These are deliberately stressful microbenchmarks. The core algorithms run sequentially; runtimes may still use background threads for garbage collection or compilation. The extended Parallelism test below is explicitly concurrent. Values near the clock resolution, especially displayed zeroes, cannot support precise speedup ratios.

### Extended benchmark results (I/O, mutability, async)

These are historical measurements, outside the September 2026 audit of the 14 core tests. They have not been revalidated with the corrected core protocol.
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

#### Extended Results

Benchmark               | Compiled JS ([official](https://github.com/purescript/purescript)) | Compiled JS ([Arista](https://github.com/aristanetworks/purescript-backend-optimizer)) | Compiled Go ([gopurs](https://github.com/0x000000000000000000001/gopurs), mature WIP) | Compiled Rust ([purust](https://github.com/0x000000000000000000001/purust), WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
File I/O                | ~ 429223 μs   | ~ 479362 μs    | ~ 476440 μs     | ~ 451274 μs     
STArray Operations      | ~ 3 μs        | ~ 3 μs         | ~ 0 μs          | ~ 1 μs          
String Operations       | ~ 2 μs        | ~ 2 μs         | ~ 1 μs          | ~ 560 μs        
Aff Operations          | ~ 11482 μs    | ~ 11378 μs     | ~ 11030 μs      | ~ 10421 μs      
Parallelism             | ~ 15113637 μs | ~ 14690018 μs  | ~ 1255501 μs    | ~ 423907 μs     
**Total Execution Time**| ~ 15554.34 ms | ~ 15180.76 ms  | ~ 1742.97 ms    | ~ 886.16 ms     

> [!NOTE]
> **Hardware Context**
> To accurately measure multi-core scaling, these extended benchmarks (specifically the 10 concurrent tasks in the *Parallelism* test) were executed on a machine equipped with **10 performance cores** (Apple M4 Pro).

## Repository structure and output files

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The isolated build directory printed by each backend runner, containing source snapshots, generated code, build logs and a manifest. `--build-only` creates an artifact; `--run-only` executes a saved artifact with output validation.
- The validated measurement logs and reproducible commands linked alongside the core tables. Historical runner outputs remain in `var/benchmark/`.

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
