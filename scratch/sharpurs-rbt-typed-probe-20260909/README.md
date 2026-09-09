# Prochain goulot : corps natifs de `ins` et `balance`

Analyse du 9 septembre 2026. La génération actuelle consacre environ **67 % du total à RBTree** : 97,77138 ms sur un total médian de 144,96 ms dans la dernière série complète à cinq processus. Le prochain candidat est le passage entièrement typé de `ins` et `balance`, actuellement encore en `obj`. Un prototype isolé confirme un signal CPU important : **79,09640 → 45,42430 ms (−42,57 %)** lorsque les deux variantes sont compilées sous FSI optimisé.

Ce résultat concerne le prototype. Le générateur et la DLL normale n'ont pas été modifiés pendant cette analyse.

## Génération examinée et comparaison au natif

La génération F# actuelle se trouve dans `altbak.pub-sharpurs/run/bak/sharp/output/Main`. Le chemin `output/purescript` cité dans la demande n'existe pas ici. Les fichiers correspondants du checkout principal `altbak.pub` sont plus anciens ; les fichiers RBTree et Lazy du worktree sont identiques aux snapshots après de l'intégration des additions/soustractions natives.

| Benchmark | Compilé, README officiel | Compilé, dernière série complète | Natif optimisé, dernière colonne du README |
| --- | ---: | ---: | ---: |
| RBTree | 100,07783 ms | 97,77138 ms | 65,156 ms |
| Lazy | 71,20333 ms | 37,78758 ms | 0,075 ms |
| Total | 185,39 ms | 144,96 ms | 69,3 ms |

Sources : [README officiel](/Users/0x1/Documents/htdocs/altbak.pub/README.md:154) et [dernière série complète](/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/scratch/sharpurs-int-arithmetic-native-20260909/README.md). Les chiffres historiques du README ne sont pas des témoins simultanés du prototype.

Le [natif RBTree optimisé](/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/src/Test/RBTreeFFICheatcode.fs:2) emploie un `SortedSet<int>` mutable, insère en ordre croissant et renvoie `Count`. Le PureScript construit un arbre persistant d'Okasaki en ordre décroissant et calcule sa profondeur, 22. L'écart avec 65,156 ms n'est donc pas entièrement imputable au générateur.

Le [natif FP-style](/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/src/Test/RBTreeFFI.fs:2), à 76,561 ms dans le README, est une référence structurelle plus proche : mêmes quatre rotations et mêmes champs typés. Il reste différent par l'ordre croissant des insertions et le parcours final de comptage.

Pour Lazy, le natif exécute seulement 1 000 incréments sans thunk et retourne 1 000, contre un million de thunks et un résultat de 1 000 000 en PureScript. Les 0,075 ms ne constituent pas une cible comparable.

## Cause repérée dans le code

Le [layout de Tree](/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/run/bak/sharp/output/Main/Test.RBTree.fs:10) est déjà natif : `Color * Tree * int * Tree`. Mais les corps chauds utilisent encore ces signatures :

```fsharp
balance_direct : obj -> obj -> obj -> obj -> obj
ins_tco : obj -> obj -> obj
```

Les [appels de balance sont déjà directs](/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/run/bak/sharp/output/Main/Test.RBTree.fs:43). Il reste toutefois des conversions de primitives, des comparaisons transformées en valeurs `obj` puis reconnues par `LitBool`, et des patterns `Unbox` dans les arbres. Tous les `box` imprimés ne signifient pas une allocation : les ADT sont des références, et certaines conversions peuvent être supprimées par le compilateur. Le signal mesuré porte sur l'ensemble du chemin typé, pas sur un coût individuel attribué au boxing.

Le TAST fournit explicitement `Int -> Tree -> Tree` pour `ins`, et `Color -> Tree -> Int -> Tree -> Tree` pour `balance`, ainsi que les champs dans `dataDecls`. Le [chemin ADT natif actuel](/Users/0x1/Documents/htdocs/sharpurs/sharpurs/src/Sharpurs/AdtKernel.purs:85) admet une fonction ordinaire avec exactement un argument ADT récursif. Cette restriction explique pourquoi `depth` et `makeBlack` sont natifs, tandis que `ins` et `balance` restent génériques. L'extension devra également résoudre les dépendances entre fonctions natives ; enlever seulement la restriction d'arité ne suffit pas.

## Expérience exécutée

Le prototype référence la DLL normale et réutilise ses types ADT, constructeurs, `depth`, `makeBlack` et runtime. Il remplace seulement les corps de `ins` et `balance` par des fonctions `int`/`Tree`/`Color` typées. Le wrapper public de `ins` reste en `obj`, et les corps générés de `insert` et de la boucle externe `buildTree` sont recopiés sans changement. Le garde d'exception `TargetInvocationException` autour de `balance` est conservé.

Le prototype exprime directement les patterns natifs et regroupe les quatre alternatives de rotation ayant le même résultat. Il conserve l'algorithme, l'immuabilité, l'ordre descendant des 100 000 insertions et le calcul final de profondeur. Cette expérience mesure donc la traduction native de ces deux corps, pas une modification isolée des signatures.

| Expérience | Témoin médian | Candidat médian | Écart |
| --- | ---: | ---: | ---: |
| Témoin DLL Release, candidat FSI | 81,86650 ms | 48,24620 ms | −41,07 % |
| **Deux corps sous FSI optimisé** | **79,09640 ms** | **45,42430 ms** | **−42,57 %** |

Chaque expérience utilise un processus FSI, trois échauffements par variante et cinq paires alternées. La seconde recopie aussi les corps témoins sous FSI pour contrôler la différence de compilation de la première. Étendues de cette seconde série : 76,93550–84,03200 ms pour le témoin, 44,69840–66,17230 ms pour le candidat. Les cinq valeurs candidates sont inférieures aux cinq témoins.

Chaque script passe **431 389 assertions**, dont égalité des arbres complets après insertion, invariants rouge/noir, persistance, les quatre rotations, doublons, bornes Int32 et tailles de 0 à 100 000. Le résultat chronométré reste 22. Il ne s'agit pas d'une validation exhaustive des contrats d'ABI et d'exceptions du futur générateur.

FSI signale FS0040 sur les objets récursifs repris de la génération ; les logs conservent ces avertissements. Aucune mesure de RAM ou de GC. Les résultats sont un signal de microbenchmark en processus partagé, pas une nouvelle mesure du total `App.main` ni une promesse de gain identique en Release.

## Prochaine micro-étape recommandée

1. Appliquer cette transformation seulement dans une copie du programme généré complet, en conservant wrappers publics, algorithme et harness.
2. Compiler avant/après en Release, contrôler valeurs, partage, applications partielles et exceptions, puis mesurer cinq processus par variante avec `App.main`.
3. Si le gain se confirme, étendre le générateur aux fonctions natives fermées à plusieurs arguments et à leurs appels internes, avec preuves TAST et fallback pour les formes non prises en charge.

RBTree est prioritaire : il pèse environ 67 % du total, contre 26 % pour Lazy. Les chantiers Records, Church et Primes représentent chacun moins de 3 ms dans la dernière série complète. Après RBTree, les appels génériques des thunks de Lazy restent le deuxième chantier à isoler.

Les scripts `probe.fsx`, `probe-same-fsi.fsx`, leurs logs et `inputs.json` sont conservés dans ce dossier. Reproduction : `/Users/0x1/.dotnet/dotnet fsi --nologo --optimize+ --exec probe-same-fsi.fsx`. Le script référence la DLL normale par chemin absolu ; vérifier son empreinte avec `inputs.json` avant de réutiliser cette comparaison.
