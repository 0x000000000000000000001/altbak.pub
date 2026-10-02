import assert from 'node:assert/strict';
import { spawnSync, execFileSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { chmodSync, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { delimiter, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { findTypedCompiler, verifyTypedOutput } from '../../../../purust/purust/tools/native-workspace.mjs';

const htdocs = resolve(dirname(fileURLToPath(import.meta.url)), '../../../..');
const workspace = resolve(process.argv[2]);
assert(!existsSync(workspace), `Workspace already exists: ${workspace}`);
mkdirSync(workspace, { recursive: true });
const inputs = join(workspace, 'inputs'); mkdirSync(inputs);
const root = join(htdocs, 'gopurs/gopurs-aff');
const lock = JSON.parse(readFileSync(join(root, 'spago.lock'), 'utf8'));
const packages = new Set(['gopurs-aff']);
for (const pkg of Object.values(lock.packages)) {
  if (pkg.type === 'local') packages.add(dirname(join(root, pkg.path)).endsWith('/gopurs') ? pkg.path.split('/').at(-1) : '');
}
assert(!packages.has(''), 'Unexpected local dependency path');
for (const name of packages) {
  const source = join(htdocs, 'gopurs', name), target = join(inputs, name);
  mkdirSync(target);
  cpSync(join(source, 'src'), join(target, 'src'), { recursive: true, filter: path => !path.endsWith('/.DS_Store') });
  for (const file of ['spago.yaml', 'spago.lock', 'package.json']) {
    if (existsSync(join(source, file))) copyFileSync(join(source, file), join(target, file));
  }
  if (name === 'gopurs-aff') cpSync(join(source, 'test'), join(target, 'test'), { recursive: true });
}
const cwd = join(inputs, 'gopurs-aff');
// The package's Go configuration has no backend hook: generate TAST once.
copyFileSync(join(root, 'spago.go.yaml'), join(cwd, 'spago.yaml'));
const hash = path => createHash('sha256').update(readFileSync(path)).digest('hex');
const git = directory => ({ directory,
  revision: execFileSync('git', ['rev-parse', 'HEAD'], { cwd: directory, encoding: 'utf8' }).trim(),
  status: execFileSync('git', ['status', '--porcelain=v1'], { cwd: directory, encoding: 'utf8' }).trim() });
const artifacts = {};
for (const [family, compiler] of [['gopurs', 'gopurs/gopurs'], ['purust', 'purust/purust']]) {
  const directory = join(htdocs, compiler), bin = join(workspace, 'compilers', family, 'bin');
  mkdirSync(bin, { recursive: true });
  artifacts[family] = { checkout: git(directory), files: [] };
  for (const name of [family, `${family}.js`, `${family}-native`]) {
    const source = join(directory, 'bin', name), target = join(bin, name);
    copyFileSync(source, target); chmodSync(target, 0o755);
    artifacts[family].files.push({ name, path: target, origin: source, sha256: hash(target) });
  }
  if (family === 'gopurs') {
    mkdirSync(join(dirname(bin), 'tools'));
    for (const name of ['ffi-runner.mjs', 'wasm_exec.js', 'ffi_gen.wasm']) {
      const source = join(directory, 'tools', name), target = join(dirname(bin), 'tools', name);
      copyFileSync(source, target);
      artifacts[family].files.push({ name: 'tools/' + name, path: target, origin: source, sha256: hash(target) });
    }
    writeFileSync(join(dirname(bin), 'package.json'), '{"type":"module"}\n');
  }
}
const purs = findTypedCompiler(join(htdocs, 'purust/purust'));
const env = { ...process.env, PATH: [dirname(purs), join(htdocs, 'purust/purust/node_modules/.bin'), process.env.PATH ?? ''].join(delimiter) };
for (const key of Object.keys(env)) if (/^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)/.test(key)) delete env[key];
const result = spawnSync('spago', ['build'], { cwd, env, encoding: 'utf8', timeout: 900000, maxBuffer: 32 * 1024 * 1024 });
writeFileSync(join(workspace, 'frontend.log'), (result.stdout ?? '') + (result.stderr ?? ''));
assert.equal(result.status, 0, result.error?.message ?? result.stderr);
const tast = verifyTypedOutput(join(cwd, 'output'));
const modules = readdirSync(join(cwd, 'output')).filter(name => existsSync(join(cwd, 'output', name, 'corefn.json'))).sort();
const files = modules.map(name => {
  const path = join(cwd, 'output', name, 'corefn.json'), bytes = readFileSync(path), value = JSON.parse(bytes);
  return { module: name, path, bytes: bytes.length, sha256: hash(path), source: value.modulePath,
    foreign: value.foreign.map(value => typeof value === 'string' ? value : value.identifier ?? value.name ?? value) };
});
const provenance = { created_at: new Date().toISOString(), workspace, project: git(root),
  frontend: { path: purs, sha256: hash(purs) }, artifacts, packages: [...packages].sort(), tast: { ...tast,
    bytes: files.reduce((sum, file) => sum + file.bytes, 0), sha256: createHash('sha256').update(files.map(file => `${file.sha256}  ${file.module}/corefn.json\n`).join('')).digest('hex'), files } };
writeFileSync(join(workspace, 'prepared.json'), JSON.stringify(provenance, null, 2) + '\n');
console.log(JSON.stringify({ workspace, packages: packages.size, tast: { ...tast, bytes: provenance.tast.bytes, sha256: provenance.tast.sha256 } }, null, 2));
