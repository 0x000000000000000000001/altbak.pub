// Compile one frozen package frontend, then relocate its TAST and native FFI for timing.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { basename, delimiter, dirname, join, relative, resolve } from 'node:path';
import { environment, hash, manifest, run, walk, writeJson } from '../gopurs-aff/common.mjs';
import { verifyTypedOutput } from '../../../../purust/purust/tools/native-workspace.mjs';

const [archiveArg, name, caseLabel = name] = process.argv.slice(2), archive = resolve(archiveArg);
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json'))), plan = campaign.plans.find(plan => plan.name === name); assert(plan);
// Isolated workspace YAML intentionally has no backend. On a retry, preserve
// the original runner's invocation from its frozen plan instead of re-inferring it.
let invocationProvenance = null;
if (campaign.baseline_archive) {
  const path = join(campaign.baseline_archive, 'campaign.json'), baseline = JSON.parse(readFileSync(path));
  const original = baseline.plans.find(item => item.name === name); assert(original);
  invocationProvenance = { path, sha256: hash(readFileSync(path)), main: original.main, flags: original.flags };
  plan.main = original.main; plan.flags = original.flags;
}
const directory = join(archive, 'cases', caseLabel), input = join(directory, 'input');
assert(!existsSync(directory)); mkdirSync(input, { recursive: true });
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
const env = { ...environment(), GHCRTS: '-N2',
  PATH: [dirname(campaign.frontend.path), dirname(campaign.spago), process.env.PATH].join(delimiter) };
const commands = [];
const execute = (label, command, args, cwd = plan.project) => {
  const record = run(directory, label, command, args, cwd, env, 600000); commands.push(record);
  writeJson(join(directory, 'frontend-commands.json'), commands); return record;
};
execute('fetch', campaign.spago, ['fetch', '--offline']);
// Spec's integration module imports Spago's generated package information.
// Let Spago create its real BuildInfo before freezing the frontend inputs.
if (name === 'purust-spec') {
  // External-backend mode tells Spago to produce CoreFn and BuildInfo without
  // asking the JavaScript generator to validate native-only FFI. The backend
  // here is deliberately a no-op: timed generation runs separately below.
  const path = join(plan.project, 'spago.yaml'), original = readFileSync(path);
  const { parse, stringify } = createRequire(join(campaign.compilers.purust.origin, 'package.json'))('yaml');
  const config = parse(original.toString()); config.workspace.backend = { cmd: '/usr/bin/true', args: [] };
  writeFileSync(join(directory, 'frontend-original-spago.yaml'), original);
  writeFileSync(join(directory, 'frontend-spago.yaml'), stringify(config));
  try {
    writeFileSync(path, stringify(config));
    execute('spago-frontend', campaign.spago, ['build', '--offline']);
  } finally { writeFileSync(path, original); }
}
const sourcesRun = execute('sources', campaign.spago, ['sources', '--offline', '--json']);
const listed = JSON.parse(readFileSync(sourcesRun.stdout)); assert(Array.isArray(listed));
const globs = listed.filter(pattern => !/^\.?\/?test\//.test(pattern));
globs.push(plan.test_pattern);
if (existsSync(join(plan.project, '.spago/BuildInfo.purs')) && !globs.includes('.spago/BuildInfo.purs')) globs.push('.spago/BuildInfo.purs');
const output = join(plan.project, 'output');
execute('frontend', campaign.frontend.path, ['compile', ...globs, '--codegen', 'corefn', '--output', output]);
const tast = verifyTypedOutput(output, campaign.frontend.path);
assert(existsSync(join(output, plan.main, 'corefn.json')), `${name}: missing entry ${plan.main}`);
const modules = [];
for (const module of readdirSync(output).sort()) {
  const file = join(output, module, 'corefn.json'); if (!existsSync(file)) continue;
  const bytes = readFileSync(file), data = JSON.parse(bytes), source = resolve(plan.project, data.modulePath);
  const relocated = `sources/${module}/${basename(source)}`;
  for (const extension of ['purs', 'rs', 'rs.cargo.json']) {
    const original = source.replace(/\.purs$/, '.' + extension);
    if (existsSync(original)) copy(original, join(input, relocated.replace(/\.purs$/, '.' + extension)));
  }
  data.modulePath = relocated;
  mkdirSync(join(input, 'output', module), { recursive: true });
  writeFileSync(join(input, 'output', module, 'corefn.json'), JSON.stringify(data) + '\n');
  modules.push({ name: module, original_source: source, original_sha256: hash(bytes), types: data.typeTable.length });
}
// Retain test data and expected outputs for the separate application validation.
for (const path of walk(join(plan.project, 'test'))) copy(path, join(input, 'test', relative(join(plan.project, 'test'), path)));
for (const path of ['spago.yaml', 'spago.lock']) if (existsSync(join(plan.project, path))) copy(join(plan.project, path), join(input, path));
const definition = { name, family: 'purust', input, input_manifest: manifest(input), sibling_inputs: [],
  invocation: ['--source', 'output', '--main', plan.main, ...plan.flags], hosts: ['js', 'rust'],
  invocation_provenance: invocationProvenance, source_corpus: plan.project, scope: plan.scope, frontend: { ...tast, commands, modules } };
const previous = campaign.baseline_archive && join(campaign.baseline_archive, 'cases', name, 'canonical-generated');
if (previous && existsSync(previous)) definition.historical_oracle = { directory: previous, files: manifest(previous) };
writeJson(join(directory, 'definition.json'), definition);
console.log(`${name}: ${tast.modules} frozen typed modules; main=${plan.main}; ${plan.flags.join(' ')}`);
