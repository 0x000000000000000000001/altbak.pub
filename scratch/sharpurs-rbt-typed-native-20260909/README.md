# RBTree : corps `ins` et `balance` entièrement typés, expérience Release

Le changement isolé réduit le temps médian **RBTree de 90,28363 à 52,59900 ms (−41,74 %)** et le **total de 139,75 à 102,11 ms (−26,93 %)**, sur cinq processus frais par variante. Ce résultat confirme le signal FSI dans le programme complet en Release. Il concerne la copie expérimentale : le générateur, ses sorties normales et sa DLL normale restent inchangés.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 139,75 ms | 102,11 ms | −26,93 % (×1,369) |
| RBTree médiane | 90,28363 ms | 52,59900 ms | −41,74 % (×1,716) |
| Étendue des totaux | 136,95–145,47 ms | 99,31–102,92 ms | |
| Étendue RBTree | 89,44692–95,63646 ms | 50,57254–53,14563 ms | |
| Lazy médiane, code inchangé | 36,96663 ms | 35,16783 ms | −4,87 % |

Les cinq temps après sont tous inférieurs aux cinq temps avant pour RBTree et le total. La baisse à la médiane RBTree est de **37,68463 ms**, proche des **37,64 ms** de baisse du total. Les **140 sorties**, leurs noms et leur ordre sont conformes à la référence historique, dont la profondeur finale **22** pour RBTree.

La paire préliminaire séparée donne **145,52 → 102,97 ms** au total et **95,78188 → 52,98971 ms** pour RBTree ; ses 28 sorties sont conformes. Conservée dans `probe-*`, elle est exclue des médianes de confirmation. Les variations des workloads inchangés, dont Lazy, ne sont pas attribuées au code RBTree ciblé. Les médianes par test ne s’additionnent pas nécessairement à celle du total. Les chiffres de la précédente intégration Int proviennent d’une autre série : l’effet mesuré ici utilise le témoin simultané `before/`.

## Transformation exacte

`before/` fige les **355 entrées de build** de la génération normale actuelle, avec les additions/soustractions Int déjà intégrées. Les empreintes correspondent exactement à la variante après de `sharpurs-int-arithmetic-native-20260909`. Dans `after/`, **seul `Test.RBTree.fs` change** :

- Ajout de `Test_RBTree_balance_adt_native`, avec paramètres `Color`, `Tree`, `int`, `Tree`, résultat `Tree` et les quatre rotations Okasaki utilisant directement les patterns et constructeurs de l’ADT existant.
- Ajout de son entrée gardée `Test_RBTree_balance_adt_native_apply`, qui conserve l’enveloppe `TargetInvocationException` de l’appel saturé courant.
- Ajout de `Test_RBTree_ins_adt_native`, récursif et typé `int -> Tree -> Tree`, appelant ce noyau de balance avec les comparaisons Int existantes.
- Remplacement du seul corps `ins_tco` par un pont entre l’ABI `obj` et ce noyau typé, avec conversion des arguments à l’entrée et boxage du résultat à la sortie.

Les anciennes entrées `balance_direct`, `balance_direct_apply` et `balance` restent identiques. Les wrappers publics, `insert`, `buildTree`, `makeBlack`, `depth`, les layouts des constructeurs, les autres modules, le FFI, les projets et le harness sont inchangés. L’arbre reste persistant : **100 000 insertions descendantes**, mêmes rotations, ordre des comparaisons, sous-arbres et profondeur finale. Les deux nouveaux corps natifs n’utilisent ni paramètres `obj` ni `sharpurs_apply`.

Cette expérience groupe le typage des paramètres et les patterns natifs ; elle ne sépare pas leurs contributions individuelles. Le patch est extrait du prototype mesuré précédemment, puis installé seulement dans la copie expérimentale. `prepare.py` contrôle l’unique définition remplacée et les anciennes entrées préservées. `native-kernels.fs.txt`, `Test.RBTree.fs.generated` et `Test.RBTree.fs.diff` conservent le code exact.

## Validation exécutée

Les deux projets complets sont compilés en Release avec le même SDK .NET **8.0.423**, runtime **8.0.29**, et `-p:NuGetAudit=false` : **zéro avertissement, zéro erreur**, en **32,30 s** avant et **32,60 s** après. `build-before.json` et `build-after.json` relient les 355 sources, le log de build et les huit fichiers runtime de chaque variante.

Les dix validations FSI optimisées chargent les vraies DLL Release, sans recompiler le programme. Elles passent avant toute mesure, avec zéro avertissement :

