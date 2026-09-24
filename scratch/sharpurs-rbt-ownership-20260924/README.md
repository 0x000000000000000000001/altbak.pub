# RBTree : le levier mutation/ownership est réfuté — 24 septembre 2026

## Question

Le prototype de motifs plats plafonnait à ~4 % sur RBTree. L'hypothèse restante était le
« levier massif » : une analyse de propriété/linéarité permettant de **muter les nœuds en place**
au lieu de les réallouer (`buildTree`/`insert`/`ins`/`balance`), comme le fait le record-borrow de
purust. Avant d'implémenter une telle analyse dans le compilateur, le potentiel a été mesuré sur
le même algorithme d'Okasaki.

## Protocole

- Deux implémentations F#, **même algorithme** (insertion Okasaki + balance 4 cas + makeBlack) :
  persistante (union `T`/`E`) et mutable (nœuds de classe, rotations et recoloriages en place).
- Équivalence vérifiée : profondeur identique pour n = 0, 1, 2, 3, 5, 10, 50, 100, 1000, 4096,
  20000 et 100000 (aucun écart), arbre valide (profondeur 22 pour 100 000).
- Mesures dans un **programme Release compilé** (meilleur de 10, 3 échauffements) et en FSI
  optimisé (meilleur de 15). Aucun code de benchmark ne les déclenche : c'est un banc isolé.

## Résultats (100 000 insertions, meilleur de 10, programme Release)

| Variante | Run 1 | Run 2 | Run 3 |
| --- | ---: | ---: | ---: |
| Persistante (génération actuelle) | 35,65 ms | 36,36 ms | 36,12 ms |
| Mutable, propriétés automatiques | 51,96 ms | 51,73 ms | 50,27 ms |
| Mutable, champs directs | 52,12 ms | 51,69 ms | 54,17 ms |

FSI optimisé : persistante 37,1 ms, mutable 46,2 ms.

**La mutation en place est ~45 % plus lente que la réallocation persistante**, quelle que soit la
représentation. Le levier « massif » n'existe pas ici.

## Pourquoi

- L'allocation dans le nursery .NET est un simple bump pointer ; les nœuds meurent jeunes et les
  collections gen0 sont quasi gratuites pour eux.
- Muter un arbre vivant touche des objets promus : chaque écriture de référence paie une barrière
  d'écriture, et le GC doit rescanner les régions modifiées. La localité se dégrade aussi (les
  nœuds anciens sont dispersés), alors que la version persistante écrit son chemin dans le nursery
  de façon contiguë.
- Le coût observé de la mutation (~6 ns par écriture de référence) dépasse le coût d'allocation +
  construction d'un nœud neuf.

## Conséquence pour le compilateur

- Une analyse de propriété/linéarité pour RBTree **dégraderait** les performances ; elle n'est pas
  justifiée par ce benchmark. Aucune modification du compilateur n'est conservée.
- L'écart avec la colonne FFI impérative (10,5 ms) ne vient pas de la mutation seule : la
  méthodologie autorise explicitement ces colonnes à préallouer, réutiliser des arènes et changer
  d'algorithme. Le code FP compilé (~40 ms) reste au niveau du F# FP écrit à la main (~37,8 ms).
- Le seul levier de codegen restant sur RBTree est la synthèse de motifs plats (~4 %), déjà
  documentée dans `sharpurs-rbt-analysis-20260924`.

## Reproduction

```sh
cd /Users/0x1/Documents/htdocs/scratch/sharpurs-rbt-ownership-20260924/bench
dotnet build -c Release -p:Optimize=true -o ../bin
dotnet ../bin/Program.dll     # persistent vs mutable
dotnet fsi --optimize+ --exec ../mutable_proto.fsx   # équivalence + mesures FSI
```

[mutable_proto.fsx](mutable_proto.fsx) · [bench/Program.fs](bench/Program.fs)
· [bench-fields/Program.fs](bench-fields/Program.fs) · [mesures](measurements.txt)
