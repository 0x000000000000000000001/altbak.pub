**Audit des allocations de gopurs — 8 septembre 2026**

L’audit confirme des possibilités concrètes de réduction des allocations **avant de développer des arènes ou une analyse d’ownership**. Certaines optimisations que nous évoquions existent déjà. Les possibilités les mieux étayées concernent les changements de représentation : copies de tableaux, boxing des petits records et appels génériques.

Un témoin isolé ArrayOps montre **7 à 8 % de temps en moins et 19,4 % d’octets alloués en moins**, en retirant seulement deux conversions successives. Pour StateMonad, le profil identifie environ **80 % des octets alloués dans le boxing de records à deux champs**. Une représentation compacte déjà disponible dans le runtime permettrait théoriquement de réduire d’environ 30 % les octets alloués par ce test ; cette modification n’a pas été effectuée ni chronométrée.

RBTree est différent : son coût vient surtout des nouveaux nœuds nécessaires à l’algorithme persistant actuellement généré. Aucune infrastructure opérationnelle d’ownership permettant de les modifier en sécurité n’a été trouvée. Le nom des fichiers `FBIP` et `UsageAnalysis` pouvait laisser penser le contraire.

**Périmètre et absence de modifications du compilateur**

Lecture de la chaîne active, du générateur, du runtime, du backend optimizer local, des tests et du Go effectivement présent dans altbak. Mesures des 14 tests de `App`, six profils d’allocations, sondes de runtime et un témoin ArrayOps. L’audit porte sur allocations, représentations et réutilisation ; il ne constitue pas une preuve formelle de correction de tous les programmes compilables.

| Élément | Version auditée |
|---|---|
| gopurs | `e1fa15c98093b0640e1af190cf16d6c0514a3086` |
| Backend optimizer local | `67ba2151e3717b27a13e95c25d3b25c8bdc645cc` |
| altbak.pub | `f447b5c73c1170e8ee586a12ed6ae27f9f621905` |
| Exécution | Go 1.27.0, Darwin arm64, Apple M4 Pro |
| Bundle gopurs SHA-256 | `2ba567e31f9980b215172e3af2ab2267547256ceb6214eaadf80a277048992af` |

Les empreintes avant/après couvrent **905 fichiers gopurs et 1 385 fichiers du backend optimizer : aucune différence**, HEAD et statuts Git inchangés et propres. Le bundle et certains modules JS ignorés par Git sont également inclus. Aucun build du compilateur, aucune régénération de l’output altbak, aucune modification des tests du dépôt. Les seuls fichiers créés sont ce dossier d’audit et des caches Go sous `/private/tmp`.

Contrôles de provenance complémentaires : les SHA-512 des 517 sources enregistrées dans le cache de compilation gopurs correspondent aux sources actuelles ; `node --check` du bundle réussit ; les passes examinées se retrouvent dans le bundle. Le runtime altbak correspond au contenu Go de `Runtime.purs`. Cela donne une cohérence forte, sans constituer une reconstruction indépendante du bundle. L’output altbak existant a été mesuré tel quel ; ses fichiers Go sont identifiés dans le [manifeste d’entrée](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/generated-input.sha256.json).

Preuves : [état initial](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/repository-before.json), [contrôle final](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/repository-after.json).

**1. Ce qui fonctionne déjà**

La chaîne exécutée est : lecture du CoreFn enrichi → collecte des instanciations et monomorphisation → optimisations PBO → génération Go avec fusion de thunks et analyse TCO → impression du Go.

Le lecteur ouvre `output/*/corefn.json`, malgré la mention de `tcorefn.json` dans la documentation. PBO provient du checkout local déclaré dans spago. Références : [Main](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Main.purs:149), [appel de buildModules](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Main.purs:443), [lecteur CoreFn](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/App.purs:64).

| Optimisation | État constaté |
|---|---|
| Scalaires Go natifs, records fermés, tableaux typés, ADT à pointeurs, enums | Actifs |
| Appels directs `Call_…`, aplatissement d’arguments et récursion terminale | Actifs |
| Monomorphisation et suppression de dictionnaires | Actives, selon les informations de types disponibles |
| Élimination de constructeurs, records et tableaux temporaires | Active dans PBO, sous conditions d’usage et de capture |
| Intrinsics map/filter/fold transformés en boucles | Actifs, avec conversions et dispatch générique encore présents |
| Réutilisation d’un constructeur déjà identique | Active, sans mutation |
| Fusion de certains producteurs de thunks | Active, volontairement étroite |
| Boxing spécialisé des retours FFI entiers | Déjà intégré |
| Getters top-level avec `sync.Once` | Actifs |

