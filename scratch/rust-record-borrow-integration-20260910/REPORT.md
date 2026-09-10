# Emprunter les projections de records vers des scalaires — intégration

10 septembre 2026. L’optimisation est intégrée dans Purust et mesurée sur le Rust réellement généré, sans retouche du benchmark. **Records : 0,625 → 0,353 ms, soit −0,272 ms (−43,5 %).** Les cinq paires sont favorables, avec des plages sans recouvrement. B14 reste jaune ; B22 reste vert. Cette étape ne réalise pas l’unboxing des records B21/B23.

## Règle générale

[RecordBorrows](../../../purust/purust/src/Purust/RecordBorrows.purs) reconnaît une chaîne de projections issue d’un local. Les rangées fermées du TAST doivent prouver chaque champ, sans label dupliqué, et le type final doit être exactement `Int`, `Number`, `Boolean` ou `Char`. Les annotations doivent rester compatibles avec les types des champs et la représentation de la racine.

[CodeGen](../../../purust/purust/src/Purust/CodeGen.purs) émet des getters `__purust_borrow_<champ>(&self) -> &UnknownType`, avec les mêmes contrôles de forme et de présence que les getters possédés. La lecture emprunte tout le chemin puis copie immédiatement le scalaire avec `unwrap_int/number/bool/char`. Aucun emprunt ne sort vers l’ABI `Value`. Le préfixe évite les collisions avec les getters ordinaires des champs `tally_ref` et `ref_tally`.

Les frontières `Typed`, opérateurs primitifs et conditions utilisent cette preuve au point d’évaluation existant. Ordre, court-circuit et chemins de copie lors des écritures sont conservés. Les racines calculées ou opaques, rangées ouvertes, annotations contradictoires, conversions et résultats possédés restent sur leur chemin existant. La règle ne dépend d’aucun nom de module ou de benchmark.

Les quatre RHS de Records changent, comme le montre le [diff généré](generated-records.diff). Les setters et le détachement du chemin restent identiques. La règle s’applique également à cinq modules de bibliothèque : `Data.Array`, `Data.Foldable`, `Data.List`, `Data.List.Lazy`, `Data.String.NonEmpty.Internal`. Avec Records et le runtime, sept sources Rust changent. **La source RBTree reste identique.** Les champs restent des `Option<Value>`.

## Mesures sur les runners complets

Ancien bundle sauvegardé avant les modifications, puis `bin/rust/run -c` pour reconstruire le compilateur, le TAST et la sortie Rust. Le bundle sauvegardé génère la référence sur ce même TAST. Les deux variantes ont le même `Cargo.lock`, le profil release O1, `debug=true` et mimalloc. La référence Cargo est reconstruite au propre ; aucun target copié n’est utilisé comme preuve de fraîcheur. Les empreintes des deux binaires sont distinctes et vérifiées.

Cinq paires alternées de processus complets, sans compilation, test ni instrumentation en parallèle. Chaque processus conserve le warm-up et le meilleur de dix du runner ; les 14 résultats sont vérifiés à chaque exécution. Les totaux ci-dessous sont les sommes des médianes de chaque benchmark.

| Mesure | Avant | Après | Écart observé |
| --- | ---: | ---: | ---: |
| Records | 0,625 ms | 0,353 ms | −0,272 ms (−43,5 %) |
| Plage Records | 0,620–0,628 ms | 0,352–0,356 ms | cinq paires favorables |
| RBTree, source inchangée | 11,348 ms | 11,254 ms | −0,094 ms |
| Suite complète | 12,455 ms | 12,084 ms | −0,371 ms (−3,0 %) |

| Paire | Records avant/après, µs | Suite avant/après, ms |
| --- | ---: | ---: |
| 1 | 624 / 355 | 12,507 / 12,149 |
| 2 | 628 / 352 | 12,459 / 12,020 |
| 3 | 620 / 356 | 12,359 / 11,944 |
| 4 | 626 / 353 | 12,304 / 12,085 |
| 5 | 625 / 353 | 12,569 / 12,100 |

Les cinq totaux sont favorables, mais l’écart global comprend la variation de RBTree et d’autres tests dont les sources n’ont pas changé. **Le gain solidement attribuable est celui de Records : 0,272 ms.** Il représente environ 2,1 % du total officiel documenté. Le [prototype précédent](../rust-record-borrows-20260910/REPORT.md) avait mesuré séparément 0,621 → 0,349 ms ; ses chiffres ne servent pas de baseline à cette intégration.

Le [README officiel](../../../altbak.pub/README.md#rust), relu le 10 septembre, indique Records **0,674 ms**, RBTree **11,754 ms**, suite **12,98 ms**. La dernière colonne native indique respectivement **0,004 / 36,070 / 36,13 ms**. L’intégration réduit le coût de Records mais reste loin de ses quatre microsecondes natives ; le compilé reste déjà devant cette référence native pour RBTree. Aucun nouveau temps n’est substitué aux baselines historiques du README.

## Comptage séparé et validation

Pour 10 000 itérations, le vrai noyau généré passe de **110 000 à 20 000 appels de clone PerceusPtr**. Les cinq clones dus à la lecture des quatre champs de vérification sont comptés séparément, après le snapshot. Les deux variantes conservent **3 allocations et 3 libérations**, aucune allocation dans la boucle. Les quatre champs sont vérifiés pour `n = 0, 1, 2, 10, 10000` ; le dernier résultat est `[10000, 20000, 30000, 20000]`.

Ces copies instrumentées O1 utilisent chacune leur propre runtime généré et le compteur d’allocations System. Les temps utilisent les runners mimalloc sans instrumentation. Les compteurs décrivent des appels dynamiques instrumentés, pas le nombre d’instructions de comptage survivant à LLVM dans les binaires chronométrés.

- `npm run build` réussi ; les avertissements des modules existants restent hors de cette modification.
- **27 tests de génération** passent, dont les refus de chemins, types et conversions non prouvés.
- **16 tests TAST** passent avec le fork explicitement sélectionné. La nouvelle fixture étendue a ensuite été relancée seule et passe.
- La fixture indépendante des benchmarks couvre les quatre types scalaires sur un à trois niveaux, négations/opérations/comparaisons, branche et court-circuit avec callbacks comptés, huit combinaisons de partage, écritures après lectures, adresses réutilisées, anciennes versions, retours possédés, captures, appels opaques évalués une fois et unwind. Elle compile le Rust O1 avec contrôles de débordement.
- `bin/rust/run -c` réussi, **14 résultats du runner** vérifiés avant/après, puis dans chaque paire.

## Artefacts

[Mesures brutes](runner-results.json), [compteurs](counts.json), [empreintes et sources modifiées](metadata.json), [validation](validation.json), [script de comparaison](compare.py), [script de comptage](count.py). Les artefacts volumineux et binaires sont sous `build/`, ignoré par Git. `compare.py build` reconstruit la référence déjà générée ; `compare.py time` répète les cinq paires. `count.py` travaille sur des copies indépendantes des deux sorties réellement générées.
