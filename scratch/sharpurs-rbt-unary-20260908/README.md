# RBTree : fonctions unaires natives

Intégration de `depth` et `makeBlack` natifs dans le générateur normal : **−4,16 % sur RBTree et −3,49 % sur le total médian**, avec une contrepartie de **+10,81 % d’octets alloués par action RBTree**. Cette étape améliore le temps mesuré ; la réduction d’allocations reste à obtenir.

## Résultats

Médianes de cinq processus indépendants par variante, conservant le score officiel best-of-10 de chaque test :

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total | 295,98 ms | 285,64 ms | −3,49 % |
| RBTree | 215,607 ms | 206,647 ms | −4,16 % |
| Lazy, code inchangé | 68,402 ms | 68,349 ms | −0,08 % |

Étendue des totaux : **292,06–299,11 ms** avant, **271,95–286,49 ms** après. Étendue RBTree : **210,709–218,918 ms** avant, **194,770–207,185 ms** après. La série contient un processus après sensiblement plus rapide ; la conclusion utilise la médiane des cinq. Le total médian après est également inférieur de 3,11 % à la baseline officielle figée de **294,80 ms** ; RBTree est inférieur de 3,33 % à ses **213,760 ms**. Cette référence historique contextualise la comparaison simultanée, sans remplacer son témoin.

Le diagnostic séparé, trois processus et trente actions par variante, retrouve **213,230 → 204,964 ms** pour la médiane des meilleurs temps et **219,049 → 212,240 ms** pour la médiane des médianes. Les allocations médianes passent cependant de **740 065 552 à 820 095 896 octets/action** (+10,81 %). Les collections cumulées sur trente actions sont Gen0 **2 670 → 2 952**, Gen1 **51 → 36**, Gen2 **18 → 12**. Il s’agit d’octets alloués cumulés sur le thread de calcul, pas de mémoire résidente.

Les **140 sorties officielles**, **60 résultats diagnostiques** et **18 échauffements locaux** concordent ; chaque action RBTree diagnostique conserve les 100 000 insertions et la profondeur **22**. Les résultats détaillés, échantillons, résumés recalculés, métadonnées et empreintes des assemblies figurent dans `comparison.json`, `metadata.json` et `executions.json`.

`ins` et `balance` manipulent encore des `obj` autour des champs natifs. Le boxing de leurs champs `Int` est une piste pour la hausse d’allocations, dont la contribution n’a pas été isolée. Le prochain groupe à examiner est leur émission native à partir du TAST/PBO, suivie de la même comparaison de temps, allocations et GC. Aucun gain de cette étape future n’est présumé.

## Changement mesuré

`AdtKernel.prepareUnary` admet un layout local fermé à partir des `dataDecls`, puis sélectionne les fonctions à un argument ADT directement récursif dont les annotations source/PBO, le résultat et le corps complet sont pris en charge. Les constructeurs natifs sont accompagnés de wrappers publics `obj`. Une fonction récursive remplacée conserve aussi son entrée `_tco` pour les appels du générateur générique.

Sur les 301 modules du benchmark, les seules fonctions sélectionnées sont `Test.RBTree.depth` et `Test.RBTree.makeBlack`. Le layout de `Tree` devient `Color * Tree * int * Tree`. `ins` et `balance` restent génériques. Parmi les 355 fichiers de build, seul `Test.RBTree.fs` change ; le prélude, le harness officiel, les autres modules et les projets sont identiques. Voir `selection.json`, `inputs.json` et `Test.RBTree.fs.diff`.

Les appels saturés utilisent directement la signature native du constructeur, qui fixe les types des `unbox`. Les applications partielles et les constructeurs passés comme valeurs gardent les wrappers curriés. La sélection ne contient aucun nom de benchmark.

## Validation

- `test-adt-unary.log` : 11 contrôles du générateur et 63 assertions F# sur une fixture compilée avec le fork TAST et PBO. Couverture du mélange natif/générique, du pont `_tco`, des refus conservateurs et des échanges entre deux modules.
- `test-adt-interop.log` : 13 contrôles du registre et 62 assertions F#, notamment ordre d’évaluation, applications partielles réutilisables, patterns imbriqués et partage physique des sous-arbres.
- `verify-tree.log` : 27 087 vérifications sur la vraie DLL finale. Ordre BST, racine noire, absence de deux rouges consécutifs, hauteurs noires égales, contenu trié et sans doublons, profondeur native/publique/`_tco`, persistance et partage. Sept familles d’entrées couvrent notamment insertions croissantes/décroissantes, doublons et bornes Int32.
- Builds finaux Release `after` et `diagnostic-after` : aucune erreur ni avertissement. Les sources du backend et son bundle exécuté sont identifiés dans `backend-inputs.json`.
- Le répertoire généré normal a ensuite été reconstruit en Release, sans erreur ni avertissement. Ses 355 entrées sont identiques à celles de la variante mesurée et sa configuration runtime concorde : `build-normal.log` et `normal-build-verification.json`. Aucun benchmark supplémentaire n’a été mêlé à la série.

## Première expérience et correction

La première paire isolée a détecté une régression : total **292,84 → 351,97 ms**, RBTree **213,56 → 272,62 ms**. Les constructions saturées passaient par quatre appels curriés. Cette variante est conservée dans `curried-after/`, avec ses sources, sa DLL, ses empreintes et les logs `curried-probe-*`. La série complète a été arrêtée après cette paire.

Le générateur a ensuite été corrigé pour appeler directement les constructeurs saturés. Les deux suites ADT et les 27 087 vérifications de l’arbre ont été rejouées après ce correctif. Les mesures finales portent sur cette version corrigée.

## Protocole et reproduction

- `before/` fige les 355 fichiers de build générés avant intégration ; `before-inputs.json` en garde les SHA-256. Les trois métadonnées CoreFn sont également conservées, avec `before-metadata-inputs.json`.
- `README.baseline.md` est le README officiel altbak.pub figé : total F#/C# 294,80 ms ; RBTree 213 760 µs.
- `dump-rbt.mjs` observe les vrais `dataDecls` et IR PBO de `Test.RBTree`, sans génération ni exécution du benchmark.
- Après génération, exécuter `python3 prepare.py`, puis compiler Release les copies `after`, `diagnostic-before`, `diagnostic-after`. `before` est déjà compilé.
- Quand aucune compilation ou autre charge de cette tâche ne tourne, lancer `python3 run.py`, puis `python3 analyze.py`.
- Mesure officielle : 5 processus par variante, ordre ABBAABBAAB, vrai `App.main`, 14 sorties vérifiées à chaque passage et meilleurs temps sur 10.
- Diagnostic séparé : vrai `Test.RBTree.act`, 100 000 insertions ; résultat attendu `22`. Trois `App.warmup`, trois actions de chauffe, puis 10 mesures dans chacun des 3 processus par variante, ordre ABBAAB. Allocations du thread et collections GC du processus, validation hors chronométrage ; pile 1 Gio, GC workstation.

Le build initial `before` a réussi sous .NET SDK 8.0.423/runtime 8.0.29 en 38,15 s. Il a émis deux avertissements NU1900 (service d’audit NuGet inaccessible) et aucune erreur. Les reconstructions suivantes utilisent `-p:NuGetAudit=false`. Les DLL `before` et `diagnostic-before` ont été conservées à l’identique pendant la correction : `diagnostic-before-preserved.json` contrôle leurs entrées et l’assembly diagnostique.
