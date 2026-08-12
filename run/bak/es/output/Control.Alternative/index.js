import * as Control$dApplicative from "../Control.Applicative/index.js";
import * as Control$dPlus from "../Control.Plus/index.js";
const guard = dictAlternative => {
  const Applicative0 = dictAlternative.Applicative0();
  const empty = dictAlternative.Plus1().empty;
  return v => {
    if (v) { return Applicative0.pure(); }
    return empty;
  };
};
const alternativeArray = {Applicative0: () => Control$dApplicative.applicativeArray, Plus1: () => Control$dPlus.plusArray};
export {alternativeArray, guard};
