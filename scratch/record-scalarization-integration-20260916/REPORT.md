# Scalarisation des records intégrée dans Purust — 16 septembre 2026

Les étapes 1 et 2 sont intégrées au backend : garder les champs d'un record
dans les paramètres scalaires d'une boucle, puis étendre cette transformation
à certains appels locaux connus. La campagne finale mesure **394,08 → 3,92 µs
sur Records**, et **−4,65 % sur le total du runner**, en comparaison appariée.
Les informations disponibles suffisent : aucune modification supplémentaire
de PBO ni du compilateur Haskell n'a été nécessaire pour ce périmètre.

## Transformation livrée

Sources dans `/Users/0x1/Documents/htdocs/purust/purust` :

- `src/Purust/RecordScalarization.purs` : reconnaissance des boucles,
  transport des champs dans des paramètres natifs, reconstruction finale.
- `src/Purust/RecordScalarCalls.purs` : résumés des dépendances et des champs
  modifiés par les fonctions locales connues, génération de workers scalaires.
- `src/Purust/CodeGen.purs` : branchement après PBO, avant la fusion des thunks
  et le renommage des variables ; déclaration des workers privés.
- `tests/codegen/record-scalarization.mjs` : tests de reconnaissance et
  compilation/exécution du Rust réellement généré sous Rc et Arc.
- `bin/purust.js` : bundle reconstruit.

La passe travaille sur l'IR final de PBO. Elle utilise les types de records,
les projections, les mises à jour et les corps de fonctions disponibles.
Elle ne transforme pas `usageCount` ou `escapes` en preuve d'unicité mémoire.

La fonction publique garde sa signature. Elle extrait les champs avant de
déplacer le record d'origine vers un worker privé. Ce record reste la base
immuable de reconstruction ; les itérations transportent les valeurs Int.
À la sortie, seules les branches contenant des champs écrits sont mises à
jour par les mécanismes COW existants. Un indicateur d'écriture préserve
le retour direct du record lorsque la boucle ne réalise aucune écriture.
Les anciennes versions partagées restent valides.

Pour un helper connu, le backend calcule les champs lus, les champs modifiés
et les dépendances scalaires. Il génère un worker retournant un Int par champ
modifié. Cela permet de conserver des appels scalaires réels avec l'IR actuel,
sans ajouter de type tuple ni de convention d'appel à PBO. Les fonctions
publiques manipulant les records restent disponibles. Seuls les workers
effectivement utilisés sont émis.

Cette décomposition peut dupliquer des calculs scalaires : elle est réservée
aux petits helpers purs et bornée par un budget de taille. Ce choix suffit
au cas mesuré ; il ne constitue pas une ABI générale de retour multivaleur.

## Mesure du runner officiel

Comparaison du binaire figé avant cette passe et d'un runner entièrement
régénéré avec le bundle final. Le témoin inclut déjà l'optimisation précédente
des décisions d'équilibrage de RBTree.

Protocole : 21 paires, ordre mélangé avec graine 916267, warm-up officiel,
meilleur de dix par test. Les 14 sorties sont validées à chaque invocation.
Les compilations sont terminées avant les mesures. Les deux manifests ont
le même profil release (`opt-level=3`, `debug=false`) ; seul le bundle Purust
diffère dans les empreintes des entrées de compilation.

| Mesure | Médiane avant | Médiane après | Variation appariée |
| --- | ---: | ---: | ---: |
| Records | 394,08 µs | 3,92 µs | −99,01 % |
| RBTree | 8 202,62 µs | 8 197,79 µs | −0,12 % |
| Total des 14 temps | 9,06451 ms | 8,66750 ms | **−4,65 %** |

Les pourcentages sont les médianes des ratios après/avant de chaque paire,
pas les ratios des deux médianes présentées. Le total utilise la somme des
temps affichés par test, avant l'arrondi du total imprimé par le runner.
Les mesures brutes sont dans `final-comparison/results.json`.

Le Rust de RBTree est identique octet pour octet entre les variantes ; sa
petite variation temporelle relève du bruit de mesure. Les fluctuations des
autres tests ne sont pas attribuées à la passe. Le gain reproductible vient
ici de Records. On ne doit pas additionner ce pourcentage aux pourcentages
de la campagne précédente comme s'ils partageaient le même dénominateur.

Repères historiques du `README.md` d'altbak.pub-purust : Records compilé
385,38 µs, Rust FP natif 98,75 µs, Rust impératif natif 4,25 µs ; RBTree
8 913,29 µs ; total 9,78 ms. Le nouveau Records se situe dans l'ordre de
grandeur du natif impératif historique. La preuve causale repose sur les
paires mesurées dans cette campagne, sans prétendre battre un natif qui
n'a pas été remesuré simultanément. L'ancien objectif de −50 % sur RBTree
n'est toujours pas atteint.

## Preuve à travers les appels

Le test construit une expression typée au niveau de l'IR après PBO et passe
par les véritables transformations et le générateur Rust. Le fixture fait
10 000 mises à jour de quatre champs, avec dépendances entre anciens et
nouveaux champs et conservation d'un alias de l'entrée.

