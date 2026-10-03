// Serialized, rotating gopurs host/candidate comparison on original frozen Aff.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, existsSync, mkdirSync, readFileSync, renameSync, rmSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [snapshotArg, outArg, selectionArg] = process.argv.slice(2);
assert(snapshotArg && outArg && selectionArg, 'compare.mjs SNAPSHOT NEW_DIRECTORY VARIANTS.json');
const snapshot = resolve(snapshotArg), out = resolve(outArg), selectionFile = resolve(selectionArg);
const selection = JSON.parse(readFileSync(selectionFile, 'utf8'));
const rounds = selection.rounds ?? 5;
assert(Number.isInteger(rounds) && rounds > 0 && selection.variants.length >= 1);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const original = JSON.parse(readFileSync(join(snapshot, 'results.json'), 'utf8'));
function copy(from, to) { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to); }
for (const file of original.frozen_files.inputs) {
  const path = join(snapshot, 'inputs', file.path);
  assert.equal(hash(readFileSync(path)), file.sha256);
  copy(path, join(out, 'inputs', file.path));
}
const variants = selection.variants.map(variant => {
  assert(/^[a-z0-9-]+$/.test(variant.name));
  const destination = join(out, 'compilers', variant.name);
  let command;
  if (variant.directory) {
    const source = resolve(dirname(selectionFile), variant.directory);
    // Freeze the installed launcher and its full runtime support, not build caches.
    for (const name of ['bin/gopurs', 'bin/gopurs.js', 'bin/gopurs-native', 'bin/gopurs-rust',
      'package.json', 'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js'])
      copy(join(source, name), join(destination, name));
    command = join(destination, 'bin/gopurs');
  } else {
    command = join(destination, 'gopurs-rust');
    copy(resolve(dirname(selectionFile), variant.binary), command);
  }
  chmodSync(command, 0o755);
  return { ...variant, command };
});
assert.equal(new Set(variants.map(v => v.name)).size, variants.length);
copy(selectionFile, join(out, 'selection.json'));
copy(fileURLToPath(import.meta.url), join(out, 'compare.mjs'));
copy(fileURLToPath(new URL('../gopurs-aff/common.mjs', import.meta.url)), join(out, 'common.mjs'));
const inputManifest = () => manifest(join(out, 'inputs'), path =>
  /\.(purs|rs|go|js)$/.test(path) || /\/corefn\.json$/.test(path) || /\/spago\.(yaml|lock)$/.test(path));
assert.deepEqual(inputManifest(), original.frozen_files.inputs);
const oracle = original.reference_generations.find(record => record.family === 'gopurs');
const expected = JSON.parse(readFileSync(oracle.generated_manifest, 'utf8'));
for (const file of expected) assert.equal(hash(readFileSync(join(oracle.output, file.path))), file.sha256);
const result = { status: 'pending', started_at: new Date().toISOString(), snapshot, tast: original.inputs.tast,
  selection, variants, host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  frozen: { inputs: inputManifest(), compilers: manifest(join(out, 'compilers')) },
  protocol: { metric: '[gopurs] backend total', target: 'Go for every variant', warmups: 1, rounds,
    order: 'serialized processes, rotating first variant', jobs: { load: 8, prepare: 8, pbo: 8, emit: 8, pipeline: true },
    includes: 'load/sort, prepare, PBO, Go generation/emission, drain and entrypoints',
    excludes: 'frontend, compiler and application builds, application execution, process startup/exit',
    cache: 'fresh outputs and .purmeta/.cache per process, warm OS file cache',
    comparison: 'every Go file and go.mod byte-identical to original frozen oracle',
    resource_diagnostic: selection.resources ?? false }, runs: [] };
const save = () => writeJson(join(out, 'results.json'), result);
save();
try {
  for (let round = 0; round <= rounds; round++) {
    const start = round ? (round - 1) % variants.length : 0;
    for (const variant of variants.slice(start).concat(variants.slice(0, start))) {
      const label = `${round ? 'run-' + round : 'warmup'}-${variant.name}`;
      const cwd = join(out, 'inputs/gopurs-aff'), output = join(cwd, 'output'), generated = join(out, 'generated', label);
      const paths = ['purescript', 'gopurs_runtime', 'main', 'Test.Main/main', 'go.mod'];
      for (const path of [...paths, 'go.sum']) rmSync(join(output, path), { recursive: true, force: true });
      for (const path of ['.purmeta', '.cache']) rmSync(join(cwd, path), { recursive: true, force: true });
      const env = { ...environment(), GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8',
        GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1', ...variant.environment };
      const command = selection.resources ? '/usr/bin/time' : variant.command;
      const args = selection.resources ? ['-l', variant.command, '--main', 'Test.Main'] : ['--main', 'Test.Main'];
      const record = run(out, label, command, args, cwd, env);
      mkdirSync(generated, { recursive: true });
      for (const path of paths) {
        assert(existsSync(join(output, path)), path);
        mkdirSync(dirname(join(generated, path)), { recursive: true });
        renameSync(join(output, path), join(generated, path));
      }
      const stderr = readFileSync(record.stderr, 'utf8');
      record.phases_ms = Object.fromEntries([...stderr.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)]
        .map(([, phase, ms]) => [phase, Number(ms)]));
      assert.equal([...stderr.matchAll(/^\[gopurs\] backend total:/gm)].length, 1);
      assert(record.phases_ms['backend total'] > 0);
      const files = manifest(generated, path => path.endsWith('.go') || path.endsWith('/go.mod'));
      record.generated_manifest = join(out, 'logs', label + '.files.json');
      writeJson(record.generated_manifest, files);
      Object.assign(record, { variant: variant.name, round, warmup: round === 0, output: generated });
      result.runs.push(record); save();
      assert.deepEqual(files, expected, `${label}: changed generated Go`);
      record.identical_files = files.length; save();
      console.log(`${label}: ${record.phases_ms['backend total']} ms; ${files.length} exact files`);
    }
  }
  const median = values => { const xs = values.toSorted((a, b) => a - b), n = xs.length;
    return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2; };
  result.summary = Object.fromEntries(variants.map(({ name }) => {
    const runs = result.runs.filter(r => r.variant === name && !r.warmup);
    const samples = runs.map(r => r.phases_ms['backend total']);
    return [name, { samples_ms: samples, median_ms: median(samples),
      mean_ms: samples.reduce((a, b) => a + b, 0) / samples.length, min_ms: Math.min(...samples), max_ms: Math.max(...samples),
      phases_median_ms: Object.fromEntries(Object.keys(runs[0].phases_ms).map(phase =>
        [phase, median(runs.map(r => r.phases_ms[phase]))])) }];
  }));
  assert.deepEqual(inputManifest(), result.frozen.inputs);
  assert.deepEqual(manifest(join(out, 'compilers')), result.frozen.compilers);
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify(result.summary, null, 2));
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
