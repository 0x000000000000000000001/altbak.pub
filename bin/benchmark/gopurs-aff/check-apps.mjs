import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { environment, hash, writeJson } from './common.mjs';

const workspace = resolve(process.argv[2]), label = process.argv[3], apps = process.argv.slice(4).map(path => resolve(path));
assert(label && apps.length, 'Usage: check-apps.mjs WORKSPACE LABEL APPLICATION...');
const expected = readFileSync(new URL('./expected.stdout', import.meta.url), 'utf8').trimEnd().split('\n').sort();
const runs = [], directory = join(workspace, label); mkdirSync(directory);
// Fixed diagnostic batch, including all failures; never retry until success.
for (let round = 1; round <= 5; round++) {
  for (let index = 0; index < apps.length; index++) {
    const application = apps[index], prefix = join(directory, `${round}-${index}`);
    const run = spawnSync(application, [], { cwd: dirname(application), env: environment(), encoding: 'utf8', timeout: 120000 });
    writeFileSync(prefix + '.stdout', run.stdout ?? ''); writeFileSync(prefix + '.stderr', run.stderr ?? '');
    const checks = (run.stdout ?? '').trimEnd().split('\n').sort();
    const passed = run.status === 0 && run.stderr === '' && JSON.stringify(checks) === JSON.stringify(expected);
    runs.push({ round, index, application, application_sha256: hash(readFileSync(application)), exit_code: run.status,
      signal: run.signal, error: run.error?.message, passed, checks: checks.length, stdout: prefix + '.stdout', stderr: prefix + '.stderr' });
    writeJson(join(directory, 'results.json'), runs);
    console.log(`${round}-${index}: ${passed ? 'passed' : 'failed'} (${checks.length} printed checks)`);
  }
}
console.log(JSON.stringify({ passed: runs.filter(run => run.passed).length, total: runs.length }));