L’analyse [Usage de PBO](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Analysis.purs:36) compte appels, accès, captures et mises à jour. Elle sert notamment à [éliminer des constructions temporaires](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Semantics.purs:1723). Elle n’établit pas l’absence d’alias dans le tas. Refaire simplement une analyse des occurrences ferait donc en partie doublon.

`Value` représente trois mots sur cette cible. Int/Float/Bool ne nécessitent pas un objet scalaire alloué. `Str` emballe un header de string, sans recopier automatiquement les caractères ; `Array` emballe un header de slice, sans recopier automatiquement le backing array. Il ne faut pas attribuer une allocation de contenu à chaque appel de ces helpers. [Runtime](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Runtime.purs:61).

**2. Mesures du code altbak original**

Chaque ligne correspond à une exécution complète de son `act`, incluant l’effet opaque et la conversion du résultat en string. Getters initialisés et trois exécutions de chauffe ; trois échantillons avec répétitions calibrées ; GC forcé avant chaque échantillon, hors chronométrage ; `GOGC=800`. Les résultats retournés restent identiques entre répétitions.

Les temps sont des médianes d’échantillons, indicatives sur cette machine. Ce protocole vise surtout les allocations et diffère du minimum de dix temps utilisé par le site. Les très petits tests comprennent une part significative de coût d’enveloppe. **Les octets ci-dessous sont cumulés par exécution, pas le pic de RAM ni le tas conservé.**

| Test | Temps médian (µs) | Octets alloués/op | Allocations/op (arrondies) | Résultat |
|---|---:|---:|---:|---|
| AstTree | 0,70 | 1 808 | 49 | 7 |
| Fib | 0,44 | 114 | 7 | 55 |
| ListOps | 9,35 | 32 594 | 1 362 | 202950 |
| TCO | 40,48 | 129 | 8 | 100000 |
| Records | 5,34 | 128 | 8 | 20000 |
| Ackermann | 16,85 | 116 | 7 | 125 |
| Church | 475,69 | 6 941 | 165 | 100000 |
| Primes | 79,23 | 257 493 | 10 732 | 21536 |
| RBTree | 23 045,21 | 79 488 696 | 2 483 959 | 22 |
| Polymorphism | 2 326,80 | 222 | 10 | 10000000 |
| StateMonad | 138,84 | 578 730 | 13 268 | 1200 |
| LazyEvaluation | 248,67 | 135 | 8 | 1000000 |
| ArrayOps | 10,02 | 77 190 | 23 | 202950 |
| RowToList | 0,06 | 112 | 6 | 5 |

Preuves : [mesures brutes](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/measurements.jsonl), [synthèse JSON](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/measurements-summary.json), [harnais](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/main.go).

Plusieurs tests effectuent déjà leur calcul principal avec très peu d’allocations : Records, TCO, Polymorphism, LazyEvaluation, Fib et Ackermann. Des arènes auraient peu de matière à optimiser sur ces chemins. Church alloue peu relativement à son temps ; cela ne permet pas d’attribuer sa durée principalement à l’allocation.

**3. ArrayOps : une possibilité mesurée, sans arène**

Dans [sumEvens généré](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Test_ArrayOps.go:137), le filtre produit un `[]Value`. Le générateur le convertit en `[]int64`, puis immédiatement en `[]Value` pour le fold. Le profil attribue **4 096 + 10 880 = 14 976 octets** à ces deux buffers pour 900 éléments d’entrée.

Le témoin d’audit copie cette fonction hors du dépôt du compilateur et enlève uniquement cet aller-retour. Il conserve range, filtre, croissance par append, appels `Apply`/`Apply2` et somme. Une copie inchangée sert de contrôle des effets de déplacement entre packages.

