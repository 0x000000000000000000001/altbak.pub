# GetCtorField : emprunter le parent local

Le 8 septembre 2026, l'extension aux extractions de champs fait passer RBTree de **54,426 à 45,978 ms (−15,5 %)** dans le runner complet. Le total passe de **170,980 à 163,252 ms (−4,5 %)**. Les tests d'invariants du Rust généré et les 14 résultats du runner sont corrects.

## Changement intégré

Dans [CodeGen.purs](../../../purust/purust/src/Purust/CodeGen.purs), `GetCtorField` partage avec `OpIsTag` la détection des variables locales natives en `Rc`. Les wrappers `Typed` et `TypeApp` sont traversés lorsqu'ils préservent la représentation. La lecture émet `(parent).as_ref()` à la place de `(parent.clone()).as_ref()` ; le champ extrait garde **`f.clone()`**, qui lui permet de survivre au parent et de rester partagé entre plusieurs versions de l'arbre.

La règle s'applique au générateur, sans exception propre à RBTree. Elle conserve le chemin d'évaluation habituel pour les bases non locales et les conversions de représentation. Les enums, allocations de couleurs et fonctions de construction restent identiques.

Le module RBTree régénéré passe de **426 à 277 occurrences statiques de `.clone()`**, soit **149 clones de parents supprimés**. Ses **260 clones de champs** sont tous conservés. Le script [check-rbtree.py](check-rbtree.py) vérifie que le diff entier de ce module correspond exclusivement à la suppression des clones locaux précédant `.as_ref()` ; voir [RBTree.diff](RBTree.diff) et [emission.json](emission.json). Ces nombres comptent des sites dans le source, pas les appels exécutés.

## Preuves et validation

Le nouveau test [field-borrows.mjs](../../../purust/purust/tests/codegen/field-borrows.mjs) a été exécuté avant le changement : son Rust compilait et ses contrôles de persistance passaient, mais l'assertion sur l'émission échouait sur `purs_local_0.clone()` ([log avant](test-before.log)). Après intégration, il passe avec les 12 autres tests du générateur ([log des 13 tests](tests.log)).

Les contrôles ciblés couvrent les parents réutilisés après extraction, le partage d'un enfant dans les deux branches, la reconstruction d'une nouvelle racine sans modifier l'ancienne, la survie de l'enfant après destruction du parent, les champs scalaires, les motifs imbriqués, les conversions via `Value` et un appel utilisé comme base évalué exactement une fois. Les tests d'`OpIsTag` couvrent aussi les branches alternatives et les lectures répétées.

Le [test du noyau RBTree régénéré](rbtree-invariants.log) reprend les contrôles de l'audit sur les fonctions extraites sans modification : 1 000 insertions mélangées, ordre des clés, égalité des hauteurs noires, absence de rouges consécutifs, insertion d'un doublon, conservation d'une ancienne version après insertion supplémentaire.

`bin/rust/run -c` a été exécuté depuis `altbak.pub-purust` : reconstruction du générateur, nettoyage, régénération et compilation Rust réussis, **14 résultats corrects** ([log](clean-run.log)). Les six exécutions chronométrées vérifient également ces 14 sorties. Les checkouts habituels d'altbak.pub et PBO n'ont pas été modifiés par cette étape.

## Mesures

Même profil `release` à `opt-level=1`, même allocateur mimalloc ; `rustc 1.96.0`. Trois processus successifs avant et trois après, sans compilation ou autre benchmark lancé par cette tâche en parallèle. Chaque processus retient le meilleur de dix essais par benchmark ; le tableau donne la médiane des trois résultats. Le total est la somme des médianes. Scripts, valeurs individuelles et empreintes des sources sont conservés dans [measure.py](measure.py), [results.json](results.json) et [metadata.json](metadata.json).

| Benchmark | Avant | Après | Variation |
| --- | ---: | ---: | ---: |
| RBTree | 54,426 ms | 45,978 ms | −15,5 % |
| LazyEvaluation | 67,668 ms | 67,999 ms | +0,5 % |
| Polymorphism | 36,755 ms | 37,054 ms | +0,8 % |
| Church | 10,835 ms | 10,907 ms | +0,7 % |
| Total | 170,980 ms | 163,252 ms | −4,5 % |

RBTree, valeurs individuelles en ms : avant **54,441 / 54,311 / 54,426** ; après **46,691 / 45,942 / 45,978**. Le gain observé se concentre sur RBTree ; les autres gros benchmarks restent proches de leur référence. Cette série part du générateur intégrant déjà Unit natif et l'emprunt des tests de constructeur. Elle mesure l'ajout des emprunts pour les champs ; on ne doit pas additionner ses pourcentages aux séries précédentes.

Le [README officiel d'altbak.pub](../../../altbak.pub/README.md#rust), relu pendant cette étape, donne désormais **RBTree 62,425 ms**, **LazyEvaluation 76,344 ms**, **Polymorphism 38,940 ms**, **Church 11,792 ms**, **total 190,94 ms**. Ces repères ont été mis à jour depuis l'audit initial (RBTree 67,125 ms, total 452,29 ms). Les mesures actuelles sont en dessous de ces repères, mais seule la comparaison avant/après ci-dessus sert à attribuer un gain à ce changement. Le score FFI manuel de 16,700 ms utilise une représentation différente et reste un repère distinct.

L'étape 2 du [todo](../../../purust/purust/todo.md) est terminée. La prochaine piste planifiée est l'éligibilité des enums sans charge utile à une représentation en valeur, à commencer par un cas minimal de `Color` fondé sur `dataDecls`.
