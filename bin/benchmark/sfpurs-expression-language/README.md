# symfony/expression-language (`sfpurs-expression-language`) application benchmark

This benchmark measures compiling, linting, and dumping a corpus of 16 realistic
[Symfony ExpressionLanguage](https://github.com/symfony/expression-language) expressions
implemented in PureScript (`sfpurs-expression-language`), compiled with the official
PureScript JavaScript compiler versus `gopurs` (Go).

It is an application measurement, outside the numeric core and extended-test totals.
All 24 alternating samples, host details, and the verified output oracle are preserved
in [`results.json`](results.json).

## Current results

| Variant | Median runtime (24 runs) | Min / Max | Ratio `/JS` |
| --- | ---: | ---: | ---: |
| Official PureScript JS (`purs 0.15.16`, Node v24.8.0 V8 JIT) | **526.24 ms** | `521.48–571.52 ms` | — |
| `gopurs` Go (`purs 0.15.17` TAST, `go1.27.0`) | **165.54 ms** | `162.27–242.11 ms` | **0.31x** (`3.18×` faster) |

## Workload & verification

- **Corpus (`corpus.tsv`):** 16 realistic Symfony ExpressionLanguage expressions covering
  security voters, routing conditions, pricing rules, null-safe (`?.`), null-coalescing (`??`),
  regex `matches`, `starts with` / `ends with` / `contains`, hashes/arrays, `**`, `..`,
  multiline comments, and bitwise/logical operators. All 16 expressions were verified
  byte-for-byte identical to Symfony `8.2.0` (`Lexer`, `Parser`, `Compiler`, `Node::dump()`).
- **Per-process work:** `1000` iterations × `16` expressions × `3` operations
  (`compileWithConfig`, `lintWithConfig`, `dumpWithConfig`) = **`48,000` operations**
  producing **`4,408,000` UTF-8 output bytes**.
- **Oracle:** Every single output byte is hashed with 32-bit FNV-1a. Both JS and Go print
  the exact same checksum: `ops=48000 bytes=4408000 checksum=902571493`.
- **Protocol:** Whole-process wall-clock time (`time.perf_counter_ns`), 24 alternating
  runs per backend (`JS, Go` / `Go, JS`), 1 excluded warm-up per backend, no outlier
  exclusions. Inputs are read from disk at runtime (`corpus.tsv`) so they remain opaque.

## Rerun

```sh
python3 bin/benchmark/sfpurs-expression-language/measure.py --runs 24 --iters 1000
```
