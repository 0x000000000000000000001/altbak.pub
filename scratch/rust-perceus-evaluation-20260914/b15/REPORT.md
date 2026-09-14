# B15 : spécialisation consommante de `depth`

L'hybride donne un **petit gain sur la suite complète** : RBTree **10,9945 → 10,8475 ms (−1,34 %, sept blocs favorables sur huit)** ; somme des médianes des quatorze tests **11,880 → 11,753 ms (−1,07 %)**. Sur vingt lectures d'une racine conservée, **17,0773 → 13,8418 ms (−18,95 %)**, mais presque tout ce gain provient du parcours emprunté : son témoin atteint déjà **13,8706 ms**. La consommation pure régresse sur cette charge partagée (**+4,11 %**).

Prototype isolé sur le Rust effectivement généré après l'intégration des permutations de champs. Le fichier de départ porte le SHA-256 `f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a`.

Seule la fonction `Test_RBTree_depth` change dans les copies de travail. Aucune source de Purust, sortie générée courante ou fonction de construction/insertion n'est modifiée. Les quatre fichiers complets utilisables par le runner sont `build/RBTree-before.rs`, `build/RBTree-borrowed.rs`, `build/RBTree-consuming.rs` et `build/RBTree-hybrid.rs`. L'ajout de la quatrième variante préserve les empreintes des trois premières.

## Variantes comparées

- **before** : fonction possédée émise par le compilateur, projections qui clonent les enfants puis destruction de la racine après le parcours.
- **borrowed** : témoin B14 identique au précédent prototype `rust-borrowed-depth-20260910` ; worker `&Tree`, propriétaire explicitement détruit avant le retour public.
- **consuming** : spécialisation B15 avec `Rc::unwrap_or_clone`. Une cellule unique est libérée et ses enfants déplacés dans les appels récursifs ; un propriétaire partagé clone le contenu et consomme seulement sa référence. L'ordre des appels reste gauche puis droite.
- **hybrid** : `Rc::try_unwrap` consomme les cellules uniques. En cas de partage, le même worker emprunté que B14 parcourt le sous-arbre sans clone de champ, puis l'owner retourné par l'échec est détruit avant le retour.

L'expérience antérieure de B14 avait retiré les clones sans démontrer un gain régulier dans le runner. Le contrôle B14 est donc inclus pour distinguer la suppression des clones de l'intérêt éventuel de fusionner parcours et consommation.

## Comptages logiques

L'instrumentation utilise le wrapper Rc d'un mot des expériences précédentes, avec contrôle des bilans globaux de propriétaires et de cellules. Il s'agit d'opérations logiques, pas d'instructions du binaire optimisé.

| Parcours puis destruction de 100 000 nœuds | before | borrowed | consuming | hybrid |
| --- | ---: | ---: | ---: | ---: |
| Clones Rc | 200 000 | 0 | 0 | 0 |
| Drops partagés | 300 000 | 100 000 | 0 | 100 000 |
| Derniers drops | 100 001 | 100 001 | 0 | 0 |
| `unwrap_or_clone` partagé | 0 | 0 | 100 000 | 0 |
| Consommations uniques (`unwrap_unique`) | 0 | 0 | 100 001 | 100 001 |
| Échecs `try_unwrap` (owner rendu) | 0 | 0 | 0 | 100 000 |
| Nouvelles allocations pendant le parcours | 0 | 0 | 0 | 0 |

La construction conserve exactement les mêmes compteurs et **100 001 allocations** dans les quatre variantes. Les 100 001 destructions existent dans toutes les variantes : `drop_last + unwrap_unique`. Pour ajouter `try_unwrap`, seule une copie du wrapper dans `build/tracked_rc.rs` change ; l'échec ne compte pas comme une libération car l'owner est rendu, et les bilans globaux continuent de passer.

Sur le parcours d'une racine conservée de 200 nœuds, `before` et `consuming` effectuent chacun **401 clones**, contre **1** pour `borrowed` et `hybrid` (le clone de l'appelant). L'hybride ajoute un seul échec `try_unwrap` sur cette racine. Le prototype B15 pur n'améliore donc pas ce chemin partagé, tandis que l'hybride conserve les opérations du parcours emprunté après le test initial.

Voir [counts.json](counts.json).

## Validations et périmètre

Les quatre variantes passent le parcours de 100 000 nœuds, les quatre rotations, les invariants RBTree, 128 cas de propriété du parent rouge, 200 versions conservées, 512 insertions avec partage mixte, les clés extrêmes, les références fortes/faibles indépendantes, les lectures répétées, les enfants partagés et leur libération finale. L'interface publique consomme toujours le propriétaire avant de retourner la profondeur.

Le prototype est limité à l'ADT natif fermé `Tree` contenant `Color` Copy, `Int` et `Rc<Tree>`. Le parcours n'exécute aucun callback ni destructeur de charge utile opaque. Aucune validité multithread n'est revendiquée.

