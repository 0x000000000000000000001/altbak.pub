import * as $runtime from "../runtime.js";
import * as Control$dMonad$dGen from "../Control.Monad.Gen/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dNonEmpty from "../Data.NonEmpty/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const genTuple = dictApply => {
  const Functor0 = dictApply.Functor0();
  return a => b => dictApply.apply(Functor0.map(Data$dTuple.Tuple)(a))(b);
};
const genNonEmpty = dictMonadRec => dictMonadGen => {
  const Bind1 = dictMonadGen.Monad0().Bind1();
  const Apply0 = Bind1.Apply0();
  const Functor0 = Bind1.Apply0().Functor0();
  return dictUnfoldable => gen => Apply0.apply(Functor0.map(Data$dNonEmpty.NonEmpty)(gen))(dictMonadGen.resize(x => {
    const $0 = x - 1 | 0;
    const v = Data$dOrd.ordInt.compare(0)($0);
    if (v === "LT") { return $0; }
    if (v === "EQ") { return 0; }
    if (v === "GT") { return 0; }
    $runtime.fail();
  })(Control$dMonad$dGen.unfoldable(dictMonadRec)(dictMonadGen)(dictUnfoldable)(gen)));
};
const genMaybe$p = dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const Bind1 = Monad0.Bind1();
  const Functor0 = Monad0.Bind1().Apply0().Functor0();
  const Applicative0 = Monad0.Applicative0();
  return bias => gen => Bind1.bind(dictMonadGen.chooseFloat(0.0)(1.0))(n => {
    if (n < bias) { return Functor0.map(Data$dMaybe.Just)(gen); }
    return Applicative0.pure(Data$dMaybe.Nothing);
  });
};
const genMaybe = dictMonadGen => genMaybe$p(dictMonadGen)(0.75);
const genIdentity = dictFunctor => dictFunctor.map(Data$dIdentity.Identity);
const genEither$p = dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const Bind1 = Monad0.Bind1();
  const Functor0 = Monad0.Bind1().Apply0().Functor0();
  return bias => genA => genB => Bind1.bind(dictMonadGen.chooseFloat(0.0)(1.0))(n => {
    if (n < bias) { return Functor0.map(Data$dEither.Left)(genA); }
    return Functor0.map(Data$dEither.Right)(genB);
  });
};
const genEither = dictMonadGen => genEither$p(dictMonadGen)(0.5);
export {genEither, genEither$p, genIdentity, genMaybe, genMaybe$p, genNonEmpty, genTuple};
