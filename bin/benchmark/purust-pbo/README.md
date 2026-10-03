# Purust compiler experiments

The campaign archive is `var/benchmark/purust-pbo-20261002/`. Its installed
reference is the qualified TAST stage 2 (`002ed7b5…`). All compiler experiments
use the original frozen 238-module `gopurs-aff` corpus and its 484-file exact
Rust output oracle. The separate 244-module corpus has its own oracle.

## Commands (from `altbak.pub`)

```sh
node bin/benchmark/purust-pbo/experiment.mjs init NEW_ARCHIVE
node bin/benchmark/purust-pbo/experiment.mjs build ARCHIVE LABEL SELECTION.json
node bin/benchmark/purust-pbo/check-pbo.mjs ARCHIVE/work/js/output NEW_TEST_ARCHIVE

node bin/benchmark/purust-tast/compare-aff.mjs ORIGINAL_AFF RESULTS.json \
  before=ARCHIVE/before.json after=ARCHIVE/candidates/LABEL/purust-native

node bin/benchmark/purust-pbo/compare-go-rust.mjs ORIGINAL_AFF NEW_ARCHIVE RUST_VARIANTS.json
```

The source-selection JSON has optional `base` (relative to the JSON, default:
the initial frozen sources) and `overrides` entries of the form
`{ "root": 1, "path": "src/…", "from": "relative/source" }`.
Root 0 is Purust; root 1 is its optimizer fork. Selection is copied before the
build, so independent live-source changes do not enter a candidate. Separate
JS/native Spago workspaces preserve their respective library selections.
Candidate bundles, Rust sources, commands and executables are archived; a
single reusable Cargo workspace avoids redundant rebuilds. Builds are O3,
without LTO, with threaded Arc and mimalloc. Builds and measurements must run
sequentially; long commands run in the background with durable logs.

Go/Rust configurations are arrays of Rust variants:
`[{ "label": "rust-native", "binary": "relative/path", "qualification": "relative/qualification.json" }]`.
`javascript: true` selects a JS bundle; optional `env` records explicit native
worker settings. The Go variants use the original frozen production launcher
and binaries. Each target must match its own historical byte oracle.

## Diagnostics

```sh
python3 bin/benchmark/purust-pbo/profile.py BINARY COMPARE_AFF_RUNS NEW_ARCHIVE
node bin/benchmark/purust-pbo/instrument.mjs GENERATED_RUST NEW_ARCHIVE COMPARE_AFF_RUNS

python3 bin/benchmark/purust-tast/diagnostic.py build --out NEW_JSON_BUILD
python3 bin/benchmark/purust-pbo/compare-json.py \
  --reference var/benchmark/purust-tast-20261002 \
  --before var/benchmark/purust-tast-20261002/candidate-array-ann \
  --after NEW_JSON_BUILD --corpus fixture12 --out NEW_JSON_RESULTS
```

Use `--corpus gopurs238` for the larger JSON diagnostic. Its Go executable,
corpus and JS-authority fingerprints come from the original TAST archive.
Three processes per variant rotate both runtime and phase order; each process
uses two warmups and five timed samples per phase. The table statistic is the
median of the three process minima, with all samples retained. Rust final-result
destruction is recorded separately. Go uses one processor and `GOGC=100` for
this isolated diagnostic, distinct from the complete compiler's production
GC policy. JSON fingerprinting and file loading are outside the timed regions.

```sh
python3 bin/benchmark/purust-pbo/json-allocations.py JSON_CAMPAIGN NEW_ALLOCATIONS
python3 bin/benchmark/purust-pbo/resources.py TWO_VARIANT_COMPARE_AFF.json NEW_RESOURCES
node bin/benchmark/purust-pbo/check-native.mjs GENERATED_RUST NEW_TEST_ARCHIVE CHECKS.json
```

`resources.py` records whole-process user/system CPU and peak RSS separately
from `backend total`. `check-native.mjs` takes an array such as
`[{ "name": "test-native-record-names" }]`; optional `args` paths are relative
to the configuration file. It archives the tests and hashes the generated Rust
and Cargo inputs before and after the serialized checks.

`PURUST_PBO_JOBS` selects the **total budget**, not the final optimizer count.
For concurrent generation, `buildConcurrency` subtracts the codegen reservation
from that budget. Thus PBO 3 / generation 5 requires `PURUST_PBO_JOBS=8` and
`PURUST_CODEGEN_JOBS=5`. A single generation slot runs inline on the coordinator.
Check `PBO jobs`, `codegen-jobs` and `budget` in the retained logs.

`instrument.mjs` temporarily instruments generated Rust, then restores and
hash-verifies every changed file and the original executable. It counts
allocation requests/cumulative bytes, cache probes/hits, and accepted/rejected
conversion attempts with per-thread CPU. Its timings are diagnostic only.
Inclusive sampled profile families overlap; they are not CPU percentages.
The optimizer's normal `attempts-ms` sums concurrent task wall durations and
must not be interpreted as recoverable elapsed time.

## Final qualification

Retained changes require native differential regressions, full Aff/runtime
execution, JS → native stage 1 → native stage 2 source equality, and a fresh
project smoke. Use the `purust-tast/qualify-{compiler,aff}.mjs` harnesses.
`regressions.mjs QUALIFICATION NEW_ARCHIVE` runs the host-portable production
codegen/TAST Node suites with the qualified JavaScript bundle and frozen typed
frontend, and retains their test sources. Five existing Docker-only b8x FFI
fixtures are explicitly listed as unavailable; they are not reported as passed.
The native FFI checks and complete Aff suite use the qualified stage 2.
Measure the qualified stage-2 binary before installing its exact hash. Retain
failed diagnostics as well as successful results. Cargo caches may be removed
after retaining executables, generated sources and build logs.

`confirmation-protocol.json` in the campaign archive fixes the final round
counts and the treatment of inconclusive candidates before confirmation runs.
`publish.mjs ARCHIVE RESULT.json --verify-only` verifies the qualified hashes,
raw logs, JSON fingerprints and exact generated outputs before installation.
After installing the qualified executable, omit `--verify-only` to additionally
verify its installed hash and write the public data file. Result files must be new.
