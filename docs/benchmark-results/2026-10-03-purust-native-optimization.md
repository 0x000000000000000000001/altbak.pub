# Purust — optimisation nocturne du compilateur natif

Campagne du **3 octobre 2026**, sur les entrées Aff figées le 2 octobre.
Le stage 2 qualifié réduit le temps de compilation native de **52,5 % sur
238 modules** et de **53,0 % sur 244 modules**. Le gain principal est confirmé
sur **15 paires, toutes favorables**, avec les fichiers générés identiques.
Le stage 2 mesuré est **installé**, avec son empreinte vérifiée. Dans la campagne
commune, **Purust natif générant du Rust** est à **2 569 ms**, contre **2 109 ms
pour gopurs natif générant du Go** : **1,218×** le temps gopurs.

**Périmètre corrigé :** ce rapport mesure l'optimisation de Purust et une
comparaison entre deux générateurs. La comparaison de **gopurs exécuté en
JavaScript, Go et Rust**, produisant du Go dans les trois cas, fait l'objet de
la [qualification des hôtes gopurs](2026-10-03-gopurs-rust-host.md).

## Compilation native avant/après

| Corpus figé | Avant | Stage 2 final | Temps en moins | Protocole |
|---|---:|---:|---:|---|
| gopurs-aff, 238 modules | 5 447 ms | **2 586 ms** | **52,5 %** | 15 paires |
| purust-aff, 244 modules | 5 604 ms | **2 632 ms** | **53,0 %** | 7 tours |

Sur le corpus principal, le compilateur est **2,11× plus rapide**. La moyenne
passe de **5 310,5 à 2 560,2 ms (−51,8 %)** ; les plages sont respectivement
4 866–5 600 ms et 2 445–2 634 ms. L'écart médian apparié est **−2 838 ms**,
l'écart moyen **−2 750,3 ms**. Aucun échantillon mesuré n'est exclu.

### Phases, 238 modules

| Phase | Avant | Stage 2 final |
|---|---:|---:|
| Chargement TAST + tri | 657 ms | **227 ms** |
| Préparation | 618 ms | **62 ms** |
| Optimisation + génération | 3 346 ms | **1 918 ms** |
| Finalisation + émission | 721 ms | **267 ms** |
| **Backend complet, fin des workers comprise** | **5 447 ms** | **2 586 ms** |

Les médianes de phases ne s'additionnent pas ; `backend total` couvre aussi
l'attente finale des workers. Les **32 sorties × 484 fichiers** sont exactes.

| Paire | Avant, ms | Après, ms | Après − avant, ms |
|---|---:|---:|---:|
| 1 | 4 866 | 2 487 | −2 379 |
| 2 | 5 447 | 2 445 | −3 002 |
| 3 | 5 271 | 2 451 | −2 820 |
| 4 | 5 600 | 2 451 | −3 149 |
| 5 | 4 999 | 2 603 | −2 396 |
| 6 | 4 983 | 2 586 | −2 397 |
| 7 | 5 537 | 2 574 | −2 963 |
| 8 | 5 537 | 2 588 | −2 949 |
| 9 | 4 970 | 2 584 | −2 386 |
| 10 | 5 492 | 2 592 | −2 900 |
| 11 | 5 000 | 2 604 | −2 396 |
| 12 | 5 540 | 2 582 | −2 958 |
| 13 | 5 459 | 2 621 | −2 838 |
| 14 | 5 433 | 2 601 | −2 832 |
| 15 | 5 524 | 2 634 | −2 890 |

### Confirmation sur les 244 modules distincts

| Phase | Rust avant | Rust final | JavaScript actuel |
|---|---:|---:|---:|
| Chargement TAST + tri | 667 ms | 233 ms | 581 ms |
| Préparation | 608 ms | 59 ms | 38 ms |
| Optimisation + génération | 3 444 ms | 1 941 ms | 4 869 ms |
| Finalisation + émission | 753 ms | 273 ms | 846 ms |
| **Backend complet** | **5 604 ms** | **2 632 ms** | **6 389 ms** |

Le natif final prend **58,8 % de temps en moins que JS** sur ce corpus. Plages
avant/après : 5 057–5 649 ms / 2 611–2 682 ms ; moyennes 5 394,6 / 2 631,4 ms.
Les **24 sorties × 496 fichiers** sont identiques à leur propre référence.
Les temps de ce corpus ne sont pas interchangeables avec ceux des 238 modules.

