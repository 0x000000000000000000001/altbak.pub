// Verify retained measurements and write the compact, reproducible public data.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';
import { compareGeneratedSources } from '../../../../purust/purust/tools/native-workspace.mjs';

const [archiveArg] = process.argv.slice(2);
assert(archiveArg && process.argv.length === 3, 'Usage: publish.mjs CAMPAIGN_DIRECTORY');
const archive = resolve(archiveArg);
const read = path => JSON.parse(readFileSync(join(archive, path), 'utf8'));
const median = values => values.toSorted((a, b) => a - b)[Math.floor(values.length / 2)];
const close = (actual, expected) => assert(Math.abs(actual - expected) < 1e-6, `${actual} != ${expected}`);
const verified = [];
function file(path, sha256) {
  assert.equal(hash(readFileSync(path)), sha256, path);
  verified.push({ path, sha256 });
}
const artifacts = read('artifacts/provenance.json');
for (const entry of artifacts.frozen) file(join(archive, entry.copied), entry.sha256);
const candidate = read('candidate-array-ann/build.json');
assert.equal(candidate.status, 'passed');
for (const entry of [candidate.timed, candidate.alloc]) file(entry.path, entry.sha256);
const sourceRoot = fileURLToPath(new URL('../../../../', import.meta.url));
for (const [path, sha256] of Object.entries(candidate.sources)) {
  const name = relative(sourceRoot, path);
  assert(!name.startsWith('..'));
  file(join(archive, 'candidate-array-ann/sources', name), sha256);
}
const compiler = read('qualification/qualification.json');
assert.equal(compiler.status, 'passed');
for (const entry of [compiler.stage1, compiler.stage2, compiler.javascript, compiler.frontend]) file(entry.path, entry.sha256);
for (const [index, source] of compiler.sources.entries()) for (const entry of source.files) {
  file(join(archive, 'qualification/sources', String(index), entry.path), entry.sha256);
  if (source.path === compiler.pbo) assert.equal(entry.sha256,
    candidate.sources[join(source.path, entry.path)] ?? entry.sha256,
    `Optimizer differs between isolated diagnostic and compiler: ${entry.path}`);
}
const aff = read('aff-qualification/qualification.json');
assert.equal(aff.status, 'passed');
assert.equal(aff.compiler_sha256, compiler.stage2.sha256);
for (const entry of aff.inputs) file(join(aff.workspace, 'purust-aff', entry.path), entry.sha256);
assert.equal(aff.commands[0].exit_code, 0);
const affLog = readFileSync(aff.commands[0].stdout, 'utf8');
assert.match(affLog, /Summary: 47 Aff checks/);
assert.match(affLog, /9 error-reporting scenarios passed\./);

