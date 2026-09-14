# Bilan consolidé : Perceus, FBIP et sticky sharing pour Purust

14 septembre 2026. Ce rapport reprend la première étude et ajoute les expériences qui lui manquaient. **Les mécanismes expérimentés dans les sections 3 à 7 ont maintenant un cas exécuté, des contrôles fonctionnels et une comparaison temporelle.** Cela couvre un corpus défini ; ce n'est ni une validation de tout Koka/Perceus, ni une intégration dans le compilateur.

**Le résultat nouveau le plus directement exploitable est ListOps : 36 → 25 µs, soit −30,56 %, en corrigeant les durées de vie puis en réutilisant les cellules exclusives.** Le gain absolu est de 11 µs et ne suffit pas à faire ressortir une amélioration du total de la suite. La même transformation ralentit les listes dont l'ancienne racine est conservée. Le choix du chemin selon la propriété et le partage est donc central.

Les parcours empruntés conservent leur intérêt sur les lectures répétées d'arbres. Les essais FBIP confirment l'intérêt de réutiliser les cellules et de conserver les champs inchangés, avec une partie de ces mécanismes déjà intégrée à Purust. Les nouvelles mesures ne justifient toujours pas de prioriser davantage le sticky sharing.

**Aucune modification de Purust ni de sa sortie compilée courante : 868 fichiers Purust et 611 fichiers générés/manifests/verrou vérifiés inchangés par SHA-256.** Les expériences restent dans ce scratch. [Vérification](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/no-product-changes.json). Le [rapport précédent](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-evaluation-20260914/REPORT.md) et ses données sont conservés.

## 1. Ce qui existe déjà dans Purust

Le nom `PerceusPtr` ne signifie pas que toute la chaîne du compilateur Perceus est intégrée. Les [pastilles actuelles](/Users/0x1/Documents/htdocs/altbak.pub/todo.md:39) et le code donnent cet état :

| Mécanisme | État actuel | Preuve |
| --- | --- | --- |
| B12, réutilisation des ADT / FBIP | 🟡 Intégré pour certains parents consommés, avec chemin partagé de secours. | [OwnedFields](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/OwnedFields.purs:45), [génération](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs:2528). |
| B13, conserver les champs inchangés | 🟡 Plusieurs formes intégrées : scalaire, appel enfant fermé, certaines permutations LL/RR. La couverture reste partielle. | [ReuseFields](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/ReuseFields.purs:15), [FieldPermutations](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/FieldPermutations.purs:50). |
| B14, déplacements / emprunts / unicité conservée | 🟡 Derniers usages déplacés et certains emprunts exclusifs conservés ; d'autres chemins gardent des clones et revérifications. | [CodeGen](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs:2381). |
| B15, placement et spécialisation des libérations, fusion dup/drop | 🔴 Pas de chaîne générale intégrée. Les déplacements existants ne suffisent pas à lui attribuer cette passe. | [Audit](/Users/0x1/Documents/htdocs/altbak.pub/optimization-audit.md:92). |
| B22 et B14, records | Setters directs intégrés, copie si partagé, certains chemins fermés et lectures scalaires empruntées ; pas un déboxing général. | [RecordUpdates](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/RecordUpdates.purs:30), [RecordBorrows](/Users/0x1/Documents/htdocs/purust/purust/src/Purust/RecordBorrows.purs:24). |
| B49, sticky sharing | 🟡 Saturation réelle uniquement dans PerceusPtr local, à `u32::MAX`. Le mode threaded délègue à Arc ; RBTree utilise Rc. | [local.rs](/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr/src/local.rs:69), [threaded.rs](/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr/src/threaded.rs:5). |

Le runtime rangé sous `tests/` est bien embarqué dans les sorties. Les optimisations voisines restent distinctes : B41 partage limité des constructeurs nullaires 🟡 ; B42 reconstruction identique renvoyée telle quelle 🔴. Elles ne sont pas créditées comme de nouvelles passes testées ici. Aucune pastille n'a été changée par cette étude.

