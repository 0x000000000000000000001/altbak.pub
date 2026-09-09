# RBTree : partager les constructeurs vides d’un même nœud

Le partage local supprime **100 000 allocations sur 300 001 (−33,3 %)** pour 100 000 insertions, parcours et destruction. Les octets demandés cumulés passent de **14 400 048 à 9 600 048** ; ce comptage ne mesure pas le pic de mémoire résidente. Toutes les allocations sont libérées, et `Tree` reste à **32 octets**.

**Aucun gain de vitesse n’est constaté.** Deux séries de cinq paires alternées montrent un léger recul de RBTree : **+2,6 %**, puis **+1,5 %**. La transformation est conservée pour la réduction des allocations, avec ce coût explicite.

## Preuve et changement intégré

Le [prototype isolé](probe.py) compte séparément les variantes du Rust initial. Une insertion dans `E` allouait deux `E` et un `T`. Le partage des deux enfants vides ramène ce cas à un `E` et un `T`, toutes les cellules libérées. Le [comptage du Rust réellement généré](count.py) confirme ensuite :

| Pour 100 000 insertions | Avant | Après |
| --- | ---: | ---: |
| Constructeurs vides `E` | 200 001 | 100 001 |
| Nœuds à données `T` | 100 000 | 100 000 |
| Allocations / libérations | 300 001 / 300 001 | 200 001 / 200 001 |
| Octets demandés cumulés | 14 400 048 | 9 600 048 |

[ShareNullaries.purs](../../../purust/purust/src/Purust/ShareNullaries.purs) introduit un `let` local lorsque plusieurs champs directs d’une construction contiennent le même constructeur natif sans champ. L’identité comprend sa représentation Rust et son constructeur ; les types TAST sont préservés sur les nouvelles références locales. La génération existante clone la première référence et déplace la dernière. Les noms locaux évitent les variables déjà présentes.

Les représentations converties et les enums déjà en valeur sont exclus. Les appels de fonctions restent des appels distincts. Chaque construction possède son propre partage, y compris lorsqu’elle est imbriquée : aucune racine globale n’est conservée. La signature publique et le layout restent identiques. Le [diff RBTree](RBTree.diff) ne change que la construction de la feuille dans `ins`.

## Mesures de vitesse

Même configuration : Rust 1.96.0, release `opt-level=1`, mimalloc d’origine. Chaque processus effectue un échauffement, puis retient le meilleur de dix essais par benchmark. Chaque série contient cinq paires de processus complets, dans un ordre alterné. Les chiffres ci-dessous sont les médianes par benchmark, et le total additionne ces médianes.

| Série / benchmark | Avant | Après | Évolution |
| --- | ---: | ---: | ---: |
| Première série — RBTree | 19,594 ms | 20,094 ms | +2,6 % |
| Première série — total | 41,326 ms | 41,760 ms | +1,1 % |
| Confirmation — RBTree | 19,251 ms | 19,531 ms | +1,5 % |
| Confirmation — total | 40,538 ms | 40,930 ms | +1,0 % |

La seconde série a été lancée pour vérifier le recul constaté dans la première. Les deux sont conservées intégralement : [première série](results.json), [confirmation](confirmation/results.json). L’agrégation des dix processus de chaque version est également disponible dans [combined-results.json](combined-results.json) ; les médianes agrégées ne sont pas la moyenne des deux médianes ci-dessus. Aucun processus n’a été exclu.

Le binaire reste à **1 509 120 octets**. Le noyau initial correspond exactement au Rust final de l’étape précédente sur les rotations ; les deux binaires complets ont été sauvegardés avant/après la régénération de cette étape.

La [baseline officielle relue](README-rust-baseline.md) indique RBTree compilé **20,099 ms**, natif optimisé **36,070 ms**, total compilé **42,74 ms**. Ce sont les chiffres historiques du README normal d’altbak.pub, distincts des mesures appariées. Le natif n’a pas été remesuré ici.

## Validation

- **16 tests de génération + 7 tests TAST réussis**.
- **`bin/rust/run -c` réussi dans altbak.pub-purust**, avec ses 14 résultats attendus. Ces mêmes résultats sont vérifiés dans les 20 processus de mesure.
- La [fixture TAST](../../../purust/purust/tests/tast/shared-nullaries.mjs) exécute le vrai fork, PBO, Purust et les crates Rust fraîches. Elle contrôle le budget de deux allocations d’une feuille, les partages locaux et imbriqués, deux types de modules différents portant un constructeur du même nom, la persistance, les références faibles, l’ordre des callbacks et toutes les libérations.
- Sur les deux noyaux RBTree : quatre rotations, invariants rouge/noir, ordre des clés, doublons, et clés exactes de **200 versions simultanément conservées**, avant/après insertion et destruction.

Les sources, preuves, résultats et exécutables mesurés sont identifiés par leurs empreintes dans [validation.json](validation.json). Les exécutables, logs et fichiers de compilation sont ignorés par Git.

## Reproduction

Depuis la racine du worktree `altbak.pub-purust`, avec les sorties Rust régénérées et les binaires sauvegardés dans ce dossier :

```sh
python3 scratch/rust-empty-sharing-20260909/probe.py
python3 scratch/rust-empty-sharing-20260909/count.py
python3 scratch/rust-empty-sharing-20260909/measure.py nouvelle-serie
```

Le prototype et le comptage reconstruisent leurs exécutables depuis les deux noyaux enregistrés. Pour refaire la comparaison complète sans les binaires locaux ignorés, reconstruire séparément les deux versions du générateur et sauvegarder leurs runners respectifs avant de lancer `measure.py`.

La prochaine expérience peut porter sur le réemploi de la valeur vide déjà reçue par `ins`, afin d’éviter d’en recréer une à chaque insertion. Son effet sur les allocations et la vitesse reste à mesurer ; cette généralisation n’est pas intégrée ici.
