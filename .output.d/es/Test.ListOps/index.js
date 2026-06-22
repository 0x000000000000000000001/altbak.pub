import * as $runtime from "../runtime.js";
import * as Data$dEuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const $List = (tag, _1, _2) => ({tag, _1, _2});
const Nil = /* #__PURE__ */ $List("Nil");
const Cons = value0 => value1 => $List("Cons", value0, value1);
const range = start => end => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const curr = go$a0, acc = go$a1;
      if (curr < start) {
        go$c = false;
        go$r = acc;
        continue;
      }
      go$a0 = curr - 1 | 0;
      go$a1 = $List("Cons", curr, acc);
    }
    return go$r;
  };
  return go(end)(Nil);
};
const foldl = foldl$a0$copy => foldl$a1$copy => foldl$a2$copy => {
  let foldl$a0 = foldl$a0$copy, foldl$a1 = foldl$a1$copy, foldl$a2 = foldl$a2$copy, foldl$c = true, foldl$r;
  while (foldl$c) {
    const v = foldl$a0, v1 = foldl$a1, v2 = foldl$a2;
    if (v2.tag === "Nil") {
      foldl$c = false;
      foldl$r = v1;
      continue;
    }
    if (v2.tag === "Cons") {
      foldl$a0 = v;
      foldl$a1 = v(v1)(v2._1);
      foldl$a2 = v2._2;
      continue;
    }
    $runtime.fail();
  }
  return foldl$r;
};
const filterEvens = lst => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0, v1 = go$a1;
      if (v.tag === "Nil") {
        go$c = false;
        go$r = v1;
        continue;
      }
      if (v.tag === "Cons") {
        if (Data$dEuclideanRing.intMod(v._1)(2) === 0) {
          go$a0 = v._2;
          go$a1 = $List("Cons", v._1, v1);
          continue;
        }
        go$a0 = v._2;
        go$a1 = v1;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go(lst)(Nil);
};
const sumEvens = /* #__PURE__ */ foldl(Data$dSemiring.intAdd)(0)(/* #__PURE__ */ filterEvens(/* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const curr = go$a0, acc = go$a1;
      if (curr < 1) {
        go$c = false;
        go$r = acc;
        continue;
      }
      go$a0 = curr - 1 | 0;
      go$a1 = $List("Cons", curr, acc);
    }
    return go$r;
  };
  return go(90000)(Nil);
})()));
const describe = /* #__PURE__ */ Effect$dConsole.log("List Processing (90k elements):");
const act = /* #__PURE__ */ Effect$dConsole.log(/* #__PURE__ */ Data$dShow.showIntImpl(sumEvens));
export {$List, Cons, Nil, act, describe, filterEvens, foldl, range, sumEvens};
