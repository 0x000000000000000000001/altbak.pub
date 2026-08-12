import * as $runtime from "../runtime.js";
import * as Control$dMonad$dGen from "../Control.Monad.Gen/index.js";
import * as Data$dChar$dGen from "../Data.Char.Gen/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
const genString = dictMonadRec => dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const Bind1 = Monad0.Bind1();
  const Functor0 = Monad0.Bind1().Apply0().Functor0();
  return genChar => dictMonadGen.sized(size => Bind1.bind(dictMonadGen.chooseInt(1)((() => {
    const v = Data$dOrd.ordInt.compare(1)(size);
    if (v === "LT") { return size; }
    if (v === "EQ") { return 1; }
    if (v === "GT") { return 1; }
    $runtime.fail();
  })()))(newSize => dictMonadGen.resize(v => newSize)(Functor0.map(Data$dString$dCodeUnits.fromCharArray)(Control$dMonad$dGen.unfoldable(dictMonadRec)(dictMonadGen)(Data$dUnfoldable.unfoldableArray)(genChar)))));
};
const genUnicodeString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genUnicodeChar(dictMonadGen));
const genDigitString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genDigitChar(dictMonadGen));
const genAsciiString$p = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genAsciiChar$p(dictMonadGen));
const genAsciiString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genAsciiChar(dictMonadGen));
const genAlphaUppercaseString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genAlphaUppercase(dictMonadGen));
const genAlphaString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genAlpha(dictMonadGen));
const genAlphaLowercaseString = dictMonadRec => dictMonadGen => genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genAlphaLowercase(dictMonadGen));
export {genAlphaLowercaseString, genAlphaString, genAlphaUppercaseString, genAsciiString, genAsciiString$p, genDigitString, genString, genUnicodeString};
