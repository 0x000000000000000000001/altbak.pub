import * as Effect$dConsole from "../Effect.Console/index.js";
import {loopE, readFileSync, writeFileSync} from "./foreign.js";
const loopIO = n => loopE(n)((() => {
  const $0 = writeFileSync("var/iotest.txt")("Hello IO Benchmarks!");
  return () => {
    $0();
    readFileSync("var/iotest.txt")();
  };
})());
const describe = /* #__PURE__ */ Effect$dConsole.log("File I/O (10k writes/reads):");
const act = /* #__PURE__ */ loopIO(10000);
export {act, describe, loopIO};
export * from "./foreign.js";
