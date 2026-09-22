# General JSON decoding — September 22, 2026

The new `JsonDecoding` diagnostic exercises ordinary PureScript Argonaut code
without a PBO dependency. Its first baseline is **57.432 ms Go / 8.258 ms JS**
for parsing and decoding the successful corpus, a **6.95×** ratio. The slowdown
therefore also occurs on this application-shaped workload; it is not confined
to the compiler's TAST decoder. This does not establish a ratio for every JSON
workload or identify the cost of each remaining compiler/runtime mechanism.

## Workload and correctness

The shared [PureScript source](../../src/Test/JsonDecoding.purs) uses ordinary
`DecodeJson` record instances and a small hand-written instance for tagged events.
Payloads include nested records, arrays, optional records/strings/integers,
booleans, numeric values, Unicode and escaped strings. The JS/Go FFI only handles
input, timing, allocation counters, result consumption and hashing.

There are **17 fixed cases**: five successful cases are timed (660,909 UTF-8
input bytes), ten cases verify precise decoding errors, and two additional
cases verify missing/null optional fields. Errors include fractional integers,
wrong types, missing fields, invalid tags, nested array paths and first-error
priority. All inputs are syntactically valid JSON; parser rejection cases are
outside this diagnostic's scope.

The independent [fixture oracle](../../test/fixtures/json-decoding/README.md)
checks every decoded field and complete error message, rather than a checksum
of selected fields. Expected results come from explicit constructors and manually
specified error paths, never from either backend's observed output. All 17 cases
match in all six measured processes. Invalid inputs are checked outside timing.

## Initial baseline

| Successful corpus | Go | JS | Go / JS | Go allocated MiB |
|---|---:|---:|---:|---:|
| Parse | 6.654 ms | 1.517 ms | 4.39× | 5.720 |
| Decode parsed JSON | 48.947 ms | 6.673 ms | 7.33× | 77.368 |
| Parse + decode | **57.432 ms** | **8.258 ms** | **6.95×** | **83.104** |

Three independent processes per runtime, ordered Go/JS/JS/Go/Go/JS, with two
warm-ups and five samples per phase. Cells are medians of process minima;
allocation cells are medians of the fifteen Go samples. Combined timing is
measured independently. File I/O and fingerprinting are outside timing, although
validation allocates between samples. Go uses GOMAXPROCS=1, GOGC=100 and PGO off;
Node retains its default background threads. No builds or agent tests ran during
this campaign. The [archive](2026-09-22-json-decoding.json) retains every sample,
fixture/source hashes, compiler/runtime versions and executable hashes.

This is a new baseline, not a before/after optimization. The existing README
TAST diagnostic (573.63 ms Go / 75.90 ms JS) uses another corpus and decoder;
its absolute times must not be compared as a speedup. The new README cells are
excluded from the historical fourteen-case totals and impose no expectation
that Go must outperform JS.

## Running the diagnostic

Use `./bin/go/run --test JsonDecoding` or `./bin/js/run --test JsonDecoding`.
Both support `--build-only`, `--run-only`, `--build-dir` and `--output`; builds
are isolated and stale sources/binaries are rejected. `./bin/run --test JsonDecoding`
selects Go and JS only. The diagnostic is excluded from the ordinary `--x` suite.

The initial paired campaign used one build containing both executables:

```sh
./bin/go/run --test JsonDecoding --build-only
python3 -B bin/benchmark/json-diagnostic.py measure --suite JsonDecoding \
  --workspace run/bak/go/modes/test-JsonDecoding/builds/20260921T231837.961195Z-31990 \
  --output var/benchmark/json-decoding/2026-09-22-initial
```

Both language builds completed without warnings. The 11 diagnostic runner tests,
four fixture/oracle checks and three publication tests pass.
