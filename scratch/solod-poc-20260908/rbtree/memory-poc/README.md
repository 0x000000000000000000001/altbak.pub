# Arènes Go : vitesse et budget mémoire

Test du 8 septembre 2026, uniquement en Go, sur le vrai RBTree généré par
gopurs. **Le compromis peut être favorable, mais dépend du volume temporaire
et de la réutilisation de la réserve.** Le backend et les sources PureScript
ne sont pas modifiés.

## Résultat principal

Confirmation indépendante après exploration des paramètres : sept mesures
de temps et trois mesures de pic RSS par variante et taille. Le temps est
par arbre complet, construction puis profondeur. `GOGC=800` est le réglage
actuel d'altbak. Les arènes allouent des blocs typés de 1 Mio à la demande.

| Charge | Variante | Temps médian | Pic RSS médian |
| --- | --- | ---: | ---: |
| 10000 insertions | Go actuel, GC800 | 1,776 ms | 38,97 Mio |
| 10000 insertions | Arène progressive, GC800 | **1,005 ms** | **10,45 Mio** |
| 10000 insertions | Pool plafonné à 8 Mio, GC200 | 1,025 ms | 11,19 Mio |
| 100000 insertions | Go actuel, GC800 | 22,634 ms | 43,86 Mio |
| 100000 insertions | Arène progressive, GC800 | **13,729 ms** | **81,41 Mio** |
| 100000 insertions | Pool plafonné à 8 Mio, GC200 | 22,923 ms | 42,02 Mio |
| 100000 insertions | Pool plafonné à 16 Mio, GC100 | 22,466 ms | 45,39 Mio |
| 100000 insertions | Pool plafonné à 24 Mio, GC100 | 21,420 ms | 62,02 Mio |
| 100000 insertions | Pool plafonné à 32 Mio, GC100 | 20,465 ms | 78,55 Mio |
| 100000 insertions | Go sans arène, GC1600 | 23,023 ms | 76,72 Mio |

À 10k, l'arène progressive donne **1,77× la vitesse et 73 % de RSS en moins**.
À 100k, elle donne **1,65× la vitesse mais 86 % de RSS en plus**. La correction
de la grosse réserve du POC précédent ramène son pic d'environ 329 à 81 Mio.
À budget proche des 44–45 Mio du Go actuel, aucun gain substantiel n'est
établi sur 100k : moins de 1 % pour le pool 16 Mio, contre environ 1 % de
ralentissement pour le pool 8 Mio/GC200. Ces petits écarts ne constituent pas
une amélioration pratique convaincante.

Le contrôle Go/GC1600 consomme environ autant que l'arène complète sans
accélérer : le gain de l'arène ne se résume donc pas à desserrer le GC.

Le pool 8 Mio/GC200 est un compromis intéressant dans ce test : il garde le
gain sur les petits arbres et reste proche du Go actuel sur les gros, en
temps et en mémoire. Cela demande encore validation sur d'autres algorithmes
et dans un processus comportant d'autres allocations.

## Premier appel et réutilisation

Les chiffres principaux amortissent la réserve sur **131 arbres à 10k** et
**9 arbres à 100k**, par processus. Sa croissance initiale et tous les
nettoyages sont compris dans le temps ; elle est ensuite réutilisée.

Un test distinct construit **un seul arbre par processus**, sans réutilisation
de l'arène, sept fois par variante. Le démarrage de Go et du programme est
inclus et compte beaucoup dans le petit cas :

| Charge isolée | Go actuel GC800 | Arène progressive GC800 | Pool 8 Mio/GC200 |
| --- | ---: | ---: | ---: |
| 10000 | 4,995 ms | 4,969 ms | 7,322 ms |
| 100000 | 27,461 ms | 20,743 ms | 27,892 ms |

À froid, **pas de gain clair à 10k** pour l'arène complète et le pool 8 Mio
est plus lent. À 100k, l'arène complète reste environ **1,32× plus rapide**.
L'intérêt est donc plus net dans un processus qui réutilise sa réserve.

## Mémoire retenue et stabilité

La variante dite complète est plafonnée à 96 Mio, mais alloue seulement les
blocs nécessaires : **76 Mio à 100k**, couvrant les 2 483 948 constructions
temporaires. Le nœud est conservé au format original de 32 octets. Les
plafonds plus petits utilisent les mêmes blocs, puis les allocations Go
ordinaires une fois la limite atteinte.

Après 1, 10 et 100 arbres de 100k, et nettoyage puis GC forcé, le tas encore
occupé reste stable :

