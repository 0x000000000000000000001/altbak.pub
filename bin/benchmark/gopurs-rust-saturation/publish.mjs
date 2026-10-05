// Derive the public table from audited production samples; keep the README edit
// as a candidate so it can be reviewed before applying it to the tracked file.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { writeJson } from '../gopurs-aff/common.mjs';
import { hash, manifest, pairedAggregate } from './common.mjs';
import { auditRuns } from './audit-runs.mjs';

const archive = resolve(process.argv[2]), site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..');
const read = path => JSON.parse(readFileSync(join(archive, path)));
const evidence = path => ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) });
const production = read('production-campaign.json'), selection = read('selection.json'), closure = read('closure.json');
for (const result of [production, selection, closure]) assert.equal(result.status, 'passed');
assert.equal(production.candidate, selection.candidate);
for (const source of selection.sources) assert.deepEqual(manifest(join(source.root, 'src')), source.files);
assert.equal(closure.decision, 'stop');
assert(new Set(closure.non_gaining_mechanisms.map(item => item.mechanism)).size >= 3);
for (const item of [...(selection.evidence ?? []), ...(closure.evidence ?? [])])
  assert.equal(hash(readFileSync(item.path)), item.sha256);
const applicationsPath = production.applications_result ?? 'production/applications/results.json';
const qualification = read('production/results.json'), applications = read(applicationsPath);
assert.equal(qualification.status, 'passed'); assert.equal(applications.status, 'passed');
assert.equal(applications.applications.length, 9); assert(applications.applications.every(item => item.status === 'passed'));
for (const tool of applications.harness ?? []) assert.equal(hash(readFileSync(tool.path)), tool.sha256);
assert.equal(read('production/contracts/results.json').status, 'passed');
assert.equal(qualification.pgo.passes.length, 3);
assert.deepEqual(manifest(production.compiler.directory), production.compiler.files);
for (const file of production.compiler.files)
  assert.equal(hash(readFileSync(join(production.compiler.origin, file.path))), file.sha256);
const labels = ['production-full-table', 'production-published-paired'];
const audits = labels.map(label => auditRuns(archive, label));
const full = read('runs/production-full-table/results.json'), paired = read('runs/production-published-paired/results.json');
assert.deepEqual(full.variants.map(item => item.name), ['js', 'go', 'rust']);
assert.deepEqual(paired.variants.map(item => item.name), ['published-rust', 'rust']);
const campaign = read('campaign.json');
assert.equal(hash(readFileSync(campaign.historical.publication)), campaign.historical.sha256);
const historical = JSON.parse(readFileSync(campaign.historical.publication));
const libraries = full.results.filter(item => item.name.startsWith('gopurs-')).sort((a, b) => a.name.localeCompare(b.name));
assert.equal(libraries.length, 50);
const b8x = full.results.find(item => item.name === 'b8x'), hosts = ['js', 'go', 'rust'];
const subtotal = Object.fromEntries(hosts.map(host => [host, libraries.reduce((sum, item) => sum + item.summary[host].median_ms, 0)]));
const total = full.total_ms;
for (const host of hosts) assert.equal(total[host], subtotal[host] + b8x.summary[host].median_ms);
const phases = ['load TAST + sort', 'prepare + monomorphize', 'PBO producer',
  'generation + writes (cumulative batches)', 'emitter drain'];
const phaseTotals = Object.fromEntries(hosts.map(host => [host, Object.fromEntries(phases.map(phase =>
  [phase, full.results.reduce((sum, item) => sum + item.summary[host].phases_median_ms[phase], 0)]))]));
assert(Object.values(phaseTotals).every(phases => Object.values(phases).every(Number.isFinite)));
const improvement = pairedAggregate(paired.results, 'published-rust', 'rust');
const improvedProjects = paired.results.filter(item => item.summary.rust.median_ms < item.summary['published-rust'].median_ms).length;
const coverage = audits.reduce((sum, audit) => ({ generations: sum.generations + audit.generations,
  measured_generations: sum.measured_generations + audit.measured_generations, exact_files: sum.exact_files + audit.exact_files }),
  { generations: 0, measured_generations: 0, exact_files: 0 });
