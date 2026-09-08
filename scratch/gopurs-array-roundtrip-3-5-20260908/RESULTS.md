**3.5 — Bilan des mesures sur le Go régénéré**

La règle intégrée réduit nettement les allocations. Le gain de vitesse observé sur Arrays est petit ; aucun gain global d'altbak n'est démontré. La mémoire résidente n'augmente pas de façon nette dans les deux charges testées. Une petite régression de temps sur Fib a également été reproduite et doit être conservée dans le bilan.

Comparaison sur Apple M4 Pro (14 cœurs, 48 Gio), Go 1.27.0, `GOGC=800`, `GOMAXPROCS=14`, `GOMEMLIMIT=off`, PGO désactivé. Dix paires alternées, processus neufs et binaires précompilés pour chaque famille de mesures. Les 299 CoreFn et les métadonnées initiales du backend sont figés. Le seul changement entre les deux arbres Go est le corps de `Call_Test_ArrayOps_sumEvens` ; les autres optimisations du compilateur sont conservées.

**Cœur sumEvens : médianes avant → après**

| Taille | Temps/op | Variation des médianes | Octets/op | Allocations/op |
|---|---:|---:|---:|---:|
| 900 | 9,934 → 9,689 µs | −2,47 % | 77 056 → 62 080 (−19,44 %) | 15 → 13 |
| 90 000 | 941,735 → 933,594 µs | −0,86 % | 9 259 014 → 7 817 217 (−15,57 %) | 29 → 27 |

À 900, huit paires sur dix sont favorables ; les quartiles des temps sont 9,712–9,996 µs avant et 9,571–9,789 µs après. La variation appariée médiane est −3,06 %. À 90 000, sept paires sont favorables ; les quartiles sont 938,645–944,163 µs avant et 923,622–939,751 µs après. La variation appariée médiane est −1,10 %, avec des quartiles −1,57 % à +0,03 % : le gain de vitesse à cette taille reste peu convaincant. Tous les échantillons, y compris les plus défavorables, sont conservés.

Les économies d'allocation sont stables : deux allocations supprimées par appel, sans nouveau buffer pour la normalisation. Le gain de 7,1 % du prototype qui supprimait les deux passes n'est pas reproduit par cette variante, qui garde une passe de normalisation.

**Les 14 tests d'altbak**

Les 14 résultats sont identiques avant/après et vérifiés dans chaque lancement du programme complet. Les résultats du noyau sont aussi vérifiés : `202950` et `2025045000`.

Avec l'enveloppe réelle `act` (`Effect`, `opaque`, `show`), ArrayOps passe de 10,092 à 9,773 µs en médiane, 77 189 à 62 212 octets/op et 23 à 21 allocations/op. La variation appariée du temps vaut −2,16 %, avec dispersion traversant zéro. Les allocations des autres tests restent identiques, à quelques octets d'instrumentation/GC près dans les gros lots.

Dans le protocole habituel « best of 10 », ArrayOps passe de 9,5 à 9 µs et le total affiché de 27,005 à 26,695 ms. Le total est une somme de minima, pas la durée du processus. Sa variation appariée médiane n'est que de −0,39 %, avec quartiles −2,34 % à +0,43 % : **aucune amélioration globale démontrée**. ArrayOps ne représente qu'environ 0,035 % de ce total ; son petit gain local ne peut expliquer un écart global de l'ordre de 1 %.

Les variations des autres temps ne sont pas toutes favorables. Fib ralentit dans les dix premières paires (299,45 → 316,55 ns/op), ListOps dans neuf (+3,03 % apparié) et TCO dans huit (+2,00 %). Le Go de ces fonctions est pourtant identique : cela n'exclut pas des effets indirects du binaire ou du contexte d'exécution.

Un contrôle isolé de Fib, dix nouvelles paires d'une seconde avec les mêmes binaires, reproduit le ralentissement : **293,0 → 301,85 ns/op**, soit environ **+3 % et +8,85 ns/op**. Toutes les paires sont défavorables ; variation appariée médiane +3,15 %, quartiles +2,08 % à +3,60 %. Ses 114 octets/op et 7 allocations/op restent identiques. La cause de cette petite régression n'est pas établie ; elle ne doit pas être qualifiée d'absence de régression sous prétexte que le Go source est inchangé.

**Pic mémoire du noyau répété**

Mesure séparée sous `/usr/bin/time -l`, avec six chauffes puis exactement le même nombre d'appels par processus avant/après. Le RSS inclut le runtime Go et le runner ; il ne mesure pas le programme altbak complet. Un Mo vaut ici 1 000 000 octets.

| Taille et travail fixe | Pic RSS médian avant → après | Plage avant | Plage après |
|---|---:|---:|---:|
| 900, 10 000 appels | 41,75 → 41,45 Mo | 41,27–42,09 Mo | 41,04–42,12 Mo |
| 90 000, 100 appels | 53,97 → 48,80 Mo | 53,67–54,41 Mo | 44,40–51,76 Mo |

À 900, le pic est stable compte tenu de la dispersion. À 90 000, les dix paires ont un pic inférieur après optimisation ; la baisse appariée médiane vaut 9,12 %. Ces mesures ne montrent donc pas de pénalité mémoire dans ces deux cas. Elles ne garantissent pas une baisse pour tout programme : la capacité excédentaire du buffer du filtre reste vivante pendant le fold.

**Portée et suite**

3.5 est mesurée. La règle présente surtout un intérêt pour réduire les allocations des programmes ayant ce motif, avec un petit gain local de temps possible. Ces résultats ne justifient aucune promesse d'accélération globale d'altbak ni d'économie de RAM universelle. La validation étendue de 3.4 reste ouverte, ainsi que l'examen des signaux de ralentissement avant de conclure à une intégration sans régression.

Protocole et reproduction : [README.md](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/README.md). Les tableaux détaillés des 14 tests, min/max, quartiles et ratios figurent dans [summary.md](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/summary.md) et [summary.json](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/summary.json). Les mesures brutes sont dans [measures.jsonl](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/measures.jsonl), et le contrôle Fib dans [fib-control-summary.json](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/fib-control-summary.json).
