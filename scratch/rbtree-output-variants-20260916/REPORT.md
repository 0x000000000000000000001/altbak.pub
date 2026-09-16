# Variantes du Rust compilé dans altbak.pub-purust/output

Expérience du 16 septembre 2026. **Meilleur résultat répété : environ 8 à 9 %
de temps en moins**, avec décisions d'équilibrage compactes et unicité supposée
aux sites mutables. **En gardant tous les contrôles Rc et la persistance :
environ 4 à 5 %.** Les variantes de stockage en région et en arène essayées
sont plus lentes. Aucun gain proche de 50 % n'est démontré.

## Source et périmètre

Source demandée, lue directement sans régénération :

`/Users/0x1/Documents/htdocs/altbak.pub-purust/output/purust_output/Purs_Test_RBTree/src/lib.rs`

Le lien `output` pointe vers `run/bak/rust/output`. SHA-256 initial et final :

`f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a`

`original-output.rs` conserve ce fichier complet. `baseline.rs` conserve les
types et toutes les fonctions du noyau à l'identique ; seuls les adaptateurs
de module/Effect sans rôle dans le noyau sont retirés pour compiler sans
l'application entière. L'extraction est reproductible avec `prepare.py` ;
les chemins et hashes sont dans `source-manifest.json`.

Aucun changement à l'output original, aux sources PureScript, au compilateur,
à PBO ou à Purust. Le seul ajout au dépôt est ce dossier scratch. Ce sont des
modifications expérimentales du Rust compilé, pas des passes de compilation
implémentées. Les essais à représentation différente sont explicitement
distingués des transformations mécaniques du code généré.

## Protocole

- 100 000 clés décroissantes, construction complète, profondeur, puis destruction.
- Entrées, arbre construit et résultat passent par `black_box` ; résultat 22 vérifié.
- Allocation initiale et libération sont dans le chrono. Chaque région/bloc/Vec
  est recréé à chaque appel ; aucun buffer de notre prototype n'est conservé
  hors chrono pour la prochaine itération. Les caches propres à l'allocateur
  restent naturellement actifs, comme dans le programme initial.
- Rustc 1.96.0, édition 2021, `opt-level=3`, `codegen-units=1` identiques par campagne.
- Trois échauffements puis dix échantillons par processus ; 21 tours avec ordre
  des variantes mélangé et graine fixe. Compilations et validations finies avant
  chaque campagne ; aucune mesure de notre tâche ne tourne en parallèle.
- Temps présenté : médiane des 21 meilleurs-de-10. Réduction : médiane des ratios
  appariés au témoin du même tour, donc légèrement différente du ratio des médianes.
  Une variation positive du temps dans les tableaux est une régression.

Une première campagne utilise l'allocateur Rust système. La lecture du `main.rs`
de l'output établit que **l'application utilise mimalloc**. Les conclusions
principales reposent donc sur deux campagnes supplémentaires avec **le même
mimalloc compilé que l'application**, lié depuis son artefact existant. Les
hashes de cette bibliothèque, des kernels, harness et binaires sont conservés
dans chaque manifest. Les résultats système restent disponibles comme expérience
exploratoire ; ils ne sont pas présentés comme résultats de l'allocateur officiel.

Les tableaux donnent le coût du noyau isolé, avec ses propres frontières d'appel
et sa compilation. Ils ne prétendent pas être une nouvelle exécution complète
du runner officiel ou une modification intégrée du backend.

## Résultats principaux avec mimalloc

| Variante | Temps médian | Variation du temps vs témoin apparié | Tours gagnés |
| --- | ---: | ---: | ---: |
| Généré inchangé | 8 404,83 µs | référence | — |
| Garde remplacée par classification compacte | 8 027,67 µs | −4,59 % | 20/21 |
| Classification partagée avec la permutation | 8 002,92 µs | −5,07 % | 20/21 |
| Classification partagée + unicité au site mutable | 7 655,17 µs | **−9,09 %** | 21/21 |
| Bloc régional, même Rc, curseur atomique | 10 507,75 µs | +25,63 % | 0/21 |
| Même bloc, curseur sans atomiques | 10 576,29 µs | +25,64 % | 0/21 |
| Bloc sans atomiques + classification partagée | 10 195,29 µs | +21,81 % | 0/21 |
| Bloc sans atomiques + classification + unicité | 9 429,67 µs | +12,41 % | 0/21 |
| Arène Vec, indices u32, capacité réservée | 12 667,17 µs | +50,99 % | 0/21 |
| Même arène, croissance amortie sans réservation | 12 484,46 µs | +48,82 % | 0/21 |
| Arène, indices usize | 13 078,58 µs | +55,66 % | 0/21 |
| Arène u32 sans vérifications des indices | 11 763,12 µs | +41,58 % | 0/21 |
| Même arène, équilibrage/rotations forcés inline | 11 662,58 µs | +39,30 % | 0/21 |

