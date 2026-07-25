import * as Effect$dConsole from "../Effect.Console/index.js";
import {benchNow, formatNumber, opaque} from "./foreign.js";
const runBench = describe => act => () => {
  describe();
  const t1 = benchNow();
  act();
  const t2 = benchNow();
  const dt = t2 - t1;
  Effect$dConsole.log("\nExecution time: " + formatNumber(dt) + " μs\n")();
  return dt;
};
export {runBench};
export * from "./foreign.js";
