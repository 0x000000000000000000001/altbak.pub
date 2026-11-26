# Les neuf ratios Go/JS passent sous 2×

Campagne du 24 novembre 2026, après les huit optimisations locales précédentes.
Les **neuf tests de la liste passent sous 2× JS**. Le contrôle élargi confirme
ce seuil pour les **37 benchmarks Stable et Incl WIP**, dans chacun des trois
processus après optimisation. Le plus grand ratio médian est désormais
**State Monad : 1,764×** ; son échantillon le plus lent atteint **1,852×**.

Les références JavaScript sont les valeurs exactes de la première colonne JS
du [README officiel](../../README.md#javascript). Le JavaScript n'a pas été
remesuré. Les références viennent des campagnes du 20 septembre et du 14
novembre 2026 ; le README reste inchangé.
Les [données brutes](2026-11-24-go-js-ratios.json) archivent les sorties, échantillons,
manifestes, sources ciblées, empreintes et journaux de validation.

## Avant/après dans la suite complète

Temps en microsecondes ; médianes de trois processus par version. La liste
est triée selon le ratio final, y compris RowToList désormais sous 1×.

| Test | Référence JS | Go avant | Go après | Go/JS après |
| --- | ---: | ---: | ---: | ---: |
| State Monad | 48,230633 | 84,883789 | 85,067711 | 1,764× |
| string/indexOf-hit-10000 | 0,190754 | 0,382634 | 0,309349 | 1,622× |
| record/set-large-1000 | 7,028198 | 9,377502 | 9,357768 | 1,331× |
| record/set-medium-1000 | 4,553304 | 5,639120 | 5,644246 | 1,240× |
| effect/ref-write-10000 | 155,670578 | 188,559234 | 188,067047 | 1,208× |
| AST Evaluation | 0,259782 | 0,304699 | 0,308730 | 1,188× |
| string/concat-1000 | 47,710773 | 55,941406 | 55,818684 | 1,170× |
| array/map-10000 | 121,935867 | 126,953453 | 128,079102 | 1,050× |
| RowToList | 0,041569 | 0,134999 | 0,025624 | 0,616× |

**RowToList : 5,27× plus rapide**, soit
-81,02 % de temps. **indexOf-hit :
1,24× plus rapide**, soit -19,15 %.

- **RowToList** : avant 0,134048 ; 0,137320 ; 0,134999 µs ; après 0,025354 ; 0,025717 ; 0,025624 µs.
- **string/indexOf-hit-10000** : avant 0,382634 ; 0,384535 ; 0,381320 µs ; après 0,302018 ; 0,309349 ; 0,320325 µs.

Les seuils stricts sont respectivement **0,083138 µs** et **0,381508 µs**.
Le Go historique publié était à **0,140972 µs** pour RowToList et
**0,404519 µs** pour indexOf-hit (environ 2,12× JS pour ce dernier).
La première baseline fraîche d'indexOf-hit donnait déjà **0,380571 µs**, soit
**1,99509×** : la marge était presque nulle. Le gain présenté ci-dessus compare
les deux binaires locaux dans la même campagne alternée.

## RowToList : évaluer le calcul entier fermé

Le Go précédent reconstruit une chaîne de dictionnaires et de fermetures pour
un résultat entièrement déterminé par la structure typée. Le prototype isolé
mesure **136,3 → 25,78 ns** et **13 → 3 allocations**, contre 78,52 ns si l'on
se contente de partager le dictionnaire. La réduction du calcul fermé apporte
une marge plus nette sous le seuil.

`Gopurs.ClosedScalar` est une passe générale sur l'IR typé, avant l'émission Go.
Elle interprète uniquement des littéraux entiers, records immuables, captures
lexicales, applications, projections, constructeurs intermédiaires et additions
Int32. Elle émet un littéral uniquement lorsque le résultat est explicitement
annoté Int par le TAST. Les noms du benchmark, de ses classes et de ses
instances n'interviennent pas dans la reconnaissance.

Tous les arguments stricts, champs et initialisations doivent être prouvés purs.
Les corps de fermetures ne sont évalués que s'ils sont appelés ; les globaux
gardent leur portée lexicale. Appels importés/étrangers, effets, récursion,
branches et opérations inconnues conservent l'expression originale.
Un budget de 512 nœuds et une profondeur de 64 bornent chaque tentative.
L'addition conserve le bouclage Int32 sur les hôtes JS et natif.

Dans le vrai Go RowToList, **l'appel `Bench.opaque 10000` est conservé** ;
seul le résultat pur de la chaîne de dictionnaires devient le littéral 5.
Une sonde séparée confirme **13 → 3 allocations par appel** sur cette action
réelle, déjà initialisée puis appelée via `Apply(action, Unit)`. L'opaque est
inclus, l'échauffement est exclu et l'oracle 5 est vérifié avant/après sur
1 000 exécutions. Les commandes et empreintes des sondes sont archivées.

## indexOf-hit : accélérer le comptage du préfixe ASCII

La recherche Go donne un offset en octets ; le port le convertit ensuite en
unités UTF-16. Le comptage du long préfixe ASCII constituait un coût évitable.
Le runtime valide d'abord 128 octets par blocs de 32, puis parcourt les longues
plages ASCII par blocs de 128. Au premier octet non ASCII, il reprend le parcours
fin puis le décodage existant. Chaque chargement est protégé par une borne.
Le port strings et les règles UTF-16/WTF-8 restent inchangés.

Le prototype retenu mesure la recherche isolée à **294,2 → 219,6 ns**
(−25,36 %), à allocation inchangée. Le comptage ASCII de 10 006 octets passe
de **154,3 à 100,3 ns**, sans allocation. La première variante sans garde
ralentissait les courtes plages ASCII ; la garde a éliminé ce ralentissement
dans les profils mesurés. Les cinq profils de comptage retenus vont de
−10,04 % à −39,28 %. Le profil nommé Unicode alterne `a😀é` : il ne mesure
pas une chaîne entièrement non ASCII et ne prouve pas un gain du décodage.

## Validation

- Build JS commun sans erreur ni avertissement ; compilateur natif reconstruit
  et installé, à partir de 543 modules TAST.
- **68 tests ciblés réussis** : 13 tests IR de la nouvelle passe et 55 tests
  de dictionnaires, fusion précédente, codecs et caches.
- Fixture indépendante vérifiée contre le JS source et un oracle explicite,
  sur **les deux hôtes du compilateur, en Release et Dev** : plusieurs tailles
  de dictionnaires renommés, captures, curry, limites Int32, effets séquencés,
  fermeture non appelée et callback opaque.
- Tests runtime et parité du port strings réussis : alignements et frontières
  des blocs, préfixes 0..512, octets hauts, UTF-8/WTF-8 malformé, tous les octets
  simples et paires d'octets, 10 000 chaînes pseudo-aléatoires.
- **222 oracles et 222 batches calibrés vérifiés** dans la comparaison finale
  des 37 benchmarks (six processus). Les 111 échantillons après restent sous
  2× leur référence JS.
- **333 modules TAST identiques** entre les deux versions. Le détail des
  changements de fonctions Go est archivé dans les données brutes : uniquement
  RowToList, le scanner ASCII et trois simplifications triviales
  `Int(Int(constante).IntVal)` en `Int(constante)`. Aucune fonction ajoutée
  ou supprimée ; le diff complet des trois fichiers confirme ce périmètre.

Les 35 autres benchmarks varient entre **-4,46 %**
(Fibonacci) et **+1,32 %** (AST Evaluation).
Le runtime ASCII est partagé par les opérations de chaînes ; les autres
écarts ne sont pas attribués à une cause établie. Les résultats complets des
37 tests, y compris leurs trois échantillons avant/après, sont archivés.
Les observations de ralentissement des campagnes antérieures restent ouvertes.

## Conditions et reproduction

Apple M4 Pro, 48 Gio RAM, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Les deux binaires AppAll finaux sont produits par le compilateur natif en
Release, sans PGO, avec `GOGC=800` et `GOMAXPROCS` par défaut. Benchmarks,
harness et port strings inchangés. Les binaires et sources Go gelés sont conservés.

```sh
go -C /tmp/gopurs-js-ratios-20261124-before/output build -pgo=off \
  -o /tmp/gopurs-js-ratios-20261124-before/benchmark ./main
go -C /tmp/gopurs-js-ratios-20261124-after/output build -pgo=off \
  -o /tmp/gopurs-js-ratios-20261124-after/benchmark ./main

python3 bin/benchmark/js-ratios.py \
  --before /tmp/gopurs-js-ratios-20261124-before \
  --after /tmp/gopurs-js-ratios-20261124-after \
  --output /tmp/gopurs-js-ratios-comparison-new
```

Le runner alterne AB/BA/AB, soit trois processus indépendants par version.
Le harness applique trois échauffements globaux, trois échauffements locaux,
une calibration puis dix batches ; chaque processus retient le minimum, et
la campagne retient la médiane des trois processus. Compilations et tests
sont arrêtés pendant les mesures. Le runner vérifie les oracles, les sources
et les empreintes des exécutables avant et après la campagne.
