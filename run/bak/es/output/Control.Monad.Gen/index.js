import * as $runtime from "../runtime.js";
import * as Control$dMonad$dRec$dClass from "../Control.Monad.Rec.Class/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dSemigroup$dLast from "../Data.Semigroup.Last/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const $LL = (tag, _1, _2) => ({tag, _1, _2});
const monoidAdditive = /* #__PURE__ */ (() => {
  const semigroupAdditive1 = {append: v => v1 => v + v1};
  return {mempty: 0.0, Semigroup0: () => semigroupAdditive1};
})();
const Cons = value0 => value1 => $LL("Cons", value0, value1);
const Nil = /* #__PURE__ */ $LL("Nil");
const unfoldable = dictMonadRec => dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const $0 = Monad0.Applicative0();
  const Bind1 = Monad0.Bind1();
  const Functor0 = Monad0.Bind1().Apply0().Functor0();
  return dictUnfoldable => gen => Functor0.map(dictUnfoldable.unfoldr(v => {
    if (v.tag === "Nil") { return Data$dMaybe.Nothing; }
    if (v.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v._1, v._2)); }
    $runtime.fail();
  }))(dictMonadGen.sized((() => {
    const $1 = dictMonadRec.tailRecM(v => {
      if (v._2 <= 0) { return $0.pure(Control$dMonad$dRec$dClass.$Step("Done", v._1)); }
      const $1 = v._1;
      const $2 = v._2;
      return Bind1.bind(gen)(x => $0.pure(Control$dMonad$dRec$dClass.$Step("Loop", Data$dTuple.$Tuple($LL("Cons", x, $1), $2 - 1 | 0))));
    });
    const $2 = Data$dTuple.Tuple(Nil);
    return x => $1($2(x));
  })()));
};
const semigroupFreqSemigroup = {
  append: v => v1 => pos => {
    const v2 = v(pos);
    if (v2._1.tag === "Just") { return v1(v2._1._1); }
    return v2;
  }
};
const fromIndex = dictFoldable1 => {
  const Foldable0 = dictFoldable1.Foldable0();
  return i => xs => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v = go$a0, v1 = go$a1;
        if (v1.tag === "Cons") {
          if (v1._2.tag === "Nil") {
            go$c = false;
            go$r = v1._1;
            continue;
          }
          if (v <= 0) {
            go$c = false;
            go$r = v1._1;
            continue;
          }
          go$a0 = v - 1 | 0;
          go$a1 = v1._2;
          continue;
        }
        if (v1.tag === "Nil") {
          go$c = false;
          go$r = dictFoldable1.foldMap1(Data$dSemigroup$dLast.semigroupLast)(Data$dSemigroup$dLast.Last)(xs);
          continue;
        }
        $runtime.fail();
      }
      return go$r;
    };
    return go(i)(Foldable0.foldr(Cons)(Nil)(xs));
  };
};
const oneOf = dictMonadGen => {
  const Bind1 = dictMonadGen.Monad0().Bind1();
  return dictFoldable1 => {
    const Foldable0 = dictFoldable1.Foldable0();
    return xs => Bind1.bind(dictMonadGen.chooseInt(0)(Foldable0.foldl(c => v => 1 + c | 0)(0)(xs) - 1 | 0))(n => fromIndex(dictFoldable1)(n)(xs));
  };
};
const freqSemigroup = v => {
  const $0 = v._1;
  const $1 = v._2;
  return pos => {
    if (pos >= $0) { return Data$dTuple.$Tuple(Data$dMaybe.$Maybe("Just", pos - $0), $1); }
    return Data$dTuple.$Tuple(Data$dMaybe.Nothing, $1);
  };
};
const frequency = dictMonadGen => {
  const Bind1 = dictMonadGen.Monad0().Bind1();
  return dictFoldable1 => {
    const foldMap = dictFoldable1.Foldable0().foldMap(monoidAdditive);
    return xs => Bind1.bind(dictMonadGen.chooseFloat(0.0)(foldMap(Data$dTuple.fst)(xs)))((() => {
      const $0 = dictFoldable1.foldMap1(semigroupFreqSemigroup)(freqSemigroup)(xs);
      return x => $0(x)._2;
    })());
  };
};
const filtered = dictMonadRec => dictMonadGen => {
  const Functor0 = dictMonadGen.Monad0().Bind1().Apply0().Functor0();
  return gen => dictMonadRec.tailRecM(v => Functor0.map(a => {
    if (a.tag === "Nothing") { return Control$dMonad$dRec$dClass.$Step("Loop", undefined); }
    if (a.tag === "Just") { return Control$dMonad$dRec$dClass.$Step("Done", a._1); }
    $runtime.fail();
  })(gen))();
};
const suchThat = dictMonadRec => dictMonadGen => {
  const Functor0 = dictMonadGen.Monad0().Bind1().Apply0().Functor0();
  const Functor0$1 = dictMonadGen.Monad0().Bind1().Apply0().Functor0();
  return gen => pred => {
    const $0 = Functor0$1.map(a => {
      if (pred(a)) { return Data$dMaybe.$Maybe("Just", a); }
      return Data$dMaybe.Nothing;
    })(gen);
    return dictMonadRec.tailRecM(v => Functor0.map(a => {
      if (a.tag === "Nothing") { return Control$dMonad$dRec$dClass.$Step("Loop", undefined); }
      if (a.tag === "Just") { return Control$dMonad$dRec$dClass.$Step("Done", a._1); }
      $runtime.fail();
    })($0))();
  };
};
const elements = dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const Bind1 = Monad0.Bind1();
  return dictFoldable1 => {
    const Foldable0 = dictFoldable1.Foldable0();
    return xs => Bind1.bind(dictMonadGen.chooseInt(0)(Foldable0.foldl(c => v => 1 + c | 0)(0)(xs) - 1 | 0))(n => Monad0.Applicative0().pure(fromIndex(dictFoldable1)(n)(xs)));
  };
};
const choose = dictMonadGen => {
  const Bind1 = dictMonadGen.Monad0().Bind1();
  const chooseBool = dictMonadGen.chooseBool;
  return genA => genB => Bind1.bind(chooseBool)(v => {
    if (v) { return genA; }
    return genB;
  });
};
export {$LL, Cons, Nil, choose, elements, filtered, freqSemigroup, frequency, fromIndex, monoidAdditive, oneOf, semigroupFreqSemigroup, suchThat, unfoldable};
