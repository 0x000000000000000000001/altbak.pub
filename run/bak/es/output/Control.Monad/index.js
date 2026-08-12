import * as $runtime from "../runtime.js";
import * as Control$dApplicative from "../Control.Applicative/index.js";
import * as Control$dBind from "../Control.Bind/index.js";
const whenM = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return mb => m => Bind1.bind(mb)(b => {
    if (b) { return m; }
    return Applicative0.pure();
  });
};
const unlessM = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return mb => m => Bind1.bind(mb)(b => {
    if (!b) { return m; }
    if (b) { return Applicative0.pure(); }
    $runtime.fail();
  });
};
const monadProxy = {Applicative0: () => Control$dApplicative.applicativeProxy, Bind1: () => Control$dBind.bindProxy};
const monadFn = {Applicative0: () => Control$dApplicative.applicativeFn, Bind1: () => Control$dBind.bindFn};
const monadArray = {Applicative0: () => Control$dApplicative.applicativeArray, Bind1: () => Control$dBind.bindArray};
const liftM1 = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return f => a => Bind1.bind(a)(a$p => Applicative0.pure(f(a$p)));
};
const ap = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  return f => a => Bind1.bind(f)(f$p => Bind1.bind(a)(a$p => Applicative0.pure(f$p(a$p))));
};
export {ap, liftM1, monadArray, monadFn, monadProxy, unlessM, whenM};
