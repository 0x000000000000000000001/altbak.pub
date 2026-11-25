# Comparaison des 14 cas : Go, JavaScript officiel et C fonctionnel

Campagne du **25 novembre 2026**, terminée à **16 h 29 (Europe/Paris)**.

Les ratios sont les temps Go divisés par la référence : au-dessus de 1×, Go est plus lent. Le tri utilise les valeurs archivées avant arrondi. Les deux classements couvrent 14/14 cas.

## Protocole

Trois processus indépendants par backend, ordre tournant Go/JS/C, JS/C/Go, C/Go/JS. Trois échauffements globaux et locaux, calibration vers 10 ms, dix batches ; chaque durée publiée est la médiane des trois minima par processus. Tous les résultats et tailles de batches ont passé le validateur commun. Les sources et artefacts ont été figés et vérifiés avant et après chaque exécution.

Go : Release, GOGC=800, sans PGO, GOMAXPROCS par défaut. JavaScript : compilateur PureScript officiel, runtime Node.js. C : référence fonctionnelle compilée séparément avant la campagne. Les versions, commandes, profils, hashes, labels/paramètres, oracles et sorties brutes sont conservés dans [l’archive autonome](2026-11-25-go-c-js-functional.json), qui inclut le manifeste de campagne, les neuf sorties brutes et les sources figées.

La machine n’est pas isolée des autres activités. La charge système au début et à la fin de chaque processus est archivée. Les intervalles min–max ci-dessous sont les ratios observés dans les trois cycles tournants, pas des intervalles de confiance. Le classement porte sur les ratios des médianes, qui peuvent différer de la médiane des ratios par cycle.

## Portée de la comparaison

La référence C traduit la structure fonctionnelle des cas. Les backends restent libres d’optimiser : une forme close légitime est un résultat de compilation, pas une erreur. Ces ratios comparent les programmes complets optimisés ; ils ne sont pas des coûts par opération source ni la preuve d’un nombre identique d’opérations machine.

Le C alloue via `malloc` avec un registre et libère les allocations dans chaque action chronométrée ; aucune arène préallouée hors mesure n’est utilisée. Go et JavaScript utilisent leur GC. Les représentations et coûts de libération restent donc différents. Les trois harnais consomment les résultats via des accès atomiques ; les instructions natives et les adaptations FFI Go/JS restent différentes. Les adaptations d’appel et harnais sont incluses ; les très petits temps peuvent être dominés par ces coûts. La correction des oracles ne prouve pas à elle seule l’équivalence structurelle : le [contrat fonctionnel](../../tmp/fp_reference_contract.md) et le [source C](../../tmp/bench_functional.c) doivent être lus avec ces résultats.

## Go / C fonctionnel — ordre décroissant

| Cas | Go (µs) | Référence (µs) | Ratio | Min–max des ratios par cycle |
| --- | ---: | ---: | ---: | ---: |
| Polymorphism | 0.032 | 0.00068 | 47.41× | 45.98–47.95× |
| RowToList | 0.026 | 0.0017 | 15.90× | 15.69–16.08× |
| Tail Call Optimization | 60.19 | 15.07 | 3.99× | 3.95–4.00× |
| Fibonacci | 0.15 | 0.086 | 1.74× | 1.71–1.79× |
| Ackermann | 16.58 | 14.25 | 1.16× | 1.09–1.18× |
| State Monad | 92.85 | 95.59 | 0.97× | 0.91–0.99× |
| Red-Black Tree | 9472.17 | 53161.00 | 0.18× | 0.17–0.19× |
| Church Numerals | 223.43 | 1879.75 | 0.12× | 0.12–0.12× |
| Array Processing | 0.028 | 0.33 | 0.084× | 0.082–0.085× |
| AST Evaluation | 0.051 | 0.73 | 0.071× | 0.070–0.073× |
| Deep Record Updates | 0.047 | 11.49 | 0.0041× | 0.0040–0.0041× |
| Prime Sieve | 0.57 | 158.39 | 0.0036× | 0.0034–0.0036× |
| List Processing | 0.028 | 21.02 | 0.0013× | 0.0013–0.0013× |
| Lazy Evaluation | 0.76 | 20301.00 | 0.000037× | 0.000036–0.000038× |

## Go / JavaScript officiel — ordre décroissant

| Cas | Go (µs) | Référence (µs) | Ratio | Min–max des ratios par cycle |
| --- | ---: | ---: | ---: | ---: |
| State Monad | 92.85 | 46.62 | 1.99× | 1.87–2.07× |
| RowToList | 0.026 | 0.036 | 0.74× | 0.73–0.77× |
| Fibonacci | 0.15 | 0.31 | 0.48× | 0.45–0.48× |
| Tail Call Optimization | 60.19 | 185.99 | 0.32× | 0.32–0.33× |
| Ackermann | 16.58 | 63.84 | 0.26× | 0.25–0.26× |
| AST Evaluation | 0.051 | 0.26 | 0.20× | 0.20–0.20× |
| Church Numerals | 223.43 | 1362.31 | 0.16× | 0.16–0.17× |
| Red-Black Tree | 9472.17 | 82980.62 | 0.11× | 0.11–0.12× |
| Prime Sieve | 0.57 | 58.25 | 0.0097× | 0.0095–0.010× |
| Array Processing | 0.028 | 6.99 | 0.0040× | 0.0039–0.0042× |
| List Processing | 0.028 | 8.70 | 0.0032× | 0.0032–0.0033× |
| Deep Record Updates | 0.047 | 72.88 | 0.00065× | 0.00063–0.00066× |
| Lazy Evaluation | 0.76 | 13813.88 | 0.000055× | 0.000053–0.000056× |
| Polymorphism | 0.032 | 11494.62 | 0.0000028× | 0.0000027–0.0000029× |

