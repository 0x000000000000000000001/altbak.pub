import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const main = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ fib(40)));
export {fib, main};
