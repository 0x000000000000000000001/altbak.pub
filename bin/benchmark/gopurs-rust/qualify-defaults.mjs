// Qualify public launcher defaults separately from explicit-worker benchmarks.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { spawnSync } from 'node:child_process';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2] ?? '');
assert(process.argv[2], 'qualify-defaults.mjs ARCHIVE_WITH_FROZEN_BEFORE');
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../gopurs/gopurs');
const aff = resolve(root, '../gopurs-aff'), pbo = resolve(root, '../../purescript-backend-optimizer-gopurs');
const previous = resolve(archive, '../gopurs-rust-optimization-20261003');
const frontend = join(previous, 'frontend/purs'), production = join(archive, 'production');
assert(existsSync(join(archive, 'baseline.json')) && !existsSync(production));
mkdirSync(production);
const sources = () => ({ gopurs: manifest(join(root, 'src')), optimizer: manifest(join(pbo, 'src')) });
const state = sources();
const result = { status: 'pending', started_at: new Date().toISOString(), sources: state, commands: [] };
const save = () => writeJson(join(archive, 'qualification.json'), result);
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true',
  TMPDIR: production, GOPURS_NATIVE_TMPDIR: production, GOPURS_KEEP_WORKSPACE: '1', GOPURS_PURS: frontend, PURS: frontend };
function command(label, cmd, args, cwd, extra = {}) {
  console.log(label);
  const record = run(archive, label, cmd, args, cwd, { ...env, ...extra }, 3600000);
  result.commands.push(record); save(); return record;
}
copyFileSync(fileURLToPath(import.meta.url), join(archive, 'qualify-defaults.mjs'));
save();
try {
  // Demonstrate that the launcher contracts catch the original discrepancy.
  const regression = join(archive, 'regression-before');
  for (const path of ['bin', 'tools']) mkdirSync(join(regression, path), { recursive: true });
  copyFileSync(join(archive, 'before/bin/gopurs'), join(regression, 'bin/gopurs'));
  copyFileSync(join(root, 'tools/launcher.test.mjs'), join(regression, 'tools/launcher.test.mjs'));
  const failed = spawnSync(process.execPath, ['--test', join(regression, 'tools/launcher.test.mjs')], { env, encoding: 'utf8' });
  assert.ifError(failed.error);
  assert.equal(failed.status, 1, 'The old launcher must fail the new behavioral contracts');
  writeFileSync(join(regression, 'stdout'), failed.stdout); writeFileSync(join(regression, 'stderr'), failed.stderr);
  result.original_launcher_regression = { exit_code: failed.status, directory: regression }; save();
  command('launcher-tests', 'npm', ['run', 'test:launcher'], root);
  command('build-go-js', 'npm', ['run', 'build:native', '--', '--keep-workspace'], root);
  const rust = command('rust-default-rebuild-aff', './bin/test', ['-c'], aff, { GOPURS_RUST: '1' });
  result.rust_bootstrap = readFileSync(rust.stdout, 'utf8').match(/^Rust-hosted gopurs workspace: (.+)$/m)?.[1];
  assert(result.rust_bootstrap);
  const defaults = '[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true';
  assert(readFileSync(rust.stderr, 'utf8').includes(defaults)); save();
  command('qualify-hosts', process.execPath, [fileURLToPath(new URL('../gopurs-aff/qualify-hosts.mjs', import.meta.url)),
    root, aff, result.rust_bootstrap, join(archive, 'hosts')], root);
  const hosts = JSON.parse(readFileSync(join(archive, 'hosts/results.json'), 'utf8'));
  for (const record of hosts.runs.filter(r => r.host)) {
    assert(readFileSync(record.stderr, 'utf8').includes(defaults), record.host);
    const current = manifest(record.go_output, file => file.endsWith('.go') || file.endsWith('/go.mod'));
    const old = manifest(join(previous, 'production/hosts/aff-rust'), file => file.endsWith('.go') || file.endsWith('/go.mod'));
    assert.deepEqual(current, old, 'Live Aff output changed');
  }
  assert.deepEqual(sources(), state);
  result.executables = Object.fromEntries(['gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, { path: join(root, 'bin', name), sha256: hash(readFileSync(join(root, 'bin', name))) }]));
  save();
  const selection = { rounds: 5, launcherDefaults: true, variants: [
    { name: 'rust-before-default', directory: 'before', environment: { GOPURS_RUST: '1' } },
    { name: 'rust-default', directory: root, environment: { GOPURS_RUST: '1' } },
    { name: 'rust-explicit-eight', directory: root, environment: { GOPURS_RUST: '1', GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' } },
    { name: 'go-default', directory: root },
    { name: 'js-default', directory: root, environment: { GOPURS_JS: '1' } },
  ] };
  writeJson(join(archive, 'defaults.json'), selection);
  command('compare-defaults', process.execPath, [fileURLToPath(new URL('./compare.mjs', import.meta.url)),
    resolve(archive, '../gopurs-purust-aff-20261002'), join(archive, 'default-runs'), join(archive, 'defaults.json')], root);
  command('verify-defaults', process.execPath, [fileURLToPath(new URL('./verify.mjs', import.meta.url)),
    join(archive, 'verification.json'), join(archive, 'default-runs')], root);
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log('Default-host qualification and comparison passed');
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
