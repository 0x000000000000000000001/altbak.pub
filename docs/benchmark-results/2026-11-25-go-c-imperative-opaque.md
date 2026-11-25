# C impératif optimisé et entrées opaques — 25 novembre 2026

La colonne C principale utilise désormais la référence impérative
`tmp/bench_native.c`. La traduction C fonctionnelle reste une expérience
historique distincte. Cette campagne remesure les 14 cas Stable en C, Go et
JavaScript officiel, avec trois processus par langage dans un ordre alterné.
Les sources PureScript et les harnais Go/JS sont inchangés.

Les totaux des médianes sont **8,758890 ms en C**, **10,224658 ms en Go** et
**111,877916 ms en JS**. Les [données et preuves complètes](2026-11-25-go-c-imperative-opaque.json)
conservent les neuf sorties, les sources figées, les manifestes et les audits.

## Les quatre indicateurs

Sur les mêmes 14 cas, avec les valeurs non arrondies :

- Go/C, ratio des totaux : **1,167346×**.
- Go/C, moyenne arithmétique μ : **20 252,989174×**.
- Go/C, médiane Q₂ : **7,382225×**.
- Go/C, moyenne géométrique G : **14,084496×**.

Pour Go/JS : total **0,091391×**, μ **0,268532×**, Q₂ **0,133956×** et
G **0,015747×**. Le README publie ces mesures dans ses tableaux existants,
avec les règles d’arrondi habituelles. Les 23 mesures Incl WIP et les autres
colonnes conservent leurs sources historiques ; elles ne sont pas remesurées ici.

Church et State représentent **99,75 % de μ/C**. Leurs temps C sont proches de
1,1 ns : le compilateur réduit les calculs impératifs à quelques opérations sur
l’entrée runtime. Go prend respectivement 225,37 µs et 90,75 µs. Ces écarts
demeurent donc des pistes d’optimisation Go ; ils ne proviennent pas d’une
réponse C précalculée pour l’entrée du harnais.

Le ratio des totaux par cycle va de 1,148× à 1,169×. Le temps C de RBTree va
de 8,425 à 8,959 ms ; State Go de 82,082 à 94,205 µs. Ces plages observées ne
sont pas des intervalles de confiance. Les moyennes portent sur les ratios
des médianes, et le ratio total sur les sommes de médianes.

## Référence impérative

Les calculs utilisent les paramètres du programme PureScript, dont **20
répétitions de 60 opérations** pour State. L’ancien C passait 60 au noyau et
fixait les répétitions à 20 ; cette inversion est corrigée.

RBTree emploie les quatre rotations d’Okasaki en place, avec une arène contenant
les nouveaux nœuds seulement. L’arène est allouée et libérée dans chaque action.
L’ancienne arène de 640 Mo préallouée hors mesure est supprimée. L’AST utilise
un stockage local pour les petites profondeurs et une arène interne au-delà ;
le crible emploie un tableau impératif. Les boucles peuvent être fusionnées et
les calculs simplifiés, conformément au [contrat impératif](../../tmp/imperative_reference_contract.md).
Le code C définit ses opérations Int32 sans débordement signé indéfini.

## Ce qui prouve l’opacité

Le binaire C mesuré conserve, dans chaque itération du batch, l’appel à
`opaque_input`, puis l’appel de l’action et le stockage atomique du résultat.
`opaque_input` reste non inliné et effectue une écriture puis une lecture
volatile. L’assemblage des noyaux utilise la valeur retournée : State calcule
`entrée × 60`, Lazy `entrée × 1000`. Il ne retourne pas les constantes 1200 ou
1000000 indépendamment de l’entrée.

Les 68 variantes des 14 actions concordent en C, Go et JS. Les variantes Go/JS
modifient uniquement le résultat de la frontière opaque dans des copies
temporaires, sans toucher les benchmarks publiés. Chaque action traverse cette
frontière une fois. Les preuves du binaire Go de production conservent l’appel
à `opaqueValue` et l’utilisation de son retour ; le code TurboFan inspecté
conserve `Atomics.store` et `Atomics.load`, même avec un appel source à 1000.

RowToList ignore explicitement son entier opaque dans le source PureScript.
Son résultat reste donc 5 pour toutes les variantes : il compte les champs
d’un record fixe. Le C dérive ce compte des descripteurs de cette structure,
plutôt que de fournir un oracle littéral. L’opacité d’une entrée n’impose pas
qu’un calcul indépendant de cette entrée change de résultat.

L’audit des autres routes distingue les niveaux de preuve :

- Erlang : 68 variantes concordantes et appel distant à `opaque`, puis
  exécution de la closure, conservés dans le BEAM inspecté.
- C++ : fournisseur canonique `bin/cpp/bench_ffi.cc` avec `noinline` et barrière
  assembleur ; sonde compilée en unités séparées avec les flags publiés. Le
  hash du harnais correspond à l’archive historique. L’ancienne copie identité
  sous `run/bak/cpp/ffi` n’est pas le fournisseur du driver.
- Rust, Java, F#, Haskell, OCaml et Koka : barrières explicites repérées dans
  les sources, sans nouvelle vérification de leurs exécutables ici.
- PHP et Scheme : aucune violation démontrée, mais le code natif final du JIT
  PHP et de la compilation globale Chez n’est pas certifié par cet audit.

La règle d’entrée opaque vaut pour toutes les colonnes. Ces contrôles ne
certifient pas rétroactivement toutes les anciennes campagnes.

## Validation et reproduction

Les 12 tests du collecteur et les 4 tests de la référence passent, dont
ASan/UBSan. Le C vérifie 14 oracles, 56 variantes supplémentaires, 257 tailles
d’arbres, les rotations, doublons et invariants rouge-noir. La campagne valide
126 résultats ainsi que les lots calibrés, sources et exécutables figés.
Une relecture indépendante confirme les quatre indicateurs et les médianes.

Clang 17, C11, `-O3`, sans LTO ; Go 1.27.0 Release, GOGC=800, sans PGO ;
PureScript officiel 0.15.16 et Node 24.8.0. Trois échauffements globaux et
locaux, calibration vers 10 ms par batch, minimum de dix batches puis médiane
de trois processus. Les builds et tests sont terminés avant les mesures ;
la machine n’est pas isolée des autres activités utilisateur.

```sh
python3 -B bin/benchmark/compare-core.py --c-kind imperative \
  --go-dir var/benchmark/imperative-opaque-20261125/builds/go \
  --js-dir run/bak/js/output-runs/pure-4290hk6q \
  --c-binary var/benchmark/imperative-opaque-20261125/builds/c/benchmark-c \
  --output var/benchmark/imperative-opaque-20261125/replay
```

Les 75,66 ms de la [référence fonctionnelle](2026-11-25-go-c-js-functional.md)
et les 9,11 ms de l’[ancienne référence impérative](2026-11-24-c-reference.md)
conservent leurs contrats et limites propres. Le passage de 75,66 à 8,76 ms
change l’implémentation C ; il ne constitue pas une accélération isolée de clang.
