const fs = require('fs');
const path = '../phpurs/phpurs/src/Phpurs/Printer.purs';
let content = fs.readFileSync(path, 'utf8');

const target = `"  }\n" <>\n      "}\\n"\n`;
const replacement = `"  }\n" <>\n      "}\\n" <>\n      "if (!\\\\function_exists(__NAMESPACE__ . '\\\\\\\\phpurs_execute_effect')) {\\n" <>\n      "  function phpurs_execute_effect($val) {\\n" <>\n      "    if (\\\\is_callable($val)) {\\n" <>\n      "      return $val($GLOBALS['Data_Unit_unit']);\\n" <>\n      "    }\\n" <>\n      "    return $val;\\n" <>\n      "  }\\n" <>\n      "}\\n"\n`;

content = content.replace(target, replacement);
fs.writeFileSync(path, content);
