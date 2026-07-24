import * as Bench from "../Bench/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Ackermann (3, 4):");
const ackermann = v => v1 => {
  if (v === 0) { return v1 + 1 | 0; }
  if (v1 === 0) { return ackermann(v - 1 | 0)(1); }
  return ackermann(v - 1 | 0)(ackermann(v)(v1 - 1 | 0));
};
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(3);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(ackermann(dummy)(4)))();
  };
})();
export {ackermann, act, describe};
