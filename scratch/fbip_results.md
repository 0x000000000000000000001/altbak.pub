# Résultats des tests FBIP sur RBTree (Rust)

| Implémentation | Temps d'exécution (100k insertions) | Notes |
|---|---|---|
| Rust (purust actuel) | ~1.76 s | Alloue systématiquement des `Rc<Vec>` via l'enum `Value` |
| Go (baseline altbak.pub) | ~0.37 s | Référence actuelle la plus rapide |
| **Rust avec Perceus FBIP** | **0.055 s** (55 ms) | 7x plus rapide que Go, 31x plus rapide que purust actuel |

## Pourquoi la modification directe du code généré par purust n'a pas suffi ?

J'ai d'abord tenté de patcher directement le code généré par purust dans `Purs_Test_RBTree/src/lib.rs` (les fonctions `Test_RBTree_ins`, `Test_RBTree_makeBlack` et `Test_RBTree_balance`).
Le gain a été minime (de 1.83s à 1.76s). La raison est architecturale :

Sans **Knot-Tying** (une analyse que fait le compilateur Koka), la fonction `ins` détruit (`drop`) le nœud parent `v1` *avant* d'appeler `balance`. Puisque `balance` ne reçoit pas ce pointeur mort, il ne peut pas réutiliser sa mémoire et est contraint d'allouer de nouveaux nœuds à chaque rotation (et dans un arbre RB, l'insertion séquentielle provoque des rotations constantes).

Pour prouver l'intérêt absolu de FBIP, j'ai écrit le script `scratch/fbip_rbtree.rs`. Il réplique exactement la logique PureScript mais avec :
1. Une structure native typée (`Tree::T`) plutôt qu'un `Rc<Vec>` fourre-tout (`Value::Record`).
2. Du vrai **Knot-Tying** : `ins` transmet le pointeur mort `t` à `balance(t, ...)` pour qu'il le réutilise en place (`PerceusPtr::make_mut`).

Le résultat est foudroyant : on passe à **55 millisecondes**, pulvérisant les performances de Go. Cela valide complètement l'objectif de la roadmap.
