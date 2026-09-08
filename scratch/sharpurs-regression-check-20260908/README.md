# Vérification du passage à 920,64 ms — 8 septembre 2026

## Résultat

La médiane actuelle est **915,20 ms**, contre **900,70 ms** pour la reconstruction exacte de la référence, soit **+14,50 ms / +1,61 %**. L’actuel présente davantage de variabilité et descend aussi sous 900 ms. La série ne permet pas d’attribuer le petit écart médian aux deux modules ajoutés ni de conclure à une régression stable. Elle établit que la référence ancienne dépasse elle aussi 900 ms dans les conditions présentes.

| Processus par variante | Référence, ms | Actuel, ms |
| --- | ---: | ---: |
| 1 | 892,85 | 886,37 |
| 2 | 907,38 | 915,20 |
| 3 | 899,29 | 874,47 |
| 4 | 905,57 | 936,54 |
| 5 | 900,70 | 927,27 |
| Médiane | **900,70** | **915,20** |
| Étendue min–max | 892,85–907,38 | 874,47–936,54 |

Les valeurs de chaque test sont les médianes de leurs cinq scores best-of-10 :

| Test | Référence, ms | Actuel, ms | Écart |
| --- | ---: | ---: | ---: |
| Polymorphism | 595,17213 | 609,25825 | +2,37 % |
| RBTree | 224,61546 | 226,39121 | +0,79 % |
| Lazy | 71,66671 | 70,86613 | −1,12 % |
| TCO | 0,04946 | 0,05200 | +0,00254 ms |

L’écart médian est surtout visible sur Polymorphism, tandis que la boucle TCO reste à environ 50 µs. Les médianes par test ne s’additionnent pas nécessairement à la médiane des totaux. Les **140 sorties affichées** concordent avec les références, et les sommes des quatorze temps concordent avec les totaux arrondis. Les deux builds sont sans erreur ni avertissement. Les sources générées de travail sont restées inchangées.

Le passage utilisateur à 920,64 ms appartient à l’étendue observée de l’actuel. Les anciens 853,12 et 884,92 ms restent des observations historiques ; ce protocole ne reproduit pas leur état de machine. La baseline officielle de 13 918,86 ms donne un gain cumulé d’environ ×15,21, sans trancher la petite différence entre versions récentes.

Le correctif PBO `TypeApp` n’a changé aucun des fichiers F#/C# exécutés par les benchmarks. Aucun correctif de performance supplémentaire n’est appliqué à partir de cette série. Les chiffres sont descriptifs : cinq processus par variante et des étendues qui se recouvrent ne démontrent ni une équivalence parfaite ni une cause précise. [Analyse complète](comparison.json), [validation](analyze.log), [progression des exécutions](run.log).

## Protocole

Le signal utilisateur est un total de 920,64 ms, après des mesures locales de 884,92 et 853,12 ms. Les fichiers générés sont comparés aux empreintes conservées dans `sharpurs-tco-measure-20260908/inputs.json` : aucun code F#/C# de benchmark n’a changé. Le noyau TCO natif est identique. Le seul fichier commun différent est `Program.fsproj`, qui inclut désormais `AppJavaFFI.fs` et `AppJavaFFICheatcode.fs` ; ce sont aussi les deux seules sources ajoutées.

`prepare.py` produit deux copies isolées :

- `reference` : les **353 fichiers**, projet et point d’entrée compris, correspondent exactement aux empreintes de la variante rapide enregistrée après le noyau TCO.
- `current` : copie des **355 fichiers** générés actuels. La différence du projet est conservée dans [project.diff](project.diff).

Les deux programmes sont compilés avec le même SDK .NET 8.0.423, en Release, puis exécutés avec le runtime 8.0.29. Les builds finissent avant la première mesure. Les configurations runtime sont identiques, sans modification du tiering, du PGO ou du GC. Aucun changement des backends, du prélude, du noyau TCO ou du code généré habituel n’est effectué.

Le véritable `App.main` est exécuté dans cinq processus frais par variante, un seul à la fois. Ordre : **R C C R R C C R R C**. Chaque processus conserve les trois échauffements globaux, les trois échauffements locaux et le minimum de dix passages par test. Le total affiché est la somme des quatorze minimums ; le temps de build et de démarrage n’y entre pas. Les comparaisons entre processus utilisent les médianes et présentent aussi les étendues.

Les quatorze sorties affichées de chaque processus sont vérifiées contre les deux logs historiques. Le harness original ne publie ni les dix temps bruts ni le résultat de chaque passage chronométré ; cette mesure conserve ce protocole pour comparer les totaux officiels. Elle ne remplace pas le diagnostic séparé des allocations.

La [baseline officielle du README](/Users/0x1/Documents/htdocs/altbak.pub/README.md:152), **13 918,86 ms**, est conservée comme repère historique des gains cumulés. Le témoin pertinent pour la régression soupçonnée est ici la reconstruction exacte du programme récent.

## Fichiers et reproduction

Les empreintes sont dans [inputs.json](inputs.json). Les versions et paramètres figurent dans [metadata.json](metadata.json). [executions.json](executions.json) contient l’ordre, les empreintes des DLL, les dates, durées murales et charges système avant/après chaque processus. Les logs complets portent les noms `reference-1.log` à `reference-5.log` et `current-1.log` à `current-5.log`.

```sh
python3 prepare.py
(cd reference && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
(cd current && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
python3 run.py
python3 analyze.py
```

`prepare.py` refuse de reconstruire la référence si les empreintes ne correspondent plus aux entrées enregistrées. Les copies et les produits de compilation sont ignorés par Git. La machine reste un poste de travail actif ; aucune activité utilisateur n’est arrêtée. Les résultats décrivent cette série et ne constituent pas une preuve d’équivalence absolue.
