const fs = require('fs');
const path = '/Users/0x1/Documents/htdocs/purust/purust/src/Main.purs';
let content = fs.readFileSync(path, 'utf8');

content = content.replace(
  'FS.writeTextFile UTF8 (modDir <> "/src/lib.rs") v',
  `let transImps = fromMaybe [] (map (\\s -> Set.toUnfoldable s :: Array String) (Map.lookup k finalTcMap))
      let newImportsRust = String.joinWith "\\n" (map (\\i -> "use Purs_" <> i <> "::*;") transImps)
      let finalCode = String.replace (Pattern "use purust_core::*;\\n") (Replacement ("use purust_core::*;\\n" <> newImportsRust <> "\\n")) v
      FS.writeTextFile UTF8 (modDir <> "/src/lib.rs") finalCode`
);

fs.writeFileSync(path, content);
