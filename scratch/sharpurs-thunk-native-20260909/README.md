# Sharpurs : intégration des thunks natifs — 9 septembre 2026

Le gain est confirmé sur la **génération normale du compilateur** et l’application complète Release. La DLL normale est recompilée à partir des mêmes 355 entrées que la copie après mesurée.

| Médianes, cinq processus par version | Avant | Après | Gain |
| --- | ---: | ---: | ---: |
| Lazy | 37,66767 ms | **5,93163 ms** | **−84,25 %** |
| Total | 94,02 ms | **62,36 ms** | **−33,67 %** |

Protocole : dix processus frais, ordre **ABBAABBAAB**, véritable `App.main`, meilleur de dix exécutions par test comme dans le benchmark existant. **140 sorties conformes** aux sorties historiques des 14 tests.

Plages Lazy : **37,01108–38,20779 ms** avant / **5,82404–6,63700 ms** après. Plages totales : **93,61–95,39 ms** / **62,01–63,60 ms**. Pas de recouvrement. Ce sont des statistiques descriptives locales ; les médianes des tests ne s’additionnent pas nécessairement à la médiane totale. RBTree, inchangé en source, varie de 42,35596 à 41,13333 ms ; cette variation n’est pas créditée à une optimisation RBTree.

## Ce qui est intégré

Un seul fichier change sur 355 : [Test.LazyEvaluation.fs](after/Test.LazyEvaluation.fs). La génération ajoute le helper **privé** `Test_LazyEvaluation_buildThunks_thunk_native`, typé `int -> (unit -> int) -> (unit -> int)`, et utilise celui-ci dans le sous-appel fermé de `runManyTimes_tco`.

Les bindings publics `force`, `defer`, `buildThunks` et `buildThunks_tco` sont **identiques octet pour octet**. La boucle extérieure garde sa forme et la charge conserve un million de thunks de profondeur 1000, avec résultat **1000000**. Le [diff](Test.LazyEvaluation.fs.diff) et [l’inventaire](generation-inventory.json) donnent le périmètre exact.

Le helper conserve les captures, le retard et la réévaluation, avec des appels et résultats natifs. La règle utilise les informations typées TAST/PBO et les arités source ; les cas non reconnus sûrs gardent le chemin générique et ses frontières d’exception. Il ne s’agit pas d’un remplacement manuel du F# ni d’une reconnaissance du nom du benchmark.

Ces résultats concernent **l’intégration**. La sonde antérieure isolée, dans `/private/tmp/sharpurs-lazy-typed-audit-20260909`, mesurait sur deux processus par version Lazy 34,54 → 5,96 ms et total 88,47 → 61,48 ms. Les deux campagnes ne sont pas mélangées.

## Repère officiel

Le [README principal figé](README.baseline.md) donne **89,03 ms total**, **35,37675 ms Lazy**, **39,94038 ms RBTree**. La dernière colonne native donne **0,075 ms Lazy**, mais son programme effectue 1000 incréments et retourne 1000 ; la charge compilée effectue un million de thunks et retourne 1000000. Le natif RBTree utilise aussi un autre algorithme/résultat (`SortedSet` mutable et Count, contre arbre persistant et profondeur). Ces chiffres historiques situent les résultats et ne sont pas des contrôles simultanés.

## Validation et provenance

- [VerifyLazy.fsx](VerifyLazy.fsx) : **333 assertions** réussissent sur chacune des DLL **avant, après et normale**, couvrant arithmetic Int32, retard/réévaluation, captures, partials, ordre d’évaluation, exceptions et charge complète.
- [Fixture du compilateur](fixture/validation.log) : **237 assertions F# + 75 assertions JS/reconnaissance** réussissent. Les dix suites existantes et les **49 contrôles Test.Main** passent ; leurs logs `regression-*.log` sont archivés.
- Trois builds Release : **31,94 s avant, 32,01 s après, 31,03 s normal**, chacun sans avertissement ni erreur. Build backend final sans avertissement ; les premiers avertissements de reconstruction Main/CodeGen étaient préexistants.
- Sources, assemblies et configuration runtime restent stables pendant les dix mesures. L’[intégrité finale](final-integrity.json) confirme aussi le backend final et la copie de la DLL normale d’avant. Aucun test/build concurrent pendant la série.
- Aucune mesure d’allocation, RAM ou GC n’est utilisée comme critère de performance.

[Comparaison complète](comparison.json) · [Processus et valeurs](executions.json) · [Environnement](metadata.json) · [Sources et bundles](inputs.json) · [Empreintes des validations](validation-files.json) · [Build normal](build-normal.json)

Les scripts `snapshot_after.py`, `capture_build.py`, `capture_normal_build.py`, `run.py` et `analyze.py` conservent la reproduction et la traçabilité. Les mesures refusent d’écraser une série existante. Les builds utilisent `/Users/0x1/.dotnet/dotnet`, Release, avec `-p:NuGetAudit=false` pour éviter l’audit réseau NuGet.
