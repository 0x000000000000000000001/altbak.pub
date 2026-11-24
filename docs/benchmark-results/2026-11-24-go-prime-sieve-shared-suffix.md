# Prime Sieve : partage du suffixe immuable des filtres

Mesures du 24 novembre 2026. Troisième optimisation de Prime Sieve, après
la construction directe des listes et le prédicat natif. La baseline inclut
aussi les calculs directs de List Processing et Array Processing.
Les [données brutes](2026-11-24-go-prime-sieve-shared-suffix.json) conservent les
manifestes, empreintes, sorties, profils, validations et le premier candidat
remplacé après les contrôles de performance des filtres.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 34,788656 µs | 15,438436 µs |
| Allocations du calcul | 5 375 | 1 604 |
| Octets alloués par calcul | 128 984 | 38 480 |
| Résultat vérifié | 21 536 | 21 536 |

Accélération **2,253×**, soit **55,62 % de temps en moins**.
Les allocations diminuent de **70,16 %**. Les temps incluent les adaptateurs
Effect, l'entrée opaque et le sink du harness. Les allocations sont mesurées
séparément sur le calcul généré `range(2,500) → sieve → sumList`.

Échantillons avant : 34,760418 ; 35,044270 ; 34,788656 µs.
Échantillons après : 15,438436 ; 15,257364 ; 15,442993 µs.
La campagne isolée donne **34,548176 → 15,091187 µs**,
soit **2,289×**.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **73,303547 µs / 67,25× C**.
Face au C historique de **1,09 µs**, le nouveau ratio est **≈ 14,16× C**,
contre **≈ 31,92× C** avant cette campagne. Face au JS historique
de **63,388023 µs**, il est **≈ 0,244× JS**. Ces références ne sont pas
remesurées ; le README reste inchangé. Prime Sieve reste la priorité absolue
et demeure au-dessus du seuil de 2× C. Le crible strict à listes est conservé ;
le C de référence utilise des tests de divisibilité jusqu'à la racine carrée.

## Goulot mesuré

Le filtre initial allouait un nœud pour chaque élément retenu. Le profil
mémoire lui attribue environ **88,6 % des objets alloués**. Le profil CPU
montre une part importante de l'allocation et du parcours, ainsi que des
attentes du runtime sur cet hôte ; les pourcentages CPU ne se confondent donc
pas avec la réduction de temps mesurée sur les binaires.

L'audit de l'entrée 500 constate **95 appels de filtre**, dont **87 sans aucun
rejet**, et **5 182 appels de prédicat**. Les filtres recopient initialement
**4 778 nœuds**. En partageant leur suffixe inchangé, seuls **1 007 nœuds**
doivent être copiés. Les **499 nœuds de la plage**, les **95 nœuds de la liste
résultat** et trois allocations annexes restent présents.

Un premier prototype a confirmé le gain et les résultats sur 605 bornes de
-4 à 600, deux bornes plus grandes et 390 scénarios de filtre. Le prototype
reprend le Go du crible en remplaçant uniquement l'appel au filtre.

Le premier builder partageait déjà le suffixe mais ajoutait du travail sur
chaque rejet. Un contrôle dédié a trouvé environ **0,49 → 0,70 µs** lorsque
tout était rejeté. Avant de retenir le changement, un second prototype a
regroupé les parcours des éléments rejetés et retenus. Il a passé 511 masques
exhaustifs et les 605 oracles, puis a amélioré chacun des six profils testés
face au premier builder. Le compilateur final utilise ce parcours par groupes.
Toutes les mesures du premier candidat restent archivées pour rendre ce choix
vérifiable ; elles ne remplacent pas les chiffres finaux ci-dessus.

## Changement du compilateur

Seul le builder de `Gopurs.ListFilter` change. Sa reconnaissance structurelle
complète, sa validation du layout TAST et ses signatures publiques/natives
restent identiques. L'optimisation s'applique aux filtres reconnus sur des
listes natives d'Int ; elle ne dépend ni de Prime Sieve ni de la limite 500.

Le builder saute un groupe d'éléments rejetés, puis parcourt un groupe retenu.
Si ce groupe atteint la fin, il raccorde directement ce suffixe immuable à
la sortie. Si un rejet suit, il copie le groupe retenu dans de nouveaux nœuds,
consomme ce rejet et poursuit. Le premier élément de chaque groupe est
explicitement consommé avant de changer de phase : les deux sites d'appel
du prédicat ne l'évaluent jamais deux fois sur un même élément.

Chaque prédicat est appelé **une fois par élément, dans l'ordre source**,
y compris pour les callbacks opaques ou avec effets. Copier un groupe ne
rappelle jamais le prédicat. Toutes les écritures de liens concernent des
nœuds fraîchement alloués ; aucun nœud d'entrée n'est modifié. Les arguments
sont capturés une fois et avant tout travail du filtre.

Cette transformation ne suppose aucune propriété exclusive des entrées et
n'utilise pas les annotations d'usage comme preuve d'ownership. La revue a
vérifié que `Rc` n'autorise pas de mutation conditionnelle et que les workers
mutables existants exigent des arbres frais de constructeurs : les résultats
de filtres et leurs aliases locaux ne satisfont pas cette preuve. Le GC Go
gère les pointeurs partagés normalement. Aucun pool, arène, changement du
runtime, du benchmark PureScript ou du harness n'est introduit.

