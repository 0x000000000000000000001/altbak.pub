# Library microbenchmarks

These 23 **Incl WIP** cases measure common array, string, record and reference
operations. All nine JS/Go columns were measured in the
[November 14 campaign](benchmark-results/2026-11-14-library-js-go.md).
They have their own total in the README; they do not change the 14-case Stable
total. Rust and the other families have no published results for these cases.

## Running

From the repository root:

```sh
./bin/go/run                       # 14 core tests
./bin/go/run --wip-included        # Core, library cases and separate diagnostics
./bin/go/run --dev --wip-included  # Go development profile
```

Use `--ffi` or `--fficc` with JS/Go to select functional or imperative native
references. These variants run the core and library cases without the three
separate diagnostics. JS/PBO and psgo also run the core and library cases.
`--wip-included` cannot be combined with `--test` or `--x`.

For a multi-backend campaign, finish all builds before measuring:

```sh
./bin/run --wip-included --build-only
./bin/run --wip-included --run-only
```

The public JS/Go/Rust command also launches Array Indexing, JSON to Typed AST
and JSON Decoding, each with its own protocol. Its combined console output is
**not** one input to the numeric collector. Use the ordinary driver in a frozen
numeric plan, for example:

```sh
python3 bin/native/driver.py go --wip-included --run-only
```

The plan's mode must match `pure`, `ffi` or `fficc`. The collector validates
core and library results separately and takes medians over three processes.
`publish.py --wip-only` updates the nine JS/Go library columns and their totals;
it requires all nine columns and preserves the core measurements.

## What is timed?

Immutable fixtures are prepared before timing. Every measured call gets its
inputs through `Bench.opaque`, performs the operation, consumes the result and
returns a checked integer. Mutable references are created again on every call.

The interval includes callbacks, Effect adapters, representation conversions,
input barriers and result consumption. Arrays are read element by element;
strings are consumed by code point, including any deferred string work.
Consumer cost is not subtracted. Very small timings can be dominated by these
boundaries, so inspect generated code before attributing a gap to one operation.

Each process performs three global and three local warm-ups, calibrates batches
toward 10 ms (up to 2^24 calls), and retains the minimum of ten batches. Runtime
memory management stays enabled; there is no forced GC after every call and
these measurements do not report peak memory.

## Native reference variants

The nine columns are official JS, JS/PBO, JS FP/imperative FFI, gopurs
Release/Dev, psgo and Go FP/imperative FFI. They share labels, fixtures and oracles.
Native implementations live beside each case in `<Name>FFI.*` or
`<Name>FFICheatcode.*`.

Functional references preserve callbacks, explicit `Maybe` traversal and
immutable record updates. Code-point length materializes the same code-point
array as PureScript; constructed strings are consumed through that array.
Imperative references may use loops and local mutation, but must construct and
consume the required results. Failure-position traversals still visit every
element; their labels do not imply early termination.

The shared wrapper calls `makeAction :: Int -> Effect (Int -> Int)`: fixture
preparation is outside timing; the returned function performs each invocation.
The psgo WIP loader adapts `Data.String.CodeUnits.length` to count UTF-16 units,
so the Unicode fixture returns 4000 rather than 3000. This adaptation is recorded
in the build manifest.

## Cases and independent oracles

Expected values come from
[`library-cases.json`](../bin/benchmark/library-cases.json), independently of the
measured implementation. The validator checks exact labels, order and results.
These checksums complement each library's correctness tests.

