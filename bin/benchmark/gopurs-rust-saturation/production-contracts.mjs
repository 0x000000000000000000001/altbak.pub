// Recheck the changed passes in the production JavaScript module graph. Native
// contracts already attach to the byte-identical generated selection sources.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, manifest } from './common.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../../gopurs/gopurs'), pbo = resolve(root, '../../purescript-backend-optimizer-gopurs');
const selection = JSON.parse(readFileSync(join(archive, 'selection.json')));
const qualification = JSON.parse(readFileSync(join(archive, 'production/results.json')));
assert.equal(selection.status, 'passed'); assert.equal(qualification.status, 'passed');
const directory = join(archive, 'production/contracts'); assert(!existsSync(directory)); mkdirSync(directory);
const output = join(root, 'output'), sourceState = manifest(join(root, 'src'));
const state = { status: 'running', started_at: new Date().toISOString(), output, candidate: selection.candidate,
  qualification_sha256: hash(readFileSync(join(archive, 'production/results.json'))), commands: [] };
const save = () => writeJson(join(directory, 'results.json'), state); save();
const command = (label, args) => {
  state.commands.push(run(directory, label, process.execPath, args, root,
    { ...environment(), GOWORK: 'off', TMPDIR: directory }, 900000)); save();
};
try {
  command('pbo', [join(here, '../gopurs-rust/check-pbo.mjs'), output, join(directory, 'pbo')]);
  for (const name of ['syntax-folds', 'global-spine-annotations', 'monomorphize-cache', 'foreign-specialization']) {
    const path = join(directory, name + '.mjs'); copy(join(pbo, 'test', name + '.mjs'), path);
    command(name, [path, output]);
  }
  for (const name of ['closed-dictionaries.test.mjs', 'go-imports.test.mjs', 'codegen-metadata.mjs'])
    writeFileSync(join(directory, name), readFileSync(join(root, 'tools', name), 'utf8').replaceAll('../output/', output + '/'));
  command('scope-and-imports', ['--test', '--test-concurrency=1', join(directory, 'closed-dictionaries.test.mjs'),
    join(directory, 'go-imports.test.mjs')]);
  const imports = join(directory, 'go-imports-differential.mjs'); copy(join(root, 'tools/go-imports-differential.mjs'), imports);
  const reference = join(directory, 'go-imports-reference.js'); copy(join(archive, 'baseline/go-imports-reference/index.js'), reference);
  command('import-differential', [imports, output, reference, join(directory, 'import-differential')]);
  const mono = join(pbo, 'src/PureScript/Backend/Optimizer/Monomorphize.purs');
  if (readFileSync(mono, 'utf8').includes('specializeTracked')) {
    for (const name of ['specialization-reads.mjs', 'evaluate-tracked-go.mjs', 'evaluate-tracked-go_test.go'])
      copy(join(pbo, 'test', name), join(directory, name));
    command('specialization-reads', [join(directory, 'specialization-reads.mjs'), output, join(directory, 'specialization-reads')]);
    command('tracked-go-bridge', [join(directory, 'evaluate-tracked-go.mjs'), root, mono.replace(/\.purs$/, '.go'), join(directory, 'tracked-go')]);
  }
  assert.deepEqual(manifest(join(root, 'src')), sourceState);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
