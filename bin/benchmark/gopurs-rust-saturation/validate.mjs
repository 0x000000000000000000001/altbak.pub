// Qualify exact measured Go bytes, outside every compiler timing boundary.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, manifest, verifyInput } from './common.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const htdocs = resolve(here, '../../../..');
const measured = JSON.parse(readFileSync(join(archive, 'runs/production-full-table/results.json')));
assert.equal(measured.status, 'passed'); assert.equal(measured.results.length, 51);
const label = process.argv[3] ?? 'applications'; assert(/^[a-z0-9-]+$/.test(label));
const destination = join(archive, 'production', label); assert(!existsSync(destination)); mkdirSync(destination);
const harness = [fileURLToPath(import.meta.url), join(here, 'common.mjs'), join(here, '../gopurs-aff/common.mjs')]
  .map((source, index) => {
    const path = join(destination, 'harness', `${index}.mjs`); copy(source, path);
    return { source, path, sha256: hash(readFileSync(path)) };
  });
const state = { status: 'running', started_at: new Date().toISOString(), harness, applications: [] };
const save = () => writeJson(join(destination, 'results.json'), state); save();
const env = { ...environment(), GOWORK: 'off', GOGC: '1000' };
for (const name of ['gopurs-aff', 'gopurs-arrays', 'gopurs-spec', 'gopurs-yoga-json', 'gopurs-enums',
  'gopurs-js-promise', 'gopurs-prelude', 'gopurs-strings', 'b8x']) {
  const result = { name, status: 'running', commands: [] }; state.applications.push(result); save();
  try {
    const definition = JSON.parse(readFileSync(join(archive, 'cases', name, 'definition.json'))); verifyInput(definition);
    const runs = measured.results.find(item => item.name === name).runs;
    assert(runs.every(run => run.identical_files === definition.oracle.files.length));
    const directory = join(destination, name), generated = join(directory, 'generated');
    for (const file of definition.oracle.files) copy(join(definition.oracle.directory, file.path), join(generated, file.path));
    assert.deepEqual(manifest(generated), definition.oracle.files);
    for (const record of runs) assert.deepEqual(JSON.parse(readFileSync(record.generated_manifest)), definition.oracle.files);
    result.generated_sha256 = hash(JSON.stringify(definition.oracle.files));
    const command = (label, binary, args, cwd, timeout = 300000) => {
      const record = run(directory, label, binary, args, cwd, env, timeout); result.commands.push(record); save(); return record;
    };
    if (name === 'b8x') {
      command('go-mod-tidy', 'go', ['mod', 'tidy'], generated);
      command('go-build', 'go', ['build', './...'], generated, 900000);
      result.validation = 'Every generated Go package and entry point builds.';
    } else {
      const binary = join(directory, 'tests');
      command('go-build', 'go', ['build', '-o', binary, './main'], generated);
      // Tests may create reports or scratch files. Give them private copies so
      // application effects never enter the frozen compiler-input directory.
      const runtime = join(directory, 'runtime'), input = join(runtime, 'input'); mkdirSync(input, { recursive: true });
      for (const file of definition.input_manifest) copy(join(definition.input, file.path), join(input, file.path));
      for (const file of definition.sibling_inputs ?? []) copy(join(dirname(definition.input), file.path), join(runtime, file.path));
      verifyInput({ ...definition, input });
      const before = manifest(runtime);
      const execution = command('go-execute', binary, [], input, 300000);
      const after = manifest(runtime), previous = new Map(before.map(file => [file.path, file]));
      const current = new Map(after.map(file => [file.path, file]));
      result.execution = { directory: input, input_sha256: hash(JSON.stringify(definition.input_manifest)),
        created: after.filter(file => !previous.has(file.path)),
        changed: after.filter(file => previous.has(file.path) && previous.get(file.path).sha256 !== file.sha256),
        removed: before.filter(file => !current.has(file.path)) };
      result.binary_sha256 = hash(readFileSync(binary));
      assert.equal(readFileSync(execution.stderr, 'utf8'), '');
      if (name === 'gopurs-aff') {
        const expected = join(directory, 'expected.stdout');
        copy(join(htdocs, 'gopurs/gopurs-aff/test/expected-main.stdout'), expected);
        const lines = text => text.trimEnd().split(/\r?\n/).sort();
        assert.deepEqual(lines(readFileSync(execution.stdout, 'utf8')), lines(readFileSync(expected, 'utf8')));
        result.checks = lines(readFileSync(expected, 'utf8')).length;
        result.expected_stdout_sha256 = hash(readFileSync(expected));
      }
      if (name === 'gopurs-yoga-json') {
        const path = join(input, '.spec-results'), tests = JSON.parse(readFileSync(path));
        assert(tests.length > 0); assert(tests.every(([, result]) => result.success === true));
        result.spec_results = { path, sha256: hash(readFileSync(path)), passed: tests.length };
      }
    }
    const after = manifest(generated);
    assert.deepEqual(after.filter(file => file.path.endsWith('.go')),
      definition.oracle.files.filter(file => file.path.endsWith('.go')), 'Application qualification changed generated Go');
    result.module_files_after_build = after.filter(file => ['go.mod', 'go.sum'].includes(file.path));
    verifyInput(definition); result.status = 'passed';
  } catch (error) { result.status = 'failed'; result.failure = error.stack; }
  result.finished_at = new Date().toISOString(); save();
}
state.status = state.applications.every(item => item.status === 'passed') ? 'passed' : 'failed';
state.finished_at = new Date().toISOString(); save();
assert.equal(state.status, 'passed');
