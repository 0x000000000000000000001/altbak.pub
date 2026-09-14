# FBIP : réutiliser la cellule puis modifier seulement les champs utiles

Étude **synthétique**, isolée de Purust et de sa sortie compilée. Les quatre variantes utilisent exactement le même arbre natif `Rc<CellData>` : un ADT `Empty | Node` avec deux enfants et deux `i64`. Le wrapper `CellData` permet de compter la destruction du payload ; il n'ajoute pas de champ au payload. Ce protocole étudie un mécanisme proche des ADT générés, sans constituer une nouvelle passe du compilateur ni une mesure du benchmark RBTree.

Les opérations alternent une modification scalaire au dernier niveau, une modification des deux scalaires à mi-profondeur et une permutation des deux enfants plus haut. L'algorithme, les chemins, les valeurs initiales, la représentation et le fallback persistant sont communs aux variantes.

| Stage | Traitement d'une cellule unique |
| --- | --- |
| `persistent` | Copie les enfants et alloue une nouvelle cellule à chaque reconstruction. |
| `reuse_rechecked` | Retire le payload complet, reconstruit, puis revérifie l'unicité avant de le remettre dans la même cellule. |
| `reuse_retained` | Même retrait/reconstruction du payload complet ; conserve l'emprunt exclusif pendant la récursion, sans seconde vérification. |
| `fields_retained` | Conserve ce même emprunt et ne modifie que les scalaires ou emplacements d'enfants concernés. |

Le passage `reuse_rechecked` → `reuse_retained` isole la seconde vérification d'unicité. Le passage `reuse_retained` → `fields_retained` isole le déplacement du payload complet. Il ne faut donc pas attribuer aux seuls champs spécialisés le gain d'allocations obtenu dès la première réutilisation.

## Couverture et validation

Profondeurs 7 et 10, chacune avec 128 et 2048 opérations, dans six scénarios : arbre unique ; huit anciennes racines conservées ; deux sous-arbres intérieurs conservés ; enfant partagé entre les deux branches de la racine ; un `Weak` à la racine ; trois `Weak` intérieurs. L'installation du scénario diamond construit d'abord un arbre complet puis remplace sa branche droite par un alias de la gauche ; ce coût commun est inclus dans les compteurs.

**96/96 cas passent dans chacun des deux exécutables**, normal et instrumenté. Un oracle indépendant `Box<Model>` vérifie l'égalité structurelle de l'arbre final et de chaque ancienne version ou sous-arbre conservé. Les résultats normaux et instrumentés concordent. Après destruction des propriétaires, tous les payloads ont expiré ; les nombres de cellules créées/détruites, de propriétaires créés/détruits et d'observateurs `Weak` créés/détruits sont équilibrés.

Les allocations comptées sont les cellules `Rc`, y compris les feuilles `Empty`. Ce ne sont pas des mesures générales de l'allocateur : les allocations temporaires de l'oracle et les blocs de contrôle conservés par `Weak` ne sont pas comptés séparément. L'expiration des propriétaires forts et la destruction de tous les observateurs sont néanmoins vérifiées.

## Compteurs à profondeur 10, 2048 opérations

| Scénario / stage | Cellules allouées | Dup de pointeurs | Tests d'unicité | Retraits du payload | Reconstructions du payload |
| --- | ---: | ---: | ---: | ---: | ---: |
| Unique / persistent | 15703 | 27312 | 0 | 0 | 13656 |
| Unique / reuse_rechecked | 2047 | 0 | 27312 | 13656 | 13656 |
| Unique / reuse_retained | 2047 | 0 | 13656 | 13656 | 13656 |
| Unique / fields_retained | 2047 | 0 | 13656 | 0 | 0 |
| Anciennes racines / persistent | 15703 | 27320 | 0 | 0 | 13656 |
| Anciennes racines / reuse_rechecked | 3929 | 3772 | 25430 | 11774 | 13656 |
| Anciennes racines / reuse_retained | 3929 | 3772 | 13656 | 11774 | 13656 |
| Anciennes racines / fields_retained | 3929 | 3772 | 13656 | 0 | 1882 |

Dans chaque scénario, les trois stages de réutilisation ont exactement les mêmes allocations et duplications. Les autres totaux d'allocations pour ces trois stages sont 2453 (partage intérieur), 2378 (diamond), 2057 (`Weak` racine) et 2068 (`Weak` intérieurs). Les cellules empêchées par un propriétaire ou un `Weak` empruntent le fallback persistant commun. Pour conserver les huit anciennes racines, 1882 cellules sont encore reconstruites dans ce protocole.

