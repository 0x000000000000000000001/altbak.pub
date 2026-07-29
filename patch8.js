const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `      Var originalName ->`,
  `      Var originalName ->
        let
          _ = if originalName == "v_0" then Debug.trace ("Var v_0 in bound? " <> show (Map.member "v_0" bound)) (\\_ -> unit) else unit
        in`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
