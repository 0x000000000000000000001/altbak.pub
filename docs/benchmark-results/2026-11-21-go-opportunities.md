# Audit des opportunités Go — 21 novembre 2026

45 tests examinés : 14 core, 23 « excluded WIP », 3 diagnostics et 5 extended.
Les temps de référence sont ceux du README au début de cet audit. Les prototypes
restent dans les artefacts de campagne ; aucun changement de production ni de
chiffre README n’est publié par cet audit.

## Priorités et gains établis

| Piste | Avant → prototype | Gain | Périmètre exact |
| --- | ---: | ---: | --- |
| `traverse Array Maybe`, succès | 217,81 → 12,24 µs | 17,8× | Noyau de traversée, mêmes callbacks |
| `traverse`, échec premier | 188,22 → 11,49 µs | 16,4× | Tous les callbacks restent exécutés |
| `traverse`, échec dernier | 188,29 → 12,04 µs | 15,6× | Même résultat et même ordre |
| `CodePoints.length` | 71,83 → 7,40 µs | 9,71× | Action originale, sans tableau temporaire |
| Concaténation de 1 000 chaînes | 176,04 → 82,13 µs | 2,14× | Action originale et checksum, builder |
| List Processing | 19,71 → 11,40 µs | 1,73× | Noyau, mêmes `range`/`filter`, fold natif |
| String Operations extended | 931,47 → 693,31 µs | 1,34× | Action actuelle, remplacement littéral + CP.take |
| CP.take + checksum | 74,38 → 63,47 µs | 1,17× | Action complète, primitives CP |
| Split + checksums | 204,74 → 180,16 µs | 1,14× | Action complète, primitives CP |

Les priorités de cet audit sont **traverse Maybe et CP.length**, puis
**concaténation et continuité des workers spécialisés de listes** : ces pistes
ont un gain mesuré. `CP.length` est une correction locale particulièrement nette :
4 017 allocations et 285 456 octets deviennent zéro. Pour traverse succès,
16 970 allocations/921 826 octets deviennent 1 003/56 632 ; les `Just` réellement
produits par les callbacks restent inclus.

La boucle traverse doit reconnaître le dictionnaire **Maybe standard**. Les tests
vérifient l’ordre, le nombre d’appels, l’entrée intacte, les tailles limites et une
exception tardive après `Nothing`. Un arrêt anticipé des callbacks changerait la
sémantique observée. Le builder de concaténation doit reconnaître l’append canonique
et conserver l’évaluation stricte des éléments. Les tests Unicode comparent les
probes au FFI Go existant, y compris les données malformées ; ils ne constituent
pas une nouvelle certification exhaustive contre un oracle JavaScript.

## Chantiers prometteurs encore sans gain de temps démontré

- **Boucles ST/Effect** : 30 008 à 40 007 allocations et 1,20 à 2,40 MB pour
  10 000 tours. Le worker tailRecM appelle encore le step boxé, son thunk et
  construit les `Step`. Abaisser le corps local dans la boucle est la priorité ;
  les mutex des `Effect.Ref` restent nécessaires. Les annotations d’usage local
  du TAST ne prouvent pas l’absence de partage.
- **Tableaux** : éliminer les conversions `[]Value ↔ []int64`, garder des workers
  scalaires et spécialiser les callbacks connus. Le fold somme alloue deux buffers
  totalisant 327 680 octets ; le tri monotone conserve 30 008 allocations/2,36 MB.
  Ne pas attribuer toutes celles-ci au comparateur avant un profil dédié.
- **State Monad** : 7 223 allocations/385 664 octets dans l’action. Les binds
  conservent `Apply` et des records dynamiques ; un worker saturé avec retour natif
  état/résultat est une piste générale.
- **JSON → TAST** : 20,20 ms au README. Le profil du binaire correspondant confirme
  29 901 260 octets et 308 118 allocations par corpus. En octets échantillonnés,
  l’index temporaire représente environ 36 %, `RecordDict2/3/4` environ 20 %, et
  la construction de table des types environ 14 % en propre. Il faut travailler
  ces représentations et éviter les records intermédiaires ; un chemin texte
  direct existe déjà. Ces proportions d’allocation ne sont pas des proportions
  CPU ni des gains de temps promis.

