import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { hash, manifest, writeJson } from './common.mjs';

const workspace = resolve(process.argv[2]);
const result = JSON.parse(readFileSync(join(workspace, 'results.json'), 'utf8'));
assert(result.finished_at && !result.failure);
assert.equal(result.runs.length, 24);
let frozenFiles = 0;
for (const [directory, entries] of Object.entries(result.frozen_files)) for (const file of entries) {
  assert.equal(hash(readFileSync(join(workspace, directory, file.path))), file.sha256, file.path); frozenFiles++;
}
const references = new Map(result.reference_generations.map(run => [run.family,
  JSON.parse(readFileSync(run.generated_manifest, 'utf8'))]));
let generatedFiles = 0;
for (const run of [...result.reference_generations, ...result.runs]) {
  const files = manifest(run.output, path => /\.(rs|go|toml)$/.test(path) || path.endsWith('/go.mod'));
  assert.deepEqual(files, references.get(run.family), run.label);
  assert.equal(hash(JSON.stringify(files)), run.generated_sha256, run.label);
  generatedFiles += files.length;
}
const median = values => values.toSorted((a, b) => a - b)[Math.floor(values.length / 2)];
for (const [variant, summary] of Object.entries(result.summary)) {
  const samples = result.runs.filter(run => run.variant === variant && !run.warmup).map(run => run.phases_ms['backend total']);
  assert.equal(samples.length, 5); assert.deepEqual(samples, summary.samples_ms); assert.equal(median(samples), summary.median_ms);
}
const expected = readFileSync(new URL('./expected.stdout', import.meta.url), 'utf8').trimEnd().split('\n').sort();
for (const run of [...result.application_validation.original_executions, ...result.application_validation.portable_executions]) {
  assert.equal(hash(readFileSync(run.application)), run.application_sha256, run.application);
  if (run.passed) {
    assert.equal(run.exit_code, 0); assert.equal(readFileSync(run.stderr, 'utf8'), '');
    assert.deepEqual(readFileSync(run.stdout, 'utf8').trimEnd().split('\n').sort(), expected);
  }
}
const verification = { at: new Date().toISOString(), status: 'passed', frozen_files: frozenFiles,
  generated_outputs: result.runs.length + result.reference_generations.length, generated_files: generatedFiles,
  go_files_per_output: references.get('gopurs').length, rust_files_per_output: references.get('purust').length,
  original_suite_go_passes: result.application_validation.original_executions.filter(run => run.index === 0 && run.passed).length,
  original_suite_rust_passes: result.application_validation.original_executions.filter(run => run.index === 1 && run.passed).length,
  portable_suite_passes: result.application_validation.portable_executions.filter(run => run.passed).length };
writeJson(join(workspace, 'verification.json'), verification);
console.log(JSON.stringify(verification, null, 2));
