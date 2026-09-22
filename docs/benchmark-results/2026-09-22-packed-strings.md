# Packed string payloads in the gopurs runtime

The `TypeString` value no longer allocates a `string` header per box. `IntVal`
holds the byte length and `UnsafePtr` the immutable string data
(`unsafe.StringData`); readers go through a single `StrValue` helper. The Go
runtime, the Argonaut FFI, `Foreign`, `Foreign.Index` and `Node.Buffer` were
adapted.

Paired campaigns on the two JSON diagnostics show a consistent decode gain,
with every process matching the fixed oracles:

| Diagnostic (paired, 5 processes per version) | Go before | Go packed | delta |
|---|---:|---:|---:|
| JSON Decoding: decode | 30.917 ms | **27.581 ms** | **−10.79 %** |
| JSON Decoding: combined | 39.505 ms | **37.683 ms** | **−4.61 %** |
| JSON → Typed AST: decode | 450.455 ms | **413.480 ms** | **−8.21 %** |
| JSON → Typed AST: combined | 548.161 ms | **524.276 ms** | **−4.36 %** |

Allocated bytes per corpus fall with the removed string headers:

| Diagnostic | phase | Go before | Go packed |
|---|---|---:|---:|
| JSON Decoding | decode | 44.33 MiB | **41.65 MiB** |
| JSON Decoding | combined | 50.06 MiB | **47.37 MiB** |
| JSON → Typed AST | decode | 488.14 MiB | **465.01 MiB** |
| JSON → Typed AST | combined | 553.41 MiB | **530.28 MiB** |

A first 5-process campaign reproduces the direction with different absolute
values: JSON decode −8.18 %, combined −5.66 %; TAST decode −7.28 %, combined
−6.41 %. A 3-process campaign shows JSON decode −7.77 %, combined −6.63 % and
TAST decode −8.67 %, combined −4.03 %. The parse phase is unchanged within
noise (it creates strings through `encoding/json`, not through `Str`).

## What changed

- `gopurs/runtime/runtime.go`: packed `Str`, new `StrValue`, and 13 readers
  (`StrVal`, `AnyVal`, `RecordGet`, `Unbox`, `ValueToAny`, `ExtractVariant`,
  `copyReflectField`).
- `gopurs/src/Gopurs/Runtime.(go|js)`: regenerated with
  `tools/embed-runtime.mjs`.
- `gopurs-argonaut-core`, `gopurs-foreign`, `gopurs-node-buffer`: 13 FFI
  `*(*string)` reads replaced by `gopurs_runtime.StrValue`.
- `gopurs/tools/ffi-generics.test.mjs`: the ordering guard now expects zero
  allocations per call instead of exactly two string boxes; the stronger
  property is what the packed representation provides.

The representation contract:

- the referenced bytes must remain immutable for the whole lifetime of the
  `Value` (`unsafe.String` requirement); every producer of strings in the Go
  runtime already returns immutable Go strings, and no FFI builds a string over
  a mutable buffer without copying;
- a substring keeps its whole backing alive; this is a retention trade-off,
  not a correctness issue;
- empty strings are represented by a nil pointer and read as `""`;
- the `Value` struct size is unchanged.

## Official-format controls

The same packed workspaces were measured once with the diagnostic protocol
(three processes per backend, interleaved Go/JS), which also re-measures JS:

| Diagnostic | Go parse | Go decode | Go combined | JS combined |
|---|---:|---:|---:|---:|
| JSON Decoding | 7.627 ms | 28.584 ms | 37.186 ms | 9.031 ms |
| JSON → Typed AST | 76.262 ms | 414.689 ms | 528.474 ms | 79.391 ms |

The JS bundles are byte-identical between the before/after workspaces (same
`js_sha256` per suite), so JS differences are session drift only. Across
sessions the JS control moved from 9.64/77.02 ms to 9.03/79.39 ms while the Go
paired deltas stayed negative, which is why the paired campaigns, not
cross-session absolutes, carry the attribution. The README cells use the
archived paired campaign (`Go packed` column above).

## Validation

- Fixed oracles: 17 general JSON cases and 12 TAST modules matched in every
  measured process; invalid-input oracles are checked outside timing.
- `node --test tools/*.test.mjs`: 170 tests, 168 pass, 0 fail, 2 skipped.
- `bin/modtest`: `argonaut-core`, `foreign`, `foreign-object`, `node-buffer`,
  `prelude` and `strings-extra` pass with the packed runtime.
- `gopurs-strings` fails on a literal with isolated surrogates. The same
  failure reproduces with the baseline bundle at the same assertion; it is the
  pre-existing `StringEdgeCases`/`StringEscapes` limitation, not a regression.
- `bin/go/run -c` from altbak: 14 reference outputs, exit 0, total 13.138 ms.
- A standalone GC stress probe of 4 M packed values (including substrings)
  survives repeated collections; the packed workspaces pass the CGo-free
  `-race`/checkptr oracle test.

## Limits

- No complete b8x build or default parallel-load measurement was re-run.
- Only the two JSON diagnostics were re-measured; the core 14 cases were
  checked for correctness and total time, not for paired deltas.
- Retention of large backing strings through small substrings is documented
  but not quantified here.
- The baseline workspaces cannot be re-measured through the runner once the
  sources are packed, because the diagnostic runner rejects stale manifests;
  the paired driver executes both frozen benchmark binaries directly.

## Reproduction and provenance

- Paired driver and raw reports:
  `scratch/packed-strings-20260922/` (`ab-campaign.py`,
  `ab-strings-results.json`, `ab-strings-raw.json`).
- Official-format campaigns: `final-json/`, `final-tast/` (results, logs,
  manifests).
- Frozen compilers: `gopurs-native-baseline`
  (`bfd37bad…`) and `gopurs-native-packed` (`ae42e79a…`) under
  `scratch/tast-profile-20260922/compilers/`.
- Raw archive for this report:
  [2026-09-22-packed-strings.json](2026-09-22-packed-strings.json).
