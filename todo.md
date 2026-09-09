# Audit des FFI — corrections et baby steps

État au 8 septembre 2026. Cette liste vient d'un audit statique, sans exécution des benchmarks. Les constats des 22 points décrivent les sources au moment de cet audit ; le suivi ci-dessous indique les modifications apportées depuis. Les reproductions et validations par exécution restent à faire.

Périmètre : 224 implémentations natives des 14 benchmarks, dans les 8 langages JS, Go, PHP, Rust, Scheme, Erlang, F# et Java ; le fichier supplémentaire `src/Test/ArrayProcessingFFI.erl` ; les 8 FFI `src/Bench.*` ; les 5 FFI `srx/Test/FileOps/FileOps.*`. Soit 238 fichiers natifs, avec leurs wrappers et références PureScript, les lanceurs et le README. Les dépendances et copies générées dans `run/` ne constituent pas des sources supplémentaires à corriger séparément.

## Suivi de la passe Go — modifications non validées par exécution

Périmètre limité à Go à la demande de l'utilisateur. Aucune commande shell, compilation, mise en forme automatique, exécution de test ou de benchmark n'a été lancée pendant cette passe. Les cases cochées ici attestent uniquement d'une modification du code et de sa relecture statique.

- [x] `src/Test/StateMonadFFI.go` : 20 chaînes indépendantes à la profondeur reçue, au lieu d'interpréter 60 comme le nombre de répétitions.
- [x] `src/Test/StateMonadFFICheatcode.go` : 20 répétitions utilisant la profondeur reçue ; suppression des bornes qui ignoraient l'entrée.
- [x] `src/Test/ChurchFFICheatcode.go` : calcul de la borne `n⁵` au lieu de `n × 10000`, avec boucle native sans closures conservée.
- [x] `src/Test/AckermannFFI.go` et `AckermannFFICheatcode.go` : utilisation de l'entrée `m` dans `ack(m, 4)`.
- [x] `src/Test/AckermannFFI.purs` et `AckermannFFICheatcode.purs` : adaptation minimale des wrappers partagés pour transmettre `Bench.opaque 3`. Les noyaux non Go restent inchangés ; leur lecture lors de l'audit a établi qu'ils ignorent cet argument.
- [x] `src/Test/RBTreeFFICheatcode.go` : suppression du pool global et de l'arène locale introduite lors de la première correction. L'arbre utilise des pointeurs et une allocation Go ordinaire par nouvelle clé ; les mêmes rotations réutilisent les nœuds par mutation locale. Les insertions descendantes et le calcul de profondeur sont conservés.
- [x] `src/Bench.go` : durée depuis une origine monotone et frontière `//go:noinline` sur la restitution de l'entrée opaque ; l'effet de cette frontière dans le code généré reste à vérifier.
- [x] `bin/go/run` : arrêt explicite après un échec de gopurs, pour éviter de poursuivre avec une sortie potentiellement périmée.
- [ ] Vérifier compilation, résultats sur plusieurs entrées, appels répétés et invariants de l'arbre après les rotations lorsque les exécutions seront autorisées. Aucun de ces contrôles n'a été lancé.
- [ ] Traiter ultérieurement le protocole partagé : validation automatique des sorties, observation des résultats chronométrés et harmonisation de la chauffe. `Bench.purs` et les points d'entrée partagés n'ont pas été modifiés dans cette passe.
- [ ] Confirmer le domaine numérique et les conventions de paramètres, puis mesurer l'effet des changements Go et le comparer aux baselines historiques. Les chiffres du README n'ont pas été modifiés par cette passe.

Les autres FFI Go relues ne présentaient pas de défaut identifié dans le scénario audité et sont conservées, notamment les fusions, la scalarisation, la spécialisation et le résultat statique de RowToList. Les corrections des autres langages et les cases globales ci-dessous restent à traiter. Pour StateMonad, les wrappers partagés continuent à fournir une profondeur de 60 ; le code Go fixe les 20 répétitions pour conserver le scénario 20 × 60 sans modifier ces wrappers.

## Cadre à respecter : les colonnes du README

Les descriptions de `README.md`, section « The 99/1 philosophy and the AOT compiler vs FFI approach », définissent le rôle des colonnes et restent la référence de ce todo :

- **Compiled** : le programme PureScript compilé. Le compilateur est libre d'éliminer les abstractions et opérations intermédiaires dès lors qu'il préserve la sémantique.
- **Native FP-style FFI** : une traduction native, idiomatique et lisible des constructions fonctionnelles. Elle conserve leur rôle dans le programme source, sans devoir reproduire littéralement le code généré ni imposer un nombre identique d'allocations machine. Le compilateur natif reste libre de l'optimiser.
- **Native hand-optimized FFI** : une implémentation manuellement optimisée du même calcul. Les changements d'algorithme, formules équivalentes, fusions, mutations locales et suppressions de closures ou de structures intermédiaires y sont autorisés. Leur réduction du travail exécuté est précisément recherchée.

Le remplacement d'un million de closures par une boucle, cité explicitement dans le README, est conforme à la troisième colonne. Cette liberté vaut aussi pour le code compilé lorsqu'il sait réaliser la transformation. Les vérifications doivent préserver le sens des paramètres, le résultat et les effets observables dans le domaine du benchmark ; elles ne doivent pas forcer les versions optimisées à reconstruire les structures de la source.

