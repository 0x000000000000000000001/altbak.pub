# Gopurs hébergé en Rust — optimisation qualifiée

Campagne du **3 octobre 2026**. Le véritable **gopurs hébergé en Rust, produisant du Go**,
passe de **5 886 à 3 592 ms (−39 %)** sur les
238 modules Aff figés. Confirmation sur **15 paires**, dont **15 favorables**.
Le compilateur optimisé est installé et la commande **`GOPURS_RUST=1 ./bin/test`** passe.

**Réglages de cette campagne :** les mesures ci-dessous fixent explicitement les
workers à 8/8/8/8. Le launcher Rust conservait alors préparation 2 / PBO 1 sans
ces variables ; la commande ordinaire avait été validée fonctionnellement.
La [correction ultérieure des valeurs par défaut](2026-10-03-gopurs-host-defaults.md)
est maintenant installée : elle aligne le parallélisme et qualifie la commande
exacte sans réglages supplémentaires à **3 683 ms**, avec une médiane figée de
**3 701 ms** sur cinq tours.

## Comparaison commune : trois hôtes, une cible Go

| Hôte de gopurs | Médiane | Moyenne | Min–max |
|---|---|---|---|
| JavaScript / Node | 7 549 ms | 7 472,6 ms | 7 031–7 676 ms |
| Go natif | 2 165,5 ms | 2 144,2 ms | 2 033–2 217 ms |
| Rust avant | 6 176 ms | 6 131,1 ms | 5 794–6 353 ms |
| Rust optimisé | 3 780,5 ms | 3 713,9 ms | 3 526–3 828 ms |

Dix tours, une chauffe par variante, ordre tournant. Dans cette même campagne,
le rapport Rust/Go passe de **2.852× à 1.746×** ;
**59,7 % de l'écart absolu est résorbé**. Rust prend encore
74,6 % de plus que Go et 49,9 %
de moins que JavaScript. La parité stricte reste à atteindre.

| Tour | JS | Go | Rust avant | Rust optimisé |
|---|---|---|---|---|
| 1 | 7 031 | 2 136 | 6 175 | 3 526 |
| 2 | 7 671 | 2 052 | 5 981 | 3 780 |
| 3 | 7 605 | 2 175 | 6 177 | 3 527 |
| 4 | 7 494 | 2 184 | 6 353 | 3 785 |
| 5 | 7 629 | 2 130 | 6 174 | 3 828 |
| 6 | 7 676 | 2 184 | 6 210 | 3 739 |
| 7 | 7 498 | 2 033 | 5 921 | 3 820 |
| 8 | 7 600 | 2 163 | 6 238 | 3 814 |
| 9 | 7 106 | 2 217 | 6 288 | 3 781 |
| 10 | 7 416 | 2 168 | 5 794 | 3 539 |

Tous ces temps sont en millisecondes. JS et Go sont reconstruits avec les mêmes
sources PureScript finales que Rust ; le quatrième exécutable est le témoin Rust
initial figé. L'[ancienne qualification des hôtes](2026-10-03-gopurs-rust-host.md)
reste la référence initiale, avec 5 782,5 ms pour Rust dans sa propre campagne.

## Confirmation avant/après

| Phase | Rust avant | Rust optimisé |
|---|---|---|
| load TAST + sort | 809 ms | 212 ms |
| prepare + monomorphize | 1 743 ms | 1 122 ms |
| runtime | 7 ms | 7 ms |
| optimize + emit | 3 432 ms | 2 221 ms |
| entry points | 19 ms | 18 ms |
| backend total | 5 886 ms | 3 592 ms |

Moyennes : **6 004,8 → 3 621,2 ms** ; plages
5 640–6 433 / 3 527–3 822 ms.
Écart apparié médian **-2 349 ms**, moyen
**-2 383,6 ms**. Aucun échantillon n'est exclu.
Les médianes de phases ne s'additionnent pas. La spécialisation transitive est
une sous-phase de préparation, déjà comprise dans sa durée.

| Paire | Rust avant | Rust optimisé | Après − avant |
|---|---|---|---|
| 1 | 5 640 | 3 531 | -2 109 |
| 2 | 6 068 | 3 562 | -2 506 |
| 3 | 6 342 | 3 694 | -2 648 |
| 4 | 6 178 | 3 822 | -2 356 |
| 5 | 6 304 | 3 592 | -2 712 |
| 6 | 5 886 | 3 684 | -2 202 |
| 7 | 5 687 | 3 626 | -2 061 |
| 8 | 5 863 | 3 527 | -2 336 |
| 9 | 5 885 | 3 536 | -2 349 |
| 10 | 5 776 | 3 549 | -2 227 |
| 11 | 5 807 | 3 581 | -2 226 |
| 12 | 6 077 | 3 527 | -2 550 |
| 13 | 5 872 | 3 622 | -2 250 |
| 14 | 6 433 | 3 699 | -2 734 |
| 15 | 6 254 | 3 766 | -2 488 |

## Changements retenus

1. **Ordonnanceur PBO** : préserver les modules déjà prêts, respecter les attentes
   implicites et limiter la spéculation aux références finalisées. Publication
   canonique et progression en cas de dépendances inversées restent préservées.
2. **Cache de substitutions Rust** : utiliser l'identité des arbres partagés
   `ExprType`/`BackendSyntax`, avec conservation de leurs propriétaires, plutôt
   que l'enveloppe `Any` temporaire reconstruite à chaque appel.
3. **Directives natives** : petit parseur ASCII pour les cas courants ; toute
   autre syntaxe et les erreurs reviennent au parseur PureScript. Go et JS
   conservent la délégation au parseur PS.
