// Serialize a frozen list of experiments. Failures retain their own evidence;
// the next independent hypothesis may still run when continue_on_failure is set.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy } from './common.mjs';
const [archiveArg, label, configArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && configArg);
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const raw = readFileSync(resolve(configArg)), config = JSON.parse(raw);
const out = join(archive, label + '-sequence.json'); assert(!existsSync(out));
copy(resolve(configArg), join(archive, label + '-sequence-config.json'));
const state = { status: 'running', started_at: new Date().toISOString(), configuration_sha256: hash(raw), commands: [] };
const save = () => writeJson(out, state); save();
try {
  for (const command of config.commands) {
    const entry = { ...command, status: 'running' }; state.commands.push(entry); save();
    try {
      entry.result = run(archive, label + '-' + command.label, command.program ?? process.execPath,
        [join(here, command.script), archive, ...command.args.map(arg => arg.replaceAll('$ARCHIVE', archive))],
        archive, environment(), 10800000);
      entry.status = 'passed';
    } catch (error) {
      entry.status = 'failed'; entry.failure = error.stack; save();
      if (!config.continue_on_failure) throw error;
    }
    save();
  }
  state.status = state.commands.every(entry => entry.status === 'passed') ? 'passed' : 'completed-with-failures';
  state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
