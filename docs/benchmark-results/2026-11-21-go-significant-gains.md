# Gains Go intégrés — 21 novembre 2026

45 cas vérifiés : 14 core, 23 excluded WIP, 3 diagnostics et 5 extended.
Le total core apparié passe de **13.085758 à 10.654139 ms**
(**−18.58 %**), contre **13,077670 ms** au README initial.
Les colonnes historiques C/JS restent les références pour les ratios publiés.
Ces ratios utilisent les valeurs affichées dans le README ; leur précision est
donc limitée pour les cas de quelques nanosecondes et les références censurées.
Les sources des benchmarks, leurs entrées et leurs oracles sont inchangés.

## Changements retenus

- PBO : conserver la spécialisation récursive de listes lorsque les variables de
  type du corps et de la signature ont des noms de portée différents ; respecter
  le masquage des variables dans les patterns.
- Compilateur : callbacks natifs map/filter, tri Int canonique, somme native,
  traverse Maybe strict, suppression des Step/thunks terminaux de tailRecM,
  fusion de conversions adjacentes et application immédiatement saturée d’Effect.
- Runtime/FFI : cellule ST sans allocation à chaque écriture, append en Value,
  concaténation linéaire, scans Unicode sans tableaux temporaires et remplacement
  regex littéral sans matérialiser les captures.
- Récurrences affines : reconnaître une boucle pure complète et calculer sa somme
  modulo 2³². Les gains extrêmes de Polymorphism/Lazy proviennent de cette réduction
  algébrique, pas d’une suppression de dictionnaires déjà absents. Garde sur les
  entrées et boucle d’origine conservées ; aucun nom ou résultat de benchmark codé.

## Comparaison complète

Les temps sont en µs/action. « Avant » et « Après » sont mesurés dans la même
campagne ; la colonne README est la référence historique distincte.
Array Indexing conserve 8 388 608 lectures, représentation boxed, taille 16 384.

| Cas | README initial | Avant | Après | Avant / après |
| --- | ---: | ---: | ---: | ---: |
| AST Evaluation | 0.350000 | 0.314621 | 0.318653 | 0.99× |
| Fibonacci | 0.210000 | 0.205587 | 0.205484 | 1.00× |
| List Processing | 17.090000 | 16.255981 | 8.618368 | 1.89× |
| Tail Call Optimization | 64.190000 | 63.836102 | 63.819984 | 1.00× |
| Deep Record Updates | 8.760000 | 8.799500 | 8.798157 | 1.00× |
| Ackermann | 16.380000 | 17.732544 | 16.801839 | 1.06× |
| Church Numerals | 233.790000 | 230.519531 | 230.440094 | 1.00× |
| Prime Sieve | 77.260000 | 73.153320 | 74.518879 | 0.98× |
| Red-Black Tree | 10037.620000 | 9909.792000 | 10157.125000 | 0.98× |
| Polymorphism | 2288.660000 | 2417.104250 | 0.042611 | 56724.89× |
| State Monad | 94.280000 | 90.451820 | 89.740234 | 1.01× |
| Lazy Evaluation | 231.110000 | 249.688156 | 0.776375 | 321.61× |
| Array Processing | 7.800000 | 7.761637 | 2.790100 | 2.78× |
| RowToList | 0.150000 | 0.142484 | 0.143687 | 0.99× |
| Array Indexing | 4596.080000 | 4550.334000 | 4630.792000 | 0.98× |
| JSON to Typed AST | 20204.830000 | 20156.896000 | 20534.729000 | 0.98× |
| JSON Decoding | 2172.480000 | 2355.291500 | 2220.354000 | 1.06× |
| array/map-10000 | 166.810000 | 162.417328 | 131.847656 | 1.23× |
| array/filter-10000 | 107.890000 | 103.297195 | 76.492508 | 1.35× |
| array/foldl-sum-10000 | 66.880000 | 65.097656 | 4.948466 | 13.16× |
| array/sort-descending-10000 | 620.040000 | 599.682312 | 154.746094 | 3.88× |
| array/eq-neq-last-10000 | 23.320000 | 23.032471 | 16.065349 | 1.43× |
| array/append-1000 | 54.390000 | 52.437824 | 30.561117 | 1.72× |
| foldable/traverse-maybe-1000 | 235.230000 | 218.710297 | 17.474895 | 12.52× |
| foldable/traverse-maybe-fail-first-1000 | 187.230000 | 177.567703 | 3.846212 | 46.17× |
| foldable/traverse-maybe-fail-last-1000 | 186.120000 | 176.681000 | 3.797404 | 46.53× |
| string/cu-length-unicode-1000 | 0.010000 | 0.009412 | 0.009910 | 0.95× |
| string/cp-length-unicode-1000 | 50.880000 | 49.047852 | 6.712056 | 7.31× |
| string/indexOf-hit-10000 | 0.410000 | 0.409704 | 0.410172 | 1.00× |
| string/indexOf-miss-10000 | 2.440000 | 2.562622 | 2.510376 | 1.02× |
| string/cp-take-unicode-1500 | 59.220000 | 57.557945 | 44.483234 | 1.29× |
| string/split-1000 | 163.980000 | 157.261063 | 133.980469 | 1.17× |
| string/concat-1000 | 125.610000 | 121.602539 | 56.956055 | 2.14× |
| record/set-small-1000 | 0.530000 | 0.549047 | 0.534730 | 1.03× |
| record/set-medium-1000 | 6.440000 | 6.685771 | 5.689575 | 1.18× |
| record/set-large-1000 | 9.970000 | 10.417155 | 10.044271 | 1.04× |
| st/ref-write-10000 | 290.970000 | 269.451813 | 7.148641 | 37.69× |
| st/ref-modify-10000 | 292.310000 | 277.514969 | 20.672283 | 13.42× |
| effect/ref-write-10000 | 404.810000 | 388.415375 | 190.794281 | 2.04× |
| effect/ref-modify-10000 | 539.510000 | 513.634125 | 192.986328 | 2.66× |
| File I/O | 442703.250000 | 513738.000000 | 475960.000000 | 1.08× |
| STArray Operations | 0.540000 | 0.333000 | 0.292000 | 1.14× |
| String Operations | 507.080000 | 762.625000 | 559.041000 | 1.36× |
| Aff Operations | 11009.250000 | 10825.625000 | 10500.209000 | 1.03× |
| Parallelism | 1182150.080000 | 998488.583000 | 965397.125000 | 1.03× |

