// Serialized qualification follow-up for the allocation campaign.
import assert from 'node:assert/strict';
import { readdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, mode] = process.argv.slice(2);
assert(archiveArg && ['selection', 'confirm'].includes(mode),
  'finish-allocation.mjs ARCHIVE selection|confirm');
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../../gopurs/gopurs');
const snapshot = resolve(archive, '../gopurs-purust-aff-20261002');
const env = { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', GHCRTS: '-N2',
  TMPDIR: archive, PBO_QUALIFIED_WORKSPACE: join(archive, 'qualified-adapters') };
const state = { status: 'pending', mode, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(archive, `finish-${mode}.json`), state);
function command(label, executable, args, cwd = root) {
  console.log(label);
  state.commands.push(run(archive, label, executable, args, cwd, env, 3600000)); save();
}
save();
try {
  if (mode === 'selection') {
    assert.equal(JSON.parse(readFileSync(join(archive, 'native-tests/summary.json'))).status, 'ok');
    command('qualified-adapters', process.execPath,
      [resolve(root, '../../purescript-backend-optimizer-gopurs/test/corefn-qualified-native.mjs'), root]);
    command('pbo-tests', process.execPath, [join(here, 'check-pbo.mjs'), join(archive, 'work/output'), join(archive, 'pbo-tests')]);
    command('qualified-selection', process.execPath,
      [join(here, 'compare.mjs'), snapshot, join(archive, 'qualified-runs'), join(archive, 'qualified.json')]);
  } else {
    assert.equal(JSON.parse(readFileSync(join(archive, 'production/results.json'))).status, 'passed');
    for (const [name, rounds, defaults, resources, variants] of [
      ['primary', 15, false, false, ['rust-before', 'rust']],
      ['common', 10, false, false, ['js', 'go', 'rust-before', 'rust']],
      ['default', 5, true, false, ['rust-before', 'rust', 'go', 'js']],
      ['resources', 3, false, true, ['rust-before', 'rust']],
    ]) {
      const selection = { rounds, launcherDefaults: defaults, resources, variants: variants.map(name => ({
        name, directory: name === 'rust-before' ? 'baseline/compiler' : root,
        environment: name.startsWith('rust') ? { GOPURS_RUST: '1' } : name === 'js' ? { GOPURS_JS: '1' } : {},
      })) };
      writeJson(join(archive, name + '.json'), selection);
      command('activity-before-' + name, 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
      command('compare-' + name, process.execPath,
        [join(here, 'compare.mjs'), snapshot, join(archive, name + '-runs'), join(archive, name + '.json')]);
      command('activity-after-' + name, 'ps', ['-axo', 'pid,ppid,pcpu,etime,command']);
      // Reflink complete immutable artifacts between campaigns, outside clocks.
      command('deduplicate-' + name, 'python3', [join(here, 'deduplicate.py'),
        join(archive, 'deduplication-' + name + '.json'), join(archive, 'baseline'),
        join(archive, 'native-delete-runs'), join(archive, name + '-runs')]);
    }
    for (const phase of ['all', 'optimize']) {
      const output = join(archive, 'profile-final-' + phase);
      command('profile-final-' + phase, 'python3', [join(here, 'profile.py'),
        join(root, 'bin/gopurs-rust'), snapshot, output, '--phase', phase]);
      command('attribute-final-' + phase, 'python3', [join(here, 'attribute.py'), join(output, 'sample.txt')]);
    }
    command('verify-final', process.execPath, [join(here, 'verify.mjs'), join(archive, 'verification-final.json'),
      ...readdirSync(archive, { withFileTypes: true })
        .filter(entry => entry.isDirectory() && entry.name.endsWith('-runs'))
        .map(entry => join(archive, entry.name)).sort()]);
  }
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
