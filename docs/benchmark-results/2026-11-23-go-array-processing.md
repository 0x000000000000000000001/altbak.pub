# Array Processing : fusion de l'intervalle, du filtre et de la somme

Mesures du 23 novembre 2026. L'état avant inclut déjà les optimisations de
List Processing et Prime Sieve. Les [données brutes](2026-11-23-go-array-processing.json)
conservent les sorties, manifestes, empreintes, diagnostics et profils.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 2,637736 µs | 0,363295 µs |
| Allocations du calcul | 6 | 0 |
| Octets alloués par calcul | 47 000 | 0 |
| Résultat vérifié | 202 950 | 202 950 |

Accélération **7,261×**, soit **86,23 % de temps en moins**.
Les temps incluent les adaptateurs Effect, l'entrée opaque et le sink du
harness. Le diagnostic d'allocations appelle directement la fonction générée
`Call_Test_ArrayOps_sumEvens(900)`, hors adaptateurs.

Échantillons avant : 2,637736 ; 2,629588 ; 2,671438 µs.
Échantillons après : 0,363295 ; 0,359172 ; 0,363635 µs.

La campagne isolée donne séparément **2,602610 →
0,368888 µs** (7,055×). La comparaison retenue
est celle de la suite complète, avec ses trois échauffements globaux et ses
quatorze tests.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **2,673269 µs / 53,47× C**.
Face à la référence C publiée de **0,05 µs**, le nouveau ratio est
**≈ 7,27× C** ; le ratio avant mesuré localement
est ≈ 52,75× C. Le C historique n'a pas été remesuré et le README
n'est pas republié ici.

## Goulot observé

Le Go généré construisait successivement un `[]int64` pour l'intervalle,
sa copie boxée, le tableau filtré natif avec croissance de capacité, puis
une nouvelle copie boxée pour la somme. Le calcul pour 900 éléments allouait
**47 000 octets et 6 objets**. Le profil mémoire attribue environ **99,95 %**
de l'espace alloué à ces constructions et copies ; les deux conversions
vers des tableaux boxés représentent ensemble environ **69,34 %**.

Un prototype Go a d'abord validé l'élimination des tableaux : **2,544 →
0,3495 µs**, avec **0 octet / 0 allocation**. Cette mesure directe exclut
les adaptateurs Effect et le harness ; elle ne remplace pas l'avant/après
de la suite complète présenté plus haut. Le profil CPU contient surtout
des échantillons d'attente du runtime sur cet hôte ; les conclusions sur
le goulot reposent sur le profil mémoire et les mesures d'allocations.

## Changement du compilateur

`Gopurs.ArrayRangeFold.emit`, appelé avant `ScalarArrayFold` par `CallExprs`,
reconnaît le pipeline immédiatement consommé
`Array.range → Array.filter → foldl intAdd`. Les fonctions doivent être les
builtins exactement qualifiés et les types doivent établir un tableau d'Int,
un prédicat `Int → Boolean` et une somme entière. La reconnaissance ne dépend
ni du nom du benchmark, ni de la limite 900.

Le résultat est une boucle sans tableau intermédiaire :

```go
for {
    if predicate(current) {
        total = runtime.IntAdd(total, current)
    }
    if current == last {
        return total
    }
    current += step
}
```

Les bornes sont inclusives, avec `step` valant 1 ou -1. L'égalité terminale
est testée avant l'incrément ; seule la somme applique le débordement Int32
de PureScript. Les captures sont évaluées une fois, dans l'ordre initial.
Le prédicat conserve son ABI, son nombre d'appels et leur ordre. Intercaler
une addition entière totale entre ses appels n'introduit aucun effet
observable. Une réduction arbitraire, notamment avec effets, reste sur le
chemin existant.

Les littéraux Int et les variables dont le contexte natif issu du TAST
établit `int64` restent reconnaissables lorsque PBO a supprimé leur annotation
immédiate. Une annotation explicite incompatible ou une information inconnue
ne suffit pas à autoriser la fusion. Les tableaux partagés et les motifs
non reconnus gardent leur traitement habituel.

Les bornes hors Int32 utilisent le chemin matérialisé d'origine, en réutilisant
les captures. Le prédicat natif conserve sa signature dans ce repli : aucun
wrapper ne provoque d'allocation sur le chemin rapide. Le benchmark PureScript,
son résultat et son harness restent inchangés.

## Validation

- 21 tests IR de reconnaissance et de refus, avec compilation et exécution
  Go utilisant les véritables traducteurs de types et d'intrinsèques.
