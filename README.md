# PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. This is made possible by the PureScript compiler architecture, which generates an intermediate representation (`CoreFn`) that can be consumed by various backends:

1. **JavaScript (V8)**: The premier asynchronous JavaScript engine ([default backend](https://github.com/purescript/purescript)).
2. **Arista ES (V8)**: A highly optimized, modern ECMAScript backend developed by Arista Networks for extreme performance (via the [`purs-backend-es` backend](https://github.com/aristanetworks/purescript-backend-optimizer)).
3. **Erlang (BEAM)**: The distributed, highly concurrent, and fault tolerant virtual machine (via the [`purerl` backend](https://github.com/purerl/purerl)).
4. **Chez Scheme**: One of the fastest Lisp compilers in the world for highly optimized native execution (via the [`purescm` backend](https://github.com/purescm/purescm)).
5. **Go**: An experimental Ahead-Of-Time (AOT) backend generating native Go binaries (via the experimental local `gopurs` backend).
6. **PHP**: Generating modern PHP 7.4+ syntax (via the experimental local `phpurs` backend). PureScript that transpiles to PHP, and targets 70% of the web (e.g. containerless VPS).

## Comprehensive benchmarks
The benchmark suite runs a wide variety of computationally intensive tasks: AST evaluation, purely recursive Fibonacci, massive list processing, tail call optimization, deep record updates, Ackermann function, Church numerals, prime sieves, red black tree insertions, heavy polymorphism (type class dictionary lookups), State monad operations, deep lazy evaluation, heavy file I/O (10,000 synchronous writes and reads), and asynchronous `Aff` operations (via the native event loop). These tests apply massive pressure on the call stack, garbage collector, disk I/O, event loop, and runtime execution engine to measure the raw ability of the compiler and the underlying virtual machine.

> [!IMPORTANT]
> **The benchmark is just an excuse**
> The primary objective here is **not** to pit these technologies against each other to declare a performance "winner". In fact, the benchmark itself is just a pretext. The real goal is to prove that we can seamlessly compile and run the exact same, unmodified PureScript code across **6 completely different runtimes** (V8 standard, V8 optimized, BEAM, Lisp, Go, and PHP). The true victory is achieving universal abstraction without sacrificing execution viability.

### Core vs extended tests (`srx/`)
To ensure fair and executable comparisons across all backends, the test suite is split into two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) that run seamlessly on all 6 backends. Executed via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Core benchmark results (pure computational)
Command: `./bin/run` (Runs on all 6 backends). New tests will gradually be added.

```text
=============================================================================================================================================================
CORE BENCHMARK RESULTS (Fibonacci, AST, tail calls, Church, primes, etc.)                                                                      
=============================================================================================================================================================
Benchmark               | JS            | Arista ES      | Go        | Scheme     | Erlang        | PHP
----------------------- | ------------- | -------------- | --------- | ---------- | ------------- | ---------
AST Evaluation          | ~ 93 μs       | ~ 75 μs        | ~ 22 μs   | ~ 9 μs     | ~ 692 μs      | ~ 15 μs
Fibonacci               | ~ 43 μs       | ~ 46 μs        | ~ 2 μs    | ~ 2 μs     | ~ 49 μs       | ~ 505 μs
List Processing         | ~ 386 μs      | ~ 368 μs       | ~ 95 μs   | ~ 10 μs    | ~ 1212 μs     | ~ 1711 μs
Tail Call Optimization  | ~ 1597 μs     | ~ 1550 μs      | ~ 1665 μs | ~ 326 μs   | ~ 1478 μs     | ~ 12278 μs
Deep Record Updates     | ~ 433 μs      | ~ 562 μs       | ~ 1759 μs | ~ 260 μs   | ~ 778 μs      | ~ 2577 μs
Ackermann               | ~ 211 μs      | ~ 210 μs       | ~ 47 μs   | ~ 28 μs    | ~ 57 μs       | ~ 1218 μs
Church Numerals         | ~ 1662 μs     | ~ 1570 μs      | ~ 549 μs  | ~ 369 μs   | ~ 617 μs      | ~ 16291 μs
Prime Sieve             | ~ 725 μs      | ~ 689 μs       | ~ 332 μs  | ~ 76 μs    | ~ 232 μs      | ~ 3526 μs
Red-Black Tree          | ~ 94574 μs    | ~ 53648 μs     | ~ 50820 μs| ~ 25020 μs | ~ 17904 μs    | ~ 1382509 μs
Polymorphism            | ~ 9029 μs     | ~ 8111 μs      | ~ 2259 μs | ~ 17884 μs | ~ 92155 μs    | ~ 9939 μs
State Monad             | ~ 425 μs      | ~ 170 μs       | ~ 63 μs   | ~ 5 μs     | ~ 108 μs      | ~ 1129 μs
Lazy Evaluation         | ~ 16372 μs    | ~ 13986 μs     | ~ 21985 μs| ~ 2868 μs  | ~ 10325 μs    | ~ 102238 μs
Array Processing        | ~ 218 μs      | ~ 222 μs       | ~ 93 μs   | ~ 13 μs    | ~ 5430 μs     | ~ 1441 μs
----------------------- | ------------- | -------------- | --------- | ---------- | ------------- | ---------
Total Execution Time    | ~ 125.77 ms   | ~ 81.21 ms     | ~ 79.69 ms| ~ 46.87 ms | ~ 131.04 ms   | ~ 1535.38 ms
```
> [!NOTE]
> **Single-Threaded Benchmark**
> All benchmarks presented here are strictly **single-threaded**. They measure raw sequential execution speed and do not take into account the powerful multi-threading capabilities inherent to languages like Go or Erlang (BEAM).

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

```text
========================================================================================
EXTENDED BENCHMARK RESULTS (file I/O, regex, STArray, asynchronous Aff)                                                                      
========================================================================================
WIP. Usually +20% everywhere, when this is supported (e.g. JS or Go, but not Scheme)
```

> [!WARNING]
> **About the PHP results**
> Please note that the `phpurs` backend is a brand new, completely experimental, and homemade compiler built entirely from scratch for this repository. Its execution time is not yet representative of PHP's actual performance limit. The currently displayed time (~1.5s) is actively undergoing optimization and debugging!


> [!NOTE]
> **A 40 year old dinosaur steals the show!**
> One of the biggest surprises here is how incredibly performant **Chez Scheme** proves to be. It completely crushes the execution times of modern, highly optimized engines. It is quite a shock to see this "40 year old dinosaur" (first released in 1984!) comfortably outpace standard Node.js, the Arista backend, and Erlang BEAM across a wide array of functional workloads.

## Repository structure and output files

> [!NOTE]
> This repository intentionally omits compiled folders from `.gitignore`. All generated code, dependencies, and compilation artifacts are deliberately committed.

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The compiled files generated by the backends for our module: `output/` (`.js` files for Node, `.erl`/`.beam` for Erlang, and Scheme libraries/executables).
- The state variables and raw benchmark results: `var/benchmark/`

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
