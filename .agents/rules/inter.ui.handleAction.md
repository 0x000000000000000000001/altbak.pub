---
trigger: glob
globs: src/Inter/Ui/**/HandleAction/**/*.purs
---

### 1. Découpage des Actions Halogen (`handleAction`)
- **Fichiers dédiés et Routeurs (`Index.purs`)** : La fonction `handleAction` principale d'un composant ne doit JAMAIS contenir de logique en ligne. Elle doit obligatoirement être un routeur nommé `HandleAction/Index.purs` (et non `HandleAction.purs`).
- **Logique Récursive / Profonde** : Chaque cas (`case`) de l'action DOIT être extrait dans son propre fichier dédié. Si une action contient des sous-cas (comme un constructeur recevant un type `Input.Output`), cette règle s'applique récursivement :
  *Exemple* :
  1. Le cas `HandleSearchInputOutput` dans `HandleAction/Index.purs` route vers `HandleAction/HandleSearchInputOutput/Index.purs`.
  2. Ce sous-routeur `HandleSearchInputOutput/Index.purs` dispatche les sous-cas (`ChangedValue`, `Focused`) vers leurs fichiers terminaux dédiés (ex: `HandleSearchInputOutput/ChangedValue.purs`).
  3. **Nommage explicite (Outputs et actions imbriquées)** : Pour les sous-actions ou les gestions d'Output (ex: `HandleLinkOutput`), le nom de la fonction terminale DOIT refléter l'arborescence logique pour éviter les conflits (ex: `handleSearchInputFocused` ou `handleSearchInputOpenSearchQueryChanged`). **En revanche, le fichier lui-même conserve un nom simple** (ex: `Focused.purs` ou `OpenSearchQueryChanged.purs`). On évite de répéter le mot `Output` dans le nom de la fonction si cela l'alourdit inutilement. *Note: Les actions de premier niveau (ex: `Initialize`, `FetchResults`) gardent des noms de fonction simples (`initialize`, `fetchResults`).*

### 2. Gestion des Outputs de Sous-Composants
- **Terminologie au Passé** : Un output représente un événement qui vient de se produire. D'un point de vue terminologique, son nom DOIT systématiquement être formulé au passé (ex: `SearchClosed`, `SearchOpened`, `Clicked` au lieu de l'impératif `CloseSearch` ou `OpenSearch`).
- **Pas de handler externe (`HandleFooOutput.purs`)** : Il est strictement interdit de créer un fichier à la racine du composant pour gérer l'output d'un composant enfant (ex: `HandleLinkOutput.purs`).
- **Action Dédiée** : L'output d'un enfant doit être récupéré via un constructeur de donnée explicite ajouté au type `Action` du parent (ex: `HandleLinkOutput Link.Output`).
- **Routage Standard** : Ce constructeur est ensuite routé via le `HandleAction/Index.purs` habituel, en respectant la règle 1 sur le découpage récursif (ex: `HandleAction/HandleLinkOutput/Index.purs`).
