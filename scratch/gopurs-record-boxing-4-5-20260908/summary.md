# Boxing des records — StateMonad, mesures intégrées 4.5

Les valeurs sont présentées sous la forme **médiane [minimum ; maximum] ; IQR**. L’IQR utilise les quartiles inclusifs et décrit la dispersion ; ce n’est pas un intervalle de confiance.

Les ratios et variations sont calculés **après/avant pour chaque paire**, puis résumés. Un ratio de temps inférieur à 1 indique moins de temps. La médiane des différences appariées est distincte de la différence entre les deux médianes.

Les écarts faibles peuvent relever du bruit de mesure. Aucun test de significativité n’est effectué ; un gain local ne permet pas de conclure à une accélération générale.

**Allocations et mémoire résidente sont distinctes** : les octets/op sont cumulés pendant une exécution ; le RSS décrit la mémoire résidente rapportée pour le processus et ne doit pas être divisé par les itérations. Les Mo/MB ci-dessous valent 1 000 000 octets ; les Mio/MiB sont aussi disponibles dans le JSON.

## Noyau StateMonad — runManyTimes(20, 0), résultat 1200

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| n20 (10 paires) | µs/op | 124.084 [120.867 ; 145.706] ; 9.331 | 82.392 [80.503 ; 92.678] ; 0.76425 | -42.084 | 0.659081 [0.609667 ; 0.682461] ; 0.0402542 | -34.0919 % |
| n20 (10 paires) | octets/op | 578560 [578560 ; 578562] ; 1 | 404800 [404800 ; 404800] ; 0 | -173760 | 0.699668 [0.699666 ; 0.699668] ; 1.20932e-06 | -30.0332 % |
| n20 (10 paires) | allocations/op | 13260 [13260 ; 13260] ; 0 | 6020 [6020 ; 6020] ; 0 | -7240 | 0.453997 [0.453997 ; 0.453997] ; 0 | -54.6003 % |