## 2. Corpus réellement exécuté

| Expérience | Validation | Mesure |
| --- | --- | --- |
| Vrai module ListOps, quatre étapes cumulatives | 288 cas × 4 variantes × 2 exécutables, plus six tailles de sumEvens par exécutable ; ratios 0/50/100 %, partage fort et Weak. Harnais temporel : 72 cas × 4 variantes supplémentaires. | Runner complet : 8 blocs × 4 variantes, 14 sorties vérifiées à chaque run. Harnais partagé : 2 tailles × 4 partages × 4 variantes × 4 blocs, 7 échantillons par processus. |
| Petit IR typé à opérations RC explicites | 56 programmes intermédiaires contrôlés ; 4 IR invalides rejetés. 96 cas × 7 étapes × 2 représentations = 1 344 exécutions avec oracle structurel et bilans RC. | 5 fonctions × 3 partages × 7 étapes × 7 blocs ; 3 échantillons par profil et processus. |
| FBIP, mises à jour d'arbres natifs | 4 variantes × 6 partages × 2 profondeurs × 2 nombres d'opérations, exécutés avec et sans compteurs : 192 cas. Oracle indépendant et bilans des propriétaires/Weak. | Les 24 combinaisons, 4 variantes, 8 blocs, 7 échantillons par processus. |
| Sticky sharing, trois variantes | 18 cas avec compteurs et 18 sans, résultats identiques et rétention/libération vérifiées. | 6 scénarios × 3 variantes × 6 blocs ; un million d'itérations ou de clones simultanés après chauffe. |
| RBTree, première étude conservée | Arbre réel généré : quatre consommateurs, invariants, anciennes versions, partage mixte, Weak et destructions. Contre-exemples à la généralisation aux effets. | Runner complet en 8 blocs ; autre scénario à 100 000 insertions et 20 lectures, en 8 blocs. |

Les comptes de cas ne sont pas des observations statistiques indépendantes. Ils décrivent la couverture fonctionnelle. Les temps viennent exclusivement des exécutables sans instrumentation.

## 3. ListOps : le goulot concret désormais prouvé

Le témoin est une copie exacte du module généré. Quatre versions cumulatives modifient seulement les deux copies du filtre, publique et intégrée dans sumEvens :

1. **Témoin** : code généré inchangé.
2. **Durées de vie** : extraire les deux arguments typés, puis libérer les propriétaires temporaires des adaptateurs avant l'appel.
3. **Consommation** : déplacer les champs des nœuds consommés via `Rc::unwrap_or_clone`.
4. **Réutilisation** : conserver la cellule `Cons` retenue quand `Rc::get_mut` établit son exclusivité, avec chemin de secours pour le partage et les Weak.

Range, fold, représentation, boxing, fermetures et autres adaptateurs restent identiques. Le prédicat est le test de parité fermé sur List Int. Ce script scratch ne constitue pas une règle générale pour toutes les fonctions ou tous les `Value::Class`.

### Runner altbak complet

| Étape | List Processing, médiane | Écart au témoin | Blocs favorables |
| --- | ---: | ---: | ---: |
| Témoin | 36 µs | — | — |
| Durées de vie | 29 µs | −19,44 % | 8/8 |
| + consommation | 29 µs | −19,44 % | 7/8 |
| + réutilisation | **25 µs** | **−30,56 %** | **8/8** |

La consommation seule n'ajoute ici aucun gain médian après correction des durées de vie. La réutilisation ajoute ensuite −13,79 % par rapport à cette étape. Le témoin varie entre 34 et 37 µs ; la version complète entre 25 et 26 µs.

Les compteurs montrent pourquoi la piste devient possible : le témoin n'atteint aucun nœud unique, contre 900 après correction des adaptateurs. La réutilisation supprime **450 allocations de Cons**, soit 1 352 → 902 cellules List et 1 818 → 1 368 allocations globales comptées pour sumEvens(900). À n=9 000, elle supprime 4 500 cellules. Toutes les cellules List sont libérées.

