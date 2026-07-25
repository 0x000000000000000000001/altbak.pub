import * as Data$dEither from "../Data.Either/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Aff Operations (Asynchronous Delays):");
const act = /* #__PURE__ */ (() => {
  const $0 = Effect$dAff._makeFiber(Effect$dAff.ffiUtil, Effect$dAff._bind(Effect$dAff._delay(Data$dEither.Right, 10.0))(() => Effect$dAff._liftEffect(Effect$dConsole.log("10"))));
  return () => {
    const fiber = $0();
    fiber.run();
  };
})();
export {act, describe};
