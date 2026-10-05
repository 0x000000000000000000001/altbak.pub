// Fresh self-trained production PGO, three-host qualification and full tables.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, manifest } from './common.mjs';
import { auditRuns } from './audit-runs.mjs';

const [archiveArg, candidate] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(candidate));
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../../gopurs/gopurs');
const path = join(archive, 'production-campaign.json'); assert(!existsSync(path));
const state = { status: 'running', started_at: new Date().toISOString(), candidate, commands: [] };
const save = () => writeJson(path, state);
const command = (label, script, args) => {
  state.commands.push(run(archive, label, process.execPath, [join(here, script), ...args], archive, environment(), 10800000)); save();
};
save();
try {
  copy(resolve(here, '../../../README.md'), join(archive, 'README-before-production.md'));
  const selection = JSON.parse(readFileSync(join(archive, 'selection.json')));
  assert.equal(selection.status, 'passed'); assert.equal(selection.candidate, candidate);
  command('production-qualification', '../gopurs-rust/qualify.mjs', [archive, candidate]);
  const qualificationPath = join(archive, 'production/results.json');
  const qualification = JSON.parse(readFileSync(qualificationPath)); assert.equal(qualification.status, 'passed');
  command('production-contracts', 'production-contracts.mjs', [archive]);
  const destination = join(archive, 'production/compiler'); assert(!existsSync(destination));
  for (const file of ['bin/gopurs', 'bin/gopurs.js', 'bin/gopurs-native', 'bin/gopurs-rust', 'package.json',
    'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js']) copy(join(root, file), join(destination, file));
  state.compiler = { directory: destination, origin: root, files: manifest(destination),
    qualification: { path: qualificationPath, sha256: hash(readFileSync(qualificationPath)) } }; save();
  const variants = ['js', 'go', 'rust'].map(name => ({ name, binary: 'production/compiler/bin/gopurs',
    compiler: 'production/compiler/bin/' + ({ js: 'gopurs.js', go: 'gopurs-native', rust: 'gopurs-rust' })[name],
    environment: { GOPURS_JS: name === 'js' ? '1' : '0', GOPURS_RUST: name === 'rust' ? '1' : '0',
      ...(name === 'go' ? { GOGC: 'off', GOMEMLIMIT: '10GiB' } : {}) } }));
  const full = join(archive, 'production/full-table.json'); writeJson(full, { rounds: 5, variants });
  command('production-full-table', 'measure.mjs', [archive, 'production-full-table', full]);
  const paired = join(archive, 'production/published-paired.json'); writeJson(paired, { rounds: 5,
    variants: [{ name: 'published-rust', binary: 'baseline/compiler/bin/gopurs-rust' }, variants[2]] });
  command('production-published-paired', 'measure.mjs', [archive, 'production-published-paired', paired]);
  command('production-applications', 'validate.mjs', [archive]);
  state.audits = ['production-full-table', 'production-published-paired'].map(label => auditRuns(archive, label));
  assert.deepEqual(manifest(destination), state.compiler.files);
  for (const file of state.compiler.files) assert.equal(hash(readFileSync(join(root, file.path))), file.sha256);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
