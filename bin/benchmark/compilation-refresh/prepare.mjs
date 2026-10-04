// Freeze the qualified compilers and all fifty library corpora for a fresh three-host campaign.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { cpus, totalmem } from 'node:os';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const here = dirname(fileURLToPath(import.meta.url));
const site = resolve(here, '../../..'), htdocs = dirname(site);
const archive = resolve(process.argv[2]);
assert(!existsSync(join(archive, 'campaign.json')));
mkdirSync(archive, { recursive: true });
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const qualified = join(site, 'var/benchmark/gopurs-packages-fixes-20261004/revision2');
const qualificationPath = join(qualified, 'build-results.json');
const qualification = JSON.parse(readFileSync(qualificationPath));
assert.equal(qualification.status, 'passed');
const priorPath = join(qualified, 'recheck/final-results.json');
const prior = JSON.parse(readFileSync(priorPath));
assert.equal(prior.status, 'passed'); assert.equal(prior.results.length, 50);
copy(join(site, 'README.md'), join(archive, 'README-before.md'));
const compilers = {};
for (const family of ['gopurs', 'purust']) {
  const source = join(htdocs, family, family), destination = join(archive, 'compilers', family);
  const files = family === 'gopurs'
    ? ['bin/gopurs', 'bin/gopurs.js', 'bin/gopurs-native', 'bin/gopurs-rust', 'package.json', 'tools/ffi-runner.mjs', 'tools/ffi_gen.wasm', 'tools/wasm_exec.js']
    : ['bin/purust', 'bin/purust.js', 'bin/purust-native', 'package.json'];
  for (const file of files) {
    const expected = qualification.binaries.find(item => item.path === file.slice(4));
    if (file.startsWith('bin/') && !(family === 'purust' && file === 'bin/purust')) {
      assert(expected, file); assert.equal(hash(readFileSync(join(source, file))), expected.sha256);
    }
    copy(join(source, file), join(destination, file));
  }
  compilers[family] = { directory: destination, origin: source, files: manifest(destination) };
}
const provenance = [];
for (const name of ['gopurs/gopurs', 'purust/purust', 'purescript-backend-optimizer-gopurs', 'purescript-backend-optimizer-purust']) {
  const source = join(htdocs, name), destination = join(archive, 'sources', name);
  const files = manifest(join(source, 'src'));
  for (const file of files) copy(join(source, 'src', file.path), join(destination, file.path));
  provenance.push({ repository: name, commit: execFileSync('git', ['rev-parse', 'HEAD'], { cwd: source, encoding: 'utf8' }).trim(),
    status: execFileSync('git', ['status', '--short'], { cwd: source, encoding: 'utf8' }), files });
}
const cases = [];
for (const old of prior.results.toSorted((a, b) => a.name.localeCompare(b.name))) {
  assert.equal(old.status, 'passed');
  const directory = join(archive, 'cases', old.name), input = join(directory, 'input');
  for (const file of old.inputs) {
    const source = join(old.input_source, file.path);
    assert.equal(hash(readFileSync(source)), file.sha256);
    copy(source, join(input, file.path));
  }
  for (const file of old.sibling_inputs ?? []) {
    const source = join(old.sibling_source, file.path);
    assert.equal(hash(readFileSync(source)), file.sha256);
    copy(source, join(directory, file.path));
  }
  assert.deepEqual(manifest(input), old.inputs);
  const oracle = manifest(old.oracle.output, path => path.endsWith('.go') || path.endsWith('/go.mod'));
  assert.deepEqual(oracle, old.oracle.generated);
  const definition = { name: old.name, family: 'gopurs', input, input_manifest: old.inputs,
    sibling_inputs: old.sibling_inputs ?? [], invocation: old.invocation, hosts: ['js', 'go', 'rust'],
    historical_oracle: { directory: old.oracle.output, files: oracle }, source_corpus: old.input_source,
    oracle_correction: old.oracle_correction ?? null };
  const path = join(directory, 'definition.json'); writeJson(path, definition); cases.push(path);
}
const state = { status: 'prepared', started_at: new Date().toISOString(),
  host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  compilers, provenance, qualification: { path: qualificationPath, sha256: hash(readFileSync(qualificationPath)),
    libraries: priorPath, libraries_sha256: hash(readFileSync(priorPath)) },
  protocol: { rounds: 5, warmups_per_host: 1, metric: 'backend total', statistic: 'median',
    includes: 'loading, preparation, PBO, generation/emission and drain',
    excludes: 'frontend, compiler/application builds, application execution and process startup/exit',
    cache: 'fresh generated files and PBO caches; warm OS filesystem cache; fresh process each invocation',
    order: 'serialized cases; six host permutations over warmup and five measured rounds',
    gopurs_jobs: { GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' },
    purust_jobs: 'public defaults: JS 1/1; Rust PBO 8 and codegen 4',
    go_gc: 'public launcher: GOGC=off and GOMEMLIMIT=10GiB on this 48-GiB host',
    exact_output: 'all generated sources and manifests compared byte-for-byte; each run retains a manifest, canonical identical output retained once per case',
    phases: 'overlapping producer, cumulative generation/write batches and drain are reported separately, never added' },
  library_cases: cases };
writeJson(join(archive, 'campaign.json'), state);
for (const file of manifest(here)) copy(join(here, file.path), join(archive, 'tools', file.path));
console.log(`Prepared ${cases.length} library cases and five qualified compiler hosts in ${archive}`);
