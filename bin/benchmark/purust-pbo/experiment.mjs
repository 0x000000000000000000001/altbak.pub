// Reusable, archived experimental builds; production qualification is separate.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, cpSync, existsSync, lstatSync, mkdirSync, readFileSync, rmSync, symlinkSync, unlinkSync, writeFileSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler, nativeWorkspaceConfig, verifyTypedOutput } from '../../../../purust/purust/tools/native-workspace.mjs';

const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust/purust');
const pbo = resolve(root, '../../purescript-backend-optimizer-purust');
const [mode, archiveArg, label, selectionArg] = process.argv.slice(2);
assert(['init', 'build'].includes(mode) && archiveArg,
  'Usage: experiment.mjs init ARCHIVE | build ARCHIVE LABEL SELECTION.json');
const archive = resolve(archiveArg), workspace = join(archive, 'work');
const frontend = join(archive, 'frontend/purs');
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  PATH: [dirname(frontend), join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter) };
const sourceRoots = [root, pbo];
function sources(roots = sourceRoots) {
  return roots.map(path => ({ path, files: [
    ...manifest(join(path, 'src')).map(file => ({ ...file, path: 'src/' + file.path })),
    ...manifest(join(path, 'tests/runtime/perceus_ptr')).map(file => ({ ...file, path: 'tests/runtime/perceus_ptr/' + file.path })),
    ...['spago.yaml', 'spago.lock', 'package.json'].filter(name => existsSync(join(path, name)))
      .map(name => ({ path: name, sha256: hash(readFileSync(join(path, name))) })),
  ] }));
}
function copy(source, destination) {
  mkdirSync(dirname(destination), { recursive: true }); copyFileSync(source, destination);
}
function snapshot(directory, state) {
  for (const [index, source] of state.entries()) for (const file of source.files)
    copy(join(source.path, file.path), join(directory, 'sources', String(index), file.path));
}
if (mode === 'init') {
  assert(!existsSync(archive)); mkdirSync(workspace, { recursive: true });
  const purs = findTypedCompiler(root, process.env.PURUST_PURS);
  copy(purs, frontend); chmodSync(frontend, 0o755);
  const state = sources(); snapshot(join(archive, 'baseline'), state);
  for (const name of ['purust', 'purust.js', 'purust-native']) copy(join(root, 'bin', name), join(archive, 'baseline', name));
  writeJson(join(archive, 'baseline.json'), { started_at: new Date().toISOString(), sources: state,
    binaries: manifest(join(archive, 'baseline'), path => /\/purust(?:\.js|-native)?$/.test(path)),
    frontend: { origin: purs, sha256: hash(readFileSync(frontend)) },
    profile: 'O3, no LTO, threaded Arc, mimalloc, CARGO_INCREMENTAL=0, eight build jobs' });
  writeFileSync(join(workspace, 'spago.yaml'), nativeWorkspaceConfig(root));
  symlinkSync(join(root, 'src'), join(workspace, 'src'), 'dir');
  const prior = resolve(archive, '../purust-tast-20261002/qualification/qualification.json');
  const qualification = JSON.parse(readFileSync(prior, 'utf8'));
  const sha256 = hash(readFileSync(join(archive, 'baseline/purust-native')));
  assert.equal(sha256, qualification.stage2.sha256);
  writeJson(join(archive, 'before.json'), { binary: 'baseline/purust-native', sha256, qualification: prior });
  console.log(`Frozen ${sha256}; workspace ${workspace}`);
} else {
  assert(label && /^[a-z0-9-]+$/.test(label) && selectionArg);
  const out = join(archive, 'candidates', label); assert(!existsSync(out)); mkdirSync(out, { recursive: true });
  const selectionPath = resolve(selectionArg), selection = JSON.parse(readFileSync(selectionPath, 'utf8'));
  const sourceBase = selection.base ? resolve(dirname(selectionPath), selection.base) : join(archive, 'baseline/sources');
  const selectedRoots = [0, 1].map(index => join(workspace, 'sources', String(index)));
  for (const [index, path] of selectedRoots.entries()) {
    rmSync(path, { recursive: true, force: true });
    cpSync(join(sourceBase, String(index)), path, { recursive: true });
  }
  // Runtime templates are inputs to the JS bundle as well as embedded Rust.
  if (!existsSync(join(selectedRoots[0], 'tests/runtime/perceus_ptr')))
    cpSync(join(root, 'tests/runtime/perceus_ptr'), join(selectedRoots[0], 'tests/runtime/perceus_ptr'), { recursive: true });
  for (const override of selection.overrides ?? []) {
    assert([0, 1].includes(override.root) && !override.path.split('/').includes('..'));
    copy(resolve(dirname(selectionPath), override.from), join(selectedRoots[override.root], override.path));
  }
  const srcLink = join(workspace, 'src');
  if (existsSync(srcLink)) { assert(lstatSync(srcLink).isSymbolicLink()); unlinkSync(srcLink); }
  symlinkSync(join(selectedRoots[0], 'src'), srcLink, 'dir');
  writeFileSync(join(workspace, 'spago.yaml'), nativeWorkspaceConfig(root).replace(JSON.stringify(pbo), JSON.stringify(selectedRoots[1])));
  const jsWorkspace = join(workspace, 'js'); mkdirSync(jsWorkspace, { recursive: true });
  if (!existsSync(join(jsWorkspace, 'src'))) symlinkSync(join(selectedRoots[0], 'src'), join(jsWorkspace, 'src'), 'dir');
  if (!existsSync(join(jsWorkspace, 'tests'))) symlinkSync(join(selectedRoots[0], 'tests'), join(jsWorkspace, 'tests'), 'dir');
  writeFileSync(join(jsWorkspace, 'spago.yaml'), readFileSync(join(selectedRoots[0], 'spago.yaml'), 'utf8')
    .replace(/^(\s+path:\s*)(.+)$/gm, (_, prefix, value) => {
      const path = resolve(root, value.trim().replace(/^["']|["']$/g, ''));
      return prefix + JSON.stringify(path === pbo ? selectedRoots[1] : path);
    }));
  const state = sources(selectedRoots); snapshot(out, state);
  const result = { status: 'pending', started_at: new Date().toISOString(), label, selection, sources: state,
    frontend_sha256: hash(readFileSync(frontend)), commands: [] };
  const save = () => writeJson(join(out, 'build.json'), result);
  const command = (name, cmd, args, cwd) => {
    result.commands.push(run(out, name, cmd, args, cwd, env, 3600000)); save();
  };
  save();
  try {
    command('js-build', 'spago', ['build'], jsWorkspace);
    command('bundle', 'spago', ['bundle', '--module', 'Main', '--platform', 'node', '--outfile',
      join(out, 'purust.js'), '--bundle-type', 'app'], jsWorkspace);
    command('tast', 'spago', ['build'], workspace);
    result.tast = verifyTypedOutput(join(workspace, 'output')); save();
    const rust = join(workspace, 'rust'), generated = join(out, 'rust');
    command('generate', process.execPath, ['--expose-gc', '--stack-size=65536', '--max-old-space-size=16384',
      join(out, 'purust.js'), '--source', join(workspace, 'output'), '--out', generated, '--main', 'Main', '--threaded'], workspace);
    result.generated = manifest(generated, path => /\.(rs|toml)$/.test(path));
    // Keep unchanged Cargo input mtimes: repeated experiments share build
    // artifacts, while each candidate retains its complete generated source.
    const nextFiles = new Set(result.generated.map(file => file.path));
    for (const previous of manifest(rust, path => !path.includes('/target/') && /\.(rs|toml)$/.test(path)))
      if (!nextFiles.has(previous.path)) rmSync(join(rust, previous.path));
    for (const file of result.generated) {
      const destination = join(rust, file.path);
      if (!existsSync(destination) || hash(readFileSync(destination)) !== file.sha256)
        copy(join(generated, file.path), destination);
    }
    save();
    command('cargo', 'cargo', ['build', '--release', '--config', 'profile.release.lto=false',
      '--config', 'profile.release.opt-level=3', '--manifest-path', join(rust, 'Cargo.toml')], workspace);
    copy(join(rust, 'target/release/purust_output'), join(out, 'purust-native'));
    chmodSync(join(out, 'purust-native'), 0o755);
    assert.deepEqual(sources(selectedRoots), state, 'Frozen source changed during candidate build');
    result.binary_sha256 = hash(readFileSync(join(out, 'purust-native')));
    result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
    console.log(`${label}: ${result.binary_sha256}`);
  } catch (error) {
    result.status = 'failed'; result.error = error.stack; save(); throw error;
  }
}
