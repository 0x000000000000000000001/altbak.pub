# B13 : enfant modifié par un appel, intégration Purust

Le 9 septembre 2026, la règle intégrée au générateur donne **RBTree 18,309 → 15,741 ms** et **suite complète 19,419 → 16,835 ms**, soit **2,584 ms gagnées (13,3 %)**. Les cinq blocs de comparaison sont favorables, pour chacune des deux étapes. Le gain propre de la modification du seul enfant est **0,607 ms** sur la suite, après la première étape. B13 et B14 restent partiels.

## Mesures du code réellement généré

| Variante | RBTree, médiane | Suite, somme des médianes | Gain suite par rapport à la précédente |
| --- | ---: | ---: | ---: |
| Avant | 18,309 ms | 19,419 ms | — |
| Étape 1 : emprunt unique conservé et reconstruction directe | 16,339 ms | 17,442 ms | 1,977 ms |
| Étape 2 : ne réinstaller que l'enfant | 15,741 ms | 16,835 ms | 0,607 ms |

Trois binaires complets, issus de trois états du générateur, sont conservés sous `build/`. Le runner initial est sauvegardé avant intégration ; le deuxième vient de la génération de l'étape 1 ; le dernier de **`bin/rust/run -c`**, avec reconstruction du backend et régénération du TAST et du Rust. La mesure n'applique aucun patch au Rust des binaires. L'[expérience précédente](../rust-child-field-20260909/REPORT.md) était un prototype séparé.

Le profil reste **O1 / mimalloc**, avec le warm-up et le meilleur de dix mesures habituels du runner. [measure.py](measure.py) exécute cinq blocs, en variant l'ordre des trois processus ; il vérifie les **14 sorties** de chacun. Aucune compilation ni instrumentation ne tourne pendant cette série. Les totaux de la table sont les sommes des médianes de chaque benchmark, avec la même méthode pour les trois variantes. [Données complètes](final-results.json), [révisions et empreintes](metadata.json).

| Bloc | Total avant | Total étape 1 | Total étape 2 |
| --- | ---: | ---: | ---: |
| 1 | 18,999 ms | 17,381 ms | 16,709 ms |
| 2 | 19,810 ms | 17,322 ms | 17,022 ms |
| 3 | 19,414 ms | 17,567 ms | 17,090 ms |
| 4 | 19,655 ms | 17,434 ms | 16,841 ms |
| 5 | 19,298 ms | 18,231 ms | 16,814 ms |

L'écart propre à B13 varie de **0,300 à 1,417 ms** entre blocs : 0,607 ms est la différence des sommes de médianes, pas une constante par exécution. La [première série exploratoire](stage1-results.json) avait donné 18,557 → 16,856 ms pour l'étape 1 ; un bref test de génération a pu chevaucher cette série. Elle n'est pas agrégée à la série finale.