| Variante | Tas retenu après GC | Pic RSS sur 100 arbres |
| --- | ---: | ---: |
| Go actuel GC800 | 0,31–0,34 Mio | 47,28 Mio |
| Pool 8 Mio/GC200 | 8,32–8,34 Mio | 43,94 Mio |
| Arène progressive GC800 | 76,32 Mio | 82,53 Mio |

**Aucune accumulation proportionnelle au nombre d'arbres n'a été observée.**
La réserve est toutefois volontairement conservée à sa taille maximale
atteinte, y compris si les calculs suivants sont plus petits. Ce prototype
ne rend pas automatiquement ces blocs au GC après chaque expression.

Le plafond d'arène ne borne pas le RSS total. L'exploration a notamment
trouvé des pics d'environ 113, 184 et 295 Mio pour des pools de 8, 16 et
32 Mio avec GC800, sans gain utile. Un pool retenu augmente la mémoire
vivante et peut relever le budget du GC. Les variantes avec
`GOMEMLIMIT=64MiB` ont aussi été testées : cette limite Go est souple,
elle ne constitue pas une limite stricte de RSS, et n'a pas donné de gain
convaincant ici. `GOGC` affecte l'ensemble du processus.

## Implémentation et vérifications

- Le noyau Go est copié à l'identique depuis
  `run/bak/go/output/purescript/Test_RBTree.go`. Seuls les 86 constructeurs
  de nœuds sont redirigés dans la variante pool : [allocation.diff](allocation.diff).
  Aucun changement des rotations, couleurs, branches ou algorithmes.
- [pool.go.txt](pool.go.txt) utilise des tableaux typés fixes derrière des
  pointeurs Go. Agrandir la liste des blocs ne déplace aucun nœud. Aucun
  `unsafe`, aucune dépendance Solod, aucune désactivation du GC.
- Tous les champs d'un nœud sont écrits lors de l'allocation. Les slots
  utilisés sont effacés à la réinitialisation pour éviter de retenir des
  références mortes. Ce nettoyage est inclus jusqu'au dernier arbre du batch.
- Les 78 validations d'arbres des trois phases vérifient ordre BST, couleurs,
  racine noire, absence rouge-rouge, hauteurs noires, nombre d'éléments et
  somme. À 100k : profondeur 22, somme 5 000 050 000, hauteur noire 17,
  allocations 2 483 948 pour toutes les variantes pool.
- La validation intervient après un GC forcé avec l'arbre encore vivant.
  Les séquences 100k → 1k → 0 → 100k passent avec plafonds 0, 1, 8, 24 et
  96 Mio. Les checksums de chaque batch et diagnostic sont également vérifiés.

Ce sont des pools globaux dédiés à un seul worker dans ce POC. Une intégration
gopurs devrait gérer explicitement le contexte de l'arène et prouver qu'aucune
racine ne survit à sa réinitialisation. Ces résultats ne prouvent pas encore
un bénéfice général sur tous les programmes ou toutes les plateformes.

## Données et reproduction

- [results-screen.json](results-screen.json) : 11 variantes, cinq passages,
  trois mesures RSS, diagnostics de GC et d'allocations séparés des timings.
- [results-confirm.json](results-confirm.json) : sept variantes sur deux tailles,
  sept passages, trois RSS, temps bruts et dispersion, paramètres exacts.
- [results-stability.json](results-stability.json) : premier appel, tailles
  variables, tas après GC et RSS sur 1/10/100 arbres.
- [sources.json](sources.json) et [versions.txt](versions.txt) : source et SHA-256,
  Go 1.27.0, révision du dépôt. Les exécutables sont identifiés dans les résultats.

Go est compilé normalement avec `-pgo=off -trimpath` sur macOS ARM64.
Chaque variante fixe explicitement `GOGC` et `GOMEMLIMIT`. Les comparaisons
utilisent le même nombre d'arbres par batch, warmup, ordre mélangé déterministe
et chronométrage externe monotone. Le RSS est le maximum du processus enfant,
lu par un parent Python frais, séparément des timings. Aucun autre benchmark
n'est lancé simultanément par ce POC.

Depuis la racine d'altbak.pub :

```sh
bash scratch/solod-poc-20260908/rbtree/memory-poc/run.sh
python3 scratch/solod-poc-20260908/rbtree/memory-poc/confirm.py
python3 scratch/solod-poc-20260908/rbtree/memory-poc/stability.py
```

Ces commandes reconstruisent puis remplacent les résultats de cette extension
uniquement. Les mesures Solod précédentes sont conservées. Pour vérifier les
arbres sans chronométrage :

```sh
bash scratch/solod-poc-20260908/rbtree/memory-poc/run.sh --verify-only
```
