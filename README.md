# 🌈 PureScript universal multi-runtime benchmark

## Project goal
This project is a **proof of concept** demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. Backends consume either standard `CoreFn` or, for the local AOT backends, the typed `TAST/tcorefn` representation exported by a local PureScript fork (may be subject to an official PR soon).

## Comprehensive benchmarks
The benchmark suite runs a **wide variety** of computationally intensive tasks: AST evaluation, massive list processing, tail call optimization, deep record updates... These tests apply massive **pressure** on the call stack, garbage collector, disk I/O, event loops, and runtime execution engines to measure the raw ability of the compiler and the underlying virtual machine. 

> [!NOTE]
> **Reading the benchmark tables**
>
> - **Status, only for gopurs, purust, phpurs, sharpurs and javapurs:** **Stable** = measured and considered fully optimized; **WIP** = measured, with optimization work remaining; **Backlog** = planned work without a published measurement or ratio. Status labels belong only to the relevant backend columns. References and external backends have no status. Stable is currently empty, including for Go: the available analyses still identify optimization opportunities.
> - **Ratios:** `/C` uses the optimized imperative C reference; `/JS` uses the first JavaScript column. **Below 1× is faster; above 1× is slower.** Extended Results uses its own JS reference and has no C reference.
> - **Totals:** sums of the available per-test medians: **Total WIP** for the controlled backends, **Total** for references and external backends. Backlog never contributes. Empty Stable groups are shown as *Coming soon...* and have no defined total, ratio or mean. These totals do not represent an application benchmark.
> - **Summaries:** `μ` = arithmetic mean, `Q₂` = median, `G` = geometric mean of the case ratios. Each case has equal weight. All four ratios use exactly the same measured cases on both sides: Go/C and Go/JS both cover all 37 measured Go cases. The full total of a reference table is not used as a denominator for a different set of cases. Very small reference times can produce large ratios.
> - **Missing data:** `—` means unavailable or not applicable, never zero. Bounds (`< …`) are excluded from summaries. `n=a/b` indicates coverage: measured cases out of listed cases for a time total, or matched references out of measured cases for a ratio. Historical diagnostic measurements remain in the reference tables; their controlled-backend counterparts belong to unmeasured Backlog.
> - **Display precision, in every table:** values ≥ 1 use two decimal places (`12.000045 → 12.00`); positive values < 1 keep two significant digits after leading zeros (`0.000045`, `0.009492 → 0.0095`). Zero is `0.00`.
> - **Opaque inputs, in every table:** inputs must remain unknown to the optimizer at runtime. Algebraic simplification is allowed; substituting the driver’s known input is not.
> - **Calculations:** use archived timings before display rounding, including for threshold checks. `≈` marks a fallback to published precision.

### Core stresstest benchmark results (pure computational)

**Core tests (`src/`)** cover features common to all tested targets (`./bin/run`)

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.26 μs | ~ 0.19 μs | ~ 0.20 μs | ~ 0.22 μs |
| ↳ /C | 2.68x | 1.96x | 2.04x | 2.28x |
| ↳ /JS | — | 0.73x | 0.76x | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.30 μs | ~ 0.31 μs | ~ 0.31 μs | ~ 0.37 μs |
| ↳ /C | 3.83x | 3.92x | 3.97x | 4.68x |
| ↳ /JS | — | 1.02x | 1.04x | 1.22x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 8.64 μs | ~ 5.06 μs | ~ 19.06 μs | ~ 0.49 μs |
| ↳ /C | 7094.12x | 4155.08x | 15654.40x | 405.52x |
| ↳ /JS | — | 0.59x | 2.21x | 0.057x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 187.10 μs | ~ 151.13 μs | ~ 65.91 μs | ~ 68.31 μs |
| ↳ /C | 5.82x | 4.70x | 2.05x | 2.12x |
| ↳ /JS | — | 0.81x | 0.35x | 0.37x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 74.34 μs | ~ 134.11 μs | ~ 42.07 μs | ~ 7.69 μs |
| ↳ /C | 22.73x | 41.01x | 12.86x | 2.35x |
| ↳ /JS | — | 1.80x | 0.57x | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 66.77 μs | ~ 64.40 μs | ~ 38.25 μs | ~ 38.48 μs |
| ↳ /C | 4.40x | 4.25x | 2.52x | 2.54x |
| ↳ /JS | — | 0.96x | 0.57x | 0.58x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1269.89 μs | ~ 1242.29 μs | ~ 1337.51 μs | ~ 49.63 μs |
| ↳ /C | 1136880.58x | 1112176.99x | 1197422.70x | 44431.84x |
| ↳ /JS | — | 0.98x | 1.05x | 0.039x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 58.01 μs | ~ 26.99 μs | ~ 86.33 μs | ~ 0.65 μs |
| ↳ /C | 508.28x | 236.48x | 756.39x | 5.71x |
| ↳ /JS | — | 0.47x | 1.49x | 0.011x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 84263.21 μs | ~ 47272.88 μs | ~ 24486.83 μs | ~ 27249.58 μs |
| ↳ /C | 9.68x | 5.43x | 2.81x | 3.13x |
| ↳ /JS | — | 0.56x | 0.29x | 0.32x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 11836.88 μs | ~ 7368.38 μs | ~ 4843.00 μs | ~ 2805.25 μs |
| ↳ /C | 10514068.50x | 6544936.86x | 4301780.14x | 2491754.85x |
| ↳ /JS | — | 0.62x | 0.41x | 0.24x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 47.34 μs | ~ 10.95 μs | ~ 46.87 μs | ~ 0.42 μs |
| ↳ /C | 42288.91x | 9778.80x | 41870.02x | 373.21x |
| ↳ /JS | — | 0.23x | 0.99x | 0.0088x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 14058.08 μs | ~ 10604.88 μs | ~ 17513.12 μs | ~ 321.48 μs |
| ↳ /C | 12536838.89x | 9457306.51x | 15618005.03x | 286690.30x |
| ↳ /JS | — | 0.75x | 1.25x | 0.023x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.07 μs | ~ 3.75 μs | ~ 3.12 μs | ~ 0.50 μs |
| ↳ /C | 5002.97x | 2653.68x | 2208.19x | 352.58x |
| ↳ /JS | — | 0.53x | 0.44x | 0.070x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.035 μs | ~ 0.017 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 31.67x | 15.32x | 53.58x | 18.28x |
| ↳ /JS | — | 0.48x | 1.69x | 0.58x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 121.94 μs | ~ 144.60 μs | ~ 92.20 μs | ~ 12.93 μs |
| ↳ /C | 284.36x | 337.22x | 215.01x | 30.16x |
| ↳ /JS | — | 1.19x | 0.76x | 0.11x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 127.59 μs | ~ 130.46 μs | ~ 74.63 μs | ~ 18.58 μs |
| ↳ /C | 28.04x | 28.67x | 16.40x | 4.08x |
| ↳ /JS | — | 1.02x | 0.58x | 0.15x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 70.25 μs | ~ 70.22 μs | ~ 44.01 μs | ~ 3.32 μs |
| ↳ /C | 252.01x | 251.92x | 157.88x | 11.91x |
| ↳ /JS | — | 1.00x | 0.63x | 0.047x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 561.69 μs | ~ 517.87 μs | ~ 114.68 μs | ~ 78.30 μs |
| ↳ /C | 120.51x | 111.11x | 24.60x | 16.80x |
| ↳ /JS | — | 0.92x | 0.20x | 0.14x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 51.63 μs | ~ 50.71 μs | ~ 41.04 μs | ~ 5.03 μs |
| ↳ /C | 81.79x | 80.34x | 65.02x | 7.96x |
| ↳ /JS | — | 0.98x | 0.79x | 0.097x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.20 μs | ~ 31.05 μs | ~ 2.05 μs | ~ 2.40 μs |
| ↳ /C | 358.69x | 368.76x | 24.38x | 28.47x |
| ↳ /JS | — | 1.03x | 0.068x | 0.079x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 80.61 μs | ~ 81.04 μs | ~ 49.83 μs | ~ 1.39 μs |
| ↳ /C | 1908.96x | 1919.26x | 1180.16x | 32.89x |
| ↳ /JS | — | 1.01x | 0.62x | 0.017x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 68.35 μs | ~ 65.14 μs | ~ 42.18 μs | ~ 1.07 μs |
| ↳ /C | 50886.05x | 48496.96x | 31398.06x | 794.31x |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 68.01 μs | ~ 64.65 μs | ~ 42.45 μs | ~ 1.08 μs |
| ↳ /C | 208.29x | 197.98x | 130.01x | 3.30x |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.010 μs | ~ 0.0087 μs | ~ 0.019 μs | ~ 0.020 μs |
| ↳ /C | 9.05x | 7.81x | 17.38x | 17.52x |
| ↳ /JS | — | 0.86x | 1.92x | 1.94x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 37.44 μs | ~ 37.44 μs | ~ 37.09 μs | ~ 8.54 μs |
| ↳ /C | 34.14x | 34.14x | 33.82x | 7.79x |
| ↳ /JS | — | 1.00x | 0.99x | 0.23x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs |
| ↳ /C | 0.18x | 0.18x | 0.18x | 0.18x |
| ↳ /JS | — | 0.99x | 1.02x | 1.00x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 63.80 μs | ~ 63.97 μs | ~ 63.94 μs | ~ 63.72 μs |
| ↳ /C | 51.64x | 51.78x | 51.76x | 51.58x |
| ↳ /JS | — | 1.00x | 1.00x | 1.00x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 45.28 μs | ~ 44.96 μs | ~ 42.10 μs | ~ 21.01 μs |
| ↳ /C | 70.21x | 69.72x | 65.28x | 32.58x |
| ↳ /JS | — | 0.99x | 0.93x | 0.46x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 145.33 μs | ~ 144.87 μs | ~ 138.40 μs | ~ 24.69 μs |
| ↳ /C | 58.93x | 58.74x | 56.12x | 10.01x |
| ↳ /JS | — | 1.00x | 0.95x | 0.17x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 47.71 μs | ~ 46.52 μs | ~ 42.64 μs | ~ 11.85 μs |
| ↳ /C | 51.64x | 50.35x | 46.15x | 12.83x |
| ↳ /JS | — | 0.97x | 0.89x | 0.25x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 2.44 μs | ~ 4.37 μs | ~ 4.32 μs | ~ 0.72 μs |
| ↳ /C | 2175.36x | 3891.09x | 3849.97x | 645.38x |
| ↳ /JS | — | 1.79x | 1.77x | 0.30x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 4.55 μs | ~ 9.34 μs | ~ 9.04 μs | ~ 0.71 μs |
| ↳ /C | 4056.27x | 8323.17x | 8051.03x | 635.92x |
| ↳ /JS | — | 2.05x | 1.98x | 0.16x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 7.03 μs | ~ 14.73 μs | ~ 14.38 μs | ~ 0.72 μs |
| ↳ /C | 5183.02x | 10861.51x | 10601.86x | 531.73x |
| ↳ /JS | — | 2.10x | 2.05x | 0.10x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 463.87 μs | ~ 106.02 μs | ~ 2.44 μs | ~ 2.48 μs |
| ↳ /C | 411097.54x | 93960.55x | 2160.24x | 2196.39x |
| ↳ /JS | — | 0.23x | 0.0053x | 0.0053x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 513.20 μs | ~ 106.03 μs | ~ 31.67 μs | ~ 3.41 μs |
| ↳ /C | 455681.30x | 94147.95x | 28123.27x | 3026.46x |
| ↳ /JS | — | 0.21x | 0.062x | 0.0066x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 155.67 μs | ~ 102.38 μs | ~ 2.39 μs | ~ 2.56 μs |
| ↳ /C | 138832.60x | 91301.70x | 2129.67x | 2285.58x |
| ↳ /JS | — | 0.66x | 0.015x | 0.016x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 211.84 μs | ~ 113.53 μs | ~ 31.14 μs | ~ 3.52 μs |
| ↳ /C | 187275.56x | 100360.42x | 27531.34x | 3114.48x |
| ↳ /JS | — | 0.54x | 0.15x | 0.017x |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 25303.33 μs | — | — | — |
| ↳ /C | 3.84x | — | — | — |
| ↳ /JS | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 80646.58 μs | — | — | — |
| ↳ /C | 7.64x | — | — | — |
| ↳ /JS | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 9283.94 μs | — | — | — |
| ↳ /C | 14.12x | — | — | — |
| ↳ /JS | — | — | — | — |
| **Total** | <strong>229.99 ms</strong> | <strong>~ 68.84 ms</strong> (n=37/40) | <strong>~ 49.41 ms</strong> (n=37/40) | <strong>~ 30.81 ms</strong> (n=37/40) |
| ↳ /C | <strong>8.65x</strong> | <strong>7.84x</strong> | <strong>5.63x</strong> | <strong>3.51x</strong> |
| ↳ μ/C | <strong>637537.37x</strong> | <strong>475303.58x</strong> | <strong>575505.57x</strong> | <strong>76690.48x</strong> |
| ↳ Q₂/C | <strong>164.40x</strong> | <strong>236.48x</strong> | <strong>130.01x</strong> | <strong>28.47x</strong> |
| ↳ G/C | <strong>507.73x</strong> | <strong>581.52x</strong> | <strong>391.00x</strong> | <strong>74.58x</strong> |
| ↳ /JS | — | <strong>0.60x</strong> | <strong>0.43x</strong> | <strong>0.27x</strong> |
| ↳ μ/JS | — | <strong>0.92x</strong> | <strong>0.86x</strong> | <strong>0.29x</strong> |
| ↳ Q₂/JS | — | <strong>0.96x</strong> | <strong>0.76x</strong> | <strong>0.11x</strong> |
| ↳ G/JS | — | <strong>0.81x</strong> | <strong>0.55x</strong> | <strong>0.10x</strong> |


