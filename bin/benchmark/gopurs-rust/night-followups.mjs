// Finish the isolated gates serially, retaining their earlier failed diagnostics.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const recordFile = join(archive, 'isolated-followups.json'); assert(!existsSync(recordFile));
const read = path => JSON.parse(readFileSync(join(archive, path)));
const state = { status: 'pending', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(recordFile, state);
const command = (label, args, overrides = {}) => {
  console.log(label);
  state.commands.push(run(archive, label, process.execPath, args, archive,
    { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', TMPDIR: archive, ...overrides }, 3600000)); save();
};
save();
try {
  const build = read('candidates/transitive-rounds/build.json');
  assert.equal(build.status, 'passed');
  assert.equal(hash(readFileSync(join(archive, 'candidates/transitive-rounds/gopurs-rust'))), build.binary_sha256);
  assert.equal(read('transitive-rounds-pbo/results.json').status, 'passed');
  assert.equal(read('transitive-proposal/evidence/oracle-record-fixed.json').status, 'passed');
  const selection = join(archive, 'transitive-rounds-selection.json');
  writeJson(selection, { rounds: 5, variants: read('transitive-proposal/transitive-config.json').variants });
  command('transitive-qualified-compare', [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, 'transitive-rounds-runs'), selection]);
  command('collections-explicit-oracles', [resolve(here, '../../../../purust/purust-ordered-collections/tools/test-native-maps.mjs'),
    join(archive, 'candidates/native-collections/rust')],
    { PURUST_NATIVE_KEEP: '1', PURUST_NATIVE_TARGET: join(archive, 'map-contracts-target') });
  command('text-v6-campaign', [join(here, 'night-candidate.mjs'), archive, 'native-tast-text-v6',
    join(archive, 'native-tast-text-v6-config.json')]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
