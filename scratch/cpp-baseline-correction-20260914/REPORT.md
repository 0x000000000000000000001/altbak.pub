# C++ benchmark correction — 2026-09-14

The three C++ columns have been rebuilt and remeasured. The corrected totals are
**951.00 ms compiled pscpp**, **133.23 ms native FP**, and **21.33 ms native handwritten**.
These replace invalid measurements; dividing the old and new totals does not measure
an optimization gain because units, native workloads and compiler flags changed.

## What was wrong

The old C++ clock returned milliseconds since an absolute epoch. `Bench.purs`
interpreted its differences as microseconds, then divided by 1000 for the total.
Every displayed C++ duration was therefore scaled incorrectly by 1000. Formatting
those millisecond differences to two decimal places also hid work below 10 μs.

The native FP implementation returned the wrong nominal result for 9 of 14 tests;
the handwritten implementation did so for 10. These were the observed discrepancies:

| Test | Old FP | Old handwritten | Required result |
| --- | ---: | ---: | ---: |
| List | 809100 | 809100 | 202950 |
| TCO | 705082704 | 705082704 | 100000 |
| Records | 60000 | 60000 | 20000 |
| Primes | 95 | 95 | 21536 |
| Red-black tree | 22 | 100000 | 22 |
| Polymorphism | -2014260032 | -2014260032 | 10000000 |
| State | 60 | 60 | 1200 |
| Lazy | 1000 | 1000 | 1000000 |
| Array | 404550 | 404550 | 202950 |
| RowToList | 0 | 0 | 5 |

The old `make main` invocation did not request `-O3`; those flags existed only in
the generated `release` recipe. Existing objects could be reused. The old runner
also generated placeholder FFI implementations after a link failure.

The previous explanation that LLVM had deleted a C++ `std::set` and reduced it
to `return n` was not supported: insertion and destruction were present in the
audited object code. The old numbers do not establish a C++ versus Rust allocation
elision advantage. [Previous README values](previous-readme-cpp.md) are retained
solely as a record of the invalid dataset.

## Corrections

- `bin/cpp/bench_ffi.cc` uses relative `steady_clock` microseconds and an opaque
  input barrier. Kernels contain no artificial `volatile` accumulator or output.
- Both native columns now implement the benchmark results and nominal workloads.
  The FP column retains persistent structures and closures. Closure captures share
  immutable environments rather than recursively copying `std::function` chains.
- The handwritten tree uses owned mutable Okasaki nodes, performs all 100,000
  insertions, and returns depth 22. Its AST has automatic ownership and destruction.
  Scalar loops remain available for normal compiler optimization.
- Ackermann now honors its input. The existing FFI ABI is preserved: State takes
  depth 60 and repeats 20 times; Lazy takes depth 1000 and repeats 1000 times.
- Each mode builds in its own workspace, with a local generated `Main`, copied
  PureScript sources, fresh native objects, and explicit C++11 `-O3 -DNDEBUG`.
  The project-wide backend links and `src/Main.purs` are not rewritten.
- Only the imported module closure is generated. No FFI stubs or incorrect
  `boxed` primitive equality operators are injected. Provider headers are preserved.
- All modes perform three global warm-up passes of their own suite. The normal
  runner verifies the 14 displayed values, labels, count and total after execution;
  a separate smoke path executes each action once without timing. A wrong result
  produces a nonzero exit status. Manifests record inputs, flags and binary hashes.

## Validation

- **176 independent result checks** across the 28 native kernels, including
  nominal and smaller inputs, passed at `-O3` and under ASan/UBSan.
- The clock unit and preservation of opaque integer, floating-point and string
  inputs passed direct checks. The unit check compares a real wait against an
  independently expressed microsecond duration.
- The handwritten tree passed all four rotation cases, ordering, node counts,
  red/black and black-height invariants, deterministic mixed insertions, duplicates,
  and the 100,000-key case under ASan/UBSan. Checks remain active under `NDEBUG`.
