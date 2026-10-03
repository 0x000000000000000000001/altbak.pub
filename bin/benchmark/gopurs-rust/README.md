# Optimisation de gopurs hébergé en Rust

Tous les candidats de ce dossier exécutent **gopurs et produisent du Go**.
Ils utilisent le fork `purescript-backend-optimizer-gopurs` ; Purust sert au
bootstrap du compilateur.

Depuis `altbak.pub`, archive du lot :

```sh
A=var/benchmark/gopurs-rust-optimization-20261003
S=var/benchmark/gopurs-purust-aff-20261002
node bin/benchmark/gopurs-rust/experiment.mjs init "$A" SHA256_RUST_QUALIFIE
python3 bin/benchmark/gopurs-rust/profile.py "$A/baseline/compiler/bin/gopurs-rust" "$S" "$A/profile-baseline"
python3 bin/benchmark/gopurs-rust/attribute.py "$A/profile-baseline/sample.txt" > "$A/profile-baseline/attribution.json"
node bin/benchmark/gopurs-rust/experiment.mjs build "$A" NOM
node bin/benchmark/gopurs-rust/compare.mjs "$S" "$A/NOM-runs" "$A/NOM.json"
node bin/benchmark/gopurs-rust/verify.mjs "$A/verification.json" "$A/NOM-runs"
```

Chaque nom de candidat et de campagne doit être neuf. `init` exige l'empreinte
SHA-256 de la référence Rust qualifiée, puis fige sources,
exécutables des trois hôtes, frontend et bootstraps Purust. `build` prend un
instantané des sources gopurs/PBO courantes et conserve son Rust généré,
exécutable, commandes et diagnostics ; un workspace Cargo réutilisable garde
les dates des sources inchangées. Il ne publie pas de binaire de production.
Le profil est O3, sans LTO ni debug, Arc et mimalloc.

`relink.mjs ARCHIVE SOURCE_CANDIDATE NEW_CANDIDATE thin|false` isole un
changement de profil Cargo : il reprend exactement le Rust généré du candidat
source, fige les fichiers et conserve son propre binaire et ses logs. Chaque
reprise après échec utilise un nouveau nom. Un profil sélectionné doit ensuite
être appliqué au build de production ; `qualify.mjs` vérifie le réglage Cargo
effectif en plus de l'identité des sources générées.

Exemple de sélection, chemins relatifs au JSON :

```json
{
  "rounds": 5,
  "variants": [
    { "name": "before", "binary": "baseline/compiler/bin/gopurs-rust" },
    { "name": "candidate", "binary": "candidates/NOM/gopurs-rust" }
  ]
}
```

Une variante peut utiliser `directory` (checkout gopurs ou installation figée)
à la place de `binary`, avec `environment: { "GOPURS_JS": "1" }` pour JS,
`{ "GOPURS_RUST": "1" }` pour Rust, et aucun sélecteur pour Go. Le harnais copie
le launcher et son support runtime. `resources: true` ajoute `/usr/bin/time -l`.

Le protocole habituel fixe les jobs chargement/préparation/PBO/émission à 8 et le pipeline
à 1. Une chauffe puis les tours sont sérialisés, premier candidat tournant.
Chaque sortie, y compris les chauffes, est comparée exactement aux 294 fichiers
Go/manifests de l'oracle original. Les sorties et caches de compilation sont
neufs par processus. Les temps proviennent de `backend total` ; les builds,
frontend, exécution applicative et démarrage/sortie du processus sont exclus.
Les profils instrumentés restent des diagnostics distincts.
`attribute.py` rattache les échantillons de collections et de copies au plus
proche appelant gopurs/PBO ; ces familles inclusives se recouvrent et ne sont
pas des pourcentages CPU.

