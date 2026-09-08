**Mesures de la règle de normalisation des tableaux — étape 3.5, 8 septembre 2026**

Ce dossier compare du Go régénéré par gopurs, avec la règle 3.3 désactivée (`before`) puis activée (`after`). Le compilateur de travail reste inchangé. Deux copies de son bundle sont conservées sous `compiler/bin` ; la seule différence est le retour positif du garde `isIntArrayFold`, remplacé par `false` dans le témoin. Le patch exact est dans `compiler-variant.patch`. Les optimisations antérieures du boxing FFI et des lambdas restent actives dans les deux versions. Aucun corps Go généré n'est réécrit pour obtenir le gain.

**Entrées et génération**

`./bin/go/run -c` a d'abord reconstruit le backend et le programme avec le `purs` typé utilisé par altbak, `GOMAXPROCS=14` et `GOFLAGS=-pgo=off`. Le lancement complet réussit ; son temps isolé n'entre pas dans la comparaison statistique.

`prepare.py` fige les 299 CoreFn et un même état initial de `.purmeta` pour les deux générations. Ce dernier point compte : PBO lit des implémentations dans ces métadonnées pour l'inlining. Des copies avec métadonnées vides modifiaient des modules sans rapport avec 3.3 ; elles ont été écartées avant toute mesure. Les empreintes CoreFn, sources PureScript/FFI et métadonnées initiales figurent dans `inputs-sha256.json`. Les bundles, versions, patches et empreintes du Go et des binaires sont conservés dans `metadata.json` et `gopurs-source.patch`.

Les deux arbres Go ne diffèrent que dans `purescript/Test_ArrayOps.go`, au corps de `Call_Test_ArrayOps_sumEvens`. `generated-go.patch` montre les deux buffers remplacés par la normalisation sur place, avec range, filtre et `Apply2` conservés. La génération de production a ensuite été refaite à partir du même état initial : elle est identique à `after` octet pour octet. Le binaire de production est reconstruit avec les mêmes options.

**Protocole**

- Apple M4 Pro, Go 1.27.0, darwin/arm64. `GOGC=800`, `GOMAXPROCS=14`, `GOMEMLIMIT=off`, `PPROF=0` ; compilation `-pgo=off -trimpath` explicite. Pas de profilage pendant les chronométrages.
- Les binaires sont compilés avant les mesures. Chaque échantillon utilise un processus neuf. Exécutions séquentielles, ordre avant/après pour les paires impaires et après/avant pour les paires paires.
- Cœur `sumEvens` : dix paires, deux tailles (900 et 90 000), une seconde par sous-benchmark Go. Six appels d'échauffement avant `ResetTimer`. Résultats attendus `202950` et `2025045000`, vérifiés avant et après les boucles mesurées.
- Les 14 vrais `act` : dix paires, 200 ms par sous-benchmark Go, référence et six échauffements hors chronométrage. Inclut les frontières `Effect`, `opaque` et `show`. Les getters sont résolus avant le chronométrage. Le temps inclut le coût des GC survenant pendant les boucles ; les compteurs d'allocations viennent du runner Go avec `ReportAllocs`.
- Programme `App` complet : dix paires, son protocole habituel conservé (trois chauffes globales, puis trois par test, minimum de dix mesures). Le total affiché est la somme des minima des tests, pas la durée totale du processus. Les 14 résultats sont comparés dans chaque exécution. Les temps affichés à deux décimales peuvent être arrondis à zéro.
- Pic RSS : dix paires par taille, processus séparés sous `/usr/bin/time -l`, six chauffes puis un nombre fixe et identique d'appels : 10 000 à 900 éléments, 100 à 90 000. Le maximum résident inclut runtime Go, runner de test, initialisation et échauffements. Cette sonde porte sur le noyau répété, pas sur le pic du programme altbak complet.

Le harnais `bench_test.go` est ajouté uniquement aux copies figées du module Go pour les mesures. Il ne remplace pas les fixtures de régression de gopurs. L'étape 3.4 reste ouverte.

**Exécution et résultats**

Depuis ce dossier, après préparation dans un dossier `generated` neuf :

```sh
python3 prepare.py
python3 measure.py check
python3 measure.py core
python3 measure.py act
python3 measure.py app
python3 measure.py rss
python3 analyze.py
```

`measure.py` ajoute les résultats à `measures.jsonl` : utiliser un nouveau fichier pour une nouvelle série complète. Chaque phase est exécutée une seule fois dans la série rapportée. Les sorties brutes sont dans `logs/`, les résultats attendus dans `expected-results.json`. `analyze.py` contrôle les paires, les inventaires, les doublons et l'identité des résultats, puis produit `summary.json` et `summary.md` : médianes, quartiles, min/max, deltas et ratios appariés. L'IQR décrit la dispersion, ce n'est pas un intervalle de confiance.

Les octets alloués par opération et la mémoire résidente sont deux mesures distinctes. Une baisse des premiers ne prouve pas une baisse générale de RAM. Les observations de cette session ne couvrent pas les autres tailles, configurations GC ou programmes.

Le bilan interprété est dans `RESULTS.md`. La série principale comprend 380 lignes validées, toutes conservées. Une tentative RSS dans le bac à sable a été interrompue par le refus de `sysctl kern.clockrate` avant de fournir une valeur RSS ; son journal est gardé dans `logs/rss-sandbox-unavailable.log`. Les 40 valeurs RSS proviennent de la série ensuite exécutée avec l'accès système nécessaire.

Un contrôle supplémentaire a été motivé par le ralentissement de Fib dans les 14 actes : `python3 check_fib.py`, dix paires d'une seconde, même harnais et mêmes binaires, uniquement `BenchmarkAct/Fib`. Ses 20 mesures et leur synthèse restent séparées dans `fib-control.jsonl` et `fib-control-summary.json` ; elles ne remplacent pas la première série.
