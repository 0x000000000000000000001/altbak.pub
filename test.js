const fs = require('fs');
let content = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
content = content.replace(/Branch branches def -> fromMaybe \(getExprType def\) \(Array\.find \(\\t -> t \/= Any\) \(map \(\\\(Pair _ b\\) -> getExprType b\) \(toArray branches\)\)\)/g, `Branch branches def -> fromMaybe (getExprType def) (Array.find (\\t -> t /= Any) (map (\\(Pair _ b) -> getExprType b) (toArray branches)))\n      Let _ _ body -> getExprType body\n      LetRec _ body -> getExprType body`);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', content);
