# Spécialisation d'un champ scalaire : intégrée dans Purust

9 septembre 2026. La recoloration d'un nœud unique ne touche désormais que son champ couleur. Les enfants, la clé et le constructeur restent en place. Cette règle générale de réutilisation est intégrée, après un prototype favorable et une confirmation dans la suite complète.

## Mesures

| Comparaison | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Prototype RBTree, 5 paires | 19,695 ms | 19,402 ms | −1,5 % |
| Prototype RBTree, confirmation 7 paires | 19,761 ms | 19,410 ms | −1,8 % |
| **RBTree dans le runner intégré, 5 paires** | **18,237 ms** | **17,644 ms** | **−3,3 %** |
| **Total du runner intégré** | **20,977 ms** | **20,354 ms** | **−3,0 %** |

Les prototypes utilisent 15 échantillons après une chauffe par processus ; construction de 100 000 nœuds, profondeur et destruction incluses. La première série a quatre paires favorables sur cinq, la confirmation sept sur sept. O1, mimalloc, véritables `std::rc::Rc`, sans instrumentation pendant les mesures. Les précédentes transformations d'emprunts ne sont pas combinées à celle-ci.

Le runner complet conserve sa chauffe et son meilleur temps sur dix par benchmark. Nous prenons la médiane de cinq processus, en alternant l'ordre avant/après ; les totaux sont les sommes des médianes par benchmark. RBTree est plus rapide dans les cinq paires. Les temps avant vont de 17,811 à 18,321 ms ; après, de 17,567 à 18,021 ms. Les petites variations des autres benchmarks ne sont pas attribuées à la règle.

Référence historique relue dans le [README normal](../../../altbak.pub/README.md#rust) : RBTree Rust compilé **18,985 ms**, natif optimisé **36,070 ms** ; totaux **21,82 et 36,13 ms**. Les gains ci-dessus proviennent des paires de cette expérience, pas d'une soustraction entre un run actuel et cette baseline. Le README normal reste inchangé.

## Règle et limites

[ReuseFields.purs](../../../purust/purust/src/Purust/ReuseFields.purs) s'appuie sur les types/layouts TAST et les preuves de dernière utilisation existantes. Après réécriture des projections de propriété, chaque champ inchangé doit désigner son propre emplacement original, sans conversion de représentation. Un seul champ peut changer : entier, nombre, booléen, caractère ou enum `Copy` sans payload identifié par ses déclarations TAST.

Le remplacement est un atome scalaire : littéral, local ou constructeur sans champ. Les appels, calculs, closures, modifications de plusieurs champs et conversions ne déclenchent pas la règle. L'expression scalaire est évaluée avant l'emprunt ; le chemin unique affecte son seul emplacement avec `Rc::get_mut`. Le chemin partagé/faible conserve la reconstruction existante. Aucun `unsafe` ajouté au générateur. Aucun nom de benchmark n'intervient dans la décision.

La compilation propre modifie deux crates générées : les deux recolorations de `Test.RBTree` et les quatre setters scalaires de `Data.Time`. Cette dernière application couvre aussi un ADT sans constructeur vide. Les autres empreintes de sources sont inchangées dans [source-comparison.json](source-comparison.json).

## Travail supprimé et vérifications

La construction RBTree supprime **100 000 extractions/reconstructions complètes et 100 000 retests d'unicité**. Les tests réussis passent de **4 967 864 à 4 867 864**. Les **100 001 allocations/libérations**, **3 060 100 clones** et **2 960 100 relâchements temporaires** restent identiques. Les refus pour partage ou référence faible et les bilans de durée de vie sont conservés. Le comptage a été refait sur le Rust réellement émis après intégration : [counts-integrated.json](counts-integrated.json).

Dans le prototype optimisé, le harnais LLVM passe de 43 à 27 sites de chargement, de 54 à 43 écritures et de 57 à 40 branchements. Ce sont des sites statiques, pas des instructions exécutées par insertion : [assembly-summary.json](assembly-summary.json). Le [diff Rust](RBTree.diff) montre directement le remplacement de l'extraction/réécriture par l'affectation de couleur.

Validation terminée :

- Build Purust et bundle réussis, sans avertissement de compilation PureScript.
- **19 tests de génération et 11 tests TAST réussis**, fork `purs` sélectionné explicitement.
- Nouvelle fixture TAST : enums, scalaires, ADT sans constructeur vide, racines uniques/partagées/faibles, enfants aliasés, adresses et durées de vie, versions conservées, permutations, callbacks, plusieurs champs et insertion du parent comme enfant. Les cas hors périmètre gardent leur génération antérieure.
- `bin/rust/run -c` réussi dans le worktree ; **14 résultats corrects**, puis à nouveau dans les dix processus mesurés.
- Rust RBTree réellement régénéré vérifié avec et sans compteurs : quatre rotations, 100 000 nœuds, 200 versions conservées, 512 insertions avec doublons et partages intermittents, faibles et libérations finales.

## Reproduction et artefacts

[metadata.json](metadata.json) contient les révisions avant changement, empreintes des sources/harnesses et informations de compilation. L'exécutable avant a été vérifié par Cargo avant la compilation propre ; les deux binaires utilisent `libmimalloc-sys 0.1.49`. Les sources, exécutables et diagnostics volumineux restent dans `build/`, ignoré par Git. Le prototype réutilise les harnesses des deux expériences précédentes et sa copie `build/RBTree-before.rs` ; si elle manque, régénérer la baseline avec la révision Purust enregistrée avant de relancer le prototype.

```sh
python3 probe.py count
python3 probe.py validate
python3 probe.py time --assembly --pairs 5 --output timings.json
python3 probe.py time --pairs 7 --reuse-binaries --output timings-confirmation.json
python3 probe.py count --integrated
python3 probe.py validate --integrated
python3 measure-runner.py
```

Exécuter les mesures après la fin de toutes les compilations et instrumentations. Les binaires complets avant/après doivent être conservés dans `build/runner-before` et `build/runner-after`. [Mesures complètes](runner-results.json), [prototype initial](timings.json), [confirmation](timings-confirmation.json), [vérification native après intégration](validation-integrated.json).

Les modifications portent sur Purust, sa fixture et le scratch du worktree. Les checkouts normaux d'altbak.pub et de PBO sont préservés. La prochaine extension reste à mesurer séparément : un enfant modifié par un appel, avec une preuve de propriété valide pendant cet appel.
