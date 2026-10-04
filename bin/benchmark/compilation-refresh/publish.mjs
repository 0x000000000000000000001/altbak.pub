// Derive every requested README cell and the public report from checked contemporary samples.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), archive = resolve(process.argv[2]);
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const groups = ['libraries', 'extra'].map(name => {
  const path = join(archive, `${name}-results.json`), result = JSON.parse(readFileSync(path));
  assert.equal(result.status, 'passed'); return { path, sha256: hash(readFileSync(path)), result };
});
const validationPath = join(archive, 'validation/results.json'), validation = JSON.parse(readFileSync(validationPath));
assert.equal(validation.status, 'passed'); assert(validation.compilers_unchanged);
const verifiedPath = resolve(dirname(campaign.qualification.path), 'verification-final.json');
const verified = JSON.parse(readFileSync(verifiedPath)); assert.equal(verified.status, 'passed');
for (const source of campaign.provenance) {
  assert.deepEqual(manifest(join(archive, 'sources', source.repository)), source.files);
  if (verified.sources[source.repository]) assert.deepEqual(source.files, verified.sources[source.repository]);
}
const results = groups.flatMap(group => group.result.results);
assert.equal(results.length, 52); assert.equal(new Set(results.map(result => result.name)).size, 52);
for (const result of results) {
  assert.equal(result.status, 'passed'); assert(result.inputs_unchanged);
  const definition = JSON.parse(readFileSync(result.definition));
  assert.equal(hash(readFileSync(result.definition)), result.definition_sha256);
  assert.deepEqual(manifest(definition.input), definition.input_manifest);
  for (const file of definition.sibling_inputs) assert.equal(hash(readFileSync(join(dirname(definition.input), file.path))), file.sha256);
  for (const host of definition.hosts) {
    const runs = result.runs.filter(run => run.host === host), samples = runs.filter(run => run.measured).map(run => run.phases_ms['backend total']);
    assert.equal(runs.length, 6); assert.equal(samples.length, 5);
    assert.deepEqual(samples, result.summary[host].samples_ms);
    assert.equal(result.summary[host].median_ms, samples.toSorted((a, b) => a - b)[2]);
  }
  for (const run of result.runs) {
    assert.equal(run.exit_code, 0); assert.deepEqual(run.different_files, []);
    const files = JSON.parse(readFileSync(run.generated_manifest));
    assert.equal(hash(JSON.stringify(files)), run.generated_sha256); assert.equal(files.length, run.identical_files);
    assert.deepEqual(manifest(run.retained_output), files);
  }
}
for (const compiler of Object.values(campaign.compilers)) {
  assert.deepEqual(manifest(compiler.directory), compiler.files);
  for (const file of compiler.files) assert.equal(hash(readFileSync(join(compiler.origin, file.path))), file.sha256);
}
const libraries = results.filter(result => result.name.startsWith('gopurs-')).sort((a, b) => a.name.localeCompare(b.name));
assert.equal(libraries.length, 50);
const b8x = results.find(result => result.name === 'b8x'), purust = results.find(result => result.name === 'purust-aff');
const hosts = ['js', 'go', 'rust'];
const subtotal = Object.fromEntries(hosts.map(host => [host, libraries.reduce((sum, result) => sum + result.summary[host].median_ms, 0)]));
const total = Object.fromEntries(hosts.map(host => [host, subtotal[host] + b8x.summary[host].median_ms]));
const runs = results.flatMap(result => result.runs);
assert.equal(runs.length, 930);
const report = { schema: 1, status: 'passed', date: '2026-10-04', archive,
  protocol: campaign.protocol, host: campaign.host, compilers: campaign.compilers, qualification: campaign.qualification,
  evidence: groups.map(({ path, sha256 }) => ({ path, sha256 })), validation: { path: validationPath, sha256: hash(readFileSync(validationPath)), ...validation },
  coverage: { gopurs_libraries: 50, gopurs_projects: 51, purust_projects: 1, generations: runs.length,
    measured_generations: runs.filter(run => run.measured).length, exact_files: runs.reduce((sum, run) => sum + run.identical_files, 0) },
  subtotal_ms: subtotal, total_ms: total,
  results: [b8x, ...libraries, purust].map(result => {
    const definition = JSON.parse(readFileSync(result.definition));
    const tast = definition.input_manifest.filter(file => file.path.endsWith('/corefn.json'));
    return { name: result.name, family: result.family,
      summary: result.summary, definition: result.definition, definition_sha256: result.definition_sha256,
      input: { modules: tast.length, tast_bytes: tast.reduce((sum, file) => sum + file.bytes, 0), tast_sha256: hash(JSON.stringify(tast)) },
      generated_sha256: result.runs[0].generated_sha256, generated_files: result.runs[0].identical_files };
  }) };
