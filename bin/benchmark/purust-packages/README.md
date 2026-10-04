# Purust compilation-table campaign

Run commands from the `altbak.pub` repository. Frontend preparation, compiler
qualification, measurements and application builds run **serially**. Each
script writes its logs and evidence to the archive; existing result files are
never replaced by a retry.

## Initial inventory

```sh
node bin/benchmark/purust-packages/libraries.mjs var/benchmark/purust-packages-20261004
```

This freezes the 57 local native ports, the typed frontend and both installed
Purust hosts. It prepares each library's primary test input, then uses the
shared `compilation-refresh/measure.mjs` runner for one warmup and five measured
runs per host. Each generation must match the canonical Rust/Cargo output.

## Foreign-type scanning correction and complete retry

The initial campaign exposed a native regex backtracking limit, a duplicate
registry/local Promise module in the isolated workspace, and a lost negative
zero in the native PBO evaluator. The first failure's source is retained as
`diagnostics/foreign-types-before.rs`; Prelude retains both generated outputs.
Spec Discovery and Yoga JSON also expose host-dependent module-root ordering;
Spec requires Spago's generated BuildInfo module. Those diagnostics remain in
the initial campaign, which finishes at 51/57 qualified library cases.

After that inventory has finished, the following command serializes the entire
correction, measurement, validation and publication-candidate sequence:

```sh
node bin/benchmark/purust-packages/finish.mjs var/benchmark/purust-packages-20261004
```

The individual stages are also available as follows:

```sh
node bin/benchmark/purust-packages/qualify-fix.mjs var/benchmark/purust-packages-20261004
node bin/benchmark/purust-packages/prepare.mjs \
  var/benchmark/purust-packages-20261004/revision1 \
  var/benchmark/purust-packages-20261004 \
  var/benchmark/purust-packages-20261004/bootstrap/qualification.json
node bin/benchmark/purust-packages/libraries.mjs var/benchmark/purust-packages-20261004/revision1
```

Qualification retains the expected failing regressions, fixed JS/Rust
differential checks, native bootstrap identity and fresh-project signed-zero
smoke tests, plus module-order invariance over 720 permutations.
The retry uses frozen package sources from the initial archive. Earlier output
is also used as the oracle whenever it exists.

The archived retry exposed a harness mistake: backend-free workspace YAML no
longer carried `--threaded`. Case preparation now reads each original frozen
runner plan. `recover-flags.mjs` compares every actual invocation with that plan,
re-measures only mismatches in new case directories, and writes
`libraries-final-results.json`; the original results remain intact. A resumed
pipeline can start at that step with a new attempt label:

```sh
node bin/benchmark/purust-packages/finish.mjs var/benchmark/purust-packages-20261004 finish3 recover-flags
```

## b8x, application validation and publication

Set `ARCHIVE` to the qualified campaign directory, then run:

```sh
node bin/benchmark/purust-packages/prepare-b8x.mjs "$ARCHIVE"
node bin/benchmark/compilation-refresh/measure.mjs "$ARCHIVE" b8x "$ARCHIVE/cases/b8x/definition.json"
node bin/benchmark/purust-packages/recover-flags.mjs "$ARCHIVE"
node bin/benchmark/purust-packages/validate.mjs "$ARCHIVE" libraries-final-results.json
node bin/benchmark/purust-packages/validate-b8x.mjs "$ARCHIVE"
node bin/benchmark/purust-packages/publish.mjs "$ARCHIVE"
```

The b8x case uses the existing Rust Core/Infra/Util test profile's real import
closure. Its corpus differs from the full Go-target row. Library validation
builds the identical output once and reuses each primary executable's existing
output/argument checks; separate secondary mains are not part of that case.
The b8x application is built outside timing; service-dependent tests remain in
the dedicated runtime profile.

The publisher checks samples, medians, input/output manifests, compilers and
application evidence. It writes the public report and an archived
`README-next.md` candidate. Review and apply that candidate, then verify the
table and totals against `publication.json`.

After applying the reviewed candidate, audit all 116 project cells, four total
cells and 60 ratios against the retained phase logs and measurement records:

```sh
node bin/benchmark/purust-packages/audit-publication.mjs "$ARCHIVE"
```

The audit also checks the published reports and the preceding README sections,
and records the applied README hash in `publication-audit.json`.

## Primed HTTP handle correction

The first application pass found a real native HTTP panic. Foreign-declaration
scanning truncated `HttpServer'`, and the native-carrier lookup used a Rust-mangled
name against source-level layout keys. The correction is qualified with red/green
JS/Rust scans, Rc/Arc carrier checks, fresh native smoke, bootstrap identity and
the actual HTTP suite before re-measuring every project with the corrected hosts:

```sh
node bin/benchmark/purust-packages/finish-primed.mjs var/benchmark/purust-packages-20261004
```

This serialized pipeline produces `revision2`. It copies the exact frontend
inputs from the completed previous series, keeps historical generated-output
oracles for cases without primed foreign declarations, and retains both outputs
for affected cases. Existing successful application evidence is reused only
when the new canonical Rust/Cargo sources match byte for byte. Failed attempts
remain intact. A new attempt label and starting stage can resume the pipeline.
When restarting qualification, also supply a new bootstrap directory label:

```sh
node bin/benchmark/purust-packages/finish-primed.mjs var/benchmark/purust-packages-20261004 finish5 qualify-primed primed-bootstrap2
```

If both compiler stages built and matched but a smoke-test fixture failed, a
final workspace-path argument reuses those retained stage binaries, verifies
their source identity again and runs a fresh smoke in the new qualification
directory before proceeding to the HTTP application gate.

Historical-output comparisons then rejected seven cases because the first
primed scan omitted the Unicode kind separator `∷`. The final qualification
compares all frozen declaration sources with the pre-prime scanner, recognizes
both kind separators, and runs the real BigInt suite as well as HTTP. A fresh
cohort can be selected with the final argument (`-` skips bootstrap reuse):

```sh
node bin/benchmark/purust-packages/finish-primed.mjs var/benchmark/purust-packages-20261004 finish7 qualify-primed unicode-bootstrap - revision3
```
