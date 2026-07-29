const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `          mbTopLevelInfo = case mbIdent of
            Just ident -> Map.lookup (modNameStr <> "." <> ident) moduleArities`,
  `          mbTopLevelInfo = case mbIdent of
            Just ident -> 
              let 
                key = modNameStr <> "." <> ident
                val = Map.lookup key moduleArities
                _ = if ident == "makeBlack" then Debug.trace ("makeBlack top level info: " <> show (isJust val)) (\\_ -> unit) else unit
                _ = if ident == "makeBlack" && isJust val then Debug.trace ("makeBlack args: " <> String.joinWith ", " (map goTypeToStr (fromMaybe { fullName: "", fArgs: [], fRet: TypeValue, arity: 0 } val).fArgs)) (\\_ -> unit) else unit
              in val`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
