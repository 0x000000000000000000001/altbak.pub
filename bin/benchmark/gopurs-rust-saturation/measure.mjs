// Compare arbitrary host candidates serially on the same published corpora.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { run, writeJson } from '../gopurs-aff/common.mjs';
import { clean, copy, emitted, environment, hash, jobs, median, paired, verifyInput, verifyOutput } from './common.mjs';

const [archiveArg, label, configurationArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && configurationArg);
const archive = resolve(archiveArg), config = JSON.parse(readFileSync(resolve(configurationArg)));
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
assert.equal(campaign.status, 'passed');
const directory = join(archive, 'runs', label); assert(!existsSync(directory)); mkdirSync(directory, { recursive: true });
const harness = [['measure.mjs', import.meta.url], ['common.mjs', new URL('./common.mjs', import.meta.url)],
  ['gopurs-aff-common.mjs', new URL('../gopurs-aff/common.mjs', import.meta.url)]].map(([name, url]) => {
  const source = fileURLToPath(url), path = join(directory, 'tools', name);
  copy(source, path);
  return { source, path, sha256: hash(readFileSync(path)) };
});
const variants = config.variants.map(variant => {
  const binary = resolve(archive, variant.binary);
  const compiler = variant.compiler ? resolve(archive, variant.compiler) : binary;
  return { ...variant, binary, sha256: hash(readFileSync(binary)),
    compiler_artifact: { path: compiler, sha256: hash(readFileSync(compiler)) } };
});
assert.equal(new Set(variants.map(variant => variant.name)).size, variants.length);
const rounds = config.rounds ?? 5;
assert(Number.isInteger(rounds) && rounds > 0);
const state = { status: 'running', started_at: new Date().toISOString(), configuration: config,
  variants, harness, campaign_sha256: hash(readFileSync(join(archive, 'campaign.json'))), results: [] };
const save = () => writeJson(join(directory, 'results.json'), state);
save();
try {
  for (const definition of campaign.cases.filter(item => !config.cases || config.cases.includes(item.name))) {
    assert.equal(hash(readFileSync(definition.definition)), definition.sha256);
    const item = JSON.parse(readFileSync(definition.definition));
    const output = join(directory, item.name); mkdirSync(output);
    const result = { name: item.name, status: 'running', runs: [] }; state.results.push(result); save();
    clean(item); verifyInput(item);
    for (let round = 0; round <= rounds; round++) {
      const offset = round % variants.length;
      const order = [...variants.slice(offset), ...variants.slice(0, offset)];
      if (variants.length > 2 && Math.floor(round / variants.length) % 2) order.reverse();
      for (const variant of order) {
        clean(item); verifyInput(item);
        const name = `${round ? 'run-' + round : 'warmup'}-${variant.name}`;
        const record = run(output, name, variant.binary, [...(variant.arguments ?? []), ...item.invocation], item.input,
          { ...environment(), ...jobs, ...variant.environment, GOWORK: 'off' }, 900000);
        Object.assign(record, { variant: variant.name, round, measured: round > 0 });
        result.runs.push(record); save();
        const raw = readFileSync(record.stderr, 'utf8');
        assert.equal([...raw.matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)].length, 1);
        record.phases_ms = Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, name, ms]) => [name, Number(ms)]));
        assert(record.phases_ms['backend total'] > 0);
        const files = emitted(item.input);
        record.generated_manifest = join(output, 'logs', name + '.files.json');
        writeJson(record.generated_manifest, files);
        record.generated_sha256 = hash(JSON.stringify(files)); save();
        verifyOutput(item, files, join(output, 'failed-' + name));
        record.identical_files = files.length;
        writeJson(join(output, 'logs', name + '.json'), record); save();
        console.log(`${label} ${item.name} ${name}: ${record.phases_ms['backend total']} ms; ${files.length} exact files`);
      }
    }
    result.summary = Object.fromEntries(variants.map(variant => {
      const runs = result.runs.filter(run => run.measured && run.variant === variant.name);
      const samples = runs.map(run => run.phases_ms['backend total']);
      const phases = Object.keys(runs[0].phases_ms);
      return [variant.name, { samples_ms: samples, median_ms: median(samples), min_ms: Math.min(...samples), max_ms: Math.max(...samples),
        phases_median_ms: Object.fromEntries(phases.map(phase => [phase, median(runs.map(run => run.phases_ms[phase]))])) }];
    }));
    result.paired = Object.fromEntries(variants.slice(1).map(variant => [variant.name,
      paired(result.summary[variants[0].name].samples_ms, result.summary[variant.name].samples_ms)]));
    clean(item); verifyInput(item); result.status = 'passed'; save();
  }
  assert.equal(state.results.length, config.cases?.length ?? 51);
  for (const variant of variants) {
    assert.equal(hash(readFileSync(variant.binary)), variant.sha256);
    assert.equal(hash(readFileSync(variant.compiler_artifact.path)), variant.compiler_artifact.sha256);
  }
  state.total_ms = Object.fromEntries(variants.map(variant => [variant.name,
    state.results.reduce((sum, result) => sum + result.summary[variant.name].median_ms, 0)]));
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, total_ms: state.total_ms }, null, 2));
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
