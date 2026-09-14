# Expérience bornée d'un IR de comptage précis et de réutilisation

Les **sept étapes mécaniques** produisent du Rust correct sur **huit fonctions et trois layouts** : listes `map/filter`, arbres binaires `map/update/choose`, puis les mêmes opérations sur un arbre dont les champs sont réordonnés. Chaque étape passe **96 scénarios** en comptage et les mêmes 96 avec le vrai `std::rc::Rc`, soit **1 344 scénarios exécutés**. La campagne temporelle couvre cinq fonctions et trois partages sur les sept étapes : **2 205 échantillons**, dont tous les résultats passent l'oracle. Le placement précis des drops puis la réutilisation apportent les plus grands gains sur les listes et arbres uniques. En revanche, le passage à reuse ralentit map/filter avec racine partagée de **6,13 à 13,05 %**, malgré la correction des résultats et des bilans mémoire.

Le point de départ est un **IR synthétique volontairement naïf**, pas le Rust actuellement émis par Purust. Ces différences ne constituent donc pas un gain annoncé sur les benchmarks existants, ni une preuve de couverture de tout Perceus. Aucun compilateur ni sortie courante n'est modifié.

## Ce qui est réellement transformé

[prototype.py](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/prototype.py) définit un petit langage fonctionnel fermé : projections typées, arithmétique `i64` avec addition wrapping, test de parité, branches ordonnées, appel récursif saturé de la fonction courante, construction et retour d'un ADT. [source-ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/source-ir.json) contient les layouts et expressions sources.

Le lowering crée un IR ANF avec `dup_field`, `copy_field`, `dup`, `drop`, appels, constructeurs et terminators explicites. Chaque passe modifie ce graphe ; le printer lit le graphe résultant. Les JSON intermédiaires sont les objets effectivement consommés par l'émetteur, pas une description ajoutée après émission.

| Étape | Transformation | Graphe émis |
| --- | --- | --- |
| naive | Dup des projections et opérandes possédés ; drops en fin de branche. | [naive.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/naive.ir.json) |
| pushdown | Descend `dup_field` dans les seuls bras qui utilisent le champ. | [pushdown.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/pushdown.ir.json) |
| precise | Calcule les usages futurs par branche et place chaque drop après le dernier usage. | [precise.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/precise.ir.json) |
| fusion | Réécrit les paires **adjacentes** `dup y x; drop x` en `move y x`. | [fusion.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/fusion.ir.json) |
| specialized_drop | Remplace projection+drop du parent par un déballage du constructeur : champs déplacés si unique, payload cloné et owner relâché si partagé. | [specialized_drop.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/specialized_drop.ir.json) |
| reuse | Introduit `reset`, un token optionnel de cellule, `reuse` et la libération du token sur les retours sans constructeur. | [reuse.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/reuse.ir.json) |
| retained_fields | Reconnaît les reconstructions dont tous les champs récursifs sont inchangés et conserve ces champs, avec RHS scalaires copiés avant écriture. | [retained_fields.ir.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/retained_fields.ir.json) |

Exemple mécanique dans `list_map` : après placement précis, `dup v3 f1; drop f1` précède l'appel récursif. La fusion remplace exactement cette paire par `move v3 f1`. À l'étape suivante, `dup_field f1 n; drop n` disparaît au profit du déballage possédé. Le même code de passe s'applique aux deux positions de champs récursifs de `Tree` et aux positions différentes de `TreeAlt`.

