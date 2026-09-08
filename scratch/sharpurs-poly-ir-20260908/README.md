# Polymorphism : état du PBO optimisé — 8 septembre 2026

**PBO spécialise déjà les valeurs nécessaires à la boucle dans `act`.** Il inline `polyLoop` avec le dictionnaire `intMonoidish` : le corps récursif contient une addition entière de `1`, sans accès au dictionnaire. Le travail restant porte sur les annotations de types après instanciation et sur la génération de cette boucle locale depuis le corps optimisé de l’action.

Cette étape capture et analyse le callback réel `onCodegenModule`, avec les mêmes options que Sharpurs. Aucun backend n’a été modifié, aucun F# régénéré, aucun benchmark exécuté. Les empreintes des 353 fichiers générés contrôlés restent identiques.

## Preuve en entrée TAST v3

Le JSON parsé conserve `classDecls`, les quantificateurs, contraintes et instanciations. `call.input.json` isole l’appel suivant, avant optimisation :

```text
polyLoop @Int Test.Polymorphism.intMonoidish dummy 0
```

- L’argument du `ExprTypeApp` est `Int` ; son annotation est `ConstrainedType(Monoidish Int, Func[Int, Int] Int)`.
- Le dictionnaire est un `Var` qualifié `Test.Polymorphism.intMonoidish`, annoté `Monoidish Int`.
- `classDecls` décrit `Monoidish a`, avec `mempty_ : a` et `mappend_ : a -> a -> a`.
- L’instance définit `mempty_ = 1` et une addition `Int`. Après optimisation de l’instance, ces membres sont un `LitInt 1` et un `PrimOp (OpIntNum OpAdd)`.

Les informations nécessaires sont donc bien fournies par le tcorefn. Le type `Int` et l’identité du dictionnaire sont tous deux établis.

## Deux formes différentes après optimisation

| Binding observé | Dictionnaire | Boucle |
| --- | --- | --- |
| `polyLoop` générique | Paramètre `dictMonoidish`, deux `Accessor` (`mappend_`, `mempty_`) | `LetRec go`, accumulateur `a` |
| `act` à l’instance `Int` | Aucun `Accessor`, aucune référence à `polyLoop` ou `intMonoidish` | `LetRec go`, primitives entières égalité/soustraction/addition |

Le sous-arbre spécialisé de `act` correspond structurellement à :

```text
let rec go n acc =
  if n == 0 then acc else go (n - 1) (acc + 1)
in go dummy 0
```

Le groupe local contient une seule fonction, au niveau lexical 1 ; ses paramètres sont aux niveaux 2 et 3. L’auto-appel est local, terminal et saturé avec deux arguments. L’entrée opaque `10000000` et la conversion du résultat en chaîne restent dans l’action.

Les seules références globales résiduelles de `act` sont `Effect.bindE`, `Bench.opaque`, `Effect.pureE` et `Data.Show.showIntImpl`. Le dump contient un `LetRec`, trois primitives entières, aucun `TypeApp` résiduel et aucun dictionnaire. `inspect-poly.py` vérifie ces propriétés et la forme de l’auto-appel par assertions.

## Point à corriger avant génération native

Les valeurs ont été spécialisées, mais certaines annotations du corps inliné restent celles de la fonction générique. `act.tree.txt` montre notamment :

```text
Typed [Int, TypeVar(a)] -> TypeVar(a)
  Abs n
    Typed [TypeVar(a)] -> TypeVar(a)
      Abs acc
        ...
        Typed TypeVar(a)
          Typed TypeVar(a)
            Typed Int
              PrimOp OpIntNum(OpAdd) acc 1
```

Le sous-arbre de `go` contient encore 11 occurrences de `TypeVar(a)`. Les annotations externes de son résultat sont pourtant `Int`. D’autres annotations génériques subsistent également autour des applications d’effets dans `act`.

La lecture ciblée de PBO montre le mécanisme concerné :

- [Convert.purs:549](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-sharpurs/src/PureScript/Backend/Optimizer/Convert.purs:549) reprend `ann.type` dans `Typed` ; à la ligne 613, il conserve `ExprTypeApp` sous forme `Syn.TypeApp`.
- [Semantics.purs:323](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-sharpurs/src/PureScript/Backend/Optimizer/Semantics.purs:323) représente l’instanciation par `SemTypeApp`.
- [evalApp:477](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-sharpurs/src/PureScript/Backend/Optimizer/Semantics.purs:477) traverse `SemTypeApp _ fn` sans appliquer une substitution aux annotations du corps.
- [quote:1429](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-sharpurs/src/PureScript/Backend/Optimizer/Semantics.purs:1429) réémet les `SemTyped ty` tels quels.

Le dump établit le résultat observé, et cette lecture identifie le traitement incomplet des instanciations. La branche particulière d’inlining des références externes n’a pas été tracée pendant cette étape.

Le convertisseur `IntKernel` actuel rejette effectivement tous les bindings du module. Il exige des paramètres et un résultat `Int`, et traite un auto-appel global ; il ne prend pas encore en charge ce `LetRec` local ni l’enveloppe d’effets de `act`. Ignorer les annotations `TypeVar` masquerait le problème au lieu d’exploiter correctement le TAST.

## Prochaines micro-étapes justifiées

1. Isoler la propagation des instanciations dans un petit fixture PBO : même fonction polymorphe appliquée à `Int` et `String`, avec une version générique témoin. Vérifier si des `Typed a` restent autour de résultats concrets. Ce fixture doit précéder le correctif.
2. Propager la substitution de types liée à chaque `TypeApp` dans les annotations du corps inliné, en respectant la portée des quantificateurs. Un remplacement global de toutes les variables nommées `a` serait incorrect. Rejouer ensuite le vrai dump de `act`.
3. Étendre la génération Sharpurs au `LetRec` local devenu entièrement typé, dans son enveloppe d’appels et d’effets optimisés, avec repli complet pour les formes non reconnues. Les primitives du corps sont déjà couvertes par le noyau Int.
4. Vérifier `10000000`, les cas génériques et au moins deux instanciations, puis mesurer l’action complète selon le protocole validé sur TCO.

Une nouvelle spécialisation des dictionnaires n’est pas nécessaire pour cet appel : PBO a déjà fait ce travail sur les valeurs. La représentation typée d’un dictionnaire encore dynamique reste une piste séparée pour les autres cas.

## Reproduction et fichiers

Depuis la racine d’`altbak.pub-sharpurs`, avec les sorties PureScript et PBO déjà compilées :

```sh
node --expose-gc --stack-size=65536 --max-old-space-size=16384 scratch/sharpurs-poly-ir-20260908/dump-poly.mjs
python3 scratch/sharpurs-poly-ir-20260908/inspect-poly.py
```

Le script optimise les 299 modules pour disposer des implémentations importées, puis capture seulement les six bindings de `Test.Polymorphism`. Son cache `.purmeta` reste isolé et ignoré. Le sérialiseur conserve les tags structurels `Typed`/`TypeApp`, que le rendu habituel des traces masque.

Les `*.optimized.json` sont les dumps structuraux intégraux ; les `*.tree.txt` sont leurs vues compactes avec annotations conservées. `input.parsed.json` contient le module d’entrée avec types résolus, `call.input.json` l’instanciation ciblée, `go.inlined.json` le sous-arbre local, et `summary.json` les comptages vérifiés. `metadata.json` conserve révisions, empreintes et résultats du convertisseur existant. Les observations de cette étape n’établissent aucun nouveau gain de performance.
