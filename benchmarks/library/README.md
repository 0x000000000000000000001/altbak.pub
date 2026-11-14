# Library microbenchmarks (excluded WIP)

These 23 original PureScript benchmarks are marked **excluded WIP** and inspired by the
library-operation families shared by **f-f on the PureScript Discord in March
2026**. The exact sources and commit of the 218-case Discord suite were not
located. These are new implementations, not a source import or a reproduction
of its reported JS/Chez/Racket timings.

The public [purescm core Racket branch](https://github.com/purescm/purescript-core/tree/b0bee39302df30afeb70560cb328d3208441f195)
was inspected during the search. Its older library benchmarks do not establish
the source or measurement protocol of the Discord table. No upstream benchmark
code was copied into this suite.

**Status:** implementation only. No compilation, test, benchmark or executable
verification was run when adding these files, at the user's request. There are
no performance measurements or qualified backend results yet.

## Running

From the repository root, the normal Go command builds and runs the 14 core
cases followed by the 23 new WIP cases:

```sh
./bin/go/run
```

The same integration is configured for JS, JS/PBO (`es`) and Rust. `./bin/run`
uses all existing backends; those without the new cases retain their existing
coverage and report the omission. The new WIP values do not enter the core total.

To run only the excluded WIP group:

```sh
./bin/go/run --wip
```

On Go, JS and Rust, `--wip` runs these 23 cases plus the three diagnostics already
marked excluded WIP in the main README: Array Indexing, JSON to Typed AST and
JSON Decoding. The latter retain their independent workspaces, protocols and
campaign results. They can also be selected individually with `--test`.
JS/PBO currently runs only the 23 library cases. `./bin/run --wip` selects all
four configured backends. This option is mutually exclusive with `--test`,
`--x`, `--ffi` and `--fficc`.
The grouped JS/Go/Rust command accepts the build/run phase flags and `--clean`,
plus `--dev` for Go. Custom artifact/build directories and Go PGO options remain
available on the ordinary run, not on this group of different protocols.

The commands also accept `--build-only` and `--run-only`. For a multi-backend
comparison, finish all builds before starting measurement:

```sh
./bin/run --wip --build-only
./bin/run --wip --run-only
```

JS/ES keep builds under `run/bak/<backend>/output-runs/pure-*` (normal run) or
`wip-*`, with the corresponding `latest-pure.json` / `latest-wip.json` pointer
and `validated.json` in the artifact directory. Go/Rust use
`run/bak/<backend>/modes/pure/` or `modes/wip/`, with `results.json`. These files
describe one process, not the median of three independent processes.

For a publication campaign, use the existing frozen-plan protocol in
[`docs/benchmark-methodology.md`](../../docs/benchmark-methodology.md), with
`"mode": "pure"` for a normal run-only command. `measure.py` validates the core
and WIP sections separately and aggregates medians per case over three
processes. The core result keeps its original shape; the new result is nested
under `library`. Keep the build manifests and payload hashes required by that
protocol. `publish.py` continues to publish the historical core totals.

A numeric plan for the 23 new cases alone uses `"mode": "wip"` and the ordinary
driver directly, e.g. `python3 bin/native/driver.py go --wip --run-only`, against
the artifact prepared by `./bin/go/run --wip --build-only`. Do not feed the public
JS/Go/Rust `--wip` group's combined console output into one numeric plan entry:
its other three diagnostics already collect their own campaigns and use
different measurement boundaries. WIP results have no aggregate score.

## Measurement boundary

`src/AppLibrary.purs` prepares all fixtures before the three global warm-ups.
Each preparation returns an `Effect Int`; the operation is performed inside that
action on every invocation. Inputs pass through `Bench.opaque` inside the action.
`Bench.runBench` provides the three local warm-ups, calibrated batches targeting
10 ms (capped at 2^24 invocations) and minimum of ten batches.

Fixture construction is excluded. The measured interval includes the operation,
the result consumer described below, input barriers, Effect/callback adapters
and the numeric result sink. Backend conversions at a generic opaque boundary
are included too. Read very small timings with that overhead in mind. Inspect
generated code before attributing an unexpected gap to the named operation alone.

Constructed arrays are consumed element by element. Constructed strings are
consumed code point by code point. This also includes deferred string work
triggered by consumption. These are operation-plus-consumption timings; no
consumer cost is subtracted. Fixture reuse is limited to immutable inputs;
mutable reference allocation is repeated inside every invocation. Runtime memory
management and normal compiler optimizations remain enabled. This protocol does
not force a final collection after each invocation or measure peak memory.

The external validator compares outputs to the independent integer oracles in
`cases.json`, in exact label order. The batch driver also checks consistency
against the warm-up result. These checksums are benchmark contracts, not a
replacement for each library's full correctness tests.

## Cases and independent oracles

Every case below is **excluded WIP**, with performance measurements pending.

Array result checksums use alternating weights 1 and 2, starting with weight 1.
For `[1 .. 2m]`, the sum is `3m² + 2m`. All intermediate sums fit signed Int32.

| Case | Measured operation and consumer | Expected |
| --- | --- | ---: |
| `array/map-10000` | Map `+1` over `[1..10000]`, weighted checksum | 75025000 |
| `array/filter-10000` | Keep even elements, weighted checksum | 37510000 |
| `array/foldl-sum-10000` | Sum `[1..10000]` | 50005000 |
| `array/sort-descending-10000` | Sort `[10000..1]` ascending, weighted checksum | 75010000 |
| `array/eq-neq-last-10000` | Compare distinct arrays differing only at the last element; false → 0 | 0 |
| `array/append-1000` | Append `[1..1000]` and `[1001..2000]`, weighted checksum | 3002000 |
| `foldable/traverse-maybe-1000` | Traverse `[1..1000]` with `Just (x+1)`, weighted checksum | 752500 |
| `foldable/traverse-maybe-fail-first-1000` | Reject value 1; `Nothing` → −1 | -1 |
| `foldable/traverse-maybe-fail-last-1000` | Reject value 1000; `Nothing` → −1 | -1 |
| `string/cu-length-unicode-1000` | UTF-16 code-unit length of `"a😀é"` repeated 1000 times | 4000 |
| `string/cp-length-unicode-1000` | Code-point length of the same recipe | 3000 |
| `string/indexOf-hit-10000` | Find `"needle"` in 10000 `a` characters followed by `"needle"` | 10000 |
| `string/indexOf-miss-10000` | Find `"absent"` in that recipe; `Nothing` → −1 | -1 |
| `string/cp-take-unicode-1500` | Take 1500 code points, sum their scalar values | 64421000 |
| `string/split-1000` | Split 1000 comma-separated `"abc"` parts, count parts plus all scalar values | 295000 |
| `string/concat-1000` | Left-fold append over 1000 `"ab"` fragments, sum all scalar values | 195000 |
| `record/set-small-1000` | 1000 immutable updates to the last of 2 fields, sum original and final fields | 1006 |
| `record/set-medium-1000` | Same with 8 fields | 1072 |
| `record/set-large-1000` | Same with 16 fields | 1272 |
| `st/ref-write-10000` | New ST ref, writes 1 through 10000, final read | 10000 |
| `st/ref-modify-10000` | New ST ref at 0, 10000 increments, final read | 10000 |
| `effect/ref-write-10000` | New Effect ref, writes 1 through 10000, final read | 10000 |
| `effect/ref-modify-10000` | New Effect ref at 0, 10000 increments, final read | 10000 |

The filter checksum is `6m² + 4m` with `m = 2500`. Mapping adds 15000 to
the 10000-element checksum; the successful traversal adds 1500 to the
1000-element checksum. Failure position does not imply that the Array
`Traversable` instance stops evaluating callbacks early.

The Unicode take returns 500 complete `a😀é` groups, so its oracle is
`500 × (97 + 128512 + 233)`. Split returns 1000 parts whose scalar values sum
to 294000, giving 295000. Concat has checksum `1000 × (97 + 98)`.
The strings are built by the documented joins/appends; no particular flat or
rope representation is assumed, and `é` is the single code point U+00E9.

Record fields start at 1 through their field count. Both the original and final
records are consumed, giving `1000 + 2 × sum(1..fieldCount)`. Normal scalar
replacement is allowed. Each action starts again from the immutable fixture.

## Source isolation

The configured drivers copy `benchmarks/library/src/` into their private `src/`
snapshot for normal runs and WIP-only runs. The extra direct dependencies
(`foldable-traversable`, `refs`, `enums`) are added to that private configuration.
Other backends and FFI/extended/individual-test modes do not compile these
sources. Sources and the oracle manifest are fingerprinted
to reject stale library artifacts and changes during collection.

To add a case, add its preparation to the appropriate family module, register
the same label and order in `AppLibrary.cases` and `cases.json`, and document the
oracle derivation and measured consumer here. Expected values must be derived
independently of the measured implementation.
