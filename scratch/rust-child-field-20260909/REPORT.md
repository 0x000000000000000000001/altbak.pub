# B13 : modifier un enfant sans reconstruire son parent

Expérience du 9 septembre 2026. **Le prototype complet gagne 2,267 ms sur la suite, soit 12,1 %.** La conservation des champs à elle seule apporte **0,513 ms sur RBTree**, puis **0,660 ms** dans une seconde série de confirmation. Le reste vient du chemin spécialisé sans rotation et du retest d’unicité supprimé. Aucun changement du compilateur ou des sorties courantes n’est intégré.

## Référence et périmètre

Le [README officiel](../../../altbak.pub/README.md#rust), relu dans le checkout normal `altbak.pub`, donne **RBTree 18,480 ms**, **total 19,62 ms**, contre **36,070 ms** et **36,13 ms** dans la dernière colonne native. RBTree représente 94,2 % du total compilé. Ces références historiques restent distinctes des mesures appariées ci-dessous.

La source récente inspectée est [Purs_Test_RBTree/src/lib.rs](../../run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs), lignes 1663/1666 pour l’insertion. Le chemin unique extrait tout `Tree::T`, écrit temporairement `E`, transmet les champs à `balance`, puis reconstruit le parent. Un parent rouge ne tourne pas dans `balance` ; seul un enfant change après l’insertion récursive.

Le prototype vise uniquement ce cas : parent rouge, clé différente, cellule unique sans référence faible. Parents noirs, partagés, faiblement référencés et clés égales gardent le chemin existant. Il n’inclut pas la précédente optimisation expérimentale de `depth`.

## Trois variantes pour distinguer les gains

1. **before** : Rust généré inchangé.
2. **red_rebuild** : mêmes gardes que le candidat, emprunt mutable conservé, extraction de tout le payload puis reconstruction directe. Le passage par `balance` et le second `get_mut` disparaissent.
3. **red_field** : même garde et même emprunt mutable ; seul l’enfant choisi est déplacé et réinstallé. Couleur, clé et frère restent dans le parent.

Le candidat est en Rust sûr. Pour conserver un emplacement initialisé pendant l’appel récursif, il y place temporairement un clone du frère :

```rust
let child = std::mem::replace(left, right.clone());
*left = Test_RBTree_ins(key, child);
```

Le véritable enfant est consommé par l’appel, ce qui préserve sa possibilité de réutilisation. Cloner cet enfant avant l’appel aurait créé un partage artificiel. La copie temporaire du frère est relâchée lors de la réinstallation.

Le témoin utilise le même `&mut Tree` durant l’appel et la reconstruction. **La comparaison témoin → candidat isole donc le choix de garder les champs en place**, avec le coût de ce clone temporaire. Le gain avant → témoin ne doit pas être attribué à B13 seul.

## Compteurs séparés des chronométrages

Pour construire 100 000 nœuds :

| Opération logique | Avant | Témoin | Champs conservés |
| --- | ---: | ---: | ---: |
| Allocations de cellules | 100 001 | 100 001 | 100 001 |
| Extractions complètes `__purust_take` | 2 383 932 | 2 383 932 | 1 668 902 |
| Reconstructions complètes comptées | 2 383 932 | 2 383 932 | 1 668 902 |
| Tests d’unicité réussis | 4 867 864 | 4 152 834 | 4 152 834 |
| Clones | 3 060 100 | 3 060 100 | 3 775 130 |
| Relâchements sans destruction | 2 960 100 | 2 960 100 | 3 675 130 |
| Passages par la nouvelle branche rouge | 0 | 715 030 | 715 030 |

Les reconstructions du témoin additionnent le helper et ses 715 030 écritures directes. Le candidat supprime donc **715 030 extractions/reconstructions**, tout en ajoutant autant de paires clone/relâchement du frère. Les cellules créées et détruites restent identiques. Les compteurs sont des opérations instrumentées, pas un décompte d’instructions machine. [Données](counts.json).

L’assembleur O1 du noyau isolé confirme le changement : `red_rebuild.s` écrit le tag temporaire `E` avant l’appel (ligne 291), puis réécrit tag, couleur, clé et enfants (428–435). `red_field.s` modifie seulement l’emplacement gauche à l’offset 16 (283–305), ou droit à l’offset 24 (378–402), avec incrément puis décrément du frère temporaire. Les champs inchangés ne sont donc plus réécrits sur ce chemin. Ce contrôle porte sur les noyaux isolés, pas sur un désassemblage du runner complet ; il ne suffit pas à expliquer quantitativement les écarts temporels.

## Noyau isolé : première preuve

Rust `Rc` natif, O1/mimalloc, construction de 100 000 clés décroissantes, profondeur et destruction incluses. Cinq blocs de trois processus, ordre tournant/inversé ; une chauffe puis 15 mesures par processus. Compilation et instrumentation terminées avant le chronométrage.

| Médiane des 75 mesures | Temps |
| --- | ---: |
| Avant | 19,079 ms |
| Témoin | 17,691 ms |
| Champs conservés | 15,200 ms |

Le candidat gagne **3,880 ms (20,3 %) sur l’avant**, dont **2,491 ms par rapport au témoin**. Toutes les médianes par processus sont favorables. Ces chiffres appartiennent au noyau extrait et recompilé dans une crate unique ; ils ne prédisent pas directement le gain de la suite. [Mesures brutes](timings.json).

## Confirmation dans le runner complet

Une copie isolée de la crate principale, d’`App` et de `Test_RBTree` est construite sous `build/full-output`. Les autres crates utilisent leurs sources existantes inchangées. Le cache de compilation est copié par APFS avant utilisation. La source RBTree courante conserve son empreinte SHA-256 initiale.

Profil conservé : **O1, debug=true, mimalloc**. Le protocole du runner reste sa chauffe et son meilleur temps sur dix exécutions par benchmark. Cinq blocs de trois processus tournent/inversent l’ordre. Les 14 résultats attendus sont vérifiés à chaque lancement. Les totaux ci-dessous sont les sommes des médianes par benchmark, pas des temps muraux de processus.

| Runner | RBTree | Total de la suite |
| --- | ---: | ---: |
| Avant | 17,592 ms | 18,680 ms |
| Témoin | 15,844 ms | 16,938 ms |
| Champs conservés | 15,331 ms | 16,413 ms |

Le chemin complet gagne **2,261 ms sur RBTree (12,9 %)** et **2,267 ms sur la suite (12,1 %)**. La conservation des champs ajoute **0,513 ms sur RBTree (3,2 % face au témoin)** ; la différence des totaux vaut 0,525 ms. Les cinq blocs sont favorables au candidat face au témoin. [Résultats complets](runner-results.json).

L’écart avec le noyau isolé justifie une seconde série ciblée : cinq paires alternées **témoin / champs conservés**, avec les mêmes binaires et sans compilation. Résultat :

| Confirmation | RBTree | Total |
| --- | ---: | ---: |
| Témoin | 15,639 ms | 16,717 ms |
| Champs conservés | 14,979 ms | 16,055 ms |
| Gain | **0,660 ms (4,2 %)** | **0,662 ms** |

Les cinq paires sont favorables. Les deux séries ont leurs propres baselines et ne sont pas soustraites entre elles. **Le gain additionnel démontré pour garder les champs est donc d’environ 0,5 à 0,7 ms dans le runner**, plus petit que le gain du noyau isolé. [Confirmation](runner-confirmation.json).

## Correction et limites

Les trois variantes passent les contrôles natifs : quatre rotations, profondeur 22 pour 100 000 nœuds, 128 combinaisons de parent rouge et de partage fort/faible de racine/enfants, 200 anciennes versions, 512 insertions avec partage mixte et doublons, bornes entières, comparaison des clés avec `BTreeSet`, invariants rouge/noir, observation des anciennes valeurs via `Weak`, expiration après les derniers propriétaires. Les 32 parents rouges uniques contrôlés conservent leur adresse dans les trois variantes. Le comptage instrumenté vérifie aussi l’équilibre des propriétaires et destructions. [Validation](validation.json), [harness supplémentaire](validation-extra.rs).

Le scénario persistant reste correct, mais les tests d’unicité refusés passent de 1 303 à 1 683 dans les deux variantes expérimentales : la garde peut ajouter un essai avant le fallback. Ses performances ne sont pas chronométrées ici. Le gain porte sur le benchmark aux insertions majoritairement uniques.

Ce prototype spécialise un noyau fermé sans callback. Il ne constitue pas une règle générale pour des appels opaques : aucun callback ni unwind injecté n’est testé. Une intégration demanderait une reconnaissance structurale à partir du TAST et des usages, une preuve de branche sans rotation, la conservation de l’ordre d’évaluation et des libérations, puis des fixtures indépendantes du nom RBTree. Les fichiers du compilateur, sa matrice et ses tests de génération ne sont pas modifiés par cette expérience.

## Reproduction

Depuis `altbak.pub-purust`, séquentiellement :

```sh
python3 scratch/rust-child-field-20260909/probe.py prepare
python3 scratch/rust-child-field-20260909/probe.py count
python3 scratch/rust-child-field-20260909/probe.py validate
python3 scratch/rust-child-field-20260909/probe.py time
python3 scratch/rust-child-field-20260909/runner.py build
python3 scratch/rust-child-field-20260909/runner.py time
python3 scratch/rust-child-field-20260909/runner.py time --variants red_rebuild red_field --output runner-confirmation.json
```

[Métadonnées](metadata.json) : source, empreinte SHA-256, révisions, version de rustc, options et binaires du noyau. Les JSON de runner conservent les empreintes de leurs binaires et les mesures brutes. Les snapshots, binaires, logs et assembleurs sont dans `build/`, ignoré par Git.
