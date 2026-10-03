# Gopurs Rust — campagne des allocations

Campagne du **3 octobre 2026** sur les **238 modules Aff figés**. La composition
retenue (insertion native + comparaisons `Qualified` empruntées) est reconstruite
par les chemins de production et qualifiée. La confirmation finale mesure
**3 623 → 3 353 ms** (réduction de
7,5 %), avec **15/15 paires favorables**.
Profil final : **O3, ThinLTO, sans debug, Arc et mimalloc**, linker Mach-O `ld64.lld` via `-fuse-ld=/Users/0x1/.rustup/toolchains/stable-aarch64-apple-darwin/lib/rustlib/aarch64-apple-darwin/bin/gcc-ld/ld64.lld`.

**Repères distincts.** L'optimisation algorithmique historique du 3 octobre a
réduit le backend Rust de 5 886 à
3 592 ms (−39 %) ;
la correction des valeurs par défaut du launcher, mesurée sur son propre témoin, donne
5 772 à 3 701 ms
(−35,9 %). Cette passe ne revendique
aucune addition à ces repères : elle sélectionne une composition de changements
d'allocation et en mesure l'effet propre. 5 fichiers externes archivés (`source-integration.json`) sont en plus inclus dans le candidat installé ; le delta net baseline → production n'est pas attribué aux seules allocations, et les sélections natives isolées restent mesurées sur l'ancien snapshot. Les trois médianes finales restent au-dessus de 3 000 ms ; la plus basse est de 3 353 ms.

## Comparaison commune : trois hôtes, une cible Go

| Hôte de gopurs | Médiane | Moyenne | Min–max |
|---|---|---|---|
| JavaScript / Node | 7 364 ms | 7 339,1 ms | 6 901–7 684 ms |
| Go natif | 1 997 ms | 2 022,1 ms | 1 961–2 120 ms |
| Rust avant | 3 572 ms | 3 604,2 ms | 3 456–3 778 ms |
| Rust final | 3 414 ms | 3 373,2 ms | 3 193–3 561 ms |

Dix tours, une chauffe par variante, ordre tournant. Dans cette même campagne,
le rapport Rust/Go passe de **1.789× à 1.710×** ;
**10 % de l'écart absolu est résorbé**. Rust prend encore
71 % de plus que Go et 53,6 %
de moins que JavaScript.

| Tour | JS | Go | Rust avant | Rust final |
|---|---|---|---|---|
| 1 | 7 307 | 2 057 | 3 625 | 3 397 |
| 2 | 7 134 | 2 024 | 3 456 | 3 445 |
| 3 | 7 313 | 2 115 | 3 511 | 3 193 |
| 4 | 7 449 | 1 967 | 3 477 | 3 211 |
| 5 | 7 053 | 2 076 | 3 766 | 3 431 |
| 6 | 7 609 | 1 967 | 3 696 | 3 561 |
| 7 | 7 684 | 2 120 | 3 778 | 3 527 |
| 8 | 7 526 | 1 961 | 3 748 | 3 512 |
| 9 | 7 415 | 1 964 | 3 466 | 3 250 |
| 10 | 6 901 | 1 970 | 3 519 | 3 205 |

## Confirmation avant/après

| Phase | Rust avant | Rust final |
|---|---|---|
| load TAST + sort | 212 ms | 206 ms |
| transitive specializations | 484 ms | 453 ms |
| prepare + monomorphize | 1 107 ms | 1 002 ms |
| runtime | 7 ms | 7 ms |
| optimize + emit | 2 283 ms | 2 142 ms |
| entry points | 18 ms | 17 ms |
| backend total | 3 623 ms | 3 353 ms |

Moyennes : **3 601,7 → 3 347,9 ms** ; plages
3 474–3 774 / 3 234–3 489 ms.
Écart apparié médian **-247 ms**, moyen
**-253,7 ms**. Aucun échantillon n'est exclu.
Les médianes de phases ne s'additionnent pas ; la spécialisation transitive est
une sous-phase de préparation.

