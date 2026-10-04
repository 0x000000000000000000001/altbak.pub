// Paired log-ratio percentile bootstrap, deterministic and reproducible.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { hash, writeJson } from '../gopurs-aff/common.mjs';
const [resultsArg, reference, candidate, destination] = process.argv.slice(2);
assert(resultsArg && reference && candidate && destination,
  'paired-summary.mjs RESULTS.json REFERENCE CANDIDATE OUTPUT.json');
const source = resolve(resultsArg), bytes = readFileSync(source), campaign = JSON.parse(bytes);
assert.equal(campaign.status, 'passed');
const pair = round => {
  const one = name => {
    const records = campaign.runs.filter(run => !run.warmup && run.round === round && run.variant === name);
    assert.equal(records.length, 1);
    const run = records[0], raw = readFileSync(run.stderr, 'utf8');
    assert.equal(run.exit_code, 0); assert.equal(run.identical_files, 294);
    const clock = Number(raw.match(/^\[gopurs\] backend total: (\d+) ms$/m)?.[1]);
    assert.equal(clock, run.phases_ms['backend total']); assert(clock > 0);
    return clock;
  };
  const before = one(reference), after = one(candidate);
  return { round, reference_ms: before, candidate_ms: after, delta_ms: after - before, ratio: after / before };
};
const pairs = Array.from({ length: campaign.protocol.rounds }, (_, index) => pair(index + 1));
const median = values => { const sorted = values.toSorted((a, b) => a - b), middle = sorted.length >> 1;
  return sorted.length % 2 ? sorted[middle] : (sorted[middle - 1] + sorted[middle]) / 2; };
const logs = pairs.map(pair => Math.log(pair.ratio));
const seed = 0x4e495447;
let state = seed;
const random = () => { state ^= state << 13; state ^= state >>> 17; state ^= state << 5; return (state >>> 0) / 0x100000000; };
const repetitions = 100000, bootstrapped = new Array(repetitions);
for (let iteration = 0; iteration < repetitions; iteration++) {
  let total = 0;
  for (let index = 0; index < logs.length; index++) total += logs[Math.floor(random() * logs.length)];
  bootstrapped[iteration] = total / logs.length;
}
bootstrapped.sort((a, b) => a - b);
const percentile = probability => {
  const position = probability * (repetitions - 1), lower = Math.floor(position), weight = position - lower;
  return Math.exp(bootstrapped[lower] * (1 - weight) + bootstrapped[Math.min(lower + 1, repetitions - 1)] * weight);
};
const summary = {
  source, source_sha256: hash(bytes), reference, candidate, pairs,
  median_reference_ms: median(pairs.map(pair => pair.reference_ms)),
  median_candidate_ms: median(pairs.map(pair => pair.candidate_ms)),
  favorable_pairs: pairs.filter(pair => pair.ratio < 1).length,
  median_paired_delta_ms: median(pairs.map(pair => pair.delta_ms)),
  geometric_mean_paired_ratio: Math.exp(logs.reduce((a, b) => a + b, 0) / logs.length),
  bootstrap: { method: 'paired resampling of log ratios; percentile CI for geometric mean ratio',
    repetitions, seed, confidence: 0.95, ratio_interval: [percentile(0.025), percentile(0.975)] },
};
summary.median_change_percent = 100 * (summary.median_candidate_ms / summary.median_reference_ms - 1);
summary.nightly_30_pair_gate = pairs.length === 30 && summary.median_candidate_ms < summary.median_reference_ms
  && summary.bootstrap.ratio_interval[1] < 1 && summary.favorable_pairs >= 27;
writeJson(resolve(destination), summary);
console.log(JSON.stringify(summary, null, 2));
