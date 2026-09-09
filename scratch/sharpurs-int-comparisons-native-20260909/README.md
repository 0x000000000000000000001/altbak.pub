# Comparaisons Int : intégration dans le générateur Sharpurs

Audit du 9 septembre 2026. L'intégration des comparaisons entières natives dans la génération normale abaisse le temps médian de **RBTree de 33,35 %** et le **total de 23,74 %**, sur cinq processus par variante. Les copies mesurées proviennent directement du générateur.

## Résultats

| Mesure | Avant | Après | Écart |
| --- | ---: | ---: | ---: |
| Total médian | 284,39 ms | 216,87 ms | −23,74 % |
| RBTree médian | 204,36346 ms | 136,20571 ms | −33,35 % (×1,500) |
| Étendue des totaux | 281,59–287,05 ms | 215,67–219,88 ms | |
| Étendue RBTree | 204,08108–207,14446 ms | 134,91113–138,26508 ms | |

La paire préalable donnait **286,74 → 215,78 ms** au total et **207,19000 → 135,72538 ms** pour RBTree. Elle est conservée séparément sous `probe-*` et n'entre pas dans les médianes de confirmation.

Les **140 sorties**, noms et ordre concordent avec la référence historique, avec la profondeur finale **22** pour RBTree. Les deux configurations runtime sont identiques. Les empreintes des sources, des assemblies et des autres fichiers runtime ont été vérifiées durant les dix processus. `comparison.json` contient les résultats détaillés ; `executions.json` conserve l'ordre, les durées murales et la charge système, et `metadata.json` identifie les runtime/configurations et artefacts exécutés.

Le retrait à la médiane RBTree est de **68,16 ms**. Les médianes par test ne s'additionnent pas nécessairement à celle du total. Les petites variations des autres workloads ne sont pas attribuées individuellement à la transformation : plusieurs tests dont le module est inchangé varient aussi. L'expérience isolée antérieure avait donné 212,23 ms de total et 130,17 ms pour RBTree ; ces chiffres viennent d'une autre série, tandis que le témoin simultané utilisé ici est `before/`.

## Provenance des copies

Le témoin `before/` fige les 355 entrées de build de `run/bak/sharp/output/Main` avant la modification du générateur. `before-inputs.json` conserve les empreintes et l'heure de capture. `capture_after.py` copie ensuite la nouvelle génération normale dans `after/`, contrôle que les listes de fichiers concordent et refuse les modifications du point d'entrée, de `App.fs`, de `Bench.fs`, des FFI et des projets/configurations de build. Les différences par module sont conservées sous `diffs/`.

La capture après constate **15 modules F# modifiés** et **87 fichiers protégés strictement identiques**. `inputs.json` conserve toutes les empreintes et la liste exacte. Les autres fichiers de build restent identiques. Chaque DLL est accompagnée d'un manifeste `build-before.json` ou `build-after.json` qui relie les 355 sources, le log de compilation et les huit fichiers du répertoire runtime.

Après les mesures, le répertoire généré normal a été compilé en Release en **32,13 s**, avec **zéro erreur, zéro avertissement**. Ses 355 entrées correspondent à la copie `after/` mesurée et sa configuration runtime est identique. La DLL normale est donc à jour ; aucune exécution supplémentaire du benchmark normal n'a été ajoutée. `build-normal.log` et `build-normal.json` conservent cette vérification.

## Validation et protocole

- Chaque copie a été compilée en Release avec le même SDK .NET **8.0.423**, runtime **8.0.29**, et `-p:NuGetAudit=false` : **zéro erreur, zéro avertissement**. Les temps de build étaient de 31,52 s et 32,07 s ; les sources et assemblies sont enregistrées via `capture_build.py`.
- La vraie DLL intégrée passe les **27 087 vérifications** de `VerifyTree.fsx` : invariants rouge/noir, ordre, contenu, profondeur, cas limites Int32, doublons et persistance. `verify-tree.log` détaille les sept ordres d'insertion.
- Une paire préalable a été exécutée puis validée via `run.py --probe` / `analyze.py --probe` avant confirmation.
- Les vrais `App.main` ont été exécutés sur cinq processus frais par variante, ordre **ABBAABBAAB**, sans compilation concurrente. Le harness conserve ses échauffements et le meilleur temps sur dix pour chacun des 14 tests.
- Les 140 sorties, noms et ordre ont été contrôlés contre `expected-output.log`, ainsi que les empreintes des entrées et des assemblies pendant toute la mesure, avec des configurations runtime identiques.
- Les médianes et étendues décrivent les temps d'exécution. Aucun diagnostic mémoire n'entre dans cette expérience.

