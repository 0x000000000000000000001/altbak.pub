# Fixture minimal TypeApp — 8 septembre 2026

Le problème est reproduit sur **un module de 13 lignes sans dépendance**, compilé par le fork PureScript local, puis optimisé par le vrai pipeline PBO utilisé par Sharpurs. Aucun code de l’optimiseur ou du générateur n’a été modifié.

`TypeAppFixture.purs` définit une identité polymorphe, deux usages concrets et un témoin générique distinct :

```purescript
identity :: forall a. a -> a
useInt :: Int -> Int
useString :: String -> String
generic :: forall b. b -> b
```

Les trois dernières fonctions appliquent la même `identity` à leur argument. Le témoin utilise volontairement `b`, afin de rendre visible la substitution `a → b` attendue.

## Résultat observé

| Binding | TypeApp en entrée | Annotations du corps après inlining | Attendu |
| --- | --- | --- | --- |
| `useInt` | `Int` | `Typed Int (Typed a (Local 0))` | seulement `Int` |
| `useString` | `String` | `Typed String (Typed a (Local 0))` | seulement `String` |
| `generic` | `b` | `Typed b (Typed a (Local 0))` | seulement `b` |
| `identity` | aucun | `Typed a (Local 0)` | `a`, inchangé |

Les quatre corps optimisés sont bien des identités : une abstraction renvoie son propre argument, sans appel global, application ni `TypeApp` résiduel. Les signatures externes des usages sont correctes. Le défaut apparaît dans les annotations internes, où la variable `a` du callee subsiste après les trois instanciations.

Le checker inspecte donc tous les `Typed` avant l’abstraction et dans son corps. Il ne fixe pas le nombre de wrappers pour le contrat futur : toute annotation restante doit être cohérente avec le type attendu.

## Vérifications exécutées

- Compilation du seul fixture par le binaire du fork `purescript` : succès, sans erreur ni avertissement.
- Exécution du Builder PBO avec les mêmes options et le même filtre de sémantiques étrangères que Sharpurs : **1 module, 4 bindings capturés**.
- Vérification de l’entrée parsée : `identity @Int`, `identity @String`, `identity @b`, tous vers le même nom qualifié ; déclarations génériques quantifiées en `a` et `b`.
- Vérification de la forme des quatre valeurs optimisées : une identité, sans appel résiduel.
- `python3 check-fixture.py` : **échec attendu, code 1**, trois bindings violent le contrat de typage. Le témoin `identity` passe. `contract.log` conserve ce résultat.
- `python3 check-fixture.py --expect-current` : **succès, code 0** ; vérifie exactement les trois régressions observées, ainsi que l’identité originale restée générique. Ce mode valide la reproduction du défaut, pas la justesse du typage.
- Les empreintes des 353 fichiers générés du benchmark restent identiques. Aucun benchmark ni génération F# exécuté.

## Portée du prochain correctif

Le fixture isole trois substitutions manquantes : **`a → Int`, `a → String`, `a → b`**. La prochaine micro-étape est de propager l’instanciation de chaque `TypeApp` dans les annotations du corps inliné, puis de faire passer le checker normal. Les instanciations doivent rester indépendantes et la définition générique originale inchangée.

Un cas avec masquage de variable par un `ForAll` interne devra accompagner ce correctif pour vérifier la portée. Le présent fixture ne prouve pas encore ce cas.

Observation distincte : les `ForAll` figurent sur les annotations des déclarations d’entrée, alors que leurs abstractions portent `Func[a]a` ou `Func[b]b`. Le chemin `toTopLevelBackendBinding` utilise l’annotation de l’expression, et les racines optimisées observées sont donc ces `Func`. Ce fixture ne traite pas la transmission des quantificateurs ; il n’exige pas leur absence dans le contrat futur.

## Reproduction

Depuis ce dossier :

```sh
../../run/bak/js/node_modules/.bin/purs compile TypeAppFixture.purs --output output --codegen corefn,js
node --expose-gc --stack-size=65536 --max-old-space-size=16384 dump-fixture.mjs
python3 check-fixture.py --expect-current
python3 check-fixture.py
```

La dernière commande échoue tant que les annotations ne sont pas correctement instanciées. Le compilateur appelé est le lien vers le binaire du fork local, vérifié avant l’expérience ; son chemin, sa version et son empreinte figurent dans `metadata.json`.

`input.parsed.json` conserve le tcorefn avec types résolus ; les quatre `*.optimized.json` proviennent directement de `onCodegenModule`. `summary.json` contient les contrats observés, `compile.log`, `run.log`, `contract.log` et `reproduce.log` les exécutions. Les dossiers `output/` et `.purmeta/` sont isolés, ignorés par Git et reproductibles. Aucun gain de performance n’est revendiqué à ce jalon.
