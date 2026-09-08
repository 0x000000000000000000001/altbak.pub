# RBTree : Go généré → Solod → C

Extension du POC du 8 septembre 2026 au vrai noyau RBTree généré par gopurs.
Le calcul conserve les insertions descendantes d'Okasaki et le parcours de
profondeur du benchmark original. Les sources PureScript, le backend et la
suite principale ne sont pas modifiés.

Suite de l'expérience : [arènes Go progressives et plafonnées, vitesse/RAM](memory-poc/README.md).
Elle corrige la grosse réserve de cette première version et compare aussi
les appels isolés et la mémoire après 100 arbres.

## Résultats

**À 100000 insertions, Solod/C avec arène est 2,06× plus rapide que le Go
original réglé comme altbak (`GOGC=800`). Face au Go utilisant la même arène,
le facteur est 1,21×.** L'essentiel du gain vient donc de la gestion mémoire.

Temps médians par arbre complet (construction puis profondeur), sept batches
par variante. Les batches contiennent 159 arbres à 10k et 11 arbres à 100k.

| Variante | 10000 insertions | 100000 insertions | Pic RSS à 100k |
| --- | ---: | ---: | ---: |
| Go original, `GOGC=100` | 2,078 ms | 36,721 ms | 14,45 Mio |
| Go original, `GOGC=800` (altbak) | 1,897 ms | 23,419 ms | 45,77 Mio |
| Go, syntaxe adaptée, `GOGC=800` | 1,915 ms | 23,526 ms | 42,30 Mio |
| Go, arène, `GOGC=800` | 0,973 ms | 13,749 ms | 329,33 Mio |
| Solod → C `-O3`, arène | **0,778 ms** | **11,392 ms** | **77,12 Mio** |

À 10k, le C est 2,44× plus rapide que Go original `GOGC=800`, et 1,25× que Go
avec arène. À 100k, la simple adaptation syntaxique change le temps de moins
de 1 %. Le code source des rotations n'a pas été réécrit à la main.

Le compromis mémoire est réel : l'arène C conserve tous les nœuds temporaires,
alors que le GC de Go peut les récupérer pendant les insertions. Le C utilise
ici plus de mémoire résidente que Go original `GOGC=800`. Le gros tableau
Go d'arène et son runtime ont encore un autre profil mémoire. Les allocations
et pages touchées diffèrent des réservations virtuelles ; le RSS concerne
l'ensemble du processus. Un relevé mémoire par variante ne mesure pas la
variabilité de ce pic.

**59 vérifications d'arbres réussies**, dont quatre sous AddressSanitizer/UBSan,
sans diagnostic. Données brutes, ordre des passages, checksums, dispersion,
RSS et SHA-256 des exécutables : [results.json](results.json).

## Adaptations

Le type du nœud et les sept fonctions `Call_Test_RBTree_*` ont été extraits
depuis `run/bak/go/output/purescript/Test_RBTree.go`, ligne 170 à la fin.
Ce noyau ne dépend pas du runtime gopurs. Son empreinte et les transformations
sont enregistrées dans [sources.json](sources.json).

La traduction directe est refusée par Solod : `stack-allocated value escapes
function frame`. C'est un obstacle réel : les nœuds restent accessibles après
le retour de la fonction qui les crée. Voir [original-translate.log](original-translate.log).

Les changements de syntaxe sont limités aux labels/continue et aux fonctions
anonymes contenant uniquement `panic` : [syntax.diff](syntax.diff).
Les 86 sites de construction `&Constructor_Test_RBTree_T{...}` sont ensuite
redirigés vers `AllocNode`, qui écrit un nœud complet dans une arène bornée :
[allocation.diff](allocation.diff). Le corps des algorithmes, les rotations,
les couleurs, les branches et le partage des sous-arbres sont conservés.

La même source Go de cette arène est compilée par Go et transpilée par Solod.
Le lanceur fournit un tableau sur le tas via Go `make` ou C `calloc`, une fois
par processus. L'arène est réutilisée entre arbres, seulement après consommation
de la profondeur du précédent. Aucun nœud intermédiaire n'est libéré pendant
la construction : cela préserve le partage structurel. L'arène n'est ni un GC
ni un mécanisme général de gestion des arbres persistants.

