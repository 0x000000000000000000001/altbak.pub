import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const zeroC = v => x => x;
const toInt = n => n(x => x + 1 | 0)(0);
const succC = n => f => x => f(n(f)(x));
const mulC = m => n => f => x => m(n(f))(x);
const fromInt = v => {
  if (v === 0) { return zeroC; }
  const $0 = fromInt(v - 1 | 0);
  return f => x => f($0(f)(x));
};
const describe = /* #__PURE__ */ Effect$dConsole.log("Church Numerals (100k Closure Applications):");
const c10 = /* #__PURE__ */ fromInt(10);
const c100 = f => x => c10(c10(f))(x);
const c10k = f => x => c10(c10(c100(f)))(x);
const c100k = f => x => c10(c10(c100(c10(f))))(x);
const addC = m => n => f => x => m(f)(n(f)(x));
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ c10(/* #__PURE__ */ c10(/* #__PURE__ */ c100(/* #__PURE__ */ c10(x => x + 1 | 0))))(0)));
export {act, addC, c10, c100, c100k, c10k, describe, fromInt, mulC, succC, toInt, zeroC};
