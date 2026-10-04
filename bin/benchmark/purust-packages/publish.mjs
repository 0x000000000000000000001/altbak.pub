// Audit all timed outputs and derive the complete Rust-target table from integer-ms medians.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), archive = resolve(process.argv[2]);
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const originalPlans = JSON.parse(readFileSync(join(campaign.baseline_archive, 'campaign.json'))).plans;
const groups = ['libraries', 'b8x'].map(name => {
  const final = join(archive, 'libraries-final-results.json');
  const path = name === 'libraries' && existsSync(final) ? final : join(archive, name + '-results.json');
  const result = JSON.parse(readFileSync(path));
  assert.equal(result.status, 'passed'); return { path, sha256: hash(readFileSync(path)), result };
});
const validationPath = join(archive, 'validation/results.json'), validation = JSON.parse(readFileSync(validationPath));
assert.equal(validation.status, 'passed'); assert.equal(validation.results.length, 57);
const b8xValidationPath = join(archive, 'b8x-validation/results.json');
const b8xValidation = JSON.parse(readFileSync(b8xValidationPath)); assert.equal(b8xValidation.status, 'passed');
const qualification = JSON.parse(readFileSync(campaign.qualification.path)); assert.equal(qualification.status, 'passed');
assert.equal(hash(readFileSync(campaign.qualification.path)), campaign.qualification.sha256);
for (const compiler of Object.values(campaign.compilers)) assert.deepEqual(manifest(compiler.directory), compiler.files);
for (const source of campaign.source_provenance) {
  assert.deepEqual(manifest(join(archive, 'sources', source.name)), source.files);
  assert.deepEqual(source.files, qualification.sources[source.name === 'purust/purust' ? 'compiler' : 'optimizer']);
}
for (const file of campaign.family_sources) assert.equal(hash(readFileSync(join(archive, 'family', file.path))), file.sha256);
const results = groups.flatMap(group => group.result.results);
assert.equal(results.length, 58); assert.equal(new Set(results.map(result => result.name)).size, 58);
const hosts = ['js', 'rust'];
for (const result of results) {
  assert.equal(result.status, 'passed'); assert(result.inputs_unchanged);
  assert.equal(hash(readFileSync(result.definition)), result.definition_sha256);
  const definition = JSON.parse(readFileSync(result.definition));
  assert.deepEqual(manifest(definition.input), definition.input_manifest);
  if (result.name !== 'b8x') {
    const plan = originalPlans.find(plan => plan.name === result.name); assert(plan);
    assert.deepEqual(definition.invocation, ['--source', 'output', '--main', plan.main, ...plan.flags]);
  }
  for (const host of hosts) {
    const runs = result.runs.filter(run => run.host === host), samples = runs.filter(run => run.measured).map(run => run.phases_ms['backend total']);
    assert.equal(runs.length, 6); assert.equal(samples.length, 5);
    assert.deepEqual(samples, result.summary[host].samples_ms);
    assert.equal(samples.toSorted((a, b) => a - b)[2], result.summary[host].median_ms);
  }
  const canonical = manifest(result.runs[0].retained_output);
  for (const run of result.runs) {
    assert.equal(run.exit_code, 0); assert.deepEqual(run.different_files, []);
    const files = JSON.parse(readFileSync(run.generated_manifest));
    assert.equal(hash(JSON.stringify(files)), run.generated_sha256);
    assert.equal(files.length, run.identical_files); assert.deepEqual(files, canonical);
    const stderr = readFileSync(run.stderr, 'utf8');
    assert.equal(Number(stderr.match(/^\[purust\] backend total: (\d+) ms$/m)?.[1]), run.phases_ms['backend total']);
  }
  if (result.name !== 'b8x') {
    const application = validation.results.find(item => item.name === result.name); assert.equal(application?.status, 'passed');
    assert.deepEqual(application.generated_manifest, canonical);
    assert.equal(hash(readFileSync(application.binary.path)), application.binary.sha256);
  }
}
const libraries = results.filter(result => result.name.startsWith('purust-')).sort((a, b) => a.name.localeCompare(b.name));
assert.equal(libraries.length, 57);
const b8x = results.find(result => result.name === 'b8x');
assert.deepEqual(b8xValidation.generated_manifest, manifest(b8x.runs[0].retained_output));
assert.equal(b8xValidation.definition_sha256, b8x.definition_sha256);
assert.equal(hash(readFileSync(b8xValidation.binary.path)), b8xValidation.binary.sha256);
const subtotal = Object.fromEntries(hosts.map(host => [host, libraries.reduce((sum, item) => sum + item.summary[host].median_ms, 0)]));
const total = Object.fromEntries(hosts.map(host => [host, subtotal[host] + b8x.summary[host].median_ms]));
const runs = results.flatMap(result => result.runs), seconds = ms => (Math.round(ms / 10) / 100).toFixed(2);
const ratio = (a, b) => (a / b).toFixed(2), link = name => `[${name}](https://github.com/0x000000000000000000001/${name})`;
const report = { schema: 1, status: 'passed', date: '2026-10-04', archive,
  protocol: campaign.protocol, host: campaign.host, compilers: campaign.compilers, qualification: campaign.qualification,
  ...(campaign.rerun ? { rerun: campaign.rerun } : {}),
  evidence: groups.map(({ path, sha256 }) => ({ path, sha256 })),
  validation: [{ path: validationPath, sha256: hash(readFileSync(validationPath)) },
    { path: b8xValidationPath, sha256: hash(readFileSync(b8xValidationPath)) }],
  coverage: { libraries: 57, projects: 58, generations: runs.length, measured_generations: runs.filter(run => run.measured).length,
    exact_files: runs.reduce((sum, run) => sum + run.identical_files, 0) },
  subtotal_ms: subtotal, total_ms: total,
  results: [b8x, ...libraries].map(result => {
    const definition = JSON.parse(readFileSync(result.definition)), tast = definition.input_manifest.filter(file => file.path.endsWith('/corefn.json'));
    return { name: result.name, scope: definition.scope, summary: result.summary,
      input: { modules: tast.length, types: definition.frontend.types, tast_bytes: tast.reduce((sum, file) => sum + file.bytes, 0), tast_sha256: hash(JSON.stringify(tast)) },
      definition: result.definition, definition_sha256: result.definition_sha256,
      generated_sha256: result.runs[0].generated_sha256, generated_files: result.runs[0].identical_files };
  }) };
