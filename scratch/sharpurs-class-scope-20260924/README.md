# Boucle Polymorphism native rétablie — 24 septembre 2026

## Résultat

Sur **trois processus officiels validés** du vrai `App.main` (`altbak.pub`, `bin/sharp/run`), la
médiane du total passe de **339,60 à 58,63 ms** (**−82,7 %, ×5,79**) et Polymorphism de
**287,90925 à 2,52327 ms** (**×114,1**). Chaque passage est validé par le validateur officiel :
les 14 sorties affichées sont identiques et les 14 valeurs concordent avec l'oracle.

| Test | Avant (README figé) | Médiane des 3 passages | Ratio |
| --- | ---: | ---: | ---: |
| Polymorphism | 287 909,25 µs | 2 523,27 µs | ×114,1 |
| Total | 339 604,00 µs | 58 630,03 µs | ×5,79 |

Dans le workspace de travail (`altbak.pub-sharpurs`, mêmes sources), la comparaison A/B à
binaires identiques donne **344–358 → 59–62 ms** et Polymorphism **287 874–293 481 → 2 223–2 253 µs**.

## Cause : le scoping des variables de type du 13 septembre

Le fork `purescript` a ajouté le scoping des variables quantifiées (`a$scopeN`, commit
`40840b3`, `scopeTypeVariables` dans `CoreFn/Desugar.hs`). Les accesseurs de classe générés par
`Sugar.TypeClasses.typeClassMemberToDictionaryAccessor` annotent leur corps avec la **variable
de classe** (`a`), alors que le binding quantifie sa **propre variable** (`a$scope0`).

Le correctif PBO du 8 septembre (`instantiateNeutralType`/`evalTypeApp`, branche `edge-sharpurs`)
substitue la variable quantifiée du binding dans tout le corps inliné ; il ne pouvait atteindre
le `a` orphelin que tant que les noms coïncidaient. Depuis le scoping, le `a` survit :

```text
Typed Int / Typed TypeVar(a) / Typed Int / PrimOp OpIntNum(OpAdd)   <- IR de act
```

`Sharpurs.IntKernel.fromInt` exige `Int` pour chaque annotation de la boucle et rejette donc
**tout le binding** : `act` retombait sur `polyLoop` générique (accès `Map.find` au dictionnaire
et `sharpurs_apply` pour les 10 millions d'itérations). Les autres noyaux natifs (TCO, RBTree,
Lazy, thunks) ne sont pas concernés.

## Correctif

`purescript-backend-optimizer-sharpurs/src/PureScript/Backend/Optimizer/Convert.purs` :

- `alignClassMemberAnnotations` s'applique avant conversion aux bindings dont l'identifiant est
  une méthode d'une classe **déclarée dans le module** et dont la signature contient la
  contrainte correspondante ;
- les variables déclarées de la classe (`classDecls[].vars`) sont mises en correspondance avec
  les arguments de la contrainte qui sont des variables (`a` → `a$scope0`) ;
- seules les annotations de type sont réécrites ; les signatures, annotations absentes, classes
  inconnues, bindings non-membres et contraintes concrètes restent inchangés.

La règle ne mentionne aucun benchmark : elle généralise la correspondance déclarée entre un
membre de classe et sa propre quantification.

## Vérifications

- `test/typeapp.mjs` : 18 assertions existantes passent (contre `sharpurs/sharpurs/output`).
- `test/class-member-scope.mjs` (nouveau) : 6 assertions — membre déclaré, annotations imbriquées,
  groupes récursifs, binding non-membre, argument concret, classe inconnue.
- `spago test` sharpurs : 43 + 6 assertions. `test:runtime` 21, `test:kernel` 695,
  `test:local-kernel` 68 + 43.
- Génération normale : **1 fichier sur 359** change (`Test.Polymorphism.fs`), preuve
  [shared-code-proof.json](shared-code-proof.json) ; le diff est
  [Test.Polymorphism.diff](Test.Polymorphism.diff).
- 3 passages officiels complets validés : [official-medians.json](official-medians.json).
- Programme compilé Release sans erreur ni avertissement ; la DLL normale correspond aux sources
  mesurées.

## Artefact de tiering JIT sur les autres tests

Sous les réglages runtime par défaut, les tests **dont le code est identique** mesurent parfois
2 à 3 fois plus lentement après le correctif (List Processing 198 → 403 µs, Records 2158 → 3258 µs,
Ackermann 110 → 180 µs, etc.). Ce n'est pas une régression de code :

- les fichiers F# sont identiques à l'octet près (un seul fichier change) ;
- en A/B à binaires construits de façon identique, seuls ces tests varient, pas Polymorphism ;
- avec `DOTNET_TieredCompilation=0`, les deux variantes donnent des temps **égaux** pour ces tests
  (List 275,6 / 277,9 µs ; Records 2 313,9 / 2 389,2 µs ; Primes 513,1 / 515,9 µs) alors que
  Polymorphism reste massivement gagnant ;
- l'ancien total bénéficiait d'un long échauffement accidentel : les 10 millions d'itérations
  pathologiques réchauffaient le JIT avant les mesures.

Autrement dit, ces cellules varient selon l'état du JIT partagé au sein du processus ; le code
généré est identique. Le protocole officiel (médiane de trois processus) publie ces valeurs
telles quelles.

## Limites connues

- Neuf suites de fixtures sharpurs (`test:thunk-kernel`, `test:int-comparison`, `test:int-arithmetic`,
  `test:constructor-typeapp`, `test:direct-call`, `test:adt-*`) échouent à la compilation des
  fixtures : le `purs` de `/Users/0x1/.local/bin` (16 septembre) ne parse pas le FFI ESM de
  `prelude-6.0.2/src/Data/Symbol.js`. L'échec précède tout code backend et existe indépendamment
  de ce correctif.
- Les valeurs du tableau officiel `altbak.pub/README.md` n'ont pas été réécrites : la campagne de
  publication passe par `bin/benchmark/measure.py` et `publish.py` avec un artefact figé.

## Reproduction

```sh
# PBO et backend
cd /Users/0x1/Documents/htdocs/purescript-backend-optimizer-sharpurs
node test/typeapp.mjs ../sharpurs/sharpurs/output
node test/class-member-scope.mjs ../sharpurs/sharpurs/output

cd ../sharpurs/sharpurs
npm run build

# Benchmark officiel (3 passages)
cd /Users/0x1/Documents/htdocs/altbak.pub
./bin/sharp/run
```

[PBO patch](pbo-fix.patch) · [nouveau test](class-member-scope.mjs) · [logs TypeApp](typeapp-tests.log)
· [logs class-scope](class-scope-tests.log) · [artefact avant](Test.Polymorphism.before.fs)
· [artefact après](Test.Polymorphism.after.fs) · [mesures officielles](official-medians.json)
· [A/B workspace](workspace-ab.json)
