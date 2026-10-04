// Re-measure only cases whose isolated retry lost a public runner flag.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const tools = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
const replacements = new Map(process.argv.slice(3).map(path => {
  path = resolve(path); return [JSON.parse(readFileSync(path)).name, path];
}));
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const originalPath = join(campaign.baseline_archive, 'campaign.json'), original = JSON.parse(readFileSync(originalPath));
const sourcePath = join(archive, 'libraries-results.json'), previous = JSON.parse(readFileSync(sourcePath));
assert.notEqual(previous.status, 'running');
const path = join(archive, 'libraries-final-results.json'); assert(!existsSync(path));
mkdirSync(join(archive, 'recovery'));
const state = { status: 'running', started_at: new Date().toISOString(),
  previous: { path: sourcePath, sha256: hash(readFileSync(sourcePath)) },
  original_plans: { path: originalPath, sha256: hash(readFileSync(originalPath)) },
  results: structuredClone(previous.results), corrections: [] };
const save = () => writeJson(path, state); save();
for (let index = 0; index < state.results.length; index++) {
  const result = state.results[index], previousPath = join(archive, 'cases', result.name, 'definition.json');
  const replacement = replacements.get(result.name), definitionPath = replacement ?? previousPath;
  if (!existsSync(definitionPath)) continue;
  const definition = JSON.parse(readFileSync(definitionPath)), plan = original.plans.find(plan => plan.name === result.name); assert(plan);
  const invocation = ['--source', 'output', '--main', plan.main, ...plan.flags];
  if (!replacement && JSON.stringify(definition.invocation) === JSON.stringify(invocation)) continue;
  let nextPath = replacement;
  if (replacement) assert.deepEqual(definition.invocation, invocation);
  else {
    const directory = join(archive, 'cases', result.name + '-flags-retry'), input = join(directory, 'input');
    assert(!existsSync(directory)); mkdirSync(input, { recursive: true });
    for (const file of definition.input_manifest) {
      const source = join(definition.input, file.path), target = join(input, file.path);
      assert.equal(hash(readFileSync(source)), file.sha256);
      mkdirSync(dirname(target), { recursive: true }); copyFileSync(source, target, constants.COPYFILE_FICLONE);
    }
    const next = { ...definition, input, input_manifest: manifest(input), invocation,
      invocation_provenance: { path: originalPath, sha256: hash(readFileSync(originalPath)), main: plan.main, flags: plan.flags } };
    nextPath = join(directory, 'definition.json'); writeJson(nextPath, next);
  }
  const correction = { name: result.name, reason: replacement ? 'Use a separately retained repaired frontend definition'
    : 'Restore the original public runner flags lost when reusing backend-free workspace YAML',
    previous_definition: existsSync(previousPath) ? previousPath : null,
    previous_definition_sha256: existsSync(previousPath) ? hash(readFileSync(previousPath)) : null,
    previous_invocation: existsSync(previousPath) ? JSON.parse(readFileSync(previousPath)).invocation : null,
    invocation, definition: nextPath, status: 'running' };
  state.corrections.push(correction); save();
  try {
    run(archive, `recover-${result.name}`, process.execPath,
      [join(tools, '../compilation-refresh/measure.mjs'), archive, `recovery/${result.name}`, nextPath], tools, environment(), 1200000);
    const measured = JSON.parse(readFileSync(join(archive, 'recovery', result.name + '-results.json')));
    assert.equal(measured.status, 'passed'); state.results[index] = measured.results[0]; correction.status = 'passed';
    console.log(`${result.name}: corrected case PASS`);
  } catch (error) {
    correction.status = 'failed'; correction.error = error.stack;
    state.results[index] = { ...result, status: 'failed', error: error.stack };
    console.error(`${result.name}: ${error.message}`);
  }
  save();
}
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'partial';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, corrected: state.corrections.map(item => item.name),
  passed: state.results.filter(item => item.status === 'passed').length }, null, 2));
if (state.status !== 'passed') process.exitCode = 1;
