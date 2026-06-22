---
description: Applique une rule
---

Applique la rule dont j'ai donné le chemin du fichier. Les sous-rules doivent aussi être appliquées là où c'est pertinent. Par exemple, si je dis d'appliquer inter.ui.md dans un dossier, il faudra appliquer aussi inter.ui.handleAction dans le dossier HandleAction, ou inter.ui.type dans le fichier Type ou le dossier Type. Si je dis d'appliquer la règle racine (_.md), il faut appliquer inter.ui aussi (qui est fille de la racine), et ses filles pertinentes (celles de inter.ui), etc. Tu peux t'appuyer sur les globs des règles pour savoir qu'elles filles sont pertinentes à chaque fois.