## Comparaison commune Go/Rust

Même TAST original de 238 modules, **dix tours**, cinq variantes, ordre tournant,
une chauffe par variante. Les exécutables Go sont ceux de la référence figée.

| Compilateur exécuté | Cible | Médiane | Moyenne | Min–max |
|---|---|---:|---:|---:|
| gopurs.js | Go | 7 750 ms | 7 751,6 ms | 7 632–7 887 ms |
| gopurs natif | Go | **2 109 ms** | 2 123,1 ms | 2 101–2 200 ms |
| purust.js actuel | Rust | 6 208 ms | 6 209,0 ms | 6 131–6 270 ms |
| Purust natif avant | Rust | 4 954 ms | 5 113,9 ms | 4 915–5 541 ms |
| Purust natif final | Rust | **2 569 ms** | 2 571,5 ms | 2 545–2 620 ms |

Le rapport Rust/Go passe de **2,349× à 1,218× dans cette même campagne**.
L'écart absolu passe de **2 845 à 460 ms**, soit **83,8 % de l'écart résorbé**.
Rust final prend encore **21,8 % de plus que Go** et **58,6 % de moins que JS**.
Le gain Rust avant/après de cette campagne est **48,1 %** ; le témoin alterne
toujours entre des régimes rapides et lents. La campagne primaire de 15 paires
et ses 52,5 % restent publiés intégralement, sans remplacer une série par l'autre.

Sur les 16 exécutions natives de chaque variante de la confirmation primaire
(chauffe comprise), le budget reste **4 PBO / 4 codegen / total 8**. Les tentatives
rejetées passent de **64–75 par compilation à zéro**. Le gain réduit le travail
réalisé et la variabilité de l'ordonnancement.

Les **22 sorties Go × 294 fichiers** et les **33 sorties Rust × 484 fichiers**
de la comparaison commune correspondent exactement aux références de leurs
cibles. Cette campagne mesure `gopurs-aff` ; les autres cellules Rust des
paquets `gopurs-*` restent non mesurées.

## Changements retenus

Le profil reste **O3 sans LTO, Arc, mimalloc, quatre workers PBO et quatre
workers de génération**. Le natif reste le défaut ; `PURUST_JS=1` choisit JS.

1. **Ordonnancement PBO** : conserver les modules déjà prêts, écarter les
   modules en attente et limiter la spéculation aux références disponibles.
   La visibilité par rang et la publication canonique sont conservées.
2. **Identité des clés de mémoïsation** : utiliser l'allocation partagée de
   `ExprType`/`BackendSyntax`, plutôt que l'enveloppe `Any` recréée à chaque appel.
   Le cache garde les propriétaires en vie et appelle les callbacks hors verrou.
3. **Rendu des types** : parcourir les arbres et les maps par emprunt, puis
   construire les chaînes dans des buffers natifs.
4. **Directives** : reconnaître le sous-ensemble ASCII courant sans reconstruire
   le lexer CST ; déléguer les autres syntaxes et les erreurs au parseur PS.
5. **Génération** : spécialiser les coercitions simples, les noms de champs et
   la canonicalisation des noms de records. Les mots-clefs ne reconstruisent
   plus une collection à chaque champ émis.
6. **Chargement TAST** : décoder les corps des modules et valider les faits
   d'usage nativement. Les portées lexicales, le premier diagnostic, les
   références de types partagées et les représentations UTF-16 restent exacts.

Les implémentations PureScript servent d'oracles et de replis. Les sorties
applicatives générées sont comparées octet pour octet aux références figées.

### Sélection et essais écartés

Les petites campagnes ci-dessous expliquent la sélection ; leurs pourcentages
marginaux ne s'additionnent pas et leurs témoins ont été remesurés à chaque lot.

| Ajout | Médiane du prédécesseur | Médiane avec ajout | Mesures |
|---|---:|---:|---|
| Ordonnanceur + identité du cache | 5 385 ms | 4 310 ms | 5 tours, référence commune |
| Rendu des types | 4 069 ms | 3 827 ms | 5 tours |
| Directives ASCII | 4 144 ms | 3 629 ms | 5 tours |
| `boxUnbox` | 3 336 ms | 3 293 ms | Confirmation 15 tours, 14 paires favorables |
| Noms de champs | 3 434 ms | 3 242 ms | 5 tours |
| Noms de records | 3 180 ms | 2 918 ms | 5 tours |
| Validation des faits d'usage | 2 866 ms | 2 613 ms | 5 tours |
| Décodage des modules | 2 507 ms | 2 404 ms | 5 tours |

