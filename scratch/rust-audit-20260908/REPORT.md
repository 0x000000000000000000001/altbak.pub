# Audit du Rust généré — 8 septembre 2026

Deux gains sont confirmés par des expériences isolées : **26 % de temps en moins sur RBTree** en combinant couleurs en valeur et emprunts pour les tests de motifs ; **64 % de temps en moins sur LazyEvaluation** en évitant la reconstruction de `Unit`. LazyEvaluation représente actuellement deux tiers du total et constitue donc la priorité pour le temps global. Pour commencer spécifiquement par RBTree, les emprunts sont la modification la plus locale.

Ces prototypes sont conservés dans ce dossier. Aucune optimisation n'a été intégrée au générateur, aux sources des benchmarks ou au projet Rust généré. Les gains mesurés sur les noyaux extraits restent à confirmer après intégration avec `bin/rust/run -c`.

## Mesures actuelles et baselines officielles

Référence historique : `/Users/0x1/Documents/htdocs/altbak.pub/README.md`, tableau Rust, lignes 132–150. Le binaire existant du worktree a été exécuté trois fois, successivement. Chaque exécution du runner retient le meilleur de dix essais ; la colonne actuelle donne la médiane des trois résultats obtenus. Les 14 résultats fonctionnels sont corrects à chaque exécution.

| Benchmark | README officiel | Actuel | Part du total actuel |
|---|---:|---:|---:|
| LazyEvaluation | 319,209 ms | 233,546 ms | 66,5 % |
| RBTree | 67,125 ms | 61,729 ms | 17,6 % |
| Polymorphism | 38,663 ms | 42,605 ms | 12,1 % |
| Church | 24,331 ms | 11,723 ms | 3,3 % |
| Total | 452,29 ms | 351,312 ms | 100 % |

Le total actuel est la somme des médianes par benchmark. Les chiffres historiques donnent un repère ; les différences avec aujourd'hui ne constituent pas une expérience contrôlée attribuable à un changement particulier. Les logs complets et les valeurs des autres benchmarks figurent dans `full-run-{1,2,3}.log` et `results-full.json`.

## RBTree : ce qui coûte et ce qui fonctionne déjà

Le code chaud emploie déjà des enums `Tree` et `Color`, des clés `i64`, des appels directs et une boucle pour `buildTree`. Il ne passe pas par `Value` ou `PerceusPtr` dans `balance`, `ins`, `insert`, `buildTree` et `depth`. Le TAST fournit bien la structure des ADT et le typage ; chercher à monomorphiser cet algorithme déjà monomorphe ne répondrait pas aux coûts observés.

En revanche, le champ couleur est un `Rc<Color>` : une couleur sans charge utile reçoit une allocation et un compteur de références. Les tests et extractions de motifs clonent aussi fréquemment le pointeur parent avant de lire son contenu. Les quatre rotations du source deviennent environ 790 lignes dans `balance`, avec 379 occurrences statiques de `.clone()`. Ce comptage décrit le code émis, pas le nombre d'opérations exécutées par insertion.

Références dans `output/purust_output/Purs_Test_RBTree/src/lib.rs` : enums ligne 44, `balance` ligne 125, `ins` ligne 915, `insert` ligne 936, `buildTree` ligne 950. Source PureScript : `src/Test/RBTree.purs`, lignes 13–20.

### Comparaison contrôlée, même profil et même allocateur

Chaque variante reprend le noyau généré et conserve l'arbre persistant en `Rc<Tree>`, les 100 000 insertions descendantes, le parcours de profondeur et la destruction. Compilation `rustc -C opt-level=1`, allocateur mimalloc identique à celui du runner. Trois processus par variante, ordre alterné, un échauffement et cinq mesures par processus : **15 mesures par variante**, chronométrées avec `Instant`.

| Variante | Médiane | Réduction du temps |
|---|---:|---:|
| Noyau généré actuel | 55,928 ms | référence |
| Couleur en valeur, enum `Copy` | 46,493 ms | 16,9 % |
| Emprunter le parent lors des tests/extractions | 47,642 ms | 14,8 % |
| Couleur en valeur + emprunts | **41,349 ms** | **26,1 %** |
| Quatre rotations factorisées, couleurs inchangées | 48,814 ms | 12,7 % |
| Rotations factorisées + couleur en valeur | 42,546 ms | 23,9 % |

La transformation d'emprunt remplace seulement `(parent.clone()).as_ref()` par `parent.as_ref()` lorsque le parent est une variable locale. Elle conserve les clones des enfants nécessaires au partage structurel. La version factorisée est une réécriture expérimentale des quatre rotations ; son gain recouvre une partie du gain des emprunts. **Ces pourcentages ne s'additionnent pas.**

