# Émission du premier noyau Int — 8 septembre 2026

Le générateur consomme maintenant les bindings optimisés de `backendMod` acceptés par `Sharpurs.IntKernel.fromBinding`. Il remplace leur déclaration CoreFn à sa position initiale ; les bindings absents ou refusés gardent le générateur existant. Les groupes de récursion mutuelle restent entièrement sur le chemin existant, car leurs corps peuvent appeler les entrées `_tco` de leurs voisins.

Chaque binding accepté contient une fonction native locale, avec paramètres et retour `int`, et conserve son nom public et une valeur `obj`. Chaque étage du wrapper curry est `obj -> obj`. Les noms internes suivent les niveaux lexicaux et sont locaux au binding.

Sur les 301 fichiers F# présents avant cette génération, seuls `Test.TCO.fs` et `Sharpurs_Prelude.fs` changent. Le seul binding accepté est `Test.TCO.deepTailRec`, sans filtre sur son nom. La boucle générée ne contient ni `box`, ni `unbox`, ni dictionnaire, ni `sharpurs_apply` ; ces conversions restent à la frontière publique. Voir `Test.TCO.generated.fs` et les deux fichiers `.diff`.

## Vérifications exécutées

- `npm run build --silent` : backend et bundle compilés. Les 15 avertissements de cette compilation incrémentale concernent du code préexistant ; le nouveau module d’émission n’en produit pas.
- `spago test` : 43 assertions du convertisseur et 6 du routage passent. La reconstruction complète signale 78 avertissements préexistants, aucun dans les nouveaux tests ou modules IntKernel.
- `DOTNET=/Users/0x1/.dotnet/dotnet npm run test:kernel` : 695 vérifications F# passent, dont 225 couples de bornes/signes/zéro pour chacune des trois primitives. Les réponses attendues viennent des implémentations JS du Prelude installé. Les tests couvrent aussi curryfication, réutilisation des applications partielles, branches non évaluées et récursion avec permutation sur plus d’un million d’appels.
- `DOTNET=/Users/0x1/.dotnet/dotnet npm run test:runtime` : les 21 tests existants d’appels et d’exceptions passent avec le prélude modifié.
- Génération réelle par `../sharpurs/sharpurs/bin/sharpurs --main App`, puis compilation Release .NET 8 : aucune erreur ni avertissement.
- `check-tco.fsx` charge cette DLL : l’action réelle `Test.TCO.act` renvoie `"100000"`. Quatre autres vérifications passent, dont un million d’itérations et la réutilisation du wrapper partiel.

La première commande .NET, lancée depuis la racine via le symlink `output/Main/Program.fsproj`, a échoué sur une référence C# intermédiaire introuvable. La compilation depuis le chemin physique `run/bak/sharp/output/Main` a réussi sans changement de source. Les deux logs sont conservés.

## Modulo

`sharpurs_int_mod` applique le modulo euclidien PureScript : résultat non négatif et diviseur zéro donnant zéro. La garde du diviseur `-1` évite l’exception CLR sur `Int32.MinValue % -1`. La correction du reste utilise son signe et celui du diviseur, sans `abs` ni addition débordante.

Une expérience préalable de 10 169 couples contre le vrai `intMod` JS a trouvé zéro écart pour ce helper, ainsi que pour les additions/soustractions Int32 F#. Résumé conservé dans `modulo-probe.json` ; la matrice de bornes est désormais couverte par le test persistant `tests/int-kernel.mjs`.

Défaut préexistant observé séparément : le FFI `sharpurs/sharpurs-prelude/src/Data/EuclideanRing.fs` diverge du Prelude JS sur 1 272 couples de cette expérience. Par exemple, `1 mod 2147483647` y donne `-1`, au lieu de `1`. Ce FFI reste inchangé dans cette étape ; le nouveau noyau utilise le helper validé.

## Rejouer la validation

Depuis `/Users/0x1/Documents/htdocs/sharpurs/sharpurs` :

```sh
npm run build --silent
spago test
DOTNET=/Users/0x1/.dotnet/dotnet npm run test:kernel
DOTNET=/Users/0x1/.dotnet/dotnet npm run test:runtime
```

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-sharpurs`, avec le TAST du benchmark déjà compilé :

```sh
../sharpurs/sharpurs/bin/sharpurs --main App
(cd run/bak/sharp/output/Main && /Users/0x1/.dotnet/dotnet build Program.fsproj -c Release --nologo)
/Users/0x1/.dotnet/dotnet fsi --nologo --exec scratch/sharpurs-int-emission-20260908/check-tco.fsx /Users/0x1/Documents/htdocs/altbak.pub-sharpurs/run/bak/sharp/output/Main/bin/Release/net8.0/Program.dll
```

Aucun chronométrage du benchmark n’a été collecté : cette étape valide la génération et les résultats. La prochaine étape est la mesure de TCO avant/après, avec les temps individuels, résultats vérifiés, allocations et GC, en se référant au README officiel d’altbak.pub. `metadata.json` conserve les empreintes des sources et du programme vérifié.
