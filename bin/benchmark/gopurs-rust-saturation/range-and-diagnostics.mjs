import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]), here = dirname(fileURLToPath(import.meta.url));
const candidate = process.argv[3] ?? 'module-range';
assert(/^[a-z0-9-]+$/.test(candidate));
assert.equal(JSON.parse(readFileSync(join(archive, 'runs/bulk-directives/results.json'))).status, 'passed');
const out = join(archive, candidate === 'module-range' ? 'range-and-diagnostics.json' : candidate + '-and-diagnostics.json'); assert(!existsSync(out));
const state = { status: 'running', commands: [] };
const save = () => writeJson(out, state); save();
try {
  for (const [label, script, args] of [
    [candidate + '-screening', 'candidate.mjs', [archive, candidate, 'bulk-directives', join(archive, 'queued', candidate)]],
    ['diagnostic-folds', 'diagnose.mjs', [archive, 'diagnostic-folds', join(archive, 'queued/diagnostic-folds')]],
  ]) {
    state.commands.push(run(archive, label, process.execPath, [join(here, script), ...args], archive, environment(), 7200000)); save();
  }
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
