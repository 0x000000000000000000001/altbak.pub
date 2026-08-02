import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const traverse = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Effect$dAff.applicativeAff))();
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const heavyTask = n => Effect$dAff._bind(Effect$dAff._delay(Data$dEither.Right, 0.0))(() => Effect$dAff._pure());
const describe = /* #__PURE__ */ Effect$dConsole.log("Parallelism (40 x Fib 35):");
const act = /* #__PURE__ */ (() => {
  const $0 = Effect$dAff._makeFiber(
    Effect$dAff.ffiUtil,
    Effect$dAff._bind(traverse(v => Effect$dAff.forkAff(Effect$dAff._bind(Effect$dAff._delay(Data$dEither.Right, 0.0))(() => Effect$dAff._pure())))(Data$dArray.replicateImpl(
      200,
      undefined
    )))(fibers => Effect$dAff._bind(traverse(Effect$dAff.joinFiber)(fibers))(() => Effect$dAff._pure()))
  );
  return () => {
    const fiber = $0();
    fiber.run();
  };
})();
export {act, describe, fib, heavyTask, traverse};