Les mentions « fusion », « spécialisation » ou « autre algorithme » ci-dessous documentent les techniques à l'intérieur des colonnes existantes. Elles ne créent pas de nouvelles colonnes et ne justifient pas l'exclusion d'une implémentation optimisée correcte.

Précision demandée pour les FFI Go écrites à la main : conserver un code idiomatique, lisible et plausible pour un développeur Go, y compris dans la colonne hand-optimized. Les boucles, fusions et mutations locales restent adaptées ; RBTree utilise des allocations ordinaires et des pointeurs, sans arène ni pool d'allocation maison. Ce choix de référence humaine ne restreint pas les optimisations du compilateur et ne modifie pas les descriptions de colonnes du README.

## Méthode et ordre de travail

- Une famille et un langage par cycle : reproduire sur une petite entrée, corriger le problème précis, vérifier le résultat et le respect du rôle de la colonne, puis passer au suivant.
- Commencer par les contrats, la sélection des programmes et la validation des résultats. Faire ensuite les corrections fonctionnelles, puis celles des références et de la mesure. Renouveler les baselines en dernier.
- Une bonne sortie sur l'unique entrée officielle ne prouve pas que la fonction est correcte sur le domaine annoncé. Varier les entrées dans ce domaine ; contrôler aussi les constructions fonctionnelles dans la source des FFI ordinaires.
- Distinguer les bugs des optimisations équivalentes, y compris les changements d'algorithme. Corriger les implémentations pour respecter les colonnes du README, sans redéfinir ces colonnes.
- Ne pas modifier la référence PureScript pour lui faire accepter une FFI incorrecte. Documenter explicitement tout changement du contrat du benchmark.
- Privilégier les vérifications ciblées de moins de cinq minutes. Avant un processus plus long, prévoir un diagnostic qui montre quelle étape fonctionne ou échoue.
- Toutes les cases sont à traiter ; une exclusion doit avoir une raison explicite, telle qu'un résultat incorrect ou un backend indisponible, et non la seule disparition d'abstractions dans une variante optimisée.

## 1. Préciser les contrats dans les colonnes existantes

Fichiers : `src/Test/*.purs`, `README.md`, `bin/java/benchmark-native`.

- [ ] Pour chaque famille, relever le sens des paramètres, leur domaine, la sortie et les effets observables. Relever séparément les constructions de la source : répétitions, profondeur, ordre d'insertion, structure parcourue.
- [ ] Appliquer les rôles `Compiled`, `Native FP-style FFI` et `Native hand-optimized FFI` définis dans le README. Réserver les exigences de traduction des constructions fonctionnelles aux sources de la colonne FP-style ; vérifier l'équivalence sémantique dans toutes les colonnes.
- [ ] Établir la correspondance entre chaque wrapper et son noyau natif. Corriger d'abord les paramètres incohérents, notamment StateMonad et Ackermann.
- [ ] Définir les conventions numériques pertinentes pour la comparaison : largeur de `Int`, débordements, conversions et petites entrées autorisées. Ne pas supposer que tous les entiers négatifs ou nuls sont acceptés par chaque référence.
- [ ] Vérifier les 224 variantes dans leur colonne prévue et annoter leurs techniques d'implémentation. Marquer les implémentations incorrectes et celles qui ne sont pas encore vérifiées ; les exclure des comparaisons validées tant qu'elles ne le sont pas.
- [ ] Conserver les descriptions et libellés de colonnes du README ainsi que l'organisation `FFI` / `FFICheatcode`. Une différence entre le code et la colonne annoncée appelle une correction de l'implémentation, pas une redéfinition des colonnes.

Critère : chaque variante respecte le contrat du calcul et le rôle de sa colonne ; les techniques d'optimisation sont documentées sans modifier cette organisation.

### Sorties du scénario officiel à vérifier

Ces valeurs figurent aussi dans le validateur Java ; elles ne remplacent pas les vérifications sur plusieurs entrées. Les paramètres décrivent le scénario source, sans imposer la matérialisation de ses opérations dans les versions optimisées.

| Famille | Paramètres du scénario PureScript | Sortie attendue |
| --- | --- | --- |
| AstTree | profondeur 3 | 7 |
| Fib | 10 | 55 |
| ListOps | 900 | 202950 |
| TCO | 100000 | 100000 |
| Records | 10000 | 20000 |
| Ackermann | m = 3, n = 4 | 125 |
| Church | 10, composition n⁵ | 100000 |
| Primes | limite 500 | 21536 |
| RBTree | insertions descendantes 100000..1, puis profondeur | 22 |
| Polymorphism | 10000000 additions de 1 | 10000000 |
| StateMonad | 20 répétitions, profondeur 60 | 1200 |
| LazyEvaluation | 1000 répétitions, profondeur 1000 | 1000000 |
| ArrayOps | 900 | 202950 |
| RowToList | type fixe à 5 champs | 5 |

## 2. Corriger la sélection effective des suites

Fichiers : `bin/run`, `bin/php/run`, `bin/es/run`, les autres `bin/*/run`, `src/AppFFI.purs`, `src/AppFFICheatcode.purs`, `src/AppJavaFFI.purs`, `src/AppJavaFFICheatcode.purs`.

Constats : PHP ignore `--ffi` et conserve `App`. ES ignore `--ffi` et `--fficc`. La branche FFI du lanceur global n'appelle actuellement ni Scheme, ni Erlang, ni Java. `AppFFI` exécute actuellement les deux variantes, pas seulement la variante ordinaire.

