import * as $runtime from "../runtime.js";
const identity = x => x;
const identity1 = x => x;
const monoidDual = /* #__PURE__ */ (() => {
  const semigroupDual1 = {append: v => v1 => x => v1(v(x))};
  return {mempty: x => x, Semigroup0: () => semigroupDual1};
})();
const monoidEndo = /* #__PURE__ */ (() => {
  const semigroupEndo1 = {append: v => v1 => x => v(v1(x))};
  return {mempty: x => x, Semigroup0: () => semigroupEndo1};
})();
const identity2 = x => x;
const bifoldr = dict => dict.bifoldr;
const bitraverse_ = dictBifoldable => dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const Functor0$1 = Apply0.Functor0();
  return f => g => dictBifoldable.bifoldr(x => {
    const $0 = f(x);
    return b => Apply0.apply(Functor0.map(v => x$1 => x$1)($0))(b);
  })(x => {
    const $0 = g(x);
    return b => Apply0.apply(Functor0$1.map(v => x$1 => x$1)($0))(b);
  })(dictApplicative.pure());
};
const bifor_ = dictBifoldable => dictApplicative => t => f => g => bitraverse_(dictBifoldable)(dictApplicative)(f)(g)(t);
const bisequence_ = dictBifoldable => dictApplicative => bitraverse_(dictBifoldable)(dictApplicative)(identity)(identity1);
const bifoldl = dict => dict.bifoldl;
const bifoldableTuple = {
  bifoldMap: dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    return f => g => v => Semigroup0.append(f(v._1))(g(v._2));
  },
  bifoldr: f => g => z => v => f(v._1)(g(v._2)(z)),
  bifoldl: f => g => z => v => g(f(z)(v._1))(v._2)
};
const bifoldableJoker = dictFoldable => (
  {
    bifoldr: v => r => u => v1 => dictFoldable.foldr(r)(u)(v1),
    bifoldl: v => r => u => v1 => dictFoldable.foldl(r)(u)(v1),
    bifoldMap: dictMonoid => v => r => v1 => dictFoldable.foldMap(dictMonoid)(r)(v1)
  }
);
const bifoldableEither = {
  bifoldr: v => v1 => v2 => v3 => {
    if (v3.tag === "Left") { return v(v3._1)(v2); }
    if (v3.tag === "Right") { return v1(v3._1)(v2); }
    $runtime.fail();
  },
  bifoldl: v => v1 => v2 => v3 => {
    if (v3.tag === "Left") { return v(v2)(v3._1); }
    if (v3.tag === "Right") { return v1(v2)(v3._1); }
    $runtime.fail();
  },
  bifoldMap: dictMonoid => v => v1 => v2 => {
    if (v2.tag === "Left") { return v(v2._1); }
    if (v2.tag === "Right") { return v1(v2._1); }
    $runtime.fail();
  }
};
const bifoldableConst = {bifoldr: f => v => z => v1 => f(v1)(z), bifoldl: f => v => z => v1 => f(z)(v1), bifoldMap: dictMonoid => f => v => v1 => f(v1)};
const bifoldableClown = dictFoldable => (
  {
    bifoldr: l => v => u => v1 => dictFoldable.foldr(l)(u)(v1),
    bifoldl: l => v => u => v1 => dictFoldable.foldl(l)(u)(v1),
    bifoldMap: dictMonoid => l => v => v1 => dictFoldable.foldMap(dictMonoid)(l)(v1)
  }
);
const bifoldMapDefaultR = dictBifoldable => dictMonoid => {
  const $0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => g => dictBifoldable.bifoldr(x => $0.append(f(x)))(x => $0.append(g(x)))(mempty);
};
const bifoldMapDefaultL = dictBifoldable => dictMonoid => {
  const Semigroup0 = dictMonoid.Semigroup0();
  const mempty = dictMonoid.mempty;
  return f => g => dictBifoldable.bifoldl(m => a => Semigroup0.append(m)(f(a)))(m => b => Semigroup0.append(m)(g(b)))(mempty);
};
const bifoldMap = dict => dict.bifoldMap;
const bifoldableFlip = dictBifoldable => (
  {
    bifoldr: r => l => u => v => dictBifoldable.bifoldr(l)(r)(u)(v),
    bifoldl: r => l => u => v => dictBifoldable.bifoldl(l)(r)(u)(v),
    bifoldMap: dictMonoid => r => l => v => dictBifoldable.bifoldMap(dictMonoid)(l)(r)(v)
  }
);
const bifoldlDefault = dictBifoldable => f => g => z => p => dictBifoldable.bifoldMap(monoidDual)(x => a => f(a)(x))(x => a => g(a)(x))(p)(z);
const bifoldrDefault = dictBifoldable => f => g => z => p => dictBifoldable.bifoldMap(monoidEndo)(x => f(x))(x => g(x))(p)(z);
const bifoldableProduct2 = dictBifoldable => dictBifoldable1 => (
  {
    bifoldr: l => r => u => m => bifoldrDefault(bifoldableProduct2(dictBifoldable)(dictBifoldable1))(l)(r)(u)(m),
    bifoldl: l => r => u => m => bifoldlDefault(bifoldableProduct2(dictBifoldable)(dictBifoldable1))(l)(r)(u)(m),
    bifoldMap: dictMonoid => {
      const Semigroup0 = dictMonoid.Semigroup0();
      return l => r => v => Semigroup0.append(dictBifoldable.bifoldMap(dictMonoid)(l)(r)(v._1))(dictBifoldable1.bifoldMap(dictMonoid)(l)(r)(v._2));
    }
  }
);
const bifold = dictBifoldable => dictMonoid => dictBifoldable.bifoldMap(dictMonoid)(identity2)(identity2);
const biany = dictBifoldable => dictBooleanAlgebra => {
  const $0 = dictBooleanAlgebra.HeytingAlgebra0();
  const semigroupDisj1 = {append: v => v1 => $0.disj(v)(v1)};
  return p => q => dictBifoldable.bifoldMap({mempty: $0.ff, Semigroup0: () => semigroupDisj1})(x => p(x))(x => q(x));
};
const biall = dictBifoldable => dictBooleanAlgebra => {
  const $0 = dictBooleanAlgebra.HeytingAlgebra0();
  const semigroupConj1 = {append: v => v1 => $0.conj(v)(v1)};
  return p => q => dictBifoldable.bifoldMap({mempty: $0.tt, Semigroup0: () => semigroupConj1})(x => p(x))(x => q(x));
};
export {
  biall,
  biany,
  bifold,
  bifoldMap,
  bifoldMapDefaultL,
  bifoldMapDefaultR,
  bifoldableClown,
  bifoldableConst,
  bifoldableEither,
  bifoldableFlip,
  bifoldableJoker,
  bifoldableProduct2,
  bifoldableTuple,
  bifoldl,
  bifoldlDefault,
  bifoldr,
  bifoldrDefault,
  bifor_,
  bisequence_,
  bitraverse_,
  identity,
  identity1,
  identity2,
  monoidDual,
  monoidEndo
};
