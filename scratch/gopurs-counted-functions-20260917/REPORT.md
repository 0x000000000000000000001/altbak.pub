# Gopurs — fusion des producteurs de fonctions comptées

Campagne du 17 septembre 2026. Le vrai générateur réduit d’environ **52 % le
temps de Church** dans la sonde isolée. Le harnais complet confirme la baisse
de cette ligne. **Une baisse stable du total n’est pas établie** dans les deux
séries complètes ; toutes les mesures sont conservées.

## Transformation et attribution

La passe reconnaît un producteur récursif typé dont la base est l’identité et
l’étape `let previous = build (n - 1) in \f x -> f (previous f x)`.
Pour un compteur non négatif, elle retourne une closure capturant le compteur
et appelant un worker récursif terminal. La TCO existante émet la boucle.
Le corps négatif reste en place. Les callbacks et les compositions sont conservés.

Le prototype modifie seulement le Go de `fromInt`. L’intégration est une passe
structurelle sans nom de benchmark : `Gopurs.FunctionFusion`, après ThunkFusion
et avant Ownership/TCO. Elle conserve les `Apply` du callback ; aucune formule
`n⁵` n’est introduite. Le natif manuscrit utilise cette formule, puis une boucle,
ce qui explique pourquoi sa durée n’est pas une promesse pour cette passe.

La référence avant est le Go validé après la précédente sélection statique des
cellules, archivé dans `gopurs-static-cells-20260917/integration/after/output`.
Les **301 CoreFn ont des empreintes identiques** avant/après. Sur les sorties Go,
seul **`purescript/Test_Church.go`** change : le producteur `fromInt` et le nouveau
worker, avec son getter. Les compositions, le runtime, les FFI et les autres
modules restent identiques. `generated.patch` et `generated-manifest.json`
conservent la différence et les empreintes ; les fichiers de sonde/test ajoutés
sont identiques dans les variantes.

## Vérification

- Build et bundle sans erreur ni avertissement.
- 68 tests outils existants et 29 nouveaux tests passent.
- Les nouveaux tests couvrent les formes acceptées/refusées, les annotations,
  la conservation du corps négatif, les collisions et l’immutabilité de l’IR.
  Deux programmes Go vérifient zéro, callbacks non linéaires et leur trace,
  fonctions sauvegardées, applications partielles, réutilisations, composition
  et exception au troisième callback suivie d’une nouvelle invocation.
- La fixture TAST `CountedFunctions` s’exécute avec l’ancien puis le nouveau
  générateur, avec des noms distincts du benchmark. Son cas capturant le compteur
  reste non transformé. Le nouveau snapshot est ensuite contrôlé sans réécriture.
  La fixture `ThunkFusion` conserve son snapshot et s’exécute aussi.
- Trois tests Go supplémentaires sur le module réel passent dans les variantes
  avant, prototype et intégrée : callbacks/rejeu/partiels, compositions pour
  `n = 0..10` et exception au troisième callback.

Le premier lancement sandboxé de la fixture n’avait pas accès au cache SQLite
de Spago. La relance avec accès au cache a réussi. Avant création du nouveau
snapshot, le runner s’arrêtait normalement sur son absence ; le Go de cette
première génération a été compilé/exécuté séparément et a affiché `Done`.
Les journaux conservés se trouvent dans `validation/`. Les autres fixtures de
compilation ne sont pas incluses dans cette vérification ciblée.

## Sonde isolée

Go 1.27.0 darwin/arm64, `GOGC=800`, `GOWORK=off`, PGO désactivé. Dix chauffes
par processus, puis sept lots de 100 calculs ; GC préalable à chaque lot, hors
chronomètre. Chaque résultat vaut 100 000. Cinq paires de processus en ordre
alterné, sans compilation concomitante. Les allocations sont mesurées sur les
mêmes lots via `runtime.MemStats`. `GOMAXPROCS` et `GOMEMLIMIT` ne sont pas fixés.

| Variante | Médiane des temps par calcul | Allocations | Octets |
| --- | ---: | ---: | ---: |
| Avant, campagne prototype | 487,565 µs | 157 | 6 800 |
| Prototype Go | 231,195 µs | 112 | 5 280 |
| Avant, campagne intégration | 490,800 µs | 157 | 6 800 |
| Vrai générateur | 236,358 µs | 112 | 5 280 |