| Version du noyau, n=900 | Temps médian, 7 passages | Octets/op | Allocations/op |
|---|---:|---:|---:|
| Fonction générée originale | 9,044 µs | 77 056 | 15 |
| Copie inchangée | 9,129 µs | 77 056 | 15 |
| Copie sans aller-retour | 8,398 µs | 62 080 | 13 |

Le temps baisse de **7,1 % par rapport à l’original**, ou 8,0 % par rapport à la copie de contrôle. Les octets alloués baissent de **19,4 %**. Ces mesures portent sur le noyau, d’où une enveloppe plus petite que le `act` de la table générale. Les variantes ont été comparées à l’original sur tous les entiers de −32 à 2 048, soit **2 081 entrées**. Les mesures sont courtes et les variantes passent dans un ordre fixe ; les octets et nombres d’allocations constituent l’évidence la plus robuste. Aucun pic RSS n’a été mesuré pour ce témoin.

Il existe donc bien un gain vitesse/allocation possible sur un cas réel, indépendamment de Solod et de toute politique d’arène. Cela ne constitue pas encore une transformation générale prouvée du compilateur : elle devra reconnaître les types et préserver effets et aliasing aux frontières FFI/ST.

Sources de génération : [boxing des tableaux](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:186), [unboxing](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:364), [intrinsics](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1818). Preuves : [témoin](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/array_probe.go), [diff limité aux conversions](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/array-counterfactual.diff), [validation](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/array-validation.txt), [benchmarks bruts](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/array-benchmark.txt).

Deux constats connexes :

- L’output comporte 375 copies identité `Value → []Value`, réparties sur 20 fichiers. Ce comptage est statique, pas une fréquence d’exécution. Le wrapper [zipWith](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Data_Array.go:104) copie notamment ses deux entrées. Un emprunt du slice ne serait valide que si le destinataire ne peut pas le modifier de façon observable.
- Les branches d’appel direct de map/filter/fold consultent `moduleArities` avec une clé `Module.name`, alors que la table est construite avec des noms simples sanitizés. L’incohérence statique est précise : [construction](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:679), [lookup map](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1774), [fold](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1831), [filter](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1880). Les boucles générées examinées utilisent des appels génériques ; l’effet de cette incohérence sur un callback nommé survivant à l’inlining n’a pas été isolé. Un témoin minimal recompilé reste nécessaire avant un correctif ; changer seulement la clé ne suffit pas pour les fonctions locales ou externes.

**4. StateMonad : boxing coûteux réellement présent dans la boucle**

Le [boxing des records natifs](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:178) utilise toujours `RecordDict` avec deux slices. Pourtant le runtime et le printer disposent déjà de helpers compacts pour zéro à cinq champs, notamment [RecordDict2](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Runtime.purs:208).

Le profil StateMonad compte **3 620 records à deux champs**. Chaque boxing observé représente 48 octets de conteneur, 32 octets pour les clés et 48 octets pour les valeurs : **128 octets et trois allocations**. Les slices et conteneurs totalisent **463 360 octets**, soit environ **80 %** des 578 730 octets/op mesurés.

Les sites chauds sont [modify](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Test_StateMonad.go:111), le [cas de base](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Test_StateMonad.go:298), et les [deux reconversions successives](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Test_StateMonad.go:334). Ce dernier chemin reconstruit aussi des records entre plusieurs représentations successives, ce qui pourrait offrir une économie supplémentaire après analyse.

À nombre de records inchangé, `RecordDict2` représenterait 80 octets et une allocation. L’économie structurelle serait **173 760 octets/op et 7 240 allocations/op**, environ **30 % des octets actuels**. C’est une estimation à partir des layouts et du profil, **pas un résultat de remplacement exécuté**, ni une estimation de baisse du RSS ou du temps.

Le benchmark nommé Records travaille, lui, principalement avec des structs natifs : la seule présence de wrappers coûteux dans son fichier ne le rend pas prioritaire. Le profil permet de distinguer ces deux situations. [Profil StateMonad](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/StateMonad.alloc_space.focused.txt).

**5. RBTree : beaucoup d’allocations, mais aucune preuve actuelle de mutation sûre**

Pour 100 000 insertions, le profil compte exactement **2 483 948 nœuds alloués**, à 32 octets chacun : **79 486 336 octets** pour le noyau. Répartition :

| Site | Nœuds alloués |
|---|---:|
| balance | 2 383 932 |
| création des feuilles dans ins | 100 000 |
| recoloration de racine dans insert | 16 |

