# Temps du compilateur gopurs hébergé en Rust — bibliothèques

Date : 2026-10-04. 49 cellules Rust examinées, **45 mesurées et 4 en échec de génération ou de validation** ; tous les hôtes génèrent du **Go**.

## Protocole

- Compilateurs et entrées TAST/FFI figés dans `/Users/0x1/Documents/htdocs/altbak.pub/var/benchmark/gopurs-packages-rust-20261004`.
- Frontend hors chronomètre ; mesure `[gopurs] backend total` incluant chargement, préparation, PBO, génération, écritures et drain.
- Pour chaque paquet validé : une chauffe Rust puis 5 mesures, exécutions sérialisées ; médiane publiée.
- Workers chargement/préparation/PBO/émission **8/8/8/8**, pipeline actif ; sorties et caches PBO réinitialisés entre les processus.
- Chaque génération est comparée octet par octet à l'hôte Go sur les mêmes entrées ; éventuel repli oracle JavaScript indiqué par ligne.
- Les colonnes JS/Go préexistantes proviennent de campagnes antérieures. Cette extension ne mesure pas de nouveaux ratios entre hôtes.
- La compilation Go des applications et l'exécution de leurs tests ne font pas partie de cette campagne de génération.

Machine : Apple M4 Pro, 14 processeurs logiques, 48 Gio de RAM.

## Résultats

| Paquet | Modules | Types | Rust médian (ms) | Min–max (ms) | Fichiers exacts | Oracle |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| gopurs-argonaut-core | 254 | 129932 | 1213 | 1202–1234 | 324 | go |
| gopurs-arrays | 167 | 79038 | 3298 | 3243–3396 | 217 | go |
| gopurs-assert | 58 | 13674 | 110 | 110–111 | 82 | go |
| gopurs-avar | 236 | 134408 | 1328 | 1323–1335 | 292 | go |
| gopurs-catenable-lists | 153 | 81907 | 827 | 822–832 | 190 | go |
| gopurs-console | 161 | 73059 | 579 | 576–602 | 210 | go |
| gopurs-datetime | 198 | 105334 | 1531 | 1524–1535 | 252 | go |
| gopurs-effect | 60 | 13929 | 123 | 122–129 | 87 | go |
| gopurs-enums | 169 | 78152 | output mismatch | — | — | go |
| gopurs-exceptions | 79 | 19638 | 185 | 185–187 | 108 | go |
| gopurs-foldable-traversable | 161 | 77931 | 852 | 851–867 | 211 | go |
| gopurs-foreign | 228 | 118639 | 1020 | 1008–1209 | 286 | go |
| gopurs-foreign-object | 253 | 133370 | 1554 | 1543–1555 | 318 | go |
| gopurs-free | 199 | 115212 | 930 | 918–936 | 239 | go |
| gopurs-functions | 60 | 14451 | 135 | 133–136 | 87 | go |
| gopurs-integers | 162 | 73190 | 765 | 760–769 | 211 | go |
| gopurs-js-bigints | 80 | 17221 | 159 | 156–190 | 113 | go |
| gopurs-js-date | 250 | 132520 | 1353 | 1330–1362 | 313 | go |
| gopurs-js-promise | 140 | 62679 | output mismatch | — | — | go |
| gopurs-js-promise-aff | 259 | 139038 | 1402 | 1382–1408 | 324 | go |
| gopurs-js-uri | 162 | 73181 | 577 | 571–581 | 212 | go |
| gopurs-lazy | 134 | 62320 | 473 | 465–478 | 167 | go |
| gopurs-node-buffer | 241 | 121585 | 1188 | 1178–1195 | 305 | go |
| gopurs-node-event-emitter | 300 | 179138 | 2379 | 2356–2438 | 372 | go |
| gopurs-node-fs | 283 | 147570 | 1711 | 1696–1730 | 361 | go |
| gopurs-node-http | 312 | 157371 | 1728 | 1709–1783 | 411 | go |
| gopurs-node-net | 299 | 152727 | 1684 | 1661–1693 | 386 | go |
| gopurs-node-path | 60 | 13798 | 122 | 119–124 | 87 | go |
| gopurs-node-process | 285 | 147900 | 1591 | 1581–1599 | 360 | go |
| gopurs-node-streams | 313 | 185333 | 2737 | 2730–2782 | 394 | go |
| gopurs-now | 287 | 148368 | 1619 | 1608–1628 | 364 | go |
| gopurs-nullable | 230 | 119180 | 1095 | 1085–1103 | 290 | go |
| gopurs-numbers | 161 | 73154 | 716 | 709–745 | 210 | go |
| gopurs-ordered-collections | 241 | 130500 | 2040 | 2031–2065 | 302 | go |
| gopurs-partial | 161 | 72998 | 568 | 564–573 | 210 | go |
| gopurs-prelude | 163 | 75446 | failed (division by zero) | — | — | go |
| gopurs-random | 162 | 73215 | 576 | 575–578 | 212 | go |
| gopurs-record | 165 | 74983 | 591 | 588–614 | 216 | go |
| gopurs-refs | 161 | 73232 | 581 | 580–584 | 210 | go |
| gopurs-run | 266 | 160844 | 1547 | 1526–1557 | 326 | go |
| gopurs-spec | 390 | 227763 | 3807 | 3767–3926 | 490 | go |
| gopurs-st | 161 | 73636 | 583 | 578–589 | 210 | go |
| gopurs-strings | 193 | 81414 | output mismatch | — | — | go |
| gopurs-strings-extra | 187 | 79322 | 2008 | 1998–2015 | 241 | go |
| gopurs-unfoldable | 161 | 73511 | 599 | 592–602 | 210 | go |
| gopurs-unsafe-coerce | 161 | 73035 | 578 | 574–578 | 210 | go |
| gopurs-uuid | 290 | 175873 | 2338 | 2309–2358 | 355 | go |
| gopurs-variant | 165 | 91537 | 949 | 948–959 | 206 | go |
| gopurs-yoga-json | 400 | 238424 | 5402 | 5358–5434 | 498 | go |

