# Correction des benchmarks natifs — 14–15 septembre 2026

Ce dossier conserve les preuves de correction des 14 tests de calcul, les exécutions complètes et la campagne de mesure. `previous-readme.md` est le README avant cette campagne. Les anciennes valeurs ne doivent pas servir à calculer un « gain » après correction : plusieurs colonnes ne calculaient pas le résultat annoncé, certaines utilisaient un autre algorithme ou une autre charge, et les protocoles de compilation et de chauffe différaient.

## Contrat contrôlé

Les 14 résultats attendus sont, dans l'ordre : `7, 55, 202950, 100000, 20000, 125, 100000, 21536, 22, 10000000, 1200, 1000000, 202950, 5`.

Les tests natifs utilisent aussi de petites entrées, dont plusieurs valeurs négatives pour `Array.range`. Ils distinguent notamment `Ackermann(m,4)`, `Church(n)=n^5`, 20 chaînes State de profondeur donnée, un nouvel ensemble de closures Lazy à chaque répétition et la profondeur de l'arbre construit par insertions descendantes. Un simple résultat nominal identique ne suffisait pas : le calendrier d'application des closures Church, l'ordre des cellules filtrées et le partage persistant de l'arbre Rust ont aussi des contrôles structurels.

## Corrections retenues

- **JavaScript** : alignement des AST, de la profondeur d'arbre, du State, de Church et de l'argument d'Ackermann ; parcours descendant des tableaux.
- **Scheme** : même contrat de calcul et de paramètres ; vraies closures Church, vecteurs, records imbriqués et dictionnaires dans la colonne FP ; horloge monotone. Le runtime privé de chaque build utilise les symboles ICU disponibles, vérifiés avec des conversions de chaînes ; aucune bibliothèque système n'a été remplacée.
- **Erlang** : suppression du résultat factice de Array, correction des thunks Lazy et des additions polymorphes, profondeur d'arbre, boucles State, paramètres et charge ; correction du module foreign en collision. Les dépendances implicites du code généré sont incluses dans le build.
- **PHP** : AST, vraies listes et crible, profondeur d'arbre, 20 chaînes State ; fonctions locales récursives évitant les redéclarations au second appel ; Church et Ackermann paramétrés. La répétition Lazy ne retient plus les anciennes chaînes dans des frames récursives, mais construit et force toujours les closures requises.
- **F#** : résultats et charges des AST, listes/tableaux, Church, State, Lazy et arbre ; paramètres Records, TCO et Ackermann. Les versions FP conservent les structures et closures fonctionnelles correspondantes.
- **Java** : Ackermann utilise son argument ; le runner sélectionne directement les points d'entrée communs AppFFI et AppFFICheatcode, avec collecte de 14 résultats par mode. Les anciennes entrées Java spécifiques devenues inutilisées sont supprimées.
- **Rust** : State exécute 20 chaînes ; Ackermann utilise son argument ; arbre FP persistant avec partage `Rc` au lieu de clonages récursifs de sous-arbres `Box` ; calendrier Church et ordre du filtrage List ; tableaux descendants. Horloge monotone en microsecondes fractionnaires.
- **Go** : tableaux descendants ; même `GOGC` pour exécution ordinaire et PGO, avec `-pgo=off` explicitement pour cette campagne.
- **C++** : les corrections précédentes de l'unité de temps, des calculs, des flags `-O3` et de l'absence de stubs restent en place. Le parcours descendant de Array a été ajouté, et les points d'entrée natifs délèguent la chauffe au code partagé pour éviter de la doubler.

Les backends compilateurs voisins n'ont pas été modifiés pour obtenir ces résultats.

## Protection des runners

`bin/benchmark/validate.py` rejette un résultat erroné, une mauvaise étiquette ou un mauvais mode, une suite partielle, une ligne dupliquée, une durée négative ou non finie, ou un total incompatible avec la somme des lignes. Le runner C++ conserve également son validateur propre. `bin/run` termine en échec si l'un des backends échoue, au lieu de masquer cet échec derrière un message.

Les builds sont isolés par mode et disposent d'artefacts identifiés. Les modes natifs exécutent la même chauffe globale que le compilé : trois suites complètes, puis trois chauffes par test et dix exécutions chronométrées. Les options non prises en charge échouent explicitement. `--test` n'écrase plus le fichier AppX du projet. Les builds psgo et Wasm utilisent le CoreFn standard attendu par ces backends ; Go/gopurs et Rust/purust conservent leur pipeline TAST.

## Preuves d'exécution

- 218 valeurs contrôlées pour chacun de JavaScript, Scheme, Erlang, Go, Rust, PHP, F# et Java : **1 744 assertions numériques**. PHP vérifie également les appels répétés qui déclenchaient auparavant des redéclarations.
- **184 assertions numériques C++**, plus horloge/opaque et invariants de l'arbre manuscrit ; Rust vérifie aussi la persistance, les observateurs Weak, le partage Rc, le filtrage et les applications Church.
- Les tests d'horloges vérifient l'unité sur une attente réelle et la monotonie. Le test psgo exécute le vrai FFI lié au programme, pas une reproduction approximative du timer.
- **9 tests** du validateur commun, **3 tests** du lanceur global, **6 tests** du runner C++ et **5 tests** d'orchestration Go/Rust passent. Le test du lanceur global a aussi révélé et permis de corriger son invocation sans argument sous Bash macOS.
- Les **30 colonnes** ont passé une exécution complète après construction. Les reconstructions requises après suppression des entrées Java inutilisées ont été réalisées ; les manifestes n'ont pas été réécrits pour faire accepter un ancien build.

