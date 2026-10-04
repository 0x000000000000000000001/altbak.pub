// PGO trained on gopurs's own compiler modules, never on the held-out Aff test.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, rmSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, label, attempt = label] = process.argv.slice(2);
assert(archiveArg && [label, attempt].every(name => /^[a-z0-9-]+$/.test(name)),
  'night-pgo.mjs ARCHIVE SOURCE_CANDIDATE [ATTEMPT]');
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const source = join(archive, 'candidates', label), directory = join(archive, 'pgo-' + attempt), work = join(archive, 'work');
assert(!existsSync(directory)); mkdirSync(directory);
const env = { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true' };
const state = { status: 'pending', source, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(directory, 'results.json'), state);
const command = (name, executable, args, cwd = archive, extra = {}) => {
  console.log(name); const record = run(directory, name, executable, args, cwd, { ...env, ...extra }, 3600000);
  state.commands.push(record); save(); return record;
};
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
save();
try {
  const build = JSON.parse(readFileSync(join(source, 'build.json'))); assert.equal(build.status, 'passed');
  const next = new Set(build.generated.map(file => file.path));
  for (const file of manifest(join(work, 'rust'), path => /\.(rs|toml)$/.test(path)))
    if (!next.has(file.path)) rmSync(join(work, 'rust', file.path));
  for (const file of build.generated) {
    const path = join(work, 'rust', file.path);
    if (!existsSync(path) || hash(readFileSync(path)) !== file.sha256) copy(join(source, 'rust', file.path), path);
  }
  const sysroot = execFileSync('rustc', ['--print', 'sysroot'], { env, encoding: 'utf8' }).trim();
  const host = execFileSync('rustc', ['-vV'], { env, encoding: 'utf8' }).match(/^host: (.+)$/m)[1];
  const profdata = join(sysroot, 'lib/rustlib', host, 'bin/llvm-profdata'); assert(existsSync(profdata));
  state.profdata_tool = { path: profdata, sha256: hash(readFileSync(profdata)) };
  command('profdata-version', profdata, ['--version']);
  const training = join(directory, 'training'), output = join(training, 'output'); mkdirSync(output, { recursive: true });
  const bootstrap = JSON.parse(readFileSync(join(archive, 'bootstrap.json'))).go_bootstrap;
  const originalInputs = [];
  for (const entry of readdirSync(join(bootstrap, 'output'), { withFileTypes: true })) {
    const path = join(bootstrap, 'output', entry.name, 'corefn.json');
    if (!entry.isDirectory() || !existsSync(path)) continue;
    assert.notEqual(entry.name, 'Test.Main', 'held-out Aff test module must not be in training');
    const bytes = readFileSync(path), json = JSON.parse(bytes), original = resolve(bootstrap, json.modulePath);
    originalInputs.push({ path, sha256: hash(bytes), source: original });
    const sourcePath = `sources/${entry.name}/${entry.name.split('.').at(-1)}.purs`;
    for (const extension of ['purs', 'go', 'js']) {
      const sibling = original.replace(/\.purs$/, '.' + extension);
      if (existsSync(sibling)) copy(sibling, join(training, sourcePath.replace(/\.purs$/, '.' + extension)));
    }
    json.modulePath = sourcePath;
    mkdirSync(join(output, entry.name)); writeJson(join(output, entry.name, 'corefn.json'), json);
  }
  state.training = { origin: bootstrap, modules: originalInputs.length, original_inputs: originalInputs,
    frozen: manifest(training), purpose: 'gopurs compiler self-compilation, excludes Test.Main; shared library modules are allowed; rewritten paths point only to frozen sibling sources' };
  const heldOut = JSON.parse(readFileSync(resolve(archive, '../gopurs-purust-aff-20261002/results.json')));
  const heldOutNames = new Set(heldOut.frozen_files.inputs.filter(file => file.path.endsWith('/corefn.json'))
    .map(file => file.path.split('/').at(-2)));
  state.training.overlap_modules = readdirSync(output).filter(name => heldOutNames.has(name)).sort();
  state.training.held_out_modules = [...heldOutNames].sort();
  assert(!state.training.overlap_modules.includes('Test.Main'));
  assert(state.training.modules > 300); save();
  const clean = () => {
    // On case-insensitive filesystems output/main aliases the Main TAST input.
    // Delete only generated Go files, preserving every frozen corefn.json.
    for (const path of walk(output)) if (/\.go$|\/go\.(mod|sum)$/.test(path)) rmSync(path);
    for (const name of ['.cache', '.purmeta']) rmSync(join(training, name), { recursive: true, force: true });
    assert.deepEqual(manifest(training), state.training.frozen);
    assert(existsSync(join(output, 'Main/corefn.json')));
  };
  const emitted = () => manifest(output, path => path.endsWith('.go') || path === 'go.mod');
  const workers = { GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' };
  clean(); command('training-oracle', join(source, 'gopurs-rust'), ['--main', 'Main'], training, workers);
  state.training.oracle = emitted(); assert(state.training.oracle.length > 400); save();
  const raw = join(directory, 'raw'); mkdirSync(raw);
  const generateProfile = join(directory, 'generate-profile.json');
  writeJson(generateProfile, { lto: 'thin', linker: 'rust-lld', rustflags: ['-C', 'profile-generate=' + raw] });
  command('instrument', process.execPath, [join(here, 'relink.mjs'), archive, label, attempt + '-pgo-instrumented', generateProfile]);
  for (let round = 1; round <= 3; round++) {
    clean(); command('train-' + round, join(archive, 'candidates', attempt + '-pgo-instrumented/gopurs-rust'), ['--main', 'Main'], training,
      { ...workers, ...(round === 3 ? { GOPURS_PREPARE_JOBS: '1', GOPURS_PBO_JOBS: '1', GOPURS_EMIT_JOBS: '1' } : {}),
        LLVM_PROFILE_FILE: join(raw, '%m-%p.profraw') });
    assert.deepEqual(emitted(), state.training.oracle, 'training generation differs under instrumentation');
  }
  state.profiles = manifest(raw); assert(state.profiles.length > 0); save();
  const merged = join(directory, 'training.profdata');
  command('merge', profdata, ['merge', '-o', merged, ...state.profiles.map(file => join(raw, file.path))]);
  state.merged_sha256 = hash(readFileSync(merged)); save();
  const useProfile = join(directory, 'use-profile.json');
  writeJson(useProfile, { lto: 'thin', linker: 'rust-lld', rustflags: ['-C', 'profile-use=' + merged] });
  command('profile-use', process.execPath, [join(here, 'relink.mjs'), archive, label, attempt + '-pgo', useProfile]);
  const selection = join(archive, attempt + '-pgo-selection.json');
  writeJson(selection, { rounds: 5, variants: [
    { name: label, binary: `candidates/${label}/gopurs-rust` },
    { name: attempt + '-pgo', binary: `candidates/${attempt}-pgo/gopurs-rust` },
    { name: 'go', directory: 'baseline/compiler', environment: {} },
  ] });
  command('held-out-compare', process.execPath, [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, attempt + '-pgo-runs'), selection]);
  clean();
  assert.deepEqual(manifest(training), state.training.frozen);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
