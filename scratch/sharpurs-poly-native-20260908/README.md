# Boucle Polymorphism native — 8 septembre 2026

## Résultat

Sur cinq processus par variante du véritable `App.main`, le total médian passe de **875,35 à 296,39 ms** (**−66,14 %, ×2,95**). Polymorphism passe de **577,58467 à 2,22158 ms**, soit **×260**. Les 140 sorties affichées concordent avec les références.

| Mesure officielle | Avant | Après |
| --- | ---: | ---: |
| Total, processus 1 | 863,58 ms | 294,10 ms |
| Total, processus 2 | 875,35 ms | 292,20 ms |
| Total, processus 3 | 854,56 ms | 298,73 ms |
| Total, processus 4 | 878,38 ms | 307,77 ms |
| Total, processus 5 | 886,40 ms | 296,39 ms |
| Médiane totale | **875,35 ms** | **296,39 ms** |
| Médiane Polymorphism | **577,58467 ms** | **2,22158 ms** |

Le diagnostic sur trois processus par variante confirme le gain de la vraie action et la suppression des allocations dans la boucle :

| Diagnostic Polymorphism | Avant | Après |
| --- | ---: | ---: |
| Médiane des trois minimums | 603,64150 ms | 2,22750 ms |
| Médiane des trois médianes | 636,68637 ms | 2,34581 ms |
| Octets alloués par action | 1 440 000 424 | 112 |
| GC Gen0 / Gen1 / Gen2, sur 30 mesures | 5 166 / 0 / 0 | 0 / 0 / 0 |

Les 60 résultats mesurés et 18 échauffements locaux valent tous `10000000`. Le diagnostic diffère du contexte officiel : après les échauffements globaux, il lance directement l’action cible, tandis qu’`App.main` exécute les tests précédents. Ses temps ne remplacent donc pas les scores officiels. Aucune fluctuation des autres tests n’est attribuée au noyau : le ratio propre à Polymorphism est calculé séparément du total.

La médiane TCO reste à **0,05158 ms**. Les médianes résiduelles sont **215,28271 ms pour RBTree** et **69,10621 ms pour Lazy** ; elles deviennent les principales cibles. La comparaison historique au README donne environ **×47 sur le total**, gains des étapes précédentes compris. [Statistiques et validations complètes](comparison.json), [résumé](analyze.log).

## Changement et contrôles

Sharpurs consomme le `LetRec` optimisé, dont le TAST porte maintenant une signature entièrement `Int`. `IntKernel.fromLocal` accepte une seule fonction locale, fermée sur ses paramètres, avec un résultat `Int`, des primitives entières prises en charge et une récursion terminale saturée. L’appel initial peut utiliser des constantes et des paramètres externes dont le type `Int` est prouvé. Les niveaux et l’identité de la référence récursive sont vérifiés.

`Sharpurs.Optimized` conserve une enveloppe `obj` pour les références, applications et lambdas autour du noyau. Le déboxage est fait à l’entrée de la boucle et le résultat est boxé à sa sortie. Les effets passent par les appels existants : la boucle reste dans la continuation exécutée après `Bench.opaque`. Aucun nom de benchmark, nombre d’itérations ou dictionnaire particulier ne déclenche la reconnaissance.

L’enveloppe ne remplace que les bindings non récursifs entièrement pris en charge et contenant au moins un noyau local. Toute forme refusée conserve le générateur CoreFn actuel. Le noyau global existant garde la priorité et les groupes mutuellement récursifs restent intacts. Ce jalon ne généralise ni les captures des boucles locales, ni les dictionnaires dynamiques, ni les ADT.

La régénération du vrai programme change **un seul fichier sur 355**, `Test.Polymorphism.fs`, et seulement son binding `act`. La fonction générique `polyLoop` et le noyau TCO sont inchangés. [Diff émis](Test.Polymorphism.diff), [nouveau module](Test.Polymorphism.generated.fs), [version de référence](Test.Polymorphism.before.fs), [empreintes](inputs.json).

Vérifications conservées :

