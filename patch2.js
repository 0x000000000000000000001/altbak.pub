const fs = require('fs');
const file = '/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');

content = content.replace(
`    in case Map.lookup ctorFqn aritiesMap of
         Just ctorTy -> fromMaybe Any (Array.index (extractAllArgTypes ctorTy) fieldIdx)
         Nothing -> Any`,
`    in case Map.lookup ctorFqn aritiesMap of
         Just ctorTy -> fromMaybe Any (Array.index (extractAllArgTypes ctorTy) fieldIdx)
         Nothing -> Debug.trace ("Warning: ctorFqn not found in aritiesMap: " <> ctorFqn) \\_ -> Any`);

fs.writeFileSync(file, content);
