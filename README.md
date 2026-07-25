# PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. This is made possible by the PureScript compiler architecture, which generates an intermediate representation (`CoreFn`) that can be consumed by various backends:

1. **JavaScript (V8)**: The premier asynchronous JavaScript engine ([default backend](https://github.com/purescript/purescript)).
2. **Arista ES (V8)**: A highly optimized, modern ECMAScript backend developed by Arista Networks for extreme performance (via the [`purs-backend-es` backend](https://github.com/aristanetworks/purescript-backend-optimizer)).
3. **Erlang (BEAM)**: The distributed, highly concurrent, and fault tolerant virtual machine (via the [`purerl` backend](https://github.com/purerl/purerl)).
4. **Chez Scheme**: One of the fastest Lisp compilers in the world for highly optimized native execution (via the [`purescm` backend](https://github.com/purescm/purescm)).
5. **PHP**: Generating modern PHP 7.4+ syntax (via the experimental local `phpurs` backend). PureScript that transpiles to PHP, and targets 70% of the web (e.g. containerless VPS).

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
Command: `./bin/run` (Runs on all 6 backends)

```text
====================================================================================================================================
CORE BENCHMARK RESULTS (Fibonacci, AST, tail calls, Church, primes, etc.)                                                                      
====================================================================================================================================
JS (V8)     | Arista ES (V8) | Chez Scheme (Native) | Erlang (BEAM) | Native Go (experimental WIP) | PHP (experimental WIP)
----------- | -------------- | -------------------- | ------------- | ---------------------------- | ----------------------
~ 122.65 ms | ~ 77.66 ms     | ~ 46.41 ms           | ~ 122.71 ms   | ~ 125.47 ms                  | ~ 11716.88 ms (WIP! A regression bug will be fixed, usually ~900ms)
```

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

```text
========================================================================================
EXTENDED BENCHMARK RESULTS (file I/O, regex, STArray, asynchronous Aff)                                                                      
========================================================================================
WIP, usually +20% everywhere
```

> [!WARNING]
> **About the PHP results**
> Please note that the `phpurs` backend is a brand new, completely experimental, and homemade compiler built entirely from scratch for this repository. Its execution time is not yet representative of PHP's actual performance limit. The currently displayed time (~11s) is due to a recent regression bug; prior to this, it was running at around **900 ms** (which itself was still a WIP milestone, not the final optimized limit). It is actively undergoing optimization and debugging!


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