- [local-kernel-tests.log](local-kernel-tests.log) : 68 assertions de conversion/routage et 43 assertions F# ; types contradictoires, portées, repli complet, récursion terminale, application partielle réutilisable, overflow, million d’itérations, effets différés/répétés, ordre des appels et exception.
- [converter-tests.log](converter-tests.log) : 43 assertions IntKernel et 6 assertions de routage existantes.
- [global-kernel-tests.log](global-kernel-tests.log) : 695 vérifications F# du noyau global existant.
- [runtime-tests.log](runtime-tests.log) : 21 vérifications du helper d’application et des exceptions.
- [generic-checks.log](generic-checks.log) : 6 vérifications sur la DLL réellement générée, avec appels génériques en `Int` et `String`, puis appels répétés de l’action native. Le dictionnaire String reste dynamique et les représentations cohabitent correctement. [Fixture](VerifyGeneric.fsx).
- Les quatre builds F# Release passent sans erreur ni avertissement. Le build PureScript conserve les avertissements existants du backend.

## Mesure

Deux copies du code réellement généré sont compilées avec .NET 8.0.423 / runtime 8.0.29, en Release. Les **355 entrées ne diffèrent que par `Test.Polymorphism.act`**. `before` et `after` conservent le vrai `EntryPoint.fs` et exécutent `App.main` sans modification : 3 échauffements globaux, 3 locaux, puis minimum de 10 temps par test. Cinq processus indépendants par variante, ordre **avant, après, après, avant, avant, après, après, avant, avant, après**. Les 14 résultats affichés sont contrôlés contre l’historique.

Les copies `diagnostic-before` et `diagnostic-after` remplacent uniquement le point d’entrée par [MeasureEntryPoint.fs](MeasureEntryPoint.fs). Celui-ci échauffe 3 fois `App.warmup`, puis 3 fois la vraie `Test.Polymorphism.act`, et conserve ses 10 temps individuels, allocations du thread courant et compteurs GC. Les sorties sont vérifiées hors chronométrage à chaque appel local. Trois processus par variante, ordre **avant, après, après, avant, avant, après**. Le thread de travail garde une pile de 1 Gio ; aucun GC forcé ni réglage de tiering/PGO n’est introduit.

Les builds et tests sont finis avant les mesures ; aucun benchmark concurrent n’est lancé. Le GC workstation est vérifié dans les diagnostics et les configurations runtime sont identiques. Les allocations sont des octets cumulés par action, pas la mémoire résidente ; les compteurs GC portent sur le processus. Les échauffements globaux ne publient pas leurs sorties. Le harness officiel ne publie que le meilleur des dix temps, tandis que le diagnostic conserve les dix observations.

La baseline officielle [README F#/C#](/Users/0x1/Documents/htdocs/altbak.pub/README.md:152) est **9 689,415 ms** pour Polymorphism et **13 918,86 ms** au total. Elle mesure l’écart historique cumulé ; la comparaison avant/après présente isole l’émission native et son enveloppe optimisée.

## Reproduction

Le répertoire de travail normal contient le programme après génération. `prepare.py` reconstruit la référence à partir des autres sources identiques et de `Test.Polymorphism.before.fs`, puis vérifie toutes les empreintes. Les copies et produits compilés sont ignorés par Git.

```sh
# Depuis sharpurs/sharpurs :
npm run build
spago test
DOTNET=/Users/0x1/.dotnet/dotnet npm run test:local-kernel
DOTNET=/Users/0x1/.dotnet/dotnet npm run test:kernel
DOTNET=/Users/0x1/.dotnet/dotnet npm run test:runtime

# Depuis altbak.pub-sharpurs, avec son output TAST actuel :
../sharpurs/sharpurs/bin/sharpurs --main App

# Depuis ce dossier :
python3 prepare.py
(cd before && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
(cd after && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
(cd diagnostic-before && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
(cd diagnostic-after && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
python3 run.py
python3 analyze.py
/Users/0x1/.dotnet/dotnet fsi --nologo --exec VerifyGeneric.fsx
```

[source-metadata.json](source-metadata.json) identifie les sources modifiées, le bundle utilisé et l’entrée TAST ; les seuls commits ne suffisent pas à identifier des changements non commités. [executions.json](executions.json) conserve ordre, DLL mesurées, heures et charges système. Les configurations sont dans [metadata.json](metadata.json), les statistiques et validations dans [comparison.json](comparison.json).
