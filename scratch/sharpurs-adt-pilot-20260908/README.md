# Pilote ADT fermé — 8 septembre 2026

`AdtPilot.purs` isole le layout récursif du benchmark RBTree :

```purescript
data Color = R | B
data Tree = E | T Color Tree Int Tree
depth :: Tree -> Int
```

Il utilise le vrai fork PureScript et le vrai Builder PBO de Sharpurs, avec la même sélection de sémantiques étrangères. Il contient aussi des valeurs témoins, une fonction de construction curried et des projections pour contrôler l’interopérabilité. Aucune insertion ni rotation d’arbre n’est optimisée à ce jalon.

## Constats dans le TAST et l’IR

- Le TAST brut est `output/AdtPilot/corefn.json`. Ses `dataDecls` nomment `Color` et `Tree`, sans variables de type ; `T.fieldTypes` est `[0, 1, 3, 1]`, résolu par `typeTable` en `Color`, `Tree`, `Int`, `Tree`.
- `input.parsed.json` conserve les déclarations et leurs annotations, résolues par le lecteur PBO. `module.metadata.json` prouve que le Builder transmet les mêmes champs structurés dans ses `dataDecls` : ADT `AdtPilot.Color`, ADT `AdtPilot.Tree`, `Int`, ADT `AdtPilot.Tree`.
- `depth.optimized.json` contient réellement `Typed (Func [Tree] Int) (Abs ...)`. Le pattern matching devient un `Branch` testé par `OpIsTag E` et `OpIsTag T`, avec `Fail` en branche de défaut.
- Dans la branche `T`, `GetCtorField` indexe les sous-arbres à **1 et 3**. Les deux appels récursifs à `depth` sont évalués et liés par des `Let` successifs, puis le `max` inliné devient `OpIntOrd OpGt` ; le résultat est incrémenté par `OpIntNum OpAdd`.
- `rootValue`, `rootColor`, `leftChild` et `singletonWith` rendent observables les positions de champs et les conversions aux frontières publiques.

## Vérifications exécutées

Compilation de la seule fixture et des sources Prelude locales : succès, sans avertissement après renommage de la projection `leftChild`. Dump de **51 modules optimisés**, dont les **14 bindings** du pilote. Les empreintes des **355 fichiers F#/C#** du benchmark sont inchangées ; aucun benchmark n’a été lancé par la préparation de cette fixture.

`check-fixture.mjs` exécute **18 assertions** sur la sortie JavaScript du même compilateur : arbre vide, singleton, arbre asymétrique de profondeur 3, champs `Int` aux deux bornes, couleurs distinctes, réutilisation d’une application partielle, répétition de valeurs, partage physique d’un sous-arbre et profondeur 1 000. Le partage est contrôlé par identité de référence, en plus des valeurs.

Le nouveau générateur est désormais validé sur la même fixture. `AdtPilot.generated.fs` contient des unions F# ordinaires avec les champs `AdtPilot_Color * AdtPilot_Tree * int * AdtPilot_Tree` et une fonction `AdtPilot_depth_adt_native : AdtPilot_Tree -> int`. Les 14 définitions natives sont exemptes de `obj`, `unbox` et `sharpurs_apply` ; seuls les wrappers publics effectuent les conversions. Les patterns et leurs projections sont construits à partir des mêmes métadonnées que les déclarations.

Le test durable `sharpurs/sharpurs/tests/adt-kernel.mjs` recompile la fixture dans un dossier temporaire, utilise le vrai callback du Builder PBO et exécute le F# effectivement émis : **33 contrôles du générateur** (dont 27 refus ciblés) et **32 assertions F#** passent. La réflexion F# confirme les quatre types de champs. Les contrôles couvrent aussi les résultats natifs et boxed, le constructeur public curried réutilisable, les couleurs, les bornes Int, le partage physique et la profondeur 1 000.

Les formes non prises en charge produisent `Nothing` pour le module entier : données polymorphes, références externes ou absentes, champs hors du sous-ensemble fermé, métadonnées contradictoires, collisions de noms, niveaux locaux incorrects, arité incorrecte et `TypeApp` non résolu. Les fonctions du pilote n’acceptent que des paramètres et résultats de premier ordre (`Int`, `Boolean`, ADT locaux fermés), avec applications saturées. Les variables polymorphes ne sont pas effacées pour obtenir un chemin natif.

Les suites précédentes passent également : 43 assertions IntKernel et 6 de routage en PureScript, 68 contrôles du noyau local et 43 à l’exécution F#, 695 contrôles Int et 21 du runtime d’application. La compilation finale des deux nouveaux modules ne produit aucun avertissement. Les journaux de cette validation sont conservés dans `validation/` et leurs empreintes dans `validation.json`.

## Portée du jalon

`Sharpurs.AdtKernel.fromModule` est un pilote du générateur appelé par le test ; il n’est pas encore activé dans `Main` ou `CodeGen`. Les modules du benchmark restent sur leur layout existant. La prochaine étape doit établir la compatibilité d’un producteur typé et d’un consommateur utilisant encore l’ABI `obj`, en particulier les constructions et patterns entre modules, avant l’intégration dans RBTree.

Aucun nouveau benchmark ni gain chiffré n’est revendiqué. Le README officiel d’altbak.pub indique actuellement **294,80 ms** au total, dont **213,760 ms** pour RBTree ; la dernière série à cinq processus donnait **296,39 ms** de médiane. Ce pilote valide la représentation et la sémantique, sans modifier ces scores.

## Reproduction

Depuis ce dossier :

```sh
../../run/bak/js/node_modules/.bin/purs compile AdtPilot.purs '../../../sharpurs/sharpurs-prelude/src/**/*.purs' --output output --codegen corefn,js
node --expose-gc --stack-size=65536 --max-old-space-size=16384 dump-fixture.mjs
node check-fixture.mjs
node emit-pilot.mjs
```

Le test durable se lance depuis `sharpurs/sharpurs`, après `spago build` ou `npm run build` :

```sh
PURS=/chemin/du/fork/purs DOTNET=/chemin/de/dotnet npm run test:adt-kernel
```

`PRELUDE_SRC` peut sélectionner un dossier de sources Prelude ; par défaut le test utilise celui installé dans `.spago/p`. Le test supprime son dossier temporaire après exécution. Le snapshot `AdtPilot.generated.fs` peut aussi être compilé directement avec `dotnet fsi --exec AdtPilot.generated.fs`.

`compile.log`, `dump.log`, `reference.log` conservent les exécutions. `metadata.json` contient les empreintes et les révisions observées lors de la capture ; `output/` et `.purmeta/` sont isolés et ignorés par Git.
