const fs = require('fs');
const content = fs.readFileSync('src/Purust/CodeGen.purs', 'utf-8');
const newContent = content.replace(
    'locals = unsafePerformEffect (Ref.read globalLocals)',
    'locals = unsafePerformEffect do\n              ls <- Ref.read globalLocals\n              pure $ Debug.trace ("CtorSaturated " <> ctorName <> " locals=" <> show (Array.fromFoldable ls :: Array String)) \\_ -> ls'
);
fs.writeFileSync('src/Purust/CodeGen.purs', newContent);
