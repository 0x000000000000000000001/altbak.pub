// Retain the production self-host qualification, without replacing the installed
// compiler before the independent output comparisons and application checks.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, rmSync } from 'node:fs';
import { cpus, totalmem } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler } from '../../../../purust/purust/tools/native-workspace.mjs';

const [archiveArg, ...flags] = process.argv.slice(2);
assert(archiveArg && flags.every(flag => flag === '--discard-stage1-cache'),
  'Usage: qualify-compiler.mjs NEW_ARCHIVE_DIRECTORY [--discard-stage1-cache]');
const archive = resolve(archiveArg);
assert(!existsSync(archive), archive);
mkdirSync(archive, { recursive: true });
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust/purust');
const pbo = resolve(root, '../../purescript-backend-optimizer-purust');
// Bring embedded runtime sources up to date before taking the source snapshot.
run(archive, 'embed-native-runtime', process.execPath, [join(root, 'tools/embed-native-runtime.mjs')], root);
const localPackages = [...readFileSync(join(root, 'spago.yaml'), 'utf8').matchAll(/^\s+path:\s*(.+)\s*$/gm)]
  .map(([, value]) => resolve(root, value.trim().replace(/^["']|["']$/g, '')));
const sourceRoots = [...new Set([root, pbo, ...localPackages,
  ...readdirSync(dirname(root)).filter(name => name.startsWith('purust-'))
    .map(name => join(dirname(root), name)).filter(path => existsSync(join(path, 'spago.yaml')))])];
const sourceFilter = path => !path.includes('/node_modules/') && !path.includes('/.spago/') && !path.includes('/target/') &&
  /\.(purs|js|mjs|rs|json|yaml|toml)$/.test(path);
const snapshotSources = () => sourceRoots.map(path => ({ path, files: [
  ...manifest(join(path, 'src'), sourceFilter).map(file => ({ ...file, path: 'src/' + file.path })),
  ...(path === root ? ['tools', 'tests/runtime/perceus_ptr'].flatMap(directory =>
    manifest(join(path, directory), sourceFilter).map(file => ({ ...file, path: directory + '/' + file.path }))) : []),
  ...['spago.yaml', 'spago.lock', 'package.json', 'package-lock.json'].filter(name => existsSync(join(path, name))).map(name =>
    ({ path: name, sha256: hash(readFileSync(join(path, name))) })),
] }));
const sources = snapshotSources();
for (const [index, source] of sources.entries()) for (const file of source.files) {
  const destination = join(archive, 'sources', String(index), file.path);
  mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(join(source.path, file.path), destination);
}
const before = {};
for (const name of ['purust', 'purust.js', 'purust-native']) {
  const source = join(root, 'bin', name), destination = join(archive, 'before', name);
  mkdirSync(dirname(destination), { recursive: true }); copyFileSync(source, destination);
  before[name] = hash(readFileSync(source));
}
const purs = findTypedCompiler(root, process.env.PURUST_PURS);
const frontend = { origin: purs, path: join(archive, 'frontend/purs'), sha256: hash(readFileSync(purs)) };
mkdirSync(dirname(frontend.path)); copyFileSync(purs, frontend.path);
const version = (command, args) => execFileSync(command, args, { encoding: 'utf8', env: environment() }).trim();
const host = { model: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), architecture: process.arch,
  os: version('sw_vers', []), node: process.version, rustc: version('rustc', ['--version']),
  cargo: version('cargo', ['--version']), spago: version(join(root, 'node_modules/.bin/spago'), ['--version']),
  purs: version(frontend.path, ['--version']) };
for (const name of ['purust-tast/qualify-compiler.mjs', 'gopurs-aff/common.mjs']) {
  const destination = join(archive, 'harness', name); mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(resolve(dirname(fileURLToPath(import.meta.url)), '..', name), destination);
}
const resultPath = join(archive, 'qualification.json');
const result = { started_at: new Date().toISOString(), status: 'pending', root, pbo, sources, before, frontend, host,
  harness: manifest(join(archive, 'harness')),
  protocol: { optimization: 'O3, no LTO', runtime: 'threaded Arc',
    stages: ['production JS build', 'JS-generated native stage 1', 'stage-1-generated native stage 2', 'independent fresh smoke'],
    output_comparison: 'Exact generated Rust sources and Cargo manifests between JS and stage 1',
    installed_native_replaced: false } };
const save = () => writeJson(resultPath, result);
save();
try {
  const buildRoot = join(archive, 'build'); mkdirSync(buildRoot);
  const binary = join(archive, 'compiler-final');
  const env = { ...environment(), PURUST_PURS: frontend.path, PURUST_NATIVE_OUTPUT: binary, PURUST_NATIVE_TMPDIR: buildRoot,
    PURUST_NATIVE_OPT_LEVEL: '3', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', GHCRTS: '-N2' };
  result.build = run(archive, 'self-host', process.execPath,
    [join(root, 'tools/build-native.mjs'), '--self-host', '--keep-workspace'], root, env, 3600000);
  save();
  assert.equal(hash(readFileSync(join(root, 'bin/purust-native'))), before['purust-native'], 'Installed native changed');
  assert.equal(hash(readFileSync(frontend.path)), frontend.sha256, 'Frozen frontend changed');
  assert.deepEqual(snapshotSources(), sources, 'Compiler, optimizer or library sources changed during qualification');
  const workspaces = readdirSync(buildRoot).filter(name => name.startsWith('purust-native-build-'));
  assert.equal(workspaces.length, 1);
  const workspace = join(buildRoot, workspaces[0]);
  const stage1 = join(workspace, 'rust/target/release/purust_output');
  const stage2 = join(workspace, 'rust-stage2/target/release/purust_output');
  copyFileSync(stage1, join(archive, 'compiler-stage1'));
  copyFileSync(join(root, 'bin/purust.js'), join(archive, 'compiler-final.js'));
  assert.equal(hash(readFileSync(binary)), hash(readFileSync(stage2)));
  Object.assign(result, { status: 'passed', finished_at: new Date().toISOString(), workspace,
    typed_metadata: JSON.parse(readFileSync(join(workspace, 'verify-tast.log'), 'utf8')),
    comparison: readFileSync(join(workspace, 'compare-native-sources.log'), 'utf8').trim(),
    stage1: { path: join(archive, 'compiler-stage1'), sha256: hash(readFileSync(stage1)) },
    stage2: { path: binary, sha256: hash(readFileSync(binary)) },
    javascript: { path: join(archive, 'compiler-final.js'), sha256: hash(readFileSync(join(archive, 'compiler-final.js'))) },
    smoke_log: join(workspace, 'stage2-smoke.log'), sources_unchanged: true });
  save();
  if (flags.includes('--discard-stage1-cache')) {
    const cache = join(workspace, 'rust/target');
    assert(existsSync(join(cache, 'CACHEDIR.TAG')));
    assert.equal(hash(readFileSync(result.stage1.path)), result.stage1.sha256);
    result.cache_cleanup = { path: cache, retained_binary: result.stage1,
      retained: 'all generated sources, manifests, build logs, stage-1 and stage-2 executables', status: 'planned' };
    save(); rmSync(cache, { recursive: true, maxRetries: 3 });
    result.cache_cleanup.status = 'complete'; save();
  }
  console.log(JSON.stringify({ ...result, sources: undefined }, null, 2));
} catch (error) {
  result.failure = { at: new Date().toISOString(), message: error.message, stack: error.stack }; save(); throw error;
}
