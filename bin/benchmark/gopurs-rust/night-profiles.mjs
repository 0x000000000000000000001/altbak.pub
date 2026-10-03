// Serialized calibration and build-profile selection on frozen generated Rust.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

assert(process.argv[2], 'night-profiles.mjs ARCHIVE');
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const snapshot = resolve(archive, '../gopurs-purust-aff-20261002');
assert.equal(JSON.parse(readFileSync(join(archive, 'candidates/control/build.json'))).status, 'passed');
const state = { status: 'pending', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(archive, 'profile-campaign.json'), state);
const command = (name, executable, args) => {
  console.log(name);
  state.commands.push(run(archive, name, executable, args, archive, environment(), 3600000)); save();
};
save();
try {
  writeJson(join(archive, 'baseline-comparison.json'), { rounds: 5, variants: [
    { name: 'go', directory: 'baseline/compiler', environment: {} },
    { name: 'rust', directory: 'baseline/compiler', environment: { GOPURS_RUST: '1' } },
  ] });
  command('activity-before-baseline', 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
  command('baseline-comparison', process.execPath, [join(here, 'compare.mjs'), snapshot,
    join(archive, 'baseline-runs'), join(archive, 'baseline-comparison.json')]);
  for (const label of ['cgu1', 'cgu4', 'native-cgu1']) {
    command('build-' + label, process.execPath, [join(here, 'relink.mjs'), archive, 'control', label,
      join(archive, 'profile-' + label + '.json')]);
  }
  writeJson(join(archive, 'profiles.json'), { rounds: 5, variants: ['control', 'cgu1', 'cgu4', 'native-cgu1']
    .map(name => ({ name, binary: `candidates/${name}/gopurs-rust` })) });
  command('activity-before-profiles', 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
  command('profiles-comparison', process.execPath, [join(here, 'compare.mjs'), snapshot,
    join(archive, 'profiles-runs'), join(archive, 'profiles.json')]);
  command('activity-after-profiles', 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