| Paire | Rust avant | Rust final | Après − avant |
|---|---|---|---|
| 1 | 3 588 | 3 240 | -348 |
| 2 | 3 730 | 3 276 | -454 |
| 3 | 3 554 | 3 353 | -201 |
| 4 | 3 625 | 3 397 | -228 |
| 5 | 3 774 | 3 470 | -304 |
| 6 | 3 672 | 3 489 | -183 |
| 7 | 3 650 | 3 345 | -305 |
| 8 | 3 637 | 3 390 | -247 |
| 9 | 3 671 | 3 300 | -371 |
| 10 | 3 523 | 3 392 | -131 |
| 11 | 3 527 | 3 441 | -86 |
| 12 | 3 486 | 3 235 | -251 |
| 13 | 3 474 | 3 235 | -239 |
| 14 | 3 491 | 3 234 | -257 |
| 15 | 3 623 | 3 422 | -201 |

## Valeurs par défaut du launcher

Cinq tours tournants sans aucun worker imposé ; seuls les sélecteurs d'hôte sont
propres aux variantes. Le launcher corrigé choisit préparation 8, PBO 8,
émission 8 et pipeline actif.

| Variante | Médiane | Moyenne | Min–max |
|---|---|---|---|
| Rust avant | 3 760 ms | 3 726,2 ms | 3 515–3 832 ms |
| Rust final | 3 394 ms | 3 343,2 ms | 3 115–3 486 ms |
| Go | 2 064 ms | 2 059,2 ms | 1 975–2 148 ms |
| JavaScript | 7 265 ms | 7 255,8 ms | 6 911–7 555 ms |

Le Rust final passe de **3 760 à
3 394 ms** (réduction de 9,7 %),
**5/5 paires favorables**.

| Tour | Rust avant | Rust final | Go | JS |
|---|---|---|---|---|
| 1 | 3 782 | 3 486 | 2 114 | 7 265 |
| 2 | 3 832 | 3 115 | 1 995 | 7 535 |
| 3 | 3 742 | 3 467 | 2 148 | 7 555 |
| 4 | 3 515 | 3 254 | 1 975 | 7 013 |
| 5 | 3 760 | 3 394 | 2 064 | 6 911 |

## Ressources

Diagnostic séparé `/usr/bin/time -l`, trois paires après chauffe, processus entier :

| Ressource | Avant | Final |
|---|---|---|
| Temps réel | 3,8 s | 3,5 s |
| CPU utilisateur + système | 7,6 s | 7 s |
| Pic RSS | 559,2 Mio | 554,9 Mio |

## Sélection et qualification

Les relectures natives et sémantiques couvrent **8 suites
natives différentielles** (396 900 paires de comparaison
`Qualified`) et **15 suites PBO**, plus les adaptateurs Go/JS
`Qualified` : tout passe.

**Suppression rejetée pour absence de gain.** Les deux premières assertions
strictes de balance échouaient aussi sur l'oracle (les jointures générées
peuvent produire un écart de hauteur de 2) ; le contrat final renforcé — forme
exacte contre l'oracle, modèle clés/valeurs et persistance — a passé
**25200 opérations** supplémentaires, et les
15 suites PBO du candidat passent. Les
médianes face au témoin reconstruit
(3 783 / 3 819 /
3 828 ms) ne montrent aucun gain : la
suppression n'est pas retenue et n'est pas présentée comme une erreur de
correction.

**Insertion native retenue.** Sélection provisoire puis confirmation :

| Campagne | Variante | Médiane | Moyenne | Min–max |
|---|---|---|---|---|
| sélection, 5 tours | baseline | 3 672 ms | 3 682 ms | 3 662–3 714 ms |
| sélection, 5 tours | native-insert | 3 641 ms | 3 646,4 ms | 3 625–3 679 ms |
| confirmation, 15 tours | baseline | 3 907 ms | 3 968 ms | 3 724–4 279 ms |
| confirmation, 15 tours | native-insert | 3 811 ms | 3 864,3 ms | 3 659–4 234 ms |

La confirmation passe de 3 907 à
3 811 ms avec
**11/15 paires favorables** (écart apparié médian
-107 ms).