#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs) <br>`--dev`<br>↓<br>Go<br>↓<br>go-build<br>`-N -l` | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build |
----------------------- | -------------------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
|  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
|  | **WIP** | **WIP** |  | **WIP** | **WIP** |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.029 μs | ~ 0.52 μs | ~ 4.17 μs | ~ 0.32 μs | ~ 0.32 μs |
| ↳ /C | 0.30x | 5.42x | 43.53x | 3.38x | 3.32x |
| ↳ /JS | 0.11x | 2.02x | 16.22x | 1.26x | 1.24x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.13 μs | ~ 0.54 μs | ~ 5.41 μs | ~ 0.16 μs | ~ 0.16 μs |
| ↳ /C | 1.64x | 6.81x | 68.20x | 1.97x | 2.08x |
| ↳ /JS | 0.43x | 1.78x | 17.81x | 0.51x | 0.54x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.0042 μs | ~ 28.86 μs | ~ 236.01 μs | ~ 24.39 μs | ~ 0.45 μs |
| ↳ /C | 3.47x | 23709.26x | 193885.62x | 20034.74x | 369.12x |
| ↳ /JS | 0.00049x | 3.34x | 27.33x | 2.82x | 0.052x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 0.0037 μs | ~ 355.46 μs | ~ 6281.96 μs | ~ 35.98 μs | ~ 35.50 μs |
| ↳ /C | 0.00012x | 11.05x | 195.31x | 1.12x | 1.10x |
| ↳ /JS | 0.000020x | 1.90x | 33.58x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 0.017 μs | ~ 186.36 μs | ~ 5461.52 μs | ~ 40.43 μs | ~ 14.87 μs |
| ↳ /C | 0.0052x | 56.98x | 1669.93x | 12.36x | 4.55x |
| ↳ /JS | 0.00023x | 2.51x | 73.46x | 0.54x | 0.20x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.04 μs | ~ 35.56 μs | ~ 402.04 μs | ~ 20.85 μs | ~ 19.87 μs |
| ↳ /C | 1.06x | 2.35x | 26.52x | 1.38x | 1.31x |
| ↳ /JS | 0.24x | 0.53x | 6.02x | 0.31x | 0.30x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 0.0048 μs | ~ 1296.16 μs | ~ 3811.48 μs | ~ 1044.60 μs | ~ 22.45 μs |
| ↳ /C | 4.34x | 1160403.26x | 3412273.69x | 935194.87x | 20096.92x |
| ↳ /JS | 0.0000038x | 1.02x | 3.00x | 0.82x | 0.018x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 0.23 μs | ~ 76.41 μs | ~ 1340.99 μs | ~ 110.87 μs | ~ 0.62 μs |
| ↳ /C | 2.00x | 669.51x | 11749.91x | 971.48x | 5.47x |
| ↳ /JS | 0.0039x | 1.32x | 23.12x | 1.91x | 0.011x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 9474.08 μs | ~ 31239.75 μs | ~ 870404.62 μs | ~ 24630.67 μs | ~ 9004.48 μs |
| ↳ /C | 1.09x | 3.59x | 99.95x | 2.83x | 1.03x |
| ↳ /JS | 0.11x | 0.37x | 10.33x | 0.29x | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0035 μs | ~ 15617.54 μs | ~ 505069.00 μs | ~ 49160.92 μs | ~ 2244.45 μs |
| ↳ /C | 3.08x | 13872232.78x | 448626015.16x | 43667035.29x | 1993628.72x |
| ↳ /JS | 0.00000029x | 1.32x | 42.67x | 4.15x | 0.19x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.0034 μs | ~ 110.57 μs | ~ 303.56 μs | ~ 52.10 μs | ~ 0.33 μs |
| ↳ /C | 3.07x | 98773.04x | 271171.26x | 46545.06x | 298.08x |
| ↳ /JS | 0.000073x | 2.34x | 6.41x | 1.10x | 0.0070x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.0035 μs | ~ 1566.95 μs | ~ 71378.08 μs | ~ 13865.38 μs | ~ 0.27 μs |
| ↳ /C | 3.09x | 1397388.13x | 63654159.93x | 12364983.21x | 239.57x |
| ↳ /JS | 0.00000025x | 0.11x | 5.08x | 0.99x | 0.000019x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.0040 μs | ~ 32.70 μs | ~ 43.81 μs | ~ 4.90 μs | ~ 0.45 μs |
| ↳ /C | 2.83x | 23132.69x | 30993.45x | 3468.57x | 319.42x |
| ↳ /JS | 0.00057x | 4.62x | 6.20x | 0.69x | 0.064x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0030 μs | ~ 0.30 μs | ~ 0.50 μs | ~ 0.073 μs | ~ 0.032 μs |
| ↳ /C | 2.66x | 268.01x | 450.32x | 64.80x | 28.55x |
| ↳ /JS | 0.084x | 8.46x | 14.22x | 2.05x | 0.90x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 4.18 μs | ~ 1452.41 μs | ~ 2146.41 μs | ~ 77.41 μs | ~ 6.61 μs |
| ↳ /C | 9.76x | 3387.13x | 5005.58x | 180.53x | 15.42x |
| ↳ /JS | 0.034x | 11.91x | 17.60x | 0.63x | 0.054x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 12.98 μs | ~ 828.92 μs | ~ 1330.39 μs | ~ 67.38 μs | ~ 6.85 μs |
| ↳ /C | 2.85x | 182.18x | 292.39x | 14.81x | 1.50x |
| ↳ /JS | 0.10x | 6.50x | 10.43x | 0.53x | 0.054x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 1.71 μs | ~ 281.15 μs | ~ 133.05 μs | ~ 49.77 μs | ~ 2.31 μs |
| ↳ /C | 6.13x | 1008.55x | 477.29x | 178.55x | 8.29x |
| ↳ /JS | 0.024x | 4.00x | 1.89x | 0.71x | 0.033x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 31.47 μs | ~ 12300.38 μs | ~ 11626.08 μs | ~ 825.35 μs | ~ 10.36 μs |
| ↳ /C | 6.75x | 2639.06x | 2494.39x | 177.08x | 2.22x |
| ↳ /JS | 0.056x | 21.90x | 20.70x | 1.47x | 0.018x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 3.86 μs | ~ 752.15 μs | ~ 98.90 μs | ~ 48.10 μs | ~ 2.53 μs |
| ↳ /C | 6.12x | 1191.57x | 156.68x | 76.20x | 4.01x |
| ↳ /JS | 0.075x | 14.57x | 1.92x | 0.93x | 0.049x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 1.05 μs | ~ 304.51 μs | ~ 372.03 μs | ~ 13.95 μs | ~ 1.25 μs |
| ↳ /C | 12.47x | 3616.33x | 4418.08x | 165.67x | 14.89x |
| ↳ /JS | 0.035x | 10.08x | 12.32x | 0.46x | 0.042x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 0.39 μs | ~ 464.65 μs | ~ 476.59 μs | ~ 18.97 μs | ~ 0.87 μs |
| ↳ /C | 9.30x | 11004.26x | 11286.91x | 449.27x | 20.59x |
| ↳ /JS | 0.0049x | 5.76x | 5.91x | 0.24x | 0.011x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 0.011 μs | ~ 298.48 μs | ~ 303.08 μs | ~ 11.64 μs | ~ 0.74 μs |
| ↳ /C | 8.53x | 222208.13x | 225635.26x | 8666.26x | 554.05x |
| ↳ /JS | 0.00017x | 4.37x | 4.43x | 0.17x | 0.011x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 0.78 μs | ~ 297.10 μs | ~ 307.26 μs | ~ 11.44 μs | ~ 0.74 μs |
| ↳ /C | 2.40x | 909.85x | 940.96x | 35.04x | 2.26x |
| ↳ /JS | 0.012x | 4.37x | 4.52x | 0.17x | 0.011x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.0089 μs | ~ 17.05 μs | ~ 3.47 μs | ~ 3.45 μs | ~ 3.45 μs |
| ↳ /C | 7.99x | 15273.19x | 3109.50x | 3089.72x | 3090.34x |
| ↳ /JS | 0.88x | 1687.77x | 343.62x | 341.43x | 341.50x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 6.50 μs | ~ 79.96 μs | ~ 62.44 μs | ~ 7.20 μs | ~ 3.41 μs |
| ↳ /C | 5.93x | 72.92x | 56.93x | 6.57x | 3.11x |
| ↳ /JS | 0.17x | 2.14x | 1.67x | 0.19x | 0.091x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.31 μs | ~ 130.75 μs | ~ 2.71 μs | ~ 0.17 μs | ~ 0.17 μs |
| ↳ /C | 0.29x | 121.49x | 2.52x | 0.16x | 0.16x |
| ↳ /JS | 1.64x | 685.42x | 14.19x | 0.90x | 0.88x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 2.42 μs | ~ 140.42 μs | ~ 2.47 μs | ~ 2.31 μs | ~ 2.32 μs |
| ↳ /C | 1.96x | 113.67x | 2.00x | 1.87x | 1.88x |
| ↳ /JS | 0.038x | 2.20x | 0.039x | 0.036x | 0.036x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 3.24 μs | ~ 105.92 μs | ~ 104.16 μs | ~ 146.89 μs | ~ 3.66 μs |
| ↳ /C | 5.03x | 164.24x | 161.52x | 227.77x | 5.68x |
| ↳ /JS | 0.072x | 2.34x | 2.30x | 3.24x | 0.081x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 11.50 μs | ~ 318.05 μs | ~ 300.93 μs | ~ 32.68 μs | ~ 9.75 μs |
| ↳ /C | 4.66x | 128.96x | 122.02x | 13.25x | 3.95x |
| ↳ /JS | 0.079x | 2.19x | 2.07x | 0.22x | 0.067x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 1.24 μs | ~ 216.07 μs | ~ 167.98 μs | ~ 74.43 μs | ~ 3.71 μs |
| ↳ /C | 1.34x | 233.87x | 181.82x | 80.56x | 4.01x |
| ↳ /JS | 0.026x | 4.53x | 3.52x | 1.56x | 0.078x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.0085 μs | ~ 5.46 μs | ~ 144.74 μs | ~ 0.72 μs | ~ 0.24 μs |
| ↳ /C | 7.57x | 4862.25x | 128844.23x | 639.88x | 217.10x |
| ↳ /JS | 0.0035x | 2.24x | 59.23x | 0.29x | 0.10x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 0.0098 μs | ~ 13.30 μs | ~ 190.03 μs | ~ 7.20 μs | ~ 0.26 μs |
| ↳ /C | 8.70x | 11850.30x | 169290.03x | 6416.55x | 229.44x |
| ↳ /JS | 0.0021x | 2.92x | 41.74x | 1.58x | 0.057x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 0.012 μs | ~ 22.40 μs | ~ 679.42 μs | ~ 12.85 μs | ~ 1.03 μs |
| ↳ /C | 8.90x | 16521.13x | 501044.55x | 9479.91x | 762.15x |
| ↳ /JS | 0.0017x | 3.19x | 96.67x | 1.83x | 0.15x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 0.0062 μs | ~ 1019.37 μs | ~ 3085.81 μs | ~ 2.25 μs | ~ 2.25 μs |
| ↳ /C | 5.52x | 903393.56x | 2734738.94x | 1994.28x | 1996.75x |
| ↳ /JS | 0.000013x | 2.20x | 6.65x | 0.0049x | 0.0049x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 0.0065 μs | ~ 1073.35 μs | ~ 4068.22 μs | ~ 2.33 μs | ~ 2.38 μs |
| ↳ /C | 5.76x | 953046.15x | 3612245.82x | 2066.17x | 2111.26x |
| ↳ /JS | 0.000013x | 2.09x | 7.93x | 0.0045x | 0.0046x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 0.0063 μs | ~ 2199.81 μs | ~ 3135.15 μs | ~ 2.26 μs | ~ 2.26 μs |
| ↳ /C | 5.66x | 1961867.08x | 2796035.37x | 2012.95x | 2012.31x |
| ↳ /JS | 0.000041x | 14.13x | 20.14x | 0.014x | 0.014x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 0.0064 μs | ~ 4998.00 μs | ~ 4026.20 μs | ~ 2.36 μs | ~ 2.36 μs |
| ↳ /C | 5.62x | 4418406.87x | 3559299.68x | 2088.43x | 2087.72x |
| ↳ /JS | 0.000030x | 23.59x | 19.01x | 0.011x | 0.011x |
| **Total WIP** | <strong>~ 9.57 ms</strong> | <strong>77.87 ms</strong> |  | <strong>~ 90.41 ms</strong> | <strong>~ 11.41 ms</strong> |
| ↳ /C | <strong>1.09x</strong> | <strong>8.87x</strong> |  | <strong>10.30x</strong> | <strong>1.30x</strong> |
| ↳ μ/C | <strong>4.54x</strong> | <strong>678617.99x</strong> |  | <strong>1542604.93x</strong> | <strong>54814.82x</strong> |
| ↳ Q₂/C | <strong>4.34x</strong> | <strong>2639.06x</strong> |  | <strong>180.53x</strong> | <strong>14.89x</strong> |
| ↳ G/C | <strong>2.29x</strong> | <strong>2982.96x</strong> |  | <strong>341.32x</strong> | <strong>40.42x</strong> |
| ↳ /JS | <strong>0.083x</strong> | <strong>0.68x</strong> |  | <strong>0.79x</strong> | <strong>0.099x</strong> |
| ↳ μ/JS | <strong>0.11x</strong> | <strong>68.97x</strong> |  | <strong>10.12x</strong> | <strong>9.38x</strong> |
| ↳ Q₂/JS | <strong>0.012x</strong> | <strong>2.92x</strong> |  | <strong>0.63x</strong> | <strong>0.054x</strong> |
| ↳ G/JS | <strong>0.0032x</strong> | <strong>4.17x</strong> |  | <strong>0.48x</strong> | <strong>0.057x</strong> |
|  | **Backlog** | **Backlog** |  | **Backlog** | **Backlog** |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — | — | — |
| **Total** |  |  | <strong>~ 1497.51 ms</strong> (n=37/40) |  |  |
| ↳ /C |  |  | <strong>170.61x</strong> |  |  |
| ↳ μ/C |  |  | <strong>14323206.63x</strong> |  |  |
| ↳ Q₂/C |  |  | <strong>3109.50x</strong> |  |  |
| ↳ G/C |  |  | <strong>7026.62x</strong> |  |  |
| ↳ /JS |  |  | <strong>13.05x</strong> |  |  |
| ↳ μ/JS |  |  | <strong>26.59x</strong> |  |  |
| ↳ Q₂/JS |  |  | <strong>10.43x</strong> |  |  |
| ↳ G/JS |  |  | <strong>9.83x</strong> |  |  |


