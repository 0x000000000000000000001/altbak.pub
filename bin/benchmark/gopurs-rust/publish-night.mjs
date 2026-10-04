// Publish the gopurs Rust night campaign only after independently re-reading
// every configured gate, then update the single [gopurs-aff] README line.
//
// Usage: node publish-night.mjs ARCHIVE GATES.json [--publish]
//
// GATES.json decides what is required. Each gate entry is
//   { name, kind, path, status, required, paired? }
// with `path` relative to ARCHIVE, `kind` selecting the report rendering
// (production, hosts, bootstrap, runs, verification, final, text, aff-bracket,
// selection, other) and `status` the value the gate file's `status` field must
// carry ("present" means existence only). Nothing about the campaign outcome is
// guessed here: gate files, counts, rejected and preliminary entries all come
// from the config and the archived evidence.
//
// Without --publish the report is written to ARCHIVE/publication-preview.json
// and .md; docs/ and README.md are never touched. With --publish every required
// gate must exist and match, the victory claim and explicit counts must be
// present and the raw backend-total samples must be re-read, then the JSON, the
// report and the single README line are updated.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { hash, writeJson } from '../gopurs-aff/common.mjs';

const [archiveArg, gatesArg, ...flags] = process.argv.slice(2);
assert(archiveArg && gatesArg && flags.every(flag => flag === '--publish'),
  'publish-night.mjs ARCHIVE GATES.json [--publish]');
const publish = flags.includes('--publish');
const archive = resolve(archiveArg), gatesPath = resolve(gatesArg);
const root = fileURLToPath(new URL('../../../', import.meta.url));
const parse = path => JSON.parse(readFileSync(path, 'utf8'));
const sha = path => hash(readFileSync(path));
const fmt = value => value === null || value === undefined ? 'en attente'
  : typeof value !== 'number' ? String(value) : Number.isInteger(value) ? String(value) : value.toFixed(2);
const ratio = value => value.toFixed(4);
const median = values => { const sorted = values.toSorted((a, b) => a - b), middle = sorted.length >> 1;
  return sorted.length % 2 ? sorted[middle] : (sorted[middle - 1] + sorted[middle]) / 2; };
const table = (headers, rows) => [
  `| ${headers.join(' | ')} |`,
  `| ${headers.map(() => '---').join(' | ')} |`,
  ...rows.map(row => `| ${row.join(' | ')} |`),
].join('\n');
const stderrPath = (resultsPath, record) => record.stderr.startsWith('/')
  ? record.stderr : resolve(dirname(resultsPath), record.stderr);

// ---------------------------------------------------------------------------
// Configured gates
// ---------------------------------------------------------------------------
const config = parse(gatesPath);
assert(config.campaign && config.date && config.output, 'config needs campaign/date/output');
assert(Array.isArray(config.gates) && config.gates.length > 0, 'config.gates');
const gates = config.gates.map(entry => {
  assert(entry.name && entry.path && entry.status, `gate ${entry.name} needs name/path/status`);
  const path = join(archive, entry.path);
  const exists = existsSync(path);
  const found = exists ? parse(path) : null;
  const foundStatus = found ? (found.status ?? 'present') : null;
  return { name: entry.name, kind: entry.kind ?? 'other', required: entry.required !== false,
    path, paired: entry.paired ?? null, sha256: exists ? sha(path) : null,
    expected_status: entry.status, found_status: foundStatus,
    matches: exists && (entry.status === 'present' || foundStatus === entry.status) };
});
const required = gates.filter(gate => gate.required);
const allRequiredPassed = required.every(gate => gate.matches);
const missingGates = required.filter(gate => !gate.matches).map(gate => gate.name);
const gateOf = name => gates.find(gate => gate.name === name);
const evidence = name => { const gate = gateOf(name); return gate && gate.matches ? parse(gate.path) : null; };

const names = { primary: 'primary', default: 'default', common: 'common',
  resources: 'resources', beforeAfter: 'before-after', ...(config.campaigns ?? {}) };
