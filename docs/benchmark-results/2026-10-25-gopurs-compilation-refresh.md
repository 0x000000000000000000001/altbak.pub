# Gopurs compilation refresh — 25 October 2026

The 51 existing project rows have new JS Release, Go Release and Go Dev measurements. **The Rust Release column is estimated from historical per-project Rust/Go ratios.**

## Measurement protocol

- JS Release, Go Release and Go Dev: one warmup and three measured fresh processes per cell; each measured cell is a median. All Go/Dev invocations were rerun in v3, alternating their order across rounds; corrected JS invocations were retained from v2 and therefore precede these native measurements in time.
- Fixed TAST and FFI inputs; fresh generated output and backend caches before every invocation. Operating-system caches remain warm.
- b8x retains the pinned Go dependency manifests used by its normal runner. Dependencies are available offline; the shared Go export/compiler cache is warmed by the warmup.
- The internal `backend total` timer excludes frontend work, compiler bootstrap, final Go builds, application execution and process startup/exit.
- Measured variants enable package generation, package overlays and build manifests, with 8 preparation/PBO/emission workers and pipelining. Go/Dev explicitly set `GOGC=off` and `GOMEMLIMIT=16GiB`; `GOMAXPROCS` remains unset. JS retains its v2 launcher environment without those explicit Go GC overrides.
- Dev means `gopurs --dev` on the Go host. Final Go compiler flags `-N -l` affect execution builds, not this table.

## Comparability with the previous table

All 153 historical timing cells match the [4–5 October saturation campaign](2026-10-05-gopurs-rust-saturation.md): b8x Go was 33,037 ms and Assert Go was 150 ms. Both campaigns report the internal `backend total` and exclude the frontend, compiler bootstrap, final application compilation, execution and process startup/exit. The older campaign used one warmup and five measured processes per cell; this campaign uses three measured processes.

The work inside that timer has changed. The historical driver generated every module and finished after writing entry points; it did not perform the current Go package/overlay processing or persistent incremental-cache publication. The current protocol explicitly enables `GOPURS_GO_PACKAGES=1`, `GOPURS_GO_PACKAGE_OVERLAY=1` and `GOPURS_BUILD_MANIFEST=1`. Its total includes cache preparation, Go dependency export preparation, package generation, cache publication and source-proof publication. Final Go application compilation remains outside the timer.

Other inputs also differ: the historical b8x corpus contained 2,685 typed modules, versus 2,707 now; Assert has 58 in both campaigns, but its frozen TAST bytes are not identical. Compiler revisions changed, and the Go host now uses an explicit 16 GiB memory limit with `GOGC=off`, compared with the historical explicit 10 GiB limit. Both use eight workers per stage, pipelining, and warm filesystem caches.

**These historical differences do not isolate a regression.** The current total covers additional work, but this campaign does not measure how much of the gap each change causes, nor establish that the whole increase is necessary for incremental builds. That would require controlled comparisons that hold the input corpus and build protocol fixed while changing a compiler version or feature.

Sources and fingerprints:

- [Historical results JSON](2026-10-05-gopurs-rust-saturation.json), SHA-256 `3e75217d0341b78f7970cf39bd268191111f5a0137d6b8e94729fa343966bbc6`. Its recorded launcher SHA-256 `a2f198ace9ee575e3528b4f2f64f7a149cd0f9d2c9e53c0ffc3b83103d6e50dd` matches `gopurs@437cd8d:bin/gopurs` exactly.
- Historical driver: `gopurs@437cd8d:src/Gopurs/Driver.purs`, SHA-256 `94b88acea81dd09637e9c00f963b9e33414324b3e2e5d94f714c3bafe87aacbb`; module regeneration is explicit in `src/Gopurs/Driver/Build.purs` at the same revision (`onSkipModule` returns `Nothing`).
- Original measurement harness: `measure-current.mjs` in the local archive, SHA-256 `7a4fc0db0271d9416f51b27498965ec30c8b9b945327d91dd7812f54a76b1ab5`. Resume harness: `measure-current-v2.mjs`, SHA-256 `3c00cfdfd38ec275649f236a6685f1d675a7e335883a5dfa3f03ec0bd277df73`. V3 native rerun harness: `measure-current-v3.mjs`, SHA-256 `7b8c5699f9115c7210a9995425fb23bd872cd5b6ee6f4249cd904bb3ae70ccf3`. Environment source: `campaign-environment-v3.mjs`, SHA-256 `9d142c82312c1145398c91fbc4d8bef87e01e3e1b3abe276bbbd6ab1bbab333e`. Compiler hashes are recorded separately for each variant.

## JS transport correction

Two JS invocations stalled in the Go package helper's `io.ReadAll(stdin)`, with the Node parent waiting in a synchronous subprocess call. The JS FFI now passes helper input through a temporary regular file descriptor instead of a pipe. The corrected JS bundle is frozen separately; the original Go executable and its snapshot remain unchanged. The source comparison permits only the FFI transport implementation, its bundled JS and its targeted test to differ.

