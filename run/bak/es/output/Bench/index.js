import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
import {benchNow, formatNumber, keepAlive, opaque} from "./foreign.js";
const runBenchAff = describe => act => Effect$dAff._bind(Effect$dAff._liftEffect(describe))(() => Effect$dAff._bind(Effect$dAff._liftEffect(benchNow))(t1 => Effect$dAff._bind(act)(() => Effect$dAff._bind(Effect$dAff._liftEffect(benchNow))(t2 => {
  const dt = t2 - t1;
  return Effect$dAff._bind(Effect$dAff._liftEffect(Effect$dConsole.log("\nExecution time: " + formatNumber(dt) + " μs\n")))(() => Effect$dAff._pure(dt));
}))));
const runBench = describe => act => () => {
  describe();
  const t1 = benchNow();
  act();
  const t2 = benchNow();
  const dt = t2 - t1;
  Effect$dConsole.log("\nExecution time: " + formatNumber(dt) + " μs\n")();
  return dt;
};
export {runBench, runBenchAff};
export * from "./foreign.js";