`run.py` refuse d'écraser une série existante. `analyze.py` vérifie les logs et produit `comparison.json`. Les copies volumineuses et leurs builds sont ignorés par Git ; scripts, logs, empreintes et diffs restent conservés.

## Référence officielle

`README.baseline.md` est une copie du README officiel d'altbak.pub : **294,80 ms** au total compilé, **213,760 ms** pour RBTree et **65,156 ms** dans sa dernière colonne native optimisée. Ce dernier chiffre emploie une structure mutable `SortedSet<int>` et retourne sa taille ; c'est un repère historique, avec un algorithme et un résultat distincts du benchmark compilé.

## Implémentation et tests

`Sharpurs.IntComparison.fromExpr` reconnaît dans le TAST source parsé par PBO les appels canoniques à `Data.Ord.lessThan` et `Data.Ord.greaterThan`. Il exige exactement trois arguments, le dictionnaire qualifié `Data.Ord.ordInt` annoté `Ord Int`, deux opérandes `Int`, un résultat `Boolean` et des annotations cohérentes sur les applications intermédiaires. Un `TypeApp Int` est accepté avec sa signature instanciée et le `ForAll` correspondant. Les noms et les types sont examinés dans l’AST ; aucun nom de benchmark n’intervient.

La règle intervient dans `CodeGen.translateExpr`, sur le chemin qui émet encore les expressions TAST source. Elle ne nécessite pas de convertir toute la fonction depuis l’IR optimisé PBO. Elle produit une comparaison F# directe, avec le `box`/`unbox<int>` nécessaire pour accepter aussi bien les arguments `obj` que les champs ADT déjà natifs, puis conserve le résultat public `obj`. Chaque opérande est évalué une fois, de gauche à droite. Les **38 sites** acceptés se répartissent dans **15 modules** ; trois sont dans RBTree (les deux comparaisons de `ins` et celle de `max`, qui n’est plus appelée par `depth` natif).

`tests/fixtures/IntCompare.purs` est compilé avec le fork TAST puis lu par le parseur PBO. Le JavaScript émis sert d’oracle pour les comparaisons, notamment aux bornes Int32 et pour l’ordre personnalisé inversé. Le harnais F# utilise des dépendances `Ord` instrumentées pour vérifier que les comparaisons natives évitent le dictionnaire, tandis que les formes génériques continuent à l’utiliser. Une annotation explicite de fonction crée ici un alias local dans le TAST ; cet alias reste sur le chemin générique.

- `test-int-comparison.log` : **32 contrôles de reconnaissance/routage**, **396 assertions F#** ; signatures absentes/contradictoires, dictionnaire local/personnalisé, propriétaire incorrect, TypeApp incohérent, applications partielles/surappliquées, ordre des opérandes, Int32, Number/String génériques et ordre inversé.
- `test-adt-unary.log` : **11 contrôles et 63 assertions F#** sur le mélange ADT natif/générique.
- `test-local-kernel.log` : **68 contrôles et 43 assertions F#** ; `test-runtime.log` : **21 assertions** d’appels/exceptions.
- `test-spago.log` : **43 assertions IntKernel et 6 assertions de routage**. La recompilation complète signale 26 avertissements sur des lignes existantes hors ajout, sans erreur ; aucun ne concerne `IntComparison` ou le nouveau bloc d’émission.
- `verify-tree.log` : **27 087 vérifications** sur la vraie DLL intégrée, couvrant invariants, résultats, doublons, ordres d’insertion, persistance et partage.

Les empreintes du backend, du bundle exécuté et des fixtures sont dans `backend-inputs.json`. Le gain expérimental précédent (deux expressions modifiées à la main, RBTree 130,17 ms) appartient à une série distincte ; le résultat retenu pour cette intégration est la comparaison simultanée documentée ici, **204,36 → 136,21 ms**.

## Reproduction

Le témoin initial est déjà figé. `python3 capture_after.py` capture une nouvelle génération une seule fois et refuse d'écraser `after/`. Compiler les deux projets avec `dotnet build before/Program.fsproj -c Release -p:NuGetAudit=false` et la commande équivalente pour `after/`, en redirigeant les sorties vers `build-before.log` et `build-after.log`, puis lancer `python3 capture_build.py before` et `python3 capture_build.py after`.

Exécuter ensuite `dotnet fsi --nologo --optimize+ --exec VerifyTree.fsx`, puis `python3 run.py` et `python3 analyze.py` sans build concurrent. Le runner refuse d'écraser les mesures existantes ; conserver les séries antérieures avant toute répétition.
