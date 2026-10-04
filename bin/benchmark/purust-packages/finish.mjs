// Continue a completed inventory with one serialized correction/publication pipeline.
import assert from 'node:assert/strict';
import { cpSync, existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
assert.notEqual(JSON.parse(readFileSync(join(archive, 'libraries-results.json'))).status, 'running');
const attempt = process.argv[3] ?? 'finish', bootstrap = attempt === 'finish' ? 'bootstrap' : attempt + '-bootstrap';
const start = process.argv[4] ?? 'qualify-fix';
const replacements = process.argv.slice(5).map(path => resolve(path));
assert.match(attempt, /^[a-zA-Z0-9-]+$/);
const path = join(archive, attempt + '-results.json'); assert(!existsSync(path));
const revision = join(archive, 'revision1');
const state = { status: 'running', started_at: new Date().toISOString(), revision, commands: [] };
const snapshot = join(archive, attempt + '-tool-snapshot');
cpSync(tools, join(snapshot, 'purust-packages'), { recursive: true });
for (const name of ['compilation-refresh/measure.mjs', 'gopurs-aff/common.mjs']) {
  cpSync(join(tools, '..', name), join(snapshot, name));
}
state.tools = manifest(snapshot);
const save = () => writeJson(path, state); save();
const steps = [
  ['qualify-fix', join(tools, 'qualify-fix.mjs'), archive, bootstrap],
  ['prepare-revision', join(tools, 'prepare.mjs'), revision, archive, join(archive, bootstrap, 'qualification.json')],
  ['libraries-revision', join(tools, 'libraries.mjs'), revision],
  ['recover-flags', join(tools, 'recover-flags.mjs'), revision, ...replacements],
  ['prepare-b8x', join(tools, 'prepare-b8x.mjs'), revision],
  ['measure-b8x', join(tools, '../compilation-refresh/measure.mjs'), revision, 'b8x', join(revision, 'cases/b8x/definition.json')],
  ['validate-libraries', join(tools, 'validate.mjs'), revision, 'libraries-final-results.json'],
  ['validate-b8x', join(tools, 'validate-b8x.mjs'), revision],
  ['publish-candidate', join(tools, 'publish.mjs'), revision],
];
const first = steps.findIndex(([label]) => label === start); assert(first >= 0, start);
state.start_stage = start; save();
try {
  for (const [label, ...args] of steps.slice(first)) {
    state.stage = label; save(); console.log(`Starting ${label}`);
    state.commands.push(run(archive, attempt + '-' + label, process.execPath, args, tools, environment(), 7200000));
    save(); console.log(`${label}: passed`);
  }
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(`Qualified README candidate: ${join(revision, 'README-next.md')}`);
} catch (error) {
  state.status = 'failed'; state.error = error.stack; state.finished_at = new Date().toISOString(); save(); throw error;
}
