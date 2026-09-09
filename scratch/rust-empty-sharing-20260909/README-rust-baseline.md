# Baseline officielle relue le 9 septembre 2026

Source : [README altbak.pub](../../../altbak.pub/README.md). Le natif n’a pas été remesuré dans cette étape.

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
