---
trigger: glob
globs: src/Inter/Ui/**/*.purs
---

### 1. Sémantique et Abstraction des Noeuds HTML
- **Abstraction systématique** : Chaque nœud DOM DOIT être abstrait dans son propre module (`Foo.purs`).
- **Pas de BEM simulé** : Interdit de styliser un enfant via concaténation (ex: `staticClass <> "-thumb"`). Un enfant avec style DOIT avoir son propre `Style/Style.purs` et `generateStaticClass`.
- **Pas de HTML brut** : Remplace `div [ class_ "col" ]` par un sous-composant `Column/Column.purs` exportant `column`.
- **Pas de BEM dynamique inline** : Pour les variantes (ex: `--portrait`), utilise `refineClass' staticClass "portrait"` et applique-le via une constante (`staticClassWhenPortrait .? do ...`). Pas d'interpolation (`staticClass <> "--" <> format`).
- **Data > Classes (Délégation tardive)** : Ne pré-calcule pas des classes (`formatClassName`) depuis un parent vers un enfant. Fais descendre la donnée brute (`format`) le plus bas possible. Seule la fonction terminale `thumb` avec props (`Style.purs`) l'évalue pour déduire la classe CSS.
- **Convention Halogen `_` (Conflits de noms)** : Le wrapper DOM initial (`Style.purs`) DOIT exporter 2 versions systématiquement :
  1. Sans suffixe avec props (`IProp`) : `authors props = div ([ class_ staticClass ] <> props)`
  2. Avec suffixe `_` sans props : `authors_ = authors []`
  **Import** :
  - Sans props : importe `authors_` (l'alias `Style.` est interdit).
  - Avec props : importe `...Style as Style` et appelle `Style.authors` (obligatoire pour gérer l'homonymie).
  *Exception* : Le composant principal garde un nom pur (`text str =`, pas `text_`). S'il utilise un nœud natif, c'est l'import Halogen qui est aliasé (`HH.text`).
- **Séparation Styles / Attributs (Injection par props)** : Dans les fichiers `Style.purs`, les fonctions d'encapsulation (sans ou avec `_`, ex: `thumb props = ...`) NE DOIVENT traiter et injecter QUE les classes CSS (ex: `class_ staticClass`). Elles ne doivent JAMAIS accepter explicitement en argument, ni injecter d'attributs de contenu HTML (comme `src`, `href`, `alt`, ou le contenu texte/enfants tel que `[ text str ]`). 
  **Clarification** : `Style.purs` n'est qu'un raccourci (shortcut) permettant de masquer la plomberie des classes CSS pour le parent. Tout le reste relève d'une logique DOM parent/enfant qui ne le regarde pas. Bien que `authors_ = authors []` accepte implicitement un tableau d'enfants via currying, c'est toujours le composant logique (`Foo.purs`) qui garde l'entière responsabilité de déclarer, assembler et gérer les noeuds finaux (ex: `text str`) ou la structure des enfants.
- **Imports allégés** : Importe les composants explicitement sans alias (`import ... as Thumb` banni) pour un rendu fluide `[ thumb ..., content ... ]`.
- **Rendu Sémantique** : L'arbre final doit être déclaratif et immaculé : `frontPage_ [ column id false [ ... ] ]`.

### 2. Arborescence vs DOM
- **Reflet DOM strict** : Les dossiers s'alignent fidèlement sur l'arbre DOM (ex: l'enfant `Title/` dans `Content/`).
- **Extraction en Composants Feuilles** : Tout nœud logique (même trivial comme un `Label` ou un `Field`) DOIT être extrait dans son propre composant feuille (`Label.purs`, `Field.purs`) dès lors qu'il porte des propriétés spécifiques (comme des event handlers `onClick`, `onFocus`, un `Ref`) ou qu'il assemble des enfants. Le composant parent s'allège en appelant simplement ce composant (ex: `label s`) et délègue au composant terminal la responsabilité de déclarer l'ensemble de ses props (en dehors de l'injection CSS du parent) et la composition de son contenu.
- **Arbre en Dur (Pas de children/nodes en arguments)** : Il est STRICTEMENT INTERDIT de passer un argument générique contenant des noeuds enfants (ex: `nodes :: Array HTML` ou `items`) à un composant. C'est un anti-pattern.
  - Le parent (`Core` par exemple) ne doit pas savoir exactement de quoi l'enfant est constitué en interne, il doit juste l'instancier et savoir qu'il est présent.
  - Le composant enfant assume ses propres imports et compose LUI-MÊME ses enfants. Par exemple, au lieu que le parent assemble et passe un tableau d'items en argument (`[ Items.items id open items ]`), le parent appelle juste `[ Items.items id open ]`. C'est le module `Items` qui importe et assemble en dur ses sous-éléments (`Item.item`, `separator`) : `items id open = Style.items id open [ Item.item ..., separator ]`.
- **Responsabilité du Composant Enfant (Component vs Module Final)** : Lorsqu'un composant (ex: `Menu`) est instancié par un parent (ex: `Router`), l'instanciation suit un pattern strict :
  - Le fichier `Component.purs` de l'enfant ne DOIT exposer que la machinerie Halogen (`component`) et les raccourcis d'instanciation du slot (`menu_` et `menu`). Il reste agnostique du contexte métier du parent.
  - Le fichier portant le nom du composant (ex: `Menu.purs`) DOIT exporter la fonction finale de rendu (`menu :: ParentState -> ComponentHTML ...`). C'est ce module qui effectue la translation "intelligente" entre l'état du parent et les `Input` exigés par l'enfant. Cela évite au parent de polluer son propre `Render.purs` avec de la logique de préparation d'inputs complexes.

### 3. Ruissellement de l'État (State Prop-Drilling)
- **Transmission du State Intégral** : Lorsqu'on transmet des données depuis le composant racine jusqu'aux descendants (rendu, style, etc.), il FAUT privilégier le passage du `State` complet plutôt que de passer chaque propriété individuellement.
  Cela laisse la responsabilité à chaque composant descendant de déstructurer ce dont il a besoin (ex: `core { id, fold, search: { open } } props = ...`) et évite de refactoriser sans cesse la liste des arguments dans toute la chaîne de composants lors de l'ajout ou du retrait d'une dépendance.

### 4. Gestion des Données et Requêtes
- **Remote Data** : Toujours wrapper la donnée distante via `Inter.Ui.Type.Remote` (constructeurs `NotAsked`, `Loading`, etc.).
- **Requêtes (Queries)** : Utiliser des ADT/structures strictes au lieu de booléens (ex: `Needed $ Unfolded {...}`, `Needed ι`, ou `NotNeeded`).

### 5. Gestion de l'Input et Lifecycle (Receive)
- **Conservation de l'Input (Pattern Receive)** : Lorsque l'on a besoin de stocker des props passées en argument au composant (le `Input`) uniquement pour vérifier les changements dans le hook `Receive` (et déclencher des rechargements), il faut stocker l'`Input` dans son intégralité dans le `State` via une clé `input :: Input`. Il ne faut pas le déstructurer en de multiples propriétés éparpillées dans le state (ex: stocker un `slug :: String`). Les autres propriétés du `State` doivent être réservées aux éléments qui évoluent localement indépendamment de l'`Input`.
