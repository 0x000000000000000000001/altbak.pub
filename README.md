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

#### JavaScript

JS Benchmark            | Compiled JS ([official](https://github.com/purescript/purescript)) | Compiled JS ([Arista](https://github.com/aristanetworks/purescript-backend-optimizer)) | Native FP-style JS FFI (WIP) | Native hand-optimized JS FFI (WIP) |
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

Go Benchmark            | Compiled Go ([gopurs](https://github.com/0x000000000000000000001/gopurs), mature WIP) | Compiled Go ([psgo](https://github.com/i-am-the-slime/purescript-native)) | Native FP-style Go FFI (WIP) | Native hand-optimized Go FFI (WIP) |
----------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
AST Evaluation          | ~ 0 μs                           | ~ 296 μs                              | ~ 20 μs       | ~ 4 μs                  |
Fibonacci               | ~ 0 μs                           | ~ 24 μs                               | ~ 2 μs        | ~ 1 μs                  |
List Processing         | ~ 9 μs                           | ~ 586 μs                              | ~ 121 μs      | ~ 1 μs                  |
Tail Call Optimization  | ~ 37 μs                          | ~ 13354 μs                            | ~ 39 μs       | ~ 27 μs                 |
Deep Record Updates     | ~ 5 μs                           | ~ 9492 μs                             | ~ 373 μs      | ~ 3 μs                  |
Ackermann               | ~ 16 μs                          | ~ 1267 μs                             | ~ 23 μs       | ~ 32 μs                 |
Church Numerals         | ~ 442 μs                         | ~ 5731 μs                             | ~ 1002 μs     | ~ 29 μs                 |
Prime Sieve             | ~ 71 μs                          | ~ 2549 μs                             | ~ 134 μs      | ~ 4 μs                  |
Red-Black Tree          | ~ 21506 μs                       | ~ 870034 μs                           | ~ 26529 μs    | ~ 24856 μs              |
Polymorphism            | ~ 2217 μs                        | ~ 669801 μs                           | ~ 58848 μs    | ~ 2480 μs               |
State Monad             | ~ 114 μs                         | ~ 536 μs                              | ~ 127 μs      | ~ 1 μs                  |
Lazy Evaluation         | ~ 229 μs                         | ~ 71586 μs                            | ~ 14816 μs    | ~ 1 μs                  |
Array Processing        | ~ 10 μs                          | ~ 61 μs                               | ~ 7 μs        | ~ 1 μs                  |
RowToList               | ~ 0 μs                           | ~ 2 μs                                | ~ 1 μs        | ~ 1 μs                  |
**Total Execution Time**    | ~ 24.66 ms                       | ~ 1645.32 ms                          | ~ 102.04 ms   | ~ 27.44 ms              |

#### Scheme

Scheme Benchmark        | Compiled Scheme ([pscm](https://github.com/purescm/purescm)) | Native FP-style Scheme FFI (WIP) | Native hand-optimized Scheme FFI (WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation          | ~ 9 μs | ~ 15 μs        | ~ 4 μs                   |
Fibonacci               | ~ 2 μs | ~ 3 μs         | ~ 2 μs                   |
List Processing         | ~ 10 μs | ~ 8 μs         | ~ 5 μs                   |
Tail Call Optimization  | ~ 326 μs | ~ 287 μs       | ~ 319 μs                 |
Deep Record Updates     | ~ 260 μs | ~ 45 μs        | ~ 40 μs                  |
Ackermann               | ~ 28 μs | ~ 13 μs        | ~ 11 μs                  |
Church Numerals         | ~ 369 μs | ~ 73 μs        | ~ 68 μs                  |
Prime Sieve             | ~ 76 μs | ~ 29 μs        | ~ 3 μs                   |
Red-Black Tree          | ~ 25020 μs | ~ 18061 μs      | ~ 12228 μs                |
Polymorphism            | ~ 17884 μs | ~ 6596 μs       | ~ 7737 μs                 |
State Monad             | ~ 5 μs | ~ 9 μs         | ~ 5 μs                   |
Lazy Evaluation         | ~ 2868 μs | ~ 1954 μs       | ~ 3 μs                   |
Array Processing        | ~ 13 μs | ~ 5 μs         | ~ 6 μs                   |
RowToList               | ~ 1 μs | ~ 1 μs         | ~ 1 μs                   |
**Total Execution Time**    | ~ 46.87 ms | ~ 27.1 ms       | ~ 20.43 ms                 |

#### Erlang

Erlang Benchmark        | Compiled Erlang ([purerl](https://github.com/purerl/purerl)) | Native FP-style Erlang FFI (WIP) | Native hand-optimized Erlang FFI (WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation          | ~ 692 μs | ~ 803 μs         | ~ 1030 μs                  |
Fibonacci               | ~ 49 μs | ~ 192 μs         | ~ 218 μs                   |
List Processing         | ~ 1212 μs | ~ 249 μs         | ~ 222 μs                   |
Tail Call Optimization  | ~ 1478 μs | ~ 318 μs         | ~ 369 μs                   |
Deep Record Updates     | ~ 778 μs | ~ 1601 μs       | ~ 261 μs                   |
Ackermann               | ~ 57 μs | ~ 280 μs         | ~ 343 μs                   |
Church Numerals         | ~ 617 μs | ~ 708 μs         | ~ 205 μs                   |
Prime Sieve             | ~ 232 μs | ~ 356 μs         | ~ 229 μs                   |
Red-Black Tree          | ~ 17904 μs | ~ 20106 μs      | ~ 43375 μs                |
Polymorphism            | ~ 92155 μs | ~ 202472 μs     | ~ 21901 μs                |
State Monad             | ~ 108 μs | ~ 335 μs         | ~ 342 μs                   |
Lazy Evaluation         | ~ 10325 μs | ~ 337 μs         | ~ 353 μs                   |
Array Processing        | ~ 5430 μs | ~ 187 μs         | ~ 316 μs                   |
**Total Execution Time**    | ~ 131.04 ms | ~ 228.18 ms      | ~ 69.4 ms                 |

#### PHP

PHP Benchmark           | Compiled PHP ([phpurs](https://github.com/0x000000000000000000001/phpurs), WIP) | Native FP-style PHP FFI (WIP) | Native hand-optimized PHP FFI (WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
AST Evaluation          | ~ 5.00 μs                 | ~ 9 μs         | ~ 10 μs                  |
Fibonacci               | ~ 2.00 μs                 | ~ 11 μs        | ~ 451 μs                 |
List Processing         | ~ 159.00 μs               | ~ 4 μs         | ~ 166 μs                 |
Tail Call Optimization  | ~ 77.00 μs                | ~ 1937 μs      | ~ 109 μs                 |
Deep Record Updates     | ~ 1449.00 μs              | ~ 4016 μs      | ~ 199 μs                 |
Ackermann               | ~ 46.00 μs                | ~ 665 μs       | ~ 579 μs                 |
Church Numerals         | ~ 2106.00 μs              | ~ 8922 μs      | ~ 105 μs                 |
Prime Sieve             | ~ 397.00 μs               | ~ 7 μs         | ~ 386 μs                 |
Red-Black Tree          | ~ 107980.00 μs            | ~ 730387 μs    | ~ 123096 μs              |
Polymorphism            | ~ 6166.00 μs              | ~ 1169540 μs   | ~ 74555 μs               |
State Monad             | ~ 430.00 μs               | ~ 358 μs       | ~ 86 μs                  |
Lazy Evaluation         | ~ 785.00 μs               | ~ 151601 μs    | ~ 368 μs                 |
Array Processing        | ~ 91.00 μs                | ~ 448 μs       | ~ 249 μs                 |
RowToList               | ~ 1.00 μs                 | ~ 4 μs         | ~ 94 μs                  |
**Total Execution Time**| ~ 119.69 ms               | ~ 2067.91 ms   | ~ 200.45 ms              |

#### Rust

Rust Benchmark          | Compiled Rust ([purust](https://github.com/0x000000000000000000001/purust), WIP) | Native FP-style Rust FFI (WIP) | Native hand-optimized Rust FFI (WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
AST Evaluation          | ~ 1.00 μs                 | ~ 38 μs      | ~ 1 μs                 |
Fibonacci               | ~ 1.00 μs                 | ~ 27 μs      | ~ 1 μs                 |
List Processing         | ~ 27.00 μs                | ~ 76 μs      | ~ 1 μs                 |
Tail Call Optimization  | ~ 34.00 μs                | ~ 114 μs     | ~ 35 μs                |
Deep Record Updates     | ~ 939.00 μs               | ~ 272 μs     | ~ 4 μs                 |
Ackermann               | ~ 19.00 μs                | ~ 38 μs      | ~ 18 μs                |
Church Numerals         | ~ 1592.00 μs              | ~ 598 μs     | ~ 1 μs                 |
Prime Sieve             | ~ 170.00 μs               | ~ 221 μs     | ~ 1 μs                 |
Red-Black Tree          | ~ 20099.00 μs             | ~ 59784 μs   | ~ 36070 μs             |
Polymorphism            | ~ 1.00 μs                 | ~ 8342 μs    | ~ 1 μs                 |
State Monad             | ~ 54.00 μs                | ~ 145 μs     | ~ 1 μs                 |
Lazy Evaluation         | ~ 19791.00 μs             | ~ 24919 μs   | ~ 0 μs                 |
Array Processing        | ~ 14.00 μs                | ~ 14 μs      | ~ 1 μs                 |
RowToList               | ~ 1.00 μs                 | ~ 27 μs      | ~ 0 μs                 |
**Total Execution Time**| ~ 42.74 ms                | ~ 94.61 ms      | ~ 36.13 ms                |

#### F#/C#

F#/C# Benchmark         | Compiled F#/C# ([sharpurs](https://github.com/0x000000000000000000001/sharpurs), WIP) | Native FP-style F#/C# FFI (WIP) | Native hand-optimized F#/C# FFI (WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation          | ~ 75.63 μs                | ~ 615 μs        | ~ 177 μs                  |
Fibonacci               | ~ 3.54 μs                 | ~ 252 μs        | ~ 63 μs                   |
List Processing         | ~ 689.13 μs               | ~ 472 μs        | ~ 77 μs                   |
Tail Call Optimization  | ~ 44.42 μs                | ~ 247 μs        | ~ 200 μs                  |
Deep Record Updates     | ~ 4599.38 μs              | ~ 314 μs        | ~ 225 μs                  |
Ackermann               | ~ 142.13 μs               | ~ 120 μs        | ~ 107 μs                  |
Church Numerals         | ~ 2053.79 μs              | ~ 80 μs         | ~ 76 μs                   |
Prime Sieve             | ~ 3520.04 μs              | ~ 447 μs        | ~ 119 μs                  |
Red-Black Tree          | ~ 39940.38 μs             | ~ 76561 μs      | ~ 65156 μs                |
Polymorphism            | ~ 2233.75 μs              | ~ 3343 μs       | ~ 2805 μs                 |
State Monad             | ~ 279.50 μs               | ~ 1638 μs       | ~ 100 μs                  |
Lazy Evaluation         | ~ 35376.75 μs             | ~ 107 μs        | ~ 75 μs                   |
Array Processing        | ~ 71.21 μs                | ~ 357 μs        | ~ 75 μs                   |
RowToList               | ~ 0.67 μs                 | ~ 47 μs         | ~ 45 μs                   |
**Total Execution Time**| ~ 89.03 ms                | ~ 84.6 ms       | ~ 69.3 ms                 |

#### Java

Java Benchmark          | Compiled Java ([javapurs](https://github.com/0x000000000000000000001/javapurs), WIP) | Native FP-style Java FFI (WIP) | Native hand-optimized Java FFI (WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
AST Evaluation          | ~ 54.00 μs                | ~ 69 μs                  | ~ 4.4 μs                       |
Fibonacci               | ~ 5.08 μs                 | ~ 2.5 μs                 | ~ 1.9 μs                       |
List Processing         | ~ 174.88 μs               | ~ 65 μs                  | ~ 6.4 μs                       |
Tail Call Optimization  | ~ 42.96 μs                | ~ 43 μs                  | ~ 42 μs                        |
Deep Record Updates     | ~ 157.75 μs               | ~ 132 μs                 | ~ 18 μs                        |
Ackermann               | ~ 19.50 μs                | ~ 10 μs                  | ~ 7.5 μs                       |
Church Numerals         | ~ 701.88 μs               | ~ 435 μs                 | ~ 1.8 μs                       |
Prime Sieve             | ~ 165.25 μs               | ~ 61 μs                  | ~ 8.9 μs                       |
Red-Black Tree          | ~ 18313.00 μs             | ~ 12278 μs               | ~ 12384 μs                     |
Polymorphism            | ~ 2.25 μs                 | ~ 148 μs                 | ~ 1.7 μs                       |
State Monad             | ~ 66.25 μs                | ~ 40 μs                  | ~ 6.5 μs                       |
Lazy Evaluation         | ~ 84.50 μs                | ~ 5168 μs                | ~ 1.5 μs                       |
Array Processing        | ~ 70.58 μs                | ~ 41 μs                  | ~ 7.1 μs                       |
RowToList               | ~ 1.67 μs                 | ~ 2.2 μs                 | ~ 1.3 μs                       |
**Total Execution Time**| ~ 19.86 ms                | ~ 18.50 ms               | ~ 12.49 ms                     |

> [!IMPORTANT]
> **The 99/1 philosophy and the AOT compiler vs FFI approach**
> 
> Using Go as an example, the three columns give a concrete idea of what the AOT compiler actually does:
> 
> 1. **Compiled Go (mature WIP)**: The actual code generated by our compiler (`gopurs`). For statically-typed AOT targets like Go, the compiled code is now **faster** than even the most optimized handwritten FFI. This is because a compiler has no constraints regarding code readability. It can systematically apply machine-level optimizations (such as deep monomorphization, aggressive inlining, or generating loops with an immoderate use of `goto` statements) that a human developer would never spontaneously write in order to keep their codebase maintainable.
> 2. **Native FP-style Go FFI**: This is what you get if a human translates PureScript's functional patterns (closures, type classes, boxed lists) directly into idiomatic, readable Go using native features like interfaces and type assertions.
> 3. **Native hand-optimized Go FFI**: A highly optimized, human-written implementation using raw imperative shortcuts. Unlike column 2, it doesn't try to faithfully replicate unoptimized functional patterns; it just runs as fast as possible. These hand-optimized implementations look wildly different for each test, making it a tough challenge for a compiler to predict them all.
>
> **The ultimate goal of the compiler** is to get as close as possible to the hand-optimized FFI (column 3). As the benchmarks now demonstrate for our most mature backend (`gopurs`), **we have actually surpassed this goal**, beating the fastest hand-written imperative code in overall performance. Reaching this milestone remains our active objective for the other experimental AOT targets (like PHP and Rust). We achieve this by detecting the structural shortcuts that a human brain naturally figures out when hand-optimizing code, and applying them ruthlessly. This relies on reproducible heuristics (unboxing, inlining, loop vectorization, TCO) that we actively carve into stone within the compiler engine, made possible by leveraging our custom TAST (Typed Abstract Syntax Tree) which preserves deep structural type information.
>
> **Why are these tests so naive?** These tests are deliberately naive to stress the runtime. For example, the Lazy Evaluation benchmark dynamically allocates and forces 1 million closures to heavily stress the garbage collector and call stack. They represent absolute worst-case scenarios. We want to maximize the performance gap between compiled and native code and use these artificially worsened gaps to drive continuous optimizations. When a hand-optimized script replaces a million closures with a raw `for` loop taking 1 µs, it's inherently unfair. But that's exactly the point: we want to see what happens when a developer makes a huge design mistake, and measure the performance ratio when several bad choices compound together.
>
> In practice, **your high-level codebase will execute faster than manually optimized native code**, letting you focus entirely on domain concepts instead of hardware details. Optimizing for catastrophic scenarios guarantees the best possible performance ratio for real-world projects, actively mitigating the impact of naive implementations. The historical need to manually optimize critical algorithmic *hot paths* by dropping down to FFI or using safe mutability abstractions (like the `ST` monad) is now virtually obsolete. You only need FFI for interacting with the outside world, not for raw computation speed. This philosophy applies universally to all backend languages benchmarked here: imperative code is kept to a strict, perfectly isolated minimum.

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