const finalLabel = gateOf('final') ? basename(gateOf('final').path) : 'final-campaign.json';
const runResult = name => { const gate = gateOf(name); return gate && gate.matches ? parse(gate.path) : null; };
const pairedPath = gate => gate.paired ? join(archive, gate.paired) : join(archive, `${gate.name}-paired.json`);
const pairedResult = name => {
  const gate = gateOf(name);
  if (!gate || !gate.matches) return null;
  const path = pairedPath(gate);
  if (!existsSync(path)) return null;
  const paired = parse(path);
  assert.equal(resolve(paired.source), resolve(gate.path), `paired ${name} campaign binding`);
  assert(existsSync(paired.source), `paired ${name} source`);
  assert.equal(sha(paired.source), paired.source_sha256, `paired ${name} source drift`);
  return paired;
};

// Raw backend totals are re-read only once every required gate passes, so a
// preview never walks the full sample set.
function verifyRawSamples() {
  const verification = evidence('verification');
  assert(verification, 'verification-final.json is required for raw re-reads');
  assert(Array.isArray(verification.campaigns) && verification.campaigns.length > 0);
  let samples = 0;
  for (const campaign of verification.campaigns ?? []) {
    const resultsPath = join(campaign.directory, 'results.json');
    assert.equal(sha(resultsPath), campaign.results_sha256, campaign.directory);
    const results = parse(resultsPath);
    assert.equal(results.status, 'passed', campaign.directory);
    for (const record of results.runs) {
      const raw = readFileSync(stderrPath(resultsPath, record), 'utf8');
      const clock = Number(raw.match(/^\[gopurs\] backend total: (\d+) ms$/m)?.[1]);
      assert.equal(record.exit_code, 0);
      assert.equal(record.identical_files, 294);
      assert.equal(clock, record.phases_ms['backend total'], record.label);
      samples += 1;
    }
  }
  return samples;
}

// ---------------------------------------------------------------------------
// Claim: 30 explicit pairs, bootstrap interval below one, favorable count,
// then default corroboration and the final campaign's own victory flag
// ---------------------------------------------------------------------------
const primary = pairedResult(names.primary);
const defaults = pairedResult(names.default);
for (const paired of [primary, defaults].filter(Boolean)) {
  assert.equal(paired.reference, 'go'); assert.equal(paired.candidate, 'rust');
}
const final = evidence('final');
const common = runResult(names.common);
const beforeAfter = pairedResult(names.beforeAfter);
const victoryRequired = config.claim?.victory_required !== false;
const minFavorable = config.claim?.min_favorable ?? 27;
const primaryPairs = config.claim?.primary_pairs ?? 30;
const primaryGate = Boolean(primary)
  && primary.pairs.length === primaryPairs
  && primary.nightly_30_pair_gate === true
  && primary.median_candidate_ms < primary.median_reference_ms
  && primary.bootstrap.ratio_interval[1] < 1
  && primary.favorable_pairs >= minFavorable;
const defaultCorroboration = Boolean(defaults)
  && defaults.pairs.length === 10
  && defaults.median_candidate_ms < defaults.median_reference_ms
  && defaults.bootstrap.ratio_interval[1] < 1;
const victory = Boolean(final?.victory) && primaryGate && defaultCorroboration;
const rawSamples = allRequiredPassed && gateOf('verification')?.matches ? verifyRawSamples() : null;

// ---------------------------------------------------------------------------
// Campaigns, phases and whole-process resources
// ---------------------------------------------------------------------------
const selectionByCampaign = Object.fromEntries((config.selections ?? []).map(item => [item.campaign, item]));
const candidates = readdirSync(archive, { withFileTypes: true })
  .filter(entry => entry.isDirectory() && entry.name.endsWith('-runs'))
  .map(entry => join(archive, entry.name, 'results.json'))
  .filter(existsSync).sort()
  .map(path => {
    const results = parse(path);
    if (results.status !== 'passed') return null;
    const campaign = basename(dirname(path)).replace(/-runs$/, '');
    const variants = (results.variants ?? []).map(variant => variant.name);
    return { campaign, directory: dirname(path), rounds: results.protocol?.rounds ?? null, variants,
      medians: Object.fromEntries(variants.map(name => [name, results.summary?.[name]?.median_ms ?? null])),
      source_sha256: sha(path),
      selection: Object.values(names).includes(campaign) && gateOf(campaign)?.matches ? 'confirmation qualifiée'
        : selectionByCampaign[campaign]?.selected === true ? 'sélectionné'
        : selectionByCampaign[campaign] ? 'mesuré' : 'mesuré, non qualifié' };
  }).filter(Boolean);

