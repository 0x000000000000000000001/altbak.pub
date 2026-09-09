# Appeler les fonctions locales natives par emprunt

L’appel par emprunt améliore le runner complet : **LazyEvaluation 19,406 → 17,399 ms (−10,3 %)**, **Church 1,661 → 1,507 ms (−9,3 %)**, **total 41,129 → 39,108 ms (−4,9 %)** sur cinq paires alternées.

## Changement intégré

Dans [CodeGen.purs](../../../purust/purust/src/Purust/CodeGen.purs), un appel complet à une fonction locale représentée par `Func1` à `Func10` utilise maintenant le récepteur sans clone. Le runtime appelle ces fonctions par une référence à `dyn Fn`. Les annotations et l’arité du TAST déterminent la représentation et permettent de distinguer un appel complet d’une application partielle.

La variable reste disponible pendant l’évaluation des arguments : si un argument transmet ou capture la fonction appelée, il conserve le clone nécessaire. Les applications partielles continuent à posséder leur fonction ; les récepteurs calculés et les conversions de représentation gardent leur chemin d’évaluation. La vérification des locaux sans conversion est partagée avec les emprunts d’ADT déjà existants.

Le [diff de LazyEvaluation](LazyEvaluation.diff) contient un seul changement : la fonction capturée est appelée sans `.clone()` dans chaque thunk. Les thunks et leur forçage restent présents. Le [diff de Church](Church.diff) supprime aussi un clone de la fonction capturée dans `fromInt` et conserve les clones nécessaires lorsque le récepteur est transmis en argument. Le Rust de RBTree reste identique octet pour octet.

## Preuve et mesures

L’[audit préalable](prototype-results.json) avait testé le changement exact sur le noyau généré : **15,897 → 14,943 ms (−6,0 %)**, trois paires alternées, quinze mesures après échauffement par processus, même O1 et mimalloc. Ces résultats isolés restent distincts des mesures du runner complet.

Le runner complet actuel a été sauvegardé et exécuté avant modification. Ses 14 résultats et ses noms de benchmarks ont été vérifiés, puis la version après a été obtenue par `bin/rust/run -c`. Le binaire avant diffère de l’empreinte d’une précédente étape ; la comparaison utilise les deux exécutables sauvegardés dans cette étape, identifiés dans [metadata.json](metadata.json) et [results.json](results.json).

Même configuration avant/après : Rust 1.96.0, release `opt-level=1`, mimalloc d’origine. Cinq paires de processus complets, avec ordre alterné ; chaque processus s’échauffe puis retient le meilleur de dix essais par benchmark. Les chiffres sont les médianes par benchmark et le total additionne ces médianes. Aucune mesure n’a été exclue.

| Benchmark | Avant | Après | Évolution |
| --- | ---: | ---: | ---: |
| LazyEvaluation | 19,406 ms | 17,399 ms | −10,3 % |
| Church | 1,661 ms | 1,507 ms | −9,3 % |
| RBTree, code inchangé | 18,750 ms | 18,892 ms | +0,8 % |
| Suite complète | 41,129 ms | 39,108 ms | −4,9 % |

Le [comptage séparé](allocations.json), sur les deux noyaux réellement générés, confirme **1 000 000 allocations et libérations**, et **32 000 000 octets demandés cumulés**, avant comme après. Ce gain porte sur le travail aux appels, pas sur le nombre de thunks. Le comptage inclut la construction, le forçage et la destruction. Les appels répétés avec un thunk d’entrée partagé et plusieurs profondeurs produisent les résultats attendus.

Le binaire passe de **1 509 120 à 1 509 088 octets**. Les résultats et noms des 14 benchmarks sont vérifiés dans la référence, dans le runner propre après régénération et dans les dix processus de mesure.

Le [README officiel relu](../../../altbak.pub/README.md) indique LazyEvaluation compilé **19,791 ms**, Church **1,592 ms** et total **42,74 ms**. La dernière colonne native affiche respectivement environ **0 µs**, **1 µs** et **36,13 ms**. La FFI native de LazyEvaluation additionne directement 1 000 par tour et ne construit pas de thunks ; son résultat arrondi ne donne pas un objectif temporel pour cette modification. Le natif n’a pas été remesuré ici. Le dossier généré du checkout normal était ancien lors de l’audit ; les essais utilisent le worktree à jour.

## Validation

**18 tests de génération + 9 tests TAST passent**, ainsi que **`bin/rust/run -c` et les 14 résultats attendus**.

La [fixture TAST](../../../purust/purust/tests/tast/function-borrows.mjs) traverse le vrai fork, PBO, Purust et des crates Rust fraîches. Des compteurs de références contrôlent l’absence de clone pendant les appels de `Func1` et `Func2`, y compris dans les captures rejouées et les appels imbriqués. Elle vérifie aussi les arguments qui transmettent ou capturent le récepteur, les usages ultérieurs, les applications partielles conservées après le retour et l’expiration des références faibles.

Le [test de génération](../../../purust/purust/tests/codegen/function-borrows.mjs) exécute le Rust pour vérifier les wrappers `Typed`/`TypeApp`, les conversions, un récepteur calculé évalué une seule fois avant son argument, et la libération des deux fonctions lorsqu’un argument déclenche une exception.

Le [patch d’intégration](integration.patch), les sources et les preuves sont identifiés dans [validation.json](validation.json). Les exécutables, logs et compilations intermédiaires sont ignorés par Git.

## Reproduction

Depuis la racine d’altbak.pub-purust, avec les dépendances du runner généré :

```sh
python3 scratch/rust-function-borrow-20260909/prototype.py
python3 scratch/rust-function-borrow-20260909/count.py
python3 scratch/rust-function-borrow-20260909/measure.py nouvelle-serie
```

Les deux premiers scripts recompilent les noyaux enregistrés. Le dernier utilise les deux runners sauvegardés. Pour reconstruire ces runners sans les exécutables locaux ignorés, utiliser les sources avant/après identifiées par les révisions de référence et `integration.patch`, puis sauvegarder chaque binaire avant le chronométrage.
