# B13/B14 : décider après le calcul de l’enfant

Intégration du 10 septembre 2026 dans Purust. Cinq paires alternées de runners complets réellement générés donnent **16,590 → 12,640 ms**, soit **3,950 ms gagnées (−23,8 %)**. Les cinq paires sont favorables. RBTree donne **15,498 → 11,528 ms (−3,970 ms, −25,6 %)**.

## Règle intégrée

[ChildBranches](../../../purust/purust/src/Purust/ChildBranches.purs) analyse les branches d’un helper local à partir du TAST. Il prouve les chemins où le helper retourne le même constructeur avec exactement les arguments d’origine, dans leur permutation connue. Les tests admis portent sur les tags de paramètres ou de champs typés ; toute projection doit être précédée d’un test du constructeur qui la contient. Les conjonctions, disjonctions, négations et branches conservent leur ordre et leur court-circuit.

[ChildBranchPrinter](../../../purust/purust/src/Purust/ChildBranchPrinter.purs) émet une garde qui emprunte les arguments. Le chemin ajouté à [CodeGen](../../../purust/purust/src/Purust/CodeGen.purs) conserve l’emprunt mutable d’un parent unique, détache un seul enfant, le calcule puis le réinstalle. La garde examine alors le résultat : si elle est vraie, le parent est retourné directement ; sinon, le worker existant reçoit les champs, y compris l’enfant déjà calculé. La récursion n’est jamais répétée. Le précédent raccourci dont la garde pouvait être évaluée avant l’appel reste prioritaire.

Les critères portent sur les types, layouts, usages et branches : **aucun nom de benchmark, de fonction de benchmark ou de constructeur particulier n’entre dans la règle**. La fixture indépendante emploie `Mode` et `Node`, des changements de tags et des branches qui permutent les enfants.

Le périmètre reste conservateur : parent consommé unique, layout récursif contenant seulement le même arbre et des champs natifs Copy, un enfant changé, appels locaux saturés et graphe fermé. Les callbacks, appels opaques/FFI, conversions, changements multiples et parents partagés/faibles conservent les chemins existants. Les branches complexes, dont les rotations, utilisent encore la reconstruction complète. Les noms auxiliaires sont réservés contre les déclarations utilisateur et les autres helpers. Les prédicats toujours faux ne sont pas émis. **B13 et B14 restent jaunes.**

## Mesures

Même TAST frais, mêmes dépendances et directives, même runtime, même `Cargo.lock`, profil release **O1**, allocateur **mimalloc**. Le compilateur de référence est le bundle conservé avant les modifications, révision `d7aa70124e7174079cbb2cf866051e174b7a7129`. Le runner après provient de la sortie finale de Purust. Tous les fichiers Rust sont identiques entre les deux sorties sauf `Purs_Test_RBTree/src/lib.rs`.

Chaque processus conserve le protocole du benchmark : sortie/warm-up puis meilleur de dix. Les quatorze résultats sont contrôlés à chaque passage. Le tableau donne les médianes de cinq processus par variante ; le total est la somme des médianes par benchmark. Aucun test, compilation ou comptage ne tourne pendant les mesures.

| Mesure | Avant | Après | Différence |
| --- | ---: | ---: | ---: |
| RBTree | 15,498 ms | 11,528 ms | −3,970 ms |
| Records | 0,616 ms | 0,629 ms | +0,013 ms |
| Church | 0,157 ms | 0,160 ms | +0,003 ms |
| Prime Sieve | 0,175 ms | 0,180 ms | +0,005 ms |
| **Suite complète** | **16,590 ms** | **12,640 ms** | **−3,950 ms** |

Les totaux par processus s’étendent de **16,469 à 16,893 ms** avant et de **12,400 à 12,787 ms** après. Les gains appariés vont de **3,749 à 4,224 ms**. Les petits écarts des autres cas sont inclus dans le total ; leurs sources n’ont pas changé.

