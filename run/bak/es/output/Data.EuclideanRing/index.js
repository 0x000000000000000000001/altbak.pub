import * as Data$dCommutativeRing from "../Data.CommutativeRing/index.js";
import {intDegree, intDiv, intMod, numDiv} from "./foreign.js";
const mod = dict => dict.mod;
const gcd = dictEq => dictEuclideanRing => {
  const zero = dictEuclideanRing.CommutativeRing0().Ring0().Semiring0().zero;
  return a => b => {
    if (dictEq.eq(b)(zero)) { return a; }
    return gcd(dictEq)(dictEuclideanRing)(b)(dictEuclideanRing.mod(a)(b));
  };
};
const euclideanRingNumber = {degree: v => 1, div: numDiv, mod: v => v1 => 0.0, CommutativeRing0: () => Data$dCommutativeRing.commutativeRingNumber};
const euclideanRingInt = {degree: intDegree, div: intDiv, mod: intMod, CommutativeRing0: () => Data$dCommutativeRing.commutativeRingInt};
const div = dict => dict.div;
const lcm = dictEq => dictEuclideanRing => {
  const Ring0 = dictEuclideanRing.CommutativeRing0().Ring0();
  const zero = Ring0.Semiring0().zero;
  const Semiring0 = Ring0.Semiring0();
  return a => b => {
    if (dictEq.eq(a)(zero) || dictEq.eq(b)(zero)) { return zero; }
    return dictEuclideanRing.div(Semiring0.mul(a)(b))(dictEq.eq(b)(dictEuclideanRing.CommutativeRing0().Ring0().Semiring0().zero)
      ? a
      : gcd(dictEq)(dictEuclideanRing)(b)(dictEuclideanRing.mod(a)(b)));
  };
};
const degree = dict => dict.degree;
export {degree, div, euclideanRingInt, euclideanRingNumber, gcd, lcm, mod};
export * from "./foreign.js";
