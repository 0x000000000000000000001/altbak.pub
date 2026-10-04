// Fill the library-table Rust measurements with frozen, output-checked runs.
import assert from 'node:assert/strict';
import { constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync,
  readdirSync, rmSync, statSync, symlinkSync, writeFileSync } from 'node:fs';
import { basename, delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { cpus, totalmem } from 'node:os';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler, verifyTypedOutput } from '../../../../gopurs/gopurs/tools/native-workspace.mjs';

const here = dirname(fileURLToPath(import.meta.url));
const site = resolve(here, '../../..'), compiler = resolve(site, '../gopurs/gopurs');
const [archiveArg, roundsArg = '5'] = process.argv.slice(2);
assert(archiveArg, 'measure.mjs NEW_ARCHIVE [ROUNDS]');
const archive = resolve(archiveArg), rounds = Number(roundsArg);
assert(Number.isInteger(rounds) && rounds >= 3); assert(!existsSync(archive)); mkdirSync(archive);
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const readme = readFileSync(join(site, 'README.md'), 'utf8');
const packages = [...readme.matchAll(/^\[(gopurs-[a-z0-9-]+)\][^\n]+\| \(WIP\)\s*$/gm)].map(match => match[1]);
assert(packages.length > 0); assert(!packages.includes('gopurs-aff'));
copy(join(site, 'README.md'), join(archive, 'README-before.md'));
copy(fileURLToPath(import.meta.url), join(archive, 'measure.mjs'));
copy(join(here, '../gopurs-aff/common.mjs'), join(archive, 'common.mjs'));
const frozenCompiler = join(archive, 'compiler');
for (const name of ['bin/gopurs', 'bin/gopurs-native', 'bin/gopurs-rust', 'bin/gopurs.js',
  'package.json', 'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js'])
  copy(join(compiler, name), join(frozenCompiler, name));
const frontend = join(archive, 'frontend/purs');
copy(findTypedCompiler(compiler, process.env.GOPURS_PURS), frontend);
const env = { ...environment(), GHCRTS: '-N2', GOWORK: 'off',
  PATH: [dirname(frontend), join(compiler, 'node_modules/.bin'), process.env.PATH].join(delimiter) };
