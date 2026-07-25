import * as Bench from "../Bench/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const range = start => end => Data$dArray.rangeImpl(start, end);
const filterEvens = arr => Data$dArray.filterImpl(x => Data$dEuclideanRing.intMod(x)(2) === 0, arr);
const sumEvens = n => Data$dFoldable.foldlArray(Data$dSemiring.intAdd)(0)(Data$dArray.filterImpl(x => Data$dEuclideanRing.intMod(x)(2) === 0, Data$dArray.rangeImpl(1, n)));
const describe = /* #__PURE__ */ Effect$dConsole.log("Array Processing (900 elements):");
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(900);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(sumEvens(dummy)))();
  };
})();
export {act, describe, filterEvens, range, sumEvens};
