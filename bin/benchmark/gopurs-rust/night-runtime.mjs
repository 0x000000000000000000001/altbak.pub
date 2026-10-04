// Isolate the Aff IO completion change on the frozen control's generated Rust.
import assert from 'node:assert/strict';
import { chmodSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { cargoProfile } from './cargo-profile.mjs';
import { threadedRust } from '../../../../purust/purust/src/Purust/Threading.js';

const archive = resolve(process.argv[2]), label = process.argv[3] ?? 'native-io-pool';
assert(/^[a-z0-9-]+$/.test(label));
const here = dirname(fileURLToPath(import.meta.url)), root = resolve(here, '../../../..');
const source = join(archive, 'candidates/control'), out = join(archive, 'candidates', label), work = join(archive, 'work');
assert(!existsSync(out)); mkdirSync(out);
const control = JSON.parse(readFileSync(join(source, 'build.json')));
const state = { status: 'pending', label, started_at: new Date().toISOString(), source_candidate: source,
  sources: control.sources, tast: control.tast, commands: [] };
const save = () => writeJson(join(out, 'build.json'), state);
const env = { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true' };
const command = (name, executable, args, cwd = work, environment = env) => {
  console.log(name); state.commands.push(run(out, name, executable, args, cwd, environment, 3600000)); save();
};
const compare = (name, variants) => {
  const selection = join(archive, name + '-selection.json');
  writeJson(selection, { rounds: 5, variants });
  command(name + '-compare', process.execPath, [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, name + '-runs'), selection]);
};
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
save();
try {
  if (label === 'native-io-pool') {
    command('directive-contracts', process.execPath, [join(here, 'check-pbo.mjs'), join(work, 'output'), join(archive, 'directive-contracts')]);
    compare('directive-chain-v2', [
      { name: 'control', binary: 'candidates/control/gopurs-rust' },
      { name: 'directive-chain', binary: 'candidates/directive-chain-v2/gopurs-rust' },
    ]);
  }
  for (const name of ['sources', 'rust']) command('freeze-' + name, 'cp', ['-cR', join(source, name), join(out, name)], archive);
  const original = readFileSync(join(archive, 'aff-original.rs'), 'utf8');
  const current = readFileSync(join(root, 'purust/purust-aff/src/Effect/Aff.rs'), 'utf8');
  copy(join(archive, 'aff-original.rs'), join(out, 'aff-original.rs'));
  writeFileSync(join(out, 'aff-candidate.rs'), current);
  const file = join(out, 'rust/Purs_Effect_Aff/src/lib.rs'), rust = readFileSync(file, 'utf8');
  assert.equal(rust.split(threadedRust(original)).length, 2);
  writeFileSync(file, rust.replace(threadedRust(original), threadedRust(current)));
  state.runtime = { original_sha256: hash(original), candidate_sha256: hash(current) };
  state.generated = manifest(join(out, 'rust'), path => /\.(rs|toml)$/.test(path));
  const next = new Set(state.generated.map(file => file.path));
  for (const file of manifest(join(work, 'rust'), path => /\.(rs|toml)$/.test(path)))
    if (!next.has(file.path)) rmSync(join(work, 'rust', file.path));
  for (const file of state.generated) {
    const destination = join(work, 'rust', file.path);
    if (!existsSync(destination) || hash(readFileSync(destination)) !== file.sha256) copy(join(out, 'rust', file.path), destination);
  }
  save();
  command('aff-contracts', 'cargo', ['test', '--release', '--config', 'profile.release.lto=false',
    '--config', 'profile.release.debug=false', '--target-dir', join(archive, 'runtime-tests-target'),
    '--manifest-path', join(work, 'rust/Cargo.toml'), '-p', 'Purs_Effect_Aff', '--lib']);
  const cargo = cargoProfile({ lto: 'thin', linker: 'rust-lld' }, env, join(work, 'rust'), join(work, 'target'));
  Object.assign(state, { profile: cargo.profile, linker: cargo.linker, rustflags: cargo.rustflags });
  command('cargo', 'cargo', cargo.args, work, cargo.env);
  copy(join(work, 'target/release/purust_output'), join(out, 'gopurs-rust')); chmodSync(join(out, 'gopurs-rust'), 0o755);
  state.binary_sha256 = hash(readFileSync(join(out, 'gopurs-rust')));
  assert.deepEqual(manifest(join(work, 'rust'), path => /\.(rs|toml)$/.test(path)), state.generated);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  compare(label, [
    { name: 'go', directory: 'baseline/compiler', environment: {} },
    { name: 'control', binary: 'candidates/control/gopurs-rust' },
    { name: label, binary: `candidates/${label}/gopurs-rust` },
  ]);
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