All original pipe-transport JS invocations, including successful ones, are archived and excluded. Every published JS sample uses the corrected bundle. Those 204 corrected JS invocations are preserved unchanged from v2. Release oracles are rehashed from their saved generated Go; all new Go-host outputs must match the retained JS outputs and the saved oracle byte-for-byte. Dev outputs independently match their saved oracle. No failed or diagnostic duration enters any median or total.

## Explicit native GC policy and replacement of v2 native measurements

The previous harness cleared `GOGC` and `GOMEMLIMIT` and relied on the launcher to infer memory capacity. A launcher-only check showed that the sandbox can deny `sysctl hw.memsize`, leaving that automatic GC policy unset. Consequently all 408 prior Go/Dev invocations, including their warmups, are archived and excluded. V3 repeats one warmup and three measurements for both native variants with explicit `GOGC=off` and `GOMEMLIMIT=16GiB`. The frozen compiler executable, TAST inputs, generated-code oracles, workers and backend feature flags are unchanged. A separate check executes only the frozen launcher environment prefix, without starting a compiler, and confirms both explicit values survive launch.

The corrected JS measurements were recorded before these native reruns. This is not one fully interleaved JS/Go/Dev campaign. The excluded native timings do not establish a PBO defect or a controlled GC speedup: their effective policy and temporal conditions were not held fixed. The current table uses only the explicitly configured native measurements.

- Launcher-only proof SHA-256: `480c94ec90de747cb03236a1c91b15f687f3165eb8bd815efdba1d95744999e9`.
- Excluded native measurements and previous publication: `/Users/0x1/Documents/htdocs/.gopurs-evidence/compilation-readme-20261025/excluded-attempts/native-gc-implicit-v2`.

- Measured JS bundle SHA-256: `116a991f7bff46d5cecc38a268eebb31f39e6763b2738d17334489e7a760c938`.
- Measured Go/Dev executable SHA-256: `42bc5a5dd8363b397860973a9ac4647aed84e40ba812a6266c617a8b0f30d482`.
- FFI source diff SHA-256: `df47d6a7b2a0cf1e9e32d875ccfba4f41414e6d9b486d7db314cba53c33035d6`; source paths and freeze provenance are included in the JSON report.

## Rust estimation

`estimated Rust ms = new Go Release median ms × historical Rust ms / historical Go ms`

- The ratio is captured independently for every project from the previous README table, before replacement.
- Rust totals are sums of the unrounded row estimates. They are not calculated from the previous aggregate ratio.
- These estimates assume historical relative speed remains applicable. Compiler changes and cold-cache work may affect the hosts differently.
- Rust values are not new measurements and have no new Rust execution, output-equality or PGO validation behind them.
- JS, Go and Dev totals are sums of their measured per-project medians.

## Evidence and limits

- 51 corpora; 612 successful generations, comprising 459 measured samples and 153 warmups. Release Go outputs match between JS and Go; Dev outputs are checked independently across repetitions.
- Input files are checked against frozen hashes. No application test suite was rerun for this timing campaign.
- Failed attempts and diagnostic runs are excluded from all medians and totals. The first b8x warmup lacked its external Go dependency manifests; the rerun used the pinned manifests described above.
- A Datetime JS warmup and a Free JS measured invocation stalled in helper input. Both were interrupted to capture Go stacks and archived. Separate exact-request Datetime replays succeeded eight times, so isolated request replay did not reproduce the intermittent stall. The subsequent transport correction and full JS replacement are described above.
- [Measured samples, compiler hashes and historical estimation inputs](2026-10-25-gopurs-compilation-refresh.json). Full local logs, inputs and manifests: `/Users/0x1/Documents/htdocs/.gopurs-evidence/compilation-readme-20261025`.

## Backend measurements and estimates