La référence est une petite boucle Rust appelant le helper public généré,
qui prend et retourne un record. La variante utilise la boucle et les
workers Int générés par la nouvelle passe. Ce n'est donc pas une comparaison
complète de deux anciens/nouveaux pipelines à partir d'un programme source ;
c'est une expérience ciblée sur les appels du code réellement généré.

| Compilation du fixture | Référence | Variante générée | Variation appariée |
| --- | ---: | ---: | ---: |
| Inlining normal | 336,084 µs | 5,167 µs | −98,41 % |
| Helpers avec `inline(never)` | 349,292 µs | 21,542 µs | **−93,62 %** |

21 séries mélangées, trois warm-ups et dix échantillons par cas, médiane des
meilleurs temps ; 21 victoires sur 21 dans chaque mode. Rust `opt-level=3`,
LTO désactivé, MiMalloc. Les quatre champs de sortie sont observés et la
destruction du résultat est comprise dans le chronométrage ; les assertions
de correction sont ensuite vérifiées hors chronométrage.

Dans le second mode, les attributs sont appliqués dans le scratch au helper
public de référence et aux quatre workers scalaires. L'assembleur de
`ScalarRecords_throughCall` contient encore les quatre appels dans sa boucle.
La référence effectue un appel au helper record par itération et la variante
quatre appels Int : le nombre d'appels n'est donc pas identique, mais le gain
subsiste même sans leur inlining. Aucun record intermédiaire n'est nécessaire
pour transporter les champs entre ces appels.

Sources, commandes de compilation, empreintes et échantillons :
`generated-call-fixture.rs`, `measure-generated-calls.py`,
`generated-calls-results.json`, `generated-calls-assembly.txt`.
Ces chiffres ne sont pas une seconde mesure du gain total officiel.

## Validation et limites

Le build final réussit. La suite de génération compte **80 tests : 76 réussis,
4 échecs environnementaux** (`crypto-hash-ffi`, `datetime-instant-ffi`,
`foreign-object-foldm`, `uuid-ffi`). Leurs logs signalent l'accès Docker refusé.
Une inspection Docker séparée avec accès autorisé confirme que le conteneur
de référence `core-api-cli-1` est arrêté. Ces quatre tests restent non validés.

Les cas ciblés finaux compilent et s'exécutent sous Rc et Arc :

- dépendances simultanées entre champs, lets scalaires et helpers avec lets ;
- champs constants et updates conditionnels, y compris une itération sans
  écriture suivie ou précédée d'itérations qui écrivent ;
- zéro itération et zéro écriture : identité et zéro allocation conservées ;
- entrée unique : boucle de mises à jour sans allocation et identité conservée ;
- alias de l'ancienne racine et de ses enfants, champs conservés ;
- worker de 16 champs/19 paramètres compilé et exécuté ;
- collisions de noms et refus des captures, appels inconnus, rangées ouvertes,
  types contradictoires, labels dupliqués et récursion mutuelle.

Le périmètre accepté est volontairement limité : records fermés à feuilles
Int, jusqu'à 16 feuilles et huit niveaux pour les boucles ; récursion terminale
sur une fonction reconnue ; paramètres scalaires Int/Boolean. Les helpers
doivent être locaux, non récursifs, purs, saturés et disponibles dans le même
module ; ils sont bornés à 12 feuilles et à un budget de duplication de 256
nœuds. Les appels opaques et les formes non prouvées gardent leur génération
existante. Les workers conservent les opérateurs et la représentation Int
du backend ; cette passe ne redéfinit pas leur sémantique.

Les résumés sont dérivés après optimisation, là où ils décrivent le programme
réellement généré. Une extension aux appels entre modules pourrait justifier
des contrats exportés et leur validation/invalidation dans PBO. Aucune donnée
manquante imposant une modification Haskell n'a été rencontrée pour le scope
livré : les étapes 3 et 4 ne sont donc pas nécessaires ici.

## Reproduction et traçabilité

Depuis `/Users/0x1/Documents/htdocs/purust/purust` :

```sh
npm run build
node tests/codegen/record-scalarization.mjs
node --test --test-concurrency=4 tests/codegen/*.mjs
```

Depuis `/Users/0x1/Documents/htdocs/altbak.pub-purust`, relancer la comparaison
des binaires figés dans un nouveau dossier de résultats :

```sh
python3 scratch/compact-guards-integration-20260916/compare.py \
  --variant before scratch/record-scalarization-integration-20260916/before-benchmark \
  --variant scalar scratch/record-scalarization-integration-20260916/final/benchmark \
  --rounds 21 --seed 916267 \
  --output scratch/record-scalarization-integration-20260916/repeat
```

`final/manifest.json` et `before-benchmark-manifest.json` fixent les entrées
et le profil des runners. `final-source-hashes.json` fixe les sources et le
bundle livrés. `before-source-hashes.json` distingue les changements déjà
présents avant cette tâche ; `Main.purs` et `ChildBranchPrinter.purs` sont
inchangés pendant cette intégration. Les sources finales et le bundle sont
archivés dans `final-sources/`. Les logs de build, du runner officiel, de la
suite complète et du test ciblé sont conservés dans ce dossier.