## Baselines historiques — comparaison séparée

La colonne C impérative historique et le catalogue restent inchangés. Elle ne remplace pas la nouvelle référence fonctionnelle. Les variations suivantes confrontent les mesures Go/JS fraîches aux baselines archivées du README ; elles ne sont pas une expérience avant/après appariée et peuvent inclure des changements de harnais, versions ou environnement.

| Cas | Go frais / Go archivé | JS frais / JS archivé | C impératif archivé (µs) |
| --- | ---: | ---: | ---: |
| AST Evaluation | 0.96× | 1.00× | 0.090 |
| Fibonacci | 0.97× | 1.00× | 0.077 |
| List Processing | 1.04× | 0.91× | 0.046 |
| Tail Call Optimization | 0.95× | 0.98× | 32.14 |
| Deep Record Updates | 1.02× | 0.98× | 3.25 |
| Ackermann | 1.00× | 0.95× | 16.83 |
| Church Numerals | 0.96× | 1.08× | 0.00068 |
| Prime Sieve | 0.95× | 0.92× | 1.01 |
| Red-Black Tree | 0.94× | 0.92× | 9053.50 |
| Polymorphism | 1.03× | 0.50× | 0.00067 |
| State Monad | 1.02× | 0.97× | 0.00067 |
| Lazy Evaluation | 0.98× | 0.95× | 0.00067 |
| Array Processing | 1.03× | 0.95× | 0.046 |
| RowToList | 1.00× | 0.86× | 0.00067 |

## Validation de la nouvelle référence

- Les 14 oracles et 58 variantes d’entrée passent en C `-O3 -Wall -Wextra -Werror`.
- La même vérification passe sous AddressSanitizer et UndefinedBehaviorSanitizer, avec contrôles Int32, records immuables, arbres persistants et dictionnaires.
- Relecture indépendante des 14 traductions et des durées de vie des captures : aucun défaut bloquant identifié.
- Huit tests du runner passent : oracles/batches, tri avant arrondi, médianes, couverture, empreintes, inventaire JavaScript, manifeste C et rotation des neuf processus.
- Les 126 résultats et 126 batches de la campagne sont validés. Les compteurs d’allocations du mode diagnostic sont absents du binaire C mesuré.
- Les sept tests de précision passent. Quatre tests de publication passent ; le cinquième attend le titre `### Application benchmark results`, absent du README dans `HEAD` comme après les changements. Cette erreur préexistante est distincte de la campagne.

## Lecture des écarts

Polymorphism et RowToList sont devenus des calculs C très courts. Leurs ratios sont particulièrement sensibles aux chemins d’appel et aux harnais ; ils ne représentent pas un coût par dictionnaire. Les très faibles ratios Go sur List, Records, Lazy et d’autres cas incluent les éliminations et transformations permises par le contrat fonctionnel. Aucune moyenne globale n’est présentée comme un score de qualité du compilateur.

State est à 1,99× JS pour le ratio des médianes, avec des ratios par cycle de 1,87 à 2,07× : cette campagne ne confirme donc pas un passage sous 2× dans tous les processus. Les comparaisons aux baselines du README sont conservées ci-dessus, sans attribuer les variations à un changement du compilateur.

## Reproduction

Le [runner](../../bin/benchmark/compare-core.py) prend trois artefacts figés et écrit une nouvelle campagne. Reconstruire Go avec `bin/native/driver.py go --build-only --build-dir <nouveau-dossier>`, puis JS officiel avec `bin/js/driver.py js --build-only`. Le C se compile ainsi :

```sh
clang -std=c11 -O3 -Wall -Wextra -Werror tmp/bench_functional.c -o benchmark-c
./benchmark-c --verify
```

Fournir à côté du binaire C un fichier `benchmark-c.manifest.json` contenant son SHA-256, `source_sha256` pour `tmp/bench_functional.c`, la version clang, la commande exacte et le profil de compilation. Le format et le manifeste utilisé sont conservés dans l’archive. Puis, après la fin de toutes les compilations :

```sh
python3 -B bin/benchmark/compare-core.py \
  --go-dir <artefact-go> --js-dir <artefact-js> \
  --c-binary <benchmark-c> --output <nouvelle-campagne>
```

Le Go du compilateur et les sources PureScript n’ont pas été modifiés pour cette campagne. La nouvelle référence C est distincte de `tmp/bench_native.c`, qui conserve les chiffres historiques.
