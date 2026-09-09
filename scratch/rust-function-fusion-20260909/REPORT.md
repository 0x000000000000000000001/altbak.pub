# Fusion d'un producteur de fonctions avec son application

9 septembre 2026. Intégré dans Purust, validé depuis `altbak.pub-purust`. Les checkouts normaux d'altbak.pub et de PBO sont restés intacts.

## Résultat

Cinq paires alternées de runners complets, O1/mimalloc, sans compilation ni instrumentation pendant les mesures. Chaque processus conserve le protocole du runner : échauffement, meilleur de dix exécutions pour chaque cas. Le tableau donne les médianes des cinq processus ; le total additionne les médianes par benchmark.

| Mesure | Avant | Après | Évolution |
| --- | ---: | ---: | ---: |
| Church | 1,439 ms | 0,164 ms | −88,6 % |
| Total des 14 benchmarks | 21,053 ms | 19,691 ms | −6,5 % |
| Allocations de Church, compteur séparé | 133 641 | 309 | −99,77 % |
| Libérations de Church, compteur séparé | 133 641 | 309 | −99,77 % |

Les cinq paires améliorent Church et le total. Le total par processus varie de 20,665 à 21,331 ms avant, de 19,485 à 19,816 ms après. Les autres variations ne sont pas attribuées à la fusion : parmi les **302 sources Rust générées**, seul `Purs_Test_Church/src/lib.rs` change. RBTree reste identique et donne 18,352 → 18,255 ms dans cette série.

Le [README officiel](../../../altbak.pub/README.md#rust), relu avant la mesure, affiche Church **1,506 ms** en Rust compilé et **≈ 0,001 ms** dans la dernière colonne native ; les totaux sont **21,82 / 36,13 ms**. Ces références historiques restent distinctes de la paire avant/après de cette expérience. Le natif de Church utilise des boucles arithmétiques sans callbacks ; des appels indirects et 309 allocations restent ici. Le total de 20,354 ms de la session précédente de recoloration n'est pas la baseline de cette nouvelle session.

## Coût supprimé et règle

Le wrapper natif saturé de `fromInt(n, f, x)` appelait d'abord un producteur récursif retournant une chaîne de closures, puis appliquait cette chaîne à `f` et `x`. Les combinateurs de Church reconstruisaient cette chaîne de nombreuses fois. Un premier prototype isolé donne **1,479 → 0,188 ms** ; le comptage réel donne les allocations ci-dessus, ce qui a justifié l'intégration.

`Purust.FunctionFusion.countedFunctionProducers` reconnaît exclusivement ce motif dans une déclaration récursive singleton :

```purescript
build 0 = identity
build n = let previous = build (n - 1) in \f x -> f (previous f x)
```

La signature doit être explicitement annotée `Int -> (Int -> Int) -> Int -> Int`. Les annotations TAST intermédiaires doivent rester compatibles ; une conversion, un appel opaque, une capture du compteur dans le callback, un décrément autre que 1, une base non identité ou une récursion mutuelle empêchent la fusion. L'identité est une lambda vérifiée ou une déclaration locale vérifiée, éventuellement polymorphe. Les noms ne participent pas à la reconnaissance.

Dans le wrapper déjà saturé, le générateur émet :

```rust
if n >= 0 {
    let mut count = n;
    let mut result = x;
    while count > 0 {
        result = f(result);
        count -= 1;
    }
    result
} else {
    // Corps original : construire la fonction, puis l'appliquer.
}
```

Le compteur positif ne peut pas déborder pendant la décrémentation. L'arithmétique du callback, son nombre d'appels et leur ordre sont conservés. La branche négative garde l'ancien producteur, avant toute application du callback. La signature publique et les wrappers des applications partielles restent identiques. Cette règle étend la spécialisation des fonctions d'ordre supérieur et les wrappers/workers (**B3/B4**, toujours partiels) ; elle n'ajoute pas de monomorphisation générale ni de mécanisme Perceus.

## Vérifications

- Construction de Purust et de l'application : zéro erreur, zéro avertissement.
- **20 tests de génération et 12 tests TAST réussis.** Le fork TAST est sélectionné explicitement avec `PURS`.
- La fixture fraîche utilise des noms différents de Church ; vérification du Rust, compilation avec les contrôles de débordement activés, puis exécution.
- Résultats et traces de callbacks non commutatifs, compteur nul, fonctions sauvegardées puis appliquées plusieurs fois, applications partielles capturant un callback, appels intermodules et libération des captures.
- Arrêt au même callback en cas de panic, débordement dans le callback conservé ; l'entrée FFI `i64::MIN` prend le chemin négatif et déborde au premier décrément sans appeler le callback.
- Motifs voisins non transformés : décrément de 2, base non identité, compteur capturé, conversions et appels opaques.
- `bin/rust/run -c` terminé avec succès ; **14 résultats vérifiés** dans chacun des dix runners appariés et dans le run propre.
- Sur le noyau Church réellement régénéré, les compteurs et contrôles reproduisent le prototype. Toutes les allocations instrumentées sont libérées, pour les entrées 0, 1, 2, 3 et 10.

## Preuves et reproduction

- [Révisions et empreintes initiales](metadata.json), [comparaison des sources](source-comparison.json), [diff Rust](Church.diff).
- [Temps appariés complets et empreintes des exécutables](runner-results.json), [prototype isolé](time.json), [comptage intégré](count-integrated.json), [contrôles intégrés](check-integrated.json), [validation](validation.json).
- `probe.py time` compile d'abord les deux variantes, puis mesure cinq paires de vingt échantillons après échauffement. `probe.py count --integrated` et `probe.py check --integrated` vérifient aussi le code courant réellement généré.
- `measure-runner.py` compare les exécutables `build/runner-before` et `build/runner-after`, préservés avant/après `bin/rust/run -c`. Les exécutables, sorties de build et journaux volumineux sont ignorés par Git ; les résultats JSON, scripts et diff sont conservés.

Baseline Purust : `089eef3088b5353ff5e10e6ad36e08fab41e058e`. Les sources et exécutables de travail sont conservés dans `build/` pour cette session ; repartir de cette révision et des dépendances indiquées dans `metadata.json` pour une nouvelle reproduction. Les temps ne sont pas garantis d'une session à l'autre.