## Mesures coordonnées

Le coordinateur a ensuite mesuré les **24 combinaisons, huit blocs chacune**, dans le binaire sans compteurs : 768 processus, chacun avec un warmup vérifié et sept échantillons. Les [mesures brutes](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/timings.json) sont complètes ; leur empreinte de `metadata.json` correspond au fichier de préparation. Les médianes de chaque processus ont été recalculées depuis les sept échantillons, les huit blocs de chaque stage et les checksums communs aux quatre stages ont été vérifiés par lecture des résultats. Aucun nouveau benchmark n'a été exécuté pour cette analyse.

Les temps ci-dessous sont les **médianes des huit médianes de processus**, en microsecondes. Chaque variation est calculée à partir de ces temps ; `k/8` compte séparément les blocs où le nouveau stage est plus rapide que le précédent. Une variation négative signifie une réduction du temps. Les pourcentages des trois étapes ne s'additionnent pas.

### Profondeur 10, 2048 opérations : tous les scénarios

| Scénario | Persistent, µs | Reuse + recheck, µs | Borrow conservé, µs | Champs conservés, µs |
| --- | ---: | ---: | ---: | ---: |
| Unique | 204.563 | 83.688 | 78.271 | 52.876 |
| Huit anciennes racines | 208.917 | 110.667 | 108.000 | 77.250 |
| Partage intérieur | 204.188 | 94.501 | 86.875 | 57.813 |
| Diamond | 205.459 | 92.771 | 85.042 | 57.042 |
| `Weak` racine | 204.146 | 83.625 | 78.125 | 52.667 |
| `Weak` intérieurs | 206.813 | 83.917 | 79.334 | 52.021 |

| Scénario | Réutilisation : persistent → reuse_rechecked | Borrow : reuse_rechecked → reuse_retained | Champs : reuse_retained → fields_retained |
| --- | ---: | ---: | ---: |
| Unique | −59.09 % ; 8/8 | −6.47 % ; 7/8 | −32.45 % ; 8/8 |
| Huit anciennes racines | −47.03 % ; 8/8 | −2.41 % ; 5/8 | −28.47 % ; 8/8 |
| Partage intérieur | −53.72 % ; 8/8 | −8.07 % ; 7/8 | −33.45 % ; 8/8 |
| Diamond | −54.85 % ; 8/8 | −8.33 % ; 7/8 | −32.92 % ; 8/8 |
| `Weak` racine | −59.04 % ; 8/8 | −6.58 % ; 6/8 | −32.59 % ; 8/8 |
| `Weak` intérieurs | −59.42 % ; 8/8 | −5.46 % ; 7/8 | −34.43 % ; 8/8 |

La première étape permet le transfert des champs possédés et la réutilisation des cellules : en unique, elle supprime 13656 allocations et 27312 duplications, tout en ajoutant les tests d'unicité. Ce résultat mesure ce mécanisme complet, pas le coût de l'allocateur isolé. La deuxième conserve les allocations et duplications, mais supprime les tests répétés : son effet est plus modeste et moins régulier. La troisième conserve aussi les tests d'unicité, mais élimine les retraits et reconstructions du payload sur les chemins uniques ; son avantage est favorable dans les 48 blocs de ces six scénarios.

### Autres dimensions

Chaque ligne résume les six scénarios à dimension identique. La médiane et la plage portent sur les **six variations de scénarios**, pas sur une somme de temps ni un nouveau benchmark global. Les blocs favorables sont comptés sur 48 comparaisons appariées.

| Profondeur / opérations | Effet supplémentaire | Variation médiane des six scénarios | Plage des six variations | Blocs favorables |
| --- | --- | ---: | ---: | ---: |
| 7 / 128 | Réutilisation | −44.73 % | −50.00 à −30.46 % | 48/48 |
| 7 / 128 | Borrow conservé | −6.15 % | −8.98 à −2.27 % | 38/48 |
| 7 / 128 | Champs conservés | −24.83 % | −28.62 à −7.78 % | 45/48 |
| 7 / 2048 | Réutilisation | −63.18 % | −65.17 à −56.30 % | 48/48 |
| 7 / 2048 | Borrow conservé | −7.60 % | −9.26 à −7.25 % | 43/48 |
| 7 / 2048 | Champs conservés | −48.24 % | −49.31 à −39.54 % | 48/48 |
| 10 / 128 | Réutilisation | −17.47 % | −20.12 à −8.73 % | 46/48 |
| 10 / 128 | Borrow conservé | −1.77 % | −2.56 à −0.48 % | 40/48 |
| 10 / 128 | Champs conservés | −7.60 % | −8.64 à −0.82 % | 41/48 |

