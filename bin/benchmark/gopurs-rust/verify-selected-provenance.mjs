// Resolve isolated work/sources/1 aliases by content, not checkout-name hints.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync } from 'node:fs';
import { dirname, join, relative, resolve } from 'node:path';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';
import { threadedRust } from '../../../../purust/purust/src/Purust/Threading.js';

const [archiveArg, label, contractsArg] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label) && contractsArg);
const archive = resolve(archiveArg), candidate = join(archive, 'candidates', label);
const contracts = resolve(archive, contractsArg), destination = join(contracts, 'provenance.json');
assert(!existsSync(destination));
const summaryPath = join(contracts, 'summary.json');
const summary = JSON.parse(readFileSync(summaryPath));
assert.equal(summary.status, 'ok'); assert.equal(summary.results.length, 8);
assert.equal(summary.failed, 0); assert.equal(summary.skipped, 0);
assert.equal(summary.generatedRust, join(candidate, 'rust'));
const buildPath = join(candidate, 'build.json'), build = JSON.parse(readFileSync(buildPath));
assert.equal(build.status, 'passed');
assert.equal(hash(readFileSync(join(candidate, 'gopurs-rust'))), build.binary_sha256);
assert.equal(hash(readFileSync(build.generator.path)), build.generator.sha256);
assert.deepEqual(manifest(join(candidate, 'rust'), path => /\.(rs|toml)$/.test(path)), build.generated);
const frozen = join(candidate, 'sources/1/src'), live = join(summary.gopursPboDir, 'src');
const sources = manifest(frozen);
assert.deepEqual(manifest(live), sources, 'Live gopurs PBO differs from selected frozen sources');
assert.deepEqual(manifest(join(archive, 'work/sources/1/src')), sources);
assert.deepEqual(build.sources[1].files.filter(file => file.path.startsWith('src/'))
  .map(file => ({ ...file, path: file.path.slice(4) })), sources);
for (const source of summary.nativeSources) {
  const bytes = readFileSync(join(frozen, relative(live, source.source)));
  assert.equal(hash(bytes), source.sha256);
  assert(readFileSync(source.generated, 'utf8').includes(threadedRust(bytes.toString()).trim()));
}
const corpus = [], pbo = [];
for (const name of readdirSync(summary.corpus).sort()) {
  const input = join(summary.corpus, name, 'corefn.json');
  if (!existsSync(input)) continue;
  const bytes = readFileSync(input), module = JSON.parse(bytes);
  const path = join(name, 'corefn.json');
  const entry = { path, sha256: hash(bytes), modulePath: module.modulePath };
  corpus.push(entry);
  const destination = join(contracts, 'frozen-tast', path);
  mkdirSync(dirname(destination), { recursive: true });
  assert(!existsSync(destination)); copyFileSync(input, destination, constants.COPYFILE_FICLONE);
  assert(!module.modulePath.includes('purescript-backend-optimizer-purust'));
  if (name === 'PureScript.Backend.Optimizer' || name.startsWith('PureScript.Backend.Optimizer.')) {
    const actual = resolve(dirname(summary.corpus), module.modulePath);
    const expected = join(archive, 'work/sources/1/src', name.replaceAll('.', '/') + '.purs');
    assert.equal(actual, expected);
    assert.equal(hash(readFileSync(actual)), hash(readFileSync(join(frozen, relative(join(archive, 'work/sources/1/src'), actual)))));
    pbo.push(entry);
  }
}
assert.equal(corpus.length, build.tast.modules);
assert.equal(pbo.length, sources.filter(file => file.path.endsWith('.purs')).length);
assert(pbo.length > 20);
const result = { status: 'passed', verified_at: new Date().toISOString(),
  explanation: 'Isolated work/sources/1 aliases verified against complete gopurs PBO source manifests, all optimizer module paths, six embedded FFI sources and the selected generated Rust manifest. Original runner hint retained unchanged.',
  original_provenance_hint: summary.provenance,
  summary: { path: summaryPath, sha256: hash(readFileSync(summaryPath)) },
  candidate_build: { path: buildPath, sha256: hash(readFileSync(buildPath)) },
  binary_sha256: build.binary_sha256, generator: build.generator,
  source_manifest: sources, native_sources: summary.nativeSources,
  corpus, pbo_modules: pbo, frozen_tast: join(contracts, 'frozen-tast') };
writeJson(destination, result);
console.log(`${corpus.length} frozen TAST inputs, ${pbo.length} optimizer module paths, ${summary.nativeSources.length} embedded native sources: verified`);
