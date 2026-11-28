# 🌈 PureScript universal multi-runtime benchmark

## Project goal
This project is a **proof of concept** demonstrating the power of abstraction and portability offered by **PureScript**. The goal is to show how the exact same pure functional code (without any manual FFI) can be compiled and executed natively on radically different ecosystems. Backends consume either standard `CoreFn` or, for the local AOT backends, the typed `TAST/tcorefn` representation exported by a local PureScript fork (may be subject to an official PR soon).

## Comprehensive benchmarks
The benchmark suite runs a **wide variety** of computationally intensive tasks: AST evaluation, massive list processing, tail call optimization, deep record updates... These tests apply massive **pressure** on the call stack, garbage collector, disk I/O, event loops, and runtime execution engines to measure the raw ability of the compiler and the underlying virtual machine. 

> [!NOTE]
> **Reading the benchmark tables**
>
> - **Status, only for gopurs, purust, phpurs, sharpurs and javapurs:** **Stable** = measured and considered fully optimized; **WIP** = measured, with optimization work remaining; **Backlog** = planned work without a published measurement or ratio. Status labels belong only to the relevant backend columns. References and external backends have no status. For Go Release (`gopurs`), 30 cases have graduated to **Stable** and 10 cases remain in **WIP** (all 40 cases measured); Stable is currently empty for the other controlled backends.
> - **Ratios:** `/C` uses the portable idiomatic C11 imperative reference (`clang -O3`, without hand-written SIMD intrinsics); `/JS` uses the first JavaScript column; `/Go` uses the first Go column. **Below 1× is faster; above 1× is slower.** Extended Results uses its own JS reference and has no C reference.
> - **Totals:** sums of the available per-test medians: **Total Stable** (30 cases), **Total WIP** (10 cases), and **Total** (all 40 cases) across all Go columns; **Total Stable** and **Total WIP** for the other controlled backends; **Total** for references and external backends. Backlog never contributes. Empty Stable groups are shown as *Coming soon...* and have no defined total, ratio or mean. These totals do not represent an application benchmark.
> - **Summaries:** `μ` = arithmetic mean, `Q₂` = median, `G` = geometric mean of the case ratios, and `Q₄` = maximum case ratio (worst case). Each case has equal weight. All ratio statistics use exactly the same measured cases on both sides: in Go, **Total Stable** matches the 30 Stable cases, **Total WIP** matches the 10 WIP cases, and **Total** matches all 40 cases across all five Go columns. `Q₄/C` and `Q₄/JS` appear wherever `Q₂/C` and `Q₂/JS` appear. They use the same C, JavaScript and Go references as the other ratios. The full total of a reference table is not used as a denominator for a different set of cases. Very small reference times can produce large ratios.
> - **Missing data:** `—` means unavailable or not applicable, never zero. Bounds (`< …`) are excluded from summaries. `n=a/b` indicates coverage: measured cases out of listed cases for a time total, or matched references out of measured cases for a ratio. Diagnostic cases (`Array Indexing`, `JSON to Typed AST`, `JSON Decoding`) are measured across all columns of JavaScript (40/40), Go (40/40), and C (40/40) (other controlled backends keep them in unmeasured Backlog).
> - **Display precision, in every table:** values ≥ 1 use two decimal places (`12.000045 → 12.00`); positive values < 1 keep two significant digits after leading zeros (`0.000045`, `0.009492 → 0.0095`). Zero is `0.00`.
> - **Opaque inputs, in every table:** inputs must remain unknown to the optimizer at runtime. Algebraic simplification is allowed; substituting the driver’s known input is not.
> - **Calculations:** use archived timings before display rounding, including for threshold checks. `≈` marks a fallback to published precision.

### Core stresstest benchmark results (pure computational)

**Core tests (`src/`)** cover features common to all tested targets (`./bin/run`)

#### JavaScript

JS Benchmark            | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written FP-style JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written imperative JS FFI<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT |
----------------------- | ---------------------- | ------------------ | ---------------------- | ---------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.26 μs | ~ 0.19 μs | ~ 0.20 μs | ~ 0.22 μs |
| ↳ /C | 2.67x | 1.95x | 2.03x | 2.27x |
| ↳ /JS | — | 0.73x | 0.76x | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.30 μs | ~ 0.31 μs | ~ 0.31 μs | ~ 0.37 μs |
| ↳ /C | 3.88x | 3.98x | 4.03x | 4.75x |
| ↳ /JS | — | 1.02x | 1.04x | 1.22x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 8.64 μs | ~ 5.06 μs | ~ 19.06 μs | ~ 0.49 μs |
| ↳ /C | 7407.51x | 4338.64x | 16345.95x | 423.43x |
| ↳ /JS | — | 0.59x | 2.21x | 0.057x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 187.10 μs | ~ 151.13 μs | ~ 65.91 μs | ~ 68.31 μs |
| ↳ /C | 5.77x | 4.66x | 2.03x | 2.11x |
| ↳ /JS | — | 0.81x | 0.35x | 0.37x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 74.34 μs | ~ 134.11 μs | ~ 42.07 μs | ~ 7.69 μs |
| ↳ /C | 22.99x | 41.47x | 13.01x | 2.38x |
| ↳ /JS | — | 1.80x | 0.57x | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 66.77 μs | ~ 64.40 μs | ~ 38.25 μs | ~ 38.48 μs |
| ↳ /C | 4.09x | 3.95x | 2.35x | 2.36x |
| ↳ /JS | — | 0.96x | 0.57x | 0.58x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1269.89 μs | ~ 1242.29 μs | ~ 1337.51 μs | ~ 49.63 μs |
| ↳ /C | 1124222.16x | 1099793.63x | 1184090.18x | 43937.12x |
| ↳ /JS | — | 0.98x | 1.05x | 0.039x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 58.01 μs | ~ 26.99 μs | ~ 86.33 μs | ~ 0.65 μs |
| ↳ /C | 785.72x | 365.56x | 1169.25x | 8.82x |
| ↳ /JS | — | 0.47x | 1.49x | 0.011x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 84263.21 μs | ~ 47272.88 μs | ~ 24486.83 μs | ~ 27249.58 μs |
| ↳ /C | 23.26x | 13.05x | 6.76x | 7.52x |
| ↳ /JS | — | 0.56x | 0.29x | 0.32x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 11836.88 μs | ~ 7368.38 μs | ~ 4843.00 μs | ~ 2805.25 μs |
| ↳ /C | 10231315.57x | 6368925.07x | 4186093.15x | 2424744.54x |
| ↳ /JS | — | 0.62x | 0.41x | 0.24x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 47.34 μs | ~ 10.95 μs | ~ 46.87 μs | ~ 0.42 μs |
| ↳ /C | 41715.85x | 9646.29x | 41302.64x | 368.16x |
| ↳ /JS | — | 0.23x | 0.99x | 0.0088x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 14058.08 μs | ~ 10604.88 μs | ~ 17513.12 μs | ~ 321.48 μs |
| ↳ /C | 12384765.90x | 9342588.52x | 15428557.21x | 283212.72x |
| ↳ /JS | — | 0.75x | 1.25x | 0.023x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.07 μs | ~ 3.75 μs | ~ 3.12 μs | ~ 0.50 μs |
| ↳ /C | 5108.96x | 2709.90x | 2254.97x | 360.05x |
| ↳ /JS | — | 0.53x | 0.44x | 0.070x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.035 μs | ~ 0.017 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 30.66x | 14.83x | 51.86x | 17.69x |
| ↳ /JS | — | 0.48x | 1.69x | 0.58x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 121.94 μs | ~ 144.60 μs | ~ 92.20 μs | ~ 12.93 μs |
| ↳ /C | 271.66x | 322.15x | 205.40x | 28.82x |
| ↳ /JS | — | 1.19x | 0.76x | 0.11x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 127.59 μs | ~ 130.46 μs | ~ 74.63 μs | ~ 18.58 μs |
| ↳ /C | 27.55x | 28.17x | 16.11x | 4.01x |
| ↳ /JS | — | 1.02x | 0.58x | 0.15x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 70.25 μs | ~ 70.22 μs | ~ 44.01 μs | ~ 3.32 μs |
| ↳ /C | 243.58x | 243.49x | 152.60x | 11.51x |
| ↳ /JS | — | 1.00x | 0.63x | 0.047x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 561.69 μs | ~ 517.87 μs | ~ 114.68 μs | ~ 78.30 μs |
| ↳ /C | 197.40x | 182.00x | 40.30x | 27.52x |
| ↳ /JS | — | 0.92x | 0.20x | 0.14x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 51.63 μs | ~ 50.71 μs | ~ 41.04 μs | ~ 5.03 μs |
| ↳ /C | 77.00x | 75.63x | 61.20x | 7.49x |
| ↳ /JS | — | 0.98x | 0.79x | 0.097x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 30.20 μs | ~ 31.05 μs | ~ 2.05 μs | ~ 2.40 μs |
| ↳ /C | 346.69x | 356.43x | 23.56x | 27.52x |
| ↳ /JS | — | 1.03x | 0.068x | 0.079x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 80.61 μs | ~ 81.04 μs | ~ 49.83 μs | ~ 1.39 μs |
| ↳ /C | 1822.20x | 1832.03x | 1126.52x | 31.40x |
| ↳ /JS | — | 1.01x | 0.62x | 0.017x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 68.35 μs | ~ 65.14 μs | ~ 42.18 μs | ~ 1.07 μs |
| ↳ /C | 50737.46x | 48355.34x | 31306.38x | 791.99x |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 68.01 μs | ~ 64.65 μs | ~ 42.45 μs | ~ 1.08 μs |
| ↳ /C | 198.60x | 188.77x | 123.96x | 3.14x |
| ↳ /JS | — | 0.95x | 0.62x | 0.016x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.010 μs | ~ 0.0087 μs | ~ 0.019 μs | ~ 0.020 μs |
| ↳ /C | 8.99x | 7.76x | 17.26x | 17.40x |
| ↳ /JS | — | 0.86x | 1.92x | 1.94x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 37.44 μs | ~ 37.44 μs | ~ 37.09 μs | ~ 8.54 μs |
| ↳ /C | 30.56x | 30.56x | 30.27x | 6.97x |
| ↳ /JS | — | 1.00x | 0.99x | 0.23x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs | ~ 0.19 μs |
| ↳ /C | 0.17x | 0.17x | 0.18x | 0.17x |
| ↳ /JS | — | 0.99x | 1.02x | 1.00x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 63.80 μs | ~ 63.97 μs | ~ 63.94 μs | ~ 63.72 μs |
| ↳ /C | 50.22x | 50.36x | 50.33x | 50.17x |
| ↳ /JS | — | 1.00x | 1.00x | 1.00x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 45.28 μs | ~ 44.96 μs | ~ 42.10 μs | ~ 21.01 μs |
| ↳ /C | 67.62x | 67.15x | 62.87x | 31.38x |
| ↳ /JS | — | 0.99x | 0.93x | 0.46x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 145.33 μs | ~ 144.87 μs | ~ 138.40 μs | ~ 24.69 μs |
| ↳ /C | 61.12x | 60.93x | 58.21x | 10.38x |
| ↳ /JS | — | 1.00x | 0.95x | 0.17x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 47.71 μs | ~ 46.52 μs | ~ 42.64 μs | ~ 11.85 μs |
| ↳ /C | 49.52x | 48.28x | 44.25x | 12.30x |
| ↳ /JS | — | 0.97x | 0.89x | 0.25x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 2.44 μs | ~ 4.37 μs | ~ 4.32 μs | ~ 0.72 μs |
| ↳ /C | 2115.43x | 3783.88x | 3743.89x | 627.60x |
| ↳ /JS | — | 1.79x | 1.77x | 0.30x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 4.55 μs | ~ 9.34 μs | ~ 9.04 μs | ~ 0.71 μs |
| ↳ /C | 3925.98x | 8055.82x | 7792.43x | 615.50x |
| ↳ /JS | — | 2.05x | 1.98x | 0.16x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 7.03 μs | ~ 14.73 μs | ~ 14.38 μs | ~ 0.72 μs |
| ↳ /C | 5082.04x | 10649.92x | 10395.32x | 521.37x |
| ↳ /JS | — | 2.10x | 2.05x | 0.10x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 463.87 μs | ~ 106.02 μs | ~ 2.44 μs | ~ 2.48 μs |
| ↳ /C | 411880.79x | 94139.57x | 2164.35x | 2200.57x |
| ↳ /JS | — | 0.23x | 0.0053x | 0.0053x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 513.20 μs | ~ 106.03 μs | ~ 31.67 μs | ~ 3.41 μs |
| ↳ /C | 456430.14x | 94302.67x | 28169.48x | 3031.43x |
| ↳ /JS | — | 0.21x | 0.062x | 0.0066x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 155.67 μs | ~ 102.38 μs | ~ 2.39 μs | ~ 2.56 μs |
| ↳ /C | 138936.00x | 91369.69x | 2131.25x | 2287.29x |
| ↳ /JS | — | 0.66x | 0.015x | 0.016x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 211.84 μs | ~ 113.53 μs | ~ 31.14 μs | ~ 3.52 μs |
| ↳ /C | 189492.19x | 101548.30x | 27857.21x | 3151.35x |
| ↳ /JS | — | 0.54x | 0.15x | 0.017x |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 25303.33 μs | ~ 13027.00 μs | ~ 5986.71 μs | ~ 3899.42 μs |
| ↳ /C | 3.84x | 1.98x | 0.91x | 0.59x |
| ↳ /JS | — | 0.51x | 0.24x | 0.15x |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 80646.58 μs | ~ 63267.67 μs | ~ 24223.88 μs | ~ 23845.50 μs |
| ↳ /C | 7.64x | 5.99x | 2.29x | 2.26x |
| ↳ /JS | — | 0.78x | 0.30x | 0.30x |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 9283.94 μs | ~ 6252.79 μs | ~ 2210.58 μs | ~ 1733.83 μs |
| ↳ /C | 14.12x | 9.51x | 3.36x | 2.64x |
| ↳ /JS | — | 0.67x | 0.24x | 0.19x |
| **Total** | <strong>229.99 ms</strong> | <strong>~ 151.38 ms</strong> | <strong>~ 81.83 ms</strong> | <strong>~ 60.29 ms</strong> |
| ↳ /C | <strong>10.70x</strong> | <strong>7.04x</strong> | <strong>3.81x</strong> | <strong>2.80x</strong> |
| ↳ μ/C | <strong>626437.34x</strong> | <strong>432104.20x</strong> | <strong>524386.88x</strong> | <strong>69164.92x</strong> |
| ↳ Q₂/C | <strong>198.00x</strong> | <strong>185.38x</strong> | <strong>62.04x</strong> | <strong>22.60x</strong> |
| ↳ Q₄/C | <strong>12384765.90x</strong> | <strong>9342588.52x</strong> | <strong>15428557.21x</strong> | <strong>2424744.54x</strong> |
| ↳ G/C | <strong>522.48x</strong> | <strong>417.78x</strong> | <strong>269.98x</strong> | <strong>57.32x</strong> |
| ↳ /JS | — | <strong>0.66x</strong> | <strong>0.36x</strong> | <strong>0.26x</strong> |
| ↳ μ/JS | — | <strong>0.90x</strong> | <strong>0.81x</strong> | <strong>0.29x</strong> |
| ↳ Q₂/JS | — | <strong>0.95x</strong> | <strong>0.69x</strong> | <strong>0.14x</strong> |
| ↳ Q₄/JS | — | <strong>2.10x</strong> | <strong>2.21x</strong> | <strong>1.94x</strong> |
| ↳ G/JS | — | <strong>0.80x</strong> | <strong>0.52x</strong> | <strong>0.11x</strong> |

