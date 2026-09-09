# RBTree : appel direct de `balance`

L’expérience isolée du 9 septembre 2026 réduit le temps médian de **RBTree de 32,48 %** et le **total de 19,51 %**, sur cinq processus frais par variante. Le seul changement est l’entrée directe de `balance` et ses deux appels saturés dans `ins`. Le gain est établi dans la copie expérimentale ; il n’est pas encore intégré au générateur.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 222,09 ms | 178,76 ms | −19,51 % (×1,242) |
| RBTree médian | 139,11992 ms | 93,93079 ms | −32,48 % (×1,481) |
| Étendue des totaux | 217,83–225,60 ms | 173,55–182,47 ms | |
| Étendue RBTree | 136,59596–141,35104 ms | 92,31313–99,93996 ms | |

La paire préliminaire donne **227,69 → 174,62 ms** au total et **143,91238 → 92,97688 ms** pour RBTree. Elle est conservée sous `probe-*` et exclue des médianes de confirmation.

Le retrait à la médiane RBTree est de **45,19 ms**. Les cinq temps RBTree après sont tous inférieurs aux cinq temps avant. Les **140 sorties**, noms et ordre correspondent à la référence historique, dont la profondeur finale **22** pour RBTree. Les médianes par test ne s’additionnent pas nécessairement à la médiane du total. Les petites variations des workloads inchangés ne sont pas attribuées au patch. Les chiffres de la précédente intégration des comparaisons Int viennent d’une autre série ; la comparaison pertinente ici emploie le témoin simultané `before/`.

`comparison.json` contient les statistiques et les dix résultats détaillés. `executions.json` conserve l’ordre des processus, les temps muraux, les charges système et les empreintes des logs. `metadata.json` identifie les versions .NET, configurations et assemblies exécutées.

## Transformation étudiée

Le témoin `before/` fige les **355 entrées de build** du programme généré actuel, avec les comparaisons Int natives déjà intégrées. `after/` modifie uniquement `Test.RBTree.fs` :

- Le corps de `balance` est conservé octet pour octet dans `Test_RBTree_balance_direct`, une fonction F# de quatre arguments `obj`.
- L’entrée publique `Test_RBTree_balance` conserve ses quatre étapes curryfiées et délègue à l’entrée directe à saturation. Les applications partielles restent disponibles.
- Les deux appels saturés depuis `ins` invoquent directement cette fonction au lieu de traverser quatre `sharpurs_apply`.

Les algorithmes, types de champs, constructeurs et comparaisons restent identiques. Le corps de `balance` reste générique ; cette expérience isole uniquement le coût des appels curryfiés. `balance-body.fs.txt`, `replacements.json`, `Test.RBTree.fs.diff` et `Test.RBTree.fs.generated` conservent la transformation exacte.

`prepare.py` capture les copies une seule fois, vérifie le corps inchangé, les deux remplacements et l’unique fichier modifié, puis enregistre `inputs.json`. Ce manifeste contient les 355 empreintes de chaque copie et celles de la génération normale. Les listes exactes de fichiers et leurs empreintes ont été contrôlées. Le générateur et ses sorties normales sont restés inchangés pendant toute l’expérience ; la génération normale correspond au témoin `before/`.

## Validation et protocole exécuté

Les deux copies ont été compilées en Release avec le même SDK .NET **8.0.423**, runtime **8.0.29**, et `-p:NuGetAudit=false` : **zéro erreur, zéro avertissement**, en **13,48 s** et **13,11 s**. `capture_build.py` relie chaque assembly à ses 355 sources, son log de build et ses huit fichiers runtime via `build-before.json` et `build-after.json`.

Avant les mesures, la vraie DLL expérimentale passe :

- **27 087 vérifications** de `VerifyTree.fsx` : ordre BST, invariants rouge/noir, contenu, profondeur, doublons, bornes Int32, persistance et partage, pour sept ordres d’insertion. Résultats dans `verify-tree.log`.
- **107 vérifications** de `VerifyBalance.fsx` : les quatre rotations, la branche par défaut, le partage des sous-arbres, les applications partielles et leur réutilisation, le report de l’exécution du corps jusqu’à saturation, l’ordre et l’évaluation unique des arguments, l’accord entre entrée directe et wrapper public. Résultats dans `verify-balance.log`.

Une paire préalable a été exécutée et analysée avant la confirmation. Les dix vrais `App.main` ont ensuite été exécutés dans des processus frais, sans build ni test concurrent, dans l’ordre **ABBAABBAAB**. Le harness conserve ses échauffements et son meilleur temps sur dix pour chacun des 14 tests. La conclusion utilise les médianes et étendues des cinq processus par variante. Aucune mesure de RAM ni de GC dans cette expérience.

Le runner contrôle les sources, assemblies et fichiers runtime avant et après chaque processus ; les configurations runtime sont identiques, sans surcharge de configuration tiering/PGO/GC dans l’environnement. Il vérifie les 14 sorties, noms et ordre contre `expected-output.log`, ainsi que la cohérence du total affiché à 0,02 ms près. La génération normale est restée identique au témoin jusqu’à la fin de la mesure.

## Référence officielle

`README.baseline.md` copie le README officiel d’altbak.pub : **294,80 ms** au total compilé, **213,760 ms** pour RBTree et **65,156 ms** dans la dernière colonne native optimisée. Le natif optimisé utilise un `SortedSet<int>` mutable et retourne sa taille ; ce repère historique a un algorithme et un résultat distincts. Le résultat expérimental de **178,76 ms** au total se compare à cette baseline historique, tandis que le gain causal mesuré ici est **222,09 → 178,76 ms** face au témoin simultané.

## Reproduction

Les copies initiales sont déjà figées. `prepare.py` refuse d’écraser `before/` ou `after/`. Compiler les deux projets avec `dotnet build before/Program.fsproj -c Release -p:NuGetAudit=false` et la commande équivalente pour `after/`, en conservant `build-before.log` et `build-after.log`. Exécuter ensuite `python3 capture_build.py before` et `python3 capture_build.py after`.

Valider les deux fixtures avec `dotnet fsi --nologo --optimize+ --exec VerifyTree.fsx` puis `dotnet fsi --nologo --optimize+ --exec VerifyBalance.fsx`. Lancer `python3 run.py --probe` puis `python3 analyze.py --probe` pour la paire préliminaire, et `python3 run.py` puis `python3 analyze.py` pour la confirmation. Aucun build ni test concurrent pendant les mesures.

Le runner refuse d’écraser une série existante ; conserver ou déplacer les anciens logs, métadonnées et comparaisons avant de répéter une série. Les copies et builds volumineux sont ignorés par Git ; scripts, preuves, empreintes et diffs restent conservés.

## Suite

Le gain justifie une extension du générateur permettant une entrée directe pour les fonctions dont la saturation et les arguments sont établis dans le TAST, tout en conservant leur interface curryfiée. Cette intégration devra être sélectionnée par forme et typage, sans règle spécifique au nom `balance` ou au benchmark, puis mesurée séparément.
