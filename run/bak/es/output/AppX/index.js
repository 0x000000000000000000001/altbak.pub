import * as Bench from "../Bench/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
import * as Test$dAckermann from "../Test.Ackermann/index.js";
import * as Test$dAffOperations from "../Test.AffOperations/index.js";
import * as Test$dArrayOps from "../Test.ArrayOps/index.js";
import * as Test$dAstTree from "../Test.AstTree/index.js";
import * as Test$dChurch from "../Test.Church/index.js";
import * as Test$dFib from "../Test.Fib/index.js";
import * as Test$dFileOps from "../Test.FileOps/index.js";
import * as Test$dLazyEvaluation from "../Test.LazyEvaluation/index.js";
import * as Test$dListOps from "../Test.ListOps/index.js";
import * as Test$dParallelism from "../Test.Parallelism/index.js";
import * as Test$dPolymorphism from "../Test.Polymorphism/index.js";
import * as Test$dPrimes from "../Test.Primes/index.js";
import * as Test$dRBTree from "../Test.RBTree/index.js";
import * as Test$dRecords from "../Test.Records/index.js";
import * as Test$dSTArray from "../Test.STArray/index.js";
import * as Test$dStateMonad from "../Test.StateMonad/index.js";
import * as Test$dStringOps from "../Test.StringOps/index.js";
import * as Test$dTCO from "../Test.TCO/index.js";
const main = /* #__PURE__ */ (() => {
  const $0 = Bench.runBench(Test$dAstTree.describe)(Test$dAstTree.act);
  return () => {
    const t1 = $0();
    const t2 = Bench.runBench(Test$dFib.describe)(Test$dFib.act)();
    const t3 = Bench.runBench(Test$dListOps.describe)(Test$dListOps.act)();
    const t4 = Bench.runBench(Test$dTCO.describe)(Test$dTCO.act)();
    const t5 = Bench.runBench(Test$dRecords.describe)(Test$dRecords.act)();
    const t6 = Bench.runBench(Test$dAckermann.describe)(Test$dAckermann.act)();
    const t7 = Bench.runBench(Test$dChurch.describe)(Test$dChurch.act)();
    const t8 = Bench.runBench(Test$dPrimes.describe)(Test$dPrimes.act)();
    const t9 = Bench.runBench(Test$dRBTree.describe)(Test$dRBTree.act)();
    const t10 = Bench.runBench(Test$dPolymorphism.describe)(Test$dPolymorphism.act)();
    const t11 = Bench.runBench(Test$dStateMonad.describe)(Test$dStateMonad.act)();
    const t12 = Bench.runBench(Test$dLazyEvaluation.describe)(Test$dLazyEvaluation.act)();
    const t13 = Bench.runBench(Test$dArrayOps.describe)(Test$dArrayOps.act)();
    const t14 = Bench.runBench(Test$dFileOps.describe)(Test$dFileOps.act)();
    const t15 = Bench.runBench(Test$dSTArray.describe)(Test$dSTArray.act)();
    const t16 = Bench.runBench(Test$dStringOps.describe)(Test$dStringOps.act)();
    const fiber = Effect$dAff._makeFiber(
      Effect$dAff.ffiUtil,
      Effect$dAff._bind(Bench.runBenchAff(Test$dAffOperations.describe)(Test$dAffOperations.act))(t17 => Effect$dAff._bind(Bench.runBenchAff(Test$dParallelism.describe)(Test$dParallelism.act))(t18 => Effect$dAff._liftEffect(Effect$dConsole.log("\nTotal exec time: " + Bench.formatNumber(t1 / 1000.0 + t2 / 1000.0 + t3 / 1000.0 + t4 / 1000.0 + t5 / 1000.0 + t6 / 1000.0 + t7 / 1000.0 + t8 / 1000.0 + t9 / 1000.0 + t10 / 1000.0 + t11 / 1000.0 + t12 / 1000.0 + t13 / 1000.0 + t14 / 1000.0 + t15 / 1000.0 + t16 / 1000.0 + t17 / 1000.0 + t18 / 1000.0) + " ms\n"))))
    )();
    fiber.run();
  };
})();
export {main};