La référence extraite à 55,9 ms diffère du benchmark complet à 61,7 ms : contexte d'allocation, programme environnant et compilation ne sont pas identiques. Le résultat démontré est la comparaison entre variantes isolées, pas un nouveau score officiel du runner.

Les vérifications hors chronométrage couvrent 1 000 insertions mélangées, l'ordre des clés, les hauteurs noires identiques, l'absence de deux nœuds rouges consécutifs, un doublon et la conservation d'une ancienne version après insertion. Tous les essais à 100 000 insertions renvoient la profondeur attendue, 22. Cela donne une validation ciblée, pas une preuve générale de correction du futur générateur.

### Allocations réellement comptées

Mesure distincte instrumentée, avec le même algorithme ; ses temps ne sont pas utilisés pour les comparaisons de vitesse.

| Variante | Appels d'allocation | Octets demandés cumulés |
|---|---:|---:|
| Généré | 3 283 867 | 145 627 200 |
| Couleur en valeur | 2 783 933 | 133 628 784 |
| Couleur en valeur + réutilisation de la racine | 2 683 933 | 128 828 784 |

Le passage de `Rc<Color>` à `Color` élimine **499 934 allocations**, soit 15,2 % des appels. La taille de `Tree` reste 32 octets dans cette expérience : le gain vient de la suppression des allocations de couleur et de leur gestion, sans réduction observée de la taille du nœud.

Un prototype utilisant `Rc::make_mut` pour recolorer la racine économise encore 100 000 allocations. Son gain temporel supplémentaire est modeste dans la série exploratoire à `opt-level=3` : 47,0 à 45,3 ms par rapport à la couleur en valeur seule. Une réutilisation plus large des nœuds pourrait aller plus loin, mais demanderait une gestion de l'unicité et du partage ; ce gain n'a pas été mesuré.

### Baby steps RBTree

1. **Emprunts pour les lectures de motifs.** Commencer par `OpIsTag` sur une variable locale puis les extractions de champs. Vérifier les branches imbriquées, les déplacements ultérieurs et les enfants encore partagés. Points à examiner : `purust/purust/src/Purust/CodeGen.purs`, lignes 1179 et 1274. Rejouer les régressions du générateur puis `bin/rust/run -c` avant d'élargir.
2. **Enums sans charge utile en valeur.** Déterminer l'éligibilité depuis `dataDecls`, en commençant par le cas de `Color`. Rendre cohérents déclarations, signatures, constructeurs, tests, champs et conversions aux frontières dynamiques. Émettre `Copy` pour ce cas ; ne pas coder une exception propre au benchmark. Points : génération des enums ligne 48 et constructeurs vers 1504–1516 dans `CodeGen.purs`. Vérifier les échanges entre modules puis mesurer séparément et avec l'étape 1.
3. **Factoriser les motifs imbriqués.** Capturer une déconstruction empruntée et réutiliser ses champs. Déterminer d'abord à quel endroit les tests sont dupliqués entre PBO et Purust. Un petit cas à deux niveaux et deux branches suffit pour démarrer ; comparer ensuite les quatre rotations et le Rust émis.
4. **Réutiliser une racine unique.** Tester une recoloration sous `Rc::make_mut`, avec une ancienne version encore vivante pour couvrir la copie en cas de partage. Élargir seulement après mesure.

Le README donne 59,784 ms pour la FFI Rust fonctionnelle et 16,700 ms pour la version manuellement optimisée. Cette dernière emploie une arène préallouée et des indices ; la FFI fonctionnelle utilise des `Box`, des déplacements et certains clones profonds. Ces implémentations ont des coûts de partage et de durée de vie différents : **16,7 ms est un repère, pas un gain promis par une simple modification de représentation.**

## LazyEvaluation : le plus gros gain confirmé sur l'ensemble

`output/purust_output/Purs_Data_Unit/src/lib.rs:13` construit chaque `unit` comme un `Value::Record_a(PerceusPtr::new(...))`. Le record générique, défini dans `purust_core/src/lib.rs:5796`, contient les nombreux champs optionnels réunis par le runtime. Taille mesurée : **`Record_a` = 6 408 octets**, `Value` = 24 octets.

Un prototype conserve le code de construction et de forçage des thunks, ainsi que son interface `UnknownType`, mais partage une valeur `Unit` initialisée une fois par thread. Il ne remplace pas le calcul par une constante. Sur neuf mesures par variante, à `opt-level=1` et avec les dépendances existantes :

| Variante | Médiane |
|---|---:|
| Code généré actuel | 235,337 ms |
| Même code, `Unit` partagé | **85,204 ms** |

Soit **63,8 % de temps en moins**, environ ×2,76. Les cas 0, 1, 10 et 1 000 vérifient le résultat et l'appel unique de la continuation initiale ; chaque passage complet produit 1 000 000. Les échantillons présentent quelques valeurs hautes, conservées dans `results-unit.json`.

