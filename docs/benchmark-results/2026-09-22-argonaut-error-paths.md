# Deferred error wrappers in argonaut-codecs

The Argonaut decoders built their error wrappers before knowing whether the
field or element actually failed:

- `Data.Argonaut.Decode.Class.gDecodeJsonCons` built `AtKey fieldName` and then
  discarded it on success (`AtKey` partial-application allocations were 198 MB
  in the allocation profile, ~3 % of the decode total).
- `Data.Argonaut.Decode.Decoders.decodeArray` built `AtIndex i` plus a
  `lmap … <<< decoder` composition per element (`decodeArray` callback 270 MB,
  ~4 %).

`gopurs-argonaut-codecs` now constructs `AtKey`/`AtIndex` **only on the failure
path** in `gDecodeJsonCons`, `decodeArray`, `getField`, `getFieldOptional` and
`getFieldOptional'`. Error values, paths and priorities are unchanged; the full
diagnostic oracles (including the ten error cases) drive every measured process.

## Results

Paired campaigns (five alternating processes per version, oracles every
process, `GOMAXPROCS=1`, `GOGC=100`).

| JsonDecoding | before | after | delta |
|---|---:|---:|---:|
| decode | 26.398 ms | **23.184 ms** | **−12.2 %** |
| combined | 36.501 ms | **32.244 ms** | **−11.7 %** |
| decode allocations | 33.81 MiB | **29.01 MiB** | **−14.2 %** |
| combined allocations | 39.53 MiB | **34.73 MiB** | **−12.1 %** |
| JsonTypedAst, decode | 428.10 ms | 426.58 ms | −0.35 % (neutral) |

Cumulative official-format campaign in one session (previous published state
`json-native-abi/after` as control, current workspaces as treatment, three Go
and three JS processes each, interleaved):

| | before | after | delta |
|---|---:|---:|---:|
| JsonDecoding Go combined | 38.871 ms | **27.874 ms** | **−28.3 %** |
| JsonDecoding Go decode | 30.416 ms | **19.372 ms** | −36.3 % |
| JsonDecoding JS combined | 9.036 ms | **8.972 ms** | −0.7 % |
| JsonTypedAst Go combined | 558.942 ms | 522.236 ms | −6.6 % (drift) |
| JsonTypedAst JS combined | 78.052 ms | 86.371 ms | +10.7 % (drift) |

The JSON decoder cell is published as **27.87 ms**: it is consistent with the
previous published cell compounded by the two paired steps
(31.579 × (1 − 0.117) = 27.9) and with the same-session control. The TAST
campaign ran in a slower session — its own JS control moved +10.7 % and its
baseline Go 558.94 matched the published 560.75 — so the TAST cells keep the
previous official measurement (505.42 ms), the codecs step being neutral there
(−0.35 %). A focused six-run JS check confirms the change is neutral to
positive on JS (−3.7 % JSON, −4.4 % TAST), ruling out a JS regression.

## Cumulative decoder progress

| JsonDecoding | before this work | current |
|---|---:|---:|
| decode | 27.61 ms | **≈ 20 ms** (paired: −17.5 % then −12.2 %) |
| combined | 31.58 ms | **27.87 ms** (official, same-session control) |
| decode allocations | 41.65 MiB | **≈ 29 MiB** |

## Validation

- 17 JSON cases and 12 TAST modules matched the fixed oracles in every process,
  Go and JS, including the error-path cases that exercise `AtKey`/`AtIndex`.
- `bin/modtest argonaut-core` still 10/10; `node --test tools/*.test.mjs` was
  re-run green for the compiler-side changes of this session.
- Static: `AtKey` call sites 9 → 1, `AtIndex` 3 → 1,
  `Data_Argonaut_Decode_Decoders.go` `Func(` 58 → 51.

## Limits

- The fork `gopurs-argonaut-codecs` is currently wired only into the diagnostic
  workspaces (`gopurs-*` path overrides). Publishing it for `bin/test` and the
  sibling runners requires package-set/lockfile updates that were not done here.
- No b8x or whole-project campaign was run.

## Provenance

- Fork: `gopurs/gopurs-argonaut-codecs` (sources copied from
  `argonaut-codecs-9.1.0`, patched in `Decode/Class.purs` and
  `Decode/Decoders.purs`).
- Workspaces and campaigns: `scratch/abi-native-20260922/`
  (`official-paired-cumulative.log`, `abi-argonaut-codecs-*.log`,
  `js-paired.py`), archive
  [2026-09-22-argonaut-error-paths.json](2026-09-22-argonaut-error-paths.json).
