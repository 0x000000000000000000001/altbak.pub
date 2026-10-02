import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { cpSync, existsSync, readFileSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { generate, hash, manifest, writeJson } from './common.mjs';

const workspace = resolve(process.argv[2]), resultPath = join(workspace, 'results.json');
assert(!existsSync(resultPath), `Existing campaign: ${resultPath}`);
const prepared = JSON.parse(readFileSync(join(workspace, 'prepared.json'), 'utf8'));
const ffi = JSON.parse(readFileSync(join(workspace, 'rust-ffi.json'), 'utf8'));
const referenceRuns = process.argv.slice(3, 5).map(path => JSON.parse(readFileSync(resolve(path), 'utf8')));
assert.deepEqual(referenceRuns.map(value => value.family).sort(), ['gopurs', 'purust']);
const qualifications = process.argv.slice(5).map(path => JSON.parse(readFileSync(resolve(path), 'utf8')));
assert.deepEqual(qualifications.map(value => value.generation.family).sort(), ['gopurs', 'purust']);
assert(qualifications.every(value => value.status === 'passed'));
const originalExecutions = JSON.parse(readFileSync(join(workspace, 'application-diagnostic/results.json'), 'utf8'));
const portableExecutions = JSON.parse(readFileSync(join(workspace, 'portable-validation/application-checks/results.json'), 'utf8'));
assert.equal(portableExecutions.length, 10);
assert(portableExecutions.every(run => run.passed));
const harness = join(workspace, 'harness-final');
cpSync(dirname(fileURLToPath(import.meta.url)), harness, { recursive: true });
const versions = (command, args) => execFileSync(command, args, { encoding: 'utf8' }).trim();
// Include only immutable inputs, not frontend caches or generated Go output.
const snapshot = () => Object.fromEntries([
  ['inputs', path => /\.(purs|rs|go|js)$/.test(path) || /\/corefn\.json$/.test(path) || /\/spago\.(yaml|lock)$/.test(path)],
  ['compilers', () => true], ['rust-ffi', () => true], ['rust-ffi-originals', () => true], ['harness-final', () => true],
].map(([directory, filter]) => [directory, manifest(join(workspace, directory), filter)]));
const frozen = snapshot();
for (const file of prepared.tast.files) assert.equal(hash(readFileSync(file.path)), file.sha256);
for (const compiler of Object.values(prepared.artifacts)) for (const file of compiler.files) assert.equal(hash(readFileSync(file.path)), file.sha256);
for (const file of ffi) assert.equal(hash(readFileSync(file.path)), file.sha256);
const variants = ['gopurs-js', 'gopurs-native', 'purust-js', 'purust-native'];
const result = {
  schema: 1, benchmark: 'gopurs-aff: gopurs and purust backend compilation', started_at: new Date().toISOString(), workspace,
  host: { model: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), architecture: process.arch,
    os: versions('sw_vers', []), node: process.version, go: versions('go', ['version']), rustc: versions('rustc', ['--version']), cargo: versions('cargo', ['--version']) },
  inputs: prepared, rust_ffi: ffi, frozen_files: frozen, frozen_sha256: hash(JSON.stringify(frozen)),
  protocol: { metric: '[gopurs|purust] backend total, monotonic rounded milliseconds', statistic: 'median',
    includes: ['TAST loading and sorting', 'backend-specific preparation', 'PBO optimization', 'target source generation and emission'],
    excludes: ['process startup and exit', 'purs frontend', 'compiler bootstrap', 'go build', 'Cargo', 'application execution'],
    warmups_per_variant: 1, measured_rounds: 5, variants,
    order: 'sequential runs; rotate first variant by one each measured round, starting with gopurs-js',
    cache: 'fresh generated output and no .purmeta/.cache per run; OS cache not flushed',
    settings: 'frozen production launchers, native defaults, explicit GOPURS_JS/PURUST_JS selectors; other compiler/GC overrides removed',
    rust_flags: ['--main', 'Test.Main', '--threaded'], go_flags: ['--main', 'Test.Main'],
    gopurs_defaults: { GOGC: 'off', GOMEMLIMIT: '10GiB', GOPURS_PBO_JOBS: 8, GOPURS_PREPARE_JOBS: 8, GOPURS_EMIT_JOBS: 8 },
    purust_defaults: { native_pbo_jobs: 4, native_codegen_jobs: 4, js_pbo_jobs: 1, js_codegen_jobs: 1 },
    note: 'gopurs prepares/monomorphizes for Go; purust uses its default Rust pipeline. These are complete backends, not identical compiler passes.' },
  runs: [], reference_generations: referenceRuns,
  application_validation: { original_executions: originalExecutions, portable_executions: portableExecutions,
    original_preflight_executions: ['qualify-go-1', 'qualify-rust-1', 'qualify-rust-release', 'qualify-rust-o3']
      .filter(label => existsSync(join(workspace, 'logs', label + '-test.json'))).map(label =>
      JSON.parse(readFileSync(join(workspace, 'logs', label + '-test.json'), 'utf8'))),
    portable_qualifications: qualifications,
    portable_test_changes: JSON.parse(readFileSync(join(workspace, 'portable-validation/changes.json'), 'utf8')),
    original_suite_status: 'Rust is unstable; all original-suite execution attempts are retained separately from timings.' },
  validation: { status: 'pending' },
};
const save = () => writeJson(resultPath, result);
const references = new Map(referenceRuns.map(value => [value.family,
  JSON.parse(readFileSync(value.generated_manifest, 'utf8'))]));
