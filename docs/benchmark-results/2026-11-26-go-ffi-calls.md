# Go : conversions des arguments FFI — 26 novembre 2026

Les sept petites actions passent de **13–14 ns à 6,5–7,8 ns**, soit
**43 à 51 % de temps en moins** dans la paire avant/après. Le compilateur
réutilise la conversion des littéraux scalaires immuables vers l’interface
polymorphe du FFI. Le résultat de l’appel opaque n’est jamais mis en cache.
Les benchmarks, le harnais, le FFI et les références C et JS sont inchangés.

Les [données et preuves](2026-11-26-go-ffi-calls.json) contiennent les sources
du compilateur avant/après, les sorties brutes, les scripts reproductibles,
les manifestes, les tests et les contrôles d’entrées variables.

## Résultats sur les petites actions

Temps de la paire App à 14 cas, avant → après ; ratios contre les mêmes
médianes C impératives avec entrée opaque :

- Church : **14,140 → 7,314 ns** ; Go/C **12,659 → 6,548×**.
- State : **13,658 → 6,764 ns** ; Go/C **12,201 → 6,042×**.
- Lazy : **13,656 → 6,827 ns** ; Go/C **12,178 → 6,088×**.
- Polymorphism : **13,712 → 7,781 ns** ; Go/C **12,180 → 6,911×**.
- RowToList : **13,359 → 6,519 ns** ; Go/C **11,934 → 5,824×**.
- List Processing : **14,282 → 7,601 ns** ; Go/C **11,733 → 6,244×**.
- Array Processing : **14,369 → 7,633 ns** ; Go/C **10,165 → 5,400×**.

La paire AppAll à 37 cas confirme les gains sur ces sept actions :
**49 à 54 % de temps en moins**. AST et Deep Record Updates bénéficient
également de la conversion évitée, avec respectivement **17,43 %** et
**20,55 %** de temps en moins dans App14.

Après initialisation, le diagnostic sur les getters réellement compilés
trouve **une allocation de moins par action sur chacun des 14 cas**.
Douze passent de 1 à 0 ; Primes passe de 2 à 1 et RBTree de 100001 à 100000.
Ce diagnostic d’allocations est distinct des chronométrages publiés.

## Les quatre indicateurs Go/C

Sur les mêmes **14 cas communs**, la paire fraîche donne :

- Ratio des totaux : **1,206713× → 1,203742×**.
- Moyenne arithmétique μ : **6,805616× → 3,938655×**.
- Médiane Q₂ : **7,818883× → 5,454538×**.
- Moyenne géométrique G : **2,970261× → 2,066376×**.

Le total Go est de **10,569465 → 10,543442 ms**, pour **8,758890 ms en C**.
Les calculs utilisent les temps non arrondis : ΣGo/ΣC, puis moyenne,
médiane et moyenne géométrique des quatorze ratios individuels.

La [baseline officielle précédente](2026-11-25-go-state.md) était
**1,117572×**, **6,399058×**, **7,187313×**, **2,783742×**,
avec un total Go de **9,788691 ms**. Le total remesuré est donc plus élevé
que celui du README précédent, y compris pour l’ancien exécutable.
Les gains attribués à cette optimisation utilisent la paire fraîche.

Le total reste dominé par RBTree. Sa variation explique 25,458 µs des
26,023 µs de baisse du total ; ce mouvement ne démontre pas un gain causé par
la conversion évitée. Les totaux par processus se recouvrent largement :
**10,322–10,707 ms avant**, **10,252–10,631 ms après**.
La réduction du coût des petites actions améliore surtout les trois autres
indicateurs. L’objectif de 1× reste à atteindre.

## Transformation et garanties

Le worker d’un effet étranger immédiatement exécuté accepte maintenant
directement une interface lorsque le TAST décrit un argument polymorphe
`TypeVar` et que le FFI Go attend `any` ou `interface{}`. L’interface transporte
toujours la représentation canonique `Value`, conformément à l’ABI existante.

Pour un littéral scalaire dont la traduction ne produit aucun statement,
le compilateur émet un getter de cache initialisé à la première utilisation.
Il conserve l’expression exacte, y compris les nombres particuliers.
Les arguments dynamiques et structurés continuent d’être convertis à chaque
utilisation ; les types `Foreign` gardent leur adaptation ordinaire.
Les noms internes des caches ne peuvent pas entrer en collision avec un
binding PureScript utilisateur.

Seule une application saturée consommée par un bind immédiatement adjacent
peut utiliser ce worker. Les actions sauvegardées ou rejouées conservent leur
représentation existante. La fonction FFI, l’effet qu’elle retourne, leur ordre
d’évaluation et leur nombre d’exécutions sont conservés.

