// Recover an interrupted final campaign without splicing a project's rounds.
// Completed projects retain their raw records; unfinished projects restart with
// a new warmup and five complete rounds under the exact same frozen compilers.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, emitted, hash, manifest } from './common.mjs';
import { auditRuns } from './audit-runs.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const read = path => JSON.parse(readFileSync(join(archive, path)));
const evidence = path => ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) });
const state = read('production-campaign.json'), original = read('runs/production-full-table/results.json');
assert.equal(state.status, 'running'); assert.equal(original.status, 'running');
const processes = execFileSync('ps', ['-axo', 'pid,ppid,state,etime,time,command'], { encoding: 'utf8' });
assert(!processes.split('\n').some(line => /gopurs-rust-saturation\/(?:production|measure)\.mjs/.test(line)),
  'A campaign process is still present; do not launch a concurrent recovery');
const directory = join(archive, 'production/recovery-1'); assert(!existsSync(directory)); mkdirSync(directory);
copy(join(archive, 'production-campaign.json'), join(directory, 'production-before.json'));
copy(join(archive, 'runs/production-full-table/results.json'), join(directory, 'full-table-before.json'));
copy(fileURLToPath(import.meta.url), join(directory, 'resume-production.mjs'));
const completed = original.results.filter(item => item.status === 'passed');
const done = new Set(completed.map(item => item.name));
const campaign = read('campaign.json'), remaining = campaign.cases.filter(item => !done.has(item.name)).map(item => item.name);
assert(remaining.length > 0 && completed.length > 0);
assert.equal(read('production/results.json').status, 'passed');
assert.equal(read('production/contracts/results.json').status, 'passed');
assert.deepEqual(manifest(state.compiler.directory), state.compiler.files);
for (const file of state.compiler.files)
  assert.equal(hash(readFileSync(join(state.compiler.origin, file.path))), file.sha256);
state.recovery = { status: 'running', detected_at: new Date().toISOString(),
  reason: 'No campaign process remains; the persisted measurement stopped during a project without a final notification. Cause not established.',
  previous: evidence('production/recovery-1/production-before.json'),
  interrupted_measurement: evidence('production/recovery-1/full-table-before.json'),
  retained_projects: completed.map(item => item.name), restarted_projects: remaining,
  superseded_checked_generations: original.results.filter(item => !done.has(item.name)).reduce((sum, item) =>
    sum + item.runs.filter(run => run.identical_files > 0).length, 0), partial_outputs: [] };
for (const item of original.results.filter(item => !done.has(item.name))) {
  const definition = read(`cases/${item.name}/definition.json`), files = emitted(definition.input);
  for (const file of files) copy(join(definition.input, 'output', file.path), join(directory, 'partial-output', item.name, file.path));
  state.recovery.partial_outputs.push({ name: item.name, files });
}
const save = () => writeJson(join(archive, 'production-campaign.json'), state); save();
const command = (label, script, args) => {
  state.commands.push(run(archive, label, process.execPath, [join(here, script), ...args], archive, environment(), 10800000)); save();
};
try {
  const configuration = join(directory, 'remaining.json');
  writeJson(configuration, { ...original.configuration, cases: remaining });
  command('production-full-table-restart-1', 'measure.mjs', [archive, 'production-full-table-restart-1', configuration]);
  const resumed = read('runs/production-full-table-restart-1/results.json');
  assert.equal(resumed.status, 'passed'); assert.deepEqual(resumed.variants, original.variants);
  assert.equal(resumed.campaign_sha256, original.campaign_sha256);
  const byName = new Map([...completed, ...resumed.results].map(item => [item.name, item]));
  const results = campaign.cases.map(item => byName.get(item.name));
  assert.equal(results.length, 51); assert(results.every(item => item?.status === 'passed'));
  const merged = { ...original, status: 'passed', finished_at: resumed.finished_at, results,
    harness: [...original.harness, ...resumed.harness],
    recovery: { original: state.recovery.interrupted_measurement,
      restarted: evidence('runs/production-full-table-restart-1/results.json'), retained_projects: [...done], restarted_projects: remaining },
    total_ms: Object.fromEntries(original.variants.map(variant => [variant.name,
      results.reduce((sum, item) => sum + item.summary[variant.name].median_ms, 0)])) };
  writeJson(join(archive, 'runs/production-full-table/results.json'), merged);
  state.recovery.table_audit = auditRuns(archive, 'production-full-table'); save();
  const paired = join(archive, 'production/published-paired.json'); assert(!existsSync(paired));
  writeJson(paired, { rounds: 5, variants: [{ name: 'published-rust', binary: 'baseline/compiler/bin/gopurs-rust' },
    original.configuration.variants.find(variant => variant.name === 'rust')] });
  command('production-published-paired', 'measure.mjs', [archive, 'production-published-paired', paired]);
  command('production-applications', 'validate.mjs', [archive]);
  state.audits = ['production-full-table', 'production-published-paired'].map(label => auditRuns(archive, label));
  assert.deepEqual(manifest(state.compiler.directory), state.compiler.files);
  for (const file of state.compiler.files)
    assert.equal(hash(readFileSync(join(state.compiler.origin, file.path))), file.sha256);
  state.recovery.status = 'passed'; state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, recovery: state.recovery, audits: state.audits }, null, 2));
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
