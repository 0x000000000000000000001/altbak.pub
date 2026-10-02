// Compare prebuilt Purust compilers on the original frozen gopurs-aff TAST.
// The source snapshot is read-only; caches and outputs live in a fresh copy.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [snapshotArg, resultArg, ...variantArgs] = process.argv.slice(2);
assert(snapshotArg && resultArg && variantArgs.length >= 2,
  'Usage: compare-aff.mjs FROZEN_GOPURS_AFF RESULT.json LABEL=BINARY_OR_CONFIG [...]\n' +
  'Config: { "binary": "relative/to/config", "env": { "PURUST_PBO_JOBS": "8" } }');
const snapshot = resolve(snapshotArg), resultPath = resolve(resultArg);
const rounds = Number(process.env.PURUST_BENCH_RUNS ?? 5);
assert(Number.isInteger(rounds) && rounds > 0, 'PURUST_BENCH_RUNS must be a positive integer');
assert(resultPath.endsWith('.json'));
const workspace = resultPath.slice(0, -5) + '-runs';
assert(!existsSync(resultPath), resultPath);
assert(!existsSync(workspace), workspace);
const priorPath = join(snapshot, 'results.json');
const prior = JSON.parse(readFileSync(priorPath, 'utf8'));
assert(prior.validation.frozen_inputs_unchanged);
const prepared = JSON.parse(readFileSync(join(snapshot, 'prepared.json'), 'utf8'));
assert.deepEqual(prepared.tast, prior.inputs.tast);
const referenceRun = prior.reference_generations.find(value => value.family === 'purust');
assert(referenceRun);
const reference = JSON.parse(readFileSync(referenceRun.generated_manifest, 'utf8'));
const referenceDirectory = referenceRun.output;
for (const file of reference) {
  assert.equal(hash(readFileSync(join(referenceDirectory, file.path))), file.sha256, file.path);
}
const variants = variantArgs.map(value => {
  const separator = value.indexOf('=');
  assert(separator > 0);
  const label = value.slice(0, separator), path = resolve(value.slice(separator + 1));
  assert(/^[A-Za-z0-9_-]+$/.test(label));
  const config = path.endsWith('.json') ? JSON.parse(readFileSync(path, 'utf8')) : { binary: path };
  const binary = resolve(dirname(path), config.binary), env = config.env ?? {};
  assert(Object.values(env).every(value => typeof value === 'string'));
  const sha256 = hash(readFileSync(binary));
  if (config.sha256) assert.equal(sha256, config.sha256, label);
  let qualification;
  if (label === 'after') {
    const qualificationPath = config.qualification ? resolve(dirname(path), config.qualification) : join(dirname(binary), 'qualification.json');
    if (existsSync(qualificationPath)) {
      const data = JSON.parse(readFileSync(qualificationPath, 'utf8'));
      assert.equal(data.status, 'passed');
      assert.equal(sha256, data.stage2.sha256, 'Candidate differs from the qualified stage 2');
      qualification = { path: qualificationPath, sha256: hash(readFileSync(qualificationPath)), stage2_sha256: data.stage2.sha256 };
    } else assert(!config.qualification, qualificationPath);
  }
  return { label, origin: binary, javascript: binary.endsWith('.js') || binary.endsWith('.mjs'), sha256, env, qualification };
});
assert.equal(new Set(variants.map(value => value.label)).size, variants.length);
assert(variants.some(value => value.label === 'before') && variants.some(value => value.label === 'after'),
  'The before and after variants are required');
for (const [label, name] of [['before', 'purust-native'], ['js', 'purust.js']]) {
  const variant = variants.find(value => value.label === label);
  if (variant) assert.equal(variant.sha256, prior.inputs.artifacts.purust.files.find(value => value.name === name).sha256,
    `${label} differs from the previously qualified compiler`);
}
mkdirSync(workspace, { recursive: true });
mkdirSync(join(workspace, 'generated'));

