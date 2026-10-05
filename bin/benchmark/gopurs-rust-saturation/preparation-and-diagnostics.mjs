import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const candidate = process.argv[3] ?? 'dynamic-preparation';
assert(/^[a-z0-9-]+$/.test(candidate));
assert.equal(JSON.parse(readFileSync(join(archive, 'runs/module-range-v2/results.json'))).status, 'passed');
const out = join(archive, candidate + '-and-diagnostics.json'); assert(!existsSync(out));
const state = { status: 'running', commands: [] };
const save = () => writeJson(out, state); save();
try {
  for (const [label, script, args] of [
    [candidate + '-screening', 'candidate.mjs', [archive, candidate, 'module-range-v2', join(archive, 'queued', candidate)]],
    ['diagnostic-transitive', 'diagnose.mjs', [archive, 'diagnostic-transitive', join(archive, 'queued/diagnostic-transitive')]],
  ]) {
    state.commands.push(run(archive, label, process.execPath, [join(here, script), ...args], archive, environment(), 7200000)); save();
  }
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
