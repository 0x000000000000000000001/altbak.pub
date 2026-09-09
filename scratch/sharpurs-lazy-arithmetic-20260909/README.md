# LazyEvaluation : arithmétique Int native

L’expérience isolée du 9 septembre 2026 réduit le temps médian de **LazyEvaluation de 46,51 %** et le **total de 17,91 %**, sur cinq processus frais par variante. Elle remplace uniquement deux expressions arithmétiques Int dans une copie du F# généré. Ce gain est établi dans la copie expérimentale ; **il n’est pas encore intégré au générateur**.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 176,29 ms | 144,71 ms | −17,91 % (×1,218) |
| LazyEvaluation médiane | 69,23067 ms | 37,03129 ms | −46,51 % (×1,870) |
| Étendue des totaux | 174,62–183,51 ms | 143,07–150,53 ms | |
| Étendue LazyEvaluation | 67,78463–72,61408 ms | 36,06463–39,21408 ms | |
| RBTree médiane, contrôle inchangé | 95,56808 ms | 94,14329 ms | −1,49 % |

La paire préliminaire donne **183,44 → 152,94 ms** au total et **73,12808 → 39,12888 ms** pour LazyEvaluation. Elle est conservée sous `probe-*` et exclue des médianes de confirmation.

Les cinq temps LazyEvaluation après sont tous inférieurs aux cinq temps avant. Le retrait de **32,20 ms** sur sa médiane suffit à expliquer l’ordre de grandeur du gain total. Les médianes par test ne s’additionnent pas nécessairement à la médiane du total. RBTree sert de contrôle secondaire inchangé ; ses fluctuations ne sont pas attribuées au patch. Les **140 sorties**, noms et ordre correspondent à la référence historique, dont **1 000 000** pour LazyEvaluation et **22** pour RBTree.

`comparison.json` contient les statistiques et dix résultats détaillés. `executions.json` conserve l’ordre des processus, temps muraux, charges système et empreintes des logs. `metadata.json` identifie les versions .NET, configurations, environnement et assemblies exécutées.

## Transformation étudiée

`before/` fige les **355 entrées de build** du programme généré actuel, qui inclut les appels directs saturés déjà intégrés. `after/` diffère dans **un seul fichier, une seule ligne** : `Test.LazyEvaluation.fs`, ligne 15, dans `Test_LazyEvaluation_buildThunks_tco`.

- La soustraction canonique `Data.Ring.sub ringInt n 1` devient une soustraction F# native entre deux `int`, puis son résultat est boxé pour conserver l’ABI `obj`.
- L’addition canonique `Data.Semiring.add semiringInt (force acc) 1` devient une addition F# native, toujours à l’intérieur de la closure différée, avec le même résultat boxé.

Ces deux remplacements retirent **six appels génériques par thunk**, soit six millions sur le workload complet, ainsi que les deux recherches de membres dans les dictionnaires par thunk. Ils conservent les opérandes, leur nombre d’évaluations, les closures, `defer`, `force`, la récursion et tous les autres fichiers. L’addition reste différée ; elle conserve l’arithmétique Int32 avec débordement circulaire. Aucun changement des autres opérations arithmétiques du module ni du harness.

`prepare.py` capture les copies une seule fois, exige les dictionnaires canoniques exacts et les formes attendues des deux opérandes, puis vérifie l’unique ligne modifiée et les six appels retirés. `replacements.json`, `Test.LazyEvaluation.fs.before`, `Test.LazyEvaluation.fs.diff` et `Test.LazyEvaluation.fs.generated` conservent les expressions et le diff exacts.

`inputs.json` identifie les 355 sources de chaque copie et la génération normale. Les ensembles exacts de fichiers et leurs empreintes sont contrôlés avant et après chaque processus. La génération normale doit correspondre au témoin `before/` et est restée inchangée. Le générateur normal n’a pas été modifié : son commit et l’empreinte du bundle correspondent encore à la capture initiale, et son dépôt reste propre.

## Validation et protocole exécuté

Les deux copies ont été compilées en Release avec le même SDK .NET **8.0.423**, runtime **8.0.29**, et `-p:NuGetAudit=false` : **zéro erreur, zéro avertissement**, en **12,31 s** et **11,87 s**. `capture_build.py` relie chaque assembly à ses 355 sources, son log de build et ses huit fichiers runtime via `build-before.json` et `build-after.json`.

