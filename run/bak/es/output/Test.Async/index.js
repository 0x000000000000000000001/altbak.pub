import * as Data$dEither from "../Data.Either/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Asynchronous Concurrency (Aff):");
const act = /* #__PURE__ */ (() => {
  const $0 = Effect$dAff._makeFiber(
    Effect$dAff.ffiUtil,
    Effect$dAff._bind(Effect$dAff.forkAff(Effect$dAff._bind(Effect$dAff._delay(Data$dEither.Right, 10.0))(() => Effect$dAff._bind(Effect$dAff._liftEffect(Effect$dConsole.log("Fiber 1 finished")))(() => Effect$dAff._pure()))))(() => Effect$dAff._bind(Effect$dAff.forkAff(Effect$dAff._bind(Effect$dAff._delay(
      Data$dEither.Right,
      20.0
    ))(() => Effect$dAff._bind(Effect$dAff._liftEffect(Effect$dConsole.log("Fiber 2 finished")))(() => Effect$dAff._pure()))))(() => Effect$dAff._bind(Effect$dAff._delay(
      Data$dEither.Right,
      30.0
    ))(() => Effect$dAff._bind(Effect$dAff._liftEffect(Effect$dConsole.log("Main fiber finished")))(() => Effect$dAff._pure()))))
  );
  return () => {
    const fiber = $0();
    fiber.run();
  };
})();
export {act, describe};
