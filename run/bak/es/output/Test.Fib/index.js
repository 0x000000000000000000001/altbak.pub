import * as Bench from "../Bench/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const describe = /* #__PURE__ */ Effect$dConsole.log("Fibonacci:");
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(10);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(fib(dummy)))();
  };
})();
export {act, describe, fib};