Le port natif de `TypeSubstitution` est **écarté** : 15 tours donnent
2 608 → 2 590 ms, médiane appariée −16 ms, 11 paires favorables. Le gain de
0,7 % ne justifie pas cette complexité supplémentaire. Son code, ses tests,
ses exécutables et ses mesures sont conservés dans l'archive expérimentale.

Balayage des workers, trois mesures par configuration : **4/4 : 2 637 ms** ;
3/5 : 2 668 ms ; 2/6 : 2 972 ms ; 6/2 : 2 965 ms ; contrôle séquentiel :
5 955 ms. Le défaut 4/4 est retenu. `PURUST_PBO_JOBS` est le **budget total**,
dont la génération concurrente est retranchée : 4/4 exige budget 8, codegen 4.
Le premier balayage mal configuré est conservé et remplacé par
`workers-exploration-v2.json`.

## JSON → Typed AST Go/Rust

Les mêmes entrées et le même pilote structurel sont utilisés avant/après.
Chaque variante lance **trois processus**, avec ordre tournant des variantes
et des phases ; chaque processus exécute **deux chauffes et cinq mesures**.
Les empreintes JSON et AST de chaque processus concordent avec l'oracle JS figé.
Lecture des fichiers et calcul des empreintes sont hors chronométrage.

Statistique historique : médiane des trois minima de processus, en millisecondes.
Les quinze échantillons et leur médiane sont également conservés dans le JSON.

| Corpus | Phase | Go natif | Rust avant | Rust final |
|---|---|---:|---:|---:|
| 12 modules | Parsing JSON | 19,126 | 17,962 | 18,328 |
| 12 modules | Décodage du JSON pré-parsé | 11,606 | 93,914 | **13,823** |
| 12 modules | Chemin texte complet | **20,003** | 116,049 | **35,351** |
| 238 modules | Parsing JSON | 92,257 | 88,225 | 88,344 |
| 238 modules | Décodage du JSON pré-parsé | 77,714 | 449,139 | **71,161** |
| 238 modules | Chemin texte complet | **108,551** | 555,096 | **175,839** |

Rust réduit le temps de décodage de **85,3 % / 84,2 %**, et le chemin complet
de **69,5 % / 68,3 %**, respectivement sur 12/238 modules. Sur 238 modules,
le décodage Rust prend 0,916× le temps Go ; le chemin complet prend 1,620×.
Go utilise un chemin texte direct, tandis que Rust construit encore un arbre
JSON intermédiaire. Les phases indépendantes ne s'additionnent pas.

Le diagnostic JSON Go utilise **GOMAXPROCS=1, GOGC=100** ; son GC suit sa
politique normale. Le profil Rust est O3 sans LTO, Arc et mimalloc. La destruction
finale Rust est mesurée séparément : médianes des quinze passages combined,
**3,646 → 3,098 ms** sur 12 modules et **19,694 → 16,380 ms** sur 238 modules.
La médiane de tous les échantillons combined confirme le gain Rust :
**116,763 → 35,738 ms** et **559,518 → 176,695 ms**.

Les processus Go sont plus variables : minima combined 238 de
147,175 / 97,548 / 108,551 ms, contre 176,488 / 174,903 / 175,839 ms pour Rust.
Tous sont conservés ; aucun processus défavorable n'est retiré.

### Allocations JSON, 238 modules

Processus instrumentés séparés, empreintes identiques. Médianes des cinq
mesures par phase, après deux chauffes ; octets alloués cumulés en Go décimaux.

| Phase | Requêtes avant | Requêtes après | Go cumulés avant | Go cumulés après |
|---|---:|---:|---:|---:|
| Parsing JSON | 5 490 878 | 5 490 878 | 0,412 | 0,412 |
| Décodage | 35 113 805 | **4 096 136** | 1,218 | **0,188** |
| Chemin complet | 41 082 523 | **10 064 854** | 1,714 | **0,684** |

