// Publish corrected cells only after re-reading the complete qualification.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, writeFileSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const archive = resolve(process.argv[2]);
const site = resolve(dirname(fileURLToPath(import.meta.url)), '../../..');
const read = path => JSON.parse(readFileSync(path));
const evidence = path => ({ path, sha256: hash(readFileSync(path)) });
const build = read(join(archive, 'build-results.json'));
const verificationPath = join(archive, existsSync(join(archive, 'verification-final.json')) ? 'verification-final.json' : 'verification.json');
const verification = read(verificationPath);
const hosts = read(join(archive, 'hosts/results.json'));
const recheckPath = join(archive, existsSync(join(archive, 'recheck/final-results.json')) ? 'recheck/final-results.json' : 'recheck/results.json');
const recheck = read(recheckPath);
const performancePath = join(archive, 'aff-performance/results.json'), performance = read(performancePath);
for (const result of [build, verification, hosts, recheck, performance]) assert.equal(result.status, 'passed');
assert.deepEqual(manifest(join(archive, 'bin')), build.binaries);
assert.deepEqual(manifest(join(archive, 'recheck/compiler')), recheck.compiler);
assert.equal(recheck.results.length, 50);
const original = read(join(recheck.original, 'results.json'));
assert.equal(hash(readFileSync(join(recheck.original, 'results.json'))), recheck.original_results_sha256);
assert.deepEqual(manifest(join(recheck.original, 'compiler')), original.compiler);
const pgoPath = join(verification.bootstrap, 'pgo-profile.json'), pgo = read(pgoPath);
assert.equal(pgo.status, 'passed'); assert.equal(pgo.passes.length, 3);
assert.equal(pgo.binary.sha256, build.binaries.find(file => file.path === 'gopurs-rust').sha256);
assert.deepEqual(manifest(join(verification.bootstrap, 'pgo/training')), pgo.training.manifest);
assert.equal(hash(readFileSync(join(verification.bootstrap, pgo.profile.merged.path))), pgo.profile.merged.sha256);
for (const pass of pgo.passes) assert.equal(pass.sha256, pgo.training.oracle.sha256);
assert.equal(performance.runs.length, 22);
assert.deepEqual(performance.compiler, recheck.compiler);
for (const run of performance.runs) {
  assert.equal(run.exit_code, 0); assert.deepEqual(manifest(run.output), run.generated);
  assert.deepEqual(run.generated, recheck.results.find(result => result.name === 'gopurs-aff').oracle.generated);
  const clocks = [...readFileSync(run.stderr, 'utf8').matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
  assert.equal(clocks.length, 1); assert.equal(Number(clocks[0][1]), run.phases_ms['backend total']);
}
const median10 = xs => { const sorted = xs.toSorted((a, b) => a - b); assert.equal(sorted.length, 10); return (sorted[4] + sorted[5]) / 2; };
const perfSamples = host => performance.runs.filter(run => run.round > 0 && run.host === host).map(run => run.phases_ms['backend total']);
const go = perfSamples('go'), rust = perfSamples('rust');
assert.deepEqual(performance.summary, { go_ms: go, rust_ms: rust, go_median_ms: median10(go), rust_median_ms: median10(rust),
  favorable_pairs: rust.filter((ms, i) => ms < go[i]).length, ratio_of_medians: median10(rust) / median10(go) });
let generations = 0, identicalFiles = 0;
const rows = [];
for (const result of recheck.results) {
  assert.equal(result.status, 'passed');
  assert.deepEqual(manifest(join(archive, 'recheck', result.name, 'input')), result.inputs);
  for (const family of new Set((result.sibling_inputs ?? []).map(file => file.path.split('/')[0]))) {
    const prefix = family + '/';
    assert.deepEqual(manifest(join(archive, 'recheck', result.name, family)).map(file => ({ ...file, path: prefix + file.path })),
      result.sibling_inputs.filter(file => file.path.startsWith(prefix)));
  }
  assert.deepEqual(manifest(result.oracle.output, path => path.endsWith('.go') || basename(path) === 'go.mod'), result.oracle.generated);
  const samples = [];
  for (const run of result.runs) {
    assert.equal(run.exit_code, 0);
    assert.deepEqual(manifest(run.output), run.generated);
    assert.deepEqual(run.generated, result.oracle.generated);
    const clocks = [...readFileSync(run.stderr, 'utf8').matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
    assert.equal(clocks.length, 1); assert.equal(Number(clocks[0][1]), run.phases_ms['backend total']);
    if (run.measured) samples.push(Number(clocks[0][1]));
    generations++; identicalFiles += run.generated.length;
  }
  if (result.measured) {
    assert.equal(samples.length, 5); assert.equal(samples.toSorted((a, b) => a - b)[2], result.summary.median_ms);
    assert.equal(result.applications.length, 2); assert(result.applications.every(run => run.exit_code === 0));
    rows.push({ package: result.name, ...result.summary,
      runs: result.runs.map(({ label, host, measured, phases_ms, stdout, stderr, output, identical_files }) =>
        ({ label, host, measured, phases_ms, stdout, stderr, output, identical_files })),
      applications: result.applications });
  }
}
assert.equal(rows.length, 4);
const readmePath = join(site, 'README.md'), before = readFileSync(readmePath, 'utf8'), lines = before.split('\n');
const edits = [];
for (const row of rows) {
  const index = lines.findIndex(line => line.startsWith(`[${row.package}]`)); assert(index >= 0);
  const old = lines[index];
  const diagnostic = row.package === 'gopurs-prelude' ? 'failed (division by zero)' : 'output mismatch';
  assert(old.endsWith('| ' + diagnostic), `Cell changed independently: ${row.package}`);
  const updated = old.slice(0, old.lastIndexOf('|') + 1) + ` ~ ${row.median_ms} ms`;
  edits.push({ line: index + 1, before: old, after: updated }); lines[index] = updated;
}
const packages = lines.filter(line => /^\[gopurs-[a-z0-9-]+\]/.test(line)); assert.equal(packages.length, 50);
const total = packages.reduce((sum, line) => {
  const value = line.match(/\| ~ (\d+) ms$/); assert(value); return sum + Number(value[1]);
}, 0);
const totalIndex = lines.findIndex(line => line.startsWith('**Total gopurs-***')); assert(totalIndex >= 0);
const oldTotal = lines[totalIndex]; assert(oldTotal.endsWith('(46/50)'));
const newTotal = oldTotal.slice(0, oldTotal.lastIndexOf('|') + 1) + ` ~ ${(total / 1000).toFixed(2)} s`;
edits.push({ line: totalIndex + 1, before: oldTotal, after: newTotal }); lines[totalIndex] = newTotal;
assert.equal(edits.length, 5);
for (const edit of edits) assert.deepEqual(edit.before.split('|').slice(0, 3), edit.after.split('|').slice(0, 3));
writeFileSync(join(archive, 'README.patch'), ['*** Begin Patch', `*** Update File: ${readmePath}`,
  ...edits.flatMap(edit => ['@@', '-' + edit.before, '+' + edit.after]), '*** End Patch', ''].join('\n'));
const report = { schema: 'gopurs-packages-fixes-1', status: 'passed', verified_at: new Date().toISOString(), archive,
  causes: {
    prelude: 'Purust lowered OpIntNum OpDivide to truncating, partial Rust division. Checked Euclidean division now matches Prelude, including zero divisors.',
    output_mismatches: 'Filesystem-dependent module enumeration changed topological ranks and therefore predecessor visibility/inlining. PBO now visits root modules in module-name order.',
    negative_zero: 'The wider three-host check exposed a pre-existing signed-zero bug in both native hosts: bootstrapping the constant evaluator specialized generic negate to zero-minus-value. A primitive FFI negation now preserves IEEE signed zero. Prelude has an explicit one-line oracle correction independently checked against JS and application execution.',
  },
  protocol: recheck.protocol, host: recheck.host, compiler: recheck.compiler, rows,
  verified: { packages: 50, newly_measured_packages: 4, generations, identical_generated_files: identicalFiles, applications: 4,
    rust_bootstrap_identical_files: hosts.bootstrap.identical_files, pbo_suites: build.stages.filter(run => run.label.startsWith('pbo-')).length,
    division_cases_per_build: 612, modulo_cases_per_build: 12, rust_test_profiles: ['debug', 'optimized'], module_permutations: 720 },
  qualification: { build: evidence(join(archive, 'build-results.json')), verification: evidence(verificationPath),
    hosts: evidence(join(archive, 'hosts/results.json')), recheck: evidence(recheckPath), pgo: evidence(pgoPath) },
  pgo: { training_modules: pgo.training.modules.length, passes: pgo.passes, sha256: pgo.profile.merged.sha256 },
  oracle_corrections: recheck.oracle_corrections,
  aff_performance: { evidence: evidence(performancePath), protocol: performance.protocol, ...performance.summary },
  aggregate: { packages: 50, successful_packages: 50, sum_of_displayed_rust_medians_ms: total,
    note: 'Sum of displayed medians from their respective campaigns, including these four corrected rows; not a timed multi-package invocation.' },
  previous_report: evidence(join(site, 'docs/benchmark-results/2026-10-04-gopurs-packages-rust.json')),
  readme: { before_sha256: hash(before), after_sha256: hash(lines.join('\n')), edits } };
writeJson(join(archive, 'publication.json'), report);
const destination = join(site, 'docs/benchmark-results/2026-10-04-gopurs-packages-fixes');
assert(!existsSync(destination + '.json') && !existsSync(destination + '.md'));
writeJson(destination + '.json', report);
writeFileSync(destination + '.md', `# gopurs : correction des quatre cellules Rust non qualifiées\n\n` +
  `Les quatre diagnostics du [rapport initial](2026-10-04-gopurs-packages-rust.md) sont résolus. Tous les hôtes génèrent du **Go**.\n\n` +
  `## Causes et corrections\n\n` +
  `- **Prelude** : la primitive de division entière de Purust utilisait \`/\` en Rust, qui panique sur zéro et tronque les quotients négatifs. Elle utilise désormais une division euclidienne contrôlée, avec zéro pour un diviseur nul. Les deux opérandes sont évalués une seule fois.\n` +
  `- **Enums, Promise et Strings** : l'ordre des répertoires différait entre hôtes. Le tri topologique conservait cet ordre pour les racines indépendantes, modifiant les rangs visibles par PBO et ses décisions d'inlining. Le tri part désormais d'un index ordonné par nom de module.\n\n` +
  `- **Zéro signé** : la comparaison supplémentaire avec JavaScript a révélé que les deux hôtes natifs transformaient la négation de zéro en soustraction à zéro lors du bootstrap de l'évaluateur PBO. Une négation primitive FFI conserve désormais le signe IEEE. L'oracle Prelude est corrigé explicitement sur une seule ligne de \`Test_Main.go\`, validée contre JavaScript et par exécution ; l'ancien oracle reste archivé. Le nouveau projet frais \`CompilerHostNumbers\` vérifie les zéros constants et dynamiques sous les trois hôtes.\n\n` +
  `## Validation\n\n` +
  `- Régressions reproduites avant correction ; **612 divisions** comparées à Prelude JavaScript et **12 cas modulo**, en Rust debug et optimisé.\n` +
  `- Tri vérifié sur **720 permutations**, imports propres/Prim et cycles ; **${report.verified.pbo_suites} suites PBO** réussies.\n` +
  `- Purust reconstruit par bootstrap indépendant avec auto-compilation et projet frais. Gopurs reconstruit dans ses trois hôtes ; **${hosts.bootstrap.identical_files} fichiers Rust/Cargo identiques** entre générateurs Purust JS et natif.\n` +
  `- PGO réentraîné sur **${pgo.training.modules.length} modules** de compilation du compilateur, en trois passes exactes ; \`Test.Main\` exclu.\n` +
  `- Qualification des hôtes, projets frais, Aff/AVar, tests du compilateur, helpers et parser Go sous détecteur de courses réussis ; détails et logs référencés dans le JSON.\n` +
  `- **50/50 paquets** revérifiés : **${generations} générations / ${identicalFiles} fichiers Go exacts** aux oracles archivés, avec la correction explicite du zéro signé de Prelude ci-dessus. Les quatre cas corrigés passent aussi les hôtes JS/Go et le mode Rust séquentiel. Leurs quatre applications Go sont compilées et exécutées avec succès, hors chronomètre.\n\n` +
  `## Nouvelles mesures\n\n` +
  `Mêmes entrées TAST/FFI figées que la campagne initiale ; une chauffe Rust puis cinq mesures sérialisées, workers **8/8/8/8**, pipeline actif. Médiane de \`backend total\` : chargement, préparation, PBO, génération/écritures et drain. Frontend, construction des compilateurs/applications, exécution des applications et démarrage/arrêt des processus exclus.\n\n` +
  `| Paquet | Rust médian | Min–max | Fichiers exacts par génération |\n| --- | ---: | ---: | ---: |\n` +
  rows.map(row => `| ${row.package} | ${row.median_ms} ms | ${row.min_ms}–${row.max_ms} ms | ${row.identical_files_per_generation} |`).join('\n') +
  `\n\nTotal de la colonne Rust : **${(total / 1000).toFixed(2)} s, 50/50 paquets**. Il s'agit de la somme des médianes affichées provenant de leurs campagnes respectives. Les cellules JS/Go et les 46 médianes Rust préexistantes gardent leurs mesures d'origine ; aucun nouveau ratio inter-hôtes n'est déduit de ce tableau.\n\n` +
  `## Confirmation de performance sur Aff\n\n` +
  `Une campagne séparée de dix paires contemporaines, après une chauffe par hôte, donne **Go ${performance.summary.go_median_ms} ms / Rust ${performance.summary.rust_median_ms} ms**, soit **${((performance.summary.ratio_of_medians - 1) * 100).toFixed(1)} %**, avec **${performance.summary.favorable_pairs}/10 paires favorables à Rust**. Ordre Go/Rust alterné, mêmes workers et entrées figées ; les 22 générations conservent les 294 fichiers Go exacts. Cette confirmation ne remplace pas la campagne historique Aff du README.\n\n` +
  `Les échecs et rapports initiaux restent archivés. Entrées, binaires, sorties, profils PGO, journaux et contrôles de publication : \`var/benchmark/gopurs-packages-fixes-20261004/\`.\n`);
console.log(JSON.stringify({ status: 'passed', ...report.verified, rust_total_s: total / 1000, patch: join(archive, 'README.patch') }, null, 2));
