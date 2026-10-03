// Reuse the qualified original Aff corpus and its per-target output oracles.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, generate, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [snapshotArg, outArg, configArg] = process.argv.slice(2);
assert(snapshotArg && outArg && configArg, 'compare-go-rust.mjs ORIGINAL_AFF NEW_DIRECTORY RUST_VARIANTS.json');
const snapshot = resolve(snapshotArg), out = resolve(outArg), configPath = resolve(configArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const original = JSON.parse(readFileSync(join(snapshot, 'results.json'), 'utf8'));
const rounds = Number(process.env.PURUST_BENCH_RUNS ?? 5);
assert(Number.isInteger(rounds) && rounds > 0);
function copy(source, destination) { mkdirSync(dirname(destination), { recursive: true }); copyFileSync(source, destination); }
for (const name of ['inputs', 'rust-ffi']) for (const file of original.frozen_files[name]) {
  assert.equal(hash(readFileSync(join(snapshot, name, file.path))), file.sha256);
  copy(join(snapshot, name, file.path), join(out, name, file.path));
}
// Preserve the frozen installation layout: bin/, sibling tools/ and the ESM
// package.json. The JS compiler resolves its WASM FFI runner from that layout.
for (const file of original.frozen_files.compilers.filter(file => file.path.startsWith('gopurs/'))) {
  const source = join(snapshot, 'compilers', file.path), destination = join(out, 'compilers', file.path);
  assert.equal(hash(readFileSync(source)), file.sha256);
  copy(source, destination);
  if (file.path.startsWith('gopurs/bin/')) chmodSync(destination, 0o755);
}
const rustVariants = JSON.parse(readFileSync(configPath, 'utf8')).map(config => {
  assert(/^[a-zA-Z0-9_-]+$/.test(config.label));
  const source = resolve(dirname(configPath), config.binary);
  const sha256 = hash(readFileSync(source));
  if (config.sha256) assert.equal(sha256, config.sha256);
  let qualification;
  if (config.qualification) {
    const path = resolve(dirname(configPath), config.qualification);
    const data = JSON.parse(readFileSync(path, 'utf8'));
    assert.equal(data.status, 'passed');
    assert.equal(sha256, config.javascript ? data.javascript.sha256 : data.stage2.sha256);
    qualification = { path, sha256: hash(readFileSync(path)) };
  }
  const binary = join(out, 'compilers', config.label, config.javascript ? 'purust.mjs' : 'purust-native');
  copy(source, binary); chmodSync(binary, 0o755);
  return { ...config, family: 'purust', source, binary, sha256, qualification };
});
const variants = [{ label: 'go-js', family: 'gopurs', javascript: true },
  { label: 'go-native', family: 'gopurs', javascript: false }, ...rustVariants];
assert.equal(new Set(variants.map(v => v.label)).size, variants.length);
const references = Object.fromEntries(original.reference_generations.map(record => {
  const files = JSON.parse(readFileSync(record.generated_manifest, 'utf8'));
  for (const file of files) assert.equal(hash(readFileSync(join(record.output, file.path))), file.sha256);
  return [record.family, files];
}));
const harness = dirname(fileURLToPath(import.meta.url));
copy(join(harness, 'compare-go-rust.mjs'), join(out, 'harness/purust-pbo/compare-go-rust.mjs'));
copy(join(harness, '../gopurs-aff/common.mjs'), join(out, 'harness/gopurs-aff/common.mjs'));
const immutable = () => ({
  inputs: manifest(join(out, 'inputs'), path => /\.(purs|rs|go|js)$/.test(path) || /\/corefn\.json$/.test(path) || /\/spago\.(yaml|lock)$/.test(path)),
  ffi: manifest(join(out, 'rust-ffi')), compilers: manifest(join(out, 'compilers')), harness: manifest(join(out, 'harness')),
});
const frozen = immutable();
assert.deepEqual(frozen.inputs, original.frozen_files.inputs);
assert.deepEqual(frozen.ffi, original.frozen_files['rust-ffi']);
const result = { status: 'pending', started_at: new Date().toISOString(), snapshot,
  original_results_sha256: hash(readFileSync(join(snapshot, 'results.json'))),
  tast: original.inputs.tast, frozen, references, variants,
  host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  protocol: { metric: 'backend total', warmups: 1, rounds, order: 'sequential rotating first variant',
    settings: 'frozen production Go launcher; Purust native defaults; explicit per-variant overrides only',
    includes: 'TAST load/sort, preparation, optimization, code generation, emission and worker drain',
    excludes: 'frontend, compiler/application builds, application execution and process startup/exit',
    cache: 'new outputs and .purmeta/.cache cleared each process; OS file cache not purged',
    note: 'Each target is checked against its own byte-identical source/manifest oracle; backend passes differ.' },
  runs: [] };
const save = () => writeJson(join(out, 'results.json'), result);
save();
try {
  for (let round = 0; round <= rounds; round++) {
    const first = round === 0 ? 0 : (round - 1) % variants.length;
    for (const variant of variants.slice(first).concat(variants.slice(0, first))) {
      const label = `${round === 0 ? 'warmup' : 'run-' + round}-${variant.label}`;
      let record;
      if (variant.family === 'gopurs') {
        record = generate(out, variant.javascript ? 'gopurs-js' : 'gopurs-native', label);
      } else {
        const cwd = join(out, 'inputs/gopurs-aff'), output = join(out, 'generated', label);
        for (const name of ['.purmeta', '.cache']) rmSync(join(cwd, name), { recursive: true, force: true });
        const args = ['--main', 'Test.Main', '--threaded', '--source', 'output', '--out', output,
          '--ffi-dir', relative(cwd, join(out, 'rust-ffi'))];
        record = run(out, label, variant.javascript ? process.execPath : variant.binary,
          variant.javascript ? ['--expose-gc', '--stack-size=65536', '--max-old-space-size=16384', variant.binary, ...args] : args,
          cwd, { ...environment(), ...variant.env });
        const stderr = readFileSync(record.stderr, 'utf8');
        record.phases_ms = Object.fromEntries([...stderr.matchAll(/^\[purust\] (.+): (\d+) ms$/gm)].map(([, phase, ms]) => [phase, Number(ms)]));
        assert.equal([...stderr.matchAll(/^\[purust\] backend total:/gm)].length, 1);
        assert(readFileSync(record.stdout, 'utf8').includes('Successfully generated Rust code.'));
        record.output = output;
        record.generated_manifest = join(out, 'logs', label + '.files.json');
        writeJson(record.generated_manifest, manifest(output, path => /\.(rs|toml)$/.test(path)));
      }
      Object.assign(record, { variant: variant.label, family: variant.family, round, warmup: round === 0 });
      result.runs.push(record); save();
      assert(record.phases_ms['backend total'] > 0);
      assert.deepEqual(JSON.parse(readFileSync(record.generated_manifest, 'utf8')), references[variant.family], label);
      record.identical_files = references[variant.family].length; save();
      console.log(`${label}: ${record.phases_ms['backend total']} ms; exact ${record.identical_files}`);
    }
  }
  const median = values => { const xs = values.toSorted((a, b) => a - b), n = xs.length;
    return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2; };
  result.summary = Object.fromEntries(variants.map(v => {
    const runs = result.runs.filter(r => r.variant === v.label && !r.warmup), times = runs.map(r => r.phases_ms['backend total']);
    return [v.label, { samples_ms: times, median_ms: median(times), mean_ms: times.reduce((s, x) => s + x, 0) / times.length,
      phases_median_ms: Object.fromEntries(Object.keys(runs[0].phases_ms).map(key => [key, median(runs.map(r => r.phases_ms[key]))])) }];
  }));
  assert.deepEqual(immutable(), frozen);
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify(result.summary, null, 2));
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
