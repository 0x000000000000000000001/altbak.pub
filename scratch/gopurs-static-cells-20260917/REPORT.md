# Rotation Go : sélection statique de trois cellules

Expérience isolée du 17 septembre 2026. **Signal modestement favorable** :
variation appariée médiane −2,137 % sur cinq paires, mais un passage défavorable
et une forte variation dans la première paire. Le résultat ne constitue pas
une preuve de gain stable de 12 %. Aucun fichier du backend n'a été modifié.

## Changement et périmètre

Le témoin est copié depuis `gopurs-field-updates-20260917-prototype/before/output` :
il ne contient pas les gardes de recoloration précédemment rejetées. Les deux
variantes conservent les mêmes 1 989 fichiers d'entrée, sauf `Test_RBTree.go`.
Son SHA initial est
`b93c7cf9a01100043f900e975de5ec4ab94f69a34515be1f1214e2ee390a6527`.

Seule `Call_Test_RBTree___gopurs_owned_balance_0_consume` est modifiée.
**21 segments** construisent trois cellules depuis un donneur et deux cellules
mortes non nulles. Au lieu de parcourir et vider le même pool trois fois, ils
prennent directement la première cellule morte pour le premier constructeur,
la seconde pour le deuxième et le donneur pour le troisième. Une allocation
subsiste uniquement si ce donneur est nul. Les **19 segments à une cellule**
restent inchangés, ainsi que les autres fonctions.

Pour chaque segment, `segments.json` conserve les chemins, noms de cellules,
ligne initiale et lectures directes établissant la non-nullité des deux cellules
mortes dans les snapshots immédiatement précédents. Tous les snapshots restent
avant toute mutation. Les 15 écritures Rc/V0/V1/V2/V3 des trois constructeurs
sont textuellement identiques avant/après. La preuve d'absence de partage est
celle des workers consommateurs existants ; ce prototype n'élargit pas leurs
points d'entrée. L'ordre de réemploi des cellules change, pas la structure
logique des rotations. Le fichier Go passe de 5 141 à 3 881 lignes.

## Validation et protocole

Sondes préparées par le `stage-harness.py` archivé le 16 septembre : même
`main.go`, mêmes adaptateurs et même construction depuis `nil`, au point
d'entrée consommant réellement choisi par le benchmark. Go 1.27.0 darwin/arm64,
`GOGC=800`, `GOWORK=off`, PGO désactivé ; GOMAXPROCS et GOMEMLIMIT non fixés.
Cache de compilation existant : `/private/tmp/gopurs-adt-reuse-gocache`.

Les **huit tests Go principaux passent avant et après** : quatre rotations,
versions persistantes, sous-arbre retenu/enfant partagé, entrée nominale,
séquences consommantes, comparaison de forme/couleurs, recoloration avec
donneur distinct et insertion avec donneur distinct. Les attentes de
recoloration sont conservées puisque ce prototype ne modifie pas `makeBlack`.

Les builds et tests précèdent les mesures. Cinq paires de processus alternent
AB, BA, AB, BA, AB, sans compilation simultanée. Chaque processus chauffe trois
appels puis mesure cinq appels de 100 000 insertions. Le GC est forcé avant
chaque appel, hors chronométrage ; le résultat reste vivant jusqu'au relevé
mémoire. Les 50 échantillons valident tous la profondeur attendue, 22.

## Résultats

Chaque ligne représente la médiane des cinq appels de chaque processus.

| Paire | Ordre | Avant | Après | Variation |
| --- | --- | ---: | ---: | ---: |
| 1 | AB | 10,751417 ms | 9,423834 ms | −12,348 % |
| 2 | BA | 9,447833 ms | 9,605792 ms | +1,672 % |
| 3 | AB | 9,468167 ms | 9,358959 ms | −1,153 % |
| 4 | BA | 9,508291 ms | 9,305084 ms | −2,137 % |
| 5 | AB | 9,629042 ms | 9,334917 ms | −3,055 % |

Quatre paires sur cinq favorisent le prototype. Variation appariée médiane :
**−2,137 %**, plage **−12,348 à +1,672 %**. Médianes des 25 échantillons par
variante : **9,508291 → 9,423834 ms**, soit **−0,888 %**. Les deux agrégations
répondent à des calculs différents et sont toutes deux conservées.

La première médiane avant est nettement supérieure aux suivantes. Aucune
cause n'a été établie et aucun échantillon n'est retiré. Cette première paire
ne doit pas devenir l'estimation du gain annoncé.

Allocations identiques : **100 000 / 3 200 000 octets** en médiane, avec des
plages **100 000–100 001 / 3 200 000–3 200 016** dans les deux variantes.

Le README officiel d'altbak indique **13,01 ms compilé / 11,09 ms manuscrit**
pour la suite, et **9 741,25 / 8 752,42 µs** pour RBTree. Cette sonde avec GC
préalable emploie un protocole différent des minima du README ; elle ne
remplace pas ces baselines. Le manuscrit et la suite complète n'ont pas été
remesurés. Une confirmation reste nécessaire avant d'affirmer un gain stable
ou de généraliser ce traitement dans le backend.

Preuves : `prototype.patch`, `segments.json`, `provenance.json`, `results.json`,
`measurements/`, `before/probe/correctness.log`, `after/probe/correctness.log`,
informations de build et `final-audit.json`. Les empreintes confirment que les
copies figées et la source conservée sont restées inchangées pendant les tests.
