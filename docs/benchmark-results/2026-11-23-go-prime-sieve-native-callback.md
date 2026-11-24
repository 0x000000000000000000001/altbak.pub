# Prime Sieve : prédicat natif du filtre

Mesures du 23 novembre 2026. Seconde optimisation de Prime Sieve, après la
[construction directe des listes filtrées](2026-11-23-go-prime-sieve.md).
La baseline inclut aussi les optimisations précédentes de List Processing et
Array Processing. Les [données brutes](2026-11-23-go-prime-sieve-native-callback.json)
conservent les sorties, manifestes, empreintes, diagnostics et profils.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 39,562988 µs | 34,773926 µs |
| Allocations du calcul | 5 470 | 5 375 |
| Octets alloués par calcul | 130 504 | 128 984 |
| Résultat vérifié | 21 536 | 21 536 |

Accélération **1,138×**, soit **12,10 % de temps en moins**.
Les temps incluent les adaptateurs Effect, l'entrée opaque et le sink du
harness. Le diagnostic d'allocations appelle directement les fonctions générées
`range(2,500) → sieve → sumList`, hors adaptateurs. Les **95 closures** supprimées
représentent **1,74 % des allocations** ; les nœuds des listes restent alloués.

Échantillons avant : 39,562988 ; 39,714520 ; 38,977703 µs.
Échantillons après : 33,870279 ; 34,773926 ; 34,802979 µs.

La campagne isolée donne **39,334145 → 33,978434 µs**
(13,62 % de temps en moins). La mesure principale reste celle de la
suite complète, avec ses trois échauffements globaux et ses quatorze tests.
Le gain intégré est plus modeste que le prototype ; il est présenté séparément.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **73,303547 µs / 67,25× C**.
Face au C historique publié de **1,09 µs**, le nouveau ratio est
**≈ 31,90× C**. Face au JS historique de **63,388023 µs**,
il est **≈ 0,549× JS**. Les références C et JS n'ont pas été
remesurées ; le README n'est pas republié dans cette tâche.

## Goulot et preuve expérimentale

Le filtre construit déjà sa sortie directement depuis la première optimisation.
Son premier paramètre restait pourtant un `runtime.Value`, avec un appel à
`runtime.Apply(predicate, runtime.Int(value))` pour chaque élément. Le crible
créait en outre une closure boxée capturant le nombre premier courant à
chacun de ses 95 étages pour l'entrée 500.

Un prototype conservant les mêmes listes, le même parcours et le même
algorithme a mesuré **42,271 → 33,475 µs** en appel direct, soit environ
**20,8 % de temps en moins**. La seule différence utile était le prédicat
`func(int64) bool`. Les allocations sont passées de **5 470 à 5 375**, avec
**130 504 → 128 984 octets** par calcul. Une variante incorporant directement
la condition dans la boucle n'apportait pas de gain supplémentaire net.
Les trois versions donnent le résultat attendu sur les 605 limites de -4 à 600.

Le profil mémoire attribue encore environ **87,6 % des objets alloués** aux
nœuds produits par le filtre. Le profil CPU comprend une part importante
d'attentes du runtime sur cet hôte ; la preuve du gain vient du prototype
exécuté et de la comparaison finale des binaires.

## Changement du compilateur

La reconnaissance structurelle et la validation du layout TAST de
`Gopurs.ListFilter` sont partagées entre le choix de signature et l'émission
du worker. Seuls les filtres entièrement reconnus sur une liste native d'Int
reçoivent un premier paramètre `func(int64) bool` dans `ModuleBindings`.
Le builder appelle ensuite directement ce prédicat, dans l'ordre source.

`CallExprs` construit une lambda native lorsque la signature publiée demande
précisément ce type, puis capture tous les arguments dans leur ordre initial.
Le résultat du corps est converti vers `bool` si nécessaire, sans retraduire
ni réévaluer la lambda. Les autres appels conservent leur traduction existante.

L'interface publique boxée reste disponible. Un adaptateur dans `GoConversions`
capture une fois le prédicat boxé et fournit la fonction native attendue.
Son corps utilise l'appel runtime habituel pour les callbacks opaques. Les
applications partielles et les fonctions stockées dans des références restent
utilisables. Les signatures importées et les caches transportent le `TypeFunc`
par les métadonnées existantes, sans nouveau registre ou schéma de cache.