## Régression découverte dans String Operations extended

Le nouveau protocole officiel donne **802,21 µs**, contre **507,08 µs** au README
(+58,2 %). Une comparaison alternée de trois processus sur les mêmes actions,
avec un harness identique, confirme l’écart entre les binaires : **587,44 µs**
pour l’ancien, **919,80 µs** pour l’actuel. L’allocation passe de 801 924 à
1 399 647 octets/action. Il s’agit donc de plus qu’une fluctuation entre protocoles.

La comparaison du code et le profil différentiel identifient une évolution du
FFI regex : l’ancien binaire utilisait `regexp`, l’actuel `regexp2` avec adaptation
UTF-16 et gestion des captures JavaScript. `newMatch`, `regexRunes`, `Groups` et
`populateOtherGroups` apparaissent parmi les allocations supplémentaires. Cette
évolution de sémantique interdit de rétablir aveuglément l’ancien moteur.

Une piste locale conserve le moteur actuel : quand la chaîne de remplacement ne
contient aucun `$`, retourner directement cette chaîne dans le callback évite la
matérialisation des groupes et le builder du remplacement. Le scan de préfixe
CP.take fournit un second gain indépendant. Trois processus par variante donnent :

| Variante sur le binaire actuel | µs/action | B/action | Allocations/action |
| --- | ---: | ---: | ---: |
| Baseline | 931,47 | 1 399 648 | 37 464 |
| Remplacement littéral | 804,78 | 1 097 684 | 33 464 |
| CP.take | 807,29 | 1 125 667 | 30 464 |
| Les deux | 693,31 | 825 670 | 26 464 |

Le prototype combiné récupère **25,6 % du temps actuel** et **41,0 % des octets**,
sans prétendre retrouver complètement l’ancien temps. Les valeurs de l’action
et les contrôles CP existants passent. Avant intégration, il faudra compléter les
cas regex de remplacement et la validation officielle. Le premier probe String
Operations de cet audit (602,70 → 473,06 µs, CP.take seul) utilisait encore
l’ancien binaire ; il reste archivé, mais le classement utilise désormais le
prototype mesuré sur le binaire actuel.

Le run complet donne aussi File I/O +24,6 %, STArray −38,3 % et Parallelism
−16,2 % face au README. Sans comparaison causale contrôlée, ces écarts ne sont
pas attribués à un changement particulier du compilateur.

## Coûts élevés qui demandent une lecture différente

RBTree et Polymorphism représentent **94,25 % du total core** de référence.
Pour RBTree, le Go impératif témoin est 12,2 % plus rapide (9,714 → 8,531 ms),
avec les mêmes 100 000 nœuds neufs/3,20 MB, alloués dans l’intervalle. Cela montre
une marge modérée ; cela ne prouve pas encore que supprimer quelques stores de
`balance` la récupérera. Le compilateur produit déjà des workers owned/consume.

Polymorphism est déjà une boucle d’entiers spécialisée, sans lookup de dictionnaire
dans son noyau. Lazy Evaluation a déjà un worker strict et seulement trois
allocations fixes. Une amélioration massive demande une réduction générale de
récurrence affine, avec preuve de terminaison et respect d’Int modulo 2^32.
Il ne faut pas annoncer une suppression de boxing ou de thunks déjà absents.

Les références C emploient parfois un travail différent : arènes AST/RBTree
préallouées hors timing, boucles sans listes/tableaux, test des diviseurs pour
Primes, expressions numériques simplifiées pour Church/State/Lazy/Poly. Les
valeurs `<0,01 µs` ne fournissent pas de dénominateur précis.

Aff Operations contient un délai demandé de 10 ms. File I/O comprend 10 000
écritures/lectures. Parallelism effectue dix fib(42) récursifs en parallèle avec
un noyau Go direct. Leurs grands temps absolus ne suffisent pas à diagnostiquer
un défaut de gopurs ; aucun gain compiler significatif n’y est établi ici.

<details>
<summary>Référence technique : revue des 45 tests, protocole et artefacts</summary>

## Revue de chacun des 45 tests

`README` reste la baseline officielle ; `nouveau run` est une vérification du
protocole complet avec le compilateur actuel, sans les prototypes. `—` désigne un
diagnostic repris de sa campagne validée du même jour. Les gains cités dans la
colonne constat proviennent uniquement de comparaisons avant/après homogènes.

