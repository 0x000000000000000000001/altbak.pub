// Serialized native differential checks with retained compiler/test provenance.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [rustArg, outArg, configArg] = process.argv.slice(2);
assert(rustArg && outArg && configArg, 'check-native.mjs GENERATED_RUST NEW_ARCHIVE CHECKS.json');
const rust = resolve(rustArg), out = resolve(outArg), configPath = resolve(configArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust/purust');
const checks = JSON.parse(readFileSync(configPath, 'utf8'));
assert(Array.isArray(checks));
const generated = () => manifest(rust, path => !path.includes('/target/') && /\.(rs|toml)$/.test(path));
const result = { status: 'pending', rust, generated: generated(), checks, runs: [] };
const save = () => writeJson(join(out, 'results.json'), result);
const tmp = join(out, 'workspaces'); mkdirSync(tmp);
mkdirSync(join(out, 'tests'));
const env = { ...environment(), PURUST_NATIVE_TMPDIR: tmp, CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0' };
save();
try {
  for (const check of checks) {
    assert(/^test-native-[a-z-]+$/.test(check.name));
    const script = join(root, 'tools', check.name + '.mjs');
    const testSource = join(root, 'tools', check.name + '.rs');
    const inputs = [script, ...(existsSync(testSource) ? [testSource] : [])].map(path => ({ path, sha256: hash(readFileSync(path)) }));
    for (const input of inputs) copyFileSync(input.path, join(out, 'tests', basename(input.path)));
    const record = run(out, check.name, process.execPath, [script, rust,
      ...(check.args ?? []).map(path => resolve(dirname(configPath), path))], root, env, 1800000);
    for (const input of inputs) assert.equal(hash(readFileSync(input.path)), input.sha256);
    result.runs.push({ ...record, inputs }); save();
    console.log(`${check.name}: passed`);
  }
  assert.deepEqual(generated(), result.generated);
  result.status = 'passed'; save();
} catch (error) { result.status = 'failed'; result.error = error.stack; save(); throw error; }
