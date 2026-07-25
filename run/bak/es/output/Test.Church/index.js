import * as Bench from "../Bench/index.js";
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
const c10 = n => fromInt(n);
const c100 = n => {
  const $0 = fromInt(n);
  const $1 = fromInt(n);
  return f => x => $0($1(f))(x);
};
const c10k = n => {
  const $0 = c100(n);
  const $1 = c100(n);
  return f => x => $0($1(f))(x);
};
const c100k = n => {
  const $0 = c10k(n);
  const $1 = fromInt(n);
  return f => x => $0($1(f))(x);
};
const addC = m => n => f => x => m(f)(n(f)(x));
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(10);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(c100k(dummy)(x => x + 1 | 0)(0)))();
  };
})();
export {act, addC, c10, c100, c100k, c10k, describe, fromInt, mulC, succC, toInt, zeroC};
