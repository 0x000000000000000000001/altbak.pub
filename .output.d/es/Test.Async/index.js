import * as Effect$dConsole from "../Effect.Console/index.js";
import {runAsyncTest} from "./foreign.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Asynchronous Concurrency (1000 forks):");
const act = /* #__PURE__ */ (() => {
  const $0 = runAsyncTest(1000);
  return () => {
    $0();
    return Effect$dConsole.log("Done")();
  };
})();
export {act, describe};
export * from "./foreign.js";
