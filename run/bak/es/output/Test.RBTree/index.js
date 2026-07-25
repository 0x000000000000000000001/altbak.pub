import * as $runtime from "../runtime.js";
import * as Bench from "../Bench/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const $Color = tag => tag;
const $Tree = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const R = /* #__PURE__ */ $Color("R");
const B = /* #__PURE__ */ $Color("B");
const E = /* #__PURE__ */ $Tree("E");
const T = value0 => value1 => value2 => value3 => $Tree("T", value0, value1, value2, value3);
const max = x => y => {
  if (x > y) { return x; }
  return y;
};
const makeBlack = v => {
  if (v.tag === "T") { return $Tree("T", B, v._2, v._3, v._4); }
  if (v.tag === "E") { return E; }
  $runtime.fail();
};
const describe = /* #__PURE__ */ Effect$dConsole.log("Red-Black Tree (100k Worst-Case Insertions):");
const depth = v => {
  if (v.tag === "E") { return 0; }
  if (v.tag === "T") {
    const $0 = depth(v._2);
    const $1 = depth(v._4);
    if ($0 > $1) { return 1 + $0 | 0; }
    return 1 + $1 | 0;
  }
  $runtime.fail();
};
const balance = v => v1 => v2 => v3 => {
  if (v === "B") {
    if (v1.tag === "T") {
      if (v1._1 === "R") {
        if (v1._2.tag === "T") {
          if (v1._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1._2._2, v1._2._3, v1._2._4), v1._3, $Tree("T", B, v1._4, v2, v3)); }
          if (v1._4.tag === "T") {
            if (v1._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1._2, v1._3, v1._4._2), v1._4._3, $Tree("T", B, v1._4._4, v2, v3)); }
            if (v3.tag === "T" && v3._1 === "R") {
              if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
              if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
            }
            return $Tree("T", v, v1, v2, v3);
          }
          if (v3.tag === "T" && v3._1 === "R") {
            if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
            if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
          }
          return $Tree("T", v, v1, v2, v3);
        }
        if (v1._4.tag === "T") {
          if (v1._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1._2, v1._3, v1._4._2), v1._4._3, $Tree("T", B, v1._4._4, v2, v3)); }
          if (v3.tag === "T" && v3._1 === "R") {
            if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
            if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
          }
          return $Tree("T", v, v1, v2, v3);
        }
        if (v3.tag === "T" && v3._1 === "R") {
          if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
          if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
        }
        return $Tree("T", v, v1, v2, v3);
      }
      if (v3.tag === "T" && v3._1 === "R") {
        if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
        if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
      }
      return $Tree("T", v, v1, v2, v3);
    }
    if (v3.tag === "T" && v3._1 === "R") {
      if (v3._2.tag === "T" && v3._2._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2._2), v3._2._3, $Tree("T", B, v3._2._4, v3._3, v3._4)); }
      if (v3._4.tag === "T" && v3._4._1 === "R") { return $Tree("T", R, $Tree("T", B, v1, v2, v3._2), v3._3, $Tree("T", B, v3._4._2, v3._4._3, v3._4._4)); }
    }
  }
  return $Tree("T", v, v1, v2, v3);
};
const ins = v => v1 => {
  if (v1.tag === "E") { return $Tree("T", R, E, v, E); }
  if (v1.tag === "T") {
    if (v < v1._3) { return balance(v1._1)(ins(v)(v1._2))(v1._3)(v1._4); }
    if (v > v1._3) { return balance(v1._1)(v1._2)(v1._3)(ins(v)(v1._4)); }
    return $Tree("T", v1._1, v1._2, v1._3, v1._4);
  }
  $runtime.fail();
};
const insert = x => s => {
  const $0 = ins(x)(s);
  if ($0.tag === "T") { return $Tree("T", B, $0._2, $0._3, $0._4); }
  if ($0.tag === "E") { return E; }
  $runtime.fail();
};
const buildTree = buildTree$a0$copy => buildTree$a1$copy => {
  let buildTree$a0 = buildTree$a0$copy, buildTree$a1 = buildTree$a1$copy, buildTree$c = true, buildTree$r;
  while (buildTree$c) {
    const v = buildTree$a0, v1 = buildTree$a1;
    if (v === 0) {
      buildTree$c = false;
      buildTree$r = v1;
      continue;
    }
    buildTree$a0 = v - 1 | 0;
    buildTree$a1 = insert(v)(v1);
  }
  return buildTree$r;
};
const act = /* #__PURE__ */ (() => {
  const $0 = Bench.opaque(100000);
  return () => {
    const dummy = $0();
    return Effect$dConsole.log(Data$dShow.showIntImpl(depth(buildTree(dummy)(E))))();
  };
})();
export {$Color, $Tree, B, E, R, T, act, balance, buildTree, depth, describe, ins, insert, makeBlack, max};
