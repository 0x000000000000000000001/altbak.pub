import * as Bench from "../Bench/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const act = () => {
  const t1 = Bench.benchNow();
  const t2 = Bench.benchNow();
  return Effect$dConsole.log("Delta: " + Data$dShow.showNumberImpl(t2 - t1))();
};
export {act};
