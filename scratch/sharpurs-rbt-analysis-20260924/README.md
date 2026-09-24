# RBTree : coût résiduel et prototype de motifs plats — 24 septembre 2026

## Résultat de l'analyse

RBTree reste le premier poste de la campagne officielle (**42,9 ms** sur 58,6 ms). Le coût est
presque entièrement dans `buildTree` :

| Mesure FSI optimisée (meilleur de 15) | Temps |
| --- | ---: |
| `buildTree 100000` | ~32 ms |
| `depth` du même arbre | 0,3–1,0 ms |

`buildTree` alloue **124 Mo** par construction : 2,58 M nœuds `T` de 48 octets pour 100 000
insertions, soit ~26 allocations par insertion. Le coût est donc dominé par l'allocation et le
GC, pas par le calcul.

## Prototypes mesurés (F# généré, sources identiques)

| Variante | `buildTree` FSI | RBTree programme | Total programme |
| --- | ---: | ---: | ---: |
| Génération actuelle | 31,6–32,8 ms | 40,85 ms (méd.) | 60,64 ms |
| `balance`/`ins` en `match` plat écrit à la main | **29,9–30,9 ms** | **39,26 ms** | **58,87 ms** |
| Idem + appels natifs directs (sans `_adt_native_apply`) | 29,0–30,7 ms | — | — (rejeté, voir limites) |
| Projections via fonctions d'accès (`isTag`/`field`) | 33,9–35,5 ms | — | pire ou neutre |
| Émetteur : `match` imbriqués avec variables de motif | 32,4–32,5 ms | 40,78 ms | 59,64 ms |

Les 6 passages programme intercalés donnent, pour le prototype `match` plat :
**RBTree 40,85 → 39,26 ms (−3,9 %)** et **total 60,64 → 58,87 ms (−2,9 %)**.

## Pourquoi l'émetteur `match` imbriqué ne suffit pas

Le gain du prototype vient de la forme **plate** des motifs :

- `balance` écrit à la main = **5 cas, un seul `match`, motifs imbriqués** → F# produit **une
  méthode de 2 874 octets**, sans découpage ;
- `balance` généré avec des `match` imbriqués = F# découpe encore la fonction en **22 méthodes
  de continuation** (`$cont@49-*`) ; le coût de dispatch annule le gain ;
- le fichier généré passe de 59 690 à 48 994 octets, mais la structure reste trop profonde.

Autrement dit, la transformation utile est une **synthèse de motifs plats** (compilation de
l'arbre de décision en motifs F#), pas un simple remplacement `if` → `match`.

## Limites et variantes rejetées

- **Appels natifs directs** : supprimer `_adt_native_apply` entre corps natifs gagne ~2,5 % mais
  change la frontière d'exception (`TargetInvocationException`) validée par les fixtures
  `adt-multi` (chaînes et profondeurs comparées à l'oracle générique) ; non retenu.
- **Fonctions d'accès** : remplacer les `match` inline par des fonctions `isTag`/`field` réduit la
  taille du fichier de 39 % mais ralentit ou reste neutre (le JIT n'inline pas ces appels).
- L'émetteur `match` imbriqué a été **retiré** et la génération restaurée à l'identique
  (empreinte de `Test.RBTree.fs` vérifiée).

## Suite proposée

1. Émettre un `match` plat par arbre de décision (synthèse de motifs), en conservant le chemin
   `if` pour les formes non reconnues. Cible mesurée : la forme `Test.RBTree.hand-match.fs`
   ci-jointe.
2. Conserver `_adt_native_apply` : aucune modification de la sémantique d'exception.
3. Valider par les invariants d'arbre, les fixtures ADT (à réparer d'abord, voir ci-dessous) et
   le validateur officiel.

## État de l'environnement (indépendant de cette analyse)

- `sharpurs/test/Main.purs` ne compilait plus contre le `Ann` actuel de PBO (champ
  `sourceUsage`) ; le littéral est corrigé pour que `spago build` fonctionne.
- Les suites de fixtures `test:adt-*` et autres échouent avant tout code backend : le `purs` de
  `/Users/0x1/.local/bin` (16 septembre) ne parse pas le FFI ESM de `prelude-6.0.2`, et
  `Builder.buildModules` attend une API plus récente. Une réparation dédiée est nécessaire avant
  de pouvoir s'appuyer sur ces suites.
- Aucune modification de l'émetteur n'est conservée par cette étape.

[Prototype plat](Test.RBTree.hand-match.fs) · [diff](Test.RBTree.hand-match.diff)
· [mesures](measurements.json)
