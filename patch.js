const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  /_ -> GoConstructorAccess \(boxGoExpr resObj.expr resObj.exprType\) \(pkgPrefix <> monoStructName\) typeArgs idx/,
  '_ -> GoRaw ("/* expectedType: " <> goTypeToStr expectedType <> " resObj.exprType: " <> goTypeToStr resObj.exprType <> " */" <> printGoExpr (GoConstructorAccess (boxGoExpr resObj.expr resObj.exprType) (pkgPrefix <> monoStructName) typeArgs idx))'
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