Le total ne montre pas de gain établi : somme des médianes 11,4935 → 11,5185 ms ; médiane des totaux réellement observés 11,5015 → 11,5325 ms. Le RBTree inchangé passe de 10,6190 à 10,6495 ms en médiane. On n'attribue pas ces différences au filtre, et on ne transforme pas ses 11 µs gagnées en promesse sur le total. [Données complètes](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/full-runner-results.json).

### Le partage change la conclusion

Ce second harnais inclut construction directe de la liste, filtre généré, somme et destructions. Il ne remplace pas le runner sumEvens. Temps par opération, en µs :

| Propriété / taille | Témoin | Durées de vie | + consommation | + réutilisation | Écart final |
| --- | ---: | ---: | ---: | ---: | ---: |
| Unique, 900 | 26,698 | 19,126 | 18,411 | 14,306 | −46,41 % |
| Racine conservée, 900 | 25,148 | 24,703 | 27,739 | 28,086 | **+11,68 %** |
| Queue conservée, 900 | 24,187 | 22,137 | 23,606 | 22,295 | −7,82 % |
| Weak intérieur, 900 | 26,059 | 19,566 | 18,940 | 15,319 | −41,21 % |
| Unique, 9 000 | 242,631 | 205,244 | 201,917 | 160,344 | −33,91 % |
| Racine conservée, 9 000 | 240,825 | 244,377 | 262,135 | 268,319 | **+11,42 %** |
| Queue conservée, 9 000 | 255,860 | 241,988 | 253,110 | 245,767 | −3,95 % |
| Weak intérieur, 9 000 | 247,481 | 215,756 | 207,098 | 166,823 | −32,59 % |

Chaque gain ou régression finale de cette table garde son signe dans les quatre blocs. Avec une racine conservée, aucune cellule n'est réutilisable. Avec une queue conservée, seul le préfixe exclusif l'est, et corriger les durées de vie sans ajouter la consommation reste légèrement meilleur que la chaîne complète. Un Weak ciblant un élément retenu interdit bien la réutilisation de sa cellule.

Le harnais révèle aussi des allocations hors List encore vivantes, identiques entre variantes : trois par appel au filtre et six par appel de somme après initialisation. Leur origine précise n'est pas étudiée ici ; on ne prétend pas avoir libéré tout le tas. [Rapport ListOps, contrôles et mesures](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/list/REPORT.md).

### Reprise du premier bilan ListOps

Le premier essai ajoutait seulement la consommation, sans libérer les propriétaires des adaptateurs : **36 → 38 µs, +5,56 %, huit blocs sur huit défavorables**. Les 944 itérations de somme et 15 072 itérations de filtre observées restaient toutes partagées. Ses 96 cas de partage/Weak et six tailles étaient corrects. Ce résultat reste valide pour cette variante ; la nouvelle étude établit précisément la condition manquante. Les sessions ne sont pas mélangées.

## 4. Perceus : les passes ont été matérialisées, pas seulement imitées à la main

Un petit langage fonctionnel fermé produit un IR ANF typé avec projections, branches, appels récursifs et opérations RC explicites. Sept graphes successifs sont réellement transformés puis utilisés par le générateur Rust : départ naïf, descente des duplications dans les branches utiles, placement des drops au dernier usage, fusion adjacente dup/drop, déballage possédé spécialisé par constructeur, reset/reuse, puis conservation des champs inchangés.

Cela s'applique à huit fonctions déclaratives sur trois layouts : map/filter de listes, map/update/choose d'arbres, puis arbre aux champs réordonnés. Le vérificateur contrôle types, propriété, déplacements et consommation des tokens. L'interpréteur indépendant vérifie les résultats complets et les anciennes valeurs partagées. Les références faibles et la libération finale sont couvertes. Les mécanismes correspondent aux sections 2.2–2.5 du [rapport Perceus](https://www.microsoft.com/en-us/research/wp-content/uploads/2020/11/perceus-tr-v4.pdf), sans en reproduire toutes les garanties formelles.