Ces opérations correspondent à des mécanismes présentés dans les sections **2.2 à 2.5** du [rapport Perceus officiel](https://www.microsoft.com/en-us/research/wp-content/uploads/2020/11/perceus-tr-v4.pdf). L'ordre expérimental est adapté pour isoler leurs contributions ; cette implémentation ne reprend ni la totalité de son calcul linéaire ni ses garanties formelles.

## Validation des types et de la propriété

Un vérificateur parcourt chaque graphe et contrôle les types des champs/opérandes, les usages après déplacement, les propriétaires non libérés, les retours et la consommation des tokens. Les **56 programmes intermédiaires** passent. Quatre IR invalides construits par mutation sont refusés : usage après drop, pointeur dans un champ scalaire, owner oublié et token échappant au retour. [ir-checks.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/ir-checks.json).

L'oracle est un interpréteur Python des expressions sources sur des tuples immuables. La comparaison porte sur la forme complète et toutes les valeurs du résultat. Deux entrées de parités différentes couvrent les deux bras ; les filtres parcourent des listes de 64 éléments et les arbres ont 15 nœuds internes plus leurs feuilles vides.

Pour chaque fonction et entrée : propriétaire unique, ancienne racine conservée, enfant indépendant conservé, Weak de racine, partage mixte et Weak d'enfant. Les anciennes valeurs restent identiques, toute Weak survivante observe son ancienne valeur, et les Weak expirent après libération des derniers propriétaires. Les observations sont faites hors de la fonction pure, jamais dans un callback pendant sa mutation.

[runtime.rs](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/runtime.rs) enveloppe le vrai Rc dans `Option<Rc<T>>`, de la taille d'un mot, afin d'observer les opérations sans doubler les drops après extraction. Les bilans sont vérifiés : `new + dup = drop_last + drop_shared + unwrap_unique`, et `new = drop_last + unwrap_unique` après libération de tous les propriétaires. Un échec `try_unwrap` retourne le propriétaire et n'est pas compté comme une destruction. Le nombre de payloads vivants est également contrôlé. Une Weak peut conserver le header d'allocation après destruction du payload ; le pic annoncé concerne les **payloads vivants**, pas les octets encore retenus par les Weak.

Chaque même source est aussi compilée avec `--cfg untracked` : les références deviennent directement des `std::rc::Rc`, et l'instrumentation disparaît. Il n'y a aucun changement de représentation entre étapes. [Validations et événements](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/validation-counts.json), [empreintes](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/metadata.json).

## Effets isolés observés

Les compteurs ci-dessous portent sur l'appel transformant la valeur ; le pic inclut l'entrée initiale. Ils ne sont pas des instructions machine et ne prouvent pas une amélioration de durée.

| `list_map`, entrée unique de 64 éléments | naive | pushdown | precise | fusion | specialized_drop | reuse | retained_fields |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Nouvelles cellules | 65 | 65 | 65 | 65 | 65 | **0** | **0** |
| Dups Rc | 256 | 256 | 256 | **64** | **0** | **0** | **0** |
| Pic de payloads vivants | 130 | 130 | **65** | 65 | 65 | 65 | 65 |

Le placement précis modifie ici le moment de destruction, sans réduire le total des dups. La fusion réduit ensuite les opérations, et le reset/reuse supprime les nouvelles allocations. Pour `tree_map` unique, la même progression donne **31 → 0 nouvelles cellules**, **105 → 30 → 0 dups**, et un pic **62 → 31**.

Pour `tree_update` unique, `reuse` effectue un reset et une réinstallation du constructeur. `retained_fields` les remplace par une seule mise à jour scalaire : **0 reset, 0 reuse, 1 scalar_update**, sans allocation ni dup. La même sélection fonctionne sur `TreeAlt` avec des indices de champs différents.

### Une régression conservée dans le résultat

`tree_choose` sur une racine partagée passe de **3 dups** à **2** après pushdown, puis **1** après fusion. Le déballage consommant remonte à **2 dups**, car son chemin partagé clone **tous les champs du constructeur**, y compris celui rejeté par le bras choisi. `reuse` et `retained_fields` gardent cette régression locale de compteur. Elle ne devient pas une régression temporelle du même passage : specialized_drop mesure ici **−1,41 %** face à fusion, avec **5/7** blocs favorables. Ce résultat montre une limite concrète du prototype et l'impossibilité de déduire la durée des seuls compteurs.

Sur racine partagée, `list_map` garde **65 allocations et 64 dups** même après reuse. Avec une Weak sur la racine unique, le reset de cette cellule est refusé et une allocation de remplacement subsiste ; les descendants admissibles restent réutilisables. Le fast path de champs conservés exige lui aussi `Rc::get_mut` et garde le fallback en présence de Weak.

Toutes les valeurs par fonction, entrée, scénario et étape sont dans [summary.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/summary.json).

## Limites et reproduction

Le DSL est first-order, pur, cycle-free, sans closures, FFI, callbacks, effets, concurrence ni charges utiles à destructeurs observables. Un ADT possède un constructeur vide et un constructeur de données, avec seulement `i64` et références récursives du même ADT. Les fonctions retournent le même type de pointeur. La réutilisation couvre au plus un constructeur retourné par chemin ; le token privé n'échappe pas. La spécialisation de champs couvre seulement la reconstruction directe dont tous les champs récursifs sont inchangés.

La fusion traite les paires adjacentes ; le pushdown ne calcule pas tous les placements possibles ; le déballage partagé ne conserve pas la précision atteinte pour tous les champs inutilisés. Le frontend contient huit fonctions déclaratives, pas un import de CoreFn/TAST. Le vérificateur et les exécutions fournissent des preuves sur ce corpus borné, pas une preuve de compilation générale. Les contraintes sur les effets restent essentielles, comme l'a montré le précédent essai B15 avec destructeurs et callbacks.

Depuis ce dossier, lancer séparément :

```sh
python3 prototype.py prepare
python3 prototype.py validate
python3 summarize.py
```

`build/<stage>.rs` contient le programme Rust complet ; `build/<stage>-native` est le binaire validé sans instrumentation. Aucune de ces commandes ne lit une horloge. La campagne temporelle décrite ci-dessous a été lancée séparément, après coordination avec les autres tâches.

## Méthode temporelle exécutée

[timing.py](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timing.py) conserve exactement les fonctions Rust validées et remplace seulement leur programme principal par [timing-main.rs](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timing-main.rs). Les sept binaires utilisent le même profil O1 avec informations de debug, mimalloc et `--cfg untracked` : les pointeurs sont directement des `std::rc::Rc`. Les empreintes des sources validées et des binaires sont contrôlées. Les sept smokes ont passé les cinq fonctions × trois modes de partage × deux graines dynamiques, sans lire l'horloge ; [timing-build.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timing-build.json) conserve ce résultat.

La matrice temporelle couvre `list_map`, `list_filter`, `tree_map`, `tree_choose` et `tree_update`, chaque fois avec propriétaire unique, ancienne racine conservée ou enfant conservé. Chaque itération construit une nouvelle entrée, recrée le partage choisi, exécute la transformation, parcourt le résultat pour son checksum et détruit résultat et propriétaires conservés. Toute cette séquence est incluse dans la mesure. Les graines viennent des arguments du processus, les entrées et résultats passent par `black_box`, et l'oracle arithmétique est calculé hors chronométrage. Les listes ont 256 éléments, les arbres de map 127 nœuds de données, ceux de choose/update 7. Les petits arbres rendent visibles les chemins de racine, mais leur construction et leur destruction peuvent diluer les gains du noyau.

La campagne utilise trois échantillons par profil dans chacun de sept blocs à rotations équilibrées : chaque étape apparaît une fois à chaque position. Elle conserve une progression flushée par profil, les temps bruts dans `build/timing/raw-*.log`, un checkpoint atomique `timings-checkpoint.json` après chaque processus et les résultats finaux dans [timings.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timings.json). Chaque échantillon exécute 2 048 itérations pour map/filter, ou 32 768 pour choose/update, après un warmup complet non chronométré. Les médianes par profil vont de 3,598 à 15,681 ms. Commande exécutée :

```sh
python3 timing.py time --blocks 7 --samples 3
```

## Durées et régularité entre blocs

Les temps ci-dessous sont en **millisecondes par échantillon**, médiane des trois échantillons de chaque processus, puis médiane des sept blocs. Les pourcentages sont le rapport de ces médianes ; un bloc est favorable lorsque sa propre médiane est inférieure à celle de l'étape comparée dans le même bloc. Ce nombre décrit la régularité observée, sans constituer un test de significativité. Une baisse du rapport des médianes peut coexister avec une minorité de blocs favorables : ces deux résumés répondent à des questions différentes.

Le mode « mixte » temporel conserve seulement le premier enfant : la racine est unique, une partie des descendants reste partagée. Il correspond à `shared_child` dans les validations, et **pas** à leur scénario nommé `mixed`, qui combine racine/enfant/Weak. Il n'y a aucune Weak dans la matrice temporelle ; les garanties Weak viennent des validations séparées. Une liste conserve ainsi presque toute sa queue ; un arbre conserve seulement son sous-arbre gauche. Aucun total des quinze profils n'est calculé, car leurs volumes et répétitions sont arbitraires et différents.

| Profil | naive | pushdown | precise | fusion | specialized_drop | reuse | retained_fields | Final / naive | Blocs favorables |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `list_map` unique | 15.280 | 15.157 | 12.708 | 12.637 | 12.647 | 11.338 | 11.316 | -25.94 % | 7/7 |
| `list_map` racine partagée | 15.445 | 15.411 | 14.845 | 14.716 | 14.702 | 15.604 | 15.559 | +0.74 % | 2/7 |
| `list_map` mixte (enfant partagé) | 15.576 | 15.459 | 14.939 | 14.841 | 14.779 | 15.681 | 15.639 | +0.40 % | 2/7 |
| `list_filter` unique | 11.577 | 11.658 | 8.852 | 8.845 | 8.825 | 8.627 | 8.541 | -26.22 % | 7/7 |
| `list_filter` racine partagée | 11.508 | 11.558 | 11.159 | 11.135 | 11.118 | 11.895 | 11.885 | +3.28 % | 0/7 |
| `list_filter` mixte (enfant partagé) | 11.536 | 11.537 | 11.192 | 11.073 | 11.116 | 11.802 | 11.731 | +1.69 % | 0/7 |
| `tree_map` unique | 7.691 | 7.598 | 7.477 | 7.363 | 7.172 | 5.448 | 5.540 | -27.97 % | 7/7 |
| `tree_map` racine partagée | 7.490 | 7.619 | 7.209 | 7.457 | 7.293 | 8.245 | 8.231 | +9.89 % | 0/7 |
| `tree_map` mixte (enfant partagé) | 7.454 | 7.562 | 7.244 | 7.146 | 7.019 | 6.909 | 6.957 | -6.67 % | 7/7 |
| `tree_choose` unique | 3.679 | 3.734 | 3.705 | 3.784 | 3.663 | 3.701 | 3.598 | -2.20 % | 5/7 |
| `tree_choose` racine partagée | 3.751 | 3.763 | 3.806 | 3.727 | 3.675 | 3.743 | 3.663 | -2.33 % | 5/7 |
| `tree_choose` mixte (enfant partagé) | 3.724 | 3.894 | 3.785 | 3.741 | 3.656 | 3.720 | 3.695 | -0.80 % | 5/7 |
| `tree_update` unique | 4.050 | 4.033 | 4.002 | 3.930 | 3.930 | 3.710 | 3.635 | -10.23 % | 6/7 |
| `tree_update` racine partagée | 4.086 | 4.029 | 3.896 | 3.952 | 3.966 | 4.009 | 4.005 | -1.99 % | 5/7 |
| `tree_update` mixte (enfant partagé) | 3.930 | 4.065 | 3.965 | 3.923 | 4.029 | 3.793 | 3.748 | -4.64 % | 7/7 |

Écarts de chaque passage face au précédent. Chaque cellule indique **variation de durée / blocs favorables** ; un signe négatif signifie plus rapide. `=` indique que le corps Rust de cette fonction reste identique : son écart ne mesure pas une transformation locale de cette fonction.

| Profil | pushdown | precise | fusion | specialized_drop | reuse | retained_fields |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| `list_map` unique | -0.80 % / 6 = | -16.16 % / 7 | -0.55 % / 7 | +0.08 % / 1 | -10.36 % / 7 | -0.19 % / 4 = |
| `list_map` racine partagée | -0.22 % / 3 = | -3.67 % / 7 | -0.87 % / 6 | -0.10 % / 4 | +6.13 % / 0 | -0.29 % / 3 = |
| `list_map` mixte (enfant partagé) | -0.75 % / 5 = | -3.36 % / 7 | -0.65 % / 7 | -0.42 % / 5 | +6.10 % / 0 | -0.27 % / 5 = |
| `list_filter` unique | +0.69 % / 3 | -24.07 % / 7 | -0.08 % / 3 | -0.22 % / 3 | -2.25 % / 7 | -0.99 % / 4 = |
| `list_filter` racine partagée | +0.44 % / 3 | -3.45 % / 7 | -0.22 % / 6 | -0.15 % / 3 | +6.99 % / 0 | -0.09 % / 4 = |
| `list_filter` mixte (enfant partagé) | +0.01 % / 4 | -2.99 % / 7 | -1.06 % / 6 | +0.39 % / 3 | +6.17 % / 1 | -0.60 % / 5 = |
| `tree_map` unique | -1.21 % / 4 = | -1.60 % / 6 | -1.52 % / 5 | -2.59 % / 5 | -24.04 % / 7 | +1.69 % / 2 = |
| `tree_map` racine partagée | +1.71 % / 2 = | -5.38 % / 7 | +3.44 % / 1 | -2.20 % / 6 | +13.05 % / 0 | -0.16 % / 4 = |
| `tree_map` mixte (enfant partagé) | +1.45 % / 2 = | -4.21 % / 7 | -1.36 % / 2 | -1.77 % / 7 | -1.56 % / 4 | +0.69 % / 3 = |
| `tree_choose` unique | +1.52 % / 3 | -0.78 % / 4 | +2.13 % / 3 | -3.20 % / 5 | +1.03 % / 3 | -2.79 % / 6 = |
| `tree_choose` racine partagée | +0.32 % / 4 | +1.16 % / 4 | -2.09 % / 5 | -1.41 % / 5 | +1.85 % / 2 | -2.12 % / 3 = |
| `tree_choose` mixte (enfant partagé) | +4.56 % / 2 | -2.81 % / 4 | -1.16 % / 3 | -2.27 % / 5 | +1.73 % / 3 | -0.68 % / 4 = |
| `tree_update` unique | -0.41 % / 5 = | -0.77 % / 6 | -1.79 % / 5 | +0.00 % / 4 | -5.59 % / 6 | -2.02 % / 4 |
| `tree_update` racine partagée | -1.40 % / 4 = | -3.29 % / 4 | +1.43 % / 3 | +0.36 % / 4 | +1.09 % / 2 | -0.11 % / 4 |
| `tree_update` mixte (enfant partagé) | +3.43 % / 2 = | -2.47 % / 5 | -1.05 % / 5 | +2.70 % / 2 | -5.87 % / 5 | -1.18 % / 4 |

## Ce que chaque passage apporte réellement

**Pushdown.** Sur les fonctions qu'il change dans cette matrice, aucun gain temporel convaincant n'apparaît : list_filter varie de +0,01 à +0,69 % avec seulement 3–4/7 blocs favorables ; tree_choose varie de +0,32 à +4,56 % avec 2–4/7 blocs favorables. La réduction des dups de branches reste prouvée par les compteurs. Les corps Rust de list_map, tree_map et tree_update sont strictement identiques avant/après pushdown ; leurs écarts ne peuvent pas être attribués à une transformation locale de ces fonctions.

**Placement précis.** C'est le premier grand gain temporel des listes uniques : list_map **15,157 → 12,708 ms (−16,16 %, 7/7)** et list_filter **11,658 → 8,852 ms (−24,07 %, 7/7)** face à pushdown. Les dups n'ont pourtant pas diminué à cette étape ; les compteurs établissent le changement du moment de destruction et du pic de payloads. Sur listes partagées ou mixtes, le gain est plus modeste mais régulier : **−2,99 à −3,67 %, 7/7**. tree_map partagé et mixte baisse de **5,38 %** et **4,21 %**, également 7/7. La mesure établit l'effet de cette version entière, sans isoler une cause microarchitecturale.

**Fusion.** Son importante baisse de dups ne produit généralement qu'un petit écart de durée : list_map unique **−0,55 %, 7/7**, list_map mixte **−0,65 %, 7/7** et list_filter unique **−0,08 %, 3/7**. Elle régresse sur tree_map à racine partagée : **7,209 → 7,457 ms (+3,44 %, 1/7 favorable)**. Sur tree_map mixte, le rapport des médianes baisse de 1,36 %, mais seulement 2/7 blocs sont favorables : aucun gain régulier n'est établi pour ce profil. Réduire les opérations de référence de l'IR ne garantit donc pas de réduire le temps du programme complet.

**Spécialisation du drop.** Supprimer les derniers dups de list_map unique ne réduit pas sa durée : **+0,08 %, 1/7 favorable** face à fusion. Les changements des listes restent proches de zéro. tree_map baisse de **2,59 % unique (5/7)**, **2,20 % partagé (6/7)** et **1,77 % mixte (7/7)**. Le cas tree_choose partagé est l'inverse du piège précédent : le compteur de dups remonte de 1 à 2, mais la durée baisse de 1,41 % avec 5/7 blocs favorables. tree_update mixte régresse de **2,70 %, 2/7 favorable**. Cette passe a donc une contribution dépendante du profil, pas un bénéfice uniforme.

**Reset/reuse.** Sur propriétaires uniques, les gains incrémentaux sont nets : list_map **12,647 → 11,338 ms (−10,36 %, 7/7)**, list_filter **8,825 → 8,627 ms (−2,25 %, 7/7)**, tree_map **7,172 → 5,448 ms (−24,04 %, 7/7)**. tree_update baisse de **5,59 % unique (6/7)** et **5,87 % mixte (5/7)**. La racine partagée donne les résultats opposés : list_map **14,702 → 15,604 ms (+6,13 %, 0/7)**, list_filter **11,118 → 11,895 ms (+6,99 %, 0/7)**, tree_map **7,293 → 8,245 ms (+13,05 %, 0/7)**. Ces régressions sont régulières sur cette campagne.

Le mécanisme du fallback est visible dans le Rust : la tentative de reset est refusée si `Rc::get_mut` échoue ; la version emprunte alors le déballage consommant et construit sans token réutilisable. Les compteurs à racine partagée confirment **zéro cellule réutilisée**, avec les mêmes allocations et dups que specialized_drop. Il y a donc un chemin supplémentaire sans économie de cellules dans ces profils. Cela explique pourquoi la baisse attendue d'allocations n'y existe pas ; les mesures seules n'attribuent cependant **pas** les +6 à +13 % à une instruction, un branchement ou un coût précis. Aucun profilage machine ni contre-essai retirant une seule opération n'a été fait.

Le mode mixte illustre la dépendance à la forme des données. Sur listes, seule la racine peut bénéficier du reset ; presque toute la queue doit suivre le chemin partagé : list_map **+6,10 %, 0/7 favorable**, list_filter **+6,17 %, 1/7**. Sur arbre, le sous-arbre droit reste unique, et les compteurs sur l'entrée de validation de 31 cellules montrent **16 cellules réutilisées** ; le résultat temporel de tree_map mixte baisse seulement de **1,56 %, 4/7**. Il est donc nettement moins convaincant que le cas entièrement unique. Enfin tree_choose n'a pas de constructeur de sortie à réutiliser et varie défavorablement de **+1,03 à +1,85 %, 2–3/7**.

**Champs conservés.** Seul tree_update est modifié dans la matrice temporelle. Face à reuse, il passe de **3,710 → 3,635 ms (−2,02 %, 4/7)** en unique, **4,009 → 4,005 ms (−0,11 %, 4/7)** en partagé et **3,793 → 3,748 ms (−1,18 %, 4/7)** en mixte. Le reset et la réinstallation du constructeur sont bien supprimés sur racine admissible, mais cette campagne établit seulement un petit écart médian, peu régulier. Les quatre autres corps de fonction sont strictement identiques à reuse : leurs différences, dont tree_map unique +1,69 % et tree_choose unique −2,79 %, ne prouvent aucun effet local de la spécialisation de champs. Elles rappellent la dispersion entre exécutions de binaires distincts.

## Portée de ces résultats

La combinaison complète passe, face à l'IR naïf synthétique, de **15,280 à 11,316 ms pour list_map unique (−25,94 %, 7/7)**, **11,577 à 8,542 ms pour list_filter unique (−26,22 %, 7/7)** et **7,691 à 5,540 ms pour tree_map unique (−27,97 %, 7/7)**. Elle régresse avec racine conservée : **+0,74 % list_map (2/7 favorables)**, **+3,28 % list_filter (0/7)** et **+9,89 % tree_map (0/7)**. Elle améliore tree_update unique de **10,23 %, 6/7**, mais la contribution propre de retained_fields est bien plus faible que cet effet cumulé.

Cette expérience prouve des transformations mécaniques, leurs bilans et leurs effets temporels sur un corpus borné. Les améliorations ne se transfèrent pas telles quelles à Purust : la baseline est un générateur persistant synthétique volontairement naïf, pas le backend actuel ni la dernière colonne native du README altbak.pub. Aucune comparaison de ces millisecondes avec les benchmarks officiels n'est pertinente sans reprendre leurs fonctions et frontières de mesure. Ce résultat ne prouve pas davantage le besoin d'une implémentation complète de Perceus ; il donne des candidats et des contre-exemples pour une future évaluation du code réellement généré.

Les calculs exacts, les sept deltas appariés et l'identité des fonctions par passage sont dans [timings-analysis.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/timings-analysis.json), produits sans réexécution par [summarize_timings.py](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/ir/summarize_timings.py). Les empreintes et résultats bruts restent conservés ; le champ `method.timings_run: false` dans timings.json est le snapshot historique de préparation du build, pas un statut de cette campagne désormais terminée. Aucun build, test ou chronométrage supplémentaire n'a été lancé pour rédiger ce bilan.
