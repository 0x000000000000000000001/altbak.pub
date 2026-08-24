const fs = require('fs');
const path = '../phpurs/phpurs/src/Phpurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

content = content.replace(
  'executeIfOpaque expr phpExpr =\n  if isEffectNode expr then phpExpr\n  else PhpCall phpExpr [ PhpRaw "$GLOBALS[\'Data_Unit_unit\']" ]',
  'executeIfOpaque expr phpExpr =\n  if isEffectNode expr then phpExpr\n  else PhpCall (PhpRaw "phpurs_execute_effect") [ phpExpr ]'
);

fs.writeFileSync(path, content);
