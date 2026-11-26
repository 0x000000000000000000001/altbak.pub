# Go : callbacks natifs pour les effets entiers — 26 novembre 2026

Les sept petites actions passent de **6,0–7,1 ns à 3,6–4,7 ns**, soit
**31 à 43 % de temps en moins** dans la paire avant/après. Le compilateur
ajoute une entrée native aux actions `Effect Int` déjà mises en cache et
la sélectionne lors de leur adaptation en callback Go. Le chemin générique
reste disponible. Benchmarks, harnais, FFI source et références C/JS sont
inchangés ; la frontière opaque reste exécutée à chaque invocation.

Les [données et preuves](2026-11-26-go-native-callbacks.json) conservent la
campagne, les changements du compilateur, les tests, les prototypes et les
contrôles d’entrées variables.

## Résultats

Temps primaires App14 avant → après, puis ratios contre les mêmes références C :

- Church : **6,718 → 4,522 ns** ; Go/C **6,014 → 4,048×**.
- State : **6,337 → 4,028 ns** ; Go/C **5,661 → 3,598×**.
- Lazy : **6,236 → 4,035 ns** ; Go/C **5,561 → 3,598×**.
- Polymorphism : **7,148 → 4,061 ns** ; Go/C **6,349 → 3,607×**.
- RowToList : **6,039 → 3,555 ns** ; Go/C **5,395 → 3,176×**.
- List Processing : **6,906 → 4,735 ns** ; Go/C **5,673 → 3,890×**.
- Array Processing : **6,777 → 4,680 ns** ; Go/C **4,794 → 3,311×**.

La paire AppAll37 confirme les gains sur ces sept cas : **32 à 40 %**.
Le diagnostic d’allocations trouve les mêmes comptes avant/après, sur la
conversion FFI réelle et le repli forcé : zéro pour douze actions, une pour
Primes et 100000 pour RBTree. Les 56 résultats vérifiés sont corrects.

Les quatre indicateurs Go/C, sur les mêmes **14 cas communs**, donnent :

- Ratio des totaux : **1,110360× → 1,096290×**.
- Moyenne arithmétique μ : **3,579933× → 2,546327×**.
- Médiane Q₂ : **4,770374× → 3,243306×**.
- Moyenne géométrique G : **1,881685× → 1,477750×**.

Total Go : **9,725522 → 9,602283 ms**, contre **8,758890 ms en C**.
Le calcul utilise les temps exacts, avant arrondi : ΣGo/ΣC, puis moyenne,
médiane et moyenne géométrique des quatorze ratios individuels.

La [baseline officielle précédente](2026-11-26-go-ffi-calls.md) était
**1,203742× / 3,938655× / 5,454538× / 2,066376×**,
avec un total Go de **10,543442 ms**. L’ancien binaire remesuré est lui aussi
plus rapide dans cette nouvelle campagne. Les gains annoncés reposent donc
sur la paire fraîche, et non sur cette différence entre campagnes.

Le total reste dominé par RBTree : sa baisse de **123,355 µs** explique
l’essentiel de celle du total (**123,239 µs**). Les plages des totaux par
processus se recouvrent : **9,649–9,787 ms avant**, **9,587–9,668 ms après**.
Cette variation ne démontre pas un gain de RBTree causé par les callbacks.
Les progrès établis sur les petites actions améliorent surtout μ, Q₂ et G.
L’objectif de 1× sur tous les indicateurs reste ouvert.

## Transformation

`ModuleBindings` sélectionne uniquement des actions déjà initialisées par
un getter de cache, avec le TAST nominal `Effect Int` et une syntaxe que le
compilateur diffère déjà. Les alias, appels de factories, `EffectDefer` et
actions créées localement gardent leur représentation habituelle.

`EffectExprs` traduit le corps une seule fois. Il conserve son entrée Value
originale et ajoute un `func() int64` si le retour est nativement `int64`, ou
s’il est exactement emballé par le constructeur pur `Int(value)`. Les autres
valeurs génériques restent exclues : lire directement leur champ entier
pourrait différer de la conversion générale `Unbox`.

Les deux corps ont les mêmes statements et calculs ; seul le retour natif
retire cet emballage. Une invocation exécute un seul corps. La métadonnée
`WithFunctionData` est initialisée avec le getter, sans exécuter l’action.
Le prototype observe une allocation supplémentaire à cette initialisation,
qui ne se répète pas lors des appels. Aucune métadonnée n’est ajoutée aux
actions construites dans une boucle.

L’adaptateur FFI capture l’action une fois, puis recherche exactement
`FunctionData[func() int64]` pour un callback `Effect Int` attendu comme
`func() int` ou `func() int64`. La conversion vers `int` reste explicite.
Sans métadonnée correspondante, le callback conserve `Apply` puis `Unbox`.

