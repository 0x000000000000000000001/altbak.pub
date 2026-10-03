// Verify final qualification and retained raw measurements before publication.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { basename, join, resolve } from 'node:path';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, destinationArg, ...flags] = process.argv.slice(2);
assert(archiveArg && destinationArg && flags.every(flag => flag === '--verify-only'),
  'publish.mjs ARCHIVE DESTINATION.json [--verify-only]');
const verifyOnly = flags.includes('--verify-only');
const archive = resolve(archiveArg), destination = resolve(destinationArg);
assert(!existsSync(destination), destination);
const median = values => { const a = values.toSorted((a, b) => a - b), m = a.length >> 1;
  return a.length % 2 ? a[m] : (a[m - 1] + a[m]) / 2; };
const close = (a, b) => assert(Math.abs(a - b) < 1e-6, `${a} != ${b}`);
const verified = [];
function load(name) {
  const path = join(archive, name), bytes = readFileSync(path);
  verified.push({ path, sha256: hash(bytes) });
  return JSON.parse(bytes);
}
function verify(path, sha256) { assert.equal(hash(readFileSync(path)), sha256, path); }

const compiler = load('qualification/qualification.json');
const aff = load('aff-qualification/qualification.json');
assert.equal(compiler.status, 'passed'); assert.equal(aff.status, 'passed');
assert.equal(aff.compiler_sha256, compiler.stage2.sha256);
for (const file of [compiler.stage1, compiler.stage2, compiler.javascript, compiler.frontend]) verify(file.path, file.sha256);
for (const [index, source] of compiler.sources.entries()) for (const file of source.files) {
  verify(join(archive, 'qualification/sources', String(index), file.path), file.sha256);
  verify(join(source.path, file.path), file.sha256);
}
assert.match(readFileSync(compiler.smoke_log, 'utf8'), /Native smoke: \d+ fresh TAST modules, \d+ identical generated files, application result verified\./);
assert.match(readFileSync(aff.commands[0].stdout, 'utf8'), /Summary: 47 Aff checks/);
assert.match(readFileSync(aff.commands[0].stdout, 'utf8'), /9 error-reporting scenarios passed\./);

let outputs = 0, generatedFiles = 0;
function verifyCampaign(data, common = false) {
  assert.equal(common ? data.status : data.validation.status, 'passed');
  const rounds = common ? data.protocol.rounds : data.protocol.measured_rounds;
  assert.equal(data.runs.length, data.variants.length * (rounds + 1));
  for (const variant of data.variants) {
    if (variant.family !== 'gopurs') verify(variant.binary, variant.sha256);
    const measured = data.runs.filter(r => r.variant === variant.label && !r.warmup);
    assert.equal(measured.length, rounds);
    assert.equal(new Set(measured.map(r => r.round)).size, rounds);
    close(data.summary[variant.label].median_ms, median(measured.map(r => r.phases_ms['backend total'])));
  }
  if (common) {
    for (const [key, folder] of [['inputs', 'inputs'], ['ffi', 'rust-ffi'], ['compilers', 'compilers'], ['harness', 'harness']])
      for (const file of data.frozen[key]) verify(join(archive, 'common-go-rust-final', folder, file.path), file.sha256);
  } else {
    for (const [key, folder] of [['frozen_inputs', 'inputs'], ['frozen_ffi', 'rust-ffi'], ['harness', 'harness']])
      for (const file of data.provenance[key]) verify(join(data.workspace, folder, file.path), file.sha256);
  }
  for (const run of data.runs) {
    assert.equal(run.exit_code, 0);
    const family = common ? run.family : 'purust';
    const phases = Object.fromEntries([...readFileSync(run.stderr, 'utf8').matchAll(
      new RegExp(`^\\[${family}\\] (.+): (\\d+) ms$`, 'gm'))].map(([, phase, ms]) => [phase, Number(ms)]));
    assert.deepEqual(phases, run.phases_ms);
    const files = manifest(run.output, path => /\.(rs|go|toml)$/.test(path) || path.endsWith('/go.mod'));
    assert.deepEqual(files, common ? data.references[family] : data.provenance.reference_files);
    assert.deepEqual(files, JSON.parse(readFileSync(run.generated_manifest, 'utf8')));
    outputs++; generatedFiles += files.length;
  }
  return { protocol: data.protocol, host: data.host, variants: data.variants,
    tast: common ? data.tast : data.provenance.tast, summary: data.summary,
    paired_after_minus_before_ms: data.paired_after_minus_before_ms,
    runs: data.runs.map(({ generated_manifest, ...run }) => run) };
}
const primary = load('aff-final-confirmation.json');
assert.equal(primary.protocol.measured_rounds, 15);
assert.equal(primary.variants.find(v => v.label === 'after').sha256, compiler.stage2.sha256);
const common = load('common-go-rust-final/results.json');
assert.equal(common.protocol.rounds, 10);
assert.equal(common.variants.find(v => v.label === 'rust-final').sha256, compiler.stage2.sha256);
assert.equal(common.variants.find(v => v.label === 'rust-js').sha256, compiler.javascript.sha256);
const compiler238 = verifyCampaign(primary), goRust = verifyCampaign(common, true);
const pairs = primary.runs.filter(r => r.variant === 'after' && !r.warmup).map(after =>
  after.phases_ms['backend total'] - primary.runs.find(r => r.variant === 'before' && r.round === after.round).phases_ms['backend total']);
