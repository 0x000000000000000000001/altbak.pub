# 🌈 PureScript universal multi runtime benchmark

## Project goal
This project is a proof of concept demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. Backends consume either standard `CoreFn` or, for the local AOT backends, the typed `TAST/tcorefn` representation exported by a local PureScript fork (may be subject to an official PR soon).

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

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT<br><br>(WIP) |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
AST Evaluation | ~ 0.259782 μs | ~ 0.188290 μs | ~ 0.195508 μs | ~ 0.218845 μs |
Fibonacci | ~ 0.309606 μs | ~ 0.310805 μs | ~ 0.314829 μs | ~ 0.371354 μs |
List Processing | ~ 9.588542 μs | ~ 5.057760 μs | ~ 19.055257 μs | ~ 0.493614 μs |
Tail Call Optimization | ~ 189.485016 μs | ~ 151.132484 μs | ~ 65.913250 μs | ~ 68.307941 μs |
Deep Record Updates | ~ 74.181641 μs | ~ 134.113281 μs | ~ 42.068848 μs | ~ 7.685384 μs |
Ackermann | ~ 67.494629 μs | ~ 64.403320 μs | ~ 38.249512 μs | ~ 38.483238 μs |
Church Numerals | ~ 1263.817625 μs | ~ 1242.291687 μs | ~ 1337.510375 μs | ~ 49.629965 μs |
Prime Sieve | ~ 63.388023 μs | ~ 26.988934 μs | ~ 86.325195 μs | ~ 0.651260 μs |
Red-Black Tree | ~ 89940.584000 μs | ~ 47272.875000 μs | ~ 24486.833000 μs | ~ 27249.584000 μs |
Polymorphism | ~ 22830.250000 μs | ~ 7368.375000 μs | ~ 4843.000000 μs | ~ 2805.250000 μs |
State Monad | ~ 48.230633 μs | ~ 10.946736 μs | ~ 46.870770 μs | ~ 0.417788 μs |
Lazy Evaluation | ~ 14608.292000 μs | ~ 10604.875000 μs | ~ 17513.125000 μs | ~ 321.477875 μs |
Array Processing | ~ 7.350688 μs | ~ 3.751200 μs | ~ 3.121460 μs | ~ 0.498405 μs |
RowToList | ~ 0.041569 μs | ~ 0.017147 μs | ~ 0.059979 μs | ~ 0.020460 μs |
[Array Indexing (multiple sizes)](docs/benchmark-results/2026-09-21-go-diagnostics.md#array-indexing) | ~ 25.303333 ms | — | — | — |
[JSON → Typed AST (parse + decode)](docs/benchmark-results/2026-09-21-go-diagnostics.md#json-to-typed-ast) | ~ 83.838958 ms | — | — | — |
**Total Execution Time** | ~ 129.103274 ms | ~ 66.885327 ms | ~ 48.482643 ms | ~ 30.543090 ms |

These two additional diagnostics use the standard PureScript JS emitter without
PBO JS optimization and are excluded from the historical 14-case total. Their
[inputs, timing boundaries and Go comparison](docs/benchmark-results/2026-09-21-go-diagnostics.md)
are documented separately; “—” means no equivalent measurement was made.

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(WIP) |
----------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
AST Evaluation | ~ 0.328265 μs | ~ 4.174571 μs | ~ 0.323737 μs | ~ 0.318045 μs |
Fibonacci | ~ 0.246470 μs | ~ 5.406087 μs | ~ 0.156069 μs | ~ 0.164510 μs |
List Processing | ~ 17.188558 μs | ~ 236.006500 μs | ~ 24.387207 μs | ~ 0.449310 μs |
Tail Call Optimization | ~ 57.182777 μs | ~ 6281.958500 μs | ~ 35.984049 μs | ~ 35.495768 μs |
Deep Record Updates | ~ 6.633260 μs | ~ 5461.520500 μs | ~ 40.430336 μs | ~ 14.867513 μs |
Ackermann | ~ 15.926188 μs | ~ 402.039062 μs | ~ 20.846680 μs | ~ 19.867514 μs |
Church Numerals | ~ 230.560547 μs | ~ 3811.479000 μs | ~ 1044.604250 μs | ~ 22.448080 μs |
Prime Sieve | ~ 79.353188 μs | ~ 1340.994750 μs | ~ 110.873695 μs | ~ 0.623807 μs |
Red-Black Tree | ~ 10161.250000 μs | ~ 870404.625000 μs | ~ 24630.666000 μs | ~ 9004.479500 μs |
Polymorphism | ~ 2292.364625 μs | ~ 505069.000000 μs | ~ 49160.916000 μs | ~ 2244.453125 μs |
State Monad | ~ 150.942383 μs | ~ 303.558594 μs | ~ 52.104168 μs | ~ 0.333684 μs |
Lazy Evaluation | ~ 247.106125 μs | ~ 71378.083000 μs | ~ 13865.375000 μs | ~ 0.268639 μs |
Array Processing | ~ 19.262085 μs | ~ 43.811848 μs | ~ 4.903117 μs | ~ 0.451534 μs |
RowToList | ~ 0.151810 μs | ~ 0.504074 μs | ~ 0.072531 μs | ~ 0.031958 μs |
[Array Indexing (multiple sizes)](docs/benchmark-results/2026-09-21-go-diagnostics.md#array-indexing) | ~ 3.962167 ms | — | — | — |
[JSON → Typed AST (parse + decode)](docs/benchmark-results/2026-09-21-go-diagnostics.md#json-to-typed-ast) | ~ 717.159500 ms | — | — | — |
**Total Execution Time** | ~ 13.278496 ms | ~ 1464.743161 ms | ~ 88.991643 ms | ~ 11.344253 ms |

The two additional diagnostics were measured on September 21, 2026 and remain
excluded from this historical 14-case total. They use different timing boundaries:
the indexing cell reports a batch of 8,388,608 reads on a boxed array of 16,384
elements, including loop/checksum overhead; the JSON cell reports the complete
fixed corpus of 12 modules. [All sizes, phase timings, JS comparisons
and reproduction commands](docs/benchmark-results/2026-09-21-go-diagnostics.md)
are recorded separately. Go uses GOMAXPROCS=1 in these diagnostics; JS retains its
default runtime background threads. “—” means no equivalent measurement was made
for that column.

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) | Hand-written imperative Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation | ~ 0.073544 μs | ~ 0.075615 μs | ~ 0.068871 μs |
Fibonacci | ~ 0.147690 μs | ~ 0.147720 μs | ~ 0.144287 μs |
List Processing | ~ 6.352051 μs | ~ 3.759033 μs | ~ 2.835693 μs |
Tail Call Optimization | ~ 277.125000 μs | ~ 322.343750 μs | ~ 329.875000 μs |
Deep Record Updates | ~ 228.437500 μs | ~ 48.070312 μs | ~ 43.574219 μs |
Ackermann | ~ 16.992188 μs | ~ 9.072266 μs | ~ 9.141113 μs |
Church Numerals | ~ 237.125000 μs | ~ 235.531250 μs | ~ 62.257812 μs |
Prime Sieve | ~ 37.494141 μs | ~ 22.679688 μs | ~ 0.799194 μs |
Red-Black Tree | ~ 22551.000000 μs | ~ 15428.000000 μs | ~ 11970.000000 μs |
Polymorphism | ~ 17812.000000 μs | ~ 27373.000000 μs | ~ 6730.500000 μs |
State Monad | ~ 1.700684 μs | ~ 2.024170 μs | ~ 0.840210 μs |
Lazy Evaluation | ~ 2989.500000 μs | ~ 1964.000000 μs | ~ 0.609741 μs |
Array Processing | ~ 7.291504 μs | ~ 5.041504 μs | ~ 2.793213 μs |
RowToList | ~ 0.000998 μs | ~ 0.014573 μs | ~ 0.000999 μs |
**Total Execution Time** | ~ 44.165240 ms | ~ 45.413760 ms | ~ 19.153440 ms |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT<br><br>(WIP) |
----------------------- | ----------------| -------------------------- | -------------------------------- |
AST Evaluation | ~ 0.110652 μs | ~ 0.114168 μs | ~ 0.100370 μs |
Fibonacci | ~ 0.208932 μs | ~ 0.231213 μs | ~ 0.242216 μs |
List Processing | ~ 37.319256 μs | ~ 5.221883 μs | ~ 0.858925 μs |
Tail Call Optimization | ~ 1341.296875 μs | ~ 103.949547 μs | ~ 104.623695 μs |
Deep Record Updates | ~ 404.766906 μs | ~ 1276.395875 μs | ~ 15.463379 μs |
Ackermann | ~ 19.244059 μs | ~ 18.813803 μs | ~ 19.677082 μs |
Church Numerals | ~ 520.080687 μs | ~ 418.191406 μs | ~ 0.011910 μs |
Prime Sieve | ~ 148.035156 μs | ~ 30.953043 μs | ~ 36.683756 μs |
Red-Black Tree | ~ 16911.709000 μs | ~ 14282.542000 μs | ~ 14426.500000 μs |
Polymorphism | ~ 66666.375000 μs | ~ 52475.291000 μs | ~ 0.009724 μs |
State Monad | ~ 20.410482 μs | ~ 14.547241 μs | ~ 7.614156 μs |
Lazy Evaluation | ~ 8739.062500 μs | ~ 6780.021000 μs | ~ 0.010030 μs |
Array Processing | ~ 27.824787 μs | ~ 16.206787 μs | ~ 1.043688 μs |
RowToList | ~ 0.039801 μs | ~ 0.044050 μs | ~ 0.009521 μs |
**Total Execution Time** | ~ 94.836484 ms | ~ 75.422523 ms | ~ 14.612848 ms |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT<br><br>(WIP) |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
AST Evaluation | ~ 2.276372 μs | ~ 3.195791 μs | ~ 2.109731 μs |
Fibonacci | ~ 1.082995 μs | ~ 2.255086 μs | ~ 1.252284 μs |
List Processing | ~ 162.933594 μs | ~ 178.479172 μs | ~ 0.631203 μs |
Tail Call Optimization | ~ 74.793293 μs | ~ 50.012367 μs | ~ 50.135906 μs |
Deep Record Updates | ~ 1495.614625 μs | ~ 1721.760375 μs | ~ 494.089844 μs |
Ackermann | ~ 48.438965 μs | ~ 99.437820 μs | ~ 53.379230 μs |
Church Numerals | ~ 2114.473875 μs | ~ 6731.145500 μs | ~ 32.569988 μs |
Prime Sieve | ~ 467.088563 μs | ~ 1017.802063 μs | ~ 1.887258 μs |
Red-Black Tree | ~ 250560.250000 μs | ~ 508200.125000 μs | ~ 115687.750000 μs |
Polymorphism | ~ 6484.645500 μs | ~ 429315.459000 μs | ~ 88163.958000 μs |
State Monad | ~ 494.477875 μs | ~ 567.742188 μs | ~ 0.488331 μs |
Lazy Evaluation | ~ 802.460938 μs | ~ 109586.125000 μs | ~ 328.287750 μs |
Array Processing | ~ 99.269203 μs | ~ 27.265381 μs | ~ 0.773900 μs |
RowToList | ~ 0.836090 μs | ~ 0.882133 μs | ~ 0.077305 μs |
**Total Execution Time** | ~ 262.808642 ms | ~ 1057.501687 ms | ~ 204.817391 ms |

#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br><br>(WIP) | PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br><br>(WIP) | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3 | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br><br>(WIP) | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br><br>(WIP) |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
AST Evaluation | ~ 0.266511 μs | ~ 4.412313 μs | ~ 0.255287 μs | ~ 0.264587 μs | ~ 0.258017 μs |
Fibonacci | ~ 0.084319 μs | ~ 12.686319 μs | ~ 0.100370 μs | ~ 0.103848 μs | ~ 0.117165 μs |
List Processing | ~ 31.668863 μs | ~ 546.453125 μs | ~ 16.268636 μs | ~ 11.379231 μs | ~ 0.492350 μs |
Tail Call Optimization | ~ 30.373779 μs | ~ 15.894124 μs | ~ 10.678508 μs | ~ 33.669434 μs | ~ 34.145752 μs |
Deep Record Updates | ~ 3.259430 μs | ~ 10530.500000 μs | ~ 116.572586 μs | ~ 115.089844 μs | ~ 3.420349 μs |
Ackermann | ~ 16.646525 μs | ~ 757.776062 μs | ~ 15.443033 μs | ~ 17.577148 μs | ~ 16.748860 μs |
Church Numerals | ~ 168.976562 μs | ~ 11906.750000 μs | ~ 2068.119750 μs | ~ 1154.705688 μs | ~ 0.013727 μs |
Prime Sieve | ~ 154.542320 μs | ~ 2549.135500 μs | ~ 115.475586 μs | ~ 82.351242 μs | ~ 0.748619 μs |
Red-Black Tree | ~ 8529.000000 μs | ~ 62062.958000 μs | ~ 52997.625000 μs | ~ 31816.709000 μs | ~ 30705.083000 μs |
Polymorphism | ~ 0.041052 μs | ~ 1903090.875000 μs | ~ 7298.333500 μs | ~ 7131.500000 μs | ~ 0.014572 μs |
State Monad | ~ 54.527832 μs | ~ 932.015625 μs | ~ 81.707031 μs | ~ 34.783529 μs | ~ 0.013889 μs |
Lazy Evaluation | ~ 0.000702 μs | ~ 18989.708000 μs | ~ 37387.542000 μs | ~ 22062.083000 μs | ~ 0.014146 μs |
Array Processing | ~ 23.612631 μs | ~ 362.148438 μs | ~ 3.133240 μs | ~ 0.579601 μs | ~ 0.247798 μs |
RowToList | ~ 0.063045 μs | ~ 1.130124 μs | ~ 0.108747 μs | ~ 0.015588 μs | ~ 0.013889 μs |
**Total Execution Time** | ~ 9.013064 ms | ~ 2011.762443 ms | ~ 100.111363 ms | ~ 62.460812 ms | ~ 30.761332 ms |

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written FP-style C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) | Hand-written imperative C++ FFI<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation | ~ 2.711395 μs | ~ 0.939227 μs | ~ 0.852976 μs |
Fibonacci | ~ 0.314994 μs | ~ 0.107427 μs | ~ 0.024341 μs |
List Processing | ~ 154.890625 μs | ~ 35.161621 μs | ~ 0.358898 μs |
Tail Call Optimization | ~ 2844.593750 μs | ~ 35.114096 μs | ~ 33.846680 μs |
Deep Record Updates | ~ 1453.145750 μs | ~ 567.682250 μs | ~ 3.409556 μs |
Ackermann | ~ 221.500641 μs | ~ 17.989665 μs | ~ 16.538249 μs |
Church Numerals | ~ 2630.260250 μs | ~ 4240.885250 μs | ~ 0.021796 μs |
Prime Sieve | ~ 974.453125 μs | ~ 257.703125 μs | ~ 1.080015 μs |
Red-Black Tree | ~ 613388.666000 μs | ~ 66049.458000 μs | ~ 22007.125000 μs |
Polymorphism | ~ 265261.000000 μs | ~ 27358.792000 μs | ~ 0.022254 μs |
State Monad | ~ 293.296875 μs | ~ 185.699219 μs | ~ 0.022342 μs |
Lazy Evaluation | ~ 36615.584000 μs | ~ 31474.584000 μs | ~ 0.023396 μs |
Array Processing | ~ 36.573404 μs | ~ 1.095034 μs | ~ 0.735326 μs |
RowToList | ~ 0.482352 μs | ~ 0.296592 μs | ~ 0.022392 μs |
**Total Execution Time** | ~ 923.877473 ms | ~ 130.225508 ms | ~ 22.064083 ms |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) | Hand-written imperative F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation | ~ 0.743284 μs | ~ 0.490854 μs | ~ 0.497316 μs |
Fibonacci | ~ 1.541061 μs | ~ 0.115755 μs | ~ 0.115859 μs |
List Processing | ~ 202.765625 μs | ~ 16.253581 μs | ~ 2.959981 μs |
Tail Call Optimization | ~ 50.342285 μs | ~ 38.425617 μs | ~ 44.323566 μs |
Deep Record Updates | ~ 2158.468750 μs | ~ 91.296547 μs | ~ 9.131267 μs |
Ackermann | ~ 109.867188 μs | ~ 40.018066 μs | ~ 39.784016 μs |
Church Numerals | ~ 1624.307375 μs | ~ 352.321625 μs | ~ 25.649984 μs |
Prime Sieve | ~ 405.609375 μs | ~ 157.565094 μs | ~ 15.448039 μs |
Red-Black Tree | ~ 41097.666000 μs | ~ 37763.917000 μs | ~ 10551.666000 μs |
Polymorphism | ~ 287909.250000 μs | ~ 25932.334000 μs | ~ 2332.765625 μs |
State Monad | ~ 172.506516 μs | ~ 79.713867 μs | ~ 3.319326 μs |
Lazy Evaluation | ~ 5992.229000 μs | ~ 10061.000000 μs | ~ 233.031250 μs |
Array Processing | ~ 53.159340 μs | ~ 31.737305 μs | ~ 16.088583 μs |
RowToList | ~ 0.257264 μs | ~ 0.059500 μs | ~ 0.025040 μs |
**Total Execution Time** | ~ 339.778713 ms | ~ 74.565249 ms | ~ 13.274806 ms |

#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT<br><br>(WIP) |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
AST Evaluation | ~ 0.143532 μs | ~ 0.106122 μs | ~ 0.101427 μs |
Fibonacci | ~ 0.258853 μs | ~ 0.098805 μs | ~ 0.076588 μs |
List Processing | ~ 5.030192 μs | ~ 4.183125 μs | ~ 0.567115 μs |
Tail Call Optimization | ~ 45.635414 μs | ~ 38.535969 μs | ~ 39.831137 μs |
Deep Record Updates | ~ 73.279945 μs | ~ 38.798504 μs | ~ 4.178223 μs |
Ackermann | ~ 18.179606 μs | ~ 7.595113 μs | ~ 5.797079 μs |
Church Numerals | ~ 310.509125 μs | ~ 375.716156 μs | ~ 0.003118 μs |
Prime Sieve | ~ 21.288414 μs | ~ 18.400187 μs | ~ 0.541996 μs |
Red-Black Tree | ~ 16788.375000 μs | ~ 12065.750000 μs | ~ 11828.167000 μs |
Polymorphism | ~ 0.005413 μs | ~ 14583.083000 μs | ~ 0.003362 μs |
State Monad | ~ 6.256673 μs | ~ 17.249023 μs | ~ 0.003258 μs |
Lazy Evaluation | ~ 5.242452 μs | ~ 8062.000000 μs | ~ 0.036899 μs |
Array Processing | ~ 2.957743 μs | ~ 3.063232 μs | ~ 0.497445 μs |
RowToList | ~ 0.012720 μs | ~ 0.006162 μs | ~ 0.002328 μs |
**Total Execution Time** | ~ 17.277175 ms | ~ 35.214585 ms | ~ 11.879807 ms |

#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>koka -O3 |
----------------------- | ----- |
AST Evaluation | ~ 0.218750 μs |
Fibonacci | ~ 0.123459 μs |
List Processing | ~ 4.275391 μs |
Tail Call Optimization | ~ 80.507812 μs |
Deep Record Updates | ~ 158.937500 μs |
Ackermann | ~ 19.376953 μs |
Church Numerals | ~ 758.687500 μs |
Prime Sieve | ~ 12.167969 μs |
Red-Black Tree | ~ 8746.500000 μs |
Polymorphism | ~ 35687.000000 μs |
State Monad | ~ 5.206055 μs |
Lazy Evaluation | ~ 8982.000000 μs |
Array Processing | ~ 6.578613 μs |
RowToList | ~ 0.001033 μs |
**Total Execution Time** | ~ 54.461581 ms |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2 |
----------------------- | ------------------------ |
AST Evaluation | ~ 0.055759 μs |
Fibonacci | ~ 0.194260 μs |
List Processing | ~ 3.579102 μs |
Tail Call Optimization | ~ 62.843750 μs |
Deep Record Updates | ~ 7.678711 μs |
Ackermann | ~ 6.749512 μs |
Church Numerals | ~ 143.898438 μs |
Prime Sieve | ~ 21.582031 μs |
Red-Black Tree | ~ 11559.000000 μs |
Polymorphism | ~ 4883.000000 μs |
State Monad | ~ 0.015631 μs |
Lazy Evaluation | ~ 0.456085 μs |
Array Processing | ~ 4.326660 μs |
RowToList | ~ 0.004070 μs |
**Total Execution Time** | ~ 16.693384 ms |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3 |
----------------------- | --------------------------- |
AST Evaluation | ~ 0.049465 μs |
Fibonacci | ~ 0.124039 μs |
List Processing | ~ 1.976562 μs |
Tail Call Optimization | ~ 49.824219 μs |
Deep Record Updates | ~ 15.847656 μs |
Ackermann | ~ 16.053711 μs |
Church Numerals | ~ 140.375000 μs |
Prime Sieve | ~ 17.937500 μs |
Red-Black Tree | ~ 10787.000000 μs |
Polymorphism | ~ 12448.000000 μs |
State Monad | ~ 10.026367 μs |
Lazy Evaluation | ~ 5690.000000 μs |
Array Processing | ~ 5.893066 μs |
RowToList | ~ 0.004679 μs |
**Total Execution Time** | ~ 29.183112 ms |

#### C (reference)

C Benchmark             | Hand-written imperative C<br>↓<br>clang -O3 |
----------------------- | -------------------- |
AST Evaluation | ~ 0.096481 μs |
Fibonacci | ~ 0.081787 μs |
List Processing | ~ 0.048588 μs |
Tail Call Optimization | ~ 33.777344 μs |
Deep Record Updates | ~ 3.406494 μs |
Ackermann | ~ 17.351562 μs |
Church Numerals | ~ 0.000715 μs |
Prime Sieve | ~ 1.068115 μs |
Red-Black Tree | ~ 9788.000000 μs |
Polymorphism | ~ 0.000710 μs |
State Monad | ~ 0.000706 μs |
Lazy Evaluation | ~ 0.000700 μs |
Array Processing | ~ 0.048042 μs |
RowToList | ~ 0.000694 μs |
**Total Execution Time** | ~ 9.843882 ms |

### Extended benchmark results (I/O, mutability, async)

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build<br><br>(mature WIP) | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br><br>(WIP)
----------------------- | ------------- | -------------- | --------------- | ---------------
File I/O | ~ 466899.834000 μs | ~ 477117.792000 μs | ~ 442703.250000 μs | ~ 468566.459000 μs |
STArray Operations | ~ 0.916000 μs | ~ 0.542000 μs | ~ 0.541000 μs | ~ 0.333000 μs |
String Operations | ~ 256.708000 μs | ~ 232.500000 μs | ~ 507.083000 μs | ~ 445.084000 μs |
Aff Operations | ~ 10125.208000 μs | ~ 11581.833000 μs | ~ 11009.250000 μs | ~ 12510.792000 μs |
Parallelism | ~ 14275052.041000 μs | ~ 14278260.458000 μs | ~ 1182150.084000 μs | ~ 574110.166000 μs |
**Total Execution Time** | ~ 14752.334707 ms | ~ 14767.193125 ms | ~ 1636.370208 ms | ~ 1055.632834 ms |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with Go/Rust's multicore runtimes.

[Methodology and comparison limits](docs/benchmark-methodology.md) · [Validated measurements — September 20, 2026](docs/benchmark-results/2026-09-20.json). All 39 columns were recomputed; these figures form a new baseline rather than a compiler speedup comparison with the older protocol.
