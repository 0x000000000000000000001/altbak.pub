// Refresh b8x's frontend from frozen current sources and freeze the current Purust Aff corpus.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync } from 'node:fs';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler, verifyTypedOutput } from '../../../../gopurs/gopurs/tools/native-workspace.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), htdocs = dirname(site);
const archive = resolve(process.argv[2]);
assert(existsSync(join(archive, 'campaign.json')));
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const git = path => ({ directory: path,
  commit: execFileSync('git', ['rev-parse', 'HEAD'], { cwd: path, encoding: 'utf8' }).trim(),
  status: execFileSync('git', ['status', '--short'], { cwd: path, encoding: 'utf8' }) });
const modules = output => readdirSync(output).sort().filter(name => existsSync(join(output, name, 'corefn.json')));
function copySources(source, destination) {
  for (const path of walk(source)) if (/\.(purs|go|js|rs|json|c|h)$/.test(path)) copy(path, join(destination, relative(source, path)));
}
function sourceFreshness(project, output, resolveSource = path => resolve(project, path)) {
  const cachePath = join(output, 'cache-db.json'), cache = JSON.parse(readFileSync(cachePath));
  const records = [];
  for (const [module, inputs] of Object.entries(cache)) for (const [path, [, expected]] of Object.entries(inputs)) {
    const source = resolveSource(path);
    const actual = existsSync(source) ? createHash('sha512').update(readFileSync(source)).digest('hex') : null;
    records.push({ module, path, expected, actual, current: actual === expected });
  }
  return { cache_path: cachePath, cache_sha256: hash(readFileSync(cachePath)), records,
    checked: records.length, changed: records.filter(record => !record.current).length };
}

const b8x = join(htdocs, 'b8x'), b8xOutput = join(b8x, 'run/bak/go/output');
const b8xDirectory = join(archive, 'cases/b8x'), input = join(b8xDirectory, 'input');
assert(!existsSync(b8xDirectory)); mkdirSync(input, { recursive: true });
const freshness = sourceFreshness(b8x, b8xOutput, path => resolve(b8x, path.startsWith('.spago/') ? 'run/bak/go/' + path : path));
writeJson(join(b8xDirectory, 'previous-frontend-freshness.json'), freshness);
copySources(join(b8x, 'src'), join(input, 'src'));
copySources(join(b8x, 'test'), join(input, 'test'));
copy(join(b8x, 'run/bak/go/spago.go.yaml'), join(input, 'spago.yaml'));
copy(join(b8x, 'run/bak/go/spago.lock'), join(input, 'spago.lock'));
const roots = new Map(), standalone = [];
for (const name of modules(b8xOutput)) {
  const metadata = JSON.parse(readFileSync(join(b8xOutput, name, 'corefn.json'))), path = metadata.modulePath;
  if (path.startsWith('src/') || path.startsWith('test/')) continue;
  if (path === '.spago/BuildInfo.purs') {
    copy(join(b8x, 'run/bak/go', path), join(input, path)); standalone.push(path); continue;
  }
  assert(path.startsWith('.spago/p/') || path.startsWith('../gopurs/gopurs-'), path);
  const sourceIndex = path.indexOf('/src/'); assert(sourceIndex > 0, path);
  const prefix = path.slice(0, sourceIndex);
  roots.set(prefix, resolve(b8x, prefix.startsWith('.spago/') ? 'run/bak/go/' + prefix : prefix));
}
for (const [prefix, source] of roots) {
  assert(existsSync(join(source, 'src')), source);
  copySources(join(source, 'src'), resolve(input, prefix, 'src'));
  assert(walk(resolve(input, prefix, 'src')).some(path => path.endsWith('.purs')), prefix);
}
const frozenSources = manifest(b8xDirectory);
const frontend = join(archive, 'frontend/purs');
copy(findTypedCompiler(join(htdocs, 'gopurs/gopurs')), frontend);
const globs = ['src/**/*.purs', 'test/**/*.purs', ...standalone, ...[...roots.keys()].sort().map(prefix => prefix + '/src/**/*.purs')];
const frontendRun = run(b8xDirectory, 'frontend', frontend, ['compile', ...globs, '--codegen', 'corefn'], input,
  { ...environment(), GHCRTS: '-N2' }, 900000);
for (const file of frozenSources) assert.equal(hash(readFileSync(join(b8xDirectory, file.path))), file.sha256);
const tast = verifyTypedOutput(join(input, 'output'), frontend);
const inputManifest = manifest(input);
const siblingInputs = manifest(join(b8xDirectory, 'gopurs')).map(file => ({ ...file, path: 'gopurs/' + file.path }));
const definition = { name: 'b8x', family: 'gopurs', input, input_manifest: inputManifest,
  sibling_inputs: siblingInputs, invocation: [], hosts: ['js', 'go', 'rust'], timeout_ms: 900000,
  source_repository: git(b8x), frontend: { ...frontendRun, compiler_sha256: hash(readFileSync(frontend)), tast },
  scope: 'Full Go-target b8x project, including all current src/test modules and its local/registry dependencies; automatic executable entry-point discovery.' };
writeJson(join(b8xDirectory, 'definition.json'), definition);
console.log(`b8x: fresh ${tast.modules} modules / ${tast.types} types (${freshness.changed} stale frontend inputs refreshed)`);

const project = join(htdocs, 'purust/purust-aff'), output = join(project, 'output');
const directory = join(archive, 'cases/purust-aff'), cwd = join(directory, 'input');
assert(!existsSync(directory)); mkdirSync(cwd, { recursive: true });
const purustFreshness = sourceFreshness(project, output);
writeJson(join(directory, 'frontend-freshness.json'), purustFreshness);
assert.equal(purustFreshness.changed, 0, 'Regenerate stale Purust Aff frontend before timing');
for (const name of modules(output)) {
  const source = join(output, name, 'corefn.json'), metadata = JSON.parse(readFileSync(source));
  copy(source, join(cwd, 'output', name, 'corefn.json'));
  for (const extension of ['purs', 'rs', 'rs.cargo.json']) {
    const path = metadata.modulePath.replace(/\.purs$/, '.' + extension), original = resolve(project, path), target = resolve(cwd, path);
    assert(target.startsWith(directory + '/'), target);
    if (existsSync(original)) copy(original, target);
  }
}
for (const path of ['spago.yaml', 'spago.lock', 'test/expected-main.stdout']) copy(join(project, path), join(cwd, path));
const siblings = manifest(directory).filter(file => !file.path.startsWith('input/') && file.path !== 'frontend-freshness.json');
writeJson(join(directory, 'definition.json'), { name: 'purust-aff', family: 'purust', input: cwd,
  input_manifest: manifest(cwd), sibling_inputs: siblings, invocation: ['--source', 'output', '--main', 'Test.Main', '--threaded'],
  hosts: ['js', 'rust'], source_repository: git(project), frontend: verifyTypedOutput(join(cwd, 'output')),
  source_freshness: { checked: purustFreshness.checked, changed: purustFreshness.changed } });
console.log(`purust-aff: frozen ${modules(output).length} current typed modules`);
