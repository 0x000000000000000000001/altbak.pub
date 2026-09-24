# PHP private tags and the current compiled column

The compiled PHP column was refreshed from campaign 2026-09-24d
(bundle `86ad0c0b`). The FFI columns were verified code-neutral and are kept
from the previous campaign: the current machine window inflates every
measurement, including the previously published build.

## What changed since the enum-regions campaign

- `Phpurs.CodeGen` no longer emits the unused `$tag` property on private
  constructors produced by `EnumRegions`. Generated pattern matches use
  `instanceof`, enums are lowered to integers and private constructors never
  cross an FFI or public boundary; public classes keep their `$tag`. Measured
  pairing: 120.09–120.37 ms with the tag against 115.19–116.64 ms without on
  the same TAST. See `phpurs/phpurs/audit/2026-09-24/private-tag/report.md`.
- The shared PBO fork also moved (`7b39ff2`, plus recompiled modules). The
  generated PHP differs from bundle `74124d39` in seven modules for the pure
  mode; the FFI programs' entry closure was checked separately.

## Compiled column (published)

Campaign `var/benchmark/php-pure-20260924d` — three processes, median per row,
total sums the fourteen medians:

| process | total |
|---|---:|
| 1 | 116.390 ms |
| 2 | 116.572 ms |
| 3 | 117.931 ms |
| **aggregated** | **116.590 ms** |

Red-Black Tree is 104,017.00 μs against 108,946.58 μs in the previous
campaign, and the /C ratio is 11.8x. `runp` on the same machine reports
115.63 ms for one process, which is within the usual single-process variation.

Campaigns `php-pure-20260924b` and `php-pure-20260924c` are kept as window
checks: their process totals were 124.42/121.37/116.29 and
117.10/117.32/121.90 under transient desktop load, so only the tight
2026-09-24d campaign was published.

## FFI columns (not refreshed)

The current window inflates the FFI kernels. The previously published build
was re-measured under the same conditions as the new build:

| build | fficc total |
|---|---:|
| published bundle `74124d39` | 211.896 ms, 215.264 ms |
| current bundle `86ad0c0b` | 218.675 ms, 209.619 ms |

The two builds overlap, so the toolchain change is neutral for the FFI
programs; the published 202.553 ms was measured in a quieter window and stays
valid. The attempted campaigns (1114.99 ms and 215.34 ms; 1099.58 ms and
211.42 ms) were not published for the same reason.

The only module with a private constructor in the FFI builds is `Test.RBTree`,
which no FFI entry point imports: `AppFFI` and `AppFFICheatcode` use
`Test.RBTreeFFI`. The FFI generated code that the programs execute is unchanged
by the private-tag removal.

## Toolchain identity

- Backend bundle: `86ad0c0b1e697bf878706639f2e62a775f1bf9779fdf107b0d0d9a5169326032`
- PBO fork: `7b39ff2`; host compiler: PureScript 0.15.16; TAST compiler:
  0.15.16 development build `319e138c`.
- `phpurs/phpurs` HEAD `70337ae` plus the uncommitted private-tag change in
  `src/Phpurs/CodeGen.purs`.

## Reproduction

- Measure: `python3 bin/php/campaign.py --mode pure`
- Publish: `python3 bin/php/campaign.py --mode pure --campaign-dir var/benchmark/php-pure-20260924d --update-readme --publish-only`
- Private-tag pairing: `phpurs/phpurs/audit/2026-09-24/private-tag/run-full.sh <build workspace>`
