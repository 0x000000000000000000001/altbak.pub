import assert from 'node:assert/strict';
import { copyFileSync, existsSync, readFileSync, writeFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

assert(process.argv[2], 'publish-defaults.mjs ARCHIVE');
const archive = resolve(process.argv[2]), root = fileURLToPath(new URL('../../../', import.meta.url));
const read = path => JSON.parse(readFileSync(join(archive, path), 'utf8'));
const q = read('qualification.json'), h = read('hosts/results.json'), t = read('default-runs/results.json');
const v = read('verification.json'), b = read('baseline.json');
for (const r of [q, h, t, v]) assert.equal(r.status, 'passed');
assert.equal(v.campaigns[0].results_sha256, hash(readFileSync(join(archive, 'default-runs/results.json'))));
assert.equal(t.protocol.rounds, 5); assert.equal(t.protocol.warmups, 1); assert.equal(t.selection.launcherDefaults, true);
assert.equal(v.generations, 30); assert.equal(v.identical_files, 8820);
assert.deepEqual(manifest(join(archive, 'before')), b.compiler);
assert.deepEqual(manifest(join(h.compiler, 'src')), q.sources.gopurs);
assert.deepEqual(manifest(resolve(h.compiler, '../../purescript-backend-optimizer-gopurs/src')), q.sources.optimizer);
const changed = q.sources.gopurs.filter(f => b.sources.find(old => old.path === f.path)?.sha256 !== f.sha256).map(f => f.path).sort();
assert.deepEqual(changed, ['Gopurs/Driver.purs', 'Gopurs/Driver/Config.purs'].sort());
for (const [name, executable] of Object.entries(q.executables)) {
  assert.equal(hash(readFileSync(executable.path)), executable.sha256);
  assert.equal(hash(readFileSync(join(archive, 'default-runs/compilers/rust-default/bin', name))), executable.sha256);
}
const workers = '[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true';
for (const r of t.runs) {
  const stderr = readFileSync(r.stderr, 'utf8');
  if (r.variant === 'rust-before-default') assert(!stderr.includes('[gopurs] pbo stats:'));
  else { assert(stderr.includes(workers), r.label); assert(stderr.includes('[gopurs] pbo stats:'), r.label); }
  if (r.variant !== 'rust-explicit-eight') {
    for (const key of ['GOPURS_JOBS', 'GOPURS_PREPARE_JOBS', 'GOPURS_PBO_JOBS', 'GOPURS_EMIT_JOBS', 'GOPURS_PIPELINE'])
      assert.equal(r.explicit_environment[key], undefined, r.label + ' ' + key);
  }
}
const clocks = record => Object.fromEntries([...readFileSync(record.stderr, 'utf8').matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)]
  .map(([, key, ms]) => [key, Number(ms)]));
