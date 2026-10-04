// Independent Purust bootstrap: JS/native generated-source identity before install.
import assert from 'node:assert/strict';
import { chmodSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync, renameSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { nativeWorkspaceConfig } from '../../../../purust/purust/tools/native-workspace.mjs';

const archive = resolve(process.argv[2]), label = process.argv[3] ?? 'purust-bootstrap-final';
assert(/^[a-z0-9-]+$/.test(label));
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust/purust');
const out = join(archive, label); assert(!existsSync(out)); mkdirSync(out);
const paths = [root, ...[...nativeWorkspaceConfig(root).matchAll(/^      path: (.+)$/gm)].map(match => JSON.parse(match[1]))];
const inputs = () => paths.map(path => ({ path, files: manifest(join(path, 'src')) }));
const state = { status: 'pending', started_at: new Date().toISOString(), sources: inputs(), commands: [] };
const copy = (source, target) => { mkdirSync(dirname(target), { recursive: true }); copyFileSync(source, target, constants.COPYFILE_FICLONE); };
for (const [index, source] of state.sources.entries())
  for (const file of source.files) copy(join(source.path, 'src', file.path), join(out, 'sources', String(index), 'src', file.path));
for (const name of ['purust-native', 'purust.js']) copy(join(root, 'bin', name), join(out, 'before', name));
const save = () => writeJson(join(out, 'results.json'), state);
save();
try {
  const record = run(out, 'self-host', 'npm', ['run', 'build:native', '--', '--self-host', '--keep-workspace'], root,
    { ...environment(), PURUST_PURS: join(archive, 'frontend/purs'), PURUST_NATIVE_TMPDIR: out,
      PURUST_NATIVE_OUTPUT: join(out, 'purust-native'), TMPDIR: out, GHCRTS: '-N2',
      CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_RELEASE_DEBUG: '0' }, 7200000);
  state.commands.push(record);
  const stdout = readFileSync(record.stdout, 'utf8');
  state.workspace = stdout.match(/^Native bootstrap workspace: (.+)$/m)?.[1];
  state.identical_generated_files = Number(stdout.match(/(\d+) generated Rust sources and Cargo manifests are byte-identical/)?.[1]);
  assert(state.workspace && Number.isInteger(state.identical_generated_files) && state.identical_generated_files > 0);
  state.tast = JSON.parse(readFileSync(join(state.workspace, 'verify-tast.log')));
  const generated = directory => manifest(directory, file =>
    (file.endsWith('.rs') || /(^|\/)Cargo\.toml$/.test(file)) && !/(^|\/)target\//.test(file));
  state.generated = generated(join(state.workspace, 'rust'));
  assert.deepEqual(generated(join(state.workspace, 'rust-stage2')), state.generated);
  assert.equal(state.generated.length, state.identical_generated_files);
  state.type_render_fixture_sha256 = hash(readFileSync(join(root, 'tools/test-native-type-render.rs')));
  state.commands.push(run(out, 'native-type-render', process.execPath,
    [join(root, 'tools/test-native-type-render.mjs'), join(state.workspace, 'rust-stage2'),
      resolve(archive, '../gopurs-purust-aff-20261002/inputs/gopurs-aff/output')], root,
    { ...environment(), TMPDIR: out, PURUST_NATIVE_TMPDIR: out,
      CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_RELEASE_DEBUG: '0' }, 1800000));
  assert.deepEqual(inputs(), state.sources, 'Compiler/runtime source changed during independent bootstrap');
  copy(join(root, 'bin/purust.js'), join(out, 'purust.js'));
  state.executables = Object.fromEntries(['purust-native', 'purust.js'].map(name => [name, hash(readFileSync(join(out, name)))]));
  const staged = join(root, 'bin/.purust-native-night');
  assert(!existsSync(staged));
  copy(join(out, 'purust-native'), staged); chmodSync(staged, 0o755);
  renameSync(staged, join(root, 'bin/purust-native'));
  state.installed = join(root, 'bin/purust-native');
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
