const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                  let _ = Debug.trace ("FUNC: " <> name <> " ARGS: " <> String.joinWith ", " (map goTypeToStr fArgsGo)) (\\_ -> unit)
                  in [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]`,
  `                  let _ignore = Debug.trace ("FUNC: " <> name <> " ARGS: " <> String.joinWith ", " (map goTypeToStr fArgsGo)) (\\_ -> unit)
                  in [ Tuple (sanitizeName name) (const { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } _ignore) ]`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
