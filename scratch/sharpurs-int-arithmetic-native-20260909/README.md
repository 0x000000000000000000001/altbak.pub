# Int add/sub : intégration dans la génération normale

La règle générale d’addition et de soustraction Int est intégrée au compilateur Sharpurs. La génération normale passe de **188,72 à 144,96 ms au total** (−23,19 %) et de **74,37625 à 37,78758 ms pour LazyEvaluation** (−49,19 %), sur cinq processus frais par variante. La DLL normale est recompilée à partir des mêmes 355 sources que la variante après mesurée.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 188,72 ms | 144,96 ms | −23,19 % (×1,302) |
| LazyEvaluation médiane | 74,37625 ms | 37,78758 ms | −49,19 % (×1,968) |
| Étendue des totaux | 184,27–192,64 ms | 139,14–150,04 ms | |
| Étendue LazyEvaluation | 70,93121–75,01458 ms | 37,22942–38,25442 ms | |
| RBTree médiane | 101,00929 ms | 97,77138 ms | −3,21 % |
| Records médiane | 4,30433 ms | 2,45288 ms | −43,01 % |
| Church médiane | 4,23229 ms | 2,17625 ms | −48,58 % |
| Ackermann médiane | 0,45496 ms | 0,15633 ms | −65,64 % |

La paire préliminaire séparée donne **183,72 → 140,93 ms** au total et **71,30592 → 38,27917 ms** pour LazyEvaluation. Elle reste sous `probe-*` et est exclue des médianes de confirmation.

Les cinq temps totaux après et les cinq temps Lazy après sont tous inférieurs aux cinq temps avant correspondants. Les 140 sorties, noms et ordre sont conformes à la référence historique, dont 1 000 000 pour LazyEvaluation et 22 pour RBTree. Les médianes par test ne s’additionnent pas nécessairement à la médiane du total. L’intégration optimise plusieurs modules ; le gain total ne doit pas être attribué uniquement aux deux opérations de `buildThunks`. RBTree reçoit aussi une conversion dans `buildTree` et n’est donc pas un contrôle strictement inchangé ; son faible écart et ses étendues qui se chevauchent ne suffisent pas à isoler son gain propre.

## Transformation validée

Le nouveau reconnaisseur `Sharpurs.IntArithmetic.fromExpr` exige l’identité canonique exacte de `Data.Semiring.add` / `Data.Ring.sub`, le dictionnaire `semiringInt` / `ringInt` et sa classe correspondante, une saturation complète, des opérandes Int, et la cohérence des annotations TAST de tous les appels intermédiaires. Il accepte une signature concrète cohérente ou une instanciation simple `TypeApp Int` dont le `ForAll` et la signature contrainte concordent.

Le générateur émet une addition ou soustraction F# entre opérandes `unbox<int> (box (...))`, puis boxe le résultat pour préserver l’ABI `obj`. Il évalue chaque opérande une seule fois, de gauche à droite. Les dictionnaires personnalisés, nombres flottants, applications partielles, annotations absentes ou contradictoires restent génériques. Aucun nom de benchmark n’intervient dans la règle. Les closures, `force`, `defer`, le harness et les résultats attendus restent identiques.

L’inventaire compte **166 sites admissibles bruts dans 38 modules**, dont trois déjà natifs (`TCO.deepTailRec` ×2, `RBTree.depth` ×1). La génération convertit effectivement **163 sites dans 37 fichiers F# : 69 additions et 94 soustractions**. LazyEvaluation contient quatre conversions, deux dans `buildThunks` et deux dans `runManyTimes` ; RBTree en contient une dans `buildTree`. `generation-inventory.json` conserve ces comptes, et les fichiers `*.fs.diff` / `*.fs.generated` conservent les changements exacts.

Le TAST v3 est bien utilisé. La fixture vérifie toutefois un cas concret du fork où le `@Int` visible entoure l’application du dictionnaire et laisse indisponible l’annotation de la fonction canonique d’origine. Cette forme reste volontairement sur le chemin générique : les preuves requises par le reconnaisseur ne sont pas présentes. Les instanciations implicites cohérentes `TypeApp Int` des opérations testées sont reconnues. Cette limite ne suppose pas un effacement général des types par le TAST.

## Validation

Les builds Release avant, après et normal utilisent le SDK .NET **8.0.423**, runtime **8.0.29**, et `-p:NuGetAudit=false`. Les trois réussissent avec **zéro erreur et zéro avertissement**, en 32,23 s, 33,40 s et 32,69 s. Les manifestes `build-before.json`, `build-after.json` et `build-normal.json` relient sources, logs et huit fichiers runtime pour chaque variante.

La nouvelle fixture `test:int-arithmetic` passe **1 648 vérifications runtime F# et 317 vérifications JavaScript**, soit 91 vérifications du reconnaisseur/générateur et 226 comparaisons de l’oracle JavaScript avec BigInt. Elle utilise le vrai fork PureScript et son TAST décodé par PBO. Elle couvre les limites et débordements Int32, les valeurs signées, l’ordre d’évaluation et les exceptions, les opérations canoniques, les formes refusées, les dictionnaires personnalisés et les applications partielles. Son dernier passage est sans avertissement PureScript ni FSI. `int-arithmetic-validation/` contient le TAST, les programmes et oracles générés, et le log final ; `validation-files.json` identifie leurs empreintes.

Sur les DLL complètes :

