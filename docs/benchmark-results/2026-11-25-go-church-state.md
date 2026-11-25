# Church et State : optimisation Go — 25 novembre 2026

Les deux optimisations de gopurs réduisent **Church de 230,346 µs à 34,923 ns**
et **State de 93,373 µs à 18,570 ns** dans une campagne avant/après alternée.
Les benchmarks, leur harnais et la référence C impérative sont inchangés.
Les entrées restent opaques à chaque invocation.

Les [données et preuves](2026-11-25-go-church-state.json) conservent les
échantillons, sorties brutes, manifestes, changements du compilateur, tests,
extraits Go et contrôles d’opacité.

## Les quatre indicateurs Go/C

Les mêmes quatorze médianes C impératives servent de dénominateurs aux deux
versions Go. La référence vient de la
[campagne impérative avec entrées opaques](2026-11-25-go-c-imperative-opaque.md).
Ni le C ni JavaScript ne sont remesurés ici.

- Ratio des totaux : **1,196025× → 1,154000×**, soit −3,51 %.
- Moyenne arithmétique μ : **20 740,274098× → 55,639467×**, soit −99,73 %.
- Médiane Q₂ : **7,403000× → 7,320738×**, soit −1,11 %.
- Moyenne géométrique G : **14,338936× → 4,127068×**, soit −71,22 %.

Le total Go passe de **10,475853 ms à 10,107760 ms**, pour **8,758890 ms en C**.
Les calculs utilisent les temps archivés avant arrondi : ΣGo/ΣC pour le ratio
total, puis la moyenne, la médiane et la moyenne géométrique des quatorze ratios.
Les 23 cas Incl WIP ne participent pas à ces quatre indicateurs.

La précédente publication donnait respectivement **1,167346×**, **20 252,989174×**,
**7,382225×** et **14,084496×**. Elle est conservée dans son archive. Les gains
attribués aux changements du compilateur reposent sur la nouvelle paire
avant/après, qui réexécute aussi l’ancien binaire.

## Résultats des deux cas

- **Church** : accélération de **6 595,83×** ; Go/C passe de **206 220,41×** à
  **31,27×**. Les trois processus après donnent 34,854–34,943 ns.
- **State** : accélération de **5 028,18×** ; Go/C passe de **83 411,16×** à
  **16,59×**. Les trois processus après donnent 18,550–19,133 ns.

Ces plages sont les valeurs observées, pas des intervalles de confiance.
Les ratios ne sont pas encore à 1×. Lazy représente désormais **86,81 % de μ/C**,
avec un ratio de **676,24×**. RBTree représente l’essentiel du temps total ;
les gains considérables sur les deux petites actions réduisent donc beaucoup
plus μ et G que le ratio des totaux.

## Transformations du compilateur

`CountedFunctionSums`, appelé depuis `FunctionFusion`, reconnaît les compositions
d’itérateurs entiers dont le callback est une translation prouvée `x + k`.
Il calcule les sommes et produits correspondants depuis les comptes runtime,
avec garde du domaine Int32 et retour au calcul original pour les domaines
non admis. Les multiplications reproduisent exactement les additions répétées
modulo 2³², y compris aux débordements ; elles ne supposent pas que la
multiplication `Int` habituelle possède cette propriété à toute amplitude.
Les callbacks inconnus, appels étrangers, effets et transformations non
linéaires conservent leur chemin d’origine.

`ClosedScalar` évalue les calculs entiers fermés avec récursion et branches
connues, y compris les projections de records. Les primitives entières donnent
le type d’un opérande dont l’annotation est implicite. Les lectures globales
strictes nécessaires sont conservées, avec des bindings frais. Le nombre
d’étapes et la profondeur de pile sont bornés séparément ; une preuve incomplète
conserve le calcul d’origine.

Dans State, la chaîne pure de profondeur constante 60 devient l’incrément 60.
La boucle conserve le compteur runtime et la lecture de `Data.Unit.unit` ;
elle ne reconstruit plus les closures et records de la chaîne. Les règles
n’utilisent aucun nom de benchmark ni aucune valeur attendue.

## Validation

- Reconstruction du compilateur natif réussie.
- **66 tests ciblés** validés : 20 pour l’évaluation fermée, 43 pour les
  itérateurs et leurs consommateurs, et 3 pour les récurrences affines.
  Une ancienne attente de niveau local a été corrigée après l’introduction
  des niveaux frais ; les logs intermédiaires restent identifiés séparément.
- La fixture de fonctions récursives produit les mêmes résultats avec les
  compilateurs JS et natif, en Release et Dev.
- **78 variantes d’entrée** concordent avant/après et avec le C inchangé.
  Les sondes temporaires substituent seulement le retour de la frontière
  opaque et vérifient exactement un passage par cette frontière par action.
- L’assemblage du binaire de production garde l’appel à `opaqueValue` via
  le worker FFI et utilise le retour dans Church et State.
- Les **306 résultats** des douze processus et leurs lots calibrés sont
  validés. Les empreintes des benchmarks, du harnais, des références et des
  exécutables sont contrôlées avant et après les mesures.

Le contrôle AppAll couvre les 37 cas. Un seul dépasse +5 % dans la campagne :
`st/ref-write-10000`, de 6,647 à 7,518 µs (+13,10 %). Ses échantillons se
chevauchent : 6,567–7,565 µs avant, 6,473–7,583 µs après. La comparaison des
sources générées ne trouve de différences que dans Church et State ; le code
de ce cas, de ses dépendances et du runtime est identique.

Une contre-mesure isolée, utilisant les mêmes getters compilés et le même
harnais, donne **6,848 → 6,766 µs** (−1,19 %) sur trois nouveaux processus
par variante. La hausse de 13,10 % n’y est pas reproduite ; aucune régression
du code généré n’est démontrée. Ce diagnostic reste séparé et ne remplace
aucun échantillon ni agrégat de la campagne principale.

## Protocole et reproduction

Go 1.27.0, darwin/arm64, Release, GOGC=800, sans PGO, GOMAXPROCS par défaut.
Chaque paire utilise trois processus par variante, en ordre AB/BA/AB.
Le harnais conserve ses échauffements, sa calibration et le minimum de dix
batches ; la valeur publiée est la médiane des trois minima de processus.
Tous les builds et tests étaient terminés avant la campagne. La machine
n’était pas isolée des autres activités utilisateur.

L’entrée `App` à 14 cas fournit les chiffres Stable publiés. Une paire séparée
`AppAll` à 37 cas contrôle les régressions et fournit les 23 chiffres Incl WIP.
Le README actualise uniquement la colonne Go Release dans ses tableaux
existants ; les autres colonnes et les références gardent leurs mesures.

Le script de collecte est inclus dans l’archive. Son rejeu intégral vérifie
aussi le catalogue antérieur à cette publication. Pour réexécuter directement
les deux paires de binaires figés dans de nouveaux répertoires :

```sh
env -u GOMAXPROCS -u PPROF GOGC=800 python3 -B bin/benchmark/measure.py \
  --plan var/benchmark/go-church-state-20261125/campaign/core-plan.json \
  --output var/benchmark/go-church-state-20261125/replay-core

env -u GOMAXPROCS -u PPROF GOGC=800 python3 -B bin/benchmark/js-ratios.py \
  --before var/benchmark/go-small-actions-20261125/after-all \
  --after var/benchmark/go-church-state-20261125/after-all \
  --output var/benchmark/go-church-state-20261125/replay-all
```

Les quatre indicateurs se recalculent avec les temps C conservés dans
`campaign.frozen_reference.c_times_us` de l’archive.
