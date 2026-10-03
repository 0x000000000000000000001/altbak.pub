// Run the production JS/codegen/TAST regressions after the qualified JS build.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [qualificationArg, outArg] = process.argv.slice(2);
assert(qualificationArg && outArg, 'regressions.mjs QUALIFICATION NEW_ARCHIVE');
const qualification = JSON.parse(readFileSync(join(resolve(qualificationArg), 'qualification.json'), 'utf8'));
assert.equal(qualification.status, 'passed');
const root = qualification.root, out = resolve(outArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
assert.equal(hash(readFileSync(join(root, 'bin/purust.js'))), qualification.javascript.sha256);
assert.equal(hash(readFileSync(qualification.stage2.path)), qualification.stage2.sha256);
const tests = ['codegen', 'tast'].flatMap(name => manifest(join(root, 'tests', name))
  .map(file => ({ ...file, path: 'tests/' + name + '/' + file.path })));
for (const file of tests) {
  const destination = join(out, file.path); mkdirSync(dirname(destination), { recursive: true });
  copyFileSync(join(root, file.path), destination);
}
// These existing b8x FFI integration fixtures require the named Linux container.
// Keep their sources and the failed discovery run; the compiler's native FFI,
// portable codegen/TAST checks and complete Aff suite run on the host.
const unavailable = ['tests/codegen/crypto-hash-ffi.mjs', 'tests/codegen/datetime-instant-ffi.mjs',
  'tests/codegen/foreign-object-foldm.mjs', 'tests/codegen/uuid-ffi.mjs', 'tests/tast/crypto-hash.mjs'];
const result = { status: 'pending', qualified_binary: qualification.stage2, tests, runs: [],
  unavailable: unavailable.map(path => ({ path, reason: 'Requires core-api-cli-1; Docker daemon unavailable on this host' })) };
const save = () => writeJson(join(out, 'results.json'), result);
save();
const env = { ...environment(), PURUST_NATIVE: qualification.stage2.path, PURUST_PURS: qualification.frontend.path,
  PURS: qualification.frontend.path, TMPDIR: out,
  PURUST_JS: '0', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_DEV_DEBUG: '0',
  CARGO_PROFILE_TEST_DEBUG: '0', CARGO_NET_OFFLINE: 'true', GHCRTS: '-N2', PURUST_NATIVE_TMPDIR: out };
try {
  for (const suite of ['codegen', 'tast']) {
    const prefix = 'tests/' + suite + '/';
    const files = tests.filter(f => f.path.startsWith(prefix) && f.path.endsWith('.mjs') && !f.path.slice(prefix.length).includes('/') && !unavailable.includes(f.path))
      .map(f => join(root, f.path));
    const record = run(out, suite, process.execPath, ['--test', '--test-concurrency=1', ...files], root, env, 1800000);
    result.runs.push(record); save();
    assert.match(readFileSync(record.stdout, 'utf8'), /(?:#|ℹ) fail 0\b/);
  }
  for (const file of tests) assert.equal(hash(readFileSync(join(root, file.path))), file.sha256, file.path);
  result.status = 'passed'; save();
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
