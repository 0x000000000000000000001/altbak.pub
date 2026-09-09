# Réutiliser un enfant de record

9 septembre 2026. Intégration dans Purust, validée dans `altbak.pub-purust` avec PBO `-purust`. Les checkouts normaux et les FFI des benchmarks ne sont pas modifiés.

## Résultat livré

Une mise à jour `r { enfant = r.enfant { ... } }` peut maintenant réutiliser **un enfant immédiat** d'une racine déplacée au dernier usage. Le layout fermé du TAST, l'identité de la variable locale et celle du champ définissent la règle ; aucun nom de benchmark n'intervient.

| Mesure | Avant, avec déplacement de la racine | Après, avec un enfant réutilisé |
| --- | ---: | ---: |
| Records, médiane de cinq runners complets | 0,861 ms | 0,767 ms |
| Allocations pour 10 000 itérations, initialisation comprise | 20 003 | 10 003 |
| Libérations, résultat détruit | 20 003 | 10 003 |
| Total, somme des médianes des 14 benchmarks | 19,735 ms | 19,733 ms |

**Records gagne 10,9 %**, avec les cinq paires favorables. Plages des mesures par processus : avant 0,845–0,868 ms ; après 0,709–0,771 ms. **Aucun gain global n'est établi** : le total reste dans la dispersion, et RBTree, dont la source est inchangée, varie de 18,303–18,917 ms avant à 18,214–18,715 ms après.

La baseline officielle relue dans le README du checkout normal reste **Records 0,969 ms / natif 0,004 ms**, **RBTree 18,449 ms / natif 36,070 ms**, **total 19,95 ms / natif 36,13 ms**. Ces chiffres historiques ne remplacent pas la comparaison appariée. Le natif Records retourne seulement `f`, permettant d'éliminer les autres accumulateurs ; les contrôles isolés observent les quatre champs du résultat complet.

## Règle et propriété

[RecordUpdates.purs](../../../purust/purust/src/Purust/RecordUpdates.purs) reconnaît un enfant record fermé dans la rangée TAST du parent. Les wrappers compatibles sont acceptés ; un autre parent, un autre champ, une base calculée par appel, une rangée ouverte, une conversion ou des labels dupliqués conservent le chemin existant. Le déplacement de racine impose toujours une base locale sans usage ultérieur ni capture réutilisable.

Le générateur calcule les remplacements de la racine et de l'enfant dans leur ordre d'origine, en gardant la racine vivante. Il déplace ensuite la racine, récupère l'enfant avec le getter existant, remplace temporairement son slot par `Value::Unit`, applique les setters de l'enfant, puis réinstalle celui-ci. Aucune expression source ni callback n'est évalué pendant cette phase de mutation. Le getter conserve l'enfant avant le détachement ; le détachement ne détruit donc pas son contenu.

Sur une racine unique, détacher l'enfant supprime la référence que la racine détenait : l'enfant unique devient réutilisable. Sur une racine partagée, les setters copient d'abord la racine et les anciennes versions conservent leurs références à l'enfant ; `PerceusPtr::make_mut` conserve alors le chemin de copie de l'enfant. Les alias retenus par les remplacements/captures obéissent au même mécanisme.

Le runtime et les layouts `Option<Value>` restent identiques. La règle ne sélectionne qu'un enfant immédiat par mise à jour ; dans Records, `b` est réutilisé tandis que `d` conserve sa copie. B14 reste jaune, B22 reste vert. Cette étape n'ajoute ni unboxing général, ni passe Perceus branche par branche sur les opérations de comptage.

## Expérience et vérifications

- [Prototype](probe.py) appliqué aux expressions du Rust réellement généré : 20 003 → 10 003 allocations/libérations, 24 combinaisons d'entrées avec conservation des anciennes racines et enfants. Cinq paires isolées suggéraient 0,903 → 0,812 ms, avec une dispersion importante au début. Ce relevé est distinct du résultat intégré ci-dessus.
- `npm run build` : compilation et bundle réussis. La recompilation affiche 64 avertissements dans CodeGen/Main ; aucun dans le nouveau module RecordUpdates. Le build cached du runner réussit également.
- **22 tests de génération et 13 tests TAST réussis**. La fixture TAST existante est étendue aux adresses uniques, aux racines partagées, au partage de l'enfant seul, à l'ordre des callbacks, à une closure qui retient l'ancienne racine et l'ancien enfant, aux exceptions et à la libération des remplacements temporaires. Les contrôles précédents, dont les débordements, restent actifs.
- `bin/rust/run -c` : régénération complète, compilation réussie, **14 résultats attendus vérifiés**. Les dix processus des cinq paires vérifient également ces 14 résultats.
- [Comparaison des sources](source-comparison.json) : **seul Records change parmi 302 bibliothèques Rust**. [Diff généré](Records.diff).
- [Comptage du code intégré](generated-checks.json), sans chronométrage : allocations égales aux libérations pour 0, 1, 2, 10 et 10 000 itérations, quatre champs vérifiés ; anciennes versions vérifiées pour 24 combinaisons d'entrées avant et après.
- [Mesures finales](runner-results.json) : cinq paires alternées avant/après, runners complets O1/mimalloc, warmup et best-of-10 habituels puis médianes entre processus. Toutes nos compilations et instrumentations étaient terminées avant ce chronométrage. Les totaux additionnent les médianes de chaque benchmark.

[Révisions et empreintes de départ](metadata.json), [validation](validation.json), [script de comptage](check-generated.py), [script de mesure](measure-runner.py). Les binaires, sources complètes et logs de travail sont ignorés sous `build/` ou `*.log`.

## Prochain bébé

La copie de `d` explique encore **10 000 allocations pour 10 000 itérations**. Prototyper la réutilisation de ce dernier niveau avant d'étendre la reconnaissance à un chemin plus profond ; conserver les contrôles de partage indépendant à chaque profondeur et l'ordre des RHS. Le plancher théorique de ce cas unique est de trois allocations initiales, mais aucun gain temporel de cette extension n'est encore démontré dans le générateur.
