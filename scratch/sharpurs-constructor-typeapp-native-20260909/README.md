# Constructeurs saturés sous TypeApp — intégration Sharpurs

La génération normale après intégration réduit Primes de **4,09400 à 0,80863 ms** (−80,25 %) et List Processing de **0,82346 à 0,44058 ms** (−46,50 %). Le total médian passe de **59,59 à 54,68 ms** (−8,24 %). Les 140 sorties des dix processus sont identiques à la référence historique.

| Mesure, en ms | Avant, médiane [min ; max] | Après, médiane [min ; max] | Variation |
| --- | --- | --- | --- |
| Primes | 4,09400 [3,85942 ; 4,32608] | 0,80863 [0,74971 ; 0,84763] | −80,25 % |
| List Processing | 0,82346 [0,75333 ; 0,87988] | 0,44058 [0,41329 ; 0,44429] | −46,50 % |
| Total | 59,59 [57,68 ; 59,94] | 54,68 [54,32 ; 55,03] | −8,24 % |
| RBTree | 38,52258 [37,63292 ; 39,29004] | 38,41175 [38,06054 ; 38,61879] | −0,29 % |
| Lazy | 5,78713 [5,66417 ; 6,11908] | 5,88779 [5,68046 ; 6,07879] | +1,74 % |

Les sources RBTree et Lazy ne changent pas ; leurs petites variations ne sont pas attribuées à cette optimisation. Les médianes de chaque test ne s'additionnent pas nécessairement au total médian. Toutes les valeurs des quatorze tests figurent dans `comparison.json`.

Le [README officiel figé](README.baseline.md) donne 60,41 ms au total, 3,86179 ms pour Primes et 0,74763 ms pour List Processing. Il s'agit d'une référence historique, pas du témoin simultané de cette série. La dernière colonne native donne 0,119 ms pour Primes avec un algorithme différent ; ce chiffre reste un repère, sans isoler le coût des constructeurs. Elle donne aussi 0,077 ms pour List Processing. Le prototype isolé antérieur à l'intégration donnait 3,754945 → 0,838130 ms pour Primes : ces chiffres sont distincts de la série finale ci-dessus.

La mesure exécute le vrai `App.main`, avec son échauffement et son meilleur temps parmi dix exécutions par test. Les copies Release démarrent chacune dans cinq processus neufs, séquentiellement, dans l'ordre **ABBAABBAAB** (`A=before`, `B=after`). Aucun build ou test concurrent de l'équipe durant ces dix processus. Aucun relevé de RAM ou de GC. Les statistiques sont descriptives, sur une seule machine ; l'amélioration de Primes et List Processing présente ici des étendues disjointes.

L'après provient entièrement de la régénération normale avec le backend intégré, sans correction manuelle des fichiers F#. Sur les **355 entrées sources/projets**, **80 fichiers F# changent** et 275 entrées restent identiques. L'inventaire compte **1 580 occurrences textuelles de `sharpurs_apply` en moins** ; ce nombre statique n'est pas un comptage des appels à l'exécution. Les fichiers de benchmark directement modifiés sont `Test.Primes.fs` (33 → 25 occurrences) et `Test.ListOps.fs` (32 → 28), auxquels s'ajoutent les modules de bibliothèque. L'optimisation reconnaît les constructeurs saturés à travers les applications de types du TAST ; elle conserve les chemins génériques lorsque le cas ne permet pas une construction directe.

Les copies avant/après et la génération normale compilent en Release avec zéro avertissement et zéro erreur .NET, respectivement en 30,60 s, 31,51 s et 31,81 s. Le build du backend consigné dans `backend-build.log` comporte 14 avertissements préexistants. Les validations sont conservées dans ce dossier :

- Nouvelle fixture : **133 assertions F# + 62 contrôles converter/JS**, zéro avertissement, avec cas importé polymorphe et `TypeApp` intercalé dans une chaîne d'applications.
- Onze suites de régression du backend passent, ainsi que les **49 assertions de Test.Main** (`regression-*.log`).
- Chaque DLL avant/après passe **150 contrôles Tuple** (construction, projections, `swap`, double `swap`, réutilisation d'une application partielle, identité des valeurs transportées) et **165 contrôles CodePoints**.
- Les DLL après et normale passent chacune les **333 contrôles Lazy** existants.
- Le benchmark valide les noms, l'ordre et les valeurs des **140 résultats**. Les empreintes des 355 entrées de chaque programme et des binaires sont contrôlées durant la série.

La baseline est celle du programme courant. Depuis l'audit thunk précédent, `AppFFI.fs` a adopté le cumul des timings et `Program.fsproj` a changé l'ordre de ses entrées ; les 353 autres entrées sont identiques. Ces deux fichiers courants ont été figés avant la régénération et restent identiques entre avant/après. Le snapshot des sources backend avant a été pris pendant le début du travail d'intégration : il peut déjà refléter une modification en cours. Les sources générées, le bundle avant et les binaires normaux avant sont en revanche capturés séparément et leur provenance est explicitée dans `inputs.json`.

`inputs.json` conserve les révisions, les empreintes des 355 entrées et les bundles avant/après. `build-{before,after,normal}.json` associe les entrées aux huit fichiers runtime de chaque build. `normal-before-runtime/` garde les huit fichiers runtime normaux initiaux. `generation-inventory.json`, les `.diff` et `.generated` montrent les changements ; les copies intégrales et binaires volumineux sont exclus par le `.gitignore` local. `integrity-final.json` atteste le contrôle final des trois programmes, de la copie runtime initiale et du backend actuel. `validation-files.json` fige les scripts et logs de validation.

Pour relire les résultats sans exécuter le programme :

```sh
python3 analyze.py
python3 verify_integrity.py
```

`run.py` lance la série complète sur les deux copies déjà compilées ; il refuse d'écraser une série existante. Les builds ont utilisé `/Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo -p:NuGetAudit=false` depuis chaque répertoire. Les scripts `capture_build.py`, `capture_normal_build.py` et `snapshot_after.py` conservent la provenance des constructions et de la génération normale.