const proposals = readdirSync(join(archive, 'proposals')).sort().map(label => {
  const proposal = read(`proposals/${label}/results.json`), buildPath = join(archive, 'candidates', label, 'build.json');
  const build = existsSync(buildPath) ? JSON.parse(readFileSync(buildPath)) : null;
  const runPath = join(archive, 'runs', label, 'results.json');
  const measured = existsSync(runPath) ? JSON.parse(readFileSync(runPath)) : null;
  return { label, status: proposal.status, reference: proposal.reference,
    evidence: evidence(`proposals/${label}/results.json`),
    source_sha256: hash(JSON.stringify(proposal.sources.map(source => source.files))),
    binary_sha256: build?.binary_sha256,
    generated_sha256: build?.generated ? hash(JSON.stringify(build.generated)) : undefined,
    failure: proposal.failure?.split('\n')[0],
    measured: measured?.status === 'passed' ? { evidence: evidence(`runs/${label}/results.json`), total_ms: measured.total_ms,
      results: measured.results.map(item => ({ name: item.name, summary: item.summary, paired: item.paired })) } : null };
});
const profiles = readdirSync(join(archive, 'profiles')).sort().flatMap(label => {
  const path = join(archive, 'profiles', label, 'summary.json');
  if (!existsSync(path)) return [];
  const profile = JSON.parse(readFileSync(path));
  assert.equal(profile.status, 'passed');
  return [{ label, path, sha256: hash(readFileSync(path)), binary_sha256: profile.binary_sha256,
    identical_go_files: profile.identical_go_files, samples: profile.samples, metric: profile.metric,
    capture_window: profile.capture_window, tool_sha256: profile.tool_sha256, sample: profile.sample }];
});
const cleanup = read('purust-reclamation.json'); assert.equal(cleanup.status, 'passed');
const profileArchives = readdirSync(archive).filter(name => /^compressed-profiles-.*\.json$/.test(name)).map(name => {
  const result = read(name); assert.equal(result.status, 'passed');
  for (const file of result.files) {
    assert(file.restored_bytes_verified && file.original_removed);
    assert.equal(hash(readFileSync(file.compressed)), file.compressed_sha256);
  }
  return { ...evidence(name), logical_bytes_reclaimed: result.logical_bytes_reclaimed, files: result.files };
});
const seconds = ms => (Math.round(ms / 10) / 100).toFixed(2), ratio = (a, b) => (a / b).toFixed(2);
const reduction = (a, b) => (100 * (1 - b / a)).toFixed(1);
const before = readFileSync(join(archive, 'README-before-production.md'), 'utf8');
assert.equal(readFileSync(join(site, 'README.md'), 'utf8'), before, 'README changed; reconcile before publishing');
const start = before.indexOf('#### ... to Go\n'), end = before.indexOf('#### ... to Rust\n', start);
assert(start >= 0 && end > start);
let section = before.slice(start, end);
const replace = (pattern, render) => {
  assert.equal([...section.matchAll(new RegExp(pattern.source, 'gm'))].length, 1);
  section = section.replace(new RegExp(pattern.source, 'm'), render);
};
for (const item of libraries) replace(new RegExp(`^(\\[${item.name}\\]\\([^\\n]+?\\))[^\\n]*$`),
  (_, link) => `${link}  | ${hosts.map(host => `~ ${item.summary[host].median_ms} ms`).join(' | ')}`);
replace(/^(\[b8x\]\([^\n]+?\))[^\n]*$/, (_, link) => `${link} | ~ ${(b8x.summary.js.median_ms / 1000).toFixed(3)} s` +
  ` | ~ ${(b8x.summary.go.median_ms / 1000).toFixed(3)} s <br>(/JS = ${ratio(b8x.summary.go.median_ms, b8x.summary.js.median_ms)}x)` +
  ` | ~ ${(b8x.summary.rust.median_ms / 1000).toFixed(3)} s <br>(/JS = ${ratio(b8x.summary.rust.median_ms, b8x.summary.js.median_ms)}x)`);
for (const [pattern, title, values] of [[/^\*\*Total gopurs-\*\*\*[^\n]*$/, '**Total gopurs-***', subtotal],
  [/^\*\*Total\*\*[^\n]*$/, '**Total**', total]])
  replace(pattern, `${title} | ~ ${seconds(values.js)} s | ~ ${seconds(values.go)} s <br>(/JS = ${ratio(values.go, values.js)}x)` +
    ` | ~ ${seconds(values.rust)} s <br>(/JS = ${ratio(values.rust, values.js)}x)`);