La capacité est une borne conservatrice `(n+1)*(6*bits(n)+4)` : pour 100000,
339 203 392 octets réservés (nœud de 32 octets). La construction utilise
exactement **2 483 948 nœuds**, soit **79 486 336 octets** de données écrites,
pour un arbre final de 100000 nœuds. Réservation, données écrites et mémoire
résidente du processus sont trois mesures différentes.

## Comparaison

Les variantes sont Go original (`GOGC=100` puis `800`), Go avec seulement les
adaptations syntaxiques (`800`), Go avec arène (`800`), C avec la même arène.
`GOGC=800` correspond au réglage de `bin/go/run` dans altbak.pub.

Go : optimisations normales, `-pgo=off -trimpath`. C : Solod v0.3.0 puis Clang
`-O3 -std=gnu11 -fwrapv`, unités de compilation séparées et sans LTO. Les sources
du C algorithmique sont entièrement produites par Solod ; [driver.c](driver.c)
ne contient que le lancement, la réserve mémoire et l'addition des résultats.
Versions exactes dans [versions.txt](versions.txt).

Le temps inclut la construction complète et le calcul de profondeur, mais
pas le validateur d'invariants. Mesure externe monotone de batches comprenant
le démarrage du processus, l'initialisation de la réserve et la sortie.
Calibration sur Go original `GOGC=800` à environ 250 ms et au moins trois
arbres, warmup puis sept passages avec ordre mélangé déterministe.
Même taille et même nombre d'arbres par batch pour toutes les variantes.

Les mesures de mémoire sont séparées des timings : un processus Python frais
lance un seul exécutable puis lit son `ru_maxrss` via `getrusage(RUSAGE_CHILDREN)`.
Sur macOS cette valeur est en octets et exclut le parent Python. Le maximum
RSS est une mesure par variante, pas une médiane de consommation. Cet accès
évite le `sysctl` refusé à `/usr/bin/time -l` dans l'environnement du POC.

## Validation

[validate.go](validate.go) vérifie ordre BST strict, couleurs, racine noire,
absence de lien rouge-rouge, égalité des hauteurs noires, nombre d'éléments et
somme. La récursion du validateur est bornée pour détecter les cycles.

Pour 100000 insertions descendantes, Go et C donnent **profondeur 22**, somme
**5 000 050 000**, hauteur noire **17** (feuille nil comprise) et 100000 éléments.
Le nombre d'allocations d'arène concorde aussi. Le binaire C de diagnostic est
compilé avec AddressSanitizer et UBSan, ainsi qu'une vérification de l'usage
de la pile après retour. Chaque batch chronométré vérifie son checksum.

Ces mesures concernent ce calcul complet de RBTree, dans une seule session
locale ARM64. Les gains contre Go original combinent compilateur, runtime et
changement de gestion mémoire ; la comparaison des deux arènes permet de
mieux distinguer ce dernier effet. Une application conservant plusieurs
versions d'arbres simultanément demanderait une stratégie de durée de vie
adaptée avant toute utilisation de cette arène.
Même avec une arène identique, Go conserve son runtime, son GC et ses barrières
d'écriture : le gain résiduel n'est pas attribuable au seul optimiseur C.

## Rejouer

Depuis la racine d'altbak.pub :

```sh
bash scratch/solod-poc-20260908/rbtree/run.sh
```

Cela réextrait le Go actuellement généré, reconstruit, vérifie et remplace les
résultats RBTree uniquement. Les trois premiers benchmarks du POC sont conservés.
Solod v0.3.0 est réutilisé ou installé localement si nécessaire.

```sh
# Reconstruction et validation, sans chronométrage
bash scratch/solod-poc-20260908/rbtree/run.sh --verify-only

# Mesures seules à partir des binaires existants (validation préalable incluse)
python3 scratch/solod-poc-20260908/rbtree/bench.py
```
