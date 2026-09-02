const fs = require('fs');
let code = fs.readFileSync('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                      expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr (getExprType val)`,
  `                      expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr (Debug.trace { msg: "LetRec val type for " <> ident, valType: getExprType val } \\_ -> getExprType val)`
);
fs.writeFileSync('../gopurs/gopurs/src/Gopurs/CodeGen.purs', code);
