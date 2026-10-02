# Rust — JSON vers TAST et chargement du compilateur

Lot du **2 octobre 2026**. Le candidat « tableaux et annotations natifs » réduit
le chemin JSON → TAST complet de **18,3 % sur 238 modules** et de **19,6 % sur
12 modules**. Le gain de compilation complète est confirmé à **2,9 % sur
238 modules** et **2,5 % sur 244 modules**. Le stage 2 qualifié est installé.

## Résultat retenu

| Compilation backend native | Avant | Après | Réduction | Protocole |
|---|---:|---:|---:|---|
| gopurs-aff, 238 modules | 5 342 ms | **5 189 ms** | **2,9 %** | Médiane de 21 paires |
| purust-aff, 244 modules | 5 540 ms | **5 399 ms** | **2,5 %** | Médiane de cinq processus |

Sur les 238 modules, la moyenne passe de **5 210,5 à 5 072,1 ms (−2,7 %)** ;
le candidat est plus rapide dans **16 paires sur 21**. La différence médiane
appariée est de **−125 ms**. Le chargement médian passe de **782 à 635 ms**,
contre **3 158 à 3 135 ms** pour optimisation/génération. Les **44 sorties ×
484 fichiers** restent identiques à la référence historique.

Cette confirmation étendue répond à une première campagne défavorable, conservée
ci-dessous. Les deux compilateurs présentent des régimes rapides/lents avec un
travail spéculatif PBO variable : le gain global est bien plus petit et variable que
le gain isolé du décodeur. Tous les échantillons, diagnostics et qualifications
sont publiés dans les [données JSON](2026-10-02-rust-json-typed-ast.json).

## Première mesure du compilateur complet

Sur les 238 modules, médianes de cinq processus après une chauffe, ordre tournant :

| Phase | Natif avant | Candidat | JavaScript |
|---|---:|---:|---:|
| Chargement TAST + tri | 788 ms | **640 ms** | 567 ms |
| Préparation | 603 ms | 612 ms | 37 ms |
| Optimisation + génération | 2 668 ms | 3 195 ms | 4 558 ms |
| Finalisation + émission | 682 ms | 695 ms | 765 ms |
| Backend complet | **4 865 ms** | **5 266 ms** | 5 984 ms |

Le chargement gagne **18,8 %**, mais le total régresse de **8,2 %**. Le candidat
déclenche 70–73 tentatives PBO différées contre 65–72 pour le témoin dans cette
campagne. Les **18 sorties × 484 fichiers**
restent identiques à la référence historique. Cette campagne est conservée
intégralement dans `aff-before-after.json`.

Le bootstrap passe : **453 modules / 282 610 types**, **914 sources/manifests
identiques** entre JS et génération native, puis stage 2 et smoke frais. La suite
Aff complète passe également : **47 checks, cinq tests Rust, neuf scénarios
d'erreur**, concurrence/Ref/AVar et durée de vie des enfants. Le candidat qualifié
est `002ed7b50aaa2eb58c844fada98c715ebdaa1378ee8f62d1a23473d5f5e25639`.

### Confirmation sur le corpus purust-aff distinct

Sur **244 modules / 138 448 types / 27 122 224 octets**, cinq processus mesurés
après une chauffe, avec ordre tournant :

| Phase | Natif avant | Candidat | JavaScript |
|---|---:|---:|---:|
| Chargement TAST + tri | 796 ms | **643 ms** | 569 ms |
| Préparation | 599 ms | 599 ms | 38 ms |
| Optimisation + génération | 3 291 ms | 3 305 ms | 4 708 ms |
| Finalisation + émission | 724 ms | 731 ms | 814 ms |
| Backend complet | **5 540 ms** | **5 399 ms** | **6 181 ms** |

Gain natif contrôlé **2,5 %** ; le candidat prend **12,7 % de moins que JS**.
RSS maximal avant/après : **515 637 248 / 506 494 976 octets**. Les **18 sorties
× 496 fichiers** sont identiques. Archive : `purust-aff-before-after.json`.
Ces entrées diffèrent du corpus 238 : les temps ne sont pas interchangeables.

### Diagnostic de la variabilité parallèle

