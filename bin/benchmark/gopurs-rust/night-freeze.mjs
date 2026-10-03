// Adopt the freshly rebuilt, retained production compiler as an immutable control.
import assert from 'node:assert/strict';
import { chmodSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

assert(process.argv[2], 'night-freeze.mjs ARCHIVE');
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const bootstrap = JSON.parse(readFileSync(join(archive, 'bootstrap.json')));
assert.equal(bootstrap.status, 'passed');
const env = environment();
run(archive, 'freeze-baseline', process.execPath, [join(here, 'experiment.mjs'), 'init', archive,
  bootstrap.executables['gopurs-rust'].sha256, join(archive, 'profile-baseline.json')], archive, env);
const baseline = JSON.parse(readFileSync(join(archive, 'baseline.json')));
const out = join(archive, 'candidates/control'), work = join(archive, 'work');
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const state = { status: 'pending', label: 'control', sources: baseline.sources,
  origin: bootstrap.rust_bootstrap, commands: [], profile: {
    opt_level: 3, debug: false, lto: 'thin', threaded: true, allocator: 'mimalloc' } };
const save = () => writeJson(join(out, 'build.json'), state);
save();
try {
  for (const [label, source, destination] of [
    ['sources', join(archive, 'baseline/sources'), join(out, 'sources')],
    ['generated', join(bootstrap.rust_bootstrap, 'rust'), join(out, 'rust')],
    ['work-rust', join(out, 'rust'), join(work, 'rust')],
    ['work-target', join(bootstrap.rust_bootstrap, 'target'), join(work, 'target')],
    ['work-output', join(bootstrap.rust_bootstrap, 'output'), join(work, 'output')],
    ['purust-sources', resolve(here, '../../../../purust/purust/src'), join(archive, 'purust-sources')],
  ]) state.commands.push(run(out, label, 'cp', ['-cRp', source, destination], archive, env, 3600000));
  const buildLog = readFileSync(bootstrap.commands.find(r => r.label === 'rust').stdout, 'utf8');
  const command = buildLog.match(/^\[cargo-build\] (cargo .+)$/m)[1].split(' ');
  state.commands.push({ label: 'cargo', command, exit_code: 0, original_log: bootstrap.commands.find(r => r.label === 'rust').stdout });
  const linker = command.find(arg => arg.startsWith('link-arg=-fuse-ld=')).slice('link-arg=-fuse-ld='.length);
  state.linker = { path: linker, sha256: hash(readFileSync(linker)), driver: 'cc', option: '-fuse-ld=' + linker };
  state.generated = manifest(join(out, 'rust'), path => /\.(rs|toml)$/.test(path));
  state.tast = JSON.parse(readFileSync(join(bootstrap.rust_bootstrap, 'verify-tast.json')));
  copyFileSync(bootstrap.executables['gopurs-rust'].path, join(out, 'gopurs-rust'), constants.COPYFILE_FICLONE);
  chmodSync(join(out, 'gopurs-rust'), 0o755);
  state.binary_sha256 = hash(readFileSync(join(out, 'gopurs-rust')));
  assert.equal(state.binary_sha256, baseline.rust_sha256);
  assert.deepEqual(manifest(join(work, 'rust'), path => /\.(rs|toml)$/.test(path)), state.generated);
  state.status = 'passed'; save();
  console.log(JSON.stringify({ status: state.status, binary: state.binary_sha256, files: state.generated.length }, null, 2));
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
