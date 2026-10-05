// Record the reviewed selection and practical stopping decision from actual
// completed experiments. This campaign-specific decision is intentionally explicit.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { writeJson } from '../gopurs-aff/common.mjs';
import { hash, manifest, pairedAggregate } from './common.mjs';
import { auditRuns } from './audit-runs.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const candidate = 'composition-v3', read = path => JSON.parse(readFileSync(join(archive, path)));
const evidence = path => ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) });
assert(!existsSync(join(archive, 'selection.json'))); assert(!existsSync(join(archive, 'closure.json')));
const build = read(`candidates/${candidate}/build.json`), proposal = read(`proposals/${candidate}/results.json`);
assert.equal(build.status, 'passed'); assert.equal(proposal.status, 'passed');
assert.equal(proposal.runtime_layout.identical, true);
assert.equal(hash(readFileSync(join(archive, 'candidates', candidate, 'gopurs-rust'))), build.binary_sha256);
const roots = [resolve(here, '../../../../gopurs/gopurs'), resolve(here, '../../../../purescript-backend-optimizer-gopurs')];
const sources = roots.map((root, index) => {
  const files = manifest(join(root, 'src'));
  assert.deepEqual(files, manifest(join(archive, 'candidates', candidate, 'sources', String(index), 'src')));
  return { root, files };
});
const audits = readdirSync(join(archive, 'runs')).sort().flatMap(label => {
  const path = join(archive, 'runs', label, 'results.json');
  if (!existsSync(path) || JSON.parse(readFileSync(path)).status !== 'passed') return [];
  return [{ label, ...auditRuns(archive, label) }];
});
writeJson(join(archive, 'selection-audits.json'), { status: 'passed', checked_at: new Date().toISOString(), audits });
const comparison = read('runs/composition-v3-full-table/results.json');
const confirmation = pairedAggregate(comparison.results, 'composition-no-range-v2', candidate);
assert.equal(confirmation.projects, 51);
assert(confirmation.ratio_of_summed_medians <= 0.99 && confirmation.bootstrap_95[1] < 1);
const arrays = comparison.results.find(item => item.name === 'gopurs-arrays');
assert(arrays.summary[candidate].median_ms / arrays.summary['composition-no-range-v2'].median_ms <= 0.97);
const retained = [
  { name: 'Direct syntax folds', experiment: 'direct-folds', reason: 'Eliminate temporary FreeMonoidTree construction; Arrays screening 3149 → 2068 ms and b8x 29798 → 28244 ms, with exact fold order and outputs.' },
  { name: 'Bulk directive removal', experiment: 'bulk-directives', reason: 'Persistent-map intersection/difference preserves rank visibility and shadowed defaults; b8x screening 28117 → 26668 ms.' },
  { name: 'Dynamic preparation scheduling', experiment: 'dynamic-preparation', reason: 'Small atomic work claims preserve ordered round merges; b8x screening 26211 → 23895 ms, with deferred/replayable bounded scheduling contracts.' },
  { name: 'Lexical closed-scope proof', experiment: 'closed-scope', reason: 'Avoid constructing free-variable sets for every subtree and reject cheaply before traversal; composition-v1 confirms Arrays 1900 → 1512 ms.' },
  { name: 'Import accumulator', experiment: 'import-accumulator', reason: 'One deduplicated file-level accumulator; full-table confirmation gives Arrays 1511 → 1442 ms with a lower overall total.' },
  { name: 'Observed specialization misses, stable runtime layout', experiment: 'precise-layout', reason: 'Invalidate only on relevant new membership, with a sequenced per-transform lookup journal and unchanged common runtime. Final composition confirms 76138 → 75168 ms overall and Arrays 1447 → 1325 ms.' },
];
const rejected = [
  { name: 'Module-range filtering', experiment: 'composition-no-range-v2', reason: 'Removed after 51-case ablation: its 0.83% overall and 2.42% b8x improvements remain below the agreed practical thresholds.' },
  { name: 'Annotation rewrite memo', experiment: 'type-rewrite-memo-v3', reason: 'All five screening cases regress by 1–2% despite passing contracts.' },
  { name: 'Lookup-first alone', experiment: 'lookup-first-v3', reason: 'Neutral alone; retained only as the lookup interface needed by the observed-miss mechanism, with no independent gain claimed.' },
  { name: 'Precise cache with extra record shapes', experiment: 'precise-dependencies-v3', reason: 'Complete-backend regressions despite faster preparation. The retained tuple variant keeps the shared generated runtime byte-identical; this does not prove a sole cause for the earlier regression.' },
  { name: 'Adaptive directive prefixes / shared range trees', experiment: 'prefix-directives, shared-range', reason: 'No gain at the practical thresholds in their paired screenings.' },
];
const nonGaining = [
  { mechanism: 'Reuse closed-type eligibility', experiment: 'cached-eligibility', reason: 'b8x +1.01%, Spec +2.08%; five-case sum 33387 → 33786 ms.' },
  { mechanism: 'Pruned contribution-suffix fold', experiment: 'suffix-contributions', reason: 'b8x −0.46%, paired interval crosses equality; five-case sum 33404 → 33291 ms, below practical thresholds.' },
  { mechanism: 'Chunked type-name construction', experiment: 'mangle-chunks', reason: 'b8x +15.38%, Arrays +5.52%; five-case sum 33484 → 38253 ms.' },
];
const promisingFinished = read('proposals/precise-layout/results.json').finished_at;
let previousFinished = promisingFinished;
for (const item of nonGaining) {
  const record = read(`proposals/${item.experiment}/results.json`);
  assert.equal(record.status, 'passed'); assert(record.started_at >= previousFinished);
  assert.equal(read(`runs/${item.experiment}/results.json`).status, 'passed'); previousFinished = record.finished_at;
}
const profiles = ['gopurs-arrays', 'b8x', 'gopurs-aff'].map(name => {
  const path = `profiles/${candidate}-${name}/summary.json`, profile = read(path);
  assert.equal(profile.status, 'passed'); assert.equal(profile.binary_sha256, build.binary_sha256);
  assert(profile.checked_at >= previousFinished);
  return { name, ...evidence(path), samples: profile.samples, capture_window: profile.capture_window,
    nearest_frames: profile.nearest_purescript_frames.slice(0, 15), allocation_callers: profile.allocation_callers.slice(0, 8) };
});
const checked_at = new Date().toISOString();
writeJson(join(archive, 'selection.json'), { status: 'passed', candidate, checked_at, sources, binary_sha256: build.binary_sha256,
  retained, rejected, confirmation, evidence: ['selection-audits.json', 'pre-selection/results.json',
    `candidates/${candidate}/build.json`, `proposals/${candidate}/results.json`, 'runs/composition-v2-full-table/results.json',
    'runs/composition-v3-full-table/results.json', 'screening-profile-provenance.json'].map(evidence) });
writeJson(join(archive, 'closure.json'), { status: 'passed', decision: 'stop', checked_at, candidate,
  non_gaining_mechanisms: nonGaining, profiles,
  rationale: 'The priority emission/allocation and preparation/monomorphization mechanisms have been investigated. After the last newly winning mechanism, three successive distinct mechanism experiments failed the practical thresholds; subsequent composition confirmation is not a new mechanism. Final profiles show the already-tested directive filtering/removal and type-name/eligibility costs on b8x, and distributed syntax traversal, analysis, scope and code generation on Arrays/Aff. No new isolated mechanism with a convincing threshold-sized benefit emerged from this review. This is practical saturation for this campaign and corpus, not a proof that further optimization is impossible.',
  evidence: ['profile-followups-sequence.json', `${candidate}-reprofile.json`, 'runs/composition-v3-full-table/results.json',
    ...nonGaining.map(item => `runs/${item.experiment}/results.json`),
    ...profiles.map(profile => `profiles/${candidate}-${profile.name}/summary.json`)].map(evidence) });
console.log(JSON.stringify({ candidate, audits: audits.length, confirmation, decision: 'stop' }, null, 2));
