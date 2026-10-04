// Serialize the independent bootstrap, public production builds and confirmation.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
const [archiveArg, candidate = 'selected-source'] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(candidate));
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const destination = join(archive, 'production-pipeline.json'); assert(!existsSync(destination));
const state = { status: 'pending', candidate, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(destination, state);
const command = (label, script, args) => {
  console.log(label);
  state.commands.push(run(archive, label, process.execPath, [join(here, script), ...args],
    archive, environment(), 10800000)); save();
};
save();
try {
  const selection = JSON.parse(readFileSync(join(archive, 'final-selection.json')));
  assert.equal(selection.status, 'passed'); assert.equal(selection.candidate, candidate);
  for (const path of [selection.source_campaign, 'selected-native-contracts/provenance.json', 'native-runtime-rest-v1/results.json',
    'native-aff-synchronized-v1/results.json'])
    assert.equal(JSON.parse(readFileSync(join(archive, path))).status, 'passed', path);
  const bootstrap = join(archive, 'purust-bootstrap-final/results.json');
  if (existsSync(bootstrap)) {
    assert.equal(JSON.parse(readFileSync(bootstrap)).status, 'passed');
    state.independent_bootstrap = bootstrap; save();
  } else command('night-independent-bootstrap', 'night-purust-qualification.mjs', [archive]);
  command('night-production-qualification', 'qualify.mjs', [archive, candidate]);
  command('night-final-confirmation', 'night-final.mjs', [archive]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
