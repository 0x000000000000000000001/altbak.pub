// Publish only after independently re-reading the final qualification and runs.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, manifest, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg] = process.argv.slice(2);
assert(archiveArg, 'publish.mjs ARCHIVE');
const archive = resolve(archiveArg), root = fileURLToPath(new URL('../../../', import.meta.url));
const read = name => JSON.parse(readFileSync(join(archive, name), 'utf8'));
const primary = read('primary-runs/results.json'), common = read('common-runs/results.json');
const resources = read('resources-runs/results.json'), production = read('production/results.json');
const hosts = read('production/hosts/results.json'), verification = read('verification-final.json');
const pbo = read('pbo-tests/results.json'), native = read('native-tests/summary.json'), maps = read('native-maps/summary.json');
for (const result of [primary, common, resources, production, hosts, verification, pbo]) assert.equal(result.status, 'passed');
for (const result of [native, maps]) { assert.equal(result.status, 'ok'); assert.equal(result.skipped, 0); }
for (const checked of verification.campaigns) assert.equal(hash(readFileSync(join(checked.directory, 'results.json'))), checked.results_sha256);
for (const [name, campaign, rounds, variants] of [
  ['primary', primary, 15, ['rust-before', 'rust']],
  ['common', common, 10, ['js', 'go', 'rust-before', 'rust']],
  ['resources', resources, 3, ['rust-before', 'rust']],
]) {
  assert.equal(campaign.protocol.rounds, rounds);
  assert.equal(campaign.protocol.warmups, 1);
  assert.deepEqual(campaign.variants.map(v => v.name), variants);
  assert.equal(campaign.runs.length, (rounds + 1) * variants.length);
  assert.deepEqual(campaign.protocol.jobs, { load: 8, prepare: 8, pbo: 8, emit: 8, pipeline: true });
  assert.deepEqual(campaign.tast, common.tast);
  assert(verification.campaigns.some(c => c.directory === join(archive, name + '-runs')));
  assert.equal(hash(readFileSync(join(archive, name + '-runs/compilers/rust/bin/gopurs-rust'))),
    production.executables['gopurs-rust'].sha256);
  assert.equal(hash(readFileSync(join(archive, name + '-runs/compilers/rust-before/gopurs-rust'))),
    hash(readFileSync(join(archive, 'baseline/compiler/bin/gopurs-rust'))));
}
const baseline = read('baseline.json');
assert.deepEqual(manifest(join(archive, 'bootstrap')), baseline.bootstrap);
assert.equal(hosts.bootstrap.purust_native_sha256, baseline.bootstrap.find(f => f.path === 'purust-native').sha256);
assert.equal(hosts.bootstrap.purust_js_sha256, baseline.bootstrap.find(f => f.path === 'purust.js').sha256);
const goFiles = directory => manifest(directory, file => file.endsWith('.go') || file.endsWith('/go.mod'));
const allRustFiles = directory => manifest(directory, file => /\.(rs|toml)$/.test(file));
const rustFiles = directory => manifest(directory, file => /\.(rs|toml)$/.test(file) && !file.endsWith('/Purs_Gopurs_FfiSupport/build.rs'));
assert.deepEqual(allRustFiles(join(production.rust_bootstrap, 'rust')), allRustFiles(join(production.candidate, 'rust')));
assert.equal(allRustFiles(join(production.candidate, 'rust')).length, production.identical_candidate_rust_files);
assert.deepEqual(rustFiles(join(production.rust_bootstrap, 'rust')), hosts.bootstrap.native_generation);
assert.deepEqual(rustFiles(join(archive, 'production/hosts/bootstrap-js')), hosts.bootstrap.native_generation);
assert.deepEqual(manifest(join(hosts.compiler, 'src')), production.sources.gopurs);
assert.deepEqual(manifest(resolve(hosts.compiler, '../../purescript-backend-optimizer-gopurs/src')), production.sources.optimizer);
for (const [name, executable] of Object.entries(production.executables)) {
  assert.equal(hash(readFileSync(executable.path)), executable.sha256);
  assert.equal(executable.sha256, hosts.executables[name].sha256);
  for (const host of ['js', 'go', 'rust'])
    assert.equal(hash(readFileSync(join(archive, 'common-runs/compilers', host, 'bin', name))), executable.sha256);
}
const affRuns = hosts.runs.filter(record => record.host);
assert.deepEqual(affRuns.map(r => r.host).sort(), ['go', 'js', 'rust']);
const expected = readFileSync(join(hosts.aff, 'test/expected-main.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort();
for (const record of affRuns) {
  assert.deepEqual(goFiles(record.go_output), record.generated);
  assert.deepEqual(record.generated, affRuns[0].generated);
  assert.deepEqual(readFileSync(join(record.go_output, 'test.stdout'), 'utf8').trimEnd().split(/\r?\n/).sort(), expected);
  assert.equal(readFileSync(join(record.go_output, 'test.stderr'), 'utf8'), '');
}
const testLog = readFileSync(join(archive, 'production/hosts/logs/compiler-tests.stdout'), 'utf8');
assert.match(testLog, /ℹ pass 37\b/); assert.match(testLog, /ℹ fail 0\b/); assert.match(testLog, /ℹ skipped 0\b/);
const preparation = production.commands.find(r => r.label.startsWith('gopurs-preparation-tests'));
assert.match(readFileSync(preparation.stdout, 'utf8'), /ℹ pass 19\b/);
assert.match(readFileSync(preparation.stdout, 'utf8'), /ℹ fail 0\b/);
assert.match(readFileSync(preparation.stdout, 'utf8'), /ℹ skipped 0\b/);
assert.equal(production.commands.find(r => r.label.startsWith('aff-rust-rebuild-default')).exit_code, 0);
assert.equal(hosts.runs.find(r => r.label === 'parser-tests').exit_code, 0);
assert.equal(production.commands.find(r => r.label.startsWith('go-memo-contract')).exit_code, 0);
assert.equal(readFileSync(join(production.rust_bootstrap, 'smoke-go-run.log'), 'utf8').trim(), 'Done');
for (const contract of [...native.nativeSources, ...maps.nativeSources])
  assert.equal(hash(readFileSync(contract.source)), contract.sha256);

// Preserve the separately qualified JSON -> Typed AST results and their raw
// structural oracles; do not relabel them as whole-compiler timings.
const historicalPath = join(root, 'docs/benchmark-results/2026-10-03-purust-native-optimization.json');
const historical = JSON.parse(readFileSync(historicalPath, 'utf8'));
const historicalJson = Object.fromEntries(['fixture12', 'gopurs238'].map(name => {
  const item = historical.json[name];
  assert.equal(item.status, 'passed');
  for (const binary of Object.values(item.binaries)) assert.equal(hash(readFileSync(binary.path)), binary.sha256);
  for (const record of item.runs) {
    assert.equal(record.exit_code, 0);
    const stdout = join(dirname(dirname(item.binaries[record.variant].path)), record.stdout);
    assert.deepEqual(JSON.parse(readFileSync(stdout, 'utf8')), record.result);
    assert.deepEqual(record.result.fingerprints, item.runs[0].result.fingerprints);
    assert.deepEqual(record.result.json_fingerprints, item.runs[0].result.json_fingerprints);
  }
  return [name, { modules: item.modules, processes: item.runs.length,
    structural_fingerprints_sha256: hash(JSON.stringify(item.runs[0].result.fingerprints)),
    binaries: item.binaries }];
}));

const median = values => { const xs = values.toSorted((a, b) => a - b), n = xs.length;
  return n % 2 ? xs[n >> 1] : (xs[(n >> 1) - 1] + xs[n >> 1]) / 2; };
const paired = (campaign, before, after) => {
  const a = campaign.summary[before].samples_ms, b = campaign.summary[after].samples_ms;
  const deltas = b.map((n, i) => n - a[i]);
  return { deltas_ms: deltas, median_delta_ms: median(deltas), mean_delta_ms: deltas.reduce((a, b) => a + b, 0) / deltas.length,
    favorable: deltas.filter(n => n < 0).length, rounds: deltas.length };
};
const resourceSummary = Object.fromEntries(['rust-before', 'rust'].map(variant => {
  const records = resources.runs.filter(r => !r.warmup && r.variant === variant).map(r => {
    const text = readFileSync(r.stderr, 'utf8');
    const clock = text.match(/([\d.]+) real\s+([\d.]+) user\s+([\d.]+) sys/);
    const rss = text.match(/(\d+)\s+maximum resident set size/);
    assert(clock && rss, r.stderr);
    return { real_s: Number(clock[1]), user_s: Number(clock[2]), sys_s: Number(clock[3]),
      cpu_s: Number(clock[2]) + Number(clock[3]), rss_bytes: Number(rss[1]) };
  });
  return [variant, { records, medians: Object.fromEntries(['real_s', 'cpu_s', 'rss_bytes'].map(k => [k, median(records.map(r => r[k]))])) }];
}));
const attempts = Object.fromEntries(['rust-before', 'rust'].map(variant => {
  const records = primary.runs.filter(r => r.variant === variant).map(r => {
    const text = readFileSync(r.stderr, 'utf8'), attempts = text.match(/pbo module attempts: (\d+), codegen: (\d+)/);
    const deferred = text.match(/deferredAttempts=(\d+)/);
    assert(attempts && deferred);
    return { attempts: Number(attempts[1]), codegen: Number(attempts[2]), deferred: Number(deferred[1]) };
  });
  return [variant, records];
}));
assert(attempts.rust.every(r => r.attempts === 238 && r.codegen === 238 && r.deferred === 0));
const before = primary.summary['rust-before'], after = primary.summary.rust, c = common.summary;
const improvement = 100 * (1 - after.median_ms / before.median_ms);
const ratio = c.rust.median_ms / c.go.median_ms, oldRatio = c['rust-before'].median_ms / c.go.median_ms;
const gap = 100 * (1 - (c.rust.median_ms - c.go.median_ms) / (c['rust-before'].median_ms - c.go.median_ms));
const marginal = read('ast-index-confirmation-runs/results.json');
const result = { status: 'passed', verified_at: new Date().toISOString(), compiler: 'gopurs',
  target: 'Go for every host', archive, host: common.host, tast: common.tast, protocol: common.protocol,
  executables: production.executables, primary: { ...primary.summary, paired: paired(primary, 'rust-before', 'rust'), improvement_percent: improvement },
  common: c, rust_to_go_ratio: ratio, before_rust_to_go_ratio: oldRatio, absolute_gap_reduction_percent: gap,
  resources: resourceSummary, pbo_attempts: attempts,
  historical_json_preserved: { path: historicalPath, sha256: hash(readFileSync(historicalPath)), corpora: historicalJson },
  ast_index_confirmation: { summary: marginal.summary, paired: paired(marginal, 'tast', 'ast-index') },
  qualification: { bootstrap: hosts.frontend, bootstrap_identical_files: hosts.bootstrap.identical_files,
    identical_candidate_rust_files: production.identical_candidate_rust_files,
    aff_checks_per_host: expected.length, avar_stress_items: 1000, aff_identical_files: affRuns[0].identical_files,
    node_tests: 37, preparation_bootstrap_tests: 19, pbo_suites: pbo.runs.length, native_suites: native.results.length + maps.results.length,
    default_rust_rebuild_aff: 'passed', parser_race: 'passed', verification },
  source_records: ['production/results.json', 'production/hosts/results.json', 'primary-runs/results.json',
    'common-runs/results.json', 'resources-runs/results.json', 'pbo-tests/results.json',
    'native-tests/summary.json', 'native-maps/summary.json', 'verification-final.json',
    'ast-index-confirmation-runs/results.json', 'confirmation-protocol.json'].map(path =>
      ({ path: join(archive, path), sha256: hash(readFileSync(join(archive, path))) })) };
const basename = '2026-10-03-gopurs-rust-optimization';
const destination = join(root, 'docs/benchmark-results', basename + '.json');
assert(!existsSync(destination));
const fmt = n => new Intl.NumberFormat('fr-FR', { maximumFractionDigits: 1 }).format(n);
const pct = n => fmt(n) + ' %';
const table = (headers, rows) => '| ' + headers.join(' | ') + ' |\n|' + headers.map(() => '---').join('|') + '|\n' +
  rows.map(row => '| ' + row.join(' | ') + ' |').join('\n');
const commonTable = table(['Hôte de gopurs', 'Médiane', 'Moyenne', 'Min–max'],
  [['js', 'JavaScript / Node'], ['go', 'Go natif'], ['rust-before', 'Rust avant'], ['rust', 'Rust optimisé']].map(([key, name]) =>
    [name, fmt(c[key].median_ms) + ' ms', fmt(c[key].mean_ms) + ' ms', `${fmt(c[key].min_ms)}–${fmt(c[key].max_ms)} ms`]));
const samples = table(['Tour', 'JS', 'Go', 'Rust avant', 'Rust optimisé'],
  Array.from({ length: 10 }, (_, i) => [i + 1, ...['js', 'go', 'rust-before', 'rust'].map(k => fmt(c[k].samples_ms[i]))]));
const phases = table(['Phase', 'Rust avant', 'Rust optimisé'],
  ['load TAST + sort', 'prepare + monomorphize', 'runtime', 'optimize + emit', 'entry points', 'backend total'].map(key =>
    [key, fmt(before.phases_median_ms[key]) + ' ms', fmt(after.phases_median_ms[key]) + ' ms']));
const cpu = resourceSummary['rust-before'].medians, finalCpu = resourceSummary.rust.medians;
const text = `# Gopurs hébergé en Rust — optimisation qualifiée

Campagne du **3 octobre 2026**. Le véritable **gopurs hébergé en Rust, produisant du Go**,
passe de **${fmt(before.median_ms)} à ${fmt(after.median_ms)} ms (−${pct(improvement)})** sur les
238 modules Aff figés. Confirmation sur **15 paires**, dont **${result.primary.paired.favorable} favorables**.
Le compilateur optimisé est installé et la commande **\`GOPURS_RUST=1 ./bin/test\`** passe.

## Comparaison commune : trois hôtes, une cible Go

${commonTable}

Dix tours, une chauffe par variante, ordre tournant. Dans cette même campagne,
le rapport Rust/Go passe de **${oldRatio.toFixed(3)}× à ${ratio.toFixed(3)}×** ;
**${pct(gap)} de l'écart absolu est résorbé**. Rust prend encore
${pct(100 * (ratio - 1))} de plus que Go et ${pct(100 * (1 - c.rust.median_ms / c.js.median_ms))}
de moins que JavaScript. La parité stricte reste à atteindre.

${samples}

Tous ces temps sont en millisecondes. JS et Go sont reconstruits avec les mêmes
sources PureScript finales que Rust ; le quatrième exécutable est le témoin Rust
initial figé. L'[ancienne qualification des hôtes](2026-10-03-gopurs-rust-host.md)
reste la référence initiale, avec 5 782,5 ms pour Rust dans sa propre campagne.

## Confirmation avant/après

${phases}

Moyennes : **${fmt(before.mean_ms)} → ${fmt(after.mean_ms)} ms** ; plages
${fmt(before.min_ms)}–${fmt(before.max_ms)} / ${fmt(after.min_ms)}–${fmt(after.max_ms)} ms.
Écart apparié médian **${fmt(result.primary.paired.median_delta_ms)} ms**, moyen
**${fmt(result.primary.paired.mean_delta_ms)} ms**. Aucun échantillon n'est exclu.
Les médianes de phases ne s'additionnent pas. La spécialisation transitive est
une sous-phase de préparation, déjà comprise dans sa durée.

${table(['Paire', 'Rust avant', 'Rust optimisé', 'Après − avant'], before.samples_ms.map((n, i) =>
    [i + 1, fmt(n), fmt(after.samples_ms[i]), fmt(after.samples_ms[i] - n)]))}

## Changements retenus

1. **Ordonnanceur PBO** : préserver les modules déjà prêts, respecter les attentes
   implicites et limiter la spéculation aux références finalisées. Publication
   canonique et progression en cas de dépendances inversées restent préservées.
2. **Cache de substitutions Rust** : utiliser l'identité des arbres partagés
   \`ExprType\`/\`BackendSyntax\`, avec conservation de leurs propriétaires, plutôt
   que l'enveloppe \`Any\` temporaire reconstruite à chaque appel.
3. **Directives natives** : petit parseur ASCII pour les cas courants ; toute
   autre syntaxe et les erreurs reviennent au parseur PureScript. Go et JS
   conservent la délégation au parseur PS.
4. **TAST Rust** : décodage natif des tableaux, annotations et modules, puis
   validation des faits d'usage ; partage des types, Unicode et ordre des
   diagnostics restent couverts par les oracles différentiels.
5. **Index AST global de gopurs** : réutiliser l'insertion native des maps de
   chaînes existante, avec les mêmes clés, ordre et valeurs. Son petit effet
   marginal est confirmé sur quinze paires : médianes
   ${fmt(marginal.summary.tast.median_ms)} → ${fmt(marginal.summary['ast-index'].median_ms)} ms,
   moyennes ${fmt(marginal.summary.tast.mean_ms)} → ${fmt(marginal.summary['ast-index'].mean_ms)} ms,
   ${result.ast_index_confirmation.paired.favorable}/15 favorables,
   écart apparié médian ${fmt(result.ast_index_confirmation.paired.median_delta_ms)} ms.

Les 16 exécutions de la confirmation finale (chauffe comprise) font exactement
**238 tentatives PBO / 238 modules, zéro rejet**, contre
**${Math.min(...attempts['rust-before'].map(r => r.deferred))}–${Math.max(...attempts['rust-before'].map(r => r.deferred))} rejets** avant.
Les profils instrumentés sont conservés à part des chronométrages. Le profil
initial attribuait 2 019 échantillons inclusifs aux substitutions, 456 aux
directives et 312 à la validation d'usage ; le profil du candidat final donne
243 / 14 / 51. Ces familles se recouvrent et ne sont pas des pourcentages CPU.

## Qualification

- Bootstrap Rust de production : **${hosts.frontend.modules} modules / ${fmt(hosts.frontend.types)} types** ;
  **${hosts.bootstrap.identical_files} fichiers Rust/Cargo identiques** entre Purust JS et natif.
  Les ${production.identical_candidate_rust_files} fichiers, script de liaison compris,
  sont aussi identiques au candidat mesuré. Projet frais Go/FFI exécuté avant
  l'installation atomique ; chemin public \`GOPURS_RUST=1 ./bin/test -c\` réussi.
- Les trois hôtes passent **${expected.length} contrôles Aff et le stress AVar de 1 000 éléments**,
  avec TAST vivant identique et **${affRuns[0].identical_files} fichiers Go exacts**.
- **37 tests CLI/codegen**, **19 tests préparation/auxiliaires**, **15 suites PBO**
  et **sept suites natives différentielles** réussissent, sans skip. Le parseur
  Go et le contrat de cache Go passent leurs contrôles \`-race\`.
- Différentiel natif : 169 directives par défaut, 352 cas générés, 33 replis ;
  238 modules natifs, 51 frontières de décodage, 110 909 annotations,
  824 tables / 149 495 entrées ; conservation des propriétaires du cache,
  réentrance, persistance des maps et clés Unicode.
- Vérification indépendante : **${verification.generations} générations /
  ${fmt(verification.identical_files)} fichiers Go/manifests** relus et exacts dans
  les campagnes de sélection, confirmation et ressources, chauffes comprises.
  Sorties, temps bruts et empreintes des exécutables sont revérifiés.

Le TAST chronométré demeure l'original ; la suite Aff synchronisée actuelle est
la qualification applicative séparée. Les résultats historiques **JSON → Typed
AST Go/Rust et leurs empreintes structurelles sont conservés** : leurs dix-huit
sorties brutes et les six exécutables archivés sont relus et revérifiés.
Le [rapport JSON/TAST historique](2026-10-03-purust-native-optimization.md#json--typed-ast-gorust)
garde son périmètre ; les nouvelles mesures portent sur la compilation complète
de gopurs-aff.

## Ressources et protocole

Diagnostic séparé \`/usr/bin/time -l\`, trois paires après chauffe, processus entier :

${table(['Ressource', 'Avant', 'Après'], [
  ['Temps réel', fmt(cpu.real_s) + ' s', fmt(finalCpu.real_s) + ' s'],
  ['CPU utilisateur + système', fmt(cpu.cpu_s) + ' s', fmt(finalCpu.cpu_s) + ' s'],
  ['Pic RSS', fmt(cpu.rss_bytes / 1048576) + ' Mio', fmt(finalCpu.rss_bytes / 1048576) + ' Mio'],
])}

- Corpus : **238 modules / 136 604 types / 26 606 491 octets**, manifeste TAST
  \`6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149\`.
- Mesure principale : \`backend total\`, chargement/tri, préparation, PBO,
  génération/émission Go, points d'entrée et attente finale des workers inclus.
  Frontend, bootstrap, builds et exécution applicative, démarrage/sortie exclus.
- Jobs chargement/préparation/PBO/émission : **8/8/8/8**, \`GOPURS_PIPELINE=1\`
  pour tous les hôtes. Ces paramètres explicites diffèrent du défaut PBO Rust
  séquentiel, également testé lors de la qualification.
- Rust : **O3 sans LTO ni debug, Arc et mimalloc**. Go garde le launcher de
  production, \`GOGC=off\` / \`GOMEMLIMIT=10GiB\` sur cette machine ; le parseur Go
  embarqué dans Rust garde son GC ordinaire.
- Processus et sorties neufs, caches \`.purmeta\`/\`.cache\` supprimés par passage,
  cache de fichiers du système chaud. Builds, profils, tests lourds et mesures
  sérialisés ; commandes longues en arrière-plan avec logs conservés.
- Apple M4 Pro, 14 cœurs logiques, 48 Gio, Node ${common.host.node}.

## Archives et diagnostics

Archive : \`var/benchmark/gopurs-rust-optimization-20261003/\`.
[Données vérifiées et empreintes](${basename}.json) ;
[harnais et commandes](../../bin/benchmark/gopurs-rust/README.md).
Les exécutions échouées restent archivées : une coupure réseau a interrompu
deux audits délégués ; le premier nouveau test de filtrage des maps omettait le
dictionnaire \`Ord\` et a été corrigé avant reprise de la qualification.
Les nettoyages de caches Cargo conservent et ré-empreintent les exécutables,
sources générées, logs, mesures et diagnostics.

${table(['Exécutable installé', 'SHA-256'], Object.entries(production.executables).map(([name, info]) => [name, '\`' + info.sha256 + '\`']))}
`;
const readmePath = join(root, 'README.md'), readme = readFileSync(readmePath, 'utf8');
const lines = readme.split('\n'), index = lines.findIndex(line => line.startsWith('[gopurs-aff]'));
assert(index >= 0);
copyFileSync(readmePath, join(archive, 'README-before-publication.md'));
copyFileSync(fileURLToPath(import.meta.url), join(archive, 'publish.mjs'));
lines[index] = `[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff)  | ~ ${Math.round(c.js.median_ms)} ms | ~ ${Math.round(c.go.median_ms)} ms | [~ ${Math.round(c.rust.median_ms)} ms](docs/benchmark-results/${basename}.md) <br>(/Go = ${ratio.toFixed(2)}x) (WIP)`;
writeJson(destination, result);
writeFileSync(join(root, 'docs/benchmark-results', basename + '.md'), text);
writeFileSync(readmePath, lines.join('\n'));
console.log(JSON.stringify({ primary_reduction_percent: improvement, common: c, rust_to_go: ratio,
  generations_verified: verification.generations, files_verified: verification.identical_files, destination }, null, 2));