const commonVariants = ['js', 'go', 'rust'];
const commonMedians = common
  ? Object.fromEntries(commonVariants.map(variant => [variant, common.summary?.[variant]?.median_ms ?? null])) : null;
const commonRatio = commonMedians?.rust != null && commonMedians?.go ? commonMedians.rust / commonMedians.go : null;
if (common) {
  assert.equal(common.protocol.rounds, 10);
  assert.deepEqual(common.variants.map(item => item.name), commonVariants);
}
const phases = common ? Object.fromEntries(commonVariants.map(variant =>
  [variant, common.summary?.[variant]?.phases_median_ms ?? null])) : null;

const resourcesGate = gateOf(names.resources);
const resources = resourcesGate && resourcesGate.matches ? parse(resourcesGate.path) : null;
if (resources) assert.equal(resources.protocol.rounds, 3);
const resourceSummary = resources ? Object.fromEntries((resources.variants ?? []).map(variant => {
  const records = (resources.runs ?? []).filter(record => !record.warmup && record.variant === variant.name)
    .map(record => {
      const text = readFileSync(stderrPath(resourcesGate.path, record), 'utf8');
      const clock = text.match(/([\d.]+) real\s+([\d.]+) user\s+([\d.]+) sys/);
      const rss = text.match(/(\d+)\s+maximum resident set size/);
      assert(clock && rss, record.label);
      return { real_s: Number(clock[1]), cpu_s: Number(clock[2]) + Number(clock[3]), rss_bytes: Number(rss[1]) };
    });
  return [variant.name, { pairs: records.length, cpu_s: median(records.map(record => record.cpu_s)),
    rss_bytes: median(records.map(record => record.rss_bytes)) }];
})) : null;

// ---------------------------------------------------------------------------
// Production, hosts, independent Purust bootstrap and operator counts
// ---------------------------------------------------------------------------
const production = evidence('production');
const hosts = evidence('hosts');
const bootstrap = evidence('bootstrap');
const sourceSelection = evidence('source-selection');
const compilerTests = hosts?.runs?.find(record => (record.label ?? '').startsWith('compiler-tests'));
const compilerPass = compilerTests ? Number(readFileSync(compilerTests.stdout, 'utf8').match(/ℹ pass (\d+)/)?.[1]) : null;
const derivedCounts = {
  production_identical_rust_files: production?.identical_candidate_rust_files ?? null,
  bootstrap_identical_files: bootstrap?.identical_generated_files ?? hosts?.bootstrap?.identical_files ?? null,
  compiler_suite_pass: Number.isInteger(compilerPass) ? compilerPass : null,
};
const counts = { ...derivedCounts, ...(config.counts ?? {}) };
const confirmations = Object.fromEntries(Object.values(names).map(name => {
  const campaign = runResult(name);
  if (!campaign) return null;
  const directory = dirname(gateOf(name).path);
  if (production) for (const variant of campaign.variants) {
    if (!['go', 'rust', 'js'].includes(variant.name)) continue;
    const binary = variant.name === 'go' ? 'gopurs-native' : variant.name === 'rust' ? 'gopurs-rust' : 'gopurs.js';
    assert.equal(sha(join(directory, 'compilers', variant.name, 'bin', binary)), production.executables[binary].sha256,
      `${name}: measured ${variant.name} must be the qualified production host`);
  }
  return [name, { path: gateOf(name).path, sha256: gateOf(name).sha256, protocol: campaign.protocol,
    summary: campaign.summary, samples: campaign.runs.map(record => ({ label: record.label, variant: record.variant,
      round: record.round, warmup: record.warmup, phases_ms: record.phases_ms, stderr: record.stderr,
      generated_manifest: record.generated_manifest, identical_files: record.identical_files })) }];
}).filter(Boolean));
const hostRows = (hosts?.runs ?? []).filter(record => record.host).map(record =>
  [record.host, record.checks, record.identical_files, record.tast_modules, record.avar_stress_items]);

