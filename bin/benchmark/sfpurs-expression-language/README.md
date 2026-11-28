# sfpurs/expression-language rewrite benchmark

This benchmark evaluates rewriting [Symfony ExpressionLanguage](https://github.com/symfony/expression-language) (`v8.2.0`)
in PureScript (`sfpurs/expression-language`) and compiling it to Go via `gopurs`.

## 1. Rewrite benchmark: Symfony 8.2 PHP vs. FrankenPHP Go extension (`gopurs`)

Measured inside the `sfpurs-frankenphp` container (PHP 8.6.0 ZTS, Apple M4 Pro arm64) on 4 realistic
Symfony workloads (`security`, `routing`, `pricing`, `complex`) × 2 operations (`compile`, `lint`) =
**8 operations**, `2000` iterations each (after warmup and exact output equivalence verification),
including the full PHP ↔ CGO ↔ Go boundary:

| Profile (total 8 ops) | Original Symfony 8.2 PHP | `sfpurs` FrankenPHP Go ext (`gopurs`) | Ratio `/Orig` |
| --- | ---: | ---: | ---: |
| JIT OFF (default prod) | **237.71 μs** | **68.97 μs** | **0.29x** (`3.45×` faster) |
| OPcache + JIT 1255 | **164.85 μs** | **63.57 μs** | **0.39x** (`2.59×` faster) |

In standalone native Go (without the PHP/CGO boundary), the same 8 operations complete in **`31.39 μs`**.

## 2. Standalone PureScript runtime check: Official JS (V8) vs. `gopurs` Go

To isolate the `gopurs` backend against official PureScript JS (`purs 0.15.16`, Node v24.8.0 V8 JIT),
the 16-expression corpus (`corpus.tsv`, all verified byte-for-byte identical to Symfony 8.2) is run
for `1000` iterations × 3 ops (`compile`, `lint`, `dump`) = **`48,000` operations** / **`4,408,000` UTF-8 bytes**
with full FNV-1a checksum verification (`902571493`), median of 24 alternating fresh-process runs
(archived in [`results.json`](results.json)):

| Variant | Median runtime (24 runs) | Min / Max | Ratio `/JS` |
| --- | ---: | ---: | ---: |
| Official PureScript JS (`purs 0.15.16`, Node v24.8.0 V8 JIT) | **526.24 ms** | `521.48–571.52 ms` | — |
| `gopurs` Go (`purs 0.15.17` TAST, `go1.27.0`) | **165.54 ms** | `162.27–242.11 ms` | **0.31x** (`3.18×` faster) |

## Rerun

```sh
# Standalone Go vs JS corpus benchmark (24 alternating runs)
python3 bin/benchmark/sfpurs-expression-language/measure.py --runs 24 --iters 1000

# FrankenPHP Symfony PHP vs Go extension benchmark
cd ../sfpurs/sfpurs-expression-language && ./bin/bench 2000
```
