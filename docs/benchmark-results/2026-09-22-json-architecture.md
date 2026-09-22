# JSON decoding architecture probes: shared code, interfaces and descriptors

**A substantial improvement is technically plausible.** On the existing general
JSON corpus, the shared-engine Go prototypes decode in **0.525–0.602 ms**, versus
**35.191 ms** for current gopurs and **7.007 ms** for the same Argonaut program in
JS. Including the existing Go JSON parser, the prototypes reach **7.871–8.064 ms**,
close to JS at **8.485 ms**.

**These are hand-written architecture prototypes, not changes delivered by the
compiler.** Their result establishes a useful performance target on this corpus.
It does not establish that interfaces alone provide that gain, nor that the full
PureScript ABI can already produce these results.

## End-to-end JSON experiment

| Model | Parse (ms) | Decode (ms) | Parse + decode (ms) | Decode allocation (MiB) |
|---|---:|---:|---:|---:|
| Current gopurs | 6.834 | 35.191 | 44.385 | 56.729 |
| Shared engine, private typed structs / interfaces | 6.708 | **0.525** | **7.871** | **1.035** |
| Same engine, layout descriptors / slots | 6.659 | **0.602** | **8.064** | **1.676** |
| Current PureScript / JS | 1.518 | 7.007 | 8.485 | — |

The current official README baseline is **45.82 ms Go / 8.65 ms JS**, from the
previous integrated optimization. The fresh current-code observations above are
consistent with those figures. The prototype observations do **not** replace the
README's compiled-PureScript cells or enter its totals.

Against the fresh current Go run, the interfaces prototype reduces combined time
by **82.27% (×5.64)** and decode allocations by **98.17%**. Decode-only speed is
about **×67** relative to current Go and **×13.3** relative to JS. This large gap
includes several architectural changes together; it is not an estimate for one
compiler patch. Relative to JS, combined time is only **7.24% lower** in this
series, with three process minima of **7.517–8.237 ms** versus **8.332–8.509 ms**.
The descriptor prototype varies more, **7.373–9.868 ms**. Describe the combined
result as reaching approximately JS performance, not a universal victory.

Both output representations use the exact same `decodeSchema` engine, with one
body for all record layouts. Descriptors specify fields and their order; the
custom Event schema specifies tag-first validation and variant field order.
Private typed structs implement composed builder/view interfaces; slot records
implement those interfaces using a shared layout and a contiguous array. Only
storage allocation and adapters differ between the two models. Typed storage
reduces decoding time **12.73%** against slots here. That is much smaller than
the shared improvement against generated Argonaut.

The shared engine also changes these mechanisms relative to current gopurs:

- Direct checks replace chains of decoder dictionary lookups and closures.
- Errors travel separately from successful values; intermediate heap-allocated
  `Either`/`Maybe` constructor chains disappear.
- Each output array is allocated once and filled in order, replacing applicative
  intermediate arrays and concatenations.
- Records are filled once in private storage before publication. Completed output
  owns its arrays; it is not a lazy proxy over the input JSON.
- The same `encoding/json.Unmarshal` parser remains. The prototype omits the
  compiler's small parser callback adapters. Most time now resides in parsing.

## Correctness and scope

Both models match all **17 fixed independent oracle cases**, including complete
values and exact error paths/priorities. Five successful cases are timed;
the ten errors and two optional-field cases are checked outside timing.
An additional **317 generated cases per model** match the actual compiled JS
Argonaut decoder, covering missing fields, wrong kinds, Int limits, nested arrays,
optionals, variants, wrong roots and simultaneous errors. Fresh storage and
independence from mutable input arrays are tested.

Timing includes construction of the whole output representation. Normalization,
JSON serialization and hashing are excluded, as in the official diagnostic;
every timed result is checked afterwards. Source/binary hashes and the current
compiled baseline manifest are retained. The current baseline is rejected if
its source fingerprints no longer match.

This is a manually described schema for the existing benchmark, with no generated
schema-specialized decoder function bodies. It does not support arbitrary custom
`DecodeJson` code, arbitrary higher-order Applicative instances, general row
updates/deletion, or the full PureScript polymorphic ABI. In particular, a direct
early error return has the same checked results here but is not a valid blanket
replacement for effectful traversals. No TAST decoding or b8x gain was measured.

## Isolated record experiment

One polymorphic operation reads an Int, a String and a Boolean. It operates over
1, 8 or 32 concrete layouts with additional heterogeneous fields. The bodies are
explicitly shared and marked `noinline` to prevent accidental specialization.
All fields, types, results, retained source values and GC liveness are checked.

Medians for **32 layouts**, three 200 ms runs:

| Operation | Current Value | Typed interfaces | Descriptor / Value slots |
|---|---:|---:|---:|
| Read / compute | 12.24 ns | 5.41 ns | **1.89 ns** |
| Construct + read | 82.73 ns | **18.01 ns** | 64.19 ns |
| Immutable insert + read | 143.50 ns | 57.76 ns | **51.85 ns** |
| Construct: allocated bytes | 224 | **64** | 208 |
| Insert: allocated bytes | 336 | **24** | 208 |

The descriptor strategy wins the isolated read despite still storing boxed
Values. Interfaces win construction through native field storage. Therefore
there is no evidence that interface embedding is the uniquely best strategy.

Interface insertion retains the original immutable record through an embedded
interface and adds one field. This avoids copying but can retain more memory and
form forwarding chains under repeated extension. Only one insertion is tested;
the result must not be generalized to arbitrary updates. `CoerceToStruct` is
also measured separately as an expensive boundary case, not used as the headline
baseline for ordinary record reads.

## Code size and shared implementation

The small interface program has exactly **one `main.score` symbol** at each size:

| Concrete layouts | Stripped executable bytes | Shared operation bodies |
|---|---:|---:|
| 1 | 1,571,554 | 1 |
| 8 | 1,605,458 | 1 |
| 32 | 1,624,994 | 1 |

Growth from 1 to 32 layouts is **53,440 bytes (+3.40%)**. Methods, constructors and
metadata still grow. This demonstrates bounded sharing in the probe, not a
whole-project binary-size guarantee or a measured comparison with monomorphization.

## Protocol, artifacts and next decision

Apple M4 Pro, Go 1.27.0, Node 24.8.0. JSON: three separate processes per model,
two warm-ups and five samples per phase; median of process minima. GOMAXPROCS=1,
GOGC=100, PGO disabled. Order Go/interfaces/descriptors/JS, reversed in the middle
round. No compilation, profiler or other benchmark ran concurrently. Record
measurements use a precompiled Go benchmark binary; size builds occur afterwards.

[Raw JSON results and provenance](2026-09-22-json-architecture.json) include both
experiments, the size checks and differential validation. Scripts and source:
`/Users/0x1/Documents/htdocs/scratch/json-architecture-20260922`:

```sh
go test .
go build -pgo=off -o architecture .
python3 -B differential.py
python3 -B compare.py
cd record-probe
GOMAXPROCS=1 GOGC=100 ./probe.test -test.run '^$' -test.bench . -test.benchtime=200ms -test.count=3 -test.benchmem
python3 size_probe.py --build
```

The next engineering step is to make the compiler generate a bounded instance
of this model: preserve native record storage across a shared polymorphic
operation, with generated adapters or layout descriptors and no map conversion.
Measure actual compiled PureScript before expanding that ABI. The much larger
JSON gap additionally requires reducing success-path ADT/closure intermediates
and fusing array construction. The probes justify that architectural work, while
leaving the representation choice and obtainable integrated gain open.
