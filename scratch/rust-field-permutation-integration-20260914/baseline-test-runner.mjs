import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';

const scratch = fileURLToPath(new URL('./', import.meta.url));
const root = '/Users/0x1/Documents/htdocs/purust/purust';
const purs = '/Users/0x1/Documents/htdocs/purescript/.stack-work/dist/aarch64-osx/ghc-9.8.4/build/purs/purs';
const directory = join(scratch, 'build/baseline-test-evidence');
mkdirSync(directory, { recursive: true });
const hash = path => createHash('sha256').update(readFileSync(path)).digest('hex');
const cases = ['shared-nullaries', 'record-root-move', 'foreign-object'];
const results = [];
for (const name of cases) {
  const test = join(root, 'tests/tast', `${name}.mjs`);
  const args = ['--import', join(scratch, 'baseline-test-hook.mjs'), test];
  const run = spawnSync(process.execPath, args, { cwd: scratch, encoding: 'utf8', timeout: 180_000,
    maxBuffer: 32 * 1024 * 1024, env: { ...process.env, PURS: purs, PURUST_BASELINE_TEST: name } });
  const log = join(directory, `${name}.log`);
  writeFileSync(log, `${run.stdout ?? ''}\n${run.stderr ?? ''}`);
  const redirects = readFileSync(join(directory, name, 'redirects.jsonl'), 'utf8').trim().split('\n').map(JSON.parse);
  assert.ok(redirects.length > 0 && redirects.every(row => row.status === 0), `${name}: backend actually redirected and succeeded`);
  const result = { name, test, testSha256: hash(test), command: process.execPath, args, cwd: scratch,
    purs, status: run.status, signal: run.signal, error: run.error?.message ?? null, log, redirects };
  results.push(result);
  writeFileSync(join(scratch, 'baseline-test-results.json'), JSON.stringify(results, null, 2) + '\n');
  console.log(JSON.stringify({ name, status: run.status, log,
    assertion: (run.stderr ?? '').match(/(?:AssertionError[^\n]*|Unknown fallback shape[^\n]*|left: \d+|right: \d+)/g) }));
}