Le prototype donne une variation appariée médiane de **−52,39 %**.
Pour le vrai générateur : **−51,96 %** apparié, **−51,84 %** sur le rapport
des médianes regroupées. Les cinq paires sont favorables dans chaque campagne.
Les variations des paires intégrées sont −51,96 %, −58,88 %, −51,35 %, −51,55 %
et −52,14 %. La deuxième paire a un avant plus lent ; elle reste incluse.

Les deux séries ont leurs propres mesures avant : elles ne sont pas mélangées.
Le worker réel utilise un accumulateur `int64`, tandis que le prototype conserve
une `Value`. Seuls les résultats du vrai générateur établissent le gain intégré.
Données : `results-prototype.json`, `results-integrated.json`, les dossiers
`measurements-*`, et `measure.py` (refuse d’écraser un dossier de mesures).

## Suite complète

Le harnais existant effectue trois chauffes globales, trois par test et retient
le minimum de dix mesures. La première campagne emploie trois processus par
variante, en ordre tournant avant/intégré/natif. Les 14 résultats numériques
de chacun des neuf processus sont validés, soit **126 résultats**.

| Première campagne : médianes par ligne | Avant | Intégré | Manuscrit |
| --- | ---: | ---: | ---: |
| Church | 486,04 µs | 233,92 µs | 25,83 µs |
| Total, somme des médianes | 13,50499 ms | 13,55479 ms | 12,00891 ms |

Le total augmente de 0,04980 ms malgré les 0,25212 ms gagnées sur Church.
RBTree passe de 9,97333 à 10,20483 ms et Polymorphism de 2,45575 à 2,51392 ms,
bien que leurs sources soient identiques. La cause de ces écarts n’est pas
établie par la seule comparaison des sources.

Une seconde série, décidée pour clarifier cette incertitude, utilise cinq paires
alternées avant/intégré : **140 résultats numériques supplémentaires validés**.

| Seconde campagne : médianes par ligne | Avant | Intégré |
| --- | ---: | ---: |
| Church | 480,54 µs | 236,29 µs |
| Total, somme des médianes | 13,58096 ms | 13,50085 ms |

La baisse du total y est de 0,08011 ms, mais seulement trois paires sur cinq sont
favorables. Les variations des totaux par processus vont de **−3,60 % à +3,61 %**
(médiane appariée −0,80 %). Ces fluctuations empêchent de conclure à une baisse
stable du total. Le résultat robuste est la réduction de Church, observée dans
les sondes et les deux séries complètes. Aucun résultat n’a été écarté.

Données et protocoles : `suite-results.json`, `suite-pair-results.json`, dossiers
`suite-measurements/` et `suite-pair-measurements/`, scripts `measure-suite.py`
et `measure-suite-pairs.py`. Le natif est le binaire archivé déjà vérifié lors de
la campagne précédente ; sa provenance et son SHA sont conservés dans `native/`.

## Baseline et reproduction

La référence officielle reste le [README](/Users/0x1/Documents/htdocs/altbak.pub/README.md:55) :
**Church 499,42 / 29,21 µs**, **total 12,26 / 11,09 ms**, compilé/manuscrit.
Les campagnes ci-dessus ont un environnement temporel différent et ne remplacent
pas cette référence. Aucun résultat du README n’a été réécrit.

La génération intégrée utilise `gopurs --main App --ffi ffi` depuis `integrated/`,
avec 81 sources FFI brutes archivées. Les probes sont compilées depuis chaque
`output/` avec `go build -pgo=off -o ../probe ./church-probe`, les suites avec
`go build -pgo=off -o ../benchmark ./main/main.go`. Le cache Go utilisé est
`/private/tmp/gopurs-adt-reuse-gocache`, avec `GOWORK=off` et `GOPROXY=off`.
Les bundles avant/après, versions, paramètres de build et empreintes binaires
sont conservés dans `provenance.json`, `final-provenance.json` et les fichiers
`*-build-info.txt`. Les scripts de mesure doivent être exécutés dans une nouvelle
campagne, car leurs sorties existantes sont protégées contre l’écrasement.

Les entrées compilateur figées, les packages Go générés, les bundles et les
exécutables restent disponibles localement et sont exclus de Git par le
`.gitignore` de cette campagne. Le rapport, les scripts, le code des sondes/tests
(`probe-sources/`), les patches, empreintes et mesures restent versionnables.
