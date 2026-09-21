# Array indexing diagnostic

Build first, then run after every other compiler has stopped:

```sh
python3 bin/benchmark/array-indexing.py --build-only --build-dir run/bak/go/modes/test-ArrayIndexing
python3 bin/benchmark/array-indexing.py --run-only --build-dir run/bak/go/modes/test-ArrayIndexing --output var/benchmark/array-indexing/campaign
```

The same `src/Test/ArrayIndexing.purs` produces the official PureScript JavaScript and
Gopurs Go kernels. The Go backend uses the existing JavaScript compiler bundle
(`GOPURS_JS=1`), and the local TAST `purs` fork has priority on `PATH`.
Installed dependencies come from the existing Go package set. Nothing rebuilds
the backend. The drivers live in `bin/benchmark/array-indexing/`.
Pass `--runtime go` or `--runtime js` to measure three processes of one runtime;
without this option, the collector measures both. A build prepares both artifacts.
The normal `bin/go/run --test ArrayIndexing` and `bin/js/run --test ArrayIndexing`
entry points select their respective runtime and preserve the build/run separation.

Each kernel reads a runtime-created array of 16, 1,024 or 16,384 integers,
wrapping its index and accumulating a checked checksum. Input creation is
outside timing. Native Go inputs use `[]int64`; opaque boxed inputs use the
runtime's `Value` containing `[]Value`. The boxed PureScript kernel casts the
input to `Array Int` immediately inside `unsafeIndex`, reproducing the typed
boxed-to-native boundary that previously copied the entire array per read.
JavaScript uses ordinary arrays for both PureScript exports.

The drivers contain input generation, clocks, allocation counters and an oracle;
the indexed-read loops are compiled from PureScript. The Go harness calls the
generated native entry points directly. It excludes the public boxed function
adapter that would convert the complete native input once per invocation.
The audited generated code is preserved in the campaign. A bounded allocation
preflight and a generated-code check reject whole-array conversions in the loop.
The compiler can optimize normally; runtime inputs and dynamic start/count
parameters prevent precomputing the checksum.

There are three independent processes per runtime, each with three warm-ups
and ten measured batches per case. Each batch performs 8,388,608 reads by
default. Seeds vary across processes. Every checksum is checked in the driver
and independently by the collector. Time is the median of the three process
minimum batch times, divided by access count. Go allocation bytes per access
come from `runtime.MemStats.TotalAlloc`, with the counters outside the timed
interval; the archive retains all batches and reports their median and maximum.
The allocation count includes any generated closure setup amortized over the
batch. JS allocation bytes are unavailable. Run settings are `GOGC=800`,
`GOMAXPROCS=1`, PGO off; tool versions and input/artifact hashes are archived.

This is an additional diagnostic, not one of the historical fourteen cases.
It does not share a timing boundary with the existing Array Processing row,
and its numbers must not be summed into the core total or presented as a
historical speedup. Source validation rejects modified inputs or stale binaries.
