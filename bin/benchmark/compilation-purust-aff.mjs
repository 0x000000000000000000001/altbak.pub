#!/usr/bin/env node
// Measure only the backend, then validate the generated Test.Main outside timing.
import assert from 'node:assert/strict';
import { execFileSync, spawnSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { chmodSync, copyFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync,
  readdirSync, rmSync, statSync, writeFileSync } from 'node:fs';
import { cpus, loadavg, tmpdir, totalmem } from 'node:os';
import { delimiter, dirname, isAbsolute, join, relative, resolve, sep } from 'node:path';
import { pathToFileURL } from 'node:url';

const [compilerArg, projectArg, resultArg] = process.argv.slice(2);
if (!compilerArg || !projectArg || !resultArg || process.argv.length !== 5) {
  throw new Error('Usage: node compilation-purust-aff.mjs COMPILER_DIR PURUST_AFF_DIR RESULT_JSON');
}
const compiler = resolve(compilerArg), project = resolve(projectArg), resultPath = resolve(resultArg);
assert(!existsSync(resultPath), `Result already exists: ${resultPath}`);
const workspace = mkdtempSync(join(process.env.TMPDIR ?? tmpdir(), 'purust-aff-compilation-'));
const snapshot = join(workspace, 'inputs'), cwd = join(snapshot, 'purust-aff');
const bin = join(workspace, 'compiler/bin');
mkdirSync(cwd, { recursive: true });
mkdirSync(bin, { recursive: true });
mkdirSync(join(workspace, 'logs'));
mkdirSync(join(workspace, 'generated'));
console.log(`Retained workspace: ${workspace}`);

const sha256 = bytes => createHash('sha256').update(bytes).digest('hex');
const fingerprint = path => ({ bytes: statSync(path).size, sha256: sha256(readFileSync(path)) });
const commandText = (command, args, directory = project) =>
  execFileSync(command, args, { cwd: directory, encoding: 'utf8' }).trim();
const revision = directory => ({ directory, commit: commandText('git', ['rev-parse', 'HEAD'], directory),
  status: commandText('git', ['status', '--porcelain=v1'], directory) });
const files = new Map();
function freeze(source, destination) {
  const key = relative(workspace, destination);
  assert(!key.startsWith('..' + sep) && !isAbsolute(key));
  if (files.has(key)) return;
  mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(source, destination);
  assert.equal(sha256(readFileSync(source)), sha256(readFileSync(destination)));
  files.set(key, { path: key, ...fingerprint(destination) });
}
const helpers = join(compiler, 'tools/native-workspace.mjs');
freeze(helpers, join(workspace, 'native-workspace.mjs'));
const { verifyTypedOutput, compareGeneratedSources } = await import(pathToFileURL(join(workspace, 'native-workspace.mjs')));

let tastBytes = 0;
const builtWith = new Set();
for (const entry of readdirSync(join(project, 'output'), { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name))) {
  const source = join(project, 'output', entry.name, 'corefn.json');
  if (!entry.isDirectory() || !existsSync(source)) continue;
  const data = JSON.parse(readFileSync(source, 'utf8'));
  assert(!isAbsolute(data.modulePath), `Expected relative modulePath: ${data.modulePath}`);
  freeze(source, join(cwd, 'output', entry.name, 'corefn.json'));
  tastBytes += statSync(source).size;
  builtWith.add(data.builtWith);
  for (const path of [data.modulePath, data.modulePath.replace(/\.purs$/, '.rs'),
    data.modulePath.replace(/\.purs$/, '.rs.cargo.json')]) {
    if (existsSync(resolve(project, path))) freeze(resolve(project, path), resolve(cwd, path));
  }
}
for (const path of ['spago.yaml', 'spago.lock', 'test/expected-main.stdout']) {
  freeze(join(project, path), join(cwd, path));
}
for (const name of ['purust', 'purust.js', 'purust-native']) {
  freeze(join(compiler, 'bin', name), join(bin, name));
}
chmodSync(join(bin, 'purust'), 0o755);
chmodSync(join(bin, 'purust-native'), 0o755);
const frozenFiles = [...files.values()].sort((a, b) => a.path < b.path ? -1 : a.path > b.path ? 1 : 0);
const aggregate = entries => sha256(entries.map(file => `${file.sha256}  ${file.path}\n`).join(''));
function verifyFrozen() {
  for (const file of frozenFiles) assert.deepEqual(fingerprint(join(workspace, file.path)),
    { bytes: file.bytes, sha256: file.sha256 }, `Frozen file changed: ${file.path}`);
}

const environment = { ...process.env, PATH: dirname(process.execPath) + delimiter + (process.env.PATH ?? '') };
const removedVariables = Object.keys(environment).filter(key => /^(PURUST_|GOPURS_|NODE_|RUST|CARGO)/.test(key));
for (const key of removedVariables) delete environment[key];
const flags = ['--source', 'output', '--main', 'Test.Main', '--threaded'];
const result = {
  schema: 1, benchmark: 'purust-aff backend compilation', started_at: new Date().toISOString(), workspace,
  host: { model: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(),
    os: commandText('sw_vers', []), architecture: process.arch,
    node: process.version, rustc: commandText('rustc', ['--version']), cargo: commandText('cargo', ['--version']) },
  repositories_at_snapshot: { compiler: revision(compiler), project: revision(project),
    optimizer: revision(resolve(compiler, '../../purescript-backend-optimizer-purust')) },
  compiler_artifacts: Object.fromEntries(['purust', 'purust.js', 'purust-native'].map(name =>
    [name, { ...fingerprint(join(bin, name)), origin: join(compiler, 'bin', name) }])),
  inputs: { ...verifyTypedOutput(join(cwd, 'output')), tast_bytes: tastBytes, built_with: [...builtWith],
    tast_sha256: aggregate(frozenFiles.filter(file => file.path.endsWith('/corefn.json'))),
    all_frozen_files_sha256: aggregate(frozenFiles), files: frozenFiles },
  protocol: { metric: '[purust] backend total (rounded monotonic milliseconds)',
    includes: ['TAST loading and sorting', 'preparation', 'PBO optimization', 'Rust generation and emission'],
    excludes: ['process startup and exit', 'purs frontend', 'compiler bootstrap', 'Cargo', 'application tests'],
    flags, warmups_per_backend: 1, measured_runs_per_backend: 5, order: 'JS, native, repeated',
    process: 'fresh process per run', cache: 'fresh output and no .purmeta before each run; OS file cache not flushed',
    statistic: 'median', cpu_affinity: 'unset', environment_variables_removed: removedVariables,
    js_selection: 'PURUST_JS=1', native_selection: 'PURUST_JS=0',
    node_flags: ['--expose-gc', '--stack-size=65536', '--max-old-space-size=16384'] },
  runs: [], validation: { status: 'pending' }
};
const save = () => {
  const text = JSON.stringify(result, null, 2) + '\n';
  writeFileSync(join(workspace, 'results.json'), text);
  mkdirSync(dirname(resultPath), { recursive: true });
  writeFileSync(resultPath, text);
};
save();

let reference;
try {
  for (let round = 0; round <= 5; round++) {
    for (const backend of ['js', 'native']) {
      const label = `${round === 0 ? 'warmup' : `run-${round}`}-${backend}`;
      const out = join(workspace, 'generated', label);
      assert(!existsSync(out));
      rmSync(join(cwd, '.purmeta'), { recursive: true, force: true });
      const argv = [...flags, '--out', out];
      const loadBefore = loadavg();
      const started = performance.now();
      const execution = spawnSync(join(bin, 'purust'), argv, { cwd,
        env: { ...environment, PURUST_JS: backend === 'js' ? '1' : '0' },
        encoding: 'utf8', timeout: 180000, maxBuffer: 32 * 1024 * 1024 });
      const wallMs = performance.now() - started;
      writeFileSync(join(workspace, 'logs', label + '.stdout'), execution.stdout ?? '');
      writeFileSync(join(workspace, 'logs', label + '.stderr'), execution.stderr ?? '');
      const phases = Object.fromEntries([...(execution.stderr ?? '').matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
        .map(([, name, ms]) => [name, Number(ms)]));
      const run = { label, backend, warmup: round === 0, round, command: [join(bin, 'purust'), ...argv],
        exit_code: execution.status, signal: execution.signal, wall_ms: wallMs, phases_ms: phases,
        load_average_before: loadBefore, stdout: execution.stdout, stderr: execution.stderr };
      result.runs.push(run);
      save();
      assert.ifError(execution.error);
      assert.equal(execution.status, 0, execution.stderr);
      assert(execution.stdout.includes('Successfully generated Rust code.'));
      assert(phases['backend total'] > 0);
      assert.equal([...execution.stderr.matchAll(/^\[purust\] backend total:/gm)].length, 1);
      if (!reference) reference = out;
      run.identical_generated_files = compareGeneratedSources(reference, out);
      save();
      console.log(`${label}: ${phases['backend total']} ms (${run.identical_generated_files} identical Rust/manifests)`);
    }
  }
  result.summary = Object.fromEntries(['js', 'native'].map(backend => {
    const samples = result.runs.filter(run => !run.warmup && run.backend === backend).map(run => run.phases_ms['backend total']);
    const sorted = samples.toSorted((a, b) => a - b);
    return [backend, { samples_ms: samples, median_ms: sorted[2], min_ms: sorted[0], max_ms: sorted[4] }];
  }));
  result.summary.native_over_js = result.summary.native.median_ms / result.summary.js.median_ms;
  result.measurement_finished_at = new Date().toISOString();
  verifyFrozen();
  save();

  // Cargo and the executable run only after all backend measurements finish.
  const generated = join(workspace, 'generated/run-5-native');
  const buildArgs = ['build', '--offline', '--manifest-path', join(generated, 'Cargo.toml')];
  const build = spawnSync('cargo', buildArgs, { cwd, env: environment, encoding: 'utf8',
    timeout: 300000, maxBuffer: 32 * 1024 * 1024 });
  writeFileSync(join(workspace, 'logs/cargo-build.stdout'), build.stdout ?? '');
  writeFileSync(join(workspace, 'logs/cargo-build.stderr'), build.stderr ?? '');
  result.validation = { status: 'building', command: ['cargo', ...buildArgs],
    cargo_exit_code: build.status, cargo_stdout: build.stdout, cargo_stderr: build.stderr };
  save();
  assert.ifError(build.error);
  assert.equal(build.status, 0, build.stderr);
  const application = join(generated, 'target/debug/purust_output');
  const execution = spawnSync(application, [], { cwd, env: environment, encoding: 'utf8', timeout: 120000 });
  writeFileSync(join(workspace, 'logs/aff.stdout'), execution.stdout ?? '');
  writeFileSync(join(workspace, 'logs/aff.stderr'), execution.stderr ?? '');
  Object.assign(result.validation, { application, application_sha256: fingerprint(application).sha256,
    exit_code: execution.status, stdout: execution.stdout, stderr: execution.stderr,
    expected_stdout_sha256: fingerprint(join(cwd, 'test/expected-main.stdout')).sha256 });
  save();
  assert.ifError(execution.error);
  assert.equal(execution.status, 0);
  assert.equal(execution.stderr, '');
  assert.equal(execution.stdout, readFileSync(join(cwd, 'test/expected-main.stdout'), 'utf8'));
  result.validation.status = 'passed';
  result.validation.aff_checks = execution.stdout.trim().split('\n').length;
  result.validation.identical_generated_files = compareGeneratedSources(reference, generated);
  verifyFrozen();
  result.validation.frozen_inputs_unchanged = true;
  result.finished_at = new Date().toISOString();
  save();
  console.log(JSON.stringify({ summary: result.summary, validation: result.validation.status,
    aff_checks: result.validation.aff_checks, result: resultPath }, null, 2));
} catch (error) {
  result.failure = { message: error.message, stack: error.stack, at: new Date().toISOString() };
  save();
  throw error;
}
