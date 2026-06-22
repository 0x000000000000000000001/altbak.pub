import * as Data$dEuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Tail Call Optimization (10M calls):");
const deepTailRec = deepTailRec$a0$copy => deepTailRec$a1$copy => {
  let deepTailRec$a0 = deepTailRec$a0$copy, deepTailRec$a1 = deepTailRec$a1$copy, deepTailRec$c = true, deepTailRec$r;
  while (deepTailRec$c) {
    const v = deepTailRec$a0, v1 = deepTailRec$a1;
    if (v === 0) {
      deepTailRec$c = false;
      deepTailRec$r = v1;
      continue;
    }
    deepTailRec$a0 = v - 1 | 0;
    deepTailRec$a1 = v1 + Data$dEuclideanRing.intMod(v)(3) | 0;
  }
  return deepTailRec$r;
};
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ deepTailRec(10000000)(0)));
export {act, deepTailRec, describe};
