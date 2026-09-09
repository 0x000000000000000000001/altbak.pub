# Perceus : premier comptage des références dans RBTree

Le premier diagnostic trouve **2 960 100 paires clone/relâchement temporaires** dans les tests de motifs de la construction de RBTree. Un prototype les supprime par emprunt, avec les mêmes résultats, allocations et décisions d'unicité. **Il ne gagne pourtant pas de temps : 19,737 ms → 20,563 ms (+4,2 %)** dans la série isolée. Ce prototype n'est pas intégré au générateur.

La prochaine piste est la propagation d'une unicité déjà établie : l'instrumentation observe **4 967 864 appels réussis à `Rc::get_mut`**, dont **2 483 932** sur des cellules temporairement vidées après extraction. Leur coût machine et le gain de leur suppression restent à vérifier.

## Ce qui a été exécuté

[probe.py](probe.py) extrait les enums, `ins`, `insert`, `balance`, les helpers de réutilisation, `buildTree` et `depth` du [Rust généré](../../run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs). Les fonctions extraites gardent leurs numéros de ligne d'origine. Le runner et ses sorties générées restent inchangés.

Pour le comptage seulement, les pointeurs passent par [tracked_rc.rs](tracked_rc.rs), un wrapper d'un mot autour du vrai `std::rc::Rc`. `Tree` reste à 32 octets. Le wrapper observe les créations, clones, relâchements, tests d'unicité et consommations par `unwrap_or_clone`, y compris les clones de champs du chemin partagé. Le dernier relâchement laisse Rust détruire normalement le contenu. Les bilans globaux des propriétaires et des destructions sont vérifiés.

Ces nombres sont des **opérations logiques exécutées avec instrumentation**, pas un décompte des instructions du binaire optimisé. Le wrapper peut empêcher des optimisations ; ses tables de comptage ont leurs propres allocations, exclues des nombres de cellules. Il n'est jamais utilisé pour chronométrer.

## Construction et parcours uniques : 100 000 insertions

| Phase / opération | Rust actuel | Prototype emprunté |
| --- | ---: | ---: |
| Construction : cellules créées | 100 001 | 100 001 |
| Construction : clones de pointeur | 3 060 100 | 100 000 |
| Construction : relâchements sans destruction du contenu | 2 960 100 | 0 |
| Construction : tests d'unicité réussis | 4 967 864 | 4 967 864 |
| Construction : tests refusés | 0 | 0 |
| Parcours `depth` : clones de pointeur | 200 000 | 200 000 |
| Parcours et destruction : relâchements sans destruction du contenu | 300 000 | 300 000 |
| Parcours et destruction : derniers relâchements | 100 001 | 100 001 |

Le résultat final reste une profondeur de **22**. Les 100 000 clones restant pendant la construction sont ceux du constructeur vide partagé. Le parcours consomme la racine et inclut la destruction complète de l'arbre.

Les sites les plus fréquents sont les lignes **1741, 1742, 1807 et 1808** du Rust généré : un enfant est cloné, utilisé comme récepteur de `as_ref()` pour un test de constructeur ou une projection, puis immédiatement relâché. Le prototype remplace le retour possédé de cette projection par un emprunt du champ, puisque le parent local reste propriétaire pendant la lecture. Il applique cette seule forme à **210 sites statiques**, sans modifier les autres extractions de champs ni les rotations.

## Chemins partagés

Le scénario persistant construit une permutation de 200 clés en gardant toutes les anciennes racines. L'ordre exact des clés et les invariants sont vérifiés pour chaque version, y compris après libération de la dernière racine.

| Construction persistante | Rust actuel | Prototype |
| --- | ---: | ---: |
| Clones de pointeur | 5 808 | 3 006 |
| Relâchements temporaires sans destruction | 2 802 | 0 |
| Tests d'unicité réussis | 1 060 | 1 060 |
| Tests refusés pour partage | 1 303 | 1 303 |
| Consommations `unwrap_or_clone` partagées | 1 303 | 1 303 |
| Cellules créées puis détruites | 1 504 | 1 504 |

Les quatre rotations, les invariants rouge/noir, les références faibles et l'équilibre des durées de vie passent dans les deux variantes. Les relâchements internes de la référence consommée par `unwrap_or_clone` sont comptés séparément, pas une seconde fois comme des `drop` du wrapper. [Événements, phases et sites complets](counts.json).

