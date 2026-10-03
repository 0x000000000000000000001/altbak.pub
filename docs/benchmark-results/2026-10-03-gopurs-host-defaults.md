# Gopurs : parallélisme par défaut des hôtes

Correction du **3 octobre 2026**, après le signalement de **2 251 ms pour Go /
6 030 ms pour Rust** avec les commandes publiques. Sur cette machine de 48 Gio,
`GOPURS_RUST=1 ./bin/test` sélectionne désormais automatiquement **préparation 8,
PBO 8, émission 8 et pipeline actif**. Son exécution complète a passé les tests
Aff avec un backend à **3 683 ms** ; le chemin
de reconstruction `GOPURS_RUST=1 ./bin/test -c` passe également.

## Cause et correction

La branche Rust faisait `exec` avant la configuration automatique du launcher.
Elle conservait les défauts internes **préparation 2 / PBO 1**, alors que Go
recevait **8 / 8**. La [campagne d'optimisation précédente](2026-10-03-gopurs-rust-optimization.md)
forçait les workers explicitement ; sa validation de la commande ordinaire
était fonctionnelle. Ses chiffres ne décrivaient donc pas l'ancien défaut Rust.

Le launcher configure maintenant le parallélisme avant la sélection Go/JS/Rust.
Le seuil reste **32 Gio** et les variables explicites gardent la priorité.
Cette décision est indépendante de `GOGC`/`GOMEMLIMIT`. La politique de GC
automatique appartient uniquement au compilateur Go ; Rust conserve notamment
le GC ordinaire de son parseur Go embarqué.

Chaque compilation indique les limites effectives :

`[gopurs] workers: prepare=8, pbo=8, emit=8, pipeline=true`

Les machines sous le seuil conservent les défauts internes, et les appels directs
aux exécutables contournent la politique du launcher. Préparation est bornée à
1–8, PBO/émission à 1–64 ; les valeurs affichées correspondent aux limites utilisées.

## Confirmation sur le TAST figé

Une chauffe puis **cinq tours tournants**, processus sérialisés, sur les mêmes
**238 modules / 136 604 types** que la campagne précédente. Le harnais retire
les variables héritées et n'impose aucun worker aux variantes « défaut ».
Seule la variante de contrôle explicite fixe 8/8/8/8 et le pipeline.

| Variante | Médiane | Moyenne | Min–max |
|---|---:|---:|---:|
| Rust — ancien défaut | 5 772 ms | 5 665 ms | 5 181–5 975 ms |
| Rust — défaut corrigé | 3 701 ms | 3 749,2 ms | 3 624–3 898 ms |
| Rust — huit workers explicites | 3 784 ms | 3 754,4 ms | 3 574–3 913 ms |
| Go — défaut | 2 194 ms | 2 220,4 ms | 2 100–2 402 ms |
| JavaScript — défaut | 7 542 ms | 7 478,6 ms | 7 268–7 641 ms |

Le défaut Rust passe de **5 772 à 3 701 ms
(−35,9 %)**, avec **5/5 paires favorables**.
Cette comparaison mesure l'effet des réglages automatiques sur gopurs déjà optimisé.
Le chargement reste à huit pour Go/Rust et à un pour JS par défaut.

| Tour | Rust ancien défaut | Rust défaut corrigé | Rust explicite 8 | Go défaut | JS défaut |
|---|---:|---:|---:|---:|---:|
| 1 | 5 181 | 3 624 | 3 574 | 2 194 | 7 387 |
| 2 | 5 772 | 3 698 | 3 714 | 2 100 | 7 268 |
| 3 | 5 975 | 3 825 | 3 787 | 2 242 | 7 555 |
| 4 | 5 892 | 3 898 | 3 913 | 2 402 | 7 542 |
| 5 | 5 505 | 3 701 | 3 784 | 2 164 | 7 641 |

Temps `backend total` uniquement : frontend, construction du compilateur,
build/exécution applicative et démarrage/sortie exclus. Sorties et caches neufs
à chaque passage, cache de fichiers du système chaud. Aucun échantillon exclu.

## Validation

- **43 tests launcher/CLI/codegen réussis, zéro skip**, dont six contrats du
  launcher. Trois de ces contrats échouent sur sa version précédente ; les
  diagnostics sont conservés. Les limites séquentielles/parallèles affichées
  sont aussi vérifiées sur les vrais compilateurs.
- JS, Go et Rust reconstruits ; bootstrap de **500 modules /
  290 702 types**, **1008 fichiers Rust/Cargo
  identiques** entre Purust JS et natif. Smoke Go/FFI frais et parseur Go `-race` réussis.
- Commandes Aff publiques des trois hôtes : **45 contrôles + stress AVar
  de 1 000 éléments**, **294 fichiers Go identiques** entre hôtes et à la qualification
  précédente. Paramètres automatiques vérifiés dans leurs logs.
- Confirmation figée : **30 générations / 8 820 fichiers exacts**, chauffes
  comprises, relus indépendamment avec les temps bruts et les exécutables.
- Les mesures historiques à workers explicites et les résultats JSON → Typed
  AST conservent leurs données et leur périmètre.

## Preuves

Archive : `var/benchmark/gopurs-host-defaults-20261003/` ;
[données vérifiées](2026-10-03-gopurs-host-defaults.json). Harnais :
`bin/benchmark/gopurs-rust/qualify-defaults.mjs`, `compare.mjs` avec
`launcherDefaults: true`, `verify.mjs` et `publish-defaults.mjs`.

| Exécutable installé | SHA-256 |
|---|---|
| gopurs.js | `2efaee7a40cc0df12ce17bdb821bd5c6f531093524846517fa1529fc29c3f673` |
| gopurs-native | `d0c2c63b563e589176fe181da4c3664234a2f60a1719a536927f6a7d710a3b02` |
| gopurs-rust | `cf7a066084a8a154312fed6e7df6c9deb3a08f0588d0c9c3e7c1bf9f51d62615` |
