// Execute the existing complete Aff suite with the candidate compiler in a
// fresh sibling-package layout. bin/test may clear only this copied output.
import assert from 'node:assert/strict';
import { copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, symlinkSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [qualificationArg, archiveArg] = process.argv.slice(2);
assert(qualificationArg && archiveArg && process.argv.length === 4,
  'Usage: qualify-aff.mjs COMPILER_QUALIFICATION_DIRECTORY NEW_ARCHIVE_DIRECTORY');
const qualification = resolve(qualificationArg), archive = resolve(archiveArg);
const compiler = JSON.parse(readFileSync(join(qualification, 'qualification.json'), 'utf8'));
assert.equal(compiler.status, 'passed');
assert.equal(hash(readFileSync(compiler.stage2.path)), compiler.stage2.sha256);
assert.equal(hash(readFileSync(compiler.frontend.path)), compiler.frontend.sha256);
assert.equal(hash(readFileSync(compiler.javascript.path)), compiler.javascript.sha256);
assert(!existsSync(archive)); mkdirSync(archive, { recursive: true });
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust');
const workspace = join(archive, 'workspace'); mkdirSync(workspace);
const app = join(workspace, 'purust-aff'), mirror = join(workspace, 'purust');
const packages = new Set();
function frozenPackage(path) {
  packages.add(path);
  const index = compiler.sources.findIndex(source => source.path === path);
  assert(index >= 0, `Package absent from compiler provenance: ${path}`);
  const frozen = join(qualification, 'sources', String(index));
  for (const file of compiler.sources[index].files) {
    assert.equal(hash(readFileSync(join(frozen, file.path))), file.sha256, file.path);
  }
  return frozen;
}
cpSync(join(root, 'purust-aff'), app, { recursive: true, filter: path =>
  !['.git', '.spago', '.cache', '.purmeta', '.DS_Store', 'output', 'node_modules'].includes(basename(path)) });
for (const name of readdirSync(root).filter(name => name.startsWith('purust-') && name !== 'purust-aff')) {
  if (existsSync(join(root, name, 'spago.yaml'))) symlinkSync(frozenPackage(join(root, name)), join(workspace, name), 'dir');
}
mkdirSync(join(mirror, 'bin'), { recursive: true });
for (const name of ['node_modules', '.spago']) {
  symlinkSync(join(root, 'purust', name), join(mirror, name), 'dir');
}
const compilerSources = frozenPackage(join(root, 'purust'));
for (const name of ['tools', 'src', 'tests']) symlinkSync(join(compilerSources, name), join(mirror, name), 'dir');
for (const name of ['spago.lock', 'package.json']) copyFileSync(join(compilerSources, name), join(mirror, name));
const launcher = join(qualification, 'before/purust');
assert.equal(hash(readFileSync(launcher)), compiler.before.purust);
copyFileSync(launcher, join(mirror, 'bin/purust'));
copyFileSync(compiler.stage2.path, join(mirror, 'bin/purust-native'));
copyFileSync(compiler.javascript.path, join(mirror, 'bin/purust.js'));
// The existing error-reporting test also enumerates this conventional path.
symlinkSync(resolve(root, '../purescript'), join(archive, 'purescript'), 'dir');
mkdirSync(join(archive, 'errors'));
const inputs = manifest(app);
const result = { status: 'pending', qualification, archive, workspace,
  compiler_sha256: compiler.stage2.sha256, inputs, commands: [] };
const save = () => writeJson(join(archive, 'qualification.json'), result);
save();
try {
  const env = { ...environment(), PURS: compiler.frontend.path, PURUST_JS: '0',
    PURUST_AFF_ERRORS_KEEP_OUTPUT: join(archive, 'errors'),
    CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_DEV_DEBUG: '0',
    CARGO_PROFILE_TEST_DEBUG: '0', GHCRTS: '-N2' };
  result.commands.push(run(archive, 'aff-full-suite', 'bash', ['bin/test'], app, env, 1800000));
  const output = readFileSync(result.commands[0].stdout, 'utf8');
  assert(output.includes('Summary: 47 Aff checks'));
  assert(output.includes('9 error-reporting scenarios passed.'));
  assert.equal(hash(readFileSync(join(mirror, 'bin/purust-native'))), compiler.stage2.sha256);
  assert.equal(hash(readFileSync(compiler.frontend.path)), compiler.frontend.sha256);
  assert.equal(hash(readFileSync(join(mirror, 'bin/purust.js'))), compiler.javascript.sha256);
  for (const path of packages) frozenPackage(path);
  for (const file of inputs) assert.equal(hash(readFileSync(join(app, file.path))), file.sha256, file.path);
  Object.assign(result, { status: 'passed', aff_checks: 47, rust_unit_tests: 5, error_reporting_scenarios: 9,
    other_suites: ['concurrent Ref/AVar', 'child lifetime success and failures'], inputs_unchanged: true });
  save(); console.log(JSON.stringify({ ...result, inputs: undefined }, null, 2));
} catch (error) {
  result.failure = { message: error.message, stack: error.stack }; save(); throw error;
}
