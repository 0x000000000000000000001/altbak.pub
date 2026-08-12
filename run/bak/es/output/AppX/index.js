import * as Bench from "../Bench/index.js";
import * as Test$dRBTree from "../Test.RBTree/index.js";
const main = /* #__PURE__ */ (() => {
  const $0 = Bench.runBench(Test$dRBTree.describe)(Test$dRBTree.act);
  return () => {$0();};
})();
export {main};
