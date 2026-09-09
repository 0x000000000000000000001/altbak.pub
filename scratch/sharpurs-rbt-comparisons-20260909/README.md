# RBTree : comparaisons Int directes dans ins

Expérience du 9 septembre 2026, limitée à deux comparaisons dans une copie du F# généré. **RBTree baisse de 36,25 % et le total de 25,00 %**, sur les médianes de cinq processus par variante. Le gain est obtenu en conservant l’algorithme, les appels à `balance` et tous les autres fichiers générés. Le changement reste expérimental ; il n’est pas intégré au générateur.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 282,96 ms | 212,23 ms | −25,00 % |
| RBTree médian | 204,19058 ms | 130,17196 ms | −36,25 % (×1,569) |
| Étendue des totaux | 280,86–288,16 ms | 208,97–213,97 ms | |
| Étendue RBTree | 201,98221–207,04246 ms | 127,50125–131,84438 ms | |

La paire préalable donnait **282,95 → 207,10 ms** au total et **205,49154 → 128,53008 ms** pour RBTree. Elle est conservée séparément dans `probe-*` et n’entre pas dans les médianes finales. Les dix processus de confirmation suivent l’ordre ABBAABBAAB.

Les **140 sorties affichées**, leurs noms et leur ordre concordent avec la référence historique ; la profondeur RBTree reste **22**. Les configurations runtime des deux variantes sont identiques. Les 355 entrées des deux copies et de la génération normale ont été vérifiées par empreintes durant toute la mesure. `comparison.json` contient les temps de chaque test et chaque processus ; `executions.json` conserve les temps muraux, la charge système et les empreintes des assemblies/logs, et `metadata.json` les versions et paramètres runtime.

Ce changement isolé retire environ **74,02 ms** à la médiane RBTree. Les médianes par test ne s’additionnent pas nécessairement à la médiane du total ; les petites variations des autres tests inchangés ne sont pas attribuées à ce patch. La comparaison démontre le coût du chemin générique des comparaisons dans ce workload. Elle ne mesure pas encore le bénéfice d’une transformation intégrée au backend.

La prochaine étape est d’émettre les comparaisons entières résolues par TAST/PBO dans le générateur, puis de retrouver ce gain dans la génération normale. La règle doit porter sur les opérations et les types reconnus, sans nom de benchmark. Le passage natif complet de `ins`/`balance` demeure un chantier distinct.

## Changement isolé

La référence provient de la génération actuelle du worktree `altbak.pub-sharpurs`, qui contient déjà le fast path `obj -> obj`, le layout ADT typé et les noyaux natifs `depth`/`makeBlack`. Elle évite de réintroduire le prélude plus ancien présent dans la génération du dépôt principal `altbak.pub`.

`prepare.py` copie les 355 entrées de build dans `before/` et `after/`. Dans `after/Test.RBTree.fs`, il remplace uniquement les expressions de comparaison de `Test_RBTree_ins_tco` :

```fsharp
box ((unbox<int> x) < y)
box ((unbox<int> x) > y)
```

Le champ `y` est déjà un `int` dans le layout natif ; `x` reste un argument `obj`. Les branches, le pattern booléen, les appels à `balance`, les constructeurs, l’algorithme, l’ordre d’insertion, la profondeur finale et le harness sont conservés. Le patch retire les appels génériques `Data_Ord_lessThan/greaterThan ordInt x y` à ces deux sites uniquement. Le backend et la génération normale restent inchangés.

`replacements.json` conserve les deux expressions exactes avant/après. `Test.RBTree.fs.diff` et `inputs.json` prouvent qu’une seule ligne d’un seul fichier de build diffère. `provenance.json` identifie les commits et les empreintes du backend.

## Validation et protocole

- Les deux copies sont compilées en Release avec le même SDK et `-p:NuGetAudit=false` : zéro erreur, zéro avertissement ; logs `build-before.log` et `build-after.log`.
- `VerifyTree.fsx`, repris du jalon précédent, vérifie la vraie DLL expérimentale : ordre BST, couleurs, hauteurs noires, contenu, profondeur, doublons, bornes Int32 et persistance. Les **27 087 assertions** passent (`verify-tree.log`). Les ordres croissant et décroissant exercent les deux comparaisons.
- Première paire isolée avec `python3 run.py --probe`, puis cinq processus frais par variante, ordre **ABBAABBAAB**, via `python3 run.py`. Aucune compilation ou autre mesure de cette tâche pendant cette série.
- Chaque processus exécute le vrai `App.main` et conserve les échauffements et les meilleurs temps sur dix du benchmark. L’analyse vérifie les noms, l’ordre et les 14 sorties affichées ; le calcul RBTree conserve 100 000 insertions et la profondeur **22**.
- Les médianes et les étendues sont calculées par `python3 analyze.py`. La sélection se fait sur le temps d’exécution ; aucun diagnostic d’allocations n’est ajouté à cette expérience.
- Les empreintes des entrées, des assemblies et des configurations runtime accompagnent les logs. Les sources générées normales sont contrôlées avant/après les exécutions.

## Référence officielle

`README.baseline.md` fige le README officiel d’altbak.pub : **294,80 ms** pour le total compilé et **213,760 ms** pour RBTree. La dernière colonne native donne **65,156 ms** pour RBTree ; elle utilise `SortedSet<int>`, un ordre ascendant et retourne le nombre d’éléments. Elle constitue un repère historique, pas une mesure du gain de ce patch sur le même algorithme. Le témoin simultané est la copie `before/`.

## Reproduction

La première préparation s’exécute une seule fois avec `python3 prepare.py` ; elle refuse d’écraser les snapshots existants. Compiler ensuite `before/Program.fsproj` et `after/Program.fsproj` avec `dotnet build -c Release -p:NuGetAudit=false`. Lancer les vérifications avec `dotnet fsi --nologo --optimize+ --exec VerifyTree.fsx`, puis `python3 run.py` et `python3 analyze.py` sans build concurrent.

Les répertoires volumineux des copies et des builds sont ignorés par Git ; les scripts, logs, empreintes, patch et résultats d’analyse restent conservés.
