// Sample a checked composition serially, outside selection measurements.
import assert from 'node:assert/strict';
import { existsSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, candidate, label = candidate] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(candidate) && /^[a-z0-9-]+$/.test(label));
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const path = join(archive, label + '-reprofile.json'); assert(!existsSync(path));
const result = { status: 'running', candidate, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(path, result); save();
try {
  for (const [name, phase] of [['gopurs-arrays', 'optimize'], ['b8x', 'prepare'], ['gopurs-aff', 'all']]) {
    const profile = label + '-' + name;
    result.commands.push(run(archive, profile, 'python3', [join(here, 'profile.py'), archive, profile, name,
      join(archive, 'candidates', candidate, 'gopurs-rust'), '--phase', phase], archive, environment(), 900000));
    save();
  }
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
