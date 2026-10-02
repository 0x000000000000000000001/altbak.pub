# gopurs-aff — compilation avec gopurs et Purust

Mesuré le **2 octobre 2026**, sur un TAST commun de **238 modules / 136 604 types**.

**gopurs natif : 2 146 ms ; Purust natif : 5 495 ms**, soit **2,56× le temps de
Go** pour Rust. Purust natif prend néanmoins **11,9 % de temps en moins que
Purust JS**. Cette comparaison indique une marge de progression importante ;
elle ne mesure pas un nouveau gain d'optimisation de Purust.

**Qualification applicative :** la suite originale passe 5/5 en Go et 2/5 en
Rust dans la série de diagnostic. Une validation séparée avec synchronisations
explicites passe **5/5 sur chaque cible**. Les temps ci-dessous portent toujours
sur le **TAST original**, sans modification des tests. La colonne Rust du README
porte un renvoi à cette réserve.

## Résultats

Médiane de cinq mesures, après une chauffe par variante. Unité : millisecondes.

| Compilateur exécuté | Cible générée | Médiane backend | Min–max | Médiane processus |
|---|---|---:|---:|---:|
| gopurs.js / Node | Go | **7 630** | 7 418–8 057 | 7 721 |
| gopurs natif / Go | Go | **2 146** | 2 085–2 220 | 2 220 |
| purust.js / Node | Rust | **6 238** | 6 116–6 341 | 6 332 |
| Purust natif / Rust | Rust | **5 495** | 5 004–5 809 | 5 526 |

- gopurs natif / gopurs.js : **0,2813**, soit **71,9 % de temps en moins**.
- Purust natif / purust.js : **0,8809**, soit **11,9 % de temps en moins**.
- Purust natif / gopurs natif : **2,5606**. Rejoindre ce temps Go demanderait
  environ **61 % de réduction supplémentaire** du temps Rust, à charge constante.

### Échantillons bruts

| Passage | gopurs.js | gopurs natif | purust.js | Purust natif |
|---|---:|---:|---:|---:|
| Chauffe, exclue | 7 782 | 2 151 | 6 221 | 5 447 |
| 1 | 7 489 | 2 128 | 6 116 | 5 534 |
| 2 | 8 012 | 2 185 | 6 341 | 5 809 |
| 3 | 7 630 | 2 146 | 6 238 | 5 265 |
| 4 | 7 418 | 2 085 | 6 268 | 5 495 |
| 5 | 8 057 | 2 220 | 6 227 | 5 004 |

Tous les échantillons de la campagne finale sont conservés, sans sélection.
Le premier compilateur tourne entre les passages : Go JS, Go natif, Rust JS,
Rust natif, puis Go JS. Les quatre compilations de chaque passage sont
séquentielles. Les moyennes de charge à une minute relevées avant les exécutions
vont de **4,16 à 11,26** ; elles sont conservées dans le JSON.

### Phases

| Phase | gopurs.js | gopurs natif | purust.js | Purust natif |
|---|---:|---:|---:|---:|
| Chargement TAST + tri | 581 | **56** | 577 | **804** |
| Préparation, avec monomorphisation chez gopurs | 484 | 703 | 38 | 613 |
| Optimisation + génération/émission principale | 6 553 | **1 386** | 4 752 | **3 299** |
| Finalisation + émission séparée chez Purust | — | — | 805 | 681 |
| **Backend total** | **7 630** | **2 146** | **6 238** | **5 495** |

Les médianes des phases ne s'additionnent pas. Les pipelines ont des découpages
différents : gopurs inclut son émission dans `optimize + emit`, tandis que Purust
a une phase finale distincte. Les spécialisations transitives de gopurs
prennent 239 ms en médiane native, **déjà incluses** dans sa préparation.

## Protocole et périmètre

- Frontend PureScript exécuté une seule fois, avec la configuration et les
  dépendances de `gopurs-aff`. Snapshot des sources, TAST, FFI et compilateurs.