Sur map d'une liste unique de 64 éléments :

| Mesure logique | Départ | Après drops précis | Après fusion | Après déballage spécialisé | Après reuse |
| --- | ---: | ---: | ---: | ---: | ---: |
| Duplications Rc | 256 | 256 | 64 | 0 | 0 |
| Nouvelles cellules | 65 | 65 | 65 | 65 | 0 |
| Pic de payloads vivants | 130 | 65 | 65 | 65 | 65 |

Le placement précis réduit ici la rétention, sans diminuer le nombre de duplications. Le pic mesure les payloads vivants, pas les headers éventuellement conservés par Weak. Sur tree_map unique, les nouvelles cellules passent de 31 à 0, les dups de 105 à 30 puis 0, le pic de 62 à 31.

### Effets temporels des étapes

Les temps incluent entrée construite à neuf, transformation, lecture du résultat et tous les drops. Les listes ont 256 éléments et tree_map 127 nœuds ; les entrées et graines passent par black_box. Ce départ synthétique volontairement naïf **n'est pas le Rust actuel de Purust**.

| Comparaison entre étapes adjacentes | Cas | Écart | Blocs favorables |
| --- | --- | ---: | ---: |
| Drops précis / étape précédente | List map unique | −16,16 % | 7/7 |
| Drops précis / étape précédente | List filter unique | −24,07 % | 7/7 |
| Fusion / drops précis | List map unique | −0,55 % | 7/7 |
| Déballage spécialisé / fusion | List map unique | +0,08 % | 1/7 |
| Reuse / déballage spécialisé | List map unique | −10,36 % | 7/7 |
| Reuse / déballage spécialisé | Tree map unique | −24,04 % | 7/7 |
| Reuse / déballage spécialisé | List map, racine partagée | **+6,13 %** | 0/7 |
| Reuse / déballage spécialisé | List filter, racine partagée | **+6,99 %** | 0/7 |
| Reuse / déballage spécialisé | Tree map, racine partagée | **+13,05 %** | 0/7 |

La chute des dups après fusion ou déballage ne produit donc pas nécessairement un gain temporel net supplémentaire en Rust. Le partage peut rendre reuse défavorable. Sur les petites mises à jour de racine, conserver les champs donne un signal faible : −2,02 %, quatre blocs sur sept favorables ; l'expérience FBIP suivante teste ce mécanisme dans un travail plus conséquent.

La régression structurelle n'a pas été masquée : dans tree_choose partagé, pushdown puis fusion réduisent les dups de 3 à 2 puis 1 ; le déballage spécialisé remonte à 2 parce qu'il clone les deux champs avant le choix. Le prototype devrait préserver cette précision pour pouvoir prétendre à une chaîne plus générale. Ce nombre de dups supplémentaire n'est pas présenté comme l'explication d'une régression temporelle non observée.

Certaines fonctions restent identiques entre deux étapes : leurs différences de durée ne sont pas attribuées à la passe. Les résultats complets et ces contrôles sont dans le [rapport IR](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/REPORT.md), avec tous les [temps bruts](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timings.json).

## 5. FBIP : isoler cellule, unicité et champs conservés

Cette expérience utilise quatre versions du **même arbre natif**, deux enfants et deux scalaires, avec les mêmes chemins et les mêmes opérations : modification scalaire, modification de deux scalaires, permutation d'enfants. Elle ajoute plusieurs régimes de partage au-delà d'une rotation particulière.

| Étape | Différence isolée |
| --- | --- |
| Persistant | Reconstruit et alloue à chaque changement de cellule. |
| Reuse avec revérification | Réutilise une cellule unique mais retire/reconstruit tout le payload, puis revérifie l'unicité. |
| Emprunt exclusif conservé | Même payload reconstruit, sans seconde vérification d'unicité. |
| Champs conservés | Même emprunt, écrit seulement les champs concernés. |

À profondeur 10 et 2 048 opérations, temps en µs, construction et destructions incluses :

