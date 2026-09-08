# Instanciation TAST dans PBO : intégration Rust

Le 8 septembre 2026, PBO préserve les arguments de type des appels jusqu'au développement de leur implémentation. Purust émet désormais l'accumulateur de Polymorphism en `i64`. Sur cinq paires alternées de runners complets, **Polymorphism passe de 37,358 ms à 1 µs** et **le total de 161,033 à 116,407 ms (−27,7 %)**. La compilation propre donne les 14 résultats attendus et les 17 tests passent.

## Changement intégré

- [Convert.purs](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/Convert.purs) conserve le `ForAll` de la déclaration, même si l'annotation du corps a déjà retiré le quantificateur. Les évaluateurs de primitives continuent à recevoir leurs seuls arguments d'exécution ; les `TypeApp` restent disponibles pour les implémentations PureScript.
- [Semantics.purs](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/Semantics.purs) transporte les arguments TAST sur chaque référence, puis substitue les paramètres dans une copie de l'implémentation au moment de l'inlining. Le cache générique reste partagé et inchangé. Les annotations identiques répétées à la racine ne sont pas prises pour des quantificateurs imbriqués.
- [TypeInstantiation.purs](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/CoreFn/TypeInstantiation.purs) substitue les types structurels en respectant les quantificateurs imbriqués, l'ordre des champs et les queues des rangées.

Cette première intégration spécialise les groupes de paramètres entièrement renseignés avec des types fermés. Les arguments encore ouverts, inconnus ou incomplets conservent la représentation générique. Il n'y a ni nom de benchmark dans la règle, ni nouveau mécanisme de spécialisation dans le générateur Rust. Le bundle CLI de Purust est reconstruit avec ce PBO.

Le [cas minimal](PolyI64.purs) et ses [traces](trace-summary.json) montrent `Int` et `Number` dans la récursion après la première passe PBO. Le [Rust réel avant/après](Polymorphism.diff) confirme un accumulateur `i64` et des conversions limitées aux frontières. La fonction générique conserve `UnknownType`. L'[assembleur de l'implémentation générée](kernel-after.s), produit sans modifier son Rust, ne contient plus de branche de boucle : LLVM réduit les incréments à une addition des paramètres. [Vérification reproductible](verify-generated.py), [résultat](generated-checks.log), [empreintes](validation.json).

## Validation

Le nouveau [test TAST](../../../purust/purust/tests/tast/type-instantiation.mjs) compile de vrais modules avec le fork et le prélude, lance le CLI normal puis compile des crates Rust neuves. Il vérifie :

- l'inlining depuis un autre module, avec les implémentations transmises par le cache PBO ;
- les accumulateurs natifs `Int` et `Number`, les appels alternés de ces deux instanciations et les dictionnaires restant génériques ;
- deux paramètres de type de sortes différentes (`Int`, `Boolean`), un callback, une application partielle et sa réutilisation ;
- zéro itération, plusieurs compteurs et des valeurs initiales positives, nulles et négatives ;
- la primitive polymorphe `Array.length`, qui doit conserver sa réduction intrinsèque ;
- le refus de spécialiser les types ouverts, les quantificateurs imbriqués et l'ordre des champs de rangées.

**14 tests de génération + 3 tests TAST passent**, dont ce nouveau test avec 153 assertions Rust et des vérifications de typage/génération. Les autres tests continuent à couvrir les effets, FFI, `Unit`, enums en valeur et arbres persistants. [Log génération](codegen-tests.log), [log TAST](tast-tests.log).

`bin/rust/run -c` réussit depuis le worktree Rust, avec reconstruction du backend, régénération et recompilation des crates. Les **14 sorties sont vérifiées** dans le [log propre](clean-run.log) et dans chacun des dix processus chronométrés. La compilation et les tests sont terminés avant les mesures retenues.

## Mesures du runner complet

Même profil release, `opt-level=1`, allocateur mimalloc et Rust 1.96.0. Le binaire de référence a été copié avant toute modification ; ses empreintes et celles du nouveau binaire sont dans les [résultats](results.json). Cinq paires de processus alternent l'ordre avant/après. Chaque processus effectue l'échauffement habituel et conserve le meilleur de dix essais par benchmark ; le tableau donne les médianes des cinq processus. Le total est la somme des médianes par benchmark.

| Benchmark | README officiel | Avant | Après |
| --- | ---: | ---: | ---: |
| Polymorphism | 38,940 ms | 37,358 ms | 0,001 ms |
| Church | 11,792 ms | 11,130 ms | 4,928 ms |
| RBTree | 62,425 ms | 40,992 ms | 40,462 ms |
| LazyEvaluation | 76,344 ms | 70,283 ms | 69,732 ms |
| Total | 190,940 ms | 161,033 ms | 116,407 ms |

La réduction mesurée est de **44,626 ms au total**. Polymorphism en représente **37,357 ms** ; Church bénéficie également de cette intégration, avec **6,202 ms de moins (−55,7 %)**. Les autres gros benchmarks restent proches de leur référence. Les baselines du [README officiel](../../../altbak.pub/README.md#rust) servent de repères historiques ; le gain de cette correction est établi par les paires avant/après, sans additionner les pourcentages des étapes précédentes.

Polymorphism, meilleurs temps en µs pour les cinq processus : voir les valeurs brutes dans [results.json](results.json). La précision du runner est à la microseconde dans cette série ; son résultat à 1 µs n'est pas directement comparable aux environ 0,10 µs du prototype isolé mesuré par lots. Le gain est confirmé dans le runner complet, avec les effets et les conversions qu'il conserve.

## Reproduction

Depuis `purust/purust` :

```sh
npm run build
npm run test:codegen
PURS="$PWD/../../altbak.pub-purust/run/bak/js/node_modules/.bin/purs" npm run test:tast
```

Depuis `altbak.pub-purust` :

```sh
bin/rust/run -c
python3 scratch/rust-poly-instantiation-20260908/measure.py
python3 scratch/rust-poly-instantiation-20260908/verify-generated.py
node --stack-size=65536 scratch/rust-poly-instantiation-20260908/trace.mjs
```

Le script de mesure utilise le binaire `before-binary` sauvegardé localement avant la correction. Les binaires, caches et traces complètes sont ignorés par Git ; sources, scripts, synthèses, empreintes et logs sont conservés. Les modifications de code concernent le worktree PBO Rust et Purust ; cette tâche ne modifie pas les checkouts habituels d'altbak.pub et de PBO.
