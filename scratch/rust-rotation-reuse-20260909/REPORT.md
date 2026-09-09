# Cellules enfants des rotations — 9 septembre 2026

Purust réutilise maintenant les cellules enfants consommées dans les constructions imbriquées des rotations. Sur cinq paires alternées du runner complet, **RBTree passe de 21,312 à 19,529 ms (−8,4 %)** et le **total de 42,800 à 41,021 ms (−4,2 %)**. Les allocations de RBTree passent de **499 957 à 300 001 (−40,0 %)**.

## Preuve et transformation

`proof.rs` vérifie le mécanisme sur un petit programme Rust : conserver une cellule extraite permet de passer de deux à une allocation ; le cas partagé garde deux allocations et ses anciennes valeurs. La première fixture TAST a ensuite montré que cette rotation directe était déjà optimisée par Purust. Le cas encore coûteux était la rotation gauche-gauche dans les branches factorisées du vrai `balance` : **trois allocations**, même avec les deux enfants uniques.

Dans ce code, des `let` recopient les champs du parent et du petit-enfant avant les trois constructions. Le générateur reconnaît désormais une chaîne de projections stricte au début d'un `let`, avec les types et champs natifs issus du TAST. Toutes les utilisations du parent dans cette expression doivent pouvoir devenir des utilisations de ses champs. Un usage du parent entier, ou sa présence dans les valeurs encore vivantes après l'expression, bloque le transfert.

Pour un parent unique, `Rc::get_mut` et le helper existant `__purust_take` extraient le contenu en conservant la cellule. `ReturnCells.reuseNestedConstructor` affecte cette cellule à une seule construction compatible, en donnant la priorité aux constructions imbriquées. Les enfants extraits peuvent à leur tour fournir leur cellule. Le générateur emploie un registre explicite des helpers éligibles, distinct des signatures ordinaires et FFI.

Cette recherche reste limitée aux constructions strictes, aux continuations des `let` et aux arguments des helpers générés. Elle ne traverse pas les closures, branches ou appels arbitraires. Si le parent est partagé ou possède une référence faible, le chemin normal est conservé. Les types sans constructeur vide approprié gardent aussi leur chemin antérieur. Aucun pool, `unsafe` de production, changement de représentation d'ADT ou modification de FFI native n'est ajouté.

## Validation

La fixture permanente passe par le vrai fork TAST, le PBO, le CLI et des crates Rust fraîches. Elle reproduit les quatre branches de `balance`. La rotation gauche-gauche passe de **3 à 1 allocation** pour deux cellules uniques ; une racine partagée reste à **3**.

Après intégration, les quatre orientations donnent exactement la même matrice :

| Propriété des deux cellules candidates | Allocations |
| --- | ---: |
| Toutes deux uniques | 1 |
| Racine conservée dans une ancienne version | 3 |
| Petit-enfant conservé dans une ancienne version | 2 |
| Référence faible sur la racine | 3 |
| Référence faible sur le petit-enfant | 2 |

Les tests vérifient les clés et couleurs, les adresses des deux cellules réutilisées, les anciennes valeurs avant et après destruction du résultat et la libération complète des allocations. Le cas sans rotation conserve ses enfants et sa couleur.

**16 tests de génération et 6 tests TAST passent**, ainsi que **`bin/rust/run -c` et ses 14 résultats**. Les tests existants couvrent aussi les callbacks, captures, conversions et collisions avec les noms FFI.

Le contrôle séparé des noyaux RBTree vérifie les quatre rotations, l'ordre des clés, les invariants rouge/noir, les doublons et les clés exactes de **200 anciennes versions retenues simultanément**, après chaque insertion et après destruction de la dernière racine. Ces contrôles passent avant et après.

## Mesures

Rust 1.96.0, O1 et mimalloc. Cinq paires alternées de processus ; chaque processus s'échauffe puis garde le meilleur de dix essais par benchmark. Les chiffres sont les médianes des cinq processus ; le total est la somme des médianes par benchmark. Les 14 sorties sont contrôlées à chaque passage. Les compteurs d'allocations sont dans des binaires distincts du chronométrage.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| RBTree | 21,312 ms | 19,529 ms |
| LazyEvaluation | 18,588 ms | 18,636 ms |
| Total | 42,800 ms | 41,021 ms |
| Allocations RBTree | 499 957 | 300 001 |
| Octets demandés cumulés RBTree | 23 997 936 | 14 400 048 |
| Taille du binaire complet | 1 423 248 octets | 1 509 120 octets |

Toutes les allocations comptées sont libérées. `Tree` occupe toujours 32 octets. Les chemins spécialisés augmentent la taille du binaire d'environ 6 % ; le gain de temps ci-dessus est mesuré avec cette version.

Le module RBTree avant est identique à celui validé après l'étape précédente de passage de cellule à `balance`. Les empreintes des deux binaires sont conservées dans `results.json`.

Le README officiel relu donne **21,720 ms** pour RBTree compilé, **36,070 ms** pour le natif optimisé et **45,25 ms** au total compilé. Ces repères historiques restent distincts de la nouvelle série appariée ; le natif n'a pas été remesuré ici.

## Suite et reproduction

Les cellules supplémentaires des rotations sont désormais réutilisées. Le prochain test doit isoler les allocations d'une insertion dans l'arbre vide, notamment les constructeurs sans champ, avant d'envisager leur partage. Le prototype exploratoire complet à 15,475 ms reste distinct du générateur intégré.

- `python3 measure.py` : cinq paires des runners sauvegardés localement, avec contrôle des sorties.
- `python3 count.py` : recompilation des noyaux, contrôles de persistance et comptage.
- `RBTree-before.rs`, `RBTree-after.rs`, `RBTree.diff` : code observé avant/après.
- `results.json`, `allocations.json`, `validation.json` : mesures et empreintes des sources validées.

Les binaires, logs et compilations intermédiaires sont ignorés par Git.
