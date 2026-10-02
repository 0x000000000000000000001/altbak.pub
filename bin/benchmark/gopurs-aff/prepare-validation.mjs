import assert from 'node:assert/strict';
import { cpSync, existsSync, mkdirSync, readFileSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, run, writeJson } from './common.mjs';

const workspace = resolve(process.argv[2]), validation = join(workspace, 'portable-validation');
assert(!existsSync(validation)); mkdirSync(validation);
cpSync(join(workspace, 'inputs'), join(validation, 'inputs'), { recursive: true,
  filter: path => !['output', '.purmeta', '.cache'].includes(path.split('/').at(-1)) });
for (const path of ['compilers', 'rust-ffi']) symlinkSync(join(workspace, path), join(validation, path), 'dir');
const htdocs = resolve(dirname(fileURLToPath(import.meta.url)), '../../../..');
const canonicalPath = join(htdocs, 'purust/purust-aff/test/Test/Main.purs');
const canonical = readFileSync(canonicalPath, 'utf8');
writeFileSync(join(validation, 'canonical-Test.Main.purs'), canonical);
const main = join(validation, 'inputs/gopurs-aff/test/Test/Main.purs');
const original = readFileSync(main, 'utf8');
const names = ['test_kill_supervise', 'test_kill_finalizer_catch', 'test_kill_finalizer_bracket', 'test_parallel',
  'test_kill_parallel', 'test_parallel_alt', 'test_parallel_alt_throw', 'test_parallel_alt_sync',
  'test_parallel_mixed', 'test_kill_parallel_alt'];
const definition = (source, name) => {
  const start = source.indexOf(name + ' ::'); assert(start >= 0, name);
  const next = /\ntest_\w+ ::/.exec(source.slice(start + 1)); assert(next, name);
  const end = start + 1 + next.index;
  return source.slice(start, end);
};
let source = original.replace('import Data.Traversable (traverse)', 'import Data.Traversable (traverse, traverse_)')
  .replace('import Effect.Aff.Compat as AC', 'import Effect.Aff.AVar as AVar\nimport Effect.Aff.Compat as AC');
for (const name of names) source = source.replace(definition(source, name), definition(canonical, name));
// The bracket test already awaits completion: inspect the reference then, rather
// than sampling it in a different fiber after a 40ms wall-clock guess.
const bracket = definition(source, 'test_bracket');
assert(bracket.includes('  fiber <- forkAff do\n    delay (Milliseconds 40.0)\n    readRef ref'));
source = source.replace(bracket, bracket.replace('  fiber <- forkAff do\n    delay (Milliseconds 40.0)\n    readRef ref\n', '')
  .replace('  joinFiber fiber <#> eq', '  readRef ref <#> eq'));
writeFileSync(main, source);
const metadata = { original_sha256: hash(original), replacement_sha256: hash(source), canonical_path: canonicalPath,
  canonical_sha256: hash(canonical), copied_tests: names, other_test_change: 'test_bracket: read after bracket completion',
  scope: 'Validation only; the timed corpus and both application libraries remain unchanged.',
  changes: ['AVar synchronization at acquisition/cancellation boundaries', 'join before observing completion',
    'scheduler-independent Alt assertions', 'wider timer margins in inherited tests', 'inspect bracket result after completion'] };
writeJson(join(validation, 'changes.json'), metadata);
const prepared = JSON.parse(readFileSync(join(workspace, 'prepared.json'), 'utf8'));
const env = { ...environment(), PATH: [dirname(prepared.frontend.path), join(htdocs, 'purust/purust/node_modules/.bin'), process.env.PATH ?? ''].join(delimiter) };
const frontend = run(validation, 'frontend', 'spago', ['build'], join(validation, 'inputs/gopurs-aff'), env, 900000);
writeJson(join(validation, 'frontend.json'), frontend);
console.log(validation);
