const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                _ = if ident == "makeBlack" then Debug.trace ("makeBlack top level info: " <> show (isJust val)) (\\_ -> unit) else unit
                _ = if ident == "makeBlack" && isJust val then Debug.trace ("makeBlack args: " <> String.joinWith ", " (map goTypeToStr (fromMaybe { fullName: "", fArgs: [], fRet: TypeValue, arity: 0 } val).fArgs)) (\\_ -> unit) else unit
              in val`,
  `                _ignore1 = if ident == "makeBlack" then Debug.trace ("makeBlack top level info: " <> show (isJust val)) (\\_ -> unit) else unit
                _ignore2 = if ident == "makeBlack" && isJust val then Debug.trace ("makeBlack args: " <> String.joinWith ", " (map goTypeToStr (fromMaybe { fullName: "", fArgs: [], fRet: TypeValue, arity: 0 } val).fArgs)) (\\_ -> unit) else unit
              in const (const val _ignore1) _ignore2`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