[scope-checks.rs](scope-checks.rs) établit concrètement pourquoi cette restriction est nécessaire : sur un ADT doté de destructeurs observables, l'ordre passe de `[racine, gauche, droite]` à `[gauche, droite, racine]`. Un callback peut observer la disparition anticipée d'une référence faible à la racine. Lors d'une panique de ce callback, les deux variantes libèrent chaque charge utile exactement une fois mais dans des ordres différents. Ces cas sont des **contre-exemples à une généralisation sans preuve**, pas des cas déclarés équivalents.

Voir [validation.json](validation.json). Les compilations et validations n'ont utilisé aucune horloge.

## Mesures

Les mesures finales utilisent huit blocs : quatre rotations de l'ordre des variantes, puis quatre rotations inversées. Les compilations, validations et instrumentations étaient arrêtées avant le chronométrage. Le runner complet conserve son échauffement et son meilleur de dix ; les nombres ci-dessous sont les médianes des huit résultats par test. Le profil est O1/debug/mimalloc et **construction, parcours et destruction sont tous inclus**.

| Runner complet | before | borrowed | consuming | hybrid |
| --- | ---: | ---: | ---: | ---: |
| RBTree, ms | 10,9945 | 11,0190 | 10,8780 | **10,8475** |
| Évolution RBTree | — | +0,22 % | −1,06 % | **−1,34 %** |
| Blocs RBTree plus rapides que before | — | 5/8 | 6/8 | **7/8** |
| Somme des médianes des 14 tests, ms | 11,8800 | 11,9065 | 11,7675 | **11,7530** |

L'hybride gagne **0,147 ms** sur RBTree et **0,127 ms** sur la somme des médianes. Le gain reste faible et les plages RBTree se recouvrent : **10,874–11,346 ms** avant, **10,588–11,010 ms** avec l'hybride. Les quelques centièmes de milliseconde entre les variantes optimisées ne démontrent pas à eux seuls une supériorité robuste de leur mécanisme. [Résultats complets](../full-runner-results.json), [synthèse calculée](../summary.json).

Le second harness [shared_runner.py](shared_runner.py) mesure la construction de 100 000 nœuds, vingt appels à `depth` sur une racine conservée, puis la destruction du dernier propriétaire dans chaque échantillon. Chaque processus effectue un échauffement puis sept mesures et vérifie la somme 440 ; le tableau donne la médiane des huit médianes de processus. Les quatre noyaux exacts partagent le même harness et le même profil Cargo. `mimalloc` 0.1.52 et `libmimalloc-sys` 0.1.49 concordent avec le runner courant. [Provenance et smoke](shared-build.json).

| Construction + 20 lectures conservées + destruction | before | borrowed | consuming | hybrid |
| --- | ---: | ---: | ---: | ---: |
| Durée médiane, ms | 17,077250 | 13,870562 | 17,778417 | **13,841792** |
| Évolution | — | **−18,78 %** | +4,11 % | **−18,95 %** |
| Blocs plus rapides que before | — | 8/8 | 0/8 | 8/8 |

Le gain partagé est net, mais il vient presque entièrement de l'emprunt : l'écart médian entre `hybrid` et `borrowed` n'est que **0,029 ms**, leurs plages se recouvrent. Une spécialisation qui consommerait systématiquement le payload partagé serait donc une régression ; le repli emprunté est essentiel pour couvrir les deux charges. [Mesures partagées finales](shared-timings.json).

Les deux séries initiales restent disponibles séparément : [runner complet initial](../full-runner-results-initial.json) et [lectures partagées initiales](shared-timings-initial.json). Elles ont motivé l'ajout de l'hybride ; elles ne sont ni mélangées aux médianes finales ni présentées comme des mesures supplémentaires indépendantes de cette sélection.

## Ce que cette expérience établit

Le mécanisme de transfert possédé existe déjà : [le `foldl` de ListOps généré](/Users/0x1/Documents/htdocs/altbak.pub-purust/run/bak/rust/output/purust_output/Purs_Test_ListOps/src/lib.rs:115) appelle `Rc::unwrap_or_clone` et déplace les champs de `Cons`. Le prototype étend cette couverture à `depth`, avec un traitement emprunté des sous-arbres partagés. Il évalue une **extension ciblée de l'extraction et de la consommation des champs**, et non l'introduction d'un mécanisme jusqu'ici absent partout.

Ces résultats ne prouvent ni la nécessité ni le bénéfice d'une grande nouvelle passe Perceus complète. Ils justifient seulement d'examiner une règle limitée aux parcours totaux d'ADT natifs fermés, avec les contraintes d'effets et de destruction décrites ci-dessus. Le gain sur la suite est modeste ; le bénéfice des lectures partagées relève surtout de B14. Le prototype reste entièrement dans le scratch.

Reproduction de la préparation, depuis ce dossier, avec commandes séparées :

```sh
python3 probe.py prepare
python3 probe.py validate
python3 probe.py count
```
