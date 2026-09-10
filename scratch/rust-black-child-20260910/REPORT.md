# Parents noirs sans rotation : goulot résiduel après B13

Audit et prototype isolé du **10 septembre 2026**. Le meilleur candidat reste le chemin qui reconstruit un parent noir de RBTree après l'insertion d'un enfant, alors qu'aucune rotation n'est nécessaire. Le prototype confirme maintenant un gain sur la suite : **16,150 → 11,807 ms**, soit **4,343 ms (−26,9 %)**. **Aucune nouvelle règle n'est intégrée à Purust.**

## État examiné et baseline

Le [README officiel](../../../altbak.pub/README.md#rust) donne **RBTree compilé 15,956 ms**, **natif de la dernière colonne 36,070 ms**, et **total compilé 17,12 ms**. RBTree représente environ 93 % de ce total. Le compilé dépasse déjà ce natif ; les nouvelles mesures comparent donc le compilé à lui-même.

`altbak.pub-rust/run/bak/rust/output/purescript` n'existe pas sur cette machine. La sortie du checkout normal date du 8 septembre et précède B13. L'audit utilise la sortie du worktree `altbak.pub-purust/run/bak/rust/output/purust_output`, générée le 9 septembre à 23:53, avec B13 rouge intégré. Son empreinte est `fc60ee33c692f8729d61ff9a5a088fc5ded8cb519cb3ba980d641879d3e92949` ; elle est identique à celle du diagnostic précédent. [Métadonnées](metadata.json).

Le nouveau runner témoin mesure **15,115 ms pour RBTree et 16,150 ms au total**. L'écart avec le README interdit de soustraire directement les résultats du prototype à la baseline historique ; le gain annoncé vient des processus comparés dans la même série.

## Coût observé

Dans [ins](../../run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs), les deux nouvelles branches de B13 conservent les champs des parents rouges. Pour les parents noirs, le code extrait encore le payload entier, appelle récursivement `ins`, puis appelle `balance__purust_reuse`. Même sa branche sans rotation reconstruit tous les champs et reteste l'unicité.

Pour 100 000 insertions, le diagnostic exécuté puis le nouveau comptage des prototypes retrouvent exactement :

| Cas après l'appel récursif | Nombre |
| --- | ---: |
| Parent noir sans rotation | **1 368 968** |
| Parent noir avec rotation | **99 978** |

Il faut décider **après** l'appel : son résultat peut créer le motif qui déclenche la rotation. La présence d'un parent noir avant l'appel ne suffit pas à justifier une reconstruction simple.

## Trois variantes isolées

[probe.py](probe.py) conserve le noyau généré et ajoute un chemin pour les parents noirs uniques. Le cas partagé ou porteur d'une référence faible continue dans le code d'origine. Les doublons continuent aussi dans le chemin d'origine.

1. **Avant** : code généré actuel, B13 rouge déjà intégré.
2. **Témoin avec reconstruction** : extraire le payload sous un emprunt mutable conservé, calculer l'enfant, réinstaller tout le payload, puis tester les quatre motifs de rotation par emprunt. Sans rotation, retourner le parent. Avec rotation, réextraire le payload et appeler le worker existant avec l'enfant déjà calculé.
3. **Champs conservés** : même décision après appel ; détacher seulement l'enfant en utilisant un clone temporaire du frère comme emplacement initialisé, calculer/réinstaller cet enfant, puis retourner le parent sans rotation. En cas de rotation, extraire le payload et appeler le même worker existant.

Le témoin isole le raccourci de contrôle et les tests empruntés, tout en gardant une reconstruction complète. Son gain combine plusieurs effets : contournement du helper pour la branche simple, suppression de ses projections temporaires et de son retest. **Ce n'est pas le gain de la seule propagation d'unicité.** La différence avec la troisième variante mesure ensuite le maintien des champs, y compris l'évitement du cycle extraction/réinstallation préalable aux rotations dans ce témoin.

## Temps mesurés

| Variante | Noyau isolé | RBTree dans la suite | Total suite |
| --- | ---: | ---: | ---: |
| Avant | 15,909 ms | 15,115 ms | 16,150 ms |
| Témoin avec reconstruction | 11,824 ms | 11,798 ms | 12,855 ms |
| Champs conservés | 10,700 ms | 10,705 ms | 11,807 ms |

Le gain total dans la suite se décompose en **3,295 ms**, puis **1,048 ms supplémentaires**. RBTree seul gagne **4,410 ms (−29,2 %)**. Les deux étapes sont favorables dans chacun des cinq blocs.

| Bloc | Avant, total | Reconstruction, total | Champs conservés, total |
| --- | ---: | ---: | ---: |
| 1 | 16,441 ms | 13,070 ms | 11,779 ms |
| 2 | 15,697 ms | 12,559 ms | 11,587 ms |
| 3 | 16,156 ms | 12,971 ms | 12,204 ms |
| 4 | 15,984 ms | 12,846 ms | 11,813 ms |
| 5 | 16,439 ms | 12,619 ms | 11,830 ms |

Le noyau utilise cinq blocs de trois processus, un échauffement et quinze échantillons par processus, avec construction, profondeur et destruction incluses. Le runner complet conserve son échauffement et son meilleur de dix essais par benchmark. L'ordre des trois variantes tourne et s'inverse ; les totaux de synthèse sont les sommes des médianes de chaque benchmark. Profil **O1/mimalloc**, sans instrumentation ni compilation concurrente aux mesures.

[runner.py](runner.py) compile des copies isolées d'App et de RBTree, avec les autres sources de crates conservées. Il vérifie les **14 résultats de chaque processus** et contrôle que le Rust généré d'origine reste intact. Les temps des autres benchmarks varient aussi : Records passe notamment de 0,563 à 0,623 ms dans cette série, bien que sa source soit inchangée. Le gain de 4,343 ms inclut cette différence ; nous ne la soustrayons pas artificiellement. [Mesures de suite](runner-results.json), [mesures du noyau](timings.json).

## Comptage indépendant

[count.py](count.py) utilise le wrapper d'observation `Rc` des expériences précédentes. Ce sont des opérations logiques, pas un décompte d'instructions machine. [Résultats par phase](counts.json).

| Construction unique de 100 000 nœuds | Avant | Reconstruction | Champs conservés |
| --- | ---: | ---: | ---: |
| Allocations | 100 001 | 100 001 | 100 001 |
| Extractions entières | 1 668 902 | 1 768 880 | 299 934 |
| Reconstructions entières, helper + écritures directes | 1 668 902 | 1 768 880 | 299 934 |
| `get_mut` réussi | 4 152 834 | 2 783 866 | 2 783 866 |
| Clones | 3 775 130 | 1 014 986 | 2 483 932 |
| Relâchements temporaires | 3 675 130 | 914 986 | 2 383 932 |

Le témoin ajoute une extraction/reconstruction sur les 99 978 rotations, puisqu'il remet le payload avant de décider. Il supprime néanmoins **2 760 144 clones temporaires** et **1 368 968 retests** en évitant le chemin complet de balance lorsque le parent peut être conservé. La variante finale ajoute **1 468 946 clones/relâchements temporaires du frère** par rapport au témoin, mais évite l'extraction/réinstallation complète avant la décision. Au bilan par rapport au code actuel, **1 368 968 extractions/reconstructions** disparaissent. Les allocations et libérations finales restent identiques.

## Validations et portée

Les trois variantes passent les quatre rotations, les invariants rouge/noir, les clés exactes comparées à `BTreeSet`, 200 versions conservées, 512 insertions avec partage mixte et doublons, les bornes `i64`, les enfants partagés indépendamment, les références faibles et la libération finale des propriétaires. Le test ajoute **128 combinaisons de partage/faiblesse sur parents noirs**, en plus des 128 cas rouges existants. Les variantes instrumentées passent également les équilibres globaux de handles et de cellules. [Validation native](validation.json).

Cela valide ce prototype sur les scénarios exécutés. La performance des charges fortement persistantes n'est pas mesurée, et aucune nouvelle fixture TAST de génération ou preuve générale sur les paniques n'est fournie à cette étape.

## Règle générale à établir avant intégration

Le raccourci expérimental mentionne explicitement les constructeurs de RBTree **dans les copies Rust du scratch**. Il ne doit pas être transféré sous cette forme au compilateur. Purust n'a reçu aucune branche propre au benchmark.

Actuellement, [ChildCalls](../../../purust/purust/src/Purust/ChildCalls.purs) ne reconnaît que des gardes sur paramètres enums `Copy`, évaluables avant l'appel. [ChildUpdates](../../../purust/purust/src/Purust/ChildUpdates.purs) exclut explicitement une garde dépendant du résultat récursif. L'extension doit reconnaître dans le TAST une **feuille qui reconstruit le même constructeur après des tests sur l'enfant calculé**, avec :

- des chemins de projections typés, protégés par leurs tests de tags et respectant l'ordre des branches ;
- un seul calcul de l'enfant, réutilisé dans la reconstruction simple comme dans les branches complexes ;
- les restrictions actuelles sur le dernier usage du parent, le graphe natif fermé, les représentations et les payloads sans destructeur opaque ;
- une fixture indépendante du benchmark couvrant les deux issues, les partages, références faibles et libérations sur panique.

**B13/B14 restent jaunes.** La prochaine étape utile est cette preuve de génération générale, puis une mesure du Rust réellement émis. Les 4,343 ms constituent un résultat du prototype, pas une promesse pour l'intégration.

Reproduction depuis ce dossier : `python3 probe.py prepare`, `python3 probe.py validate`, `python3 count.py`, `python3 probe.py time`, `python3 runner.py build`, puis `python3 runner.py time`. Terminer les compilations/comptages avant chaque chronométrage. Les snapshots, binaires et logs sont conservés sous `build/`, ignoré par Git.
