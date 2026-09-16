# Intérêt chiffré des résumés de champs à travers les appels

16 septembre 2026. **Preuve expérimentale dans des variantes Rust isolées, sans modification de Purust, PBO, Haskell ni du benchmark source.** Le test dérive de la fonction Records réellement générée. Il ajoute une frontière d'appel contrôlée pour mesurer le maintien des champs scalaires entre appels. Ce n'est pas une nouvelle mesure de gain global du benchmark officiel.

Résultat principal : **379,29 µs → 44,46 µs pour 10 000 appels**, soit **−88,26 % en comparaison appariée** (environ ×8,5). Les 10 000 appels subsistent. À nombre d'appels identique, le même worker scalaire avec reconstruction à chaque tour prend **378,42 µs** : l'essentiel du gain provient bien de la conservation de la représentation scalaire entre les appels.

## Expérience et ablations

Le source de départ est `output/purust_output/Purs_Test_Records/src/lib.rs`, SHA256 `6343d7b109a260c8360b22ceee33be861bafe80972b0c01b580a271836d447f3`. La fonction originale est extraite sans modification. `record_step` reprend exactement son expression de mise à jour, avec seulement un renommage des deux paramètres. `scalar_step` effectue les quatre mêmes opérations sur un tableau de quatre entiers.

Les deux workers sont dans une crate Rust compilée séparément, avec `#[inline(never)]`, `opt-level=3`, `lto=off`, `embed-bitcode=no`. Le caller est compilé avec les mêmes options. Le corps du callee n'est donc pas disponible pour une optimisation interprocédurale du caller dans cette expérience. Les deux variantes ont réellement la même frontière d'appel.

| Variante, sans snapshot intermédiaire | Temps médian du meilleur de 10 | Rôle |
|---|---:|---|
| Rust généré original, sans frontière ajoutée | 368,33 µs | Vérifier que la frontière n'a pas créé artificiellement l'essentiel du coût |
| Appel `record_step(n, record)` à chaque tour | 379,29 µs | Référence avec frontière d'appel |
| Appel `scalar_step`, extraction/reconstruction à chaque tour | 378,42 µs | Même worker que la variante optimisée, sans conservation des scalaires entre appels |
| Appel `scalar_step`, scalaires conservés pendant la boucle | **44,46 µs** | Reconstruction une seule fois à la sortie |
| Ancien prototype scalaire local, sans appel par tour | 2,92 µs | Limite de comparaison : LLVM peut simplifier et vectoriser davantage sans la frontière |

Le gain apparié du maintien des scalaires par rapport au témoin utilisant **le même worker** est **−88,23 %**. Changer uniquement le worker tout en reconstruisant à chaque tour n'apporte pas de gain net ici (+0,19 % face au worker record, dans le bruit).

Cela isole le mécanisme utile : enlever les projections, manipulations de `Value` et mises à jour COW répétées de la boucle appelante. Le nombre d'allocations ne baisse pas dans ce cas : **3 allocations dans les deux variantes**, car le réemploi existant était déjà efficace.

## Pourquoi l'information sur la rétention compte

On conserve des snapshots complets après certaines itérations. Leur création et leur destruction sont présentes dans toutes les variantes comparées. La version scalaire reconstruit le record avant chaque snapshot et maintient les anciennes versions intactes.

| Fréquence des snapshots | Worker record | Scalaires conservés | Variation appariée | Reconstructions de la variante scalaire |
|---|---:|---:|---:|---:|
| Aucun | 379,29 µs | **44,46 µs** | **−88,26 %** | 1 |
| Tous les 100 tours | 386,75 µs | **48,54 µs** | **−87,39 %** | 100 |
| À chaque tour | 733,00 µs | **603,29 µs** | **−16,87 %** | 10 000 |

Quand chaque état doit être conservé, la plupart des reconstructions sont obligatoires et le gros gain disparaît. Le gain restant peut notamment venir de l'absence de relecture des champs à chaque tour ; la présente mesure n'isole pas chaque instruction responsable de ces 17 %.

Les allocations/désallocations sont identiques entre variantes pour chaque scénario : 3 sans snapshots, 306 pour 100 snapshots, 30 013 pour 10 000 snapshots, allocations du vecteur de snapshots incluses. Les compteurs sont compilés dans un binaire distinct de celui utilisé pour le temps.

**Limite de portée : les rétentions sont explicites dans le caller expérimental.** Cette expérience mesure leur coût et vérifie qu'on peut les respecter ; elle ne démontre pas encore qu'une future analyse Haskell détectera les captures à l'intérieur de fonctions arbitraires.

