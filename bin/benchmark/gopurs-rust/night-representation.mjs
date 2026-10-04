// Qualify the shared-class generator, then isolate it over the IO-pool candidate.
import assert from 'node:assert/strict';
import { chmodSync, copyFileSync, existsSync, globSync, mkdirSync, readFileSync, symlinkSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../..'), purust = join(root, 'purust/purust');
const directory = join(archive, process.argv[3] ?? 'representation');
assert(!existsSync(directory)); mkdirSync(directory);
const tools = join(directory, 'host-tools'); mkdirSync(tools);
const adapter = join(here, 'local-container-command.mjs'); chmodSync(adapter, 0o755); symlinkSync(adapter, join(tools, 'docker'));
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  NIGHT_HOST_ADAPTER_LOGS: join(directory, 'host-adapter-commands'),
  TMPDIR: directory, PATH: [tools, join(archive, 'frontend'), join(purust, 'node_modules/.bin'), process.env.PATH].join(delimiter) };
const state = { status: 'pending', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(directory, 'results.json'), state);
const command = (name, executable, args, cwd = purust, extra = {}) => {
  console.log(name); state.commands.push(run(directory, name, executable, args, cwd, { ...env, ...extra }, 3600000)); save();
};
save();
try {
  state.sources = manifest(join(purust, 'src'));
  command('embed-runtime', process.execPath, ['tools/embed-native-runtime.mjs']);
  command('js-build', 'spago', ['build']);
  command('js-bundle', 'spago', ['bundle', '--module', 'Main', '--platform', 'node', '--outfile', 'bin/purust.js', '--bundle-type', 'app']);
  const generator = join(directory, 'purust.js'); copyFileSync(join(purust, 'bin/purust.js'), generator);
  state.generator_sha256 = hash(readFileSync(generator)); save();
  for (const suite of ['codegen', 'tast']) command(suite, process.execPath,
    ['--test', '--test-concurrency=1', ...globSync(`tests/${suite}/*.mjs`, { cwd: purust }).sort()]);
  if (process.argv.includes('--tests-only')) {
    state.status = 'passed'; state.tests_only = true; state.finished_at = new Date().toISOString(); save();
    process.exit(0);
  }
  const sources = join(directory, 'sources');
  command('freeze-base-sources', 'cp', ['-cR', join(archive, 'candidates/control/sources'), sources], archive);
  for (const name of ['BoundedMemo.rs', 'Monomorphize.rs', 'NativeMaps.rs', 'CoreFn/Usage.rs', 'CoreFn/Json.rs']) {
    const path = 'src/PureScript/Backend/Optimizer/' + name;
    copyFileSync(join(root, 'purescript-backend-optimizer-gopurs', path), join(sources, '1', path));
  }
  command('shared-classes-build', process.execPath, [join(here, 'experiment.mjs'), 'build', archive, 'shared-classes',
    join(archive, 'profile-baseline.json')], archive, {
    GOPURS_EXPERIMENT_GENERATOR: generator,
    GOPURS_EXPERIMENT_SOURCE_ROOTS: JSON.stringify([join(sources, '0'), join(sources, '1')]),
  });
  const selection = join(archive, 'shared-classes-selection.json');
  writeJson(selection, { rounds: 5, variants: [
    { name: 'native-io-pool', binary: 'candidates/native-io-pool/gopurs-rust' },
    { name: 'shared-classes', binary: 'candidates/shared-classes/gopurs-rust' },
    { name: 'go', directory: 'baseline/compiler', environment: {} },
  ] });
  command('compare-shared-classes', process.execPath, [join(here, 'compare.mjs'), resolve(archive, '../gopurs-purust-aff-20261002'),
    join(archive, 'shared-classes-runs'), selection], archive);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
