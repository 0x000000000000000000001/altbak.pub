# Polymorphism : diagnostic de l'accumulateur et prototype i64

Le 8 septembre 2026, le cas minimal confirme que la spécialisation de la récursion locale se perd dans **la première passe d'optimisation de PBO**. Le prototype limité à l'accumulateur du Rust déjà généré passe de **38,433 ms à environ 0,10 µs** pour l'effet `act`. L'assembleur explique cet écart : LLVM remplace les dix millions d'incréments par une addition. Ce prototype est isolé ; aucune optimisation n'est encore intégrée au générateur.

## Localisation par une exécution du pipeline

[PolyI64.purs](PolyI64.purs) contient la même récursion que le benchmark, une instance `Int` (+1), une instance `Number` (+0,5) et deux fonctions appelantes. [trace.mjs](trace.mjs) compile ces sources avec le fork TAST et les dépendances réelles du prélude, utilise le Builder PBO avec ses traces d'optimisation, puis lance le CLI Purust normal. Les 32 modules compilent ; les trois fonctions tracées ont respectivement trois, quatre et quatre états.

Les [états résumés](trace-summary.json) montrent :

1. À l'entrée de l'optimiseur, `intLoop` contient `TypeApp Int` et `numberLoop` contient `TypeApp Number`. Le TAST et son décodage transmettent donc bien les instanciations.
2. Après la première passe, ces `TypeApp` ont disparu et le corps développé contient une récursion de type `Int -> a -> a`. Le `a` n'a été remplacé ni par `Int`, ni par `Number`.
3. À la sortie de PBO, l'enveloppe extérieure est concrète, mais la fonction récursive et sa branche de retour portent encore `TypeVar a`. Le Rust produit des entrées natives `i64`/`f64` et des accumulateurs locaux `UnknownType`, avec les conversions à chaque tour ([Rust émis](PolyI64-generated.rs)).

Dans PBO, [Convert.purs](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/Convert.purs) préserve `ExprTypeApp` en `Syn.TypeApp` (ligne 613). [Semantics.purs](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/Semantics.purs) le transforme en `SemTypeApp` (ligne 323), mais `evalApp` descend dans la fonction en ignorant l'argument de type (ligne 477). Les chemins d'inlining filtrent également les `ExternTypeApp` (ligne 1150). Ces chemins sont les points à corriger et couvrir ; cette étape n'a pas encore isolé une modification suffisante à eux seuls.

La branche `Syn.TypeApp` de Purust n'est donc pas le premier endroit où intervenir pour ce cas : les traces montrent que la récursion y arrive déjà avec ses variables de type non substituées. La prochaine micro-étape devra préserver l'instanciation lors du développement du corps dans PBO, séparément pour chaque appel, tout en gardant `polyLoop` générique.

## Prototype mesuré

[prototype.py](prototype.py) prend le module réel `Purs_Test_Polymorphism`, en conserve une copie et compile deux variantes isolées contre les mêmes dépendances Rust du runner. Le [diff](prototype.diff) change seulement le paramètre et le retour de l'accumulateur local en `i64`, retire l'aller-retour `Value` dans l'incrément et conserve les conversions à l'entrée/sortie de la closure. Les dictionnaires génériques, effets, thunks, adaptateurs et `Bench.opaque` sont conservés.

Conditions : Rust 1.96.0, `opt-level=1`, informations de debug et mimalloc, comme le runner. Sept paires de processus alternées ; un échauffement puis dix échantillons par processus ; médiane des meilleurs échantillons. Pour dépasser la résolution du chronomètre, chaque échantillon natif regroupe 1 024 appels et est normalisé par appel, contre un seul appel de la référence. Les résultats sont vérifiés pendant les mesures. Les compilations et le traçage étaient terminés avant les échantillons retenus.

| Effet complet `act`, module isolé | Médiane | Meilleurs échantillons des sept processus |
| --- | ---: | --- |
| Rust généré actuel | 38 433,375 µs | 38 457,750 / 38 433,375 / 38 427,667 / 38 445,292 / 38 383,583 / 38 357,042 / 38 520,500 |
| Accumulateur `i64` | 0,103 µs | 0,102743 / 0,103312 / 0,102661 / 0,102295 / 0,103313 / 0,103027 / 0,102214 |

La [référence assembleur](kernel-before.s) garde une boucle ; le [prototype](kernel-native.s) contient une addition des deux paramètres, sans branche de boucle. Il calcule bien à partir d'arguments d'exécution, pas à partir d'un résultat codé en dur. Ce résultat démontre le potentiel sur cette boucle particulière, pas une accélération de plusieurs ordres de grandeur pour toute fonction polymorphe. [Échantillons et méthode](results.json), [log](prototype.log).

Le [README officiel d'altbak.pub](../../../altbak.pub/README.md#rust), relu pour cette étape, donne **38,940 ms** pour Polymorphism et **190,94 ms** au total. La dernière série après les enums donnait **40,261 ms** pour Polymorphism et **170,126 ms** au total. Le prototype n'a pas été intégré dans le runner complet : aucun nouveau total n'est mesuré ou revendiqué.

## Vérifications et reproduction

- Le Rust du cas minimal passe **60 assertions** : appels spécialisés `Int` et `Number`, appels génériques avec les deux dictionnaires, zéro itération et valeurs initiales variées.
- Chaque processus de benchmark vérifie **18 cas** sur une copie de la fonction récursive de sa variante : six compteurs, dont zéro et 100 000, et trois valeurs initiales. Chaque exécution chronométrée renvoie `10000000`.
- L'extraction de l'assembleur vérifie l'addition et l'absence de branche de boucle dans le prototype arm64. [Script](check-fixture.py), [log](fixture-checks.log).

Depuis la racine d'`altbak.pub-purust`, avec les dépendances release du runner déjà construites :

```sh
node --stack-size=65536 scratch/rust-poly-i64-20260908/trace.mjs
python3 scratch/rust-poly-i64-20260908/prototype.py
python3 scratch/rust-poly-i64-20260908/check-fixture.py
```

Ces commandes écrivent seulement dans ce dossier scratch. Les sorties compilées, caches PBO et traces complètes sont ignorés par Git ; scripts, synthèse, copies des petits modules et mesures permettent de les reproduire. Les [empreintes et références Git](validation.json) identifient les sources et dépendances utilisées, avec vérification que le Rust du runner est resté identique. Les sources de PBO et de Purust restent inchangées ; cette étape met uniquement à jour le suivi dans `purust/todo.md`.
