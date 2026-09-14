# ListOps : durées de vie, consommation et réutilisation

La libération des propriétaires conservés par les adaptateurs rend les cellules du vrai filtre généré uniques. Leur consommation seule ne réduit pas le nombre d'allocations. La réutilisation des `Cons` retenus supprime ensuite exactement une allocation par cellule admissible. Sur `sumEvens(900)`, cela représente 450 cellules ; sur `sumEvens(9000)`, 4 500. Les mesures coordonnées du runner complet donnent **36 → 29 → 29 → 25 µs** pour les quatre étapes, soit **−30,56 %** sur List Processing. Les profils de partage montrent la limite : la chaîne complète ralentit de **11,4 à 11,7 %** quand une ancienne racine reste vivante. Ce résultat ne justifie donc pas son application systématique à une liste arbitrairement partagée.

## Périmètre et variantes cumulatives

Le fichier source est la copie figée `rust-perceus-evaluation-20260914/build/frozen/Purs_Test_ListOps/src/lib.rs`, SHA-256 `687280582b00e16f9e698a7e14719144ad829819aae34d81e2647fcd6e48ec48`. `ListOps-baseline.rs` lui est identique octet pour octet. Tous les changements restent dans ce dossier scratch ; aucune modification de Purust ou de sa sortie live.

| Module complet | Changement supplémentaire |
| --- | --- |
| `ListOps-baseline.rs` | Copie exacte de la source figée. |
| `ListOps-lifetimes.rs` | Dans les deux adaptateurs du filtre, acquérir les deux arguments `Rc<List>`, puis détruire `_a1` et `_a0` avant l'appel à `_f`. |
| `ListOps-consuming.rs` | Ajouter la substitution de consommation précédente : obtenir tête et queue avec `Rc::unwrap_or_clone` et déplacer les champs consommés. |
| `ListOps-reuse.rs` | Pour un élément retenu, si `Rc::get_mut` réussit, conserver sa tête et remplacer sa queue par l'accumulateur. La cellule devient le nouvel accumulateur. Les autres cas suivent le chemin de consommation. |

Les deux boucles traitées sont le `filterEvens` public et sa copie introduite dans `sumEvens` par l'inlining. `range`, `foldl`, les autres adaptateurs, les closures/thunks, le boxing `Value` et les clones de l'accumulateur restent inchangés. Il ne s'agit pas d'une fusion range/filter/fold.

Les conversions des deux arguments précèdent tout `drop`, y compris celui du premier propriétaire : on ne raccourcit pas sa durée de vie avant l'acquisition du second argument. Le périmètre prouvé est ce filtre fermé `List Int`, dont la tête contient un entier et dont le prédicat est le test de parité généré. Il ne prouve pas qu'un même déplacement de `drop` serait correct pour n'importe quel `Value::Class`, payload avec destruction observable ou callback arbitraire.

La réutilisation requiert l'exclusivité vérifiée par `Rc::get_mut`, donc l'absence de propriétaire fort supplémentaire **et** de `Weak`. Une cellule avec observateur faible suit le chemin existant. Une racine forte conservée empêche toute réutilisation ; une queue conservée permet uniquement celle du préfixe exclusif.

## Validation native exécutée

`python3 probe.py build` puis `python3 probe.py check` ont réussi. Pour chaque variante, le module ordinaire et le module instrumenté ont été compilés et exécutés : 288 cas par binaire, plus six appels à `sumEvens`. Les tailles sont 0, 1, 2, 31, 900 et 9 000 ; les proportions d'éléments retenus sont 0 %, 50 % et 100 %. Les seize combinaisons de partage couvrent racine forte conservée, queue indépendante à `n/3`, `Weak` de la racine et `Weak` de cette queue.

Les assertions comparent la liste complète dans l'ordre à une référence Rust immutable, vérifient l'intégrité des anciennes versions et des queues, les possibilités d'upgrade des `Weak`, leur expiration finale, les chemins uniques/partagés et le nombre exact de nouvelles cellules. Les cas 100 % retenus vérifient notamment qu'une cellule observée par `Weak` ne peut être réutilisée. Les cas 0 et 1 couvrent les frontières où racine et queue peuvent désigner la même cellule. Le harnais utilise une pile de 64 Mio pour la destruction récursive des longues listes conservées, sans changer la génération.

