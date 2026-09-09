# Conserver la preuve d'unicité dans RBTree

9 septembre 2026. Expérience isolée sur le Rust généré actuel, profil O1 et mimalloc. **Aucune variante intégrée au générateur :** le gain local reste trop faible et variable pour conclure ; le passage de l'emprunt au worker ralentit le noyau.

## Résultats temporels

Chaque processus construit 100 000 nœuds, calcule la profondeur 22 et détruit l'arbre. Une itération de chauffe, puis 15 mesures. Les variantes alternent leur ordre entre paires. Les compilations, comptages et vérifications sont terminés avant les mesures. Les nombres ci-dessous sont les médianes des échantillons, sans suppression des valeurs extrêmes.

| Expérience | Paires | Avant | Prototype | Écart |
| --- | ---: | ---: | ---: | ---: |
| Emprunt local, premier relevé | 3 | 19,699 ms | 19,358 ms | −1,7 % |
| Emprunt local, confirmation | 5 | 19,555 ms | 19,398 ms | −0,8 % |
| Emprunt transmis au worker privé | 3 | 19,282 ms | 20,778 ms | +7,8 % |

La confirmation locale donne trois paires favorables et deux défavorables. Les médianes par processus avant/après sont respectivement : 19,597/19,078 ; 19,968/19,398 ; 19,423/20,117 ; 19,573/19,266 ; 19,285/19,477 ms. Plage des échantillons : 18,524–23,003 ms avant, 18,290–20,751 ms après. Cette série n'établit pas un gain suffisamment stable pour justifier une intégration.

Le worker est plus lent dans les trois paires : 19,671/20,911 ; 19,242/20,787 ; 19,160/20,587 ms. Plages : 18,683–24,937 ms avant, 18,492–26,957 ms après. Les variantes sont indépendantes : l'essai worker ne contient pas la transformation locale ni la suppression des clones du précédent bébé.

