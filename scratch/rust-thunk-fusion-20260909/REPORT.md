# Fusion des thunks Int immédiatement forcés

La fusion est intégrée dans Purust : **LazyEvaluation 17,280 ms → 1 µs**, **total 38,529 → 21,395 ms (−44,5 %)** sur cinq paires alternées de runners complets. Le noyau LazyEvaluation passe de **1 000 000 allocations à zéro**.

## Transformation et périmètre

[ThunkFusion.purs](../../../purust/purust/src/Purust/ThunkFusion.purs) reprend la reconnaissance de GoPurs, après les optimisations PBO et avant le renommage local et la génération Rust. Les types du TAST identifient un producteur récursif de `Unit -> Int` et son forçage immédiat. Le nouveau worker privé remplace l'accumulateur fonctionnel par un entier ; le générateur existant transforme sa récursion terminale en boucle.

Dans LazyEvaluation, l'appel qui construit puis force une chaîne de profondeur 1 000 avec une graine pure 0 devient un appel au worker entier. Le [diff généré](LazyEvaluation.diff) ajoute ce worker et remplace cet appel. Le producteur initial, son interface et les autres usages restent présents.

Cette première intégration accepte les arguments entiers fermés. Une preuve de 4 096 étapes au maximum vérifie que le worker termine et que chaque opération reste dans les bornes Int PureScript. Cette preuve empêche aussi le déplacement d'un débordement observable lorsque les vérifications arithmétiques Rust sont activées. Le résultat calculé pendant la preuve n'est pas substitué au programme : Purust émet le worker et laisse Rust optimiser la boucle.

Les paramètres inconnus, opérations hors bornes, dépassements du budget de preuve, graines opaques, appels intermédiaires inconnus et portées `LetRec` conservent le chemin initial. Chaque thunk doit demander son prédécesseur exactement une fois, sans branche conditionnelle dans cette demande. Les producteurs dont le résultat est conservé comme fonction restent disponibles. Le budget porte sur le travail du compilateur et n'impose aucune limite d'exécution au programme. Aucune règle ne dépend du nom du benchmark ou de sa profondeur particulière.

La cellule **B15 est donc jaune** dans le todo du worktree. Étendre la fusion aux arguments inconnus demanderait une preuve plus générale de terminaison et de sûreté arithmétique.

## Mesures du runner complet

Chaque processus conserve le meilleur de dix essais après échauffement. Les cinq paires alternent l'ordre avant/après ; les valeurs ci-dessous sont les médianes par benchmark. Le total additionne ces médianes. O1 et mimalloc sont conservés. Les noms et résultats des 14 benchmarks sont vérifiés dans les dix processus, dans la référence initiale et dans le run propre.

| Benchmark | Avant | Après |
| --- | ---: | ---: |
| LazyEvaluation | 17,280 ms | 0,001 ms |
| RBTree | 18,540 ms | 18,597 ms |
| Church | 1,463 ms | 1,485 ms |
| Deep Record Updates | 0,918 ms | 0,988 ms |
| **Total** | **38,529 ms** | **21,395 ms** |

[Résultats complets et mesures brutes](results.json). Le binaire passe de **1 509 088 à 1 506 832 octets**.

La [comparaison des sources](source-comparison.json) régénère les mêmes entrées TAST avec le bundle sauvegardé avant modification et avec le nouveau bundle : **seul Test.LazyEvaluation change parmi 302 fichiers Rust de crates**. La régénération de référence reproduit exactement la source LazyEvaluation sauvegardée initialement. Aucun gain n'est revendiqué sur les autres benchmarks ; l'écart observé de 70 µs sur les records est conservé dans les résultats malgré l'identité de leur source Rust.

Le [README officiel](../../../altbak.pub/README.md#rust) indique LazyEvaluation compilé **19,791 ms**, natif optimisé environ **0 µs**, total compilé **42,74 ms**, total natif **36,13 ms**. Ce sont des repères historiques distincts des paires ci-dessus. Le natif calcule directement une somme ; son zéro est un arrondi. Aucun nouveau temps du runner natif complet n'est revendiqué.

Le prototype préalable isolé donnait **15,351 ms → moins de 1 µs**, avec zéro allocation pour le worker strict et le témoin natif, sur trois tours alternés O1/mimalloc. Les temps sous la microseconde atteignaient la résolution de l'horloge. Ce prototype précédait l'intégration et ne mesurait pas le total.

## Allocations et validation

Le [comptage séparé](allocations.json) utilise les noyaux réellement générés avant/après, avec un allocateur instrumenté distinct des exécutables chronométrés. Il couvre construction, forçage et destruction :

| Noyau LazyEvaluation | Allocations | Libérations | Octets demandés cumulés |
| --- | ---: | ---: | ---: |
| Avant | 1 000 000 | 1 000 000 | 32 000 000 |
| Après | 0 | 0 | 0 |

Ces nombres excluent l'affichage, la conversion en chaîne et le reste du runner. Les appels répétés avec une graine opaque partagée et plusieurs profondeurs sont également contrôlés ; ils conservent le producteur initial.

**19 tests de génération et 10 tests TAST passent**, ainsi que **`bin/rust/run -c` et les 14 résultats attendus**. Le [test de la passe](../../../purust/purust/tests/codegen/thunk-fusion.mjs) vérifie les exclusions, les bornes, le budget, les noms réservés après sanitation, les demandes multiples et la conservation du producteur. La [fixture TAST](../../../purust/purust/tests/tast/thunk-fusion.mjs) compile des crates fraîches avec les vérifications de débordement activées. Elle vérifie notamment les transformations non commutatives, l'utilisation des anciennes valeurs des paramètres, les bornes Int, les graines opaques rejouables, inutilisées ou forcées conditionnellement, les erreurs et la préservation d'un débordement sur le chemin opaque.

La compilation initiale du générateur signalait 64 avertissements dans du code existant, aucun dans le nouveau module. La régénération propre aboutit sans erreur. Les sources, bundles et preuves sont identifiés par [validation.json](validation.json), la référence initiale par [metadata.json](metadata.json), et la modification reproductible par [integration.patch](integration.patch).

## Reproduction

Depuis `altbak.pub-purust`, avec les dépendances du runner disponibles :

```sh
python3 scratch/rust-thunk-fusion-20260909/count.py
python3 scratch/rust-thunk-fusion-20260909/measure.py nouvelle-serie
```

Le premier script recompile les sources Rust sauvegardées. Le second réutilise les deux runners sauvegardés. Les exécutables, le bundle de référence, les logs et le répertoire de compilation sont ignorés par Git. Pour reconstruire les runners ailleurs, utiliser les révisions de `metadata.json`, puis `integration.patch`, et sauvegarder chaque exécutable avant les mesures.

Les changements de cette étape concernent Purust et les documents/preuves du worktree altbak.pub-purust. Des modifications documentaires concurrentes concernant Sharpurs ont été observées dans le checkout normal d'altbak.pub et laissées intactes. Les checkouts de PBO sont inchangés.
