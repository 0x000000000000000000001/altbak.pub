import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { environment, generate, hash, run, writeJson } from './common.mjs';

const workspace = resolve(process.argv[2]), variant = process.argv[3], label = process.argv[4];
assert(variant && label, 'Usage: qualify.mjs WORKSPACE VARIANT LABEL');
const generation = generate(workspace, variant, label);
const release = process.argv.includes('--release');
const env = { ...environment(), CARGO_PROFILE_DEV_DEBUG: '0', CARGO_INCREMENTAL: '0', CARGO_BUILD_JOBS: '8',
  CARGO_PROFILE_RELEASE_OPT_LEVEL: '3', CARGO_PROFILE_RELEASE_LTO: 'false', CARGO_PROFILE_RELEASE_DEBUG: '0',
  CARGO_TARGET_DIR: join(workspace, 'cargo-target') };
const rust = generation.family === 'purust', cwd = generation.output;
const application = rust ? join(env.CARGO_TARGET_DIR, release ? 'release/purust_output' : 'debug/purust_output') : join(cwd, 'aff-test');
const build = rust
  ? run(workspace, label + '-build', 'cargo', ['build', '--offline', ...(release ? ['--release'] : [])], cwd, env, 900000)
  : run(workspace, label + '-build', 'go', ['build', '-o', application, './main'], cwd, env, 900000);
const execution = run(workspace, label + '-test', application, [], cwd, env, 180000);
const stdout = readFileSync(execution.stdout, 'utf8');
assert.equal(readFileSync(execution.stderr, 'utf8'), '');
assert.deepEqual(stdout.trimEnd().split('\n').sort(), readFileSync(new URL('./expected.stdout', import.meta.url), 'utf8').trimEnd().split('\n').sort());
writeJson(join(workspace, label + '-qualification.json'), { generation, build, execution,
  application, application_sha256: hash(readFileSync(application)), aff_checks: 45, avar_stress_items: 1000, status: 'passed' });
console.log(`${label}: application passed`);
