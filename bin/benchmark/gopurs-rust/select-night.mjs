// Bind the source composition and qualified profile experiment before production.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), candidate = 'selected-source-v2-pgo';
const path = join(archive, 'final-selection.json'); assert(!existsSync(path));
const records = {};
const load = relative => {
  const path = join(archive, relative), bytes = readFileSync(path);
  records[relative] = hash(bytes); return JSON.parse(bytes);
};
assert.equal(load('selected-source-campaign.json').status, 'passed');
assert.equal(load('selected-native-contracts/provenance.json').status, 'passed');
const pgo = load('pgo-selected-source-v2/results.json'); assert.equal(pgo.status, 'passed');
assert(pgo.commands.every(record => record.exit_code === 0));
assert.equal(pgo.commands.filter(record => /^train-/.test(record.label)).length, 3);
assert(!pgo.training.overlap_modules.includes('Test.Main'));
assert(pgo.training.frozen.some(file => file.path === 'output/Main/corefn.json'));
assert.deepEqual(manifest(join(archive, 'pgo-selected-source-v2/training')), pgo.training.frozen);
assert.equal(hash(readFileSync(join(archive, 'pgo-selected-source-v2/training.profdata'))), pgo.merged_sha256);
const results = load(candidate + '-runs/results.json'); assert.equal(results.status, 'passed');
const summary = results.summary, reference = 'selected-source';
let favorable = 0;
for (let round = 1; round <= results.protocol.rounds; round++) {
  const time = variant => {
    const runs = results.runs.filter(run => run.round === round && run.variant === variant);
    assert.equal(runs.length, 1); const run = runs[0]; assert.equal(run.identical_files, 294);
    const raw = Number(readFileSync(run.stderr, 'utf8').match(/^\[gopurs\] backend total: (\d+) ms$/m)?.[1]);
    assert.equal(raw, run.phases_ms['backend total']); return raw;
  };
  if (time(candidate) < time(reference)) favorable++;
}
assert.equal(results.protocol.rounds, 5); assert.equal(favorable, 5);
assert(summary[candidate].median_ms < summary[reference].median_ms * 0.95);
const build = load('candidates/' + candidate + '/build.json'); assert.equal(build.status, 'passed');
const source = load('candidates/selected-source/build.json'); assert.equal(source.status, 'passed');
assert.deepEqual(build.generated, source.generated); assert.deepEqual(build.sources, source.sources);
assert.equal(hash(readFileSync(join(archive, 'candidates', candidate, 'gopurs-rust'))), build.binary_sha256);
const result = { status: 'passed', selected_at: new Date().toISOString(), candidate,
  source_candidate: 'selected-source', source_campaign: 'selected-source-campaign.json',
  reason: 'Full source composition passes differential gates. Complete independent compiler-self PGO training passes integrity/output checks; five held-out pairs favor PGO by more than five percent in median.',
  binary_sha256: build.binary_sha256, profile: { ...build.profile, pgo: 'compiler-self-training' },
  measurement: { reference, candidate, reference_median_ms: summary[reference].median_ms,
    candidate_median_ms: summary[candidate].median_ms, favorable_pairs: favorable, pairs: 5,
    change_percent: 100 * (summary[candidate].median_ms / summary[reference].median_ms - 1) },
  training: { modules: pgo.training.modules, shared_library_modules: pgo.training.overlap_modules,
    excluded_entrypoint: 'Test.Main', included_entrypoint: 'Main',
    merged_sha256: pgo.merged_sha256,
    note: '500 frozen compiler-self modules; shared library modules allowed. Profile archive also retains instrumented Cargo build-script/proc-macro process profiles. No held-out Aff invocation used the instrumented host.' },
  production_policy: 'Rebuild with the same frozen compiler sources, O3/ThinLTO and self-contained PGO training on current compiler-self inputs; independently qualify and measure the resulting installed binaries against freshly rebuilt Go.',
  evidence_sha256: records };
writeJson(path, result); console.log(JSON.stringify(result.measurement, null, 2));
