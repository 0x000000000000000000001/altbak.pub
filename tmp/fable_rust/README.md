# Fable Rust numeric benchmarks

Two separate columns measure two source routes:

- PureScript -> sharpurs -> generated F# with compatibility adaptations -> locally patched Fable -> Rust.
- Hand-written, typed F# -> unmodified Fable -> Rust, following the native Koka/Haskell/OCaml reference implementations.

## Hand-written F# reference

```sh
python3 tmp/run_fable_native_benchmark.py --dotnet /path/to/dotnet10/dotnet --update-readme
```

This builds [NativeBench.fs](NativeBench.fs) directly with the unmodified Fable 5.17.2 NuGet compiler and compiles its generated Rust with optimization level 3 and mimalloc. It needs .NET 10, Rust/Cargo, the cached Fable package and Cargo dependencies. It does not need sharpurs output or a Fable source checkout. Logs, source/compiler fingerprints, generated code and three validated measurement processes are saved in a fresh `var/benchmark/fable-native-rust/` directory.

The source mirrors the native Haskell/OCaml implementations: recursive AST evaluation and Fibonacci, nested immutable records, a linked-list prime sieve and a persistent red-black tree. As in those references, List/Array, Church, Polymorphism, State and Lazy use simplified numeric loops, and RowToList returns the known field count. These rows do not measure the original PureScript abstractions. State takes depth 60 with 20 repetitions; RowToList takes 0, matching the native inputs. The Rust driver supplies timing, opaque inputs and result validation; all numeric kernels are compiled from F#.

The tree's four balancing cases use separate matches with named children to avoid incorrect variable shadowing in Fable 5.17.2's translation of deeply nested patterns. This preserves the original rotations and tree structure; the compiler and its generated Rust remain unmodified.

The timing protocol below is shared with the sharpurs/Fable column. The historical Koka/Haskell/OCaml tables use one process each; both Fable columns report medians of three processes. The C reference also differs algorithmically in its prime test and arena allocation, so these timings compare the published implementations, not identical allocations across languages.

## PureScript through sharpurs

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
