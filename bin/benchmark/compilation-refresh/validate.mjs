// Build and execute representative generated applications after all timing runs.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), htdocs = dirname(site);
const archive = resolve(process.argv[2]);
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const groups = ['libraries', 'extra'].map(name => JSON.parse(readFileSync(join(archive, `${name}-results.json`))));
for (const group of groups) assert.equal(group.status, 'passed');
const cases = groups.flatMap(group => group.results);
const destination = join(archive, 'validation'); assert(!existsSync(destination)); mkdirSync(destination);
const state = { status: 'running', started_at: new Date().toISOString(), applications: [] };
const save = () => writeJson(join(destination, 'results.json'), state);
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const env = { ...environment(), GOWORK: 'off', GOGC: '1000', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_DEV_DEBUG: '0' };
save();
for (const name of ['gopurs-aff', 'gopurs-arrays', 'gopurs-enums', 'gopurs-js-promise', 'gopurs-prelude', 'gopurs-strings', 'b8x', 'purust-aff']) {
  const measured = cases.find(item => item.name === name); assert(measured);
  const definition = JSON.parse(readFileSync(measured.definition));
  const directory = join(destination, name), generated = join(directory, 'generated'); mkdirSync(directory);
  const original = measured.runs[0].retained_output, files = JSON.parse(readFileSync(measured.runs[0].generated_manifest));
  for (const file of files) {
    assert.equal(hash(readFileSync(join(original, file.path))), file.sha256);
    copy(join(original, file.path), join(generated, file.path));
  }
  assert.deepEqual(manifest(generated), files);
  const result = { name, status: 'running', generated_sha256: measured.runs[0].generated_sha256, commands: [] };
  state.applications.push(result); save();
  const execute = (label, command, args, cwd, timeout = 300000) => {
    const record = run(directory, label, command, args, cwd, env, timeout); result.commands.push(record); save(); return record;
  };
  try {
    if (definition.family === 'gopurs') {
      if (name === 'b8x') {
        execute('go-mod-tidy', 'go', ['mod', 'tidy'], generated);
        execute('go-build', 'go', ['build', './...'], generated, 900000);
        result.validation = 'All generated packages and executable entry points build.';
      } else {
        const binary = join(directory, 'tests');
        execute('go-build', 'go', ['build', '-o', binary, './main'], generated);
        const application = execute('go-execute', binary, [], definition.input, 120000);
        assert.equal(readFileSync(application.stderr, 'utf8'), '');
        result.binary_sha256 = hash(readFileSync(binary));
        if (name === 'gopurs-aff') {
          const expected = join(directory, 'expected.stdout');
          copy(join(htdocs, 'gopurs/gopurs-aff/test/expected-main.stdout'), expected);
          const lines = text => text.trimEnd().split(/\r?\n/).sort();
          assert.deepEqual(lines(readFileSync(application.stdout, 'utf8')), lines(readFileSync(expected, 'utf8')));
          result.expected_stdout_sha256 = hash(readFileSync(expected)); result.aff_checks = 45;
        }
      }
    } else {
      execute('cargo-build', 'cargo', ['build', '--offline', '--manifest-path', join(generated, 'Cargo.toml')], definition.input, 900000);
      const binary = join(generated, 'target/debug/purust_output'), retained = join(directory, 'aff-tests');
      copy(binary, retained);
      const application = execute('aff-execute', retained, [], definition.input, 120000);
      assert.equal(readFileSync(application.stderr, 'utf8'), '');
      const expected = join(definition.input, 'test/expected-main.stdout');
      assert.equal(readFileSync(application.stdout, 'utf8'), readFileSync(expected, 'utf8'));
      result.binary_sha256 = hash(readFileSync(retained)); result.expected_stdout_sha256 = hash(readFileSync(expected));
      result.aff_checks = readFileSync(expected, 'utf8').trim().split('\n').length;
    }
    result.status = 'passed';
  } catch (error) {
    result.status = 'failed'; result.error = error.stack; console.error(`${name}: ${error.message}`);
  }
  result.finished_at = new Date().toISOString(); save();
  console.log(`${name}: ${result.status}`);
}
for (const compiler of Object.values(campaign.compilers)) {
  assert.deepEqual(manifest(compiler.directory), compiler.files);
  for (const file of compiler.files) assert.equal(hash(readFileSync(join(compiler.origin, file.path))), file.sha256);
}
state.compilers_unchanged = true;
state.status = state.applications.every(result => result.status === 'passed') ? 'passed' : 'failed';
state.finished_at = new Date().toISOString(); save();
if (state.status !== 'passed') process.exitCode = 1;
