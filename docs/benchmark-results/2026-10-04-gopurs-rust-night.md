# gopurs Rust nuit — 2026-10-04

Statut : **publié** ;
tous les gates requis passent.

## Résultat

Les trois hôtes exécutent **gopurs et génèrent du Go**. Le corpus figé
`gopurs-aff` contient **238 modules / 136604 types**
(26606491 octets), manifeste TAST `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
Machine : Apple M4 Pro, 14 processeurs logiques,
48 Gio de RAM.

La revendication de victoire est établie : campagne finale `final-campaign.json`
     avec `victory=true`, **30/30 paires favorables**, intervalle bootstrap
      [0.8418, 0.8468] entièrement sous 1, corroboration aux défauts publics
      (10/10 paires, intervalle [0.8402, 0.8463]).

| Contrôle | Valeur |
| --- | --- |
| Campagne finale | final-campaign.json |
| Victoire | true |
| Paires primaires / favorables | 30 / 30 |
| Intervalle bootstrap primaire (ratio Rust/Go) | 0.8418 … 0.8468 |
| Médianes primaires Go → Rust | 1771 → 1495 ms |
| Écart médian Rust/Go | -15.58 % |
| Corroboration défauts | 10/10 favorables, 1787.50 → 1505.50 ms |
| Rust début de nuit → production (campagne indépendante) | 3218 → 1502.50 ms (-53.31 %) |
| Échantillons bruts revérifiés | 147 |

## Candidats mesurés (dossiers *-runs passés)

| Campagne | Tours | Variantes (médianes ms) | Statut |
| --- | --- | --- | --- |
| aff-resume-pool | 5 | go: 1992 ; control: 3306 ; aff-resume-pool: 3275 | mesuré, non qualifié |
| baseline | 5 | go: 2050 ; rust: 3361 | mesuré, non qualifié |
| before-after | 10 | rust-before: 3218 ; rust: 1502.50 | confirmation qualifiée |
| common | 10 | js: 6893 ; go: 1787 ; rust: 1507.50 | confirmation qualifiée |
| default | 10 | go: 1787.50 ; rust: 1505.50 | confirmation qualifiée |
| directive-chain-v2 | 5 | control: 3308 ; directive-chain: 3265 | mesuré, non qualifié |
| native-aff-io-v2 | 5 | control: 3253 ; native-aff-io-v2: 2523 ; go: 1967 | mesuré |
| native-collections | 5 | shared-classes: 2119 ; native-collections: 1935 ; go: 1955 | mesuré |
| native-io-pool | 5 | go: 1981 ; control: 3303 ; native-io-pool: 2534 | mesuré, non qualifié |
| native-tast-text-v6 | 5 | shared-classes: 2108 ; native-tast-text-v6: 1995 ; go: 1946 | mesuré |
| parallel-rewrite | 5 | shared-classes: 2188 ; parallel-rewrite: 2040 ; go: 2062 | mesuré |
| primary | 30 | go: 1771 ; rust: 1495 | confirmation qualifiée |
| profiles | 5 | control: 3368 ; cgu1: 3406 ; cgu4: 3430 ; native-cgu1: 3431 | mesuré, non qualifié |
| resources | 3 | go: 1815 ; rust: 1496 | confirmation qualifiée |
| selected-source-pgo | 5 | selected-source: 1679 ; selected-source-pgo: 1494 ; go: 1967 | mesuré, non qualifié |
| selected-source | 5 | transitive-rounds: 1898 ; selected-source: 1675 ; go: 1984 | mesuré |
| selected-source-v2-pgo | 5 | selected-source: 1677 ; selected-source-v2-pgo: 1507 ; go: 1968 | sélectionné |
| shared-classes | 5 | native-aff-io-v2: 2517 ; shared-classes: 2122 ; go: 1952 | mesuré |
| transitive-rounds | 5 | native-collections: 1907 ; transitive-rounds: 1877 ; go: 1930 | mesuré |

Les campagnes mesurées ne sont pas qualifiées par cette extraction : seule la campagne
finale et les gates configurés décident.

## Rejetés

- **native-io-pool** (2534 ms) — déplaçait des callbacks Effect hors de leurs checkpoints et changeait leur ordre observable ; remplacé par les opérations Node.FS.Aff ciblées
- **directive-chain-v2** — signal insuffisant (3/5 paires favorables) ; préfixe/range retiré des sources vivantes
- **CGU / CPU natif** — aucun gain démontré face au profil O3 + ThinLTO sans debug
- **première série Aff 35/35** — comparaison invalidée : le candidat lançait Test.Lifetime et les témoins Test.Main ; logs conservés, aucun résultat de cette série ne qualifie bracket

## Préliminaires / non qualifiés

- **sélections isolées à cinq paires** — témoins contemporains distincts, gains non additionnables ; qualification et publication fondées sur la composition reconstruite
- **native-aff-io-v2** (2523 ms), 17 tests — mesure antérieure au correctif de reprise des timers ; la composition utilise le runtime corrigé, qualifié par 19 contrats et la suite Aff complète
- **test bracket à échéance de 40 ms** — course d'observation sous charge conservée dans les diagnostics ; test synchronisé par joinFiber, état et résultat vérifiés immédiatement, même test validé en Rust et JavaScript
- **textes v1–v5** — échecs de fixtures, validation foreignAnnotations et résolution d'alias de packages archivés ; v6 passe les 4030 comparaisons avec un oracle PS indépendant
- **première tentative PGO** (1494 ms) — invalidée par le contrôle du manifeste : le nettoyage de output/main supprimait Main/corefn.json sur APFS insensible à la casse ; entraînement amputé, résultat conservé comme failed et non sélectionné

## Phases (p50, non additives)

| Variante | backend total | load + tri | transitive | prepare + mono. | PBO optimize + emit | entry points |
| --- | --- | --- | --- | --- | --- | --- |
| js | 6893 | 528 | 174 | 442 | 5913 | 1 |
| go | 1787 | 57 | 200 | 489.50 | 1238.50 | 0 |
| rust | 1507.50 | 57 | 167 | 441.50 | 993 | 13 |

| Variante | Producteur PBO (avec attente émission) | Génération + écritures (lots cumulés) | Drain final émission |
| --- | --- | --- | --- |
| js | 4843.50 | 5854.50 | 1071.50 |
| go | 894.50 | 1106 | 343 |
| rust | 752.50 | 791 | 238 |

`backend total` couvre chargement/tri, préparation, PBO, génération/émission Go, points
d'entrée et drain des workers ; `prepare + monomorphize` inclut `transitive
specializations`. Les phases **ne s'additionnent pas** et les familles de profil se
recouvrent.
Le producteur PBO inclut sa contre-pression vers l'émetteur ; les temps cumulés
des lots de génération se superposent à son exécution en mode pipeline. Le drain
mesure l'attente finale de l'émetteur et reste inclus dans `optimize + emit`.

## Ressources (processus entier, CPU + RSS)

| Variante | Paires | CPU médian (s) | RSS médian (Mio) |
| --- | --- | --- | --- |
| go | 3 | 7.62 | 9634.50 |
| rust | 3 | 5.03 | 452.42 |

Parsing brut de `/usr/bin/time -l` (Mac), processus entier avec descendants.

Sur cette machine, le lanceur applique à l'hôte Go sa politique mémoire
`GOGC=off` / `GOMEMLIMIT=10GiB` (seuil d'au moins 32 Gio de RAM).
Les chiffres de ressources reflètent cette configuration publique ; le réglage
automatique du GC Go n'est pas appliqué à l'hôte Rust.

## Qualification

| Compte | Valeur |
| --- | --- |
| production_identical_rust_files | 1009 |
| bootstrap_identical_files | 914 |
| compiler_suite_pass | 43 |
| native_aff_contracts | 19 |
| native_pbo_suites | 8 |
| generator_codegen_tests | 95 |
| generator_tast_suites | 43 |
| text_full_oracle_cases | 4030 |
| text_frozen_modules_without_fallback | 238 |
| composition_text_cases | 289 |
| production_helper_tests | 47 |
| go_target_aff_checks_per_host | 45 |
| go_target_avar_stress_items_per_host | 1000 |
| native_aff_test_main_js_oracle_checks | 47 |
| gopurs_bootstrap_tast_modules | 500 |
| gopurs_bootstrap_tast_types | 291171 |
| gopurs_js_native_identical_rust_files | 1008 |
| purust_self_host_tast_modules | 453 |
| purust_self_host_tast_types | 282796 |
| native_type_render_synthetic_checks | 948 |
| native_type_render_corpus_checks | 409812 |
| production_pgo_training_generated_files_per_pass | 612 |
| final_verified_generations | 147 |
| final_identical_files | 43218 |

| Hôte | Contrôles Aff | Fichiers Go exacts | Modules TAST | Stress AVar |
| --- | --- | --- | --- | --- |
| rust | 45 | 294 | 238 | 1000 |
| go | 45 | 294 | 238 | 1000 |
| js | 45 | 294 | 238 | 1000 |

Candidat production `selected-source-v2-pgo`, profil `{"opt_level":3,"debug":false,"lto":"thin","threaded":true,"allocator":"mimalloc","pgo":true}`.

Bootstrap Purust indépendant : 914 fichiers identiques, TAST 453 modules.


### PGO de production

Le build public reconstruit son profil sur **500 modules
d'auto-compilation du compilateur**, puis installe le binaire qualifié.
Les **234 modules de bibliothèque communs** avec
le corpus tenu à l'écart sont nommés dans le rapport JSON ; `Test.Main` est exclu.
Les trois passes d'entraînement conservent le manifeste des entrées et reproduisent
exactement les sorties Go du compilateur sans profil.

- Sources Rust générées : `cc097135936ccfd40e9f727db333e4d72016dba8a72de0a219901a4bb4208bfe`.
- Entrées d'entraînement : `8b589139388aac64ef2eea98c201c354402c6a2adbaff4d4e7b211f124120171`.
- Profil fusionné : `db05826ede6d5884d9361d44260c1ec585c3696f132791ba2199697fca6ec93c`.
- Binaire installé : `526a5d1adf0f86279ec146d8333126e6a82d07d2f5b1ed0a817c89cf9073237d`.
- Métadonnées complètes : `/Users/0x1/Documents/htdocs/altbak.pub/var/benchmark/gopurs-rust-night-20261004/production/gopurs-rust-build-Jhjgzy/pgo-profile.json`.

La sélection isolée et ce build public ont des profils distincts : les mesures
finales ci-dessus portent sur le **binaire installé**. Le profil se régénère par
`npm run build:rust` ; `GOPURS_RUST_PGO=0` permet de reconstruire le témoin sans PGO.


## Historique

Rapport `2026-10-03-purust-native-optimization.json` conservé (sha256 `a275404028d5318cf0d538e6c16395db3857861528c997612d148dfbf6369d59`) :
fixture12 12 modules / 9 exécutions, gopurs238 238 modules / 9 exécutions.
L'axe **Purust → Rust** (`goRust.summary.rust-final.median_ms` = 2569 ms)
reste distinct de l'axe **gopursRust → Go** mesuré ici : ce chiffre n'est ni redéfini ni
réutilisé.

| Rapport JSON conservé | SHA-256 |
| --- | --- |
| 2026-10-03-purust-native-optimization.json | `a275404028d5318cf0d538e6c16395db3857861528c997612d148dfbf6369d59` |
| 2026-10-03-gopurs-rust-optimization.json | `d0cff61ba8a7d735b9b780e27e22ec778e660e2b4a308a0ed7f3a40c6f086f8b` |
| 2026-10-03-gopurs-host-defaults.json | `814e75970e0756a3e44f3bd83dd69759d85308f226f1d67c9a8b12de2c44896c` |
| 2026-10-03-gopurs-rust-allocation.json | `695a4ac5172770321b5e1724d297be2d2881ccc19542d963737c7651904f1ff1` |

## Méthode, gates et reproduction

- Commande : `node altbak.pub/bin/benchmark/gopurs-rust/publish-night.mjs gopurs-rust-night-20261004 publication-gates.json --publish`
- Métrique : `backend total` (chargement/tri, préparation, PBO, génération/émission Go, points d'entrée, drain) ;
  frontend, builds, exécution et démarrage/sortie du processus exclus.
- Corpus et compilateurs gelés, un échauffement par variante, processus sérialisés,
  ordre de passage tournant, sorties et caches PBO recréés à chaque génération.
  Comparaisons explicites : chargement/préparation/PBO/émission 8/8/8/8 et pipeline actif.
  Les dix paires « default » passent par le lanceur public sans ces surcharges.
- Intervalle à 95 % : rééchantillonnage apparié des logarithmes des ratios Rust/Go,
  100 000 réplications déterministes ; chaque paire reste groupée.
- Échantillons bruts relus : 147.
- Archive : `/Users/0x1/Documents/htdocs/altbak.pub/var/benchmark/gopurs-rust-night-20261004`.

| Gate | Type | Attendu | Trouvé | Requis | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| production | production | passed | passed | oui | `5a8bb3c81e96bb30990ea5ff389164ee2c2ddff81089d1920ff6cf03354ba251` |
| hosts | hosts | passed | passed | oui | `ac33ff525aec6cfd73aff55616e54862a5249fa263d7a9f03d6a1bee99473be7` |
| bootstrap | bootstrap | passed | passed | oui | `f3647a23d10d29ca28577d1a87c897f8bb64d7fa2595d581bea94aee49819883` |
| primary | runs | passed | passed | oui | `d4ce09affd439e813722501e251c728a73ddb5164942040eac5851a9edd48f5d` |
| default | runs | passed | passed | oui | `4bd0628d103e513770d6a9a15e7455a7996cab55c15b8b533e66b0b546f0371c` |
| common | runs | passed | passed | oui | `4c441fa153370d13cec22eb25a811e0aad2acf459897a9f26eb0f7e97455cb55` |
| resources | runs | passed | passed | oui | `8ad31fee29cfbd53b3a81dd1230071bbcb6bd4b61d5ff8c1c23c20f57445a1af` |
| before-after | runs | passed | passed | oui | `4d6756d3553be10aa1bdbdefa3ba07321da70135f39a3d547203f3f745e8e74c` |
| verification | verification | passed | passed | oui | `607c468c01d2eb1d0ba1109dba2896e81871835c479e47811f9eb11198f0c943` |
| final | final | passed | passed | oui | `ee15a2321eaf0163286102f676649596198bb2affd8d49272c1b76002c0d7419` |
| source-composition | other | passed | passed | oui | `7ecf841501e42dcdf9fe4f7e623662f73243a78e5c357308ecdcd2c4b36ccc38` |
| native-pbo | other | ok | ok | oui | `289764d4f461404d53d0df366de7b13422c7204c3fa0a141a759392c92b464ef` |
| native-pbo-provenance | other | passed | passed | oui | `08c094e29e3af18454c751604e57e4ba3fd33e765d30ef433b5169ed7181d600` |
| full-text-oracle | text | passed | passed | oui | `46eaa7dda5ddc72f8bc3fd35326886777e95368e95b6f747bbc216d3542ffdf4` |
| aff-synchronized | other | passed | passed | oui | `dd1ada7cfbc2aa0940b934270889ccb5bd1e3b4465ae2b2afdf816907a7f16fa` |
| runtime-fs-promise-unfoldable | other | passed | passed | oui | `ca29330222cc0089c616613542b5151e5239016c932122d10f0a139c8219c014` |
| generator-representation | other | passed | passed | oui | `27b1a035f07c333a2c18873439b64ce19ffc015310e68ee75599b0d1a6a0d5ea` |
| source-selection | selection | passed | passed | oui | `56061e5cb2934364b3b8d791885bb7794d8fa3613006a1aa42641fd00cb2aa39` |
| pgo-training | other | passed | passed | oui | `991d9fe9a875df69dbd03eff27d034bdf6cd6ec769ec974d294de83c6c160a93` |
