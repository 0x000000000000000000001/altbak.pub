// Frozen experimental builds of gopurs's Go generator on the Rust host.
import assert from 'node:assert/strict';
import { chmodSync, constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler, nativeWorkspaceConfig, verifyTypedOutput } from '../../../../gopurs/gopurs/tools/native-workspace.mjs';
import { cargoProfile, readProfile } from './cargo-profile.mjs';

const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../gopurs/gopurs');
const pbo = resolve(root, '../../purescript-backend-optimizer-gopurs');
const purust = resolve(root, '../../purust/purust');
const [mode, archiveArg, label, profileArg] = process.argv.slice(2);
assert(['init', 'build'].includes(mode) && archiveArg,
  'experiment.mjs init ARCHIVE EXPECTED_RUST_SHA256 [PROFILE.json] | build ARCHIVE LABEL [PROFILE.json]');
const archive = resolve(archiveArg), work = join(archive, 'work'), frontend = join(archive, 'frontend/purs');
function copy(from, to) { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); }
function sources(roots = [root, pbo]) {
  return roots.map(path => ({ path, files: [
    ...manifest(join(path, 'src')).map(file => ({ ...file, path: 'src/' + file.path })),
    ...['spago.yaml', 'spago.lock', 'package.json'].filter(name => existsSync(join(path, name)))
      .map(name => ({ path: name, sha256: hash(readFileSync(join(path, name))) })),
  ] }));
}
function snapshot(out, state) {
  for (const [index, source] of state.entries()) for (const file of source.files)
    copy(join(source.path, file.path), join(out, 'sources', String(index), file.path));
}
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  CARGO_NET_OFFLINE: 'true', GOWORK: 'off',
  PATH: [dirname(frontend), join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter) };