- [ ] Établir une matrice arguments → point d'entrée → suites réellement exécutées, en commençant par PHP et ES.
- [ ] Définir sans ambiguïté si `--ffi` signifie la suite ordinaire ou les deux suites ; appliquer cette convention aux lanceurs et à leurs libellés.
- [ ] Corriger PHP, puis vérifier que `--ffi` sélectionne le programme prévu, sans lancer toute la suite globale.
- [ ] Corriger ES pour les deux modes, puis vérifier ses points d'entrée séparément.
- [ ] Vérifier la couverture globale et ajouter les backends pris en charge manquants, ou rendre leur exclusion explicite. Ne pas créer une FFI de calcul Wasm uniquement pour remplir une colonne.
- [ ] Faire apparaître le backend, le mode et le point d'entrée dans les sorties ; signaler les options non prises en charge et les échecs de façon exploitable par le validateur.

Critère : aucune commande annoncée comme FFI ne lance silencieusement la suite PureScript compilée.

## 3. Vérifier les résultats avant de chronométrer

Fichiers : `src/Bench.purs`, les points d'entrée `src/App*.purs`, `bin/java/benchmark-native`, les lanceurs.

- [ ] S'appuyer sur les 14 cas du validateur Java pour établir une vérification commune des sorties, sans recopier 8 jeux de constantes indépendants.
- [ ] Faire échouer une vérification ciblée avec un cas déjà incorrect, par exemple ArrayOps Erlang ou ListOps PHP, pour prouver que le défaut est détecté.
- [ ] Vérifier séparément les 14 variantes ordinaires et les 14 variantes optimisées de chaque langage disponible. Détecter aussi un cas manquant ou une suite interrompue.
- [ ] Ajouter de petites entrées qui révèlent les coïncidences : TCO à 2, Records à 1/2/6, Church à 2, ainsi que plusieurs tailles pour les familles concernées.
- [ ] Contrôler dans la source PureScript et les sources FP-style les dimensions et constructions annoncées, notamment pour StateMonad, LazyEvaluation et RBTree, avec un diagnostic ciblé hors mesure. Pour les variantes optimisées, vérifier leur interprétation des paramètres et leur résultat, sans exiger les mêmes opérations ou structures intermédiaires. Ne pas contrer les optimisations du compilateur natif pour forcer une charge machine identique.
- [ ] Empêcher une suite incorrecte ou incomplète de produire un total présenté comme validé.

Critère : les erreurs connues sont détectées avant les mesures ; une sortie officielle correcte à elle seule ne suffit pas à valider la variante.

## 4. Rendre les Cheatcode PHP rappelables

Fichiers : `src/Test/AstTreeFFICheatcode.php`, `FibFFICheatcode.php`, `AckermannFFICheatcode.php`, `RBTreeFFICheatcode.php`.

Constat : ces fonctions déclarent des fonctions nommées pendant leur appel. Le harness appelle chaque action trois fois en chauffe, puis dix fois en mesure, dans le même processus. Les déclarations sont également conservées dans la sortie PHP générée examinée.

- [ ] Reproduire deux appels successifs d'une seule variante et relever l'erreur de redéclaration.
- [ ] Corriger la portée des fonctions auxiliaires dans ce fichier en conservant le calcul.
- [ ] Vérifier plusieurs appels successifs, puis appliquer la même démarche aux trois autres fichiers.
- [ ] Vérifier l'absence de collisions entre ces variantes lorsqu'elles tournent dans la même suite.

Critère : les quatre variantes supportent les appels répétés sans redéclaration ni état résiduel.

## 5. AST : corriger le calcul et restaurer l'arbre dans les FFI ordinaires

Fichiers : `src/Test/AstTree.purs`, `AstTreeFFI.{js,php,fs,ss}`, `AstTreeFFICheatcode.{fs,ss}` ; PHP optimisé relève aussi du point 4.

Constats : JS/PHP ordinaires et les deux F# construisent seulement un arbre Add/Value, donnant 8 au lieu de 7. Les deux Scheme construisent une profondeur 6 et utilisent l'entrée comme nombre de répétitions, donnant 1420 à l'entrée 3.

- [ ] Dans les FFI ordinaires, comparer la forme d'un petit arbre à la référence `Val/Add/Mul/Sub`, avant de comparer sa seule valeur.
- [ ] Corriger JS ordinaire, vérifier les profondeurs 1, 2 et 3, puis traiter PHP et les deux F# séparément.
- [ ] Corriger le sens du paramètre dans chaque Scheme : il désigne la profondeur de l'expression à évaluer. La version ordinaire doit construire puis évaluer l'arbre ; la version optimisée peut utiliser une transformation équivalente sans construire cet arbre.
- [ ] Vérifier les résultats sur plusieurs petites profondeurs, puis la sortie officielle 7 ; contrôler aussi construction et évaluation dans les sources FP-style. Cette correction de résultat s'applique également au F# optimisé, sans lui imposer l'arbre de la référence.
- [ ] Documenter les optimisations de représentation Java/Erlang et le fait que Go/Rust optimisés sont essentiellement identiques aux ordinaires ; ne pas inventer une différence de performance entre deux copies.

## 6. Fibonacci et Ackermann : vérifier les résultats et aligner les entrées

Fichiers : `src/Test/Fib*`, `src/Test/Ackermann*`, `src/Bench.*`.

Constats : les Fibonacci ordinaires et optimisés restent récursifs, sans cache ni formule. Les 16 Ackermann natifs ignorent leur argument et inscrivent `ack(3,4)` dans le code ; la référence reçoit `m = 3` via `Bench.opaque`, alors que les wrappers FFI transmettent un dummy 0.

