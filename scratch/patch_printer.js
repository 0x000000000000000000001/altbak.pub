const fs = require('fs');

const path = '../phpurs/phpurs/src/Phpurs/Printer.purs';
let content = fs.readFileSync(path, 'utf8');

const executeEffect = `    }
    return phpurs_curry_fallback($fn, $merged, $expected);
  };
}
if (!\\function_exists(__NAMESPACE__ . '\\\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}
`;

content = content.replace(/    return phpurs_curry_fallback\(\$fn, \$merged, \$expected\);\n    };\n  }\n}\n"/g, executeEffect + '}\n"');

fs.writeFileSync(path, content);
