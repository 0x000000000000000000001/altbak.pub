# Dedicated `caseJson*` FFI in Argonaut

The six single-type case helpers in `Data.Argonaut.Core`
(`caseJsonNull`, `caseJsonBoolean`, `caseJsonNumber`, `caseJsonString`,
`caseJsonArray`, `caseJsonObject`) were defined through the generic `caseJson`
by passing five or six `const d` callbacks. On Go, every call built six
closures before `CaseJsonImpl` applied one of them.

Each helper now has a dedicated FFI implementation that selects the matching
branch directly, with exactly the dispatch semantics of `CaseJsonImpl`
(null, boolean, number, string, array, then object for records and
constructors). The public PureScript API is unchanged.

## Published cells

Official diagnostic campaigns (three processes per backend, interleaved
Go/JS, median of process minima, oracles checked every process):

| Diagnostic | phase | before | after | delta |
|---|---|---:|---:|---:|
| JSON Decoding | decode | 27.61 ms | **22.91 ms** | **−17.0 %** |
| JSON Decoding | combined | 37.19 ms | **31.58 ms** | **−15.1 %** |
| JSON Decoding | decode allocations | 41.65 MiB | **33.81 MiB** | −18.8 % |
| JSON Decoding | combined allocations | 47.37 MiB | **39.53 MiB** | −16.6 % |
| JSON → Typed AST | decode | 460.79 ms | **414.73 ms** | **−10.0 %** |
| JSON → Typed AST | combined | 560.75 ms | **505.42 ms** | **−9.9 %** |
| JSON → Typed AST | decode allocations | 488.14 MiB | **454.56 MiB** | −6.9 % |

JS controls, same campaigns: JSON 9.03 ms combined (was 9.64 in the previous
cell) and TAST 80.94 ms (was 77.02). The JS bundles changed with the library,
so both backends were validated against the oracles; the Go deltas above are
also confirmed by alternating paired campaigns (−17.5 % decode on JSON,
−8.6 % decode on TAST, five processes per version).

Static effect in the generated decoder: `Data_Argonaut_Core.go` drops from
114 to 24 `gopurs_runtime.Func(` allocations (−79 %) and from 119 662 to
81 696 bytes (−32 %).

## Validation

- 17 general JSON cases and 12 TAST modules matched the fixed oracles in every
  measured process, on both backends.
- `bin/modtest argonaut-core`: 10/10 tests pass with the new FFI.
- `node --test tools/*.test.mjs`: 170 tests, 168 pass, 0 fail, 2 skipped.
- The dedicated FFI keeps the exact `CaseJsonImpl` fallthrough, including
  `TypeAny` null handling and the object branch for constructors.

## Limits

- The remaining decode cost is dominated by boxed dictionary dispatch
  (`Apply`/`Apply2` account for ~90 % of allocation cumulatively); this change
  removes one large adapter but does not devirtualise the decoder.
- No b8x campaign was re-run; the library change is shared by every Argonaut
  user, so whole-project effects were not measured.

## Provenance

[Raw archive](2026-09-22-casejson-ffi.json) with campaign medians, paired
results, source hashes and workspace hashes. Local workspaces:
`scratch/abi-native-20260922/` (`new-json`, `new-json2`, `final-json`,
`new-tast`, `new-tast2`, `published-json`, `published-tast`).
