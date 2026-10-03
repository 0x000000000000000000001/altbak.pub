// Publish the gopurs Rust allocation campaign only after independently
// re-reading every proof, then update the single [gopurs-aff] README line.
// The final production qualification is expected from qualify.mjs:
// production/results.json (candidate, build_profile, cargo command, sources,
// executables) and production/hosts/results.json (three-host Aff + tests).
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, readdirSync, readFileSync, writeFileSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg] = process.argv.slice(2);
assert(archiveArg, 'publish-allocation.mjs ARCHIVE');
const archive = resolve(archiveArg);
const root = fileURLToPath(new URL('../../../', import.meta.url));
const parse = path => JSON.parse(readFileSync(path, 'utf8'));
const evidence = path => {
  const full = join(archive, path);
  assert(existsSync(full), 'Missing evidence: ' + full);
  return parse(full);
};
const sha = path => hash(readFileSync(path));
const median = values => {
  const xs = values.toSorted((a, b) => a - b), n = xs.length;
  return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2;
};
const paired = (campaign, before, after) => {
  const a = campaign.summary[before].samples_ms, b = campaign.summary[after].samples_ms;
  const deltas = b.map((n, i) => n - a[i]);
  return { deltas_ms: deltas, median_delta_ms: median(deltas),
    mean_delta_ms: deltas.reduce((x, y) => x + y, 0) / deltas.length,
    favorable: deltas.filter(n => n < 0).length, rounds: deltas.length };
};
const isPassed = result => { assert.equal(result.status, 'passed'); return result; };

// ---------------------------------------------------------------------------
// Evidence
// ---------------------------------------------------------------------------
const baseline = evidence('baseline.json');
const provenance = evidence('baseline-provenance.json');
const protocol = evidence('confirmation-protocol.json');
const finish = evidence('finish-selection.json');
const diagnostics = evidence('diagnostics.json');
const nativeDeleteDecision = evidence('native-delete-decision.json');
const nativeInsertDecision = evidence('native-insert-decision.json');
const native = evidence('native-tests/summary.json');
const pbo = evidence('pbo-tests/results.json');
const cleanup = evidence('purust-cleanup-results.json');
const production = evidence('production/results.json');
const finalSelection = evidence('final-selection.json');
const hosts = evidence('production/hosts/results.json');
const verification = evidence('verification-final.json');
const finishConfirmation = isPassed(evidence('finish-confirm.json'));
const primary = evidence('primary-runs/results.json');
const common = evidence('common-runs/results.json');
const defaultCampaign = evidence('default-runs/results.json');
const resources = evidence('resources-runs/results.json');
const selected = Object.fromEntries(['native-delete', 'native-insert', 'native-insert-confirmation', 'qualified']
  .map(name => [name, evidence(name + '-runs/results.json')]));

// ---------------------------------------------------------------------------
// Frozen baseline, optional diagnostics and Purust binaries
// ---------------------------------------------------------------------------
assert.equal(sha(join(archive, 'baseline/compiler/bin/gopurs-rust')), baseline.rust_sha256);
assert.equal(baseline.rust_sha256, protocol.baseline_sha256);
assert.deepEqual(manifest(join(archive, 'bootstrap')), baseline.bootstrap);
assert.equal(provenance.compiler_hashes['gopurs-rust'], baseline.rust_sha256);
const bootstrapSha = name => baseline.bootstrap.find(file => file.path === name).sha256;
assert.equal(hosts.bootstrap.purust_native_sha256, bootstrapSha('purust-native'));
assert.equal(hosts.bootstrap.purust_js_sha256, bootstrapSha('purust.js'));

// Purust bins are independently protected by the cleanup audit; the archived
// bootstraps above must keep the same hash.
const protectedFiles = evidence('purust-cleanup-protected.json').files;
const protectedPurustBins = protectedFiles.filter(file => /\/purust\/purust\/bin\/(purust-native|purust\.js)$/.test(file.path));
assert.equal(protectedPurustBins.length, 2);
for (const file of protectedPurustBins) assert.equal(file.sha256, bootstrapSha(basename(file.path)));
for (const file of protectedPurustBins) if (existsSync(file.path)) assert.equal(sha(file.path), file.sha256);

// ---------------------------------------------------------------------------
// Selection campaigns (binary variants frozen in candidates/)
// ---------------------------------------------------------------------------
function checkSelection(name, results, { rounds, variants }) {
  assert.equal(results.status, 'passed', name);
  assert.equal(results.protocol.rounds, rounds, name);
  assert.equal(results.protocol.warmups, 1, name);
  assert.deepEqual(results.variants.map(variant => variant.name), variants, name);
  assert.equal(results.runs.length, (rounds + 1) * variants.length, name);
  for (const [variant, summary] of Object.entries(results.summary))
    assert.equal(summary.samples_ms.length, rounds, name + ' ' + variant);
  for (const record of results.runs) {
    assert.equal(record.exit_code, 0, record.label);
    assert.equal(record.identical_files, 294, record.label);
  }
  for (const variant of results.variants) {
    if (!variant.binary) continue;
    const frozen = join(archive, name + '-runs/compilers', variant.name, 'gopurs-rust');
    assert.equal(sha(frozen), sha(resolve(archive, variant.binary)), name + ' ' + variant.name);
  }
  return results;
}
const nativeDelete = checkSelection('native-delete', selected['native-delete'],
  { rounds: 5, variants: ['baseline', 'control', 'native-delete'] });
const nativeInsert = checkSelection('native-insert', selected['native-insert'],
  { rounds: 5, variants: ['baseline', 'native-insert'] });
const insertConfirmation = checkSelection('native-insert-confirmation', selected['native-insert-confirmation'],
  { rounds: 15, variants: ['baseline', 'native-insert'] });
const qualified = checkSelection('qualified', selected.qualified,
  { rounds: 5, variants: ['baseline', 'native-insert', 'qualified'] });

// Native delete: rejected after 18 identical generations, no median gain.
assert.equal(nativeDeleteDecision.decision, 'not retained alone');
assert.equal(nativeDeleteDecision.baseline_ms, nativeDelete.summary.baseline.median_ms);
assert.equal(nativeDeleteDecision.control_ms, nativeDelete.summary.control.median_ms);
assert.equal(nativeDeleteDecision.candidate_ms, nativeDelete.summary['native-delete'].median_ms);
assert.equal(nativeDeleteDecision.identical_generations, nativeDelete.runs.length);
assert.equal(nativeDeleteDecision.native_contract, 'passed');
assert.equal(nativeDeleteDecision.semantic_suites, 15);
assert.equal(nativeDelete.runs.length, 18);

// Native insert: provisional selection then 15-pair confirmation; the
// decision file must reproduce the confirmation campaign exactly.
assert.equal(nativeInsertDecision.rounds, insertConfirmation.protocol.rounds);
assert.equal(nativeInsertDecision.generations, insertConfirmation.runs.length);
assert.equal(nativeInsertDecision.median_before_ms, insertConfirmation.summary.baseline.median_ms);
assert.equal(nativeInsertDecision.median_after_ms, insertConfirmation.summary['native-insert'].median_ms);
const insertConfirmationPaired = paired(insertConfirmation, 'baseline', 'native-insert');
assert.deepEqual(nativeInsertDecision.paired_deltas_ms, insertConfirmationPaired.deltas_ms);
assert.equal(nativeInsertDecision.paired_median_delta_ms, insertConfirmationPaired.median_delta_ms);
assert(Math.abs(nativeInsertDecision.paired_mean_delta_ms - insertConfirmationPaired.mean_delta_ms) < 1e-9);
assert.equal(nativeInsertDecision.favorable, insertConfirmationPaired.favorable);
assert.equal(nativeInsertDecision.identical_files, insertConfirmation.runs.reduce((n, record) => n + record.identical_files, 0));

// Qualified = native-insert + borrowed Qualified comparisons. It must beat
// native-insert in every selection round (five rotating rounds).
const composedPaired = paired(qualified, 'native-insert', 'qualified');
assert.equal(composedPaired.favorable, composedPaired.rounds);
assert(composedPaired.deltas_ms.every(delta => delta < 0));
const qualifiedLog = readFileSync(join(archive, 'native-tests/qualified.console.log'), 'utf8');
const qualifiedPairs = Number(qualifiedLog.match(/Native Qualified comparison: (\d+) pairs/)?.[1]);
assert.equal(qualifiedPairs, 396900);
assert.equal(finish.status, 'passed');
assert.deepEqual(finish.commands.map(record => record.label), ['qualified-adapters', 'pbo-tests', 'qualified-selection']);
for (const record of finish.commands) assert.equal(record.exit_code, 0, record.label);

