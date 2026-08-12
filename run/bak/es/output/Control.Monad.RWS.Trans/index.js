// | This module defines the reader-writer-state monad transformer, `RWST`.
import * as $runtime from "../runtime.js";
import * as Control$dMonad$dRec$dClass from "../Control.Monad.Rec.Class/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const $RWSResult = (_1, _2, _3) => ({tag: "RWSResult", _1, _2, _3});
const RWSResult = value0 => value1 => value2 => $RWSResult(value0, value1, value2);
const RWST = x => x;
const withRWST = f => m => r => s => {
  const $0 = f(r)(s);
  return m($0._1)($0._2);
};
const runRWST = v => v;
const newtypeRWST = {Coercible0: () => {}};
const monadTransRWST = dictMonoid => (
  {
    lift: dictMonad => {
      const Bind1 = dictMonad.Bind1();
      return m => v => s => Bind1.bind(m)(a => dictMonad.Applicative0().pure($RWSResult(s, a, dictMonoid.mempty)));
    }
  }
);
const mapRWST = f => v => r => s => f(v(r)(s));
const lazyRWST = {defer: f => r => s => f()(r)(s)};
const functorRWST = dictFunctor => ({map: f => v => r => s => dictFunctor.map(v1 => $RWSResult(v1._1, f(v1._2), v1._3))(v(r)(s))});
const execRWST = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return v => r => s => Bind1.bind(v(r)(s))(v1 => Applicative0.pure(Data$dTuple.$Tuple(v1._1, v1._3)));
};
const evalRWST = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return v => r => s => Bind1.bind(v(r)(s))(v1 => Applicative0.pure(Data$dTuple.$Tuple(v1._2, v1._3)));
};
const applyRWST = dictBind => {
  const Apply0 = dictBind.Apply0();
  const Functor0 = Apply0.Functor0();
  const $0 = Apply0.Functor0();
  const functorRWST1 = {map: f => v => r => s => $0.map(v1 => $RWSResult(v1._1, f(v1._2), v1._3))(v(r)(s))};
  return dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    return {
      apply: v => v1 => r => s => dictBind.bind(v(r)(s))(v2 => {
        const $1 = v2._3;
        return Functor0.map(v3 => $RWSResult(v3._1, v2._2(v3._2), Semigroup0.append($1)(v3._3)))(v1(r)(v2._1));
      }),
      Functor0: () => functorRWST1
    };
  };
};
const bindRWST = dictBind => {
  const Functor0 = dictBind.Apply0().Functor0();
  const applyRWST1 = applyRWST(dictBind);
  return dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    const applyRWST2 = applyRWST1(dictMonoid);
    return {
      bind: v => f => r => s => dictBind.bind(v(r)(s))(v1 => {
        const $0 = v1._3;
        return Functor0.map(v3 => $RWSResult(v3._1, v3._2, Semigroup0.append($0)(v3._3)))(f(v1._2)(r)(v1._1));
      }),
      Apply0: () => applyRWST2
    };
  };
};
const semigroupRWST = dictBind => {
  const applyRWST1 = applyRWST(dictBind);
  return dictMonoid => {
    const applyRWST2 = applyRWST1(dictMonoid);
    return dictSemigroup => (
      {
        append: (() => {
          const Functor0 = applyRWST2.Functor0();
          const $0 = dictSemigroup.append;
          return a => b => applyRWST2.apply(Functor0.map($0)(a))(b);
        })()
      }
    );
  };
};
const applicativeRWST = dictMonad => {
  const applyRWST1 = applyRWST(dictMonad.Bind1());
  return dictMonoid => {
    const applyRWST2 = applyRWST1(dictMonoid);
    return {pure: a => v => s => dictMonad.Applicative0().pure($RWSResult(s, a, dictMonoid.mempty)), Apply0: () => applyRWST2};
  };
};
const monadRWST = dictMonad => {
  const applicativeRWST1 = applicativeRWST(dictMonad);
  const bindRWST1 = bindRWST(dictMonad.Bind1());
  return dictMonoid => {
    const applicativeRWST2 = applicativeRWST1(dictMonoid);
    const bindRWST2 = bindRWST1(dictMonoid);
    return {Applicative0: () => applicativeRWST2, Bind1: () => bindRWST2};
  };
};
const monadAskRWST = dictMonad => {
  const monadRWST1 = monadRWST(dictMonad);
  return dictMonoid => {
    const monadRWST2 = monadRWST1(dictMonoid);
    return {ask: r => s => dictMonad.Applicative0().pure($RWSResult(s, r, dictMonoid.mempty)), Monad0: () => monadRWST2};
  };
};
const monadReaderRWST = dictMonad => {
  const monadAskRWST1 = monadAskRWST(dictMonad);
  return dictMonoid => {
    const monadAskRWST2 = monadAskRWST1(dictMonoid);
    return {local: f => m => r => s => m(f(r))(s), MonadAsk0: () => monadAskRWST2};
  };
};
const monadEffectRWS = dictMonoid => dictMonadEffect => {
  const Monad0 = dictMonadEffect.Monad0();
  const monadRWST1 = monadRWST(Monad0)(dictMonoid);
  return {
    liftEffect: (() => {
      const Bind1 = Monad0.Bind1();
      return x => {
        const $0 = dictMonadEffect.liftEffect(x);
        return v => s => Bind1.bind($0)(a => Monad0.Applicative0().pure($RWSResult(s, a, dictMonoid.mempty)));
      };
    })(),
    Monad0: () => monadRWST1
  };
};
const monadRecRWST = dictMonadRec => {
  const Monad0 = dictMonadRec.Monad0();
  const Bind1 = Monad0.Bind1();
  const Applicative0 = Monad0.Applicative0();
  const monadRWST1 = monadRWST(Monad0);
  return dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    const monadRWST2 = monadRWST1(dictMonoid);
    return {
      tailRecM: k => a => r => s => dictMonadRec.tailRecM(v => {
        const $0 = v._3;
        return Bind1.bind(k(v._2)(r)(v._1))(v2 => Applicative0.pure((() => {
          if (v2._2.tag === "Loop") { return Control$dMonad$dRec$dClass.$Step("Loop", $RWSResult(v2._1, v2._2._1, Semigroup0.append($0)(v2._3))); }
          if (v2._2.tag === "Done") { return Control$dMonad$dRec$dClass.$Step("Done", $RWSResult(v2._1, v2._2._1, Semigroup0.append($0)(v2._3))); }
          $runtime.fail();
        })()));
      })($RWSResult(s, a, dictMonoid.mempty)),
      Monad0: () => monadRWST2
    };
  };
};
const monadStateRWST = dictMonad => {
  const monadRWST1 = monadRWST(dictMonad);
  return dictMonoid => {
    const monadRWST2 = monadRWST1(dictMonoid);
    return {
      state: f => v => s => {
        const v1 = f(s);
        return dictMonad.Applicative0().pure($RWSResult(v1._2, v1._1, dictMonoid.mempty));
      },
      Monad0: () => monadRWST2
    };
  };
};
const monadTellRWST = dictMonad => {
  const monadRWST1 = monadRWST(dictMonad);
  return dictMonoid => {
    const Semigroup0 = dictMonoid.Semigroup0();
    const monadRWST2 = monadRWST1(dictMonoid);
    return {tell: w => v => s => dictMonad.Applicative0().pure($RWSResult(s, undefined, w)), Semigroup0: () => Semigroup0, Monad1: () => monadRWST2};
  };
};
const monadWriterRWST = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  const monadTellRWST1 = monadTellRWST(dictMonad);
  return dictMonoid => {
    const monadTellRWST2 = monadTellRWST1(dictMonoid);
    return {
      listen: m => r => s => Bind1.bind(m(r)(s))(v => Applicative0.pure($RWSResult(v._1, Data$dTuple.$Tuple(v._2, v._3), v._3))),
      pass: m => r => s => Bind1.bind(m(r)(s))(v => Applicative0.pure($RWSResult(v._1, v._2._1, v._2._2(v._3)))),
      Monoid0: () => dictMonoid,
      MonadTell1: () => monadTellRWST2
    };
  };
};
const monadThrowRWST = dictMonadThrow => {
  const Monad0 = dictMonadThrow.Monad0();
  const monadRWST1 = monadRWST(dictMonadThrow.Monad0());
  return dictMonoid => {
    const monadRWST2 = monadRWST1(dictMonoid);
    return {
      throwError: e => {
        const Bind1 = Monad0.Bind1();
        const $0 = dictMonadThrow.throwError(e);
        return v => s => Bind1.bind($0)(a => Monad0.Applicative0().pure($RWSResult(s, a, dictMonoid.mempty)));
      },
      Monad0: () => monadRWST2
    };
  };
};
const monadErrorRWST = dictMonadError => {
  const monadThrowRWST1 = monadThrowRWST(dictMonadError.MonadThrow0());
  return dictMonoid => {
    const monadThrowRWST2 = monadThrowRWST1(dictMonoid);
    return {catchError: m => h => r => s => dictMonadError.catchError(m(r)(s))(e => h(e)(r)(s)), MonadThrow0: () => monadThrowRWST2};
  };
};
const monadSTRWST = dictMonoid => dictMonadST => {
  const Monad0 = dictMonadST.Monad0();
  const monadRWST1 = monadRWST(Monad0)(dictMonoid);
  return {
    liftST: (() => {
      const Bind1 = Monad0.Bind1();
      return x => {
        const $0 = dictMonadST.liftST(x);
        return v => s => Bind1.bind($0)(a => Monad0.Applicative0().pure($RWSResult(s, a, dictMonoid.mempty)));
      };
    })(),
    Monad0: () => monadRWST1
  };
};
const monoidRWST = dictMonad => {
  const applicativeRWST1 = applicativeRWST(dictMonad);
  const semigroupRWST1 = semigroupRWST(dictMonad.Bind1());
  return dictMonoid => {
    const applicativeRWST2 = applicativeRWST1(dictMonoid);
    const semigroupRWST2 = semigroupRWST1(dictMonoid);
    return dictMonoid1 => {
      const semigroupRWST3 = semigroupRWST2(dictMonoid1.Semigroup0());
      return {mempty: applicativeRWST2.pure(dictMonoid1.mempty), Semigroup0: () => semigroupRWST3};
    };
  };
};
const altRWST = dictAlt => {
  const $0 = dictAlt.Functor0();
  const functorRWST1 = {map: f => v => r => s => $0.map(v1 => $RWSResult(v1._1, f(v1._2), v1._3))(v(r)(s))};
  return {alt: v => v1 => r => s => dictAlt.alt(v(r)(s))(v1(r)(s)), Functor0: () => functorRWST1};
};
const plusRWST = dictPlus => {
  const empty = dictPlus.empty;
  const $0 = dictPlus.Alt0();
  const $1 = $0.Functor0();
  const altRWST1 = (() => {
    const functorRWST1 = {map: f => v => r => s => $1.map(v1 => $RWSResult(v1._1, f(v1._2), v1._3))(v(r)(s))};
    return {alt: v => v1 => r => s => $0.alt(v(r)(s))(v1(r)(s)), Functor0: () => functorRWST1};
  })();
  return {empty: v => v1 => empty, Alt0: () => altRWST1};
};
const alternativeRWST = dictMonoid => dictAlternative => {
  const $0 = dictAlternative.Plus1();
  const empty = $0.empty;
  const plusRWST1 = (() => {
    const $1 = $0.Alt0();
    const $2 = $1.Functor0();
    const functorRWST1 = {map: f => v => r => s => $2.map(v1 => $RWSResult(v1._1, f(v1._2), v1._3))(v(r)(s))};
    const altRWST1 = {alt: v => v1 => r => s => $1.alt(v(r)(s))(v1(r)(s)), Functor0: () => functorRWST1};
    return {empty: v => v1 => empty, Alt0: () => altRWST1};
  })();
  return dictMonad => {
    const applicativeRWST1 = applicativeRWST(dictMonad)(dictMonoid);
    return {Applicative0: () => applicativeRWST1, Plus1: () => plusRWST1};
  };
};
export {
  $RWSResult,
  RWSResult,
  RWST,
  altRWST,
  alternativeRWST,
  applicativeRWST,
  applyRWST,
  bindRWST,
  evalRWST,
  execRWST,
  functorRWST,
  lazyRWST,
  mapRWST,
  monadAskRWST,
  monadEffectRWS,
  monadErrorRWST,
  monadRWST,
  monadReaderRWST,
  monadRecRWST,
  monadSTRWST,
  monadStateRWST,
  monadTellRWST,
  monadThrowRWST,
  monadTransRWST,
  monadWriterRWST,
  monoidRWST,
  newtypeRWST,
  plusRWST,
  runRWST,
  semigroupRWST,
  withRWST
};
