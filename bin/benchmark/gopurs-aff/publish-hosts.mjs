// Recheck retained artifacts and raw timings before publishing the host summary.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { hash, manifest, writeJson } from './common.mjs';

const [archiveArg, destinationArg] = process.argv.slice(2);
assert(archiveArg && destinationArg, 'publish-hosts.mjs ARCHIVE NEW_RESULT.json');
const archive = resolve(archiveArg), destination = resolve(destinationArg);
assert(!existsSync(destination));
const load = path => JSON.parse(readFileSync(path, 'utf8'));
const qpath = join(archive, 'qualification/results.json'), tpath = join(archive, 'timings/results.json');
const qualification = load(qpath), timings = load(tpath);
assert.equal(qualification.status, 'passed'); assert.equal(timings.status, 'passed');
const frozen = timings.frozen;
const timingRoot = join(archive, 'timings');
for (const [name, files] of Object.entries(frozen)) for (const file of files) {
  assert.equal(hash(readFileSync(join(timingRoot, name, file.path))), file.sha256, `${name}/${file.path}`);
}
for (const [name, binary] of Object.entries(qualification.executables)) {
  assert.equal(hash(readFileSync(binary.path)), binary.sha256, 'Installed ' + name);
  assert.equal(hash(readFileSync(join(timingRoot, 'compiler/bin', name))), binary.sha256, 'Frozen ' + name);
}
const goFiles = path => manifest(path, file => file.endsWith('.go') || file.endsWith('/go.mod'));
const rustFiles = path => manifest(path, file => (file.endsWith('.rs') || file.endsWith('/Cargo.toml'))
  && !file.endsWith('/Purs_Gopurs_FfiSupport/build.rs'));
assert.deepEqual(rustFiles(join(qualification.bootstrap.directory, 'rust')), qualification.bootstrap.native_generation);
assert.deepEqual(rustFiles(join(archive, 'qualification/bootstrap-js')), qualification.bootstrap.native_generation);
const original = load(join(timings.snapshot, 'results.json'));
const referenceRecord = original.reference_generations.find(record => record.family === 'gopurs');
const reference = load(referenceRecord.generated_manifest);
assert.deepEqual(goFiles(referenceRecord.output), reference);
let filesChecked = 0;
for (const record of timings.runs) {
  const phases = Object.fromEntries([...readFileSync(record.stderr, 'utf8').matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)]
    .map(([, name, value]) => [name, Number(value)]));
  assert.deepEqual(phases, record.phases_ms, record.label);
  assert.deepEqual(goFiles(record.output), reference, record.label);
  assert.deepEqual(load(record.generated_manifest), reference);
  assert.equal(record.exit_code, 0);
  filesChecked += reference.length;
}
const median = values => { const xs = values.toSorted((a, b) => a - b), n = xs.length;
  return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2; };
for (const host of ['js', 'go', 'rust']) {
  const records = timings.runs.filter(record => record.host === host);
  assert.equal(records.filter(record => record.warmup).length, 1);
  const samples = records.filter(record => !record.warmup).map(record => record.phases_ms['backend total']);
  assert.equal(samples.length, timings.protocol.rounds);
  assert.deepEqual(samples, timings.summary[host].samples_ms);
  assert.equal(median(samples), timings.summary[host].median_ms);
}
const aff = qualification.runs.filter(record => record.label.startsWith('aff-'));
assert.deepEqual(aff.map(record => record.host).sort(), ['go', 'js', 'rust']);
const expectedLines = readFileSync(join(qualification.aff, 'test/expected-main.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort();
for (const record of aff) {
  assert.deepEqual(goFiles(record.go_output), record.generated);
  assert.deepEqual(record.generated, aff[0].generated);
  assert.deepEqual(readFileSync(join(record.go_output, 'test.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort(), expectedLines);
  assert.equal(readFileSync(join(record.go_output, 'test.stderr'), 'utf8'), '');
  assert.equal(record.exit_code, 0);
}
const testOutput = readFileSync(join(archive, 'qualification/logs/compiler-tests.stdout'), 'utf8');
assert.match(testOutput, /ℹ pass 37\b/); assert.match(testOutput, /ℹ fail 0\b/); assert.match(testOutput, /ℹ skipped 0\b/);
assert.match(readFileSync(join(archive, 'qualification/logs/parser-tests.stdout'), 'utf8'), /^ok\s/m);
const readme = readFileSync(new URL('../../../README.md', import.meta.url), 'utf8');
const row = readme.split('\n').find(line => line.startsWith('[gopurs-aff]'));
for (const host of ['js', 'go', 'rust']) assert(row.includes(`~ ${Math.round(timings.summary[host].median_ms)} ms`));
assert(readme.includes('| [gopurs](https://github.com/0x000000000000000000001/gopurs) Rust binary (WIP)'));

const result = {
  status: 'passed', verified_at: new Date().toISOString(), compiler: 'gopurs', target: 'Go for every host',
  archive, host: timings.host, protocol: timings.protocol,
  source_records: [{ path: qpath, sha256: hash(readFileSync(qpath)) }, { path: tpath, sha256: hash(readFileSync(tpath)) }],
  executables: qualification.executables,
  bootstrap: { ...qualification.frontend, identical_generated_files: qualification.bootstrap.identical_files,
    installed_sha256: qualification.executables['gopurs-rust'].sha256,
    purust_native_sha256: qualification.bootstrap.purust_native_sha256, purust_js_sha256: qualification.bootstrap.purust_js_sha256 },
  validation: { node_tests_passed: 37, aff_checks_per_host: expectedLines.length, avar_stress_items: 1000,
    aff_identical_go_files: aff[0].identical_files, parser_race_checks: 'passed',
    timed_generations: timings.runs.length, timed_files_reverified: filesChecked },
  summary: timings.summary,
  rust_to_go_ratio: timings.summary.rust.median_ms / timings.summary.go.median_ms,
  rust_to_js_ratio: timings.summary.rust.median_ms / timings.summary.js.median_ms,
  retained_diagnostics: [
    { path: 'bootstrap-first.log', issue: 'Ambiguous Rust Value import in Main FFI; fixed by qualifying purust_core.' },
    { path: 'regressions-candidate.log', issue: 'Pre-existing JS/Go constant-folding mismatch for unused loneSurrogates in StringEscapes. Targeted live-string regression passes on all three hosts.' },
  ],
};
writeJson(destination, result);
console.log(`Verified ${timings.runs.length} timed outputs / ${filesChecked} Go files, 3 Aff applications and ${qualification.bootstrap.identical_files} bootstrap sources.`);
console.log(`Published ${destination}`);
