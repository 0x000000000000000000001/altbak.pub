# Enums en valeur : activation globale

Le 8 septembre 2026, les enums sans champs sont activés dans le runner Rust d'altbak. RBTree passe de **47,509 à 42,687 ms (−10,1 %)** sur trois exécutions avant/après. Le total reste **170,025 → 170,126 ms**, globalement stable dans cette série. Le comptage indépendant confirme **499 934 allocations supprimées** sur le noyau RBTree.

## Intégration

[DataLayout.purs](../../../purust/purust/src/Purust/DataLayout.purs) collecte les types qualifiés éligibles de tous les modules. [Main.purs](../../../purust/purust/src/Main.purs) transmet cet ensemble à l'émission des modules et aux signatures des FFI de remplacement. La sélection utilise les `dataDecls` du TAST, sans traitement propre à RBTree.

Huit enums générés sont vérifiés comme `Copy` : `Weekday`, `Month`, `NoArguments`, `Parity`, `DurationComponent`, `Ordering`, `Color` et `Proxy` ([liste qualifiée](eligible-enums.json)). Les ADT contenant des champs conservent leur `Rc`, et `Unit` conserve `()`.

Les FFI réelles [Data.Ord](../../../purust/purust-prelude/src/Data/Ord.rs) et [Data.String.Common](../../../purust/purust-strings/src/Data/String/Common.rs) utilisaient explicitement `Rc<Ordering>`. Leurs signatures passent à `Ordering` en valeur. Leurs corps et leur comportement sont conservés, y compris le cas `NaN` de la comparaison numérique. La comparaison aux copies antérieures vérifie que seules ces occurrences de types ont changé dans ces deux fichiers.

À la frontière polymorphe existante, une couleur est stockée dans `Value::Class(Rc::new(Color))`, puis récupérée comme `Color`. Cette frontière conserve son allocation dynamique ; les chemins natifs utilisent la valeur directement.

## Tests et résultat fonctionnel

Le [test d'intégration](../../../purust/purust/tests/tast/value-enums-interop.mjs) compile de vrais modules PureScript avec le fork TAST, exécute le CLI Purust, puis compile séparément les crates Rust produites. Il vérifie :

- une couleur déclarée dans un module et construite/utilisée dans un autre ;
- les paramètres, retours et champs d'arbre ;
- un appel polymorphe passant par `Value`, la copie de cette boîte et la destruction de l'original ;
- une FFI recevant et renvoyant une couleur, et la signature d'une FFI de remplacement ;
- les FFI réelles de comparaison des entiers, nombres, caractères, chaînes et booléens ;
- `Unit`, les sous-arbres partagés et la conservation d'une ancienne racine.

Le test détectait l'ancien `Rc<Color>` avant activation ([log avant](interop-before.log)) et passe désormais ([log après](interop-after.log)). Il utilise un répertoire temporaire pour ses sorties et métadonnées PBO.

Les [14 tests de génération](codegen-tests.log) et les [2 tests TAST](tast-tests.log) passent. `bin/rust/run -c`, exécuté dans `altbak.pub-purust`, régénère et compile le runner avec **14 résultats corrects** ([log](clean-run.log)). Les six exécutions chronométrées vérifient également les 14 sorties. Les checkouts habituels d'altbak.pub et PBO n'ont pas été modifiés par cette tâche.

## Temps du runner complet

Même profil `release`, `opt-level=1`, allocateur mimalloc, `rustc 1.96.0`. Trois processus successifs avant et trois après ; chaque processus retient le meilleur de dix essais par benchmark. Le tableau donne la médiane des trois valeurs ; le total est la somme des médianes. Le comptage des allocations et les compilations ont été exécutés séparément des mesures après changement.

| Benchmark | Avant | Après | Variation |
| --- | ---: | ---: | ---: |
| RBTree | 47,509 ms | 42,687 ms | −10,1 % |
| Polymorphism | 38,245 ms | 40,261 ms | +5,3 % |
| LazyEvaluation | 71,837 ms | 74,395 ms | +3,6 % |
| Church | 11,157 ms | 11,478 ms | +2,9 % |
| Total | 170,025 ms | 170,126 ms | +0,1 % |

RBTree, valeurs individuelles en ms : **47,509 / 47,817 / 47,218** avant ; **42,687 / 43,233 / 41,685** après. Le bénéfice observé sur RBTree est compensé dans le total par des temps supérieurs sur les autres gros benchmarks. Cette série ne suffit pas à expliquer ces autres écarts ; aucun gain global n'est revendiqué. Les groupes avant/après sont successifs, pas alternés. [Script](measure.py), [résultats complets](results.json), [versions et empreintes](metadata.json).

Le [README officiel d'altbak.pub](../../../altbak.pub/README.md#rust), relu lors de l'étape, donne **RBTree 62,425 ms**, **Polymorphism 38,940 ms**, **LazyEvaluation 76,344 ms**, **Church 11,792 ms**, **total 190,94 ms**. Ces valeurs servent de repères historiques ; seule la série avant/après ci-dessus mesure ce changement. Les pourcentages des étapes précédentes ne s'additionnent pas à celui-ci.

## Allocations et invariants RBTree

[check-rbtree.py](check-rbtree.py) extrait les déclarations et fonctions du noyau généré avant/après sans les réécrire. Il les compile à `opt-level=1` avec mimalloc et un compteur d'allocations, puis vérifie les invariants hors comptage. Le test couvre les quatre rotations, 1 000 insertions mélangées, l'ordre des clés, une racine noire, l'égalité des hauteurs noires, l'absence de rouges consécutifs, un doublon et la conservation d'une ancienne racine après insertion et destruction de la nouvelle.

| Mesure pour 100 000 insertions + profondeur + destruction | Avant | Après |
| --- | ---: | ---: |
| Allocations | 3 283 867 | 2 783 933 |
| Libérations | 3 283 867 | 2 783 933 |
| Octets demandés cumulés | 145 627 200 | 133 628 784 |
| Allocations de 24 octets | 499 934 | 0 |
| Taille de `Tree` | 32 octets | 32 octets |

La réduction est de **499 934 allocations** et **11 998 416 octets demandés cumulés**. Ces octets ne représentent pas une baisse mesurée de mémoire simultanément résidente. Les invariants passent pour les deux versions ; allocations et libérations s'équilibrent. [Résultats](allocations.json), [log](rbtree-checks.log), [diff du Rust](RBTree.diff), [validation](validation.json).
