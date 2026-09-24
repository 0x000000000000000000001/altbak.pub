# Direct text-to-schema application decoding

This cycle starts from the retained application schema workers and their
**3.19 ms** official combined result. Preserved sources, builds and measurements
are under `var/benchmark/json-text-schema-20260924/` (the cycle crossed midnight).
The baseline is `json-schema-20260924/published-json`.

The object-slot implementation is retained. The final official combined result
is **2.629042 ms**, published as **2.63 ms** in both benchmark READMEs, versus
**3.189625 ms** previously (**−17.6%**). Three complete-workload campaigns win
**18/18 pairs**; the final one includes the stricter callback guard and rebuilt
toolchain. The isolated object-slot comparison improves **5.2%** over the first
indexed candidate, winning **5/6**.

## Public API and implementation

`Data.Argonaut.Decode.Parser` adds:

```purescript
decodeJsonString
  :: forall a. DecodeJson a => String -> Either JsonDecodeError a

decodeJsonStringWith
  :: forall a. (Json -> Either JsonDecodeError a)
  -> String -> Either JsonDecodeError a
```

Their ordinary meaning is `parseJson text >>= decoder`, including the existing
`TypeMismatch "JSON"` syntax-error boundary. JavaScript uses that composition.

When the text ABI is available and every operation is proven,
`Gopurs.DecoderSchemas` emits a second worker
from the same resolved dictionaries and proven custom decoder bodies as the
existing DOM worker. The original getter and DOM shape guard are retained. A
stricter text guard requires standard child methods to have their expected tags;
only custom bodies already proven by the compiler can bypass those tags. Text
metadata is immutable and attached to the same callable decoder. Untagged,
opaque-callback or guard-rejected trees use the complete ordinary composition.

The first implementation validates syntax and builds an offset-only index in
one pass, then constructs final records, arrays, optionals and custom variants
through typed cursors. It builds no generic JSON tree for the recognized
application schema. Every member is syntax-checked, including ignored fields,
overflowing numbers and malformed escapes, before any decoding callback runs.

Final strings own their storage. When a public `Json` result needs a subtree,
the ordinary owned parser materializes it at that point.
This preserves its compact-object representation, insertion/duplicate handling
and public `Foreign.Object` ownership behavior. The cursor/index never escapes
as a public JSON value.

Early candidates materialized callback inputs individually. The final version
declines text specialization for every tree containing an opaque callback:
callbacks can mutate their JSON arguments, and repeated reads must preserve
the aliases created by one complete ordinary parse. An executable guard test
also checks that an opaque replacement takes the whole-composition fallback.

The `object-slots` variant additionally gathers each object's statically known
fields in one reverse traversal. Stack-local cursor slots preserve the last
normalized input key; values still decode in the proven schema/program order.
Repeated schema labels decode separately and keep their existing reverse-
insertion precedence in the final record.

The diagnostic's combined phase calls the public `decodeText` function. Parse
and already-parsed decoding remain separately measured controls. The complete
final results are retained and fingerprinted after every timed sample.

## Validation

- The indexed, object-slot and final published builds pass **2,022** full
  application/schema cases against the ordinary native decoder, including exact
  errors.
- The object-slot and final builds additionally match **829** pre-fusion archived
  native fingerprints, including the previously documented native-64-bit Int
  behavior.
- **10,012** deterministic syntax/mutation cases, bit-exact numeric reads and
  six nesting-depth boundaries agree with `encoding/json`.
- Complete-result input-overwrite, concurrent decoding and syntax-before-custom-
  callback checks pass under `-race`.
- Generated DOM/text workers preserve duplicate schema labels and owned strings
  in an executable Go regression under `-race`.
- The codec fixture passes in JavaScript and generated Go under `-race`, covering
  standard schemas, proven custom branches, arbitrary callbacks, public `Json`
  and `Foreign.Object Json`, ignored invalid members and duplicate input keys.
- The existing record-plan callback/symbol contract suite passes under `-race`.
- The final compiler schema suite passes **9/9 tests**, including executable
  emitted-worker and callback-guard regressions.

The public-ownership audit exercises a specialized record containing an escaping
`Json`, and a fallback record returning a `Foreign.Object Json`. It overwrites
the input buffer and mutates one returned map to verify independent storage.
Two initial test assumptions were corrected: a foreign-object schema can decline
specialization, and object maps can contain boxed JSON values, so equality uses
the public JSON conversion. The failed test workspaces/logs are preserved.

## Initial paired measurement

The indexed candidate's first six-pair campaign reports:

| Phase | Previous production | Indexed text worker |
| --- | ---: | ---: |
| Parse control | 2.255396 ms | 2.250188 ms |
| Decode control | 0.824209 ms | 0.796584 ms |
| Complete text-to-value | 3.082896 ms | 2.743708 ms |

Combined elapsed time improves **11.0%**, winning **5/6 pairs**. Minimum combined
allocations decrease **4,311,856 → 3,483,048 B (−19.2%)**. Parse and decode
allocation minima remain **2,435,696 / 1,876,160 B** respectively.

## Retained object-slot measurements

