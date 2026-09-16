# Recoloration Go : suppression conditionnelle des écritures inchangées

Expérience du 17 septembre 2026. **Aucun gain temporel démontré** pour ce
prototype limité à deux fonctions. Aucun fichier du backend n'a été modifié.

## Changement isolé

Deux copies des 1 989 fichiers de l'output courant ont été conservées depuis
`altbak.pub-gopurs/run/bak/go/modes/pure/output`. Seul `Test_RBTree.go` diffère.
Le SHA initial de ce fichier est
`b93c7cf9a01100043f900e975de5ec4ab94f69a34515be1f1214e2ee390a6527`.

Dans `__gopurs_owned_insert_0_consume`, les écritures V1/V2/V3 sont exécutées
seulement lorsque `__cell_12 != __let_tree_5`. Dans
`__gopurs_owned_makeBlack_0_consume`, la garde est `__cell_6 != __arg0`.
La sélection des cellules, Rc et la couleur restent identiques. Une cellule
donneuse distincte reçoit donc toujours tous les champs. Les lectures des
valeurs sont conservées à leur place initiale. Voir `prototype.patch`.

## Protocole

Sondes préparées avec `stage-harness.py` de la campagne du 16 septembre ;
`main.go` et les adaptateurs sont identiques avant/après. L'entrée construit
depuis `nil` avec le worker réellement choisi par le benchmark, puis mesure la
profondeur. Go 1.27.0 darwin/arm64, `GOGC=800`, `GOWORK=off`, PGO désactivé ;
GOMAXPROCS et GOMEMLIMIT non fixés. Cache de compilation partagé existant :
`/private/tmp/gopurs-adt-reuse-gocache`.

Les deux variantes ont été compilées et testées avant les mesures. Ensuite,
cinq paires de processus avant/après alternent leur ordre : AB, BA, AB, BA, AB.
Chaque processus chauffe trois appels puis mesure cinq appels de 100 000
insertions. Le GC est forcé avant chaque appel, hors chronométrage. Le résultat
reste vivant jusqu'au relevé mémoire. Aucun build simultané n'a été lancé.

`measure.py --pair N` conserve une paire sans écraser les précédentes ;
`measure.py --summary` produit `results.json` à partir des cinq paires.

## Résultats

Chaque ligne donne la médiane des cinq appels de chaque processus. Variation
positive : prototype plus lent.

| Paire | Ordre | Avant | Après | Variation |
| --- | --- | ---: | ---: | ---: |
| 1 | AB | 10,681792 ms | 10,660791 ms | −0,197 % |
| 2 | BA | 10,126750 ms | 9,978042 ms | −1,468 % |
| 3 | AB | 9,881250 ms | 9,886667 ms | +0,055 % |
| 4 | BA | 9,922792 ms | 10,017583 ms | +0,955 % |
| 5 | AB | 9,832125 ms | 10,046334 ms | +2,179 % |

La variation appariée médiane est **+0,055 %**, avec une plage de −1,468 à
+2,179 %. Deux paires sont favorables, trois défavorables. La médiane des
25 échantillons par variante passe de **9,992250 à 10,050541 ms** (+0,583 %).
Ces observations ne démontrent ni gain, ni régression temporelle stable.
Les deux médianes de la première paire sont supérieures à celles des suivantes.
La cause de cet écart n'a pas été établie ; il n'est pas attribué à une chauffe
insuffisante, au GC ou à un autre mécanisme.

Les allocations restent identiques : **100 000 allocations / 3 200 000 octets**
en médiane, avec des plages 100 000–100 001 / 3 200 000–3 200 016 dans les deux
variantes. Les 50 résultats mesurés donnent tous la profondeur attendue, 22.

Le README officiel d'altbak indique **13,01 ms compilé contre 11,09 ms
manuscrit** pour la suite complète, et 9 741,25 contre 8 752,42 µs pour RBTree.
Cette sonde emploie un protocole différent des minima du README et ne remplace
pas ses baselines. Le manuscrit et la suite complète n'ont pas été remesurés.

## Correction et conclusion

Les **huit tests Go principaux passent dans chaque variante** : quatre
rotations, historiques persistants, sous-arbre retenu/enfant partagé, entrée
nominale, séquences consommantes, comparaison des formes/couleurs, recoloration
avec donneur distinct, insertion avec donneur distinct. Les six premiers
proviennent du harnais archivé ; les deux derniers vérifient spécialement le
repli où toutes les écritures doivent subsister.

Une inspection bornée des binaires de mesure confirme que les écritures n'ont
pas été éliminées par Go : l'`insert` avant comporte `STP` pour V1/V2, `STR`
pour V3, et un chemin conditionnel vers `runtime.gcWriteBarrier4`. Après, le
`CMP` et le `B.EQ` ajoutés contournent les écritures et ce chemin lorsque la
cellule retenue est la cellule source. Les lectures initiales restent présentes.
Le cadre de pile passe de 0x50 à 0x60 octets ; le code de la fonction grossit.
La barrière et les affectations restent présentes dans le repli qui copie vers
une autre cellule. L'état dynamique de la barrière pendant les mesures n'a pas
été instrumenté.

Le symbole `makeBlack` n'est pas lié dans les binaires de cette sonde ; la
recoloration effectivement mesurée se trouve dans `insert`. Les tests Go
appellent aussi `makeBlack` directement. L'outil `go tool objdump` est absent
de cette installation ; `nm` et le `llvm-objdump` de Xcode ont permis cette
inspection sans recompilation. Extraits et symboles : `assembly/`.

Les résultats ne justifient pas l'intégration de cette garde comme optimisation
temporelle. Ils ne concluent rien sur une future suppression statique des
écritures ou sur une modification de l'équilibrage, qui n'ont pas été testées.

Preuves conservées : `provenance.json`, `prototype.patch`, `results.json`,
`measurements/`, `before/probe/correctness.log`, `after/probe/correctness.log`
et `final-audit.json`. Les empreintes finales confirment que les deux copies
figées et le RBTree d'origine sont restés inchangés pendant la validation.
