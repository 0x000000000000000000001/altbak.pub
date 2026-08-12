import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const monoidDual = /* #__PURE__ */ (() => {
  const semigroupDual1 = {append: v => v1 => x => v1(v(x))};
  return {mempty: x => x, Semigroup0: () => semigroupDual1};
})();
const monoidEndo = /* #__PURE__ */ (() => {
  const semigroupEndo1 = {append: v => v1 => x => v(v1(x))};
  return {mempty: x => x, Semigroup0: () => semigroupEndo1};
})();
const monoidEndo1 = /* #__PURE__ */ (() => {
  const semigroupEndo1 = {append: v => v1 => x => v(v1(x))};
  return {mempty: x => x, Semigroup0: () => semigroupEndo1};
})();
const foldrWithIndex = dict => dict.foldrWithIndex;
const traverseWithIndex_ = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  return dictFoldableWithIndex => f => dictFoldableWithIndex.foldrWithIndex(i => {
    const $1 = f(i);
    return x => {
      const $2 = $1(x);
      return b => $0.apply(Functor0.map(v => x$1 => x$1)($2))(b);
    };
  })(dictApplicative.pure());
};
const forWithIndex_ = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  return dictFoldableWithIndex => b => a => dictFoldableWithIndex.foldrWithIndex(i => {
    const $1 = a(i);
    return x => {
      const $2 = $1(x);
      return b$1 => $0.apply(Functor0.map(v => x$1 => x$1)($2))(b$1);
    };
  })(dictApplicative.pure())(b);
};
const foldrDefault = dictFoldableWithIndex => f => dictFoldableWithIndex.foldrWithIndex(v => f);
const foldlWithIndex = dict => dict.foldlWithIndex;
const foldlDefault = dictFoldableWithIndex => f => dictFoldableWithIndex.foldlWithIndex(v => f);
const foldableWithIndexTuple = {
  foldrWithIndex: f => z => v => f()(v._2)(z),
  foldlWithIndex: f => z => v => f()(z)(v._2),
  foldMapWithIndex: dictMonoid => f => v => f()(v._2),
  Foldable0: () => Data$dFoldable.foldableTuple
};
const foldableWithIndexMultiplicative = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => $0(v)(z);
  },
  foldlWithIndex: f => f(),
  foldMapWithIndex: dictMonoid => f => f(),
  Foldable0: () => Data$dFoldable.foldableMultiplicative
};
const foldableWithIndexMaybe = {
  foldrWithIndex: f => {
    const $0 = f();
    return v1 => v2 => {
      if (v2.tag === "Nothing") { return v1; }
      if (v2.tag === "Just") { return $0(v2._1)(v1); }
      $runtime.fail();
    };
  },
  foldlWithIndex: f => {
    const $0 = f();
    return v1 => v2 => {
      if (v2.tag === "Nothing") { return v1; }
      if (v2.tag === "Just") { return $0(v1)(v2._1); }
      $runtime.fail();
    };
  },
  foldMapWithIndex: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return f => {
      const $0 = f();
      return v1 => {
        if (v1.tag === "Nothing") { return mempty; }
        if (v1.tag === "Just") { return $0(v1._1); }
        $runtime.fail();
      };
    };
  },
  Foldable0: () => Data$dFoldable.foldableMaybe
};
const foldableWithIndexLast = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => {
      if (v.tag === "Nothing") { return z; }
      if (v.tag === "Just") { return $0(v._1)(z); }
      $runtime.fail();
    };
  },
  foldlWithIndex: f => {
    const $0 = f();
    return z => v => {
      if (v.tag === "Nothing") { return z; }
      if (v.tag === "Just") { return $0(z)(v._1); }
      $runtime.fail();
    };
  },
  foldMapWithIndex: dictMonoid => f => {
    const $0 = f();
    return v => {
      if (v.tag === "Nothing") { return dictMonoid.mempty; }
      if (v.tag === "Just") { return $0(v._1); }
      $runtime.fail();
    };
  },
  Foldable0: () => Data$dFoldable.foldableLast
};
const foldableWithIndexIdentity = {
  foldrWithIndex: f => z => v => f()(v)(z),
  foldlWithIndex: f => z => v => f()(z)(v),
  foldMapWithIndex: dictMonoid => f => v => f()(v),
  Foldable0: () => Data$dFoldable.foldableIdentity
};
const foldableWithIndexFirst = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => {
      if (v.tag === "Nothing") { return z; }
      if (v.tag === "Just") { return $0(v._1)(z); }
      $runtime.fail();
    };
  },
  foldlWithIndex: f => {
    const $0 = f();
    return z => v => {
      if (v.tag === "Nothing") { return z; }
      if (v.tag === "Just") { return $0(z)(v._1); }
      $runtime.fail();
    };
  },
  foldMapWithIndex: dictMonoid => f => {
    const $0 = f();
    return v => {
      if (v.tag === "Nothing") { return dictMonoid.mempty; }
      if (v.tag === "Just") { return $0(v._1); }
      $runtime.fail();
    };
  },
  Foldable0: () => Data$dFoldable.foldableFirst
};
const foldableWithIndexEither = {
  foldrWithIndex: v => v1 => v2 => {
    if (v2.tag === "Left") { return v1; }
    if (v2.tag === "Right") { return v()(v2._1)(v1); }
    $runtime.fail();
  },
  foldlWithIndex: v => v1 => v2 => {
    if (v2.tag === "Left") { return v1; }
    if (v2.tag === "Right") { return v()(v1)(v2._1); }
    $runtime.fail();
  },
  foldMapWithIndex: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return v => v1 => {
      if (v1.tag === "Left") { return mempty; }
      if (v1.tag === "Right") { return v()(v1._1); }
      $runtime.fail();
    };
  },
  Foldable0: () => Data$dFoldable.foldableEither
};
const foldableWithIndexDual = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => $0(v)(z);
  },
  foldlWithIndex: f => f(),
  foldMapWithIndex: dictMonoid => f => f(),
  Foldable0: () => Data$dFoldable.foldableDual
};
const foldableWithIndexDisj = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => $0(v)(z);
  },
  foldlWithIndex: f => f(),
  foldMapWithIndex: dictMonoid => f => f(),
  Foldable0: () => Data$dFoldable.foldableDisj
};
const foldableWithIndexConst = {
  foldrWithIndex: v => z => v1 => z,
  foldlWithIndex: v => z => v1 => z,
  foldMapWithIndex: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return v => v1 => mempty;
  },
  Foldable0: () => Data$dFoldable.foldableConst
};
const foldableWithIndexConj = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => $0(v)(z);
  },
  foldlWithIndex: f => f(),
  foldMapWithIndex: dictMonoid => f => f(),
  Foldable0: () => Data$dFoldable.foldableConj
};
const foldableWithIndexAdditive = {
  foldrWithIndex: f => {
    const $0 = f();
    return z => v => $0(v)(z);
  },
  foldlWithIndex: f => f(),
  foldMapWithIndex: dictMonoid => f => f(),
  Foldable0: () => Data$dFoldable.foldableAdditive
};
const foldWithIndexM = dictFoldableWithIndex => dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return f => a0 => dictFoldableWithIndex.foldlWithIndex(i => ma => b => Bind1.bind(ma)((() => {
    const $0 = f(i);
    return a => $0(a)(b);
  })()))(Applicative0.pure(a0));
};
const foldMapWithIndexDefaultR = dictFoldableWithIndex => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => dictFoldableWithIndex.foldrWithIndex(i => x => acc => Semigroup0.append(f(i)(x))(acc))(mempty);
};
const foldableWithIndexArray = {
  foldrWithIndex: f => z => {
    const $0 = Data$dFoldable.foldrArray(v => {
      const $0 = v._1;
      const $1 = v._2;
      return y => f($0)($1)(y);
    })(z);
    const $1 = Data$dFunctorWithIndex.mapWithIndexArray(Data$dTuple.Tuple);
    return x => $0($1(x));
  },
  foldlWithIndex: f => z => {
    const $0 = Data$dFoldable.foldlArray(y => v => f(v._1)(y)(v._2))(z);
    const $1 = Data$dFunctorWithIndex.mapWithIndexArray(Data$dTuple.Tuple);
    return x => $0($1(x));
  },
  foldMapWithIndex: dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    const mempty = dictMonoid.mempty;
    return f => foldableWithIndexArray.foldrWithIndex(i => x => acc => Semigroup0.append(f(i)(x))(acc))(mempty);
  },
  Foldable0: () => Data$dFoldable.foldableArray
};
const foldMapWithIndexDefaultL = dictFoldableWithIndex => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => dictFoldableWithIndex.foldlWithIndex(i => acc => x => Semigroup0.append(acc)(f(i)(x)))(mempty);
};
const foldMapWithIndex = dict => dict.foldMapWithIndex;
const foldableWithIndexApp = dictFoldableWithIndex => {
  const $0 = dictFoldableWithIndex.Foldable0();
  const foldableApp = {foldr: f => i => v => $0.foldr(f)(i)(v), foldl: f => i => v => $0.foldl(f)(i)(v), foldMap: dictMonoid => f => v => $0.foldMap(dictMonoid)(f)(v)};
  return {
    foldrWithIndex: f => z => v => dictFoldableWithIndex.foldrWithIndex(f)(z)(v),
    foldlWithIndex: f => z => v => dictFoldableWithIndex.foldlWithIndex(f)(z)(v),
    foldMapWithIndex: dictMonoid => f => v => dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(f)(v),
    Foldable0: () => foldableApp
  };
};
const foldableWithIndexCompose = dictFoldableWithIndex => {
  const $0 = dictFoldableWithIndex.Foldable0();
  return dictFoldableWithIndex1 => {
    const $1 = dictFoldableWithIndex1.Foldable0();
    const foldableCompose1 = {
      foldr: f => i => v => $0.foldr((() => {
        const $2 = $1.foldr(f);
        return b => a => $2(a)(b);
      })())(i)(v),
      foldl: f => i => v => $0.foldl($1.foldl(f))(i)(v),
      foldMap: dictMonoid => f => v => $0.foldMap(dictMonoid)($1.foldMap(dictMonoid)(f))(v)
    };
    return {
      foldrWithIndex: f => i => v => dictFoldableWithIndex.foldrWithIndex(a => {
        const $2 = dictFoldableWithIndex1.foldrWithIndex(b => f(Data$dTuple.$Tuple(a, b)));
        return b => a$1 => $2(a$1)(b);
      })(i)(v),
      foldlWithIndex: f => i => v => dictFoldableWithIndex.foldlWithIndex(x => dictFoldableWithIndex1.foldlWithIndex(b => f(Data$dTuple.$Tuple(x, b))))(i)(v),
      foldMapWithIndex: dictMonoid => {
        const foldMapWithIndex2 = dictFoldableWithIndex1.foldMapWithIndex(dictMonoid);
        return f => v => dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(x => foldMapWithIndex2(b => f(Data$dTuple.$Tuple(x, b))))(v);
      },
      Foldable0: () => foldableCompose1
    };
  };
};
const foldableWithIndexCoproduct = dictFoldableWithIndex => {
  const foldableCoproduct = Data$dFoldable.foldableCoproduct(dictFoldableWithIndex.Foldable0());
  return dictFoldableWithIndex1 => {
    const foldableCoproduct1 = foldableCoproduct(dictFoldableWithIndex1.Foldable0());
    return {
      foldrWithIndex: f => z => {
        const $0 = dictFoldableWithIndex.foldrWithIndex(x => f(Data$dEither.$Either("Left", x)))(z);
        const $1 = dictFoldableWithIndex1.foldrWithIndex(x => f(Data$dEither.$Either("Right", x)))(z);
        return v2 => {
          if (v2.tag === "Left") { return $0(v2._1); }
          if (v2.tag === "Right") { return $1(v2._1); }
          $runtime.fail();
        };
      },
      foldlWithIndex: f => z => {
        const $0 = dictFoldableWithIndex.foldlWithIndex(x => f(Data$dEither.$Either("Left", x)))(z);
        const $1 = dictFoldableWithIndex1.foldlWithIndex(x => f(Data$dEither.$Either("Right", x)))(z);
        return v2 => {
          if (v2.tag === "Left") { return $0(v2._1); }
          if (v2.tag === "Right") { return $1(v2._1); }
          $runtime.fail();
        };
      },
      foldMapWithIndex: dictMonoid => f => {
        const $0 = dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(x => f(Data$dEither.$Either("Left", x)));
        const $1 = dictFoldableWithIndex1.foldMapWithIndex(dictMonoid)(x => f(Data$dEither.$Either("Right", x)));
        return v2 => {
          if (v2.tag === "Left") { return $0(v2._1); }
          if (v2.tag === "Right") { return $1(v2._1); }
          $runtime.fail();
        };
      },
      Foldable0: () => foldableCoproduct1
    };
  };
};
const foldableWithIndexProduct = dictFoldableWithIndex => {
  const $0 = dictFoldableWithIndex.Foldable0();
  return dictFoldableWithIndex1 => {
    const $1 = dictFoldableWithIndex1.Foldable0();
    const foldableProduct1 = {
      foldr: f => z => v => $0.foldr(f)($1.foldr(f)(z)(v._2))(v._1),
      foldl: f => z => v => $1.foldl(f)($0.foldl(f)(z)(v._1))(v._2),
      foldMap: dictMonoid => {
        const Semigroup0 = dictMonoid.Semigroup0();
        return f => v => Semigroup0.append($0.foldMap(dictMonoid)(f)(v._1))($1.foldMap(dictMonoid)(f)(v._2));
      }
    };
    return {
      foldrWithIndex: f => z => v => dictFoldableWithIndex.foldrWithIndex(x => f(Data$dEither.$Either("Left", x)))(dictFoldableWithIndex1.foldrWithIndex(x => f(Data$dEither.$Either(
        "Right",
        x
      )))(z)(v._2))(v._1),
      foldlWithIndex: f => z => v => dictFoldableWithIndex1.foldlWithIndex(x => f(Data$dEither.$Either("Right", x)))(dictFoldableWithIndex.foldlWithIndex(x => f(Data$dEither.$Either(
        "Left",
        x
      )))(z)(v._1))(v._2),
      foldMapWithIndex: dictMonoid => {
        const Semigroup0 = dictMonoid.Semigroup0();
        return f => v => Semigroup0.append(dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(x => f(Data$dEither.$Either("Left", x)))(v._1))(dictFoldableWithIndex1.foldMapWithIndex(dictMonoid)(x => f(Data$dEither.$Either(
          "Right",
          x
        )))(v._2));
      },
      Foldable0: () => foldableProduct1
    };
  };
};
const foldlWithIndexDefault = dictFoldableWithIndex => c => u => xs => dictFoldableWithIndex.foldMapWithIndex(monoidDual)(i => {
  const $0 = c(i);
  return x => a => $0(a)(x);
})(xs)(u);
const foldrWithIndexDefault = dictFoldableWithIndex => c => u => xs => dictFoldableWithIndex.foldMapWithIndex(monoidEndo)(i => c(i))(xs)(u);
const surroundMapWithIndex = dictFoldableWithIndex => dictSemigroup => d => t => f => dictFoldableWithIndex.foldMapWithIndex(monoidEndo1)(i => a => m => dictSemigroup.append(d)(dictSemigroup.append(t(i)(a))(m)))(f)(d);
const foldMapDefault = dictFoldableWithIndex => dictMonoid => f => dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(v => f);
const findWithIndex = dictFoldableWithIndex => p => dictFoldableWithIndex.foldlWithIndex(v => v1 => v2 => {
  if (v1.tag === "Nothing" && p(v)(v2)) { return Data$dMaybe.$Maybe("Just", {index: v, value: v2}); }
  return v1;
})(Data$dMaybe.Nothing);
const findMapWithIndex = dictFoldableWithIndex => f => dictFoldableWithIndex.foldlWithIndex(v => v1 => v2 => {
  if (v1.tag === "Nothing") { return f(v)(v2); }
  return v1;
})(Data$dMaybe.Nothing);
const anyWithIndex = dictFoldableWithIndex => dictHeytingAlgebra => {
  const semigroupDisj1 = {append: v => v1 => dictHeytingAlgebra.disj(v)(v1)};
  return t => dictFoldableWithIndex.foldMapWithIndex({mempty: dictHeytingAlgebra.ff, Semigroup0: () => semigroupDisj1})(i => t(i));
};
const allWithIndex = dictFoldableWithIndex => dictHeytingAlgebra => {
  const semigroupConj1 = {append: v => v1 => dictHeytingAlgebra.conj(v)(v1)};
  return t => dictFoldableWithIndex.foldMapWithIndex({mempty: dictHeytingAlgebra.tt, Semigroup0: () => semigroupConj1})(i => t(i));
};
export {
  allWithIndex,
  anyWithIndex,
  findMapWithIndex,
  findWithIndex,
  foldMapDefault,
  foldMapWithIndex,
  foldMapWithIndexDefaultL,
  foldMapWithIndexDefaultR,
  foldWithIndexM,
  foldableWithIndexAdditive,
  foldableWithIndexApp,
  foldableWithIndexArray,
  foldableWithIndexCompose,
  foldableWithIndexConj,
  foldableWithIndexConst,
  foldableWithIndexCoproduct,
  foldableWithIndexDisj,
  foldableWithIndexDual,
  foldableWithIndexEither,
  foldableWithIndexFirst,
  foldableWithIndexIdentity,
  foldableWithIndexLast,
  foldableWithIndexMaybe,
  foldableWithIndexMultiplicative,
  foldableWithIndexProduct,
  foldableWithIndexTuple,
  foldlDefault,
  foldlWithIndex,
  foldlWithIndexDefault,
  foldrDefault,
  foldrWithIndex,
  foldrWithIndexDefault,
  forWithIndex_,
  monoidDual,
  monoidEndo,
  monoidEndo1,
  surroundMapWithIndex,
  traverseWithIndex_
};