| Partage | Persistant | Reuse | Emprunt conservé | Champs conservés |
| --- | ---: | ---: | ---: | ---: |
| Unique | 204,562 | 83,688 | 78,271 | 52,876 |
| Huit anciennes racines conservées | 208,917 | 110,666 | 107,999 | 77,250 |
| Sous-arbres intérieurs conservés | 204,188 | 94,501 | 86,874 | 57,812 |
| Enfant commun aux deux branches | 205,458 | 92,770 | 85,041 | 57,042 |
| Weak de racine | 204,145 | 83,625 | 78,125 | 52,667 |
| Weak intérieurs | 206,812 | 83,916 | 79,334 | 52,020 |

La première réutilisation gagne 47,03–59,42 %. Conserver ensuite le borrow ajoute 2,41–8,33 %, plus variable. Conserver enfin les champs ajoute 28,47–34,43 % par rapport au borrow déjà conservé. Pour reuse et champs, les six scénarios sont favorables dans les huit blocs. Les petites dimensions ont des effets plus fragiles, documentés séparément.

Les compteurs isolent les contributions : en unique, allocations **15 703 → 2 047** dès reuse ; ensuite le borrow évite **13 656 secondes vérifications**, puis les champs conservés évitent **13 656 extractions et reconstructions de payload**. Ces deux dernières étapes ne changent ni les allocations ni les dups. Avec anciennes racines, 1 882 copies restent nécessaires, et les invariants persistants passent.

**Ce sont des preuves de l'intérêt des mécanismes, pas une promesse de −74 % sur le RBTree de Purust.** Son témoin bénéficie déjà de plusieurs optimisations B12/B13/B14. Il faudrait identifier un motif admissible encore non couvert pour convertir ce résultat synthétique en gain supplémentaire. [Rapport FBIP, toutes dimensions](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/REPORT.md).

## 6. RBTree : reprise du bilan précédent

Le précédent prototype modifiait uniquement la fonction depth du vrai module généré. Construction, insertion, rotations, représentation et allocateur restaient identiques. Quatre versions : témoin avec clones Rc, parcours par emprunt, consommation par unwrap_or_clone, hybride consommant les cellules uniques et empruntant les sous-arbres partagés.

| Variante | Runner complet : RBTree | Écart | Construction + 20 lectures partagées + destruction | Écart |
| --- | ---: | ---: | ---: | ---: |
| Témoin | 10,9945 ms | — | 17,077 ms | — |
| Emprunt | 11,0190 ms | +0,22 % | 13,871 ms | −18,78 % |
| Consommation | 10,8780 ms | −1,06 % | 17,778 ms | +4,11 % |
| Hybride | 10,8475 ms | −1,34 % | 13,842 ms | −18,95 % |

Huit blocs pour chaque expérience. L'hybride est favorable 7/8 sur le RBTree du runner, avec plages qui se recouvrent ; son gain est de 0,147 ms. La somme des quatorze médianes passe de 11,880 à 11,753 ms, soit −1,07 %. C'est un petit signal, sans supériorité systématique établie sur les autres variantes.

Sur vingt lectures d'une racine conservée après 100 000 insertions, emprunt et hybride sont favorables 8/8, avec plages entièrement sous le témoin. Le bénéfice vient essentiellement de l'emprunt : l'hybride n'ajoute que −0,21 % par rapport à lui, cinq blocs sur huit favorables. La consommation systématique est défavorable 8/8.

La première série avant hybride donnait consommation 11,0435 → 10,7835 ms (−2,35 %, 7/8) et emprunt 10,820 ms (−2,02 %, 6/8). Elle reste distincte et montre la variabilité des petits gains ; elle n'est pas fusionnée avec la série finale.

En unique, les prototypes suppriment 200 000 clones et libèrent les mêmes 100 001 cellules. Sur une racine partagée de 200 nœuds, témoin/consommation font 401 clones contre un clone d'appelant pour emprunt/hybride. Les invariants rouge/noir, 200 versions persistantes, 512 insertions en partage mixte, les Weak et libérations finales passent. Des contre-exemples exécutés montrent qu'avec destructeurs observables ou callbacks, changer ces durées de vie peut changer le résultat observable. Le périmètre fermé typé reste essentiel. [Rapport et données initiaux](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-evaluation-20260914/b15/REPORT.md).

