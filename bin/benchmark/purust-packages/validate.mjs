// Build the canonical two-host output once and run each primary application outside timing.
import assert from 'node:assert/strict';
import { constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, symlinkSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { basename, delimiter, dirname, join, resolve } from 'node:path';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, resultsArg = 'libraries-results.json', group = 'validation', ...names] = process.argv.slice(2);
const archive = resolve(archiveArg), campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const results = JSON.parse(readFileSync(join(archive, resultsArg)));
const root = join(archive, group); assert(!existsSync(root)); mkdirSync(root, { recursive: true });
const state = { status: 'running', started_at: new Date().toISOString(), results: [] };
const save = () => writeJson(join(root, 'results.json'), state); save();
const target = join(archive, 'validation-cargo-target');
const env = { ...environment(), CARGO_NET_OFFLINE: 'true', CARGO_TARGET_DIR: target,
  CARGO_PROFILE_DEV_DEBUG: '0', CARGO_PROFILE_DEV_INCREMENTAL: 'false', CARGO_BUILD_JOBS: '8',
  RUSTFLAGS: '-Awarnings', RUST_MIN_STACK: '67108864',
  PATH: [dirname(campaign.frontend.path), dirname(campaign.spago), process.env.PATH].join(delimiter) };
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const ignored = new Set(['.git', '.spago', '.cache', '.purmeta', 'node_modules', 'output', 'target', '.DS_Store']);
const copyProject = (from, to) => cpSync(from, to, { recursive: true, mode: constants.COPYFILE_FICLONE,
  filter: path => !ignored.has(basename(path)) });
