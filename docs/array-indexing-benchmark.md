# Array indexing diagnostic

This diagnostic measures indexed reads from arrays of **16, 1,024 and 16,384
integers**. It compares native and boxed representations in Go, JS, C and Rust.
It is separate from Array Processing and excluded from the core total.

## Run it

Build first, then measure after all other builds and benchmarks have stopped:

```sh
python3 bin/benchmark/array-indexing.py --build-only --build-dir run/bak/go/modes/test-ArrayIndexing
python3 bin/benchmark/array-indexing.py --run-only --build-dir run/bak/go/modes/test-ArrayIndexing --output var/benchmark/array-indexing/campaign
```

A build prepares every runtime. Add `--runtime go`, `js`, `c` or `rust` to
measure only that runtime; otherwise all four are measured. The ordinary
`bin/go/run --test ArrayIndexing` and `bin/js/run --test ArrayIndexing` commands
select their respective runtime and preserve this build/run separation.

Backends must already be built. Go uses its existing JS compiler bundle
(`GOPURS_JS=1`); the local typed `purs` fork takes priority on `PATH`.
Rust resolves FFI from its own library ports. The
[driver directory](../bin/benchmark/array-indexing/) contains the harnesses.

## What is measured?

Each compiled kernel repeatedly reads a runtime-created array, wraps its index
and accumulates a checked checksum. Input construction is outside timing.
Runtime inputs and dynamic start/count arguments prevent a precomputed answer.
The Go harness calls the native entry point directly, excluding the public
boxed adapter's one-time input conversion.

| Representation | Access path |
| --- | --- |
| Go native | `[]int64` |
| Go boxed | `Value` containing `[]Value`, cast to `Array Int` at `unsafeIndex` |
| JavaScript | Ordinary arrays for both exports |
| Rust | Typed `IntArray`/borrowed slice where admitted; boxed values use the borrowing array accessor |

Rust's earlier generic representation used the same `Rc<Vec<Value>>` path for
both rows. Read each campaign's generated-code evidence when comparing versions.
A bounded allocation preflight and code audit reject whole-array conversions
inside the read loop. The generated sources are retained with the results.

## Protocol and checks

- Three independent processes per runtime, with varying seeds.
- Three warm-ups, then ten batches per case; normally 8,388,608 reads per batch.
- Published time: median of the three process minima, divided by access count.
- Every checksum checked by both driver and collector.
- Go: `GOGC=800`, `GOMAXPROCS=1`, PGO off. Rust: optimization level 3, thin LTO.

Go allocation bytes come from `runtime.MemStats.TotalAlloc`; Rust uses a counting
allocator around mimalloc. Counters are read outside timing and include amortized
closure setup. The archive keeps every batch and reports median and maximum
bytes per access. JS allocation bytes are unavailable.

Sources, oracles, compiler versions and artifact hashes are recorded. Validation
rejects changed inputs, stale executables and incorrect results. These timing
boundaries differ from historical core benchmarks; do not present their ratios
as speedups over the Array Processing row.
