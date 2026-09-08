# Conversions de tableaux — mesures intégrées 3.5

Les valeurs sont présentées sous la forme **médiane [minimum ; maximum] ; IQR**. L’IQR utilise les quartiles inclusifs et décrit la dispersion ; ce n’est pas un intervalle de confiance.

Les ratios et variations sont calculés **après/avant pour chaque paire**, puis résumés. Un ratio de temps inférieur à 1 indique moins de temps. La médiane des différences appariées est distincte de la différence entre les deux médianes.

Les écarts faibles peuvent relever du bruit de mesure. Aucun test de significativité n’est effectué ; un gain local ne permet pas de conclure à une accélération générale.

**Allocations et mémoire résidente sont distinctes** : les octets/op sont cumulés pendant une exécution ; le RSS décrit la mémoire résidente rapportée pour le processus et ne doit pas être divisé par les itérations. Les Mo/MB ci-dessous valent 1 000 000 octets ; les Mio/MiB sont aussi disponibles dans le JSON.

## Noyau sumEvens

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| n900 (10 paires) | µs/op | 9.934 [9.547 ; 10.606] ; 0.284 | 9.689 [9.279 ; 10.49] ; 0.21875 | -0.295 | 0.969426 [0.900999 ; 1.09488] ; 0.0148151 | -3.05742 % |
| n900 (10 paires) | octets/op | 77056 [77056 ; 77056] ; 0 | 62080 [62080 ; 62080] ; 0 | -14976 | 0.805648 [0.805648 ; 0.805648] ; 0 | -19.4352 % |
| n900 (10 paires) | allocations/op | 15 [15 ; 15] ; 0 | 13 [13 ; 13] ; 0 | -2 | 0.866667 [0.866667 ; 0.866667] ; 0 | -13.3333 % |
| n90000 (10 paires) | µs/op | 941.735 [924.556 ; 959.224] ; 5.5175 | 933.593 [913.474 ; 991.08] ; 16.1285 | -10.334 | 0.989025 [0.977068 ; 1.03321] ; 0.0160294 | -1.09749 % |
| n90000 (10 paires) | octets/op | 9.25901e+06 [9.25901e+06 ; 9.25902e+06] ; 3 | 7.81722e+06 [7.81722e+06 ; 7.81722e+06] ; 4 | -1.4418e+06 | 0.844282 [0.844282 ; 0.844283] ; 5.23196e-07 | -15.5718 % |
| n90000 (10 paires) | allocations/op | 29 [29 ; 29] ; 0 | 27 [27 ; 27] ; 0 | -2 | 0.931034 [0.931034 ; 0.931034] ; 0 | -6.89655 % |

