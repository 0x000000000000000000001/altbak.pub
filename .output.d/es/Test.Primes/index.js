import * as $runtime from "../runtime.js";
import * as Data$dEuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const $List = (tag, _1, _2) => ({tag, _1, _2});
const Nil = /* #__PURE__ */ $List("Nil");
const Cons = value0 => value1 => $List("Cons", value0, value1);
const sumList = lst => {
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
        go$a0 = v._2;
        go$a1 = v1 + v._1 | 0;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go(lst)(0);
};
const reverse = lst => {
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
        go$a0 = v._2;
        go$a1 = $List("Cons", v._1, v1);
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go(lst)(Nil);
};
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
const filter = p => lst => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0, v1 = go$a1;
      if (v.tag === "Nil") {
        const go$1 = go$1$a0$copy => go$1$a1$copy => {
          let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
          while (go$1$c) {
            const v$1 = go$1$a0, v1$1 = go$1$a1;
            if (v$1.tag === "Nil") {
              go$1$c = false;
              go$1$r = v1$1;
              continue;
            }
            if (v$1.tag === "Cons") {
              go$1$a0 = v$1._2;
              go$1$a1 = $List("Cons", v$1._1, v1$1);
              continue;
            }
            $runtime.fail();
          }
          return go$1$r;
        };
        go$c = false;
        go$r = go$1(v1)(Nil);
        continue;
      }
      if (v.tag === "Cons") {
        if (p(v._1)) {
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
const sieve = v => {
  if (v.tag === "Nil") { return Nil; }
  if (v.tag === "Cons") {
    const $0 = v._1;
    return $List(
      "Cons",
      $0,
      sieve((() => {
        const go = go$a0$copy => go$a1$copy => {
          let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
          while (go$c) {
            const v$1 = go$a0, v1 = go$a1;
            if (v$1.tag === "Nil") {
              const go$1 = go$1$a0$copy => go$1$a1$copy => {
                let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
                while (go$1$c) {
                  const v$2 = go$1$a0, v1$1 = go$1$a1;
                  if (v$2.tag === "Nil") {
                    go$1$c = false;
                    go$1$r = v1$1;
                    continue;
                  }
                  if (v$2.tag === "Cons") {
                    go$1$a0 = v$2._2;
                    go$1$a1 = $List("Cons", v$2._1, v1$1);
                    continue;
                  }
                  $runtime.fail();
                }
                return go$1$r;
              };
              go$c = false;
              go$r = go$1(v1)(Nil);
              continue;
            }
            if (v$1.tag === "Cons") {
              if (Data$dEuclideanRing.intMod(v$1._1)($0) !== 0) {
                go$a0 = v$1._2;
                go$a1 = $List("Cons", v$1._1, v1);
                continue;
              }
              go$a0 = v$1._2;
              go$a1 = v1;
              continue;
            }
            $runtime.fail();
          }
          return go$r;
        };
        return go(v._2)(Nil);
      })())
    );
  }
  $runtime.fail();
};
const describe = /* #__PURE__ */ Effect$dConsole.log("Prime Sieve (sum primes up to 20k):");
const act = /* #__PURE__ */ (() => {
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
        go$a0 = v._2;
        go$a1 = v1 + v._1 | 0;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return Effect$dConsole.log(Data$dShow.showIntImpl(go(sieve((() => {
    const go$1 = go$1$a0$copy => go$1$a1$copy => {
      let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
      while (go$1$c) {
        const curr = go$1$a0, acc = go$1$a1;
        if (curr < 2) {
          go$1$c = false;
          go$1$r = acc;
          continue;
        }
        go$1$a0 = curr - 1 | 0;
        go$1$a1 = $List("Cons", curr, acc);
      }
      return go$1$r;
    };
    return go$1(20000)(Nil);
  })()))(0)));
})();
export {$List, Cons, Nil, act, describe, filter, range, reverse, sieve, sumList};