- [ ] Après le point 4, vérifier quelques petits Fibonacci dans les deux variantes et documenter leur récursion actuellement conservée. La source FP-style traduit le calcul récursif ; la version hand-optimized peut adopter un algorithme plus rapide ou une formule équivalente. Une absence actuelle d'optimisation n'est pas un bug, mais la récursion n'est pas une contrainte pour cette dernière colonne.
- [ ] Aligner le sens du paramètre Ackermann entre référence, wrappers et noyaux natifs, en gardant le scénario officiel `(3,4)`.
- [ ] Vérifier plusieurs petites valeurs de `m` à `n` fixé, puis 125 pour le scénario officiel.
- [ ] Vérifier que l'entrée opaque atteint réellement le noyau Ackermann ; distinguer une possibilité de calcul constant d'une élimination effectivement observée.
- [ ] Conserver l'élimination de récursion terminale du Cheatcode Java si son équivalence est confirmée ; autoriser également toute autre optimisation équivalente dans la colonne hand-optimized.

## 7. ListOps : réparer les résultats et assainir la référence structurelle

Fichiers : `src/Test/ListOps.purs`, `ListOpsFFI.php`, `ListOpsFFI.fs`, `ListOpsFFICheatcode.fs`, `ListOpsFFI.rs`, `ListOpsFFI.ss`.

- [ ] Montrer que `range` PHP ne construit qu'un `Cons`, puis corriger sa récursion. Vérifier une petite plage et la somme 202950 pour 900.
- [ ] Remplacer, dans les deux F#, le calcul `5 × somme(1..n)` par la somme des pairs. Vérifier une borne paire et une borne impaire.
- [ ] Examiner le `Vec` intermédiaire du Rust ordinaire et le `reverse` supplémentaire du Scheme ordinaire ; corriger le surtravail accidentel tout en conservant une traduction FP-style idiomatique, sans imposer une copie littérale du code PureScript.
- [ ] Vérifier que les sept autres Cheatcode calculent la somme pour l'entrée reçue et documenter leur fusion de parcours dans la colonne hand-optimized. Une formule équivalente supprimant également la boucle y est autorisée.

Critère : même somme sur la plage demandée ; la source FP-style conserve le traitement fonctionnel des listes et la version hand-optimized peut supprimer les allocations et parcours.

## 8. ArrayOps : supprimer le stub et résoudre le doublon Erlang

Fichiers : `src/Test/ArrayOpsFFI.erl`, `ArrayProcessingFFI.erl`, `ArrayOpsFFI.fs`, `ArrayOpsFFICheatcode.fs`, `ArrayOpsFFI.ss`, les wrappers ArrayOps.

- [ ] Vérifier quel fichier Erlang est effectivement associé à `Test.ArrayOpsFFI`. Le fichier homonyme retourne 0 ; `ArrayProcessingFFI.erl` contient un calcul mais déclare le même module.
- [ ] Garder une seule implémentation source correctement nommée et raccordée ; vérifier le module chargé avant de retirer le doublon inutile.
- [ ] Vérifier une petite plage et la sortie officielle 202950 en Erlang.
- [ ] Corriger séparément les deux F# qui calculent 2027250 à l'entrée 900 ; vérifier bornes paires et impaires.
- [ ] Aligner le Scheme ordinaire déjà fusionné sur le rôle FP-style en rétablissant un traitement fonctionnel des tableaux dans sa source. Conserver la fusion dans la colonne hand-optimized.
- [ ] Conserver comme optimisations équivalentes les fusions des autres Cheatcode, après vérification de leur domaine.

## 9. TCO : supprimer la coïncidence `+1`

Fichiers : `src/Test/TCOFFI.{erl,fs}`, `TCOFFICheatcode.{erl,fs}`, référence `TCO.purs`.

- [ ] Vérifier le contre-exemple `n = 2` : la référence donne 3, les quatre variantes donnent 2.
- [ ] Restaurer l'accumulation de `n mod 3` dans les sources FP-style. Corriger les variantes hand-optimized pour calculer la même somme ; une formule équivalente peut y remplacer la boucle et les modulos.
- [ ] Vérifier plusieurs tailles non multiples de 3, puis le scénario officiel à 100000.
- [ ] Documenter les variantes ordinaires/optimisées identiques et les transformations TCO réellement présentes.

Critère : la fonction est correcte sur les petites entrées, pas seulement sur le total officiel qui masquait l'erreur.

## 10. Records : corriger F# et distinguer reconstruction et scalarisation

Fichiers : `src/Test/RecordsFFI.fs`, `RecordsFFICheatcode.fs`, `RecordsFFI.ss`, référence `Records.purs` et autres variantes Records.

- [ ] Reproduire `n = 1`, `2` et `6`, qui révèlent la substitution incorrecte du modulo par `+2` dans les deux F#.
- [ ] Corriger le calcul F# ; conserver les records imbriqués dans la source FP-style, tout en autorisant scalarisation ou formule équivalente dans la variante hand-optimized.
- [ ] Vérifier la sortie 20000 à 10000 après les contre-exemples, et non comme unique preuve.
- [ ] Rétablir une représentation idiomatique des records imbriqués dans la source Scheme FP-style. La scalarisation manuelle reste autorisée dans la colonne hand-optimized et le compilateur natif peut optimiser les deux sources.
- [ ] Documenter les structs par valeur, la mutation locale et la scalarisation comme optimisations possibles lorsque les anciennes versions ne s'échappent pas. Ne pas réintroduire artificiellement des allocations dans une variante optimisée.

