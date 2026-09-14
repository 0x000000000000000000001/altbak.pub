# Sticky sharing : preuves et mesures isolées

14 septembre 2026. **Aucune accélération convaincante : la saturation forcée ralentit la lecture possédée de 9,97 % et sa variante avec fermeture de 2,89 %.** Les mises à jour restent dans la variation mesurée. Aucun fichier du compilateur, du runtime source ou de la sortie Rust active n'a été modifié.

## Ce qui est testé

Le [runtime local réel](../../../../purust/purust/tests/runtime/perceus_ptr/src/local.rs) considère `u32::MAX` comme immortel : `clone` et `drop` cessent de modifier ce compteur, `is_unique` reste faux et `make_mut` copie. Le compteur suit les propriétaires simultanés, pas le cumul des clones exécutés. Le test de cycle de vie vérifie explicitement qu'un partage normal à deux propriétaires redevient unique après un `drop`.

Le [runtime threaded](../../../../purust/purust/tests/runtime/perceus_ptr/src/threaded.rs) délègue à `Arc`. Il expose la même constante et borne la valeur renvoyée par `count()`, mais ne contient aucun chemin d'immortalisation sticky. Les ADT natifs tels que RBTree et les enveloppes de fermetures `Func1::Shared` utilisent aussi les pointeurs standard `Rc`/`Arc`, pas cette branche locale de `PerceusPtr`.

[probe.py](probe.py) copie les sources exactes de `purust_core` et extrait **sans réécriture** `Test_Records_updateRec` de la sortie générée active. Les empreintes se trouvent dans [metadata.json](metadata.json). Les entrées sont les trois mêmes types de records fermés et quatre entiers non nuls ; le noyau reste celui des 10 000 mises à jour du benchmark. La création de chaque entrée est volontairement fraîche pour séparer le cas unique du cas où une ancienne racine reste conservée, comme avec la valeur initiale de module.

La seule modification de runtime concerne sa **copie scratch** : exposition du setter de compteur déjà réservé aux tests et instrumentation optionnelle. Deux exécutables sont compilés en O1 avec contrôles de débordement et allocateur Rust par défaut `System` : l'un compte les événements ; l'autre conserve les branches du runtime sans compteurs. Le runner complet utilise `mimalloc` ; ces profils ne sont pas interchangeables. Les deux exécutables vérifient les mêmes résultats. Les dépendances tierces inchangées proviennent des bibliothèques release déjà compilées ; `purust_core` et le runtime sont recompilés séparément pour chacun des deux exécutables. Les sources scratch et les deux binaires sont figés par SHA-256 dans `metadata.json` et revérifiés avant tout timing.

La variante forcée place les **trois cellules initiales** directement à **4 294 967 295** avant le travail. C'est artificiel et explicitement étiqueté `forced-sticky` ; le seuil n'est jamais abaissé. Les nouvelles cellules créées par `make_mut` commencent normalement à 1. Il n'y a pas de resaturation cachée à chaque itération.

## Compteurs observés, 10 000 itérations

| Charge | Pic de propriétaires normal | Clones normal | Écritures compteur normal / forcé | Copies `make_mut` normal / forcé | Allocations / libérations normal | Allocations / libérations forcé |
|---|---:|---:|---:|---:|---:|---:|
| Lecture possédée, chaîne de getters | 2 | 30 000 | 60 003 / 0 | 0 / 0 | 3 / 3 | 3 / 0 |
| Lecture empruntée, getters scalaires | 1 | 0 | 3 / 0 | 0 / 0 | 3 / 3 | 3 / 0 |
| Fermeture capturant le record, lecture possédée | 3 | 30 001 | 60 005 / 0 | 0 / 0 | 3 / 3 | 3 / 0 |
| Noyau Records, entrée unique | 2 | 20 000 | 40 003 / 39 999 | 0 / 3 | 3 / 3 | 6 / 3 |
| Noyau Records, ancienne racine conservée | 3 | 20 003 | 40 012 / 39 999 | 3 / 3 | 6 / 6 | 6 / 3 |

Ces compteurs portent sur les seules cellules `PerceusPtr`. L'allocation `Rc` de la fermeture n'y est pas incluse et sa politique de comptage ne change pas. Les nombres d'appels instrumentés ne préjugent pas des instructions que LLVM conserve dans un exécutable chronométré.

**Aucune des charges normales n'atteint naturellement le seuil sticky : leur pic vaut 1 à 3 propriétaires.** Les 30 000 clones d'une lecture ne signifient donc jamais un compteur à 30 000.

En lecture possédée, l'état forcé supprime les écritures logiques de compteur, mais les mesures ci-dessous ne montrent aucun gain machine. En lecture empruntée, le travail de comptage a déjà disparu de la boucle. Dans le noyau Records unique, la saturation initiale impose trois copies dès la première écriture ; les nouveaux records redeviennent normaux, donc presque toutes les écritures de compteur des 10 000 mises à jour restent présentes.

