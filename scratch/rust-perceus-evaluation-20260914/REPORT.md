# Intérêt de nouvelles optimisations inspirées de Perceus

14 septembre 2026 — prototypes dans `altbak.pub-purust/scratch/rust-perceus-evaluation-20260914`.

**Verdict : une extension ciblée des parcours empruntés mérite d'être envisagée ; une intégration générale de Perceus et un nouvel investissement dans le sticky sharing ne sont pas justifiés par ces résultats.** Le meilleur compromis essayé consomme les cellules uniques et emprunte les sous-arbres partagés. Il donne un petit signal favorable sur le benchmark actuel, et un gain net lorsque le même arbre est lu plusieurs fois. L'essentiel de ce second gain est déjà obtenu par le simple emprunt.

Aucun changement dans Purust ni dans sa sortie Rust courante : **868 fichiers de Purust et 611 fichiers générés/manifests/verrou sont inchangés**, vérifiés par SHA-256. Les prototypes, binaires et rapports restent dans ce scratch. [Vérification](no-product-changes.json).

## 1. RBTree : quatre versions de la même fonction générée

Le témoin inclut les optimisations FBIP et de réutilisation des champs déjà intégrées. On remplace seulement `Test_RBTree_depth` dans une copie de son module ; construction, insertions, rotations, représentation et allocator sont conservés.

| Variante | Comportement |
| --- | --- |
| Témoin | Parcourt les enfants via des clones de Rc, puis libère le propriétaire. |
| Emprunt | Parcourt `&Tree`, puis libère le propriétaire avant de retourner. |
| Consommation | `Rc::unwrap_or_clone` déplace les champs des cellules uniques ; les cellules partagées clonent leur contenu. |
| Hybride | `Rc::try_unwrap` consomme les cellules uniques ; en cas de partage, un worker parcourt le sous-arbre par emprunt. |