Cette répartition corrige une piste trop générale de notre échange : **optimiser encore makeBlack n’est pas le grand levier ici**. La règle de réutilisation d’une racine déjà noire est déjà présente dans l’output actuel, et les recolorations restantes représentent 16 nœuds sur près de 2,5 millions. [Profil objets](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/RBTree.alloc_objects.focused.txt), [insert](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Test_RBTree.go:1569).

`constructorReuse` retourne le nœud original quand une reconstruction avec un champ constant ne change effectivement rien. Sinon, il alloue. Il ne modifie aucun champ et ne teste pas `Rc == 1`. [Règle](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:411), [utilisation](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:2763).

Que `buildTree` remplace son accumulateur à chaque insertion ne prouve pas l’unicité transitive des nœuds. Un caller peut conserver l’arbre initial ; les versions partagent leurs sous-arbres ; `insert` doit conserver son comportement persistant hors du benchmark. Une spécialisation depuis un arbre vide pourrait être un cas favorable, mais demande une preuve adaptée des alias et des durées de vie.

Les modules `Gopurs.UsageAnalysis`, `Gopurs.FBIP` et `Gopurs.OptimizeTAST` ne sont pas raccordés à la chaîne active. Le champ `Rc` est émis et `IncRef` existe, mais sans instrumentation complète des références, sans décrément ni décision de réutilisation : **ce n’est pas un comptage de références fonctionnel**. Retirer `Rc` ne garantit même pas une baisse de taille, à cause de l’alignement.

Deux défauts latents ont été vérifiés par appels purs aux modules JS déjà compilés, en mémoire :

- [UsageAnalysis](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/UsageAnalysis.purs:118) ignore les sous-expressions de `PrimEffect`. `Local x` donne un usage et une variable libre, mais `RefRead x` donne zéro pour les deux. Un comptage syntaxique correct resterait insuffisant pour prouver l’ownership, notamment dans une closure réutilisée.
- [FBIP](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/FBIP.purs:39) peut extraire une projection hors de la lambda qui lie sa variable : `Abs(x, Accessor(x))` devient schématiquement `Let(p, Accessor(x), Abs(x,p))`. Il déplace alors un accès hors de sa portée. Cette passe est une extraction de projections, pas une gestion de mémoire opérationnelle.

Ces défauts sont **dormants** : ils ne démontrent pas une régression du Go actif. Ils interdisent en revanche de considérer qu’il suffirait de réactiver ces passes pour disposer d’ownership.

Le précédent POC arènes reste cohérent avec ces constats : à 100k insertions, l’arène progressive complète donnait 13,729 ms et 81,41 Mio de pic RSS contre 22,634 ms et 43,86 Mio pour le Go original ; le pool de 8 Mio donnait 22,923 ms et 42,02 Mio. Le compromis favorable vitesse/RAM observé à 10k ne se généralisait donc pas à 100k. Ces chiffres proviennent du [POC précédent](/Users/0x1/Documents/htdocs/altbak.pub/scratch/solod-poc-20260908/rbtree/memory-poc/README.md), sans nouvelle exécution dans cet audit.

**6. Runtime : rétention et concurrence confirmées**

Chaque création via `Func`…`Func11` écrit la fonction dans une globale `EscapeSink`, utilisée pour forcer son échappement. [Runtime.purs](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Runtime.purs:451).

Une sonde capture un tableau de 16 Mio, abandonne la valeur fonction, puis force le GC. Le tas passe d’environ 383 ko à 17,12 Mo. Après création d’une autre fonction sans cette capture et nouveau GC, il redescend à environ 347 ko. La dernière closure conserve donc effectivement son environnement jusqu’à remplacement de la globale. Il s’agit d’une rétention de la dernière capture, pas d’une croissance cumulative de toutes les closures. [Résultat de rétention](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/retention.txt).

Deux goroutines créant des closures reproduisent une **data race** sur `forceEscape`. `go test -race` échoue comme attendu avec une trace sur cette écriture. Cela concerne notamment un appel concurrent via FFI Go ; les benchmarks monothread ne démontrent pas une corruption liée à cette course. [Trace race](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/race.txt).

