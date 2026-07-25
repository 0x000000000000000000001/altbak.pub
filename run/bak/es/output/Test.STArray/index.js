import * as $runtime from "../runtime.js";
import * as Data$dArray$dST from "../Data.Array.ST/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const sumArray = /* #__PURE__ */ (() => {
  const arr = [];
  arr.push(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
  const x = Data$dArray$dST.popImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, arr);
  if (x.tag === "Just") { return x._1; }
  if (x.tag === "Nothing") { return 0; }
  $runtime.fail();
})();
const describe = /* #__PURE__ */ Effect$dConsole.log("STArray Operations:");
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(sumArray));
export {act, describe, sumArray};
