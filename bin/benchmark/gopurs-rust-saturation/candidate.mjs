// Snapshot first, then build, check, and screen one candidate without concurrent
// CPU-heavy work. Later live edits cannot alter this candidate's build inputs.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, run, writeJson } from '../gopurs-aff/common.mjs';
import { copy, hash, manifest } from './common.mjs';

const [archiveArg, label, reference = 'control', sourceBase] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && /^[a-z0-9-]+$/.test(reference));
const archive = resolve(archiveArg), here = dirname(fileURLToPath(import.meta.url));
const htdocs = resolve(here, '../../../..');
const liveRoots = [join(htdocs, 'gopurs/gopurs'), join(htdocs, 'purescript-backend-optimizer-gopurs')];
const roots = sourceBase ? [0, 1].map(index => resolve(sourceBase, String(index))) : liveRoots;
const proposal = join(archive, 'proposals', label); assert(!existsSync(proposal)); mkdirSync(proposal, { recursive: true });
const sources = roots.map((root, index) => {
  const files = [...manifest(join(root, 'src')).map(file => ({ ...file, path: 'src/' + file.path })),
    ...['spago.yaml', 'spago.lock', 'package.json'].filter(name => existsSync(join(root, name)))
      .map(path => ({ path, sha256: hash(readFileSync(join(root, path))) }))];
  for (const file of files) copy(join(root, file.path), join(proposal, 'sources', String(index), file.path));
  return { root, files };
});
const state = { status: 'running', label, reference, sources, started_at: new Date().toISOString(), commands: [] };
const save = () => writeJson(join(proposal, 'results.json'), state);
const command = (name, args) => {
  state.commands.push(run(proposal, name, process.execPath, args, archive, environment(), 3600000)); save();
};
save();
try {
  command('build', [join(here, 'build.mjs'), archive, label, join(proposal, 'sources')]);
  if (readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/Monomorphize.purs'), 'utf8')
    .includes('Reuse Tuple and the existing prepared-record field set')) {
    const files = [reference, label].map(name => {
      const path = join(archive, 'candidates', name, 'rust/purust_core/src/lib.rs');
      return { candidate: name, path, sha256: hash(readFileSync(path)) };
    });
    state.runtime_layout = { files, identical: files[0].sha256 === files[1].sha256 }; save();
    assert(state.runtime_layout.identical, 'Layout-control hypothesis must preserve the common generated runtime byte-for-byte');
  }
  const folds = join(liveRoots[1], 'test/syntax-folds.mjs');
  if (existsSync(folds)) {
    const frozen = join(proposal, 'syntax-folds.mjs'); copy(folds, frozen);
    command('syntax-folds', [frozen, join(archive, 'work/output')]);
  }
  if (readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/NativeMaps.purs'), 'utf8').includes('foreign import filterRangeImpl')) {
    for (const name of ['module-directive-range.mjs', 'native-range.mjs', 'native-range.rs'])
      copy(join(liveRoots[1], 'test', name), join(proposal, name));
    command('module-directive-range', [join(proposal, 'module-directive-range.mjs'), join(archive, 'work/output')]);
    command('native-range', [join(proposal, 'native-range.mjs'), join(archive, 'candidates', label, 'rust'), join(proposal, 'native-range')]);
  }
  if (!readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/Monomorphize.purs'))
    .equals(readFileSync(join(archive, 'candidates', reference, 'sources/1/src/PureScript/Backend/Optimizer/Monomorphize.purs')))) {
    for (const name of ['global-spine-annotations', 'global-type-quantifiers', 'static-args', 'static-argument-keys',
      'static-dictionaries', 'monomorphize-transitive', 'monomorphize-callsite', 'monomorphize-cache',
      'foreign-specialization', 'spine-annotations', 'local-specializations', 'binding-order']) {
      const frozen = join(proposal, name + '.mjs'); copy(join(liveRoots[1], 'test', name + '.mjs'), frozen);
      command(name, [frozen, join(archive, 'work/output')]);
    }
  }
  if (readFileSync(join(proposal, 'sources/0/src/Gopurs/Preparation.purs'), 'utf8').includes("Ref.modify'")) {
    for (const name of ['preparation.test.mjs', 'monomorphization.test.mjs']) {
      const frozen = join(proposal, name);
      writeFileSync(frozen, readFileSync(join(liveRoots[0], 'tools', name), 'utf8').replaceAll('../output/', join(archive, 'work/output') + '/'));
      command(name, ['--test', '--test-timeout=60000', frozen]);
    }
    for (const name of ['preparation-native.mjs', 'preparation-native.rs']) copy(join(liveRoots[0], 'tools', name), join(proposal, name));
    command('native-preparation', [join(proposal, 'preparation-native.mjs'), join(archive, 'candidates', label, 'rust'), join(proposal, 'native-preparation')]);
  }
  if (readFileSync(join(proposal, 'sources/0/src/Gopurs/ClosedDictionaries/Scope.purs'), 'utf8').includes('closedIn')) {
    for (const name of ['closed-dictionaries.test.mjs', 'codegen-metadata.mjs']) {
      writeFileSync(join(proposal, name), readFileSync(join(liveRoots[0], 'tools', name), 'utf8')
        .replaceAll('../output/', join(archive, 'work/output') + '/'));
    }
    command('closed-dictionaries', ['--test', '--test-timeout=60000', join(proposal, 'closed-dictionaries.test.mjs')]);
  }
  if (readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/Monomorphize.purs'), 'utf8').includes('specializeTracked')) {
    const frozen = join(proposal, 'specialization-reads.mjs'); copy(join(liveRoots[1], 'test/specialization-reads.mjs'), frozen);
    command('specialization-reads', [frozen, join(archive, 'work/output'), join(proposal, 'specialization-reads')]);
    for (const name of ['specialization-reads-native.mjs', 'specialization-reads-native.rs'])
      copy(join(liveRoots[1], 'test', name), join(proposal, name));
    command('specialization-reads-native', [join(proposal, 'specialization-reads-native.mjs'),
      join(archive, 'candidates', label, 'rust'), join(proposal, 'specialization-reads-native')]);
  }
  if (readFileSync(join(proposal, 'sources/0/src/Gopurs/GoImports.purs'), 'utf8').includes('type Imports = Set String')) {
    writeFileSync(join(proposal, 'go-imports.test.mjs'), readFileSync(join(liveRoots[0], 'tools/go-imports.test.mjs'), 'utf8')
      .replaceAll('../output/', join(archive, 'work/output') + '/'));
    command('go-imports', ['--test', join(proposal, 'go-imports.test.mjs')]);
    const frozen = join(proposal, 'go-imports-differential.mjs');
    copy(join(liveRoots[0], 'tools/go-imports-differential.mjs'), frozen);
    const reference = join(proposal, 'go-imports-reference.js');
    copy(join(archive, 'baseline/go-imports-reference/index.js'), reference);
    command('go-imports-differential', [frozen, join(archive, 'work/output'), reference, join(proposal, 'go-imports-differential')]);
  }
  if (readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/NativeMaps.purs'), 'utf8').includes('foreign import foldrIntSuffixImpl')) {
    for (const name of ['int-suffix.mjs', 'native-int-suffix.mjs', 'native-int-suffix.rs'])
      copy(join(liveRoots[1], 'test', name), join(proposal, name));
    command('int-suffix', [join(proposal, 'int-suffix.mjs'), join(archive, 'work/output')]);
    command('native-int-suffix', [join(proposal, 'native-int-suffix.mjs'),
      join(archive, 'candidates', label, 'rust'), join(proposal, 'native-int-suffix')]);
  }
  if (readFileSync(join(proposal, 'sources/1/src/PureScript/Backend/Optimizer/Monomorphize.purs'), 'utf8').includes('mangleTypeChunks')) {
    const frozen = join(proposal, 'mangle-chunks.mjs'); copy(join(liveRoots[1], 'test/mangle-chunks.mjs'), frozen);
    const reference = join(proposal, 'mangle-reference.js'); copy(join(archive, 'baseline/mangle-reference/index.js'), reference);
    command('mangle-chunks', [frozen, join(archive, 'work/output'), reference, join(proposal, 'mangle-chunks')]);
  }
  const configuration = { rounds: 5, cases: ['gopurs-arrays', 'b8x', 'gopurs-aff', 'gopurs-spec', 'gopurs-yoga-json'],
    variants: [{ name: reference, binary: `candidates/${reference}/gopurs-rust` },
      { name: label, binary: `candidates/${label}/gopurs-rust` }] };
  const configPath = join(proposal, 'screening.json'); writeJson(configPath, configuration);
  command('measure', [join(here, 'measure.mjs'), archive, label, configPath]);
  state.status = 'passed'; state.finished_at = new Date().toISOString(); save();
} catch (error) { state.status = 'failed'; state.failure = error.stack; save(); throw error; }