#### Go

Go Benchmark            | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs) <br>`--dev`<br>↓<br>Go<br>↓<br>go-build<br>`-N -l` | Hand-written PureScript<br>↓<br>[psgo](https://github.com/i-am-the-slime/purescript-native)<br>↓<br>Go<br>↓<br>go-build | Hand-written FP-style Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written imperative Go FFI<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build |
----------------------- | -------------------------------- | -------------------------------- | ------------------------------------- | ---------------------- | ---------------------------- |
|  | <div align="center">**Stable**<br>↓</div> |  |  |  |  |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.011 μs | ~ 0.044 μs | ~ 4.15 μs | ~ 0.28 μs | ~ 0.28 μs |
| ↳ /C | 0.11x | 0.46x | 43.09x | 2.95x | 2.91x |
| ↳ /JS | 0.042x | 0.17x | 16.13x | 1.11x | 1.09x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs | ~ 0.19 μs | ~ 5.46 μs | ~ 0.14 μs | ~ 0.14 μs |
| ↳ /C | 1.48x | 2.40x | 69.81x | 1.83x | 1.82x |
| ↳ /JS | 0.38x | 0.62x | 17.98x | 0.47x | 0.47x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 39.00 μs | ~ 298.01 μs | ~ 6300.02 μs | ~ 35.19 μs | ~ 34.91 μs |
| ↳ /C | 1.20x | 9.19x | 194.35x | 1.09x | 1.08x |
| ↳ /JS | 0.21x | 1.59x | 33.67x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.47 μs | ~ 193.05 μs | ~ 5348.88 μs | ~ 39.64 μs | ~ 14.43 μs |
| ↳ /C | 2.31x | 59.70x | 1654.13x | 12.26x | 4.46x |
| ↳ /JS | 0.10x | 2.60x | 71.95x | 0.53x | 0.19x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.76 μs | ~ 27.05 μs | ~ 400.98 μs | ~ 19.34 μs | ~ 19.26 μs |
| ↳ /C | 0.97x | 1.66x | 24.59x | 1.19x | 1.18x |
| ↳ /JS | 0.24x | 0.41x | 6.01x | 0.29x | 0.29x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 1581.81 μs | ~ 5534.02 μs | ~ 856171.62 μs | ~ 22952.46 μs | ~ 8739.88 μs |
| ↳ /C | 0.44x | 1.53x | 236.35x | 6.34x | 2.41x |
| ↳ /JS | 0.019x | 0.066x | 10.16x | 0.27x | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0015 μs | ~ 0.017 μs | ~ 527241.79 μs | ~ 46836.00 μs | ~ 2226.40 μs |
| ↳ /C | 1.28x | 14.46x | 455726460.47x | 40483142.40x | 1924411.00x |
| ↳ /JS | 0.00000012x | 0.0000014x | 44.54x | 3.96x | 0.19x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.0015 μs | ~ 0.017 μs | ~ 291.63 μs | ~ 46.58 μs | ~ 0.30 μs |
| ↳ /C | 1.30x | 14.84x | 256981.79x | 41045.62x | 262.22x |
| ↳ /JS | 0.000031x | 0.00036x | 6.16x | 0.98x | 0.0063x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.0015 μs | ~ 0.017 μs | ~ 70764.92 μs | ~ 13404.08 μs | ~ 0.24 μs |
| ↳ /C | 1.30x | 14.73x | 62341847.63x | 11808609.91x | 211.25x |
| ↳ /JS | 0.00000010x | 0.0000012x | 5.03x | 0.95x | 0.000017x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0012 μs | ~ 0.011 μs | ~ 0.51 μs | ~ 0.045 μs | ~ 0.0068 μs |
| ↳ /C | 1.06x | 9.93x | 442.62x | 39.03x | 5.87x |
| ↳ /JS | 0.035x | 0.32x | 14.44x | 1.27x | 0.19x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 2.83 μs | ~ 34.91 μs | ~ 1411.95 μs | ~ 68.31 μs | ~ 6.73 μs |
| ↳ /C | 0.61x | 7.54x | 304.84x | 14.75x | 1.45x |
| ↳ /JS | 0.022x | 0.27x | 11.07x | 0.54x | 0.053x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 0.24 μs | ~ 0.27 μs | ~ 134.15 μs | ~ 48.61 μs | ~ 2.28 μs |
| ↳ /C | 0.84x | 0.92x | 465.13x | 168.56x | 7.89x |
| ↳ /JS | 0.0034x | 0.0038x | 1.91x | 0.69x | 0.032x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 1.45 μs | ~ 37.88 μs | ~ 11598.04 μs | ~ 831.80 μs | ~ 10.50 μs |
| ↳ /C | 0.51x | 13.31x | 4075.98x | 292.32x | 3.69x |
| ↳ /JS | 0.0026x | 0.067x | 20.65x | 1.48x | 0.019x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 0.48 μs | ~ 0.53 μs | ~ 103.52 μs | ~ 47.85 μs | ~ 3.58 μs |
| ↳ /C | 0.71x | 0.79x | 154.38x | 71.36x | 5.34x |
| ↳ /JS | 0.0093x | 0.010x | 2.01x | 0.93x | 0.069x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 0.0024 μs | ~ 0.049 μs | ~ 308.76 μs | ~ 11.31 μs | ~ 0.79 μs |
| ↳ /C | 1.81x | 36.31x | 229189.48x | 8392.83x | 583.04x |
| ↳ /JS | 0.000036x | 0.00072x | 4.52x | 0.17x | 0.011x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 0.36 μs | ~ 5.10 μs | ~ 322.27 μs | ~ 11.17 μs | ~ 0.75 μs |
| ↳ /C | 1.04x | 14.88x | 941.02x | 32.60x | 2.18x |
| ↳ /JS | 0.0052x | 0.075x | 4.74x | 0.16x | 0.011x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.0015 μs | ~ 0.030 μs | ~ 3.49 μs | ~ 3.47 μs | ~ 3.48 μs |
| ↳ /C | 1.32x | 26.98x | 3107.15x | 3084.36x | 3097.54x |
| ↳ /JS | 0.15x | 3.00x | 345.65x | 343.11x | 344.58x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 0.54 μs | ~ 0.59 μs | ~ 66.55 μs | ~ 6.95 μs | ~ 3.39 μs |
| ↳ /C | 0.44x | 0.48x | 54.32x | 5.68x | 2.77x |
| ↳ /JS | 0.014x | 0.016x | 1.78x | 0.19x | 0.091x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 0.16 μs | ~ 0.28 μs | ~ 2.78 μs | ~ 0.17 μs | ~ 0.16 μs |
| ↳ /C | 0.15x | 0.26x | 2.52x | 0.15x | 0.15x |
| ↳ /JS | 0.86x | 1.49x | 14.58x | 0.87x | 0.86x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 0.17 μs | ~ 0.29 μs | ~ 2.54 μs | ~ 2.31 μs | ~ 2.34 μs |
| ↳ /C | 0.14x | 0.23x | 2.00x | 1.82x | 1.84x |
| ↳ /JS | 0.0027x | 0.0046x | 0.040x | 0.036x | 0.037x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 1.23 μs | ~ 4.77 μs | ~ 107.79 μs | ~ 145.98 μs | ~ 3.71 μs |
| ↳ /C | 1.84x | 7.12x | 160.97x | 218.00x | 5.54x |
| ↳ /JS | 0.027x | 0.11x | 2.38x | 3.22x | 0.082x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 1.89 μs | ~ 11.12 μs | ~ 306.20 μs | ~ 32.50 μs | ~ 9.63 μs |
| ↳ /C | 0.79x | 4.68x | 128.78x | 13.67x | 4.05x |
| ↳ /JS | 0.013x | 0.077x | 2.11x | 0.22x | 0.066x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 2.05 μs | ~ 6.58 μs | ~ 175.24 μs | ~ 74.24 μs | ~ 3.77 μs |
| ↳ /C | 2.13x | 6.83x | 181.88x | 77.05x | 3.91x |
| ↳ /JS | 0.043x | 0.14x | 3.67x | 1.56x | 0.079x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.0027 μs | ~ 0.051 μs | ~ 145.45 μs | ~ 0.71 μs | ~ 0.24 μs |
| ↳ /C | 2.33x | 44.56x | 125911.05x | 617.33x | 206.74x |
| ↳ /JS | 0.0011x | 0.021x | 59.52x | 0.29x | 0.098x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 0.0027 μs | ~ 0.050 μs | ~ 196.70 μs | ~ 7.22 μs | ~ 0.25 μs |
| ↳ /C | 2.35x | 43.50x | 169596.61x | 6225.45x | 217.59x |
| ↳ /JS | 0.00060x | 0.011x | 43.20x | 1.59x | 0.055x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 0.0030 μs | ~ 0.055 μs | ~ 711.00 μs | ~ 12.74 μs | ~ 0.26 μs |
| ↳ /C | 2.15x | 40.06x | 514117.60x | 9210.32x | 188.73x |
| ↳ /JS | 0.00042x | 0.0079x | 101.16x | 1.81x | 0.037x |
| [st/ref-write-10000](src/Test/STRefWrite10000/STRefWrite10000.purs) | ~ 0.0015 μs | ~ 0.029 μs | ~ 3098.28 μs | ~ 2.25 μs | ~ 2.25 μs |
| ↳ /C | 1.31x | 25.89x | 2751020.57x | 1996.91x | 1996.97x |
| ↳ /JS | 0.0000032x | 0.000063x | 6.68x | 0.0048x | 0.0048x |
| [st/ref-modify-10000](src/Test/STRefModify10000/STRefModify10000.purs) | ~ 0.0015 μs | ~ 0.030 μs | ~ 4076.21 μs | ~ 2.31 μs | ~ 2.39 μs |
| ↳ /C | 1.31x | 26.49x | 3625287.65x | 2058.56x | 2129.07x |
| ↳ /JS | 0.0000029x | 0.000058x | 7.94x | 0.0045x | 0.0047x |
| [effect/ref-write-10000](src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | ~ 0.0015 μs | ~ 444.95 μs | ~ 3183.40 μs | ~ 2.24 μs | ~ 2.26 μs |
| ↳ /C | 1.31x | 397118.37x | 2841180.88x | 2001.01x | 2015.13x |
| ↳ /JS | 0.0000095x | 2.86x | 20.45x | 0.014x | 0.015x |
| [effect/ref-modify-10000](src/Test/EffectRefModify10000/EffectRefModify10000.purs) | ~ 0.0015 μs | ~ 1230.03 μs | ~ 4042.36 μs | ~ 2.37 μs | ~ 2.34 μs |
| ↳ /C | 1.32x | 1100263.60x | 3615889.44x | 2119.43x | 2092.74x |
| ↳ /JS | 0.0000070x | 5.81x | 19.08x | 0.011x | 0.011x |
| **Total Stable** | <strong>~ 1.66 ms</strong> | <strong>~ 7.83 ms</strong> | <strong>~ 1496.53 ms</strong> | <strong>~ 84.65 ms</strong> | <strong>~ 11.10 ms</strong> |
| ↳ /C | <strong>0.45x</strong> | <strong>2.12x</strong> | <strong>405.45x</strong> | <strong>22.93x</strong> | <strong>3.01x</strong> |
| ↳ μ/C | <strong>1.20x</strong> | <strong>49927.06x</strong> | <strong>17740324.24x</strong> | <strong>1745648.83x</strong> | <strong>64582.35x</strong> |
| ↳ Q₂/C | <strong>1.29x</strong> | <strong>11.62x</strong> | <strong>703.08x</strong> | <strong>122.81x</strong> | <strong>5.44x</strong> |
| ↳ Q₄/C | <strong>2.35x</strong> | <strong>1100263.60x</strong> | <strong>455726460.47x</strong> | <strong>40483142.40x</strong> | <strong>1924411.00x</strong> |
| ↳ G/C | <strong>0.95x</strong> | <strong>13.48x</strong> | <strong>5050.84x</strong> | <strong>208.99x</strong> | <strong>30.70x</strong> |
| ↳ /JS | <strong>0.015x</strong> | <strong>0.069x</strong> | <strong>13.22x</strong> | <strong>0.75x</strong> | <strong>0.098x</strong> |
| ↳ μ/JS | <strong>0.072x</strong> | <strong>0.66x</strong> | <strong>29.97x</strong> | <strong>12.23x</strong> | <strong>11.63x</strong> |
| ↳ Q₂/JS | <strong>0.0072x</strong> | <strong>0.071x</strong> | <strong>10.61x</strong> | <strong>0.53x</strong> | <strong>0.068x</strong> |
| ↳ Q₄/JS | <strong>0.86x</strong> | <strong>5.81x</strong> | <strong>345.65x</strong> | <strong>343.11x</strong> | <strong>344.58x</strong> |
| ↳ G/JS | <strong>0.0018x</strong> | <strong>0.026x</strong> | <strong>9.81x</strong> | <strong>0.41x</strong> | <strong>0.060x</strong> |
|  | <div align="center">**WIP**<br>↓</div> |  |  |  |  |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.40 μs | ~ 28.35 μs | ~ 244.02 μs | ~ 23.20 μs | ~ 0.40 μs |
| ↳ /C | 346.83x | 24315.95x | 209321.08x | 19897.43x | 345.34x |
| ↳ /JS | 0.047x | 3.28x | 28.26x | 2.69x | 0.047x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 231.71 μs | ~ 1338.40 μs | ~ 3907.90 μs | ~ 974.67 μs | ~ 22.28 μs |
| ↳ /C | 205128.31x | 1184878.64x | 3459637.67x | 862869.09x | 19725.39x |
| ↳ /JS | 0.18x | 1.05x | 3.08x | 0.77x | 0.018x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 15.07 μs | ~ 75.64 μs | ~ 1283.86 μs | ~ 103.76 μs | ~ 0.57 μs |
| ↳ /C | 204.06x | 1024.58x | 17389.55x | 1405.42x | 7.78x |
| ↳ /JS | 0.26x | 1.30x | 22.13x | 1.79x | 0.0099x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.31 μs | ~ 5.69 μs | ~ 43.86 μs | ~ 4.32 μs | ~ 0.42 μs |
| ↳ /C | 221.85x | 4109.26x | 31684.63x | 3118.44x | 299.96x |
| ↳ /JS | 0.043x | 0.80x | 6.20x | 0.61x | 0.059x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 1.75 μs | ~ 55.97 μs | ~ 2130.15 μs | ~ 79.15 μs | ~ 6.65 μs |
| ↳ /C | 3.89x | 124.68x | 4745.76x | 176.33x | 14.81x |
| ↳ /JS | 0.014x | 0.46x | 17.47x | 0.65x | 0.055x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 0.21 μs | ~ 7.62 μs | ~ 384.55 μs | ~ 13.93 μs | ~ 1.26 μs |
| ↳ /C | 2.43x | 87.49x | 4413.97x | 159.91x | 14.43x |
| ↳ /JS | 0.0070x | 0.25x | 12.73x | 0.46x | 0.042x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 0.18 μs | ~ 5.62 μs | ~ 509.78 μs | ~ 18.76 μs | ~ 1.12 μs |
| ↳ /C | 4.05x | 127.13x | 11524.30x | 424.11x | 25.33x |
| ↳ /JS | 0.0022x | 0.070x | 6.32x | 0.23x | 0.014x |
| [Array Indexing](src/Test/ArrayIndexing/ArrayIndexing.purs) | ~ 8139.67 μs | ~ 583895.08 μs | ~ 1595438.29 μs | ~ 3827.17 μs | ~ 3241.50 μs |
| ↳ /C | 1.24x | 88.62x | 242.14x | 0.58x | 0.49x |
| ↳ /JS | 0.32x | 23.08x | 63.05x | 0.15x | 0.13x |
| [JSON to Typed AST](src/Test/JsonTypedAst/JsonTypedAst.purs) | ~ 19733.00 μs | ~ 51932.58 μs | ~ 1509892.75 μs | ~ 98639.50 μs | ~ 15878.58 μs |
| ↳ /C | 1.87x | 4.92x | 142.98x | 9.34x | 1.50x |
| ↳ /JS | 0.24x | 0.64x | 18.72x | 1.22x | 0.20x |
| [JSON Decoding](src/Test/JsonDecoding/JsonDecoding.purs) | ~ 1276.75 μs | ~ 9196.50 μs | ~ 116940.08 μs | ~ 7035.00 μs | ~ 3924.25 μs |
| ↳ /C | 1.94x | 13.98x | 177.80x | 10.70x | 5.97x |
| ↳ /JS | 0.14x | 0.99x | 12.60x | 0.76x | 0.42x |
| **Total WIP** | <strong>~ 29.40 ms</strong> | <strong>~ 646.54 ms</strong> | <strong>~ 3230.78 ms</strong> | <strong>~ 110.72 ms</strong> | <strong>~ 23.08 ms</strong> |
| ↳ /C | <strong>1.65x</strong> | <strong>36.31x</strong> | <strong>181.43x</strong> | <strong>6.22x</strong> | <strong>1.30x</strong> |
| ↳ μ/C | <strong>20591.65x</strong> | <strong>121477.53x</strong> | <strong>373927.99x</strong> | <strong>88807.14x</strong> | <strong>2044.10x</strong> |
| ↳ Q₂/C | <strong>3.97x</strong> | <strong>125.91x</strong> | <strong>8135.03x</strong> | <strong>300.22x</strong> | <strong>14.62x</strong> |
| ↳ Q₄/C | <strong>205128.31x</strong> | <strong>1184878.64x</strong> | <strong>3459637.67x</strong> | <strong>862869.09x</strong> | <strong>19725.39x</strong> |
| ↳ G/C | <strong>29.80x</strong> | <strong>502.44x</strong> | <strong>7538.78x</strong> | <strong>373.07x</strong> | <strong>28.70x</strong> |
| ↳ /JS | <strong>0.25x</strong> | <strong>5.53x</strong> | <strong>27.66x</strong> | <strong>0.95x</strong> | <strong>0.20x</strong> |
| ↳ μ/JS | <strong>0.13x</strong> | <strong>3.19x</strong> | <strong>19.06x</strong> | <strong>0.93x</strong> | <strong>0.099x</strong> |
| ↳ Q₂/JS | <strong>0.092x</strong> | <strong>0.90x</strong> | <strong>15.10x</strong> | <strong>0.70x</strong> | <strong>0.051x</strong> |
| ↳ Q₄/JS | <strong>0.32x</strong> | <strong>23.08x</strong> | <strong>63.05x</strong> | <strong>2.69x</strong> | <strong>0.42x</strong> |
| ↳ G/JS | <strong>0.055x</strong> | <strong>0.92x</strong> | <strong>13.80x</strong> | <strong>0.68x</strong> | <strong>0.053x</strong> |
| **Total** | <strong>~ 31.05 ms</strong> | <strong>~ 654.37 ms</strong> | <strong>~ 4727.30 ms</strong> | <strong>~ 195.37 ms</strong> | <strong>~ 34.17 ms</strong> |
| ↳ /C | <strong>1.44x</strong> | <strong>30.44x</strong> | <strong>219.89x</strong> | <strong>9.09x</strong> | <strong>1.59x</strong> |
| ↳ μ/C | <strong>5148.81x</strong> | <strong>67814.67x</strong> | <strong>13398725.17x</strong> | <strong>1331438.40x</strong> | <strong>48947.79x</strong> |
| ↳ Q₂/C | <strong>1.31x</strong> | <strong>14.78x</strong> | <strong>2380.64x</strong> | <strong>172.45x</strong> | <strong>6.87x</strong> |
| ↳ Q₄/C | <strong>205128.31x</strong> | <strong>1184878.64x</strong> | <strong>455726460.47x</strong> | <strong>40483142.40x</strong> | <strong>1924411.00x</strong> |
| ↳ G/C | <strong>2.24x</strong> | <strong>33.30x</strong> | <strong>5582.75x</strong> | <strong>241.57x</strong> | <strong>30.19x</strong> |
| ↳ /JS | <strong>0.14x</strong> | <strong>2.85x</strong> | <strong>20.55x</strong> | <strong>0.85x</strong> | <strong>0.15x</strong> |
| ↳ μ/JS | <strong>0.086x</strong> | <strong>1.29x</strong> | <strong>27.24x</strong> | <strong>9.41x</strong> | <strong>8.75x</strong> |
| ↳ Q₂/JS | <strong>0.014x</strong> | <strong>0.15x</strong> | <strong>12.66x</strong> | <strong>0.63x</strong> | <strong>0.057x</strong> |
| ↳ Q₄/JS | <strong>0.86x</strong> | <strong>23.08x</strong> | <strong>345.65x</strong> | <strong>343.11x</strong> | <strong>344.58x</strong> |
| ↳ G/JS | <strong>0.0043x</strong> | <strong>0.064x</strong> | <strong>10.69x</strong> | <strong>0.46x</strong> | <strong>0.058x</strong> |

#### Scheme

Scheme Benchmark        | Hand-written PureScript<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written FP-style Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 | Hand-written Scheme FFI<br>↓<br>[pscm](https://github.com/purescm/purescm)<br>↓<br>Scheme<br>↓<br>Chez-O3 |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.074 μs | ~ 0.076 μs | ~ 0.069 μs |
| ↳ /C | 0.76x | 0.78x | 0.71x |
| ↳ /JS | 0.29x | 0.29x | 0.27x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.15 μs | ~ 0.15 μs | ~ 0.14 μs |
| ↳ /C | 1.89x | 1.89x | 1.85x |
| ↳ /JS | 0.49x | 0.49x | 0.48x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 6.35 μs | ~ 3.76 μs | ~ 2.84 μs |
| ↳ /C | 5448.91x | 3224.57x | 2432.51x |
| ↳ /JS | 0.74x | 0.44x | 0.33x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 277.12 μs | ~ 322.34 μs | ~ 329.88 μs |
| ↳ /C | 8.55x | 9.94x | 10.18x |
| ↳ /JS | 1.48x | 1.72x | 1.76x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 228.44 μs | ~ 48.07 μs | ~ 43.57 μs |
| ↳ /C | 70.64x | 14.87x | 13.48x |
| ↳ /JS | 3.07x | 0.65x | 0.59x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.99 μs | ~ 9.07 μs | ~ 9.14 μs |
| ↳ /C | 1.04x | 0.56x | 0.56x |
| ↳ /JS | 0.25x | 0.14x | 0.14x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 237.12 μs | ~ 235.53 μs | ~ 62.26 μs |
| ↳ /C | 209925.39x | 208514.45x | 55116.48x |
| ↳ /JS | 0.19x | 0.19x | 0.049x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 37.49 μs | ~ 22.68 μs | ~ 0.80 μs |
| ↳ /C | 507.85x | 307.19x | 10.82x |
| ↳ /JS | 0.65x | 0.39x | 0.014x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 22551.00 μs | ~ 15428.00 μs | ~ 11970.00 μs |
| ↳ /C | 6.23x | 4.26x | 3.30x |
| ↳ /JS | 0.27x | 0.18x | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 17812.00 μs | ~ 27373.00 μs | ~ 6730.50 μs |
| ↳ /C | 15395971.74x | 23660113.09x | 5817571.74x |
| ↳ /JS | 1.50x | 2.31x | 0.57x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 1.70 μs | ~ 2.02 μs | ~ 0.84 μs |
| ↳ /C | 1498.65x | 1783.70x | 740.40x |
| ↳ /JS | 0.036x | 0.043x | 0.018x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 2989.50 μs | ~ 1964.00 μs | ~ 0.61 μs |
| ↳ /C | 2633663.14x | 1730227.26x | 537.16x |
| ↳ /JS | 0.21x | 0.14x | 0.000043x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 7.29 μs | ~ 5.04 μs | ~ 2.79 μs |
| ↳ /C | 5267.45x | 3642.03x | 2017.84x |
| ↳ /JS | 1.03x | 0.71x | 0.39x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs | ~ 0.015 μs | ~ 0.0010 μs |
| ↳ /C | 0.86x | 12.60x | 0.86x |
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
| ↳ /C | <strong>12.02x</strong> | <strong>12.36x</strong> | <strong>5.21x</strong> |
| ↳ μ/C | <strong>1303740.93x</strong> | <strong>1829132.66x</strong> | <strong>419889.85x</strong> |
| ↳ Q₂/C | <strong>289.25x</strong> | <strong>161.03x</strong> | <strong>12.15x</strong> |
| ↳ Q₄/C | <strong>15395971.74x</strong> | <strong>23660113.09x</strong> | <strong>5817571.74x</strong> |
| ↳ G/C | <strong>370.18x</strong> | <strong>346.50x</strong> | <strong>89.02x</strong> |
| ↳ /JS | <strong>0.39x</strong> | <strong>0.41x</strong> | <strong>0.17x</strong> |
| ↳ μ/JS | <strong>0.73x</strong> | <strong>0.58x</strong> | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>0.39x</strong> | <strong>0.40x</strong> | <strong>0.20x</strong> |
| ↳ Q₄/JS | <strong>3.07x</strong> | <strong>2.31x</strong> | <strong>1.76x</strong> |
| ↳ G/JS | <strong>0.38x</strong> | <strong>0.36x</strong> | <strong>0.091x</strong> |

#### Erlang

Erlang Benchmark        | Hand-written PureScript<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written FP-style Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT | Hand-written imperative Erlang FFI<br>↓<br>[purerl](https://github.com/purerl/purerl)<br>↓<br>Erlang<br>↓<br>BEAM JIT |
----------------------- | ----------------| -------------------------- | -------------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.11 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.15x | 1.19x | 1.04x |
| ↳ /JS | 0.43x | 0.44x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.21 μs | ~ 0.23 μs | ~ 0.24 μs |
| ↳ /C | 2.67x | 2.96x | 3.10x |
| ↳ /JS | 0.69x | 0.76x | 0.80x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 37.32 μs | ~ 5.22 μs | ~ 0.86 μs |
| ↳ /C | 32013.14x | 4479.43x | 736.80x |
| ↳ /JS | 4.32x | 0.60x | 0.099x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 1341.30 μs | ~ 103.95 μs | ~ 104.62 μs |
| ↳ /C | 41.38x | 3.21x | 3.23x |
| ↳ /JS | 7.17x | 0.56x | 0.56x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 404.77 μs | ~ 1276.40 μs | ~ 15.46 μs |
| ↳ /C | 125.17x | 394.72x | 4.78x |
| ↳ /JS | 5.44x | 17.17x | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.24 μs | ~ 18.81 μs | ~ 19.68 μs |
| ↳ /C | 1.18x | 1.15x | 1.21x |
| ↳ /JS | 0.29x | 0.28x | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 520.08 μs | ~ 418.19 μs | ~ 0.012 μs |
| ↳ /C | 460424.42x | 370222.43x | 10.54x |
| ↳ /JS | 0.41x | 0.33x | 0.0000094x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 148.04 μs | ~ 30.95 μs | ~ 36.68 μs |
| ↳ /C | 2005.09x | 419.25x | 496.87x |
| ↳ /JS | 2.55x | 0.53x | 0.63x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16911.71 μs | ~ 14282.54 μs | ~ 14426.50 μs |
| ↳ /C | 4.67x | 3.94x | 3.98x |
| ↳ /JS | 0.20x | 0.17x | 0.17x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 66666.38 μs | ~ 52475.29 μs | ~ 0.0097 μs |
| ↳ /C | 57623715.78x | 45357517.25x | 8.41x |
| ↳ /JS | 5.63x | 4.43x | 0.00000082x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 20.41 μs | ~ 14.55 μs | ~ 7.61 μs |
| ↳ /C | 17985.77x | 12819.06x | 6709.61x |
| ↳ /JS | 0.43x | 0.31x | 0.16x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8739.06 μs | ~ 6780.02 μs | ~ 0.010 μs |
| ↳ /C | 7698861.61x | 5973002.64x | 8.84x |
| ↳ /JS | 0.62x | 0.48x | 0.00000071x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 27.82 μs | ~ 16.21 μs | ~ 1.04 μs |
| ↳ /C | 20100.87x | 11707.92x | 753.97x |
| ↳ /JS | 3.93x | 2.29x | 0.15x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.040 μs | ~ 0.044 μs | ~ 0.0095 μs |
| ↳ /C | 34.41x | 38.09x | 8.23x |
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
| ↳ /C | <strong>25.81x</strong> | <strong>20.52x</strong> | <strong>3.98x</strong> |
| ↳ μ/C | <strong>4703951.24x</strong> | <strong>3695043.80x</strong> | <strong>625.04x</strong> |
| ↳ Q₂/C | <strong>1065.13x</strong> | <strong>406.99x</strong> | <strong>8.32x</strong> |
| ↳ Q₄/C | <strong>57623715.78x</strong> | <strong>45357517.25x</strong> | <strong>6709.61x</strong> |
| ↳ G/C | <strong>1209.23x</strong> | <strong>760.73x</strong> | <strong>20.81x</strong> |
| ↳ /JS | <strong>0.85x</strong> | <strong>0.67x</strong> | <strong>0.13x</strong> |
| ↳ μ/JS | <strong>2.37x</strong> | <strong>2.11x</strong> | <strong>0.27x</strong> |
| ↳ Q₂/JS | <strong>0.91x</strong> | <strong>0.54x</strong> | <strong>0.19x</strong> |
| ↳ Q₄/JS | <strong>7.17x</strong> | <strong>17.17x</strong> | <strong>0.80x</strong> |
| ↳ G/JS | <strong>1.24x</strong> | <strong>0.78x</strong> | <strong>0.021x</strong> |

#### PHP

PHP Benchmark           | Hand-written PureScript<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT | Hand-written FP-style PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT | Hand-written imperative PHP FFI<br>↓<br>[phpurs](https://github.com/0x000000000000000000001/phpurs)<br>↓<br>PHP<br>↓<br>Zend JIT |
----------------------- | ------------------------- | ----------------------- | ----------------------------- |
|  | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |  |  |
|  | <div align="center">**WIP**<br>↓</div> |  |  |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.13 μs | ~ 3.19 μs | ~ 2.11 μs |
| ↳ /C | 22.11x | 33.12x | 21.90x |
| ↳ /JS | 8.27x | 12.39x | 8.20x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 1.08 μs | ~ 2.26 μs | ~ 1.27 μs |
| ↳ /C | 13.82x | 28.91x | 16.25x |
| ↳ /JS | 3.56x | 7.45x | 4.18x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 137.53 μs | ~ 174.47 μs | ~ 0.63 μs |
| ↳ /C | 117975.75x | 149663.56x | 540.43x |
| ↳ /JS | 15.93x | 20.20x | 0.073x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 72.18 μs | ~ 49.87 μs | ~ 50.30 μs |
| ↳ /C | 2.23x | 1.54x | 1.55x |
| ↳ /JS | 0.39x | 0.27x | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1413.71 μs | ~ 1676.49 μs | ~ 491.74 μs |
| ↳ /C | 437.19x | 518.45x | 152.07x |
| ↳ /JS | 19.02x | 22.55x | 6.61x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 45.89 μs | ~ 100.16 μs | ~ 52.48 μs |
| ↳ /C | 2.81x | 6.14x | 3.22x |
| ↳ /JS | 0.69x | 1.50x | 0.79x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1941.41 μs | ~ 6824.38 μs | ~ 32.84 μs |
| ↳ /C | 1718719.01x | 6041584.04x | 29073.06x |
| ↳ /JS | 1.53x | 5.37x | 0.026x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 698.16 μs | ~ 1022.38 μs | ~ 1.89 μs |
| ↳ /C | 9456.36x | 13847.82x | 25.60x |
| ↳ /JS | 12.04x | 17.62x | 0.033x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 53877.92 μs | ~ 509843.88 μs | ~ 113830.46 μs |
| ↳ /C | 14.87x | 140.74x | 31.42x |
| ↳ /JS | 0.64x | 6.05x | 1.35x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 6237.63 μs | ~ 432965.96 μs | ~ 87761.96 μs |
| ↳ /C | 5391554.86x | 374238248.60x | 75857885.47x |
| ↳ /JS | 0.53x | 36.58x | 7.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 467.12 μs | ~ 560.06 μs | ~ 0.48 μs |
| ↳ /C | 411627.29x | 493526.25x | 422.98x |
| ↳ /JS | 9.87x | 11.83x | 0.010x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 791.27 μs | ~ 111183.96 μs | ~ 326.09 μs |
| ↳ /C | 697086.01x | 97949856.89x | 287275.87x |
| ↳ /JS | 0.056x | 7.91x | 0.023x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 89.94 μs | ~ 26.94 μs | ~ 0.80 μs |
| ↳ /C | 64973.44x | 19461.69x | 577.93x |
| ↳ /JS | 12.72x | 3.81x | 0.11x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.84 μs | ~ 0.85 μs | ~ 0.080 μs |
| ↳ /C | 726.32x | 734.97x | 69.17x |
| ↳ /JS | 23.69x | 23.98x | 2.26x |
| **Total WIP** | <strong>65.78 ms</strong> | <strong>1064.43 ms</strong> | <strong>202.55 ms</strong> |
| ↳ /C | <strong>17.90x</strong> | <strong>289.66x</strong> | <strong>55.12x</strong> |
| ↳ μ/C | <strong>600900.86x</strong> | <strong>34207689.48x</strong> | <strong>5441149.78x</strong> |
| ↳ Q₂/C | <strong>5091.34x</strong> | <strong>7291.40x</strong> | <strong>110.62x</strong> |
| ↳ Q₄/C | <strong>5391554.86x</strong> | <strong>374238248.60x</strong> | <strong>75857885.47x</strong> |
| ↳ G/C | <strong>2791.27x</strong> | <strong>7598.57x</strong> | <strong>378.94x</strong> |
| ↳ /JS | <strong>0.59x</strong> | <strong>9.51x</strong> | <strong>1.81x</strong> |
| ↳ μ/JS | <strong>7.78x</strong> | <strong>12.68x</strong> | <strong>2.24x</strong> |
| ↳ Q₂/JS | <strong>5.92x</strong> | <strong>9.87x</strong> | <strong>0.53x</strong> |
| ↳ Q₄/JS | <strong>23.69x</strong> | <strong>36.58x</strong> | <strong>8.20x</strong> |
| ↳ G/JS | <strong>2.86x</strong> | <strong>7.79x</strong> | <strong>0.39x</strong> |
|  | <div align="center">**Backlog**<br>↓</div> |  |  |
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
|  | <div align="center">**Stable**<br>↓<br>Coming soon...</div> | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |  |  |  |
|  | <div align="center">**WIP**<br>↓</div> | <div align="center">**WIP**<br>↓</div> |  |  |  |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.24 μs | ~ 3.68 μs | ~ 0.23 μs | ~ 0.22 μs | ~ 0.23 μs |
| ↳ /C | 2.49x | 38.20x | 2.39x | 2.28x | 2.39x |
| ↳ /JS | 0.93x | 14.30x | 0.89x | 0.85x | 0.89x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.090 μs | ~ 11.78 μs | ~ 0.090 μs | ~ 0.12 μs | ~ 0.10 μs |
| ↳ /C | 1.15x | 150.71x | 1.15x | 1.54x | 1.28x |
| ↳ /JS | 0.30x | 38.82x | 0.30x | 0.40x | 0.33x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.11 μs | ~ 459.33 μs | ~ 17.07 μs | ~ 11.05 μs | ~ 0.51 μs |
| ↳ /C | 94.36x | 394021.69x | 14642.96x | 9478.89x | 437.49x |
| ↳ /JS | 0.013x | 53.19x | 1.98x | 1.28x | 0.059x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 29.88 μs | ~ 23547.50 μs | ~ 10.93 μs | ~ 34.70 μs | ~ 34.63 μs |
| ↳ /C | 0.92x | 726.42x | 0.34x | 1.07x | 1.07x |
| ↳ /JS | 0.16x | 125.86x | 0.058x | 0.19x | 0.19x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.14 μs | ~ 8728.98 μs | ~ 114.96 μs | ~ 107.97 μs | ~ 3.47 μs |
| ↳ /C | 0.97x | 2699.43x | 35.55x | 33.39x | 1.07x |
| ↳ /JS | 0.042x | 117.42x | 1.55x | 1.45x | 0.047x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.45 μs | ~ 625.00 μs | ~ 15.60 μs | ~ 18.59 μs | ~ 18.36 μs |
| ↳ /C | 1.01x | 38.33x | 0.96x | 1.14x | 1.13x |
| ↳ /JS | 0.25x | 9.36x | 0.23x | 0.28x | 0.27x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 159.52 μs | ~ 11792.17 μs | ~ 2249.51 μs | ~ 1033.78 μs | ~ 0.010 μs |
| ↳ /C | 141222.13x | 10439539.72x | 1991478.16x | 915199.44x | 8.85x |
| ↳ /JS | 0.13x | 9.29x | 1.77x | 0.81x | 0.0000079x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 144.00 μs | ~ 2194.96 μs | ~ 120.43 μs | ~ 78.88 μs | ~ 0.78 μs |
| ↳ /C | 1950.44x | 29730.06x | 1631.19x | 1068.41x | 10.56x |
| ↳ /JS | 2.48x | 37.84x | 2.08x | 1.36x | 0.013x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8136.98 μs | ~ 659207.04 μs | ~ 50862.42 μs | ~ 28737.38 μs | ~ 28247.96 μs |
| ↳ /C | 2.25x | 181.98x | 14.04x | 7.93x | 7.80x |
| ↳ /JS | 0.097x | 7.82x | 0.60x | 0.34x | 0.34x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.030 μs | ~ 1588595.46 μs | ~ 7471.33 μs | ~ 0.010 μs | ~ 0.010 μs |
| ↳ /C | 25.93x | 1373117606.48x | 6457915.20x | 8.64x | 8.64x |
| ↳ /JS | 0.0000025x | 134.21x | 0.63x | 0.00000084x | 0.00000084x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 44.57 μs | ~ 826.12 μs | ~ 73.12 μs | ~ 27.36 μs | ~ 0.010 μs |
| ↳ /C | 39275.19x | 727978.97x | 64433.52x | 24109.70x | 8.81x |
| ↳ /JS | 0.94x | 17.45x | 1.54x | 0.58x | 0.00021x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.010 μs | ~ 159056.38 μs | ~ 38342.83 μs | ~ 23675.83 μs | ~ 0.010 μs |
| ↳ /C | 8.81x | 140124075.97x | 33778925.59x | 20857722.28x | 8.81x |
| ↳ /JS | 0.00000071x | 11.31x | 2.73x | 1.68x | 0.00000071x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.030 μs | ~ 298.70 μs | ~ 3.28 μs | ~ 0.58 μs | ~ 0.25 μs |
| ↳ /C | 21.67x | 215783.47x | 2369.50x | 419.00x | 180.60x |
| ↳ /JS | 0.0042x | 42.24x | 0.46x | 0.082x | 0.035x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.070 μs | ~ 0.95 μs | ~ 0.10 μs | ~ 0.020 μs | ~ 0.010 μs |
| ↳ /C | 60.53x | 821.44x | 86.47x | 17.29x | 8.65x |
| ↳ /JS | 1.97x | 26.80x | 2.82x | 0.56x | 0.28x |
| **Total WIP** | <strong>8.54 ms</strong> | <strong>2455.35 ms</strong> |  | <strong>53.73 ms</strong> | <strong>28.31 ms</strong> |
| ↳ /C | <strong>2.32x</strong> | <strong>668.17x</strong> |  | <strong>14.62x</strong> | <strong>7.70x</strong> |
| ↳ μ/C | <strong>13047.70x</strong> | <strong>108932385.20x</strong> |  | <strong>1557719.36x</strong> | <strong>49.08x</strong> |
| ↳ Q₂/C | <strong>15.24x</strong> | <strong>16214.74x</strong> |  | <strong>25.34x</strong> | <strong>8.65x</strong> |
| ↳ Q₄/C | <strong>141222.13x</strong> | <strong>1373117606.48x</strong> |  | <strong>20857722.28x</strong> | <strong>437.49x</strong> |
| ↳ G/C | <strong>32.95x</strong> | <strong>28550.30x</strong> |  | <strong>210.30x</strong> | <strong>7.34x</strong> |
| ↳ /JS | <strong>0.076x</strong> | <strong>21.95x</strong> |  | <strong>0.48x</strong> | <strong>0.25x</strong> |
| ↳ μ/JS | <strong>0.52x</strong> | <strong>46.13x</strong> |  | <strong>0.70x</strong> | <strong>0.18x</strong> |
| ↳ Q₂/JS | <strong>0.14x</strong> | <strong>32.32x</strong> |  | <strong>0.57x</strong> | <strong>0.053x</strong> |
| ↳ Q₄/JS | <strong>2.48x</strong> | <strong>134.21x</strong> |  | <strong>1.68x</strong> | <strong>0.89x</strong> |
| ↳ G/JS | <strong>0.034x</strong> | <strong>29.26x</strong> |  | <strong>0.22x</strong> | <strong>0.0075x</strong> |
|  | <div align="center">**Backlog**<br>↓</div> | <div align="center">**Backlog**<br>↓</div> |  |  |  |
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
| ↳ /C |  |  | <strong>27.02x</strong> |  |  |
| ↳ μ/C |  |  | <strong>3022252.64x</strong> |  |  |
| ↳ Q₂/C |  |  | <strong>858.83x</strong> |  |  |
| ↳ Q₄/C |  |  | <strong>33778925.59x</strong> |  |  |
| ↳ G/C |  |  | <strong>818.56x</strong> |  |  |
| ↳ /JS |  |  | <strong>0.89x</strong> |  |  |
| ↳ μ/JS |  |  | <strong>1.26x</strong> |  |  |
| ↳ Q₂/JS |  |  | <strong>1.22x</strong> |  |  |
| ↳ Q₄/JS |  |  | <strong>2.82x</strong> |  |  |
| ↳ G/JS |  |  | <strong>0.84x</strong> |  |  |

#### F#/C#

F#/C# Benchmark         | Hand-written PureScript<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release | Hand-written FP-style F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release | Hand-written F#/C# FFI<br>↓<br>[sharpurs](https://github.com/0x000000000000000000001/sharpurs)<br>↓<br>F#/C#<br>↓<br>dotnet -Release |
----------------------- | ------------------------- | ------------------------- | ------------------------------- |
|  | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |  |  |
|  | <div align="center">**WIP**<br>↓</div> |  |  |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 1.12 μs | ~ 0.50 μs | ~ 0.50 μs |
| ↳ /C | 11.63x | 5.19x | 5.19x |
| ↳ /JS | 4.35x | 1.94x | 1.94x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.36 μs | ~ 0.12 μs | ~ 0.11 μs |
| ↳ /C | 4.61x | 1.54x | 1.41x |
| ↳ /JS | 1.19x | 0.40x | 0.36x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 560.47 μs | ~ 16.05 μs | ~ 3.05 μs |
| ↳ /C | 480781.44x | 13767.98x | 2616.35x |
| ↳ /JS | 64.90x | 1.86x | 0.35x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 57.45 μs | ~ 37.60 μs | ~ 49.49 μs |
| ↳ /C | 1.77x | 1.16x | 1.53x |
| ↳ /JS | 0.31x | 0.20x | 0.26x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 2545.27 μs | ~ 87.33 μs | ~ 7.21 μs |
| ↳ /C | 787.12x | 27.01x | 2.23x |
| ↳ /JS | 34.24x | 1.17x | 0.097x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 15.15 μs | ~ 7.70 μs | ~ 7.66 μs |
| ↳ /C | 0.93x | 0.47x | 0.47x |
| ↳ /JS | 0.23x | 0.12x | 0.11x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 1678.49 μs | ~ 342.79 μs | ~ 25.24 μs |
| ↳ /C | 1485957.46x | 303470.00x | 22344.83x |
| ↳ /JS | 1.32x | 0.27x | 0.020x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 525.55 μs | ~ 45.26 μs | ~ 1.80 μs |
| ↳ /C | 7118.41x | 613.03x | 24.38x |
| ↳ /JS | 9.06x | 0.78x | 0.031x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 19071.54 μs | ~ 36659.92 μs | ~ 9499.96 μs |
| ↳ /C | 5.26x | 10.12x | 2.62x |
| ↳ /JS | 0.23x | 0.44x | 0.11x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 2523.93 μs | ~ 23242.04 μs | ~ 2254.54 μs |
| ↳ /C | 2181582.92x | 20089478.50x | 1948733.11x |
| ↳ /JS | 0.21x | 1.96x | 0.19x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 169.78 μs | ~ 78.22 μs | ~ 3.21 μs |
| ↳ /C | 149610.55x | 68927.66x | 2828.66x |
| ↳ /JS | 3.59x | 1.65x | 0.068x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 6576.73 μs | ~ 9665.58 μs | ~ 235.04 μs |
| ↳ /C | 5793909.14x | 8515096.76x | 207063.45x |
| ↳ /JS | 0.47x | 0.69x | 0.017x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 62.07 μs | ~ 29.54 μs | ~ 15.63 μs |
| ↳ /C | 44839.91x | 21339.95x | 11291.25x |
| ↳ /JS | 8.78x | 4.18x | 2.21x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.24 μs | ~ 0.060 μs | ~ 0.020 μs |
| ↳ /C | 207.52x | 51.88x | 17.29x |
| ↳ /JS | 6.77x | 1.69x | 0.56x |
| **Total WIP** | <strong>33.79 ms</strong> | <strong>70.21 ms</strong> | <strong>12.10 ms</strong> |
| ↳ /C | <strong>9.19x</strong> | <strong>19.11x</strong> | <strong>3.29x</strong> |
| ↳ μ/C | <strong>724629.91x</strong> | <strong>2072342.23x</strong> | <strong>156780.91x</strong> |
| ↳ Q₂/C | <strong>3952.77x</strong> | <strong>332.46x</strong> | <strong>20.84x</strong> |
| ↳ Q₄/C | <strong>5793909.14x</strong> | <strong>20089478.50x</strong> | <strong>1948733.11x</strong> |
| ↳ G/C | <strong>2101.20x</strong> | <strong>794.44x</strong> | <strong>163.47x</strong> |
| ↳ /JS | <strong>0.30x</strong> | <strong>0.63x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>9.69x</strong> | <strong>1.24x</strong> | <strong>0.45x</strong> |
| ↳ Q₂/JS | <strong>2.45x</strong> | <strong>0.98x</strong> | <strong>0.15x</strong> |
| ↳ Q₄/JS | <strong>64.90x</strong> | <strong>4.18x</strong> | <strong>2.21x</strong> |
| ↳ G/JS | <strong>2.15x</strong> | <strong>0.81x</strong> | <strong>0.17x</strong> |
|  | <div align="center">**Backlog**<br>↓</div> |  |  |
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
|  | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |  |  |
|  | <div align="center">**WIP**<br>↓</div> |  |  |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.14 μs | ~ 0.11 μs | ~ 0.10 μs |
| ↳ /C | 1.49x | 1.10x | 1.05x |
| ↳ /JS | 0.56x | 0.41x | 0.39x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.26 μs | ~ 0.099 μs | ~ 0.077 μs |
| ↳ /C | 3.31x | 1.26x | 0.98x |
| ↳ /JS | 0.85x | 0.33x | 0.25x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 5.03 μs | ~ 4.18 μs | ~ 0.57 μs |
| ↳ /C | 4314.99x | 3588.36x | 486.48x |
| ↳ /JS | 0.58x | 0.48x | 0.066x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 45.64 μs | ~ 38.54 μs | ~ 39.83 μs |
| ↳ /C | 1.41x | 1.19x | 1.23x |
| ↳ /JS | 0.24x | 0.21x | 0.21x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 73.28 μs | ~ 38.80 μs | ~ 4.18 μs |
| ↳ /C | 22.66x | 12.00x | 1.29x |
| ↳ /JS | 0.99x | 0.52x | 0.056x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 18.18 μs | ~ 7.60 μs | ~ 5.80 μs |
| ↳ /C | 1.11x | 0.47x | 0.36x |
| ↳ /JS | 0.27x | 0.11x | 0.087x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 310.51 μs | ~ 375.72 μs | ~ 0.0031 μs |
| ↳ /C | 274891.93x | 332619.33x | 2.76x |
| ↳ /JS | 0.24x | 0.30x | 0.0000025x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.29 μs | ~ 18.40 μs | ~ 0.54 μs |
| ↳ /C | 288.35x | 249.22x | 7.34x |
| ↳ /JS | 0.37x | 0.32x | 0.0093x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 16788.38 μs | ~ 12065.75 μs | ~ 11828.17 μs |
| ↳ /C | 4.63x | 3.33x | 3.27x |
| ↳ /JS | 0.20x | 0.14x | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0054 μs | ~ 14583.08 μs | ~ 0.0034 μs |
| ↳ /C | 4.68x | 12605026.60x | 2.91x |
| ↳ /JS | 0.00000046x | 1.23x | 0.00000028x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 6.26 μs | ~ 17.25 μs | ~ 0.0033 μs |
| ↳ /C | 5513.40x | 15199.88x | 2.87x |
| ↳ /JS | 0.13x | 0.36x | 0.000069x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5.24 μs | ~ 8062.00 μs | ~ 0.037 μs |
| ↳ /C | 4618.45x | 7102389.11x | 32.51x |
| ↳ /JS | 0.00037x | 0.57x | 0.0000026x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 2.96 μs | ~ 3.06 μs | ~ 0.50 μs |
| ↳ /C | 2136.70x | 2212.91x | 359.36x |
| ↳ /JS | 0.42x | 0.43x | 0.070x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.013 μs | ~ 0.0062 μs | ~ 0.0023 μs |
| ↳ /C | 11.00x | 5.33x | 2.01x |
| ↳ /JS | 0.36x | 0.17x | 0.066x |
| **Total WIP** | <strong>~ 17.28 ms</strong> | <strong>~ 35.21 ms</strong> | <strong>~ 11.88 ms</strong> |
| ↳ /C | <strong>4.70x</strong> | <strong>9.58x</strong> | <strong>3.23x</strong> |
| ↳ μ/C | <strong>20843.86x</strong> | <strong>1432950.72x</strong> | <strong>64.60x</strong> |
| ↳ Q₂/C | <strong>16.83x</strong> | <strong>130.61x</strong> | <strong>2.82x</strong> |
| ↳ Q₄/C | <strong>274891.93x</strong> | <strong>12605026.60x</strong> | <strong>486.48x</strong> |
| ↳ G/C | <strong>83.33x</strong> | <strong>325.06x</strong> | <strong>4.82x</strong> |
| ↳ /JS | <strong>0.15x</strong> | <strong>0.31x</strong> | <strong>0.11x</strong> |
| ↳ μ/JS | <strong>0.37x</strong> | <strong>0.40x</strong> | <strong>0.097x</strong> |
| ↳ Q₂/JS | <strong>0.32x</strong> | <strong>0.34x</strong> | <strong>0.066x</strong> |
| ↳ Q₄/JS | <strong>0.99x</strong> | <strong>1.23x</strong> | <strong>0.39x</strong> |
| ↳ G/JS | <strong>0.085x</strong> | <strong>0.33x</strong> | <strong>0.0049x</strong> |
|  | <div align="center">**Backlog**<br>↓</div> |  |  |
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
| ↳ /C | 2.27x |
| ↳ /JS | 0.85x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.58x |
| ↳ /JS | 0.41x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 4.28 μs |
| ↳ /C | 3667.51x |
| ↳ /JS | 0.50x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 80.51 μs |
| ↳ /C | 2.48x |
| ↳ /JS | 0.43x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 158.94 μs |
| ↳ /C | 49.15x |
| ↳ /JS | 2.14x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 19.38 μs |
| ↳ /C | 1.19x |
| ↳ /JS | 0.29x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 758.69 μs |
| ↳ /C | 671661.64x |
| ↳ /JS | 0.60x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 12.17 μs |
| ↳ /C | 164.81x |
| ↳ /JS | 0.21x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 8746.50 μs |
| ↳ /C | 2.41x |
| ↳ /JS | 0.10x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 35687.00 μs |
| ↳ /C | 30846398.13x |
| ↳ /JS | 3.01x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 5.21 μs |
| ↳ /C | 4587.59x |
| ↳ /JS | 0.11x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 8982.00 μs |
| ↳ /C | 7912882.53x |
| ↳ /JS | 0.64x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 6.58 μs |
| ↳ /C | 4752.45x |
| ↳ /JS | 0.93x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0010 μs |
| ↳ /C | 0.89x |
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
| ↳ /C | <strong>14.82x</strong> |
| ↳ μ/C | <strong>2817441.05x</strong> |
| ↳ Q₂/C | <strong>106.98x</strong> |
| ↳ Q₄/C | <strong>30846398.13x</strong> |
| ↳ G/C | <strong>397.10x</strong> |
| ↳ /JS | <strong>0.49x</strong> |
| ↳ μ/JS | <strong>0.73x</strong> |
| ↳ Q₂/JS | <strong>0.46x</strong> |
| ↳ Q₄/JS | <strong>3.01x</strong> |
| ↳ G/JS | <strong>0.41x</strong> |

#### Haskell

Haskell Benchmark       | Hand-written Haskell<br>↓<br>GHC -O2 |
----------------------- | ------------------------ |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.056 μs |
| ↳ /C | 0.58x |
| ↳ /JS | 0.22x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.19 μs |
| ↳ /C | 2.49x |
| ↳ /JS | 0.64x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 3.58 μs |
| ↳ /C | 3070.22x |
| ↳ /JS | 0.41x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 62.84 μs |
| ↳ /C | 1.94x |
| ↳ /JS | 0.34x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 7.68 μs |
| ↳ /C | 2.37x |
| ↳ /JS | 0.10x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 6.75 μs |
| ↳ /C | 0.41x |
| ↳ /JS | 0.10x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 143.90 μs |
| ↳ /C | 127392.45x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 21.58 μs |
| ↳ /C | 292.32x |
| ↳ /JS | 0.37x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 11559.00 μs |
| ↳ /C | 3.19x |
| ↳ /JS | 0.14x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 4883.00 μs |
| ↳ /C | 4220667.53x |
| ↳ /JS | 0.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.016 μs |
| ↳ /C | 13.77x |
| ↳ /JS | 0.00033x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.46 μs |
| ↳ /C | 401.80x |
| ↳ /JS | 0.000032x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 4.33 μs |
| ↳ /C | 3125.62x |
| ↳ /JS | 0.61x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0041 μs |
| ↳ /C | 3.52x |
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
| ↳ /C | <strong>4.54x</strong> |
| ↳ μ/C | <strong>311069.87x</strong> |
| ↳ Q₂/C | <strong>8.65x</strong> |
| ↳ Q₄/C | <strong>4220667.53x</strong> |
| ↳ G/C | <strong>76.99x</strong> |
| ↳ /JS | <strong>0.15x</strong> |
| ↳ μ/JS | <strong>0.26x</strong> |
| ↳ Q₂/JS | <strong>0.18x</strong> |
| ↳ Q₄/JS | <strong>0.64x</strong> |
| ↳ G/JS | <strong>0.079x</strong> |

#### OCaml

OCaml Benchmark         | Hand-written OCaml<br>↓<br>ocamlopt -O3 |
----------------------- | --------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.049 μs |
| ↳ /C | 0.51x |
| ↳ /JS | 0.19x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.12 μs |
| ↳ /C | 1.59x |
| ↳ /JS | 0.41x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 1.98 μs |
| ↳ /C | 1695.53x |
| ↳ /JS | 0.23x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 49.82 μs |
| ↳ /C | 1.54x |
| ↳ /JS | 0.27x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 15.85 μs |
| ↳ /C | 4.90x |
| ↳ /JS | 0.21x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.05 μs |
| ↳ /C | 0.98x |
| ↳ /JS | 0.24x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 140.38 μs |
| ↳ /C | 124273.17x |
| ↳ /JS | 0.11x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 17.94 μs |
| ↳ /C | 242.96x |
| ↳ /JS | 0.31x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 10787.00 μs |
| ↳ /C | 2.98x |
| ↳ /JS | 0.13x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 12448.00 μs |
| ↳ /C | 10759547.28x |
| ↳ /JS | 1.05x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 10.03 μs |
| ↳ /C | 8835.26x |
| ↳ /JS | 0.21x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 5690.00 μs |
| ↳ /C | 5012725.63x |
| ↳ /JS | 0.40x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 5.89 μs |
| ↳ /C | 4257.20x |
| ↳ /JS | 0.83x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0047 μs |
| ↳ /C | 4.05x |
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
| ↳ /C | <strong>7.94x</strong> |
| ↳ μ/C | <strong>1136542.40x</strong> |
| ↳ Q₂/C | <strong>123.93x</strong> |
| ↳ Q₄/C | <strong>10759547.28x</strong> |
| ↳ G/C | <strong>262.97x</strong> |
| ↳ /JS | <strong>0.26x</strong> |
| ↳ μ/JS | <strong>0.34x</strong> |
| ↳ Q₂/JS | <strong>0.23x</strong> |
| ↳ Q₄/JS | <strong>1.05x</strong> |
| ↳ G/JS | <strong>0.27x</strong> |

#### C++

C++ Benchmark           | Hand-written PureScript<br>↓<br>[pscpp](https://github.com/purescript-native/purescript)<br>↓<br>C++<br>↓<br>clang++ -O3 |
----------------------- | ------------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 2.71 μs |
| ↳ /C | 28.15x |
| ↳ /JS | 10.53x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.31 μs |
| ↳ /C | 4.03x |
| ↳ /JS | 1.04x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 154.89 μs |
| ↳ /C | 132868.02x |
| ↳ /JS | 17.94x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 2844.59 μs |
| ↳ /C | 87.75x |
| ↳ /JS | 15.20x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 1453.15 μs |
| ↳ /C | 449.38x |
| ↳ /JS | 19.55x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 221.50 μs |
| ↳ /C | 13.58x |
| ↳ /JS | 3.32x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 2630.26 μs |
| ↳ /C | 2328554.15x |
| ↳ /JS | 2.07x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 974.45 μs |
| ↳ /C | 13198.67x |
| ↳ /JS | 16.80x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 613388.67 μs |
| ↳ /C | 169.33x |
| ↳ /JS | 7.28x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 265261.00 μs |
| ↳ /C | 229280870.17x |
| ↳ /JS | 22.41x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 293.30 μs |
| ↳ /C | 258453.93x |
| ↳ /JS | 6.20x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 36615.58 μs |
| ↳ /C | 32257271.76x |
| ↳ /JS | 2.60x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 36.57 μs |
| ↳ /C | 26420.94x |
| ↳ /JS | 5.17x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.48 μs |
| ↳ /C | 417.08x |
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
| ↳ /C | <strong>251.41x</strong> |
| ↳ μ/C | <strong>18878486.21x</strong> |
| ↳ Q₂/C | <strong>6824.03x</strong> |
| ↳ Q₄/C | <strong>229280870.17x</strong> |
| ↳ G/C | <strong>7134.31x</strong> |
| ↳ /JS | <strong>8.26x</strong> |
| ↳ μ/JS | <strong>10.27x</strong> |
| ↳ Q₂/JS | <strong>8.91x</strong> |
| ↳ Q₄/JS | <strong>22.41x</strong> |
| ↳ G/JS | <strong>7.31x</strong> |

#### C (reference)

C Benchmark             | [Hand-written imperative C](tmp/bench_native.c)<br>↓<br>clang -O3 |
----------------------- | -------------------- |
| [AST Evaluation](src/Test/AstTree/AstTree.purs) | ~ 0.096 μs |
| ↳ /C | — |
| ↳ /JS | 0.37x |
| [Fibonacci](src/Test/Fib/Fib.purs) | ~ 0.078 μs |
| ↳ /C | — |
| ↳ /JS | 0.26x |
| [List Processing](src/Test/ListOps/ListOps.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.00013x |
| [Tail Call Optimization](src/Test/TCO/TCO.purs) | ~ 32.42 μs |
| ↳ /C | — |
| ↳ /JS | 0.17x |
| [Deep Record Updates](src/Test/Records/Records.purs) | ~ 3.23 μs |
| ↳ /C | — |
| ↳ /JS | 0.043x |
| [Ackermann](src/Test/Ackermann/Ackermann.purs) | ~ 16.31 μs |
| ↳ /C | — |
| ↳ /JS | 0.24x |
| [Church Numerals](src/Test/Church/Church.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.00000089x |
| [Prime Sieve](src/Test/Primes/Primes.purs) | ~ 0.074 μs |
| ↳ /C | — |
| ↳ /JS | 0.0013x |
| [Red-Black Tree](src/Test/RBTree/RBTree.purs) | ~ 3622.50 μs |
| ↳ /C | — |
| ↳ /JS | 0.043x |
| [Polymorphism](src/Test/Polymorphism/Polymorphism.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.000000098x |
| [State Monad](src/Test/StateMonad/StateMonad.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.000024x |
| [Lazy Evaluation](src/Test/LazyEvaluation/LazyEvaluation.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.000000081x |
| [Array Processing](src/Test/ArrayOps/ArrayOps.purs) | ~ 0.0014 μs |
| ↳ /C | — |
| ↳ /JS | 0.00020x |
| [RowToList](src/Test/RowToList/RowToList.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.033x |
| [array/map-10000](src/Test/ArrayMap10000/ArrayMap10000.purs) | ~ 0.45 μs |
| ↳ /C | — |
| ↳ /JS | 0.0037x |
| [array/filter-10000](src/Test/ArrayFilter10000/ArrayFilter10000.purs) | ~ 4.63 μs |
| ↳ /C | — |
| ↳ /JS | 0.036x |
| [array/foldl-sum-10000](src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | ~ 0.29 μs |
| ↳ /C | — |
| ↳ /JS | 0.0041x |
| [array/sort-descending-10000](src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | ~ 2.85 μs |
| ↳ /C | — |
| ↳ /JS | 0.0051x |
| [array/eq-neq-last-10000](src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | ~ 0.67 μs |
| ↳ /C | — |
| ↳ /JS | 0.013x |
| [array/append-1000](src/Test/ArrayAppend1000/ArrayAppend1000.purs) | ~ 0.087 μs |
| ↳ /C | — |
| ↳ /JS | 0.0029x |
| [foldable/traverse-maybe-1000](src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | ~ 0.044 μs |
| ↳ /C | — |
| ↳ /JS | 0.00055x |
| [foldable/traverse-maybe-fail-first-1000](src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | ~ 0.0013 μs |
| ↳ /C | — |
| ↳ /JS | 0.000020x |
| [foldable/traverse-maybe-fail-last-1000](src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | ~ 0.34 μs |
| ↳ /C | — |
| ↳ /JS | 0.0050x |
| [string/cu-length-unicode-1000](src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | ~ 0.0011 μs |
| ↳ /C | — |
| ↳ /JS | 0.11x |
| [string/cp-length-unicode-1000](src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | ~ 1.23 μs |
| ↳ /C | — |
| ↳ /JS | 0.033x |
| [string/indexOf-hit-10000](src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | ~ 1.11 μs |
| ↳ /C | — |
| ↳ /JS | 5.79x |
| [string/indexOf-miss-10000](src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | ~ 1.27 μs |
| ↳ /C | — |
| ↳ /JS | 0.020x |
| [string/cp-take-unicode-1500](src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | ~ 0.67 μs |
| ↳ /C | — |
| ↳ /JS | 0.015x |
| [string/split-1000](src/Test/StringSplit1000/StringSplit1000.purs) | ~ 2.38 μs |
| ↳ /C | — |
| ↳ /JS | 0.016x |
| [string/concat-1000](src/Test/StringConcat1000/StringConcat1000.purs) | ~ 0.96 μs |
| ↳ /C | — |
| ↳ /JS | 0.020x |
| [record/set-small-1000](src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.00047x |
| [record/set-medium-1000](src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | ~ 0.0012 μs |
| ↳ /C | — |
| ↳ /JS | 0.00025x |
| [record/set-large-1000](src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | ~ 0.0014 μs |
| ↳ /C | — |
| ↳ /JS | 0.00020x |
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
| **Total** | <strong>21.50 ms</strong> |
| ↳ /C | — |
| ↳ μ/C | — |
| ↳ Q₂/C | — |
| ↳ Q₄/C | — |
| ↳ G/C | — |
| ↳ /JS | <strong>0.093x</strong> |
| ↳ μ/JS | <strong>0.19x</strong> |
| ↳ Q₂/JS | <strong>0.0051x</strong> |
| ↳ Q₄/JS | <strong>5.79x</strong> |
| ↳ G/JS | <strong>0.0019x</strong> |

### Extended benchmark results (I/O, mutability, async)

**Extended tests (`srx/`)** cover features shared by a smaller subset of targets (`./bin/run --x`). For example, `Aff` is supported by the JavaScript, Rust and Go backends, but not the Scheme backend.

#### Extended Results

Benchmark               | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[Arista](https://github.com/aristanetworks/purescript-backend-optimizer)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build | Hand-written PureScript<br>↓<br>[purust](https://github.com/0x000000000000000000001/purust)<br>↓<br>Rust<br>↓<br>rustc -O3<br>+ thin LTO
----------------------- | ------------- | -------------- | --------------- | ---------------
| | | | <div align="center">**Stable**<br>↓<br>Coming soon...</div> | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |
| | | | <div align="center">**WIP**<br>↓</div> | <div align="center">**WIP**<br>↓</div> |
[File I/O](srx/Test/FileOps/FileOps.purs) | ~ 466899.83 μs | ~ 477117.79 μs | ~ 397954.62 μs | ~ 542377.75 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.02x | 0.85x | 1.16x |
[STArray Operations](srx/Test/STArray.purs) | ~ 0.92 μs | ~ 0.54 μs | ~ 0.21 μs | ~ 0.29 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.59x | 0.23x | 0.32x |
[String Operations](srx/Test/StringOps.purs) | ~ 256.71 μs | ~ 232.50 μs | ~ 535.29 μs | ~ 414.58 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 0.91x | 2.09x | 1.61x |
[Aff Operations](srx/Test/AffOperations.purs) | ~ 10125.21 μs | ~ 11581.83 μs | ~ 10676.88 μs | ~ 10903.21 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.14x | 1.05x | 1.08x |
[Parallelism](srx/Test/Parallelism.purs) | ~ 14275052.04 μs | ~ 14278260.46 μs | ~ 628306.62 μs | ~ 478303.67 μs |
↳ /C | — | — | — | — |
↳ /JS | — | 1.00x | 0.044x | 0.030x |
**Total Execution Time** | <strong>~ 14752.33 ms</strong> | <strong>~ 14767.19 ms</strong> | <strong>~ 1037.47 ms</strong> | <strong>~ 1032.00 ms</strong> |
↳ /C | — | — | — | — |
↳ μ/C | — | — | — | — |
↳ Q₂/C | — | — | — | — |
↳ Q₄/C | — | — | — | — |
↳ G/C | — | — | — | — |
↳ /JS | — | <strong>1.00x</strong> | <strong>0.070x</strong> | <strong>0.070x</strong> |
↳ μ/JS | — | <strong>0.93x</strong> | <strong>0.85x</strong> | <strong>0.84x</strong> |
↳ Q₂/JS | — | <strong>1.00x</strong> | <strong>0.85x</strong> | <strong>1.08x</strong> |
↳ Q₄/JS | — | <strong>1.14x</strong> | <strong>2.09x</strong> | <strong>1.61x</strong> |
↳ G/JS | — | <strong>0.91x</strong> | <strong>0.45x</strong> | <strong>0.46x</strong> |

> [!NOTE]
> **Hardware Context**
> Measurements ran on an **Apple M4 Pro with 10 performance cores and 4 efficiency cores**, without explicit CPU affinity. The extended *Parallelism* row compares JavaScript's single-thread scheduler with the Go/Rust multicore runtimes.

### Application benchmark results

#### ... to Go (instead of JS)

App scenario | Hand-written PureScript<br>↓<br>[official](https://github.com/purescript/purescript)<br>↓<br>JS<br>↓<br>V8 JIT | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go<br>↓<br>go-build
--- | --- | --- |
| | | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |
| | | <div align="center">**WIP**<br>↓</div> |
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
↳ Q₄/JS | — | <strong>2.07x</strong> |
↳ G/JS | — | <strong>1.40x</strong> |

### Rewrite benchmark results

#### From PHP

Tool | Hand-written PHP<br>↓<br>Zend / FrankenPHP | Hand-written PureScript<br>↓<br>[gopurs](https://github.com/0x000000000000000000001/gopurs)<br>↓<br>Go (FrankenPHP ext)<br>↓<br>go-build
--- | --- | --- |
| | | <div align="center">**Stable**<br>↓<br>Coming soon...</div> |
| | | <div align="center">**WIP**<br>↓</div> |
[symfony/expression-language](https://github.com/symfony/expression-language) → [sfpurs](https://github.com/0x000000000000000000001/sfpurs-expression-language): 8 ops, JIT OFF | ~ 237.71 μs | ~ 68.97 μs |
↳ /PHP | — | 0.29x |
[symfony/expression-language](https://github.com/symfony/expression-language) → [sfpurs](https://github.com/0x000000000000000000001/sfpurs-expression-language): 8 ops, JIT ON | ~ 164.85 μs | ~ 63.57 μs |
↳ /PHP | — | 0.39x |
**Total** | <strong>~ 0.40 ms</strong> | <strong>~ 0.13 ms</strong> |
↳ /PHP | — | <strong>0.33x</strong> |
↳ μ/PHP | — | <strong>0.34x</strong> |
↳ Q₂/PHP | — | <strong>0.34x</strong> |
↳ Q₄/PHP | — | <strong>0.39x</strong> |
↳ G/PHP | — | <strong>0.33x</strong> |

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
↳ Q₄/JS | — | <strong>0.53x</strong> | <strong>0.47x</strong> | <strong>0.37x</strong>
↳ G/JS | — | <strong>0.37x</strong> | <strong>0.32x</strong> | <strong>0.29x</strong>
**Total** | <strong>~ 812.48 s</strong> | <strong>~ 332.71 s</strong> | <strong>~ 265.63 s</strong> | <strong>~ 256.92 s</strong>
↳ /JS | — | <strong>0.41x</strong> | <strong>0.33x</strong> | <strong>0.32x</strong>
↳ μ/JS | — | <strong>0.38x</strong> | <strong>0.32x</strong> | <strong>0.30x</strong>
↳ Q₂/JS | — | <strong>0.37x</strong> | <strong>0.31x</strong> | <strong>0.30x</strong>
↳ Q₄/JS | — | <strong>0.53x</strong> | <strong>0.47x</strong> | <strong>0.37x</strong>
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
↳ Q₄/JS | — | <strong>1.44x</strong>
↳ G/JS | — | <strong>0.34x</strong>
**Total** | <strong>~ 348.34 s</strong> | <strong>~ 140.73 s</strong>
↳ /JS | — | <strong>0.40x</strong>
↳ μ/JS | — | <strong>0.36x</strong>
↳ Q₂/JS | — | <strong>0.34x</strong>
↳ Q₄/JS | — | <strong>1.44x</strong>
↳ G/JS | — | <strong>0.35x</strong>