Avant les mesures, `VerifyLazy.fsx` passe **333 vérifications sur chaque vraie DLL**, avant et après, soit 666 assertions réussies. La fixture couvre les profondeurs 0, 1, 3 et 1 000 ; les bornes Int32 avec un oracle arithmétique indépendant calculé en Int64 ; le report de l’évaluation ; le nombre de forçages ; l’absence de mémoïsation ; les applications partielles réutilisées ; l’accord entre ABI directe et curryfiée ; l’identité de la cause des exceptions et leurs enveloppes d’application ; l’ordre des arguments ; le résultat du workload complet à un million de thunks. Résultats dans `verify-before.log` et `verify-after.log`.

Une paire préalable a été exécutée et analysée avant la confirmation. Les dix vrais `App.main` ont ensuite été exécutés dans des processus frais, sans build ni test concurrent, dans l’ordre **ABBAABBAAB**. Le harness conserve ses échauffements et son meilleur temps sur dix pour chacun des 14 tests. La conclusion utilise les médianes et étendues des cinq processus par variante. Aucune mesure de RAM ni de GC dans cette expérience.

Le runner vérifie les 14 sorties, noms et ordre contre `expected-output.log`, ainsi que la cohérence du total affiché à 0,02 ms près. Les configurations runtime sont identiques, sans surcharge tiering/PGO/GC dans l’environnement. Les manifestes de sources, logs de build, assemblies et fichiers runtime sont vérifiés pendant toute la série. La référence de sorties vient de l’audit précédent `sharpurs-rbt-balance-direct-20260909`.

## Référence officielle

`README.baseline.md` copie le README officiel d’`altbak.pub`, empreinte SHA-256 `7317993e8512d5d0cc4253383e87272f9d86ae9da446f45d44be3e097d3e7a3d`. Le parseur vérifie explicitement **185,39 ms** au total compilé, **71,20333 ms** pour LazyEvaluation, **100,07783 ms** pour RBTree et **0,075 ms** pour LazyEvaluation dans la dernière colonne native optimisée.

Le natif optimisé n’exécute que **1 000 incréments** et renvoie **1 000**, tandis que le programme PureScript force **1 000 000 de thunks** et renvoie **1 000 000**. Son temps ne constitue donc pas une cible directement comparable. Le résultat expérimental total de **144,71 ms** est inférieur de 21,94 % au repère historique de 185,39 ms ; le gain attribuable au patch est celui mesuré face au témoin simultané : **176,29 → 144,71 ms**, soit 17,91 %.

## Reproduction

Les copies sont déjà figées. `prepare.py` refuse d’écraser `before/` ou `after/`. Compiler les deux projets avec `dotnet build before/Program.fsproj -c Release -p:NuGetAudit=false` et la commande équivalente pour `after/`, en conservant `build-before.log` et `build-after.log`. Exécuter ensuite `python3 capture_build.py before` et `python3 capture_build.py after`.

Valider les DLL avec `dotnet fsi --nologo --optimize+ --define:BASELINE --exec VerifyLazy.fsx`, puis la même commande sans `--define:BASELINE`. Lancer `python3 run.py --probe` puis `python3 analyze.py --probe` pour la paire préliminaire ; `python3 run.py` puis `python3 analyze.py` pour la confirmation. Aucun build ni test concurrent pendant les mesures.

Le runner refuse d’écraser une série existante ; conserver ou déplacer ses logs, métadonnées et comparaisons avant de répéter les mesures. Les copies et builds volumineux sont ignorés par Git ; scripts, preuves, empreintes et diff restent conservés.

## Suite

Le gain justifie d’intégrer une règle générale pour l’addition et la soustraction Int canoniques, sélectionnée par le TAST : types et instanciations `TypeApp`, dictionnaire exact, saturation des appels. Préserver les chemins génériques pour les types, dictionnaires ou formes non prouvés. L’intégration devra être indépendante des noms `LazyEvaluation` et `buildThunks`, conserver les applications partielles, puis être validée et remesurée séparément sur la génération normale.
