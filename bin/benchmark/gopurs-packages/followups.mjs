// Correct the library-only assert invocation and retain confirmations of failures.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { dirname, join, resolve, basename } from 'node:path';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), original = JSON.parse(readFileSync(join(archive, 'results.json')));
const out = join(archive, 'followups'); assert(!existsSync(out)); mkdirSync(out);
const env = { ...environment(), ...original.protocol.jobs };
const state = { status: 'pending', original_results_sha256: hash(readFileSync(join(archive, 'results.json'))), results: [] };
const save = () => writeJson(join(out, 'results.json'), state);
const copy = (source, destination) => { mkdirSync(dirname(destination), { recursive: true }); copyFileSync(source, destination, constants.COPYFILE_FICLONE); };
const clean = input => {
  for (const path of walk(join(input, 'output'))) if (path.endsWith('.go') || ['go.mod', 'go.sum'].includes(basename(path))) rmSync(path);
  for (const name of ['.cache', '.purmeta']) rmSync(join(input, name), { recursive: true, force: true });
};
for (const previous of original.results.filter(result => result.status === 'failed')) {
  const directory = join(out, previous.name), input = join(directory, 'input'); mkdirSync(input, { recursive: true });
  const oldInput = join(archive, 'packages', previous.name, 'input');
  const inputs = previous.inputs?.manifest ?? manifest(oldInput, path => path.includes('/sources/') || path.endsWith('/corefn.json'));
  for (const file of inputs) { assert.equal(hash(readFileSync(join(oldInput, file.path))), file.sha256); copy(join(oldInput, file.path), join(input, file.path)); }
  const library = previous.name === 'gopurs-assert';
  const result = { name: previous.name, status: 'pending', started_at: new Date().toISOString(), inputs,
    invocation: library ? [] : ['--main', 'Test.Main'], runs: [],
    reason: library ? 'The package bin/test invokes gopurs without --main; it is a library-only generation.' : 'Confirm the original failure with unchanged compilers and frozen inputs.' };
  state.results.push(result); save();
  const attempts = library ? ['go', 'rust', ...Array(5).fill('rust')] : ['go', 'rust'];
  for (const [index, host] of attempts.entries()) {
    clean(input); assert.deepEqual(manifest(input), inputs);
    const label = `${index}-${host}`;
    let record;
    try {
      record = run(directory, label, join(archive, 'compiler/bin/gopurs'), result.invocation, input,
        { ...env, GOPURS_RUST: host === 'rust' ? '1' : '0', GOPURS_JS: '0', ...(previous.name === 'gopurs-prelude' ? { RUST_BACKTRACE: '1' } : {}) });
    } catch (error) {
      record = JSON.parse(readFileSync(join(directory, 'logs', label + '.json')));
      record.failure = error.stack;
    }
    record.host = host; record.warmup = index < 2;
    const raw = readFileSync(record.stderr, 'utf8');
    record.phases_ms = Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, phase, ms]) => [phase, Number(ms)]));
    record.generated = manifest(join(input, 'output'), path => path.endsWith('.go') || basename(path) === 'go.mod');
    record.output = join(directory, 'generated', label);
    for (const file of record.generated) copy(join(input, 'output', file.path), join(record.output, file.path));
    if (record.exit_code === 0 && host === 'rust') {
      const oracle = result.runs[0].generated;
      record.different_files = [...new Set([...oracle, ...record.generated].map(file => file.path))].filter(path =>
        oracle.find(file => file.path === path)?.sha256 !== record.generated.find(file => file.path === path)?.sha256);
      record.identical_files = record.different_files.length === 0 ? record.generated.length : null;
    }
    result.runs.push(record); save();
    console.log(`${result.name} ${label}: exit ${record.exit_code}, ${record.phases_ms['backend total'] ?? 'no total'} ms, ${record.different_files?.length ?? 0} different files`);
  }
  clean(input); assert.deepEqual(manifest(input), inputs);
  if (library) {
    assert(result.runs.every(record => record.exit_code === 0 && (record.host !== 'rust' || record.different_files.length === 0)));
    const samples = result.runs.filter(record => !record.warmup).map(record => record.phases_ms['backend total']);
    assert.equal(samples.length, 5); assert(samples.every(ms => ms > 0));
    result.summary = { samples_ms: samples, median_ms: samples.toSorted((a, b) => a - b)[2],
      min_ms: Math.min(...samples), max_ms: Math.max(...samples), identical_files_per_generation: result.runs[0].generated.length };
    result.status = 'passed';
  } else {
    const rust = result.runs.find(record => record.host === 'rust');
    assert(rust.exit_code !== 0 || rust.different_files.length > 0, 'Original failure did not reproduce; investigate before publishing');
    result.status = rust.exit_code !== 0 ? 'backend_failed' : 'output_mismatch';
    result.different_files = rust.different_files ?? [];
    result.readme_cell = rust.exit_code !== 0 ? 'failed (division by zero)' : 'output mismatch';
    if (rust.exit_code !== 0) assert(readFileSync(rust.stderr, 'utf8').includes('attempt to divide by zero'));
  }
  result.finished_at = new Date().toISOString(); save();
}
assert.deepEqual(manifest(join(archive, 'compiler')), original.compiler);
state.status = 'completed'; state.finished_at = new Date().toISOString(); save();
