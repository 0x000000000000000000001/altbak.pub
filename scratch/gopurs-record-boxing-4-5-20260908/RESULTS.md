# Boxing compact des records natifs — bilan de 4.5

Le boxing compact intégré réduit nettement le coût de **StateMonad**, sans pénalité de mémoire résidente observée. Le noyau passe d'environ **124 à 82 µs**, avec **30 % d'octets alloués et 55 % d'allocations en moins**. Le temps total d'altbak reste comparable : ce changement bénéficie au motif de StateMonad, trop petit dans ce benchmark pour accélérer visiblement l'ensemble.

La comparaison porte sur le boxing natif compact complet des étapes **4.3 et 4.4** : désactivé dans le témoin, activé dans la version actuelle. Elle ne mesure pas le seul incrément de 4.4 par rapport au cas à deux champs déjà intégré en 4.3. Dans StateMonad, les neuf sites modifiés deviennent tous des `RecordDict2`.

## Temps et allocations

Chaque série comprend dix paires séquentielles, avec ordre avant/après alterné. Les colonnes avant/après donnent les médianes de chaque version ; les pourcentages donnent la médiane des variations calculées paire par paire. Ces deux calculs peuvent différer, notamment pour l'enveloppe `act`. Les quartiles décrivent la dispersion, sans constituer un intervalle de confiance.

| Mesure | Avant | Après | Variation appariée médiane |
|---|---:|---:|---:|
| Noyau StateMonad, temps/op | 124,084 µs | 82,392 µs | −34,1 % |
| Noyau, octets alloués/op | 578 560 | 404 800 | −30,0 % |
| Noyau, allocations/op | 13 260 | 6 020 | −54,6 % |
| StateMonad avec son `act`, temps/op | 141,950 µs | 86,389 µs | −35,7 % |
| StateMonad avec son `act`, octets/op | 578 729 | 404 957 | −30,0 % |
| StateMonad avec son `act`, allocations/op | 13 268 | 6 028 | −54,6 % |
| StateMonad dans App, minimum de dix mesures | 124 µs | 89,5 µs | −27,0 % |
| Total affiché par App | 27,495 ms | 27,540 ms | −0,6 % |

Les dix paires sont favorables pour StateMonad dans chacune des trois séries. Sur le noyau, les variations vont de **−39,0 à −31,8 %**, avec quartiles **−37,1 et −33,1 %** : le gain dépasse nettement les petites variations des contrôles.

**Aucun gain global n'est démontré.** Les totaux médians sont pratiquement identiques ; les variations appariées vont de −5,8 à +4,6 %. Le total d'App additionne les minima par test et ne mesure pas la durée du processus. StateMonad économise environ 35 µs dans 27 500 µs, soit seulement 0,13 % de ce total ; RBTree en occupe toujours l'essentiel.

## Mémoire résidente

La sonde exécute exactement **10 000 appels du noyau**, après six échauffements, dans un processus neuf sous `/usr/bin/time -l`. Le RSS maximal médian passe de **44,30 à 43,67 Mo**, pour une variation appariée médiane de −1,0 %. La conclusion défendable est une mémoire résidente comparable, autour de 44 Mo, **sans pénalité observée**.

Une valeur atypique avant optimisation, **93,44 Mo à la paire 6**, est conservée dans les données ; les neuf autres valeurs avant sont proches de 44 Mo. Elle ne permet pas d'annoncer une division par deux de la RAM. Le faible déplacement des médianes ne justifie pas non plus de revendiquer précisément 1 % de RAM gagnée. Les 30 % d'octets alloués en moins mesurent le volume cumulé d'allocations, pas une baisse équivalente de la mémoire résidente. Cette sonde ne mesure pas le pic de l'App complète.

## Ce que confirme le profil

Les profils sont collectés séparément des chronométrages, avec échantillonnage exhaustif, sur vingt appels. Ils confirment **72 400 boxings**, soit **3 620 par appel**, avant comme après. Pour chaque record à deux champs, le conteneur et les deux slices génériques coûtaient 128 octets et trois allocations ; `RecordDict2` coûte 80 octets et une allocation sur cette cible arm64.

| Coût par appel du noyau | Avant | Après |
|---|---:|---:|
| Records : octets / allocations | 463 360 / 10 860 | 289 600 / 3 620 |
| Closures `Apply` et `chainModifications` | 115 200 / 2 400 | 115 200 / 2 400 |
| **Total : octets / allocations** | **578 560 / 13 260** | **404 800 / 6 020** |

L'économie exacte, **173 760 octets et 7 240 allocations/op**, concorde avec les compteurs du benchmark et le delta `MemStats`. Les boxings subsistent : 1 200 dans `modify`, 1 200 à chacune des deux reconversions successives et 20 dans le cas de base. Les reconversions coûtent encore 192 000 octets et 2 400 allocations/op après optimisation ; ce constat ne démontre pas leur éliminabilité et aucun travail de suppression n'a été engagé.

Les exports pprof utilisent `focus=Test_StateMonad` : un filtre limité au seul appel externe `runManyTimes` exclurait les piles profondes dont cette frame est tronquée. `trim_path` et `source_path` permettent de relier les binaires compilés avec `-trimpath` aux sources figées. Les profils incluent aussi les allocations du dispositif de collecte ; les comptes ci-dessus correspondent à la partie chaude filtrée. Les profils bruts sont conservés, et seuls leurs rapports texte ont été régénérés pour corriger ce filtrage et la localisation des sources.

## Les 14 contrôles et la portée du résultat

Les quatorze `act` et l'App complète donnent les résultats attendus dans les deux variantes. Pour les treize autres `act`, les comptes médians d'allocations sont inchangés et les variations temporelles appariées médianes restent entre **−2,6 et +0,33 %**, avec dispersion. Aucun ralentissement soutenu n'apparaît dans cette session ; ces petits écarts ne sont pas revendiqués comme des gains. RBTree conserve notamment environ 2,48 millions d'allocations/op. Le résultat de Fib observé à l'étape 3.5 concerne une autre optimisation et n'est pas réinterprété ici.

Le [README historique d'altbak](/Users/0x1/Documents/htdocs/altbak.pub/README.md) indique environ **114 µs pour StateMonad et 24,66 ms au total**. Il donne un repère d'échelle ; les mesures causales sont les paires de cette session, dont le témoin donne déjà 124 µs et 27,495 ms dans App.

Le protocole conserve les mêmes **299 CoreFn typés**, métadonnées `.purmeta` initiales, FFI et autres optimisations. La comparaison AST des **386 fichiers Go** confirme uniquement **269 substitutions de helpers dans 116 fichiers** ; runtime et FFI sont identiques. Le Go de production régénéré correspond à la variante optimisée. Apple M4 Pro, Go 1.27.0, `GOGC=800`, `GOMAXPROCS=14`, `GOMEMLIMIT=off`, PGO désactivé ; binaires compilés avant les séries, exécutions séquentielles.

Cette étape ne modifie pas les sources du compilateur. Aucun test de la suite `passing` n'a été lancé ; les validations portent sur le Go régénéré et les benchmarks concernés. Les résultats sont propres à ce workload, cette cible et ces réglages du GC.

Les **340 observations principales**, les profils et les scripts restent disponibles : [statistiques détaillées](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-record-boxing-4-5-20260908/summary.md), [données brutes](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-record-boxing-4-5-20260908/measures.jsonl), [delta des profils](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-record-boxing-4-5-20260908/profiles/summary.json) et [protocole de reproduction](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-record-boxing-4-5-20260908/README.md). Le pilote court est conservé séparément.
