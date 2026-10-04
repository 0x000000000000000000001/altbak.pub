// Recheck corrected compilers against the original, frozen library campaign.
// Measure only previously unqualified rows; all remaining rows are regressions.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { cpus, totalmem } from 'node:os';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';

const [originalArg, archiveArg] = process.argv.slice(2);
assert(originalArg && archiveArg, 'recheck.mjs ORIGINAL_ARCHIVE FIX_ARCHIVE');
const original = resolve(originalArg), archive = resolve(archiveArg);
const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..');
const root = resolve(site, '../gopurs/gopurs');
const previous = JSON.parse(readFileSync(join(original, 'results.json')));
const followups = JSON.parse(readFileSync(join(original, 'followups/results.json')));
const build = JSON.parse(readFileSync(join(archive, 'build-results.json')));
assert.equal(build.status, 'passed'); assert.deepEqual(manifest(join(archive, 'bin')), build.binaries);
const correctionsPath = join(archive, 'oracle-corrections.json');
const corrections = existsSync(correctionsPath) ? JSON.parse(readFileSync(correctionsPath)) : [];
const destination = join(archive, 'recheck'); assert(!existsSync(destination)); mkdirSync(destination);
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
copy(fileURLToPath(import.meta.url), join(destination, 'recheck.mjs'));
const compiler = join(destination, 'compiler');
for (const file of ['bin/gopurs', 'bin/gopurs-native', 'bin/gopurs-rust', 'bin/gopurs.js',
  'package.json', 'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js']) copy(join(root, file), join(compiler, file));
for (const name of ['gopurs', 'gopurs-native', 'gopurs-rust', 'gopurs.js'])
  assert.equal(hash(readFileSync(join(compiler, 'bin', name))), build.binaries.find(file => file.path === name).sha256);
const jobs = previous.protocol.jobs;
const env = { ...environment(), ...jobs, GOWORK: 'off' };
const cases = previous.results.map(old => {
  const replacement = followups.results.find(result => result.name === old.name);
  const chosen = old.status === 'passed' ? old : replacement;
  assert(chosen);
  return { name: old.name, input_source: join(original, old.status === 'passed' ? 'packages' : 'followups', old.name, 'input'),
    inputs: chosen.inputs.manifest ?? chosen.inputs, invocation: chosen.invocation ?? ['--main', 'Test.Main'],
    measured: chosen.status !== 'passed', oracle: chosen.oracle ?? chosen.runs[0] };
});
const affArchive = resolve(original, '../gopurs-purust-aff-20261002');
const aff = JSON.parse(readFileSync(join(affArchive, 'results.json')));
const affOracle = aff.reference_generations.find(record => record.variant === 'gopurs-native');
cases.push({ name: 'gopurs-aff', input_source: join(affArchive, 'inputs/gopurs-aff'),
  inputs: aff.frozen_files.inputs.filter(file => file.path.startsWith('gopurs-aff/')).map(file => ({ ...file, path: file.path.slice('gopurs-aff/'.length) })),
  sibling_source: join(affArchive, 'inputs'),
  sibling_inputs: aff.frozen_files.inputs.filter(file => !file.path.startsWith('gopurs-aff/')),
  invocation: ['--main', 'Test.Main'], measured: false,
  oracle: { ...affOracle, generated: JSON.parse(readFileSync(affOracle.generated_manifest)) } });
for (const correction of corrections) {
  assert.equal(correction.package, 'gopurs-prelude');
  const item = cases.find(item => item.name === correction.package);
  assert.deepEqual(item.oracle.generated, correction.previous_oracle.generated);
  const changed = correction.oracle.generated.filter(file => item.oracle.generated.find(old => old.path === file.path)?.sha256 !== file.sha256);
  assert.deepEqual(changed.map(file => file.path), ['purescript/Test_Main.go']);
  const old = readFileSync(join(item.oracle.output, changed[0].path), 'utf8');
  const corrected = readFileSync(join(correction.oracle.output, changed[0].path), 'utf8');
  const literal = 'Call_Data_Ord_signum__2002100468(0.0)';
  assert.equal(old.split(literal).length, 2);
  assert.equal(corrected, old.replace(literal, 'Call_Data_Ord_signum__2002100468(gopurs_runtime.NegativeZero())'));
  item.previous_oracle = item.oracle; item.oracle = correction.oracle; item.oracle_correction = correction.reason;
}
assert.equal(cases.length, 50); assert.equal(cases.filter(item => item.measured).length, 4);
cases.sort((a, b) => Number(b.measured) - Number(a.measured) || a.name.localeCompare(b.name));
const state = { status: 'running', started_at: new Date().toISOString(),
  original, original_results_sha256: hash(readFileSync(join(original, 'results.json'))),
  compiler: manifest(compiler), host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  protocol: { ...previous.protocol, rounds: 5,
    scope: 'Four corrected rows measured; all 50 packages checked against their retained Go-output oracle.',
    oracle: 'Byte-exact retained Go-host output, with any independently diagnosed correction explicitly recorded; corrected Go and JS hosts additionally checked on the four corrected rows.',
    validation: 'Corrected rows also checked with sequential Rust workers, then their generated Go applications built and executed outside timing.' },
  oracle_corrections: corrections, results: [] };