**Composition qualifiée.** Le candidat `qualified` compose l'insertion native
et les comparaisons empruntées `Qualified Ident` ; il bat `native-insert` dans
les **5/5 tours** de la sélection
(écart apparié médian -53 ms).

| Variante | Médiane | Moyenne | Min–max |
|---|---|---|---|
| baseline | 3 690 ms | 3 710,4 ms | 3 639–3 809 ms |
| native-insert | 3 682 ms | 3 711,8 ms | 3 615–3 836 ms |
| qualified | 3 629 ms | 3 660,2 ms | 3 568–3 770 ms |

**Intégration externe.** 5 fichiers divergent du snapshot `qualified` : corrections concurrentes de labels Unicode (GoAst, GoConversions, Printer, RecordExprs) et générateur JSON (CoreFn/Json/Text.go). Ils sont préservés dans le candidat installé, et la campagne `integration-runs` compare les candidats figés à l'intégration sur les mêmes entrées. Le delta net baseline → production inclut donc ces changements archivés ; les sélections natives ci-dessus restent mesurées sur l'ancien snapshot.

**ThinLTO.** Le premier essai `thin-lto` a échoué par disque plein (« No space left on device »). Le retry `thin-lto-retry` a échoué à la liaison : le bitcode Rust LLVM 22 n'est pas lisible par le linker Apple LLVM 17. Le candidat `thin-lto-rust-lld` passe avec le linker Mach-O `ld64.lld` fourni par la toolchain Rust (`-fuse-ld=/Users/0x1/.rustup/toolchains/stable-aarch64-apple-darwin/lib/rustlib/aarch64-apple-darwin/bin/gcc-ld/ld64.lld`, SHA-256 `695f239b52eef3c6551fdda84147710e54eeeed066cd545dd85291ce7545eb30`). Comparaison à cinq tours contre baseline et qualified :

| Variante | Médiane | Moyenne | Min–max |
|---|---|---|---|
| baseline | 3 767 ms | 3 784,4 ms | 3 727–3 838 ms |
| qualified | 3 675 ms | 3 655 ms | 3 539–3 708 ms |
| thin-lto | 3 521 ms | 3 518,8 ms | 3 479–3 569 ms |

Confirmation du profil sur les sources intégrées : **3 662 → 3 495 ms (−4,6 %)**, **14/15 paires favorables**. Moyennes **3 662,5 → 3 491,2 ms** ; les 32 générations, chauffes comprises, sont exactes.

## Profils et limites

Les profils instrumentés restent **séparés des chronométrages** ; les familles
inclusives se recouvrent et ne sont pas des pourcentages CPU.
L'inlining de ThinLTO modifie aussi la visibilité des fonctions dans les piles.

| Famille inclusive | Avant (tout) | Final (tout) | Avant (optimize) | Final (optimize) |
|---|---|---|---|---|
| PBO | 6 234 | 4 914 | 3 576 | 3 060 |
| allocation | 1 487 | 1 136 | 873 | 719 |
| Value clone | 905 | 614 | 546 | 363 |
| String clone | 671 | 212 | 292 | 103 |
| Map/Set | 1 868 | 1 436 | 989 | 913 |
| type substitution | 360 | 260 | 260 | 247 |
| directives | 13 | 11 | 1 | 0 |
| usage validation | 69 | 62 | 16 | 22 |
| decoding | 101 | 48 | 2 | 3 |
| gopurs | 3 323 | 2 784 | 1 613 | 1 292 |

Le profil final porte l'empreinte de l'exécutable de production
(`24ed33ad2624f8442080c2c0a256bae7b5e041ab37cfb6f86253c1ffb0152f3a`) ; celui du témoin, `cf7a066084a8a154312fed6e7df6c9deb3a08f0588d0c9c3e7c1bf9f51d62615`.
Objectif < 3 000 ms : Les trois médianes finales restent au-dessus de 3 000 ms ; la plus basse est de 3 353 ms.

## Protocole et empreintes