const median = values => values.toSorted((a, b) => a - b)[Math.floor(values.length / 2)];
save();
try {
  for (let round = 0; round <= 5; round++) {
    const start = round === 0 ? 0 : (round - 1) % variants.length;
    const order = variants.slice(start).concat(variants.slice(0, start));
    for (const variant of order) {
      const label = 'final-' + (round === 0 ? 'warmup' : `run-${round}`) + '-' + variant;
      const record = generate(workspace, variant, label);
      Object.assign(record, { round, warmup: round === 0 });
      result.runs.push(record); save();
      const files = JSON.parse(readFileSync(record.generated_manifest, 'utf8'));
      assert.deepEqual(files, references.get(record.family), `Generated output mismatch: ${label}`);
      record.identical_generated_files = files.length; save();
    }
  }
  result.summary = Object.fromEntries(variants.map(variant => {
    const runs = result.runs.filter(run => run.variant === variant && !run.warmup);
    const samples = runs.map(run => run.phases_ms['backend total']);
    const phases = Object.keys(runs[0].phases_ms);
    return [variant, { samples_ms: samples, median_ms: median(samples), min_ms: Math.min(...samples), max_ms: Math.max(...samples),
      median_wall_ms: median(runs.map(run => run.wall_ms)),
      phases_median_ms: Object.fromEntries(phases.map(phase => [phase, median(runs.map(run => run.phases_ms[phase]))])) }];
  }));
  const summary = result.summary;
  result.ratios = { purust_native_over_gopurs_native: summary['purust-native'].median_ms / summary['gopurs-native'].median_ms,
    gopurs_native_over_js: summary['gopurs-native'].median_ms / summary['gopurs-js'].median_ms,
    purust_native_over_js: summary['purust-native'].median_ms / summary['purust-js'].median_ms };
  assert.deepEqual(snapshot(), frozen, 'Frozen inputs changed');
  result.validation = { status: 'passed-with-original-rust-suite-instability', frozen_inputs_unchanged: true, outputs: result.runs.length,
    generated_files_per_go_output: references.get('gopurs').length, generated_files_per_rust_output: references.get('purust').length,
    comparison: 'Every output equals the built/executed original-corpus reference for its target, byte for byte.',
    portable_suite: '5/5 passes per target; 45 printed Aff checks and the 1000-item AVar stress assertion per pass.' };
  result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ summary: result.summary, ratios: result.ratios, validation: result.validation }, null, 2));
} catch (error) {
  result.failure = { at: new Date().toISOString(), message: error.message, stack: error.stack }; save(); throw error;
}
