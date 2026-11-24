# C reference precision — 24 November 2026

All 14 core reference timings are positive and validated in three independent
processes. The five timings previously published as `< 0.01 μs` can now be
included in ratios using their recorded values. Raw times, outputs, batch counts,
compiler version and hashes are in the [measurement archive](2026-11-24-c-reference.json).

| Case | Median (µs, display precision) | Three process minima (µs) |
| --- | ---: | --- |
| AST Evaluation | 0.090 | 0.089943, 0.086990, 0.090553 |
| Fibonacci | 0.077 | 0.075287, 0.077171, 0.077148 |
| List Processing | 0.046 | 0.045300, 0.046219, 0.046761 |
| Tail Call Optimization | 32.14 | 31.869141, 32.140625, 32.833984 |
| Deep Record Updates | 3.25 | 3.209717, 3.272949, 3.252930 |
| Ackermann | 16.83 | 16.883789, 16.828125, 16.432617 |
| Church Numerals | 0.00068 | 0.000688, 0.000678, 0.000669 |
| Prime Sieve | 1.01 | 1.016541, 0.996887, 1.005493 |
| Red-Black Tree | 9053.50 | 9185.000000, 9053.500000, 9027.500000 |
| Polymorphism | 0.00067 | 0.000670, 0.000677, 0.000667 |
| State Monad | 0.00067 | 0.000668, 0.000669, 0.000668 |
| Lazy Evaluation | 0.00067 | 0.000665, 0.000697, 0.000668 |
| Array Processing | 0.046 | 0.045197, 0.046844, 0.045589 |
| RowToList | 0.00067 | 0.000667, 0.000673, 0.000668 |

Total of median rows: **9.106989424 ms**.

## Protocol and provenance

Apple M4 Pro, darwin/arm64, Apple clang 17.0.0, `clang -O3`. The source
`tmp/bench_native.c` was copied unchanged into `/tmp/altbak-c-reference-20261124`
and compiled once. Each process performs three global warmups, three local
invocations, power-of-two calibration toward 10 ms (maximum 16,777,216
iterations), then ten timed batches. The table uses the median of three
per-process minimum times per call. All 42 returned values and batch counts
passed the ordinary benchmark validator; `publish.verify` accepted the aggregate.

The compiler input SHA-256 is `66156bcdee70d092c9f31298ea97aabac08e794760c8fd7a0176c9014a535af0`.
The executable SHA-256 is `78e246f60bc026b10f42686ea6eb27923f2676b7232056721cbf4858fa8e2162`.
The driver prints six decimal places in microseconds; the display keeps two
significant fractional digits after leading zeros. Ratios use archived values,
not the shortened display.

## Interpretation and retained limitations

This is a fresh reference campaign, distinct from the previous historical C
column. References for the excluded JSON and Array Indexing diagnostics were
not remeasured and are unchanged. Updating the core C reference consequently
updates the core `/C` ratios in every language table.

The C driver retains its historical imperative algorithms and allocation model.
In particular, AST and red-black tree arenas are allocated before timing and
reset/reused across calls. No arena, pool, benchmark-source edit or arithmetic
shortcut was introduced for this measurement. These comparisons include the
existing source-level choices; they are not allocator-equivalent measurements
of the generated functional program.

The five extremely small reference kernels are heavily simplified by the C
compiler. Their restored ratios can therefore be very large. Removing the old
formatting bound reveals these comparisons; it does not establish a new
regression in a compiled backend. Genuine missing data or bounds in other
columns must still retain their incomplete-coverage annotation.
