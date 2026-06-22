import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
import * as Effect$dNow from "../Effect.Now/index.js";
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const main = () => {
  const time = Effect$dNow.now();
  Effect$dConsole.log("Now: (Instant (Milliseconds " + Data$dShow.showNumberImpl(time) + "))")();
  return Effect$dConsole.log(Data$dShow.showIntImpl(fib(40)))();
};
export {fib, main};
