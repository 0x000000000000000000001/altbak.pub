# Gopurs exécuté en JavaScript, Go et Rust

Qualification du **3 octobre 2026**. Les trois exécutables compilent le même
TAST de `gopurs-aff` et **produisent du Go**.

Cette première qualification sert de référence à l'[optimisation qualifiée
ultérieure](2026-10-03-gopurs-rust-optimization.md), désormais installée.

| Hôte de gopurs | Médiane | Moyenne | Min–max |
|---|---:|---:|---:|
| JavaScript / Node | 7 718,5 ms | 7 711,6 ms | 7 588–7 919 ms |
| Go natif | **2 116,5 ms** | 2 121,5 ms | 2 077–2 184 ms |
| Rust natif | **5 782,5 ms** | 5 763,9 ms | 5 577–5 885 ms |

Dans cette campagne, Rust prend **25,1 % de temps en moins que JavaScript** et
**2,732× le temps du natif Go**. Ce premier port de gopurs en Rust est fonctionnel
et qualifié ; les optimisations propres au générateur Rust de Purust ne constituent
pas des optimisations mesurées de ce compilateur Go.

## Commandes et construction

Dans `gopurs/gopurs-aff` :

```sh
./bin/test                 # gopurs natif Go → application Go
GOPURS_JS=1 ./bin/test     # gopurs JavaScript → application Go
GOPURS_RUST=1 ./bin/test   # gopurs natif Rust → application Go
GOPURS_RUST=1 ./bin/test -c # reconstruit gopurs en Rust avant le test
```

`npm run build:rust`, dans `gopurs/gopurs`, produit `bin/gopurs-rust` à partir
des sources PureScript de gopurs, de `purescript-backend-optimizer-gopurs` et
des bibliothèques hôtes `purust-*`. Purust traduit ce compilateur en Rust ; Cargo
construit l'exécutable en **O3, sans LTO**, avec Arc et mimalloc. Le parseur FFI
Go existant est lié par une archive C ; le compilateur Rust l'appelle directement.

`PURUST_JS=1` choisit le compilateur de bootstrap de Purust lorsqu'on reconstruit
gopurs. L'exécution de `gopurs-rust` sélectionnée par `GOPURS_RUST=1` est native.
Les anciens scripts qui faisaient générer une application Rust par Purust ont
été retirés du runner Aff.

## Protocole

- Corpus original du 2 octobre : **238 modules / 136 604 types /
  26 606 491 octets** de TAST. Entrées et FFI recopiées depuis l'archive figée,
  avec toutes les empreintes revérifiées.
- Une chauffe et **dix mesures par hôte**, processus sérialisés et ordre tournant.
  Aucun échantillon mesuré n'est exclu.
- `GOPURS_JOBS`, `GOPURS_PREPARE_JOBS`, `GOPURS_PBO_JOBS` et `GOPURS_EMIT_JOBS`
  sont fixés à **8** pour les trois hôtes ; `GOPURS_PIPELINE=1`.
- Ces paramètres explicites diffèrent notamment du défaut PBO Rust, qui est
  séquentiel. Le launcher Go applique sa politique de GC pour cette machine
  de 48 Gio : `GOGC=off`, `GOMEMLIMIT=10GiB`. Le parseur embarqué de l'hôte Rust
  conserve les paramètres GC Go ordinaires.
- Mesure : `[gopurs] backend total`, incluant chargement/tri, préparation,
  optimisation, génération/émission Go et fin des workers. Frontend, bootstrap,
  compilation/exécution applicative et démarrage/sortie du processus sont exclus.
- Sorties et caches `.purmeta`/`.cache` neufs par processus ; cache de fichiers
  du système non purgé. Apple M4 Pro, 14 processeurs logiques, Node 24.8.0.
- **33 générations × 294 fichiers Go/manifests = 9 702 fichiers** identiques
  octet par octet à l'oracle Go du corpus original.

### Échantillons bruts, millisecondes

| Tour | JavaScript | Go | Rust |
|---:|---:|---:|---:|
| 1 | 7 744 | 2 126 | 5 770 |
| 2 | 7 588 | 2 113 | 5 634 |
| 3 | 7 701 | 2 122 | 5 577 |
| 4 | 7 613 | 2 184 | 5 847 |
| 5 | 7 794 | 2 146 | 5 797 |
| 6 | 7 632 | 2 077 | 5 795 |
| 7 | 7 600 | 2 115 | 5 740 |
| 8 | 7 789 | 2 118 | 5 885 |
| 9 | 7 736 | 2 111 | 5 737 |
| 10 | 7 919 | 2 103 | 5 857 |