Le comptage des allocations `List` utilise les adresses brutes rendues par l'allocateur, dans un tableau statique sans allocation. Il n'ajoute aucun propriétaire ni `Weak` qui désactiverait artificiellement la réutilisation. Les modules destinés aux mesures du runner sont les fichiers sans suffixe `-counted`.

Profil de cette validation : `rustc -C opt-level=1 -C debuginfo=2`, dépendances Rust figées de l'étude précédente, allocateur `System`. Le runner complet applique son propre profil avec mimalloc. Les hashes des sources, du harnais, du script et des huit exécutables sont dans `metadata.json` ; les observations et assertions finales sont dans `validation.json` et les logs individuels.

## Résultats des compteurs

Pour le vrai `sumEvens`, incluant son `range`, son filtre et son fold inchangés :

| Variante | n | Visites uniques / partagées | Cellules réutilisées | Allocations totales | Allocations `List` |
| --- | ---: | ---: | ---: | ---: | ---: |
| baseline | 900 | 0 / 900 | 0 | 1 818 | 1 352 |
| lifetimes | 900 | 900 / 0 | 0 | 1 818 | 1 352 |
| consuming | 900 | 900 / 0 | 0 | 1 818 | 1 352 |
| reuse | 900 | 900 / 0 | 450 | 1 368 | 902 |
| baseline | 9 000 | 0 / 9 000 | 0 | 18 018 | 13 502 |
| lifetimes | 9 000 | 9 000 / 0 | 0 | 18 018 | 13 502 |
| consuming | 9 000 | 9 000 / 0 | 0 | 18 018 | 13 502 |
| reuse | 9 000 | 9 000 / 0 | 4 500 | 13 518 | 9 002 |

Pour `filterEvens` appelé sur 900 éléments, dont 50 % retenus :

| Partage externe | Cellules admissibles réutilisées | Nouvelles cellules `List`, baseline → reuse |
| --- | ---: | ---: |
| Aucun | 450 | 451 → 1 |
| Racine conservée | 0 | 451 → 451 |
| Queue conservée à l'index 300 | 150 | 451 → 301 |

La cellule nouvelle restante dans le cas exclusif est le `Nil` initial de l'accumulateur. Les proportions 0 % et 100 % suivent la même règle : une allocation supprimée par `Cons` retenu et exclusif, zéro pour un élément rejeté ou partagé.

Toutes les cellules `List` sont libérées après destruction du résultat, des anciennes versions et des observateurs faibles. Le harnais observe aussi des allocations **hors `List`** encore vivantes : trois par appel au filtre, six par appel à `sumEvens` après son premier appel. Ce premier appel (`n=0`) en conserve douze. Ces résidus sont identiques entre les quatre variantes pour chaque observation correspondante ; leur origine précise n'est pas étudiée ici. Il ne faut donc pas conclure à la libération de tout le tas.

## Passage au runner complet

`full-runner-variants.json` contient les quatre modules complets sous les noms `list-baseline`, `list-lifetimes`, `list-consuming` et `list-reuse`. Le coordinateur a construit et mesuré ces mêmes modules avec sa config sous les noms `before`, `list_lifetimes`, `list_consuming` et `list_reuse`. Les étapes restent séparées afin d'attribuer les effets de durée de vie, consommation et réutilisation. Les [résultats du runner](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/full-runner-results.json) et ceux du harnais de partage sont analysés ci-dessous.

Ce prototype établit une condition manquante dans la première expérience : les propriétaires temporaires de l'adaptateur masquaient l'unicité réelle du parcours. Une règle générique éventuelle doit prouver la légalité de ces durées de vie depuis l'IR typé, puis la consommation et le remplacement des champs ; les noms et le motif précis de ListOps utilisés par ce script scratch ne constituent pas une règle de compilateur.

## Harnais séparé des profils de partage

`share-timing.rs` utilise exactement les quatre modules ordinaires déjà livrés, avec mimalloc et `-C opt-level=1 -C debuginfo=2`. Ses dépendances viennent du runner complet figé de la présente étude. `share_timing.py build` enregistre les hashes des dépendances directes, modules, harnais, script et exécutables dans `sharing-build.json`. Il ne modifie aucun des modules livrés.