Dans le runtime, `Apply` déballe directement une fonction unaire portant
une métadonnée et rejoint sa branche ordinaire. Les autres arités et les
cas nil conservent le comportement précédent. Aucun test supplémentaire
n’est ajouté au chemin des fonctions ordinaires.

Cette correction du runtime vient du prototype : une première variante
ralentissait le repli générique de 27–32 %. Conserver le corps Value original
supprime un appel intermédiaire ; le déballage direct supprime la récursion
restante. Dans la sonde finale, les fonctions ordinaires restent à ±0,4 %,
et le chemin natif conserve son gain. Ces prototypes sont séparés des mesures
officielles et ne remplacent aucun de leurs échantillons.

## Vérifications et limite observée

- **51 tests ciblés** passent : six nouveaux tests, 39 tests FFI et six tests
  runtime, dont les vérifications avec détecteur de courses.
- La fixture source passe sur les hôtes JS et natif, en Release et Dev :
  valeurs, effets différés, captures retenues, actions locales, alias,
  rejouabilité et chemins de repli.
- Une différence d’ordre de construction des factories dans `traverse_`
  existe entre JS et Go. L’ancien compilateur gelé produit exactement la
  même sortie Go que le nouveau. Valeurs et ordre des effets à l’intérieur
  des thunks concordent avec JS ; cette différence préexistante est conservée
  explicitement dans la preuve et dans les assertions de la fixture.
- **88 entrées × trois chemins × deux exécutables = 528 contrôles** concordent
  avec le C inchangé. Chaque chemin impose un appel opaque par action.
  Les métadonnées natives sont présentes sur **0/14 actions avant** et
  **14/14 après** ; le repli forcé les retire bien.
- Les **306 résultats** des douze processus et leurs batches calibrés passent.
  Les sources, références et exécutables sont protégés par leurs empreintes.
- La revue exhaustive explique tous les changements générés : **17 getters**
  dans chaque build conservent leur corps Value et ajoutent sa copie native ;
  seuls deux adaptateurs de `Bench_MeasureBatch` changent. Le runtime correspond
  au patch éprouvé. Autres déclarations, fichiers, imports et directives sont
  inchangés. La source FFI du harnais et son algorithme de mesure ne changent pas.

Une hausse dépasse 5 % dans AppAll : `record/set-large-1000`,
**0,371126 → 0,411767 µs** (+10,95 %). Son code de calcul est inchangé.
Une contre-mesure isolée, avec les getters réels et le même `Bench.runBench`,
donne **0,403505 → 0,401560 µs** (−0,48 %), avec six résultats corrects.
La hausse n’est pas reproduite dans ce contrôle ; cette observation ne prouve
pas l’absence de toute régression. Les données principales restent intactes,
y compris la hausse publiée ; le diagnostic ne les remplace pas.

## Protocole et reproduction

Apple M4 Pro, 48 Gio, macOS 26.1, Go 1.27.0 darwin/arm64, Node 24.8.0.
Baseline compilateur gelée à `25768d7`, incluant l’optimisation des conversions
FFI ; quatre fichiers de production modifiés. Le fork PureScript `41fb570`
produit le TAST. Le bootstrap natif vérifie 550 modules et 342182 types.
Il conserve le seul avertissement `ShadowedName` préexistant dans
`PrivateSTRegions.purs:266`. Le contrôle documentaire passe.

Release, GOGC=800, GOMAXPROCS par défaut, sans PGO. Builds séparés et terminés
avant la collecte principale, ainsi que les tests, prototypes et revues CPU.
Le diagnostic isolé ultérieur a sa propre séquence sans build concurrent. Trois processus
par variante, en ordre AB/BA/AB, pour App14 puis AppAll37. Échauffements,
calibration et minimum de dix batches du harnais conservés ; statistique
publiée : médiane des trois minima de processus. Compilation, démarrage et
sortie du processus sont hors chronomètre. La machine n’est pas isolée des
autres activités utilisateur.

C et JS ne sont pas remesurés. Le README reprend les 14 temps primaires App
et les 23 temps supplémentaires AppAll dans les 37 lignes Go WIP ; Go/C
porte sur 14 cas, Go/JS sur 37. Seule la colonne Go Release est actualisée,
sans ajout de tableau ni changement des statuts ou du Backlog non mesuré.
Le catalogue supprimé par l’utilisateur n’est pas recréé.

Depuis `var/benchmark/go-native-callbacks-20261126`, pour réexécuter les mêmes
binaires figés dans un nouveau répertoire :

```sh
python3 -B proofs/campaign-scripts/measure-pair.py \
  --after-core after-core --after-all after-all --output replay --execute
```

Scripts, commandes, manifestes, sorties brutes et preuves sont archivés.
Les gros overlays générés et la copie du bundle de référence restent dans le
workspace ; l’archive conserve leurs empreintes, leur provenance et les moyens
de reconstruction, sans dupliquer leur contenu intégral.
