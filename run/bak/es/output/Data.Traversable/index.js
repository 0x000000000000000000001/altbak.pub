import * as $runtime from "../runtime.js";
import * as Data$dConst from "../Data.Const/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dFunctor$dApp from "../Data.Functor.App/index.js";
import * as Data$dFunctor$dCompose from "../Data.Functor.Compose/index.js";
import * as Data$dFunctor$dProduct from "../Data.Functor.Product/index.js";
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMaybe$dFirst from "../Data.Maybe.First/index.js";
import * as Data$dMaybe$dLast from "../Data.Maybe.Last/index.js";
import * as Data$dMonoid$dAdditive from "../Data.Monoid.Additive/index.js";
import * as Data$dMonoid$dConj from "../Data.Monoid.Conj/index.js";
import * as Data$dMonoid$dDisj from "../Data.Monoid.Disj/index.js";
import * as Data$dMonoid$dDual from "../Data.Monoid.Dual/index.js";
import * as Data$dMonoid$dMultiplicative from "../Data.Monoid.Multiplicative/index.js";
import * as Data$dTraversable$dAccum$dInternal from "../Data.Traversable.Accum.Internal/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import {traverseArrayImpl} from "./foreign.js";
const identity = x => x;
const traverse = dict => dict.traverse;
const traversableTuple = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dTuple.Tuple(v._1))(f(v._2));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dTuple.Tuple(v._1))(v._2);
  },
  Functor0: () => Data$dTuple.functorTuple,
  Foldable1: () => Data$dFoldable.foldableTuple
};
const traversableMultiplicative = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMonoid$dMultiplicative.Multiplicative)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMonoid$dMultiplicative.Multiplicative)(v);
  },
  Functor0: () => Data$dMonoid$dMultiplicative.functorMultiplicative,
  Foldable1: () => Data$dFoldable.foldableMultiplicative
};
const traversableMaybe = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => v1 => {
      if (v1.tag === "Nothing") { return dictApplicative.pure(Data$dMaybe.Nothing); }
      if (v1.tag === "Just") { return Functor0.map(Data$dMaybe.Just)(v(v1._1)); }
      $runtime.fail();
    };
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => {
      if (v.tag === "Nothing") { return dictApplicative.pure(Data$dMaybe.Nothing); }
      if (v.tag === "Just") { return Functor0.map(Data$dMaybe.Just)(v._1); }
      $runtime.fail();
    };
  },
  Functor0: () => Data$dMaybe.functorMaybe,
  Foldable1: () => Data$dFoldable.foldableMaybe
};
const traversableIdentity = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dIdentity.Identity)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dIdentity.Identity)(v);
  },
  Functor0: () => Data$dIdentity.functorIdentity,
  Foldable1: () => Data$dFoldable.foldableIdentity
};
const traversableEither = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => v1 => {
      if (v1.tag === "Left") { return dictApplicative.pure(Data$dEither.$Either("Left", v1._1)); }
      if (v1.tag === "Right") { return Functor0.map(Data$dEither.Right)(v(v1._1)); }
      $runtime.fail();
    };
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => {
      if (v.tag === "Left") { return dictApplicative.pure(Data$dEither.$Either("Left", v._1)); }
      if (v.tag === "Right") { return Functor0.map(Data$dEither.Right)(v._1); }
      $runtime.fail();
    };
  },
  Functor0: () => Data$dEither.functorEither,
  Foldable1: () => Data$dFoldable.foldableEither
};
const traversableDual = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMonoid$dDual.Dual)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMonoid$dDual.Dual)(v);
  },
  Functor0: () => Data$dMonoid$dDual.functorDual,
  Foldable1: () => Data$dFoldable.foldableDual
};
const traversableDisj = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMonoid$dDisj.Disj)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMonoid$dDisj.Disj)(v);
  },
  Functor0: () => Data$dMonoid$dDisj.functorDisj,
  Foldable1: () => Data$dFoldable.foldableDisj
};
const traversableConst = {
  traverse: dictApplicative => v => v1 => dictApplicative.pure(v1),
  sequence: dictApplicative => v => dictApplicative.pure(v),
  Functor0: () => Data$dConst.functorConst,
  Foldable1: () => Data$dFoldable.foldableConst
};
const traversableConj = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMonoid$dConj.Conj)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMonoid$dConj.Conj)(v);
  },
  Functor0: () => Data$dMonoid$dConj.functorConj,
  Foldable1: () => Data$dFoldable.foldableConj
};
const traversableCompose = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const $1 = dictTraversable.Foldable1();
  return dictTraversable1 => {
    const $2 = dictTraversable1.Functor0();
    const functorCompose1 = {map: f => v => $0.map($2.map(f))(v)};
    const $3 = dictTraversable1.Foldable1();
    const foldableCompose1 = {
      foldr: f => i => v => $1.foldr((() => {
        const $4 = $3.foldr(f);
        return b => a => $4(a)(b);
      })())(i)(v),
      foldl: f => i => v => $1.foldl($3.foldl(f))(i)(v),
      foldMap: dictMonoid => f => v => $1.foldMap(dictMonoid)($3.foldMap(dictMonoid)(f))(v)
    };
    return {
      traverse: dictApplicative => {
        const Functor0 = dictApplicative.Apply0().Functor0();
        return f => v => Functor0.map(Data$dFunctor$dCompose.Compose)(dictTraversable.traverse(dictApplicative)(dictTraversable1.traverse(dictApplicative)(f))(v));
      },
      sequence: dictApplicative => traversableCompose(dictTraversable)(dictTraversable1).traverse(dictApplicative)(identity),
      Functor0: () => functorCompose1,
      Foldable1: () => foldableCompose1
    };
  };
};
const traversableAdditive = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMonoid$dAdditive.Additive)(f(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMonoid$dAdditive.Additive)(v);
  },
  Functor0: () => Data$dMonoid$dAdditive.functorAdditive,
  Foldable1: () => Data$dFoldable.foldableAdditive
};
const sequenceDefault = dictTraversable => dictApplicative => dictTraversable.traverse(dictApplicative)(identity);
const traversableArray = {
  traverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    return traverseArrayImpl(Apply0.apply)(Apply0.Functor0().map)(dictApplicative.pure);
  },
  sequence: dictApplicative => traversableArray.traverse(dictApplicative)(identity),
  Functor0: () => Data$dFunctor.functorArray,
  Foldable1: () => Data$dFoldable.foldableArray
};
const sequence = dict => dict.sequence;
const traversableApp = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const $1 = dictTraversable.Foldable1();
  const foldableApp = {foldr: f => i => v => $1.foldr(f)(i)(v), foldl: f => i => v => $1.foldl(f)(i)(v), foldMap: dictMonoid => f => v => $1.foldMap(dictMonoid)(f)(v)};
  return {
    traverse: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return f => v => Functor0.map(Data$dFunctor$dApp.App)(dictTraversable.traverse(dictApplicative)(f)(v));
    },
    sequence: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => Functor0.map(Data$dFunctor$dApp.App)(dictTraversable.sequence(dictApplicative)(v));
    },
    Functor0: () => $0,
    Foldable1: () => foldableApp
  };
};
const traversableCoproduct = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const foldableCoproduct = Data$dFoldable.foldableCoproduct(dictTraversable.Foldable1());
  return dictTraversable1 => {
    const $1 = dictTraversable1.Functor0();
    const functorCoproduct1 = {
      map: f => v => {
        const $2 = $0.map(f);
        const $3 = $1.map(f);
        if (v.tag === "Left") { return Data$dEither.$Either("Left", $2(v._1)); }
        if (v.tag === "Right") { return Data$dEither.$Either("Right", $3(v._1)); }
        $runtime.fail();
      }
    };
    const foldableCoproduct1 = foldableCoproduct(dictTraversable1.Foldable1());
    return {
      traverse: dictApplicative => {
        const Functor0 = dictApplicative.Apply0().Functor0();
        return f => {
          const $2 = Functor0.map(x => Data$dEither.$Either("Left", x));
          const $3 = dictTraversable.traverse(dictApplicative)(f);
          const $4 = Functor0.map(x => Data$dEither.$Either("Right", x));
          const $5 = dictTraversable1.traverse(dictApplicative)(f);
          return v2 => {
            if (v2.tag === "Left") { return $2($3(v2._1)); }
            if (v2.tag === "Right") { return $4($5(v2._1)); }
            $runtime.fail();
          };
        };
      },
      sequence: dictApplicative => {
        const Functor0 = dictApplicative.Apply0().Functor0();
        const $2 = Functor0.map(x => Data$dEither.$Either("Left", x));
        const $3 = dictTraversable.sequence(dictApplicative);
        const $4 = Functor0.map(x => Data$dEither.$Either("Right", x));
        const $5 = dictTraversable1.sequence(dictApplicative);
        return v2 => {
          if (v2.tag === "Left") { return $2($3(v2._1)); }
          if (v2.tag === "Right") { return $4($5(v2._1)); }
          $runtime.fail();
        };
      },
      Functor0: () => functorCoproduct1,
      Foldable1: () => foldableCoproduct1
    };
  };
};
const traversableFirst = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMaybe$dFirst.First)(traversableMaybe.traverse(dictApplicative)(f)(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMaybe$dFirst.First)(traversableMaybe.sequence(dictApplicative)(v));
  },
  Functor0: () => Data$dMaybe.functorMaybe,
  Foldable1: () => Data$dFoldable.foldableFirst
};
const traversableLast = {
  traverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dMaybe$dLast.Last)(traversableMaybe.traverse(dictApplicative)(f)(v));
  },
  sequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dMaybe$dLast.Last)(traversableMaybe.sequence(dictApplicative)(v));
  },
  Functor0: () => Data$dMaybe.functorMaybe,
  Foldable1: () => Data$dFoldable.foldableLast
};
const traversableProduct = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const $1 = dictTraversable.Foldable1();
  return dictTraversable1 => {
    const $2 = dictTraversable1.Functor0();
    const functorProduct1 = {map: f => v => Data$dTuple.$Tuple($0.map(f)(v._1), $2.map(f)(v._2))};
    const $3 = dictTraversable1.Foldable1();
    const foldableProduct1 = {
      foldr: f => z => v => $1.foldr(f)($3.foldr(f)(z)(v._2))(v._1),
      foldl: f => z => v => $3.foldl(f)($1.foldl(f)(z)(v._1))(v._2),
      foldMap: dictMonoid => {
        const Semigroup0 = dictMonoid.Semigroup0();
        return f => v => Semigroup0.append($1.foldMap(dictMonoid)(f)(v._1))($3.foldMap(dictMonoid)(f)(v._2));
      }
    };
    return {
      traverse: dictApplicative => {
        const Apply0 = dictApplicative.Apply0();
        return f => v => Apply0.apply(Apply0.Functor0().map(Data$dFunctor$dProduct.product)(dictTraversable.traverse(dictApplicative)(f)(v._1)))(dictTraversable1.traverse(dictApplicative)(f)(v._2));
      },
      sequence: dictApplicative => {
        const Apply0 = dictApplicative.Apply0();
        return v => Apply0.apply(Apply0.Functor0().map(Data$dFunctor$dProduct.product)(dictTraversable.sequence(dictApplicative)(v._1)))(dictTraversable1.sequence(dictApplicative)(v._2));
      },
      Functor0: () => functorProduct1,
      Foldable1: () => foldableProduct1
    };
  };
};
const traverseDefault = dictTraversable => {
  const Functor0 = dictTraversable.Functor0();
  return dictApplicative => f => ta => dictTraversable.sequence(dictApplicative)(Functor0.map(f)(ta));
};
const mapAccumR = dictTraversable => f => s0 => xs => dictTraversable.traverse(Data$dTraversable$dAccum$dInternal.applicativeStateR)(a => s => f(s)(a))(xs)(s0);
const scanr = dictTraversable => f => b0 => xs => dictTraversable.traverse(Data$dTraversable$dAccum$dInternal.applicativeStateR)(a => s => {
  const b$p = f(a)(s);
  return {accum: b$p, value: b$p};
})(xs)(b0).value;
const mapAccumL = dictTraversable => f => s0 => xs => dictTraversable.traverse(Data$dTraversable$dAccum$dInternal.applicativeStateL)(a => s => f(s)(a))(xs)(s0);
const scanl = dictTraversable => f => b0 => xs => dictTraversable.traverse(Data$dTraversable$dAccum$dInternal.applicativeStateL)(a => s => {
  const b$p = f(s)(a);
  return {accum: b$p, value: b$p};
})(xs)(b0).value;
const $$for = dictApplicative => dictTraversable => x => f => dictTraversable.traverse(dictApplicative)(f)(x);
export {
  $$for as for,
  identity,
  mapAccumL,
  mapAccumR,
  scanl,
  scanr,
  sequence,
  sequenceDefault,
  traversableAdditive,
  traversableApp,
  traversableArray,
  traversableCompose,
  traversableConj,
  traversableConst,
  traversableCoproduct,
  traversableDisj,
  traversableDual,
  traversableEither,
  traversableFirst,
  traversableIdentity,
  traversableLast,
  traversableMaybe,
  traversableMultiplicative,
  traversableProduct,
  traversableTuple,
  traverse,
  traverseDefault
};
export * from "./foreign.js";