const save = () => writeJson(join(destination, 'results.json'), state);
const clean = input => {
  for (const path of walk(join(input, 'output'))) if (path.endsWith('.go') || ['go.mod', 'go.sum'].includes(basename(path))) rmSync(path);
  for (const name of ['.cache', '.purmeta']) rmSync(join(input, name), { recursive: true, force: true });
};
save();
for (const item of cases) {
  const directory = join(destination, item.name), input = join(directory, 'input');
  const result = { ...item, status: 'running', runs: [], applications: [] };
  state.results.push(result); save();
  try {
    assert.deepEqual(manifest(item.oracle.output, path => path.endsWith('.go') || basename(path) === 'go.mod'), item.oracle.generated);
    // The historical Aff corpus keeps modulePath references to ../gopurs-*.
    // Keep its entire frozen sibling family when relocating the project.
    for (const file of item.sibling_inputs ?? []) {
      const source = join(item.sibling_source, file.path);
      assert.equal(hash(readFileSync(source)), file.sha256); copy(source, join(directory, file.path));
    }
    for (const file of item.inputs) {
      const source = join(item.input_source, file.path);
      assert.equal(hash(readFileSync(source)), file.sha256); copy(source, join(input, file.path));
    }
    assert.deepEqual(manifest(input), item.inputs);
    const attempts = item.measured ? [
      { label: 'go', host: 'go' }, { label: 'js', host: 'js' },
      { label: 'rust-sequential', host: 'rust', jobs: { ...jobs, GOPURS_JOBS: '1', GOPURS_PREPARE_JOBS: '1', GOPURS_PBO_JOBS: '1', GOPURS_EMIT_JOBS: '1' } },
      { label: 'rust-warmup', host: 'rust' },
      ...Array.from({ length: 5 }, (_, i) => ({ label: `rust-${i + 1}`, host: 'rust', measured: true })),
    ] : [{ label: 'rust-regression', host: 'rust' }];
    for (const attempt of attempts) {
      clean(input); assert.deepEqual(manifest(input), item.inputs);
      const record = run(directory, attempt.label, join(compiler, 'bin/gopurs'), item.invocation, input,
        { ...env, ...attempt.jobs, GOPURS_RUST: attempt.host === 'rust' ? '1' : '0', GOPURS_JS: attempt.host === 'js' ? '1' : '0' });
      const raw = readFileSync(record.stderr, 'utf8');
      const totals = [...raw.matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
      assert.equal(totals.length, 1); assert(Number(totals[0][1]) > 0);
      record.phases_ms = Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, name, ms]) => [name, Number(ms)]));
      record.host = attempt.host; record.measured = attempt.measured ?? false;
      record.generated = manifest(join(input, 'output'), path => path.endsWith('.go') || basename(path) === 'go.mod');
      record.output = join(directory, 'generated', attempt.label);
      for (const file of record.generated) copy(join(input, 'output', file.path), join(record.output, file.path));
      result.runs.push(record); save();
      const expected = new Map(item.oracle.generated.map(file => [file.path, file.sha256]));
      const actual = new Map(record.generated.map(file => [file.path, file.sha256]));
      record.different_files = [...new Set([...expected.keys(), ...actual.keys()])].filter(path => expected.get(path) !== actual.get(path));
      save(); assert.deepEqual(record.different_files, [], `${item.name}: ${attempt.label} differs from original Go oracle`);
      record.identical_files = record.generated.length; save();
      console.log(`${item.name} ${attempt.label}: ${record.phases_ms['backend total']} ms, ${record.identical_files} exact files`);
    }
    if (item.measured) {
      const samples = result.runs.filter(record => record.measured).map(record => record.phases_ms['backend total']);
      assert.equal(samples.length, 5);
      result.summary = { samples_ms: samples, median_ms: samples.toSorted((a, b) => a - b)[2],
        min_ms: Math.min(...samples), max_ms: Math.max(...samples), identical_files_per_generation: item.oracle.generated.length };
      save();
      const application = join(directory, 'application'); mkdirSync(application);
      for (const [label, command, args, cwd] of [
        ['go-build', 'go', ['build', '-o', join(application, 'tests'), './main'], join(input, 'output')],
        ['go-execute', join(application, 'tests'), [], input],
      ]) {
        result.applications.push(run(directory, label, command, args, cwd, { ...env, GOGC: '1000' })); save();
      }
    }
    clean(input); assert.deepEqual(manifest(input), item.inputs);
    for (const family of new Set((item.sibling_inputs ?? []).map(file => file.path.split('/')[0]))) {
      const prefix = family + '/';
      assert.deepEqual(manifest(join(directory, family)).map(file => ({ ...file, path: prefix + file.path })),
        item.sibling_inputs.filter(file => file.path.startsWith(prefix)));
    }
    result.status = 'passed';
  } catch (error) {
    result.status = 'failed'; result.error = error.stack; console.error(`${item.name}: ${error.message}`);
  }
  result.finished_at = new Date().toISOString(); save();
}
assert.deepEqual(manifest(compiler), state.compiler);
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'failed';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, packages: state.results.length,
  failed: state.results.filter(result => result.status !== 'passed').map(result => result.name),
  measured: state.results.filter(result => result.measured).map(({ name, summary, status }) => ({ name, summary, status })) }, null, 2));
if (state.status !== 'passed') process.exitCode = 1;
