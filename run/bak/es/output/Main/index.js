import * as Effect$dConsole from "../Effect.Console/index.js";
const main = /* #__PURE__ */ (() => {
  const $0 = Effect$dConsole.log("Hello");
  return () => {
    $0();
    return Effect$dConsole.log("World")();
  };
})();
export {main};
