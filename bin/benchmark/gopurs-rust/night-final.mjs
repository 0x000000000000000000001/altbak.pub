// The declared overnight confirmation protocol, run only after production gates.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, resume] = process.argv.slice(2);
assert(archiveArg && (!resume || resume === '--resume'), 'night-final.mjs ARCHIVE [--resume]');
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../../gopurs/gopurs');
const snapshot = resolve(archive, '../gopurs-purust-aff-20261002');
const path = join(archive, 'final-campaign.json');
const production = JSON.parse(readFileSync(join(archive, 'production/results.json')));
assert.equal(production.status, 'passed');
const bootstrap = JSON.parse(readFileSync(join(archive, 'purust-bootstrap-final/results.json')));
assert.equal(bootstrap.status, 'passed');
const hosts = JSON.parse(readFileSync(join(archive, 'production/hosts/results.json')));
assert.equal(hosts.status, 'passed');
assert.equal(hosts.bootstrap.purust_native_sha256, bootstrap.executables['purust-native']);
const executables = Object.fromEntries(['gopurs', 'gopurs.js', 'gopurs-native', 'gopurs-rust']
  .map(name => [name, hash(readFileSync(join(root, 'bin', name)))]));
for (const [name, binary] of Object.entries(production.executables)) assert.equal(executables[name], binary.sha256,
  'Installed host changed after qualification: ' + name);
const state = resume ? JSON.parse(readFileSync(path))
  : { status: 'pending', started_at: new Date().toISOString(), executables, commands: [] };
assert.deepEqual(state.executables, executables);
if (resume) { assert.equal(state.status, 'failed'); state.resumed_at = new Date().toISOString(); }
else assert(!existsSync(path));
const save = () => writeJson(path, state);
const command = (label, executable, args) => {
  console.log(label);
  state.commands.push(run(archive, label, executable, args, root, environment(), 3600000)); save();
};
save();
try {
  for (const [name, rounds, defaults, resources, variants] of [
    ['primary', 30, false, false, ['go', 'rust']],
    ['default', 10, true, false, ['go', 'rust']],
    ['common', 10, false, false, ['js', 'go', 'rust']],
    ['resources', 3, false, true, ['go', 'rust']],
    ['before-after', 10, false, false, ['rust-before', 'rust']],
  ]) {
    const directory = join(archive, name + '-runs');
    if (existsSync(directory)) {
      assert(resume, directory);
      assert.equal(JSON.parse(readFileSync(join(directory, 'results.json'))).status, 'passed',
        'Retain failed diagnostics and choose a fresh campaign directory before resuming');
      continue;
    }
    const selection = { rounds, launcherDefaults: defaults, resources,
      variants: variants.map(name => ({ name, directory: name === 'rust-before' ? join(archive, 'baseline/compiler') : root,
        environment: name.startsWith('rust') ? { GOPURS_RUST: '1' } : name === 'js' ? { GOPURS_JS: '1' } : {} })) };
    writeJson(join(archive, name + '.json'), selection);
    command('activity-before-' + name, 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
    command('compare-' + name, process.execPath,
      [join(here, 'compare.mjs'), snapshot, directory, join(archive, name + '.json')]);
    command('activity-after-' + name, 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
  }
  for (const name of ['primary', 'default', 'common']) command('paired-' + name, process.execPath,
    [join(here, 'paired-summary.mjs'), join(archive, name + '-runs/results.json'), 'go', 'rust',
      join(archive, name + '-paired.json')]);
  command('paired-before-after', process.execPath, [join(here, 'paired-summary.mjs'),
    join(archive, 'before-after-runs/results.json'), 'rust-before', 'rust', join(archive, 'before-after-paired.json')]);
  const campaigns = ['primary', 'default', 'common', 'resources', 'before-after']
    .map(name => join(archive, name + '-runs'));
  command('verify-night', process.execPath, [join(here, 'verify.mjs'), join(archive, 'verification-final.json'), ...campaigns]);
  assert.deepEqual(Object.fromEntries(Object.keys(executables).map(name =>
    [name, hash(readFileSync(join(root, 'bin', name)))])), executables);
  const primary = JSON.parse(readFileSync(join(archive, 'primary-paired.json')));
  const defaults = JSON.parse(readFileSync(join(archive, 'default-paired.json')));
  state.victory = primary.nightly_30_pair_gate && defaults.median_candidate_ms < defaults.median_reference_ms
    && defaults.bootstrap.ratio_interval[1] < 1;
  state.status = 'passed'; state.finished_at = new Date().toISOString(); delete state.failure; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