// Copy only the files covered by the earlier integrity record, avoiding its
// frontend caches, generated target code and Cargo/Go build products.
for (const directory of ['inputs', 'rust-ffi']) {
  const files = prior.frozen_files[directory];
  assert(files.length > 0);
  for (const file of files) {
    assert(!file.path.split('/').includes('..') && !file.path.startsWith('/'));
    const source = join(snapshot, directory, file.path), destination = join(workspace, directory, file.path);
    assert.equal(hash(readFileSync(source)), file.sha256, source);
    mkdirSync(dirname(destination), { recursive: true });
    copyFileSync(source, destination);
  }
}
for (const variant of variants) {
  variant.binary = join(workspace, 'compilers', variant.label, variant.javascript ? 'purust.mjs' : 'purust-native');
  mkdirSync(dirname(variant.binary), { recursive: true });
  copyFileSync(variant.origin, variant.binary);
  chmodSync(variant.binary, 0o755);
}
for (const name of ['purust-tast/compare-aff.mjs', 'gopurs-aff/common.mjs']) {
  const destination = join(workspace, 'harness', name);
  mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(resolve(dirname(fileURLToPath(import.meta.url)), '..', name), destination);
}
const inputManifest = () => manifest(join(workspace, 'inputs'), path =>
  /\.(purs|rs|go|js)$/.test(path) || /\/corefn\.json$/.test(path) || /\/spago\.(yaml|lock)$/.test(path));
