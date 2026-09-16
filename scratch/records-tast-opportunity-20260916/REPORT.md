# Deep Record Updates: scalar replacement proof

2026-09-16. Scratch-only diagnostic. Original source and compiler/backend unchanged. The diagnostic agent initially supplied counts only; the parent subsequently ran the serialized timing campaign below. Original source and production output remain unchanged.

## Concrete current cost

The original `output/purust_output/Purs_Test_Records/src/lib.rs` uses `UnknownType = Value` throughout `updateRec`, despite `output/Test.Records/corefn.json` containing the exact closed structural record in typeTable entries 0–8. Each record stores its fields as `Option<Value>`, not native integer/record fields. Measured sizes are 24 bytes per Value and 48 bytes per two-field record payload.

`get_b` / `get_d` clone the child pointer before the old slot is cleared; each setter dispatches on Value and calls PerceusPtr::make_mut. The existing COW/reuse is already good: with the original seed retained, the full 10000-step call allocates and releases only 3 nodes, 168 bytes total. There is no per-iteration allocation problem to fix here.

## Controlled variant

`kernel-original.rs` is the exact extracted generated updateRec function. `kernel-scalar.rs` loads the four integer fields into local variables, runs the same descending loop and same additions/modulo, then writes the final four fields into the original representation once. It retains the same COW getters/setters and does not require uniqueness: retained input snapshots remain unchanged. No closed-form replacement of the benchmark loop was written.

Runtime instrumentation (`probe-count`, not the timing binary) counts the calls emitted by those two functions:

| 10000 updates | Generated | Scalar replacement |
|---|---:|---:|
| Borrowed field projections | 90000 | 9 |
| Cloning child getters | 20000 | 2 |
| Record setters / make_mut calls | 80000 | 8 |
| Heap allocations | 3 | 3 |
| Allocated bytes | 168 | 168 |

These are emitted operation counts, not a claim that each survives LLVM as an independent machine instruction. Measurements are in `operation-counts.txt` and `results.txt`.

## Validation

Thirty combinations: iteration counts 0, 1, 2, 3, 4, 5, 6, 31, 100, 10000; starting fields [0,0,0,0], [3,19,-100,7], [-5,999,12,-50]. Every output field matches generated Rust and both retained aliases of each original input remain unchanged. At 10000 updates from zero the complete result is [10000,20000,30000,20000]. Zero iterations perform no writes or allocations. Supplementary checks retain the b and d child records independently, and pass a root with an unrelated c=12345 field; both variants preserve these aliases and the extra field. Negative counts require an impractically long wrapping countdown in this release representation and are not tested; the workload and validated domain use nonnegative counts.

## TAST prerequisites

For this concrete worker, the required field shapes and primitive types already exist in the TAST. No new Haskell annotation is needed to know the record is closed, nor to resolve the four integer fields. The function is a pure tail-recursive loop with no calls observing intermediate records. A backend scalar-replacement / worker-wrapper transformation can preserve the Value ABI while carrying scalars around the loop and reconstructing at exit.

For a general pass spanning calls, extra summaries could be useful: which field paths a callee reads/writes, whether a parameter or intermediate record is retained/returned, and which aliases can observe mutations. Such summaries would justify keeping fields in scalar locals across calls. They are unnecessary on this self-contained example and must not be confused with the already-available structural typing.

This is a better immediate candidate than adding an unconditional uniqueness bit: the variant preserves COW and still removes almost all loop-level record operations.

## Reproduction and parent-controlled timing

Run `python3 prepare.py` to regenerate both extracted kernels, compile with the existing output's purust_core/perceus_ptr libraries, and run validation/counting. SHA and compiler arguments are stored in `manifest.json`.

Timing is exposed only for the parent to schedule without concurrency:

```
./probe-time original
./probe-time scalar
```

Each mode uses a retained shared seed, three warmups, then fifteen nanosecond samples. Each sample covers update, inspection of all four fields, and dropping the result. The 10000 count, seed input, and result values pass through black_box. The dedicated probe-time binary links the output’s native mimalloc allocator and contains no operation counters or allocation-counter wrapper. The retained shared seed matches Test_Records_initial: the real generated initializer caches the record in a module-value cell and returns a clone. Both use an opaque 10000 input. The kernel harness additionally inspects all four result fields; the official act instead extracts f and formats a String through its effect wrapper. Thus this remains a kernel diagnostic and its absolute times must not be compared directly to the README.

`probe-count` contains atomic operation counters and must never be used for speed comparisons.

Official README context: Deep Record Updates is 385.38 µs for generated Rust, 98.75 µs for native functional Rust, and 4.25 µs for the native imperative version (README.md:150). Those historical figures motivate this probe but are not current A/B measurements.

## Measured timing campaign and independent audit

The parent ran 21 rounds in alternating/shuffled variant order, 15 measured samples per variant per round after the three warmups: **630 measurements** total. Raw data and summaries are in `timing-results.json`.

| Same mimalloc kernel harness | Generated original | Scalar replacement |
|---|---:|---:|
| Median of each round's best sample | 388.750 µs | 3.000 µs |
| Reported sample median | 413.042 µs | 3.333 µs |