assert.equal(report.coverage.generations, 696); assert.equal(report.coverage.measured_generations, 580);
const table = [
  'Benchmark   | ' + link('purust') + ' JS binary (WIP) | ' + link('purust') + ' Rust binary (WIP)',
  '----------- | --------- | ---------------------',
  `[b8x — Rust test profile](https://github.com/0x000000000000000000001/b8x.pub) | ~ ${(b8x.summary.js.median_ms / 1000).toFixed(3)} s | ~ ${(b8x.summary.rust.median_ms / 1000).toFixed(3)} s <br>(/JS = ${ratio(b8x.summary.rust.median_ms, b8x.summary.js.median_ms)}x)`,
  ...libraries.map(result => `${link(result.name)} | ~ ${result.summary.js.median_ms} ms | ~ ${result.summary.rust.median_ms} ms <br>(/JS = ${ratio(result.summary.rust.median_ms, result.summary.js.median_ms)}x)`),
  `**Total purust-*** | ~ ${seconds(subtotal.js)} s | ~ ${seconds(subtotal.rust)} s <br>(/JS = ${ratio(subtotal.rust, subtotal.js)}x)`,
  `**Total** | ~ ${seconds(total.js)} s | ~ ${seconds(total.rust)} s <br>(/JS = ${ratio(total.rust, total.js)}x)`,
].join('\n');
const note = `> Purust timings: median of five backend-only runs per host after one warmup, on frozen inputs with byte-identical Rust/Cargo output. Totals are sums of medians. b8x uses the Rust Core/Infra/Util test profile (${report.results[0].input.modules.toLocaleString('en-US')} modules), a different corpus from the full Go-target row. [Protocol, scopes and validation](docs/benchmark-results/2026-10-04-purust-packages.md).`;
const path = join(site, 'README.md'), current = readFileSync(path, 'utf8'), marker = '#### ... to Rust\n';
const before = readFileSync(join(archive, 'README-before.md'), 'utf8');
assert.equal(current.split(marker).length, 2); assert.equal(current.split(marker)[1], before.split(marker)[1], 'Rust table changed during campaign; reconcile first');
const next = current.split(marker)[0] + marker + '\n' + table + '\n\n' + note + '\n\nMore to come...\n';
report.publication = { readme_before_sha256: hash(current), readme_after_sha256: hash(next), project_cells: 116, total_cells: 4 };
const document = `# Complete Purust compilation table — 4 October 2026

The **57 local Purust libraries** and the **b8x Rust test profile** are measured
under the JavaScript-hosted and Rust-hosted versions of Purust. Both generate Rust.

## Results

| Backend-only median | JavaScript host | Rust host | Rust / JS |
| --- | ---: | ---: | ---: |
| b8x Rust test profile | ${seconds(b8x.summary.js.median_ms)} s | ${seconds(b8x.summary.rust.median_ms)} s | ${ratio(b8x.summary.rust.median_ms, b8x.summary.js.median_ms)}x |
| 57 libraries, sum of medians | ${seconds(subtotal.js)} s | ${seconds(subtotal.rust)} s | ${ratio(subtotal.rust, subtotal.js)}x |
| Complete table, sum of medians | ${seconds(total.js)} s | ${seconds(total.rust)} s | ${ratio(total.rust, total.js)}x |

## Protocol and scope

- ${campaign.host.cpu}, ${campaign.host.logical_cpus} logical CPUs, ${campaign.host.memory_bytes / 1024 ** 3} GiB RAM; Node ${campaign.host.node}; no explicit CPU affinity.
- One warmup, then five measured fresh processes per host. Projects and hosts
  run serially, alternating JS/Rust order. Public worker defaults: JavaScript
  1/1; native budget 8, split into four PBO and four code-generation workers.
- The internal **backend total** includes TAST loading/sorting, preparation,
  optimization, generation/emission and final drain. Frontend work, application
  and compiler builds, application execution and process startup/exit are excluded.
  Overlapping phase clocks are retained individually and never added together.
- Fresh generated output and PBO caches; warm operating-system filesystem cache.
  Each library retains its primary runner's threaded/non-threaded generation mode.
- The frontend is rebuilt from frozen native-family sources. Isolated Spago
  configurations resolve the Purust ports, including configurations inherited
  from Go projects. Root package names are normalized to their dependency names.
- Argonaut Core uses its existing compact-DOM fixture; Argonaut Codecs uses its
  existing typed JSON plans fixture. ArrayBuffer Types has only foreign type
  declarations and uses an explicit minimal executable harness.
- b8x is the existing **Core/Infra/Util Rust test profile**, not the larger
  Go-target application corpus. Its actual import closure and native FFI are frozen.
  Module/type counts and per-row scopes are recorded below and in the JSON.
- All measurements use the qualified compiler rebuild after replacing the
  backtracking foreign-type scan with a linear-time regex engine and preserving
  IEEE negative zero in PBO's native constant evaluator. Module roots are sorted
  canonically so spec registration and Cargo dependencies have identical order.
  The original failures, red/green regressions, 720 module permutations,
  self-hosting identity and fresh-project checks are retained with compiler hashes.
- The initial preparation's broken historical hello-world symlink and the
  Promise workspace's duplicate registry modules are recorded as harness issues.
  Spec's required BuildInfo module is produced by Spago before freezing its TAST.
  Earlier attempts and their statuses remain archived; the publication uses a
  complete new two-host campaign with the corrected configuration and compiler.
  A retry restored the original threaded-mode flags from the frozen runner plans;
  affected cases were re-measured, while correctly configured series were retained.

## Validation and evidence

- **${report.coverage.generations} exact-output-checked generations**, including
  **${report.coverage.measured_generations} measured runs**, checking
  **${report.coverage.exact_files.toLocaleString('en-US')} generated files**.
- Every library's canonical application builds and its primary executable
  checks pass outside timing. Existing golden-output, marker and argument-scenario
  checks are retained; independently generated secondary mains are outside each
  primary compilation case. JavaScript and Rust produce byte-identical sources.
${campaign.rerun ? `- The first application pass exposed a native HTTP panic: a trailing apostrophe
  in \`HttpServer'\` was dropped by foreign-declaration scanning, and a mangled
  identifier was used to look up the source-level FFI layout. Both are corrected.
  Historical-output checks also rejected a subsequent scan that omitted the
  Unicode kind separator \`∷\`. The final scan preserves both \`::\` and \`∷\`;
  declaration/layout parity is checked over the entire frozen source corpus.
  JS/Rust differential checks, Rc/Arc carrier checks, a fresh primed-handle smoke
  test, self-hosting identity and the real HTTP/upgrade/cookie/HTTPS and BigInt suites qualify
  the rebuilt hosts before a complete new 58-project measurement campaign.
  The previous measurements and failed HTTP executable/backtrace remain archived.
  Application evidence is reused only for byte-identical Rust/Cargo source,
  with explicit provenance to the retained executable and validation record.
` : ''}- The b8x native import closure and Test.Rust.Main executable build outside
  timing; service-dependent execution is covered by the separate runtime profile.
