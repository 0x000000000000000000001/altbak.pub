# PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. This is made possible by the PureScript compiler architecture, which generates an intermediate representation (`CoreFn`) that can be consumed by various backends:

1. **JavaScript (V8)**: The premier asynchronous JavaScript engine ([default backend](https://github.com/purescript/purescript)).
2. **Arista ES (V8)**: A highly optimized, modern ECMAScript backend developed by Arista Networks for extreme performance (via the [`purs-backend-es` backend](https://github.com/aristanetworks/purescript-backend-optimizer)).
3. **Erlang (BEAM)**: The distributed, highly concurrent, and fault tolerant virtual machine (via the [`purerl` backend](https://github.com/purerl/purerl)).
4. **Chez Scheme**: One of the fastest Lisp compilers in the world for highly optimized native execution (via the [`purescm` backend](https://github.com/purescm/purescm)).
5. **Go**: An experimental Ahead-Of-Time (AOT) backend generating native Go binaries (via the experimental local `gopurs` backend).
6. **PHP**: Generating modern PHP 7.4+ syntax (via the experimental local `phpurs` backend). PureScript that transpiles to PHP, and targets 70% of the web (e.g. containerless VPS).
7. **Rust**: An experimental Ahead-Of-Time (AOT) backend generating native Rust binaries (via the experimental local `purust` backend).

## Comprehensive benchmarks
The benchmark suite runs a wide variety of computationally intensive tasks: AST evaluation, purely recursive Fibonacci, massive list processing, tail call optimization, deep record updates, Ackermann function, Church numerals, prime sieves, red black tree insertions, heavy polymorphism (type class dictionary lookups), State monad operations, deep lazy evaluation, heavy file I/O (10,000 synchronous writes and reads), and asynchronous `Aff` operations (via the native event loop). These tests apply massive pressure on the call stack, garbage collector, disk I/O, event loop, and runtime execution engine to measure the raw ability of the compiler and the underlying virtual machine. 

> [!IMPORTANT]
> **Worst-case scenario driven design**
> The code in these benchmarks is deliberately designed to be as naive, inefficient, and stressful as possible. The goal is to obtain the maximum possible performance gap between the compiled code and its native FFI equivalent. By studying these artificially worsened gaps, we can continually improve the code generation of our AOT compilers.
> 
> In the context of a real-world project, the idea is to provide the best possible performance ratio even when a developer makes a catastrophic design error. This mitigates the impact of such errors and delays as long as possible the need to drop down to FFI or use mutable abstractions for the rare algorithmic *hot paths* of a project.

### Core vs extended tests (`srx/`)
To ensure fair and executable comparisons across all backends, the test suite is split into two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) that run seamlessly on all 6 backends. Executed via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Core stresstest benchmark results (pure computational)

Command: `./bin/run` (Runs on all 7 backends). New tests will gradually be added.

#### JavaScript

JS Benchmark            | Compiled JS ([official](https://github.com/purescript/purescript)) | Compiled JS ([Arista](https://github.com/aristanetworks/purescript-backend-optimizer)) | Native JS FFI | Native JS FFI Cheatcode |
----------------------- | ---------------------- | ------------------ | ------------- | ----------------------- |
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
RowToList               | -                      | -                  | ~ 58 μs       | ~ 19 μs                 |
**Total Execution Time**    | ~ 125.77 ms            | ~ 81.21 ms         | ~ 73.45 ms    | ~ 34.19 ms              |
> *Read the IMPORTANT notice below!*

#### Go

Go Benchmark            | Compiled Go ([gopurs](https://github.com/0x000000000000000000001/gopurs), mature WIP) | Compiled Go ([psgo](https://github.com/i-am-the-slime/purescript-native)) | Native Go FFI | Native Go FFI Cheatcode |
----------------------- | -------------------------------- | ------------------------------------- | ------------- | ----------------------- |
AST Evaluation          | ~ 3 μs                           | ~ 296 μs                              | ~ 20 μs       | ~ 4 μs                  |
Fibonacci               | ~ 1 μs                           | ~ 24 μs                               | ~ 2 μs        | ~ 1 μs                  |
List Processing         | ~ 43 μs                          | ~ 586 μs                              | ~ 121 μs      | ~ 1 μs                  |
Tail Call Optimization  | ~ 720 μs                         | ~ 13354 μs                            | ~ 39 μs       | ~ 27 μs                 |
Deep Record Updates     | ~ 259 μs                         | ~ 9492 μs                             | ~ 373 μs      | ~ 3 μs                  |
Ackermann               | ~ 22 μs                          | ~ 1267 μs                             | ~ 23 μs       | ~ 32 μs                 |
Church Numerals         | ~ 2277 μs                        | ~ 5731 μs                             | ~ 1002 μs     | ~ 29 μs                 |
Prime Sieve             | ~ 163 μs                         | ~ 2549 μs                             | ~ 134 μs      | ~ 4 μs                  |
Red-Black Tree          | ~ 25320 μs                       | ~ 870034 μs                           | ~ 26529 μs    | ~ 24856 μs              |
Polymorphism            | ~ 2252 μs                        | ~ 669801 μs                           | ~ 58848 μs    | ~ 2480 μs               |
State Monad             | ~ 234 μs                         | ~ 536 μs                              | ~ 127 μs      | ~ 1 μs                  |
Lazy Evaluation         | ~ 17599 μs                       | ~ 71586 μs                            | ~ 14816 μs    | ~ 1 μs                  |
Array Processing        | ~ 34 μs                          | ~ 61 μs                               | ~ 7 μs        | ~ 1 μs                  |
RowToList               | ~ 1 μs                           | ~ 2 μs                                | ~ 1 μs        | ~ 1 μs                  |
**Total Execution Time**    | ~ 48.93 ms                       | ~ 1645.32 ms                          | ~ 102.04 ms   | ~ 27.44 ms              |
> *Read the IMPORTANT notice below!*

#### Scheme

Scheme Benchmark        | Compiled Scheme ([pscm](https://github.com/purescm/purescm)) | Native Scheme FFI | Native Scheme FFI Cheatcode |
----------------------- | ----------------| ----------------- | --------------------------- |
AST Evaluation          | ~ 9 μs | ~ 15 μs        | ~ 4 μs                   |
Fibonacci               | ~ 2 μs | ~ 3 μs         | ~ 2 μs                   |
List Processing         | ~ 10 μs | ~ 8 μs         | ~ 5 μs                   |
Tail Call Optimization  | ~ 326 μs | ~ 287 μs       | ~ 319 μs                 |
Deep Record Updates     | ~ 260 μs | ~ 45 μs        | ~ 40 μs                  |
Ackermann               | ~ 28 μs | ~ 13 μs        | ~ 11 μs                  |
Church Numerals         | ~ 369 μs | ~ 73 μs        | ~ 68 μs                  |
Prime Sieve             | ~ 76 μs | ~ 29 μs        | ~ 3 μs                   |
Red-Black Tree          | ~ 25020 μs | ~ 18.061 ms      | ~ 12.228 ms                |
Polymorphism            | ~ 17884 μs | ~ 6.596 ms       | ~ 7.737 ms                 |
State Monad             | ~ 5 μs | ~ 9 μs         | ~ 5 μs                   |
Lazy Evaluation         | ~ 2868 μs | ~ 1.954 ms       | ~ 3 μs                   |
Array Processing        | ~ 13 μs | ~ 5 μs         | ~ 6 μs                   |
**Total Execution Time**    | ~ 46.87 ms | ~ 27.10 ms       | ~ 20.43 ms                 |
> *Read the IMPORTANT notice below!*

#### Erlang

Erlang Benchmark        | Compiled Erlang ([purerl](https://github.com/purerl/purerl)) | Native Erlang FFI | Native Erlang FFI Cheatcode |
----------------------- | ----------------| ----------------- | --------------------------- |
AST Evaluation          | ~ 692 μs | ~ 803 μs         | ~ 1.03 ms                  |
Fibonacci               | ~ 49 μs | ~ 192 μs         | ~ 218 μs                   |
List Processing         | ~ 1212 μs | ~ 249 μs         | ~ 222 μs                   |
Tail Call Optimization  | ~ 1478 μs | ~ 318 μs         | ~ 369 μs                   |
Deep Record Updates     | ~ 778 μs | ~ 1.601 ms       | ~ 261 μs                   |
Ackermann               | ~ 57 μs | ~ 280 μs         | ~ 343 μs                   |
Church Numerals         | ~ 617 μs | ~ 708 μs         | ~ 205 μs                   |
Prime Sieve             | ~ 232 μs | ~ 356 μs         | ~ 229 μs                   |
Red-Black Tree          | ~ 17904 μs | ~ 20.106 ms      | ~ 43.375 ms                |
Polymorphism            | ~ 92155 μs | ~ 202.472 ms     | ~ 21.901 ms                |
State Monad             | ~ 108 μs | ~ 335 μs         | ~ 342 μs                   |
Lazy Evaluation         | ~ 10325 μs | ~ 337 μs         | ~ 353 μs                   |
Array Processing        | ~ 5430 μs | ~ 187 μs         | ~ 316 μs                   |
**Total Execution Time**    | ~ 131.04 ms | ~ 228.18 ms      | ~ 69.40 ms                 |
> *Read the IMPORTANT notice below!*

#### PHP

PHP Benchmark           | Compiled PHP ([phpurs](https://github.com/0x000000000000000000001/phpurs), WIP) | Native PHP FFI | Native PHP FFI Cheatcode |
----------------------- | ------------------------- | -------------- | ------------------------ |
AST Evaluation          | ~ 18 μs                   | ~ 9 μs         | ~ 10 μs                  |
Fibonacci               | ~ 354 μs                  | ~ 11 μs        | ~ 451 μs                 |
List Processing         | ~ 2189 μs                 | ~ 4 μs         | ~ 166 μs                 |
Tail Call Optimization  | ~ 23244 μs                | ~ 1937 μs      | ~ 109 μs                 |
Deep Record Updates     | ~ 4943 μs                 | ~ 4016 μs      | ~ 199 μs                 |
Ackermann               | ~ 418 μs                  | ~ 665 μs       | ~ 579 μs                 |
Church Numerals         | ~ 17684 μs                | ~ 8922 μs      | ~ 105 μs                 |
Prime Sieve             | ~ 7576 μs                 | ~ 7 μs         | ~ 386 μs                 |
Red-Black Tree          | ~ 300566 μs               | ~ 730387 μs    | ~ 123096 μs              |
Polymorphism            | ~ 10591 μs                | ~ 1169540 μs   | ~ 74555 μs               |
State Monad             | ~ 539 μs                  | ~ 358 μs       | ~ 86 μs                  |
Lazy Evaluation         | ~ 95661 μs                | ~ 151601 μs    | ~ 368 μs                 |
Array Processing        | ~ 1375 μs                 | ~ 448 μs       | ~ 249 μs                 |
RowToList               | -                         | ~ 4 μs         | ~ 94 μs                  |
**Total Execution Time**    | ~ 465.16 ms               | ~ 2067.91 ms   | ~ 200.45 ms              |
> *Read the IMPORTANT notice below!*

#### Rust

Rust Benchmark          | Compiled Rust ([purust](https://github.com/0x000000000000000000001/purust), young WIP) | Native Rust FFI | Native Rust FFI Cheatcode |
----------------------- | ------------------------- | --------------- | ------------------------- |
AST Evaluation          | ~ 48 μs                   | ~ 38.00 μs      | ~ 4.00 μs                 |
Fibonacci               | ~ 5 μs                    | ~ 27.00 μs      | ~ 4.00 μs                 |
List Processing         | ~ 146 μs                  | ~ 76.00 μs      | ~ 4.00 μs                 |
Tail Call Optimization  | ~ 762 μs                  | ~ 114.00 μs     | ~ 47.00 μs                |
Deep Record Updates     | ~ 2166 μs                 | ~ 272.00 μs     | ~ 7.00 μs                 |
Ackermann               | ~ 26 μs                   | ~ 38.00 μs      | ~ 29.00 μs                |
Church Numerals         | ~ 16363 μs                | ~ 598.00 μs     | ~ 3.00 μs                 |
Prime Sieve             | ~ 265 μs                  | ~ 221.00 μs     | ~ 4.00 μs                 |
Red-Black Tree          | ~ 65765 μs                | ~ 59784.00 μs   | ~ 16700.00 μs             |
Polymorphism            | ~ 38713 μs                | ~ 8342.00 μs    | ~ 4.00 μs                 |
State Monad             | ~ 419 μs                  | ~ 145.00 μs     | ~ 3.00 μs                 |
Lazy Evaluation         | ~ 324862 μs               | ~ 24919.00 μs   | ~ 3.00 μs                 |
Array Processing        | ~ 35 μs                   | ~ 14.00 μs      | ~ 3.00 μs                 |
RowToList               | ~ 2 μs                    | ~ 27.00 μs      | ~ 2.00 μs                 |
**Total Execution Time**    | ~ 449.58 ms               | ~ 94.61 ms      | ~ 16.82 ms                |
> *Read the IMPORTANT notice below!*

> [!IMPORTANT]
> **The 99/1 philosophy and the AOT compiler vs FFI vs cheatcode approach**
> 
> Using Go as an example, the three columns give a concrete idea of what the AOT compiler actually does:
> 
> 1. **Compiled Go (mature WIP)**: The actual code generated by our compiler (`gopurs`). For statically-typed AOT targets like Go, the compiled code is often **faster** than basic handwritten FFI. 
> 2. **Native Go FFI**: This is what you get if a human translates PureScript's functional patterns (closures, type classes, boxed lists) directly into idiomatic Go using native features like interfaces and type assertions.
> 3. **Native Go FFI cheatcode**: The theoretical limit of the hardware, using raw imperative shortcuts. Unlike column 2, it doesn't try to faithfully replicate unoptimized functional patterns; it just runs as fast as possible. Cheatcodes look wildly different for each test, making it a tough challenge for a compiler to predict them all. Some patterns are predictable though, and those can be integrated into the compiler engine. 
>
> **The ultimate goal of the compiler** isn't just to match or outperform the faithful native FFI (column 2). Ultimately, the goal is to see how much unoptimized design we can cancel out to match the performance of column 3. We want to get as close as possible to the cheatcode, proving that the generated code stays within a highly competitive order of magnitude despite the high-level expressiveness of PureScript or bad upfront design. We achieve this by detecting the structural shortcuts that a human brain naturally figures out when writing a cheatcode. Generating a perfect cheatcode for every arbitrary design is obviously impossible, but a large part of it relies on reproducible heuristics (unboxing, inlining, loop vectorization, TCO) that we actively carve into stone within the compiler engine, made possible by leveraging our custom TAST (Typed Abstract Syntax Tree) which preserves deep structural type information.
>
> **Why are these tests so naive?** These tests are deliberately naive to stress the runtime. For example, the Lazy Evaluation benchmark dynamically allocates and forces 1 million closures to heavily stress the garbage collector and call stack. They represent absolute worst-case scenarios. We want to maximize the performance gap between compiled and native code and use these artificially worsened gaps to drive continuous optimizations. When the cheatcode replaces a million closures with a raw `for` loop taking 1 µs, it's inherently unfair. But that's exactly the point: we want to see what happens when a developer makes a huge design mistake, and measure the performance ratio when several bad choices compound together.
>
> In practice, **99% of your codebase will be as fast as optimized native code**, letting you focus entirely on domain concepts instead of hardware details. Optimizing for catastrophic scenarios guarantees the best possible performance ratio for real-world projects. This actively mitigates the impact of naive implementations and delays the need to manually optimize the remaining 1% (critical algorithmic *hot paths*). For those rare hot paths, you still have three options: accept a 2x slowdown that the compiler actively tries to minimize, use safe mutability abstractions like the `ST` monad (which compile down to fast imperative loops), or drop down into FFI to write native code as close to the metal as needed. This philosophy applies universally to all backend languages benchmarked here: imperative code is kept to a strict, perfectly isolated minimum.

> [!NOTE]
> **Single-threaded benchmark**
> All benchmarks presented here are strictly **single-threaded**. They measure raw sequential execution speed and do not take into account the powerful multi-threading capabilities inherent to languages like Go or Erlang (BEAM).

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

#### Extended Results

Benchmark               | Compiled JS ([official](https://github.com/purescript/purescript)) | Compiled JS ([Arista](https://github.com/aristanetworks/purescript-backend-optimizer)) | Compiled Go ([gopurs](https://github.com/0x000000000000000000001/gopurs), mature WIP)
----------------------- | ------------- | -------------- | --------------- 
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
**Total Execution Time**    | ~ 15682.46 ms | ~ 15267.91 ms  | ~ 1823.52 ms    

> [!NOTE]
> **Hardware Context**
> To accurately measure multi-core scaling, these extended benchmarks (specifically the 10 concurrent tasks in the *Parallelism* test) were executed on a machine equipped with **10 performance cores** (Apple M4 Pro).

## Repository structure and output files

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The compiled files generated by the backends for our module: `output/` (`.js` files for Node, `.erl`/`.beam` for Erlang, and Scheme libraries/executables).
- The state variables and raw benchmark results: `var/benchmark/`

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
