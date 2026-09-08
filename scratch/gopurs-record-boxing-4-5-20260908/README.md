# Boxing des records natifs — mesure 4.5

Comparaison du compilateur courant avec le seul boxing natif compact de zéro à cinq champs désactivé (`before`), puis activé (`after`). Les helpers du printer, la normalisation des tableaux 3.3, les lambdas et le boxing FFI restent actifs dans les deux versions. Aucun Go généré n'est réécrit pour obtenir le gain.

## Entrées et preuve du changement

`bin/go/run -c` a reconstruit le programme avec le `purs` typé d'altbak avant de figer les entrées. `prepare.py` conserve les 299 CoreFn, les empreintes des sources PureScript/FFI et une copie physique du même état initial de `.purmeta` pour chaque génération. Le patch du bundle témoin est dans `compiler-variant.patch` ; les versions, options et empreintes sont dans `metadata.json` et `inputs-sha256.json`.

La vérification AST `verify_codegen.go` compare les 386 fichiers Go, en normalisant les appels compacts vers leur forme générique dans les deux arbres. Elle vérifie qu'aucune autre différence de code n'existe, indépendamment des commentaires et positions. Rapport : `logs/verify-codegen.log`. Les 116 fichiers modifiés totalisent 269 remplacements ; StateMonad contient neuf sites natifs qui passent de `RecordDict` à `RecordDict2`. Runtime et FFI restent identiques. La génération de production est également identique à la copie `after`.

## Protocole

Apple M4 Pro, 14 cœurs, 48 Gio ; Go 1.27.0 darwin/arm64. `GOGC=800`, `GOMAXPROCS=14`, `GOMEMLIMIT=off`, `PPROF=0`. Binaires précompilés avec `-pgo=off -trimpath`. Les mesures sont séquentielles et utilisent un processus neuf : dix paires, ordre avant/après aux paires impaires et après/avant aux paires paires.

- Noyau réel `Call_Test_StateMonad_runManyTimes(20, 0)` : résultat attendu 1200, six chauffes, une seconde par benchmark Go, compteurs d'allocations.
- Les 14 vrais `act` : résultats connus vérifiés, getters résolus et six chauffes avant le chronométrage ; 200 ms par sous-benchmark. Les frontières Effect, opaque et show sont conservées.
- Programme App : protocole habituel inchangé, trois chauffes globales puis trois par test, minimum de dix mesures. Le total affiché additionne ces minima ; il ne mesure pas la durée du processus. Les 14 résultats sont contrôlés à chaque lancement.
- Pic RSS : processus séparés sous `/usr/bin/time -l`, six chauffes puis exactement 10 000 appels du noyau n20. Le maximum résident inclut le runtime et le harnais ; ce n'est pas le pic de l'App complète.
- Profils d'allocations : exécutions séparées avec échantillonnage exhaustif, avant/après la boucle chaude, puis soustraction et filtrage sur StateMonad. Leurs temps n'entrent pas dans les mesures de vitesse.

Le pilote court est isolé dans `pilot.jsonl` et `pilot-logs`. Les mesures principales restent dans `measures.jsonl`, sans écrasement ni doublons. Les quartiles sont inclusifs ; l'IQR décrit la dispersion et n'est pas un intervalle de confiance.

## Reproduction

Préparer un dossier neuf avec ces scripts et le harnais, sous `altbak.pub/scratch`, puis :

```sh
python3 prepare.py
python3 measure.py check
python3 measure.py core
python3 measure.py act
python3 measure.py app
python3 measure.py rss
python3 profile.py
python3 analyze.py
```

Les accès nécessaires au cache Spago et aux compteurs système de macOS doivent être disponibles. Les scripts refusent de réutiliser les générations ou journaux existants. Aucun test de la suite `passing` n'est lancé dans cette étape.

Le README d'altbak conserve la baseline historique (State Monad ~114 µs, total ~24,66 ms). Elle sert de contexte, pas de témoin statistique pour cette session : l'attribution repose sur les deux variantes mesurées ensemble. Le bilan interprété est dans `RESULTS.md` ; statistiques complètes et données brutes sont conservées à côté.