Une campagne distincte vérifie les gains positifs et teste une dernière réduction
du travail d'équilibrage : ne regarder que l'enfant modifié par l'insertion.

| Variante | Temps médian | Variation du temps vs témoin apparié | Tours gagnés |
| --- | ---: | ---: | ---: |
| Généré inchangé, témoin de cette campagne | 8 300,88 µs | référence | — |
| Classification partagée | 7 978,00 µs | −3,99 % | 19/21 |
| Classification partagée + unicité | 7 660,88 µs | **−8,01 %** | 21/21 |
| Arène u32 de référence | 12 627,33 µs | +51,77 % | 0/21 |
| Arène u32, équilibrage du côté modifié | 11 688,62 µs | +40,66 % | 0/21 |
| Arène sans vérifications, équilibrage inline | 11 771,04 µs | +42,54 % | 0/21 |
| Même arène, équilibrage du côté modifié | 11 175,96 µs | +35,64 % | 0/21 |

L'équilibrage directionnel améliore son propre témoin d'arène, mais ne rejoint
pas le Rust généré. La suppression des vérifications d'indices et l'inlining
améliorent également certaines variantes sans les rendre compétitives.

La différence entre garde compacte et classification partagée est petite ;
aucune supériorité générale de la seconde sur la première n'est revendiquée.
Les plages/échantillons individuels et le bootstrap descriptif des gains sont
dans les JSON. Les statistiques décrivent cette session sur cette machine ;
elles ne sont pas des garanties de gains universels.

## Ce qui a effectivement changé

### Décisions d'équilibrage : gain sans nouvelle preuve TAST

Le code original produit une longue expression de garde et reconnaît ensuite
à nouveau le cas LL pour la permutation sur place. Les variantes utilisent
une classification `aucun / LL / LR / RL / RR`, puis éventuellement partagent
ce résultat entre garde et permutation. Les autres rotations conservent le
fallback généré. Allocations, Rc, signatures, descente, parcours et destruction
restent identiques. Aucun unsafe n'est introduit.

Les informations nécessaires existent déjà dans les motifs et la dominance
des branches. Ce gain peut donc orienter une amélioration du générateur, sans
enrichissement préalable du TAST. Si des faits de motifs sont transportés,
ils doivent rester liés aux champs non modifiés qui ont été inspectés.

### Unicité au site mutable : petit gain supplémentaire conditionnel

Le meilleur prototype compose la classification avec le remplacement de 103
sites textuels `get_mut`. Il suppose que chaque cellule atteinte par une
mutation est exclusive, y compris les cellules temporairement vidées en E.
Les E partagés ne doivent pas atteindre ces accès. Une version de validation
vérifie forte=1 et faible=0 à chaque accès, notamment 2 483 932 accès sur le cas
décroissant. L'algorithme, le layout et les allocations ne changent pas.

Cela requiert une preuve de provenance/exclusivité au site, et sa préservation
par les appels et les rotations. `usageCount + escapes` ne la fournit pas.
Le prototype unsafe est confiné au scratch ; la variante générale recevant
des arbres partagés ou des Weak doit conserver ses contrôles/fallbacks.

### Région d'allocation : essai causal négatif

Le noyau Rc reste inchangé. Pendant son exécution, l'allocateur place les
allocations dans un bloc neuf de 6,4 Mo pour n=100 000, puis libère ce bloc
après les drops Rc ordinaires. 4 800 048 octets sont effectivement utilisés,
sans débordement sur ce cas. Le débordement garde un fallback vers l'allocateur
amont. Le curseur sans atomiques élimine un coût propre au premier prototype,
sans rétablir un gain.

Le fait TAST pertinent serait **aucune allocation ne survit à la région**,
avec un contrat d'exécution fermé. Cette optimisation n'a pas besoin de prouver
chaque cellule unique ; les partages internes peuvent vivre et mourir dans la
région. Le prototype global exige explicitement un contexte mono-thread sans
allocation concurrente ou échappante. Une application générale nécessiterait
un mécanisme de région adapté à ces autres contextes.

### Arène par indices : mémoire plus compacte, temps moins bon dans ces essais

Les liens deviennent des indices dans un Vec. E est l'indice zéro. Les rotations
conservent exactement les formes et couleurs de l'algorithme source. Cela
combine stockage régional, représentation, accès mutables et code de sélection
des rotations : le résultat ne s'attribue pas à une seule métadonnée.

La taille du nœud est 24 octets avec u32, 32 avec usize ; ce sont des tailles de
layout, pas une mesure du RSS. La variante u32 réservée demande une capacité
de 100 001 nœuds. La croissance amortie montre que connaître cette borne ne
constitue pas, ici, une source de gros gain temporel.

Les obligations supplémentaires seraient : région fermée, absence de versions
anciennes nécessaires aux écritures, absence d'observation des adresses, durée
de vie compatible et destruction sans effet observable. Pour éliminer les
vérifications d'indices, le backend doit établir l'invariant d'appartenance des
indices à l'arène. Cet invariant peut être assuré par sa propre construction ;
il n'exige pas nécessairement une nouvelle annotation Haskell.

