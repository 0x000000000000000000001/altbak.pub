// Isolate a Rust build-profile change from the frozen generated source.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { chmodSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, sourceLabel, label, lto, linkerArg] = process.argv.slice(2);
assert(archiveArg && [sourceLabel, label].every(x => /^[a-z0-9-]+$/.test(x)) && ['thin', 'false'].includes(lto)
  && (!linkerArg || linkerArg === '--rust-lld'),
  'relink.mjs ARCHIVE SOURCE_CANDIDATE NEW_CANDIDATE thin|false [--rust-lld]');
const archive = resolve(archiveArg), source = join(archive, 'candidates', sourceLabel);
const out = join(archive, 'candidates', label), work = join(archive, 'work');
assert(!existsSync(out)); mkdirSync(out);
const original = JSON.parse(readFileSync(join(source, 'build.json')));
assert.equal(original.status, 'passed');
assert.deepEqual(manifest(join(work, 'rust'), p => /\.(rs|toml)$/.test(p)), original.generated);
const result = { status: 'pending', label, source_candidate: source, started_at: new Date().toISOString(),
  sources: original.sources, generated: original.generated, tast: original.tast,
  profile: { opt_level: 3, debug: false, lto, threaded: true, allocator: 'mimalloc' }, commands: [] };
const save = () => writeJson(join(out, 'build.json'), result);
const env = { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true' };
let linker;
if (linkerArg) {
  assert.equal(process.platform, 'darwin');
  const sysroot = execFileSync('rustc', ['--print', 'sysroot'], { env, encoding: 'utf8' }).trim();
  const host = execFileSync('rustc', ['-vV'], { env, encoding: 'utf8' }).match(/^host: (.+)$/m)[1];
  linker = join(sysroot, 'lib/rustlib', host, 'bin/gcc-ld/ld64.lld');
  assert(existsSync(linker));
  result.linker = { path: linker, sha256: hash(readFileSync(linker)), driver: 'cc', option: '-fuse-ld=' + linker };
}
save();
try {
  copyFileSync(new URL(import.meta.url), join(out, 'relink.mjs'));
  for (const name of ['sources', 'rust']) {
    result.commands.push(run(out, 'freeze-' + name, 'cp', ['-cR', join(source, name), join(out, name)], archive, env)); save();
  }
  result.commands.push(run(out, 'cargo', 'cargo', [linker ? 'rustc' : 'build', '--release', '--target-dir', join(work, 'target'),
    '--config', `profile.release.lto=${lto === 'thin' ? '"thin"' : 'false'}`,
    '--config', 'profile.release.opt-level=3', '--config', 'profile.release.debug=false',
    '--manifest-path', join(work, 'rust/Cargo.toml'),
    ...(linker ? ['--bin', 'purust_output', '--', '-C', 'link-arg=-fuse-ld=' + linker] : [])], work, env, 3600000)); save();
  const binary = join(out, 'gopurs-rust');
  copyFileSync(join(work, 'target/release/purust_output'), binary, constants.COPYFILE_FICLONE);
  chmodSync(binary, 0o755);
  assert.deepEqual(manifest(join(work, 'rust'), p => /\.(rs|toml)$/.test(p)), original.generated);
  result.binary_sha256 = hash(readFileSync(binary)); result.status = 'passed';
  result.finished_at = new Date().toISOString(); save();
  console.log(`${label}: ${result.binary_sha256}`);
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