4. **TAST Rust** : décodage natif des tableaux, annotations et modules, puis
   validation des faits d'usage ; partage des types, Unicode et ordre des
   diagnostics restent couverts par les oracles différentiels.
5. **Index AST global de gopurs** : réutiliser l'insertion native des maps de
   chaînes existante, avec les mêmes clés, ordre et valeurs. Son petit effet
   marginal est confirmé sur quinze paires : médianes
   3 747 → 3 596 ms,
   moyennes 3 704,6 → 3 634,9 ms,
   11/15 favorables,
   écart apparié médian -42 ms.

Les 16 exécutions de la confirmation finale (chauffe comprise) font exactement
**238 tentatives PBO / 238 modules, zéro rejet**, contre
**114–121 rejets** avant.
Les profils instrumentés sont conservés à part des chronométrages. Le profil
initial attribuait 2 019 échantillons inclusifs aux substitutions, 456 aux
directives et 312 à la validation d'usage ; le profil du candidat final donne
243 / 14 / 51. Ces familles se recouvrent et ne sont pas des pourcentages CPU.

## Qualification

- Bootstrap Rust de production : **500 modules / 290 647 types** ;
  **1008 fichiers Rust/Cargo identiques** entre Purust JS et natif.
  Les 1009 fichiers, script de liaison compris,
  sont aussi identiques au candidat mesuré. Projet frais Go/FFI exécuté avant
  l'installation atomique ; chemin public `GOPURS_RUST=1 ./bin/test -c` réussi.
- Les trois hôtes passent **45 contrôles Aff et le stress AVar de 1 000 éléments**,
  avec TAST vivant identique et **294 fichiers Go exacts**.
- **37 tests CLI/codegen**, **19 tests préparation/auxiliaires**, **15 suites PBO**
  et **sept suites natives différentielles** réussissent, sans skip. Le parseur
  Go et le contrat de cache Go passent leurs contrôles `-race`.
- Différentiel natif : 169 directives par défaut, 352 cas générés, 33 replis ;
  238 modules natifs, 51 frontières de décodage, 110 909 annotations,
  824 tables / 149 495 entrées ; conservation des propriétaires du cache,
  réentrance, persistance des maps et clés Unicode.
- Vérification indépendante : **208 générations /
  61 152 fichiers Go/manifests** relus et exacts dans
  les campagnes de sélection, confirmation et ressources, chauffes comprises.
  Sorties, temps bruts et empreintes des exécutables sont revérifiés.

Le TAST chronométré demeure l'original ; la suite Aff synchronisée actuelle est
la qualification applicative séparée. Les résultats historiques **JSON → Typed
AST Go/Rust et leurs empreintes structurelles sont conservés** : leurs dix-huit
sorties brutes et les six exécutables archivés sont relus et revérifiés.
Le [rapport JSON/TAST historique](2026-10-03-purust-native-optimization.md#json--typed-ast-gorust)
garde son périmètre ; les nouvelles mesures portent sur la compilation complète
de gopurs-aff.

## Ressources et protocole

Diagnostic séparé `/usr/bin/time -l`, trois paires après chauffe, processus entier :

| Ressource | Avant | Après |
|---|---|---|
| Temps réel | 6,2 s | 3,8 s |
| CPU utilisateur + système | 13,3 s | 7,6 s |
| Pic RSS | 594,4 Mio | 566,6 Mio |

- Corpus : **238 modules / 136 604 types / 26 606 491 octets**, manifeste TAST
  `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
- Mesure principale : `backend total`, chargement/tri, préparation, PBO,
  génération/émission Go, points d'entrée et attente finale des workers inclus.
  Frontend, bootstrap, builds et exécution applicative, démarrage/sortie exclus.
- Jobs chargement/préparation/PBO/émission : **8/8/8/8**, `GOPURS_PIPELINE=1`
  pour tous les hôtes. Lors de cette campagne, le launcher Rust gardait encore
  un défaut PBO séquentiel, testé fonctionnellement lors de la qualification.
  La correction liée en tête de rapport a ensuite aligné ce défaut sur Go.
- Rust : **O3 sans LTO ni debug, Arc et mimalloc**. Go garde le launcher de
  production, `GOGC=off` / `GOMEMLIMIT=10GiB` sur cette machine ; le parseur Go
  embarqué dans Rust garde son GC ordinaire.
- Processus et sorties neufs, caches `.purmeta`/`.cache` supprimés par passage,
  cache de fichiers du système chaud. Builds, profils, tests lourds et mesures
  sérialisés ; commandes longues en arrière-plan avec logs conservés.
- Apple M4 Pro, 14 cœurs logiques, 48 Gio, Node v24.8.0.

## Archives et diagnostics

Archive : `var/benchmark/gopurs-rust-optimization-20261003/`.
[Données vérifiées et empreintes](2026-10-03-gopurs-rust-optimization.json) ;
[harnais et commandes](../../bin/benchmark/gopurs-rust/README.md).
Les exécutions échouées restent archivées : une coupure réseau a interrompu
deux audits délégués ; le premier nouveau test de filtrage des maps omettait le
dictionnaire `Ord` et a été corrigé avant reprise de la qualification.
Les nettoyages de caches Cargo conservent et ré-empreintent les exécutables,
sources générées, logs, mesures et diagnostics.

| Exécutable installé | SHA-256 |
|---|---|
| gopurs.js | `20aedc162a4de8cd6ddb8beeaf7ddad577e38c3fec8aefbcfe29b12bff37d6a6` |
| gopurs-native | `20a537798879f6a19f8716a06986649a207afc189d8ef9bd79bea6f9772dd0d1` |
| gopurs-rust | `f099077b7f6daf4fa0606e3b66919046f5fa205c17c5f3f09449a71e0ff23478` |