Chaque graphe forcé laisse **3 cellules, 168 octets de stockage direct**, non libérées. Chaque processus de comptage utilise un seul graphe ; chaque processus chronométré utilise un graphe de chauffe puis un graphe mesuré, soit **6 cellules et 336 octets directement retenus** après ces deux charges. Les compteurs couvrent ici l'en-tête et le payload direct des cellules ; un record réel retenant d'autres ressources pourrait retenir davantage. Les compteurs normaux libèrent toutes leurs cellules.

Les quatre champs du résultat sont vérifiés, pas seulement `f` ; la version initiale conservée reste inchangée. Un payload muni d'un destructeur confirme indépendamment : partage à deux puis retour à l'unicité normal, perte persistante d'unicité après saturation, copie correcte sur `make_mut`, ancien contenu immuable, destructeur de la version copiée exécuté et destructeur du payload immortel jamais exécuté. [Résultats bruts](counts.json).

## Mesures : sept paires, un million d'itérations

Le harness empêche l'élimination de la charge avec des entrées/sorties `black_box`, des lecteurs `inline(never)`, l'observation des quatre champs et un appel de fermeture réellement indirect. L'instrumentation et le chronométrage sont séparés. Le coordinateur a exécuté la commande distincte après les autres préparations :

```sh
python3 probe.py time --timing-authorized
```

Elle utilise sept paires de processus alternées normal/forcé et l'exécutable sans compteurs. Chaque processus commence par une chauffe explicite de **10 000 itérations du même scénario et du même mode**, sans lecture d'horloge, puis mesure **1 000 000 d'itérations**. Création/saturation de la nouvelle entrée mesurée et capture de fermeture précèdent le timer. Pour **tous** les scénarios, le segment mesuré inclut la boucle, l'observation et la vérification du résultat puis la destruction de tous les propriétaires finaux : résultat, fermeture et captures, racine et ancienne version retenue selon le cas. Les cellules immortelles restent retenues, y compris les trois cellules propres à la chauffe forcée. Chaque paire compare le même scénario et le même protocole.

| Charge | Médiane normale | Médiane forcée | Écart forcé | Paires favorables au forcé |
|---|---:|---:|---:|---:|
| Lecture possédée | 13,866 ms | 15,249 ms | **+9,97 %** | **0 / 7** |
| Lecture empruntée | 4,626 ms | 4,623 ms | −0,07 % | 3 / 7 |
| Lecture possédée via fermeture | 15,659 ms | 16,111 ms | **+2,89 %** | **0 / 7** |
| Noyau Records, entrée unique | 41,067 ms | 40,783 ms | −0,69 % | 4 / 7 |
| Noyau Records, ancienne racine conservée | 41,464 ms | 41,707 ms | +0,59 % | 4 / 7 |

La lecture possédée ralentit dans les sept paires ; ses plages sont disjointes : **13,759–14,564 ms** normales contre **14,893–15,644 ms** forcées. La fermeture ralentit également dans chaque paire, avec des plages globales qui se recouvrent. Les autres écarts sont petits, les plages se recouvrent et le sens varie entre les paires : aucun gain stable n'est établi pour la lecture empruntée ou les mises à jour. [Temps bruts](timings.json), [résumé avec plages et paires](../summary.json).

**Supprimer les écritures logiques de compteur ne suffit pas à démontrer une accélération machine.** Les compteurs instrumentés et les temps de l'exécutable sans compteurs apportent deux observations distinctes. Aucune cause précise du ralentissement n'a été profilée ; cette étude ne l'attribue ni à LLVM, ni aux branches, ni au cache sans preuve supplémentaire.

Le seuil naturel n'est jamais atteint dans les charges observées, qui conservent au plus trois propriétaires simultanés. Le chemin forcé retient de la mémoire et retire l'unicité, sans bénéfice temporel convaincant ici. **Cette expérience ne justifie donc aucune nouvelle intégration ou activation forcée de sticky sharing dans le compilateur.** Elle ne modifie pas la protection par saturation déjà présente dans le runtime local.

La baseline officielle reste le [README d'altbak.pub](../../../../altbak.pub/README.md#rust) : Records **382 µs**, RBTree **11 631 µs**, total **12,56 ms** ; dernière colonne native **4 µs / 36 070 µs / 36,13 ms**. Ces mesures scratch à un million d'itérations, avec saturation artificielle et allocateur `System`, ne remplacent pas ces références, ne prédisent pas un gain du runner complet et ne s'appliquent pas au `Rc<Tree>` de RBTree.

Reproduction sans timing : `python3 probe.py build`, puis `python3 probe.py check`.