### Core (14)

| Test | README µs | Nouveau run µs | Priorité | Constat |
| --- | ---: | ---: | --- | --- |
| [AST Evaluation](../../src/Test/AstTree/AstTree.purs) | 0.35 | 0.311 | Faible | Construction/évaluation de 43 nœuds ; petit coût absolu. FFI Go historique 0,32 µs contre 0,35 µs. Une représentation ADT plus compacte est une piste générale, sans gros gain établi ici. |
| [Fibonacci](../../src/Test/Fib/Fib.purs) | 0.21 | 0.202 | Faible | Récursion directe sur entiers natifs, sans allocation ni dictionnaire dans le noyau. Peu de résidu propre au backend. |
| [List Processing](../../src/Test/ListOps/ListOps.purs) | 17.09 | 15.629 | Gain mesuré | Noyau 19,71 → 11,40 µs (1,73×), 1 805 → 1 357 allocations. Même range/filter et addition Int ; seul le retour au fold générique avec rebox est supprimé. |
| [Tail Call Optimization](../../src/Test/TCO/TCO.purs) | 64.19 | 62.092 | Moyenne | Boucle native déjà générée. IntMod garde les sémantiques euclidiennes générales ; analyse de bornes/signe pour simplifier modulo 3. Référence FFI Go historique 35,50 µs ; potentiel non mesuré. |
| [Deep Record Updates](../../src/Test/Records/Records.purs) | 8.76 | 8.631 | Secondaire | Copies de structs imbriqués et champs intermédiaires non consommés. Scalarisation/projection possible ; coût absolu limité. |
| [Ackermann](../../src/Test/Ackermann/Ackermann.purs) | 16.38 | 16.048 | Faible | Récursion native et appel terminal transformé en boucle ; déjà proche du C et meilleur que la FFI Go historique. |
| [Church Numerals](../../src/Test/Church/Church.purs) | 233.79 | 227.702 | Moyenne | Boucle comptée déjà présente, mais callback Value/Apply conservé. Spécialiser les fonctions connues ; gain non mesuré. |
| [Prime Sieve](../../src/Test/Primes/Primes.purs) | 77.26 | 72.939 | Moyenne | Le filtre construit puis reverse reconstruit des Cons. Fusion ou réemploi sous preuve de consommation ; le C utilise un autre algorithme (diviseurs). |
| [Red-Black Tree](../../src/Test/RBTree/RBTree.purs) | 10037.62 | 10024.375 | Marge modérée mesurée | Généré 9,714 ms contre FFI native 8,531 ms (−12,2 %), avec les mêmes 100 000 nœuds et 3,20 MB. Témoin de marge, pas prototype de passe compilateur. Simplification de balance/stores à démontrer. |
| [Polymorphism](../../src/Test/Polymorphism/Polymorphism.purs) | 2288.66 | 2222.094 | Recherche | Boucle spécialisée int32 sans dictionnaire dans le noyau. Réduction générale de récurrence affine possible sous preuve de terminaison et overflow ; pas une simple suppression de boxing. |
| [State Monad](../../src/Test/StateMonad/StateMonad.purs) | 94.28 | 86.476 | Élevée, non mesurée | Chaque bind conserve Value/Apply et RecordDict2. Worker saturé et retour natif pour le couple état/résultat : mécanisme général à explorer. |
| [Lazy Evaluation](../../src/Test/LazyEvaluation/LazyEvaluation.purs) | 231.11 | 229.410 | Recherche | Worker strict déjà présent ; absence de chaîne de thunks. Réduction affine possible pour les boucles numériques, comme Polymorphism. |
| [Array Processing](../../src/Test/ArrayOps/ArrayOps.purs) | 7.80 | 7.637 | Moyenne | range/filter/fold alternent tableaux natifs et Value, avec callbacks dynamiques. Même chantier de représentation que les WIP arrays ; petit coût absolu ici. |
| [RowToList](../../src/Test/RowToList/RowToList.purs) | 0.15 | 0.141 | Faible absolu | Dictionnaires et closures autour de cinq champs connus. Évaluation partielle/inlining possible, mais seulement 0,15 µs au total. |

