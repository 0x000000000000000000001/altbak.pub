import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from './common.mjs';

const [compilerArg, affArg, bootstrapArg, directoryArg] = process.argv.slice(2);
assert(compilerArg && affArg && bootstrapArg && directoryArg,
  'qualify-hosts.mjs GOPURS_CHECKOUT AFF_CHECKOUT RUST_BOOTSTRAP NEW_DIRECTORY');
const compiler = resolve(compilerArg), aff = resolve(affArg), bootstrap = resolve(bootstrapArg), directory = resolve(directoryArg);
assert(!existsSync(directory)); mkdirSync(directory, { recursive: true });
const frontend = JSON.parse(readFileSync(join(bootstrap, 'verify-tast.json'), 'utf8'));
const purust = resolve(compiler, '../../purust/purust');
const env = { ...environment(), GHCRTS: '-N2', PURS: frontend.compiler, CARGO_NET_OFFLINE: 'true',
  TMPDIR: directory, GOPURS_TEST_KEEP_WORKSPACE: '1' };
const result = { status: 'pending', started_at: new Date().toISOString(), compiler, aff, bootstrap, frontend,
  executables: Object.fromEntries(['gopurs.js', 'gopurs-native', 'gopurs-rust'].map(name =>
    [name, { path: join(compiler, 'bin', name), sha256: hash(readFileSync(join(compiler, 'bin', name))) }])),
  sources: { gopurs: manifest(join(compiler, 'src')), optimizer: manifest(resolve(compiler, '../../purescript-backend-optimizer-gopurs/src')),
    aff: manifest(join(aff, 'test')) }, runs: [] };
const save = () => writeJson(join(directory, 'results.json'), result);
const check = (label, command, args, cwd, overrides = {}) => {
  const record = run(directory, label, command, args, cwd, { ...env, ...overrides }, 900000);
  result.runs.push(record); save(); return record;
};
function copy(source, destination) { mkdirSync(dirname(destination), { recursive: true }); copyFileSync(source, destination); }
const generated = output => manifest(output, path => path.endsWith('.go') || path.endsWith('/go.mod'));
const generatedRust = output => manifest(output, path => (path.endsWith('.rs') || path.endsWith('/Cargo.toml'))
  && !path.endsWith('/Purs_Gopurs_FfiSupport/build.rs'));
const sortedLines = text => text.trimEnd().split(/\r?\n/).sort();
for (const path of ['qualify-hosts.mjs', 'common.mjs']) copy(join(dirname(fileURLToPath(import.meta.url)), path), join(directory, 'harness', path));
save();
try {
  const pgoPath = join(bootstrap, 'pgo-profile.json');
  const pgo = existsSync(pgoPath) ? JSON.parse(readFileSync(pgoPath)) : null;
  if (pgo) {
    assert.equal(pgo.status, 'passed');
    assert.equal(pgo.binary.sha256, result.executables['gopurs-rust'].sha256);
  }
  assert.equal(hash(readFileSync(resolve(bootstrap, pgo?.binary.path ?? 'target/release/purust_output'))), result.executables['gopurs-rust'].sha256);
  assert.equal(readFileSync(join(bootstrap, 'smoke-go-run.log'), 'utf8').trim(), 'Done');
  result.bootstrap = { directory: bootstrap, frontend,
    binary_sha256: result.executables['gopurs-rust'].sha256,
    native_generation: generatedRust(join(bootstrap, 'rust')),
    purust_native_sha256: hash(readFileSync(join(purust, 'bin/purust-native'))),
    purust_js_sha256: hash(readFileSync(join(purust, 'bin/purust.js'))),
    pgo: pgo ? { path: pgoPath, sha256: hash(readFileSync(pgoPath)), binary_sha256: pgo.binary.sha256,
      profile_sha256: pgo.profile.merged.sha256, training_sha256: pgo.training.sha256 } : null,
  };
  const javascript = join(directory, 'bootstrap-js');
  check('bootstrap-js', join(purust, 'bin/purust'), ['--source', join(bootstrap, 'output'), '--out', javascript,
    '--main', 'Main', '--threaded'], bootstrap, { PURUST_JS: '1' });
  assert.deepEqual(generatedRust(javascript), result.bootstrap.native_generation, 'JS/native Purust generated different gopurs Rust');
  result.bootstrap.identical_files = result.bootstrap.native_generation.length;
  save();

  let reference, inputs;
  const expected = sortedLines(readFileSync(join(aff, 'test/expected-main.stdout'), 'utf8'));
  for (const host of ['rust', 'go', 'js']) {
    const record = check('aff-' + host, './bin/test', [], aff,
      { GOPURS_RUST: host === 'rust' ? '1' : '0', GOPURS_JS: host === 'js' ? '1' : '0' });
    const output = join(aff, 'output'), files = generated(output);
    const tast = manifest(output, path => path.endsWith('/corefn.json'));
    const saved = join(directory, 'aff-' + host);
    for (const file of files) copy(join(output, file.path), join(saved, file.path));
    for (const path of ['test.stdout', 'test.stderr', 'go_test_app']) copy(join(output, path), join(saved, path));
    if (!inputs) {
      inputs = tast;
      for (const file of tast) copy(join(output, file.path), join(directory, 'aff-tast', file.path));
    } else assert.deepEqual(tast, inputs, 'Live Aff frontend inputs changed between hosts');
    assert.deepEqual(sortedLines(readFileSync(join(output, 'test.stdout'), 'utf8')), expected);
    assert.equal(readFileSync(join(output, 'test.stderr'), 'utf8'), '');
    if (reference) assert.deepEqual(files, reference, 'Live Aff generated different Go');
    else reference = files;
    Object.assign(record, { host, go_output: saved, identical_files: files.length, tast_modules: tast.length,
      checks: expected.length, avar_stress_items: 1000, generated: files });
    save();
    console.log(`${host}: ${expected.length} Aff checks + AVar stress, ${files.length} exact Go files`);
  }
  check('compiler-tests', 'npm', ['run', 'test:rust'], compiler);
  check('parser-tests', 'go', ['test', '-race', '-tags=carchive', '-count=1', './...'], join(compiler, 'tools/ffi-gen'));
  assert.deepEqual(manifest(join(aff, 'test')), result.sources.aff);
  assert.equal(hash(readFileSync(join(compiler, 'bin/gopurs-rust'))), result.executables['gopurs-rust'].sha256);
  result.status = 'passed'; result.finished_at = new Date().toISOString(); save();
  console.log(`Qualification passed: ${directory}`);
} catch (error) { result.status = 'failed'; result.failure = error.stack; save(); throw error; }