const workers = { GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' };
const family = join(archive, 'family'); mkdirSync(family);
const omitted = new Set(['.git', '.spago', '.cache', '.purmeta', 'node_modules', 'output', '.stack-work', '.pulp-cache', '.DS_Store']);
for (const name of readdirSync(dirname(compiler)).sort()) {
  if (!name.startsWith('gopurs-')) continue;
  const source = join(dirname(compiler), name), target = join(family, name);
  if (!statSync(source).isDirectory()) continue;
  cpSync(source, target, { recursive: true, verbatimSymlinks: true, mode: constants.COPYFILE_FICLONE,
    filter: path => !omitted.has(basename(path)) });
  if (existsSync(join(target, 'spago.go.yaml'))) {
    rmSync(join(target, 'spago.yaml'), { force: true });
    copy(join(target, 'spago.go.yaml'), join(target, 'spago.yaml'));
  }
}
const state = { status: 'pending', started_at: new Date().toISOString(),
  packages, host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  protocol: { metric: '[gopurs] backend total', target: 'Go', rounds, rust_warmups: 1,
    jobs: workers, oracle: 'Go host on the same frozen inputs; JavaScript fallback if Go generation fails',
    statistic: 'median of successful Rust backend generations, each byte-exact to the oracle',
    order: 'serialized packages, frontend outside timing, one oracle and one Rust warmup before measured Rust runs',
    excludes: 'frontend, compiler/application builds, application execution and process startup/exit',
    cache: 'fresh generated output and PBO caches for each process, warm OS filesystem cache',
    scope: 'Only unmeasured gopurs-* Rust cells; existing JS/Go and gopurs-aff cells keep their previous campaigns. No new cross-host speed-ratio claim.' },
  compiler: manifest(frozenCompiler), frontend_sha256: hash(readFileSync(frontend)),
  family_sources: manifest(family), results: [] };
const save = () => writeJson(join(archive, 'results.json'), state);
save();
const generatedFiles = output => manifest(output, path => path.endsWith('.go') || basename(path) === 'go.mod');
const clean = input => {
  for (const path of walk(join(input, 'output'))) if (path.endsWith('.go') || ['go.mod', 'go.sum'].includes(basename(path))) rmSync(path);
  for (const name of ['.cache', '.purmeta']) rmSync(join(input, name), { force: true, recursive: true });
};
function freeze(project, input) {
  const modules = [];
  for (const name of readdirSync(join(project, 'output')).sort()) {
    const file = join(project, 'output', name, 'corefn.json'); if (!existsSync(file)) continue;
    const bytes = readFileSync(file), json = JSON.parse(bytes);
    const original = resolve(project, json.modulePath), source = `sources/${name}/${basename(original)}`;
    assert(existsSync(original));
    for (const extension of ['purs', 'go', 'js']) {
      const sibling = original.replace(/\.purs$/, '.' + extension);
      if (existsSync(sibling)) copy(sibling, join(input, source.replace(/\.purs$/, '.' + extension)));
    }
    json.modulePath = source;
    const output = join(input, 'output', name, 'corefn.json'); mkdirSync(dirname(output), { recursive: true });
    writeFileSync(output, JSON.stringify(json) + '\n');
    modules.push({ name, original_path: original, original_sha256: hash(bytes),
      frozen_sha256: hash(readFileSync(output)), bytes: statSync(output).size, types: json.typeTable.length });
  }
  assert(modules.some(module => module.name === 'Test.Main'));
  return { modules, manifest: manifest(input) };
}
for (const name of packages) {
  const directory = join(archive, 'packages', name), project = join(family, name), input = join(directory, 'input');
  mkdirSync(directory, { recursive: true });
  const result = { name, status: 'pending', started_at: new Date().toISOString(), commands: [], runs: [] };
  state.results.push(result); save();
  const retain = () => { writeJson(join(directory, 'results.json'), result); save(); };
  const execute = (label, command, args, cwd, overrides = {}, timeout = 300000) => {
    const record = run(directory, label, command, args, cwd, { ...env, ...overrides }, timeout);
    result.commands.push(record); retain(); return record;
  };
  const generate = (label, host, warmup) => {
    clean(input); assert.deepEqual(manifest(input), result.inputs.manifest);
    const record = execute(label, join(frozenCompiler, 'bin/gopurs'), ['--main', 'Test.Main'], input,
      { ...workers, GOPURS_RUST: host === 'rust' ? '1' : '0', GOPURS_JS: host === 'js' ? '1' : '0' });
    const raw = readFileSync(record.stderr, 'utf8');
    const clocks = [...raw.matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
    assert.equal(clocks.length, 1); assert(Number(clocks[0][1]) > 0);
    const files = generatedFiles(join(input, 'output')); assert(files.length > 0);
    const output = join(directory, 'generated', label);
    for (const file of files) copy(join(input, 'output', file.path), join(output, file.path));
    Object.assign(record, { host, warmup, output, generated: files,
      phases_ms: Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, phase, ms]) => [phase, Number(ms)])) });
    result.runs.push(record); retain();
    if (result.oracle) assert.deepEqual(files, result.oracle.generated, `${name}: ${host} generated different Go`);
    record.identical_files = files.length; retain();
    clean(input); assert.deepEqual(manifest(input), result.inputs.manifest);
    console.log(`${name} ${label}: ${record.phases_ms['backend total']} ms; ${files.length} exact files`);
    return record;
  };
  try {
    console.log(`${name}: frontend`);
    execute('frontend', 'spago', ['build'], project, {}, 600000);
    result.tast = verifyTypedOutput(join(project, 'output'), frontend);
    result.inputs = freeze(project, input); retain();
    try { result.oracle = generate('oracle-go', 'go', true); }
    catch (error) { result.go_oracle_failure = error.stack; retain(); result.oracle = generate('oracle-js', 'js', true); }
    retain();
    generate('warmup-rust', 'rust', true);
    for (let round = 1; round <= rounds; round++) generate(`run-${round}-rust`, 'rust', false);
    const runs = result.runs.filter(record => record.host === 'rust' && !record.warmup);
    const samples = runs.map(record => record.phases_ms['backend total']), sorted = samples.toSorted((a, b) => a - b);
    result.summary = { samples_ms: samples, median_ms: sorted.length % 2 ? sorted[sorted.length >> 1]
      : (sorted[sorted.length / 2 - 1] + sorted[sorted.length / 2]) / 2,
      min_ms: sorted[0], max_ms: sorted.at(-1), identical_files_per_generation: result.oracle.generated.length };
    result.status = 'passed'; result.finished_at = new Date().toISOString(); retain();
    console.log(`${name}: PASS, median ${result.summary.median_ms} ms`);
  } catch (error) {
    result.status = 'failed'; result.failure = error.stack; result.finished_at = new Date().toISOString(); retain();
    console.error(`${name}: FAILED: ${error.message}`);
  }
}
assert.deepEqual(manifest(frozenCompiler), state.compiler);
assert.equal(hash(readFileSync(frontend)), state.frontend_sha256);
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'partial';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, passed: state.results.filter(result => result.status === 'passed').length,
  failed: state.results.filter(result => result.status !== 'passed').map(result => result.name) }, null, 2));
