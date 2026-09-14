# Sticky sharing : saturation, contrôle du surcoût et partage élevé

Les six scénarios ont été exécutés avec trois variantes. **Aucun intérêt supplémentaire du sticky sharing n'est établi sur ce corpus**, même avec une saturation forcée ou un million de propriétaires conservés simultanément. Supprimer les gardes de saturation n'améliore pas non plus systématiquement le runtime normal.

## Protocole

Le noyau de mise à jour est celui du vrai Records généré, avec le runtime local `PerceusPtr` et le même `purust_core` que l'étude précédente. Le profil isolé reste O1, vérification des débordements activée et allocateur System. Ces temps ne sont pas comparés directement à ceux du runner altbak, qui emploie mimalloc.

| Variante | Changement |
| --- | --- |
| `normal` | Runtime actuel à compteur `u32` saturant. |
| `checked` | Même représentation ; les deux gardes de sticky sont supprimées, et l'incrément devient un `checked_add` qui refuse le débordement. |
| `forced` | Même binaire que `normal` ; seuls les compteurs des trois cellules initiales sont forcés à `u32::MAX`. Les copies nouvelles repartent avec un compteur normal. |

Les comportements normal et checked sont équivalents sur les valeurs de compteur parcourues ici. L'expérience checked est une ablation des chemins de saturation avec protection de débordement, pas une proposition de changement de sémantique du runtime.

Les 18 cas instrumentés concordent avec 18 exécutions sans compteurs. Normal et checked effectuent exactement les mêmes opérations de comptage et libèrent leurs trois cellules initiales. Le test `fanout` valide **100 001 propriétaires forts simultanés** avec compteurs, puis mesure un million de clones conservés ensemble, au lieu de clones successifs immédiatement détruits. Le seuil effectif reste 4 294 967 295 ; il n'est atteint naturellement dans aucun cas. [Compteurs](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/counts.json), [harnais](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/harness.rs), [préparation et mesure](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/probe.py).

Chaque scénario temporel a six blocs équilibrés : trois rotations, puis leurs inverses. Un processus fait 10 000 itérations de chauffe et un échantillon d'un million d'itérations, ou d'un million de clones simultanés pour fanout. Setup et saturation sont exclus ; travail, checksum et libération finale des propriétaires sont inclus. Fanout inclut l'allocation et la destruction de son Vec. Les compteurs, séparés des temps, concernent uniquement les cellules PerceusPtr : ils n'incluent pas ce Vec ni les cellules Rc des fermetures.

## Mesures de cette campagne

Médiane des six processus, en millisecondes ; écart par rapport à normal.

| Scénario | Normal | Checked | Écart checked | Forced | Écart forced |
| --- | ---: | ---: | ---: | ---: | ---: |
| Lectures avec clones | 12,964 | 14,331 | +10,54 % | 13,937 | +7,50 % |
| Lectures empruntées | 4,624 | 4,620 | −0,09 % | 4,611 | −0,28 % |
| Lectures via fermeture | 13,561 | 15,159 | +11,78 % | 14,774 | +8,94 % |
| Mises à jour uniques | 39,799 | 40,876 | +2,71 % | 39,763 | −0,09 % |
| Version initiale conservée | 39,528 | 40,674 | +2,90 % | 39,949 | +1,07 % |
| Un million de clones simultanés | 8,570 | 8,422 | −1,73 % | 8,763 | +2,25 % |

Checked est défavorable dans les six blocs des lectures possédées, des lectures via fermeture et des deux mises à jour. Son petit gain médian sur fanout n'est favorable que dans quatre blocs sur six, avec plages qui se recouvrent. Forced est défavorable dans les six blocs des lectures possédées et via fermeture. Ses minuscules gains médians sur lectures empruntées et updates uniques ne sont favorables que dans trois blocs sur six.

La saturation forcée conserve **trois cellules, 168 octets par graphe**, donc 336 octets pour le graphe de chauffe et celui mesuré d'un processus. Les nouvelles copies sont libérables normalement. La première mise à jour d'un graphe initialement unique impose trois copies supplémentaires. L'élimination logique des écritures de compteur dans les lectures saturées ne suffit pas à accélérer le binaire ; la cause machine du ralentissement n'a pas été profilée. [Mesures brutes](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/timings.json), [empreintes](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/sticky/metadata.json).

## Raccord avec la première étude

La première campagne de sept paires donnait déjà, pour normal → forced : lectures possédées 13,866 → 15,249 ms (+9,97 %) ; empruntées 4,626 → 4,623 ms ; fermetures 15,659 → 16,111 ms (+2,89 %) ; updates uniques 41,067 → 40,783 ms ; version conservée 41,464 → 41,707 ms. Les deux régressions de lecture étaient présentes dans les sept paires. Ces résultats sont conservés comme **session distincte**, sans fusion des temps ni sélection de la meilleure série. [Premier rapport](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-evaluation-20260914/sticky/REPORT.md).

Cette étude ne mesure pas une saturation naturelle à plusieurs milliards de propriétaires, ni les threads. Le runtime threaded délègue à Arc et n'a pas ce comportement saturant effectif ; les ADT natifs comme RBTree utilisent Rc. Les résultats ne justifient ni d'étendre sticky sharing pour ces benchmarks, ni de supprimer le mécanisme existant pour gagner du temps.

## Reproduction

Depuis ce dossier, `python3 probe.py build`, puis `python3 probe.py check`. Lorsque les autres builds, tests et mesures sont arrêtés : `python3 probe.py time --coordinated`. Toutes les variantes et leurs dépendances restent dans le scratch.
