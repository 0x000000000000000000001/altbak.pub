# Benchmark methodology

## What the columns compare

The compiled columns start with the same 14 PureScript kernels in `src/Test`.
The Haskell, Koka, OCaml and native F# references translate their functional
structures according to the [source contract](../tmp/fp_reference_contract.md).
The FP-style FFI columns also retain these structures, callbacks and explicit
dictionaries. In FFI columns, the backend translates the PureScript wrapper;
the native kernel is supplied as hand-written code. Their host-language control loops may implement tail recursion;
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

## Additional Go diagnostics

The Go and JS tables include additional diagnostics, with a [separate initial measurement
archive](benchmark-results/2026-09-21-go-diagnostics.md). The existing 14-case
results and totals remain the published core baseline. These diagnostics do not
assert that Go must outperform JavaScript.

Sources use `src/Test/ArrayIndexing.purs`, `src/Test/JsonTypedAst.*` and
`src/Test/JsonDecoding.*`.
Run them through `bin/go/run --test ArrayIndexing`, `--test JsonTypedAst` or `--test JsonDecoding`,
and the equivalent `bin/js/run` commands. They are opt-in cases with their
recorded protocols, not additions to the ordinary fourteen-case or `--x` totals.
PBO dependencies are resolved only in the `JsonTypedAst` isolated workspace.
`JsonDecoding` uses ordinary Argonaut codecs without a PBO dependency.
Historical measurement manifests retain their original source paths; rebuild
the diagnostics after their relocation from `srx/Test` to `src/Test`.

- **Array Indexing (multiple sizes):** repeatedly read elements from arrays of
  16, 1,024 and 16,384 elements, exercising both native and boxed representations
  and the generated conversion between them. Obtain inputs through an optimization
  barrier, consume a checked checksum, and keep input construction outside the
  indexed-read interval. Report time and allocated bytes per access for each size,
  rather than hiding the size dependence in one aggregate. Check the generated
  code actually exercises the intended conversion; allocations must not grow
  with the entire array length for a single read. Compare the same PureScript
  workload compiled to Go and JS.
- **JSON → Typed AST (parse + decode):** use a fixed, versioned corpus of real
  TAST/tcorefn JSON and the same PBO decoding logic compiled to Go and JS. Measure
  JSON parsing, decoding of the parsed objects into the typed AST, and their
  combined duration separately. File reading and compiler optimization/emission
  are outside these intervals. Preserve and verify the types and decoded
  structure, and recreate decoded state on every invocation. Report corpus size,
  time and Go allocations. Use one worker for the sequential comparison; record
  any parallel variant separately. Parsing includes the selected JSON library
  and runtime costs, so a parsing-only result cannot explain all AST decoding.
- **JSON Decoding (records and arrays):** compile the same ordinary Argonaut
  record codecs and a tagged event codec to Go and JS. Exercise nested records,
  arrays, optional fields, Unicode and numeric values. Measure parsing, decoding
  of already parsed JSON and their combination separately. Missing fields, wrong
  types and nested error paths are checked outside timing. A fixed independent
  oracle verifies every decoded field and complete error message; fast failure
  cannot improve the timed successful workload. Sources and fixture generation
  are documented in [the corpus](../test/fixtures/json-decoding/README.md).

Freeze sources, inputs, binaries and runtime settings, validate outputs, and
collect three independent processes per configuration after warm-up.
The measured array protocol uses three warm-ups and ten batches, GOGC=800; JSON
uses two warm-ups and five corpus samples per phase, GOGC=100. Both report the
median of three process minima with Go GOMAXPROCS=1. This caps Go runtime execution
as well as application work; Node retains its default background threads, so the
comparison does not establish equal total CPU use. Array drivers call generated
native Go kernels directly; JSON retains generic callback adapters. These
boundaries are recorded rather than equated to every historical table cell.

Keep per-size and per-phase results alongside the table summary. Compiler-workload
measurements use a different boundary from the numeric core suite and must remain
outside its historical total.

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

These are warmed-up minimum batch durations, not mean latency or tail-latency
measurements. Some short JIT kernels vary noticeably between processes: in the
September 20 campaign, Java AST ranged from 0.140650 to 0.279165 microseconds.
The archive retains all three measurements. Printed decimal places preserve
small nonzero values; they do not establish that precision or justify rankings
based on differences smaller than the observed variation.

The numeric references, including the published Purust and sharpurs/Fable Rust
columns, call numeric functions directly. Other compiled columns and the FFI
columns retain their generated Effect/call adapters. The additional `rust-pure`
diagnostic in the measurement archive retains Effect and is not the published
Purust numeric column. Adapter overhead is included where present,
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

The [September 20, 2026 archive](benchmark-results/2026-09-20.json) records the
39 published columns and two additional diagnostics (`rust-pure` and
`wasm-pure`): 123 process runs and 1,614 validated case outputs. It includes all
per-process measurements, the common source fingerprints, build identities and
recorded tool configurations. Missing version information is marked explicitly.
Complete raw logs and manifests remain in
`var/benchmark/fairness-20260920/{core-measurements,extended-measurements}`;
their hashes are in the archive. These corrected tables establish a new baseline.

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

`./bin/run` and `./bin/run --x` run the configured backends. They do not alone
produce all the published native references or aggregate three processes.
After collecting the core and extended plans from the same frozen sources,
publish their validated medians with:

```sh
python3 bin/benchmark/publish.py --campaign path/to/core-campaign --campaign path/to/extended-campaign
```

These deliberately small synthetic programs help examine particular compiler
transformations. Neither their aggregate total nor a faster cell establishes
that a compiler beats every hand-written implementation, or removes the need to
profile real applications and implement appropriate FFI.
