const parTraverse_ = dictParallel => dictApplicative => dictFoldable => f => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  const $1 = dictFoldable.foldr(x => {
    const $1 = dictParallel.parallel(f(x));
    return b => $0.apply(Functor0.map(v => x$1 => x$1)($1))(b);
  })(dictApplicative.pure());
  return x => dictParallel.sequential($1(x));
};
const parTraverse = dictParallel => dictApplicative => dictTraversable => f => {
  const $0 = dictTraversable.traverse(dictApplicative)(x => dictParallel.parallel(f(x)));
  return x => dictParallel.sequential($0(x));
};
const parSequence_ = dictParallel => dictApplicative => dictFoldable => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  const $1 = dictFoldable.foldr(x => {
    const $1 = dictParallel.parallel(x);
    return b => $0.apply(Functor0.map(v => x$1 => x$1)($1))(b);
  })(dictApplicative.pure());
  return x => dictParallel.sequential($1(x));
};
const parSequence = dictParallel => dictApplicative => dictTraversable => {
  const $0 = dictTraversable.traverse(dictApplicative)(x => dictParallel.parallel(x));
  return x => dictParallel.sequential($0(x));
};
const parOneOfMap = dictParallel => dictAlternative => {
  const Plus1 = dictAlternative.Plus1();
  return dictFoldable => dictFunctor => f => {
    const $0 = dictFoldable.foldr(x => Plus1.Alt0().alt(dictParallel.parallel(f(x))))(Plus1.empty);
    return x => dictParallel.sequential($0(x));
  };
};
const parOneOf = dictParallel => dictAlternative => {
  const Plus1 = dictAlternative.Plus1();
  return dictFoldable => dictFunctor => {
    const $0 = dictFoldable.foldr(x => Plus1.Alt0().alt(dictParallel.parallel(x)))(Plus1.empty);
    return x => dictParallel.sequential($0(x));
  };
};
const parApply = dictParallel => {
  const Apply1 = dictParallel.Apply1();
  return mf => ma => dictParallel.sequential(Apply1.apply(dictParallel.parallel(mf))(dictParallel.parallel(ma)));
};
export {parApply, parOneOf, parOneOfMap, parSequence, parSequence_, parTraverse, parTraverse_};