const configuration = readProfile(profileArg);
if (mode === 'init') {
  assert(label && /^[a-f0-9]{64}$/.test(label), 'Supply the qualified Rust baseline SHA-256');
  assert.equal(hash(readFileSync(join(root, 'bin/gopurs-rust'))), label);
  assert(!existsSync(join(archive, 'baseline')));
  mkdirSync(work, { recursive: true });
  const purs = findTypedCompiler(root, process.env.GOPURS_PURS);
  copy(purs, frontend); chmodSync(frontend, 0o755);
  const state = sources(); snapshot(join(archive, 'baseline'), state);
  for (const name of ['bin/gopurs', 'bin/gopurs.js', 'bin/gopurs-native', 'bin/gopurs-rust', 'package.json',
    'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js'])
    copy(join(root, name), join(archive, 'baseline/compiler', name));
  copy(join(purust, 'bin/purust-native'), join(archive, 'bootstrap/purust-native'));
  copy(join(purust, 'bin/purust.js'), join(archive, 'bootstrap/purust.js'));
  const baseline = { started_at: new Date().toISOString(), sources: state,
    compiler: manifest(join(archive, 'baseline/compiler')), bootstrap: manifest(join(archive, 'bootstrap')),
    frontend: { origin: purs, sha256: hash(readFileSync(frontend)) },
    profile: `O3, LTO=${configuration.lto}, no debug, threaded Arc, mimalloc, CARGO_INCREMENTAL=0, eight build jobs`,
    profile_configuration: configuration };
  assert.equal(hash(readFileSync(join(archive, 'baseline/compiler/bin/gopurs-rust'))), label);
  baseline.rust_sha256 = label;
  copy(fileURLToPath(import.meta.url), join(archive, 'experiment.mjs'));
  copy(fileURLToPath(new URL('./cargo-profile.mjs', import.meta.url)), join(archive, 'cargo-profile.mjs'));
  writeJson(join(archive, 'baseline.json'), baseline);
  console.log('Frozen gopurs hosts, source trees, bootstrap compilers and frontend');
} else {
  assert(label && /^[a-z0-9-]+$/.test(label));
  const out = join(archive, 'candidates', label); assert(!existsSync(out)); mkdirSync(out, { recursive: true });
  copy(fileURLToPath(import.meta.url), join(out, 'build-harness.mjs'));
  copy(fileURLToPath(new URL('./cargo-profile.mjs', import.meta.url)), join(out, 'cargo-profile.mjs'));
  const inputRoots = process.env.GOPURS_EXPERIMENT_SOURCE_ROOTS
    ? JSON.parse(process.env.GOPURS_EXPERIMENT_SOURCE_ROOTS).map(path => resolve(path)) : [root, pbo];
  assert.equal(inputRoots.length, 2);
  const state = sources(inputRoots); snapshot(out, state);
  for (const i of [0, 1]) {
    const destination = join(work, 'sources', String(i));
    rmSync(destination, { recursive: true, force: true });
    cpSync(join(out, 'sources', String(i)), destination, { recursive: true });
  }
  if (!existsSync(join(work, 'src'))) symlinkSync(join(work, 'sources/0/src'), join(work, 'src'), 'dir');
  let workspaceConfig = nativeWorkspaceConfig(root, { runtime: 'rust', purust });
  for (const [from, to] of Object.entries(JSON.parse(process.env.GOPURS_EXPERIMENT_RUNTIME_OVERRIDES ?? '{}'))) {
    const original = 'path: ' + JSON.stringify(resolve(from));
    assert(workspaceConfig.includes(original), 'Unknown runtime override: ' + from);
    workspaceConfig = workspaceConfig.replace(original, 'path: ' + JSON.stringify(resolve(to)));
  }
  workspaceConfig = workspaceConfig.replace(JSON.stringify(pbo), JSON.stringify(join(work, 'sources/1')));
  writeFileSync(join(work, 'spago.yaml'), workspaceConfig);
  writeFileSync(join(out, 'spago.yaml'), workspaceConfig);
  const result = { status: 'pending', label, started_at: new Date().toISOString(), sources: state, commands: [] };
  const save = () => writeJson(join(out, 'build.json'), result);
  const command = (name, cmd, args, cwd = work) => {
    console.log(`[${label}] ${name}`);
    result.commands.push(run(out, name, cmd, args, cwd, env, 3600000)); save();
  };
  save();
  try {
    command('tast', 'spago', ['build']);
    result.tast = verifyTypedOutput(join(work, 'output'), frontend); save();
    const rust = join(work, 'rust'), generated = join(out, 'rust');
    const generatorSource = resolve(process.env.GOPURS_EXPERIMENT_GENERATOR ?? join(archive, 'bootstrap/purust-native'));
    const javascriptGenerator = generatorSource.endsWith('.js');
    const generator = join(out, javascriptGenerator ? 'generator.js' : 'generator');
    copy(generatorSource, generator); chmodSync(generator, 0o755);
    result.generator = { origin: generatorSource, path: generator, sha256: hash(readFileSync(generator)) };
    const runtimeInputs = [...workspaceConfig.matchAll(/^      path: (.+)$/gm)]
      .map(match => JSON.parse(match[1])).filter(path => path !== join(work, 'sources/1'));
    const runtimeState = sources(runtimeInputs);
    snapshot(join(out, 'runtime'), runtimeState);
    result.runtime_sources = runtimeState;
    result.purust_sources = manifest(join(purust, 'src'));
    for (const file of result.purust_sources) copy(join(purust, 'src', file.path), join(out, 'purust-sources', file.path));
    save();
    command('generate', javascriptGenerator ? process.execPath : generator,
      [...(javascriptGenerator ? ['--stack-size=65536', generator] : []), '--source', join(work, 'output'),
        '--out', generated, '--main', 'Main', '--threaded']);
    assert.deepEqual(sources(runtimeInputs), runtimeState, 'Runtime source changed during generation');
    copy(join(root, 'tools/ffi-gen/rust-build.rs'), join(generated, 'Purs_Gopurs_FfiSupport/build.rs'));
    result.generated = manifest(generated, path => /\.(rs|toml)$/.test(path));
    const next = new Set(result.generated.map(file => file.path));
    for (const file of manifest(rust, path => !path.includes('/target/') && /\.(rs|toml)$/.test(path)))
      if (!next.has(file.path)) rmSync(join(rust, file.path));
    for (const file of result.generated) {
      const destination = join(rust, file.path);
      if (!existsSync(destination) || hash(readFileSync(destination)) !== file.sha256)
        copy(join(generated, file.path), destination);
    }
    const parser = join(rust, 'Purs_Gopurs_FfiSupport/native-parser/libgopurs_ffi.a');
    mkdirSync(dirname(parser), { recursive: true });
    if (!existsSync(parser)) command('parser', 'go', ['build', '-trimpath', '-tags=carchive',
      '-buildmode=c-archive', '-o', parser, '.'], join(root, 'tools/ffi-gen'));
    result.parser_sha256 = hash(readFileSync(parser));
    const cargo = cargoProfile(configuration, env, rust, join(work, 'target'));
    Object.assign(result, { profile: cargo.profile, linker: cargo.linker, rustflags: cargo.rustflags }); save();
    result.commands.push(run(out, 'cargo', 'cargo', cargo.args, work, cargo.env, 3600000)); save();
    copy(join(work, 'target/release/purust_output'), join(out, 'gopurs-rust'));
    chmodSync(join(out, 'gopurs-rust'), 0o755);
    // Compilation reads the frozen copy, so later live edits may prepare the
    // next candidate while this one builds. Verify the actual build inputs.
    assert.deepEqual(sources([0, 1].map(i => join(work, 'sources', String(i)))).map(s => s.files),
      state.map(s => s.files), 'Frozen source changed during candidate build');
    result.binary_sha256 = hash(readFileSync(join(out, 'gopurs-rust')));
    result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
    console.log(`${label}: ${result.binary_sha256}`);
  } catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
}
