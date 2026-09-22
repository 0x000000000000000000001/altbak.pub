# Comparing the JSON performance directions

**Recommended direction: a native calling convention with shared polymorphic
workers, stable record storage, and fused construction.** Interfaces are one
adapter mechanism in that design, not a universal representation to impose.
The biggest isolated architectural effect comes from removing Value-wrapped
arguments/results, rather than merely replacing dictionary lookup. The strongest
concrete change tested on existing generated Go is array-decoder fusion.

All changes in this campaign are **scratch prototypes**. No new compiler or
library optimization was integrated. The official README cells stay unchanged.

## 1. Actual generated Go: what transfers immediately

The existing general JSON diagnostic was rebuilt from the same generated Go,
with narrowly scoped edits. No replacement schema decoder was used here.

| Generated-code variant | Decode (ms) | Parse + decode (ms) | Decode allocations (MiB) |
|---|---:|---:|---:|
| Current Go, unchanged | 38.327 | 45.364 | 56.729 |
| Case application distributed + beta reduction | 35.124 | 44.020 | 54.489 |
| Same, forwarding an existing Left | 35.090 | 45.266 | 54.489 |
| Fused array decoding | **26.939** | **35.193** | 42.107 |
| Array fusion + beta/Left forwarding | 27.425 | **34.709** | **39.867** |
| Current JS | 7.307 | 9.051 | — |

**Array fusion reduces combined time 22.42% and decode allocations 25.77%.**
It replaces the composition/traversal machinery of `decodeArray` with one
ordered loop and one output array. The supplied decoder, input/output ABI and
error constructors remain unchanged. Every callback is still evaluated in order,
including after an earlier failure; only the first error is returned. This is
not an unconditional early-exit rewrite of generic Applicative traversal.

The beta prototype removes an immediately applied closure pattern inside
`gDecodeJsonCons` (three closure construction sites and two generic Apply sites).
Its combined improvement is **2.96%**, with timing ranges overlapping. The
existing-Left variant has the same successful path and allocations; its different
timings are noise, not evidence that forwarding Left hurts valid inputs.

The combination observes **23.49% lower combined time** and **29.72% fewer decode
bytes**. It provides no convincing additional throughput win over fusion alone
in this small series, despite the extra allocation reduction. Improvements cannot
be added. All four modified executables are no larger than baseline: baseline
6,458,962 bytes; beta 6,458,658; fusion 6,420,482; combination 6,420,194.

Historical README baseline: **45.82 ms Go / 8.65 ms JS**. Current controls are
45.364 / 9.051 ms. The fresh paired controls above govern attribution. Even the
combined generated-code prototype remains about **3.8× slower than JS** here;
these two changes do not deliver the previous full architectural prototype's gain.

## 2. Native dictionaries versus the argument/result ABI

All four modes below share one schema decoder and the same native record storage.
Dictionaries and their closures are built once, outside timing. The middle two
Value modes use identical argument/result boxing, isolating lookup and dispatch.

| Call boundary | Decode (ms) | Decode allocations (MiB) |
|---|---:|---:|
| Direct shared decoder | 0.580 | 1.035 |
| Native dictionary, native `(value, error)` result | **0.613** | **1.035** |
| Native dictionary, Value input/output retained | 2.664 | 4.659 |
| Boxed dictionary + RecordGet/Apply, same Value ABI | 2.835 | 4.659 |

Replacing only RecordGet/Apply saves **6.02%** in this probe, with no allocation
change. Removing the Value wrappers at that same native dictionary boundary
improves decode time by **×4.35 (−77.00%)**, with **77.77% fewer allocated bytes**.
This supports prioritizing the calling convention and result representation.
Native dictionaries alone retain nearly all of the boxing cost if their methods
continue to accept/return the same wrapped values.

This is a controlled boundary experiment using the real current runtime's
Value/Any/RecordGet/Apply. The boxed result is an opaque native tuple, not the
exact generated Either layout. The ratio is not a measured whole-gopurs gain.
Its purpose is to separate storage/lookup from representation changes.

## 3. Array construction and transient result allocations

A second shared decoder independently toggles array construction and heap-boxed
recursive results. Output record storage, schemas and input JSON stay constant.

| Array construction / recursive result | Decode (ms) | Decode allocations (MiB) |
|---|---:|---:|
| Direct fill / value result | **0.656** | **1.035** |
| Concatenation tree / value result | 1.141 | 2.714 |
| Direct fill / boxed result | 1.013 | 2.847 |
| Concatenation tree / boxed result | 1.502 | 4.526 |

With value results, direct array construction is **42.50% faster** and allocates
**61.85% fewer bytes** than the concatenation tree. With direct arrays, returning
the recursive result by value is **35.22% faster** and allocates **63.63% fewer
bytes** than the box model. Together the direct/value variant is **56.34% faster**
than split/boxed; these percentages overlap and cannot be summed.

The split uses the production traversal's leaf sizes and pivot, but omits its
generic Applicative machinery. The boxed model deliberately uses a non-inlined
pointer-returning constructor to force a heap allocation, confirmed by escape
analysis; the value model has a matching non-inlined boundary. It is a mechanism
study, not an exact count/size of generated Either or Maybe objects. Optional
field storage is unchanged: **a separate Maybe-specific gain is not established**.
The direct model may stop on an error because its decoder is pure. Tests expose
the different post-error callback counts; such early exit is not claimed safe
for arbitrary effectful traversal. The actual generated-Go fusion above retains
those calls instead.

