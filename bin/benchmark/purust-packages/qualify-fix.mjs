// Qualify the unbounded foreign-type scan, retaining the failing and fixed runs.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { constants, copyFileSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), root = join(dirname(site), 'purust/purust');
const archive = resolve(process.argv[2]), directory = join(archive, process.argv[3] ?? 'bootstrap'); mkdirSync(directory);
const state = { status: 'running', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(directory, 'qualification.json'), state); save();
const env = { ...environment(), CARGO_BUILD_JOBS: '8', GHCRTS: '-N2', RUSTFLAGS: '-Awarnings',
  TMPDIR: directory, PURUST_FFI_KEEP: '1',
  PURUST_FFI_CORPUS: join(archive, 'cases/purust-foldable-traversable/input/sources') };
try {
  const source = join(archive, 'diagnostics/foreign-types-before.rs');
  const red = spawnSync(process.execPath, ['tests/codegen/native-compiler-ffi.mjs'], {
    cwd: root, env: { ...env, PURUST_FFI_SOURCE: source }, encoding: 'utf8', timeout: 240000, maxBuffer: 64 * 1024 * 1024 });
  assert.ifError(red.error);
  writeFileSync(join(directory, 'regression-before.stdout'), red.stdout ?? '');
  writeFileSync(join(directory, 'regression-before.stderr'), red.stderr ?? '');
  assert.notEqual(red.status, 0); assert.match(red.stderr, /BacktrackLimitExceeded/);
  state.regression_before = { status: 'expected_failure', exit_code: red.status, original_source_sha256: hash(readFileSync(source)) }; save();
  state.commands.push(run(directory, 'regressions-after', process.execPath,
    ['--test', '--test-concurrency=1', 'tests/codegen/foreign-types.mjs', 'tests/codegen/native-compiler-ffi.mjs'], root, env, 360000)); save();
  const oldSmoke = spawnSync(process.execPath, ['tools/test-native.mjs', '--keep-workspace'], {
    cwd: root, env: { ...env, PURUST_NATIVE: join(archive, 'compilers/purust/bin/purust-native'),
      PURUST_NATIVE_TMPDIR: directory, PURUST_PURS: join(archive, 'frontend/purs') },
    encoding: 'utf8', timeout: 360000, maxBuffer: 64 * 1024 * 1024 });
  assert.ifError(oldSmoke.error);
  writeFileSync(join(directory, 'negative-zero-before.stdout'), oldSmoke.stdout ?? '');
  writeFileSync(join(directory, 'negative-zero-before.stderr'), oldSmoke.stderr ?? '');
  assert.notEqual(oldSmoke.status, 0); assert.match(oldSmoke.stderr, /Node\/native output mismatch/);
  state.negative_zero_before = { status: 'expected_failure', exit_code: oldSmoke.status }; save();
  const oldSort = spawnSync(process.execPath, ['--test-reporter=tap', join(dirname(site), 'purescript-backend-optimizer-purust/test/module-sort.mjs'), join(root, 'output')], {
    cwd: root, env, encoding: 'utf8', timeout: 120000 });
  assert.ifError(oldSort.error);
  writeFileSync(join(directory, 'module-sort-before.stdout'), oldSort.stdout ?? '');
  writeFileSync(join(directory, 'module-sort-before.stderr'), oldSort.stderr ?? '');
  assert.notEqual(oldSort.status, 0); assert.match(oldSort.stdout, /not ok.*module ranks/);
  state.module_sort_before = { status: 'expected_failure', exit_code: oldSort.status }; save();
  state.commands.push(run(directory, 'self-host', process.execPath,
    ['tools/build-native.mjs', '--self-host', '--keep-workspace'], root,
    { ...env, PURUST_NATIVE_TMPDIR: directory, PURUST_NATIVE_OUTPUT: join(directory, 'purust-native'),
      PURUST_PURS: join(archive, 'frontend/purs') }, 3600000)); save();
  const workspace = readdirSync(directory).find(name => name.startsWith('purust-native-build-')); assert(workspace);
  state.workspace = join(directory, workspace);
  state.commands.push(run(directory, 'pbo-numeric-negation', process.execPath,
    [join(dirname(site), 'purescript-backend-optimizer-purust/test/numeric-negate.mjs'), join(root, 'output')], root, env, 120000)); save();
  state.commands.push(run(directory, 'pbo-module-sort', process.execPath,
    [join(dirname(site), 'purescript-backend-optimizer-purust/test/module-sort.mjs'), join(root, 'output')], root, env, 120000)); save();
  state.tast = JSON.parse(readFileSync(join(state.workspace, 'verify-tast.log')));
  state.bootstrap_identity = readFileSync(join(state.workspace, 'compare-native-sources.log'), 'utf8').trim();
  assert.match(state.bootstrap_identity, /byte-identical/);
  copyFileSync(join(directory, 'purust-native'), join(root, 'bin/purust-native'), constants.COPYFILE_FICLONE);
  state.compilers = { purust: { directory: root, files: manifest(join(root, 'bin'), path => /\/(purust|purust\.js|purust-native)$/.test(path))
    .map(file => ({ ...file, path: 'bin/' + file.path })) } };
  state.sources = { compiler: manifest(join(root, 'src')), optimizer: manifest(join(dirname(site), 'purescript-backend-optimizer-purust/src')) };
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, tast: state.tast, identity: state.bootstrap_identity, compilers: state.compilers }, null, 2));
} catch (error) {
  state.status = 'failed'; state.error = error.stack; save(); throw error;
}
