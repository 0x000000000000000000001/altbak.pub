// Verify the applied README candidate, public reports and preserved historical evidence.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..'), archive = resolve(process.argv[2]);
const publication = JSON.parse(readFileSync(join(archive, 'publication.json')));
const campaign = JSON.parse(readFileSync(join(archive, 'campaign.json')));
const gates = [];
const gate = (name, action) => { action(); gates.push(name); };
const readme = readFileSync(join(site, 'README.md'), 'utf8');
const before = readFileSync(join(archive, 'README-before.md'), 'utf8');
gate('README hashes and exact generated candidate', () => {
  assert.equal(hash(before), publication.publication.readme_before_sha256);
  assert.equal(hash(readme), publication.publication.readme_after_sha256);
  assert.equal(readme, readFileSync(join(archive, 'README-next.md'), 'utf8'));
});
gate('Only requested table rows and protocol note changed', () => {
  assert.equal(readme.split('### Compilation times...')[0], before.split('### Compilation times...')[0]);
  assert.equal(publication.publication.changed_table_lines.length, 54);
  const goTable = readme.split('#### ... to Go')[1].split('#### ... to Rust')[0];
  assert.equal((goTable.match(/^\[(b8x|gopurs-[a-z-]+)\]/gm) ?? []).length, 51);
  assert(!goTable.includes('(failed)'));
  for (const line of goTable.split('\n').filter(line => /^\[|^\*\*Total/.test(line))) {
    assert.equal((line.match(/\| ~ \d+(?:\.\d+)? (?:ms|s)/g) ?? []).length, 3, line);
  }
});
gate('Public report matches complete local publication', () => {
  assert.deepEqual(JSON.parse(readFileSync(join(site, 'docs/benchmark-results/2026-10-04-compilation-refresh.json'))), publication);
  assert.equal(readFileSync(join(site, 'docs/benchmark-results/2026-10-04-compilation-refresh.md'), 'utf8'), readFileSync(join(archive, 'report.md'), 'utf8'));
});
gate('930 generations, 775 measured, five samples per cell', () => {
  assert.equal(publication.coverage.generations, 930); assert.equal(publication.coverage.measured_generations, 775);
  assert.equal(publication.results.length, 52);
  for (const result of publication.results) for (const summary of Object.values(result.summary)) {
    assert.equal(summary.samples_ms.length, 5);
    assert.equal(summary.median_ms, summary.samples_ms.toSorted((a, b) => a - b)[2]);
  }
});
gate('All three totals derived from measured medians', () => {
  for (const host of ['js', 'go', 'rust']) {
    const subtotal = publication.results.filter(result => result.name.startsWith('gopurs-')).reduce((sum, result) => sum + result.summary[host].median_ms, 0);
    assert.equal(subtotal, publication.subtotal_ms[host]);
    assert.equal(subtotal + publication.results.find(result => result.name === 'b8x').summary[host].median_ms, publication.total_ms[host]);
  }
});
gate('Installed compiler artifacts remain the measured artifacts', () => {
  for (const compiler of Object.values(campaign.compilers)) {
    assert.deepEqual(manifest(compiler.directory), compiler.files);
    for (const file of compiler.files) assert.equal(hash(readFileSync(join(compiler.origin, file.path))), file.sha256);
  }
});
gate('Earlier qualified build and library evidence preserved', () => {
  assert.equal(hash(readFileSync(campaign.qualification.path)), campaign.qualification.sha256);
  assert.equal(hash(readFileSync(campaign.qualification.libraries)), campaign.qualification.libraries_sha256);
  const original = JSON.parse(readFileSync(join(site, 'var/benchmark/gopurs-packages-rust-20261004/results.json')));
  assert.equal(original.status, 'partial');
  const correction = join(site, 'var/benchmark/gopurs-packages-fixes-20261004');
  for (const path of ['recheck/results.json', 'revision2/recheck/results.json']) assert.equal(JSON.parse(readFileSync(join(correction, path))).status, 'failed');
  if (publication.restart) assert.equal(hash(readFileSync(publication.restart.interrupted)), publication.restart.interrupted_sha256);
});
gate('Current repository whitespace checks', () => {
  execFileSync('git', ['diff', '--check'], { cwd: site });
  execFileSync('git', ['diff', '--check'], { cwd: resolve(site, '../purust/purust') });
});
writeJson(join(archive, 'publication-audit.json'), { status: 'passed', at: new Date().toISOString(), gates,
  readme_sha256: hash(readme), report_sha256: hash(readFileSync(join(site, 'docs/benchmark-results/2026-10-04-compilation-refresh.json'))) });
console.log(`Publication audit passed: ${gates.length} gates; ${publication.coverage.generations} generations / ${publication.coverage.exact_files} exact files.`);
