// Reuse the exact published frontend inputs and Go oracles in writable copies.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, jobs, manifest, verifyInput } from './common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..');
const archive = resolve(process.argv[2]);
const historical = join(site, 'var/benchmark/compilation-refresh-20261004');
const publicationPath = join(historical, 'publication.json');
const publication = JSON.parse(readFileSync(publicationPath));
const output = join(archive, 'campaign.json');
assert(!existsSync(output));
mkdirSync(archive, { recursive: true });
const campaign = { status: 'preparing', started_at: new Date().toISOString(),
  historical: { publication: publicationPath, sha256: hash(readFileSync(publicationPath)) },
  protocol: { metric: 'backend total', jobs, warmups: 1, screening_rounds: 5,
    practical_threshold: { aggregate: 0.01, important_case: 0.03 },
    exact_output: 'Every generated Go file and go.mod compared byte-for-byte',
    stop: 'After all priority mechanisms and a final profile review, three distinct mechanisms without a retained gain' },
  cases: [] };
writeJson(output, campaign);
for (const result of publication.results.filter(result => result.family === 'gopurs')) {
  const original = JSON.parse(readFileSync(result.definition));
  assert.equal(hash(readFileSync(result.definition)), result.definition_sha256);
  const directory = join(archive, 'cases', result.name), input = join(directory, 'input');
  assert(!existsSync(directory)); mkdirSync(directory, { recursive: true });
  for (const file of original.input_manifest) {
    const source = join(original.input, file.path);
    assert.equal(hash(readFileSync(source)), file.sha256, source);
    copy(source, join(input, file.path));
  }
  for (const file of original.sibling_inputs ?? []) {
    const source = join(dirname(original.input), file.path);
    assert.equal(hash(readFileSync(source)), file.sha256, source);
    const destination = join(directory, file.path);
    assert(destination.startsWith(directory + '/'), destination);
    copy(source, destination);
  }
  const canonical = join(dirname(result.definition), 'canonical-generated');
  const files = manifest(canonical);
  assert.equal(hash(JSON.stringify(files)), result.generated_sha256);
  const oracle = join(directory, 'oracle');
  for (const file of files) copy(join(canonical, file.path), join(oracle, file.path));
  const item = { name: result.name, input, input_manifest: original.input_manifest,
    sibling_inputs: original.sibling_inputs ?? [], invocation: original.invocation,
    oracle: { directory: oracle, files }, original_definition: result.definition,
    original_definition_sha256: result.definition_sha256, historical_summary: result.summary,
    scope: original.scope, input_counts: result.input };
  verifyInput(item);
  const definition = join(directory, 'definition.json'); writeJson(definition, item);
  campaign.cases.push({ name: item.name, definition, sha256: hash(readFileSync(definition)) });
  writeJson(output, campaign);
}
assert.equal(campaign.cases.length, 51);
campaign.status = 'passed'; campaign.finished_at = new Date().toISOString();
writeJson(output, campaign);
writeJson(join(archive, 'profile-baseline.json'), { lto: 'thin', linker: 'rust-lld' });
console.log(`Frozen ${campaign.cases.length} published compilation cases in ${archive}`);