const cuts = {
  'purust-arrays': 'echo "=== Checking Rust array FFI',
  'purust-avar': 'echo "=== Checking cancellation cleanup',
  'purust-foldable-traversable': 'echo "=== Checking the four Rust FFI',
  'purust-foreign': 'echo "=== Checking Rust Foreign FFI',
  'purust-js-promise': 'echo "=== Running the native Promise contract',
  'purust-node-streams': 'echo "=== Running the overflow PassThrough',
  'purust-random': 'echo "=== Checking the Rust random source',
  'purust-spec-discovery': 'echo "=== Running the empty-result discovery',
  'purust-unfoldable': 'echo "=== Checking native FFI step counts',
};
for (const result of results.results.filter(result => result.status === 'passed' && (!names.length || names.includes(result.name)))) {
  const item = { name: result.name, status: 'running', commands: [] }; state.results.push(item); save();
  const directory = join(root, result.name); mkdirSync(directory);
  const definition = JSON.parse(readFileSync(result.definition));
  const execute = (label, command, args, cwd, overrides = {}) => {
    const record = run(directory, label, command, args, cwd, { ...env, ...overrides }, 900000);
    item.commands.push(record); save(); return record;
  };
  const build = (generated, label) => {
    execute(label, 'cargo', ['build', '--manifest-path', join(generated, 'Cargo.toml')], generated);
    const binary = join(directory, label + '-binary'); copy(join(target, 'debug/purust_output'), binary);
    mkdirSync(join(generated, 'target/debug'), { recursive: true });
    symlinkSync(binary, join(generated, 'target/debug/purust_output'));
    return { path: binary, sha256: hash(readFileSync(binary)) };
  };
  try {
    const project = join(directory, 'project'); copyProject(definition.source_corpus, project);
    const generated = join(project, 'output/purust_output'), canonical = join(dirname(result.definition), 'canonical-generated');
    for (const file of manifest(canonical)) copy(join(canonical, file.path), join(generated, file.path));
    item.generated_manifest = manifest(canonical); item.definition_sha256 = hash(readFileSync(result.definition));
    item.binary = build(generated, 'build');
    if (result.name === 'purust-spec-node') {
      const { parse, stringify } = createRequire(join(campaign.compilers.purust.origin, 'package.json'))('yaml');
      item.fixture_binaries = [];
      for (const name of ['project', 'issue-2-non-identity-generator-monad']) {
        const fixture = join(project, 'test-fixtures', name), config = parse(readFileSync(join(fixture, 'spago.yaml'), 'utf8'));
        delete config.workspace.backend;
        config.workspace.extraPackages = Object.fromEntries(campaign.plans.map(plan => [plan.name.slice('purust-'.length), { path: plan.project }]));
        writeFileSync(join(fixture, 'spago.yaml'), stringify(config));
        execute(`${name}-fetch`, campaign.spago, ['fetch', '--offline'], fixture);
        const sources = execute(`${name}-sources`, campaign.spago, ['sources', '--offline', '--json'], fixture);
        execute(`${name}-frontend`, campaign.frontend.path, ['compile', ...JSON.parse(readFileSync(sources.stdout)), '--codegen', 'corefn'], fixture, { GHCRTS: '-N2' });
        const out = join(fixture, 'output/purust_output');
        execute(`${name}-generate`, join(campaign.compilers.purust.directory, 'bin/purust'), ['--source', 'output', '--out', out, '--main', 'Test.Main', '--threaded'], fixture);
        item.fixture_binaries.push(build(out, `${name}-build`));
      }
    }
    const runnerPath = join(project, 'bin/test');
    if (['purust-aff', 'purust-arraybuffer-types', 'purust-argonaut-core', 'purust-argonaut-codecs', 'purust-strings-extra'].includes(result.name)) {
      const record = execute('application', item.binary.path, [], project);
      const stdout = readFileSync(record.stdout, 'utf8'); assert.equal(readFileSync(record.stderr, 'utf8'), '');
      if (result.name === 'purust-aff') assert.equal(stdout, readFileSync(join(project, 'test/expected-main.stdout'), 'utf8'));
      if (result.name === 'purust-argonaut-core') assert(stdout.includes('compact DOM interoperability: Done'), stdout);
      if (result.name === 'purust-argonaut-codecs') assert(stdout.includes('typed plans: Done'), stdout);
      if (result.name === 'purust-strings-extra') for (const marker of [
        'camelCase', 'kebabCase', 'pascalCase', 'snakeCase', 'upperCaseFirst', 'words', 'levenshtein', 'sorensenDiceCoefficient',
      ]) assert(stdout.split('\n').includes(marker), `Missing Strings.Extra group: ${marker}`);
      item.application = record;
    } else {
      const original = readFileSync(runnerPath, 'utf8');
      let start = original.indexOf('\nBINARY=');
      if (['purust-partial', 'purust-unsafe-coerce'].includes(result.name)) start = original.indexOf("\nnode --input-type=module <<'NODE'");
      assert(start >= 0, `No primary test boundary for ${result.name}`);
      let body = original.slice(start + 1);
      if (cuts[result.name]) { const end = body.indexOf(cuts[result.name]); assert(end > 0); body = body.slice(0, end); }
      const script = join(directory, 'primary-test.sh');
      writeFileSync(script, '#!/usr/bin/env bash\nset -euo pipefail\nWITH_INTEGRATION=0\n' + body + '\necho PRIMARY_APPLICATION_PASSED\n');
      item.runner = { original_sha256: hash(original), extracted_sha256: hash(readFileSync(script)),
        scope: 'Existing primary executable checks, including its argument scenarios; separately generated entry points and standalone FFI tests are outside this compilation case.' };
      item.application = execute('application', 'bash', [script], project);
      assert(readFileSync(item.application.stdout, 'utf8').includes('PRIMARY_APPLICATION_PASSED'));
    }
    assert.deepEqual(manifest(canonical), item.generated_manifest);
    item.status = 'passed'; console.log(`${result.name}: primary application PASS`);
  } catch (error) {
    item.status = 'failed'; item.error = error.stack; console.error(`${result.name}: ${error.message}`);
  }
  item.finished_at = new Date().toISOString(); save();
}
state.status = state.results.every(result => result.status === 'passed') ? 'passed' : 'partial';
state.finished_at = new Date().toISOString(); save();
console.log(JSON.stringify({ status: state.status, passed: state.results.filter(result => result.status === 'passed').length,
  failed: state.results.filter(result => result.status !== 'passed').map(result => result.name) }, null, 2));
if (state.status !== 'passed') process.exitCode = 1;
