// Freeze the native Purust family and prepare a reproducible two-host campaign.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { constants, copyFileSync, cpSync, existsSync, lstatSync, mkdirSync, readFileSync, readdirSync, readlinkSync, statSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { cpus, totalmem } from 'node:os';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';
import { findTypedCompiler } from '../../../../purust/purust/tools/native-workspace.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), htdocs = dirname(site);
const root = join(htdocs, 'purust/purust');
const [archiveArg, sourceArchiveArg, qualificationArg] = process.argv.slice(2);
const sourceArchive = sourceArchiveArg ? resolve(sourceArchiveArg) : null;
const baselineCampaign = sourceArchive ? JSON.parse(readFileSync(join(sourceArchive, 'campaign.json'))) : null;
const familyRoot = sourceArchive ? join(sourceArchive, 'family') : dirname(root);
const { parse, stringify } = createRequire(join(root, 'package.json'))('yaml');
const archive = resolve(archiveArg); assert(!existsSync(archive)); mkdirSync(archive, { recursive: true });
const copy = (from, to) => { mkdirSync(dirname(to), { recursive: true }); copyFileSync(from, to, constants.COPYFILE_FICLONE); };
copy(join(site, 'README.md'), join(archive, 'README-before.md'));
const qualificationPath = qualificationArg ? resolve(qualificationArg) : join(site, 'var/benchmark/compilation-refresh-20261004/campaign.json');
const qualification = JSON.parse(readFileSync(qualificationPath));
const compiler = join(archive, 'compilers/purust');
for (const file of qualification.compilers.purust.files) {
  assert.equal(hash(readFileSync(join(root, file.path))), file.sha256);
  copy(join(root, file.path), join(compiler, file.path));
}
const frontend = join(archive, 'frontend/purs'); copy(sourceArchive ? join(sourceArchive, 'frontend/purs') : findTypedCompiler(root), frontend);
const packages = readdirSync(familyRoot).filter(name => name.startsWith('purust-') && statSync(join(familyRoot, name)).isDirectory()).sort();
assert.equal(packages.length, 57);
const family = join(archive, 'family'); mkdirSync(family);
const excluded = new Set(['.git', '.spago', '.cache', '.purmeta', 'node_modules', 'output', 'target', '.DS_Store']);
const plans = [], ignoredSymlinks = [];
for (const name of packages) {
  const source = join(familyRoot, name), project = join(family, name);
  cpSync(source, project, { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE,
    filter: path => {
      if (excluded.has(basename(path)) || basename(path) === 'spago.lock') return false;
      if (!existsSync(path)) {
        assert(lstatSync(path).isSymbolicLink());
        assert.equal(path, join(familyRoot, 'purust-strings-extra/src/lib.rs'));
        ignoredSymlinks.push({ path, target: readlinkSync(path), reason: 'Broken historical hello-world link, unrelated to a PureScript module FFI' });
        return false;
      }
      return true;
    } });
  const originalText = readFileSync(join(source, 'spago.yaml'), 'utf8'), original = parse(originalText);
  copy(join(source, 'spago.yaml'), join(archive, 'original-configs', name + '.yaml'));
  const config = { package: original.package, workspace: {
    packageSet: original.workspace?.packageSet ?? { registry: '77.10.1' },
    extraPackages: Object.fromEntries(packages.filter(other => other !== name).map(other => [other.slice('purust-'.length), { path: '../' + other }])) } };
  // A port may keep a purust-* package name, while its siblings import the
  // registry name. Normalize the isolated root to prevent a second JS copy.
  config.package.name = name.slice('purust-'.length);
  let main = original.workspace?.backend?.args?.[original.workspace.backend.args.indexOf('--main') + 1];
  if (!original.workspace?.backend?.args?.includes('--main')) main = config.package.test?.main ?? 'Main';
  const flags = baselineCampaign?.plans.find(plan => plan.name === name)?.flags
    ?? (original.workspace?.backend?.args?.includes('--threaded') || name === 'purust-aff' ? ['--threaded'] : []);
  let testPattern = 'test/**/*.purs', scope = 'Package sources and primary test entry point';
  if (name === 'purust-argonaut-core') {
    main = 'Main'; testPattern = 'test/compact-dom.purs';
    config.package.test = { main, dependencies: [...new Set([...(config.package.test?.dependencies ?? []), 'assert', 'console', 'effect', 'st'])] };
    scope = 'Existing compact-DOM fixture; the legacy Test.Main foreign constants have no Rust implementation';
  }
  if (name === 'purust-argonaut-codecs') {
    main = 'Main'; testPattern = 'test/typed-plans.purs';
    config.package.test = { main, dependencies: ['assert', 'console', 'effect', 'refs', 'numbers'] };
    scope = 'Existing typed JSON decoding plans fixture';
  }
  if (name === 'purust-arraybuffer-types') {
    main = 'Benchmark.Main';
    config.package.test = { main, dependencies: ['effect', 'prelude'] };
    mkdirSync(join(project, 'test'), { recursive: true });
    writeFileSync(join(project, 'test/Main.purs'), 'module Benchmark.Main where\n\nimport Prelude (Unit, pure, unit)\nimport Data.ArrayBuffer.Types (ArrayBuffer)\nimport Effect (Effect)\n\nretain :: ArrayBuffer -> ArrayBuffer\nretain value = value\n\nmain :: Effect Unit\nmain = pure unit\n');
    scope = 'Foreign-type-only library with an explicit minimal executable harness';
  }
  writeFileSync(join(project, 'spago.yaml'), stringify(config));
  plans.push({ name, project, main, flags, test_pattern: testPattern, scope,
    original_configuration_sha256: hash(originalText), configuration_sha256: hash(readFileSync(join(project, 'spago.yaml'))) });
}
const sourceProvenance = [];
for (const name of ['purust/purust', 'purescript-backend-optimizer-purust']) {
  const source = join(htdocs, name), destination = join(archive, 'sources', name), files = manifest(join(source, 'src'));
  for (const file of files) copy(join(source, 'src', file.path), join(destination, file.path));
  sourceProvenance.push({ name, files, commit: execFileSync('git', ['rev-parse', 'HEAD'], { cwd: source, encoding: 'utf8' }).trim(),
    status: execFileSync('git', ['status', '--short'], { cwd: source, encoding: 'utf8' }) });
}
const campaign = { status: 'prepared', started_at: new Date().toISOString(),
  host: { cpu: cpus()[0].model, logical_cpus: cpus().length, memory_bytes: totalmem(), node: process.version },
  protocol: { rounds: 5, warmups_per_host: 1, metric: 'backend total', statistic: 'median',
    includes: 'TAST loading/sorting, preparation, PBO/generation, final emission and drain',
    excludes: 'frontend, compiler/application builds, application execution and process startup/exit',
    order: 'Serialized projects; JS/Rust and Rust/JS alternate over warmup and five measured rounds',
    cache: 'Fresh generated output and PBO caches; warm OS filesystem cache; fresh process per run',
    purust_jobs: 'Public defaults: JS 1/1; native worker budget 8 split into 4 PBO and 4 codegen workers',
    output_mode: 'Primary package runner flags; threaded where requested by the package',
    exact_output: 'Per-run manifests and byte comparisons; canonical identical output retained once per project' },
  compilers: { purust: { directory: compiler, origin: root, files: manifest(compiler) } },
  qualification: { path: qualificationPath, sha256: hash(readFileSync(qualificationPath)) },
  baseline_archive: sourceArchive,
  frontend: { path: frontend, sha256: hash(readFileSync(frontend)) },
  spago: join(root, 'node_modules/.bin/spago'), source_provenance: sourceProvenance,
  family_sources: manifest(family), ignored_symlinks: ignoredSymlinks, plans,
  library_cases: plans.map(plan => join(archive, 'cases', plan.name, 'definition.json')) };
writeJson(join(archive, 'campaign.json'), campaign);
console.log(`Prepared ${packages.length} native library projects in ${archive}`);
