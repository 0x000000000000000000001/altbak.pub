// Reuse the exact frozen frontend corpus while moving to one corrected compiler cohort.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), previous = resolve(process.argv[3]);
const campaignPath = join(archive, 'campaign.json'), campaign = JSON.parse(readFileSync(campaignPath));
const groups = ['libraries-final-results.json', 'b8x-results.json'].map(name => {
  const path = join(previous, name), result = JSON.parse(readFileSync(path)); assert.equal(result.status, 'passed');
  return { path, sha256: hash(readFileSync(path)), result };
});
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const changes = [];
for (const row of groups.flatMap(group => group.result.results)) {
  const old = JSON.parse(readFileSync(row.definition)), directory = join(archive, 'cases', row.name), input = join(directory, 'input');
  assert(!existsSync(directory));
  for (const file of old.input_manifest) {
    assert.equal(hash(readFileSync(join(old.input, file.path))), file.sha256); copy(join(old.input, file.path), join(input, file.path));
  }
  assert.deepEqual(manifest(input), old.input_manifest);
  const primed = old.input_manifest.filter(file => file.path.endsWith('.purs') &&
    /^foreign\s+import\s+data\s+[A-Z][A-Za-z0-9_']*'[A-Za-z0-9_']*\s*(?:::|∷)/m.test(readFileSync(join(input, file.path), 'utf8'))).map(file => file.path);
  const definition = { ...old, input,
    source_corpus: row.name === 'b8x' ? old.source_corpus : join(archive, 'family', row.name),
    historical_oracle: primed.length ? undefined : { directory: row.runs[0].retained_output, files: manifest(row.runs[0].retained_output) },
    prior_definition: { path: row.definition, sha256: row.definition_sha256 },
    correction: primed.length ? { reason: 'Preserve complete primed foreign identifiers and their native FFI carrier', sources: primed } : null };
  writeJson(join(directory, 'definition.json'), definition);
  if (primed.length) changes.push({ name: row.name, sources: primed });
}
campaign.rerun = { previous_archive: previous, evidence: groups.map(({ path, sha256 }) => ({ path, sha256 })),
  reason: 'A native HTTP application exposed a truncated primed foreign identifier; all projects are re-measured with the corrected qualified hosts.',
  changed_output_cases: changes };
writeJson(campaignPath, campaign);
console.log(JSON.stringify({ frozen_cases: 58, changed_output_cases: changes }, null, 2));
