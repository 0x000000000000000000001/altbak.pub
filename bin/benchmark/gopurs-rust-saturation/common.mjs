import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { basename, dirname, join } from 'node:path';
import { environment, hash, manifest, walk } from '../gopurs-aff/common.mjs';

export { environment, hash, manifest };
export const jobs = Object.freeze({ GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8',
  GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' });
export function copy(from, to) {
  mkdirSync(dirname(to), { recursive: true });
  copyFileSync(from, to, constants.COPYFILE_FICLONE);
}
export function clean(item) {
  const protectedPaths = new Set(item.input_manifest.map(file => file.path));
  for (const path of walk(join(item.input, 'output'))) {
    if (path.endsWith('.go') || ['go.mod', 'go.sum'].includes(basename(path))) {
      assert(!protectedPaths.has(path.slice(item.input.length + 1)), path);
      rmSync(path);
    }
  }
  for (const name of ['.cache', '.purmeta']) rmSync(join(item.input, name), { recursive: true, force: true });
}
export function verifyInput(item) {
  assert.deepEqual(manifest(item.input), item.input_manifest, item.name);
  for (const file of item.sibling_inputs ?? [])
    assert.equal(hash(readFileSync(join(dirname(item.input), file.path))), file.sha256, file.path);
}
export function emitted(input) {
  return manifest(join(input, 'output'), path => path.endsWith('.go') || basename(path) === 'go.mod');
}
export function verifyOutput(item, files, failedDirectory) {
  try {
    assert.deepEqual(files, item.oracle.files, `${item.name}: generated manifest differs`);
    for (const file of files) assert(readFileSync(join(item.input, 'output', file.path))
      .equals(readFileSync(join(item.oracle.directory, file.path))), file.path);
  } catch (error) {
    assert(!existsSync(failedDirectory));
    for (const file of files) copy(join(item.input, 'output', file.path), join(failedDirectory, file.path));
    throw error;
  }
}
export const median = values => {
  const sorted = values.toSorted((a, b) => a - b), middle = Math.floor(sorted.length / 2);
  assert(sorted.length);
  return sorted.length % 2 ? sorted[middle] : (sorted[middle - 1] + sorted[middle]) / 2;
};

// Paired rounds stay together. The interval describes this case's log-ratio;
// it is not a prediction from historical runs or a sum of overlapping phases.
export function paired(before, after) {
  assert.equal(before.length, after.length);
  const logs = before.map((value, index) => Math.log(after[index] / value));
  let seed = 0x74af0193;
  const random = () => { seed ^= seed << 13; seed ^= seed >>> 17; seed ^= seed << 5; return (seed >>> 0) / 4294967296; };
  const samples = Array.from({ length: 20000 }, () => Math.exp(logs.reduce(sum => sum + logs[Math.floor(random() * logs.length)], 0) / logs.length))
    .sort((a, b) => a - b);
  return { pairs: logs.length, favorable: logs.filter(value => value < 0).length,
    ratio_of_medians: median(after) / median(before),
    geometric_mean_ratio: Math.exp(logs.reduce((a, b) => a + b, 0) / logs.length),
    bootstrap_95: [samples[500], samples[19499]], seed: '0x74af0193', replicates: samples.length };
}

// Resample paired rounds within each project, then sum project medians. This
// quantifies observed within-project variability without treating per-project
// times as samples of one large compiler invocation.
export function pairedAggregate(results, beforeName, afterName) {
  const cases = results.map(result => {
    const before = result.summary[beforeName].samples_ms, after = result.summary[afterName].samples_ms;
    assert.equal(before.length, after.length); assert(before.length > 0);
    return { before, after };
  });
  const totals = cases.reduce((sum, item) => ({ before: sum.before + median(item.before), after: sum.after + median(item.after) }),
    { before: 0, after: 0 });
  let seed = 0x6a09e667;
  const random = () => { seed ^= seed << 13; seed ^= seed >>> 17; seed ^= seed << 5; return (seed >>> 0) / 4294967296; };
  const samples = Array.from({ length: 20000 }, () => {
    let before = 0, after = 0;
    for (const item of cases) {
      const indexes = Array.from({ length: item.before.length }, () => Math.floor(random() * item.before.length));
      before += median(indexes.map(index => item.before[index]));
      after += median(indexes.map(index => item.after[index]));
    }
    return after / before;
  }).sort((a, b) => a - b);
  return { projects: cases.length, before_ms: totals.before, after_ms: totals.after,
    ratio_of_summed_medians: totals.after / totals.before, bootstrap_95: [samples[500], samples[19499]],
    seed: '0x6a09e667', replicates: samples.length, method: 'Paired resampling within projects; ratio of sums of project medians' };
}
