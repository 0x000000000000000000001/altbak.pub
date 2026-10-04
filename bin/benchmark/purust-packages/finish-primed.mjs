// Serialize the final foreign-handle correction, uniform rerun and publication candidate.
import assert from 'node:assert/strict';
import { cpSync, existsSync, readFileSync, watch } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
const attempt = process.argv[3] ?? 'finish4', start = process.argv[4] ?? 'qualify-primed';
assert.match(attempt, /^[a-zA-Z0-9-]+$/);
const path = join(archive, attempt + '-results.json'); assert(!existsSync(path));
const bootstrap = process.argv[5] ?? 'primed-bootstrap'; assert.match(bootstrap, /^[a-zA-Z0-9-]+$/);
const resumeWorkspace = process.argv[6] && process.argv[6] !== '-' ? [resolve(process.argv[6])] : [];
const revisionLabel = process.argv[7] ?? 'revision2'; assert.match(revisionLabel, /^[a-zA-Z0-9-]+$/);
const revision = join(archive, revisionLabel), qualification = join(archive, bootstrap, 'qualification.json');
const state = { status: 'running', started_at: new Date().toISOString(), revision, commands: [] };
const snapshot = join(archive, attempt + '-tool-snapshot');
cpSync(tools, join(snapshot, 'purust-packages'), { recursive: true });
for (const name of ['compilation-refresh/measure.mjs', 'gopurs-aff/common.mjs']) cpSync(join(tools, '..', name), join(snapshot, name));
state.tools = manifest(snapshot);
const save = () => writeJson(path, state); save();
// Filesystem notification, rather than polling or overlapping the prior Cargo campaign.
const previous = join(archive, 'finish3-results.json');
if (JSON.parse(readFileSync(previous)).status === 'running') {
  state.stage = 'waiting-for-previous-validation'; save();
  await new Promise((done, fail) => {
    const watcher = watch(archive, (_event, name) => { if (name === 'finish3-results.json') check(); });
    watcher.on('error', fail);
    function check() {
      let result;
      try { result = JSON.parse(readFileSync(previous)); } catch { return; }
      if (result.status !== 'running') { watcher.close(); done(); }
    }
    check();
  });
}
const steps = [
  ['qualify-primed', join(tools, 'qualify-primed.mjs'), archive, bootstrap, ...resumeWorkspace],
  ['prepare-revision', join(tools, 'prepare.mjs'), revision, archive, qualification],
  ['reuse-frontend', join(tools, 'prepare-rerun.mjs'), revision, join(archive, 'revision1')],
  ['measure-libraries', join(tools, '../compilation-refresh/measure.mjs'), revision, 'libraries'],
  ['measure-b8x', join(tools, '../compilation-refresh/measure.mjs'), revision, 'b8x', join(revision, 'cases/b8x/definition.json')],
  ['validate-libraries', join(tools, 'reuse-validation.mjs'), revision],
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
