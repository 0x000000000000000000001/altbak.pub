import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const here = dirname(fileURLToPath(import.meta.url)), archive = resolve(process.argv[2]);
const output = join(archive, 'start-results.json'); assert(!existsSync(output));
const reclamation = JSON.parse(readFileSync(join(archive, 'purust-reclamation.json')));
assert.equal(reclamation.status, 'passed');
const state = { status: 'running', started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(output, state);
const command = (label, executable, args) => {
  state.commands.push(run(archive, label, executable, args, archive, environment(), 1800000)); save();
};
save();
try {
  writeJson(join(archive, 'profile-baseline.json'), { lto: 'thin', linker: 'rust-lld' });
  command('freeze-compilers', process.execPath, [join(here, '../gopurs-rust/experiment.mjs'), 'init', archive,
    'e3fa820c4e78fb0f0c5f4d8d7ed91314c73e5b6c0e0b7e3b0e780ec8401a9aa1', join(archive, 'profile-baseline.json')]);
  command('freeze-cases', process.execPath, [join(here, 'freeze.mjs'), archive]);
  for (const name of ['gopurs-arrays', 'b8x', 'gopurs-aff'])
    command('profile-' + name, 'python3', [join(here, 'profile.py'), archive, 'baseline-' + name, name,
      join(archive, 'baseline/compiler/bin/gopurs-rust')]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