## 4. Interfaces versus descriptors

The earlier storage alternatives were rerun in the same campaign:

| Shared decoder output storage | Decode (ms) | Combined (ms) | Decode allocations (MiB) |
|---|---:|---:|---:|
| Private typed structs / interfaces | **0.548** | 8.501 | **1.035** |
| Layout descriptors / generic slots | 0.708 | 8.324 | 1.676 |

Native typed storage improves decoding and memory. Combined times are dominated
by parsing and vary enough that they do not establish a winner. Both again reach
approximately JS combined performance. These still are manually described schema
prototypes, not compiler-produced implementations of arbitrary DecodeJson instances.

The prior isolated read study also favored indexed descriptors (1.89 ns versus
5.41 ns through interfaces at 32 layouts), while construction favored typed structs
(18.01 ns versus 64.19 ns). Hence the recommended hybrid: direct typed fields where
known, shared layout/accessor metadata for open records, interfaces where they
provide a cheap stable view. Avoid generic Value reconstruction at every boundary.

### Repeated immutable updates expose a wrapper-chain problem

A final microbenchmark replaces the same existing score field 1/8/32 times,
across 32 layouts. It times only reading four final fields; construction is
outside timing. The interface variant wraps the previous record at every update.

| Immutable replacements | Current Value (ns) | Chained interface (ns) | Flat descriptor/slots (ns) |
|---|---:|---:|---:|
| 1 | 19.76 | 52.05 | 2.37 |
| 8 | 21.12 | 153.40 | 2.83 |
| 32 | 20.11 | 525.30 | 2.85 |

Three 100 ms runs per cell, medians, zero read allocations. Forwarded field
access becomes progressively slower, and the final wrapper retains all previous
wrappers (updates + 1 links, including initial score insertion). Tests verify
all seven fields, constant width, prior-version immutability and structural
reachability after GC. This is not a retained-heap-size measurement.

The problem is **naive chaining**, not a necessary cost of every interface design.
Replacement here is not a Row.Lacks insertion. A production design must flatten
or avoid such chains; forcing this wrapper model onto all records would create
a regression. This strengthens the case for choosing native storage and its
adapters separately, and for flat layouts where updates require them.

## 5. Integration order and limits

1. **Integrate a general, semantically guarded array-construction fusion.** The
   generated-code experiment is the strongest near-term evidence. Preserve strict
   callback order and first-error priority; prove the relevant Applicative instance.
2. **Establish a native worker boundary for shared polymorphic code.** Keep typed
   records and native success/error results across known calls, with explicit
   adapters at genuinely generic/FFI boundaries. Use native dictionaries, but do
   not expect the lookup change alone to close the gap. Keep code bodies shared;
   layout/accessor metadata or methods may differ per shape.
3. **Eliminate remaining immediately consumed closures and ADT wrappers across
   those boundaries.** The specific beta pattern is useful but insufficient alone.
   Escaping closures/values and unknown callbacks need a valid generic fallback.
4. **Choose record adapters per use, rather than replacing everything with
   interfaces.** Respect row polymorphism, immutable updates, layout/field types,
   cross-module ABI and FFI. Avoid unbounded forwarding chains and retained old
   wrappers. Measure code size on actual generated applications.

The previous size probe retained one shared operation body for 1/8/32 layouts,
with only 3.40% stripped-binary growth, but it is not a complete compiler-size
guarantee. Neither the architectural schemas nor the hand-edited generated Go
prove integration into the compiler. No new TAST or b8x run was performed. The
parser remains unchanged and becomes the dominant cost near the architectural
target; only then is replacing it a likely large end-to-end lever.

## Validation and reproducibility

Sixteen variants, three processes each, two warm-ups and five samples per phase;
median of process minima. GOMAXPROCS=1, GOGC=100, PGO disabled. Apple M4 Pro,
Go 1.27.0, Node 24.8.0. Round two reverses order; round three rotates it. No
compilation, tests, profiler or competing benchmark ran during these 48 processes.
Complete output hashes are checked outside timing for all timed results.

All variants match the **17 fixed oracle cases**. Prototype mechanisms also match
**317 JS differential cases**. Generated-code changes match current Go on all
317 cases; ten cases outside signed 32-bit Int bounds already differ between
current Go and JS. That pre-existing native-64-bit difference is documented in
gopurs testing exclusions and is unchanged by these probes; it is not hidden as
a new successful JS parity check. None of the timed inputs has that discrepancy.

Actual array fusion is compared with the original function on sizes 0,1,2,3,4,7,
8,9,31 and multiple error positions, checking callback traces and first errors.
The abstract traversal model covers sizes 0–65, all first-error positions and
later errors. Ownership and output materialization checks are retained.

[Raw results, patches, scripts and provenance](2026-09-22-json-options.json).
Local executables and logs are under
`/Users/0x1/Documents/htdocs/scratch/json-options-20260922`.
`python3 compare.py` reruns the controlled campaign after the documented builds;
`verify-generated.py` checks both fused generated-code variants against current
Go. Per-experiment READMEs describe their models and limitations. The earlier
[architecture report](2026-09-22-json-architecture.md) covers the storage and size
experiments in detail.
