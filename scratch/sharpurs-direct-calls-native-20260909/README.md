# Appels directs intégrés au générateur Sharpurs

Audit de l’intégration au compilateur de la règle générale d’appels saturés directs. La mesure porte sur les sorties normales du générateur ; aucune variante F# n’est réécrite manuellement dans cet audit.

Le backend part du commit propre `310a09b48416d0bccf2ed9ab077a8323f0cd1758`.

La reconnaissance est dans [DirectCall.purs](/Users/0x1/Documents/htdocs/sharpurs/sharpurs/src/Sharpurs/DirectCall.purs), la sélection et l’émission dans [CodeGen.purs](/Users/0x1/Documents/htdocs/sharpurs/sharpurs/src/Sharpurs/CodeGen.purs). La nouvelle suite est [tests/direct-call.mjs](/Users/0x1/Documents/htdocs/sharpurs/sharpurs/tests/direct-call.mjs), disponible via `npm run test:direct-call` après reconstruction du backend, avec `PURS` pointant vers le fork TAST et `DOTNET` vers le runtime .NET.

L’intégration confirme le gain : **219,95 → 182,44 ms au total (−17,05 %)** et **138,04950 → 97,79479 ms pour RBTree (−29,16 %)**, médianes de cinq processus par variante. Les 140 sorties sont conformes et tous les hashes de sources/runtime sont restés identiques durant la série.

| Mesure | Avant : médiane [min ; max] | Après : médiane [min ; max] | Évolution |
| --- | ---: | ---: | ---: |
| Total | 219,95 [216,58 ; 228,68] ms | 182,44 [173,84 ; 184,23] ms | −17,05 % |
| RBTree | 138,04950 [135,72363 ; 143,05104] ms | 97,79479 [91,53483 ; 99,17942] ms | −29,16 % |

La paire préliminaire donnait 213,97 → 172,97 ms au total et 133,33833 → 90,82950 ms pour RBTree. Elle est conservée séparément et exclue des médianes finales. L’écart entre les processus figure dans les étendues ; les statistiques descriptives ne sont pas une promesse d’un temps fixe à chaque lancement.

## Changement effectivement mesuré

Le générateur reconnaît les fonctions non récursives à plusieurs arguments, avec des lambdas consécutives et une signature fermée vérifiée dans le TAST. Leurs paramètres et leur résultat doivent être des `Int`, `Boolean` ou des types ADT acceptés par le layout natif. Les applications sont dirigées vers le helper uniquement si elles sont qualifiées, exactement saturées et cohérentes avec toutes les annotations intermédiaires. Le chemin générique reste utilisé pour les autres formes, notamment les applications partielles et les `TypeApp`.

Le corps conserve ses opérations et ses représentations `obj`. Celui de `balance` est identique octet par octet entre les deux générations (`source-checks.json`). Un helper à plusieurs paramètres exécute ce corps, le wrapper public conserve ses étapes curryfiées, et les sites directs passent par `_direct_apply`, qui reproduit l’enveloppe `TargetInvocationException` de `sharpurs_apply`.

Les 355 entrées de build ne diffèrent que dans deux modules. `direct-call-sites.json` recense cinq sites : deux appels de `balance` et un appel de `insert` dans RBTree, deux appels de `unsurrogate` dans `Data.String.CodePoints`. `max` reçoit aussi un helper, sans appel direct correspondant. Ce résultat provient de la règle générale du compilateur ; le générateur ne cible aucun de ces noms.

La génération normale a aussi été compilée en Release après intégration. Ses 355 entrées correspondent exactement à celles de la variante mesurée `after`, et `build-normal.json` conserve ses hashes runtime ainsi que ceux du backend et de son bundle.

## Validation

