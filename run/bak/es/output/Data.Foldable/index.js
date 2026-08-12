import * as $runtime from "../runtime.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMaybe$dFirst from "../Data.Maybe.First/index.js";
import {foldlArray, foldrArray} from "./foreign.js";
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
const identity1 = x => x;
const identity2 = x => x;
const foldr = dict => dict.foldr;
const indexr = dictFoldable => idx => {
  const $0 = dictFoldable.foldr(a => cursor => {
    if (cursor.elem.tag === "Just") { return cursor; }
    if (cursor.pos === idx) { return {elem: Data$dMaybe.$Maybe("Just", a), pos: cursor.pos}; }
    return {pos: cursor.pos + 1 | 0, elem: cursor.elem};
  })({elem: Data$dMaybe.Nothing, pos: 0});
  return x => $0(x).elem;
};
const $$null = dictFoldable => dictFoldable.foldr(v => v1 => false)(true);
const oneOf = dictFoldable => dictPlus => dictFoldable.foldr(dictPlus.Alt0().alt)(dictPlus.empty);
const oneOfMap = dictFoldable => dictPlus => {
  const empty = dictPlus.empty;
  return f => dictFoldable.foldr(x => dictPlus.Alt0().alt(f(x)))(empty);
};
const traverse_ = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  return dictFoldable => f => dictFoldable.foldr(x => {
    const $1 = f(x);
    return b => $0.apply(Functor0.map(v => x$1 => x$1)($1))(b);
  })(dictApplicative.pure());
};
const for_ = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  return dictFoldable => b => a => dictFoldable.foldr(x => {
    const $1 = a(x);
    return b$1 => $0.apply(Functor0.map(v => x$1 => x$1)($1))(b$1);
  })(dictApplicative.pure())(b);
};
const sequence_ = dictApplicative => dictFoldable => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  return dictFoldable.foldr(x => b => $0.apply(Functor0.map(v => x$1 => x$1)(x))(b))(dictApplicative.pure());
};
const foldl = dict => dict.foldl;
const indexl = dictFoldable => idx => {
  const $0 = dictFoldable.foldl(cursor => a => {
    if (cursor.elem.tag === "Just") { return cursor; }
    if (cursor.pos === idx) { return {elem: Data$dMaybe.$Maybe("Just", a), pos: cursor.pos}; }
    return {pos: cursor.pos + 1 | 0, elem: cursor.elem};
  })({elem: Data$dMaybe.Nothing, pos: 0});
  return x => $0(x).elem;
};
const intercalate = dictFoldable => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return sep => xs => dictFoldable.foldl(v => v1 => {
    if (v.init) { return {init: false, acc: v1}; }
    return {init: false, acc: Semigroup0.append(v.acc)(Semigroup0.append(sep)(v1))};
  })({init: true, acc: mempty})(xs).acc;
};
const length = dictFoldable => dictSemiring => dictFoldable.foldl(c => v => dictSemiring.add(dictSemiring.one)(c))(dictSemiring.zero);
const maximumBy = dictFoldable => cmp => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v1); }
  if (v.tag === "Just") { return Data$dMaybe.$Maybe("Just", cmp(v._1)(v1) === "GT" ? v._1 : v1); }
  $runtime.fail();
})(Data$dMaybe.Nothing);
const maximum = dictOrd => dictFoldable => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v1); }
  if (v.tag === "Just") { return Data$dMaybe.$Maybe("Just", dictOrd.compare(v._1)(v1) === "GT" ? v._1 : v1); }
  $runtime.fail();
})(Data$dMaybe.Nothing);
const minimumBy = dictFoldable => cmp => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v1); }
  if (v.tag === "Just") { return Data$dMaybe.$Maybe("Just", cmp(v._1)(v1) === "LT" ? v._1 : v1); }
  $runtime.fail();
})(Data$dMaybe.Nothing);
const minimum = dictOrd => dictFoldable => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v1); }
  if (v.tag === "Just") { return Data$dMaybe.$Maybe("Just", dictOrd.compare(v._1)(v1) === "LT" ? v._1 : v1); }
  $runtime.fail();
})(Data$dMaybe.Nothing);
const product = dictFoldable => dictSemiring => dictFoldable.foldl(dictSemiring.mul)(dictSemiring.one);
const sum = dictFoldable => dictSemiring => dictFoldable.foldl(dictSemiring.add)(dictSemiring.zero);
const foldableTuple = {foldr: f => z => v => f(v._2)(z), foldl: f => z => v => f(z)(v._2), foldMap: dictMonoid => f => v => f(v._2)};
const foldableMultiplicative = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldableMaybe = {
  foldr: v => v1 => v2 => {
    if (v2.tag === "Nothing") { return v1; }
    if (v2.tag === "Just") { return v(v2._1)(v1); }
    $runtime.fail();
  },
  foldl: v => v1 => v2 => {
    if (v2.tag === "Nothing") { return v1; }
    if (v2.tag === "Just") { return v(v1)(v2._1); }
    $runtime.fail();
  },
  foldMap: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return v => v1 => {
      if (v1.tag === "Nothing") { return mempty; }
      if (v1.tag === "Just") { return v(v1._1); }
      $runtime.fail();
    };
  }
};
const foldableIdentity = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldableEither = {
  foldr: v => v1 => v2 => {
    if (v2.tag === "Left") { return v1; }
    if (v2.tag === "Right") { return v(v2._1)(v1); }
    $runtime.fail();
  },
  foldl: v => v1 => v2 => {
    if (v2.tag === "Left") { return v1; }
    if (v2.tag === "Right") { return v(v1)(v2._1); }
    $runtime.fail();
  },
  foldMap: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return v => v1 => {
      if (v1.tag === "Left") { return mempty; }
      if (v1.tag === "Right") { return v(v1._1); }
      $runtime.fail();
    };
  }
};
const foldableDual = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldableDisj = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldableConst = {
  foldr: v => z => v1 => z,
  foldl: v => z => v1 => z,
  foldMap: dictMonoid => {
    const mempty = dictMonoid.mempty;
    return v => v1 => mempty;
  }
};
const foldableConj = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldableAdditive = {foldr: f => z => v => f(v)(z), foldl: f => z => v => f(z)(v), foldMap: dictMonoid => f => v => f(v)};
const foldMapDefaultR = dictFoldable => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => dictFoldable.foldr(x => acc => Semigroup0.append(f(x))(acc))(mempty);
};
const foldableArray = {
  foldr: foldrArray,
  foldl: foldlArray,
  foldMap: dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    const mempty = dictMonoid.mempty;
    return f => foldableArray.foldr(x => acc => Semigroup0.append(f(x))(acc))(mempty);
  }
};
const foldMapDefaultL = dictFoldable => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => dictFoldable.foldl(acc => x => Semigroup0.append(acc)(f(x)))(mempty);
};
const foldMap = dict => dict.foldMap;
const foldableApp = dictFoldable => (
  {foldr: f => i => v => dictFoldable.foldr(f)(i)(v), foldl: f => i => v => dictFoldable.foldl(f)(i)(v), foldMap: dictMonoid => f => v => dictFoldable.foldMap(dictMonoid)(f)(v)}
);
const foldableCompose = dictFoldable => dictFoldable1 => (
  {
    foldr: f => i => v => dictFoldable.foldr((() => {
      const $0 = dictFoldable1.foldr(f);
      return b => a => $0(a)(b);
    })())(i)(v),
    foldl: f => i => v => dictFoldable.foldl(dictFoldable1.foldl(f))(i)(v),
    foldMap: dictMonoid => f => v => dictFoldable.foldMap(dictMonoid)(dictFoldable1.foldMap(dictMonoid)(f))(v)
  }
);
const foldableCoproduct = dictFoldable => dictFoldable1 => (
  {
    foldr: f => z => {
      const $0 = dictFoldable.foldr(f)(z);
      const $1 = dictFoldable1.foldr(f)(z);
      return v2 => {
        if (v2.tag === "Left") { return $0(v2._1); }
        if (v2.tag === "Right") { return $1(v2._1); }
        $runtime.fail();
      };
    },
    foldl: f => z => {
      const $0 = dictFoldable.foldl(f)(z);
      const $1 = dictFoldable1.foldl(f)(z);
      return v2 => {
        if (v2.tag === "Left") { return $0(v2._1); }
        if (v2.tag === "Right") { return $1(v2._1); }
        $runtime.fail();
      };
    },
    foldMap: dictMonoid => f => {
      const $0 = dictFoldable.foldMap(dictMonoid)(f);
      const $1 = dictFoldable1.foldMap(dictMonoid)(f);
      return v2 => {
        if (v2.tag === "Left") { return $0(v2._1); }
        if (v2.tag === "Right") { return $1(v2._1); }
        $runtime.fail();
      };
    }
  }
);
const foldableFirst = {
  foldr: f => z => v => {
    if (v.tag === "Nothing") { return z; }
    if (v.tag === "Just") { return f(v._1)(z); }
    $runtime.fail();
  },
  foldl: f => z => v => {
    if (v.tag === "Nothing") { return z; }
    if (v.tag === "Just") { return f(z)(v._1); }
    $runtime.fail();
  },
  foldMap: dictMonoid => f => v => {
    if (v.tag === "Nothing") { return dictMonoid.mempty; }
    if (v.tag === "Just") { return f(v._1); }
    $runtime.fail();
  }
};
const foldableLast = {
  foldr: f => z => v => {
    if (v.tag === "Nothing") { return z; }
    if (v.tag === "Just") { return f(v._1)(z); }
    $runtime.fail();
  },
  foldl: f => z => v => {
    if (v.tag === "Nothing") { return z; }
    if (v.tag === "Just") { return f(z)(v._1); }
    $runtime.fail();
  },
  foldMap: dictMonoid => f => v => {
    if (v.tag === "Nothing") { return dictMonoid.mempty; }
    if (v.tag === "Just") { return f(v._1); }
    $runtime.fail();
  }
};
const foldableProduct = dictFoldable => dictFoldable1 => (
  {
    foldr: f => z => v => dictFoldable.foldr(f)(dictFoldable1.foldr(f)(z)(v._2))(v._1),
    foldl: f => z => v => dictFoldable1.foldl(f)(dictFoldable.foldl(f)(z)(v._1))(v._2),
    foldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(dictFoldable.foldMap(dictMonoid)(f)(v._1))(dictFoldable1.foldMap(dictMonoid)(f)(v._2));
    }
  }
);
const foldlDefault = dictFoldable => c => u => xs => dictFoldable.foldMap(monoidDual)(x => a => c(a)(x))(xs)(u);
const foldrDefault = dictFoldable => c => u => xs => dictFoldable.foldMap(monoidEndo)(x => c(x))(xs)(u);
const lookup = dictFoldable => dictEq => a => dictFoldable.foldMap(Data$dMaybe$dFirst.monoidFirst)(v => {
  if (dictEq.eq(a)(v._1)) { return Data$dMaybe.$Maybe("Just", v._2); }
  return Data$dMaybe.Nothing;
});
const surroundMap = dictFoldable => dictSemigroup => d => t => f => dictFoldable.foldMap(monoidEndo1)(a => m => dictSemigroup.append(d)(dictSemigroup.append(t(a))(m)))(f)(d);
const surround = dictFoldable => dictSemigroup => d => surroundMap(dictFoldable)(dictSemigroup)(d)(identity1);
const foldM = dictFoldable => dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return f => b0 => dictFoldable.foldl(b => a => Bind1.bind(b)(a$1 => f(a$1)(a)))(Applicative0.pure(b0));
};
const fold = dictFoldable => dictMonoid => dictFoldable.foldMap(dictMonoid)(identity1);
const findMap = dictFoldable => p => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing") { return p(v1); }
  return v;
})(Data$dMaybe.Nothing);
const find = dictFoldable => p => dictFoldable.foldl(v => v1 => {
  if (v.tag === "Nothing" && p(v1)) { return Data$dMaybe.$Maybe("Just", v1); }
  return v;
})(Data$dMaybe.Nothing);
const any = dictFoldable => dictHeytingAlgebra => dictFoldable.foldMap((() => {
  const semigroupDisj1 = {append: v => v1 => dictHeytingAlgebra.disj(v)(v1)};
  return {mempty: dictHeytingAlgebra.ff, Semigroup0: () => semigroupDisj1};
})());
const elem = dictFoldable => {
  const any1 = dictFoldable.foldMap((() => {
    const semigroupDisj1 = {append: v => v1 => v || v1};
    return {mempty: false, Semigroup0: () => semigroupDisj1};
  })());
  return dictEq => x => any1(dictEq.eq(x));
};
const notElem = dictFoldable => dictEq => x => {
  const $0 = dictFoldable.foldMap((() => {
    const semigroupDisj1 = {append: v => v1 => v || v1};
    return {mempty: false, Semigroup0: () => semigroupDisj1};
  })())(dictEq.eq(x));
  return x$1 => !$0(x$1);
};
const or = dictFoldable => dictHeytingAlgebra => dictFoldable.foldMap((() => {
  const semigroupDisj1 = {append: v => v1 => dictHeytingAlgebra.disj(v)(v1)};
  return {mempty: dictHeytingAlgebra.ff, Semigroup0: () => semigroupDisj1};
})())(identity2);
const all = dictFoldable => dictHeytingAlgebra => dictFoldable.foldMap((() => {
  const semigroupConj1 = {append: v => v1 => dictHeytingAlgebra.conj(v)(v1)};
  return {mempty: dictHeytingAlgebra.tt, Semigroup0: () => semigroupConj1};
})());
const and = dictFoldable => dictHeytingAlgebra => dictFoldable.foldMap((() => {
  const semigroupConj1 = {append: v => v1 => dictHeytingAlgebra.conj(v)(v1)};
  return {mempty: dictHeytingAlgebra.tt, Semigroup0: () => semigroupConj1};
})())(identity2);
export {
  all,
  and,
  any,
  elem,
  find,
  findMap,
  fold,
  foldM,
  foldMap,
  foldMapDefaultL,
  foldMapDefaultR,
  foldableAdditive,
  foldableApp,
  foldableArray,
  foldableCompose,
  foldableConj,
  foldableConst,
  foldableCoproduct,
  foldableDisj,
  foldableDual,
  foldableEither,
  foldableFirst,
  foldableIdentity,
  foldableLast,
  foldableMaybe,
  foldableMultiplicative,
  foldableProduct,
  foldableTuple,
  foldl,
  foldlDefault,
  foldr,
  foldrDefault,
  for_,
  identity1,
  identity2,
  indexl,
  indexr,
  intercalate,
  length,
  lookup,
  maximum,
  maximumBy,
  minimum,
  minimumBy,
  monoidDual,
  monoidEndo,
  monoidEndo1,
  notElem,
  $$null as null,
  oneOf,
  oneOfMap,
  or,
  product,
  sequence_,
  sum,
  surround,
  surroundMap,
  traverse_
};
export * from "./foreign.js";