Pour mesurer le comportement du launcher sans réglages forcés, une sélection
peut ajouter `"launcherDefaults": true`. Le harnais retire alors les variables
héritées sans injecter de nombres de workers ni de valeur du pipeline ; seuls
les éventuels `environment` propres aux variantes sont appliqués. Le protocole
archivé indique explicitement ce mode. La correction des valeurs par défaut du
3 octobre est qualifiée par `qualify-defaults.mjs` : ancien launcher figé,
reconstruction des hôtes, chemin public Rust/Aff `-c`, tests CLI/FFI et cinq
tours comparant les défauts aux huit workers explicites, avec vérification exacte.

Lancer les campagnes longues en arrière-plan avec redirection vers des logs
durables. Sérialiser les builds, profils, tests lourds et chronométrages.
`confirmation-protocol.json` dans l'archive fixe les confirmations avant la
sélection finale. La qualification de production utilise aussi
`../gopurs-aff/qualify-hosts.mjs` (bootstrap, Aff et régressions des trois hôtes).

## Qualification et confirmation finales

```sh
# Reconstruit JS/Go, puis exécute le chemin public Rust/Aff avec -c.
node bin/benchmark/gopurs-rust/qualify.mjs "$A" ast-index
# Une reprise explicite conserve les anciens logs et le diagnostic d'échec :
# node bin/benchmark/gopurs-rust/qualify.mjs "$A" ast-index --resume

# Exécuter ces trois campagnes successivement, jamais en parallèle.
node bin/benchmark/gopurs-rust/compare.mjs "$S" "$A/primary-runs" "$A/primary.json"
node bin/benchmark/gopurs-rust/compare.mjs "$S" "$A/common-runs" "$A/common.json"
node bin/benchmark/gopurs-rust/compare.mjs "$S" "$A/resources-runs" "$A/resources.json"
node bin/benchmark/gopurs-rust/verify.mjs "$A/verification-final.json" \
  "$A/baseline-aa-runs" "$A/scheduler-runs" "$A/memo-runs" \
  "$A/directives-runs" "$A/tast-runs" "$A/ast-index-runs" \
  "$A/ast-index-confirmation-runs" "$A/primary-runs" \
  "$A/common-runs" "$A/resources-runs"
node bin/benchmark/gopurs-rust/publish.mjs "$A"
```

`primary.json` contient quinze paires Rust avant/après ; `common.json`, dix
tours JS/Go/Rust-avant/Rust-final ; `resources.json`, trois paires processus
entier avec `/usr/bin/time -l`. Chacune commence par une chauffe par variante.
La qualification compare aussi le Rust de production au candidat sélectionné,
le bootstrap Purust JS/natif, les trois exécutions Aff et les fixtures CLI/FFI.
`publish` relit les justificatifs et les empreintes, puis produit le rapport
Markdown/JSON et actualise uniquement la ligne `gopurs-aff` du README.

Les contrats différentiels natifs sont exécutés par
`gopurs/gopurs/tools/test-native-pbo.mjs` contre les crates Rust conservées du
compilateur gopurs. Ses huit suites réutilisent les oracles PureScript et
vérifient les FFI du fork **gopurs** (cache, directives, usage, annotations,
modules, tables de types, maps et comparaisons `Qualified Ident`). Le contrat
des maps couvre aussi `insertWith`, son ordre de fusion et les versions
persistantes. `check-pbo.mjs` exécute les quinze suites
sémantiques PBO sur le JavaScript compilé correspondant.

### Passe allocations du 3 octobre

Le lot `gopurs-rust-allocation-20261003` reprend la référence aux défauts publics
corrigés. `finish-allocation.mjs ARCHIVE selection` exécute les replis Go/JS,
les suites PBO et la sélection composée ; le mode `confirm`, après qualification
de production, enchaîne quinze paires principales, dix tours communs, cinq
tours aux défauts publics et trois paires de ressources. Les processus sont
inventoriés entre campagnes ; la vérification finale relit tous les dossiers
`*-runs`, y compris les essais de profil de liaison. `publish-allocation.mjs`
publie ensuite les preuves vérifiées et la seule ligne `gopurs-aff` du tableau.