L’assemblage final conserve les appels à `opaqueValue`. Les calculs dépendant
de l’entrée reçoivent la valeur retournée ; RowToList conserve son résultat
invariant. La conversion allouante apparaît dans l’initialisation du cache,
plus dans son chemin répété. La transformation ne dépend d’aucun nom de
benchmark ni de son résultat attendu.

L’adaptation des callbacks par `Apply`/`Unbox` reste présente. Son remplacement
par un chemin natif typé constitue la prochaine étape distincte du plan.

## Validation

- **47 tests ciblés** passent, dont l’oracle Go d’ordre des arguments,
  de la fonction FFI et de son effet, ainsi que les refus de spécialisation.
- Une fixture indépendante concorde avec JavaScript sur les hôtes du
  compilateur JS et natif, en Release et Dev. Elle couvre effets répétés,
  captures variables, valeurs structurées, callbacks et actions sauvegardées.
- **88 variantes d’entrée par exécutable** concordent avant/après et avec C ;
  chaque action franchit exactement une fois la frontière opaque.
- Les **306 résultats** des douze processus et leurs lots calibrés sont validés.
  Les empreintes des sources, références et exécutables sont contrôlées.
- Les **28 résultats** du diagnostic d’allocations sont corrects.
- Le bootstrap natif et le contrôle des 44 documents et 133 entrées du
  module-map passent. Le seul avertissement du bootstrap est le
  `ShadowedName` préexistant dans `PrivateSTRegions.purs:266`.

La première tentative de fixture JS n’exerçait pas la nouvelle règle : son
bind était éliminé avant l’émission. L’exemple a été rendu discriminant et
rejoué avec succès. Cette tentative et les prototypes de départ sont archivés
séparément ; ils ne fournissent aucun échantillon publié.

La comparaison exhaustive du Go généré explique tous les changements :
46 caches de littéraux dans App14, 57 dans AppAll, la signature du worker
opaque, quatre conversions locales dynamiques dans AppAll et des déplacements
de déclarations ou renumérotations de variables. Après normalisation de ces
changements, les autres déclarations sont identiques. Les sections `Bench.go`,
le runtime, les entrées, l’initialisation, les imports et les directives Go
restent inchangés.

Trois hausses dépassent 5 % dans la paire principale AppAll :

- `st/ref-write-10000` : **6,575480 → 7,649109 µs** (+16,33 %).
- `record/set-large-1000` : **0,400007 → 0,439283 µs** (+9,82 %).
- `string/cp-length-unicode-1000` : **6,944010 → 7,622722 µs** (+9,77 %).

Une contre-mesure isolée conserve les getters compilés et le harnais d’origine,
avec trois processus par variante et par cas. Les 18 résultats sont corrects.
Elle donne respectivement **7,562052 → 6,463094 µs** (−14,53 %),
**0,456258 → 0,455802 µs** (−0,10 %) et
**6,829366 → 6,792582 µs** (−0,54 %).
Aucune des hausses n’est reproduite dans ce contrôle. Le calcul mesuré de ces
trois actions est inchangé ; seuls les arguments préparatoires des deux
premières changent de conversion. Ces observations ne prouvent pas l’absence
de toute régression. Les échantillons et agrégats principaux restent intacts,
y compris les hausses ; les diagnostics ne les remplacent pas.

## Protocole et reproduction

Apple M4 Pro, 48 Gio, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Compilateur natif reconstruit depuis `f28f181` avec les deux modules modifiés
archivés ; TAST produit par le fork PureScript `41fb570`.
Release, GOGC=800, GOMAXPROCS par défaut, sans PGO. Les sorties Go sont
reconstruites dans des répertoires distincts avant la collecte.

Trois processus par variante, en ordre AB/BA/AB, pour App puis AppAll.
Le harnais conserve échauffements, calibration et minimum de dix batches ;
la statistique publiée est la médiane des trois minima de processus.
La compilation, le démarrage et la sortie du processus sont
hors chronomètre. Builds, tests et sondes d’allocations sont terminés avant
la collecte. La machine n’est pas isolée des activités utilisateur.

Le README reprend les 14 temps primaires App et les 23 temps supplémentaires
AppAll dans ses **37 lignes Go WIP**. Go/C porte sur 14 cas, Go/JS sur 37.
Les autres colonnes, les statuts et le Backlog non mesuré sont conservés.
Aucun tableau n’est ajouté au README. Le catalogue supprimé par l’utilisateur
n’est pas recréé : les références JS exactes proviennent de l’archive State.

Depuis `var/benchmark/go-ffi-calls-20261126`, une nouvelle collecte des mêmes
binaires figés peut être lancée dans un répertoire inexistant :

```sh
python3 -B proofs/campaign-scripts/measure-pair.py \
  --after-core after-core --after-all after-all --output replay --execute
```

Les scripts de collecte, de publication et leurs audits sont conservés dans
le workspace et intégrés à l’archive JSON liée plus haut.
