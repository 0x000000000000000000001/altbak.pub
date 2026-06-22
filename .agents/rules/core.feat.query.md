---
trigger: glob
globs: src/Core/Feat/*/Message/Query/**./*.purs
---

Tout ce qui est demandé dans la requête doit être un Need. Tout ce qui est retourné doit être un Return. Si l'un est présent d'un côté, il doit l'être de l'autre.

## Mutualisation des types (Opt / Result)
Il est fondamental d'identifier les structures partagées entre différentes queries d'une même feature pour les centraliser dans des dossiers dédiés (souvent sous `src/Core/Feat/.../Message/Query/_/Opt.purs` pour les besoins ou `src/Core/Feat/.../Message/Query/_/Result/` pour les retours).

**Exemple imaginaire :**
Si `GetPost` et `SearchPosts` ont tous les deux la possibilité de remonter des informations sur l'auteur du post :
1. **NE PAS** dupliquer `type AuthorOpt_ = { id :: Need Ɩ, name :: Need Ɩ }` dans les `Payload.purs` ou `Needs.purs` de chaque query. Centraliser cela dans un `Opt.purs` partagé.
2. **NE PAS** dupliquer `type Author = { id :: Return AuthorId, name :: Return String }` dans les `Result.purs` de chaque query. Créer un fichier `src/Core/Feat/Blog/Message/Query/_/Result/Authors/Authors.purs` (qui définira le type et potentiellement ses Lenses) et l'importer dans `GetPost.Result` et `SearchPosts.Result`.