const live = h.runs.filter(r => r.host);
const goFiles = path => manifest(path, file => file.endsWith('.go') || file.endsWith('/go.mod'));
const expected = readFileSync(join(h.aff, 'test/expected-main.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort();
for (const r of live) {
  assert(readFileSync(r.stderr, 'utf8').includes(workers));
  assert.deepEqual(goFiles(r.go_output), r.generated);
  assert.deepEqual(r.generated, live[0].generated);
  assert.deepEqual(readFileSync(join(r.go_output, 'test.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort(), expected);
  assert.equal(readFileSync(join(r.go_output, 'test.stderr'), 'utf8'), '');
}
const tests = readFileSync(join(archive, 'hosts/logs/compiler-tests.stdout'), 'utf8');
assert.match(tests, /ℹ pass 43\b/); assert.match(tests, /ℹ fail 0\b/); assert.match(tests, /ℹ skipped 0\b/);
assert.match(readFileSync(join(archive, 'regression-before/stdout'), 'utf8'), /ℹ fail 3\b/);
const rustFiles = path => manifest(path, file => /\.(rs|toml)$/.test(file) && !file.endsWith('/Purs_Gopurs_FfiSupport/build.rs'));
assert.deepEqual(rustFiles(join(q.rust_bootstrap, 'rust')), h.bootstrap.native_generation);
assert.deepEqual(rustFiles(join(archive, 'hosts/bootstrap-js')), h.bootstrap.native_generation);
const oldReport = join(root, 'docs/benchmark-results/2026-10-03-gopurs-rust-optimization.json');
assert.equal(hash(readFileSync(oldReport)), b.report_sha256);
const s = t.summary, old = s['rust-before-default'], current = s['rust-default'];
const deltas = current.samples_ms.map((n, i) => n - old.samples_ms[i]);
const result = { status: 'passed', verified_at: new Date().toISOString(), archive,
  scope: 'Go-target gopurs, public launcher defaults; frozen Aff and separate live Aff execution',
  protocol: t.protocol, host: t.host, summary: s,
  default_rust_reduction_percent: 100 * (1 - current.median_ms / old.median_ms),
  paired_deltas_ms: deltas, favorable_pairs: deltas.filter(n => n < 0).length,
  live: Object.fromEntries(live.map(r => [r.host, clocks(r)])),
  rebuild_rust_default: clocks(q.commands.find(r => r.label === 'rust-default-rebuild-aff')),
  executables: q.executables, launcher_sha256: hash(readFileSync(join(h.compiler, 'bin/gopurs'))),
  qualification: { node_tests: 43, old_launcher_failed_contracts: 3, aff_checks: expected.length, avar_stress_items: 1000,
    bootstrap: h.frontend, identical_rust_files: h.bootstrap.identical_files, verification: v },
  source_records: ['baseline.json', 'qualification.json', 'hosts/results.json', 'default-runs/results.json', 'verification.json']
    .map(path => ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) })) };
const base = '2026-10-03-gopurs-host-defaults', destination = join(root, 'docs/benchmark-results', base);
assert(!existsSync(destination + '.json'));
const fmt = n => new Intl.NumberFormat('fr-FR', { maximumFractionDigits: 1 }).format(n);
const names = { 'rust-before-default': 'Rust — ancien défaut', 'rust-default': 'Rust — défaut corrigé',
  'rust-explicit-eight': 'Rust — huit workers explicites', 'go-default': 'Go — défaut', 'js-default': 'JavaScript — défaut' };
const text = `# Gopurs : parallélisme par défaut des hôtes

Correction du **3 octobre 2026**, après le signalement de **2 251 ms pour Go /
6 030 ms pour Rust** avec les commandes publiques. Sur cette machine de 48 Gio,
\`GOPURS_RUST=1 ./bin/test\` sélectionne désormais automatiquement **préparation 8,
PBO 8, émission 8 et pipeline actif**. Son exécution complète a passé les tests
Aff avec un backend à **${fmt(result.live.rust['backend total'])} ms** ; le chemin
de reconstruction \`GOPURS_RUST=1 ./bin/test -c\` passe également.

## Cause et correction

La branche Rust faisait \`exec\` avant la configuration automatique du launcher.
Elle conservait les défauts internes **préparation 2 / PBO 1**, alors que Go
recevait **8 / 8**. La [campagne d'optimisation précédente](2026-10-03-gopurs-rust-optimization.md)
forçait les workers explicitement ; sa validation de la commande ordinaire
était fonctionnelle. Ses chiffres ne décrivaient donc pas l'ancien défaut Rust.

Le launcher configure maintenant le parallélisme avant la sélection Go/JS/Rust.
Le seuil reste **32 Gio** et les variables explicites gardent la priorité.
Cette décision est indépendante de \`GOGC\`/\`GOMEMLIMIT\`. La politique de GC
automatique appartient uniquement au compilateur Go ; Rust conserve notamment
le GC ordinaire de son parseur Go embarqué.

Chaque compilation indique les limites effectives :

\`[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true\`

Les machines sous le seuil conservent les défauts internes, et les appels directs
aux exécutables contournent la politique du launcher. Préparation est bornée à
1–8, PBO/émission à 1–64 ; les valeurs affichées correspondent aux limites utilisées.

## Confirmation sur le TAST figé

Une chauffe puis **cinq tours tournants**, processus sérialisés, sur les mêmes
**238 modules / 136 604 types** que la campagne précédente. Le harnais retire
les variables héritées et n'impose aucun worker aux variantes « défaut ».
Seule la variante de contrôle explicite fixe 8/8/8/8 et le pipeline.

| Variante | Médiane | Moyenne | Min–max |
|---|---:|---:|---:|
${Object.entries(names).map(([key, name]) => `| ${name} | ${fmt(s[key].median_ms)} ms | ${fmt(s[key].mean_ms)} ms | ${fmt(s[key].min_ms)}–${fmt(s[key].max_ms)} ms |`).join('\n')}

Le défaut Rust passe de **${fmt(old.median_ms)} à ${fmt(current.median_ms)} ms
(−${fmt(result.default_rust_reduction_percent)} %)**, avec **${result.favorable_pairs}/5 paires favorables**.
Cette comparaison mesure l'effet des réglages automatiques sur gopurs déjà optimisé.
Le chargement reste à huit pour Go/Rust et à un pour JS par défaut.

| Tour | Rust ancien défaut | Rust défaut corrigé | Rust explicite 8 | Go défaut | JS défaut |
|---|---:|---:|---:|---:|---:|
${Array.from({ length: 5 }, (_, i) => `| ${i + 1} | ${Object.keys(names).map(k => fmt(s[k].samples_ms[i])).join(' | ')} |`).join('\n')}

Temps \`backend total\` uniquement : frontend, construction du compilateur,
build/exécution applicative et démarrage/sortie exclus. Sorties et caches neufs
à chaque passage, cache de fichiers du système chaud. Aucun échantillon exclu.

## Validation

- **43 tests launcher/CLI/codegen réussis, zéro skip**, dont six contrats du
  launcher. Trois de ces contrats échouent sur sa version précédente ; les
  diagnostics sont conservés. Les limites séquentielles/parallèles affichées
  sont aussi vérifiées sur les vrais compilateurs.
- JS, Go et Rust reconstruits ; bootstrap de **${h.frontend.modules} modules /
  ${fmt(h.frontend.types)} types**, **${h.bootstrap.identical_files} fichiers Rust/Cargo
  identiques** entre Purust JS et natif. Smoke Go/FFI frais et parseur Go \`-race\` réussis.
- Commandes Aff publiques des trois hôtes : **${expected.length} contrôles + stress AVar
  de 1 000 éléments**, **294 fichiers Go identiques** entre hôtes et à la qualification
  précédente. Paramètres automatiques vérifiés dans leurs logs.
- Confirmation figée : **30 générations / 8 820 fichiers exacts**, chauffes
  comprises, relus indépendamment avec les temps bruts et les exécutables.
- Les mesures historiques à workers explicites et les résultats JSON → Typed
  AST conservent leurs données et leur périmètre.

## Preuves

Archive : \`var/benchmark/gopurs-host-defaults-20261003/\` ;
[données vérifiées](${base}.json). Harnais :
\`bin/benchmark/gopurs-rust/qualify-defaults.mjs\`, \`compare.mjs\` avec
\`launcherDefaults: true\`, \`verify.mjs\` et \`publish-defaults.mjs\`.

| Exécutable installé | SHA-256 |
|---|---|
${Object.entries(q.executables).map(([name, e]) => `| ${name} | \`${e.sha256}\` |`).join('\n')}
`;
copyFileSync(fileURLToPath(import.meta.url), join(archive, 'publish-defaults.mjs'));
writeJson(destination + '.json', result); writeFileSync(destination + '.md', text);
console.log(JSON.stringify({ live: result.live, medians: Object.fromEntries(Object.entries(s).map(([k, v]) => [k, v.median_ms])),
  improvement_percent: result.default_rust_reduction_percent, favorable_pairs: result.favorable_pairs, destination }, null, 2));