## 11. Church : corriger la fonction et expliciter la disparition des closures

Fichiers : `src/Test/ChurchFFI.{fs,ss}`, `ChurchFFICheatcode.{js,go,php,ss,fs}`, les autres Church et référence `Church.purs`.

- [ ] Vérifier à `n = 2` que la référence vaut 32 ; les compteurs `n × 10000` ne sont pas des remplacements corrects de `n⁵`.
- [ ] Corriger les deux F# qui ne font que `n` incréments et retournent 10 au scénario officiel.
- [ ] Corriger le calcul du Scheme ordinaire et des Cheatcode JS/Go/PHP/Scheme, en gardant des paramètres effectivement utilisés.
- [ ] Conserver les fonctions Church dans les sources FP-style ; les rétablir dans le Scheme ordinaire, qui n'est actuellement pas un témoin de ces constructions.
- [ ] Documenter dans la colonne hand-optimized les Cheatcode qui comptent jusqu'à `n⁵` et l'Erlang qui calcule directement le produit. Les deux approches y sont admises : supprimer les applications Church est conforme au README, et aucune boucle d'incréments n'est obligatoire pour une formule équivalente.
- [ ] Préserver les optimisations valides de partage en Erlang et d'applications partielles en Rust ordinaire, en documentant leur effet sur les allocations.
- [ ] Vérifier plusieurs petites entrées dans le domaine convenu, puis 100000 à l'entrée 10 ; contrôler les débordements si ce domaine est élargi.

## 12. Primes : réparer PHP et documenter les algorithmes dans leur colonne

Fichiers : `src/Test/PrimesFFI.php`, `PrimesFFI.erl`, `PrimesFFI.fs`, tous les `PrimesFFICheatcode.*`, référence `Primes.purs`.

- [ ] Reproduire la plage PHP réduite à `[limit]`, puis corriger sa construction avant le crible.
- [ ] Vérifier une petite borne composée et une borne première, puis la somme 21536 jusqu'à 500.
- [ ] Aligner la source Erlang ordinaire sur le crible fonctionnel par filtrages attendu dans la colonne FP-style. Les essais de division restent autorisés dans la colonne hand-optimized.
- [ ] Documenter le filtre direct du F# ordinaire, qui évite l'inversion de la référence mais change allocations et profondeur de pile.
- [ ] Documenter les six Cheatcode par marquage des multiples et les variantes par essais de division Erlang/F# ; vérifier qu'ils calculent bien les premiers pour plusieurs limites. Ces algorithmes restent tous admissibles dans la même colonne hand-optimized.

Critère : la colonne FP-style traduit le crible fonctionnel et la colonne hand-optimized peut employer un meilleur algorithme. L'écart entre les colonnes reste un objectif légitime d'amélioration du compilateur, conformément au README ; aucune réponse erronée n'entre dans les mesures.

## 13. RBTree : corriger la profondeur retournée et le scénario FP-style

Fichiers : `src/Test/RBTreeFFI.{js,php,fs}`, `RBTreeFFICheatcode.{erl,fs}`, référence `RBTree.purs` ; PHP optimisé relève aussi du point 4.

Constats : JS/PHP ordinaires insèrent `0..n-1`, ajoutent `n` recherches et somment les valeurs ; F# ordinaire compte les nœuds après insertions ascendantes. Les Cheatcode Erlang/F# utilisent `gb_trees` / `SortedSet` et retournent leur taille, pas la profondeur de l'arbre de référence.

- [ ] Dans les sources FP-style, vérifier sur un petit arbre l'ordre des insertions, les rotations et la nature de la sortie avant de traiter 100000 éléments.
- [ ] Corriger JS puis PHP ordinaires : insertions `n..1`, puis parcours calculant la profondeur ; retirer les recherches propres à l'autre scénario.
- [ ] Corriger F# ordinaire pour le même contrat.
- [ ] Corriger les Cheatcode Erlang/F# qui retournent une taille au lieu de la profondeur demandée. Leur changement de structure est autorisé s'il permet de calculer le même résultat ; il ne justifie pas de retourner une autre valeur. Une méthode équivalente sans arbre matérialisé est également admissible.
- [ ] Vérifier l'ordre des clés et les invariants rouge-noir dans les implémentations qui utilisent cette structure. Pour toutes les variantes, comparer le résultat à la profondeur de référence sur de petites entrées, puis vérifier 22 à 100000, sans imposer les mêmes rotations ou un parcours final aux versions optimisées.
- [ ] Si une implémentation n'est valable que pour les insertions descendantes du scénario, documenter ce périmètre plutôt que de la présenter comme un arbre général validé.

## 14. RBTree : corriger les copies excessives et simplifier l'allocation Go

Fichiers : `src/Test/RBTreeFFI.rs`, `RBTreeFFICheatcode.go`, `RBTreeFFICheatcode.rs`, `RBTreeFFICheatcode.java`.