#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.074 μs | ~ 0.076 μs | ~ 0.069 μs |
| ↳ /C | 0.77x | 0.79x | 0.72x |
| ↳ /JS | 0.29x | 0.29x | 0.27x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs | ~ 0.15 μs | ~ 0.14 μs |
| ↳ /C | 1.86x | 1.86x | 1.82x |
| ↳ /JS | 0.49x | 0.49x | 0.48x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 6.35 μs | ~ 3.76 μs | ~ 2.84 μs |
| ↳ /C | 5218.38x | 3088.15x | 2329.60x |
| ↳ /JS | 0.74x | 0.44x | 0.33x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 277.12 μs | ~ 322.34 μs | ~ 329.88 μs |
| ↳ /C | 8.62x | 10.02x | 10.26x |
| ↳ /JS | 1.48x | 1.72x | 1.76x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 228.44 μs | ~ 48.07 μs | ~ 43.57 μs |
| ↳ /C | 69.85x | 14.70x | 13.32x |
| ↳ /JS | 3.07x | 0.65x | 0.59x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.99 μs | ~ 9.07 μs | ~ 9.14 μs |
| ↳ /C | 1.12x | 0.60x | 0.60x |
| ↳ /JS | 0.25x | 0.14x | 0.14x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 237.12 μs | ~ 235.53 μs | ~ 62.26 μs |
| ↳ /C | 212289.09x | 210862.26x | 55737.08x |
| ↳ /JS | 0.19x | 0.19x | 0.049x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 37.49 μs | ~ 22.68 μs | ~ 0.80 μs |
| ↳ /C | 328.53x | 198.72x | 7.00x |
| ↳ /JS | 0.65x | 0.39x | 0.014x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 22551.00 μs | ~ 15428.00 μs | ~ 11970.00 μs |
| ↳ /C | 2.59x | 1.77x | 1.37x |
| ↳ /JS | 0.27x | 0.18x | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 17812.00 μs | ~ 27373.00 μs | ~ 6730.50 μs |
| ↳ /C | 15821455.25x | 24313984.65x | 5978346.32x |
| ↳ /JS | 1.50x | 2.31x | 0.57x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 1.70 μs | ~ 2.02 μs | ~ 0.84 μs |
| ↳ /C | 1519.23x | 1808.21x | 750.57x |
| ↳ /JS | 0.036x | 0.043x | 0.018x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 2989.50 μs | ~ 1964.00 μs | ~ 0.61 μs |
| ↳ /C | 2666001.99x | 1751472.79x | 543.76x |
| ↳ /JS | 0.21x | 0.14x | 0.000043x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.29 μs | ~ 5.04 μs | ~ 2.79 μs |
| ↳ /C | 5158.17x | 3566.47x | 1975.98x |
| ↳ /JS | 1.03x | 0.71x | 0.39x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs | ~ 0.015 μs | ~ 0.0010 μs |
| ↳ /C | 0.89x | 13.02x | 0.89x |
| ↳ /JS | 0.028x | 0.41x | 0.028x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| **Total** | <strong>~ 44.17 ms</strong> (n=14/40) | <strong>~ 45.41 ms</strong> (n=14/40) | <strong>~ 19.15 ms</strong> (n=14/40) |
| ↳ /C | <strong>5.04x</strong> | <strong>5.18x</strong> | <strong>2.19x</strong> |
| ↳ μ/C | <strong>1336575.45x</strong> | <strong>1877501.71x</strong> | <strong>431408.52x</strong> |
| ↳ Q₂/C | <strong>199.19x</strong> | <strong>106.71x</strong> | <strong>11.79x</strong> |
| ↳ G/C | <strong>339.28x</strong> | <strong>317.58x</strong> | <strong>81.59x</strong> |
| ↳ /JS | <strong>0.39x</strong> | <strong>0.41x</strong> | <strong>0.17x</strong> |
| ↳ μ/JS | <strong>0.73x</strong> | <strong>0.58x</strong> | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>0.39x</strong> | <strong>0.40x</strong> | <strong>0.20x</strong> |
| ↳ G/JS | <strong>0.38x</strong> | <strong>0.36x</strong> | <strong>0.091x</strong> |


#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.11 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.15x | 1.19x | 1.05x |
| ↳ /JS | 0.43x | 0.44x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.21 μs | ~ 0.23 μs | ~ 0.24 μs |
| ↳ /C | 2.64x | 2.92x | 3.06x |
| ↳ /JS | 0.69x | 0.76x | 0.80x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 37.32 μs | ~ 5.22 μs | ~ 0.86 μs |
| ↳ /C | 30658.76x | 4289.92x | 705.63x |
| ↳ /JS | 4.32x | 0.60x | 0.099x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 1341.30 μs | ~ 103.95 μs | ~ 104.62 μs |
| ↳ /C | 41.70x | 3.23x | 3.25x |
| ↳ /JS | 7.17x | 0.56x | 0.56x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 404.77 μs | ~ 1276.40 μs | ~ 15.46 μs |
| ↳ /C | 123.76x | 390.27x | 4.73x |
| ↳ /JS | 5.44x | 17.17x | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.24 μs | ~ 18.81 μs | ~ 19.68 μs |
| ↳ /C | 1.27x | 1.24x | 1.30x |
| ↳ /JS | 0.29x | 0.28x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 520.08 μs | ~ 418.19 μs | ~ 0.012 μs |
| ↳ /C | 465608.66x | 374391.03x | 10.66x |
| ↳ /JS | 0.41x | 0.33x | 0.0000094x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 148.04 μs | ~ 30.95 μs | ~ 36.68 μs |
| ↳ /C | 1297.10x | 271.21x | 321.43x |
| ↳ /JS | 2.55x | 0.53x | 0.63x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16911.71 μs | ~ 14282.54 μs | ~ 14426.50 μs |
| ↳ /C | 1.94x | 1.64x | 1.66x |
| ↳ /JS | 0.20x | 0.17x | 0.17x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 66666.38 μs | ~ 52475.29 μs | ~ 0.0097 μs |
| ↳ /C | 59216206.42x | 46611018.88x | 8.64x |
| ↳ /JS | 5.63x | 4.43x | 0.00000082x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 20.41 μs | ~ 14.55 μs | ~ 7.61 μs |
| ↳ /C | 18232.84x | 12995.16x | 6801.78x |
| ↳ /JS | 0.43x | 0.31x | 0.16x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8739.06 μs | ~ 6780.02 μs | ~ 0.010 μs |
| ↳ /C | 7793396.22x | 6046345.36x | 8.94x |
| ↳ /JS | 0.62x | 0.48x | 0.00000071x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 27.82 μs | ~ 16.21 μs | ~ 1.04 μs |
| ↳ /C | 19683.86x | 11465.03x | 738.33x |
| ↳ /JS | 3.93x | 2.29x | 0.15x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.040 μs | ~ 0.044 μs | ~ 0.0095 μs |
| ↳ /C | 35.56x | 39.35x | 8.51x |
| ↳ /JS | 1.12x | 1.24x | 0.27x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |
| **Total** | <strong>~ 94.84 ms</strong> (n=14/40) | <strong>~ 75.42 ms</strong> (n=14/40) | <strong>~ 14.61 ms</strong> (n=14/40) |
| ↳ /C | <strong>10.83x</strong> | <strong>8.61x</strong> | <strong>1.67x</strong> |
| ↳ μ/C | <strong>4824663.71x</strong> | <strong>3790086.89x</strong> | <strong>615.64x</strong> |
| ↳ Q₂/C | <strong>710.43x</strong> | <strong>330.74x</strong> | <strong>8.57x</strong> |
| ↳ G/C | <strong>1108.31x</strong> | <strong>697.24x</strong> | <strong>19.07x</strong> |
| ↳ /JS | <strong>0.85x</strong> | <strong>0.67x</strong> | <strong>0.13x</strong> |
| ↳ μ/JS | <strong>2.37x</strong> | <strong>2.11x</strong> | <strong>0.27x</strong> |
| ↳ Q₂/JS | <strong>0.91x</strong> | <strong>0.54x</strong> | <strong>0.19x</strong> |
| ↳ G/JS | <strong>1.24x</strong> | <strong>0.78x</strong> | <strong>0.021x</strong> |


