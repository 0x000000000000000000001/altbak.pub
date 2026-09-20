# Fable Rust numeric benchmark

Run from the repository root:

```sh
python3 tmp/run_fable_rust_benchmark.py --dotnet /path/to/dotnet10/dotnet --update-readme
```

Prerequisites: .NET 10, Rust/Cargo, the Fable 5.17.2 NuGet package in its default cache location, the sibling `Fable` source checkout, and cached Cargo dependencies. The launcher accepts explicit `--fable-package`, `--fable-source`, and `--source` paths. It consumes `run/bak/sharp/output/Main`, the F# output produced by sharpurs; generate that output with the project's sharpurs workflow when changing PureScript sources. It rebuilds Fable and Rust, records the supplied F# source hashes, and does not rebuild PureScript itself.

Each invocation creates a fresh directory under `var/benchmark/fable-rust/`, with compiler/build logs, source and executable hashes, a correctness check, three process logs, and aggregated results. Only verified measurements update the 14 Fable cells and their total. Existing results are preserved.

## Compatibility work

This is **Fable 5.17.2 with a local Rust emitter patch**, not the unchanged NuGet compiler. The patch is applied to a temporary copy of `Fable2Rust.fs`; the source checkout and installed package remain untouched. It fixes ownership of boxed values, explicit generic/lambda types, null values, recursive getters and closures, and casts between boxed objects and reference ADTs. The generated algorithms retain their boxing and reference counting.

The extractor retains the declarations reachable from the 14 numeric kernels and the existing F# FFI members they need. The adapter removes seven unused local closure aliases. The pure runtime supports the generated curried function shapes without .NET reflection, uses the built-in F# `unbox`, and expresses literal type tests in a form Fable handles. Two .NET reflection exception wrappers in generated RBTree call adapters propagate their original exception instead. These adjustments preserve successful benchmark calculations; they are not a general replacement for sharpurs' Effect/Aff runtime.

## Measurement

The Rust harness uses opaque inputs, consumes results, validates all expected outputs, performs three global and three local warm-ups, and measures ten calibrated batches of at least 10 ms (capped at 2^24 calls). Each table cell is the median across three processes of the minimum per-call batch time. The displayed total sums the median cells.

Rust uses release optimization level 3, mimalloc (as in the Purust numeric harness), and a benchmark thread with a 1 GiB stack reservation (as in sharpurs' native entry point). Thread startup and build time are outside the timed kernels. The existing Rust FFI columns use their older measurement protocol.
