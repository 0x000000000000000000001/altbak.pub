# Prime Sieve : construction directe des listes filtrées

Mesures du 23 novembre 2026. Le benchmark PureScript et son harness restent
inchangés. L'état avant inclut déjà l'optimisation de List Processing.
Les [données brutes](2026-11-23-go-prime-sieve.json) conservent les sorties,
manifestes, empreintes des binaires et sources, profils et prototypes.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 70,447594 µs | 39,927246 µs |
| Allocations du calcul | 10 533 | 5 470 |
| Octets alloués par calcul | 254 296 | 130 504 |
| Résultat vérifié | 21 536 | 21 536 |

Accélération **1,764×**, soit **43,32 % de temps en moins**, avec **48,07 %
d'allocations en moins**. Les temps incluent les adaptateurs Effect, l'entrée
opaque et le sink du harness. Le diagnostic d'allocations appelle directement
les fonctions générées `range → sieve → sumList`, hors adaptateurs Effect.

Échantillons avant : 70,447594 ; 70,551594 ; 70,190918 µs.
Échantillons après : 39,932617 ; 39,927246 ; 39,561199 µs.

La campagne isolée donne séparément **71,158531 → 39,900879 µs** (1,783×).
La mesure retenue pour le README est celle de la suite complète, avec ses
trois échauffements globaux et ses quatorze tests.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **73,303547 µs / 67,25× C**.
La reproduction initiale du test isolé donnait 68,784180 µs. La comparaison
finale utilise les deux variantes figées de la même campagne alternée.
Face à la référence C publiée de **1,09 µs**, le nouveau ratio est
**≈ 36,63× C** ; le ratio avant mesuré localement est ≈ 64,63× C.
Le C historique n'a pas été remesuré et le README n'est pas republié ici.

## Goulot observé

Le Go était déjà typé avec des `Cons[int64]`. Cependant, chaque filtre
construisait un accumulateur inversé puis reconstruisait une seconde liste
pour rétablir l'ordre. Le profil attribue **91,64 % des objets alloués** à
ces deux constructions. L'allocateur représente **39,86 %** des échantillons
CPU cumulés, et `gopurs_runtime.Apply` **10,87 %**.

Pour l'entrée de 500, les 95 filtrages retiennent au total 4 778 éléments.
Leurs deux constructions produisaient 9 556 maillons. Il reste une seule
construction de 4 778 maillons après transformation, ainsi que les listes
du range et du résultat du crible. Des fermetures locales disparaissent aussi.

Un prototype Go conservant le callback boxé a d'abord mesuré environ
**69,75 → 40,59 µs**, avec les mêmes comptes d'allocations que la version
intégrée. Une variante expérimentale à prédicat natif descendait à environ
32,53 µs ; cette seconde transformation n'est pas intégrée. Elle constitue
une piste distincte, avec ses propres obligations de spécialisation.

Le C de référence utilise des tests de divisibilité jusqu'à la racine carrée,
alors que le programme PureScript conserve son crible à listes strictes.
Le ratio inclut cette différence d'algorithme ; le benchmark n'a pas été
remplacé par le calcul impératif de référence.

## Changement du compilateur

`Gopurs.ListFilter.worker`, appelé par `ModuleWorkers`, reconnaît le corps
complet d'un filtre avec accumulation inversée et reverse final. La signature
concrète et les métadonnées TAST établissent le layout : constructeur avec
champs `Int`/liste et alternative vide représentée par `nil`.

Le worker construit directement une liste neuve dans l'ordre source. Il
maintient une tête et une queue privées, et n'affecte que les liens de ses
propres nœuds fraîchement alloués. Il ne modifie pas l'entrée, même lorsqu'elle
est partagée ou capturée par le prédicat. Chaque callback garde son ABI boxé,
son nombre d'invocations et son ordre. Les constructeurs conservent `Rc=1`.

La reconnaissance est structurelle, indépendante des noms du benchmark et de
la limite 500. Les types quantifiés, payloads autres que Int, layouts inconnus,
reverse incomplets et autres motifs non prouvés conservent le chemin ordinaire.
Les noms internes sont isolés dans une fonction à paramètres explicites.

## Validation

- 19 tests IR de reconnaissance et de refus ; revue indépendante du layout,
  des callbacks et de la mutation limitée à la sortie privée.
- Oracle indépendant des nombres premiers pour 605 entrées, de -4 à 600.
- Fixture renommée : 11 entrées et 110 observations, listes vides, filtres
  retenant tout/rien, diviseurs nuls/négatifs, captures et partage. Un callback
  avec effets journalise les éléments visités et l'empreinte de l'entrée.
- 25 contrôles d'allocations, avec exactement un nœud supplémentaire par
  élément retenu et un coût constant borné, entrée construite hors mesure.
- Tests d'intégration réussis avec les hôtes JS et natif, en Release et Dev.
- Six processus de la suite complète : les 14 oracles sont corrects à chaque
  exécution. List Processing conserve son gain : 0,428607 → 0,427517 µs.
- Les différences Go se limitent à `Test.Primes` : worker de filtre remplacé
  et renumérotation des identifiants locaux de `act`.
- Compilateur natif par défaut reconstruit et installé.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1 ; Go 1.27.0 darwin/arm64,
Node 24.8.0. Les deux générations finales utilisent l'hôte **natif**, Release,
PGO désactivé, `GOGC=800`, `GOMAXPROCS` par défaut (14 dans les diagnostics).
La baseline gopurs est le commit `a355e98d05831b95eb9b5c59d7d4b542af575133`
avec la [fusion de List Processing](2026-11-23-go-list-processing.md) déjà
appliquée. Les empreintes distinguent cet état du commit seul et du candidat.

Construire un workspace avec chaque version du compilateur, avant les mesures :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-20261123-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-20261123-after

python3 bin/benchmark/list-processing.py --test Primes \
  --before /tmp/gopurs-primes-20261123-before \
  --after /tmp/gopurs-primes-20261123-after \
  --output /tmp/prime-sieve-isolated-comparison-new
```

Le runner reste compatible avec ListOps par défaut. Pour la suite complète,
construire aussi `core-benchmark` depuis le même Go, en remplaçant uniquement
`Get_AppX_main()` par `Get_App_main()` dans une copie du point d'entrée :

```sh
for variant in before after; do
  workspace="/tmp/gopurs-primes-20261123-$variant"
  mkdir -p "$workspace/output/coremain"
  sed 's/Get_AppX_main()/Get_App_main()/' "$workspace/output/main/main.go" \
    > "$workspace/output/coremain/main.go"
  go -C "$workspace/output" build -pgo=off -o "$workspace/core-benchmark" ./coremain
done

python3 bin/benchmark/list-processing.py --test Primes --core \
  --before /tmp/gopurs-primes-20261123-before \
  --after /tmp/gopurs-primes-20261123-after \
  --output /tmp/prime-sieve-core-comparison-new
```

Dans cette campagne les deux entrées appartiennent au même package Go ; le
runner vérifie les empreintes des binaires, la configuration et les sources.
Ordre **AB / BA / AB**, sans compilation concurrente. Chaque processus garde
la calibration et le minimum de dix batches ; le résultat est la médiane
des trois minima. Le mode core vérifie tous les oracles et conserve les trois
échauffements globaux. Les sources des diagnostics et leurs résultats sont
archivés dans le JSON pour reproduire le comptage d'allocations.