`share_timing.py smoke` vérifie, sans lecture d'horloge, 72 cas par variante : six tailles, trois proportions et quatre partages. Les quatre scénarios sont une entrée exclusive, une racine conservée, une queue conservée à `n/3`, ou un `Weak` sans propriétaire fort supplémentaire. Pour le scénario `Weak`, l'observateur cible un élément pair conservé, à l'index `(n/3)|1` pour les tailles mesurées : il exclut réellement une cellule de la réutilisation.

La commande de mesure, à lancer uniquement après coordination, est `python3 share_timing.py time --coordinated`. La matrice temporelle contient deux tailles (900, 9 000), quatre scénarios et 50 % d'éléments retenus. Chaque processus réalise une chauffe d'un batch complet puis sept échantillons ; un batch contient 100 opérations pour n=900 et 10 pour n=9 000. Chaque opération chronométrée comprend construction de l'entrée et des propriétaires/Weak, filtre, somme empruntée du résultat, puis destruction du résultat, des propriétaires conservés et du Weak. Aucun propriétaire `List` de l'opération ne survit à cette limite. La préparation du thread, sa pile de 64 Mio, la chauffe, les assertions de checksum après le chrono et l'impression sont exclues.

Quatre tours ont fait tourner l'ordre des quatre variantes, à profil fixe entre les tours et entre les profils. Cela représente 128 processus, chacun avec une chauffe et sept échantillons. Les [résultats bruts](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/list/sharing-timings.json) sont sauvegardés après chaque profil ; le [résumé](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/list/sharing-summary.json) agrège d'abord les sept échantillons de chaque processus par médiane, puis les quatre médianes de processus par profil/variante. Il y a donc quatre répétitions de processus par comparaison, chacune comprenant sept échantillons internes ; les 28 échantillons ne sont pas présentés comme 28 expériences indépendantes.

## Temps du runner complet

Huit blocs ont exécuté les quatre variantes avec ordre tournant puis inversé, en conservant le warmup et le meilleur de dix du benchmark. Le tableau donne la médiane des huit résultats, en microsecondes. La résolution de cette sortie est de 1 µs.

| Variante | List Processing, n=900 | Écart à la baseline appariée | Écart à l'étape précédente |
| --- | ---: | ---: | ---: |
| baseline | 36 | — | — |
| lifetimes | 29 | −19,44 % | −19,44 % |
| consuming | 29 | −19,44 % | 0 % |
| reuse | 25 | −30,56 % | −13,79 % |

La durée de vie précise gagne dans huit blocs sur huit par rapport à la baseline. La consommation ajoutée ne produit pas de gain médian supplémentaire : un bloc est favorable, quatre égaux et trois défavorables par rapport à `lifetimes`. La réutilisation gagne dans huit blocs sur huit par rapport à `consuming`, et dans huit sur huit par rapport à la baseline. Les différences sont attribuées aux étapes comparées ; ces mesures ne profilent pas les instructions responsables du gain de `lifetimes`, dont le nombre d'allocations reste identique.

**Le total du runner ne mesure pas le gain de ListOps.** Sa médiane observée passe de 11 501,5 à 11 532,5 µs entre baseline et reuse, alors que le seul benchmark modifié gagne 11 µs. Le RBTree inchangé représente environ 92 % du total ; sa médiane varie de 10 619 à 10 649,5 µs dans cette campagne. Ces variations dominent le total : elles ne sont pas attribuées à la transformation ListOps.

Le [README officiel, section Rust](/Users/0x1/Documents/htdocs/altbak.pub-purust/README.md:132) documente historiquement List Processing à environ **34 µs compilé**, **76 µs natif de style fonctionnel** et **1 µs dans la dernière colonne native écrite à la main**. Le 36 µs de la campagne est une nouvelle baseline appariée ; les pourcentages ci-dessus sont calculés à partir d'elle, sans mélanger les sessions. Le prototype à 25 µs laisse donc un écart important avec cette dernière référence historique. Il conserve le range, le boxing et les adaptateurs du code généré ; ce résultat ne constitue pas une comparaison contemporaine à travail et représentation identiques avec le natif.

## Temps par profil de partage

Les valeurs suivantes sont des **microsecondes par opération complète** du harnais séparé : construction, filtre, somme empruntée de sortie, puis libération des cellules, des propriétaires conservés et du Weak. Elles ne doivent pas être substituées aux temps de `sumEvens` du runner : le harnais construit directement l'entrée et emploie sa propre somme. Pour chaque profil, seuls les quatre modules générés varient.

