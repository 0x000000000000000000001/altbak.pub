import * as $runtime from "../runtime.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Data$dString$dRegex from "../Data.String.Regex/index.js";
import * as Data$dString$dRegex$dFlags from "../Data.String.Regex.Flags/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const regexPattern = /* #__PURE__ */ (() => {
  const v = Data$dString$dRegex.regex("(hello|world)[0-9]+")(Data$dString$dRegex$dFlags.noFlags);
  if (v.tag === "Right") { return v._1; }
  $runtime.fail();
})();
const runStringOps = n => {
  const loop = loop$a0$copy => loop$a1$copy => loop$a2$copy => {
    let loop$a0 = loop$a0$copy, loop$a1 = loop$a1$copy, loop$a2 = loop$a2$copy, loop$c = true, loop$r;
    while (loop$c) {
      const v = loop$a0, v1 = loop$a1, v2 = loop$a2;
      if (v === 0) {
        loop$c = false;
        loop$r = v2;
        continue;
      }
      const concatted = v1 + Data$dShow.showIntImpl(v) + "world" + Data$dShow.showIntImpl(v + 1 | 0);
      loop$a0 = v - 1 | 0;
      loop$a1 = Data$dString$dCodePoints.take(10)(concatted);
      loop$a2 = v2 + Data$dString$dCommon.split("e")(Data$dString$dRegex.replace(regexPattern)("matched")(concatted)).length | 0;
    }
    return loop$r;
  };
  return loop(n)("hello")(0);
};
const describe = /* #__PURE__ */ Effect$dConsole.log("String Operations (1k Regex/Split):");
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ runStringOps(1000)));
export {act, describe, regexPattern, runStringOps};