#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
|  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
|  | **WIP** | **WIP** | **WIP** |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.13 μs | ~ 3.19 μs | ~ 2.11 μs |
| ↳ /C | 22.21x | 33.26x | 22.00x |
| ↳ /JS | 8.27x | 12.39x | 8.20x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 1.08 μs | ~ 2.26 μs | ~ 1.27 μs |
| ↳ /C | 13.62x | 28.51x | 16.02x |
| ↳ /JS | 3.56x | 7.45x | 4.18x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 137.53 μs | ~ 174.47 μs | ~ 0.63 μs |
| ↳ /C | 112984.56x | 143331.75x | 517.56x |
| ↳ /JS | 15.93x | 20.20x | 0.073x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 72.18 μs | ~ 49.87 μs | ~ 50.30 μs |
| ↳ /C | 2.24x | 1.55x | 1.56x |
| ↳ /JS | 0.39x | 0.27x | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1413.71 μs | ~ 1676.49 μs | ~ 491.74 μs |
| ↳ /C | 432.26x | 512.61x | 150.36x |
| ↳ /JS | 19.02x | 22.55x | 6.61x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 45.89 μs | ~ 100.16 μs | ~ 52.48 μs |
| ↳ /C | 3.03x | 6.61x | 3.46x |
| ↳ /JS | 0.69x | 1.50x | 0.79x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1941.41 μs | ~ 6824.38 μs | ~ 32.84 μs |
| ↳ /C | 1738071.30x | 6109610.55x | 29400.42x |
| ↳ /JS | 1.53x | 5.37x | 0.026x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 698.16 μs | ~ 1022.38 μs | ~ 1.89 μs |
| ↳ /C | 6117.34x | 8958.18x | 16.56x |
| ↳ /JS | 12.04x | 17.62x | 0.033x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 53877.92 μs | ~ 509843.88 μs | ~ 113830.46 μs |
| ↳ /C | 6.19x | 58.55x | 13.07x |
| ↳ /JS | 0.64x | 6.05x | 1.35x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 6237.63 μs | ~ 432965.96 μs | ~ 87761.96 μs |
| ↳ /C | 5540556.02x | 384580707.45x | 77954296.14x |
| ↳ /JS | 0.53x | 36.58x | 7.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 467.12 μs | ~ 560.06 μs | ~ 0.48 μs |
| ↳ /C | 417281.93x | 500305.96x | 428.79x |
| ↳ /JS | 9.87x | 11.83x | 0.010x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 791.27 μs | ~ 111183.96 μs | ~ 326.09 μs |
| ↳ /C | 705645.56x | 99152586.81x | 290803.34x |
| ↳ /JS | 0.056x | 7.91x | 0.023x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 89.94 μs | ~ 26.94 μs | ~ 0.80 μs |
| ↳ /C | 63625.51x | 19057.94x | 565.94x |
| ↳ /JS | 12.72x | 3.81x | 0.11x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.84 μs | ~ 0.85 μs | ~ 0.080 μs |
| ↳ /C | 750.42x | 759.35x | 71.47x |
| ↳ /JS | 23.69x | 23.98x | 2.26x |
| **Total WIP** | <strong>65.78 ms</strong> | <strong>1064.43 ms</strong> | <strong>202.55 ms</strong> |
| ↳ /C | <strong>7.51x</strong> | <strong>121.53x</strong> | <strong>23.13x</strong> |
| ↳ μ/C | <strong>613250.87x</strong> | <strong>35036854.22x</strong> | <strong>5591164.76x</strong> |
| ↳ Q₂/C | <strong>3433.88x</strong> | <strong>4858.77x</strong> | <strong>110.91x</strong> |
| ↳ G/C | <strong>2558.31x</strong> | <strong>6964.40x</strong> | <strong>347.31x</strong> |
| ↳ /JS | <strong>0.59x</strong> | <strong>9.51x</strong> | <strong>1.81x</strong> |
| ↳ μ/JS | <strong>7.78x</strong> | <strong>12.68x</strong> | <strong>2.24x</strong> |
| ↳ Q₂/JS | <strong>5.92x</strong> | <strong>9.87x</strong> | <strong>0.53x</strong> |
| ↳ G/JS | <strong>2.86x</strong> | <strong>7.79x</strong> | <strong>0.39x</strong> |
|  | **Backlog** | **Backlog** | **Backlog** |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |


#### Rust

Rust Benchmark          | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F# (adapted)<br>↓<br>[Fable (patched)](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written F#<br>↓<br>[Fable](https://github.com/fable-compiler/fable)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written FP-style Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO | Hand-written imperative Rust FFI<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO |
----------------------- | ------------------------- | -------------------------- | ------------------------------ | ------------------------ | ------------------------------ |
|  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
|  | **WIP** | **WIP** |  | **WIP** | **WIP** |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.24 μs | ~ 3.68 μs | ~ 0.23 μs | ~ 0.22 μs | ~ 0.23 μs |
| ↳ /C | 2.50x | 38.37x | 2.40x | 2.29x | 2.40x |
| ↳ /JS | 0.93x | 14.30x | 0.89x | 0.85x | 0.89x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.090 μs | ~ 11.78 μs | ~ 0.090 μs | ~ 0.12 μs | ~ 0.10 μs |
| ↳ /C | 1.14x | 148.61x | 1.14x | 1.51x | 1.26x |
| ↳ /JS | 0.30x | 38.82x | 0.30x | 0.40x | 0.33x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.11 μs | ~ 459.33 μs | ~ 17.07 μs | ~ 11.05 μs | ~ 0.51 μs |
| ↳ /C | 90.37x | 377351.83x | 14023.46x | 9077.87x | 418.98x |
| ↳ /JS | 0.013x | 53.19x | 1.98x | 1.28x | 0.059x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 29.88 μs | ~ 23547.50 μs | ~ 10.93 μs | ~ 34.70 μs | ~ 34.63 μs |
| ↳ /C | 0.93x | 732.11x | 0.34x | 1.08x | 1.08x |
| ↳ /JS | 0.16x | 125.86x | 0.058x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.14 μs | ~ 8728.98 μs | ~ 114.96 μs | ~ 107.97 μs | ~ 3.47 μs |
| ↳ /C | 0.96x | 2669.00x | 35.15x | 33.01x | 1.06x |
| ↳ /JS | 0.042x | 117.42x | 1.55x | 1.45x | 0.047x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.45 μs | ~ 625.00 μs | ~ 15.60 μs | ~ 18.59 μs | ~ 18.36 μs |
| ↳ /C | 1.09x | 41.23x | 1.03x | 1.23x | 1.21x |
| ↳ /JS | 0.25x | 9.36x | 0.23x | 0.28x | 0.27x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 159.52 μs | ~ 11792.17 μs | ~ 2249.51 μs | ~ 1033.78 μs | ~ 0.010 μs |
| ↳ /C | 142812.25x | 10557085.96x | 2013901.63x | 925504.32x | 8.95x |
| ↳ /JS | 0.13x | 9.29x | 1.77x | 0.81x | 0.0000079x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 144.00 μs | ~ 2194.96 μs | ~ 120.43 μs | ~ 78.88 μs | ~ 0.78 μs |
| ↳ /C | 1261.74x | 19232.42x | 1055.22x | 691.15x | 6.83x |
| ↳ /JS | 2.48x | 37.84x | 2.08x | 1.36x | 0.013x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8136.98 μs | ~ 659207.04 μs | ~ 50862.42 μs | ~ 28737.38 μs | ~ 28247.96 μs |
| ↳ /C | 0.93x | 75.70x | 5.84x | 3.30x | 3.24x |
| ↳ /JS | 0.097x | 7.82x | 0.60x | 0.34x | 0.34x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.030 μs | ~ 1588595.46 μs | ~ 7471.33 μs | ~ 0.010 μs | ~ 0.010 μs |
| ↳ /C | 26.65x | 1411065123.60x | 6636386.33x | 8.88x | 8.88x |
| ↳ /JS | 0.0000025x | 134.21x | 0.63x | 0.00000084x | 0.00000084x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 44.57 μs | ~ 826.12 μs | ~ 73.12 μs | ~ 27.36 μs | ~ 0.010 μs |
| ↳ /C | 39814.73x | 737979.43x | 65318.67x | 24440.90x | 8.93x |
| ↳ /JS | 0.94x | 17.45x | 1.54x | 0.58x | 0.00021x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.010 μs | ~ 159056.38 μs | ~ 38342.83 μs | ~ 23675.83 μs | ~ 0.010 μs |
| ↳ /C | 8.92x | 141844664.70x | 34193698.27x | 21113835.03x | 8.92x |
| ↳ /JS | 0.00000071x | 11.31x | 2.73x | 1.68x | 0.00000071x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.030 μs | ~ 298.70 μs | ~ 3.28 μs | ~ 0.58 μs | ~ 0.25 μs |
| ↳ /C | 21.22x | 211306.86x | 2320.34x | 410.30x | 176.86x |
| ↳ /JS | 0.0042x | 42.24x | 0.46x | 0.082x | 0.035x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.070 μs | ~ 0.95 μs | ~ 0.10 μs | ~ 0.020 μs | ~ 0.010 μs |
| ↳ /C | 62.53x | 848.69x | 89.34x | 17.87x | 8.93x |
| ↳ /JS | 1.97x | 26.80x | 2.82x | 0.56x | 0.28x |
| **Total WIP** | <strong>8.54 ms</strong> | <strong>2455.35 ms</strong> |  | <strong>53.73 ms</strong> | <strong>28.31 ms</strong> |
| ↳ /C | <strong>0.97x</strong> | <strong>280.33x</strong> |  | <strong>6.13x</strong> | <strong>3.23x</strong> |
| ↳ μ/C | <strong>13150.43x</strong> | <strong>111772664.18x</strong> |  | <strong>1576716.34x</strong> | <strong>46.97x</strong> |
| ↳ Q₂/C | <strong>15.07x</strong> | <strong>10950.71x</strong> |  | <strong>25.44x</strong> | <strong>7.86x</strong> |
| ↳ G/C | <strong>30.20x</strong> | <strong>26167.52x</strong> |  | <strong>192.74x</strong> | <strong>6.73x</strong> |
| ↳ /JS | <strong>0.076x</strong> | <strong>21.95x</strong> |  | <strong>0.48x</strong> | <strong>0.25x</strong> |
| ↳ μ/JS | <strong>0.52x</strong> | <strong>46.13x</strong> |  | <strong>0.70x</strong> | <strong>0.18x</strong> |
| ↳ Q₂/JS | <strong>0.14x</strong> | <strong>32.32x</strong> |  | <strong>0.57x</strong> | <strong>0.053x</strong> |
| ↳ G/JS | <strong>0.034x</strong> | <strong>29.26x</strong> |  | <strong>0.22x</strong> | <strong>0.0075x</strong> |
|  | **Backlog** | **Backlog** |  | **Backlog** | **Backlog** |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — | — | — |
| **Total** |  |  | <strong>99.28 ms</strong> (n=14/40) |  |  |
| ↳ /C |  |  | <strong>11.33x</strong> |  |  |
| ↳ μ/C |  |  | <strong>3066202.80x</strong> |  |  |
| ↳ Q₂/C |  |  | <strong>572.28x</strong> |  |  |
| ↳ G/C |  |  | <strong>750.25x</strong> |  |  |
| ↳ /JS |  |  | <strong>0.89x</strong> |  |  |
| ↳ μ/JS |  |  | <strong>1.26x</strong> |  |  |
| ↳ Q₂/JS |  |  | <strong>1.22x</strong> |  |  |
| ↳ G/JS |  |  | <strong>0.84x</strong> |  |  |


