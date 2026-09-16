# Preuve par ablation des métadonnées actuelles — 16 septembre 2026

## Résultat principal

Sur le noyau généré de `Test.RBTree`, les champs TAST `usageCount` et
`escapes` actuels n'apportent **aucune différence de code Rust**. Leur retrait
dans toutes les entrées TAST reproduit le fichier RBTree octet pour octet.
Ce résultat prouve une contribution directe nulle de ces deux champs au
noyau actuel ; il ne prouve pas que toute analyse de consommation serait
inutile, ni qu'une exploitation future de ces informations serait impossible.

Le noyau obtenu des deux côtés est aussi identique au noyau utilisé dans le
benchmark frais précédent. SHA-256 :
`f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a`.

## Expérience reproductible

Exécuter `python3 ablate_usage.py` depuis ce dossier ou par son chemin absolu.
Le script ne modifie pas le compilateur, PBO, Purust ni le benchmark source.
Il copie les 302 `corefn.json` du benchmark frais dans deux dossiers isolés.
Les `modulePath` sont absolutisés de la même manière des deux côtés pour
retrouver les sources et FFI ; seule l'ablation de `usageCount` et `escapes`
diffère entre les deux entrées.

Il retire exactement **26 658 occurrences de chacun des deux champs**,
puis exécute le même bundle Purust existant deux fois, avec caches `.purmeta`
isolés et sorties Rust isolées. Bundle SHA-256 :
`4de4f7da2f8931fe9b4461bc1133a73f1277b7bd23ff8dae4f60998204eb7cd1`.

Les SHA des entrées, commandes, répertoires de travail et sorties sont dans
`ablation-results.json`. Aucune compilation Rust ni mesure d'exécution n'est
nécessaire pour constater l'identité exacte de ce noyau. Les temps de
génération enregistrés sont des informations de diagnostic, sans conclusion
de performance faute de protocole répété.

## Portée précise de l'identité

- 612 fichiers générés, dont 307 Rust.
- 610 fichiers identiques, dont **305 fichiers Rust identiques**.
- `Purs_Test_RBTree/src/lib.rs` et les fonctions `ins`, `balance`, `insert`,
  `buildTree`, `makeBlack`, `depth` sont identiques.
- `Purs_Data_Enum/src/lib.rs` diffère uniquement par deux commentaires `Typed`.
- `Purs_Data_Array/src/lib.rs` diffère réellement dans `intersperse` :
  réassociation de blocs `EffectBind` et d'une capture. Ce n'est pas un code
  exécuté par l'algorithme RBTree. La différence montre toutefois que les
  wrappers source peuvent encore influencer une transformation avant leur
  invalidation ; on ne peut pas affirmer une neutralité totale du pipeline.

Les deux diff sont conservés dans `Purs_Data_Array.diff` et
`Purs_Data_Enum.diff`. Aucun gain général à d'autres programmes n'a été mesuré
ou revendiqué.

## Ce que contient réellement Test.RBTree

Le détail des 135 annotations est dans `metadata.json` avec fonction, nature
du nœud, valeur, intervalle source et chemin JSON.

| Nature | Nombre |
| --- | ---: |
| Liaison NonRec | 10 |
| Liaison récursive | 3 |
| Paramètre Abs | 15 |
| Variable de motif VarBinder | 46 |
| Occurrence Var | 61 |

| Couple usageCount / escapes | Nombre |
| --- | ---: |
| 1 / true | 101 |
| 1 / false | 10 |
| 0 / false | 13 |
| -1 / true | 7 |
| 2 / true | 2 |
| 3 / true | 2 |

Les dix `(1, false)` ne représentent pas dix opérations indépendantes à
optimiser : ce sont **cinq paramètres et leurs cinq occurrences**, employés
comme scrutinees de `case` dans `makeBlack`, `depth`, `balance` (dernier
paramètre), `ins` (arbre) et `buildTree` (accumulateur).

Les enfants `a` et `b` de `ins` et de `depth` portent `(1, true)`. L'analyse
actuelle marque les arguments d'un appel comme échappants de façon
conservatrice (`Usage.hs`, lignes 112–115). La simple conjonction
`usageCount == 1 && escapes == false` ne sélectionne donc pas les clones
temporaires d'enfants qui restent dans les chemins chauds.

## Pourquoi les hints n'améliorent pas ce noyau aujourd'hui

- `Semantics.purs:392` invalide `UsageMeta` lors de l'évaluation PBO, car les
  transformations peuvent changer les usages de la syntaxe source.
- `CodeGen.purs:99` nettoie récursivement ces hints avant les analyses Rust.
- `CodeGen.purs:2403` décide déplacement ou clone à partir de l'ensemble
  `alive` du programme transformé. Les déplacements sûrs des variables
  mortes sont donc déjà possibles sans les deux champs source.

Enfin, dernière utilisation d'une variable et absence d'échappement local
ne prouvent pas l'unicité du tas. `depth` a un paramètre `(1, false)` même si
son appelant conserve une autre référence au même arbre. Des enfants peuvent
aussi être partagés entre eux. Ces deux champs ne justifient donc pas, seuls,
de retirer les contrôles `Rc::get_mut`, de déplacer un champ hors d'une cellule
partagée ou de remplacer globalement `Rc` par `Box`.

Pour les clones de la descente, une optimisation correcte doit notamment
spécifier comment la référence est extraite du parent, comment le parent
est maintenu valide durant l'appel récursif et quoi faire en présence de
partage. Pour `depth`, un résumé de paramètre purement lu permettrait un
worker par emprunt ; ce résumé n'est pas le booléen `escapes` actuel.

## Lien avec les chiffres de performance

La baseline officielle du README est **8 788,42 µs** et la cible de −50 %
**4 394,21 µs**. Le rapport frais précédent mesure une médiane de
**8 536,29 µs** et constate 100 001 allocations, 2 483 932 succès sur autant
d'appels `get_mut`, 2 183 976 clones temporaires durant la descente et 200 000
dans `depth`. Ces nombres sont cités comme résultats antérieurs, pas comme
de nouvelles mesures réalisées par ce script.

L'ablation donne une réponse causale plus précise à la question actuelle :
les bons résultats de réemploi de ce noyau proviennent déjà des analyses et
transformations de Purust, et ne peuvent pas être attribués à la présence
des champs source `usageCount` + `escapes` dans ce build.
