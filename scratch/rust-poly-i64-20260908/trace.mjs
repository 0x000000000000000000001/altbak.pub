// Reproduce the actual TAST -> PBO -> Rust path without editing either backend.
import assert from 'node:assert/strict';
import { readFileSync, writeFileSync, readdirSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import { spawnSync } from 'node:child_process';

const here = dirname(fileURLToPath(import.meta.url));
const altbak = resolve(here, '../..');
const purust = resolve(altbak, '../purust/purust');
const existing = join(altbak, 'run/bak/rust/output');
const output = join(here, 'tast');
const paths = new Set();
const visited = new Set();
function dependency(name) {
  if (visited.has(name) || name === 'Prim' || name.startsWith('Prim.')) return;
  visited.add(name);
  const json = JSON.parse(readFileSync(join(existing, name, 'corefn.json'), 'utf8'));
  paths.add(resolve(altbak, json.modulePath));
  for (const imported of json.imports) dependency(imported.moduleName.join('.'));
}
dependency('Prelude');
function run(command, args, name) {
  const result = spawnSync(command, args, { cwd: here, encoding: 'utf8', maxBuffer: 32 * 1024 * 1024 });
  writeFileSync(join(here, `${name}.log`), `${result.stdout}\n${result.stderr}`);
  assert.equal(result.status, 0, `${name}: ${result.error ?? ''}\n${result.stderr}`);
}
run(join(altbak, 'run/bak/js/node_modules/.bin/purs'),
  ['compile', ...paths, join(here, 'PolyI64.purs'), '--codegen', 'corefn', '--output', output], 'fixture-compile');
process.chdir(here);
const imp = name => import(pathToFileURL(join(purust, 'output', name, 'index.js')).href);
const [json, cf, sort, fold, maybe, either, set, effect, builder, defaults, directives, foreign] =
  await Promise.all(['PureScript.Backend.Optimizer.CoreFn.Json', 'PureScript.Backend.Optimizer.CoreFn',
    'PureScript.Backend.Optimizer.CoreFn.Sort', 'Data.Foldable', 'Data.Maybe', 'Data.Either',
    'Data.Set', 'Effect', 'PureScript.Backend.Optimizer.Builder',
    'PureScript.Backend.Optimizer.Directives.Defaults', 'PureScript.Backend.Optimizer.Directives',
    'PureScript.Backend.Optimizer.Semantics.Foreign'].map(imp));
const modules = readdirSync(output).filter(n => !n.startsWith('.') && !n.endsWith('.json')).flatMap(name => {
  let data;
  try { data = JSON.parse(readFileSync(join(output, name, 'corefn.json'), 'utf8')); } catch { return []; }
  const decoded = json.decodeModule(data);
  assert.ok(decoded instanceof either.Right, name);
  return [decoded.value0];
});
const names = ['polyLoop', 'intLoop', 'numberLoop'];
const traced = set.fromFoldable(fold.foldableArray)(cf.ordQualified(cf.ordIdent))(
  names.map(name => new cf.Qualified(new maybe.Just('PolyI64'), name)));
function compact(value) {
  if (value == null || typeof value !== 'object') return value;
  if (Array.isArray(value)) return value.map(compact);
  if (value.constructor.name === 'ExprSyntax') return compact(value.value1);
  const tag = value.constructor.name;
  return { ...(tag === 'Object' ? {} : { tag }),
    ...Object.fromEntries(Object.entries(value).map(([k, v]) => [k, compact(v)])) };
}
function nodes(value) {
  if (value == null || typeof value !== 'object') return [];
  return [value, ...Object.values(value).flatMap(nodes)];
}
function hasVariable(value, name) {
  return nodes(value).some(n => n.tag === 'TypeVar' && n.value0 === name);
}
const summary = {};
builder.buildModules(effect.monadEffect)({
  analyzeCustom: () => () => maybe.Nothing.value,
  directives: directives.parseDirectiveFile(defaults.defaultDirectives).directives,
  foreignSemantics: foreign.coreForeignSemantics,
  traceIdents: traced,
  rewriteLimit: 10000,
  onPrepareModule: () => mod => () => mod,
  onSkipModule: () => () => () => maybe.Nothing.value,
  onCodegenModule: () => mod => backend => steps => () => {
    if (mod.name !== 'PolyI64') return;
    writeFileSync(join(here, 'pbo-trace.json'), JSON.stringify(compact({ bindings: backend.bindings, steps }), null, 2) + '\n');
    writeFileSync(join(here, 'decoded-tast.json'), JSON.stringify(compact(mod.decls), null, 2) + '\n');
    for (const step of compact(steps)) {
      const name = step.value0.value1;
      summary[name] = step.value1.map((stage, index) => ({
        stage: index,
        typeApplications: nodes(stage).filter(n => n.tag === 'TypeApp').map(n => n.value1),
        recursiveFunctionTypes: nodes(stage).filter(n => n.tag === 'LetRec').flatMap(rec =>
          nodes(rec.value1).filter(n => n.tag === 'Typed' && n.value1?.tag === 'Abs').map(n => n.value0)),
      }));
      if (name !== 'polyLoop') {
        const expected = name === 'intLoop' ? 'Int' : '$$Number';
        assert.ok(summary[name][0].typeApplications.some(t => t.tag === expected));
        assert.equal(summary[name][1].typeApplications.length, 0);
        assert.ok(summary[name][1].recursiveFunctionTypes.some(t => hasVariable(t, 'a')));
        assert.ok(summary[name].at(-1).recursiveFunctionTypes.some(t => hasVariable(t, 'a')));
      }
    }
    writeFileSync(join(here, 'trace-summary.json'), JSON.stringify(summary, null, 2) + '\n');
    console.log(`PBO trace: ${steps.length} bindings, ${modules.length} modules, ${steps.map(s => s.value1.length).join('/')} stages`);
  },
})(sort.sortModules(fold.foldableArray)(modules))();
run(process.execPath, ['--stack-size=65536', join(purust, 'bin/purust.js'),
  '--source', output, '--out', join(here, 'rust'), '--main', 'PolyI64'], 'fixture-codegen');
writeFileSync(join(here, 'PolyI64-generated.rs'), readFileSync(join(here, 'rust/Purs_PolyI64/src/lib.rs')));
console.log('Fixture compiled and emitted with the normal CLI.');