// ---------------------------------------------------------------------------
// Native differential and PBO semantic suites
// ---------------------------------------------------------------------------
assert.equal(native.status, 'ok');
assert.equal(native.skipped, 0);
assert.equal(native.failed, 0);
assert(native.results.length >= 8);
for (const result of native.results) { assert.equal(result.status, 'ok'); assert.equal(result.exitCode, 0); }
for (const contract of native.nativeSources) assert.equal(sha(contract.source), contract.sha256);
assert.equal(pbo.status, 'passed');
assert(pbo.runs.length >= 15);
for (const record of pbo.runs) {
  assert.equal(record.exit_code, 0, record.label);
  assert.equal(sha(join(archive, 'pbo-tests/tests', record.label + '.mjs')), record.test_sha256, record.label);
}
// Candidate-level diagnostics retained with the rejected/selected patches.
const nativeDeleteContract = evidence('native-delete-contract/summary.json');
const nativeDeleteTests = evidence('native-delete-tests/summary.json');
const nativeInsertContract = evidence('native-insert-contract/summary.json');
assert.equal(nativeDeleteContract.status, 'ok');
assert.equal(nativeDeleteContract.failed, 0);
assert.equal(nativeInsertContract.status, 'ok');
assert.equal(nativeInsertContract.failed, 0);
// The rejected native-delete patch is still exercised by the reinforced
// contract: exact shape against the generated oracle, key/value model and
// persistence over 25200 mixed operations. The earlier strict balance
// assertions failed identically on the oracle (the generated joins can produce
// a height difference of 2), so their failed summaries stay as raw diagnostics
// and are never presented as semantic incorrectness.
const nativeDeleteContractLog = readFileSync(join(archive, 'native-delete-contract/maps.log'), 'utf8');
assert.match(nativeDeleteContractLog, /exact AVL shape and persistent versions passed/);
const nativeDeleteSupplementOperations = Number(nativeDeleteContractLog.match(/Native Maps: (\d+) mixed updates\/deletes/)?.[1]);
assert.equal(nativeDeleteSupplementOperations, 25200);
assert.match(readFileSync(join(archive, 'native-delete-tests/maps.log'), 'utf8'), /unbalanced deletion/);
assert.match(readFileSync(join(archive, 'native-delete-maps-delete.rs'), 'utf8'), /height difference of 2/);
for (const name of ['native-delete-pbo', 'native-insert-pbo']) {
  const candidateSuites = evidence(name + '/results.json');
  assert.equal(candidateSuites.status, 'passed', name);
  assert.equal(candidateSuites.runs.length, 15, name);
  for (const record of candidateSuites.runs) assert.equal(record.exit_code, 0, name + ' ' + record.label);
}

// ---------------------------------------------------------------------------
// Production qualification: profile, sources, Rust generation, hosts, tests
// ---------------------------------------------------------------------------
isPassed(production);
isPassed(hosts);
assert(production.candidate && existsSync(join(production.candidate, 'build.json')), production.candidate);
const candidateBuild = parse(join(production.candidate, 'build.json'));
assert.equal(candidateBuild.status, 'passed');
assert.equal(basename(production.candidate), finalSelection.candidate);
assert.equal(candidateBuild.binary_sha256, finalSelection.binary_sha256);

// The final profile is derived from the selected candidate build (explicit
// profile metadata when available, otherwise the Cargo command), then
// cross-checked against the qualification build_profile and the exact Cargo
// command used by the public Rust/Aff -c path. A rust-lld linker override is
// distinct metadata (result.linker), never part of build_profile.
function cargoRecord(build) {
  return build.commands.find(record => ['cargo', 'cargo-rustc', 'rustc'].includes(record.label) && record.exit_code === 0)
    ?? build.commands.find(record => Array.isArray(record.command) && record.command[0] === 'cargo' && record.exit_code === 0);
}
function buildProfile(candidate) {
  const build = parse(join(candidate, 'build.json'));
  const cargo = cargoRecord(build);
  assert(cargo, 'candidate Cargo command: ' + candidate);
  const explicit = Object.fromEntries(['lto', 'opt-level', 'debug'].map(key => {
    const argument = cargo.command.find(value => value.startsWith('profile.release.' + key + '='));
    return [key, argument ? argument.slice(argument.indexOf('=') + 1) : undefined];
  }));
  if (explicit['opt-level'] !== undefined) assert.equal(JSON.parse(explicit['opt-level']), 3);
  if (explicit.debug !== undefined) assert.equal(JSON.parse(explicit.debug), false);
  const profile = { opt_level: 3, debug: false,
    lto: explicit.lto !== undefined ? JSON.parse(explicit.lto) : (build.profile?.lto ?? false),
    threaded: true, allocator: 'mimalloc' };
  if (build.profile) {
    for (const key of ['opt_level', 'debug', 'lto', 'threaded', 'allocator'])
      assert.equal(build.profile[key], profile[key], candidate + ' profile ' + key);
    assert.equal(Object.keys(build.profile).length, 5, candidate + ' profile fields');
  }
  return profile;
}
const profile = buildProfile(production.candidate);
assert.deepEqual(profile, finalSelection.profile);
assert([false, 'thin'].includes(profile.lto));
const candidateBuildCargo = cargoRecord(candidateBuild);
const candidateLinker = candidateBuild.linker ?? null;
if (candidateLinker) {
  assert.equal(candidateLinker.driver, 'cc');
  assert.match(candidateLinker.option, /^-fuse-ld=.*ld64\.lld$/);
  assert.match(candidateLinker.path, /gcc-ld\/ld64\.lld$/);
  if (existsSync(candidateLinker.path)) assert.equal(sha(candidateLinker.path), candidateLinker.sha256);
  assert(candidateBuildCargo.command.includes('rustc'), 'lld candidate uses cargo rustc');
  assert(candidateBuildCargo.command.includes('link-arg=' + candidateLinker.option), 'lld link-arg');
}
const productionLinker = production.linker ?? candidateLinker;
if (production.linker && candidateLinker) assert.deepEqual(production.linker, candidateLinker);
if (profile.lto === 'thin') {
  assert(productionLinker, 'thin LTO profile without linker metadata');
  assert.equal(productionLinker.driver, 'cc');
  assert.match(productionLinker.option, /^-fuse-ld=.*ld64\.lld$/);
  if (existsSync(productionLinker.path)) assert.equal(sha(productionLinker.path), productionLinker.sha256);
}
if (production.build_profile) {
  for (const key of ['opt_level', 'debug', 'lto', 'threaded', 'allocator'])
    assert.equal(production.build_profile[key], profile[key], 'production build_profile ' + key);
  assert(!('linker' in production.build_profile), 'linker metadata must stay outside build_profile');
}
if (production.cargo_command) {
  assert.match(production.cargo_command, /^cargo /);
  assert(production.cargo_command.includes('profile.release.opt-level=3'));
  assert(production.cargo_command.includes('profile.release.debug=false'));
  assert(production.cargo_command.includes(`profile.release.lto=${JSON.stringify(profile.lto)}`));
  if (profile.lto === 'thin') {
    assert.match(production.cargo_command, /^cargo rustc /);
    assert(production.cargo_command.includes(productionLinker.option), 'public cargo command uses the lld option');
  }
}
if (production.build_rust_sha256) {
  const copies = readdirSync(join(archive, 'production')).filter(name => /^build-rust.*\.mjs$/.test(name));
  assert(copies.length >= 1);
  assert(copies.some(name => sha(join(archive, 'production', name)) === production.build_rust_sha256));
}
assert(production.rust_bootstrap && existsSync(production.rust_bootstrap));
assert(production.go_bootstrap && existsSync(production.go_bootstrap));

for (const [name, executable] of Object.entries(production.executables)) {
  assert.equal(sha(executable.path), executable.sha256, name);
  assert.equal(hosts.executables[name].sha256, executable.sha256, name);
}
assert.deepEqual(manifest(join(hosts.compiler, 'src')), production.sources.gopurs);
assert.deepEqual(manifest(resolve(hosts.compiler, '../../purescript-backend-optimizer-gopurs/src')), production.sources.optimizer);
assert.deepEqual(manifest(join(production.candidate, 'sources/0/src')), production.sources.gopurs);
assert.deepEqual(manifest(join(production.candidate, 'sources/1/src')), production.sources.optimizer);
assert.deepEqual(manifest(join(hosts.aff, 'test')), hosts.sources.aff);
assert.equal(sha(join(archive, 'frontend/purs')), production.frontend_sha256);
assert.equal(sha(join(production.rust_bootstrap, 'target/release/purust_output')), production.executables['gopurs-rust'].sha256);
assert.equal(readFileSync(join(production.rust_bootstrap, 'smoke-go-run.log'), 'utf8').trim(), 'Done');
assert.equal(sha(join(production.go_bootstrap, 'gopurs-native')), production.executables['gopurs-native'].sha256);