const frozenInputs = inputManifest(), frozenFfi = manifest(join(workspace, 'rust-ffi'));
assert.deepEqual(frozenInputs, prior.frozen_files.inputs);
assert.deepEqual(frozenFfi, prior.frozen_files['rust-ffi']);
function verify() {
  assert.deepEqual(inputManifest(), frozenInputs, 'Frozen input files changed');
  assert.deepEqual(manifest(join(workspace, 'rust-ffi')), frozenFfi, 'Frozen FFI files changed');
  for (const variant of variants) assert.equal(hash(readFileSync(variant.binary)), variant.sha256, variant.label);
}
verify();
const cwd = join(workspace, 'inputs/gopurs-aff');
const result = {
  schema: 1, benchmark: 'Purust TAST optimization: complete frozen gopurs-aff backend',
  started_at: new Date().toISOString(), snapshot, workspace,
  provenance: { prior_result_sha256: hash(readFileSync(priorPath)),
    tast: { modules: prepared.tast.modules, types: prepared.tast.types, bytes: prepared.tast.bytes, sha256: prepared.tast.sha256 },
    frozen_inputs: frozenInputs, frozen_ffi: frozenFfi, harness: manifest(join(workspace, 'harness')),
    reference_output: referenceDirectory, reference_files: reference },
  host: { model: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(),
    architecture: process.arch, node: process.version },
  variants,
  protocol: { metric: 'Purust backend total, monotonic rounded milliseconds', statistic: 'median',
    warmups_per_variant: 1, measured_rounds: rounds, order: 'sequential; rotate first variant each measured round',
    includes: ['TAST loading and sorting', 'preparation', 'PBO', 'generation and emission', 'worker drain'],
    excludes: ['process startup and exit', 'purs frontend', 'compiler bootstrap', 'Cargo', 'application execution'],
    cache: 'fresh generated output and no .purmeta/.cache per run; OS cache not flushed',
    settings: 'prebuilt binaries; compiler/GC overrides removed, then explicit per-variant env applied',
    flags: ['--main', 'Test.Main', '--threaded', '--source', 'output', '--ffi-dir', 'relative frozen FFI directory'],
    phase_note: 'Medians of individual phases are not additive.' },
  runs: [], validation: { status: 'pending' },
};
const save = () => writeJson(resultPath, result);
const median = values => {
  const sorted = values.toSorted((a, b) => a - b), mid = Math.floor(sorted.length / 2);
  return sorted.length % 2 ? sorted[mid] : (sorted[mid - 1] + sorted[mid]) / 2;
};
const mean = values => values.reduce((sum, value) => sum + value, 0) / values.length;
save();
try {
  for (let round = 0; round <= rounds; round++) {
    const first = round === 0 ? 0 : (round - 1) % variants.length;
    for (const variant of variants.slice(first).concat(variants.slice(0, first))) {
      const label = `${round === 0 ? 'warmup' : `run-${round}`}-${variant.label}`;
      const output = join(workspace, 'generated', label);
      for (const directory of ['.purmeta', '.cache']) rmSync(join(cwd, directory), { recursive: true, force: true });
      const args = ['--main', 'Test.Main', '--threaded', '--source', 'output', '--out', output,
        '--ffi-dir', relative(cwd, join(workspace, 'rust-ffi'))];
      const command = variant.javascript ? process.execPath : variant.binary;
      const commandArgs = variant.javascript
        ? ['--expose-gc', '--stack-size=65536', '--max-old-space-size=16384', variant.binary, ...args] : args;
      const record = run(workspace, label, command, commandArgs, cwd, { ...environment(), ...variant.env });
      Object.assign(record, { variant: variant.label, round, warmup: round === 0, output });
      result.runs.push(record); save();
      const stderr = readFileSync(record.stderr, 'utf8');
      record.phases_ms = Object.fromEntries([...stderr.matchAll(/^\[purust\] (.+): (\d+) ms$/gm)].map(([, name, ms]) => [name, Number(ms)]));
      assert.equal([...stderr.matchAll(/^\[purust\] backend total:/gm)].length, 1);
      assert(record.phases_ms['backend total'] > 0);
      assert(readFileSync(record.stdout, 'utf8').includes('Successfully generated Rust code.'));
      const files = manifest(output, path => /\.(rs|toml)$/.test(path));
      record.generated_manifest = join(workspace, 'logs', label + '.files.json');
      writeJson(record.generated_manifest, files);
      assert.deepEqual(files, reference, `Generated output mismatch: ${label}`);
      record.identical_generated_files = files.length;
      writeJson(join(workspace, 'logs', label + '.json'), record);
      save();
      console.log(`${label}: ${record.phases_ms['backend total']} ms; ${files.length} identical files`);
    }
  }
  result.summary = Object.fromEntries(variants.map(({ label }) => {
    const runs = result.runs.filter(run => run.variant === label && !run.warmup);
    const samples = runs.map(run => run.phases_ms['backend total']);
    return [label, { samples_ms: samples, median_ms: median(samples), mean_ms: mean(samples),
      min_ms: Math.min(...samples), max_ms: Math.max(...samples),
      median_wall_ms: median(runs.map(run => run.wall_ms)),
      phases_median_ms: Object.fromEntries(Object.keys(runs[0].phases_ms).map(phase => [phase, median(runs.map(run => run.phases_ms[phase]))])) }];
  }));
  const paired = Array.from({ length: rounds }, (_, index) => {
    const time = label => result.runs.find(run => run.round === index + 1 && run.variant === label).phases_ms['backend total'];
    return time('after') - time('before');
  });
  result.paired_after_minus_before_ms = { samples: paired, median: median(paired), mean: mean(paired),
    after_faster_pairs: paired.filter(value => value < 0).length, pairs: rounds };
  verify();
  result.validation = { status: 'passed', frozen_inputs_unchanged: true, binaries_unchanged: true,
    outputs: result.runs.length, files_per_output: reference.length,
    comparison: 'Every Rust source and Cargo manifest equals the previously qualified original-corpus reference byte for byte.' };
  result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ summary: result.summary, validation: result.validation }, null, 2));
} catch (error) {
  result.failure = { at: new Date().toISOString(), message: error.message, stack: error.stack }; save(); throw error;
}