The median paired change in best samples is **−99.2464%**. All 21 rounds favor the scalar variant; individual paired changes range from −99.4261% to −99.1358%. Early baseline rounds were substantially slower than later rounds, so absolute distributions show warm-up/environment drift, but the effect exceeds that variability by two orders of magnitude. This is a demonstrated kernel opportunity, not an integrated backend result.

Independent read-only disassembly of `probe-time` confirms the work was not constant-folded away. `scalar_replacement` receives runtime n and the input record; LLVM emits a real modulo-5 accumulation loop with a backward branch, vectorized/unrolled eight iterations at a time, plus a scalar remainder loop. The a/c/e constant-increment recurrences become additions/multiplications of runtime n under the release integer semantics. Those simplifications are performed by LLVM after scalar replacement; the handwritten variant contains the same iterative additions and checked Euclidean remainder, not a benchmark-specific mathematical formula. All four output fields are observed and black-boxed before destruction, so dead-field elimination does not explain the result.

Audit commands use `xcrun llvm-objdump --syms probe-time` and `--disassemble-symbols=<scalar_replacement symbol>`. In the current binary the vector loop branches from 0x100004d40 to 0x100004c0c; the scalar remainder loop branches from 0x100004d8c to 0x100004d68. The ordinary field getter/unwrap and setter calls occur before/after the loops, not inside them.

### Integer and semantic limits

Both compared kernels use i64 exactly as the existing output does. `purust_core::Value::Int` stores i64, `unwrap_int` returns i64, and `mk_int(i64)` stores it without narrowing. Both use `checked_rem_euclid(5).unwrap_or(0)`. The benchmark inputs and checked examples remain in range, so no overflow occurs. Thus the experiment has not introduced an i32-versus-i64 semantic change. It is also **not a proof that the existing backend's treatment of all PureScript Int boundary cases is correct**; that broader issue is outside this probe.

A generic implementation must preserve the language/backend's exact integer width and wrapping/checking semantics, evaluate simultaneous record updates against the previous logical iteration state, and maintain exit materialization on every path. It must not apply the transformation across a call that can observe/retain intermediate records without a supporting analysis. Here there are no such calls, effects, or escaping intermediate versions, and the COW boundary still preserves existing shared inputs and child aliases.

### Generic compiler transformation supported by the present TAST

Implement a scalar worker for a pure tail-recursive record loop, keeping the existing Value entry/exit wrapper:

1. Discover the fixed primitive field paths used/updated by the loop using the closed row and primitive types already in the TAST.
2. Load their initial values at the worker entry and represent them as scalar loop-carried locals.
3. Translate each recursive record update into new scalar values, preserving old-value dependencies and recursion/branch behavior.
4. Materialize the record once at each escaping/return boundary, preserving untouched fields and COW behavior.

This is scalar replacement / worker-wrapper transformation, rather than an annotation naming this benchmark. No new Haskell data is required for this example. Field-sensitive read/write/retention summaries would help extend it across other functions. The experiment shows that the current structural types are already sufficient to unlock a large opportunity when the backend carries them through to the loop representation.

A full-runner scratch copy was compiled successfully in 13.83 seconds at `full-scalar`, with only the generated updateRec worker replaced. All other copied generated files are byte-identical to the parent’s initial guard `after` workspace. The executable is `full-scalar/benchmark` (also `full-scalar/target/release/purust_output`); build flags are opt-level=3 and debug=false, and the original mimalloc runner is retained. See `full-scalar.patch`, `full-scalar-manifest.json`, and `full-scalar-build.log`.

## Full official runner: measured impact

The parent subsequently ran 21 shuffled paired rounds of those two full binaries, with no concurrent builds or other benchmark processes. Each invocation retains the official global warmup and best-of-10 per benchmark. The official validator checked all 14 expected results and the total in all 42 invocations. Raw logs, executable SHA256 values, the command and all timings are saved in `full-comparison/results.json` and its adjacent logs.

| Quantity | Original | Scalar variant | Median paired time change |
|---|---:|---:|---:|
| Records, median reported time | 376.17 µs | 3.54 µs | −99.024% |
| Total, median sum of displayed rows | 9.01534 ms | 8.76976 ms | −4.041% |

The paired percentage is the median of each round's scalar/original ratio. It is deliberately not the ratio of the two separately calculated medians: the total includes the much larger and noisier RBTree workload. The Records time reduction is about 373 µs, consistent with a roughly four-percent contribution to this suite. Other per-case fluctuations in this campaign are not attributed to the record transformation.

Reproduce from the repository root:

```sh
python3 scratch/compact-guards-integration-20260916/compare.py --variant original scratch/compact-guards-integration-20260916/after/benchmark --variant scalar scratch/records-tast-opportunity-20260916/full-scalar/benchmark --rounds 21 --output scratch/records-tast-opportunity-20260916/full-comparison-repeat
```

This confirms a substantial performance opportunity with the current TAST. The scalar-worker transformation is still a scratch Rust variant, not a production Purust pass. The historical README numbers (385.38 µs compiled / 4.25 µs native imperative) provide context, while the paired measurements establish the present effect.
