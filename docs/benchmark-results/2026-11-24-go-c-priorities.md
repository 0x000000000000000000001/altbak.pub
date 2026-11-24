# Trois priorités Go/C sous 2× — 24 novembre 2026

AST Evaluation, Fibonacci et Deep Record Updates passent sous **2× C dans chacun des trois processus après optimisation**. JSON Decoding reste une **exception provisoire autorisée par l'utilisateur** ; aucune optimisation JSON de cette campagne n'est intégrée.

## Avant/après natif

Temps en µs, médiane des minima de trois processus. Comparaison des deux exécutables figés en ordre **AB/BA/AB**, mêmes sources de benchmarks et mêmes paramètres. Les ratios C sont approximatifs : références historiques arrondies du README, non remesurées pour cette campagne finale.

| Test | Avant | Après | Variation | Accélération | Après/C |
| --- | ---: | ---: | ---: | ---: | ---: |
| AST Evaluation | 0.341745 | 0.053361 | -84.39 % | 6.40× | ≈ 0.5336× |
| Fibonacci | 0.207413 | 0.153941 | -25.78 % | 1.35× | ≈ 1.9243× |
| Deep Record Updates | 8.861593 | 0.046134 | -99.48 % | 192.08× | ≈ 0.0135× |

Références C : AST 0,10 µs ; Fibonacci 0,08 µs ; Deep Record Updates 3,42 µs. Les anciens temps Go publiés étaient respectivement 0,325774, 0,214848 et 8,678508 µs. Ils restent distincts de l'avant fraîchement remesuré ci-dessus.

Échantillons après :

- AST Evaluation : 0.055949, 0.053361, 0.051628 µs.
- Fibonacci : 0.153941, 0.156737, 0.150918 µs.
- Deep Record Updates : 0.046134, 0.046983, 0.045674 µs.

Les durées très courtes et la précision publiée du C limitent la précision des ratios ; la marge de Fibonacci reste modeste.

## Goulots et changements

- **AST Evaluation** : l'arbre fraîchement construit était immédiatement consommé. `ConstructorFold` fusionne le producteur d'ADT et son consommateur Int à partir des layouts et types du TAST. Les champs doivent tous être consommés exactement une fois, dans l'ordre ; appels opaques, arbres partagés et opérations partielles conservent leur chemin d'origine. Le prototype isolé du calcul passe de 43 allocations à zéro ; avec l'action complète, de 46 à 3.
- **Fibonacci** : `TailAddLoops` transforme l'appel récursif terminal à droite d'une addition Int en boucle avec accumulateur. Les opérandes et arguments sont évalués dans leur ordre initial, avec une garde pour les bases Int64 non canoniques et le débordement Int32. Les retours directs contribuent aussi au gain. Le calcul conserve zéro allocation.
- **Deep Record Updates** : les copies de records n'étaient pas le coût dominant ; Go les éliminait déjà. `RecordAffineLoops` remplace des récurrences pures prouvées par leurs formules arithmétiques, notamment les sommes de restes. Les bornes, champs inchangés, références empruntées et chemins non admis sont conservés. Le calcul reste sans allocation.

Les transformations sont générales et couvertes par des fixtures renommées et des contre-exemples. Aucun benchmark, résultat attendu, FFI de chronométrage ou référence C/JS n'a été modifié pour obtenir ces gains.

## Validation et périmètre

- **101 tests IR et de régression réussis**.
- Fixtures AST, records et Fibonacci validées par les compilateurs hébergé JS et natif, en **Release et Dev**, contre le JavaScript officiel et des résultats explicites ; vérification de l'émission effective des chemins optimisés.
- Bootstrap natif réussi : 546 modules TAST. Le build JS ne produit aucun avertissement ; le frontend TAST signale sept avertissements dans deux modules inchangés (`GoImports`, `PrivateSTRegions`), aucun dans les modules modifiés.
- **37 résultats et batches vérifiés dans les six processus** ; les 37 ratios Go/JS restent sous 2× dans chacun des trois processus après.
- Audit du Go : **333 TAST identiques**, sept fonctions modifiées et deux helpers ajoutés, aucune fonction supprimée. Les sept changements concernent uniquement les trois optimisations et les trois tests `record/set` qui bénéficient de la règle générale.
- Aucun ralentissement supérieur à 5 % dans la comparaison finale des 37 médianes.