Le décodage supprime **88,3 % des requêtes d'allocation** et **84,6 % des octets
cumulés** ; le chemin complet supprime respectivement **75,5 %** et **60,1 %**.
Ces compteurs ne représentent pas la mémoire résidente. Les durées du binaire
instrumenté ne servent pas à calculer les gains de performance.

## CPU, mémoire et allocations du compilateur

Diagnostic séparé du compilateur non instrumenté, **trois paires après une
chauffe par variante**, `/usr/bin/time -l`. Médianes du processus entier,
démarrage et destruction compris, sur les mêmes 238 modules et sorties exactes :

| Ressource | Avant | Après | Réduction |
|---|---:|---:|---:|
| Temps réel du processus | 5,52 s | **2,59 s** | 53,1 % |
| CPU utilisateur + système | 13,55 s | **7,58 s** | **44,1 %** |
| Pic RSS | 472,47 Mio | **425,41 Mio** | **10,0 %** |

Le diagnostic d'allocations final utilise une copie instrumentée des sources
du stage 2, ensuite restaurées et revérifiées avec son exécutable. Ses sorties
restent exactes. Compteurs finaux :

| Phase | Requêtes d'allocation | Octets cumulés, Go |
|---|---:|---:|
| Chargement | 10 390 863 | 0,774 |
| Préparation | 1 156 068 | 0,072 |
| Optimisation + génération | 313 383 762 | 13,842 |
| Finalisation | 7 395 110 | 0,799 |
| **Backend entier** | **339 694 649** | **15,805** |

Les intervalles des phases ne couvrent pas tout le total. **238 tentatives
acceptées, zéro rejet**, et **167 133 succès / 181 181 sondes de mémo (92,2 %)**.
Le total de conversion seul est de 135 308 088 requêtes / 4,905 Go.
Le premier diagnostic après ordonnanceur + cache mesurait 472 405 467 requêtes /
22,763 Go : les spécialisations suivantes retirent encore **28,1 % des requêtes**
et **30,6 % des octets**. Ce point intermédiaire n'est pas la référence native
initiale. Les compteurs historiques de 244 modules restent séparés.

## Protocole du compilateur complet

