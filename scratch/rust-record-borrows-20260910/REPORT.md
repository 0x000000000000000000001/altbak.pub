# Projections de records empruntées : prochaine piste B14

10 septembre 2026. **Prototype isolé, non intégré à Purust.** Les lectures de Records demandent encore des copies de références alors que leur résultat final est un entier. En empruntant le chemin jusqu'au scalaire, Records passe de **0,621 à 0,349 ms**, soit **0,272 ms gagnée (−43,8 %)** sur cinq blocs favorables. Cela correspond à environ **2,1 %** du total de 12,98 ms documenté dans le README, avant toute confirmation par une intégration générale.

## Code et coût visé

La sortie actuelle est `run/bak/rust/output/purust_output`, dans `altbak.pub-purust`. Le chemin demandé `altbak.pub-rust/run/bak/rust/output/purescript` n'existe pas ; la sortie du checkout normal est ancienne. Le noyau Records récent lit ses quatre valeurs aux lignes 56–59 :

```rust
// Forme du chemin de lecture généré, avant les setters :
purs_local_1.clone().get_b().get_d().get_f().unwrap_int()
```

Le clone de la racine possède une référence supplémentaire. Chaque getter intermédiaire retourne également une `Value` possédée : par exemple, `get_b` retourne `r.b.clone().unwrap()`. L'ensemble des quatre expressions demande quatre clones de racine, trois de `b` et deux de `d` par itération. Les getters des entiers clonent aussi leur `Value`, sans allouer un entier séparé. Les appels à `PerceusPtr::clone` sont comptés séparément dans des binaires instrumentés.

L'essai change uniquement ces lectures :

```rust
(&purs_local_1).get_b_ref().get_d_ref().get_f_ref().unwrap_int()
```

Les nouveaux getters retournent `&Value`, avec les mêmes variantes de records admises et les mêmes vérifications de présence. Ils sont ajoutés dans une copie isolée de `purust_core`, avec la même visibilité que les getters existants. Les quatre entiers sont calculés avant la première écriture. Les huit setters, leurs gardes d'unicité, les détachements/réinstallations et les layouts `Option<Value>` restent identiques.

Conserver un `b` ou `d` **possédé** pendant les setters serait une autre transformation : cela rendrait le chemin partagé et pourrait réintroduire des copies. Le prototype ne conserve que des emprunts pendant les lectures.

## Résultats

Trois binaires de runners complets, même source récente et mêmes dépendances, profil **release O1/debug=true/mimalloc**. Le workspace de mesure est créé sans copier `target`. Les six getters empruntés sont présents dans les trois variantes ; seules les quatre expressions de lecture du module Records varient. Chaque binaire est construit après écriture de sa propre source.

Cinq blocs font tourner puis inversent l'ordre des trois variantes. Tous les builds, tests et comptages sont arrêtés avant le chronométrage. Chaque runner conserve son warm-up et son meilleur de dix, et vérifie les quatorze résultats. Les chiffres ci-dessous sont les médianes entre les cinq processus ; le total additionne les médianes par benchmark.

| Variante | Records dans le runner | Noyau observant les quatre champs | Total de la suite |
| --- | ---: | ---: | ---: |
| Avant | 0,621 ms | 0,660 ms | 12,567 ms |
| Emprunter seulement la racine | 0,512 ms | 0,525 ms | 12,443 ms |
| Emprunter le chemin complet | **0,349 ms** | **0,374 ms** | **12,162 ms** |

Les plages de Records sont disjointes : **0,583–0,673 ms** avant, **0,466–0,558 ms** avec la racine empruntée et **0,346–0,384 ms** avec le chemin complet. Les cinq blocs sont favorables à chacune des deux étapes sur Records. Les contrôles séparés observent les quatre champs, initialisation et destruction incluses, et confirment environ **0,286 ms** de gain sur le noyau complet.

La suite baisse ici de **0,405 ms**, mais RBTree, dont le code est inchangé, varie aussi de **11,441 à 11,330 ms**. Je retiens donc le gain local de **0,272 ms** comme résultat directement attribuable à cette piste ; les **0,405 ms** observées sur le total ne constituent pas une prévision de gain garanti après intégration. Les différentes séries de prototypes conservent leurs propres références et ne doivent pas être additionnées.

## Comptage séparé

Le [compteur](count.py) observe **110 000 → 70 000 → 20 000 appels à `PerceusPtr::clone`** pendant les 10 000 itérations. Les deux détachements de champs conservent leurs 20 000 clones ; les lectures n’en demandent plus dans la dernière variante. La vérification des quatre champs ajoute cinq clones, relevés séparément et exclus du comptage de la boucle. [Résultats](counts.json).

Il s’agit d’appels dynamiques dans des binaires instrumentés O1. Le compteur atomique peut empêcher certaines éliminations LLVM : ces chiffres ne prouvent pas que 90 000 instructions de comptage survivaient dans le binaire original non instrumenté. Le gain temporel provient exclusivement des binaires sans instrumentation.

## Référence native et pastilles

Le [README officiel](../../../altbak.pub/README.md#rust) donne **Records 0,674 ms**, contre **0,004 ms** dans sa dernière colonne native ; Rust total **12,98 ms**. Le natif Records calcule des accumulateurs scalaires et ne retourne que `f`, ce qui laisse le compilateur éliminer les autres. Notre validation observe les quatre champs du record complet ; aucune promesse d'atteindre 4 µs n'est faite.

Cette piste étend **B14**, encore jaune : éviter le comptage explicite lors de lectures. **B22** reste vert sur son périmètre de setters spécialisés ; il ne reste déjà que trois allocations initiales sur cette boucle. **B21/B23**, les layouts et champs typés, désignent un chantier différent : les layouts actuels sont identifiés par leurs labels et conservent `Option<Value>`. Les remplacer par des champs numériques demanderait de distinguer aussi les types des champs et de traiter les frontières de représentation. Ce prototype ne le fait pas.

Une règle générale pourrait commencer par les chaînes de projections dont le TAST prouve un résultat scalaire natif Copy, sans retour/capture de la référence empruntée et sans appel opaque pendant la lecture. Les contextes qui demandent une valeur possédée garderaient le getter actuel. La reconnaissance ne dépendrait d'aucun nom de benchmark ou de champ particulier. Il reste à intégrer cette règle, couvrir ses frontières dans des fixtures TAST puis mesurer la sortie réellement générée.

## Validation et artefacts

- Quatorze résultats corrects pour chaque runner et chaque passage.
- Résultat complet observé : `a`, `c`, `e` et `f`, pour plusieurs seeds et nombres d'itérations.
- Anciennes racines, enfants et feuilles conservés, ainsi que les huit combinaisons indépendantes de partage.
- **Trois allocations et trois libérations** pour chacune des trois variantes, y compris 10 000 itérations ; aucune copie de chemin réintroduite.
- Sources du compilateur, FFI et sortie courante inchangées. Tous les fichiers expérimentaux restent dans ce scratch, avec `build/` ignoré.

[Script](probe.py), [mesures brutes](results.json), [validation](validation.json), [empreintes et configuration](metadata.json). Les binaires de temps du noyau utilisent vingt échantillons après un warm-up par processus et incluent la destruction.

Le [prototype de parcours RBTree emprunté](../rust-borrowed-depth-20260910/REPORT.md), mesuré séparément, supprime 200 000 clones mais ne montre qu'un petit écart de **0,103 ms** sur la suite, avec une paire défavorable. Records fournit ici la piste la plus nette, avec un potentiel limité à quelques dixièmes de milliseconde.
