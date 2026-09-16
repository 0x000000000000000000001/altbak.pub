# Preuve expérimentale des possibilités TAST / FBIP — 16 septembre 2026

**Résultat : les champs actuels `usageCount + escapes` ne changent pas le
noyau RBTree produit. Une preuve d'exclusivité plus précise permet un gain
expérimental d'environ 3,2 %. Aucun des essais ne démontre −50 %.**

Le compilateur Haskell, PBO, Purust et les sources du benchmark n'ont pas été
modifiés pendant cette expérience. Les variantes, binaires et rapports sont
confinés à ce dossier scratch. Les corrections antérieures de production
restent en place. Les variantes futures sont préparées manuellement à partir
du Rust généré ; aucune nouvelle analyse Haskell n'est implémentée ici.

## 1. Informations présentes : usageCount et escapes

### Question et expérience

Quelle contribution ces deux champs ont-ils aujourd'hui sur RBTree ? Deux
générations utilisent le même bundle et les mêmes 302 TAST. Dans la seconde,
on enlève seulement les **26 658 paires** `usageCount` / `escapes`. Les caches
et sorties sont isolés ; les chemins source/FFI sont normalisés identiquement.

**Le fichier Rust RBTree est identique octet pour octet**, avec et sans les
champs, et identique au fichier du benchmark frais de la session précédente :

`f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a`

C'est une preuve de contribution directe nulle au code du noyau actuel,
plus précise qu'un chronométrage comparant deux programmes identiques. Ce
n'est pas une preuve d'inutilité générale de l'analyse d'usage. Sur l'ensemble
des sorties, 305/307 fichiers Rust sont identiques : Data.Enum change dans des
commentaires et Data.Array.intersperse dans une réassociation EffectBind.
Ne pas étendre le résultat RBTree à toutes les applications du backend.

### Pourquoi ce résultat

RBTree contient 135 annotations, mais seulement **cinq occurrences de variable**
avec `(usageCount=1, escapes=false)`, plus leurs cinq paramètres. Ce sont des
scrutinees de case. Les enfants dans `ins` et `depth` sont `(1,true)` :
`Usage.hs:112` considère conservativement qu'un argument d'appel échappe.
Le prédicat actuel ne sélectionne donc pas les clones chauds des enfants.

Purust calcule déjà les déplacements depuis la liveness du programme optimisé
(`CodeGen.purs:2411`). Les hints source sont invalidés par PBO puis nettoyés
par `withoutSourceUsage` (`CodeGen.purs:98`). Cela empêche un compteur source
de contourner les usages effectivement créés par les transformations.

**Dernière utilisation d'une variable ne signifie pas allocation unique.**
Un paramètre de `depth` peut n'être lu qu'une fois dans un case alors que
l'appelant garde un alias de son arbre. Le booléen `escapes` actuel ne décrit
ni les alias préexistants, ni le partage entre enfants.

### Ce que fait déjà le code actuel

Sur les 100 000 insertions décroissantes, les compteurs séparés confirment :

- 100 001 allocations : une cellule E initiale et un nœud par clé ;
- 2 483 932 appels à `get_mut`, tous réussis ;
- 99 978 rotations sur place ;
- aucune allocation supplémentaire causée par un échec d'unicité sur ce cas.

Il n'y a donc pas de taux de réemploi restant à faire passer à 100 % sur ce
workload. L'hypothèse initiale « enlever des clones pour éviter les allocations
de reconstruction » ne correspond pas au coût actuel de ce benchmark.

Preuves détaillées : [ablation](current/REPORT.md), [inventaire](current/metadata.json),
[hashes et commandes](current/ablation-results.json), [script](current/ablate_usage.py).

## 2. Informations supplémentaires et transformations expérimentales

### Protocole commun

Le noyau de référence est extrait du Rust réellement généré, en retirant
uniquement les adaptateurs applicatifs sans rapport avec le chemin testé.
Toutes les variantes chronométrées utilisent le Rc natif, sans compteurs,
ou le Box natif pour l'essai indiqué. Rustc 1.96.0, édition 2021, niveau
d'optimisation 3, une unité de génération de code, mêmes flags par campagne.

Chaque exécution construit 100 000 clés décroissantes, calcule la profondeur
et **détruit complètement l'arbre dans l'intervalle mesuré**. Les entrées,
l'arbre construit et le résultat traversent `black_box`. Résultat vérifié : 22.
Trois échauffements puis dix mesures par processus ; 21 tours, ordre des
variantes mélangé avec graine fixe, aucun benchmark ni build de notre tâche
en parallèle. Les compilations précèdent les mesures.

Les temps du tableau sont la médiane des 21 meilleurs-de-10. La variation est
la médiane des variations appariées au témoin **du même tour** ; elle peut
légèrement différer du ratio des deux médianes. Une variation positive du
temps est une régression. Les deux campagnes ont chacune leur témoin.