Référence historique, relue dans le [README du checkout normal](../../../altbak.pub/README.md#rust) : **RBTree compilé 18,985 ms, natif optimisé 36,070 ms ; totaux 21,82 et 36,13 ms**. Nos mesures portent sur un noyau extrait, pas sur les 14 benchmarks du runner. Elles ne remplacent pas ces baselines et ne démontrent aucun nouveau gain global.

## Deux transformations en Rust sûr

**Locale :** dans `makeBlack`, `insert` et la branche de clé égale de `ins`, conserver le `Option<&mut Tree>` issu du premier `Rc::get_mut`. Extraire les champs via cet emprunt, puis écrire le nouveau payload via le même emprunt. Le fallback partagé/faible reste identique. Trois sites statiques sont modifiés ; la construction descendante du benchmark en exécute un 100 000 fois.

**Worker privé :** `ins` conserve son `Rc` et passe le `&mut Tree` à `balance__purust_reuse`. Celui-ci remplit l'emplacement et retourne `()`. Ses enfants restent des `Rc` possédés ; le propriétaire retourne sa cellule après la fin de l'emprunt. Les deux appels et les 82 sites de retour du worker sont adaptés. Un helper de remplissage écrit directement le payload. Le worker public et les chemins partagés/faibles restent identiques. Les cellules des rotations internes gardent leurs tests actuels.

Le compilateur Rust vérifie la durée de l'emprunt et interdit d'aliasser ou déplacer le propriétaire pendant son utilisation. Aucun `unsafe`, pointeur brut d'écriture, compteur supposé unique, ni changement de représentation n'est ajouté. Cela ne constitue pas une règle générale du générateur : ses contraintes de flux, d'effets et de libération seraient à formaliser avant toute extension à d'autres programmes.

## Opérations réellement comptées

Comptage séparé avec le wrapper autour de `std::rc::Rc` du [bébé précédent](../rust-perceus-counts-20260909/REPORT.md). Ce sont des appels logiques exécutés, pas des instructions machine.

| Construction de 100 000 nœuds | Avant | Local | Worker |
| --- | ---: | ---: | ---: |
| `get_mut` réussis | 4 967 864 | 4 867 864 | 2 783 888 |
| Retests supprimés | — | 100 000 | 2 183 976 |
| Allocations | 100 001 | 100 001 | 100 001 |
| Clones | 3 060 100 | 3 060 100 | 3 060 100 |
| Relâchements partagés pendant la construction | 2 960 100 | 2 960 100 | 2 960 100 |
| Libérations finales après profondeur | 100 001 | 100 001 | 100 001 |

Dans toutes les phases instrumentées, créations, clones, libérations, consommations et refus d'unicité pour partage ou référence faible restent identiques. Les bilans globaux des handles et des payloads sont vérifiés. Le worker ne supprime aucun test dans la construction qui conserve chaque ancienne racine : à ces appels, la racine est partagée.

## Code machine et validité

L'[extrait arm64](assembly-excerpt.txt) confirme que le helper original relit les deux compteurs et les compare à 1 ; LLVM O1 n'efface donc pas tous les retests. Le helper de remplissage du prototype reçoit directement l'adresse du payload et ne refait pas ces contrôles. Les tests locaux diminuent aussi le code optimisé du harnais (`main` : 43→41 sites de chargement, 57→55 branchements).

Le [résumé LLVM](assembly-summary.json) montre également que le changement de signature du worker modifie son appelant : `ins` passe de 83 à 86 chargements statiques, de 68 à 74 écritures et de 98 à 107 branchements. Ces comptes statiques ne prouvent pas la cause du ralentissement. Supprimer une catégorie d'opérations ne garantit pas que l'ensemble du code machine soit plus rapide.

Les trois variantes passent les vérifications instrumentées et les vérifications avec le **véritable `std::rc::Rc`** : quatre rotations, invariants rouge/noir et ordre, profondeur et cardinalité des 100 000 nœuds, clés exactes des 200 versions conservées, recoloration avec les quatre combinaisons de partage fort/faible, adresses conservées sur le chemin unique, durée de vie des enfants. Une séquence supplémentaire de 512 insertions mêle doublons, racines uniques, snapshots espacés et références faibles ; toutes les anciennes versions et leurs libérations sont vérifiées.

## Décision et suite

Ne pas intégrer ces prototypes pour l'instant. L'emprunt local reste une simplification possible si un autre changement le rend utile ; le worker ne justifie pas une généralisation sous cette forme.

Le prochain bébé proposé reste dans la réutilisation Perceus/FBIP : **recolorer uniquement le champ couleur d'une cellule unique**, au lieu d'extraire les enfants, d'écrire `E`, puis de reconstruire tout `T`. Mesurer cette spécialisation indépendamment ; le nombre d'allocations est déjà minimal et aucun gain temporel n'est présumé.

## Reproduction

Depuis ce dossier, avec le même Rust généré et les dépendances mimalloc du runner :

```sh
python3 probe.py count
python3 validate.py
python3 probe.py time --assembly --variants before local --pairs 3 --output timings-local.json
python3 probe.py time --assembly --variants before worker --pairs 3 --output timings-worker.json
python3 probe.py time --variants before local --pairs 5 --reuse-binaries --output timings-local-confirmation.json
python3 summarize.py
```

Exécuter ces commandes séquentiellement, sans autre compilation ou instrumentation pendant les mesures. Les sources natives, binaires, LLVM et assembleur sont régénérables dans `build/`, ignoré par Git. Les scripts réutilisent `probe.py`, `tracked_rc.rs` et `count_harness.rs` du bébé précédent ; leurs empreintes sont dans [metadata.json](metadata.json), avec celles du Rust généré, des binaires chronométrés, les révisions et la version de Rust. [Comptages](counts.json), [vérifications natives](validation.json), [mesures locales initiales](timings-local.json), [confirmation locale](timings-local-confirmation.json) et [mesures worker](timings-worker.json) conservent les résultats détaillés.

Seuls les artefacts de cette expérience et le todo Purust sont modifiés. Aucun changement du générateur, du Rust généré par le runner ou des checkouts normaux d'altbak.pub/PBO. Pas de compilation complète de la suite pour cette expérience non intégrée.
