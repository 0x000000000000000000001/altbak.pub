# Unit sans allocation — intégration du 8 septembre 2026

`Unit` est désormais représenté par `()` dans le code Rust typé et par la variante immédiate `Value::Unit` aux frontières dynamiques. Les conversions `mk_unit` et `unwrap_unit` n'allouent rien ; elles conservent l'évaluation de l'expression convertie. `unwrap_unit` vérifie la variante reçue.

La génération des tableaux convertit maintenant leurs éléments natifs en `Value` : `[unit]` devient un tableau contenant `Value::Unit`. Les effets utilisent cette même valeur comme argument d'exécution. La FFI `Data.Unit` retourne `()`, et les vrais résultats/arguments `Unit` des FFI Effect, Console, Array et Assert sont alignés, y compris leurs copies présentes dans `purust-strings`.

Les littéraux de record vide et les témoins de contraintes `Partial` conservent leur représentation de record. Les cas `PrimUndefined` et record vide ne deviennent `()` que lorsque leur annotation est explicitement `Unit`. L'algorithme de LazyEvaluation et les rotations de RBTree n'ont pas été modifiés.

## Résultat dans le runner complet

Trois exécutions avant le changement et trois après la version finale. Chaque processus retient le meilleur de dix essais par benchmark ; le tableau donne la médiane des trois résultats. Les 14 valeurs fonctionnelles sont vérifiées à chaque exécution. Le total est la somme des médianes.

| Benchmark | Baseline README officiel | Avant | Après |
| --- | ---: | ---: | ---: |
| LazyEvaluation | 319,209 ms | 220,394 ms | **67,315 ms** |
| RBTree | 67,125 ms | 57,558 ms | 57,541 ms |
| Polymorphism | 38,663 ms | 39,948 ms | 36,842 ms |
| Church | 24,331 ms | 11,036 ms | 11,086 ms |
| State Monad | 0,514 ms | 0,396 ms | 0,063 ms |
| Total | 452,29 ms | **330,603 ms** | **174,147 ms** |

Le temps de LazyEvaluation baisse de **69,5 %** et le total de **47,3 %** dans cette comparaison. RBTree reste stable. La baseline historique vient du [README officiel](../../../altbak.pub/README.md#rust) ; les écarts historiques ne sont pas attribuables à ce seul changement. Les très petits benchmarks ont une résolution de mesure insuffisante pour interpréter leurs pourcentages de variation.

Même profil release `opt-level=1`, même allocateur mimalloc ; version du compilateur et empreintes du Rust généré dans `metadata.json`. Les processus de benchmark lancés pour cette tâche sont successifs. L'environnement n'est pas verrouillé en fréquence ou en affinité CPU ; les séries ne donnent pas d'intervalle de confiance. Les données brutes des 14 benchmarks sont conservées dans `results.json` et les logs `before-{1,2,3}.log` / `after-{1,2,3}.log`.

## Validation

- Build de Purust réussi, sans erreur ni avertissement PureScript.
- `npm run test:codegen` : **11 fichiers de tests passent**.
- `altbak.pub-purust/bin/rust/run -c` : réussite sur la version finale, **14 résultats corrects** ; log `clean-run-final.log`.
- Le test `purust/purust/tests/codegen/unit-values.mjs` vérifie les signatures natives, les conversions polymorphes, les tableaux, les records vides distincts, les continuations et les effets rejouables. Il compile aussi les FFI concernées, dont celles non chargées par la suite principale.
- Le compteur d'allocations est exécuté séparément des benchmarks temporels : **zéro allocation, zéro octet demandé et zéro libération pour 1 000 constructions de Unit**, contre 1 000 allocations et 56 000 octets dans le runtime minimal avant changement. Le passage d'une valeur et les allers-retours `()` / `Value` sont aussi sans allocation. Les allocations des tableaux et des closures ne sont pas annoncées comme supprimées.

Les FFI sont dans les checkouts Purust partagés, conformément au montage existant. Les builds altbak utilisent `altbak.pub-purust` et le PBO configuré dans `purescript-backend-optimizer-purust`. Cette tâche n'a pas modifié les checkouts habituels d'altbak.pub et PBO.

## Reproduction et suite

Depuis `purust/purust`, reconstruire avec le Spago du projet puis lancer `npm run test:codegen`. Depuis `altbak.pub-purust` :

```sh
bin/rust/run -c
python3 scratch/rust-unit-20260908/measure.py
```

Le script rejoue seulement la série après changement et la compare aux logs avant changement conservés ; il remplace `after-*.log`, `results.json` et `metadata.json`. L'[audit initial](../rust-audit-20260908/REPORT.md) et les sources de référence dans `originals/` permettent de retrouver l'état précédent. Les résultats intermédiaires avant correction du cas des tableaux restent dans `results-first-integration.json` et `after-first-integration-*.log`.

La prochaine micro-étape du chantier Unit est de compter les allocations restantes dans LazyEvaluation avant de reprendre les adaptateurs de thunks. Aucun gain supplémentaire sur ces adaptateurs n'est établi par cette intégration.