function isolated(name) {
  const result = read(name + '/results.json');
  assert(result.validation.startsWith('passed'));
  assert.equal(result.runs.length, 6);
  file(join(archive, name, 'corpus.json'), result.corpus_sha256);
  file(join(archive, name, 'oracle.json'), result.oracle_sha256);
  const oracle = read(name + '/oracle.json');
  for (const entry of Object.values(result.binaries)) file(entry.path, entry.sha256);
  assert.equal(result.binaries.after.sha256, candidate.timed.sha256);
  for (const label of ['before', 'after']) {
    const runs = result.runs.filter(run => run.variant === label);
    assert.equal(runs.length, 3);
    for (const run of runs) {
      assert.equal(run.status, 0);
      assert.deepEqual(read(name + '/' + run.stdout), run.result);
      for (const key of ['modules', 'fingerprints', 'json_fingerprints']) assert.deepEqual(run.result[key], oracle[key]);
    }
    for (const phase of ['parse', 'decode', 'combined']) {
      for (const run of runs) close(run.result.phases[phase].time_us,
        Math.min(...run.result.phases[phase].samples.map(sample => sample.time_us)));
      close(result.summary[label][phase].median_process_minimum_us, median(runs.map(run => run.result.phases[phase].time_us)));
      close(result.summary[label][phase].median_all_samples_us,
        median(runs.flatMap(run => run.result.phases[phase].samples.map(sample => sample.time_us))));
      close(result.summary[label][phase].median_drop_us,
        median(runs.flatMap(run => run.result.phases[phase].samples.map(sample => sample.drop_us))));
    }
  }
  for (const phase of ['parse', 'decode', 'combined']) close(result.after_over_before[phase],
    result.summary.after[phase].median_process_minimum_us / result.summary.before[phase].median_process_minimum_us);
  return result;
}
const diagnostic12 = isolated('array-ann-12'), diagnostic238 = isolated('array-ann-238');
const backend238 = read('aff-before-after.json'), backend244 = read('purust-aff-before-after.json');
assert.equal(backend238.validation.status, 'passed');
assert(!backend244.failure && backend244.finished_at);
assert.equal(backend238.runs.length, 18); assert.equal(backend244.runs.length, 18);
let generatedFiles = 0;
for (const variant of backend238.variants) {
  file(variant.binary, variant.sha256);
  const runs = backend238.runs.filter(run => !run.warmup && run.variant === variant.label);
  close(backend238.summary[variant.label].median_ms, median(runs.map(run => run.phases_ms['backend total'])));
  if (variant.label === 'after') assert.equal(variant.sha256, compiler.stage2.sha256);
}
for (const directory of ['inputs', 'rust-ffi']) {
  const entries = directory === 'inputs' ? backend238.provenance.frozen_inputs : backend238.provenance.frozen_ffi;
  for (const entry of entries) file(join(backend238.workspace, directory, entry.path), entry.sha256);
}
for (const run of backend238.runs) {
  assert.equal(run.exit_code, 0);
  const phases = Object.fromEntries([...readFileSync(run.stderr, 'utf8').matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
    .map(([, name, ms]) => [name, Number(ms)]));
  assert.deepEqual(run.phases_ms, phases);
  const entries = manifest(run.output, path => /\.(rs|toml)$/.test(path));
  assert.deepEqual(entries, backend238.provenance.reference_files);
  assert.equal(entries.length, run.identical_generated_files);
  generatedFiles += entries.length;
}
const snapshot244 = JSON.parse(readFileSync(join(backend244.snapshot, 'manifest.json'), 'utf8'));
for (const entry of snapshot244.inputs.files) file(join(backend244.snapshot, entry.path), entry.sha256);
const runs244 = join(archive, 'purust-aff-before-after-runs');
let reference244;
for (const variant of backend244.compilers) {
  file(variant.binary, variant.sha256);
  const runs = backend244.runs.filter(run => !run.warmup && run.compiler === variant.label);
  close(backend244.summary[variant.label].phases_median_ms['backend total'], median(runs.map(run => run.phases_ms['backend total'])));
  if (variant.label === 'after') assert.equal(variant.sha256, compiler.stage2.sha256);
}
for (const run of backend244.runs) {
  assert.equal(run.status, 0);
  const phases = Object.fromEntries([...run.stderr.matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
    .map(([, name, ms]) => [name, Number(ms)]));
  assert.deepEqual(run.phases_ms, phases);
  const output = join(runs244, `${run.round}-${run.compiler}`);
  reference244 ??= output;
  const count = compareGeneratedSources(reference244, output);
  assert.equal(count, run.identical_files); generatedFiles += count;
}
const confirmation = read('aff-confirmation-21.json');
assert.equal(confirmation.protocol.measured_rounds, 21);
assert.equal(confirmation.validation.status, 'passed');
assert.equal(confirmation.runs.length, 44);
const serial = read('aff-serial-before-after.json');
assert.equal(serial.validation.status, 'passed');
assert.equal(serial.runs.length, 12);
for (const campaign of [confirmation, serial]) {
  assert.deepEqual(campaign.provenance.tast, backend238.provenance.tast);
  for (const variant of campaign.variants) {
    file(variant.binary, variant.sha256);
    assert.equal(variant.sha256, backend238.variants.find(entry => entry.label === variant.label).sha256);
    const runs = campaign.runs.filter(run => !run.warmup && run.variant === variant.label);
    close(campaign.summary[variant.label].median_ms, median(runs.map(run => run.phases_ms['backend total'])));
  }
  for (const run of campaign.runs) {
    assert.equal(run.exit_code, 0);
    const phases = Object.fromEntries([...readFileSync(run.stderr, 'utf8').matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
      .map(([, name, ms]) => [name, Number(ms)]));
    assert.deepEqual(run.phases_ms, phases);
    const files = manifest(run.output, path => /\.(rs|toml)$/.test(path));
    assert.deepEqual(files, backend238.provenance.reference_files);
    generatedFiles += files.length;
  }
  for (const directory of ['inputs', 'rust-ffi']) {
    const entries = directory === 'inputs' ? campaign.provenance.frozen_inputs : campaign.provenance.frozen_ffi;
    for (const entry of entries) file(join(campaign.workspace, directory, entry.path), entry.sha256);
  }
}
const ablation = read('ablation2/results.json');
assert.equal(ablation.status, 'passed'); assert.equal(ablation.restored, true);
assert.equal(ablation.runs.length, 20);
file(ablation.binary.path, ablation.binary.sha256);
file(ablation.source, ablation.source_sha256);
file(join(archive, 'ablation2/instrumented-json.rs'), ablation.instrumented_sha256);
for (const run of ablation.runs) {
  assert.equal(run.exit_code, 0);
  const files = manifest(join(archive, 'ablation2/generated', run.label), path => /\.(rs|toml)$/.test(path));
  assert.deepEqual(files, backend238.provenance.reference_files);
  generatedFiles += files.length;
}
const codegen = readFileSync(join(archive, 'logs/codegen.log'), 'utf8');
const retry = readFileSync(join(archive, 'logs/codegen-docker-retry2.log'), 'utf8');
assert.match(codegen, /pass 90\b/); assert.match(codegen, /fail 4\b/);
assert.match(retry, /pass 4\b/); assert.match(retry, /fail 0\b/);
const annotationTests = readFileSync(join(archive, 'logs/array-ann.log'), 'utf8').split('\n').find(line => line.startsWith('Native TAST:'));
assert(annotationTests?.includes('110909 native, 0 errors, 0 declined'));
const extendedTests = readFileSync(join(archive, 'logs/native-tast-extended.log'), 'utf8');
assert(extendedTests.includes('52 numeric annotations / 16 packed arrays passed'));
assert(extendedTests.includes('110909 native, 0 errors, 0 declined'));
for (const [suite, count] of [['json-fields', 7], ['source-usage', 12], ['numeric-negate', 5]]) {
  const log = readFileSync(join(archive, 'logs', `pbo-${suite}.log`), 'utf8');
  assert.match(log, new RegExp(`pass ${count}\\b`)); assert.match(log, /fail 0\b/);
}
const typeTests = readFileSync(join(archive, 'logs/type-table.log'), 'utf8');
assert.match(typeTests, /exact errors and shared references passed/);
const baseline = {};
for (const name of ['campaign-12', 'campaign-238']) {
  const data = read(name + '/results.json');
  const { build, ...measurements } = data;
  const oracle = read((name === 'campaign-12' ? 'array-ann-12' : 'array-ann-238') + '/oracle.json');
  for (const [runtime, processes] of Object.entries(data.results)) {
    assert.equal(processes.length, 3);
    for (const process of processes) for (const key of ['modules', 'fingerprints', 'json_fingerprints']) {
      assert.deepEqual(process[key], oracle[key]);
    }
    for (const phase of ['parse', 'decode', 'combined']) close(data.medians_us[runtime][phase],
      median(processes.map(process => process.phases[phase].time_us)));
  }
  baseline[name] = { ...measurements, build_artifacts: { go_sha256: build.binary_sha256,
    javascript_sha256: build.js_sha256, rust_sha256: build.rust.binary_sha256, rust_profile: build.rust.profile } };
}
const allocations = {};
for (const [label, directory] of [['before', 'alloc-238'], ['after', 'array-ann-alloc-238']]) {
  const raw = read(directory + '/alloc-parse_decode_combined.json');
  const summary = read(directory + '/alloc-parse_decode_combined-summary.json');
  const oracle = read('array-ann-238/oracle.json');
  assert.equal(raw.alloc_diag, true);
  for (const key of ['modules', 'fingerprints', 'json_fingerprints']) assert.deepEqual(raw[key], oracle[key]);
  for (const [phase, data] of Object.entries(raw.phases)) {
    for (const key of ['allocations', 'allocated_bytes', 'deallocations', 'deallocated_bytes',
      'drop_allocated_bytes', 'drop_deallocated_bytes']) {
      assert.deepEqual(summary.phases[phase][key], data.samples.map(sample => sample[key]));
    }
    for (const key of ['allocations', 'allocated_bytes', 'drop_deallocated_bytes']) close(
      summary.phases[phase]['median_' + key], median(data.samples.map(sample => sample[key])));
  }
  allocations[label] = { ...summary, raw };
}
const jsonDecoding = read('json-decoding-regression-results/results.json');
const jsonDecodingOracle = JSON.parse(readFileSync(fileURLToPath(new URL('../../../test/fixtures/json-decoding/expected.json', import.meta.url)), 'utf8'));
for (const processes of Object.values(jsonDecoding.results)) {
  assert.equal(processes.length, 3);
  for (const process of processes) for (const key of ['modules', 'names', 'timed_cases', 'fingerprints', 'json_fingerprints']) {
    assert.deepEqual(process[key], jsonDecodingOracle[key]);
  }
}
const installationPath = join(archive, 'installation.json');
assert(existsSync(installationPath), 'Install the qualified, measured stage 2 before publishing final data');
const installation = read('installation.json');
assert.equal(installation.sha256, compiler.stage2.sha256); file(installation.destination, installation.sha256);
const verification = { status: 'passed', verified_at: new Date().toISOString(),
  frozen_files_checked: verified.length, generated_outputs: 112, generated_files_checked: generatedFiles,
  diagnostics: 'Every process fingerprint and historical median recomputed',
  compiler: 'Measured and installed binary equals the qualified stage 2', files: verified };
writeJson(join(archive, 'verification.json'), verification);
const publication = {
  schema: 1, title: 'Rust JSON to Typed AST and complete Aff compilation', archive,
  baseline, diagnostic_before_after: { fixture12: diagnostic12, gopurs238: diagnostic238 },
  allocations,
  profiles_before: { combined: read('profile-238/sample-combined-summary.json'), decode: read('profile-238/sample-decode-summary.json') },
  backend_before_after: { gopurs238: backend238, purust244: backend244 },
  confirmation_gopurs238: confirmation,
  scheduling_diagnostics: { serial_gopurs238: serial, same_binary_ablation: ablation },
  qualification: { compiler: { ...compiler, sources: undefined }, aff: { ...aff, inputs: undefined },
    annotation_tests: annotationTests, extended_annotation_tests: extendedTests.trim(),
    pbo_contract_tests: { json_fields: 7, source_usage: 12, numeric_negate: 5, type_table_script: 'passed' },
    type_table_tests: typeTests.trim(), codegen: { passed: 94,
      initial_passed: 90, initial_environment_failures: 4, retry_passed: 4,
      note: 'Docker service and then its existing test container were stopped; all failed attempts are retained.' },
    json_decoding_regression: { cases: 17, processes: 9, status: 'passed' } },
  installation, verification: { ...verification, files: undefined },
};
const destination = fileURLToPath(new URL('../../../docs/benchmark-results/2026-10-02-rust-json-typed-ast.json', import.meta.url));
writeJson(destination, publication);
console.log(JSON.stringify({ destination, verification: publication.verification }, null, 2));
