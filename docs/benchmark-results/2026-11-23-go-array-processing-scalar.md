# Array Processing : somme filtrée en temps constant

Mesures du 23 novembre 2026. Seconde optimisation d'Array Processing, après
la [fusion des tableaux](2026-11-23-go-array-processing.md). La baseline inclut
les cinq campagnes précédentes, dont le calcul direct de List Processing.
Les [données brutes](2026-11-23-go-array-processing-scalar.json) conservent les
manifestes, empreintes, sorties, diagnostics, preuves Go et journaux de validation.

## Résultat dans la suite complète

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Médiane de trois processus | 0,356887 µs | 0,025920 µs |
| Allocations du calcul | 0 | 0 |
| Octets alloués par calcul | 0 | 0 |
| Résultat vérifié | 202 950 | 202 950 |

Accélération **13,769×**, soit **92,74 % de temps en moins**.
Ces temps incluent les adaptateurs Effect, l'entrée opaque et le sink du
harness. Les allocations sont mesurées séparément sur `sumEvens(900)`.

Échantillons avant : 0,355927 ; 0,356887 ; 0,357412 µs.
Échantillons après : 0,025830 ; 0,026114 ; 0,025920 µs.
La campagne isolée confirme **0,355235 → 0,025912 µs**, soit **13,709×**.

