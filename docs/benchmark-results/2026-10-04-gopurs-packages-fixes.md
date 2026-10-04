# gopurs : correction des quatre cellules Rust non qualifiées

Les quatre diagnostics du [rapport initial](2026-10-04-gopurs-packages-rust.md) sont résolus. Tous les hôtes génèrent du **Go**.

## Causes et corrections

- **Prelude** : la primitive de division entière de Purust utilisait `/` en Rust, qui panique sur zéro et tronque les quotients négatifs. Elle utilise désormais une division euclidienne contrôlée, avec zéro pour un diviseur nul. Les deux opérandes sont évalués une seule fois.
- **Enums, Promise et Strings** : l'ordre des répertoires différait entre hôtes. Le tri topologique conservait cet ordre pour les racines indépendantes, modifiant les rangs visibles par PBO et ses décisions d'inlining. Le tri part désormais d'un index ordonné par nom de module.

- **Zéro signé** : la comparaison supplémentaire avec JavaScript a révélé que les deux hôtes natifs transformaient la négation de zéro en soustraction à zéro lors du bootstrap de l'évaluateur PBO. Une négation primitive FFI conserve désormais le signe IEEE. L'oracle Prelude est corrigé explicitement sur une seule ligne de `Test_Main.go`, validée contre JavaScript et par exécution ; l'ancien oracle reste archivé. Le nouveau projet frais `CompilerHostNumbers` vérifie les zéros constants et dynamiques sous les trois hôtes.

## Validation

- Régressions reproduites avant correction ; **612 divisions** comparées à Prelude JavaScript et **12 cas modulo**, en Rust debug et optimisé.
- Tri vérifié sur **720 permutations**, imports propres/Prim et cycles ; **17 suites PBO** réussies.
- Purust reconstruit par bootstrap indépendant avec auto-compilation et projet frais. Gopurs reconstruit dans ses trois hôtes ; **1008 fichiers Rust/Cargo identiques** entre générateurs Purust JS et natif.
- PGO réentraîné sur **500 modules** de compilation du compilateur, en trois passes exactes ; `Test.Main` exclu.
- Qualification des hôtes, projets frais, Aff/AVar, tests du compilateur, helpers et parser Go sous détecteur de courses réussis ; détails et logs référencés dans le JSON.
- **50/50 paquets** revérifiés : **84 générations / 20283 fichiers Go exacts** aux oracles archivés, avec la correction explicite du zéro signé de Prelude ci-dessus. Les quatre cas corrigés passent aussi les hôtes JS/Go et le mode Rust séquentiel. Leurs quatre applications Go sont compilées et exécutées avec succès, hors chronomètre.

## Nouvelles mesures

Mêmes entrées TAST/FFI figées que la campagne initiale ; une chauffe Rust puis cinq mesures sérialisées, workers **8/8/8/8**, pipeline actif. Médiane de `backend total` : chargement, préparation, PBO, génération/écritures et drain. Frontend, construction des compilateurs/applications, exécution des applications et démarrage/arrêt des processus exclus.

| Paquet | Rust médian | Min–max | Fichiers exacts par génération |
| --- | ---: | ---: | ---: |
| gopurs-enums | 814 ms | 804–826 ms | 219 |
| gopurs-js-promise | 454 ms | 453–457 ms | 177 |
| gopurs-prelude | 775 ms | 768–782 ms | 214 |
| gopurs-strings | 1286 ms | 1275–1299 ms | 248 |

Total de la colonne Rust : **61.99 s, 50/50 paquets**. Il s'agit de la somme des médianes affichées provenant de leurs campagnes respectives. Les cellules JS/Go et les 46 médianes Rust préexistantes gardent leurs mesures d'origine ; aucun nouveau ratio inter-hôtes n'est déduit de ce tableau.

## Confirmation de performance sur Aff

Une campagne séparée de dix paires contemporaines, après une chauffe par hôte, donne **Go 1896.5 ms / Rust 1501.5 ms**, soit **-20.8 %**, avec **10/10 paires favorables à Rust**. Ordre Go/Rust alterné, mêmes workers et entrées figées ; les 22 générations conservent les 294 fichiers Go exacts. Cette confirmation ne remplace pas la campagne historique Aff du README.

Les échecs et rapports initiaux restent archivés. Entrées, binaires, sorties, profils PGO, journaux et contrôles de publication : `var/benchmark/gopurs-packages-fixes-20261004/`.
