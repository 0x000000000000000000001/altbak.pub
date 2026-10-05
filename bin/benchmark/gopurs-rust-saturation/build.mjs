// Screening keeps the published training profile fixed for both controls and
// candidates. The selected composition must subsequently retrain production PGO.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync, symlinkSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash } from './common.mjs';

const [archiveArg, label, sourcesArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && sourcesArg);
const archive = resolve(archiveArg), sources = resolve(sourcesArg), here = dirname(fileURLToPath(import.meta.url));
const site = resolve(here, '../../..');
const nodeModules = resolve(here, '../../../../gopurs/gopurs/node_modules');
// Differential JS fixtures resolve esbuild from the compiled output's parent.
// Spago alone does not create this link in the private native build workspace.
mkdirSync(join(archive, 'work'), { recursive: true });
if (!existsSync(join(archive, 'work/node_modules'))) symlinkSync(nodeModules, join(archive, 'work/node_modules'), 'dir');
const old = join(site, 'var/benchmark/gopurs-packages-fixes-20261004/revision2/gopurs-rust-build-yJznhT');
const pgo = JSON.parse(readFileSync(join(old, 'pgo-profile.json')));
assert.equal(pgo.status, 'passed');
assert.equal(pgo.binary.sha256, JSON.parse(readFileSync(join(archive, 'baseline.json'))).rust_sha256);
const profile = join(archive, 'screening.profdata');
if (!existsSync(profile)) copy(join(old, pgo.profile.merged.path), profile);
assert.equal(hash(readFileSync(profile)), pgo.profile.merged.sha256);
const configuration = join(archive, 'profile-screening.json');
if (!existsSync(configuration)) {
  writeJson(configuration, { lto: 'thin', linker: 'rust-lld', rustflags: ['-C', 'profile-use=' + profile] });
  writeJson(join(archive, 'screening-profile-provenance.json'), { source: join(old, 'pgo-profile.json'),
    metadata_sha256: hash(readFileSync(join(old, 'pgo-profile.json'))), profile_sha256: pgo.profile.merged.sha256,
    policy: 'Fixed historical compiler-self profile for screening; independent freshly retrained PGO required for production' });
}
const output = join(archive, label + '-build-results.json'); assert(!existsSync(output));
const state = { status: 'running', label, sources, started_at: new Date().toISOString(), commands: [],
  test_dependencies: { node_modules: nodeModules, esbuild_package_sha256: hash(readFileSync(join(nodeModules, 'esbuild/package.json'))) } };
const save = () => writeJson(output, state); save();
try {
  state.commands.push(run(archive, label + '-build', process.execPath,
    [join(here, '../gopurs-rust/experiment.mjs'), 'build', archive, label, configuration], archive,
    { ...environment(), GOPURS_EXPERIMENT_SOURCE_ROOTS: JSON.stringify([join(sources, '0'), join(sources, '1')]) }, 3600000)); save();
  state.commands.push(run(archive, label + '-pbo-checks', process.execPath,
    [join(here, '../gopurs-rust/check-pbo.mjs'), join(archive, 'work/output'), join(archive, 'checks', label + '-pbo')], archive,
    environment(), 1800000));
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
