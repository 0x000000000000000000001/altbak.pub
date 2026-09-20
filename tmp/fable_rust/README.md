# Fable Rust numeric benchmarks

Two separate columns measure two source routes:

- PureScript -> sharpurs -> generated F# with compatibility adaptations -> locally patched Fable -> Rust.
- Hand-written, typed F# -> unmodified Fable -> Rust, translating the PureScript functional kernels directly.

## Hand-written F# reference

```sh
python3 tmp/run_fable_native_benchmark.py --dotnet /path/to/dotnet10/dotnet --update-readme
```

This builds [NativeBench.fs](NativeBench.fs) directly with the unmodified Fable 5.17.2 NuGet compiler and compiles its generated Rust with optimization level 3 and mimalloc. It needs .NET 10, Rust/Cargo, the cached Fable package and Cargo dependencies. It does not need sharpurs output or a Fable source checkout. Logs, source/compiler fingerprints, generated code and three validated measurement processes are saved in a fresh `var/benchmark/fable-native-rust/` directory.

The source follows the [functional reference contract](../fp_reference_contract.md): recursive AST evaluation and Fibonacci, immutable lists and nested records, a linked-list sieve, a persistent tree, Church functions, generic Monoidish dictionaries, State closures, non-memoizing thunks, allocated arrays and typed recursive row dictionaries. State takes 20 repetitions with depth 60 and fresh initial state zero; RowToList takes the opaque input 10000, matching PureScript. The Rust driver supplies timing, opaque inputs and result validation; all numeric kernels are compiled from F#.

The tree's four balancing cases use separate matches with named children to avoid incorrect variable shadowing in Fable 5.17.2's translation of deeply nested patterns. This preserves the original rotations and tree structure; the compiler and its generated Rust remain unmodified.

Two additional source helpers retain the original operations while satisfying the stock Rust emitter: `applyChurchStep` avoids borrowing and moving the same Church function in one expression, and the generic `recordKeys` factory keeps the row dictionary's phantom type in scope. Neither substitutes a result or disables compiler optimization.

The timing protocol below is shared with the sharpurs/Fable column and the corrected Haskell/Koka/OCaml references: three-process medians. Older F# results using numerical shortcuts (57.46 ms total) are superseded. C remains an imperative reference with different algorithms and allocation strategies.

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
