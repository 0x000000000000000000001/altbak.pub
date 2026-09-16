# Arène exclusive : variante expérimentale de RBTree

Source de référence : `/Users/0x1/Documents/htdocs/altbak.pub-purust/output/purust_output/Purs_Test_RBTree/src/lib.rs`.

SHA-256 : `f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a`.

`prepare.py` extrait les déclarations et fonctions autonomes du Rust généré ; seule leur qualification `crate::` devient `crate::generated::` pour permettre une comparaison dans le même exécutable. Le programme de référence reste donc celui de l'output demandé.

## Transformation

`kernel.rs` est une réécriture expérimentale du même algorithme Okasaki : même descente selon les clés, même priorité des quatre rotations, même couleur des nœuds et même forme finale. Les pointeurs `Rc` et les cellules individuelles deviennent un `Vec<Node>` ; zéro représente le constructeur vide, les autres indices désignent des nœuds. Il n'y a pas de spécialisation aux entrées descendantes. Toute l'implémentation est du Rust sûr, avec les vérifications ordinaires d'index et de conversion vers u32.

Cela combine représentation, stockage dans une région, accès mutables et sélection directe des rotations. Un gain contre le généré ne peut donc pas être attribué uniquement à la suppression des allocations, ni directement à une métadonnée déjà disponible.

API de benchmark : `pub fn run(n: i64) -> i64`, `#[inline(never)]`. Allocation, insertion n..1, profondeur et destruction du Vec sont dans `run`. `black_box(tree)` empêche de remplacer le calcul de la profondeur par une connaissance de la construction.

Configurations à compiler avec les mêmes options :

- défaut : indices u32, capacité `max(n,0)+1`, 24 octets par nœud ; une allocation Vec initiale, aucune croissance pour ce programme ;
- `--cfg arena_grow` : indices u32, croissance amortie normale du Vec, sans connaissance de cardinalité ;
- `--cfg arena_usize` : indices usize, témoin pour la taille des liens ;
- les deux cfg sont combinables.

La réservation est un upper bound du nombre de clés insérées, pas une connaissance des formes ou du résultat. Sa durée et sa libération restent chronométrées. Le gain du défaut par rapport à `arena_grow` isole approximativement l'intérêt de cette borne, sans séparer tous les effets de placement mémoire de l'allocateur.

## Conditions pour une application automatique

Le TAST devrait prouver une région fermée : toutes les cellules restent internes jusqu'à production de l'entier final, aucune adresse n'est observée/retournée par du FFI, pas de Weak, aucun consommateur externe conservant une version antérieure. L'analyse devrait également prouver l'exclusivité des cellules lors des écritures et leur durée de vie commune. La destruction groupée suppose des champs sans destructeur observable — ici entiers et couleur.

La version qui réserve ajoute une borne sur les allocations distinctes : au plus un nouveau nœud par insertion, elle-même appelée au plus max(n,0) fois. La version à croissance amortie n'a pas besoin de cette borne. Les indices u32 gardent un contrôle d'overflow ; une borne statique pourrait supprimer ce contrôle dans une spécialisation bornée.

Ce sont des obligations de preuve suggérées, pas des preuves implémentées dans Haskell. La transformation ne préserve pas une API publique d'arbres persistants permettant de conserver des snapshots. Elle spécialise exclusivement le calcul fermé `depth (buildTree n E)`.

La comparaison se limite à `n >= 0` : le comportement de `buildTree` généré sur une entrée négative n'est pas reproduit par la boucle Rust. Pour une génération en production, il faudrait conserver ce domaine par une précondition prouvée ou une garde avec fallback. De même, un arbre trop grand pour les indices u32 nécessite une borne prouvée ou un fallback à une représentation plus large ; le prototype conserve une conversion vérifiée lors de chaque allocation, qui panique si la capacité des indices est dépassée.

## Variante isolant les contrôles de bornes

`prepare_unchecked.py` dérive mécaniquement `kernel-unchecked.rs` du noyau sûr. Les seuls changements portent sur le stockage : un wrapper `NodeStore` implémente `Index` et `IndexMut` via `get_unchecked` / `get_unchecked_mut`. Les corps des méthodes d'insertion, rotation et parcours sont identiques. La conversion checked vers u32 lors des ajouts reste présente.

L'invariant nécessaire est inductif : zéro est toujours le nœud sentinelle ; tout nouvel indice désigne un élément qui vient d'être ajouté ; les liens proviennent exclusivement de ces indices ; aucune opération ne retire de nœud ou ne raccourcit le Vec. La réallocation ne pose pas de problème puisque seuls les indices sont conservés, jamais des pointeurs à travers un ajout. Les emprunts Rust ordinaires empêchent les références mutables simultanées. Il s'agit d'une propriété interne du futur backend d'arène ; ajouter un booléen au TAST ne la rendrait pas vraie automatiquement.

`--cfg arena_verify` réactive une assertion de bornes à chaque accès, dans le même wrapper. `check-unchecked.rs` reprend toute la comparaison de formes, couleurs et invariants ; `validation-unchecked-verified.tsv` consigne le résultat avec ces assertions, et `validation-unchecked.tsv` celui du binaire sans assertions. Aucun changement d'algorithme n'est associé à cette variante.

## Équilibrage spécialisé selon l'enfant modifié

`prepare_directional.py` produit `kernel-directional.rs` (indices vérifiés) et `kernel-unchecked-directional-inline.rs` (composition avec les indices sans vérification et les indications d'inlining du témoin correspondant). La décision d'équilibrage conserve exactement les quatre rotations et leur réécriture de couleurs. Après insertion à gauche, elle recherche seulement LL/LR ; après insertion à droite, seulement RL/RR. La direction est un paramètre constant générique, donc le programme ne doit plus la tester à l'exécution.

La justification est structurelle : avant l'insertion, chaque sous-arbre satisfait l'absence de deux nœuds rouges adjacents ; seule la branche modifiée peut enfreindre cette propriété. Un parent noir n'a donc pas besoin de rechercher une violation dans l'autre branche. Cela ne dépend ni de l'ordre des clés ni de leur nombre. Il faut conserver l'invariant comme précondition de la spécialisation : la fonction générique balance appliquée à un arbre arbitrairement invalide ne peut pas être remplacée sans autre preuve. Une analyse Haskell correspondante serait plus ambitieuse qu'un simple compteur d'usages, puisqu'elle doit préserver un invariant récursif entre les états avant/après une fonction.

## Validation

`check.rs` compare l'intégralité du préordre (constructeurs vides compris), les clés et les couleurs avec le généré. Il vérifie séparément ordre BST, racine noire, enfants noirs des nœuds rouges, hauteurs noires égales, cardinalité et profondeur. Cas : quatre rotations minimales, vide, 100 000 clés ascendantes, 100 000 descendantes, 100 000 mélangées (seed fixe), 100 000 insertions contenant des doublons (4 096 clés distinctes).

```sh
python3 prepare.py
rustc --edition=2021 -C opt-level=3 check.rs -o check-u32
./check-u32
rustc --edition=2021 -C opt-level=3 --cfg arena_grow --cfg arena_usize check.rs -o check-usize-grow
./check-usize-grow
```

Les sorties sont conservées dans `validation-u32.tsv` et `validation-usize-grow.tsv`. Aucun benchmark concurrent n'a été lancé par ce sous-agent ; les mesures sont centralisées par l'agent parent.
