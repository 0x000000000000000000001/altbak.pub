# List Processing : somme filtrée en temps constant

Mesures du 23 novembre 2026. Seconde optimisation de List Processing, après
la [fusion des listes](2026-11-23-go-list-processing.md). La baseline inclut
aussi les deux optimisations de Prime Sieve et celle d'Array Processing.
Les [données brutes](2026-11-23-go-list-processing-scalar.json) conservent les
manifestes, empreintes, sorties, diagnostics, tests et deux campagnes complètes.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 0,424458 µs | 0,025805 µs |
| Allocations du calcul | 0 | 0 |
| Octets alloués par calcul | 0 | 0 |
| Résultat vérifié | 202 950 | 202 950 |

Accélération **16,449×**, soit **93,92 % de temps en moins**.
Les temps incluent les adaptateurs Effect, l'entrée opaque et le sink du
harness. Les allocations sont mesurées séparément sur `sumEvens(900)`.

Échantillons avant : 0,423804 ; 0,425687 ; 0,424458 µs.
Échantillons après : 0,025784 ; 0,026144 ; 0,025805 µs.

Une seconde campagne sur les mêmes binaires donne
**0,425307 → 0,025863 µs**, soit **16,445×**. La première campagne reste la
mesure principale ; la confirmation est archivée séparément. La campagne
isolée donne **0,433858 → 0,025905 µs**, soit **16,748×**.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **8,711060 µs / 174,22× C**.
Face au C historique publié de **0,05 µs**, le nouveau ratio est
**≈ 0,516× C**, contre **≈ 8,489× C** avant cette seconde optimisation.
Face au JS historique de **9,588542 µs**, il est **≈ 0,00269× JS**.
Ces références C et JS ne sont pas remesurées ; le README reste inchangé.
List Processing sort donc des priorités locales au seuil de 2×.

## Goulot et preuve expérimentale

La première optimisation avait supprimé les listes intermédiaires et leurs
allocations. Le Go parcourait encore tous les entiers de 900 à 1, testait le
modulo et ajoutait les valeurs retenues à un accumulateur Int32.

Un prototype isolé a comparé cette boucle, un parcours des seules valeurs
retenues et une somme arithmétique directe. Avec entrée chargée par opération
atomique et fonctions prototypes non inlinables, les trois mesures de la
boucle générée sont proches de **404 ns**, celles du parcours espacé entre
**198 et 257 ns**, et celles de la formule proches de **1,68 ns**.
Ces timings de diagnostic ne sont pas ceux du harness ; le gain applicatif
retenu est celui de la suite complète ci-dessus. Deux premières mesures avec
entrée constante ou prototype inlinable ont été écartées.

Avant toute modification du compilateur, le prototype a passé 10 000 cas
aléatoires déterministes, six cas extrêmes et une comparaison avec le vrai
ListOps sur les 1 030 bornes de -4 à 1 025. Un oracle BigInt vérifie les
grandes sommes sans arrondi flottant. Les diagnostics `go test` servent aussi
à contrôler les allocations ; leurs temps ne se substituent pas aux mesures
des exécutables figés.

## Changement du compilateur

`Gopurs.FilteredSumLoops`, raccordé à `AffineLoops.wrap`, reconnaît la
récurrence complète et pure suivante, avec deux paramètres Int64 :

```text
si i < borne alors acc
sinon self(i - 1, si i mod module == résidu alors acc + i sinon acc)
```

La borne et le module doivent être des constantes positives Int32, et le
résidu compris entre zéro inclus et le module exclu. Les deux ordres de
paramètres et les opérandes commutés de l'égalité/addition sont reconnus.
Aucun autre appel, effet, modification d'état ou prédicat n'est admis.

À l'exécution, le compteur doit être compris entre la borne et MaxInt32,
et le seed doit être représentable en Int32. Le chemin rapide calcule :

```text
premier = borne + (résidu - borne % module + module) % module
si premier > i : retourner acc
nombre = (i - premier) / module + 1
dernier = premier + (nombre - 1) * module
retourner Int32(acc + nombre * (premier + dernier) / 2)
```

Les opérations intermédiaires sont émises en Go `int64`, y compris les
constantes : le premier candidat peut dépasser MaxInt32. Si une valeur est
retenue, `nombre*(premier+dernier) ≤ 2*MaxInt32² < MaxInt64` ; la division
par deux est exacte et l'ajout du seed reste sûr. La conversion finale
reproduit les additions modulaires Int32. `runtime.IntMul` ne convient pas
ici, car il reproduit les arrondis de la multiplication JavaScript.

La formule est isolée dans une fonction Go immédiatement appelée afin de
préserver les noms locaux. La boucle originale reste le repli pour toute
entrée non admise. La transformation ne connaît ni ListOps ni la limite 900 ;
elle s'applique également à une récurrence scalaire sans listes. Le benchmark
PureScript, le runtime et le harness restent identiques.

