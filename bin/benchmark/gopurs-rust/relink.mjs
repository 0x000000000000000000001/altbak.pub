// Isolate a Rust build-profile change from the frozen generated source.
import assert from 'node:assert/strict';
import { chmodSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
import { cargoProfile, readProfile } from './cargo-profile.mjs';

const [archiveArg, sourceLabel, label, lto, linkerArg] = process.argv.slice(2);
assert(archiveArg && [sourceLabel, label].every(x => /^[a-z0-9-]+$/.test(x)) && lto
  && (!linkerArg || linkerArg === '--rust-lld'),
  'relink.mjs ARCHIVE SOURCE_CANDIDATE NEW_CANDIDATE thin|false|PROFILE.json [--rust-lld]');
const archive = resolve(archiveArg), source = join(archive, 'candidates', sourceLabel);
const out = join(archive, 'candidates', label), work = join(archive, 'work');
assert(!existsSync(out)); mkdirSync(out);
const original = JSON.parse(readFileSync(join(source, 'build.json')));
assert.equal(original.status, 'passed');
assert.deepEqual(manifest(join(work, 'rust'), p => /\.(rs|toml)$/.test(p)), original.generated);
const result = { status: 'pending', label, source_candidate: source, started_at: new Date().toISOString(),
  sources: original.sources, generated: original.generated, tast: original.tast,
  commands: [] };
const save = () => writeJson(join(out, 'build.json'), result);
const env = { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true' };
const configuration = readProfile(lto);
if (linkerArg) configuration.linker = 'rust-lld';
const cargo = cargoProfile(configuration, env, join(work, 'rust'), join(work, 'target'));
Object.assign(result, { profile: cargo.profile, linker: cargo.linker, rustflags: cargo.rustflags });
save();
try {
  copyFileSync(new URL(import.meta.url), join(out, 'relink.mjs'));
  copyFileSync(new URL('./cargo-profile.mjs', import.meta.url), join(out, 'cargo-profile.mjs'));
  for (const name of ['sources', 'rust']) {
    result.commands.push(run(out, 'freeze-' + name, 'cp', ['-cR', join(source, name), join(out, name)], archive, env)); save();
  }
  result.commands.push(run(out, 'cargo', 'cargo', cargo.args, work, cargo.env, 3600000)); save();
  const binary = join(out, 'gopurs-rust');
  copyFileSync(join(work, 'target/release/purust_output'), binary, constants.COPYFILE_FICLONE);
  chmodSync(binary, 0o755);
  assert.deepEqual(manifest(join(work, 'rust'), p => /\.(rs|toml)$/.test(p)), original.generated);
  result.binary_sha256 = hash(readFileSync(binary)); result.status = 'passed';
  result.finished_at = new Date().toISOString(); save();
  console.log(`${label}: ${result.binary_sha256}`);
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
