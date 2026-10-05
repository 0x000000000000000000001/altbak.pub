# Saturation du compilateur gopurs hébergé en Rust

Campagne autorisée dans la nuit du 4 au 5 octobre 2026. Les trois hôtes gopurs
génèrent du **Go**. Archive : `var/benchmark/gopurs-rust-saturation-20261004/`.

## Protocole

- `reclaim-purust.py` inventorie les caches Purust régénérables, puis vérifie les
  empreintes des sources et exécutables conservés. Les preuves antérieures et
  les métadonnées suivies par Git sont conservées.
- `start.mjs` fige les compilateurs, les sources et les **51 corpus publiés**
  dans `compilation-refresh-20261004`, ainsi que leurs sorties Go canoniques.
  Il profile ensuite Arrays, b8x et Aff, hors de toute mesure de sélection.
- `build.mjs` reconstruit un contrôle ou candidat O3 + ThinLTO avec le même
  profil PGO historique, à partir d'un instantané privé. La composition retenue
  doit ensuite être reconstruite avec son propre entraînement PGO de production.
- `candidate.mjs` fige les sources avant de construire et vérifier un candidat,
  puis lance un screening sur Arrays, b8x, Aff, Spec et Yoga JSON.
  `snapshot.mjs` prépare les hypothèses dans `queued/` ; ces copies restent
  éditables jusqu'au gel fait par `candidate.mjs`. `sequence.mjs` sérialise une
  liste figée de constructions/mesures et conserve chaque échec séparément.
- `measure.mjs ARCHIVE LABEL CONFIG.json` réalise une chauffe et cinq paires
  (nombre configurable), dans l'ordre alterné, projets et variantes sérialisés.
  Les workers sont 8/8/8/8, pipeline actif. Chaque génération est comparée
  **octet par octet**, avec contrôle des manifestes, aux sorties figées.
- `profile.py` conserve l'échantillonnage de tous les threads, les chronologies
  et les sorties contrôlées. Les familles inclusives se chevauchent : ce ne sont
  pas des pourcentages CPU et ces profils ne servent pas de chronomètre.
  `diagnose.mjs` ajoute les temps par module/tranche et les compteurs de cache
  dans des compilateurs dédiés au diagnostic.
  `reprofile.mjs` reprend les trois cas prioritaires sur une composition figée.
  Le paramètre `--phase` fixe le début de la capture, qui continue jusqu'à la
  sortie du processus et inclut donc les phases ultérieures.
  `compress-profiles.py` archive les grands `sample.txt` en gzip, puis contrôle
  l'empreinte et la taille des octets décompressés avant de retirer le fichier
  brut. L'inventaire garde les deux empreintes ; `../gopurs-rust/attribute.py`
  lit indifféremment un échantillonnage brut ou compressé.
- `production.mjs` exige la sélection qualifiée, reconstruit les trois hôtes,
  réentraîne le PGO de production, lance les régressions et l'identité de
  bootstrap, puis mesure les 51 projets. Une seconde campagne complète apparie
  le nouvel hôte Rust avec le binaire publié figé. `validate.mjs` construit et
  exécute les applications représentatives hors chronomètre.
- `audit-runs.mjs` recalcule les médianes et totaux à partir des journaux bruts,
  vérifie les phases, l'ordre des hôtes, les paramètres, les manifestes de chaque
  génération et les empreintes des artefacts. `seal-selection.mjs` consigne la
  sélection revue et la clôture, après identité des sources installées avec le
  candidat mesuré, audit des campagnes achevées et dernière revue des profils.
  `production-contracts.mjs` vérifie les passes modifiées dans le graphe JS de
  production ainsi que la nouvelle frontière Effect du bridge Go sous `-race`.
  `resume-production.mjs` reprend une interruption sans processus restant :
  résultats initiaux et sorties partielles archivés, projets complets conservés,
  projets inachevés remesurés intégralement avec chauffe. Les empreintes des
  compilateurs restent identiques, puis la table fusionnée est réauditée.
  `validate.mjs` isole l'exécution des applications dans des copies privées des
  entrées ; ses journaux et outils sont figés par tentative. `retry-applications.mjs`
  termine la qualification après le correctif du répertoire d'exécution, avec
  précontrôle archivé, échec initial conservé et campagnes de mesures inchangées.
  `summarize.mjs` affiche les screenings en quelques lignes.

La mesure est `backend total` : chargement, préparation, PBO, émission et
drainage. Frontend, construction des exécutables, démarrage/arrêt des processus
et exécution des applications sont hors chronomètre. Les sous-phases ne sont
pas additives. Les totaux sont les **sommes des médianes par projet**.

Seuils pratiques : ≥ 1 % sur le total complet, ou ≥ 3 % sur un cas important
sans régression significative du total. Les intervalles bootstrap appariés
(ratios logarithmiques, 20 000 réplications déterministes) évaluent le bruit
observé ; un screening sur cinq projets n'est pas le total de la publication.
Les pistes prometteuses sont confirmées avant sélection. Après exploration
des familles prioritaires, trois mécanismes distincts sans nouveau gain retenu
déclenchent une dernière revue des profils et la décision de clôture.

Les échecs, propositions rejetées, binaires, sources générées et diagnostics
restent dans des répertoires nommés distincts. Ne jamais supprimer récursivement
`output/main` : sur APFS insensible à la casse, c'est aussi `output/Main`.

## Configuration de mesure

```json
{
  "rounds": 5,
  "cases": ["gopurs-arrays", "b8x", "gopurs-aff", "gopurs-spec", "gopurs-yoga-json"],
  "variants": [
    { "name": "control", "binary": "candidates/control/gopurs-rust" },
    { "name": "candidate", "binary": "candidates/candidate/gopurs-rust" }
  ]
}
```

Omettre `cases` pour mesurer les 51 projets. Une configuration archivée ne doit
pas être modifiée après exécution. La publication finale conserve les anciennes
qualifications et n'affiche les ratios `/JS` que sur b8x, le sous-total et le total.