## Fixture TAST minimale

[PerceusProbe.purs](PerceusProbe.purs) définit un autre ADT et des motifs imbriqués, sans reprendre le benchmark. [fixture.py](fixture.py) le compile avec le fork TAST, contrôle ses deux `dataDecls`, puis lance le vrai bundle Purust. Le prototype ne modifie que le Rust fraîchement émis.

Les 8 sites de projection empruntée suppriment **10 paires** dans les cas uniques et **4 paires** dans chacun des scénarios partagé et faiblement référencé. Les trois résultats possibles, la priorité des branches, les appels répétés, le contenu des anciennes valeurs et les durées de vie des racines/enfants sont vérifiés. Allocations et destructions restent identiques. [Résultats de la fixture](fixture.json).

## Chronométrage sans instrumentation

[time.py](time.py) compile deux noyaux utilisant le vrai `std::rc::Rc`, avec **O1 et mimalloc**. Chaque mesure inclut construction, profondeur et destruction. Trois paires de processus alternent l'ordre ; chaque processus effectue un échauffement puis 15 mesures. La série retenue commence après la fin du comptage, de la fixture et des compilations de diagnostic ; aucun autre processus de cette tâche n'est lancé pendant les mesures.

| Médiane par processus | Rust actuel | Prototype |
| --- | ---: | ---: |
| Paire 1 | 19,632 ms | 21,058 ms |
| Paire 2 | 19,737 ms | 20,530 ms |
| Paire 3 | 19,824 ms | 20,460 ms |
| **Médiane des 45 mesures** | **19,737 ms** | **20,563 ms** |

Un premier passage exploratoire, lancé avec le comptage/compilation en concurrence, allait dans le même sens ; il est exclu de cette série. Les médianes exploratoires et toutes les mesures de la série retenue sont dans [timings.json](timings.json).

Le code machine confirme que la transformation retire des opérations de comptage dans le motif inspecté : [extrait assembleur](assembly-excerpt.txt). L'IR optimisé du worker de réutilisation contient aussi moins de sites d'instructions : [résumé statique](assembly-summary.json). Cela ne suffit pas à expliquer ni prédire le temps total ; aucune cause précise du ralentissement n'est établie ici.

Le [README officiel](../../../altbak.pub/README.md#rust), relu pour cette étape, indique **RBTree 18,985 ms**, **total 21,82 ms**, contre **36,070 ms** et **36,13 ms** dans la dernière colonne native. Ces repères restent distincts des noyaux isolés. Aucune nouvelle mesure du runner complet ni amélioration de son total n'est revendiquée.

## Suite proposée

Les **2 483 932** tests d'unicité sur des cellules au tag temporaire `E` se répartissent entre le helper de reconstruction, ligne **64** (**2 383 932**), et la recoloration, ligne **1682** (**100 000**). Ils suivent les extractions autorisées par un premier `get_mut` réussi. Le prochain bébé consiste à vérifier un de ces retests dans le code machine, puis mesurer une variante conservant la preuve d'unicité. La preuve doit rester valide à travers les usages intermédiaires ; les cas partagés et faiblement référencés gardent leur comportement.

Cette étape fournit un diagnostic et une fixture, sans nouvelle passe Perceus, changement de runtime, modification du générateur ou nouveau classement de la matrice. Elle évite d'intégrer une réduction du comptage qui ralentit le noyau mesuré.

## Reproduction

Depuis `altbak.pub-purust`, exécuter les commandes séparément, dans cet ordre :

```sh
python3 scratch/rust-perceus-counts-20260909/probe.py
python3 scratch/rust-perceus-counts-20260909/fixture.py
EMIT_IR=1 python3 scratch/rust-perceus-counts-20260909/time.py
python3 scratch/rust-perceus-counts-20260909/summarize_ir.py
```

Les scripts lisent les sorties générées actuelles ; leurs empreintes et les révisions de départ sont dans [metadata.json](metadata.json). `build/` contient les sources extraites, les binaires et les diagnostics volumineux ignorés par Git. [validation.json](validation.json) identifie les scripts et résultats conservés.
