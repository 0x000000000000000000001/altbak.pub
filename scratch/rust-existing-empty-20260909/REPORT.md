# RBTree : réemployer la valeur vide déjà testée

Le réemploi de l’entrée vide d’`ins` supprime **100 000 allocations sur 200 001 (−50,0 %)** et améliore le runner complet : **RBTree 21,222 → 19,569 ms (−7,8 %)**, **total 44,380 → 42,923 ms (−3,3 %)** sur cinq paires alternées.

## Preuve puis intégration

La référence fraîche correspond exactement au Rust final de l’étape précédente sur le partage local. Les sorties générées ayant été nettoyées entre les deux étapes, le runner de référence a été reconstruit avant toute modification du générateur.

Le [prototype](RBTree-prototype.rs) change uniquement la liaison de la feuille dans `ins` : la valeur locale dont le tag `E` vient d’être testé remplace `Rc::new(Tree::E)`. Les deux enfants reçoivent cette référence, avec un clone pour le premier et un déplacement pour le dernier. Le prototype passe les invariants et les 200 versions conservées. Trois paires de processus du noyau isolé, quinze mesures après échauffement par processus, donnent **21,612 → 19,905 ms (−7,9 %)** à O1 avec mimalloc. Ces [mesures exploratoires](prototype-time.json) sont distinctes du runner complet.

La règle intégrée reconnaît un **test de tag positif sur une variable locale native empruntée**. Dans sa branche vraie, [ShareNullaries.purs](../../../purust/purust/src/Purust/ShareNullaries.purs) remplace les constructeurs sans champ correspondants par cette variable. L’identité comprend la représentation Rust et le constructeur ; les annotations TAST sont conservées. La réécriture intervient avant le calcul des usages nécessaires dans les branches et les champs, afin de garder les clones requis par les utilisations ultérieures.

La règle parcourt seulement les constructions strictes et les annotations sans changement de représentation. Elle s’arrête aux appels et aux nouvelles portées ; elle n’ajoute aucune capture à un corps différé. Les appels testés directement, les conversions, les tags de constructeurs à données et les branches `else` sont exclus. Une entrée partagée ou faiblement référencée reste réutilisable comme valeur immuable : son contenu n’est pas modifié et aucune racine globale n’est retenue.

Dans le [diff Rust réellement généré](RBTree.diff), `ins` réemploie maintenant l’entrée vide pour ses deux enfants. La même règle simplifie aussi le retour `E` des branches vides de `makeBlack` et `insert`. Le [patch du générateur propre à cette étape](integration.patch) permet de distinguer ce changement du partage local précédent ; ses sources de référence correspondent aux empreintes enregistrées lors de cette précédente étape.

## Allocations et correction

Le [comptage](count.py) porte sur les deux noyaux réellement générés, indépendamment du chronométrage. Il inclut 100 000 insertions, parcours et destruction.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Constructeurs vides `E` alloués | 100 001 | 1 |
| Nœuds à données `T` alloués | 100 000 | 100 000 |
| Allocations / libérations | 200 001 / 200 001 | 100 001 / 100 001 |
| Octets demandés cumulés | 9 600 048 | 4 800 048 |
| Taille de `Tree` | 32 octets | 32 octets |

Toutes les allocations comptées sont libérées. Les octets demandés cumulés ne mesurent pas le pic de mémoire résidente. Sur les deux noyaux : quatre rotations, invariants rouge/noir, ordre des clés, doublons, et clés exactes de **200 versions simultanément conservées**, avant/après insertion et destruction.

La [fixture TAST](../../../purust/purust/tests/tast/known-nullaries.mjs) traverse le vrai fork, PBO, Purust et les crates Rust fraîches. Elle vérifie **une allocation par insertion dans `Empty`**, avec une entrée unique, partagée, faiblement référencée, ou partagée et faiblement référencée. Les adresses des enfants correspondent à celle de l’entrée. Elle couvre aussi les constructions imbriquées, les constructeurs homonymes de modules différents, les autres tags, les callbacks, l’évaluation unique du prédicat calculé et les usages ultérieurs de l’entrée. Toutes les cellules et tous les blocs de contrôle faibles sont libérés.

Le [test de génération](../../../purust/purust/tests/codegen/known-nullaries.mjs) conserve volontairement des appels directement dans les conditions, sans transformation préalable de PBO. Il exécute le Rust pour contrôler l’évaluation unique, les conversions, la branche `else`, les tags à données et l’absence de capture ajoutée dans les corps différés.

## Mesures du runner complet

Même configuration avant/après : Rust 1.96.0, release `opt-level=1`, mimalloc d’origine. Cinq paires alternées de processus complets ; chaque processus effectue un échauffement puis retient le meilleur de dix essais par benchmark. Les chiffres sont les médianes de cinq processus par benchmark ; le total additionne ces médianes. Aucune mesure n’a été exclue.

| Benchmark | Avant | Après | Évolution |
| --- | ---: | ---: | ---: |
| RBTree | 21,222 ms | 19,569 ms | −7,8 % |
| Suite complète | 44,380 ms | 42,923 ms | −3,3 % |

Les [résultats complets](results.json) comprennent chaque processus et chaque benchmark. Le binaire reste à **1 509 120 octets**. Les **14 résultats attendus** sont contrôlés dans le runner propre après régénération, dans la reconstruction de référence et dans les dix processus de mesure.

Le [README officiel relu](../../../altbak.pub/README.md) indique RBTree compilé **20,099 ms**, natif optimisé **36,070 ms**, total compilé **42,74 ms**. Ces références historiques sont distinctes de la série appariée ci-dessus. Le natif n’a pas été remesuré ici. La comparaison de cette étape est faite avec le partage local déjà intégré, dont la précédente mesure ne montrait pas de gain de vitesse.

## Validation et reproduction

**17 tests de génération + 8 tests TAST passent**, ainsi que **`bin/rust/run -c` et les 14 résultats attendus**. Les empreintes des sources, des noyaux, des binaires et des preuves sont consignées dans [validation.json](validation.json).

Depuis la racine d’altbak.pub-purust :

```sh
python3 scratch/rust-existing-empty-20260909/prototype-check.py
python3 scratch/rust-existing-empty-20260909/prototype-time.py
python3 scratch/rust-existing-empty-20260909/count.py
python3 scratch/rust-existing-empty-20260909/measure.py nouvelle-serie
```

Les trois premiers scripts reconstruisent leurs exécutables depuis les noyaux enregistrés, avec la bibliothèque mimalloc du runner généré. `measure.py` utilise les deux runners complets sauvegardés avant/après. Les exécutables, logs et compilations intermédiaires sont ignorés par Git. Pour refaire les mesures sans les exécutables locaux, reconstruire séparément les deux versions identifiées par `integration.patch`, puis sauvegarder leurs runners avant de lancer le script.