// ---------------------------------------------------------------------------
// Historical reports: fingerprints preserved, Purust -> Rust axis kept apart
// ---------------------------------------------------------------------------
const historicalPath = join(root, 'docs/benchmark-results/2026-10-03-purust-native-optimization.json');
let historical = null;
if (existsSync(historicalPath)) {
  const document = parse(historicalPath);
  const corpora = ['fixture12', 'gopurs238'].map(name => {
    const item = document.json[name];
    assert(item && item.status === 'passed', `historical ${name}`);
    for (const binary of Object.values(item.binaries)) assert.equal(sha(binary.path), binary.sha256, binary.path);
    for (const record of item.runs) {
      assert.equal(record.exit_code, 0, name);
      const stdout = join(dirname(dirname(item.binaries[record.variant].path)), record.stdout);
      assert.deepEqual(parse(stdout), record.result, `${name} raw result`);
      assert.deepEqual(record.result.fingerprints, item.runs[0].result.fingerprints, `${name} fingerprints`);
      assert.deepEqual(record.result.json_fingerprints, item.runs[0].result.json_fingerprints, `${name} json fingerprints`);
    }
    return { name, modules: item.modules, runs: item.runs.length,
      structural_fingerprints_sha256: hash(JSON.stringify(item.runs[0].result.fingerprints)) };
  });
  historical = { path: historicalPath, sha256: sha(historicalPath), corpora,
    purust_to_rust: { axis: 'Purust -> Rust', field: 'goRust.summary.rust-final.median_ms',
      value: document.goRust.summary['rust-final'].median_ms } };
}
const historicalReports = ['2026-10-03-purust-native-optimization.json', '2026-10-03-gopurs-rust-optimization.json',
  '2026-10-03-gopurs-host-defaults.json', '2026-10-03-gopurs-rust-allocation.json']
  .map(name => join(root, 'docs/benchmark-results', name)).filter(existsSync)
  .map(path => ({ path, sha256: sha(path) }));

// ---------------------------------------------------------------------------
// Result and report
// ---------------------------------------------------------------------------
const result = {
  schema: 'gopurs-rust-night-publication-1',
  date: config.date, campaign: config.campaign, archive, mode: publish ? 'publish' : 'preview',
  generated_at: new Date().toISOString(),
  gates, all_required_gates_passed: allRequiredPassed, missing_gates: missingGates,
  claim: { victory, victory_required: victoryRequired, primary_gate: primaryGate,
    default_corroboration: defaultCorroboration, min_favorable: minFavorable, primary_pairs: primaryPairs,
     primary: primary ?? null, default: defaults ?? null, final_victory: final?.victory ?? null,
     before_after: beforeAfter,
    raw_samples_verified: rawSamples },
  common: common ? { host: common.host, tast: common.tast, medians: commonMedians,
    ratio_rust_go: commonRatio, phases_median_ms: phases } : null,
  resources: resourceSummary,
  confirmations,
  candidates, rejected: config.rejected ?? [], preliminary: config.preliminary ?? [], counts,
  production: production ? { candidate: production.candidate, build_profile: production.build_profile,
    identical_candidate_rust_files: production.identical_candidate_rust_files,
    executables: production.executables, pgo: production.pgo ?? null } : null,
  source_selection: sourceSelection,
  hosts: hosts ? { bootstrap_identical_files: hosts.bootstrap?.identical_files, host_rows: hostRows } : null,
  purust_bootstrap: bootstrap ? { status: bootstrap.status,
    identical_generated_files: bootstrap.identical_generated_files, tast: bootstrap.tast } : null,
  historical, historical_reports: historicalReports,
  source_records: [...gates.filter(gate => gate.sha256).map(gate => ({ path: gate.path, sha256: gate.sha256 })),
    ...historicalReports],
};

const claimLines = victory
  ? [`La revendication de victoire est établie : campagne finale \`${finalLabel}\`
     avec \`victory=true\`, **${primary.favorable_pairs}/${primaryPairs} paires favorables**, intervalle bootstrap
      [${primary.bootstrap.ratio_interval.map(ratio).join(', ')}] entièrement sous 1, corroboration aux défauts publics
      (${defaults.favorable_pairs}/${defaults.pairs.length} paires, intervalle [${defaults.bootstrap.ratio_interval.map(ratio).join(', ')}]).`]
  : [`Aucune revendication de victoire : ${final ? `final-campaign.victory=${String(final.victory)}` : 'final-campaign.json absent'},
     primaire 30 paires ${primaryGate ? 'établie' : 'non établie'} (${primary ? `${primary.favorable_pairs}/${primaryPairs} favorables, intervalle [${primary.bootstrap.ratio_interval.map(ratio).join(', ')}]` : 'absente'}),
     corroboration défauts ${defaultCorroboration ? 'établie' : 'non établie'}.`];
