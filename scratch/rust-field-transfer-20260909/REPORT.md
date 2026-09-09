# Transfert des champs consommés — 9 septembre 2026

Purust déplace maintenant les champs d'un nœud consommé pour préserver l'unicité des enfants pendant les appels. Sur cinq paires alternées du runner complet, **RBTree passe de 39,970 à 34,794 ms (−13,0 %)** et le **total de 61,685 à 56,623 ms (−8,2 %)**.

## Génération

`Purust.OwnedFields` reconnaît les projections directes de locaux natifs, avec les types des champs issus du TAST. Toutes les utilisations du parent dans l'expression doivent pouvoir être remplacées par des projections du même constructeur. Une utilisation du parent entier, y compris dans une capture, désactive le transfert anticipé. Les champs deviennent des locaux typés ; l'analyse des utilisations conserve les clones nécessaires lorsqu'un champ est réutilisé.

- Pour une reconstruction, `Rc::get_mut` vérifie l'unicité. Un constructeur sans champ provenant de `dataDecls` permet d'extraire temporairement le contenu de la cellule avec `mem::replace`, puis de réécrire le résultat dans la même cellule. Les types sans constructeur vide gardent le chemin précédent. Les références partagées conservent la reconstruction avec copie.
- Pour un appel standard dont l'expression utilise tous les champs, `Rc::unwrap_or_clone` déplace le contenu d'un parent unique, ou le clone lorsqu'il est partagé. Les enfants peuvent ainsi arriver uniques dans l'appel récursif. Les projections partielles gardent le chemin précédent.

Les déclarations des ADT et les signatures publiques des fonctions gardent leurs représentations. Un helper de prise du contenu est généré depuis les déclarations ; il utilise explicitement `std::option::Option` pour éviter les collisions avec un type du programme. Aucun `unsafe` ni constructeur synthétique n'est ajouté aux ADT.

## Fixture et validation

Un prototype isolé sur la fixture réellement générée a d'abord vérifié le transfert sans allocation. L'intégration passe par le vrai fork TAST, le PBO, le CLI et des crates Rust fraîches.

| 1 000 mises à jour | Avant cette étape | Après |
| --- | ---: | ---: |
| Racine unique | 0 allocation | 0 |
| Racine partagée à chaque appel | 1 000 allocations | 1 000 |
| Chemin de deux nœuds uniques | 1 000 allocations | 0 |

Les tests contrôlent les valeurs et la persistance, les enfants partagés ou réutilisés, les anciennes racines conservées dans le résultat, les références faibles lors de la reconstruction et les captures différées. Un callback vérifie que chaque enfant unique a bien un seul propriétaire à son entrée. Une déclaration PureScript nommée `Option` vérifie les noms des helpers. Toutes les allocations de la fixture sont libérées.

**15 tests de génération et 5 tests TAST passent**, ainsi que **`bin/rust/run -c` et ses 14 résultats**. Les tests de génération couvrent aussi les bases calculées, les conversions et les utilisations ultérieures qui conservent le chemin précédent.

Le contrôle séparé des noyaux RBTree vérifie les quatre rotations, l'ordre des clés, les invariants rouge/noir, les doublons et les clés exactes de **200 anciennes versions retenues simultanément**, après chaque insertion et après destruction de la dernière racine. Il passe avant et après.

## Mesures

Même Rust 1.96.0, O1 et mimalloc. Cinq paires alternées de runners complets ; chaque processus s'échauffe puis garde le meilleur de dix essais par benchmark. Le tableau donne les médianes sur cinq processus ; le total est la somme de ces médianes. Les 14 sorties sont vérifiées à chaque passage. Le binaire avant correspond exactement au binaire après validé lors de l'étape précédente.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| RBTree | 39,970 ms | 34,794 ms |
| LazyEvaluation | 18,823 ms | 18,969 ms |
| Total | 61,685 ms | 56,623 ms |
| Allocations RBTree | 2 683 933 | 2 683 933 |
| Octets demandés cumulés RBTree | 128 828 784 | 128 828 784 |

Le comptage est effectué dans des binaires distincts du chronométrage, sur les fonctions générées extraites sans réécriture. Toutes les allocations comptées sont libérées ; `Tree` occupe toujours 32 octets. Le gain de temps ne provient donc pas d'une réduction du nombre d'allocations dans RBTree à cette étape.

Dans `ins`, les deux branches récursives prennent désormais les champs du parent avant d'appeler `balance`, au lieu de cloner les références vers les enfants tout en gardant le parent. La cellule de ce parent est libérée ; `balance` alloue encore la nouvelle cellule et celles des rotations. Conserver ces cellules à travers les appels reste le prochain travail. Le prototype exploratoire à 15,475 ms réalisait aussi cette réutilisation et ne décrit pas encore le générateur intégré.

Le README officiel relu indique **41,482 ms** pour RBTree compilé, **36,070 ms** pour le natif optimisé et **64,69 ms** au total compilé. Ces repères restent distincts de la série avant/après. Le pool natif a été retiré par l'utilisateur ; l'ancien chiffre natif de 16,700 ms n'est plus la référence actuelle. La FFI native n'est pas modifiée par cette étape.

## Reproduction

- `python3 measure.py` : cinq paires des binaires sauvegardés localement, avec contrôle des sorties.
- `python3 count.py` : recompilation des noyaux sauvegardés avec mimalloc, invariants et comptage.
- `RBTree-before.rs`, `RBTree-after.rs`, `RBTree.diff` : Rust observé et changement exact.
- `results.json`, `allocations.json`, `validation.json` : observations et empreintes.

Les binaires, logs et fichiers de compilation sont ignorés par Git. Les sources et tests sont dans Purust ; les preuves sont dans le worktree Rust d'altbak.
