// Serialize frontend preparation and all backend measurements; retain every failed attempt.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
if (!existsSync(join(archive, 'campaign.json'))) {
  const prepared = spawnSync(process.execPath, [join(tools, 'prepare.mjs'), archive], { stdio: 'inherit' });
  assert.ifError(prepared.error); assert.equal(prepared.status, 0);
}
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const path = join(archive, 'libraries-results.json'); assert(!existsSync(path));
mkdirSync(join(archive, 'measurements'), { recursive: true });
const state = { status: 'running', started_at: new Date().toISOString(), results: [] };
const save = () => writeJson(path, state); save();
for (const plan of campaign.plans) {
  const result = { name: plan.name, status: 'running' }; state.results.push(result); save();
  try {
    run(archive, `prepare-${plan.name}`, process.execPath, [join(tools, 'prepare-case.mjs'), archive, plan.name], tools, environment(), 900000);
    run(archive, `measure-${plan.name}`, process.execPath, [join(tools, '../compilation-refresh/measure.mjs'), archive,
      `measurements/${plan.name}`, join(archive, 'cases', plan.name, 'definition.json')], tools, environment(), 900000);
    const measured = JSON.parse(readFileSync(join(archive, 'measurements', plan.name + '-results.json')));
    assert.equal(measured.status, 'passed'); assert.equal(measured.results.length, 1);
    Object.assign(result, measured.results[0]);
    console.log(`${plan.name}: PASS JS ${result.summary.js.median_ms} / Rust ${result.summary.rust.median_ms} ms`);
  } catch (error) {
    result.status = 'failed'; result.error = error.stack; console.error(`${plan.name}: FAILED ${error.message}`);
  }
  save();
}
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'partial';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, passed: state.results.filter(result => result.status === 'passed').length,
  failed: state.results.filter(result => result.status !== 'passed').map(result => result.name) }, null, 2));
if (state.status !== 'passed') process.exitCode = 1;
