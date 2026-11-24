# Prime Sieve : objectif inférieur à 2× C atteint

Mesures du 24 novembre 2026. Quatrième campagne Prime Sieve, après le partage
du suffixe immuable des filtres. L'objectif demandé est strictement inférieur
à **2× le C historique de 1,09 µs**, donc **moins de 2,18 µs**.

Le résultat principal est **0,577380 µs**, soit **≈ 0,530× C**. Les trois
processus après optimisation restent sous 0,59 µs. Les
[données brutes](2026-11-24-go-prime-sieve-progression-fusion.json) conservent
les sorties, manifestes, empreintes, preuves d'admission, prototypes et tests.

## Avant/après dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 15,234578 µs | **0,577380 µs** |
| Ratio au C historique | ≈ 13,98× | **≈ 0,530×** |
| Résultat vérifié | 21 536 | 21 536 |
| Allocations de l'action Effect complète | 1 606 | **5** |
| Octets alloués par action | 38 520 | **592** |

Accélération **26,386×**, soit **−96,21 % de temps**. Échantillons avant :
15,319661 ; 15,234578 ; 15,146322 µs. Après :
0,570887 ; 0,577380 ; 0,585762 µs.

La mesure isolée confirme **15,215373 → 0,565222 µs**, soit **26,919×**.
La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **73,303547 µs / 67,25× C**.
Le ratio au JavaScript historique de **63,388023 µs** devient
**≈ 0,00911× JS**. C et JS ne sont pas remesurés ; le README reste inchangé.

Les allocations ci-dessus sont mesurées séparément sur le même appel public
`Get_Test_Primes_act()` mis en cache, puis `Apply(action, Unit)`, avant et après.
Elles incluent l'entrée opaque et les adaptateurs Effect. Ce périmètre diffère
de la sonde de calcul direct des campagnes précédentes (1 604 allocations
avant cette fusion). Le helper seul utilise un tableau de 512 octets et une
allocation à la borne 500 ; la fonction pure de la fixture conserve également
une fermeture du fold pour le repli, soit au plus deux allocations.

## Preuve empirique avant implémentation

Le profil du Go précédent confirme le poids du filtre et des allocations :
environ 73,3 % du CPU cumulé dans le filtre, 48,5 % dans l'allocation et 6,9 %
à plat dans `IntMod`. Ces pourcentages se recouvrent et ne s'additionnent pas.

Quatre prototypes temporaires ont été comparés dans le même binaire,
sans modifier le benchmark PureScript :

| Calcul direct, cinq répétitions | Médiane | Allocations | Octets |
| --- | ---: | ---: | ---: |
| Go précédent | 15,573 µs | 1 604 | 38 480 |
| Tableau compacté à chaque filtre | 3,411 µs | 1 | 4 096 |
| Valeurs testées contre les têtes retenues | 3,929 µs | 1 | 4 096 |
| Marquage dans un tableau de booléens | **0,520 µs** | **1** | **512** |
| Marquage dans un bitset | 1,101 µs | 1 | 64 |

Les oracles vérifient notamment les bornes −4 à 1 000, les bornes basses
2 à 50 avec plusieurs largeurs et seeds, les limites Int32 et les grandes
entrées jusqu'à 500 000 avec débordement de la somme.
La variante booléenne remplace les 5 182 modulos du cas 500 par
499 lectures de marque et 913 écritures de progression, doublons compris.

## Transformation générale du compilateur

`Gopurs.ListProgressionFold` intervient sur l'IR typé avant TCO. Il reconnaît
la composition complète suivante par ses récurrences et le layout TAST,
sans dépendre des noms du module, des fonctions, des constructeurs ou de 500 :

1. Une plage ascendante fraîche, construite par une boucle descendante de
   pas un qui préfixe les valeurs à une liste vide.
2. Une récurrence `Cons(head, self(filter(predicate, tail)))`, avec branche
   vide exacte, filtre stable et inversion reconnus intégralement.
3. Le prédicat primitif exact `not (value mod head == 0)`, sans autre appel,
   capture ni effet.
4. Une consommation immédiate par une somme `IntAdd` à accumulateur.

La représentation des listes intermédiaires est remplacée par un tableau
privé de booléens. Chaque valeur encore active est ajoutée à la somme, puis
ses multiples futurs sont marqués à partir de **2× cette valeur**. La somme
conserve le bouclage Int32 à chaque addition.

Le départ n'est pas le carré de la tête : pour une plage commençant après 2,
des nombres composés peuvent rester dans le résultat original. Par exemple,
`4..10` conserve `[4,5,6,7,9]`, de somme 31. Le marquage depuis `2×head`
préserve cette sémantique sans supposer que les têtes sont premières.

