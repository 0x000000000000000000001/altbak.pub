# Réutiliser le dernier niveau de Records

9 septembre 2026. Intégration dans Purust, validation dans `altbak.pub-purust` avec PBO `-purust`. Les checkouts normaux et les FFI des benchmarks ne sont pas modifiés.

## Résultat livré

La règle de réutilisation suit maintenant **un chemin de records fermés**, avec au plus un enfant choisi par niveau. Elle réutilise aussi `d` dans Records, après la racine et `b` déjà couverts. Le chemin est prouvé par le layout TAST et les projections depuis la même variable locale ; aucun nom de benchmark ni limite de profondeur spécifique à Records n'intervient.

| Mesure | Avant : racine et un enfant | Après : chemin jusqu'à la feuille |
| --- | ---: | ---: |
| Records, médiane de cinq runners complets | 0,768 ms | 0,624 ms |
| Allocations pour 10 000 itérations, initialisation comprise | 10 003 | 3 |
| Libérations, résultat détruit | 10 003 | 3 |
| Total, somme des médianes des 14 benchmarks | 19,434 ms | 19,322 ms |

**Records gagne 18,8 %**, avec les cinq paires favorables et des plages disjointes : avant 0,756–0,775 ms ; après 0,595–0,634 ms. Les trois allocations restantes construisent le record initial ; aucune allocation supplémentaire n'est comptée pendant ces 10 000 mises à jour uniques.

La baisse du total est **0,112 ms (−0,6 %) sur cette série**, avec quatre paires favorables et une défavorable. Elle reste petite face aux variations observées, notamment sur RBTree (source inchangée) : 17,956–18,206 ms avant, 18,187–18,260 ms après. Le gain local est établi par cette expérience ; cinq paires ne permettent pas d'affirmer un gain global stable de cette ampleur.

Le README officiel du checkout normal, relu pour cette étape, affiche **Records 0,969 ms / natif 0,004 ms**, **RBTree 18,449 ms / natif 36,070 ms**, **total 19,95 ms / natif 36,13 ms**. Ces baselines historiques restent distinctes des mesures appariées ci-dessus. Le natif Records retourne seulement `f`, ce qui permet d'éliminer les autres accumulateurs ; nos contrôles isolés observent les quatre champs du record complet.

## Règle et limites

[RecordUpdates.purs](../../../purust/purust/src/Purust/RecordUpdates.purs) construit un plan récursif : à chaque niveau, le premier enfant éligible est suivi et les autres valeurs de remplacement conservent leur génération habituelle. Chaque projection doit désigner le même chemin depuis la racine ; des wrappers compatibles sont acceptés. Un autre parent, un autre chemin/champ, une base calculée par appel, une rangée ouverte, une conversion ou des labels dupliqués arrêtent la réutilisation à ce niveau. La racine conserve les gardes précédentes de local fermé, sans conversion ni usage ultérieur/capture réutilisable.

[CodeGen.purs](../../../purust/purust/src/Purust/CodeGen.purs) aplatit les remplacements du plan dans leur ordre d'origine et les évalue tous en gardant la racine vivante. Il déplace ensuite la racine et détache les enfants de l'extérieur vers l'intérieur, au moyen des getters/setters existants et des slots temporaires `Value::Unit`. Après modification, il réinstalle les enfants de l'intérieur vers l'extérieur. Aucune expression source ni callback n'est évalué pendant cette mutation.

Les tests d'unicité existants de `PerceusPtr::make_mut` restent actifs à chaque niveau. Une racine partagée entraîne la copie de son chemin modifié ; un enfant ou une feuille partagé seul conserve également son ancienne version. Les remplacements qui capturent une ancienne version déclenchent le même mécanisme. Le runtime, les champs `Option<Value>` et les layouts ne changent pas.

Cette extension couvre un chemin, pas toutes les branches d'une mise à jour. **B14 reste jaune et B22 reste vert**. Elle ne constitue ni un unboxing des records, ni une passe Perceus branche par branche sur les opérations de comptage.

## Expérience et validation

- [Prototype](probe.py) à partir des expressions du Rust réellement généré : **10 003 → 3 allocations/libérations**. Les huit combinaisons indépendantes de partage racine/enfant/feuille et les 24 combinaisons de seed/itérations passent avant et après. Cinq paires isolées suggéraient **0,835 → 0,655 ms** avec une forte dispersion au début ; le résultat livré vient des runners complets ci-dessus.
- `npm run build` : compilation et bundle réussis ; 64 avertissements dans CodeGen/Main lors de la recompilation, aucun dans RecordUpdates, zéro erreur. La compilation PureScript propre du runner réussit sans avertissement.
- **22 tests de génération et 13 tests TAST réussis.** La génération couvre jusqu'à trois niveaux enfants, les wrappers, les chemins différents, les feuilles ouvertes/calculées/converties et le maintien d'un seul chemin parmi des frères éligibles. La fixture TAST fraîche vérifie les adresses des trois cellules, les huit combinaisons de partage, l'ordre des callbacks, les captures des anciennes versions jusque dans la feuille, les exceptions et la libération des valeurs temporaires. Les tests précédents, dont les débordements, restent actifs.
- `bin/rust/run -c` : régénération et compilation complètes ; **14 résultats attendus vérifiés**. Les dix runners de la comparaison vérifient aussi ces 14 résultats à chaque exécution.
- [Comparaison de sources](source-comparison.json) : **seul Records change parmi 302 bibliothèques Rust**, runtime compris. [Diff généré](Records.diff).
- [Comptage du code intégré](generated-checks.json), séparé du chronométrage : **3 allocations et 3 libérations** pour 0, 1, 2, 10 et 10 000 itérations ; quatre champs vérifiés. Les huit combinaisons de partage et 24 combinaisons d'entrées passent avant et après.
- [Mesures finales](runner-results.json) : cinq paires alternées de runners complets O1/mimalloc, warmup et best-of-10 habituels, puis médianes entre processus. Toutes nos compilations et instrumentations étaient terminées. Le total additionne les médianes des 14 benchmarks, comme dans les expériences précédentes.

[Révisions, empreintes et baselines](metadata.json), [validation](validation.json), [script de comptage](check-generated.py), [contrôles de partage](sharing-check.rs), [script de mesure](measure-runner.py). Les binaires, sources complètes et logs de travail sont ignorés sous `build/` ou `*.log`.

## Suite

La piste des allocations de copies imbriquées est terminée pour ce chemin unique : il ne reste que l'initialisation. Avant une nouvelle optimisation Records, isoler le coût des accès répétés aux champs et des valeurs dynamiques `Option<Value>` dans le code machine O1, puis mesurer un prototype qui conserve le résultat complet. Les layouts typés/unboxés B21/B23 sont une piste distincte, sans gain attribué à cette étape. Les cas d'ADT et de parcours empruntés du plan restent ouverts.
