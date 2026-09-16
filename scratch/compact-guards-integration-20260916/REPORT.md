# Intégration des décisions d'équilibrage — 16 septembre 2026

L'optimisation générique est intégrée dans `purust/purust/src/Purust/ChildBranchPrinter.purs`, et le bundle `bin/purust.js` est recompilé. Deux campagnes sur le runner complet mesurent **−2,96 % puis −4,03 % sur RBTree**, et **−2,73 % puis −3,65 % sur le temps total**. Aucun changement Haskell/PBO ni suppression de contrôle d'unicité n'est nécessaire à ce gain.

## Changement retenu

Le générateur valide d'abord les chemins et constructeurs du prédicat initial. Il construit ensuite un arbre borné de décisions, simplifie les tests déjà connus et emprunte les champs dans les branches où leur constructeur est établi. Une continuation répétée est émise une seule fois lorsqu'elle peut être évaluée depuis les arguments racines, sans déplacer un emprunt hors de sa portée. Un bloc Rust étiqueté exprime ce partage sans allocation ni closure.

La limite porte sur 128 décisions au total. Les prédicats non pris en charge conservent l'ancien rendu. Les constructeurs qualifiés et non qualifiés sont comparés par leur identité native validée avant toute déduction d'exclusivité.

Sur le Rust RBTree réellement généré, `balance__purust_child_rebuilds` passe de 11 593 caractères (57 tests `matches!` et 70 projections imbriquées par `match`) à 2 608 caractères et 17 `match`, sans projection `unreachable!`. Ces nombres décrivent le source, pas le nombre d'instructions exécutées. Les **103 occurrences de `Rc::get_mut` et 652 `.clone()` restent inchangées** dans ce module.

Une première version ne faisait que développer l'arbre de décisions : 57 `match`, 7 637 caractères. Elle n'apportait aucun gain mesurable (+0,25 % RBTree, +0,02 % total sur 21 paires). Elle a été remplacée par la version qui partage la continuation. Son exécutable et ses résultats restent dans `after/` et `ab-results.json` afin de conserver cette expérience négative.

## Mesure sur le benchmark complet

Avant/après : mêmes sources PureScript, même runner natif MiMalloc, même profil `opt-level=3`, `debug=false`. Les manifests diffèrent sur un seul fingerprint d'entrée : le bundle `purust.js`. Le témoin est gelé dans `before-benchmark`; la version intégrée est dans `compact/benchmark`.

Chaque campagne comporte 21 paires avec ordre tiré au sort, sans autre build/benchmark lancé en parallèle. Chaque invocation utilise le warm-up global et le meilleur de dix mesures par cas du runner officiel. Les **14 valeurs attendues et la somme des temps ont été validées à chacune des 84 invocations**. Un passage supplémentaire par `./bin/rust/run --run-only --build-dir .../compact` valide également le manifest frais et les résultats.

| Campagne | RBTree avant, médiane | RBTree après, médiane | Variation RBTree appariée | Total avant, médiane | Total après, médiane | Variation totale appariée |
|---|---:|---:|---:|---:|---:|---:|
| 1 | 8 512,92 µs | 8 271,46 µs | −2,96 % | 9,38500 ms | 9,13312 ms | −2,73 % |
| 2 | 8 204,67 µs | 7 938,50 µs | −4,03 % | 9,03650 ms | 8,77297 ms | −3,65 % |

Les pourcentages sont les médianes des rapports après/avant de chaque paire, pas les rapports entre deux médianes indépendantes. Les temps absolus changent entre campagnes ; on attribue le gain aux comparaisons appariées. Le résultat intégré est légèrement inférieur au prototype manuel antérieur (environ −4 à −5 % sur RBTree) ; on ne remplace pas ces mesures par celles du prototype.

Contexte historique obligatoire : `altbak.pub-purust/README.md` affiche RBTree à **8 913,29 µs**, Records à **385,38 µs** et le total Rust à **9,78 ms**. Le README du projet initial `altbak.pub` donnait RBTree à **8 788,42 µs**. Ces baselines situent l'ordre de grandeur ; elles ne servent pas de témoin causal pour le changement présent. L'objectif historique de −50 % sur RBTree n'est pas atteint.

Les logs, commandes exactes, SHA256 des exécutables et résultats par module sont dans `final-comparison/results.json` et `final-repeat/results.json`. Reproduction depuis `altbak.pub-purust` :

```sh
python3 scratch/compact-guards-integration-20260916/compare.py --variant before scratch/compact-guards-integration-20260916/before-benchmark --variant compact scratch/compact-guards-integration-20260916/compact/benchmark --rounds 21 --output scratch/compact-guards-integration-20260916/reproduction
```

## Validation et limites

- Backend compilé et bundle généré ; runner complet recompilé depuis un workspace isolé.
- Nouveau test `tests/codegen/child-branch-printer.mjs` : 1 078 décisions ordonnées vérifiées sous Rc et Arc, qualification des constructeurs, partage effectif d'une continuation, refus de sortir un emprunt local, métadonnées invalides et limite globale de taille.
- Suite `node --test --test-concurrency=4 tests/codegen/*.mjs` : **75 réussites / 79 tests**. Les quatre tests FFI `crypto-hash-ffi`, `datetime-instant-ffi`, `foreign-object-foldm` et `uuid-ffi` sont bloqués par l'accès au socket Docker depuis le sandbox. Une inspection séparée du conteneur de référence `core-api-cli-1` l'avait confirmé arrêté. Ils ne sont pas comptés comme réussis.
- Tests de permutations et de réemploi existants conservés, y compris partage, références faibles, Rc/Arc et chemins de panique.
- Aucun `unsafe` ajouté, aucune preuve d'unicité déduite des compteurs source, aucun changement à `CodeGen.purs` pour cette intégration.

## Autre bloc de performance identifié

Le prototype Records, conservé séparément et non intégré au backend, passe de **376,17 µs à 3,54 µs** dans le runner complet : **−99,02 % sur Records et −4,04 % sur le total** en médiane appariée sur 21 paires. Il garde les quatre champs entiers dans des variables locales pendant la boucle et reconstruit le record à la sortie en conservant le comportement COW.

Les types structurels nécessaires sont **déjà présents dans le TAST**. La prochaine étape est une transformation générique de boucle dans le backend, pas l'ajout d'une annotation d'unicité. Des résumés de lectures/écritures par champ et de rétention des paramètres aideraient à généraliser cette transformation à travers les appels ; aucun gain supplémentaire n'est mesuré pour ces futures métadonnées.

Voir `../records-tast-opportunity-20260916/REPORT.md` pour la preuve, les alias conservés, les comptages et les limites sémantiques. Les gains des deux transformations n'ont pas été mesurés ensemble : ne pas les additionner comme un résultat acquis.
