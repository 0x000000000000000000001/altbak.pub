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
> **The benchmark is just an excuse**
> The primary objective here is **not** to pit these technologies against each other to declare a performance "winner". In fact, the benchmark itself is just a pretext. The real goal is to prove that we can seamlessly compile and run the exact same, unmodified PureScript code across **6 completely different runtimes** (V8 standard, V8 optimized, BEAM, Lisp, Go, and PHP). The true victory is achieving universal abstraction without sacrificing execution viability.

### Core vs extended tests (`srx/`)
To ensure fair and executable comparisons across all backends, the test suite is split into two parts:
1. **Core tests (`src/`)**: Pure computational tasks (AST, Fibonacci, recursion) that run seamlessly on all 6 backends. Executed via `./bin/run`.
2. **Extended tests (`srx/`)**: Tests relying heavily on Javascript/PHP FFI bindings (like `Effect.Aff`, mutable `STArray`, and regex). Since Scheme and Erlang lack FFI implementations for these specific libraries in their package sets, they are isolated in the `srx/` directory. **Note that this is completely normal and expected:** Scheme is targeted here for raw computation, and Erlang's BEAM already natively handles concurrency and multithreading at the VM level (making JS style `Aff` workarounds irrelevant). Executed via `./bin/run --x` (which dynamically injects `srx/` into the compilation step and skips Scheme/Erlang).

### Core stresstest benchmark results (pure computational)

Command: `./bin/run` (Runs on all 7 backends). New tests will gradually be added.

```text
=============================================================================================================================================================
CORE STRESSTEST BENCHMARK RESULTS (Fibonacci, AST, tail calls, Church, primes, etc.)                                                                      
=============================================================================================================================================================
Benchmark               | Compiled JS   | Compiled Arista JS | Native JS FFI
----------------------- | ------------- | ------------------ | -------------
AST Evaluation          | ~ 93 μs       | ~ 74 μs            | ~ 119 μs
Fibonacci               | ~ 43 μs       | ~ 46 μs            | ~ 68 μs
List Processing         | ~ 386 μs      | ~ 368 μs           | ~ 54 μs
Tail Call Optimization  | ~ 1597 μs     | ~ 1550 μs          | ~ 635 μs
Deep Record Updates     | ~ 433 μs      | ~ 562 μs           | ~ 606 μs
Ackermann               | ~ 211 μs      | ~ 210 μs           | ~ 152 μs
Church Numerals         | ~ 1662 μs     | ~ 1570 μs          | ~ 544 μs
Prime Sieve             | ~ 725 μs      | ~ 689 μs           | ~ 188 μs
Red-Black Tree          | ~ 94574 μs    | ~ 53648 μs         | ~ 33658 μs
Polymorphism            | ~ 9029 μs     | ~ 8111 μs          | ~ 4246 μs
State Monad             | ~ 425 μs      | ~ 170 μs           | ~ 36 μs
Lazy Evaluation         | ~ 16372 μs    | ~ 13986 μs         | ~ 357 μs
Array Processing        | ~ 218 μs      | ~ 222 μs           | ~ 33 μs
----------------------- | ------------- | ------------------ | -------------
Total Execution Time    | ~ 125.77 ms   | ~ 81.21 ms         | ~ 40.72 ms

Benchmark               | Compiled Go (mature WIP) | Native Go FFI
----------------------- | ------------------------ | -------------
AST Evaluation          | ~ 10 μs                  | ~ 6 μs
Fibonacci               | ~ 2 μs                   | ~ 2 μs
List Processing         | ~ 133 μs                 | ~ 2 μs
Tail Call Optimization  | ~ 897 μs                 | ~ 39 μs
Deep Record Updates     | ~ 1034 μs                | ~ 5 μs
Ackermann               | ~ 58 μs                  | ~ 67 μs
Church Numerals         | ~ 947 μs                 | ~ 79 μs
Prime Sieve             | ~ 225 μs                 | ~ 5 μs
Red-Black Tree          | ~ 37305 μs               | ~ 27079 μs
Polymorphism            | ~ 2326 μs                | ~ 4541 μs
State Monad             | ~ 181 μs                 | ~ 1 μs
Lazy Evaluation         | ~ 21525 μs               | ~ 273 μs
Array Processing        | ~ 38 μs                  | ~ 2 μs
----------------------- | ------------------------ | -------------
Total Execution Time    | ~ 64.68 ms               | ~ 32.10 ms

Benchmark               | Compiled Scheme | Native Scheme FFI
----------------------- | ----------------| -----------------
AST Evaluation          | ~ 9 μs          | 
Fibonacci               | ~ 2 μs          | 
List Processing         | ~ 10 μs         | 
Tail Call Optimization  | ~ 326 μs        | 
Deep Record Updates     | ~ 260 μs        | 
Ackermann               | ~ 28 μs         | 
Church Numerals         | ~ 369 μs        | 
Prime Sieve             | ~ 76 μs         | 
Red-Black Tree          | ~ 25020 μs      | 
Polymorphism            | ~ 17884 μs      | 
State Monad             | ~ 5 μs          | 
Lazy Evaluation         | ~ 2868 μs       | 
Array Processing        | ~ 13 μs         | 
----------------------- | ----------------| -----------------
Total Execution Time    | ~ 46.87 ms      | 

Benchmark               | Compiled Erlang | Native Erlang FFI
----------------------- | ----------------| -----------------
AST Evaluation          | ~ 692 μs        | 
Fibonacci               | ~ 49 μs         | 
List Processing         | ~ 1212 μs       | 
Tail Call Optimization  | ~ 1478 μs       | 
Deep Record Updates     | ~ 778 μs        | 
Ackermann               | ~ 57 μs         | 
Church Numerals         | ~ 617 μs        | 
Prime Sieve             | ~ 232 μs        | 
Red-Black Tree          | ~ 17904 μs      | 
Polymorphism            | ~ 92155 μs      | 
State Monad             | ~ 108 μs        | 
Lazy Evaluation         | ~ 10325 μs      | 
Array Processing        | ~ 5430 μs       | 
----------------------- | ----------------| -----------------
Total Execution Time    | ~ 131.04 ms     | 

Benchmark               | Compiled PHP (normal WIP) | Native PHP FFI
----------------------- | ------------------------ | --------------
AST Evaluation          | ~ 18 μs                  | 
Fibonacci               | ~ 354 μs                 | 
List Processing         | ~ 2189 μs                | 
Tail Call Optimization  | ~ 23244 μs               | 
Deep Record Updates     | ~ 4943 μs                | 
Ackermann               | ~ 418 μs                 | 
Church Numerals         | ~ 17684 μs               | 
Prime Sieve             | ~ 7576 μs                | 
Red-Black Tree          | ~ 300566 μs              | 
Polymorphism            | ~ 10591 μs               | 
State Monad             | ~ 539 μs                 | 
Lazy Evaluation         | ~ 95661 μs               | 
Array Processing        | ~ 1375 μs                | 
----------------------- | ------------------------ | --------------
Total Execution Time    | ~ 465.16 ms              | 

Benchmark               | Compiled Rust (young WIP)| Native Rust FFI
----------------------- | ------------------------ | ---------------
AST Evaluation          | ~ 15 μs                  | 
Fibonacci               | ~ 4 μs                   | 
List Processing         | ~ 172 μs                 | 
Tail Call Optimization  | ~ 746 μs                 | 
Deep Record Updates     | ~ 2295 μs                | 
Ackermann               | ~ 46 μs                  | 
Church Numerals         | ~ 46762 μs               | 
Prime Sieve             | ~ 311 μs                 | 
Red-Black Tree          | ~ 60546 μs               | 
Polymorphism            | ~ 38606 μs               | 
State Monad             | ~ 406 μs                 | 
Lazy Evaluation         | ~ 310128 μs              | 
Array Processing        | ~ 42 μs                  | 
RowToList               | ~ 2 μs                   | 
----------------------- | ------------------------ | ---------------
Total Execution Time    | ~ 460.08 ms              | 
```

