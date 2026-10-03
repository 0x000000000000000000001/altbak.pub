// Build a source candidate, check PBO contracts, then compare serially with control.
import assert from 'node:assert/strict';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, label] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label), 'night-candidate.mjs ARCHIVE LABEL');
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const state = { status: 'pending', label, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(archive, label + '-campaign.json'), state);
const command = (name, args) => {
  console.log(name);
  state.commands.push(run(archive, label + '-' + name, process.execPath, args, archive, environment(), 3600000)); save();
};
save();
try {
  command('build', [join(here, 'experiment.mjs'), 'build', archive, label, join(archive, 'profile-baseline.json')]);
  command('pbo', [join(here, 'check-pbo.mjs'), join(archive, 'work/output'), join(archive, label + '-pbo')]);
  const selection = join(archive, label + '-selection.json');
  writeJson(selection, { rounds: 5, variants: [
    { name: 'control', binary: 'candidates/control/gopurs-rust' },
    { name: label, binary: `candidates/${label}/gopurs-rust` },
  ] });
  command('compare', [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, label + '-runs'), selection]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