- [ ] Sur une petite entrée Rust ordinaire, compter ou tracer les copies de sous-arbres provoquées par les `.clone()` de `Option<Box<Node>>` dans les rotations.
- [ ] Choisir une représentation qui conserve le contrat et le partage attendu sans copie profonde supplémentaire ; vérifier les rotations et la sortie avant toute mesure.
- [x] Pour Go optimisé, remplacer le pool puis l'arène par des allocations ordinaires de nœuds, des pointeurs et des rotations en place. La relecture statique est faite ; la validation par exécution reste à faire.
- [ ] Pour Go, vérifier chaque cas de rotation sur une petite entrée : ordre des clés, couleurs et profondeur comparée à la référence ; vérifier ensuite les insertions descendantes du scénario officiel.
- [ ] Pour Go, vérifier les appels répétés avec des tailles différentes ; confirmer que chaque appel construit son propre arbre et inclut ses allocations dans le calcul chronométré.
- [ ] Lors de la passe Rust, relever le nombre de nœuds réellement nécessaires face à la capacité réservée de 5000000, puis ajuster ou justifier cette réserve ; préciser le coût d'allocation et d'initialisation inclus dans la mesure.
- [ ] Conserver la mutation locale Go et Java tant que les anciennes versions de l'arbre ne sont pas observables. L'examen des choix d'allocation des autres langages reste reporté à leurs passes respectives.

Critère : les témoins ne sont pas ralentis par des copies accidentelles ; la version Go optimisée utilise une gestion mémoire ordinaire et maintenable, sans arène ni pool maison. Les coûts d'allocation des variantes sont inclus et documentés dans la mesure.

## 15. Polymorphism : corriger Erlang et respecter le rôle des deux FFI

Fichiers : `src/Test/PolymorphismFFI.erl`, `PolymorphismFFICheatcode.erl`, `PolymorphismFFI.{ss,fs}`, les autres variantes Polymorphism.

- [ ] Reproduire sur une petite entrée que les deux Erlang additionnent l'aire d'un cercle au lieu d'additionner 1 via `Monoidish Int`.
- [ ] Corriger les deux calculs avec les mêmes paramètres numériques que la référence ; vérifier plusieurs petits nombres d'itérations, puis 10000000.
- [ ] Restaurer une traduction idiomatique du dictionnaire dans les sources FP-style Scheme/F#, qui font actuellement de simples incréments. Conserver la spécialisation manuelle dans la colonne hand-optimized ; le compilateur natif reste libre de dévirtualiser les sources FP-style.
- [ ] Retirer les restes inutiles de l'ancien scénario, notamment le dictionnaire `Show` inutilisé en F#, lors de la correction du fichier concerné.
- [ ] Documenter les interfaces et appels non curryfiés JS/Go/PHP/Rust/Java comme optimisations de représentation possibles. Ne pas empêcher artificiellement une dévirtualisation autorisée par les types connus.

## 16. StateMonad : unifier paramètres, répétitions et profondeur

Fichiers : `src/Test/StateMonad.purs`, `StateMonadFFI.purs`, `StateMonadFFICheatcode.purs`, toutes les variantes natives StateMonad.

- [ ] Fixer le sens de l'entrée commune : nombre de répétitions ou profondeur. Le scénario source reste 20 chaînes indépendantes de profondeur 60 ; les versions optimisées doivent en calculer le résultat sans devoir matérialiser ces chaînes.
- [ ] Adapter tous les noyaux au contrat retenu, y compris ceux qui donnent déjà 1200. Si l'entrée devient le nombre de répétitions, changer seulement le wrapper de 60 à 20 casserait notamment Java ordinaire et les Cheatcode JS/PHP/Rust/Erlang/Java qui l'interprètent actuellement comme une profondeur.
- [ ] Avec de petites dimensions distinctes, observer le nombre de chaînes et leur profondeur dans les sources FP-style. Dans les versions optimisées, vérifier la fonction des paramètres et le résultat, sans imposer cette répartition à l'exécution.
- [ ] Corriger les ordinaires JS/Go/Rust/Scheme/F# qui font actuellement 60 × 60, puis PHP qui ne fait qu'une chaîne de 60.
- [ ] Corriger l'ordinaire Erlang qui fait une seule chaîne de 1200 ; le Java ordinaire actuel constitue un témoin utile du scénario 20 × 60.
- [ ] Corriger les Cheatcode Scheme/F# qui font 3600 incréments ; rendre le Go optimisé dépendant de l'entrée au lieu de deux bornes constantes 60 et 20.
- [ ] Vérifier que l'usage du dictionnaire de processus dans le Cheatcode Erlang ne laisse pas d'effet observable ni n'écrase un état préexistant pour une FFI déclarée pure ; le remplacer par un état local si nécessaire.
- [ ] Vérifier les paramètres et les résultats sur de petites dimensions, puis 1200 dans toutes les variantes ; contrôler aussi la répartition des chaînes dans les sources FP-style. Documenter les substitutions par boucles ou formules dans la colonne hand-optimized sans exiger 1200 incréments effectifs.

## 17. LazyEvaluation : corriger le résultat et les chaînes des FFI ordinaires

Fichiers : `src/Test/LazyEvaluationFFI.{erl,fs}`, `LazyEvaluationFFICheatcode.{erl,fs}`, les autres variantes LazyEvaluation et référence `LazyEvaluation.purs`.