> [!IMPORTANT]
> **The 99/1 philosophy and the FFI approach**
> Using Go as an example: the *Native Go FFI* column demonstrates that the code generated by the compiler remains within a highly competitive order of magnitude, especially given the high-level expressiveness of PureScript (i.e., how much low-level machinery we can safely ignore). Since these tests are deliberately naive and designed to stress the runtime, they represent absolute worst-case scenarios for the compiler. The vast majority of real-world code will perform significantly better, yielding a far more favorable ratio between compiled and native execution times.
>
> In practice, **99% of your codebase will be nearly as fast as native code**, allowing you to focus entirely on domain concepts rather than hardware details. For the remaining 1% (critical algorithmic *hot paths*), you have three options: accept a minimal Nx slowdown that the compiler actively works to mitigate (n.b. N ~ 2 or 3), use safe mutability abstractions like the `ST` monad (which compile down to highly efficient imperative loops), or seamlessly drop down into FFI to write native code as close to the metal as needed. This philosophy applies universally to all backend languages benchmarked here: imperative code is kept to a strict, perfectly isolated minimum. At this stage, optimization gains are often on the scale of mere microseconds; while further improvement is always possible, chasing these micro-optimizations is no longer a priority.

> [!NOTE]
> **Single-threaded benchmark**
> All benchmarks presented here are strictly **single-threaded**. They measure raw sequential execution speed and do not take into account the powerful multi-threading capabilities inherent to languages like Go or Erlang (BEAM).

### Extended benchmark results (I/O, mutability, async)
Command: `./bin/run --x` (Skips runtimes lacking necessary FFI bindings like Scheme and Erlang)

```text
========================================================================================
EXTENDED BENCHMARK RESULTS (file I/O, regex, STArray, asynchronous Aff)                                                                      
========================================================================================
Benchmark               | JS            | Arista ES      | Go              
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
----------------------- | ------------- | -------------- | --------------- 
Total Execution Time    | ~ 15682.46 ms | ~ 15267.91 ms  | ~ 1823.52 ms    
```

> [!NOTE]
> **Hardware Context**
> To accurately measure multi-core scaling, these extended benchmarks (specifically the 10 concurrent tasks in the *Parallelism* test) were executed on a machine equipped with **10 performance cores** (Apple M4 Pro).

## Repository structure and output files

The purpose of this approach is to allow an educational exploration of how the backends work, without needing to install the local compilers yourself. You can directly inspect:

- The compiled files generated by the backends for our module: `output/` (`.js` files for Node, `.erl`/`.beam` for Erlang, and Scheme libraries/executables).
- The state variables and raw benchmark results: `var/benchmark/`

The main orchestration script is `bin/run`. It calls the backend specific runners which manage compilation and execute the timed results.