| n | Partage | baseline | lifetimes | consuming | reuse | reuse / baseline |
| ---: | --- | ---: | ---: | ---: | ---: | ---: |
| 900 | Exclusif | 26,70 | 19,13 | 18,41 | 14,31 | −46,41 % |
| 900 | Racine conservée | 25,15 | 24,70 | 27,74 | 28,09 | +11,68 % |
| 900 | Queue conservée à n/3 | 24,19 | 22,14 | 23,61 | 22,30 | −7,82 % |
| 900 | Weak sur élément retenu | 26,06 | 19,57 | 18,94 | 15,32 | −41,21 % |
| 9 000 | Exclusif | 242,63 | 205,24 | 201,92 | 160,34 | −33,91 % |
| 9 000 | Racine conservée | 240,82 | 244,38 | 262,14 | 268,32 | +11,42 % |
| 9 000 | Queue conservée à n/3 | 255,86 | 241,99 | 253,11 | 245,77 | −3,95 % |
| 9 000 | Weak sur élément retenu | 247,48 | 215,76 | 207,10 | 166,82 | −32,59 % |

Les profils exclusifs et Weak gagnent dans les quatre tours. Le Weak n'observe qu'une cellule : celle-ci est exclue de la réutilisation, les autres cellules restent admissibles. Ce résultat ne s'étend pas à une liste dont tous les nœuds seraient observés par des Weak. La consommation ajoutée aux durées de vie apporte ici un petit gain : de −1,62 à −4,01 % selon taille et profil, puis la réutilisation apporte de −19,12 à −22,30 % supplémentaires.

Une racine conservée empêche toute réutilisation et la chaîne complète ralentit dans les quatre tours, aux deux tailles. Le principal surcoût observé apparaît à l'étape `consuming` : **+12,29 %** à n=900 et **+7,27 %** à n=9 000 par rapport à `lifetimes`. Le passage supplémentaire à `reuse` change encore la médiane de +1,25 % et +2,36 %, avec des tours de signes mixtes. Les compteurs prouvent le maintien du chemin partagé ; aucune analyse machine ne décompose précisément son surcoût.

Une queue conservée n'autorise la réutilisation que dans le premier tiers. La chaîne complète gagne face à la baseline dans quatre tours sur quatre, mais **ne bat pas l'étape `lifetimes` seule** : +0,71 % à n=900 et +1,56 % à n=9 000 en médiane. Dans ce mélange de préfixe unique et de suffixe partagé, la consommation augmente les médianes de +6,63 % et +4,60 % ; l'étape de réutilisation en récupère une partie, sans que le protocole mesure séparément le temps du préfixe et du suffixe. Dans ce profil, garder seulement les durées de vie précises reste la meilleure médiane mesurée.

Les ratios 0 % et 100 % sont validés fonctionnellement et par compteurs, sans mesure temporelle. Les trois allocations hors `List` restant vivantes par opération du filtre sont inchangées ; elles s'accumulent pendant les batches et sont récupérées à la fin du processus par le système. La borne du protocole et les mêmes nombres d'opérations limitent ce biais entre variantes, mais ces résultats ne démontrent pas une consommation mémoire stable dans un processus de durée arbitraire. Toutes les cellules `List`, y compris celles retenues ou observées, sont bien libérées selon la validation distincte.

## Rapport avec l'essai précédent

L'[étude précédente](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-evaluation-20260914/REPORT.md:58) reste un résultat séparé : **36 → 38 µs (+5,56 %)** pour la consommation seule, avec huit blocs défavorables sur huit. Elle conservait les propriétaires `_a0`/`_a1` dans les adaptateurs pendant `_f` ; aucun nœud n'atteignait donc le chemin unique. Il ne faut ni remplacer son chiffre par celui-ci, ni présenter ses 38 µs comme une étape de la nouvelle série appariée.

La nouvelle chaîne commence par lever cette rétention temporaire. Les compteurs passent alors de zéro à 900 nœuds uniques dans `sumEvens(900)`, puis la réutilisation retire 450 allocations. La conclusion change pour **cette chaîne et cette entrée consommée** : elle accélère le vrai benchmark. Elle ne transforme pas la consommation systématique d'un nœud partagé en bonne stratégie ; la régression de 11 % avec racine persistante confirme cette limite. Une éventuelle règle générique devra établir les durées de vie légales puis conserver un traitement approprié aux chemins partagés. Aucun changement de compilateur n'est inclus dans cette étude scratch.