Le contrôle des sorties compare la valeur publiée pendant la chauffe, puis les étiquettes, le nombre de tests, les durées et le total. Les dix résultats individuels chronométrés restent consommés comme dans le protocole existant ; ils ne sont pas tous imprimés ni comparés séparément. Les tests natifs hors chronométrage complètent ce contrôle sur plusieurs entrées et appels.

## Méthode des nouvelles mesures

La campagne utilise des artefacts déjà compilés, sans compiler simultanément. Chaque colonne est exécutée dans trois processus indépendants et les modes tournent dans un ordre différent à chaque répétition. Pour chaque test, le README retient la médiane des trois minima de dix exécutions. Le total est la somme de ces 14 médianes, pas le minimum d'un total choisi séparément. Les sorties brutes, les valeurs validées et les commandes sont conservées dans `measure-*.log`, `measure-*.json`, `measurement-plan.json` et `results.json`.

Rust utilise explicitement `opt-level=3`, `debug=false` dans les trois colonnes ; le précédent runner héritait de `opt-level=1`, `debug=true` du Cargo généré. C'est un changement de profil de mesure, pas un progrès de purust. Go et psgo utilisent `GOGC=800`, sans PGO dans cette campagne. C++ utilise `-std=c++11 -O3 -DNDEBUG`, sans LTO. PHP utilise OPcache CLI et JIT 1255, avec un buffer de 128 Mio ; son activation réelle est enregistrée dans `php-jit-profile.json`. Les trois colonnes Java utilisent `-Xss100M`. F# utilise .NET 8 en Release avec optimisation. Les autres options et versions exactes sont conservées dans les manifestes, `environment.json` et `php-fsharp-java-native-versions.json`.

## Limites

Le périmètre validé ici est la suite de **14 calculs core**. Le chemin PGO Go est couvert par les contrôles d'orchestration, mais aucune campagne PGO réelle n'a été exécutée ; les 30 colonnes de cette campagne utilisent les profils indiqués, sans PGO. Le validateur du mode étendu `--x` vérifie sa structure et ses durées ; il ne prétend pas certifier les résultats des effets, notamment le parallélisme et ses différences de largeur d'entiers. Les tests natifs vérifient le domaine fini utilisé dans ces benchmarks et des entrées voisines, pas tous les débordements possibles de `Int`.

La colonne FP reste une traduction manuelle avec structures fonctionnelles, pas une identité instruction par instruction. La colonne manuscrite autorise volontairement les simplifications algorithmiques qui respectent le résultat et l'argument. Les temps proches de la résolution du timer et les calculs fortement simplifiés par le compilateur doivent être interprétés comme tels ; une valeur affichée `0.00 µs` ne prouve pas une absence de coût. La frontière Effect `opaque` de JavaScript n'offre pas une garantie d'interdiction d'inlining par V8.

Un minimum de dix exécutions réduit certaines perturbations, mais ne fournit ni intervalle de confiance ni garantie d'absence de bruit système. Les trois répétitions et leurs totaux sont conservés pour rendre cette variabilité visible.

## Totaux de la campagne validée

| Cible | Compilé | FP natif | Manuscrit natif |
|---|---:|---:|---:|
| JavaScript | 129.10 ms | 49.79 ms | 31.38 ms |
| Go | 44.73 ms | 98.04 ms | 11.72 ms |
| Scheme | 45.78 ms | 46.51 ms | 19.20 ms |
| Erlang | 90.65 ms | 74.90 ms | 15.34 ms |
| PHP | 125.41 ms | 1081.66 ms | 211.11 ms |
| Rust | 9.78 ms | 61.35 ms | 31.89 ms |
| C++ | 946.32 ms | 134.54 ms | 21.11 ms |
| F#/C# | 346.59 ms | 75.62 ms | 13.41 ms |
| Java | 18.03 ms | 18.12 ms | 9.19 ms |

Autres colonnes compilées : Arista ES **70.31 ms**, psgo **1514.83 ms**, Wasm GC **579.83 ms**.

## Comparaison avec les publications antérieures

Le précédent README publiait 24,02 ms pour Go compilé et 60,41 ms pour F# compilé ; cette campagne donne respectivement 44,73 ms et 346,59 ms. Ce sont des différences historiques importantes, pas des régressions attribuées par cette seule campagne à une modification précise du compilateur. Les sources et profils nouvellement mesurés sont conservés pour permettre cette attribution séparément.

Rust compilé passe de la publication à 12,56 ms à 9,78 ms ici, avec le changement explicite O1→O3 ; cette différence ne prouve pas un gain de purust. Les valeurs C++ déjà corrigées avant cette campagne restent du même ordre : 951,00→946,32 ms pour le compilé, 133,23→134,54 ms pour le FP et 21,33→21,11 ms pour le manuscrit.

L'ancienne conclusion générale sur Go doit être retirée : dans cette campagne, Go compilé (44,73 ms) est plus lent que cette implémentation manuscrite (11,72 ms). Rust compilé (9,78 ms) est plus rapide que les deux implémentations natives présentées (61,35 et 31,89 ms). Ces constats portent sur les programmes mesurés, sans établir une supériorité universelle d'un langage ou d'un compilateur.

Le contrôle final `verify-publication.py` relit les 90 sorties brutes, revalide les 1 260 valeurs, recalcule les médianes et compare les 420 temps et les 30 totaux publiés. Les empreintes des sources du benchmark sont inchangées depuis leur gel avant mesure. Résultat : `publication-validation.log`.
