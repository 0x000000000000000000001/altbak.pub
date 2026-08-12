import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data$dFunctor$dApp from "../Data.Functor.App/index.js";
import * as Data$dFunctor$dCompose from "../Data.Functor.Compose/index.js";
import * as Data$dFunctor$dProduct from "../Data.Functor.Product/index.js";
import * as Data$dFunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTraversable$dAccum$dInternal from "../Data.Traversable.Accum.Internal/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const traverseWithIndexDefault = dictTraversableWithIndex => {
  const FunctorWithIndex0 = dictTraversableWithIndex.FunctorWithIndex0();
  return dictApplicative => {
    const sequence1 = dictTraversableWithIndex.Traversable2().sequence(dictApplicative);
    return f => {
      const $0 = FunctorWithIndex0.mapWithIndex(f);
      return x => sequence1($0(x));
    };
  };
};
const traverseWithIndex = dict => dict.traverseWithIndex;
const traverseDefault = dictTraversableWithIndex => dictApplicative => f => dictTraversableWithIndex.traverseWithIndex(dictApplicative)(v => f);
const traversableWithIndexTuple = {
  traverseWithIndex: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dTuple.Tuple(v._1))(f()(v._2));
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexTuple,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexTuple,
  Traversable2: () => Data$dTraversable.traversableTuple
};
const traversableWithIndexProduct = dictTraversableWithIndex => {
  const $0 = dictTraversableWithIndex.FunctorWithIndex0();
  const $1 = $0.Functor0();
  const foldableWithIndexProduct = Data$dFoldableWithIndex.foldableWithIndexProduct(dictTraversableWithIndex.FoldableWithIndex1());
  const traversableProduct = Data$dTraversable.traversableProduct(dictTraversableWithIndex.Traversable2());
  return dictTraversableWithIndex1 => {
    const $2 = dictTraversableWithIndex1.FunctorWithIndex0();
    const $3 = $2.Functor0();
    const functorWithIndexProduct1 = (() => {
      const functorProduct1 = {map: f => v => Data$dTuple.$Tuple($1.map(f)(v._1), $3.map(f)(v._2))};
      return {
        mapWithIndex: f => v => Data$dTuple.$Tuple($0.mapWithIndex(x => f(Data$dEither.$Either("Left", x)))(v._1), $2.mapWithIndex(x => f(Data$dEither.$Either("Right", x)))(v._2)),
        Functor0: () => functorProduct1
      };
    })();
    const foldableWithIndexProduct1 = foldableWithIndexProduct(dictTraversableWithIndex1.FoldableWithIndex1());
    const traversableProduct1 = traversableProduct(dictTraversableWithIndex1.Traversable2());
    return {
      traverseWithIndex: dictApplicative => {
        const Apply0 = dictApplicative.Apply0();
        return f => v => Apply0.apply(Apply0.Functor0().map(Data$dFunctor$dProduct.product)(dictTraversableWithIndex.traverseWithIndex(dictApplicative)(x => f(Data$dEither.$Either(
          "Left",
          x
        )))(v._1)))(dictTraversableWithIndex1.traverseWithIndex(dictApplicative)(x => f(Data$dEither.$Either("Right", x)))(v._2));
      },
      FunctorWithIndex0: () => functorWithIndexProduct1,
      FoldableWithIndex1: () => foldableWithIndexProduct1,
      Traversable2: () => traversableProduct1
    };
  };
};
const traversableWithIndexMultiplicative = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableMultiplicative.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexMultiplicative,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexMultiplicative,
  Traversable2: () => Data$dTraversable.traversableMultiplicative
};
const traversableWithIndexMaybe = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableMaybe.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexMaybe,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexMaybe,
  Traversable2: () => Data$dTraversable.traversableMaybe
};
const traversableWithIndexLast = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableLast.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexLast,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexLast,
  Traversable2: () => Data$dTraversable.traversableLast
};
const traversableWithIndexIdentity = {
  traverseWithIndex: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => Functor0.map(Data$dIdentity.Identity)(f()(v));
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexIdentity,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexIdentity,
  Traversable2: () => Data$dTraversable.traversableIdentity
};
const traversableWithIndexFirst = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableFirst.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexFirst,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexFirst,
  Traversable2: () => Data$dTraversable.traversableFirst
};
const traversableWithIndexEither = {
  traverseWithIndex: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => v1 => {
      if (v1.tag === "Left") { return dictApplicative.pure(Data$dEither.$Either("Left", v1._1)); }
      if (v1.tag === "Right") { return Functor0.map(Data$dEither.Right)(v()(v1._1)); }
      $runtime.fail();
    };
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexEither,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexEither,
  Traversable2: () => Data$dTraversable.traversableEither
};
const traversableWithIndexDual = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableDual.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexDual,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexDual,
  Traversable2: () => Data$dTraversable.traversableDual
};
const traversableWithIndexDisj = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableDisj.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexDisj,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexDisj,
  Traversable2: () => Data$dTraversable.traversableDisj
};
const traversableWithIndexCoproduct = dictTraversableWithIndex => {
  const functorWithIndexCoproduct = Data$dFunctorWithIndex.functorWithIndexCoproduct(dictTraversableWithIndex.FunctorWithIndex0());
  const foldableWithIndexCoproduct = Data$dFoldableWithIndex.foldableWithIndexCoproduct(dictTraversableWithIndex.FoldableWithIndex1());
  const traversableCoproduct = Data$dTraversable.traversableCoproduct(dictTraversableWithIndex.Traversable2());
  return dictTraversableWithIndex1 => {
    const functorWithIndexCoproduct1 = functorWithIndexCoproduct(dictTraversableWithIndex1.FunctorWithIndex0());
    const foldableWithIndexCoproduct1 = foldableWithIndexCoproduct(dictTraversableWithIndex1.FoldableWithIndex1());
    const traversableCoproduct1 = traversableCoproduct(dictTraversableWithIndex1.Traversable2());
    return {
      traverseWithIndex: dictApplicative => {
        const Functor0 = dictApplicative.Apply0().Functor0();
        return f => {
          const $0 = Functor0.map(x => Data$dEither.$Either("Left", x));
          const $1 = dictTraversableWithIndex.traverseWithIndex(dictApplicative)(x => f(Data$dEither.$Either("Left", x)));
          const $2 = Functor0.map(x => Data$dEither.$Either("Right", x));
          const $3 = dictTraversableWithIndex1.traverseWithIndex(dictApplicative)(x => f(Data$dEither.$Either("Right", x)));
          return v2 => {
            if (v2.tag === "Left") { return $0($1(v2._1)); }
            if (v2.tag === "Right") { return $2($3(v2._1)); }
            $runtime.fail();
          };
        };
      },
      FunctorWithIndex0: () => functorWithIndexCoproduct1,
      FoldableWithIndex1: () => foldableWithIndexCoproduct1,
      Traversable2: () => traversableCoproduct1
    };
  };
};
const traversableWithIndexConst = {
  traverseWithIndex: dictApplicative => v => v1 => dictApplicative.pure(v1),
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexConst,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexConst,
  Traversable2: () => Data$dTraversable.traversableConst
};
const traversableWithIndexConj = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableConj.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexConj,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexConj,
  Traversable2: () => Data$dTraversable.traversableConj
};
const traversableWithIndexCompose = dictTraversableWithIndex => {
  const $0 = dictTraversableWithIndex.FunctorWithIndex0();
  const $1 = $0.Functor0();
  const foldableWithIndexCompose = Data$dFoldableWithIndex.foldableWithIndexCompose(dictTraversableWithIndex.FoldableWithIndex1());
  const traversableCompose = Data$dTraversable.traversableCompose(dictTraversableWithIndex.Traversable2());
  return dictTraversableWithIndex1 => {
    const $2 = dictTraversableWithIndex1.FunctorWithIndex0();
    const $3 = $2.Functor0();
    const functorWithIndexCompose1 = (() => {
      const functorCompose1 = {map: f => v => $1.map($3.map(f))(v)};
      return {mapWithIndex: f => v => $0.mapWithIndex(x => $2.mapWithIndex(b => f(Data$dTuple.$Tuple(x, b))))(v), Functor0: () => functorCompose1};
    })();
    const foldableWithIndexCompose1 = foldableWithIndexCompose(dictTraversableWithIndex1.FoldableWithIndex1());
    const traversableCompose1 = traversableCompose(dictTraversableWithIndex1.Traversable2());
    return {
      traverseWithIndex: dictApplicative => {
        const Functor0 = dictApplicative.Apply0().Functor0();
        const traverseWithIndex2 = dictTraversableWithIndex1.traverseWithIndex(dictApplicative);
        return f => v => Functor0.map(Data$dFunctor$dCompose.Compose)(dictTraversableWithIndex.traverseWithIndex(dictApplicative)(x => traverseWithIndex2(b => f(Data$dTuple.$Tuple(
          x,
          b
        ))))(v));
      },
      FunctorWithIndex0: () => functorWithIndexCompose1,
      FoldableWithIndex1: () => foldableWithIndexCompose1,
      Traversable2: () => traversableCompose1
    };
  };
};
const traversableWithIndexArray = {
  traverseWithIndex: dictApplicative => {
    const FunctorWithIndex0 = traversableWithIndexArray.FunctorWithIndex0();
    const sequence1 = traversableWithIndexArray.Traversable2().sequence(dictApplicative);
    return f => {
      const $0 = FunctorWithIndex0.mapWithIndex(f);
      return x => sequence1($0(x));
    };
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexArray,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexArray,
  Traversable2: () => Data$dTraversable.traversableArray
};
const traversableWithIndexApp = dictTraversableWithIndex => {
  const $0 = dictTraversableWithIndex.FunctorWithIndex0();
  const $1 = $0.Functor0();
  const functorWithIndexApp = {mapWithIndex: f => v => $0.mapWithIndex(f)(v), Functor0: () => $1};
  const $2 = dictTraversableWithIndex.FoldableWithIndex1();
  const $3 = $2.Foldable0();
  const foldableWithIndexApp = (() => {
    const foldableApp = {foldr: f => i => v => $3.foldr(f)(i)(v), foldl: f => i => v => $3.foldl(f)(i)(v), foldMap: dictMonoid => f => v => $3.foldMap(dictMonoid)(f)(v)};
    return {
      foldrWithIndex: f => z => v => $2.foldrWithIndex(f)(z)(v),
      foldlWithIndex: f => z => v => $2.foldlWithIndex(f)(z)(v),
      foldMapWithIndex: dictMonoid => f => v => $2.foldMapWithIndex(dictMonoid)(f)(v),
      Foldable0: () => foldableApp
    };
  })();
  const traversableApp = Data$dTraversable.traversableApp(dictTraversableWithIndex.Traversable2());
  return {
    traverseWithIndex: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return f => v => Functor0.map(Data$dFunctor$dApp.App)(dictTraversableWithIndex.traverseWithIndex(dictApplicative)(f)(v));
    },
    FunctorWithIndex0: () => functorWithIndexApp,
    FoldableWithIndex1: () => foldableWithIndexApp,
    Traversable2: () => traversableApp
  };
};
const traversableWithIndexAdditive = {
  traverseWithIndex: dictApplicative => {
    const traverse8 = Data$dTraversable.traversableAdditive.traverse(dictApplicative);
    return f => traverse8(f());
  },
  FunctorWithIndex0: () => Data$dFunctorWithIndex.functorWithIndexAdditive,
  FoldableWithIndex1: () => Data$dFoldableWithIndex.foldableWithIndexAdditive,
  Traversable2: () => Data$dTraversable.traversableAdditive
};
const mapAccumRWithIndex = dictTraversableWithIndex => f => s0 => xs => dictTraversableWithIndex.traverseWithIndex(Data$dTraversable$dAccum$dInternal.applicativeStateR)(i => a => s => f(i)(s)(a))(xs)(s0);
const scanrWithIndex = dictTraversableWithIndex => f => b0 => xs => dictTraversableWithIndex.traverseWithIndex(Data$dTraversable$dAccum$dInternal.applicativeStateR)(i => a => s => {
  const b$p = f(i)(a)(s);
  return {accum: b$p, value: b$p};
})(xs)(b0).value;
const mapAccumLWithIndex = dictTraversableWithIndex => f => s0 => xs => dictTraversableWithIndex.traverseWithIndex(Data$dTraversable$dAccum$dInternal.applicativeStateL)(i => a => s => f(i)(s)(a))(xs)(s0);
const scanlWithIndex = dictTraversableWithIndex => f => b0 => xs => dictTraversableWithIndex.traverseWithIndex(Data$dTraversable$dAccum$dInternal.applicativeStateL)(i => a => s => {
  const b$p = f(i)(s)(a);
  return {accum: b$p, value: b$p};
})(xs)(b0).value;
const forWithIndex = dictApplicative => dictTraversableWithIndex => {
  const $0 = dictTraversableWithIndex.traverseWithIndex(dictApplicative);
  return b => a => $0(a)(b);
};
export {
  forWithIndex,
  mapAccumLWithIndex,
  mapAccumRWithIndex,
  scanlWithIndex,
  scanrWithIndex,
  traversableWithIndexAdditive,
  traversableWithIndexApp,
  traversableWithIndexArray,
  traversableWithIndexCompose,
  traversableWithIndexConj,
  traversableWithIndexConst,
  traversableWithIndexCoproduct,
  traversableWithIndexDisj,
  traversableWithIndexDual,
  traversableWithIndexEither,
  traversableWithIndexFirst,
  traversableWithIndexIdentity,
  traversableWithIndexLast,
  traversableWithIndexMaybe,
  traversableWithIndexMultiplicative,
  traversableWithIndexProduct,
  traversableWithIndexTuple,
  traverseDefault,
  traverseWithIndex,
  traverseWithIndexDefault
};
