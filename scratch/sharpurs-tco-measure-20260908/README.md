# Mesure du noyau TCO Int — 8 septembre 2026

Le noyau typé réduit TCO d’environ **5,6 ms à 0,055 ms**, soit **×100 à ×103 dans le benchmark complet**, confirmé dans deux paires de processus. Les résultats restent identiques. Le gain propre à TCO est d’environ **5,5 ms** sur le score total.

## Références et mesures

La [baseline officielle F#/C#](/Users/0x1/Documents/htdocs/altbak.pub/README.md:159) donne **184 283 µs** pour TCO. Les mesures locales avant l’étape 1 donnaient 93 463,88 µs ; l’appel direct intégré avait déjà ramené TCO à 5 334,83 puis 5 589,13 µs. La présente comparaison conserve cette optimisation d’appel des deux côtés et isole le remplacement du binding TCO par son noyau `int`.

Mesures détaillées : trois processus par variante, chacun avec 3 échauffements globaux d’`App`, 3 échauffements TCO et 10 mesures de l’action complète.

| Indicateur | Avant noyau Int | Après noyau Int |
| --- | ---: | ---: |
| Minimum des 10 passages, processus 1 | 5 487,000 µs | 49,666 µs |
| Minimum des 10 passages, processus 2 | 5 594,542 µs | 52,542 µs |
| Minimum des 10 passages, processus 3 | 5 563,500 µs | 48,875 µs |
| Médiane des trois minimums | 5 563,500 µs | 49,666 µs |
| Médiane des trois médianes | 5 649,751 µs | 50,813 µs |
| Allocation par action, sur chacun des 30 passages | 19 204 552 octets | 160 octets |
| Collections GC sur les 30 passages, générations 0 / 1 / 2 | 69 / 0 / 0 | 0 / 0 / 0 |

Le ratio des médianes des minimums est **×112,02** ; celui des médianes des médianes est **×111,19**. Les allocations sont cumulées par action sur le thread de travail, pas une mesure de mémoire résidente. Les 160 octets restants incluent les frontières du noyau et la production du résultat de l’action.

Confirmation par le véritable `App.main`, avec son harness original et les autres benchmarks exécutés dans leur ordre habituel :

| Paire | TCO avant | TCO après | Ratio | Total avant | Total après |
| --- | ---: | ---: | ---: | ---: | ---: |
| 1 | 5 596,83 µs | 55,75 µs | ×100,39 | 897,40 ms | 884,92 ms |
| 2 | 5 645,75 µs | 54,88 µs | ×102,87 | 886,71 ms | 853,12 ms |

Les 14 sorties affichées concordent avec les résultats de référence dans les quatre exécutions complètes. Le harness officiel conserve son minimum de dix mesures ; il n’effectue pas de validation après chaque passage. Les séries détaillées, elles, vérifient `100000` après chacun des 60 passages chronométrés et des 18 échauffements TCO locaux.

La variation des totaux dépasse les 5,5 ms économisées par TCO : elle inclut les fluctuations des autres tests. Il serait incorrect de l’attribuer entièrement au noyau Int. Dans la deuxième exécution après changement, Polymorphism prend 563,722 ms, RBTree 212,694 ms et Lazy 68,243 ms : **99,01 % du total**. La prochaine cible utile est donc le binding polymorphe `polyLoop`, à examiner dans le PBO optimisé avec ses instanciations TAST v3.

Les colonnes FFI historiques (247 et 200 µs) restent indicatives : elles accumulent `1`, tandis que le test PureScript accumule `n mod 3`. Même entrée opaque et même résultat ne constituent pas ici une charge arithmétique identique.

## Protocole et contrôle des variantes

- Deux copies du programme F# généré, compilées en Release, SDK .NET 8.0.423, runtime 8.0.29, macOS arm64. Le GC workstation et ses paramètres sont identiques ; aucun GC forcé ni réglage du tiering ajouté.
- `prepare.py` vérifie les empreintes du TCO et du prélude déjà validés, puis restaure uniquement l’ancien binding TCO dans la copie `before` en inversant son diff enregistré. **Un seul fichier sur les 353 entrées copiées diffère entre variantes : `Test.TCO.fs`.** Les préludes et le harness de mesure sont identiques.
- Le vrai `Test.TCO.act` est invoqué, avec `Bench.opaque 100000`, le wrapper public et `show`. Aucun appel du seul noyau nu ne remplace cette action.
- `MeasureEntryPoint.fs` utilise le même thread de travail avec une pile de 1 Gio que le programme normal. Les compteurs d’allocation entourent l’action et les lectures d’horloge ; la validation, les objets de rapport et leur sérialisation restent en dehors.
- Trois processus indépendants par variante, ordre **avant, après, après, avant, avant, après**. Tous les builds sont terminés avant les mesures ; aucune mesure n’est lancée en parallèle par ces scripts.
- Le diagnostic échauffe les mêmes fonctions partagées via 3 `App.warmup`, puis fait les 3 échauffements locaux TCO. Le programme complet exécute en plus les benchmarks AST/Fib/List avant TCO : les séries complètes sont donc conservées comme confirmation distincte.
- Les sources générées habituelles sont inchangées après cette mesure, vérifiées par empreintes.

Une première expérience avec seulement 6 échauffements TCO isolés donnait environ 14,3 ms avant et 47,5 µs après. Son résultat « avant » différait nettement de celui d’`App.main`. L’ajout des échauffements globaux réels ramène la mesure détaillée à environ 5,6 ms. Cela établit l’importance du contexte d’échauffement ; la cause précise n’a pas été tracée. Les données initiales et leur harness exact sont conservés dans `isolated/` et ne servent pas au ratio principal.

## Reproduction

Depuis ce dossier, avec les entrées générées dont les empreintes sont enregistrées :

```sh
python3 prepare.py
(cd before && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
(cd after && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
python3 run.py
python3 analyze.py
```

`run.py` conserve les 10 temps individuels, allocations et compteurs GC de chaque processus dans les fichiers `.jsonl`, puis lance une paire complète officielle. `analyze.py` vérifie les résultats et produit `comparison.json` ; la première paire officielle conservée sous `isolated/` reste une observation historique et ne sera pas relancée par cette commande.

`inputs.json` contient les empreintes des 353 entrées, du TAST TCO, du README officiel et du harness. `executions.json` conserve l’ordre d’exécution, les durées murales et les empreintes des DLL mesurées ; `metadata.json` contient les révisions et paramètres du runtime. Les dossiers compilés `before/` et `after/` sont ignorés par Git et reproductibles.
