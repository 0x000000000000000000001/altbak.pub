// Qualify complete foreign identifiers and their native ABI before the final rerun.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { compareGeneratedSources } from '../../../../purust/purust/tools/native-workspace.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), site = resolve(tools, '../../..');
const root = resolve(site, '../purust/purust'), archive = resolve(process.argv[2]);
const directory = join(archive, process.argv[3] ?? 'primed-bootstrap'); assert(!existsSync(directory)); mkdirSync(directory);
const resumeWorkspace = process.argv[4] ? resolve(process.argv[4]) : null;
const before = join(archive, 'diagnostics/http-primed-before');
for (const name of ['regression-js', 'regression-codegen']) {
  assert.equal(JSON.parse(readFileSync(join(before, name + '.json'))).status, 'expected_failure');
}
const state = { status: 'running', started_at: new Date().toISOString(), diagnostics: before, commands: [],
  sources_before: { compiler: manifest(join(root, 'src')), optimizer: manifest(resolve(site, '../purescript-backend-optimizer-purust/src')) } };
const save = () => writeJson(join(directory, 'qualification.json'), state); save();
const env = { ...environment(), CARGO_BUILD_JOBS: '8', GHCRTS: '-N2', RUSTFLAGS: '-Awarnings',
  TMPDIR: directory, PURUST_FFI_KEEP: '1',
  PURUST_FFI_CORPUS: join(archive, 'cases/purust-foldable-traversable/input/sources') };
