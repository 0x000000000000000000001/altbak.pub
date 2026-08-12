import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dInterval$dDuration from "../Data.Interval.Duration/index.js";
import * as Data$dList from "../Data.List/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dNonEmpty from "../Data.NonEmpty/index.js";
import * as Data$dNumber from "../Data.Number/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
const $$$Error = (tag, _1) => ({tag, _1});
const foldMap = /* #__PURE__ */ (() => Data$dList$dTypes.foldableList.foldMap(Data$dList$dTypes.monoidList))();
const monoidAdditive = /* #__PURE__ */ (() => {
  const semigroupAdditive1 = {append: v => v1 => v + v1};
  return {mempty: 0.0, Semigroup0: () => semigroupAdditive1};
})();
const monoidFn = /* #__PURE__ */ (() => {
  const semigroupFn = {append: f => g => x => Data$dList$dTypes.foldableList.foldr(Data$dList$dTypes.Cons)(g(x))(f(x))};
  return {mempty: v => Data$dList$dTypes.Nil, Semigroup0: () => semigroupFn};
})();
const IsEmpty = /* #__PURE__ */ $$$Error("IsEmpty");
const InvalidWeekComponentUsage = /* #__PURE__ */ $$$Error("InvalidWeekComponentUsage");
const ContainsNegativeValue = value0 => $$$Error("ContainsNegativeValue", value0);
const InvalidFractionalUse = value0 => $$$Error("InvalidFractionalUse", value0);
const unIsoDuration = v => v;
const showIsoDuration = {show: v => "(IsoDuration (Duration " + Data$dInterval$dDuration.showMap.show(v) + "))"};
const showError = {
  show: v => {
    if (v.tag === "IsEmpty") { return "(IsEmpty)"; }
    if (v.tag === "InvalidWeekComponentUsage") { return "(InvalidWeekComponentUsage)"; }
    if (v.tag === "ContainsNegativeValue") {
      if (v._1 === "Minute") { return "(ContainsNegativeValue Minute)"; }
      if (v._1 === "Second") { return "(ContainsNegativeValue Second)"; }
      if (v._1 === "Hour") { return "(ContainsNegativeValue Hour)"; }
      if (v._1 === "Day") { return "(ContainsNegativeValue Day)"; }
      if (v._1 === "Week") { return "(ContainsNegativeValue Week)"; }
      if (v._1 === "Month") { return "(ContainsNegativeValue Month)"; }
      if (v._1 === "Year") { return "(ContainsNegativeValue Year)"; }
      $runtime.fail();
    }
    if (v.tag === "InvalidFractionalUse") {
      if (v._1 === "Minute") { return "(InvalidFractionalUse Minute)"; }
      if (v._1 === "Second") { return "(InvalidFractionalUse Second)"; }
      if (v._1 === "Hour") { return "(InvalidFractionalUse Hour)"; }
      if (v._1 === "Day") { return "(InvalidFractionalUse Day)"; }
      if (v._1 === "Week") { return "(InvalidFractionalUse Week)"; }
      if (v._1 === "Month") { return "(InvalidFractionalUse Month)"; }
      if (v._1 === "Year") { return "(InvalidFractionalUse Year)"; }
    }
    $runtime.fail();
  }
};
const prettyError = v => {
  if (v.tag === "IsEmpty") { return "Duration is empty (has no components)"; }
  if (v.tag === "InvalidWeekComponentUsage") { return "Week component of Duration is used with other components"; }
  if (v.tag === "ContainsNegativeValue") {
    if (v._1 === "Minute") { return "Component `Minute` contains negative value"; }
    if (v._1 === "Second") { return "Component `Second` contains negative value"; }
    if (v._1 === "Hour") { return "Component `Hour` contains negative value"; }
    if (v._1 === "Day") { return "Component `Day` contains negative value"; }
    if (v._1 === "Week") { return "Component `Week` contains negative value"; }
    if (v._1 === "Month") { return "Component `Month` contains negative value"; }
    if (v._1 === "Year") { return "Component `Year` contains negative value"; }
    $runtime.fail();
  }
  if (v.tag === "InvalidFractionalUse") {
    if (v._1 === "Minute") { return "Invalid usage of Fractional value at component `Minute`"; }
    if (v._1 === "Second") { return "Invalid usage of Fractional value at component `Second`"; }
    if (v._1 === "Hour") { return "Invalid usage of Fractional value at component `Hour`"; }
    if (v._1 === "Day") { return "Invalid usage of Fractional value at component `Day`"; }
    if (v._1 === "Week") { return "Invalid usage of Fractional value at component `Week`"; }
    if (v._1 === "Month") { return "Invalid usage of Fractional value at component `Month`"; }
    if (v._1 === "Year") { return "Invalid usage of Fractional value at component `Year`"; }
  }
  $runtime.fail();
};
const eqIsoDuration = {eq: x => y => Data$dInterval$dDuration.eqMap.eq(x)(y)};
const ordIsoDuration = {compare: x => y => Data$dInterval$dDuration.ordMap.compare(x)(y), Eq0: () => eqIsoDuration};
const eqError = {
  eq: x => y => {
    if (x.tag === "IsEmpty") { return y.tag === "IsEmpty"; }
    if (x.tag === "InvalidWeekComponentUsage") { return y.tag === "InvalidWeekComponentUsage"; }
    if (x.tag === "ContainsNegativeValue") {
      return y.tag === "ContainsNegativeValue" && (() => {
        if (x._1 === "Second") { return y._1 === "Second"; }
        if (x._1 === "Minute") { return y._1 === "Minute"; }
        if (x._1 === "Hour") { return y._1 === "Hour"; }
        if (x._1 === "Day") { return y._1 === "Day"; }
        if (x._1 === "Week") { return y._1 === "Week"; }
        if (x._1 === "Month") { return y._1 === "Month"; }
        return x._1 === "Year" && y._1 === "Year";
      })();
    }
    return x.tag === "InvalidFractionalUse" && y.tag === "InvalidFractionalUse" && (() => {
      if (x._1 === "Second") { return y._1 === "Second"; }
      if (x._1 === "Minute") { return y._1 === "Minute"; }
      if (x._1 === "Hour") { return y._1 === "Hour"; }
      if (x._1 === "Day") { return y._1 === "Day"; }
      if (x._1 === "Week") { return y._1 === "Week"; }
      if (x._1 === "Month") { return y._1 === "Month"; }
      return x._1 === "Year" && y._1 === "Year";
    })();
  }
};
const ordError = {
  compare: x => y => {
    if (x.tag === "IsEmpty") {
      if (y.tag === "IsEmpty") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsEmpty") { return Data$dOrdering.GT; }
    if (x.tag === "InvalidWeekComponentUsage") {
      if (y.tag === "InvalidWeekComponentUsage") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "InvalidWeekComponentUsage") { return Data$dOrdering.GT; }
    if (x.tag === "ContainsNegativeValue") {
      if (y.tag === "ContainsNegativeValue") {
        if (x._1 === "Second") {
          if (y._1 === "Second") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Second") { return Data$dOrdering.GT; }
        if (x._1 === "Minute") {
          if (y._1 === "Minute") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Minute") { return Data$dOrdering.GT; }
        if (x._1 === "Hour") {
          if (y._1 === "Hour") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Hour") { return Data$dOrdering.GT; }
        if (x._1 === "Day") {
          if (y._1 === "Day") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Day") { return Data$dOrdering.GT; }
        if (x._1 === "Week") {
          if (y._1 === "Week") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Week") { return Data$dOrdering.GT; }
        if (x._1 === "Month") {
          if (y._1 === "Month") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "Month") { return Data$dOrdering.GT; }
        if (x._1 === "Year" && y._1 === "Year") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "ContainsNegativeValue") { return Data$dOrdering.GT; }
    if (x.tag === "InvalidFractionalUse" && y.tag === "InvalidFractionalUse") {
      if (x._1 === "Second") {
        if (y._1 === "Second") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Second") { return Data$dOrdering.GT; }
      if (x._1 === "Minute") {
        if (y._1 === "Minute") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Minute") { return Data$dOrdering.GT; }
      if (x._1 === "Hour") {
        if (y._1 === "Hour") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Hour") { return Data$dOrdering.GT; }
      if (x._1 === "Day") {
        if (y._1 === "Day") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Day") { return Data$dOrdering.GT; }
      if (x._1 === "Week") {
        if (y._1 === "Week") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Week") { return Data$dOrdering.GT; }
      if (x._1 === "Month") {
        if (y._1 === "Month") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "Month") { return Data$dOrdering.GT; }
      if (x._1 === "Year" && y._1 === "Year") { return Data$dOrdering.EQ; }
    }
    $runtime.fail();
  },
  Eq0: () => eqError
};
const checkWeekUsage = v => {
  if (
    (() => {
      const go = go$a0$copy => {
        let go$a0 = go$a0$copy, go$c = true, go$r;
        while (go$c) {
          const v$1 = go$a0;
          if (v$1.tag === "Leaf") {
            go$c = false;
            go$r = Data$dMaybe.Nothing;
            continue;
          }
          if (v$1.tag === "Node") {
            if (v$1._3 === "Second") {
              go$a0 = v$1._6;
              continue;
            }
            if (v$1._3 === "Minute") {
              go$a0 = v$1._6;
              continue;
            }
            if (v$1._3 === "Hour") {
              go$a0 = v$1._6;
              continue;
            }
            if (v$1._3 === "Day") {
              go$a0 = v$1._6;
              continue;
            }
            if (v$1._3 === "Week") {
              go$c = false;
              go$r = Data$dMaybe.$Maybe("Just", v$1._4);
              continue;
            }
            go$a0 = v$1._5;
            continue;
          }
          $runtime.fail();
        }
        return go$r;
      };
      const $0 = go(v.asMap);
      return (() => {
        if ($0.tag === "Nothing") { return false; }
        if ($0.tag === "Just") { return true; }
        $runtime.fail();
      })() && (() => {
        if (v.asMap.tag === "Leaf") { return false; }
        if (v.asMap.tag === "Node") { return v.asMap._2 > 1; }
        $runtime.fail();
      })();
    })()
  ) {
    return Data$dList$dTypes.$List("Cons", InvalidWeekComponentUsage, Data$dList$dTypes.Nil);
  }
  return Data$dList$dTypes.Nil;
};
const checkNegativeValues = v => foldMap(v1 => {
  if (v1._2 >= 0.0) { return Data$dList$dTypes.Nil; }
  return Data$dList$dTypes.$List("Cons", $$$Error("ContainsNegativeValue", v1._1), Data$dList$dTypes.Nil);
})(v.asList);
const checkFractionalUse = v => {
  const $0 = Data$dList.span(x => Data$dNumber.floor(x._2) === x._2)(v.asList);
  if ($0.rest.tag === "Cons" && Data$dList$dTypes.foldableList.foldMap(monoidAdditive)(x => Data$dNumber.abs(x._2))($0.rest._2) > 0.0) {
    return Data$dList$dTypes.$List("Cons", $$$Error("InvalidFractionalUse", $0.rest._1._1), Data$dList$dTypes.Nil);
  }
  return Data$dList$dTypes.Nil;
};
const checkEmptiness = v => {
  if (v.asList.tag === "Nil") { return Data$dList$dTypes.$List("Cons", IsEmpty, Data$dList$dTypes.Nil); }
  return Data$dList$dTypes.Nil;
};
const checkValidIsoDuration = v => Data$dFoldable.foldableArray.foldMap(monoidFn)(Data$dFoldable.identity1)([
  checkWeekUsage,
  checkEmptiness,
  checkFractionalUse,
  checkNegativeValues
])({
  asList: (() => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v$1 = go$a0, v1 = go$a1;
        if (v1.tag === "Nil") {
          go$c = false;
          go$r = v$1;
          continue;
        }
        if (v1.tag === "Cons") {
          go$a0 = Data$dList$dTypes.$List("Cons", v1._1, v$1);
          go$a1 = v1._2;
          continue;
        }
        $runtime.fail();
      }
      return go$r;
    };
    return go(Data$dList$dTypes.Nil)((() => {
      const go$1 = go$1$a0$copy => go$1$a1$copy => {
        let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
        while (go$1$c) {
          const source = go$1$a0, memo = go$1$a1;
          const v$1 = Data$dMap$dInternal.stepUnfoldr(source);
          if (v$1.tag === "Nothing") {
            const go$2 = go$2$a0$copy => go$2$a1$copy => {
              let go$2$a0 = go$2$a0$copy, go$2$a1 = go$2$a1$copy, go$2$c = true, go$2$r;
              while (go$2$c) {
                const b = go$2$a0, v$2 = go$2$a1;
                if (v$2.tag === "Nil") {
                  go$2$c = false;
                  go$2$r = b;
                  continue;
                }
                if (v$2.tag === "Cons") {
                  go$2$a0 = Data$dList$dTypes.$List("Cons", v$2._1, b);
                  go$2$a1 = v$2._2;
                  continue;
                }
                $runtime.fail();
              }
              return go$2$r;
            };
            go$1$c = false;
            go$1$r = go$2(Data$dList$dTypes.Nil)(memo);
            continue;
          }
          if (v$1.tag === "Just") {
            go$1$a0 = v$1._1._2;
            go$1$a1 = Data$dList$dTypes.$List("Cons", v$1._1._1, memo);
            continue;
          }
          $runtime.fail();
        }
        return go$1$r;
      };
      return go$1(Data$dMap$dInternal.$MapIter("IterNode", v, Data$dMap$dInternal.IterLeaf))(Data$dList$dTypes.Nil);
    })());
  })(),
  asMap: v
});
const mkIsoDuration = d => {
  const $0 = checkValidIsoDuration(d);
  if ($0.tag === "Nil") { return Data$dEither.$Either("Right", d); }
  if ($0.tag === "Cons") { return Data$dEither.$Either("Left", Data$dNonEmpty.$NonEmpty($0._1, $0._2)); }
  $runtime.fail();
};
export {
  $$$Error,
  ContainsNegativeValue,
  InvalidFractionalUse,
  InvalidWeekComponentUsage,
  IsEmpty,
  checkEmptiness,
  checkFractionalUse,
  checkNegativeValues,
  checkValidIsoDuration,
  checkWeekUsage,
  eqError,
  eqIsoDuration,
  foldMap,
  mkIsoDuration,
  monoidAdditive,
  monoidFn,
  ordError,
  ordIsoDuration,
  prettyError,
  showError,
  showIsoDuration,
  unIsoDuration
};
