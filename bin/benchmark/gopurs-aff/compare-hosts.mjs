// Compare one Go code generator hosted in JS, Go and Rust on frozen Aff TAST.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, existsSync, mkdirSync, readFileSync, renameSync, rmSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from './common.mjs';

const [snapshotArg, directoryArg, compilerArg] = process.argv.slice(2);
assert(snapshotArg && directoryArg && compilerArg, 'compare-hosts.mjs ORIGINAL_AFF NEW_DIRECTORY GOPURS_CHECKOUT');
const snapshot = resolve(snapshotArg), directory = resolve(directoryArg), compiler = resolve(compilerArg);
assert(!existsSync(directory), 'Use a new results directory');
mkdirSync(directory, { recursive: true });
const original = JSON.parse(readFileSync(join(snapshot, 'results.json'), 'utf8'));
const rounds = Number(process.env.GOPURS_BENCH_RUNS ?? 10);
assert(Number.isInteger(rounds) && rounds > 0);
function copy(source, destination) {
  mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(source, destination);
}
for (const file of original.frozen_files.inputs) {
  const source = join(snapshot, 'inputs', file.path);
  assert.equal(hash(readFileSync(source)), file.sha256);
  copy(source, join(directory, 'inputs', file.path));
}
for (const path of ['bin/gopurs', 'bin/gopurs.js', 'bin/gopurs-native', 'bin/gopurs-rust', 'package.json',
  'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js']) {
  copy(join(compiler, path), join(directory, 'compiler', path));
  if (path.startsWith('bin/')) chmodSync(join(directory, 'compiler', path), 0o755);
}
const harness = dirname(fileURLToPath(import.meta.url));
for (const path of ['compare-hosts.mjs', 'common.mjs']) copy(join(harness, path), join(directory, 'harness', path));
const inputManifest = () => manifest(join(directory, 'inputs'), path =>
  /\.(purs|rs|go|js)$/.test(path) || /\/corefn\.json$/.test(path) || /\/spago\.(yaml|lock)$/.test(path));
const frozen = { inputs: inputManifest(), compiler: manifest(join(directory, 'compiler')), harness: manifest(join(directory, 'harness')) };
assert.deepEqual(frozen.inputs, original.frozen_files.inputs);
const referenceRecord = original.reference_generations.find(record => record.family === 'gopurs');
const reference = JSON.parse(readFileSync(referenceRecord.generated_manifest, 'utf8'));
for (const file of reference) assert.equal(hash(readFileSync(join(referenceRecord.output, file.path))), file.sha256);
const hosts = ['js', 'go', 'rust'];
const result = {
  status: 'pending', started_at: new Date().toISOString(), snapshot, tast: original.inputs.tast,
  original_results_sha256: hash(readFileSync(join(snapshot, 'results.json'))), frozen,
  sources: {
    gopurs: manifest(join(compiler, 'src')),
    optimizer: manifest(resolve(compiler, '../../purescript-backend-optimizer-gopurs/src')),
  },
  host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  protocol: {
    metric: '[gopurs] backend total', target: 'Go for all three hosts', warmups: 1, rounds,
    order: 'serial processes, rotating first host',
    includes: 'TAST loading/sort, preparation, runtime, PBO, Go generation/emission, worker drain, entrypoints',
    excludes: 'frontend, bootstrap, application build/execution, process startup/exit',
    jobs: { load: 8, prepare: 8, pbo: 8, emit: 8, pipeline: true },
    cache: 'new generated output and .purmeta/.cache per process; warm OS file cache',
    comparison: 'each generated Go file and go.mod must match the original frozen gopurs oracle byte-for-byte',
  }, runs: [],
};
const save = () => writeJson(join(directory, 'results.json'), result);
save();
try {
  for (let round = 0; round <= rounds; round++) {
    const start = round === 0 ? 0 : (round - 1) % hosts.length;
    for (const host of hosts.slice(start).concat(hosts.slice(0, start))) {
      const label = `${round ? 'run-' + round : 'warmup'}-${host}`;
      const cwd = join(directory, 'inputs/gopurs-aff'), output = join(cwd, 'output');
      const generated = join(directory, 'generated', label);
      const paths = ['purescript', 'gopurs_runtime', 'main', 'Test.Main/main', 'go.mod'];
      for (const path of [...paths, 'go.sum']) rmSync(join(output, path), { recursive: true, force: true });
      for (const path of ['.purmeta', '.cache']) rmSync(join(cwd, path), { recursive: true, force: true });
      const env = { ...environment(), GOPURS_JS: host === 'js' ? '1' : '0', GOPURS_RUST: host === 'rust' ? '1' : '0',
        GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' };
      const record = run(directory, label, join(directory, 'compiler/bin/gopurs'), ['--main', 'Test.Main'], cwd, env);
      mkdirSync(generated, { recursive: true });
      for (const path of paths) {
        assert(existsSync(join(output, path)), `Missing ${path}`);
        mkdirSync(dirname(join(generated, path)), { recursive: true });
        renameSync(join(output, path), join(generated, path));
      }
      const stderr = readFileSync(record.stderr, 'utf8');
      record.phases_ms = Object.fromEntries([...stderr.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, name, time]) => [name, Number(time)]));
      assert.equal([...stderr.matchAll(/^\[gopurs\] backend total:/gm)].length, 1);
      assert(record.phases_ms['backend total'] > 0);
      const files = manifest(generated, path => path.endsWith('.go') || path.endsWith('/go.mod'));
      record.generated_manifest = join(directory, 'logs', label + '.files.json');
      writeJson(record.generated_manifest, files);
      Object.assign(record, { host, round, warmup: round === 0, output: generated });
      result.runs.push(record); save();
      assert.deepEqual(files, reference, `${label}: changed generated Go`);
      record.identical_files = files.length; save();
      console.log(`${label}: ${record.phases_ms['backend total']} ms, ${files.length} exact Go files/manifests`);
    }
  }
  const median = values => { const xs = values.toSorted((a, b) => a - b), n = xs.length;
    return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2; };
  result.summary = Object.fromEntries(hosts.map(host => {
    const runs = result.runs.filter(record => record.host === host && !record.warmup);
    const values = runs.map(record => record.phases_ms['backend total']);
    return [host, { samples_ms: values, median_ms: median(values), mean_ms: values.reduce((a, b) => a + b, 0) / values.length,
      min_ms: Math.min(...values), max_ms: Math.max(...values),
      phases_median_ms: Object.fromEntries(Object.keys(runs[0].phases_ms).map(phase => [phase, median(runs.map(r => r.phases_ms[phase]))])) }];
  }));
  assert.deepEqual(inputManifest(), frozen.inputs);
  assert.deepEqual(manifest(join(directory, 'compiler')), frozen.compiler);
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify(result.summary, null, 2));
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