#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release | Hand-written F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
|  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
|  | **WIP** | **WIP** | **WIP** |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 1.12 μs | ~ 0.50 μs | ~ 0.50 μs |
| ↳ /C | 11.68x | 5.21x | 5.21x |
| ↳ /JS | 4.35x | 1.94x | 1.94x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.36 μs | ~ 0.12 μs | ~ 0.11 μs |
| ↳ /C | 4.54x | 1.51x | 1.39x |
| ↳ /JS | 1.19x | 0.40x | 0.36x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 560.47 μs | ~ 16.05 μs | ~ 3.05 μs |
| ↳ /C | 460441.03x | 13185.50x | 2505.66x |
| ↳ /JS | 64.90x | 1.86x | 0.35x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.45 μs | ~ 37.60 μs | ~ 49.49 μs |
| ↳ /C | 1.79x | 1.17x | 1.54x |
| ↳ /JS | 0.31x | 0.20x | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 2545.27 μs | ~ 87.33 μs | ~ 7.21 μs |
| ↳ /C | 778.25x | 26.70x | 2.20x |
| ↳ /JS | 34.24x | 1.17x | 0.097x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.15 μs | ~ 7.70 μs | ~ 7.66 μs |
| ↳ /C | 1.00x | 0.51x | 0.51x |
| ↳ /JS | 0.23x | 0.12x | 0.11x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1678.49 μs | ~ 342.79 μs | ~ 25.24 μs |
| ↳ /C | 1502688.92x | 306886.98x | 22596.42x |
| ↳ /JS | 1.32x | 0.27x | 0.020x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 525.55 μs | ~ 45.26 μs | ~ 1.80 μs |
| ↳ /C | 4604.91x | 396.57x | 15.77x |
| ↳ /JS | 9.06x | 0.78x | 0.031x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 19071.54 μs | ~ 36659.92 μs | ~ 9499.96 μs |
| ↳ /C | 2.19x | 4.21x | 1.09x |
| ↳ /JS | 0.23x | 0.44x | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2523.93 μs | ~ 23242.04 μs | ~ 2254.54 μs |
| ↳ /C | 2241873.21x | 20644671.89x | 2002588.35x |
| ↳ /JS | 0.21x | 1.96x | 0.19x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 169.78 μs | ~ 78.22 μs | ~ 3.21 μs |
| ↳ /C | 151665.80x | 69874.53x | 2867.52x |
| ↳ /JS | 3.59x | 1.65x | 0.068x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 6576.73 μs | ~ 9665.58 μs | ~ 235.04 μs |
| ↳ /C | 5865052.77x | 8619653.95x | 209605.99x |
| ↳ /JS | 0.47x | 0.69x | 0.017x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 62.07 μs | ~ 29.54 μs | ~ 15.63 μs |
| ↳ /C | 43909.67x | 20897.24x | 11057.00x |
| ↳ /JS | 8.78x | 4.18x | 2.21x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.24 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 214.41x | 53.60x | 17.87x |
| ↳ /JS | 6.77x | 1.69x | 0.56x |
| **Total WIP** | <strong>33.79 ms</strong> | <strong>70.21 ms</strong> | <strong>12.10 ms</strong> |
| ↳ /C | <strong>3.86x</strong> | <strong>8.02x</strong> | <strong>1.38x</strong> |
| ↳ μ/C | <strong>733660.72x</strong> | <strong>2119689.97x</strong> | <strong>160804.75x</strong> |
| ↳ Q₂/C | <strong>2691.58x</strong> | <strong>225.09x</strong> | <strong>16.82x</strong> |
| ↳ G/C | <strong>1925.83x</strong> | <strong>728.13x</strong> | <strong>149.83x</strong> |
| ↳ /JS | <strong>0.30x</strong> | <strong>0.63x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>9.69x</strong> | <strong>1.24x</strong> | <strong>0.45x</strong> |
| ↳ Q₂/JS | <strong>2.45x</strong> | <strong>0.98x</strong> | <strong>0.15x</strong> |
| ↳ G/JS | <strong>2.15x</strong> | <strong>0.81x</strong> | <strong>0.17x</strong> |
|  | **Backlog** | **Backlog** | **Backlog** |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |


#### Java

