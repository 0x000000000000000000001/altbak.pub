const fs = require('fs');
const file = '/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');

content = content.replace(
`         Just ctorTy -> fromMaybe Any (Array.index (extractAllArgTypes ctorTy) fieldIdx)`,
`         Just ctorTy -> 
           let args = extractAllArgTypes ctorTy
           in case Array.index args fieldIdx of
             Just t -> t
             Nothing -> Debug.trace ("Warning: fieldIdx " <> show fieldIdx <> " out of bounds for " <> ctorFqn <> " (args len: " <> show (Array.length args) <> ")") \\_ -> Any`);

fs.writeFileSync(file, content);
