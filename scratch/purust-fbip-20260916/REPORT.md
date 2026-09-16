# Purust FBIP — corrections et mesures du 16 septembre 2026

La génération Rust fonctionne de nouveau. La cible de réduction de 50 % du
temps d'exécution de RBTree n'est pas atteinte : le réemploi dynamique était
déjà effectif sur toutes les cellules testées du scénario unique.

## Corrections

- `purescript/src/Language/PureScript/CoreFn/Usage.hs` : indications d'usage
  source conservatrices, `escapes` conservé, qualification respectée, gardes
  avec repli comptées, aucune unicité fabriquée pour App/Constructor.
- `purescript-backend-optimizer-purust` : UsageMeta invalidé au passage dans
  l'IR sémantique, y compris pour les grandes expressions. Les informations
  de type restent présentes. Les usages de l'expression transformée sont
  recalculés par PBO.
- `purust/purust` : vue uniforme sans compteurs source périmés avant les
  passes Rust ; décisions de déplacement issues des ensembles `alive` ;
  option de diagnostic `--trace-phases`.

## Diagnostic du blocage

`before-sampled.log` : l'ancien bundle, dont les traces récursives ont été
remplacées par un compteur borné et des repères de phase dans une copie,
dépasse 200 000 appels de génération après 78,02 s. Il progresse entre modules
jusqu'à Data.Map.Internal. Cette expérience ne démontre pas une boucle infinie.

`after-old-tast.log` : le nouveau bundle termine sur les mêmes fichiers TAST
en 9,52 s. Ce diagnostic compare la terminaison ; il ne constitue pas un
benchmark contrôlé du temps de compilation.

## Validation

- Compilateur : cinq des six nouvelles régressions échouaient avant
  correction ; les 46 exemples CoreFn passent après correction.
- PBO : les 12 régressions UsageMeta échouaient avant correction ; elles
  passent, ainsi que 35 vérifications existantes (types, applications,
  négation, cache), soit 47 vérifications.
- Génération Rust : 74 tests passent dans la suite complète. Quatre tests
  FFI externes ne peuvent pas aboutir parce que leur conteneur préexistant
  `core-api-cli-1` est arrêté. Voir `codegen-tests.log` et
  `codegen-docker-tests.log` ; ce ne sont pas des échecs de compilation Rust.
- Pipeline frais TAST → PBO → Rust : `tests/tast/consumed-nodes.mjs` passe,
  avec 0 allocation pour 1 000 mutations uniques et pour 2 000 reconstructions
  via appels ; les versions partagées sont conservées correctement.

## Benchmark sans instrumentation

Baseline officielle : `altbak.pub/README.md`, 8 788,42 µs. Cible : 4 394,21 µs.
Compilation isolée O3 ; le compilateur Haskell local et Purust ont été
reconstruits. Les SHA-256 des exécutables utilisés sont conservés dans
`toolchain-fingerprints.json` ; le manifest complet est dans
`benchmark/manifest.json`.

```sh
cd /Users/0x1/Documents/htdocs/altbak.pub
./bin/rust/run --test Test.RBTree --build-only \
  --build-dir /Users/0x1/Documents/htdocs/altbak.pub/scratch/purust-fbip-20260916/benchmark
./bin/rust/run --test Test.RBTree --run-only \
  --build-dir /Users/0x1/Documents/htdocs/altbak.pub/scratch/purust-fbip-20260916/benchmark
```

Première mesure : 8 525,38 µs. Cinq séries suivantes :
8 538,08 ; 8 329,88 ; 8 608,62 ; 8 536,29 ; 8 348,96 µs.
Chaque série utilise 3 warm-ups puis le meilleur de 10. Médiane des cinq
résultats : 8 536,29 µs. Toutes produisent la profondeur attendue, 22.
Ce retour au niveau de la baseline historique ne démontre pas un gain
significatif par rapport au backend fonctionnel antérieur.

## Comptage séparé

Voir `allocation-probe/REPORT.md` et `allocation-probe/results.json` pour
la méthode, la source Rust identifiée par SHA-256 et les compteurs détaillés.

Sur 100 000 insertions : 100 001 allocations ; 2 483 932 contrôles get_mut,
tous réussis ; 99 978 rotations sur place. Les invariants rouge/noir,
les clés et 200 versions persistantes sont vérifiés séparément.

Les coûts encore présents incluent 2 183 976 clones temporaires d'enfants
pendant la récursion et 200 000 clones dans depth. Ces clones ne déclenchent
pas d'allocations supplémentaires. Leur suppression doit faire l'objet d'un
A/B temporel : diminuer un compteur ne suffit pas à prouver une accélération.