- Corpus : **238 modules / 136 604 types /
  26 606 491 octets**, manifeste TAST
  `6a30fb919f104df813feb7f6a884fca0a7f768e44a99f1129884d8992b832149`.
- Machine : **Apple M4 Pro**, 14 cœurs logiques,
  48 Gio, Node v24.8.0.
- Jobs chargement/préparation/PBO/émission : **8/8/8/8**, `GOPURS_PIPELINE=1`
  pour les campagnes principales ; les tours « défaut » n'imposent rien et le
  launcher choisit seul ses limites.
- Temps `backend total` : chargement/tri, préparation, PBO, génération/émission
  Go, points d'entrée et attente finale des workers. Frontend, construction du
  compilateur et de l'application, exécution applicative et démarrage/sortie du
  processus sont exclus. Cache de fichiers système chaud.
- Sorties neuves et caches `.purmeta`/`.cache` supprimés à chaque passage ;
  **294 fichiers Go/manifests exacts** dans chaque génération, chauffes comprises.
- Vérification indépendante : **262 générations /
  77 028 fichiers** relus et exacts dans
  11 campagnes `*-runs` (sélection, confirmation,
  qualification, finale, ThinLTO, intégration), chauffes comprises.
- Tests : **43 tests compilateur**, **19 tests
  préparation/auxiliaires**, **15 suites PBO**,
  **8 suites natives**, zéro échec et zéro skip ;
  parseur Go `-race`, contrat de cache Go et rebuild Rust/Aff publics passés.
- Les trois hôtes passent **45 contrôles Aff + stress AVar
  1 000 éléments**, sur le même TAST vivant et avec 294 fichiers Go identiques.
- Production : candidat `integrated-thin-lto`,
  **1009 fichiers Rust/Cargo** (dont
  `build.rs`, soit 1008 hors
  `build.rs`) identiques entre candidat et production, bootstrap Purust JS/natif
  identique (1008 fichiers), sources gopurs/PBO
  figées, profil `{"opt_level":3,"debug":false,"lto":"thin","threaded":true,"allocator":"mimalloc"}`, linker `/Users/0x1/.rustup/toolchains/stable-aarch64-apple-darwin/lib/rustlib/aarch64-apple-darwin/bin/gcc-ld/ld64.lld` (cc, `-fuse-ld=/Users/0x1/.rustup/toolchains/stable-aarch64-apple-darwin/lib/rustlib/aarch64-apple-darwin/bin/gcc-ld/ld64.lld`).
- Nettoyage Purust : audit `passed`, 9 cibles Cargo
  régénérables supprimées, 37,4 Gio
  libérés, 6186 fichiers protégés revérifiés par
  empreinte (sources et bins Purust inchangés).
- JSON → Typed AST historique conservé : 18 sorties brutes et 6 exécutables relus
  et revérifiés, portée et empreintes structurelles intactes.

## Archives et diagnostics

Archive : `var/benchmark/gopurs-rust-allocation-20261003/` ;
[données vérifiées et empreintes](2026-10-03-gopurs-rust-allocation.json) ;
[harnais et commandes](../../bin/benchmark/gopurs-rust/README.md).
Les stdout/stderr des commandes de qualification sont référencés dans le JSON
(`helpers`) et leurs empreintes figurent dans `source_records`.
Les essais rejetés restent archivés : suppression native non retenue (absence
de gain, contrat renforcé passé), premier build ThinLTO interrompu par un disque
plein, second par le linker Apple LLVM 17 face au bitcode LLVM 22.
L'intégration externe du snapshot `qualified`
est conservée et mesurée séparément.

| Exécutable installé | SHA-256 |
|---|---|
| gopurs.js | `0654f8f6eb8e853d8f87ae94eea64484bc0a3201c13dcf1c28932c9073624d14` |
| gopurs-native | `d99ada8d641c4a68522f483e9bd15c8444d99b3f81c2b43971c1b903b6e367e3` |
| gopurs-rust | `24ed33ad2624f8442080c2c0a256bae7b5e041ab37cfb6f86253c1ffb0152f3a` |
