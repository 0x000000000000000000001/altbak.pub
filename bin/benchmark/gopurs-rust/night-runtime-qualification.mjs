// Complete native Aff/FS/Promise execution after selecting the runtime changes.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../../purust');
const out = join(archive, process.argv[3] ?? 'native-runtime-qualification');
const resume = process.argv[5] === '--resume';
assert(!process.argv[5] || resume);
if (resume) assert(existsSync(out));
else { assert(!existsSync(out)); mkdirSync(out); }
const generator = join(root, 'purust/bin/purust.js'), generatorHash = hash(readFileSync(generator));
const supported = ['purust-aff', 'purust-node-fs', 'purust-js-promise', 'purust-unfoldable'];
const packages = process.argv[4]?.split(',') ?? supported;
assert(packages.length && packages.every(name => supported.includes(name)));
const inputs = () => Object.fromEntries(packages.map(name => [name, manifest(join(root, name, 'src'))]));
const testInputs = () => Object.fromEntries(packages.map(name => [name, {
  tests: manifest(join(root, name, 'test')), runner_sha256: hash(readFileSync(join(root, name, 'bin/test'))),
}]));
const state = resume ? JSON.parse(readFileSync(join(out, 'results.json'))) :
  { status: 'pending', packages, started_at: new Date().toISOString(), generator_sha256: generatorHash,
  sources: inputs(), test_sources: testInputs(), commands: [], generator_host: 'JavaScript', generated_application_host: 'Rust' };
const suffix = resume ? '-resume-' + Date.now() : '';
if (resume) {
  assert.equal(state.status, 'failed');
  assert.equal(state.generator_sha256, generatorHash);
  assert.deepEqual(state.packages, packages);
  assert.deepEqual(state.sources, inputs());
  assert.deepEqual(state.test_sources, testInputs());
  copyFileSync(join(out, 'results.json'), join(out, 'failed-results' + suffix + '.json'));
  state.status = 'pending'; state.resumed_at = new Date().toISOString(); delete state.failure;
}
const save = () => writeJson(join(out, 'results.json'), state);
const env = { ...environment(), PURUST_JS: '1', PURS: join(archive, 'frontend/purs'), GHCRTS: '-N2',
  CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0', CARGO_PROFILE_DEV_DEBUG: '0', CARGO_PROFILE_TEST_DEBUG: '0',
  TMPDIR: out };
save();
try {
  for (const name of packages) {
    console.log(name);
    if (!state.commands.some(record => record.label === name && record.exit_code === 0)) {
      state.commands.push(run(out, name + suffix, join(root, name, 'bin/test'), [], join(root, name), env, 3600000)); save();
    }
    if (name === 'purust-node-fs' || name === 'purust-aff') {
      const entry = name === 'purust-node-fs' ? 'TestAffNative' : 'Test.Main';
      const label = name.replace('purust-', '');
      let module = join(root, name, 'output', entry, 'index.js');
      if (!existsSync(module)) {
        const output = join(root, name, 'output'), js = join(out, label + '-js-output');
        const modules = new Map(readdirSync(output, { withFileTypes: true }).filter(entry => entry.isDirectory())
          .map(entry => join(output, entry.name, 'corefn.json')).filter(existsSync)
          .map(file => { const mod = JSON.parse(readFileSync(file)); return [mod.moduleName.join('.'), mod]; }));
        const visited = new Set(), sources = [];
        const visit = moduleName => {
          if (visited.has(moduleName) || moduleName === 'Prim' || moduleName.startsWith('Prim.')) return;
          visited.add(moduleName);
          const mod = modules.get(moduleName); assert(mod, 'Missing JS oracle dependency ' + moduleName);
          sources.push(resolve(root, name, mod.modulePath));
          for (const dependency of mod.imports) visit(dependency.moduleName.join('.'));
        };
        // Native-only diagnostics in the same workspace need no JS FFI. Compile
        // the actual oracle's dependency closure, not unrelated test modules.
        visit(entry);
        assert(sources.length > 0);
        state.commands.push(run(out, label + '-js-frontend' + suffix, env.PURS,
          ['compile', ...new Set(sources), '--codegen', 'js', '--output', js], join(root, name), env, 300000)); save();
        module = join(js, entry, 'index.js');
      }
      assert(existsSync(module));
      const record = run(out, label + '-js-oracle' + suffix, process.execPath,
        ['--input-type=module', '-e', `import { main } from ${JSON.stringify(pathToFileURL(module).href)}; main();`],
        join(root, name), environment(), 120000);
      if (name === 'purust-node-fs') assert.match(readFileSync(record.stdout, 'utf8'), /FS Aff native: all good/);
      else {
        assert.equal(readFileSync(record.stdout, 'utf8'), readFileSync(join(root, name, 'test/expected-main.stdout'), 'utf8'));
        assert.equal(readFileSync(record.stderr, 'utf8'), '');
      }
      state.commands.push(record); save();
    }
  }
  assert.equal(hash(readFileSync(generator)), generatorHash);
  assert.deepEqual(inputs(), state.sources);
  // FS creates and removes temporary files inside its test tree. Snapshot only
  // after its runner has cleaned them, and require the actual source files to
  // retain their hashes; generated artifacts may have appeared beside them.
  for (const [name, expected] of Object.entries(state.test_sources)) {
    assert.equal(hash(readFileSync(join(root, name, 'bin/test'))), expected.runner_sha256);
    for (const file of expected.tests.filter(file => /\.(purs|rs|js)$/.test(file.path)))
      assert.equal(hash(readFileSync(join(root, name, 'test', file.path))), file.sha256);
  }
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