const allRustFiles = directory => manifest(directory, file => /\.(rs|toml)$/.test(file));
const rustFiles = directory => manifest(directory, file => /\.(rs|toml)$/.test(file)
  && !file.endsWith('/Purs_Gopurs_FfiSupport/build.rs'));
const candidateRust = allRustFiles(join(production.candidate, 'rust'));
assert.deepEqual(candidateRust, candidateBuild.generated);
assert.deepEqual(allRustFiles(join(production.rust_bootstrap, 'rust')), candidateRust);
assert.equal(production.identical_candidate_rust_files, candidateRust.length);
assert.equal(candidateRust.filter(file => file.path === 'Purs_Gopurs_FfiSupport/build.rs').length, 1,
  'exactly one FfiSupport build.rs');
const productionRustOnly = rustFiles(join(production.rust_bootstrap, 'rust'));
assert.equal(candidateRust.length - productionRustOnly.length, 1, 'only build.rs is excluded from the Rust manifest');
assert.deepEqual(productionRustOnly, hosts.bootstrap.native_generation);
assert.deepEqual(rustFiles(join(archive, 'production/hosts/bootstrap-js')), hosts.bootstrap.native_generation);
assert.equal(hosts.bootstrap.identical_files, hosts.bootstrap.native_generation.length);

const preparation = production.commands.find(record => record.label.startsWith('gopurs-preparation-tests'));
assert(preparation && preparation.exit_code === 0);
const preparationLog = readFileSync(preparation.stdout, 'utf8');
const preparationPass = Number(preparationLog.match(/ℹ pass (\d+)/)?.[1]);
assert(preparationPass >= 19, 'preparation pass count: ' + preparationPass);
assert.match(preparationLog, /ℹ fail 0\b/);
assert.match(preparationLog, /ℹ skipped 0\b/);
const rustRebuild = production.commands.find(record => record.label.startsWith('aff-rust-rebuild-default'));
assert(rustRebuild && rustRebuild.exit_code === 0);
const memoContract = production.commands.find(record => record.label.startsWith('go-memo-contract'));
assert(memoContract && memoContract.exit_code === 0);

