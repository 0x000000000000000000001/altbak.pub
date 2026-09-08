**Reproduction de l’audit isolé**

Les commandes suivantes écrivent uniquement dans ce dossier et les caches indiqués. Le Go altbak existant est une dépendance locale en lecture ; ne pas lancer les scripts de build altbak/gopurs pour reproduire ces mesures.

```sh
cd /Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908
export GOCACHE=/private/tmp/altbak-solod-gocache
export GOMODCACHE=/private/tmp/altbak-solod-modcache
export GOGC=800
go build -o audit ./
./audit > measurements.jsonl
go test -run TestArrayCounterfactual -v -count=1
go test -run '^$' -bench BenchmarkArray -benchmem -benchtime=150ms -count=7
go test -run TestLastClosureRetention -v -count=1
go test -run 'TestSaturatedCallAllocations|TestValueToAnyFloatCompatibility' -v -count=1
go test -race -run TestConcurrentFunctionWrapping -count=1
```

Le dernier test doit échouer avec une data race dans le runtime audité. La sonde ValueToAny est conçue pour enregistrer le panic actuellement observé. Ce ne sont pas des tests validant la correction de ces deux comportements.

Profil d’allocations d’un test, hors chronométrage :

```sh
go build -o pprof cmd/pprof
./audit -test ArrayOps -profile ArrayOps.pprof -count 1
./pprof -top -nodecount=25 -nodefraction=0 -relative_percentages -sample_index=alloc_space -focus=gopurs/output/ -base ArrayOps.pprof.before ./audit ArrayOps.pprof
./pprof -top -nodecount=25 -nodefraction=0 -relative_percentages -sample_index=alloc_objects -focus=gopurs/output/ -base ArrayOps.pprof.before ./audit ArrayOps.pprof
./pprof -list=Call_Test_ArrayOps -sample_index=alloc_space -base ArrayOps.pprof.before ./audit ArrayOps.pprof
```

`pprof` a été construit dans le dossier d’audit depuis la bibliothèque de commandes Go installée, car cette installation n’exposait pas `go tool pprof`. Aucun téléchargement ni installation globale. Les autres profils utilisent RBTree, StateMonad, Primes, ListOps et Church à la place d’ArrayOps.

Les profils bruts sans filtrage contiennent aussi le coût d’écriture du profil de référence. Utiliser les versions `*.focused.txt` pour l’attribution au programme. Les fichiers `*.lines.txt` permettent de rattacher les octets aux lignes ; ne pas interpréter leur pourcentage global comme une part du seul programme.

Le témoin `array_probe.go` contient une copie de la fonction originale avec seulement renommage et qualification des imports, suivie de la variante supprimant l’aller-retour `[]Value → []int64 → []Value`. La fonction originale du module altbak reste appelée comme référence. Aucun correctif de compilateur n’est présent.