Le [README officiel du checkout normal](../../../altbak.pub/README.md#rust), relu le 9 septembre, donne **18,480 ms** pour RBTree compilé et **36,070 ms** dans la dernière colonne native ; totaux **19,62 et 36,13 ms**. Notre baseline appariée de 18,309 / 19,419 ms est proche de ce relevé. Le compilé dépassait déjà ce natif ; le gain ci-dessus compare le compilé à lui-même dans la même série. Le README historique est conservé.

## Règle et portée

[ChildCalls](../../../purust/purust/src/Purust/ChildCalls.purs) reconnaît une branche par défaut d'un helper qui reconstruit ses paramètres, après des tests de tags sur des enums `Copy`. Les alias locaux, annotations préservant la représentation et permutations bijectives sont analysés dans l'IR typé. Les autres branches du helper restent en place. Dans RBTree, cela reconnaît le cas où le parent rouge ne peut pas déclencher une rotation.

[ChildUpdates](../../../purust/purust/src/Purust/ChildUpdates.purs) exige le même constructeur et layout TAST, un parent consommé reconnu par l'analyse existante, tous les champs inchangés sauf un enfant, et un appel natif direct saturé qui consomme cet enfant une seule fois. Ses autres arguments sont des scalaires simples, éventuellement des expressions entières `+`, `-`, `*` sans conversion. Le graphe des appels doit être fermé dans le module : appels opaques ou étrangers, callbacks, applications partielles et closures le rejettent, y compris transitivement. Tous les champs du type récursif sont des scalaires `Copy` ou des enfants du même type.

[CodeGen](../../../purust/purust/src/Purust/CodeGen.purs) émet deux sites dans RBTree. Sur le chemin unique, `Rc::get_mut` reste emprunté pendant l'appel. La première étape reconstituait directement le nœud entier avec cet emprunt, en évitant le helper de balance et son retest. Son gain combine ces effets ; il ne démontre pas que la seule suppression d'un retest suffirait.

La deuxième étape laisse les autres champs à leur place. `mem::replace` déplace l'enfant vers l'appel récursif ; un clone temporaire du frère garde son emplacement initialisé, puis le résultat est réinstallé dans le seul champ modifié. Cela conserve l'unicité réelle de l'enfant transmis. Le Rust émis est sûr, sans `unsafe` ajouté. Sur panique, le parent et son emplacement temporaire restent destructibles ; aucun callback ou destructeur opaque du payload ne peut observer cet état dans le périmètre accepté.

Les parents partagés ou porteurs d'une référence faible gardent le chemin existant. Un enfant partagé reste soumis à sa propre décision de copie. Rotations, conversions, plusieurs champs changés et graphes d'appels non prouvés restent hors de cette extension. Le générateur ne contient aucun nom de benchmark ni reconnaissance de texte Rust.

## Comptage séparé

[count.py](count.py) instrumente les trois noyaux réellement générés avec le wrapper `Rc` de l'expérience précédente. Les nombres ci-dessous concernent la construction unique des 100 000 nœuds ; ce sont des opérations logiques, pas des instructions machine. [Résultats par phase](counts.json).

| Opération | Avant | Étape 1 | Étape 2 |
| --- | ---: | ---: | ---: |
| Allocations | 100 001 | 100 001 | 100 001 |
| Extractions du payload entier | 2 383 932 | 2 383 932 | 1 668 902 |
| Reconstructions entières, helper ou écriture directe | 2 383 932 | 2 383 932 | 1 668 902 |
| `get_mut` réussi | 4 867 864 | 4 152 834 | 4 152 834 |
| Clones | 3 060 100 | 3 060 100 | 3 775 130 |
| Relâchements temporaires | 2 960 100 | 2 960 100 | 3 675 130 |

**715 030** appels utilisent la nouvelle branche. La deuxième étape retire autant d'extractions/reconstructions complètes, au prix d'autant de clones et relâchements temporaires du frère. Les allocations et libérations finales restent identiques. Le gain provient donc de moins de travail sur les payloads, sans diminution des allocations ni du nombre total de clones.

## Validation et reproduction

La reconstruction complète réussit, suivie de **23 tests de génération et 14 tests TAST**, tous réussis. Les 14 résultats du runner sont également vérifiés dans chaque processus mesuré.

- [Contrats de génération](../../../purust/purust/tests/codegen/child-call-eligibility.mjs) : alias et annotations, gardes, bijections, conversions refusées, récursivité fermée, dépendances opaques transitives, callbacks/closures/applications partielles refusés. Un appel local fermé qui panique vérifie les durées de vie sur le chemin optimisé, des deux côtés de l'arbre.
- [Fixture TAST fraîche](../../../purust/purust/tests/tast/child-call-reuse.mjs) : descentes gauche/droite, 128 combinaisons de partage et de références faibles, versions persistantes, ordre des callbacks sur les chemins exclus, paniques et libérations ; 1 000 mises à jour uniques sans allocation.
- Noyaux générés instrumentés : quatre rotations, ordre et invariants rouge/noir, clés exactes de 200 versions conservées, exclusion des références faibles et équilibre global des allocations/handles.

Depuis Purust : `npm run test:codegen`, puis `PURS=/chemin/vers/le/fork/purs npm run test:tast`. Depuis le worktree de benchmark : `bin/rust/run -c`. Les snapshots locaux permettent ensuite `python3 scratch/rust-child-field-integration-20260909/count.py` et, séparément, `python3 scratch/rust-child-field-integration-20260909/measure.py before stage1 stage2 --output final-results.json`.

Le gain est établi pour le runner actuel, dominé par la construction unique de RBTree. La validité des versions fortement persistantes est vérifiée, mais leur performance ne l'est pas : les tests d'unicité supplémentaires et le partage temporaire du frère peuvent changer les coûts. Les résultats ne justifient ni B13/B14 verts, ni une optimisation générale des rotations ou des appels avec callbacks.
