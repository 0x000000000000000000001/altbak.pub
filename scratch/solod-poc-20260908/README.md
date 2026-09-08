# POC altbak.pub : Go généré → Solod → C

Test local du 8 septembre 2026. **Gain sur deux noyaux numériques sur trois,
mais la sortie Go complète ne passe pas dans Solod.** Aucun changement au
backend gopurs, aux sources PureScript ou au benchmark principal.

Extension demandée ensuite : [test RBTree avec 100000 insertions](rbtree/README.md),
avec comparaison de la gestion mémoire Go et d'une arène commune Go/C.

## Résultats mesurés

Temps médians par appel, sept batches par variante ; plus petit = mieux.
Les entrées alternent entre les deux valeurs indiquées afin d'éviter un
calcul invariant dans la boucle du driver.

| Calcul extrait | Entrées | Go original | Go adapté | Solod + Clang -O3 | Accélération vs Go original |
| --- | --- | ---: | ---: | ---: | ---: |
| Fibonacci | 10 / 11 | 0,288 µs | 0,287 µs | 0,123 µs | **2,33×** |
| Ackermann | (3,4) / (3,5) | 50,38 µs | 50,00 µs | 55,95 µs | **0,90×**, soit 11 % plus lent |
| Récursion terminale, somme des `n % 3` | 100000 / 100001 | 41,15 µs | 40,90 µs | 30,46 µs | **1,35×** |

Go original et adapté diffèrent de moins de 1 % dans ces mesures. Les gains
observés ne proviennent donc pas d'une réécriture manuelle des algorithmes.
Les timings bruts, tailles de batches, ordre d'exécution, checksums, dispersion
et empreintes des exécutables sont conservés dans [results.json](results.json).

Environnement : macOS 26.1, ARM64, Go 1.27.0, Solod **v0.3.0**, Apple Clang 17.
Go utilise ses optimisations normales, `-pgo=off -trimpath` ; C utilise `-O3
-std=gnu11 -fwrapv`, sans LTO et sans fast-math. Voir [versions.txt](versions.txt).

## Ce qui a été essayé

1. Transpilation du programme complet depuis `run/bak/go/output/Main` : échec
   dans `gopurs_runtime/runtime.go:853`, fonction `ExtractVariant` avec trois
   valeurs de retour : `multi-return must have exactly 2 values`.
   Diagnostic exact : [full-translate.log](full-translate.log).
2. Extraction à l'identique des fonctions `Call_Test_Fib_fib`,
   `Call_Test_Ackermann_ackermann`, `Call_Test_TCO_deepTailRec`. Leur traduction
   directe échoue sur `labeled continue is not supported` :
   [direct-kernels.log](direct-kernels.log).
3. Copie adaptée : suppression des labels des boucles uniques et des
   `if false` associés ; remplacement des `continue label` par `continue` ;
   suppression des IIFE `panic("unreachable")` situées après `continue`.
   Les autres instructions, les signatures et les algorithmes sont préservés.
   Diff intégral : [adaptations.diff](adaptations.diff).
4. Compilation Go des deux copies, transpilation de la copie adaptée avec
   Solod, puis compilation du C produit. Le fichier [driver.c](driver.c)
   contient uniquement le lanceur et l'addition des résultats. Les algorithmes
   se trouvent dans `generated/kernels.c`, produit automatiquement.

Les sources d'origine, leurs lignes et SHA-256 sont dans
[sources.json](sources.json). Révision du dépôt :
`f447b5c73c1170e8ee586a12ed6ae27f9f621905` ; le Go généré est ignoré par Git,
donc ses empreintes et les copies extraites identifient précisément l'entrée.

## Vérification et limites

- 38 cas vérifiés par exécutable contre des oracles indépendants : **152
  vérifications réussies**, y compris un binaire AddressSanitizer + UBSan.
  Chaque appel mesuré vérifie également le checksum attendu.
  [verification.json](verification.json) conserve la vérification après
  reconstruction. Les trois binaires chronométrés ont conservé leurs SHA-256 ;
  le binaire de diagnostic sanitizer reconstruit possède une nouvelle empreinte,
  enregistrée avec cette vérification, et ne participe jamais aux mesures.
- Calibration à environ 150 ms côté Go, warmup puis sept batches, ordre
  mélangé avec graine fixe. Chronométrage externe monotone, démarrage du
  processus et driver inclus. Les batches C Fibonacci durent environ 64 ms.
- Le C du driver et celui des noyaux sont compilés en unités séparées sans
  LTO. Les entrées viennent de la ligne de commande et varient à chaque appel.
- Ce sont trois fonctions numériques sans allocation applicative. Ces premières
  mesures ne couvrent ni le GC, ni les listes, ni RBTree, ni la suite complète.
  Fibonacci est très court, donc le coût du driver compte davantage.
- Binaire du POC : Go 2 413 698 octets, C 34 584 octets. Ce rapport concerne
  uniquement ces petits exécutables et leurs bibliothèques de lancement.
  Il ne prédit pas la taille du programme complet.
- Une seule session sur cette machine : charge concurrente, fréquence CPU et
  température peuvent affecter les mesures. Les résultats ne sont pas une
  garantie de performances pour d'autres plateformes.
  Les entrées alternées diffèrent aussi des entrées uniques de la suite initiale :
  les temps absolus ne sont pas à comparer directement au README principal.

La conversion globale demanderait un vrai travail sur le runtime : closures
dans les wrappers, `sync.Once`, réflexion, allocations et représentation des
fonctions via `unsafe.Pointer` (`runtime.go:268`, `454`, `478`). La
[spécification officielle de Solod](https://github.com/solod-dev/solod/blob/v0.3.0/doc/spec.md)
décrit un sous-ensemble de Go, avec notamment l'absence de fonctions anonymes.
Corriger seulement le premier message d'erreur ne rendrait pas le runtime compatible.

## Rejouer

Depuis la racine d'altbak.pub :

```sh
bash scratch/solod-poc-20260908/run.sh
```

Le script installe Solod v0.3.0 dans le dossier du POC si nécessaire, extrait
à nouveau le Go actuellement généré, conserve les échecs de traduction,
compile les variantes, vérifie les calculs et remplace `results.json`.
La première installation nécessite le réseau. Les caches Go sont temporaires.
Pour reconstruire et vérifier sans refaire les mesures :

```sh
bash scratch/solod-poc-20260908/run.sh --verify-only
```

Pour refaire uniquement les mesures sur les exécutables présents :

```sh
UBSAN_OPTIONS=halt_on_error=1 python3 scratch/solod-poc-20260908/bench.py
```