- Les builds .NET avant, après et normal réussissent avec zéro avertissement et zéro erreur.
- La vraie DLL après intégration passe 27 087 vérifications RBTree, 138 vérifications de `balance` et 165 vérifications CodePoints, soit 27 390 vérifications ciblées. Cela couvre notamment l’ABI publique partielle, les rotations et le partage, l’ordre d’évaluation et les exceptions du helper direct, ainsi que les deux sites appelants Unicode.
- La nouvelle fixture du backend rapporte 130 vérifications runtime et 67 vérifications du convertisseur. `sharpurs-direct-test-direct-call.log` est explicitement un relevé des résultats observés, accompagné des trois artefacts F# de `direct-call-tests/` ; ce fichier n’est pas présenté comme un journal brut.
- Les journaux conservés des suites existantes passent : runtime apply (21), IntKernel (695), local IntKernel (43 runtime + 68 convertisseur), ADT kernel (32 + 33), ADT interop (62 + 13), ADT unary (63 + 11), Int comparisons (396 + 32), et `spago test` (43 + 6).
- Le build PureScript du backend conserve ses avertissements existants ; les journaux PureScript sont inclus. La mention « zéro avertissement » ci-dessus concerne les builds .NET, pas l’ensemble des builds de ce tour.

Les résultats détaillés, les 10 journaux, leur ordre et leurs hashes sont dans `comparison.json`, `executions.json` et `metadata.json`.

## Protocole

- `before/` contient les 355 entrées de build du générateur normal avant cette intégration, avec les comparaisons Int natives déjà intégrées. Le témoin a été copié avant toute régénération.
- Ses 355 hashes sont identiques au témoin de `../sharpurs-rbt-balance-direct-20260909` ; l’expérience précédente n’a donc pas remplacé les sources normales.
- `capture_after.py` a copié les entrées du générateur normal après intégration dans `after/`. Son manifeste est identique à celui de `run/bak/sharp/output/Main`.
- Deux builds Release indépendants, `dotnet build Program.fsproj -c Release -p:NuGetAudit=false`, produisent les journaux `build-before.log` et `build-after.log`. `capture_build.py` relie chaque ensemble de sources à ses fichiers runtime et son journal par SHA-256.
- `VerifyTree.fsx` contrôle les invariants de l’arbre, sa profondeur, sa persistance et les ponts d’ABI sur la génération réelle. `VerifyBalance.fsx` contrôle les rotations, les sous-arbres partagés, les applications partielles, le report d’exécution du corps et l’ordre d’évaluation.
- `VerifyCodePoints.fsx` vérifie les conversions de 25 couples de surrogates, les applications partielles réutilisées, puis les deux fonctions appelantes `uncons` et `unsafeCodePointAt0Fallback` sur des chaînes Unicode et des surrogates isolés.
- Aucun build ni test n’accompagne les mesures. `run.py --probe` exécute une paire préliminaire. Ensuite `run.py` lance cinq processus neufs par variante, séquentiellement dans l’ordre ABBAABBAAB.
- Chaque processus appelle `App.main` sans modification : meilleur de 10 pour chacun des 14 tests. Le runner vérifie les noms, l’ordre et les sorties contre `expected-output.log`, référence conservée dans les audits précédents. Les 10 processus donnent 140 sorties à vérifier.
- `analyze.py` présente les médianes et les étendues des mesures par processus. La médiane de la somme n’est pas nécessairement la somme des médianes. La paire préliminaire ne fonde pas la conclusion finale.
- Les sources, la génération normale, les configurations runtime, les DLL, les journaux de build et les références sont vérifiés par hash avant et après chaque processus.

Le critère de gain est le temps d’exécution. Cet audit ne mesure ni allocations ni GC.

## Repères officiels

`README.baseline.md` est une copie du README de `/Users/0x1/Documents/htdocs/altbak.pub` au début de ce tour : 216,81 ms au total, 135,882 ms pour RBTree compilé et 65,156 ms dans la dernière colonne native optimisée. Ce README a été actualisé depuis l’expérience précédente, qui utilisait 294,80 ms et 213,760 ms. Le garde-fou a détecté cette évolution avant toute mesure et l’analyse utilise les chiffres officiels actuels.

Ces chiffres historiques arrondis donnent du contexte ; seul le témoin A/B simultané permet d’attribuer l’effet à cette intégration. Le RBTree natif utilise un `SortedSet<int>` mutable, insère les valeurs dans l’ordre croissant et renvoie sa taille ; le compilé utilise un arbre persistant, insère dans l’ordre décroissant et renvoie sa profondeur. L’écart au natif ne mesure donc pas uniquement le coût du code généré.

Par rapport au README officiel actuel, la médiane après intégration est inférieure de 15,85 % au total et de 28,03 % pour RBTree. Les gains attribués au changement restent ceux du tableau A/B : −17,05 % et −29,16 %.