const execute = (label, command, args, cwd = root, overrides = {}, timeout = 360000) => {
  const record = run(directory, label, command, args, cwd, { ...env, ...overrides }, timeout);
  state.commands.push(record); save(); return record;
};
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
try {
  execute('foreign-corpus', process.execPath, [join(tools, 'verify-foreign-corpus.mjs'), archive, join(directory, 'foreign-corpus')]);
  env.PURUST_FFI_CORPUS = join(directory, 'foreign-corpus');
  state.foreign_corpus = { path: join(env.PURUST_FFI_CORPUS, 'corpus.json'), sha256: hash(readFileSync(join(env.PURUST_FFI_CORPUS, 'corpus.json'))) };
  for (const [label, source] of [
    ['native-before', join(before, 'src/Purust/ForeignTypes.rs')],
    ['unicode-before', join(archive, 'diagnostics/unicode-kind-before/src/Purust/ForeignTypes.rs')],
  ]) {
    const red = spawnSync(process.execPath, ['tests/codegen/native-compiler-ffi.mjs'], {
      cwd: root, env: { ...env, PURUST_FFI_SOURCE: source },
      encoding: 'utf8', timeout: 240000, maxBuffer: 64 * 1024 * 1024 });
    writeFileSync(join(directory, label + '.stdout'), red.stdout ?? '');
    writeFileSync(join(directory, label + '.stderr'), red.stderr ?? '');
    assert.ifError(red.error); assert.notEqual(red.status, 0); assert.match(red.stderr, /Handle/);
    state[label.replaceAll('-', '_')] = { status: 'expected_failure', exit_code: red.status, source_sha256: hash(readFileSync(source)) }; save();
  }
  execute('ffi-regressions', process.execPath,
    ['--test', '--test-concurrency=1', 'tests/codegen/foreign-types.mjs', 'tests/codegen/native-compiler-ffi.mjs']);
  if (resumeWorkspace) {
    const previousQualification = join(dirname(resumeWorkspace), 'qualification.json');
    const previous = JSON.parse(readFileSync(previousQualification));
    assert.equal(previous.status, 'failed'); assert.match(previous.error, /stage2-smoke/);
    assert.deepEqual(state.sources_before, previous.sources_before, 'Compiler sources changed: a fresh bootstrap is required');
    state.workspace = resumeWorkspace;
    const stage1 = join(resumeWorkspace, 'rust'), stage2 = join(resumeWorkspace, 'rust-stage2');
    state.reused_bootstrap = { previous_qualification: { path: previousQualification, sha256: hash(readFileSync(previousQualification)) },
      identical_files: compareGeneratedSources(stage1, stage2),
      binaries: [stage1, stage2].map(directory => {
        const path = join(directory, 'target/release/purust_output'); return { path, sha256: hash(readFileSync(path)) };
      }) };
    copy(join(stage2, 'target/release/purust_output'), join(directory, 'purust-native')); save();
    execute('stage2-smoke-retry', process.execPath, ['tools/test-native.mjs', '--keep-workspace'], root,
      { PURUST_NATIVE: join(directory, 'purust-native'), PURUST_NATIVE_TMPDIR: directory,
        PURUST_PURS: join(archive, 'frontend/purs') });
  } else {
    execute('self-host', process.execPath, ['tools/build-native.mjs', '--self-host', '--keep-workspace'], root,
      { PURUST_NATIVE_TMPDIR: directory, PURUST_NATIVE_OUTPUT: join(directory, 'purust-native'),
        PURUST_PURS: join(archive, 'frontend/purs') }, 3600000);
    state.workspace = join(directory, readdirSync(directory).find(name => name.startsWith('purust-native-build-')));
  }
  state.tast = JSON.parse(readFileSync(join(state.workspace, 'verify-tast.log')));
  state.bootstrap_identity = readFileSync(join(state.workspace, 'compare-native-sources.log'), 'utf8').trim();
  assert.match(state.bootstrap_identity, /byte-identical/);
  execute('class-shared-rc-arc', process.execPath, ['tests/codegen/class-shared.mjs']);

  // The original package suite checks plain HTTP, upgrades, cookies and local HTTPS.
  const parent = join(archive, 'revision1'), campaign = JSON.parse(readFileSync(join(parent, 'campaign.json')));
  const previousRows = JSON.parse(readFileSync(join(parent, 'libraries-final-results.json'))).results;
  for (const [label, name] of [['http', 'purust-node-http'], ['bigints', 'purust-js-bigints']]) {
    const previous = previousRows.find(row => row.name === name);
    const old = JSON.parse(readFileSync(previous.definition)), smoke = join(directory, label); mkdirSync(smoke);
    const caseRoot = join(smoke, 'cases', name), input = join(caseRoot, 'input');
    for (const file of old.input_manifest) copy(join(old.input, file.path), join(input, file.path));
    const definition = { ...old, input, input_manifest: manifest(input), historical_oracle: undefined };
    const definitionPath = join(caseRoot, 'definition.json'); writeJson(definitionPath, definition);
    const generated = join(caseRoot, 'canonical-generated'), native = join(caseRoot, 'native-generated');
    execute(label + '-js-generate', process.execPath, ['--expose-gc', '--stack-size=65536', '--max-old-space-size=16384',
      join(root, 'bin/purust.js'), ...definition.invocation, '--out', generated], input);
    execute(label + '-native-generate', join(directory, 'purust-native'), [...definition.invocation, '--out', native], input);
    state[label + '_identical_files'] = compareGeneratedSources(generated, native);
    if (label === 'bigints') assert.deepEqual(manifest(generated), manifest(previous.runs[0].retained_output));
    save();
    writeJson(join(smoke, 'campaign.json'), campaign);
    writeJson(join(smoke, 'application-inputs.json'), { scope: 'Untimed compiler qualification', results: [
      { name: definition.name, status: 'passed', definition: definitionPath },
    ] });
    execute(label + '-application', process.execPath, [join(tools, 'validate.mjs'), smoke, 'application-inputs.json'], tools, {}, 1200000);
    state[label + '_validation'] = { path: join(smoke, 'validation/results.json'), sha256: hash(readFileSync(join(smoke, 'validation/results.json'))) };
    assert.equal(JSON.parse(readFileSync(state[label + '_validation'].path)).status, 'passed');
  }
  copy(join(directory, 'purust-native'), join(root, 'bin/purust-native'));
  state.compilers = { purust: { directory: root, files: manifest(join(root, 'bin'), path => /\/(purust|purust\.js|purust-native)$/.test(path))
    .map(file => ({ ...file, path: 'bin/' + file.path })) } };
  state.sources = { compiler: manifest(join(root, 'src')), optimizer: manifest(resolve(site, '../purescript-backend-optimizer-purust/src')) };
  assert.deepEqual(state.sources, state.sources_before);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, identity: state.bootstrap_identity, http_files: state.http_identical_files }, null, 2));
} catch (error) {
  state.status = 'failed'; state.error = error.stack; state.finished_at = new Date().toISOString(); save(); throw error;
}