- [ ] Fixer séparément nombre de répétitions et profondeur ; la source du scénario officiel reconstruit et force 1000 chaînes de profondeur 1000. Ces dimensions définissent le résultat attendu, sans imposer la construction de thunks dans la colonne hand-optimized.
- [ ] Reproduire sur de petites dimensions le calcul différent d'Erlang (somme d'une liste paresseuse) et la simple boucle de F#.
- [ ] Corriger les deux FFI ordinaires pour construire et forcer les chaînes prévues ; vérifier une profondeur supérieure à 1 et plusieurs répétitions.
- [ ] Corriger la valeur et les paramètres des deux Cheatcode Erlang/F#, puis vérifier 1000000 au scénario officiel.
- [ ] Documenter les autres Cheatcode dans leur colonne actuelle : un million d'incréments en JS/PHP/Java, mille additions de 1000 en Go/Rust/Scheme. Ces suppressions de thunks sont explicitement voulues par le README ; une multiplication équivalente supprimant aussi les boucles est autorisée. Leur écart avec la version fonctionnelle sert précisément à mesurer le potentiel d'optimisation.
- [ ] Documenter les différences de représentation des FFI ordinaires, dont les wrappers de closure supplémentaires en Scheme, avant d'interpréter leurs écarts.

## 18. RowToList : conserver la constante légitime et corriger les libellés

Fichiers : `src/Test/RowToList.purs`, tous les `RowToListFFI.*` et `RowToListFFICheatcode.*`.

- [ ] Consigner que `keys` ignore les valeurs du record et que son type fixe possède cinq champs ; le paramètre opaque n'est pas une taille dynamique dans la référence.
- [ ] Reconnaître les retours directs de 5 comme une spécialisation statique légitime. Dans la colonne FP-style, vérifier si la constante découle du compilateur natif appliqué à une traduction fonctionnelle, ou remplace déjà cette traduction dans la source, comme en Rust/Scheme/F# ; traiter cette dernière différence selon le rôle de la colonne, sans qualifier la valeur de triche.
- [ ] Documenter les ordinaires JS/Go/PHP/Java comme parcours de dictionnaires et l'Erlang comme énumération dynamique d'une map. Aligner les sources FP-style sur une traduction idiomatique de la résolution par le type, sans imposer un parcours après compilation.
- [ ] Conserver les constantes manuelles dans la colonne hand-optimized et toute réduction statique produite par les compilateurs. Ne pas inventer une dépendance à l'entrée opaque que la référence ignore déjà.
- [ ] Vérifier 5 dans toutes les variantes ; les vérifications ne doivent pas empêcher la réduction à cette constante ni forcer des allocations dans les versions optimisées.

## 19. Protéger les calculs mesurés contre leur élimination

Fichiers : `src/Bench.purs`, les 8 `src/Bench.*`, les wrappers, `src/Bench.java` et `bin/psgo/run` pour sa FFI Bench embarquée.

Constats : `runBench` ignore les résultats des dix appels mesurés. Les implémentations d'`opaque` sont hétérogènes : plusieurs sont de simples identités encapsulées ; Rust utilise `black_box` à l'entrée ; les noyaux Java ont une barrière volatile d'entrée et de résultat. Le fait qu'une élimination soit possible ne prouve pas qu'elle a eu lieu.

- [ ] Sur un seul petit noyau par backend, examiner le code généré et sa sortie native optimisée pour établir si les entrées prévues restent opaques et si le résultat chronométré reste observable. Le calcul peut légitimement subsister sous une forme simplifiée.
- [ ] Définir un mécanisme d'observation des résultats qui fonctionne réellement dans chaque backend ; l'appliquer de manière comparable au compilé et aux FFI.
- [ ] Vérifier les résultats mesurés sans inclure des logs ou des comparaisons coûteuses dans chaque intervalle ; expliciter le coût inclus de la barrière.
- [ ] Traiter Ackermann après alignement de ses paramètres et StateMonad Go après suppression des bornes fixes ; ne pas protéger seulement un dummy inutilisé.
- [ ] Harmoniser aussi la FFI Bench embarquée du lanceur psgo si ce backend reste dans les comparaisons.
- [ ] Distinguer la suppression d'un calcul parce que son résultat n'est pas observé d'une simplification équivalente dont le résultat reste utilisé. Autoriser les réductions de boucles et d'abstractions dans le compilé comme dans le hand-optimized, notamment RowToList constant ; les microtemps correspondants doivent être interprétés comme tels.

## 20. Harmoniser horloges, chauffe et statistiques

Fichiers : `src/Bench.*`, `src/App.purs`, `src/AppFFI.purs`, `src/AppFFICheatcode.purs`, `src/AppJavaFFI.purs`, `src/AppJavaFFICheatcode.purs`, `bin/java/benchmark-native`.

- [ ] Relever pour chaque horloge sa monotonie, son unité et sa résolution : Go/PHP/Erlang/Rust utilisent actuellement l'heure civile ; Rust tronque à la microseconde entière. Vérifier aussi Scheme et le fallback JS.
- [ ] Adopter une horloge adaptée aux durées dans chaque backend ; vérifier une courte mesure ciblée avant de comparer les noyaux.
- [ ] Pour les opérations proches du coût de mesure, déterminer un nombre de répétitions permettant une durée lisible et documenter le coût fixe. Ne pas attribuer une précision que l'horloge ne fournit pas.
- [ ] Aligner ou justifier la chauffe : `App` fait trois passages globaux, les points d'entrée FFI génériques ne les font pas ; Java possède ses points d'entrée dédiés.
- [ ] Vérifier que les appels de chauffe recalculent bien la même tâche sans état résiduel ni coût déplacé hors mesure.
- [ ] Définir la statistique publiée et la variabilité à conserver ; distinguer le minimum de dix échantillons de la médiane de plusieurs processus frais utilisée en Java.
- [ ] Corriger l'affirmation de `Bench.purs` selon laquelle toutes les perturbations ne peuvent qu'ajouter du temps et le minimum représente nécessairement le coût le plus exact.