## Validation

- 9 tests ciblés : 4 pour cette passe, 3 pour AffineLoops, 2 pour Int32Loops.
- Le Go issu des tests IR passe 658 cas sémantiques et 49 refus de garde :
  paramètres permutés, appels locaux/globaux, sommes au-delà de 2^53, premier
  candidat hors Int32, module égal à 1 ou MaxInt32, valeurs larges et replis.
- Fixture compilée par les hôtes JS et natif, en Release et Dev :
  139 résultats par profil contre un oracle indépendant BigInt.
- Prédicat `mod 5 == 3`, petites plages proches de MaxInt32, récurrence sans
  listes, cas historiques négatifs, partage des listes et ordre des folds.
- Par hôte Release : 70 seeds Int64 hors Int32 et 35 contrôles d'allocations
  sur cinq fonctions ; tous restent sans allocation.
- Assertions sur le Go réel confirmant la formule pour trois fonctions
  distinctes, dont la récurrence sans listes.
- Vrai ListOps : 1 030 petites bornes, cinq bornes Int64 hors Int32 à exécution
  courte, puis sept grandes bornes jusqu'à MaxInt32 sur le candidat.
- Les 14 oracles de la suite complète passent dans les douze processus des
  deux campagnes. Compilateur natif reconstruit et installé.

## Contrôle des autres benchmarks

Les **307 modules CoreFn typés sont identiques**. Le Go généré ne change
que dans **`Call_Test_ListOps_sumEvens`**. Voici les écarts des autres tests,
un signe positif indiquant un ralentissement :

| Test | Campagne principale | Confirmation |
| --- | ---: | ---: |
| AST Evaluation | −1,29 % | −0,03 % |
| Fibonacci | +0,00 % | −0,50 % |
| Tail Call Optimization | +1,17 % | +1,08 % |
| Deep Record Updates | +0,02 % | +0,27 % |
| Ackermann | +1,21 % | +2,87 % |
| Church Numerals | −2,46 % | −2,01 % |
| Prime Sieve | +0,01 % | +0,16 % |
| Red-Black Tree | −0,99 % | +1,24 % |
| Polymorphism | +0,20 % | +0,73 % |
| State Monad | −0,16 % | −0,49 % |
| Lazy Evaluation | −1,28 % | −1,25 % |
| Array Processing | +2,07 % | +0,46 % |
| RowToList | −0,12 % | +0,87 % |

La confirmation a été lancée pour contrôler les +2,07 % d'Array Processing ;
cet écart n'est pas retrouvé à la même amplitude. Les variations restantes,
notamment Ackermann, sont conservées sans leur attribuer de cause établie.
Cette campagne n'explique pas les anciens écarts Prime/RBTree observés pendant
l'optimisation Array Processing.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Les deux générations finales utilisent l'hôte natif, Release, sans PGO,
`GOGC=800`, et `GOMAXPROCS` par défaut (14 dans les diagnostics).
La baseline gopurs est `a355e98d05831b95eb9b5c59d7d4b542af575133` avec les
quatre campagnes précédentes appliquées ; les empreintes archivées identifient
cet état et le candidat. Fork PureScript :
`41fb570fa58e3698646ea6e6179571348e4acd12`.

Construire chaque workspace avec la version correspondante du compilateur :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ListOps --build-only \
  --build-dir /tmp/gopurs-listops-scalar-20261123-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ListOps --build-only \
  --build-dir /tmp/gopurs-listops-scalar-20261123-after
```

Dans chaque workspace, copier `output/main/main.go` dans
`output/coremain/main.go` et remplacer uniquement `Get_AppX_main()` par
`Get_App_main()`, puis construire `core-benchmark` :

```sh
go -C /tmp/gopurs-listops-scalar-20261123-before/output build -pgo=off \
  -o /tmp/gopurs-listops-scalar-20261123-before/core-benchmark ./coremain
go -C /tmp/gopurs-listops-scalar-20261123-after/output build -pgo=off \
  -o /tmp/gopurs-listops-scalar-20261123-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test ListOps --core \
  --before /tmp/gopurs-listops-scalar-20261123-before \
  --after /tmp/gopurs-listops-scalar-20261123-after \
  --output /tmp/list-scalar-core-comparison-new
```

Omettre `--core` et choisir un autre répertoire de sortie pour la mesure
isolée. Le runner alterne AB/BA/AB : trois processus indépendants par variante,
trois échauffements locaux, calibration, minimum de dix batches, puis médiane
des trois processus. La suite complète ajoute trois échauffements globaux.
Compilations et tests étaient arrêtés pendant les mesures finales.
