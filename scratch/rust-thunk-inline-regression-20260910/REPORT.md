# Régression de fusion des thunks après validation des bibliothèques

10 septembre 2026. La restriction d’inlining ajoutée pour protéger la portée des types polymorphes empêchait un constructeur de newtype de disparaître. La fusion des thunks ne reconnaissait plus le producteur. **Le correctif final ramène le total de 112,203 à 12,723 ms**, soit **99,480 ms récupérées**, sur cinq paires alternées du runner complet.

## Cause et correction

Le Rust avant correction reconstruit des closures `Rc` imbriquées avec conversions `i64 → Value → i64`. Après correction, le consommateur appelle de nouveau le worker strict sur deux entiers. Le benchmark définit son propre newtype ; il n’utilise pas les mutex de Ref/AVar ni le runtime Tokio. Le runner reste sans `--threaded` et le runtime local Perceus est inchangé.

Le [PBO](../../../purescript-backend-optimizer-purust/src/PureScript/Backend/Optimizer/Semantics.purs) reconnaît uniquement une implémentation structurée comme `Typed* → Abs [paramètre] → Typed* → Local même niveau`. Pour une application à un seul argument, il restitue exactement la valeur sémantique de l’appelant, sans évaluer le corps annoté ni recopier ses types génériques. Les directives `InlineNever` et `InlineArity` restent respectées. La garde des `ForAll` non instanciés conserve les autres fonctions et les références nues.

La règle ne dépend d’aucun nom de module, de newtype ou de benchmark. Aucune modification du reconnaisseur `ThunkFusion`, des bibliothèques ou du runtime n’est nécessaire. Les sources Rust de 42 modules de bibliothèque changent également ; parmi les modules de benchmark, seul `Test.LazyEvaluation` change.

La revalidation a aussi détecté une erreur préexistante dans `purust-functions` : les déclarations `Fn2` à `Fn10` exposent un résultat opaque `Value`, alors que leur corps optimisé produit une fonction native. L’ancien bundle et celui avec le seul correctif d’identité génèrent exactement le même Rust erroné. Le générateur convertit désormais le corps des bindings sans argument vers leur type de retour déclaré, comme pour les autres retours. Cette correction est incluse dans le binaire final et dans les validations ci-dessous.

## Mesures

Les bundles avant et après correction génèrent chacun leur sortie depuis les mêmes fichiers TAST, dont les empreintes sont vérifiées. Le Rust Lazy régénéré par l’ancien bundle est identique à celui du run signalé. La référence Rust est compilée dans un nouveau répertoire, avec le même `Cargo.lock`, le même runtime, mimalloc et le profil release `opt-level=1`, `debug=true`. Aucun build ni test n’accompagne les mesures.

Le protocole conserve le warm-up et le meilleur de dix du runner. Les quatorze résultats sont contrôlés à chaque processus. Les agrégats sont les sommes des médianes par benchmark.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Lazy Evaluation | 99,536 ms | 0,000 ms |
| Total | 112,203 ms | 12,723 ms |
| Plage des cinq totaux | 110,661–112,971 ms | 12,699–12,901 ms |

| Paire | Avant | Après |
| --- | ---: | ---: |
| 1 | 112,971 ms | 12,901 ms |
| 2 | 112,346 ms | 12,699 ms |
| 3 | 110,661 ms | 12,716 ms |
| 4 | 110,966 ms | 12,726 ms |
| 5 | 112,575 ms | 12,768 ms |

Une [première série avec la seule correction d’identité](identity-only-runner-results.json) mesurait 108,727 → 12,507 ms, et Lazy 96,273 → 0,001 ms. Elle confirme l’origine de la récupération. Les deux séries ont leurs propres baselines ; leurs totaux ne se soustraient pas entre eux.

Le [README officiel](../../../altbak.pub/README.md#rust) indique **12,56 ms** au total et **1 µs** pour Lazy Evaluation. Le résultat corrigé retrouve cet ordre de grandeur ; ces mesures ne démontrent pas une amélioration par rapport à cette baseline. Les valeurs de 0–1 µs sont à la résolution du chronomètre, pas une absence de travail.

Le nettoyage complet via `./bin/rust/run -c` réussit également avec la correction d’identité, les quatorze résultats corrects et Lazy à 0 µs. Son total de 13,79 ms a été obtenu pendant les validations en parallèle et ne remplace pas les mesures isolées ci-dessus. Après la correction des valeurs Fn, la sortie finale est régénérée et compilée sur le même TAST, dont les 301 empreintes restent identiques après le nettoyage.

## Validation finale

**41 vérifications codegen, 16 suites TAST et 13 runners de paquets passent** avec le compilateur final : arrays, foldable-traversable, unfoldable, unsafe-coerce, partial, assert, console, effect, refs, exceptions, functions, avar et aff. Les runners de paquets utilisent `bin/test`, qui régénère leur TAST et leur sortie Rust ; Cargo utilise les dépendances disponibles hors ligne. Aucun test existant n’est supprimé ni affaibli.

Aff conserve ses 45 tests, les contrôles de concurrence Ref/AVar et les quatre scénarios de durée de vie des enfants. Le contrôle de portée des types `Wrapped 42` passe toujours. La nouvelle régression des valeurs décurryfiées compile et exécute les arités 2 et 10, les captures et le rejeu. Les [résultats de validation](validation.json) enregistrent les commandes et résumés des runners.

## Reproduction et preuves

- [Script de mesure](compare.py) : `python3 scratch/rust-thunk-inline-regression-20260910/compare.py`.
- [Résultats détaillés](runner-results.json), [métadonnées de comparaison](comparison-metadata.json), [empreintes initiales](before-metadata.json), [sources Rust modifiées](changed-rust-sources.json).
- `build/` contient les deux exécutables comparés, le bundle précédent, les sorties Rust et les dix logs. Ce répertoire est ignoré par Git.
- La fixture TAST indépendante ajoute `Suspended a`, un alias polymorphe de constructeur et un consommateur polymorphe. Elle échoue avec le bundle précédent et passe avec le correctif, tout en conservant les cas préexistants.
- Quatre contrôles sémantiques vérifient la conservation exacte des annotations de l’appelant, l’absence d’exécution des fonctions/effets, le refus des autres formes et le respect des directives d’inlining.
