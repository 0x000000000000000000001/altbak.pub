# Validation de la réutilisation ADT Go — 16 septembre 2026

Les sources du backend et le README d'altbak ne sont pas modifiés par cette
préparation. Les fichiers de ce dossier sont des preuves et outils de validation.

## Référence conservée

`before-pure` et `before-fficc` conservent les binaires existants, leur manifeste,
les sources de l'application, les 388 fichiers Go nécessaires, le module Go,
les artefacts CoreFn de RBTree et les résultats antérieurs. Les empreintes sont
dans `saved-sha256.json` et `preservation.json`. Les binaires correspondent aux
SHA du manifeste avant copie. Leur provenance est celle des builds conservés,
pas celle du HEAD actuel de chaque compilateur.

Le README sauvegardé indique 24,02 ms compilé et 11,09 ms manuscrit. Le binaire
compilé présent avait pour résultat archivé 25,703 ms ; ce n'est donc pas
l'artefact établissant les 24,02 ms. Les mesures actuelles ne réécrivent pas cette
baseline historique.

Après coordination avec le responsable du build, six processus existants ont
été exécutés sans compilation simultanée : trois par mode, en ordre alterné.
Les 84 résultats numériques sont validés. Go : `go1.27.0 darwin/arm64` ;
`GOGC=800`, `GOWORK=off`, sans PGO, sans `PPROF`, `GOMAXPROCS` non fixé.
Le harnais de chaque application effectue trois suites de chauffe globales,
trois chauffes par test, puis retient le minimum de dix mesures.

| Mesure | Compilé conservé | Manuscrit conservé |
| --- | ---: | ---: |
| Total processus 1 | 25,41980 ms | 10,76954 ms |
| Total processus 2 | 25,45882 ms | 10,71759 ms |
| Total processus 3 | 25,74868 ms | 11,00841 ms |
| Somme des médianes par ligne | 25,45087 ms | 10,75403 ms |
| Médiane RBTree | 22,26379 ms | 8,43225 ms |

`baseline-measurements.json` contient les données, dates UTC et SHA ; les six
sorties brutes sont conservées à côté. Une comparaison finale devra alterner
à nouveau avant/après/manuscrit dans la même fenêtre, sans builds concurrents.

## Points d'entrée de compilation

Depuis altbak, avec le bundle gopurs déjà reconstruit et le créneau de build
coordonné :

```sh
GOGC=800 bin/go/run --test RBTree --expected 22 --build-only --build-dir scratch/gopurs-adt-reuse-validation-20260916/after-rbtree
GOGC=800 bin/go/run --build-only --build-dir scratch/gopurs-adt-reuse-validation-20260916/after-pure
```

`-c` reconstruirait aussi le backend : ne pas le demander en concurrence avec
un autre build du backend. Les workspaces ci-dessus sont isolés ; ils ne
changent pas les symlinks du projet. `--run-only` contrôle les SHA des entrées,
du bundle et du binaire : il refuse un artefact devenu périmé. La mesure d'un
ancien binaire archivé se fait directement avec son SHA et sa provenance,
sans modifier son manifeste pour contourner ce contrôle.

L'ancien `altbak/bin/go/test` modifie les liens de sortie, certaines sources et
les snapshots. Il n'est pas utilisé ici. Le runner moderne de gopurs est isolé.
Pour le nouveau test, depuis gopurs, après reconstruction du bundle :

```sh
./bin/test OwnedTrees --update-snapshots --keep-workspace
```

Le snapshot est nouveau ; le runner ne l'écrit qu'après compilation et
exécution réussies. La vérification suivante utilisera le mode normal.
La fixture `OwnedTrees.purs` emploie des noms génériques, des frontières
`inline never` et des entrées rendues dynamiques par `Effect.Ref`. Elle couvre
les quatre rotations, l'ordre, les couleurs, les hauteurs noires, les doublons,
des insertions ascendantes/descendantes/permutées, 32 snapshots conservés, un
sous-arbre retenu et un parent frais contenant un enfant déjà partagé. Les
fonctions de descente, équilibrage et recoloration sont top-level pour permettre
l'analyse des travailleurs spécialisés. Les appels publics restent persistants.

## Sonde Go indépendante

`harness/` contient les contrôles structurels sur le Go généré et une sonde
`runtime.MemStats` : allocations, octets cumulés, durée et profondeur. Une entrée
`buildFresh` sélectionne avant le `buildTree` public, et après le worker
`__gopurs_owned_buildTree_0` réellement choisi par le `act` généré. Les deux
commencent depuis `nil`. Le corps de `main.go` est identique entre les sondes ;
seul `entrypoints.go` s'adapte au point d'entrée de chaque version. Le staging
vérifie que le worker sélectionné figure bien dans le `act` généré.
La sonde peut aussi appeler le FFI manuscrit. Elle ne transmet jamais d'arbre
partagé à un worker de possession et conserve le résultat jusqu'à la fin de
la mesure. Les tests de persistance restent sur les fonctions publiques.

Depuis ce dossier, préparer une copie de la sonde sans compiler :

```sh
python3 stage-harness.py --generated before-pure/output --target probe-before-owned
python3 stage-harness.py --generated after-pure/output --target probe-after-owned
```

Puis, seulement pendant un créneau coordonné, depuis chaque `probe-*` :

```sh
GOWORK=off GOCACHE=/private/tmp/gopurs-adt-reuse-gocache go test -pgo=off -count=1 -v ./...
GOWORK=off GOCACHE=/private/tmp/gopurs-adt-reuse-gocache go build -pgo=off -o probe .
GOGC=800 ./probe -mode pure -n 100000 -iterations 5
GOGC=800 ./probe -mode fficc -n 100000 -iterations 5
```