- **6 runner regression tests** passed, covering malformed, missing, duplicated
  and incorrect output, a total with the wrong unit, and an actual CLI process that
  must exit 1 when its benchmark returns a wrong value.
- Each of the three rebuilt binaries passed all **14 smoke results**. All nine
  subsequent benchmark processes also passed result and total validation.

Evidence: [native validation](native-validation.log),
[runner regression validation](runner-validation.log),
[compiled smoke](pure-smoke.log), [FP smoke](ffi-smoke.log),
[handwritten smoke](fficc-smoke.log).

The native checks cover the documented workloads and selected small inputs, not
all possible machine integers. ASan/UBSan success is not a complete proof of memory
safety or leak freedom. Some transitively imported library namespaces have no FFI
provider; they are listed in each build manifest, receive no fabricated implementation,
and were not called by the validated suite. Extended C++ library support remains WIP.

## Measurement protocol and results

Environment: macOS 26.1 arm64, Apple Clang 17.0.0 (`clang-1700.3.19.1`), C++11,
`-O3 -DNDEBUG`, no LTO. This legacy pscpp backend uses its PureScript 0.14.4 CoreFn
frontend. These measurements do not involve the separate Purust/TAST compiler.

All three binaries were built before measurement. No builds or validation test
processes were run concurrently with timings. The execution orders were:

1. compiled, FP, handwritten;
2. FP, handwritten, compiled;
3. handwritten, compiled, FP.

Each process uses the existing `Bench.purs` protocol: three global warm-ups, three
per-test warm-ups, then the best of ten measured iterations. Each README row is the
median of its three process results. A column total sums its row medians; it is
not process wall time or necessarily the median of the three process totals.

| Column | Sum of row medians | Individual process totals |
| --- | ---: | --- |
| Compiled pscpp | 950.99520 ms | 949.70372 / 964.71191 / 946.24903 ms |
| Native FP | 133.22988 ms | 134.12134 / 133.32732 / 133.25883 ms |
| Native handwritten | 21.33117 ms | 19.11641 / 21.32934 / 21.89203 ms |

These durations include the common PureScript harness, opaque input call and
result conversion. Sub-microsecond rows are close to harness and clock overhead;
they do not by themselves prove the elimination of an algorithm. The handwritten
tree accounts for about 21.27 ms of its 21.33 ms total. Its variation across runs
is visible in the raw data; three runs are not a confidence-interval estimate.

Full row data and ordering are in [results.json](results.json). Its `archived_log`
fields name all nine raw logs in this directory. Build manifests:
[compiled](pure-build.json), [FP](ffi-build.json), [handwritten](fficc-build.json).

## Reproduce

From the repository root, with the local pscpp binary, PureScript 0.14.4 frontend
and a real [cpp-ffi checkout](https://github.com/purescript-native/cpp-ffi) available:

```sh
bin/cpp/run --build-only
bin/cpp/run --ffi --build-only
bin/cpp/run --fficc --build-only
bin/cpp/test-native
bin/cpp/test-native --sanitize
python3 test/cpp-runner.py
bin/cpp/run --run-only --smoke
bin/cpp/run --ffi --run-only --smoke
bin/cpp/run --fficc --run-only --smoke
bin/cpp/run --run-only
bin/cpp/run --ffi --run-only
bin/cpp/run --fficc --run-only
```

Default workspaces are `run/bak/cpp/modes/{pure,ffi,fficc}`. The driver accepts
`--purs`, `--pscpp`, `--cxx`, `--ffi-root`, `--dependency-root` and `--build-dir`
overrides. It uses the existing C++ dependency source cache when present, otherwise
fetches the configured Spago package set. The cpp-ffi checkout used here was at
commit `268567879eed9a1da8450f07277382a76c658a90`; consumed provider hashes are recorded
in the manifests. [measure.py](measure.py) documents the balanced collection and
aggregation, and refuses to overwrite this dated result set.
