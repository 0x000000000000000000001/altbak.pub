const fs = require('fs');
const file = '/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');

// Replace signatures
content = content.replace('unboxGoExpr :: GoExpr -> GoType -> GoType -> GoExpr', 'unboxGoExpr :: String -> GoExpr -> GoType -> GoType -> GoExpr');

// Replace unboxGoExpr definitions (there are a few clauses)
content = content.replace('unboxGoExpr expr currentType desiredType =', 'unboxGoExpr modNameStr expr currentType desiredType =');
// Inside unboxGoExpr, there are recursive calls to unboxGoExpr and boxGoExpr. We need to replace them carefully.
// It's better to replace globally and then fix the definition of boxGoExprImpl.
content = content.replace(/\bunboxGoExpr\b(?!\s*::)/g, 'unboxGoExpr modNameStr');
content = content.replace(/\bboxGoExpr\b(?!\s*::)/g, 'boxGoExpr modNameStr');

// Wait, the global replacement will also replace `boxGoExpr modNameStr` inside its OWN definition if we're not careful.