const rejectedText = (config.rejected ?? []).length
  ? (config.rejected ?? []).map(item => `- **${item.label}**${item.median_ms ? ` (${fmt(item.median_ms)} ms)` : ''} — ${item.reason}`).join('\n')
  : '_Aucun rejet configuré._';
const preliminaryText = (config.preliminary ?? []).length
  ? (config.preliminary ?? []).map(item => `- **${item.label}**${item.median_ms ? ` (${fmt(item.median_ms)} ms)` : ''}${item.tests ? `, ${item.tests} tests` : ''} — ${item.reason}`).join('\n')
  : '_Aucun essai préliminaire configuré._';
const text = `# gopurs Rust nuit — ${config.date}

Statut : **${publish ? 'publié' : 'prévisualisation, aucune publication'}** ;
${allRequiredPassed ? 'tous les gates requis passent.' : `gates requis en attente : ${missingGates.join(', ')}.`}

## Résultat

Les trois hôtes exécutent **gopurs et génèrent du Go**. ${common ? `Le corpus figé
\`gopurs-aff\` contient **${common.tast.modules} modules / ${common.tast.types} types**
(${common.tast.bytes} octets), manifeste TAST \`${common.tast.sha256}\`.
Machine : ${common.host.cpu}, ${common.host.logical_cpus} processeurs logiques,
${common.host.memory_bytes / 1073741824} Gio de RAM.` : ''}

${claimLines.join('\n')}

${table(['Contrôle', 'Valeur'], [
  ['Campagne finale', finalLabel],
  ['Victoire', String(victory)],
  ['Paires primaires / favorables', primary ? `${primary.pairs.length} / ${primary.favorable_pairs}` : 'en attente'],
  ['Intervalle bootstrap primaire (ratio Rust/Go)', primary ? primary.bootstrap.ratio_interval.map(ratio).join(' … ') : 'en attente'],
  ['Médianes primaires Go → Rust', primary ? `${fmt(primary.median_reference_ms)} → ${fmt(primary.median_candidate_ms)} ms` : 'en attente'],
  ['Écart médian Rust/Go', primary ? `${fmt(primary.median_change_percent)} %` : 'en attente'],
  ['Corroboration défauts', defaults ? `${defaults.favorable_pairs}/${defaults.pairs.length} favorables, ${fmt(defaults.median_reference_ms)} → ${fmt(defaults.median_candidate_ms)} ms` : 'en attente'],
  ['Rust début de nuit → production (campagne indépendante)', beforeAfter ? `${fmt(beforeAfter.median_reference_ms)} → ${fmt(beforeAfter.median_candidate_ms)} ms (${fmt(beforeAfter.median_change_percent)} %)` : 'en attente'],
  ['Échantillons bruts revérifiés', rawSamples === null ? 'non relus (gates incomplets)' : String(rawSamples)],
])}

## Candidats mesurés (dossiers *-runs passés)

${table(['Campagne', 'Tours', 'Variantes (médianes ms)', 'Statut'], candidates.length
  ? candidates.map(row => [row.campaign, row.rounds,
    Object.entries(row.medians).map(([name, ms]) => `${name}: ${fmt(ms)}`).join(' ; '), row.selection])
  : [['—', '—', '—', 'aucun']])}

Les campagnes mesurées ne sont pas qualifiées par cette extraction : seule la campagne
finale et les gates configurés décident.

## Rejetés

${rejectedText}

## Préliminaires / non qualifiés

${preliminaryText}

## Phases (p50, non additives)

${common ? table(['Variante', 'backend total', 'load + tri', 'transitive', 'prepare + mono.', 'PBO optimize + emit', 'entry points'],
  commonVariants.map(variant => [variant, fmt(commonMedians[variant]),
    fmt(phases[variant]?.['load TAST + sort']), fmt(phases[variant]?.['transitive specializations']),
    fmt(phases[variant]?.['prepare + monomorphize']), fmt(phases[variant]?.['optimize + emit']),
     fmt(phases[variant]?.['entry points'])])) : '_campagne common absente_'}

