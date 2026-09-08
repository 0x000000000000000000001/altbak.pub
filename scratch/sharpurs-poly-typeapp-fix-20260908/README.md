# Polymorphism après correction TypeApp — 8 septembre 2026

Dans le vrai `act` optimisé, le `LetRec go` est désormais annoté **`Int -> Int -> Int`**. Son sous-arbre, appel initial compris, ne contient plus de `TypeVar`. La fonction générique `polyLoop` conserve son quantificateur et ses deux accès au dictionnaire.

Le [checker](inspect-poly.py) vérifie les mêmes propriétés de calcul que le [dump initial](../sharpurs-poly-ir-20260908/README.md) : deux paramètres, récursion locale terminale saturée, comparaison entière à zéro, soustraction de 1, addition de 1, appel initial `go dummy 0`. L’action conserve les mêmes quatre références globales et aucun accès au dictionnaire. Les données d’entrée décodées de Polymorphism sont identiques à celles de la preuve initiale.

Une différence volontaire : `Bench.opaque @Int` reste maintenant explicite dans l’IR, car cette FFI n’a pas de corps à instancier. Le checker exige cette référence précise. La boucle locale ne contient aucun `TypeApp` résiduel.

Les deux annotations d’effet en variables `a`/`b` restent présentes autour de `Effect.pureE`/`Effect.bindE`. Elles proviennent de champs de dictionnaires référant directement à ces fonctions sans `TypeApp` et ne font pas partie de la boucle spécialisée. Leur traitement relève d’une étape distincte.

Preuves : [assertions.log](assertions.log), [arbre de act](act.tree.txt), [boucle extraite](go.inlined.json), [summary.json](summary.json) et [metadata.json](metadata.json). Le dump a traversé 301 modules de l’output courant. Les 355 fichiers F#/C#/projets contrôlés sont inchangés.

```sh
# Depuis ce dossier, après compilation de Sharpurs avec le PBO corrigé :
node --expose-gc --stack-size=65536 --max-old-space-size=16384 dump-poly.mjs
python3 inspect-poly.py
```

Le callback utilisé observe l’entrée réelle de Sharpurs, force l’optimisation et écrit les caches `.purmeta` dans ce dossier isolé. Il ne génère pas de F#. Les anciennes preuves restent intactes. Aucun gain de benchmark revendiqué : la prochaine étape consiste à émettre cette récursion locale en F# natif puis à la mesurer.