La transformation ne dépend ni du nom du benchmark ni de la limite 500.
Le filtre emprunte l'entrée immuable et ne modifie que ses nouveaux nœuds de
sortie. Les callbacks conservent leur nombre d'invocations, leur ordre et leurs
effets. Le crible à listes strictes, le benchmark PureScript et le harness
restent identiques. La référence C historique emploie des tests de divisibilité
jusqu'à la racine carrée ; le ratio inclut cette différence d'algorithme.

## Validation

- 26 tests IR : reconnaissance/refus, signature publiée, callbacks natifs et
  opaques, conversion du retour, ordre des captures, imports et appels partiels.
- 2 tests ciblés des codecs de modules et des références Go, dont `TypeFunc`.
- Fixture renommée avec un filtre dans un module réellement importé et
  conservé par une directive `never` : 11 entrées, 143 observations par profil.
- Lambdas avec captures, factories du prédicat et de la source, callbacks
  stockés dans `Ref`, fonction importée et application partielle réutilisée,
  listes vides, alias et fallback sur records.
- 25 scénarios d'allocations par hôte Release ; le contrat antérieur est
  conservé : un nœud par élément retenu et un coût constant borné.
- Tests d'intégration JS et natifs, en Release et Dev ; oracle indépendant
  du crible sur 605 limites ; contrôle des allocations du véritable benchmark.
- Tous les résultats de la suite complète sont validés contre les 14 oracles.
  Compilateur natif par défaut reconstruit et installé.

## Contrôle des autres benchmarks

Les **307 modules CoreFn typés** sont identiques avant/après. Le Go généré ne
change que dans le filtre spécialisé de Prime Sieve, son getter public et
son appel depuis `sieve`.

| Test | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| List Processing | 0,424637 µs | 0,424671 µs | 0,01 % |
| Array Processing | 0,359651 µs | 0,362181 µs | 0,70 % |
| Red-Black Tree | 9776,271000 µs | 9694,708500 µs | -0,83 % |

Un écart positif indique un ralentissement. Les treize autres benchmarks
varient de moins de 1,8 % dans cette campagne ; aucun changement de leur Go
n'est observé. Les 14 résultats sont corrects dans les six processus.
Red-Black Tree ne ralentit pas dans cette campagne (−0,83 %). La cause du
ralentissement relevé lors de la campagne Array Processing reste inconnue ;
cette nouvelle mesure ne l'explique pas et n'établit pas sa correction.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1 ; Go 1.27.0 darwin/arm64,
Node 24.8.0. Les deux générations finales utilisent l'hôte **natif**, Release,
PGO désactivé, `GOGC=800`, `GOMAXPROCS` par défaut (14 dans les diagnostics).
La baseline gopurs est le commit `a355e98d05831b95eb9b5c59d7d4b542af575133`
avec les trois optimisations précédentes déjà appliquées. Les empreintes
archivées identifient cet état et le candidat. Le fork PureScript est au commit
`41fb570fa58e3698646ea6e6179571348e4acd12`.

Construire un workspace avec chaque version du compilateur avant les mesures :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-native-callback-20261123-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-native-callback-20261123-after

python3 bin/benchmark/list-processing.py --test Primes \
  --before /tmp/gopurs-primes-native-callback-20261123-before \
  --after /tmp/gopurs-primes-native-callback-20261123-after \
  --output /tmp/prime-native-callback-isolated-comparison-new
```

Pour la suite complète, copier `output/main/main.go` dans
`output/coremain/main.go`, remplacer uniquement `Get_AppX_main()` par
`Get_App_main()`, puis construire `core-benchmark` depuis chaque workspace :

```sh
go -C /tmp/gopurs-primes-native-callback-20261123-before/output build -pgo=off \
  -o /tmp/gopurs-primes-native-callback-20261123-before/core-benchmark ./coremain
go -C /tmp/gopurs-primes-native-callback-20261123-after/output build -pgo=off \
  -o /tmp/gopurs-primes-native-callback-20261123-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test Primes --core \
  --before /tmp/gopurs-primes-native-callback-20261123-before \
  --after /tmp/gopurs-primes-native-callback-20261123-after \
  --output /tmp/prime-native-callback-core-comparison-new
```

Le runner alterne AB/BA/AB avec trois processus indépendants par variante.
Chaque test utilise trois échauffements locaux, une calibration, puis le
minimum de dix batches ; la suite complète ajoute trois échauffements
globaux. Le résultat final est la médiane des trois processus. Les compilations
et les tests de cette campagne sont arrêtés pendant les mesures finales.