const restartPath = join(archive, 'restart.json');
if (existsSync(restartPath)) {
  report.restart = JSON.parse(readFileSync(restartPath));
  assert.equal(hash(readFileSync(report.restart.interrupted)), report.restart.interrupted_sha256);
}
const seconds = ms => (Math.round(ms / 10) / 100).toFixed(2), ratio = (left, right) => (left / right).toFixed(2);
const before = readFileSync(join(archive, 'README-before.md'), 'utf8');
assert.equal(readFileSync(join(site, 'README.md'), 'utf8'), before, 'README changed during campaign; reconcile before publication');
let next = before;
const replace = (pattern, render) => {
  const matches = [...next.matchAll(new RegExp(pattern.source, 'gm'))]; assert.equal(matches.length, 1, pattern.source);
  next = next.replace(new RegExp(pattern.source, 'm'), render);
};
for (const result of libraries) replace(new RegExp(`^(\\[${result.name}\\]\\([^\\n]+?\\))[^\\n]*$`),
  (_, link) => `${link}  | ${hosts.map(host => `~ ${result.summary[host].median_ms} ms`).join(' | ')}`);
replace(/^(\[b8x\]\([^\n]+?\)) \| ~[^\n]*$/,
  (_, link) => `${link} | ~ ${(b8x.summary.js.median_ms / 1000).toFixed(3)} s | ~ ${(b8x.summary.go.median_ms / 1000).toFixed(3)} s <br>(/JS = ${ratio(b8x.summary.go.median_ms, b8x.summary.js.median_ms)}x) | ~ ${(b8x.summary.rust.median_ms / 1000).toFixed(3)} s`);
replace(/^\*\*Total gopurs-\*\*\*[^\n]*$/,
  `**Total gopurs-*** | ~ ${seconds(subtotal.js)} s | ~ ${seconds(subtotal.go)} s <br>(/JS = ${ratio(subtotal.go, subtotal.js)}x) | ~ ${seconds(subtotal.rust)} s`);
replace(/^\*\*Total\*\* \| ~[^\n]*$/,
  `**Total** | ~ ${seconds(total.js)} s | ~ ${seconds(total.go)} s <br>(/JS = ${ratio(total.go, total.js)}x) | ~ ${seconds(total.rust)} s`);
replace(/^(\[purust-aff\]\([^\n]+?\))[^\n]*$/,
  (_, link) => `${link}  | ~ ${purust.summary.js.median_ms} ms | ~ ${purust.summary.rust.median_ms} ms <br>(/JS = ${ratio(purust.summary.rust.median_ms, purust.summary.js.median_ms)}x)`);