## Les 14 actes (avec leur enveloppe)

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| AstTree (10 paires) | µs/op | 0.37645 [0.3592 ; 0.3991] ; 0.026025 | 0.3806 [0.3625 ; 0.7975] ; 0.01885 | 0.00465 | 1.01216 [0.983397 ; 1.99825] ; 0.0335227 | 1.21588 % |
| AstTree (10 paires) | octets/op | 1808 [1808 ; 1808] ; 0 | 1808 [1808 ; 1808] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| AstTree (10 paires) | allocations/op | 49 [49 ; 49] ; 0 | 49 [49 ; 49] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Fib (10 paires) | µs/op | 0.29945 [0.2916 ; 0.3119] ; 0.00365 | 0.31655 [0.3059 ; 0.3827] ; 0.0144 | 0.0172 | 1.05642 [1.0241 ; 1.27227] ; 0.0412976 | 5.64194 % |
| Fib (10 paires) | octets/op | 114 [114 ; 114] ; 0 | 114 [114 ; 114] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Fib (10 paires) | allocations/op | 7 [7 ; 7] ; 0 | 7 [7 ; 7] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ListOps (10 paires) | µs/op | 9.2395 [8.983 ; 9.813] ; 0.27475 | 9.559 [9.206 ; 14.866] ; 0.413 | 0.2775 | 1.03033 [0.990108 ; 1.61289] ; 0.0383677 | 3.03254 % |
| ListOps (10 paires) | octets/op | 32594 [32594 ; 32594] ; 0 | 32594 [32594 ; 32594] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ListOps (10 paires) | allocations/op | 1362 [1362 ; 1362] ; 0 | 1362 [1362 ; 1362] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| TCO (10 paires) | µs/op | 41.655 [40.963 ; 42.658] ; 0.5985 | 42.1685 [41.363 ; 44.979] ; 1.06375 | 0.8245 | 1.02004 [0.985431 ; 1.05441] ; 0.0184474 | 2.00426 % |
| TCO (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| TCO (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Records (10 paires) | µs/op | 5.477 [5.403 ; 5.615] ; 0.0965 | 5.5105 [5.425 ; 5.67] ; 0.08225 | 0.0365 | 1.00671 [0.979697 ; 1.02089] ; 0.0193158 | 0.671151 % |
| Records (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Records (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Ackermann (10 paires) | µs/op | 16.7915 [16.422 ; 17.671] ; 0.61475 | 17.062 [16.316 ; 17.929] ; 0.62975 | -0.1135 | 0.99317 [0.969229 ; 1.0411] ; 0.0344512 | -0.683035 % |
| Ackermann (10 paires) | octets/op | 115 [115 ; 115] ; 0 | 115 [115 ; 115] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Ackermann (10 paires) | allocations/op | 7 [7 ; 7] ; 0 | 7 [7 ; 7] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Church (10 paires) | µs/op | 530.644 [487.466 ; 541.306] ; 40.481 | 530.016 [487.805 ; 549.608] ; 44.8398 | 3.0115 | 1.00596 [0.913634 ; 1.08123] ; 0.0218276 | 0.596088 % |
| Church (10 paires) | octets/op | 6928 [6928 ; 6928] ; 0 | 6928 [6928 ; 6928] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Church (10 paires) | allocations/op | 165 [165 ; 165] ; 0 | 165 [165 ; 165] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Primes (10 paires) | µs/op | 77.843 [76.083 ; 81.276] ; 3.78475 | 77.3735 [75.883 ; 82.994] ; 2.76875 | -1.2135 | 0.984276 [0.937742 ; 1.06954] ; 0.0353889 | -1.57243 % |
| Primes (10 paires) | octets/op | 257490 [257490 ; 257491] ; 0 | 257490 [257490 ; 257491] ; 0.75 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Primes (10 paires) | allocations/op | 10732 [10732 ; 10732] ; 0 | 10732 [10732 ; 10732] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| RBTree (10 paires) | µs/op | 24693.3 [24003.6 ; 25690.9] ; 647.713 | 24628 [24116.1 ; 25820.3] ; 397.272 | -72.435 | 0.99706 [0.960002 ; 1.02649] ; 0.0452187 | -0.29402 % |
| RBTree (10 paires) | octets/op | 7.94887e+07 [7.94887e+07 ; 7.94894e+07] ; 447.75 | 7.94888e+07 [7.94887e+07 ; 7.94893e+07] ; 436.25 | 6.5 | 1 [0.999992 ; 1.00001] ; 1.24546e-06 | 8.17726e-06 % |
| RBTree (10 paires) | allocations/op | 2.48396e+06 [2.48396e+06 ; 2.48396e+06] ; 0 | 2.48396e+06 [2.48396e+06 ; 2.48396e+06] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Polymorphism (10 paires) | µs/op | 2333.54 [2300.21 ; 2536.24] ; 97.0435 | 2345.58 [2272.64 ; 2522.72] ; 171.821 | -4.804 | 0.997925 [0.981207 ; 1.02527] ; 0.0173081 | -0.207467 % |
| Polymorphism (10 paires) | octets/op | 160 [160 ; 160] ; 0 | 160 [160 ; 160] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Polymorphism (10 paires) | allocations/op | 10 [10 ; 10] ; 0 | 10 [10 ; 10] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| StateMonad (10 paires) | µs/op | 140.028 [135.223 ; 144.889] ; 7.28525 | 139.771 [136.151 ; 150.566] ; 5.007 | -0.3365 | 0.99743 [0.968986 ; 1.06734] ; 0.0518625 | -0.256993 % |
| StateMonad (10 paires) | octets/op | 578729 [578728 ; 578730] ; 0.75 | 578729 [578728 ; 578729] ; 0 | 0 | 1 [0.999998 ; 1] ; 1.72792e-06 | 0 % |
| StateMonad (10 paires) | allocations/op | 13268 [13268 ; 13268] ; 0 | 13268 [13268 ; 13268] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| LazyEvaluation (10 paires) | µs/op | 245.651 [236.088 ; 262.826] ; 13.7363 | 249.679 [235.699 ; 257.388] ; 9.802 | 1.9575 | 1.00802 [0.949079 ; 1.08124] ; 0.0544039 | 0.802088 % |
| LazyEvaluation (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| LazyEvaluation (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ArrayOps (10 paires) | µs/op | 10.092 [9.816 ; 11.318] ; 0.20875 | 9.773 [9.522 ; 20.266] ; 0.45575 | -0.218 | 0.978371 [0.841315 ; 1.96605] ; 0.028895 | -2.16287 % |
| ArrayOps (10 paires) | octets/op | 77189 [77189 ; 77189] ; 0 | 62212 [62212 ; 62212] ; 0 | -14977 | 0.80597 [0.80597 ; 0.80597] ; 0 | -19.403 % |
| ArrayOps (10 paires) | allocations/op | 23 [23 ; 23] ; 0 | 21 [21 ; 21] ; 0 | -2 | 0.913043 [0.913043 ; 0.913043] ; 0 | -8.69565 % |
| RowToList (10 paires) | µs/op | 0.068195 [0.06655 ; 0.06988] ; 0.00176 | 0.068785 [0.06709 ; 0.1109] ; 0.0017775 | 0.000945 | 1.01405 [0.964491 ; 1.58701] ; 0.0507759 | 1.40523 % |
| RowToList (10 paires) | octets/op | 112 [112 ; 112] ; 0 | 112 [112 ; 112] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| RowToList (10 paires) | allocations/op | 6 [6 ; 6] ; 0 | 6 [6 ; 6] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |

## Application : temps de chaque test

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| AstTree (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |
| Fib (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |
| ListOps (10 paires) | µs | 8 [7 ; 9] ; 0.75 | 8 [7 ; 10] ; 0.1875 | 0.5 | 1.0625 [0.777778 ; 1.25] ; 0.22619 | 6.25 % |
| TCO (10 paires) | µs | 40 [36 ; 41] ; 3.1875 | 39.5 [36 ; 42] ; 3 | 1 | 1.02608 [0.878049 ; 1.08609] ; 0.130413 | 2.6084 % |
| Records (10 paires) | µs | 5 [5 ; 6] ; 0 | 5 [4.75 ; 6] ; 0 | 0 | 1 [0.833333 ; 1.2] ; 0 | 0 % |
| Ackermann (10 paires) | µs | 17 [15 ; 17] ; 0 | 17 [15 ; 18] ; 0 | 0 | 1 [0.882353 ; 1.13333] ; 0.0441176 | 0 % |
| Church (10 paires) | µs | 477.5 [468 ; 487] ; 3.5 | 465 [453 ; 480] ; 12.25 | -11.5 | 0.975817 [0.947699 ; 1.01068] ; 0.0357793 | -2.41829 % |
| Primes (10 paires) | µs | 81 [72 ; 89] ; 10.75 | 82 [67 ; 90] ; 11.5 | 2.5 | 1.03284 [0.77907 ; 1.22222] ; 0.224148 | 3.2841 % |
| RBTree (10 paires) | µs | 23619 [22966 ; 24084] ; 483 | 23463.5 [22816 ; 23843] ; 173 | -173.5 | 0.992621 [0.963107 ; 1.03797] ; 0.0217577 | -0.737893 % |
| Polymorphism (10 paires) | µs | 2364.5 [2216 ; 2471] ; 207 | 2310.5 [2216.75 ; 2459.25] ; 129.188 | -30.5 | 0.986736 [0.909622 ; 1.10927] ; 0.0687232 | -1.3264 % |
| StateMonad (10 paires) | µs | 122.5 [114 ; 152] ; 6.0625 | 123 [107 ; 134] ; 6.5 | -5 | 0.959797 [0.822368 ; 1.08943] ; 0.0986896 | -4.0203 % |
| LazyEvaluation (10 paires) | µs | 248 [229 ; 263] ; 19.5 | 236.125 [229 ; 258] ; 21 | -8.375 | 0.965816 [0.870722 ; 1.09565] ; 0.0827289 | -3.41838 % |
| ArrayOps (10 paires) | µs | 9.5 [9 ; 12] ; 1.75 | 9 [8 ; 11] ; 0.9375 | -1 | 0.894444 [0.770833 ; 1.22222] ; 0.141414 | -10.5556 % |
| RowToList (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |

**Temps affichés arrondis à zéro : 60 observations (AstTree, Fib, RowToList).** Un affichage de 0,00 µs ne démontre pas un coût nul. Les valeurs restent dans les statistiques absolues ; les paires dont le temps avant vaut zéro sont exclues des ratios et pourcentages. Le nombre de paires utilisables est indiqué dans leur cellule. Les ratios impliquant un temps après arrondi à zéro ne prouvent pas une suppression totale du coût.

## Application : total rapporté

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| Total (10 paires) | ms | 27.005 [26.23 ; 27.54] ; 0.4575 | 26.695 [26.28 ; 27.27] ; 0.265 | -0.105 | 0.99609 [0.964779 ; 1.03088] ; 0.0277723 | -0.391023 % |

Le total de l’application est celui fourni par son protocole : il n’est pas reconstruit en additionnant des médianes. S’il provient de Bench.runBench, il additionne les minima de dix mesures par test, et ne représente pas le temps mural de tout le processus.

**Les 14 résultats de l’application sont identiques dans toutes les lignes avant et après.**

| Test | Résultat |
|---|---|
| AstTree | 7 |
| Fib | 55 |
| ListOps | 202950 |
| TCO | 100000 |
| Records | 20000 |
| Ackermann | 125 |
| Church | 100000 |
| Primes | 21536 |
| RBTree | 22 |
| Polymorphism | 10000000 |
| StateMonad | 1200 |
| LazyEvaluation | 1000000 |
| ArrayOps | 202950 |
| RowToList | 5 |

## RSS rapporté par processus

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| n900 (10 paires) | Mo / MB | 41.7464 [41.2713 ; 42.0905] ; 0.372736 | 41.4515 [41.0419 ; 42.1233] ; 0.311296 | -0.262144 | 0.993715 [0.982738 ; 1.01469] ; 0.0198627 | -0.628491 % |
| n90000 (10 paires) | Mo / MB | 53.9689 [53.674 ; 54.4113] ; 0.229376 | 48.7997 [44.4006 ; 51.7571] ; 3.60858 | -4.89882 | 0.908772 [0.823458 ; 0.954958] ; 0.0654902 | -9.12275 % |

- n900 : itérations par processus 10000 ; même nombre avant/après au sein de chaque paire.
- n90000 : itérations par processus 100 ; même nombre avant/après au sein de chaque paire.

## Lecture prudente des temps

- Noyau n900 : -3.05742 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Noyau n90000 : -1.09749 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte AstTree : 1.21588 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte ListOps : 3.03254 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Acte TCO : 2.00426 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Acte Records : 0.671151 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Ackermann : -0.683035 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Church : 0.596088 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Primes : -1.57243 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte RBTree : -0.29402 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Polymorphism : -0.207467 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte StateMonad : -0.256993 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte LazyEvaluation : 0.802088 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte ArrayOps : -2.16287 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte RowToList : 1.40523 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Application total : -0.391023 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.

Le seuil de 5 % est un repère de lecture, pas un seuil statistique. Les statistiques complètes, différences, ratios et numéros des paires figurent dans summary.json.
