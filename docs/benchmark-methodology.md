# Benchmark methodology

## What the columns compare

The compiled columns start with the same 14 PureScript kernels in `src/Test`.
The Haskell, Koka, OCaml and native F# references translate their functional
structures according to the [source contract](../tmp/fp_reference_contract.md).
The FP-style FFI columns also retain these structures, callbacks and explicit
dictionaries. Their host-language control loops may implement tail recursion;
they do not demonstrate that the host compiler optimizes recursive source.
Dynamic languages use explicit dictionaries without static row proofs.

The imperative columns, including C, intentionally permit other algorithms,
mutable data, preallocation and numerical simplifications. Their ratios include
those source-level choices. In particular, C allocates its AST/tree arenas before
timing and reuses them. These are useful alternative implementations, not tests
of compiling the same FP program.

Normal compiler optimizations are permitted: specialization, fusion, constant
folding, conversion to loops, allocation reuse and invariant-code motion.
A compiler reducing a generic computation to a constant is valid. Replacing that
computation manually with its expected result in an FP reference is not.

## Core timing boundary

Core entry points return integers. Conversion to text, logging, compilation and
process startup are outside the interval being measured. `Bench.measureBatch`
calls the supplied effect for every iteration; constructing the effect must not
execute or cache the measured computation. Every result is consumed through a
native sink or optimization barrier. Each batch checks its final result against
the warm-up value; the external validator independently checks all 14 oracles.

There are three global warm-up suites and three local warm-ups. Calibration
doubles the invocation count until a batch takes at least 10 ms, capped at
2^24 invocations. Each process records the minimum per-invocation time of ten
batches. The published cell is the median of three independent processes, run
sequentially after all compilation has finished. The total sums these cell
medians before rounding; it is not a separately selected fastest total.

Direct numeric references call native functions; compiled programs and the FFI
columns retain their generated Effect/call adapters. That overhead is included,
as are input/result barriers. Very short kernels can mostly measure this
overhead. Native representations, integer widths, garbage collectors and
allocators differ. Rust numeric references use mimalloc; Fable's driver reserves
a 1 GiB stack. These choices belong to the measured configuration, not to a
universal language ranking. All current core values fit the common integer range.

The native FFI State entry historically accepts depth 60 with 20 repetitions;
the PureScript entry accepts 20 repetitions with depth 60. Work is equivalent
at the benchmark point, but the optimizer sees different opaque parameters.
This interface difference remains recorded, rather than hidden by the labels.

The older protocol timed individual `Effect String` calls, discarded their
results and included formatting. Those numbers must not be mixed with this
numeric batch protocol or presented as compiler speedups after this correction.
Recorded zeros from the old microsecond clocks are not zero-cost executions.

## Fable routes

The hand-written F# route uses unmodified Fable. The sharpurs route consumes a
generated F# snapshot, extracts reachable numeric kernels, adapts its runtime
and uses a locally patched Fable Rust emitter. It is a compatibility experiment,
not an unchanged end-to-end build with the public Fable package. Both retain
their numeric algorithms. Exact adaptations and commands are described in
[the Fable benchmark documentation](../tmp/fable_rust/README.md).

## Extended tests

Extended tests retain an end-to-end Effect/Aff boundary and ten individual
samples, with three independent processes. They are a separate experiment from
the numeric core. Inputs for STArray, StringOps and the async computations are
obtained inside each effect, preventing module initialization from doing the
work before the timer. Every timed result is stored in an observable sink and
compared with its warm-up value before the timer stops; the external validator
checks the warm-up value against an independent oracle. These short string
checks and the Effect/Aff adapters are part of the end-to-end interval.

- File I/O performs 10,000 writes and reads and checks each read's content. It
  measures synchronous filesystem APIs with OS caching, without an `fsync`
  durability guarantee.
- STArray creates and modifies a small mutable array on every invocation.
- StringOps performs 1,000 regex/split iterations on every invocation.
- Aff Operations requests a 10 ms delay. This is a timer/runtime check, not a
  measurement of general async throughput.
- Parallelism launches ten `fib 42` tasks. Its checksum is reduced safely within
  the common 32-bit range and is checked. JavaScript's single-thread scheduler
  and Go/Rust's multicore runtimes use different CPU resources; this row compares
  runtime parallelism, not sequential code-generation efficiency.

## Evidence and reproduction

Every published campaign keeps its build plan, source fingerprints, compiler
commands/profiles, executable identities, raw stdout/stderr, validated outputs
and per-process measurements. The measurement command is:

```sh
python3 bin/benchmark/measure.py --plan path/to/plan.json --output path/to/new-campaign
```

The plan names frozen artifacts and run-only commands. Numeric executables carry
their build SHA. The collector checks source stability, outputs, batch counts
and positive durations, then aggregates three processes. It never builds while
measuring. `--resume` accepts only the same source snapshot and plan and verifies
the saved stdout before reusing a completed process.

Backend runners expose `--build-only` and `--run-only`. The numeric launchers in
`tmp/run_benchmarks.py`, `tmp/run_purust_benchmark.py` and both Fable launchers also
accept `--build-only`. Source fidelity has separate native tests in `test/native`;
`test/benchmark-batches.py` checks deferred execution, exact invocation counts,
recomputation and rejection of incorrect results.

These deliberately small synthetic programs help examine particular compiler
transformations. Neither their aggregate total nor a faster cell establishes
that a compiler beats every hand-written implementation, or removes the need to
profile real applications and implement appropriate FFI.
