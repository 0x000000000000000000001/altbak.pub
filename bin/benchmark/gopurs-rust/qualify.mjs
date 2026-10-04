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
const purust = resolve(root, '../../purust/purust');
const independentBootstrap = join(archive, 'purust-bootstrap-final/results.json');
if (existsSync(independentBootstrap)) {
  const bootstrap = JSON.parse(readFileSync(independentBootstrap));
  assert.equal(bootstrap.status, 'passed');
  for (const name of ['purust-native', 'purust.js'])
    assert.equal(hash(readFileSync(join(purust, 'bin', name))), bootstrap.executables[name]);
}
const frontend = join(archive, 'frontend/purs');
const candidateBuild = JSON.parse(readFileSync(join(candidate, 'build.json'), 'utf8'));
assert.equal(candidateBuild.status, 'passed');
const candidateCargo = candidateBuild.commands.find(record => record.label === 'cargo').command;
const profile = { opt_level: 3, debug: false,
  lto: JSON.parse(candidateCargo.find(arg => arg.startsWith('profile.release.lto=')).split('=')[1]),
  threaded: true, allocator: 'mimalloc',
  pgo: (candidateBuild.rustflags ?? []).some(flag => flag.startsWith('profile-use=')) };
assert(candidateCargo.includes('profile.release.opt-level=3'));
assert(candidateCargo.includes('profile.release.debug=false'));
const sources = () => ({ gopurs: manifest(join(root, 'src')), optimizer: manifest(join(pbo, 'src')) });
const state = sources();
const buildTools = () => Object.fromEntries(['build-rust.mjs', 'pgo.mjs', 'native-workspace.mjs', 'command-runner.mjs']
  .map(name => [name, hash(readFileSync(join(root, 'tools', name)))]));
assert.deepEqual(state.gopurs, manifest(join(candidate, 'sources/0/src')));
assert.deepEqual(state.optimizer, manifest(join(candidate, 'sources/1/src')));
const result = resumeArg ? JSON.parse(readFileSync(join(directory, 'results.json'), 'utf8'))
  : { status: 'pending', started_at: new Date().toISOString(), candidate, sources: state,
    frontend_sha256: hash(readFileSync(frontend)), build_profile: profile, linker: candidateBuild.linker ?? null,
    build_rust_sha256: hash(readFileSync(join(root, 'tools/build-rust.mjs'))), build_tools: buildTools(), commands: [] };
