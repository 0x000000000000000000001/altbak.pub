import * as Effect$dConsole from "../Effect.Console/index.js";
import {runAsyncTest} from "./foreign.js";
const describe = /* #__PURE__ */ Effect$dConsole.log("Asynchronous Concurrency (1000 forks):");
const act = /* #__PURE__ */ runAsyncTest(1000);
export {act, describe};
export * from "./foreign.js";