| Project | JS Release (measured) | Go Release (measured) | Rust Release (estimated) | Go Dev (measured) |
| --- | ---: | ---: | ---: | ---: |
| b8x | 146389 ms | 78037 ms | 52921 ms (est.) | 53611 ms |
| gopurs-aff | 16023 ms | 6312 ms | 5002 ms (est.) | 5437 ms |
| gopurs-argonaut-core | 15672 ms | 5930 ms | 4800 ms (est.) | 4962 ms |
| gopurs-arrays | 12758 ms | 5072 ms | 4001 ms (est.) | 3782 ms |
| gopurs-assert | 2514 ms | 1305 ms | 922 ms (est.) | 1187 ms |
| gopurs-avar | 15097 ms | 5865 ms | 4770 ms (est.) | 4971 ms |
| gopurs-catenable-lists | 10188 ms | 3918 ms | 3024 ms (est.) | 3387 ms |
| gopurs-console | 9476 ms | 3207 ms | 2513 ms (est.) | 2880 ms |
| gopurs-datetime | 14497 ms | 5560 ms | 4324 ms (est.) | 4768 ms |
| gopurs-effect | 4161 ms | 1390 ms | 994 ms (est.) | 1296 ms |
| gopurs-enums | 10858 ms | 3936 ms | 3125 ms (est.) | 3300 ms |
| gopurs-exceptions | 4671 ms | 1665 ms | 1195 ms (est.) | 1519 ms |
| gopurs-foldable-traversable | 10499 ms | 3841 ms | 3057 ms (est.) | 3357 ms |
| gopurs-foreign | 12846 ms | 5309 ms | 4347 ms (est.) | 4621 ms |
| gopurs-foreign-object | 15874 ms | 6228 ms | 5136 ms (est.) | 5090 ms |
| gopurs-free | 11960 ms | 4840 ms | 3881 ms (est.) | 4252 ms |
| gopurs-functions | 4231 ms | 1417 ms | 960 ms (est.) | 1299 ms |
| gopurs-integers | 10024 ms | 3470 ms | 2707 ms (est.) | 3098 ms |
| gopurs-js-bigints | 4625 ms | 1596 ms | 1179 ms (est.) | 1443 ms |
| gopurs-js-date | 16044 ms | 6227 ms | 5007 ms (est.) | 5297 ms |
| gopurs-js-promise | 8189 ms | 2787 ms | 2196 ms (est.) | 2531 ms |
| gopurs-js-promise-aff | 16617 ms | 6434 ms | 5158 ms (est.) | 5511 ms |
| gopurs-js-uri | 9730 ms | 3275 ms | 2569 ms (est.) | 2948 ms |
| gopurs-lazy | 7605 ms | 2758 ms | 2178 ms (est.) | 2526 ms |
| gopurs-node-buffer | 14716 ms | 5351 ms | 4366 ms (est.) | 4743 ms |
| gopurs-node-event-emitter | 21845 ms | 8936 ms | 7607 ms (est.) | 6631 ms |
| gopurs-node-fs | 18796 ms | 6999 ms | 5685 ms (est.) | 5912 ms |
| gopurs-node-http | 20650 ms | 7475 ms | 6173 ms (est.) | 6248 ms |
| gopurs-node-net | 19243 ms | 6945 ms | 5630 ms (est.) | 5865 ms |
| gopurs-node-path | 4257 ms | 1302 ms | 914 ms (est.) | 1204 ms |
| gopurs-node-process | 17942 ms | 6684 ms | 5391 ms (est.) | 5685 ms |
| gopurs-node-streams | 23385 ms | 9204 ms | 7938 ms (est.) | 6884 ms |
| gopurs-now | 18036 ms | 6756 ms | 5480 ms (est.) | 5772 ms |
| gopurs-nullable | 14155 ms | 5343 ms | 4399 ms (est.) | 4659 ms |
| gopurs-numbers | 10142 ms | 3507 ms | 2749 ms (est.) | 3170 ms |
| gopurs-ordered-collections | 16381 ms | 6449 ms | 5047 ms (est.) | 5138 ms |
| gopurs-partial | 9519 ms | 3188 ms | 2493 ms (est.) | 2859 ms |
| gopurs-prelude | 10711 ms | 3557 ms | 2792 ms (est.) | 3053 ms |
| gopurs-random | 9902 ms | 3224 ms | 2525 ms (est.) | 2880 ms |
| gopurs-record | 10133 ms | 3356 ms | 2627 ms (est.) | 3019 ms |
| gopurs-refs | 9881 ms | 3300 ms | 2594 ms (est.) | 2960 ms |
| gopurs-run | 17665 ms | 6775 ms | 5433 ms (est.) | 5776 ms |
| gopurs-spec | 30416 ms | 12821 ms | 10954 ms (est.) | 8658 ms |
| gopurs-st | 9659 ms | 3302 ms | 2597 ms (est.) | 2980 ms |
| gopurs-strings | 12680 ms | 4489 ms | 3805 ms (est.) | 3824 ms |
| gopurs-strings-extra | 13757 ms | 6990 ms | 4827 ms (est.) | 6343 ms |
| gopurs-unfoldable | 9565 ms | 3322 ms | 2607 ms (est.) | 2979 ms |
| gopurs-unsafe-coerce | 9841 ms | 3290 ms | 2562 ms (est.) | 2941 ms |
| gopurs-uuid | 20972 ms | 8748 ms | 7340 ms (est.) | 6505 ms |
| gopurs-variant | 11455 ms | 4451 ms | 3279 ms (est.) | 3881 ms |
| gopurs-yoga-json | 36226 ms | 16572 ms | 13134 ms (est.) | 11991 ms |
| **Total gopurs-*** | ~ 666.09 s | ~ 254.68 s | ~ 203.99 s (est.) | ~ 212.02 s |
| **Total** | ~ 812.48 s | ~ 332.71 s | ~ 256.92 s (est.) | ~ 265.63 s |
