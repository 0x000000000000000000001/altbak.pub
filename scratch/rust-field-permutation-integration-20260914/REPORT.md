# B13 — permutation de champs intégrée dans Purust

14 septembre 2026. La règle générale est intégrée au backend et au bundle `bin/purust.js`. Sur cinq paires alternées, RBTree passe de **12,469 à 11,928 ms** en médiane (**−4,34 %**). La somme des médianes des 14 benchmarks passe de **13,427 à 12,888 ms** (**−4,01 %**). Tous les résultats fonctionnels sont conservés.

## Règle et limites

[FieldPermutations.purs](../../../purust/purust/src/Purust/FieldPermutations.purs) reconnaît à partir des types TAST une reconstruction de trois cellules du même constructeur, avec deux champs récursifs à des positions quelconques. La première branche gardée doit décrire une chaîne LL ou son miroir RR et préserver chacun des quatre sous-arbres frontières exactement une fois. Les autres champs sont des scalaires Copy, provenant de projections prouvées ou de constructeurs d'enums Copy sans champs. Aucun nom de module, de fonction ou de constructeur du benchmark n'est codé dans la règle.

[FieldPermutationPrinter.purs](../../../purust/purust/src/Purust/FieldPermutationPrinter.purs) conserve les trois cellules et l'adresse de la racine. Les tests de constructeurs et d'exclusivité précèdent toute écriture ; toutes les sources scalaires modifiées sont copiées avant les affectations, puis trois échanges de pointeurs réorganisent les champs. La règle s'insère dans le chemin existant après l'appel enfant, exécuté une seule fois.

Les cellules partagées ou observées par `Weak`, les autres branches et topologies, les conversions, appels opaques et scalaires calculés conservent le chemin existant. L'extension couvre la première branche LL de `balance` dans RBTree ; la fixture indépendante vérifie aussi RR en première branche. **B13 reste jaune** : ni toutes les rotations ni la fusion générale du comptage de références ne sont couvertes.

## Comparaison mesurée

Les deux bundles sont issus du même code courant ; le témoin désactive uniquement la nouvelle analyse. Générés depuis le dossier projet `altbak.pub-purust`, tous les fichiers Rust et manifests sont identiques entre variantes sauf `Purs_Test_RBTree/src/lib.rs`. Ce module témoin est également identique au Rust initial fourni pour l'analyse. [Empreintes et comparaison](generation-proof.json).

Les runners complets utilisent les mêmes dépendances, options release et verrou Cargo. Les trois crates concernées sont nettoyées avant chaque construction ; sources, binaires et verrou sont vérifiés avant mesure. Le warm-up global et le meilleur de dix du runner sont conservés. Cinq paires alternent avant/après puis après/avant, sans compilation ni instrumentation concurrente de cette tâche. Construction, profondeur et destruction restent incluses. [Protocole](runner.py), [empreintes des builds](runner-build.json), [mesures complètes](runner-results.json).

| Paire | RBTree avant | RBTree après | Réduction |
| --- | ---: | ---: | ---: |
| 1 | 12,707 ms | 11,511 ms | 9,41 % |
| 2 | 12,462 ms | 11,843 ms | 4,97 % |
| 3 | 12,469 ms | 11,928 ms | 4,34 % |
| 4 | 12,304 ms | 11,955 ms | 2,84 % |
| 5 | 13,019 ms | 12,215 ms | 6,18 % |
| Médiane | **12,469 ms** | **11,928 ms** | **4,34 %** |

Le total de 13,427 → 12,888 ms est la somme des médianes par benchmark ; la médiane des totaux observés vaut 13,425 → 12,884 ms. Chaque exécution vérifie les 14 valeurs attendues, y compris profondeur RBTree = 22. Les autres benchmarks ont un code identique ; leurs petits écarts ne sont pas attribués à cette règle. Aucun intervalle de confiance n'est établi par ces cinq paires.

Le [README officiel](../../README.md) documente RBTree Rust **11,631 ms**, total **12,56 ms**, et le natif de dernière colonne **36,070 ms**, total **36,13 ms**. Le dernier run utilisateur donnait **12,198 ms / 13,17 ms**. La comparaison contrôlée ci-dessus établit le gain de la règle dans cette session ; elle ne démontre pas un nouveau record contre l'historique. Les anciennes mesures du prototype ne sont pas substituées à celles de l'intégration.

Le dossier de sortie utilisé par le projet a été régénéré et reconstruit après qualification. Son RBTree a la même empreinte que la variante après, et son [run final](evidence/live-final.log) valide les 14 résultats (total indicatif 12,67 ms). Une première préparation depuis un mauvais répertoire ne retrouvait pas les FFI : elle a été remplacée avant toute mesure retenue.

## Cause vérifiée par instrumentation

Sur la construction de 100 000 éléments, l'instrumentation du code effectivement émis observe :

