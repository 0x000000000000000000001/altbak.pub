# Go ST.Ref and Effect.Ref optimization

Le compilateur réduit les boucles d'écriture et d'incrémentation d'une référence
locale lorsque son usage complet est connu : création, appel d'un helper
vérifié, puis lecture finale. Le nouveau passage `Gopurs.RefLoopRegions`
travaille sur l'IR typé avant la conversion TCO. Il reconnaît les opérations
ST.Ref et Effect.Ref et la structure de `tailRecM`, sans dépendre des noms des
benchmarks ni de leurs bornes.

Pour une incrémentation, le résultat est la valeur initiale plus le produit du
nombre positif d'itérations et de l'incrément, avec la sémantique Int32 du
backend. Pour une écriture croissante inclusive, le résultat est la dernière
borne écrite, ou la valeur initiale si la boucle ne tourne pas. La borne
`2147483647` conserve le code original : dans ce cas, le compteur de la boucle
source déborde et la boucle ne termine pas.

Le passage n'accepte que des arguments Int déjà évalués ou des littéraux pour
éviter de déplacer, supprimer ou répéter l'évaluation d'un appel. Les paramètres
du helper sont substitués simultanément, sans confusion avec les niveaux locaux
de l'appelant. Il conserve les helpers publics et refuse les régions contenant
des usages supplémentaires de la référence, des observations intermédiaires ou
des opérations inconnues. Les actions restent différées et répétables ; la
lecture opaque de leur borne reste dans chaque invocation.

Benchmarks, harnesses, library FFI and C references are unchanged. Runtime inputs remain opaque.

## Reference workloads

| Case | Official Go baseline (µs) | Paired before (µs) | After (µs) | Speedup in this pair |
| --- | ---: | ---: | ---: | ---: |
| `st/ref-write-10000` | 7.14 | 7.11 | 0.0086 | 822.80× |
| `st/ref-modify-10000` | 20.23 | 22.03 | 0.011 | 2053.27× |
| `effect/ref-write-10000` | 194.25 | 220.24 | 0.016 | 13888.52× |
| `effect/ref-modify-10000` | 179.83 | 196.97 | 0.011 | 18587.77× |

## All 37 matched Go/C cases

| Indicator | Official baseline | Paired before | After |
| --- | ---: | ---: | ---: |
| Ratio of totals | 1.21× | 1.31× | 1.29× |
| Arithmetic mean μ | 9755.93× | 10844.96× | 132.26× |
| Median Q₂ | 8.44× | 9.12× | 7.65× |
| Geometric mean G | 23.18× | 25.19× | 10.20× |

All four indicators use the same 37 measured Go/C pairs: 14 core medians from App and 23 library medians from AppAll. The C denominator remains frozen at 8.777346938669927 ms. Differences between the official baseline and the fresh before measurement reflect measurement variation, not this compiler change.

## Measurement and validation

Three independent processes per build, in AB/BA/AB order; three global warm-ups, three local warm-ups, calibrated batches and the minimum of ten batches per process. Published rows use medians across the three processes. Release Go, GOGC=800, default GOMAXPROCS, no PGO. All builds and tests finish before measurement. The host is not isolated.

The App/AppAll campaign validates 306 measured outputs. Source and binary hashes are checked around each collection. The archived validation summary and runtime-input proofs record the targeted tests, varied inputs, repeatability and opaque-call checks.

Full commands, original outputs, compiler changes, input proofs and validation logs are preserved in [the JSON archive](2026-11-26-go-ref-loops.json) and `var/benchmark/go-ref-loops-20261126`.

## Regression review

The primary 37-case combination (App14 and AppAll23) contains these increases above 5%; they are retained in the measurements:

- `record/set-large-1000`: +10.85%.
- `string/cu-length-unicode-1000`: +10.00%.
- `string/indexOf-miss-10000`: +5.21%.
- `Polymorphism`: +5.08%.

The four increases above 5% were checked in a separate diagnostic campaign.
The generated target modules and the selected original benchmark-harness sections
are byte-identical before/after; normalization also confirms equality. The probes
call the actual generated action and unchanged `Bench.runBench`. All eight probe
builds finished before the 24 timed diagnostic processes, ordered AB/BA/AB for
each case. Library probes use AppAll; Polymorphism uses the primary App14 build.

- `record/set-large-1000`: primary +10.85%; isolated **0.464465 → 0.454062 µs**
  (**−2.24%**). The increase above 5% did not reproduce.
- `string/cu-length-unicode-1000`: primary +10.00%; isolated
  **0.010248 → 0.010162 µs** (**−0.84%**). The increase above 5% did not reproduce.
- `string/indexOf-miss-10000`: primary +5.21%; isolated
  **2.547455 → 2.522563 µs** (**−0.98%**). The increase above 5% did not reproduce.
- `Polymorphism` (App14): primary +5.08%; isolated
  **0.004349 → 0.004386 µs** (**+0.85%**). The increase above 5% did not reproduce.

All 24 diagnostic outputs pass their oracles. These results are separate evidence;
none of the primary samples, medians or four published aggregates was replaced.
The diagnostic does not prove that execution layout has no effect; it shows that
the original increases were not reproduced when these unchanged actions were
measured separately.

RBTree also matters because it dominates the sum. Its primary App14 median was
already **10,291.084 µs before** this change, versus **9,525.8125 µs** in the
official historical baseline (+8.03%). It measured **10,526.625 µs after**,
another +2.29% (+235.541 µs) within this pair. The generated RBTree and selected
harness sections are byte-identical. No isolated RBTree timing was run, so the
cause cannot be assigned confidently to noise, layout or machine state. Its
variation is retained in the total ratio, explaining why the fresh after total
must be compared with both the fresh before value and the historical baseline.

Evidence: `proofs/isolated-regressions/results.json`,
`proofs/isolated-regressions/generated-comparison.json` and
`proofs/rbtree-comparison.json`. The original generated sources, executables,
benchmark inputs and primary campaign results remain unchanged.