Ne pas supprimer mécaniquement ce mécanisme : la représentation utilise des pointeurs unsafe pour reconstituer les fonctions. Une évolution doit garantir la durée de vie des captures, y compris après retour et GC. La synchronisation seule traiterait la course mais pas la rétention ; l’audit ne choisit pas une nouvelle représentation.

Deux autres sondes précises :

- Un appel saturé à `Func5` via `Apply5` alloue zéro objet ; l’équivalent à six arguments via `Apply6` alloue **cinq closures intermédiaires**. Le chemin [Apply6…10](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Runtime.purs:811) utilise toujours des applications successives. Il existe 17 appels statiques `Apply6…10` dans l’output, sans preuve qu’ils dominent les tests actuels.
- `ValueToAny(Float(1.5))` panique : cet ancien helper lit `UnsafePtr` alors que Float encode ses bits dans `IntVal`. Aucun appel à ce helper n’a été trouvé dans le Go altbak généré ; ce défaut concerne ce chemin runtime, sans expliquer les performances actuelles. [Helper](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Runtime.purs:826).

[Sondes exécutables](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/runtime_test.go), [résultats](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/runtime-probes.txt). La sonde de compatibilité enregistre intentionnellement le panic observé ; son statut PASS signifie que le défaut a été reproduit, pas que le helper est correct.

**7. Autres possibilités documentées, sans gain promis**

| Constat | Conséquence et niveau de preuve |
|---|---|
| Reboxing d’ADT paramétrés | Des helpers reconstruisent récursivement les sous-arbres entre représentations spécialisées et polymorphes. Coût visible dans le code, poids applicatif non mesuré. |
| FFI records | Certaines conversions passent par `RecordToMap` puis une seconde map. Coût potentiel aux frontières fréquentes. |
| FFI callbacks scalaires | Certains paramètres passent encore par `Box`, alors que les retours sont spécialisés. À profiler sur une application qui utilise ces callbacks. |
| Primes et ListOps | Les profils attribuent surtout les allocations aux listes construites par range/filter. Une fusion producteur-consommateur serait une possibilité distincte des arènes, sans résultat de transformation dans cet audit. |
| Church | Les objets alloués concernent surtout applications partielles et construction des fonctions ; son temps ne justifie pas à lui seul une politique d’arène. |

Le reboxing récursif est visible dans [Data.Map.Internal](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Data_Map_Internal.go:6709), y compris des [aller-retour imbriqués](/Users/0x1/Documents/htdocs/altbak.pub/run/bak/go/output/purescript/Data_Map_Internal.go:2786). Les helpers ne mémoïsent pas les sous-structures partagées. Les 991 définitions `Rebox_…` recensées ne représentent pas 991 chemins chauds. Le helper réflexif `ReboxStruct` n’est pas appelé dans cet output. Éviter de confondre ce coût avec une réflexion omniprésente, ou de proposer un cast unsafe entre layouts différents.

Sources FFI : [tableaux](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:3396), [records](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:3419), [paramètres de callbacks](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:3110).

**8. Limites du pipeline et des preuves de performance**

Le backend optimizer contient des directives d’inlining nommant explicitement `Test.RBTree`, `Test.Polymorphism` et `Test.LazyEvaluation`. Elles sont actives et intégrées dans le bundle. Les scores altbak bénéficient donc déjà de réglages par nom ; leur extrapolation à une autre application demande prudence. [Directives](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Directives/Defaults.purs:39).

La monomorphisation [exporte toutes les déclarations finales](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Monomorphize.purs:728), tandis que [la DCE conserve les bindings exportés](/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Convert.purs:177). Cela limite l’élimination top-level amont et peut augmenter volume généré et temps de compilation. Aucune pénalité RAM à l’exécution n’en découle automatiquement : le linker Go peut éliminer du code inaccessible.

`shouldUncurryAbs` est un stub, mais d’autres mécanismes d’aplatissement des applications et abstractions sont actifs. De même, `CodeGenBackend` et l’ancien `Main.runBuild` ne sont pas la chaîne exécutée. Les lire seuls donnerait une image erronée des capacités actuelles.

Le cache gopurs est écrit mais son réemploi est désactivé par [res <- pure Nothing](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Main.purs:454). Le cache PBO est distinct. Les options `--output`, `--bundle`, `--autoload-path` sont parsées mais non exploitées par Main ; pour une future compilation isolée, il faut un cwd séparé avec son propre `output`, pas simplement un `--output` différent.