| Case | Measured operation and consumer | Expected |
| --- | --- | ---: |
| [`array/map-10000`](../src/Test/ArrayMap10000/ArrayMap10000.purs) | Map `+1` over `[1..10000]`, weighted checksum | 75025000 |
| [`array/filter-10000`](../src/Test/ArrayFilter10000/ArrayFilter10000.purs) | Keep even elements, weighted checksum | 37510000 |
| [`array/foldl-sum-10000`](../src/Test/ArrayFoldlSum10000/ArrayFoldlSum10000.purs) | Sum `[1..10000]` | 50005000 |
| [`array/sort-descending-10000`](../src/Test/ArraySortDescending10000/ArraySortDescending10000.purs) | Sort `[10000..1]` ascending, weighted checksum | 75010000 |
| [`array/eq-neq-last-10000`](../src/Test/ArrayEqNeqLast10000/ArrayEqNeqLast10000.purs) | Compare distinct arrays differing only at the last element; false → 0 | 0 |
| [`array/append-1000`](../src/Test/ArrayAppend1000/ArrayAppend1000.purs) | Append `[1..1000]` and `[1001..2000]`, weighted checksum | 3002000 |
| [`foldable/traverse-maybe-1000`](../src/Test/TraverseMaybe1000/TraverseMaybe1000.purs) | Traverse `[1..1000]` with `Just (x+1)`, weighted checksum | 752500 |
| [`foldable/traverse-maybe-fail-first-1000`](../src/Test/TraverseMaybeFailFirst1000/TraverseMaybeFailFirst1000.purs) | Reject value 1; `Nothing` → −1 | -1 |
| [`foldable/traverse-maybe-fail-last-1000`](../src/Test/TraverseMaybeFailLast1000/TraverseMaybeFailLast1000.purs) | Reject value 1000; `Nothing` → −1 | -1 |
| [`string/cu-length-unicode-1000`](../src/Test/StringCodeUnitsLengthUnicode1000/StringCodeUnitsLengthUnicode1000.purs) | UTF-16 code-unit length of `"a😀é"` repeated 1000 times | 4000 |
| [`string/cp-length-unicode-1000`](../src/Test/StringCodePointsLengthUnicode1000/StringCodePointsLengthUnicode1000.purs) | Code-point length of the same recipe | 3000 |
| [`string/indexOf-hit-10000`](../src/Test/StringIndexOfHit10000/StringIndexOfHit10000.purs) | Find `"needle"` in 10000 `a` characters followed by `"needle"` | 10000 |
| [`string/indexOf-miss-10000`](../src/Test/StringIndexOfMiss10000/StringIndexOfMiss10000.purs) | Find `"absent"` in that recipe; `Nothing` → −1 | -1 |
| [`string/cp-take-unicode-1500`](../src/Test/StringTakeCodePointsUnicode1500/StringTakeCodePointsUnicode1500.purs) | Take 1500 code points, sum their scalar values | 64421000 |
| [`string/split-1000`](../src/Test/StringSplit1000/StringSplit1000.purs) | Split 1000 comma-separated `"abc"` parts, count parts plus all scalar values | 295000 |
| [`string/concat-1000`](../src/Test/StringConcat1000/StringConcat1000.purs) | Left-fold append over 1000 `"ab"` fragments, sum all scalar values | 195000 |
| [`record/set-small-1000`](../src/Test/RecordSetSmall1000/RecordSetSmall1000.purs) | 1000 immutable updates to the last of 2 fields, sum original and final fields | 1006 |
| [`record/set-medium-1000`](../src/Test/RecordSetMedium1000/RecordSetMedium1000.purs) | Same with 8 fields | 1072 |
| [`record/set-large-1000`](../src/Test/RecordSetLarge1000/RecordSetLarge1000.purs) | Same with 16 fields | 1272 |
| [`st/ref-write-10000`](../src/Test/STRefWrite10000/STRefWrite10000.purs) | New ST ref, writes 1 through 10000, final read | 10000 |
| [`st/ref-modify-10000`](../src/Test/STRefModify10000/STRefModify10000.purs) | New ST ref at 0, 10000 increments, final read | 10000 |
| [`effect/ref-write-10000`](../src/Test/EffectRefWrite10000/EffectRefWrite10000.purs) | New Effect ref, writes 1 through 10000, final read | 10000 |
| [`effect/ref-modify-10000`](../src/Test/EffectRefModify10000/EffectRefModify10000.purs) | New Effect ref at 0, 10000 increments, final read | 10000 |

<details>
<summary>Oracle derivations and fixture details</summary>

Array checksums alternate weights 1 and 2. For `[1 .. 2m]`, the checksum is
`3m² + 2m`. Filter uses `6m² + 4m` with `m = 2500`; map adds 15000 to the
10000-element checksum; successful traversal adds 1500 to the 1000-element
checksum. All intermediate sums fit Int32.

The Unicode fixture repeats `a😀é` 1000 times (`é` is U+00E9). Taking 1500 code
points returns 500 complete groups, hence `500 × (97 + 128512 + 233)`.
Split's 1000 parts contain scalar values summing to 294000, giving 295000.
Concat returns `1000 × (97 + 98)`. No flat or rope representation is assumed.

Record fields start at 1 through their field count. Both original and final
records are consumed: `1000 + 2 × sum(1..fieldCount)`. Each invocation restarts
from the immutable fixture; normal scalar replacement remains allowed.

</details>

## Source selection

Cases live in `src/Test/<Name>/`; shared helpers are in `src/Bench/`.
[`AppLibrary`](../src/AppLibrary.purs) and [`AppAll`](../src/AppAll.purs) select
them only for `--wip-included`. Default, extended and individual core-test runs
omit them. Drivers add the required dependencies to their isolated workspace;
sources and the oracle manifest are fingerprinted to reject stale artifacts.

To add a case, write its `prepare` function, register the same label and order
in `AppLibrary.cases` and `library-cases.json`, then document its consumer and
independent oracle here.

## Origins

These are original implementations inspired by operation families discussed on
PureScript Discord in March 2026, not a reproduction of that 218-case suite.
Its exact source was not found during the recorded search. The identified
[August 2024 string benchmark](https://github.com/anttih/purescm-string-benchmarks/blob/b792b00fa3fc8f50bccec5d28c0964221721ca89/src/Main.purs)
is a different suite. No upstream benchmark code was copied.