### Médianes des phases

| Phase | JavaScript | Go | Rust |
|---|---:|---:|---:|
| Chargement TAST + tri | 532,5 ms | 55 ms | 730 ms |
| Préparation + monomorphisation | 462 ms | 694,5 ms | 1 646,5 ms |
| Runtime | 2 ms | 1 ms | 7 ms |
| Optimisation + émission | 6 714 ms | 1 359,5 ms | 3 345 ms |
| Points d'entrée | 1 ms | 0 ms | 18 ms |

Les médianes des phases ne s'additionnent pas. La spécialisation transitive est
une sous-phase de préparation, déjà incluse dans sa durée.

## Qualification

- Bootstrap : **500 modules / 290 636 types** ; les générations par Purust JS
  et Purust natif donnent **1 008 fichiers Rust/Cargo identiques**. Un projet
  frais avec FFI Go est compilé et exécuté avant l'installation atomique.
- `GOPURS_RUST=1 ./bin/test -c` passe la reconstruction complète et l'exécution
  de l'application Go.
- Suite Aff actuelle : **45 contrôles et stress AVar de 1 000 éléments**,
  réussis pour chaque hôte. Le TAST des trois exécutions est identique et les
  **294 fichiers Go** produits sont identiques. Les tests synchronisés actuels
  constituent la qualification fonctionnelle ; le corpus historique reste celui
  des chronométrages.
- **37 tests Node sans échec ni skip** : succès et diagnostics CLI séquentiels/
  parallèles, puis cinq fixtures exécutées via les trois hôtes : chaînes et
  surrogates, retours FFI, workers de records, arbres possédés et plan JSON.
- **Neuf contrôles auxiliaires** de bootstrap Go et d'embarquement du runtime
  passent également. Le parseur Go passe `go test -race -tags=carchive`.
- Le nouveau snapshot `CompilerHostStrings` est admis par le runner normal après
  compilation et exécution réussies.

### Diagnostics conservés

La première construction a rencontré une ambiguïté entre deux imports Rust
nommés `Value` dans la FFI de `Main`. Le nom est maintenant qualifié avec
`purust_core` ; le diagnostic initial et les reprises restent archivés.

Le premier contrôle sur `StringEscapes` a détecté un écart **déjà présent entre
gopurs JS et gopurs Go** : la constante inutilisée `loneSurrogates` est repliée
différemment. Son assertion était déjà désactivée dans cette fixture. Cette
comparaison échouée est conservée et n'est pas comptée comme réussie. La nouvelle
fixture `CompilerHostStrings` conserve les chaînes à travers une référence
mutable et vérifie leur émission/exécution sur les trois hôtes, avec sorties Go
strictement identiques, y compris les surrogates isolés.

## Exécutables et archives

| Exécutable | SHA-256 |
|---|---|
| gopurs JS | `b21ddc47cc5cc6f48a830c355812d3be328301363a7f8db695679e65e9e534ac` |
| gopurs Go | `bae4835bdb63ee802f172253f6a5e88ef03bc141bb15330318ba4d1a54674c16` |
| gopurs Rust de cette qualification initiale | `ad7fd6a9592ccef0e5a33615c1b79ede669589898aa3742f0bba184de07e7e8e` |

- [Résumé JSON vérifié](2026-10-03-gopurs-rust-host.json).
- Archive locale : `var/benchmark/gopurs-rust-host-20261003/` : bootstrap,
  `qualification/`, `timings/`, logs, sources, exécutables et diagnostics initiaux.
- Les caches Cargo temporaires des constructions terminées sont nettoyés après
  vérification des empreintes ; les exécutables qualifiés restent conservés.
- Harnais : `bin/benchmark/gopurs-aff/{qualify-hosts,compare-hosts,publish-hosts}.mjs`.

Les résultats historiques de l'[optimisation de Purust](2026-10-03-purust-native-optimization.md)
restent valides pour **Purust générant du Rust**. La ligne `gopurs-aff` du tableau
« compilation vers Go » utilise désormais exclusivement les trois hôtes de gopurs.
