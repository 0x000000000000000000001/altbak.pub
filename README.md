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
AST Evaluation          | ~ 93 μs       | ~ 74 μs        | ~ 4 μs    | ~ 9 μs     | ~ 692 μs      | ~ 18 μs
Fibonacci               | ~ 43 μs       | ~ 46 μs        | ~ 2 μs    | ~ 2 μs     | ~ 49 μs       | ~ 354 μs
List Processing         | ~ 386 μs      | ~ 368 μs       | ~ 124 μs  | ~ 10 μs    | ~ 1212 μs     | ~ 2189 μs
Tail Call Optimization  | ~ 1597 μs     | ~ 1550 μs      | ~ 1106 μs | ~ 326 μs   | ~ 1478 μs     | ~ 23244 μs
Deep Record Updates     | ~ 433 μs      | ~ 562 μs       | ~ 1113 μs | ~ 260 μs   | ~ 778 μs      | ~ 4943 μs
Ackermann               | ~ 211 μs      | ~ 210 μs       | ~ 20 μs   | ~ 28 μs    | ~ 57 μs       | ~ 418 μs
Church Numerals         | ~ 1662 μs     | ~ 1570 μs      | ~ 505 μs  | ~ 369 μs   | ~ 617 μs      | ~ 17684 μs
Prime Sieve             | ~ 725 μs      | ~ 689 μs       | ~ 274 μs  | ~ 76 μs    | ~ 232 μs      | ~ 7576 μs
Red-Black Tree          | ~ 94574 μs    | ~ 53648 μs     | ~ 40704 μs| ~ 25020 μs | ~ 17904 μs    | ~ 300566 μs
Polymorphism            | ~ 9029 μs     | ~ 8111 μs      | ~ 2314 μs | ~ 17884 μs | ~ 92155 μs    | ~ 10591 μs
State Monad             | ~ 425 μs      | ~ 170 μs       | ~ 43 μs   | ~ 5 μs     | ~ 108 μs      | ~ 539 μs
Lazy Evaluation         | ~ 16372 μs    | ~ 13986 μs     | ~ 21286 μs| ~ 2868 μs  | ~ 10325 μs    | ~ 95661 μs
Array Processing        | ~ 218 μs      | ~ 222 μs       | ~ 70 μs   | ~ 13 μs    | ~ 5430 μs     | ~ 1375 μs
----------------------- | ------------- | -------------- | --------- | ---------- | ------------- | ---------
Total Execution Time    | ~ 125.77 ms   | ~ 81.21 ms     | ~ 67.56 ms| ~ 46.87 ms | ~ 131.04 ms   | ~ 465.16 ms
```
> [!NOTE]
> **Single-Threaded Benchmark**
> All benchmarks presented here are strictly **single-threaded**. They measure raw sequential execution speed and do not take into account the powerful multi-threading capabilities inherent to languages like Go or Erlang (BEAM).

> [!WARNING]
> **About the PHP results**
> Please note that the `phpurs` backend is a brand new, completely experimental, and homemade compiler built entirely from scratch for this repository. Its execution time is not yet representative of PHP's actual performance limit. The currently displayed time is actively undergoing optimization and debugging!

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

```text
========================================================================================
EXTENDED BENCHMARK RESULTS (file I/O, regex, STArray, asynchronous Aff)                                                                      
========================================================================================
Benchmark               | JS            | Arista ES      | Go        
----------------------- | ------------- | -------------- | --------- 
AST Evaluation          | ~ 96 μs       | ~ 68 μs        | ~ 44 μs   
Fibonacci               | ~ 49 μs       | ~ 38 μs        | ~ 2 μs    
List Processing         | ~ 409 μs      | ~ 374 μs       | ~ 222 μs  
Tail Call Optimization  | ~ 1195 μs     | ~ 1572 μs      | ~ 1837 μs 
Deep Record Updates     | ~ 414 μs      | ~ 597 μs       | ~ 1780 μs 
Ackermann               | ~ 218 μs      | ~ 224 μs       | ~ 32 μs   
Church Numerals         | ~ 1812 μs     | ~ 1654 μs      | ~ 698 μs  
Prime Sieve             | ~ 693 μs      | ~ 656 μs       | ~ 434 μs  
Red-Black Tree          | ~ 99135 μs    | ~ 58721 μs     | ~ 49991 μs
Polymorphism            | ~ 8466 μs     | ~ 8930 μs      | ~ 2509 μs 
State Monad             | ~ 491 μs      | ~ 765 μs       | ~ 33 μs   
Lazy Evaluation         | ~ 14911 μs    | ~ 13361 μs     | ~ 22896 μs
Array Processing        | ~ 223 μs      | ~ 189 μs       | ~ 67 μs   
File I/O                | ~ 429223 μs   | ~ 479362 μs    | ~ 476440 μs
STArray Operations      | ~ 3 μs        | ~ 3 μs         | ~ 0 μs    
String Operations       | ~ 2 μs        | ~ 2 μs         | ~ 1 μs    
Aff Operations          | ~ 11482 μs    | ~ 11378 μs     | ~ 11030 μs
Parallelism             | ~ 15113637 μs | ~ 14690018 μs  | ~ 1255501 μs
----------------------- | ------------- | -------------- | --------- 
Total Execution Time    | ~ 15682.46 ms | ~ 15267.91 ms  | ~ 1823.52 ms
```

> [!NOTE]
> **Hardware Context**
> To accurately measure multi-core scaling, these extended benchmarks (specifically the 10 concurrent tasks in the *Parallelism* test) were executed on a machine equipped with **10 performance cores** (Apple M4 Pro).

## Repository structure and output files

> [!NOTE]
> This repository intentionally omits compiled folders from `.gitignore`. All generated code, dependencies, and compilation artifacts are deliberately committed.

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The compiled files generated by the backends for our module: `output/` (`.js` files for Node, `.erl`/`.beam` for Erlang, and Scheme libraries/executables).
- The state variables and raw benchmark results: `var/benchmark/`

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
