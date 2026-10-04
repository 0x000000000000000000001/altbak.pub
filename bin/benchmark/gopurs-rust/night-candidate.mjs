// Build a source candidate, check PBO contracts, then compare serially with control.
import assert from 'node:assert/strict';
import { constants, cpSync, existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, label, configurationArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label), 'night-candidate.mjs ARCHIVE LABEL [CONFIGURATION.json]');
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const configuration = configurationArg ? JSON.parse(readFileSync(resolve(configurationArg))) : {};
// Isolated follow-ups must not inherit runtime edits made after the reference
// build. Recover every library from that candidate's immutable source snapshot.
if (configuration.frozenRuntime) {
  const frozen = join(archive, 'candidates', configuration.frozenRuntime);
  const build = JSON.parse(readFileSync(join(frozen, 'build.json')));
  assert.equal(build.status, 'passed');
  const libraryNames = new Map([...readFileSync(join(frozen, 'spago.yaml'), 'utf8')
    .matchAll(/^    ([^:\n]+):\n      path: (.+)$/gm)]
    .map(([, name, path]) => [JSON.parse(path), name]));
  const runtime = join(archive, label + '-runtime-inputs');
  assert(!existsSync(runtime)); mkdirSync(runtime);
  const overrides = Object.fromEntries(build.runtime_sources.map((source, index) => {
    const origin = join(frozen, 'runtime/sources', String(index));
    for (const file of source.files) assert.equal(hash(readFileSync(join(origin, file.path))), file.sha256);
    // Build tools may create package-local metadata. Keep it outside the frozen
    // reference, with a disposable writable copy for this candidate only.
    const directory = join(runtime, String(index));
    cpSync(origin, directory, { recursive: true, mode: constants.COPYFILE_FICLONE });
    // Workspace aliases are authoritative: a few ports deliberately retain a
    // `purust-` prefix in package.name despite an unprefixed workspace alias.
    const name = libraryNames.get(source.path);
    assert(name, 'Missing frozen workspace alias for ' + source.path);
    assert(/^[a-z][a-z0-9-]*$/.test(name));
    return [resolve(here, '../../../../purust/purust-' + name), directory];
  }));
  configuration.environment = { ...configuration.environment,
    GOPURS_EXPERIMENT_RUNTIME_OVERRIDES: JSON.stringify({ ...overrides,
      ...JSON.parse(configuration.environment?.GOPURS_EXPERIMENT_RUNTIME_OVERRIDES ?? '{}') }) };
}
const state = { status: 'pending', label, configuration, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(archive, label + '-campaign.json'), state);
const command = (name, args, executable = process.execPath, env = {}) => {
  console.log(name);
  state.commands.push(run(archive, label + '-' + name, executable, args, archive, { ...environment(), ...env }, 3600000)); save();
};
save();
try {
  command('build', [join(here, 'experiment.mjs'), 'build', archive, label, join(archive, 'profile-baseline.json')], process.execPath,
    configuration.environment ?? {});
  if (configuration.pbo !== false)
    command('pbo', [join(here, 'check-pbo.mjs'), join(archive, 'work/output'), join(archive, label + '-pbo')]);
  if (configuration.affTests) command('aff-contracts', ['test', '--release', '--config', 'profile.release.lto=false',
    '--config', 'profile.release.debug=false', '--target-dir', join(archive, 'runtime-tests-target'),
    '--manifest-path', join(archive, 'work/rust/Cargo.toml'), '-p', 'Purs_Effect_Aff', '--lib'], 'cargo',
    { CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_NET_OFFLINE: 'true' });
  if (configuration.preparationTests) {
    const source = resolve(here, '../../../../gopurs/gopurs/tools/monomorphization.test.mjs');
    const script = join(archive, 'candidates', label, 'preparation.test.mjs');
    writeFileSync(script, readFileSync(source, 'utf8').replaceAll('../output/', join(archive, 'work/output') + '/'));
    command('preparation-contracts', ['--test', '--test-timeout=30000', script]);
  }
  for (const check of configuration.checks ?? [])
    command(check.label, check.args, check.command ?? process.execPath, check.environment ?? {});
  const selection = join(archive, label + '-selection.json');
  writeJson(selection, { rounds: configuration.rounds ?? 5, variants: configuration.variants ?? [
    { name: 'control', binary: 'candidates/control/gopurs-rust' },
    { name: label, binary: `candidates/${label}/gopurs-rust` },
  ] });
  command('compare', [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, label + '-runs'), selection]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
