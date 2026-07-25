import * as Bench from "../Bench/index.js";
import * as Test$dTCO from "../Test.TCO/index.js";
const main = /* #__PURE__ */ (() => {
  const $0 = Bench.runBench(Test$dTCO.describe)(Test$dTCO.act);
  return () => {$0();};
})();
export {main};
