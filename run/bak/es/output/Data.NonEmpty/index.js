// | This module defines a generic non-empty data structure, which adds an
// | additional element to any container type.
import * as $runtime from "../runtime.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
const $NonEmpty = (_1, _2) => ({tag: "NonEmpty", _1, _2});
const NonEmpty = value0 => value1 => $NonEmpty(value0, value1);
const unfoldable1NonEmpty = dictUnfoldable => (
  {
    unfoldr1: f => b => {
      const $0 = f(b);
      return $NonEmpty(
        $0._1,
        dictUnfoldable.unfoldr(v1 => {
          if (v1.tag === "Just") { return Data$dMaybe.$Maybe("Just", f(v1._1)); }
          return Data$dMaybe.Nothing;
        })($0._2)
      );
    }
  }
);
const tail = v => v._2;
const singleton = dictPlus => {
  const empty = dictPlus.empty;
  return a => $NonEmpty(a, empty);
};
const showNonEmpty = dictShow => dictShow1 => ({show: v => "(NonEmpty " + dictShow.show(v._1) + " " + dictShow1.show(v._2) + ")"});
const semigroupNonEmpty = dictApplicative => dictSemigroup => (
  {append: v => v1 => $NonEmpty(v._1, dictSemigroup.append(v._2)(dictSemigroup.append(dictApplicative.pure(v1._1))(v1._2)))}
);
const oneOf = dictAlternative => {
  const Alt0 = dictAlternative.Plus1().Alt0();
  const Applicative0 = dictAlternative.Applicative0();
  return v => Alt0.alt(Applicative0.pure(v._1))(v._2);
};
const head = v => v._1;
const functorNonEmpty = dictFunctor => ({map: f => m => $NonEmpty(f(m._1), dictFunctor.map(f)(m._2))});
const functorWithIndex = dictFunctorWithIndex => {
  const $0 = dictFunctorWithIndex.Functor0();
  const functorNonEmpty1 = {map: f => m => $NonEmpty(f(m._1), $0.map(f)(m._2))};
  return {
    mapWithIndex: f => v => $NonEmpty(f(Data$dMaybe.Nothing)(v._1), dictFunctorWithIndex.mapWithIndex(x => f(Data$dMaybe.$Maybe("Just", x)))(v._2)),
    Functor0: () => functorNonEmpty1
  };
};
const fromNonEmpty = f => v => f(v._1)(v._2);
const foldableNonEmpty = dictFoldable => (
  {
    foldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(f(v._1))(dictFoldable.foldMap(dictMonoid)(f)(v._2));
    },
    foldl: f => b => v => dictFoldable.foldl(f)(f(b)(v._1))(v._2),
    foldr: f => b => v => f(v._1)(dictFoldable.foldr(f)(b)(v._2))
  }
);
const foldableWithIndexNonEmpty = dictFoldableWithIndex => {
  const $0 = dictFoldableWithIndex.Foldable0();
  const foldableNonEmpty1 = {
    foldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(f(v._1))($0.foldMap(dictMonoid)(f)(v._2));
    },
    foldl: f => b => v => $0.foldl(f)(f(b)(v._1))(v._2),
    foldr: f => b => v => f(v._1)($0.foldr(f)(b)(v._2))
  };
  return {
    foldMapWithIndex: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(f(Data$dMaybe.Nothing)(v._1))(dictFoldableWithIndex.foldMapWithIndex(dictMonoid)(x => f(Data$dMaybe.$Maybe("Just", x)))(v._2));
    },
    foldlWithIndex: f => b => v => dictFoldableWithIndex.foldlWithIndex(x => f(Data$dMaybe.$Maybe("Just", x)))(f(Data$dMaybe.Nothing)(b)(v._1))(v._2),
    foldrWithIndex: f => b => v => f(Data$dMaybe.Nothing)(v._1)(dictFoldableWithIndex.foldrWithIndex(x => f(Data$dMaybe.$Maybe("Just", x)))(b)(v._2)),
    Foldable0: () => foldableNonEmpty1
  };
};
const traversableNonEmpty = dictTraversable => {
  const $0 = dictTraversable.Functor0();
  const functorNonEmpty1 = {map: f => m => $NonEmpty(f(m._1), $0.map(f)(m._2))};
  const $1 = dictTraversable.Foldable1();
  const foldableNonEmpty1 = {
    foldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(f(v._1))($1.foldMap(dictMonoid)(f)(v._2));
    },
    foldl: f => b => v => $1.foldl(f)(f(b)(v._1))(v._2),
    foldr: f => b => v => f(v._1)($1.foldr(f)(b)(v._2))
  };
  return {
    sequence: dictApplicative => {
      const Apply0 = dictApplicative.Apply0();
      const Functor0 = dictApplicative.Apply0().Functor0();
      return v => Apply0.apply(Functor0.map(NonEmpty)(v._1))(dictTraversable.sequence(dictApplicative)(v._2));
    },
    traverse: dictApplicative => {
      const Apply0 = dictApplicative.Apply0();
      const Functor0 = dictApplicative.Apply0().Functor0();
      return f => v => Apply0.apply(Functor0.map(NonEmpty)(f(v._1)))(dictTraversable.traverse(dictApplicative)(f)(v._2));
    },
    Functor0: () => functorNonEmpty1,
    Foldable1: () => foldableNonEmpty1
  };
};
const traversableWithIndexNonEmpty = dictTraversableWithIndex => {
  const $0 = dictTraversableWithIndex.FunctorWithIndex0();
  const $1 = $0.Functor0();
  const functorWithIndex1 = (() => {
    const functorNonEmpty1 = {map: f => m => $NonEmpty(f(m._1), $1.map(f)(m._2))};
    return {mapWithIndex: f => v => $NonEmpty(f(Data$dMaybe.Nothing)(v._1), $0.mapWithIndex(x => f(Data$dMaybe.$Maybe("Just", x)))(v._2)), Functor0: () => functorNonEmpty1};
  })();
  const foldableWithIndexNonEmpty1 = foldableWithIndexNonEmpty(dictTraversableWithIndex.FoldableWithIndex1());
  const traversableNonEmpty1 = traversableNonEmpty(dictTraversableWithIndex.Traversable2());
  return {
    traverseWithIndex: dictApplicative => {
      const Apply0 = dictApplicative.Apply0();
      const Functor0 = dictApplicative.Apply0().Functor0();
      return f => v => Apply0.apply(Functor0.map(NonEmpty)(f(Data$dMaybe.Nothing)(v._1)))(dictTraversableWithIndex.traverseWithIndex(dictApplicative)(x => f(Data$dMaybe.$Maybe(
        "Just",
        x
      )))(v._2));
    },
    FunctorWithIndex0: () => functorWithIndex1,
    FoldableWithIndex1: () => foldableWithIndexNonEmpty1,
    Traversable2: () => traversableNonEmpty1
  };
};
const foldable1NonEmpty = dictFoldable => {
  const foldableNonEmpty1 = {
    foldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return f => v => Semigroup0.append(f(v._1))(dictFoldable.foldMap(dictMonoid)(f)(v._2));
    },
    foldl: f => b => v => dictFoldable.foldl(f)(f(b)(v._1))(v._2),
    foldr: f => b => v => f(v._1)(dictFoldable.foldr(f)(b)(v._2))
  };
  return {
    foldMap1: dictSemigroup => f => v => dictFoldable.foldl(s => a1 => dictSemigroup.append(s)(f(a1)))(f(v._1))(v._2),
    foldr1: f => v => {
      const $0 = f(v._1);
      const $1 = dictFoldable.foldr(a1 => {
        const $1 = f(a1);
        return x => Data$dMaybe.$Maybe(
          "Just",
          (() => {
            if (x.tag === "Nothing") { return a1; }
            if (x.tag === "Just") { return $1(x._1); }
            $runtime.fail();
          })()
        );
      })(Data$dMaybe.Nothing)(v._2);
      if ($1.tag === "Nothing") { return v._1; }
      if ($1.tag === "Just") { return $0($1._1); }
      $runtime.fail();
    },
    foldl1: f => v => dictFoldable.foldl(f)(v._1)(v._2),
    Foldable0: () => foldableNonEmpty1
  };
};
const foldl1 = dictFoldable => foldable1NonEmpty(dictFoldable).foldl1;
const eqNonEmpty = dictEq1 => dictEq => ({eq: x => y => dictEq.eq(x._1)(y._1) && dictEq1.eq1(dictEq)(x._2)(y._2)});
const ordNonEmpty = dictOrd1 => {
  const $0 = dictOrd1.Eq10();
  return dictOrd => {
    const $1 = dictOrd.Eq0();
    const eqNonEmpty2 = {eq: x => y => $1.eq(x._1)(y._1) && $0.eq1($1)(x._2)(y._2)};
    return {
      compare: x => y => {
        const v = dictOrd.compare(x._1)(y._1);
        if (v === "LT") { return Data$dOrdering.LT; }
        if (v === "GT") { return Data$dOrdering.GT; }
        return dictOrd1.compare1(dictOrd)(x._2)(y._2);
      },
      Eq0: () => eqNonEmpty2
    };
  };
};
const eq1NonEmpty = dictEq1 => ({eq1: dictEq => x => y => dictEq.eq(x._1)(y._1) && dictEq1.eq1(dictEq)(x._2)(y._2)});
const ord1NonEmpty = dictOrd1 => {
  const $0 = dictOrd1.Eq10();
  const eq1NonEmpty1 = {eq1: dictEq => x => y => dictEq.eq(x._1)(y._1) && $0.eq1(dictEq)(x._2)(y._2)};
  return {
    compare1: dictOrd => x => y => {
      const v = dictOrd.compare(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      return dictOrd1.compare1(dictOrd)(x._2)(y._2);
    },
    Eq10: () => eq1NonEmpty1
  };
};
export {
  $NonEmpty,
  NonEmpty,
  eq1NonEmpty,
  eqNonEmpty,
  foldable1NonEmpty,
  foldableNonEmpty,
  foldableWithIndexNonEmpty,
  foldl1,
  fromNonEmpty,
  functorNonEmpty,
  functorWithIndex,
  head,
  oneOf,
  ord1NonEmpty,
  ordNonEmpty,
  semigroupNonEmpty,
  showNonEmpty,
  singleton,
  tail,
  traversableNonEmpty,
  traversableWithIndexNonEmpty,
  unfoldable1NonEmpty
};
