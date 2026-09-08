# Correction des annotations après TypeApp — 8 septembre 2026

Le même fixture de 13 lignes passe désormais son contrat : `identity` reste en `a`, `useInt` ne contient que des annotations `Int` dans son corps, `useString` des `String`, et `generic` des `b`. La preuve initiale du défaut reste dans [le dossier précédent](../sharpurs-typeapp-fixture-20260908/README.md).

PBO conserve le `ForAll` de la déclaration et l’annotation de la référence sous `ExprTypeApp`. Il transporte les arguments de type jusqu’à l’implémentation, puis substitue ses annotations **avant** l’évaluation et l’injection des arguments de valeur. La substitution est simultanée, respecte les quantificateurs internes et renomme les variables liées pour éviter la capture. Le corps générique partagé reste intact. Les types non résolus et les références `InlineNever` gardent leur `TypeApp`.

Les évaluateurs de primitives FFI reçoivent leur épine d’arguments runtime habituelle. Les arguments de type restent disponibles pour l’évaluation de l’implémentation syntaxique.

Validation :

- [contract.log](contract.log) : les 4 bindings passent, sans le mode de caractérisation `--expect-current`.
- [scope-tests.log](scope-tests.log) : 18 groupes de régression PBO, dont `ForAll` interne homonyme, capture de `b` lors de `@b @String`, rangées ouvertes, contraintes, instanciations indépendantes et replis curried/uncurried/effets.
- [sharpurs-tests.log](sharpurs-tests.log) : 43 assertions IntKernel et 6 de routage passent.
- [build.log](build.log) : compilation et bundle réussis. Les avertissements existants du projet ne sont pas traités dans cette étape.
- [metadata.json](metadata.json) : empreintes des entrées, du PBO source/compilé et du test de portée. Les 355 fichiers F#/C#/projets contrôlés sont inchangés.
- [Vrai dump Polymorphism corrigé](../sharpurs-poly-typeapp-fix-20260908/README.md).

Reproduction après `npm run build` dans `sharpurs/sharpurs` :

```sh
# Depuis ce dossier ; le purs utilisé est le fork TAST de l’utilisateur.
../../run/bak/js/node_modules/.bin/purs compile TypeAppFixture.purs --output output --codegen corefn,js
node --expose-gc --stack-size=65536 --max-old-space-size=16384 dump-fixture.mjs
python3 check-fixture.py

# Depuis purescript-backend-optimizer-sharpurs :
npm run test:typeapp -- ../sharpurs/sharpurs/output
```

Ce jalon corrige les instanciations explicites des implémentations globales disponibles. Il ne généralise pas l’instanciation des closures locales ni la spécialisation des champs de dictionnaires sans `TypeApp`. Aucune émission F# ni mesure de performance dans cette étape.
