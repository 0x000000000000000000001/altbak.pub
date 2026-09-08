# Emprunts des tests de constructeur — 8 septembre 2026

Le générateur emprunte désormais les variables locales `Rc<ADT>` pour `OpIsTag`, au lieu de les cloner avant `matches!((parent).as_ref(), ...)`. Les wrappers `Typed` ne sont traversés que s'ils conservent la représentation Rust à chaque niveau ; `TypeApp` reste transparent. Les accesseurs, les conversions de représentation et les appels conservent leur chemin précédent.

Le retrait du local de l'ensemble `alive` concerne uniquement l'émission de cet opérande. Le parent reste disponible dans les branches et après le test. Le changement est dans `purust/purust/src/Purust/CodeGen.purs`, bloc `PrimOp (Op1 op a)` ; le bundle `bin/purust.js` est reconstruit.

## Vérifications

- Avant modification, le nouveau test compile et exécute correctement le code, mais échoue sur le clone temporaire encore émis pour la lecture de constructeur. Après modification, il passe.
- `tests/codegen/tag-borrows.mjs` couvre les constructeurs vides et avec champ, les parents partagés, la réutilisation dans les branches, les lectures répétées, les annotations imbriquées, la conversion `Value` vers ADT et les appels non locaux évalués une seule fois, y compris lorsque leur argument est réutilisé ensuite.
- Le test vérifie aussi l'absence de clones temporaires dans cinq lectures locales : leur coût serait invisible dans un simple comptage de références après le retour.
- Build Purust réussi sans erreur ni avertissement PureScript ; `npm run test:codegen` : **12 tests passent**.
- `bin/rust/run -c` dans le worktree : **réussite, 14 résultats corrects**. Log : `clean-run.log`.

Le module RBTree généré passe de **440 à 426 occurrences statiques de `.clone()`**, soit 14 supprimées. Ce comptage ne représente pas le nombre de clones exécutés par insertion. Les extractions de champs ne sont pas modifiées. Le diff du Rust émis figure dans `RBTree.diff`, avec les deux versions de référence.

## Mesures

Trois exécutions successives avant, puis trois après la modification. Chaque runner retient le meilleur de dix essais par benchmark ; le tableau donne la médiane des trois résultats. Les 14 valeurs sont contrôlées à chaque exécution. Même allocateur mimalloc et profil release `opt-level=1`.

| Mesure | Baseline officielle | Avant ce baby step | Après |
| --- | ---: | ---: | ---: |
| RBTree | 67,125 ms | 60,445 ms | **57,148 ms** |
| LazyEvaluation | 319,209 ms | 69,848 ms | 71,625 ms |
| Polymorphism | 38,663 ms | 37,433 ms | 38,050 ms |
| Church | 24,331 ms | 11,214 ms | 11,686 ms |
| Total | 452,29 ms | 180,281 ms | **179,837 ms** |

La baisse observée sur RBTree est de **5,5 %** ; le total reste globalement stable. Les autres variations ne sont pas attribuées à cette seule modification. Le total est la somme des médianes. Le [README officiel d'altbak.pub](../../../altbak.pub/README.md#rust) fournit les baselines historiques ; cette série est distincte des mesures précédentes d'Unit natif et du prototype combinant tous les emprunts. Aucun gain de 15 % n'est revendiqué pour `OpIsTag` seul.

Fréquence et affinité CPU ne sont pas verrouillées. Les résultats restent des observations locales, sans intervalle de confiance. Données détaillées : `results.json`, `before-{1,2,3}.log`, `after-{1,2,3}.log`. Version du compilateur et empreintes du Rust généré : `metadata.json`.

Depuis `altbak.pub-purust`, `python3 scratch/rust-tag-borrow-20260908/measure.py` rejoue les trois runs après changement et les compare aux logs avant changement conservés. Il remplace les fichiers `after-*.log`, `results.json` et `metadata.json`.

La prochaine micro-étape est l'emprunt du parent lors des extractions de champs, en conservant les clones d'enfants nécessaires au partage. Les checkouts habituels d'altbak.pub et PBO n'ont pas été modifiés par cette tâche.
