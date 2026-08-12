// | A module for formatting numbers as strings.
// |
// | Usage:
// | ``` purs
// | > let x = 1234.56789
// |
// | > toStringWith (precision 6) x
// | "1234.57"
// |
// | > toStringWith (fixed 3) x
// | "1234.568"
// |
// | > toStringWith (exponential 2) x
// | "1.23e+3"
// | ```
// |
// | The main method of this module is the `toStringWith` function that accepts
// | a `Format` argument which can be constructed through one of the smart
// | constructors `precision`, `fixed` and `exponential`. Internally, the
// | number will be formatted with JavaScripts `toPrecision`, `toFixed` or
// | `toExponential`.
import * as $runtime from "../runtime.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import {toExponentialNative, toFixedNative, toPrecisionNative, toString} from "./foreign.js";
const $Format = (tag, _1) => ({tag, _1});
const Precision = value0 => $Format("Precision", value0);
const Fixed = value0 => $Format("Fixed", value0);
const Exponential = value0 => $Format("Exponential", value0);
const toStringWith = v => {
  if (v.tag === "Precision") { return toPrecisionNative(v._1); }
  if (v.tag === "Fixed") { return toFixedNative(v._1); }
  if (v.tag === "Exponential") { return toExponentialNative(v._1); }
  $runtime.fail();
};
const precision = x => $Format(
  "Precision",
  (() => {
    const v = Data$dOrd.ordInt.compare(1)(x);
    const $0 = (() => {
      if (v === "LT") { return x; }
      if (v === "EQ") { return 1; }
      if (v === "GT") { return 1; }
      $runtime.fail();
    })();
    const v$1 = Data$dOrd.ordInt.compare(21)($0);
    if (v$1 === "LT") { return 21; }
    if (v$1 === "EQ") { return 21; }
    if (v$1 === "GT") { return $0; }
    $runtime.fail();
  })()
);
const fixed = x => $Format(
  "Fixed",
  (() => {
    const v = Data$dOrd.ordInt.compare(0)(x);
    const $0 = (() => {
      if (v === "LT") { return x; }
      if (v === "EQ") { return 0; }
      if (v === "GT") { return 0; }
      $runtime.fail();
    })();
    const v$1 = Data$dOrd.ordInt.compare(20)($0);
    if (v$1 === "LT") { return 20; }
    if (v$1 === "EQ") { return 20; }
    if (v$1 === "GT") { return $0; }
    $runtime.fail();
  })()
);
const exponential = x => $Format(
  "Exponential",
  (() => {
    const v = Data$dOrd.ordInt.compare(0)(x);
    const $0 = (() => {
      if (v === "LT") { return x; }
      if (v === "EQ") { return 0; }
      if (v === "GT") { return 0; }
      $runtime.fail();
    })();
    const v$1 = Data$dOrd.ordInt.compare(20)($0);
    if (v$1 === "LT") { return 20; }
    if (v$1 === "EQ") { return 20; }
    if (v$1 === "GT") { return $0; }
    $runtime.fail();
  })()
);
export {$Format, Exponential, Fixed, Precision, exponential, fixed, precision, toStringWith};
export * from "./foreign.js";