| Campaign | Baseline combined | Object slots combined | Gain | Winning pairs |
| --- | ---: | ---: | ---: | ---: |
| Previous production vs object slots | 3.100167 ms | 2.467146 ms | 20.4% | 6/6 |
| Confirmation | 3.002625 ms | 2.465708 ms | 17.9% | 6/6 |
| Final strict-guard build vs previous production | 3.634396 ms | 2.593605 ms | 28.6% | 6/6 |
| Indexed candidate vs object slots | 2.616459 ms | 2.481375 ms | 5.2% | 5/6 |

Object slots add no measured allocation: combined minima remain **3,483,048 B**.
These comparisons distinguish the complete text path's benefit from the
incremental change to object lookup.

All campaigns use preserved, fingerprinted binaries, `GOMAXPROCS=1`, `GOGC=100`,
alternating process order, all six phase orders, two warmups and five retained
samples per phase. Each six-pair application campaign validates **900 sampled
outputs**. Final construction, conversion, copying and normal GC are timed;
fingerprinting and retention-buffer allocation are outside the interval.
Each process reports the best of five post-warmup samples; campaign results are
the medians of these process observations. Background-load variation is visible
in the final campaign, especially its baseline, so its 28.6% is reported
separately from the earlier paired gains and the official measurement.

## Final official measurement

| Phase | Go | JavaScript validation run |
| --- | ---: | ---: |
| Parse | 2.245208 ms | 1.877459 ms |
| Decode already-parsed JSON | 0.855458 ms | 8.078459 ms |
| Complete text-to-value | **2.629042 ms** | 8.791834 ms |

These are the medians of three fresh official runs. Go allocation minima are
**2,435,696 / 1,876,160 / 3,483,048 B** for parse/decode/combined. Phase medians
come from separate executions with different allocator histories and must not
be added. The new public composition is also exercised by the JavaScript run.
The three diagnostic rows remain excluded from totals and `/C` ratios.

## Construction, lifetime and release

Six alternating pairs of fresh processes construct the public decoder, retain
two complete five-input output batches, fingerprint them, then measure the final
collection after release. An untimed preparatory collection clears fingerprint
scratch while both batches remain alive. Input loading, fingerprinting,
retention-buffer allocation and that preparatory collection are outside the
measured intervals.

| Measured interval | Previous production | Final text worker |
| --- | ---: | ---: |
| Public getter construction | 53.854 µs | 72.3955 µs |
| First combined corpus decode | 3.812604 ms | 2.705917 ms |
| Second combined corpus decode | 3.733854 ms | 3.010521 ms |
| Final release collection | 0.436792 ms | 0.448396 ms |
| Per-process interval subtotal, then median | **8.035020 ms** | **6.245001 ms** |

The subtotal improves **22.3%**, winning **6/6 pairs**. It is computed within
each process before taking the median, rather than by adding phase medians.
Getter construction allocations increase **42,696 → 79,040 B**. The slightly
higher construction and final-collection costs are included in the subtotal.
Package initialization remains **10,120 B / 43 allocations** in separate
`inittrace` inventories. Heap-allocation medians with both result batches alive
are **5,306,152 → 5,281,944 B**; after release they are **813,928 → 825,456 B**.
The lifecycle audit validates **140 retained outputs**, including its separate
initialization-inventory processes.

## Distribution and bootstrap

The final application executable changes **6,645,186 → 7,001,506 B (+5.36%)**.
Its generated Go source changes **8,436,829 → 8,613,556 B** across **326 → 327 files**.
The `go.mod` files are byte-identical; the text path adds no external dependency.
The native compiler changes **28,980,610 → 29,174,738 B (+0.67%)**, relative to
the compiler preserved at this cycle's start.

The rebuilt native compiler regenerates **569/569 Go files byte-identically**
to Node in its preserved bootstrap workspace.

The first publication attempt stopped before timing because the external
PureScript compiler binary had changed. `publication-source-changes.json`
records the detected hashes; the final build uses the new toolchain and the
stricter callback guard.

Final executable SHA-256:
`d37d08ce58337bed932af002dac8359a59607973066f6c457d8435c8e2719109`.
Rebuilt native compiler SHA-256:
`63e0da0701c126c8457300285d942c688ab7ae368b1a53b52833e7197b6b9969`.

## Reproduction and archive verification

The final preserved diagnostic is `json-text-schema-20260924/published`.
`validation-final`, `ownership-published`, `bootstrap-final-parity`,
`paired-final`, `official-final`, `official-final-js` and `lifecycle-published`
contain its validation, paired timing and lifecycle evidence. The earlier
`indexed` and `object-slots` builds remain preserved with their own manifests.

```sh
python3 bin/benchmark/json-diagnostic/text-schema/finalize.py \
  --cycle var/benchmark/json-text-schema-20260924
```

This checks archived sources and binaries, all five paired campaigns, official
results, generated-source validation parity, bootstrap regeneration and
lifecycle aggregates. The campaigns, official runs and lifecycle audit account
for **5,090 retained, fingerprint-checked sampled outputs**, in addition to
warmups and the semantic/ownership suites. The final source archive records the
integrated local implementation and both README updates.
