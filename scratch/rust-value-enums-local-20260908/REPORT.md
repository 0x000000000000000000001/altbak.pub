# Enums en valeur : émission locale

Le 8 septembre 2026, le générateur sait émettre les enums sans champs en valeur dans un module minimal. `Color` est `Copy` et occupe un octet dans ce cas. La création et la lecture des couleurs n'allouent rien ; les nœuds d'arbre conservent leur `Rc` et leur partage structurel.

## Périmètre intégré

[DataLayout.purs](../../../purust/purust/src/Purust/DataLayout.purs) construit un ensemble de types qualifiés depuis les `dataDecls`. [CodeGen.purs](../../../purust/purust/src/Purust/CodeGen.purs) reçoit ce contexte explicitement via `codegenModuleWithValueEnums` et le transmet à ses fonctions d'émission. `codegenExprTypeWithValueEnums` rend les signatures correspondantes. La sélection reste indépendante du nom `Color` ; un type homonyme d'un autre module reste distinct.

Le contexte pilote la déclaration `#[derive(Clone, Copy)]`, les signatures des fonctions et callbacks, les champs des ADT, les constructeurs définis et saturés et les tests de constructeur. Les tests sur une couleur utilisent directement la valeur ; les champs des arbres sont toujours lus via un emprunt du parent. `Unit` conserve `()`.

L'entrée par défaut de Purust utilise encore un contexte vide. La prochaine étape transmettra le contexte global à tous les modules et aux signatures FFI, puis vérifiera les conversions dynamiques avant d'activer ce chemin dans le runner. Cette limite correspond au baby step local convenu.

## Preuves exécutées

Le [test minimal](../../../purust/purust/tests/codegen/value-enums.mjs) génère un module Rust avec `Color = R | B` et `Tree = E | T Color Tree Int Tree`, puis le compile et l'exécute avec le runtime du projet. Avant le changement, `rustc` rejetait l'exigence `Color: Copy` et les signatures attendues en valeur ([log avant](test-before.log)). Le nouveau chemin passe ([log après](test-after.log)).

Le test vérifie la réutilisation d'une couleur, les motifs appliqués à une variable et au résultat d'un appel, un callback natif, une fermeture capturant la couleur, les champs lus après construction, le partage d'un enfant entre plusieurs racines et la survie de la couleur après destruction de sa racine. Les assertions d'émission contrôlent l'absence de `Rc<Color>` et de conversions dynamiques dans ce cas typé.

Un allocateur instrumenté comptabilise les appels à l'allocateur système, avec `rustc` sans optimisation :

| Opération dans le cas minimal | Allocations |
| --- | ---: |
| 1 000 itérations de construction, passage et lecture des couleurs, avec `Unit` | 0 |
| Construction d'un nœud à partir d'une couleur et de sous-arbres existants | 1 |

La fermeture capturée est vérifiée séparément du comptage des couleurs. Ces mesures concernent le cas minimal et ne remplacent pas une mesure des allocations de RBTree complet.

Les [14 tests de génération](codegen-tests.log) et le [test TAST](tast-tests.log) passent. `bin/rust/run -c`, lancé dans `altbak.pub-purust`, réussit avec les **14 résultats attendus** ([log](clean-run.log)). Les empreintes des **301 sources Rust générées** sont identiques à celles d'avant l'étape ; le runner utilise encore son ABI actuelle ([validation](validation.json), [empreintes avant](rust-before.json)). Aucun gain de temps supplémentaire n'est attribué à cette étape. Les checkouts habituels d'altbak.pub et PBO n'ont pas été modifiés par cette tâche.
