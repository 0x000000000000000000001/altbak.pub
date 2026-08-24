const fs = require('fs');
const path = '../phpurs/phpurs/src/Phpurs/CodeGen.purs';
let content = fs.readFileSync(path, 'utf8');

content = content.replace(
  'PhpRaw "\\\\\\\\phpurs_execute_effect"',
  'PhpRaw "phpurs_execute_effect"'
);

fs.writeFileSync(path, content);
