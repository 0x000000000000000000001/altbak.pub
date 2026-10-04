// Independently verify every applied table cell against the retained measurement records.
import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, writeJson } from '../gopurs-aff/common.mjs';

const script = fileURLToPath(import.meta.url), site = resolve(dirname(script), '../../..');
assert(process.argv[2], 'Usage: node audit-publication.mjs ARCHIVE');
const archive = resolve(process.argv[2]), output = join(archive, 'publication-audit.json');
assert(!existsSync(output), 'Retain the existing publication audit');
const reportPath = join(archive, 'publication.json'), report = JSON.parse(readFileSync(reportPath));
assert.equal(report.status, 'passed');
const publicJson = join(site, `docs/benchmark-results/${report.date}-purust-packages.json`);
const publicMarkdown = join(site, `docs/benchmark-results/${report.date}-purust-packages.md`);
assert(readFileSync(publicJson).equals(readFileSync(reportPath)));
assert(readFileSync(publicMarkdown).equals(readFileSync(join(archive, 'report.md'))));
const records = report.evidence.flatMap(evidence => {
  const bytes = readFileSync(evidence.path); assert.equal(hash(bytes), evidence.sha256);
  const group = JSON.parse(bytes); assert.equal(group.status, 'passed'); return group.results;
});
for (const evidence of [...report.validation, report.qualification]) {
  const bytes = readFileSync(evidence.path); assert.equal(hash(bytes), evidence.sha256);
  assert.equal(JSON.parse(bytes).status, 'passed');
}
assert.equal(records.length, 58); assert.equal(new Set(records.map(record => record.name)).size, 58);
const hosts = ['js', 'rust'], subtotal = { js: 0, rust: 0 }, total = { js: 0, rust: 0 };
let generations = 0, measured = 0, exactFiles = 0;
for (const row of report.results) {
  const record = records.find(record => record.name === row.name); assert.equal(record?.status, 'passed');
  for (const run of record.runs) {
    assert.equal(run.exit_code, 0); assert.deepEqual(run.different_files, []);
    const stderr = readFileSync(run.stderr, 'utf8');
    const phases = Object.fromEntries([...stderr.matchAll(/^\[purust\] (.+): (\d+) ms$/gm)]
      .map(([, name, ms]) => [name, Number(ms)]));
    assert.deepEqual(phases, run.phases_ms);
    for (const phase of ['load TAST + sort', 'prepare', 'optimize + generate', 'finalize + emit', 'backend total']) {
      assert(Number.isInteger(phases[phase]) && phases[phase] >= 0, `${record.name}: missing ${phase}`);
    }
    generations++; measured += Number(run.measured); exactFiles += run.identical_files;
  }
  for (const host of hosts) {
    const runs = record.runs.filter(run => run.host === host); assert.equal(runs.length, 6);
    const samples = runs.filter(run => run.measured).map(run => run.phases_ms['backend total']);
    assert.equal(samples.length, 5); assert.deepEqual(samples, row.summary[host].samples_ms);
    const median = samples.toSorted((a, b) => a - b)[2]; assert.equal(median, row.summary[host].median_ms);
    total[host] += median;
    if (row.name !== 'b8x') subtotal[host] += median;
  }
}
assert.deepEqual(report.coverage, { libraries: 57, projects: 58, generations, measured_generations: measured, exact_files: exactFiles });
assert.equal(generations, 696); assert.equal(measured, 580);
assert.deepEqual(subtotal, report.subtotal_ms); assert.deepEqual(total, report.total_ms);
const readmePath = join(site, 'README.md'), readme = readFileSync(readmePath, 'utf8');
assert.equal(readme, readFileSync(join(archive, 'README-next.md'), 'utf8'));
assert.equal(hash(readme), report.publication.readme_after_sha256);
const marker = '#### ... to Rust\n', before = readFileSync(join(archive, 'README-before.md'), 'utf8');
assert.equal(readme.split(marker).length, 2); assert.equal(before.split(marker).length, 2);
assert.equal(readme.split(marker)[0], before.split(marker)[0], 'Earlier README sections changed');
const section = readme.split(marker)[1];
const rows = section.split('\n').filter(line => /^(\[|\*\*Total)/.test(line)).map(line => line.split('|').map(cell => cell.trim()));
assert.equal(rows.length, 60); assert(rows.every(row => row.length === 3));
assert.equal(report.results[0].name, 'b8x');
const libraryNames = report.results.slice(1).map(row => row.name);
assert.equal(libraryNames.length, 57); assert(libraryNames.every(name => name.startsWith('purust-')));
assert.deepEqual(libraryNames, [...new Set(libraryNames)].sort((a, b) => a.localeCompare(b)));
const checks = [];
for (let index = 0; index < rows.length; index++) {
  const [label, ...cells] = rows[index], summary = report.results[index];
  const name = index < 58 ? summary.name : index === 58 ? '**Total purust-***' : '**Total**';
  if (index < 58) assert.equal(label, index === 0
    ? '[b8x — Rust test profile](https://github.com/0x000000000000000000001/b8x.pub)'
    : `[${name}](https://github.com/0x000000000000000000001/${name})`);
  else assert.equal(label, name);
  const values = index < 58 ? Object.fromEntries(hosts.map(host => [host, summary.summary[host].median_ms]))
    : index === 58 ? subtotal : total;
  for (const [hostIndex, host] of hosts.entries()) {
    const time = index === 0 ? `${(values[host] / 1000).toFixed(3)} s`
      : index < 58 ? `${values[host]} ms` : `${(Math.round(values[host] / 10) / 100).toFixed(2)} s`;
    const expected = `~ ${time}` + (host === 'rust' ? ` <br>(/JS = ${(values.rust / values.js).toFixed(2)}x)` : '');
    assert.equal(cells[hostIndex], expected, `${name} ${host}`);
  }
  checks.push({ name, median_ms: values, rust_js_ratio: (values.rust / values.js).toFixed(2) });
}
assert(section.includes('Totals are sums of medians.'));
assert(section.includes('Rust Core/Infra/Util test profile'));
assert(section.includes(`docs/benchmark-results/${report.date}-purust-packages.md`));
assert.equal(report.publication.project_cells, 116); assert.equal(report.publication.total_cells, 4);
const audit = { status: 'passed', checked_at: new Date().toISOString(), archive,
  script: { path: script, sha256: hash(readFileSync(script)) },
  readme: { path: readmePath, sha256: hash(readme), preceding_sections_sha256: hash(readme.split(marker)[0]) },
  reports: [reportPath, publicJson, publicMarkdown].map(path => ({ path, sha256: hash(readFileSync(path)) })),
  coverage: report.coverage, project_cells: 116, total_cells: 4, ratios: 60, rows: checks };
writeJson(output, audit);
console.log(JSON.stringify({ status: audit.status, project_cells: 116, total_cells: 4, ratios: 60, subtotal, total, audit: output }, null, 2));
