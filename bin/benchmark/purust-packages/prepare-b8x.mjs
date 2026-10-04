// Freeze the existing complete Rust test profile, selecting its real import closure.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, globSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { basename, delimiter, dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';
import { verifyTypedOutput } from '../../../../purust/purust/tools/native-workspace.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), htdocs = dirname(site);
const archive = resolve(process.argv[2]), campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const root = join(htdocs, 'b8x'), profile = join(root, 'run/bak/rust');
const project = join(archive, 'b8x-project'), directory = join(archive, 'cases/b8x'), input = join(directory, 'input');
assert(!existsSync(project)); assert(!existsSync(directory)); mkdirSync(project); mkdirSync(input, { recursive: true });
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const sourceInputs = [];
for (const part of ['src', 'test']) for (const path of walk(join(root, part))) {
  if (!/\.(purs|rs|js|json)$/.test(path)) continue;
  const destination = join(project, 'b8x', part, relative(join(root, part), path));
  copy(path, destination); sourceInputs.push({ original: path, destination, sha256: hash(readFileSync(path)) });
}
for (const file of manifest(join(profile, 'ffi'))) copy(join(profile, 'ffi', file.path), join(project, 'ffi', file.path));
copy(join(profile, 'src/DefaultMain.purs'), join(project, 'src/DefaultMain.purs'));
for (const name of ['profile.mjs', 'shared.mjs']) copy(join(profile, 'driver', name), join(directory, 'original-driver', name));
const { parse, stringify } = createRequire(join(htdocs, 'purust/purust/package.json'))('yaml');
const original = parse(readFileSync(join(profile, 'spago.yaml'), 'utf8'));
const config = { package: original.package, workspace: {
  packageSet: original.workspace.packageSet,
  extraPackages: Object.fromEntries(campaign.plans.map(plan => [plan.name.slice('purust-'.length), { path: relative(project, plan.project) }])) } };
writeFileSync(join(project, 'spago.yaml'), stringify(config));
const sourceManifest = manifest(project);
const env = { ...environment(), GHCRTS: '-N2', PATH: [dirname(campaign.frontend.path), dirname(campaign.spago), process.env.PATH].join(delimiter) };
const commands = [];
const execute = (label, command, args) => {
  const result = run(directory, label, command, args, project, env, 900000); commands.push(result);
  writeJson(join(directory, 'frontend-commands.json'), commands); return result;
};
execute('fetch', campaign.spago, ['fetch', '--offline']);
const sourcesRun = execute('sources', campaign.spago, ['sources', '--offline', '--json']);
const globs = JSON.parse(readFileSync(sourcesRun.stdout));
const { candidateClosure, selectSources } = await import(pathToFileURL(join(profile, 'driver/profile.mjs')));
const candidates = [...new Set([
  ...globSync(globs.map(pattern => resolve(project, pattern))),
  ...globSync(join(project, 'b8x/src/**/*.purs')), join(project, 'b8x/test/Util/Assert.purs'),
])].sort();
const main = 'Test.Rust.Main';
const graphRun = execute('graph', campaign.frontend.path, ['graph', ...candidateClosure(candidates, main)]);
const selected = selectSources(JSON.parse(readFileSync(graphRun.stdout)), 'default', project);
execute('frontend', campaign.frontend.path, ['compile', ...selected.map(source => source.path), '--codegen', 'corefn', '--output', join(project, 'output')]);
const tast = verifyTypedOutput(join(project, 'output'), campaign.frontend.path);
assert.equal(tast.modules, selected.length);
for (const file of sourceManifest) assert.equal(hash(readFileSync(join(project, file.path))), file.sha256);
// Resolve fallback/override FFI before relocation; each selected file becomes adjacent to its module.
process.chdir(project);
const { findFfiFileImpl } = await import(pathToFileURL(join(htdocs, 'purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/FfiSupport.js')));
const resolver = findFfiFileImpl('.rs')([])('ffi'), foreign = [];
for (const source of selected) {
  const path = join(project, 'output', source.module, 'corefn.json'), data = JSON.parse(readFileSync(path));
  const relocated = `sources/${source.module}/${basename(source.path)}`;
  copy(source.path, join(input, relocated));
  const ffi = resolver(source.module)(source.path)();
  if (ffi) {
    copy(ffi, join(input, relocated.replace(/\.purs$/, '.rs')));
    if (existsSync(ffi + '.cargo.json')) copy(ffi + '.cargo.json', join(input, relocated.replace(/\.purs$/, '.rs.cargo.json')));
  }
  foreign.push({ module: source.module, source: source.path, ffi, ffi_sha256: ffi ? hash(readFileSync(ffi)) : null,
    foreign: data.foreign, original_corefn_sha256: hash(readFileSync(path)) });
  data.modulePath = relocated; mkdirSync(join(input, 'output', source.module), { recursive: true });
  writeFileSync(join(input, 'output', source.module, 'corefn.json'), JSON.stringify(data) + '\n');
}
writeJson(join(directory, 'source-provenance.json'), { source_inputs: sourceInputs, frozen_project: sourceManifest, selected, foreign });
writeJson(join(directory, 'definition.json'), { name: 'b8x', family: 'purust', input, input_manifest: manifest(input), sibling_inputs: [],
  invocation: ['--source', 'output', '--main', main, '--threaded'], hosts: ['js', 'rust'], timeout_ms: 900000,
  source_corpus: project, scope: 'Existing Rust default test profile: Core, Infra and Util (286 tests), selected import closure; distinct from the full Go-target corpus.',
  frontend: { ...tast, commands }, expected_tests: 286, main });
console.log(`b8x Rust profile: ${tast.modules} modules / ${tast.types} types, ${foreign.filter(item => item.ffi).length} frozen native FFI files`);
