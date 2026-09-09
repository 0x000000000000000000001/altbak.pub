# Fonctions retournées et récursion native — 9 septembre 2026

Purust génère désormais une fonction interne avec les paramètres des abstractions présentes dans l'AST lorsqu'une définition récursive renvoie une fonction calculée par une branche ou un let. L'interface publique reste identique. Les appels terminaux à cette fonction interne deviennent des affectations suivies de `continue`, au lieu d'adaptateurs `Func3`/`Func2` et d'une closure pour chaque appel récursif.

## Résultat mesuré

Cinq paires alternées des runners complets avant/après ; chaque processus effectue son échauffement et garde le meilleur de dix essais par benchmark. Le tableau donne les médianes sur cinq processus ; le total est la somme des médianes. Même Rust 1.96.0, `opt-level=1` et mimalloc. Les 14 sorties sont contrôlées pour chaque processus. Les allocations sont comptées séparément des temps.

| Benchmark | Avant | Après |
| --- | ---: | ---: |
| Lazy Evaluation (1M Thunks Forced, 1k Depth) | 67.352 ms | 19.136 ms |
| Church Numerals (100k Closure Applications) | 4.919 ms | 1.584 ms |
| Red-Black Tree (100k Worst-Case Insertions) | 39.934 ms | 40.151 ms |
| State Monad (1.2k Binds, 60 Stack Depth) | 0.063 ms | 0.049 ms |
| Total | 113.499 ms | 62.136 ms |

Le total baisse de 45.3 %. Le README officiel indique LazyEvaluation **72,002 ms**, Church **5,096 ms** et total **119,19 ms**. Sa dernière colonne indique **3 µs** pour LazyEvaluation : cette référence native a éliminé la chaîne. Le changement intégré conserve les thunks utiles et leur forçage. La comparaison avant/après ci-dessus mesure uniquement le changement sur le même environnement ; elle ne remplace pas les baselines officielles.

## Allocations du Rust réellement généré

Les deux noyaux sont extraits sans réécriture des modules Rust avant/après et compilés avec les mêmes dépendances récentes, O1 et mimalloc. Pour 1 000 répétitions de profondeur 1 000 :

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Allocations | 5003000 | 1000000 |
| Octets demandés cumulés | 160088000 | 32000000 |
| Libérations | 5003000 | 1000000 |

Il s'agit d'octets demandés cumulés, pas de mémoire simultanément résidente. Aucun cache de résultat ni formule remplaçant le calcul n'est introduit.

## Portée de la génération

- La transformation concerne les groupes récursifs d'une seule définition dont les paramètres explicites ne couvrent qu'une partie de la signature fonctionnelle. Elle conserve le chemin existant pour les autres groupes.
- Les annotations TAST déterminent les types des paramètres et de la fonction renvoyée. Une fonction locale masque le nom public dans son seul corps de génération ; les métadonnées exportées et la signature publique restent inchangées.
- Les appels directs et les références `FuncN::Static` utilisent le nombre de paramètres de l'appel natif, sans compter ceux de la fonction renvoyée.
- L'inférence d'un `App` consomme aussi ses arguments supplémentaires dans le type du résultat. La première reconstruction a révélé ce besoin dans StateMonad ; une fixture avec retour de record reproduit la panique `Expected Func1` avant cette correction.
- Les fonctions locales récursives et les groupes mutuellement récursifs ne sont pas étendus dans cette étape.

## Vérification

**14 tests de génération + 4 tests TAST passent**, ainsi que **`bin/rust/run -c` avec les 14 résultats corrects**. La nouvelle fixture compile le vrai TAST, le PBO et des crates Rust fraîches. Elle vérifie les profondeurs 0/1/2/17/1000, captures et nombre d'appels, applications partielles, forçage répété d'une closure partagée, fonctions retournées à deux arguments, récursion sous une closure, records et appels polymorphes Int/Number. Son budget d'allocation autorise au plus un objet par thunk utile afin de permettre de futures optimisations supplémentaires.

Les tests ont d'abord reproduit 5 003 allocations pour 1 000 thunks, puis la panique du retour de record. Le comptage séparé du runner vérifie également plusieurs profondeurs, la réutilisation d'un thunk capturé et l'absence de fuite.

## Fichiers et reproduction

- [Rust avant](LazyEvaluation-before.rs), [Rust après](LazyEvaluation-after.rs), [diff](LazyEvaluation.diff).
- [Mesures complètes](results.json), [allocations](allocations.json), [empreintes et validations](validation.json).
- `python3 measure.py` : cinq paires des binaires sauvegardés localement `before-binary`/`after-binary`.
- `python3 count.py` : recompilation des extraits avant/après avec les dépendances du runner puis vérification et comptage.

Les binaires, produits de compilation et logs sont ignorés par Git. Les modifications concernent Purust et les preuves dans altbak.pub-purust ; le PBO et les checkouts normaux restent inchangés.
