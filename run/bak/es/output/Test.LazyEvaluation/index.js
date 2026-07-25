import * as Bench from "../Bench/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const Lazy = x => x;
const force = v => v();
const describe = /* #__PURE__ */ Effect$dConsole.log("Lazy Evaluation (1M Thunks Forced, 1k Depth):");
const defer = f => f;
const buildThunks = buildThunks$a0$copy => buildThunks$a1$copy => {
  let buildThunks$a0 = buildThunks$a0$copy, buildThunks$a1 = buildThunks$a1$copy, buildThunks$c = true, buildThunks$r;
  while (buildThunks$c) {
    const v = buildThunks$a0, v1 = buildThunks$a1;
    if (v === 0) {
      buildThunks$c = false;
      buildThunks$r = v1;
      continue;
    }
    buildThunks$a0 = v - 1 | 0;
    buildThunks$a1 = v2 => v1() + 1 | 0;
  }
  return buildThunks$r;
};
const runManyTimes = runManyTimes$a0$copy => runManyTimes$a1$copy => {
  let runManyTimes$a0 = runManyTimes$a0$copy, runManyTimes$a1 = runManyTimes$a1$copy, runManyTimes$c = true, runManyTimes$r;
  while (runManyTimes$c) {
    const v = runManyTimes$a0, v1 = runManyTimes$a1;
    if (v === 0) {
      runManyTimes$c = false;
      runManyTimes$r = v1;
      continue;
    }
    runManyTimes$a0 = v - 1 | 0;
    runManyTimes$a1 = v1 + buildThunks(1000)(v2 => 0)() | 0;
  }
  return runManyTimes$r;
};
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(1000);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(runManyTimes(dummy)(0)))();
  };
})();
export {Lazy, act, buildThunks, defer, describe, force, runManyTimes};