## Protocole et validation

Les suites ordinaires utilisent le runner officiel, trois processus par version,
ordre alterné/rotatif, GOGC=800, PGO désactivé, médiane des résultats par ligne.
Le total core est la somme des 14 médianes ; les WIP n’ont pas de score agrégé.
Les diagnostics JSON utilisent six paires, GOGC=100/GOMAXPROCS=1, deux warmups et
cinq échantillons par phase ; leurs sorties conservées sont vérifiées hors timing.
Les phases ne sont pas additionnées. Array Indexing vérifie les six cas à chaque
processus avec trois seeds, dix batches et GOGC=800/GOMAXPROCS=1.
Toutes les constructions ont fini avant les comparaisons appariées.

- `./bin/go/run --wip-included` a réussi avec le compilateur natif reconstruit.
- 87 tests ciblés du compilateur ; 43 tests PBO ; runtime normal et race detector.
- Compatibilité ST : 54 combinaisons anciennes/nouvelles, alias et snapshots,
  références imbriquées, zéro allocation pour les écritures nouvelles répétées.
- FFI Effect.Ref : différé/rejeu, callback, exception, newWithSelf et concurrence
  testés via le bridge réel sous race detector ; mutex conservés.
- Unicode : oracles JS/Go, surrogates isolés/appariés, WTF-8 et octets malformés.
  Concaténation comparée à l’append strict sur 10 000 séquences aléatoires.
- Traverse évalue tous les callbacks dans l’ordre, même après Nothing ; tri stable,
  entrées immuables et ordre des convertisseurs vérifiés par exécution du Go généré.
- Bootstrap : préparation des dépendances Go à leurs versions déclarées ajoutée ;
  11 tests vérifient aussi la conservation du binaire précédent en cas d’échec.

## Limites et pistes non retenues

RBTree domine désormais le core ; il alloue déjà les 100 000 nœuds nécessaires.
Supprimer d’autres stores demande des preuves d’alias/provenance. StateMonad garde
des closures et records intermédiaires : obtenir des workers à deux arguments
exige une analyse de stricte évaluation, les types seuls ne l’autorisent pas.
TCO/modulo et scalarisation des records nécessitent des preuves de bornes ou de
durée de vie. Aucun gain supplémentaire n’est promis pour ces pistes.

L’index TAST est déjà rempli à 94–99 %. Un cache Ann borné ne gagne que 2,24 %
dans son probe apparié ; il n’est pas intégré. Les entrées File I/O/Parallelism
subissent la variabilité système ; leurs écarts ne sont pas attribués au patch.
Aff Operations inclut un délai explicite de 10 ms. Les temps tiny sont conservés
en pleine précision dans l’archive, malgré l’arrondi d’affichage du README.

String Operations extended conserve regexp2 et les sémantiques JavaScript.
Son ancien temps README (507,08 µs) utilisait un moteur différent : les gains
présents sont comparés au binaire courant préservé, pas revendiqués contre cet
ancien moteur. Les références historiques restent visibles dans le tableau.

Les écritures Effect.Ref stockées dans un let gardent leur action ; seule une
application immédiatement exécutée est fusionnée. Les conversions autour d’un
Value opaque restent présentes. La compatibilité ST couvre la FFI publiée,
pas une FFI tierce inspectant directement un payload interne.

## Artefacts

Pour reconstruire ce compilateur en conservant la correction PBO locale :

```sh
cd ../gopurs/gopurs
GOPURS_PBO_DIR="$PWD/../../purescript-backend-optimizer" \
GOPURS_LIBRARY_ROOT="$PWD/.." \
GOGC=100 GOMEMLIMIT=8GiB GOMAXPROCS=4 npm run build:native
cd ../../altbak.pub
GOPURS_GO_PACKAGES=0 GOPURS_GO_PACKAGE_OVERLAY=0 ./bin/go/run --wip-included
```

Le bootstrap par défaut utilise les versions épinglées ; l’override PBO est donc
nécessaire tant que la correction locale n’est pas publiée et réépinglée.

[Archive chiffrée et empreintes](2026-11-21-go-significant-gains.json).
Campagne complète : `var/benchmark/go-ratio-gains-20261121/` ; probes antérieurs :
`var/benchmark/go-opportunities-20261121/`. Les caches de compilation temporaires
ont été supprimés pour libérer le disque ; exécutables, Go généré et preuves sont
conservés. Le build natif a été repris après préparation de regexp2 v1.11.5, puis
installé atomiquement. Les suites core, extended et JSON utilisent ce binaire
natif ; Array Indexing conserve le host JS imposé par son runner, reconstruit
depuis les mêmes sources et le même PBO local.
