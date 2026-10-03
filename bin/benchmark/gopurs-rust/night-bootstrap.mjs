// Fresh, retained production hosts after branch synchronization.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

assert(process.argv[2], 'night-bootstrap.mjs NEW_ARCHIVE');
const archive = resolve(process.argv[2]);
assert(!existsSync(archive)); mkdirSync(archive, { recursive: true });
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../gopurs/gopurs');
const pbo = resolve(root, '../../purescript-backend-optimizer-gopurs');
const purust = resolve(root, '../../purust/purust');
const sources = () => ({ gopurs: manifest(join(root, 'src')), optimizer: manifest(join(pbo, 'src')),
  purust: manifest(join(purust, 'src')) });
const state = { status: 'pending', started_at: new Date().toISOString(), sources: sources(),
  previous_executables: Object.fromEntries(['gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, hash(readFileSync(join(root, 'bin', name)))])),
  commands: [] };
const save = () => writeJson(join(archive, 'bootstrap.json'), state);
const env = { ...environment(), GHCRTS: '-N2', CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0',
  CARGO_NET_OFFLINE: 'true', GOPURS_KEEP_WORKSPACE: '1', GOPURS_NATIVE_TMPDIR: archive, TMPDIR: archive };
copyFileSync(fileURLToPath(import.meta.url), join(archive, 'night-bootstrap.mjs'));
save();
try {
  for (const [label, command, args, cwd] of [
    ['activity-before', 'ps', ['-axo', 'pid,ppid,pcpu,etime,command'], root],
    ['revision-gopurs', 'git', ['rev-parse', 'HEAD'], root],
    ['revision-pbo', 'git', ['rev-parse', 'HEAD'], pbo],
    ['revision-purust', 'git', ['rev-parse', 'HEAD'], purust],
    ['go-js', 'npm', ['run', 'build:native', '--', '--keep-workspace'], root],
    ['rust', 'npm', ['run', 'build:rust', '--', '--keep-workspace'], root],
  ]) {
    state.commands.push(run(archive, label, command, args, cwd, env, 3600000)); save();
  }
  state.go_bootstrap = readFileSync(state.commands.find(r => r.label === 'go-js').stdout, 'utf8')
    .match(/^Native bootstrap workspace: (.+)$/m)?.[1];
  state.rust_bootstrap = readFileSync(state.commands.find(r => r.label === 'rust').stdout, 'utf8')
    .match(/^Rust-hosted gopurs workspace: (.+)$/m)?.[1];
  assert(state.go_bootstrap && state.rust_bootstrap);
  assert.deepEqual(sources(), state.sources);
  state.executables = Object.fromEntries(['gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, { path: join(root, 'bin', name), sha256: hash(readFileSync(join(root, 'bin', name))) }]));
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
  console.log(JSON.stringify({ status: state.status, go: state.go_bootstrap, rust: state.rust_bootstrap,
    executables: state.executables }, null, 2));
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