Ces workers opèrent pour **n >= 0**. Les indices u32 demandent une borne ou un
fallback vers une représentation plus large ; le prototype garde un contrôle
de dépassement. Ces limites doivent figurer dans une spécialisation réelle,
pas être extrapolées à tous les Int PureScript.

### Résumé de validité pour l'équilibrage directionnel

Une insertion dans un arbre valide ne peut créer un besoin d'équilibrage dans
l'enfant laissé intact. On spécialise donc l'analyse vers le seul enfant modifié.
Les rotations, leurs priorités et les couleurs restent identiques.

Le contrat dépasse l'absence d'échappement : il faut savoir que l'arbre entrant
satisfait l'invariant récursif avant l'insertion et que les autres transformations
le préservent. Un résumé interprocédural de formes et de postconditions pourrait
exprimer ce fait ; son inférence automatique serait nettement plus ambitieuse
que le comptage d'occurrences. Aucun ordre particulier des clés n'est supposé.

## Validation et portée

Les checks comparent au généré original les **clés, couleurs, constructeurs
vides, forme complète et profondeur**, avec vérification indépendante de l'ordre
BST, de la racine noire, des hauteurs noires et de l'absence d'arêtes rouge/rouge.
Cas : quatre rotations minimales, vide lorsque pertinent, 100k insertions
croissantes/décroissantes, permutations déterministes et doublons.

Les décisions compactes sont aussi comparées exhaustivement aux gardes originales
sur 722 combinaisons de formes/couleurs, y compris hors invariant rouge/noir.
Elles préservent 200 snapshots, les Weak racine et enfant. Les validations de
région incluent également snapshots internes, Weak, alignements 1..4096,
débordements et réinitialisations, y compris avec mimalloc.

Les variantes d'unicité n'acceptent pas les snapshots/Weak externes. Les arènes
spécialisent le calcul fermé et ne remplacent pas une API publique persistante.
Leur version sans vérifications d'indices a été exécutée aussi avec assertions
à chaque accès. Les assertions prouvent les exécutions vérifiées, pas tous les
programmes futurs ni une passe Haskell encore inexistante.

## Baselines et conclusion pratique

Le README de **ce dépôt altbak.pub-purust**, ligne 154, documente **8 913,29 µs**.
La référence historique de la discussion précédente, dans altbak.pub, était
**8 788,42 µs**. Les témoins mimalloc de cette expérience sont respectivement
8 404,83 et 8 300,88 µs. Les gains annoncés sont toujours calculés contre le
témoin contemporain de la même campagne, jamais contre un chiffre historique
plus favorable. Le seuil de −50 % du README de ce dépôt serait 4 456,65 µs ;
aucune variante ne l'approche.

L'amélioration justifiée la plus directe est la génération de décisions compactes
et partagées : quelques pour cent mesurés, sans nouvelle preuve d'ownership.
La composition avec une preuve d'unicité porte les essais vers 8–9 %. Les pistes
mémoire et les invariants d'équilibrage restent des expériences négatives dans
les implémentations présentes. Cela ne prouve pas que toute autre arène ou
génération serait lente, mais ne permet pas de promettre de gros gains en ajoutant
simplement ces champs au TAST. Le coût exact de leurs régressions n'a pas été
profilé et n'est pas attribué sans preuve à une cause unique.

## Reproduire

Les kernels sont déjà conservés et prêts à compiler. Depuis ce dossier :

```sh
python3 measure.py --allocator mimalloc --out timing-mimalloc --variants generated,guard_simple,guard_fused,guard_unique,region,region_local,region_local_fused,region_local_unique,arena_u32,arena_grow,arena_usize,arena_unchecked,arena_unchecked_inline
python3 measure.py --allocator mimalloc --out timing-directional --variants generated,guard_fused,guard_unique,arena_u32,arena_directional,arena_unchecked_inline,arena_unchecked_directional_inline
python3 analyze.py
```

Le script lie l'artefact mimalloc existant du mode pure de ce dépôt ; son chemin
et son hash sont dans les manifests. Il compile toutes les variantes avant de
mesurer. `--compile-only` et `--skip-compile` permettent de séparer ces phases.
Les scripts de préparation et checks sont dans `decisions/`, `region/`, `arena/`.
`prepare_inline.py` et `arena/prepare_directional.py` conservent les transformations
supplémentaires. Les README associés explicitent leurs contrats.

Données brutes : `timing/results.json` (système), `timing-mimalloc/results.json`,
`timing-directional/results.json`. Au total, **5 880 échantillons chronométrés**,
hors échauffements, dans les trois campagnes. `statistics.json` conserve le
bootstrap descriptif reproductible. Les logs de correction sont séparés des
binaires non instrumentés utilisés pour les temps.
