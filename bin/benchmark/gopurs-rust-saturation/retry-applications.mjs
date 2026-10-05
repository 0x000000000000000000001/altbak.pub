// Complete a measured production campaign after the application-cwd correction.
// The failed attempt and both successful measurement campaigns stay untouched.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, manifest } from './common.mjs';
import { auditRuns } from './audit-runs.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const read = path => JSON.parse(readFileSync(join(archive, path)));
const evidence = path => ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) });
const preflightPath = 'production/application-recovery-1/preflight.json', preflight = read(preflightPath);
assert.equal(preflight.status, 'passed'); assert.equal(preflight.frozen_projects_verified, 51);
for (const file of [preflight.previous, preflight.validator_before, preflight.failed_applications,
  preflight.artifact, ...preflight.measurements]) assert.equal(hash(readFileSync(file.path)), file.sha256);
const state = read('production-campaign.json'); assert.equal(state.status, 'failed');
assert.equal(hash(readFileSync(join(archive, 'production-campaign.json'))), preflight.previous.sha256);
assert.equal(read('production/results.json').status, 'passed');
assert.equal(read('production/contracts/results.json').status, 'passed');
const processes = execFileSync('ps', ['-axo', 'pid,ppid,state,etime,time,command'], { encoding: 'utf8' });
assert(!processes.split('\n').some(line => /gopurs-rust-saturation\/(?:production|resume-production|measure|validate)\.mjs/.test(line)),
  'Another campaign process is still present');
const verifyCompiler = () => {
  assert.deepEqual(manifest(state.compiler.directory), state.compiler.files);
  for (const file of state.compiler.files) assert.equal(hash(readFileSync(join(state.compiler.origin, file.path))), file.sha256);
};
verifyCompiler();
copy(fileURLToPath(import.meta.url), join(archive, 'production/application-recovery-1/retry-applications.mjs'));
state.application_recovery = { status: 'running', started_at: new Date().toISOString(),
  reason: preflight.reason, preflight: evidence(preflightPath), failed_applications: preflight.failed_applications,
  previous: preflight.previous, preserved_measurements: preflight.measurements };
state.status = 'running'; delete state.failure;
const save = () => writeJson(join(archive, 'production-campaign.json'), state); save();
try {
  state.commands.push(run(archive, 'production-applications-retry-1', process.execPath,
    [join(here, 'validate.mjs'), archive, 'applications-retry-1'], archive, environment(), 3600000)); save();
  const resultPath = 'production/applications-retry-1/results.json', applications = read(resultPath);
  assert.equal(applications.status, 'passed'); assert.equal(applications.applications.length, 9);
  assert(applications.applications.every(item => item.status === 'passed'));
  state.audits = ['production-full-table', 'production-published-paired'].map(label => auditRuns(archive, label));
  for (const file of preflight.measurements) assert.equal(hash(readFileSync(file.path)), file.sha256);
  verifyCompiler();
  state.applications_result = resultPath;
  state.application_recovery.result = evidence(resultPath);
  state.application_recovery.status = 'passed'; state.application_recovery.finished_at = new Date().toISOString();
  state.recovery.status = 'passed'; state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, applications: applications.applications.map(item =>
    ({ name: item.name, status: item.status, checks: item.checks, spec_results: item.spec_results, execution: item.execution })),
    audits: state.audits }, null, 2));
} catch (error) {
  state.application_recovery.status = 'failed'; state.status = 'failed'; state.failure = error.stack; save(); throw error;
}
