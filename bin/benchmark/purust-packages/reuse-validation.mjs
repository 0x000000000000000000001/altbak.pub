// Reuse application evidence only when all generated source bytes still match.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const measured = JSON.parse(readFileSync(join(archive, 'libraries-results.json'))); assert.equal(measured.status, 'passed');
const qualification = JSON.parse(readFileSync(campaign.qualification.path));
const paths = [join(campaign.rerun.previous_archive, 'validation/results.json'), qualification.http_validation.path];
const prior = paths.map(path => ({ path, sha256: hash(readFileSync(path)), ...JSON.parse(readFileSync(path)) }));
assert(prior.every(group => group.status !== 'running'));
const root = join(archive, 'validation'); assert(!existsSync(root)); mkdirSync(root);
const state = { status: 'running', started_at: new Date().toISOString(), results: [], reused: [], commands: [] };
const save = () => writeJson(join(root, 'results.json'), state); save();
const fresh = [];
for (const row of measured.results) {
  const files = manifest(row.runs[0].retained_output);
  const group = prior.find(group => group.results.some(item => item.name === row.name && item.status === 'passed' &&
    JSON.stringify(item.generated_manifest) === JSON.stringify(files)));
  if (!group) { fresh.push(row.name); continue; }
  const item = group.results.find(item => item.name === row.name);
  assert.equal(hash(readFileSync(item.binary.path)), item.binary.sha256);
  const reuse = { path: group.path, sha256: group.sha256, name: item.name, reason: 'Byte-identical canonical Rust/Cargo source' };
  state.results.push({ ...item, definition_sha256: row.definition_sha256, reused_from: reuse });
  state.reused.push(reuse); save();
}
try {
  if (fresh.length) {
    state.commands.push(run(archive, 'validate-changed-applications', process.execPath,
      [join(tools, 'validate.mjs'), archive, 'libraries-results.json', 'validation-new', ...fresh], tools, environment(), 7200000));
    const result = JSON.parse(readFileSync(join(archive, 'validation-new/results.json'))); assert.equal(result.status, 'passed');
    state.results.push(...result.results);
  }
  state.results.sort((a, b) => a.name.localeCompare(b.name));
  assert.equal(state.results.length, 57); assert.equal(new Set(state.results.map(item => item.name)).size, 57);
  assert(state.results.every(item => item.status === 'passed'));
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ passed: 57, reused: state.reused.length, fresh }, null, 2));
} catch (error) {
  state.status = 'failed'; state.error = error.stack; state.finished_at = new Date().toISOString(); save(); throw error;
}