## 7. Sticky sharing : confirmation et ablation supplémentaire

Le runtime local sature à **4 294 967 295 propriétaires**. Les lectures et mises à jour naturelles restent à un–trois propriétaires. Le nouveau cas fanout conserve les clones ensemble : 100 001 propriétaires simultanés vérifiés avec compteurs ; un million de clones dans la mesure. Cela n'active toujours pas naturellement sticky sharing.

Trois variantes gardent la même représentation : runtime normal ; gardes sticky retirées avec incrément checked ; trois cellules initiales forcées au seuil. Six blocs équilibrés, un million d'itérations après chauffe, O1/overflow checks/System, séparément du runner mimalloc.

| Scénario | Normal | Sans gardes sticky, checked | Forcé sticky |
| --- | ---: | ---: | ---: |
| Lectures avec clones | 12,964 ms | 14,331 ms (+10,54 %) | 13,937 ms (+7,50 %) |
| Lectures empruntées | 4,624 ms | 4,620 ms (−0,09 %) | 4,611 ms (−0,28 %) |
| Lectures via fermeture | 13,561 ms | 15,159 ms (+11,78 %) | 14,774 ms (+8,94 %) |
| Updates uniques | 39,799 ms | 40,876 ms (+2,71 %) | 39,763 ms (−0,09 %) |
| Version initiale conservée | 39,528 ms | 40,674 ms (+2,90 %) | 39,949 ms (+1,07 %) |
| Un million de clones simultanés | 8,570 ms | 8,422 ms (−1,73 %) | 8,763 ms (+2,25 %) |

Les ralentissements de lecture possédée et via fermeture sous saturation forcée sont présents dans les six blocs. Les petits gains médians ne sont pas réguliers. Retirer les gardes ne donne pas non plus une amélioration générale ; checked conserve une protection de débordement et n'est pas une proposition de nouveau runtime.

La première série de sept paires donnait normal → forcé : **13,866 → 15,249 ms** en lectures possédées (+9,97 %) ; **4,626 → 4,623** par emprunt ; **15,659 → 16,111** via fermeture (+2,89 %) ; **41,067 → 40,783** en update unique ; **41,464 → 41,707** avec version initiale conservée. Les deux régressions de lecture étaient présentes 7/7. Ces résultats antérieurs sont conservés comme session distincte.

Forcer les trois cellules initiales les rend non libérables : **168 octets par graphe**, 336 octets par processus avec chauffe. La première mise à jour unique ajoute trois copies ; elles retrouvent ensuite un compteur normal. Le compteur logique des écritures diminue sous saturation mais le binaire ne devient pas plus rapide ; la cause machine n'a pas été profilée. Aucun résultat n'est extrapolé aux threads, à Arc ou au RBTree Rc. [Rapport sticky complet](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/REPORT.md).

## 8. Baselines officielles et conditions de mesure

Le [README officiel](/Users/0x1/Documents/htdocs/altbak.pub/README.md:133), relu pour ce bilan, donne :

| Benchmark | Rust compilé | Dernière colonne native |
| --- | ---: | ---: |
| List Processing | 34 µs | 1 µs |
| Deep Record Updates | 382 µs | 4 µs |
| RBTree | 11 631 µs | 36 070 µs |
| Total | 12,56 ms | 36,13 ms |

RBTree étant déjà plus rapide que cette référence native, ses essais comparent le même code généré à ses variantes. ListOps reste loin de la référence native même après le progrès mesuré ; ce prototype n'en supprime pas tous les coûts. Le run utilisateur à 12,198 ms pour RBTree et 13,17 ms au total est encore une autre session. Aucun gain n'est calculé en soustrayant deux sessions historiques.

