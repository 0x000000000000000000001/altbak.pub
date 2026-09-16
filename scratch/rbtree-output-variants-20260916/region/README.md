# Allocation en région, noyau généré inchangé

Cette expérience ne change ni les rotations, ni les appels à `Rc::get_mut`, ni les
clones, ni le layout de `Tree` ou de `Rc`. `kernel.rs` est une extraction fidèle des
fonctions générées depuis l'output demandé. Le manifeste donne les SHA256.

`region_alloc.rs` remplace seulement l'allocateur global pendant le noyau fermé.
Chaque `run(n)` :

1. alloue un bloc neuf via `System`, de `(n + 1) * 64` octets ;
2. active l'allocation par déplacement d'un curseur dans ce bloc ;
3. appelle exactement `buildTree(n, Rc::new(E))`, puis `depth(tree)` ;
4. exécute tous les drops et destructions récursives ordinaires ;
5. libère le bloc avec `System` avant de retourner l'entier.

Toute cette fonction doit être chronométrée. Il n'y a pas de buffer persistant,
de réserve allouée avant le chrono, de destruction oubliée, ni de mémoire confiée
à la fin du processus. Les dealloc individuels vers l'intérieur du bloc ne font
rien ; ceux du fallback système continuent d'être libérés individuellement.
La capacité n'est pas une condition de correction : un débordement revient à
`System`. Les alignements demandés sont respectés.

Le prototype utilise des atomiques pour le curseur et la plage de mémoire. Leur
coût reste inclus. La portée expérimentale est un programme mono-thread, sans
allocation retenue au-delà de la région. Cette restriction appartient au contrat
de l'API unsafe `Region::new` ; ce n'est pas un allocateur général réutilisable sans
analyse supplémentaire. Aucune preuve statique de ce contrat n'est implémentée.

## Reproduction

`python3 prepare.py` extrait le noyau, compile avec `rustc --edition=2021 -C opt-level=3`
et lance uniquement les validations, sans mesure de performance.

Pour le harness de mesure partagé :

```rust
#![allow(warnings)]
include!("kernel.rs");
mod region_alloc;
// chronométrer tout l'appel :
// let depth = region_alloc::run(std::hint::black_box(100000));
```

## Validation obtenue

- Forme exacte, couleurs et clés identiques au noyau original avec allocateur
  système : 100 000 insertions croissantes, décroissantes et mélangées ; doublons.
- Invariants rouge/noir, ordre strict des clés, racine noire, quatre rotations.
- 4 800 048 octets occupés pour 100 000 clés ; aucun fallback à capacité normale.
- Débordement volontaire : retour système correct, arbre identique.
- Quatre exécutions consécutives avec région fraîche ; profondeur 22.
- 200 versions persistantes retenues simultanément ; références `Weak`.
- Alignements d'allocation de 1 à 4096 octets.

Toutes les références fortes/faibles et les vecteurs temporaires de validation
sont détruits avant la région. Les arbres de référence sont alloués avant son
activation et détruits après sa désactivation.

## Information TAST associée

Une région de durée de vie fermée permettrait de choisir cette stratégie : les
objets construits ne sont plus accessibles après le calcul de l'entier final,
aucun FFI ne conserve leurs adresses, aucun autre thread n'y accède. Il n'est pas
nécessaire de prouver que chaque cellule est unique : les références partagées,
la persistance et les `Weak` peuvent exister **dans** la région, leurs mécanismes
restant inchangés ici. Il faut prouver l'absence de fuite de l'ensemble des
allocations de cette région. Les destructions restent exécutées, donc cette
expérience ne suppose pas leur absence d'effets.

Ce prototype teste la réduction des allocations système et la localité mémoire.
Il ne teste pas encore la suppression des refcounts, la suppression des drops,
les indices compacts, ni un layout différent. Les timings sont orchestrés par le
rapport parent pour éviter des benchmarks concurrents.

## Variante locale sans atomiques

`region_local.rs` fournit la même API avec un curseur dans `UnsafeCell`, sans
atomiques. Le contrat mono-thread fermé devient une précondition de toutes les
opérations de cet exécutable, y compris hors région ; l'implémentation `Sync`
explicite repose sur cette précondition. Aucun callback de signal ou accès
réentrant à l'allocateur n'est admis. Cette variante n'est donc pas une bibliothèque
d'allocateur global sûre pour un programme arbitraire.

`python3 prepare_local.py` répète toutes les validations avec cette variante.
Le fichier `validation-local.txt` confirme les mêmes résultats. Pour choisir
l'allocateur du binaire de production comme allocateur amont, compiler avec
`--cfg upstream_mimalloc` et la dépendance `mimalloc`. Sans ce cfg, l'amont reste
`std::alloc::System`. Allocation fraîche et libération du bloc restent dans `run`.