- `VerifyLazy.fsx` : **333 vérifications avant et 333 après**, incluant paresse, débordement, exceptions, ABI, absence de mémoïsation et workload complet.
- `VerifyTree.fsx` : **27 087 vérifications après**, incluant invariants RBTree, profondeurs, ordres d’insertion variés, doublons, bornes signées et partage.
- `VerifyBalance.fsx` : **138 vérifications après**, pour rotations, partage, application partielle, report d’évaluation et ordre des arguments.
- `VerifyCodePoints.fsx` : **165 vérifications après**, pour paires de surrogates, applications partielles et appelants Unicode.

Les huit suites existantes passent également : noyau Int 695 ; noyau local 43 runtime + 68 converter ; noyau ADT 32 + 33 ; interop ADT 62 + 13 ; ADT unaire 63 + 11 ; comparaisons Int 396 + 32 ; appels directs 130 + 67 ; application runtime 21. `spago test` réussit avec 43 + 6 assertions. Les logs `sharpurs-int-arithmetic-*.log` sont archivés ici. Les builds PureScript du backend conservent leurs avertissements existants ; ils se distinguent des builds Release .NET et de la nouvelle fixture sans avertissement.

## Protocole et provenance

`before/` fige les 355 entrées du programme généré avant intégration. Elles correspondent exactement au témoin avant de `sharpurs-lazy-arithmetic-20260909`. Le commit backend de référence est `77ce414759f9b9a11a7ee0bbfd818d8874c4df6e`. `after/` capture les 355 entrées de la nouvelle génération normale : aucun patch F# expérimental n’y est appliqué. Seuls les 37 fichiers F# inventoriés changent ; les autres entrées et le harness restent identiques.

`inputs.json` contient les empreintes avant/après, celles de la génération normale, ainsi que le commit, le bundle et les sources du backend au moment de chaque capture. `sourceGeneratedSha256` doit être identique à `afterSha256`. Les fichiers exacts, leurs empreintes, les logs de build, assemblies et autres fichiers runtime sont vérifiés avant et après chaque processus. Les configurations runtime des variantes sont identiques, sans surcharge de tiering, PGO ou GC dans l’environnement.

Après une paire préliminaire valide, dix vrais `App.main` sont exécutés séquentiellement dans des processus frais, dans l’ordre **ABBAABBAAB**, sans build ni test concurrent. Le harness conserve ses échauffements et son meilleur temps sur dix pour chacun des 14 tests. Les statistiques finales utilisent médianes et étendues des cinq processus par variante. Chaque sortie est comparée à la référence historique, et le total est vérifié contre la somme des temps affichés à 0,02 ms près. Aucune mesure de RAM ou GC n’est utilisée.

`comparison.json` conserve les statistiques et dix résultats détaillés. `executions.json` conserve l’ordre, les temps muraux, les charges système et les empreintes des logs. `metadata.json` identifie les versions .NET, configurations, environnement et assemblies. La nouvelle fixture, dont deux imports inutiles ont été nettoyés, a été réexécutée une dernière fois après la fin des benchmarks ; ce nettoyage ne modifie ni le générateur ni les sources ou DLL mesurées.

## Repère officiel

`README.baseline.md` copie le README officiel d’`altbak.pub`, SHA-256 `7317993e8512d5d0cc4253383e87272f9d86ae9da446f45d44be3e097d3e7a3d`. Le parseur exige **185,39 ms** au total compilé, **71,20333 ms** pour LazyEvaluation, **100,07783 ms** pour RBTree et **0,075 ms** pour LazyEvaluation dans la dernière colonne native optimisée.

Le natif Lazy exécute seulement **1 000 incréments** et renvoie **1 000**, tandis que PureScript force **1 000 000 de thunks** et renvoie **1 000 000**. Son temps n’est pas une cible directement comparable. Le nouveau total de 144,96 ms est inférieur de 21,81 % au repère historique de 185,39 ms ; l’écart attribuable à cette intégration est celui mesuré face au témoin simultané : **188,72 → 144,96 ms**, soit 23,19 %.

## Reproduction

Les snapshots sont figés et `snapshot_before.py` / `snapshot_after.py` refusent leur écrasement. Pour reconstruire les copies, lancer `dotnet build before/Program.fsproj -c Release -p:NuGetAudit=false` et la commande équivalente après, en conservant `build-before.log` et `build-after.log`, puis exécuter `python3 capture_build.py before` et `python3 capture_build.py after`. La génération normale se construit dans `run/bak/sharp/output/Main` ; conserver son log dans `build-normal.log` et exécuter `python3 capture_normal_build.py`.

Exécuter les fixtures FSI avec `dotnet fsi --nologo --optimize+ --exec VerifyLazy.fsx` et les trois autres fichiers `Verify*.fsx`. Ajouter `--define:BASELINE` pour la validation Lazy avant. Exécuter la nouvelle fixture du backend avec `npm run test:int-arithmetic` ; ses artefacts ici correspondent au passage final archivé.

Pour les mesures, lancer `python3 run.py --probe` puis `python3 analyze.py --probe`, ensuite `python3 run.py` puis `python3 analyze.py`. Le runner refuse d’écraser une série existante : conserver ou déplacer ses fichiers avant une répétition. Aucun build ou test concurrent pendant les mesures. Les copies et builds volumineux sont ignorés par Git ; scripts, résultats, empreintes, diff et fixtures restent conservés.
