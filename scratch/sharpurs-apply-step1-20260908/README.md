# Étape 1 : appel direct des closures Sharpurs

Intégration du 8 septembre 2026. Seule l’application générique change ; aucune transformation TAST ni modification du protocole de benchmark.

## Changement livré

`sharpurs/sharpurs/src/Main.purs` émet un appel direct pour les valeurs `obj -> obj`. Les autres fonctions gardent le chemin réfléchi d’origine. Le contrôle de valeur nulle reste avant l’appel. Le chemin direct recrée une `TargetInvocationException` seulement si la fonction lève une exception, ce qui conserve type, message, enveloppes imbriquées et identité de la cause. Les frames des stacks reflètent désormais l’appel direct.

Le test `npm run test:runtime` extrait le vrai `fsPrelude` et utilise les vrais wrappers FFI et `Effect.Exception`. Ses 21 assertions couvrent les closures, retours de fonction boxés ou non, captures, applications partielles, fallback typé et délégué .NET, effets, récupération et exceptions imbriquées/de handler. Il crée puis supprime son script F# temporaire.

## Validation et performances

- Backend reconstruit, bundle vérifié, puis vraie commande `bin/sharp/run -c` exécutée depuis le worktree altbak.pub-sharpurs.
- Spago et .NET : builds réussis, zéro erreur et zéro avertissement.
- Les 14 sorties affichées sont identiques à celles de l’audit initial, aussi lors d’une seconde exécution indépendante du programme compilé.
- Score historique README altbak.pub : **13 918,86 ms**.
- Score avant modification, mesuré pendant l’audit : **13 766,97 ms**.
- Après régénération complète : **856,55 ms** (×16,07 face à l’avant local).
- Confirmation sans reconstruction : **869,22 ms** (≈×15,84).

Ces scores sont des sommes de minimums best-of-10. Le temps mural complet de `run -c` est **36,74 s**, compilation comprise. Celui de la confirmation du programme déjà compilé est **14,61 s**, échauffements et répétitions compris.

| Cas | Avant local | Après run -c |
| --- | ---: | ---: |
| Polymorphism | 9 853,619 ms | 561,211 ms |
| RBTree | 2 489,231 ms | 214,098 ms |
| LazyEvaluation | 1 220,149 ms | 67,636 ms |

Environnement inchangé depuis l’audit : macOS arm64, SDK .NET 8.0.423, runtime 8.0.29, Release net8.0. Les entrées App/Bench et les types TAST ne sont pas modifiés.

## Comptage diagnostique séparé

Une copie des sources générées ajoute uniquement deux compteurs et un EntryPoint limité aux trois actions. Un passage original par cas, sortie vérifiée, initialisations des modules hors comptage. Aucune mesure de performance ne vient de cette version instrumentée.

| Cas | Appels directs | Appels réfléchis |
| --- | ---: | ---: |
| Polymorphism | 100 000 027 | 0 |
| RBTree | 23 639 803 | 100 001 |
| LazyEvaluation | 9 010 005 | 1 001 |

La réflexion restante concerne donc une minorité des appels. Ces nombres ne mesurent pas sa part du temps CPU ; aucune mise en cache ou normalisation supplémentaire n’a été ajoutée dans cette étape.

## Reproduction

Depuis `/Users/0x1/Documents/htdocs/sharpurs/sharpurs` :

```sh
DOTNET="$HOME/.dotnet/dotnet" npm run test:runtime
```

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-sharpurs` :

```sh
./bin/sharp/run -c
"$HOME/.dotnet/dotnet" output/Main/bin/Release/net8.0/Program.dll
python3 scratch/sharpurs-apply-step1-20260908/prepare-dispatch-profile.py
"$HOME/.dotnet/dotnet" build scratch/sharpurs-apply-step1-20260908/dispatch-profile/Program.fsproj -c Release -p:NuGetAudit=false --ignore-failed-sources
"$HOME/.dotnet/dotnet" scratch/sharpurs-apply-step1-20260908/dispatch-profile/bin/Release/net8.0/Program.dll
```

Exécuter les benchmarks sans build concurrent. `dispatch-profile/` est ignoré par Git et ne remplace pas le programme généré normal.

## Preuves

`runtime-test.log` ; `backend-build.log` ; `run-c.log` et `run-c.stderr` ; `confirmation.log` et `confirmation.time` ; `comparison.json` ; `dispatch-counts.jsonl` ; `profile-build.log` ; `metadata.json` et `runtime.patch`.

Le plan `/Users/0x1/Documents/htdocs/sharpurs/todo.md` consigne l’étape terminée. Les étapes TAST suivantes ne sont pas commencées. Les changements de chemins Spago préexistants sont conservés.