## Contrôle des profils de filtre

Filtre généré seul sur une entrée de 500 nœuds construite avant le chronométrage.
Médianes de trois répétitions `go test`, durées en µs ; allocations par appel.
Ces diagnostics distincts de la suite complète mesurent le compromis général :
un rejet tardif oblige à parcourir puis copier un préfixe presque complet.

| Profil | Avant (µs) | Après (µs) | Allocations |
| --- | ---: | ---: | ---: |
| Tout retenu | 2,987 | 0,405 | 500 → 0 |
| Tout rejeté | 0,503 | 0,403 | 0 → 0 |
| Premier rejeté | 2,998 | 0,412 | 499 → 0 |
| Dernier rejeté | 2,976 | 3,375 | 499 → 499 |
| Rejet au milieu | 2,967 | 1,901 | 499 → 249 |
| Un élément sur deux | 1,649 | 1,718 | 250 → 249 |

Le gain n'est pas uniforme selon le profil : le rejet du dernier élément
coûte **13,41 %** de plus, et l'alternance **4,18 %**
de plus. Dans ces cas, le suffixe partageable est vide ou très court et le
parcours supplémentaire coûte plus que les allocations évitées. Ces surcoûts
du filtre général sont conservés explicitement avec toutes les répétitions.
Le résultat principal reste le crible complet.

## Validation

- 27 tests IR du filtre et 2 tests ciblés de codecs : les deux ABI, les phases,
  les transitions consommées une seule fois, les écritures sur nœuds neufs,
  l'absence de callback pendant la copie et les captures inchangées.
- Hôtes JS et natif, Release et Dev : 143 sorties par profil, avec callbacks
  importés/opaques, factories, captures de la source et partiels réutilisés.
- Par hôte Release : 25 contrôles d'allocations adaptés au préfixe copié,
  511 masques exhaustifs pour les longueurs 0 à 8, et trois positions de panic.
- Vérification du préfixe neuf et du suffixe partagé par identité, des liens
  et valeurs de chaque nœud source, de l'ordre et du nombre des callbacks,
  et du nombre exact d'allocations dans les cas exhaustifs.
- Vrai crible généré contre un oracle indépendant sur 605 bornes, puis cinq
  entrées supplémentaires jusqu'à 3 000 sur le candidat natif.
- Tous les résultats de la suite complète sont vérifiés contre les 14 oracles.
  Compilateur natif reconstruit et installé.

## Contrôle des autres benchmarks

Les **307 modules CoreFn typés sont identiques**. Le Go généré ne change que
dans **`Call_Test_Primes_filter__394169086`**. Le crible, ses appels de prédicat,
sa plage initiale et sa somme finale restent identiques.

| Test | Avant (µs) | Après (µs) | Écart |
| --- | ---: | ---: | ---: |
| List Processing | 0,025797 | 0,025990 | 0,75 % |
| Array Processing | 0,025998 | 0,026261 | 1,01 % |
| Red-Black Tree | 9648,812500 | 9705,625000 | 0,59 % |

Les treize autres benchmarks varient de **-2,15 % à 1,01 %** dans
la campagne principale. Les valeurs complètes sont archivées ; leur Go reste
inchangé. Cette campagne n'explique pas les écarts des campagnes antérieures
et ne les présente pas comme corrigés.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Les générations finales utilisent l'hôte natif, Release, sans PGO, `GOGC=800`
et `GOMAXPROCS` par défaut (14 dans les diagnostics). Baseline gopurs :
`a355e98d05831b95eb9b5c59d7d4b542af575133`, avec les six campagnes précédentes
appliquées. Fork PureScript : `41fb570fa58e3698646ea6e6179571348e4acd12`.
Les empreintes archivées identifient les sources et les exécutables exacts.

Pour réserver l'espace nécessaire aux compilations, d'anciens workspaces ont
été conservés en archives `.tar.gz` et trois objets de cache Go obsolètes de
nos builds du compilateur ont été supprimés. Aucun résultat de benchmark ou
fichier source n'a été supprimé. Les emplacements figurent dans le JSON.

Construire chaque workspace avec la version correspondante du compilateur :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-sharing-20261124-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test Primes --build-only \
  --build-dir /tmp/gopurs-primes-sharing-20261124-after
```

Dans chaque workspace, copier `output/main/main.go` dans
`output/coremain/main.go`, remplacer uniquement `Get_AppX_main()` par
`Get_App_main()`, puis construire `core-benchmark` :

```sh
go -C /tmp/gopurs-primes-sharing-20261124-before/output build -pgo=off \
  -o /tmp/gopurs-primes-sharing-20261124-before/core-benchmark ./coremain
go -C /tmp/gopurs-primes-sharing-20261124-after/output build -pgo=off \
  -o /tmp/gopurs-primes-sharing-20261124-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test Primes --core \
  --before /tmp/gopurs-primes-sharing-20261124-before \
  --after /tmp/gopurs-primes-sharing-20261124-after \
  --output /tmp/prime-sharing-core-comparison-new
```

Omettre `--core` et choisir un autre répertoire de sortie pour la mesure isolée.
Le runner alterne AB/BA/AB : trois processus indépendants par variante,
trois échauffements locaux, calibration, minimum de dix batches, puis médiane
des trois processus. La suite complète ajoute trois échauffements globaux.
Compilations et tests étaient arrêtés pendant les mesures finales.
