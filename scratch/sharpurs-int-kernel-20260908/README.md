# Baby step 2.2 — représentation typée des noyaux Int

Réalisé le 8 septembre 2026, après l’observation du binding optimisé TCO. Cette étape ajoute la représentation et sa conversion conservatrice. Elle ne branche pas encore l’émission F# sur `backendMod` et ne change pas le runtime du benchmark.

## Changement

Nouveau module : `/Users/0x1/Documents/htdocs/sharpurs/sharpurs/src/Sharpurs/IntKernel.purs`.

`IntKernel` porte l’identité qualifiée, les paramètres identifiés par leurs niveaux lexicaux et un corps entier. Paramètres et résultat sont implicitement tous `Int`. `IntCondition` distingue les comparaisons des expressions entières ; `IntModulo` conserve l’opération sémantique sans la convertir prématurément en `%` F#.

`fromBinding` accepte uniquement un binding entièrement pris en charge et retourne `Nothing` sinon. Il exige une signature globale `Func [Int, …] Int` avec au moins un paramètre, parcourt les `Abs` réelles à travers les `Typed`, valide chaque annotation et vérifie les niveaux des paramètres. Les noms de variables et de benchmarks ne décident pas de l’acceptation.

Le corps peut contenir des littéraux, références locales, additions/soustractions/modulos, branches à égalité entière et auto-appels saturés en position terminale. Les arguments d’appel, opérandes et conditions ne sont pas des positions terminales. Les annotations contradictoires, variables libres, niveaux dupliqués ou négatifs, arités incohérentes, appels externes, `TypeApp` résiduels, effets et autres nœuds sont rejetés.

Le résultat `Nothing` prépare le repli vers le générateur existant lors du branchement suivant ; aucun appel à ce convertisseur n’est encore ajouté à `Main.purs`.

## Vérification empirique

- Le vrai binding `Test.TCO.deepTailRec` enregistré au baby step 2.1 a été restauré avec ses constructeurs PBO puis donné au convertisseur compilé : accepté.
- Le résultat `deepTailRec.int-kernel.json` contient deux paramètres `Int`, `IntIf`/`IntEqual`, puis un `IntTailCall` avec `IntSubtract`, `IntAdd`, `IntModulo`.
- Le test PureScript `/Users/0x1/Documents/htdocs/sharpurs/sharpurs/test/Main.purs` contient **43 assertions**. Elles passent avec `spago test`, sans dépendance ni framework ajouté.
- La fixture TCO vérifie le résultat structurel complet. Les autres cas couvrent renommage, identité lexicale, annotations, arités, rejet intégral des branches non prises en charge, limites des littéraux entiers et positions de récursion.
- Construction du module et bundle réussie. Validation finale du fichier de test : zéro erreur et zéro avertissement. La première compilation complète avait aussi exposé 79 avertissements dans le code existant ; ce baby step ne modifie pas ces fichiers.
- Revue indépendante : aucun défaut concret trouvé dans le périmètre.

Aucune mesure de performance nouvelle : il n’y a pas encore de code natif émis par cette représentation. Les tests des valeurs limites vérifient leur conservation dans l’IR, pas la sémantique arithmétique F# future.

## Reproduction

Depuis `/Users/0x1/Documents/htdocs/sharpurs/sharpurs` :

```sh
npm run build --silent
spago test
```

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-sharpurs` :

```sh
node scratch/sharpurs-int-kernel-20260908/convert-dump.mjs
```

Ce dernier script utilise le dump structurel daté de 2.1 ; le test du backend lui-même est autonome et ne dépend pas du dossier scratch. Les logs, le résultat de conversion et les empreintes sont conservés ici.

## Prochain baby step

Émettre ce noyau depuis `backendMod`, préserver son wrapper public `obj` et utiliser le générateur existant pour les bindings rejetés. La traduction de `IntModulo` et des entiers 32 bits doit être validée avant de revendiquer un gain. Seuls le dump et la représentation sont cochés au point 2 du todo.