// Three-host qualification: exact Aff oracles, TAST identity, 294 files.
const live = hosts.runs.filter(record => record.host);
assert.deepEqual(live.map(record => record.host).sort(), ['go', 'js', 'rust']);
const expectedAff = readFileSync(join(hosts.aff, 'test/expected-main.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort();
for (const record of live) {
  assert(readFileSync(record.stderr, 'utf8').includes('[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true'));
  assert.deepEqual(manifest(record.go_output, file => file.endsWith('.go') || file.endsWith('/go.mod')), record.generated);
  assert.deepEqual(record.generated, live[0].generated);
  assert.equal(record.generated.length, 294);
  assert.deepEqual(readFileSync(join(record.go_output, 'test.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort(), expectedAff);
  assert.equal(readFileSync(join(record.go_output, 'test.stderr'), 'utf8'), '');
}
const compilerTests = hosts.runs.find(record => record.label === 'compiler-tests');
assert(compilerTests && compilerTests.exit_code === 0);
const compilerLog = readFileSync(compilerTests.stdout, 'utf8');
const compilerPass = Number(compilerLog.match(/ℹ pass (\d+)/)?.[1]);
assert(compilerPass >= 43, 'compiler test pass count: ' + compilerPass);
assert.match(compilerLog, /ℹ fail 0\b/);
assert.match(compilerLog, /ℹ skipped 0\b/);
const parserTests = hosts.runs.find(record => record.label === 'parser-tests');
assert(parserTests && parserTests.exit_code === 0);

// ---------------------------------------------------------------------------
// Final timing campaigns: primary, common, defaults, resources
// ---------------------------------------------------------------------------
const workerEnvironment = { GOPURS_JOBS: '8', GOPURS_PREPARE_JOBS: '8', GOPURS_PBO_JOBS: '8', GOPURS_EMIT_JOBS: '8', GOPURS_PIPELINE: '1' };
const finalCampaigns = [
  ['primary', primary, 15, ['rust-before', 'rust']],
  ['common', common, 10, ['js', 'go', 'rust-before', 'rust']],
  ['default', defaultCampaign, 5, ['rust-before', 'rust', 'go', 'js']],
  ['resources', resources, 3, ['rust-before', 'rust']],
];
for (const [name, campaign, rounds, variants] of finalCampaigns) {
  assert.equal(campaign.status, 'passed', name);
  assert.equal(campaign.protocol.rounds, rounds, name);
  assert.equal(campaign.protocol.warmups, 1, name);
  assert.deepEqual(campaign.variants.map(variant => variant.name), variants, name);
  assert(campaign.variants.every(variant => variant.directory), name + ': directory variants');
  assert.equal(campaign.runs.length, (rounds + 1) * variants.length, name);
  assert.equal(campaign.selection.resources ?? false, name === 'resources', name);
  assert.equal(campaign.selection.launcherDefaults ?? false, name === 'default', name);
  if (name === 'default') assert.match(campaign.protocol.jobs, /^launcher defaults/);
  else assert.deepEqual(campaign.protocol.jobs, { load: 8, prepare: 8, pbo: 8, emit: 8, pipeline: true }, name);
  assert.deepEqual(campaign.tast, common.tast, name);
  assert.deepEqual(campaign.frozen.inputs, common.frozen.inputs, name);
  for (const [variant, summary] of Object.entries(campaign.summary))
    assert.equal(summary.samples_ms.length, rounds, name + ' ' + variant);
  const runDirectory = join(archive, name + '-runs');
  for (const record of campaign.runs) {
    assert.equal(record.exit_code, 0, record.label);
    assert.equal(record.identical_files, 294, record.label);
    assert.equal(parse(record.generated_manifest).length, 294, record.label);
  }
  // Every final variant uses directory, so the frozen baseline compiler is
  // copied to compilers/rust-before (bin/gopurs-rust), not gopurs-rust.
  assert.equal(sha(join(runDirectory, 'compilers/rust/bin/gopurs-rust')), production.executables['gopurs-rust'].sha256, name);
  assert.equal(sha(join(runDirectory, 'compilers/rust-before/bin/gopurs-rust')), baseline.rust_sha256, name);
  if (variants.includes('js'))
    assert.equal(sha(join(runDirectory, 'compilers/js/bin/gopurs.js')), production.executables['gopurs.js'].sha256, name);
  if (variants.includes('go'))
    assert.equal(sha(join(runDirectory, 'compilers/go/bin/gopurs-native')), production.executables['gopurs-native'].sha256, name);
}
for (const record of [...primary.runs, ...common.runs, ...resources.runs]) {
  for (const [key, value] of Object.entries(workerEnvironment))
    assert.equal(record.explicit_environment[key], value, record.label + ' ' + key);
}
for (const record of defaultCampaign.runs) {
  assert(readFileSync(record.stderr, 'utf8').includes('[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true'));
  const selector = record.variant === 'js' ? 'GOPURS_JS' : record.variant.startsWith('rust') ? 'GOPURS_RUST' : null;
  assert.deepEqual(Object.keys(record.explicit_environment).sort(), selector ? [selector] : [], record.label);
  if (selector) assert.equal(record.explicit_environment[selector], '1', record.label);
}

const before = primary.summary['rust-before'], after = primary.summary.rust, c = common.summary;
const primaryPaired = paired(primary, 'rust-before', 'rust');
const improvement = 100 * (1 - after.median_ms / before.median_ms);
const ratio = c.rust.median_ms / c.go.median_ms;
const oldRatio = c['rust-before'].median_ms / c.go.median_ms;
const gap = 100 * (1 - (c.rust.median_ms - c.go.median_ms) / (c['rust-before'].median_ms - c.go.median_ms));
const defaultPaired = paired(defaultCampaign, 'rust-before', 'rust');
const defaultReduction = 100 * (1 - defaultCampaign.summary.rust.median_ms / defaultCampaign.summary['rust-before'].median_ms);

const attempts = Object.fromEntries(['rust-before', 'rust'].map(variant => [variant,
  primary.runs.filter(record => record.variant === variant).map(record => {
    const text = readFileSync(record.stderr, 'utf8');
    const parsed = text.match(/pbo module attempts: (\d+), codegen: (\d+)/);
    const deferred = text.match(/deferredAttempts=(\d+)/);
    assert(parsed && deferred, record.stderr);
    return { attempts: Number(parsed[1]), codegen: Number(parsed[2]), deferred: Number(deferred[1]) };
  })]));
assert(attempts.rust.every(record => record.attempts === 238 && record.codegen === 238 && record.deferred === 0));

const resourceSummary = Object.fromEntries(['rust-before', 'rust'].map(variant => {
  const records = resources.runs.filter(record => !record.warmup && record.variant === variant).map(record => {
    const text = readFileSync(record.stderr, 'utf8');
    const clock = text.match(/([\d.]+) real\s+([\d.]+) user\s+([\d.]+) sys/);
    const rss = text.match(/(\d+)\s+maximum resident set size/);
    assert(clock && rss, record.stderr);
    return { real_s: Number(clock[1]), user_s: Number(clock[2]), sys_s: Number(clock[3]),
      cpu_s: Number(clock[2]) + Number(clock[3]), rss_bytes: Number(rss[1]) };
  });
  return [variant, { records, medians: Object.fromEntries(['real_s', 'cpu_s', 'rss_bytes']
    .map(key => [key, median(records.map(record => record[key]))])) }];
}));

// ---------------------------------------------------------------------------
// Independent verification of every *-runs campaign actually created
// ---------------------------------------------------------------------------
isPassed(verification);
for (const checked of verification.campaigns) {
  assert(resolve(checked.directory).startsWith(archive + '/'), checked.directory);
  assert.equal(sha(join(checked.directory, 'results.json')), checked.results_sha256);
  const campaign = parse(join(checked.directory, 'results.json'));
  assert.equal(campaign.status, 'passed', checked.directory);
  assert.equal(campaign.runs.length, checked.generations, checked.directory);
  assert.equal(campaign.runs.reduce((n, record) => n + record.identical_files, 0), checked.identical_files, checked.directory);
}
const discoveredCampaigns = readdirSync(archive, { withFileTypes: true })
  .filter(entry => entry.isDirectory() && entry.name.endsWith('-runs') && existsSync(join(archive, entry.name, 'results.json')))
  .map(entry => entry.name).sort();
const verifiedDirectories = new Set(verification.campaigns.map(checked => resolve(checked.directory)));
for (const name of discoveredCampaigns)
  assert(verifiedDirectories.has(join(archive, name)),
    `verification-final.json is missing ${name}; rerun verify.mjs with every *-runs campaign`);
assert.equal(verification.generations, verification.campaigns.reduce((n, checked) => n + checked.generations, 0));
assert.equal(verification.identical_files, verification.campaigns.reduce((n, checked) => n + checked.identical_files, 0));

// ---------------------------------------------------------------------------
// Final profiles (separate from timing) and cleanup audit
// ---------------------------------------------------------------------------
const finalAll = evidence('profile-final-all/summary.json');
const finalOptimize = evidence('profile-final-optimize/summary.json');
const finalProfileByPhase = { all: finalAll, optimize: finalOptimize };
for (const [phase, summary] of Object.entries(finalProfileByPhase)) {
  assert.equal(summary.phase, phase);
  assert.equal(summary.binary_sha256, production.executables['gopurs-rust'].sha256);
  assert.equal(summary.identical_go_files, 294);
  assert(summary.samples > 0);
}
const baselineProfile = evidence('profile-baseline/summary.json');
const baselineProfileAll = evidence('profile-baseline-all/summary.json');
const baselineByPhase = { all: baselineProfileAll, optimize: baselineProfile };
for (const [phase, summary] of Object.entries(baselineByPhase)) {
  const entry = diagnostics.find(item => item.phase === phase);
  assert(entry, phase);
  assert.equal(entry.samples, summary.samples);
  assert.deepEqual(entry.inclusive_families, summary.inclusive_families);
  assert.equal(entry.identical_go_files, 294);
  assert.equal(summary.binary_sha256, baseline.rust_sha256);
}

assert.equal(cleanup.status, 'passed');
assert(cleanup.free_bytes_gained > 0);
assert(cleanup.targets.length >= 1 && cleanup.targets.every(target => target.removed));
assert(cleanup.protected_files_verified > 0);
assert(cleanup.repository_statuses_unchanged_except_finder_metadata > 0);

// ---------------------------------------------------------------------------
// ThinLTO: two retained build failures (full disk, then system linker) and the
// rust-lld candidate. A failed thin-lto-retry folder never implies campaign
// success: the measured binary is discovered from the campaign selection and
// its build metadata, not from an assumed label.
// ---------------------------------------------------------------------------
function retainedFailure(label, pattern) {
  const path = join(archive, 'candidates', label, 'build.json');
  if (!existsSync(path)) return null;
  const build = parse(path);
  assert.equal(build.status, 'failed', label);
  const logs = [join(archive, 'build-' + label + '.log'), join(archive, 'candidates', label, 'logs/cargo.stderr'),
    join(archive, 'candidates', label, 'logs/cargo.stdout')].filter(existsSync);
  const text = logs.map(log => readFileSync(log, 'utf8')).join('\n');
  assert.match(text, pattern, label);
  return { label, profile: build.profile ?? null, failure: build.failure ?? null, logs };
}
const thinDiagnostics = [
  retainedFailure('thin-lto', /No space left on device/),
  retainedFailure('thin-lto-retry', /LLVM APPLE_1_1700|could not parse bitcode|Unknown attribute kind \(102\)/),
].filter(Boolean);

const lldDirectory = join(archive, 'candidates/thin-lto-rust-lld');
let thinCandidate = null;
if (existsSync(join(lldDirectory, 'build.json'))) {
  const build = parse(join(lldDirectory, 'build.json'));
  assert.equal(build.status, 'passed');
  const candidateProfile = buildProfile(lldDirectory);
  assert.equal(candidateProfile.lto, 'thin');
  assert(build.linker, 'thin-lto-rust-lld linker metadata');
  assert.equal(build.linker.driver, 'cc');
  assert.match(build.linker.option, /^-fuse-ld=.*ld64\.lld$/);
  const binary = join(lldDirectory, 'gopurs-rust');
  assert.equal(sha(binary), build.binary_sha256);
  thinCandidate = { label: build.label ?? 'thin-lto-rust-lld', binary, binary_sha256: build.binary_sha256,
    profile: candidateProfile, linker: build.linker };
}
if (existsSync(join(archive, 'thin-lto.json')))
  assert(existsSync(join(archive, 'thin-lto-runs/results.json')),
    'thin-lto.json exists but thin-lto-runs/results.json is missing');

let thinComparison = null, thinMeasured = [], thinConfirmation = null;
if (existsSync(join(archive, 'thin-lto-runs/results.json'))) {
  const runs = evidence('thin-lto-runs/results.json');
  assert.equal(runs.status, 'passed');
  assert.equal(runs.protocol.rounds, 5);
  assert.equal(runs.runs.length, (runs.protocol.rounds + 1) * runs.variants.length);
  assert(runs.variants.some(variant => variant.name === 'baseline'));
  assert(runs.variants.some(variant => variant.name === 'qualified'));
  for (const record of runs.runs) {
    assert.equal(record.exit_code, 0, record.label);
    assert.equal(record.identical_files, 294, record.label);
  }
  for (const variant of runs.variants) {
    if (!variant.binary) continue;
    assert.equal(sha(join(archive, 'thin-lto-runs/compilers', variant.name, 'gopurs-rust')),
      sha(resolve(archive, variant.binary)), variant.name);
  }
  const measuredVariants = runs.variants.filter(variant => !['baseline', 'qualified'].includes(variant.name));
  assert(measuredVariants.length >= 1, 'thin-lto measured candidate');
  thinMeasured = measuredVariants.map(variant => {
    assert(variant.binary, 'thin-lto measured binary: ' + variant.name);
    const binary = resolve(archive, variant.binary);
    const build = parse(join(dirname(binary), 'build.json'));
    assert.equal(build.status, 'passed', variant.name);
    assert.equal(buildProfile(dirname(binary)).lto, 'thin', variant.name);
    assert.equal(sha(binary), build.binary_sha256, variant.name);
    if (build.linker) {
      assert.equal(build.linker.driver, 'cc', variant.name);
      assert.match(build.linker.option, /^-fuse-ld=.*ld64\.lld$/, variant.name);
    }
    return { name: variant.name, binary, binary_sha256: build.binary_sha256, linker: build.linker ?? null };
  });
  thinComparison = { protocol: runs.protocol, selection: runs.selection, summary: runs.summary };
}
if (existsSync(join(archive, 'thin-lto-confirmation-runs/results.json'))) {
  const confirmation = evidence('thin-lto-confirmation-runs/results.json');
  assert.equal(confirmation.status, 'passed');
  assert.equal(confirmation.protocol.rounds, 15);
  for (const record of confirmation.runs) assert.equal(record.identical_files, 294, record.label);
  thinConfirmation = { protocol: confirmation.protocol, selection: confirmation.selection, summary: confirmation.summary };
}
const thinLto = { diagnostics: thinDiagnostics, candidate: thinCandidate, measured: thinMeasured,
  comparison: thinComparison, confirmation: thinConfirmation };

// ---------------------------------------------------------------------------
// Concurrent source integration: preserved in the production candidate
// ---------------------------------------------------------------------------
const integrationPath = join(archive, 'source-integration.json');
const integration = existsSync(integrationPath) ? parse(integrationPath) : null;
if (integration) {
  assert(integration.selection_candidate, 'integration selection_candidate');
  const integrationChanges = integration.changes.flatMap(group => group.changes.map(change => ({ index: group.index, ...change })));
  assert(integrationChanges.length > 0);
  const selectionCandidate = join(archive, 'candidates', integration.selection_candidate);
  assert.equal(parse(join(selectionCandidate, 'build.json')).status, 'passed');
  const sourceManifest = index => production.sources[index === 0 ? 'gopurs' : 'optimizer'];
  integration.preservation = integrationChanges.map(change => {
    assert.equal(sha(join(selectionCandidate, 'sources', String(change.index), 'src', change.before.path)), change.before.sha256,
      change.before.path);
    const entry = sourceManifest(change.index).find(file => file.path === change.after.path);
    assert(entry, change.after.path);
    return { index: change.index, path: change.after.path,
      status: entry.sha256 === change.after.sha256 ? 'exact' : entry.sha256 === change.before.sha256 ? 'missing' : 'diverged' };
  });
  assert(integration.preservation.every(item => item.status !== 'missing'),
    'production candidate does not include the source-integration changes');
  integration.changes_count = integrationChanges.length;
  integration.exact = integration.preservation.every(item => item.status === 'exact');
}
let integrationRuns = null, integrationIncludesProduction = false;
if (existsSync(join(archive, 'integration-runs/results.json'))) {
  const runs = evidence('integration-runs/results.json');
  assert.equal(runs.status, 'passed');
  const rounds = runs.protocol.rounds;
  assert(Number.isInteger(rounds) && rounds >= 5);
  assert.equal(runs.runs.length, (rounds + 1) * runs.variants.length);
  for (const record of runs.runs) {
    assert.equal(record.exit_code, 0, record.label);
    assert.equal(record.identical_files, 294, record.label);
  }
  for (const variant of runs.variants) {
    if (!variant.binary) continue;
    assert.equal(sha(join(archive, 'integration-runs/compilers', variant.name, 'gopurs-rust')),
      sha(resolve(archive, variant.binary)), variant.name);
  }
  const productionHash = production.executables['gopurs-rust'].sha256;
  integrationIncludesProduction = runs.variants.some(variant => {
    if (variant.binary) {
      const binary = resolve(archive, variant.binary);
      return binary.startsWith(production.candidate + '/') || sha(binary) === productionHash;
    }
    if (variant.directory) {
      const frozen = join(archive, 'integration-runs/compilers', variant.name, 'bin/gopurs-rust');
      return existsSync(frozen) && sha(frozen) === productionHash;
    }
    return false;
  });
  assert(integrationIncludesProduction, 'integration campaign does not include the production candidate');
  integrationRuns = { protocol: runs.protocol, selection: runs.selection, summary: runs.summary };
}
if (integration) assert(integrationRuns,
  'source-integration.json changes are installed but integration-runs/results.json is missing');
if (profile.lto === 'thin')
  assert(thinMeasured.some(entry => entry.binary_sha256 === production.executables['gopurs-rust'].sha256)
    || integrationIncludesProduction, 'production ThinLTO binary is not bound to a measured comparison');
const integrationAttribution = integration
  ? 'the installed baseline -> production delta includes the archived external source integration, not allocations alone; isolated native selections remain measured on the older snapshot'
  : 'the installed baseline -> production delta reflects the measured allocation composition';

// ---------------------------------------------------------------------------
// Historical JSON -> Typed AST preservation and separate historical gains
// ---------------------------------------------------------------------------
const historicalPath = join(root, 'docs/benchmark-results/2026-10-03-purust-native-optimization.json');
const historical = parse(historicalPath);
const historicalJson = Object.fromEntries(['fixture12', 'gopurs238'].map(name => {
  const item = historical.json[name];
  assert.equal(item.status, 'passed');
  for (const binary of Object.values(item.binaries)) assert.equal(sha(binary.path), binary.sha256);
  for (const record of item.runs) {
    assert.equal(record.exit_code, 0);
    const stdout = join(dirname(dirname(item.binaries[record.variant].path)), record.stdout);
    assert.deepEqual(parse(stdout), record.result);
    assert.deepEqual(record.result.fingerprints, item.runs[0].result.fingerprints);
    assert.deepEqual(record.result.json_fingerprints, item.runs[0].result.json_fingerprints);
  }
  return [name, { modules: item.modules, processes: item.runs.length,
    structural_fingerprints_sha256: hash(JSON.stringify(item.runs[0].result.fingerprints)),
    binaries: item.binaries }];
}));
const historicalOptimizationPath = join(root, 'docs/benchmark-results/2026-10-03-gopurs-rust-optimization.json');
const historicalDefaultsPath = join(root, 'docs/benchmark-results/2026-10-03-gopurs-host-defaults.json');
const historicalOptimization = parse(historicalOptimizationPath);
const historicalDefaults = parse(historicalDefaultsPath);
const historicalGains = {
  optimization: { path: historicalOptimizationPath, sha256: sha(historicalOptimizationPath),
    improvement_percent: historicalOptimization.primary.improvement_percent,
    before_ms: historicalOptimization.primary['rust-before'].median_ms,
    after_ms: historicalOptimization.primary.rust.median_ms },
  defaults: { path: historicalDefaultsPath, sha256: sha(historicalDefaultsPath),
    reduction_percent: historicalDefaults.default_rust_reduction_percent,
    before_ms: historicalDefaults.summary['rust-before-default'].median_ms,
    after_ms: historicalDefaults.summary['rust-default'].median_ms },
};

// The milestone concerns the installed compiler, not superseded selections.
const subThreeCampaigns = {
  primary: after.median_ms,
  common: c.rust.median_ms,
  default: defaultCampaign.summary.rust.median_ms,
};
const finalSamples = [primary, common, defaultCampaign].flatMap(campaign => campaign.summary.rust.samples_ms);
const subThreeBelow = Object.entries(subThreeCampaigns).filter(([, ms]) => ms < 3000).map(([name]) => name);
const subThree = { threshold_ms: 3000, campaigns: subThreeCampaigns, below: subThreeBelow,
  samples: finalSamples.length, samples_below: finalSamples.filter(ms => ms < 3000).length,
  claim: finalSamples.every(ms => ms < 3000) };

// ---------------------------------------------------------------------------
// Helper command references and source-record fingerprints
// ---------------------------------------------------------------------------
const helperRecords = [...finish.commands, ...finishConfirmation.commands, ...production.commands, ...hosts.runs]
  .map(({ label, exit_code, stdout, stderr }) => ({ label, exit_code, stdout, stderr }));
for (const helper of helperRecords) for (const key of ['stdout', 'stderr']) {
  assert(helper[key], helper.label + ' ' + key);
  assert(existsSync(helper[key]), 'Missing helper log: ' + helper[key]);
}
const relativeEvidence = [
  'final-selection.json', 'profile-selection.json', 'thin-lto-confirmation-summary.json',
  'baseline.json', 'baseline-provenance.json', 'confirmation-protocol.json', 'finish-selection.json',
  'diagnostics.json', 'native-delete-decision.json', 'native-insert-decision.json',
  'native-tests/summary.json', 'pbo-tests/results.json', 'purust-cleanup-results.json',
  'native-delete-contract/summary.json', 'native-delete-tests/summary.json', 'native-insert-contract/summary.json',
  'native-delete-pbo/results.json', 'native-insert-pbo/results.json', 'native-delete-maps-delete.rs',
  'production/results.json', 'production/hosts/results.json', 'verification-final.json',
  'profile-baseline/summary.json', 'profile-baseline-all/summary.json',
  'profile-final-all/summary.json', 'profile-final-optimize/summary.json',
  'candidates/qualified/build.json',
  ...finalCampaigns.map(([name]) => name + '-runs/results.json'),
  ...Object.keys(selected).map(name => name + '-runs/results.json'),
];
for (const optional of ['source-integration.json', 'thin-lto.json', 'integration.json',
  'candidates/thin-lto/build.json', 'candidates/thin-lto-retry/build.json', 'candidates/thin-lto-rust-lld/build.json',
  'thin-lto-runs/results.json', 'thin-lto-confirmation-runs/results.json', 'integration-runs/results.json'])
  if (existsSync(join(archive, optional))) relativeEvidence.push(optional);
const diagnosticLogs = thinLto.diagnostics.flatMap(diagnostic => diagnostic.logs);
const sourceRecords = [...new Set([...relativeEvidence.map(path => join(archive, path)),
  join(production.candidate, 'build.json'), join(archive, 'finish-confirm.json'),
  ...helperRecords.flatMap(helper => [helper.stdout, helper.stderr]), ...diagnosticLogs,
  historicalPath, historicalOptimizationPath, historicalDefaultsPath])]
  .map(path => ({ path, sha256: sha(path) }));
for (const record of sourceRecords) assert.equal(sha(record.path), record.sha256);

// ---------------------------------------------------------------------------
// Published result
// ---------------------------------------------------------------------------
const selection = {
  native_delete: { decision: nativeDeleteDecision, summary: nativeDelete.summary,
    contract: { status: nativeDeleteContract.status, failed: nativeDeleteContract.failed, skipped: nativeDeleteContract.skipped,
      operations: nativeDeleteSupplementOperations, log: join(archive, 'native-delete-contract/maps.log'),
      source: join(archive, 'native-delete-maps-delete.rs') },
    differential: { status: nativeDeleteTests.status, failed: nativeDeleteTests.failed,
      note: 'raw legacy strict-balance diagnostics; the same invariant failed on the generated oracle' },
    verification: verification.campaigns.find(checked => resolve(checked.directory) === join(archive, 'native-delete-runs')) },
  native_insert: { selection: nativeInsert.summary, confirmation: insertConfirmation.summary,
    paired: insertConfirmationPaired, decision: nativeInsertDecision,
    contract: { status: nativeInsertContract.status, failed: nativeInsertContract.failed, skipped: nativeInsertContract.skipped } },
  qualified: { summary: qualified.summary, versus_native_insert: composedPaired,
    native_qualified_pairs: qualifiedPairs },
  thin_lto: thinLto,
};
const qualification = {
  production: { candidate: production.candidate, rust_bootstrap: production.rust_bootstrap,
    go_bootstrap: production.go_bootstrap, identical_candidate_rust_files: production.identical_candidate_rust_files,
    build_profile: production.build_profile ?? profile, build_rust_sha256: production.build_rust_sha256,
    cargo_command: production.cargo_command, linker: productionLinker },
  hosts: { compiler: hosts.compiler, aff: hosts.aff, frontend: hosts.frontend, bootstrap: hosts.bootstrap,
    executables: hosts.executables, aff_checks_per_host: expectedAff.length, avar_stress_items: 1000,
    identical_files: live[0].generated.length },
  compiler_tests: { pass: compilerPass, fail: 0, skipped: 0, log: compilerTests.stdout },
  preparation_tests: { pass: preparationPass, fail: 0, skipped: 0, log: preparation.stdout },
  pbo_suites: pbo.runs.length, native_suites: native.results.length,
  parser_race: 'passed', go_memo_contract: 'passed', rust_rebuild_default: 'passed',
  cleanup: { status: cleanup.status, free_bytes_before: cleanup.free_bytes_before,
    free_bytes_after: cleanup.free_bytes_after, free_bytes_gained: cleanup.free_bytes_gained,
    protected_files_verified: cleanup.protected_files_verified,
    repository_statuses_unchanged_except_finder_metadata: cleanup.repository_statuses_unchanged_except_finder_metadata,
    targets_removed: cleanup.targets.length, purust_bins: protectedPurustBins },
  verification: { path: join(archive, 'verification-final.json'), generations: verification.generations,
    identical_files: verification.identical_files, campaigns: verification.campaigns },
};
const result = {
  status: 'passed', verified_at: new Date().toISOString(), compiler: 'gopurs', target: 'Go for every host',
  date: '2026-10-03', archive,
  scope: 'gopurs hosted in Rust, JavaScript and Go, always emitting Go; allocation-candidate composition'
    + (integration ? ' plus preserved concurrent source integration' : '') + ' on the frozen 238-module gopurs-aff',
  host: common.host, profile, linker: productionLinker,
  build_rust_sha256: production.build_rust_sha256, cargo_command: production.cargo_command,
  tast: common.tast, protocol: { scope: protocol.scope, metric: protocol.metric, primary: protocol.primary,
    common: protocol.common, defaults: protocol.defaults, resources: protocol.resources },
  executables: production.executables,
  baseline: { rust_sha256: baseline.rust_sha256, profile: baseline.profile,
    sources: baseline.sources.map(source => ({ path: source.path, files: source.files.length })),
    bootstrap: baseline.bootstrap, provenance },
  selection, final_selection: finalSelection,
  primary: { ...primary.summary, paired: primaryPaired, improvement_percent: improvement },
  common: { ...c, rust_to_go_ratio: ratio, before_rust_to_go_ratio: oldRatio, absolute_gap_reduction_percent: gap },
  defaults: { summary: defaultCampaign.summary, paired: defaultPaired, reduction_percent: defaultReduction,
    launcher_defaults: true },
  resources: resourceSummary,
  pbo_attempts: attempts,
  qualification,
  profiles: {
    note: 'instrumented stack sampling is a separate diagnostic; inclusive families overlap and are not CPU percentages',
    baseline: { all: baselineByPhase.all, optimize: baselineByPhase.optimize },
    final: { all: finalAll, optimize: finalOptimize },
  },
  historical_gains: historicalGains,
  historical_json_preserved: { path: historicalPath, sha256: sha(historicalPath), corpora: historicalJson },
  integration: integration
    ? { path: integrationPath, state: integration, runs: integrationRuns, attribution: integrationAttribution }
    : { runs: integrationRuns, attribution: integrationAttribution },
  sub_three_second: subThree,
  helpers: { selection: finish.commands.map(({ label, exit_code, stdout, stderr }) => ({ label, exit_code, stdout, stderr })),
    confirmation: finishConfirmation.commands.map(({ label, exit_code, stdout, stderr }) => ({ label, exit_code, stdout, stderr })),
    production: production.commands.map(({ label, exit_code, stdout, stderr }) => ({ label, exit_code, stdout, stderr })),
    hosts: hosts.runs.map(({ label, exit_code, stdout, stderr }) => ({ label, exit_code, stdout, stderr })) },
  source_records: sourceRecords,
};

// ---------------------------------------------------------------------------
// Markdown report
// ---------------------------------------------------------------------------
const basenameOut = '2026-10-03-gopurs-rust-allocation';
const destination = join(root, 'docs/benchmark-results', basenameOut);
assert(!existsSync(destination + '.json'));
assert(!existsSync(destination + '.md'));
const fmt = n => new Intl.NumberFormat('fr-FR', { maximumFractionDigits: 1 }).format(n);
const pct = n => fmt(n) + ' %';
const table = (headers, rows) => '| ' + headers.join(' | ') + ' |\n|' + headers.map(() => '---').join('|') + '|\n'
  + rows.map(row => '| ' + row.join(' | ') + ' |').join('\n');
const ltoLabel = profile.lto === false ? 'sans LTO' : 'ThinLTO';
const subThreeEntries = Object.entries(subThreeCampaigns);
const subThreeText = subThree.claim
  ? `Les ${subThree.samples} mesures du compilateur installé restent sous 3 000 ms dans les trois campagnes finales.`
  : subThreeBelow.length
    ? `${subThreeBelow.length}/${subThreeEntries.length} médianes et ${subThree.samples_below}/${subThree.samples} mesures finales passent sous 3 000 ms ; ce seuil n'est pas encore tenu sur tous les passages.`
    : `Les trois médianes finales restent au-dessus de 3 000 ms ; la plus basse est de `
      + `${fmt(Math.min(...subThreeEntries.map(([, ms]) => ms)))} ms.`;
const commonRows = [['js', 'JavaScript / Node'], ['go', 'Go natif'], ['rust-before', 'Rust avant'], ['rust', 'Rust final']];
const phases = Object.keys(before.phases_median_ms);
const resourceRows = [
  ['Temps réel', fmt(resourceSummary['rust-before'].medians.real_s) + ' s', fmt(resourceSummary.rust.medians.real_s) + ' s'],
  ['CPU utilisateur + système', fmt(resourceSummary['rust-before'].medians.cpu_s) + ' s', fmt(resourceSummary.rust.medians.cpu_s) + ' s'],
  ['Pic RSS', fmt(resourceSummary['rust-before'].medians.rss_bytes / 1048576) + ' Mio', fmt(resourceSummary.rust.medians.rss_bytes / 1048576) + ' Mio'],
];
const familyNames = ['PBO', 'allocation', 'Value clone', 'String clone', 'Map/Set', 'type substitution',
  'directives', 'usage validation', 'decoding', 'gopurs'];
const familyValue = (summary, key) => summary.inclusive_families[key] ?? 0;
const profileRows = familyNames.map(name => [name,
  fmt(familyValue(baselineProfileAll, name)), fmt(familyValue(finalAll, name)),
  fmt(familyValue(baselineProfile, name)), fmt(familyValue(finalOptimize, name))]);
const thinParts = [];
if (thinLto.diagnostics.some(diagnostic => diagnostic.label === 'thin-lto'))
  thinParts.push('Le premier essai `thin-lto` a échoué par disque plein (« No space left on device »).');
if (thinLto.diagnostics.some(diagnostic => diagnostic.label === 'thin-lto-retry'))
  thinParts.push("Le retry `thin-lto-retry` a échoué à la liaison : le bitcode Rust LLVM 22 n'est pas lisible par le linker Apple LLVM 17.");
if (thinLto.candidate)
  thinParts.push(`Le candidat \`${thinLto.candidate.label}\` passe avec le linker Mach-O \`ld64.lld\` fourni par la toolchain Rust (\`${thinLto.candidate.linker.option}\`, SHA-256 \`${thinLto.candidate.binary_sha256}\`).`);
if (thinLto.comparison) thinParts.push(`Comparaison à cinq tours contre baseline et qualified :

${table(['Variante', 'Médiane', 'Moyenne', 'Min–max'], Object.entries(thinLto.comparison.summary)
    .map(([key, summary]) => [key, fmt(summary.median_ms) + ' ms', fmt(summary.mean_ms) + ' ms',
      `${fmt(summary.min_ms)}–${fmt(summary.max_ms)} ms`]))}`);
else if (thinParts.length) thinParts.push('La comparaison à cinq tours reste à lancer.');
const thinText = thinParts.length ? '**ThinLTO.** ' + thinParts.join(' ') : '**ThinLTO.** Aucun essai ThinLTO dans ce lot.';
const thinConfirmationText = thinLto.confirmation ? (() => {
  const campaign = thinLto.confirmation;
  const [beforeName, afterName] = campaign.selection.variants.map(variant => variant.name);
  const before = campaign.summary[beforeName], after = campaign.summary[afterName];
  const comparison = paired(campaign, beforeName, afterName);
  return `\n\nConfirmation du profil sur les sources intégrées : **${fmt(before.median_ms)} → ${fmt(after.median_ms)} ms (−${pct(100 * (1 - after.median_ms / before.median_ms))})**, **${comparison.favorable}/${comparison.rounds} paires favorables**. Moyennes **${fmt(before.mean_ms)} → ${fmt(after.mean_ms)} ms** ; les 32 générations, chauffes comprises, sont exactes.`;
})() : '';
const text = `# Gopurs Rust — campagne des allocations

Campagne du **3 octobre 2026** sur les **238 modules Aff figés**. La composition
retenue (insertion native + comparaisons \`Qualified\` empruntées) est reconstruite
par les chemins de production et qualifiée. La confirmation finale mesure
**${fmt(before.median_ms)} → ${fmt(after.median_ms)} ms** (réduction de
${pct(improvement)}), avec **${primaryPaired.favorable}/15 paires favorables**.
Profil final : **O3, ${ltoLabel}, sans debug, Arc et mimalloc**${productionLinker ? `, linker Mach-O \`ld64.lld\` via \`${productionLinker.option}\`` : ''}.

**Repères distincts.** L'optimisation algorithmique historique du 3 octobre a
réduit le backend Rust de ${fmt(historicalOptimization.primary['rust-before'].median_ms)} à
${fmt(historicalOptimization.primary.rust.median_ms)} ms (−${fmt(historicalGains.optimization.improvement_percent)} %) ;
la correction des valeurs par défaut du launcher, mesurée sur son propre témoin, donne
${fmt(historicalGains.defaults.before_ms)} à ${fmt(historicalGains.defaults.after_ms)} ms
(−${fmt(historicalGains.defaults.reduction_percent)} %). Cette passe ne revendique
aucune addition à ces repères : elle sélectionne une composition de changements
d'allocation et en mesure l'effet propre.${integration ? ` ${integration.changes_count} fichiers externes archivés (\`source-integration.json\`) sont en plus inclus dans le candidat installé ; le delta net baseline → production n'est pas attribué aux seules allocations, et les sélections natives isolées restent mesurées sur l'ancien snapshot.` : ''} ${subThreeText}

## Comparaison commune : trois hôtes, une cible Go

${table(['Hôte de gopurs', 'Médiane', 'Moyenne', 'Min–max'],
  commonRows.map(([key, name]) => [name, fmt(c[key].median_ms) + ' ms', fmt(c[key].mean_ms) + ' ms',
    `${fmt(c[key].min_ms)}–${fmt(c[key].max_ms)} ms`]))}

Dix tours, une chauffe par variante, ordre tournant. Dans cette même campagne,
le rapport Rust/Go passe de **${oldRatio.toFixed(3)}× à ${ratio.toFixed(3)}×** ;
**${pct(gap)} de l'écart absolu est résorbé**. Rust prend encore
${pct(100 * (ratio - 1))} de plus que Go et ${pct(100 * (1 - c.rust.median_ms / c.js.median_ms))}
de moins que JavaScript.

${table(['Tour', 'JS', 'Go', 'Rust avant', 'Rust final'],
  Array.from({ length: 10 }, (_, i) => [i + 1, ...['js', 'go', 'rust-before', 'rust'].map(key => fmt(c[key].samples_ms[i]))]))}

## Confirmation avant/après

${table(['Phase', 'Rust avant', 'Rust final'],
  phases.map(key => [key, fmt(before.phases_median_ms[key]) + ' ms', fmt(after.phases_median_ms[key]) + ' ms']))}

Moyennes : **${fmt(before.mean_ms)} → ${fmt(after.mean_ms)} ms** ; plages
${fmt(before.min_ms)}–${fmt(before.max_ms)} / ${fmt(after.min_ms)}–${fmt(after.max_ms)} ms.
Écart apparié médian **${fmt(primaryPaired.median_delta_ms)} ms**, moyen
**${fmt(primaryPaired.mean_delta_ms)} ms**. Aucun échantillon n'est exclu.
Les médianes de phases ne s'additionnent pas ; la spécialisation transitive est
une sous-phase de préparation.

${table(['Paire', 'Rust avant', 'Rust final', 'Après − avant'],
  before.samples_ms.map((n, i) => [i + 1, fmt(n), fmt(after.samples_ms[i]), fmt(after.samples_ms[i] - n)]))}

## Valeurs par défaut du launcher

Cinq tours tournants sans aucun worker imposé ; seuls les sélecteurs d'hôte sont
propres aux variantes. Le launcher corrigé choisit préparation 8, PBO 8,
émission 8 et pipeline actif.

${table(['Variante', 'Médiane', 'Moyenne', 'Min–max'],
  [['rust-before', 'Rust avant'], ['rust', 'Rust final'], ['go', 'Go'], ['js', 'JavaScript']].map(([key, name]) =>
    [name, fmt(defaultCampaign.summary[key].median_ms) + ' ms', fmt(defaultCampaign.summary[key].mean_ms) + ' ms',
      `${fmt(defaultCampaign.summary[key].min_ms)}–${fmt(defaultCampaign.summary[key].max_ms)} ms`]))}

Le Rust final passe de **${fmt(defaultCampaign.summary['rust-before'].median_ms)} à
${fmt(defaultCampaign.summary.rust.median_ms)} ms** (réduction de ${pct(defaultReduction)}),
**${defaultPaired.favorable}/5 paires favorables**.

${table(['Tour', 'Rust avant', 'Rust final', 'Go', 'JS'],
  Array.from({ length: 5 }, (_, i) => [i + 1, ...['rust-before', 'rust', 'go', 'js']
    .map(key => fmt(defaultCampaign.summary[key].samples_ms[i]))]))}

## Ressources

Diagnostic séparé \`/usr/bin/time -l\`, trois paires après chauffe, processus entier :

${table(['Ressource', 'Avant', 'Final'], resourceRows)}

## Sélection et qualification

Les relectures natives et sémantiques couvrent **${native.results.length} suites
natives différentielles** (${fmt(qualifiedPairs)} paires de comparaison
\`Qualified\`) et **${pbo.runs.length} suites PBO**, plus les adaptateurs Go/JS
\`Qualified\` : tout passe.

**Suppression rejetée pour absence de gain.** Les deux premières assertions
strictes de balance échouaient aussi sur l'oracle (les jointures générées
peuvent produire un écart de hauteur de 2) ; le contrat final renforcé — forme
exacte contre l'oracle, modèle clés/valeurs et persistance — a passé
**${nativeDeleteSupplementOperations} opérations** supplémentaires, et les
${nativeDeleteDecision.semantic_suites} suites PBO du candidat passent. Les
médianes face au témoin reconstruit
(${fmt(nativeDeleteDecision.baseline_ms)} / ${fmt(nativeDeleteDecision.control_ms)} /
${fmt(nativeDeleteDecision.candidate_ms)} ms) ne montrent aucun gain : la
suppression n'est pas retenue et n'est pas présentée comme une erreur de
correction.

**Insertion native retenue.** Sélection provisoire puis confirmation :

${table(['Campagne', 'Variante', 'Médiane', 'Moyenne', 'Min–max'], [
  ...['baseline', 'native-insert'].map(key => ['sélection, 5 tours', key, fmt(nativeInsert.summary[key].median_ms) + ' ms',
    fmt(nativeInsert.summary[key].mean_ms) + ' ms', `${fmt(nativeInsert.summary[key].min_ms)}–${fmt(nativeInsert.summary[key].max_ms)} ms`]),
  ...['baseline', 'native-insert'].map(key => ['confirmation, 15 tours', key, fmt(insertConfirmation.summary[key].median_ms) + ' ms',
    fmt(insertConfirmation.summary[key].mean_ms) + ' ms', `${fmt(insertConfirmation.summary[key].min_ms)}–${fmt(insertConfirmation.summary[key].max_ms)} ms`]),
])}

La confirmation passe de ${fmt(nativeInsertDecision.median_before_ms)} à
${fmt(nativeInsertDecision.median_after_ms)} ms avec
**${nativeInsertDecision.favorable}/15 paires favorables** (écart apparié médian
${fmt(nativeInsertDecision.paired_median_delta_ms)} ms).

**Composition qualifiée.** Le candidat \`qualified\` compose l'insertion native
et les comparaisons empruntées \`Qualified Ident\` ; il bat \`native-insert\` dans
les **${composedPaired.favorable}/${composedPaired.rounds} tours** de la sélection
(écart apparié médian ${fmt(composedPaired.median_delta_ms)} ms).

${table(['Variante', 'Médiane', 'Moyenne', 'Min–max'],
  ['baseline', 'native-insert', 'qualified'].map(key => [key, fmt(qualified.summary[key].median_ms) + ' ms',
    fmt(qualified.summary[key].mean_ms) + ' ms', `${fmt(qualified.summary[key].min_ms)}–${fmt(qualified.summary[key].max_ms)} ms`]))}

${integration ? `**Intégration externe.** ${integration.changes_count} fichiers divergent du snapshot \`${integration.selection_candidate}\` : corrections concurrentes de labels Unicode (GoAst, GoConversions, Printer, RecordExprs) et générateur JSON (CoreFn/Json/Text.go). Ils sont préservés dans le candidat installé, et la campagne \`integration-runs\` compare les candidats figés à l'intégration sur les mêmes entrées. Le delta net baseline → production inclut donc ces changements archivés ; les sélections natives ci-dessus restent mesurées sur l'ancien snapshot.` : ''}

${thinText}${thinConfirmationText}

## Profils et limites

Les profils instrumentés restent **séparés des chronométrages** ; les familles
inclusives se recouvrent et ne sont pas des pourcentages CPU.
L'inlining de ThinLTO modifie aussi la visibilité des fonctions dans les piles.

${table(['Famille inclusive', 'Avant (tout)', 'Final (tout)', 'Avant (optimize)', 'Final (optimize)'], profileRows)}

Le profil final porte l'empreinte de l'exécutable de production
(\`${production.executables['gopurs-rust'].sha256}\`) ; celui du témoin, \`${baseline.rust_sha256}\`.
Objectif < 3 000 ms : ${subThreeText}

## Protocole et empreintes

- Corpus : **${common.tast.modules} modules / ${fmt(common.tast.types)} types /
  ${fmt(common.tast.bytes)} octets**, manifeste TAST
  \`${common.tast.sha256}\`.
- Machine : **${common.host.cpu}**, ${common.host.logical_cpus} cœurs logiques,
  ${fmt(common.host.memory_bytes / 1073741824)} Gio, Node ${common.host.node}.
- Jobs chargement/préparation/PBO/émission : **8/8/8/8**, \`GOPURS_PIPELINE=1\`
  pour les campagnes principales ; les tours « défaut » n'imposent rien et le
  launcher choisit seul ses limites.
- Temps \`backend total\` : chargement/tri, préparation, PBO, génération/émission
  Go, points d'entrée et attente finale des workers. Frontend, construction du
  compilateur et de l'application, exécution applicative et démarrage/sortie du
  processus sont exclus. Cache de fichiers système chaud.
- Sorties neuves et caches \`.purmeta\`/\`.cache\` supprimés à chaque passage ;
  **294 fichiers Go/manifests exacts** dans chaque génération, chauffes comprises.
- Vérification indépendante : **${verification.generations} générations /
  ${fmt(verification.identical_files)} fichiers** relus et exacts dans
  ${discoveredCampaigns.length} campagnes \`*-runs\` (sélection, confirmation,
  qualification, finale, ThinLTO${integrationRuns ? ', intégration' : ''}), chauffes comprises.
- Tests : **${compilerPass} tests compilateur**, **${preparationPass} tests
  préparation/auxiliaires**, **${pbo.runs.length} suites PBO**,
  **${native.results.length} suites natives**, zéro échec et zéro skip ;
  parseur Go \`-race\`, contrat de cache Go et rebuild Rust/Aff publics passés.
- Les trois hôtes passent **${expectedAff.length} contrôles Aff + stress AVar
  1 000 éléments**, sur le même TAST vivant et avec 294 fichiers Go identiques.
- Production : candidat \`${basename(production.candidate)}\`,
  **${production.identical_candidate_rust_files} fichiers Rust/Cargo** (dont
  \`build.rs\`, soit ${production.identical_candidate_rust_files - 1} hors
  \`build.rs\`) identiques entre candidat et production, bootstrap Purust JS/natif
  identique (${hosts.bootstrap.identical_files} fichiers), sources gopurs/PBO
  figées, profil \`${JSON.stringify(profile)}\`${productionLinker ? `, linker \`${productionLinker.path}\` (${productionLinker.driver}, \`${productionLinker.option}\`)` : ''}.
- Nettoyage Purust : audit \`passed\`, ${cleanup.targets.length} cibles Cargo
  régénérables supprimées, ${fmt(cleanup.free_bytes_gained / 1073741824)} Gio
  libérés, ${cleanup.protected_files_verified} fichiers protégés revérifiés par
  empreinte (sources et bins Purust inchangés).
- JSON → Typed AST historique conservé : 18 sorties brutes et 6 exécutables relus
  et revérifiés, portée et empreintes structurelles intactes.

## Archives et diagnostics

Archive : \`var/benchmark/gopurs-rust-allocation-20261003/\` ;
[données vérifiées et empreintes](${basenameOut}.json) ;
[harnais et commandes](../../bin/benchmark/gopurs-rust/README.md).
Les stdout/stderr des commandes de qualification sont référencés dans le JSON
(\`helpers\`) et leurs empreintes figurent dans \`source_records\`.
Les essais rejetés restent archivés : suppression native non retenue (absence
de gain, contrat renforcé passé), premier build ThinLTO interrompu par un disque
plein, second par le linker Apple LLVM 17 face au bitcode LLVM 22.
L'intégration externe du snapshot \`${integration?.selection_candidate ?? 'qualified'}\`
est conservée et mesurée séparément.

${table(['Exécutable installé', 'SHA-256'],
  Object.entries(production.executables).map(([name, info]) => [name, '`' + info.sha256 + '`']))}
`;

// ---------------------------------------------------------------------------
// Publication: JSON, report and the single README line
// ---------------------------------------------------------------------------
const readmePath = join(root, 'README.md'), readme = readFileSync(readmePath, 'utf8');
const lines = readme.split('\n'), indices = lines.map((line, index) => line.startsWith('[gopurs-aff]') ? index : -1)
  .filter(index => index >= 0);
assert.equal(indices.length, 1, 'README gopurs-aff line');
copyFileSync(readmePath, join(archive, 'README-before-publication.md'));
copyFileSync(fileURLToPath(import.meta.url), join(archive, 'publish-allocation.mjs'));
writeJson(destination + '.json', result);
writeFileSync(destination + '.md', text);
lines[indices[0]] = `[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff)  | ~ ${Math.round(c.js.median_ms)} ms | ~ ${Math.round(c.go.median_ms)} ms | [~ ${Math.round(c.rust.median_ms)} ms](docs/benchmark-results/${basenameOut}.md) <br>(/Go = ${ratio.toFixed(2)}x) (WIP)`;
writeFileSync(readmePath, lines.join('\n'));
console.log(JSON.stringify({
  primary_ms: { before: before.median_ms, after: after.median_ms },
  improvement_percent: improvement,
  common: Object.fromEntries(['js', 'go', 'rust-before', 'rust'].map(key => [key, c[key].median_ms])),
  defaults_reduction_percent: defaultReduction,
  qualified_versus_native_insert: composedPaired,
  profile,
  linker: productionLinker,
  integration_changes: integration ? integration.changes_count : 0,
  integration_runs: integrationRuns ? 'present' : 'missing',
  thin_lto_candidate: thinLto.candidate ? thinLto.candidate.binary_sha256 : null,
  thin_lto_measured: thinLto.measured.length,
  sub_three_second_claim: subThree.claim,
  generations_verified: verification.generations,
  files_verified: verification.identical_files,
  campaigns: discoveredCampaigns,
  destination: destination + '.md',
}, null, 2));