${common ? table(['Variante', 'Producteur PBO (avec attente émission)', 'Génération + écritures (lots cumulés)', 'Drain final émission'],
  commonVariants.map(variant => [variant, fmt(phases[variant]?.['PBO producer']),
    fmt(phases[variant]?.['generation + writes (cumulative batches)']),
    fmt(phases[variant]?.['emitter drain'])])) : ''}

\`backend total\` couvre chargement/tri, préparation, PBO, génération/émission Go, points
d'entrée et drain des workers ; \`prepare + monomorphize\` inclut \`transitive
specializations\`. Les phases **ne s'additionnent pas** et les familles de profil se
recouvrent.
Le producteur PBO inclut sa contre-pression vers l'émetteur ; les temps cumulés
des lots de génération se superposent à son exécution en mode pipeline. Le drain
mesure l'attente finale de l'émetteur et reste inclus dans \`optimize + emit\`.

## Ressources (processus entier, CPU + RSS)

${resourceSummary ? table(['Variante', 'Paires', 'CPU médian (s)', 'RSS médian (Mio)'],
  Object.entries(resourceSummary).map(([variant, summary]) => [variant, summary.pairs,
    fmt(summary.cpu_s), fmt(summary.rss_bytes / 1048576)]))
  + '\n\nParsing brut de \`/usr/bin/time -l\` (Mac), processus entier avec descendants.' : '_campagne resources absente_'}

Sur cette machine, le lanceur applique à l'hôte Go sa politique mémoire
\`GOGC=off\` / \`GOMEMLIMIT=10GiB\` (seuil d'au moins 32 Gio de RAM).
Les chiffres de ressources reflètent cette configuration publique ; le réglage
automatique du GC Go n'est pas appliqué à l'hôte Rust.

## Qualification

${table(['Compte', 'Valeur'], Object.entries(counts).map(([key, value]) => [key, fmt(value)]))}

${hosts ? table(['Hôte', 'Contrôles Aff', 'Fichiers Go exacts', 'Modules TAST', 'Stress AVar'],
  hostRows.map(row => row.map(fmt))) + '\n' : ''}
${production ? `Candidat production \`${basename(production.candidate)}\`, profil \`${JSON.stringify(production.build_profile)}\`.\n` : ''}
${bootstrap ? `Bootstrap Purust indépendant : ${fmt(bootstrap.identical_generated_files)} fichiers identiques, TAST ${fmt(bootstrap.tast?.modules)} modules.\n` : ''}

${production?.pgo ? `### PGO de production

Le build public reconstruit son profil sur **${production.pgo.training_modules} modules
d'auto-compilation du compilateur**, puis installe le binaire qualifié.
Les **${production.pgo.shared_library_modules.length} modules de bibliothèque communs** avec
le corpus tenu à l'écart sont nommés dans le rapport JSON ; \`Test.Main\` est exclu.
Les trois passes d'entraînement conservent le manifeste des entrées et reproduisent
exactement les sorties Go du compilateur sans profil.

- Sources Rust générées : \`${production.pgo.source_sha256}\`.
- Entrées d'entraînement : \`${production.pgo.training_sha256}\`.
- Profil fusionné : \`${production.pgo.profile_sha256}\`.
- Binaire installé : \`${production.pgo.binary_sha256}\`.
- Métadonnées complètes : \`${production.pgo.path}\`.

La sélection isolée et ce build public ont des profils distincts : les mesures
finales ci-dessus portent sur le **binaire installé**. Le profil se régénère par
\`npm run build:rust\` ; \`GOPURS_RUST_PGO=0\` permet de reconstruire le témoin sans PGO.
` : ''}

## Historique

${historical ? `Rapport \`${basename(historical.path)}\` conservé (sha256 \`${historical.sha256}\`) :
${historical.corpora.map(corpus => `${corpus.name} ${corpus.modules} modules / ${corpus.runs} exécutions`).join(', ')}.
L'axe **Purust → Rust** (\`${historical.purust_to_rust.field}\` = ${fmt(historical.purust_to_rust.value)} ms)
reste distinct de l'axe **gopursRust → Go** mesuré ici : ce chiffre n'est ni redéfini ni
réutilisé.` : '_rapport historique absent_'}
${historicalReports.length ? `\n${table(['Rapport JSON conservé', 'SHA-256'], historicalReports.map(report => [basename(report.path), '`' + report.sha256 + '`']))}` : ''}