La sonde chauffe trois appels puis force un GC avant chaque appel mesuré ; ce
protocole sert à comparer les allocations à paramètres égaux. Ses durées ne
sont pas interchangeables avec les minima du harnais publié. L'appel nominal
doit produire 22. Les checks Go supposent le layout typé actuel de RBTree ; un
changement de représentation demandera une adaptation explicite du harnais.

Après coordination, la sonde baseline a été compilée et exécutée. Les quatre
tests Go passent : rotations, historiques persistants (six séquences), sous-arbre
retenu/enfant partagé et entrée nominale (100 000 nœuds accessibles, profondeur
22). Le journal est `probe-before/correctness-with-nominal.log`.

| Médiane de cinq appels directs | Compilé conservé | Manuscrit conservé |
| --- | ---: | ---: |
| Allocations | 2 483 949 | 100 000 |
| Octets cumulés | 79 486 448 | 3 200 000 |
| Durée | 21,786375 ms | 8,510875 ms |
| Profondeur, chacun des cinq appels | 22 | 22 |

Les dix échantillons sont dans `probe-before/*-allocations.jsonl`, avec les
médianes dans `allocations-summary.json`. Les durées suivent le protocole de
sonde décrit plus haut, distinct du harnais publié.

La fixture PureScript `OwnedTrees` a ensuite passé la compilation PureScript,
la génération Go, la compilation Go et l'exécution (`Done`) avec le bundle
existant, sans rebuild du backend. Le nouveau snapshot a été écrit après
réussite. Les fonctions top-level `rebalance`, `descend`, `blacken`, `put`,
`build` et `mixed` sont présentes dans le Go conservé. Sources, Go, CoreFn,
binaire, logs et empreintes sont archivés dans `owned-trees-before/`.
Un premier essai s'était arrêté avant compilation sur l'ouverture du cache
SQLite Spago ; l'essai autorisé avec accès à ce cache a réussi.

## Validation et campagne après optimisation

`OwnedTrees` passe également après reconstruction du bundle. Le snapshot ajoute
sept workers consommants et leurs sept wrappers ; seules `Get_Main_main` et
`Call_Main_three` changent parmi les fonctions préexistantes, pour sélectionner
des appels depuis `Tip`. Les autres fonctions publiques sont identiques à la
référence. Les preuves correspondantes sont dans `owned-trees-after/`.
Le contrôle final `./bin/test OwnedTrees --keep-workspace`, sans mise à jour de
snapshot, passe aussi : snapshot existant identique, compilation et exécution
`Done`. Son journal et son SHA sont dans `owned-trees-final-check/`.

Les builds `after-pure` et `after-fficc` ont réussi avec le bundle existant,
sans `-c`, et leurs manifestes identifient sources, profil et binaires. Le Go
RBTree ajoute cinq workers consommants ; son `act` appelle la construction
consommante depuis `nil`. Les fonctions publiques restent persistantes.

Les six tests Go de `probe-after-owned/correctness.log` passent : les quatre
contrôles précédents, neuf séquences transmises exclusivement au worker (dont
les quatre rotations, des doublons et les trois ordres), et la comparaison de
forme/couleurs entre construction consommante et persistante sur sept tailles
de 0 à 4 096. L'entrée nominale confirme 100 000 nœuds et profondeur 22.

Après libération du CPU par les autres agents, `measure-comparison.py` a exécuté
trois processus par variante, dans un ordre tournant, sans compilation. Les
126 résultats des neuf suites sont validés. Les conditions restent celles de
la référence (Go 1.27.0 darwin/arm64, GOGC=800, sans PGO).

| Médiane par ligne, puis somme | Avant compilé | Après compilé | Manuscrit |
| --- | ---: | ---: | ---: |
| Suite complète | 27,51207 ms | 13,21992 ms | 11,27350 ms |
| RBTree | 24,19912 ms | 9,94929 ms | 8,92221 ms |

Dans cette campagne, le total baisse de 51,95 % et RBTree de 58,89 %. Le RBTree
compilé reste 11,51 % plus lent que le manuscrit. Ce sont les comparaisons
contrôlées de cette campagne ; elles ne remplacent pas les nombres historiques
24,02/11,09 ms du README et ne démontrent pas une supériorité générale.

Les sondes utilisent trois processus par variante et cinq appels par processus,
soit 15 échantillons, avec le GC préalable décrit plus haut :

| Médiane de la sonde RBTree | Avant compilé | Après compilé | Manuscrit |
| --- | ---: | ---: | ---: |
| Allocations | 2 483 949 | 100 000 | 100 000 |
| Octets cumulés | 79 486 352 | 3 200 000 | 3 200 000 |
| Durée | 23,154625 ms | 10,068500 ms | 8,860250 ms |
| Profondeur de chaque échantillon | 22 | 22 | 22 |

Après et manuscrit varient de 100 000 à 100 001 allocations et de 3 200 000 à
3 200 016 octets ; la médiane est une cellule de 32 octets par clé. Les
allocations Runtime mesurées ne sont pas un compteur instrumenté des seuls
constructeurs. `main.go` et `persistence_test.go` sont identiques entre les deux
sondes ; les adaptateurs et noms des workers figurent dans `entrypoints.json`.
Les 18 sorties d'exécution, leurs stderr, les SHA et toutes les valeurs sont
conservés dans `comparison/`, notamment `comparison/results.json`.

Le backend et le README d'altbak n'ont pas été modifiés par cette sous-tâche.