const note = '> Compilation timings refreshed on **4 October 2026**: median of five backend-only runs per host after one warmup, with frozen inputs and byte-exact generated-output checks. Totals sum the displayed project medians. [Protocol and full results](docs/benchmark-results/2026-10-04-compilation-refresh.md).';
next = next.replace('\nMore to come...', `\n${note}\n\nMore to come...`);
const beforeLines = before.split('\n'), nextLines = next.split('\n');
const edits = beforeLines.flatMap((line, index) => index < beforeLines.length - 3 && line !== nextLines[index] ? [{ line: index + 1, before: line, after: nextLines[index] }] : []);
assert.equal(edits.length, 54);
report.publication = { readme_before_sha256: hash(before), readme_after_sha256: hash(next), changed_table_lines: edits };
writeFileSync(join(archive, 'README-next.md'), next);
writeJson(join(archive, 'publication.json'), report);
const rows = [b8x, ...libraries].map(result => `| ${result.name} | ${hosts.map(host => result.summary[host].median_ms).join(' | ')} |`).join('\n');
const document = `# Complete compilation-table refresh — 4 October 2026

All **153 gopurs project cells** (51 projects × three hosts), both three-column
totals, and both **purust-aff** cells are refreshed from this campaign.

## Results

| Backend-only median | JavaScript host | Go host | Rust host |
| --- | ---: | ---: | ---: |
| b8x → Go | ${seconds(b8x.summary.js.median_ms)} s | ${seconds(b8x.summary.go.median_ms)} s | ${seconds(b8x.summary.rust.median_ms)} s |
| 50 gopurs libraries → Go (sum of medians) | ${seconds(subtotal.js)} s | ${seconds(subtotal.go)} s | ${seconds(subtotal.rust)} s |
| Complete gopurs table (sum of medians) | ${seconds(total.js)} s | ${seconds(total.go)} s | ${seconds(total.rust)} s |
| purust-aff → Rust | ${purust.summary.js.median_ms} ms | — | ${purust.summary.rust.median_ms} ms |

## Protocol and provenance

- ${campaign.host.cpu}, ${campaign.host.logical_cpus} logical CPUs, ${campaign.host.memory_bytes / 1024 ** 3} GiB memory; Node ${campaign.host.node}; no explicit CPU affinity.
- One warmup and five measured fresh processes per host; serialized cases and
  six host permutations over the six rounds. The JavaScript/Rust pair alternates
  for Purust. Each cell is the median of its five successful measurements.
- The timer is the compiler's internal **backend total**: loading and sorting,
  preparation, PBO, generation/emission and final drain. Frontend compilation,
  compiler bootstrap, application builds/execution and process startup/exit are
  outside this timer. Phase samples and medians are retained in the JSON report;
  overlapping producer, cumulative emission and drain clocks are never added.
- gopurs uses workers **8/8/8/8**, pipeline enabled. The public launcher applies
  Go's **GOGC=off / GOMEMLIMIT=10GiB** policy only to the Go host. Purust uses
  its public defaults (JS 1/1; native PBO 8, generation 4). OS caches stay warm;
  generated output and PBO caches are reset between runs.
- All five compiler artifacts match the most recent qualified rebuild, including
  the division, deterministic module ordering and IEEE negative-zero fixes.
  Rust-hosted gopurs uses O3, ThinLTO, threaded Arc, mimalloc and self-trained PGO.
  Binary hashes, frozen compiler sources and earlier bootstrap/test evidence are
  linked in the JSON report and retained in the archive.
- The 50 library rows use their retained reference TAST/FFI corpora. Every host
  must match the previously qualified Go source oracle, including the explicitly
  corrected Prelude signed-zero oracle. Historical failed evidence is retained.
- b8x uses the full local Go-target application corresponding to the public
  b8x project. Its frontend is regenerated from frozen current sources because
  20 cached source inputs were stale; its automatic entry-point selection is
  retained. Purust Aff's cached frontend was verified against all 296 source
  hashes before freezing its TAST and Rust FFI.
- Every generation is checked byte-for-byte. A manifest is kept for every run;
  identical source bytes are archived once per case, with every run referring
  to that canonical output. Failed runs, if any, retain their own diagnostics
  and generated files. All frozen inputs and compilers are rechecked afterward.
- The first b8x series was interrupted by a machine/server restart with 16
  complete, output-identical generations retained. Its raw state and logs are
  archived separately. The published b8x medians use a complete new warmup and
  five-run series on the same frozen inputs; the 900 completed library runs
  retain their original, uninterrupted campaign.

## Validation

- **${report.coverage.generations} checked generations**, including **${report.coverage.measured_generations} measured runs**;
  **${report.coverage.exact_files.toLocaleString('en-US')} exact generated files** checked.
- Generated Go applications for Aff, Arrays, Enums, Promise, Prelude and Strings
  build and execute successfully outside timing. Aff prints all 45 checks.
- All generated b8x Go packages and entry points build outside timing.
- The Purust Aff application builds offline and executes all 47 checks, with
  exact expected output. JavaScript and Rust hosts generate identical Rust/Cargo files.
- The README totals are **sums of per-project medians**, not timings of one
  multi-project invocation.

## All gopurs medians (milliseconds)

| Project | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
${rows}

## Evidence

- [Machine-readable results](2026-10-04-compilation-refresh.json)
- Local archive: \`var/benchmark/compilation-refresh-20261004/\`
- Reproduction tools: \`bin/benchmark/compilation-refresh/\`
- Source qualification: [corrected package campaign](2026-10-04-gopurs-packages-fixes.md)
`;
writeFileSync(join(archive, 'report.md'), document);
const publicJson = join(site, 'docs/benchmark-results/2026-10-04-compilation-refresh.json');
const publicMd = join(site, 'docs/benchmark-results/2026-10-04-compilation-refresh.md');
assert(!existsSync(publicJson)); assert(!existsSync(publicMd));
writeJson(publicJson, report); writeFileSync(publicMd, document);
console.log(JSON.stringify({ coverage: report.coverage, subtotal, total,
  b8x: b8x.summary, purust: purust.summary, readme_edits: edits, readme_candidate: join(archive, 'README-next.md') }, null, 2));
