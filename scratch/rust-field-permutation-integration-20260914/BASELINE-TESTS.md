# Trois échecs TAST reproduits sans la nouvelle règle

Le 14 septembre 2026, les trois tests ont été rejoués **sans modifier leurs sources ni leurs assertions**, avec le bundle courant dont la seule nouvelle analyse `fieldPermutation` retourne `Nothing`. Ils échouent aux mêmes endroits que dans la suite courante. Ces échecs ne sont donc pas causés par l'activation de la permutation de champs.

Le préchargement [baseline-test-hook.mjs](baseline-test-hook.mjs) redirige exclusivement l'argument absolu `purust/purust/bin/purust.js` des appels `spawnSync` vers `build/purust-rule-disabled.js`, puis synchronise les exports ESM de Node. Le compilateur `PURS` est fixé explicitement au fork local `.stack-work/dist/aarch64-osx/ghc-9.8.4/build/purs/purs`. Chaque test compile un TAST frais. Chaque invocation du backend redirigé termine avec le code **0** avant l'échec du test.

| Test inchangé | Résultat sans règle | Observation établie |
| --- | --- | --- |
| `shared-nullaries.mjs` | Sortie Node 1 ; binaire Rust 101 | `checks.rs:101` échoue sur « All cells and weak control blocks are freed » : **23 allocations / 17 libérations**. Le défaut de comptage ou de durée de vie reste à diagnostiquer séparément ; ce contrôle ne prétend pas en établir toute la cause. |
| `record-root-move.mjs` | Sortie Node 1 à la ligne 38 | `updateDeep` contient `_base.set_b(purust_core::Value::Unit);` ; l'assertion impose le spelling `_base.set_b(crate::Value::Unit);`. L'échec est une attente textuelle devenue différente du code généré, avant compilation Rust. |
| `foreign-object.mjs` | Sortie Node 1 à la ligne 43 | `guardObject` traite `_foldM` comme un fallback et ne reconnaît pas sa forme. Le fichier généré contient une vraie implémentation multiligne de `Foreign_Object__foldM`, avec parcours des clés et callbacks ; le test ne passe donc pas son instrumentation de fallback. |

Les journaux complets, les empreintes du bundle effectif et des tests, les commandes redirigées et les chemins des sources copiées avant nettoyage sont conservés dans [baseline-test-results.json](baseline-test-results.json) et `build/baseline-test-evidence/`. Le diagnostic Foreign Object est également conservé par le test dans son propre répertoire temporaire, indiqué dans son journal.

Reproduction depuis ce dossier : `node baseline-test-runner.mjs`. Aucun correctif produit, aucune modification des tests de régression et aucune mesure de performance n'ont été effectués dans cette vérification.
