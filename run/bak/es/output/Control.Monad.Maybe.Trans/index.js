// | This module defines the `MaybeT` monad transformer.
import * as $runtime from "../runtime.js";
import * as Control$dMonad$dRec$dClass from "../Control.Monad.Rec.Class/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const MaybeT = x => x;
const runMaybeT = v => v;
const newtypeMaybeT = {Coercible0: () => {}};
const monadTransMaybeT = {
  lift: dictMonad => {
    const Bind1 = dictMonad.Bind1();
    const Applicative0 = dictMonad.Applicative0();
    return x => Bind1.bind(x)(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
  }
};
const mapMaybeT = f => v => f(v);
const functorMaybeT = dictFunctor => (
  {
    map: f => v => dictFunctor.map(v1 => {
      if (v1.tag === "Just") { return Data$dMaybe.$Maybe("Just", f(v1._1)); }
      return Data$dMaybe.Nothing;
    })(v)
  }
);
const monadMaybeT = dictMonad => ({Applicative0: () => applicativeMaybeT(dictMonad), Bind1: () => bindMaybeT(dictMonad)});
const bindMaybeT = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return {
    bind: v => f => Bind1.bind(v)(v1 => {
      if (v1.tag === "Nothing") { return Applicative0.pure(Data$dMaybe.Nothing); }
      if (v1.tag === "Just") { return f(v1._1); }
      $runtime.fail();
    }),
    Apply0: () => applyMaybeT(dictMonad)
  };
};
const applyMaybeT = dictMonad => {
  const $0 = dictMonad.Bind1().Apply0().Functor0();
  const functorMaybeT1 = {
    map: f => v => $0.map(v1 => {
      if (v1.tag === "Just") { return Data$dMaybe.$Maybe("Just", f(v1._1)); }
      return Data$dMaybe.Nothing;
    })(v)
  };
  return {
    apply: (() => {
      const Bind1 = bindMaybeT(dictMonad);
      const Applicative0 = applicativeMaybeT(dictMonad);
      return f => a => Bind1.bind(f)(f$p => Bind1.bind(a)(a$p => Applicative0.pure(f$p(a$p))));
    })(),
    Functor0: () => functorMaybeT1
  };
};
const applicativeMaybeT = dictMonad => ({pure: x => dictMonad.Applicative0().pure(Data$dMaybe.$Maybe("Just", x)), Apply0: () => applyMaybeT(dictMonad)});
const semigroupMaybeT = dictMonad => {
  const applyMaybeT1 = applyMaybeT(dictMonad);
  return dictSemigroup => (
    {
      append: (() => {
        const Functor0 = applyMaybeT1.Functor0();
        const $0 = dictSemigroup.append;
        return a => b => applyMaybeT1.apply(Functor0.map($0)(a))(b);
      })()
    }
  );
};
const monadAskMaybeT = dictMonadAsk => {
  const $0 = dictMonadAsk.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT($0), Bind1: () => bindMaybeT($0)};
  return {
    ask: (() => {
      const $1 = dictMonadAsk.Monad0();
      const Applicative0 = $1.Applicative0();
      return $1.Bind1().bind(dictMonadAsk.ask)(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    })(),
    Monad0: () => monadMaybeT1
  };
};
const monadReaderMaybeT = dictMonadReader => {
  const monadAskMaybeT1 = monadAskMaybeT(dictMonadReader.MonadAsk0());
  return {local: f => dictMonadReader.local(f), MonadAsk0: () => monadAskMaybeT1};
};
const monadContMaybeT = dictMonadCont => {
  const $0 = dictMonadCont.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT($0), Bind1: () => bindMaybeT($0)};
  return {callCC: f => dictMonadCont.callCC(c => f(a => c(Data$dMaybe.$Maybe("Just", a)))), Monad0: () => monadMaybeT1};
};
const monadEffectMaybe = dictMonadEffect => {
  const Monad0 = dictMonadEffect.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT(Monad0), Bind1: () => bindMaybeT(Monad0)};
  return {
    liftEffect: (() => {
      const Bind1 = Monad0.Bind1();
      const Applicative0 = Monad0.Applicative0();
      return x => Bind1.bind(dictMonadEffect.liftEffect(x))(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    })(),
    Monad0: () => monadMaybeT1
  };
};
const monadRecMaybeT = dictMonadRec => {
  const Monad0 = dictMonadRec.Monad0();
  const Bind1 = Monad0.Bind1();
  const Applicative0 = Monad0.Applicative0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT(Monad0), Bind1: () => bindMaybeT(Monad0)};
  return {
    tailRecM: f => dictMonadRec.tailRecM(a => Bind1.bind(f(a))(m$p => Applicative0.pure((() => {
      if (m$p.tag === "Nothing") { return Control$dMonad$dRec$dClass.$Step("Done", Data$dMaybe.Nothing); }
      if (m$p.tag === "Just") {
        if (m$p._1.tag === "Loop") { return Control$dMonad$dRec$dClass.$Step("Loop", m$p._1._1); }
        if (m$p._1.tag === "Done") { return Control$dMonad$dRec$dClass.$Step("Done", Data$dMaybe.$Maybe("Just", m$p._1._1)); }
      }
      $runtime.fail();
    })()))),
    Monad0: () => monadMaybeT1
  };
};
const monadStateMaybeT = dictMonadState => {
  const Monad0 = dictMonadState.Monad0();
  const $0 = dictMonadState.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT($0), Bind1: () => bindMaybeT($0)};
  return {
    state: f => {
      const Applicative0 = Monad0.Applicative0();
      return Monad0.Bind1().bind(dictMonadState.state(f))(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    },
    Monad0: () => monadMaybeT1
  };
};
const monadTellMaybeT = dictMonadTell => {
  const Monad1 = dictMonadTell.Monad1();
  const Semigroup0 = dictMonadTell.Semigroup0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT(Monad1), Bind1: () => bindMaybeT(Monad1)};
  return {
    tell: (() => {
      const Bind1 = Monad1.Bind1();
      const Applicative0 = Monad1.Applicative0();
      return x => Bind1.bind(dictMonadTell.tell(x))(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    })(),
    Semigroup0: () => Semigroup0,
    Monad1: () => monadMaybeT1
  };
};
const monadWriterMaybeT = dictMonadWriter => {
  const MonadTell1 = dictMonadWriter.MonadTell1();
  const Monad1 = MonadTell1.Monad1();
  const Bind1 = Monad1.Bind1();
  const Applicative0 = Monad1.Applicative0();
  const Monoid0 = dictMonadWriter.Monoid0();
  const monadTellMaybeT1 = monadTellMaybeT(MonadTell1);
  return {
    listen: v => Bind1.bind(dictMonadWriter.listen(v))(v$1 => Monad1.Applicative0().pure(v$1._1.tag === "Just"
      ? Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v$1._1._1, v$1._2))
      : Data$dMaybe.Nothing)),
    pass: v => dictMonadWriter.pass(Bind1.bind(v)(a => Applicative0.pure((() => {
      if (a.tag === "Nothing") { return Data$dTuple.$Tuple(Data$dMaybe.Nothing, x => x); }
      if (a.tag === "Just") { return Data$dTuple.$Tuple(Data$dMaybe.$Maybe("Just", a._1._1), a._1._2); }
      $runtime.fail();
    })()))),
    Monoid0: () => Monoid0,
    MonadTell1: () => monadTellMaybeT1
  };
};
const monadThrowMaybeT = dictMonadThrow => {
  const Monad0 = dictMonadThrow.Monad0();
  const $0 = dictMonadThrow.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT($0), Bind1: () => bindMaybeT($0)};
  return {
    throwError: e => {
      const Applicative0 = Monad0.Applicative0();
      return Monad0.Bind1().bind(dictMonadThrow.throwError(e))(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    },
    Monad0: () => monadMaybeT1
  };
};
const monadErrorMaybeT = dictMonadError => {
  const monadThrowMaybeT1 = monadThrowMaybeT(dictMonadError.MonadThrow0());
  return {catchError: v => h => dictMonadError.catchError(v)(a => h(a)), MonadThrow0: () => monadThrowMaybeT1};
};
const monadSTMaybeT = dictMonadST => {
  const Monad0 = dictMonadST.Monad0();
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT(Monad0), Bind1: () => bindMaybeT(Monad0)};
  return {
    liftST: (() => {
      const Bind1 = Monad0.Bind1();
      const Applicative0 = Monad0.Applicative0();
      return x => Bind1.bind(dictMonadST.liftST(x))(a$p => Applicative0.pure(Data$dMaybe.$Maybe("Just", a$p)));
    })(),
    Monad0: () => monadMaybeT1
  };
};
const monoidMaybeT = dictMonad => {
  const applicativeMaybeT1 = applicativeMaybeT(dictMonad);
  const semigroupMaybeT1 = semigroupMaybeT(dictMonad);
  return dictMonoid => {
    const semigroupMaybeT2 = semigroupMaybeT1(dictMonoid.Semigroup0());
    return {mempty: applicativeMaybeT1.pure(dictMonoid.mempty), Semigroup0: () => semigroupMaybeT2};
  };
};
const altMaybeT = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  const $0 = dictMonad.Bind1().Apply0().Functor0();
  const functorMaybeT1 = {
    map: f => v => $0.map(v1 => {
      if (v1.tag === "Just") { return Data$dMaybe.$Maybe("Just", f(v1._1)); }
      return Data$dMaybe.Nothing;
    })(v)
  };
  return {
    alt: v => v1 => Bind1.bind(v)(m => {
      if (m.tag === "Nothing") { return v1; }
      return Applicative0.pure(m);
    }),
    Functor0: () => functorMaybeT1
  };
};
const plusMaybeT = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  const altMaybeT1 = (() => {
    const $0 = dictMonad.Bind1().Apply0().Functor0();
    const functorMaybeT1 = {
      map: f => v => $0.map(v1 => {
        if (v1.tag === "Just") { return Data$dMaybe.$Maybe("Just", f(v1._1)); }
        return Data$dMaybe.Nothing;
      })(v)
    };
    return {
      alt: v => v1 => Bind1.bind(v)(m => {
        if (m.tag === "Nothing") { return v1; }
        return Applicative0.pure(m);
      }),
      Functor0: () => functorMaybeT1
    };
  })();
  return {empty: dictMonad.Applicative0().pure(Data$dMaybe.Nothing), Alt0: () => altMaybeT1};
};
const alternativeMaybeT = dictMonad => {
  const applicativeMaybeT1 = applicativeMaybeT(dictMonad);
  const plusMaybeT1 = plusMaybeT(dictMonad);
  return {Applicative0: () => applicativeMaybeT1, Plus1: () => plusMaybeT1};
};
const monadPlusMaybeT = dictMonad => {
  const monadMaybeT1 = {Applicative0: () => applicativeMaybeT(dictMonad), Bind1: () => bindMaybeT(dictMonad)};
  const alternativeMaybeT1 = alternativeMaybeT(dictMonad);
  return {Monad0: () => monadMaybeT1, Alternative1: () => alternativeMaybeT1};
};
export {
  MaybeT,
  altMaybeT,
  alternativeMaybeT,
  applicativeMaybeT,
  applyMaybeT,
  bindMaybeT,
  functorMaybeT,
  mapMaybeT,
  monadAskMaybeT,
  monadContMaybeT,
  monadEffectMaybe,
  monadErrorMaybeT,
  monadMaybeT,
  monadPlusMaybeT,
  monadReaderMaybeT,
  monadRecMaybeT,
  monadSTMaybeT,
  monadStateMaybeT,
  monadTellMaybeT,
  monadThrowMaybeT,
  monadTransMaybeT,
  monadWriterMaybeT,
  monoidMaybeT,
  newtypeMaybeT,
  plusMaybeT,
  runMaybeT,
  semigroupMaybeT
};