| Opération | Avant | Après |
| --- | ---: | ---: |
| Extractions `__purust_take` | 299 934 | 0 |
| Reconstructions par les helpers | 299 934 | 0 |
| Permutations générées | 0 | 99 978 |
| Clones Rc | 2 483 932 | 2 283 976 |
| `get_mut` réussis | 2 783 866 | 2 483 932 |
| Allocations de cellules | 100 001 | 100 001 |

[Compteurs](counts.json), [script](count.py). Ces compteurs sont mesurés séparément du chronométrage. Les empreintes des sources correspondent aux variantes finales, et les scénarios persistants conservent leurs compteurs.

## Validation et échecs préexistants

- Build du backend réussi, bundle actualisé. Le rebuild des anciens modules a signalé leurs avertissements existants ; la dernière compilation incrémentale termine sans avertissement.
- Nouvelle [suite de preuve et d'émission](../../../purust/purust/tests/codegen/field-permutations.mjs) : 53 cas acceptés/rejetés, six dispositions/directions exécutées sous Rc et Arc, avec cellules uniques, partagées et faibles ; adresses et sous-arbres préservés.
- Nouvelle [fixture TAST indépendante](../../../purust/purust/tests/tast/field-permutations.mjs) : 320 cas de partage, 32 cas de gardes/ordre, dispositions de champs différentes, refus des formes non linéaires, versions persistantes, adresses, absence d'allocations sur le chemin prouvé, callbacks, unwind et libérations.
- [Validation RBTree générée](validation.json) : invariants, comparaison BTreeSet, 200 versions persistantes, 512 insertions à partage mixte, références faibles et durées de vie, ainsi que 54 cas d'exclusivité/gardes et six motifs absents.
- Suite codegen : **77/77 passent**, après relance des tests Docker et téléchargement de dépendances manquantes. [Run initial](evidence/codegen.log), [relance](evidence/codegen-retry.log), [UUID](evidence/codegen-uuid.log).
- Suite TAST : **33/39 passent**, bilan consolidé après sélection explicite du fork TAST local et récupération d'une dépendance Rust. [Relance avec le fork](evidence/tast-retry.log), [nouvelle fixture](evidence/tast-fork-check.log), [crypto](evidence/tast-crypto.log). Les six échecs ci-dessous restent ouverts ; aucun test n'a été assoupli pour les masquer.

| Test TAST | Preuve de préexistence |
| --- | --- |
| `json-read-object` | Les 171 sources Rust Rc et Arc produites avec le bundle initial sont identiques ; le témoin Rc reproduit 16/17 avec le même défaut d'ordre des clés. [Comparaison](evidence/json-order-before-report.json), [exécution](evidence/json-order-before-test-normal.json). |
| `bigint` | Le scénario négatif exige une erreur mais le FFI courant émet déjà un type opaque. L'émetteur extrait de HEAD reproduit le suffixe exact. [Preuve](evidence/bigint-missing-ffi-head.json). |
| `nullable` | Même décalage du scénario négatif avec le type opaque émis par HEAD. [Preuve](evidence/nullable-missing-ffi-head.json). |
| `shared-nullaries` | Le témoin avec règle désactivée reproduit le même échec : 23 allocations / 17 libérations. La cause précise reste à diagnostiquer. |
| `record-root-move` | Le témoin reproduit l'attente textuelle `crate::Value::Unit`, alors que le code émet `purust_core::Value::Unit`. |
| `foreign-object` | Le témoin reproduit l'instrumentation qui attend un fallback pour `_foldM`, déjà remplacé par une implémentation. |

Les trois derniers témoins, commandes et résultats sont détaillés dans [BASELINE-TESTS.md](BASELINE-TESTS.md) et [baseline-test-results.json](baseline-test-results.json). Les journaux complets et binaires temporaires restent dans `build/` du scratch `altbak.pub-purust` ; seules les preuves utiles sont archivées ici.

## Reproduction

Le compilateur de tests est explicitement `/Users/0x1/Documents/htdocs/purescript/.stack-work/dist/aarch64-osx/ghc-9.8.4/build/purs/purs` ; le `purs` standard présent dans le PATH ne produit pas le TAST requis.

Depuis `altbak.pub-purust`, générer les variantes avec le bundle courant et [disabled-baseline.mjs](disabled-baseline.mjs), en conservant ce répertoire comme cwd pour la résolution des FFI. Les options sont `--main App --source run/bak/rust/output --out scratch/rust-field-permutation-integration-20260914/build/{after,disabled}`. Puis, depuis ce scratch : `python3 runner.py build`, `python3 runner.py time`. Les validations/instrumentations se lancent séparément avec `python3 count.py validate` et `python3 count.py count`. Les scripts utilisent les ressources des expériences voisines déjà présentes dans le projet.