assert.deepEqual(pairs, primary.paired_after_minus_before_ms.samples);
close(median(pairs), primary.paired_after_minus_before_ms.median);
assert.equal(pairs.filter(value => value < 0).length, primary.paired_after_minus_before_ms.after_faster_pairs);
const secondary = load('purust-aff-final.json');
assert(!secondary.failure && secondary.finished_at);
assert.equal(secondary.compilers.find(v => v.label === 'after').sha256, compiler.stage2.sha256);
for (const variant of secondary.compilers) {
  verify(variant.binary, variant.sha256);
  const runs = secondary.runs.filter(r => r.compiler === variant.label && !r.warmup);
  assert.equal(runs.length, 7);
  close(secondary.summary[variant.label].phases_median_ms['backend total'], median(runs.map(r => r.phases_ms['backend total'])));
}
const secondaryManifest = JSON.parse(readFileSync(join(secondary.snapshot, 'manifest.json'), 'utf8'));
for (const file of secondaryManifest.inputs.files) verify(join(secondary.snapshot, file.path), file.sha256);
const secondaryRuns = join(archive, 'purust-aff-final-runs');
const secondaryOracle = manifest(join(secondaryRuns, `0-${secondary.compilers[0].label}`), path => /\.(rs|toml)$/.test(path));
assert.equal(secondaryOracle.length, 496);
assert.deepEqual(secondaryOracle, manifest(join(secondaryManifest.workspace, 'generated/warmup-js'),
  path => /\.(rs|toml)$/.test(path)));