- **26 606 491 octets de TAST**, 238 modules, 136 604 types ; SHA-256 du manifeste :
  `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
- Même corpus original pour les quatre variantes, `--main Test.Main` ; Purust
  reçoit aussi `--threaded` et les adaptateurs FFI Rust du snapshot.
- Un processus neuf, sortie générée neuve et suppression de `.purmeta`/`.cache`
  à chaque passage. Le cache de fichiers du système n'est pas purgé.
- Une chauffe par variante, puis cinq passages, ordre tournant, médiane.
- Mesure primaire : `[gopurs|purust] backend total`, chronomètre monotone interne.
  Elle inclut chargement/tri, préparation, PBO, génération et émission jusqu'à
  la fin du travail des workers. Démarrage/sortie du processus, frontend `purs`,
  bootstrap, `go build`, Cargo et exécution applicative sont exclus.
- Les launchers de production sont figés. Go natif : `GOGC=off`, limite mémoire
  `10GiB`, 8 workers de préparation, 8 PBO, 8 d'émission. Purust natif stage 2 :
  O3 sans LTO, budget **4 PBO + 4 codegen** ; Purust JS reste séquentiel.
  Les variables de surcharge de l'environnement sont retirées avant lancement.
- Les réglages de production sont comparés, sans égalisation des budgets mémoire
  ou des phases. gopurs effectue sa monomorphisation ; Purust utilise son pipeline
  Rust par défaut. L'équivalence des entrées ne signifie pas des passes identiques.

Machine : **Apple M4 Pro, 14 cœurs, 48 Gio, arm64, macOS 26.1** ; Node **24.8.0**,
Go **1.27.0**, Rust/Cargo **1.96.0**. Pas d'affinité CPU imposée. Les builds de
qualification sont terminés avant la campagne finale de 15:44:32 à 15:46:44 UTC.

## FFI Rust et validation

Les interfaces PureScript de `gopurs-aff` et `purust-aff` diffèrent, notamment
pour les fibers et les callbacks Aff. Le benchmark conserve les bibliothèques
PureScript de gopurs et emploie les implémentations Rust existantes de Purust
avec **six adaptateurs locaux** :

- `Effect.Aff` : handles de fibers et callbacks erreur/succès ;
- `Effect.AVar` : arguments des constructeurs de résultats ;
- `Effect.Ref` : `modify_` atomique ;
- `Control.Monad.ST.Internal` et `Data.Array.ST` : noms des imports étrangers ;
- `Data.Traversable` : argument supplémentaire d'append des tableaux.

Les 49 FFI Rust importées et leurs versions originales sont archivées et
empreintées. Trois modules non exercés par `Test.Main` (`Control.Extend`,
`Data.Int.Bits`, `Performance.Minibench`) reçoivent des traps explicites qui
paniqueraient s'ils étaient appelés, plutôt que des résultats par défaut.
La validation porte sur le programme Aff exercé, pas sur toutes les API de ces
trois modules.

### Identité des fichiers générés

- **12 sorties Go × 294 fichiers**, identiques à la référence Go construite.
- **12 sorties Rust × 484 fichiers**, identiques à la référence Rust construite.
- Les comparaisons incluent les variantes JS et natives, les sources et les
  manifests. Aucune comparaison textuelle n'est attendue entre Go et Rust.
- Vérification finale : **1 032 fichiers figés**, **26 sorties avec les deux
  références / 10 114 fichiers générés**, empreintes et médianes vérifiées.

### Exécution de la suite originale

Les applications Go et Rust ont été compilées. Pour Rust, la série contrôlée
utilise un binaire applicatif **O3, sans LTO ni debug info**, hors chronométrage.

| Série originale, cinq passages | Go | Rust |
|---|---:|---:|
| Suite complète : 45 checks imprimés + stress AVar de 1 000 échanges | **5/5** | **2/5** |

Échecs Rust de cette série : `kill/finalizer/bracket` aux passages 1 et 5,
`parallel/mixed` au passage 4. Le premier suppose qu'une acquisition a déjà
commencé lors de l'annulation ; le second impose les gagnants de courses dont
les timers ne sont espacés que de 1 ms. Trois tentatives préliminaires sont
également archivées : `parallel/mixed` puis `bracket` en développement,
`bracket` en O3. Aucun de ces échecs n'est effacé ou transformé en succès.

### Validation commune avec synchronisations explicites

Dans un workspace **séparé**, dix tests reprennent les variantes de `purust-aff`
qui attendent explicitement les acquisitions/terminaisons, utilisent des AVars,
élargissent certains écarts de timers et vérifient les invariants indépendants
de l'ordre des branches. Le test `bracket` lit sa référence après la fin du
bracket plutôt que dans une autre fiber après un délai supposé suffisant.

Seul **`Test.Main` change : les 237 autres TAST sont identiques à l'original**.
Les 45 checks, le stress AVar et les adaptateurs sont communs aux deux cibles.
Après un premier passage de qualification réussi, une série fixée à cinq
passages par cible donne **Go 5/5, Rust 5/5**. Ce résultat valide les adaptateurs
sur cette suite synchronisée ; il ne rend pas stable la suite originale.
Le corpus synchronisé n'entre dans **aucun temps publié**.

## Ce que suggère la comparaison pour la suite

1. **Chargement TAST** : 804 ms en Rust contre 56 ms en Go, soit un écart de
   748 ms. Reprofiler parsing JSON et construction modules/annotations, au-delà
   du décodage natif de la table de types déjà optimisé.
2. **Optimisation/génération Rust** : 3 299 ms, plus 681 ms de finalisation.
   Reprofiler les conversions de types, boxing/unboxing, allocations PBO et
   construction des chaînes, qui restent les pistes du plan Purust.
3. Mesurer chaque nouveau lot sur ce corpus figé puis sur `purust-aff`, en gardant
   les comparaisons exactes et la qualification auto-hébergée habituelle.

Les anciens **5 743 ms** de `purust-aff` correspondent à **244 modules et
138 448 types** : la différence avec les 5 495 ms présents ne doit pas être
présentée comme un gain de compilateur.

## Provenance et reproduction

| Élément | Révision ou SHA-256 |
|---|---|
| Source gopurs-aff | `6ed72325e4d56ecc0c9c52fc43524fe8a5534a2d` |
| Checkout gopurs au snapshot | `6aec7c603f68cbc71fb64c33fbff736e05261692` |
| Checkout Purust au snapshot | `3ef079d2cfe28596c4c3cf9617a1a464f10a1c46` |
| Binaire gopurs natif | `bae4835bdb63ee802f172253f6a5e88ef03bc141bb15330318ba4d1a54674c16` |
| Bundle gopurs JS | `c665f10f5c7e42551b0da5e90f1b9016e69d76c1c0156df177f1d41bec4c5298` |
| Binaire Purust natif | `96cf27f08bb888a0677d1cd267ded9c52a5537301587ceb1715572e2018bea26` |
| Bundle Purust JS | `969bf49f5da903506d6821caff0342f0f50161c48156a41078cbd1b9721c0940` |

- [Mesures brutes, manifests et résultats de validation](2026-10-02-gopurs-aff-go-rust-compilation.json).
- [Harness et commandes](../../bin/benchmark/gopurs-aff/README.md).
- Archive locale : `var/benchmark/gopurs-purust-aff-20261002/`.
  `results.json`, `verification.json`, `logs/`, `generated/`, `compilers/`,
  `inputs/`, `rust-ffi-originals/`, `rust-ffi/`, `application-diagnostic/`,
  `portable-validation/` et `harness-final/` conservent les éléments nécessaires.
- `ffi-path-diagnostic.json` conserve une première chauffe interrompue : un
  `--ffi-dir` absolu n'était pas résolu de la même manière en JS et en natif.
  Le harness emploie désormais un chemin relatif. Les trois échantillons de
  cette tentative sont exclus des médianes, car les sorties n'étaient pas
  équivalentes. La campagne finale repasse toutes les comparaisons exactes.

Les autres paquets `gopurs-*` ne sont pas mesurés en Rust dans cette campagne.
Les totaux JS/Go du README sont les sommes des valeurs affichées, mélangeant
les lignes historiques et cette nouvelle ligne Aff ; ils ne constituent pas
une nouvelle campagne complète.
