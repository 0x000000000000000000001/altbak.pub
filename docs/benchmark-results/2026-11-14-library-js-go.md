# JS and Go library benchmarks — November 14, 2026

The README now contains all **207 cells** for the 23 new tests across four JS
columns and five Go columns. Each value is the median of three independent
process minima: **27 processes, 621 validated WIP outputs**. The 14 core cases
were also validated in each process (378 outputs); their published cells and
totals were preserved. The new cases remain excluded WIP and have no total.

The [JSON archive](2026-11-14-library-js-go.json) contains every per-process
result, the frozen plan, source fingerprints, build manifests and raw-log hashes.
The full local campaign is
`var/benchmark/library-js-go-20261114T162925Z/measurements-final`.
Sources were frozen across the campaign; all builds finished before sequential
measurement began. Plan ordering rotates between repetitions.

## Environment and columns

- Host: Apple M4 Pro; macOS-26.1-arm64-arm-64bit.
- Official PureScript: 0.15.16; Node: v24.8.0.
- Go: go version go1.27.0 darwin/arm64; GOGC=800, PGO disabled.
- JS: official compiler, Arista/PBO, FP-style FFI and imperative FFI.
- Go: gopurs release, gopurs development, psgo, FP-style FFI and imperative FFI.
- The development profile passes `-N -l` to the generated Go packages.
- Campaign: 2026-11-14T16:31:29.924907+00:00 to 2026-11-14T16:35:59.546325+00:00.

## Workload and qualification

Immutable fixtures are prepared before timing. Each measured invocation performs
the operation and consumes its result against an independent integer oracle.
Mutable references are allocated anew on every invocation. The harness uses
three global warm-ups, three local warm-ups, calibrated batches and ten timed
batches per process. [Suite documentation](../library-benchmarks.md) specifies
the operations, consumers, native factories and independent oracles.

FP-style references retain functional intermediate results, callbacks, immutable
record updates and balanced traversal. Imperative references may use direct loops
and mutation, with the same result consumers and fresh mutable state. PureScript
cases pass their inputs through `Bench.opaque`; the native factory closes over
immutable fixtures and passes the scalar bound through that barrier. These
columns therefore include differences in representation and FFI boundaries.

Four reference loops use `tailRecM` so all compiled routes execute them without
stack overflow. JS and Arista use the stock compiler; dependency fetching does
not leave TAST artifacts from the local compiler in their output.

The psgo workspace includes a local correction to `Data.String.CodeUnits.length`:
its upstream FFI counted code points rather than UTF-16 units. The Unicode case
requires 4000 units for 3000 code points. The correction is part of the recorded
build sources, and no upstream checkout was changed.

All nine measured columns passed their result validators. Publication checks
require three processes, correct per-process values, unchanged source/build
hashes and exact cell medians. The publication audit confirmed that only the
46 new table rows changed numerically and that every source link was preserved.
The archive's base commit identifies the checkout before the new FFI and runner
changes; the recorded source fingerprints identify the exact measured working tree.

## Reproduction

```sh
python3 bin/benchmark/measure.py \
  --plan var/benchmark/library-js-go-20261114T162925Z/plan-final.json \
  --output path/to/new-campaign
python3 bin/benchmark/publish.py --wip-only --campaign path/to/new-campaign
```

The frozen plan references local artifacts and verifies their hashes. Rebuild and
freeze a new plan after changing sources or compilers. The historical core
baseline remains in the README; the core timings collected here are archived for
qualification and do not replace that baseline.
