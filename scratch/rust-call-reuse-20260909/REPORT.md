# Cellule transmise à la fonction appelée — 9 septembre 2026

Purust conserve maintenant la cellule du parent à travers un appel de reconstruction. Dans RBTree, `ins` la transmet à une variante privée de `balance`, qui y écrit le constructeur retourné. Sur cinq paires alternées du runner complet, **RBTree passe de 35,391 à 21,536 ms (−39,1 %)** et le **total de 57,762 à 43,840 ms (−24,1 %)**.

## Transformation intégrée

La génération découvre les fonctions locales non récursives dont tous les chemins de retour construisent un ADT natif local. `Purust.ReturnCells` remplace uniquement les constructeurs en position de retour, à travers `Typed`, `Branch` et `Let`, par des appels à des helpers de reconstruction. Le type des paramètres, du résultat et des champs provient du TAST. Les helpers écrivent le nouveau contenu dans la cellule disponible ; les constructeurs imbriqués restent générés normalement.

Une variante privée reçoit les arguments habituels et une cellule supplémentaire du même type natif. Les fonctions publiques conservent leurs signatures et leur corps normal. Les noms déjà présents, y compris dans les signatures FFI sans corps PureScript, désactivent la génération des helpers ou variantes en conflit.

Un appel direct saturé utilise cette variante seulement si le parent consommé a la même représentation native que le résultat et si toutes ses utilisations sont des projections de ses champs. `Rc::get_mut` vérifie son unicité. Le contenu est déplacé en utilisant un vrai constructeur vide de `dataDecls` comme contenu temporaire ; les enfants peuvent être transmis uniques, tandis que la cellule reste disponible jusqu'à l'appel. Le partage ou des références faibles conservent le chemin normal. Les types sans constructeur vide, les appels partiels et les fonctions retournant une valeur existante restent hors de cette règle.

Aucun pool, `unsafe`, changement de disposition des ADT ou modification de FFI native n'est introduit. `Tree` occupe toujours 32 octets.

## Preuve et validation

`proof.rs` a d'abord testé le passage explicite de la cellule à une fonction séparée : 1 000 allocations avant, zéro avec un parent unique, 1 000 avec un parent partagé. Le test vérifie valeurs, adresse de la cellule et libération.

La fixture permanente passe ensuite par le vrai fork TAST, le PBO, le CLI et des crates Rust fraîches. Ses directives empêchent le PBO d'intégrer les petites fonctions de reconstruction dans leur appelant : les tests conservent donc une vraie frontière d'appel.

| Fixture | Avant cette étape | Après |
| --- | ---: | ---: |
| 2 000 appels, parents uniques | 2 000 allocations | 0 |
| 2 000 appels, anciennes racines conservées | 2 000 allocations | 2 000 |
| 1 000 mises à jour sur deux nœuds uniques | 0 allocation | 0 |

Les assertions contrôlent les adresses réutilisées, les branches et valeurs retournées, la conservation des anciennes racines dans le PureScript, l'unicité des enfants à l'entrée des callbacks et leur ordre d'exécution. Les exceptions traversant les appels, uniques et partagés, libèrent toutes les allocations mesurées. Le mécanisme de panic Rust est initialisé avant ce comptage.

Un test supplémentaire compile les déclarations de noms FFI concurrents avec le Rust généré. Il reproduisait une double définition avant la garde finale et passe après. **16 tests de génération et 5 tests TAST passent.** `bin/rust/run -c` passe avec ses **14 résultats**, puis une régénération finale vérifie la garde des noms.

Le comptage séparé sur les noyaux générés vérifie les quatre rotations, l'ordre des clés, les invariants rouge/noir, les doublons et les clés exactes de **200 anciennes versions conservées simultanément**, après chaque insertion et après destruction de la dernière racine. Ces contrôles passent avant et après.

## Mesures

Rust 1.96.0, O1 et mimalloc. Cinq paires alternées de processus ; chaque processus s'échauffe puis garde le meilleur de dix essais par benchmark. Les chiffres ci-dessous sont les médianes des cinq processus ; le total est la somme des médianes par benchmark. Les 14 sorties sont contrôlées à chaque passage. Les compteurs d'allocations sont dans des binaires séparés du chronométrage.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| RBTree | 35,391 ms | 21,536 ms |
| LazyEvaluation | 19,432 ms | 19,370 ms |
| Total | 57,762 ms | 43,840 ms |
| Allocations RBTree | 2 683 933 | 499 957 |
| Octets demandés cumulés RBTree | 128 828 784 | 23 997 936 |

Toutes les allocations comptées de RBTree sont libérées. Le module RBTree avant est identique à celui validé après l'étape précédente de transfert des champs. Les empreintes des binaires mesurés sont conservées dans `results.json`.

Le README officiel relu donne **35,323 ms** pour RBTree compilé, **36,070 ms** pour le natif optimisé et **58,39 ms** au total compilé. Ces références historiques restent distinctes des nouvelles mesures appariées ; le natif n'a pas été remesuré ici.

## Suite et reproduction

La cellule reçue est réutilisée pour le résultat de `balance`. Les deux constructeurs imbriqués des rotations allouent encore : la prochaine étape consiste à isoler la réutilisation d'une cellule enfant sur une rotation, avec un enfant unique puis partagé, avant d'étendre le mécanisme. Le prototype exploratoire complet à 15,475 ms reste distinct de cette intégration.

- `python3 measure.py` : cinq paires des runners sauvegardés localement, avec contrôle des sorties.
- `python3 count.py` : recompilation des noyaux, contrôles de persistance et comptage.
- `RBTree-before.rs`, `RBTree-after.rs`, `RBTree.diff` : code observé avant/après.
- `results.json`, `allocations.json`, `validation.json` : mesures et empreintes des sources validées.

Les binaires, logs et compilations intermédiaires sont ignorés par Git.
