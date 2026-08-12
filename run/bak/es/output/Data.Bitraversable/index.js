import * as $runtime from "../runtime.js";
import * as Data$dBifoldable from "../Data.Bifoldable/index.js";
import * as Data$dBifunctor from "../Data.Bifunctor/index.js";
import * as Data$dConst from "../Data.Const/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFunctor$dClown from "../Data.Functor.Clown/index.js";
import * as Data$dFunctor$dFlip from "../Data.Functor.Flip/index.js";
import * as Data$dFunctor$dJoker from "../Data.Functor.Joker/index.js";
import * as Data$dFunctor$dProduct2 from "../Data.Functor.Product2/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const identity = x => x;
const identity1 = x => x;
const bitraverse = dict => dict.bitraverse;
const lfor = dictBitraversable => dictApplicative => {
  const pure = dictApplicative.pure;
  return t => f => dictBitraversable.bitraverse(dictApplicative)(f)(pure)(t);
};
const ltraverse = dictBitraversable => dictApplicative => {
  const pure = dictApplicative.pure;
  return f => dictBitraversable.bitraverse(dictApplicative)(f)(pure);
};
const rfor = dictBitraversable => dictApplicative => {
  const pure = dictApplicative.pure;
  return t => f => dictBitraversable.bitraverse(dictApplicative)(pure)(f)(t);
};
const rtraverse = dictBitraversable => dictApplicative => dictBitraversable.bitraverse(dictApplicative)(dictApplicative.pure);
const bitraversableTuple = {
  bitraverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => g => v => Apply0.apply(Functor0.map(Data$dTuple.Tuple)(f(v._1)))(g(v._2));
  },
  bisequence: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Apply0.apply(Functor0.map(Data$dTuple.Tuple)(v._1))(v._2);
  },
  Bifunctor0: () => Data$dBifunctor.bifunctorTuple,
  Bifoldable1: () => Data$dBifoldable.bifoldableTuple
};
const bitraversableJoker = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const bifunctorJoker = {bimap: v => g => v1 => $0.map(g)(v1)};
  const $1 = dictTraversable.Foldable1();
  const bifoldableJoker = {
    bifoldr: v => r => u => v1 => $1.foldr(r)(u)(v1),
    bifoldl: v => r => u => v1 => $1.foldl(r)(u)(v1),
    bifoldMap: dictMonoid => v => r => v1 => $1.foldMap(dictMonoid)(r)(v1)
  };
  return {
    bitraverse: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => r => v1 => Functor0.map(Data$dFunctor$dJoker.Joker)(dictTraversable.traverse(dictApplicative)(r)(v1));
    },
    bisequence: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => Functor0.map(Data$dFunctor$dJoker.Joker)(dictTraversable.sequence(dictApplicative)(v));
    },
    Bifunctor0: () => bifunctorJoker,
    Bifoldable1: () => bifoldableJoker
  };
};
const bitraversableEither = {
  bitraverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => v1 => v2 => {
      if (v2.tag === "Left") { return Functor0.map(Data$dEither.Left)(v(v2._1)); }
      if (v2.tag === "Right") { return Functor0.map(Data$dEither.Right)(v1(v2._1)); }
      $runtime.fail();
    };
  },
  bisequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => {
      if (v.tag === "Left") { return Functor0.map(Data$dEither.Left)(v._1); }
      if (v.tag === "Right") { return Functor0.map(Data$dEither.Right)(v._1); }
      $runtime.fail();
    };
  },
  Bifunctor0: () => Data$dBifunctor.bifunctorEither,
  Bifoldable1: () => Data$dBifoldable.bifoldableEither
};
const bitraversableConst = {
  bitraverse: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return f => v => v1 => Functor0.map(Data$dConst.Const)(f(v1));
  },
  bisequence: dictApplicative => {
    const Functor0 = dictApplicative.Apply0().Functor0();
    return v => Functor0.map(Data$dConst.Const)(v);
  },
  Bifunctor0: () => Data$dBifunctor.bifunctorConst,
  Bifoldable1: () => Data$dBifoldable.bifoldableConst
};
const bitraversableClown = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const bifunctorClown = {bimap: f => v => v1 => $0.map(f)(v1)};
  const $1 = dictTraversable.Foldable1();
  const bifoldableClown = {
    bifoldr: l => v => u => v1 => $1.foldr(l)(u)(v1),
    bifoldl: l => v => u => v1 => $1.foldl(l)(u)(v1),
    bifoldMap: dictMonoid => l => v => v1 => $1.foldMap(dictMonoid)(l)(v1)
  };
  return {
    bitraverse: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return l => v => v1 => Functor0.map(Data$dFunctor$dClown.Clown)(dictTraversable.traverse(dictApplicative)(l)(v1));
    },
    bisequence: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => Functor0.map(Data$dFunctor$dClown.Clown)(dictTraversable.sequence(dictApplicative)(v));
    },
    Bifunctor0: () => bifunctorClown,
    Bifoldable1: () => bifoldableClown
  };
};
const bisequenceDefault = dictBitraversable => dictApplicative => dictBitraversable.bitraverse(dictApplicative)(identity)(identity1);
const bisequence = dict => dict.bisequence;
const bitraversableFlip = dictBitraversable => {
  const $0 = dictBitraversable.Bifunctor0();
  const bifunctorFlip = {bimap: f => g => v => $0.bimap(g)(f)(v)};
  const $1 = dictBitraversable.Bifoldable1();
  const bifoldableFlip = {
    bifoldr: r => l => u => v => $1.bifoldr(l)(r)(u)(v),
    bifoldl: r => l => u => v => $1.bifoldl(l)(r)(u)(v),
    bifoldMap: dictMonoid => r => l => v => $1.bifoldMap(dictMonoid)(l)(r)(v)
  };
  return {
    bitraverse: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return r => l => v => Functor0.map(Data$dFunctor$dFlip.Flip)(dictBitraversable.bitraverse(dictApplicative)(l)(r)(v));
    },
    bisequence: dictApplicative => {
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => Functor0.map(Data$dFunctor$dFlip.Flip)(dictBitraversable.bisequence(dictApplicative)(v));
    },
    Bifunctor0: () => bifunctorFlip,
    Bifoldable1: () => bifoldableFlip
  };
};
const bitraversableProduct2 = dictBitraversable => {
  const $0 = dictBitraversable.Bifunctor0();
  const bifoldableProduct2 = Data$dBifoldable.bifoldableProduct2(dictBitraversable.Bifoldable1());
  return dictBitraversable1 => {
    const $1 = dictBitraversable1.Bifunctor0();
    const bifunctorProduct21 = {bimap: f => g => v => Data$dFunctor$dProduct2.$Product2($0.bimap(f)(g)(v._1), $1.bimap(f)(g)(v._2))};
    const bifoldableProduct21 = bifoldableProduct2(dictBitraversable1.Bifoldable1());
    return {
      bitraverse: dictApplicative => {
        const Apply0 = dictApplicative.Apply0();
        const Functor0 = dictApplicative.Apply0().Functor0();
        return l => r => v => Apply0.apply(Functor0.map(Data$dFunctor$dProduct2.Product2)(dictBitraversable.bitraverse(dictApplicative)(l)(r)(v._1)))(dictBitraversable1.bitraverse(dictApplicative)(l)(r)(v._2));
      },
      bisequence: dictApplicative => {
        const Apply0 = dictApplicative.Apply0();
        const Functor0 = dictApplicative.Apply0().Functor0();
        return v => Apply0.apply(Functor0.map(Data$dFunctor$dProduct2.Product2)(dictBitraversable.bisequence(dictApplicative)(v._1)))(dictBitraversable1.bisequence(dictApplicative)(v._2));
      },
      Bifunctor0: () => bifunctorProduct21,
      Bifoldable1: () => bifoldableProduct21
    };
  };
};
const bitraverseDefault = dictBitraversable => {
  const Bifunctor0 = dictBitraversable.Bifunctor0();
  return dictApplicative => f => g => t => dictBitraversable.bisequence(dictApplicative)(Bifunctor0.bimap(f)(g)(t));
};
const bifor = dictBitraversable => dictApplicative => t => f => g => dictBitraversable.bitraverse(dictApplicative)(f)(g)(t);
export {
  bifor,
  bisequence,
  bisequenceDefault,
  bitraversableClown,
  bitraversableConst,
  bitraversableEither,
  bitraversableFlip,
  bitraversableJoker,
  bitraversableProduct2,
  bitraversableTuple,
  bitraverse,
  bitraverseDefault,
  identity,
  identity1,
  lfor,
  ltraverse,
  rfor,
  rtraverse
};
