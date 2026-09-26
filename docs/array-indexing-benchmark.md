# Array indexing diagnostic

Build first, then run after every other compiler has stopped:

```sh
python3 bin/benchmark/array-indexing.py --build-only --build-dir run/bak/go/modes/test-ArrayIndexing
python3 bin/benchmark/array-indexing.py --run-only --build-dir run/bak/go/modes/test-ArrayIndexing --output var/benchmark/array-indexing/campaign
```

The same `src/Test/ArrayIndexing.purs` produces the official PureScript JavaScript,
Gopurs Go, C reference and purust Rust kernels. The Go backend uses the existing
JavaScript compiler bundle (`GOPURS_JS=1`), and the local TAST `purs` fork has
priority on `PATH`. Installed dependencies come from the existing Go package set;
the Rust kernel compiles the same sources against the purust library ports, so
purust resolves the matching `.rs` FFI. Nothing rebuilds a backend. The drivers
live in `bin/benchmark/array-indexing/`.
Pass `--runtime go`, `--runtime js`, `--runtime c` or `--runtime rust` to measure
three processes of one runtime; without this option, the collector measures all
four. A build prepares every artifact.
The normal `bin/go/run --test ArrayIndexing` and `bin/js/run --test ArrayIndexing`
entry points select their respective runtime and preserve the build/run separation.

The Rust runtime has one array representation: `Rc<Vec<Value>>`. Both rows
exercise it, because the opaque `BoxedArray` maps to the same runtime value, and
the generated loop reads through the borrowing `Value::array_get` accessor
without cloning the buffer. A generated-code audit rejects whole-array
conversions, and a counting global allocator reports requested bytes per access
(amortized closure setup remains visible). A bounded preflight rejects a
whole-array copy before a long batch could allocate terabytes on the largest
input. Native and boxed rows are therefore expected to be equal for Rust; they
are kept separate to compare with the runtimes that do have distinct
representations.

`Array Int` values may also keep their elements unboxed (`Value::IntArray`).
When a self-recursive loop captures such an array and reads it by index, the
generator emits the loop twice: the generic one, and one over a borrowed
`&[i64]` slice that is selected once per loop call. `Array Int` reads are typed
as `Int`, so the integer is only boxed where a consumer needs a `Value`. The
native row materializes its input as `IntArray`, so it measures the slice path,
while the boxed row passes a boxed array and keeps the accessor path.

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
come from `runtime.MemStats.TotalAlloc`; Rust bytes come from a counting global
allocator wrapping mimalloc. Both counters are read outside the timed interval;
the archive retains all batches and reports their median and maximum. The
allocation count includes any generated closure setup amortized over the batch.
JS allocation bytes are unavailable. Run settings are `GOGC=800`,
`GOMAXPROCS=1`, PGO off, and the Rust profile is `opt-level=3` with thin LTO;
tool versions and input/artifact hashes are archived.

This is an additional diagnostic, not one of the historical fourteen cases.
It does not share a timing boundary with the existing Array Processing row,
and its numbers must not be summed into the core total or presented as a
historical speedup. Source validation rejects modified inputs or stale binaries.