Aucune des 24 variations de médianes ne régresse pour chacune des trois étapes. Cela ne signifie pas que tous les blocs sont favorables : sur la matrice entière, les scores sont **190/192** pour la réutilisation, **160/192** pour le borrow et **182/192** pour les champs. Il reste respectivement 2, 31 et 10 inversions ponctuelles, plus une égalité pour le borrow.

Les plus petites différences ne justifient pas une promesse ferme. En unique à 10/128, conserver le borrow donne **30.437 → 30.291 µs (−0.48 %, 5/8)**. Avec anciennes racines à 10/128, conserver les champs donne **35.688 → 35.395 µs (−0.82 %, 5/8)** ; la réutilisation elle-même donne **39.876 → 36.396 µs (−8.73 %, 6/8)**. À 7/128 avec anciennes racines, les champs donnent **8.042 → 7.417 µs (−7.78 %, 5/8)**. La construction et la destruction de l'arbre font partie du chronométrage ; avec seulement 128 opérations et une profondeur de 10, leur coût commun pèse davantage. Les scores par bloc rendent visibles ces cas moins réguliers sans les masquer derrière une moyenne générale.

### Ce que ces temps disent du code actuel

La baseline `persistent` est volontairement synthétique et reconstruit systématiquement. Purust dispose déjà de réutilisation d'ADT et de spécialisations scalaires, d'enfants et de certaines permutations : [état B12/B13/B14](/Users/0x1/Documents/htdocs/altbak.pub/todo.md#perceus-fbip-and-sticky-sharing). Le nouveau résultat ne permet donc pas d'annoncer −47 à −59 %, ni −28 à −34 %, sur son RBTree actuel. Il montre l'intérêt d'étendre ces preuves à d'autres formes admissibles et de séparer leurs contributions.

Le [README officiel relu pendant cette étude](/Users/0x1/Documents/htdocs/altbak.pub/README.md#rust) indique **RBTree compilé 11.631 ms / natif 36.070 ms**, total **12.56 / 36.13 ms**, Records **382 / 4 µs** et List **34 / 1 µs**. Ces chiffres historiques situent la priorité ; ce ne sont pas les contrôles appariés de cette expérience. Le [RBTree actuellement généré](/Users/0x1/Documents/htdocs/altbak.pub-purust/run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs) contient déjà le helper de permutation et utilise des `std::rc::Rc` ; `depth` conserve encore des projections d'enfants clonées. Aucun de ces fichiers n'a été modifié ici.

## Reproduction et limites

`python3 probe.py build` compile hors réseau deux exécutables, avec et sans compteurs. `python3 probe.py validate` rejoue la matrice. Le profil release utilise `opt-level=1`, `debug=true` et mimalloc ; le lock résout `mimalloc 0.1.52` / `libmimalloc-sys 0.1.49`, comme l'étude précédente. Les empreintes de sources, lock et exécutables sont dans [metadata.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/metadata.json), les résultats dans [validation.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/validation.json) et [counts.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/counts.json).

La commande de mesure coordonnée est `python3 probe.py time --coordinated --blocks 8`, éventuellement avec `--scenario`, `--depth` et `--updates`. Quatre rotations puis leurs inverses équilibrent les positions des stages. Chaque processus normal réalise un warmup vérifié puis sept échantillons. Le segment chronométré comprend construction native, installation du partage, opérations, checksum et destructions ; la préparation des chemins et de l'oracle est exclue. `metadata.json` est le manifeste conservé de la préparation : son champ `timings_run: false` précède la campagne du coordinateur ; [timings.json](/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-complete-evaluation-20260914/fbip/timings.json) contient les résultats ultérieurs et référence l'empreinte de ce manifeste inchangé.

Cette expérience couvre des mises à jour de chemin et de champs au-delà de la seule rotation LL. Les gains temporels concernent ce programme fermé, ces dimensions et ce protocole. Elle ne démontre ni l'extraction automatique de ces preuves depuis le TAST, ni la sûreté avec des callbacks opaques, ni une gestion générale des panics, ni un gain sur le runner altbak. Le borrow exclusif conservé est une propriété du programme Rust fermé ici, pas une nouvelle analyse interprocédurale de Purust.