## 21. I/O étendues : contrôler les erreurs et comparer les mêmes opérations

Fichiers : `srx/Test/FileOps/FileOps.{purs,js,go,php,erl,ss}`, `srx/AppX.purs`, les lanceurs `--x`.

Constats : les cinq FFI font de vrais accès aux fichiers. PHP ignore les échecs de lecture/écriture. Scheme supprime puis recrée le fichier. Le point d'entrée étendu actuel importe `Test.BenchCheck`, absent des sources examinées.

- [ ] Vérifier le raccordement effectif de `--x` et son point d'entrée avant d'utiliser les résultats étendus du README ; corriger l'import absent ou la sélection du scénario.
- [ ] Définir le contrat de succès/échec des primitives et vérifier un aller-retour sur un petit fichier temporaire ainsi qu'une erreur d'accès.
- [ ] Corriger PHP pour qu'un échec d'écriture ou de lecture ne soit pas compté comme une opération réussie et ne retourne pas `false` à la place du `String` annoncé.
- [ ] Comparer la suppression/recréation Scheme à l'ouverture/troncature des autres langages ; aligner les opérations si la comparaison l'exige, sinon documenter cette différence.
- [ ] Vérifier le contenu lu et le nombre de répétitions hors chronométrage, puis le scénario 10000 écritures/lectures.
- [ ] Décrire la portée réelle de la mesure : appels synchrones au système de fichiers, sans garantie de persistance physique faute de synchronisation explicite. Documenter l'état du cache lors des mesures, sans prétendre que les variantes évitent les accès.
- [ ] Vérifier les types FFI des chaînes dans chaque backend effectivement activé ; ne pas publier de ligne pour une variante non raccordée ou non validée.

## 22. Mettre à jour les descriptions puis renouveler les baselines

Fichiers : `README.md`, commentaires et `describe` dans `src/Test/*`, scripts de mesure.

- [ ] Corriger les descriptions périmées : Primes mentionne encore 20000 dans ses commentaires alors que la limite est 500 ; Records parle d'un million de mises à jour au lieu de 10000 ; LazyEvaluation contient encore une profondeur commentée de 4000 alors que le scénario utilise 1000.
- [ ] Revoir les libellés qui annoncent des nombres exacts de closures ou de lookups : distinguer les opérations source de ce que le compilateur élimine ou alloue réellement.
- [ ] Conserver les libellés et descriptions `Compiled`, `Native FP-style FFI` et `Native hand-optimized FFI` du README. Ajouter si utile des annotations par ligne sur les techniques employées ; elles ne remplacent pas ces colonnes et ne leur imposent pas la même charge d'exécution.
- [ ] Conserver les chiffres historiques officiels du README avec leur provenance et marquer les comparaisons affectées comme non validées pour les sources actuelles. Ne pas remplacer ces baselines par un run isolé.
- [ ] Après les corrections et validations ciblées, mesurer une famille à la fois dans des conditions documentées, puis seulement les suites complètes.
- [ ] Conserver les sorties, versions des runtimes/backends, révisions des sources, options de compilation/JIT, paramètres, méthode de chauffe, statistique et variabilité des nouvelles mesures.
- [ ] Comparer les nouvelles mesures aux baselines historiques en explicitant les corrections de calcul, les optimisations et les changements de protocole. Écarter les gains provenant d'un résultat faux ou de paramètres altérés ; reconnaître comme gains légitimes les suppressions d'opérations et changements d'algorithme qui préservent la sémantique.
- [ ] Recalculer les totaux avec les lignes validées selon le rôle de chaque colonne ; expliciter les variantes absentes et documenter les algorithmes employés. Une différence de charge due à une optimisation correcte ne rend pas les colonnes incomparables au regard de l'objectif annoncé.
- [ ] Mettre à jour les constats chiffrés et les affirmations de performance affectées par les nouvelles mesures, en conservant la philosophie 99/1 et l'objectif de rapprocher le compilé du meilleur natif optimisé, voire de le dépasser. Fonder ces constats sur les résultats validés ; l'audit ne redéfinit pas l'objectif du projet.

## Critères de clôture

- [ ] Chaque implémentation auditée est corrigée et vérifiée dans le rôle de sa colonne, ou explicitement exclue avec une raison. Les descriptions de colonnes du README sont conservées.
- [ ] Chaque lanceur exécute la suite annoncée et les erreurs de résultat sont détectées avant mesure.
- [ ] Les contre-exemples TCO, Records et Church passent et toutes les variantes calculent le bon résultat pour les paramètres convenus. Les sources FP-style traduisent les chaînes StateMonad/LazyEvaluation et la structure RBTree ; les variantes optimisées peuvent les supprimer ou les remplacer par un calcul équivalent.
- [ ] Les suites supportent les appels répétés et les calculs chronométrés restent observables.
- [ ] Toutes les optimisations équivalentes restent autorisées dans le compilé et le hand-optimized : RowToList constant, fusion des parcours, TCO, spécialisation fondée sur les types, mutation locale, formules et changements d'algorithme. Aucun contrôle n'impose de conserver les opérations source après optimisation.
- [ ] Les FFI Go écrites à la main restent idiomatiques et maintenables ; RBTree utilise des allocations ordinaires sans arène ni pool maison, conformément au choix de référence demandé.
- [ ] Les baselines et conclusions sont mises à jour à partir de mesures reproductibles, comparées à l'historique officiel du README.