### Campagne A : emprunts, contrôles et représentation

| Variante | Temps médian | Variation du temps | Observation |
| --- | ---: | ---: | --- |
| Rust généré, témoin A | 9 335,58 µs | référence | noyau actuel |
| Parcours `depth` emprunté | 9 255,67 µs | −0,26 % | gain non établi |
| Insertion par emplacement emprunté | 11 257,04 µs | +21,91 % | régression |
| Emprunt insertion + depth | 11 456,75 µs | +22,93 % | régression |
| Unicité supposée des nœuds non vides, garde E conservée | 9 242,33 µs | −0,50 % | gain non établi |
| Autre génération manuelle, Rc | 12 926,54 µs | +39,36 % | plusieurs transformations combinées |
| Même génération manuelle, Box | 13 470,46 µs | +43,28 % | Box 4,43 % plus lent que son témoin Rc |

Les compteurs logiques expliquent le travail enlevé, mais ne sont pas des
compteurs d'instructions machine après optimisation LLVM :

| Opérations | Généré | Depth emprunté | Insertion empruntée | Deux emprunts |
| --- | ---: | ---: | ---: | ---: |
| Clones pendant construction | 2 283 976 | 2 283 976 | 200 000 | 200 000 |
| Clones pendant depth | 200 000 | 0 | 200 000 | 0 |
| Allocations construction | 100 001 | 100 001 | 100 001 | 100 001 |

Enlever ces opérations source ne garantit donc pas un meilleur code machine.
Le coût exact de la régression des workers n'a pas été profilé ; on ne l'attribue
pas arbitrairement à un seul mécanisme. Le prototype d'insertion conserve les
gardes et permutations générées, mais son ABI et ses emprunts changent.

### Campagne B : deux vérifications ciblées

La première variante d'unicité conservait une garde sur E. Une information
plus précise sur la **provenance de chaque cellule mutable** permet aussi
de retirer cette garde : même les cellules temporairement vidées en E sont
uniques ; les feuilles E partagées n'atteignent pas ces sites dans les cas
validés. La variante remplace les mêmes 103 sites textuels `get_mut`, sans
changer rotations, allocations, layout ou parcours.

| Variante | Temps médian | Variation du temps |
| --- | ---: | ---: |
| Rust généré, témoin B | 9 033,96 µs | référence |
| Insertion empruntée initiale | 10 956,96 µs | +21,32 % |
| Insertion empruntée : test de couleur avancé et adaptateur inliné | 10 304,46 µs | +12,85 % |
| Unicité supposée des nœuds non vides, garde E conservée | 9 066,71 µs | −0,18 % |
| Exclusivité au site mutable, sans garde E | 8 792,17 µs | **−3,16 %** |

La dernière variante gagne dans **19 des 21 tours**. Bootstrap descriptif de
la médiane des gains appariés : intervalle indicatif 95 % **[2,59 ; 4,40] %**.
Les intervalles correspondants de `depth` et de la première variante d'unicité
traversent zéro. Ce sont des observations sur cette machine et cette session,
pas une garantie universelle ni une borne maximale des gains futurs.

### Correction et limites des preuves

- Les variantes empruntées vérifient ordre et clés exactes, invariants rouge/noir,
  quatre rotations, ordres croissant/décroissant/mélangé, doublons, références
  faibles, 200 versions persistantes et bilans de références équilibrés.
- Le fallback partagé du prototype d'insertion est conservateur : 7 311 clones
  contre 5 808 sur le test de 200 versions, mêmes 1 504 allocations. Le petit
  test Weak ajoute deux allocations. Aucun gain partagé n'est revendiqué.
- Les variantes d'unicité contiennent un accès unsafe confiné au scratch. Les
  versions de validation assertent `strong_count==1` et `weak_count==0` à chaque
  accès spécialisé, sur rotations, 100k clés croissantes/décroissantes/mélangées
  et doublons. La variante finale le vérifie aussi sur les E transitoires.
  Une API générale recevant des snapshots doit conserver son fallback.
- Les variantes Rc/Box manuelles sont entièrement safe et ont les mêmes corps
  d'algorithme, sauf l'alias du propriétaire et la fonction d'accès mutable.
  Leur structure complète, couleurs et clés sont comparées au généré. Rc
  préserve 200 versions ; Box suppose une région sans partage. Leur différence
  avec Purust combine représentation des feuilles, emprunts et rotations :
  elle n'est pas imputable à une seule métadonnée.

**Les assertions vérifient les exécutions testées, pas tous les programmes.**
Les chronométrages montrent la valeur possible d'un contrat statique valide.
Ils ne constituent pas l'implémentation ni la preuve de correction d'une
nouvelle passe du compilateur.

## 3. Qu'ajouter réellement au TAST Haskell ?

Les types profonds, `dataDecls`, `classDecls` et `TypeApp` existent déjà. Les
informations manquantes ici sont des propriétés de flux, d'alias et de
consommation. Elles pourraient aussi être inférées dans PBO ; les exporter
depuis Haskell permettrait de les partager entre backends.

