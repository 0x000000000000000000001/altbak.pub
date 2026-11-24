# List Processing : fusion des listes intermédiaires

Mesure du 23 novembre 2026. Sources PureScript et harness inchangés.
Les [données brutes](2026-11-23-go-list-processing.json) contiennent les
échantillons, sorties vérifiées, manifestes, empreintes et profils.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus, harness altbak | 8,329569 µs | 0,429392 µs |
| Allocations du calcul `sumEvens(900)` | 1 352 | 0 |
| Octets alloués par calcul | 32 432 | 0 |
| Résultat vérifié | 202 950 | 202 950 |

Accélération **19,40×**, soit **94,84 % de temps en moins**.
Les temps incluent les adaptateurs `Effect`, l'entrée opaque et le sink du
harness. Le comptage d'allocations appelle directement la fonction générée ;
il exclut ces adaptateurs.

Échantillons avant : 8,329916 ; 8,317261 ; 8,329569 µs.
Échantillons après : 0,429392 ; 0,428755 ; 0,432672 µs.

La mesure isolée `bin/go/run --test ListOps` donne séparément
**8,232727 → 0,280910 µs**, soit **29,31×**. Ses échantillons avant sont
8,192688 ; 8,232727 ; 8,310709 µs, et après 0,288314 ; 0,278667 ; 0,280910 µs.
Les deux périmètres sont conservés dans les données brutes ; le résultat
retenu pour la comparaison au README est celui de la suite complète.

La baseline officielle du [README](../../README.md#go), détaillée dans
[la campagne du même jour](2026-11-23-go.json), est **8,711060 µs / 174,22× C**.
La mesure initiale de cette tâche, avant toute modification, était
8,683390 µs (médiane de trois processus). La campagne alternée finale donne
8,329569 µs dans la suite complète pour le même code généré avant ; le gain
annoncé utilise exclusivement les deux variantes de cette campagne alternée.

Rapporté à la référence C publiée de **0,05 µs**, le résultat après donne
**environ 8,59× C**. Le C n'a pas été remesuré. La campagne complète conserve
les trois échauffements globaux et les quatorze cas du README. Le ratio avant
mesuré localement est 166,59×, contre 174,22× dans la baseline publiée.
Cette note ne remplace pas automatiquement la cellule publiée dans le README.

## Goulot et transformation

Le Go généré utilisait déjà des maillons typés `Cons[int64]` et des boucles
issues de la récursion terminale. Il matérialisait néanmoins une liste de
900 éléments, puis une seconde liste de 450 éléments pairs avant de sommer.
Le profil d'allocations attribue **99,86 % des objets** aux deux constructions.
Le profil CPU attribue **53,36 % des échantillons cumulés** à l'allocateur
`mallocgcSmallScanNoHeaderSC3`.

Des mesures séparées situaient la construction vers 5,3 µs, le filtre vers
3,0 µs, et le fold vers 0,43 µs. Un prototype isolé supprimant les listes,
mais conservant `IntMod`, `IntSub` et `IntAdd`, a confirmé le gain avant
l'implémentation de la passe.

La nouvelle passe `Gopurs.ListFusion` reconnaît dans l'IR final :

1. Un ADT local à constructeur vide et constructeur `[Int, liste]`, décrit
   par `dataDecls`, avec des signatures concrètes.
2. Un producteur descendant qui ajoute chaque entier en tête.
3. Un filtre fermé qui ajoute en tête les éléments retenus.
4. Une réduction entière par addition consommant immédiatement le résultat.

Elle remplace cette composition par une boucle scalaire. Les deux inversions
successives permettent de conserver l'ordre de visite du producteur. Les
opérations PureScript sur les entiers et leur débordement à 32 bits restent
inchangés ; les optimisations Go existantes peuvent ensuite simplifier la
boucle. Aucun nom de benchmark, aucune taille fixe ni résultat précalculé
n'intervient dans la reconnaissance.

Les fonctions originales restent disponibles pour les autres appels. La passe
refuse notamment les effets, appels opaques, prédicats capturant une valeur,
types quantifiés, bornes inférieures dynamiques, pas autres que 1 et formes
de réduction non prouvées. Elle ne modifie aucun maillon partagé.

## Validation

- 21 tests de reconnaissance et de refus de la passe, dont deux fusions dans
  un même binding et l'immutabilité de l'IR d'entrée.
- Oracle arithmétique indépendant pour 1 030 entrées de ListOps, de -4 à 1 025.
- Fixture avec noms d'ADT et fonctions différents : 110 observations par
  profil, plages vides/négatives, filtres pairs/impairs/multiples de trois,
  débordements Int32, réduction non commutative et listes partagées.
- Zéro allocation dans les trois variantes admissibles sur sept tailles.
- Fixture validée en Release et Dev avec les deux hôtes du compilateur, JS
  et natif parallèle ; le binaire natif par défaut a été reconstruit.
- Les 14 résultats attendus de la suite principale sont identiques avant/après.
- Comparaison de tous les fichiers Go générés : seul `Test_ListOps.go` change.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1 ; Go 1.27.0 darwin/arm64,
Node 24.8.0. Profil Release, PGO désactivé, `GOGC=800`, `GOMAXPROCS` par défaut
(14 dans les diagnostics Go). La campagne principale utilise l'hôte JS du
compilateur pour les deux générations.

Révisions initiales : gopurs `a355e98d05831b95eb9b5c59d7d4b542af575133`,
altbak `8733d0a6124028649e64fa7dd8eabc4dcb213d49`, compilateur PureScript typé
`41fb570fa58e3698646ea6e6179571348e4acd12`.

Construire un workspace avant la modification, puis un autre après :

```sh
GOPURS_JS=1 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ListOps --build-only \
  --build-dir /tmp/gopurs-listops-20261123-before

GOPURS_JS=1 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ListOps --build-only \
  --build-dir /tmp/gopurs-listops-20261123-after

python3 bin/benchmark/list-processing.py \
  --before /tmp/gopurs-listops-20261123-before \
  --after /tmp/gopurs-listops-20261123-after \
  --output /tmp/list-processing-comparison-new
```

Pour la confirmation dans la suite complète, le même Go généré expose déjà
`App.main`. Construire les deux points d'entrée avant de lancer les mesures :

```sh
for variant in before after; do
  workspace="/tmp/gopurs-listops-20261123-$variant"
  mkdir -p "$workspace/output/coremain"
  sed 's/Get_AppX_main()/Get_App_main()/' "$workspace/output/main/main.go" \
    > "$workspace/output/coremain/main.go"
  go -C "$workspace/output" build -pgo=off -o "$workspace/core-benchmark" ./coremain
done

python3 bin/benchmark/list-processing.py --core \
  --before /tmp/gopurs-listops-20261123-before \
  --after /tmp/gopurs-listops-20261123-after \
  --output /tmp/list-processing-core-comparison-new
```

Le script vérifie les binaires figés, la même configuration et les mêmes
sources. Il exécute les variantes dans l'ordre **AB / BA / AB**, sans
compilation concurrente. Chaque processus fait trois échauffements locaux,
une calibration puis dix batches ; la statistique finale est la médiane des
trois minima. Le mode `--core` ajoute les trois échauffements globaux de la
suite et vérifie ses quatorze oracles ; le mode isolé vérifie l'oracle 202 950.

Un premier build après et un test d'intégration ont échoué faute d'espace
disque, avant mesure. Après nettoyage du cache de compilation Go régénérable,
ils ont été repris. Aucun échec n'entre dans les statistiques.
