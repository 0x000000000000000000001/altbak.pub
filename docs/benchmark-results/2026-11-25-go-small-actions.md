# Polymorphism et RowToList : optimisation Go — 25 novembre 2026

Les optimisations de gopurs réduisent le temps de **Polymorphism de 55,66 %** et celui de **RowToList de 51,92 %** dans une campagne avant/après. Les benchmarks, leur harnais et les références C/JS sont inchangés. La moyenne arithmétique Go/C passe de **5,045× à 2,611×** ; la cible de 1× reste à atteindre.

## Résultats des deux cas

Les temps sont les médianes de trois minima de processus, en nanosecondes par action complète. Les ratios emploient les mêmes médianes C fonctionnelles archivées : 0,679731 ns pour Polymorphism et 1,658201 ns pour RowToList.

| Cas | Go avant (ns) | Go après (ns) | Accélération | Go/C avant | Go/C après | Allocations avant → après |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Polymorphism | 31.030 | 13.760 | 2.255× | 45.650× | 20.243× | 4 → 1 |
| RowToList | 26.994 | 12.980 | 2.080× | 16.279× | 7.828× | 3 → 1 |

Les trois processus après donnent 13,735–13,923 ns pour Polymorphism et 12,825–13,442 ns pour RowToList. Ce sont des plages observées, pas des intervalles de confiance. Les allocations proviennent de sondes séparées sur les actions Go réellement générées, via `runtime.Apply`, sans substitution du code et avec trois répétitions concordantes.

## Les quatre indicateurs Go/C

Pour chaque cas, le ratio est la médiane Go divisée par la médiane C figée. Le ratio des totaux vaut ΣGo/ΣC ; μ est la moyenne arithmétique des 14 ratios ; Q₂ leur médiane ; G leur moyenne géométrique. Les 23 cas de bibliothèque servent au contrôle des régressions et ne sont pas inclus dans ces quatre calculs.

| Indicateur | Avant | Après | Variation |
| --- | ---: | ---: | ---: |
| Ratio des totaux | 0.139809× | 0.139882× | +0.053 % |
| Moyenne arithmétique μ | 5.045127× | 2.610828× | -48.251 % |
| Médiane Q₂ | 0.157630× | 0.157916× | +0.182 % |
| Moyenne géométrique G | 0.135576× | 0.104848× | -22.665 % |

Le ratio des totaux et la médiane restent pratiquement stables dans cette campagne. Polymorphism et RowToList sont des actions très courtes : leurs gros ratios pesaient fortement sur μ, mais très peu sur le temps total. Une allocation subsiste par action ; leurs ratios Go/C restent respectivement de 20,24× et 7,83×.

## Comparaison avec le bilan publié

Le [bilan Go/C/JS précédent](2026-11-25-go-c-js-functional.md), repris dans le README, mesurait l’entrée `App` à 14 cas. La présente campagne mesure deux binaires `AppAll` identiques par leurs sources (37 cas), avant et après modification du compilateur. La colonne « avant » est donc une nouvelle mesure, et ne remplace pas les chiffres historiques. Les dénominateurs C sont exactement ceux du bilan précédent.

| Indicateur | Bilan publié (App) | Nouvelle mesure avant (AppAll, 14 Stable) |
| --- | ---: | ---: |
| Ratio des totaux | 0.130413× | 0.139809× |
| μ | 5.116700× | 5.045127× |
| Q₂ | 0.148519× | 0.157630× |
| G | 0.133042× | 0.135576× |

Le bilan publié donnait 32,225 ns / 47,41× pour Polymorphism et 26,366 ns / 15,90× pour RowToList. L’accélération attribuée à ce changement repose sur les deux variantes de la nouvelle campagne alternée.

## Changements et validation

- Les effets FFI construits puis immédiatement exécutés passent par un appel direct qui évite les closures intermédiaires. L’ordre d’évaluation et les conversions polymorphes sont conservés ; les actions sauvegardées ou rejouées gardent leur représentation existante.
- Les fonctions locales appelées uniquement à saturation gardent leur fonction native et n’allouent plus un adaptateur currié inutilisé. Les applications partielles, alias, valeurs retournées et usages indirects conservent leur adaptateur. Les types publiés des adaptateurs récursifs sont aussi corrigés.
- Ces règles sont générales : aucun nom de benchmark ni aucune valeur attendue n’est utilisé dans le compilateur.
- Reconstruction native réussie ; 18 tests ciblés passent ; les programmes de régression produisent les mêmes résultats via les compilateurs JS et natif, en Release et Dev.
- Les six processus valident les 37 oracles et les lots calibrés : 222 résultats vérifiés. Aucun des 37 cas ne présente une hausse de médiane supérieure à 5 %. La plus forte hausse observée est `string/concat-1000` (+3,13 %).
- Les mêmes transformations bénéficient aussi à List Processing (−49,81 %), Array Processing (−48,88 %), Deep Record Updates (−28,69 %) et AST Evaluation (−24,15 %).
- Les 126 empreintes figées des sources de benchmarks et références sont inchangées ; le runner vérifie également les sources `AppAll`, les profils, les manifestes et les exécutables. Le catalogue historique des temps reste inchangé.

## Protocole et preuves

Go 1.27.0, darwin/arm64, Release, GOGC=800, sans PGO, GOMAXPROCS par défaut. Trois processus indépendants par variante dans l’ordre AB/BA/AB, échauffements et calibration du harnais existant, minimum des dix lots chronométrés puis médiane des trois minima de processus. Les builds, tests et sondes d’allocation étaient terminés avant le chronométrage ; la machine n’était pas isolée des autres activités utilisateur.

Le C n’a pas été relancé. L’[archive de référence](2026-11-25-go-c-js-functional.json) conserve le SHA-256 `87efaa7336c9b727a6298da9fb3be68c9fd6f6820c492faf00d29c301c9a0350`. Les limites de comparaison des modèles d’allocation et des harnais restent celles du contrat fonctionnel.

L’[archive de cette optimisation](2026-11-25-go-small-actions.json) contient les six sorties brutes, les échantillons des 37 cas, les quatre agrégats, les manifestes de build, le changement du compilateur et ses sources, les actions Go générées, les sondes d’allocation et les logs de validation. Les données JS incidentes du runner proviennent de références historiques : elles ne constituent pas une nouvelle mesure JS.

Pour rejouer les binaires figés dans un nouveau répertoire de résultats :

```sh
env -u GOMAXPROCS -u PPROF GOGC=800 python3 bin/benchmark/js-ratios.py \
  --before var/benchmark/go-small-actions-20261125/before-all \
  --after var/benchmark/go-small-actions-20261125/after-all \
  --output var/benchmark/go-small-actions-20261125/replay
```

Les quatre métriques se recalculent à partir des 14 entrées `category == "Stable"` de `campaign.all_tests`, de leurs `medians_us` et de `frozen_reference.c_functional_times_us`.
