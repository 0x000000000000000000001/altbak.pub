// Reuse the immutable preparation binary after correcting the JS error fixture.
import assert from 'node:assert/strict';
import { readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const candidate = join(archive, 'candidates/parallel-rewrite');
const build = JSON.parse(readFileSync(join(candidate, 'build.json')));
assert.equal(build.status, 'passed');
assert.equal(hash(readFileSync(join(candidate, 'gopurs-rust'))), build.binary_sha256);
const state = { status: 'pending', candidate, binary_sha256: build.binary_sha256, commands: [] };
const save = () => writeJson(join(archive, 'preparation-check-v2.json'), state);
const command = (name, args) => {
  state.commands.push(run(archive, name, process.execPath, args, archive, environment(), 300000)); save();
};
save();
try {
  const source = resolve(here, '../../../../gopurs/gopurs/tools/monomorphization.test.mjs');
  const script = join(candidate, 'preparation-v2.test.mjs');
  writeFileSync(script, readFileSync(source, 'utf8').replaceAll('../output/', join(archive, 'work/output') + '/'));
  command('preparation-contracts-v2', ['--test', '--test-timeout=30000', script]);
  const selection = join(archive, 'parallel-rewrite-selection.json');
  writeJson(selection, { rounds: 5, variants: JSON.parse(readFileSync(join(archive, 'parallel-rewrite-config.json'))).variants });
  command('preparation-compare-v2', [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, 'parallel-rewrite-runs'), selection]);
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
