import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, run, writeJson } from '../gopurs-aff/common.mjs';

const [outputArg, outArg] = process.argv.slice(2);
assert(outputArg && outArg, 'check-pbo.mjs COMPILED_JS_OUTPUT NEW_ARCHIVE');
const output = resolve(outputArg), out = resolve(outArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const pbo = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purescript-backend-optimizer-gopurs');
const tests = ['parallel-scheduling', 'parallel-visibility', 'directives-removal', 'implementation-lookup',
  'purmeta-build-cache', 'purmeta-stats', 'typeapp', 'bounded-memo', 'json-fields', 'source-usage',
  'numeric-negate', 'type-table', 'transitive-parallel', 'usage-metadata', 'parallel-load'];
const result = { status: 'pending', output, pbo, runs: [] };
const save = () => writeJson(join(out, 'results.json'), result);
save(); mkdirSync(join(out, 'tests'));
try {
  for (const name of tests) {
    const script = join(pbo, 'test', name + '.mjs');
    copyFileSync(script, join(out, 'tests', name + '.mjs'));
    const record = run(out, name, process.execPath, [script, output], pbo);
    record.test_sha256 = hash(readFileSync(script));
    result.runs.push(record); save(); console.log(`${name}: passed`);
  }
  result.status = 'passed'; save();
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
