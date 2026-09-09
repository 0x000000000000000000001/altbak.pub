# Reconstruction des nœuds consommés — 9 septembre 2026

Purust réutilise désormais la cellule d'un ADT natif lorsqu'une reconstruction projette les champs d'un local du même type et du même constructeur, à sa dernière utilisation. Les types du TAST déterminent la représentation native ; les usages vivants déterminent si le local peut être consommé.

Tous les nouveaux champs sont évalués avant de toucher l'ancien nœud. Le local reste vivant pendant cette évaluation, même lorsqu'un champ conserve sa valeur. Ensuite, `Rc::get_mut` permet de remplacer le contenu d'une cellule unique. En présence d'autres références fortes ou faibles, une nouvelle cellule contient les champs déjà calculés. Cela évite de cloner l'ancien contenu pour l'écraser aussitôt.

Cette première règle ne transfère pas encore les champs vers les appels récursifs. Elle garde le chemin précédent pour les bases calculées, les conversions de représentation, les locaux utilisés ultérieurement, les changements de constructeur, les classes et les enums en valeur.

## Fixture et régressions

La fixture passe par le vrai fork TAST, le PBO, le CLI et des crates Rust compilées dans un répertoire temporaire frais. Les allocations ci-dessous excluent la construction des entrées et sont comptées avec l'allocateur System, sans chronométrage.

| 1 000 mises à jour | Avant | Après |
| --- | ---: | ---: |
| Racine unique | 1 000 | 0 |
| Racine partagée à chaque appel | 1 000 | 1 000 |
| Chemin de deux nœuds uniques | 2 000 | 1 000 |

Le budget de zéro allocation reproduit d'abord un échec avec l'ancien générateur, puis passe avec la règle. Le cas à deux nœuds montre la limite restante : l'extraction de l'enfant le clone alors que le parent le possède encore, empêchant sa réutilisation pendant l'appel imbriqué.

Les contrôles vérifient les valeurs, les enfants partagés, les anciennes versions, les références faibles, les profondeurs de boucle, l'absence de cycle lorsque le résultat conserve l'ancienne racine, et l'équilibre allocations/libérations. Un test de génération complémentaire couvre les wrappers qui préservent la représentation, les conversions, les bases calculées évaluées une seule fois, les utilisations ultérieures et les autres constructeurs.

**15 tests de génération + 5 tests TAST passent**, ainsi que **`bin/rust/run -c` et ses 14 résultats**.

## Mesure dans altbak.pub-purust

Les temps proviennent de cinq paires alternées de runners complets, avec Rust 1.96.0, O1 et mimalloc. Chaque processus s'échauffe puis garde le meilleur de dix essais par benchmark. Le tableau donne les médianes sur cinq processus ; les totaux sont les sommes de ces médianes. Les 14 sorties sont contrôlées à chaque passage.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| RBTree | 40,714 ms | 40,603 ms |
| Total | 62,595 ms | 63,231 ms |
| Allocations RBTree | 2 783 933 | 2 683 933 |
| Octets demandés cumulés RBTree | 133 628 784 | 128 828 784 |

**100 000 allocations et 4 800 000 octets demandés cumulés sont supprimés. Aucun gain de vitesse n'est revendiqué.** RBTree reste stable et le total est supérieur de 1,0 % dans cette série. LazyEvaluation contribue principalement à cet écart (18,996 → 19,688 ms) ; son module Rust généré est identique avant/après. La série ne permet pas d'attribuer cet écart à une cause précise.

Les compteurs tournent séparément sur les noyaux Rust générés extraits sans réécriture, avec O1 et mimalloc. Toutes les allocations comptées sont libérées. Les quatre rotations, l'ordre des clés, la racine noire, les hauteurs noires, l'absence de rouges consécutifs, les doublons et la persistance après insertion et destruction passent sur les deux versions.

Dans RBTree, la règle apparaît dans `makeBlack`, la branche de doublon d'`ins` et la recoloration d'`insert`. Les rotations de `balance` restent inchangées. Les 100 000 insertions du benchmark sont distinctes ; la suppression d'allocations mesurée vient donc de la recoloration dans `insert`.

Le README officiel relu indique RBTree compilé **41,482 ms**, natif optimisé **36,070 ms**, total compilé **64,69 ms**. L'utilisateur a retiré le pool de la FFI native : les anciens chiffres natifs de l'audit ne décrivent plus cette référence. Cette étape mesure exclusivement le généré avant/après et ne modifie pas la FFI.

Le prototype exploratoire à 15,475 ms réutilisait aussi les cellules du chemin et des rotations. Cette extension reste à intégrer ; son résultat n'est pas attribué à la règle actuelle.

## Reproduction

- `python3 measure.py` : cinq paires des binaires sauvegardés localement, sorties contrôlées.
- `python3 count.py` : recompilation des noyaux sauvegardés avec mimalloc, contrôles et comptage.
- `RBTree-before.rs`, `RBTree-after.rs`, `RBTree.diff` : Rust observé et changement exact.
- `results.json`, `allocations.json`, `validation.json` : observations et empreintes.

Les binaires, logs et fichiers de compilation sont ignorés par Git. Les sources et tests de la règle se trouvent dans Purust ; les preuves sont dans le worktree Rust d'altbak.