const reportName = '2026-10-05-gopurs-rust-saturation';
section = section.trimEnd() + `\n\n> Refreshed on **5 October 2026** after the Rust-host optimization campaign. Median of five backend-only runs per host after one warmup; totals sum project medians. All hosts generate byte-identical Go. [Protocol, qualification and results](docs/benchmark-results/${reportName}.md).\n\n`;
const next = before.slice(0, start) + section + before.slice(end);
for (const line of section.split('\n').filter(line => line.startsWith('[gopurs-'))) assert(!line.includes('/JS'));
assert.equal((section.match(/\/JS =/g) ?? []).length, 6);
const report = { schema: 1, status: 'passed', date: '2026-10-05', archive,
  protocol: campaign.protocol, compilers: production.compiler, qualification: evidence('production/results.json'),
  production_contracts: evidence('production/contracts/results.json'),
  applications: { ...evidence(applicationsPath), harness: applications.harness, results: applications.applications },
  selection: { ...selection, evidence_file: evidence('selection.json') }, closure: { ...closure, evidence_file: evidence('closure.json') },
  cleanup: { ...evidence('purust-reclamation.json'), logical_bytes_removed: cleanup.removed.reduce((sum, file) => sum + file.bytes, 0) },
  historical: campaign.historical, audits, coverage, subtotal_ms: subtotal, total_ms: total, phase_totals_ms: phaseTotals,
  recovery: production.recovery ?? null,
  application_recovery: production.application_recovery ?? null,
  published_paired: { ...improvement, improved_projects: improvedProjects }, proposals, profiles, profile_archives: profileArchives,
  results: [b8x, ...libraries].map(item => {
    const definition = read(`cases/${item.name}/definition.json`), tast = definition.input_manifest.filter(file => file.path.endsWith('/corefn.json'));
    return { name: item.name, summary: item.summary, paired: paired.results.find(result => result.name === item.name).paired.rust,
      input: { modules: tast.length, tast_bytes: tast.reduce((sum, file) => sum + file.bytes, 0), sha256: hash(JSON.stringify(tast)) },
      generated_files: definition.oracle.files.length, generated_sha256: hash(JSON.stringify(definition.oracle.files)) };
  }), publication: { readme_before_sha256: hash(before), readme_after_sha256: hash(next), project_cells: 153, total_cells: 6 } };
const publicJson = join(site, 'docs/benchmark-results', reportName + '.json');
const publicMd = join(site, 'docs/benchmark-results', reportName + '.md');
assert(!existsSync(publicJson)); assert(!existsSync(publicMd)); assert(!existsSync(join(archive, 'publication.json')));
const rows = [b8x, ...libraries].map(item => `| ${item.name} | ${hosts.map(host => item.summary[host].median_ms).join(' | ')} |`).join('\n');
const document = `# Rust-hosted gopurs saturation campaign — 4–5 October 2026

All three gopurs hosts generate **Go**. The final production table covers
**50 libraries plus b8x**, with newly built JavaScript, Go and Rust compilers.

## Qualified results

| Backend-only median | JavaScript host | Go host | Rust host |
| --- | ---: | ---: | ---: |
| b8x → Go | ${seconds(b8x.summary.js.median_ms)} s | ${seconds(b8x.summary.go.median_ms)} s | ${seconds(b8x.summary.rust.median_ms)} s |
| 50 libraries, sum of medians | ${seconds(subtotal.js)} s | ${seconds(subtotal.go)} s | ${seconds(subtotal.rust)} s |
| Complete table, sum of medians | ${seconds(total.js)} s | ${seconds(total.go)} s | ${seconds(total.rust)} s |

The final Rust host is **${reduction(total.go, total.rust)}% below the Go host**
on the complete table. The prior published Rust total was
**${seconds(historical.total_ms.rust)} s**. A separate contemporary paired
campaign measures the frozen published Rust binary at **${seconds(improvement.before_ms)} s**
and the new production binary at **${seconds(improvement.after_ms)} s**:
**${reduction(improvement.before_ms, improvement.after_ms)}% lower**. Its paired
bootstrap ratio interval is **[${improvement.bootstrap_95.map(value => value.toFixed(4)).join(', ')}]**.
The new Rust median is lower on **${improvedProjects}/${paired.results.length} projects** in that paired campaign.

### Recorded backend phases

These values sum each project's phase median across the complete three-host
table, in seconds. Preparation includes monomorphization. PBO, generation/writes
and drain overlap; the rows must not be added to reconstruct the backend total.

| Phase, sum of project medians | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
${phases.map(phase => `| ${phase} | ${hosts.map(host => seconds(phaseTotals[host][phase])).join(' | ')} |`).join('\n')}

## Selection and stopping decision

${selection.retained.map(item => `- **${item.name}** — ${item.reason}`).join('\n')}

${(selection.rejected ?? []).length ? 'Other selection decisions:\n' + selection.rejected.map(item => `- **${item.name}** — ${item.reason}`).join('\n') : ''}

Screenings compare one mechanism with a contemporary control; their percentages
are not added together. The final composition receives its own confirmation and
fresh production PGO training. ${closure.rationale}

The last non-retained mechanisms were:
${closure.non_gaining_mechanisms.map(item => `- **${item.mechanism}** (${item.experiment}): ${item.reason}`).join('\n').replaceAll('\u001f', '`')}