Le signal isolé de ralentissement de `effect/ref-write-10000` dans le preview n'est pas confirmé par la campagne appariée finale :

| Test | Variation avant/après | Après/JS |
| --- | ---: | ---: |
| effect/ref-write-10000 | -0.23 % | 1.2954× |
| record/set-small-1000 | -91.87 % | 0.0171× |
| record/set-medium-1000 | -97.61 % | 0.0307× |
| record/set-large-1000 | -95.97 % | 0.0565× |
| State Monad | -4.29 % | 1.8904× |

## Protocole reproductible

Go 1.27.0, darwin/arm64, Apple M4 Pro, Release, sans PGO, **GOGC=800**, GOMAXPROCS par défaut. Trois échauffements globaux, trois locaux, calibration, minimum de dix batches par processus ; médiane des trois minima. Références JS exactes issues du catalogue historique ; C aux précisions publiées. Les neuf références C mesurables de la suite principale restent sous 2×.

```sh
cd /Users/0x1/Documents/htdocs/gopurs/gopurs
npm run build:native
GOPURS_TEST_MODES=native-parallel node --test --test-concurrency=1 \
  tools/constructor-fold.test.mjs tools/record-affine-loops-e2e.test.mjs \
  tools/tail-add-loop-e2e.test.mjs

cd /Users/0x1/Documents/htdocs/altbak.pub
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  python3 bin/native/driver.py go --wip-included --build-only \
  --build-dir /tmp/gopurs-c-ratios-20261124-after
env -u GOMAXPROCS -u PPROF python3 bin/benchmark/js-ratios.py \
  --before /tmp/gopurs-js-ratios-20261124-after \
  --after /tmp/gopurs-c-ratios-20261124-after \
  --output /tmp/gopurs-c-ratios-20261124-comparison
```

Les dossiers de sortie doivent être nouveaux ; les chemins ci-dessus désignent les artefacts figés de cette campagne. L'[archive complète](2026-11-24-go-c-priorities.json) conserve les six sorties brutes, les manifestes, les versions, les hashes, les oracles, les références, l'audit du code et les journaux de validation. Le README actualise seulement la colonne gopurs optimisé et ses agrégats ; les autres colonnes et les diagnostics exclus gardent leurs mesures publiées.

## JSON Decoding : exception provisoire

La mesure publiée reste **2 237,333 µs**, soit **3,40× C** face à 657,69 µs. Une baseline locale isolée donnait 2 123,667 µs ; elle ne remplace pas le chiffre publié.

L'analyse a identifié le coût de l'index intermédiaire et des allocations. Les essais portent sur un indexeur itératif, l'analyse unique des nombres, des résultats de décodeur plus compacts, puis un décodeur généré produisant directement les records et ADT. Aucun de ces prototypes n'est intégré au compilateur ou au runtime livré.

Le dernier prototype, sur le véritable diagnostic JSON (**GOMAXPROCS=1, GOGC=100**), donne 1 401,208 / 1 358,875 / 1 412,416 µs, soit une médiane de **1 401,208 µs ≈ 2,1305× C**. Il reste au-dessus du seuil de 1 315,38 µs. Ses 17 oracles, 707 comparaisons de fallback et 21 variantes de branches passent ; cela ne constitue pas une validation d'intégration. Les mesures exploratoires plus rapides sous GOGC=800/GOMAXPROCS par défaut ne sont pas comparables à ce diagnostic.

Preuves de reprise locales : `/tmp/gopurs-c-priorities-experiments-report.md`, `/tmp/gopurs-c-json-streaming-fast-harness.json`, `/tmp/gopurs-c-json-streaming-fast-tests.log`, `/tmp/gopurs-c-json-streaming-source-hashes.json`. Les hashes et les résultats essentiels sont aussi conservés dans l'archive de cette campagne. L'exception reste ouverte dans le `todo.md` ; les trois optimisations validées sont cochées.
