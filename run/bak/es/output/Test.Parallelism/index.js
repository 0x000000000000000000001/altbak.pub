import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const traverse = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Effect$dAff.applicativeAff))();
const sum = /* #__PURE__ */ Data$dFoldable.foldlArray(Data$dSemiring.intAdd)(0);
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const heavyTask = n => Effect$dAff._bind(Effect$dAff._delay(Data$dEither.Right, 0.0))(() => Effect$dAff._pure(fib(n)));
const describe = /* #__PURE__ */ Effect$dConsole.log("Parallelism (4 x Fib 42):");
const act = /* #__PURE__ */ Effect$dAff._bind(/* #__PURE__ */ traverse(v => Effect$dAff.forkAff(heavyTask(42)))(/* #__PURE__ */ Data$dArray.replicateImpl(4, undefined)))(fibers => Effect$dAff._bind(traverse(Effect$dAff.joinFiber)(fibers))(results => Effect$dAff._liftEffect(Effect$dConsole.log("Sum of results: " + Data$dShow.showIntImpl(sum(results))))));
export {act, describe, fib, heavyTask, sum, traverse};