for (const run of secondary.runs) {
  assert.equal(run.status, 0);
  assert.deepEqual(Object.fromEntries([...run.stderr.matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
    .map(([, phase, ms]) => [phase, Number(ms)])), run.phases_ms);
  const files = manifest(join(secondaryRuns, `${run.round}-${run.compiler}`), path => /\.(rs|toml)$/.test(path));
  assert.deepEqual(files, secondaryOracle);
  outputs++; generatedFiles += files.length;
}

const jsonBuild = load('json-final-build/build.json');
assert.equal(jsonBuild.status, 'passed');
for (const file of [jsonBuild.timed, jsonBuild.alloc]) verify(file.path, file.sha256);
// Bind the isolated decoder FFI and PureScript sources to the qualified compiler.
for (const [index, source] of compiler.sources.entries()) if (source.path === compiler.pbo) {
  for (const file of source.files) {
    const key = join(source.path, file.path);
    if (jsonBuild.sources[key]) assert.equal(file.sha256, jsonBuild.sources[key], key);
  }
}
function jsonCampaign(name) {
  const result = load(name + '/results.json'); assert.equal(result.status, 'passed');
  assert.equal(result.binaries.after.sha256, jsonBuild.timed.sha256);
  verify(join(archive, name, 'corpus.json'), result.corpus_sha256);
  verify(join(archive, name, 'oracle.json'), result.oracle_sha256);
  const oracle = load(name + '/oracle.json');
  for (const binary of Object.values(result.binaries)) verify(binary.path, binary.sha256);
  for (const run of result.runs) {
    assert.equal(run.exit_code, 0);
    assert.deepEqual(load(name + '/' + run.stdout), run.result);
    for (const key of ['modules', 'fingerprints', 'json_fingerprints']) assert.deepEqual(run.result[key], oracle[key]);
  }
  for (const label of Object.keys(result.binaries)) {
    const runs = result.runs.filter(r => r.variant === label);
    assert.equal(runs.length, 3);
    for (const phase of ['parse', 'decode', 'combined']) {
      for (const run of runs) close(run.result.phases[phase].time_us, Math.min(...run.result.phases[phase].samples.map(s => s.time_us)));
      close(result.summary[label][phase].median_process_minimum_us, median(runs.map(r => r.result.phases[phase].time_us)));
    }
  }
  return result;
}
const json12 = jsonCampaign('json-final-12'), json238 = jsonCampaign('json-final-238');
const nativeTests = load('tests-final/results.json'), pboTests = load('pbo-tests-final/results.json');
assert.equal(nativeTests.status, 'passed'); assert.equal(pboTests.status, 'passed');
assert.equal(nativeTests.runs.length, 12); assert.equal(pboTests.runs.length, 12);
assert.equal(nativeTests.rust, join(compiler.workspace, 'rust-stage2'));
assert.deepEqual(nativeTests.generated, manifest(nativeTests.rust,
  path => !path.includes('/target/') && /\.(rs|toml)$/.test(path)));
for (const run of nativeTests.runs) {
  assert.equal(run.exit_code, 0);
  for (const input of run.inputs) {
    verify(input.path, input.sha256);
    verify(join(archive, 'tests-final/tests', basename(input.path)), input.sha256);
  }
}
for (const run of pboTests.runs) {
  assert.equal(run.exit_code, 0);
  verify(run.command[1], run.test_sha256);
  verify(join(archive, 'pbo-tests-final/tests', basename(run.command[1])), run.test_sha256);
}
const regressions = load('regressions-final/results.json'); assert.equal(regressions.status, 'passed');
assert.equal(regressions.qualified_binary.sha256, compiler.stage2.sha256);
for (const file of regressions.tests) {
  verify(join(archive, 'regressions-final', file.path), file.sha256);
  verify(join(compiler.root, file.path), file.sha256);
}
for (const run of regressions.runs) {
  assert.equal(run.exit_code, 0);
  assert.match(readFileSync(run.stdout, 'utf8'), /(?:#|ℹ) fail 0\b/);
}
const resources = load('resources-final-v2/results.json'); assert.equal(resources.status, 'passed');
assert.equal(resources.variants.find(v => v.label === 'after').sha256, compiler.stage2.sha256);
verify(resources.source_comparison, resources.source_sha256);
assert.equal(resources.runs.length, 8);
for (const run of resources.runs) {
  assert.equal(run.exit_code, 0);
  const log = readFileSync(join(archive, 'resources-final-v2', run.stderr), 'utf8');
  const timing = log.match(/(\d+\.\d+) real\s+(\d+\.\d+) user\s+(\d+\.\d+) sys/);
  assert(timing);
  close(run.wall_seconds, Number(timing[1])); close(run.user_seconds, Number(timing[2]));
  close(run.system_seconds, Number(timing[3])); close(run.cpu_seconds, Number(timing[2]) + Number(timing[3]));
  close(run.max_rss_bytes, Number(log.match(/(\d+)\s+maximum resident set size/)[1]));
  const files = manifest(join(archive, 'resources-final-v2/generated', `${run.round}-${run.label}`), path => /\.(rs|toml)$/.test(path));
  assert.deepEqual(files, primary.provenance.reference_files);
  outputs++; generatedFiles += files.length;
}
for (const [label, summary] of Object.entries(resources.summary)) for (const [key, value] of Object.entries(summary))
  close(value, median(resources.runs.filter(r => r.label === label && !r.warmup).map(r => r[key])));
const diagnostic = load('diagnostic-final/results.json'); assert.equal(diagnostic.status, 'passed'); assert(diagnostic.restored);
assert.equal(diagnostic.binary_sha256, compiler.stage2.sha256);
verify(join(archive, 'diagnostic-final/compiler-counted'), diagnostic.instrumented_sha256);
verify(join(nativeTests.rust, 'target/release/purust_output'), compiler.stage2.sha256);
const countedFiles = manifest(join(archive, 'diagnostic-final/generated'), path => /\.(rs|toml)$/.test(path));
assert.deepEqual(countedFiles, primary.provenance.reference_files);
outputs++; generatedFiles += countedFiles.length;
const jsonAllocations = load('json-allocations-238/results.json'); assert.equal(jsonAllocations.status, 'passed');
assert.equal(jsonAllocations.runs.after.binary.sha256, jsonBuild.alloc.sha256);
for (const [label, run] of Object.entries(jsonAllocations.runs)) {
  verify(run.binary.path, run.binary.sha256);
  assert.equal(run.exit_code, 0);
  assert.deepEqual(load('json-allocations-238/' + label + '.stdout'), run.raw);
  assert(run.raw.alloc_diag);
  for (const key of ['modules', 'fingerprints', 'json_fingerprints'])
    assert.deepEqual(run.raw[key], json238.runs[0].result[key]);
  for (const [phase, summary] of Object.entries(run.summary)) for (const [key, value] of Object.entries(summary))
    close(value, median(run.raw.phases[phase].samples.map(sample => sample[key])));
}
const nativePath = resolve(compiler.root, 'bin/purust-native');
if (!verifyOnly) verify(nativePath, compiler.stage2.sha256);
verify(resolve(compiler.root, 'bin/purust.js'), compiler.javascript.sha256);
const report = {
  schema: 1, date: new Date().toISOString(), benchmark: 'Purust native compiler: nocturnal PBO, code generation and TAST optimization',
  archive, qualified: { stage1: compiler.stage1, stage2: compiler.stage2, javascript: compiler.javascript,
    frontend: compiler.frontend, comparison: compiler.comparison, typed_metadata: compiler.typed_metadata,
    recovery: compiler.recovery,
    aff: { checks: aff.aff_checks, rust_tests: aff.rust_unit_tests, error_scenarios: aff.error_reporting_scenarios },
    native_tests: nativeTests.runs.map(r => ({ label: r.label, stdout: r.stdout, inputs: r.inputs })),
    pbo_tests: pboTests.runs.map(r => ({ label: r.label, stdout: r.stdout, test_sha256: r.test_sha256 })),
    regressions: { runs: regressions.runs, unavailable: regressions.unavailable },
    installed: verifyOnly ? null : { path: nativePath, sha256: compiler.stage2.sha256 } },
  compiler238, goRust, compiler244: secondary,
  json: { profile: jsonBuild.profile, fixture12: json12, gopurs238: json238, allocations238: jsonAllocations },
  resources, diagnostic, validation: { outputs_rechecked: outputs, generated_files_rechecked: generatedFiles, verified },
};
writeJson(destination, report);
console.log(`Verified ${outputs} outputs / ${generatedFiles} generated files; ${destination}`);
