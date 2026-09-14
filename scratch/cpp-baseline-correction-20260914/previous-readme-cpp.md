#### C++

C++ Benchmark           | Compiled C++ ([pscpp](https://github.com/purescript-native/purescript), WIP) | Native FP-style C++ FFI (WIP) | Native hand-written C++ FFI (WIP) |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
AST Evaluation          | ~ 0.03 μs                 | ~ 0.01 μs                 | ~ 0.00 μs                       |
Fibonacci               | ~ 0.01 μs                 | ~ 0.00 μs                 | ~ 0.00 μs                       |
List Processing         | ~ 1.39 μs                 | ~ 0.28 μs                 | ~ 0.02 μs                       |
Tail Call Optimization  | ~ 25.06 μs                | ~ 0.61 μs                 | ~ 0.04 μs                       |
Deep Record Updates     | ~ 14.78 μs                | ~ 0.82 μs                 | ~ 0.01 μs                       |
Ackermann               | ~ 1.39 μs                 | ~ 0.04 μs                 | ~ 0.03 μs                       |
Church Numerals         | ~ 22.36 μs                | ~ 116.20 μs               | ~ 0.00 μs                       |
Prime Sieve             | ~ 9.32 μs                 | ~ 0.76 μs                 | ~ 0.01 μs                       |
Red-Black Tree          | ~ 6076.46 μs              | ~ 459.54 μs               | ~ 22.99 μs                      |
Polymorphism            | ~ 3088.77 μs              | ~ 280.09 μs               | ~ 4.95 μs                       |
State Monad             | ~ 2.47 μs                 | ~ 0.63 μs                 | ~ 0.00 μs                       |
Lazy Evaluation         | ~ 341.13 μs               | ~ 155.60 μs               | ~ 0.00 μs                       |
Array Processing        | ~ 0.29 μs                 | ~ 1.94 μs                 | ~ 0.02 μs                       |
RowToList               | ~ 0.01 μs                 | ~ 0.00 μs                 | ~ 0.00 μs                       |
**Total Execution Time**| ~ 9.58 ms                 | ~ 1.02 ms                 | ~ 0.03 ms                       |

