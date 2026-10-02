# Qualify a Purust TAST-loading optimization

The isolated decoder diagnostic is built and measured with
`../json-diagnostic.py`. This directory contains the complete-backend check on
the same frozen **238-module gopurs-aff** corpus. Run it only after building and
freezing the candidate compiler, with no concurrent build or benchmark.

From `altbak.pub`:

```sh
python3 bin/benchmark/purust-tast/diagnostic.py build \
  --out var/benchmark/purust-tast-20261002/candidate-array-ann
python3 bin/benchmark/purust-tast/diagnostic.py measure \
  --baseline var/benchmark/purust-tast-20261002 \
  --candidate var/benchmark/purust-tast-20261002/candidate-array-ann \
  --corpus gopurs238 --out var/benchmark/purust-tast-20261002/array-ann-238
python3 bin/benchmark/purust-tast/diagnostic.py measure \
  --baseline var/benchmark/purust-tast-20261002 \
  --candidate var/benchmark/purust-tast-20261002/candidate-array-ann \
  --corpus fixture12 --out var/benchmark/purust-tast-20261002/array-ann-12

node bin/benchmark/purust-tast/qualify-compiler.mjs \
  var/benchmark/purust-tast-20261002/qualification --discard-stage1-cache
node bin/benchmark/purust-tast/qualify-aff.mjs \
  var/benchmark/purust-tast-20261002/qualification \
  var/benchmark/purust-tast-20261002/aff-qualification

node bin/benchmark/purust-tast/compare-aff.mjs \
  var/benchmark/gopurs-purust-aff-20261002 \
  var/benchmark/purust-tast-20261002/aff-before-after.json \
  before=var/benchmark/gopurs-purust-aff-20261002/compilers/purust/bin/purust-native \
  after=var/benchmark/purust-tast-20261002/qualification/compiler-final \
  js=var/benchmark/gopurs-purust-aff-20261002/compilers/purust/bin/purust.js

node ../purust/purust/tools/compare-native.mjs \
  var/benchmark/purust-pipeline-20261002/reference \
  var/benchmark/purust-tast-20261002/purust-aff-before-after.json \
  before=var/benchmark/gopurs-purust-aff-20261002/compilers/purust/bin/purust-native \
  after=var/benchmark/purust-tast-20261002/qualification/compiler-final \
  js=var/benchmark/purust-tast-20261002/qualification/compiler-final.js
```

The Rust-only candidate build snapshots all fingerprinted input files and
retains separate timed and counting-allocator binaries. The diagnostic
before/after comparison alternates the frozen executables, checks the same
driver/profile and validates every fingerprint against the frozen oracle.
Both the historical statistic and the median of all 15 samples are recorded;
result destruction stays separate. Build before starting the measured runs.

`qualify-compiler.mjs` uses the production `tools/build-native.mjs --self-host`
workflow. It retains the source/FFI snapshots, stage-1 and stage-2 binaries,
fresh smoke workspace and logs. The resulting compiler is placed in the new
archive directory; the installed native executable is checked unchanged.
`PURUST_PURS` may explicitly select the typed frontend, which is then copied
and hashed. The production JS bundle is rebuilt in its checkout. Embedded
runtime sources are refreshed before their snapshot is taken.
The optional `--discard-stage1-cache` frees this build's stage-1 Cargo cache
after stage 2 and the smoke pass; it retains both executables, all generated
sources and logs, and records the cleanup in `qualification.json`.
`qualify-aff.mjs` copies the Aff project to a fresh sibling-package layout and
runs its existing complete `bin/test` with the qualified compiler: 47 printed
checks, Rust unit tests, concurrency/lifetime tests and nine error-reporting
scenarios. Logs and fresh TAST/application outputs are retained there.

The runner makes a new, integrity-checked copy of the TAST, PureScript sources,
FFI adapters and compiler binaries. It preserves the previous snapshot and
refuses to overwrite a campaign. Each variant receives one warm-up and five
measured runs, executed sequentially with rotating first variant. Timings use
the compiler's `backend total`; all phase samples and process logs are retained.
Every generated Rust source and Cargo manifest must match the original,
previously qualified output byte for byte. Input and binary hashes are checked
again after the campaign.
`PURUST_BENCH_RUNS` can set a predeclared longer confirmation. This lot uses
`PURUST_BENCH_RUNS=21` for `aff-confirmation-21.json` with only `before` and
`after`, after observing bimodal speculative-PBO work in its first five-run
campaign. The latter remains in `aff-before-after.json`. The extended result
also records the mean and all paired deltas, alongside the median.
The `before` and optional `js` variants must match the compiler hashes in the
original campaign. When `after` has an adjacent `qualification.json`, its hash
must match that qualified stage 2. A configuration JSON can provide explicit
`sha256` and `qualification` fields; the latter is relative to that config.

The isolated diagnostic's historical statistic (median of three process minima)
differs from this complete-compiler statistic (median of five process runs).
Result destruction and process shutdown are excluded from `backend total`.

This output comparison complements the candidate's native decoder differential
tests, full `purust-aff` execution and JS → native stage 1 → stage 2 → independent
smoke qualification. The original gopurs-aff application has known
scheduling-sensitive assertions, documented in `../gopurs-aff/README.md`; the
frozen timing corpus remains the original one.

The second complete-backend command confirms the change on the separately
frozen **244-module purust-aff** corpus. Compare before and after within each
campaign: the two corpora have different inputs and generated file counts.

`publish.mjs CAMPAIGN_DIRECTORY` verifies the retained input/source/binary hashes,
all process fingerprints, allocation counters, medians, qualifications and
generated outputs, then writes the public JSON and an archival verification
record. It requires `installation.json` to bind the installed native executable
to the measured, qualified stage 2.
