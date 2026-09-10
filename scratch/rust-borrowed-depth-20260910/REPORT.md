# Prototype isolé : emprunter l'arbre pendant `depth`

Préparé puis mesuré le 10 septembre 2026. Cinq paires de runners complets donnent **RBTree 11,197 → 11,091 ms**, **total 12,296 → 12,193 ms**. Le gain observé est petit (**0,103 ms sur la suite**) et la cinquième paire est défavorable. Les plages se recouvrent ; cette expérience ne démontre pas un gain suffisamment régulier pour en faire la prochaine priorité.

## Périmètre

La référence est le Rust généré après l'intégration des gardes après appel, copié depuis `run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs`. Son SHA-256 est `d3330b6fd36a36c9dd7f321383ab64c510328add641b903afd29bda52a493936`.

Seule la fonction `Test_RBTree_depth` change dans la copie expérimentale : son interface publique possède toujours `Rc<Tree>`, appelle un worker privé qui reçoit `&Tree`, puis détruit explicitement le propriétaire avant de retourner la profondeur. Le worker parcourt les références des deux enfants dans le même ordre. Les fonctions d'insertion et de construction restent identiques.

Le prototype est local à ce dossier. Aucun fichier de compilateur ou de sortie courante n'est modifié. Les sources copiées et les binaires sont dans `build/`, ignoré par Git. Les empreintes sont conservées dans [metadata.json](metadata.json) et [runner-build.json](runner-build.json).

## Preuves sans chronométrage

Le comptage utilise le même wrapper d'un mot autour de `std::rc::Rc` que les expériences précédentes. Il mesure des opérations logiques, pas des instructions machine. Le noyau est extrait du code généré en conservant toutes ses fonctions utiles et ses helpers. [Comptage et bilans](counts.json).

| Parcours puis destruction de l'arbre de 100 000 nœuds | Référence | Emprunt |
| --- | ---: | ---: |
| Clones | 200 000 | 0 |
| Relâchements sans destruction | 300 000 | 100 000 |
| Destructions de cellules | 100 001 | 100 001 |

La construction conserve exactement les mêmes opérations. Les bilans globaux de propriétaires et de destructions passent, y compris sur 200 versions conservées.

La [fixture supplémentaire](validation-depth.rs) vérifie la profondeur 22 pour 100 000 nœuds, la destruction du propriétaire avant le retour de l'appel, vingt lectures successives d'un arbre conservé, une insertion ultérieure, des enfants partagés par les deux branches et des références fortes/faibles indépendantes. Les anciens contrôles des quatre rotations, des invariants rouge/noir, de 200 versions persistantes, des bornes `i64`, de 512 insertions avec partage mixte et de l'expiration finale des références faibles passent aussi pour les deux variantes. [Résultats natifs](validation.json).

## Runners prêts

Les deux runners complets utilisent le profil existant **O1, debug=true, mimalloc**, le même `Cargo.lock` et les autres crates inchangées. `Purs_App`, `Purs_Test_RBTree` et la crate principale sont isolées. Leur cache de compilation est explicitement nettoyé pour chaque variante afin d'éviter une fausse fraîcheur Cargo. Les deux binaires ont des empreintes différentes. Aucun des deux n'a été exécuté pendant la préparation.

La mesure doit porter sur le runner complet : **construction, parcours et destruction restent tous dans la durée RBTree**. Une durée du parcours qui exclut la destruction ne constitue pas une comparaison équivalente. Les deux variantes doivent être mesurées sans compilation, instrumentation ni autre test concurrent.

Depuis la racine d'`altbak.pub-purust`, après coordination :

```sh
python3 scratch/rust-borrowed-depth-20260910/runner.py time
```

Cette commande a effectué cinq paires alternées, conservé la chauffe et le meilleur de dix mesures propres au runner et vérifié les quatorze résultats. [Mesures](runner-results.json) : totaux par processus **12,262–12,415 ms** avant, **12,129–12,335 ms** après. Construction, parcours et destruction sont inclus, sans compilation ou autre test concurrent. Aucun code de production n'est modifié.

Le README officiel relu pour cette expérience donne **11,754 ms RBTree et 12,98 ms total**, contre **36,070 ms et 36,13 ms** dans sa dernière colonne native. Ces références historiques restent distinctes des mesures appariées. La piste demeure ouverte, mais les [lectures empruntées de Records](../rust-record-borrows-20260910/REPORT.md) montrent un gain local plus net dans l'essai parallèle, chronométré séparément.

Pour reproduire la préparation, séquentiellement :

```sh
python3 scratch/rust-borrowed-depth-20260910/probe.py prepare
python3 scratch/rust-borrowed-depth-20260910/probe.py count
python3 scratch/rust-borrowed-depth-20260910/probe.py validate
python3 scratch/rust-borrowed-depth-20260910/runner.py build
```