Le comptage séparé donne 6 004 000 allocations et 6 582 504 000 octets demandés pour le code actuel, contre 5 003 000 allocations et 160 088 000 octets pour le prototype après échauffement. La différence représente **1 001 000 allocations et 6 422 416 000 octets cumulés**. Il ne s'agit pas de mémoire simultanément résidente. La suppression du coût répété de `Unit` est confirmée ; les cinq millions d'allocations restantes justifient une étude ultérieure des adaptateurs.

Baby steps : reproduire d'abord la construction de `Unit` dans un petit module ; introduire ensuite une représentation dédiée sans allocation, telle que `()` dans le code typé et une variante immédiate aux frontières `Value`. Vérifier les passages par `Effect`, les fonctions polymorphes et les FFI, puis relancer le benchmark complet. Le cache du prototype sert à isoler la cause ; il ne tranche pas le choix définitif d'ABI.

**Piste non confirmée :** un autre prototype reconstruit une chaîne avec un seul thunk par étape au lieu des adaptateurs visibles dans le code généré. Il ralentit le test de 230,252 à 300,579 ms, soit environ 31 %. Il ne doit donc pas être intégré en l'état. Le nombre de closures visibles ne suffit pas à prédire le temps d'exécution.

## Autres pistes, encore non mesurées

**Polymorphism — 42,6 ms, 12,1 % du total.** La boucle chaude a déjà remplacé l'appel de dictionnaire par `+ 1`, mais son accumulateur reste un `Value` avec `unwrap_int` et `mk_int` à chaque tour (`Purs_Test_Polymorphism/src/lib.rs:141`). `mk_int` construit une variante immédiate, pas une allocation sur le tas. Première étape : suivre l'instanciation `polyLoop<Int>` et garder l'accumulateur en `i64` dans un cas minimal de récursion locale. Puis vérifier que le générateur produit cette représentation et mesurer.

**Church — 11,7 ms, 3,3 % du total.** Des conversions `Func1<i64, i64>` → `Func1<Value, Value>` → `Func1<i64, i64>` apparaissent dans `Purs_Test_Church/src/lib.rs:95`, puis dans les constructions de `c100`, `c10k` et `c100k`. Première étape : reproduire un aller-retour sur une fonction typée et suivre ses arguments `TypeApp`, puis vérifier si les adaptateurs inverses peuvent être supprimés sans modifier l'ordre des effets. Aucun facteur d'accélération n'est établi.

Le TAST v3 fournit ici les instanciations utiles. Dans `CodeGen.purs:982`, la branche `Syn.TypeApp a ty` descend actuellement dans `a` sans exploiter directement `ty` à cet endroit. C'est un point précis à tracer pour ces deux pistes ; cela ne signifie pas que les annotations typées ne sont utilisées nulle part ailleurs dans la chaîne.

**Profil Rust.** Le projet généré utilise `opt-level=1` et `debug=true`. Une comparaison isolée RBTree donne 55,3 ms à O1 contre 59,8 ms à O3 dans la même série exploratoire. Passer globalement à O3 n'est donc pas un gain démontré. LTO et les unités de codegen n'ont pas été évalués dans cet audit.

## Reproduction et traçabilité

Outils : `rustc 1.96.0 (ac68faa20 2026-05-25)`, mimalloc lié depuis les dépendances release existantes. Aucun benchmark CPU n'a été exécuté en parallèle par cet audit. Les mesures ne comprennent ni compilation ni contrôles préalables. Elles n'ont pas été réalisées avec affinité CPU ou fréquence verrouillée.

Depuis la racine d'`altbak.pub-purust` :

```sh
RUST_AUDIT_VARIANTS=generated-o1,color-copy-o1,borrow-tests-o1,color-borrow-o1,compact-balance-o1,color-compact-o1 python3 scratch/rust-audit-20260908/measure.py
RUST_AUDIT_LAZY_VARIANTS=lazy-generated,lazy-cached-unit RUST_AUDIT_LAZY_RESULTS=results-unit.json python3 scratch/rust-audit-20260908/measure-lazy.py
python3 scratch/rust-audit-20260908/measure-unit-alloc.py
```

`measure.py` contient aussi les variantes O3 et les versions instrumentées pour les allocations RBTree. Les résultats de l'audit sont archivés dans `results-o1.json`, `results-o3.json`, `results-lazy.json`, `results-unit.json` et `results-unit-alloc.json`. Relancer les scripts régénère les copies Rust et les exécutables dans ce seul dossier. Ils demandent une version non ambiguë des bibliothèques `.rlib` déjà compilées.

Empreintes SHA-256 des sources générées étudiées :

- RBTree : `60616a63e65f281c0a481476a2db856f18fd8f61485a9163e2344210178a77ca`.
- LazyEvaluation : `dc12651db8283b5eb9de7441580f9fc330b3ec4a06885f3e2d342acae02beb8e`.
