# Complete Purust compilation table — 4 October 2026

The **57 local Purust libraries** and the **b8x Rust test profile** are measured
under the JavaScript-hosted and Rust-hosted versions of Purust. Both generate Rust.

## Results

| Backend-only median | JavaScript host | Rust host | Rust / JS |
| --- | ---: | ---: | ---: |
| b8x Rust test profile | 44.32 s | 20.62 s | 0.47x |
| 57 libraries, sum of medians | 304.03 s | 120.12 s | 0.40x |
| Complete table, sum of medians | 348.34 s | 140.73 s | 0.40x |

## Protocol and scope

- Apple M4 Pro, 14 logical CPUs, 48 GiB RAM; Node v24.8.0; no explicit CPU affinity.
- One warmup, then five measured fresh processes per host. Projects and hosts
  run serially, alternating JS/Rust order. Public worker defaults: JavaScript
  1/1; native budget 8, split into four PBO and four code-generation workers.
- The internal **backend total** includes TAST loading/sorting, preparation,
  optimization, generation/emission and final drain. Frontend work, application
  and compiler builds, application execution and process startup/exit are excluded.
  Overlapping phase clocks are retained individually and never added together.
- Fresh generated output and PBO caches; warm operating-system filesystem cache.
  Each library retains its primary runner's threaded/non-threaded generation mode.
- The frontend is rebuilt from frozen native-family sources. Isolated Spago
  configurations resolve the Purust ports, including configurations inherited
  from Go projects. Root package names are normalized to their dependency names.
- Argonaut Core uses its existing compact-DOM fixture; Argonaut Codecs uses its
  existing typed JSON plans fixture. ArrayBuffer Types has only foreign type
  declarations and uses an explicit minimal executable harness.
- b8x is the existing **Core/Infra/Util Rust test profile**, not the larger
  Go-target application corpus. Its actual import closure and native FFI are frozen.
  Module/type counts and per-row scopes are recorded below and in the JSON.
- All measurements use the qualified compiler rebuild after replacing the
  backtracking foreign-type scan with a linear-time regex engine and preserving
  IEEE negative zero in PBO's native constant evaluator. Module roots are sorted
  canonically so spec registration and Cargo dependencies have identical order.
  The original failures, red/green regressions, 720 module permutations,
  self-hosting identity and fresh-project checks are retained with compiler hashes.
- The initial preparation's broken historical hello-world symlink and the
  Promise workspace's duplicate registry modules are recorded as harness issues.
  Spec's required BuildInfo module is produced by Spago before freezing its TAST.
  Earlier attempts and their statuses remain archived; the publication uses a
  complete new two-host campaign with the corrected configuration and compiler.
  A retry restored the original threaded-mode flags from the frozen runner plans;
  affected cases were re-measured, while correctly configured series were retained.

## Validation and evidence

- **696 exact-output-checked generations**, including
  **580 measured runs**, checking
  **323,712 generated files**.
- Every library's canonical application builds and its primary executable
  checks pass outside timing. Existing golden-output, marker and argument-scenario
  checks are retained; independently generated secondary mains are outside each
  primary compilation case. JavaScript and Rust produce byte-identical sources.
- The first application pass exposed a native HTTP panic: a trailing apostrophe
  in `HttpServer'` was dropped by foreign-declaration scanning, and a mangled
  identifier was used to look up the source-level FFI layout. Both are corrected.
  Historical-output checks also rejected a subsequent scan that omitted the
  Unicode kind separator `∷`. The final scan preserves both `::` and `∷`;
  declaration/layout parity is checked over the entire frozen source corpus.
  JS/Rust differential checks, Rc/Arc carrier checks, a fresh primed-handle smoke
  test, self-hosting identity and the real HTTP/upgrade/cookie/HTTPS and BigInt suites qualify
  the rebuilt hosts before a complete new 58-project measurement campaign.
  The previous measurements and failed HTTP executable/backtrace remain archived.
  Application evidence is reused only for byte-identical Rust/Cargo source,
  with explicit provenance to the retained executable and validation record.
- The b8x native import closure and Test.Rust.Main executable build outside
  timing; service-dependent execution is covered by the separate runtime profile.
- Frozen input and compiler manifests are checked before and after generation;
  every run retains phase samples, stdout/stderr, exit status and a file manifest.
  Identical output bytes are retained once per project, with separate failed outputs.
- Totals are **sums of per-project medians**, not a timed multi-project invocation.
- [Machine-readable report](2026-10-04-purust-packages.json)
- Local archive: `var/benchmark/purust-packages-20261004/revision3`
- Reproduction scripts: `bin/benchmark/purust-packages/`

## All medians (milliseconds)