Les gardes exigent `2 ≤ lower ≤ upper ≤ MaxInt32`, au plus **1 048 576**
éléments et un seed Int32 canonique. Les calculs d'indices et de multiples
utilisent des intermédiaires int64. Les autres entrées conservent l'application
originale. Les listes empruntées, callbacks opaques, autres récursions et folds
non additifs ne sont pas fusionnés.

Les valeurs dupliquées dans les gardes sont exclusivement des littéraux ou
des locaux déjà évalués. Les masquages de binders sont traités pour les lambdas,
lets, récursions et binds Effect. Aucune annotation d'usage ne sert de preuve
d'exclusivité. Les noms des helpers sont réservés contre les collisions.

## Validation et limites mesurées

- **109 tests ciblés passent**, dont 22 pour la nouvelle reconnaissance,
  les cas négatifs, les portées, les collisions et le repli conservé.
- Hôtes JS et natif, Release et Dev : **105 observations par profil**,
  vérifiées contre les résultats PureScript/JS et un oracle indépendant.
- Par hôte Release : **15 cas d'allocation et 9 grandes progressions**, jusqu'à
  1 048 576 éléments, avec au plus deux allocations pour la fonction pure.
- **135 comparaisons** de replis sur bornes/seeds larges et vérification
  de l'identité des nœuds empruntés par hôte Release.
- Tous les résultats des **14 benchmarks** sont vérifiés dans les six
  processus de la comparaison principale.

Un premier test d'ordre utilisait des expressions `unsafePerformEffect`
passées comme arguments. Les 105 sorties Go avant et après étaient identiques,
mais 15 traces différaient du JS source. Ce comportement préexistant est archivé.
Le test final utilise des binds Effect explicites pour vérifier la séquence
seed → lower → upper. La factory opaque initiale reste présente et sa
non-admission par la fusion est vérifiée.

La fermeture du fold reste initialisée avant la garde pour le repli. Son coût
constant a été attribué en comparant la fonction complète au helper : une
allocation de plus, indépendamment de la longueur. Il n'est pas présenté
comme supprimé par cette campagne.

## Autres benchmarks

Les **307 modules CoreFn typés sont identiques**. Une comparaison des fonctions
Go par leur AST trouve uniquement **`Get_Test_Primes_act` modifiée** et le helper
`Call_Test_Primes___gopurs_list_progression_fold_0` ajouté. Les fonctions publiques
range, filtre, crible et somme restent identiques pour leurs autres appels.

| Test | Avant (µs) | Après (µs) | Écart |
| --- | ---: | ---: | ---: |
| List Processing | 0,025882 | 0,026239 | +1,38 % |
| Array Processing | 0,026120 | 0,026399 | +1,07 % |
| Red-Black Tree | 9 518,000000 | 9 553,437500 | +0,37 % |

Les treize autres tests varient de **−1,84 % à +1,38 %**. Leur Go reste inchangé.
Ces variations ne sont pas attribuées à une cause établie, et les écarts des
campagnes antérieures ne sont pas présentés comme résolus.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Les deux binaires finaux proviennent du compilateur natif, en Release, sans
PGO, avec `GOGC=800` et `GOMAXPROCS` par défaut. La baseline inclut les sept
campagnes locales précédentes ; les empreintes identifient les sources et
les binaires exacts.

Le bootstrap natif a produit 542 modules TAST, puis échoué à l'édition de liens
faute d'espace disque. Après suppression de trois anciens objets de cache Go
du compilateur, l'étape `go build -trimpath` a été reprise avec succès dans le
même workspace généré. Le binaire a été installé par copie puis renommage
atomique, comme dans le script de bootstrap. Les journaux de l'échec, de la
reprise et les empreintes d'installation sont archivés. Sources, anciens
binaires de benchmark et résultats ont été conservés.

Les workspaces avant/après et leurs Go générés sont conservés. Pour reconstruire
les exécutables de la suite à partir de ces sources gelées :

```sh
go -C /tmp/gopurs-primes-sharing-20261124-after/output build -pgo=off \
  -o /tmp/gopurs-primes-sharing-20261124-after/core-benchmark ./coremain
go -C /tmp/gopurs-primes-progression-20261124-after/output build -pgo=off \
  -o /tmp/gopurs-primes-progression-20261124-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test Primes --core \
  --before /tmp/gopurs-primes-sharing-20261124-after \
  --after /tmp/gopurs-primes-progression-20261124-after \
  --output /tmp/prime-progression-core-comparison-new
```

Le runner alterne AB/BA/AB : trois processus indépendants par variante,
trois échauffements globaux de la suite, puis trois échauffements locaux,
calibration et minimum de dix batches. Le chiffre principal est la médiane
des trois processus. Omettre `--core` et choisir une autre destination pour
la comparaison isolée. Compilations, tests et profils étaient arrêtés pendant
les mesures finales.