- Prédicats boxés, `func(Value) Value` et `func(int64) bool` ; captures
  évaluées une seule fois, plages ascendantes, descendantes et singletons,
  bornes proches des extrêmes Int32 et Int64.
- Fixture renommée : 15 entrées, 75 observations par exécution, tableaux
  partagés, filtres retenant tout/rien, diviseurs nuls/négatifs, sommes avec
  débordement, factories et callbacks avec effets journalisés.
- 17 contrôles de zéro allocation par hôte Release, dont une lambda capturant
  son diviseur. 90 comparaisons avec un calcul matérialisé, comprenant
  81 entrées non canoniques et 63 passages avec des bornes hors Int32.
- Tests d'intégration réussis avec les hôtes JS et natif, en Release et Dev,
  et 3 tests voisins sur les folds scalaires et l'ordre des intrinsèques.
- Oracle indépendant du résultat d'Array Processing pour 3 008 limites :
  de -1 000 à 2 000, plus sept grandes valeurs positives et négatives.
- Six processus de la suite complète, avec les 14 oracles corrects à chaque
  exécution. Compilateur natif par défaut reconstruit et installé.

## Contrôle des autres benchmarks

Les empreintes des **307 modules CoreFn typés** sont identiques avant/après.
Le seul changement du Go généré est `Call_Test_ArrayOps_sumEvens`.
List Processing conserve son gain. Deux autres tests présentent néanmoins
un ralentissement mesuré, confirmé dans une seconde campagne indépendante :

| Test | Première campagne avant → après | Écart | Écart du contrôle |
| --- | ---: | ---: | ---: |
| List Processing | 0,430436 → 0,428904 µs | -0,36 % | 0,02 % |
| Prime Sieve | 39,494141 → 40,949547 µs | 3,69 % | 3,35 % |
| Red-Black Tree | 9605,062500 → 9969,959000 µs | 3,80 % | 2,95 % |

Un écart positif indique un ralentissement. La cause de ces écarts sur Prime
Sieve et Red-Black Tree n'est pas établie ; leur Go est inchangé. Ces régressions
mesurées sont conservées dans le bilan et leur cause reste à analyser.
Les douze processus des deux campagnes valident tous les résultats calculés.

Cette seconde campagne utilise les mêmes binaires et le même ordre AB/BA/AB.
Elle donne pour Array Processing **2,649465 →
0,359336 µs** (7,373×). La première campagne
reste la mesure principale ; toutes les sorties du contrôle sont archivées.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1 ; Go 1.27.0 darwin/arm64,
Node 24.8.0. Les deux générations finales utilisent l'hôte **natif**, Release,
PGO désactivé, `GOGC=800`, `GOMAXPROCS` par défaut (14 dans les diagnostics).
La baseline gopurs est le commit `a355e98d05831b95eb9b5c59d7d4b542af575133`
avec les optimisations de List Processing et Prime Sieve déjà appliquées.
Les empreintes distinguent cet état du commit seul et du candidat. Le fork
PureScript est au commit `41fb570fa58e3698646ea6e6179571348e4acd12`.

Construire un workspace avec chaque version du compilateur avant les mesures :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ArrayOps --build-only \
  --build-dir /tmp/gopurs-arrayops-20261123-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ArrayOps --build-only \
  --build-dir /tmp/gopurs-arrayops-20261123-after

python3 bin/benchmark/list-processing.py --test ArrayOps \
  --before /tmp/gopurs-arrayops-20261123-before \
  --after /tmp/gopurs-arrayops-20261123-after \
  --output /tmp/array-processing-isolated-comparison-new
```

Pour mesurer la suite complète à partir des mêmes modules Go générés, copier
`output/main/main.go` dans `output/coremain/main.go`, remplacer uniquement
`Get_AppX_main()` par `Get_App_main()`, puis construire chaque binaire :

```sh
go -C /tmp/gopurs-arrayops-20261123-before/output build -pgo=off \
  -o /tmp/gopurs-arrayops-20261123-before/core-benchmark ./coremain
go -C /tmp/gopurs-arrayops-20261123-after/output build -pgo=off \
  -o /tmp/gopurs-arrayops-20261123-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test ArrayOps --core \
  --before /tmp/gopurs-arrayops-20261123-before \
  --after /tmp/gopurs-arrayops-20261123-after \
  --output /tmp/array-processing-core-comparison-new
```

Le runner alterne AB/BA/AB, avec trois processus indépendants par variante.
Chaque test utilise trois échauffements locaux, une calibration, puis le
minimum de dix batches ; la suite complète ajoute trois échauffements
globaux. Le résultat final est la médiane des trois processus. Aucune
compilation ni suite de tests ne tourne en parallèle des mesures finales.