## Ce que les futurs contrats devraient fournir

Pour choisir cette représentation à travers une frontière opaque, le caller aurait besoin d'un contrat décrivant :

- Les feuilles nécessaires au calcul : `a`, `b.c`, `b.d.e`, `b.d.f`, avec leurs types déjà disponibles dans le TAST.
- La relation résultat/entrée : ces quatre feuilles sont remplacées ; les autres champs proviennent de l'entrée et restent préservés.
- Les conservations et observations du record ou de ses sous-records : captures, retours, callbacks, appels inconnus ou effets imposant une matérialisation.
- **Un worker scalaire effectivement exporté par le backend**, avec la correspondance entre ses paramètres/résultats et les chemins du record.

Un résumé seul n'accélère pas une fonction opaque dont l'unique interface reçoit un `Value`. Il faut que le callee fournisse la version spécialisée correspondante. Dans l'expérience, les deux workers sont écrits dans les variantes de test et leur équivalence est vérifiée ; aucun générateur automatique de workers ni lecteur de nouveaux résumés TAST n'a été implémenté.

Lorsque le corps de la fonction est disponible dans PBO/Purust, ces faits peuvent être déduits du programme existant. **L'expérience ne prouve donc pas qu'un ajout Haskell est nécessaire pour ce cas.** Pour la compilation séparée ou les frontières où le corps n'est pas disponible, exporter contrat et worker est une manière de rendre cette optimisation accessible sans ouvrir le corps au caller. Les faits Haskell et la description du worker généré par le backend sont deux parties distinctes de cette interface.

Le cas `Test.Records` actuel n'a pas cette frontière : son prototype local reste plus rapide, à environ 3 µs. Les −88 % mesurés ici ne s'ajoutent pas au précédent gain Records et ne doivent pas être appliqués au total du README. Ils prouvent l'intérêt d'étendre la transformation aux programmes contenant ces appels.

## Protocole et validation

- 21 tours avec ordre aléatoire des 11 couples variante/scénario ; 3 warm-ups puis 10 échantillons par couple : **2 310 échantillons chronométrés**.
- Allocateur natif MiMalloc, entrée partagée conservée, nombre d'itérations 10 000 passé par `black_box`, quatre champs du résultat observés, destruction du résultat et des snapshots incluse dans le temps.
- Les pourcentages sont les médianes des rapports appariés, pas les rapports entre les médianes indépendantes du tableau.
- **880 cas validés et 246 072 snapshots vérifiés** : 10 tailles de boucle dont zéro, quatre initialisations, deux formes de records, trois politiques de rétention et les variantes applicables.
- Oracle arithmétique indépendant pour les quatre champs, aliases de la racine et des sous-records conservés, champ supplémentaire inchangé. Des cas proches de `i64::MAX` vérifient la conservation du comportement wrapping de la représentation Rust actuelle ; cela ne valide pas globalement la sémantique Int de tous les backends.
- Compteur séparé : exactement **10 000 appels** au worker record ou scalaire dans chaque scénario. Le champ brut `materializations` de `counts.json` ne compte que les appels au helper du caller ; sa valeur zéro pour `record` signifie que les mises à jour sont réalisées dans `record_step`, appelé 10 000 fois.
- Désassemblage du binaire chronométré : appels `bl record_step` et `bl scalar_step` présents dans les boucles, avec branche arrière et décrément du compteur. Le calcul n'a pas été remplacé par une formule fermée. Voir `assembly.txt` et `assembly-summary.json`.

Contexte historique : le README officiel du fork `altbak.pub-purust` donne **385,38 µs** pour Records généré, **98,75 µs** pour la version native fonctionnelle et **4,25 µs** pour la version native impérative. Le témoin actuel à 368,33 µs est cohérent avec cet ordre de grandeur. Les conclusions causales reposent sur les comparaisons présentes, pas sur les écarts avec les chiffres historiques.

## Reproduction

Depuis le répertoire du présent rapport :

```sh
python3 prepare_callee.py
python3 build.py
python3 measure.py
```

`build.py` compile les crates séparées, valide les résultats et enregistre les compteurs. `measure.py` utilise uniquement le binaire sans instrumentation. Les sources originales et transformées, commandes exactes, SHA256, version Rust, données brutes et résultats se trouvent dans ce dossier : `callee-manifest.json`, `build-manifest.json`, `toolchain.txt`, `validation.txt`, `counts.json` et `timings.json`.