| Fixture | Avant | Après | Vérification |
| --- | ---: | ---: | --- |
| `VerifyIns.fsx` | 522 281 | 524 959 | Structure, ordre BST, invariants rouge/noir, persistance, ABI, ordre d’évaluation, charge complète |
| `VerifyTree.fsx` | 27 087 | 27 087 | Sept ordres d’insertion, profondeurs, doublons, bornes Int32, partage |
| `VerifyBalance.fsx` | 138 | 190 | Quatre rotations, branche par défaut, partage, applications partielles et réutilisation, report et ordre d’évaluation ; helpers natifs après |
| `VerifyLazy.fsx` | 333 | 333 | Arithmétique, paresse, exceptions, ABI et charge complète inchangées |
| `VerifyCodePoints.fsx` | 165 | 165 | Surrogates, applications partielles et appelants Unicode inchangés |

`VerifyIns.fsx` compare **1 067 observations exactes** de la DLL après avec un oracle enregistré par la DLL avant : formes complètes des arbres, y compris la charge de 100 000 insertions et sa profondeur 22. Les logs `validation-{before,after}-*.log`, fixtures et oracle sont conservés ; `validation-files.json` identifie leurs empreintes finales. Les preuves s’appliquent aux valeurs conformes aux types PureScript ; les entrées FFI délibérément mal typées restent des diagnostics séparés, hors critère de cette expérience.

## Protocole et provenance

Après validation puis une paire préliminaire, dix vrais `App.main` s’exécutent séquentiellement dans des processus frais, dans l’ordre **ABBAABBAAB**, sans build ni test concurrent. Le harness conserve ses échauffements et son meilleur temps sur dix pour chacun des 14 tests. La conclusion utilise les médianes et étendues des cinq processus par variante. Aucune mesure de RAM ou de GC n’est utilisée.

`run.py` vérifie les ensembles exacts des 355 fichiers et leurs empreintes avant/après/normal avant et après chaque processus. Il exige que seule `Test.RBTree.fs` diffère et que la génération normale corresponde au témoin avant. Il vérifie aussi les assemblies et autres fichiers runtime mesurés, leurs logs de build, la **DLL normale et ses fichiers runtime**, ainsi que le **bundle et les sources backend** figés dans `inputs.json`. Tous ces contrôles réussissent jusqu’à la fin de la série.

Les configurations runtime sont identiques, sans surcharge de tiering, PGO ou GC dans l’environnement. Le runner compare les 14 sorties, noms et ordre à `expected-output.log`, et contrôle le total affiché contre la somme des tests à 0,02 ms près. `comparison.json` conserve toutes les valeurs et statistiques ; `executions.json` l’ordre, les temps muraux, charges système et empreintes des logs ; `metadata.json` les versions, configurations, environnement et assemblies.

## Repère officiel

`README.baseline.md` fige le README principal d’`altbak.pub`, SHA-256 `b426ab9570200d8024fda4549142c10c3fcbc59c147fc9777f1d5b0233916a76` : **185,39 ms** au total compilé, **100,07783 ms** pour RBTree, et **65,156 ms** pour RBTree dans la dernière colonne native optimisée.

Le natif optimisé utilise un `SortedSet<int>` mutable, des insertions ascendantes et retourne `Count = 100000`. Le programme compilé conserve un arbre Okasaki persistant, des insertions descendantes et retourne la profondeur 22. Ces algorithmes et résultats distincts limitent la comparaison directe ; **52,599 ms** dans l’expérience n’établit pas une supériorité à charge identique sur le natif. Le gain attribuable à cette transformation se mesure face au témoin simultané : **90,28363 → 52,59900 ms** pour RBTree.

## Reproduction

`prepare.py` refuse d’écraser les snapshots ou `inputs.json`. Les captures existantes permettent de reconstruire `before/Program.fsproj` puis `after/Program.fsproj` avec `dotnet build <projet> -c Release -p:NuGetAudit=false`, en conservant respectivement `build-before.log` et `build-after.log`, puis `python3 capture_build.py before` et `python3 capture_build.py after`.

Valider chaque fixture avec `dotnet fsi --nologo --optimize+ --define:BASELINE --exec VerifyIns.fsx` puis sans `--define:BASELINE`, et faire de même pour les quatre autres fixtures. La première exécution de `VerifyIns.fsx` enregistre l’oracle avant ; l’exécution après le compare.

Lancer `python3 run.py --probe` puis `python3 analyze.py --probe`, ensuite `python3 run.py` puis `python3 analyze.py`. Le runner refuse d’écraser une série existante ; conserver ou déplacer ses logs et métadonnées avant toute répétition. Ne pas lancer de build ou test concurrent pendant les mesures. Les copies et builds volumineux sont ignorés par Git ; scripts, preuves, empreintes, diffs et validations restent conservés.

## Suite

Le résultat justifie l’intégration d’une règle générale de noyaux ADT multiarguments dans le générateur. Elle devra utiliser les annotations et déclarations TAST, préserver l’interface curryfiée, la sémantique des appels et la voie générique lorsque les preuves de types manquent, sans cibler un nom de benchmark. La génération normale devra ensuite faire l’objet d’une validation et d’une mesure propres.