La référence 238 présente déjà un passage lent (3 199 ms d'optimisation/génération)
parmi ses cinq mesures. L'ordonnanceur lance des tentatives spéculatives quand
aucun module prêt n'est disponible. Si elles découvrent une dépendance encore
absente, leur conversion est achevée puis rejetée, avant une nouvelle tentative.
`attempts-ms` additionne des durées murales de tâches, pas du temps CPU.

Deux diagnostics complètent les campagnes principales :

- **Un seul worker, 238 modules**, cinq processus : total 8 966 → 8 887 ms ;
  chargement 770 → 613 ms ; optimisation/génération 6 797 → 6 855 ms. Ce contrôle
  ne montre pas le surcoût de 527 ms observé dans le premier passage parallèle.
- **Ablation dans un même exécutable**, trois processus par variante après une
  chauffe : aucun chemin natif 5 289 ms, tableaux seuls 5 273 ms, annotations
  seules 4 713 ms, les deux 5 119 ms ; ancien exécutable 5 333 ms. Les médianes
  de chargement sont respectivement 763 / 719 / 665 / 618 / 775 ms. Tous les
  chemins rencontrent les régimes rapide/lent ; ces petites séries ne permettent
  pas de classer sûrement les variantes sur le total. Les 20 sorties sont exactes.

Le graphe de valeurs et le partage des types sont conservés. L'hypothèse est
qu'un état d'allocateur différent déplace les temps relatifs des tâches, puis
que la spéculation amplifie cet écart. L'ordonnanceur reste identique. Une
confirmation à **21 paires**, fixée avant son lancement et alternant l'ordre,
donne les résultats retenus en tête de rapport ; le premier passage reste
archivé, sans exclusion de ses échantillons.

## Avant/après contrôlé

Exécutables figés, même pilote et même profil, trois processus par variante :
avant/après, après/avant, avant/après. Chaque processus fait deux chauffes et
cinq passages par phase. Les valeurs et erreurs sont vérifiées hors timing.
Les pourcentages proviennent du témoin remesuré dans cette même campagne.

| Corpus | Phase | Rust avant | Rust après | Temps en moins |
|---|---|---:|---:|---:|
| 12 modules | Decode | 122,698 ms | 93,988 ms | **23,4 %** |
| 12 modules | Combined | 145,331 ms | **116,851 ms** | **19,6 %** |
| 238 modules | Decode | 580,648 ms | 455,077 ms | **21,6 %** |
| 238 modules | Combined | 692,722 ms | **565,790 ms** | **18,3 %** |

Contrôle parsing : 18,098 → 18,279 ms sur 12 modules et 88,251 → 90,159 ms sur
238 modules. Le parseur n'a pas été modifié. La médiane des quinze échantillons
confirme le gain combined : **146,748 → 118,676 ms** et **700,371 → 571,280 ms**.

Minima de processus pour combined, en millisecondes :

| Corpus / variante | Processus 1 | Processus 2 | Processus 3 |
|---|---:|---:|---:|
| 12 / avant | 146,748 | 145,305 | 145,331 |
| 12 / après | 116,452 | 116,851 | 117,174 |
| 238 / avant | 714,259 | 692,722 | 685,785 |
| 238 / après | 570,380 | 565,790 | 556,844 |

### Changements et correction

- `decodeArrayImpl` exécute une boucle native : un appel du décodeur par
  élément, ordre conservé, arrêt au premier échec avec la même erreur `AtIndex`.
- `decodeAnnWithUsageImpl` construit directement les métadonnées, annotations
  et faits d'usage. Les références aux types restent partagées avec leur table.
  Tout cas non couvert par ce chemin se replie entièrement sur PureScript,
  préservant la priorité et les messages d'erreur.
- Le décodage global du module et la validation des portées restent ceux de
  PBO. Les valeurs temporaires et résultats conservent leur durée de vie.

Les tests Rust différentiels passent : **82 cas de tableaux**, **70 cas
d'annotations** et **110 909 annotations du corpus complet**, toutes traitées
nativement. Ils vérifient les valeurs, erreurs, compteurs d'appels, surrogates
UTF-16, bornes entières, `null`, possession des chaînes et partage des types.
Les six processus avant/après de chaque corpus concordent tous avec l'oracle
structurel figé.
La relecture indépendante a été complétée par **52 annotations numériques en
représentation native** (`Int`, véritable zéro négatif, nombres non finis) et
**16 tableaux compacts**. Les arbres d'erreur sont aussi comparés structurellement.
Ces tests passent, ainsi que 94 checks codegen, 824 tables différentielles /
149 495 entrées, et les contrats PBO : sept checks de champs, douze de faits
d'usage, cinq de négation numérique et le script de tables de types.

## Référence initiale

Le benchmark appelle les vraies fonctions PBO `jsonParser`, `decodeModule` et
`parseModule`. Rust utilise le fork **purescript-backend-optimizer-purust** et
les mêmes bibliothèques natives que le compilateur. Le pilote FFI orchestre les
mesures et la validation structurelle.

Médiane des minima de trois processus, chacun avec deux chauffes et cinq
passages par phase. Unité : millisecondes.

| Corpus | Runtime | Parsing JSON | Décodage du JSON pré-parsé | Chemin texte complet |
|---|---|---:|---:|---:|
| 12 modules | JavaScript | 20,345 | 59,083 | 80,647 |
| 12 modules | Go | 19,377 | 11,601 | 20,237 |
| 12 modules | Rust initial | 18,608 | 125,987 | **151,162** |
| 12 modules | Référence C++ | 4,314 | 6,296 | 10,560 |
| 238 modules | JavaScript | 101,633 | 289,095 | 392,464 |
| 238 modules | Go | 95,415 | 60,142 | 101,347 |
| 238 modules | Rust initial | 90,380 | 595,734 | **717,887** |

Les trois phases sont mesurées indépendamment ; leurs médianes ne s'additionnent
pas. La métrique prioritaire est le chemin complet `parseModule`.

Le parsing Rust est déjà comparable au Go. Le décodage initial prend environ
596 ms, contre 90 ms pour le parsing ; il est la première cible de ce lot.
Cette première campagne décrit la référence initiale ; les gains contrôlés
figurent dans la section avant/après ci-dessus.

## Allocations et durée de vie

Diagnostic séparé sur les 238 modules, avec un allocateur compteur autour de
mimalloc. Ses durées perturbées sont exclues des résultats de performance.

| Phase Rust initiale | Requêtes d'allocation | Octets cumulés alloués |
|---|---:|---:|
| Parse | 5 490 878 | 412 036 481 |
| Decode | 46 413 892 | 1 596 925 637 |
| Combined | 52 382 610 | 2 092 673 148 |

Après optimisation :

| Phase | Requêtes d'allocation | Octets cumulés alloués |
|---|---:|---:|
| Parse | 5 490 878 | 412 036 481 |
| Decode | **35 113 805** | **1 218 117 925** |
| Combined | **41 082 523** | **1 713 865 436** |

Soit **11 300 087 requêtes et 378 807 712 octets de moins** par passage :
**−24,3 % de requêtes pour decode**, **−21,6 % pour combined**. Les cinq passages
instrumentés du candidat donnent exactement ces mêmes compteurs. Le parsing
conserve ses compteurs initiaux.

Ces octets sont un volume cumulé, pas la mémoire résidente. Le coût de libération
finale des résultats est mesuré séparément : environ **20,8 ms** pour le chemin
complet initial sur 238 modules. Les allocations et destructions temporaires
effectuées dans le décodeur restent dans son chronométrage.
Dans la comparaison contrôlée, la médiane de destruction finale passe de
**20,621 à 19,844 ms** sur 238 modules et de **3,692 à 3,500 ms** sur 12 modules.
Rust et C++ mesurent cette destruction séparément ; les pilotes JS/Go historiques
laissent le ramasse-miettes fonctionner selon leur protocole, sans mesure isolée
de la destruction finale. Les comparaisons inter-runtimes conservent cette
différence de frontière, explicitée dans les résultats bruts.

Les profils CPU démarrent sur un signal envoyé après chargement et validation.
Le noyau répété ne calcule aucune empreinte. Les piles des workers inactifs sont
séparées ; les profils de calcul confirment le coût des allocations, libérations
et copies. Les symboles propres au parseur ne représentent qu'une partie de son
coût : leurs auto-échantillons ne bornent pas le gain d'un éventuel chemin texte
évitant le JSON intermédiaire.

## Corpus, oracles et profil de construction

- Corpus historique : **12 modules / 5 545 093 octets** de CoreFn, fixture
  `test/fixtures/json-typed-ast/`, oracle structurel historique conservé.
- Corpus Aff complet : **238 modules / 136 604 types / 26 606 491 octets**,
  TAST original figé lors de la comparaison gopurs/Purust. Manifeste :
  `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
- L'oracle 238 est produit par le décodeur PureScript JavaScript, puis figé.
  Go et Rust concordent sur toutes les empreintes JSON et TAST, à chaque
  processus. Les empreintes parcourent les modules, expressions, annotations,
  types et faits d'usage ; elles sont calculées hors chronométrage.
- La référence C++ passe les douze modules, mais diffère sur la canonicalisation
  de `1.0e-6` dans `Data.Number.Approximate` sur le corpus 238. Ce diagnostic est
  conservé ; C++ n'est pas l'oracle de ce corpus et ses temps 238 sont exclus.
- Rust : **O3, sans LTO, runtime threadé Arc, mimalloc**, `debug=false`,
  `codegen-units=16`. Ce profil de diagnostic est explicitement distinct de
  l'en-tête générique « O3 + thin LTO » de la table Rust du README.
- Go : `GOMAXPROCS=1`, `GOGC=100`, PGO désactivé. Entrées identiques pour tous
  les runtimes ; I/O du corpus et empreintes exclues. Les processus et les
  campagnes sont séquentiels, sans build concurrent.

Machine : Apple M4 Pro, 14 cœurs, 48 Gio, arm64 macOS 26.1 ; Node 24.8.0,
Go 1.27.0, Rust/Cargo 1.96.0. Aucune affinité CPU imposée.

## Reproduction et qualification

Le [harness du lot](../../bin/benchmark/purust-tast/README.md) fournit le build
Rust isolé, la comparaison avant/après des exécutables figés, la qualification
JS → stage 1 → stage 2 → smoke et la compilation Aff complète.

La comparaison isolée conserve la médiane de trois minima de processus, et
rapporte aussi la médiane des quinze échantillons. La compilation complète
emploie une autre statistique : médiane de cinq processus, puis confirmation
238 étendue à 21 paires, après une chauffe par variante, avec ordre tournant
et sorties neuves. Elle mesure `backend total`, du
chargement TAST à la fin des workers, hors frontend, Cargo, bootstrap et
démarrage/sortie du processus.

Archive locale : `var/benchmark/purust-tast-20261002/`. Les sous-répertoires
`artifacts/`, `campaign-12/`, `campaign-238/`, `alloc-238/` et `profile-238/`
conservent les références, échantillons, oracles et diagnostics initiaux.
Le binaire Rust initial du diagnostic est `39b1078f3876fca6…` ; le compilateur
natif initial est `96cf27f08bb888a0677d1cd267ded9c52a5537301587ceb1715572e2018bea26`.
Le candidat mesuré est `f67b5dadd6e0d4b8534e4b32d6a67d7f584b00f79a0c81db64c390c89f0549da`.
`candidate-array-ann/`, `array-ann-12/`, `array-ann-238/` et
`array-ann-alloc-238/` conservent son build et tous les résultats avant/après.
Le compilateur installé et mesuré est le stage 2 `002ed7b5…`, lié par
`installation.json` à `qualification/qualification.json`. Le lanceur utilise
le natif par défaut et `PURUST_JS=1` sélectionne explicitement JavaScript.

Vérification finale : **5 429 empreintes de fichiers figés**, **112 sorties /
54 424 fichiers générés** comparés ; toutes les empreintes de processus et les
médianes publiées sont recalculées. `verification.json` conserve le détail des
empreintes vérifiées. Les échecs d'environnement des quatre checks Docker,
l'entrée `Prim` sans `corefn.json` dans le test des tables et les diagnostics
de chemin du harness sont conservés avec leurs relances réussies.

Le runner de l'autre suite `JsonDecoding` a aussi été reconstruit et contrôlé :
les 17 cas conservent leurs empreintes dans neuf processus, trois par runtime
JS/Go/Rust. Ces temps de contrôle ne remplacent pas les résultats publiés de
cette autre suite.
