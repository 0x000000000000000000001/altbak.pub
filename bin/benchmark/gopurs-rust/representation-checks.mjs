// Resume only the two truncated host-adapter checks, then run the TAST suite.
import { globSync, readFileSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
import assert from 'node:assert/strict';
const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const purust = resolve(here, '../../../../purust/purust'), directory = join(archive, 'representation-v2');
const prior = JSON.parse(readFileSync(join(directory, 'results.json')));
assert.equal(hash(readFileSync(join(purust, 'bin/purust.js'))), prior.generator_sha256);
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  TMPDIR: directory, NIGHT_HOST_ADAPTER_LOGS: join(directory, 'host-adapter-commands'),
  PATH: [join(directory, 'host-tools'), join(archive, 'frontend'), join(purust, 'node_modules/.bin'), process.env.PATH].join(delimiter) };
const state = { status: 'pending', prior: join(directory, 'results.json'), generator_sha256: prior.generator_sha256, commands: [] };
const save = () => writeJson(join(directory, 'checks-resumed.json'), state);
save();
try {
  for (const [label, paths] of [
    ['codegen-resumed', ['tests/codegen/datetime-instant-ffi.mjs', 'tests/codegen/uuid-ffi.mjs']],
    ['tast', globSync('tests/tast/*.mjs', { cwd: purust }).sort()],
  ]) {
    console.log(label);
    state.commands.push(run(directory, label, process.execPath, ['--test', '--test-concurrency=1', ...paths], purust, env, 3600000)); save();
  }
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
