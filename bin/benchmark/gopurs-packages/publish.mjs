// Re-read every sample and output before preparing the README cell-only patch.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]);
const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..');
const source = join(archive, 'results.json'), state = JSON.parse(readFileSync(source));
assert(['passed', 'partial'].includes(state.status));
const followupPath = join(archive, 'followups/results.json');
const followups = existsSync(followupPath) ? JSON.parse(readFileSync(followupPath)) : null;
if (followups) {
  assert.equal(followups.status, 'completed');
  assert.equal(followups.original_results_sha256, hash(readFileSync(source)));
}
assert.deepEqual(manifest(join(archive, 'compiler')), state.compiler);
assert.equal(hash(readFileSync(join(archive, 'frontend/purs'))), state.frontend_sha256);
const rows = [], median = values => { const xs = values.toSorted((a, b) => a - b);
  return xs.length % 2 ? xs[xs.length >> 1] : (xs[xs.length / 2 - 1] + xs[xs.length / 2]) / 2; };
let generations = 0, identical = 0, diagnosticRuns = 0;
for (const previous of state.results) {
  const replaced = previous.status !== 'passed';
  const result = replaced ? followups?.results.find(result => result.name === previous.name) : previous;
  assert(result, 'Missing resolution: ' + previous.name);
  assert(['passed', 'output_mismatch', 'backend_failed'].includes(result.status));
  const directory = join(archive, replaced ? 'followups' : 'packages', result.name);
  const evidence = replaced ? followupPath : join(directory, 'results.json');
  if (!replaced) assert.deepEqual(JSON.parse(readFileSync(evidence)), result);
  assert.deepEqual(manifest(join(directory, 'input')), replaced ? result.inputs : result.inputs.manifest);
  const oracle = result.oracle ?? result.runs[0];
  assert.equal(oracle.host, 'go'); assert.equal(oracle.exit_code, 0);
  const samples = [];
  for (const run of result.runs) {
    const clocks = [...readFileSync(run.stderr, 'utf8').matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
    assert.deepEqual(manifest(run.output), run.generated);
    if (run.exit_code === 0) {
      assert.equal(clocks.length, 1); assert.equal(Number(clocks[0][1]), run.phases_ms['backend total']);
    }
    if (result.status === 'passed') {
      assert.equal(run.exit_code, 0);
      assert.deepEqual(run.generated, oracle.generated);
      if (run.host === 'rust') assert.equal(run.identical_files, run.generated.length);
      generations++; identical += run.generated.length;
      if (run.host === 'rust' && !run.warmup) samples.push(Number(clocks[0][1]));
    } else diagnosticRuns++;
  }
  if (result.status === 'passed') {
    assert.equal(samples.length, state.protocol.rounds);
    assert.equal(median(samples), result.summary.median_ms);
  } else {
    const rust = result.runs.find(run => run.host === 'rust'); assert(rust);
    if (result.status === 'output_mismatch') {
      assert.equal(rust.exit_code, 0); assert(rust.different_files.length > 0);
      assert.notDeepEqual(rust.generated, oracle.generated);
    } else {
      assert.notEqual(rust.exit_code, 0);
      assert(readFileSync(rust.stderr, 'utf8').includes('attempt to divide by zero'));
    }
  }
  rows.push({ package: result.name, status: result.status, modules: previous.tast.modules, types: previous.tast.types,
    median_ms: samples.length ? median(samples) : null, samples_ms: samples,
    min_ms: samples.length ? Math.min(...samples) : null, max_ms: samples.length ? Math.max(...samples) : null,
    identical_files_per_generation: result.status === 'passed' ? oracle.generated.length : null,
    oracle_host: oracle.host, readme_cell: result.readme_cell ?? null,
    different_files: result.different_files ?? [], original_failure: replaced ? previous.failure : null,
    evidence, evidence_sha256: hash(readFileSync(evidence)) });
}
assert.deepEqual(rows.map(row => row.package), state.packages);
const readmePath = join(site, 'README.md'), readme = readFileSync(readmePath, 'utf8'), lines = readme.split('\n');
const edits = [];
for (const row of rows) {
  const indices = lines.map((line, index) => line.startsWith(`[${row.package}]`) ? index : -1).filter(index => index >= 0);
  assert.equal(indices.length, 1); const index = indices[0], before = lines[index];
  assert(/\| \(WIP\)$/.test(before), `Rust cell changed independently: ${row.package}`);
  const cell = row.status === 'passed' ? `~ ${Math.round(row.median_ms)} ms` : row.readme_cell;
  assert(cell);
  const after = before.replace(/\| \(WIP\)$/, `| ${cell}`);
  assert.deepEqual(before.split('|').slice(0, 3), after.split('|').slice(0, 3));
  edits.push({ line: index + 1, before, after }); lines[index] = after;
}
const allRows = lines.filter(line => /^\[gopurs-[a-z0-9-]+\]/.test(line));
const successfulRows = allRows.filter(line => /\| ~ (\d+) ms$/.test(line));
const displayedTotal = successfulRows.reduce((total, line) => total + Number(line.match(/\| ~ (\d+) ms$/)[1]), 0);
const totalIndex = lines.findIndex(line => line.startsWith('**Total gopurs-***'));
assert(totalIndex >= 0); const before = lines[totalIndex]; assert(/\| \(WIP\)$/.test(before));
const coverage = successfulRows.length === allRows.length ? '' : ` (${successfulRows.length}/${allRows.length})`;
const after = before.replace(/\| \(WIP\)$/, `| ~ ${(displayedTotal / 1000).toFixed(2)} s${coverage}`);
edits.push({ line: totalIndex + 1, before, after }); lines[totalIndex] = after;
const patch = ['*** Begin Patch', '*** Update File: ' + readmePath,
  ...edits.flatMap(edit => ['@@', '-' + edit.before, '+' + edit.after]), '*** End Patch', ''].join('\n');
writeFileSync(join(archive, 'README.patch'), patch);
const report = { schema: 'gopurs-packages-rust-1', status: rows.every(row => row.status === 'passed') ? 'passed' : 'completed_with_failures', verified_at: new Date().toISOString(),
  archive, source_sha256: hash(readFileSync(source)), protocol: state.protocol, host: state.host,
  compiler: state.compiler, frontend_sha256: state.frontend_sha256, rows,
  verified: { packages: rows.length, measured_packages: rows.filter(row => row.status === 'passed').length,
    failed_packages: rows.filter(row => row.status !== 'passed').map(row => row.package),
    generations, identical_generated_files: identical, diagnostic_runs: diagnosticRuns },
  aggregate: { packages: allRows.length, successful_packages: successfulRows.length, sum_of_displayed_rust_medians_ms: displayedTotal,
    note: 'Subtotal of successful package medians, including the previously qualified gopurs-aff README value; failed rows excluded. Not a measured multi-package invocation.' },
  readme: { before_sha256: hash(readme), after_sha256: hash(lines.join('\n')), edits } };
writeJson(join(archive, 'publication.json'), report);
const name = basenameForDate(archive);
const destination = join(site, 'docs/benchmark-results', name);
assert(!existsSync(destination + '.json') && !existsSync(destination + '.md'));
writeJson(destination + '.json', report);
writeFileSync(destination + '.md', `# Temps du compilateur gopurs hébergé en Rust — bibliothèques\n\n` +
  `Date : ${new Date().toISOString().slice(0, 10)}. ${rows.length} cellules Rust examinées, **${report.verified.measured_packages} mesurées et ${report.verified.failed_packages.length} en échec de génération ou de validation** ; tous les hôtes génèrent du **Go**.\n\n` +
  `## Protocole\n\n` +
  `- Compilateurs et entrées TAST/FFI figés dans \`${archive}\`.\n` +
  `- Frontend hors chronomètre ; mesure \`[gopurs] backend total\` incluant chargement, préparation, PBO, génération, écritures et drain.\n` +
  `- Pour chaque paquet validé : une chauffe Rust puis ${state.protocol.rounds} mesures, exécutions sérialisées ; médiane publiée.\n` +
  `- Workers chargement/préparation/PBO/émission **8/8/8/8**, pipeline actif ; sorties et caches PBO réinitialisés entre les processus.\n` +
  `- Chaque génération est comparée octet par octet à l'hôte Go sur les mêmes entrées ; éventuel repli oracle JavaScript indiqué par ligne.\n` +
  `- Les colonnes JS/Go préexistantes proviennent de campagnes antérieures. Cette extension ne mesure pas de nouveaux ratios entre hôtes.\n` +
  `- La compilation Go des applications et l'exécution de leurs tests ne font pas partie de cette campagne de génération.\n\n` +
  `Machine : ${state.host.cpu}, ${state.host.logical_cpus} processeurs logiques, ${state.host.memory_bytes / 1073741824} Gio de RAM.\n\n` +
  `## Résultats\n\n| Paquet | Modules | Types | Rust médian (ms) | Min–max (ms) | Fichiers exacts | Oracle |\n` +
  `| --- | ---: | ---: | ---: | ---: | ---: | --- |\n` +
  rows.map(row => `| ${row.package} | ${row.modules} | ${row.types} | ${row.median_ms ?? row.readme_cell} | ${row.median_ms === null ? '—' : `${row.min_ms}–${row.max_ms}`} | ${row.identical_files_per_generation ?? '—'} | ${row.oracle_host} |`).join('\n') +
  `\n\n**${generations} générations / ${identical} fichiers générés exacts** relus après mesure.\n\n` +
  `Sous-total de la colonne Rust \`gopurs-*\` : **${(displayedTotal / 1000).toFixed(2)} s** pour **${successfulRows.length}/${allRows.length} paquets**. C'est la somme des médianes validées affichées, incluant \`gopurs-aff\` déjà qualifié ; les ${report.verified.failed_packages.length} échecs sont exclus. Ce n'est pas une invocation globale chronométrée.\n\n` +
  `## Diagnostics conservés\n\n` +
  `\`gopurs-assert\` est une bibliothèque : son runner public appelle gopurs sans \`--main\`. La première tentative du harness exigeait à tort \`Test.Main\` ; la reprise utilise la commande réelle et conserve cinq mesures exactes.\n\n` +
  rows.filter(row => row.status !== 'passed').map(row => `- **${row.package}** : ${row.readme_cell}.${row.different_files.length ? ` Fichiers différents : ${row.different_files.map(path => '`' + path + '`').join(', ')}.` : ''} Diagnostic reproduit avec le même binaire et les mêmes entrées. Aucune médiane valide n'est publiée pour cette ligne.`).join('\n') +
  `\n\nLes différences de sources générées constituent un échec du contrôle d'identité, pas une preuve à elles seules de différence sémantique. La division par zéro de Prelude interrompt effectivement le backend Rust avant son horloge totale.\n\n` +
  `Les échantillons, empreintes et chemins des sorties brutes sont dans le rapport JSON associé.\n`);
console.log(JSON.stringify({ status: report.status, report: destination + '.md', patch: join(archive, 'README.patch'),
  ...report.verified, rust_total_s: displayedTotal / 1000 }, null, 2));
function basenameForDate(path) {
  const match = path.match(/gopurs-packages-rust-(\d{4})(\d{2})(\d{2})$/); assert(match);
  return `${match[1]}-${match[2]}-${match[3]}-gopurs-packages-rust`;
}
