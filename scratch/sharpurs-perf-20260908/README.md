# Audit de performance Sharpurs — 8 septembre 2026

Objectif : identifier des améliorations du runtime et préparer le plan de `sharpurs/todo.md`. Aucun fichier source du backend, de PBO ou du benchmark n’a été modifié par cet audit. La copie `direct-call/` et ses produits de build sont ignorés par Git.

## Résultat du test contrôlé

Seule la fonction `sharpurs_apply` change dans la copie générée : si la valeur est une closure `obj -> obj`, appel direct ; sinon, conservation du code de réflexion initial. Voir [le patch](direct-call.patch).

| Score officiel (somme des 14 best-of-10) | Temps |
| --- | ---: |
| Baseline historique README altbak.pub | 13 918,860 ms |
| Programme original local | 13 766,970 ms |
| Copie avec fast path | 849,810 ms |

Gain expérimental global : ×16,20, soit −93,83 %. Les 14 sorties affichées ont été extraites et comparées, toutes identiques. Le protocole original ne vérifie pas les dix résultats chronométrés ; le harness complémentaire ci-dessous les vérifie pour les trois cas prioritaires.

Temps muraux des exécutables déjà compilés : 227,73 s et 14,19 s. Ce ne sont pas des temps de `run -c`. La compilation isolée de la copie a pris 31,78 s, mesurée séparément ; aucun gain de compilation n’est revendiqué. Build réussi avec zéro erreur et zéro avertissement MSBuild ; la plateforme a également émis une ligne CSSM dans stderr.

## Confirmation indépendante sur les trois actions originales

Harness `measure-actions.fsx`, processus séparé par variante : trois échauffements, cinq passages mesurés, entrées originales, résultats vérifiés à chaque passage. Un thread de pile 1 Gio reproduit la convention de l’EntryPoint généré.

| Cas | Médiane originale | Médiane fast path | Allocations originales / passage | Allocations fast path / passage |
| --- | ---: | ---: | ---: | ---: |
| Polymorphism | 11 118,265 ms | 581,918 ms | ≈31,840 Go | ≈1,440 Go |
| RBTree | 2 613,527 ms | 217,761 ms | ≈7,927 Go | ≈0,740 Go |
| LazyEvaluation | 985,999 ms | 53,460 ms | ≈2,908 Go | ≈0,169 Go |

Go décimaux : octets cumulativement alloués par le thread pendant une action, pas mémoire résidente. Le JSON conserve les valeurs exactes et les compteurs GC. Les sorties attendues et vérifiées sont respectivement `10000000`, `22`, `1000000`.

Cette série est diagnostique : F# Interactive utilise le GC serveur (`FSharp/fsi.runtimeconfig.json`), contrairement au programme officiel dont le runtimeconfig ne l’active pas. Les comparaisons de cette table sont entre les deux variantes du même harness ; elles ne remplacent pas le score officiel. Les compteurs GC sont ceux du processus entier, les octets ceux du thread de travail. La résolution de l’action a lieu avant la mesure ; un unique appel réfléchi mis en cache entoure chaque action, identiquement dans les deux variantes.

## Reproduction

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-sharpurs`, avec les sources et versions consignées dans `metadata.json` :

```sh
# Générer la copie depuis output/Main, puis compiler séparément.
python3 scratch/sharpurs-perf-20260908/prepare-probe.py
cd scratch/sharpurs-perf-20260908/direct-call
"$HOME/.dotnet/dotnet" build Program.fsproj -c Release -p:NuGetAudit=false --ignore-failed-sources
```

Exécuter séquentiellement les programmes, sans build concurrent :

```sh
"$HOME/.dotnet/dotnet" /Users/0x1/Documents/htdocs/altbak.pub-sharpurs/output/Main/bin/Release/net8.0/Program.dll
"$HOME/.dotnet/dotnet" /Users/0x1/Documents/htdocs/altbak.pub-sharpurs/scratch/sharpurs-perf-20260908/direct-call/bin/Release/net8.0/Program.dll
```

Depuis ce dossier d’audit, confirmation avec allocations :

```sh
"$HOME/.dotnet/dotnet" fsi --exec measure-actions.fsx /Users/0x1/Documents/htdocs/altbak.pub-sharpurs/output/Main/bin/Release/net8.0/Program.dll
"$HOME/.dotnet/dotnet" fsi --exec measure-actions.fsx /Users/0x1/Documents/htdocs/altbak.pub-sharpurs/scratch/sharpurs-perf-20260908/direct-call/bin/Release/net8.0/Program.dll
```

Un argument facultatif `Polymorphism`, `RBTree` ou `LazyEvaluation` limite le harness. L’original vérifié par `input-sha256.json` est resté inchangé pendant l’expérience. Pour reproduire après d’autres modifications, reconstruire les révisions enregistrées et contrôler les empreintes ; ne pas supposer qu’un ancien cache représente le nouveau code.

## Pièces conservées

- `baseline-full.log`, `direct-call-full.log`, leurs `.time` et `full-comparison.json` : score et sorties des 14 tests.
- `baseline-actions.jsonl`, `direct-call-actions.jsonl`, leurs `.time` et `actions-comparison.json` : passages individuels, médianes, allocations et GC.
- `direct-call-build.log`, `direct-call-build.time`, `direct-call.patch`, `prepare-probe.py` : expérience isolée.
- `metadata.json`, `input-sha256.json` : versions et identité des entrées.
- `tast-evidence.json` : déclarations et types effectifs pour TCO, Polymorphism, RBTree et Lazy, avec exemples de `TypeApp`. Ceux-ci existent bien dans le JSON et ne sont pas déduits du F#.

## Limites et suite

Le test valide une réduction du coût de dispatch sur le corpus pur actuel. Avant intégration, vérifier les représentations curriées, les FFI et les exceptions : un appel direct peut supprimer une enveloppe `TargetInvocationException`. Le plan détaillé est dans `/Users/0x1/Documents/htdocs/sharpurs/todo.md`.

Le protocole et les scores du README historique restent la référence datée. Les colonnes FFI comportent des différences de charge ou de résultat, notamment Lazy et RBTree ; aucune convergence chiffrée vers ces colonnes n’est promise ici.

La fonction F# est représentée par `FSharpFunc<T,U>`, qui dispose d’un appel `Invoke` typé : [documentation FSharp.Core](https://fsharp.github.io/fsharp-core-docs/reference/fsharp-core-fsharpfunc-2.html). Microsoft documente aussi les différences entre réflexion répétée et invocation typée : [billet .NET 7, section réflexion](https://devblogs.microsoft.com/dotnet/performance_improvements_in_net_7/). Les gains annoncés ici viennent des mesures locales, pas de ces références générales.
