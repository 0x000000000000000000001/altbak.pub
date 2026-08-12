import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Type$dProxy from "../Type.Proxy/index.js";
import {arrayApply} from "./foreign.js";
const applyProxy = {apply: v => v1 => Type$dProxy.Proxy, Functor0: () => Data$dFunctor.functorProxy};
const applyFn = {apply: f => g => x => f(x)(g(x)), Functor0: () => Data$dFunctor.functorFn};
const applyArray = {apply: arrayApply, Functor0: () => Data$dFunctor.functorArray};
const apply = dict => dict.apply;
const applyFirst = dictApply => {
  const Functor0 = dictApply.Functor0();
  return a => b => dictApply.apply(Functor0.map(Data$dFunction.const)(a))(b);
};
const applySecond = dictApply => {
  const Functor0 = dictApply.Functor0();
  return a => b => dictApply.apply(Functor0.map(v => x => x)(a))(b);
};
const lift2 = dictApply => {
  const Functor0 = dictApply.Functor0();
  return f => a => b => dictApply.apply(Functor0.map(f)(a))(b);
};
const lift3 = dictApply => {
  const Functor0 = dictApply.Functor0();
  return f => a => b => c => dictApply.apply(dictApply.apply(Functor0.map(f)(a))(b))(c);
};
const lift4 = dictApply => {
  const Functor0 = dictApply.Functor0();
  return f => a => b => c => d => dictApply.apply(dictApply.apply(dictApply.apply(Functor0.map(f)(a))(b))(c))(d);
};
const lift5 = dictApply => {
  const Functor0 = dictApply.Functor0();
  return f => a => b => c => d => e => dictApply.apply(dictApply.apply(dictApply.apply(dictApply.apply(Functor0.map(f)(a))(b))(c))(d))(e);
};
export {apply, applyArray, applyFirst, applyFn, applyProxy, applySecond, lift2, lift3, lift4, lift5};
export * from "./foreign.js";