Every proposal, including failed construction/test attempts, remains in the
machine-readable report and local archive with its source, generated-code and
binary fingerprints. Initial fixed-profile screening and final self-trained
PGO are recorded separately.

## Method

- Apple M4 Pro, 10 performance and 4 efficiency cores; no explicit affinity.
- Frozen TAST/FFI inputs from the qualified compilation refresh. Every generation
  is checked byte-for-byte against its canonical Go oracle; input manifests and
  sibling FFI inputs are rechecked between runs.
- One warmup and five measured fresh processes per project and host. Projects,
  builds and measurement campaigns are serialized. The three-host table cycles
  through all six host permutations; the old/new Rust comparison alternates.
- Workers **8/8/8/8**, pipeline enabled. Go alone uses
  **GOGC=off / GOMEMLIMIT=10GiB**. The public launcher still selects Go by default,
  JavaScript with GOPURS_JS=1 and Rust with GOPURS_RUST=1.
- The clock is the internal **backend total**: loading/sorting, preparation,
  monomorphization, PBO, generation/writes and drain. Frontend work, compiler and
  application builds, application execution and process startup/exit are outside
  timing. Producer, emission and drain phases overlap and are not additive.
- Each total is a sum of per-project medians. Bootstrap resampling preserves
  paired rounds within each project. Practical selection thresholds are 1% of
  the full-table total, or 3% on an important case with no significant overall
  regression, above observed variability.
${production.recovery ? `- The final three-host campaign was interrupted after ${production.recovery.retained_projects.length} complete projects. The remaining ${production.recovery.restarted_projects.length} projects were restarted with a fresh warmup and five whole rounds, using identical compiler artifacts. The interrupted records and ${production.recovery.superseded_checked_generations} superseded checked generations remain archived; they are excluded from the published medians and final coverage.` : ''}
- Diagnostic sampling covers all threads, excluding known blocking leaves.
  Inclusive categories overlap and are not CPU percentages. Instrumented runs
  are never used as selection timings.
- Before the campaign, **${report.cleanup.logical_bytes_removed.toLocaleString('en-US')} logical bytes**
  of regenerable Purust intermediates were reclaimed after inventory. Preserved
  sources, executables and historical evidence were rehashed.

## Qualification

- **${coverage.generations} final checked generations**, including
  **${coverage.measured_generations} measured runs** and
  **${coverage.exact_files.toLocaleString('en-US')} byte-exact generated files**,
  across the production table and separate old/new paired campaign.
- Production Rust: O3, ThinLTO, threaded Arc, mimalloc and fresh compiler-self
  PGO, with three checked training passes excluding Test.Main. Bootstrap outputs
  are compared with the unprofiled compiler; independent JavaScript/native Purust
  generation checks establish Rust source identity.
- Affected fold, directive visibility, monomorphization, scheduling, scope and
  native FFI contracts are retained with their exact compiler builds.
- Fresh-project, launcher, parser/race and three-host native regression fixtures
  pass. Aff is checked on all three hosts. Representative measured applications
  (Aff, Arrays, Spec, Yoga JSON, Enums, Promise, Prelude and Strings) build and run;
  all generated b8x packages and entry points build outside timing.
${production.application_recovery ? '- The first application qualification detected Yoga JSON writing a `.spec-results` report in the frozen-input working directory. All application commands succeeded and no pre-existing input changed. The report and failed qualification remain archived; the corrected harness reruns all applications in private working copies, checks Yoga JSON\'s recorded assertions and revalidates every frozen corpus. Both timing campaigns are preserved unchanged.' : ''}

## Complete table (milliseconds)

| Project | JavaScript | Go | Rust |
| --- | ---: | ---: | ---: |
${rows}

## Evidence

- [Machine-readable report](${reportName}.json)
- Local archive: var/benchmark/gopurs-rust-saturation-20261004/
- Reproduction tools: bin/benchmark/gopurs-rust-saturation/
- [Previous qualified compilation table](2026-10-04-compilation-refresh.md)
- [Earlier gopurs optimization campaign](2026-10-04-gopurs-rust-night.md)
`.replaceAll('\u001f', '`');
writeFileSync(join(archive, 'README-next.md'), next);
writeFileSync(join(archive, 'report.md'), document); writeJson(join(archive, 'publication.json'), report);
writeJson(publicJson, report); writeFileSync(publicMd, document);
console.log(JSON.stringify({ coverage, subtotal_ms: subtotal, total_ms: total, published_paired: improvement,
  readme_candidate: join(archive, 'README-next.md'), readme_sha256: hash(next) }, null, 2));
