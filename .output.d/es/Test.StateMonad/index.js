import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const State = x => x;
const runState = v => s => v(s);
const put = s => v => ({val: undefined, state: s});
const pureState = a => s => ({val: a, state: s});
const $$get = s => ({val: s, state: s});
const describe = /* #__PURE__ */ Effect$dConsole.log("State Monad (12M Binds, 6k Stack Depth):");
const bindState = v => g => s => {
  const r1 = v(s);
  return g(r1.val)(r1.state);
};
const modify = f => s => ({val: undefined, state: f(s)});
const chainModifications = v => {
  if (v === 0) { return s => ({val: undefined, state: s}); }
  return s => chainModifications(v - 1 | 0)(s + 1 | 0);
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
    runManyTimes$a1 = v1 + chainModifications(6000)(0).state | 0;
  }
  return runManyTimes$r;
};
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ runManyTimes(2000)(0)));
export {State, act, bindState, chainModifications, describe, $$get as get, modify, pureState, put, runManyTimes, runState};
