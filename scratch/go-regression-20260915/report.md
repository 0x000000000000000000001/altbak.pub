# Régression Go corrigée — 15 septembre 2026

La régression de la fusion des thunks est corrigée dans le PBO utilisé par Gopurs. Le total médian passe de **42,69 à 27,03 ms**, soit **36,7 % de temps en moins**. L’ancien binaire conservé mesure **26,50 ms** dans cette même série. Le résultat corrigé reste donc 2 % derrière cet ancien binaire ; **25,50 ms n’a pas été mesuré dans cette session**.

`bin/go/run --clean` reconstruit désormais Gopurs avant de compiler le benchmark. Le même comportement est restauré pour `bin/rust/run --clean` et Purust. L’empreinte du compilateur est enregistrée après sa reconstruction ; un échec arrête le build et ne produit pas de manifeste valide.

## Mesures comparables

Chaque variante a été exécutée dans cinq processus, dans un ordre alterné, sans compilation simultanée. Les benchmarks gardent leurs échauffements et le meilleur de dix mesures par test. Go 1.27.0, darwin/arm64, GOGC=800, sans PGO. Les 210 résultats numériques des 15 suites sont validés.

| Variante | Médiane des totaux | Plage des totaux | Lazy, médiane |
| --- | ---: | ---: | ---: |
| Ancien binaire conservé | 26,50 ms | 26,20–26,53 ms | 228,96 µs |
| Binaire régressé de la précédente campagne | 42,69 ms | 42,23–43,92 ms | 16 799,42 µs |
| Reconstruction corrigée | 27,03 ms | 26,45–27,31 ms | 229,08 µs |

Le README historique indiquait 24,02 ms ; l’utilisateur avait observé 25,50 ms. La campagne précédente publiait 44,73 ms. Ces observations historiques ne remplacent pas la comparaison alternée ci-dessus.

La colonne Gopurs du README est actualisée à partir des cinq mesures corrigées. Son total de **27,02 ms** est la somme des médianes des 14 lignes, selon la convention du tableau ; **27,03 ms** est la médiane des totaux de processus. Les autres colonnes conservent leurs mesures précédentes.

Sources : [résultats agrégés](results.json), [15 sorties validées](final-runs.json), [commandes et protocole](measurement-plan.json), fichiers `final-{old,regressed,fixed}-{1..5}.log` et `.json`.

## Cause établie et correctif

Le TAST de `Test.LazyEvaluation` est identique entre l’ancienne génération rapide et la campagne régressée : SHA-256 `f685d904e30741f555b6bc81747c58597e425310273d1df9093b77469b91f622`. Il contient déjà l’instanciation explicite `Int`.

La monomorphisation calculait une substitution pour les variables de type de la **définition**, puis l’appliquait à l’annotation du **site d’appel**. Les quantificateurs ont des noms distincts, par exemple `a$scope3` et `a$scope2`. L’appel spécialisé conservait donc un type générique `Unit -> a$scope2`, qui masquait `Unit -> Int` et faisait échouer les gardes de ThunkFusion.

Dans `purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/Monomorphize.purs`, l’annotation du site d’appel utilise maintenant sa propre substitution `subst`, au lieu de `info.subst`. Le traitement des dictionnaires statiques et dynamiques est conservé.

La preuve isolée change cette seule substitution dans une copie du bundle antérieur : avec le même TAST, le worker strict absent réapparaît et est appelé. **Toutes les gardes de ThunkFusion restent identiques**. Voir [résultat de la preuve](monomorph-callsite-proof/result.json), [patch appliqué](monomorph-callsite-proof/patch.json) et [diff Go](monomorph-callsite-proof/generated.diff).

Une simple reconstruction du compilateur avant ce correctif avait produit le même SHA de bundle et n’avait pas rétabli la fusion. Le cache et l’isolation du runner ne sont donc pas présentés comme la cause démontrée. Un premier essai de suppression de gardes a servi à localiser le blocage ; cette modification a été entièrement retirée du code final.

## Validation finale

- Build réel avec `bin/go/run --clean --build-only` : reconstruction du compilateur, TAST, génération et compilation Go réussies ; zéro erreur et zéro avertissement du compilateur.
- PBO : **56 tests passent**, dont trois nouveaux cas avec quantificateurs distincts au site d’appel et à la définition, sans dictionnaire, avec dictionnaire statique et avec dictionnaire dynamique. [Journal](pbo-tests.log).
- Test Go `ThunkFusion` existant : exécution et snapshot inchangé validés. [Journal](thunk-tests.log).
- Nouveau test Go `ThunkFusionNewtype` : résultats corrects, trois workers stricts, ordre de calcul conservé, fermeture échappée conservée et effets forcés deux fois. [Journal](thunk-newtype-tests.log).
- Runner : **7 tests simulés passent**, couvrant notamment la reconstruction Go/Rust, l’arrêt sur erreur, les empreintes et le rejet de `--run-only --clean`. La reconstruction réelle a été exécutée pour Go ; celle de Rust est couverte ici par les tests simulés.
- Benchmark final : 15 suites, 210 sorties numériques correctes. [Manifeste corrigé](final-manifest.json), [empreintes des sources](source-sha256.json).

Les sources Go de RBTree, LazyEvaluation et Polymorphism, le runtime et le chronomètre FFI sont identiques à ceux de l’ancien binaire rapide. Le petit écart résiduel de 2 % sur le total n’a pas été attribué par profilage ; il ne faut pas l’expliquer arbitrairement par le bruit ou prétendre avoir remesuré exactement 25,50 ms.

## Reproduction

Depuis `altbak.pub`, reconstruire et exécuter avec `bin/go/run --clean`, puis répéter avec `bin/go/run --run-only`. Le build par défaut est disponible dans `run/bak/go/modes/pure`.

Les tests Go doivent utiliser le fork PureScript typé : placer `altbak.pub/run/bak/js/node_modules/.bin` en tête de `PATH`, puis lancer `bin/test ThunkFusion ThunkFusionNewtype` depuis `gopurs/gopurs`.
