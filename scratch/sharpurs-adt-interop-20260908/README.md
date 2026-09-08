# Compatibilité ADT entre modules — 8 septembre 2026

La fixture compile deux vrais modules PureScript : `AdtPilot` produit une union récursive F# avec des champs `Color * Tree * int * Tree`, et `AdtConsumer` conserve ses fonctions et valeurs publiques en `obj`. Le producteur utilise le pilote natif ; le consommateur utilise le générateur habituel avec un routage explicite des constructeurs importés vers leurs wrappers publics.

## Problème reproduit

`before/fsi.log` conserve l’échec du générateur antérieur sur la fixture minimale : **FS0001**, `Color` attendu mais `obj` reçu lors d’une construction saturée dans le consommateur. Celui-ci appelait directement le constructeur F# natif en lui passant des arguments boxed. Les sources réellement générées sont dans `before/`.

Les patterns n’exigent pas de modification pour ce jalon : `before/patterns-only.fsx` compile et vérifie les patterns imbriqués de couleur, arbre, entier et alias, avec partage du sous-arbre. Les conversions des patterns existants acceptent les champs natifs.

## Changement

`Sharpurs.CodeGen` conserve la table complète des arités pour reconnaître les constructeurs dans les patterns. Un registre distinct sélectionne les constructeurs dont les expressions doivent passer par leur fonction publique `obj -> obj`. Les références, applications saturées et partielles utilisent ainsi le wrapper qui convertit les champs à la frontière. Les autres constructeurs gardent leur génération existante.

`Sharpurs.AdtInterop` expose un `NativeProducer` opaque : seul un module entièrement accepté par `AdtKernel`, possédant effectivement tous ses wrappers de constructeurs, peut être enregistré. La traduction du consommateur contrôle les arités, les propriétaires des types et les collisions de noms après conversion des noms de modules en F#. Une sélection vide produit exactement le code du générateur habituel.

`Main` n’active pas encore cette sélection. L’intégration progressive des fonctions de RBTree reste le jalon suivant.

## Validation

Le test durable `sharpurs/sharpurs/tests/adt-interop.mjs` compile les fixtures avec le fork TAST puis le vrai Builder PBO, et exécute le F# effectivement émis. Il passe **13 contrôles du registre et 62 assertions F#**, sans avertissement F#.

Les assertions couvrent :

- Construction saturée et application partielle réutilisée, constructeur passé à une fonction polymorphe.
- Entiers aux deux bornes, projections et patterns imbriqués, couleurs et branches qui ne correspondent pas.
- Partage physique des sous-arbres et passage du résultat du consommateur au `depth` natif du producteur.
- ADT local boxed du consommateur contenant l’arbre natif du producteur, avec patterns traversant les deux représentations.
- Évaluation des arguments saturés dans l’ordre `1, 2, 3, 4` ; évaluation immédiate des arguments `1, 2` d’une application partielle, sans répétition lors de sa réutilisation. Les fonctions étrangères traçantes sont confinées au test.
- Refus des arités absentes ou incorrectes, des producteurs invalides ou dupliqués, des wrappers absents et des collisions de noms `Adt.Pilot` / `Adt_Pilot`.

Les suites précédentes passent : 43 assertions IntKernel et 6 de routage en PureScript ; 68 contrôles du noyau local et 43 assertions F# ; 695 contrôles Int ; 21 du runtime d’application ; 33 contrôles et 32 assertions F# du premier pilote ADT. Les journaux sont dans `validation/`.

`check-default.mjs` réexécute le vrai Builder et la génération habituelle dans ce dossier isolé. Les corps des **301 modules F#** correspondent exactement à ceux du benchmark existant (`default-generation.json`). Les empreintes des **355 fichiers générés** correspondent également à la dernière variante mesurée ; `validation.json` conserve cette vérification. Aucun fichier généré normal du benchmark n’est réécrit.

Ce jalon ne mesure pas de nouveau gain. La baseline officielle actuelle d’altbak.pub est **294,80 ms**, dont **213,760 ms** pour RBTree. La dernière série à cinq processus donnait **296,39 ms** de médiane. Le prochain jalon doit intégrer les premières fonctions natives dans RBTree avant de refaire une comparaison.

## Reproduction

Depuis `sharpurs/sharpurs`, après `spago build` ou `npm run build` :

```sh
PURS=/chemin/du/fork/purs DOTNET=/chemin/de/dotnet npm run test:adt-interop
```

`PRELUDE_SRC` peut préciser les sources Prelude ; le défaut utilise `.spago/p`. `ADT_INTEROP_ARTIFACTS` peut désigner un dossier où conserver les sources F#, le journal et les empreintes. Sans cette option, les fichiers et caches du test sont temporaires et supprimés après exécution.

`after/metadata.json` conserve les versions du compilateur/runtime et les empreintes des sources du backend, du test et du F# exécuté. La baseline `before/` contient la fixture minimale qui reproduit l’échec ; la fixture durable dans `after/` ajoute les cas de partage, mélange de représentations et ordre d’évaluation décrits ci-dessus.
