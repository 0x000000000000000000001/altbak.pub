import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { existsSync, mkdirSync, readFileSync, readdirSync, renameSync, rmSync, statSync, writeFileSync } from 'node:fs';
import { dirname, join, relative } from 'node:path';
import { loadavg } from 'node:os';

export const hash = bytes => createHash('sha256').update(bytes).digest('hex');
export function walk(directory) {
  if (!existsSync(directory)) return [];
  return readdirSync(directory, { withFileTypes: true }).sort((a, b) => a.name < b.name ? -1 : 1).flatMap(entry => {
    const path = join(directory, entry.name);
    return entry.isDirectory() ? walk(path) : entry.isFile() ? [path] : [];
  });
}
export const manifest = (directory, filter = () => true) => walk(directory).filter(filter).map(path =>
  ({ path: relative(directory, path), bytes: statSync(path).size, sha256: hash(readFileSync(path)) }));
export const writeJson = (path, value) => writeFileSync(path, JSON.stringify(value, null, 2) + '\n');
export function environment() {
  const env = { ...process.env };
  for (const key of Object.keys(env)) if (/^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)/.test(key)) delete env[key];
  return env;
}
export function run(workspace, label, command, args, cwd, env = environment(), timeout = 300000) {
  const logs = join(workspace, 'logs'); mkdirSync(logs, { recursive: true });
  const load = loadavg(), started = performance.now();
  const result = spawnSync(command, args, { cwd, env, encoding: 'utf8', timeout, maxBuffer: 64 * 1024 * 1024 });
  const record = { label, command: [command, ...args], cwd, exit_code: result.status, signal: result.signal,
    wall_ms: performance.now() - started, load_average_before: load,
    explicit_environment: Object.fromEntries(Object.entries(env).filter(([key]) => /^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)/.test(key))),
    stdout: join(logs, label + '.stdout'), stderr: join(logs, label + '.stderr') };
  writeFileSync(record.stdout, result.stdout ?? ''); writeFileSync(record.stderr, result.stderr ?? '');
  writeJson(join(logs, label + '.json'), record);
  assert.ifError(result.error);
  assert.equal(result.status, 0, `Failed ${label}; see ${record.stderr}\n${result.stderr?.slice(-8000)}`);
  return record;
}
export function generate(workspace, variant, label) {
  const family = variant.startsWith('gopurs') ? 'gopurs' : 'purust', js = variant.endsWith('-js');
  const cwd = join(workspace, 'inputs/gopurs-aff'), output = join(cwd, 'output');
  const out = join(workspace, 'generated', label);
  assert(!existsSync(out), out); mkdirSync(dirname(out), { recursive: true });
  for (const directory of ['.purmeta', '.cache']) rmSync(join(cwd, directory), { recursive: true, force: true });
  if (family === 'gopurs') {
    for (const path of ['purescript', 'gopurs_runtime', 'main', 'Test.Main/main', 'go.mod', 'go.sum']) {
      rmSync(join(output, path), { recursive: true, force: true });
    }
  }
  const args = ['--main', 'Test.Main'];
  // The JS FFI resolver joins this option to cwd, so pass a relative directory
  // that resolves identically in JavaScript and Rust.
  if (family === 'purust') args.push('--threaded', '--source', 'output', '--out', out, '--ffi-dir', relative(cwd, join(workspace, 'rust-ffi')));
  const env = { ...environment(), [family.toUpperCase() + '_JS']: js ? '1' : '0' };
  const record = run(workspace, label, join(workspace, 'compilers', family, 'bin', family), args, cwd, env);
  if (family === 'gopurs') {
    mkdirSync(out);
    for (const path of ['purescript', 'gopurs_runtime', 'main', 'Test.Main/main', 'go.mod']) {
      assert(existsSync(join(output, path)), path); mkdirSync(dirname(join(out, path)), { recursive: true });
      renameSync(join(output, path), join(out, path));
    }
  }
  const stderr = readFileSync(record.stderr, 'utf8');
  record.phases_ms = Object.fromEntries([...stderr.matchAll(new RegExp(`^\\[${family}\\] (.+): (\\d+) ms$`, 'gm'))].map(([, name, ms]) => [name, Number(ms)]));
  assert(record.phases_ms['backend total'] > 0, stderr);
  assert.equal([...stderr.matchAll(new RegExp(`^\\[${family}\\] backend total:`, 'gm'))].length, 1);
  if (family === 'purust') assert(readFileSync(record.stdout, 'utf8').includes('Successfully generated Rust code.'));
  const files = manifest(out, path => /\.(rs|go|toml)$/.test(path) || path.endsWith('/go.mod'));
  Object.assign(record, { variant, family, output: out, generated_files: files.length,
    generated_sha256: hash(JSON.stringify(files)), generated_manifest: join(workspace, 'logs', label + '.files.json') });
  writeJson(record.generated_manifest, files);
  writeJson(join(workspace, 'logs', label + '.json'), record);
  console.log(`${label}: ${record.phases_ms['backend total']} ms; ${files.length} files`);
  return record;
}