| Contrat proposé, inexistant actuellement | Analyse Haskell envisageable | Transformation essayée et résultat |
| --- | --- | --- |
| `parameterEffects`: lecture seule, pas de rétention, pas de retour d'alias | suivre usages et projections, résumés transitifs des appels, point fixe sur fonctions récursives | depth par emprunt : aucun gain établi sur RBTree |
| `consumesParameter` + origine du résultat + retour dans le même emplacement | suivre un jeton d'ownership à travers cases, appels et constructeurs | insertion par emplacement : régression dans les prototypes testés |
| `resultOrigins` et précondition `preservesDisjointOwnership` | distinguer allocation fraîche, alias d'argument, projection ; séparer les jetons des champs lors d'un case ; conserver leur provenance au réemploi | unicité au site mutable : environ 3,2 % mesuré |
| Exclusivité d'une région complète et absence d'alias externe/Weak/FFI retenant | propager fraîcheur et non-partage dans un groupe fermé de fonctions, avec spécialisation conditionnelle des appels | représentation Box : aucun gain dans cette génération manuelle |

Le premier contrat utile pour `insert` doit être conditionnel : « si l'arbre
entrant est exclusivement possédé avec enfants disjoints, la fonction préserve
cette propriété ». `depth` a un contrat de lecture. Ce ne sont pas deux booléens
locaux sur une occurrence ; il faut relier paramètres, champs et résultats.

Point d'intégration existant : `purescript/src/Language/PureScript/Make/Actions.hs:256`
appelle `Usage.computeUsage`. Une passe de résumés de fonctions pourrait y être
ajoutée. `CoreFn/Ann.hs:42` porte les annotations et `CoreFn/ToJSON.hs:135` les
sérialise. Les résumés peuvent vivre à la racine du module, référencés par des
identifiants de liaison stables, afin d'éviter un nouveau wrapper de syntaxe à
chaque nœud. `simplifyType` doit continuer à traiter les types : l'unicité
demande une analyse des expressions et des appels.

Les analyses doivent avoir un état inconnu conservateur, traiter les appels
étrangers/inconnus comme retenant potentiellement leurs arguments et résoudre
les groupes récursifs par point fixe. Chaque fait doit préciser sa phase de
validité. Après inlining, duplication, extraction de closure ou ajout de clones,
PBO/Purust doivent préserver formellement, recalculer ou invalider la propriété.
Une annotation Haskell périmée n'autorise jamais un accès mutable unchecked.

Les identifiants stables, la multiplicité des closures et des résumés d'appels
plus précis sont des extensions plausibles, mais **aucun gain chiffré ne leur
est attribué dans ce rapport**, faute d'expérience correspondante.

## 4. Comparaison avec la baseline officielle

Le README d'altbak.pub, ligne 145, documente **8 788,42 µs**. La cible −50 %
est donc **4 394,21 µs**. La session précédente obtenait une médiane officielle
de 8 536,29 µs. Trois nouvelles exécutions du runner officiel inchangé donnent
8 204,83 ; 8 374,71 ; 8 139,83 µs, médiane **8 204,83 µs**. Elles vérifient
toutes le résultat 22 ; logs dans `official-checks/`.

Ces variations de baseline ne sont pas un gain du TAST : aucun changement de
production n'a été fait ici. Le harness isolé a ses propres conditions de
compilation et frontières d'appel ; ses temps absolus ne sont pas substitués
aux chiffres officiels. Les gains expérimentaux sont calculés seulement contre
leur témoin apparié. L'objectif de 50 % n'est démontré par aucune comparaison.

L'étape justifiée est de considérer le contrat d'exclusivité au site mutable
comme une piste de quelques pour cent sur ce test. Une ambition plus grande
demande une autre expérience sur le code machine, la navigation ou la
représentation ; les annotations seules ne fournissent pas cette preuve.

## Reproduction et données

Depuis ce dossier :

```sh
python3 current/ablate_usage.py
python3 borrow/prepare.py
python3 borrow/prepare-ins.py
python3 borrow/check-orders.py
python3 unique/prepare.py
python3 measure.py --with-owned
python3 measure.py --final-checks
python3 analyze.py
```

Les kernels et tous les contrôles supplémentaires sont conservés dans
`borrow/`, `unique/` et `unique/owned/`, avec leur README. `measure.py` compile
les kernels existants avant toute mesure. Il ne lance aucune mesure
instrumentée. `timing/results.json` et `timing-checks/results.json` contiennent
les 2 520 temps individuels mesurés, l'ordre des essais, flags, version rustc,
hashes des kernels/harness/binaires et les agrégats. `bootstrap.json` complète
l'analyse statistique descriptive. Les fichiers `timing*.log` conservent
l'avancement et les résultats des deux campagnes.