- Frozen input and compiler manifests are checked before and after generation;
  every run retains phase samples, stdout/stderr, exit status and a file manifest.
  Identical output bytes are retained once per project, with separate failed outputs.
- Totals are **sums of per-project medians**, not a timed multi-project invocation.
- [Machine-readable report](2026-10-04-purust-packages.json)
- Local archive: \`${archive.slice(site.length + 1)}\`
- Reproduction scripts: \`bin/benchmark/purust-packages/\`

## All medians (milliseconds)

| Project | Typed modules | Types | JavaScript | Rust | Rust / JS |
| --- | ---: | ---: | ---: | ---: | ---: |
${report.results.map(result => `| ${result.name} | ${result.input.modules} | ${result.input.types} | ${result.summary.js.median_ms} | ${result.summary.rust.median_ms} | ${ratio(result.summary.rust.median_ms, result.summary.js.median_ms)}x |`).join('\n')}
`;
const json = join(site, 'docs/benchmark-results/2026-10-04-purust-packages.json'); assert(!existsSync(json));
writeFileSync(join(archive, 'README-next.md'), next); writeJson(join(archive, 'publication.json'), report);
writeFileSync(join(archive, 'report.md'), document); writeJson(json, report);
writeFileSync(join(site, 'docs/benchmark-results/2026-10-04-purust-packages.md'), document);
console.log(JSON.stringify({ coverage: report.coverage, subtotal, total, candidate: join(archive, 'README-next.md') }, null, 2));
