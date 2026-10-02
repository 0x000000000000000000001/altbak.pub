# Compile the same gopurs-aff TAST with Go and Rust backends

This local harness freezes the installed **gopurs** and **purust** launchers,
JavaScript bundles and native executables. It generates one TAST from
`gopurs-aff`, then measures the four backend variants sequentially. The primary
metric is the compiler's `backend total`, excluding `purs`, process startup/exit,
`go build`, Cargo and application execution.

Run from `altbak.pub` in the sibling-repository `htdocs` layout:

```sh
W=var/benchmark/gopurs-purust-aff-NEW
node bin/benchmark/gopurs-aff/prepare.mjs "$W"
node bin/benchmark/gopurs-aff/prepare-ffi.mjs "$W"
node bin/benchmark/gopurs-aff/qualify.mjs "$W" gopurs-native qualify-go-1
node bin/benchmark/gopurs-aff/qualify.mjs "$W" purust-native qualify-rust-o3 --release
```

`qualify.mjs` builds the generated application and checks the 45 printed Aff
results, allowing the first independent fibers to print in any order. The final
AVar stress test has an additional assertion and no success message. A failed
application exits the harness nonzero; its generated files and logs are retained.
In the October 2 campaign, the original suite was unstable on Rust. The two
development builds (including the historical label `qualify-rust-release`,
which was invoked **without** `--release`) and the O3 build all have archived
preflight logs.

Use a fixed five-run diagnostic batch, retaining successes and failures:

```sh
node bin/benchmark/gopurs-aff/check-apps.mjs "$W" application-diagnostic \
  "$W/generated/qualify-go-1/aff-test" "$W/cargo-target/release/purust_output"
```

## Rust FFI compatibility and application validation

`prepare-ffi.mjs` snapshots the existing Purust library implementations and
adapts six modules to gopurs' PureScript interfaces. These are local benchmark
adapters: Aff callbacks/native fibers, AVar constructor arguments, ST names,
Traversable's extra append argument, and atomic `Ref.modify_`. Three unused FFI
modules have explicit panic traps rather than silent fallback values. This
qualifies the exercised Aff program, not every API in those packages.

The original Aff tests include assumptions about startup order and closely
spaced timer races. `prepare-validation.mjs` creates a **separate** validation
workspace: ten tests use the synchronization-aware versions already in
`purust-aff`; the bracket test reads its result after completion instead of a
40 ms guess. Both backends receive the same validation sources. The timed TAST
is untouched.

```sh
node bin/benchmark/gopurs-aff/prepare-validation.mjs "$W"
V="$W/portable-validation"
node bin/benchmark/gopurs-aff/qualify.mjs "$V" gopurs-native qualify-go
node bin/benchmark/gopurs-aff/qualify.mjs "$V" purust-native qualify-rust --release
node bin/benchmark/gopurs-aff/check-apps.mjs "$V" application-checks \
  "$V/generated/qualify-go/aff-test" "$V/cargo-target/release/purust_output"
```

The portable validation batch must pass all ten executions before measurement.
Original-suite failures remain part of the published validation record.

## Measurement and verification

```sh
node bin/benchmark/gopurs-aff/measure.mjs "$W" \
  "$W/logs/qualify-go-1.json" "$W/logs/qualify-rust-o3.json" \
  "$V/qualify-go-qualification.json" "$V/qualify-rust-qualification.json"
node bin/benchmark/gopurs-aff/verify.mjs "$W"
```

There is one warm-up per variant and five measured rounds, rotating the first
variant. Every run starts a fresh compiler process with fresh generated output
and no `.purmeta`/`.cache`; the OS file cache is not flushed. The native launchers'
production worker/GC defaults are retained. All Go outputs must equal their
original-corpus reference byte for byte; the same comparison applies to all
Rust outputs. Cross-language source files are not compared to each other.

Use a relative `--ffi-dir`: the current JavaScript resolver joins it to the
working directory, while the native resolver also accepts an absolute path.
The first October 2 warm-up campaign exposed this difference and was archived
as `ffi-path-diagnostic.json`; none of its samples enter the final medians.

The workspace retains inputs, compiler artifacts, generated sources, manifests,
application binaries, logs, failed diagnostics, all timing samples and the
`verification.json` integrity check. Run builds and campaigns in the background,
and run the final timing campaign without concurrent builds or other benchmarks.
