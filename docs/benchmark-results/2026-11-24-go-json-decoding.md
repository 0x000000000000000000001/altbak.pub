# JSON Decoding intégré — 24 novembre 2026

Le décodeur généré direct est intégré au compilateur et au port Argonaut. **JSON Decoding passe de 2 412,583 à 1 403,375 µs (−41,83 %, 1,72× plus rapide)**. Le ratio reste **2,1338× C** : l’exception provisoire demeure ouverte.

| Phase | Avant (µs) | Après (µs) | Variation | Octets alloués avant → après |
| --- | ---: | ---: | ---: | ---: |
| Parse seul | 1968.542 | 2067.583 | +5.03 % | 2435696 → 2435696 |
| Decode seul | 642.083 | 669.792 | +4.32 % | 1563120 → 1563120 |
| Parse + decode | 2412.583 | 1403.375 | −41.83 % | 3170008 → 2253992 |

L’optimisation cible le chemin combiné et élimine son index intermédiaire. Les phases isolées conservent leur chemin et leurs allocations ; leurs médianes sont ici plus lentes de 4–5 %. Cette campagne ne démontre pas de gain sur ces phases. Le chemin combiné alloue **28,90 % de moins**.

## Références et dispersion

Le README publiait **2237.333 µs**, soit **3.40× C**. La nouvelle valeur est 37,27 % plus basse que cette baseline historique, distincte de l’avant remesuré. Les références JSON restent inchangées : **C 657.69 µs**, **JS 9283.94 µs**, soit **2.133794× C** et **0.151162× JS** après intégration.

Les minima des trois processus sont, avant : 2412.583 / 2797.750 / 2352.792 µs ; après : 1269.209 / 1403.375 / 1535.250 µs. Seul le premier processus après est sous 2× C ; la médiane et les deux autres restent au-dessus. Aucun nouveau réglage du protocole n’a été utilisé pour franchir le seuil.

## Goulot et garanties

L’ancien chemin construisait un index de tokens, puis le parcourait pour produire le résultat typé. Les schémas prouvés purs produisent désormais directement leurs records et ADT pendant la lecture. Les scalaires numériques simples évitent une conversion générale redondante ; les autres utilisent `strconv.ParseFloat`.

Un résultat n’est accepté qu’après validation complète de l’entrée et de la profondeur. Les champs inconnus ou répétés, schémas non admis et erreurs reviennent au décodeur historique ; les callbacks opaques ne sont pas exécutés pendant la tentative. Les chaînes sont possédées et les conteneurs indépendants. Un contre-exemple a révélé un partage de tableaux quand le même champ était décodé deux fois : ce cas est maintenant exclu de l’optimisation, avec test de mutation vérifiant l’indépendance. Les branches exclusives peuvent toujours partager un emplacement temporaire.

Le marqueur optionnel `streamDecoderABI1` conserve le layout et l’ABI texte existants. Aucun benchmark, oracle, corpus ou FFI de chronométrage n’a été modifié. Aucune arène ni pool ajouté.

## Validation

- **35 tests décodeur + 4 tests d’imports réussis**, avec régression d’alias et exécution Go/race du code généré.
- Tests du port Go/race : grammaire et précision numérique, chaînes/Unicode, propriété des résultats, appels successifs, profondeur, fin d’entrée, erreurs et callbacks de repli.
- Fixture PureScript comparée au JavaScript officiel en **Release/Dev**, avec hôtes JS et natif. Le dernier garde-fou sur les lectures répétées est couvert par les tests IR et par la validation native finale.
- Ancien compilateur/nouveau port et nouveau compilateur/ancien port : réussis.
- Bootstrap natif final réussi, **547 modules TAST**.
- **17 oracles vérifiés dans chacun des six processus** du benchmark Release.
- Les 37 autres benchmarks sont inchangés : régénération forcée, **333 TAST, 9 fichiers Go complets et 11 933 fonctions identiques**. Leurs mesures précédentes restent valides ; aucune nouvelle campagne de chronométrage AppAll n’était nécessaire.

## Protocole et reproduction

Apple M4 Pro, Go 1.27.0, darwin/arm64, Release sans PGO, **GOMAXPROCS=1, GOGC=100**. Deux échauffements et cinq échantillons par phase/processus ; médiane des minima de trois processus par version, ordre **AB/BA/AB**. Cinq documents chronométrés, résultats possédés ; E/S et empreinte de vérification hors chronométrage.

```sh
python3 bin/benchmark/json-diagnostic.py build --suite JsonDecoding --runtime go \
  --workspace /tmp/gopurs-json-integration-20261124-after
python3 bin/benchmark/json-comparison.py \
  --before /tmp/gopurs-json-integration-20261124-before \
  --after /tmp/gopurs-json-integration-20261124-after \
  --output /tmp/gopurs-json-integration-20261124-comparison
```

Les chemins désignent les artefacts figés ; utiliser de nouveaux dossiers pour une nouvelle campagne et construire l’avant avec l’ancienne version. Le comparateur vérifie les empreintes des binaires, TAST, sources du benchmark, corpus et harness, tout en autorisant les changements de compilateur et de port.

L’[archive complète](2026-11-24-go-json-decoding.json) contient sorties brutes, manifestes avant/après, protocole, sources du comparateur, empreintes, audit AppAll et journaux de validation. La [nouvelle campagne C14](2026-11-24-c-reference.md) concerne séparément la suite principale ; elle ne remesure pas la référence C JSON.
