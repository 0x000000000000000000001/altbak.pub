| Profil | naive | pushdown | precise | fusion | specialized_drop | reuse | retained_fields | Final / naive | Blocs favorables |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `list_map` unique | 15.280 | 15.157 | 12.708 | 12.637 | 12.647 | 11.338 | 11.316 | -25.94 % | 7/7 |
| `list_map` racine partagée | 15.445 | 15.411 | 14.845 | 14.716 | 14.702 | 15.604 | 15.559 | +0.74 % | 2/7 |
| `list_map` mixte (enfant partagé) | 15.576 | 15.459 | 14.939 | 14.841 | 14.779 | 15.681 | 15.639 | +0.40 % | 2/7 |
| `list_filter` unique | 11.577 | 11.658 | 8.852 | 8.845 | 8.825 | 8.627 | 8.541 | -26.22 % | 7/7 |
| `list_filter` racine partagée | 11.508 | 11.558 | 11.159 | 11.135 | 11.118 | 11.895 | 11.885 | +3.28 % | 0/7 |
| `list_filter` mixte (enfant partagé) | 11.536 | 11.537 | 11.192 | 11.073 | 11.116 | 11.802 | 11.731 | +1.69 % | 0/7 |
| `tree_map` unique | 7.691 | 7.598 | 7.477 | 7.363 | 7.172 | 5.448 | 5.540 | -27.97 % | 7/7 |
| `tree_map` racine partagée | 7.490 | 7.619 | 7.209 | 7.457 | 7.293 | 8.245 | 8.231 | +9.89 % | 0/7 |
| `tree_map` mixte (enfant partagé) | 7.454 | 7.562 | 7.244 | 7.146 | 7.019 | 6.909 | 6.957 | -6.67 % | 7/7 |
| `tree_choose` unique | 3.679 | 3.734 | 3.705 | 3.784 | 3.663 | 3.701 | 3.598 | -2.20 % | 5/7 |
| `tree_choose` racine partagée | 3.751 | 3.763 | 3.806 | 3.727 | 3.675 | 3.743 | 3.663 | -2.33 % | 5/7 |
| `tree_choose` mixte (enfant partagé) | 3.724 | 3.894 | 3.785 | 3.741 | 3.656 | 3.720 | 3.695 | -0.80 % | 5/7 |
| `tree_update` unique | 4.050 | 4.033 | 4.002 | 3.930 | 3.930 | 3.710 | 3.635 | -10.23 % | 6/7 |
| `tree_update` racine partagée | 4.086 | 4.029 | 3.896 | 3.952 | 3.966 | 4.009 | 4.005 | -1.99 % | 5/7 |
| `tree_update` mixte (enfant partagé) | 3.930 | 4.065 | 3.965 | 3.923 | 4.029 | 3.793 | 3.748 | -4.64 % | 7/7 |

Écarts de chaque passage face au précédent. Chaque cellule indique **variation de durée / blocs favorables** ; un signe négatif signifie plus rapide. `=` indique que le corps Rust de cette fonction reste identique : son écart ne mesure pas une transformation locale de cette fonction.

| Profil | pushdown | precise | fusion | specialized_drop | reuse | retained_fields |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| `list_map` unique | -0.80 % / 6 = | -16.16 % / 7 | -0.55 % / 7 | +0.08 % / 1 | -10.36 % / 7 | -0.19 % / 4 = |
| `list_map` racine partagée | -0.22 % / 3 = | -3.67 % / 7 | -0.87 % / 6 | -0.10 % / 4 | +6.13 % / 0 | -0.29 % / 3 = |
| `list_map` mixte (enfant partagé) | -0.75 % / 5 = | -3.36 % / 7 | -0.65 % / 7 | -0.42 % / 5 | +6.10 % / 0 | -0.27 % / 5 = |
| `list_filter` unique | +0.69 % / 3 | -24.07 % / 7 | -0.08 % / 3 | -0.22 % / 3 | -2.25 % / 7 | -0.99 % / 4 = |
| `list_filter` racine partagée | +0.44 % / 3 | -3.45 % / 7 | -0.22 % / 6 | -0.15 % / 3 | +6.99 % / 0 | -0.09 % / 4 = |
| `list_filter` mixte (enfant partagé) | +0.01 % / 4 | -2.99 % / 7 | -1.06 % / 6 | +0.39 % / 3 | +6.17 % / 1 | -0.60 % / 5 = |
| `tree_map` unique | -1.21 % / 4 = | -1.60 % / 6 | -1.52 % / 5 | -2.59 % / 5 | -24.04 % / 7 | +1.69 % / 2 = |
| `tree_map` racine partagée | +1.71 % / 2 = | -5.38 % / 7 | +3.44 % / 1 | -2.20 % / 6 | +13.05 % / 0 | -0.16 % / 4 = |
| `tree_map` mixte (enfant partagé) | +1.45 % / 2 = | -4.21 % / 7 | -1.36 % / 2 | -1.77 % / 7 | -1.56 % / 4 | +0.69 % / 3 = |
| `tree_choose` unique | +1.52 % / 3 | -0.78 % / 4 | +2.13 % / 3 | -3.20 % / 5 | +1.03 % / 3 | -2.79 % / 6 = |
| `tree_choose` racine partagée | +0.32 % / 4 | +1.16 % / 4 | -2.09 % / 5 | -1.41 % / 5 | +1.85 % / 2 | -2.12 % / 3 = |
| `tree_choose` mixte (enfant partagé) | +4.56 % / 2 | -2.81 % / 4 | -1.16 % / 3 | -2.27 % / 5 | +1.73 % / 3 | -0.68 % / 4 = |
| `tree_update` unique | -0.41 % / 5 = | -0.77 % / 6 | -1.79 % / 5 | +0.00 % / 4 | -5.59 % / 6 | -2.02 % / 4 |
| `tree_update` racine partagée | -1.40 % / 4 = | -3.29 % / 4 | +1.43 % / 3 | +0.36 % / 4 | +1.09 % / 2 | -0.11 % / 4 |
| `tree_update` mixte (enfant partagé) | +3.43 % / 2 = | -2.47 % / 5 | -1.05 % / 5 | +2.70 % / 2 | -5.87 % / 5 | -1.18 % / 4 |
