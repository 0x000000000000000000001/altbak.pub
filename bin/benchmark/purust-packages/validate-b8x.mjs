// Build the output-identical b8x native application outside the backend timer.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), root = join(archive, 'b8x-validation');
assert(!existsSync(root)); mkdirSync(root);
const results = JSON.parse(readFileSync(join(archive, 'b8x-results.json'))); assert.equal(results.status, 'passed');
const result = results.results[0]; assert.equal(result.name, 'b8x');
const canonical = join(dirname(result.definition), 'canonical-generated'), generated = join(root, 'generated');
const files = manifest(canonical), copy = (from, to) => {
  mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE);
};
const state = { status: 'running', started_at: new Date().toISOString(), definition_sha256: hash(readFileSync(result.definition)),
  generated_manifest: files, scope: 'Build the native import closure and the Test.Rust.Main application; service-dependent execution belongs to the separate b8x runtime qualification.', commands: [] };
const save = () => writeJson(join(root, 'results.json'), state); save();
try {
  for (const file of files) copy(join(canonical, file.path), join(generated, file.path));
  const env = { ...environment(), CARGO_TARGET_DIR: join(archive, 'validation-cargo-target'), CARGO_BUILD_JOBS: '8',
    CARGO_PROFILE_DEV_DEBUG: '0', CARGO_PROFILE_DEV_INCREMENTAL: 'false', RUSTFLAGS: '-Awarnings' };
  state.commands.push(run(root, 'build', 'cargo', ['build', '--manifest-path', join(generated, 'Cargo.toml'),
    '-p', 'purust_output', '--bin', 'purust_output'], generated, env, 1800000)); save();
  const binary = join(root, 'b8x-tests'); copy(join(env.CARGO_TARGET_DIR, 'debug/purust_output'), binary);
  state.binary = { path: binary, sha256: hash(readFileSync(binary)) };
  assert.deepEqual(manifest(canonical), files);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(`b8x: native application builds; retained ${binary}`);
} catch (error) {
  state.status = 'failed'; state.error = error.stack; save(); throw error;
}
