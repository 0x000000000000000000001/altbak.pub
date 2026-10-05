// Diagnostic-only compiler: per-module emission and per-round preparation chunks.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, label, sourcesArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && sourcesArg);
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const out = join(archive, label + '-diagnostics.json'); assert(!existsSync(out));
const state = { status: 'running', label, diagnostic_only: true, commands: [] };
const save = () => writeJson(out, state); save();
try {
  state.commands.push(run(archive, label + '-build-diagnostic', process.execPath,
    [join(here, 'build.mjs'), archive, label, resolve(sourcesArg)], archive, environment(), 3600000)); save();
  for (const name of ['gopurs-arrays', 'b8x']) {
    state.commands.push(run(archive, label + '-profile-' + name, 'python3',
      [join(here, 'profile.py'), archive, label + '-' + name, name, join(archive, 'candidates', label, 'gopurs-rust')],
      archive, environment(), 900000)); save();
    const summary = JSON.parse(readFileSync(join(archive, 'profiles', label + '-' + name, 'summary.json')));
    assert.equal(summary.status, 'passed');
    const lines = readFileSync(join(archive, 'profiles', label + '-' + name, 'stderr'), 'utf8');
    const modules = [...lines.matchAll(/^\[gopurs\] diagnostic emit module (.*): (\d+) ms$/gm)]
      .map(([, name, time]) => ({ name, ms: Number(time) })).sort((a, b) => b.ms - a.ms);
    let pending = [];
    const rounds = [];
    for (const line of lines.split('\n')) {
      const chunk = line.match(/^\[gopurs\] diagnostic prepare chunk (\d+) jobs=(\d+): (\d+) ms$/);
      if (chunk) pending.push({ chunk: Number(chunk[1]), jobs: Number(chunk[2]), ms: Number(chunk[3]) });
      const end = line.match(/^\[gopurs\] diagnostic prepare dispatch jobs=(\d+): (\d+) ms$/);
      if (end) { rounds.push({ jobs: Number(end[1]), ms: Number(end[2]), chunks: pending }); pending = []; }
    }
    assert.equal(pending.length, 0);
    const directives = [...lines.matchAll(/^\[gopurs\] diagnostic directive rank=(\d+) before=(\d+) after=(\d+)$/gm)]
      .map(([, rank, before, after]) => ({ rank: Number(rank), before: Number(before), after: Number(after) }));
    const entries = [...lines.matchAll(/^\[gopurs\] diagnostic transitive entry (\S+) key=(\S+) reused=(true|false): (\d+) ms$/gm)]
      .map(([, name, key, reused, time]) => ({ name, key, reused: reused === 'true', ms: Number(time) }));
    const reuse = [...lines.matchAll(/^\[gopurs\] diagnostic transitive results=(\d+) reused=(\d+)$/gm)]
      .map(([, jobs, reused]) => ({ jobs: Number(jobs), reused: Number(reused) }));
    writeJson(join(archive, 'profiles', label + '-' + name, 'tasks.json'), { diagnostic_only: true, modules, rounds, directives, transitive: { entries, reuse } });
  }
  state.status = 'passed'; save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
