---
trigger: glob
globs: 
  - src/Inter/Ui/**/Style.purs
  - src/Inter/Ui/**/Style/**/*.purs
---

### 1. Gestion des Styles (Statique vs Dynamique)
- **Styles Statiques (Bottom-Up)** :
  Définis via `staticStyle :: CSS.CSS` (`Style.purs` ou `Style/Style.purs`). Agrégés vers le haut dans `Style/Index.purs`. Injectés globalement une fois.
  *Règles* : Ne jamais nommer l'agrégation `sheet` (réservé aux styles Halogen dynamiques). Ne jamais importer un composant local sous `as Core` (utilise le nom réel, ex: `as Thumb`).
- **Structure des dossiers de Style (Index vs Style unique)** :
  - **Composant Feuille** (pas de sous-composants visuels à agréger) : Il ne possède **PAS** d'index. Son style est défini dans un fichier `Style.purs` unique, placé directement au premier niveau de son dossier logique (ex: `MatchingWord/Style.purs`).
  - **Composant Parent** (possède des sous-composants à agréger) : Il possède obligatoirement un index. Dans ce cas, **absolument tout** ce qui a trait au style doit être rangé dans un sous-dossier `Style/`. Son propre style devient `Style/Style.purs` et l'agrégation se fait dans `Style/Index.purs`. *Il est strictement interdit d'avoir un fichier `Style.purs` et un dossier `Style/` qui cohabitent au même niveau.*
- **Délégation stricte des arbres de styles (Index.purs)** : Même si un dossier logique n'a pas besoin de styliser de composant (ex: un dossier de logique métier qui regroupe des sous-composants visuels), il DOIT posséder un fichier `Style/Index.purs` chargé de collecter et remonter les styles de ses sous-dossiers. Ainsi, chaque parent collecte systématiquement l'index de ses enfants directs, sans jamais sauter de niveau dans l'arborescence, garantissant une lecture fluide et prédictible de l'agrégation.
- **Styles Dynamiques** :
  Le staticStyle étant évalué 1 fois, interdit d'y placer du CSS dynamique par instance. Interdit d'utiliser `style="..."`. Préfère les raffinements de classes (ex: `refineClass`, `inferInstanceClass` de `Util.Style.Style`).

### 2. Scoping CSS / InstanceId
- **État via `withId`** : L'état `initialState` Halogen DOIT être initialisé avec `withId` (ex: `withId $ κ { ... }`), créant un `InstanceId` unique.
- **Propagation de l'ID** : Cet ID doit descendre jusqu'au DOM via `class' id` pour cloisonner l'application CSS et empêcher les conflits entre instances.

### 3. DSL CSS et Utilitaires
- **Réflexe `Util.Style.Style`** : Toujours utiliser les utilitaires du projet avant les primitives natives (`ClassName`, `classes []`). Ex: l'utilitaire `classes` prend directement un tableau de chaînes.
- **Interdiction de `raw` / CSS Natif** : L'usage de `raw "max-width" ""` ou des primitives Purescript (`maxWidth (rem 6.0)`) est strictement interdit si un raccourci existe (ex: `maxWidthRem 6.0`).
- **Base Vanilla Interdite** : Préférer les alias structurés existants (`widthPct100`, `displayFlex`).
- **Tout en `rem`** : Unités `em` et `px` bannies (sauf nécessité typo). Utiliser prioritairement les helpers incluant déjà l'unité (`paddingRight 2.4`).
- **Arborescence SCSS simulée** : 
  Ouvrir avec `.?`. L'utilisation de l'opérateur `?` avec des sélecteurs concaténés inline (ex: `(staticClass .& after) ? do`) est **strictement interdite**. Tous les sélecteurs complexes (variantes, pseudo-classes comme `hover`, ou pseudo-éléments comme `after`, `before`) DOIVENT systématiquement être déportés dans un bloc `where`, appelés avec l'opérateur `:?` et un préfixe `__` imitant l'indentation SCSS.
  Les petits-enfants ajoutent `__` par profondeur (ex: `____article`).
  ```purescript
  staticStyle = do
    staticClass .? do
      -- styles principaux
      __after :? do
        -- styles du pseudo-élément
    where
      __after = staticClass .& after
  ```
- **Imports CSS allégés (Pas de préfixe `CSS.`)** : Importer les fonctions CSS (`color`, `flexGrow`, `maxWidth`, etc.) explicitement et éviter le préfixe `CSS.` quand c'est possible et que cela ne cause pas de conflit (ex: conflits avec `top` ou `bottom` de Purescript/Proem).

### 4. Génération de classes statiques
- **Constante `fullModuleName`** : Au lieu de passer le nom du module directement en chaîne de caractères dans `generateStaticClass`, on doit toujours déclarer une constante `fullModuleName :: String` (contenant le nom complet du module, ex: `"Inter.Ui.Page.MonComposant.Style.Style"`) en haut du fichier, et l'utiliser comme argument. Cela clarifie l'origine de la classe générée.
