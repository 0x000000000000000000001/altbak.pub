const fs = require('fs');
const path = '/Users/0x1/Documents/htdocs/purust/purust/src/Purust/ASTCollector.purs';
let content = fs.readFileSync(path, 'utf8');

content = content.replace(
  'ExprUpdate _ expr props -> Set.union (collectModulesExpr expr) (Array.foldl (\\acc (Prop _ v) -> Set.union acc (collectModulesExpr v)) Set.empty props)',
  'ExprUpdate _ expr props -> Set.insert "Record.Unsafe" (Set.union (collectModulesExpr expr) (Array.foldl (\\acc (Prop _ v) -> Set.union acc (collectModulesExpr v)) Set.empty props))'
);

content = content.replace(
  'ExprAccessor _ expr _ -> collectModulesExpr expr',
  'ExprAccessor _ expr _ -> Set.insert "Record.Unsafe" (collectModulesExpr expr)'
);

fs.writeFileSync(path, content);
