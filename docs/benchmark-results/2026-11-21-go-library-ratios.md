# Go library ratios — November 21, 2026

The Go backend and native libraries now avoid several costs exposed by the
excluded WIP suite: dynamic array equality, repeated sorting work on monotone
inputs, the private reference used by `tailRecM`, and byte-by-byte ASCII scans.
The benchmarks and their immutable input fixtures were not changed.

All results below are medians of three independent process minima. The before
and after Go binaries and freshly built official JS/PBO programs were measured
sequentially in rotating order, after every build and diagnostic test finished.
All **444 outputs** passed the existing independent oracles: 14 core and 23 WIP
outputs in each of 12 processes. WIP results still have no aggregate score.

## Measured changes

Times are in microseconds. The README column records the historical Go baseline;
**before → after** measures the changes made in this campaign. JS ratios use the
same freshly measured JS denominator for both Go versions, rather than mixing
measurements from different dates. Lower ratios are better.

| Case | README Go | Go before | Go after | Speedup | Go / JS before → after | Go after / PBO |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| `string/cu-length-unicode-1000` | 2.66 | 2.902 | 2.563 | 1.13× | 275.57× → 243.37× | 296.16× |
| `string/indexOf-hit-10000` | 2.56 | 2.730 | 0.413 | 6.60× | 13.64× → 2.07× | 2.12× |
| `array/sort-descending-10000` | 4566.00 | 4695.146 | 607.038 | 7.73× | 8.21× → 1.06× | 1.16× |
| `array/eq-neq-last-10000` | 208.80 | 206.780 | 23.262 | 8.89× | 3.94× → 0.44× | 0.44× |
| `st/ref-write-10000` | 536.37 | 557.404 | 289.660 | 1.92× | 1.15× → 0.60× | 2.72× |
| `st/ref-modify-10000` | 538.57 | 561.691 | 294.722 | 1.91× | 1.06× → 0.56× | 2.77× |
| `effect/ref-write-10000` | 925.27 | 1017.349 | 399.021 | 2.55× | 6.38× → 2.50× | 3.71× |
| `effect/ref-modify-10000` | 1031.72 | 1134.076 | 543.413 | 2.09× | 5.37× → 2.57× | 4.54× |

The largest relative gap remains UTF-16 code-unit length. It improved by **11.7%**
in the complete PureScript benchmark, but still takes **243×** the fresh JS
measurement. Native strings retain UTF-8/WTF-8, and this implementation still
scans the input to count UTF-16 units. This change does not make length constant
time. The historical README reported 266×; the matched baseline in this campaign
is 276×, which is the appropriate before value for the measured improvement.

ST references now beat official JS in these cases but remain approximately
2.7× slower than PBO. Effect reference loops remain approximately 2.5× official
JS and 3.7–4.5× PBO. Their callbacks and `Step` values still cross the boxed ABI;
their remaining allocations were not eliminated by this change.

Other WIP rows remain within 3% of the matched baseline. The 14-case core total
is **13.454 → 13.298 ms**, compared with the historical README's **13.28 ms**.
These small differences do not establish a separate core performance gain.
No C measurements were added for these WIP cases.

## Implementation and semantic checks

- `gopurs-strings/src/Data/String/CodeUnits.go`: scan ASCII runs in 32/8-byte
  blocks and classify non-ASCII sequence widths with a table. Preserve UTF-16
  counting, isolated surrogates and malformed-byte behavior. JS/Go parity,
  all 65,536 two-byte inputs, 10,000 deterministic random strings and transitions
  around word boundaries pass. No allocation was added by the native helper.
- `gopurs-arrays/src/Data/Array.go`: detect monotone runs before the generic
  stable sort. Only strictly decreasing runs are reversed, preserving the order
  of equal keys; the input remains unchanged. Fifty cases, including large
  random and duplicate-key inputs, match the JS sort.
- `Gopurs.ArrayEquality`: recognize canonical scalar `Eq` dictionaries and emit
  a direct element loop. Custom instances retain their callbacks. Tests cover
  empty/different-length arrays, unequal final elements, NaN, signed zero,
  Unicode/WTF-8, booleans, and argument evaluation order. The actual altbak Go
  output contains the specialized loop.
- `Gopurs.TailRecEffect`: replace the private `Step` reference of the standard
  Effect/ST `tailRecM` instances with loop-local state. User references and their
  mutexes retain their normal behavior. Tests verify deferred effects, capture
  order, replay of one action, immediate completion, distinct state/result
  types, exception propagation and rejection of custom instances. The four
  unchanged altbak reference cases pass in release and development modes.

The JS compiler builds without errors or warnings, and the native compiler was
rebuilt successfully with the local TAST fork. Targeted tests, six adjacent array
checks and seven recursive-function/callback checks passed. Independent code
reviews found no semantic defect in the string scan or Effect/ST transformation.
`git diff --check` passes in all three implementation repositories.

## Evidence and reproduction

The [JSON archive](2026-11-21-go-library-ratios.json) contains all process results,
source revisions and patches, historical README timings, build manifests,
source/payload fingerprints and raw-log hashes. Complete local logs and binaries
are in `var/benchmark/go-ratios-20261121/`, including `strings-diagnostic/` and
`refs-diagnostic/`. The main README retains its historical multi-column table;
this report records the new Go candidate and the fresh JS references separately.

Host: Apple M4 Pro, macOS 26.1, Go 1.27.0, Node 24.8.0. Application runs use
`GOGC=800`, PGO off and no explicit CPU affinity. Go before/after use the same
monolithic output mode. The initial default incremental build failed because
`regexp2` was missing from its readonly module plan; the successful builds used
`GOPURS_GO_PACKAGES=0 GOPURS_GO_PACKAGE_OVERLAY=0`, followed by the existing
driver's `go mod tidy`. That failed build contributed no measurement.

The source and native-compiler changes span `gopurs`, `gopurs-arrays` and
`gopurs-strings`; their precise base revisions and working-tree patches are in
the archive. No benchmark-specific special case, result cache or input mutation
was introduced.

```sh
# Run from altbak.pub after rebuilding gopurs with npm run build:native.
GOPURS_GO_PACKAGES=0 GOPURS_GO_PACKAGE_OVERLAY=0 \
  python3 bin/native/driver.py go --wip-included --build-only \
  --build-dir var/benchmark/go-ratios-20261121/after

# The archived plan verifies the frozen binaries and JS payloads.
GOGC=800 python3 bin/benchmark/measure.py \
  --plan var/benchmark/go-ratios-20261121/plan.json \
  --output path/to/new-campaign
```

The frozen plan is specific to this machine. Rebuild and freeze a new plan when
sources, compilers, paths or binaries change. Do not rebuild an artifact in place
and reuse an old plan with mismatched hashes.