### Diagnostics (3)

| Test | README µs | Nouveau run µs | Priorité | Constat |
| --- | ---: | ---: | --- | --- |
| [Array Indexing](../../src/Test/ArrayIndexing/ArrayIndexing.purs) | 4596.08 | — | Faible | 8 388 608 lectures : environ 0,55 ns/lecture à la baseline. Les conversions hors intervalle sont vérifiées. Pas de gros défaut démontré. |
| [JSON to Typed AST](../../src/Test/JsonTypedAst/JsonTypedAst.purs) | 20204.83 | — | Élevée, gain non mesuré | Profil du chemin combiné : 29 901 260 B et 308 118 allocations/corpus. Échantillonnage des octets : index temporaire ~36 %, RecordDict2/3/4 ~20 %, table des types ~14 % en propre. Chemin texte direct déjà présent ; compacter les intermédiaires et produire les records natifs. |
| [JSON Decoding](../../src/Test/JsonDecoding/JsonDecoding.purs) | 2172.48 | — | Déjà corrigé | La régression précédente est corrigée : 2,172 ms, contre 2,241 ms historique. Six paires validées, chemin texte typé rétabli ; conserver la couverture de ce chemin. |

### Excluded WIP (23)

| Test | README µs | Nouveau run µs | Priorité | Constat |
| --- | ---: | ---: | --- | --- |
| [array/map-10000](../../src/Test/ArrayMap10000/ArrayMap10000.purs) | 166.81 | 159.179 | Élevée, non mesurée | 819 Ko/action et 5 allocations dans le probe : résultat boxé, conversions natif/Value et copie pour checksum ; callback Apply par élément. |
| [array/filter-10000](../../src/Test/ArrayFilter10000/ArrayFilter10000.purs) | 107.89 | 102.326 | Moyenne | 456 Ko et 17 allocations : croissance du buffer sans capacité et conversions. Préserver une évaluation du prédicat par élément. |
| [array/foldl-sum-10000](../../src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | 66.88 | 64.586 | Élevée, non mesurée | Deux buffers int64/Value totalisent 327 680 octets. Lire la représentation existante et conserver un accumulateur scalaire natif. |
| [array/sort-descending-10000](../../src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | 620.04 | 591.714 | Élevée, non mesurée | 30 008 allocations et 2,36 MB malgré le chemin monotone. Spécialiser le comparateur connu et ses adaptateurs ; conserver stabilité et input immuable. |
| [array/eq-neq-last-10000](../../src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | 23.32 | 21.896 | Moyenne | La boucle primitive existe. Les conversions autour des entrées allouent encore ; le probe et le protocole officiel ont des paramètres GC différents. |
| [array/append-1000](../../src/Test/ArrayAppend1000/ArrayAppend1000.purs) | 54.39 | 51.550 | Moyenne | 2 015 allocations/310 Ko : passerelle []interface{} et boxing. Garder une sortie fraîche lorsque les deux entrées sont non vides. |
| [foldable/traverse-maybe-1000](../../src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | 235.23 | 211.685 | Gain mesuré | Noyau 217,81 → 12,24 µs (17,8×) ; 16 970 → 1 003 allocations. Boucle dédiée au Maybe standard, tous callbacks conservés. |
| [foldable/traverse-maybe-fail-first-1000](../../src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | 187.23 | 171.397 | Gain mesuré | Noyau 188,22 → 11,49 µs (16,4×). Les 1 000 callbacks restent évalués en ordre, même après Nothing. |
| [foldable/traverse-maybe-fail-last-1000](../../src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | 186.12 | 170.738 | Gain mesuré | Noyau 188,29 → 12,04 µs (15,6×). Résultat, ordre, input intact et exception tardive vérifiés. |
| [string/cu-length-unicode-1000](../../src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | 0.01 | 0.010 | Déjà corrigé | Lecture de longueur UTF-16 en O(1) avec métadonnée immuable. La construction JoinWith avait augmenté de 47 % ; ne pas étendre ce cache sans mesurer lecture et construction. |
| [string/cp-length-unicode-1000](../../src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | 50.88 | 47.395 | Gain mesuré | Action 71,83 → 7,40 µs (9,71×) avec CP.length seul ; ensemble CP 6,73 µs. Compter sans tableau : 4 017 → 0 allocations. Parité avec FFI existant sur Unicode/WTF-8/malformed. |
| [string/indexOf-hit-10000](../../src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | 0.41 | 0.399 | Faible absolu | Recherche native et comptage UTF-16 du préfixe. Ancien gain déjà obtenu (2,56 → 0,41 µs) ; résidu petit. |
| [string/indexOf-miss-10000](../../src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | 2.44 | 2.339 | Faible | Recherche native déjà très favorable face au JS. Aucun gain significatif du probe. |
| [string/cp-take-unicode-1500](../../src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | 59.22 | 56.368 | Gain mesuré modéré | Action avec checksum 74,38 → 63,47 µs (1,17×). Scanner seulement le préfixe utile et améliorer toCodePointArray ; préserver normalisation malformed. |
| [string/split-1000](../../src/Test/StringSplit1000/StringSplit1000.purs) | 163.98 | 154.329 | Gain mesuré modéré | Action avec tous checksums 204,74 → 180,16 µs (1,14×). Le coût ne se limite pas à strings.Split : conversions Unicode des 1 000 morceaux. |
| [string/concat-1000](../../src/Test/StringConcat1000/StringConcat1000.purs) | 125.61 | 118.054 | Gain mesuré | Action avec checksum 176,04 → 82,13 µs (2,14×) via builder pour fold appendString canonique. Les primitives CP seules ne gagnent que 2 %. |
| [record/set-small-1000](../../src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | 0.53 | 0.509 | Faible | Seulement 3 allocations fixes/64 octets pour 1 000 updates. Boucle native proche du plancher. |
| [record/set-medium-1000](../../src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | 6.44 | 6.134 | Secondaire | 3 allocations fixes/64 octets. Copies de struct 8 champs dans la boucle ; scalariser son état privé, sans muter un record source partagé. |
| [record/set-large-1000](../../src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | 9.97 | 9.798 | Secondaire | Même mécanisme pour struct 16 champs. Gain possible mais coût absolu inférieur à 10 µs. |
| [st/ref-write-10000](../../src/Test/STRefWrite10000/STRefWrite10000.purs) | 290.97 | 260.687 | Élevée, non mesurée | 30 008 allocations/1,20 MB par action. Abaisser le corps local tailRecM dans la boucle et supprimer Step/thunks intermédiaires. |
| [st/ref-modify-10000](../../src/Test/STRefModify10000/STRefModify10000.purs) | 292.31 | 263.773 | Élevée, non mesurée | Même 30 008 allocations/1,20 MB. Conserver l’appel et l’ordre du callback de modification. |
| [effect/ref-write-10000](../../src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | 404.81 | 380.529 | Élevée, non mesurée | 40 007 allocations/1,76 MB. Même chantier tailRecM/Effect ; conserver les mutex. |
| [effect/ref-modify-10000](../../src/Test/EffectRefModify10000/EffectRefModify10000.purs) | 539.51 | 492.762 | Élevée, non mesurée | 40 007 allocations/2,40 MB, adaptateur state/value inclus. Worker Effect natif possible ; usage local ne prouve pas absence de partage. |

### Extended (5)

| Test | README µs | Nouveau run µs | Priorité | Constat |
| --- | ---: | ---: | --- | --- |
| [File I/O](../../srx/Test/FileOps/FileOps.purs) | 442703.25 | 551794.667 | Faible pour gopurs | 10 000 écritures/lectures de fichier, principalement système. Go proche du JS ; aucune optimisation compiler prioritaire démontrée. |
| [STArray Operations](../../srx/Test/STArray.purs) | 0.54 | 0.333 | Faible absolu | Création, pushAll de 10 éléments, pop et show : 0,54 µs. Passerelles ST encore présentes, mais très faible coût absolu. |
| [String Operations](../../srx/Test/StringOps.purs) | 507.08 | 802.208 | Régression et gain mesurés | Ancien binaire 587,44 → actuel 919,80 µs dans le même harness : passage regexp standard à regexp2 et adaptateurs UTF-16/captures. Sur actuel, remplacement littéral + CP.take donnent 931,47 → 693,31 µs (1,34×), sans remplacer le moteur. 37 464 → 26 464 allocations. |
| [Aff Operations](../../srx/Test/AffOperations.purs) | 11009.25 | 10925.375 | Plancher du test | Le test demande explicitement un délai de 10 ms. Son temps ~11 ms ne représente pas le débit Aff ; peu de gain mesurable sans changer le travail. |
| [Parallelism](../../srx/Test/Parallelism.purs) | 1182150.08 | 990267.208 | Faible pour gopurs | 10 fib(42) récursifs exécutés en parallèle. Noyau déjà direct int64/IntAdd/IntSub ; changer Fibonacci ou partager ses résultats changerait le travail. Écart Rust restant non attribué. |

## Protocole, validation et limites

- Baseline figée : `var/benchmark/go-opportunities-20261121/README.before.md`.
- Core/WIP et extended recompilés via le driver officiel, puis exécutés sans
  compilation concurrente. Trois processus indépendants par suite, résultats
  authentifiés par `bin/benchmark/measure.py`, `GOGC=800`, PGO désactivé.
- Les trois diagnostics reprennent leurs dernières campagnes validées du jour.
  JSON Decoding dispose déjà de six paires avant/après et reste à 2,172 ms contre
  2,241 ms historiquement. TAST a reçu un nouveau profil des allocations sur son
  binaire figé, avec les 12 fingerprints vérifiés et 20 parcours du corpus.
  Il n’a pas été recompilé avec le dernier correctif Argonaut de JSON Decoding.
- Strings : 18 processus, trois par variante, `GOGC=100`, `GOMAXPROCS=1`,
  `go testing` 100 ms, préparation hors timing et checksums d’origine inclus.
  Les baselines de ce harness sont souvent 19–41 % au-dessus du README : les
  facteurs publiés ici comparent uniquement les variantes du même harness.
- Arrays/traverse : trois répétitions de 100 ms par cas dans le même processus,
  protocole exploratoire. Les 16 actions originales et leurs valeurs sont
  vérifiées ; les variantes traverse comparent seulement le noyau. Ces facteurs
  devront être confirmés sur les actions complètes et trois processus officiels.
- Core probes : trois processus alternés par variante ListOps/RBTree,
  `GOGC=100`, `GOMAXPROCS=1`, cinq échantillons après deux warmups, GC avant chaque
  échantillon. Les 14 autres relevés sont des diagnostics d’allocation dans un
  processus, pas de nouvelles baselines publiables. 12 563 vérifications
  différentielles réussies dans le driver (dont les contrôles numériques préparés).
- Les builds indépendants core et extended ont été faits en parallèle, puis
  toutes les mesures ont été sérialisées sans compilation concurrente pour
  éviter la contention. Les facteurs exploratoires ne sont pas multipliés ensemble et ne
  permettent pas de calculer un nouveau total officiel.

## Artefacts

Le dossier [de campagne](../../var/benchmark/go-opportunities-20261121/) contient
les manifests, hashes, sorties brutes, sources des probes et vérifications :

- [Couverture structurée](../../var/benchmark/go-opportunities-20261121/coverage.json)
- [Mesures officielles réexécutées](../../var/benchmark/go-opportunities-20261121/measurements/results.json)
- [Arrays, records, refs et traverse](../../var/benchmark/go-opportunities-20261121/arrays-probes/summary.json)
- [Strings](../../var/benchmark/go-opportunities-20261121/strings-probes/summary.md)
- [Core et témoins](../../var/benchmark/go-opportunities-20261121/core-probes/results.md)
- [String Operations actuel, variantes et régression](../../var/benchmark/go-opportunities-20261121/extended-current-probe/candidates.json)
- [Profil des allocations TAST](../../var/benchmark/go-opportunities-20261121/tast-probe/allocations-top.txt)
- [Preuves dans le Go généré core](../../var/benchmark/go-opportunities-20261121/core-audit.md)
- [Correction JSON Decoding précédente](2026-11-21-go-json-decoding-restored.md)

Les prochains changements devraient être génériques : décisions fondées sur le
TAST et les opérations canoniques, sans reconnaissance des noms des benchmarks.
Chaque intégration doit ensuite passer les contrôles sémantiques appropriés et
`./bin/go/run --wip-included` avant de publier des chiffres dans le README.

</details>
