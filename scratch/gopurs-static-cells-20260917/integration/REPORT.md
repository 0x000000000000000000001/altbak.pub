# Validation du vrai générateur — 17 septembre 2026

L’intégration générique produit un gain mesuré sur RBTree. La sonde indépendante
donne −3,53 % pour la médiane des variations des cinq paires ; les médianes des
25 appels regroupés passent de 9,370458 à 9,080833 ms (−3,09 %). La suite complète
confirme une baisse de RBTree de 3,31 % et du total de 2,49 %. Ces résultats sont
une nouvelle campagne ; ils ne sont pas les mesures du prototype parent.

## Artefacts et attribution

- `before/output` copie le Go avant le prototype, sans la garde de recoloration
  abandonnée. `after/output` est régénéré par le vrai bundle Gopurs avec
  `gopurs --main App --ffi ffi`, depuis `after/`.
- Les 301 fichiers CoreFn ont exactement les mêmes SHA. Le staging conserve
  81 sources FFI brutes pour résoudre leurs chemins depuis le dossier isolé.
- Sur 388 fichiers Go avant/après, **seul `purescript/Test_RBTree.go` change**.
  Runtime, types, main, autres modules et FFI générés sont identiques.
  Seules les fonctions `balance_0_consume` et `makeBlack_0_consume` changent ;
  le fichier passe de 5 141 à 3 932 lignes. Voir `generated.patch` et
  `go-diff-summary.json`.
- Bundle générateur SHA-256 :
  `6fa0beb7fed381fa8d37ab2a76265269d4f18313baad3457193da7019dd799c4`.
- La sonde avant est le binaire archivé du dossier parent, recopié sans
  reconstruction. La sonde après et les deux suites compilées utilisent
  `go build -pgo=off`, Go `go1.27.0 darwin/arm64`, le même cache Go.
- Le binaire natif provient de la campagne du 16 septembre, `after-fficc`.
  Son SHA correspond au manifeste archivé. Ses sources Go correspondent à
  celles de `before/output`, hormis les fichiers d’entrée main sélectionnant
  `AppFFICheatcode` au lieu de `App`. Voir `native-provenance.json`.
- `provenance.json`, `final-audit.json` et les fichiers `*-build-info.txt`
  conservent les empreintes et les informations de compilation.

La première invocation depuis le mauvais répertoire a échoué immédiatement,
sans modifier les sorties ; son diagnostic est conservé dans
`failed-generation-wrong-cwd.log`. La génération correcte a réussi.

## Vérification fonctionnelle

Les huit tests Go passent avant et après : quatre rotations, historiques
persistants, sous-arbres conservés et enfant partagé, construction nominale
de 100 000 nœuds, séquences exclusives et comparaison au constructeur persistant,
recoloration avec donneur distinct, insertion avec donneur distinct.

Une seule assertion de la sonde change, identiquement avant/après : la
recoloration peut rendre la cellule d’entrée **ou** le donneur disponible.
Le choix précis d’adresse n’est pas contractuel ; la règle générique préfère
maintenant la cellule connue non nulle. Les assertions Rc, couleurs, valeurs,
enfants, forme, ordre, partage et persistance sont conservées. Les journaux
complets se trouvent dans `before/probe/correctness.log` et
`after/probe/correctness.log` (8/8 chacun).

## Sonde RBTree

Cinq paires de processus, ordre alterné ; trois chauffes puis cinq appels
mesurés par processus, 100 000 insertions depuis `nil`, profondeur 22 validée
pour chaque résultat. Un GC est forcé avant chaque appel, hors chronomètre.
`GOGC=800`, `GOWORK=off`, PGO désactivé, `PPROF` absent ; `GOMAXPROCS` et
`GOMEMLIMIT` non fixés. Aucun build concomitant. Ces durées ne sont pas
interchangeables avec les minima de la suite complète.

| Paire | Avant (ms) | Après (ms) | Variation |
| --- | ---: | ---: | ---: |
| 1 | 10,043834 | 10,077791 | +0,338 % |
| 2 | 9,296125 | 8,943500 | −3,793 % |
| 3 | 9,345833 | 9,015667 | −3,533 % |
| 4 | 9,370458 | 9,046208 | −3,460 % |
| 5 | 9,355125 | 9,017875 | −3,605 % |

Quatre paires sur cinq sont favorables. La première est plus lente des deux
côtés ; aucune cause précise n’est attribuée à cet écart. Les 25 appels par
variante sont tous conservés, sans exclusion. La médiane des variations est
−3,533 %, différente du rapport des médianes regroupées, −3,091 %.

Les allocations sont inchangées : médiane 100 000 allocations / 3 200 000 octets
des deux côtés ; bornes observées 100 000–100 001 et 3 200 000–3 200 016 octets.
Les écarts temporels restent soumis au bruit de mesure ; ce petit échantillon
ne fournit pas un intervalle de confiance. Données : `results.json` et
`measurements/`. Reproduction : `measure.py --pair N`, puis `--summary` dans
un nouveau dossier de campagne (le script refuse d’écraser une paire).

## Suite complète et comparaison native

Trois processus par variante, ordre tournant avant/après/natif, sans build
concurrent. Chaque processus effectue trois chauffes globales, trois par test,
puis le minimum de dix mesures. Les 14 résultats numériques de chacun des
neuf processus sont validés, soit 126 résultats. Même environnement que la
sonde ; comparaison par médiane de chaque ligne, puis somme de ces médianes.

| Mesure | Avant | Après | Go natif manuscrit |
| --- | ---: | ---: | ---: |
| Total, somme des médianes | 12,58267 ms | 12,26888 ms | 10,59596 ms |
| RBTree, médiane | 9,40708 ms | 9,09525 ms | 8,28300 ms |

Le total baisse de 0,31379 ms (−2,494 %), dont 0,31183 ms pour RBTree
(−3,315 %). Les variations des autres lignes sont compatibles avec le bruit
observé ; leurs sources sont identiques. Le compilé reste 15,79 % plus lent
que le natif pour le total de cette campagne.

La baseline officielle d’`altbak.pub/README.md` reste **13,01 ms / 11,09 ms**,
et RBTree **9,74125 ms / 8,75242 ms** (lignes 71 et 65). Les chiffres plus bas
de cette courte campagne ne remplacent pas cette référence : la comparaison
causale porte sur avant/après exécutés ensemble. Aucun README officiel n’est
modifié ici. Données complètes : `suite-results.json`, `suite-measurements/`
et script `measure-suite.py`.
