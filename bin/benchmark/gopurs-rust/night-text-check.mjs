// Resume the text decoder's contract gate after a fixture-only correction.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), revision = process.argv[3] ?? 'v2';
assert(/^[a-z0-9-]+$/.test(revision));
const here = dirname(fileURLToPath(import.meta.url));
const candidate = join(archive, 'candidates/native-tast-text');
const build = JSON.parse(readFileSync(join(candidate, 'build.json')));
assert.equal(build.status, 'passed');
assert.equal(hash(readFileSync(join(candidate, 'gopurs-rust'))), build.binary_sha256);
const state = { status: 'pending', candidate, binary_sha256: build.binary_sha256, commands: [] };
const save = () => writeJson(join(archive, `text-check-${revision}.json`), state);
const command = (name, args) => {
  state.commands.push(run(archive, name, process.execPath, args, archive,
    { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', PBO_NATIVE_KEEP: '1', TMPDIR: archive }, 3600000)); save();
};
save();
try {
  command(`text-contracts-${revision}`, [join(archive, 'text-proposal/tests/run-text-differential.mjs'),
    '--workspace', join(candidate, 'rust'), '--corpus', resolve(archive, '../gopurs-purust-aff-20261002/inputs/gopurs-aff/output'),
    '--corpus-name', 'gopurs-aff-238', '--log', join(archive, `text-contracts-${revision}`),
    '--target-dir', join(archive, 'native-contracts-target')]);
  const selection = join(archive, 'native-tast-text-selection.json');
  writeJson(selection, { rounds: 5, variants: JSON.parse(readFileSync(join(archive, 'native-tast-text-config.json'))).variants });
  command(`text-compare-${revision}`, [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, 'native-tast-text-runs'), selection]);
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