- Corpus principal **gopurs-aff : 238 modules / 136 604 types / 26 606 491 octets**,
  manifeste SHA-256
  `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
- Corpus secondaire distinct **purust-aff : 244 modules / 138 448 types /
  27 122 224 octets**, manifeste TAST
  `b05785ab746528af700ab808066550cbd13df059aa2c7773ff5560d10ede0786`.
- Protocole pré-déclaré : **15 paires** Rust avant/après sur le premier corpus,
  **sept tours** sur le second, **dix tours communs Go/Rust**. Une chauffe par
  variante, processus et sorties neufs, premier compilateur tournant.
- Mesure : `backend total`, de chargement/tri à la fin de l'émission et des
  workers. Frontend, builds Cargo/Go, bootstrap, exécution applicative et
  démarrage/sortie du processus sont exclus.
- Les caches applicatifs `.purmeta` et `.cache` sont vidés à chaque passage ;
  le cache de fichiers du système n'est pas purgé. Builds et mesures sont
  sérialisés. Profils instrumentés et mesures de performance restent distincts.
- Comparaison des configurations de production : Go garde son launcher figé,
  `GOGC=off`, limite mémoire 10 Gio et ses workers 8/8/8. Rust garde 4 PBO +
  4 génération. Les passes et budgets mémoire des deux backends diffèrent.
- Machine : **Apple M4 Pro, 14 cœurs, 48 Gio, arm64, macOS 26.1** ; Node 24.8.0,
  Go 1.27.0, Rust/Cargo 1.96.0 ; aucune affinité CPU imposée.

## Qualification et intégrité

La chaîne **JS → stage 1 natif → stage 2 natif → projet frais** passe :
**453 modules / 282 637 types**, **914 fichiers du compilateur identiques**
entre JS et natif, puis **152 modules frais / 312 fichiers identiques** et
résultat applicatif du smoke vérifié. Les régressions PBO/codegen/TAST, les
tests différentiels des FFI et la suite Aff complètent cette qualification.

Les **12 suites natives** et les **12 suites PBO** passent. Couverture native :

- **409 812 comparaisons de rendu** sur les 136 604 types du corpus, plus
  948 cas synthétiques et sept sondes de délégation ;
- **61 600 comparaisons de noms de champs**, 1 032 cas de noms de records,
  2 028 coercitions scalaires/gardes et 60 cas fonctions/ADT ;
- **169 directives par défaut**, 352 cas générés et 33 replis exacts ;
- **238 modules décodés nativement** et validés, 51 cas frontières de modules,
  **110 909 annotations**, 824 tables / 149 495 entrées différentielles ;
- possession et libération des clés de cache, appels réentrants, Maps
  persistantes et **12 996 recherches de layout**, dont huit lecteurs concurrents.

La première exécution codegen a passé 88 checks et trouvé deux fixtures
standalone qui injectaient désormais trop de dépendances FFI. Leurs périmètres
ont été corrigés : sanitizer isolé, identités de types opaques dans la fixture
des services hôtes ; le test natif de mémo utilise les véritables types générés.
Quatre fixtures codegen b8x et une fixture TAST crypto exigent le conteneur
`core-api-cli-1`, indisponible sur cette machine. Elles sont consignées comme
non exécutées dans la qualification locale, avec les premiers échecs Docker
conservés ; les autres régressions sont exécutées sur l'hôte.
La passe codegen portable corrigée réussit **90/90 checks**, dont
**133 148 cas différentiels de sanitizer** et **63 assertions FFI JS/Rust**.
Les **42/42 tests TAST portables** passent également avec le frontend figé.
La suite Aff complète passe avec le stage 2 : **47 checks, cinq tests Rust,
neuf scénarios d'erreur**, concurrence Ref/AVar et durée de vie des enfants.

La vérification de publication a relu **120 sorties / 54 188 fichiers** des
campagnes primaires, secondaires, communes et des diagnostics. Elle contrôle
les empreintes des sources/binaires/tests, les temps bruts, leurs médianes et
les empreintes JSON/AST. Le binaire installé a exactement l'empreinte du stage 2
mesuré ; les lancements par défaut natif et `PURUST_JS=1` ont été vérifiés.

Le corpus chronométré gopurs-aff reste l'original. Ses anciennes assertions
sensibles à l'ordonnancement et la validation applicative synchronisée restent
documentées dans le [rapport de référence](2026-10-02-gopurs-aff-go-rust-compilation.md#validation-commune-avec-synchronisations-explicites).
L'identité exacte du code généré relie les mesures à ces validations ; la suite
purust-aff complète est exécutée séparément avec le nouveau compilateur.

Les diagnostics d'échec sont conservés : disque plein historique, erreurs de
harnais corrigées, clés dupliquées dans un test de permutation et échec DNS de
Cargo au lancement du build stage 2. Cette dernière étape est reprise hors
ligne après revérification du gel et des sources JS/natives identiques.
La première chauffe Go/JS commune est conservée dans `common-go-rust` : le
harnais avait copié `tools/` sous `bin/`. La campagne complète, après copie
fidèle et empreintée de l'installation Go, est `common-go-rust-final`.
Le premier diagnostic CPU/RSS a échoué parce que son parent `generated/`
manquait ; il est conservé dans `resources-final`. La reprise complète après
correction est `resources-final-v2`.

## Provenance et reproduction

- Archive : `var/benchmark/purust-pbo-20261002/`.
- Référence native qualifiée :
  `002ed7b50aaa2eb58c844fada98c715ebdaa1378ee8f62d1a23473d5f5e25639`.
- Nouveau stage 2 :
  `d170ad6a1b2b18627a70dfd873d35105b947835e03f32da998f218206cbcd703`.
- Bundle JS reconstruit :
  `df995c72762c1bdcb2cbb1c3c49d2cb59bfc8e7192724748c771a1dea956c181`.
- Binaire Go figé :
  `bae4835bdb63ee802f172253f6a5e88ef03bc141bb15330318ba4d1a54674c16`.
- [Harness et commandes](../../bin/benchmark/purust-pbo/README.md).
- [Données, échantillons bruts et vérifications](2026-10-03-purust-native-optimization.json).
- `confirmation-protocol.json` fixe les protocoles ; `candidates/` conserve
  chaque sélection immuable ; `rejected-substitution/` conserve le port écarté.
- Les caches Cargo supprimés font l'objet de rapports de nettoyage avec copie
  et vérification des exécutables. Sources, sorties et logs restent archivés.
