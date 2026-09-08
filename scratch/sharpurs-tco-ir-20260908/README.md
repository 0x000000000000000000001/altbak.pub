# Baby step 2.1 — observer le binding optimisé de TCO

Le 8 septembre 2026, exécution effective du builder PBO utilisé par Sharpurs, avec ses mêmes directives, limite de réécriture et filtre de sémantiques étrangères. Le callback `onCodegenModule` capture uniquement `Test.TCO.deepTailRec`. Les 299 modules chargés sont optimisés pour disposer du même contexte d’imports ; aucune génération F#, compilation ou exécution du benchmark n’a lieu.

Le cache `.purmeta` de ce passage est isolé dans ce dossier scratch et ignoré par Git. Aucun code du backend ou de PBO n’est modifié.

## Résultat observé

- Signature résolue : `Func [Int, Int] Int`, soit `Int -> Int -> Int`.
- Groupe marqué `recursive: true`.
- Deux `Abs` imbriquées, chacune avec un paramètre : `v` au niveau lexical 0, `v1` au niveau 1.
- Une branche teste l’égalité entière du premier paramètre à zéro et retourne le second.
- Sinon, appel saturé à `Test.TCO.deepTailRec` avec deux arguments : la soustraction entière `v - 1` et l’addition entière `v1 + modInt(v, 3)`.
- Cet appel est terminal. La récursion est portée par le groupe et la référence globale ; aucun `LetRec` local n’est présent.
- Aucun accès à un dictionnaire, appel FFI ou effet dans ce binding optimisé. Sa seule référence qualifiée est sa propre fonction.

Notation du résultat observé (pseudocode, pas un émetteur F#) :

```text
deepTailRec : Int -> Int -> Int
deepTailRec(v, v1) =
  if eqInt(v, 0) then v1
  else deepTailRec(subInt(v, 1), addInt(v1, modInt(v, 3)))
```

Le binding d’entrée contient 3 nœuds d’expression `TypeApp` ; le binding optimisé en contient 0. Ce cas est monomorphe et ses appels arithmétiques sont déjà ramenés aux primitives entières. Cette absence résiduelle ne signifie pas que le TAST v3 est perdu.

## Sous-ensemble minimal pour le prochain baby step

| Nœud d’expression | Nombre | Usage |
| --- | ---: | --- |
| `Typed` | 17 | Types structurels résolus ; enveloppes parfois répétées |
| `Abs` | 2 | Deux paramètres curriés avec niveaux lexicaux |
| `Branch` | 1 | Condition et branche alternative |
| `PrimOp` | 4 | Égalité, soustraction, addition, modulo entiers |
| `App` | 1 | Appel récursif saturé |
| `Var` | 1 | Référence qualifiée à la fonction elle-même |
| `Local` | 5 | Références aux deux paramètres |
| `Lit` / `LitInt` | 3 | Constantes 0, 1 et 3 |

Les opérateurs exacts sont `Op2 (OpIntOrd OpEq)`, `Op2 (OpIntNum OpSubtract)`, `Op2 (OpIntNum OpAdd)` et `Op2 (OpIntNum OpMod)`.

Certains nœuds synthétisés n’ont pas de wrapper `Typed` immédiatement au-dessus : le test d’égalité, le modulo, la référence récursive, un accès au second paramètre et le littéral zéro. Leurs types sont néanmoins déterminés par les opérateurs, la signature globale et les paramètres. Le futur émetteur doit donc transporter un environnement de types pour les niveaux lexicaux et connaître les signatures des primitives.

L’arité totale est deux, mais le premier `Abs` ne contient qu’un paramètre ; il faut traverser les `Typed` intermédiaires pour atteindre le second. Ce résultat n’autorise pas à traiter toute fonction renvoyée comme un argument supplémentaire dans les futurs cas polymorphes ou newtypes.

La traduction F# de `OpMod` devra respecter sa sémantique, notamment sur les entiers négatifs. Ce jalon observe l’opérateur ; il ne choisit ni ne valide encore son implémentation native.

## Prochaine action proposée

Définir une petite représentation F# typée couvrant exactement ce sous-ensemble pur fermé, avec rejet vers le générateur existant dès qu’un binding sort du sous-ensemble. L’émission native, les tests de sémantique entière et les mesures de performance appartiennent aux baby steps suivants.

Le seul point coché dans l’étape 2 de `sharpurs/todo.md` est l’extraction de ce dump. Aucun gain de performance nouveau n’est revendiqué.

## Preuves et reproduction

- `deepTailRec.optimized.json` : binding exact reçu par `onCodegenModule`, tags des constructeurs et types conservés.
- `summary.json` : inventaire des nœuds et références.
- `metadata.json` : commits, runtime Node et empreintes des modules compilés utilisés ainsi que du JSON d’entrée.
- `dump-tco.mjs` et `run.log` : script et confirmation d’exécution.

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-sharpurs` :

```sh
node --expose-gc --stack-size=65536 --max-old-space-size=16384 scratch/sharpurs-tco-ir-20260908/dump-tco.mjs
```

Le script utilise les modules JS compilés du backend, tels que référencés dans `metadata.json`. Il faut reconstruire le backend si ses sources ou PBO ont changé avant une nouvelle observation.

Le printer habituel de traces PBO masque volontairement `Typed` et `TypeApp` ; le dump structurel est conservé pour éviter de confondre une omission d’affichage avec une absence de type.