**9. Tests et garanties qui restent à établir**

L’inventaire contient 376 fixtures PureScript de premier niveau, 367 snapshots Go principaux et un snapshot FFI. Les neuf fixtures sans snapshot correspondent à la blacklist du runner. La mention de couverture de 100 % dans la documentation nécessite donc une précision sur le périmètre.

La suite compile, compare les snapshots et exécute les résultats ; elle ne compare pas systématiquement au backend JS. Aucun seuil de coût mémoire, test `AllocsPerRun`, benchmark d’allocations ou passage `-race` n’a été trouvé dans la suite auditée. Le test [RBTree](/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/passing/RBTree.purs:46) affiche sa profondeur, sans asserter les invariants de l’arbre. Les vérifications d’invariants du POC antérieur apportent une couverture supplémentaire, limitée aux variantes de ce POC.

Des tests ciblés sont déjà utiles : [ConstructorReuse](/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/passing/ConstructorReuse.purs:111) vérifie notamment anciennes versions conservées, sources mélangées, plusieurs changements et −0.0 ; ThunkFusion vérifie des refus, l’ordre des effets et des portées. Leur existence ne mesure pas l’économie obtenue.

Attention au snapshot RBTree du dépôt : il ne contient pas la réutilisation visible dans la source et le Go altbak actuels. Il ne doit pas servir seul à décrire ce qui a été mesuré. La suite complète du compilateur **n’a pas été relancée**, car son runner écrit dans le dépôt ; l’audit ne revendique pas sa réussite sur cette révision.

Les profils sont pris dans des processus séparés, avec `MemProfileRate=1`, après chauffe et soustraction d’un profil de référence. Le filtrage `gopurs/output/` écarte les allocations du mécanisme de profilage. Leurs pourcentages se rapportent au sous-graphe filtré ; les petits écarts avec les mesures répétées viennent notamment des GC explicites et du refroidissement des pools de formatage. Les temps profilés n’ont pas été utilisés comme benchmarks. Aucune mesure concurrente entre variantes.

**10. Priorités proposées pour une décision ultérieure**

| Priorité | Sujet | Pourquoi | Ce qui manque avant intégration |
|---|---|---|---|
| Fiabilité | Durée de vie des closures et `EscapeSink` | Rétention et course reproduites | Représentation correcte sous GC/concurrence, puis mesures |
| Performance 1 | Aller-retour de représentation ArrayOps | Gain temps/octets mesuré | Règle générale étroite, tests d’effets et frontières |
| Performance 1 | Boxing compact des petits records | Coût chaud mesuré dans StateMonad | Remplacement isolé, comportement FFI, chrono/RSS |
| Performance 2 | Résolution des appels dans intrinsics | Incohérence statique localisée | Témoin recompilé, traitement des fonctions locales/externes |
| Performance 2 | Appels saturés à six–dix arguments | Allocations intermédiaires reproduites | Mesurer un utilisateur réel de ces arités |
| Selon application | Reboxing récursif, FFI, copies identité | Coûts potentiellement importants | Profil de chemins réellement utilisés |
| Recherche distincte | Ownership et mutation des ADT | Potentiel RBTree, risque sémantique et ampleur élevés | Analyse d’alias, preuves de durée de vie, rentabilité |
| Pas de priorité générale établie | Arènes | Compromis vitesse/RAM dépendant de la charge | Cas d’usage à durée de vie groupée démontrée |

Mon avis est donc plus précis que la recommandation initiale : **oui à des réductions ciblées des conversions et du boxing ; pas de justification actuelle pour lancer un système général d’arènes ou d’ownership dans gopurs**. Le témoin ArrayOps prouve un gain local sans compromis d’octets alloués. Il ne promet ni une accélération globale de tous les programmes, ni un pourcentage identique de RAM économisée.

**Reproduction sans modifier gopurs**

Se placer dans ce dossier d’audit. Le module Go importe l’output existant avec une directive `replace`, sans le régénérer. Les commandes exactes sont dans [REPRODUCTION.md](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/REPRODUCTION.md). Les artefacts de mesure ne dépendent pas de Solod.
