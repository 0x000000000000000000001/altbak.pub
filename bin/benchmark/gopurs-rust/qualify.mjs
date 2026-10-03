// Rebuild the selected sources through the production paths, then qualify hosts.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, candidateArg, resumeArg] = process.argv.slice(2);
assert(archiveArg && candidateArg && (!resumeArg || resumeArg === '--resume'),
  'qualify.mjs ARCHIVE CANDIDATE_NAME [--resume]');
const archive = resolve(archiveArg), candidate = join(archive, 'candidates', candidateArg);
const directory = join(archive, 'production');
if (!resumeArg) assert(!existsSync(directory));
mkdirSync(directory, { recursive: true });
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../gopurs/gopurs');
const pbo = resolve(root, '../../purescript-backend-optimizer-gopurs'), aff = resolve(root, '../gopurs-aff');
const frontend = join(archive, 'frontend/purs');
const candidateBuild = JSON.parse(readFileSync(join(candidate, 'build.json'), 'utf8'));
assert.equal(candidateBuild.status, 'passed');
const candidateCargo = candidateBuild.commands.find(record => record.label === 'cargo').command;
const profile = { opt_level: 3, debug: false,
  lto: JSON.parse(candidateCargo.find(arg => arg.startsWith('profile.release.lto=')).split('=')[1]),
  threaded: true, allocator: 'mimalloc' };
assert(candidateCargo.includes('profile.release.opt-level=3'));
assert(candidateCargo.includes('profile.release.debug=false'));
const sources = () => ({ gopurs: manifest(join(root, 'src')), optimizer: manifest(join(pbo, 'src')) });
const state = sources();
assert.deepEqual(state.gopurs, manifest(join(candidate, 'sources/0/src')));
assert.deepEqual(state.optimizer, manifest(join(candidate, 'sources/1/src')));
const result = resumeArg ? JSON.parse(readFileSync(join(directory, 'results.json'), 'utf8'))
  : { status: 'pending', started_at: new Date().toISOString(), candidate, sources: state,
    frontend_sha256: hash(readFileSync(frontend)), build_profile: profile,
    build_rust_sha256: hash(readFileSync(join(root, 'tools/build-rust.mjs'))), commands: [] };
const suffix = resumeArg ? '-resume-' + Date.now() : '';
if (resumeArg) {
  assert.equal(result.status, 'failed');
  assert.deepEqual(result.sources, state);
  assert.equal(result.candidate, candidate);
  assert.deepEqual(result.build_profile, profile);
  assert.equal(result.build_rust_sha256, hash(readFileSync(join(root, 'tools/build-rust.mjs'))));
  copyFileSync(join(directory, 'results.json'), join(directory, 'failed-results' + suffix + '.json'));
  result.status = 'pending'; result.resumed_at = new Date().toISOString(); delete result.failure;
}
const save = () => writeJson(join(directory, 'results.json'), result);
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  CARGO_NET_OFFLINE: 'true', TMPDIR: directory, GOPURS_NATIVE_TMPDIR: directory,
  GOPURS_KEEP_WORKSPACE: '1', GOPURS_PURS: frontend, PURS: frontend };
function command(label, cmd, args, cwd, extra = {}) {
  console.log(label);
  const record = run(directory, label + suffix, cmd, args, cwd, { ...env, ...extra }, 3600000);
  result.commands.push(record); save(); return record;
}
copyFileSync(fileURLToPath(import.meta.url), join(directory, 'qualify' + suffix + '.mjs'));
copyFileSync(join(root, 'tools/build-rust.mjs'), join(directory, 'build-rust' + suffix + '.mjs'));
save();
try {
  if (!result.go_bootstrap) {
    const go = command('build-go-and-js', 'npm', ['run', 'build:native', '--', '--keep-workspace'], root);
    result.go_bootstrap = readFileSync(go.stdout, 'utf8').match(/^Native bootstrap workspace: (.+)$/m)?.[1];
  }
  assert(result.go_bootstrap);
  assert.equal(hash(readFileSync(join(result.go_bootstrap, 'gopurs-native'))), hash(readFileSync(join(root, 'bin/gopurs-native'))));
  command('gopurs-preparation-tests', process.execPath, ['--test', '--test-concurrency=1',
    'tools/monomorphization.test.mjs', 'tools/preparation.test.mjs',
    'tools/build-native.test.mjs', 'tools/embed-runtime.test.mjs'], root);
  // Exact public -c path, with ordinary runtime defaults (no worker overrides).
  const rust = command('aff-rust-rebuild-default', './bin/test', ['-c'], aff, { GOPURS_RUST: '1' });
  const rustLog = readFileSync(rust.stdout, 'utf8');
  result.rust_bootstrap = rustLog.match(/^Rust-hosted gopurs workspace: (.+)$/m)?.[1];
  result.cargo_command = rustLog.match(/^\[cargo-build\] (cargo .+)$/m)?.[1];
  assert(result.cargo_command);
  for (const setting of ['profile.release.opt-level=3', 'profile.release.debug=false',
    `profile.release.lto=${JSON.stringify(profile.lto)}`]) assert(result.cargo_command.includes(setting), setting);
  assert(result.rust_bootstrap);
  const generated = directory => manifest(directory, file => /\.(rs|toml)$/.test(file));
  assert.deepEqual(generated(join(result.rust_bootstrap, 'rust')), generated(join(candidate, 'rust')),
    'Production Rust source differs from measured selection');
  result.identical_candidate_rust_files = generated(join(candidate, 'rust')).length; save();
  command('qualify-three-hosts', process.execPath, [
    fileURLToPath(new URL('../gopurs-aff/qualify-hosts.mjs', import.meta.url)),
    root, aff, result.rust_bootstrap, join(directory, 'hosts'),
  ], root);
  // These tests exercise the generic/native map FFI bridge and its Go race contract.
  command('go-memo-contract', process.execPath, [join(pbo, 'test/bounded-memo-native.mjs'), root], pbo);
  assert.deepEqual(sources(), state);
  assert.equal(result.build_rust_sha256, hash(readFileSync(join(root, 'tools/build-rust.mjs'))));
  result.executables = Object.fromEntries(['gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, { path: join(root, 'bin', name), sha256: hash(readFileSync(join(root, 'bin', name))) }]));
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: result.status, bootstrap: result.rust_bootstrap, executables: result.executables }, null, 2));
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