Machine locale macOS arm64, rustc 1.96.0. Runner, ListOps, IR et FBIP emploient O1/debug/mimalloc ; sticky emploie le profil isolé System décrit plus haut. Dans chaque expérience, représentation, algorithme, charge et allocateur sont constants entre variantes, hormis les transformations explicitement comparées. IR/FBIP synthétiques et code réellement généré restent séparés.

Les préparations/builds/smokes sont terminés avant les mesures ; les cinq campagnes ont été exécutées **en série**, sans builds, validations ou instrumenteurs de cette étude en parallèle. L'activité extérieure n'est pas contrôlée. Les ordres tournent entre variantes ; runner/FBIP/sticky ajoutent des rotations inversées. Les listes tournent sur quatre blocs et l'IR sur sept. Les temps rapportés sont les médianes des médianes de processus, sauf le runner qui conserve son protocole meilleur de dix. Les données individuelles, ordres et empreintes sont conservés.

La validation est bornée : le petit IR est first-order, pur, sans cycles, effets, concurrence, FFI ni destructeurs opaques ; sa fusion traite les paires adjacentes et ses tokens n'échappent pas. Les tests natifs n'établissent pas une analyse interprocédurale générale ni la correction sur les panics. Aucun frontend TAST n'est branché au prototype ; une éventuelle intégration devra tirer ses preuves du TAST typé existant, sans supposer un effacement des types.

## 9. Priorités que ces preuves justifient

1. **Corriger les durées de vie des adaptateurs admissibles**, puis étendre la réutilisation aux filtres exclusifs. C'est la piste nouvelle la mieux reliée à un benchmark actuel : −19,44 % avec les durées de vie seules, −30,56 % avec la chaîne complète. Le chemin partagé doit être traité séparément pour éviter le +11 % observé.
2. **Étendre les parcours empruntés aux lectures d'arbres partagés**, éventuellement avec une entrée hybride : près de −19 % sur construction + vingt lectures, mais seulement un petit signal sur le benchmark actuel.
3. **Étendre B13/B14 là où des reconstructions admissibles restent présentes**. La conservation des champs et de l'emprunt a un intérêt isolé établi. Son gain supplémentaire dans Purust reste à mesurer pour chaque motif encore manquant ; le FBIP intégré n'est pas à refaire.
4. **Utiliser le prototype IR comme démonstrateur et corpus de régression pour B15**, pas comme justification automatique d'un port complet de Perceus. Le bénéfice du placement précis est réel sur ce corpus, alors que fusion/déballage apportent parfois peu de temps gagné et que reuse peut pénaliser le partage.
5. **Ne pas prioriser de nouveau travail sticky sharing pour ces charges.** Le seuil naturel n'est pas atteint et l'essai forcé ne convainc pas. Les mesures ne justifient pas non plus de retirer le mécanisme existant.

Le premier bilan était donc incomplet sur les passes générales et les interactions ; cette étude comble ces essais manquants. Elle change surtout la conclusion sur ListOps, en identifiant puis en testant la condition d'unicité absente de la première variante. Elle confirme la nécessité de distinguer structures uniques et partagées.

## Reproduction et données

[Résumé calculé](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/summary.json), [script de résumé](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/summarize.py), [snapshot commun](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/full-runner-snapshot.json), [constructions du runner](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/full-runner-build.json), [verrou](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/full-runner-lock.json), [vérification des sources](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/verify_unchanged.py). Les grands binaires et copies sont dans `build/`, ignoré par Git. Les quatre sous-rapports décrivent leur préparation et leurs contrôles.

Depuis ce dossier, après préparation et arrêt des autres mesures, lancer séparément et successivement :

```sh
python3 full_runner.py time --blocks 8 --coordinated
python3 list/share_timing.py time --coordinated
python3 fbip/probe.py time --coordinated --blocks 8
python3 ir/timing.py time --blocks 7 --samples 3
python3 sticky/probe.py time --coordinated
python3 summarize.py
python3 verify_unchanged.py
```

Les anciennes données restent dans le scratch de la première étude. Les nouvelles mesures n'en écrasent aucune.
