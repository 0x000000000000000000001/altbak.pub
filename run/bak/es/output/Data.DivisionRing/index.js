import * as Data$dRing from "../Data.Ring/index.js";
const recip = dict => dict.recip;
const rightDiv = dictDivisionRing => {
  const Semiring0 = dictDivisionRing.Ring0().Semiring0();
  return a => b => Semiring0.mul(a)(dictDivisionRing.recip(b));
};
const leftDiv = dictDivisionRing => {
  const Semiring0 = dictDivisionRing.Ring0().Semiring0();
  return a => b => Semiring0.mul(dictDivisionRing.recip(b))(a);
};
const divisionringNumber = {recip: x => 1.0 / x, Ring0: () => Data$dRing.ringNumber};
export {divisionringNumber, leftDiv, recip, rightDiv};
