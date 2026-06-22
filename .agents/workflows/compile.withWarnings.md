---
description: Fait compiler le code en forçant l'affichage de tous les warnings
---

Fais compiler le code de façon à voir l'intégralité des warnings. 
Puisque le compilateur met en cache les fichiers non modifiés, tu dois d'abord nettoyer le dossier `output` pour forcer le compilateur à tout revérifier.

// turbo-all
Exécute la commande `rm -rf output && zsh -lc "./bin/build"` pour déclencher ce build complet.

Une fois que tu as la liste exhaustive de tous les avertissements ou erreurs (e.g. imports non utilisés), corrige-les.
