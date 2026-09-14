# Prototype B13 : conserver les champs pendant une rotation gauche-gauche

14 septembre 2026. Le prototype isolé améliore les cinq comparaisons du runner complet : **RBTree 12,368 → 11,485 ms (−7,1 %)**, **suite 13,975 → 13,104 ms (−0,871 ms, −6,2 %)**. Les **299 934 extractions/reconstructions complètes** de la construction descendante disparaissent. Les **100 001 allocations** restent identiques. Aucune règle du compilateur n'est modifiée ou intégrée ; **B13 reste jaune**.

## Référence et provenance

Le [README officiel](../../../altbak.pub/README.md#rust) donne **11,631 ms pour RBTree compilé**, **36,070 ms pour le natif de la dernière colonne**, et **12,56 / 36,13 ms pour leurs suites**. RBTree représente 92,6 % du total compilé documenté. Le natif n'a pas été remesuré.

Le chemin annoncé `altbak.pub-rust/run/bak/rust/output/purescript` n'existe pas ici, et le worktree `altbak.pub-purust` n'avait plus sa sortie. L'audit précédent a donc généré du Rust avec le bundle Purust courant depuis le TAST existant dans `altbak.pub/run/bak/rust/output`, daté du 8 septembre. La source PureScript de RBTree est inchangée depuis le 29 août. Il s'agit d'une nouvelle génération backend, sans nouvelle compilation du frontend dans cette expérience.

La sortie complète est archivée dans `build/generated`, et le module RBTree de référence dans `build/RBTree-original.rs`. Son SHA-256 est `33ae6022e29f228a26e5d55a15edab26ac2d34ee17e0bc302b60918cc81fc7fb`. Les métadonnées et empreintes sont dans [metadata.json](metadata.json) et [runner-build.json](runner-build.json).

La nouvelle suite témoin mesure **13,975 ms**, pas les 12,56 ms historiques. Records vaut notamment **1,037 ms** dans ce témoin contre **0,382 ms** dans le README. L'origine de cet écart n'est pas étudiée ici ; toutes les variantes utilisent les mêmes autres modules. Le gain annoncé vient uniquement de la comparaison appariée, et ne doit pas être soustrait au total du README.

## Transformation et témoin

Le motif est `G = B(P = R(C = R(a,x,b),y,c),z,d)`. Le résultat est `G = R(C = B(a,x,b),y,P = B(c,z,d))` : la cellule racine G reste racine, C devient son fils gauche, P son fils droit.

Le prototype [rotation.rs](rotation.rs) est appelé dans les deux chemins uniques après insertion de l'enfant, seulement lorsque la garde existante a détecté une rotation. Il vérifie le motif gauche-gauche et l'unicité sans référence faible de P et C avant toute mutation. G est déjà sous l'emprunt exclusif du code généré. Trois permutations de pointeurs, une permutation de clés et trois écritures de couleur réalisent la rotation. Les sous-arbres `a,b,c,d` peuvent rester partagés. Après les gardes, aucune allocation, clone, destruction de pointeur ni opération susceptible de paniquer n'intervient.

Les autres motifs ou les cellules partagées/faiblement référencées conservent le code généré existant. Le workload descendant exerce **99 978 rotations gauche-gauche**, toutes prises par ce chemin. Cette expérience ne mesure pas le gain sur les trois autres orientations ni sur une charge fortement persistante.

Le Rust initial réutilise P pour le fils gauche et C pour le fils droit : son placement diffère du prototype. Le témoin [rebuild.rs](rebuild.rs) emploie **le même placement et les mêmes gardes que le prototype**, mais extrait et réinstalle intégralement les trois payloads. La comparaison `fields`/`rebuild` isole donc la stratégie sur les champs ; `fields`/`before` mesure le gain net du prototype, qui comprend aussi le changement de placement et le contournement du worker général.

## Mesures

Rust 1.96.0, profil existant **O1, debug=true, mimalloc**. Les trois variantes partagent le même lockfile et les mêmes versions d'allocateur. Chaque série comporte cinq blocs, ordre tournant puis inversé. Aucune compilation ni instrumentation de cette tâche ne tourne pendant les chronométrages. Aucun échantillon n'est écarté.

| Variante | Noyau isolé | RBTree dans la suite | Total suite |
| --- | ---: | ---: | ---: |
| Rust généré initial | 11,772 ms | 12,368 ms | 13,975 ms |
| Témoin, reconstructions complètes | 12,104 ms | 12,072 ms | 13,688 ms |
| Permutations de champs | 11,344 ms | 11,485 ms | 13,104 ms |

Le noyau comprend construction, profondeur et destruction : une chauffe puis quinze échantillons par processus, médiane des médianes de processus. Le runner complet garde sa chauffe et son meilleur de dix essais par benchmark. Ses totaux ci-dessus additionnent les médianes des quatorze benchmarks ; les médianes des totaux observés, conservées séparément, sont 13,973 / 13,693 / 13,111 ms.

Les permutations gagnent **dans chacun des cinq blocs**, face au Rust initial comme au témoin, dans le noyau et dans la suite. Le témoin seul est plus lent que l'initial dans le noyau mais plus rapide dans la suite : son effet dépend du contexte de compilation. Il ne faut pas lui attribuer un gain général.

| Bloc | Total initial | Total témoin | Total permutations |
| --- | ---: | ---: | ---: |
| 1 | 14,155 ms | 13,720 ms | 13,094 ms |
| 2 | 13,970 ms | 13,792 ms | 13,111 ms |
| 3 | 13,973 ms | 13,693 ms | 13,239 ms |
| 4 | 14,020 ms | 13,668 ms | 13,132 ms |
| 5 | 13,829 ms | 13,669 ms | 13,048 ms |

Dans la suite, la conservation des champs ajoute **0,587 ms de gain RBTree** par rapport au témoin. Le gain net face à l'initial est **0,883 ms RBTree**, et **0,871 ms sur le total** ; les autres modules inchangés varient légèrement. Les plages RBTree par processus sont **12,228–12,540 / 12,071–12,168 / 11,428–11,622 ms**.

Les quinze processus du runner complet valident chacun les **14 résultats attendus**. Sources et binaires sont hachés, les crates RBTree/App/root sont explicitement nettoyées puis reconstruites pour chaque variante, et les sources générées d'entrée restent identiques. [Mesures noyau](timings.json), [mesures suite](runner-results.json).

## Comptage séparé et correction

| Construction de 100 000 nœuds | Initial | Témoin | Permutations |
| --- | ---: | ---: | ---: |
| Rotations du prototype | 0 | 99 978 | 99 978 |
| Extractions complètes | 299 934 | 299 934 | **0** |
| Reconstructions complètes | 299 934 | 299 934 | **0** |
| Clones | 2 483 932 | 2 283 976 | 2 283 976 |
| `get_mut` réussis | 2 783 866 | 2 483 932 | 2 483 932 |
| Allocations | 100 001 | 100 001 | 100 001 |

Le comptage utilise un wrapper d'observation autour du véritable `std::rc::Rc`. Ce sont des opérations logiques exécutées, pas des instructions machine. Les **199 956 clones** et **299 934 tests d'unicité** évités sont communs au témoin et aux permutations ; ils ne doivent pas être attribués à la seule conservation des champs. Les bilans de propriétaires et de cellules passent. [Compteurs](counts.json).

Les trois noyaux passent les quatre orientations, les invariants rouge/noir, les clés comparées à `BTreeSet`, 200 versions persistantes, 512 insertions avec partage mixte et doublons, les bornes `i64`, les enfants indépendamment partagés, les références faibles et les libérations finales. Les deux prototypes passent en plus **54 combinaisons d'ownership et 6 motifs absents**, avec vérification des adresses, sous-arbres inchangés, refus sans mutation, snapshots conservés et disparition des références faibles après destruction. [Fixture ciblée](validation-extra.rs), [résultats](validation.json).

## Suite et reproduction

Le signal justifie maintenant une règle générale dérivée du TAST : reconnaître une reconstruction imbriquée comme permutation de champs de cellules uniques, prouver tous les accès avant mutation, puis garder le fallback existant lorsque la preuve ou l'unicité manque. Les noms et constructeurs de RBTree dans ce scratch ne doivent pas devenir des heuristiques propres au benchmark dans le compilateur. Une fixture TAST indépendante et une nouvelle mesure du Rust réellement émis seront nécessaires à l'intégration.

Depuis la racine d'`altbak.pub-purust`, exécuter séquentiellement :

```sh
python3 scratch/rust-rotation-fields-20260914/probe.py prepare
python3 scratch/rust-rotation-fields-20260914/probe.py validate
python3 scratch/rust-rotation-fields-20260914/probe.py count
python3 scratch/rust-rotation-fields-20260914/probe.py time
python3 scratch/rust-rotation-fields-20260914/runner.py build
python3 scratch/rust-rotation-fields-20260914/runner.py time
```

La première commande requiert l'archive `build/generated`. Sources générées, binaires et logs sont conservés localement sous `build/`, ignoré par Git. Les scripts, fixtures, mesures et ce rapport sont les seuls nouveaux fichiers hors de ce dossier. Aucun changement du compilateur, des sources de benchmark ni du README officiel.
