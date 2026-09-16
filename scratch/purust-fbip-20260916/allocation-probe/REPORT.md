# Comptage RBTree après réparation des métadonnées

Le noyau fraîchement généré construit 100 000 clés avec **100 001 allocations de
cellules** et réussit **les 2 483 932 vérifications `Rc::get_mut`**. Aucun refus
pour partage ou référence faible ne survient pendant ce scénario unique.

| Fonction / chemin | Appels | Allocations | `get_mut` réussis | Clones de pointeurs |
| --- | ---: | ---: | ---: | ---: |
| `ins` | 2 283 976 | 100 000 | 2 183 976 | 2 283 976 |
| `insert`, recoloration incluse | 100 000 | 0 | 100 000 | 0 |
| Rotation par permutation des champs | 99 978 | 0 | 199 956 | 0 |
| Garde de reconstruction après récursion | 1 468 946 | 0 | 0 | 0 |
| Racine vide initiale du harnais | — | 1 | 0 | 0 |
| `depth`, phase suivante | 200 001 | 0 | 0 | 200 000 |

Les workers généraux `balance` et `balance__purust_reuse` ne sont jamais appelés
pendant les insertions décroissantes uniques. Le chemin spécialisé met les
champs à jour et effectue les rotations sur place.

Dans `ins`, 100 000 clones partagent le constructeur vide entre les deux enfants
du nouveau nœud. Les 2 183 976 autres clones fournissent un substitut temporaire
au champ enfant déplacé pendant la récursion (`mem::replace(child,
sibling.clone())`). Ils sont suivis d'autant de relâchements sans destruction.

La suppression de clones ne peut pas améliorer le taux de réemploi ou le nombre
d'allocations de ce scénario : ils sont déjà au niveau attendu. Le coût restant
porte notamment sur les visites récursives, les gardes, les tests d'unicité,
les substitutions de champs et le parcours final. Leur contribution temporelle
demande une mesure séparée. Ces compteurs ne prouvent aucun gain de temps.

Les essais historiques isolés de septembre 2026 comparant extraction complète
et modification de champ favorisaient cette dernière malgré son clone
temporaire : `red_field` 15,20 ms contre `red_rebuild` 17,69 ms, et `black_field`
10,70 ms contre `black_rebuild` 11,82 ms. Ces anciens noyaux ne remplacent pas la
baseline officielle actuelle et ne doivent pas être comparés directement au
nouveau runner.

## Vérifications

Le harnais vérifie 100 000 nœuds, une profondeur de 22, l'ordre des clés, les
invariants rouge/noir, les quatre rotations, les références faibles et les clés
exactes de 200 anciennes versions conservées simultanément. Les bilans globaux
des références et des destructions sont également équilibrés.

Le wrapper de comptage conserve la taille d'un pointeur et appelle le vrai
`std::rc::Rc`. Il enregistre les opérations logiques, pas les instructions
machine ; ses propres allocations de journal sont exclues. Il n'est jamais
utilisé pour chronométrer.

Exécution depuis `altbak.pub` :

```sh
python3 scratch/purust-fbip-20260916/allocation-probe/probe.py
```

La source originale reste inchangée. `results.json` enregistre son empreinte,
les phases, les fonctions et les sites ; `events.tsv` contient les compteurs
bruts. Le noyau instrumenté conserve les numéros de ligne du Rust émis.