const suffix = resumeArg ? '-resume-' + Date.now() : '';
if (resumeArg) {
  assert.equal(result.status, 'failed');
  assert.deepEqual(result.sources, state);
  assert.equal(result.candidate, candidate);
  assert.deepEqual(result.build_profile, profile);
  assert.deepEqual(result.linker, candidateBuild.linker ?? null);
  assert.equal(result.build_rust_sha256, hash(readFileSync(join(root, 'tools/build-rust.mjs'))));
  assert.deepEqual(result.build_tools, buildTools());
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
for (const name of Object.keys(result.build_tools))
  copyFileSync(join(root, 'tools', name), join(directory, name.replace(/\.mjs$/, suffix + '.mjs')));
save();
try {
  if (!result.go_bootstrap) {
    const go = command('build-go-and-js', 'npm', ['run', 'build:native', '--', '--keep-workspace'], root);
    result.go_bootstrap = readFileSync(go.stdout, 'utf8').match(/^Native bootstrap workspace: (.+)$/m)?.[1];
  }
  assert(result.go_bootstrap);
  assert.equal(hash(readFileSync(join(result.go_bootstrap, 'gopurs-native'))), hash(readFileSync(join(root, 'bin/gopurs-native'))));
  command('gopurs-preparation-tests', process.execPath, ['--test', '--test-concurrency=1', '--test-timeout=60000',
    'tools/monomorphization.test.mjs', 'tools/preparation.test.mjs',
    'tools/emission.test.mjs', 'tools/build-native.test.mjs', 'tools/embed-runtime.test.mjs', 'tools/pgo.test.mjs'], root);
  // Exact public -c path, with ordinary runtime defaults (no worker overrides).
  const rust = command('aff-rust-rebuild-default', './bin/test', ['-c'], aff, { GOPURS_RUST: '1' });
  const rustLog = readFileSync(rust.stdout, 'utf8');
  result.rust_bootstrap = rustLog.match(/^Rust-hosted gopurs workspace: (.+)$/m)?.[1];
  result.cargo_command = rustLog.match(/^\[cargo-build\] (cargo .+)$/m)?.[1];
  assert(result.cargo_command);
  for (const setting of ['profile.release.opt-level=3', 'profile.release.debug=false',
    `profile.release.lto=${JSON.stringify(profile.lto)}`]) assert(result.cargo_command.includes(setting), setting);
  if (result.linker) {
    assert(result.cargo_command.includes('link-arg=' + result.linker.option));
    assert.equal(hash(readFileSync(result.linker.path)), result.linker.sha256);
  }
  assert(result.rust_bootstrap);
  const generated = directory => manifest(directory, file => /\.(rs|toml)$/.test(file));
  assert.deepEqual(generated(join(result.rust_bootstrap, 'rust')), generated(join(candidate, 'rust')),
    'Production Rust source differs from measured selection');
  result.identical_candidate_rust_files = generated(join(candidate, 'rust')).length; save();
  if (profile.pgo) {
    const path = join(result.rust_bootstrap, 'pgo-profile.json');
    const pgo = JSON.parse(readFileSync(path));
    assert.equal(pgo.status, 'passed'); assert.equal(pgo.enabled, true);
    assert.deepEqual(pgo.source.manifest, candidateBuild.generated);
    assert.equal(pgo.source.sha256, hash(JSON.stringify(pgo.source.manifest)));
    assert.equal(pgo.passes.length, 3);
    for (const pass of pgo.passes) {
      assert.equal(pass.files, pgo.training.oracle.files);
      assert.equal(pass.sha256, pgo.training.oracle.sha256);
    }
    assert(pgo.training.oracle.files > 400);
    assert(pgo.training.modules.some(entry => entry.module === 'Main'));
    assert(!pgo.training.modules.some(entry => entry.module === 'Test.Main'));
    assert.deepEqual(manifest(join(result.rust_bootstrap, 'pgo/training')), pgo.training.manifest);
    assert.equal(hash(readFileSync(resolve(result.rust_bootstrap, pgo.profile.merged.path))), pgo.profile.merged.sha256);
    assert.deepEqual(pgo.profile.use_flags, ['-C', 'profile-use=' + resolve(result.rust_bootstrap, pgo.profile.merged.path)]);
    assert.notEqual(pgo.binary.sha256, pgo.binary.bootstrap_sha256);
    assert.equal(hash(readFileSync(resolve(result.rust_bootstrap, pgo.binary.path))), pgo.binary.sha256);
    assert.equal(hash(readFileSync(join(root, 'bin/gopurs-rust'))), pgo.binary.sha256);
    const heldOut = new Set(JSON.parse(readFileSync(resolve(archive, '../gopurs-purust-aff-20261002/results.json')))
      .frozen_files.inputs.filter(file => file.path.endsWith('/corefn.json')).map(file => file.path.split('/').at(-2)));
    result.pgo = { path, sha256: hash(readFileSync(path)), binary_sha256: pgo.binary.sha256,
      training_modules: pgo.training.modules.length, training_sha256: pgo.training.sha256,
      shared_library_modules: pgo.training.modules.map(entry => entry.module).filter(name => heldOut.has(name)).sort(),
      profile_sha256: pgo.profile.merged.sha256, source_sha256: pgo.source.sha256, passes: pgo.passes };
    assert(!result.pgo.shared_library_modules.includes('Test.Main'));
    save();
  }
  command('qualify-three-hosts', process.execPath, [
    fileURLToPath(new URL('../gopurs-aff/qualify-hosts.mjs', import.meta.url)),
    root, aff, result.rust_bootstrap, join(directory, 'hosts'),
  ], root);
  // These tests exercise the generic/native map FFI bridge and its Go race contract.
  command('go-memo-contract', process.execPath, [join(pbo, 'test/bounded-memo-native.mjs'), root], pbo);
  assert.deepEqual(sources(), state);
  assert.equal(result.build_rust_sha256, hash(readFileSync(join(root, 'tools/build-rust.mjs'))));
  assert.deepEqual(result.build_tools, buildTools());
  result.executables = Object.fromEntries(['gopurs', 'gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, { path: join(root, 'bin', name), sha256: hash(readFileSync(join(root, 'bin', name))) }]));
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: result.status, bootstrap: result.rust_bootstrap, executables: result.executables }, null, 2));
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
