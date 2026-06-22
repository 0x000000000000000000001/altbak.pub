import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const mempty_ = dict => dict.mempty_;
const mappend_ = dict => dict.mappend_;
const polyLoop = dictMonoidish => {
  const mempty_1 = dictMonoidish.mempty_;
  return n_init => acc_init => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v = go$a0, v1 = go$a1;
        if (v === 0) {
          go$c = false;
          go$r = v1;
          continue;
        }
        go$a0 = v - 1 | 0;
        go$a1 = dictMonoidish.mappend_(v1)(mempty_1);
      }
      return go$r;
    };
    return go(n_init)(acc_init);
  };
};
const intMonoidish = {mempty_: 1, mappend_: x => y => x + y | 0};
const describe = /* #__PURE__ */ Effect$dConsole.log("Polymorphism (10M Type Class Dict Lookups):");
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(/* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0, v1 = go$a1;
      if (v === 0) {
        go$c = false;
        go$r = v1;
        continue;
      }
      go$a0 = v - 1 | 0;
      go$a1 = v1 + 1 | 0;
    }
    return go$r;
  };
  return go(10000000)(0);
})()));
export {act, describe, intMonoidish, mappend_, mempty_, polyLoop};