Java Benchmark          | Hand-written PureScript<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT | Hand-written FP-style Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT | Hand-written imperative Java FFI<br>↓<br>[javapurs](https://github.com/0x000000000000000000001/javapurs)<br>↓<br>Java<br>↓<br>HotSpot JIT |
----------------------- | ------------------------- | ------------------------ | ------------------------------ |
|  | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
|  | **WIP** | **WIP** | **WIP** |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.14 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.50x | 1.11x | 1.06x |
| ↳ /JS | 0.56x | 0.41x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.26 μs | ~ 0.099 μs | ~ 0.077 μs |
| ↳ /C | 3.27x | 1.25x | 0.97x |
| ↳ /JS | 0.85x | 0.33x | 0.25x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 5.03 μs | ~ 4.18 μs | ~ 0.57 μs |
| ↳ /C | 4132.44x | 3436.55x | 465.90x |
| ↳ /JS | 0.58x | 0.48x | 0.066x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 45.64 μs | ~ 38.54 μs | ~ 39.83 μs |
| ↳ /C | 1.42x | 1.20x | 1.24x |
| ↳ /JS | 0.24x | 0.21x | 0.21x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 73.28 μs | ~ 38.80 μs | ~ 4.18 μs |
| ↳ /C | 22.41x | 11.86x | 1.28x |
| ↳ /JS | 0.99x | 0.52x | 0.056x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 18.18 μs | ~ 7.60 μs | ~ 5.80 μs |
| ↳ /C | 1.20x | 0.50x | 0.38x |
| ↳ /JS | 0.27x | 0.11x | 0.087x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 310.51 μs | ~ 375.72 μs | ~ 0.0031 μs |
| ↳ /C | 277987.13x | 336364.53x | 2.79x |
| ↳ /JS | 0.24x | 0.30x | 0.0000025x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.29 μs | ~ 18.40 μs | ~ 0.54 μs |
| ↳ /C | 186.53x | 161.22x | 4.75x |
| ↳ /JS | 0.37x | 0.32x | 0.0093x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16788.38 μs | ~ 12065.75 μs | ~ 11828.17 μs |
| ↳ /C | 1.93x | 1.39x | 1.36x |
| ↳ /JS | 0.20x | 0.14x | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0054 μs | ~ 14583.08 μs | ~ 0.0034 μs |
| ↳ /C | 4.81x | 12953379.47x | 2.99x |
| ↳ /JS | 0.00000046x | 1.23x | 0.00000028x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 6.26 μs | ~ 17.25 μs | ~ 0.0033 μs |
| ↳ /C | 5589.13x | 15408.69x | 2.91x |
| ↳ /JS | 0.13x | 0.36x | 0.000069x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5.24 μs | ~ 8062.00 μs | ~ 0.037 μs |
| ↳ /C | 4675.16x | 7189599.60x | 32.91x |
| ↳ /JS | 0.00037x | 0.57x | 0.0000026x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.96 μs | ~ 3.06 μs | ~ 0.50 μs |
| ↳ /C | 2092.37x | 2167.00x | 351.90x |
| ↳ /JS | 0.42x | 0.43x | 0.070x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.013 μs | ~ 0.0062 μs | ~ 0.0023 μs |
| ↳ /C | 11.36x | 5.50x | 2.08x |
| ↳ /JS | 0.36x | 0.17x | 0.066x |
| **Total WIP** | <strong>~ 17.28 ms</strong> | <strong>~ 35.21 ms</strong> | <strong>~ 11.88 ms</strong> |
| ↳ /C | <strong>1.97x</strong> | <strong>4.02x</strong> | <strong>1.36x</strong> |
| ↳ μ/C | <strong>21050.76x</strong> | <strong>1464324.28x</strong> | <strong>62.32x</strong> |
| ↳ Q₂/C | <strong>16.88x</strong> | <strong>86.54x</strong> | <strong>2.44x</strong> |
| ↳ G/C | <strong>76.38x</strong> | <strong>297.93x</strong> | <strong>4.42x</strong> |
| ↳ /JS | <strong>0.15x</strong> | <strong>0.31x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>0.37x</strong> | <strong>0.40x</strong> | <strong>0.097x</strong> |
| ↳ Q₂/JS | <strong>0.32x</strong> | <strong>0.34x</strong> | <strong>0.066x</strong> |
| ↳ G/JS | <strong>0.085x</strong> | <strong>0.33x</strong> | <strong>0.0049x</strong> |
|  | **Backlog** | **Backlog** | **Backlog** |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — | — | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — | — | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — | — | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — | — | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — | — | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — | — | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — | — | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — | — | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — | — | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — | — | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — | — | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — | — | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — | — | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — | — | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — | — | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — | — | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — | — | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — | — | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — | — | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — | — | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — | — | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — | — | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — | — | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — | — | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — | — | — |


#### Koka

Koka Benchmark          | Hand-written Koka<br>↓<br>koka -O3 |
----------------------- | ----- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.22 μs |
| ↳ /C | 2.28x |
| ↳ /JS | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.56x |
| ↳ /JS | 0.41x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 4.28 μs |
| ↳ /C | 3512.35x |
| ↳ /JS | 0.50x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 80.51 μs |
| ↳ /C | 2.50x |
| ↳ /JS | 0.43x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 158.94 μs |
| ↳ /C | 48.60x |
| ↳ /JS | 2.14x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.38 μs |
| ↳ /C | 1.28x |
| ↳ /JS | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 758.69 μs |
| ↳ /C | 679224.36x |
| ↳ /JS | 0.60x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 12.17 μs |
| ↳ /C | 106.62x |
| ↳ /JS | 0.21x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8746.50 μs |
| ↳ /C | 1.00x |
| ↳ /JS | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 35687.00 μs |
| ↳ /C | 31698870.06x |
| ↳ /JS | 3.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 5.21 μs |
| ↳ /C | 4650.61x |
| ↳ /JS | 0.11x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8982.00 μs |
| ↳ /C | 8010045.11x |
| ↳ /JS | 0.64x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 6.58 μs |
| ↳ /C | 4653.85x |
| ↳ /JS | 0.93x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs |
| ↳ /C | 0.92x |
| ↳ /JS | 0.029x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| **Total** | <strong>~ 54.46 ms</strong> (n=14/40) |
| ↳ /C | <strong>6.22x</strong> |
| ↳ μ/C | <strong>2885794.36x</strong> |
| ↳ Q₂/C | <strong>77.61x</strong> |
| ↳ G/C | <strong>363.96x</strong> |
| ↳ /JS | <strong>0.49x</strong> |
| ↳ μ/JS | <strong>0.73x</strong> |
| ↳ Q₂/JS | <strong>0.46x</strong> |
| ↳ G/JS | <strong>0.41x</strong> |


#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2 |
----------------------- | ------------------------ |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.056 μs |
| ↳ /C | 0.58x |
| ↳ /JS | 0.22x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.19 μs |
| ↳ /C | 2.45x |
| ↳ /JS | 0.64x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 3.58 μs |
| ↳ /C | 2940.33x |
| ↳ /JS | 0.41x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 62.84 μs |
| ↳ /C | 1.95x |
| ↳ /JS | 0.34x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.68 μs |
| ↳ /C | 2.35x |
| ↳ /JS | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 6.75 μs |
| ↳ /C | 0.45x |
| ↳ /JS | 0.10x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 143.90 μs |
| ↳ /C | 128826.86x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.58 μs |
| ↳ /C | 189.10x |
| ↳ /JS | 0.37x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 11559.00 μs |
| ↳ /C | 1.33x |
| ↳ /JS | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 4883.00 μs |
| ↳ /C | 4337310.02x |
| ↳ /JS | 0.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.016 μs |
| ↳ /C | 13.96x |
| ↳ /JS | 0.00033x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.46 μs |
| ↳ /C | 406.73x |
| ↳ /JS | 0.000032x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 4.33 μs |
| ↳ /C | 3060.77x |
| ↳ /JS | 0.61x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0041 μs |
| ↳ /C | 3.64x |
| ↳ /JS | 0.11x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| **Total** | <strong>~ 16.69 ms</strong> (n=14/40) |
| ↳ /C | <strong>1.91x</strong> |
| ↳ μ/C | <strong>319482.89x</strong> |
| ↳ Q₂/C | <strong>8.80x</strong> |
| ↳ G/C | <strong>70.56x</strong> |
| ↳ /JS | <strong>0.15x</strong> |
| ↳ μ/JS | <strong>0.26x</strong> |
| ↳ Q₂/JS | <strong>0.18x</strong> |
| ↳ G/JS | <strong>0.079x</strong> |


#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3 |
----------------------- | --------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.049 μs |
| ↳ /C | 0.52x |
| ↳ /JS | 0.19x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.56x |
| ↳ /JS | 0.41x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 1.98 μs |
| ↳ /C | 1623.80x |
| ↳ /JS | 0.23x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 49.82 μs |
| ↳ /C | 1.55x |
| ↳ /JS | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 15.85 μs |
| ↳ /C | 4.85x |
| ↳ /JS | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.05 μs |
| ↳ /C | 1.06x |
| ↳ /JS | 0.24x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 140.38 μs |
| ↳ /C | 125672.45x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 17.94 μs |
| ↳ /C | 157.17x |
| ↳ /JS | 0.31x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10787.00 μs |
| ↳ /C | 1.24x |
| ↳ /JS | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 12448.00 μs |
| ↳ /C | 11056898.44x |
| ↳ /JS | 1.05x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 10.03 μs |
| ↳ /C | 8956.63x |
| ↳ /JS | 0.21x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5690.00 μs |
| ↳ /C | 5074277.07x |
| ↳ /JS | 0.40x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 5.89 μs |
| ↳ /C | 4168.88x |
| ↳ /JS | 0.83x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0047 μs |
| ↳ /C | 4.18x |
| ↳ /JS | 0.13x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| **Total** | <strong>~ 29.18 ms</strong> (n=14/40) |
| ↳ /C | <strong>3.33x</strong> |
| ↳ μ/C | <strong>1162269.24x</strong> |
| ↳ Q₂/C | <strong>81.01x</strong> |
| ↳ G/C | <strong>241.02x</strong> |
| ↳ /JS | <strong>0.26x</strong> |
| ↳ μ/JS | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>0.23x</strong> |
| ↳ G/JS | <strong>0.27x</strong> |


#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3 |
----------------------- | ------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.71 μs |
| ↳ /C | 28.27x |
| ↳ /JS | 10.53x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs |
| ↳ /C | 3.97x |
| ↳ /JS | 1.04x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 154.89 μs |
| ↳ /C | 127246.77x |
| ↳ /JS | 17.94x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 2844.59 μs |
| ↳ /C | 88.44x |
| ↳ /JS | 15.20x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1453.15 μs |
| ↳ /C | 444.32x |
| ↳ /JS | 19.55x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 221.50 μs |
| ↳ /C | 14.61x |
| ↳ /JS | 3.32x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 2630.26 μs |
| ↳ /C | 2354773.00x |
| ↳ /JS | 2.07x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 974.45 μs |
| ↳ /C | 8538.24x |
| ↳ /JS | 16.80x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 613388.67 μs |
| ↳ /C | 70.44x |
| ↳ /JS | 7.28x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 265261.00 μs |
| ↳ /C | 235617282.80x |
| ↳ /JS | 22.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 293.30 μs |
| ↳ /C | 262004.38x |
| ↳ /JS | 6.20x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 36615.58 μs |
| ↳ /C | 32653359.99x |
| ↳ /JS | 2.60x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 36.57 μs |
| ↳ /C | 25872.82x |
| ↳ /JS | 5.17x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.48 μs |
| ↳ /C | 430.91x |
| ↳ /JS | 13.61x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | — |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | — |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | — |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | — |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | — |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | — |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | — |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | — |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | — |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | — |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | — |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | — |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | — |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | — |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | — |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | — |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | — |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | — |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | — |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | — |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | — |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | — |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | — |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | — |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | — |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | — |
| **Total** | <strong>~ 923.88 ms</strong> (n=14/40) |
| ↳ /C | <strong>105.48x</strong> |
| ↳ μ/C | <strong>19360725.64x</strong> |
| ↳ Q₂/C | <strong>4491.28x</strong> |
| ↳ G/C | <strong>6538.89x</strong> |
| ↳ /JS | <strong>8.26x</strong> |
| ↳ μ/JS | <strong>10.27x</strong> |
| ↳ Q₂/JS | <strong>8.91x</strong> |
| ↳ G/JS | <strong>7.31x</strong> |


#### C (reference)

The 23 [library references](tmp/c-library) complement the existing core reference. [Measurements and validation](docs/benchmark-results/2026-11-26-c-library.md).

C Benchmark             | [Hand-written imperative C](tmp/bench_native.c)<br>↓<br>clang -O3 |
----------------------- | -------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.096 μs |
| ↳ /C | — |
| ↳ /JS | 0.37x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.079 μs |
| ↳ /C | — |
| ↳ /JS | 0.26x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.00014x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 32.16 μs |
| ↳ /C | — |
| ↳ /JS | 0.17x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.27 μs |
| ↳ /C | — |
| ↳ /JS | 0.044x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.16 μs |
| ↳ /C | — |
| ↳ /JS | 0.23x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.00000088x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 0.11 μs |
| ↳ /C | — |
| ↳ /JS | 0.0020x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8708.00 μs |
| ↳ /C | — |
| ↳ /JS | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.000000095x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.000024x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.000000080x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.0014 μs |
| ↳ /C | — |
| ↳ /JS | 0.00020x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.032x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 0.43 μs |
| ↳ /C | — |
| ↳ /JS | 0.0035x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 4.55 μs |
| ↳ /C | — |
| ↳ /JS | 0.036x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 0.28 μs |
| ↳ /C | — |
| ↳ /JS | 0.0040x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 4.66 μs |
| ↳ /C | — |
| ↳ /JS | 0.0083x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 0.63 μs |
| ↳ /C | — |
| ↳ /JS | 0.012x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 0.084 μs |
| ↳ /C | — |
| ↳ /JS | 0.0028x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 0.042 μs |
| ↳ /C | — |
| ↳ /JS | 0.00052x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 0.0013 μs |
| ↳ /C | — |
| ↳ /JS | 0.000020x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 0.33 μs |
| ↳ /C | — |
| ↳ /JS | 0.0048x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.11x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 1.10 μs |
| ↳ /C | — |
| ↳ /JS | 0.029x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 1.08 μs |
| ↳ /C | — |
| ↳ /JS | 5.64x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 1.24 μs |
| ↳ /C | — |
| ↳ /JS | 0.019x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 0.64 μs |
| ↳ /C | — |
| ↳ /JS | 0.014x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 2.47 μs |
| ↳ /C | — |
| ↳ /JS | 0.017x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 0.92 μs |
| ↳ /C | — |
| ↳ /JS | 0.019x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.00046x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.00025x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 0.0014 μs |
| ↳ /C | — |
| ↳ /JS | 0.00019x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.0000024x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.0000022x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.0000072x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.0000053x |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 6589.00 μs |
| ↳ /C | — |
| ↳ /JS | 0.26x |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 10560.13 μs |
| ↳ /C | — |
| ↳ /JS | 0.13x |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 657.69 μs |
| ↳ /C | — |
| ↳ /JS | 0.071x |
| **Total** | <strong>26.58 ms</strong> |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ G/C | — |
| ↳ /JS | <strong>0.12x</strong> |
| ↳ μ/JS | <strong>0.19x</strong> |
| ↳ Q₂/JS | <strong>0.0065x</strong> |
| ↳ G/JS | <strong>0.0020x</strong> |


### Extended benchmark results (I/O, mutability, async)

**Extended tests (`srx/`)** cover features shared by a smaller subset of targets (`./bin/run --x`). For example, `Aff` is supported by the JavaScript, Rust and Go backends, but not the Scheme backend.

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO
----------------------- | ------------- | -------------- | --------------- | ---------------
| | | | **Stable**<br>Coming soon... | **Stable**<br>Coming soon... |
| | | | **WIP** | **WIP** |
[File I/O](srx/Test/FileOps/FileOps.purs) | ~ 466899.83 μs | ~ 477117.79 μs | ~ 475960.00 μs | ~ 542377.75 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.02x | 1.02x | 1.16x |
[STArray Operations](srx/Test/STArray.purs) | ~ 0.92 μs | ~ 0.54 μs | ~ 0.29 μs | ~ 0.29 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.59x | 0.32x | 0.32x |
[String Operations](srx/Test/StringOps.purs) | ~ 256.71 μs | ~ 232.50 μs | ~ 559.04 μs | ~ 414.58 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.91x | 2.18x | 1.61x |
[Aff Operations](srx/Test/AffOperations.purs) | ~ 10125.21 μs | ~ 11581.83 μs | ~ 10500.21 μs | ~ 10903.21 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.14x | 1.04x | 1.08x |
[Parallelism](srx/Test/Parallelism.purs) | ~ 14275052.04 μs | ~ 14278260.46 μs | ~ 965397.12 μs | ~ 478303.67 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.00x | 0.070x | 0.030x |
**Total Execution Time** | <strong>~ 14752.33 ms</strong> | <strong>~ 14767.19 ms</strong> | <strong>~ 1452.42 ms</strong> | <strong>~ 1032.00 ms</strong> |
↳ /C | — | — | — | — |
↳ μ/C | — | — | — | — |
↳ Q₂/C | — | — | — | — |
↳ G/C | — | — | — | — |
↳ /JS | — | <strong>1.00x</strong> | <strong>0.10x</strong> | <strong>0.070x</strong> |
↳ μ/JS | — | <strong>0.93x</strong> | <strong>0.92x</strong> | <strong>0.84x</strong> |
↳ Q₂/JS | — | <strong>1.00x</strong> | <strong>1.02x</strong> | <strong>1.08x</strong> |
↳ G/JS | — | <strong>0.91x</strong> | <strong>0.55x</strong> | <strong>0.46x</strong> |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with the Go/Rust multicore runtimes.

### Application benchmark results

App scenario | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build
--- | --- | --- |
| | | **Stable**<br>Coming soon... |
| | | **WIP** |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (corpus parsing) | ~ 62.17 ms | ~ 128.79 ms |
↳ /JS | — | 2.07x |
[CST](https://github.com/natefaubion/purescript-language-cst-parser) (fixture parsing) | ~ 8.94 ms | ~ 17.67 ms |
↳ /JS | — | 1.98x |
[tidy](https://github.com/natefaubion/purescript-tidy) (parsing + formatting) | ~ 14.23 ms | ~ 23.65 ms |
↳ /JS | — | 1.66x |
[tidy](https://github.com/natefaubion/purescript-tidy) (10 iterations) | ~ 316.12 ms | ~ 239.07 ms |
↳ /JS | — | 0.76x |
[b8x](https://github.com/0x000000000000000000001/b8x.pub) (287 tests) | ~ 10.43 s | ~ 11.05 s |
↳ /JS | — | 1.06x |
**Total** | <strong>~ 10.83 s</strong> | <strong>~ 11.46 s</strong> |
↳ /JS | — | <strong>1.06x</strong> |
↳ μ/JS | — | <strong>1.51x</strong> |
↳ Q₂/JS | — | <strong>1.66x</strong> |
↳ G/JS | — | <strong>1.40x</strong> |

### Compilation times...

#### ... to Go

Benchmark   | [gopurs](https://github.com/0x000000000000000000001/gopurs) JS binary<br>(official compiler)<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Go binary<br>with `--dev`<br>(compiled by [gopurs](https://github.com/0x000000000000000000001/gopurs))<br>(WIP) | [gopurs](https://github.com/0x000000000000000000001/gopurs) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ------------------ | --------------------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 146.39 s | ~ 78.04 s | ~ 53.61 s | ~ 52.92 s
↳ /JS | — | 0.53x | 0.37x | 0.36x
[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff) | ~ 16023.00 ms | ~ 6312.00 ms | ~ 5437.00 ms | ~ 5002.00 ms
↳ /JS | — | 0.39x | 0.34x | 0.31x
[gopurs-argonaut-core](https://github.com/0x000000000000000000001/gopurs-argonaut-core) | ~ 15672.00 ms | ~ 5930.00 ms | ~ 4962.00 ms | ~ 4800.00 ms
↳ /JS | — | 0.38x | 0.32x | 0.31x
[gopurs-arrays](https://github.com/0x000000000000000000001/gopurs-arrays) | ~ 12758.00 ms | ~ 5072.00 ms | ~ 3782.00 ms | ~ 4001.00 ms
↳ /JS | — | 0.40x | 0.30x | 0.31x
[gopurs-assert](https://github.com/0x000000000000000000001/gopurs-assert) | ~ 2514.00 ms | ~ 1305.00 ms | ~ 1187.00 ms | ~ 922.00 ms
↳ /JS | — | 0.52x | 0.47x | 0.37x
[gopurs-avar](https://github.com/0x000000000000000000001/gopurs-avar) | ~ 15097.00 ms | ~ 5865.00 ms | ~ 4971.00 ms | ~ 4770.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.32x
[gopurs-catenable-lists](https://github.com/0x000000000000000000001/gopurs-catenable-lists) | ~ 10188.00 ms | ~ 3918.00 ms | ~ 3387.00 ms | ~ 3024.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.30x
[gopurs-console](https://github.com/0x000000000000000000001/gopurs-console) | ~ 9476.00 ms | ~ 3207.00 ms | ~ 2880.00 ms | ~ 2513.00 ms
↳ /JS | — | 0.34x | 0.30x | 0.27x
[gopurs-datetime](https://github.com/0x000000000000000000001/gopurs-datetime) | ~ 14497.00 ms | ~ 5560.00 ms | ~ 4768.00 ms | ~ 4324.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.30x
[gopurs-effect](https://github.com/0x000000000000000000001/gopurs-effect) | ~ 4161.00 ms | ~ 1390.00 ms | ~ 1296.00 ms | ~ 994.00 ms
↳ /JS | — | 0.33x | 0.31x | 0.24x
[gopurs-enums](https://github.com/0x000000000000000000001/gopurs-enums) | ~ 10858.00 ms | ~ 3936.00 ms | ~ 3300.00 ms | ~ 3125.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.29x
[gopurs-exceptions](https://github.com/0x000000000000000000001/gopurs-exceptions) | ~ 4671.00 ms | ~ 1665.00 ms | ~ 1519.00 ms | ~ 1195.00 ms
↳ /JS | — | 0.36x | 0.33x | 0.26x
[gopurs-foldable-traversable](https://github.com/0x000000000000000000001/gopurs-foldable-traversable) | ~ 10499.00 ms | ~ 3841.00 ms | ~ 3357.00 ms | ~ 3057.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.29x
[gopurs-foreign](https://github.com/0x000000000000000000001/gopurs-foreign) | ~ 12846.00 ms | ~ 5309.00 ms | ~ 4621.00 ms | ~ 4347.00 ms
↳ /JS | — | 0.41x | 0.36x | 0.34x
[gopurs-foreign-object](https://github.com/0x000000000000000000001/gopurs-foreign-object) | ~ 15874.00 ms | ~ 6228.00 ms | ~ 5090.00 ms | ~ 5136.00 ms
↳ /JS | — | 0.39x | 0.32x | 0.32x
[gopurs-free](https://github.com/0x000000000000000000001/gopurs-free) | ~ 11960.00 ms | ~ 4840.00 ms | ~ 4252.00 ms | ~ 3881.00 ms
↳ /JS | — | 0.40x | 0.36x | 0.32x
[gopurs-functions](https://github.com/0x000000000000000000001/gopurs-functions) | ~ 4231.00 ms | ~ 1417.00 ms | ~ 1299.00 ms | ~ 960.00 ms
↳ /JS | — | 0.33x | 0.31x | 0.23x
[gopurs-integers](https://github.com/0x000000000000000000001/gopurs-integers) | ~ 10024.00 ms | ~ 3470.00 ms | ~ 3098.00 ms | ~ 2707.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-js-bigints](https://github.com/0x000000000000000000001/gopurs-js-bigints) | ~ 4625.00 ms | ~ 1596.00 ms | ~ 1443.00 ms | ~ 1179.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.25x
[gopurs-js-date](https://github.com/0x000000000000000000001/gopurs-js-date) | ~ 16044.00 ms | ~ 6227.00 ms | ~ 5297.00 ms | ~ 5007.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.31x
[gopurs-js-promise](https://github.com/0x000000000000000000001/gopurs-js-promise) | ~ 8189.00 ms | ~ 2787.00 ms | ~ 2531.00 ms | ~ 2196.00 ms
↳ /JS | — | 0.34x | 0.31x | 0.27x
[gopurs-js-promise-aff](https://github.com/0x000000000000000000001/gopurs-js-promise-aff) | ~ 16617.00 ms | ~ 6434.00 ms | ~ 5511.00 ms | ~ 5158.00 ms
↳ /JS | — | 0.39x | 0.33x | 0.31x
[gopurs-js-uri](https://github.com/0x000000000000000000001/gopurs-js-uri) | ~ 9730.00 ms | ~ 3275.00 ms | ~ 2948.00 ms | ~ 2569.00 ms
↳ /JS | — | 0.34x | 0.30x | 0.26x
[gopurs-lazy](https://github.com/0x000000000000000000001/gopurs-lazy) | ~ 7605.00 ms | ~ 2758.00 ms | ~ 2526.00 ms | ~ 2178.00 ms
↳ /JS | — | 0.36x | 0.33x | 0.29x
[gopurs-node-buffer](https://github.com/0x000000000000000000001/gopurs-node-buffer) | ~ 14716.00 ms | ~ 5351.00 ms | ~ 4743.00 ms | ~ 4366.00 ms
↳ /JS | — | 0.36x | 0.32x | 0.30x
[gopurs-node-event-emitter](https://github.com/0x000000000000000000001/gopurs-node-event-emitter) | ~ 21845.00 ms | ~ 8936.00 ms | ~ 6631.00 ms | ~ 7607.00 ms
↳ /JS | — | 0.41x | 0.30x | 0.35x
[gopurs-node-fs](https://github.com/0x000000000000000000001/gopurs-node-fs) | ~ 18796.00 ms | ~ 6999.00 ms | ~ 5912.00 ms | ~ 5685.00 ms
↳ /JS | — | 0.37x | 0.31x | 0.30x
[gopurs-node-http](https://github.com/0x000000000000000000001/gopurs-node-http) | ~ 20650.00 ms | ~ 7475.00 ms | ~ 6248.00 ms | ~ 6173.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.30x
[gopurs-node-net](https://github.com/0x000000000000000000001/gopurs-node-net) | ~ 19243.00 ms | ~ 6945.00 ms | ~ 5865.00 ms | ~ 5630.00 ms
↳ /JS | — | 0.36x | 0.30x | 0.29x
[gopurs-node-path](https://github.com/0x000000000000000000001/gopurs-node-path) | ~ 4257.00 ms | ~ 1302.00 ms | ~ 1204.00 ms | ~ 914.00 ms
↳ /JS | — | 0.31x | 0.28x | 0.21x
[gopurs-node-process](https://github.com/0x000000000000000000001/gopurs-node-process) | ~ 17942.00 ms | ~ 6684.00 ms | ~ 5685.00 ms | ~ 5391.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.30x
[gopurs-node-streams](https://github.com/0x000000000000000000001/gopurs-node-streams) | ~ 23385.00 ms | ~ 9204.00 ms | ~ 6884.00 ms | ~ 7938.00 ms
↳ /JS | — | 0.39x | 0.29x | 0.34x
[gopurs-now](https://github.com/0x000000000000000000001/gopurs-now) | ~ 18036.00 ms | ~ 6756.00 ms | ~ 5772.00 ms | ~ 5480.00 ms
↳ /JS | — | 0.37x | 0.32x | 0.30x
[gopurs-nullable](https://github.com/0x000000000000000000001/gopurs-nullable) | ~ 14155.00 ms | ~ 5343.00 ms | ~ 4659.00 ms | ~ 4399.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.31x
[gopurs-numbers](https://github.com/0x000000000000000000001/gopurs-numbers) | ~ 10142.00 ms | ~ 3507.00 ms | ~ 3170.00 ms | ~ 2749.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-ordered-collections](https://github.com/0x000000000000000000001/gopurs-ordered-collections) | ~ 16381.00 ms | ~ 6449.00 ms | ~ 5138.00 ms | ~ 5047.00 ms
↳ /JS | — | 0.39x | 0.31x | 0.31x
[gopurs-partial](https://github.com/0x000000000000000000001/gopurs-partial) | ~ 9519.00 ms | ~ 3188.00 ms | ~ 2859.00 ms | ~ 2493.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-prelude](https://github.com/0x000000000000000000001/gopurs-prelude) | ~ 10711.00 ms | ~ 3557.00 ms | ~ 3053.00 ms | ~ 2792.00 ms
↳ /JS | — | 0.33x | 0.29x | 0.26x
[gopurs-random](https://github.com/0x000000000000000000001/gopurs-random) | ~ 9902.00 ms | ~ 3224.00 ms | ~ 2880.00 ms | ~ 2525.00 ms
↳ /JS | — | 0.33x | 0.29x | 0.25x
[gopurs-record](https://github.com/0x000000000000000000001/gopurs-record) | ~ 10133.00 ms | ~ 3356.00 ms | ~ 3019.00 ms | ~ 2627.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-refs](https://github.com/0x000000000000000000001/gopurs-refs) | ~ 9881.00 ms | ~ 3300.00 ms | ~ 2960.00 ms | ~ 2594.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-run](https://github.com/0x000000000000000000001/gopurs-run) | ~ 17665.00 ms | ~ 6775.00 ms | ~ 5776.00 ms | ~ 5433.00 ms
↳ /JS | — | 0.38x | 0.33x | 0.31x
[gopurs-spec](https://github.com/0x000000000000000000001/gopurs-spec) | ~ 30416.00 ms | ~ 12821.00 ms | ~ 8658.00 ms | ~ 10954.00 ms
↳ /JS | — | 0.42x | 0.28x | 0.36x
[gopurs-st](https://github.com/0x000000000000000000001/gopurs-st) | ~ 9659.00 ms | ~ 3302.00 ms | ~ 2980.00 ms | ~ 2597.00 ms
↳ /JS | — | 0.34x | 0.31x | 0.27x
[gopurs-strings](https://github.com/0x000000000000000000001/gopurs-strings) | ~ 12680.00 ms | ~ 4489.00 ms | ~ 3824.00 ms | ~ 3805.00 ms
↳ /JS | — | 0.35x | 0.30x | 0.30x
[gopurs-strings-extra](https://github.com/0x000000000000000000001/gopurs-strings-extra) | ~ 13757.00 ms | ~ 6990.00 ms | ~ 6343.00 ms | ~ 4827.00 ms
↳ /JS | — | 0.51x | 0.46x | 0.35x
[gopurs-unfoldable](https://github.com/0x000000000000000000001/gopurs-unfoldable) | ~ 9565.00 ms | ~ 3322.00 ms | ~ 2979.00 ms | ~ 2607.00 ms
↳ /JS | — | 0.35x | 0.31x | 0.27x
[gopurs-unsafe-coerce](https://github.com/0x000000000000000000001/gopurs-unsafe-coerce) | ~ 9841.00 ms | ~ 3290.00 ms | ~ 2941.00 ms | ~ 2562.00 ms
↳ /JS | — | 0.33x | 0.30x | 0.26x
[gopurs-uuid](https://github.com/0x000000000000000000001/gopurs-uuid) | ~ 20972.00 ms | ~ 8748.00 ms | ~ 6505.00 ms | ~ 7340.00 ms
↳ /JS | — | 0.42x | 0.31x | 0.35x
[gopurs-variant](https://github.com/0x000000000000000000001/gopurs-variant) | ~ 11455.00 ms | ~ 4451.00 ms | ~ 3881.00 ms | ~ 3279.00 ms
↳ /JS | — | 0.39x | 0.34x | 0.29x
[gopurs-yoga-json](https://github.com/0x000000000000000000001/gopurs-yoga-json) | ~ 36226.00 ms | ~ 16572.00 ms | ~ 11991.00 ms | ~ 13134.00 ms
↳ /JS | — | 0.46x | 0.33x | 0.36x
**Total gopurs-*** | <strong>~ 666.09 s</strong> | <strong>~ 254.68 s</strong> | <strong>~ 212.02 s</strong> | <strong>~ 203.99 s</strong>
↳ /JS | — | <strong>0.38x</strong> | <strong>0.32x</strong> | <strong>0.31x</strong>
↳ μ/JS | — | <strong>0.37x</strong> | <strong>0.32x</strong> | <strong>0.29x</strong>
↳ Q₂/JS | — | <strong>0.37x</strong> | <strong>0.31x</strong> | <strong>0.30x</strong>
↳ G/JS | — | <strong>0.37x</strong> | <strong>0.32x</strong> | <strong>0.29x</strong>
**Total** | <strong>~ 812.48 s</strong> | <strong>~ 332.71 s</strong> | <strong>~ 265.63 s</strong> | <strong>~ 256.92 s</strong>
↳ /JS | — | <strong>0.41x</strong> | <strong>0.33x</strong> | <strong>0.32x</strong>
↳ μ/JS | — | <strong>0.38x</strong> | <strong>0.32x</strong> | <strong>0.30x</strong>
↳ Q₂/JS | — | <strong>0.37x</strong> | <strong>0.31x</strong> | <strong>0.30x</strong>
↳ G/JS | — | <strong>0.37x</strong> | <strong>0.32x</strong> | <strong>0.29x</strong>

#### ... to Rust

Benchmark   | [purust](https://github.com/0x000000000000000000001/purust) JS binary<br>(official compiler)<br>(WIP) | [purust](https://github.com/0x000000000000000000001/purust) Rust binary<br>(compiled by [purust](https://github.com/0x000000000000000000001/purust))<br>(WIP)
----------- | --------- | ---------------------
[b8x](https://github.com/0x000000000000000000001/b8x.pub) | ~ 44.31 s | ~ 20.62 s
↳ /JS | — | 0.47x
[purust-aff](https://github.com/0x000000000000000000001/purust-aff) | ~ 6218.00 ms | ~ 2076.00 ms
↳ /JS | — | 0.33x
[purust-argonaut-codecs](https://github.com/0x000000000000000000001/purust-argonaut-codecs) | ~ 6090.00 ms | ~ 2120.00 ms
↳ /JS | — | 0.35x
[purust-argonaut-core](https://github.com/0x000000000000000000001/purust-argonaut-core) | ~ 5122.00 ms | ~ 1682.00 ms
↳ /JS | — | 0.33x
[purust-arraybuffer-types](https://github.com/0x000000000000000000001/purust-arraybuffer-types) | ~ 591.00 ms | ~ 158.00 ms
↳ /JS | — | 0.27x
[purust-arrays](https://github.com/0x000000000000000000001/purust-arrays) | ~ 15134.00 ms | ~ 4291.00 ms
↳ /JS | — | 0.28x
[purust-assert](https://github.com/0x000000000000000000001/purust-assert) | ~ 631.00 ms | ~ 170.00 ms
↳ /JS | — | 0.27x
[purust-avar](https://github.com/0x000000000000000000001/purust-avar) | ~ 5942.00 ms | ~ 1938.00 ms
↳ /JS | — | 0.33x
[purust-catenable-lists](https://github.com/0x000000000000000000001/purust-catenable-lists) | ~ 3474.00 ms | ~ 1177.00 ms
↳ /JS | — | 0.34x
[purust-console](https://github.com/0x000000000000000000001/purust-console) | ~ 664.00 ms | ~ 183.00 ms
↳ /JS | — | 0.28x
[purust-datetime](https://github.com/0x000000000000000000001/purust-datetime) | ~ 5659.00 ms | ~ 1902.00 ms
↳ /JS | — | 0.34x
[purust-effect](https://github.com/0x000000000000000000001/purust-effect) | ~ 1240.00 ms | ~ 475.00 ms
↳ /JS | — | 0.38x
[purust-enums](https://github.com/0x000000000000000000001/purust-enums) | ~ 2950.00 ms | ~ 923.00 ms
↳ /JS | — | 0.31x
[purust-exceptions](https://github.com/0x000000000000000000001/purust-exceptions) | ~ 924.00 ms | ~ 276.00 ms
↳ /JS | — | 0.30x
[purust-exists](https://github.com/0x000000000000000000001/purust-exists) | ~ 629.00 ms | ~ 171.00 ms
↳ /JS | — | 0.27x
[purust-foldable-traversable](https://github.com/0x000000000000000000001/purust-foldable-traversable) | ~ 3593.00 ms | ~ 1220.00 ms
↳ /JS | — | 0.34x
[purust-foreign](https://github.com/0x000000000000000000001/purust-foreign) | ~ 4949.00 ms | ~ 1647.00 ms
↳ /JS | — | 0.33x
[purust-foreign-object](https://github.com/0x000000000000000000001/purust-foreign-object) | ~ 6320.00 ms | ~ 2217.00 ms
↳ /JS | — | 0.35x
[purust-free](https://github.com/0x000000000000000000001/purust-free) | ~ 4570.00 ms | ~ 1520.00 ms
↳ /JS | — | 0.33x
[purust-functions](https://github.com/0x000000000000000000001/purust-functions) | ~ 1365.00 ms | ~ 508.00 ms
↳ /JS | — | 0.37x
[purust-integers](https://github.com/0x000000000000000000001/purust-integers) | ~ 4359.00 ms | ~ 1398.00 ms
↳ /JS | — | 0.32x
[purust-js-bigints](https://github.com/0x000000000000000000001/purust-js-bigints) | ~ 6620.00 ms | ~ 2336.00 ms
↳ /JS | — | 0.35x
[purust-js-date](https://github.com/0x000000000000000000001/purust-js-date) | ~ 5642.00 ms | ~ 1933.00 ms
↳ /JS | — | 0.34x
[purust-js-promise](https://github.com/0x000000000000000000001/purust-js-promise) | ~ 6029.00 ms | ~ 2018.00 ms
↳ /JS | — | 0.33x
[purust-js-promise-aff](https://github.com/0x000000000000000000001/purust-js-promise-aff) | ~ 6165.00 ms | ~ 2078.00 ms
↳ /JS | — | 0.34x
[purust-lazy](https://github.com/0x000000000000000000001/purust-lazy) | ~ 2140.00 ms | ~ 693.00 ms
↳ /JS | — | 0.32x
[purust-node-buffer](https://github.com/0x000000000000000000001/purust-node-buffer) | ~ 4894.00 ms | ~ 1601.00 ms
↳ /JS | — | 0.33x
[purust-node-child-process](https://github.com/0x000000000000000000001/purust-node-child-process) | ~ 7651.00 ms | ~ 2609.00 ms
↳ /JS | — | 0.34x
[purust-node-event-emitter](https://github.com/0x000000000000000000001/purust-node-event-emitter) | ~ 8202.00 ms | ~ 2809.00 ms
↳ /JS | — | 0.34x
[purust-node-fs](https://github.com/0x000000000000000000001/purust-node-fs) | ~ 7430.00 ms | ~ 2508.00 ms
↳ /JS | — | 0.34x
[purust-node-http](https://github.com/0x000000000000000000001/purust-node-http) | ~ 7105.00 ms | ~ 2452.00 ms
↳ /JS | — | 0.35x
[purust-node-net](https://github.com/0x000000000000000000001/purust-node-net) | ~ 6805.00 ms | ~ 2310.00 ms
↳ /JS | — | 0.34x
[purust-node-os](https://github.com/0x000000000000000000001/purust-node-os) | ~ 6032.00 ms | ~ 2086.00 ms
↳ /JS | — | 0.35x
[purust-node-path](https://github.com/0x000000000000000000001/purust-node-path) | ~ 650.00 ms | ~ 175.00 ms
↳ /JS | — | 0.27x
[purust-node-process](https://github.com/0x000000000000000000001/purust-node-process) | ~ 7196.00 ms | ~ 2443.00 ms
↳ /JS | — | 0.34x
[purust-node-streams](https://github.com/0x000000000000000000001/purust-node-streams) | ~ 9197.00 ms | ~ 3489.00 ms
↳ /JS | — | 0.38x
[purust-now](https://github.com/0x000000000000000000001/purust-now) | ~ 6780.00 ms | ~ 2310.00 ms
↳ /JS | — | 0.34x
[purust-nullable](https://github.com/0x000000000000000000001/purust-nullable) | ~ 4653.00 ms | ~ 1557.00 ms
↳ /JS | — | 0.33x
[purust-numbers](https://github.com/0x000000000000000000001/purust-numbers) | ~ 4401.00 ms | ~ 1408.00 ms
↳ /JS | — | 0.32x
[purust-ordered-collections](https://github.com/0x000000000000000000001/purust-ordered-collections) | ~ 8217.00 ms | ~ 2986.00 ms
↳ /JS | — | 0.36x
[purust-partial](https://github.com/0x000000000000000000001/purust-partial) | ~ 871.00 ms | ~ 268.00 ms
↳ /JS | — | 0.31x
[purust-prelude](https://github.com/0x000000000000000000001/purust-prelude) | ~ 2793.00 ms | ~ 985.00 ms
↳ /JS | — | 0.35x
[purust-quickcheck](https://github.com/0x000000000000000000001/purust-quickcheck) | ~ 4753.00 ms | ~ 1624.00 ms
↳ /JS | — | 0.34x
[purust-random](https://github.com/0x000000000000000000001/purust-random) | ~ 2660.00 ms | ~ 898.00 ms
↳ /JS | — | 0.34x
[purust-record](https://github.com/0x000000000000000000001/purust-record) | ~ 2641.00 ms | ~ 896.00 ms
↳ /JS | — | 0.34x
[purust-refs](https://github.com/0x000000000000000000001/purust-refs) | ~ 718.00 ms | ~ 196.00 ms
↳ /JS | — | 0.27x
[purust-run](https://github.com/0x000000000000000000001/purust-run) | ~ 6914.00 ms | ~ 2262.00 ms
↳ /JS | — | 0.33x
[purust-spec](https://github.com/0x000000000000000000001/purust-spec) | ~ 12717.00 ms | ~ 4633.00 ms
↳ /JS | — | 0.36x
[purust-spec-discovery](https://github.com/0x000000000000000000001/purust-spec-discovery) | ~ 11401.00 ms | ~ 4318.00 ms
↳ /JS | — | 0.38x
[purust-spec-node](https://github.com/0x000000000000000000001/purust-spec-node) | ~ 11931.00 ms | ~ 4307.00 ms
↳ /JS | — | 0.36x
[purust-st](https://github.com/0x000000000000000000001/purust-st) | ~ 2554.00 ms | ~ 866.00 ms
↳ /JS | — | 0.34x
[purust-strings](https://github.com/0x000000000000000000001/purust-strings) | ~ 6589.00 ms | ~ 1851.00 ms
↳ /JS | — | 0.28x
[purust-strings-extra](https://github.com/0x000000000000000000001/purust-strings-extra) | ~ 7666.00 ms | ~ 11028.00 ms
↳ /JS | — | 1.44x
[purust-unfoldable](https://github.com/0x000000000000000000001/purust-unfoldable) | ~ 2651.00 ms | ~ 920.00 ms
↳ /JS | — | 0.35x
[purust-unsafe-coerce](https://github.com/0x000000000000000000001/purust-unsafe-coerce) | ~ 2089.00 ms | ~ 669.00 ms
↳ /JS | — | 0.32x
[purust-uuid](https://github.com/0x000000000000000000001/purust-uuid) | ~ 7613.00 ms | ~ 2588.00 ms
↳ /JS | — | 0.34x
[purust-variant](https://github.com/0x000000000000000000001/purust-variant) | ~ 4537.00 ms | ~ 1405.00 ms
↳ /JS | — | 0.31x
[purust-yoga-json](https://github.com/0x000000000000000000001/purust-yoga-json) | ~ 16694.00 ms | ~ 14548.00 ms
↳ /JS | — | 0.87x
**Total purust-*** | <strong>~ 304.03 s</strong> | <strong>~ 120.12 s</strong>
↳ /JS | — | <strong>0.40x</strong>
↳ μ/JS | — | <strong>0.36x</strong>
↳ Q₂/JS | — | <strong>0.34x</strong>
↳ G/JS | — | <strong>0.34x</strong>
**Total** | <strong>~ 348.34 s</strong> | <strong>~ 140.73 s</strong>
↳ /JS | — | <strong>0.40x</strong>
↳ μ/JS | — | <strong>0.36x</strong>
↳ Q₂/JS | — | <strong>0.34x</strong>
↳ G/JS | — | <strong>0.35x</strong>