## Les 14 actes (avec leur enveloppe)

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| AstTree (10 paires) | µs/op | 0.3792 [0.3422 ; 0.4171] ; 0.0357 | 0.36505 [0.3434 ; 0.4019] ; 0.025075 | -0.00735 | 0.981015 [0.867418 ; 1.07835] ; 0.0953338 | -1.89847 % |
| AstTree (10 paires) | octets/op | 1808 [1808 ; 1808] ; 0 | 1808 [1808 ; 1808] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| AstTree (10 paires) | allocations/op | 49 [49 ; 49] ; 0 | 49 [49 ; 49] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Fib (10 paires) | µs/op | 0.30825 [0.2995 ; 0.3324] ; 0.013125 | 0.30775 [0.2962 ; 0.3183] ; 0.008525 | -0.005 | 0.983476 [0.944344 ; 1.0467] ; 0.0214941 | -1.6524 % |
| Fib (10 paires) | octets/op | 114 [114 ; 114] ; 0 | 114 [114 ; 114] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Fib (10 paires) | allocations/op | 7 [7 ; 7] ; 0 | 7 [7 ; 7] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ListOps (10 paires) | µs/op | 9.247 [8.572 ; 10.081] ; 0.782 | 9.335 [8.634 ; 10.123] ; 0.82 | -0.215 | 0.977344 [0.916573 ; 1.12179] ; 0.0540108 | -2.26555 % |
| ListOps (10 paires) | octets/op | 32594 [32594 ; 32594] ; 0 | 32594 [32594 ; 32594] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ListOps (10 paires) | allocations/op | 1362 [1362 ; 1362] ; 0 | 1362 [1362 ; 1362] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| TCO (10 paires) | µs/op | 41.2895 [40.813 ; 42.352] ; 0.67225 | 41.2985 [40.744 ; 43.15] ; 0.835 | -0.2705 | 0.99343 [0.983283 ; 1.04414] ; 0.0129972 | -0.65699 % |
| TCO (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| TCO (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Records (10 paires) | µs/op | 5.3935 [5.371 ; 5.576] ; 0.05975 | 5.485 [5.359 ; 5.626] ; 0.14025 | -0.004 | 0.999261 [0.983859 ; 1.04459] ; 0.0221965 | -0.0738559 % |
| Records (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Records (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Ackermann (10 paires) | µs/op | 16.7925 [15.833 ; 17.984] ; 1.13625 | 17.2585 [15.756 ; 17.857] ; 1.741 | -0.1565 | 0.990772 [0.918985 ; 1.09714] ; 0.0392396 | -0.922757 % |
| Ackermann (10 paires) | octets/op | 115 [115 ; 115] ; 0 | 115 [115 ; 115] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Ackermann (10 paires) | allocations/op | 7 [7 ; 7] ; 0 | 7 [7 ; 7] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Church (10 paires) | µs/op | 488.382 [477.647 ; 504.7] ; 17.1247 | 482.962 [476.717 ; 496.444] ; 8.854 | -6.534 | 0.986864 [0.952677 ; 1.02549] ; 0.0366475 | -1.31361 % |
| Church (10 paires) | octets/op | 6928 [6928 ; 6928] ; 0 | 6928 [6928 ; 6928] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Church (10 paires) | allocations/op | 165 [165 ; 165] ; 0 | 165 [165 ; 165] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Primes (10 paires) | µs/op | 77.5885 [74.046 ; 83.426] ; 5.0345 | 76.576 [74.106 ; 82.062] ; 3.523 | -1.624 | 0.979418 [0.888932 ; 1.09174] ; 0.0260841 | -2.0582 % |
| Primes (10 paires) | octets/op | 257490 [257489 ; 257492] ; 0.75 | 257490 [257489 ; 257491] ; 1 | 0 | 1 [0.999992 ; 1.00001] ; 5.82546e-06 | 0 % |
| Primes (10 paires) | allocations/op | 10732 [10732 ; 10732] ; 0 | 10732 [10732 ; 10732] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| RBTree (10 paires) | µs/op | 24941.9 [22888.2 ; 26462] ; 1833.54 | 24758.8 [22625.8 ; 25877.6] ; 2257.27 | -667.003 | 0.97407 [0.885429 ; 1.09208] ; 0.0608059 | -2.59297 % |
| RBTree (10 paires) | octets/op | 7.94887e+07 [7.94887e+07 ; 7.94894e+07] ; 519.75 | 7.94887e+07 [7.94887e+07 ; 7.94899e+07] ; 447.25 | -8.5 | 1 [0.999992 ; 1.00002] ; 3.00041e-06 | -1.06933e-05 % |
| RBTree (10 paires) | allocations/op | 2.48396e+06 [2.48396e+06 ; 2.48396e+06] ; 0.75 | 2.48396e+06 [2.48396e+06 ; 2.48396e+06] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Polymorphism (10 paires) | µs/op | 2387.42 [2239.81 ; 2508.6] ; 197.73 | 2365.62 [2231.74 ; 2472.5] ; 153.721 | -15.5425 | 0.993465 [0.912506 ; 1.07152] ; 0.0228191 | -0.653548 % |
| Polymorphism (10 paires) | octets/op | 160 [160 ; 160] ; 0 | 160 [160 ; 160] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| Polymorphism (10 paires) | allocations/op | 10 [10 ; 10] ; 0 | 10 [10 ; 10] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| StateMonad (10 paires) | µs/op | 141.95 [124.621 ; 155.962] ; 8.1485 | 86.3885 [82.002 ; 102.392] ; 8.73275 | -51.8305 | 0.643171 [0.546236 ; 0.72908] ; 0.0868827 | -35.6829 % |
| StateMonad (10 paires) | octets/op | 578729 [578728 ; 578730] ; 0.75 | 404957 [404956 ; 404957] ; 0.75 | -173772 | 0.699735 [0.699732 ; 0.699736] ; 1.20909e-06 | -30.0265 % |
| StateMonad (10 paires) | allocations/op | 13268 [13268 ; 13268] ; 0 | 6028 [6028 ; 6028] ; 0 | -7240 | 0.454326 [0.454326 ; 0.454326] ; 0 | -54.5674 % |
| LazyEvaluation (10 paires) | µs/op | 241.909 [230.349 ; 257.923] ; 19.5598 | 245.184 [229.572 ; 259.128] ; 13.1125 | -1.2965 | 0.994392 [0.929995 ; 1.10884] ; 0.0324446 | -0.560832 % |
| LazyEvaluation (10 paires) | octets/op | 128 [128 ; 128] ; 0 | 128 [128 ; 128] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| LazyEvaluation (10 paires) | allocations/op | 8 [8 ; 8] ; 0 | 8 [8 ; 8] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ArrayOps (10 paires) | µs/op | 9.3165 [8.949 ; 9.743] ; 0.38575 | 9.235 [8.927 ; 9.831] ; 0.266 | -0.0375 | 0.995952 [0.949913 ; 1.0537] ; 0.0266663 | -0.404774 % |
| ArrayOps (10 paires) | octets/op | 62212 [62212 ; 62212] ; 0 | 62212 [62212 ; 62212] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| ArrayOps (10 paires) | allocations/op | 21 [21 ; 21] ; 0 | 21 [21 ; 21] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| RowToList (10 paires) | µs/op | 0.067165 [0.0646 ; 0.07049] ; 0.002685 | 0.066455 [0.06365 ; 0.0716] ; 0.003135 | 0.00022 | 1.0033 [0.94576 ; 1.05773] ; 0.0533792 | 0.330236 % |
| RowToList (10 paires) | octets/op | 112 [112 ; 112] ; 0 | 112 [112 ; 112] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |
| RowToList (10 paires) | allocations/op | 6 [6 ; 6] ; 0 | 6 [6 ; 6] ; 0 | 0 | 1 [1 ; 1] ; 0 | 0 % |

## Application : temps de chaque test

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| AstTree (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |
| Fib (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |
| ListOps (10 paires) | µs | 8.5 [7 ; 10] ; 1.75 | 8 [7 ; 8] ; 1 | -1 | 0.888889 [0.7 ; 1] ; 0.18125 | -11.1111 % |
| TCO (10 paires) | µs | 41 [40 ; 41] ; 0.8125 | 41 [40 ; 42] ; 0.25 | 0 | 1 [0.97561 ; 1.025] ; 0.0243996 | 0 % |
| Records (10 paires) | µs | 5 [4.75 ; 5] ; 0 | 5 [5 ; 5] ; 0 | 0 | 1 [1 ; 1.05263] ; 0 | 0 % |
| Ackermann (10 paires) | µs | 17 [16 ; 17] ; 0.75 | 17 [16 ; 17] ; 0.75 | 0 | 1 [0.941176 ; 1.0625] ; 0 | 0 % |
| Church (10 paires) | µs | 475 [471 ; 486] ; 4.75 | 472.5 [470 ; 481] ; 4.25 | -1.5 | 0.996831 [0.971193 ; 1.01907] ; 0.0121111 | -0.316903 % |
| Primes (10 paires) | µs | 87.5 [70.75 ; 90.25] ; 17.75 | 77 [72 ; 90] ; 11 | 0.5 | 1.00575 [0.808989 ; 1.08451] ; 0.128552 | 0.574713 % |
| RBTree (10 paires) | µs | 24013.5 [22443 ; 24848] ; 891.25 | 24109.5 [22140 ; 25098] ; 1077.75 | -145.5 | 0.993918 [0.934927 ; 1.05809] ; 0.0195797 | -0.608202 % |
| Polymorphism (10 paires) | µs | 2473 [2284 ; 2521] ; 69 | 2478.5 [2380 ; 2503.75] ; 24.25 | 14 | 1.00566 [0.950859 ; 1.08187] ; 0.0366476 | 0.566114 % |
| StateMonad (10 paires) | µs | 124 [114 ; 134] ; 3.75 | 89.5 [81 ; 98] ; 9.25 | -34.5 | 0.730327 [0.648 ; 0.790323] ; 0.0446568 | -26.9673 % |
| LazyEvaluation (10 paires) | µs | 255 [229 ; 258] ; 21.0625 | 255 [230 ; 258] ; 7.5 | -1.5 | 0.994186 [0.901961 ; 1.12664] ; 0.0123639 | -0.581395 % |
| ArrayOps (10 paires) | µs | 10 [8.75 ; 11] ; 0.75 | 10 [8.75 ; 11] ; 0.625 | 0 | 1 [0.875 ; 1.25714] ; 0.179672 | 0 % |
| RowToList (10 paires) | µs | 0 [0 ; 0] ; 0 | 0 [0 ; 0] ; 0 | 0 | non défini (0/10 paires) | non définie |

**Temps affichés arrondis à zéro : 60 observations (AstTree, Fib, RowToList).** Un affichage de 0,00 µs ne démontre pas un coût nul. Les valeurs restent dans les statistiques absolues ; les paires dont le temps avant vaut zéro sont exclues des ratios et pourcentages. Le nombre de paires utilisables est indiqué dans leur cellule. Les ratios impliquant un temps après arrondi à zéro ne prouvent pas une suppression totale du coût.

## Application : total rapporté

| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |
|---|---|---:|---:|---:|---:|---:|
| Total (10 paires) | ms | 27.495 [25.96 ; 28.31] ; 0.8025 | 27.54 [25.58 ; 28.47] ; 1.0675 | -0.16 | 0.994021 [0.942173 ; 1.04554] ; 0.0158627 | -0.597885 % |

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
| n20 (10 paires) | Mo / MB | 44.3023 [43.778 ; 93.438] ; 0.405504 | 43.6716 [43.1555 ; 44.1713] ; 0.323584 | -0.425984 | 0.990328 [0.461862 ; 0.998132] ; 0.00902636 | -0.967162 % |

- n20 : itérations par processus 10000 ; même nombre avant/après au sein de chaque paire.

## Lecture prudente des temps

- Acte AstTree : -1.89847 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Fib : -1.6524 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Acte ListOps : -2.26555 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte TCO : -0.65699 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Records : -0.0738559 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Ackermann : -0.922757 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Church : -1.31361 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte Primes : -2.0582 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Acte RBTree : -2.59297 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.
- Acte Polymorphism : -0.653548 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte LazyEvaluation : -0.560832 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte ArrayOps : -0.404774 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Acte RowToList : 0.330236 % ; écart médian inférieur à 5 %, à ne pas surinterpréter ; les quartiles des variations encadrent zéro.
- Application total : -0.597885 % ; écart médian inférieur à 5 %, à ne pas surinterpréter.

Le seuil de 5 % est un repère de lecture, pas un seuil statistique. Les statistiques complètes, différences, ratios et numéros des paires figurent dans summary.json.