Le [README officiel](../../../altbak.pub/README.md#rust), relu le 10 septembre, donne **15,956 ms RBTree et 17,12 ms au total** pour le compilé ; sa dernière colonne native donne **36,070 ms et 36,13 ms**. Le compilé était donc déjà plus rapide que ce natif sur RBTree. Ces chiffres historiques contextualisent la mesure ; le gain de cette intégration est calculé sur la paire fraîche ci-dessus. Ni le natif ni le README n’ont été remesurés/modifiés ici. Les résultats du prototype précédent et des intégrations antérieures ne s’ajoutent pas à ce gain.

Une première série a été invalidée : la copie du cache Cargo avait donné deux binaires identiques malgré des sources différentes. Le binaire avant a ensuite été **recompilé entièrement dans son propre workspace**, et les empreintes des deux binaires sont distinctes. Seule cette série corrigée figure dans [runner-results.json](runner-results.json). Le script impose désormais cette reconstruction et vérifie les empreintes.

## Travail exécuté

Les compteurs sont ajoutés séparément au noyau réellement généré. Ils observent des opérations logiques du runtime, pas des instructions machine.

| Construction de 100 000 nœuds | Avant | Après |
| --- | ---: | ---: |
| Extractions / reconstructions complètes | 1 668 902 | **299 934** |
| Clones de pointeurs | 3 775 130 | **2 483 932** |
| Relâchements partagés | 3 675 130 | **2 383 932** |
| Tests d’unicité réussis | 4 152 834 | **2 783 866** |
| Allocations / libérations finales | 100 001 | **100 001** |
| Appels `ins` | 2 283 976 | **2 283 976** |

La nouvelle garde est vraie **1 368 968 fois** et conserve le fallback de rotation **99 978 fois**. Le chemin précédent couvre toujours **715 030** parents. Les bilans de propriétaires et de destructions sont équilibrés.

La correction des versions persistantes est vérifiée, **leur accélération n’est pas établie**. Dans le petit comptage conservant 200 versions, les tests d’unicité refusés pour partage passent de 1 683 à 2 606 : le nouveau chemin peut tester l’unicité avant de rejoindre le fallback existant. Les allocations et clones de cette phase restent identiques. Le gain annoncé concerne le runner mesuré, dont RBTree consomme les versions successives.

## Validation

- Build du compilateur réussi ; **26 tests de génération et 15 tests TAST** passent. Les avertissements préexistants du compilateur restent hors de cette modification.
- `bin/rust/run -c` puis régénérations finales : les **14 sorties** sont correctes. La dernière régénération et les sources instrumentées ont les mêmes empreintes.
- [Fixture TAST](../../../purust/purust/tests/tast/post-call-child-reuse.mjs) : 320 cas de partage, deux directions, tags modifiés pendant l’appel, branches simples/complexes, un seul appel, récursion, persistance et chemins exclus.
- [Analyse des branches](../../../purust/purust/tests/codegen/child-branches.mjs) : chemins typés, aliases, projections dominées, ordre des branches, négation/conjonction/disjonction et refus des expressions inconnues.
- [Paniques sur le chemin accepté](../../../purust/purust/tests/codegen/post-child-panic.mjs) : enfants fermés qui échouent, parents uniques/partagés/faibles, durées de vie équilibrées. [Collision de noms](../../../purust/purust/tests/codegen/post-child-helpers.mjs) : compilation et exécution Rust.
- Validation du noyau natif réellement généré avant/après : quatre rotations, 100 000 clés/profondeur 22, 256 combinaisons de partage fort/faible, 200 anciennes versions, bornes `i64`, 512 insertions avec partage mixte et doublons.

## Reproduction et artefacts

[compare.py](compare.py) construit et compare les runners ; [count.py](count.py) compte les opérations et [validate.py](validate.py) vérifie les noyaux natifs. Les révisions, empreintes du TAST, des sources et des binaires sont dans [metadata.json](metadata.json), [counts.json](counts.json), [validation.json](validation.json) et [runner-results.json](runner-results.json). Les snapshots, logs et binaires restent dans `build/`, ignoré par Git.

Depuis la racine d’`altbak.pub-purust`, après le build propre avec le compilateur final :

```sh
# Restaurer seulement une copie du bundle de référence dans le scratch.
mkdir -p scratch/rust-post-call-child-integration-20260910/build
git -C ../purust/purust show d7aa70124e7174079cbb2cf866051e174b7a7129:bin/purust.js > scratch/rust-post-call-child-integration-20260910/build/purust-before.mjs
node --expose-gc --stack-size=65536 --max-old-space-size=16384 scratch/rust-post-call-child-integration-20260910/build/purust-before.mjs --main App --source run/bak/rust/output --out scratch/rust-post-call-child-integration-20260910/build/before-output
python3 scratch/rust-post-call-child-integration-20260910/compare.py build
python3 scratch/rust-post-call-child-integration-20260910/compare.py time
python3 scratch/rust-post-call-child-integration-20260910/count.py
python3 scratch/rust-post-call-child-integration-20260910/validate.py
```

Les checkouts normaux `altbak.pub` et `purescript-backend-optimizer` n’ont pas été modifiés.