| Project | Typed modules | Types | JavaScript | Rust | Rust / JS |
| --- | ---: | ---: | ---: | ---: | ---: |
| b8x | 1404 | 426499 | 44315 | 20619 | 0.47x |
| purust-aff | 243 | 138550 | 6218 | 2076 | 0.33x |
| purust-argonaut-codecs | 262 | 137872 | 6090 | 2120 | 0.35x |
| purust-argonaut-core | 252 | 129390 | 5122 | 1682 | 0.33x |
| purust-arraybuffer-types | 56 | 13062 | 591 | 158 | 0.27x |
| purust-arrays | 169 | 79880 | 15134 | 4291 | 0.28x |
| purust-assert | 58 | 13756 | 631 | 170 | 0.27x |
| purust-avar | 237 | 134884 | 5942 | 1938 | 0.33x |
| purust-catenable-lists | 152 | 81922 | 3474 | 1177 | 0.34x |
| purust-console | 59 | 13972 | 664 | 183 | 0.28x |
| purust-datetime | 197 | 105447 | 5659 | 1902 | 0.34x |
| purust-effect | 64 | 17656 | 1240 | 475 | 0.38x |
| purust-enums | 168 | 78160 | 2950 | 923 | 0.31x |
| purust-exceptions | 82 | 20217 | 924 | 276 | 0.30x |
| purust-exists | 60 | 13938 | 629 | 171 | 0.27x |
| purust-foldable-traversable | 162 | 78854 | 3593 | 1220 | 0.34x |
| purust-foreign | 227 | 118882 | 4949 | 1647 | 0.33x |
| purust-foreign-object | 252 | 136611 | 6320 | 2217 | 0.35x |
| purust-free | 198 | 115364 | 4570 | 1520 | 0.33x |
| purust-functions | 61 | 16316 | 1365 | 508 | 0.37x |
| purust-integers | 161 | 73198 | 4359 | 1398 | 0.32x |
| purust-js-bigints | 266 | 129675 | 6620 | 2336 | 0.35x |
| purust-js-date | 249 | 132606 | 5642 | 1933 | 0.34x |
| purust-js-promise | 258 | 139265 | 6029 | 2018 | 0.33x |
| purust-js-promise-aff | 272 | 143081 | 6165 | 2078 | 0.34x |
| purust-lazy | 133 | 62314 | 2140 | 693 | 0.32x |
| purust-node-buffer | 240 | 121502 | 4894 | 1601 | 0.33x |
| purust-node-child-process | 299 | 153890 | 7651 | 2609 | 0.34x |
| purust-node-event-emitter | 299 | 179888 | 8202 | 2809 | 0.34x |
| purust-node-fs | 283 | 147767 | 7430 | 2508 | 0.34x |
| purust-node-http | 310 | 156924 | 7105 | 2452 | 0.35x |
| purust-node-net | 296 | 152177 | 6805 | 2310 | 0.34x |
| purust-node-os | 274 | 138644 | 6032 | 2086 | 0.35x |
| purust-node-path | 59 | 13795 | 650 | 175 | 0.27x |
| purust-node-process | 284 | 148060 | 7196 | 2443 | 0.34x |
| purust-node-streams | 340 | 194528 | 9197 | 3489 | 0.38x |
| purust-now | 286 | 148033 | 6780 | 2310 | 0.34x |
| purust-nullable | 229 | 119097 | 4653 | 1557 | 0.33x |
| purust-numbers | 160 | 73162 | 4401 | 1408 | 0.32x |
| purust-ordered-collections | 240 | 133531 | 8217 | 2986 | 0.36x |
| purust-partial | 83 | 19863 | 871 | 268 | 0.31x |
| purust-prelude | 162 | 75454 | 2793 | 985 | 0.35x |
| purust-quickcheck | 232 | 123935 | 4753 | 1624 | 0.34x |
| purust-random | 161 | 73279 | 2660 | 898 | 0.34x |
| purust-record | 164 | 74962 | 2641 | 896 | 0.34x |
| purust-refs | 61 | 14381 | 718 | 196 | 0.27x |
| purust-run | 265 | 161007 | 6914 | 2262 | 0.33x |
| purust-spec | 391 | 236592 | 12717 | 4633 | 0.36x |
| purust-spec-discovery | 380 | 229754 | 11401 | 4318 | 0.38x |
| purust-spec-node | 382 | 232926 | 11931 | 4307 | 0.36x |
| purust-st | 160 | 73691 | 2554 | 866 | 0.34x |
| purust-strings | 192 | 81426 | 6589 | 1851 | 0.28x |
| purust-strings-extra | 186 | 79484 | 7666 | 11028 | 1.44x |
| purust-unfoldable | 162 | 74226 | 2651 | 920 | 0.35x |
| purust-unsafe-coerce | 135 | 60889 | 2089 | 669 | 0.32x |
| purust-uuid | 289 | 176135 | 7613 | 2588 | 0.34x |
| purust-variant | 164 | 91522 | 4537 | 1405 | 0.31x |
| purust-yoga-json | 386 | 233689 | 16694 | 14548 | 0.87x |