## Méthode, gates et reproduction

- Commande : \`node altbak.pub/bin/benchmark/gopurs-rust/publish-night.mjs ${basename(archive)} ${basename(gatesPath)} --publish\`
- Métrique : \`backend total\` (chargement/tri, préparation, PBO, génération/émission Go, points d'entrée, drain) ;
  frontend, builds, exécution et démarrage/sortie du processus exclus.
- Corpus et compilateurs gelés, un échauffement par variante, processus sérialisés,
  ordre de passage tournant, sorties et caches PBO recréés à chaque génération.
  Comparaisons explicites : chargement/préparation/PBO/émission 8/8/8/8 et pipeline actif.
  Les dix paires « default » passent par le lanceur public sans ces surcharges.
- Intervalle à 95 % : rééchantillonnage apparié des logarithmes des ratios Rust/Go,
  100 000 réplications déterministes ; chaque paire reste groupée.
- Échantillons bruts relus : ${rawSamples === null ? 'non relus tant que tous les gates requis ne passent pas' : `${rawSamples}`}.
- Archive : \`${archive}\`.

${table(['Gate', 'Type', 'Attendu', 'Trouvé', 'Requis', 'SHA-256'],
  gates.map(gate => [gate.name, gate.kind, gate.expected_status, gate.found_status ?? 'absent',
    gate.required ? 'oui' : 'non', gate.sha256 ? '`' + gate.sha256 + '`' : '—']))}
`;

if (!publish) {
  writeJson(join(archive, 'publication-preview.json'), result);
  writeFileSync(join(archive, 'publication-preview.md'), text);
  console.log(JSON.stringify({ mode: 'preview', missing_gates: missingGates, victory,
    preview: join(archive, 'publication-preview.md') }, null, 2));
} else {
  assert(allRequiredPassed, `required gates not passed: ${missingGates.join(', ')}`);
  assert(!victoryRequired || victory, 'victory claim is not established');
  assert(config.counts && Object.keys(config.counts).length > 0, 'explicit counts required');
  assert(historical && historical.corpora.length === 2 && historicalReports.length === 4,
    'all historical qualifications must remain available and verified');
  assert(rawSamples && rawSamples > 0, 'raw backend total samples not verified');
  assert(!existsSync(join(root, 'docs/benchmark-results', config.output + '.json'))
    && !existsSync(join(root, 'docs/benchmark-results', config.output + '.md')), 'publication already exists');
  assert(commonMedians?.js != null && commonMedians?.go != null && commonMedians?.rust != null,
    'common medians are required for the README line');
  const destination = join(root, 'docs/benchmark-results', config.output);
  writeJson(destination + '.json', result);
  writeFileSync(destination + '.md', text);
  const readmePath = join(root, 'README.md'), readme = readFileSync(readmePath, 'utf8');
  const lines = readme.split('\n');
  const indices = lines.map((line, index) => line.startsWith('[gopurs-aff]') ? index : -1)
    .filter(index => index >= 0);
  assert.equal(indices.length, 1, 'README gopurs-aff line');
  copyFileSync(readmePath, join(archive, 'README-before-publication.md'));
  copyFileSync(fileURLToPath(import.meta.url), join(archive, 'publish-night.mjs'));
  copyFileSync(gatesPath, join(archive, 'gates.json'));
  // Only this line changes; every other byte of README.md is preserved.
  lines[indices[0]] = `[gopurs-aff](https://github.com/0x000000000000000000001/gopurs-aff)  | ~ ${Math.round(commonMedians.js)} ms | ~ ${Math.round(commonMedians.go)} ms | [~ ${Math.round(commonMedians.rust)} ms](docs/benchmark-results/${config.output}.md) <br>(/Go = ${commonRatio.toFixed(2)}x) (WIP)`;
  writeFileSync(readmePath, lines.join('\n'));
  console.log(JSON.stringify({ mode: 'publish', missing_gates: [], victory,
    common: commonMedians, ratio_rust_go: commonRatio, destination: destination + '.md',
    raw_samples_verified: rawSamples }, null, 2));
}
