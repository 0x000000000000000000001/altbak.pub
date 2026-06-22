import * as Data$dEuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const updateRec = updateRec$a0$copy => updateRec$a1$copy => {
  let updateRec$a0 = updateRec$a0$copy, updateRec$a1 = updateRec$a1$copy, updateRec$c = true, updateRec$r;
  while (updateRec$c) {
    const v = updateRec$a0, v1 = updateRec$a1;
    if (v === 0) {
      updateRec$c = false;
      updateRec$r = v1;
      continue;
    }
    updateRec$a0 = v - 1 | 0;
    updateRec$a1 = {...v1, a: v1.a + 1 | 0, b: {...v1.b, c: v1.b.c + 2 | 0, d: {...v1.b.d, e: v1.b.d.e + 3 | 0, f: v1.b.d.f + Data$dEuclideanRing.intMod(v)(5) | 0}}};
  }
  return updateRec$r;
};
const initial = {a: 0, b: {c: 0, d: {e: 0, f: 0}}};
const describe = /* #__PURE__ */ Effect$dConsole.log("Deep Record Updates (1M iterations):");
const act = /* #__PURE__ */ (() => Effect$dConsole.log(Data$dShow.showIntImpl(updateRec(1000000)(initial).b.d.f)))();
export {act, describe, initial, updateRec};
