# Records : les clones temporaires empêchent la réutilisation

Audit du 9 septembre 2026, sans modification du compilateur, des FFI ou des sources du benchmark. Deux prototypes Rust isolés ont été exécutés ; aucune règle nouvelle n'est intégrée.

## Comparaison avec le README

Le [README officiel](../../../altbak.pub/README.md#rust), relu pour cet audit, donne :

| Benchmark | Rust compilé | Natif optimisé, dernière colonne | Lecture pour la suite |
| --- | ---: | ---: | --- |
| RBTree | 18,449 ms | 36,070 ms | Déjà plus rapide ; examiner le travail résiduel du compilé indépendamment du natif. |
| Deep Record Updates | 0,969 ms | 0,004 ms | Écart important, avec des copies évitables effectivement mesurées ci-dessous. |
| Church | 0,173 ms | 0,001 ms | La fusion récente a réduit ce coût ; le budget résiduel est plus petit que celui de Records. |
| Total | 19,95 ms | 36,13 ms | RBTree représente environ 92,5 % du total. |

Le répertoire demandé `altbak.pub/run/bak/rust/output/purescript` n'existe pas : le Rust se trouve dans `purust_output`. La sortie du checkout normal date du **8 septembre** et conserve notamment un `Unsupported Op2` dans Records. L'audit et les expériences utilisent donc la sortie du **9 septembre** de `altbak.pub-purust`, qui inclut les intégrations récentes. Les sources PureScript/FFI et le README normal ont été lus, sans être modifiés.

## Goulot constaté

Dans [Test_Records_updateRec](../../run/bak/rust/output/purust_output/Purs_Test_Records/src/lib.rs), la mise à jour commence par :

```rust
let mut _base = purs_local_1.clone();
_base.set_a(mk_int(purs_local_1.get_a().unwrap_int() + 1));
// Puis calcul/mise à jour des deux niveaux imbriqués à partir de purs_local_1.
```

Le clone conserve deux propriétaires de la racine pendant `set_a`. Le setter utilise `PerceusPtr::make_mut`, qui copie le record s'il n'est pas unique. Les niveaux `b` et `b.d` sont eux aussi partagés pendant leurs mises à jour. Le comptage réel retrouve exactement **3 nouvelles cellules par itération**, en plus des 3 cellules initiales : **30 003 allocations et 30 003 libérations pour 10 000 itérations**.

Le point d'émission est la branche `Update base props` de [CodeGen](../../../purust/purust/src/Purust/CodeGen.purs). Elle génère la base avant les valeurs des champs ; `aliveForBase` inclut les références présentes dans ces valeurs, ce qui impose ici le clone. Le runtime sait déjà réutiliser une cellule unique : c'est l'ordre des opérations et la durée de vie des propriétaires générés qui empêchent cette réutilisation.

Les champs restent `Option<Value>`, avec des enfants sous `PerceusPtr`. `mk_int` construit un `Value::Int(i64)` immédiat : **les 30 000 allocations supplémentaires ne sont pas une allocation par entier**. Les structures et les types profonds sont disponibles dans le TAST : `typeTable[0]` décrit `Int -> DeepRecord -> DeepRecord`, et les entrées 2–7 décrivent les trois records fermés et leurs quatre champs `Int`. Le `UnknownType` émis correspond à un choix de représentation du backend, pas à une absence de types dans le TAST.

Le [natif](../../../altbak.pub/src/Test/RecordsFFICheatcode.rs) utilise quatre variables `i64` et ne renvoie que `f`. Les autres accumulations peuvent disparaître à l'optimisation. Son chiffre de 4 µs n'est donc pas une prévision pour un worker qui conserve un résultat record complet et sa persistance.

## Expériences isolées

La variante « racine déplacée » calcule les valeurs des champs dans leur ordre, puis déplace la racine au dernier usage au lieu de la cloner avant ces calculs. Les deux niveaux intérieurs gardent leurs setters actuels.

La variante « chemin réutilisé » calcule d'abord les quatre valeurs scalaires, puis utilise successivement `PerceusPtr::make_mut` sur la racine, `b` et `b.d`. Les chemins partagés restent copiés par le runtime. Elle conserve les mêmes `Value`, `Option<Value>`, structures et pointeurs ; elle n'introduit pas de records unboxés, de pool ou de calcul fermé du résultat.

| Noyau, 10 000 itérations | Temps médian | Gain local | Allocations / libérations |
| --- | ---: | ---: | ---: |
| Rust généré actuel | 0,960 ms | — | 30 003 / 30 003 |
| Déplacer seulement la racine | 0,686 ms | −28,6 % | 20 003 / 20 003 |
| Réutiliser les trois niveaux | 0,360 ms | −62,5 % | 3 / 3 |

Protocole : cinq triplets de processus, ordre alterné, vingt mesures après une chauffe par processus ; médiane des cent échantillons par variante. O1/mimalloc, bibliothèques du build actuel. Tous les binaires sont compilés avant le chronométrage. Les quatre champs du résultat sont observés, et création/destruction sont incluses. Les compteurs d'allocations sont exécutés séparément. Les trois variantes ralentissent au premier triplet, mais les deux prototypes sont favorables dans chacun des cinq triplets ; les données brutes sont conservées.

Les gains correspondent à **0,275 ms** et **0,600 ms** sur ce noyau. Rapportés au total historique de 19,95 ms, cela représenterait approximativement **1,4 %** et **3,0 %** si ces différences se reproduisaient dans la suite. **Aucune mesure du runner complet après intégration n'existe à cette étape.**

Contrôles réalisés :

- Comptage séparé pour 0, 1, 2, 10 et 10 000 itérations ; toutes les allocations sont libérées.
- Les quatre valeurs du record, pour 24 combinaisons de graines/comptages, y compris des records initiaux non nuls.
- Conservation d'une ancienne racine, d'un sous-record et d'une feuille partagés ; leurs valeurs restent inchangées.
- Même représentation et mêmes résultats dans les trois variantes.

Les prototypes ne constituent pas une preuve générale de transformation : une intégration doit encore couvrir l'ordre des callbacks opaques, les débordements/exceptions, les références retenues dans les champs de remplacement et les usages ultérieurs. Le cas testé n'a que des mises à jour numériques pures et des compteurs non négatifs.

## Bébé recommandé et pastilles

**Commencer par la racine déplacée.** Dans `Update`, reconnaître une base locale à son dernier usage, évaluer les valeurs de remplacement dans l'ordre puis déplacer cette base. Garder le chemin existant lorsque la preuve manque. Cela supprime déjà **10 000 copies** dans le prototype, sans changement de layout.

1. Ajouter une fixture TAST fraîche : racines uniques/partagées, anciens résultats conservés, base encore utilisée après la mise à jour, champ qui retient la base, callbacks et exceptions.
2. Intégrer cette seule règle locale ; build, tests de génération/TAST et `bin/rust/run -c` dans le worktree.
3. Confirmer sur cinq paires de runners complets et recompter les allocations réellement générées.
4. Traiter ensuite les enfants imbriqués avec une preuve d'accès exclusif au chemin ; le prototype à 3 allocations fournit une cible à vérifier, pas une intégration déjà acquise.

La table est cohérente avec ce diagnostic : **B14 est partiel** pour les transferts au dernier usage et les emprunts ; cette règle en étendrait la couverture. **B22 est déjà vert** pour les setters de records sans copie générique par clé : leur présence ne garantit pas l'unicité de la base. **B21/B23** désignent un chantier ultérieur de layouts/valeurs strictement typés, dont les prototypes ci-dessus n'ont pas besoin. La spécialisation d'un champ d'ADT en **B13** est un autre chemin de génération.

RBTree demeure prioritaire en budget absolu. L'extension aux enfants modifiés par un appel et le parcours `depth` emprunté restent à mesurer. Les expériences précédentes sur les paires clone/drop et les preuves d'unicité n'ont pas établi de gain suffisant ; elles ne justifient pas de promettre une amélioration de RBTree sur la seule lecture des `.clone()`. **Records est la prochaine piste disposant ici d'un gain local positif mesuré.**

## Artefacts

[Prototype reproductible](probe.py), [source générée analysée](generated-before.rs), [temps bruts](time.json), [allocations](count.json), [contrôles de persistance](check.json), [révision/empreintes/TAST](metadata.json). Les binaires et variantes Rust sont conservés sous `build/`, ignoré par Git.

Depuis ce dossier : `python3 probe.py count`, `python3 probe.py check`, puis `python3 probe.py time` une fois les autres compilations terminées. La source avant est figée ; les bibliothèques de `purust_output/target/release/deps` doivent correspondre aux types de cette source. Le script échoue en cas d'ambiguïté sur les bibliothèques.
