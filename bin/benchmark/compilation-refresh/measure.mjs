// Run a serialized, exact-output-checked compilation campaign on frozen inputs.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, rmSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, group = 'libraries', ...definitionArgs] = process.argv.slice(2);
assert(archiveArg);
const archive = resolve(archiveArg), campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const outputPath = join(archive, `${group}-results.json`); assert(!existsSync(outputPath));
const definitions = definitionArgs.length ? definitionArgs.map(resolvePath => resolve(resolvePath)) : campaign.library_cases;
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const state = { status: 'running', started_at: new Date().toISOString(), campaign_sha256: hash(readFileSync(join(archive, 'campaign.json'))), results: [] };
const save = () => writeJson(outputPath, state);
const sourceFiles = (output, family) => manifest(output, path => family === 'gopurs'
  ? path.endsWith('.go') || basename(path) === 'go.mod' : /\.(rs|toml)$/.test(path));
const clean = (input, family) => {
  if (family === 'gopurs') for (const path of walk(join(input, 'output'))) {
    if (path.endsWith('.go') || ['go.mod', 'go.sum'].includes(basename(path))) rmSync(path);
  }
  for (const name of ['.cache', '.purmeta']) rmSync(join(input, name), { recursive: true, force: true });
};
const verifyInput = item => {
  assert.deepEqual(manifest(item.input), item.input_manifest);
  for (const file of item.sibling_inputs ?? []) assert.equal(hash(readFileSync(join(dirname(item.input), file.path))), file.sha256);
};
for (const compiler of Object.values(campaign.compilers)) assert.deepEqual(manifest(compiler.directory), compiler.files);
save();
for (const definition of definitions) {
  const item = JSON.parse(readFileSync(definition)), directory = dirname(definition), family = item.family;
  const result = { name: item.name, family, status: 'running', definition, definition_sha256: hash(readFileSync(definition)), runs: [] };
  state.results.push(result); save();
  try {
    clean(item.input, family); verifyInput(item);
    const orders = item.hosts.length === 3 ? [
      ['js', 'go', 'rust'], ['go', 'rust', 'js'], ['rust', 'js', 'go'],
      ['js', 'rust', 'go'], ['rust', 'go', 'js'], ['go', 'js', 'rust'],
    ] : [['js', 'rust'], ['rust', 'js'], ['js', 'rust'], ['rust', 'js'], ['js', 'rust'], ['rust', 'js']];
    let oracle = item.historical_oracle;
    for (let round = 0; round <= campaign.protocol.rounds; round++) {
      for (const host of orders[round]) {
        assert(item.hosts.includes(host));
        const label = `${round ? `run-${round}` : 'warmup'}-${host}`;
        clean(item.input, family); verifyInput(item);
        const generated = family === 'gopurs' ? join(item.input, 'output') : join(directory, 'scratch-generated');
        if (family === 'purust') { assert(!existsSync(generated)); }
        const args = [...item.invocation, ...(family === 'purust' ? ['--out', generated] : [])];
        const env = { ...environment(), GOWORK: 'off',
          ...(family === 'gopurs' ? { ...campaign.protocol.gopurs_jobs, GOPURS_JS: host === 'js' ? '1' : '0', GOPURS_RUST: host === 'rust' ? '1' : '0' }
            : { PURUST_JS: host === 'js' ? '1' : '0' }) };
        const record = run(directory, label, join(campaign.compilers[family].directory, 'bin', family), args, item.input, env, item.timeout_ms ?? 600000);
        record.host = host; record.round = round; record.measured = round > 0;
        result.runs.push(record); save();
        const raw = readFileSync(record.stderr, 'utf8');
        const total = [...raw.matchAll(new RegExp(`^\\[${family}\\] backend total: (\\d+) ms$`, 'gm'))];
        assert.equal(total.length, 1); assert(Number(total[0][1]) > 0);
        record.phases_ms = Object.fromEntries([...raw.matchAll(new RegExp(`^\\[${family}\\] (.+): (\\d+) ms$`, 'gm'))].map(([, name, ms]) => [name, Number(ms)]));
        const files = sourceFiles(generated, family); assert(files.length > 0);
        record.generated_manifest = join(directory, 'logs', `${label}.files.json`);
        record.generated_sha256 = hash(JSON.stringify(files)); writeJson(record.generated_manifest, files); save();
        if (!oracle) {
          const output = join(directory, 'canonical-generated');
          for (const file of files) copy(join(generated, file.path), join(output, file.path));
          oracle = { directory: output, files };
        }
        const expected = new Map(oracle.files.map(file => [file.path, file.sha256]));
        const actual = new Map(files.map(file => [file.path, file.sha256]));
        record.different_files = [...new Set([...expected.keys(), ...actual.keys()])].filter(path => expected.get(path) !== actual.get(path));
        if (record.different_files.length) {
          const failure = join(directory, 'failed-generated', label);
          for (const file of files) copy(join(generated, file.path), join(failure, file.path));
          record.failed_output = failure; save();
        }
        assert.deepEqual(record.different_files, [], `${item.name} ${label}: output mismatch`);
        for (const file of files) assert(readFileSync(join(generated, file.path)).equals(readFileSync(join(oracle.directory, file.path))), file.path);
        const canonical = join(directory, 'canonical-generated');
        if (!existsSync(canonical)) for (const file of files) copy(join(generated, file.path), join(canonical, file.path));
        record.retained_output = canonical; record.identical_files = files.length;
        writeJson(join(directory, 'logs', `${label}.json`), record); save();
        console.log(`${item.name} ${label}: ${record.phases_ms['backend total']} ms, ${files.length} exact files`);
        if (family === 'purust') rmSync(generated, { recursive: true });
      }
    }
    result.summary = Object.fromEntries(item.hosts.map(host => {
      const runs = result.runs.filter(run => run.measured && run.host === host), samples = runs.map(run => run.phases_ms['backend total']);
      assert.equal(samples.length, 5);
      const phases = [...new Set(runs.flatMap(run => Object.keys(run.phases_ms)))];
      return [host, { samples_ms: samples, median_ms: samples.toSorted((a, b) => a - b)[2], min_ms: Math.min(...samples), max_ms: Math.max(...samples),
        phases_median_ms: Object.fromEntries(phases.map(phase => [phase, runs.map(run => run.phases_ms[phase]).toSorted((a, b) => a - b)[2]])) }];
    }));
    clean(item.input, family); verifyInput(item);
    result.inputs_unchanged = true; result.status = 'passed';
  } catch (error) {
    result.status = 'failed'; result.error = error.stack; console.error(`${item.name}: ${error.message}`);
  }
  result.finished_at = new Date().toISOString(); save();
}
for (const compiler of Object.values(campaign.compilers)) assert.deepEqual(manifest(compiler.directory), compiler.files);
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'failed';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, cases: state.results.length, failed: state.results.filter(result => result.status !== 'passed').map(result => result.name) }, null, 2));
if (state.status !== 'passed') process.exitCode = 1;
