// Independently recompute measurements and exact-output evidence from raw logs.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, jobs, manifest, median, verifyInput } from './common.mjs';

export function auditRuns(archive, label) {
  const path = join(archive, 'runs', label, 'results.json'), result = JSON.parse(readFileSync(path));
  assert.equal(result.status, 'passed');
  const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
  assert.equal(hash(readFileSync(join(archive, 'campaign.json'))), result.campaign_sha256);
  const definitions = new Map(campaign.cases.map(item => [item.name, item]));
  assert.equal(result.results.length, result.configuration.cases?.length ?? 51);
  assert.equal(new Set(result.results.map(item => item.name)).size, result.results.length);
  let generations = 0, measured = 0, files = 0;
  const totals = Object.fromEntries(result.variants.map(variant => [variant.name, 0]));
  for (const variant of result.variants) {
    assert.equal(hash(readFileSync(variant.binary)), variant.sha256);
    if (variant.compiler_artifact) assert.equal(hash(readFileSync(variant.compiler_artifact.path)), variant.compiler_artifact.sha256);
  }
  for (const tool of result.harness ?? []) assert.equal(hash(readFileSync(tool.path)), tool.sha256);
  for (const item of result.results) {
    assert.equal(item.status, 'passed');
    const definition = definitions.get(item.name); assert(definition);
    assert.equal(hash(readFileSync(definition.definition)), definition.sha256);
    const corpus = JSON.parse(readFileSync(definition.definition)); verifyInput(corpus);
    assert.deepEqual(manifest(corpus.oracle.directory), corpus.oracle.files);
    for (let round = 0; round <= (result.configuration.rounds ?? 5); round++) {
      const variants = result.variants, offset = round % variants.length;
      const order = [...variants.slice(offset), ...variants.slice(0, offset)].map(variant => variant.name);
      if (variants.length > 2 && Math.floor(round / variants.length) % 2) order.reverse();
      assert.deepEqual(item.runs.filter(run => run.round === round).map(run => run.variant), order);
    }
    for (const variant of result.variants) {
      const runs = item.runs.filter(run => run.variant === variant.name);
      assert.equal(runs.length, (result.configuration.rounds ?? 5) + 1);
      assert.deepEqual(runs.map(run => run.round), Array.from({ length: runs.length }, (_, i) => i));
      const samples = [];
      for (const run of runs) {
        assert.equal(run.exit_code, 0); assert.equal(run.measured, run.round > 0);
        assert.equal(run.cwd, corpus.input);
        assert.deepEqual(run.command, [variant.binary, ...(variant.arguments ?? []), ...corpus.invocation]);
        for (const [name, value] of Object.entries({ ...jobs, ...variant.environment }))
          assert.equal(run.explicit_environment[name], value, `${item.name}/${variant.name}: ${name}`);
        const raw = readFileSync(run.stderr, 'utf8');
        const timing = [...raw.matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)]; assert.equal(timing.length, 1);
        assert.equal(Number(timing[0][1]), run.phases_ms['backend total']);
        const phases = Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)]
          .map(([, name, milliseconds]) => [name, Number(milliseconds)]));
        assert.deepEqual(phases, run.phases_ms);
        const generated = JSON.parse(readFileSync(run.generated_manifest));
        assert.equal(hash(JSON.stringify(generated)), run.generated_sha256);
        assert.equal(generated.length, run.identical_files);
        assert.deepEqual(generated, corpus.oracle.files);
        if (run.measured) { samples.push(Number(timing[0][1])); measured++; }
        generations++; files += generated.length;
      }
      assert.deepEqual(samples, item.summary[variant.name].samples_ms);
      const ms = median(samples); assert.equal(ms, item.summary[variant.name].median_ms);
      const measuredRuns = runs.filter(run => run.measured);
      assert.deepEqual(Object.fromEntries(Object.keys(measuredRuns[0].phases_ms).map(phase =>
        [phase, median(measuredRuns.map(run => run.phases_ms[phase]))])), item.summary[variant.name].phases_median_ms);
      totals[variant.name] += ms;
    }
  }
  assert.deepEqual(totals, result.total_ms);
  return { path, sha256: hash(readFileSync(path)), status: 'passed', projects: result.results.length,
    auditor_sha256: hash(readFileSync(fileURLToPath(import.meta.url))),
    generations, measured_generations: measured, exact_files: files, total_ms: totals };
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  const [archive, ...labels] = process.argv.slice(2);
  for (const label of labels) console.log(JSON.stringify(auditRuns(resolve(archive), label), null, 2));
}
