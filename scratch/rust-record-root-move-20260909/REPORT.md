# Records : déplacement de la racine intégré

9 septembre 2026. Premier bébé de l'[audit Records](../rust-records-audit-20260909/REPORT.md) : calculer les valeurs de remplacement avant de déplacer la racine au dernier usage. Le compilateur conserve la représentation actuelle et les copies des deux niveaux intérieurs.

## Résultats

| Série de runners complets | Records avant | Records après | Écart | Total avant → après |
| --- | ---: | ---: | ---: | ---: |
| Cinq paires initiales | 0,947 ms | 0,848 ms | −10,5 % | 19,108 → 19,148 ms |
| Cinq paires de confirmation, mêmes binaires | 0,921 ms | 0,828 ms | −10,1 % | 19,219 → 18,951 ms |
| Dix paires réunies | 0,938 ms | 0,837 ms | −10,8 % | 19,112 → 19,077 ms |

Le gain sur Records est confirmé dans les deux séries. **Aucun gain global n'est établi au regard de la dispersion**, dominée par RBTree. La première série ne permettait pas de conclure sur le total ; une seconde série avec les mêmes binaires a vérifié ce point. Les deux séries sont conservées, sans sélectionner la plus favorable. Le total additionne les médianes par benchmark, pas les médianes des totaux de processus.

Le comptage séparé sur le code réellement généré donne **30 003 → 20 003 allocations et libérations** pour 10 000 itérations : **10 000 copies de racine supprimées**. Les trois cellules initiales et les deux copies de sous-records par itération subsistent. Les 302 empreintes de sources sont comparées : **seul `Purs_Test_Records/src/lib.rs` change**.

Référence historique relue dans le [README normal](../../../altbak.pub/README.md#rust) : Records **0,969 ms** compilé / **0,004 ms** natif optimisé ; RBTree **18,449 / 36,070 ms** ; totaux **19,95 / 36,13 ms**. Les comparaisons appariées ci-dessus ont leur propre baseline. Le natif de Records n'expose que l'accumulateur final `f`, tandis que le code généré conserve le record complet.

Le prototype d'audit donnait **−28,6 %** localement pour le déplacement de racine. L'intégration démontre environ **−10 %** dans le runner complet ; le chiffre du prototype n'est pas repris comme gain livré. Le [diff généré](Records.diff) montre notamment que la base reste vivante pendant toutes les valeurs de remplacement, ce qui conserve un clone temporaire supplémentaire sur la dernière lecture. Les temps du prototype et du runner ne suffisent pas à attribuer à eux seuls toute la différence à ce clone.

## Règle

Dans [CodeGen, branche Update](../../../purust/purust/src/Purust/CodeGen.purs), la nouvelle émission s'applique uniquement si :

- Le TAST indique un record fermé pour la base.
- La base est un local, éventuellement sous des annotations sans conversion de représentation.
- Ce local n'est plus vivant après la mise à jour, y compris parmi les captures d'une fonction réutilisable.
- Au moins une valeur de remplacement lit ce local ; les bases déjà déplaçables gardent leur émission habituelle.

Les valeurs sont générées dans leur ordre initial et stockées dans des temporaires. La base reste vivante pendant leur évaluation, puis est déplacée avant les setters. Un alias conservé par un résultat ou par une closure maintient le partage : le runtime copie alors la racine. Aucun accès mutable à un enfant ni `unsafe` supplémentaire n'est introduit.

```rust
let new_a = /* calcul depuis l'ancienne racine */;
let new_b = /* calcul depuis l'ancienne racine */;
let mut base = old_root;
base.set_a(new_a);
base.set_b(new_b);
base
```

Les bases produites par un appel, les rows ouvertes, les conversions et les bases utilisées plus tard gardent le chemin précédent. La règle ne contient aucun nom de benchmark. B14 reste jaune : sa couverture s'étend aux racines de records ; B22 garde les setters existants. B21/B23 et les sous-records restent des chantiers distincts.

## Validation

- `npm run build` réussi. La première recompilation signale 64 avertissements à des lignes inchangées de `Main.purs` et `CodeGen.purs` ; leurs emplacements ont été comparés à la révision de départ. Aucun avertissement ne vise les lignes ajoutées. Le build incrémental du runner passe également.
- **21 tests de génération, 13 tests TAST**, tous réussis. Les nouveaux tests ciblés passent aussi séparément.
- Fixture fraîche compilée avec le fork TAST, puis Rust O1 avec les contrôles de débordement activés : adresses uniques réutilisées, anciennes racines partagées conservées, valeurs permutées lues avant les écritures, usages ultérieurs, appels de base, callbacks dans l'ordre, closure retenant l'ancienne racine, fonction sauvegardée appelée plusieurs fois, panic/overflow et libération d'un résultat temporaire si le RHS suivant échoue.
- `bin/rust/run -c` réussi dans le worktree. Les **14 sorties** sont vérifiées dans ce run et chacun des **20 processus appariés**.
- Comptage avant/après pour 0, 1, 2, 10 et 10 000 itérations : toutes les allocations sont libérées.
- Le noyau généré avant/après passe 24 combinaisons de graines/comptages, avec les quatre champs observés et conservation de racines, de sous-records et de feuilles partagés.

Un premier intitulé de champ dans la fixture (`{ a :: Int | r }`) exposait une collision avec le nom runtime `Record_a`, [également reproduite avec le bundle avant modification](existing-record-name-collision.json). La fixture de row ouverte utilise `score` ; la règle de déplacement refuse toujours les rows ouvertes. La collision de nom n'a pas été corrigée dans cette étape.

## Mesure et reproduction

Les deux runners complets sont figés avant/après reconstruction. O1/mimalloc, protocole du runner conservé : échauffement puis meilleur temps de dix par benchmark, ordre avant/après alterné entre processus. Toutes les compilations et instrumentations de cette tâche sont terminées avant le chronométrage.

- [Métadonnées et empreintes initiales](metadata.json), [comparaison des sources](source-comparison.json), [diff Rust](Records.diff).
- [Première série](runner-results.json), [confirmation](runner-confirmation.json), [médianes des dix paires](runner-combined.json).
- [Comptage et persistance du Rust généré](generated-checks.json), [validation](validation.json).
- `python3 check-generated.py` compile les noyaux avant/après puis exécute les contrôles séparés, avec les harnesses de l'audit.
- `python3 measure-runner.py` mesure cinq paires ; `python3 measure-runner.py --confirmation` écrit la confirmation séparément.

Les sources et binaires avant/après sont conservés dans `build/`, ignoré par Git. Baseline Purust : `cce9ca16eb04fc9a786d2a667b046000f4f9d07d`. Les checkouts normaux d'altbak.pub et de PBO, les FFI et les sources PureScript du benchmark n'ont pas été modifiés.

## Suite

Les **20 000 copies restantes** viennent des deux niveaux intérieurs. Le prochain bébé consiste à réutiliser **un enfant** avec une preuve d'accès exclusif au chemin, puis à couvrir les alias conservés et les exceptions. Le prototype de réutilisation complète à 3 allocations reste une cible expérimentale ; il n'est pas intégré ici.