Ces prototypes évaluent une forme de spécialisation des libérations/transfert des champs et sa combinaison avec les emprunts. **Ils n'implémentent pas l'algorithme Perceus complet.** `unwrap_or_clone` est déjà émis ailleurs, notamment dans `foldl` de ListOps ; une extension ciblée de la couverture des consommateurs pourrait donc suffire. Les états de la matrice ne sont pas modifiés. Le [papier Perceus](https://www.microsoft.com/en-us/research/wp-content/uploads/2020/11/perceus-tr-v4.pdf) distingue le comptage précis, l'analyse de réutilisation et les spécialisations ; les emprunts seuls ne démontrent pas toute cette chaîne.

### Runner altbak complet : construction + profondeur + destruction

Huit blocs, quatre variantes par bloc ; quatre rotations de l'ordre puis leurs miroirs. Chaque variante occupe chaque position deux fois. Le warm-up et le meilleur de dix du runner sont conservés. Les 14 valeurs attendues sont vérifiées à chaque exécution.

| Variante | RBTree, médiane | Écart au témoin | Blocs favorables | Total, somme des médianes |
| --- | ---: | ---: | ---: | ---: |
| Témoin | 10,9945 ms | — | — | 11,8800 ms |
| Emprunt | 11,0190 ms | +0,22 % | 5/8 | 11,9065 ms |
| Consommation | 10,8780 ms | −1,06 % | 6/8 | 11,7675 ms |
| Hybride | **10,8475 ms** | **−1,34 %** | **7/8** | **11,7530 ms** |

L'hybride représente **0,147 ms sur RBTree** et **0,127 ms sur le total (−1,07 %)** dans cette série. Le signal est modeste : les plages se recouvrent, avec 10,874–11,346 ms pour le témoin et 10,588–11,010 ms pour l'hybride. Les petites variations des autres benchmarks, dont le code est identique, ne sont pas attribuées au prototype. Il n'est pas établi que l'hybride soit systématiquement meilleur que les deux autres transformations dans ce workload.

Une première série distincte, avant l'ajout de l'hybride, donnait 11,0435 → 10,7835 ms pour la consommation (−2,35 %, 7/8) et 10,820 ms pour l'emprunt (−2,02 %, 6/8). Elle confirme que le gain sur ce benchmark est petit et variable ; elle n'est pas fusionnée artificiellement avec la seconde série. [Série finale](full-runner-results.json), [série initiale](full-runner-results-initial.json).

### Même arbre, vingt lectures d'une racine conservée

Ce scénario reste proche d'altbak : mêmes fonctions générées de construction et profondeur, 100 000 insertions, puis vingt appels de profondeur sur une racine conservée, et destruction finale. Tout ce travail est inclus dans chaque échantillon. La somme des profondeurs vaut 440 dans les quatre versions. Huit blocs équilibrés, un warm-up puis sept mesures par processus ; le tableau donne la médiane des médianes des processus.

| Variante | Temps | Écart | Blocs favorables |
| --- | ---: | ---: | ---: |
| Témoin | 17,077 ms | — | — |
| Emprunt | **13,871 ms** | **−18,78 %** | **8/8** |
| Consommation | 17,778 ms | +4,11 % | 0/8 |
| Hybride | **13,842 ms** | **−18,95 %** | **8/8** |

Le gain des deux variantes empruntées est net : leurs plages restent entièrement sous celle du témoin. L'écart entre hybride et simple emprunt est minuscule et leurs plages se recouvrent largement. **La consommation systématique pénalise ce cas partagé ; le bénéfice établi vient principalement de l'emprunt.** [Mesures complètes](b15/shared-timings.json), [harness](b15/shared_runner.py).

### Pourquoi les compteurs ne suffisent pas

Dans le parcours puis la destruction de l'arbre unique, les trois prototypes suppriment les **200 000 clones Rc** du témoin. Ils n'ajoutent aucune allocation, et les **100 001 cellules** sont libérées dans toutes les variantes. La construction reste inchangée.

Sur le parcours d'une racine conservée de 200 nœuds, témoin et consommation effectuent chacun 401 clones ; emprunt et hybride n'en effectuent qu'un, celui de l'appelant. L'hybride ajoute un test d'extraction initial qui échoue, puis emprunte. [Compteurs](b15/counts.json).

Les validations couvrent les invariants rouge/noir, les rotations, 200 versions persistantes, 512 insertions à partage mixte, les enfants partagés, références faibles, lectures répétées et libérations finales. Elles passent pour les quatre variantes. La portée reste un arbre natif fermé, à champs scalaires et récursifs, sans callback ni destructeur opaque. Des contre-exemples exécutés montrent qu'une généralisation sans cette preuve peut changer l'ordre observable des destructions et l'observation d'une référence faible. [Détail et validations](b15/REPORT.md).

## 2. List Processing : spécialiser le filtre seul ralentit

Le test reprend le vrai code généré, y compris ses adaptateurs, fermetures et valeurs récursives. Seules les deux copies du filtre, publique et intégrée dans `sumEvens`, consomment le nœud avec `unwrap_or_clone`. Range, fold, accumulation et boxing restent inchangés. [Diff exact](list/variant.diff).

Dans la première série du runner complet, **List Processing passe de 36 à 38 µs (+5,56 %), avec huit blocs sur huit défavorables**. La résolution du runner est de 1 µs ; il s'agit d'un petit coût absolu, pas d'un gain à intégrer.

Les compteurs donnent une explication structurelle : **aucune itération ne trouve de nœud unique**. Les 944 itérations des six tailles de `sumEvens` et les 15 072 itérations des cas du filtre empruntent toutes le chemin partagé. Les adaptateurs gardent le propriétaire qui sert à extraire puis cloner leur argument, pendant l'appel. Le prototype ne rétablit donc pas les transferts de champs attendus. Ce comptage ne prétend pas expliquer chaque instruction responsable du ralentissement.

96 cas de partage/références faibles et six tailles de somme donnent les résultats attendus, dont 202 950 pour n=900. [Validation et compteurs](list/validation.json), [transformation reproductible](list_probe.py). Une amélioration des durées de vie dans ces adaptateurs serait une autre expérience ; ce test n'en établit pas le gain.

## 3. Sticky sharing : aucun intérêt démontré sur ces charges

Le runtime local actuel sature à `u32::MAX`, soit **4 294 967 295 propriétaires**. Les scénarios normaux mesurés culminent à **un, deux ou trois propriétaires simultanés**. Un million de clones successifs ne rapproche pas nécessairement le compteur du seuil : les drops intermédiaires le font redescendre.

Pour mesurer le mécanisme malgré cela, l'expérience force uniquement les trois cellules initiales à ce seuil dans une copie privée du runtime. Les nouvelles cellules conservent leur compteur normal. C'est un cas artificiel de valeurs immortelles, explicitement nommé `forced-sticky`, et non un accélérateur observé dans le benchmark normal.

Sept paires alternées, un million d'itérations par scénario, après une chauffe non chronométrée de 10 000 itérations. Setup, saturation et capture sont hors timer ; travail, observation des résultats et destruction des propriétaires sont inclus. Le noyau de mise à jour est celui réellement généré pour Records ; les lectures sont des scénarios complémentaires sur ses records imbriqués.

| Scénario | Compteur normal | Saturation forcée | Écart |
| --- | ---: | ---: | ---: |
| Lectures avec clones | 13,866 ms | 15,249 ms | **+9,97 %**, 7/7 défavorables |
| Lectures empruntées | 4,626 ms | 4,623 ms | −0,07 %, variation |
| Lectures via fermeture | 15,659 ms | 16,111 ms | **+2,89 %**, 7/7 défavorables |
| Updates, initialement unique | 41,067 ms | 40,783 ms | −0,69 %, variation |
| Updates, version initiale conservée | 41,464 ms | 41,707 ms | +0,59 %, variation |

La saturation forcée conserve **168 octets par graphe initial** et impose trois copies supplémentaires à la première mise à jour d'un graphe initialement unique. Les cellules copiées redeviennent normales. Chaque processus chronométré crée un graphe de chauffe puis un mesuré, donc retient 336 octets sous saturation forcée.

Même dans ce cas artificiel, aucune accélération convaincante n'apparaît. Les écritures logiques de compteur disparaissent dans la boucle de lecture possédée, mais le binaire ralentit ; la cause machine précise n'a pas été profilée. La variante multithread de PerceusPtr délègue à Arc et n'implémente pas cette saturation ; RBTree emploie Rc. **Aucun gain n'est extrapolé au runner complet.** [Rapport sticky](sticky/REPORT.md), [mesures](sticky/timings.json), [compteurs](sticky/counts.json).

## Décision proposée

1. **Étudier une extension ciblée des parcours empruntés**, avec éventuellement un chemin consommant lorsque l'unicité est prouvée. Elle a un bénéfice clair sur les lectures répétées, et l'hybride évite la pénalité du prototype consommant sur le partage. Le bénéfice sur la suite actuelle reste autour d'un dixième de milliseconde dans la série finale.
2. **Ne pas intégrer la consommation systématique du filtre de ListOps telle quelle** : elle ralentit, sans atteindre son chemin unique.
3. **Ne pas prioriser davantage le sticky sharing pour ces workloads** : le seuil naturel n'est pas approché et l'essai forcé n'apporte pas de gain convaincant.
4. **Ne pas considérer cette étude comme une justification d'une grande passe Perceus générale**. Elle identifie des consommateurs précis à améliorer et les conditions de leur intérêt ; une éventuelle intégration devra prouver leur admissibilité dans le TAST.

## Protocole, références et reproduction

Les runners complets utilisent un snapshot de **611 fichiers**, les mêmes dépendances et le profil existant **O1, debug=true, mimalloc**. Les manifests isolés éliminent du lock les paquets hors fermeture sans changer aucune version ni checksum restant. Les quatre crates pertinentes sont nettoyées avant chaque construction. FFI Console/Bench vérifiées, empreintes de sources/binaires/verrou vérifiées. [Snapshot](full-runner-snapshot.json), [builds](full-runner-build.json), [lock](full-runner-lock.json).

Le harness de lectures partagées utilise les mêmes versions de mimalloc et le même profil. L'expérience sticky est distincte : **O1, overflow-checks=yes et allocator System**. Ses variantes ont les mêmes options, mais leurs temps ne se comparent pas directement à ceux du runner altbak. Les compteurs sont toujours exécutés séparément des mesures ; aucun build, test ou instrumenteur de cette tâche ne tourne pendant les séries retenues. L'activité extérieure à cette tâche n'est pas contrôlée. [Environnement](environment.json), [résumé calculé](summary.json).

Le [README officiel](../../../altbak.pub/README.md#rust), relu pour cette étude, indique Rust **RBTree 11,631 ms / Records 0,382 ms / total 12,56 ms**, contre **36,070 / 0,004 / 36,13 ms** dans sa dernière colonne native. Ces références historiques sont distinctes de nos comparaisons appariées. Le run utilisateur **12,198 ms / 13,17 ms total** appartient également à une autre session. Aucun pourcentage d'amélioration n'est déduit en soustrayant ces sessions.

Les scripts de préparation restent dans ce dossier et n'appellent pas le générateur Purust : `full_runner.py snapshot`, `b15/probe.py prepare`, `b15/probe.py validate`, `b15/probe.py count`, `list_validate.py`, puis les commandes build des runners. Les sources figées et binaires volumineux sont conservés dans `build/`, ignoré par Git.

Les commandes de mesure s'exécutent **séquentiellement**, lorsque les autres tâches de cette étude sont arrêtées :

```sh
python3 full_runner.py time --variant before --variant borrowed --variant consuming --variant hybrid --blocks 8 --coordinated
python3 b15/shared_runner.py time
python3 sticky/probe.py time --timing-authorized
python3 summarize.py
```

Les données initiales et finales sont toutes conservées ; les résultats défavorables restent dans le compte-rendu.
