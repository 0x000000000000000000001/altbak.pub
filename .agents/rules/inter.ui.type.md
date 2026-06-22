---
trigger: glob
globs: 
  - src/Inter/Ui/**/Type.purs
  - src/Inter/Ui/**/Type/**/*.purs
---

### 1. Élimination de `Void` / `Unit`
- **Sécurité Types (Slots)** : Banni d'utiliser `Unit` ou `Void` nus dans les signatures Halogen. Toujours recourir aux abstractions standard `Inter.Ui.Type` : `NoInput`, `NoOutput`, `NoQuery` (avec `ι`).

### 2. Fichier vs Dossier (Cohérence d'arborescence des Types)
- **Exclusivité** : Il est strictement interdit d'avoir un fichier `Type.purs` et un dossier `Type/` simultanément. Il ne doit pas y avoir non plus de `Type/Index.purs` qui sert d'agrégateur.
- **Fichier unique** : Si le domaine est simple, on utilise un fichier unique `Type.purs`.
- **Dossier éclaté (Splitté)** : S'il est nécessaire de scinder les types (ex: `Type/State.purs`, `Type/Action.purs`), on crée un dossier `Type/`. Dans ce cas, **il n'y a plus aucun fichier centralisateur** (ni `Type.purs`, ni `Index.purs`). Les fichiers externes doivent importer spécifiquement le type dont ils ont besoin (ex: `import ...Type.Action (Action)`). Les types utilitaires restants (comme `Query`, `UiM`) doivent être placés dans leurs propres fichiers (ex: `Type/Query.purs`).

### 3. State et Input
- **Stockage de l'Input complet** : De manière générale, un `State` qui se contente de "mirror" (copier un à un) les champs de l'`Input` ne doit pas exister. Lorsqu'un composant a besoin des données de l'`Input` (notamment pour les comparer dans un `Receive`), on ne doit pas stocker des champs isolés. À la place, on stocke systématiquement l'objet `input` complet sous le champ `input` dans le `State` afin de faciliter la mise à jour et la comparaison (`oldInput <- gets _.input`, `when (oldInput /= input) ...`).