La baseline officielle du [README](../../README.md#go), précisée dans
[la campagne publiée](2026-11-23-go.json), est **2,673269 µs / 53,47× C**.
Face au C historique publié de **0,05 µs**, le nouveau ratio est
**≈ 0,518× C**, contre **≈ 7,138× C** dans la mesure locale avant changement.
Face au JS historique exact de **7,350688 µs**, il est **≈ 0,00353× JS**.
Les références C et JS ne sont pas remesurées ; le README reste inchangé.
Array Processing sort donc des priorités locales au seuil de 2×.

## Goulot et preuve expérimentale

La première optimisation avait supprimé les tableaux intermédiaires. Le Go
parcourait encore les 900 entiers, appelait le prédicat modulo et additionnait
chaque valeur retenue. Il s'agissait désormais d'un coût de parcours, sans
allocation.

Un prototype a remplacé cette boucle par la somme exacte d'une progression
arithmétique. Il a passé 10 000 cas aléatoires déterministes dans les deux
directions, six cas extrêmes et une comparaison avec le vrai ArrayOps sur les
3 008 bornes de -8 à 2 999. Un oracle BigInt vérifie les grandes sommes.
Le prototype conserve le parcours pour les bornes non positives ou les seeds
hors Int32 ; les grandes plages positives sont calculées directement.

Avec entrée chargée par opération atomique et prototype non inlinable, le
diagnostic a mesuré **332 à 422 ns** pour le Go initial et **1,98 à 2,00 ns**
pour la formule. Ces mesures ne sont pas celles du harness. Le résultat
applicatif retenu reste **13,769×** dans la suite complète ; les diagnostics
`go test` servent aussi à contrôler les allocations.

## Changement du compilateur

`Gopurs.CongruenceSum` partage avec List Processing la reconnaissance du
prédicat et le corps du calcul exact. `ArrayRangeFold` n'élimine les appels du
prédicat que pour une lambda directe à un paramètre dont le corps complet est
`valeur mod diviseur == résidu`. Le diviseur doit être une constante positive
Int32 et le résidu compris entre zéro inclus et le diviseur exclu.
Les appels, factories, captures utilisées dans le calcul, effets et opérations
supplémentaires ne sont pas admis dans ce prédicat.

Tous les arguments restent capturés une seule fois dans leur ordre source,
avant les gardes. Dans le chemin où les bornes sont déjà prouvées Int32,
le calcul direct exige en plus deux bornes strictement positives et un seed
Int32. Une fonction Go immédiatement appelée normalise ses bornes privées :
`Array.range` est inclusive dans les deux directions. La somme Int32 est
commutative et le prédicat reconnu est pur ; cette normalisation préserve donc
le résultat.

```text
bas = min(début, fin)
haut = max(début, fin)
premier = bas + (résidu - bas % diviseur + diviseur) % diviseur
si premier > haut : retourner seed
nombre = (haut - premier) / diviseur + 1
dernier = premier + (nombre - 1) * diviseur
retourner Int32(seed + nombre * (premier + dernier) / 2)
```

Les opérations intermédiaires sont émises en Go `int64`, y compris les
constantes. Le premier candidat peut dépasser MaxInt32 : il est testé avant
de calculer le nombre de termes. Lorsqu'un terme existe,
`nombre*(premier+dernier) ≤ 2*MaxInt32² < MaxInt64`. La division par deux
est exacte et l'ajout du seed reste sûr. La conversion finale reproduit les
additions modulaires Int32. `runtime.IntMul`, qui reproduit les arrondis de
la multiplication JavaScript, n'est pas utilisé.

Les bornes non positives et les seeds hors Int32 gardent la boucle fusionnée.
Les bornes hors Int32 gardent la construction de tableaux existante, sur les
arguments déjà capturés. Les autres callbacks conservent leurs appels et leur
ordre. Le benchmark PureScript, le runtime et le harness ne changent pas.
La transformation ne dépend ni du nom ArrayOps ni de la limite 900.

## Validation

- **33 tests ciblés** : 24 ArrayRangeFold, 4 FilteredSumLoops, 3 AffineLoops,
  2 Int32Loops. Compilation JS sans erreur ni avertissement.
- **715 cas exécutés sur le Go émis**, dont 120 grandes sommes vérifiées par
  BigInt : deux sens, seeds Int64, bornes négatives/larges, ordre et unicité
  des captures, reconnaissance stricte et replis.
- Fixture compilée par les hôtes JS et natif, en Release et Dev :
  **120 sorties par profil**, avec valeurs, tableaux partagés et traces d'effets.
- Par hôte Release : **51 contrôles sans allocation**, **42 grandes plages**
  avec oracle BigInt et **240 contrôles des replis** contre le calcul matérialisé.
- Prédicats constants `mod 2 == 0` et `mod 5 == 3`, singletons, sélection vide,
  bornes proches de MinInt32/MaxInt32, prédicats dynamiques et opaques,
  captures observables et fold à deux phases.
- Assertions sur le Go réel confirmant l'admission de trois fonctions et
  le refus des callbacks dynamiques/opaques.
- Vrai ArrayOps : 3 008 petites bornes, puis sept grandes bornes jusqu'à
  MaxInt32 sur le candidat. Pas de matérialisation de ces grandes plages
  dans le JavaScript de référence.
- Les 14 résultats de la suite complète passent leurs oracles dans les six
  processus. Compilateur natif reconstruit et installé.

## Contrôle des autres benchmarks

Les **307 modules CoreFn typés sont identiques**. Le Go généré ne change que
dans **`Call_Test_ArrayOps_sumEvens`**. Le package contenant List Processing
reste identique octet par octet malgré le partage du calcul dans le compilateur.

| Test | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| List Processing | 0,026047 µs | 0,025936 µs | −0,43 % |
| Prime Sieve | 34,734863 µs | 34,596762 µs | −0,40 % |
| Red-Black Tree | 9689,666500 µs | 9618,958000 µs | −0,73 % |

Les treize autres benchmarks varient entre **−0,90 % et +0,16 %** dans cette
campagne, sans changement de leur Go ; les valeurs complètes sont archivées.
Cette mesure n'explique pas les écarts relevés lors de campagnes antérieures
et ne les présente pas comme corrigés.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Les deux générations finales utilisent l'hôte natif, Release, sans PGO,
`GOGC=800`, et `GOMAXPROCS` par défaut (14 dans les diagnostics).
La baseline gopurs est `a355e98d05831b95eb9b5c59d7d4b542af575133` avec les
cinq campagnes précédentes appliquées ; les empreintes archivées identifient
cet état et le candidat. Fork PureScript :
`41fb570fa58e3698646ea6e6179571348e4acd12`.

La dernière étape Go de la reconstruction native a échoué deux fois avec un
journal vide. Après libération d'espace disque, sa reprise avec trace détaillée
a réussi ; la cause initiale n'est pas établie. Le binaire a été installé par
remplacement atomique avec vérification de son empreinte, puis validé par les
tests et générations natifs ci-dessus. Les logs sont conservés dans le JSON.
Six anciens workspaces des premières campagnes ont été archivés sans perte
dans `/tmp/gopurs-{listops,primes,arrayops}-20261123-{before,after}.tar.gz` ;
les workspaces avant/après de cette campagne restent disponibles.

Construire chaque workspace avec la version correspondante du compilateur :

```sh
GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ArrayOps --build-only \
  --build-dir /tmp/gopurs-arrayops-scalar-20261123-before

GOPURS_JS=0 GOPURS_DAEMON=0 GOPURS_RESIDENT=0 GOGC=800 \
  bin/go/run --test ArrayOps --build-only \
  --build-dir /tmp/gopurs-arrayops-scalar-20261123-after
```

Dans chaque workspace, copier `output/main/main.go` dans
`output/coremain/main.go` et remplacer uniquement `Get_AppX_main()` par
`Get_App_main()`, puis construire `core-benchmark` :

```sh
go -C /tmp/gopurs-arrayops-scalar-20261123-before/output build -pgo=off \
  -o /tmp/gopurs-arrayops-scalar-20261123-before/core-benchmark ./coremain
go -C /tmp/gopurs-arrayops-scalar-20261123-after/output build -pgo=off \
  -o /tmp/gopurs-arrayops-scalar-20261123-after/core-benchmark ./coremain

python3 bin/benchmark/list-processing.py --test ArrayOps --core \
  --before /tmp/gopurs-arrayops-scalar-20261123-before \
  --after /tmp/gopurs-arrayops-scalar-20261123-after \
  --output /tmp/array-scalar-core-comparison-new
```

Omettre `--core` et choisir un autre répertoire de sortie pour la mesure
isolée. Le runner alterne AB/BA/AB : trois processus indépendants par variante,
trois échauffements locaux, calibration, minimum de dix batches, puis médiane
des trois processus. La suite complète ajoute trois échauffements globaux.
Compilations et tests étaient arrêtés pendant les mesures finales.
