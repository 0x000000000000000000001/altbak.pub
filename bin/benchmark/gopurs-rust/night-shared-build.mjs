// Measure shared ADT owners separately from the qualified Aff resumption change.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../..'), directory = join(archive, 'shared-build');
assert(!existsSync(directory)); mkdirSync(directory);
const reference = process.argv[3] ?? 'native-aff-io';
assert.equal(JSON.parse(readFileSync(join(archive, 'candidates', reference, 'build.json'))).status, 'passed');
const runtime = JSON.parse(readFileSync(join(archive, 'candidates', reference, 'build.json'))).runtime_sources;
const affSource = runtime.find(source => source.path === join(root, 'purust/purust-aff'));
assert.equal(hash(readFileSync(join(root, 'purust/purust-aff/src/Effect/Aff.rs'))),
  affSource.files.find(file => file.path === 'src/Effect/Aff.rs').sha256);
const checks = JSON.parse(readFileSync(join(archive, 'representation-v3/results.json')));
assert.equal(checks.status, 'passed');
const generator = join(archive, 'representation-v3/purust.js');
assert.equal(hash(readFileSync(generator)), checks.generator_sha256);
const state = { status: 'pending', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(directory, 'results.json'), state);
const command = (label, executable, args, extra = {}) => {
  console.log(label); state.commands.push(run(directory, label, executable, args, archive, { ...environment(), ...extra }, 3600000)); save();
};
save();
try {
  const sources = join(directory, 'sources');
  command('freeze-sources', 'cp', ['-cR', join(archive, 'candidates/control/sources'), sources]);
  for (const name of ['BoundedMemo.rs', 'Monomorphize.rs', 'NativeMaps.rs', 'CoreFn/Usage.rs', 'CoreFn/Json.rs']) {
    const path = 'src/PureScript/Backend/Optimizer/' + name;
    copyFileSync(join(root, 'purescript-backend-optimizer-gopurs', path), join(sources, '1', path));
  }
  command('build', process.execPath, [join(here, 'experiment.mjs'), 'build', archive, 'shared-classes', join(archive, 'profile-baseline.json')], {
    GOPURS_EXPERIMENT_GENERATOR: generator,
    GOPURS_EXPERIMENT_SOURCE_ROOTS: JSON.stringify([join(sources, '0'), join(sources, '1')]),
  });
  const selection = join(archive, 'shared-classes-selection.json');
  writeJson(selection, { rounds: 5, variants: [
    { name: reference, binary: `candidates/${reference}/gopurs-rust` },
    { name: 'shared-classes', binary: 'candidates/shared-classes/gopurs-rust' },
    { name: 'go', directory: 'baseline/compiler', environment: {} },
  ] });
  command('compare', process.execPath, [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, 'shared-classes-runs'), selection]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