**315 générations / 81753 fichiers générés exacts** relus après mesure.

Sous-total de la colonne Rust `gopurs-*` : **58.66 s** pour **46/50 paquets**. C'est la somme des médianes validées affichées, incluant `gopurs-aff` déjà qualifié ; les 4 échecs sont exclus. Ce n'est pas une invocation globale chronométrée.

## Diagnostics conservés

`gopurs-assert` est une bibliothèque : son runner public appelle gopurs sans `--main`. La première tentative du harness exigeait à tort `Test.Main` ; la reprise utilise la commande réelle et conserve cinq mesures exactes.

- **gopurs-enums** : output mismatch. Fichiers différents : `purescript/Control_Bind.go`. Diagnostic reproduit avec le même binaire et les mêmes entrées. Aucune médiane valide n'est publiée pour cette ligne.
- **gopurs-js-promise** : output mismatch. Fichiers différents : `purescript/Data_Traversable.go`, `purescript/Promise_Lazy.go`, `purescript/Test_Main.go`. Diagnostic reproduit avec le même binaire et les mêmes entrées. Aucune médiane valide n'est publiée pour cette ligne.
- **gopurs-prelude** : failed (division by zero). Diagnostic reproduit avec le même binaire et les mêmes entrées. Aucune médiane valide n'est publiée pour cette ligne.
- **gopurs-strings** : output mismatch. Fichiers différents : `purescript/Test_Assert.go`. Diagnostic reproduit avec le même binaire et les mêmes entrées. Aucune médiane valide n'est publiée pour cette ligne.

Les différences de sources générées constituent un échec du contrôle d'identité, pas une preuve à elles seules de différence sémantique. La division par zéro de Prelude interrompt effectivement le backend Rust avant son horloge totale.

Les échantillons, empreintes et chemins des sorties brutes sont dans le rapport JSON associé.
